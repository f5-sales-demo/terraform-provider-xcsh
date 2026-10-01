---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 17254, "body_sha256": "sha256:853a7e5de100543568664e04352203b09cc1524a97ac026b646869c3eb8c47a6", "canonical_id": "xcsh-docs:resources:authentication:reference", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params", "xcsh-docs:resources:authentication:properties:oidc_auth", "xcsh-docs:resources:authentication:properties:timeouts"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:reference", "parent_id": "xcsh-docs:resources:authentication:fundamentals", "path": "docs/guides/resources--authentication--reference.md", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [cookie_params](resources--authentication--properties--cookie_params.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Authentication. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Authentication is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [oidc_auth](resources--authentication--properties--oidc_auth.md): complete subsection reference.

- [timeouts](resources--authentication--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--authentication--reference.md#schema-annotations) |
| `cookie_params` | [cookie_params](resources--authentication--properties--cookie_params.md#section) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](resources--authentication--properties--cookie_params--auth_hmac.md#section) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](resources--authentication--properties--cookie_params--auth_hmac--prim_key.md#section) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md#section) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--decryption_provider) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--location) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--store_provider) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--clear_secret_info.md#section) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](resources--authentication--properties--cookie_params--auth_hmac--prim_key--clear_secret_info.md#schema-cookie_params--auth_hmac--prim_key--clear_secret_info--provider_ref) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](resources--authentication--properties--cookie_params--auth_hmac--prim_key--clear_secret_info.md#schema-cookie_params--auth_hmac--prim_key--clear_secret_info--url) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](resources--authentication--properties--cookie_params--auth_hmac.md#schema-cookie_params--auth_hmac--prim_key_expiry) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](resources--authentication--properties--cookie_params--auth_hmac--sec_key.md#section) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md#section) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--decryption_provider) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--location) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--store_provider) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md#section) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](resources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md#schema-cookie_params--auth_hmac--sec_key--clear_secret_info--provider_ref) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](resources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md#schema-cookie_params--auth_hmac--sec_key--clear_secret_info--url) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](resources--authentication--properties--cookie_params--auth_hmac.md#schema-cookie_params--auth_hmac--sec_key_expiry) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](resources--authentication--properties--cookie_params.md#schema-cookie_params--cookie_expiry) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](resources--authentication--properties--cookie_params.md#schema-cookie_params--cookie_refresh_interval) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](resources--authentication--properties--cookie_params--kms_key_hmac.md#section) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](resources--authentication--properties--cookie_params.md#schema-cookie_params--session_expiry) |
| `description` | [description](resources--authentication--reference.md#schema-description) |
| `disable` | [disable](resources--authentication--reference.md#schema-disable) |
| `id` | [id](resources--authentication--reference.md#schema-id) |
| `labels` | [labels](resources--authentication--reference.md#schema-labels) |
| `name` | [name](resources--authentication--reference.md#schema-name) |
| `namespace` | [namespace](resources--authentication--reference.md#schema-namespace) |
| `oidc_auth` | [oidc_auth](resources--authentication--properties--oidc_auth.md#section) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](resources--authentication--properties--oidc_auth--client_secret.md#section) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md#section) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md#schema-oidc_auth--client_secret--blindfold_secret_info--decryption_provider) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md#schema-oidc_auth--client_secret--blindfold_secret_info--location) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md#schema-oidc_auth--client_secret--blindfold_secret_info--store_provider) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](resources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md#section) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](resources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md#schema-oidc_auth--client_secret--clear_secret_info--provider_ref) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](resources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md#schema-oidc_auth--client_secret--clear_secret_info--url) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](resources--authentication--properties--oidc_auth--oidc_auth_params.md#section) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](resources--authentication--properties--oidc_auth--oidc_auth_params.md#schema-oidc_auth--oidc_auth_params--auth_endpoint_url) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](resources--authentication--properties--oidc_auth--oidc_auth_params.md#schema-oidc_auth--oidc_auth_params--end_session_endpoint_url) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](resources--authentication--properties--oidc_auth--oidc_auth_params.md#schema-oidc_auth--oidc_auth_params--token_endpoint_url) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](resources--authentication--properties--oidc_auth.md#schema-oidc_auth--oidc_client_id) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](resources--authentication--properties--oidc_auth.md#schema-oidc_auth--oidc_well_known_config_url) |
| `timeouts` | [timeouts](resources--authentication--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--authentication--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--authentication--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--authentication--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--authentication--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [cookie_params](resources--authentication--properties--cookie_params.md)
- [oidc_auth](resources--authentication--properties--oidc_auth.md)
- [timeouts](resources--authentication--properties--timeouts.md)
- [xcsh_authentication](../resources/authentication.md)
