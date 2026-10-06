---
page_title: "reauth_timeout_days"
subcategory: ""
description: "Set Duration in days."
xcsh_docs: {"aliases": ["duration", "reauth timeout days"], "body_bytes": 2018, "body_sha256": "sha256:9feb05b8be61c71f935d771ffb595276ef724211ca8d5fe0526d134ef9f7ba54", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "documentation/resources/ike1/properties/reauth_timeout_days/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0220323021313113-1033010000130003-0220310211031020-1021200012332310-0001102230113202-1001013232032120-1010122331133030-2212010211021003", "registry_path": "docs/guides/resources--ike1--reference--group-001.md", "relationships": [{"anchor": "schema-reauth_timeout_days--duration", "enforcement": "provider-schema", "group": "reauth_timeout_days:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_timeout_days"], "schema_version": 1, "sections": [{"aliases": ["duration", "reauth timeout days duration"], "anchor": "schema-reauth_timeout_days--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_timeout_days", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_timeout_days/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Set Duration in days.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["ike1CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_days

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/)
- reauth_timeout_days

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Additional upstream details:

Set Duration in days.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

Terraform syntax:

```terraform
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-reauth_timeout_days--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```
