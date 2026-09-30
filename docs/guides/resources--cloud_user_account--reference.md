---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 13007, "body_sha256": "sha256:0e33ef243fb9c21a3307b4d76c7a353a408072531efd569b0ac11d86595882c2", "canonical_id": "xcsh-docs:resources:cloud_user_account:reference", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider", "xcsh-docs:resources:cloud_user_account:properties:timeouts"], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:reference", "parent_id": "xcsh-docs:resources:cloud_user_account:fundamentals", "path": "docs/guides/resources--cloud_user_account--reference.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
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

- [aws_provider](resources--cloud_user_account--properties--aws_provider.md): complete subsection reference.

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

Name of the Cloud User Account. Must be unique within the namespace.

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

Namespace where the Cloud User Account is created.

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

- [timeouts](resources--cloud_user_account--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_user_account--reference.md#schema-annotations) |
| `aws_provider` | [aws_provider](resources--cloud_user_account--properties--aws_provider.md#section) |
| `aws_provider.aws_account_number` | [aws_provider.aws_account_number](resources--cloud_user_account--properties--aws_provider.md#schema-aws_provider--aws_account_number) |
| `aws_provider.aws_assume_role` | [aws_provider.aws_assume_role](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#section) |
| `aws_provider.aws_assume_role.custom_external_id` | [aws_provider.aws_assume_role.custom_external_id](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#schema-aws_provider--aws_assume_role--custom_external_id) |
| `aws_provider.aws_assume_role.duration_seconds` | [aws_provider.aws_assume_role.duration_seconds](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#schema-aws_provider--aws_assume_role--duration_seconds) |
| `aws_provider.aws_assume_role.external_id_is_optional` | [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_optional.md#section) |
| `aws_provider.aws_assume_role.external_id_is_tenant_id` | [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_tenant_id.md#section) |
| `aws_provider.aws_assume_role.role_arn` | [aws_provider.aws_assume_role.role_arn](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#schema-aws_provider--aws_assume_role--role_arn) |
| `aws_provider.aws_assume_role.session_name` | [aws_provider.aws_assume_role.session_name](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#schema-aws_provider--aws_assume_role--session_name) |
| `aws_provider.aws_assume_role.session_tags` | [aws_provider.aws_assume_role.session_tags](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md#schema-aws_provider--aws_assume_role--session_tags) |
| `aws_provider.aws_secret_key` | [aws_provider.aws_secret_key](resources--cloud_user_account--properties--aws_provider--aws_secret_key.md#section) |
| `aws_provider.aws_secret_key.access_key` | [aws_provider.aws_secret_key.access_key](resources--cloud_user_account--properties--aws_provider--aws_secret_key.md#schema-aws_provider--aws_secret_key--access_key) |
| `aws_provider.aws_secret_key.secret_key` | [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md#section) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md#section) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_provider--aws_secret_key--secret_key--blindfold_secret_info--decryption_provider) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_provider--aws_secret_key--secret_key--blindfold_secret_info--location) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_provider--aws_secret_key--secret_key--blindfold_secret_info--store_provider) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info` | [aws_provider.aws_secret_key.secret_key.clear_secret_info](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md#section) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md#schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--provider_ref) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.url](resources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md#schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--url) |
| `description` | [description](resources--cloud_user_account--reference.md#schema-description) |
| `disable` | [disable](resources--cloud_user_account--reference.md#schema-disable) |
| `id` | [id](resources--cloud_user_account--reference.md#schema-id) |
| `labels` | [labels](resources--cloud_user_account--reference.md#schema-labels) |
| `name` | [name](resources--cloud_user_account--reference.md#schema-name) |
| `namespace` | [namespace](resources--cloud_user_account--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--cloud_user_account--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--cloud_user_account--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_user_account--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--cloud_user_account--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--cloud_user_account--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [aws_provider](resources--cloud_user_account--properties--aws_provider.md)
- [timeouts](resources--cloud_user_account--properties--timeouts.md)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
