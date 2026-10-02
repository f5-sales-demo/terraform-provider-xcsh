---
page_title: "cookie_params.auth_hmac"
subcategory: ""
description: "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid."
xcsh_docs: {"aliases": ["cookie params auth hmac"], "body_bytes": 3980, "body_sha256": "sha256:b90811e0c7099027bffc33abd229f1722f99da07c0875a3d21954d30e90c0a0c", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params", "path": "documentation/resources/authentication/properties/cookie_params/auth_hmac/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "schema-cookie_params--auth_hmac--prim_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}, {"anchor": "schema-cookie_params--auth_hmac--sec_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_params", "auth_hmac"], "schema_version": 1, "sections": [{"aliases": ["prim key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["cookie_params", "auth_hmac", "prim_key"], "syntax": "block", "type": "object"}, {"aliases": ["prim key expiry"], "anchor": "schema-cookie_params--auth_hmac--prim_key_expiry", "description": "Primary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "auth_hmac", "prim_key_expiry"], "syntax": "attribute", "type": "string"}, {"aliases": ["sec key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["cookie_params", "auth_hmac", "sec_key"], "syntax": "block", "type": "object"}, {"aliases": ["sec key expiry"], "anchor": "schema-cookie_params--auth_hmac--sec_key_expiry", "description": "Secondary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "auth_hmac", "sec_key_expiry"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/)
- cookie_params.auth_hmac

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prim_key_expiry",
    "sec_key_expiry")}
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
auth_hmac {
  # Configure direct properties listed below.
}
```

## Direct properties

- [prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/): complete subsection reference.

<a id="schema-cookie_params--auth_hmac--prim_key_expiry"></a>

### prim_key_expiry property

Type: `"string"`. Optional.

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

- [sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/sec_key/): complete subsection reference.

<a id="schema-cookie_params--auth_hmac--sec_key_expiry"></a>

### sec_key_expiry property

Type: `"string"`. Optional.

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

- [cookie_params.auth_hmac.prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/)
- [cookie_params.auth_hmac.sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/sec_key/)
- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
