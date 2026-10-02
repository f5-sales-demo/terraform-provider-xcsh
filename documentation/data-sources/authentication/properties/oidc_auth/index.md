---
page_title: "oidc_auth"
subcategory: ""
description: "OIDCAuthType."
xcsh_docs: {"aliases": ["oidc auth"], "body_bytes": 3972, "body_sha256": "sha256:730b393804d2370d6938a14e718b76a0dc6844406bf5a0a4899cf64368535a91", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "parent_id": "xcsh-docs:data-sources:authentication:reference", "path": "documentation/data-sources/authentication/properties/oidc_auth/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth"], "schema_version": 1, "sections": [{"aliases": ["client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "client_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["oidc auth params"], "anchor": "section", "description": "Configuration parameter for oidc auth params.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:oidc_auth_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "oidc_auth_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["oidc client id"], "anchor": "schema-oidc_auth--oidc_client_id", "description": "Client ID used while sending the Authorization Request to OIDC server.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc well known config url"], "anchor": "schema-oidc_auth--oidc_well_known_config_url", "description": "Exclusive with An OIDC well-known configuration URL that will be used to fetch authentication related endpoints.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oidc_auth", "oidc_well_known_config_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OIDCAuthType.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [oidc_auth.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/)
- [oidc_auth.oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
