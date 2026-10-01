---
page_title: "xcsh_api_crawler reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler reference."
---

# xcsh_api_crawler reference

<a id="canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2231993ea9f69d7bbb661d34943fb4c235f12f2cd35694a7a852380669b4c3f"></a>

## Property reference — Property reference / bc165e95dabb / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- Property reference

<a id="canonical-729d9a0c3d2afc723ae583e2eb9ebba54f41cae4c7fdfacf450493f8fb9ba94c"></a>

## Direct properties — Property reference / bc165e95dabb / 3

<a id="canonical-0090ac99960f8c6aded64d213a5e67c37459e04f9810e16d86dd92117f0bdc84"></a>

<a id="canonical-04305f173e6ac2b11fdf5403d0c229ea4dcdba0ffeef720adfc8f7a347aa591c"></a>

## annotations property — Property reference / bc165e95dabb / 4

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

<a id="canonical-1a0357878db29d0484e93f81b5cf794c21c49c44a8721f820555c7ef670c5bff"></a>

<a id="canonical-89d23e36b88f6a987f2dc39a1b01f5bcc826b4c2580084947999a7599cc14a00"></a>

## description property — Property reference / bc165e95dabb / 5

Type: `"string"`. Computed.

Description of the APICrawler.

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

- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44): complete subsection reference.

<a id="canonical-3332c2697586df75d4c43298d5f4609558e2ffe7029108caaa947cb2a963238e"></a>

<a id="canonical-2d3cf6e203717564e3236563aa16697eb8e7764ec12db88485ad670077f0aa87"></a>

## id property — Property reference / bc165e95dabb / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b9edbb06fc93cd02781b85df1c3c1c22e493a27feb44ad28285a1acf71199390"></a>

<a id="canonical-7c5a1035a3a25131b2a18ec7c783038f36feaba73249c214bb5fae4ef7b0f77c"></a>

## labels property — Property reference / bc165e95dabb / 7

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

<a id="canonical-dc5e41cad83ab29d80e48e4d2fe1862859820c2a2221081c3971e45019e3406f"></a>

<a id="canonical-c5fe158ee844be1cd344ce5cc3f8cf4a48f1a0ed6e7dca423cce6afa23886d03"></a>

## name property — Property reference / bc165e95dabb / 8

Type: `"string"`. Required.

Name of the APICrawler.

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

<a id="canonical-dd1c6eee23935035d8b87c11ec9d40cb5ae75e3c9902d149a5b1ab40b48f4eb5"></a>

<a id="canonical-453b0c7bd6a24db98fdf60e1496c2721e96e1731401922d5a08689043871ef77"></a>

## namespace property — Property reference / bc165e95dabb / 9

Type: `"string"`. Required.

Namespace where the APICrawler exists.

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

<a id="canonical-544fca49c043b83efb0e426a0fc77dd47d26bcc0713d34fbb96814c002c27f0d"></a>

