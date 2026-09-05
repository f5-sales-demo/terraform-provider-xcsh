// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type smsv2GenerationManifest struct {
	ContractID      string `json:"contract_id"`
	ContractVersion string `json:"contract_version"`
	Release         struct {
		Tag    string `json:"tag"`
		Commit string `json:"commit"`
	} `json:"release"`
}

// GenerateSMSv2ContractConstants binds the provider binary to one immutable
// API release. Runtime data sources cannot silently drift to mutable assets.
func GenerateSMSv2ContractConstants(specDir, outputDir string) ([]SMSv2DataSourceTemplate, error) {
	contractJSON, err := os.ReadFile(filepath.Join(specDir, "smsv2-contract.json"))
	if err != nil {
		return nil, fmt.Errorf("read SMSv2 contract: %w", err)
	}
	templates, err := SMSv2DataSourceTemplates(contractJSON)
	if err != nil {
		return nil, err
	}
	var contract smsv2ReleaseContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract: %w", err)
	}
	manifestJSON, err := os.ReadFile(filepath.Join(specDir, "smsv2-contract-manifest.json"))
	if err != nil {
		return nil, fmt.Errorf("read SMSv2 contract manifest: %w", err)
	}
	var manifest smsv2GenerationManifest
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract manifest: %w", err)
	}
	contractMajor, _, contractFound := strings.Cut(contract.Version, ".")
	releaseMajor, _, releaseFound := strings.Cut(strings.TrimPrefix(manifest.Release.Tag, "v"), ".")
	if manifest.ContractID != contract.ContractID || manifest.ContractVersion != contract.Version ||
		!strings.HasPrefix(manifest.Release.Tag, "v") || !contractFound || !releaseFound ||
		contractMajor != releaseMajor || len(manifest.Release.Commit) != 40 {
		return nil, fmt.Errorf("SMSv2 contract manifest identity mismatch")
	}
	capabilities := fmt.Sprintf("map[string]string{%q: %q, %q: %q, %q: %q, %q: %q}",
		"aws_ce_create", contract.Providers.AWS.Capabilities["aws_ce_create"],
		"runtime_status", contract.Providers.AWS.Capabilities["runtime_status"],
		"site_upgrade", contract.Providers.AWS.Capabilities["site_upgrade"],
		"tgw_connect", contract.Providers.AWS.Capabilities["tgw_connect"])
	f5xcAuthorities := goStringSlice(contract.Providers.AWS.Authorities["f5xc"])
	awsAuthorities := goStringSlice(contract.Providers.AWS.Authorities["aws"])
	source := fmt.Sprintf(`// Code generated from api-specs-enriched %s smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID = %q
	smsv2ContractVersion = %q
	smsv2APIReleaseTag = %q
	smsv2SourceCommit = %q
	smsv2TelemetrySchemaID = %q
)

var smsv2ContractCapabilities = %s
var smsv2ContractF5XCAuthorities = %s
var smsv2ContractAWSAuthorities = %s
`, manifest.Release.Tag, contract.ContractID, contract.Version, manifest.Release.Tag, manifest.Release.Commit, contract.Providers.AWS.Telemetry.SchemaID, capabilities, f5xcAuthorities, awsAuthorities)
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return nil, fmt.Errorf("format SMSv2 generated constants: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outputDir, "smsv2_contract_generated.go"), formatted, 0o644); err != nil {
		return nil, fmt.Errorf("write SMSv2 generated constants: %w", err)
	}
	return templates, nil
}

func goStringSlice(values []string) string {
	quoted := make([]string, len(values))
	for index, value := range values {
		quoted[index] = strconv.Quote(value)
	}
	return "[]string{" + strings.Join(quoted, ", ") + "}"
}
