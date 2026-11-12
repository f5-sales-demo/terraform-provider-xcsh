---
page_title: "reauth_timeout_days"
subcategory: ""
description: "Set Duration in days."
xcsh_docs: {"aliases": ["duration", "reauth timeout days"], "body_bytes": 2018, "body_sha256": "sha256:e15351cec8027474fb4c7ef046c0a71726d2192d466cfa0991cf5f4b5ac94bf7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "documentation/resources/ike1/properties/reauth_timeout_days/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0220323021313113-1033010000130003-0220310211031020-1021200012332310-0001102230113202-1001013232032120-1010122331133030-2212010211021003", "registry_path": "docs/guides/resources--ike1--reference--group-001.md", "relationships": [{"anchor": "schema-reauth_timeout_days--duration", "enforcement": "provider-schema", "group": "reauth_timeout_days:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_timeout_days"], "schema_version": 1, "sections": [{"aliases": ["duration", "reauth timeout days duration"], "anchor": "schema-reauth_timeout_days--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_timeout_days", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_timeout_days/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Set Duration in days.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["ike1CreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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
