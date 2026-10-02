---
page_title: "https_auto_cert.http_protocol_options.http_protocol_enable_v1_only"
subcategory: "Load Balancing"
description: "HTTP/1.1 Protocol OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["https auto cert http protocol options http protocol enable v1 only"], "body_bytes": 1952, "body_sha256": "sha256:21ef39e53eaeae222e2e58bd2b8053e04f4a2b61051de2ba630a676c86f71d04", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:http_protocol_options", "path": "documentation/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "sections": [{"aliases": ["header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https_auto_cert/)
- [https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

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

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/): complete subsection reference.

## Next pages

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
