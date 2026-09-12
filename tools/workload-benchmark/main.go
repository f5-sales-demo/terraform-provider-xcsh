// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/workloadbench"
)

const checkoutActionSHA = "3d3c42e5aac5ba805825da76410c181273ba90b1"

var variantPattern = regexp.MustCompile(`^(d8|d16)-(current|n(1|2|4|8|16))(?:-examples-c(1|2|4|8))?$`)

type rawProfile struct {
	Repository    string  `json:"repository"`
	Commit        string  `json:"commit"`
	RunID         string  `json:"run_id"`
	RunAttempt    string  `json:"run_attempt"`
	JobID         string  `json:"job_id"`
	RunnerProfile string  `json:"runner_profile"`
	ImageDigest   string  `json:"image_digest"`
	Duration      float64 `json:"duration_seconds"`
	CPU           struct {
		UsageUsec     int64 `json:"usage_usec"`
		UserUsec      int64 `json:"user_usec"`
		SystemUsec    int64 `json:"system_usec"`
		ThrottledUsec int64 `json:"throttled_usec"`
	} `json:"cpu"`
	Memory struct {
		PeakBytes      *int64           `json:"peak_bytes"`
		LimitBytes     *int64           `json:"limit_bytes"`
		PeakLimitRatio *float64         `json:"peak_limit_ratio"`
		Events         map[string]int64 `json:"events"`
	} `json:"memory"`
	Exit struct {
		Code int `json:"code"`
	} `json:"exit"`
}

type runOptions struct {
	profile, workload, variant, commit string
	sourceRoot, specSource             string
	cacheRoot, outputDir, baselineDir  string
}

type toolIdentity struct {
	specRelease, specPin, goVersion, terraformVersion, documentationTool string
}

func main() {
	if len(os.Args) < 2 {
		fatal(errors.New("usage: workload-benchmark <run|aggregate>"))
	}
	var err error
	switch os.Args[1] {
	case "run":
		err = runCommand(os.Args[2:])
	case "run-example-suite":
		err = runExampleSuiteCommand(os.Args[2:])
	case "aggregate":
		err = aggregateCommand(os.Args[2:])
	default:
		err = fmt.Errorf("unknown fixed subcommand %q", os.Args[1])
	}
	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "workload benchmark: %v\n", err)
	os.Exit(1)
}

