---
page_title: "oidc_auth"
subcategory: ""
description: "OIDCAuthType."
xcsh_docs: {"aliases": ["oidc auth"], "body_bytes": 3989, "body_sha256": "sha256:9d836e27a0badf8470caee2d3f8ffd06977914dff71665248b980f6e058d87a6", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth", "parent_id": "xcsh-docs:resources:authentication:reference", "path": "documentation/resources/authentication/properties/oidc_auth/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "schema-oidc_auth--oidc_well_known_config_url", "enforcement": "provider-schema", "group": "oidc_auth:ConflictingObjectAttributes:oidc_auth_params,oidc_well_known_config_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth:ConflictingObjectAttributes:oidc_auth_params,oidc_well_known_config_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "conflicts"}, {"anchor": "schema-oidc_auth--oidc_client_id", "enforcement": "provider-schema", "group": "oidc_auth:RequiredObjectAttributes:oidc_client_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth"], "schema_version": 1, "sections": [{"aliases": ["oidc auth client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "type": "conflicts"}], "schema_path": ["oidc_auth", "client_secret"], "syntax": "block", "type": "object"}, {"aliases": ["oidc auth oidc auth params"], "anchor": "section", "description": "Configuration parameter for oidc auth params.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}], "schema_path": ["oidc_auth", "oidc_auth_params"], "syntax": "block", "type": "object"}, {"aliases": ["oidc auth oidc client id"], "anchor": "schema-oidc_auth--oidc_client_id", "description": "Client ID used while sending the Authorization Request to OIDC server.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc well known config url"], "anchor": "schema-oidc_auth--oidc_well_known_config_url", "description": "Exclusive with An OIDC well-known configuration URL that will be used to fetch authentication related endpoints.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_well_known_config_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "OIDCAuthType.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["authenticationCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- oidc_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

OIDCAuthType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("oidc_client_id"),
  validators.ConflictingObjectAttributes("oidc_auth_params",
    "oidc_well_known_config_url")}
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
  "x-ves-oneof-field-auth_params_choice": "[\"oidc_auth_params\",\"oidc_well_known_config_url\"]"
}
```

Terraform syntax:

```terraform
oidc_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/): complete subsection reference.

- [oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/oidc_auth_params/): complete subsection reference.

<a id="schema-oidc_auth--oidc_client_id"></a>

### oidc_client_id property

Type: `"string"`. Optional.

Client ID used while sending the Authorization Request to OIDC server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-oidc_auth--oidc_well_known_config_url"></a>

### oidc_well_known_config_url property

Type: `"string"`. Optional.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```
