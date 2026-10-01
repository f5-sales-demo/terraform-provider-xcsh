---
page_title: "xcsh_rate_limiter reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter reference."
---

# xcsh_rate_limiter reference

<a id="canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3a30b17604f76a36d0730d5bb413629648b6a20bf45172bb707c6dee735c284"></a>

## Property reference — Property reference / ae2990700e76 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- Property reference

<a id="canonical-4cc7234ffeab2bcc992e117f174ca39a1c8f591ef8bfd7b7df9c3f226d4bae75"></a>

## Direct properties — Property reference / ae2990700e76 / 3

<a id="canonical-63c93c37dabc25b5816745cf5de4b44fd843558ec4b06c2833c94456f9844e1e"></a>

<a id="canonical-dca67297fac905b16e45ad22cb6772d1e085516c15bde0135caa44d84335b804"></a>

## annotations property — Property reference / ae2990700e76 / 4

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

<a id="canonical-88f48a0b1f9b77a0d8868b8e25884e047eb02c009d4ec0214874f6f796ad6669"></a>

<a id="canonical-b958b3ba9c8600c1680e51f9dc990e6481e1e0e3b07fbe0e47aaba60a9d86866"></a>

## description property — Property reference / ae2990700e76 / 5

Type: `"string"`. Computed.

Description of the RateLimiter.

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

<a id="canonical-7e280590d8d4c7079858da6018ddb20fa8539e4dba085edde5e7befa2dea17c4"></a>

<a id="canonical-93e173762942badacac3dee3d290718619661a4eea117f460ef74b34c2787f47"></a>

## id property — Property reference / ae2990700e76 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-013c62effff9d15455dd6f1588d7f79f91221d29d1a539bdd976afafa158d729"></a>

<a id="canonical-ee122dd8d1d6893ca5b1229a0671ca11d6241c26d2885d3537c9d4a557751254"></a>

## labels property — Property reference / ae2990700e76 / 7

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

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5): complete subsection reference.

<a id="canonical-f03b930429b7a1917a43df3c770f5c7db276c4c783ba1ed3150c408fb742fe72"></a>

<a id="canonical-59f037c90ce165cf73da934ce72746b083b3c57c4a11d174de498cb4ec4a457e"></a>

## name property — Property reference / ae2990700e76 / 8

Type: `"string"`. Required.

Name of the RateLimiter.

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

<a id="canonical-278f0ae100eb4e7b104d2d87867676eb715c07b2e5bd481b3648b576f8552949"></a>

<a id="canonical-b8ef9be8056a3a869ebc850a6ff0f82e97bff936510550b15b2dd6867ef848b8"></a>

## namespace property — Property reference / ae2990700e76 / 9

Type: `"string"`. Required.

Namespace where the RateLimiter exists.

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

