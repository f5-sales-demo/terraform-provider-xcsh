---
page_title: "jwt_validation.reserved_claims"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4008, "body_sha256": "sha256:e2eec052bc6b393a8c8e44a44e53f735be2943d4f7e0767888600ca51498dfca", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.reserved_claims

<a id="section"></a>

Type: `"single"`. Computed.

Configurable Validation of reserved Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

## Direct properties

- [audience](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/): complete subsection reference.

- [audience_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/): complete subsection reference.

<a id="schema-jwt_validation--reserved_claims--issuer"></a>

### issuer property

Type: `"string"`. Computed.

Exact Match. Exclusive with \[issuer\_disable\]

Upstream description:

Exclusive with \[issuer\_disable\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [issuer_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/): complete subsection reference.

- [validate_period_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_disable/): complete subsection reference.

- [validate_period_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_enable/): complete subsection reference.

## Next pages

- [jwt_validation.reserved_claims.audience](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/)
- [jwt_validation.reserved_claims.audience_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/)
- [jwt_validation.reserved_claims.issuer_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/)
- [jwt_validation.reserved_claims.validate_period_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_disable/)
- [jwt_validation.reserved_claims.validate_period_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_enable/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
