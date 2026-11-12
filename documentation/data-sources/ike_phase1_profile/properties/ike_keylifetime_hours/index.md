---
page_title: "ike_keylifetime_hours"
subcategory: ""
description: "Input Hours."
xcsh_docs: {"aliases": ["ike keylifetime hours"], "body_bytes": 2304, "body_sha256": "sha256:589ec9691ef18f2004065b9ceb9af7aad83c7601e7205ec4edad8fb860de0985", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:properties:ike_keylifetime_hours", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:reference", "path": "documentation/data-sources/ike_phase1_profile/properties/ike_keylifetime_hours/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0123031303212103-3332123222221131-2132130102220331-1232032333220210-0103020112011230-2202330010130302-3233013302103330-2202121013020130", "registry_path": "docs/guides/data-sources--ike_phase1_profile--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_hours"], "schema_version": 1, "sections": [{"aliases": ["ike keylifetime hours duration"], "anchor": "schema-ike_keylifetime_hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:ike_phase1_profile:properties:ike_keylifetime_hours", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/properties/ike_keylifetime_hours/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Input Hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_hours

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/)
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

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/ike_keylifetime_hours/#section)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/ike_keylifetime_minutes/#section)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/use_default_keylifetime/#section)

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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```
