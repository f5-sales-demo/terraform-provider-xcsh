---
page_title: "xcsh_api_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery reference."
---

# xcsh_api_discovery reference

<a id="canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ee4cfce0492eabc25e21b28f9664aff732fcc5b727656ddfc4298b829b27cfa"></a>

## Property reference — Property reference / 97b9f7c33aba / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- Property reference

<a id="canonical-b36a1214eaef1809687d9a381ead2aa3b5f2be502d76c04271f370592a5cb873"></a>

## Direct properties — Property reference / 97b9f7c33aba / 3

<a id="canonical-9550d929d2279c518d0bfdbe183d187a0fc84334bf55cd42a5f1807c8a14300b"></a>

<a id="canonical-ded955a3db0c176c6c828975ebfa240f137a64a04a13ef9c0c4442cce02296aa"></a>

## annotations property — Property reference / 97b9f7c33aba / 4

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

- [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-0c1417b41c2d178afa90ca3239194e2337a08fbc1044143f4747b3e0dfc62bc5): complete subsection reference.

<a id="canonical-267a24731331927a5f2c8eae5776b1d99723fb672269a94dfbf6a3b0b798fd32"></a>

<a id="canonical-58b6fc0dbb9756907a999309ee4b89e9e314b5bdce7c8721d644863b2110e782"></a>

## description property — Property reference / 97b9f7c33aba / 5

Type: `"string"`. Computed.

Description of the APIDiscovery.

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

<a id="canonical-ecaf97851fc64c9522a3dd3eb0eaae9f6fc13695dfdf343d3bb33063f159315a"></a>

<a id="canonical-ae920681f3e75335997fe19b2ffb0f7b95292b612a943eda016a5b4d46a2b42f"></a>

## id property — Property reference / 97b9f7c33aba / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5d328231dd8e08ff1fdce66b01aa306b71ab23e46c52544dd5efbbb2544ca015"></a>

<a id="canonical-d4e8e7d04ff7a92a6bcf2f3df1a582ba1f5cf2d2d4087ae59778677af1f37bca"></a>

## labels property — Property reference / 97b9f7c33aba / 7

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

<a id="canonical-a4aaea1af6dae86d01d1b58af306946f01d63d7392cef0093c33746e52d5e4c5"></a>

<a id="canonical-0c5183e75a322a7d652a54843fc0cff3e15d970422e03d6983711020ff654838"></a>

## name property — Property reference / 97b9f7c33aba / 8

Type: `"string"`. Required.

Name of the APIDiscovery.

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

<a id="canonical-003ebc4f042252de1e8a0fe148e3208befbd1ed543f24a287389f5d70cac4391"></a>

<a id="canonical-92b9d07058d5e86329adbac74d6705c76c38136c7384a06cb87620922ebde271"></a>

## namespace property — Property reference / 97b9f7c33aba / 9

Type: `"string"`. Required.

Namespace where the APIDiscovery exists.

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

- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f): complete subsection reference.

<a id="canonical-20cf4c5172b0b2f920bfb04788590b59e592d2f723f3a5d65a6e8ca073acf8be"></a>

