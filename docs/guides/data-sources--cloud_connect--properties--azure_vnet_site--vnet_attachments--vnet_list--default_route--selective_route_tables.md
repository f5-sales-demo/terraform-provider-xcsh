---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 2506, "body_sha256": "sha256:99837e8960fbd6a13799880be0db5f4899a44df0fe2bb39fdd4be8e4830e1841", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "path": "docs/guides/data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "default_route", "selective_route_tables"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/selective_route_tables/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for selective route tables.

Upstream description:

Azure Route Table.

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

<a id="schema-azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables--route_table_id"></a>

### route_table_id property

Type: `["list", "string"]`. Computed.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
