package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cli/go-gh/v2"
)

const (
	gitRepoName = "abekoh/go-ecr-deploy"
	ecrRepoName = "go-ecr-deploy"
)

func getJSONFields[T any]() string {
	var d T
	t := reflect.TypeOf(d)
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	fields := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("json")
		if field == "" {
			return ""
		}
		fields[i] = field
	}
	return strings.Join(fields, ",")
}

func runRepoGH(ctx context.Context, args ...string) error {
	commandArgs := append(args, "--repo", gitRepoName)
	log.Printf("[CMD] gh %s", strings.Join(commandArgs, " "))
	_, _, err := gh.ExecContext(ctx, commandArgs...)
	return err
}

func runRepoGHWithResponse[T any](ctx context.Context, args ...string) (T, error) {
	commandArgs := append(args, "--repo", gitRepoName, "--json", getJSONFields[T]())
	log.Printf("[CMD] gh %s", strings.Join(commandArgs, " "))
	var res T
	commandRes, _, err := gh.ExecContext(ctx, commandArgs...)
	if err != nil {
		return res, err
	}
	if err := json.Unmarshal(commandRes.Bytes(), &res); err != nil {
		return res, err
	}
	return res, nil
}

type RunListElement struct {
	Number       int       `json:"number"`
	WorkflowName string    `json:"workflowName"`
	StartedAt    time.Time `json:"startedAt"`
	HeadBranch   string    `json:"headBranch"`
	DatabaseID   int       `json:"databaseId"`
}

type (
	RunView struct {
		DatabaseID int               `json:"databaseId"`
		StartedAt  time.Time         `json:"startedAt"`
		UpdatedAt  time.Time         `json:"updatedAt"`
		Status     RunViewStatus     `json:"status"`
		Conclusion RunViewConclusion `json:"conclusion"`
	}
	RunViewStatus     string
	RunViewConclusion string
)

const (
	RunViewStatusCompleted RunViewStatus = "completed"
)

const (
	RunViewConclusionSuccess   RunViewConclusion = "success"
	RunViewConclusionFailure   RunViewConclusion = "failure"
	RunViewConclusionCancelled RunViewConclusion = "cancelled"
)

func runJobAndMeasure(ctx context.Context, targetJob, branch string) (time.Duration, error) {
	workflowStartedAt := time.Now()
	if err := runRepoGH(ctx, "workflow", "run", "deploy.yml", "-f", "targetJob="+targetJob, "--ref", branch); err != nil {
		return 0, fmt.Errorf("failed to run gh command: %w", err)
	}
	time.Sleep(5 * time.Second)

	var run RunListElement
	attempts := 0
	for {
		runs, err := runRepoGHWithResponse[[]RunListElement](ctx, "run", "list", "--limit", "1", "--json", getJSONFields[RunListElement]())
		if err != nil {
			return 0, fmt.Errorf("failed to run gh command: %w", err)
		}
		if len(runs) == 0 {
			continue
		}
		if runs[0].StartedAt.After(workflowStartedAt) {
			run = runs[0]
			break
		}
		if attempts >= 5 {
			return 0, fmt.Errorf("failed to find a run after %d attempts", attempts)
		}
		attempts++
		time.Sleep(5 * time.Second)
	}

	var runView RunView
	attempts = 0
	for {
		rv, err := runRepoGHWithResponse[RunView](ctx, "run", "view", fmt.Sprintf("%d", run.DatabaseID), "--json", getJSONFields[RunView]())
		if err != nil {
			return 0, fmt.Errorf("failed to run gh command: %w", err)
		}
		runView = rv
		if runView.Status == RunViewStatusCompleted {
			break
		}
		if attempts >= 500 {
			return 0, fmt.Errorf("failed to find a completed run after %d attempts", attempts)
		}
		attempts++
		time.Sleep(5 * time.Second)
	}
	if runView.Conclusion != RunViewConclusionSuccess {
		return 0, fmt.Errorf("run failed: %s", runView.Conclusion)
	}
	return runView.UpdatedAt.Sub(runView.StartedAt), nil
}

func clearCache(ctx context.Context) error {
	log.Printf("[CMD] gh cache delete --all")
	_, _, _ = gh.ExecContext(ctx, "cache", "delete", "--all")
	log.Printf("[AWS] remove all images from ECR")
	if err := tryNTimes(func() error {
		resp, err := ecrClient.DescribeImages(ctx, &ecr.DescribeImagesInput{
			RepositoryName: aws.String(ecrRepoName),
		})
		if err != nil {
			return err
		}
		if len(resp.ImageDetails) == 0 {
			return nil
		}

		imageIds := make([]types.ImageIdentifier, 0, len(resp.ImageDetails))
		for _, image := range resp.ImageDetails {
			imageIds = append(imageIds, types.ImageIdentifier{ImageDigest: image.ImageDigest})
		}

		if _, err := ecrClient.BatchDeleteImage(ctx, &ecr.BatchDeleteImageInput{
			RepositoryName: aws.String(ecrRepoName),
			ImageIds:       imageIds,
		}); err != nil {
			return err
		}
		return nil
	}, 1, 3); err != nil {
		return fmt.Errorf("failed to remove all images from ECR: %w", err)
	}
	return nil
}

