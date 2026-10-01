---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 30330, "body_sha256": "sha256:f75be538e7fe6861f090aaf1f2af31841c22287dab947133e06634f9bad2cbf0", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:any_server", "xcsh-docs:resources:rate_limiter_policy:properties:rules", "xcsh-docs:resources:rate_limiter_policy:properties:server_name_matcher", "xcsh-docs:resources:rate_limiter_policy:properties:server_selector", "xcsh-docs:resources:rate_limiter_policy:properties:timeouts"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:reference", "parent_id": "xcsh-docs:resources:rate_limiter_policy:fundamentals", "path": "documentation/resources/rate_limiter_policy/properties/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
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

- [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/any_server/): complete subsection reference.

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

Name of the Rate Limiter Policy. Must be unique within the namespace.

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

Namespace where the Rate Limiter Policy is created.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/): complete subsection reference.

<a id="schema-server_name"></a>

### server_name property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_name_matcher/): complete subsection reference.

- [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_selector/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-annotations) |
| `any_server` | [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/any_server/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-namespace) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/#section) |
| `rules.metadata` | [rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/metadata/#section) |
| `rules.metadata.description_spec` | [rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/metadata/#schema-rules--metadata--description_spec) |
| `rules.metadata.name` | [rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/metadata/#schema-rules--metadata--name) |
| `rules.spec` | [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/#section) |
| `rules.spec.any_asn` | [rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/any_asn/#section) |
| `rules.spec.any_country` | [rules.spec.any_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/any_country/#section) |
| `rules.spec.any_ip` | [rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/any_ip/#section) |
| `rules.spec.apply_rate_limiter` | [rules.spec.apply_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/apply_rate_limiter/#section) |
| `rules.spec.asn_list` | [rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_list/#section) |
| `rules.spec.asn_list.as_numbers` | [rules.spec.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_list/#schema-rules--spec--asn_list--as_numbers) |
| `rules.spec.asn_matcher` | [rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/#section) |
| `rules.spec.asn_matcher.asn_sets` | [rules.spec.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#section) |
| `rules.spec.asn_matcher.asn_sets.kind` | [rules.spec.asn_matcher.asn_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#schema-rules--spec--asn_matcher--asn_sets--kind) |
| `rules.spec.asn_matcher.asn_sets.name` | [rules.spec.asn_matcher.asn_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#schema-rules--spec--asn_matcher--asn_sets--name) |
| `rules.spec.asn_matcher.asn_sets.namespace` | [rules.spec.asn_matcher.asn_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#schema-rules--spec--asn_matcher--asn_sets--namespace) |
| `rules.spec.asn_matcher.asn_sets.tenant` | [rules.spec.asn_matcher.asn_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#schema-rules--spec--asn_matcher--asn_sets--tenant) |
| `rules.spec.asn_matcher.asn_sets.uid` | [rules.spec.asn_matcher.asn_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/asn_sets/#schema-rules--spec--asn_matcher--asn_sets--uid) |
| `rules.spec.bypass_rate_limiter` | [rules.spec.bypass_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/#section) |
| `rules.spec.country_list` | [rules.spec.country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/country_list/#section) |
| `rules.spec.country_list.country_codes` | [rules.spec.country_list.country_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/country_list/#schema-rules--spec--country_list--country_codes) |
| `rules.spec.country_list.invert_match` | [rules.spec.country_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/country_list/#schema-rules--spec--country_list--invert_match) |
| `rules.spec.custom_rate_limiter` | [rules.spec.custom_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/#section) |
| `rules.spec.custom_rate_limiter.name` | [rules.spec.custom_rate_limiter.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/#schema-rules--spec--custom_rate_limiter--name) |
| `rules.spec.custom_rate_limiter.namespace` | [rules.spec.custom_rate_limiter.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/#schema-rules--spec--custom_rate_limiter--namespace) |
| `rules.spec.custom_rate_limiter.tenant` | [rules.spec.custom_rate_limiter.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/#schema-rules--spec--custom_rate_limiter--tenant) |
| `rules.spec.domain_matcher` | [rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/domain_matcher/#section) |
| `rules.spec.domain_matcher.exact_values` | [rules.spec.domain_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/domain_matcher/#schema-rules--spec--domain_matcher--exact_values) |
| `rules.spec.domain_matcher.regex_values` | [rules.spec.domain_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/domain_matcher/#schema-rules--spec--domain_matcher--regex_values) |
| `rules.spec.headers` | [rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/#section) |
| `rules.spec.headers.check_not_present` | [rules.spec.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/check_not_present/#section) |
| `rules.spec.headers.check_present` | [rules.spec.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/check_present/#section) |
| `rules.spec.headers.invert_matcher` | [rules.spec.headers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/#schema-rules--spec--headers--invert_matcher) |
| `rules.spec.headers.item` | [rules.spec.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/item/#section) |
| `rules.spec.headers.item.exact_values` | [rules.spec.headers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/item/#schema-rules--spec--headers--item--exact_values) |
| `rules.spec.headers.item.regex_values` | [rules.spec.headers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/item/#schema-rules--spec--headers--item--regex_values) |
| `rules.spec.headers.item.transformers` | [rules.spec.headers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/item/#schema-rules--spec--headers--item--transformers) |
| `rules.spec.headers.name` | [rules.spec.headers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/headers/#schema-rules--spec--headers--name) |
| `rules.spec.http_method` | [rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/http_method/#section) |
| `rules.spec.http_method.invert_matcher` | [rules.spec.http_method.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/http_method/#schema-rules--spec--http_method--invert_matcher) |
| `rules.spec.http_method.methods` | [rules.spec.http_method.methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/http_method/#schema-rules--spec--http_method--methods) |
| `rules.spec.ip_matcher` | [rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/#section) |
| `rules.spec.ip_matcher.invert_matcher` | [rules.spec.ip_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/#schema-rules--spec--ip_matcher--invert_matcher) |
| `rules.spec.ip_matcher.prefix_sets` | [rules.spec.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#section) |
| `rules.spec.ip_matcher.prefix_sets.kind` | [rules.spec.ip_matcher.prefix_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#schema-rules--spec--ip_matcher--prefix_sets--kind) |
| `rules.spec.ip_matcher.prefix_sets.name` | [rules.spec.ip_matcher.prefix_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#schema-rules--spec--ip_matcher--prefix_sets--name) |
| `rules.spec.ip_matcher.prefix_sets.namespace` | [rules.spec.ip_matcher.prefix_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#schema-rules--spec--ip_matcher--prefix_sets--namespace) |
| `rules.spec.ip_matcher.prefix_sets.tenant` | [rules.spec.ip_matcher.prefix_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#schema-rules--spec--ip_matcher--prefix_sets--tenant) |
| `rules.spec.ip_matcher.prefix_sets.uid` | [rules.spec.ip_matcher.prefix_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_matcher/prefix_sets/#schema-rules--spec--ip_matcher--prefix_sets--uid) |
| `rules.spec.ip_prefix_list` | [rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/#section) |
| `rules.spec.ip_prefix_list.invert_match` | [rules.spec.ip_prefix_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/#schema-rules--spec--ip_prefix_list--invert_match) |
| `rules.spec.ip_prefix_list.ip_prefixes` | [rules.spec.ip_prefix_list.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/#schema-rules--spec--ip_prefix_list--ip_prefixes) |
| `rules.spec.path` | [rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#section) |
| `rules.spec.path.encoded_path_matcher` | [rules.spec.path.encoded_path_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--encoded_path_matcher) |
| `rules.spec.path.exact_values` | [rules.spec.path.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--exact_values) |
| `rules.spec.path.invert_matcher` | [rules.spec.path.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--invert_matcher) |
| `rules.spec.path.prefix_values` | [rules.spec.path.prefix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--prefix_values) |
| `rules.spec.path.regex_values` | [rules.spec.path.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--regex_values) |
| `rules.spec.path.suffix_values` | [rules.spec.path.suffix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--suffix_values) |
| `rules.spec.path.transformers` | [rules.spec.path.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/path/#schema-rules--spec--path--transformers) |
| `rules.spec.segment_policy` | [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/#section) |
| `rules.spec.segment_policy.dst_any` | [rules.spec.segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/#section) |
| `rules.spec.segment_policy.dst_segments` | [rules.spec.segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/#section) |
| `rules.spec.segment_policy.dst_segments.segments` | [rules.spec.segment_policy.dst_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/segments/#section) |
| `rules.spec.segment_policy.dst_segments.segments.name` | [rules.spec.segment_policy.dst_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/segments/#schema-rules--spec--segment_policy--dst_segments--segments--name) |
| `rules.spec.segment_policy.dst_segments.segments.namespace` | [rules.spec.segment_policy.dst_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/segments/#schema-rules--spec--segment_policy--dst_segments--segments--namespace) |
| `rules.spec.segment_policy.dst_segments.segments.tenant` | [rules.spec.segment_policy.dst_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/segments/#schema-rules--spec--segment_policy--dst_segments--segments--tenant) |
| `rules.spec.segment_policy.intra_segment` | [rules.spec.segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/#section) |
| `rules.spec.segment_policy.src_any` | [rules.spec.segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/#section) |
| `rules.spec.segment_policy.src_segments` | [rules.spec.segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/#section) |
| `rules.spec.segment_policy.src_segments.segments` | [rules.spec.segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/segments/#section) |
| `rules.spec.segment_policy.src_segments.segments.name` | [rules.spec.segment_policy.src_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/segments/#schema-rules--spec--segment_policy--src_segments--segments--name) |
| `rules.spec.segment_policy.src_segments.segments.namespace` | [rules.spec.segment_policy.src_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/segments/#schema-rules--spec--segment_policy--src_segments--segments--namespace) |
| `rules.spec.segment_policy.src_segments.segments.tenant` | [rules.spec.segment_policy.src_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/segments/#schema-rules--spec--segment_policy--src_segments--segments--tenant) |
| `server_name` | [server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/#schema-server_name) |
| `server_name_matcher` | [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_name_matcher/#section) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_name_matcher/#schema-server_name_matcher--exact_values) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_name_matcher/#schema-server_name_matcher--regex_values) |
| `server_selector` | [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_selector/#section) |
| `server_selector.expressions` | [server_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_selector/#schema-server_selector--expressions) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/any_server/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/)
- [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_name_matcher/)
- [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/server_selector/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/timeouts/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