func runCommand(args []string) error {
	set := flag.NewFlagSet("run", flag.ContinueOnError)
	options := runOptions{}
	set.StringVar(&options.profile, "profile", "", "fixed profile: d8 or d16")
	set.StringVar(&options.workload, "workload", "", "fixed workload ID")
	set.StringVar(&options.variant, "variant", "", "fixed configuration variant")
	set.StringVar(&options.commit, "commit", "", "exact pull request head")
	set.StringVar(&options.sourceRoot, "source-root", ".", "source repository")
	set.StringVar(&options.specSource, "spec-source", "docs/specifications/api", "verified specification directory")
	set.StringVar(&options.cacheRoot, "cache-root", "", "run-local cache root")
	set.StringVar(&options.outputDir, "output-dir", "", "sanitized receipt directory")
	set.StringVar(&options.baselineDir, "baseline-dir", "", "D8 validation receipt directory")
	if err := set.Parse(args); err != nil {
		return err
	}
	if set.NArg() != 0 {
		return errors.New("positional commands are forbidden")
	}
	configuration, err := parseVariant(options.profile, options.workload, options.variant)
	if err != nil {
		return err
	}
	if _, ok := workloadbench.FixedArgv(options.workload); !ok {
		return fmt.Errorf("unknown fixed workload %q", options.workload)
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(options.commit) {
		return errors.New("commit must be an exact 40-character SHA")
	}
	for name, value := range map[string]string{"cache root": options.cacheRoot, "output directory": options.outputDir} {
		if value == "" {
			return fmt.Errorf("%s is required", name)
		}
	}
	root, err := filepath.Abs(options.sourceRoot)
	if err != nil {
		return err
	}
	options.sourceRoot = root
	spec, err := filepath.Abs(options.specSource)
	if err != nil {
		return err
	}
	options.specSource = spec
	if err := requireDirectory(options.specSource); err != nil {
		return fmt.Errorf("verified specification bundle: %w", err)
	}
	if err := os.MkdirAll(options.outputDir, 0o755); err != nil {
		return err
	}
	tools, err := inspectToolchain(options.sourceRoot)
	if err != nil {
		return err
	}
	runner, err := attestRunner(options.profile)
	if err != nil {
		return err
	}
	expectedBaselineDigest := ""
	if options.profile == "d8" && options.variant != "d8-current" {
		baselineOptions := baselineProbeOptions(options)
		baselineConfiguration := workloadbench.Configuration{VariantID: "d8-current", ExampleConcurrency: 1}
		baselinePlan, planErr := workloadbench.PlanCache(options.cacheRoot, "d8", options.workload, "d8-current", "cold", 1)
		if planErr != nil {
			return planErr
		}
		baselinePlan.Root = filepath.Join(options.cacheRoot, "d8", options.workload, options.variant, "baseline-probe")
		baselinePlan.GoBuild = filepath.Join(baselinePlan.Root, "go-build")
		baselinePlan.GoModules = filepath.Join(baselinePlan.Root, "go-mod")
		baselinePlan.Terraform = filepath.Join(baselinePlan.Root, "terraform")
		baselinePlan.PackageManager = filepath.Join(baselinePlan.Root, "package-manager")
		baseline, probeErr := executeSample(baselineOptions, tools, runner, baselineConfiguration, baselinePlan, "validation", "cold", 0)
		if probeErr != nil {
			return probeErr
		}
		// Retain the baseline-probe receipt with the job evidence even when the
		// probe fails. The aggregate ignores validation samples for its measured
		// distributions, while the retained receipt prevents an upload failure
		// from hiding an intended benchmark rejection.
		// Apply memory and OOM eligibility gates to the candidate's own
		// validation receipt below; otherwise a resource-ineligible baseline can
		// suppress the candidate receipt entirely.
		if !baselineProbeProvidesDigest(baseline) {
			return nil
		}
		expectedBaselineDigest = baseline.OutputTreeSHA256
	}

	validationPlan, err := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, options.variant, "cold", 1)
	if err != nil {
		return err
	}
	validationPlan.Root = filepath.Join(options.cacheRoot, options.profile, options.workload, options.variant, "validation")
	validationPlan.GoBuild = filepath.Join(validationPlan.Root, "go-build")
	validationPlan.GoModules = filepath.Join(validationPlan.Root, "go-mod")
	validationPlan.Terraform = filepath.Join(validationPlan.Root, "terraform")
	validationPlan.PackageManager = filepath.Join(validationPlan.Root, "package-manager")
	validation, err := executeSample(options, tools, runner, configuration, validationPlan, "validation", "cold", 0)
	if err != nil {
		return err
	}
	if !validationAllowsMeasurement(options.profile, options.variant, validation) {
		return nil
	}
	if expectedBaselineDigest != "" && validation.OutputTreeSHA256 != expectedBaselineDigest {
		return nil
	}
	if options.baselineDir != "" {
		baselineDigest, err := findBaselineDigest(options.baselineDir, options.workload)
		if err != nil {
			return err
		}
		if validation.OutputTreeSHA256 != baselineDigest {
			return nil
		}
	}

	for _, state := range []string{"cold", "warm"} {
		if state == "warm" {
			plan, planErr := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, options.variant, state, 1)
			if planErr != nil {
				return planErr
			}
			if err := executeSeed(options, configuration, plan); err != nil {
				return fmt.Errorf("warm cache seed: %w", err)
			}
		}
		for sample := 1; sample <= 5; sample++ {
			plan, planErr := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, options.variant, state, sample)
			if planErr != nil {
				return planErr
			}
			receipt, sampleErr := executeSample(options, tools, runner, configuration, plan, "measured", state, sample)
			if sampleErr != nil {
				return sampleErr
			}
			if receipt.Exit.Code != 0 {
				return nil
			}
		}
	}
	return nil
}

func baselineProbeOptions(options runOptions) runOptions {
	probe := options
	probe.variant = "d8-current"
	return probe
}

func baselineProbeProvidesDigest(receipt workloadbench.Receipt) bool {
	return receipt.Exit.Code == 0
}

func validationAllowsMeasurement(profile, variant string, receipt workloadbench.Receipt) bool {
	if receipt.Exit.Code != 0 || receipt.Metrics.OOMEvents != 0 {
		return false
	}
	// Always establish complete D8-current distributions. A resource-ineligible
	// baseline must yield a reproducible negative aggregate rather than prevent
	// the aggregate from reporting why no compute route may be selected.
	return profile == "d8" && variant == "d8-current" || receipt.Metrics.PeakMemoryRatio < 0.80
}

