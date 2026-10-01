---
page_title: "xcsh_ike1 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 reference."
---

# xcsh_ike1 reference

<a id="canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3f2b964c88ac8755a716583cdc245e992a9c25518945b372f75f1d56e0eb846"></a>

## Property reference — Property reference / 674a40595e0a / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- Property reference

<a id="canonical-ada52f55ae8675ef403e9bca6f4a2d2293b4d5dd7202c1c82da190b5a6d84bb5"></a>

## Direct properties — Property reference / 674a40595e0a / 3

<a id="canonical-e2fe2aaa8510599367e799fb378a2084ffe4377add66966b9a8f73c698124733"></a>

<a id="canonical-0d9fbe4304771ac2964c6c2a03344675da79bb42f210ccca95197db57bce4572"></a>

## annotations property — Property reference / 674a40595e0a / 4

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

<a id="canonical-e41b3a0fb3e3c395629c4158cf444e491049a256415400a4374f0f9aa8733820"></a>

<a id="canonical-b6a0c72a0df6ec67070ab19cebd17f940fde74f3a7fb77eeb9396c0f53c975ed"></a>

## description property — Property reference / 674a40595e0a / 5

Type: `"string"`. Computed.

Description of the Ike1.

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

<a id="canonical-0e7a84e83d210f4bd655015c70038a8fff8e08c919ed67a5b31138e29e7ad913"></a>

<a id="canonical-3ff1cf9935c058b56a475754bd6b13204897a0cef92e43358dd70f4183d74e19"></a>

## id property — Property reference / 674a40595e0a / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-1837670d5eafda209f40c41913ab8b39bc277eb15f5089ad642caacc3c1553a0): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-6ecb1ddbeba9ac926cf97dda50f3e8622846fb69344104611aa728b8960559e5): complete subsection reference.

<a id="canonical-061efd15243557774b8986a8a1322aa0926bed65a112fbe497cdeecdb4e64889"></a>

<a id="canonical-8c16ce667af30fe9766b3c9e1861fc4c7a4d931d91d91c98dac8cfdb14aced62"></a>

## labels property — Property reference / 674a40595e0a / 7

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

<a id="canonical-4f0b16f4b3c528f309646e99c2b3a542458514449fc1b1bed7c636b5c9062dcd"></a>

<a id="canonical-28b55bedc9e058daa0f5bd4b5cc8a4b2c66713f711750f3e22f24a16d9cffabd"></a>

## name property — Property reference / 674a40595e0a / 8

Type: `"string"`. Required.

Name of the Ike1.

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

<a id="canonical-513b8b812c5908a5356300b4ce5054e605e743d0c27ba7f5cc1034e530134559"></a>

<a id="canonical-9a18264c736223863cd0f479e92dc60964a6d44b63ddb7ecc78489ca87f61e3c"></a>

## namespace property — Property reference / 674a40595e0a / 9

Type: `"string"`. Required.

Namespace where the Ike1 exists.

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

- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-cfdadc4976f94df40a0f88f8f55f8eacb3d022cd33b7fc13837ae4f6dcb390f6): complete subsection reference.

- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-9f5780cb1a126497a40e27d626a991556e465226d8305aba42bb1d69a9c4058a): complete subsection reference.

- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-121d040c70e58f5da1c83ef962d9b2078ce3d433d7f7c1d8c40ca2e74e036094): complete subsection reference.

- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-02566e77a149f9bf9f5020567c55112e03ce4c8eaf5b0048292559d2a549f8a4): complete subsection reference.

<a id="canonical-1cd83b13bfae97378e7a1fe9917c1c08f95b444ad091c64447be10bcab69af8a"></a>

