---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 23136, "body_sha256": "sha256:4dee101f9952d9b3209f5333af148ba8408e055402ee5ce42e6c4d9f75baffa1", "canonical_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:allow_all", "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:deny_all", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "xcsh-docs:resources:enhanced_firewall_policy:properties:timeouts"], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:fundamentals", "path": "docs/guides/resources--enhanced_firewall_policy--reference.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
- Property reference

## Direct properties

- [allow_all](resources--enhanced_firewall_policy--properties--allow_all.md): complete subsection reference.

- [allowed_destinations](resources--enhanced_firewall_policy--properties--allowed_destinations.md): complete subsection reference.

- [allowed_sources](resources--enhanced_firewall_policy--properties--allowed_sources.md): complete subsection reference.

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

- [denied_destinations](resources--enhanced_firewall_policy--properties--denied_destinations.md): complete subsection reference.

- [denied_sources](resources--enhanced_firewall_policy--properties--denied_sources.md): complete subsection reference.

- [deny_all](resources--enhanced_firewall_policy--properties--deny_all.md): complete subsection reference.

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

Name of the Enhanced Firewall Policy. Must be unique within the namespace.

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

Namespace where the Enhanced Firewall Policy is created.

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

- [rule_list](resources--enhanced_firewall_policy--properties--rule_list.md): complete subsection reference.

