---
page_title: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route"
subcategory: "Infrastructure"
description: "Defines a static route, configuring a list of prefixes and a next-hop to be used for them."
xcsh_docs: {"aliases": ["voltstack cluster ar outside static routes static route list custom static route"], "body_bytes": 4578, "body_sha256": "sha256:d1088171e9c9055773b7f0571db2e8b74f7766f5847a460fe18d6d098ef6e6c9", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:labels", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2101021223322030-1330023313301112-1332300232030121-2231231213100230-1231021020020133-1002213132311212-0020102103203121-0123202031213111", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route"], "schema_version": 1, "sections": [{"aliases": ["attrs"], "anchor": "schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--attrs", "description": "List of route attributes associated with the static route.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "attrs"], "syntax": "attribute", "type": "list"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this Static Route, these labels can be used in network policy.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["nexthop"], "anchor": "section", "description": "Identifies the next-hop for a route.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnets"], "anchor": "section", "description": "List of route prefixes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Defines a static route, configuring a list of prefixes and a next-hop to be used for them.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route

<a id="section"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

<a id="schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--attrs"></a>

### attrs property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/labels/): complete subsection reference.

- [nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/): complete subsection reference.

- [subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/labels/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
