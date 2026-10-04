---
page_title: "oidc_auth"
subcategory: ""
description: "OIDCAuthType."
xcsh_docs: {"aliases": ["oidc auth"], "body_bytes": 4599, "body_sha256": "sha256:2dce4d532b9ad46e027b541d8f984888a3ae9320d69dc6cf519242a6b8be6a4f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth", "parent_id": "xcsh-docs:resources:authentication:reference", "path": "documentation/resources/authentication/properties/oidc_auth/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "schema-oidc_auth--oidc_well_known_config_url", "enforcement": "provider-schema", "group": "oidc_auth:ConflictingObjectAttributes:oidc_auth_params,oidc_well_known_config_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth:ConflictingObjectAttributes:oidc_auth_params,oidc_well_known_config_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "conflicts"}, {"anchor": "schema-oidc_auth--oidc_client_id", "enforcement": "provider-schema", "group": "oidc_auth:RequiredObjectAttributes:oidc_client_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth"], "schema_version": 1, "sections": [{"aliases": ["oidc auth client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "type": "conflicts"}], "schema_path": ["oidc_auth", "client_secret"], "syntax": "block", "type": "object"}, {"aliases": ["oidc auth oidc auth params"], "anchor": "section", "description": "Configuration parameter for oidc auth params.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}], "schema_path": ["oidc_auth", "oidc_auth_params"], "syntax": "block", "type": "object"}, {"aliases": ["oidc auth oidc client id"], "anchor": "schema-oidc_auth--oidc_client_id", "description": "Client ID used while sending the Authorization Request to OIDC server.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc well known config url"], "anchor": "schema-oidc_auth--oidc_well_known_config_url", "description": "Exclusive with An OIDC well-known configuration URL that will be used to fetch authentication related endpoints.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_well_known_config_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "OIDCAuthType.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Upstream description:

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [oidc_auth.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/)
- [oidc_auth.oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/oidc_auth_params/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
