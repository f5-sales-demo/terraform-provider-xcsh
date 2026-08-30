// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package workloadbench validates and aggregates sanitized provider benchmark receipts.
package workloadbench

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	ReceiptSchemaVersion = 1
	expectedRepository   = "f5-sales-demo/terraform-provider-xcsh"
	expectedRunnerImage  = "sha256:8817d93949ce0429b16bbcae686065b81d976b43df22e15e90379b5978c6dc2b"
)

var (
	commitPattern         = regexp.MustCompile(`^[0-9a-f]{40}$`)
	currentVariantPattern = regexp.MustCompile(`^d(8|16)-current(?:-examples-c(1|2|4|8))?$`)
	digestPattern         = regexp.MustCompile(`^sha256:[0-9a-f]{64}$`)
	workloadArgv          = map[string][]string{
		"go-build":                     {"scripts/run-fixed-benchmark-workload.sh", "go-build"},
		"go-vet":                       {"scripts/run-fixed-benchmark-workload.sh", "go-vet"},
		"go-race":                      {"scripts/run-fixed-benchmark-workload.sh", "go-race"},
		"provider-generation":          {"scripts/run-fixed-benchmark-workload.sh", "provider-generation"},
		"documentation-generation":     {"scripts/run-fixed-benchmark-workload.sh", "documentation-generation"},
		"terraform-example-validation": {"scripts/run-fixed-benchmark-workload.sh", "terraform-example-validation"},
		"release-source-reproduction":  {"scripts/run-fixed-benchmark-workload.sh", "release-source-reproduction"},
	}
)

func FixedArgv(workload string) ([]string, bool) {
	argv, ok := workloadArgv[workload]
	return append([]string(nil), argv...), ok
}

func FixedRunner(profile string) (RunnerAttestation, bool) {
	profiles := map[string]RunnerAttestation{
		"d8":  {Label: "managed-socketless", Profile: "socketless", VMSize: "Standard_D8ads_v5", CPULimit: 7, MemoryLimit: 28 * 1024 * 1024 * 1024},
		"d16": {Label: "terraform-provider-xcsh-compute", Profile: "compute", VMSize: "Standard_D16ads_v5", CPULimit: 15, MemoryLimit: 56 * 1024 * 1024 * 1024},
	}
	runner, ok := profiles[profile]
	return runner, ok
}

type Identity struct {
	Repository string `json:"repository"`
	Commit     string `json:"commit"`
	RunID      string `json:"run_id"`
	JobID      string `json:"job_id"`
	RunAttempt string `json:"run_attempt"`
	PairID     string `json:"pair_id"`
}

type Toolchain struct {
	SpecRelease       string `json:"spec_release"`
	SpecPinSHA256     string `json:"spec_pin_sha256"`
	GoVersion         string `json:"go_version"`
	TerraformVersion  string `json:"terraform_version"`
	DocumentationTool string `json:"documentation_tool"`
	ActionCheckout    string `json:"action_checkout"`
	RunnerImage       string `json:"runner_image"`
}

type RunnerAttestation struct {
	Label        string `json:"label"`
	Profile      string `json:"profile"`
	VMSize       string `json:"vm_size"`
	CPULimit     int    `json:"cpu_limit"`
	MemoryLimit  int64  `json:"memory_limit"`
	DockerSocket bool   `json:"docker_socket"`
}

type WorkloadIdentity struct {
	ID   string   `json:"id"`
	Argv []string `json:"argv"`
}

type Configuration struct {
	VariantID          string `json:"variant_id"`
	GOFLAGS            string `json:"goflags"`
	GOMAXPROCS         int    `json:"gomaxprocs"`
	ExampleConcurrency int    `json:"example_concurrency"`
}

type CacheIdentity struct {
	State string `json:"state"`
	Key   string `json:"key"`
}

type Metrics struct {
	ElapsedSeconds   float64 `json:"elapsed_seconds"`
	CPUUsageUsec     int64   `json:"cpu_usage_usec"`
	CPUUserUsec      int64   `json:"cpu_user_usec"`
	CPUSystemUsec    int64   `json:"cpu_system_usec"`
	CPUThrottledUsec int64   `json:"cpu_throttled_usec"`
	PeakMemoryBytes  int64   `json:"peak_memory_bytes"`
	MemoryLimitBytes int64   `json:"memory_limit_bytes"`
	PeakMemoryRatio  float64 `json:"peak_memory_ratio"`
	OOMEvents        int64   `json:"oom_events"`
}