## All schema paths — Property reference / bc165e95dabb / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_crawler--reference--group-001.md#canonical-0090ac99960f8c6aded64d213a5e67c37459e04f9810e16d86dd92117f0bdc84) |
| `description` | [description](data-sources--api_crawler--reference--group-001.md#canonical-1a0357878db29d0484e93f81b5cf794c21c49c44a8721f820555c7ef670c5bff) |
| `domains` | [domains](data-sources--api_crawler--reference--group-001.md#canonical-dd9ebf993d0fc9fef2ccb0fb462c227af362017a76e438faeb306a402e0653a8) |
| `domains.domain` | [domains.domain](data-sources--api_crawler--reference--group-001.md#canonical-7da2ffa3293801b55aaa6b9ae6e96b5f958e9d98cd11c4901c0db947d3593497) |
| `domains.simple_login` | [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-235d945f69ea549313b3bb49801e2c82c40dc2da895a173e1f8de9d90ca9e312) |
| `domains.simple_login.password` | [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-1434a77895e660d6faced17cd07e414de650edc09aaa76677a3781d0a68f409e) |
| `domains.simple_login.password.blindfold_secret_info` | [domains.simple_login.password.blindfold_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-bbcea523b89cc0a1707137a7540b79ed974b1fa24682cab3f59f78653c742280) |
| `domains.simple_login.password.blindfold_secret_info.decryption_provider` | [domains.simple_login.password.blindfold_secret_info.decryption_provider](data-sources--api_crawler--reference--group-001.md#canonical-08ddde8f793010b320ca09fab69c15faea82cfade85297d91f722d69244e6875) |
| `domains.simple_login.password.blindfold_secret_info.location` | [domains.simple_login.password.blindfold_secret_info.location](data-sources--api_crawler--reference--group-001.md#canonical-b94f299942041840401e84df61a1d3ec6e359b827186c899480086eff571bd97) |
| `domains.simple_login.password.blindfold_secret_info.store_provider` | [domains.simple_login.password.blindfold_secret_info.store_provider](data-sources--api_crawler--reference--group-001.md#canonical-fa25fbda81b9e779c0fbe532294316baabca4a581538dccfa104e95997fc560a) |
| `domains.simple_login.password.clear_secret_info` | [domains.simple_login.password.clear_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-00140828af538ea9c6cf85db012aa698ccd8d669b0e12ab0f63a31e7d0d8e399) |
| `domains.simple_login.password.clear_secret_info.provider_ref` | [domains.simple_login.password.clear_secret_info.provider_ref](data-sources--api_crawler--reference--group-001.md#canonical-167642f0d04eab8252b7e201e613e81d1671fce69645e4f8b97207fd3227457c) |
| `domains.simple_login.password.clear_secret_info.url` | [domains.simple_login.password.clear_secret_info.url](data-sources--api_crawler--reference--group-001.md#canonical-63d361e4ad2d7f425f7c65b76d0e2768dd65e2d95b128a5d9da4299af88819ac) |
| `domains.simple_login.user` | [domains.simple_login.user](data-sources--api_crawler--reference--group-001.md#canonical-a0656c4bc75e1ce52614f4a949e660a241783931a5ee7a9b52dcf891ba8f4d46) |
| `id` | [id](data-sources--api_crawler--reference--group-001.md#canonical-3332c2697586df75d4c43298d5f4609558e2ffe7029108caaa947cb2a963238e) |
| `labels` | [labels](data-sources--api_crawler--reference--group-001.md#canonical-b9edbb06fc93cd02781b85df1c3c1c22e493a27feb44ad28285a1acf71199390) |
| `name` | [name](data-sources--api_crawler--reference--group-001.md#canonical-dc5e41cad83ab29d80e48e4d2fe1862859820c2a2221081c3971e45019e3406f) |
| `namespace` | [namespace](data-sources--api_crawler--reference--group-001.md#canonical-dd1c6eee23935035d8b87c11ec9d40cb5ae75e3c9902d149a5b1ab40b48f4eb5) |

<a id="canonical-2a179a437d02ff47eedab4a82b577e6abdd9b06dd99669da4019b9b7ca333e76"></a>

## Next pages — Property reference / bc165e95dabb / 11

- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2e202f3864d68a777fe39e4d93768a9c0722ede566675dfa9febbd3f54dcb09"></a>

## domains — domains / 808f217d4604 / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- domains

<a id="canonical-dd9ebf993d0fc9fef2ccb0fb462c227af362017a76e438faeb306a402e0653a8"></a>

Type: `"list"`. Computed.

API Crawler. API Crawler Configuration.

Upstream description:

API Crawler Configuration.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-a56a81a1b9165fec80c8086ab73686ad81d29827027ac597d63d5374537ad01e"></a>

## Direct properties — domains / 808f217d4604 / 3

<a id="canonical-7da2ffa3293801b55aaa6b9ae6e96b5f958e9d98cd11c4901c0db947d3593497"></a>

<a id="canonical-bfd3f35072b3823fd358d208e0ffe7b66e43c6c306f8d23114d93a7af0307e90"></a>

## domain property — domains / 808f217d4604 / 4

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31): complete subsection reference.

<a id="canonical-4402bef8272e6ccd28bc20efd8e3b33f2f4c55413436cd4f6fb7f1de9e565f17"></a>

## Next pages — domains / 808f217d4604 / 5

- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5f5e10618d1a9023bde8fb4f85ecf658774022ed2cf5dc71300146b4a49daed"></a>

## domains.simple_login — domains.simple_login / ac8eeaeb5a91 / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- domains.simple_login

<a id="canonical-235d945f69ea549313b3bb49801e2c82c40dc2da895a173e1f8de9d90ca9e312"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-fb730a9cef23cd35ac18dabd6edeac8320e746f622cd5aa813a2a0649229ceb3"></a>

## Direct properties — domains.simple_login / ac8eeaeb5a91 / 3

- [password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2): complete subsection reference.

<a id="canonical-a0656c4bc75e1ce52614f4a949e660a241783931a5ee7a9b52dcf891ba8f4d46"></a>

<a id="canonical-b3ebc3bdb59d5c09d4574523cf7b95d8a6dd63c75b17b605a36172a1b5595b5c"></a>

## user property — domains.simple_login / ac8eeaeb5a91 / 4

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-e805e64e8ce35e81e357ab09cfa8e13f01d9660a68e86bd39044eb82292b7269"></a>

## Next pages — domains.simple_login / ac8eeaeb5a91 / 5

- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68b08d22188fc10ef64e133bcec6ecc64594b0ed856aa9e705f0785839b2102d"></a>

## domains.simple_login.password — domains.simple_login.password / 7c06f7042eea / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31)
- domains.simple_login.password

<a id="canonical-1434a77895e660d6faced17cd07e414de650edc09aaa76677a3781d0a68f409e"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-b077731c046c49c9a889a90c412d157b01fc2014079a0ee969b01850e453366a"></a>

## Direct properties — domains.simple_login.password / 7c06f7042eea / 3

- [blindfold_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-7ff42553a8a803158baaf54654b9dff4b001276c95e28cdb9d8231da13a6b8d4): complete subsection reference.

- [clear_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-3ab32d8cb66385b672c3ffbe720bc251fbf50ccf200d45e03d317d97f5d893af): complete subsection reference.

<a id="canonical-19fa5fc3064b8aa2af5cdc43f31ee69b2513a7ff100500738ad5de0492b0c588"></a>

## Next pages — domains.simple_login.password / 7c06f7042eea / 4

- [domains.simple_login.password.blindfold_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-7ff42553a8a803158baaf54654b9dff4b001276c95e28cdb9d8231da13a6b8d4)
- [domains.simple_login.password.clear_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-3ab32d8cb66385b672c3ffbe720bc251fbf50ccf200d45e03d317d97f5d893af)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-7ff42553a8a803158baaf54654b9dff4b001276c95e28cdb9d8231da13a6b8d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fecc9ba9fc682a16e949f0e3086cf98542619173cdf11cced7e0624056c2cba"></a>

## domains.simple_login.password.blindfold_secret_info — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31)
- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2)
- domains.simple_login.password.blindfold_secret_info

<a id="canonical-bbcea523b89cc0a1707137a7540b79ed974b1fa24682cab3f59f78653c742280"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-65575004328e000a7a8a33432b2caa1dafea5987339ee83b73666147732bc030"></a>

## Direct properties — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 3

<a id="canonical-08ddde8f793010b320ca09fab69c15faea82cfade85297d91f722d69244e6875"></a>

<a id="canonical-e5544868dda2c14ba29a618bb0ea0140b0ba7ce6a75888e6db9ae4fe5eccac2d"></a>

## decryption_provider property — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-b94f299942041840401e84df61a1d3ec6e359b827186c899480086eff571bd97"></a>

<a id="canonical-95e30c14e4035c59d7aee4ce75d93ac43730b06a82922f1fad22015baef4c598"></a>

## location property — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-fa25fbda81b9e779c0fbe532294316baabca4a581538dccfa104e95997fc560a"></a>

<a id="canonical-db64e0dbfd99b772a30404480756d919454aaf964b26b8550778d0ab79719c62"></a>

## store_provider property — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-f0880e5cbd081a2adf058ffcdb63a88547e264b69c873ae2bf2825017dfd7130"></a>

## Next pages — domains.simple_login.password.blindfold_secret_info / 3c0cfe464463 / 7

- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-3ab32d8cb66385b672c3ffbe720bc251fbf50ccf200d45e03d317d97f5d893af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26193e6f733517b07d202d5d87f052c90285b493197bdc699600f167db6e55c2"></a>

## domains.simple_login.password.clear_secret_info — domains.simple_login.password.clear_secret_info / ddc5af52688d / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-d627055a7582c05efef412a79b2ed1dfa360e2e277934617f2b77eea539baa44)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-18b74863581441bff5c54b83116973d3048a1c605f9bf473fd89fd2381951d31)
- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2)
- domains.simple_login.password.clear_secret_info

<a id="canonical-00140828af538ea9c6cf85db012aa698ccd8d669b0e12ab0f63a31e7d0d8e399"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-47abef36161605bb6fc13c099fd66fa42ae82e3239856e385e7dcbd8fe2f1c2c"></a>

## Direct properties — domains.simple_login.password.clear_secret_info / ddc5af52688d / 3

<a id="canonical-167642f0d04eab8252b7e201e613e81d1671fce69645e4f8b97207fd3227457c"></a>

<a id="canonical-6b694ad21159ae7f4435ba6c8487c51ba63631923593b76b00e29d241a2c3f3a"></a>

## provider_ref property — domains.simple_login.password.clear_secret_info / ddc5af52688d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-63d361e4ad2d7f425f7c65b76d0e2768dd65e2d95b128a5d9da4299af88819ac"></a>

<a id="canonical-0a88df739dfa5381f7645d14fc5c9e4c432c3fa60b5c03262a5b1625cd57e7f9"></a>

## url property — domains.simple_login.password.clear_secret_info / ddc5af52688d / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-a739eeb84747b25f3be8401d01476248b02c33da05b7b2e75f25749bfc59aee2"></a>

## Next pages — domains.simple_login.password.clear_secret_info / ddc5af52688d / 6

- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-528513b9857e0aeb2c4d3027dcf9e433395ccc55e93646b269278f3468a1edb2)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
