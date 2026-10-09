---
page_title: "cname_pool.disable_health_check"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cname pool disable health check"], "body_bytes": 1021, "body_sha256": "sha256:e0be577e5e77bf915cba51026ae13ff0777df6ebdc54d08036c16ac987d0b424", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "parent_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "path": "documentation/resources/dns_lb_pool/properties/cname_pool/disable_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2213312121000132-2321302000121322-3231020123322033-2011210210232100-3222320312103232-2001130103001013-1223301213121001-2023133333112302", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool", "disable_health_check"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/cname_pool/disable_health_check/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool.disable_health_check

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/)
- cname_pool.disable_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.
