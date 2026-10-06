package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

// HasHTTPStatus classifies transport status without interpreting backend message text.
func HasHTTPStatus(err error, status int) bool {
	var apiErr *xcsherrors.XCSHError
	return errors.As(err, &apiErr) && apiErr.StatusCode == status
}

func protectedDomainVerificationError(reason string) error {
	return fmt.Errorf("protected-domain verification unavailable: %s; retain ownership and verify the namespace collection projection", reason)
}

func validProtectedDomainIdentity(namespace, name, root string) error {
	if namespace == "" || name == "" || strings.ContainsAny(namespace+name, "/?#") {
		return protectedDomainVerificationError("invalid namespace or object identity")
	}
	if root != "" && (strings.TrimSpace(root) != root || strings.ContainsAny(root, "/?# \t\r\n") || !strings.Contains(root, ".")) {
		return protectedDomainVerificationError("invalid protected root")
	}
	return nil
}

func validateProtectedDomainProjection(result *ProtectedDomain, namespace, name, root string) error {
	if result == nil || result.Spec == nil {
		return protectedDomainVerificationError("missing domain projection")
	}
	observed, ok := result.Spec["protected_domain"].(string)
	if !ok || observed == "" {
		return protectedDomainVerificationError("missing protected root")
	}
	if root != "" && observed != root {
		return protectedDomainVerificationError("protected root does not match configured ownership")
	}
	if result.Metadata.Namespace != "" && result.Metadata.Namespace != namespace {
		return protectedDomainVerificationError("conflicting namespace")
	}
	if result.Metadata.Name != "" && result.Metadata.Name != name {
		return protectedDomainVerificationError("conflicting object identity")
	}
	return nil
}

// VerifyProtectedDomain prefers GET-by-name and uses the namespace List RPC only
// for typed HTTP 501. The released RPC lists the set in one namespace, exposes
// repeated report_fields and collection errors, and has no pagination input or
// continuation output. An inferred enrichment maxItems is not a pagination cap.
// Unknown response envelope fields fail closed rather than guessing continuation.
func (c *Client) VerifyProtectedDomain(ctx context.Context, namespace, name, root string) (*ProtectedDomain, error) {
	if err := validProtectedDomainIdentity(namespace, name, root); err != nil {
		return nil, err
	}
	result, err := c.GetProtectedDomain(ctx, namespace, name)
	if err == nil {
		if err := validateProtectedDomainProjection(result, namespace, name, root); err != nil {
			return nil, err
		}
		return result, nil
	}
	if !HasHTTPStatus(err, http.StatusNotImplemented) {
		return nil, err
	}

	query := url.Values{}
	for _, field := range []string{"get_spec", "metadata", "name", "namespace", "uid"} {
		query.Add("report_fields", field)
	}
	endpoint := fmt.Sprintf("/api/shape/csd/namespaces/%s/protected_domains?%s", url.PathEscape(namespace), query.Encode())
	var envelope map[string]json.RawMessage
	if err := c.Get(ctx, endpoint, &envelope); err != nil {
		// Collection 404 is not authoritative absence of an individual object.
		return nil, protectedDomainVerificationError("namespace collection request failed")
	}
	for key := range envelope {
		if key != "items" && key != "errors" {
			return nil, protectedDomainVerificationError("unknown collection coverage")
		}
	}
	if raw, ok := envelope["errors"]; ok {
		var collectionErrors []json.RawMessage
		if string(raw) == "null" || json.Unmarshal(raw, &collectionErrors) != nil || len(collectionErrors) != 0 {
			return nil, protectedDomainVerificationError("collection reports errors")
		}
	}
	rawItems, ok := envelope["items"]
	if !ok || string(rawItems) == "null" {
		return nil, protectedDomainVerificationError("missing collection items")
	}
	var items []struct {
		Metadata  Metadata               `json:"metadata"`
		Name      string                 `json:"name"`
		Namespace string                 `json:"namespace"`
		UID       string                 `json:"uid"`
		GetSpec   map[string]interface{} `json:"get_spec"`
	}
	if json.Unmarshal(rawItems, &items) != nil {
		return nil, protectedDomainVerificationError("malformed collection projection")
	}
	var match *ProtectedDomain
	roots := map[string]bool{}
	for _, item := range items {
		observed, ok := item.GetSpec["protected_domain"].(string)
		if !ok || observed == "" {
			return nil, protectedDomainVerificationError("incomplete item projection")
		}
		if roots[observed] {
			return nil, protectedDomainVerificationError("duplicate protected roots")
		}
		roots[observed] = true
		if (item.Namespace != "" && item.Namespace != namespace) || (item.Metadata.Namespace != "" && item.Metadata.Namespace != namespace) {
			return nil, protectedDomainVerificationError("collection contains conflicting scope")
		}
		if item.Name != "" && item.Metadata.Name != "" && item.Name != item.Metadata.Name {
			return nil, protectedDomainVerificationError("conflicting collection identifiers")
		}
		metadata := item.Metadata
		if metadata.Name == "" {
			metadata.Name = item.Name
		}
		if metadata.Namespace == "" {
			metadata.Namespace = item.Namespace
		}
		selected := root != "" && observed == root
		if root == "" {
			if metadata.Name == "" {
				return nil, protectedDomainVerificationError("collection lacks lookup identifiers")
			}
			selected = metadata.Name == name
		}
		if !selected {
			if metadata.Name == name {
				return nil, protectedDomainVerificationError("owned name has a different protected root")
			}
			continue
		}
		candidate := &ProtectedDomain{Metadata: metadata, Spec: item.GetSpec}
		if err := validateProtectedDomainProjection(candidate, namespace, name, root); err != nil {
			return nil, err
		}
		if match != nil {
			return nil, protectedDomainVerificationError("ambiguous registration")
		}
		match = candidate
	}
	if match == nil {
		return nil, xcsherrors.NewAPIError(http.StatusNotFound, nil, "protected_domain", "verified namespace collection")
	}
	return match, nil
}

