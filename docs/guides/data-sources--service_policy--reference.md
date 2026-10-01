---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 63340, "body_sha256": "sha256:028450d39a6f891f7631bb4091de5855954d21b80d02170310cc4de265312ff5", "canonical_id": "xcsh-docs:data-sources:service_policy:reference", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:allow_all_requests", "xcsh-docs:data-sources:service_policy:properties:allow_list", "xcsh-docs:data-sources:service_policy:properties:any_server", "xcsh-docs:data-sources:service_policy:properties:deny_all_requests", "xcsh-docs:data-sources:service_policy:properties:deny_list", "xcsh-docs:data-sources:service_policy:properties:rule_list", "xcsh-docs:data-sources:service_policy:properties:server_name_matcher", "xcsh-docs:data-sources:service_policy:properties:server_selector"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:reference", "parent_id": "xcsh-docs:data-sources:service_policy:fundamentals", "path": "docs/guides/data-sources--service_policy--reference.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- Property reference

## Direct properties

- [allow_all_requests](data-sources--service_policy--properties--allow_all_requests.md): complete subsection reference.

- [allow_list](data-sources--service_policy--properties--allow_list.md): complete subsection reference.

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

- [any_server](data-sources--service_policy--properties--any_server.md): complete subsection reference.

- [deny_all_requests](data-sources--service_policy--properties--deny_all_requests.md): complete subsection reference.

- [deny_list](data-sources--service_policy--properties--deny_list.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ServicePolicy.

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

Name of the ServicePolicy.

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

Namespace where the ServicePolicy exists.

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

- [rule_list](data-sources--service_policy--properties--rule_list.md): complete subsection reference.

<a id="schema-server_name"></a>

### server_name property

Type: `"string"`. Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is..

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is directed to a virtual K8s service, the actual names also contain the name of that service. The
predicate evaluates to true if any of the actual names is the same as the expected server name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](data-sources--service_policy--properties--server_name_matcher.md): complete subsection reference.

