---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 19475, "body_sha256": "sha256:a0a331428d05b8198f27601c129eeedce58d20dc3e09efcba67f38cbf64830bb", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:reference", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key", "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:reference", "parent_id": "xcsh-docs:data-sources:cloud_credentials:fundamentals", "path": "docs/guides/data-sources--cloud_credentials--reference.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
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

- [aws_assume_role](data-sources--cloud_credentials--properties--aws_assume_role.md): complete subsection reference.

- [aws_secret_key](data-sources--cloud_credentials--properties--aws_secret_key.md): complete subsection reference.

- [azure_client_secret](data-sources--cloud_credentials--properties--azure_client_secret.md): complete subsection reference.

- [azure_pfx_certificate](data-sources--cloud_credentials--properties--azure_pfx_certificate.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the CloudCredentials.

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

- [gcp_cred_file](data-sources--cloud_credentials--properties--gcp_cred_file.md): complete subsection reference.

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

Name of the CloudCredentials.

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

Namespace where the CloudCredentials exists.

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
| `annotations` | [annotations](data-sources--cloud_credentials--reference.md#schema-annotations) |
| `aws_assume_role` | [aws_assume_role](data-sources--cloud_credentials--properties--aws_assume_role.md#section) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](data-sources--cloud_credentials--properties--aws_assume_role.md#schema-aws_assume_role--custom_external_id) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](data-sources--cloud_credentials--properties--aws_assume_role.md#schema-aws_assume_role--duration_seconds) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](data-sources--cloud_credentials--properties--aws_assume_role--external_id_is_optional.md#section) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](data-sources--cloud_credentials--properties--aws_assume_role--external_id_is_tenant_id.md#section) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](data-sources--cloud_credentials--properties--aws_assume_role.md#schema-aws_assume_role--role_arn) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](data-sources--cloud_credentials--properties--aws_assume_role.md#schema-aws_assume_role--session_name) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](data-sources--cloud_credentials--properties--aws_assume_role.md#schema-aws_assume_role--session_tags) |
| `aws_secret_key` | [aws_secret_key](data-sources--cloud_credentials--properties--aws_secret_key.md#section) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](data-sources--cloud_credentials--properties--aws_secret_key.md#schema-aws_secret_key--access_key) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](data-sources--cloud_credentials--properties--aws_secret_key--secret_key.md#section) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md#section) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_secret_key--secret_key--blindfold_secret_info--decryption_provider) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_secret_key--secret_key--blindfold_secret_info--location) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--blindfold_secret_info.md#schema-aws_secret_key--secret_key--blindfold_secret_info--store_provider) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--clear_secret_info.md#section) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--clear_secret_info.md#schema-aws_secret_key--secret_key--clear_secret_info--provider_ref) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](data-sources--cloud_credentials--properties--aws_secret_key--secret_key--clear_secret_info.md#schema-aws_secret_key--secret_key--clear_secret_info--url) |
| `azure_client_secret` | [azure_client_secret](data-sources--cloud_credentials--properties--azure_client_secret.md#section) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](data-sources--cloud_credentials--properties--azure_client_secret.md#schema-azure_client_secret--client_id) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](data-sources--cloud_credentials--properties--azure_client_secret--client_secret.md#section) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--blindfold_secret_info.md#section) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--blindfold_secret_info.md#schema-azure_client_secret--client_secret--blindfold_secret_info--decryption_provider) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--blindfold_secret_info.md#schema-azure_client_secret--client_secret--blindfold_secret_info--location) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--blindfold_secret_info.md#schema-azure_client_secret--client_secret--blindfold_secret_info--store_provider) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--clear_secret_info.md#section) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--clear_secret_info.md#schema-azure_client_secret--client_secret--clear_secret_info--provider_ref) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](data-sources--cloud_credentials--properties--azure_client_secret--client_secret--clear_secret_info.md#schema-azure_client_secret--client_secret--clear_secret_info--url) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](data-sources--cloud_credentials--properties--azure_client_secret.md#schema-azure_client_secret--subscription_id) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](data-sources--cloud_credentials--properties--azure_client_secret.md#schema-azure_client_secret--tenant_id) |
| `azure_pfx_certificate` | [azure_pfx_certificate](data-sources--cloud_credentials--properties--azure_pfx_certificate.md#section) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](data-sources--cloud_credentials--properties--azure_pfx_certificate.md#schema-azure_pfx_certificate--certificate_url) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](data-sources--cloud_credentials--properties--azure_pfx_certificate.md#schema-azure_pfx_certificate--client_id) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](data-sources--cloud_credentials--properties--azure_pfx_certificate--password.md#section) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md#section) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md#schema-azure_pfx_certificate--password--blindfold_secret_info--decryption_provider) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md#schema-azure_pfx_certificate--password--blindfold_secret_info--location) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md#schema-azure_pfx_certificate--password--blindfold_secret_info--store_provider) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md#section) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md#schema-azure_pfx_certificate--password--clear_secret_info--provider_ref) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](data-sources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md#schema-azure_pfx_certificate--password--clear_secret_info--url) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](data-sources--cloud_credentials--properties--azure_pfx_certificate.md#schema-azure_pfx_certificate--subscription_id) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](data-sources--cloud_credentials--properties--azure_pfx_certificate.md#schema-azure_pfx_certificate--tenant_id) |
| `description` | [description](data-sources--cloud_credentials--reference.md#schema-description) |
| `gcp_cred_file` | [gcp_cred_file](data-sources--cloud_credentials--properties--gcp_cred_file.md#section) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file.md#section) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md#section) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md#schema-gcp_cred_file--credential_file--blindfold_secret_info--decryption_provider) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md#schema-gcp_cred_file--credential_file--blindfold_secret_info--location) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--blindfold_secret_info.md#schema-gcp_cred_file--credential_file--blindfold_secret_info--store_provider) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md#section) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md#schema-gcp_cred_file--credential_file--clear_secret_info--provider_ref) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file--clear_secret_info.md#schema-gcp_cred_file--credential_file--clear_secret_info--url) |
| `id` | [id](data-sources--cloud_credentials--reference.md#schema-id) |
| `labels` | [labels](data-sources--cloud_credentials--reference.md#schema-labels) |
| `name` | [name](data-sources--cloud_credentials--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cloud_credentials--reference.md#schema-namespace) |

## Next pages

- [aws_assume_role](data-sources--cloud_credentials--properties--aws_assume_role.md)
- [aws_secret_key](data-sources--cloud_credentials--properties--aws_secret_key.md)
- [azure_client_secret](data-sources--cloud_credentials--properties--azure_client_secret.md)
- [azure_pfx_certificate](data-sources--cloud_credentials--properties--azure_pfx_certificate.md)
- [gcp_cred_file](data-sources--cloud_credentials--properties--gcp_cred_file.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
