---
page_title: "voltstack_cluster.active_enhanced_firewall_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster.active_enhanced_firewall_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1640, "body_sha256": "sha256:815a9271356c9450e49a4b412128bbcfc381c424e194cc45cdacbe3599017f38", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.active_enhanced_firewall_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
