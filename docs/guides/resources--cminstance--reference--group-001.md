---
page_title: "xcsh_cminstance reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance reference."
---

# xcsh_cminstance reference

<a id="canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-543cc122df6582f79aa49e222682c21ce49b7210da734bc16b02470ed15c2d55"></a>

## Property reference — Property reference / 175ad1f18c29 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- Property reference

<a id="canonical-2e12dcab53fff7548071c318cc5080bb00872822105af799cbac89a188d58d96"></a>

## Direct properties — Property reference / 175ad1f18c29 / 3

<a id="canonical-c2a38e6115b7f4c09046d8b5b672038b8b8ff0534a87efa1f83a65270e86a32e"></a>

<a id="canonical-f7f483a060cc94636ee9d21c1088e6c00f256a777d5a4f87b01ade49a47ed925"></a>

## annotations property — Property reference / 175ad1f18c29 / 4

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

- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456): complete subsection reference.

<a id="canonical-fc46c0eb6c25dc443da1b0ab31a914b641d6ab8f98c80547af96e24c32edd521"></a>

<a id="canonical-d6b6976bb0cae694f0d288be68214484896d2e729f66ad2da49a46f54c205469"></a>

## description property — Property reference / 175ad1f18c29 / 5

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

<a id="canonical-f585b572b1d89e7b2ae5c45d1211eca08e0aefd25960fcb84018c89ec833e671"></a>

<a id="canonical-671bb0fd10e3f737f3317c1072df73a14dea2ac48debbfb212557829a0a2ad60"></a>

## disable property — Property reference / 175ad1f18c29 / 6

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

<a id="canonical-9f2d3d8fb132e4ad575cb04ec5d284b7ecb816a20178ac2be89f9cb9a85ae13e"></a>

<a id="canonical-3e4e504322ba8d00f00a9bbc8555cf30c075237537bd7c8d0de3045a48405524"></a>

## id property — Property reference / 175ad1f18c29 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip](resources--cminstance--reference--group-001.md#canonical-4208ac3bdeba809c65b0e7a640952ddc9421205ab8ed9d36e182f4a0058ed12d): complete subsection reference.

<a id="canonical-af7b6f52710d56972d0e581928f9c4a107b5e69f2826584969982f95e3d5f56d"></a>

<a id="canonical-2861db7e9ee27b6b6de7f0a4e2a84156abea907e61ea9847c9b33b675088def0"></a>

## labels property — Property reference / 175ad1f18c29 / 8

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

<a id="canonical-03c529df917e6b60ac293b06ab08cfce60fc654a5075ea18447d36746ff07bd5"></a>

<a id="canonical-4fe42c3f49910ff907769df35e6894f71cdea6fc4e052791bac6842d574ef604"></a>

## name property — Property reference / 175ad1f18c29 / 9

Type: `"string"`. Required.

Name of the Cminstance. Must be unique within the namespace.

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

<a id="canonical-52d791175bc750c5c5b283dffb3a53229417c65b2c8f04687124807756a93bc8"></a>

<a id="canonical-b13f23dde29a30d90213bc0e8b9082788638f4922b192535ff907b7f056e524b"></a>

## namespace property — Property reference / 175ad1f18c29 / 10

Type: `"string"`. Required.

Namespace where the Cminstance is created.

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

- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558): complete subsection reference.

<a id="canonical-a86058f1041f5e1abfdd55d16aaaf59f0efc5185257dec349153d1d310b6ed8e"></a>

<a id="canonical-3338dc57a34020c6eddaf5d2c4653027542d75f000fed5625050e1d6472f277a"></a>

## port property — Property reference / 175ad1f18c29 / 11

Type: `"number"`. Required.

