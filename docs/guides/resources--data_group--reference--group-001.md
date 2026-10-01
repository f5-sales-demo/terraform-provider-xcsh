---
page_title: "xcsh_data_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group reference."
---

# xcsh_data_group reference

<a id="canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13f189694f768052a6aa19147eafab1e88e955f531f1164b5100178b709b012e"></a>

## Property reference — Property reference / 6f3faae483d9 / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- Property reference

<a id="canonical-c847f274e61a8c954d0958808a8e7b38e7a0a9300f4358cec036d93ad46ce0f0"></a>

## Direct properties — Property reference / 6f3faae483d9 / 3

- [address_records](resources--data_group--reference--group-001.md#canonical-602d0d90ceccd05c0a362fbcf6fde4cd7aa988da6724568503241d71060f495c): complete subsection reference.

<a id="canonical-92380fcba6e886d2f06f0d44f307441e8ab6c5cd33b1361c5c11d45ad3dfd5f5"></a>

<a id="canonical-2a4ae7c28aa5afabcad96728c304b1c5881f75240b4ad255545e8cb1d5753177"></a>

## annotations property — Property reference / 6f3faae483d9 / 4

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

<a id="canonical-7d46d9f64b5869e44947264e87ed3a27c078a01b1861dcd9666c7795a36854be"></a>

<a id="canonical-3f7a1af2e83e9b5dcb1f7f7d9738c0a43f16fb2099440fa3bc838f6d4aa4e768"></a>

## description property — Property reference / 6f3faae483d9 / 5

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

<a id="canonical-c95f851fe9107448ee292179632d6ee3712b94462bf10fa4d067a146916a33da"></a>

<a id="canonical-0a4ed8446a28666d83ecd38760495c7b639edd2c19bd22ca3b906f2486bff591"></a>

## disable property — Property reference / 6f3faae483d9 / 6

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

<a id="canonical-4ce3f72728baa48c40739b90e888bab13b3d19c28889796c98b294357e546229"></a>

<a id="canonical-6966a173ab75ed711c45a38d1faf7fa9e965a321151c59b53d9d222ac35491c3"></a>

## id property — Property reference / 6f3faae483d9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [integer_records](resources--data_group--reference--group-001.md#canonical-b0968acd82eee4cf6e53f78293a0b1c0e17b10e30fe00f086044402b97580c6a): complete subsection reference.

<a id="canonical-81c2dcf8bb2269ac64df07e83ecc7e0270030544328af6088706441cf35e4ccc"></a>

<a id="canonical-ba1dc984228855039e0e76850d1936443387eb9ce692e286293608bdbe57a657"></a>

## labels property — Property reference / 6f3faae483d9 / 8

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

<a id="canonical-d9ea7640c67f746f36d4b898ada5a4c74d9372d2f142b559fd9b2fff063067bd"></a>

<a id="canonical-5b9416db04b74164b0d8dc62022e254f3aea869bd6db0ed6ee6551c0f8c467f0"></a>

## name property — Property reference / 6f3faae483d9 / 9

Type: `"string"`. Required.

Name of the Data Group. Must be unique within the namespace.

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

<a id="canonical-0b7d88b68a615eb5942c6586453c2da824edb145a7bb560f7d0c9cac8523898b"></a>

<a id="canonical-e3d0b5f562687e9c1bf3d964402584edb6413353c348e4fe0968de9e195df289"></a>

## namespace property — Property reference / 6f3faae483d9 / 10

Type: `"string"`. Required.

Namespace where the Data Group is created.

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

- [string_records](resources--data_group--reference--group-001.md#canonical-c3156183f9aeb02c49a4e945883fddb6d3af27b67771bb753bb1e36e67678cbb): complete subsection reference.

- [timeouts](resources--data_group--reference--group-001.md#canonical-5ee286714b1f6498cde053581f86737dd8fe3e1aaa63f3a41798fcd44a4a6041): complete subsection reference.

<a id="canonical-6e42a2cacc707e5f6bb1fbdd2f49e927892cd948a6d217594a764e3c7b6d25e1"></a>

## All schema paths — Property reference / 6f3faae483d9 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_records` | [address_records](resources--data_group--reference--group-001.md#canonical-3541d101a7d24425962a529ffb6cc2a18850503287588ee02fe840bc32bf05e8) |
| `address_records.records` | [address_records.records](resources--data_group--reference--group-001.md#canonical-cf9713387787593642467ae35d00eebebe46454e544bd9ffabe9a7bf4adc9cfe) |
| `annotations` | [annotations](resources--data_group--reference--group-001.md#canonical-92380fcba6e886d2f06f0d44f307441e8ab6c5cd33b1361c5c11d45ad3dfd5f5) |
| `description` | [description](resources--data_group--reference--group-001.md#canonical-7d46d9f64b5869e44947264e87ed3a27c078a01b1861dcd9666c7795a36854be) |
| `disable` | [disable](resources--data_group--reference--group-001.md#canonical-c95f851fe9107448ee292179632d6ee3712b94462bf10fa4d067a146916a33da) |
| `id` | [id](resources--data_group--reference--group-001.md#canonical-4ce3f72728baa48c40739b90e888bab13b3d19c28889796c98b294357e546229) |
| `integer_records` | [integer_records](resources--data_group--reference--group-001.md#canonical-03531bd7030803bc59e35c623d0a078c3900688eb8e2fc52b65ce870c0bf1c6c) |
| `integer_records.records` | [integer_records.records](resources--data_group--reference--group-001.md#canonical-4a4a7cc14950c49e7f34a9cce094eaebb19991abcbbe781b6d021078c14d1cc9) |
| `labels` | [labels](resources--data_group--reference--group-001.md#canonical-81c2dcf8bb2269ac64df07e83ecc7e0270030544328af6088706441cf35e4ccc) |
| `name` | [name](resources--data_group--reference--group-001.md#canonical-d9ea7640c67f746f36d4b898ada5a4c74d9372d2f142b559fd9b2fff063067bd) |
| `namespace` | [namespace](resources--data_group--reference--group-001.md#canonical-0b7d88b68a615eb5942c6586453c2da824edb145a7bb560f7d0c9cac8523898b) |
| `string_records` | [string_records](resources--data_group--reference--group-001.md#canonical-0688350d83195bd6c5f5a88332f79911d8e43ac84d367aa23446f8d6f2a11730) |
| `string_records.records` | [string_records.records](resources--data_group--reference--group-001.md#canonical-a8ff98bb4367424c3e1286f90d6f8645e469410000b32a7464ac9efbf0f872d2) |
| `timeouts` | [timeouts](resources--data_group--reference--group-001.md#canonical-63c80158ea0fa0f2ddee253a3a78a404e98bd11d480a7adc099c95fa505e9392) |
| `timeouts.create` | [timeouts.create](resources--data_group--reference--group-001.md#canonical-0532650da064a1c9a45a6ed5703c33c69114bd651739316bfbd5a58425dfd56e) |
| `timeouts.delete` | [timeouts.delete](resources--data_group--reference--group-001.md#canonical-f606159784ccf3ba63e0d2db2b3030b58325f3dc6836da77af664a82d38676ee) |
| `timeouts.read` | [timeouts.read](resources--data_group--reference--group-001.md#canonical-3f5c316eed73e1507be78e37b77dbb060456e139817907809a6cbf16a20a154a) |
| `timeouts.update` | [timeouts.update](resources--data_group--reference--group-001.md#canonical-65d440c590b2646ca442467c4cd8d14c31df37aeff52217879c080913667318f) |

<a id="canonical-ca1924a09d0e28ab4ad78f8e2317444c75ad07661c77ee0149859d0c5276c718"></a>

## Next pages — Property reference / 6f3faae483d9 / 12

- [address_records](resources--data_group--reference--group-001.md#canonical-602d0d90ceccd05c0a362fbcf6fde4cd7aa988da6724568503241d71060f495c)
- [integer_records](resources--data_group--reference--group-001.md#canonical-b0968acd82eee4cf6e53f78293a0b1c0e17b10e30fe00f086044402b97580c6a)
- [string_records](resources--data_group--reference--group-001.md#canonical-c3156183f9aeb02c49a4e945883fddb6d3af27b67771bb753bb1e36e67678cbb)
- [timeouts](resources--data_group--reference--group-001.md#canonical-5ee286714b1f6498cde053581f86737dd8fe3e1aaa63f3a41798fcd44a4a6041)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)

<a id="canonical-602d0d90ceccd05c0a362fbcf6fde4cd7aa988da6724568503241d71060f495c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6c7cd2c38e32e2749b15b5ff71d678b5036357627c1d70d0f171aaf530bca81"></a>

## address_records — address_records / 42ba1b9a1e58 / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- address_records

<a id="canonical-3541d101a7d24425962a529ffb6cc2a18850503287588ee02fe840bc32bf05e8"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: address\_records, integer\_records, string\_records\] Address Record. Data group with
address record List.

Upstream description:

Data group with address record List.

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

- [address_records](resources--data_group--reference--group-001.md#canonical-3541d101a7d24425962a529ffb6cc2a18850503287588ee02fe840bc32bf05e8)
- [integer_records](resources--data_group--reference--group-001.md#canonical-03531bd7030803bc59e35c623d0a078c3900688eb8e2fc52b65ce870c0bf1c6c)
- [string_records](resources--data_group--reference--group-001.md#canonical-0688350d83195bd6c5f5a88332f79911d8e43ac84d367aa23446f8d6f2a11730)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
address_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-64145360585bec336f5be33f0d5a0cd0e724c1203d181bbc2d47602a9bb10000"></a>

## Direct properties — address_records / 42ba1b9a1e58 / 3

<a id="canonical-cf9713387787593642467ae35d00eebebe46454e544bd9ffabe9a7bf4adc9cfe"></a>

<a id="canonical-f4f88853d8512b5cf210e4900220b6667d8c6f912675eee911745f88a59c900b"></a>

## records property — address_records / 42ba1b9a1e58 / 4

Type: `["map", "string"]`. Optional.

Address records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-2964b994722197609190f4c964e534d9279dce1fa707565d5900d53ca65a95f9"></a>

## Next pages — address_records / 42ba1b9a1e58 / 5

- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)

<a id="canonical-b0968acd82eee4cf6e53f78293a0b1c0e17b10e30fe00f086044402b97580c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b19eb9e0e46752c30c32cbf94dae2262b1b22cbffd01096924f23c68b9beda5d"></a>

## integer_records — integer_records / b4470d996495 / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- integer_records

<a id="canonical-03531bd7030803bc59e35c623d0a078c3900688eb8e2fc52b65ce870c0bf1c6c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for integer records.

Upstream description:

Data group with integer record List.

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
integer_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-2febf85c9bf6150d6b41129c2b92a55f7c2718e57ac42972cb84510047a8b246"></a>

## Direct properties — integer_records / b4470d996495 / 3

<a id="canonical-4a4a7cc14950c49e7f34a9cce094eaebb19991abcbbe781b6d021078c14d1cc9"></a>

<a id="canonical-6022d89bed28c1e5ce1e24407d9a4fabfdf3b83a797837d61f92e56b94557443"></a>

## records property — integer_records / b4470d996495 / 4

Type: `["map", "string"]`. Optional.

Integer records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  }
}
```

<a id="canonical-3b7dee1edab6c0dfe16f8245e07dee4e9762ed2880d7cd92cc039b550e955b01"></a>

## Next pages — integer_records / b4470d996495 / 5

- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)

<a id="canonical-c3156183f9aeb02c49a4e945883fddb6d3af27b67771bb753bb1e36e67678cbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ac4a59f9556d7cdd7d4461973f85c416321851c16b94cffa758faac6d17e504"></a>

## string_records — string_records / 60ab37879ed1 / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- string_records

<a id="canonical-0688350d83195bd6c5f5a88332f79911d8e43ac84d367aa23446f8d6f2a11730"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for string records.

Upstream description:

Data group with strings record List.

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
string_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-5566b406810c81f091dedf7e892e9764438763e330f96b3ef12f5b2e5b64237e"></a>

## Direct properties — string_records / 60ab37879ed1 / 3

<a id="canonical-a8ff98bb4367424c3e1286f90d6f8645e469410000b32a7464ac9efbf0f872d2"></a>

<a id="canonical-68e65021dfa1641fd9171cd16074153020d925f1a452baac56ecf737b498e8c5"></a>

## records property — string_records / 60ab37879ed1 / 4

Type: `["map", "string"]`. Optional.

String records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-d8e08737b03347778f6bd7c9bf62c23141a7ff7eaaf9c898e7792e22c0e4e104"></a>

## Next pages — string_records / 60ab37879ed1 / 5

- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)

<a id="canonical-5ee286714b1f6498cde053581f86737dd8fe3e1aaa63f3a41798fcd44a4a6041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c7a95dc1aca58ee2f29168c6a3ddfb7e682b2903e9e0c55d2a2f5b9f45e1761"></a>

## timeouts — timeouts / ea545c83f5cb / 2

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- timeouts

<a id="canonical-63c80158ea0fa0f2ddee253a3a78a404e98bd11d480a7adc099c95fa505e9392"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-40d9440461746a1861163e22963ecfb8005af207cffd668fdd35f0735f6d3d9e"></a>

## Direct properties — timeouts / ea545c83f5cb / 3

<a id="canonical-0532650da064a1c9a45a6ed5703c33c69114bd651739316bfbd5a58425dfd56e"></a>

<a id="canonical-ab8ffe65045eb9eb41bfc35c7acc5ef76e0a37c41feeb04ed169214b88c5ef02"></a>

## create property — timeouts / ea545c83f5cb / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f606159784ccf3ba63e0d2db2b3030b58325f3dc6836da77af664a82d38676ee"></a>

<a id="canonical-6f42c97477e435eb3072f687aa39be6eac46f72f403fb747027e6ce6ee3f8734"></a>

## delete property — timeouts / ea545c83f5cb / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3f5c316eed73e1507be78e37b77dbb060456e139817907809a6cbf16a20a154a"></a>

<a id="canonical-fa40ddbf0a806e79bee5bf179947c78cbc99e44b450ed4452500f05a9140e24e"></a>

## read property — timeouts / ea545c83f5cb / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-65d440c590b2646ca442467c4cd8d14c31df37aeff52217879c080913667318f"></a>

<a id="canonical-b6f365dfcd82fc92216b7ecf8b042e36a9949a9434124affa38b97fb36c9186f"></a>

## update property — timeouts / ea545c83f5cb / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-62e92289266ea06e00215dda85e9bee13986e122c04d654324d732d659eafe5b"></a>

## Next pages — timeouts / ea545c83f5cb / 8

- [Property reference](resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- [xcsh_data_group](../resources/data_group.md#canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd)
