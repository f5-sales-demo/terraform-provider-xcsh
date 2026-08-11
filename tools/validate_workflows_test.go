// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const repositoryRunnerLabel = "terraform-provider-xcsh"

type jobContract struct {
	workflow    string
	job         string
	runsOn      []string
	environment string
	permissions map[string]string
	triggers    []string
	jobIf       *string
	secrets     []string
}

func strptr(value string) *string { return &value }

var protectedJobs = []jobContract{
	{"acc-tests.yml", "mock-tests", []string{"ubuntu-latest"}, "", map[string]string{"checks": "write", "contents": "read"}, []string{"pull_request", "schedule", "workflow_dispatch"}, nil, nil},
	{"acc-tests.yml", "real-api-tests", []string{"self-hosted", "Linux", "X64", repositoryRunnerLabel}, "acceptance-tests", map[string]string{"checks": "write", "contents": "read"}, []string{"pull_request", "schedule", "workflow_dispatch"}, strptr("always() &&\ngithub.event_name != 'pull_request' &&\n((github.event_name == 'schedule' &&\n  needs.mock-tests.result == 'success') ||\n (github.event_name == 'workflow_dispatch' &&\n  github.event.inputs.mode == 'full' &&\n  needs.mock-tests.result == 'success') ||\n (github.event_name == 'workflow_dispatch' &&\n  github.event.inputs.mode == 'real-only' &&\n  needs.mock-tests.result == 'skipped'))\n"), []string{"XCSH_API_TOKEN", "XCSH_API_URL"}},
	{"acc-tests.yml", "cleanup", []string{"self-hosted", "Linux", "X64", repositoryRunnerLabel}, "acceptance-tests", map[string]string{"contents": "read"}, []string{"pull_request", "schedule", "workflow_dispatch"}, strptr("always() &&\ngithub.event_name != 'pull_request' &&\nneeds.real-api-tests.result != 'skipped'\n"), []string{"XCSH_API_TOKEN", "XCSH_API_URL"}},
	{"discover-defaults.yml", "discover", []string{"self-hosted", "Linux", "X64", repositoryRunnerLabel}, "default-discovery", map[string]string{}, []string{"schedule", "workflow_dispatch"}, nil, []string{"REPO_SYNC_TOKEN", "XCSH_API_TOKEN", "XCSH_API_URL"}},
	{"sync-openapi.yml", "sync", []string{"ubuntu-latest"}, "provider-delivery", map[string]string{}, []string{"repository_dispatch", "workflow_dispatch"}, nil, []string{"GITHUB_TOKEN", "REPO_SYNC_TOKEN"}},
	{"on-merge.yml", "create-regeneration-pr", []string{"ubuntu-latest"}, "provider-delivery", map[string]string{"contents": "read"}, []string{"push", "workflow_dispatch"}, nil, []string{"REPO_SYNC_TOKEN"}},
	{"on-merge.yml", "receipt-spec-delivery", []string{"ubuntu-latest"}, "provider-delivery", map[string]string{"contents": "read"}, []string{"push", "workflow_dispatch"}, nil, []string{"REPO_SYNC_TOKEN"}},
	{"auto-merge.yml", "require-token", []string{"ubuntu-latest"}, "repository-settings", map[string]string{}, []string{"pull_request"}, nil, []string{"REPO_SYNC_TOKEN"}},
}

type workflowDocument struct {
	On          yaml.Node                 `yaml:"on"`
	Permissions map[string]string         `yaml:"permissions"`
	Jobs        map[string]map[string]any `yaml:"jobs"`
}

var secretReference = regexp.MustCompile(`\bsecrets\.([A-Za-z_][A-Za-z0-9_]*)\b`)

var environmentBoundSecrets = map[string]bool{
	"REPO_SYNC_TOKEN": true,
	"XCSH_API_TOKEN":  true,
	"XCSH_API_URL":    true,
}

