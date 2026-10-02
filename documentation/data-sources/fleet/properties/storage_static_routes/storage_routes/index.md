---
page_title: "storage_static_routes.storage_routes"
subcategory: ""
description: "List of storage static routes."
xcsh_docs: {"aliases": ["storage static routes storage routes"], "body_bytes": 4261, "body_sha256": "sha256:c7ed136bbc1a48d809c26e8295de1c59d85d3512129d81dab31f637071a46fb2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:labels", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes", "path": "documentation/data-sources/fleet/properties/storage_static_routes/storage_routes/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110", "registry_path": "docs/guides/data-sources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes"], "schema_version": 1, "sections": [{"aliases": ["attrs"], "anchor": "schema-storage_static_routes--storage_routes--attrs", "description": "List of route attributes associated with the static route.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "attrs"], "syntax": "attribute", "type": "list"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this Static Route, these labels can be used in network policy.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["nexthop"], "anchor": "section", "description": "Identifies the next-hop for a route.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnets"], "anchor": "section", "description": "List of route prefixes.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "subnets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of storage static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/labels/): complete subsection reference.

- [nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/): complete subsection reference.

- [subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/labels/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [storage_static_routes.storage_routes.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
