// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
	"fmt"
	"net/url"
)

// SMSv2Observation preserves an operational response for strict correlation.
type SMSv2Observation map[string]interface{}

func escapeSMSv2Path(value string) string { return url.PathEscape(value) }

func (c *Client) GetSMSv2Configuration(ctx context.Context, namespace, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_site_v2s/%s", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2Health(ctx context.Context, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/operate/namespaces/system/sites/%s/vpm/debug/global/health", escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2BGPPeers(ctx context.Context, namespace, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/operate/namespaces/%s/sites/%s/ver/bgp_peers", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2BGPRoutes(ctx context.Context, namespace, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/operate/namespaces/%s/sites/%s/ver/bgp_routes", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2SimplifiedRoutes(ctx context.Context, namespace, site, role string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/operate/namespaces/%s/sites/%s/ver/simplified_routes", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	body := map[string]interface{}{"namespace": namespace, "site": site, "all_nodes": map[string]interface{}{}}
	switch role {
	case "slo", "sli":
		body[role] = map[string]interface{}{}
	default:
		return nil, fmt.Errorf("unsupported SMSv2 route role %q", role)
	}
	err := c.Post(ctx, path, body, &result)
	return result, err
}

func (c *Client) GetSMSv2SiteUpgradeStatus(ctx context.Context, namespace, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/config/namespaces/%s/sites/%s", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2UpgradableSoftwareVersions(ctx context.Context, currentOS, currentSoftware string) (SMSv2Observation, error) {
	var result SMSv2Observation
	query := url.Values{}
	query.Set("current_os_version", currentOS)
	query.Set("current_sw_version", currentSoftware)
	err := c.Get(ctx, "/api/maurice/upgradable_sw_versions?"+query.Encode(), &result)
	return result, err
}

func (c *Client) GetSMSv2PreUpgradeCheck(ctx context.Context, namespace, site, softwareVersion string) (SMSv2Observation, error) {
	var result SMSv2Observation
	query := url.Values{}
	query.Set("sw_version", softwareVersion)
	path := fmt.Sprintf("/api/maurice/namespaces/%s/sites/%s/pre_upgrade_check?%s", escapeSMSv2Path(namespace), escapeSMSv2Path(site), query.Encode())
	err := c.Get(ctx, path, &result)
	return result, err
}

func (c *Client) GetSMSv2UpgradeProgress(ctx context.Context, namespace, site string) (SMSv2Observation, error) {
	var result SMSv2Observation
	path := fmt.Sprintf("/api/maurice/namespaces/%s/sites/%s/upgrade_status", escapeSMSv2Path(namespace), escapeSMSv2Path(site))
	err := c.Get(ctx, path, &result)
	return result, err
}
