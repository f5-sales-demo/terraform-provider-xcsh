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
	"time"
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
		{"xcsh_container_registry", func(name string) error {
			_, err := c.CreateContainerRegistry(ctx, &client.ContainerRegistry{Metadata: client.Metadata{Name: name, Namespace: "shared"}, Spec: map[string]interface{}{"registry": "registry.example.com", "user_name": "testuser", "email": "test@example.com", "password": map[string]interface{}{"clear_secret_info": map[string]interface{}{"url": "string:///dGVzdC1wYXNzd29yZA=="}}}})
			return err
		}},
	}
	for _, operation := range operations {
		probeName := CompactRandomName(20)
		err := operation.create(probeName)
		if err == nil {
			journal, loadErr := acceptanceJournal()
			if loadErr != nil {
				return errors.New("quota probe ownership unavailable")
			}
			journal.mu.Lock()
			var owned OwnedResource
			for _, r := range journal.Resources {
				if r.Type == operation.typ && r.Name == probeName && r.Status == "created" {
					owned = r
				}
			}
			journal.mu.Unlock()
			if owned.Name == "" {
				return errors.New("quota probe owned identity missing")
			}
			backend := apiOwnershipBackend{c: c, run: journal.RunID}
			if err := releaseQuotaProbe(ctx, backend, owned, 12, time.Second); err != nil {
				return err
			}
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

func releaseQuotaProbe(ctx context.Context, backend OwnershipBackend, owned OwnedResource, attempts int, delay time.Duration) error {
	if err := backend.Delete(ctx, owned); err != nil && !errors.Is(err, errOwnedAbsent) {
		return errors.New("quota probe deletion failed")
	}
	for attempt := 0; attempt < attempts; attempt++ {
		_, err := backend.Read(ctx, owned)
		if errors.Is(err, errOwnedAbsent) {
			return nil
		}
		if err != nil {
			return errors.New("quota probe cleanup read failed")
		}
		if waitContext(ctx, delay) != nil {
			return errors.New("quota probe cleanup interrupted")
		}
	}
	return errors.New("quota probe cleanup incomplete")
}
