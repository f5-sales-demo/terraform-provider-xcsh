// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"strings"
)

// deleteOwnedNamespace uses the command endpoint only after exact namespace
// identity and an explicitly empty referring-object inventory are verified.
func deleteOwnedNamespace(ctx context.Context, c *client.Client, r OwnedResource) error {
	var document struct {
		Metadata       client.Metadata `json:"metadata"`
		SystemMetadata struct {
			UID string `json:"uid"`
		} `json:"system_metadata"`
		ReferringObjects []json.RawMessage `json:"referring_objects"`
	}
	if err := c.Get(ctx, "/api/web/namespaces/"+r.Name, &document); err != nil {
		return err
	}
	uid := document.SystemMetadata.UID
	if uid == "" {
		uid = document.Metadata.UID
	}
	if document.Metadata.Name != r.Name || uid == "" || r.UID == "" || ownershipUIDDigest(uid) != r.UID || document.ReferringObjects == nil || len(document.ReferringObjects) > 0 {
		return errors.New("namespace cascade ownership or empty inventory unverified")
	}
	var envelope map[string]json.RawMessage
	var response struct {
		Items []struct {
			ErrorMessage string `json:"error_message"`
		} `json:"items"`
	}
	if err := c.Post(ctx, "/api/web/namespaces/"+r.Name+"/cascade_delete", struct{}{}, &envelope); err != nil {
		if namespaceDeletionInProgress(err) {
			return nil
		}
		return err
	}
	if envelope == nil {
		return errors.New("namespace cascade response malformed")
	}
	for key := range envelope {
		if key != "items" {
			return errors.New("namespace cascade response malformed")
		}
	}
	// Proto JSON omits empty repeated fields. An empty object acknowledges the
	// command; the ownership journal still requires typed absence in both passes.
	if raw, present := envelope["items"]; present {
		if string(raw) == "null" || json.Unmarshal(raw, &response.Items) != nil {
			return errors.New("namespace cascade response malformed")
		}
	}
	for _, item := range response.Items {
		if item.ErrorMessage != "" {
			return errors.New("namespace cascade item failed")
		}
	}
	return nil
}

func namespaceDeletionInProgress(err error) bool {
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 400 {
		return false
	}
	message := strings.ToLower(apiErr.Message)
	if raw, ok := apiErr.Details["raw_response"].(string); ok {
		var envelope struct {
			Message string `json:"message"`
		}
		if json.Unmarshal([]byte(raw), &envelope) != nil {
			return false
		}
		message = strings.ToLower(envelope.Message)
	}
	if !strings.Contains(message, "namespace") {
		return false
	}
	for _, signature := range []string{"already mark for deletion", "already marked for deletion", "already scheduled for deletion", "already queued for deletion", "already pending deletion", "already undergoing deletion"} {
		if strings.Contains(message, signature) {
			return true
		}
	}
	return false
}
