---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 54000, "body_sha256": "sha256:9986bb3377c3468e07a570771c11f5cf57d7ed62b029a909f45e10273ff6c73b", "canonical_id": "xcsh-docs:resources:route:reference", "child_ids": ["xcsh-docs:resources:route:properties:routes", "xcsh-docs:resources:route:properties:timeouts"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:reference", "parent_id": "xcsh-docs:resources:route:fundamentals", "path": "docs/guides/resources--route--reference.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_route](../resources/route.md)
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

Name of the Route. Must be unique within the namespace.

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

Namespace where the Route is created.

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

- [routes](resources--route--properties--routes.md): complete subsection reference.

- [timeouts](resources--route--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--route--reference.md#schema-annotations) |
| `description` | [description](resources--route--reference.md#schema-description) |
| `disable` | [disable](resources--route--reference.md#schema-disable) |
| `id` | [id](resources--route--reference.md#schema-id) |
| `labels` | [labels](resources--route--reference.md#schema-labels) |
| `name` | [name](resources--route--reference.md#schema-name) |
| `namespace` | [namespace](resources--route--reference.md#schema-namespace) |
| `routes` | [routes](resources--route--properties--routes.md#section) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](resources--route--properties--routes--bot_defense_javascript_injection.md#section) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](resources--route--properties--routes--bot_defense_javascript_injection.md#schema-routes--bot_defense_javascript_injection--javascript_location) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](resources--route--properties--routes--bot_defense_javascript_injection--javascript_tags.md#section) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](resources--route--properties--routes--bot_defense_javascript_injection--javascript_tags.md#schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--route--properties--routes--bot_defense_javascript_injection--javascript_tags--tag_attributes.md#section) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](resources--route--properties--routes--bot_defense_javascript_injection--javascript_tags--tag_attributes.md#schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--javascript_tag) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](resources--route--properties--routes--bot_defense_javascript_injection--javascript_tags--tag_attributes.md#schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--tag_value) |
| `routes.disable_location_add` | [routes.disable_location_add](resources--route--properties--routes.md#schema-routes--disable_location_add) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](resources--route--properties--routes--inherited_bot_defense_javascript_injection.md#section) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](resources--route--properties--routes--inherited_waf_exclusion.md#section) |
| `routes.match` | [routes.match](resources--route--properties--routes--match.md#section) |
| `routes.match.headers` | [routes.match.headers](resources--route--properties--routes--match--headers.md#section) |
| `routes.match.headers.exact` | [routes.match.headers.exact](resources--route--properties--routes--match--headers.md#schema-routes--match--headers--exact) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](resources--route--properties--routes--match--headers.md#schema-routes--match--headers--invert_match) |
| `routes.match.headers.name` | [routes.match.headers.name](resources--route--properties--routes--match--headers.md#schema-routes--match--headers--name) |
| `routes.match.headers.presence` | [routes.match.headers.presence](resources--route--properties--routes--match--headers.md#schema-routes--match--headers--presence) |
| `routes.match.headers.regex` | [routes.match.headers.regex](resources--route--properties--routes--match--headers.md#schema-routes--match--headers--regex) |
| `routes.match.http_method` | [routes.match.http_method](resources--route--properties--routes--match.md#schema-routes--match--http_method) |
| `routes.match.incoming_port` | [routes.match.incoming_port](resources--route--properties--routes--match--incoming_port.md#section) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](resources--route--properties--routes--match--incoming_port--no_port_match.md#section) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](resources--route--properties--routes--match--incoming_port.md#schema-routes--match--incoming_port--port) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](resources--route--properties--routes--match--incoming_port.md#schema-routes--match--incoming_port--port_ranges) |
| `routes.match.path` | [routes.match.path](resources--route--properties--routes--match--path.md#section) |
| `routes.match.path.path` | [routes.match.path.path](resources--route--properties--routes--match--path.md#schema-routes--match--path--path) |
| `routes.match.path.prefix` | [routes.match.path.prefix](resources--route--properties--routes--match--path.md#schema-routes--match--path--prefix) |
| `routes.match.path.regex` | [routes.match.path.regex](resources--route--properties--routes--match--path.md#schema-routes--match--path--regex) |
| `routes.match.query_params` | [routes.match.query_params](resources--route--properties--routes--match--query_params.md#section) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](resources--route--properties--routes--match--query_params.md#schema-routes--match--query_params--exact) |
| `routes.match.query_params.key` | [routes.match.query_params.key](resources--route--properties--routes--match--query_params.md#schema-routes--match--query_params--key) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](resources--route--properties--routes--match--query_params.md#schema-routes--match--query_params--regex) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](resources--route--properties--routes--request_cookies_to_add.md#section) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](resources--route--properties--routes--request_cookies_to_add.md#schema-routes--request_cookies_to_add--name) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](resources--route--properties--routes--request_cookies_to_add.md#schema-routes--request_cookies_to_add--overwrite) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](resources--route--properties--routes--request_cookies_to_add--secret_value.md#section) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--properties--routes--request_cookies_to_add--secret_value--clear_secret_info.md#schema-routes--request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](resources--route--properties--routes--request_cookies_to_add--secret_value--clear_secret_info.md#schema-routes--request_cookies_to_add--secret_value--clear_secret_info--url) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](resources--route--properties--routes--request_cookies_to_add.md#schema-routes--request_cookies_to_add--value) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](resources--route--properties--routes.md#schema-routes--request_cookies_to_remove) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](resources--route--properties--routes--request_headers_to_add.md#section) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](resources--route--properties--routes--request_headers_to_add.md#schema-routes--request_headers_to_add--append) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](resources--route--properties--routes--request_headers_to_add.md#schema-routes--request_headers_to_add--name) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](resources--route--properties--routes--request_headers_to_add--secret_value.md#section) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](resources--route--properties--routes--request_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--properties--routes--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--properties--routes--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--properties--routes--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](resources--route--properties--routes--request_headers_to_add--secret_value--clear_secret_info.md#section) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--properties--routes--request_headers_to_add--secret_value--clear_secret_info.md#schema-routes--request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](resources--route--properties--routes--request_headers_to_add--secret_value--clear_secret_info.md#schema-routes--request_headers_to_add--secret_value--clear_secret_info--url) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](resources--route--properties--routes--request_headers_to_add.md#schema-routes--request_headers_to_add--value) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](resources--route--properties--routes.md#schema-routes--request_headers_to_remove) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](resources--route--properties--routes--response_cookies_to_add.md#section) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--add_domain) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--add_expiry) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](resources--route--properties--routes--response_cookies_to_add--add_httponly.md#section) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](resources--route--properties--routes--response_cookies_to_add--add_partitioned.md#section) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--add_path) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](resources--route--properties--routes--response_cookies_to_add--add_secure.md#section) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](resources--route--properties--routes--response_cookies_to_add--ignore_domain.md#section) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](resources--route--properties--routes--response_cookies_to_add--ignore_expiry.md#section) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](resources--route--properties--routes--response_cookies_to_add--ignore_httponly.md#section) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](resources--route--properties--routes--response_cookies_to_add--ignore_max_age.md#section) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](resources--route--properties--routes--response_cookies_to_add--ignore_partitioned.md#section) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](resources--route--properties--routes--response_cookies_to_add--ignore_path.md#section) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](resources--route--properties--routes--response_cookies_to_add--ignore_samesite.md#section) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](resources--route--properties--routes--response_cookies_to_add--ignore_secure.md#section) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](resources--route--properties--routes--response_cookies_to_add--ignore_value.md#section) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--max_age_value) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--name) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--overwrite) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](resources--route--properties--routes--response_cookies_to_add--samesite_lax.md#section) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](resources--route--properties--routes--response_cookies_to_add--samesite_none.md#section) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](resources--route--properties--routes--response_cookies_to_add--samesite_strict.md#section) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](resources--route--properties--routes--response_cookies_to_add--secret_value.md#section) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](resources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](resources--route--properties--routes--response_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--properties--routes--response_cookies_to_add--secret_value--clear_secret_info.md#schema-routes--response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](resources--route--properties--routes--response_cookies_to_add--secret_value--clear_secret_info.md#schema-routes--response_cookies_to_add--secret_value--clear_secret_info--url) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](resources--route--properties--routes--response_cookies_to_add.md#schema-routes--response_cookies_to_add--value) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](resources--route--properties--routes.md#schema-routes--response_cookies_to_remove) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](resources--route--properties--routes--response_headers_to_add.md#section) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](resources--route--properties--routes--response_headers_to_add.md#schema-routes--response_headers_to_add--append) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](resources--route--properties--routes--response_headers_to_add.md#schema-routes--response_headers_to_add--name) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](resources--route--properties--routes--response_headers_to_add--secret_value.md#section) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](resources--route--properties--routes--response_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--properties--routes--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--properties--routes--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--properties--routes--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](resources--route--properties--routes--response_headers_to_add--secret_value--clear_secret_info.md#section) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--properties--routes--response_headers_to_add--secret_value--clear_secret_info.md#schema-routes--response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](resources--route--properties--routes--response_headers_to_add--secret_value--clear_secret_info.md#schema-routes--response_headers_to_add--secret_value--clear_secret_info--url) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](resources--route--properties--routes--response_headers_to_add.md#schema-routes--response_headers_to_add--value) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](resources--route--properties--routes.md#schema-routes--response_headers_to_remove) |
| `routes.route_destination` | [routes.route_destination](resources--route--properties--routes--route_destination.md#section) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](resources--route--properties--routes--route_destination.md#schema-routes--route_destination--auto_host_rewrite) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](resources--route--properties--routes--route_destination--buffer_policy.md#section) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](resources--route--properties--routes--route_destination--buffer_policy.md#schema-routes--route_destination--buffer_policy--disabled) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](resources--route--properties--routes--route_destination--buffer_policy.md#schema-routes--route_destination--buffer_policy--max_request_bytes) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](resources--route--properties--routes--route_destination--cors_policy.md#section) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--allow_credentials) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--allow_headers) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--allow_methods) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--allow_origin) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--allow_origin_regex) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--disabled) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--expose_headers) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](resources--route--properties--routes--route_destination--cors_policy.md#schema-routes--route_destination--cors_policy--maximum_age) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](resources--route--properties--routes--route_destination--csrf_policy.md#section) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](resources--route--properties--routes--route_destination--csrf_policy--all_load_balancer_domains.md#section) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](resources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md#section) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](resources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md#schema-routes--route_destination--csrf_policy--custom_domain_list--domains) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](resources--route--properties--routes--route_destination--csrf_policy--disabled.md#section) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](resources--route--properties--routes--route_destination--destinations.md#section) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](resources--route--properties--routes--route_destination--destinations--cluster.md#section) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](resources--route--properties--routes--route_destination--destinations--cluster.md#schema-routes--route_destination--destinations--cluster--kind) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](resources--route--properties--routes--route_destination--destinations--cluster.md#schema-routes--route_destination--destinations--cluster--name) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](resources--route--properties--routes--route_destination--destinations--cluster.md#schema-routes--route_destination--destinations--cluster--namespace) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](resources--route--properties--routes--route_destination--destinations--cluster.md#schema-routes--route_destination--destinations--cluster--tenant) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](resources--route--properties--routes--route_destination--destinations--cluster.md#schema-routes--route_destination--destinations--cluster--uid) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](resources--route--properties--routes--route_destination--destinations--endpoint_subsets.md#section) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](resources--route--properties--routes--route_destination--destinations.md#schema-routes--route_destination--destinations--priority) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](resources--route--properties--routes--route_destination--destinations.md#schema-routes--route_destination--destinations--weight) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](resources--route--properties--routes--route_destination--do_not_retract_cluster.md#section) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](resources--route--properties--routes--route_destination--endpoint_subsets.md#section) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](resources--route--properties--routes--route_destination--hash_policy.md#section) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](resources--route--properties--routes--route_destination--hash_policy--cookie.md#section) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](resources--route--properties--routes--route_destination--hash_policy--cookie--add_httponly.md#section) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](resources--route--properties--routes--route_destination--hash_policy--cookie--add_secure.md#section) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](resources--route--properties--routes--route_destination--hash_policy--cookie--ignore_httponly.md#section) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](resources--route--properties--routes--route_destination--hash_policy--cookie--ignore_samesite.md#section) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](resources--route--properties--routes--route_destination--hash_policy--cookie--ignore_secure.md#section) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](resources--route--properties--routes--route_destination--hash_policy--cookie.md#schema-routes--route_destination--hash_policy--cookie--name) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](resources--route--properties--routes--route_destination--hash_policy--cookie.md#schema-routes--route_destination--hash_policy--cookie--path) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](resources--route--properties--routes--route_destination--hash_policy--cookie--samesite_lax.md#section) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](resources--route--properties--routes--route_destination--hash_policy--cookie--samesite_none.md#section) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](resources--route--properties--routes--route_destination--hash_policy--cookie--samesite_strict.md#section) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](resources--route--properties--routes--route_destination--hash_policy--cookie.md#schema-routes--route_destination--hash_policy--cookie--ttl) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](resources--route--properties--routes--route_destination--hash_policy.md#schema-routes--route_destination--hash_policy--header_name) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](resources--route--properties--routes--route_destination--hash_policy.md#schema-routes--route_destination--hash_policy--source_ip) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](resources--route--properties--routes--route_destination--hash_policy.md#schema-routes--route_destination--hash_policy--terminal) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](resources--route--properties--routes--route_destination.md#schema-routes--route_destination--host_rewrite) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](resources--route--properties--routes--route_destination--mirror_policy.md#section) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#section) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#schema-routes--route_destination--mirror_policy--cluster--kind) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#schema-routes--route_destination--mirror_policy--cluster--name) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#schema-routes--route_destination--mirror_policy--cluster--namespace) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#schema-routes--route_destination--mirror_policy--cluster--tenant) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](resources--route--properties--routes--route_destination--mirror_policy--cluster.md#schema-routes--route_destination--mirror_policy--cluster--uid) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](resources--route--properties--routes--route_destination--mirror_policy--percent.md#section) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](resources--route--properties--routes--route_destination--mirror_policy--percent.md#schema-routes--route_destination--mirror_policy--percent--denominator) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](resources--route--properties--routes--route_destination--mirror_policy--percent.md#schema-routes--route_destination--mirror_policy--percent--numerator) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](resources--route--properties--routes--route_destination.md#schema-routes--route_destination--prefix_rewrite) |
| `routes.route_destination.priority` | [routes.route_destination.priority](resources--route--properties--routes--route_destination.md#schema-routes--route_destination--priority) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](resources--route--properties--routes--route_destination--query_params.md#section) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](resources--route--properties--routes--route_destination--query_params--remove_all_params.md#section) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](resources--route--properties--routes--route_destination--query_params.md#schema-routes--route_destination--query_params--replace_params) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](resources--route--properties--routes--route_destination--query_params--retain_all_params.md#section) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](resources--route--properties--routes--route_destination--regex_rewrite.md#section) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](resources--route--properties--routes--route_destination--regex_rewrite.md#schema-routes--route_destination--regex_rewrite--pattern) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](resources--route--properties--routes--route_destination--regex_rewrite.md#schema-routes--route_destination--regex_rewrite--substitution) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](resources--route--properties--routes--route_destination--retract_cluster.md#section) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](resources--route--properties--routes--route_destination--retry_policy.md#section) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](resources--route--properties--routes--route_destination--retry_policy--back_off.md#section) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](resources--route--properties--routes--route_destination--retry_policy--back_off.md#schema-routes--route_destination--retry_policy--back_off--base_interval) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](resources--route--properties--routes--route_destination--retry_policy--back_off.md#schema-routes--route_destination--retry_policy--back_off--max_interval) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](resources--route--properties--routes--route_destination--retry_policy.md#schema-routes--route_destination--retry_policy--num_retries) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](resources--route--properties--routes--route_destination--retry_policy.md#schema-routes--route_destination--retry_policy--per_try_timeout) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](resources--route--properties--routes--route_destination--retry_policy.md#schema-routes--route_destination--retry_policy--retriable_status_codes) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](resources--route--properties--routes--route_destination--retry_policy.md#schema-routes--route_destination--retry_policy--retry_condition) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](resources--route--properties--routes--route_destination--spdy_config.md#section) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](resources--route--properties--routes--route_destination--spdy_config.md#schema-routes--route_destination--spdy_config--use_spdy) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](resources--route--properties--routes--route_destination.md#schema-routes--route_destination--timeout) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](resources--route--properties--routes--route_destination--web_socket_config.md#section) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](resources--route--properties--routes--route_destination--web_socket_config.md#schema-routes--route_destination--web_socket_config--use_websocket) |
| `routes.route_direct_response` | [routes.route_direct_response](resources--route--properties--routes--route_direct_response.md#section) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](resources--route--properties--routes--route_direct_response.md#schema-routes--route_direct_response--response_body_encoded) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](resources--route--properties--routes--route_direct_response.md#schema-routes--route_direct_response--response_code) |
| `routes.route_redirect` | [routes.route_redirect](resources--route--properties--routes--route_redirect.md#section) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--host_redirect) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--path_redirect) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--prefix_rewrite) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--proto_redirect) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](resources--route--properties--routes--route_redirect--remove_all_params.md#section) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--replace_params) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](resources--route--properties--routes--route_redirect.md#schema-routes--route_redirect--response_code) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](resources--route--properties--routes--route_redirect--retain_all_params.md#section) |
| `routes.service_policy` | [routes.service_policy](resources--route--properties--routes--service_policy.md#section) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](resources--route--properties--routes--service_policy.md#schema-routes--service_policy--disable_spec) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](resources--route--properties--routes--waf_exclusion_policy.md#section) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](resources--route--properties--routes--waf_exclusion_policy.md#schema-routes--waf_exclusion_policy--name) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](resources--route--properties--routes--waf_exclusion_policy.md#schema-routes--waf_exclusion_policy--namespace) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](resources--route--properties--routes--waf_exclusion_policy.md#schema-routes--waf_exclusion_policy--tenant) |
| `routes.waf_type` | [routes.waf_type](resources--route--properties--routes--waf_type.md#section) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](resources--route--properties--routes--waf_type--app_firewall.md#section) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#section) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#schema-routes--waf_type--app_firewall--app_firewall--kind) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#schema-routes--waf_type--app_firewall--app_firewall--name) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#schema-routes--waf_type--app_firewall--app_firewall--namespace) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#schema-routes--waf_type--app_firewall--app_firewall--tenant) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](resources--route--properties--routes--waf_type--app_firewall--app_firewall.md#schema-routes--waf_type--app_firewall--app_firewall--uid) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](resources--route--properties--routes--waf_type--disable_waf.md#section) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](resources--route--properties--routes--waf_type--inherit_waf.md#section) |
| `timeouts` | [timeouts](resources--route--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--route--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--route--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--route--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--route--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [routes](resources--route--properties--routes.md)
- [timeouts](resources--route--properties--timeouts.md)
- [xcsh_route](../resources/route.md)
