---
page_title: "xcsh_cloud_user_account reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account reference."
---

# xcsh_cloud_user_account reference

<a id="canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bcaa546164c93c484e92b09ee178f4952ee42406d6cf081e59305fadfc5b2fa"></a>

## Property reference — Property reference / 3977d040a00e / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- Property reference

<a id="canonical-f959207fbc57d49373432fbfd4d8081d86e2684d6c293fc977fc2a496d097adf"></a>

## Direct properties — Property reference / 3977d040a00e / 3

<a id="canonical-3aac5f4da40ac8976e6ce248c38f829014e7cd926279924b4e280eb6549669e4"></a>

<a id="canonical-4bf99401f228aef0bf0e4ef259a3a02c96a1e5a74b2c93823b7948f0da0bfadb"></a>

## annotations property — Property reference / 3977d040a00e / 4

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

- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654): complete subsection reference.

<a id="canonical-ca04287c721475c8b64115f20d3e6281cb660102661d3242a96c3cdfecf2dd21"></a>

<a id="canonical-635e1c052761adf0a3beced264350318045a660cbe56340614248b9340671884"></a>

## description property — Property reference / 3977d040a00e / 5

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

<a id="canonical-3d854c5161e923d79f70b60104bb19829ded5cd4db81bb6589f173100db2dc1b"></a>

<a id="canonical-471b2f62ba163f257fa9bfc738336047a71413f2ddd61bf26d814921c4b7faf3"></a>

## disable property — Property reference / 3977d040a00e / 6

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

<a id="canonical-df1499ab3eef09aab86c368d16878ee69165fdc7d70a58ad38efe193336fe042"></a>

<a id="canonical-9dc039b8b6d7c5a58fb24e42e414a14c2ca1b9781fb4382a1dda8c8015e52f68"></a>

## id property — Property reference / 3977d040a00e / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d3e88152d7f7075dee855c882d754fdd8fd6c6c8fc6039d26064210fcdda3983"></a>

<a id="canonical-e69ea708f83e48713a0b83b4e81455dd85e43785be8fbc2a4d5a768a24f3106c"></a>

## labels property — Property reference / 3977d040a00e / 8

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

<a id="canonical-ff529a1b6162c33e5b8d6d8b4df53a498390c22f0400359e30fcd853fd6686aa"></a>

<a id="canonical-8d01fa7187f7b6b8da171424e3e7e3b8b6a91398621260afe4084291c55e9845"></a>

## name property — Property reference / 3977d040a00e / 9

Type: `"string"`. Required.

Name of the Cloud User Account. Must be unique within the namespace.

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

<a id="canonical-50057b1eda6338788da29e53e703353ade1fece871f4421bdd6c0be49498ff3e"></a>

<a id="canonical-bd4c8007d29f7bafefcef53e89e1eedd9f80081e4ea0187989bed26ad1d1341d"></a>

## namespace property — Property reference / 3977d040a00e / 10

Type: `"string"`. Required.

Namespace where the Cloud User Account is created.

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

- [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-295f7970adb4a94de98e96380d058dc8db8b4f2d7a14edf838f9d30a8324b8ef): complete subsection reference.

<a id="canonical-1fdeec821a06fbb72359d571cef714bea4f9ccc876e4789eee5cffa09c27438d"></a>

