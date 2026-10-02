---
page_title: "jwt_validation.reserved_claims"
subcategory: "Load Balancing"
description: "Configurable Validation of reserved Claims."
xcsh_docs: {"aliases": ["jwt validation reserved claims"], "body_bytes": 4434, "body_sha256": "sha256:135d2ef9b120d6b94767088bf32ec569f88cff08a89ffaf10163811e28afe2d7", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--issuer", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "reserved_claims"], "schema_version": 1, "sections": [{"aliases": ["audience"], "anchor": "section", "description": "Audiences", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--audience--audiences", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims.audience:RequiredObjectAttributes:audiences", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "requires"}], "schema_path": ["jwt_validation", "reserved_claims", "audience"], "syntax": "block", "type": "object"}, {"aliases": ["audience disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "audience_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["issuer"], "anchor": "schema-jwt_validation--reserved_claims--issuer", "description": "Exclusive with", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "issuer"], "syntax": "attribute", "type": "string"}, {"aliases": ["issuer disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "issuer_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["validate period disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "validate_period_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["validate period enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "validate_period_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configurable Validation of reserved Claims.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
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

- [audience](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/): complete subsection reference.

- [audience_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/): complete subsection reference.

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

- [issuer_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/): complete subsection reference.

- [validate_period_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_disable/): complete subsection reference.

- [validate_period_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_enable/): complete subsection reference.

## Next pages

- [jwt_validation.reserved_claims.audience](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/)
- [jwt_validation.reserved_claims.audience_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/)
- [jwt_validation.reserved_claims.issuer_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/issuer_disable/)
- [jwt_validation.reserved_claims.validate_period_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_disable/)
- [jwt_validation.reserved_claims.validate_period_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/validate_period_enable/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