- [server_selector](data-sources--service_policy--properties--server_selector.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](data-sources--service_policy--properties--allow_all_requests.md#section) |
| `allow_list` | [allow_list](data-sources--service_policy--properties--allow_list.md#section) |
| `allow_list.asn_list` | [allow_list.asn_list](data-sources--service_policy--properties--allow_list--asn_list.md#section) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](data-sources--service_policy--properties--allow_list--asn_list.md#schema-allow_list--asn_list--as_numbers) |
| `allow_list.asn_set` | [allow_list.asn_set](data-sources--service_policy--properties--allow_list--asn_set.md#section) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](data-sources--service_policy--properties--allow_list--asn_set.md#schema-allow_list--asn_set--name) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](data-sources--service_policy--properties--allow_list--asn_set.md#schema-allow_list--asn_set--namespace) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](data-sources--service_policy--properties--allow_list--asn_set.md#schema-allow_list--asn_set--tenant) |
| `allow_list.country_list` | [allow_list.country_list](data-sources--service_policy--properties--allow_list.md#schema-allow_list--country_list) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](data-sources--service_policy--properties--allow_list--default_action_allow.md#section) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](data-sources--service_policy--properties--allow_list--default_action_deny.md#section) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](data-sources--service_policy--properties--allow_list--default_action_next_policy.md#section) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](data-sources--service_policy--properties--allow_list--ip_prefix_set.md#section) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](data-sources--service_policy--properties--allow_list--ip_prefix_set.md#schema-allow_list--ip_prefix_set--name) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](data-sources--service_policy--properties--allow_list--ip_prefix_set.md#schema-allow_list--ip_prefix_set--namespace) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](data-sources--service_policy--properties--allow_list--ip_prefix_set.md#schema-allow_list--ip_prefix_set--tenant) |
| `allow_list.prefix_list` | [allow_list.prefix_list](data-sources--service_policy--properties--allow_list--prefix_list.md#section) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](data-sources--service_policy--properties--allow_list--prefix_list.md#schema-allow_list--prefix_list--prefixes) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](data-sources--service_policy--properties--allow_list.md#schema-allow_list--tls_fingerprint_classes) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](data-sources--service_policy--properties--allow_list.md#schema-allow_list--tls_fingerprint_values) |
| `annotations` | [annotations](data-sources--service_policy--reference.md#schema-annotations) |
| `any_server` | [any_server](data-sources--service_policy--properties--any_server.md#section) |
| `deny_all_requests` | [deny_all_requests](data-sources--service_policy--properties--deny_all_requests.md#section) |
| `deny_list` | [deny_list](data-sources--service_policy--properties--deny_list.md#section) |
| `deny_list.asn_list` | [deny_list.asn_list](data-sources--service_policy--properties--deny_list--asn_list.md#section) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](data-sources--service_policy--properties--deny_list--asn_list.md#schema-deny_list--asn_list--as_numbers) |
| `deny_list.asn_set` | [deny_list.asn_set](data-sources--service_policy--properties--deny_list--asn_set.md#section) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](data-sources--service_policy--properties--deny_list--asn_set.md#schema-deny_list--asn_set--name) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](data-sources--service_policy--properties--deny_list--asn_set.md#schema-deny_list--asn_set--namespace) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](data-sources--service_policy--properties--deny_list--asn_set.md#schema-deny_list--asn_set--tenant) |
| `deny_list.country_list` | [deny_list.country_list](data-sources--service_policy--properties--deny_list.md#schema-deny_list--country_list) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](data-sources--service_policy--properties--deny_list--default_action_allow.md#section) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](data-sources--service_policy--properties--deny_list--default_action_deny.md#section) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](data-sources--service_policy--properties--deny_list--default_action_next_policy.md#section) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](data-sources--service_policy--properties--deny_list--ip_prefix_set.md#section) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](data-sources--service_policy--properties--deny_list--ip_prefix_set.md#schema-deny_list--ip_prefix_set--name) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](data-sources--service_policy--properties--deny_list--ip_prefix_set.md#schema-deny_list--ip_prefix_set--namespace) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](data-sources--service_policy--properties--deny_list--ip_prefix_set.md#schema-deny_list--ip_prefix_set--tenant) |
| `deny_list.prefix_list` | [deny_list.prefix_list](data-sources--service_policy--properties--deny_list--prefix_list.md#section) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](data-sources--service_policy--properties--deny_list--prefix_list.md#schema-deny_list--prefix_list--prefixes) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](data-sources--service_policy--properties--deny_list.md#schema-deny_list--tls_fingerprint_classes) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](data-sources--service_policy--properties--deny_list.md#schema-deny_list--tls_fingerprint_values) |
| `description` | [description](data-sources--service_policy--reference.md#schema-description) |
| `id` | [id](data-sources--service_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--service_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--service_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--service_policy--reference.md#schema-namespace) |
| `rule_list` | [rule_list](data-sources--service_policy--properties--rule_list.md#section) |
| `rule_list.rules` | [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md#section) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--service_policy--properties--rule_list--rules--metadata.md#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--service_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--service_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.spec` | [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md#section) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](data-sources--service_policy--properties--rule_list--rules--spec.md#schema-rule_list--rules--spec--action) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](data-sources--service_policy--properties--rule_list--rules--spec--any_asn.md#section) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](data-sources--service_policy--properties--rule_list--rules--spec--any_client.md#section) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](data-sources--service_policy--properties--rule_list--rules--spec--any_ip.md#section) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](data-sources--service_policy--properties--rule_list--rules--spec--api_group_matcher.md#section) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--api_group_matcher.md#schema-rule_list--rules--spec--api_group_matcher--invert_matcher) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](data-sources--service_policy--properties--rule_list--rules--spec--api_group_matcher.md#schema-rule_list--rules--spec--api_group_matcher--match) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers.md#section) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--check_not_present.md#section) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--check_present.md#section) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers.md#schema-rule_list--rules--spec--arg_matchers--invert_matcher) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--item.md#section) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--item.md#schema-rule_list--rules--spec--arg_matchers--item--exact_values) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--item.md#schema-rule_list--rules--spec--arg_matchers--item--regex_values) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers--item.md#schema-rule_list--rules--spec--arg_matchers--item--transformers) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](data-sources--service_policy--properties--rule_list--rules--spec--arg_matchers.md#schema-rule_list--rules--spec--arg_matchers--name) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](data-sources--service_policy--properties--rule_list--rules--spec--asn_list.md#section) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](data-sources--service_policy--properties--rule_list--rules--spec--asn_list.md#schema-rule_list--rules--spec--asn_list--as_numbers) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher.md#section) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#section) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#schema-rule_list--rules--spec--asn_matcher--asn_sets--kind) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#schema-rule_list--rules--spec--asn_matcher--asn_sets--name) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#schema-rule_list--rules--spec--asn_matcher--asn_sets--namespace) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#schema-rule_list--rules--spec--asn_matcher--asn_sets--tenant) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](data-sources--service_policy--properties--rule_list--rules--spec--asn_matcher--asn_sets.md#schema-rule_list--rules--spec--asn_matcher--asn_sets--uid) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md#section) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md#schema-rule_list--rules--spec--body_matcher--exact_values) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md#schema-rule_list--rules--spec--body_matcher--regex_values) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](data-sources--service_policy--properties--rule_list--rules--spec--body_matcher.md#schema-rule_list--rules--spec--body_matcher--transformers) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](data-sources--service_policy--properties--rule_list--rules--spec--bot_action.md#section) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--bot_skip_processing.md#section) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--none.md#section) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](data-sources--service_policy--properties--rule_list--rules--spec.md#schema-rule_list--rules--spec--client_name) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md#section) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md#schema-rule_list--rules--spec--client_name_matcher--exact_values) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md#schema-rule_list--rules--spec--client_name_matcher--regex_values) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](data-sources--service_policy--properties--rule_list--rules--spec--client_name_matcher.md#schema-rule_list--rules--spec--client_name_matcher--transformers) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](data-sources--service_policy--properties--rule_list--rules--spec--client_selector.md#section) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](data-sources--service_policy--properties--rule_list--rules--spec--client_selector.md#schema-rule_list--rules--spec--client_selector--expressions) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers.md#section) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--check_not_present.md#section) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--check_present.md#section) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers.md#schema-rule_list--rules--spec--cookie_matchers--invert_matcher) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--item.md#section) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--item.md#schema-rule_list--rules--spec--cookie_matchers--item--exact_values) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--item.md#schema-rule_list--rules--spec--cookie_matchers--item--regex_values) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers--item.md#schema-rule_list--rules--spec--cookie_matchers--item--transformers) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](data-sources--service_policy--properties--rule_list--rules--spec--cookie_matchers.md#schema-rule_list--rules--spec--cookie_matchers--name) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md#section) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md#schema-rule_list--rules--spec--domain_matcher--exact_values) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md#schema-rule_list--rules--spec--domain_matcher--regex_values) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](data-sources--service_policy--properties--rule_list--rules--spec--domain_matcher.md#schema-rule_list--rules--spec--domain_matcher--transformers) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](data-sources--service_policy--properties--rule_list--rules--spec.md#schema-rule_list--rules--spec--expiration_timestamp) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](data-sources--service_policy--properties--rule_list--rules--spec--headers.md#section) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](data-sources--service_policy--properties--rule_list--rules--spec--headers--check_not_present.md#section) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](data-sources--service_policy--properties--rule_list--rules--spec--headers--check_present.md#section) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--headers.md#schema-rule_list--rules--spec--headers--invert_matcher) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](data-sources--service_policy--properties--rule_list--rules--spec--headers--item.md#section) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--headers--item.md#schema-rule_list--rules--spec--headers--item--exact_values) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--headers--item.md#schema-rule_list--rules--spec--headers--item--regex_values) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](data-sources--service_policy--properties--rule_list--rules--spec--headers--item.md#schema-rule_list--rules--spec--headers--item--transformers) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](data-sources--service_policy--properties--rule_list--rules--spec--headers.md#schema-rule_list--rules--spec--headers--name) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](data-sources--service_policy--properties--rule_list--rules--spec--http_method.md#section) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--http_method.md#schema-rule_list--rules--spec--http_method--invert_matcher) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](data-sources--service_policy--properties--rule_list--rules--spec--http_method.md#schema-rule_list--rules--spec--http_method--methods) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher.md#section) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher.md#schema-rule_list--rules--spec--ip_matcher--invert_matcher) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#section) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#schema-rule_list--rules--spec--ip_matcher--prefix_sets--kind) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#schema-rule_list--rules--spec--ip_matcher--prefix_sets--name) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#schema-rule_list--rules--spec--ip_matcher--prefix_sets--namespace) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#schema-rule_list--rules--spec--ip_matcher--prefix_sets--tenant) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](data-sources--service_policy--properties--rule_list--rules--spec--ip_matcher--prefix_sets.md#schema-rule_list--rules--spec--ip_matcher--prefix_sets--uid) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_prefix_list.md#section) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](data-sources--service_policy--properties--rule_list--rules--spec--ip_prefix_list.md#schema-rule_list--rules--spec--ip_prefix_list--invert_match) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](data-sources--service_policy--properties--rule_list--rules--spec--ip_prefix_list.md#schema-rule_list--rules--spec--ip_prefix_list--ip_prefixes) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](data-sources--service_policy--properties--rule_list--rules--spec--ip_threat_category_list.md#section) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](data-sources--service_policy--properties--rule_list--rules--spec--ip_threat_category_list.md#schema-rule_list--rules--spec--ip_threat_category_list--ip_threat_categories) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](data-sources--service_policy--properties--rule_list--rules--spec--ja4_tls_fingerprint.md#section) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--ja4_tls_fingerprint.md#schema-rule_list--rules--spec--ja4_tls_fingerprint--exact_values) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims.md#section) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--check_not_present.md#section) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--check_present.md#section) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims.md#schema-rule_list--rules--spec--jwt_claims--invert_matcher) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--item.md#section) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--item.md#schema-rule_list--rules--spec--jwt_claims--item--exact_values) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--item.md#schema-rule_list--rules--spec--jwt_claims--item--regex_values) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims--item.md#schema-rule_list--rules--spec--jwt_claims--item--transformers) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](data-sources--service_policy--properties--rule_list--rules--spec--jwt_claims.md#schema-rule_list--rules--spec--jwt_claims--name) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](data-sources--service_policy--properties--rule_list--rules--spec--label_matcher.md#section) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](data-sources--service_policy--properties--rule_list--rules--spec--label_matcher.md#schema-rule_list--rules--spec--label_matcher--keys) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](data-sources--service_policy--properties--rule_list--rules--spec.md#schema-rule_list--rules--spec--log_rule_evaluation) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](data-sources--service_policy--properties--rule_list--rules--spec--mum_action.md#section) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](data-sources--service_policy--properties--rule_list--rules--spec--mum_action--default.md#section) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--mum_action--skip_processing.md#section) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](data-sources--service_policy--properties--rule_list--rules--spec--path.md#section) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--encoded_path_matcher) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--exact_values) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--invert_matcher) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--prefix_values) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--regex_values) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--suffix_values) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](data-sources--service_policy--properties--rule_list--rules--spec--path.md#schema-rule_list--rules--spec--path--transformers) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](data-sources--service_policy--properties--rule_list--rules--spec--port_matcher.md#section) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--port_matcher.md#schema-rule_list--rules--spec--port_matcher--invert_matcher) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](data-sources--service_policy--properties--rule_list--rules--spec--port_matcher.md#schema-rule_list--rules--spec--port_matcher--ports) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](data-sources--service_policy--properties--rule_list--rules--spec--query_params.md#section) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](data-sources--service_policy--properties--rule_list--rules--spec--query_params--check_not_present.md#section) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](data-sources--service_policy--properties--rule_list--rules--spec--query_params--check_present.md#section) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](data-sources--service_policy--properties--rule_list--rules--spec--query_params.md#schema-rule_list--rules--spec--query_params--invert_matcher) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](data-sources--service_policy--properties--rule_list--rules--spec--query_params--item.md#section) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--query_params--item.md#schema-rule_list--rules--spec--query_params--item--exact_values) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--query_params--item.md#schema-rule_list--rules--spec--query_params--item--regex_values) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](data-sources--service_policy--properties--rule_list--rules--spec--query_params--item.md#schema-rule_list--rules--spec--query_params--item--transformers) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](data-sources--service_policy--properties--rule_list--rules--spec--query_params.md#schema-rule_list--rules--spec--query_params--key) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_cookie_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_count_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_cookie_key_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_key_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_cookie_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_value_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_header_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_count_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_header_key_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_key_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_header_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_value_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_parameter_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_count_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_parameter_name_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_name_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_parameter_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_value_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_query_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_query_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_request_line_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_line_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_request_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_size_none.md#section) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md#schema-rule_list--rules--spec--request_constraints--max_url_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_url_size_none.md#section) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy.md#section) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_any.md#section) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments.md#section) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments--segments.md#section) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments--segments.md#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--name) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments--segments.md#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--namespace) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments--segments.md#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--tenant) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--intra_segment.md#section) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_any.md#section) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments.md#section) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md#section) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md#schema-rule_list--rules--spec--segment_policy--src_segments--segments--name) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md#schema-rule_list--rules--spec--segment_policy--src_segments--segments--namespace) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md#schema-rule_list--rules--spec--segment_policy--src_segments--segments--tenant) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md#section) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md#schema-rule_list--rules--spec--tls_fingerprint_matcher--classes) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md#schema-rule_list--rules--spec--tls_fingerprint_matcher--exact_values) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](data-sources--service_policy--properties--rule_list--rules--spec--tls_fingerprint_matcher.md#schema-rule_list--rules--spec--tls_fingerprint_matcher--excluded_values) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](data-sources--service_policy--properties--rule_list--rules--spec--user_identity_matcher.md#section) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](data-sources--service_policy--properties--rule_list--rules--spec--user_identity_matcher.md#schema-rule_list--rules--spec--user_identity_matcher--exact_values) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](data-sources--service_policy--properties--rule_list--rules--spec--user_identity_matcher.md#schema-rule_list--rules--spec--user_identity_matcher--regex_values) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](data-sources--service_policy--properties--rule_list--rules--spec--waf_action.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--exclude_attack_type) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--exclude_violation) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--none.md#section) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--waf_skip_processing.md#section) |
| `server_name` | [server_name](data-sources--service_policy--reference.md#schema-server_name) |
| `server_name_matcher` | [server_name_matcher](data-sources--service_policy--properties--server_name_matcher.md#section) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](data-sources--service_policy--properties--server_name_matcher.md#schema-server_name_matcher--exact_values) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](data-sources--service_policy--properties--server_name_matcher.md#schema-server_name_matcher--regex_values) |
| `server_selector` | [server_selector](data-sources--service_policy--properties--server_selector.md#section) |
| `server_selector.expressions` | [server_selector.expressions](data-sources--service_policy--properties--server_selector.md#schema-server_selector--expressions) |

## Next pages

- [allow_all_requests](data-sources--service_policy--properties--allow_all_requests.md)
- [allow_list](data-sources--service_policy--properties--allow_list.md)
- [any_server](data-sources--service_policy--properties--any_server.md)
- [deny_all_requests](data-sources--service_policy--properties--deny_all_requests.md)
- [deny_list](data-sources--service_policy--properties--deny_list.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [server_name_matcher](data-sources--service_policy--properties--server_name_matcher.md)
- [server_selector](data-sources--service_policy--properties--server_selector.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
