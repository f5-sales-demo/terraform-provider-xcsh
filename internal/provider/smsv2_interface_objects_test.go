// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"strings"
	"testing"
)

func runtimeInterfaceObjects(configuration client.SMSv2Observation) client.SMSv2Observation {
	configured, err := extractSMSv2ConfiguredInterfaces(configuration)
	if err != nil {
		panic(err)
	}
	items := []interface{}{}
	for _, iface := range configured {
		role := "site_local_network"
		if iface.Role == "sli" {
			role = "site_local_inside_network"
		}
		items = append(items, map[string]interface{}{
			"name":       fmt.Sprintf("ves-io-securemesh-site-v2-lab-site-network-%s-%s-0", iface.Node, iface.Device),
			"namespace":  "system",
			"owner_view": map[string]interface{}{"kind": "securemesh_site_v2", "name": "lab-site", "namespace": "system", "uid": "site-uid"},
			"get_spec":   map[string]interface{}{"ethernet_interface": map[string]interface{}{"node": iface.Node, "device": iface.Device, "mtu": float64(iface.MTU), role: map[string]interface{}{}}},
		})
	}
	return client.SMSv2Observation{"items": items, "errors": []interface{}{}}
}

func resolvedRuntimeFixture(configuration client.SMSv2Observation) ([]smsv2ConfiguredInterface, error) {
	configured, err := extractSMSv2ConfiguredInterfaces(configuration)
	if err != nil {
		return nil, err
	}
	return resolveSMSv2InterfaceObjects(configuration, configured, runtimeInterfaceObjects(configuration))
}

func TestSMSv2ResolvesActualInterfaceObjectNames(t *testing.T) {
	config := runtimeConfiguration()
	configured, err := extractSMSv2ConfiguredInterfaces(config)
	if err != nil {
		t.Fatal(err)
	}
	objects := runtimeInterfaceObjects(config)
	objects["items"].([]interface{})[0].(map[string]interface{})["name"] = "actual-platform-interface"
	got, err := resolveSMSv2InterfaceObjects(config, configured, objects)
	if err != nil || got[0].Name != "actual-platform-interface" {
		t.Fatalf("got=%v err=%v", got, err)
	}
	if configured[0].Name != "" {
		t.Fatal("resolver mutated its input")
	}
}

func TestSMSv2InterfaceDiscoveryRejectsUnprovenIdentity(t *testing.T) {
	for _, mutation := range []string{"missing", "duplicate", "foreign UID", "foreign site", "foreign namespace", "wrong device", "wrong node", "wrong role", "wrong MTU", "partial errors", "missing items"} {
		t.Run(mutation, func(t *testing.T) {
			config := runtimeConfiguration()
			configured, err := extractSMSv2ConfiguredInterfaces(config)
			if err != nil {
				t.Fatal(err)
			}
			objects := runtimeInterfaceObjects(config)
			items := objects["items"].([]interface{})
			item := items[0].(map[string]interface{})
			owner := item["owner_view"].(map[string]interface{})
			eth := item["get_spec"].(map[string]interface{})["ethernet_interface"].(map[string]interface{})
			switch mutation {
			case "missing":
				objects["items"] = items[1:]
			case "duplicate":
				objects["items"] = append(items, item)
			case "foreign UID":
				owner["uid"] = "old-site-uid"
			case "foreign site":
				owner["name"] = "other-site"
			case "foreign namespace":
				owner["namespace"] = "other"
			case "wrong device":
				eth["device"] = "eth9"
			case "wrong node":
				eth["node"] = "other-node"
			case "wrong role":
				eth["site_local_inside_network"] = map[string]interface{}{}
			case "wrong MTU":
				eth["mtu"] = float64(9000)
			case "partial errors":
				objects["errors"] = []interface{}{map[string]interface{}{"message": "private server diagnostic"}}
			case "missing items":
				delete(objects, "items")
			}
			_, err = resolveSMSv2InterfaceObjects(config, configured, objects)
			if err == nil {
				t.Fatal("accepted unproven identity")
			}
			if strings.Contains(err.Error(), "private server diagnostic") {
				t.Fatal("leaked server diagnostics")
			}
		})
	}
}