var environmentBoundWorkflows = map[string]bool{
	"acc-tests.yml":         true,
	"discover-defaults.yml": true,
	"on-merge.yml":          true,
	"sync-openapi.yml":      true,
}

func scalarString(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func stringSlice(value any) ([]string, error) {
	switch typed := value.(type) {
	case string:
		return []string{typed}, nil
	case []any:
		result := make([]string, 0, len(typed))
		for _, item := range typed {
			text, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("non-literal runs-on item %T", item)
			}
			result = append(result, text)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("runs-on must be a literal string or list, got %T", value)
	}
}

func stringMap(value any) (map[string]string, error) {
	if value == nil {
		return nil, nil
	}
	raw, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("must be a mapping, got %T", value)
	}
	result := map[string]string{}
	for key, item := range raw {
		text, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("%s must be a string", key)
		}
		result[key] = text
	}
	return result, nil
}

func triggerMapping(node yaml.Node) (map[string]any, error) {
	var value any
	if err := node.Decode(&value); err != nil {
		return nil, err
	}
	switch typed := value.(type) {
	case string:
		return map[string]any{typed: nil}, nil
	case []any:
		out := map[string]any{}
		for _, item := range typed {
			name, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("non-string trigger")
			}
			out[name] = nil
		}
		return out, nil
	case map[string]any:
		return typed, nil
	default:
		return nil, fmt.Errorf("on must be a trigger mapping, got %T", value)
	}
}

func recursivelyCollectSecrets(value any, path []string, found map[string]bool, errors *[]string) {
	switch typed := value.(type) {
	case map[string]any:
		for key, child := range typed {
			recursivelyCollectSecrets(child, append(path, key), found, errors)
		}
	case []any:
		for index, child := range typed {
			recursivelyCollectSecrets(child, append(path, fmt.Sprint(index)), found, errors)
		}
	case string:
		matches := secretReference.FindAllStringSubmatch(typed, -1)
		for _, match := range matches {
			found[match[1]] = true
		}
		if len(matches) > 0 {
			allowed := false
			for _, component := range path {
				if component == "env" || component == "with" {
					allowed = true
				}
			}
			for _, component := range path {
				if component == "run" {
					allowed = false
				}
			}
			if !allowed {
				*errors = append(*errors, "secret reference in unsupported location "+strings.Join(path, "."))
			}
		}
	}
}

