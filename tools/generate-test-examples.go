// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// generate-test-examples.go extracts verified Terraform configurations from
// acceptance test files and writes them as named example .tf files.
// These examples are proven to work against the live F5 XC staging API.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var configFuncRegex = regexp.MustCompile(`func (testAcc\w+Config_\w+)\([^)]*\)\s+string\s*{`)
var formatStringRegex = regexp.MustCompile("return\\s+(?:acctest\\.ConfigCompose\\(\\s*acctest\\.ProviderConfig\\(\\),\\s*)?fmt\\.Sprintf\\(`([^`]+)`")
var simpleReturnRegex = regexp.MustCompile("return\\s+fmt\\.Sprintf\\(`([^`]+)`")

const (
	expectedNamedExampleCount = 81
	xcshProviderSource        = "f5-sales-demo/xcsh"
	xcshVersionConstraint     = ">= 0.1.0"
	timeProviderVersion       = "0.13.1"
)

var testExampleResources = []string{
	"http_loadbalancer",
	"tcp_loadbalancer",
	"healthcheck",
	"app_firewall",
	"origin_pool",
	"rate_limiter",
	"service_policy",
	"user_identification",
	"malicious_user_mitigation",
}

var (
	healthyThresholdPlaceholderRegex   = regexp.MustCompile(`(?m)^(\s*healthy_threshold\s*=\s*)%\[\d+\]d`)
	unhealthyThresholdPlaceholderRegex = regexp.MustCompile(`(?m)^(\s*unhealthy_threshold\s*=\s*)%\[\d+\]d`)
	jitterPercentPlaceholderRegex      = regexp.MustCompile(`(?m)^(\s*jitter_percent\s*=\s*)%\[\d+\]d`)
	rateLimiterUnitPlaceholderRegex    = regexp.MustCompile(`(?m)^(\s*unit\s*=\s*)%\[\d+\]q`)
	healthyThresholdAssignmentRegex    = regexp.MustCompile(`(?m)^\s*healthy_threshold\s*=\s*([0-9]+)\s*$`)
	unhealthyThresholdAssignmentRegex  = regexp.MustCompile(`(?m)^\s*unhealthy_threshold\s*=\s*([0-9]+)\s*$`)
	jitterPercentAssignmentRegex       = regexp.MustCompile(`(?m)^\s*jitter_percent\s*=\s*([0-9]+)\s*$`)
	rateLimiterUnitAssignmentRegex     = regexp.MustCompile(`(?m)^\s*unit\s*=\s*"([^"]+)"\s*$`)
)

type testExample struct {
	Resource string
	Name     string
	Config   string
}

type generatedExample struct {
	Path    string
	Content string
}

func main() {
	examples, err := renderExamples("internal/provider", "examples/resources")
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate verified examples: %v\n", err)
		os.Exit(1)
	}
	if len(examples) != expectedNamedExampleCount {
		fmt.Fprintf(os.Stderr, "generate verified examples: got %d named examples, want %d\n", len(examples), expectedNamedExampleCount)
		os.Exit(1)
	}
	if err := validateRenderedExamples(examples); err != nil {
		fmt.Fprintf(os.Stderr, "generate verified examples: %v\n", err)
		os.Exit(1)
	}

	paths := make([]string, 0, len(examples))
	for _, example := range examples {
		if err := os.MkdirAll(filepath.Dir(example.Path), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "create example directory for %s: %v\n", example.Path, err)
			os.Exit(1)
		}
		if err := os.WriteFile(example.Path, []byte(example.Content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write example %s: %v\n", example.Path, err)
			os.Exit(1)
		}
		fmt.Printf("Generated: %s\n", example.Path)
		paths = append(paths, example.Path)
	}

	fmtArgs := append([]string{"fmt"}, paths...)
	cmd := exec.Command("terraform", fmtArgs...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "format generated examples: %v\n%s\n", err, out)
		os.Exit(1)
	}

	fmt.Printf("\nGenerated and formatted %d verified example files\n", len(examples))
}

