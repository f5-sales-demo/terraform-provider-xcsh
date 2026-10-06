---
page_title: "ike_keylifetime_minutes"
subcategory: ""
description: "Set IKE Key Lifetime in minutes."
xcsh_docs: {"aliases": ["ike keylifetime minutes"], "body_bytes": 2099, "body_sha256": "sha256:ff3507d7981eb50407261d69c66b5ee460bcb64f9c7708dce6806e515ab831e4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_minutes", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "documentation/resources/ike_phase1_profile/properties/ike_keylifetime_minutes/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1233332313021110-1030011230320011-0013110302032120-0111103012113333-0302233032003013-2121303331212323-0211011132210010-2131100332012121", "registry_path": "docs/guides/resources--ike_phase1_profile--reference--group-001.md", "relationships": [{"anchor": "schema-ike_keylifetime_minutes--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_minutes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_minutes"], "schema_version": 1, "sections": [{"aliases": ["ike keylifetime minutes duration"], "anchor": "schema-ike_keylifetime_minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_minutes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/ike_keylifetime_minutes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Set IKE Key Lifetime in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_minutes

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- ike_keylifetime_minutes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ike_keylifetime_minutes--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```
