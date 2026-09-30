---
page_title: "vip_params_per_az"
subcategory: "Infrastructure"
description: "vip_params_per_az for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1794, "body_sha256": "sha256:68e2d00630b03b94b0f710a236c8c50a163a0d38e85c6349cb9c636579649fe4", "canonical_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "docs/guides/data-sources--site--properties--vip_params_per_az.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vip_params_per_az"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/vip_params_per_az/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vip_params_per_az for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vip_params_per_az

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- vip_params_per_az

<a id="section"></a>

Type: `"list"`. Computed.

Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in
advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined
will be used to publish to external systems like K8s, Consul.

## Direct properties

<a id="schema-vip_params_per_az--az_name"></a>

### az_name property

Type: `"string"`. Computed.

AZ Name. Name of the Availability zone.

<a id="schema-vip_params_per_az--inside_vip"></a>

### inside_vip property

Type: `["list", "string"]`. Computed.

Inside VIP(s). List of Inside VIPs for an AZ.

<a id="schema-vip_params_per_az--inside_vip_cname"></a>

### inside_vip_cname property

Type: `"string"`. Computed.

CNAME value for the inside VIP, These are usually public cloud generated CNAME.

<a id="schema-vip_params_per_az--inside_vip_v6"></a>

### inside_vip_v6 property

Type: `["list", "string"]`. Computed.

Optional list of Inside IPv6 VIPs for an AZ.

<a id="schema-vip_params_per_az--outside_vip"></a>

### outside_vip property

Type: `["list", "string"]`. Computed.

Outside VIP(s). List of Outside VIPs for an AZ.

<a id="schema-vip_params_per_az--outside_vip_cname"></a>

### outside_vip_cname property

Type: `"string"`. Computed.

CNAME value for the outside VIP These are usually public cloud generated CNAME.

<a id="schema-vip_params_per_az--outside_vip_v6"></a>

### outside_vip_v6 property

Type: `["list", "string"]`. Computed.

Optional list of Outside IPv6 VIPs for an AZ.

## Next pages

- [Property reference](data-sources--site--reference.md)
- [xcsh_site](../data-sources/site.md)
