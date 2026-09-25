// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package networkallowlist extracts the release-pinned F5XC network allowlist
// extension into the compact artifact compiled into the Terraform provider.
package networkallowlist

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"net/netip"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const extensionName = "x-f5xc-network-allowlist"

var expectedServices = []string{
	"bot_defense",
	"cdn",
	"data_intelligence",
	"dnslb_health_checks",
	"global_controller_sso_egress",
	"global_log_receiver",
	"regional_edges",
	"secondary_dns_zone_transfer",
}

type Artifact struct {
	APIReleaseTag string                  `json:"api_release_tag"`
	SourceURL     string                  `json:"source_url"`
	GeneratedAt   string                  `json:"generated_at"`
	SourceSHA256  string                  `json:"source_sha256"`
	Services      map[string]ServiceGroup `json:"services"`
	CustomerEdge  CustomerEdge            `json:"customer_edge"`
}

type ServiceGroup struct {
	SourceEntries []string                `json:"source_entries,omitempty"`
	Domains       []string                `json:"domains,omitempty"`
	Regions       map[string]ServiceGroup `json:"regions,omitempty"`
}

type CustomerEdge struct {
	Defaults     CustomerEdgeDefaults `json:"defaults"`
	SecureMeshV2 CustomerEdgeEgress   `json:"secure_mesh_v2"`
}

type CustomerEdgeDefaults struct {
	DNSServers []string `json:"dns_servers"`
	NTPServers []string `json:"ntp_servers"`
}

type CustomerEdgeEgress struct {
	RegistrationAddresses []string `json:"registration_addresses"`
	Domains               []string `json:"domains"`
}

type extension struct {
	SourceURL string          `json:"source_url"`
	SHA256    string          `json:"sha256"`
	Manifest  json.RawMessage `json:"manifest"`
}

type pinnedRelease struct {
	ReleaseTag string            `json:"release_tag"`
	Version    string            `json:"version"`
	Commit     string            `json:"target_commit"`
	Assets     map[string]string `json:"assets"`
}