## All schema paths — Property reference / 674a40595e0a / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike1--reference--group-001.md#canonical-e2fe2aaa8510599367e799fb378a2084ffe4377add66966b9a8f73c698124733) |
| `description` | [description](data-sources--ike1--reference--group-001.md#canonical-e41b3a0fb3e3c395629c4158cf444e491049a256415400a4374f0f9aa8733820) |
| `id` | [id](data-sources--ike1--reference--group-001.md#canonical-0e7a84e83d210f4bd655015c70038a8fff8e08c919ed67a5b31138e29e7ad913) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-7232a2b64c0376754e80bcbcf319daf78b12bf86c451a60a3e797c5da4b49208) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike1--reference--group-001.md#canonical-81586cf1d496ef5686a3c498ff7c5407a7730e2723a954324ba6981541d5fa49) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-7c9626b25dba2f7e9ea8cdd5b572cf16f2353cab1e571602190264e364cbd997) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike1--reference--group-001.md#canonical-64f960acdefd00d6f25d2b3dc37c111e94399c7a1fabbe92dfa74c7c64e1672f) |
| `labels` | [labels](data-sources--ike1--reference--group-001.md#canonical-061efd15243557774b8986a8a1322aa0926bed65a112fbe497cdeecdb4e64889) |
| `name` | [name](data-sources--ike1--reference--group-001.md#canonical-4f0b16f4b3c528f309646e99c2b3a542458514449fc1b1bed7c636b5c9062dcd) |
| `namespace` | [namespace](data-sources--ike1--reference--group-001.md#canonical-513b8b812c5908a5356300b4ce5054e605e743d0c27ba7f5cc1034e530134559) |
| `reauth_disabled` | [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-faaa6eb445887fbd3606b355cb167798fcade087319f964ea5b3e7a2e29168f8) |
| `reauth_timeout_days` | [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-a976b6591c1d07313ac11b437f73bb50db1c475e52d610ff78d37ae046a5b530) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](data-sources--ike1--reference--group-001.md#canonical-90480891043ea9c9696934433bdea7cf7c4d0b410380bd954d466e98d1210737) |
| `reauth_timeout_hours` | [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-d302cea8a0faa46fb9a23bb6e6f248ba89a919a6278f015ae4beb1ee1cac5e2c) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](data-sources--ike1--reference--group-001.md#canonical-72de5e63ac2d86be63bfc349bf0c6062e330de789baff41efddd77d93c815165) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0add465ad63432526486844528211c65e36673c2e70bdbc13b652fc71f961ad5) |

<a id="canonical-a814422f81d20291c2ab43b3714089bed2e00c479411429d6bea5f1a780fe80d"></a>

## Next pages — Property reference / 674a40595e0a / 11

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-1837670d5eafda209f40c41913ab8b39bc277eb15f5089ad642caacc3c1553a0)
- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-6ecb1ddbeba9ac926cf97dda50f3e8622846fb69344104611aa728b8960559e5)
- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-cfdadc4976f94df40a0f88f8f55f8eacb3d022cd33b7fc13837ae4f6dcb390f6)
- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-9f5780cb1a126497a40e27d626a991556e465226d8305aba42bb1d69a9c4058a)
- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-121d040c70e58f5da1c83ef962d9b2078ce3d433d7f7c1d8c40ca2e74e036094)
- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-02566e77a149f9bf9f5020567c55112e03ce4c8eaf5b0048292559d2a549f8a4)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-1837670d5eafda209f40c41913ab8b39bc277eb15f5089ad642caacc3c1553a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88c3dc5c5085735b4c9f95f6295ae53173d5f5197e83bd30818a1231429bf638"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 692f47aaeda6 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- ike_keylifetime_hours

<a id="canonical-7232a2b64c0376754e80bcbcf319daf78b12bf86c451a60a3e797c5da4b49208"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-7232a2b64c0376754e80bcbcf319daf78b12bf86c451a60a3e797c5da4b49208)
- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-7c9626b25dba2f7e9ea8cdd5b572cf16f2353cab1e571602190264e364cbd997)
- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0add465ad63432526486844528211c65e36673c2e70bdbc13b652fc71f961ad5)

Select alternatives according to the provider validators above.

<a id="canonical-822b13cb99bd586443bbef3b78ac8695364b5496c0811d42aa7016a976018639"></a>

## Direct properties — ike_keylifetime_hours / 692f47aaeda6 / 3

<a id="canonical-81586cf1d496ef5686a3c498ff7c5407a7730e2723a954324ba6981541d5fa49"></a>

<a id="canonical-000629fd0de90566b4d7c0e9f820828f9059ddfc31e9fa077a03fe5e64c573b6"></a>

## duration property — ike_keylifetime_hours / 692f47aaeda6 / 4

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
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-e6f09b5f24c4cf92d605ac351ba074f15e5f33e26d6eeef57a9156924912a4fc"></a>

## Next pages — ike_keylifetime_hours / 692f47aaeda6 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-6ecb1ddbeba9ac926cf97dda50f3e8622846fb69344104611aa728b8960559e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74bf609cc2a2c98ace9dee45335d0e9c94b83bba594b9ffc4dfe7d9d081e9f70"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 60587c454c11 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- ike_keylifetime_minutes

<a id="canonical-7c9626b25dba2f7e9ea8cdd5b572cf16f2353cab1e571602190264e364cbd997"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-9251c89358a81807604d6ae3b5a84e1175d12691856ff409432139eec92c3f20"></a>

## Direct properties — ike_keylifetime_minutes / 60587c454c11 / 3

<a id="canonical-64f960acdefd00d6f25d2b3dc37c111e94399c7a1fabbe92dfa74c7c64e1672f"></a>

<a id="canonical-65ba2e4981e87ef9fee05bf974433b7e95b5037edf1809d4aa7c65346fd1597b"></a>

## duration property — ike_keylifetime_minutes / 60587c454c11 / 4

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
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-8965611653ad62208201802950000f814109092768a7a7d7f3b46022f0ca1029"></a>