Port of the Central Manager instance to connect to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [timeouts](resources--cminstance--reference--group-001.md#canonical-c618a58cf4b317ee41d5f985fe5bdfe9ce0372370876b4438635ccc4ccfaa964): complete subsection reference.

<a id="canonical-e84dae64c91d6a92bd7a1e29e482be51f695b83597c19e267dcc66c2b1a5eb11"></a>

<a id="canonical-a585ec388ebea9550b5ce4a4bb48cfc16ea94716e7b32617286f08bc13ab7b1d"></a>

## username property — Property reference / 175ad1f18c29 / 12

Type: `"string"`. Required.

Username for the Central Manager instance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-99617af7a7f48151241d2bb3f27c3801206e9b5e1818af784742139af2e24dbf"></a>

## All schema paths — Property reference / 175ad1f18c29 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cminstance--reference--group-001.md#canonical-c2a38e6115b7f4c09046d8b5b672038b8b8ff0534a87efa1f83a65270e86a32e) |
| `api_token` | [api_token](resources--cminstance--reference--group-001.md#canonical-77b0aa76b8474fd9cb5f7b3618fc3d307d83f01f201834a54cb20d9b9339cab5) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-9d1f311dd6d60f48f476142a742a343b89dc71e93247f44e662d469382f5cffe) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](resources--cminstance--reference--group-001.md#canonical-beac995e93143ab7b1b20644cf27dd9d7b4e175bc98615b95039dd35561017e7) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](resources--cminstance--reference--group-001.md#canonical-02890753eac524b44cc431ee902e1163554c37dca4e89800cae65a45ef6df84e) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](resources--cminstance--reference--group-001.md#canonical-488765200506e5715bd945bf3772c800fa2c764a7a7cd8f70f619ab2f180a7f8) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-09b08b75e784d33fea7799b59168e74ef5b378633e3553a42ec337d7fe5dc93b) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](resources--cminstance--reference--group-001.md#canonical-230aba866b8723e34a11cb0964ad92d7371b740dc8950dde9121f9dba44d9348) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](resources--cminstance--reference--group-001.md#canonical-d3dd999f6e730aed8e752f3d99c0f57c3d4a095997f30f30c18469513a1471b5) |
| `description` | [description](resources--cminstance--reference--group-001.md#canonical-fc46c0eb6c25dc443da1b0ab31a914b641d6ab8f98c80547af96e24c32edd521) |
| `disable` | [disable](resources--cminstance--reference--group-001.md#canonical-f585b572b1d89e7b2ae5c45d1211eca08e0aefd25960fcb84018c89ec833e671) |
| `id` | [id](resources--cminstance--reference--group-001.md#canonical-9f2d3d8fb132e4ad575cb04ec5d284b7ecb816a20178ac2be89f9cb9a85ae13e) |
| `ip` | [ip](resources--cminstance--reference--group-001.md#canonical-20ef244fc0b6e21ba84c0848f04a4c2536c760657478317c272113db2565c5a8) |
| `ip.addr` | [ip.addr](resources--cminstance--reference--group-001.md#canonical-7d6e9b230fce7f6a103025074fb899e870898086ed4f918077c9c4aee11bc401) |
| `labels` | [labels](resources--cminstance--reference--group-001.md#canonical-af7b6f52710d56972d0e581928f9c4a107b5e69f2826584969982f95e3d5f56d) |
| `name` | [name](resources--cminstance--reference--group-001.md#canonical-03c529df917e6b60ac293b06ab08cfce60fc654a5075ea18447d36746ff07bd5) |
| `namespace` | [namespace](resources--cminstance--reference--group-001.md#canonical-52d791175bc750c5c5b283dffb3a53229417c65b2c8f04687124807756a93bc8) |
| `password` | [password](resources--cminstance--reference--group-001.md#canonical-7d36e31e30fba19cfd43be27e85d6a9be99b2ac0e688726ffcf043b1765d071f) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-b0a3abde8ca62d8ebace3e5e3179b29b2745462f6db19cfc5a0bd3ca72b43a08) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](resources--cminstance--reference--group-001.md#canonical-5e356a9804a4853c40d4880947ec27624dc735d4f2c71e5f1fcd6fe764d5d499) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](resources--cminstance--reference--group-001.md#canonical-3b50f8a553adcb50263876bc866badacaad85a65c46c81d804924eaed2549b7a) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](resources--cminstance--reference--group-001.md#canonical-69cb8d527a0e7af7b59eaa8b591ce466e0ec7d11a6d7575bfd6ab8231f8968e7) |
| `password.clear_secret_info` | [password.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-2635f65ddf3f9d7df5423735d6bad08bada0bae385f483a922a01f06b4a5b6d7) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](resources--cminstance--reference--group-001.md#canonical-dafc105a9e4b2a71acbfff8e524177c6dbc1882aa7875ba0c1a9482ca1992749) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](resources--cminstance--reference--group-001.md#canonical-7456c13c870691f7f9fab4093eef92cdef1e890d4c055561689a9b970586d783) |
| `port` | [port](resources--cminstance--reference--group-001.md#canonical-a86058f1041f5e1abfdd55d16aaaf59f0efc5185257dec349153d1d310b6ed8e) |
| `timeouts` | [timeouts](resources--cminstance--reference--group-001.md#canonical-4dc34b93a6fbec662d7961402d9c964d83ee2205ddd51464a2d5302645cdb24d) |
| `timeouts.create` | [timeouts.create](resources--cminstance--reference--group-001.md#canonical-fc26f78eef477c69dd27a65193faf0203fa06c9dbfd178b713120978c642546a) |
| `timeouts.delete` | [timeouts.delete](resources--cminstance--reference--group-001.md#canonical-8590a2ffb0e346f95b9ff8ca74afc5f4cbcc7804ee5bfe23fe59dcd5641067e9) |
| `timeouts.read` | [timeouts.read](resources--cminstance--reference--group-001.md#canonical-9d004d8c8b5e2f9a7e0a2644f0b4bf3005a54bee5cd4659e79ff3a21069d946b) |
| `timeouts.update` | [timeouts.update](resources--cminstance--reference--group-001.md#canonical-9c1c0714cc7a459df497d5295dd5743bad0e62f742168af82f8d751bd4e5fa19) |
| `username` | [username](resources--cminstance--reference--group-001.md#canonical-e84dae64c91d6a92bd7a1e29e482be51f695b83597c19e267dcc66c2b1a5eb11) |

<a id="canonical-83258f1babad9ffbe14b868070f179f0209cda9f3c75ed7ff4ec6c2e4bb4334b"></a>

## Next pages — Property reference / 175ad1f18c29 / 14

- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456)
- [ip](resources--cminstance--reference--group-001.md#canonical-4208ac3bdeba809c65b0e7a640952ddc9421205ab8ed9d36e182f4a0058ed12d)
- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558)
- [timeouts](resources--cminstance--reference--group-001.md#canonical-c618a58cf4b317ee41d5f985fe5bdfe9ce0372370876b4438635ccc4ccfaa964)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-775f22586a091dc8af635ff58211bb63e529c6d1aae0ad0b097f2f08001ffedb"></a>

## api_token — api_token / 0eeeae0361f2 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- api_token

<a id="canonical-77b0aa76b8474fd9cb5f7b3618fc3d307d83f01f201834a54cb20d9b9339cab5"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-8db5583e1bf95457e2602942b5231d1f91274b5f2533dd5d332ffed1386f257c"></a>

## Direct properties — api_token / 0eeeae0361f2 / 3

- [blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-1a7cd43c28c3c30ea645c68f2d48f5c417919ec689f4bc7feacab4c63aa0baeb): complete subsection reference.

- [clear_secret_info](resources--cminstance--reference--group-001.md#canonical-6fa3aa1a288a2469f0662a03ba2be55ab9e77f73819001633643c2d2f04783a2): complete subsection reference.

<a id="canonical-871fe706f25ca0be2db776f8ffec05f85eb4488beae3aad5437159e3ccaeb2a3"></a>

## Next pages — api_token / 0eeeae0361f2 / 4

- [api_token.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-1a7cd43c28c3c30ea645c68f2d48f5c417919ec689f4bc7feacab4c63aa0baeb)
- [api_token.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-6fa3aa1a288a2469f0662a03ba2be55ab9e77f73819001633643c2d2f04783a2)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-1a7cd43c28c3c30ea645c68f2d48f5c417919ec689f4bc7feacab4c63aa0baeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c36923103780bd49514e1143e672d186763db2a464e1f2d40e2e8856d27e7642"></a>

## api_token.blindfold_secret_info — api_token.blindfold_secret_info / ff9644f88122 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456)
- api_token.blindfold_secret_info

<a id="canonical-9d1f311dd6d60f48f476142a742a343b89dc71e93247f44e662d469382f5cffe"></a>

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

<a id="canonical-f2baab1e27f2edec3dc580311bae866cdd88a1138cf7281b94dfe1acdf283af4"></a>

## Direct properties — api_token.blindfold_secret_info / ff9644f88122 / 3

<a id="canonical-beac995e93143ab7b1b20644cf27dd9d7b4e175bc98615b95039dd35561017e7"></a>

<a id="canonical-b2eeb43f370196d19f1fd20f58a74cf5507dd034402f1db171e0a84e582c7ad3"></a>

## decryption_provider property — api_token.blindfold_secret_info / ff9644f88122 / 4

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

<a id="canonical-02890753eac524b44cc431ee902e1163554c37dca4e89800cae65a45ef6df84e"></a>

<a id="canonical-b6b564b46a5ce2593facce416ec6d004f54b39fd421e0ffd18881cd414b2260c"></a>

## location property — api_token.blindfold_secret_info / ff9644f88122 / 5

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

<a id="canonical-488765200506e5715bd945bf3772c800fa2c764a7a7cd8f70f619ab2f180a7f8"></a>

<a id="canonical-5bd72bfc892018df2aa1e23461eab31899e4a10fd36eb335d1d353ba6426fe31"></a>

## store_provider property — api_token.blindfold_secret_info / ff9644f88122 / 6

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

<a id="canonical-740c925c23abe8f2e510315939fdf3b96cce7fcdcdabff79d13bb3de7a343eda"></a>

## Next pages — api_token.blindfold_secret_info / ff9644f88122 / 7

- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-6fa3aa1a288a2469f0662a03ba2be55ab9e77f73819001633643c2d2f04783a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94017c83032bbc5ec54fcb5a6bca77d44a14dfd221e9283725ee3ea0db6b2507"></a>

## api_token.clear_secret_info — api_token.clear_secret_info / 36b8db4b5332 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456)
- api_token.clear_secret_info

<a id="canonical-09b08b75e784d33fea7799b59168e74ef5b378633e3553a42ec337d7fe5dc93b"></a>

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

<a id="canonical-b2161a6764b6e5159007b468f6ce8924ec1ebd2d75ad7d4bf1cb78d11d9a217b"></a>

## Direct properties — api_token.clear_secret_info / 36b8db4b5332 / 3

<a id="canonical-230aba866b8723e34a11cb0964ad92d7371b740dc8950dde9121f9dba44d9348"></a>

<a id="canonical-686c6925d090a0ca3613cb6f5a5880c1b7200d253b186d1e1f7bd3b5a49644b0"></a>

## provider_ref property — api_token.clear_secret_info / 36b8db4b5332 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d3dd999f6e730aed8e752f3d99c0f57c3d4a095997f30f30c18469513a1471b5"></a>

<a id="canonical-114929a1881cdd6f9a50805b19e1186e2b7237e78d904807a213ea58e843e25d"></a>

## url property — api_token.clear_secret_info / 36b8db4b5332 / 5

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

<a id="canonical-0441474692134043b54098851735a166fffaa52a43e97d0a7733a013fc35e866"></a>

## Next pages — api_token.clear_secret_info / 36b8db4b5332 / 6

- [api_token](resources--cminstance--reference--group-001.md#canonical-ea33742e2022969dd31e7c6fb3bd2bc25dbe6edf5650bb9b90ed1c0561324456)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-4208ac3bdeba809c65b0e7a640952ddc9421205ab8ed9d36e182f4a0058ed12d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5220dadf29c95882caf355910b0d9dcee58a200067b3dbee65787b4c05b6bcb7"></a>

## ip — ip / 23c11a5887d9 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- ip

<a id="canonical-20ef244fc0b6e21ba84c0848f04a4c2536c760657478317c272113db2565c5a8"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a5712dd4a1ac7317bdc6eab6dbeb50fab34ae1f718d768d36a5331de27fc8b9"></a>

## Direct properties — ip / 23c11a5887d9 / 3

<a id="canonical-7d6e9b230fce7f6a103025074fb899e870898086ed4f918077c9c4aee11bc401"></a>

<a id="canonical-b043b55b92f3be2cad6ba7408f2421ca1d1fa946c098aa46ef1eb754c11ad5e8"></a>

## addr property — ip / 23c11a5887d9 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-222bb322e5ac03d84d024ce0293af0d4c10c66240fad9244ce03dca9f53a1602"></a>

## Next pages — ip / 23c11a5887d9 / 5

- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c414c407ba76325a2e5dbb065b4ee8de6e69eb31d073a681060bce0c4c35978"></a>

## password — password / cb56e46f1e40 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- password

<a id="canonical-7d36e31e30fba19cfd43be27e85d6a9be99b2ac0e688726ffcf043b1765d071f"></a>

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

<a id="canonical-15f3ac0c104db909fca693e8d0ea2420f1a18d1ef6a2425a0b44792f8c4f290c"></a>

## Direct properties — password / cb56e46f1e40 / 3

- [blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-757113d7b1afc0f2b052c1ed36509e5c6b9b3a9860686bbe4e8e4608aaf53c81): complete subsection reference.

- [clear_secret_info](resources--cminstance--reference--group-001.md#canonical-5675a1e009f8431b7727bf8bb0fcf0d15b277e13a69dc5646ae60f8fbbbf1595): complete subsection reference.

<a id="canonical-003b382fdd4621dabc3c965fb069b8c3cdc632584a515c1af651dee921a7671e"></a>

## Next pages — password / cb56e46f1e40 / 4

- [password.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-757113d7b1afc0f2b052c1ed36509e5c6b9b3a9860686bbe4e8e4608aaf53c81)
- [password.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-5675a1e009f8431b7727bf8bb0fcf0d15b277e13a69dc5646ae60f8fbbbf1595)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-757113d7b1afc0f2b052c1ed36509e5c6b9b3a9860686bbe4e8e4608aaf53c81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9a96eee2fbde4df06266f6d7ddd455f1f70c637f10f4b940bd1d2c458a8d370"></a>

## password.blindfold_secret_info — password.blindfold_secret_info / 091a2160df52 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558)
- password.blindfold_secret_info

<a id="canonical-b0a3abde8ca62d8ebace3e5e3179b29b2745462f6db19cfc5a0bd3ca72b43a08"></a>

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

<a id="canonical-2d749ac57ff55ba00bfc4dc1ad83673f283fbd878bb31113cee7da9bb31a505e"></a>

## Direct properties — password.blindfold_secret_info / 091a2160df52 / 3

<a id="canonical-5e356a9804a4853c40d4880947ec27624dc735d4f2c71e5f1fcd6fe764d5d499"></a>

<a id="canonical-f81e99eccd913ae467475279e39df994e1c10ae14370c72222ab84aff536d329"></a>

## decryption_provider property — password.blindfold_secret_info / 091a2160df52 / 4

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

<a id="canonical-3b50f8a553adcb50263876bc866badacaad85a65c46c81d804924eaed2549b7a"></a>

<a id="canonical-51213fa5b3545ec563f1c9e85f2b2723639ae94998b5aca83e5180c10631be7a"></a>

## location property — password.blindfold_secret_info / 091a2160df52 / 5

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

<a id="canonical-69cb8d527a0e7af7b59eaa8b591ce466e0ec7d11a6d7575bfd6ab8231f8968e7"></a>

<a id="canonical-4137b593eee18a193284bd33b7c0a1290adcc96c743adf57f5f392aebf5acf98"></a>

## store_provider property — password.blindfold_secret_info / 091a2160df52 / 6

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

<a id="canonical-eb3cb76111f571da82869f6722aae408d28bff2bf6f4bf0caeefddac8e5b8281"></a>

## Next pages — password.blindfold_secret_info / 091a2160df52 / 7

- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-5675a1e009f8431b7727bf8bb0fcf0d15b277e13a69dc5646ae60f8fbbbf1595"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f23685a9e9347e23eefd6969fabfe4472f79594ee7e75f2a9b25bf07ed388723"></a>

## password.clear_secret_info — password.clear_secret_info / 210b528dfa2d / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558)
- password.clear_secret_info

<a id="canonical-2635f65ddf3f9d7df5423735d6bad08bada0bae385f483a922a01f06b4a5b6d7"></a>

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

<a id="canonical-85cc1eff6ac2d0e6481ddef3918ad70998d9f6cc125f53abd941442b06876a4c"></a>

## Direct properties — password.clear_secret_info / 210b528dfa2d / 3

<a id="canonical-dafc105a9e4b2a71acbfff8e524177c6dbc1882aa7875ba0c1a9482ca1992749"></a>

<a id="canonical-b80e0423cdc1f124b4e3d3306fd3a985d215e85f605437c9b05de20e695853cb"></a>

## provider_ref property — password.clear_secret_info / 210b528dfa2d / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7456c13c870691f7f9fab4093eef92cdef1e890d4c055561689a9b970586d783"></a>

<a id="canonical-a5d6aa7591a34e3bd4fcbc5a15b949e592721d8d6d4080e33e3c9b8762c4d8f4"></a>

## url property — password.clear_secret_info / 210b528dfa2d / 5

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

<a id="canonical-353ec1952ceb8b0f856f40efb8b843556e825c352c1b8b300235bf5a33af4917"></a>

## Next pages — password.clear_secret_info / 210b528dfa2d / 6

- [password](resources--cminstance--reference--group-001.md#canonical-9f344e7c141b5de8499fe4be4e8555ec1be72b5b1c3adef5c3d59012b93a7558)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-c618a58cf4b317ee41d5f985fe5bdfe9ce0372370876b4438635ccc4ccfaa964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4afa640164ee2ffb4f7d9b44db5dd1c785cd6da51459b3b62b6aca3b3821a7c9"></a>

## timeouts — timeouts / 3be3dc4acc96 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- timeouts

<a id="canonical-4dc34b93a6fbec662d7961402d9c964d83ee2205ddd51464a2d5302645cdb24d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0cd2ebd91d9e002bd07435fdf7c4857b394d8821823d3f3072bb0e1c21e95c89"></a>

## Direct properties — timeouts / 3be3dc4acc96 / 3

<a id="canonical-fc26f78eef477c69dd27a65193faf0203fa06c9dbfd178b713120978c642546a"></a>

<a id="canonical-dd03cc6d354a68c7503bc2d467cba8e73c95bcf6ec1f6f4bdb86d25f2ff58689"></a>

## create property — timeouts / 3be3dc4acc96 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8590a2ffb0e346f95b9ff8ca74afc5f4cbcc7804ee5bfe23fe59dcd5641067e9"></a>

<a id="canonical-a86ce1be927437a2bc745a67524ac4b6e63caa8403a2411f4176653b11514458"></a>

## delete property — timeouts / 3be3dc4acc96 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-9d004d8c8b5e2f9a7e0a2644f0b4bf3005a54bee5cd4659e79ff3a21069d946b"></a>

<a id="canonical-a4b0866f35177af295a0fde0c77581d2ed437452b3c3d94c7acd8215dacd984c"></a>

## read property — timeouts / 3be3dc4acc96 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9c1c0714cc7a459df497d5295dd5743bad0e62f742168af82f8d751bd4e5fa19"></a>

<a id="canonical-16b8948abbe69773ee576016cf166b241702d9b6bd4ed0feeacbaa9c0e351b61"></a>

## update property — timeouts / 3be3dc4acc96 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3d6d5e181d5dcfdacc87e269d8d79cdf4df18c20c753e9cdacdb554e30926b12"></a>

## Next pages — timeouts / 3be3dc4acc96 / 8

- [Property reference](resources--cminstance--reference--group-001.md#canonical-1a3303bad1df452da182bfce55ae7d71e5cba24da48067fcf379e36045f7be59)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
