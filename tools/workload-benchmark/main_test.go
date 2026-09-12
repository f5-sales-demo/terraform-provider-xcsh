package main

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/workloadbench"
)

func TestBaselineProbeGateOnlyRequiresSuccessfulOutput(t *testing.T) {
	receipt := workloadbench.Receipt{
		Exit: workloadbench.Exit{Code: 0},
		Metrics: workloadbench.Metrics{
			PeakMemoryRatio: 0.99,
			OOMEvents:       1,
		},
	}
	if !baselineProbeProvidesDigest(receipt) {
		t.Fatal("resource-ineligible baseline probe must still provide the canonical digest")
	}
	receipt.Exit.Code = 1
	if baselineProbeProvidesDigest(receipt) {
		t.Fatal("failed baseline probe cannot provide a canonical digest")
	}
}

func TestBaselineProbeOptionsRetainEvidenceDirectory(t *testing.T) {
	options := runOptions{variant: "d8-n4", outputDir: "/receipts"}
	probe := baselineProbeOptions(options)
	if probe.variant != "d8-current" {
		t.Fatalf("probe variant = %q, want d8-current", probe.variant)
	}
	if probe.outputDir != options.outputDir {
		t.Fatalf("probe output directory = %q, want %q", probe.outputDir, options.outputDir)
	}
}

func TestValidationAllowsMeasurementForResourceIneligibleD8Baseline(t *testing.T) {
	receipt := workloadbench.Receipt{Exit: workloadbench.Exit{Code: 0}, Metrics: workloadbench.Metrics{PeakMemoryRatio: 0.95}}
	if !validationAllowsMeasurement("d8", "d8-current", receipt) {
		t.Fatal("D8 current baseline must be measured so an ineligible route is reported deterministically")
	}
	if validationAllowsMeasurement("d8", "d8-n4", receipt) {
		t.Fatal("resource-ineligible candidate must not receive measured samples")
	}
	receipt.Metrics.OOMEvents = 1
	if validationAllowsMeasurement("d8", "d8-current", receipt) {
		t.Fatal("OOM baseline must not receive additional samples")
	}
}
