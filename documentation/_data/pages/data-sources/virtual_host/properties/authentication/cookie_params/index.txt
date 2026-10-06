---
page_title: "authentication.cookie_params"
subcategory: ""
description: "Specifies different cookie related config parameters for authentication."
xcsh_docs: {"aliases": ["authentication cookie params"], "body_bytes": 4299, "body_sha256": "sha256:e82efeeef02825ed1ac68efe71be923d97872d23942a226bd04155e46796d9d6", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:kms_key_hmac"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:authentication", "path": "documentation/data-sources/virtual_host/properties/authentication/cookie_params/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params"], "schema_version": 1, "sections": [{"aliases": ["authentication cookie params auth hmac"], "anchor": "section", "description": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication cookie params cookie expiry"], "anchor": "schema-authentication--cookie_params--cookie_expiry", "description": "Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the session cookie. This will act as an expiry duration on the client side after which client will not be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "cookie_expiry"], "syntax": "attribute", "type": "number"}, {"aliases": ["authentication cookie params cookie refresh interval", "login", "login result", "sign in"], "anchor": "schema-authentication--cookie_params--cookie_refresh_interval", "description": "Specifies in seconds refresh interval for session cookie. This is used to keep the active user active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to expire falls behind this interval, RE-issue a cookie with new expiry and with the same original session expiry. Default refresh", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "cookie_refresh_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["authentication cookie params kms key hmac"], "anchor": "section", "description": "Reference to KMS Key Object.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "kms_key_hmac"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication cookie params session expiry", "login", "login result", "sign in"], "anchor": "schema-authentication--cookie_params--session_expiry", "description": "Specifies in seconds max lifetime of an authenticated session after which the user will be forced to login again. Default session expiry is 86400 seconds(24 hours).", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "session_expiry"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/cookie_params/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Specifies different cookie related config parameters for authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/)
- authentication.cookie_params

<a id="section"></a>

Type: `"single"`. Computed.

Specifies different cookie related config parameters for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

## Direct properties

- [auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/): complete subsection reference.

<a id="schema-authentication--cookie_params--cookie_expiry"></a>

### cookie_expiry property

Type: `"number"`. Computed.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="schema-authentication--cookie_params--cookie_refresh_interval"></a>

### cookie_refresh_interval property

Type: `"number"`. Computed.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/): complete subsection reference.

<a id="schema-authentication--cookie_params--session_expiry"></a>

### session_expiry property

Type: `"number"`. Computed.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```
