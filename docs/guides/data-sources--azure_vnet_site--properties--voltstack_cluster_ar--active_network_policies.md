---
page_title: "voltstack_cluster_ar.active_network_policies"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.active_network_policies for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1299, "body_sha256": "sha256:dd5e58750736e30cd7a52bf4fdd427e36904508395b1cb81b5ec39cc4a97b364", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies:network_policies"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.active_network_policies for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- voltstack_cluster_ar.active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

- [network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
