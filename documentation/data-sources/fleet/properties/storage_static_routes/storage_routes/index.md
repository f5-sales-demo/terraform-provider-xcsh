---
page_title: "storage_static_routes.storage_routes"
subcategory: ""
description: "List of storage static routes."
xcsh_docs: {"aliases": ["storage static routes storage routes"], "body_bytes": 4261, "body_sha256": "sha256:a139f99df3957e205834718f3944dc865c454a4f31e61c7641ced12b5c311a29", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:labels", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes", "path": "documentation/data-sources/fleet/properties/storage_static_routes/storage_routes/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110", "registry_path": "docs/guides/data-sources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes"], "schema_version": 1, "sections": [{"aliases": ["storage static routes storage routes attrs"], "anchor": "schema-storage_static_routes--storage_routes--attrs", "description": "List of route attributes associated with the static route.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "attrs"], "syntax": "attribute", "type": "list"}, {"aliases": ["storage static routes storage routes labels"], "anchor": "section", "description": "Add Labels for this Static Route, these labels can be used in network policy.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage static routes storage routes nexthop"], "anchor": "section", "description": "Identifies the next-hop for a route.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage static routes storage routes subnets"], "anchor": "section", "description": "List of route prefixes.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "subnets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of storage static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/)
- storage_static_routes.storage_routes

<a id="section"></a>

Type: `"list"`. Computed.

List of Static Routes. List of storage static routes.

Upstream description:

List of storage static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-storage_static_routes--storage_routes--attrs"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/labels/): complete subsection reference.

- [nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/): complete subsection reference.

- [subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/labels/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [storage_static_routes.storage_routes.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
