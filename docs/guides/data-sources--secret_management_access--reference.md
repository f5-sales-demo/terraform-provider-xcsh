---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 42802, "body_sha256": "sha256:c786d94cb7cce4cbdc1303e9f761de17dcff92b60f75797f9077bee7db91108c", "canonical_id": "xcsh-docs:data-sources:secret_management_access:reference", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info", "xcsh-docs:data-sources:secret_management_access:properties:where"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:reference", "parent_id": "xcsh-docs:data-sources:secret_management_access:fundamentals", "path": "docs/guides/data-sources--secret_management_access--reference.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- Property reference

## Direct properties

- [access_info](data-sources--secret_management_access--properties--access_info.md): complete subsection reference.

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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the SecretManagementAccess.

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

Name of the SecretManagementAccess.

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

Namespace where the SecretManagementAccess exists.

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

<a id="schema-provider_name"></a>

### provider_name property

Type: `"string"`. Computed.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [where](data-sources--secret_management_access--properties--where.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](data-sources--secret_management_access--properties--access_info.md#section) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](data-sources--secret_management_access--properties--access_info--rest_auth_info.md#section) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md#section) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password.md#section) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--blindfold_secret_info.md#section) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--blindfold_secret_info.md#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--decryption_provider) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--blindfold_secret_info.md#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--location) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--blindfold_secret_info.md#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--store_provider) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--clear_secret_info.md#section) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--clear_secret_info.md#schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--provider_ref) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth--password--clear_secret_info.md#schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--url) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md#schema-access_info--rest_auth_info--basic_auth--username) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--headers_auth.md#section) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](data-sources--secret_management_access--properties--access_info--rest_auth_info--headers_auth--headers.md#section) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md#section) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth--query_params.md#section) |
| `access_info.scheme` | [access_info.scheme](data-sources--secret_management_access--properties--access_info.md#schema-access_info--scheme) |
| `access_info.server_endpoint` | [access_info.server_endpoint](data-sources--secret_management_access--properties--access_info.md#schema-access_info--server_endpoint) |
| `access_info.tls_config` | [access_info.tls_config](data-sources--secret_management_access--properties--access_info--tls_config.md#section) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](data-sources--secret_management_access--properties--access_info--tls_config--cert_params.md#section) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#section) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#schema-access_info--tls_config--cert_params--certificates--kind) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#schema-access_info--tls_config--cert_params--certificates--name) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#schema-access_info--tls_config--cert_params--certificates--namespace) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#schema-access_info--tls_config--cert_params--certificates--tenant) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--certificates.md#schema-access_info--tls_config--cert_params--certificates--uid) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](data-sources--secret_management_access--properties--access_info--tls_config--cert_params.md#schema-access_info--tls_config--cert_params--cipher_suites) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](data-sources--secret_management_access--properties--access_info--tls_config--cert_params.md#schema-access_info--tls_config--cert_params--maximum_protocol_version) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](data-sources--secret_management_access--properties--access_info--tls_config--cert_params.md#schema-access_info--tls_config--cert_params--minimum_protocol_version) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--skip_server_verification.md#section) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md#section) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md#schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca.md#section) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#section) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--kind) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--name) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--namespace) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--tenant) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--uid) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md#schema-access_info--tls_config--cert_params--tls_validation_params--verify_subject_alt_names) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](data-sources--secret_management_access--properties--access_info--tls_config--cert_params--volterra_trusted_ca.md#section) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md#section) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md#schema-access_info--tls_config--common_params--cipher_suites) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md#schema-access_info--tls_config--common_params--maximum_protocol_version) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md#schema-access_info--tls_config--common_params--minimum_protocol_version) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md#section) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md#schema-access_info--tls_config--common_params--tls_certificates--certificate_url) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms.md#section) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms.md#schema-access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md#schema-access_info--tls_config--common_params--tls_certificates--description_spec) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--disable_ocsp_stapling.md#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md#schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md#schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--use_system_defaults.md#section) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params.md#section) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params.md#schema-access_info--tls_config--common_params--validation_params--skip_hostname_verification) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca.md#section) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#section) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params.md#schema-access_info--tls_config--common_params--validation_params--trusted_ca_url) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](data-sources--secret_management_access--properties--access_info--tls_config--common_params--validation_params.md#schema-access_info--tls_config--common_params--validation_params--verify_subject_alt_names) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](data-sources--secret_management_access--properties--access_info--tls_config--default_session_key_caching.md#section) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](data-sources--secret_management_access--properties--access_info--tls_config--disable_session_key_caching.md#section) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](data-sources--secret_management_access--properties--access_info--tls_config--disable_sni.md#section) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](data-sources--secret_management_access--properties--access_info--tls_config.md#schema-access_info--tls_config--max_session_keys) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](data-sources--secret_management_access--properties--access_info--tls_config.md#schema-access_info--tls_config--sni) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](data-sources--secret_management_access--properties--access_info--tls_config--use_host_header_as_sni.md#section) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](data-sources--secret_management_access--properties--access_info--vault_auth_info.md#section) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md#section) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md#schema-access_info--vault_auth_info--app_role_auth--role_id) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--decryption_provider) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--location) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--store_provider) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info.md#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info.md#schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--provider_ref) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info.md#schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--url) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](data-sources--secret_management_access--properties--access_info--vault_auth_info--token.md#section) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--blindfold_secret_info.md#section) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--blindfold_secret_info.md#schema-access_info--vault_auth_info--token--blindfold_secret_info--decryption_provider) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--blindfold_secret_info.md#schema-access_info--vault_auth_info--token--blindfold_secret_info--location) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--blindfold_secret_info.md#schema-access_info--vault_auth_info--token--blindfold_secret_info--store_provider) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--clear_secret_info.md#section) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--clear_secret_info.md#schema-access_info--vault_auth_info--token--clear_secret_info--provider_ref) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](data-sources--secret_management_access--properties--access_info--vault_auth_info--token--clear_secret_info.md#schema-access_info--vault_auth_info--token--clear_secret_info--url) |
| `annotations` | [annotations](data-sources--secret_management_access--reference.md#schema-annotations) |
| `description` | [description](data-sources--secret_management_access--reference.md#schema-description) |
| `id` | [id](data-sources--secret_management_access--reference.md#schema-id) |
| `labels` | [labels](data-sources--secret_management_access--reference.md#schema-labels) |
| `name` | [name](data-sources--secret_management_access--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--secret_management_access--reference.md#schema-namespace) |
| `provider_name` | [provider_name](data-sources--secret_management_access--reference.md#schema-provider_name) |
| `where` | [where](data-sources--secret_management_access--properties--where.md#section) |
| `where.site` | [where.site](data-sources--secret_management_access--properties--where--site.md#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--secret_management_access--properties--where--site--disable_internet_vip.md#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--secret_management_access--properties--where--site--enable_internet_vip.md#section) |
| `where.site.network_type` | [where.site.network_type](data-sources--secret_management_access--properties--where--site.md#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](data-sources--secret_management_access--properties--where--site--ref.md#section) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--secret_management_access--properties--where--site--ref.md#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--secret_management_access--properties--where--site--ref.md#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--secret_management_access--properties--where--site--ref.md#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--secret_management_access--properties--where--site--ref.md#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--secret_management_access--properties--where--site--ref.md#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](data-sources--secret_management_access--properties--where--virtual_network.md#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--secret_management_access--properties--where--virtual_network--ref.md#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--secret_management_access--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--secret_management_access--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--secret_management_access--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--secret_management_access--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--secret_management_access--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](data-sources--secret_management_access--properties--where--virtual_site.md#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--secret_management_access--properties--where--virtual_site--disable_internet_vip.md#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--secret_management_access--properties--where--virtual_site--enable_internet_vip.md#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--secret_management_access--properties--where--virtual_site.md#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--secret_management_access--properties--where--virtual_site--ref.md#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--secret_management_access--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--secret_management_access--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--secret_management_access--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--secret_management_access--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--secret_management_access--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--uid) |

## Next pages

- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [where](data-sources--secret_management_access--properties--where.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
