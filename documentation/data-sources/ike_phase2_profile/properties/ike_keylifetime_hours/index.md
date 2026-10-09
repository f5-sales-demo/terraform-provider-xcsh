---
page_title: "ike_keylifetime_hours"
subcategory: ""
description: "Input Hours."
xcsh_docs: {"aliases": ["ike keylifetime hours"], "body_bytes": 2304, "body_sha256": "sha256:07054288fc0ffa112de36d82b8c4c38025f20e9b54689fd510824c34ce628c0e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:properties:ike_keylifetime_hours", "parent_id": "xcsh-docs:data-sources:ike_phase2_profile:reference", "path": "documentation/data-sources/ike_phase2_profile/properties/ike_keylifetime_hours/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0221231222222301-0211130333012203-0003230320030031-1010003201022212-1021323232023230-3102301221130200-0002230220230222-1213032102021021", "registry_path": "docs/guides/data-sources--ike_phase2_profile--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_hours"], "schema_version": 1, "sections": [{"aliases": ["ike keylifetime hours duration"], "anchor": "schema-ike_keylifetime_hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:ike_phase2_profile:properties:ike_keylifetime_hours", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/properties/ike_keylifetime_hours/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Input Hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_hours

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/)
- ike_keylifetime_hours

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/ike_keylifetime_hours/#section)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/ike_keylifetime_minutes/#section)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/use_default_keylifetime/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-ike_keylifetime_hours--duration"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
