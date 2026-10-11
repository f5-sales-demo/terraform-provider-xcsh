---
page_title: "jwt_validation"
subcategory: "Load Balancing"
description: "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid."
xcsh_docs: {"aliases": ["jwt validation"], "body_bytes": 2313, "body_sha256": "sha256:fd3468809c4d28d6a32f06d5f8dced8dc414367e20970664b2b5e26a647a089f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:jwks_config", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:mandatory_claims", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/jwt_validation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation"], "schema_version": 1, "sections": [{"aliases": ["jwt validation action"], "anchor": "section", "description": "Action", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "action"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation authorization server"], "anchor": "section", "description": "Reference to Authorization Server object.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:authorization_server", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "authorization_server"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation jwks config"], "anchor": "section", "description": "The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:jwks_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "jwks_config"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation mandatory claims"], "anchor": "section", "description": "Configurable Validation of mandatory Claims.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:mandatory_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "mandatory_claims"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation reserved claims"], "anchor": "section", "description": "Configurable Validation of reserved Claims.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:reserved_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "reserved_claims"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation target"], "anchor": "section", "description": "Define endpoints for which JWT token validation will be performed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:target", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target"], "syntax": "block", "type": "object"}, {"aliases": ["jwt validation token location"], "anchor": "section", "description": "Location of JWT in HTTP request.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:token_location", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "token_location"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- jwt_validation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/action/): complete subsection reference.

- [authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/authorization_server/): complete subsection reference.

- [jwks_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/jwks_config/): complete subsection reference.

- [mandatory_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/mandatory_claims/): complete subsection reference.

- [reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/reserved_claims/): complete subsection reference.

- [target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/target/): complete subsection reference.

- [token_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/token_location/): complete subsection reference.
