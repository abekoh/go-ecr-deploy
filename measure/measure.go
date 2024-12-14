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

func runGHCommand[T any](ctx context.Context, args ...string) (T, error) {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: measure <command>")
		os.Exit(1)
	}
	target := os.Args[1]

	ctx := context.Background()
	res, err := runGHCommand[[]RunListElement](ctx, "run", "list", "--branch", target, "--limit", "1")
	if err != nil {
		log.Fatalf("failed to run gh command: %v", err)
	}
	fmt.Printf("%+v\n", res)
}
