---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.custom_routing"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list.custom_routing for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1804, "body_sha256": "sha256:0f249ceaad546306696ef166a185fb0fa9b9797ec8ab10c5be8aef67053be290", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing:route_tables"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "parent_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "path": "docs/guides/resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "custom_routing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list.custom_routing for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list.custom_routing

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [azure_vnet_site](resources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List Azure Route Table with Static Route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
```

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

Terraform syntax:

```terraform
custom_routing {
  # Configure direct properties listed below.
}
```

## Direct properties

- [route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables.md): complete subsection reference.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables.md)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
