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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: measure <target>")
		os.Exit(1)
	}
	target := os.Args[1]

	ctx := context.Background()

	workflowStartedAt := time.Now()
	if err := runGH(ctx, "workflow", "run", "deploy.yml", "-f", "target="+target); err != nil {
		log.Fatalf("failed to run gh command: %v", err)
	}
	time.Sleep(5 * time.Second)

	var run RunListElement
	attempts := 0
	for {
		runs, err := runGHWithResponse[[]RunListElement](ctx, "run", "list", "--limit", "1", "--json", getJSONFields[RunListElement]())
		if err != nil {
			log.Fatalf("failed to run gh command: %v", err)
		}
		if len(runs) == 0 {
			continue
		}
		if runs[0].StartedAt.After(workflowStartedAt) {
			run = runs[0]
			break
		}
		if attempts >= 3 {
			log.Fatalf("failed to find a run after %d attempts", attempts)
		}
		attempts++
		time.Sleep(5 * time.Second)
	}

	var runView RunView
	attempts = 0
	for {
		rv, err := runGHWithResponse[RunView](ctx, "run", "view", fmt.Sprintf("%d", run.DatabaseID), "--json", getJSONFields[RunView]())
		if err != nil {
			log.Fatalf("failed to run gh command: %v", err)
		}
		runView = rv
		if runView.Status == RunViewStatusCompleted {
			break
		}
		if attempts >= 100 {
			log.Fatalf("failed to find a completed run after %d attempts", attempts)
		}
		attempts++
		time.Sleep(5 * time.Second)
	}
	if runView.Conclusion != RunViewConclusionSuccess {
		log.Fatalf("run failed: %s", runView.Conclusion)
	}
	elapsed := runView.UpdatedAt.Sub(runView.StartedAt)
	log.Printf("run succeeded in %s", elapsed)
}