func runExampleSuiteCommand(args []string) error {
	set := flag.NewFlagSet("run-example-suite", flag.ContinueOnError)
	options := runOptions{workload: "terraform-example-validation"}
	set.StringVar(&options.profile, "profile", "", "fixed profile: d8 or d16")
	set.StringVar(&options.commit, "commit", "", "exact pull request head")
	set.StringVar(&options.sourceRoot, "source-root", ".", "source repository")
	set.StringVar(&options.specSource, "spec-source", "docs/specifications/api", "verified specification directory")
	set.StringVar(&options.cacheRoot, "cache-root", "", "run-local cache root")
	set.StringVar(&options.outputDir, "output-dir", "", "sanitized receipt directory")
	set.StringVar(&options.baselineDir, "baseline-dir", "", "D8 validation receipt directory")
	if err := set.Parse(args); err != nil {
		return err
	}
	if set.NArg() != 0 || (options.profile != "d8" && options.profile != "d16") {
		return errors.New("example suite accepts only a fixed d8 or d16 profile")
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(options.commit) || options.cacheRoot == "" || options.outputDir == "" {
		return errors.New("example suite requires an exact commit and run-local cache/output roots")
	}
	if options.profile == "d16" && options.baselineDir == "" {
		return errors.New("d16 example suite requires D8 baseline receipts")
	}
	var err error
	options.sourceRoot, err = filepath.Abs(options.sourceRoot)
	if err != nil {
		return err
	}
	options.specSource, err = filepath.Abs(options.specSource)
	if err != nil {
		return err
	}
	if err := requireDirectory(options.specSource); err != nil {
		return fmt.Errorf("verified specification bundle: %w", err)
	}
	if err := os.MkdirAll(options.outputDir, 0o755); err != nil {
		return err
	}
	tools, err := inspectToolchain(options.sourceRoot)
	if err != nil {
		return err
	}
	runner, err := attestRunner(options.profile)
	if err != nil {
		return err
	}

	variants := []string{options.profile + "-current", options.profile + "-n1", options.profile + "-n2", options.profile + "-n4", options.profile + "-n8"}
	if options.profile == "d16" {
		variants = append(variants, "d16-n16")
	}
	baselineDigest := ""
	if options.profile == "d16" {
		baselineDigest, err = findBaselineDigest(options.baselineDir, options.workload)
		if err != nil {
			return err
		}
	}

	type screened struct {
		configuration workloadbench.Configuration
		duration      float64
	}
	var candidates []screened
	for _, variant := range variants {
		configuration, parseErr := parseVariant(options.profile, options.workload, variant)
		if parseErr != nil {
			return parseErr
		}
		probeOptions := options
		probeOptions.variant = variant
		plan, planErr := validationCache(options, variant)
		if planErr != nil {
			return planErr
		}
		receipt, probeErr := executeSample(probeOptions, tools, runner, configuration, plan, "validation", "cold", 0)
		if probeErr != nil {
			return probeErr
		}
		if options.profile == "d8" && variant == "d8-current" {
			baselineDigest = receipt.OutputTreeSHA256
		}
		if receipt.Exit.Code == 0 && receipt.Metrics.PeakMemoryRatio < 0.80 && receipt.Metrics.OOMEvents == 0 && receipt.OutputTreeSHA256 == baselineDigest {
			candidates = append(candidates, screened{configuration: configuration, duration: receipt.Metrics.ElapsedSeconds})
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].duration != candidates[j].duration {
			return candidates[i].duration < candidates[j].duration
		}
		if candidates[i].configuration.GOMAXPROCS != candidates[j].configuration.GOMAXPROCS {
			return candidates[i].configuration.GOMAXPROCS < candidates[j].configuration.GOMAXPROCS
		}
		return candidates[i].configuration.VariantID < candidates[j].configuration.VariantID
	})
	selected := candidates[0].configuration

	if options.profile == "d8" {
		baselineConfiguration, _ := parseVariant("d8", options.workload, "d8-current")
		baselineOptions := options
		baselineOptions.variant = baselineConfiguration.VariantID
		if err := measureSamples(baselineOptions, tools, runner, baselineConfiguration); err != nil {
			return err
		}
	}
	for _, concurrency := range []int{1, 2, 4, 8} {
		configuration := selected
		if concurrency > 1 {
			configuration.VariantID += "-examples-c" + strconv.Itoa(concurrency)
		}
		configuration.ExampleConcurrency = concurrency
		if options.profile == "d8" && configuration.VariantID == "d8-current" {
			continue
		}
		variantOptions := options
		variantOptions.variant = configuration.VariantID
		plan, planErr := validationCache(options, configuration.VariantID)
		if planErr != nil {
			return planErr
		}
		validation, validationErr := executeSample(variantOptions, tools, runner, configuration, plan, "validation", "cold", 0)
		if validationErr != nil {
			return validationErr
		}
		if validation.Exit.Code != 0 || validation.Metrics.PeakMemoryRatio >= 0.80 || validation.Metrics.OOMEvents != 0 || validation.OutputTreeSHA256 != baselineDigest {
			continue
		}
		if err := measureSamples(variantOptions, tools, runner, configuration); err != nil {
			return err
		}
	}
	return nil
}

