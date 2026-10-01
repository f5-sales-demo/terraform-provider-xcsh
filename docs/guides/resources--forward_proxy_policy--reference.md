---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 27768, "body_sha256": "sha256:5c2e0522881eeca850bca5b1efafe47aba146dc4233708445f2f0735c2e39593", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:reference", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:allow_all", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "xcsh-docs:resources:forward_proxy_policy:properties:any_proxy", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "xcsh-docs:resources:forward_proxy_policy:properties:drp_http_connect", "xcsh-docs:resources:forward_proxy_policy:properties:network_connector", "xcsh-docs:resources:forward_proxy_policy:properties:proxy_label_selector", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "xcsh-docs:resources:forward_proxy_policy:properties:timeouts"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:reference", "parent_id": "xcsh-docs:resources:forward_proxy_policy:fundamentals", "path": "docs/guides/resources--forward_proxy_policy--reference.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- Property reference

## Direct properties

- [allow_all](resources--forward_proxy_policy--properties--allow_all.md): complete subsection reference.

- [allow_list](resources--forward_proxy_policy--properties--allow_list.md): complete subsection reference.

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

- [any_proxy](resources--forward_proxy_policy--properties--any_proxy.md): complete subsection reference.

- [deny_list](resources--forward_proxy_policy--properties--deny_list.md): complete subsection reference.

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

- [drp_http_connect](resources--forward_proxy_policy--properties--drp_http_connect.md): complete subsection reference.

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

Name of the Forward Proxy Policy. Must be unique within the namespace.

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

Namespace where the Forward Proxy Policy is created.

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

- [network_connector](resources--forward_proxy_policy--properties--network_connector.md): complete subsection reference.

- [proxy_label_selector](resources--forward_proxy_policy--properties--proxy_label_selector.md): complete subsection reference.

- [rule_list](resources--forward_proxy_policy--properties--rule_list.md): complete subsection reference.

