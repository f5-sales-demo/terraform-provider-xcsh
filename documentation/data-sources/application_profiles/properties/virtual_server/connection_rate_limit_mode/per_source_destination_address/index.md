---
page_title: "virtual_server.connection_rate_limit_mode.per_source_destination_address"
subcategory: ""
description: "Destination and Source Address Mask."
xcsh_docs: {"aliases": ["virtual server connection rate limit mode per source destination address"], "body_bytes": 2973, "body_sha256": "sha256:8bd16dc9878eecb192c24325ce4dade5b41c9a3b4a84d42cabe4a90258ae336f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "path": "documentation/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3330032322333222-0102301301210031-1303211030222203-1011111011221201-0201131101002113-0322122120130331-3322302200001000-1031202013332002", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_destination_address"], "schema_version": 1, "sections": [{"aliases": ["virtual server connection rate limit mode per source destination address destination mask"], "anchor": "schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--destination_mask", "description": "Configuration parameter for destination mask", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_destination_address", "destination_mask"], "syntax": "attribute", "type": "number"}, {"aliases": ["virtual server connection rate limit mode per source destination address source mask"], "anchor": "schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--source_mask", "description": "Configuration parameter for source mask", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_destination_address", "source_mask"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Destination and Source Address Mask.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["application_profilesCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode.per_source_destination_address

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/)
- virtual_server.connection_rate_limit_mode.per_source_destination_address

<a id="section"></a>

Type: `"single"`. Computed.

Destination and Source Address Mask.

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

<a id="schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--destination_mask"></a>

### destination_mask property

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--source_mask"></a>

### source_mask property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```
