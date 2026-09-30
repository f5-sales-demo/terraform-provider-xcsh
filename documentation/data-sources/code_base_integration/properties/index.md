---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 32877, "body_sha256": "sha256:854bef9239fe736807422ff6d534100ab2cb77c7991b069d8f9047bf60d5e762", "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration"], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:reference", "parent_id": "xcsh-docs:data-sources:code_base_integration:fundamentals", "path": "documentation/data-sources/code_base_integration/properties/index.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the CodeBaseIntegration.

Upstream description:

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the CodeBaseIntegration.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the CodeBaseIntegration exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-annotations) |
| `code_base_integration` | [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/#section) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/#section) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/#section) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/blindfold_secret_info/#section) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/blindfold_secret_info/#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/blindfold_secret_info/#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--location) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/blindfold_secret_info/#schema-code_base_integration--azure_repos--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/clear_secret_info/#section) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/clear_secret_info/#schema-code_base_integration--azure_repos--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/clear_secret_info/#schema-code_base_integration--azure_repos--access_token--clear_secret_info--url) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/#section) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/#section) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/#section) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--decryption_provider) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--location) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--store_provider) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/clear_secret_info/#section) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/clear_secret_info/#schema-code_base_integration--bitbucket--passwd--clear_secret_info--provider_ref) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/clear_secret_info/#schema-code_base_integration--bitbucket--passwd--clear_secret_info--url) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/#schema-code_base_integration--bitbucket--username) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/#section) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/#section) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/blindfold_secret_info/#section) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--decryption_provider) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--location) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/blindfold_secret_info/#schema-code_base_integration--bitbucket_server--passwd--blindfold_secret_info--store_provider) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/clear_secret_info/#section) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/clear_secret_info/#schema-code_base_integration--bitbucket_server--passwd--clear_secret_info--provider_ref) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/clear_secret_info/#schema-code_base_integration--bitbucket_server--passwd--clear_secret_info--url) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/#schema-code_base_integration--bitbucket_server--url) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/#schema-code_base_integration--bitbucket_server--username) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/#schema-code_base_integration--bitbucket_server--verify_ssl) |
| `code_base_integration.github` | [code_base_integration.github](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/#section) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/#section) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/#section) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/#schema-code_base_integration--github--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/#schema-code_base_integration--github--access_token--blindfold_secret_info--location) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/#schema-code_base_integration--github--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/clear_secret_info/#section) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/clear_secret_info/#schema-code_base_integration--github--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/clear_secret_info/#schema-code_base_integration--github--access_token--clear_secret_info--url) |
| `code_base_integration.github.username` | [code_base_integration.github.username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/#schema-code_base_integration--github--username) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/#schema-code_base_integration--github--verify_ssl) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/#section) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/#section) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/#section) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--location) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--github_enterprise--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/clear_secret_info/#section) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/clear_secret_info/#schema-code_base_integration--github_enterprise--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/clear_secret_info/#schema-code_base_integration--github_enterprise--access_token--clear_secret_info--url) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/#schema-code_base_integration--github_enterprise--hostname) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/#schema-code_base_integration--github_enterprise--username) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/#section) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/#section) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/blindfold_secret_info/#section) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--location) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/clear_secret_info/#section) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/clear_secret_info/#schema-code_base_integration--gitlab--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/clear_secret_info/#schema-code_base_integration--gitlab--access_token--clear_secret_info--url) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/#section) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/#section) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/#section) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--decryption_provider) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--location) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/#schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--store_provider) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/#section) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/#schema-code_base_integration--gitlab_enterprise--access_token--clear_secret_info--provider_ref) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/#schema-code_base_integration--gitlab_enterprise--access_token--clear_secret_info--url) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/#schema-code_base_integration--gitlab_enterprise--url) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/#schema-namespace) |

## Next pages

- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
