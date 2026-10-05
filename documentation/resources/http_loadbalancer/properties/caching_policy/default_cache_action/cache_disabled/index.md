---
page_title: "caching_policy.default_cache_action.cache_disabled"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["caching policy default cache action cache disabled"], "body_bytes": 1544, "body_sha256": "sha256:5b346d6d3f4769c5e10a047fbca0d4c89bf16ca61d5d7a3a55a3705139458fa2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action:cache_disabled", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "path": "documentation/resources/http_loadbalancer/properties/caching_policy/default_cache_action/cache_disabled/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy", "default_cache_action", "cache_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/default_cache_action/cache_disabled/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.default_cache_action.cache_disabled

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/)
- [caching_policy.default_cache_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/default_cache_action/)
- caching_policy.default_cache_action.cache_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
cache_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [caching_policy.default_cache_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/default_cache_action/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
