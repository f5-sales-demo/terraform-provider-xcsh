---
page_title: "jwt_validation.reserved_claims.audience"
subcategory: "Load Balancing"
description: "Audiences"
xcsh_docs: {"aliases": ["jwt validation reserved claims audience"], "body_bytes": 2908, "body_sha256": "sha256:2e763aca8537dbafc4c0b55a1ebde06401836f3cbc294afb1e363442e15f061c", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--audience--audiences", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims.audience:RequiredObjectAttributes:audiences", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "audience"], "schema_version": 1, "sections": [{"aliases": ["audiences"], "anchor": "schema-jwt_validation--reserved_claims--audience--audiences", "description": "Configuration parameter for audiences", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "reserved_claims", "audience", "audiences"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/audience/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Audiences", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.audience

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- [jwt_validation.reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/)
- jwt_validation.reserved_claims.audience

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("audiences")}
```

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
audience {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-jwt_validation--reserved_claims--audience--audiences"></a>

### audiences property

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [jwt_validation.reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
