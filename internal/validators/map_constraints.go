// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package validators

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type mapScope struct {
	Type      string      `json:"type"`
	MinLength *int        `json:"minLength"`
	MaxLength *int        `json:"maxLength"`
	Pattern   string      `json:"pattern"`
	Format    string      `json:"format"`
	Minimum   *uint64     `json:"minimum"`
	Maximum   *uint64     `json:"maximum"`
	Ranges    [][2]uint64 `json:"ranges"`
}
type mapContract struct {
	Keys        *mapScope `json:"keys"`
	Values      *mapScope `json:"values"`
	Cardinality struct {
		Min *int `json:"minProperties"`
		Max *int `json:"maxProperties"`
	} `json:"cardinality"`
	CrossEntry struct {
		Unique bool `json:"uniqueValues"`
	} `json:"crossEntry"`
}
type mapConstraintsValidator struct {
	contract mapContract
	err      error
}

// MapConstraintsValidator preserves key, value and cardinality scopes in plan validation.
func MapConstraintsValidator(raw string) validator.Map {
	v := mapConstraintsValidator{}
	v.err = json.Unmarshal([]byte(raw), &v.contract)
	return v
}
func (v mapConstraintsValidator) Description(context.Context) string {
	return "map keys, values and pair count must satisfy the published API contract"
}
func (v mapConstraintsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func validateMapText(text string, scope *mapScope) error {
	if scope == nil {
		return nil
	}
	length := utf8.RuneCountInString(text)
	if scope.MinLength != nil && length < *scope.MinLength {
		return fmt.Errorf("below minimum length %d", *scope.MinLength)
	}
	if scope.MaxLength != nil && length > *scope.MaxLength {
		return fmt.Errorf("exceeds maximum length %d", *scope.MaxLength)
	}
	if scope.Pattern != "" {
		r, err := regexp.Compile(scope.Pattern)
		if err != nil {
			return err
		}
		if !r.MatchString(text) {
			return fmt.Errorf("does not match required pattern")
		}
	}
	if scope.Type == "uint32-string" {
		n, err := strconv.ParseUint(text, 10, 32)
		if err != nil {
			return fmt.Errorf("must be a uint32 key")
		}
		if scope.Minimum != nil && n < *scope.Minimum {
			return fmt.Errorf("below minimum key")
		}
		if scope.Maximum != nil && n > *scope.Maximum {
			return fmt.Errorf("above maximum key")
		}
		if len(scope.Ranges) > 0 {
			matched := false
			for _, r := range scope.Ranges {
				if n >= r[0] && n <= r[1] {
					matched = true
				}
			}
			if !matched {
				return fmt.Errorf("outside exact key ranges")
			}
		}
	}
	switch scope.Format {
	case "ipv4":
		if ip := net.ParseIP(text); ip == nil || ip.To4() == nil {
			return fmt.Errorf("must be IPv4")
		}
	case "ipv6":
		if ip := net.ParseIP(text); ip == nil || ip.To4() != nil {
			return fmt.Errorf("must be IPv6")
		}
	case "ip-address":
		if net.ParseIP(text) == nil {
			return fmt.Errorf("must be an IP address")
		}
	case "mac-address":
		if _, err := net.ParseMAC(text); err != nil {
			return err
		}
	case "regex":
		if _, err := regexp.Compile(text); err != nil {
			return err
		}
	case "uri-reference":
		if _, err := url.Parse(text); err != nil {
			return err
		}
		if regexp.MustCompile(`[\x00-\x20\x7f-\xff]`).MatchString(text) {
			return fmt.Errorf("must be a URI-reference")
		}
	case "k8s-label-value":
		if len(text) > 63 || !regexp.MustCompile(`^(?:[A-Za-z0-9](?:[A-Za-z0-9_.-]*[A-Za-z0-9])?)?$`).MatchString(text) {
			return fmt.Errorf("must be a Kubernetes label value")
		}
	}
	return nil
}
func (v mapConstraintsValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if v.err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid Map Constraint Contract", v.err.Error())
		return
	}
	entries := req.ConfigValue.Elements()
	if v.contract.Cardinality.Min != nil && len(entries) < *v.contract.Cardinality.Min {
		resp.Diagnostics.AddAttributeError(req.Path, "Too Few Map Pairs", v.Description(ctx))
	}
	if v.contract.Cardinality.Max != nil && len(entries) > *v.contract.Cardinality.Max {
		resp.Diagnostics.AddAttributeError(req.Path, "Too Many Map Pairs", v.Description(ctx))
	}
	seen := map[string]bool{}
	for key, value := range entries {
		if err := validateMapText(key, v.contract.Keys); err != nil {
			resp.Diagnostics.AddAttributeError(req.Path.AtMapKey(key), "Invalid Map Key", err.Error())
		}
		if value.IsNull() || value.IsUnknown() {
			continue
		}
		str, ok := value.(types.String)
		if !ok {
			if v.contract.Values != nil {
				resp.Diagnostics.AddAttributeError(req.Path.AtMapKey(key), "Invalid Map Value", "Expected string map value")
			}
			continue
		}
		text := str.ValueString()
		if err := validateMapText(text, v.contract.Values); err != nil {
			resp.Diagnostics.AddAttributeError(req.Path.AtMapKey(key), "Invalid Map Value", err.Error())
		}
		if v.contract.CrossEntry.Unique && seen[text] {
			resp.Diagnostics.AddAttributeError(req.Path.AtMapKey(key), "Duplicate Map Value", "Map values must be unique")
		}
		seen[text] = true
	}
}
