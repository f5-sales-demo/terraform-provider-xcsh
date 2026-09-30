---
page_title: "jwt_validation.reserved_claims.validate_period_enable"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims.validate_period_enable for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1166, "body_sha256": "sha256:fbbae369c6c8d7377a417f2301a316d41e5fef336cbf2b64684a030bfb6dc473", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_enable.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "validate_period_enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims.validate_period_enable for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# jwt_validation.reserved_claims.validate_period_enable

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--properties--jwt_validation--reserved_claims.md)
- jwt_validation.reserved_claims.validate_period_enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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
validate_period_enable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.reserved_claims](resources--http_loadbalancer--properties--jwt_validation--reserved_claims.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
