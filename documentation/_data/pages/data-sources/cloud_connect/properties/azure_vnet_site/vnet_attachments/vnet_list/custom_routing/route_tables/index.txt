---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables"
subcategory: ""
description: "Route Tables with static routes."
xcsh_docs: {"aliases": ["azure vnet site vnet attachments vnet list custom routing route tables"], "body_bytes": 5448, "body_sha256": "sha256:d6751db1095eb637c4d8d17938e5893f02d0cd314de1e3774f030d9ddcee57d9", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing:route_tables", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing", "path": "documentation/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/route_tables/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2232200020030213-2233122101320313-0020210103330312-0131131212000202-2313322032131132-0101021311112010-0331323202330030-3023202020332330", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "custom_routing", "route_tables"], "schema_version": 1, "sections": [{"aliases": ["route table id"], "anchor": "schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--route_table_id", "description": "Route table ID in the format /<resource-group-name>/<route-table-name>", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing:route_tables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "custom_routing", "route_tables", "route_table_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["static routes"], "anchor": "schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--static_routes", "description": "List of Static Routes.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:custom_routing:route_tables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "custom_routing", "route_tables", "static_routes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/route_tables/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Route Tables with static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/)
- [azure_vnet_site.vnet_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/)
- [azure_vnet_site.vnet_attachments.vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/)
- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

<a id="section"></a>

Type: `"list"`. Computed.

List of route tables with static routes. Route Tables with static routes.

Upstream description:

Route Tables with static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--route_table_id"></a>

### route_table_id property

Type: `"string"`. Computed.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  }
}
```

<a id="schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--static_routes"></a>

### static_routes property

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
