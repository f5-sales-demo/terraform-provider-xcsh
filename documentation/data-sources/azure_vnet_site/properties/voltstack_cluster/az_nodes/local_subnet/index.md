---
page_title: "voltstack_cluster.az_nodes.local_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["voltstack cluster az nodes local subnet"], "body_bytes": 2210, "body_sha256": "sha256:4c5de74e9ddfdeacc69541fdc2834d216985d28ee17119d728fa9fb1206816e4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2110330231022032-0122022333213303-2332230220332101-0130330112332330-3002323021110113-1210132123202102-0120301223023230-1012002301102110", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes.local_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/)
- voltstack_cluster.az_nodes.local_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

## Direct properties

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/)
- [voltstack_cluster.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
