---
page_title: "jwt_validation.reserved_claims"
subcategory: "Load Balancing"
description: "jwt_validation.reserved_claims for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3704, "body_sha256": "sha256:44740e14bd1fc92b7566b2f32444b21a4024014e5f82d3cd46a14be007798412", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:audience", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--reserved_claims.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "reserved_claims"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/reserved_claims/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.reserved_claims for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- jwt_validation.reserved_claims

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("audience",
    "audience_disable"),
  validators.ConflictingObjectAttributes("issuer",
    "issuer_disable"),
  validators.ConflictingObjectAttributes("validate_period_disable",
    "validate_period_enable")}
```

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

Terraform syntax:

```terraform
reserved_claims {
  # Configure direct properties listed below.
}
```

## Direct properties

- [audience](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--audience.md): complete subsection reference.

- [audience_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--audience_disable.md): complete subsection reference.

<a id="schema-jwt_validation--reserved_claims--issuer"></a>

### issuer property

Type: `"string"`. Optional.

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

- [issuer_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md): complete subsection reference.

- [validate_period_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_disable.md): complete subsection reference.

- [validate_period_enable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_enable.md): complete subsection reference.

## Next pages

- [jwt_validation.reserved_claims.audience](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--audience.md)
- [jwt_validation.reserved_claims.audience_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--audience_disable.md)
- [jwt_validation.reserved_claims.issuer_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--issuer_disable.md)
- [jwt_validation.reserved_claims.validate_period_disable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_disable.md)
- [jwt_validation.reserved_claims.validate_period_enable](resources--http_loadbalancer--properties--jwt_validation--reserved_claims--validate_period_enable.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
