---
page_title: "graphql_rules.graphql_settings"
subcategory: "Load Balancing"
description: "GraphQL configuration."
xcsh_docs: {"aliases": ["graphql rules graphql settings"], "body_bytes": 4390, "body_sha256": "sha256:718e1ccb608da59493e002915d558aa4486b116ff44749f6cae6edbf072abc71", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules", "path": "documentation/data-sources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0133031100210221-1323313022121310-3033102310033213-0203322113001232-2310323131112130-0100031120121320-3200031210201120-0002020120000011", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "graphql_settings"], "schema_version": 1, "sections": [{"aliases": ["graphql rules graphql settings disable introspection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "disable_introspection"], "syntax": "attribute", "type": "object"}, {"aliases": ["graphql rules graphql settings enable introspection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "enable_introspection"], "syntax": "attribute", "type": "object"}, {"aliases": ["graphql rules graphql settings max batched queries"], "anchor": "schema-graphql_rules--graphql_settings--max_batched_queries", "description": "Specify maximum number of queries in a single batched request.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_batched_queries"], "syntax": "attribute", "type": "number"}, {"aliases": ["graphql rules graphql settings max depth"], "anchor": "schema-graphql_rules--graphql_settings--max_depth", "description": "Specify maximum depth for the GraphQL query.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_depth"], "syntax": "attribute", "type": "number"}, {"aliases": ["graphql rules graphql settings max total length"], "anchor": "schema-graphql_rules--graphql_settings--max_total_length", "description": "Specify maximum length in bytes for the GraphQL query.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_total_length"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "GraphQL configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.graphql_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/graphql_rules/)
- graphql_rules.graphql_settings

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for graphql settings.

Additional upstream details:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

## Direct properties

- [disable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/): complete subsection reference.

- [enable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/enable_introspection/): complete subsection reference.

<a id="schema-graphql_rules--graphql_settings--max_batched_queries"></a>

### max_batched_queries property

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="schema-graphql_rules--graphql_settings--max_depth"></a>

### max_depth property

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="schema-graphql_rules--graphql_settings--max_total_length"></a>

### max_total_length property

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```
