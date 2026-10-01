---
page_title: "xcsh_malicious_user_mitigation reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation reference."
---

# xcsh_malicious_user_mitigation reference

<a id="canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24a8a6f15d485af349bc4425291171692f322eb03c04b3b02e940490b0c6eee3"></a>

## Property reference — Property reference / 7dc19f615073 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- Property reference

<a id="canonical-83667a64d2a81a7efb9959f8d7628d547c55052f8cad1913b1a0668ef32dbee2"></a>

## Direct properties — Property reference / 7dc19f615073 / 3

<a id="canonical-cb2846d2d634b88f3e5528efe2ba2ee089e47a9182dbf95a168287f085588780"></a>

<a id="canonical-0ef8a7ef70f3a2f6af10c46835d41efedda3ddddb8f5ab281bf352ce11a0f1a6"></a>

## annotations property — Property reference / 7dc19f615073 / 4

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

<a id="canonical-454bf37e19d59436479a1688b2550b71ad2387e9f21d06169fb6d408e0732533"></a>

<a id="canonical-9113d32bc3fe06ccc3c8f033aa38d3e6cd0a3d3c2390ba26ee37c22d8514964f"></a>

## description property — Property reference / 7dc19f615073 / 5

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

<a id="canonical-efe055887b9a86966368f48cf380ddfe5164c8536c784a218a1e33f6dd8a9d50"></a>

<a id="canonical-2a7ed5a7c3b4cb5e87b039bb71723609692dab1ca84e8b894467ffc9daf3a931"></a>

## disable property — Property reference / 7dc19f615073 / 6

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

<a id="canonical-a2b486d0b0479ad7d6d4f9631dd2b299aae37ba476d8b49f3652eacf024ef839"></a>

<a id="canonical-27c5267ed173d96968b327b70fb85a82ba4bd8f84fd769dbeff6892439f57efc"></a>

## id property — Property reference / 7dc19f615073 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a3cfd9a5414ba6294420b02bb30aa515895135a095b6ed9ecb0b1186ccc23b0e"></a>

<a id="canonical-96fd9f2416c5153fa211af1332cb9c4aadba30e46f9725a41a588cac37b2262c"></a>

## labels property — Property reference / 7dc19f615073 / 8

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

- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec): complete subsection reference.

<a id="canonical-c537e776c0e06ede8e54dc139bf31b948248d17552d1cbfb647140822fd5ec1c"></a>

<a id="canonical-cdf55c930599c86f5a4a8d0be019867b34f7f2c209d885434d6e64a124b3fc04"></a>

## name property — Property reference / 7dc19f615073 / 9

Type: `"string"`. Required.

Name of the Malicious User Mitigation. Must be unique within the namespace.

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

<a id="canonical-4454a874fe1a931d372c009b3af8b2c65785d38610d8c1b42255fbe731d5c384"></a>

<a id="canonical-718ac6750e7608019f810e8c2fc873cc8ea5c162f87d4e412a35c46683cce50f"></a>

## namespace property — Property reference / 7dc19f615073 / 10

Type: `"string"`. Required.

Namespace where the Malicious User Mitigation is created.

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

- [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-955ad4ebf252780b24cb8db4b64412cd8d8b8383f4ae33233547a54e893a0716): complete subsection reference.

<a id="canonical-3215dcb3171c6fa3728fbe6708744b3b0c7c027fd10b7f4e5cdd98a4bc682215"></a>

