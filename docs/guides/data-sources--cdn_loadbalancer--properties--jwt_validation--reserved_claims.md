---
page_title: "jwt_validation.reserved_claims"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3261, "body_sha256": "sha256:a88927c38c2583e2a0ec8948a31fc01bb4db6ff2a9d687494fb0b4e7855b9c40", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
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

- [audience](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience.md): complete subsection reference.

- [audience_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience_disable.md): complete subsection reference.

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

- [issuer_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md): complete subsection reference.

- [validate_period_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_disable.md): complete subsection reference.

- [validate_period_enable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_enable.md): complete subsection reference.

## Next pages

- [jwt_validation.reserved_claims.audience](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience.md)
- [jwt_validation.reserved_claims.audience_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--audience_disable.md)
- [jwt_validation.reserved_claims.issuer_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md)
- [jwt_validation.reserved_claims.validate_period_disable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_disable.md)
- [jwt_validation.reserved_claims.validate_period_enable](data-sources--cdn_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_enable.md)
- [jwt_validation](data-sources--cdn_loadbalancer--properties--jwt_validation.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