func validationCache(options runOptions, variant string) (workloadbench.CachePlan, error) {
	plan, err := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, variant, "cold", 1)
	if err != nil {
		return workloadbench.CachePlan{}, err
	}
	plan.Root = filepath.Join(options.cacheRoot, options.profile, options.workload, variant, "validation")
	plan.GoBuild = filepath.Join(plan.Root, "go-build")
	plan.GoModules = filepath.Join(plan.Root, "go-mod")
	plan.Terraform = filepath.Join(plan.Root, "terraform")
	plan.PackageManager = filepath.Join(plan.Root, "package-manager")
	return plan, nil
}

func measureSamples(options runOptions, tools toolIdentity, runner workloadbench.RunnerAttestation, configuration workloadbench.Configuration) error {
	for _, state := range []string{"cold", "warm"} {
		if state == "warm" {
			plan, err := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, options.variant, state, 1)
			if err != nil {
				return err
			}
			if err := executeSeed(options, configuration, plan); err != nil {
				return fmt.Errorf("warm cache seed: %w", err)
			}
		}
		for sample := 1; sample <= 5; sample++ {
			plan, err := workloadbench.PlanCache(options.cacheRoot, options.profile, options.workload, options.variant, state, sample)
			if err != nil {
				return err
			}
			receipt, err := executeSample(options, tools, runner, configuration, plan, "measured", state, sample)
			if err != nil {
				return err
			}
			if receipt.Exit.Code != 0 {
				return nil
			}
		}
	}
	return nil
}

func parseVariant(profile, workload, variant string) (workloadbench.Configuration, error) {
	matches := variantPattern.FindStringSubmatch(variant)
	if len(matches) == 0 || matches[1] != profile {
		return workloadbench.Configuration{}, errors.New("variant is not a fixed member of its runner profile")
	}
	configuration := workloadbench.Configuration{VariantID: variant, ExampleConcurrency: 1}
	if matches[2] != "current" {
		workers, _ := strconv.Atoi(matches[3])
		if profile == "d8" && workers == 16 {
			return workloadbench.Configuration{}, errors.New("d8 does not permit n16")
		}
		configuration.GOMAXPROCS = workers
		configuration.GOFLAGS = "-p=" + strconv.Itoa(workers)
	}
	if matches[4] != "" {
		if workload != "terraform-example-validation" {
			return workloadbench.Configuration{}, errors.New("example concurrency is restricted to Terraform example validation")
		}
		configuration.ExampleConcurrency, _ = strconv.Atoi(matches[4])
	}
	return configuration, nil
}

func executeSeed(options runOptions, configuration workloadbench.Configuration, cache workloadbench.CachePlan) error {
	if err := prepareCache(cache); err != nil {
		return err
	}
	worktree, cleanup, err := detachedWorktree(options.sourceRoot, options.commit, options.specSource)
	if err != nil {
		return err
	}
	defer cleanup()
	argv, _ := workloadbench.FixedArgv(options.workload)
	command := exec.Command("bash", filepath.Join(worktree, argv[0]), argv[1])
	command.Dir = worktree
	command.Env = benchmarkEnvironment(configuration, cache)
	return command.Run()
}

