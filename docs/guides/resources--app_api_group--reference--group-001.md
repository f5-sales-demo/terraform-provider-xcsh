---
page_title: "xcsh_app_api_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group reference."
---

# xcsh_app_api_group reference

<a id="canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d181beab7112abac7bb7bdbceb4ce8c7bb3cbca39e3c19e495c32c88700ca045"></a>

## Property reference — Property reference / eb8801fe5326 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- Property reference

<a id="canonical-0bb17d02b5d8859fb24e17ee40c1954903688c8f49a15d7cd66c514acad02ba1"></a>

## Direct properties — Property reference / eb8801fe5326 / 3

<a id="canonical-16fda29c3351fa907cb0c393f42f369ea72ce4c21ea27c13bd5cd622dc569083"></a>

<a id="canonical-dd9458d118da44c78337c2ecf47c1826657c39fa3e27f7ff9e3c8bca414df338"></a>

## annotations property — Property reference / eb8801fe5326 / 4

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

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-4da9c95349bdbd11d45b9aa30ddec8bdafe31d948ab9baf2912bf79d79d38665): complete subsection reference.

- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-9ad14a19759041b89c41886386e9b3d16043982d9d1dac1d8ec223b393ed9dbf): complete subsection reference.

<a id="canonical-6e6ac5b87a656462b58a68a8933e5eae72db729e572995edcf03e222f4047b3d"></a>

<a id="canonical-23cc43368087d18af9fb742b31c4461f8cffeb99974e0bfa89ddd47fa75abcf4"></a>

## description property — Property reference / eb8801fe5326 / 5

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

<a id="canonical-0d23d4311518aec01a8116076aef20f35d9da77670358b7d571958a12e1ecdcf"></a>

<a id="canonical-877366ce465cf019a590647117f6033214282d5bf8279fb23b7094c7f34536e4"></a>

## disable property — Property reference / eb8801fe5326 / 6

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

- [elements](resources--app_api_group--reference--group-001.md#canonical-60598adaa18ffa381545f0e38c3d5b7a9a43d5b31ede39e728ca654f13b3b201): complete subsection reference.

- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-157613b0b1812ead82ff72ffc24dcba5266d41484cdd1c0de946731b0ece2bce): complete subsection reference.

<a id="canonical-6885183fd37bbf1429b247961e98f77f2f20ecf5e7f17523827e2ac23a06c756"></a>

<a id="canonical-fae78829c0c24aef61f158a5bccf79856d6f0b8a7141ba777164386ab02dada7"></a>

## id property — Property reference / eb8801fe5326 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f3f5cd06bcfc52740195044708517cc55647b02b2a97818c43dac1a9dfa20256"></a>

<a id="canonical-20169e70cfd1e83ff9f62ac4b7b8800f6cc50dd98ac935cf68c81aa607d1bb57"></a>

## labels property — Property reference / eb8801fe5326 / 8

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

<a id="canonical-17a70b42084a7fe082f214b93b19b56f2568fec8beb9665d694b15dfb97def78"></a>

<a id="canonical-723ff77ec2060969777bfdf72a907419a807e3251a330677eece6e4182fafc66"></a>

## name property — Property reference / eb8801fe5326 / 9

Type: `"string"`. Required.

Name of the App API Group. Must be unique within the namespace.

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

<a id="canonical-0d589a6c57ae45e2ff67499f01ad49ee3ba269c6edc1d917a61f3e53fed8260f"></a>

<a id="canonical-c677a1dead676a71447400b89ac3c11ee4ed5b3136e803c708a3b6311b00ce84"></a>

## namespace property — Property reference / eb8801fe5326 / 10

Type: `"string"`. Required.

Namespace where the App API Group is created.

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

