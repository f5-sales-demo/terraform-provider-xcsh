---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 24974, "body_sha256": "sha256:031f1ae99308321d008bd91ede400faeae9be75d7c4282068d1b98591521d425", "canonical_id": "xcsh-docs:data-sources:fast_acl:reference", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:protocol_policer", "xcsh-docs:data-sources:fast_acl:properties:re_acl", "xcsh-docs:data-sources:fast_acl:properties:site_acl"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:reference", "parent_id": "xcsh-docs:data-sources:fast_acl:fundamentals", "path": "docs/guides/data-sources--fast_acl--reference.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
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

Description of the FastACL.

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

Name of the FastACL.

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

Type: `"string"`. Optional, Computed.

Namespace where the FastACL exists.

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

- [protocol_policer](data-sources--fast_acl--properties--protocol_policer.md): complete subsection reference.

- [re_acl](data-sources--fast_acl--properties--re_acl.md): complete subsection reference.

- [site_acl](data-sources--fast_acl--properties--site_acl.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--fast_acl--reference.md#schema-annotations) |
| `description` | [description](data-sources--fast_acl--reference.md#schema-description) |
| `id` | [id](data-sources--fast_acl--reference.md#schema-id) |
| `labels` | [labels](data-sources--fast_acl--reference.md#schema-labels) |
| `name` | [name](data-sources--fast_acl--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--fast_acl--reference.md#schema-namespace) |
| `protocol_policer` | [protocol_policer](data-sources--fast_acl--properties--protocol_policer.md#section) |
| `protocol_policer.name` | [protocol_policer.name](data-sources--fast_acl--properties--protocol_policer.md#schema-protocol_policer--name) |
| `protocol_policer.namespace` | [protocol_policer.namespace](data-sources--fast_acl--properties--protocol_policer.md#schema-protocol_policer--namespace) |
| `protocol_policer.tenant` | [protocol_policer.tenant](data-sources--fast_acl--properties--protocol_policer.md#schema-protocol_policer--tenant) |
| `re_acl` | [re_acl](data-sources--fast_acl--properties--re_acl.md#section) |
| `re_acl.all_public_vips` | [re_acl.all_public_vips](data-sources--fast_acl--properties--re_acl--all_public_vips.md#section) |
| `re_acl.default_tenant_vip` | [re_acl.default_tenant_vip](data-sources--fast_acl--properties--re_acl--default_tenant_vip.md#section) |
| `re_acl.fast_acl_rules` | [re_acl.fast_acl_rules](data-sources--fast_acl--properties--re_acl--fast_acl_rules.md#section) |
| `re_acl.fast_acl_rules.action` | [re_acl.fast_acl_rules.action](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action.md#section) |
| `re_acl.fast_acl_rules.action.policer_action` | [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action.md#section) |
| `re_acl.fast_acl_rules.action.policer_action.ref` | [re_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#section) |
| `re_acl.fast_acl_rules.action.policer_action.ref.kind` | [re_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#schema-re_acl--fast_acl_rules--action--policer_action--ref--kind) |
| `re_acl.fast_acl_rules.action.policer_action.ref.name` | [re_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#schema-re_acl--fast_acl_rules--action--policer_action--ref--name) |
| `re_acl.fast_acl_rules.action.policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#schema-re_acl--fast_acl_rules--action--policer_action--ref--namespace) |
| `re_acl.fast_acl_rules.action.policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#schema-re_acl--fast_acl_rules--action--policer_action--ref--tenant) |
| `re_acl.fast_acl_rules.action.policer_action.ref.uid` | [re_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--policer_action--ref.md#schema-re_acl--fast_acl_rules--action--policer_action--ref--uid) |
| `re_acl.fast_acl_rules.action.protocol_policer_action` | [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action.md#section) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#section) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-re_acl--fast_acl_rules--action--protocol_policer_action--ref--kind) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-re_acl--fast_acl_rules--action--protocol_policer_action--ref--name) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-re_acl--fast_acl_rules--action--protocol_policer_action--ref--namespace) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-re_acl--fast_acl_rules--action--protocol_policer_action--ref--tenant) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-re_acl--fast_acl_rules--action--protocol_policer_action--ref--uid) |
| `re_acl.fast_acl_rules.action.simple_action` | [re_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--properties--re_acl--fast_acl_rules--action.md#schema-re_acl--fast_acl_rules--action--simple_action) |
| `re_acl.fast_acl_rules.ip_prefix_set` | [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set.md#section) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref` | [re_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#section) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [re_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-re_acl--fast_acl_rules--ip_prefix_set--ref--kind) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.name` | [re_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-re_acl--fast_acl_rules--ip_prefix_set--ref--name) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [re_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-re_acl--fast_acl_rules--ip_prefix_set--ref--namespace) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [re_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-re_acl--fast_acl_rules--ip_prefix_set--ref--tenant) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [re_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-re_acl--fast_acl_rules--ip_prefix_set--ref--uid) |
| `re_acl.fast_acl_rules.metadata` | [re_acl.fast_acl_rules.metadata](data-sources--fast_acl--properties--re_acl--fast_acl_rules--metadata.md#section) |
| `re_acl.fast_acl_rules.metadata.description_spec` | [re_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--properties--re_acl--fast_acl_rules--metadata.md#schema-re_acl--fast_acl_rules--metadata--description_spec) |
| `re_acl.fast_acl_rules.metadata.name` | [re_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--properties--re_acl--fast_acl_rules--metadata.md#schema-re_acl--fast_acl_rules--metadata--name) |
| `re_acl.fast_acl_rules.port` | [re_acl.fast_acl_rules.port](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port.md#section) |
| `re_acl.fast_acl_rules.port.all` | [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--all.md#section) |
| `re_acl.fast_acl_rules.port.dns` | [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--dns.md#section) |
| `re_acl.fast_acl_rules.port.user_defined` | [re_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port.md#schema-re_acl--fast_acl_rules--port--user_defined) |
| `re_acl.fast_acl_rules.prefix` | [re_acl.fast_acl_rules.prefix](data-sources--fast_acl--properties--re_acl--fast_acl_rules--prefix.md#section) |
| `re_acl.fast_acl_rules.prefix.prefix` | [re_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--properties--re_acl--fast_acl_rules--prefix.md#schema-re_acl--fast_acl_rules--prefix--prefix) |
| `re_acl.selected_tenant_vip` | [re_acl.selected_tenant_vip](data-sources--fast_acl--properties--re_acl--selected_tenant_vip.md#section) |
| `re_acl.selected_tenant_vip.default_tenant_vip` | [re_acl.selected_tenant_vip.default_tenant_vip](data-sources--fast_acl--properties--re_acl--selected_tenant_vip.md#schema-re_acl--selected_tenant_vip--default_tenant_vip) |
| `re_acl.selected_tenant_vip.public_ip_refs` | [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md#section) |
| `re_acl.selected_tenant_vip.public_ip_refs.name` | [re_acl.selected_tenant_vip.public_ip_refs.name](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md#schema-re_acl--selected_tenant_vip--public_ip_refs--name) |
| `re_acl.selected_tenant_vip.public_ip_refs.namespace` | [re_acl.selected_tenant_vip.public_ip_refs.namespace](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md#schema-re_acl--selected_tenant_vip--public_ip_refs--namespace) |
| `re_acl.selected_tenant_vip.public_ip_refs.tenant` | [re_acl.selected_tenant_vip.public_ip_refs.tenant](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md#schema-re_acl--selected_tenant_vip--public_ip_refs--tenant) |
| `site_acl` | [site_acl](data-sources--fast_acl--properties--site_acl.md#section) |
| `site_acl.all_services` | [site_acl.all_services](data-sources--fast_acl--properties--site_acl--all_services.md#section) |
| `site_acl.fast_acl_rules` | [site_acl.fast_acl_rules](data-sources--fast_acl--properties--site_acl--fast_acl_rules.md#section) |
| `site_acl.fast_acl_rules.action` | [site_acl.fast_acl_rules.action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action.md#section) |
| `site_acl.fast_acl_rules.action.policer_action` | [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action.md#section) |
| `site_acl.fast_acl_rules.action.policer_action.ref` | [site_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#section) |
| `site_acl.fast_acl_rules.action.policer_action.ref.kind` | [site_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#schema-site_acl--fast_acl_rules--action--policer_action--ref--kind) |
| `site_acl.fast_acl_rules.action.policer_action.ref.name` | [site_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#schema-site_acl--fast_acl_rules--action--policer_action--ref--name) |
| `site_acl.fast_acl_rules.action.policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#schema-site_acl--fast_acl_rules--action--policer_action--ref--namespace) |
| `site_acl.fast_acl_rules.action.policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#schema-site_acl--fast_acl_rules--action--policer_action--ref--tenant) |
| `site_acl.fast_acl_rules.action.policer_action.ref.uid` | [site_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action--ref.md#schema-site_acl--fast_acl_rules--action--policer_action--ref--uid) |
| `site_acl.fast_acl_rules.action.protocol_policer_action` | [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action.md#section) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#section) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-site_acl--fast_acl_rules--action--protocol_policer_action--ref--kind) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-site_acl--fast_acl_rules--action--protocol_policer_action--ref--name) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-site_acl--fast_acl_rules--action--protocol_policer_action--ref--namespace) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-site_acl--fast_acl_rules--action--protocol_policer_action--ref--tenant) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action--ref.md#schema-site_acl--fast_acl_rules--action--protocol_policer_action--ref--uid) |
| `site_acl.fast_acl_rules.action.simple_action` | [site_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action.md#schema-site_acl--fast_acl_rules--action--simple_action) |
| `site_acl.fast_acl_rules.ip_prefix_set` | [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set.md#section) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref` | [site_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#section) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [site_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-site_acl--fast_acl_rules--ip_prefix_set--ref--kind) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.name` | [site_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-site_acl--fast_acl_rules--ip_prefix_set--ref--name) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [site_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-site_acl--fast_acl_rules--ip_prefix_set--ref--namespace) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [site_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-site_acl--fast_acl_rules--ip_prefix_set--ref--tenant) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [site_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md#schema-site_acl--fast_acl_rules--ip_prefix_set--ref--uid) |
| `site_acl.fast_acl_rules.metadata` | [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--properties--site_acl--fast_acl_rules--metadata.md#section) |
| `site_acl.fast_acl_rules.metadata.description_spec` | [site_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--properties--site_acl--fast_acl_rules--metadata.md#schema-site_acl--fast_acl_rules--metadata--description_spec) |
| `site_acl.fast_acl_rules.metadata.name` | [site_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--properties--site_acl--fast_acl_rules--metadata.md#schema-site_acl--fast_acl_rules--metadata--name) |
| `site_acl.fast_acl_rules.port` | [site_acl.fast_acl_rules.port](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port.md#section) |
| `site_acl.fast_acl_rules.port.all` | [site_acl.fast_acl_rules.port.all](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port--all.md#section) |
| `site_acl.fast_acl_rules.port.dns` | [site_acl.fast_acl_rules.port.dns](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port--dns.md#section) |
| `site_acl.fast_acl_rules.port.user_defined` | [site_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port.md#schema-site_acl--fast_acl_rules--port--user_defined) |
| `site_acl.fast_acl_rules.prefix` | [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--properties--site_acl--fast_acl_rules--prefix.md#section) |
| `site_acl.fast_acl_rules.prefix.prefix` | [site_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--properties--site_acl--fast_acl_rules--prefix.md#schema-site_acl--fast_acl_rules--prefix--prefix) |
| `site_acl.inside_network` | [site_acl.inside_network](data-sources--fast_acl--properties--site_acl--inside_network.md#section) |
| `site_acl.interface_services` | [site_acl.interface_services](data-sources--fast_acl--properties--site_acl--interface_services.md#section) |
| `site_acl.outside_network` | [site_acl.outside_network](data-sources--fast_acl--properties--site_acl--outside_network.md#section) |
| `site_acl.vip_services` | [site_acl.vip_services](data-sources--fast_acl--properties--site_acl--vip_services.md#section) |

## Next pages

- [protocol_policer](data-sources--fast_acl--properties--protocol_policer.md)
- [re_acl](data-sources--fast_acl--properties--re_acl.md)
- [site_acl](data-sources--fast_acl--properties--site_acl.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
