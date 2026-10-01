---
page_title: "voltstack_cluster.az_nodes.local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster.az_nodes.local_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1971, "body_sha256": "sha256:e4502c32175eb569e3f2e02400f9486b0a705b6ec0746258a61e6e3e5a3a2d04", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.az_nodes.local_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes.local_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes.md)
- voltstack_cluster.az_nodes.local_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
```

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

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [subnet](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet.md): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet.subnet](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet.md)
- [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--properties--voltstack_cluster--az_nodes.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