- [timeouts](resources--app_api_group--reference--group-001.md#canonical-33040bf565bb7fa870d0b92a51ce1ca2d296c992dcc47fcc2258d293e450f757): complete subsection reference.

<a id="canonical-4896feb094984a7fd6357fcaba66dccb8240a52973ae310cfd9cb5e2ab2b2337"></a>

## All schema paths — Property reference / eb8801fe5326 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_api_group--reference--group-001.md#canonical-16fda29c3351fa907cb0c393f42f369ea72ce4c21ea27c13bd5cd622dc569083) |
| `bigip_virtual_server` | [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-efa8c9ac150e1d59d933109719b95d2677846ca042d79b5d135088ba875533ae) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-4076f449450b3e908f65230f0c6d452d274413a870e4e1c285d938cea154d6a8) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](resources--app_api_group--reference--group-001.md#canonical-a0cb808d82a4b32aaf79145beecfd74250556beae6848b055a6c0614cf706ff1) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](resources--app_api_group--reference--group-001.md#canonical-3b2d5de8de2c38c983070c4c70701da80c6f09d97db2f481f9e0aaec4ae9bfcc) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](resources--app_api_group--reference--group-001.md#canonical-0c584c711cd1be941e5959161b7a2c76acea7f9953bc563e2be36ccd69ea7134) |
| `cdn_loadbalancer` | [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-8fa158e1a991154f5239984935a422abe70b9f6816fc86d809c874776f46b43c) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-af6fdca61f5bcab354a377684f03094e018d0a6f17afb79abeff86f2ab725d76) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](resources--app_api_group--reference--group-001.md#canonical-2d2ff5c75575ce4ca391da3f926cbbc9df5a6aa37a63cdbbaac08f3bf3051414) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](resources--app_api_group--reference--group-001.md#canonical-93e677a01fdcc7e6742a1e8f24d9a233e7c8672b7ff36b8843eeff88da8a47b4) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](resources--app_api_group--reference--group-001.md#canonical-908633b723c26fa9b6d1091d191ff232ef1a82709be9ab0924881db68c8cc0d0) |
| `description` | [description](resources--app_api_group--reference--group-001.md#canonical-6e6ac5b87a656462b58a68a8933e5eae72db729e572995edcf03e222f4047b3d) |
| `disable` | [disable](resources--app_api_group--reference--group-001.md#canonical-0d23d4311518aec01a8116076aef20f35d9da77670358b7d571958a12e1ecdcf) |
| `elements` | [elements](resources--app_api_group--reference--group-001.md#canonical-4c3f62b99cba5b4aa5099dce9f442dff2cfc45b4661466f09ffb9c01e4d13ae9) |
| `elements.methods` | [elements.methods](resources--app_api_group--reference--group-001.md#canonical-5ed62c499e97698a08a4aeb0f7cf33ec8493c77e56fc193ff24d5c8f8a38b3fc) |
| `elements.path_regex` | [elements.path_regex](resources--app_api_group--reference--group-001.md#canonical-555e9ddfcca15716247ed9177352f91b1efded2ececda241886a1f6cb964e4ea) |
| `http_loadbalancer` | [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-37dde52b55c04dd4a403430c35212bb49cc363a7929eea330fb804ec2e171b16) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-7de8a9d88a7c78cc8c60ada37d7d8491941300fdad15035b82a42f1300cf2a45) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](resources--app_api_group--reference--group-001.md#canonical-50bbde334f1beac1e9b2d854204b7c37c03a987e73bb30891acc6b2597aebc00) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](resources--app_api_group--reference--group-001.md#canonical-6a3cf156ab24c0d7739f6370aabc5145a19e27a445e08e016662844181396ada) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](resources--app_api_group--reference--group-001.md#canonical-7de24ec28299aef238f6066c65bb3cb01ef75877b60d5d358ee43413c697041e) |
| `id` | [id](resources--app_api_group--reference--group-001.md#canonical-6885183fd37bbf1429b247961e98f77f2f20ecf5e7f17523827e2ac23a06c756) |
| `labels` | [labels](resources--app_api_group--reference--group-001.md#canonical-f3f5cd06bcfc52740195044708517cc55647b02b2a97818c43dac1a9dfa20256) |
| `name` | [name](resources--app_api_group--reference--group-001.md#canonical-17a70b42084a7fe082f214b93b19b56f2568fec8beb9665d694b15dfb97def78) |
| `namespace` | [namespace](resources--app_api_group--reference--group-001.md#canonical-0d589a6c57ae45e2ff67499f01ad49ee3ba269c6edc1d917a61f3e53fed8260f) |
| `timeouts` | [timeouts](resources--app_api_group--reference--group-001.md#canonical-da7f7f5aeef941a430fe48669bdf2dd9bed4e6e5f57fc0d9df009545925fdbda) |
| `timeouts.create` | [timeouts.create](resources--app_api_group--reference--group-001.md#canonical-02f940832f7a9e5cd8f86540fc59687d96503b62f2267966c4757f8c3e426700) |
| `timeouts.delete` | [timeouts.delete](resources--app_api_group--reference--group-001.md#canonical-49d598c4cc92a96a3fe543d5dd0556caf06148112ea9eedac3aaae6d7c6908bf) |
| `timeouts.read` | [timeouts.read](resources--app_api_group--reference--group-001.md#canonical-f95dc5a6fe786ec85327329a395bb42717dda5f546d46613c7dbc9c6ea561600) |
| `timeouts.update` | [timeouts.update](resources--app_api_group--reference--group-001.md#canonical-7cff0b7937fe7afe7749f8d40ef8be9741aca007a7e59b3a1d1dea04f3ae2fa0) |

<a id="canonical-1047985b02c64c273b51ad68369b260edafd1c9e204596f48f3e9282e7a484ae"></a>

## Next pages — Property reference / eb8801fe5326 / 12

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-4da9c95349bdbd11d45b9aa30ddec8bdafe31d948ab9baf2912bf79d79d38665)
- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-9ad14a19759041b89c41886386e9b3d16043982d9d1dac1d8ec223b393ed9dbf)
- [elements](resources--app_api_group--reference--group-001.md#canonical-60598adaa18ffa381545f0e38c3d5b7a9a43d5b31ede39e728ca654f13b3b201)
- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-157613b0b1812ead82ff72ffc24dcba5266d41484cdd1c0de946731b0ece2bce)
- [timeouts](resources--app_api_group--reference--group-001.md#canonical-33040bf565bb7fa870d0b92a51ce1ca2d296c992dcc47fcc2258d293e450f757)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-4da9c95349bdbd11d45b9aa30ddec8bdafe31d948ab9baf2912bf79d79d38665"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e9008aa6d4385b8ca75798f345878673ec488c6c17f4af6829f65f05840262"></a>

## bigip_virtual_server — bigip_virtual_server / 1cd64cbab9d7 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- bigip_virtual_server

<a id="canonical-efa8c9ac150e1d59d933109719b95d2677846ca042d79b5d135088ba875533ae"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

Upstream description:

Set the scope of the API Group to a specific BIG-IP Virtual Server.

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

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-efa8c9ac150e1d59d933109719b95d2677846ca042d79b5d135088ba875533ae)
- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-8fa158e1a991154f5239984935a422abe70b9f6816fc86d809c874776f46b43c)
- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-37dde52b55c04dd4a403430c35212bb49cc363a7929eea330fb804ec2e171b16)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bigip_virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-d9a5b8dacf4b601cc764c4968c7e3b2f41722e195001c74a19e42eb0a5511ed2"></a>

## Direct properties — bigip_virtual_server / 1cd64cbab9d7 / 3

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-1383c37d5840ab9611145a3b2bb7125c2a09d5df91f4997b8d444d08fe09c455): complete subsection reference.

<a id="canonical-2a477544d8aca88179d9d164dd8d6c437d0ed2a7ecb4e847d3ec5255b912995d"></a>

## Next pages — bigip_virtual_server / 1cd64cbab9d7 / 4

- [bigip_virtual_server.bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-1383c37d5840ab9611145a3b2bb7125c2a09d5df91f4997b8d444d08fe09c455)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-1383c37d5840ab9611145a3b2bb7125c2a09d5df91f4997b8d444d08fe09c455"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61f52da9057597a056efb5681c2c0044576134778980c6e104884ca3e917fbe0"></a>

## bigip_virtual_server.bigip_virtual_server — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-4da9c95349bdbd11d45b9aa30ddec8bdafe31d948ab9baf2912bf79d79d38665)
- bigip_virtual_server.bigip_virtual_server

<a id="canonical-4076f449450b3e908f65230f0c6d452d274413a870e4e1c285d938cea154d6a8"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
bigip_virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-16fab9d7efedda2ca488391e62044f377a92fa443388ae81086445bcd07c4052"></a>

## Direct properties — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 3

<a id="canonical-a0cb808d82a4b32aaf79145beecfd74250556beae6848b055a6c0614cf706ff1"></a>

<a id="canonical-f0c88e04ac3e5e93693d7def0c7f5fec9aea479c844f7cd360685cae2fa66f56"></a>

## name property — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3b2d5de8de2c38c983070c4c70701da80c6f09d97db2f481f9e0aaec4ae9bfcc"></a>

<a id="canonical-24643f2682912092452649eb46b6f9200ff1825869231bf093a78737733a4051"></a>

## namespace property — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0c584c711cd1be941e5959161b7a2c76acea7f9953bc563e2be36ccd69ea7134"></a>

<a id="canonical-0b7daf6c3a7e1c08fd0322e56c3d4f7a20424d9ef56bfb476b2b9ee70ea8abd8"></a>

## tenant property — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-ff462d9a791dac4b53acd9106925a905525bdc416f752338015b6574fe9792ad"></a>

## Next pages — bigip_virtual_server.bigip_virtual_server / 5ee53c6e66f7 / 7

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-4da9c95349bdbd11d45b9aa30ddec8bdafe31d948ab9baf2912bf79d79d38665)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-9ad14a19759041b89c41886386e9b3d16043982d9d1dac1d8ec223b393ed9dbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51e2c9dffbcb73f0bddb9874bfc354c5ae140738dd4ac2033f867d904f193c55"></a>

## cdn_loadbalancer — cdn_loadbalancer / e2360fb95fbb / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- cdn_loadbalancer

<a id="canonical-8fa158e1a991154f5239984935a422abe70b9f6816fc86d809c874776f46b43c"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific CDN Loadbalancer.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-55ef02b1313e8ec75e1187853bf0ab65fcb874d515bf5674efe04b687e496995"></a>

## Direct properties — cdn_loadbalancer / e2360fb95fbb / 3

- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-42c2359a5a889d8395ca45530f8f48376fc505091161f87c78f9dd396fca7b4f): complete subsection reference.

<a id="canonical-1c0b82ed713c47e90a1f524c95f92062b84aa8bfeb9192e0f5ebfc1a4f5776e3"></a>

## Next pages — cdn_loadbalancer / e2360fb95fbb / 4

- [cdn_loadbalancer.cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-42c2359a5a889d8395ca45530f8f48376fc505091161f87c78f9dd396fca7b4f)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-42c2359a5a889d8395ca45530f8f48376fc505091161f87c78f9dd396fca7b4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ee220328fc36750cb8814e706d107df43ccf99a59611a65e18ff258af83811c"></a>

## cdn_loadbalancer.cdn_loadbalancer — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-9ad14a19759041b89c41886386e9b3d16043982d9d1dac1d8ec223b393ed9dbf)
- cdn_loadbalancer.cdn_loadbalancer

<a id="canonical-af6fdca61f5bcab354a377684f03094e018d0a6f17afb79abeff86f2ab725d76"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ef5fdb85cfe7690bb327f88ddd90d5772b8f92d644bc65d8b159bd0e0e37156"></a>

## Direct properties — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 3

<a id="canonical-2d2ff5c75575ce4ca391da3f926cbbc9df5a6aa37a63cdbbaac08f3bf3051414"></a>

<a id="canonical-e2288631097c9a90ec48dc9afd1442f1a004c6343f3f3c3a462cca0bcd998b16"></a>

## name property — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-93e677a01fdcc7e6742a1e8f24d9a233e7c8672b7ff36b8843eeff88da8a47b4"></a>

<a id="canonical-bbc05d8b706e3a8063144efcada87662b435fa53999525cfd7ea17e12489a271"></a>

## namespace property — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-908633b723c26fa9b6d1091d191ff232ef1a82709be9ab0924881db68c8cc0d0"></a>

<a id="canonical-b90cf5b1ca53767ff3a5bb5b439f962b8327c32c6b896b1f6462ef9caf830973"></a>

## tenant property — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-46f2099bbb52606a14763c2670ae189fab07ac96f8ed158b0f81b16827f78e15"></a>

## Next pages — cdn_loadbalancer.cdn_loadbalancer / 6dbd12b388de / 7

- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-9ad14a19759041b89c41886386e9b3d16043982d9d1dac1d8ec223b393ed9dbf)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-60598adaa18ffa381545f0e38c3d5b7a9a43d5b31ede39e728ca654f13b3b201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-415e0bf2caf83056326cbe5a4238a4689b1ca76ff951d6ea2262becb3ae7a043"></a>

## elements — elements / 666ea51bc177 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- elements

<a id="canonical-4c3f62b99cba5b4aa5099dce9f442dff2cfc45b4661466f09ffb9c01e4d13ae9"></a>

Type: `"object"`. list nested block, Optional.

List of API group elements with methods and path regex for matching requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("methods",
    "path_regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  }
}
```

Terraform syntax:

```terraform
elements {
  # Configure direct properties listed below.
}
```

<a id="canonical-d24b7357757aeca925676e7894e0948e8b2c5187a294864d768d1e995d32289e"></a>

## Direct properties — elements / 666ea51bc177 / 3

<a id="canonical-5ed62c499e97698a08a4aeb0f7cf33ec8493c77e56fc193ff24d5c8f8a38b3fc"></a>

<a id="canonical-bc04319f9f65f6df97599b3191c05684e759f1de3b0c75794d179c2720b1c3ff"></a>

## methods property — elements / 666ea51bc177 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of method values to
match the input request API method against. The match is considered to succeed if the input request
API method is a member of the list. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`,
\`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of method values to match the input request API method against. The match is considered to
succeed if the input request API method is a member of the list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-555e9ddfcca15716247ed9177352f91b1efded2ececda241886a1f6cb964e4ea"></a>

<a id="canonical-b74883cf204073e55d69c5a44679539fede4d7591b03e2cef4e0260257aade1e"></a>

## path_regex property — elements / 666ea51bc177 / 5

Type: `"string"`. Optional.

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

Upstream description:

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2ddb190826897d96512a52c3880894c2845b5845e3882aa8d2470ebe2427b9ee"></a>

## Next pages — elements / 666ea51bc177 / 6

- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-157613b0b1812ead82ff72ffc24dcba5266d41484cdd1c0de946731b0ece2bce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ded2de6ff993be1ef2e4fa0dc05673e2ed2262dd697a3c7bdce65ee97f8edfb2"></a>

## http_loadbalancer — http_loadbalancer / e38755d83004 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- http_loadbalancer

<a id="canonical-37dde52b55c04dd4a403430c35212bb49cc363a7929eea330fb804ec2e171b16"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee122816db4dc2d4f163874ef33c058109073c9023a23c9a9c04feb07d448860"></a>

## Direct properties — http_loadbalancer / e38755d83004 / 3

- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-23e862c0b97afced8edbc5bd7df3f3632e785c0b0b135e9e2290fbdcf6e6c80a): complete subsection reference.

<a id="canonical-ff14491cba73b7b1fe09fa3050ade29f3650d087dc3652a9addc93f6a84f31e1"></a>

## Next pages — http_loadbalancer / e38755d83004 / 4

- [http_loadbalancer.http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-23e862c0b97afced8edbc5bd7df3f3632e785c0b0b135e9e2290fbdcf6e6c80a)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-23e862c0b97afced8edbc5bd7df3f3632e785c0b0b135e9e2290fbdcf6e6c80a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-818a5fc6273dda431d188a7db476ac473a90cf2add26cc75b992ce90cd58a49a"></a>

## http_loadbalancer.http_loadbalancer — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-157613b0b1812ead82ff72ffc24dcba5266d41484cdd1c0de946731b0ece2bce)
- http_loadbalancer.http_loadbalancer

<a id="canonical-7de8a9d88a7c78cc8c60ada37d7d8491941300fdad15035b82a42f1300cf2a45"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee136ec119f36394164437e00d71bb2111085787cc7623be8e50c4546cf26eef"></a>

## Direct properties — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 3

<a id="canonical-50bbde334f1beac1e9b2d854204b7c37c03a987e73bb30891acc6b2597aebc00"></a>

<a id="canonical-78427dff951d25124ff62fb749713fa46868cb4edba0ddb0efec2ef59bc8c1b2"></a>

## name property — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-6a3cf156ab24c0d7739f6370aabc5145a19e27a445e08e016662844181396ada"></a>

<a id="canonical-fca4b933dadd84acc7309bf12e1cf172a056983949cab5c74ead61f3a630dff3"></a>

## namespace property — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-7de24ec28299aef238f6066c65bb3cb01ef75877b60d5d358ee43413c697041e"></a>

<a id="canonical-34a551effa11dbcfa18560a01085b0750982607e202c5b543b5c917aa9fc16e5"></a>

## tenant property — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2b3be34562a1496d35186768949185d86e8d175d5c883d57c59ccb1cc2b8972c"></a>

## Next pages — http_loadbalancer.http_loadbalancer / c36efa15ee89 / 7

- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-157613b0b1812ead82ff72ffc24dcba5266d41484cdd1c0de946731b0ece2bce)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-33040bf565bb7fa870d0b92a51ce1ca2d296c992dcc47fcc2258d293e450f757"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9f1e62471d94332e3d259e1d44ba5ff15f4c05fb65b7592e26b14c99815fd41"></a>

## timeouts — timeouts / f962d744b129 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- timeouts

<a id="canonical-da7f7f5aeef941a430fe48669bdf2dd9bed4e6e5f57fc0d9df009545925fdbda"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-db562d371351d71dfe765797b2e76b3e4e2ea5aed492907ba117b7be6774f994"></a>

## Direct properties — timeouts / f962d744b129 / 3

<a id="canonical-02f940832f7a9e5cd8f86540fc59687d96503b62f2267966c4757f8c3e426700"></a>

<a id="canonical-c4f9587badadc5000c60aa721cb537a3ee5a4d6a3afa3c10ce96d3e1be070352"></a>

## create property — timeouts / f962d744b129 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-49d598c4cc92a96a3fe543d5dd0556caf06148112ea9eedac3aaae6d7c6908bf"></a>

<a id="canonical-f517047e8bf95c6eec81b6eabfee457395b10d03d4c495fd6506b1e4d502a2b1"></a>

## delete property — timeouts / f962d744b129 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-f95dc5a6fe786ec85327329a395bb42717dda5f546d46613c7dbc9c6ea561600"></a>

<a id="canonical-636fe9f340530c1515f5f671fced479962c175c320240c83c0eeadb60f2d6e76"></a>

## read property — timeouts / f962d744b129 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-7cff0b7937fe7afe7749f8d40ef8be9741aca007a7e59b3a1d1dea04f3ae2fa0"></a>

<a id="canonical-b5464561f24e1121c9390758dbe2db67949dd58bce94c30ce64dfaa56993be89"></a>

## update property — timeouts / f962d744b129 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a3aa9eb39924bafd756ba894a25dc4ec867f1b274957abf74cfd0ccb2d01b7c4"></a>

## Next pages — timeouts / f962d744b129 / 8

- [Property reference](resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
