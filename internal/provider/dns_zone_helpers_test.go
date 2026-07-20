// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import "testing"

// group is a tiny helper to build an rr_set_group element as the flatten sees it
// (a map[string]interface{} decoded from the API response).
func rrSetGroup(name string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{"name": name},
	}
}

func TestIsSystemManagedRrSetGroup(t *testing.T) {
	tests := []struct {
		name string
		item map[string]interface{}
		want bool
	}{
		{
			name: "reserved system group",
			item: map[string]interface{}{
				"metadata": map[string]interface{}{
					"name":        "x-ves-io-managed",
					"description": "Special RRSetGroup managed by F5XC",
				},
			},
			want: true,
		},
		{name: "user group", item: rrSetGroup("demo-records"), want: false},
		{name: "no metadata", item: map[string]interface{}{}, want: false},
		{
			name: "metadata missing name",
			item: map[string]interface{}{"metadata": map[string]interface{}{}},
			want: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isSystemManagedRrSetGroup(tc.item); got != tc.want {
				t.Fatalf("isSystemManagedRrSetGroup(%v) = %v, want %v", tc.item, got, tc.want)
			}
		})
	}
}

func TestFilterSystemManagedRrSetGroups(t *testing.T) {
	sys := map[string]interface{}{
		"metadata": map[string]interface{}{"name": "x-ves-io-managed"},
	}

	t.Run("drops the system group, keeps the user group", func(t *testing.T) {
		in := []interface{}{rrSetGroup("demo-records"), sys}
		got := filterSystemManagedRrSetGroups(in)
		if len(got) != 1 {
			t.Fatalf("len = %d, want 1 (%v)", len(got), got)
		}
		md := got[0].(map[string]interface{})["metadata"].(map[string]interface{})
		if md["name"] != "demo-records" {
			t.Fatalf("kept %q, want demo-records", md["name"])
		}
	})

	t.Run("preserves user group order when system group is first", func(t *testing.T) {
		in := []interface{}{sys, rrSetGroup("demo-records")}
		got := filterSystemManagedRrSetGroups(in)
		if len(got) != 1 {
			t.Fatalf("len = %d, want 1", len(got))
		}
		md := got[0].(map[string]interface{})["metadata"].(map[string]interface{})
		if md["name"] != "demo-records" {
			t.Fatalf("index 0 = %q, want demo-records", md["name"])
		}
	})

	t.Run("no system group leaves the list unchanged", func(t *testing.T) {
		in := []interface{}{rrSetGroup("demo-records"), rrSetGroup("more-records")}
		got := filterSystemManagedRrSetGroups(in)
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
	})

	t.Run("empty input yields empty output", func(t *testing.T) {
		if got := filterSystemManagedRrSetGroups(nil); len(got) != 0 {
			t.Fatalf("len = %d, want 0", len(got))
		}
	})
}
