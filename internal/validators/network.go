// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// The network validators below translate a spec `format` label into a plan-time
// check. Each skips null/unknown/empty values (Optional fields), and validates
// non-empty values with the Go `net` stdlib rather than regex so the check is
// robust and identical on every workstation.

// IPv4Validator returns a validator that requires a valid IPv4 address.
func IPv4Validator() validator.String { return &ipv4Validator{} }

type ipv4Validator struct{}

func (v ipv4Validator) Description(ctx context.Context) string {
	return "must be a valid IPv4 address"
}

func (v ipv4Validator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipv4Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	value, ok := nonEmptyString(req)
	if !ok {
		return
	}
	if ip := net.ParseIP(value); ip == nil || ip.To4() == nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Address",
			fmt.Sprintf("Value %q is not a valid IPv4 address (e.g. 192.168.0.1).", value),
		)
	}
}

// IPv6Validator returns a validator that requires a valid IPv6 address.
func IPv6Validator() validator.String { return &ipv6Validator{} }

type ipv6Validator struct{}

func (v ipv6Validator) Description(ctx context.Context) string {
	return "must be a valid IPv6 address"
}

func (v ipv6Validator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipv6Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	value, ok := nonEmptyString(req)
	if !ok {
		return
	}
	if ip := net.ParseIP(value); ip == nil || ip.To4() != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Address",
			fmt.Sprintf("Value %q is not a valid IPv6 address (e.g. 2001:db8::1).", value),
		)
	}
}

// IPValidator returns a validator that requires a valid IP address (v4 or v6).
func IPValidator() validator.String { return &ipValidator{} }

type ipValidator struct{}

func (v ipValidator) Description(ctx context.Context) string {
	return "must be a valid IP address (IPv4 or IPv6)"
}

func (v ipValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	value, ok := nonEmptyString(req)
	if !ok {
		return
	}
	if net.ParseIP(value) == nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Address",
			fmt.Sprintf("Value %q is not a valid IP address (IPv4 or IPv6).", value),
		)
	}
}

// CIDRValidator returns a validator that requires a valid CIDR range.
func CIDRValidator() validator.String { return &cidrValidator{} }

type cidrValidator struct{}

func (v cidrValidator) Description(ctx context.Context) string {
	return "must be a valid CIDR range (e.g. 192.168.0.0/24)"
}

func (v cidrValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v cidrValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	value, ok := nonEmptyString(req)
	if !ok {
		return
	}
	if _, _, err := net.ParseCIDR(value); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid CIDR Range",
			fmt.Sprintf("Value %q is not a valid CIDR range (e.g. 192.168.0.0/24 or 2001:db8::/32).", value),
		)
	}
}

// MACValidator returns a validator that requires a valid MAC (hardware) address.
func MACValidator() validator.String { return &macValidator{} }

type macValidator struct{}

func (v macValidator) Description(ctx context.Context) string {
	return "must be a valid MAC address (e.g. 7C-1E-52-7F-F8-12)"
}

func (v macValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v macValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	value, ok := nonEmptyString(req)
	if !ok {
		return
	}
	if _, err := net.ParseMAC(value); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid MAC Address",
			fmt.Sprintf("Value %q is not a valid MAC address (e.g. 7C-1E-52-7F-F8-12).", value),
		)
	}
}

// nonEmptyString returns the config value and true when it is a known, non-null,
// non-empty string worth validating; otherwise it returns ("", false) so callers
// skip null/unknown/empty (Optional) values.
func nonEmptyString(req validator.StringRequest) (string, bool) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return "", false
	}
	value := req.ConfigValue.ValueString()
	if value == "" {
		return "", false
	}
	return value, true
}
