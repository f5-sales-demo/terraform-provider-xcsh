---
page_title: "xcsh_ike_phase1_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile reference."
---

# xcsh_ike_phase1_profile reference

<a id="canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c91ecc3c1736c26183c56fee790ac6ead4d92165fdaef16a47c922aca66c3eda"></a>

## Property reference — Property reference / fc17ac3396c7 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- Property reference

<a id="canonical-31db3ba0758db75e30038925165713f10cac72efc53b26a85a7c7b9f169e558b"></a>

## Direct properties — Property reference / fc17ac3396c7 / 3

<a id="canonical-899ce136ea268909e19dc4dbc78c87f7213938cf83ed03ec517cd3eb21926f4f"></a>

<a id="canonical-fb503f5b2a62ab77695386eafdcd472d13825a9aa4280984b6ca8a5f71b187e2"></a>

## annotations property — Property reference / fc17ac3396c7 / 4

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

<a id="canonical-a6013a0030bdb8d0053d6d3fec154f777b1dc856711f086e126fde3e6af6a4cc"></a>

<a id="canonical-6fcb9f6a8757f491e1a5d3ed3b6ca062c9e625ac0b065f2579821501511cd29e"></a>

## authentication_algos property — Property reference / fc17ac3396c7 / 5

Type: `["list", "string"]`. Computed.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm
encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-493b923c5508091d60744ce6fe69d474710dbc60679b717423ed8d9c572bee5a"></a>

<a id="canonical-1d58b45cd5672106ddcc93691b00322ed42a02783dbac90994edb14f528c5658"></a>

## description property — Property reference / fc17ac3396c7 / 6

Type: `"string"`. Computed.

Description of the IKEPhase1Profile.

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

<a id="canonical-11bbe163ae44aa1c359b801e50e8c4c9be0f47494a9211d739484c1971a03135"></a>

<a id="canonical-1dac71ceb088a87adc880a057a6e7c54d9229b1ed5b8a19fa8a7487b6627c4f5"></a>

## dh_group property — Property reference / fc17ac3396c7 / 7

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-8044ebc1481e10d7555c9327a0a90a5b17be847737ea0a2259a189cd97961034"></a>

<a id="canonical-e0281963e1e836fde6b16e86429c8625998222df86c32ab928dc9a14e4962260"></a>

## encryption_algos property — Property reference / fc17ac3396c7 / 8

Type: `["list", "string"]`. Computed.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-4dcef52e0981e2205cd29ebb57ae03a8840f152fccc511dabd009b80dc725475"></a>

<a id="canonical-116d4b29cfe14476b297206f7e556f05b11d3d85ce3cc7d1a714044f9e6c996f"></a>

## id property — Property reference / fc17ac3396c7 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1b373993fe6eaa5d9e712a3d6e3bfa241321616ca2f04732ef1f24fca264721c): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-f86d6add4eec8c679321cd96a5ba2693a6588034563641d3d751260483a03f73): complete subsection reference.

<a id="canonical-31ac83b37cd85d8134ed824263393459064ca1b5acb613d14ad8a6d89b55a77b"></a>

<a id="canonical-0b2d1bb4001c9af6313136eee869870cfa188690c9d801d304a045427ed5f16d"></a>

## labels property — Property reference / fc17ac3396c7 / 10

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

<a id="canonical-77fc57cb15ea3a68efa1a2b15743c1978af9dc94b2ccddf625e6fed6fda5acc4"></a>

<a id="canonical-82f997ca68d1032892657ad758998625e98d9bb7f5d9dd50e80c9326a070cc8c"></a>

## name property — Property reference / fc17ac3396c7 / 11

Type: `"string"`. Required.

Name of the IKEPhase1Profile.

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

<a id="canonical-60d834e1afe6a2776bb886e704501571cabcf0851a2861c580b9ba116e6f37f4"></a>

<a id="canonical-202124ae7df36b5e23cc6d48209569a019fbc2d21314eeb8f439946866ab514c"></a>

## namespace property — Property reference / fc17ac3396c7 / 12

Type: `"string"`. Required.

Namespace where the IKEPhase1Profile exists.

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

<a id="canonical-02032d7107ba1b9c03c62dcf3521d7a0ddb977a0e054d3c31b0bd6d3041a4320"></a>

