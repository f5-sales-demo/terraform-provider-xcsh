---
page_title: "jwt_validation.reserved_claims.audience_disable"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims.audience_disable for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1192, "body_sha256": "sha256:8c4cbc48a99aca74eb6a9247da233c8556c15e38af0fe3d664b6507a87f3553e", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience_disable.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "audience_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims.audience_disable for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.audience_disable

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- jwt_validation.reserved_claims.audience_disable

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for audience disable.

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

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
