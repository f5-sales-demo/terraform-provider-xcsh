---
page_title: "http_protocol_options.http_protocol_enable_v1_only"
subcategory: ""
description: "HTTP/1.1 Protocol OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["http protocol options http protocol enable v1 only"], "body_bytes": 1627, "body_sha256": "sha256:d9981175daeaaafae8d56ce0757f8d0e90c038c876872ac7f5418e3040fedff5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options", "path": "documentation/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "sections": [{"aliases": ["http protocol options http protocol enable v1 only header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/)
- http_protocol_options.http_protocol_enable_v1_only

<a id="section"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/http_protocol_options/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
