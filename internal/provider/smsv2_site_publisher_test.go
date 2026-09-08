// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Synthetic identities preserve the site-publisher format observed on three AWS SMSv2 sites.
func sitePublisherPhysicalFixture() client.SMSv2Observation {
	v := runtimePhysicalLinkStatus()
	v["system_metadata"] = map[string]interface{}{"uid": "physical-site-uid"}
	status := v["status"].([]interface{})[0].(map[string]interface{})
	metadata := physicalFixtureMetadata(v)
	metadata["creator_id"] = "lab-site"
	metadata["status_id"] = "master-0_SiteStatusMgr"
	status["ver_status"].(map[string]interface{})["ver_instance_name"] = "master-0-lab-site"
	status["object_refs"] = []interface{}{map[string]interface{}{"kind": "ves.io.vega.cfg.site.Object", "uid": "physical-site-uid"}}
	return v
}

func TestSMSv2SitePublisherPhysicalLinks(t *testing.T) {
	configuration := runtimeConfiguration()
	configured, err := resolvedRuntimeFixture(configuration)
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string]smsv2BindingModel{"slo": {Node: types.StringValue("master-0"), Role: types.StringValue("slo"), MAC: types.StringValue("02:aa:bb:cc:dd:01")}}
	for _, tc := range []struct {
		name   string
		change func(client.SMSv2Observation)
		valid  bool
	}{
		{"valid", func(client.SMSv2Observation) {}, true},
		{"foreign publisher", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["creator_id"] = "foreign-site" }, false},
		{"wrong status node", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["status_id"] = "other_SiteStatusMgr" }, false},
		{"stale", func(v client.SMSv2Observation) { physicalFixtureMetadata(v)["vtrp_stale"] = true }, false},
		{"down", func(v client.SMSv2Observation) { physicalFixtureInterface(v)["link_state"] = false }, false},
		{"missing physical UID", func(v client.SMSv2Observation) { delete(v, "system_metadata") }, false},
		{"wrong instance node", func(v client.SMSv2Observation) {
			v["status"].([]interface{})[0].(map[string]interface{})["ver_status"].(map[string]interface{})["ver_instance_name"] = "other-lab-site"
		}, false},
		{"missing reference", func(v client.SMSv2Observation) {
			delete(v["status"].([]interface{})[0].(map[string]interface{}), "object_refs")
		}, false},
		{"foreign reference", func(v client.SMSv2Observation) {
			v["status"].([]interface{})[0].(map[string]interface{})["object_refs"].([]interface{})[0].(map[string]interface{})["uid"] = "foreign-uid"
		}, false},
		{"duplicate reference", func(v client.SMSv2Observation) {
			s := v["status"].([]interface{})[0].(map[string]interface{})
			refs := s["object_refs"].([]interface{})
			s["object_refs"] = append(refs, refs[0])
		}, false},
		{"duplicate publication", func(v client.SMSv2Observation) { s := v["status"].([]interface{}); v["status"] = append(s, s[0]) }, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			v := sitePublisherPhysicalFixture()
			tc.change(v)
			err := validateSMSv2PhysicalLinks(configuration, v, configured, bindings)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v error=%v", tc.valid, err)
			}
		})
	}
}
