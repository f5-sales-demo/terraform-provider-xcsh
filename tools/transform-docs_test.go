// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

package main

import (
	"strings"
	"testing"
)

// Promotion is driven by document structure, not by punctuation: a bold line
// becomes an H4 only when the next non-blank line is the nested-block marker
// ("A `foo` block supports the following:"). An earlier version of this
// transform promoted every standalone bold line that did not end in a colon,
// which turned ordinary labels such as "**Note**" into spurious H4 headings and
// their anchors into registry navigation targets. Anything not introducing a
// block — colon-terminated labels included — stays bold.
func TestConvertBoldToH4HeadersPromotesOnlyNestedBlockHeadings(t *testing.T) {
	input := strings.Join([]string{
		"**Example (API format):**",
		"",
		"**Note**",
		"",
		"**Possible state values:**",
		"",
		"**CORS Policy**",
		"",
		"A `cors_policy` block supports the following:",
	}, "\n")
	want := strings.Join([]string{
		"**Example (API format):**",
		"",
		"**Note**",
		"",
		"**Possible state values:**",
		"",
		"#### CORS Policy",
		"",
		"A `cors_policy` block supports the following:",
	}, "\n")

	got := convertBoldToH4Headers(input)
	if got != want {
		t.Fatalf("convertBoldToH4Headers() = %q, want %q", got, want)
	}
	if second := convertBoldToH4Headers(got); second != got {
		t.Fatalf("convertBoldToH4Headers() is not idempotent:\nfirst:  %q\nsecond: %q", got, second)
	}
}

func TestEscapeEmphasisMarkersTouchesOnlyMarkdownProse(t *testing.T) {
	input := strings.Join([]string{
		"Contact alerts@example.com for object * values and **Rules**.",
		"Already protected: `alerts@example.com`.",
		"```yaml",
		"email_address: alerts@example.com",
		"```",
	}, "\n")
	want := strings.Join([]string{
		"Contact `alerts@example.com` for object \\* values and **Rules**.",
		"Already protected: `alerts@example.com`.",
		"```yaml",
		"email_address: alerts@example.com",
		"```",
	}, "\n")

	first := escapeEmphasisMarkersInContent(input)
	if first != want {
		t.Fatalf("escapeEmphasisMarkersInContent() = %q, want %q", first, want)
	}
	if second := escapeEmphasisMarkersInContent(first); second != first {
		t.Fatalf("escapeEmphasisMarkersInContent() is not idempotent:\nfirst:  %q\nsecond: %q", first, second)
	}
}

func TestNormalizeDocumentProseIsIdempotent(t *testing.T) {
	input := "# xcsh_example (Data Source)\n\nManages a javascript API."
	first := normalizeDocumentProse(input, "docs/data-sources/example.md")
	second := normalizeDocumentProse(first, "docs/data-sources/example.md")

	if second != first {
		t.Fatalf("normalizeDocumentProse() is not idempotent:\nfirst:  %q\nsecond: %q", first, second)
	}
	if strings.Contains(first, "Manages") || !strings.Contains(first, "Retrieves information") {
		t.Fatalf("normalizeDocumentProse() did not normalize data-source prose: %q", first)
	}
}

func TestDescriptionDeduplicationPreservesCollapsedSection(t *testing.T) {
	input := strings.Join([]string{
		"#### One Two Three Four Five Six Seven JavaScript",
		"",
		"Nested content.",
		"",
		"## Import",
		"",
	}, "\n")

	first := applyDescriptionDeduplication(input)
	second := applyDescriptionDeduplication(first)
	if second != first {
		t.Fatalf("applyDescriptionDeduplication() is not idempotent:\nfirst:\n%s\nsecond:\n%s", first, second)
	}
	if !strings.Contains(first, "\n---\n\n## Common Types\n") {
		t.Fatalf("applyDescriptionDeduplication() lost the Common Types separator: %q", first)
	}
}
