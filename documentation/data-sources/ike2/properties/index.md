---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike2."
xcsh_docs: {"aliases": [], "body_bytes": 9661, "body_sha256": "sha256:e6ec09298690995cb289fb7b9957090f98187e6651a7ce11c39a8c7abc1c44fa", "child_ids": ["xcsh-docs:data-sources:ike2:properties:dh_group_set", "xcsh-docs:data-sources:ike2:properties:disable_pfs", "xcsh-docs:data-sources:ike2:properties:ike_keylifetime_hours", "xcsh-docs:data-sources:ike2:properties:ike_keylifetime_minutes", "xcsh-docs:data-sources:ike2:properties:use_default_keylifetime"], "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike2:reference", "parent_id": "xcsh-docs:data-sources:ike2:fundamentals", "path": "documentation/data-sources/ike2/properties/index.md", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_ike2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/)
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

Description of the Ike2.

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

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/dh_group_set/): complete subsection reference.

- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/disable_pfs/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_hours/): complete subsection reference.

- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_minutes/): complete subsection reference.

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

Name of the Ike2.

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

Namespace where the Ike2 exists.

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

- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/use_default_keylifetime/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-description) |
| `dh_group_set` | [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/dh_group_set/#section) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/dh_group_set/#schema-dh_group_set--dh_groups) |
| `disable_pfs` | [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/disable_pfs/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_hours/#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_hours/#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_minutes/#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_minutes/#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/#schema-namespace) |
| `use_default_keylifetime` | [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/use_default_keylifetime/#section) |

## Next pages

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/dh_group_set/)
- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/disable_pfs/)
- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_hours/)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/ike_keylifetime_minutes/)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/use_default_keylifetime/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/)