func executeSample(options runOptions, tools toolIdentity, runner workloadbench.RunnerAttestation, configuration workloadbench.Configuration, cache workloadbench.CachePlan, kind, state string, sample int) (workloadbench.Receipt, error) {
	worktree, cleanup, err := detachedWorktree(options.sourceRoot, options.commit, options.specSource)
	if err != nil {
		return workloadbench.Receipt{}, err
	}
	defer cleanup()
	if err := prepareCache(cache); err != nil {
		return workloadbench.Receipt{}, err
	}
	argv, _ := workloadbench.FixedArgv(options.workload)
	rawPath := filepath.Join(options.outputDir, ".raw-"+options.profile+"-"+options.workload+"-"+options.variant+"-"+kind+"-"+state+"-"+strconv.Itoa(sample)+".json")
	pairID := strings.Join([]string{options.workload, state, strconv.Itoa(sample)}, "/")
	profileArgs := []string{"--name", options.workload, "--output", rawPath, "--cache-state", state, "--variant", options.variant, "--pair-id", pairID, "--commit", options.commit, "--"}
	profileArgs = append(profileArgs, filepath.Join(worktree, argv[0]))
	profileArgs = append(profileArgs, argv[1:]...)
	command := exec.Command("runner-profile", profileArgs...)
	command.Dir = worktree
	command.Env = benchmarkEnvironment(configuration, cache)
	runErr := command.Run()
	raw, err := readRawProfile(rawPath)
	_ = os.Remove(rawPath)
	if err != nil {
		return workloadbench.Receipt{}, fmt.Errorf("read runner measurement: %w", err)
	}
	if raw.RunnerProfile != runner.Profile {
		return workloadbench.Receipt{}, errors.New("measured runner profile differs from the attested profile")
	}
	digest, digestErr := workloadbench.CanonicalTreeDigest(worktree)
	category := "none"
	if runErr != nil {
		category = "command_failed"
	}
	if digestErr != nil {
		category = "digest_failed"
		digest = "sha256:" + strings.Repeat("0", 64)
	}
	peak, limit, peakRatio := int64(0), int64(0), float64(0)
	if raw.Memory.PeakBytes != nil {
		peak = *raw.Memory.PeakBytes
	}
	if raw.Memory.LimitBytes != nil {
		limit = *raw.Memory.LimitBytes
	}
	if raw.Memory.PeakLimitRatio != nil {
		peakRatio = *raw.Memory.PeakLimitRatio
	}
	receipt := workloadbench.Receipt{
		SchemaVersion:    workloadbench.ReceiptSchemaVersion,
		Identity:         workloadbench.Identity{Repository: raw.Repository, Commit: raw.Commit, RunID: raw.RunID, JobID: raw.JobID, RunAttempt: raw.RunAttempt, PairID: pairID},
		Toolchain:        workloadbench.Toolchain{SpecRelease: tools.specRelease, SpecPinSHA256: tools.specPin, GoVersion: tools.goVersion, TerraformVersion: tools.terraformVersion, DocumentationTool: tools.documentationTool, ActionCheckout: checkoutActionSHA, RunnerImage: raw.ImageDigest},
		Runner:           runner,
		Workload:         workloadbench.WorkloadIdentity{ID: options.workload, Argv: argv},
		Configuration:    configuration,
		Cache:            workloadbench.CacheIdentity{State: state, Key: cacheKey(options, state, sample)},
		Metrics:          workloadbench.Metrics{ElapsedSeconds: raw.Duration, CPUUsageUsec: raw.CPU.UsageUsec, CPUUserUsec: raw.CPU.UserUsec, CPUSystemUsec: raw.CPU.SystemUsec, CPUThrottledUsec: raw.CPU.ThrottledUsec, PeakMemoryBytes: peak, MemoryLimitBytes: limit, PeakMemoryRatio: peakRatio, OOMEvents: raw.Memory.Events["oom"] + raw.Memory.Events["oom_kill"]},
		OutputTreeSHA256: digest,
		Exit:             workloadbench.Exit{Code: raw.Exit.Code, ErrorCategory: category},
		SampleKind:       kind, SampleIndex: sample,
	}
	encoded, err := workloadbench.MarshalReceipt(receipt)
	if err != nil {
		return workloadbench.Receipt{}, err
	}
	name := strings.Join([]string{options.profile, options.workload, options.variant, kind, state, strconv.Itoa(sample)}, "-") + ".json"
	if err := os.WriteFile(filepath.Join(options.outputDir, name), encoded, 0o600); err != nil {
		return workloadbench.Receipt{}, err
	}
	return receipt, nil
}

func cacheKey(options runOptions, state string, sample int) string {
	leaf := strconv.Itoa(sample)
	if state == "warm" {
		leaf = "shared"
	}
	return strings.Join([]string{options.profile, options.workload, options.variant, state, leaf}, "/")
}

