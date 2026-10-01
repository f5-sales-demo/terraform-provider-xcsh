---
page_title: "xcsh_api_crawler reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler reference."
---

# xcsh_api_crawler reference

<a id="canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57f8d2a299806d147193d9d7ca826a1b1fd4ea048fb72e242e0935dc00947d10"></a>

## Property reference — Property reference / ae2127cd2da0 / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- Property reference

<a id="canonical-c39342cb65abd803821f0cade19b05b572b9bd4cfb2c9e7949a2b6d8837a9883"></a>

## Direct properties — Property reference / ae2127cd2da0 / 3

<a id="canonical-9fcd7520589d5705f77db05aa5f9c97570c69927114b94aba98d6efe24026095"></a>

<a id="canonical-85a0f8a4983f97d72700a415047006dd3c3a1c1832de31785fb6a8c23425920a"></a>

## annotations property — Property reference / ae2127cd2da0 / 4

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

<a id="canonical-614dc86d17500901bca520a4e953bbcc1210642e349bc9d9e3fee9e7964a2f76"></a>

<a id="canonical-a016c82ed96c83f295cc9baf241a8e684cffc2521d1dcef8c33f5cf442efb4fa"></a>

## description property — Property reference / ae2127cd2da0 / 5

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

<a id="canonical-821f540aa6cf672c11a5d674cf9f0d91c2fcb67432a5ccec47e6c0507f52f569"></a>

<a id="canonical-a84aa007931f5d438faefffc4654167439543a464097ed5bac552d11c040ff71"></a>

## disable property — Property reference / ae2127cd2da0 / 6

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

- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70): complete subsection reference.

<a id="canonical-272489ddd29ed7138ed2d66aceed920a24a417a6197ec9954c5847c0cb4f035a"></a>

<a id="canonical-64b942c9181d682101e017d3ed1b65158b888d1029e4b2b1108ce333d9ccb892"></a>

## id property — Property reference / ae2127cd2da0 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-e47cfe7d72a7377e15259ff99c0a9015a5af7f0d192062582d9cd32df5ae3d26"></a>

<a id="canonical-37d98ca073b1e31782cffd504144057849ce93cb207a8eacd98feee64bdfa397"></a>

## labels property — Property reference / ae2127cd2da0 / 8

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

<a id="canonical-0c8f1ab01af48a53bc72d47dae8248bdd03477cac4fbccf873643a4d1a311e66"></a>

<a id="canonical-c5f1c36062d03b48e7d3720d3f1ab943906c57589fed3c8a11e8ca25efa41443"></a>

## name property — Property reference / ae2127cd2da0 / 9

Type: `"string"`. Required.

Name of the API Crawler. Must be unique within the namespace.

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

<a id="canonical-26ddd18eae146c8b89c2c6d2e63a5e9377559cd98e944a5c5753fafc94ccdf46"></a>

<a id="canonical-415ef1577c6e90a38a055c05dc310acc604c4a3e029681cf92480cb9875364c8"></a>

## namespace property — Property reference / ae2127cd2da0 / 10

Type: `"string"`. Required.

Namespace where the API Crawler is created.

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