## All schema paths — Property reference / 7dc19f615073 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--malicious_user_mitigation--reference--group-001.md#canonical-cb2846d2d634b88f3e5528efe2ba2ee089e47a9182dbf95a168287f085588780) |
| `description` | [description](resources--malicious_user_mitigation--reference--group-001.md#canonical-454bf37e19d59436479a1688b2550b71ad2387e9f21d06169fb6d408e0732533) |
| `disable` | [disable](resources--malicious_user_mitigation--reference--group-001.md#canonical-efe055887b9a86966368f48cf380ddfe5164c8536c784a218a1e33f6dd8a9d50) |
| `id` | [id](resources--malicious_user_mitigation--reference--group-001.md#canonical-a2b486d0b0479ad7d6d4f9631dd2b299aae37ba476d8b49f3652eacf024ef839) |
| `labels` | [labels](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3cfd9a5414ba6294420b02bb30aa515895135a095b6ed9ecb0b1186ccc23b0e) |
| `mitigation_type` | [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-1858ae903a9d60ca48822a5ec06673fcbd6b18a99988caaee6b2b6b5d4b39861) |
| `mitigation_type.rules` | [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-72d05b5877d48d2bc52c49c785ca21dd23b02ba4202058919de58a2f31ba07fa) |
| `mitigation_type.rules.mitigation_action` | [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-dec53abc665bdb360c4b84e1702c77eeacde47238372622605f92392b4b9f24b) |
| `mitigation_type.rules.mitigation_action.block_temporarily` | [mitigation_type.rules.mitigation_action.block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-1702a0b2f1f32baa4bb74dc25db8cb5803bd1a671228d305e3538bf0a6c0b2fc) |
| `mitigation_type.rules.mitigation_action.captcha_challenge` | [mitigation_type.rules.mitigation_action.captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-0f1916e934b3dfd11bc400503c91d49adafcbd60d1226074b7b458052e3bb0ca) |
| `mitigation_type.rules.mitigation_action.javascript_challenge` | [mitigation_type.rules.mitigation_action.javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-5143a2b6971bf4d0bf8047e2592bc068fa9dba5a83b73ae6baecbe00286206ba) |
| `mitigation_type.rules.threat_level` | [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3b75c9d8008b82e2d2534ea338a899a5f7ec9628f9d611d63f838ff78bd6c96c) |
| `mitigation_type.rules.threat_level.high` | [mitigation_type.rules.threat_level.high](resources--malicious_user_mitigation--reference--group-001.md#canonical-91f1d17966d4f8636290166c50a623d830707b69e6ab58ae14f56bb4da6b4b4d) |
| `mitigation_type.rules.threat_level.low` | [mitigation_type.rules.threat_level.low](resources--malicious_user_mitigation--reference--group-001.md#canonical-e489aa61e99f73736cc1d27bf03dc7350e3c2b01b5210deeeef1e70040464132) |
| `mitigation_type.rules.threat_level.medium` | [mitigation_type.rules.threat_level.medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-e6446b8cd7161ea021d76fe26a3476b943b704d42a84e2cfa30724f8e842e8b8) |
| `name` | [name](resources--malicious_user_mitigation--reference--group-001.md#canonical-c537e776c0e06ede8e54dc139bf31b948248d17552d1cbfb647140822fd5ec1c) |
| `namespace` | [namespace](resources--malicious_user_mitigation--reference--group-001.md#canonical-4454a874fe1a931d372c009b3af8b2c65785d38610d8c1b42255fbe731d5c384) |
| `timeouts` | [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-e7d8c81d563c08fa731ff3f71dcfa170d1f556527df74d232aa4bd7b250cccd6) |
| `timeouts.create` | [timeouts.create](resources--malicious_user_mitigation--reference--group-001.md#canonical-b0d585d69b64ddbb842a9ae1645c6139f96a15745ef88d72f46ad805ae1a8565) |
| `timeouts.delete` | [timeouts.delete](resources--malicious_user_mitigation--reference--group-001.md#canonical-e7a3b4eb388a4f649e45c743b594520b720ee65095c0378f56982740e4a21129) |
| `timeouts.read` | [timeouts.read](resources--malicious_user_mitigation--reference--group-001.md#canonical-7e51834809c34cadcf25b8d691db6ad389608ae42a3eb9b8955f0be3831527ea) |
| `timeouts.update` | [timeouts.update](resources--malicious_user_mitigation--reference--group-001.md#canonical-232a462da6d6e03c3aa9d2f5d26e101381f76ec9fae14f8ce58ca2bbff070d98) |

<a id="canonical-ff90724fe6c010303088502b0f0e68a04bad04bd56a418ca95b577637df173b9"></a>

## Next pages — Property reference / 7dc19f615073 / 12

- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-955ad4ebf252780b24cb8db4b64412cd8d8b8383f4ae33233547a54e893a0716)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57d4acdbd3d15fc6e80f526e6cf10989739f5e6605dbde413178e01517e7ac3a"></a>

## mitigation_type — mitigation_type / 84f27f88e802 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- mitigation_type

<a id="canonical-1858ae903a9d60ca48822a5ec06673fcbd6b18a99988caaee6b2b6b5d4b39861"></a>

Type: `"object"`. single nested block, Optional.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Upstream description:

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. The settings defined in malicious user
mitigation specify what mitigation actions to take for user determined to be at different threat
levels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
mitigation_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-7bf23e61110428c7260996c61c1d93cd0504ff3dbc028cddc4ad2a95a82023df"></a>

## Direct properties — mitigation_type / 84f27f88e802 / 3

- [rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73): complete subsection reference.

<a id="canonical-569472e471258ee16762e96edb270cbaeee5d588e9104373ab3cea9b67c7c724"></a>

## Next pages — mitigation_type / 84f27f88e802 / 4

- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81aea79ab430733764f87535017d53d5821ceb3ba91bab001438f93957245907"></a>

## mitigation_type.rules — mitigation_type.rules / a51b443d41a8 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- mitigation_type.rules

<a id="canonical-72d05b5877d48d2bc52c49c785ca21dd23b02ba4202058919de58a2f31ba07fa"></a>

Type: `"object"`. list nested block, Optional.

Define the threat levels and the corresponding mitigation actions to be taken.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb7de900517d2f41c67b6bf10f64239060945717375296c5e4f8c585d175b162"></a>

## Direct properties — mitigation_type.rules / a51b443d41a8 / 3

- [mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b): complete subsection reference.

- [threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949): complete subsection reference.

<a id="canonical-4604ef11b6d652567ba73c974388b2db833826498aae697b8f51d55eb4f661a9"></a>

## Next pages — mitigation_type.rules / a51b443d41a8 / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f98cfe05dd618f0bb02189d960f8e864ffc0eea5a4b47732ea22ec817358fad5"></a>

## mitigation_type.rules.mitigation_action — mitigation_type.rules.mitigation_action / 35791eb9f7cd / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- mitigation_type.rules.mitigation_action

<a id="canonical-dec53abc665bdb360c4b84e1702c77eeacde47238372622605f92392b4b9f24b"></a>

Type: `"object"`. single nested block, Optional.

Supported actions that can be taken to mitigate malicious activity from a user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_temporarily",
    "captcha_challenge"),
  validators.ConflictingObjectAttributes("block_temporarily",
    "javascript_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge",
    "javascript_challenge")}
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
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

Terraform syntax:

```terraform
mitigation_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-5d1c1be8d5329776b71a7dfe13bcef2b91888ffa87f9c231abbb9c97c985d08f"></a>

## Direct properties — mitigation_type.rules.mitigation_action / 35791eb9f7cd / 3

- [block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-5681d27777fa206d2524741433e2de2863f046b4cc3047c4f45243ed1262ed95): complete subsection reference.

- [captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-df120dbb94864fec613b74ccf8b1631d44298a3a3c52884a69d2fa78fc664759): complete subsection reference.

- [javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-de0e67e82df8c26de8478d46d05298ed6ebdc50b95fd905dc871f7f10be7c4db): complete subsection reference.

<a id="canonical-2a5070a9b7c5e72abc6100d59e34fcc29584d7439b69fc96d94dd489d9cd50ab"></a>

## Next pages — mitigation_type.rules.mitigation_action / 35791eb9f7cd / 4

- [mitigation_type.rules.mitigation_action.block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-5681d27777fa206d2524741433e2de2863f046b4cc3047c4f45243ed1262ed95)
- [mitigation_type.rules.mitigation_action.captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-df120dbb94864fec613b74ccf8b1631d44298a3a3c52884a69d2fa78fc664759)
- [mitigation_type.rules.mitigation_action.javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-de0e67e82df8c26de8478d46d05298ed6ebdc50b95fd905dc871f7f10be7c4db)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-5681d27777fa206d2524741433e2de2863f046b4cc3047c4f45243ed1262ed95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b07280ab9d07ec540365be81a632bfda83d0281c2e3749492b4794890890283"></a>

## mitigation_type.rules.mitigation_action.block_temporarily — mitigation_type.rules.mitigation_action.block_temporarily / 81317a89a5bb / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- mitigation_type.rules.mitigation_action.block_temporarily

<a id="canonical-1702a0b2f1f32baa4bb74dc25db8cb5803bd1a671228d305e3538bf0a6c0b2fc"></a>

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
block_temporarily = {}
```

<a id="canonical-e07fb0e8ab044f3be50509f3278804b09b44edea081c21da52fdbd4a2d4ff299"></a>

## Direct properties — mitigation_type.rules.mitigation_action.block_temporarily / 81317a89a5bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-540085911b6a99157c7cc6d00a5cc5272bdff47ce9da8fd9b9bad2632419a88a"></a>

## Next pages — mitigation_type.rules.mitigation_action.block_temporarily / 81317a89a5bb / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-df120dbb94864fec613b74ccf8b1631d44298a3a3c52884a69d2fa78fc664759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce7445f2b6be5476f7d9d070cb65ca59befdc32f2276a87a732839cd8ac6b0a6"></a>

## mitigation_type.rules.mitigation_action.captcha_challenge — mitigation_type.rules.mitigation_action.captcha_challenge / 03f2a5aacbca / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- mitigation_type.rules.mitigation_action.captcha_challenge

<a id="canonical-0f1916e934b3dfd11bc400503c91d49adafcbd60d1226074b7b458052e3bb0ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for captcha challenge.

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
captcha_challenge = {}
```

<a id="canonical-e46201c2ea06ff9ddd27fe4583a1595096f70a48f97965c0f5d57fdb630fcf6c"></a>

## Direct properties — mitigation_type.rules.mitigation_action.captcha_challenge / 03f2a5aacbca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78c883dd20659e62494fb29a45a3222cded4eb66b166d07fb129d70c0eaf86af"></a>

## Next pages — mitigation_type.rules.mitigation_action.captcha_challenge / 03f2a5aacbca / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-de0e67e82df8c26de8478d46d05298ed6ebdc50b95fd905dc871f7f10be7c4db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-145b86ed5a0f722bb16ba8f8c52a552ae88409213351192260d2279927de3ff6"></a>

## mitigation_type.rules.mitigation_action.javascript_challenge — mitigation_type.rules.mitigation_action.javascript_challenge / 47d16025105c / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- mitigation_type.rules.mitigation_action.javascript_challenge

<a id="canonical-5143a2b6971bf4d0bf8047e2592bc068fa9dba5a83b73ae6baecbe00286206ba"></a>

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
javascript_challenge = {}
```

<a id="canonical-3abdf91a752cdc23b98cb8d8d69b8a3784af9ac654c6ba38ad77a488c3f81c2c"></a>

## Direct properties — mitigation_type.rules.mitigation_action.javascript_challenge / 47d16025105c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b68795ce707ed6a174e67d51a81f04d803c27030f5861fa2db7c1a39e1d90ad8"></a>

## Next pages — mitigation_type.rules.mitigation_action.javascript_challenge / 47d16025105c / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-647b639c89946475992f17bb28836301729effc02229c80ff085a4d82ab5817b)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fead9cfe57101810a2f31a2d29a172b19a3760d62196bc4e52d302131cf88b8"></a>

## mitigation_type.rules.threat_level — mitigation_type.rules.threat_level / 5bf6aca58a8a / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- mitigation_type.rules.threat_level

<a id="canonical-3b75c9d8008b82e2d2534ea338a899a5f7ec9628f9d611d63f838ff78bd6c96c"></a>

Type: `"object"`. single nested block, Optional.

Threat level estimated for each user based on the user's activity and reputation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("high",
    "low"),
  validators.ConflictingObjectAttributes("high",
    "medium"),
  validators.ConflictingObjectAttributes("low",
    "medium")}
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
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

Terraform syntax:

```terraform
threat_level {
  # Configure direct properties listed below.
}
```

<a id="canonical-259d1e9e01a76d20a0c3a4f7163a5221ea05e1a5ef4b8b29dbf7f75ec0595efc"></a>

## Direct properties — mitigation_type.rules.threat_level / 5bf6aca58a8a / 3

- [high](resources--malicious_user_mitigation--reference--group-001.md#canonical-10fe7b69666edaaef09db27deda2a4044031ff22bab35041ef52721ee11a5b53): complete subsection reference.

- [low](resources--malicious_user_mitigation--reference--group-001.md#canonical-c65a2fcb9aedaa742a16f4b94421829c0d7596e401c46dbbdf918c901e60d946): complete subsection reference.

- [medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-0d6f1d940d1a4a097443dbe087c47972714f78cf0b81c53cb7120a5f405fce22): complete subsection reference.

<a id="canonical-cac022300835d06083c20e8cdeaf50384528d92753caed1843154a295abb080c"></a>

## Next pages — mitigation_type.rules.threat_level / 5bf6aca58a8a / 4

- [mitigation_type.rules.threat_level.high](resources--malicious_user_mitigation--reference--group-001.md#canonical-10fe7b69666edaaef09db27deda2a4044031ff22bab35041ef52721ee11a5b53)
- [mitigation_type.rules.threat_level.low](resources--malicious_user_mitigation--reference--group-001.md#canonical-c65a2fcb9aedaa742a16f4b94421829c0d7596e401c46dbbdf918c901e60d946)
- [mitigation_type.rules.threat_level.medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-0d6f1d940d1a4a097443dbe087c47972714f78cf0b81c53cb7120a5f405fce22)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-10fe7b69666edaaef09db27deda2a4044031ff22bab35041ef52721ee11a5b53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-deb179f75060e7ba2d55e7a9d8a4c8d79efbe0090ed9c95cf385ac009490904b"></a>

## mitigation_type.rules.threat_level.high — mitigation_type.rules.threat_level.high / cc8591331f68 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- mitigation_type.rules.threat_level.high

<a id="canonical-91f1d17966d4f8636290166c50a623d830707b69e6ab58ae14f56bb4da6b4b4d"></a>

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
high = {}
```

<a id="canonical-5e5a61579d8bc44000fb37e9971210bbe99127afde86b81e11418defa9741504"></a>

## Direct properties — mitigation_type.rules.threat_level.high / cc8591331f68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd7f46aab30e0280180fbb2220ec30855241489096f72698a7e91cbae06524d2"></a>

## Next pages — mitigation_type.rules.threat_level.high / cc8591331f68 / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-c65a2fcb9aedaa742a16f4b94421829c0d7596e401c46dbbdf918c901e60d946"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a12ba5dead29dec1888a4285a38d536d320410a43355e82eeaae4d842a3ca5c"></a>

## mitigation_type.rules.threat_level.low — mitigation_type.rules.threat_level.low / 6f221d4867ad / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- mitigation_type.rules.threat_level.low

<a id="canonical-e489aa61e99f73736cc1d27bf03dc7350e3c2b01b5210deeeef1e70040464132"></a>

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
low = {}
```

<a id="canonical-40f330fbb53cf460e8b4885f4111900e6036a38edb44d8f9c9493688e78a1697"></a>

## Direct properties — mitigation_type.rules.threat_level.low / 6f221d4867ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2ed36b2bdc3b4462f8c2f75cfd203515c6cf784c8f01c36f87e02f59b11df73"></a>

## Next pages — mitigation_type.rules.threat_level.low / 6f221d4867ad / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-0d6f1d940d1a4a097443dbe087c47972714f78cf0b81c53cb7120a5f405fce22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-108d2f6df3229b13a27ad62a0acbdb6bf0a25a5d7771cbe72364511b3ffd8b32"></a>

## mitigation_type.rules.threat_level.medium — mitigation_type.rules.threat_level.medium / fcd632568083 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-18837081a8150c1ba590aedacdb79d81919a9b235f1736af59596cf7c5821fec)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-a3b833d323fdab6d7a2b4b6c163974ddac79a90e00f89f06aafa62927ff9ff73)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- mitigation_type.rules.threat_level.medium

<a id="canonical-e6446b8cd7161ea021d76fe26a3476b943b704d42a84e2cfa30724f8e842e8b8"></a>

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
medium = {}
```

<a id="canonical-22b5e72dc5c9636f9cc4a198dafa4460936304d45baa1fcbee8136947c15f2af"></a>

## Direct properties — mitigation_type.rules.threat_level.medium / fcd632568083 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5e900fafaf89c5a3758421b399e8c9f87dcaf0f76da4b2df9ee803289dfd130"></a>

## Next pages — mitigation_type.rules.threat_level.medium / fcd632568083 / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-fd989a8b1386aced4e911ba6e82cb0556924f7dcce33a1d9327f41facc677949)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)

<a id="canonical-955ad4ebf252780b24cb8db4b64412cd8d8b8383f4ae33233547a54e893a0716"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e9109bf3c258ce41129b9975fe3c6eae34a231680590a8d878a505b4063c559"></a>

## timeouts — timeouts / 9dbce0c7628a / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- timeouts

<a id="canonical-e7d8c81d563c08fa731ff3f71dcfa170d1f556527df74d232aa4bd7b250cccd6"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5f50d3b2a5903bbbf1ab481e302012a7f14766f8b2d73cc8016109e7a35c7a7"></a>

## Direct properties — timeouts / 9dbce0c7628a / 3

<a id="canonical-b0d585d69b64ddbb842a9ae1645c6139f96a15745ef88d72f46ad805ae1a8565"></a>

<a id="canonical-724ed66a6c47150ede8d337b3b12c16ec9e4388ae8e304667737d16a5c65924b"></a>

## create property — timeouts / 9dbce0c7628a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e7a3b4eb388a4f649e45c743b594520b720ee65095c0378f56982740e4a21129"></a>

<a id="canonical-49e27c76c8cdfff0e63d1c3647f3cb8f2fedfd6b5ee9fc057e829e8ad7b054bb"></a>

## delete property — timeouts / 9dbce0c7628a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7e51834809c34cadcf25b8d691db6ad389608ae42a3eb9b8955f0be3831527ea"></a>

<a id="canonical-b80ea1db6654f1955706c52b1152b70daf4d177542d58a7d240c2138f35d8d9f"></a>

## read property — timeouts / 9dbce0c7628a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-232a462da6d6e03c3aa9d2f5d26e101381f76ec9fae14f8ce58ca2bbff070d98"></a>

<a id="canonical-604d0306a82c979d0e4312ffc4ecfd6eea78ddd928c530d1acaaf7377ef169b0"></a>

## update property — timeouts / 9dbce0c7628a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f2e512f2a0648655d6f2f661f9dd838988e01d8e022ccdfc3b16caaa2917f46c"></a>

## Next pages — timeouts / 9dbce0c7628a / 8

- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db)
