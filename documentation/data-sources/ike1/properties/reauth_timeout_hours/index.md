---
page_title: "reauth_timeout_hours"
subcategory: ""
description: "Input Hours."
xcsh_docs: {"aliases": ["duration", "reauth timeout hours"], "body_bytes": 1548, "body_sha256": "sha256:80bd2045ce33f0e386b1066e785d204299f5d6bd46db002e83c80c84f6bd6c2c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_hours", "parent_id": "xcsh-docs:data-sources:ike1:reference", "path": "documentation/data-sources/ike1/properties/reauth_timeout_hours/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0102013100100030-1300321120331131-2201302003323321-1202312123020013-2030320331100303-3113331330013120-3010003022023213-1032000312002110", "registry_path": "docs/guides/data-sources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_timeout_hours"], "schema_version": 1, "sections": [{"aliases": ["duration", "reauth timeout hours duration"], "anchor": "schema-reauth_timeout_hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_hours", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_timeout_hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/properties/reauth_timeout_hours/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Input Hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["ike1CreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_hours

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/)
- reauth_timeout_hours

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Additional upstream details:

Input Hours.

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

<a id="schema-reauth_timeout_hours--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```
