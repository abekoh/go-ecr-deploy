package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/cli/go-gh/v2"
)

const repositoryName = "abekoh/go-ecr-deploy"

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

func runGH(ctx context.Context, args ...string) error {
	commandArgs := append(args, "--repo", repositoryName)
	log.Printf("gh %s", strings.Join(commandArgs, " "))
	_, _, err := gh.ExecContext(ctx, commandArgs...)
	return err
}

func runGHWithResponse[T any](ctx context.Context, args ...string) (T, error) {
	commandArgs := append(args, "--repo", repositoryName, "--json", getJSONFields[T]())
	log.Printf("gh %s", strings.Join(commandArgs, " "))
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
	if err := runGH(ctx, "workflow", "run", "deploy.yml", "-f", "targetJob="+targetJob, "--ref", branch); err != nil {
		return 0, fmt.Errorf("failed to run gh command: %w", err)
	}
	time.Sleep(5 * time.Second)

	var run RunListElement
	attempts := 0
	for {
		runs, err := runGHWithResponse[[]RunListElement](ctx, "run", "list", "--limit", "1", "--json", getJSONFields[RunListElement]())
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
		if attempts >= 3 {
			return 0, fmt.Errorf("failed to find a run after %d attempts", attempts)
		}
		attempts++
		time.Sleep(5 * time.Second)
	}

	var runView RunView
	attempts = 0
	for {
		rv, err := runGHWithResponse[RunView](ctx, "run", "view", fmt.Sprintf("%d", run.DatabaseID), "--json", getJSONFields[RunView]())
		if err != nil {
			return 0, fmt.Errorf("failed to run gh command: %w", err)
		}
		runView = rv
		if runView.Status == RunViewStatusCompleted {
			break
		}
		if attempts >= 100 {
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
	if err := runGH(ctx, "cache", "clear", "--all"); err != nil {
		return fmt.Errorf("failed to run gh command: %w", err)
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
			continue
		}
		count += 1
		if count >= maxCount {
			break
		}
	}
	return nil
}

func main() {
	targetJobs := []string{
		"multistage-copy-nocache",
		"multistage-copy-layercache-inline",
		"multistage-copy-layercache-registry",
		"multistage-copy-layercache-gha",
		"multistage-copy-layercache-local",
		"multistage-mount-layercache-gha",
		"multistage-mount-layercache-gocache-gha",
		"runneronly-layercache-gocache-gha",
		"ko",
	}

	ctx := context.Background()
	maxCount := 3
	maxAttempts := 3

	type Result struct {
		TargetJob           string
		NoCache             []time.Duration
		UseCacheNoChanges   []time.Duration
		UseCachePkgChanges  []time.Duration
		UseCacheCodeChanges []time.Duration
	}
	results := make([]Result, 0, len(targetJobs))

	for _, targetJob := range targetJobs {
		result := Result{TargetJob: targetJob}

		// No cache
		if err := tryNTimes(func() error {
			if err := clearCache(ctx); err != nil {
				return err
			}
			duration, err := runJobAndMeasure(ctx, targetJob, "main")
			if err != nil {
				return err
			}
			result.NoCache = append(result.NoCache, duration)
			return nil
		}, maxCount, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}

		// Use cache (no changes)
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

		// Use cache (pkg changes)
		if err := tryNTimes(func() error {
			if err := clearCache(ctx); err != nil {
				return err
			}
			if _, err := runJobAndMeasure(ctx, targetJob, "main"); err != nil {
				return err
			}
			duration, err := runJobAndMeasure(ctx, targetJob, "pkg-changes")
			if err != nil {
				return err
			}
			result.UseCachePkgChanges = append(result.UseCachePkgChanges, duration)
			return nil
		}, maxCount, maxAttempts); err != nil {
			log.Fatalf("failed to tryNTimes: %v", err)
		}

		// Use cache (code changes)
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

		results = append(results, result)
	}
}
