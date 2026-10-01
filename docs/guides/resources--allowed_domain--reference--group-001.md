---
page_title: "xcsh_allowed_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain reference."
---

# xcsh_allowed_domain reference

<a id="canonical-b4f5d2267a3810dd15ca31a5e12ece95139a49ccbbc3defd7a52e44af9653777"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd44032d3c03e1598cb29ba3bed3835e5d78ff25333f6fcacfc66410c78103b1"></a>

## Property reference — Property reference / 2b066732bdc7 / 2

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
- Property reference

<a id="canonical-74cf43e689fda075896095f7e27986d6354059f70b88183dfaf95f1a025f4763"></a>

## Direct properties — Property reference / 2b066732bdc7 / 3

<a id="canonical-b92e108d2c7ff60af4c0f8bb981bc710f5557acc3e6756d3990713e0ed1cd24a"></a>

<a id="canonical-b434ae44f9bf16972c034af2374d9ce31917408cf6cd2f0dc78c1211d55be818"></a>

## allowed_domain property — Property reference / 2b066732bdc7 / 4

Type: `"string"`. Required.

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.

Upstream description:

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.
Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-b11ad97ba060c0c8812cf7744546c89c46fff3d05d935009efdfa82d12488a8c"></a>

<a id="canonical-c3f0281a55aa28b9d7af58e1f1a003713bc6d4047d7f9a5b1cdc5dd6ca104c4b"></a>

## annotations property — Property reference / 2b066732bdc7 / 5

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

<a id="canonical-528c39270faf58f2d5332ea0d269615fd758b88a1193324d5cd1fb869b090adc"></a>

<a id="canonical-d0e1c25cfbddbed4f426ab2de4497960e99b67e69c65ca6e41622bc3bc1f8747"></a>

## description property — Property reference / 2b066732bdc7 / 6

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

<a id="canonical-39c17961d6a9593fbf83f6b9253aebd1c4af06c5bc83661eb316b5ee6eecec58"></a>

<a id="canonical-2db180c8820e1cf2858de7187c266055133d808f45b1141fe7a65d14090e004e"></a>

## disable property — Property reference / 2b066732bdc7 / 7

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

<a id="canonical-6462d66f99ff95de60c7d2f8f32a8ac270c7dc9e4fdb11c612951fc544eedfd8"></a>

<a id="canonical-17676eac18154ebc5667600dd0a31c4c2a2a85ff887165235aeb34ea17d89ef4"></a>

## id property — Property reference / 2b066732bdc7 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5fe203a730b4a14f91c25a0f21be08e3f4885fee34ce15d0054adeec77f63eac"></a>

<a id="canonical-c057b894ffab8b56887f9cd3b8b8da2e73848edc82c4566658e71a01c49852e1"></a>

## labels property — Property reference / 2b066732bdc7 / 9

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

<a id="canonical-b7eabe3eabcf72f4b65a6f6c4d114741fb514a9922d7e25ca4b3ad5f7374a2aa"></a>

<a id="canonical-05ed66ef0ed6831db93918beb9ac8652e79437c22bd4dbf4c63146c724218107"></a>

## name property — Property reference / 2b066732bdc7 / 10

Type: `"string"`. Required.

Name of the Allowed Domain. Must be unique within the namespace.

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

<a id="canonical-535fdb5a85728d61c421a2e657476e4bcc74eb96595bfd149e0d935438bdeefd"></a>

<a id="canonical-1ebd9fb42ebfc8b2b4405b7e466e743e7161ef3cac2d810df8c2883f40cf36ad"></a>

## namespace property — Property reference / 2b066732bdc7 / 11

Type: `"string"`. Required.

Namespace where the Allowed Domain is created.

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

