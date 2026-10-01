---
page_title: "jwt_validation.reserved_claims.issuer_disable"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims.issuer_disable for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1225, "body_sha256": "sha256:e162b0f3f795a4a7a75992361f55ef9590c21f98d85e453281e4253e207ef488", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "path": "docs/guides/resources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "issuer_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims.issuer_disable for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.issuer_disable

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [jwt_validation](resources--cdn_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- jwt_validation.reserved_claims.issuer_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
issuer_disable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
