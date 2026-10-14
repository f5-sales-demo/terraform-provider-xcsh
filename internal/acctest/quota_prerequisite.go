// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"strings"
)

type quotaPrerequisite struct{ Unconfigured, Exhausted bool }

func readQuotaPrerequisite(ctx context.Context, c *client.Client, typ, namespace string) (quotaPrerequisite, error) {
	var document struct {
		Objects map[string]json.RawMessage `json:"objects"`
	}
	if err := c.Get(ctx, "/api/web/namespaces/system/quota/limits", &document); err != nil {
		return quotaPrerequisite{}, err
	}
	if document.Objects == nil {
		return quotaPrerequisite{}, errors.New("quota inventory malformed")
	}
	key := strings.TrimPrefix(typ, "xcsh_")
	raw, found := document.Objects[key]
	if !found {
		return quotaPrerequisite{Unconfigured: true}, nil
	}
	var entry struct {
		Limit struct {
			Maximum *int `json:"maximum"`
		} `json:"limit"`
	}
	if json.Unmarshal(raw, &entry) != nil || entry.Limit.Maximum == nil {
		return quotaPrerequisite{}, errors.New("quota limit unavailable")
	}
	route, ok := ownedCollectionRoutes[typ]
	if !ok {
		return quotaPrerequisite{}, errors.New("quota collection unavailable")
	}
	var listing struct {
		Items []json.RawMessage `json:"items"`
	}
	if err := c.Get(ctx, strings.ReplaceAll(route, "%s", namespace), &listing); err != nil {
		return quotaPrerequisite{}, err
	}
	if listing.Items == nil {
		return quotaPrerequisite{}, errors.New("quota collection malformed")
	}
	maximum := *entry.Limit.Maximum
	return quotaPrerequisite{Exhausted: maximum >= 0 && len(listing.Items) >= maximum}, nil
}