// ValidateProtectedDomainCreateResponse rejects conflicting successful writes.
// Blank response metadata/spec is permitted only pending authoritative verification.
func ValidateProtectedDomainCreateResponse(result *ProtectedDomain, namespace, name, root string) error {
	if result == nil {
		return protectedDomainVerificationError("missing create response")
	}
	if result.Metadata.Name != "" && result.Metadata.Name != name {
		return protectedDomainVerificationError("create returned conflicting identity")
	}
	if result.Metadata.Namespace != "" && result.Metadata.Namespace != namespace {
		return protectedDomainVerificationError("create returned conflicting namespace")
	}
	if result.Spec != nil {
		if observed, ok := result.Spec["protected_domain"]; ok && observed != root {
			return protectedDomainVerificationError("create returned conflicting protected root")
		}
	}
	return nil
}

// CreateProtectedDomainRegistration distinguishes a landed HTTP-successful POST
// from a response decoding failure, so Terraform can retain ownership and retry
// verification without replaying the create.
func (c *Client) CreateProtectedDomainRegistration(ctx context.Context, resource *ProtectedDomain) (*ProtectedDomain, bool, error) {
	endpoint := fmt.Sprintf("/api/shape/csd/namespaces/%s/protected_domains", url.PathEscape(resource.Metadata.Namespace))
	body, err := c.doRequestWithRetry(ctx, http.MethodPost, endpoint, resource, false)
	if err != nil {
		return nil, false, err
	}
	result := &ProtectedDomain{}
	if len(body) == 0 || json.Unmarshal(body, result) != nil {
		return nil, true, protectedDomainVerificationError("successful create response is incomplete")
	}
	return result, true, nil
}

// ProtectedDomainDiagnostic exposes verification context without response bodies.
func ProtectedDomainDiagnostic(err error) string {
	var apiErr *xcsherrors.XCSHError
	if errors.As(err, &apiErr) {
		return fmt.Sprintf("Protected-domain operation failed (HTTP %d). Verify permissions, namespace scope and exact root before another mutation.", apiErr.StatusCode)
	}
	return "Protected-domain verification is unavailable or ambiguous. Retain ownership; verify the exact namespace collection, projection and root before another mutation."
}

// DeleteProtectedDomainRegistration uses the observed domain-keyed backend
// contract. Terraform's configured name remains its ownership key. Verify the
// complete namespace collection before returning deletion success.
func (c *Client) DeleteProtectedDomainRegistration(ctx context.Context, namespace, name, root string) error {
	if err := validProtectedDomainIdentity(namespace, name, root); err != nil {
		return err
	}
	if root == "" {
		return protectedDomainVerificationError("deletion requires the configured protected root")
	}
	endpoint := fmt.Sprintf("/api/shape/csd/namespaces/%s/protected_domains/%s", url.PathEscape(namespace), url.PathEscape(root))
	if err := c.Delete(ctx, endpoint); err != nil && !HasHTTPStatus(err, http.StatusNotFound) {
		return err
	}
	_, err := c.VerifyProtectedDomain(ctx, namespace, name, root)
	if HasHTTPStatus(err, http.StatusNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return protectedDomainVerificationError("delete completed but the registration is still present")
}
