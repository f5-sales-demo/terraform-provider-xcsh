---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 16073, "body_sha256": "sha256:20d8e47fcd5dbf2dc906439705d16e2cbc4802926a2bc2c750fec88983dc5fe5", "child_ids": ["xcsh-docs:resources:cminstance:properties:api_token", "xcsh-docs:resources:cminstance:properties:ip", "xcsh-docs:resources:cminstance:properties:password", "xcsh-docs:resources:cminstance:properties:timeouts"], "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:reference", "parent_id": "xcsh-docs:resources:cminstance:fundamentals", "path": "documentation/resources/cminstance/properties/index.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
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

- [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/): complete subsection reference.

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

- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/ip/): complete subsection reference.

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

Name of the Cminstance. Must be unique within the namespace.

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

Namespace where the Cminstance is created.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/): complete subsection reference.

<a id="schema-port"></a>

### port property

Type: `"number"`. Required.

Port of the Central Manager instance to connect to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/): complete subsection reference.

<a id="schema-username"></a>

### username property

Type: `"string"`. Required.

Username for the Central Manager instance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-annotations) |
| `api_token` | [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/#section) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/blindfold_secret_info/#section) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--decryption_provider) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--location) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--store_provider) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/clear_secret_info/#section) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/clear_secret_info/#schema-api_token--clear_secret_info--provider_ref) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/clear_secret_info/#schema-api_token--clear_secret_info--url) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/ip/#section) |
| `ip.addr` | [ip.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/ip/#schema-ip--addr) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-namespace) |
| `password` | [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/#section) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/#section) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--decryption_provider) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--location) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--store_provider) |
| `password.clear_secret_info` | [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/#section) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/#schema-password--clear_secret_info--provider_ref) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/#schema-password--clear_secret_info--url) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-port) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/#schema-timeouts--update) |
| `username` | [username](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/#schema-username) |

## Next pages

- [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/api_token/)
- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/ip/)
- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/timeouts/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
