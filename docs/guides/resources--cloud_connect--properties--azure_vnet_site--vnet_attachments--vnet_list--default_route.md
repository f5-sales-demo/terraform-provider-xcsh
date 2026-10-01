---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.default_route"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list.default_route for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 2407, "body_sha256": "sha256:6f3c84d199d8d5e6b5eb58471fbd83d3d10165aefe7fe54aca773d5abd708066", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:all_route_tables", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "parent_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "path": "docs/guides/resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "default_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list.default_route for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list.default_route

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [azure_vnet_site](resources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
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
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--all_route_tables.md): complete subsection reference.

- [selective_route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md): complete subsection reference.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--all_route_tables.md)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
