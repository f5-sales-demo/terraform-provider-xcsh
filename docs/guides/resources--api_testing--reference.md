---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 22147, "body_sha256": "sha256:f2088e0e156c1ae46d1e613af7253cef51e755bdf4a4cec373df5bc1541eb192", "canonical_id": "xcsh-docs:resources:api_testing:reference", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains", "xcsh-docs:resources:api_testing:properties:every_day", "xcsh-docs:resources:api_testing:properties:every_month", "xcsh-docs:resources:api_testing:properties:every_week", "xcsh-docs:resources:api_testing:properties:timeouts"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:reference", "parent_id": "xcsh-docs:resources:api_testing:fundamentals", "path": "docs/guides/resources--api_testing--reference.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
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

<a id="schema-custom_header_value"></a>

### custom_header_value property

Type: `"string"`. Optional, Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

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

- [domains](resources--api_testing--properties--domains.md): complete subsection reference.

- [every_day](resources--api_testing--properties--every_day.md): complete subsection reference.

- [every_month](resources--api_testing--properties--every_month.md): complete subsection reference.

- [every_week](resources--api_testing--properties--every_week.md): complete subsection reference.

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

Name of the API Testing. Must be unique within the namespace.

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

Namespace where the API Testing is created.

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

- [timeouts](resources--api_testing--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_testing--reference.md#schema-annotations) |
| `custom_header_value` | [custom_header_value](resources--api_testing--reference.md#schema-custom_header_value) |
| `description` | [description](resources--api_testing--reference.md#schema-description) |
| `disable` | [disable](resources--api_testing--reference.md#schema-disable) |
| `domains` | [domains](resources--api_testing--properties--domains.md#section) |
| `domains.allow_destructive_methods` | [domains.allow_destructive_methods](resources--api_testing--properties--domains.md#schema-domains--allow_destructive_methods) |
| `domains.credentials` | [domains.credentials](resources--api_testing--properties--domains--credentials.md#section) |
| `domains.credentials.admin` | [domains.credentials.admin](resources--api_testing--properties--domains--credentials--admin.md#section) |
| `domains.credentials.api_key` | [domains.credentials.api_key](resources--api_testing--properties--domains--credentials--api_key.md#section) |
| `domains.credentials.api_key.key` | [domains.credentials.api_key.key](resources--api_testing--properties--domains--credentials--api_key.md#schema-domains--credentials--api_key--key) |
| `domains.credentials.api_key.value` | [domains.credentials.api_key.value](resources--api_testing--properties--domains--credentials--api_key--value.md#section) |
| `domains.credentials.api_key.value.blindfold_secret_info` | [domains.credentials.api_key.value.blindfold_secret_info](resources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md#section) |
| `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` | [domains.credentials.api_key.value.blindfold_secret_info.decryption_provider](resources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md#schema-domains--credentials--api_key--value--blindfold_secret_info--decryption_provider) |
| `domains.credentials.api_key.value.blindfold_secret_info.location` | [domains.credentials.api_key.value.blindfold_secret_info.location](resources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md#schema-domains--credentials--api_key--value--blindfold_secret_info--location) |
| `domains.credentials.api_key.value.blindfold_secret_info.store_provider` | [domains.credentials.api_key.value.blindfold_secret_info.store_provider](resources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md#schema-domains--credentials--api_key--value--blindfold_secret_info--store_provider) |
| `domains.credentials.api_key.value.clear_secret_info` | [domains.credentials.api_key.value.clear_secret_info](resources--api_testing--properties--domains--credentials--api_key--value--clear_secret_info.md#section) |
| `domains.credentials.api_key.value.clear_secret_info.provider_ref` | [domains.credentials.api_key.value.clear_secret_info.provider_ref](resources--api_testing--properties--domains--credentials--api_key--value--clear_secret_info.md#schema-domains--credentials--api_key--value--clear_secret_info--provider_ref) |
| `domains.credentials.api_key.value.clear_secret_info.url` | [domains.credentials.api_key.value.clear_secret_info.url](resources--api_testing--properties--domains--credentials--api_key--value--clear_secret_info.md#schema-domains--credentials--api_key--value--clear_secret_info--url) |
| `domains.credentials.basic_auth` | [domains.credentials.basic_auth](resources--api_testing--properties--domains--credentials--basic_auth.md#section) |
| `domains.credentials.basic_auth.password` | [domains.credentials.basic_auth.password](resources--api_testing--properties--domains--credentials--basic_auth--password.md#section) |
| `domains.credentials.basic_auth.password.blindfold_secret_info` | [domains.credentials.basic_auth.password.blindfold_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md#section) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md#schema-domains--credentials--basic_auth--password--blindfold_secret_info--decryption_provider) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.location` | [domains.credentials.basic_auth.password.blindfold_secret_info.location](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md#schema-domains--credentials--basic_auth--password--blindfold_secret_info--location) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.store_provider](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md#schema-domains--credentials--basic_auth--password--blindfold_secret_info--store_provider) |
| `domains.credentials.basic_auth.password.clear_secret_info` | [domains.credentials.basic_auth.password.clear_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md#section) |
| `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` | [domains.credentials.basic_auth.password.clear_secret_info.provider_ref](resources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md#schema-domains--credentials--basic_auth--password--clear_secret_info--provider_ref) |
| `domains.credentials.basic_auth.password.clear_secret_info.url` | [domains.credentials.basic_auth.password.clear_secret_info.url](resources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md#schema-domains--credentials--basic_auth--password--clear_secret_info--url) |
| `domains.credentials.basic_auth.user` | [domains.credentials.basic_auth.user](resources--api_testing--properties--domains--credentials--basic_auth.md#schema-domains--credentials--basic_auth--user) |
| `domains.credentials.bearer_token` | [domains.credentials.bearer_token](resources--api_testing--properties--domains--credentials--bearer_token.md#section) |
| `domains.credentials.bearer_token.token` | [domains.credentials.bearer_token.token](resources--api_testing--properties--domains--credentials--bearer_token--token.md#section) |
| `domains.credentials.bearer_token.token.blindfold_secret_info` | [domains.credentials.bearer_token.token.blindfold_secret_info](resources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md#section) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider](resources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md#schema-domains--credentials--bearer_token--token--blindfold_secret_info--decryption_provider) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.location` | [domains.credentials.bearer_token.token.blindfold_secret_info.location](resources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md#schema-domains--credentials--bearer_token--token--blindfold_secret_info--location) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.store_provider](resources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md#schema-domains--credentials--bearer_token--token--blindfold_secret_info--store_provider) |
| `domains.credentials.bearer_token.token.clear_secret_info` | [domains.credentials.bearer_token.token.clear_secret_info](resources--api_testing--properties--domains--credentials--bearer_token--token--clear_secret_info.md#section) |
| `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` | [domains.credentials.bearer_token.token.clear_secret_info.provider_ref](resources--api_testing--properties--domains--credentials--bearer_token--token--clear_secret_info.md#schema-domains--credentials--bearer_token--token--clear_secret_info--provider_ref) |
| `domains.credentials.bearer_token.token.clear_secret_info.url` | [domains.credentials.bearer_token.token.clear_secret_info.url](resources--api_testing--properties--domains--credentials--bearer_token--token--clear_secret_info.md#schema-domains--credentials--bearer_token--token--clear_secret_info--url) |
| `domains.credentials.credential_name` | [domains.credentials.credential_name](resources--api_testing--properties--domains--credentials.md#schema-domains--credentials--credential_name) |
| `domains.credentials.login_endpoint` | [domains.credentials.login_endpoint](resources--api_testing--properties--domains--credentials--login_endpoint.md#section) |
| `domains.credentials.login_endpoint.json_payload` | [domains.credentials.login_endpoint.json_payload](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload.md#section) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md#section) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md#schema-domains--credentials--login_endpoint--json_payload--blindfold_secret_info--decryption_provider) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md#schema-domains--credentials--login_endpoint--json_payload--blindfold_secret_info--location) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md#schema-domains--credentials--login_endpoint--json_payload--blindfold_secret_info--store_provider) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info` | [domains.credentials.login_endpoint.json_payload.clear_secret_info](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--clear_secret_info.md#section) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--clear_secret_info.md#schema-domains--credentials--login_endpoint--json_payload--clear_secret_info--provider_ref) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.url](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload--clear_secret_info.md#schema-domains--credentials--login_endpoint--json_payload--clear_secret_info--url) |
| `domains.credentials.login_endpoint.method` | [domains.credentials.login_endpoint.method](resources--api_testing--properties--domains--credentials--login_endpoint.md#schema-domains--credentials--login_endpoint--method) |
| `domains.credentials.login_endpoint.path` | [domains.credentials.login_endpoint.path](resources--api_testing--properties--domains--credentials--login_endpoint.md#schema-domains--credentials--login_endpoint--path) |
| `domains.credentials.login_endpoint.token_response_key` | [domains.credentials.login_endpoint.token_response_key](resources--api_testing--properties--domains--credentials--login_endpoint.md#schema-domains--credentials--login_endpoint--token_response_key) |
| `domains.credentials.standard` | [domains.credentials.standard](resources--api_testing--properties--domains--credentials--standard.md#section) |
| `domains.domain` | [domains.domain](resources--api_testing--properties--domains.md#schema-domains--domain) |
| `every_day` | [every_day](resources--api_testing--properties--every_day.md#section) |
| `every_month` | [every_month](resources--api_testing--properties--every_month.md#section) |
| `every_week` | [every_week](resources--api_testing--properties--every_week.md#section) |
| `id` | [id](resources--api_testing--reference.md#schema-id) |
| `labels` | [labels](resources--api_testing--reference.md#schema-labels) |
| `name` | [name](resources--api_testing--reference.md#schema-name) |
| `namespace` | [namespace](resources--api_testing--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--api_testing--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--api_testing--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--api_testing--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--api_testing--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--api_testing--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [domains](resources--api_testing--properties--domains.md)
- [every_day](resources--api_testing--properties--every_day.md)
- [every_month](resources--api_testing--properties--every_month.md)
- [every_week](resources--api_testing--properties--every_week.md)
- [timeouts](resources--api_testing--properties--timeouts.md)
- [xcsh_api_testing](../resources/api_testing.md)
