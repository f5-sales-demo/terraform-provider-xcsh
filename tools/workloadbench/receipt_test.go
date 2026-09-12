// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package workloadbench

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func validReceipt(variant, profile, cache string, sample int, duration float64) Receipt {
	cpu := 7
	memory := int64(28 * 1024 * 1024 * 1024)
	label := "managed-socketless"
	vm := "Standard_D8ads_v5"
	if profile == "d16" {
		cpu = 15
		memory = 56 * 1024 * 1024 * 1024
		label = "terraform-provider-xcsh-compute"
		vm = "Standard_D16ads_v5"
	}
	return Receipt{
		SchemaVersion: ReceiptSchemaVersion,
		Identity: Identity{
			Repository: "f5-sales-demo/terraform-provider-xcsh",
			Commit:     strings.Repeat("a", 40),
			RunID:      "123",
			JobID:      "benchmark-" + profile,
			RunAttempt: "1",
			PairID:     "go-build/" + cache + "/" + string(rune('0'+sample)),
		},
		Toolchain: Toolchain{
			SpecRelease:   "v5.0.1",
			SpecPinSHA256: "sha256:f606c2b53ec7d073c8d1d7b0bb5670a0660cb9082cfb07173fef8cda6934e477",
			GoVersion:     "go1.25.12", TerraformVersion: "1.15.8",
			DocumentationTool: "github.com/hashicorp/terraform-plugin-docs@v0.25.0+dirty",
			ActionCheckout:    "3d3c42e5aac5ba805825da76410c181273ba90b1",
			RunnerImage:       "sha256:2a0243be5404daa0f52bae16384f53dbc04554e31406ed0db45152d92f6187e1",
		},
		Runner:           RunnerAttestation{Label: label, Profile: map[string]string{"d8": "socketless", "d16": "compute"}[profile], VMSize: vm, CPULimit: cpu, MemoryLimit: memory, DockerSocket: false},
		Workload:         WorkloadIdentity{ID: "go-build", Argv: []string{"scripts/run-fixed-benchmark-workload.sh", "go-build"}},
		Configuration:    Configuration{VariantID: variant, GOFLAGS: "-p=4", GOMAXPROCS: 4, ExampleConcurrency: 1},
		Cache:            CacheIdentity{State: cache, Key: profile + "/go-build/" + variant + "/" + cache},
		Metrics:          Metrics{ElapsedSeconds: duration, CPUUsageUsec: 1000, PeakMemoryBytes: memory / 2, MemoryLimitBytes: memory, PeakMemoryRatio: 0.5},
		OutputTreeSHA256: "sha256:" + strings.Repeat("c", 64),
		Exit:             Exit{Code: 0, ErrorCategory: "none"},
		SampleKind:       "measured",
		SampleIndex:      sample,
	}
}

func TestReceiptRoundTripRejectsMalformedOrIncompleteSamples(t *testing.T) {
	receipt := validReceipt("d8-n4", "d8", "cold", 1, 10)
	encoded, err := MarshalReceipt(receipt)
	if err != nil {
		t.Fatalf("marshal valid receipt: %v", err)
	}
	parsed, err := ParseReceipt(encoded)
	if err != nil {
		t.Fatalf("parse valid receipt: %v", err)
	}
	if parsed.Identity.Commit != receipt.Identity.Commit {
		t.Fatalf("commit = %q, want %q", parsed.Identity.Commit, receipt.Identity.Commit)
	}

	for name, payload := range map[string][]byte{
		"unknown field": append(encoded[:len(encoded)-2], []byte(",\"secret\":true}\n")...),
		"malformed":     []byte("{not-json}"),
		"incomplete":    []byte(`{"schema_version":1}`),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := ParseReceipt(payload); err == nil {
				t.Fatal("invalid receipt was accepted")
			}
		})
	}
}

func TestReceiptRequiresIdentityAndExactProfileAttestation(t *testing.T) {
	for _, mutate := range []func(*Receipt){
		func(r *Receipt) { r.Identity.Commit = "" },
		func(r *Receipt) { r.Identity.PairID = "" },
		func(r *Receipt) { r.Runner.Profile = "compute" },
		func(r *Receipt) { r.Runner.CPULimit = 4 },
		func(r *Receipt) { r.Runner.MemoryLimit = 8 * 1024 * 1024 * 1024 },
		func(r *Receipt) { r.Runner.DockerSocket = true },
		func(r *Receipt) { r.Toolchain.RunnerImage = "sha256:" + strings.Repeat("d", 64) },
	} {
		receipt := validReceipt("d8-n4", "d8", "cold", 1, 10)
		mutate(&receipt)
		if err := receipt.Validate(); err == nil {
			t.Fatal("invalid identity or attestation was accepted")
		}
	}
}

