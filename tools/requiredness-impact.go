// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/requiredness"
)

func main() {
	baselineProvider := flag.String("baseline-provider", "", "baseline generated internal/provider directory")
	candidateProvider := flag.String("candidate-provider", "internal/provider", "candidate generated internal/provider directory")
	baselineSpecs := flag.String("baseline-specs", "", "baseline released spec bundle directory")
	candidateSpecs := flag.String("candidate-specs", "docs/specifications/api", "candidate released spec bundle directory")
	baselineProviderRef := flag.String("baseline-provider-ref", "", "immutable baseline provider reference for the report")
	baselineSpecVersion := flag.String("baseline-spec-version", "", "baseline spec version")
	candidateSpecVersion := flag.String("candidate-spec-version", "", "candidate spec version")
	output := flag.String("output", "", "output JSON path")
	flag.Parse()

	if *baselineProvider == "" || *baselineSpecs == "" || *output == "" || *baselineProviderRef == "" || *baselineSpecVersion == "" || *candidateSpecVersion == "" {
		fmt.Fprintln(os.Stderr, "baseline provider/specs/reference, both spec versions, and output are required")
		os.Exit(2)
	}
	report, err := requiredness.Compare(*baselineProvider, *candidateProvider, *baselineSpecs, *candidateSpecs, requiredness.Report{
		BaselineProvider: *baselineProviderRef,
		BaselineSpec:     *baselineSpecVersion,
		CandidateSpec:    *candidateSpecVersion,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := requiredness.Write(*output, report); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("verified %d Required-to-Optional transitions\n", len(report.Transitions))
}
