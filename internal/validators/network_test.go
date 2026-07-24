// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// runStringValidator drives a validator.String against a plain string value and
// reports whether it produced a diagnostic error.
func runStringValidator(v validator.String, in string) bool {
	resp := &validator.StringResponse{}
	v.ValidateString(context.Background(),
		validator.StringRequest{ConfigValue: types.StringValue(in)}, resp)
	return resp.Diagnostics.HasError()
}

func TestIPv4Validator(t *testing.T) {
	v := IPv4Validator()
	// value -> wantErr
	for in, wantErr := range map[string]bool{
		"192.168.0.1":     false,
		"0.0.0.0":         false,
		"255.255.255.255": false,
		"::1":             true, // IPv6, not IPv4
		"2001:db8::1":     true,
		"not-an-ip":       true,
		"192.168.0.256":   true,
		"192.168.0.1/24":  true,  // CIDR, not a bare IP
		"":                false, // empty skipped (Optional fields)
	} {
		if got := runStringValidator(v, in); got != wantErr {
			t.Errorf("IPv4 %q: hasErr=%v want %v", in, got, wantErr)
		}
	}
}

func TestIPv6Validator(t *testing.T) {
	v := IPv6Validator()
	for in, wantErr := range map[string]bool{
		"::1":         false,
		"2001:db8::1": false,
		"fe80::1":     false,
		"192.168.0.1": true, // IPv4, not IPv6
		"not-an-ip":   true,
		"":            false,
	} {
		if got := runStringValidator(v, in); got != wantErr {
			t.Errorf("IPv6 %q: hasErr=%v want %v", in, got, wantErr)
		}
	}
}

func TestIPValidator(t *testing.T) {
	v := IPValidator()
	for in, wantErr := range map[string]bool{
		"192.168.0.1": false,
		"::1":         false,
		"2001:db8::1": false,
		"not-an-ip":   true,
		"1.2.3":       true,
		"":            false,
	} {
		if got := runStringValidator(v, in); got != wantErr {
			t.Errorf("IP %q: hasErr=%v want %v", in, got, wantErr)
		}
	}
}

func TestCIDRValidator(t *testing.T) {
	v := CIDRValidator()
	for in, wantErr := range map[string]bool{
		"192.168.0.0/24": false,
		"10.0.0.0/8":     false,
		"2001:db8::/32":  false,
		"192.168.0.1":    true, // bare IP, no prefix
		"not-a-cidr":     true,
		"192.168.0.0/33": true, // invalid prefix length
		"":               false,
	} {
		if got := runStringValidator(v, in); got != wantErr {
			t.Errorf("CIDR %q: hasErr=%v want %v", in, got, wantErr)
		}
	}
}

func TestMACValidator(t *testing.T) {
	v := MACValidator()
	for in, wantErr := range map[string]bool{
		"7C-1E-52-7F-F8-12": false,
		"7c:1e:52:7f:f8:12": false,
		"not-a-mac":         true,
		"7C-1E-52-7F-F8":    true, // too short
		"":                  false,
	} {
		if got := runStringValidator(v, in); got != wantErr {
			t.Errorf("MAC %q: hasErr=%v want %v", in, got, wantErr)
		}
	}
}
