// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func logAcceptanceAPIDiagnostic(response *http.Response) {
	if response == nil || response.StatusCode < 400 || response.StatusCode >= 600 {
		return
	}
	data, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		return
	}
	// Preserve the original response for the API client; only fixed facts are logged.
	response.Body = io.NopCloser(bytes.NewReader(data))
	text := strings.ToLower(string(data))
	words := regexp.MustCompile(`[a-z][a-z0-9_]*`).FindAllString(text, -1)
	vocabulary := strings.Fields("message code details quota quotas limit limits rate maximum minimum available allowed exceeded exceeding exhausted exhausted_resource exhaustion configured configuration configure namespace namespaces system shared default kind resource resources global tenant number count current total not no none missing found unknown unsupported invalid denied permission forbidden must should requires required need needs labels label characters character length longer exceeds char dns1035 valid format wrong invalid tenant namespace slash copied reference expanding expansion resolver store read could found resolve cidr interface classification virtualsites virtual_site ref_id fully qualified tenant_id cannot unable failed failure create creation capacity max reached greater than too many exists exist existing object objects name value ref reference referred referring selector site virtual customer regional edge ce re segment network dns config field constraint constraints rule match selected expansion resolve internal server error operation format response status type specification pool port in from to for is are the a of be on with as undefined bad out range larger less zero negative per exhaustedresource enabled feature entitlement entitlements quota_limit quota_usage quota_used limit_exceeded")
	allowed := map[string]bool{}
	for _, word := range vocabulary {
		allowed[word] = true
	}
	safe := []string{}
	for _, word := range words {
		if allowed[word] {
			safe = append(safe, word)
		} else {
			safe = append(safe, "_")
		}
	}
	if private := os.Getenv("XCSH_ACC_PRIVATE"); private != "" {
		record := map[string]interface{}{"status": response.StatusCode, "words": safe}
		limits := []int{}
		for _, match := range regexp.MustCompile(`(?i)(?:no (?:more|longer) than|at most|maximum length|length must be|must be no more than) ([0-9]{1,3})`).FindAllStringSubmatch(text, -1) {
			n, _ := strconv.Atoi(match[1])
			if n > 0 && n <= 128 {
				limits = append(limits, n)
			}
		}
		record["length_limits"] = limits
		if response.Request != nil {
			for typ, route := range ownedCollectionRoutes {
				parts := strings.Split(response.Request.URL.Path, "/")
				ns := ""
				for i, part := range parts {
					if part == "namespaces" && i+1 < len(parts) {
						ns = parts[i+1]
						break
					}
				}
				if strings.ReplaceAll(route, "%s", ns) == response.Request.URL.Path {
					record["resource_type"] = typ
					break
				}
			}
		}
		encoded, _ := json.Marshal(record)
		f, err := os.OpenFile(private+"/fixed-api-errors.jsonl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err == nil {
			_, _ = f.Write(append(encoded, '\n'))
			_ = f.Close()
		}
	}
	signatures := map[string][]string{
		"resource_dependency_missing": {"not found", "does not exist", "failed to resolve", "no such"},
		"quota_unconfigured":          {"quota not configured"},
		"resource_limit_exhausted":    {"resource exhausted", "resource_exhausted", "limit exceeded", "exceeds quota", "maximum allowed", "quota"},
		"namespace_required_default":  {"default namespace", "namespace default"},
		"namespace_required_shared":   {"shared namespace", "namespace shared"},
		"site_selector_unavailable":   {"no matching site", "no sites", "site selector", "site-selector", "virtual-site"},
		"quota_exceeded":              {"quota exceeded", "quota limit", "maximum number"},
		"document_fetch_failed":       {"unable to download", "failed to download", "failed to fetch", "cannot fetch", "fetch failed", "file not found", "not found"},
		"document_parse_failed":       {"failed to parse", "unable to parse", "parse error", "invalid spec", "invalid schema", "invalid format"},
		"swagger_reference":           {"swagger", "openapi"},
		"namespace_constraint":        {"namespace"},
		"missing_required_input":      {"required", "must specify", "cannot be empty"},
		"schema_validation":           {"validation", "invalid request", "bad request"},
		"authentication_denied":       {"unauthorized", "authentication", "invalid token"},
	}
	facts := []string{}
	for name, patterns := range signatures {
		for _, pattern := range patterns {
			if strings.Contains(text, pattern) {
				facts = append(facts, name)
				log.Printf("[API_DIAGNOSTIC:%s]", name)
				logDiagnosticFact(name)
				break
			}
		}
	}
	if schema, err := LoadFixtureSchema(os.Getenv("XCSH_ACC_SCHEMA")); err == nil {
		if response.Request != nil {
			for typ, route := range ownedCollectionRoutes {
				parts := strings.Split(response.Request.URL.Path, "/")
				ns := ""
				for i, p := range parts {
					if p == "namespaces" && i+1 < len(parts) {
						ns = parts[i+1]
						break
					}
				}
				if strings.ReplaceAll(route, "%s", ns) == response.Request.URL.Path {
					schema = &FixtureSchema{Resources: map[string]fixtureBlockSchema{typ: schema.Resources[typ]}}
					break
				}
			}
		}
		for _, field := range diagnosticSchemaFields(schema, text) {
			acceptanceDiagnosticTest.RLock()
			test := acceptanceDiagnosticTest.test
			acceptanceDiagnosticTest.RUnlock()
			if test != nil {
				test.Logf("[SCHEMA_FIELD:%s]", field)
			}
		}
	}
	sort.Strings(facts)
	var body map[string]interface{}
	if json.Unmarshal(data, &body) == nil {
		message, _ := body["message"].(string)
		for _, fact := range facts {
			message += " [API_DIAGNOSTIC:" + fact + "]"
		}
		body["message"] = message
		if serialized, err := json.Marshal(body); err == nil {
			response.Body = io.NopCloser(bytes.NewReader(serialized))
			response.ContentLength = int64(len(serialized))
		}
	}
}

func diagnosticSchemaFields(schema *FixtureSchema, message string) []string {
	tokens := map[string]bool{}
	for _, token := range regexp.MustCompile(`[a-z][a-z0-9_]*`).FindAllString(strings.ToLower(message), -1) {
		tokens[token] = true
	}
	result := []string{}
	seen := map[string]bool{}
	var walk func(fixtureBlockSchema, string)
	walk = func(block fixtureBlockSchema, prefix string) {
		for name := range block.Attributes {
			if tokens[name] && !seen[prefix+name] {
				seen[prefix+name] = true
				result = append(result, prefix+name)
			}
		}
		for name, nested := range block.Blocks {
			if tokens[name] && !seen[prefix+name] {
				seen[prefix+name] = true
				result = append(result, prefix+name)
			}
			walk(nested.Block, prefix+name+".")
		}
	}
	for _, block := range schema.Resources {
		walk(block, "")
	}
	sort.Strings(result)
	return result
}
