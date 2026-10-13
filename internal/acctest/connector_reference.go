// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

type connectorEndpoint struct {
	Site      string `json:"site"`
	Node      string `json:"node"`
	Interface string `json:"interface"`
}

func discoverConnectorEndpoint(ctx context.Context, c *client.Client, ref readOnlySiteReference) (connectorEndpoint, error) {
	path := "/api/config/namespaces/" + ref.Namespace + "/sites/" + ref.Name
	metadata, err := readOwnershipProjection(ctx, c, path, ref.Namespace, ref.Name)
	if err != nil {
		return connectorEndpoint{}, err
	}
	var site struct {
		SystemMetadata struct {
			Owner struct {
				Kind      string `json:"kind"`
				UID       string `json:"uid"`
				Namespace string `json:"namespace"`
				Name      string `json:"name"`
			} `json:"owner_view"`
		} `json:"system_metadata"`
		Spec struct {
			Type  string `json:"site_type"`
			Nodes []struct {
				Name string `json:"name"`
			} `json:"main_nodes"`
		} `json:"spec"`
	}
	if err = c.Get(ctx, path, &site); err != nil {
		return connectorEndpoint{}, err
	}
	if site.Spec.Type != "CUSTOMER_EDGE" || len(site.Spec.Nodes) == 0 || site.Spec.Nodes[0].Name == "" {
		return connectorEndpoint{}, errEligibleInfrastructureAbsent
	}
	ownerUID := metadata.UID
	owner := site.SystemMetadata.Owner
	if owner.UID != "" {
		collections := map[string]string{"aws_vpc_site": "aws_vpc_sites", "azure_vnet_site": "azure_vnet_sites", "gcp_vpc_site": "gcp_vpc_sites", "securemesh_site_v2": "securemesh_site_v2s"}
		collection, ok := collections[owner.Kind]
		if !ok || owner.Name == "" || owner.Namespace != "system" {
			return connectorEndpoint{}, errors.New("connector site backing owner unverified")
		}
		observed, readErr := readOwnershipProjection(ctx, c, "/api/config/namespaces/system/"+collection+"/"+owner.Name, "system", owner.Name)
		if readErr != nil || observed.UID != owner.UID {
			return connectorEndpoint{}, errors.New("connector site backing owner readback failed")
		}
		ownerUID = owner.UID
	}
	var inventory struct {
		Items []struct {
			Name  string `json:"name"`
			Owner struct {
				UID string `json:"uid"`
			} `json:"owner_view"`
			Spec struct {
				Type     string `json:"type"`
				Device   string `json:"device_name"`
				Networks []struct {
					Kind      string `json:"kind"`
					Name      string `json:"name"`
					Namespace string `json:"namespace"`
				} `json:"virtual_network"`
			} `json:"get_spec"`
		} `json:"items"`
		Errors []json.RawMessage `json:"errors"`
	}
	if err = c.Get(ctx, "/api/config/namespaces/system/network_interfaces?report_fields=get_spec&report_fields=system_metadata", &inventory); err != nil {
		return connectorEndpoint{}, err
	}
	if inventory.Items == nil || len(inventory.Errors) != 0 {
		return connectorEndpoint{}, errors.New("connector interface inventory malformed")
	}
	devices := []string{}
	for _, item := range inventory.Items {
		if item.Owner.UID == ownerUID && item.Spec.Type == "NETWORK_INTERFACE_ETHERNET" && item.Spec.Device != "" && item.Name != "" {
			inside := false
			for _, network := range item.Spec.Networks {
				if network.Kind != "virtual_network" || network.Name == "" || network.Namespace != "system" {
					continue
				}
				var observed struct {
					Spec struct {
						Inside map[string]interface{} `json:"site_local_inside_network"`
					} `json:"spec"`
				}
				if err = c.Get(ctx, "/api/config/namespaces/system/virtual_networks/"+network.Name, &observed); err != nil {
					return connectorEndpoint{}, err
				}
				if observed.Spec.Inside != nil {
					inside = true
				}
			}
			if inside {
				devices = append(devices, item.Name)
			}
		}
	}
	if len(devices) == 0 {
		return connectorEndpoint{}, errEligibleInfrastructureAbsent
	}
	sort.Strings(devices)
	return connectorEndpoint{Site: ref.Name, Node: site.Spec.Nodes[0].Name, Interface: devices[0]}, nil
}
func RequireConnectorEndpoint(t *testing.T) {
	t.Helper()
	SkipIfNotAccTest(t)
	if IsMockMode() {
		return
	}
	c, err := GetTestClient()
	if err != nil {
		t.Fatal("connector client unavailable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	refs, err := discoverCustomerEdgeSites(ctx, c)
	if err != nil {
		if errors.Is(err, errEligibleInfrastructureAbsent) {
			t.Log("[PREREQUISITE:infrastructure_absent method=GET type=connector_endpoint]")
			t.Skip("typed inventory contains no customer edge")
		}
		t.Fatal("connector site inventory unverified")
	}
	var endpoint connectorEndpoint
	found := false
	for _, ref := range refs {
		endpoint, err = discoverConnectorEndpoint(ctx, c, ref)
		if errors.Is(err, errEligibleInfrastructureAbsent) {
			continue
		}
		if err != nil {
			t.Fatal("connector node and interface prerequisite unverified")
		}
		if err = saveSiteReference(ref); err != nil {
			t.Fatal("connector site receipt unavailable")
		}
		found = true
		break
	}
	if !found {
		t.Log("[PREREQUISITE:infrastructure_absent method=GET type=connector_endpoint]")
		t.Skip("typed customer-edge inventories contain no owned inside interface")
	}

	data, err := json.Marshal(endpoint)
	if err != nil {
		t.Fatal("connector endpoint encoding failed")
	}
	if os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "connector-endpoint.json"), data, 0600) != nil {
		t.Fatal("connector endpoint receipt unavailable")
	}
	t.Log("[PREREQUISITE:tested method=GET type=connector_endpoint]")
}
func ReadOnlyConnectorNode() string      { return connectorEndpointField(true) }
func ReadOnlyConnectorInterface() string { return connectorEndpointField(false) }
func connectorEndpointField(node bool) string {
	var ref connectorEndpoint
	data, err := os.ReadFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "connector-endpoint.json"))
	if err == nil && json.Unmarshal(data, &ref) == nil {
		if node {
			return ref.Node
		}
		return ref.Interface
	}
	if node {
		return "fixture-node"
	}
	return "eth0"
}