type Exit struct {
	Code          int    `json:"code"`
	ErrorCategory string `json:"error_category"`
}

type Receipt struct {
	SchemaVersion    int               `json:"schema_version"`
	Identity         Identity          `json:"identity"`
	Toolchain        Toolchain         `json:"toolchain"`
	Runner           RunnerAttestation `json:"runner"`
	Workload         WorkloadIdentity  `json:"workload"`
	Configuration    Configuration     `json:"configuration"`
	Cache            CacheIdentity     `json:"cache"`
	Metrics          Metrics           `json:"metrics"`
	OutputTreeSHA256 string            `json:"output_tree_sha256"`
	Exit             Exit              `json:"exit"`
	SampleKind       string            `json:"sample_kind"`
	SampleIndex      int               `json:"sample_index"`
}

func MarshalReceipt(receipt Receipt) ([]byte, error) {
	if err := receipt.Validate(); err != nil {
		return nil, err
	}
	encoded, err := json.Marshal(receipt)
	if err != nil {
		return nil, fmt.Errorf("marshal receipt: %w", err)
	}
	return append(encoded, '\n'), nil
}

func ParseReceipt(payload []byte) (Receipt, error) {
	var receipt Receipt
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&receipt); err != nil {
		return Receipt{}, fmt.Errorf("decode receipt: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return Receipt{}, errors.New("receipt contains trailing JSON")
	}
	if err := receipt.Validate(); err != nil {
		return Receipt{}, err
	}
	return receipt, nil
}

func (receipt Receipt) Validate() error {
	if receipt.SchemaVersion != ReceiptSchemaVersion {
		return fmt.Errorf("unsupported receipt schema %d", receipt.SchemaVersion)
	}
	identity := receipt.Identity
	if identity.Repository != expectedRepository || !commitPattern.MatchString(identity.Commit) ||
		identity.RunID == "" || identity.JobID == "" || identity.RunAttempt == "" || identity.PairID == "" {
		return errors.New("receipt identity is incomplete or untrusted")
	}
	toolchain := receipt.Toolchain
	if toolchain.SpecRelease != "v4.0.2" || toolchain.SpecPinSHA256 != "sha256:a3ed1000e5bf0b4d0f694fd70041e07f4f7b59e46f699282389d58b3a31d6972" ||
		toolchain.GoVersion != "go1.25.12" || toolchain.TerraformVersion != "1.15.8" ||
		toolchain.DocumentationTool != "github.com/hashicorp/terraform-plugin-docs@v0.25.0+dirty" ||
		toolchain.ActionCheckout != "3d3c42e5aac5ba805825da76410c181273ba90b1" || toolchain.RunnerImage != expectedRunnerImage {
		return errors.New("toolchain identity is incomplete or unexpected")
	}
	if err := validateRunner(receipt.Runner); err != nil {
		return err
	}
	expectedVariantPrefix, expectedJob := "d8-", "benchmark-d8"
	if receipt.Runner.Label == "terraform-provider-xcsh-compute" {
		expectedVariantPrefix, expectedJob = "d16-", "benchmark-d16"
	}
	if !strings.HasPrefix(receipt.Configuration.VariantID, expectedVariantPrefix) || identity.JobID != expectedJob {
		return errors.New("variant or job identity does not match the attested runner")
	}
	wantArgv, ok := workloadArgv[receipt.Workload.ID]
	if !ok || strings.Join(receipt.Workload.Argv, "\x00") != strings.Join(wantArgv, "\x00") {
		return errors.New("workload command is not a fixed workload")
	}
	if err := validateConfiguration(receipt.Configuration, receipt.Workload.ID); err != nil {
		return err
	}
	if receipt.Cache.State != "cold" && receipt.Cache.State != "warm" {
		return errors.New("cache state must be cold or warm")
	}
	if receipt.Cache.Key == "" || !map[string]bool{"validation": true, "measured": true}[receipt.SampleKind] {
		return errors.New("sample or cache identity is incomplete")
	}
	if (receipt.SampleKind == "validation" && receipt.SampleIndex != 0) ||
		(receipt.SampleKind == "measured" && (receipt.SampleIndex < 1 || receipt.SampleIndex > 5)) {
		return errors.New("sample index does not match its kind")
	}
	metrics := receipt.Metrics
	if metrics.ElapsedSeconds <= 0 || metrics.CPUUsageUsec < 0 || metrics.PeakMemoryBytes < 0 ||
		metrics.MemoryLimitBytes != receipt.Runner.MemoryLimit || metrics.PeakMemoryRatio < 0 {
		return errors.New("receipt metrics are incomplete")
	}
	if !digestPattern.MatchString(receipt.OutputTreeSHA256) {
		return errors.New("output tree digest is malformed")
	}
	allowedCategory := map[string]bool{"none": true, "command_failed": true, "attestation_failed": true, "digest_failed": true, "infrastructure_error": true}
	if receipt.Exit.Code < 0 || receipt.Exit.Code > 255 || !allowedCategory[receipt.Exit.ErrorCategory] {
		return errors.New("exit status is malformed")
	}
	return nil
}

