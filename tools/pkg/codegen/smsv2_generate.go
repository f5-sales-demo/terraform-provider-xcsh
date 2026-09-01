// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

type smsv2GenerationContract struct {
	Version    string `json:"version"`
	ContractID string `json:"contract_id"`
	Providers  struct {
		AWS struct {
			Telemetry struct {
				SchemaID string `json:"schema_id"`
				Complete bool   `json:"complete"`
			} `json:"telemetry_intake"`
		} `json:"aws"`
	} `json:"providers"`
}

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
	var contract smsv2GenerationContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract: %w", err)
	}
	if contract.Providers.AWS.Telemetry.SchemaID != "f5xc-smsv2-aws-tgw-telemetry/v1" || !contract.Providers.AWS.Telemetry.Complete {
		return nil, fmt.Errorf("SMSv2 telemetry contract is incomplete")
	}
	manifestJSON, err := os.ReadFile(filepath.Join(specDir, "smsv2-contract-manifest.json"))
	if err != nil {
		return nil, fmt.Errorf("read SMSv2 contract manifest: %w", err)
	}
	var manifest smsv2GenerationManifest
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		return nil, fmt.Errorf("decode SMSv2 contract manifest: %w", err)
	}
	if manifest.ContractID != contract.ContractID || manifest.ContractVersion != contract.Version || manifest.Release.Tag != "v"+contract.Version || len(manifest.Release.Commit) != 40 {
		return nil, fmt.Errorf("SMSv2 contract manifest identity mismatch")
	}
	source := fmt.Sprintf(`// Code generated from api-specs-enriched %s smsv2-contract.json. DO NOT EDIT.

package provider

const (
	smsv2ContractID = %q
	smsv2ContractVersion = %q
	smsv2APIReleaseTag = %q
	smsv2SourceCommit = %q
	smsv2TelemetrySchemaID = %q
)

var smsv2ContractCapabilities = map[string]string{"aws_ce_create": "available", "runtime_status": "available", "tgw_connect": "available"}
var smsv2ContractF5XCAuthorities = []string{"smsv2_configuration", "runtime_health", "bgp_peers", "bgp_routes", "simplified_routes"}
var smsv2ContractAWSAuthorities = []string{"eni", "transit_gateway", "transit_gateway_connect", "gre_endpoints", "bgp_inside_cidrs"}
`, manifest.Release.Tag, contract.ContractID, contract.Version, manifest.Release.Tag, manifest.Release.Commit, contract.Providers.AWS.Telemetry.SchemaID)
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return nil, fmt.Errorf("format SMSv2 generated constants: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outputDir, "smsv2_contract_generated.go"), formatted, 0o644); err != nil {
		return nil, fmt.Errorf("write SMSv2 generated constants: %w", err)
	}
	return templates, nil
}
