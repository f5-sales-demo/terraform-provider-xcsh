---
page_title: "http_proxy"
subcategory: ""
description: "Parameters for HTTP Connect Proxy."
xcsh_docs: {"aliases": ["http proxy"], "body_bytes": 1583, "body_sha256": "sha256:255b87998be3aa0e7e5a4cd4121e22b3fd4ecf02d2891d3b8738a373da2aefc8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:http_proxy:enable_http", "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:http_proxy", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "documentation/data-sources/proxy/properties/http_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302", "registry_path": "docs/guides/data-sources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy"], "schema_version": 1, "sections": [{"aliases": ["enable http"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:enable_http", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "enable_http"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "more_option"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/http_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for HTTP Connect Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- http_proxy

<a id="section"></a>

Type: `"single"`. Computed.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

## Direct properties

- [enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/enable_http/): complete subsection reference.

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/): complete subsection reference.

## Next pages

- [http_proxy.enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/enable_http/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
