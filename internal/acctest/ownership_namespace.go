// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
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
	var response struct {
		Items []struct {
			ErrorMessage string `json:"error_message"`
		} `json:"items"`
	}
	if err := c.Post(ctx, "/api/web/namespaces/"+r.Name+"/cascade_delete", struct{}{}, &response); err != nil {
		return err
	}
	if response.Items == nil {
		return errors.New("namespace cascade response malformed")
	}
	for _, item := range response.Items {
		if item.ErrorMessage != "" {
			return errors.New("namespace cascade item failed")
		}
	}
	return nil
}
