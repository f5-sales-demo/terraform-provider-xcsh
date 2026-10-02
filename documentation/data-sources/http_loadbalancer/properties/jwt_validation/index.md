---
page_title: "jwt_validation"
subcategory: "Load Balancing"
description: "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid."
xcsh_docs: {"aliases": ["backend servers", "jwt validation", "origin servers", "upstream servers"], "body_bytes": 3863, "body_sha256": "sha256:cf7839905fd0139bb0cbda0ab23600fa98559066c8bd1ac8a1dcdb87719429ce", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:mandatory_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:token_location"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "Action", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "action"], "syntax": "attribute", "type": "object"}, {"aliases": ["authorization server"], "anchor": "section", "description": "Reference to Authorization Server object.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:authorization_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "authorization_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwks config"], "anchor": "section", "description": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:jwks_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "jwks_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["mandatory claims"], "anchor": "section", "description": "Configurable Validation of mandatory Claims.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:mandatory_claims", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "mandatory_claims"], "syntax": "attribute", "type": "object"}, {"aliases": ["reserved claims"], "anchor": "section", "description": "Configurable Validation of reserved Claims.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "reserved_claims"], "syntax": "attribute", "type": "object"}, {"aliases": ["target"], "anchor": "section", "description": "Define endpoints for which JWT token validation will be performed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target"], "syntax": "attribute", "type": "object"}, {"aliases": ["token location"], "anchor": "section", "description": "Location of JWT in HTTP request.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:token_location", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "token_location"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- jwt_validation

<a id="section"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

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

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/action/): complete subsection reference.

- [authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/): complete subsection reference.

- [jwks_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/): complete subsection reference.

- [mandatory_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/mandatory_claims/): complete subsection reference.

- [reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/): complete subsection reference.

- [target](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/): complete subsection reference.

- [token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/token_location/): complete subsection reference.

## Next pages

- [jwt_validation.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/action/)
- [jwt_validation.authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/authorization_server/)
- [jwt_validation.jwks_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/jwks_config/)
- [jwt_validation.mandatory_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/mandatory_claims/)
- [jwt_validation.reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/)
- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/)
- [jwt_validation.token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/token_location/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
