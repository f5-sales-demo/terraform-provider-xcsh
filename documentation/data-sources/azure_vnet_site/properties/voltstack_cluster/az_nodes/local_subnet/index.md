---
page_title: "voltstack_cluster.az_nodes.local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster.az_nodes.local_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2111, "body_sha256": "sha256:f1809419f4ed78017f9598c53b23c02dd556f288599794c13efc0f0516fdfb41", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.az_nodes.local_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