## All schema paths — Property reference / 3977d040a00e / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_user_account--reference--group-001.md#canonical-3aac5f4da40ac8976e6ce248c38f829014e7cd926279924b4e280eb6549669e4) |
| `aws_provider` | [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-8f4d3447ec0bfc52b741b5d8123cb8638fdf69cbc35a8e79c28b9c44f40349ba) |
| `aws_provider.aws_account_number` | [aws_provider.aws_account_number](resources--cloud_user_account--reference--group-001.md#canonical-44602df315deff361c76bc99a524147a2344a0421375d6d52db879ad00f8a129) |
| `aws_provider.aws_assume_role` | [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-3b2fd0fb8a0125dcbec833ab65e80d0ba1eaf10b742f6304bd1cccc000238e57) |
| `aws_provider.aws_assume_role.custom_external_id` | [aws_provider.aws_assume_role.custom_external_id](resources--cloud_user_account--reference--group-001.md#canonical-961bf47c07daa37b32e0fc10f32485d89ff35994a9de9e291d4c4a789c7b3ef4) |
| `aws_provider.aws_assume_role.duration_seconds` | [aws_provider.aws_assume_role.duration_seconds](resources--cloud_user_account--reference--group-001.md#canonical-aac2e8a82743964e28cc707bec88e89b5f760636780d328c06746bcb6cca09a1) |
| `aws_provider.aws_assume_role.external_id_is_optional` | [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-ae74b6afc4cdb3cf7b7ce295ad1b6ed405e6ac49dc601b20184e4690f16a7319) |
| `aws_provider.aws_assume_role.external_id_is_tenant_id` | [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-8ff006c672de846b2a01ad13b4653ecfd56c078b62ea50b389f0eabeef3ca8c2) |
| `aws_provider.aws_assume_role.role_arn` | [aws_provider.aws_assume_role.role_arn](resources--cloud_user_account--reference--group-001.md#canonical-67af5acb30c03df60e667211716cfe8edfbafe6428e6c6e24c156c1297b02340) |
| `aws_provider.aws_assume_role.session_name` | [aws_provider.aws_assume_role.session_name](resources--cloud_user_account--reference--group-001.md#canonical-41aaf7001eebe7c2c714cc6f884684374a03f90d44c8b2773a4577a8a7945934) |
| `aws_provider.aws_assume_role.session_tags` | [aws_provider.aws_assume_role.session_tags](resources--cloud_user_account--reference--group-001.md#canonical-11d37508a71d3a82cc4c775b15984d737e8d15d3651c777f94a44c910c540328) |
| `aws_provider.aws_secret_key` | [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-9bac1e67a1680cb75bf7cbf349b8ef58e0454cb3c9c2446f1d607b267ac73e70) |
| `aws_provider.aws_secret_key.access_key` | [aws_provider.aws_secret_key.access_key](resources--cloud_user_account--reference--group-001.md#canonical-2c46d51ed52ef9fc608a4b88d5e9a071dc94b9b10c103eacd1157b1839effd60) |
| `aws_provider.aws_secret_key.secret_key` | [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-7aa363d5dbc83cc4af6e12e03b6daeb9949f719dc7006ec56dbde5f3cb5f9fd3) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-39bf284886e189ed2a360a445fd0beb95a33686247d3a0f7267a009be790f65b) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](resources--cloud_user_account--reference--group-001.md#canonical-594047a6f8350c71e6906bcd3d2cd3da5c8add6fa4524e1fa2e85866ffbd4aaf) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location](resources--cloud_user_account--reference--group-001.md#canonical-8f17da804134a131e844dda6a63954a041448cbb6e6c51e616d997f4891be50a) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider](resources--cloud_user_account--reference--group-001.md#canonical-12856e7840c809ee06d78fd46db8af1bbd8a26c265536a1fb3b4d6bba222f7b7) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info` | [aws_provider.aws_secret_key.secret_key.clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-c6c0635a6afe754545f703ed17a0b23d9d2a4b8d0916cff0fcd3496f43ce4592) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref](resources--cloud_user_account--reference--group-001.md#canonical-5a5b870f675e2cd35ca31693ef2ac0bf7a0c3a190d34e276f68b0bfd1733f6ea) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.url](resources--cloud_user_account--reference--group-001.md#canonical-cd1354937833c67e235aab2270f089bde036375beef0404e51082a1f37c04082) |
| `description` | [description](resources--cloud_user_account--reference--group-001.md#canonical-ca04287c721475c8b64115f20d3e6281cb660102661d3242a96c3cdfecf2dd21) |
| `disable` | [disable](resources--cloud_user_account--reference--group-001.md#canonical-3d854c5161e923d79f70b60104bb19829ded5cd4db81bb6589f173100db2dc1b) |
| `id` | [id](resources--cloud_user_account--reference--group-001.md#canonical-df1499ab3eef09aab86c368d16878ee69165fdc7d70a58ad38efe193336fe042) |
| `labels` | [labels](resources--cloud_user_account--reference--group-001.md#canonical-d3e88152d7f7075dee855c882d754fdd8fd6c6c8fc6039d26064210fcdda3983) |
| `name` | [name](resources--cloud_user_account--reference--group-001.md#canonical-ff529a1b6162c33e5b8d6d8b4df53a498390c22f0400359e30fcd853fd6686aa) |
| `namespace` | [namespace](resources--cloud_user_account--reference--group-001.md#canonical-50057b1eda6338788da29e53e703353ade1fece871f4421bdd6c0be49498ff3e) |
| `timeouts` | [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-2632692b502039ebbcd48f892a32c66d43d720f2a2954f136c2c9af4adc1bc54) |
| `timeouts.create` | [timeouts.create](resources--cloud_user_account--reference--group-001.md#canonical-c9f0bb22681ff2783a8afaeb6b70f65a5562d0a0ee489fe44e9d26c9f726548b) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_user_account--reference--group-001.md#canonical-1a3771343f629b1b3f632606c505a31921794dddf7bd678639cd10cadec9251d) |
| `timeouts.read` | [timeouts.read](resources--cloud_user_account--reference--group-001.md#canonical-58e73ff0d03fece550d5b5bc8a1cd47e2bc53cb74e908c88ee5752f59210ef04) |
| `timeouts.update` | [timeouts.update](resources--cloud_user_account--reference--group-001.md#canonical-c26adf69b56e3e11c5437d79978aaff63d4954cd5ff44cba4ddd6faff3c03656) |

<a id="canonical-bbc76f85a16be981139d9776a56c4bf804b47ccd77c8d750c66d3dd3a41b809e"></a>

## Next pages — Property reference / 3977d040a00e / 12

- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-295f7970adb4a94de98e96380d058dc8db8b4f2d7a14edf838f9d30a8324b8ef)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7882a5697015a50fad507a5404832841a89fbdebcd5fbe6c723f691961144dc9"></a>

## aws_provider — aws_provider / 094ee9334c74 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- aws_provider

<a id="canonical-8f4d3447ec0bfc52b741b5d8123cb8638fdf69cbc35a8e79c28b9c44f40349ba"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws provider.

Upstream description:

Create AWS Provider Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_account_number"),
  validators.ConflictingObjectAttributes("aws_assume_role",
    "aws_secret_key")}
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
  "x-ves-oneof-field-aws_authentication_type": "[\"aws_assume_role\",\"aws_secret_key\"]"
}
```

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

<a id="canonical-60f1406397eabd05d8a17345fec53906ab890134db7aa981301c6f054669dd0c"></a>

## Direct properties — aws_provider / 094ee9334c74 / 3

<a id="canonical-44602df315deff361c76bc99a524147a2344a0421375d6d52db879ad00f8a129"></a>

<a id="canonical-03ba02df1986e493af73ca102b3b86aad8a8de9eafb11c8b03eb481fbb219b76"></a>

## aws_account_number property — aws_provider / 094ee9334c74 / 4

Type: `"string"`. Optional.

Account Number. 12 Digit Account Number.

Upstream description:

12 Digit Account Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  }
}
```

- [aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d): complete subsection reference.

- [aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc): complete subsection reference.

<a id="canonical-2ef641ceb3aa279ddec2cecbe3dd340e59b436c05fe9fceee9d722b162e382f4"></a>

## Next pages — aws_provider / 094ee9334c74 / 5

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba80a224eb67b5a3cebfce5553533a9b3f0896ae73a0d616b6190a76d4c708d0"></a>

## aws_provider.aws_assume_role — aws_provider.aws_assume_role / c4f680fa5d33 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- aws_provider.aws_assume_role

<a id="canonical-3b2fd0fb8a0125dcbec833ab65e80d0ba1eaf10b742f6304bd1cccc000238e57"></a>

Type: `"object"`. single nested block, Optional.

AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration_seconds",
    "role_arn",
    "session_name"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_optional"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_tenant_id"),
  validators.ConflictingObjectAttributes("external_id_is_optional",
    "external_id_is_tenant_id")}
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
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-10f279f5ae00005fdb0eac1928cac14e6f021038bc0e48119dda3f40fd83af8e"></a>

## Direct properties — aws_provider.aws_assume_role / c4f680fa5d33 / 3

<a id="canonical-961bf47c07daa37b32e0fc10f32485d89ff35994a9de9e291d4c4a789c7b3ef4"></a>

<a id="canonical-cea6791b4ddf88b9e013845630eb71b8c4c5c94075b00ddb97f321178bd5e20f"></a>

## custom_external_id property — aws_provider.aws_assume_role / c4f680fa5d33 / 4

Type: `"string"`. Optional.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="canonical-aac2e8a82743964e28cc707bec88e89b5f760636780d328c06746bcb6cca09a1"></a>

<a id="canonical-c56c7a4ac2d3ef2a98efa2df2748dc3833c2a3e9a2928552e593e06d54e3afde"></a>

## duration_seconds property — aws_provider.aws_assume_role / c4f680fa5d33 / 5

Type: `"number"`. Optional.

The duration, in seconds of the role session.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-15f6c18ef047dcf212c8e6795e0d3dfb2d364cf4036c36c5a990639b40f0c5fe): complete subsection reference.

- [external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-bfae67d44327d207a63e37e519c5a199ac5268168fcc5af0232b6e9b107d1d58): complete subsection reference.

<a id="canonical-67af5acb30c03df60e667211716cfe8edfbafe6428e6c6e24c156c1297b02340"></a>

<a id="canonical-efed09cfb532b17c709b97712f67b6c4972bb8b83fdf2ab783604ef0ac4707c0"></a>

## role_arn property — aws_provider.aws_assume_role / c4f680fa5d33 / 6

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="canonical-41aaf7001eebe7c2c714cc6f884684374a03f90d44c8b2773a4577a8a7945934"></a>

<a id="canonical-bb3f4e55d71305cca76ed4f5955ecb0838345e4df729cfa6a31884f36ed1bb3b"></a>

## session_name property — aws_provider.aws_assume_role / c4f680fa5d33 / 7

Type: `"string"`. Optional.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="canonical-11d37508a71d3a82cc4c775b15984d737e8d15d3651c777f94a44c910c540328"></a>

<a id="canonical-05195750bf14cb3e7755cf3da0318c2b62a9d29f88e42bd0c1e7fa9acb9ddb94"></a>

## session_tags property — aws_provider.aws_assume_role / c4f680fa5d33 / 8

Type: `["map", "string"]`. Optional.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-0e31937e9af6517879412734ee2525bbc7ddcf0c7b1c7807761b62ab5f6f1d9d"></a>

## Next pages — aws_provider.aws_assume_role / c4f680fa5d33 / 9

- [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-15f6c18ef047dcf212c8e6795e0d3dfb2d364cf4036c36c5a990639b40f0c5fe)
- [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-bfae67d44327d207a63e37e519c5a199ac5268168fcc5af0232b6e9b107d1d58)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-15f6c18ef047dcf212c8e6795e0d3dfb2d364cf4036c36c5a990639b40f0c5fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62b02af7b7baf7e40c5c9c0cc1e4347e48693d3c60838359fc6c380e5cd74e4d"></a>

## aws_provider.aws_assume_role.external_id_is_optional — aws_provider.aws_assume_role.external_id_is_optional / 8b42517659c2 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d)
- aws_provider.aws_assume_role.external_id_is_optional

<a id="canonical-ae74b6afc4cdb3cf7b7ce295ad1b6ed405e6ac49dc601b20184e4690f16a7319"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external id is optional.

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
external_id_is_optional = {}
```

<a id="canonical-a996db8cfbaf18e62ec34d561b0f3212de6cf2d0bfc5911f3502143bdfc09e03"></a>

## Direct properties — aws_provider.aws_assume_role.external_id_is_optional / 8b42517659c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed7922d23c1e8d60d381cd44e1b09c597bb8bdf7542f44a71b5c7e01adeacedd"></a>

## Next pages — aws_provider.aws_assume_role.external_id_is_optional / 8b42517659c2 / 4

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-bfae67d44327d207a63e37e519c5a199ac5268168fcc5af0232b6e9b107d1d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80cce3b3826e34b7c39fd0c0c16dba3f7f6d6f052f7f4abb117da0a17cdea60e"></a>

## aws_provider.aws_assume_role.external_id_is_tenant_id — aws_provider.aws_assume_role.external_id_is_tenant_id / 12d84a9cf1d4 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d)
- aws_provider.aws_assume_role.external_id_is_tenant_id

<a id="canonical-8ff006c672de846b2a01ad13b4653ecfd56c078b62ea50b389f0eabeef3ca8c2"></a>

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
external_id_is_tenant_id = {}
```

<a id="canonical-d30bc47e5e07d8f3b7a8304f04ef89f53189d032d82a288a3c50aa85e260dde8"></a>

## Direct properties — aws_provider.aws_assume_role.external_id_is_tenant_id / 12d84a9cf1d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c1454d27847b9c7f34230d92bad9f17facfe60e0f789603020f17eea4c4a87b"></a>

## Next pages — aws_provider.aws_assume_role.external_id_is_tenant_id / 12d84a9cf1d4 / 4

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-2368ea2424180340d053ee7f4d0e5fe4ec7adf8b00e4a6fffe7d06970fc6980d)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d93f23bba2902047dad01bd520b054c362eb3508b7f94e7d3bfa4d26f42bfd88"></a>

## aws_provider.aws_secret_key — aws_provider.aws_secret_key / c4e3e013d565 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- aws_provider.aws_secret_key

<a id="canonical-9bac1e67a1680cb75bf7cbf349b8ef58e0454cb3c9c2446f1d607b267ac73e70"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("access_key")}
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
aws_secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-fd3419e793e37392e7ed0e40e272fe38f3a6aeb51069e9844881761d3b0924e1"></a>

## Direct properties — aws_provider.aws_secret_key / c4e3e013d565 / 3

<a id="canonical-2c46d51ed52ef9fc608a4b88d5e9a071dc94b9b10c103eacd1157b1839effd60"></a>

<a id="canonical-28252669694e4b4477cb781a71610d0171e3bbb16de04abbfd67d90439354dae"></a>

## access_key property — aws_provider.aws_secret_key / c4e3e013d565 / 4

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2): complete subsection reference.

<a id="canonical-3e6b3396b5707d718d03ce27d00565f2b286e21688b2719e48cb69d710503e98"></a>

## Next pages — aws_provider.aws_secret_key / c4e3e013d565 / 5

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b90dad9d81fc712f20e9dbec59544bfb583cdc39281c1a947183fd82da627d9f"></a>

## aws_provider.aws_secret_key.secret_key — aws_provider.aws_secret_key.secret_key / b0e7cf76ffc2 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc)
- aws_provider.aws_secret_key.secret_key

<a id="canonical-7aa363d5dbc83cc4af6e12e03b6daeb9949f719dc7006ec56dbde5f3cb5f9fd3"></a>

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
secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0796b3a098e0a95e2efa91b95c6e8422477b10a023588338522033935e53077f"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key / b0e7cf76ffc2 / 3

- [blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-0dfd98926ce2e793b02f5a3cd5f505ffdc388fa0b01aa71bb1c9c0b4582c1b40): complete subsection reference.

- [clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-5fdd10d8c7413d33add8d56f2aa85e94efe1fe78a53976f142f7aebdb8ecccd7): complete subsection reference.

<a id="canonical-5b7d4d19f0540bef76eade66ab4fd0220c1b6ee63045e4a29f3b49f3952cc329"></a>

## Next pages — aws_provider.aws_secret_key.secret_key / b0e7cf76ffc2 / 4

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-0dfd98926ce2e793b02f5a3cd5f505ffdc388fa0b01aa71bb1c9c0b4582c1b40)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-5fdd10d8c7413d33add8d56f2aa85e94efe1fe78a53976f142f7aebdb8ecccd7)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-0dfd98926ce2e793b02f5a3cd5f505ffdc388fa0b01aa71bb1c9c0b4582c1b40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32f661e6d21546013fb13c8fa5c92950ebe8ee0944e010247cbd79932bb7e4f4"></a>

## aws_provider.aws_secret_key.secret_key.blindfold_secret_info — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc)
- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2)
- aws_provider.aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-39bf284886e189ed2a360a445fd0beb95a33686247d3a0f7267a009be790f65b"></a>

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

<a id="canonical-d745f2fae538251673647ff22411b66ae9025fb02507d2d0241acd4a26dc83dd"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 3

<a id="canonical-594047a6f8350c71e6906bcd3d2cd3da5c8add6fa4524e1fa2e85866ffbd4aaf"></a>

<a id="canonical-119e1c1a873a1c6415e6749f3d7f7ab8227b376a245cfb528126d7c9a4fc4bb6"></a>

## decryption_provider property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 4

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

<a id="canonical-8f17da804134a131e844dda6a63954a041448cbb6e6c51e616d997f4891be50a"></a>

<a id="canonical-8222e4f6adc733dac3f59c926c70201130e1be5f02b5615721bfb71a576fb922"></a>

## location property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 5

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

<a id="canonical-12856e7840c809ee06d78fd46db8af1bbd8a26c265536a1fb3b4d6bba222f7b7"></a>

<a id="canonical-e4612e7c3765ed4324d13c011eaa7ecf692785155aeffcd8876d93d4502eeec0"></a>

## store_provider property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 6

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

<a id="canonical-c5e5470ef33c20268f6b155f8656022c1fbd8766abdaad68043397c108aa90c7"></a>

## Next pages — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / cd6fe96f0bba / 7

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-5fdd10d8c7413d33add8d56f2aa85e94efe1fe78a53976f142f7aebdb8ecccd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29b226e5de6b81b9b06c291dbe4fb00361c7768ec268b93a91b1996bc347af23"></a>

## aws_provider.aws_secret_key.secret_key.clear_secret_info — aws_provider.aws_secret_key.secret_key.clear_secret_info / 086086dfef8e / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-061af349d0eb9a7f8e5fdc5ff6b441755e47e6910aeecc86bc84d790cecc3654)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-e46346254feadac1e304ec15d59d2c52107ac1d046f17332528f11456e6b0edc)
- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="canonical-c6c0635a6afe754545f703ed17a0b23d9d2a4b8d0916cff0fcd3496f43ce4592"></a>

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

<a id="canonical-09cd536186fdabcfd49bf3b67acb42499333505d529c2408ffb11ea45607b69a"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key.clear_secret_info / 086086dfef8e / 3

<a id="canonical-5a5b870f675e2cd35ca31693ef2ac0bf7a0c3a190d34e276f68b0bfd1733f6ea"></a>

<a id="canonical-1acd9a5e38a2127cc58c801026900fca00541d5d18b72d0aee0bc078dd348cf1"></a>

## provider_ref property — aws_provider.aws_secret_key.secret_key.clear_secret_info / 086086dfef8e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-cd1354937833c67e235aab2270f089bde036375beef0404e51082a1f37c04082"></a>

<a id="canonical-ed4682e76abd6354175d5dfa56af6da9a8ea633c1b1b59134b1608b6e5320583"></a>

## url property — aws_provider.aws_secret_key.secret_key.clear_secret_info / 086086dfef8e / 5

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

<a id="canonical-aa7befe0eb5cf2141c281f7165c95f9d0845205b8b7e9c4354956878b88be5c0"></a>

## Next pages — aws_provider.aws_secret_key.secret_key.clear_secret_info / 086086dfef8e / 6

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-81a72fa5984175973e89ca051ac8a14687e0596396a8e4cae1be1e266a754ec2)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-295f7970adb4a94de98e96380d058dc8db8b4f2d7a14edf838f9d30a8324b8ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08a6e510332d5a477cd714d195f3927fa2dcc93214f12709e534b54f02a99804"></a>

## timeouts — timeouts / 099c9dcacc66 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- timeouts

<a id="canonical-2632692b502039ebbcd48f892a32c66d43d720f2a2954f136c2c9af4adc1bc54"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0b5fe143abba37deb631784e3ada5a3c1793194a1bdf5b16840c2afb8405af1f"></a>

## Direct properties — timeouts / 099c9dcacc66 / 3

<a id="canonical-c9f0bb22681ff2783a8afaeb6b70f65a5562d0a0ee489fe44e9d26c9f726548b"></a>

<a id="canonical-80152a8b754649073a4f62f2f4c6c610ecaff3a1b6244c6fba124fa89697acee"></a>

## create property — timeouts / 099c9dcacc66 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1a3771343f629b1b3f632606c505a31921794dddf7bd678639cd10cadec9251d"></a>

<a id="canonical-da7bdb31f3d76b2308c7df8b36d0977218817ddd7342a3a557ed59e98dd8d657"></a>

## delete property — timeouts / 099c9dcacc66 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-58e73ff0d03fece550d5b5bc8a1cd47e2bc53cb74e908c88ee5752f59210ef04"></a>

<a id="canonical-a539a1dcb0b7739c4ae31d4f1acc0fb7fabae56f8664fa2f7984ff768f9e244a"></a>

## read property — timeouts / 099c9dcacc66 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-c26adf69b56e3e11c5437d79978aaff63d4954cd5ff44cba4ddd6faff3c03656"></a>

<a id="canonical-77883bea8ef75f619d0e42467d9bcf80e3a8149ec7a2e3c66ef108f836f43078"></a>

## update property — timeouts / 099c9dcacc66 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6bc4ec389d5eb88865855e713312062d0f902a3fea1e74d3007115dff64409d5"></a>

## Next pages — timeouts / 099c9dcacc66 / 8

- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
