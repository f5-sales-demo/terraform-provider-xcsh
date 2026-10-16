package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// SecretPolicy uses the published secret_management API, not config CRUD.
type SecretPolicy struct {
	Metadata       SecretPolicyMetadata `json:"metadata"`
	Spec           SecretPolicySpec     `json:"spec"`
	SystemMetadata struct {
		UID string `json:"uid"`
	} `json:"system_metadata,omitempty"`
	ResourceVersion json.RawMessage `json:"resource_version,omitempty"`
}
type SecretPolicyMetadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
type SecretPolicySpec struct {
	Algorithm           string                `json:"algo,omitempty"`
	AllowF5XC           bool                  `json:"allow_f5xc"`
	DecryptCacheTimeout string                `json:"decrypt_cache_timeout"`
	RuleList            *SecretPolicyRuleList `json:"rule_list,omitempty"`
	LegacyRuleList      *struct {
		Rules []SecretPolicyMetadata `json:"rules"`
	} `json:"legacy_rule_list,omitempty"`
	MarkedForDelete bool   `json:"marked_for_delete,omitempty"`
	DeletionTime    string `json:"deletion_time,omitempty"`
}
type SecretPolicyRuleList struct {
	Rules []SecretPolicyRule `json:"rules"`
}
type SecretPolicyRule struct {
	Metadata SecretPolicyMetadata `json:"metadata"`
	Spec     SecretPolicyRuleSpec `json:"spec"`
}
type SecretPolicyRuleSpec struct {
	ClientName        string                `json:"client_name,omitempty"`
	ClientNameMatcher json.RawMessage       `json:"client_name_matcher,omitempty"`
	Action            string                `json:"action"`
	ClientSelector    *SecretPolicySelector `json:"client_selector,omitempty"`
}
type SecretPolicySelector struct {
	Expressions []string `json:"expressions"`
}

func secretPolicyItem(ns, name string) string {
	return "/api/secret_management/namespaces/" + url.PathEscape(ns) + "/secret_policys/" + url.PathEscape(name)
}

// GetSecretPolicy normalizes replace-form and legacy child references into inline rules.
// Tombstones and identity always come from the full response, even with response_format=2.
func (c *Client) GetSecretPolicy(ctx context.Context, ns, name string) (*SecretPolicy, error) {
	var full SecretPolicy
	if err := c.Get(ctx, secretPolicyItem(ns, name), &full); err != nil {
		return nil, err
	}
	if full.Metadata.Name != name || full.Metadata.Namespace != ns {
		return nil, errors.New("secret policy readback identity mismatch")
	}
	if full.Spec.Algorithm != "FIRST_MATCH" {
		return nil, errors.New("secret policy FIRST_MATCH algorithm not verified")
	}
	var body json.RawMessage
	if err := c.Get(ctx, secretPolicyItem(ns, name)+"?response_format=2", &body); err != nil {
		if full.Spec.MarkedForDelete {
			full.Spec.RuleList = nil
			full.Spec.LegacyRuleList = nil
			return &full, nil
		}
		return nil, err
	}
	var envelope struct {
		SecretPolicy
		ReplaceForm *SecretPolicy `json:"replace_form"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, errors.New("invalid Secret Policy response")
	}
	p := envelope.SecretPolicy
	if envelope.ReplaceForm != nil {
		p = *envelope.ReplaceForm
	}
	if len(full.ResourceVersion) > 0 && len(envelope.ResourceVersion) > 0 && !bytes.Equal(full.ResourceVersion, envelope.ResourceVersion) {
		return nil, errors.New("secret policy changed during read normalization")
	}
	p.SystemMetadata = full.SystemMetadata
	p.ResourceVersion = full.ResourceVersion
	p.Spec.Algorithm = full.Spec.Algorithm
	p.Spec.MarkedForDelete = full.Spec.MarkedForDelete
	p.Spec.DeletionTime = full.Spec.DeletionTime
	if p.Metadata.Name != name || p.Metadata.Namespace != ns {
		return nil, errors.New("secret policy readback identity mismatch")
	}
	if p.Spec.RuleList == nil && p.Spec.LegacyRuleList != nil && !p.Spec.MarkedForDelete {
		rules := []SecretPolicyRule{}
		for _, ref := range p.Spec.LegacyRuleList.Rules {
			if ref.Name == "" || (ref.Namespace != "" && ref.Namespace != ns) {
				return nil, errors.New("invalid Secret Policy rule reference")
			}
			var rule SecretPolicyRule
			path := "/api/secret_management/namespaces/" + url.PathEscape(ns) + "/secret_policy_rules/" + url.PathEscape(ref.Name)
			if err := c.Get(ctx, path, &rule); err != nil {
				return nil, fmt.Errorf("secret policy rule read failed: %w", err)
			}
			if rule.Metadata.Name != ref.Name || (rule.Metadata.Namespace != "" && rule.Metadata.Namespace != ns) {
				return nil, errors.New("secret policy child identity mismatch")
			}
			rules = append(rules, rule)
		}
		p.Spec.RuleList = &SecretPolicyRuleList{Rules: rules}
	}
	if p.Spec.RuleList != nil {
		for i := range p.Spec.RuleList.Rules {
			rule := &p.Spec.RuleList.Rules[i]
			if rule.Spec.ClientName != "" || len(rule.Spec.ClientNameMatcher) != 0 {
				return nil, errors.New("unsupported Secret Policy name predicate; adoption refused")
			}
			if rule.Spec.Action == "" {
				rule.Spec.Action = "DENY"
			}
			if rule.Spec.Action != "ALLOW" && rule.Spec.Action != "DENY" {
				return nil, errors.New("unsupported Secret Policy chaining action; adoption refused")
			}
		}
	}
	p.Spec.LegacyRuleList = nil
	if p.Spec.RuleList == nil && !p.Spec.MarkedForDelete {
		return nil, errors.New("secret policy inline rules unavailable")
	}
	return &p, nil
}

// CreateSecretPolicy sends exactly one create. An ambiguous result needs named readback/import.
func (c *Client) CreateSecretPolicy(ctx context.Context, p *SecretPolicy) error {
	path := "/api/secret_management/namespaces/" + url.PathEscape(p.Metadata.Namespace) + "/secret_policys"
	_, err := c.doRequestWithRetry(ctx, http.MethodPost, path, secretPolicyWrite(p), false)
	return err
}
func (c *Client) ReplaceSecretPolicy(ctx context.Context, p *SecretPolicy) error {
	return c.PutOnce(ctx, secretPolicyItem(p.Metadata.Namespace, p.Metadata.Name), secretPolicyWrite(p), nil)
}
func (c *Client) DeleteSecretPolicy(ctx context.Context, ns, name string) error {
	_, err := c.doRequestWithRetry(ctx, http.MethodDelete, secretPolicyItem(ns, name), nil, false)
	return err
}

func secretPolicyWrite(p *SecretPolicy) map[string]any {
	body := map[string]any{"metadata": p.Metadata, "spec": map[string]any{
		"allow_f5xc":            p.Spec.AllowF5XC,
		"decrypt_cache_timeout": p.Spec.DecryptCacheTimeout,
		"rule_list":             p.Spec.RuleList,
	}}
	if len(p.ResourceVersion) > 0 {
		body["resource_version"] = p.ResourceVersion
	}
	return body
}
