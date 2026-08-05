package docfmt

import (
	"strings"
	"testing"
)

func TestCanonicalizeNetworkLiterals(t *testing.T) {
	t.Parallel()

	publicAddress := strings.Join([]string{"8", "8", "4", "4"}, ".")
	publicPrefix := strings.Join([]string{"11", "0", "0", "0"}, ".") + "/8"
	input := strings.Join([]string{
		"public=" + publicAddress,
		"prefix=" + publicPrefix,
		"private=10.0.0.1",
		"documentation=198.51.100.7",
		"invalid=999.1.1.1",
	}, " ")
	want := "public=192.0.2.1 prefix=192.0.2.0/24 private=10.0.0.1 " +
		"documentation=198.51.100.7 invalid=999.1.1.1"

	if got := CanonicalizeNetworkLiterals(input); got != want {
		t.Fatalf("CanonicalizeNetworkLiterals() = %q, want %q", got, want)
	}
	if again := CanonicalizeNetworkLiterals(want); again != want {
		t.Fatalf("CanonicalizeNetworkLiterals() is not idempotent: %q", again)
	}
}
