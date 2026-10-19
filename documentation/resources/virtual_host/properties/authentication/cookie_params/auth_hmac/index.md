---
page_title: "authentication.cookie_params.auth_hmac"
subcategory: ""
description: "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid."
xcsh_docs: {"aliases": ["authentication cookie params auth hmac"], "body_bytes": 3184, "body_sha256": "sha256:2e21e217a884a26ed8bd501eb4a3c20c634a45673f53526fc8b392ee901c7dfe", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac"], "schema_version": 1, "sections": [{"aliases": ["authentication cookie params auth hmac prim key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key"], "syntax": "block", "type": "object"}, {"aliases": ["authentication cookie params auth hmac prim key expiry"], "anchor": "schema-authentication--cookie_params--auth_hmac--prim_key_expiry", "description": "Primary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key_expiry"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication cookie params auth hmac sec key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key"], "syntax": "block", "type": "object"}, {"aliases": ["authentication cookie params auth hmac sec key expiry"], "anchor": "schema-authentication--cookie_params--auth_hmac--sec_key_expiry", "description": "Secondary HMAC Key Expiry time.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key_expiry"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- authentication.cookie_params.auth_hmac

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
auth_hmac {
  # Configure direct properties listed below.
}
```

## Direct properties

- [prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/): complete subsection reference.

<a id="schema-authentication--cookie_params--auth_hmac--prim_key_expiry"></a>

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

- [sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/): complete subsection reference.

<a id="schema-authentication--cookie_params--auth_hmac--sec_key_expiry"></a>

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
