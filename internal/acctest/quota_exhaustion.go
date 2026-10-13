// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"os"
	"path/filepath"
	"strings"
)

func confirmedObjectExhaustion(err error) bool {
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 429 {
		return false
	}
	message := apiErr.Message
	if raw, ok := apiErr.Details["raw_response"].(string); ok {
		var envelope struct {
			Message string `json:"message"`
		}
		if json.Unmarshal([]byte(raw), &envelope) != nil {
			return false
		}
		message = envelope.Message
	}
	return strings.Contains(strings.ToLower(message), "object kind") && strings.Contains(strings.ToLower(message), "exhausted limits")
}

func provisionQuotaBoundCapabilities(ctx context.Context, c *client.Client) error {
	operations := []struct {
		typ    string
		create func(string) error
	}{
		{"xcsh_api_crawler", func(name string) error {
			_, err := c.CreateAPICrawler(ctx, &client.APICrawler{Metadata: client.Metadata{Name: name, Namespace: "system"}, Spec: map[string]interface{}{"domains": []map[string]interface{}{{"domain": "example.com"}}}})
			return err
		}},
		{"xcsh_crl", func(name string) error {
			_, err := c.CreateCRL(ctx, &client.CRL{Metadata: client.Metadata{Name: name, Namespace: "system"}, Spec: map[string]interface{}{"server_address": "crl.example.com", "server_port": 80, "refresh_interval": 24, "timeout": 30, "http_access": map[string]interface{}{"path": "/fixture.crl"}}})
			return err
		}},
	}
	for _, operation := range operations {
		err := operation.create(CompactRandomName(20))
		if err == nil {
			continue
		}
		if !confirmedObjectExhaustion(err) {
			return errors.New("quota capability probe failed without authoritative exhaustion")
		}
		data, _ := json.Marshal(map[string]interface{}{"version": 1, "resource_type": operation.typ, "status": 429, "outcome": "object_limits_exhausted"})
		if err = os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), operation.typ+"-quota.json"), data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func exhaustedCapabilityReceipt(typ string) bool {
	data, err := os.ReadFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), typ+"-quota.json"))
	if err != nil {
		return false
	}
	var receipt struct {
		Version int    `json:"version"`
		Type    string `json:"resource_type"`
		Status  int    `json:"status"`
		Outcome string `json:"outcome"`
	}
	return json.Unmarshal(data, &receipt) == nil && receipt.Version == 1 && receipt.Type == typ && receipt.Status == 429 && receipt.Outcome == "object_limits_exhausted"
}
