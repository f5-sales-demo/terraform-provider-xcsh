---
page_title: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes simple route advanced options specific hash policy hash policy cookie ignore httponly"], "body_bytes": 2178, "body_sha256": "sha256:9077c3c332a762acae2718edbcff30c9ef6ba8b50594d0fc1bfa8bc0dcc07aa9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_httponly", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/ignore_httponly/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0230100213022323-1112030322111222-0013302010101012-0330311013321030-1200221132300221-3223032320322202-1023320020320121-3211310002023302", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "cookie", "ignore_httponly"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/ignore_httponly/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.specific_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.
