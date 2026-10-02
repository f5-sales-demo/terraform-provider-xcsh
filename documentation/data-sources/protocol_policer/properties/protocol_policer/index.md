---
page_title: "protocol_policer"
subcategory: ""
description: "List of L4 protocol match condition and associated traffic rate limits."
xcsh_docs: {"aliases": ["protocol policer"], "body_bytes": 2186, "body_sha256": "sha256:861dfb3196ab047e8518ac9a09c7b050970f387bf38fa084fb5ba0d9931db1e3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:policer", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer", "parent_id": "xcsh-docs:data-sources:protocol_policer:reference", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer"], "schema_version": 1, "sections": [{"aliases": ["policer"], "anchor": "section", "description": "Reference to policer object to apply traffic rate limits.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:policer", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["protocol_policer", "policer"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol"], "anchor": "section", "description": "Protocol and protocol specific flags to be matched in packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of L4 protocol match condition and associated traffic rate limits.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [protocol_policer.policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/policer/)
- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
