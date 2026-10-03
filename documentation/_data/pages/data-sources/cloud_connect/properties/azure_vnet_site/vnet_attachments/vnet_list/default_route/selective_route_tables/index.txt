---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables"
subcategory: ""
description: "Azure Route Table."
xcsh_docs: {"aliases": ["azure vnet site vnet attachments vnet list default route selective route tables"], "body_bytes": 3006, "body_sha256": "sha256:5a11440cc560c3cebfda90889bd6de8a5db476a7fc8dc4d021d44c190a2f5ed7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route", "path": "documentation/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/selective_route_tables/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0020001020221203-3323031110111102-3223132120232113-3031230333001003-2022221320110213-0220331100120123-3332110302231002-1322333311212331", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "default_route", "selective_route_tables"], "schema_version": 1, "sections": [{"aliases": ["azure vnet site vnet attachments vnet list default route selective route tables route table id"], "anchor": "schema-azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables--route_table_id", "description": "Route table ID in the format /<resource-group-name>/<route-table-name>", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:default_route:selective_route_tables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "default_route", "selective_route_tables", "route_table_id"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/selective_route_tables/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Azure Route Table.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/)
- [azure_vnet_site.vnet_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/)
- [azure_vnet_site.vnet_attachments.vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/)
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

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
