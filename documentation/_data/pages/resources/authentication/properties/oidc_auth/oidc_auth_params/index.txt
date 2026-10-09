---
page_title: "oidc_auth.oidc_auth_params"
subcategory: ""
description: "Configuration parameter for oidc auth params."
xcsh_docs: {"aliases": ["oidc auth oidc auth params"], "body_bytes": 5199, "body_sha256": "sha256:e4582780fd992ea2c2e082ee32b6a6cf5dc74349e7d6b66b91319b80fdf6611c", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "parent_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "path": "documentation/resources/authentication/properties/oidc_auth/oidc_auth_params/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2121211332103330-0132113033112313-0200321101212122-0022032210032230-2112313233102322-2101230323030132-2030321103332312-1323132000201232", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}, {"anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "enforcement": "provider-schema", "group": "oidc_auth.oidc_auth_params:RequiredObjectAttributes:auth_endpoint_url,end_session_endpoint_url,token_endpoint_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth", "oidc_auth_params"], "schema_version": 1, "sections": [{"aliases": ["oidc auth oidc auth params auth endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "description": "URL of the authorization server's authorization endpoint.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "auth_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params end session endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "description": "URL of the authorization server's Logout endpoint.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "end_session_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params token endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "description": "URL of the authorization server's Token endpoint.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "token_endpoint_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/oidc_auth_params/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for oidc auth params.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["authenticationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth.oidc_auth_params

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/)
- oidc_auth.oidc_auth_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for oidc auth params.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_endpoint_url",
    "end_session_endpoint_url",
    "token_endpoint_url")}
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
oidc_auth_params {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-oidc_auth--oidc_auth_params--auth_endpoint_url"></a>

### auth_endpoint_url property

Type: `"string"`. Optional.

URL of the authorization server's authorization endpoint.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-oidc_auth--oidc_auth_params--end_session_endpoint_url"></a>

### end_session_endpoint_url property

Type: `"string"`. Optional.

URL of the authorization server's Logout endpoint.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-oidc_auth--oidc_auth_params--token_endpoint_url"></a>

### token_endpoint_url property

Type: `"string"`. Optional.

URL of the authorization server's Token endpoint.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```
