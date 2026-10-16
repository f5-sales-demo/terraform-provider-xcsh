---
page_title: "custom_data_types"
subcategory: "Security"
description: "Select your custom data types to be monitored in the API discovery."
xcsh_docs: {"aliases": ["custom data types"], "body_bytes": 1417, "body_sha256": "sha256:ba9364af8490be7f2cbd1cbaca8d57fffec786d3bdc36ec08eb61c91e2631bf1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types", "parent_id": "xcsh-docs:data-sources:sensitive_data_policy:reference", "path": "documentation/data-sources/sensitive_data_policy/properties/custom_data_types/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2213331322133102-0313231331030120-0210321330120000-3203032221020333-2021113001131100-2132023212032102-1123203103322320-1230012010303112", "registry_path": "docs/guides/data-sources--sensitive_data_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_data_types"], "schema_version": 1, "sections": [{"aliases": ["custom data types custom data type ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:sensitive_data_policy:properties:custom_data_types:custom_data_type_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_data_types", "custom_data_type_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/properties/custom_data_types/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Select your custom data types to be monitored in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_data_types

Breadcrumbs:

- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/)
- custom_data_types

<a id="section"></a>

Type: `"list"`. Computed.

Select your custom data types to be monitored in the API discovery. Defaults to \`\[\]\`. Server
applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [custom_data_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/custom_data_types/custom_data_type_ref/): complete subsection reference.