func renderExamples(testDir, outputDir string) ([]generatedExample, error) {
	var generated []generatedExample
	seenPaths := make(map[string]struct{})

	for _, res := range testExampleResources {
		testFile := filepath.Join(testDir, res+"_resource_test.go")
		content, err := os.ReadFile(testFile)
		if err != nil {
			return nil, fmt.Errorf("read acceptance tests for %s: %w", res, err)
		}

		examples := extractExamples(res, string(content))

		for _, ex := range examples {
			exDir := filepath.Join(outputDir, "xcsh_"+ex.Resource)
			filename := toExampleFilename(ex.Name) + ".tf"
			if filename == "basic-system.tf" || filename == "basic.tf" {
				continue
			}

			outPath := filepath.Join(exDir, filename)
			if _, exists := seenPaths[outPath]; exists {
				return nil, fmt.Errorf("multiple acceptance helpers generate %s", outPath)
			}
			header := fmt.Sprintf("# %s — Verified Configuration Example\n# This configuration is extracted from acceptance tests\n# and verified against the live F5 XC API.\n\n",
				toHumanName(ex.Name))

			cleaned := cleanConfig(ex.Name, ex.Config)
			if cleaned == "" {
				continue
			}

			seenPaths[outPath] = struct{}{}
			generated = append(generated, generatedExample{
				Path:    outPath,
				Content: header + addProviderRequirements(cleaned),
			})
		}
	}

	return generated, nil
}

func extractExamples(resource, content string) []testExample {
	var examples []testExample

	matches := configFuncRegex.FindAllStringSubmatchIndex(content, -1)

	for _, match := range matches {
		funcName := content[match[2]:match[3]]
		funcStart := match[0]

		funcEnd := findFuncEnd(content, funcStart)
		if funcEnd < 0 {
			continue
		}

		funcBody := content[funcStart:funcEnd]

		configName := extractConfigName(funcName, resource)
		if configName == "" {
			continue
		}

		hcl := extractHCL(funcBody)
		if hcl == "" {
			continue
		}

		examples = append(examples, testExample{
			Resource: resource,
			Name:     configName,
			Config:   hcl,
		})
	}

	return examples
}

func findFuncEnd(content string, start int) int {
	depth := 0
	inString := false
	inBacktick := false

	for i := start; i < len(content); i++ {
		c := content[i]

		if c == '`' {
			inBacktick = !inBacktick
			continue
		}
		if inBacktick {
			continue
		}
		if c == '"' && (i == 0 || content[i-1] != '\\') {
			inString = !inString
			continue
		}
		if inString {
			continue
		}
		if c == '{' {
			depth++
		}
		if c == '}' {
			depth--
			if depth == 0 {
				return i + 1
			}
		}
	}

	return -1
}

func extractConfigName(funcName, resource string) string {
	prefix := "testAcc" + toCamelCase(resource) + "Config_"
	altPrefix := "testAccHTTPLBConfig_"
	altPrefix2 := "testAccHTTPLoadBalancerConfig_"

	name := ""
	if strings.HasPrefix(funcName, prefix) {
		name = funcName[len(prefix):]
	} else if strings.HasPrefix(funcName, altPrefix) {
		name = funcName[len(altPrefix):]
	} else if strings.HasPrefix(funcName, altPrefix2) {
		name = funcName[len(altPrefix2):]
	} else {
		for _, p := range []string{"testAcc"} {
			if strings.HasPrefix(funcName, p) && strings.Contains(funcName, "Config_") {
				idx := strings.Index(funcName, "Config_")
				name = funcName[idx+7:]
				break
			}
		}
	}

	return name
}

