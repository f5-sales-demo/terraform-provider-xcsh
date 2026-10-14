// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestAssertionFieldsRetainSchemaNamesWithoutValues(t *testing.T) {
	original := publicSchemaFields
	t.Cleanup(func() { publicSchemaFields = original })
	publicSchemaFields = map[string]bool{"name": true, "labels": true}
	for _, input := range []string{
		`Check failed: Attribute 'name' expected "SECRET_SENTINEL", got "OTHER_SENTINEL"`,
		`Check failed: Attribute 'labels.environment' expected "SECRET_SENTINEL", got "OTHER_SENTINEL"`,
	} {
		f := classifyFailureFingerprint(input)
		want := []string{"name"}
		if strings.Contains(input, "labels.environment") {
			want = []string{"labels"}
		}
		if f.Class != FailureClassAssertion || !reflect.DeepEqual(f.SchemaFields, want) {
			t.Fatal("assertion schema field lost")
		}
		if strings.Contains(reportJSON(t, &Report{FailedTests: []TestResult{{FailureFingerprint: &f}}}), "SENTINEL") {
			t.Fatal("assertion value leaked")
		}
	}
}
