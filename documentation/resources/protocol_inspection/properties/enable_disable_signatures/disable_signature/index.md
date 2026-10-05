---
page_title: "enable_disable_signatures.disable_signature"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable disable signatures disable signature"], "body_bytes": 1404, "body_sha256": "sha256:99d1e6dfbc1efb690b8c90b5d98017e115b82fcff4a0836d006c8c01dfe07ee9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "parent_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures", "path": "documentation/resources/protocol_inspection/properties/enable_disable_signatures/disable_signature/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3330133210221230-3222233212003302-1012112101212302-2203233110220020-0211212000001121-3002003303032113-1122010212123113-0113301221233212", "registry_path": "docs/guides/resources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_signatures", "disable_signature"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_signatures/disable_signature/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_signatures.disable_signature

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
- [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/)
- enable_disable_signatures.disable_signature

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable signature.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
disable_signature = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
