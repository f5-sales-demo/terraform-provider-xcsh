---
page_title: "cookie_params.auth_hmac"
subcategory: ""
description: "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid."
xcsh_docs: {"aliases": ["cookie params auth hmac"], "body_bytes": 3154, "body_sha256": "sha256:602f3f5aea6d3cc9c51f9e3c7c2a2655670fd52de1bbd7649e61a0f070fc4c9d", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params", "path": "documentation/resources/authentication/properties/cookie_params/auth_hmac/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "schema-cookie_params--auth_hmac--prim_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}, {"anchor": "schema-cookie_params--auth_hmac--sec_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_params", "auth_hmac"], "schema_version": 1, "sections": [{"aliases": ["cookie params auth hmac prim key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["cookie_params", "auth_hmac", "prim_key"], "syntax": "block", "type": "object"}, {"aliases": ["cookie params auth hmac prim key expiry"], "anchor": "schema-cookie_params--auth_hmac--prim_key_expiry", "description": "Primary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "auth_hmac", "prim_key_expiry"], "syntax": "attribute", "type": "string"}, {"aliases": ["cookie params auth hmac sec key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["cookie_params", "auth_hmac", "sec_key"], "syntax": "block", "type": "object"}, {"aliases": ["cookie params auth hmac sec key expiry"], "anchor": "schema-cookie_params--auth_hmac--sec_key_expiry", "description": "Secondary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "auth_hmac", "sec_key_expiry"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
