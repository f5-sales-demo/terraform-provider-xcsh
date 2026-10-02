---
page_title: "voltstack_cluster.az_nodes.local_subnet.subnet_param"
subcategory: "Infrastructure"
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["voltstack cluster az nodes local subnet subnet param"], "body_bytes": 2610, "body_sha256": "sha256:426b5563af6fe588f7ef2be30d7e28dc2c561f6d8d45dfd84784aa6f67950844", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0311000113233111-2012211331232121-2000313102223232-2033210112120033-2213012021210020-3010102320001023-1233210130223120-3130123330323211", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes.local_subnet.subnet_param

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/)
- [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/)
- voltstack_cluster.az_nodes.local_subnet.subnet_param

<a id="section"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

## Next pages

- [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
