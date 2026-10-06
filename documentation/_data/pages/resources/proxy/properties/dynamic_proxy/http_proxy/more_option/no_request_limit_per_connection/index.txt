---
page_title: "dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dynamic proxy http proxy more option no request limit per connection"], "body_bytes": 1401, "body_sha256": "sha256:0d86948eb9cb6680e963e448cf01178a8c8da076358ee9624cf003fc5a0e69cb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "path": "documentation/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/no_request_limit_per_connection/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1300132130201131-0303023011210331-2213112112010120-1033210011003131-0103131233101201-1213132101102322-3112320303333013-0033212231102002", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "no_request_limit_per_connection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/)
- [dynamic_proxy.http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.
