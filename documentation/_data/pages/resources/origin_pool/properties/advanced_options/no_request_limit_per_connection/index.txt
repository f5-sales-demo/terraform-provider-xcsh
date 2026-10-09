---
page_title: "advanced_options.no_request_limit_per_connection"
subcategory: "Load Balancing"
description: "Configures no request limit per connection to the origin server."
xcsh_docs: {"aliases": ["advanced options no request limit per connection", "disable max requests per connection", "no request limit", "no request limit per connection"], "body_bytes": 1161, "body_sha256": "sha256:ed5ba473c2ab29ba668465f8ed0b57653f6e5d99e9f8c02a88c423d856385708", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule", "reviewed-summary"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "documentation/resources/origin_pool/properties/advanced_options/no_request_limit_per_connection/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1330213333313110-3100031301022021-2123130200313332-3201013222020131-2223030121213030-2220120203131331-1221101000313013-1112103003000113", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "no_request_limit_per_connection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configures no request limit per connection to the origin server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["origin_poolCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- advanced_options.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
