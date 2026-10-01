---
page_title: "xcsh_nginx_service_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery reference."
---

# xcsh_nginx_service_discovery reference

<a id="canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7db2c38db22d4a3d928f98f2e50ecc1922011c81ba28145d6962c5ee54260b45"></a>

## Property reference — Property reference / b63700e5a8d7 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- Property reference

<a id="canonical-48e5e0babfb0baefb3c24f38ea7891823d0c3695b7695a57210792a38fc3a223"></a>

## Direct properties — Property reference / b63700e5a8d7 / 3

<a id="canonical-026aca834a5039bd2693680214e4d160a94c4c86379632dd6476fd9ae558391a"></a>

<a id="canonical-d116044a745d3bece9d0a062175c8f5ab6abddc9f428eb1769f16470fd0a5be9"></a>

## annotations property — Property reference / b63700e5a8d7 / 4

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

<a id="canonical-a575456414991cf0314347691a407b8a2e63d14649098a019297dbd54e4f4584"></a>

<a id="canonical-b8075cfbbcd77716d09505efb3db8aee1a006dfbe90f47c61d32478ecad53971"></a>

## description property — Property reference / b63700e5a8d7 / 5

Type: `"string"`. Computed.

Description of the NginxServiceDiscovery.

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

- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365): complete subsection reference.

<a id="canonical-566d300d0084480e42f2d0c311f536abef08662406ea7b8c1b40d0a754e0181f"></a>

<a id="canonical-b35f864fb4deb13b1e6136eba33fc5331715c9ed5db15603e94b42e38d242bf9"></a>

## id property — Property reference / b63700e5a8d7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1e494c3ada4d725e34b0941a714a4d791bd1765c64f8a8bfa6df394183dd7d1f"></a>

<a id="canonical-4098808fd87f9bf436f68032a705b591adcbe622a9290a7c2bb6f51477ffcd30"></a>

## labels property — Property reference / b63700e5a8d7 / 7

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

<a id="canonical-6315c0028c947c62fb51f86f8cd7e15bf7d4f368ee66d0488b3b14a9c3e316fa"></a>

<a id="canonical-7faa4dc4eb0c3f473130e9029c304a13e8ef978efc693e1916469604484e6c64"></a>

## name property — Property reference / b63700e5a8d7 / 8

Type: `"string"`. Required.

Name of the NginxServiceDiscovery.

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

<a id="canonical-ae995acaccde3e2a85e4d7302095eba73104912d0fa801949a09792b9bcec6ac"></a>

<a id="canonical-8ea47cb3a244590ca8c5ac600dc3098b2b27d0768a98514d41e75542ed5f2170"></a>

## namespace property — Property reference / b63700e5a8d7 / 9

Type: `"string"`. Required.

Namespace where the NginxServiceDiscovery exists.

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