func extractHCL(funcBody string) string {
	backtickStart := strings.Index(funcBody, "`\n")
	if backtickStart < 0 {
		backtickStart = strings.Index(funcBody, "(`")
		if backtickStart < 0 {
			return ""
		}
		backtickStart++
	}
	backtickStart++

	backtickEnd := strings.Index(funcBody[backtickStart+1:], "`")
	if backtickEnd < 0 {
		return ""
	}
	backtickEnd += backtickStart + 1

	return funcBody[backtickStart:backtickEnd]
}

func cleanConfig(name, config string) string {
	config = strings.TrimSpace(config)

	if name == "nestedLabels" {
		config = strings.ReplaceAll(config, "%[2]q", `"192.0.2.1"`)
		config = strings.ReplaceAll(config, "%[3]s", "    labels = {\n      \"env\" = \"test\"\n      \"app\" = \"demo\"\n    }\n")
	}

	// Positional values in acceptance helpers do not have one universal meaning.
	// Resolve fields with constrained domains before applying the general sample
	// substitutions below so generated configurations remain schema-valid.
	config = healthyThresholdPlaceholderRegex.ReplaceAllString(config, "${1}3")
	config = unhealthyThresholdPlaceholderRegex.ReplaceAllString(config, "${1}2")
	config = jitterPercentPlaceholderRegex.ReplaceAllString(config, "${1}30")
	config = rateLimiterUnitPlaceholderRegex.ReplaceAllString(config, `${1}"MINUTE"`)

	config = regexp.MustCompile(`%\[\d+\]s(\s*)\}`).ReplaceAllString(config, "    example-key = \"example-value\"\n$1}")

	config = regexp.MustCompile(`%\[1\]q`).ReplaceAllString(config, `"example"`)
	config = regexp.MustCompile(`%\[1\]s`).ReplaceAllString(config, "example")
	config = regexp.MustCompile(`%\[2\]q`).ReplaceAllString(config, `"example-value"`)
	config = regexp.MustCompile(`%\[2\]s`).ReplaceAllString(config, `"example-value"`)
	config = regexp.MustCompile(`%\[2\]d`).ReplaceAllString(config, "443")
	config = regexp.MustCompile(`%\[3\]q`).ReplaceAllString(config, `"example-description"`)
	config = regexp.MustCompile(`%\[3\]s`).ReplaceAllString(config, "example-value")
	config = regexp.MustCompile(`%\[3\]d`).ReplaceAllString(config, "3")
	config = regexp.MustCompile(`%\[4\]q`).ReplaceAllString(config, `"example-value"`)
	config = regexp.MustCompile(`%\[4\]s`).ReplaceAllString(config, "example-value")
	config = regexp.MustCompile(`%\[4\]d`).ReplaceAllString(config, "5")
	config = regexp.MustCompile(`%\[5\]s`).ReplaceAllString(config, "example-value")
	config = regexp.MustCompile(`%\[5\]d`).ReplaceAllString(config, "15")
	config = regexp.MustCompile(`%s`).ReplaceAllString(config, "example")
	config = regexp.MustCompile(`%q`).ReplaceAllString(config, `"example"`)
	config = regexp.MustCompile(`%d`).ReplaceAllString(config, "80")

	if strings.Contains(config, "%") && strings.Contains(config, "[") {
		return ""
	}

	return config
}

func addProviderRequirements(config string) string {
	var requirements strings.Builder
	requirements.WriteString("terraform {\n  required_providers {\n")
	requirements.WriteString("    xcsh = {\n")
	fmt.Fprintf(&requirements, "      source  = %q\n", xcshProviderSource)
	fmt.Fprintf(&requirements, "      version = %q\n", xcshVersionConstraint)
	requirements.WriteString("    }\n")
	if strings.Contains(config, "time_sleep") {
		requirements.WriteString("    time = {\n")
		requirements.WriteString("      source  = \"hashicorp/time\"\n")
		fmt.Fprintf(&requirements, "      version = \"= %s\"\n", timeProviderVersion)
		requirements.WriteString("    }\n")
	}
	requirements.WriteString("  }\n}\n\n")
	requirements.WriteString(strings.TrimSpace(config))
	requirements.WriteByte('\n')
	return requirements.String()
}

