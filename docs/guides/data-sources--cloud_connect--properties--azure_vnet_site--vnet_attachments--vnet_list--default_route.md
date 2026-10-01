---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.default_route"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list.default_route for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 2130, "body_sha256": "sha256:b3f473e48f5fc103a263941f27faf2e1a5866bb06556e25c7cc23f5fb55811de", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:all_route_tables", "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables"], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "path": "docs/guides/data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "default_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list.default_route for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list.default_route

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

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

## Direct properties

- [all_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--all_route_tables.md): complete subsection reference.

- [selective_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md): complete subsection reference.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--all_route_tables.md)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
