// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestSMSv2BGPExpectedRouteRequiresSessionPath(t *testing.T) {
	for _, tc := range []struct {
		name     string
		paths    interface{}
		imported bool
		want     bool
	}{
		{"matching peer", []interface{}{map[string]interface{}{"peer": "169.254.10.1"}}, true, true},
		{"another session only", []interface{}{map[string]interface{}{"peer": "169.254.10.2"}}, true, false},
		{"both session paths", []interface{}{map[string]interface{}{"peer": "169.254.10.2"}, map[string]interface{}{"peer": "169.254.10.1"}}, true, true},
		{"missing paths", nil, true, false},
		{"empty paths", []interface{}{}, true, false},
		{"malformed path", []interface{}{"invalid"}, true, false},
		{"missing peer", []interface{}{map[string]interface{}{}}, true, false},
		{"exported route only", []interface{}{map[string]interface{}{"peer": "169.254.10.1"}}, false, false},
		{"canonical mapped peer", []interface{}{map[string]interface{}{"peer": "::ffff:169.254.10.1"}}, true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			configured, err := resolvedRuntimeFixture(runtimeConfiguration())
			if err != nil {
				t.Fatal(err)
			}
			expected := map[string]smsv2ExpectedPeerModel{}
			if d := bgpExpectedPeers(t).ElementsAs(context.Background(), &expected, false); d.HasError() {
				t.Fatal(d)
			}
			peers, err := extractSMSv2BGPPeers(bgpPeerFixture("Established"))
			if err != nil {
				t.Fatal(err)
			}
			routes := bgpRoutesFixture()
			table := routes["ver"].([]interface{})[0].(map[string]interface{})["ri_table"].([]interface{})[0].(map[string]interface{})["rt_table"].([]interface{})[0].(map[string]interface{})
			route := map[string]interface{}{"subnet": "10.10.0.0/16"}
			if tc.paths != nil {
				route["path"] = tc.paths
			}
			if tc.imported {
				table["imported"] = []interface{}{route}
			} else {
				table["imported"] = []interface{}{}
				table["exported"] = append(table["exported"].([]interface{}), route)
			}
			simplified := simplifiedRoutesFixture()
			n := simplified["ver_routes"].([]interface{})[0].(map[string]interface{})
			n["route"] = append(n["route"].([]interface{}), map[string]interface{}{"prefix": "10.10.0.0/16"})
			_, got, reason := convergeSMSv2BGP(expected, configured, peers, client.SMSv2Observation(routes), simplified, simplified)
			if got != tc.want {
				t.Fatalf("converged=%v want=%v reason=%q", got, tc.want, reason)
			}
		})
	}
}
