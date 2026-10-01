---
page_title: "xcsh_cloud_user_account reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account reference."
---

# xcsh_cloud_user_account reference

<a id="canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2682d46788f117495e2a2493585279f268e3f3f44f5746c75ca9cc1f3452fba0"></a>

## Property reference — Property reference / 8533095ed6b2 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- Property reference

<a id="canonical-aee3c912ef6b4d92b93cdcb7c7df19ecb8abdb0939dfe630a334ad4c8ae094bf"></a>

## Direct properties — Property reference / 8533095ed6b2 / 3

<a id="canonical-58bf41d14cfa82316a1cd85adf154fa8ab4b3a463a14e595d2c0ffc4096b8829"></a>

<a id="canonical-ebe2190c77de5dca474032af97ecfcc98818d602924c486bbe029bf21601d8da"></a>

## annotations property — Property reference / 8533095ed6b2 / 4

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

- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45): complete subsection reference.

<a id="canonical-d304d29b2f5e11fbbb6dd8a558bf570b7a40573c5022f46159901842bec85383"></a>

<a id="canonical-655303a8c48932dc0cc242c3768cb6cbeed99f49ef64b0db5b5376771bd805c3"></a>

## description property — Property reference / 8533095ed6b2 / 5

Type: `"string"`. Computed.

Description of the CloudUserAccount.

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

<a id="canonical-262b20a223e44933770d4f43a3cbe30beec14bbf716b2db88eb9eaa8bf73124d"></a>

<a id="canonical-bf9dde05377e4c8845b662a074e449814581d24e15007f94ffc44ed80676d578"></a>

## id property — Property reference / 8533095ed6b2 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-271535f5f720d517d91441a494bd9b92c6b25800358ae242249224e0bbd33d9a"></a>

<a id="canonical-14f6784cf5c824cd975160b10483361418a54560e4bbf3d19fdf0d86570f848e"></a>

## labels property — Property reference / 8533095ed6b2 / 7

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

<a id="canonical-32aff63849652bbc01708cd8c6d2ee0c07e6d42b7294b06030900afd0098b98f"></a>

<a id="canonical-d362d3f7e29b21475363739e1b778ef7e6cd2fd0046af882a84768e136b634c5"></a>

## name property — Property reference / 8533095ed6b2 / 8

Type: `"string"`. Required.

Name of the CloudUserAccount.

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

<a id="canonical-2ec2d9564252198d55bbedbd11444c45d605e81abe313dc7befa75f4f599c81f"></a>

<a id="canonical-a4f6ab185f1269c1701c465572cd894f6949059d555c28852049e625a8937275"></a>

## namespace property — Property reference / 8533095ed6b2 / 9

Type: `"string"`. Required.

Namespace where the CloudUserAccount exists.

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

<a id="canonical-3b587b9fe375ee7fd5dd8773246dd00a148ce348acaa0e628991aa94e924e9ed"></a>

