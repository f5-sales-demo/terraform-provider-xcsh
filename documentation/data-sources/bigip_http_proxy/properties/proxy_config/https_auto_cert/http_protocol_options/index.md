---
page_title: "proxy_config.https_auto_cert.http_protocol_options"
subcategory: ""
description: "HTTP protocol configuration OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["proxy config https auto cert http protocol options"], "body_bytes": 3012, "body_sha256": "sha256:36e13390487e7994322c5ec8123543bbb504f941833e82f03474f0f38a7adc0e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "http_protocol_options"], "schema_version": 1, "sections": [{"aliases": ["proxy config https auto cert http protocol options http protocol enable v1 only"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https auto cert http protocol options http protocol enable v1 v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_v2"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https auto cert http protocol options http protocol enable v2 only"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v2_only"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "HTTP protocol configuration OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https_auto_cert.http_protocol_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- proxy_config.https_auto_cert.http_protocol_options

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

- [http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/): complete subsection reference.

- [http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_v2/): complete subsection reference.

- [http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v2_only/): complete subsection reference.

## Next pages

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_v2/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v2_only/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
