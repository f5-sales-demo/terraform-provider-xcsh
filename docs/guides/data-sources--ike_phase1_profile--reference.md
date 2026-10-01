---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 13443, "body_sha256": "sha256:2622e3671d448150dbb51e84bcbb85b92eb3df10b72aa720d84f02db8c5796bf", "canonical_id": "xcsh-docs:data-sources:ike_phase1_profile:reference", "child_ids": ["xcsh-docs:data-sources:ike_phase1_profile:properties:ike_keylifetime_hours", "xcsh-docs:data-sources:ike_phase1_profile:properties:ike_keylifetime_minutes", "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_disabled", "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_days", "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_hours", "xcsh-docs:data-sources:ike_phase1_profile:properties:use_default_keylifetime"], "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:reference", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:fundamentals", "path": "docs/guides/data-sources--ike_phase1_profile--reference.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md)
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

<a id="schema-authentication_algos"></a>

### authentication_algos property

Type: `["list", "string"]`. Computed.

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

Type: `"string"`. Computed.

Description of the IKEPhase1Profile.

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

<a id="schema-dh_group"></a>

### dh_group property

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile.

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

<a id="schema-encryption_algos"></a>

### encryption_algos property

Type: `["list", "string"]`. Computed.

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

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--properties--ike_keylifetime_hours.md): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--properties--ike_keylifetime_minutes.md): complete subsection reference.

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

Name of the IKEPhase1Profile.

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

Namespace where the IKEPhase1Profile exists.

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

<a id="schema-prf"></a>

### prf property

Type: `["list", "string"]`. Computed.

\[Enum: PRF\_DEFAULT|PRFSHA256|PRFSHA384|PRFSHA512\] PseudoRandomFunction. Select
PseudoRandomFunction for IKE SA. Possible values are \`PRF\_DEFAULT\`, \`PRFSHA256\`, \`PRFSHA384\`,
\`PRFSHA512\`. Defaults to \`PRF\_DEFAULT\`.

Upstream description:

Select PseudoRandomFunction for IKE SA.

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

- [reauth_disabled](data-sources--ike_phase1_profile--properties--reauth_disabled.md): complete subsection reference.

- [reauth_timeout_days](data-sources--ike_phase1_profile--properties--reauth_timeout_days.md): complete subsection reference.

- [reauth_timeout_hours](data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md): complete subsection reference.

- [use_default_keylifetime](data-sources--ike_phase1_profile--properties--use_default_keylifetime.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike_phase1_profile--reference.md#schema-annotations) |
| `authentication_algos` | [authentication_algos](data-sources--ike_phase1_profile--reference.md#schema-authentication_algos) |
| `description` | [description](data-sources--ike_phase1_profile--reference.md#schema-description) |
| `dh_group` | [dh_group](data-sources--ike_phase1_profile--reference.md#schema-dh_group) |
| `encryption_algos` | [encryption_algos](data-sources--ike_phase1_profile--reference.md#schema-encryption_algos) |
| `id` | [id](data-sources--ike_phase1_profile--reference.md#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike_phase1_profile--properties--ike_keylifetime_hours.md#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike_phase1_profile--properties--ike_keylifetime_hours.md#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike_phase1_profile--properties--ike_keylifetime_minutes.md#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike_phase1_profile--properties--ike_keylifetime_minutes.md#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](data-sources--ike_phase1_profile--reference.md#schema-labels) |
| `name` | [name](data-sources--ike_phase1_profile--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--ike_phase1_profile--reference.md#schema-namespace) |
| `prf` | [prf](data-sources--ike_phase1_profile--reference.md#schema-prf) |
| `reauth_disabled` | [reauth_disabled](data-sources--ike_phase1_profile--properties--reauth_disabled.md#section) |
| `reauth_timeout_days` | [reauth_timeout_days](data-sources--ike_phase1_profile--properties--reauth_timeout_days.md#section) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](data-sources--ike_phase1_profile--properties--reauth_timeout_days.md#schema-reauth_timeout_days--duration) |
| `reauth_timeout_hours` | [reauth_timeout_hours](data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md#section) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md#schema-reauth_timeout_hours--duration) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike_phase1_profile--properties--use_default_keylifetime.md#section) |

## Next pages

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--properties--ike_keylifetime_hours.md)
- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--properties--ike_keylifetime_minutes.md)
- [reauth_disabled](data-sources--ike_phase1_profile--properties--reauth_disabled.md)
- [reauth_timeout_days](data-sources--ike_phase1_profile--properties--reauth_timeout_days.md)
- [reauth_timeout_hours](data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md)
- [use_default_keylifetime](data-sources--ike_phase1_profile--properties--use_default_keylifetime.md)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md)
