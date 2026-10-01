---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 25285, "body_sha256": "sha256:ecfa1848473d07fd6179703e93e47f3465dbca331417385c711a08b11e526cc4", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:reference", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "xcsh-docs:data-sources:policy_based_routing:properties:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:reference", "parent_id": "xcsh-docs:data-sources:policy_based_routing:fundamentals", "path": "docs/guides/data-sources--policy_based_routing--reference.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the PolicyBasedRouting.

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

- [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--properties--forwarding_class_list.md): complete subsection reference.

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

Name of the PolicyBasedRouting.

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

Namespace where the PolicyBasedRouting exists.

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

- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--policy_based_routing--reference.md#schema-annotations) |
| `description` | [description](data-sources--policy_based_routing--reference.md#schema-description) |
| `forward_proxy_pbr` | [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_destinations.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--all_sources.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--namespace) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--tenant) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--any_path.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_prefix_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--suffix_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--namespace) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--tenant) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--label_selector.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--label_selector--expressions) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--metadata--description_spec) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--metadata.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--metadata--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list--prefixes) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list.md#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list.md#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--suffix_value) |
| `forwarding_class_list` | [forwarding_class_list](data-sources--policy_based_routing--properties--forwarding_class_list.md#section) |
| `forwarding_class_list.name` | [forwarding_class_list.name](data-sources--policy_based_routing--properties--forwarding_class_list.md#schema-forwarding_class_list--name) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](data-sources--policy_based_routing--properties--forwarding_class_list.md#schema-forwarding_class_list--namespace) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](data-sources--policy_based_routing--properties--forwarding_class_list.md#schema-forwarding_class_list--tenant) |
| `id` | [id](data-sources--policy_based_routing--reference.md#schema-id) |
| `labels` | [labels](data-sources--policy_based_routing--reference.md#schema-labels) |
| `name` | [name](data-sources--policy_based_routing--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--policy_based_routing--reference.md#schema-namespace) |
| `network_pbr` | [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md#section) |
| `network_pbr.any` | [network_pbr.any](data-sources--policy_based_routing--properties--network_pbr--any.md#section) |
| `network_pbr.label_selector` | [network_pbr.label_selector](data-sources--policy_based_routing--properties--network_pbr--label_selector.md#section) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](data-sources--policy_based_routing--properties--network_pbr--label_selector.md#schema-network_pbr--label_selector--expressions) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules.md#section) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_tcp_traffic.md#section) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_traffic.md#section) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_udp_traffic.md#section) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--any.md#section) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--applications.md#section) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--applications.md#schema-network_pbr--network_pbr_rules--applications--applications) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules.md#schema-network_pbr--network_pbr_rules--dns_name) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md#section) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md#schema-network_pbr--network_pbr_rules--forwarding_class_list--name) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md#schema-network_pbr--network_pbr_rules--forwarding_class_list--namespace) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md#schema-network_pbr--network_pbr_rules--forwarding_class_list--tenant) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set.md#section) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#section) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--kind) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--name) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--namespace) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--tenant) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--uid) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--metadata.md#section) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--metadata.md#schema-network_pbr--network_pbr_rules--metadata--description_spec) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--metadata.md#schema-network_pbr--network_pbr_rules--metadata--name) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--prefix_list.md#section) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--prefix_list.md#schema-network_pbr--network_pbr_rules--prefix_list--prefixes) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--protocol_port_range.md#section) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--protocol_port_range.md#schema-network_pbr--network_pbr_rules--protocol_port_range--port_ranges) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--protocol_port_range.md#schema-network_pbr--network_pbr_rules--protocol_port_range--protocol) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](data-sources--policy_based_routing--properties--network_pbr--prefix_list.md#section) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](data-sources--policy_based_routing--properties--network_pbr--prefix_list.md#schema-network_pbr--prefix_list--prefixes) |

## Next pages

- [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md)
- [forwarding_class_list](data-sources--policy_based_routing--properties--forwarding_class_list.md)
- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