func validateWorkflowBytes(filename string, content []byte) []string {
	var workflow workflowDocument
	if err := yaml.Unmarshal(content, &workflow); err != nil {
		return []string{err.Error()}
	}
	triggers, err := triggerMapping(workflow.On)
	if err != nil {
		return []string{err.Error()}
	}
	triggerNames := make([]string, 0, len(triggers))
	for name := range triggers {
		triggerNames = append(triggerNames, name)
	}
	sort.Strings(triggerNames)
	errors := []string{}

	contracts := map[string]jobContract{}
	for _, contract := range protectedJobs {
		if contract.workflow == filename {
			contracts[contract.job] = contract
		}
	}
	if len(contracts) > 0 {
		for _, name := range []string{"pull_request_target", "workflow_run"} {
			if _, ok := triggers[name]; ok {
				errors = append(errors, "forbidden trigger "+name)
			}
		}
		if push, ok := triggers["push"]; ok {
			mapping, ok := push.(map[string]any)
			if !ok || !reflect.DeepEqual(mapping["branches"], []any{"main"}) {
				errors = append(errors, "push is not bounded exactly to main")
			}
		}
		if filename == "acc-tests.yml" {
			dispatch, _ := triggers["workflow_dispatch"].(map[string]any)
			inputs, _ := dispatch["inputs"].(map[string]any)
			if _, ok := inputs["runner"]; ok {
				errors = append(errors, "workflow_dispatch must not expose runner selection")
			}
			pullRequest, _ := triggers["pull_request"].(map[string]any)
			if !reflect.DeepEqual(pullRequest["branches"], []any{"main"}) {
				errors = append(errors, "pull_request is not bounded exactly to main")
			}
		}
		if filename == "sync-openapi.yml" {
			dispatch, _ := triggers["repository_dispatch"].(map[string]any)
			if !reflect.DeepEqual(dispatch["types"], []any{"enriched-specs-updated"}) {
				errors = append(errors, "repository_dispatch types are not exact")
			}
		}
	}
	for jobID, job := range workflow.Jobs {
		runsOn, runErr := stringSlice(job["runs-on"])
		isSelfHosted := runErr == nil && slicesContain(runsOn, "self-hosted")
		if runErr == nil {
			for _, label := range runsOn {
				if strings.Contains(label, "${{") {
					errors = append(errors, jobID+": dynamic runs-on is forbidden")
				}
			}
		}
		contract, listed := contracts[jobID]
		if isSelfHosted && !listed {
			errors = append(errors, "unlisted self-hosted job "+jobID)
		}
		if !listed {
			continue
		}
		if runErr != nil {
			errors = append(errors, jobID+": "+runErr.Error())
			continue
		}
		if !reflect.DeepEqual(runsOn, contract.runsOn) {
			errors = append(errors, fmt.Sprintf("%s: runs-on %v", jobID, runsOn))
		}
		if isSelfHosted && !reflect.DeepEqual(runsOn, []string{"self-hosted", "Linux", "X64", repositoryRunnerLabel}) {
			errors = append(errors, jobID+": invalid self-hosted tuple")
		}
		if scalarString(job["environment"]) != contract.environment {
			errors = append(errors, jobID+": environment mismatch")
		}
		permissions, permErr := stringMap(job["permissions"])
		if permErr != nil {
			errors = append(errors, jobID+": "+permErr.Error())
		} else {
			if permissions == nil {
				permissions = workflow.Permissions
			}
			if !reflect.DeepEqual(permissions, contract.permissions) {
				errors = append(errors, fmt.Sprintf("%s: permissions %v", jobID, permissions))
			}
		}
		expectedTriggers := append([]string(nil), contract.triggers...)
		sort.Strings(expectedTriggers)
		if !reflect.DeepEqual(triggerNames, expectedTriggers) {
			errors = append(errors, fmt.Sprintf("%s: triggers %v", jobID, triggerNames))
		}
		if contract.jobIf != nil && scalarString(job["if"]) != *contract.jobIf {
			errors = append(errors, jobID+": exact if mismatch")
		}
		found := map[string]bool{}
		secretErrors := []string{}
		recursivelyCollectSecrets(job, nil, found, &secretErrors)
		errors = append(errors, secretErrors...)
		actualSecrets := make([]string, 0, len(found))
		for name := range found {
			actualSecrets = append(actualSecrets, name)
		}
		sort.Strings(actualSecrets)
		expectedSecrets := append([]string(nil), contract.secrets...)
		sort.Strings(expectedSecrets)
		if strings.Join(actualSecrets, "\x00") != strings.Join(expectedSecrets, "\x00") {
			errors = append(errors, fmt.Sprintf("%s: secrets %v", jobID, actualSecrets))
		}
		steps, _ := job["steps"].([]any)
		for _, raw := range steps {
			step, _ := raw.(map[string]any)
			uses, _ := step["uses"].(string)
			if strings.HasPrefix(uses, "actions/checkout@") {
				with, _ := step["with"].(map[string]any)
				if value, ok := with["persist-credentials"].(bool); !ok || value {
					errors = append(errors, jobID+": checkout persists credentials")
				}
			}
		}
	}
	if environmentBoundWorkflows[filename] {
		for jobID, job := range workflow.Jobs {
			found := map[string]bool{}
			secretErrors := []string{}
			recursivelyCollectSecrets(job, nil, found, &secretErrors)
			for secret := range found {
				if !environmentBoundSecrets[secret] {
					continue
				}
				contract, ok := contracts[jobID]
				if !ok || !slicesContain(contract.secrets, secret) {
					errors = append(errors, fmt.Sprintf("%s: environment-bound secret %s is not in the exact job contract", jobID, secret))
				}
			}
		}
	}
	for jobID := range contracts {
		if _, ok := workflow.Jobs[jobID]; !ok {
			errors = append(errors, "missing protected job "+jobID)
		}
	}
	return errors
}

