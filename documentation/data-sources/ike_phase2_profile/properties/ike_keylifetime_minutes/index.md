---
page_title: "ike_keylifetime_minutes"
subcategory: ""
description: "Set IKE Key Lifetime in minutes."
xcsh_docs: {"aliases": ["ike keylifetime minutes"], "body_bytes": 1631, "body_sha256": "sha256:bd135df3e7eb664b814a47917ecd714b0d73ab532ec83b1c1b9e15821a875795", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:properties:ike_keylifetime_minutes", "parent_id": "xcsh-docs:data-sources:ike_phase2_profile:reference", "path": "documentation/data-sources/ike_phase2_profile/properties/ike_keylifetime_minutes/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3110023100310321-3211000132212003-1211001313310313-2323132132231223-2131032031310313-1022320312031330-0302321200131310-2133333103210022", "registry_path": "docs/guides/data-sources--ike_phase2_profile--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_minutes"], "schema_version": 1, "sections": [{"aliases": ["ike keylifetime minutes duration"], "anchor": "schema-ike_keylifetime_minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:ike_phase2_profile:properties:ike_keylifetime_minutes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/properties/ike_keylifetime_minutes/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Set IKE Key Lifetime in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_minutes

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/)
- ike_keylifetime_minutes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

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

<a id="schema-ike_keylifetime_minutes--duration"></a>

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
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
