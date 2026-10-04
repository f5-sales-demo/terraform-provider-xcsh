---
page_title: "enable_disable_signatures"
subcategory: ""
description: "Enable Disable Signature Choice."
xcsh_docs: {"aliases": ["enable disable signatures"], "body_bytes": 1911, "body_sha256": "sha256:f6a76dcc65fc20122274f9672c1b0fd60d55e6375f057682b55ac1e6305b6f7a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "documentation/data-sources/protocol_inspection/properties/enable_disable_signatures/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220", "registry_path": "docs/guides/data-sources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_signatures"], "schema_version": 1, "sections": [{"aliases": ["enable disable signatures disable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "disable_signature"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable disable signatures enable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "enable_signature"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_signatures/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Enable Disable Signature Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

Upstream description:

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

## Next pages

- [enable_disable_signatures.disable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/disable_signature/)
- [enable_disable_signatures.enable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/enable_signature/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