## Next pages — ike_keylifetime_minutes / 60587c454c11 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-cfdadc4976f94df40a0f88f8f55f8eacb3d022cd33b7fc13837ae4f6dcb390f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8833c462a8c842ee22e221d99f5a726e93b27836139f37da2fe3460c3c5fb88"></a>

## reauth_disabled — reauth_disabled / 00ddb94c9486 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- reauth_disabled

<a id="canonical-faaa6eb445887fbd3606b355cb167798fcade087319f964ea5b3e7a2e29168f8"></a>

Type: `["object", {}]`. Computed.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

OneOf alternatives in this subsection:

- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-faaa6eb445887fbd3606b355cb167798fcade087319f964ea5b3e7a2e29168f8)
- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-a976b6591c1d07313ac11b437f73bb50db1c475e52d610ff78d37ae046a5b530)
- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-d302cea8a0faa46fb9a23bb6e6f248ba89a919a6278f015ae4beb1ee1cac5e2c)

Select alternatives according to the provider validators above.

<a id="canonical-341bce7b74fb83a3bf57bab2fa8161d04e795cef0f0fb7923bbde51bec9f176c"></a>

## Direct properties — reauth_disabled / 00ddb94c9486 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f31d945fb9d929e6e33b50d1415c63f463eb84b0f0fde94cd0a9dd0d715bea14"></a>

## Next pages — reauth_disabled / 00ddb94c9486 / 4

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-9f5780cb1a126497a40e27d626a991556e465226d8305aba42bb1d69a9c4058a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-836377871d73e4b78b3f78d73133cb3786677092c946689b12ac53186caca7dc"></a>

## reauth_timeout_days — reauth_timeout_days / cd81fccb9fd2 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- reauth_timeout_days

<a id="canonical-a976b6591c1d07313ac11b437f73bb50db1c475e52d610ff78d37ae046a5b530"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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

<a id="canonical-62ddd17be6f59074500129989b40aa864fa518af0ecfb2d873f1d70937c4446a"></a>

## Direct properties — reauth_timeout_days / cd81fccb9fd2 / 3

<a id="canonical-90480891043ea9c9696934433bdea7cf7c4d0b410380bd954d466e98d1210737"></a>

<a id="canonical-c644230e0ec878d800f6b77c450b5c30bbd0cbe56cfdeba1300d60525d464409"></a>

## duration property — reauth_timeout_days / cd81fccb9fd2 / 4

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
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-6d31df86b060b0be64e88b2179b732d0e5ffaba011a81224c0b1a1861cad0c82"></a>

## Next pages — reauth_timeout_days / cd81fccb9fd2 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-121d040c70e58f5da1c83ef962d9b2078ce3d433d7f7c1d8c40ca2e74e036094"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abe278a187ae6c71e628491daee4d0d15b6cf9d565229a03e6247a7aff2d3885"></a>

## reauth_timeout_hours — reauth_timeout_hours / f366dd688cc4 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- reauth_timeout_hours

<a id="canonical-d302cea8a0faa46fb9a23bb6e6f248ba89a919a6278f015ae4beb1ee1cac5e2c"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

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

<a id="canonical-6563acaeed3bf8311916e156b6470da3f4e4edf5a5b920bf18aac2a49d904499"></a>

## Direct properties — reauth_timeout_hours / f366dd688cc4 / 3

<a id="canonical-72de5e63ac2d86be63bfc349bf0c6062e330de789baff41efddd77d93c815165"></a>

<a id="canonical-7557070c1867e75f411fce8f2a0d8fb3deea1b13d3657bbf596a1f77e44331a2"></a>

## duration property — reauth_timeout_hours / f366dd688cc4 / 4

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
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-4c2fcea53453f3c80ae6aebbc1cbce99dded31b8947556fb42d4a4c9e6e1838b"></a>

## Next pages — reauth_timeout_hours / f366dd688cc4 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)

<a id="canonical-02566e77a149f9bf9f5020567c55112e03ce4c8eaf5b0048292559d2a549f8a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a82a4cecbb2a46c69d960f02c7e8cba9ea4ce9896296298389001cd228dd618"></a>

## use_default_keylifetime — use_default_keylifetime / f4bbad4e0022 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- use_default_keylifetime

<a id="canonical-0add465ad63432526486844528211c65e36673c2e70bdbc13b652fc71f961ad5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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

<a id="canonical-e3d15417d58fa15535cdf4d908468f07b85080a5176abb14398cc1c6a7a25857"></a>

## Direct properties — use_default_keylifetime / f4bbad4e0022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4d12f34860c6cc4cb8016dabeb706e231a19cd94e14e30aeb8ddce2f37039d0"></a>

## Next pages — use_default_keylifetime / f4bbad4e0022 / 4

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [xcsh_ike1](../data-sources/ike1.md#canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7)