- [timeouts](resources--allowed_domain--reference--group-001.md#canonical-1f4c7de6d11fafe5a60c7358873da13c37c3304824dbdc85b5de263bc437a2d6): complete subsection reference.

<a id="canonical-10cc69bb4984964b1162e13ed0edb8b6478a616addf3f67cb88066f587b52006"></a>

## All schema paths — Property reference / 2b066732bdc7 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_domain` | [allowed_domain](resources--allowed_domain--reference--group-001.md#canonical-b92e108d2c7ff60af4c0f8bb981bc710f5557acc3e6756d3990713e0ed1cd24a) |
| `annotations` | [annotations](resources--allowed_domain--reference--group-001.md#canonical-b11ad97ba060c0c8812cf7744546c89c46fff3d05d935009efdfa82d12488a8c) |
| `description` | [description](resources--allowed_domain--reference--group-001.md#canonical-528c39270faf58f2d5332ea0d269615fd758b88a1193324d5cd1fb869b090adc) |
| `disable` | [disable](resources--allowed_domain--reference--group-001.md#canonical-39c17961d6a9593fbf83f6b9253aebd1c4af06c5bc83661eb316b5ee6eecec58) |
| `id` | [id](resources--allowed_domain--reference--group-001.md#canonical-6462d66f99ff95de60c7d2f8f32a8ac270c7dc9e4fdb11c612951fc544eedfd8) |
| `labels` | [labels](resources--allowed_domain--reference--group-001.md#canonical-5fe203a730b4a14f91c25a0f21be08e3f4885fee34ce15d0054adeec77f63eac) |
| `name` | [name](resources--allowed_domain--reference--group-001.md#canonical-b7eabe3eabcf72f4b65a6f6c4d114741fb514a9922d7e25ca4b3ad5f7374a2aa) |
| `namespace` | [namespace](resources--allowed_domain--reference--group-001.md#canonical-535fdb5a85728d61c421a2e657476e4bcc74eb96595bfd149e0d935438bdeefd) |
| `timeouts` | [timeouts](resources--allowed_domain--reference--group-001.md#canonical-df139cc6b3d988bea01704894c19ef8b09f0a48cc5f349eaa5df847023233b6c) |
| `timeouts.create` | [timeouts.create](resources--allowed_domain--reference--group-001.md#canonical-76bb69eb7836437edff69b09b38157421a608b4e1dbc93881530964759316d09) |
| `timeouts.delete` | [timeouts.delete](resources--allowed_domain--reference--group-001.md#canonical-9fe98d145f26b027da2d6945baafaaf1dd9d2ab75b856902d360174849e90ce4) |
| `timeouts.read` | [timeouts.read](resources--allowed_domain--reference--group-001.md#canonical-bf4bc724f78c040b187a6f622c51edb5544ff056454f1ccb88f8cbb3f377584c) |
| `timeouts.update` | [timeouts.update](resources--allowed_domain--reference--group-001.md#canonical-9b2daaff1149c90aa44b42b42e32f2092825e5df0c2e277a7926dd76ff1fdc5e) |

<a id="canonical-cb944b0a334574a5e42c7694d55d1fdbf6b50fcc8a1e22ba68f24e072af5aa81"></a>

## Next pages — Property reference / 2b066732bdc7 / 13

- [timeouts](resources--allowed_domain--reference--group-001.md#canonical-1f4c7de6d11fafe5a60c7358873da13c37c3304824dbdc85b5de263bc437a2d6)
- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)

<a id="canonical-1f4c7de6d11fafe5a60c7358873da13c37c3304824dbdc85b5de263bc437a2d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e16b8782f3b0bbdcb1e198913f5b6d30ebea6f44e0b63abe9518e296030a89e2"></a>

## timeouts — timeouts / 4ae2e27e80b7 / 2

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
- [Property reference](resources--allowed_domain--reference--group-001.md#canonical-b4f5d2267a3810dd15ca31a5e12ece95139a49ccbbc3defd7a52e44af9653777)
- timeouts

<a id="canonical-df139cc6b3d988bea01704894c19ef8b09f0a48cc5f349eaa5df847023233b6c"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3970f7cc799e7c479cef59f7bffc50ad4d3c10c5e125dc124801d63743699e6a"></a>

## Direct properties — timeouts / 4ae2e27e80b7 / 3

<a id="canonical-76bb69eb7836437edff69b09b38157421a608b4e1dbc93881530964759316d09"></a>

<a id="canonical-46c43a7fde05401fe259ae94642b2c8a0266b0dc2e920d4c67f54b6deb418521"></a>

## create property — timeouts / 4ae2e27e80b7 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9fe98d145f26b027da2d6945baafaaf1dd9d2ab75b856902d360174849e90ce4"></a>

<a id="canonical-75ac79fb62f50364ae0ba6d105fc55910d6cc7c176285ec0c3033c17a5a8c7f5"></a>

## delete property — timeouts / 4ae2e27e80b7 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-bf4bc724f78c040b187a6f622c51edb5544ff056454f1ccb88f8cbb3f377584c"></a>

<a id="canonical-bd3b60471f32181b294b88a026d3e60726226dbe772b8d864765726e9fb5c7c8"></a>

## read property — timeouts / 4ae2e27e80b7 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9b2daaff1149c90aa44b42b42e32f2092825e5df0c2e277a7926dd76ff1fdc5e"></a>

<a id="canonical-2910e2e8aea9a6d7149808d7d40d11fede23ff3b9cff08e87d3fd2e171b944bc"></a>

## update property — timeouts / 4ae2e27e80b7 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-94bec10eb843c9cc2fc37ce7a4f6a778e9ab46b9b4f53d03a55163b340b347ff"></a>

## Next pages — timeouts / 4ae2e27e80b7 / 8

- [Property reference](resources--allowed_domain--reference--group-001.md#canonical-b4f5d2267a3810dd15ca31a5e12ece95139a49ccbbc3defd7a52e44af9653777)
- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
