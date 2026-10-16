---
page_title: "enable_disable_signatures"
subcategory: ""
description: "Enable Disable Signature Choice."
xcsh_docs: {"aliases": ["enable disable signatures"], "body_bytes": 1271, "body_sha256": "sha256:c5a0bb7a6fbcbd82af67de33911a77e26a4a51e424bcf24d5135592f7c689ff2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "documentation/data-sources/protocol_inspection/properties/enable_disable_signatures/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220", "registry_path": "docs/guides/data-sources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_signatures"], "schema_version": 1, "sections": [{"aliases": ["enable disable signatures disable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "disable_signature"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable disable signatures enable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "enable_signature"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_signatures/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Enable Disable Signature Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_signatures

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/)
- enable_disable_signatures

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for enable disable signatures.

Additional upstream details:

Enable Disable Signature Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

## Direct properties

- [disable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/disable_signature/): complete subsection reference.

- [enable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/enable_signature/): complete subsection reference.