func TestCachePlanIsolatesCandidatesAndColdSamples(t *testing.T) {
	root := t.TempDir()
	cold1, err := PlanCache(root, "d16", "go-race", "d16-n8", "cold", 1)
	if err != nil {
		t.Fatal(err)
	}
	cold2, err := PlanCache(root, "d16", "go-race", "d16-n8", "cold", 2)
	if err != nil {
		t.Fatal(err)
	}
	warm, err := PlanCache(root, "d16", "go-race", "d16-n8", "warm", 1)
	if err != nil {
		t.Fatal(err)
	}
	other, err := PlanCache(root, "d16", "go-race", "d16-n16", "warm", 1)
	if err != nil {
		t.Fatal(err)
	}
	paths := []string{cold1.Root, cold2.Root, warm.Root, other.Root}
	for i, left := range paths {
		for j, right := range paths {
			if i != j && (left == right || strings.HasPrefix(left+string(os.PathSeparator), right+string(os.PathSeparator))) {
				t.Fatalf("cache roots overlap: %q and %q", left, right)
			}
		}
	}
}

func TestCanonicalTreeDigestIncludesPathModeAndBytes(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "a"), []byte("same"), 0o600); err != nil {
		t.Fatal(err)
	}
	first, err := CanonicalTreeDigest(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(filepath.Join(root, "a"), 0o644); err != nil {
		t.Fatal(err)
	}
	second, err := CanonicalTreeDigest(root)
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatal("mode change did not alter canonical digest")
	}
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".git", "ignored"), []byte("noise"), 0o600); err != nil {
		t.Fatal(err)
	}
	third, err := CanonicalTreeDigest(root)
	if err != nil {
		t.Fatal(err)
	}
	if second != third {
		t.Fatal("excluded Git metadata altered canonical digest")
	}
}

func TestAggregateUsesNearestRankP95AndDeterministicRanking(t *testing.T) {
	var receipts []Receipt
	for _, cache := range []string{"cold", "warm"} {
		for sample, base := range []float64{10, 11, 12, 13, 14} {
			receipts = append(receipts,
				validReceipt("d8-current", "d8", cache, sample+1, base),
				validReceipt("d16-n8", "d16", cache, sample+1, base*0.70),
				validReceipt("d16-n4", "d16", cache, sample+1, base*0.75),
			)
		}
	}
	result, err := Aggregate(receipts, "d8-current")
	if err != nil {
		t.Fatalf("aggregate valid receipts: %v", err)
	}
	if result.Winner != "d16-n8" {
		t.Fatalf("winner = %q, want d16-n8", result.Winner)
	}
	if result.Baseline.Cold.P95 != 14 || result.Baseline.Warm.P95 != 14 {
		t.Fatalf("nearest-rank p95 = cold %.1f warm %.1f, want 14", result.Baseline.Cold.P95, result.Baseline.Warm.P95)
	}
}

func TestAggregateRejectsDigestDriftIncompleteSamplesAndMemory(t *testing.T) {
	for name, mutate := range map[string]func([]Receipt) []Receipt{
		"digest drift": func(receipts []Receipt) []Receipt {
			receipts[len(receipts)-1].OutputTreeSHA256 = "sha256:" + strings.Repeat("d", 64)
			return receipts
		},
		"incomplete": func(receipts []Receipt) []Receipt { return receipts[:len(receipts)-1] },
		"memory": func(receipts []Receipt) []Receipt {
			receipts[len(receipts)-1].Metrics.PeakMemoryRatio = 0.80
			return receipts
		},
	} {
		t.Run(name, func(t *testing.T) {
			var receipts []Receipt
			for _, cache := range []string{"cold", "warm"} {
				for sample := 1; sample <= 5; sample++ {
					receipts = append(receipts,
						validReceipt("d8-current", "d8", cache, sample, 10),
						validReceipt("d16-n8", "d16", cache, sample, 7),
					)
				}
			}
			result, err := Aggregate(mutate(receipts), "d8-current")
			if err != nil {
				t.Fatalf("aggregate: %v", err)
			}
			if result.Winner != "" || result.Candidates[0].Eligible {
				t.Fatal("invalid candidate qualified")
			}
		})
	}
}

func TestAggregateReportsAnIneligibleBaselineWithoutDiscardingEvidence(t *testing.T) {
	var receipts []Receipt
	for _, cache := range []string{"cold", "warm"} {
		for sample := 1; sample <= 5; sample++ {
			baseline := validReceipt("d8-current", "d8", cache, sample, 10)
			baseline.Metrics.PeakMemoryRatio = 0.90
			receipts = append(receipts, baseline, validReceipt("d16-n8", "d16", cache, sample, 7))
		}
	}
	result, err := Aggregate(receipts, "d8-current")
	if err != nil {
		t.Fatalf("aggregate must report a resource-ineligible baseline: %v", err)
	}
	if result.Winner != "" {
		t.Fatalf("winner = %q, want no route when baseline breaches memory gate", result.Winner)
	}
	if len(result.Baseline.RejectionReasons) == 0 || result.Baseline.RejectionReasons[0] != "memory gate failed" {
		t.Fatalf("baseline rejection reasons = %v, want memory gate failure", result.Baseline.RejectionReasons)
	}
	if len(result.Candidates) != 1 || result.Candidates[0].Eligible || !strings.Contains(strings.Join(result.Candidates[0].RejectionReasons, ","), "baseline did not satisfy qualification gates") {
		t.Fatalf("candidate must be explicitly rejected by the baseline gate: %+v", result.Candidates)
	}
}
