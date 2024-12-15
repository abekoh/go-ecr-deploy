package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func summary() {
	jsonF, err := os.Open("summary.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonF.Close()

	var results []MeasureResult
	if err := json.NewDecoder(jsonF).Decode(&results); err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	sb.WriteString("# Summary\n\n")
	sb.WriteString("## Averages\n\n")
	sb.WriteString("| | noCache | useCacheNoChanges | useCachePkgChanges | useCacheCodeChanges |\n")
	sb.WriteString("| --- | --- | --- | --- | --- |\n")
	for _, r := range results {
		sb.WriteString("| " + r.TargetJob)
		sb.WriteString(" | " + r.NoCache.AverageElapsed().String())
		sb.WriteString(" | " + r.UseCacheNoChanges.AverageElapsed().String())
		sb.WriteString(" | " + r.UseCachePkgChanges.AverageElapsed().String())
		sb.WriteString(" | " + r.UseCacheCodeChanges.AverageElapsed().String())
		sb.WriteString(" |\n")
	}
	sb.WriteString("\n")

	summaryMDFilename := "summary.md"
	_ = os.Remove(summaryMDFilename)
	markdownF, err := os.Create(summaryMDFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer markdownF.Close()
	if _, err := markdownF.WriteString(sb.String()); err != nil {
		log.Fatal(err)
	}

}
