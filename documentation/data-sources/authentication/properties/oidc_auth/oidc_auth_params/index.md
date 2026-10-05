---
page_title: "oidc_auth.oidc_auth_params"
subcategory: ""
description: "Configuration parameter for oidc auth params."
xcsh_docs: {"aliases": ["oidc auth oidc auth params"], "body_bytes": 4573, "body_sha256": "sha256:945114f6e2be5790d3001ea190721a165dc8015004900cf83843cb4fb2d9d503", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "parent_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "path": "documentation/data-sources/authentication/properties/oidc_auth/oidc_auth_params/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1221330010102002-2102120122013232-1202123012303113-0312030123202321-0102003103331033-0213013330332320-1012102300010132-2232213123132012", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth", "oidc_auth_params"], "schema_version": 1, "sections": [{"aliases": ["oidc auth oidc auth params auth endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--auth_endpoint_url", "description": "URL of the authorization server's authorization endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "auth_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params end session endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--end_session_endpoint_url", "description": "URL of the authorization server's Logout endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "end_session_endpoint_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc auth params token endpoint url"], "anchor": "schema-oidc_auth--oidc_auth_params--token_endpoint_url", "description": "URL of the authorization server's Token endpoint.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params", "token_endpoint_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/oidc_auth_params/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for oidc auth params.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
