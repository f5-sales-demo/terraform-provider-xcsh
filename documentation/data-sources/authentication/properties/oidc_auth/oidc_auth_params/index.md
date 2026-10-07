---
page_title: "oidc_auth.oidc_auth_params"
subcategory: ""
description: "Configuration parameter for oidc auth params."
xcsh_docs: {"aliases": ["oidc auth oidc auth params"], "body_bytes": 4326, "body_sha256": "sha256:2cc2d50081868b8a58be9f860442d311b2e8d2ae5df5d76a31b4ab60a2ce0936", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "parent_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "path": "documentation/data-sources/authentication/properties/oidc_auth/oidc_auth_params/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1221330010102002-2102120122013232-1202123012303113-0312030123202321-0102003103331033-0213013330332320-1012102300010132-2232213123132012", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth", "oidc_auth_params"], "schema_version": 1, "sections": [{"aliases": ["oidc auth oidc auth params auth endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "description": "URL of the authorization server's authorization endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "auth_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params end session endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "description": "URL of the authorization server's Logout endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "end_session_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params token endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "description": "URL of the authorization server's Token endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "token_endpoint_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/oidc_auth_params/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration parameter for oidc auth params.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["authenticationCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth.oidc_auth_params

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/)
- oidc_auth.oidc_auth_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for oidc auth params.

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

## Direct properties

<a id="schema-oidc_auth--oidc_auth_params--auth_endpoint_url"></a>

### auth_endpoint_url property

Type: `"string"`. Computed.

URL of the authorization server's authorization endpoint.

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

Type: `"string"`. Computed.

URL of the authorization server's Logout endpoint.

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

Type: `"string"`. Computed.

URL of the authorization server's Token endpoint.

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
