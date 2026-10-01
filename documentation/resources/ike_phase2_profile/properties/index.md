---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 14249, "body_sha256": "sha256:2d7e9534c845170cc42348f2786e058ed17040a1601e785bf47d8368a865df95", "child_ids": ["xcsh-docs:resources:ike_phase2_profile:properties:dh_group_set", "xcsh-docs:resources:ike_phase2_profile:properties:disable_pfs", "xcsh-docs:resources:ike_phase2_profile:properties:ike_keylifetime_hours", "xcsh-docs:resources:ike_phase2_profile:properties:ike_keylifetime_minutes", "xcsh-docs:resources:ike_phase2_profile:properties:timeouts", "xcsh-docs:resources:ike_phase2_profile:properties:use_default_keylifetime"], "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase2_profile:reference", "parent_id": "xcsh-docs:resources:ike_phase2_profile:fundamentals", "path": "documentation/resources/ike_phase2_profile/properties/index.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
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

<a id="schema-authentication_algos"></a>

### authentication_algos property

Type: `["list", "string"]`. Required.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm
encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/dh_group_set/): complete subsection reference.

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

- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/disable_pfs/): complete subsection reference.

<a id="schema-encryption_algos"></a>

### encryption_algos property

Type: `["list", "string"]`. Required.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_hours/): complete subsection reference.

- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/): complete subsection reference.

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

Name of the IKE Phase2 Profile. Must be unique within the namespace.

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

Namespace where the IKE Phase2 Profile is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/): complete subsection reference.

- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/use_default_keylifetime/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-annotations) |
| `authentication_algos` | [authentication_algos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-authentication_algos) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-description) |
| `dh_group_set` | [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/dh_group_set/#section) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/dh_group_set/#schema-dh_group_set--dh_groups) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-disable) |
| `disable_pfs` | [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/disable_pfs/#section) |
| `encryption_algos` | [encryption_algos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-encryption_algos) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_hours/#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_hours/#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/#schema-timeouts--update) |
| `use_default_keylifetime` | [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/use_default_keylifetime/#section) |

## Next pages

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/dh_group_set/)
- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/disable_pfs/)
- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_hours/)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/timeouts/)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/use_default_keylifetime/)
- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
