---
page_title: "cookie_params"
subcategory: ""
description: "Specifies different cookie related config parameters for authentication."
xcsh_docs: {"aliases": ["cookie params"], "body_bytes": 6072, "body_sha256": "sha256:367452c03dd69d40e44053ce15eb6216a5cdff66cafdd36d5b006f6fdf85960f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "xcsh-docs:resources:authentication:properties:cookie_params:kms_key_hmac"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params", "parent_id": "xcsh-docs:resources:authentication:reference", "path": "documentation/resources/authentication/properties/cookie_params/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params:ConflictingObjectAttributes:auth_hmac,kms_key_hmac", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params:ConflictingObjectAttributes:auth_hmac,kms_key_hmac", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:kms_key_hmac", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_params"], "schema_version": 1, "sections": [{"aliases": ["cookie params auth hmac"], "anchor": "section", "description": "HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cookie_params--auth_hmac--prim_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}, {"anchor": "schema-cookie_params--auth_hmac--sec_key_expiry", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac:RequiredObjectAttributes:prim_key_expiry,sec_key_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "type": "requires"}], "schema_path": ["cookie_params", "auth_hmac"], "syntax": "block", "type": "object"}, {"aliases": ["cookie params cookie expiry"], "anchor": "schema-cookie_params--cookie_expiry", "description": "Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the session cookie. This will act as an expiry duration on the client side after which client will not be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "cookie_expiry"], "syntax": "attribute", "type": "number"}, {"aliases": ["cookie params cookie refresh interval", "login", "login result", "sign in"], "anchor": "schema-cookie_params--cookie_refresh_interval", "description": "Specifies in seconds refresh interval for session cookie. This is used to keep the active user active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to expire falls behind this interval, RE-issue a cookie with new expiry and with the same original session expiry. Default refresh", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "cookie_refresh_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["cookie params kms key hmac"], "anchor": "section", "description": "Reference to KMS Key Object.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:kms_key_hmac", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "kms_key_hmac"], "syntax": "attribute", "type": "object"}, {"aliases": ["cookie params session expiry", "login", "login result", "sign in"], "anchor": "schema-cookie_params--session_expiry", "description": "Specifies in seconds max lifetime of an authenticated session after which the user will be forced to login again. Default session expiry is 86400 seconds(24 hours).", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookie_params", "session_expiry"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Specifies different cookie related config parameters for authentication.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- cookie_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies different cookie related config parameters for authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_hmac",
    "kms_key_hmac")}
```

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

Terraform syntax:

```terraform
cookie_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/): complete subsection reference.

<a id="schema-cookie_params--cookie_expiry"></a>

### cookie_expiry property

Type: `"number"`. Optional.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-cookie_params--cookie_refresh_interval"></a>

### cookie_refresh_interval property

Type: `"number"`. Optional.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [kms_key_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/kms_key_hmac/): complete subsection reference.

<a id="schema-cookie_params--session_expiry"></a>

### session_expiry property

Type: `"number"`. Optional.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1296000),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/)
- [cookie_params.kms_key_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/kms_key_hmac/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