func slicesContain(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func TestProviderWorkflowContracts(t *testing.T) {
	workflowDir := filepath.Join("..", ".github", "workflows")
	entries, err := filepath.Glob(filepath.Join(workflowDir, "*.y*ml"))
	if err != nil {
		t.Fatal(err)
	}
	selfHosted := map[string]bool{}
	for _, path := range entries {
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		filename := filepath.Base(path)
		for _, issue := range validateWorkflowBytes(filename, content) {
			t.Errorf("%s: %s", filename, issue)
		}
		var workflow workflowDocument
		if err := yaml.Unmarshal(content, &workflow); err != nil {
			t.Fatal(err)
		}
		for jobID, job := range workflow.Jobs {
			runsOn, _ := stringSlice(job["runs-on"])
			if slicesContain(runsOn, "self-hosted") {
				selfHosted[filename+"/"+jobID] = true
			}
			steps, _ := job["steps"].([]any)
			for _, raw := range steps {
				step, _ := raw.(map[string]any)
				uses, _ := step["uses"].(string)
				if strings.HasPrefix(uses, "actions/checkout@") {
					with, _ := step["with"].(map[string]any)
					if value, ok := with["persist-credentials"].(bool); !ok || value {
						t.Errorf("%s/%s: checkout must set literal persist-credentials: false", filename, jobID)
					}
				}
			}
		}
	}
	expected := map[string]bool{"acc-tests.yml/real-api-tests": true, "acc-tests.yml/cleanup": true, "discover-defaults.yml/discover": true}
	if !reflect.DeepEqual(selfHosted, expected) {
		t.Fatalf("self-hosted inventory mismatch: %v", selfHosted)
	}
}

func TestProviderWorkflowMutationsFail(t *testing.T) {
	base, err := os.ReadFile(filepath.Join("..", ".github", "workflows", "acc-tests.yml"))
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]func(string) string{
		"obsolete runner input": func(s string) string {
			return strings.Replace(s, "      timeout:\n", "      runner:\n        default: ubuntu-latest\n        type: string\n      timeout:\n", 1)
		},
		"removed PR exclusion": func(s string) string {
			return strings.Replace(s, "github.event_name != 'pull_request' &&", "true &&", 1)
		},
		"OR guard": func(s string) string {
			return strings.Replace(s, "github.event_name != 'pull_request' &&", "github.event_name != 'pull_request' ||", 1)
		},
		"repository dispatch": func(s string) string {
			return strings.Replace(s, "  pull_request:\n", "  repository_dispatch:\n\n  pull_request:\n", 1)
		},
		"unbounded push": func(s string) string {
			return strings.Replace(s, "  pull_request:\n", "  push:\n\n  pull_request:\n", 1)
		},
		"direct run secret": func(s string) string {
			return strings.Replace(s, "          echo \"Checking API credentials...\"", "          echo \"${{ secrets.XCSH_API_TOKEN }}\"", 1)
		},
		"persist checkout": func(s string) string {
			return strings.Replace(s, "persist-credentials: false", "persist-credentials: true", 1)
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			mutated := mutate(string(base))
			if mutated == string(base) {
				t.Fatal("mutation did not apply")
			}
			if issues := validateWorkflowBytes("acc-tests.yml", []byte(mutated)); len(issues) == 0 {
				t.Fatal("unsafe mutation passed validation")
			}
		})
	}
}
