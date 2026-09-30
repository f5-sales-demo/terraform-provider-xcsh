---
page_title: "authentication.cookie_params.auth_hmac"
subcategory: ""
description: "authentication.cookie_params.auth_hmac for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 3435, "body_sha256": "sha256:4a937a1608ff00ff2b1cdd049f93f0dc146481f087a37a5e3882cb3ec435d565", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:auth_hmac", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "path": "docs/guides/data-sources--virtual_host--properties--authentication--cookie_params--auth_hmac.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/cookie_params/auth_hmac/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication.cookie_params.auth_hmac for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# authentication.cookie_params.auth_hmac

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [authentication](data-sources--virtual_host--properties--authentication.md)
- [authentication.cookie_params](data-sources--virtual_host--properties--authentication--cookie_params.md)
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

- [prim_key](data-sources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key.md): complete subsection reference.

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

- [sec_key](data-sources--virtual_host--properties--authentication--cookie_params--auth_hmac--sec_key.md): complete subsection reference.

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

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key.md)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--properties--authentication--cookie_params--auth_hmac--sec_key.md)
- [authentication.cookie_params](data-sources--virtual_host--properties--authentication--cookie_params.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
