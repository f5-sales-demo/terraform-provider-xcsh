---
page_title: "xcsh_rate_limiter reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter reference."
---

# xcsh_rate_limiter reference

<a id="canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45f44a23e4a0ce29629b124009044572d0c2b0fb61498dfce8db141b3a69630f"></a>

## Property reference — Property reference / ed7503977601 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- Property reference

<a id="canonical-78c7a29a9dd24fcf8c8c9e03a68c14be9ccbe91ad6dea2f7acf306f95c7a6736"></a>

## Direct properties — Property reference / ed7503977601 / 3

<a id="canonical-8313c3cfde8f0c7b881b595c0530e28e57e82cf28bdab163849d1dc90c3baf83"></a>

<a id="canonical-04a91f6fe0a8604b98281f1b833a2ea6f27733a0ce4c8195f895e37a623237f6"></a>

## annotations property — Property reference / ed7503977601 / 4

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

<a id="canonical-9a8e00ce2886abbeb3b87f2a1b3c4d40042e8ad9b66190c50ca17ac5fe827fd4"></a>

<a id="canonical-83ca40ec84e6710c504285900a69b7e2d74202e479de9626fcc2c1c6b9aa572c"></a>

## description property — Property reference / ed7503977601 / 5

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

<a id="canonical-1e1fc3d4886a5b8f514019063435809f8ad5a75110d59a31eed21936f9c25c48"></a>

<a id="canonical-a0ad43a61e182c4f8eb9dd87609057134eee4a1098f11752753715f5c4ec3bb8"></a>

## disable property — Property reference / ed7503977601 / 6

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

<a id="canonical-df4a5cd99a96bef62e0a24d5272a7eca703cca6db7b6babf63ae9679f8d39393"></a>

<a id="canonical-7532cb7b3a39db455d2c5cdc01cc5ca8fa7762714b263c7762898af61b3ea8d7"></a>

## id property — Property reference / ed7503977601 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a23b93e93ec7c107580836249ccea2554859a6d272b3ae0ad38fd67e0447f5e9"></a>

<a id="canonical-9a8e9f56ad2ff2f694142bca315e4aef2fdde7140bf5dd13a9a0efedb3ad0cac"></a>

## labels property — Property reference / ed7503977601 / 8

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

- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a): complete subsection reference.

<a id="canonical-3487bae0b621c4dc3c796c071fe352b3d192f834b930c64d9100ff526e528cb2"></a>

<a id="canonical-509041611f8b22bb1ae2718557c4bcf902eea748ede985176f0718eb39d7cb50"></a>

## name property — Property reference / ed7503977601 / 9

Type: `"string"`. Required.

Name of the Rate Limiter. Must be unique within the namespace.

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

<a id="canonical-0bee3117b7d6b9133962ccf74155a78f27ae7539a0a4200ee37f00af57e20c0a"></a>

<a id="canonical-b7bad3f68d463ccf18dbba1f867c02f93ca79442ae3e69904ec111d57e69ac7d"></a>

## namespace property — Property reference / ed7503977601 / 10

Type: `"string"`. Required.

Namespace where the Rate Limiter is created.

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

