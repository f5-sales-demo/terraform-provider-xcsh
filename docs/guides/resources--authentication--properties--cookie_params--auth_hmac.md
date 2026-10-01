---
page_title: "cookie_params.auth_hmac"
subcategory: ""
description: "cookie_params.auth_hmac for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 3527, "body_sha256": "sha256:7bac2abc2e4bb23c112ca16f7dbf74951e2046090fb380143b655341af7bfd86", "canonical_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params", "path": "docs/guides/resources--authentication--properties--cookie_params--auth_hmac.md", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_params", "auth_hmac"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_params.auth_hmac for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- [Property reference](resources--authentication--reference.md)
- [cookie_params](resources--authentication--properties--cookie_params.md)
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

- [prim_key](resources--authentication--properties--cookie_params--auth_hmac--prim_key.md): complete subsection reference.

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

- [sec_key](resources--authentication--properties--cookie_params--auth_hmac--sec_key.md): complete subsection reference.

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

- [cookie_params.auth_hmac.prim_key](resources--authentication--properties--cookie_params--auth_hmac--prim_key.md)
- [cookie_params.auth_hmac.sec_key](resources--authentication--properties--cookie_params--auth_hmac--sec_key.md)
- [cookie_params](resources--authentication--properties--cookie_params.md)
- [xcsh_authentication](../resources/authentication.md)
