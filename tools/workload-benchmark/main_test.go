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
