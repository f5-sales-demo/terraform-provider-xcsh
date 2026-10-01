---
page_title: "xcsh_api_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery reference."
---

# xcsh_api_discovery reference

<a id="canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a83b3950bd4cc596b2010b5cb20775780921dc5377cf1135c622b7a504149021"></a>

## Property reference — Property reference / 5d794140beb1 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- Property reference

<a id="canonical-d21752457a03556112c0016eed22b962b19e14185f78966e7ef29c27707e31be"></a>

## Direct properties — Property reference / 5d794140beb1 / 3

<a id="canonical-bb2952315c4209adea0213adab59602d55dcb4b57fd72860fc513d9b0020959d"></a>

<a id="canonical-fceff7f44fd18f6be88241335fc847a0dad609fd3fbd4630a1acf089b4bb2b50"></a>

## annotations property — Property reference / 5d794140beb1 / 4

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

- [custom_auth_types](resources--api_discovery--reference--group-001.md#canonical-a665f31bf74e531f7e4763d20bd1087c5d711f4d0d584c4828319fecf024d789): complete subsection reference.

<a id="canonical-a13c57240c21468b6eaf10359f373eecc4369c737ca1f9a361e508eef1badbb1"></a>

<a id="canonical-42e83d3d63de1aa4ff6c331b8205281f642898b70d7619a46d8873dbdab7b793"></a>

## description property — Property reference / 5d794140beb1 / 5

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

<a id="canonical-97e4f97732e91947995b2bb8deff14db464466f7d3ab721db0b61a74fd334184"></a>

<a id="canonical-46fa6605f425c1ece0fa39607c2fea9a7bc83e9863a7d02bff8e3e0a041ca626"></a>

## disable property — Property reference / 5d794140beb1 / 6

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

<a id="canonical-2b45365c40cdf2bc8e5e2f9de778f160f043f880d0cd02940bdb124e71f617e3"></a>

<a id="canonical-2c50f98f8778683ae8053076f0f8a221bad7160764cfc71c6febb4c3e1674876"></a>

## id property — Property reference / 5d794140beb1 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b1a1bb68393beda923b4edb7d6acfa3c41fe3075762a3335cdcd8b936627274d"></a>

<a id="canonical-9f1c11efb3cc0d90a8a1acefacf69f576d23b37662aecacec8a82ccf602e46f7"></a>

## labels property — Property reference / 5d794140beb1 / 8

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

<a id="canonical-d6ec85efc35f560e41717c9f4a089f6f11e438f3b277bbe001b75b1279f07104"></a>

<a id="canonical-5de09e00bf1ac672c5f79401301dc2e63e82c7bbae2657b3708344182d8d69e8"></a>

## name property — Property reference / 5d794140beb1 / 9

Type: `"string"`. Required.

Name of the API Discovery. Must be unique within the namespace.

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

<a id="canonical-da7b5d695cc0d17c98cf4f72b7322df54f0cbcdddd49c7a5f2f934b3bc9ede84"></a>

<a id="canonical-1e0e5e6f2b16e407bfd64acff3f495e54e1a2c2ad557ca9e700c88f99ad9344d"></a>

## namespace property — Property reference / 5d794140beb1 / 10

Type: `"string"`. Required.

Namespace where the API Discovery is created.

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

- [timeouts](resources--api_discovery--reference--group-001.md#canonical-3c0afe996f3d4c432c32163b150b30e8f135cb9bbaf05d80b644a76df1490621): complete subsection reference.

- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da): complete subsection reference.

<a id="canonical-cab39b45e606ae237862c4546b6acd51a9a96ed567388b5bcd9dd231b1f33a1e"></a>

## All schema paths — Property reference / 5d794140beb1 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_discovery--reference--group-001.md#canonical-bb2952315c4209adea0213adab59602d55dcb4b57fd72860fc513d9b0020959d) |
| `custom_auth_types` | [custom_auth_types](resources--api_discovery--reference--group-001.md#canonical-f895eb34bb63fdbeeb750a67ce6baadf5a850386d0410dc83e99ad6d8f3de32b) |
| `custom_auth_types.parameter_name` | [custom_auth_types.parameter_name](resources--api_discovery--reference--group-001.md#canonical-75326a4fedd25f78b3d896e7054af462e0545a0b3025273392fddab5021c6a56) |
| `custom_auth_types.parameter_type` | [custom_auth_types.parameter_type](resources--api_discovery--reference--group-001.md#canonical-2cf7b9e64449fbd4dd1f247f45b01b828ab2a3d80429e5579aa8e876efed6c52) |
| `description` | [description](resources--api_discovery--reference--group-001.md#canonical-a13c57240c21468b6eaf10359f373eecc4369c737ca1f9a361e508eef1badbb1) |
| `disable` | [disable](resources--api_discovery--reference--group-001.md#canonical-97e4f97732e91947995b2bb8deff14db464466f7d3ab721db0b61a74fd334184) |
| `id` | [id](resources--api_discovery--reference--group-001.md#canonical-2b45365c40cdf2bc8e5e2f9de778f160f043f880d0cd02940bdb124e71f617e3) |
| `labels` | [labels](resources--api_discovery--reference--group-001.md#canonical-b1a1bb68393beda923b4edb7d6acfa3c41fe3075762a3335cdcd8b936627274d) |
| `name` | [name](resources--api_discovery--reference--group-001.md#canonical-d6ec85efc35f560e41717c9f4a089f6f11e438f3b277bbe001b75b1279f07104) |
| `namespace` | [namespace](resources--api_discovery--reference--group-001.md#canonical-da7b5d695cc0d17c98cf4f72b7322df54f0cbcdddd49c7a5f2f934b3bc9ede84) |
| `timeouts` | [timeouts](resources--api_discovery--reference--group-001.md#canonical-5192c5df2f835dc7fd1bd26e3035e203e0bb6ab2e31f326362af80c5b1481c42) |
| `timeouts.create` | [timeouts.create](resources--api_discovery--reference--group-001.md#canonical-dd72e5f29377e72691878240c2c933a2a2e5e40a964b9a4a20fb106021e80fff) |
| `timeouts.delete` | [timeouts.delete](resources--api_discovery--reference--group-001.md#canonical-f963e093ed981e71c03d9dd9241ba5efdf58feb62b3914bf5888e0211aa5739a) |
| `timeouts.read` | [timeouts.read](resources--api_discovery--reference--group-001.md#canonical-07b9550ab45d9cf6dd7fd151806cde2f9a3062ed0c9f5a6e8be9aa950e65be4e) |
| `timeouts.update` | [timeouts.update](resources--api_discovery--reference--group-001.md#canonical-1b6e4e484d071bca3ad3cb16cf6e4a8144b2ba97a7658b212cf308f5005633cc) |
| `user_defined_api_discovery_policy` | [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-31430f5b353f1e44f75d34e5022a8e6c451d93bc511df67f4a7b8a1013f7690c) |
| `user_defined_api_discovery_policy.discovery_rules` | [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-bb8b38596024ce4af53787d313968a0e234e186eba2128dac17372f238ffd800) |
| `user_defined_api_discovery_policy.discovery_rules.labels` | [user_defined_api_discovery_policy.discovery_rules.labels](resources--api_discovery--reference--group-001.md#canonical-5ea8588e0910c47612d702e56858df917fe6e1a890df24274be68697e5995912) |
| `user_defined_api_discovery_policy.discovery_rules.metadata` | [user_defined_api_discovery_policy.discovery_rules.metadata](resources--api_discovery--reference--group-001.md#canonical-87dcd9bd8c97d19286dc2213ca2410f14e70861d869e39ae1045a0eabd4a6780) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` | [user_defined_api_discovery_policy.discovery_rules.metadata.description_spec](resources--api_discovery--reference--group-001.md#canonical-b2e4361f880e929e0144418ed18cf3c93bdadce2e74d854cf1eda5aa1c79784b) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.name` | [user_defined_api_discovery_policy.discovery_rules.metadata.name](resources--api_discovery--reference--group-001.md#canonical-6e3a05d0ae380bbd15e4f65fe207424b6b7e69d151a7c2336b24e91f19df21e6) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties` | [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-fab47f91141bb36b45cc6380adf6ccc6883b8170855d12ae67aaf513079615c4) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-0adb592cf3f69e4f87b0a6142e842002c909f9d71cfb69a6967352456f9a2a6a) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](resources--api_discovery--reference--group-001.md#canonical-1449627de64b9d903d788c127a46c0d7a2360f04bd5cec748be336980209a714) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](resources--api_discovery--reference--group-001.md#canonical-579413246f55e319a2e82b45884027069ac6ebae89b3fec144c1ac809f962885) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](resources--api_discovery--reference--group-001.md#canonical-ef2290efb5b85bd03f9abee3cee66f02e312f32a37f4eb104cca8088871ebdce) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name](resources--api_discovery--reference--group-001.md#canonical-ef12866b26ea380074f4ea4ed8374ae4e631d3040e53939b2714b1b3756ab65e) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location](resources--api_discovery--reference--group-001.md#canonical-e6b6621bc2787799849963e576423496c87f319bd02a6da8b0ded6d4ac7f4c5f) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type](resources--api_discovery--reference--group-001.md#canonical-125c671a3d756b485af70f1dd5f3f8bb37ac5fb5b59b6ce7015d8d0b4564663a) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value](resources--api_discovery--reference--group-001.md#canonical-89bf675842de5ed22a04df2e1254d1e725e0a864e7f0dd0da7043448078e47ba) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](resources--api_discovery--reference--group-001.md#canonical-eca9aa09abcd33b21c4755e58e1998aab6cdde5c2aa4c04a3c16a3d8dba9c2e2) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern](resources--api_discovery--reference--group-001.md#canonical-d9b6b12529bb06e3cf6799d0a4551f1a153f94918d8a3f7c064c5a93a4126829) |
| `user_defined_api_discovery_policy.exclusive` | [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-cbec79f92b65ead2fc36f394bd4cfaa39c839af9f5670cb7aa1a01d62bc8c129) |
| `user_defined_api_discovery_policy.exclusive.archive` | [user_defined_api_discovery_policy.exclusive.archive](resources--api_discovery--reference--group-001.md#canonical-29b0943f65d31292692ef35367e8c34972b053e3e4d503cae748b72459d9d50d) |
| `user_defined_api_discovery_policy.exclusive.ignore` | [user_defined_api_discovery_policy.exclusive.ignore](resources--api_discovery--reference--group-001.md#canonical-38dfafee41402c831aa5d0fd455380adbfa32d2708ee330efb9e469c0c36cd08) |
| `user_defined_api_discovery_policy.inclusive` | [user_defined_api_discovery_policy.inclusive](resources--api_discovery--reference--group-001.md#canonical-520c55a97de2900b073e946ed60cf14b2cd569b6c6ab5af9ff8227127b5ef7ed) |

<a id="canonical-280b8456e0fbf89bda327219c3a02936418381d16d5f72bb31b47512e96ea4f0"></a>

## Next pages — Property reference / 5d794140beb1 / 12

- [custom_auth_types](resources--api_discovery--reference--group-001.md#canonical-a665f31bf74e531f7e4763d20bd1087c5d711f4d0d584c4828319fecf024d789)
- [timeouts](resources--api_discovery--reference--group-001.md#canonical-3c0afe996f3d4c432c32163b150b30e8f135cb9bbaf05d80b644a76df1490621)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-a665f31bf74e531f7e4763d20bd1087c5d711f4d0d584c4828319fecf024d789"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-070799d683828992b9d226cfd74548d47006df1340fa0b3a0babf52f7031729f"></a>

## custom_auth_types — custom_auth_types / 4b9a99e42aff / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- custom_auth_types

<a id="canonical-f895eb34bb63fdbeeb750a67ce6baadf5a850386d0410dc83e99ad6d8f3de32b"></a>

Type: `"object"`. list nested block, Optional.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

Upstream description:

Select your custom authentication types to be detected in the API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("parameter_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  }
}
```

Terraform syntax:

```terraform
custom_auth_types {
  # Configure direct properties listed below.
}
```

<a id="canonical-8610640f455354debe218a3b87a625a2ee3c1c2eb135c28095131209d9417dce"></a>

## Direct properties — custom_auth_types / 4b9a99e42aff / 3

<a id="canonical-75326a4fedd25f78b3d896e7054af462e0545a0b3025273392fddab5021c6a56"></a>

<a id="canonical-02cb93641be9dd6c21baec93c832a01878bbee29b0d2242c4af586239749e324"></a>

## parameter_name property — custom_auth_types / 4b9a99e42aff / 4

Type: `"string"`. Optional.

Parameter Name. The authentication parameter name.

Upstream description:

The authentication parameter name.

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
    },
    "pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  }
}
```

<a id="canonical-2cf7b9e64449fbd4dd1f247f45b01b828ab2a3d80429e5579aa8e876efed6c52"></a>

<a id="canonical-56b295d961d107120176fccb9f4c1a2f42b844fe1bc5874933afb8d2ab38952f"></a>

## parameter_type property — custom_auth_types / 4b9a99e42aff / 5

Type: `"string"`. Optional.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

Upstream description:

Enumeration for authentication parameter types.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("QUERY_PARAMETER",
    "HEADER",
    "COOKIE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "QUERY_PARAMETER",
  "enum": [
    "QUERY_PARAMETER",
    "HEADER",
    "COOKIE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-03c373bade76d1d9047ba63eced4c6c69b1bdb21d73f4087734327f29031e430"></a>

## Next pages — custom_auth_types / 4b9a99e42aff / 6

- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-3c0afe996f3d4c432c32163b150b30e8f135cb9bbaf05d80b644a76df1490621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba23d9ea904bd54e845d45ef6d7ee6ef909461c871deeb3488879414c19e8d93"></a>

## timeouts — timeouts / 7627835c1459 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- timeouts

<a id="canonical-5192c5df2f835dc7fd1bd26e3035e203e0bb6ab2e31f326362af80c5b1481c42"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e6edf82b2f9bc40a16829b4032d3e4058625cdfb5dc1f5ef713f38f5f2ec560"></a>

## Direct properties — timeouts / 7627835c1459 / 3

<a id="canonical-dd72e5f29377e72691878240c2c933a2a2e5e40a964b9a4a20fb106021e80fff"></a>

<a id="canonical-ed568a2dc77f08b4da78835b20d8f33c7477346dd4e20720912268a17714675d"></a>

## create property — timeouts / 7627835c1459 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f963e093ed981e71c03d9dd9241ba5efdf58feb62b3914bf5888e0211aa5739a"></a>

<a id="canonical-72b63f968813e5c045680e8c4ef7050bb4f17b4e50e074eba3dec9b18974b370"></a>

## delete property — timeouts / 7627835c1459 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-07b9550ab45d9cf6dd7fd151806cde2f9a3062ed0c9f5a6e8be9aa950e65be4e"></a>

<a id="canonical-71e522d3de7b069a4cc6dc34df63cf5ea360b61461ba74b392051fb29880ec4b"></a>

## read property — timeouts / 7627835c1459 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1b6e4e484d071bca3ad3cb16cf6e4a8144b2ba97a7658b212cf308f5005633cc"></a>

<a id="canonical-6ddb66de2961728ae41c62a6b5f781c369cb9d9a495eea3893a1963132b42864"></a>

## update property — timeouts / 7627835c1459 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-242a028febfb47629b6d76395eda9154f15ab0e8ff0161a3a104be0ce940fdb5"></a>

## Next pages — timeouts / 7627835c1459 / 8

- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0ff86f4240c459713ac056dd2cbb0f75e1c5735f3c97e0d3f734bf659639f77"></a>

## user_defined_api_discovery_policy — user_defined_api_discovery_policy / a1faba526f6a / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- user_defined_api_discovery_policy

<a id="canonical-31430f5b353f1e44f75d34e5022a8e6c451d93bc511df67f4a7b8a1013f7690c"></a>

Type: `"object"`. single nested block, Optional.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusive",
    "inclusive")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

Terraform syntax:

```terraform
user_defined_api_discovery_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6a823a7826380842e80d2e5fdbf474c1b39f11801600c7cb00c2a8c6c00ef1a"></a>

## Direct properties — user_defined_api_discovery_policy / a1faba526f6a / 3

- [discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37): complete subsection reference.

- [exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2): complete subsection reference.

- [inclusive](resources--api_discovery--reference--group-001.md#canonical-a7398775e05bea14c9a617c2451e7da60f0d3d47385a42bc309eefbaa4de46bf): complete subsection reference.

<a id="canonical-49d03026efe5b9696382e7e7e5c2b6a454acdad0b6bd4688ddc96ed824aa5b72"></a>

## Next pages — user_defined_api_discovery_policy / a1faba526f6a / 4

- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2)
- [user_defined_api_discovery_policy.inclusive](resources--api_discovery--reference--group-001.md#canonical-a7398775e05bea14c9a617c2451e7da60f0d3d47385a42bc309eefbaa4de46bf)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-126d4bd784fbf6770462b4bab6cb150721d411f49b65cde4c729dbe5d34a1a30"></a>

## user_defined_api_discovery_policy.discovery_rules — user_defined_api_discovery_policy.discovery_rules / 6909a6f35fad / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- user_defined_api_discovery_policy.discovery_rules

<a id="canonical-bb8b38596024ce4af53787d313968a0e234e186eba2128dac17372f238ffd800"></a>

Type: `"object"`. list nested block, Optional.

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
discovery_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-fbb2fae87119f3865697b815c8f9a0ed616d94dca8f6dfbcd3df6a2a0d355ec4"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules / 6909a6f35fad / 3

- [labels](resources--api_discovery--reference--group-001.md#canonical-93880d9f0fb1b10b1fbcb86f4ea390b9fd80a9ccbc952f3c346d90d51a646a2a): complete subsection reference.

- [metadata](resources--api_discovery--reference--group-001.md#canonical-90523bff2e453de634d884f297e7555f5601cd23ed41ac05b790323bda368027): complete subsection reference.

- [rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce): complete subsection reference.

<a id="canonical-02ad5467f17a4d8bd2cf8ee98e5e0a2cff7c303615937b77349c7bcb4823f590"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules / 6909a6f35fad / 4

- [user_defined_api_discovery_policy.discovery_rules.labels](resources--api_discovery--reference--group-001.md#canonical-93880d9f0fb1b10b1fbcb86f4ea390b9fd80a9ccbc952f3c346d90d51a646a2a)
- [user_defined_api_discovery_policy.discovery_rules.metadata](resources--api_discovery--reference--group-001.md#canonical-90523bff2e453de634d884f297e7555f5601cd23ed41ac05b790323bda368027)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-93880d9f0fb1b10b1fbcb86f4ea390b9fd80a9ccbc952f3c346d90d51a646a2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63d1fd8b0d44e1fdc57f2c744cf6a4c8ac0f39b7837dcde2d8354c6a04849e53"></a>

## user_defined_api_discovery_policy.discovery_rules.labels — user_defined_api_discovery_policy.discovery_rules.labels / 0a8284afeb11 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- user_defined_api_discovery_policy.discovery_rules.labels

<a id="canonical-5ea8588e0910c47612d702e56858df917fe6e1a890df24274be68697e5995912"></a>

Type: `"object"`. single nested block, Optional.

Map of string keys and values that can be used to organize and categorize the rule.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-068deb002a80600adea8b35a264888165094a76c3f7395ae830fbc30fc3d3233"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.labels / 0a8284afeb11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8d993bbe0eb1ed2dff1045e89a59c3de37cc54bddc5f60965e11a1244448816"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.labels / 0a8284afeb11 / 4

- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-90523bff2e453de634d884f297e7555f5601cd23ed41ac05b790323bda368027"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fb6e2cc55cedd908a95ea5ab455e1c77a501de373fc9009eac21fadd6469d77"></a>

## user_defined_api_discovery_policy.discovery_rules.metadata — user_defined_api_discovery_policy.discovery_rules.metadata / 460753a85a47 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- user_defined_api_discovery_policy.discovery_rules.metadata

<a id="canonical-87dcd9bd8c97d19286dc2213ca2410f14e70861d869e39ae1045a0eabd4a6780"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bf29ae0408a2b012ba4016e734e54cd1c8cdc9dc21f72c3282d728eda6481a5"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.metadata / 460753a85a47 / 3

<a id="canonical-b2e4361f880e929e0144418ed18cf3c93bdadce2e74d854cf1eda5aa1c79784b"></a>

<a id="canonical-5651cbc996f293616f5281934906e415a6d78a84391a98277e325dda13fb0e51"></a>

## description_spec property — user_defined_api_discovery_policy.discovery_rules.metadata / 460753a85a47 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-6e3a05d0ae380bbd15e4f65fe207424b6b7e69d151a7c2336b24e91f19df21e6"></a>

<a id="canonical-18822e5e0522d35b6414d9895887ce8021aa2c85885d13bbc5acb048113535fb"></a>

## name property — user_defined_api_discovery_policy.discovery_rules.metadata / 460753a85a47 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-60388e05e4bfcec7c35226ffb0edcdce3fbba75c789ced310103a278e1d6fed3"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.metadata / 460753a85a47 / 6

- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3dd14f9bda1f373ca01c5b854e5b00b20178f86d4750e8d28dac224a52dd6ed"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties — user_defined_api_discovery_policy.discovery_rules.rule_properties / 3ce7697dca8a / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="canonical-fab47f91141bb36b45cc6380adf6ccc6883b8170855d12ae67aaf513079615c4"></a>

Type: `"object"`. single nested block, Optional.

Determines whether matching endpoints are included in API Discovery or excluded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusion",
    "inclusion"),
  validators.ConflictingObjectAttributes("http_header_criteria",
    "pattern")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-criteria": "[\"http_header_criteria\",\"pattern\"]",
  "x-ves-oneof-field-rule_type_choice": "[\"exclusion\",\"inclusion\"]"
}
```

Terraform syntax:

```terraform
rule_properties {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0497ebf68b9ea7982e6fcd353577a35865edd86b82f4ab40a80424cd0af99ee"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties / 3ce7697dca8a / 3

- [exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2): complete subsection reference.

- [http_header_criteria](resources--api_discovery--reference--group-001.md#canonical-17fd45f8b70a4e7023236f090d019b7a4108bca3e7587cf868ff9d7255131f41): complete subsection reference.

- [inclusion](resources--api_discovery--reference--group-001.md#canonical-b7f9d8edbcc9f6b6430fa9cf7376f6ddeb5261778f590224e51ce9f347a5b829): complete subsection reference.

<a id="canonical-d9b6b12529bb06e3cf6799d0a4551f1a153f94918d8a3f7c064c5a93a4126829"></a>

<a id="canonical-e78cc919bf7cb3442efdc41e138da69434f122b0d95a94b72c2324aec79acaf2"></a>

## pattern property — user_defined_api_discovery_policy.discovery_rules.rule_properties / 3ce7697dca8a / 4

Type: `"string"`. Optional.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Upstream description:

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-d4a912fa579d04ec2bbf6c54a96d6fe99bfd24b0577bb13d37561ae8b557e154"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties / 3ce7697dca8a / 5

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](resources--api_discovery--reference--group-001.md#canonical-17fd45f8b70a4e7023236f090d019b7a4108bca3e7587cf868ff9d7255131f41)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](resources--api_discovery--reference--group-001.md#canonical-b7f9d8edbcc9f6b6430fa9cf7376f6ddeb5261778f590224e51ce9f347a5b829)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45f25cf4dc06b445c2929dfe48030a370cd479fe7c23377dd2b19aaf691b404a"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 84dfa0e4d1fa / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

<a id="canonical-0adb592cf3f69e4f87b0a6142e842002c909f9d71cfb69a6967352456f9a2a6a"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

Terraform syntax:

```terraform
exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb7f2f16bffd053eeb9085a143dda092308d960cfd1150bfa19a5888ad8ab151"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 84dfa0e4d1fa / 3

- [archive](resources--api_discovery--reference--group-001.md#canonical-cd7a4807c813e2b22dc39cdbe41df8164ca371e531f4d3271eff43e4615929dc): complete subsection reference.

- [ignore](resources--api_discovery--reference--group-001.md#canonical-c88714dd5f9e563f245cb0a297ec364fbec316e9ba0e019471902764eca0885e): complete subsection reference.

<a id="canonical-a53b8083f13ced737471e08a5115eebb96e427656c3f46e1974d69c44bd0da1e"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 84dfa0e4d1fa / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](resources--api_discovery--reference--group-001.md#canonical-cd7a4807c813e2b22dc39cdbe41df8164ca371e531f4d3271eff43e4615929dc)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](resources--api_discovery--reference--group-001.md#canonical-c88714dd5f9e563f245cb0a297ec364fbec316e9ba0e019471902764eca0885e)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-cd7a4807c813e2b22dc39cdbe41df8164ca371e531f4d3271eff43e4615929dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d26e95bbf6e94119fa0b2f95310ab15d9b12cfc5198fd4906174827003e83d0"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / 9cc8452c56f9 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive

<a id="canonical-1449627de64b9d903d788c127a46c0d7a2360f04bd5cec748be336980209a714"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
archive = {}
```

<a id="canonical-05d428e5c4cafba2719b0bf708d6ff29299aa7efd439708f97339e7c3a977240"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / 9cc8452c56f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a5bb9b23c7af9d1724ac021aca9a3fdeab4b16256158d48c73e0fc87e1084b3"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / 9cc8452c56f9 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-c88714dd5f9e563f245cb0a297ec364fbec316e9ba0e019471902764eca0885e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25e3055da42100c6455b8ac04fc595d58771f31ee598489e26b6cbb5b78db419"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 6161abf5219f / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore

<a id="canonical-579413246f55e319a2e82b45884027069ac6ebae89b3fec144c1ac809f962885"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
ignore = {}
```

<a id="canonical-fd8966e89e2acfeb10a3a0d4164b14718c87e6fddada084080bafa618047f34c"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 6161abf5219f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2473065795d5a4596b71447562f2d471d7be8a90c1e7683c07ee621027ae917b"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 6161abf5219f / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-9610133a9d11c36e977c5fb78256206f648bacc3b70c30bd66e5eee8e13a97e2)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-17fd45f8b70a4e7023236f090d019b7a4108bca3e7587cf868ff9d7255131f41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43b436598c7ddc1833e5a272853f9997defb0f172132784a39fa25d621b52684"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

<a id="canonical-ef2290efb5b85bd03f9abee3cee66f02e312f32a37f4eb104cca8088871ebdce"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header criteria.

Upstream description:

Criteria for matching HTTP headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("field_name",
    "value")}
```

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

Terraform syntax:

```terraform
http_header_criteria {
  # Configure direct properties listed below.
}
```

<a id="canonical-6eebfc35d9ce0045baea5e5f5f8c19b10d1ea9135028596fa84d14af6f1a4994"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 3

<a id="canonical-ef12866b26ea380074f4ea4ed8374ae4e631d3040e53939b2714b1b3756ab65e"></a>

<a id="canonical-1beadbbc0501c092c3ab6fb864d4d5c3b8309cb293a23fd481f3cbc235512227"></a>

## field_name property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 4

Type: `"string"`. Optional.

HTTP Header Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-e6b6621bc2787799849963e576423496c87f319bd02a6da8b0ded6d4ac7f4c5f"></a>

<a id="canonical-3744e446ac7d46f383b05feccf80fa6e2a0cafd0f85cd40f2a1e7b226e650da7"></a>

## location property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 5

Type: `"string"`. Optional.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

Upstream description:

Specifies whether the rule criteria should be evaluated against request or response

Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back
to the client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("REQUEST",
    "RESPONSE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "REQUEST",
  "enum": [
    "REQUEST",
    "RESPONSE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-125c671a3d756b485af70f1dd5f3f8bb37ac5fb5b59b6ce7015d8d0b4564663a"></a>

<a id="canonical-3e43c5759a1f2a20221ad28d2af0ec800b3df48ceef4e970d35f509da0c07ed5"></a>

## match_type property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 6

Type: `"string"`. Optional.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

Upstream description:

Specifies how the value should be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EXACT_MATCH",
    "SUBSTRING",
    "REGEX"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EXACT_MATCH",
  "enum": [
    "EXACT_MATCH",
    "SUBSTRING",
    "REGEX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-89bf675842de5ed22a04df2e1254d1e725e0a864e7f0dd0da7043448078e47ba"></a>

<a id="canonical-a0541412efe4bfe4a9a2c4ff6b8dda96f003ec71ef4b62d4e58206ee4d1e9490"></a>

## value property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 7

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2c47d96feab317aca8941dc31e17491c80e6633c2ce3ca9abe78cec7d78b9c30"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 9808e473a8c0 / 8

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-b7f9d8edbcc9f6b6430fa9cf7376f6ddeb5261778f590224e51ce9f347a5b829"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813f0950edb46d8d7524407b43a2ec0a2740fc3bb2a39839f9dc77d97d6d851f"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 0bf62b5b26aa / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-50fbe6ff910de18ef8070786641fcf48b9e53cc47a604bb48176b14204719a37)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion

<a id="canonical-eca9aa09abcd33b21c4755e58e1998aab6cdde5c2aa4c04a3c16a3d8dba9c2e2"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
inclusion = {}
```

<a id="canonical-4e774347146b16629f5fe6575ae5b511b5ea69ee44bd3a6a11b14693ee1547f1"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 0bf62b5b26aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8fedcdf201cc2da84b69cb36f2aae63ee0c96cf6e84496a61ec8e37f54077708"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 0bf62b5b26aa / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-b9f42b59311a2bdd24a82671984817d452b3f0dd56828ced1efda73def0fd2ce)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ffb59e82fc5607830f6900d59311cd8b5e10a94bedc92f6aa39eb11f0989c3"></a>

## user_defined_api_discovery_policy.exclusive — user_defined_api_discovery_policy.exclusive / ddc529fb0add / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- user_defined_api_discovery_policy.exclusive

<a id="canonical-cbec79f92b65ead2fc36f394bd4cfaa39c839af9f5670cb7aa1a01d62bc8c129"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

Terraform syntax:

```terraform
exclusive {
  # Configure direct properties listed below.
}
```

<a id="canonical-6bf4ff30b4b017893eb745452d908872ad52503b9d215c77709c51dd74180242"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive / ddc529fb0add / 3

- [archive](resources--api_discovery--reference--group-001.md#canonical-2f59022a46052b9934b4c847ea95112808c2c7230dfadf0f2648321a6237ad06): complete subsection reference.

- [ignore](resources--api_discovery--reference--group-001.md#canonical-cd7f0991016e7c397fc60549d343d319b9435bf6484113b4d4b96ce0aabbb542): complete subsection reference.

<a id="canonical-d08ace72e81e1354a25adaefe0fa22901aab7e8e12766e1b8b46f898e227c255"></a>

## Next pages — user_defined_api_discovery_policy.exclusive / ddc529fb0add / 4

- [user_defined_api_discovery_policy.exclusive.archive](resources--api_discovery--reference--group-001.md#canonical-2f59022a46052b9934b4c847ea95112808c2c7230dfadf0f2648321a6237ad06)
- [user_defined_api_discovery_policy.exclusive.ignore](resources--api_discovery--reference--group-001.md#canonical-cd7f0991016e7c397fc60549d343d319b9435bf6484113b4d4b96ce0aabbb542)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-2f59022a46052b9934b4c847ea95112808c2c7230dfadf0f2648321a6237ad06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02a816a7d2e86c5291b88bcd2cc377c730581916ae6e2a0d9ebcb505e67cf8b4"></a>

## user_defined_api_discovery_policy.exclusive.archive — user_defined_api_discovery_policy.exclusive.archive / 13bc1eaf3c85 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2)
- user_defined_api_discovery_policy.exclusive.archive

<a id="canonical-29b0943f65d31292692ef35367e8c34972b053e3e4d503cae748b72459d9d50d"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
archive = {}
```

<a id="canonical-71b6b67f628bc6185674010414f9321158bfacaca2a17bf8280f37e59d81e3f3"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive.archive / 13bc1eaf3c85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c66ca1a44c81e960212444a0cb33ed6d53d823ee53c6a91c6ba57ddfa7a36467"></a>

## Next pages — user_defined_api_discovery_policy.exclusive.archive / 13bc1eaf3c85 / 4

- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-cd7f0991016e7c397fc60549d343d319b9435bf6484113b4d4b96ce0aabbb542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40abfff9121844b474cbeced9991fd52e781b7e52a1c2d7d2e9af49a2a4ea55a"></a>

## user_defined_api_discovery_policy.exclusive.ignore — user_defined_api_discovery_policy.exclusive.ignore / f8251d6aa5d5 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2)
- user_defined_api_discovery_policy.exclusive.ignore

<a id="canonical-38dfafee41402c831aa5d0fd455380adbfa32d2708ee330efb9e469c0c36cd08"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
ignore = {}
```

<a id="canonical-02d814c9688983f492a639669b6f27075163eee24f9f2996ce30c6ef6c865659"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive.ignore / f8251d6aa5d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1cbc0f43623eb8192249ed282b52da005966eac65dfbf796c6b3ef664452a86"></a>

## Next pages — user_defined_api_discovery_policy.exclusive.ignore / f8251d6aa5d5 / 4

- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-7b54a6425e7eb691ab1c8e0a2da00c76e4b23f4e8351d75198433c50e1942fc2)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)

<a id="canonical-a7398775e05bea14c9a617c2451e7da60f0d3d47385a42bc309eefbaa4de46bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b83fc24c00f1280a7e6f90cae81c2b132b4841c812e51215d955cf929323c90"></a>

## user_defined_api_discovery_policy.inclusive — user_defined_api_discovery_policy.inclusive / a19952a8e797 / 2

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-575bca73ed45a994813627df03f207d6700fcf5d58cbd77144034ddaffad7a83)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- user_defined_api_discovery_policy.inclusive

<a id="canonical-520c55a97de2900b073e946ed60cf14b2cd569b6c6ab5af9ff8227127b5ef7ed"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
inclusive = {}
```

<a id="canonical-42d9472927fe1e00586317615c4e697e3115e24332ff5f31e8111624b3d29712"></a>

## Direct properties — user_defined_api_discovery_policy.inclusive / a19952a8e797 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7258f2c8b5225007bf52981037d0c0d46d1c16964cf4ef1acbbf481883979252"></a>

## Next pages — user_defined_api_discovery_policy.inclusive / a19952a8e797 / 4

- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-945dbdf060dc38630c8a559fe588309a3a933c419e7b7dadb04588c152f0a2da)
- [xcsh_api_discovery](../resources/api_discovery.md#canonical-fff0c8d1271786e411169d435430f2430b433fc66759a1a7fca2bad027a99fb4)
