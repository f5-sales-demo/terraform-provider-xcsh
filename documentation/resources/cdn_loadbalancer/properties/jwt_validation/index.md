---
page_title: "jwt_validation"
subcategory: "Load Balancing"
description: "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid."
xcsh_docs: {"aliases": ["backend servers", "jwt validation", "origin servers", "upstream servers"], "body_bytes": 4093, "body_sha256": "sha256:bd7634e0748de7cbc039c511edd9fd26e50175fd3e7ecd9293028c6d112cef6b", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:jwks_config", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:mandatory_claims", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:token_location"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation:ConflictingObjectAttributes:authorization_server,jwks_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation:ConflictingObjectAttributes:authorization_server,jwks_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:jwks_config", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "Action", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.action:ConflictingObjectAttributes:block,report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.action:ConflictingObjectAttributes:block,report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report", "type": "conflicts"}], "schema_path": ["jwt_validation", "action"], "syntax": "block", "type": "object"}, {"aliases": ["authorization server"], "anchor": "section", "description": "Reference to Authorization Server object.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.authorization_server:RequiredObjectAttributes:authorization_servers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:authorization_server:authorization_servers", "type": "requires"}], "schema_path": ["jwt_validation", "authorization_server"], "syntax": "block", "type": "object"}, {"aliases": ["jwks config"], "anchor": "section", "description": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:jwks_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "jwks_config"], "syntax": "block", "type": "object"}, {"aliases": ["mandatory claims"], "anchor": "section", "description": "Configurable Validation of mandatory Claims.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:mandatory_claims", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "mandatory_claims"], "syntax": "block", "type": "object"}, {"aliases": ["reserved claims"], "anchor": "section", "description": "Configurable Validation of reserved Claims.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-jwt_validation--reserved_claims--issuer", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:audience,audience_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:issuer,issuer_disable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:issuer_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.reserved_claims:ConflictingObjectAttributes:validate_period_disable,validate_period_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:reserved_claims:validate_period_enable", "type": "conflicts"}], "schema_path": ["jwt_validation", "reserved_claims"], "syntax": "block", "type": "object"}, {"aliases": ["target"], "anchor": "section", "description": "Define endpoints for which JWT token validation will be performed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:all_endpoint,api_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:all_endpoint,base_paths", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:all_endpoint,api_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:api_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:api_groups,base_paths", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:api_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:all_endpoint,base_paths", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:base_paths", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.target:ConflictingObjectAttributes:api_groups,base_paths", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:base_paths", "type": "conflicts"}], "schema_path": ["jwt_validation", "target"], "syntax": "block", "type": "object"}, {"aliases": ["token location"], "anchor": "section", "description": "Location of JWT in HTTP request.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:token_location", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "token_location"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- jwt_validation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("authorization_server",
    "jwks_config")}
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
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/action/): complete subsection reference.

- [authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/authorization_server/): complete subsection reference.

- [jwks_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/jwks_config/): complete subsection reference.

- [mandatory_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/mandatory_claims/): complete subsection reference.

- [reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/): complete subsection reference.

- [target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/): complete subsection reference.

- [token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/token_location/): complete subsection reference.

## Next pages

- [jwt_validation.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/action/)
- [jwt_validation.authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/authorization_server/)
- [jwt_validation.jwks_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/jwks_config/)
- [jwt_validation.mandatory_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/mandatory_claims/)
- [jwt_validation.reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/reserved_claims/)
- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/)
- [jwt_validation.token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/token_location/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