func validateRunner(runner RunnerAttestation) error {
	d8, _ := FixedRunner("d8")
	d16, _ := FixedRunner("d16")
	expected := map[string]RunnerAttestation{d8.Label: d8, d16.Label: d16}
	want, ok := expected[runner.Label]
	if !ok || runner.Profile != want.Profile || runner.VMSize != want.VMSize || runner.CPULimit != want.CPULimit ||
		runner.MemoryLimit != want.MemoryLimit || runner.DockerSocket {
		return errors.New("runner attestation does not match a fixed socketless profile")
	}
	return nil
}

func validateConfiguration(configuration Configuration, workload string) error {
	if configuration.VariantID == "" {
		return errors.New("variant identity is empty")
	}
	if configuration.GOMAXPROCS == 0 {
		if configuration.GOFLAGS != "" || !currentVariantPattern.MatchString(configuration.VariantID) {
			return errors.New("current configuration is malformed")
		}
	} else if configuration.GOFLAGS != "-p="+strconv.Itoa(configuration.GOMAXPROCS) ||
		!map[int]bool{1: true, 2: true, 4: true, 8: true, 16: true}[configuration.GOMAXPROCS] {
		return errors.New("CPU configuration is not fixed")
	}
	if !map[int]bool{1: true, 2: true, 4: true, 8: true}[configuration.ExampleConcurrency] {
		return errors.New("example concurrency is not fixed")
	}
	if workload != "terraform-example-validation" && configuration.ExampleConcurrency != 1 {
		return errors.New("example concurrency applies only to example validation")
	}
	return nil
}

type CachePlan struct {
	Root           string
	GoBuild        string
	GoModules      string
	Terraform      string
	PackageManager string
}

func PlanCache(root, profile, workload, variant, state string, sample int) (CachePlan, error) {
	if root == "" || !map[string]bool{"d8": true, "d16": true}[profile] || workloadArgv[workload] == nil || variant == "" ||
		!map[string]bool{"cold": true, "warm": true}[state] || sample < 1 || sample > 5 {
		return CachePlan{}, errors.New("invalid fixed cache identity")
	}
	leaf := "shared"
	if state == "cold" {
		leaf = fmt.Sprintf("sample-%d", sample)
	}
	cacheRoot := filepath.Join(root, profile, workload, variant, state, leaf)
	return CachePlan{
		Root: cacheRoot, GoBuild: filepath.Join(cacheRoot, "go-build"), GoModules: filepath.Join(cacheRoot, "go-mod"),
		Terraform: filepath.Join(cacheRoot, "terraform"), PackageManager: filepath.Join(cacheRoot, "package-manager"),
	}, nil
}

