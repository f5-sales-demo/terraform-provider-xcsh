---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 20881, "body_sha256": "sha256:96b216bc8fff42637288ca0cb33565c9a939e258cada030b7749968181f4fdf4", "canonical_id": "xcsh-docs:resources:network_policy_view:reference", "child_ids": ["xcsh-docs:resources:network_policy_view:properties:egress_rules", "xcsh-docs:resources:network_policy_view:properties:endpoint", "xcsh-docs:resources:network_policy_view:properties:ingress_rules", "xcsh-docs:resources:network_policy_view:properties:timeouts"], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:reference", "parent_id": "xcsh-docs:resources:network_policy_view:fundamentals", "path": "docs/guides/resources--network_policy_view--reference.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
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

- [egress_rules](resources--network_policy_view--properties--egress_rules.md): complete subsection reference.

- [endpoint](resources--network_policy_view--properties--endpoint.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](resources--network_policy_view--properties--ingress_rules.md): complete subsection reference.

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

Name of the Network Policy View. Must be unique within the namespace.

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

Type: `"string"`. Optional, Computed.

Namespace for the Network Policy View. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [timeouts](resources--network_policy_view--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_policy_view--reference.md#schema-annotations) |
| `description` | [description](resources--network_policy_view--reference.md#schema-description) |
| `disable` | [disable](resources--network_policy_view--reference.md#schema-disable) |
| `egress_rules` | [egress_rules](resources--network_policy_view--properties--egress_rules.md#section) |
| `egress_rules.action` | [egress_rules.action](resources--network_policy_view--properties--egress_rules.md#schema-egress_rules--action) |
| `egress_rules.adv_action` | [egress_rules.adv_action](resources--network_policy_view--properties--egress_rules--adv_action.md#section) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](resources--network_policy_view--properties--egress_rules--adv_action.md#schema-egress_rules--adv_action--action) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](resources--network_policy_view--properties--egress_rules--all_tcp_traffic.md#section) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](resources--network_policy_view--properties--egress_rules--all_traffic.md#section) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](resources--network_policy_view--properties--egress_rules--all_udp_traffic.md#section) |
| `egress_rules.any` | [egress_rules.any](resources--network_policy_view--properties--egress_rules--any.md#section) |
| `egress_rules.applications` | [egress_rules.applications](resources--network_policy_view--properties--egress_rules--applications.md#section) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](resources--network_policy_view--properties--egress_rules--applications.md#schema-egress_rules--applications--applications) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](resources--network_policy_view--properties--egress_rules--inside_endpoints.md#section) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](resources--network_policy_view--properties--egress_rules--ip_prefix_set.md#section) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#section) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#schema-egress_rules--ip_prefix_set--ref--kind) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#schema-egress_rules--ip_prefix_set--ref--name) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#schema-egress_rules--ip_prefix_set--ref--namespace) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#schema-egress_rules--ip_prefix_set--ref--tenant) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md#schema-egress_rules--ip_prefix_set--ref--uid) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](resources--network_policy_view--properties--egress_rules--label_matcher.md#section) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](resources--network_policy_view--properties--egress_rules--label_matcher.md#schema-egress_rules--label_matcher--keys) |
| `egress_rules.label_selector` | [egress_rules.label_selector](resources--network_policy_view--properties--egress_rules--label_selector.md#section) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](resources--network_policy_view--properties--egress_rules--label_selector.md#schema-egress_rules--label_selector--expressions) |
| `egress_rules.metadata` | [egress_rules.metadata](resources--network_policy_view--properties--egress_rules--metadata.md#section) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](resources--network_policy_view--properties--egress_rules--metadata.md#schema-egress_rules--metadata--description_spec) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](resources--network_policy_view--properties--egress_rules--metadata.md#schema-egress_rules--metadata--name) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](resources--network_policy_view--properties--egress_rules--outside_endpoints.md#section) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](resources--network_policy_view--properties--egress_rules--prefix_list.md#section) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](resources--network_policy_view--properties--egress_rules--prefix_list.md#schema-egress_rules--prefix_list--prefixes) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](resources--network_policy_view--properties--egress_rules--protocol_port_range.md#section) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](resources--network_policy_view--properties--egress_rules--protocol_port_range.md#schema-egress_rules--protocol_port_range--port_ranges) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](resources--network_policy_view--properties--egress_rules--protocol_port_range.md#schema-egress_rules--protocol_port_range--protocol) |
| `endpoint` | [endpoint](resources--network_policy_view--properties--endpoint.md#section) |
| `endpoint.any` | [endpoint.any](resources--network_policy_view--properties--endpoint--any.md#section) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](resources--network_policy_view--properties--endpoint--inside_endpoints.md#section) |
| `endpoint.label_selector` | [endpoint.label_selector](resources--network_policy_view--properties--endpoint--label_selector.md#section) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](resources--network_policy_view--properties--endpoint--label_selector.md#schema-endpoint--label_selector--expressions) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](resources--network_policy_view--properties--endpoint--outside_endpoints.md#section) |
| `endpoint.prefix_list` | [endpoint.prefix_list](resources--network_policy_view--properties--endpoint--prefix_list.md#section) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](resources--network_policy_view--properties--endpoint--prefix_list.md#schema-endpoint--prefix_list--prefixes) |
| `id` | [id](resources--network_policy_view--reference.md#schema-id) |
| `ingress_rules` | [ingress_rules](resources--network_policy_view--properties--ingress_rules.md#section) |
| `ingress_rules.action` | [ingress_rules.action](resources--network_policy_view--properties--ingress_rules.md#schema-ingress_rules--action) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](resources--network_policy_view--properties--ingress_rules--adv_action.md#section) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](resources--network_policy_view--properties--ingress_rules--adv_action.md#schema-ingress_rules--adv_action--action) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](resources--network_policy_view--properties--ingress_rules--all_tcp_traffic.md#section) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](resources--network_policy_view--properties--ingress_rules--all_traffic.md#section) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](resources--network_policy_view--properties--ingress_rules--all_udp_traffic.md#section) |
| `ingress_rules.any` | [ingress_rules.any](resources--network_policy_view--properties--ingress_rules--any.md#section) |
| `ingress_rules.applications` | [ingress_rules.applications](resources--network_policy_view--properties--ingress_rules--applications.md#section) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](resources--network_policy_view--properties--ingress_rules--applications.md#schema-ingress_rules--applications--applications) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](resources--network_policy_view--properties--ingress_rules--inside_endpoints.md#section) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](resources--network_policy_view--properties--ingress_rules--ip_prefix_set.md#section) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#section) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#schema-ingress_rules--ip_prefix_set--ref--kind) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#schema-ingress_rules--ip_prefix_set--ref--name) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#schema-ingress_rules--ip_prefix_set--ref--namespace) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#schema-ingress_rules--ip_prefix_set--ref--tenant) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md#schema-ingress_rules--ip_prefix_set--ref--uid) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](resources--network_policy_view--properties--ingress_rules--label_matcher.md#section) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](resources--network_policy_view--properties--ingress_rules--label_matcher.md#schema-ingress_rules--label_matcher--keys) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](resources--network_policy_view--properties--ingress_rules--label_selector.md#section) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](resources--network_policy_view--properties--ingress_rules--label_selector.md#schema-ingress_rules--label_selector--expressions) |
| `ingress_rules.metadata` | [ingress_rules.metadata](resources--network_policy_view--properties--ingress_rules--metadata.md#section) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](resources--network_policy_view--properties--ingress_rules--metadata.md#schema-ingress_rules--metadata--description_spec) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](resources--network_policy_view--properties--ingress_rules--metadata.md#schema-ingress_rules--metadata--name) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](resources--network_policy_view--properties--ingress_rules--outside_endpoints.md#section) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](resources--network_policy_view--properties--ingress_rules--prefix_list.md#section) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](resources--network_policy_view--properties--ingress_rules--prefix_list.md#schema-ingress_rules--prefix_list--prefixes) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](resources--network_policy_view--properties--ingress_rules--protocol_port_range.md#section) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](resources--network_policy_view--properties--ingress_rules--protocol_port_range.md#schema-ingress_rules--protocol_port_range--port_ranges) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](resources--network_policy_view--properties--ingress_rules--protocol_port_range.md#schema-ingress_rules--protocol_port_range--protocol) |
| `labels` | [labels](resources--network_policy_view--reference.md#schema-labels) |
| `name` | [name](resources--network_policy_view--reference.md#schema-name) |
| `namespace` | [namespace](resources--network_policy_view--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--network_policy_view--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--network_policy_view--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_view--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--network_policy_view--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--network_policy_view--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- [endpoint](resources--network_policy_view--properties--endpoint.md)
- [ingress_rules](resources--network_policy_view--properties--ingress_rules.md)
- [timeouts](resources--network_policy_view--properties--timeouts.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