func benchmarkEnvironment(configuration workloadbench.Configuration, cache workloadbench.CachePlan) []string {
	environment := append([]string{}, os.Environ()...)
	environment = setEnvironment(environment, "GOCACHE", cache.GoBuild)
	environment = setEnvironment(environment, "GOMODCACHE", cache.GoModules)
	environment = setEnvironment(environment, "TF_DATA_DIR", cache.Terraform)
	environment = setEnvironment(environment, "TF_PLUGIN_CACHE_DIR", filepath.Join(cache.Terraform, "plugins"))
	environment = setEnvironment(environment, "npm_config_cache", cache.PackageManager)
	environment = setEnvironment(environment, "EXAMPLE_WORKERS", strconv.Itoa(configuration.ExampleConcurrency))
	environment = setEnvironment(environment, "GOFLAGS", configuration.GOFLAGS)
	if configuration.GOMAXPROCS == 0 {
		environment = removeEnvironment(environment, "GOMAXPROCS")
	} else {
		environment = setEnvironment(environment, "GOMAXPROCS", strconv.Itoa(configuration.GOMAXPROCS))
	}
	return environment
}

func setEnvironment(environment []string, name, value string) []string {
	environment = removeEnvironment(environment, name)
	return append(environment, name+"="+value)
}

func removeEnvironment(environment []string, name string) []string {
	prefix := name + "="
	result := environment[:0]
	for _, item := range environment {
		if !strings.HasPrefix(item, prefix) {
			result = append(result, item)
		}
	}
	return result
}

func prepareCache(cache workloadbench.CachePlan) error {
	for _, directory := range []string{cache.GoBuild, cache.GoModules, cache.Terraform, filepath.Join(cache.Terraform, "plugins"), cache.PackageManager} {
		if err := os.MkdirAll(directory, 0o755); err != nil {
			return err
		}
	}
	return nil
}

func detachedWorktree(sourceRoot, commit, specSource string) (string, func(), error) {
	parent := filepath.Join(os.TempDir(), "xcsh-benchmark-worktrees")
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return "", nil, err
	}
	worktree, err := os.MkdirTemp(parent, "sample-")
	if err != nil {
		return "", nil, err
	}
	if err := os.Remove(worktree); err != nil {
		return "", nil, err
	}
	command := exec.Command("git", "-C", sourceRoot, "worktree", "add", "--detach", worktree, commit)
	if output, err := command.CombinedOutput(); err != nil {
		return "", nil, fmt.Errorf("create detached worktree: %w: %s", err, strings.TrimSpace(string(output)))
	}
	cleanup := func() {
		_ = os.RemoveAll(filepath.Join(worktree, "docs", "specifications", "api"))
		_ = exec.Command("git", "-C", sourceRoot, "worktree", "remove", "--force", worktree).Run()
	}
	destination := filepath.Join(worktree, "docs", "specifications", "api")
	if err := copyTree(specSource, destination); err != nil {
		cleanup()
		return "", nil, err
	}
	return worktree, cleanup, nil
}

func copyTree(source, destination string) error {
	return filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		return copyFile(path, target, info.Mode().Perm())
	})
}

func copyFile(source, destination string, mode fs.FileMode) error {
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	output, err := os.OpenFile(destination, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		_ = input.Close()
		return err
	}
	_, copyErr := io.Copy(output, input)
	inputCloseErr := input.Close()
	outputCloseErr := output.Close()
	if copyErr != nil {
		return copyErr
	}
	if inputCloseErr != nil {
		return inputCloseErr
	}
	return outputCloseErr
}