func Extract(openAPIRaw, releasePinRaw []byte, releaseTag string) (Artifact, error) {
	var pin pinnedRelease
	if err := strictDecode(releasePinRaw, &pin); err != nil {
		return Artifact{}, fmt.Errorf("parse release pin: %w", err)
	}
	if pin.ReleaseTag != releaseTag {
		return Artifact{}, fmt.Errorf("release pin tag %q does not match %q", pin.ReleaseTag, releaseTag)
	}
	expectedOpenAPI := strings.TrimPrefix(pin.Assets["openapi.json"], "sha256:")
	actualOpenAPISum := sha256.Sum256(openAPIRaw)
	actualOpenAPI := hex.EncodeToString(actualOpenAPISum[:])
	if len(expectedOpenAPI) != 64 || actualOpenAPI != expectedOpenAPI {
		return Artifact{}, fmt.Errorf("openapi.json digest mismatch: got %s, want %s", actualOpenAPI, expectedOpenAPI)
	}

	var document struct {
		Info map[string]json.RawMessage `json:"info"`
	}
	if err := json.Unmarshal(openAPIRaw, &document); err != nil {
		return Artifact{}, fmt.Errorf("parse OpenAPI document: %w", err)
	}
	rawExtension, ok := document.Info[extensionName]
	if !ok {
		return Artifact{}, fmt.Errorf("missing info.%s extension", extensionName)
	}
	var ext extension
	if err := strictDecode(rawExtension, &ext); err != nil {
		return Artifact{}, fmt.Errorf("parse info.%s: %w", extensionName, err)
	}
	parsedURL, err := url.Parse(ext.SourceURL)
	if err != nil || parsedURL.Scheme != "https" || parsedURL.Host == "" {
		return Artifact{}, fmt.Errorf("source_url must be an absolute HTTPS URL")
	}
	var rawManifest map[string]json.RawMessage
	if err := strictDecode(ext.Manifest, &rawManifest); err != nil {
		return Artifact{}, fmt.Errorf("parse network allowlist manifest: %w", err)
	}
	if err := exactKeys(rawManifest, "manifest", "generated_at", "manifest_type", "manifest_version", "reference", "schema_version", "services", "customer_edge"); err != nil {
		return Artifact{}, err
	}
	var generatedAt string
	if err := json.Unmarshal(rawManifest["generated_at"], &generatedAt); err != nil {
		return Artifact{}, fmt.Errorf("parse manifest generated_at: %w", err)
	}
	if _, err := time.Parse(time.RFC3339, generatedAt); err != nil {
		return Artifact{}, fmt.Errorf("invalid manifest generated_at %q: %w", generatedAt, err)
	}
	var manifestAny any
	if err := json.Unmarshal(ext.Manifest, &manifestAny); err != nil {
		return Artifact{}, fmt.Errorf("parse manifest for digest: %w", err)
	}
	canonical, err := json.Marshal(manifestAny)
	if err != nil {
		return Artifact{}, fmt.Errorf("canonicalize manifest: %w", err)
	}
	sourceSum := sha256.Sum256(canonical)
	actualSourceDigest := hex.EncodeToString(sourceSum[:])
	if len(ext.SHA256) != 64 || actualSourceDigest != ext.SHA256 {
		return Artifact{}, fmt.Errorf("network allowlist digest mismatch: got %s, want %s", actualSourceDigest, ext.SHA256)
	}

	var rawServices map[string]json.RawMessage
	if err := strictDecode(rawManifest["services"], &rawServices); err != nil {
		return Artifact{}, fmt.Errorf("parse manifest services: %w", err)
	}
	if err := exactKeys(rawServices, "services", expectedServices...); err != nil {
		return Artifact{}, err
	}
	services := make(map[string]ServiceGroup, len(rawServices))
	for _, name := range []string{"cdn", "secondary_dns_zone_transfer", "global_log_receiver", "dnslb_health_checks", "global_controller_sso_egress"} {
		group, err := decodeIPGroup(name, rawServices[name])
		if err != nil {
			return Artifact{}, err
		}
		services[name] = group
	}
	bot, err := decodeDomainGroup("bot_defense", rawServices["bot_defense"])
	if err != nil {
		return Artifact{}, err
	}
	services["bot_defense"] = bot
	regional, err := decodeRegionGroup("regional_edges", rawServices["regional_edges"], []string{"americas", "asia", "europe"})
	if err != nil {
		return Artifact{}, err
	}
	services["regional_edges"] = regional
	intelligence, err := decodeRegionGroup("data_intelligence", rawServices["data_intelligence"], []string{"eu", "us"})
	if err != nil {
		return Artifact{}, err
	}
	services["data_intelligence"] = intelligence

	customerEdge, err := decodeCustomerEdge(rawManifest["customer_edge"])
	if err != nil {
		return Artifact{}, err
	}
	return Artifact{
		APIReleaseTag: releaseTag,
		SourceURL:     ext.SourceURL,
		GeneratedAt:   generatedAt,
		SourceSHA256:  ext.SHA256,
		Services:      services,
		CustomerEdge:  customerEdge,
	}, nil
}

func Load(openAPIPath, releasePinPath, versionPath string) (Artifact, error) {
	openAPIRaw, err := os.ReadFile(openAPIPath)
	if err != nil {
		return Artifact{}, fmt.Errorf("read pinned OpenAPI: %w", err)
	}
	releasePinRaw, err := os.ReadFile(releasePinPath)
	if err != nil {
		return Artifact{}, fmt.Errorf("read release pin: %w", err)
	}
	versionRaw, err := os.ReadFile(versionPath)
	if err != nil {
		return Artifact{}, fmt.Errorf("read spec version: %w", err)
	}
	return Extract(openAPIRaw, releasePinRaw, strings.TrimSpace(string(versionRaw)))
}

