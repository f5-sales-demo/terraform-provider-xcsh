---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": ["enhanced firewall policy"], "body_bytes": 28019, "body_sha256": "sha256:f957c993fd24dc9f885a47f5f52668e83e5ff4681aeb7983509168d922119bb8", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:allow_all", "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:deny_all", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "xcsh-docs:resources:enhanced_firewall_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:fundamentals", "path": "documentation/resources/enhanced_firewall_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:allow_all", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowed destinations"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_destinations", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowed_destinations"], "syntax": "block", "type": "object"}, {"aliases": ["allowed sources"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_sources", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowed_sources"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["denied destinations"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_destinations", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["denied_destinations"], "syntax": "block", "type": "object"}, {"aliases": ["denied sources"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:denied_sources", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["denied_sources"], "syntax": "block", "type": "object"}, {"aliases": ["deny all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:deny_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list"], "anchor": "section", "description": "Custom Enhanced Firewall Policy Rules.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "type": "requires"}], "schema_path": ["rule_list"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_enhanced_firewall_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- Property reference

## Direct properties

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allow_all/): complete subsection reference.

- [allowed_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_destinations/): complete subsection reference.

- [allowed_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_sources/): complete subsection reference.

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

- [denied_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_destinations/): complete subsection reference.

- [denied_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_sources/): complete subsection reference.

- [deny_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/deny_all/): complete subsection reference.

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

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allow_all/#section) |
| `allowed_destinations` | [allowed_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_destinations/#section) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_destinations/#schema-allowed_destinations--prefix) |
| `allowed_sources` | [allowed_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_sources/#section) |
| `allowed_sources.prefix` | [allowed_sources.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_sources/#schema-allowed_sources--prefix) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-annotations) |
| `denied_destinations` | [denied_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_destinations/#section) |
| `denied_destinations.prefix` | [denied_destinations.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_destinations/#schema-denied_destinations--prefix) |
| `denied_sources` | [denied_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_sources/#section) |
| `denied_sources.prefix` | [denied_sources.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_sources/#schema-denied_sources--prefix) |
| `deny_all` | [deny_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/deny_all/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/#schema-namespace) |
| `rule_list` | [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/#section) |
| `rule_list.rules` | [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/#section) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/#section) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/#schema-rule_list--rules--advanced_action--action) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_destinations/#section) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sli_vips/#section) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/#section) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sources/#section) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_tcp_traffic/#section) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_traffic/#section) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_udp_traffic/#section) |
| `rule_list.rules.allow` | [rule_list.rules.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/allow/#section) |
| `rule_list.rules.applications` | [rule_list.rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/applications/#section) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/applications/#schema-rule_list--rules--applications--applications) |
| `rule_list.rules.deny` | [rule_list.rules.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/deny/#section) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/#section) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/#schema-rule_list--rules--destination_aws_vpc_ids--vpc_id) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/#section) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#section) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#schema-rule_list--rules--destination_ip_prefix_set--ref--kind) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#schema-rule_list--rules--destination_ip_prefix_set--ref--name) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#schema-rule_list--rules--destination_ip_prefix_set--ref--namespace) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#schema-rule_list--rules--destination_ip_prefix_set--ref--tenant) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/ref/#schema-rule_list--rules--destination_ip_prefix_set--ref--uid) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/#section) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/#schema-rule_list--rules--destination_label_selector--expressions) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/#section) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/#schema-rule_list--rules--destination_prefix_list--prefixes) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/#section) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/nfv_service/#section) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/nfv_service/#schema-rule_list--rules--insert_service--nfv_service--name) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/nfv_service/#schema-rule_list--rules--insert_service--nfv_service--namespace) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/nfv_service/#schema-rule_list--rules--insert_service--nfv_service--tenant) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_destinations/#section) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/#section) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/#section) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/#schema-rule_list--rules--label_matcher--keys) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/metadata/#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_destinations/#section) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_sources/#section) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/#section) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/#schema-rule_list--rules--protocol_port_range--port_ranges) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/#schema-rule_list--rules--protocol_port_range--protocol) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/#section) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/#schema-rule_list--rules--source_aws_vpc_ids--vpc_id) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/#section) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#section) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#schema-rule_list--rules--source_ip_prefix_set--ref--kind) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#schema-rule_list--rules--source_ip_prefix_set--ref--name) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#schema-rule_list--rules--source_ip_prefix_set--ref--namespace) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#schema-rule_list--rules--source_ip_prefix_set--ref--tenant) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/#schema-rule_list--rules--source_ip_prefix_set--ref--uid) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/#section) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/#schema-rule_list--rules--source_label_selector--expressions) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/#section) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/#schema-rule_list--rules--source_prefix_list--prefixes) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allow_all/)
- [allowed_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_destinations/)
- [allowed_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/allowed_sources/)
- [denied_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_destinations/)
- [denied_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/denied_sources/)
- [deny_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/deny_all/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/timeouts/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