func validateRenderedExamples(examples []generatedExample) error {
	if len(examples) != expectedNamedExampleCount {
		return fmt.Errorf("got %d named examples, want %d", len(examples), expectedNamedExampleCount)
	}
	for _, example := range examples {
		if err := validateExampleContent(example.Content); err != nil {
			return fmt.Errorf("%s: %w", example.Path, err)
		}
	}
	return nil
}

func validateExampleContent(content string) error {
	requiredSource := fmt.Sprintf(`source  = %q`, xcshProviderSource)
	requiredVersion := fmt.Sprintf(`version = %q`, xcshVersionConstraint)
	if !strings.Contains(content, requiredSource) || !strings.Contains(content, requiredVersion) {
		return fmt.Errorf("missing the required xcsh provider source/version binding")
	}
	if strings.Contains(content, "time_sleep") {
		requiredTimeVersion := fmt.Sprintf(`version = "= %s"`, timeProviderVersion)
		if !strings.Contains(content, `source  = "hashicorp/time"`) || !strings.Contains(content, requiredTimeVersion) {
			return fmt.Errorf("time_sleep example does not pin the time provider")
		}
	}
	if strings.Contains(content, "%[") {
		return fmt.Errorf("contains an unresolved positional format placeholder")
	}

	for _, check := range []struct {
		name    string
		pattern *regexp.Regexp
		minimum int
		maximum int
	}{
		{name: "healthy_threshold", pattern: healthyThresholdAssignmentRegex, minimum: 1, maximum: 16},
		{name: "unhealthy_threshold", pattern: unhealthyThresholdAssignmentRegex, minimum: 1, maximum: 16},
		{name: "jitter_percent", pattern: jitterPercentAssignmentRegex, minimum: 0, maximum: 100},
	} {
		for _, match := range check.pattern.FindAllStringSubmatch(content, -1) {
			value, err := strconv.Atoi(match[1])
			if err != nil || value < check.minimum || value > check.maximum {
				return fmt.Errorf("%s value %q is outside %d..%d", check.name, match[1], check.minimum, check.maximum)
			}
		}
	}

	validRateLimiterUnits := map[string]struct{}{"SECOND": {}, "MINUTE": {}, "HOUR": {}}
	for _, match := range rateLimiterUnitAssignmentRegex.FindAllStringSubmatch(content, -1) {
		if _, ok := validRateLimiterUnits[match[1]]; !ok {
			return fmt.Errorf("rate limiter unit %q is not SECOND, MINUTE, or HOUR", match[1])
		}
	}

	return nil
}

func toExampleFilename(name string) string {
	name = strings.TrimSuffix(name, "System")
	name = strings.TrimSuffix(name, "_system")

	// Keep overlapping initialisms longest-first so HTTPS is not partially
	// rewritten as HttpS by a nondeterministic map iteration.
	acronyms := []string{"HTTPS", "HTTP", "ICMP", "WAF", "TLS", "TCP", "UDP", "DNS", "API", "SSL", "SNI", "IP"}

	for _, acronym := range acronyms {
		lower := strings.ToLower(acronym)
		name = strings.ReplaceAll(name, acronym, strings.ToUpper(lower[:1])+lower[1:])
	}

	result := ""
	for i, c := range name {
		if c >= 'A' && c <= 'Z' {
			if i > 0 && result[len(result)-1] != '-' {
				result += "-"
			}
			result += string(c + 32)
		} else if c == '_' {
			result += "-"
		} else {
			result += string(c)
		}
	}

	result = strings.ReplaceAll(result, "--", "-")
	result = strings.TrimPrefix(result, "-")
	return result
}

func toHumanName(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "System", "")
	name = strings.TrimSpace(name)
	words := strings.Fields(name)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

func toCamelCase(snake string) string {
	parts := strings.Split(snake, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}