- [timeouts](resources--enhanced_firewall_policy--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--enhanced_firewall_policy--properties--allow_all.md#section) |
| `allowed_destinations` | [allowed_destinations](resources--enhanced_firewall_policy--properties--allowed_destinations.md#section) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](resources--enhanced_firewall_policy--properties--allowed_destinations.md#schema-allowed_destinations--prefix) |
| `allowed_sources` | [allowed_sources](resources--enhanced_firewall_policy--properties--allowed_sources.md#section) |
| `allowed_sources.prefix` | [allowed_sources.prefix](resources--enhanced_firewall_policy--properties--allowed_sources.md#schema-allowed_sources--prefix) |
| `annotations` | [annotations](resources--enhanced_firewall_policy--reference.md#schema-annotations) |
| `denied_destinations` | [denied_destinations](resources--enhanced_firewall_policy--properties--denied_destinations.md#section) |
| `denied_destinations.prefix` | [denied_destinations.prefix](resources--enhanced_firewall_policy--properties--denied_destinations.md#schema-denied_destinations--prefix) |
| `denied_sources` | [denied_sources](resources--enhanced_firewall_policy--properties--denied_sources.md#section) |
| `denied_sources.prefix` | [denied_sources.prefix](resources--enhanced_firewall_policy--properties--denied_sources.md#schema-denied_sources--prefix) |
| `deny_all` | [deny_all](resources--enhanced_firewall_policy--properties--deny_all.md#section) |
| `description` | [description](resources--enhanced_firewall_policy--reference.md#schema-description) |
| `disable` | [disable](resources--enhanced_firewall_policy--reference.md#schema-disable) |
| `id` | [id](resources--enhanced_firewall_policy--reference.md#schema-id) |
| `labels` | [labels](resources--enhanced_firewall_policy--reference.md#schema-labels) |
| `name` | [name](resources--enhanced_firewall_policy--reference.md#schema-name) |
| `namespace` | [namespace](resources--enhanced_firewall_policy--reference.md#schema-namespace) |
| `rule_list` | [rule_list](resources--enhanced_firewall_policy--properties--rule_list.md#section) |
| `rule_list.rules` | [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md#section) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](resources--enhanced_firewall_policy--properties--rule_list--rules--advanced_action.md#section) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](resources--enhanced_firewall_policy--properties--rule_list--rules--advanced_action.md#schema-rule_list--rules--advanced_action--action) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--enhanced_firewall_policy--properties--rule_list--rules--all_destinations.md#section) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](resources--enhanced_firewall_policy--properties--rule_list--rules--all_sli_vips.md#section) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](resources--enhanced_firewall_policy--properties--rule_list--rules--all_slo_vips.md#section) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--enhanced_firewall_policy--properties--rule_list--rules--all_sources.md#section) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](resources--enhanced_firewall_policy--properties--rule_list--rules--all_tcp_traffic.md#section) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](resources--enhanced_firewall_policy--properties--rule_list--rules--all_traffic.md#section) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](resources--enhanced_firewall_policy--properties--rule_list--rules--all_udp_traffic.md#section) |
| `rule_list.rules.allow` | [rule_list.rules.allow](resources--enhanced_firewall_policy--properties--rule_list--rules--allow.md#section) |
| `rule_list.rules.applications` | [rule_list.rules.applications](resources--enhanced_firewall_policy--properties--rule_list--rules--applications.md#section) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](resources--enhanced_firewall_policy--properties--rule_list--rules--applications.md#schema-rule_list--rules--applications--applications) |
| `rule_list.rules.deny` | [rule_list.rules.deny](resources--enhanced_firewall_policy--properties--rule_list--rules--deny.md#section) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_aws_vpc_ids.md#section) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_aws_vpc_ids.md#schema-rule_list--rules--destination_aws_vpc_ids--vpc_id) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set.md#section) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#section) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#schema-rule_list--rules--destination_ip_prefix_set--ref--kind) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#schema-rule_list--rules--destination_ip_prefix_set--ref--name) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#schema-rule_list--rules--destination_ip_prefix_set--ref--namespace) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#schema-rule_list--rules--destination_ip_prefix_set--ref--tenant) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set--ref.md#schema-rule_list--rules--destination_ip_prefix_set--ref--uid) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_label_selector.md#section) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_label_selector.md#schema-rule_list--rules--destination_label_selector--expressions) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_prefix_list.md#section) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](resources--enhanced_firewall_policy--properties--rule_list--rules--destination_prefix_list.md#schema-rule_list--rules--destination_prefix_list--prefixes) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](resources--enhanced_firewall_policy--properties--rule_list--rules--insert_service.md#section) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](resources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md#section) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](resources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md#schema-rule_list--rules--insert_service--nfv_service--name) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](resources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md#schema-rule_list--rules--insert_service--nfv_service--namespace) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](resources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md#schema-rule_list--rules--insert_service--nfv_service--tenant) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](resources--enhanced_firewall_policy--properties--rule_list--rules--inside_destinations.md#section) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](resources--enhanced_firewall_policy--properties--rule_list--rules--inside_sources.md#section) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](resources--enhanced_firewall_policy--properties--rule_list--rules--label_matcher.md#section) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](resources--enhanced_firewall_policy--properties--rule_list--rules--label_matcher.md#schema-rule_list--rules--label_matcher--keys) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--enhanced_firewall_policy--properties--rule_list--rules--metadata.md#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--enhanced_firewall_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--enhanced_firewall_policy--properties--rule_list--rules--metadata.md#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](resources--enhanced_firewall_policy--properties--rule_list--rules--outside_destinations.md#section) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](resources--enhanced_firewall_policy--properties--rule_list--rules--outside_sources.md#section) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](resources--enhanced_firewall_policy--properties--rule_list--rules--protocol_port_range.md#section) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](resources--enhanced_firewall_policy--properties--rule_list--rules--protocol_port_range.md#schema-rule_list--rules--protocol_port_range--port_ranges) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](resources--enhanced_firewall_policy--properties--rule_list--rules--protocol_port_range.md#schema-rule_list--rules--protocol_port_range--protocol) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](resources--enhanced_firewall_policy--properties--rule_list--rules--source_aws_vpc_ids.md#section) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--properties--rule_list--rules--source_aws_vpc_ids.md#schema-rule_list--rules--source_aws_vpc_ids--vpc_id) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set.md#section) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#section) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#schema-rule_list--rules--source_ip_prefix_set--ref--kind) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#schema-rule_list--rules--source_ip_prefix_set--ref--name) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#schema-rule_list--rules--source_ip_prefix_set--ref--namespace) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#schema-rule_list--rules--source_ip_prefix_set--ref--tenant) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md#schema-rule_list--rules--source_ip_prefix_set--ref--uid) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](resources--enhanced_firewall_policy--properties--rule_list--rules--source_label_selector.md#section) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](resources--enhanced_firewall_policy--properties--rule_list--rules--source_label_selector.md#schema-rule_list--rules--source_label_selector--expressions) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](resources--enhanced_firewall_policy--properties--rule_list--rules--source_prefix_list.md#section) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](resources--enhanced_firewall_policy--properties--rule_list--rules--source_prefix_list.md#schema-rule_list--rules--source_prefix_list--prefixes) |
| `timeouts` | [timeouts](resources--enhanced_firewall_policy--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--enhanced_firewall_policy--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--enhanced_firewall_policy--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--enhanced_firewall_policy--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--enhanced_firewall_policy--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [allow_all](resources--enhanced_firewall_policy--properties--allow_all.md)
- [allowed_destinations](resources--enhanced_firewall_policy--properties--allowed_destinations.md)
- [allowed_sources](resources--enhanced_firewall_policy--properties--allowed_sources.md)
- [denied_destinations](resources--enhanced_firewall_policy--properties--denied_destinations.md)
- [denied_sources](resources--enhanced_firewall_policy--properties--denied_sources.md)
- [deny_all](resources--enhanced_firewall_policy--properties--deny_all.md)
- [rule_list](resources--enhanced_firewall_policy--properties--rule_list.md)
- [timeouts](resources--enhanced_firewall_policy--properties--timeouts.md)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
