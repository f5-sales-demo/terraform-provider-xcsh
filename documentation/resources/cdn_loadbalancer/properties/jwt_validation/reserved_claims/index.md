---
page_title: "jwt_validation.reserved_claims"
subcategory: "Load Balancing"
description: "Configurable Validation of reserved Claims."
xcsh_docs: {"aliases": ["jwt validation reserved claims"], "body_bytes": 3162, "body_sha256": "sha256:3cc1b8e17f869c5d6c9d73330c1d6fad5e5fbc0690e31a155b58c92e9ad203a1", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--issuer", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "reserved_claims"], "schema_version": 1, "sections": [{"aliases": ["jwt validation reserved claims audience"], "anchor": "section", "description": "Audiences", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--audience--audiences", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims.audience:RequiredObjectAttributes:audiences", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "requires"}], "schema_path": ["jwt_validation", "reserved_claims", "audience"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation reserved claims audience disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "audience_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation reserved claims issuer"], "anchor": "schema-jwt_validation--reserved_claims--issuer", "description": "Exclusive with", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "issuer"], "syntax": "attribute", "type": "string"}, {"aliases": ["jwt validation reserved claims issuer disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "issuer_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation reserved claims validate period disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "validate_period_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation reserved claims validate period enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "validate_period_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configurable Validation of reserved Claims.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
