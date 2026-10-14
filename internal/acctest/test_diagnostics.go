// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"os"
	"strings"
	"sync"
	"testing"
)

var acceptanceDiagnosticTest struct {
	sync.RWMutex
	test *testing.T
}

func bindDiagnosticTest(t *testing.T) {
	acceptanceDiagnosticTest.Lock()
	acceptanceDiagnosticTest.test = t
	acceptanceDiagnosticTest.Unlock()
}
func logDiagnosticFact(fact string) {
	acceptanceDiagnosticTest.RLock()
	t := acceptanceDiagnosticTest.test
	acceptanceDiagnosticTest.RUnlock()
	if t != nil {
		t.Logf("[API_DIAGNOSTIC:%s]", fact)
	}
}
func CaptureImportFieldDiagnostics(t *testing.T, typ, address string) (resource.TestCheckFunc, resource.ImportStateCheckFunc) {
	var prior map[string]string
	capture := func(state *terraform.State) error {
		rs := state.RootModule().Resources[address]
		if rs == nil || rs.Primary == nil {
			return nil
		}
		prior = map[string]string{}
		for k, v := range rs.Primary.Attributes {
			prior[k] = v
		}
		return nil
	}
	check := func(states []*terraform.InstanceState) error {
		schema, err := LoadFixtureSchema(os.Getenv("XCSH_ACC_SCHEMA"))
		if err != nil {
			return nil
		}
		block, ok := schema.Resources[typ]
		if !ok {
			return nil
		}
		allowed := map[string]bool{}
		var walk func(fixtureBlockSchema, string)
		walk = func(b fixtureBlockSchema, prefix string) {
			for name := range b.Attributes {
				allowed[prefix+name] = true
			}
			for name, nested := range b.Blocks {
				allowed[prefix+name] = true
				walk(nested.Block, prefix+name+".")
			}
		}
		walk(block, "")
		for _, state := range states {
			for key, before := range prior {
				if state.Attributes[key] == before {
					continue
				}
				normalized := normalizeSchemaPath(key)
				if allowed[normalized] {
					t.Logf("[SCHEMA_FIELD:%s]", normalized)
				}
			}
			for key, after := range state.Attributes {
				if prior[key] == after {
					continue
				}
				normalized := normalizeSchemaPath(key)
				if allowed[normalized] {
					t.Logf("[SCHEMA_FIELD:%s]", normalized)
				}
			}
		}
		return nil
	}
	return capture, check
}
func normalizeSchemaPath(key string) string {
	parts := []string{}
	for _, p := range strings.Split(key, ".") {
		if p == "#" || p == "%" {
			continue
		}
		numeric := p != ""
		for _, r := range p {
			if r < '0' || r > '9' {
				numeric = false
			}
		}
		if !numeric {
			parts = append(parts, p)
		}
	}
	return strings.Join(parts, ".")
}