- [timeouts](resources--api_crawler--reference--group-001.md#canonical-69eb8dcbcfa7b59611e1e7504d9f07b75f916ffac60640efbeef0f5630242895): complete subsection reference.

<a id="canonical-b85140a62b7ee0d33c4878749ee3a07bffa6f1f3e927d9ffdbbadbd806e46888"></a>

## All schema paths — Property reference / ae2127cd2da0 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_crawler--reference--group-001.md#canonical-9fcd7520589d5705f77db05aa5f9c97570c69927114b94aba98d6efe24026095) |
| `description` | [description](resources--api_crawler--reference--group-001.md#canonical-614dc86d17500901bca520a4e953bbcc1210642e349bc9d9e3fee9e7964a2f76) |
| `disable` | [disable](resources--api_crawler--reference--group-001.md#canonical-821f540aa6cf672c11a5d674cf9f0d91c2fcb67432a5ccec47e6c0507f52f569) |
| `domains` | [domains](resources--api_crawler--reference--group-001.md#canonical-2334026b649df3c7aa3385997b21a7a56db1f1edd74c87f0f41282f1350259ce) |
| `domains.domain` | [domains.domain](resources--api_crawler--reference--group-001.md#canonical-c44c79f43d29c584847454ab9a93a9096c7c4e5778c9ba6945795d4a1ad490f4) |
| `domains.simple_login` | [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-5bff1e398f60ff109a3488df641625da363e4d432ee3eadd6b01c1931c153d10) |
| `domains.simple_login.password` | [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-cddabc031b9f750eea995e81ac5919a9167d93eafa5019d070e33bb6f71f94ec) |
| `domains.simple_login.password.blindfold_secret_info` | [domains.simple_login.password.blindfold_secret_info](resources--api_crawler--reference--group-001.md#canonical-9b852cc2db0671442ef976539c26edc968a17b4084db216ff289ec29e321851b) |
| `domains.simple_login.password.blindfold_secret_info.decryption_provider` | [domains.simple_login.password.blindfold_secret_info.decryption_provider](resources--api_crawler--reference--group-001.md#canonical-d57a00b130913df7f51de3fe1b6d29543fbd9ed3e6ae344931382356b0cd73db) |
| `domains.simple_login.password.blindfold_secret_info.location` | [domains.simple_login.password.blindfold_secret_info.location](resources--api_crawler--reference--group-001.md#canonical-f51b5dcd2c392c2a1c8dec0768847f00f252a9a34a732128da63c984fa2f6539) |
| `domains.simple_login.password.blindfold_secret_info.store_provider` | [domains.simple_login.password.blindfold_secret_info.store_provider](resources--api_crawler--reference--group-001.md#canonical-d9369e0f9f2ab4309cf29662d8fafee5f01c37ded0ca2b714cbccf5bc5a13d28) |
| `domains.simple_login.password.clear_secret_info` | [domains.simple_login.password.clear_secret_info](resources--api_crawler--reference--group-001.md#canonical-90f112b17a5e48c190ec5377e3be0aef1109f22115cf3828d21c5137d7ec87aa) |
| `domains.simple_login.password.clear_secret_info.provider_ref` | [domains.simple_login.password.clear_secret_info.provider_ref](resources--api_crawler--reference--group-001.md#canonical-b9a748c83c0fe21d9211cab3b0e24bf370c8d2251b2c86ffcd6319321b8a16e9) |
| `domains.simple_login.password.clear_secret_info.url` | [domains.simple_login.password.clear_secret_info.url](resources--api_crawler--reference--group-001.md#canonical-ccdb115a1789fa65addb6213732ece5b7c9ddd8d84dbf9acfb625cc9622a41e6) |
| `domains.simple_login.user` | [domains.simple_login.user](resources--api_crawler--reference--group-001.md#canonical-62f03181757f97cc45e30846124211ec88123966ae5d9ae9b4db63625a886a18) |
| `id` | [id](resources--api_crawler--reference--group-001.md#canonical-272489ddd29ed7138ed2d66aceed920a24a417a6197ec9954c5847c0cb4f035a) |
| `labels` | [labels](resources--api_crawler--reference--group-001.md#canonical-e47cfe7d72a7377e15259ff99c0a9015a5af7f0d192062582d9cd32df5ae3d26) |
| `name` | [name](resources--api_crawler--reference--group-001.md#canonical-0c8f1ab01af48a53bc72d47dae8248bdd03477cac4fbccf873643a4d1a311e66) |
| `namespace` | [namespace](resources--api_crawler--reference--group-001.md#canonical-26ddd18eae146c8b89c2c6d2e63a5e9377559cd98e944a5c5753fafc94ccdf46) |
| `timeouts` | [timeouts](resources--api_crawler--reference--group-001.md#canonical-fabd4c1948459389081727d8cd3e6ffc71ed2abde9578fd6a8f1d060e5d90c83) |
| `timeouts.create` | [timeouts.create](resources--api_crawler--reference--group-001.md#canonical-340e98014cdd4ad1b2a0734a7bd908177272a2e9fd88d82c4d2203c4b4597ecb) |
| `timeouts.delete` | [timeouts.delete](resources--api_crawler--reference--group-001.md#canonical-a0e2ea55d362f50e7fde7290a742e254b5afea2c1a9500dd868aaec705571ebd) |
| `timeouts.read` | [timeouts.read](resources--api_crawler--reference--group-001.md#canonical-63670bf0d38699fab124cb325310461233cd834425a826e5166959f5eb11ad26) |
| `timeouts.update` | [timeouts.update](resources--api_crawler--reference--group-001.md#canonical-d3a97e1f38ade5cb084239f2db028022c133661abd44034df6356174246de790) |

<a id="canonical-a6222720dce6357e8d4bda10bda36ac2b18d0fb19f699e695d8716a4a4e104c7"></a>

## Next pages — Property reference / ae2127cd2da0 / 12

- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- [timeouts](resources--api_crawler--reference--group-001.md#canonical-69eb8dcbcfa7b59611e1e7504d9f07b75f916ffac60640efbeef0f5630242895)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-034200d12602882d41c2670b16a8f57ddb8aa3e271ace77ed9cc913d1f13117d"></a>

## domains — domains / 6f8f4cadfe18 / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- domains

<a id="canonical-2334026b649df3c7aa3385997b21a7a56db1f1edd74c87f0f41282f1350259ce"></a>

Type: `"object"`. list nested block, Optional.

API Crawler. API Crawler Configuration.

Upstream description:

API Crawler Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

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

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-c09845e7212bb8bca6c592376de3a1bbf3a9720fb2e512f032f150e2491e4f22"></a>

## Direct properties — domains / 6f8f4cadfe18 / 3

<a id="canonical-c44c79f43d29c584847454ab9a93a9096c7c4e5778c9ba6945795d4a1ad490f4"></a>

<a id="canonical-3316d9051d876b193e9fe48df59b67c3dc2a2c1fb0c261c108d0f988692f0c7b"></a>

## domain property — domains / 6f8f4cadfe18 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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

- [simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87): complete subsection reference.

<a id="canonical-43f3b053cdbb8b5b3c85a483f46b69bad24a865976d7a632feca35ef467e28b0"></a>

## Next pages — domains / 6f8f4cadfe18 / 5

- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c3d946bb59bed9b7a164f49ad50c9acc2fb270e0376fcd8c573c10048e5b5a5"></a>

## domains.simple_login — domains.simple_login / 86f8531eddf4 / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- domains.simple_login

<a id="canonical-5bff1e398f60ff109a3488df641625da363e4d432ee3eadd6b01c1931c153d10"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a402091cd039bdddc4fe2a009f4aa7cc954d1e7ac29e1a6589d2543d2e47ec1"></a>

## Direct properties — domains.simple_login / 86f8531eddf4 / 3

- [password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995): complete subsection reference.

<a id="canonical-62f03181757f97cc45e30846124211ec88123966ae5d9ae9b4db63625a886a18"></a>

<a id="canonical-a48b2fdb770474e5e31bdfa0dc5bea5b6622f7498568236c54b0a58bde5bdc0e"></a>

## user property — domains.simple_login / 86f8531eddf4 / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-613ea736f67503b3f80aa4e36c0041a14a5f51af0c0e31a682af27c773ae62a9"></a>

## Next pages — domains.simple_login / 86f8531eddf4 / 5

- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995)
- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8ef2969935f5630ecbf44cbbe928ca96da6c352aa8f3ae653ac6876d32ee833"></a>

## domains.simple_login.password — domains.simple_login.password / 9d7ab63481fe / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87)
- domains.simple_login.password

<a id="canonical-cddabc031b9f750eea995e81ac5919a9167d93eafa5019d070e33bb6f71f94ec"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2bdd6844354938fc9a31694cac44ad1e0ff524c1caa8f79583cf1bf69308a627"></a>

## Direct properties — domains.simple_login.password / 9d7ab63481fe / 3

- [blindfold_secret_info](resources--api_crawler--reference--group-001.md#canonical-8c6552b990825679088e5ff93c0ae3feccfe0677021b63d1e420f39f778fb62d): complete subsection reference.

- [clear_secret_info](resources--api_crawler--reference--group-001.md#canonical-99ea22a0b729cb18c071d92cfb0dc17ebae32676b6b64e6c51289bb5b4898fab): complete subsection reference.

<a id="canonical-17a11a612a6804c0c73563f486c83663c4e5c80623fcf0f4633a8ae31da2e7c6"></a>

## Next pages — domains.simple_login.password / 9d7ab63481fe / 4

- [domains.simple_login.password.blindfold_secret_info](resources--api_crawler--reference--group-001.md#canonical-8c6552b990825679088e5ff93c0ae3feccfe0677021b63d1e420f39f778fb62d)
- [domains.simple_login.password.clear_secret_info](resources--api_crawler--reference--group-001.md#canonical-99ea22a0b729cb18c071d92cfb0dc17ebae32676b6b64e6c51289bb5b4898fab)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-8c6552b990825679088e5ff93c0ae3feccfe0677021b63d1e420f39f778fb62d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9778f2eec5e077ab1fb99f15786d4aee81c1ff65e9bbd7044b26b360e50c0314"></a>

## domains.simple_login.password.blindfold_secret_info — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87)
- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995)
- domains.simple_login.password.blindfold_secret_info

<a id="canonical-9b852cc2db0671442ef976539c26edc968a17b4084db216ff289ec29e321851b"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-ddb6778142294738d31eed29f8ec6f9234df6f6718ed9ba3cd40ae6a39f30b1f"></a>

## Direct properties — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 3

<a id="canonical-d57a00b130913df7f51de3fe1b6d29543fbd9ed3e6ae344931382356b0cd73db"></a>

<a id="canonical-c451e5546c095033febebfb9b9c8fe3200a218ef5e760da4027df3db58d32dbc"></a>

## decryption_provider property — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 4

Type: `"string"`. Optional.

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

<a id="canonical-f51b5dcd2c392c2a1c8dec0768847f00f252a9a34a732128da63c984fa2f6539"></a>

<a id="canonical-e46dca9fcd8685e00a58a7084bb675384a051176d34c384cf1f02e8a04d2bbb1"></a>

## location property — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-d9369e0f9f2ab4309cf29662d8fafee5f01c37ded0ca2b714cbccf5bc5a13d28"></a>

<a id="canonical-ff15000dba24e0ba824c82b32e2644b3da2a8c300d99ddbb79c8b51bfb3a14d9"></a>

## store_provider property — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 6

Type: `"string"`. Optional.

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

<a id="canonical-2c068110f5e1500c037979154704d91a4cd9bf146caf6af27d710dcc33c8bec7"></a>

## Next pages — domains.simple_login.password.blindfold_secret_info / d7b5de49e62c / 7

- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-99ea22a0b729cb18c071d92cfb0dc17ebae32676b6b64e6c51289bb5b4898fab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce8cdbc36ab630fbac6ddc1765bc2837d2fee899491c8f39eb30b772c92bda99"></a>

## domains.simple_login.password.clear_secret_info — domains.simple_login.password.clear_secret_info / 13837237fc9e / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [domains](resources--api_crawler--reference--group-001.md#canonical-8b8a0ca0e843202153144fb3ac85a171948a0d322f9866bd6d7185887572cb70)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-263da810ad3a7ab7fbf196076f01da1e241b2d627c3d2e9559fb1b6d3b484a87)
- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995)
- domains.simple_login.password.clear_secret_info

<a id="canonical-90f112b17a5e48c190ec5377e3be0aef1109f22115cf3828d21c5137d7ec87aa"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0ca90730fb0a0c3cea4fcb35cabd435469882356d70aeebe4e944fdc63a659d"></a>

## Direct properties — domains.simple_login.password.clear_secret_info / 13837237fc9e / 3

<a id="canonical-b9a748c83c0fe21d9211cab3b0e24bf370c8d2251b2c86ffcd6319321b8a16e9"></a>

<a id="canonical-dd67167e1299bcc34d8a173d7b5896dca1e68dc62e68b4e23f4d3370157a0022"></a>

## provider_ref property — domains.simple_login.password.clear_secret_info / 13837237fc9e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ccdb115a1789fa65addb6213732ece5b7c9ddd8d84dbf9acfb625cc9622a41e6"></a>

<a id="canonical-a8789f1089958405cb530055ebfda6159cdd06d312376e15c7e122b48d81c82f"></a>

## url property — domains.simple_login.password.clear_secret_info / 13837237fc9e / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-f02a885491922da2800b577f787b799cedb0cea7d875f1528450becfed63b3a1"></a>

## Next pages — domains.simple_login.password.clear_secret_info / 13837237fc9e / 6

- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-9c6d6de11abc2ff536663083cc94bd6755166ec78e22e15d7c906a61e630c995)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-69eb8dcbcfa7b59611e1e7504d9f07b75f916ffac60640efbeef0f5630242895"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-452aadc72b431b9aa9ee3b1f19d4ec3f0b9a3ba596bdb3675c084db94d5ff657"></a>

## timeouts — timeouts / 6f212d1f7bdc / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- timeouts

<a id="canonical-fabd4c1948459389081727d8cd3e6ffc71ed2abde9578fd6a8f1d060e5d90c83"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3aa4e70aeacbabe64e8577c181f96312414e339fa294639fd1f8772fb9ea27b8"></a>

## Direct properties — timeouts / 6f212d1f7bdc / 3

<a id="canonical-340e98014cdd4ad1b2a0734a7bd908177272a2e9fd88d82c4d2203c4b4597ecb"></a>

<a id="canonical-5c5de113b6d27502f970de82e6804903db73418ca09a88de81e3a584a96dc462"></a>

## create property — timeouts / 6f212d1f7bdc / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a0e2ea55d362f50e7fde7290a742e254b5afea2c1a9500dd868aaec705571ebd"></a>

<a id="canonical-25ccfd9952627a39439e0a4506466352b7f45e48684b01b26665a850742f7f83"></a>

## delete property — timeouts / 6f212d1f7bdc / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-63670bf0d38699fab124cb325310461233cd834425a826e5166959f5eb11ad26"></a>

<a id="canonical-541e0859dcf3489e9c2b0f0f078a3abf59c8f21504c695dafec7305bca37d2e3"></a>

## read property — timeouts / 6f212d1f7bdc / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d3a97e1f38ade5cb084239f2db028022c133661abd44034df6356174246de790"></a>

<a id="canonical-af696144dfa2f65a425499e11e1c62fd52cb2aec5115cc808f2c97f40da19cca"></a>

## update property — timeouts / 6f212d1f7bdc / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-361821afc6f99fcdd821f787d51430ffe96ffa41041017f48f23c236cd5dcc42"></a>

## Next pages — timeouts / 6f212d1f7bdc / 8

- [Property reference](resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