func WriteGo(path string, artifact Artifact) error {
	encoded, err := json.Marshal(artifact)
	if err != nil {
		return fmt.Errorf("encode bundled network allowlist: %w", err)
	}
	source := fmt.Sprintf("// Code generated by generate-all-schemas.go. DO NOT EDIT.\n// Source: info.%s in the pinned OpenAPI release.\n\npackage provider\n\nconst bundledNetworkAllowlistJSON = %q\n", extensionName, string(encoded))
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return fmt.Errorf("format bundled network allowlist: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	temp, err := os.CreateTemp(filepath.Dir(path), ".network-allowlist-*.go")
	if err != nil {
		return fmt.Errorf("create temporary output: %w", err)
	}
	tempName := temp.Name()
	defer os.Remove(tempName)
	if _, err := temp.Write(formatted); err != nil {
		_ = temp.Close()
		return fmt.Errorf("write temporary output: %w", err)
	}
	if err := temp.Chmod(0o644); err != nil {
		_ = temp.Close()
		return fmt.Errorf("chmod temporary output: %w", err)
	}
	if err := temp.Close(); err != nil {
		return fmt.Errorf("close temporary output: %w", err)
	}
	if err := os.Rename(tempName, path); err != nil {
		return fmt.Errorf("replace generated output: %w", err)
	}
	return nil
}

func strictDecode(data []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return fmt.Errorf("multiple JSON values")
		}
		return fmt.Errorf("trailing JSON data: %w", err)
	}
	return nil
}

func exactKeys(values map[string]json.RawMessage, context string, expected ...string) error {
	expectedSet := make(map[string]bool, len(expected))
	for _, name := range expected {
		expectedSet[name] = true
		if _, ok := values[name]; !ok {
			return fmt.Errorf("%s is missing required key %q", context, name)
		}
	}
	for name := range values {
		if !expectedSet[name] {
			return fmt.Errorf("%s contains unsupported key %q", context, name)
		}
	}
	return nil
}

func decodeIPGroup(name string, raw json.RawMessage) (ServiceGroup, error) {
	var fields map[string]json.RawMessage
	if err := strictDecode(raw, &fields); err != nil {
		return ServiceGroup{}, fmt.Errorf("parse service %s: %w", name, err)
	}
	if len(fields) != 1 {
		return ServiceGroup{}, fmt.Errorf("service %s must contain exactly one address list", name)
	}
	for _, key := range []string{"ipv4_ips", "ipv4_cidrs"} {
		if entriesRaw, ok := fields[key]; ok {
			var entries []string
			if err := strictDecode(entriesRaw, &entries); err != nil {
				return ServiceGroup{}, fmt.Errorf("parse service %s.%s: %w", name, key, err)
			}
			if err := validateAddresses(name, key, entries); err != nil {
				return ServiceGroup{}, err
			}
			return ServiceGroup{SourceEntries: entries}, nil
		}
	}
	return ServiceGroup{}, fmt.Errorf("service %s has no supported IPv4 address field", name)
}

func decodeDomainGroup(name string, raw json.RawMessage) (ServiceGroup, error) {
	var fields map[string]json.RawMessage
	if err := strictDecode(raw, &fields); err != nil {
		return ServiceGroup{}, fmt.Errorf("parse service %s: %w", name, err)
	}
	if err := exactKeys(fields, "service "+name, "domains"); err != nil {
		return ServiceGroup{}, err
	}
	var domains []string
	if err := strictDecode(fields["domains"], &domains); err != nil {
		return ServiceGroup{}, fmt.Errorf("parse service %s domains: %w", name, err)
	}
	if err := validateDomains(name, domains); err != nil {
		return ServiceGroup{}, err
	}
	return ServiceGroup{Domains: domains}, nil
}

func decodeRegionGroup(name string, raw json.RawMessage, expectedRegions []string) (ServiceGroup, error) {
	var fields map[string]json.RawMessage
	if err := strictDecode(raw, &fields); err != nil {
		return ServiceGroup{}, fmt.Errorf("parse service %s: %w", name, err)
	}
	if err := exactKeys(fields, "service "+name, "regions"); err != nil {
		return ServiceGroup{}, err
	}
	var regions map[string]json.RawMessage
	if err := strictDecode(fields["regions"], &regions); err != nil {
		return ServiceGroup{}, fmt.Errorf("parse service %s regions: %w", name, err)
	}
	if err := exactKeys(regions, "service "+name+" regions", expectedRegions...); err != nil {
		return ServiceGroup{}, err
	}
	result := ServiceGroup{Regions: make(map[string]ServiceGroup, len(regions))}
	for _, region := range expectedRegions {
		group, err := decodeIPGroup(name+"."+region, regions[region])
		if err != nil {
			return ServiceGroup{}, err
		}
		result.Regions[region] = group
	}
	return result, nil
}

