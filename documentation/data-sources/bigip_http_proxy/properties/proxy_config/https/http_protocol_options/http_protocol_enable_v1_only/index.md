---
page_title: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only"
subcategory: ""
description: "HTTP/1.1 Protocol OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["proxy config https http protocol options http protocol enable v1 only"], "body_bytes": 2104, "body_sha256": "sha256:04ec066f1d3f430fef6df85e0c9ae27e1a7cb8b7efc7d3757dbd8993267b6c90", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "sections": [{"aliases": ["header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

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

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/): complete subsection reference.

## Next pages

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [proxy_config.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
