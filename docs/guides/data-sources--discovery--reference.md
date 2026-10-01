---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 32084, "body_sha256": "sha256:6eda49d401e5b78d2af9c49eb826ce728e2635f13f557b96d95d1e5878372d14", "canonical_id": "xcsh-docs:data-sources:discovery:reference", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul", "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "xcsh-docs:data-sources:discovery:properties:no_cluster_id", "xcsh-docs:data-sources:discovery:properties:where"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:reference", "parent_id": "xcsh-docs:data-sources:discovery:fundamentals", "path": "docs/guides/data-sources--discovery--reference.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
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

<a id="schema-cluster_id"></a>

### cluster_id property

Type: `"string"`. Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

Upstream description:

Exclusive with \[no\_cluster\_id\] Specify identifier for discovery cluster. This identifier can be
specified in endpoint object to discover only from this discovery object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [cluster_id](data-sources--discovery--reference.md#schema-cluster_id)
- [no_cluster_id](data-sources--discovery--properties--no_cluster_id.md#section)

Select alternatives according to the provider validators above.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Discovery.

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

- [discovery_consul](data-sources--discovery--properties--discovery_consul.md): complete subsection reference.

- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md): complete subsection reference.

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

Name of the Discovery.

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

Namespace where the Discovery exists.

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

- [no_cluster_id](data-sources--discovery--properties--no_cluster_id.md): complete subsection reference.

- [where](data-sources--discovery--properties--where.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--discovery--reference.md#schema-annotations) |
| `cluster_id` | [cluster_id](data-sources--discovery--reference.md#schema-cluster_id) |
| `description` | [description](data-sources--discovery--reference.md#schema-description) |
| `discovery_consul` | [discovery_consul](data-sources--discovery--properties--discovery_consul.md#section) |
| `discovery_consul.access_info` | [discovery_consul.access_info](data-sources--discovery--properties--discovery_consul--access_info.md#section) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info.md#section) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](data-sources--discovery--properties--discovery_consul--access_info--connection_info.md#schema-discovery_consul--access_info--connection_info--api_server) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md#section) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md#schema-discovery_consul--access_info--connection_info--tls_info--certificate) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url.md#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--decryption_provider) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--location) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--store_provider) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md#schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md#schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md#schema-discovery_consul--access_info--connection_info--tls_info--server_name) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md#schema-discovery_consul--access_info--connection_info--tls_info--trusted_ca_url) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url.md#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--decryption_provider) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--location) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info.md#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--store_provider) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info.md#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info.md#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info--provider_ref) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info.md#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info--url) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](data-sources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md#schema-discovery_consul--access_info--http_basic_auth_info--user_name) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](data-sources--discovery--properties--discovery_consul--publish_info.md#section) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](data-sources--discovery--properties--discovery_consul--publish_info--disable_spec.md#section) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](data-sources--discovery--properties--discovery_consul--publish_info--publish.md#section) |
| `discovery_k8s` | [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md#section) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](data-sources--discovery--properties--discovery_k8s--access_info.md#section) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info.md#section) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](data-sources--discovery--properties--discovery_k8s--access_info--connection_info.md#schema-discovery_k8s--access_info--connection_info--api_server) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md#section) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--certificate) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url.md#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--decryption_provider) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--location) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--store_provider) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info.md#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info--url) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--server_name) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md#schema-discovery_k8s--access_info--connection_info--tls_info--trusted_ca_url) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](data-sources--discovery--properties--discovery_k8s--access_info--isolated.md#section) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url.md#section) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info.md#section) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--decryption_provider) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--location) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info.md#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--store_provider) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--clear_secret_info.md#section) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--clear_secret_info.md#schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--provider_ref) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--clear_secret_info.md#schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--url) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](data-sources--discovery--properties--discovery_k8s--access_info--reachable.md#section) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](data-sources--discovery--properties--discovery_k8s--default_all.md#section) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](data-sources--discovery--properties--discovery_k8s--namespace_mapping.md#section) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](data-sources--discovery--properties--discovery_k8s--namespace_mapping--items.md#section) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](data-sources--discovery--properties--discovery_k8s--namespace_mapping--items.md#schema-discovery_k8s--namespace_mapping--items--namespace) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](data-sources--discovery--properties--discovery_k8s--namespace_mapping--items.md#schema-discovery_k8s--namespace_mapping--items--namespace_regex) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](data-sources--discovery--properties--discovery_k8s--publish_info.md#section) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](data-sources--discovery--properties--discovery_k8s--publish_info--disable_spec.md#section) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](data-sources--discovery--properties--discovery_k8s--publish_info--dns_delegation.md#section) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](data-sources--discovery--properties--discovery_k8s--publish_info--dns_delegation.md#schema-discovery_k8s--publish_info--dns_delegation--dns_mode) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](data-sources--discovery--properties--discovery_k8s--publish_info--dns_delegation.md#schema-discovery_k8s--publish_info--dns_delegation--subdomain) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](data-sources--discovery--properties--discovery_k8s--publish_info--publish.md#section) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](data-sources--discovery--properties--discovery_k8s--publish_info--publish.md#schema-discovery_k8s--publish_info--publish--namespace) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](data-sources--discovery--properties--discovery_k8s--publish_info--publish_fqdns.md#section) |
| `id` | [id](data-sources--discovery--reference.md#schema-id) |
| `labels` | [labels](data-sources--discovery--reference.md#schema-labels) |
| `name` | [name](data-sources--discovery--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--discovery--reference.md#schema-namespace) |
| `no_cluster_id` | [no_cluster_id](data-sources--discovery--properties--no_cluster_id.md#section) |
| `where` | [where](data-sources--discovery--properties--where.md#section) |
| `where.site` | [where.site](data-sources--discovery--properties--where--site.md#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--discovery--properties--where--site--disable_internet_vip.md#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--discovery--properties--where--site--enable_internet_vip.md#section) |
| `where.site.network_type` | [where.site.network_type](data-sources--discovery--properties--where--site.md#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](data-sources--discovery--properties--where--site--ref.md#section) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--discovery--properties--where--site--ref.md#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--discovery--properties--where--site--ref.md#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--discovery--properties--where--site--ref.md#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--discovery--properties--where--site--ref.md#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--discovery--properties--where--site--ref.md#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](data-sources--discovery--properties--where--virtual_network.md#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--discovery--properties--where--virtual_network--ref.md#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--discovery--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--discovery--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--discovery--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--discovery--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--discovery--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](data-sources--discovery--properties--where--virtual_site.md#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--discovery--properties--where--virtual_site--disable_internet_vip.md#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--discovery--properties--where--virtual_site--enable_internet_vip.md#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--discovery--properties--where--virtual_site.md#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--discovery--properties--where--virtual_site--ref.md#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--discovery--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--discovery--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--discovery--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--discovery--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--discovery--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--uid) |

## Next pages

- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- [no_cluster_id](data-sources--discovery--properties--no_cluster_id.md)
- [where](data-sources--discovery--properties--where.md)
- [xcsh_discovery](../data-sources/discovery.md)
