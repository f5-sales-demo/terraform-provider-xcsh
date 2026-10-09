---
page_title: "http_proxy.more_option.no_request_limit_per_connection"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http proxy more option no request limit per connection"], "body_bytes": 1200, "body_sha256": "sha256:b618df3bfbc8833e827708ca78bf102503010ad3f3e772956f503615525d1f74", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "path": "documentation/resources/proxy/properties/http_proxy/more_option/no_request_limit_per_connection/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "no_request_limit_per_connection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- http_proxy.more_option.no_request_limit_per_connection

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