- [timeouts](resources--forward_proxy_policy--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--forward_proxy_policy--properties--allow_all.md#section) |
| `allow_list` | [allow_list](resources--forward_proxy_policy--properties--allow_list.md#section) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](resources--forward_proxy_policy--properties--allow_list--default_action_allow.md#section) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](resources--forward_proxy_policy--properties--allow_list--default_action_deny.md#section) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](resources--forward_proxy_policy--properties--allow_list--default_action_next_policy.md#section) |
| `allow_list.dest_list` | [allow_list.dest_list](resources--forward_proxy_policy--properties--allow_list--dest_list.md#section) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--properties--allow_list--dest_list.md#schema-allow_list--dest_list--ipv6_prefixes) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](resources--forward_proxy_policy--properties--allow_list--dest_list.md#schema-allow_list--dest_list--port_ranges) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](resources--forward_proxy_policy--properties--allow_list--dest_list.md#schema-allow_list--dest_list--prefixes) |
| `allow_list.http_list` | [allow_list.http_list](resources--forward_proxy_policy--properties--allow_list--http_list.md#section) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](resources--forward_proxy_policy--properties--allow_list--http_list--any_path.md#section) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--exact_value) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--path_exact_value) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--path_prefix_value) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--path_regex_value) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--regex_value) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](resources--forward_proxy_policy--properties--allow_list--http_list.md#schema-allow_list--http_list--suffix_value) |
| `allow_list.tls_list` | [allow_list.tls_list](resources--forward_proxy_policy--properties--allow_list--tls_list.md#section) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](resources--forward_proxy_policy--properties--allow_list--tls_list.md#schema-allow_list--tls_list--exact_value) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](resources--forward_proxy_policy--properties--allow_list--tls_list.md#schema-allow_list--tls_list--regex_value) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](resources--forward_proxy_policy--properties--allow_list--tls_list.md#schema-allow_list--tls_list--suffix_value) |
| `annotations` | [annotations](resources--forward_proxy_policy--reference.md#schema-annotations) |
| `any_proxy` | [any_proxy](resources--forward_proxy_policy--properties--any_proxy.md#section) |
| `deny_list` | [deny_list](resources--forward_proxy_policy--properties--deny_list.md#section) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](resources--forward_proxy_policy--properties--deny_list--default_action_allow.md#section) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](resources--forward_proxy_policy--properties--deny_list--default_action_deny.md#section) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](resources--forward_proxy_policy--properties--deny_list--default_action_next_policy.md#section) |
| `deny_list.dest_list` | [deny_list.dest_list](resources--forward_proxy_policy--properties--deny_list--dest_list.md#section) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--properties--deny_list--dest_list.md#schema-deny_list--dest_list--ipv6_prefixes) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](resources--forward_proxy_policy--properties--deny_list--dest_list.md#schema-deny_list--dest_list--port_ranges) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](resources--forward_proxy_policy--properties--deny_list--dest_list.md#schema-deny_list--dest_list--prefixes) |
| `deny_list.http_list` | [deny_list.http_list](resources--forward_proxy_policy--properties--deny_list--http_list.md#section) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](resources--forward_proxy_policy--properties--deny_list--http_list--any_path.md#section) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--exact_value) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--path_exact_value) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--path_prefix_value) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--path_regex_value) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--regex_value) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](resources--forward_proxy_policy--properties--deny_list--http_list.md#schema-deny_list--http_list--suffix_value) |
| `deny_list.tls_list` | [deny_list.tls_list](resources--forward_proxy_policy--properties--deny_list--tls_list.md#section) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](resources--forward_proxy_policy--properties--deny_list--tls_list.md#schema-deny_list--tls_list--exact_value) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](resources--forward_proxy_policy--properties--deny_list--tls_list.md#schema-deny_list--tls_list--regex_value) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](resources--forward_proxy_policy--properties--deny_list--tls_list.md#schema-deny_list--tls_list--suffix_value) |
| `description` | [description](resources--forward_proxy_policy--reference.md#schema-description) |
| `disable` | [disable](resources--forward_proxy_policy--reference.md#schema-disable) |
| `drp_http_connect` | [drp_http_connect](resources--forward_proxy_policy--properties--drp_http_connect.md#section) |
| `id` | [id](resources--forward_proxy_policy--reference.md#schema-id) |
| `labels` | [labels](resources--forward_proxy_policy--reference.md#schema-labels) |
| `name` | [name](resources--forward_proxy_policy--reference.md#schema-name) |
| `namespace` | [namespace](resources--forward_proxy_policy--reference.md#schema-namespace) |
| `network_connector` | [network_connector](resources--forward_proxy_policy--properties--network_connector.md#section) |
| `network_connector.name` | [network_connector.name](resources--forward_proxy_policy--properties--network_connector.md#schema-network_connector--name) |
| `network_connector.namespace` | [network_connector.namespace](resources--forward_proxy_policy--properties--network_connector.md#schema-network_connector--namespace) |
| `network_connector.tenant` | [network_connector.tenant](resources--forward_proxy_policy--properties--network_connector.md#schema-network_connector--tenant) |
| `proxy_label_selector` | [proxy_label_selector](resources--forward_proxy_policy--properties--proxy_label_selector.md#section) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](resources--forward_proxy_policy--properties--proxy_label_selector.md#schema-proxy_label_selector--expressions) |
| `rule_list` | [rule_list](resources--forward_proxy_policy--properties--rule_list.md#section) |
| `rule_list.rules` | [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md#section) |
| `rule_list.rules.action` | [rule_list.rules.action](resources--forward_proxy_policy--properties--rule_list--rules.md#schema-rule_list--rules--action) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--forward_proxy_policy--properties--rule_list--rules--all_destinations.md#section) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--forward_proxy_policy--properties--rule_list--rules--all_sources.md#section) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_list.md#section) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_list.md#schema-rule_list--rules--dst_asn_list--as_numbers) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md#section) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md#schema-rule_list--rules--dst_asn_set--name) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md#schema-rule_list--rules--dst_asn_set--namespace) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](resources--forward_proxy_policy--properties--rule_list--rules--dst_asn_set.md#schema-rule_list--rules--dst_asn_set--tenant) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](resources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md#section) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](resources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md#schema-rule_list--rules--dst_ip_prefix_set--name) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](resources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md#schema-rule_list--rules--dst_ip_prefix_set--namespace) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](resources--forward_proxy_policy--properties--rule_list--rules--dst_ip_prefix_set.md#schema-rule_list--rules--dst_ip_prefix_set--tenant) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](resources--forward_proxy_policy--properties--rule_list--rules--dst_label_selector.md#section) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](resources--forward_proxy_policy--properties--rule_list--rules--dst_label_selector.md#schema-rule_list--rules--dst_label_selector--expressions) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](resources--forward_proxy_policy--properties--rule_list--rules--dst_prefix_list.md#section) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](resources--forward_proxy_policy--properties--rule_list--rules--dst_prefix_list.md#schema-rule_list--rules--dst_prefix_list--prefixes) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](resources--forward_proxy_policy--properties--rule_list--rules--http_list.md#section) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#section) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list--any_path.md#section) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--exact_value) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--path_exact_value) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--path_prefix_value) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--path_regex_value) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--regex_value) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](resources--forward_proxy_policy--properties--rule_list--rules--http_list--http_list.md#schema-rule_list--rules--http_list--http_list--suffix_value) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](resources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md#section) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](resources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md#schema-rule_list--rules--ip_prefix_set--name) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](resources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md#schema-rule_list--rules--ip_prefix_set--namespace) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](resources--forward_proxy_policy--properties--rule_list--rules--ip_prefix_set.md#schema-rule_list--rules--ip_prefix_set--tenant) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](resources--forward_proxy_policy--properties--rule_list--rules--label_selector.md#section) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](resources--forward_proxy_policy--properties--rule_list--rules--label_selector.md#schema-rule_list--rules--label_selector--expressions) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--forward_proxy_policy--properties--rule_list--rules--metadata.md#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--forward_proxy_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--forward_proxy_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](resources--forward_proxy_policy--properties--rule_list--rules--no_http_connect_port.md#section) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](resources--forward_proxy_policy--properties--rule_list--rules--port_matcher.md#section) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](resources--forward_proxy_policy--properties--rule_list--rules--port_matcher.md#schema-rule_list--rules--port_matcher--invert_matcher) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](resources--forward_proxy_policy--properties--rule_list--rules--port_matcher.md#schema-rule_list--rules--port_matcher--ports) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](resources--forward_proxy_policy--properties--rule_list--rules--prefix_list.md#section) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](resources--forward_proxy_policy--properties--rule_list--rules--prefix_list.md#schema-rule_list--rules--prefix_list--prefixes) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](resources--forward_proxy_policy--properties--rule_list--rules--tls_list.md#section) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md#section) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md#schema-rule_list--rules--tls_list--tls_list--exact_value) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md#schema-rule_list--rules--tls_list--tls_list--regex_value) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md#schema-rule_list--rules--tls_list--tls_list--suffix_value) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](resources--forward_proxy_policy--properties--rule_list--rules--url_category_list.md#section) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](resources--forward_proxy_policy--properties--rule_list--rules--url_category_list.md#schema-rule_list--rules--url_category_list--url_categories) |
| `timeouts` | [timeouts](resources--forward_proxy_policy--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--forward_proxy_policy--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--forward_proxy_policy--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--forward_proxy_policy--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--forward_proxy_policy--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [allow_all](resources--forward_proxy_policy--properties--allow_all.md)
- [allow_list](resources--forward_proxy_policy--properties--allow_list.md)
- [any_proxy](resources--forward_proxy_policy--properties--any_proxy.md)
- [deny_list](resources--forward_proxy_policy--properties--deny_list.md)
- [drp_http_connect](resources--forward_proxy_policy--properties--drp_http_connect.md)
- [network_connector](resources--forward_proxy_policy--properties--network_connector.md)
- [proxy_label_selector](resources--forward_proxy_policy--properties--proxy_label_selector.md)
- [rule_list](resources--forward_proxy_policy--properties--rule_list.md)
- [timeouts](resources--forward_proxy_policy--properties--timeouts.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