func inspectToolchain(root string) (toolIdentity, error) {
	pinPath := filepath.Join(root, "tools", "spec-release.json")
	pin, err := os.ReadFile(pinPath)
	if err != nil {
		return toolIdentity{}, err
	}
	var release struct {
		ReleaseTag string `json:"release_tag"`
	}
	if err := json.Unmarshal(pin, &release); err != nil || release.ReleaseTag == "" {
		return toolIdentity{}, errors.New("spec release pin is malformed")
	}
	pinHash := sha256.Sum256(pin)
	goVersion, err := commandOutput(root, "go", "env", "GOVERSION")
	if err != nil {
		return toolIdentity{}, err
	}
	terraJSON, err := commandOutput(root, "terraform", "version", "-json")
	if err != nil {
		return toolIdentity{}, err
	}
	var terraform struct {
		Version string `json:"terraform_version"`
	}
	if err := json.Unmarshal([]byte(terraJSON), &terraform); err != nil || terraform.Version == "" {
		return toolIdentity{}, errors.New("Terraform identity is malformed")
	}
	docPath, err := exec.LookPath("tfplugindocs")
	if err != nil {
		return toolIdentity{}, err
	}
	docBuild, err := commandOutput(root, "go", "version", "-m", docPath)
	if err != nil {
		return toolIdentity{}, err
	}
	docIdentity := ""
	for _, line := range strings.Split(docBuild, "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == "mod" {
			docIdentity = fields[1] + "@" + fields[2]
			break
		}
	}
	if docIdentity != "github.com/hashicorp/terraform-plugin-docs@v0.25.0+dirty" {
		return toolIdentity{}, fmt.Errorf("unexpected documentation tool %q", docIdentity)
	}
	return toolIdentity{specRelease: release.ReleaseTag, specPin: "sha256:" + hex.EncodeToString(pinHash[:]), goVersion: goVersion, terraformVersion: terraform.Version, documentationTool: docIdentity}, nil
}

func commandOutput(directory, name string, args ...string) (string, error) {
	var command *exec.Cmd
	switch name {
	case "go":
		command = exec.Command("go", args...)
	case "terraform":
		command = exec.Command("terraform", args...)
	default:
		return "", fmt.Errorf("identity command %q is not fixed", name)
	}
	command.Dir = directory
	output, err := command.Output()
	if err != nil {
		return "", fmt.Errorf("inspect %s identity: %w", name, err)
	}
	return strings.TrimSpace(string(output)), nil
}

func attestRunner(profile string) (workloadbench.RunnerAttestation, error) {
	expected, ok := workloadbench.FixedRunner(profile)
	if !ok {
		return workloadbench.RunnerAttestation{}, errors.New("profile must be d8 or d16")
	}
	if os.Getenv("RUNNER_PROFILE") != expected.Profile {
		return workloadbench.RunnerAttestation{}, errors.New("observed runner profile mismatch")
	}
	image := os.Getenv("RUNNER_IMAGE_DIGEST")
	if strings.Contains(image, "@") {
		image = image[strings.LastIndex(image, "@")+1:]
	}
	if image != "sha256:2a0243be5404daa0f52bae16384f53dbc04554e31406ed0db45152d92f6187e1" {
		return workloadbench.RunnerAttestation{}, errors.New("observed runner image mismatch")
	}
	cgroup, err := currentCgroup()
	if err != nil {
		return workloadbench.RunnerAttestation{}, err
	}
	cpuFields, err := fieldsFromFile(filepath.Join(cgroup, "cpu.max"))
	if err != nil || len(cpuFields) != 2 || cpuFields[0] == "max" {
		return workloadbench.RunnerAttestation{}, errors.New("CPU cgroup limit is unavailable")
	}
	quota, quotaErr := strconv.Atoi(cpuFields[0])
	period, periodErr := strconv.Atoi(cpuFields[1])
	if quotaErr != nil || periodErr != nil || period == 0 || quota/period != expected.CPULimit {
		return workloadbench.RunnerAttestation{}, errors.New("CPU cgroup limit mismatch")
	}
	memoryFields, err := fieldsFromFile(filepath.Join(cgroup, "memory.max"))
	if err != nil || len(memoryFields) != 1 {
		return workloadbench.RunnerAttestation{}, errors.New("memory cgroup limit is unavailable")
	}
	memory, err := strconv.ParseInt(memoryFields[0], 10, 64)
	if err != nil || memory != expected.MemoryLimit {
		return workloadbench.RunnerAttestation{}, errors.New("memory cgroup limit mismatch")
	}
	for _, socket := range []string{"/var/run/docker.sock", "/run/docker.sock", "/run/containerd/containerd.sock"} {
		if _, err := os.Lstat(socket); err == nil || !errors.Is(err, os.ErrNotExist) {
			return workloadbench.RunnerAttestation{}, errors.New("container socket is visible")
		}
	}
	vmSize, err := azureVMSize()
	if err != nil || vmSize != expected.VMSize {
		return workloadbench.RunnerAttestation{}, errors.New("Azure VM size mismatch or unavailable")
	}
	return expected, nil
}

