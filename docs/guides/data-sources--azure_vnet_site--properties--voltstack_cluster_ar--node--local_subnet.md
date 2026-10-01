---
page_title: "voltstack_cluster_ar.node.local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.node.local_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1706, "body_sha256": "sha256:ead4307ec58700d78a53009c7eb5b3afb56c4a107d762f082c9bb3f524e92242", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "node", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.node.local_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.node.local_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md)
- voltstack_cluster_ar.node.local_subnet

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

- [subnet](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md)
- [voltstack_cluster_ar.node.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet_param.md)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
