---
page_title: "http_proxy.more_option.response_cookies_to_add.samesite_none"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http proxy more option response cookies to add samesite none"], "body_bytes": 1335, "body_sha256": "sha256:d87420b25908094361f6e5c31ab3861d9fa83ed23e1d236204cd34c26b1250e4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add:samesite_none", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:response_cookies_to_add", "path": "documentation/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/samesite_none/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "response_cookies_to_add", "samesite_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/samesite_none/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["proxyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.response_cookies_to_add.samesite_none

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [http_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/response_cookies_to_add/)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.