func CanonicalTreeDigest(root string) (string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if relative == "." {
			return nil
		}
		parts := strings.Split(filepath.ToSlash(relative), "/")
		for _, part := range parts {
			if part == ".git" || part == ".terraform" || part == ".cache" || part == "benchmark-receipts" {
				if entry.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if entry.IsDir() || strings.HasSuffix(entry.Name(), ".log") {
			return nil
		}
		paths = append(paths, relative)
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("walk output tree: %w", err)
	}
	sort.Strings(paths)
	hash := sha256.New()
	for _, relative := range paths {
		path := filepath.Join(root, relative)
		info, err := os.Lstat(path)
		if err != nil {
			return "", err
		}
		var data []byte
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return "", err
			}
			data = []byte(target)
		} else {
			data, err = os.ReadFile(path)
			if err != nil {
				return "", err
			}
		}
		writeDigestField(hash, []byte(filepath.ToSlash(relative)))
		writeDigestField(hash, []byte(info.Mode().String()))
		writeDigestField(hash, data)
	}
	return "sha256:" + hex.EncodeToString(hash.Sum(nil)), nil
}

func writeDigestField(writer io.Writer, value []byte) {
	_ = binary.Write(writer, binary.BigEndian, uint64(len(value)))
	_, _ = writer.Write(value)
}

type Distribution struct {
	Median float64 `json:"median"`
	P95    float64 `json:"p95"`
}

type AggregateVariant struct {
	VariantID        string       `json:"variant_id"`
	Cold             Distribution `json:"cold"`
	Warm             Distribution `json:"warm"`
	ColdImprovement  float64      `json:"cold_improvement"`
	WarmImprovement  float64      `json:"warm_improvement"`
	WorstP95Ratio    float64      `json:"worst_p95_ratio"`
	PeakMemoryRatio  float64      `json:"peak_memory_ratio"`
	Eligible         bool         `json:"eligible"`
	RejectionReasons []string     `json:"rejection_reasons,omitempty"`
	digest           string
}

type AggregateResult struct {
	Workload   string             `json:"workload"`
	Baseline   AggregateVariant   `json:"baseline"`
	Candidates []AggregateVariant `json:"candidates"`
	Winner     string             `json:"winner,omitempty"`
}

func Aggregate(receipts []Receipt, baselineID string) (AggregateResult, error) {
	if len(receipts) == 0 || baselineID == "" {
		return AggregateResult{}, errors.New("receipts and baseline are required")
	}
	workload := receipts[0].Workload.ID
	groups := map[string]map[string][]Receipt{}
	for _, receipt := range receipts {
		if err := receipt.Validate(); err != nil {
			return AggregateResult{}, err
		}
		if receipt.Workload.ID != workload {
			return AggregateResult{}, errors.New("aggregate contains multiple workloads")
		}
		if groups[receipt.Configuration.VariantID] == nil {
			groups[receipt.Configuration.VariantID] = map[string][]Receipt{}
		}
		if receipt.SampleKind == "measured" {
			groups[receipt.Configuration.VariantID][receipt.Cache.State] = append(groups[receipt.Configuration.VariantID][receipt.Cache.State], receipt)
		}
	}
	baselineGroups, ok := groups[baselineID]
	if !ok {
		return AggregateResult{}, errors.New("baseline variant is absent")
	}
	baseline := aggregateVariant(baselineID, baselineGroups, nil)
	if len(baseline.RejectionReasons) != 0 {
		return AggregateResult{}, fmt.Errorf("baseline is incomplete: %s", strings.Join(baseline.RejectionReasons, ", "))
	}
	result := AggregateResult{Workload: workload, Baseline: baseline}
	for variantID, samples := range groups {
		if variantID == baselineID {
			continue
		}
		candidate := aggregateVariant(variantID, samples, &baseline)
		result.Candidates = append(result.Candidates, candidate)
	}
	sort.Slice(result.Candidates, func(i, j int) bool { return rankLess(result.Candidates[i], result.Candidates[j]) })
	for _, candidate := range result.Candidates {
		if candidate.Eligible {
			result.Winner = candidate.VariantID
			break
		}
	}
	return result, nil
}

