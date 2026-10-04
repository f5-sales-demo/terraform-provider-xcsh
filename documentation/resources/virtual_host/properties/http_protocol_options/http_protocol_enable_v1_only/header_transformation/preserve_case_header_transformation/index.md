---
page_title: "http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation"
subcategory: ""
description: "Preserve HTTP header-name case when upstream case must remain unchanged."
xcsh_docs: {"aliases": ["http protocol options http protocol enable v1 only header transformation preserve case header transformation"], "body_bytes": 1973, "body_sha256": "sha256:8bea8cf7da14b6416f42f59816c8859c5205db8caa98ca51ddefd6f277b897c5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "path": "documentation/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1103321023013312-0322220013332302-2311321011032331-1223322102031303-3123213113230031-0311331333022222-1311032212300232-0003200002333323", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "preserve_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Preserve HTTP header-name case when upstream case must remain unchanged.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/)
- [http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