- [user_identification](data-sources--rate_limiter--reference--group-001.md#canonical-2e785194ce5d49cc5a46f78f660eb49b5e85ffaf21f4d12510336cde4ee4f8d2): complete subsection reference.

<a id="canonical-0b8548f3c4e2318e1d5faed9820766534b51b0b0af9a547404cf9a56a05cd1ff"></a>

## All schema paths — Property reference / ae2990700e76 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--rate_limiter--reference--group-001.md#canonical-63c93c37dabc25b5816745cf5de4b44fd843558ec4b06c2833c94456f9844e1e) |
| `description` | [description](data-sources--rate_limiter--reference--group-001.md#canonical-88f48a0b1f9b77a0d8868b8e25884e047eb02c009d4ec0214874f6f796ad6669) |
| `id` | [id](data-sources--rate_limiter--reference--group-001.md#canonical-7e280590d8d4c7079858da6018ddb20fa8539e4dba085edde5e7befa2dea17c4) |
| `labels` | [labels](data-sources--rate_limiter--reference--group-001.md#canonical-013c62effff9d15455dd6f1588d7f79f91221d29d1a539bdd976afafa158d729) |
| `limits` | [limits](data-sources--rate_limiter--reference--group-001.md#canonical-976faa9acc79533b69bd704d8a83d1081215b6095efe6dd84216d5aa497b5433) |
| `limits.action_block` | [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-6e5b0e4555b846f0ea7fc8f16333f3af3ed0400a39c695392345b9ad7268c030) |
| `limits.action_block.hours` | [limits.action_block.hours](data-sources--rate_limiter--reference--group-001.md#canonical-556a4cf36d3ecded750ff4e7352806563cda494e9b0573e3cdb749be3df43d0c) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](data-sources--rate_limiter--reference--group-001.md#canonical-e8adf03b6f90173a6a1108ee16fffd82a29c81d1979674fea49e2e3b7465ffe0) |
| `limits.action_block.minutes` | [limits.action_block.minutes](data-sources--rate_limiter--reference--group-001.md#canonical-ce25579aac2643d97a676ec6771343cf5db732a66157fc2646c45b90d6bedbca) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](data-sources--rate_limiter--reference--group-001.md#canonical-12097ce1f76d8115a2581bfcae3782db8acb0bc9ab226e4ef52cd88ea71eda77) |
| `limits.action_block.seconds` | [limits.action_block.seconds](data-sources--rate_limiter--reference--group-001.md#canonical-04f941b4e7989142e6f6ea93dda5efb3b1ad6f10e8921596d0f80a8b2e989602) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](data-sources--rate_limiter--reference--group-001.md#canonical-3d2c73542182574913d445bb99cbb0f75b6e9dfd6c77eea90c08df0d668bb5c1) |
| `limits.burst_multiplier` | [limits.burst_multiplier](data-sources--rate_limiter--reference--group-001.md#canonical-b75f03fc691192b13fdbba5398ef5a63e8890cd2a62d778e697d9a4d8e3fd466) |
| `limits.disabled` | [limits.disabled](data-sources--rate_limiter--reference--group-001.md#canonical-bbfb19a6e6b13b04858d5065e391875852f5483ed7e409c3798d09bcfd1d853e) |
| `limits.leaky_bucket` | [limits.leaky_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-e675c2273e716347b455ba44df272ca003f516a4f24018a2893fa537dd187b37) |
| `limits.period_multiplier` | [limits.period_multiplier](data-sources--rate_limiter--reference--group-001.md#canonical-7a6c80acf18325ee2cc2661107343b6f85e135d5f06968e1a8e4026acfb817a8) |
| `limits.token_bucket` | [limits.token_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-9edc36d3e31bc221d7f665c3c008d41573fb44b9e029c28141628ea60aaf1c30) |
| `limits.total_number` | [limits.total_number](data-sources--rate_limiter--reference--group-001.md#canonical-b6641007f85d155028ddde38f564c7b9f1e8fa389c95f0cdb52fecd4bc33d3c3) |
| `limits.unit` | [limits.unit](data-sources--rate_limiter--reference--group-001.md#canonical-84113d015e01f4b8993eac7057f1c6048f36ac673b47a1ebcc2ad37fb74a5ebe) |
| `name` | [name](data-sources--rate_limiter--reference--group-001.md#canonical-f03b930429b7a1917a43df3c770f5c7db276c4c783ba1ed3150c408fb742fe72) |
| `namespace` | [namespace](data-sources--rate_limiter--reference--group-001.md#canonical-278f0ae100eb4e7b104d2d87867676eb715c07b2e5bd481b3648b576f8552949) |
| `user_identification` | [user_identification](data-sources--rate_limiter--reference--group-001.md#canonical-afc60cd90fdc3ecf4579911779819ae9dce098377961ff6c07b2bd283822c623) |
| `user_identification.kind` | [user_identification.kind](data-sources--rate_limiter--reference--group-001.md#canonical-597ae2544ff0674f482cff9364bda474af9a3c9f3fc30d7eb66287166adda0d5) |
| `user_identification.name` | [user_identification.name](data-sources--rate_limiter--reference--group-001.md#canonical-eeb8e62407b0e306efe364af141871f300786a0569549f0e34b203b39daebae6) |
| `user_identification.namespace` | [user_identification.namespace](data-sources--rate_limiter--reference--group-001.md#canonical-b7281a5e398187898376f033ebbbe3c921bf73181d79ab37f5fca5b0ea7a711b) |
| `user_identification.tenant` | [user_identification.tenant](data-sources--rate_limiter--reference--group-001.md#canonical-1b822416d7ba7ff6eee3e5e8a7d0e9142cd0447f7764030c7c2bc29049df7289) |
| `user_identification.uid` | [user_identification.uid](data-sources--rate_limiter--reference--group-001.md#canonical-bdeedf407eb4145f5e995774315157dd27b4955c68159ebaf1053d89997011eb) |

<a id="canonical-a5d7d8379ae86038913753febf14325370b6e715629d46bb27bca62a6c595ad7"></a>

## Next pages — Property reference / ae2990700e76 / 11

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [user_identification](data-sources--rate_limiter--reference--group-001.md#canonical-2e785194ce5d49cc5a46f78f660eb49b5e85ffaf21f4d12510336cde4ee4f8d2)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e49b5bde1810567c2f53164d9197241d01166e444d086e4811beb6a98894dbf"></a>

## limits — limits / 8faaf2fea690 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- limits

<a id="canonical-976faa9acc79533b69bd704d8a83d1081215b6095efe6dd84216d5aa497b5433"></a>

Type: `"list"`. Computed.

List of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Upstream description:

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-5d4401366bc8cd1af3a0a7c9e562bd7c0e3ca2aef9ccd958f375ddad691e8f0f"></a>

## Direct properties — limits / 8faaf2fea690 / 3

- [action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46): complete subsection reference.

<a id="canonical-b75f03fc691192b13fdbba5398ef5a63e8890cd2a62d778e697d9a4d8e3fd466"></a>

<a id="canonical-d3fb2ee0003e58aec6c4c81e697a1528fb30331541240c57593dcc97bf39b41a"></a>

## burst_multiplier property — limits / 8faaf2fea690 / 4

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--rate_limiter--reference--group-001.md#canonical-6b2a38a26c654fdec27c34d32767fec384969c115acf264bf52ab70cc4cfdaf4): complete subsection reference.

- [leaky_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-16d0bdc86ee1735087dbca414f8554ea37d4e47b8cc9a94fdfa27690f7b0e813): complete subsection reference.

<a id="canonical-7a6c80acf18325ee2cc2661107343b6f85e135d5f06968e1a8e4026acfb817a8"></a>

<a id="canonical-02a15e92c4746b0d701de839b9c8d478ab6a12b79ddab3bd24aed27d050cbb77"></a>

## period_multiplier property — limits / 8faaf2fea690 / 5

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-a65ca8deb9f7f8dfd2a93809340d53b4f88178ff732814479a942850f23df879): complete subsection reference.

<a id="canonical-b6641007f85d155028ddde38f564c7b9f1e8fa389c95f0cdb52fecd4bc33d3c3"></a>

<a id="canonical-6f239ff2d4b7e87949037287a353f66f69be258c6ecae2eed62debf269e1096a"></a>

## total_number property — limits / 8faaf2fea690 / 6

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-84113d015e01f4b8993eac7057f1c6048f36ac673b47a1ebcc2ad37fb74a5ebe"></a>

<a id="canonical-7431708a3b07a6956d93f6f047157c1bf799d174fed3a219472239b79121cc5a"></a>

## unit property — limits / 8faaf2fea690 / 7

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2741176d06d9d97281f70f7b19beeddb25c3bf59868b378bdb0aa0b5fcd04907"></a>

## Next pages — limits / 8faaf2fea690 / 8

- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- [limits.disabled](data-sources--rate_limiter--reference--group-001.md#canonical-6b2a38a26c654fdec27c34d32767fec384969c115acf264bf52ab70cc4cfdaf4)
- [limits.leaky_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-16d0bdc86ee1735087dbca414f8554ea37d4e47b8cc9a94fdfa27690f7b0e813)
- [limits.token_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-a65ca8deb9f7f8dfd2a93809340d53b4f88178ff732814479a942850f23df879)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3dee6eb972f3ae01b6819ab0d051a3c32d76113cf3155e468ee7f3053e8de1b"></a>

## limits.action_block — limits.action_block / 3c2f4c3470e7 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- limits.action_block

<a id="canonical-6e5b0e4555b846f0ea7fc8f16333f3af3ed0400a39c695392345b9ad7268c030"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-1e8d8c50eb1506527b426f5a177ae98c8a9cbad97987f5bd0a77ec98ea5afb6f"></a>

## Direct properties — limits.action_block / 3c2f4c3470e7 / 3

- [hours](data-sources--rate_limiter--reference--group-001.md#canonical-5ed99723e85814a9c3961f1e4d9bd4445fe786c3944faa44943c98855f8e945f): complete subsection reference.

- [minutes](data-sources--rate_limiter--reference--group-001.md#canonical-49cbe5280d0e15246b5386ab496dbab3c573cbacf175bb4fc70c065cfe73007a): complete subsection reference.

- [seconds](data-sources--rate_limiter--reference--group-001.md#canonical-1a0714caf9933410bb0fa7ab445fd54a8b79c7cc9071f1c98b22c51fde58c19a): complete subsection reference.

<a id="canonical-e144f1fbb2871dd036e9b2c8c3335c5c175cb4631b15b99fe7a825f3466e6171"></a>

## Next pages — limits.action_block / 3c2f4c3470e7 / 4

- [limits.action_block.hours](data-sources--rate_limiter--reference--group-001.md#canonical-5ed99723e85814a9c3961f1e4d9bd4445fe786c3944faa44943c98855f8e945f)
- [limits.action_block.minutes](data-sources--rate_limiter--reference--group-001.md#canonical-49cbe5280d0e15246b5386ab496dbab3c573cbacf175bb4fc70c065cfe73007a)
- [limits.action_block.seconds](data-sources--rate_limiter--reference--group-001.md#canonical-1a0714caf9933410bb0fa7ab445fd54a8b79c7cc9071f1c98b22c51fde58c19a)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-5ed99723e85814a9c3961f1e4d9bd4445fe786c3944faa44943c98855f8e945f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad4f6c19315b1c2389eb1230c7426526d9672784f4d6c49952fa226f03247ae6"></a>

## limits.action_block.hours — limits.action_block.hours / 67de57790b4c / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- limits.action_block.hours

<a id="canonical-556a4cf36d3ecded750ff4e7352806563cda494e9b0573e3cdb749be3df43d0c"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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

<a id="canonical-72bb1b8ee636590a71527b7039855c81f04b5e3db5ef166a69ceb2e3bdb0a81b"></a>

## Direct properties — limits.action_block.hours / 67de57790b4c / 3

<a id="canonical-e8adf03b6f90173a6a1108ee16fffd82a29c81d1979674fea49e2e3b7465ffe0"></a>

<a id="canonical-9ad3e34595c6fa22733d7e00135aeaa4eb1a9521902aa384a33218c9cdd56e83"></a>

## duration property — limits.action_block.hours / 67de57790b4c / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-5831cd8108a719a04b49e97d4b2c337a561a37d6dbbe79ba1b8b7168e3dfa418"></a>

## Next pages — limits.action_block.hours / 67de57790b4c / 5

- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-49cbe5280d0e15246b5386ab496dbab3c573cbacf175bb4fc70c065cfe73007a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2704e05a577527c99e1cc6c36251c9e4bdfd794f3f612d48377e178978a10787"></a>

## limits.action_block.minutes — limits.action_block.minutes / 925b91d90499 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- limits.action_block.minutes

<a id="canonical-ce25579aac2643d97a676ec6771343cf5db732a66157fc2646c45b90d6bedbca"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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

<a id="canonical-922f416f37949f67182dee20817088216b93d1021a4ae463bfd505c4c17014c5"></a>

## Direct properties — limits.action_block.minutes / 925b91d90499 / 3

<a id="canonical-12097ce1f76d8115a2581bfcae3782db8acb0bc9ab226e4ef52cd88ea71eda77"></a>

<a id="canonical-03cb835f5bb86a6f88209ca7784ad6f85dbacb86086511c8772c8b6b19ac4b35"></a>

## duration property — limits.action_block.minutes / 925b91d90499 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-3dd25aa3d74c41657ef6f5a61ab1387abecec773c2a2df355ae7b3e488067c61"></a>

## Next pages — limits.action_block.minutes / 925b91d90499 / 5

- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-1a0714caf9933410bb0fa7ab445fd54a8b79c7cc9071f1c98b22c51fde58c19a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbd19603275ed05c22601756744911877b3b93961f09b763eeaa7fbcf389c1af"></a>

## limits.action_block.seconds — limits.action_block.seconds / 78cb490a02f7 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- limits.action_block.seconds

<a id="canonical-04f941b4e7989142e6f6ea93dda5efb3b1ad6f10e8921596d0f80a8b2e989602"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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

<a id="canonical-ed95558302292d56a33252049be1337e8c596e99f54f1263e1bfc51f26766c05"></a>

## Direct properties — limits.action_block.seconds / 78cb490a02f7 / 3

<a id="canonical-3d2c73542182574913d445bb99cbb0f75b6e9dfd6c77eea90c08df0d668bb5c1"></a>

<a id="canonical-06ae28e1c8add785af73ee6c8f494afc3fd768e731d0ea2d4b9e131809e10ac4"></a>

## duration property — limits.action_block.seconds / 78cb490a02f7 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-1003b23c681b80492783b4fb1fab9328611a5bf0f87d35462f23de41a447d4e1"></a>

## Next pages — limits.action_block.seconds / 78cb490a02f7 / 5

- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-e97292133612203b89a9479f07f51d44861d821c3be640c88ee8f024e5661a46)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-6b2a38a26c654fdec27c34d32767fec384969c115acf264bf52ab70cc4cfdaf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4041f5820527c36be426b0de2736a0d45385c53c50c95a6fc6fc0d858828368"></a>

## limits.disabled — limits.disabled / c94a0fc709c3 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- limits.disabled

<a id="canonical-bbfb19a6e6b13b04858d5065e391875852f5483ed7e409c3798d09bcfd1d853e"></a>

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

<a id="canonical-d423d19d12735c0fc0a155fa406612c4790bf3a26243277fa666f2c97e0e9819"></a>

## Direct properties — limits.disabled / c94a0fc709c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed62cde5b8059a2575d106ac5cae3549d5d210d612e06c7c951de2a1ebf8b8de"></a>

## Next pages — limits.disabled / c94a0fc709c3 / 4

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-16d0bdc86ee1735087dbca414f8554ea37d4e47b8cc9a94fdfa27690f7b0e813"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-067eb633ca9cec728dc02b633cd6037432cee5ee6c02a9acff5b78a1a8825608"></a>

## limits.leaky_bucket — limits.leaky_bucket / 41487350b813 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- limits.leaky_bucket

<a id="canonical-e675c2273e716347b455ba44df272ca003f516a4f24018a2893fa537dd187b37"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

<a id="canonical-8d5ea738a7e1059e2c8844c406444c339d8775523f810541074f46b2302cefd5"></a>

## Direct properties — limits.leaky_bucket / 41487350b813 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ea624bfdcbe9d18103db70753e907b48af15e63f398708115edece6945d60b8"></a>

## Next pages — limits.leaky_bucket / 41487350b813 / 4

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-a65ca8deb9f7f8dfd2a93809340d53b4f88178ff732814479a942850f23df879"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0e06dcb1f570ef4444e0ceb74bffa3c109eade2c7bac1439e6b8d04f460105"></a>

## limits.token_bucket — limits.token_bucket / 8adea24f229d / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- limits.token_bucket

<a id="canonical-9edc36d3e31bc221d7f665c3c008d41573fb44b9e029c28141628ea60aaf1c30"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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

<a id="canonical-bcaab567b1a1e71b36470dbf22f7429d1c05096cf795bea85608024438e7b656"></a>

## Direct properties — limits.token_bucket / 8adea24f229d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a2e41ec5dcb9ee35db53635ea391cc6264a8e8baecfd3fd87928b0d2093bb76"></a>

## Next pages — limits.token_bucket / 8adea24f229d / 4

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-64a4522a6c876e8a071ad77c6983b89135f42c799839dc9776dd35e28cb4b4c5)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-2e785194ce5d49cc5a46f78f660eb49b5e85ffaf21f4d12510336cde4ee4f8d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-278e8880dd1aec0d39ce057335f8c030d7c728df4e691a493f4808c94eb5cad8"></a>

## user_identification — user_identification / ccfd94696f5a / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- user_identification

<a id="canonical-afc60cd90fdc3ecf4579911779819ae9dce098377961ff6c07b2bd283822c623"></a>

Type: `"list"`. Computed.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-9afc80d2ab5afb3697a60cf66a64586b36a4043e5f5766e911b7446459ffbb0d"></a>

## Direct properties — user_identification / ccfd94696f5a / 3

<a id="canonical-597ae2544ff0674f482cff9364bda474af9a3c9f3fc30d7eb66287166adda0d5"></a>

<a id="canonical-d64d1471524c1f5785eedccdef497d31e8ccc534203cf9a416090bd74aaa6095"></a>

## kind property — user_identification / ccfd94696f5a / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-eeb8e62407b0e306efe364af141871f300786a0569549f0e34b203b39daebae6"></a>

<a id="canonical-8af68d941f71ae857f33924c4a63c13d3a30ed96c0e2d26d275fb4f4f0f3d568"></a>

## name property — user_identification / ccfd94696f5a / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-b7281a5e398187898376f033ebbbe3c921bf73181d79ab37f5fca5b0ea7a711b"></a>

<a id="canonical-115d055957b9edd9f712df523eaffa1aa9480fceaa0e986bb71741e5f35a110e"></a>

## namespace property — user_identification / ccfd94696f5a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1b822416d7ba7ff6eee3e5e8a7d0e9142cd0447f7764030c7c2bc29049df7289"></a>

<a id="canonical-2df2d423db75ff81acdf49060395ba98e32fbedc627ef8afec87528010d51025"></a>

## tenant property — user_identification / ccfd94696f5a / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-bdeedf407eb4145f5e995774315157dd27b4955c68159ebaf1053d89997011eb"></a>

<a id="canonical-d31e8806e31a4d44651424421d8fbbf9b64ec10db49f53736e743179208b6d45"></a>

## uid property — user_identification / ccfd94696f5a / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-9897870038ca69dc9d0028d96a3c253a08ee2b785cdec4872d31bfe301867940"></a>

## Next pages — user_identification / ccfd94696f5a / 9

- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
