---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 30330, "body_sha256": "sha256:af6c77972be1b45cf6f6de4b3313e8f2ef8a10b3546d888e08c56dda9a0009c3", "canonical_id": "xcsh-docs:resources:code_base_integration:reference", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration", "xcsh-docs:resources:code_base_integration:properties:timeouts"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:reference", "parent_id": "xcsh-docs:resources:code_base_integration:fundamentals", "path": "docs/guides/resources--code_base_integration--reference.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
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

- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md): complete subsection reference.

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

Name of the Code Base Integration. Must be unique within the namespace.

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

Namespace where the Code Base Integration is created.

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

- [timeouts](resources--code_base_integration--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--code_base_integration--reference.md#schema-annotations) |
| `code_base_integration` | [code_base_integration](resources--code_base_integration--properties--code_base_integration.md#section) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](resources--code_base_integration--properties--code_base_integration--azure_repos.md#section) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token.md#section) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md#section) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--location) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--clear_secret_info.md#section) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--clear_secret_info.md#schema-code_base_integration--azure_repos--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--azure_repos--access_token--clear_secret_info.md#schema-code_base_integration--azure_repos--access_token--clear_secret_info--url) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](resources--code_base_integration--properties--code_base_integration--bitbucket.md#section) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd.md#section) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--blindfold_secret_info.md#section) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--decryption_provider) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--location) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--store_provider) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--clear_secret_info.md#section) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--clear_secret_info.md#schema-code_base_integration--bitbucket--passwd--clear_secret_info--provider_ref) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--bitbucket--passwd--clear_secret_info.md#schema-code_base_integration--bitbucket--passwd--clear_secret_info--url) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](resources--code_base_integration--properties--code_base_integration--bitbucket.md#schema-code_base_integration--bitbucket--username) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md#section) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd.md#section) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md#section) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--decryption_provider) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--location) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--store_provider) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--clear_secret_info.md#section) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--clear_secret_info.md#schema-code_base_integration--bitbucket_server--passwd--clear_secret_info--provider_ref) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--clear_secret_info.md#schema-code_base_integration--bitbucket_server--passwd--clear_secret_info--url) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md#schema-code_base_integration--bitbucket_server--url) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md#schema-code_base_integration--bitbucket_server--username) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md#schema-code_base_integration--bitbucket_server--verify_ssl) |
| `code_base_integration.github` | [code_base_integration.github](resources--code_base_integration--properties--code_base_integration--github.md#section) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](resources--code_base_integration--properties--code_base_integration--github--access_token.md#section) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--github--access_token--blindfold_secret_info.md#section) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--github--access_token--blindfold_secret_info.md#schema-code_base_integration--github--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--github--access_token--blindfold_secret_info.md#schema-code_base_integration--github--access_token--blindfold_secret_info--location) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--github--access_token--blindfold_secret_info.md#schema-code_base_integration--github--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--github--access_token--clear_secret_info.md#section) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--github--access_token--clear_secret_info.md#schema-code_base_integration--github--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--github--access_token--clear_secret_info.md#schema-code_base_integration--github--access_token--clear_secret_info--url) |
| `code_base_integration.github.username` | [code_base_integration.github.username](resources--code_base_integration--properties--code_base_integration--github.md#schema-code_base_integration--github--username) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](resources--code_base_integration--properties--code_base_integration--github.md#schema-code_base_integration--github--verify_ssl) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](resources--code_base_integration--properties--code_base_integration--github_enterprise.md#section) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token.md#section) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md#section) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--location) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--clear_secret_info.md#section) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--clear_secret_info.md#schema-code_base_integration--github_enterprise--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--clear_secret_info.md#schema-code_base_integration--github_enterprise--access_token--clear_secret_info--url) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](resources--code_base_integration--properties--code_base_integration--github_enterprise.md#schema-code_base_integration--github_enterprise--hostname) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](resources--code_base_integration--properties--code_base_integration--github_enterprise.md#schema-code_base_integration--github_enterprise--username) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](resources--code_base_integration--properties--code_base_integration--gitlab.md#section) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](resources--code_base_integration--properties--code_base_integration--gitlab--access_token.md#section) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md#section) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--location) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--clear_secret_info.md#section) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--clear_secret_info.md#schema-code_base_integration--gitlab--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--gitlab--access_token--clear_secret_info.md#schema-code_base_integration--gitlab--access_token--clear_secret_info--url) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md#section) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token.md#section) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info.md#section) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--location) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info.md#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--clear_secret_info.md#section) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--clear_secret_info.md#schema-code_base_integration--gitlab_enterprise--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token--clear_secret_info.md#schema-code_base_integration--gitlab_enterprise--access_token--clear_secret_info--url) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md#schema-code_base_integration--gitlab_enterprise--url) |
| `description` | [description](resources--code_base_integration--reference.md#schema-description) |
| `disable` | [disable](resources--code_base_integration--reference.md#schema-disable) |
| `id` | [id](resources--code_base_integration--reference.md#schema-id) |
| `labels` | [labels](resources--code_base_integration--reference.md#schema-labels) |
| `name` | [name](resources--code_base_integration--reference.md#schema-name) |
| `namespace` | [namespace](resources--code_base_integration--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--code_base_integration--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--code_base_integration--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--code_base_integration--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--code_base_integration--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--code_base_integration--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- [timeouts](resources--code_base_integration--properties--timeouts.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
