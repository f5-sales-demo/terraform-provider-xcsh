---
page_title: "jwt_validation.reserved_claims.issuer_disable"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims.issuer_disable for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1194, "body_sha256": "sha256:a8236ee32521aa34c7e61e85d20338729dcef7dfb557d097249b10f0b22c98f9", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims", "path": "docs/guides/data-sources--http_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "issuer_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims.issuer_disable for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.issuer_disable

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--properties--jwt_validation--reserved_claims.md)
- jwt_validation.reserved_claims.issuer_disable

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for issuer disable.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