func currentCgroup() (string, error) {
	content, err := os.ReadFile("/proc/self/cgroup")
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.SplitN(line, ":", 3)
		if len(parts) == 3 && parts[0] == "0" {
			return filepath.Join("/sys/fs/cgroup", strings.TrimPrefix(parts[2], "/")), nil
		}
	}
	return "", errors.New("unified cgroup path is unavailable")
}

func fieldsFromFile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	return strings.Fields(string(content)), err
}

func azureVMSize() (string, error) {
	command := exec.Command(
		"curl", "--fail", "--silent", "--show-error", "--max-time", "10",
		"--noproxy", "*", "-H", "Metadata:true",
		"http://169.254.169.254/metadata/instance/compute/vmSize?api-version=2021-02-01&format=text",
	)
	output, err := command.Output()
	if err != nil {
		return "", fmt.Errorf("query Azure VM size: %w", err)
	}
	if len(output) > 128 {
		return "", errors.New("Azure VM size response is oversized")
	}
	return strings.TrimSpace(string(output)), nil
}

func readRawProfile(path string) (rawProfile, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return rawProfile{}, err
	}
	var profile rawProfile
	if err := json.Unmarshal(payload, &profile); err != nil {
		return rawProfile{}, err
	}
	return profile, nil
}

func findBaselineDigest(root, workload string) (string, error) {
	var matches []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && strings.HasSuffix(path, ".json") {
			payload, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			receipt, err := workloadbench.ParseReceipt(payload)
			if err == nil && receipt.Workload.ID == workload && receipt.Configuration.VariantID == "d8-current" && receipt.SampleKind == "validation" {
				matches = append(matches, receipt.OutputTreeSHA256)
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if len(matches) != 1 {
		return "", fmt.Errorf("expected one D8 validation digest for %s, found %d", workload, len(matches))
	}
	return matches[0], nil
}

func aggregateCommand(args []string) error {
	set := flag.NewFlagSet("aggregate", flag.ContinueOnError)
	input, output, markdown := "", "", ""
	set.StringVar(&input, "input", "", "receipt directory")
	set.StringVar(&output, "output", "", "aggregate JSON")
	set.StringVar(&markdown, "markdown", "", "aggregate Markdown")
	if err := set.Parse(args); err != nil {
		return err
	}
	if set.NArg() != 0 || input == "" || output == "" || markdown == "" {
		return errors.New("aggregate requires fixed input, output, and markdown paths")
	}
	byWorkload := map[string][]workloadbench.Receipt{}
	err := filepath.WalkDir(input, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		payload, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		receipt, err := workloadbench.ParseReceipt(payload)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		byWorkload[receipt.Workload.ID] = append(byWorkload[receipt.Workload.ID], receipt)
		return nil
	})
	if err != nil {
		return err
	}
	workloads := make([]string, 0, len(byWorkload))
	for workload := range byWorkload {
		workloads = append(workloads, workload)
	}
	sort.Strings(workloads)
	results := make([]workloadbench.AggregateResult, 0, len(workloads))
	for _, workload := range workloads {
		result, err := workloadbench.Aggregate(byWorkload[workload], "d8-current")
		if err != nil {
			return fmt.Errorf("aggregate %s: %w", workload, err)
		}
		results = append(results, result)
	}
	receiptBundleDigest, err := workloadbench.CanonicalTreeDigest(input)
	if err != nil {
		return fmt.Errorf("digest receipt bundle: %w", err)
	}
	payload, err := json.MarshalIndent(map[string]any{"schema_version": 1, "receipt_bundle_sha256": receiptBundleDigest, "workloads": results}, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	if err := os.WriteFile(output, payload, 0o600); err != nil {
		return err
	}
	hash := sha256.Sum256(payload)
	var summary strings.Builder
	summary.WriteString("## Provider workload benchmark\n\n")
	summary.WriteString("Receipt bundle SHA-256: `" + receiptBundleDigest + "`\n\n")
	summary.WriteString("Aggregate SHA-256: `sha256:" + hex.EncodeToString(hash[:]) + "`\n\n")
	summary.WriteString("| Workload | Winner | Production compute eligible |\n|---|---|---|\n")
	for _, result := range results {
		eligible := strings.HasPrefix(result.Winner, "d16-")
		winner := result.Winner
		if winner == "" {
			winner = "none"
		}
		summary.WriteString(fmt.Sprintf("| `%s` | `%s` | `%t` |\n", result.Workload, winner, eligible))
	}
	return os.WriteFile(markdown, []byte(summary.String()), 0o600)
}

func requireDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("path is not a directory")
	}
	return nil
}