- [server_block_filters](data-sources--nginx_service_discovery--reference--group-001.md#canonical-f15c85d7318332866b803af7bf7d9529e778ac3afb78c0c3a20d92f8579b85f2): complete subsection reference.

<a id="canonical-4cae7d0a2fca8e76f685ef1e9ec45c0b8bec6dae9f67cc5f01e4f46d4f558a34"></a>

## All schema paths — Property reference / b63700e5a8d7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_service_discovery--reference--group-001.md#canonical-026aca834a5039bd2693680214e4d160a94c4c86379632dd6476fd9ae558391a) |
| `description` | [description](data-sources--nginx_service_discovery--reference--group-001.md#canonical-a575456414991cf0314347691a407b8a2e63d14649098a019297dbd54e4f4584) |
| `discovery_target` | [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-df50e99ad9169135d31e6573baac1c344858bd5564560cf5d1147a873351998f) |
| `discovery_target.config_sync_group` | [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1665e9513d05689fe97ac44bae1dd5b45fce0fdcbf0b11234d2a84b60573a59c) |
| `discovery_target.config_sync_group.config_sync_group` | [discovery_target.config_sync_group.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-dda2a775c484343297312f41b7038bdb52b99e6f175927c45dc2d77e772c0fb9) |
| `discovery_target.config_sync_group.config_sync_group.kind` | [discovery_target.config_sync_group.config_sync_group.kind](data-sources--nginx_service_discovery--reference--group-001.md#canonical-fd557371eaddf52d858d1fae3363193e6736bef6d1c07c7de564fb1f62eee9f3) |
| `discovery_target.config_sync_group.config_sync_group.name` | [discovery_target.config_sync_group.config_sync_group.name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-b10d468a7fe446ec8d06547a29a8eb2d272c1b564c185ec9c1cf55e01a2fbc02) |
| `discovery_target.config_sync_group.config_sync_group.namespace` | [discovery_target.config_sync_group.config_sync_group.namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-907f9251705262beaafd86c98156a0c141d23dd5fb2d90fda4b45ae65d04f4c8) |
| `discovery_target.config_sync_group.config_sync_group.tenant` | [discovery_target.config_sync_group.config_sync_group.tenant](data-sources--nginx_service_discovery--reference--group-001.md#canonical-fd1ac596e90e03bee7a3ede32350309ed6d412394900f403648416b7844dbb1b) |
| `discovery_target.config_sync_group.config_sync_group.uid` | [discovery_target.config_sync_group.config_sync_group.uid](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2fc00dcdc349aca5092404744197164a1028851c96af75c3f5b72385e99025c4) |
| `discovery_target.nginx_instance` | [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2b2da8085b6a119323a02684d32c36cf90734d7f6ae0f10887cbb45554bf4cc1) |
| `discovery_target.nginx_instance.nginx_instance` | [discovery_target.nginx_instance.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-93a79813669bce397ead21cb9016a9046bec9b1480a1f74d73075d0e4fbc4394) |
| `discovery_target.nginx_instance.nginx_instance.kind` | [discovery_target.nginx_instance.nginx_instance.kind](data-sources--nginx_service_discovery--reference--group-001.md#canonical-22ddeb2a5ed830c13e865c0697b67e4d4650fcd68532e53ce80d5f1b18139070) |
| `discovery_target.nginx_instance.nginx_instance.name` | [discovery_target.nginx_instance.nginx_instance.name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1d76bbbecdf5c5ff0f23e0e345786e4d4b733f0fd1a1a4b7ec7dc086e5617d73) |
| `discovery_target.nginx_instance.nginx_instance.namespace` | [discovery_target.nginx_instance.nginx_instance.namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-fe725669b7b2a8875342334dba43ae6980a3327f0f14aa14652fba6fb409ff44) |
| `discovery_target.nginx_instance.nginx_instance.tenant` | [discovery_target.nginx_instance.nginx_instance.tenant](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1e34ce1e04f1f439910052f21d25a1f2112045d6190a7cf611058532333648be) |
| `discovery_target.nginx_instance.nginx_instance.uid` | [discovery_target.nginx_instance.nginx_instance.uid](data-sources--nginx_service_discovery--reference--group-001.md#canonical-14b9513ac1d9d1721980bcfb1aead7c7b05be40833506cea9316941d84c34bbc) |
| `id` | [id](data-sources--nginx_service_discovery--reference--group-001.md#canonical-566d300d0084480e42f2d0c311f536abef08662406ea7b8c1b40d0a754e0181f) |
| `labels` | [labels](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1e494c3ada4d725e34b0941a714a4d791bd1765c64f8a8bfa6df394183dd7d1f) |
| `name` | [name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-6315c0028c947c62fb51f86f8cd7e15bf7d4f368ee66d0488b3b14a9c3e316fa) |
| `namespace` | [namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-ae995acaccde3e2a85e4d7302095eba73104912d0fa801949a09792b9bcec6ac) |
| `server_block_filters` | [server_block_filters](data-sources--nginx_service_discovery--reference--group-001.md#canonical-88cc217792f73beea10a1dcf9d01165332e60d8313f12602f1684167fd6509d6) |
| `server_block_filters.name_regex` | [server_block_filters.name_regex](data-sources--nginx_service_discovery--reference--group-001.md#canonical-e5d1491fe069fcf2e50ab9a4b65c74f939daf1cd6649dbd09272d52a9c1a2d95) |
| `server_block_filters.port_ranges` | [server_block_filters.port_ranges](data-sources--nginx_service_discovery--reference--group-001.md#canonical-ff96280429062254b79586344713ddbce255eae16ac2f873ba26b1961ffb84cb) |

<a id="canonical-523d822f2d764a96afb015df625304b21cfb36beab826d11eb7a64c4debbfe82"></a>

## Next pages — Property reference / b63700e5a8d7 / 11

- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- [server_block_filters](data-sources--nginx_service_discovery--reference--group-001.md#canonical-f15c85d7318332866b803af7bf7d9529e778ac3afb78c0c3a20d92f8579b85f2)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd9af58e40ab47a58a98f697ee1f5cde9a23dd7766a77f99e82576c2e68dcacb"></a>

## discovery_target — discovery_target / f9be7b5827b6 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- discovery_target

<a id="canonical-df50e99ad9169135d31e6573baac1c344858bd5564560cf5d1147a873351998f"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery target.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"config_sync_group\",\"nginx_instance\"]"
}
```

<a id="canonical-b264f59860f5652c4da4e8682190be1e5aadd4dd1866f826d46e5190c5cf9be1"></a>

## Direct properties — discovery_target / f9be7b5827b6 / 3

- [config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-59826215131aa0158675c1c91bb66d34c938a26b21ac1906940f9086a20592a9): complete subsection reference.

- [nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-30b9170c01bb0964331171f1fa7673e1af07aa29a2b2025a6e33041c7350251d): complete subsection reference.

<a id="canonical-7aaf53f3b49c59e4b2093e60955e3304fd9e088f6ae2d32c97fa89eb20417f0a"></a>

## Next pages — discovery_target / f9be7b5827b6 / 4

- [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-59826215131aa0158675c1c91bb66d34c938a26b21ac1906940f9086a20592a9)
- [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-30b9170c01bb0964331171f1fa7673e1af07aa29a2b2025a6e33041c7350251d)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-59826215131aa0158675c1c91bb66d34c938a26b21ac1906940f9086a20592a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3722811b19c333d50becbef3c87e7e8d860db59ca0ad937718774fbd777ba273"></a>

## discovery_target.config_sync_group — discovery_target.config_sync_group / 8b7f3f2a1c43 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- discovery_target.config_sync_group

<a id="canonical-1665e9513d05689fe97ac44bae1dd5b45fce0fdcbf0b11234d2a84b60573a59c"></a>

Type: `"single"`. Computed.

Configuration parameter for config sync group.

Upstream description:

Select new ConfigSyncGroup.

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

<a id="canonical-684e77568e2bff1183e488724018326ee81aade4443b63003dc385865bc711cc"></a>

## Direct properties — discovery_target.config_sync_group / 8b7f3f2a1c43 / 3

- [config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-621643620ded985ef65095913c6f24c9086d74b3beaea35a4b338d66b77dda0f): complete subsection reference.

<a id="canonical-8c1aa7b9b3adc68e1fb11e6e7159965e1e73be5500baa3a08737411aad6e20ae"></a>

## Next pages — discovery_target.config_sync_group / 8b7f3f2a1c43 / 4

- [discovery_target.config_sync_group.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-621643620ded985ef65095913c6f24c9086d74b3beaea35a4b338d66b77dda0f)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-621643620ded985ef65095913c6f24c9086d74b3beaea35a4b338d66b77dda0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709f65d071ca558c4aaf0411cc858a11233e5d98aec508136fc43a2da99dcc93"></a>

## discovery_target.config_sync_group.config_sync_group — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-59826215131aa0158675c1c91bb66d34c938a26b21ac1906940f9086a20592a9)
- discovery_target.config_sync_group.config_sync_group

<a id="canonical-dda2a775c484343297312f41b7038bdb52b99e6f175927c45dc2d77e772c0fb9"></a>

Type: `"list"`. Computed.

Reference. Select new ConfigSyncGroup.

Upstream description:

Select new ConfigSyncGroup.

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

<a id="canonical-bf36d036d13088a20691e7a1cb922af02c25c03c6ca1d558f08e3108da87bda4"></a>

## Direct properties — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 3

<a id="canonical-fd557371eaddf52d858d1fae3363193e6736bef6d1c07c7de564fb1f62eee9f3"></a>

<a id="canonical-ccc7c1c98124f4898f131435e6199b91211982636f42c86564e7b12efdd77e93"></a>

## kind property — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 4

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

<a id="canonical-b10d468a7fe446ec8d06547a29a8eb2d272c1b564c185ec9c1cf55e01a2fbc02"></a>

<a id="canonical-18d8a5fc486e0653284c8397d6a552fbb5c88c08ab7832b41525093ade4154e0"></a>

## name property — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 5

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

<a id="canonical-907f9251705262beaafd86c98156a0c141d23dd5fb2d90fda4b45ae65d04f4c8"></a>

<a id="canonical-3fbe6182aadc760dfa78228d07901c06aa47d0d1745baac9f8ffa9aab49ef79d"></a>

## namespace property — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 6

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

<a id="canonical-fd1ac596e90e03bee7a3ede32350309ed6d412394900f403648416b7844dbb1b"></a>

<a id="canonical-ae6dc508298086318eeb0baed178dcd4f31507e6b10b54196176693f4acd8ce9"></a>

## tenant property — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 7

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

<a id="canonical-2fc00dcdc349aca5092404744197164a1028851c96af75c3f5b72385e99025c4"></a>

<a id="canonical-20fddd1cb6eec3b46d902dcbfea25fe64d7b50aa50b283a1a34d4f40317b04d7"></a>

## uid property — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 8

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

<a id="canonical-b97940fed082aa21deb58d799bca153d0bc8b4fb9b44c2c8aa287c7662cfad6e"></a>

## Next pages — discovery_target.config_sync_group.config_sync_group / 7300f414db34 / 9

- [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-59826215131aa0158675c1c91bb66d34c938a26b21ac1906940f9086a20592a9)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-30b9170c01bb0964331171f1fa7673e1af07aa29a2b2025a6e33041c7350251d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-347f519e7765a06bb205c3d3bbc390318aed53e8126fcab13c063c4a442b595d"></a>

## discovery_target.nginx_instance — discovery_target.nginx_instance / a959c85f5e12 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- discovery_target.nginx_instance

<a id="canonical-2b2da8085b6a119323a02684d32c36cf90734d7f6ae0f10887cbb45554bf4cc1"></a>

Type: `"single"`. Computed.

NGINXInstance Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

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

<a id="canonical-99a2b34b219a2c0e19f8c58cabd51333c81b913141a1638c496d4b89763fbe17"></a>

## Direct properties — discovery_target.nginx_instance / a959c85f5e12 / 3

- [nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-89a9eb2a7472920d9468cfdcf8e6c42f5c4f130a169b894e05a59cf7fa23817b): complete subsection reference.

<a id="canonical-38d3eb6171ae4c19a19c74ebe2bbc051caf5a314eeb2b52df0c7ff6fca562400"></a>

## Next pages — discovery_target.nginx_instance / a959c85f5e12 / 4

- [discovery_target.nginx_instance.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-89a9eb2a7472920d9468cfdcf8e6c42f5c4f130a169b894e05a59cf7fa23817b)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-89a9eb2a7472920d9468cfdcf8e6c42f5c4f130a169b894e05a59cf7fa23817b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1ff0b71dc93c0dec3c533c7fb4c73d5c44f2ec9185647003c758944f6d1905e"></a>

## discovery_target.nginx_instance.nginx_instance — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-28db11dc7b5bfd314bd7112055e2a79296492810e0503d4dcd9bdbc70848f365)
- [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-30b9170c01bb0964331171f1fa7673e1af07aa29a2b2025a6e33041c7350251d)
- discovery_target.nginx_instance.nginx_instance

<a id="canonical-93a79813669bce397ead21cb9016a9046bec9b1480a1f74d73075d0e4fbc4394"></a>

Type: `"list"`. Computed.

Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

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

<a id="canonical-5365d50a08644b7001b0c7e1c67dc551aa428d145492b2672a6498153b693f5a"></a>

## Direct properties — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 3

<a id="canonical-22ddeb2a5ed830c13e865c0697b67e4d4650fcd68532e53ce80d5f1b18139070"></a>

<a id="canonical-5eef0f72bd4b0349aca686ae56a0d8c6f21a13f179b376fcee27a53e9b0b0099"></a>

## kind property — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 4

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

<a id="canonical-1d76bbbecdf5c5ff0f23e0e345786e4d4b733f0fd1a1a4b7ec7dc086e5617d73"></a>

<a id="canonical-7a5587834ef75282f2f1313571bb548d3a707c82d692dcde1bc05a5881eee18c"></a>

## name property — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 5

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

<a id="canonical-fe725669b7b2a8875342334dba43ae6980a3327f0f14aa14652fba6fb409ff44"></a>

<a id="canonical-87d28437d27d9194157b1f791e0d0c76f441353d9b1dffb1f07c6d2c9554b034"></a>

## namespace property — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 6

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

<a id="canonical-1e34ce1e04f1f439910052f21d25a1f2112045d6190a7cf611058532333648be"></a>

<a id="canonical-4d9a77349cf8da20dac10eca049f6668b86c53c202721bd027027368921d636f"></a>

## tenant property — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 7

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

<a id="canonical-14b9513ac1d9d1721980bcfb1aead7c7b05be40833506cea9316941d84c34bbc"></a>

<a id="canonical-5c2e05c4f9d95a5fc35010ea10f6f9392e12c4ff99431af542d2adcddfa407a4"></a>

## uid property — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 8

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

<a id="canonical-5310e28b23ab944dcddc394ef2dbf28bb7ae155cec4d792fa4c99cb6d94abdea"></a>

## Next pages — discovery_target.nginx_instance.nginx_instance / cf45fa523d01 / 9

- [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-30b9170c01bb0964331171f1fa7673e1af07aa29a2b2025a6e33041c7350251d)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-f15c85d7318332866b803af7bf7d9529e778ac3afb78c0c3a20d92f8579b85f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-861f1b705692470ccdf0a2e503a99fb24f5b8fb0cdd83c99bcef6d44f8d9be1a"></a>

## server_block_filters — server_block_filters / aef7f4f30d00 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- server_block_filters

<a id="canonical-88cc217792f73beea10a1dcf9d01165332e60d8313f12602f1684167fd6509d6"></a>

Type: `"list"`. Computed.

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks
will be discovered by default.

Upstream description:

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter.

X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.

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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7da73415eee5e371a64ffc3f06c47201366e3a152735b4d16ca49fbd469a9e2e"></a>

## Direct properties — server_block_filters / aef7f4f30d00 / 3

<a id="canonical-e5d1491fe069fcf2e50ab9a4b65c74f939daf1cd6649dbd09272d52a9c1a2d95"></a>

<a id="canonical-75f87233655e72be32901ebbff3afa869f0e3446636b071b11c0a9aabb249084"></a>

## name_regex property — server_block_filters / aef7f4f30d00 / 4

Type: `"string"`. Computed.

Regular expression to match the server name or domain that must be discovered.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-ff96280429062254b79586344713ddbce255eae16ac2f873ba26b1961ffb84cb"></a>

<a id="canonical-4377a976e82588af994f9eee4101cc7d15071deabab69c730a24f6e97c9addd3"></a>

## port_ranges property — server_block_filters / aef7f4f30d00 / 5

Type: `"string"`. Computed.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Upstream description:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". For example, 8000-8191. Maximum
number of ports allowed is 1024.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-8a99c5a75751854a7f644efdcef9b66052d90ba55c8314438503e4b6ee45b02c"></a>

## Next pages — server_block_filters / aef7f4f30d00 / 6

- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
