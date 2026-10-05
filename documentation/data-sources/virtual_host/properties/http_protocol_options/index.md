---
page_title: "http_protocol_options"
subcategory: ""
description: "HTTP protocol configuration OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["http protocol options"], "body_bytes": 2316, "body_sha256": "sha256:6933fc781a8ff30ec0b1338386906d8da5a0d082a36b4339f324d32f59e82cb8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/http_protocol_options/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_protocol_options"], "schema_version": 1, "sections": [{"aliases": ["http protocol options http protocol enable v1 only"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only"], "syntax": "attribute", "type": "object"}, {"aliases": ["http protocol options http protocol enable v1 v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_protocol_options", "http_protocol_enable_v1_v2"], "syntax": "attribute", "type": "object"}, {"aliases": ["http protocol options http protocol enable v2 only"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_protocol_options", "http_protocol_enable_v2_only"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/http_protocol_options/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "HTTP protocol configuration OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- http_protocol_options

<a id="section"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

## Direct properties

- [http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/): complete subsection reference.

- [http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_v2/): complete subsection reference.

- [http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v2_only/): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/)
- [http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_v2/)
- [http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v2_only/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
