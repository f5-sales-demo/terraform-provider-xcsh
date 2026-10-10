---
page_title: "filter_fields.filter_expression_field"
subcategory: ""
description: "Filter Expression Field."
xcsh_docs: {"aliases": ["filter fields filter expression field"], "body_bytes": 1658, "body_sha256": "sha256:04fb59178da18dc3990c849755a4c3fe298a2e14135a12f3c2d01659058270a4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:filter_expression_field", "parent_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "path": "documentation/data-sources/filter_set/properties/filter_fields/filter_expression_field/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2012103130001230-2322121013103302-0100211001222230-3300002022000120-2102320103311120-2120113212331331-3032232312010321-1112122332010332", "registry_path": "docs/guides/data-sources--filter_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "filter_expression_field"], "schema_version": 1, "sections": [{"aliases": ["filter fields filter expression field expression"], "anchor": "schema-filter_fields--filter_expression_field--expression", "description": "Expression is a Kubernetes style label expression for selections, but differs in that it allows special characters in the keys and values.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:filter_expression_field", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "filter_expression_field", "expression"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/filter_fields/filter_expression_field/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Filter Expression Field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["filter_setCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.filter_expression_field

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/)
- filter_fields.filter_expression_field

<a id="section"></a>

Type: `"single"`. Computed.

Filter Expression Field.

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

<a id="schema-filter_fields--filter_expression_field--expression"></a>

### expression property

Type: `"string"`. Computed.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
