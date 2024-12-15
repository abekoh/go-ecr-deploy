package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func secondStr(d time.Duration) string {
	s := d.Round(time.Second)
	return fmt.Sprintf("%ds", s/time.Second)
}

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
	sb.WriteString("| | noCache | useCacheNoChanges | useCacheCodeChanges | useCachePkgChanges |\n")
	sb.WriteString("| :-- | --: | --: | --: | --: |\n")
	for _, r := range results {
		sb.WriteString("| " + r.TargetJob)
		sb.WriteString(" | " + secondStr(r.NoCache.AverageElapsed()))
		sb.WriteString(" | " + secondStr(r.UseCacheNoChanges.AverageElapsed()))
		sb.WriteString(" | " + secondStr(r.UseCacheCodeChanges.AverageElapsed()))
		sb.WriteString(" | " + secondStr(r.UseCachePkgChanges.AverageElapsed()))
		sb.WriteString(" |\n")
	}
	sb.WriteString("\n")

	sb.WriteString("## Details\n\n")
	sb.WriteString("| | noCache | useCacheNoChanges | useCacheCodeChanges | useCachePkgChanges |\n")
	sb.WriteString("| :-- | :-: | :-: | :-: | :-: |\n")
	for _, r := range results {
		sb.WriteString("| " + r.TargetJob)
		sb.WriteString(" | " + fmt.Sprintf("%s / %s / %s", r.NoCache[0].DatabaseMDLink(), r.NoCache[1].DatabaseMDLink(), r.NoCache[2].DatabaseMDLink()))
		sb.WriteString(" | " + fmt.Sprintf("%s / %s / %s", r.UseCacheNoChanges[0].DatabaseMDLink(), r.UseCacheNoChanges[1].DatabaseMDLink(), r.UseCacheNoChanges[2].DatabaseMDLink()))
		sb.WriteString(" | " + fmt.Sprintf("%s / %s / %s", r.UseCacheCodeChanges[0].DatabaseMDLink(), r.UseCacheCodeChanges[1].DatabaseMDLink(), r.UseCacheCodeChanges[2].DatabaseMDLink()))
		sb.WriteString(" | " + fmt.Sprintf("%s / %s / %s", r.UseCachePkgChanges[0].DatabaseMDLink(), r.UseCachePkgChanges[1].DatabaseMDLink(), r.UseCachePkgChanges[2].DatabaseMDLink()))
		sb.WriteString(" |\n")
	}
	sb.WriteString("|\n")

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