- [timeouts](resources--rate_limiter--reference--group-001.md#canonical-71f67a13190659f6652797231d0bd7bcdc4a9539bcaf307aa577563a18c60677): complete subsection reference.

- [user_identification](resources--rate_limiter--reference--group-001.md#canonical-c71857479c5fdec24ae2de7db1c7d7c2fadd6c396ee88f9da33557471ea0d42b): complete subsection reference.

<a id="canonical-e7115bf6aefd7b796655e4e4bc25be4d4930cbd55adbb99a0f3230331a22ad3d"></a>

## All schema paths — Property reference / ed7503977601 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--rate_limiter--reference--group-001.md#canonical-8313c3cfde8f0c7b881b595c0530e28e57e82cf28bdab163849d1dc90c3baf83) |
| `description` | [description](resources--rate_limiter--reference--group-001.md#canonical-9a8e00ce2886abbeb3b87f2a1b3c4d40042e8ad9b66190c50ca17ac5fe827fd4) |
| `disable` | [disable](resources--rate_limiter--reference--group-001.md#canonical-1e1fc3d4886a5b8f514019063435809f8ad5a75110d59a31eed21936f9c25c48) |
| `id` | [id](resources--rate_limiter--reference--group-001.md#canonical-df4a5cd99a96bef62e0a24d5272a7eca703cca6db7b6babf63ae9679f8d39393) |
| `labels` | [labels](resources--rate_limiter--reference--group-001.md#canonical-a23b93e93ec7c107580836249ccea2554859a6d272b3ae0ad38fd67e0447f5e9) |
| `limits` | [limits](resources--rate_limiter--reference--group-001.md#canonical-b7fef236206496682885cc848067f085299aff3f883043a5c61363f5cfd68d1b) |
| `limits.action_block` | [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-0103d2e4e5bf6771f9d8d008bfefdada288122cec036c586e9735d8e465de59c) |
| `limits.action_block.hours` | [limits.action_block.hours](resources--rate_limiter--reference--group-001.md#canonical-4ce929d78af865971f520c19efe9f8d385e828514bd0b8b0d82e6c82491b14fa) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](resources--rate_limiter--reference--group-001.md#canonical-44eb79b7898e20d193049011e1ef34895be90863cc1141e80f6de08ca1f35168) |
| `limits.action_block.minutes` | [limits.action_block.minutes](resources--rate_limiter--reference--group-001.md#canonical-12c5faa791bad12bb26ec8e955d39114d0c189701e6e1dac789b891b19c3c9aa) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](resources--rate_limiter--reference--group-001.md#canonical-1a836af83eabde45767b35db2af91b3a99466c8756a46f09113fbcd206722c0e) |
| `limits.action_block.seconds` | [limits.action_block.seconds](resources--rate_limiter--reference--group-001.md#canonical-2751e8f9d13897708897de57c7ca5e057b6255e0857c70006c756818f3c6bd54) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](resources--rate_limiter--reference--group-001.md#canonical-29dc165987e03867dbd5288081cc7d67c4308f989679d14fb6217e06ec1db823) |
| `limits.burst_multiplier` | [limits.burst_multiplier](resources--rate_limiter--reference--group-001.md#canonical-98f3bcad54a3c7aad1e5b8ce37e7187cf2d7636cadb1c6d0c7c2bc5f8aa508f0) |
| `limits.disabled` | [limits.disabled](resources--rate_limiter--reference--group-001.md#canonical-6be65be12a2ee424f7ec6446f758378ceea8cfd2c56285bc3f98d2b917d6f569) |
| `limits.leaky_bucket` | [limits.leaky_bucket](resources--rate_limiter--reference--group-001.md#canonical-14f05a7aa223ea93257cef71335c8f072a8a84139ae2c483253b82822ebdf4c6) |
| `limits.period_multiplier` | [limits.period_multiplier](resources--rate_limiter--reference--group-001.md#canonical-5cb373ef2b78e4b17369e2f0e1fc872690bb7d84a81ae09c4762d0128c95053a) |
| `limits.token_bucket` | [limits.token_bucket](resources--rate_limiter--reference--group-001.md#canonical-53057ff219fac970fdd5cc864939b75379ffc55abfa5d232618a37d52a5aa8b5) |
| `limits.total_number` | [limits.total_number](resources--rate_limiter--reference--group-001.md#canonical-f3b6a2b0946ebab09f73de068b2611cadc62b1a7b7161df38131051debbb48cb) |
| `limits.unit` | [limits.unit](resources--rate_limiter--reference--group-001.md#canonical-1a7ae8f5a4dfb0e401b78f04e326931d65f80e0c2a220aab1f4e1eecc9afaef1) |
| `name` | [name](resources--rate_limiter--reference--group-001.md#canonical-3487bae0b621c4dc3c796c071fe352b3d192f834b930c64d9100ff526e528cb2) |
| `namespace` | [namespace](resources--rate_limiter--reference--group-001.md#canonical-0bee3117b7d6b9133962ccf74155a78f27ae7539a0a4200ee37f00af57e20c0a) |
| `timeouts` | [timeouts](resources--rate_limiter--reference--group-001.md#canonical-8995368b12a8264d5d73b395fee3ff1202f6fe8010cbbd055df355fd35e0c76f) |
| `timeouts.create` | [timeouts.create](resources--rate_limiter--reference--group-001.md#canonical-3307419e2c34ee9d7d4fbd04d69f8d86ac5438d54fda98200b80c65d6ac013b0) |
| `timeouts.delete` | [timeouts.delete](resources--rate_limiter--reference--group-001.md#canonical-f139a64290b180fbef4b8aa2609918e3aaec85dc369962928bc95d39339842e9) |
| `timeouts.read` | [timeouts.read](resources--rate_limiter--reference--group-001.md#canonical-4847522c7ac30e29d81ab3948641c65673a9f876d57f1dc530bb5c9bb963d006) |
| `timeouts.update` | [timeouts.update](resources--rate_limiter--reference--group-001.md#canonical-881b4e142b89f1268564a42617710074695f472ab63dc1e512e6dd7fe9a05a4a) |
| `user_identification` | [user_identification](resources--rate_limiter--reference--group-001.md#canonical-f60c6bb90a5fe22fae476e39c42aff138534dc4bab2c273599b99f0845df850a) |
| `user_identification.kind` | [user_identification.kind](resources--rate_limiter--reference--group-001.md#canonical-005a55d6d8c23467eb983d85e0b0c529f57de755465dbcf2c53d797b8f1cd7a9) |
| `user_identification.name` | [user_identification.name](resources--rate_limiter--reference--group-001.md#canonical-2234a6025d598f3aa48910f89232367bcbe94c87ba511f604dd29fef70b25b1b) |
| `user_identification.namespace` | [user_identification.namespace](resources--rate_limiter--reference--group-001.md#canonical-f0da8eb213b5c3308390c69054693ab7fde522cf31219f6f529187deb455687e) |
| `user_identification.tenant` | [user_identification.tenant](resources--rate_limiter--reference--group-001.md#canonical-f792fe7eb63a5439fa5eeb3bc383959fe8b60391693c1c32423d31932fd35d74) |
| `user_identification.uid` | [user_identification.uid](resources--rate_limiter--reference--group-001.md#canonical-61043647d1f67837b7848e6d18d07abf59098825441efd0bdf4328970bc6aa0c) |

<a id="canonical-d9ed0e0c4d12fc2c9a723e93a31cc8f72db31e59be15ab6ab97ba6739fbac82f"></a>

## Next pages — Property reference / ed7503977601 / 12

- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [timeouts](resources--rate_limiter--reference--group-001.md#canonical-71f67a13190659f6652797231d0bd7bcdc4a9539bcaf307aa577563a18c60677)
- [user_identification](resources--rate_limiter--reference--group-001.md#canonical-c71857479c5fdec24ae2de7db1c7d7c2fadd6c396ee88f9da33557471ea0d42b)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4831d1f3e06ce1f0db82757477a9c71e039d420d514782834f6be7fdd2e01ba0"></a>

## limits — limits / 7761d9d0bbfd / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- limits

<a id="canonical-b7fef236206496682885cc848067f085299aff3f883043a5c61363f5cfd68d1b"></a>

Type: `"object"`. list nested block, Optional.

List of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Upstream description:

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("total_number"),
  validators.ConflictingListObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingListObjectAttributes("leaky_bucket",
    "token_bucket")}
```

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

Terraform syntax:

```terraform
limits {
  # Configure direct properties listed below.
}
```

<a id="canonical-359f5c1706164b65fcd9c2e6c3af5c325fda9bd96d661c45a512041678460bec"></a>

## Direct properties — limits / 7761d9d0bbfd / 3

- [action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f): complete subsection reference.

<a id="canonical-98f3bcad54a3c7aad1e5b8ce37e7187cf2d7636cadb1c6d0c7c2bc5f8aa508f0"></a>

<a id="canonical-6a28674f5dfca89dca02d065823696e3fe5d3748f79606eb06bc899c1a989502"></a>

## burst_multiplier property — limits / 7761d9d0bbfd / 4

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

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

- [disabled](resources--rate_limiter--reference--group-001.md#canonical-90fde68ff5ef47582e5936112e6658c79e8175f67bce55a3d015920ead4b2cf0): complete subsection reference.

- [leaky_bucket](resources--rate_limiter--reference--group-001.md#canonical-0f915de64bbb8272e4c98c582fde7600740e976b21fcc3acf628fb3a42010caf): complete subsection reference.

<a id="canonical-5cb373ef2b78e4b17369e2f0e1fc872690bb7d84a81ae09c4762d0128c95053a"></a>

<a id="canonical-f8ddd128e12d0de72074400d63b541b77926b64f3483dbd25638d8521d906f05"></a>

## period_multiplier property — limits / 7761d9d0bbfd / 5

Type: `"number"`. Optional.

Setting, combined with Per Period units, provides a duration.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

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

- [token_bucket](resources--rate_limiter--reference--group-001.md#canonical-72b9ce499fefa810dc273ce14ef21e22f66f6167abc9a3162b92655bc59d878f): complete subsection reference.

<a id="canonical-f3b6a2b0946ebab09f73de068b2611cadc62b1a7b7161df38131051debbb48cb"></a>

<a id="canonical-f37742c1168acf279f60df74abfc2187ae0a20d2ec918c2741467b540cd8602d"></a>

## total_number property — limits / 7761d9d0bbfd / 6

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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

<a id="canonical-1a7ae8f5a4dfb0e401b78f04e326931d65f80e0c2a220aab1f4e1eecc9afaef1"></a>

<a id="canonical-808e6a530c0cfd2a45a8fd7131a1982de8e3acda9a75b7c7ea057902a25abe8f"></a>

## unit property — limits / 7761d9d0bbfd / 7

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

<a id="canonical-b8824f32a0191fa7e7ced25807b9103a3eda057cbc3799e915c5dcd997e23c65"></a>

## Next pages — limits / 7761d9d0bbfd / 8

- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- [limits.disabled](resources--rate_limiter--reference--group-001.md#canonical-90fde68ff5ef47582e5936112e6658c79e8175f67bce55a3d015920ead4b2cf0)
- [limits.leaky_bucket](resources--rate_limiter--reference--group-001.md#canonical-0f915de64bbb8272e4c98c582fde7600740e976b21fcc3acf628fb3a42010caf)
- [limits.token_bucket](resources--rate_limiter--reference--group-001.md#canonical-72b9ce499fefa810dc273ce14ef21e22f66f6167abc9a3162b92655bc59d878f)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-385963139c4164de298038eaacafc57081f661445998eee2cc39fafe750ad1bc"></a>

## limits.action_block — limits.action_block / f1e9d562e799 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- limits.action_block

<a id="canonical-0103d2e4e5bf6771f9d8d008bfefdada288122cec036c586e9735d8e465de59c"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-a9d908581e4277c2ca421ef6781bdf8e14c35a7f4b30231a5c0acaffd5961dca"></a>

## Direct properties — limits.action_block / f1e9d562e799 / 3

- [hours](resources--rate_limiter--reference--group-001.md#canonical-b72227b0b1faa3fc2e08fc596d8c94132bb83d3fc7e7715d3b9097dc16d82c51): complete subsection reference.

- [minutes](resources--rate_limiter--reference--group-001.md#canonical-06476eecefd3b64453e69b0df8b517f59809dad1ec3d0fc7ca5866d793b55f90): complete subsection reference.

- [seconds](resources--rate_limiter--reference--group-001.md#canonical-059ac76ecb6d941646ff731dd7d283117053c5c113d1868d1b58ec4d70c2dc5e): complete subsection reference.

<a id="canonical-f674231408156b45587ae3d8fad181a93e1e66543b768db31c02d8852ef6b2f8"></a>

## Next pages — limits.action_block / f1e9d562e799 / 4

- [limits.action_block.hours](resources--rate_limiter--reference--group-001.md#canonical-b72227b0b1faa3fc2e08fc596d8c94132bb83d3fc7e7715d3b9097dc16d82c51)
- [limits.action_block.minutes](resources--rate_limiter--reference--group-001.md#canonical-06476eecefd3b64453e69b0df8b517f59809dad1ec3d0fc7ca5866d793b55f90)
- [limits.action_block.seconds](resources--rate_limiter--reference--group-001.md#canonical-059ac76ecb6d941646ff731dd7d283117053c5c113d1868d1b58ec4d70c2dc5e)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-b72227b0b1faa3fc2e08fc596d8c94132bb83d3fc7e7715d3b9097dc16d82c51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afab6b96657a6835b3cf3d13a9015070ab6a03beb9717b26af069c928e2f292c"></a>

## limits.action_block.hours — limits.action_block.hours / 28a68eb95536 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- limits.action_block.hours

<a id="canonical-4ce929d78af865971f520c19efe9f8d385e828514bd0b8b0d82e6c82491b14fa"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-ce391fd90059f4ca1eeba5d682b2740fdc378e02c2480166da8c373d4b7e940a"></a>

## Direct properties — limits.action_block.hours / 28a68eb95536 / 3

<a id="canonical-44eb79b7898e20d193049011e1ef34895be90863cc1141e80f6de08ca1f35168"></a>

<a id="canonical-6fd3cc8e84551db799eb169cd2cf9cdef880bf90fae6ae07f316f6eb8cf70154"></a>

## duration property — limits.action_block.hours / 28a68eb95536 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

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

<a id="canonical-d819915a6858b2053f77a491d23e562cca127d46ae0482ea9637543c8001697e"></a>

## Next pages — limits.action_block.hours / 28a68eb95536 / 5

- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-06476eecefd3b64453e69b0df8b517f59809dad1ec3d0fc7ca5866d793b55f90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f584cc2ed89f9ecdf41d38eb539bf64184898f9ac3f8e8701ba76c2b150cd843"></a>

## limits.action_block.minutes — limits.action_block.minutes / dcb5618ea9f7 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- limits.action_block.minutes

<a id="canonical-12c5faa791bad12bb26ec8e955d39114d0c189701e6e1dac789b891b19c3c9aa"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f4734c56bcbebd9879832f6e8ddf2c0320fb322070f45460473a238986d726b"></a>

## Direct properties — limits.action_block.minutes / dcb5618ea9f7 / 3

<a id="canonical-1a836af83eabde45767b35db2af91b3a99466c8756a46f09113fbcd206722c0e"></a>

<a id="canonical-2852eb6e0f4127e9e0bdc2b333b9ae5efac923f603d564b1fa49695e4a319a22"></a>

## duration property — limits.action_block.minutes / dcb5618ea9f7 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

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

<a id="canonical-9324768d103e48d6d877e4ad8dc95d4c47121a2295ad42da73cb1bcd0179b669"></a>

## Next pages — limits.action_block.minutes / dcb5618ea9f7 / 5

- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-059ac76ecb6d941646ff731dd7d283117053c5c113d1868d1b58ec4d70c2dc5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a6ca456c9034061ce654bb13063ed0cf61d3c48c875ca7c210ddde9da0c8bfb"></a>

## limits.action_block.seconds — limits.action_block.seconds / fc0b585207a2 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- limits.action_block.seconds

<a id="canonical-2751e8f9d13897708897de57c7ca5e057b6255e0857c70006c756818f3c6bd54"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fd2e51811234d3ff1f4409f749b6192528334a9bd2ef0abaf8d5f38c6cc7647"></a>

## Direct properties — limits.action_block.seconds / fc0b585207a2 / 3

<a id="canonical-29dc165987e03867dbd5288081cc7d67c4308f989679d14fb6217e06ec1db823"></a>

<a id="canonical-e2b5f017b9be539ecea95cf26e10f9b192c1e8bd7a45c7f2a59358e5a46abf75"></a>

## duration property — limits.action_block.seconds / fc0b585207a2 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

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

<a id="canonical-739902334d72e02b7bfb26df4a59c1292cb3473060d770470d8cb59fcd624686"></a>

## Next pages — limits.action_block.seconds / fc0b585207a2 / 5

- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-fcb4709e42fe13cbe3c77c9da10de22801ed3e15fc5e3b7065c8777a86ee474f)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-90fde68ff5ef47582e5936112e6658c79e8175f67bce55a3d015920ead4b2cf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8f28457fa5db0c6c11b294484b6924ebfc5db6056f200b304c95adf918e9e4e"></a>

## limits.disabled — limits.disabled / a8aa701252ba / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- limits.disabled

<a id="canonical-6be65be12a2ee424f7ec6446f758378ceea8cfd2c56285bc3f98d2b917d6f569"></a>

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
disabled = {}
```

<a id="canonical-29799112e95595f5262ad877485e2317cc9db93f1cd33afc7734a7625650775d"></a>

## Direct properties — limits.disabled / a8aa701252ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f98227f8ebba3aafa794bd689d7e83ba5f309738d33f387ddb01438e255eab7e"></a>

## Next pages — limits.disabled / a8aa701252ba / 4

- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-0f915de64bbb8272e4c98c582fde7600740e976b21fcc3acf628fb3a42010caf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bb4e56aa7841b149bfd690a4013d7d935a6b75b08c6914e69a5c9d65068999f"></a>

## limits.leaky_bucket — limits.leaky_bucket / a9c397e387f4 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- limits.leaky_bucket

<a id="canonical-14f05a7aa223ea93257cef71335c8f072a8a84139ae2c483253b82822ebdf4c6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
leaky_bucket = {}
```

<a id="canonical-848e01cea391a9496c70b0096253b2b83002378bceef4e95a76e6594bf9cde71"></a>

## Direct properties — limits.leaky_bucket / a9c397e387f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7007311f8039f2cbb5c752ffa2af6a1cc154dd69e0330dff87f6fb83d367754e"></a>

## Next pages — limits.leaky_bucket / a9c397e387f4 / 4

- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-72b9ce499fefa810dc273ce14ef21e22f66f6167abc9a3162b92655bc59d878f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68c8cec921ee1be3d487ec9a44afa6cb95b86ad5ff64dadb74cbbcd4818f9b99"></a>

## limits.token_bucket — limits.token_bucket / 5910a7e02ad8 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- limits.token_bucket

<a id="canonical-53057ff219fac970fdd5cc864939b75379ffc55abfa5d232618a37d52a5aa8b5"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
token_bucket = {}
```

<a id="canonical-d5dfe7254efbd96a64a56be8c54a2b80df3eb8571faa9c59337e024f833e9997"></a>

## Direct properties — limits.token_bucket / 5910a7e02ad8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e95e3acbb8a5095da3cfae60e44e3af2ecf6ed2f095848aeedcfda00ab2a866c"></a>

## Next pages — limits.token_bucket / 5910a7e02ad8 / 4

- [limits](resources--rate_limiter--reference--group-001.md#canonical-f4740452ab2d3bc1b8bf722aeec28f483045e2df55384e15d135e505f5dd082a)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-71f67a13190659f6652797231d0bd7bcdc4a9539bcaf307aa577563a18c60677"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0612dfb29e8394a5246509f559e1b0d192617265d4c4008ae4069f93a16114cb"></a>

## timeouts — timeouts / 1302681c2cb5 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- timeouts

<a id="canonical-8995368b12a8264d5d73b395fee3ff1202f6fe8010cbbd055df355fd35e0c76f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-81abce661a5f32a853e38e1e6a1e685ea5231f608d6e57a193b36f2341fa301f"></a>

## Direct properties — timeouts / 1302681c2cb5 / 3

<a id="canonical-3307419e2c34ee9d7d4fbd04d69f8d86ac5438d54fda98200b80c65d6ac013b0"></a>

<a id="canonical-494b4054c8b773121b85155c17f7d85aacbc740d7deaa6b8f33daf7108ecfba6"></a>

## create property — timeouts / 1302681c2cb5 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f139a64290b180fbef4b8aa2609918e3aaec85dc369962928bc95d39339842e9"></a>

<a id="canonical-ace5585fb0ad774a01d5c8134ff764b804fb88c62ea8c3529b0faa2bed569fdb"></a>

## delete property — timeouts / 1302681c2cb5 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-4847522c7ac30e29d81ab3948641c65673a9f876d57f1dc530bb5c9bb963d006"></a>

<a id="canonical-b3dc493621d0ec3decbef44028d1d9a8201e06262c466a55dc3df08ee4253e82"></a>

## read property — timeouts / 1302681c2cb5 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-881b4e142b89f1268564a42617710074695f472ab63dc1e512e6dd7fe9a05a4a"></a>

<a id="canonical-40cdb531d15758fab79ad74be2f5dc7b7ab7e9c17d4fc250885f05a2c6d1afae"></a>

## update property — timeouts / 1302681c2cb5 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-35111a1fa5f495d139fc292caa883d235781d793b2b4f82a42ce08f70a39e9aa"></a>

## Next pages — timeouts / 1302681c2cb5 / 8

- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-c71857479c5fdec24ae2de7db1c7d7c2fadd6c396ee88f9da33557471ea0d42b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5705c79e3c3c409c54a8fb5bbaa646003460aa3636cb7bd9c6c611e299c8434e"></a>

## user_identification — user_identification / 91759a27a57a / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- user_identification

<a id="canonical-f60c6bb90a5fe22fae476e39c42aff138534dc4bab2c273599b99f0845df850a"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-c0903da781e49cb00aa25d994aa09bb4a6c40be070b94419d057a6cf80bae0d5"></a>

## Direct properties — user_identification / 91759a27a57a / 3

<a id="canonical-005a55d6d8c23467eb983d85e0b0c529f57de755465dbcf2c53d797b8f1cd7a9"></a>

<a id="canonical-7f92750fa191ea1158545cecfbcdabea5ecfa53e16f16173b315f0cc15d99722"></a>

## kind property — user_identification / 91759a27a57a / 4

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

<a id="canonical-2234a6025d598f3aa48910f89232367bcbe94c87ba511f604dd29fef70b25b1b"></a>

<a id="canonical-1191e5d28a26507282746047ef5bad5f9e8b812a50a597d0da788b698d8c0187"></a>

## name property — user_identification / 91759a27a57a / 5

Type: `"string"`. Optional.

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

<a id="canonical-f0da8eb213b5c3308390c69054693ab7fde522cf31219f6f529187deb455687e"></a>

<a id="canonical-8f424ff4967d6bb047433ca14fb73162608a1bbb98d551ca13f8c5da5ff827b2"></a>

## namespace property — user_identification / 91759a27a57a / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-f792fe7eb63a5439fa5eeb3bc383959fe8b60391693c1c32423d31932fd35d74"></a>

<a id="canonical-0be8d12e05ea966297f8cc918f8bbedd2c13d3a0a20206abaeeccc51f54c691c"></a>

## tenant property — user_identification / 91759a27a57a / 7

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

<a id="canonical-61043647d1f67837b7848e6d18d07abf59098825441efd0bdf4328970bc6aa0c"></a>

<a id="canonical-fab55f74709be33e6a05747ae6931004e1e44ea8432c4d942910db937ac363f9"></a>

## uid property — user_identification / 91759a27a57a / 8

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

<a id="canonical-9eb6f756f4fbd884dca66f5bbcf4292de1834f584ee27ac17e2acd5d0e1471c4"></a>

## Next pages — user_identification / 91759a27a57a / 9

- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
