---
page_title: "default_pool.advanced_options.http1_config"
subcategory: "Load Balancing"
description: "HTTP/1.1 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["default pool advanced options http1 config"], "body_bytes": 1792, "body_sha256": "sha256:a70cd186096e22802c68dbb3cc69bf578a7d1d291f06eb3079fa99b6c08669b8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config"], "schema_version": 1, "sections": [{"aliases": ["default pool advanced options http1 config header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
- default_pool.advanced_options.http1_config

<a id="section"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
