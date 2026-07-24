// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package docsterm

import (
	"strings"
	"testing"
)

// TestFixUpstreamTerminology_PreservesEncodingBase64 is the regression test for
// the S2b fix: the terminology normaliser must NOT lowercase "Base64" inside the
// real API enum token "EncodingBase64" (which would corrupt it into
// "Encodingbase64"). Go's regexp has no lookbehind, so a naive Base64->base64
// rewrite cannot exclude the enum token; the rewrite was therefore removed.
func TestFixUpstreamTerminology_PreservesEncodingBase64(t *testing.T) {
	input := "Possible values are `EncodingNone`, `EncodingBase64`"
	got := FixUpstreamTerminology(input)

	if strings.Contains(got, "Encodingbase64") {
		t.Errorf("FixUpstreamTerminology corrupted enum token: got %q, must not contain %q", got, "Encodingbase64")
	}
	if !strings.Contains(got, "EncodingBase64") {
		t.Errorf("FixUpstreamTerminology dropped enum token: got %q, must contain %q", got, "EncodingBase64")
	}
}

// TestFixUpstreamTerminology_NormalisesInternet is a characterization test proving
// the function still performs a real terminology transform (lowercasing prose
// "Internet" -> "internet"), i.e. the extraction did not gut the logic.
func TestFixUpstreamTerminology_NormalisesInternet(t *testing.T) {
	input := "Traffic reaches the Internet directly."
	want := "Traffic reaches the internet directly."
	got := FixUpstreamTerminology(input)
	if got != want {
		t.Errorf("FixUpstreamTerminology(%q) = %q, want %q", input, got, want)
	}
}