<a id="canonical-299730bd4e2bde83d52ed478108655ea423e93b8873cec3d13336e249704f0bb"></a>

## prf property — Property reference / fc17ac3396c7 / 13

Type: `["list", "string"]`. Computed.

\[Enum: PRF\_DEFAULT|PRFSHA256|PRFSHA384|PRFSHA512\] PseudoRandomFunction. Select
PseudoRandomFunction for IKE SA. Possible values are \`PRF\_DEFAULT\`, \`PRFSHA256\`, \`PRFSHA384\`,
\`PRFSHA512\`. Defaults to \`PRF\_DEFAULT\`.

Upstream description:

Select PseudoRandomFunction for IKE SA.

Receipt-pinned upstream constraints:

```json
{
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

- [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-89ceae40686c00b1e6d586841e1058cd1b0940b8e7cac4233e283853125268b2): complete subsection reference.

- [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-c86d108dd9d41d82e437353de084c84587ccba24264612d5c5dc8004a053a6c0): complete subsection reference.

- [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-6ca185d07dc5cb21643520089952ad9bc4f288fb1cb7d4aa59cccd1a662d4a25): complete subsection reference.

- [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-e47b7caf259b20fa0f5838daf0c8610633a7c9b16e9a070cdf17804e27072a86): complete subsection reference.

<a id="canonical-c2ef2f73c99298f4ad56da06ea5e281251030e5046560d7169c61f39fdba1548"></a>

## All schema paths — Property reference / fc17ac3396c7 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike_phase1_profile--reference--group-001.md#canonical-899ce136ea268909e19dc4dbc78c87f7213938cf83ed03ec517cd3eb21926f4f) |
| `authentication_algos` | [authentication_algos](data-sources--ike_phase1_profile--reference--group-001.md#canonical-a6013a0030bdb8d0053d6d3fec154f777b1dc856711f086e126fde3e6af6a4cc) |
| `description` | [description](data-sources--ike_phase1_profile--reference--group-001.md#canonical-493b923c5508091d60744ce6fe69d474710dbc60679b717423ed8d9c572bee5a) |
| `dh_group` | [dh_group](data-sources--ike_phase1_profile--reference--group-001.md#canonical-11bbe163ae44aa1c359b801e50e8c4c9be0f47494a9211d739484c1971a03135) |
| `encryption_algos` | [encryption_algos](data-sources--ike_phase1_profile--reference--group-001.md#canonical-8044ebc1481e10d7555c9327a0a90a5b17be847737ea0a2259a189cd97961034) |
| `id` | [id](data-sources--ike_phase1_profile--reference--group-001.md#canonical-4dcef52e0981e2205cd29ebb57ae03a8840f152fccc511dabd009b80dc725475) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-720ccc0ae76eab392fa99f7625adf09ca7f76306a8375d3a707ff2799db54cd6) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-c126e07fc3dfdf0444df8dab4efacf6a15a9b8deb664e9f9149859f65fde528e) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-58b178655a495d4d8282921323d01ffa982e49c5e6c4bdb2b61c411d0dc4d95f) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-240c824a31f0cd21088acb4f1c6985355cecab2f101fae60c1e7fa9edc6f53de) |
| `labels` | [labels](data-sources--ike_phase1_profile--reference--group-001.md#canonical-31ac83b37cd85d8134ed824263393459064ca1b5acb613d14ad8a6d89b55a77b) |
| `name` | [name](data-sources--ike_phase1_profile--reference--group-001.md#canonical-77fc57cb15ea3a68efa1a2b15743c1978af9dc94b2ccddf625e6fed6fda5acc4) |
| `namespace` | [namespace](data-sources--ike_phase1_profile--reference--group-001.md#canonical-60d834e1afe6a2776bb886e704501571cabcf0851a2861c580b9ba116e6f37f4) |
| `prf` | [prf](data-sources--ike_phase1_profile--reference--group-001.md#canonical-02032d7107ba1b9c03c62dcf3521d7a0ddb977a0e054d3c31b0bd6d3041a4320) |
| `reauth_disabled` | [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3d9a9d804764976d971fc4e61652d75e6e854046be745fd74406f4998972205a) |
| `reauth_timeout_days` | [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-e6aaf7f795da55f7e05dbdb9a879b055476847afab7137a26cc3300e6b3df829) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-cdab320fba8c7c3adb19296d0861404755e25485245d5105c14ecf889c9b004d) |
| `reauth_timeout_hours` | [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-a70c23d5e0b0829958ccae393b18e7f933fcc7699e343a2b2c3f937da8a20783) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-e86fe5eda6af196e9030b6ad57b42b94744afdbb12fb4fd48650688311bf6c60) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-6327285e6f8dcdb7c517d68ed9b1a3da046a1198c4685ef6992770af892c030d) |

<a id="canonical-5c34317a19bb43a60e1906d3c2b011e45315b76222ff5e97aa8177303f503ba8"></a>

## Next pages — Property reference / fc17ac3396c7 / 15

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1b373993fe6eaa5d9e712a3d6e3bfa241321616ca2f04732ef1f24fca264721c)
- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-f86d6add4eec8c679321cd96a5ba2693a6588034563641d3d751260483a03f73)
- [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-89ceae40686c00b1e6d586841e1058cd1b0940b8e7cac4233e283853125268b2)
- [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-c86d108dd9d41d82e437353de084c84587ccba24264612d5c5dc8004a053a6c0)
- [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-6ca185d07dc5cb21643520089952ad9bc4f288fb1cb7d4aa59cccd1a662d4a25)
- [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-e47b7caf259b20fa0f5838daf0c8610633a7c9b16e9a070cdf17804e27072a86)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-1b373993fe6eaa5d9e712a3d6e3bfa241321616ca2f04732ef1f24fca264721c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33135351ddb8422fa2a27935681120177e740007cec0b33965d2615f193355d5"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / cc1c9b157391 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- ike_keylifetime_hours

<a id="canonical-720ccc0ae76eab392fa99f7625adf09ca7f76306a8375d3a707ff2799db54cd6"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-720ccc0ae76eab392fa99f7625adf09ca7f76306a8375d3a707ff2799db54cd6)
- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-58b178655a495d4d8282921323d01ffa982e49c5e6c4bdb2b61c411d0dc4d95f)
- [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-6327285e6f8dcdb7c517d68ed9b1a3da046a1198c4685ef6992770af892c030d)

Select alternatives according to the provider validators above.

<a id="canonical-d609874d5ee710d46fd3e5e472af188babc3226a81f977d4e3a3eb339034e98c"></a>

## Direct properties — ike_keylifetime_hours / cc1c9b157391 / 3

<a id="canonical-c126e07fc3dfdf0444df8dab4efacf6a15a9b8deb664e9f9149859f65fde528e"></a>

<a id="canonical-7d6c5a96d087bc7e4dcef0fc5126a87f0426aba4a93346123d33feb932584875"></a>

## duration property — ike_keylifetime_hours / cc1c9b157391 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-e73d4de59188e3666b24243dfeb73e8372b76949ab0fa63c9488203c28d594b1"></a>

## Next pages — ike_keylifetime_hours / cc1c9b157391 / 5

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-f86d6add4eec8c679321cd96a5ba2693a6588034563641d3d751260483a03f73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c656def51d66afdd8c165c1fba0a6aee69e96a2c59e91b6d070a10bc3a2c807a"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 099317603608 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- ike_keylifetime_minutes

<a id="canonical-58b178655a495d4d8282921323d01ffa982e49c5e6c4bdb2b61c411d0dc4d95f"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-f9074e01abb6d23a19ff62c3ef5d72250bcfd3ebb9deede741028e01061f4952"></a>

## Direct properties — ike_keylifetime_minutes / 099317603608 / 3

<a id="canonical-240c824a31f0cd21088acb4f1c6985355cecab2f101fae60c1e7fa9edc6f53de"></a>

<a id="canonical-fba94a6b88d6393d6eef6aeed83d06d07c16529556b2ade886cb93fa0589a080"></a>

## duration property — ike_keylifetime_minutes / 099317603608 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

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
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-f0ebf70ddadf3ea1ba882db468d981249c04315830005ed9c120e4eb0b316416"></a>

## Next pages — ike_keylifetime_minutes / 099317603608 / 5

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-89ceae40686c00b1e6d586841e1058cd1b0940b8e7cac4233e283853125268b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d50bd9277a429b4ed6efcc78d0daacaa03e241f637831ce435e67efef9734339"></a>

## reauth_disabled — reauth_disabled / 1ff5270b06e3 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- reauth_disabled

<a id="canonical-3d9a9d804764976d971fc4e61652d75e6e854046be745fd74406f4998972205a"></a>

Type: `["object", {}]`. Computed.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

OneOf alternatives in this subsection:

- [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3d9a9d804764976d971fc4e61652d75e6e854046be745fd74406f4998972205a)
- [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-e6aaf7f795da55f7e05dbdb9a879b055476847afab7137a26cc3300e6b3df829)
- [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-a70c23d5e0b0829958ccae393b18e7f933fcc7699e343a2b2c3f937da8a20783)

Select alternatives according to the provider validators above.

<a id="canonical-60512df34e3a260eac1e3dc0679ec53bf770c74fbf33761163705f07fecb4bf8"></a>

## Direct properties — reauth_disabled / 1ff5270b06e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20632cfcf174225dd0cc96c3da5e07e1570bda5e080f55cb78ef2b1f1ffa2903"></a>

## Next pages — reauth_disabled / 1ff5270b06e3 / 4

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-c86d108dd9d41d82e437353de084c84587ccba24264612d5c5dc8004a053a6c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91cb049b9144df82fa4958852111667443087cc33a2dfb33e906728a26c4c89e"></a>

## reauth_timeout_days — reauth_timeout_days / 19bc8411bbac / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- reauth_timeout_days

<a id="canonical-e6aaf7f795da55f7e05dbdb9a879b055476847afab7137a26cc3300e6b3df829"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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

<a id="canonical-1e5ceb69edfce6571dfe6e11a8f4c3e821e9f20eaf96d8b02d45913fb27cbf92"></a>

## Direct properties — reauth_timeout_days / 19bc8411bbac / 3

<a id="canonical-cdab320fba8c7c3adb19296d0861404755e25485245d5105c14ecf889c9b004d"></a>

<a id="canonical-b615bf2820577057cdef0e96c0e28dc4e73ce3d51814ce9d4609bd40f036ff43"></a>

## duration property — reauth_timeout_days / 19bc8411bbac / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-0ebcc1157c1fb6452847e2b16add8654120c8a62a4c54bb8112e7eebc5d546fe"></a>

## Next pages — reauth_timeout_days / 19bc8411bbac / 5

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-6ca185d07dc5cb21643520089952ad9bc4f288fb1cb7d4aa59cccd1a662d4a25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d016cc3241ab1dfbd683e4b3865f97539c93a8387b30b86252eedcb59c7ee33"></a>

## reauth_timeout_hours — reauth_timeout_hours / 0c8502539e6b / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- reauth_timeout_hours

<a id="canonical-a70c23d5e0b0829958ccae393b18e7f933fcc7699e343a2b2c3f937da8a20783"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

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

<a id="canonical-57abe95c04f8a6c759f366eff39eb73a539d6df1f3c59b684199cbba31898dbd"></a>

## Direct properties — reauth_timeout_hours / 0c8502539e6b / 3

<a id="canonical-e86fe5eda6af196e9030b6ad57b42b94744afdbb12fb4fd48650688311bf6c60"></a>

<a id="canonical-618fb0c4131727fe96f9827d709ebed7fd55e160ff355adc368beef217f99e12"></a>

## duration property — reauth_timeout_hours / 0c8502539e6b / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-788d5fc0684baa1c40922941d69ba539a95716823bf3d380ed8a632818ad15ea"></a>

## Next pages — reauth_timeout_hours / 0c8502539e6b / 5

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)

<a id="canonical-e47b7caf259b20fa0f5838daf0c8610633a7c9b16e9a070cdf17804e27072a86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c518390d9b818daf941c2026a195605f6e250b4680eb9dd2b4fdd189f2c9f123"></a>

## use_default_keylifetime — use_default_keylifetime / bc3faa41fbb8 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- use_default_keylifetime

<a id="canonical-6327285e6f8dcdb7c517d68ed9b1a3da046a1198c4685ef6992770af892c030d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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

<a id="canonical-2f73e95d1bcc2eeb54ba3b7d2b17b5f81f03ed9925a3114f5b389c4d353c4e8b"></a>

## Direct properties — use_default_keylifetime / bc3faa41fbb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-668f5a5e5034f278a551b522cfbff7db7b3b5af770cd851333c9448e3d32c316"></a>

## Next pages — use_default_keylifetime / bc3faa41fbb8 / 4

- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b)
