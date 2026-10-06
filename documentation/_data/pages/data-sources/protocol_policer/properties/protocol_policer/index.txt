---
page_title: "protocol_policer"
subcategory: ""
description: "List of L4 protocol match condition and associated traffic rate limits."
xcsh_docs: {"aliases": ["protocol policer"], "body_bytes": 1626, "body_sha256": "sha256:20595ccca74993ee4a04ecf6786fd9bd664bf1e8ae5a583109db7ebe0090e74f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:policer", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer", "parent_id": "xcsh-docs:data-sources:protocol_policer:reference", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer"], "schema_version": 1, "sections": [{"aliases": ["protocol policer policer"], "anchor": "section", "description": "Reference to policer object to apply traffic rate limits.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:policer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["protocol_policer", "policer"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol policer protocol"], "anchor": "section", "description": "Protocol and protocol specific flags to be matched in packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of L4 protocol match condition and associated traffic rate limits.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- protocol_policer

<a id="section"></a>

Type: `"list"`. Computed.

List of L4 protocol match condition and associated traffic rate limits.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/policer/): complete subsection reference.

- [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/): complete subsection reference.
