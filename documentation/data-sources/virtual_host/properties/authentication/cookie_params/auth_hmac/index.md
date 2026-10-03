---
page_title: "authentication.cookie_params.auth_hmac"
subcategory: ""
description: "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid."
xcsh_docs: {"aliases": ["authentication cookie params auth hmac"], "body_bytes": 4032, "body_sha256": "sha256:20cc4849b518eae9efe4fc633189131324eaf5d4550c5a6a2ca219d70ab8059c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "path": "documentation/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac"], "schema_version": 1, "sections": [{"aliases": ["authentication cookie params auth hmac prim key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication cookie params auth hmac prim key expiry"], "anchor": "schema-authentication--cookie_params--auth_hmac--prim_key_expiry", "description": "Primary HMAC Key Expiry time.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key_expiry"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication cookie params auth hmac sec key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication cookie params auth hmac sec key expiry"], "anchor": "schema-authentication--cookie_params--auth_hmac--sec_key_expiry", "description": "Secondary HMAC Key Expiry time.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key_expiry"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/)
- authentication.cookie_params.auth_hmac

<a id="section"></a>

Type: `"single"`. Computed.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

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

- [prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/): complete subsection reference.

<a id="schema-authentication--cookie_params--auth_hmac--prim_key_expiry"></a>

### prim_key_expiry property

Type: `"string"`. Computed.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/): complete subsection reference.

<a id="schema-authentication--cookie_params--auth_hmac--sec_key_expiry"></a>

### sec_key_expiry property

Type: `"string"`. Computed.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [authentication.cookie_params.auth_hmac.prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/)
- [authentication.cookie_params.auth_hmac.sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
