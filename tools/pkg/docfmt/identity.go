package docfmt

import (
	"regexp"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/identity"
)

const identityFieldPattern = `tenant(?:_name|_id)?|customer(?:_name|_id)?|` +
	`account(?:_name|_id)?|subscription(?:_name|_id)?|project(?:_name|_id)?|namespace`

var quotedIdentityLiteralPattern = regexp.MustCompile(
	`(?m)(\b(` + identityFieldPattern + `)\b\s*=\s*)"([^"]*)"`,
)

var yamlIdentityLiteralPattern = regexp.MustCompile(
	`(?m)^(\s*(` + identityFieldPattern + `)\s*:\s*)([^#\r\n]+?)\s*$`,
)

// CanonicalizeIdentityLiterals replaces captured organization identifiers in
// generated HCL snippets while preserving reserved namespaces and expressions.
func CanonicalizeIdentityLiterals(content string) string {
	content = quotedIdentityLiteralPattern.ReplaceAllStringFunc(content, func(match string) string {
		parts := quotedIdentityLiteralPattern.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}
		replacement := identity.Canonical(parts[2], parts[3])
		return parts[1] + `"` + replacement + `"`
	})
	return yamlIdentityLiteralPattern.ReplaceAllStringFunc(content, func(match string) string {
		parts := yamlIdentityLiteralPattern.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}
		return parts[1] + identity.Canonical(parts[2], parts[3])
	})
}
