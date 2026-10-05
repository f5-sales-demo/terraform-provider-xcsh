---
page_title: "enable_disable_signatures"
subcategory: ""
description: "Enable Disable Signature Choice."
xcsh_docs: {"aliases": ["enable disable signatures"], "body_bytes": 1911, "body_sha256": "sha256:f6a76dcc65fc20122274f9672c1b0fd60d55e6375f057682b55ac1e6305b6f7a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "documentation/data-sources/protocol_inspection/properties/enable_disable_signatures/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220", "registry_path": "docs/guides/data-sources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_signatures"], "schema_version": 1, "sections": [{"aliases": ["enable disable signatures disable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "disable_signature"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable disable signatures enable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "enable_signature"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_signatures/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Enable Disable Signature Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