func tryNTimes(f func() error, maxCount, maxAttempt int) error {
	attempt := 0
	count := 0
	for {
		if err := f(); err != nil {
			attempt += 1
			if attempt >= maxAttempt {
				return fmt.Errorf("reached max attempt: %d, error: %w", attempt, err)
			}
			log.Printf("retrying after error: %v", err)
			continue
		}
		count += 1
		if count >= maxCount {
			break
		}
	}
	return nil
}

var (
	ecrClient *ecr.Client
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	cfg.Region = "us-west-2"
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	ecrClient = ecr.NewFromConfig(cfg)
}

func main() {
	targetJobs := []string{
		//"multistage-copy-nocache",
		"multistage-copy-layercache-inline",
		"multistage-copy-layercache-registry",
		"multistage-copy-layercache-gha",
		"multistage-copy-layercache-local",
		"multistage-mount-layercache-gha",
		"multistage-mount-layercache-gocache-gha",
		"runneronly-layercache-gocache-gha",
		"ko",
	}

	log.Printf("targetJobs: %v", targetJobs)

	ctx := context.Background()
	maxCount := 3
	maxAttempts := 3

	log.Printf("maxCount: %d, maxAttempts: %d", maxCount, maxAttempts)

	type Result struct {
		TargetJob           string          `json:"targetJob"`
		NoCache             []time.Duration `json:"noCache"`
		UseCacheNoChanges   []time.Duration `json:"useCacheNoChanges"`
		UseCachePkgChanges  []time.Duration `json:"useCachePkgChanges"`
		UseCacheCodeChanges []time.Duration `json:"useCacheCodeChanges"`
	}
	results := make([]Result, 0, len(targetJobs))

	for _, targetJob := range targetJobs {
		result := Result{TargetJob: targetJob}

		// Use cache (no changes)
		log.Printf("%v / Use cache (no changes): start", targetJob)
		if err := tryNTimes(func() error {
			if err := clearCache(ctx); err != nil {
				return err
			}
			if _, err := runJobAndMeasure(ctx, targetJob, "main"); err != nil {
				return err
			}
			return nil
		}, 1, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}
		if err := tryNTimes(func() error {
			duration, err := runJobAndMeasure(ctx, targetJob, "main")
			if err != nil {
				return err
			}
			result.UseCacheNoChanges = append(result.UseCacheNoChanges, duration)
			return nil
		}, maxCount, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}
		log.Printf("%v / Use cache (no changes): %v", targetJob, result.UseCacheNoChanges)

		// No cache & Use cache (pkg changes)
		log.Printf("%v / Use cache (pkg changes): start", targetJob)
		if err := tryNTimes(func() error {
			if err := clearCache(ctx); err != nil {
				return err
			}
			noCacheDuration, err := runJobAndMeasure(ctx, targetJob, "main")
			if err != nil {
				return err
			}
			useCachePkgChangesDuration, err := runJobAndMeasure(ctx, targetJob, "pkg-changes")
			if err != nil {
				return err
			}
			result.NoCache = append(result.NoCache, noCacheDuration)
			result.UseCachePkgChanges = append(result.UseCachePkgChanges, useCachePkgChangesDuration)
			return nil
		}, maxCount, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}
		log.Printf("%v / Use cache (pkg changes): %v", targetJob, result.UseCachePkgChanges)

		// Use cache (code changes)
		log.Printf("%v / Use cache (code changes): start", targetJob)
		if err := tryNTimes(func() error {
			if err := clearCache(ctx); err != nil {
				return err
			}
			if _, err := runJobAndMeasure(ctx, targetJob, "main"); err != nil {
				return err
			}
			duration, err := runJobAndMeasure(ctx, targetJob, "code-changes")
			if err != nil {
				return err
			}
			result.UseCacheCodeChanges = append(result.UseCacheCodeChanges, duration)
			return nil
		}, maxCount, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}
		log.Printf("%v / Use cache (code changes): %v", targetJob, result.UseCacheCodeChanges)

		results = append(results, result)

		// Output current results
		resultJSON, err := json.Marshal(results)
		if err != nil {
			log.Fatalf("failed to marshal results: %v", err)
		}
		fmt.Println(string(resultJSON))

		outJSONFile, err := os.Create(fmt.Sprintf("%s_results_(%d_of_%d).json", time.Now().Format("20060102_150405"), len(results), len(targetJobs)))
		if err != nil {
			log.Fatalf("failed to create outJSONFile: %v", err)
		}
		if _, err := outJSONFile.Write(resultJSON); err != nil {
			log.Fatalf("failed to write resultJSON: %v", err)
		}
		if err := outJSONFile.Close(); err != nil {
			log.Fatalf("failed to close outJSONFile: %v", err)
		}
	}

}