func decodeCustomerEdge(raw json.RawMessage) (CustomerEdge, error) {
	var root map[string]json.RawMessage
	if err := strictDecode(raw, &root); err != nil {
		return CustomerEdge{}, fmt.Errorf("parse customer_edge: %w", err)
	}
	if err := exactKeys(root, "customer_edge", "defaults", "site_types"); err != nil {
		return CustomerEdge{}, err
	}
	var defaults struct {
		DNSServers []string `json:"dns_servers"`
		NTPServers []string `json:"ntp_servers"`
	}
	if err := strictDecode(root["defaults"], &defaults); err != nil {
		return CustomerEdge{}, fmt.Errorf("parse customer_edge.defaults: %w", err)
	}
	if err := validateAddresses("customer_edge.defaults", "dns_servers", defaults.DNSServers); err != nil {
		return CustomerEdge{}, err
	}
	if err := validateAddresses("customer_edge.defaults", "ntp_servers", defaults.NTPServers); err != nil {
		return CustomerEdge{}, err
	}
	var siteTypes map[string]json.RawMessage
	if err := strictDecode(root["site_types"], &siteTypes); err != nil {
		return CustomerEdge{}, fmt.Errorf("parse customer_edge.site_types: %w", err)
	}
	if err := exactKeys(siteTypes, "customer_edge.site_types", "legacy", "secure_mesh_v2"); err != nil {
		return CustomerEdge{}, err
	}
	var secure struct {
		EgressDomainRules struct {
			Domains []string `json:"domains"`
		} `json:"egress_domain_rules"`
		EgressIPRules struct {
			RegistrationUpdates struct {
				IPv4IPs []string `json:"ipv4_ips"`
			} `json:"registration_updates"`
		} `json:"egress_ip_rules"`
	}
	if err := strictDecode(siteTypes["secure_mesh_v2"], &secure); err != nil {
		return CustomerEdge{}, fmt.Errorf("parse customer_edge.site_types.secure_mesh_v2: %w", err)
	}
	if err := validateAddresses("customer_edge.secure_mesh_v2", "registration_updates", secure.EgressIPRules.RegistrationUpdates.IPv4IPs); err != nil {
		return CustomerEdge{}, err
	}
	if err := validateDomains("customer_edge.secure_mesh_v2", secure.EgressDomainRules.Domains); err != nil {
		return CustomerEdge{}, err
	}
	return CustomerEdge{
		Defaults: CustomerEdgeDefaults{DNSServers: defaults.DNSServers, NTPServers: defaults.NTPServers},
		SecureMeshV2: CustomerEdgeEgress{
			RegistrationAddresses: secure.EgressIPRules.RegistrationUpdates.IPv4IPs,
			Domains:               secure.EgressDomainRules.Domains,
		},
	}, nil
}

func validateAddresses(group, field string, values []string) error {
	if len(values) == 0 {
		return fmt.Errorf("%s.%s must not be empty", group, field)
	}
	seen := map[string]bool{}
	for _, value := range values {
		valid := false
		if strings.Contains(value, "/") {
			if prefix, err := netip.ParsePrefix(value); err == nil && prefix.Addr().Is4() {
				valid = true
			}
		} else if address, err := netip.ParseAddr(value); err == nil && address.Is4() {
			valid = true
		}
		if !valid {
			return fmt.Errorf("%s.%s contains invalid IPv4 address or CIDR %q", group, field, value)
		}
		if seen[value] {
			return fmt.Errorf("%s.%s contains duplicate %q", group, field, value)
		}
		seen[value] = true
	}
	return nil
}

func validateDomains(group string, values []string) error {
	if len(values) == 0 {
		return fmt.Errorf("%s.domains must not be empty", group)
	}
	seen := map[string]bool{}
	for _, value := range values {
		check := strings.TrimPrefix(value, ".")
		if value == "" || strings.ContainsAny(value, " /:@") || !strings.Contains(check, ".") {
			return fmt.Errorf("%s.domains contains invalid domain %q", group, value)
		}
		if seen[value] {
			return fmt.Errorf("%s.domains contains duplicate %q", group, value)
		}
		seen[value] = true
	}
	return nil
}

func SortedServiceNames(services map[string]ServiceGroup) []string {
	names := make([]string, 0, len(services))
	for name := range services {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