## All schema paths — Property reference / 8533095ed6b2 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_user_account--reference--group-001.md#canonical-58bf41d14cfa82316a1cd85adf154fa8ab4b3a463a14e595d2c0ffc4096b8829) |
| `aws_provider` | [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-715e3006dbfc5ab7ce7351411dd292f67f8c629cc779e929a9654a95c69d898d) |
| `aws_provider.aws_account_number` | [aws_provider.aws_account_number](data-sources--cloud_user_account--reference--group-001.md#canonical-ee58bf05b3d0d73f96434c60935082602be23a41e2f04ecb4ab679b15641b44d) |
| `aws_provider.aws_assume_role` | [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-ca4a001b467d9ebde6e85aa0abe092bb28f17d168b78f300eb2cd639a85e30a7) |
| `aws_provider.aws_assume_role.custom_external_id` | [aws_provider.aws_assume_role.custom_external_id](data-sources--cloud_user_account--reference--group-001.md#canonical-0e90ad7ed42e3d41268fb74b228f9167ed14447da9a6e9719ab31fe2167a08e5) |
| `aws_provider.aws_assume_role.duration_seconds` | [aws_provider.aws_assume_role.duration_seconds](data-sources--cloud_user_account--reference--group-001.md#canonical-b94ee4ec31ed320084af264cae256614fae41008dff28552e68ef25a2c3b9e64) |
| `aws_provider.aws_assume_role.external_id_is_optional` | [aws_provider.aws_assume_role.external_id_is_optional](data-sources--cloud_user_account--reference--group-001.md#canonical-0de3b92a1e0634a4c327d15cff3e7b699a1a43dde7e289f36d1a6564e1adfa9e) |
| `aws_provider.aws_assume_role.external_id_is_tenant_id` | [aws_provider.aws_assume_role.external_id_is_tenant_id](data-sources--cloud_user_account--reference--group-001.md#canonical-b73e757eff81659f5b3f6065a67b7e54ffaa408c9b32b301f588c9768fb84361) |
| `aws_provider.aws_assume_role.role_arn` | [aws_provider.aws_assume_role.role_arn](data-sources--cloud_user_account--reference--group-001.md#canonical-0fd6d72ee207f02f25a78e0397b8542ff6794b9f15f8b9d8b5b866584428858d) |
| `aws_provider.aws_assume_role.session_name` | [aws_provider.aws_assume_role.session_name](data-sources--cloud_user_account--reference--group-001.md#canonical-3d527c44f2b9dc08919448878e49f8704c43a98f157993e39e01c1467bd94b25) |
| `aws_provider.aws_assume_role.session_tags` | [aws_provider.aws_assume_role.session_tags](data-sources--cloud_user_account--reference--group-001.md#canonical-ac5150e4311bd537d7e82672f3777967bf96d5e7faeb9fc460cdae9b49e9448b) |
| `aws_provider.aws_secret_key` | [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-f2f1ac48310a68e02280ccbb499acaf83d06355fb47a2b9e40db785942be7239) |
| `aws_provider.aws_secret_key.access_key` | [aws_provider.aws_secret_key.access_key](data-sources--cloud_user_account--reference--group-001.md#canonical-e4e8a254d7a4b91a9436512bac161f1d09ba26bcf8c27beecc6a7cfdbadec9c2) |
| `aws_provider.aws_secret_key.secret_key` | [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-6d684de748a49bfe14b0454d6b0bccdb82e8f67ff0c451cefa9c434e139666f1) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-dc0ea03991ca60f800d9fcadaa6ff5048f77e172fb9388cb8eee182657c73968) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-8ceb8aa5c2854f1d83e001bf9c059e3f0ba78f01124a0d687660113c5f1a4ef1) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location](data-sources--cloud_user_account--reference--group-001.md#canonical-71d781c295cc1fb43f20f5357fe1bb3bc17b5fc99b7ee9de91eeb933bcdf3828) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-50fb741a9be92aa20424798c8e85a015fd60085de9d8e59949a026df0b00e529) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info` | [aws_provider.aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-bca62256d87476ced878889085e4300746f5b987b988686450640ca093cafb12) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref](data-sources--cloud_user_account--reference--group-001.md#canonical-e55490684f722af81000a7eb81cefd23517a5629089d911316271b0c61613f3c) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.url](data-sources--cloud_user_account--reference--group-001.md#canonical-8c469f384de9fc9921bc50fd8e2c11eef162c9286225a91c9a37cbf4212a32d9) |
| `description` | [description](data-sources--cloud_user_account--reference--group-001.md#canonical-d304d29b2f5e11fbbb6dd8a558bf570b7a40573c5022f46159901842bec85383) |
| `id` | [id](data-sources--cloud_user_account--reference--group-001.md#canonical-262b20a223e44933770d4f43a3cbe30beec14bbf716b2db88eb9eaa8bf73124d) |
| `labels` | [labels](data-sources--cloud_user_account--reference--group-001.md#canonical-271535f5f720d517d91441a494bd9b92c6b25800358ae242249224e0bbd33d9a) |
| `name` | [name](data-sources--cloud_user_account--reference--group-001.md#canonical-32aff63849652bbc01708cd8c6d2ee0c07e6d42b7294b06030900afd0098b98f) |
| `namespace` | [namespace](data-sources--cloud_user_account--reference--group-001.md#canonical-2ec2d9564252198d55bbedbd11444c45d605e81abe313dc7befa75f4f599c81f) |

<a id="canonical-868cd81500442cedea093865db80219f0144ec54768fdd2227c30e2ece906926"></a>

## Next pages — Property reference / 8533095ed6b2 / 11

- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41c900b3f3aa5d410426e79fc5550bb18174ab045192729fd74944e7b04d92c2"></a>

## aws_provider — aws_provider / e4a46084d967 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- aws_provider

<a id="canonical-715e3006dbfc5ab7ce7351411dd292f67f8c629cc779e929a9654a95c69d898d"></a>

Type: `"single"`. Computed.

Configuration parameter for aws provider.

Upstream description:

Create AWS Provider Type.

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

<a id="canonical-f80eace5a2d60f04091def00a4a73ecbf0d7a99fb0e896c6616b4355a38fccc5"></a>

## Direct properties — aws_provider / e4a46084d967 / 3

<a id="canonical-ee58bf05b3d0d73f96434c60935082602be23a41e2f04ecb4ab679b15641b44d"></a>

<a id="canonical-e3e16779ccca5f75d5d67993f1c7d88f679b46746db112a9b8d6dfe875cb4f51"></a>

## aws_account_number property — aws_provider / e4a46084d967 / 4

Type: `"string"`. Computed.

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

- [aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625): complete subsection reference.

- [aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da): complete subsection reference.

<a id="canonical-78ded9f13d3fee28230676392a98cec06eb09ff154d3cf3f7bc844444d15c837"></a>

## Next pages — aws_provider / e4a46084d967 / 5

- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b827b30f161d0a9ab41f9e556ec45f82f0fc13c8f17d15364bf1feaade9b3096"></a>

## aws_provider.aws_assume_role — aws_provider.aws_assume_role / 48fd10249cb8 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- aws_provider.aws_assume_role

<a id="canonical-ca4a001b467d9ebde6e85aa0abe092bb28f17d168b78f300eb2cd639a85e30a7"></a>

Type: `"single"`. Computed.

AWS Assume Role to Handle Delegated Access.

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

<a id="canonical-34bb5d18d5f100fe428b2884d38c2afdc807a94c2febc0fe4eb381029cdd658f"></a>

## Direct properties — aws_provider.aws_assume_role / 48fd10249cb8 / 3

<a id="canonical-0e90ad7ed42e3d41268fb74b228f9167ed14447da9a6e9719ab31fe2167a08e5"></a>

<a id="canonical-9b35770d48ecbca83a7e193f211d4fcd75baff650f5d47fca99f87c8bbe50276"></a>

## custom_external_id property — aws_provider.aws_assume_role / 48fd10249cb8 / 4

Type: `"string"`. Computed.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

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

<a id="canonical-b94ee4ec31ed320084af264cae256614fae41008dff28552e68ef25a2c3b9e64"></a>

<a id="canonical-95bc1da8bae98db147f60b2305e1875634fadab63aecae47dff3257f82eb7114"></a>

## duration_seconds property — aws_provider.aws_assume_role / 48fd10249cb8 / 5

Type: `"number"`. Computed.

The duration, in seconds of the role session.

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

- [external_id_is_optional](data-sources--cloud_user_account--reference--group-001.md#canonical-f3c45f1af2b931c312946a6d46a57ded67fce1da515f4a98c129a38b610289f1): complete subsection reference.

- [external_id_is_tenant_id](data-sources--cloud_user_account--reference--group-001.md#canonical-2e8f6b536381c481eb89dc1ee2d17ccd163b7082b9ce3c617705f49ae573c461): complete subsection reference.

<a id="canonical-0fd6d72ee207f02f25a78e0397b8542ff6794b9f15f8b9d8b5b866584428858d"></a>

<a id="canonical-88a96a3b5ef537748369aed3df5e6152de8090934be276854437192fbcb05d23"></a>

## role_arn property — aws_provider.aws_assume_role / 48fd10249cb8 / 6

Type: `"string"`. Computed.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

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

<a id="canonical-3d527c44f2b9dc08919448878e49f8704c43a98f157993e39e01c1467bd94b25"></a>

<a id="canonical-64f7e2ad7a458c324679a4573e4353fb68f00b8cf3d8fa75539677201f309de3"></a>

## session_name property — aws_provider.aws_assume_role / 48fd10249cb8 / 7

Type: `"string"`. Computed.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

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

<a id="canonical-ac5150e4311bd537d7e82672f3777967bf96d5e7faeb9fc460cdae9b49e9448b"></a>

<a id="canonical-f2def2452571051d2a7e7741763ffa454db762c5c994b0ef471e82799bc90a54"></a>

## session_tags property — aws_provider.aws_assume_role / 48fd10249cb8 / 8

Type: `["map", "string"]`. Computed.

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

<a id="canonical-6fae4584612cfdb071c188fe9552e4ba12a10efeac49054ecc9c068b5bf107b4"></a>

## Next pages — aws_provider.aws_assume_role / 48fd10249cb8 / 9

- [aws_provider.aws_assume_role.external_id_is_optional](data-sources--cloud_user_account--reference--group-001.md#canonical-f3c45f1af2b931c312946a6d46a57ded67fce1da515f4a98c129a38b610289f1)
- [aws_provider.aws_assume_role.external_id_is_tenant_id](data-sources--cloud_user_account--reference--group-001.md#canonical-2e8f6b536381c481eb89dc1ee2d17ccd163b7082b9ce3c617705f49ae573c461)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-f3c45f1af2b931c312946a6d46a57ded67fce1da515f4a98c129a38b610289f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ef1b88bb3d89394c0efa367a2a6671cd229da5bb8939b27953598e8c98b0273"></a>

## aws_provider.aws_assume_role.external_id_is_optional — aws_provider.aws_assume_role.external_id_is_optional / 4f150e17bee2 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625)
- aws_provider.aws_assume_role.external_id_is_optional

<a id="canonical-0de3b92a1e0634a4c327d15cff3e7b699a1a43dde7e289f36d1a6564e1adfa9e"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fba57897d59abcd65fabb2c92634ba8de6416b02ae85dffe7559c9d14cb08102"></a>

## Direct properties — aws_provider.aws_assume_role.external_id_is_optional / 4f150e17bee2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e8bc331943a8ba3b918ecb5a05a45ed76f9990aed7fa0ce93b2e78c2f096fed"></a>

## Next pages — aws_provider.aws_assume_role.external_id_is_optional / 4f150e17bee2 / 4

- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-2e8f6b536381c481eb89dc1ee2d17ccd163b7082b9ce3c617705f49ae573c461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-576af698e8c1ca7092dfeb7ce96f2708470efc1cd5632c800a804fd6b7f9a965"></a>

## aws_provider.aws_assume_role.external_id_is_tenant_id — aws_provider.aws_assume_role.external_id_is_tenant_id / f995d252dbaf / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625)
- aws_provider.aws_assume_role.external_id_is_tenant_id

<a id="canonical-b73e757eff81659f5b3f6065a67b7e54ffaa408c9b32b301f588c9768fb84361"></a>

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

<a id="canonical-471e0b076b2fbebb51907a62905ab3a59d2b4d4e328456c0e0f434431e85e9cc"></a>

## Direct properties — aws_provider.aws_assume_role.external_id_is_tenant_id / f995d252dbaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-248667955b6760a2e86bf1f46c2181a73ac0a2def7d6e61515e8ad478e3a02aa"></a>

## Next pages — aws_provider.aws_assume_role.external_id_is_tenant_id / f995d252dbaf / 4

- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-a8d7593e962ff42fb5cd499a94c47e735d96531f0dfebe44fd1d4021a25d8625)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2db445addee0602fc3a66db5768f7a916b0569f47c9ebad15cd50d3d9440fb78"></a>

## aws_provider.aws_secret_key — aws_provider.aws_secret_key / fac3ffba99d2 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- aws_provider.aws_secret_key

<a id="canonical-f2f1ac48310a68e02280ccbb499acaf83d06355fb47a2b9e40db785942be7239"></a>

Type: `"single"`. Computed.

AWS Programmatic Access Credentials type.

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

<a id="canonical-9af565c68805fa4d8aed2cce8e7ebc92404b60cca23b2637caebd8069e7e1f2a"></a>

## Direct properties — aws_provider.aws_secret_key / fac3ffba99d2 / 3

<a id="canonical-e4e8a254d7a4b91a9436512bac161f1d09ba26bcf8c27beecc6a7cfdbadec9c2"></a>

<a id="canonical-9650c288e1553cb6884819920f2836f069e1ac2bb9b1c242769683e99cbe50d2"></a>

## access_key property — aws_provider.aws_secret_key / fac3ffba99d2 / 4

Type: `"string"`. Computed.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

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

- [secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b): complete subsection reference.

<a id="canonical-0d64dec1fc7e1366fa08c62f52227ce8545e99484988eecd9201e62d5cb3ea76"></a>

## Next pages — aws_provider.aws_secret_key / fac3ffba99d2 / 5

- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-426a49cdbc14fe7f14a56dbe9a1d9df79cc0bdc4d3e960cd1c3aecbe4c944c9f"></a>

## aws_provider.aws_secret_key.secret_key — aws_provider.aws_secret_key.secret_key / 70ae9872ed1a / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da)
- aws_provider.aws_secret_key.secret_key

<a id="canonical-6d684de748a49bfe14b0454d6b0bccdb82e8f67ff0c451cefa9c434e139666f1"></a>

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

<a id="canonical-8e5f04d57c896ca20a4d079edbf6130fe21ef169b34f321b64ac114d9d5c5763"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key / 70ae9872ed1a / 3

- [blindfold_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-65481a9e20886657a6165e8d572faba00fbddee404a32eb4e19b11bb19e86456): complete subsection reference.

- [clear_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-c71a0cbd744686b1331014ccc4ed5f9c3c3b7b78ca66104f151f3db544b65a7e): complete subsection reference.

<a id="canonical-ef3ab8668c722480f4460524c009e61f0130f0feaceb0bac5065401441239b35"></a>

## Next pages — aws_provider.aws_secret_key.secret_key / 70ae9872ed1a / 4

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-65481a9e20886657a6165e8d572faba00fbddee404a32eb4e19b11bb19e86456)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-c71a0cbd744686b1331014ccc4ed5f9c3c3b7b78ca66104f151f3db544b65a7e)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-65481a9e20886657a6165e8d572faba00fbddee404a32eb4e19b11bb19e86456"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e553d007ff1f0d88c7e9a86d3f3201abf1d24abe864f3cf984a2b4be59485ab2"></a>

## aws_provider.aws_secret_key.secret_key.blindfold_secret_info — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da)
- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b)
- aws_provider.aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-dc0ea03991ca60f800d9fcadaa6ff5048f77e172fb9388cb8eee182657c73968"></a>

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

<a id="canonical-49fef5881f1bd30d5a7e6578aa53b37078f678e010ee39156afe9d1ecfdf211e"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 3

<a id="canonical-8ceb8aa5c2854f1d83e001bf9c059e3f0ba78f01124a0d687660113c5f1a4ef1"></a>

<a id="canonical-d69f8d1594f52a9dc6c80f2396ecc9f8bc0397f085f618ae1ba89c508a5d2eec"></a>

## decryption_provider property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 4

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

<a id="canonical-71d781c295cc1fb43f20f5357fe1bb3bc17b5fc99b7ee9de91eeb933bcdf3828"></a>

<a id="canonical-99551e0435d9495a8fc98d7560c4f47592370c0a31465d2bf92b67bc559a7802"></a>

## location property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 5

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

<a id="canonical-50fb741a9be92aa20424798c8e85a015fd60085de9d8e59949a026df0b00e529"></a>

<a id="canonical-2e77a9dd3aa397b0da01d48cc7208370ff6c1a6e519c4d57145f2652d80c16c5"></a>

## store_provider property — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 6

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

<a id="canonical-bc87feb7aa415ff67465caf482c84ecbb467c3740964da6792ba191384d43911"></a>

## Next pages — aws_provider.aws_secret_key.secret_key.blindfold_secret_info / e4f8b9a3a0e7 / 7

- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-c71a0cbd744686b1331014ccc4ed5f9c3c3b7b78ca66104f151f3db544b65a7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e5812883833f76af09dee2ecf1531c9c385da5f9290675590fdb819eaec91f1"></a>

## aws_provider.aws_secret_key.secret_key.clear_secret_info — aws_provider.aws_secret_key.secret_key.clear_secret_info / 2c7da5b9911d / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-12950f6c794fd0a06e4903ccb92de67ed85b8b864def70fea269d3a1fd332f45)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-ceb7a2fe2b8a2aa8c3e1582f80b83603c95f7bb666377b5025249b8ba7c0a7da)
- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="canonical-bca62256d87476ced878889085e4300746f5b987b988686450640ca093cafb12"></a>

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

<a id="canonical-4542b6296e58b1cbcd9eb042756b825d64acfffa607f7f71f3b0884d06b6f505"></a>

## Direct properties — aws_provider.aws_secret_key.secret_key.clear_secret_info / 2c7da5b9911d / 3

<a id="canonical-e55490684f722af81000a7eb81cefd23517a5629089d911316271b0c61613f3c"></a>

<a id="canonical-7383259ab6d7cea7c00bdaebb4ca13c5001158e9196f440822d00e27d52c384c"></a>

## provider_ref property — aws_provider.aws_secret_key.secret_key.clear_secret_info / 2c7da5b9911d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8c469f384de9fc9921bc50fd8e2c11eef162c9286225a91c9a37cbf4212a32d9"></a>

<a id="canonical-fbd3fd8a37bbf5a382270b509e2e4bb0be6e94be3997ca4a48bff27eeafa4f54"></a>

## url property — aws_provider.aws_secret_key.secret_key.clear_secret_info / 2c7da5b9911d / 5

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

<a id="canonical-6d3cf84d7cb4b6e79659566c37be21a8b95f5a8903ae3bf7f0fbc95ffef80680"></a>

## Next pages — aws_provider.aws_secret_key.secret_key.clear_secret_info / 2c7da5b9911d / 6

- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-a16ce34cf77a0c58c3f3a244378e33485f3fa0b9df3f4fe3537c6f9c260b858b)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