func aggregateVariant(variant string, groups map[string][]Receipt, baseline *AggregateVariant) AggregateVariant {
	result := AggregateVariant{VariantID: variant}
	all := append(append([]Receipt{}, groups["cold"]...), groups["warm"]...)
	result.Cold = distribution(groups["cold"])
	result.Warm = distribution(groups["warm"])
	if len(groups["cold"]) != 5 || len(groups["warm"]) != 5 {
		result.RejectionReasons = append(result.RejectionReasons, "requires five cold and five warm samples")
	}
	seenSamples := map[string]bool{}
	digest := ""
	for _, receipt := range all {
		key := receipt.Cache.State + "/" + strconv.Itoa(receipt.SampleIndex)
		if seenSamples[key] {
			result.RejectionReasons = append(result.RejectionReasons, "duplicate sample identity")
		}
		seenSamples[key] = true
		if receipt.Exit.Code != 0 || receipt.Exit.ErrorCategory != "none" {
			result.RejectionReasons = append(result.RejectionReasons, "sample failed")
		}
		if receipt.Metrics.PeakMemoryRatio >= 0.80 || receipt.Metrics.OOMEvents != 0 {
			result.RejectionReasons = append(result.RejectionReasons, "memory gate failed")
		}
		result.PeakMemoryRatio = math.Max(result.PeakMemoryRatio, receipt.Metrics.PeakMemoryRatio)
		if digest == "" {
			digest = receipt.OutputTreeSHA256
		} else if digest != receipt.OutputTreeSHA256 {
			result.RejectionReasons = append(result.RejectionReasons, "output digest mismatch")
		}
	}
	result.digest = digest
	if baseline == nil {
		result.RejectionReasons = uniqueStrings(result.RejectionReasons)
		return result
	}
	result.ColdImprovement = improvement(baseline.Cold.Median, result.Cold.Median)
	result.WarmImprovement = improvement(baseline.Warm.Median, result.Warm.Median)
	result.WorstP95Ratio = math.Max(ratio(result.Cold.P95, baseline.Cold.P95), ratio(result.Warm.P95, baseline.Warm.P95))
	if result.ColdImprovement < 0.20 || result.WarmImprovement < 0.20 {
		result.RejectionReasons = append(result.RejectionReasons, "median improvement below 20 percent")
	}
	if result.Cold.P95 > baseline.Cold.P95 || result.Warm.P95 > baseline.Warm.P95 {
		result.RejectionReasons = append(result.RejectionReasons, "p95 regression")
	}
	if result.digest == "" || result.digest != baseline.digest {
		result.RejectionReasons = append(result.RejectionReasons, "output digest differs from baseline")
	}
	result.RejectionReasons = uniqueStrings(result.RejectionReasons)
	result.Eligible = len(result.RejectionReasons) == 0
	return result
}

func distribution(receipts []Receipt) Distribution {
	values := make([]float64, 0, len(receipts))
	for _, receipt := range receipts {
		values = append(values, receipt.Metrics.ElapsedSeconds)
	}
	if len(values) == 0 {
		return Distribution{}
	}
	sort.Float64s(values)
	median := values[len(values)/2]
	p95Index := int(math.Ceil(0.95*float64(len(values)))) - 1
	return Distribution{Median: median, P95: values[p95Index]}
}

func improvement(baseline, candidate float64) float64 {
	if baseline <= 0 {
		return 0
	}
	return (baseline - candidate) / baseline
}

func ratio(candidate, baseline float64) float64 {
	if baseline <= 0 {
		return math.Inf(1)
	}
	return candidate / baseline
}

func rankLess(left, right AggregateVariant) bool {
	if left.Eligible != right.Eligible {
		return left.Eligible
	}
	leftMinimum := math.Min(left.ColdImprovement, left.WarmImprovement)
	rightMinimum := math.Min(right.ColdImprovement, right.WarmImprovement)
	if leftMinimum != rightMinimum {
		return leftMinimum > rightMinimum
	}
	if left.WorstP95Ratio != right.WorstP95Ratio {
		return left.WorstP95Ratio < right.WorstP95Ratio
	}
	if left.PeakMemoryRatio != right.PeakMemoryRatio {
		return left.PeakMemoryRatio < right.PeakMemoryRatio
	}
	leftN, rightN := variantN(left.VariantID), variantN(right.VariantID)
	if leftN != rightN {
		return leftN < rightN
	}
	return left.VariantID < right.VariantID
}

func variantN(variant string) int {
	index := strings.LastIndex(variant, "-n")
	if index < 0 {
		return math.MaxInt
	}
	value, err := strconv.Atoi(variant[index+2:])
	if err != nil {
		return math.MaxInt
	}
	return value
}

func uniqueStrings(values []string) []string {
	sort.Strings(values)
	result := values[:0]
	for _, value := range values {
		if len(result) == 0 || result[len(result)-1] != value {
			result = append(result, value)
		}
	}
	return result
}
