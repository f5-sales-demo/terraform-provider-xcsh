---
page_title: "oidc_auth"
subcategory: ""
description: "OIDCAuthType."
xcsh_docs: {"aliases": ["oidc auth"], "body_bytes": 3262, "body_sha256": "sha256:729f551d4e47232a98e3fd6ebd2aff8b29dc02e715475b3e1dd714ef7f8e9282", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "parent_id": "xcsh-docs:data-sources:authentication:reference", "path": "documentation/data-sources/authentication/properties/oidc_auth/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth"], "schema_version": 1, "sections": [{"aliases": ["oidc auth client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "client_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["oidc auth oidc auth params"], "anchor": "section", "description": "Configuration parameter for oidc auth params.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["oidc auth oidc client id"], "anchor": "schema-oidc_auth--oidc_client_id", "description": "Client ID used while sending the Authorization Request to OIDC server.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth oidc well known config url"], "anchor": "schema-oidc_auth--oidc_well_known_config_url", "description": "Exclusive with An OIDC well-known configuration URL that will be used to fetch authentication related endpoints.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_well_known_config_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "OIDCAuthType.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["authenticationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- oidc_auth

<a id="section"></a>

Type: `"single"`. Computed.

OIDCAuthType.

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

## Direct properties

- [client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/): complete subsection reference.

- [oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/): complete subsection reference.

<a id="schema-oidc_auth--oidc_client_id"></a>

### oidc_client_id property

Type: `"string"`. Computed.

Client ID used while sending the Authorization Request to OIDC server.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"string"`. Computed.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
