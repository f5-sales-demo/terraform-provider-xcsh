---
page_title: "jwt_validation.token_location.bearer_token"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["jwt validation token location bearer token"], "body_bytes": 1171, "body_sha256": "sha256:ba0e3f96be9cb16a76e2357cd9b96ed78b151ca31091643f249e6dda3226373f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:token_location:bearer_token", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:token_location", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/token_location/bearer_token/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1032111232021013-2223332323312312-0230231122101003-0210202333210322-0212233233233120-3120300221122332-3033230201120130-3220010301201222", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "token_location", "bearer_token"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/token_location/bearer_token/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.token_location.bearer_token

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- [jwt_validation.token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/token_location/)
- jwt_validation.token_location.bearer_token

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

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

This is an empty object or choice marker. It has no direct properties.
