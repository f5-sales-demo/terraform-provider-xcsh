---
page_title: "ike_keylifetime_hours"
subcategory: ""
description: "Input Hours."
xcsh_docs: {"aliases": ["ike keylifetime hours"], "body_bytes": 2758, "body_sha256": "sha256:6ca9afd8a49730b31f705bf7c9946ab1edc3a35b233f871b0ebfe47e7700b6ca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "documentation/resources/ike_phase1_profile/properties/ike_keylifetime_hours/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0202020210020100-0010323001013223-3102230231101231-0332020330300301-3131031101323222-2122211233201101-2032303103300212-2333331323213101", "registry_path": "docs/guides/resources--ike_phase1_profile--reference--group-001.md", "relationships": [{"anchor": "schema-ike_keylifetime_hours--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_hours"], "schema_version": 1, "sections": [{"aliases": ["ike keylifetime hours duration"], "anchor": "schema-ike_keylifetime_hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/ike_keylifetime_hours/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Input Hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_hours

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- ike_keylifetime_hours

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Additional upstream details:

Input Hours.

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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/ike_keylifetime_hours/#section)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/ike_keylifetime_minutes/#section)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/use_default_keylifetime/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ike_keylifetime_hours--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
