---
page_title: "Property reference"
subcategory: "DNS"
description: "Property reference for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 21844, "body_sha256": "sha256:8b4271338f7aab0f2a030388e6a6c49e669c112dba266961652d585b2886fc91", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:fallback_pool", "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "xcsh-docs:resources:dns_load_balancer:properties:rule_list", "xcsh-docs:resources:dns_load_balancer:properties:timeouts"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:reference", "parent_id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "path": "documentation/resources/dns_load_balancer/properties/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
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

- [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/): complete subsection reference.

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

Name of the DNS Load Balancer. Must be unique within the namespace.

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

Namespace for the DNS Load Balancer. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

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

<a id="schema-record_type"></a>

### record_type property

Type: `"string"`. Optional, Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "A",
  "enum": [
    "A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/): complete subsection reference.

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-disable) |
| `fallback_pool` | [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/#section) |
| `fallback_pool.name` | [fallback_pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--name) |
| `fallback_pool.namespace` | [fallback_pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--namespace) |
| `fallback_pool.tenant` | [fallback_pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-namespace) |
| `record_type` | [record_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/#schema-record_type) |
| `response_cache` | [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/#section) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/#section) |
| `response_cache.disable_spec` | [response_cache.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/disable_spec/#section) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/#section) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_cidr_ipv4) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_cidr_ipv6) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_ttl) |
| `rule_list` | [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/#section) |
| `rule_list.rules` | [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/#section) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_list/#section) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_list/#schema-rule_list--rules--asn_list--as_numbers) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/#section) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#section) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--kind) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--name) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--namespace) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--tenant) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--uid) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/#section) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/#schema-rule_list--rules--geo_location_label_selector--expressions) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#section) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--name) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--namespace) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--tenant) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#section) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#schema-rule_list--rules--ip_prefix_list--invert_match) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#schema-rule_list--rules--ip_prefix_list--ip_prefixes) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/#section) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/#schema-rule_list--rules--ip_prefix_set--invert_matcher) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#section) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--kind) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--name) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--namespace) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--tenant) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--uid) |
| `rule_list.rules.pool` | [rule_list.rules.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/#section) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--name) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--namespace) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--tenant) |
| `rule_list.rules.score` | [rule_list.rules.score](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/#schema-rule_list--rules--score) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/fallback_pool/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/timeouts/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