## All schema paths — Property reference / 97b9f7c33aba / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_discovery--reference--group-001.md#canonical-9550d929d2279c518d0bfdbe183d187a0fc84334bf55cd42a5f1807c8a14300b) |
| `custom_auth_types` | [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-030d886b996c56b328eb9fc9373e96be1aef6d539b6ee27972915786e137485d) |
| `custom_auth_types.parameter_name` | [custom_auth_types.parameter_name](data-sources--api_discovery--reference--group-001.md#canonical-e72c2aca4c39193224f404e84b812a0681ea1bc995d9df10dc5fe98e0c95275b) |
| `custom_auth_types.parameter_type` | [custom_auth_types.parameter_type](data-sources--api_discovery--reference--group-001.md#canonical-ebe630b92879be45657da88b4fe50bc186d9e4e185d9f8a27f9be76df5ce3f10) |
| `description` | [description](data-sources--api_discovery--reference--group-001.md#canonical-267a24731331927a5f2c8eae5776b1d99723fb672269a94dfbf6a3b0b798fd32) |
| `id` | [id](data-sources--api_discovery--reference--group-001.md#canonical-ecaf97851fc64c9522a3dd3eb0eaae9f6fc13695dfdf343d3bb33063f159315a) |
| `labels` | [labels](data-sources--api_discovery--reference--group-001.md#canonical-5d328231dd8e08ff1fdce66b01aa306b71ab23e46c52544dd5efbbb2544ca015) |
| `name` | [name](data-sources--api_discovery--reference--group-001.md#canonical-a4aaea1af6dae86d01d1b58af306946f01d63d7392cef0093c33746e52d5e4c5) |
| `namespace` | [namespace](data-sources--api_discovery--reference--group-001.md#canonical-003ebc4f042252de1e8a0fe148e3208befbd1ed543f24a287389f5d70cac4391) |
| `user_defined_api_discovery_policy` | [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-a1daaf4759d2eecf4e4a26b30c2c8fdfb5b25eed1345f20c4d8330ede7d6f6bd) |
| `user_defined_api_discovery_policy.discovery_rules` | [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-cebe1ecc3966ac40908956a54e7e7683523a492b60d49fad9e1ab72a3c416f29) |
| `user_defined_api_discovery_policy.discovery_rules.labels` | [user_defined_api_discovery_policy.discovery_rules.labels](data-sources--api_discovery--reference--group-001.md#canonical-fe6d05b54da95948a3306780db1a605a27d56bd603590b204a3fe9281711ee12) |
| `user_defined_api_discovery_policy.discovery_rules.metadata` | [user_defined_api_discovery_policy.discovery_rules.metadata](data-sources--api_discovery--reference--group-001.md#canonical-49cdf0310f6cb1ac4a588a3cc319f552a444ce296b42d3cc9aa89b92def3ab16) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` | [user_defined_api_discovery_policy.discovery_rules.metadata.description_spec](data-sources--api_discovery--reference--group-001.md#canonical-cca9603d1b8b2d9babd9c524b20993a7a10bf402ef22bd370d9bd85bf614edb1) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.name` | [user_defined_api_discovery_policy.discovery_rules.metadata.name](data-sources--api_discovery--reference--group-001.md#canonical-57b65dd893f1a040d4c77aa552f7987d7fee5c314db42e9a1db5c47594c48606) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties` | [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-b9ddf479f034861952d33b90b12063bce59c036342b6043005a3a1bc520b78c2) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-55db5bd2a8c67182207db72968e901ffb54762358a509aa253c04e2da839098b) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](data-sources--api_discovery--reference--group-001.md#canonical-49048d9014d4b917ffa6c73094020827fa97441752ccac511096733bbddae843) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](data-sources--api_discovery--reference--group-001.md#canonical-d2c44bef0cac43c056d7b3f01d35616b82d77517b16dbc73e154ba8d1fe52f7f) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-f75e87e22e9fdd21812092a41f0d0de97e80e97431a45ac190cee2d14b928d2a) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name](data-sources--api_discovery--reference--group-001.md#canonical-72b5f6a67b6083f9a61629c03532a0c09f540634417dc60dced10093b368838a) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location](data-sources--api_discovery--reference--group-001.md#canonical-916c54b2798a0453d87b6e380224e233d12f5efd4205e7109c8c6049d469715e) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type](data-sources--api_discovery--reference--group-001.md#canonical-f38942eb939f94c5b718aa3f785f464e120aab68c49d8649860fc3f80cdef1aa) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value](data-sources--api_discovery--reference--group-001.md#canonical-a529e017e2a61670e2eeda41c322676149ec45c71369923598056a5f2a72dc47) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](data-sources--api_discovery--reference--group-001.md#canonical-fcbacfbdb299034beb04e4ad819fb69effa6688115de6a5df375dd59c3e05c4a) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern](data-sources--api_discovery--reference--group-001.md#canonical-f930dbf319779eb7c6e5b6481c8b0b7705ab476c209f893d768904bdb1121392) |
| `user_defined_api_discovery_policy.exclusive` | [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b63035187bf70e0ec907d80e1555b4074fa922161fea9799b6aea43731bc5194) |
| `user_defined_api_discovery_policy.exclusive.archive` | [user_defined_api_discovery_policy.exclusive.archive](data-sources--api_discovery--reference--group-001.md#canonical-d95f3acaf0e55473d8c68e9b91687ea02418b33899184ab1e183581daa39ede2) |
| `user_defined_api_discovery_policy.exclusive.ignore` | [user_defined_api_discovery_policy.exclusive.ignore](data-sources--api_discovery--reference--group-001.md#canonical-bf15e05cbf45a7fc74ef3d6a5cb7c28d5cec97112f4fbbda9d5a93b4bc371c99) |
| `user_defined_api_discovery_policy.inclusive` | [user_defined_api_discovery_policy.inclusive](data-sources--api_discovery--reference--group-001.md#canonical-78a9c420c6cd63e788e70962e0b1772a6c4697326d148fbbdd69aaf0001455f7) |

<a id="canonical-33a3839dcacfd951cbfef8f13aa54e156a7eb541a256bbc1ef64ad1dcc9d6b7c"></a>

## Next pages — Property reference / 97b9f7c33aba / 11

- [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-0c1417b41c2d178afa90ca3239194e2337a08fbc1044143f4747b3e0dfc62bc5)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-0c1417b41c2d178afa90ca3239194e2337a08fbc1044143f4747b3e0dfc62bc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-303a46dc2f5bc9f682abc2bfa185e35d9605f57f7e4027981e6768d62e4093c8"></a>

## custom_auth_types — custom_auth_types / dca934492cba / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- custom_auth_types

<a id="canonical-030d886b996c56b328eb9fc9373e96be1aef6d539b6ee27972915786e137485d"></a>

Type: `"list"`. Computed.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

Upstream description:

Select your custom authentication types to be detected in the API discovery.

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

<a id="canonical-e0dbd0b6df3098c3b2fd0abc28032923e6f74f0a13ee67959172ff34a3f43800"></a>

## Direct properties — custom_auth_types / dca934492cba / 3

<a id="canonical-e72c2aca4c39193224f404e84b812a0681ea1bc995d9df10dc5fe98e0c95275b"></a>

<a id="canonical-3afb50b745f7d405a6bbdd1839adc234258b4b72c3d1dafcd86f2903deada762"></a>

## parameter_name property — custom_auth_types / dca934492cba / 4

Type: `"string"`. Computed.

Parameter Name. The authentication parameter name.

Upstream description:

The authentication parameter name.

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

<a id="canonical-ebe630b92879be45657da88b4fe50bc186d9e4e185d9f8a27f9be76df5ce3f10"></a>

<a id="canonical-8c475b8d34efc6fce45f48c55226b67baff9c3b8a75c7c759b941bf0016622d0"></a>

## parameter_type property — custom_auth_types / dca934492cba / 5

Type: `"string"`. Computed.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

Upstream description:

Enumeration for authentication parameter types.

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

<a id="canonical-aa5469ca9b1dc9ab55b499391729f206dd57878bdf86a0a8b1ea5820c929f49b"></a>

## Next pages — custom_auth_types / dca934492cba / 6

- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-044c4ec59e90132172e954c34d01a62d3cdee53648e0f7d46c570175d3f9089e"></a>

## user_defined_api_discovery_policy — user_defined_api_discovery_policy / 1e773e2559aa / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- user_defined_api_discovery_policy

<a id="canonical-a1daaf4759d2eecf4e4a26b30c2c8fdfb5b25eed1345f20c4d8330ede7d6f6bd"></a>

Type: `"single"`. Computed.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

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

<a id="canonical-3a83dba3be17835394c640a115dfc1e69713b5e48f44c2f0f418e9a867554a66"></a>

## Direct properties — user_defined_api_discovery_policy / 1e773e2559aa / 3

- [discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff): complete subsection reference.

- [exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7): complete subsection reference.

- [inclusive](data-sources--api_discovery--reference--group-001.md#canonical-3ce833875c4473db0ce0956084f54f6b79efdfb1debe8621cbe9001b71e7b948): complete subsection reference.

<a id="canonical-240b6440fab9436e06cae325f1b8fc69f883838dd69745f81053a6dd6e5f652b"></a>

## Next pages — user_defined_api_discovery_policy / 1e773e2559aa / 4

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7)
- [user_defined_api_discovery_policy.inclusive](data-sources--api_discovery--reference--group-001.md#canonical-3ce833875c4473db0ce0956084f54f6b79efdfb1debe8621cbe9001b71e7b948)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a283e6337c692f933f57df25e741aaf0c06e1809415bcd6cbf631061b660ad49"></a>

## user_defined_api_discovery_policy.discovery_rules — user_defined_api_discovery_policy.discovery_rules / 11cc3635d71b / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- user_defined_api_discovery_policy.discovery_rules

<a id="canonical-cebe1ecc3966ac40908956a54e7e7683523a492b60d49fad9e1ab72a3c416f29"></a>

Type: `"list"`. Computed.

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

<a id="canonical-cfeb6e02043e588c8a7a2719031ffd5f56ac1c8db322e5d56746c50123f75488"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules / 11cc3635d71b / 3

- [labels](data-sources--api_discovery--reference--group-001.md#canonical-5d2b192bc63579d3878f6e402a554262d10d42b4a758d0d727569791b88c2354): complete subsection reference.

- [metadata](data-sources--api_discovery--reference--group-001.md#canonical-dab6594c5eb0f02a1d0e7b6bcc5d1664219716c03d655f6727f0140ef4bcbc02): complete subsection reference.

- [rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6): complete subsection reference.

<a id="canonical-810b4cbebb61be9e84b2350f865313bcee0cad542b358c5cf239e6ec7cf64d2c"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules / 11cc3635d71b / 4

- [user_defined_api_discovery_policy.discovery_rules.labels](data-sources--api_discovery--reference--group-001.md#canonical-5d2b192bc63579d3878f6e402a554262d10d42b4a758d0d727569791b88c2354)
- [user_defined_api_discovery_policy.discovery_rules.metadata](data-sources--api_discovery--reference--group-001.md#canonical-dab6594c5eb0f02a1d0e7b6bcc5d1664219716c03d655f6727f0140ef4bcbc02)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-5d2b192bc63579d3878f6e402a554262d10d42b4a758d0d727569791b88c2354"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f59e346581d12aa04c84fc0ad471d4bbe4d250022426e75b2be62b2e110502c"></a>

## user_defined_api_discovery_policy.discovery_rules.labels — user_defined_api_discovery_policy.discovery_rules.labels / 48f8a9383b14 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- user_defined_api_discovery_policy.discovery_rules.labels

<a id="canonical-fe6d05b54da95948a3306780db1a605a27d56bd603590b204a3fe9281711ee12"></a>

Type: `"single"`. Computed.

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

<a id="canonical-62fcbe17e35c9a1e9c604e02a0bc8e484a8c0051c2b45aeaeafdde81e5162462"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.labels / 48f8a9383b14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64d944fb3bc0ba1d03c711317a753910eda475fd028623e82d0e70f6d4851699"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.labels / 48f8a9383b14 / 4

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-dab6594c5eb0f02a1d0e7b6bcc5d1664219716c03d655f6727f0140ef4bcbc02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-871e482134a5887d327ee77918120aaa21eb41bcc3f445fc6899855a7ca1232d"></a>

## user_defined_api_discovery_policy.discovery_rules.metadata — user_defined_api_discovery_policy.discovery_rules.metadata / 57343a3db9df / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- user_defined_api_discovery_policy.discovery_rules.metadata

<a id="canonical-49cdf0310f6cb1ac4a588a3cc319f552a444ce296b42d3cc9aa89b92def3ab16"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-100cdaeea61b00f81d7dfbf9e3ec14b02f374c6231563e17518448e08c1a354c"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.metadata / 57343a3db9df / 3

<a id="canonical-cca9603d1b8b2d9babd9c524b20993a7a10bf402ef22bd370d9bd85bf614edb1"></a>

<a id="canonical-16db01d5d03b2434032eec0b09461e8c1b1f35b14ad994d2fe31aaefcf81bd2a"></a>

## description_spec property — user_defined_api_discovery_policy.discovery_rules.metadata / 57343a3db9df / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-57b65dd893f1a040d4c77aa552f7987d7fee5c314db42e9a1db5c47594c48606"></a>

<a id="canonical-b580b582a015f766cce861497e1c470e600e81cec490303f00ae6089105a423b"></a>

## name property — user_defined_api_discovery_policy.discovery_rules.metadata / 57343a3db9df / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3a9d0598773fb8082147f3823f626d47578426cc211559291cb2ec669cd2617d"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.metadata / 57343a3db9df / 6

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4ea615a39972ca3347239d2bc7c1e5636cc99b84c30f9864069db7442c0a962"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties — user_defined_api_discovery_policy.discovery_rules.rule_properties / 9111e53cdf74 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="canonical-b9ddf479f034861952d33b90b12063bce59c036342b6043005a3a1bc520b78c2"></a>

Type: `"single"`. Computed.

Determines whether matching endpoints are included in API Discovery or excluded.

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

<a id="canonical-14ddba266abd0cad6ba9068ab115aa61cfd63f626aa624493e5fe7665af88bea"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties / 9111e53cdf74 / 3

- [exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9): complete subsection reference.

- [http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-2a22f2cfcbf71953e9a3b999a299b96b8b8068152460d6788cd09f55af0f5059): complete subsection reference.

- [inclusion](data-sources--api_discovery--reference--group-001.md#canonical-a20b9742fc4b0e67494e59440d65f8d480e29ad40133d592f83922e45f68317d): complete subsection reference.

<a id="canonical-f930dbf319779eb7c6e5b6481c8b0b7705ab476c209f893d768904bdb1121392"></a>

<a id="canonical-0fc06d489333a7b0cf0ffee41c1b7d5f6421e6bafb917b8b503f0fdd005f2fa8"></a>

## pattern property — user_defined_api_discovery_policy.discovery_rules.rule_properties / 9111e53cdf74 / 4

Type: `"string"`. Computed.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Upstream description:

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

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

<a id="canonical-c12fc9117c15e28cf2f13b422bc11732f652c8e7ee7ab101cf6660e89ec27bc4"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties / 9111e53cdf74 / 5

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-2a22f2cfcbf71953e9a3b999a299b96b8b8068152460d6788cd09f55af0f5059)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](data-sources--api_discovery--reference--group-001.md#canonical-a20b9742fc4b0e67494e59440d65f8d480e29ad40133d592f83922e45f68317d)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abdd475932d47b1c6cb20ff704d0175215cb8b259a2e8a9be8fcb11a20f63ade"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 14ea1fe3c5be / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

<a id="canonical-55db5bd2a8c67182207db72968e901ffb54762358a509aa253c04e2da839098b"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

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

<a id="canonical-248f5f9603ae4d3bc73eee97ca8fa0e96667304e06a8d28626d84f9b60b4473a"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 14ea1fe3c5be / 3

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-65339affc5ecbb4b985eb93dc00f7a8f10154e2a5b3c77b7728ba8e6e83e308e): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-ab8a911448c4f81cf5b9af22e138abacae15db3667793eaef45825979c99e666): complete subsection reference.

<a id="canonical-35e65110748e3b4e907c8cfeda0accd8771ba80c509bc748e0a26757de358e8a"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion / 14ea1fe3c5be / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](data-sources--api_discovery--reference--group-001.md#canonical-65339affc5ecbb4b985eb93dc00f7a8f10154e2a5b3c77b7728ba8e6e83e308e)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](data-sources--api_discovery--reference--group-001.md#canonical-ab8a911448c4f81cf5b9af22e138abacae15db3667793eaef45825979c99e666)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-65339affc5ecbb4b985eb93dc00f7a8f10154e2a5b3c77b7728ba8e6e83e308e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5e2e94090c345daa223209af0058bc46b22d6710545c845861b80bbeddf4d99"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / cbf60f7f6ab5 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive

<a id="canonical-49048d9014d4b917ffa6c73094020827fa97441752ccac511096733bbddae843"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-81dbd1c79c8a1c51d2bcdea874655e2cb41d6039358d2ebd6b22ee44f909dd1f"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / cbf60f7f6ab5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07f9163ae2b7fa150c97beccb0bb67d481e6050deca03ac2ab03e95ab65291f5"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.arch / cbf60f7f6ab5 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-ab8a911448c4f81cf5b9af22e138abacae15db3667793eaef45825979c99e666"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d65b1741c6c4a443d001dbc5a84e2c698df8fdd556ebd82b3165f4f996fcfcf"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 9f28455d3c0e / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore

<a id="canonical-d2c44bef0cac43c056d7b3f01d35616b82d77517b16dbc73e154ba8d1fe52f7f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2799e9c2c373de16630c32d61e7fbc07b08e888b5120c8d0c5d3097e1885c130"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 9f28455d3c0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a58e1b2ca14d65dd10fcc5bebcd15baaed88f304f27f184e108470eada75f459"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.igno / 9f28455d3c0e / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-e9d3a8bf69b199db56deef60be86878bfe6df60e357b89f21a10f205151057a9)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-2a22f2cfcbf71953e9a3b999a299b96b8b8068152460d6788cd09f55af0f5059"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-736c72350ce1163145218bdc49ab7379fd227aa52f6ddb56fd994e811f188fca"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

<a id="canonical-f75e87e22e9fdd21812092a41f0d0de97e80e97431a45ac190cee2d14b928d2a"></a>

Type: `"single"`. Computed.

Configuration parameter for http header criteria.

Upstream description:

Criteria for matching HTTP headers.

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

<a id="canonical-361dd0aa09ac0565aa8d6877cf078176c039643b8ebed641c0616fa07444bafb"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 3

<a id="canonical-72b5f6a67b6083f9a61629c03532a0c09f540634417dc60dced10093b368838a"></a>

<a id="canonical-1fbaa4b7ebc1ff5d90847580924d2b43a83440c3ce14981fbd7310c8f1e1275e"></a>

## field_name property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 4

Type: `"string"`. Computed.

HTTP Header Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-916c54b2798a0453d87b6e380224e233d12f5efd4205e7109c8c6049d469715e"></a>

<a id="canonical-bae47555c23952dc13f0e843d4e82e4d518abaee568bb86ed3e74fabbfc7c2cf"></a>

## location property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 5

Type: `"string"`. Computed.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

Upstream description:

Specifies whether the rule criteria should be evaluated against request or response

Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back
to the client.

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

<a id="canonical-f38942eb939f94c5b718aa3f785f464e120aab68c49d8649860fc3f80cdef1aa"></a>

<a id="canonical-88f1012a1870c95ab19e1b46e94321949c2f88f638be4da09a056529099bfb58"></a>

## match_type property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 6

Type: `"string"`. Computed.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

Upstream description:

Specifies how the value should be matched.

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

<a id="canonical-a529e017e2a61670e2eeda41c322676149ec45c71369923598056a5f2a72dc47"></a>

<a id="canonical-6ce76495c2d96ef93ba45b1b928a7476f9e512a5e995280e06cd75c8a9488af9"></a>

## value property — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 7

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

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

<a id="canonical-6e5c56002742f0a9650563d45a3b7d26596862de9188915c64a8d19937b7396b"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_cr / 5830e7c28c1f / 8

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-a20b9742fc4b0e67494e59440d65f8d480e29ad40133d592f83922e45f68317d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-045e44ce0708a7665c1d531bc9a2ceaecab95202c9459ed6bd9eb1e5183816ca"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 1ddb7c5e95fc / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-783e29ef30d911ce42db5c00200d0b8b9a5012a39522f11d16d29ee491fc3aff)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion

<a id="canonical-fcbacfbdb299034beb04e4ad819fb69effa6688115de6a5df375dd59c3e05c4a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-26dedb2b37bcf17efdccbd9c9172616214f8258d9658c1401720855adb3df005"></a>

## Direct properties — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 1ddb7c5e95fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-694d1c50159326322ec143892ef2b51dfa3dc7a549a7e423a9898f0084758c57"></a>

## Next pages — user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion / 1ddb7c5e95fc / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-fbcc5acef872b8584b26ef657449999f5c1a520cc8601c3baeb0a4f91d6324b6)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f2f06595c7b087b16415c807e098451ffb0c0f013217c15293089bf33259c40"></a>

## user_defined_api_discovery_policy.exclusive — user_defined_api_discovery_policy.exclusive / e6d0cd0fdb99 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- user_defined_api_discovery_policy.exclusive

<a id="canonical-b63035187bf70e0ec907d80e1555b4074fa922161fea9799b6aea43731bc5194"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

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

<a id="canonical-2265891cdf67ec43c2069f7a50601421eeec36eda2cdf51d43fb51237c0c2aff"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive / e6d0cd0fdb99 / 3

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-4d3b9087dd1fd2a0e60b0af2a53a8669e7a3cdf8c56e331f9824f79ce7f7ea17): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-20cd638f789ce8b12a363dc62cefe987416c07f599bd566db6b9db91a9eecda9): complete subsection reference.

<a id="canonical-fceda958aecf4013b049c5aaac323e5e98d21a37ded49cb4c24f21511415e3d4"></a>

## Next pages — user_defined_api_discovery_policy.exclusive / e6d0cd0fdb99 / 4

- [user_defined_api_discovery_policy.exclusive.archive](data-sources--api_discovery--reference--group-001.md#canonical-4d3b9087dd1fd2a0e60b0af2a53a8669e7a3cdf8c56e331f9824f79ce7f7ea17)
- [user_defined_api_discovery_policy.exclusive.ignore](data-sources--api_discovery--reference--group-001.md#canonical-20cd638f789ce8b12a363dc62cefe987416c07f599bd566db6b9db91a9eecda9)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-4d3b9087dd1fd2a0e60b0af2a53a8669e7a3cdf8c56e331f9824f79ce7f7ea17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5a380c99b4b59b4dc3e090b32ebef1566fc2b70926f23d464f3ef3cceb80c3e"></a>

## user_defined_api_discovery_policy.exclusive.archive — user_defined_api_discovery_policy.exclusive.archive / e33d59e902dd / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7)
- user_defined_api_discovery_policy.exclusive.archive

<a id="canonical-d95f3acaf0e55473d8c68e9b91687ea02418b33899184ab1e183581daa39ede2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-98a54b6ebd7ec1f3cac017ba51b142a07ed1a424e097dc1e0bf8edc34d3cb290"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive.archive / e33d59e902dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fcfcf6afc4db5f4f9719fb9103b9d3c037d71668f013cd604263fed6219e26d"></a>

## Next pages — user_defined_api_discovery_policy.exclusive.archive / e33d59e902dd / 4

- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-20cd638f789ce8b12a363dc62cefe987416c07f599bd566db6b9db91a9eecda9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60e8629bb6ddeb76708bc994b3e93f1b24ac9a5d9de7116aae1ab3ae39aaf31c"></a>

## user_defined_api_discovery_policy.exclusive.ignore — user_defined_api_discovery_policy.exclusive.ignore / e57a5ae4f827 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7)
- user_defined_api_discovery_policy.exclusive.ignore

<a id="canonical-bf15e05cbf45a7fc74ef3d6a5cb7c28d5cec97112f4fbbda9d5a93b4bc371c99"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8d0c48ddb349170e7ce8266d0f9bc2ef449e90871ff1335c0e521e7b5262839f"></a>

## Direct properties — user_defined_api_discovery_policy.exclusive.ignore / e57a5ae4f827 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-095cef639294cf14a704e85001a64c1ebff395bfbd8227bca2733b397fd4f181"></a>

## Next pages — user_defined_api_discovery_policy.exclusive.ignore / e57a5ae4f827 / 4

- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-b0e8660221ba99a49b969245a004a469bdf69216047532d1c30b7df4cc486cb7)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-3ce833875c4473db0ce0956084f54f6b79efdfb1debe8621cbe9001b71e7b948"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1a21f20cce8ad5f2e06456f335d56e244688f9d5d5694b1cc72207856d59e8a"></a>

## user_defined_api_discovery_policy.inclusive — user_defined_api_discovery_policy.inclusive / 28a8a1501d34 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- user_defined_api_discovery_policy.inclusive

<a id="canonical-78a9c420c6cd63e788e70962e0b1772a6c4697326d148fbbdd69aaf0001455f7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5c0eb4322c2b9548129f56835484669482a665970e56f590f23bc4994cd89965"></a>

## Direct properties — user_defined_api_discovery_policy.inclusive / 28a8a1501d34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72f611d179021c6684972bfbcef30ef5ffabf3201525b35c9909ce4253270fe1"></a>

## Next pages — user_defined_api_discovery_policy.inclusive / 28a8a1501d34 / 4

- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-3abef7bd02e08d9a5e69bca4ccced59ccd3bca8006fcb46fc445bea42cae609f)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
