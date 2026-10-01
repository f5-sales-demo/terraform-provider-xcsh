---
page_title: "xcsh_data_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type reference."
---

# xcsh_data_type reference

<a id="canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af14e530967015bcac7f1c85302c51b9835972887f0c29f279650a725c9d570e"></a>

## Property reference — Property reference / 88f4c39ef970 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- Property reference

<a id="canonical-8bc7d1fa27fa191fe9c04c7fa500557162a9edfac24ebfc0b92f50172d2eeeb2"></a>

## Direct properties — Property reference / 88f4c39ef970 / 3

<a id="canonical-1ce62dd9263fac4208e9cc45263e6b27355d64b805bc203fa73ad1f8279f6f26"></a>

<a id="canonical-505dec8048dcf3c3e3664b1e0b311ce39db9101d1bd7bcfdfe88986a51036605"></a>

## annotations property — Property reference / 88f4c39ef970 / 4

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

<a id="canonical-fcdc3981dd030105d3c50c60f7e2ce7d0d80e3763cbc1c31e022cee47dccfdf5"></a>

<a id="canonical-d05a1efc403d5c615ee23694b23f81d2e36e5e789e951fc7d1eefcdb4068e20e"></a>

## compliances property — Property reference / 88f4c39ef970 / 5

Type: `["list", "string"]`. Optional.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach. Possible values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`,
\`APPI\`, \`HIPAA\`, \`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`,
\`ISO\_IEC\_27701\`, \`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Upstream description:

Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-aa2dbfdfd06210d46d6a1166dc000a816f161466a6f18a569f56b39abaa1a331"></a>

<a id="canonical-b3a40e97c653f6fe529fb15bed3ce2e426ee9f895e49b515761f0af3d5433d87"></a>

## description property — Property reference / 88f4c39ef970 / 6

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

<a id="canonical-7c38100754a050d703f42b272ff0f384e67b7e173969e7cb65b8249ece525c87"></a>

<a id="canonical-967a5dde3df5c0d9ae0f7b473eca16d500dc824749141f32c81cc2da36d282fa"></a>

## disable property — Property reference / 88f4c39ef970 / 7

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

<a id="canonical-a38a32393b84ef53e195d990e23d636ba1da52e94a609b7ec76fe36114feac61"></a>

<a id="canonical-511994aa82aaed9be56b06edb3ead23775066e69d76554a3f54488489a9287a6"></a>

## id property — Property reference / 88f4c39ef970 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-217da0fa75b1d9528d89be1d73ec555a35304dc7a1f09cc2994257745684216d"></a>

<a id="canonical-32fa2eea6b2f65fccf334e78bb3fa58bc9d6d6047d6d2ce1157f0389de1ec371"></a>

## is_pii property — Property reference / 88f4c39ef970 / 9

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as personally identifiable information (PII).

Upstream description:

Select this option to classify the custom data type as personally identifiable information (PII)

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

<a id="canonical-732a04d01c63dc02538c72f55d78f85d6929b60bab5d1d83c6bad1f880253d45"></a>

<a id="canonical-5e89f90ed9afe002e60fed1436b6de69a4a56c768adaa7078f5ae53244fcb341"></a>

## is_sensitive_data property — Property reference / 88f4c39ef970 / 10

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as sensitive, enabling detection of API
vulnerabilities related to this data type.

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

<a id="canonical-6a9e3dd92419d199d04ff8d40ce570fac7e7f3b2a5a04da1bfc6911a687b7342"></a>

<a id="canonical-6fa8b30183168b717dd0053682a81be1a8c13e5334c52291ff6bb881cb100cb1"></a>

## labels property — Property reference / 88f4c39ef970 / 11

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

<a id="canonical-6e04c03dd2734d7f8e0db2d2fdf2d64e70aec75c1de6406dbb65891fae909c0d"></a>

<a id="canonical-044e68e284b232a67b33e6e0449c08d2f5431cf0978d8fdcd5eda5f26c5e7622"></a>

## name property — Property reference / 88f4c39ef970 / 12

Type: `"string"`. Required.

Name of the Data Type. Must be unique within the namespace.

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

<a id="canonical-234aa9ef7d78afcccd74a7cd062c9636d39f0a7dc5c563cf54131462074ed0e7"></a>

<a id="canonical-b00168abb5e0a15f262f6b1d5888dbfc02d757dfb4d86c49440c9231545d5b99"></a>

## namespace property — Property reference / 88f4c39ef970 / 13

Type: `"string"`. Required.

Namespace where the Data Type is created.

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

- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f): complete subsection reference.

- [timeouts](resources--data_type--reference--group-001.md#canonical-d6857eddf697e8ad5f64e5efd56156f523af26e99dd74a845587b1a363cb8292): complete subsection reference.

<a id="canonical-1401e91488edf022ca5a53dd245ba8997e1846366cfd73a6ed862d9673a4dc61"></a>

## All schema paths — Property reference / 88f4c39ef970 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--data_type--reference--group-001.md#canonical-1ce62dd9263fac4208e9cc45263e6b27355d64b805bc203fa73ad1f8279f6f26) |
| `compliances` | [compliances](resources--data_type--reference--group-001.md#canonical-fcdc3981dd030105d3c50c60f7e2ce7d0d80e3763cbc1c31e022cee47dccfdf5) |
| `description` | [description](resources--data_type--reference--group-001.md#canonical-aa2dbfdfd06210d46d6a1166dc000a816f161466a6f18a569f56b39abaa1a331) |
| `disable` | [disable](resources--data_type--reference--group-001.md#canonical-7c38100754a050d703f42b272ff0f384e67b7e173969e7cb65b8249ece525c87) |
| `id` | [id](resources--data_type--reference--group-001.md#canonical-a38a32393b84ef53e195d990e23d636ba1da52e94a609b7ec76fe36114feac61) |
| `is_pii` | [is_pii](resources--data_type--reference--group-001.md#canonical-217da0fa75b1d9528d89be1d73ec555a35304dc7a1f09cc2994257745684216d) |
| `is_sensitive_data` | [is_sensitive_data](resources--data_type--reference--group-001.md#canonical-732a04d01c63dc02538c72f55d78f85d6929b60bab5d1d83c6bad1f880253d45) |
| `labels` | [labels](resources--data_type--reference--group-001.md#canonical-6a9e3dd92419d199d04ff8d40ce570fac7e7f3b2a5a04da1bfc6911a687b7342) |
| `name` | [name](resources--data_type--reference--group-001.md#canonical-6e04c03dd2734d7f8e0db2d2fdf2d64e70aec75c1de6406dbb65891fae909c0d) |
| `namespace` | [namespace](resources--data_type--reference--group-001.md#canonical-234aa9ef7d78afcccd74a7cd062c9636d39f0a7dc5c563cf54131462074ed0e7) |
| `rules` | [rules](resources--data_type--reference--group-001.md#canonical-1541c2d365a427cd53ebf7b2985df75ab8c8664d4b5997520502c2cd97fb3220) |
| `rules.key_pattern` | [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-d2ab56d6586b32915fb5bf6ea9a21c7b9cc487b020f92889a44f32b80048660a) |
| `rules.key_pattern.exact_values` | [rules.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-63f35e7cc060e8ebee84cb603bf390f5ca6e32328578397bbe272e3dfac09f08) |
| `rules.key_pattern.exact_values.exact_values` | [rules.key_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-2788ba9334176caa4aa9de5243ad43f96dbf4eae3df19a4f30c6e9112fb94d68) |
| `rules.key_pattern.regex_value` | [rules.key_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-af7cffc53a17492d4e6e01eb135233ea96422c6b583ba6fdb1a1ba8b1cb5f610) |
| `rules.key_pattern.substring_value` | [rules.key_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-2a2033ec61377234ed32edbb142e545cbf407b14aeefc5f787974771cf50a8ac) |
| `rules.key_value_pattern` | [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a7bd1a5867de7de77bf28f11c1ef3efd79035580ee2103e1e50e1c3c6d4bf91b) |
| `rules.key_value_pattern.key_pattern` | [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-06c67125772877c33c332759446e2afd9bc2aa1e7a2e0f0d8e21e82fd290435f) |
| `rules.key_value_pattern.key_pattern.exact_values` | [rules.key_value_pattern.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-c7935cf5e22eecccd7372ced807f628b9594c27881a00b55b70f8feb1fe872a6) |
| `rules.key_value_pattern.key_pattern.exact_values.exact_values` | [rules.key_value_pattern.key_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-19e3d03ed2eaba1c4637b559fda1fd79b2cfb96df7d0baf8dd900e3223750d8d) |
| `rules.key_value_pattern.key_pattern.regex_value` | [rules.key_value_pattern.key_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-77afda751e5807ace53eeb14efb8a84b4c4ad2a96ed824d66724c3a894dc729e) |
| `rules.key_value_pattern.key_pattern.substring_value` | [rules.key_value_pattern.key_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-3cb4e1150d9c0c997e29d4717aba55aa63f174a0c3951e6c759edf0f01197ff2) |
| `rules.key_value_pattern.value_pattern` | [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-72d7cca122df66cebb9c777c0dac7d6e1777134a278fd64c39c364cee5bb191d) |
| `rules.key_value_pattern.value_pattern.exact_values` | [rules.key_value_pattern.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-d2c42aaaa5763dc651cd096e468c93bc0a30749d4889f9fccaa983d527d00d7c) |
| `rules.key_value_pattern.value_pattern.exact_values.exact_values` | [rules.key_value_pattern.value_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-6f450765571d7e2885c045085d892efa1f4fbbee6b46fefe17f2eda7335247a7) |
| `rules.key_value_pattern.value_pattern.regex_value` | [rules.key_value_pattern.value_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-ead8f55bee4e15c854bb0c311d94979727ce0af2d5d8b337d76c8dc244e662b6) |
| `rules.key_value_pattern.value_pattern.substring_value` | [rules.key_value_pattern.value_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-c9256ee21aaa86c677eec3cde1e6bf8b701b69e47a6324de5df056cde3161f87) |
| `rules.value_pattern` | [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-2395fd4a58d5110a6561eadd89fb8193e2b25ef9f129f3217ce95b707b96c37f) |
| `rules.value_pattern.exact_values` | [rules.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-721dcf583dcc39c11078ccd9b4f55498ae9921f668e87cf53b6cdb7bc14cdb84) |
| `rules.value_pattern.exact_values.exact_values` | [rules.value_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-e9494cb2704bf820377b13678b5e924a95697c3dd3ed9f593971f77d6057ff28) |
| `rules.value_pattern.regex_value` | [rules.value_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-6e550d172a9259f9c8c7621a7bf06331a0fa58bca5338504664e7fbaadb4e013) |
| `rules.value_pattern.substring_value` | [rules.value_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-db60b42db49e739aea274d8e45c69fab35b69185e17f0c2d6de67833793230fb) |
| `timeouts` | [timeouts](resources--data_type--reference--group-001.md#canonical-1024f000576cb37d6c6960f738fa4f70ecc53b390baa39683adb5d93182d8b6a) |
| `timeouts.create` | [timeouts.create](resources--data_type--reference--group-001.md#canonical-9c92d2a972dded9510d43756bde53af29c0d1d0fa69a86f8774861c9608ae738) |
| `timeouts.delete` | [timeouts.delete](resources--data_type--reference--group-001.md#canonical-10ce832519c442254f2d207d8ffc9b80d9a4bf8ffa1eacd8aa019049ee324dae) |
| `timeouts.read` | [timeouts.read](resources--data_type--reference--group-001.md#canonical-4e4d15d0b910c0a72fe5338c49eced81a71647d8d09334f92fa7f243fcc49435) |
| `timeouts.update` | [timeouts.update](resources--data_type--reference--group-001.md#canonical-150bf9a4a6a491ef2e84072db520dc49c8e5cad39c4362472069c76c9121e126) |

<a id="canonical-1cb5bcf519aa20a594f3488203839d3843bee9a87c1420e3b17948e0385f801e"></a>

## Next pages — Property reference / 88f4c39ef970 / 15

- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [timeouts](resources--data_type--reference--group-001.md#canonical-d6857eddf697e8ad5f64e5efd56156f523af26e99dd74a845587b1a363cb8292)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08056126ae4d923b987ef78926af318765fab027a51de0630883f0293d49a3e9"></a>

## rules — rules / bf4c0de68c48 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- rules

<a id="canonical-1541c2d365a427cd53ebf7b2985df75ab8c8664d4b5997520502c2cd97fb3220"></a>

Type: `"object"`. list nested block, Optional.

Configure key/value or regex match rules to enable the platform to detect this custom data type in
the API request or response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("key_pattern",
    "key_value_pattern"),
  validators.ConflictingListObjectAttributes("key_pattern",
    "value_pattern"),
  validators.ConflictingListObjectAttributes("key_value_pattern",
    "value_pattern")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-37c83abf66d10dac0251fd92fce468e94c08404dbb3cbb5034fe7b8854491a85"></a>

## Direct properties — rules / bf4c0de68c48 / 3

- [key_pattern](resources--data_type--reference--group-001.md#canonical-7b453bdc4fae6ae63e906f99f39415c9d4a8eb7215e62bf724af8e6a67a9cdbe): complete subsection reference.

- [key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075): complete subsection reference.

- [value_pattern](resources--data_type--reference--group-001.md#canonical-247b23e08738a5f0c30111f98604a9ab4e371303bf9e7480c75bc6830f1a46a6): complete subsection reference.

<a id="canonical-fe9754a04e6096556a602cbe53e4337d7d599edefadb88ae4d79064ca59bc00c"></a>

## Next pages — rules / bf4c0de68c48 / 4

- [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-7b453bdc4fae6ae63e906f99f39415c9d4a8eb7215e62bf724af8e6a67a9cdbe)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-247b23e08738a5f0c30111f98604a9ab4e371303bf9e7480c75bc6830f1a46a6)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-7b453bdc4fae6ae63e906f99f39415c9d4a8eb7215e62bf724af8e6a67a9cdbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b2310cd9411f45e3eec9d782e6e6a69ffdb701d6e3e22d029f8c06a130e8d77"></a>

## rules.key_pattern — rules.key_pattern / ec70af0d1bb2 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- rules.key_pattern

<a id="canonical-d2ab56d6586b32915fb5bf6ea9a21c7b9cc487b020f92889a44f32b80048660a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
key_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-61017ce480a08c1dc5bdb9838244904a573257248663dc4df6a4c32649097d96"></a>

## Direct properties — rules.key_pattern / ec70af0d1bb2 / 3

- [exact_values](resources--data_type--reference--group-001.md#canonical-ea23e817dca5dda3940f4519dd1ec5b84bb62606c8c516ed0e73faf6bb3b064b): complete subsection reference.

<a id="canonical-af7cffc53a17492d4e6e01eb135233ea96422c6b583ba6fdb1a1ba8b1cb5f610"></a>

<a id="canonical-ac9b2f13def7099cff1b99cfb44e79295e81c3e7557ec06d0e1e1c6f3b635dbc"></a>

## regex_value property — rules.key_pattern / ec70af0d1bb2 / 4

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2a2033ec61377234ed32edbb142e545cbf407b14aeefc5f787974771cf50a8ac"></a>

<a id="canonical-1b85db798ed9346451f7d8cde5b82ec491965ffd7792ab1bb916886c32c1eee0"></a>

## substring_value property — rules.key_pattern / ec70af0d1bb2 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-72deadc153f3b7ad699e0ad695e22cc38fedc2a1d9af7afa7022781320ae8793"></a>

## Next pages — rules.key_pattern / ec70af0d1bb2 / 6

- [rules.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-ea23e817dca5dda3940f4519dd1ec5b84bb62606c8c516ed0e73faf6bb3b064b)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-ea23e817dca5dda3940f4519dd1ec5b84bb62606c8c516ed0e73faf6bb3b064b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa062b9d7a38e818d3ce16ef06941807ccebdc0afa6fda8b08bfdc2ecc1a584a"></a>

## rules.key_pattern.exact_values — rules.key_pattern.exact_values / 3a08cbbe27ec / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-7b453bdc4fae6ae63e906f99f39415c9d4a8eb7215e62bf724af8e6a67a9cdbe)
- rules.key_pattern.exact_values

<a id="canonical-63f35e7cc060e8ebee84cb603bf390f5ca6e32328578397bbe272e3dfac09f08"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-d51e07e5cafe015924159ff98ca82cf6953e78c8ec6b3552aae45dddbbd1a44e"></a>

## Direct properties — rules.key_pattern.exact_values / 3a08cbbe27ec / 3

<a id="canonical-2788ba9334176caa4aa9de5243ad43f96dbf4eae3df19a4f30c6e9112fb94d68"></a>

<a id="canonical-4ee6a8a930bc6d7847de657fd2c2d852218a1ddcd73794f4c35ee11e46601515"></a>

## exact_values property — rules.key_pattern.exact_values / 3a08cbbe27ec / 4

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f4e181ab4094df92d5a6041983c25f1c802bead87119f77d655648cbeb9caeb1"></a>

## Next pages — rules.key_pattern.exact_values / 3a08cbbe27ec / 5

- [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-7b453bdc4fae6ae63e906f99f39415c9d4a8eb7215e62bf724af8e6a67a9cdbe)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abbcc260f7e2925731f143d5644955ba1207bfee9b8a9ac934d49fad87e24886"></a>

## rules.key_value_pattern — rules.key_value_pattern / b48fa2df9abb / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- rules.key_value_pattern

<a id="canonical-a7bd1a5867de7de77bf28f11c1ef3efd79035580ee2103e1e50e1c3c6d4bf91b"></a>

Type: `"object"`. single nested block, Optional.

Search for specific key &amp; value patterns in the specified sections.

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
key_value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-461415dd9ca8f78b337a61e364f96a558fa7ea0952c07d58e91bae3dc89dfc25"></a>

## Direct properties — rules.key_value_pattern / b48fa2df9abb / 3

- [key_pattern](resources--data_type--reference--group-001.md#canonical-a3b62e26ddd1cefe2500959d240b92425f80b4795b2d20fbc818ac5a30105267): complete subsection reference.

- [value_pattern](resources--data_type--reference--group-001.md#canonical-9d041350217effbe2fcd8bde6f1709b135a0dc57146b651fccb7d6332e955b12): complete subsection reference.

<a id="canonical-0e88dacdde352d293501d951f645c8d47d9d138225a668219e6984cfd8889e37"></a>

## Next pages — rules.key_value_pattern / b48fa2df9abb / 4

- [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-a3b62e26ddd1cefe2500959d240b92425f80b4795b2d20fbc818ac5a30105267)
- [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-9d041350217effbe2fcd8bde6f1709b135a0dc57146b651fccb7d6332e955b12)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-a3b62e26ddd1cefe2500959d240b92425f80b4795b2d20fbc818ac5a30105267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ef3a2bc032c800974f80abac2b0d808fd2ce462e7817d455077a05a22a27476"></a>

## rules.key_value_pattern.key_pattern — rules.key_value_pattern.key_pattern / 95dc78bc1ce3 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- rules.key_value_pattern.key_pattern

<a id="canonical-06c67125772877c33c332759446e2afd9bc2aa1e7a2e0f0d8e21e82fd290435f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
key_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-48ae513bb263d24af5ea1374161eefb26b2fa00b2dc4c25d5f13cdc2ed05cf74"></a>

## Direct properties — rules.key_value_pattern.key_pattern / 95dc78bc1ce3 / 3

- [exact_values](resources--data_type--reference--group-001.md#canonical-32feea47059900693d68e8afd3bab4220f9caa54440bf17894f223f94ae2842f): complete subsection reference.

<a id="canonical-77afda751e5807ace53eeb14efb8a84b4c4ad2a96ed824d66724c3a894dc729e"></a>

<a id="canonical-da0e4909527e963ab1c45f2106aa3a855875b66de4dd0de5eea7a183dcf0a845"></a>

## regex_value property — rules.key_value_pattern.key_pattern / 95dc78bc1ce3 / 4

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3cb4e1150d9c0c997e29d4717aba55aa63f174a0c3951e6c759edf0f01197ff2"></a>

<a id="canonical-71b90ef5cc17f8a00632c3788ebc519d8ef94b8c6d7c5476f005deb06e8e2307"></a>

## substring_value property — rules.key_value_pattern.key_pattern / 95dc78bc1ce3 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-9d6304f479d7b80fae46ebb3c466f65f1aa5aba044d63c4001fc99c52ca05062"></a>

## Next pages — rules.key_value_pattern.key_pattern / 95dc78bc1ce3 / 6

- [rules.key_value_pattern.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-32feea47059900693d68e8afd3bab4220f9caa54440bf17894f223f94ae2842f)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-32feea47059900693d68e8afd3bab4220f9caa54440bf17894f223f94ae2842f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97a58dbaaa2abf3bbc618ddafe6d172bccc9d8036a5876b87d08cf551e527547"></a>

## rules.key_value_pattern.key_pattern.exact_values — rules.key_value_pattern.key_pattern.exact_values / d688a6ba71ee / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-a3b62e26ddd1cefe2500959d240b92425f80b4795b2d20fbc818ac5a30105267)
- rules.key_value_pattern.key_pattern.exact_values

<a id="canonical-c7935cf5e22eecccd7372ced807f628b9594c27881a00b55b70f8feb1fe872a6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-e66a18bb23a355cbb340a028c12fa59f9d96ff17e7625c03b8c967a3de58442e"></a>

## Direct properties — rules.key_value_pattern.key_pattern.exact_values / d688a6ba71ee / 3

<a id="canonical-19e3d03ed2eaba1c4637b559fda1fd79b2cfb96df7d0baf8dd900e3223750d8d"></a>

<a id="canonical-c0f43a541707dc7e2c3e13ffeb53f7656988185736348972ae754d95aab1d544"></a>

## exact_values property — rules.key_value_pattern.key_pattern.exact_values / d688a6ba71ee / 4

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ffc6aea9b088d2a5d568fff1debb61caacfef70ea977f0a1b2c0647751d8d293"></a>

## Next pages — rules.key_value_pattern.key_pattern.exact_values / d688a6ba71ee / 5

- [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-a3b62e26ddd1cefe2500959d240b92425f80b4795b2d20fbc818ac5a30105267)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-9d041350217effbe2fcd8bde6f1709b135a0dc57146b651fccb7d6332e955b12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07598cfb303f4e9847542e448bc67b49ba0fc4eb3ec334d799dd321661c01bce"></a>

## rules.key_value_pattern.value_pattern — rules.key_value_pattern.value_pattern / 78f71c13df5a / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- rules.key_value_pattern.value_pattern

<a id="canonical-72d7cca122df66cebb9c777c0dac7d6e1777134a278fd64c39c364cee5bb191d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7f5fa4c25a13470bcff102455886467783025d275f7404853a281f5c66914d5"></a>

## Direct properties — rules.key_value_pattern.value_pattern / 78f71c13df5a / 3

- [exact_values](resources--data_type--reference--group-001.md#canonical-92c12d7e05005ab7b7aca022c39050ca48efbcbd91b1c9a85fb6d749710b8d20): complete subsection reference.

<a id="canonical-ead8f55bee4e15c854bb0c311d94979727ce0af2d5d8b337d76c8dc244e662b6"></a>

<a id="canonical-1fbbf41bb4fa0541462659871c94681dac504d29e625bf2163759024fccf45cb"></a>

## regex_value property — rules.key_value_pattern.value_pattern / 78f71c13df5a / 4

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-c9256ee21aaa86c677eec3cde1e6bf8b701b69e47a6324de5df056cde3161f87"></a>

<a id="canonical-00fd746e14391e6161ceef8cf1fd3f0ce3be280da790c61ec5f516b03c0fabba"></a>

## substring_value property — rules.key_value_pattern.value_pattern / 78f71c13df5a / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-05e6a45ab7c7904162f90a7908963ea050300301da9a54f3a2ab3ae548c5eab1"></a>

## Next pages — rules.key_value_pattern.value_pattern / 78f71c13df5a / 6

- [rules.key_value_pattern.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-92c12d7e05005ab7b7aca022c39050ca48efbcbd91b1c9a85fb6d749710b8d20)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-92c12d7e05005ab7b7aca022c39050ca48efbcbd91b1c9a85fb6d749710b8d20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f14328948a8403829dcc295937c6b6d68aaaf5df24f8b9cbdc79022f22f17e3"></a>

## rules.key_value_pattern.value_pattern.exact_values — rules.key_value_pattern.value_pattern.exact_values / 970d3b7299fc / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-a3e81a9e98d2d46c86313d37bfb332775276cd48ab9b4b23f3eadc431f38f075)
- [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-9d041350217effbe2fcd8bde6f1709b135a0dc57146b651fccb7d6332e955b12)
- rules.key_value_pattern.value_pattern.exact_values

<a id="canonical-d2c42aaaa5763dc651cd096e468c93bc0a30749d4889f9fccaa983d527d00d7c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b53660e14ed53e108b19b62a332b29c6df778f8484a92b5102375e9dc4a19ab"></a>

## Direct properties — rules.key_value_pattern.value_pattern.exact_values / 970d3b7299fc / 3

<a id="canonical-6f450765571d7e2885c045085d892efa1f4fbbee6b46fefe17f2eda7335247a7"></a>

<a id="canonical-25d628efd3e5dff972131407fd83cd65bdc354b2c5569cb4174f7350f3b6e235"></a>

## exact_values property — rules.key_value_pattern.value_pattern.exact_values / 970d3b7299fc / 4

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e1c9b6748e84b16165a24b93185cf5a3aa9a909abc5c4846c0e5e4b10dedd974"></a>

## Next pages — rules.key_value_pattern.value_pattern.exact_values / 970d3b7299fc / 5

- [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-9d041350217effbe2fcd8bde6f1709b135a0dc57146b651fccb7d6332e955b12)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-247b23e08738a5f0c30111f98604a9ab4e371303bf9e7480c75bc6830f1a46a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-291783e97a678adf64b2896d9170ba92936bd5df05b9c99152fc11d952d1d60e"></a>

## rules.value_pattern — rules.value_pattern / 6c1cf658ada1 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- rules.value_pattern

<a id="canonical-2395fd4a58d5110a6561eadd89fb8193e2b25ef9f129f3217ce95b707b96c37f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-205f146c1a8e32c463b49aaefc3693cce5998d9f4f45c22c563719f5e7e8820e"></a>

## Direct properties — rules.value_pattern / 6c1cf658ada1 / 3

- [exact_values](resources--data_type--reference--group-001.md#canonical-355e17b43ce734d817407ed6c535c3f9105d3f32d16365667a160830ab26f3d9): complete subsection reference.

<a id="canonical-6e550d172a9259f9c8c7621a7bf06331a0fa58bca5338504664e7fbaadb4e013"></a>

<a id="canonical-ac800e3d8af6fc2923760ff191a1e2a54b0601aee3ea7a8ae8a940ceff0669bb"></a>

## regex_value property — rules.value_pattern / 6c1cf658ada1 / 4

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-db60b42db49e739aea274d8e45c69fab35b69185e17f0c2d6de67833793230fb"></a>

<a id="canonical-01f11c0d3197288078c6c2520cfa4254b1f3358f6acb7720612629a0e305670f"></a>

## substring_value property — rules.value_pattern / 6c1cf658ada1 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-15107b6cb7b1f26c375ceeefa4f3dddeb0fde8240f5b916e05e67da5bef54fa7"></a>

## Next pages — rules.value_pattern / 6c1cf658ada1 / 6

- [rules.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-355e17b43ce734d817407ed6c535c3f9105d3f32d16365667a160830ab26f3d9)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-355e17b43ce734d817407ed6c535c3f9105d3f32d16365667a160830ab26f3d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e98a1c38f9955bff8f5c57bf5957d4fc9bf65df7c75612b0f0bad9a6a18e388"></a>

## rules.value_pattern.exact_values — rules.value_pattern.exact_values / 320292108d54 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [rules](resources--data_type--reference--group-001.md#canonical-074aeb3e6c5915de0271d727586c4f776b5aaa774a41e11cc3460b5b04b0ab7f)
- [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-247b23e08738a5f0c30111f98604a9ab4e371303bf9e7480c75bc6830f1a46a6)
- rules.value_pattern.exact_values

<a id="canonical-721dcf583dcc39c11078ccd9b4f55498ae9921f668e87cf53b6cdb7bc14cdb84"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e52a9fb5701e6f90aaea2577aa22c1991b6285ae7847ae1c579e99fc4bf56a7"></a>

## Direct properties — rules.value_pattern.exact_values / 320292108d54 / 3

<a id="canonical-e9494cb2704bf820377b13678b5e924a95697c3dd3ed9f593971f77d6057ff28"></a>

<a id="canonical-b60eefc8e8380a413a8ff547f1b9794f5a9f5a2c5865079c8d30b7d4dc93a3e5"></a>

## exact_values property — rules.value_pattern.exact_values / 320292108d54 / 4

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f997f3380872a353c3e1d1b5b29ac1eebdf958a033a4cf181a036dccdf60a290"></a>

## Next pages — rules.value_pattern.exact_values / 320292108d54 / 5

- [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-247b23e08738a5f0c30111f98604a9ab4e371303bf9e7480c75bc6830f1a46a6)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)

<a id="canonical-d6857eddf697e8ad5f64e5efd56156f523af26e99dd74a845587b1a363cb8292"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-918cbd934a8a4f3b72fd5ecaca56bd8160aefa0c44a1b997c2fef198490c61cd"></a>

## timeouts — timeouts / 7f694c1dd699 / 2

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- timeouts

<a id="canonical-1024f000576cb37d6c6960f738fa4f70ecc53b390baa39683adb5d93182d8b6a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3df46f44c74d307afe3c5b2f1261625952597e56c37f16b932e2d6bbc276a8e"></a>

## Direct properties — timeouts / 7f694c1dd699 / 3

<a id="canonical-9c92d2a972dded9510d43756bde53af29c0d1d0fa69a86f8774861c9608ae738"></a>

<a id="canonical-8c1f4417ed61eed56a7b85f0b0be30443ac9400a41c505822c3655377f81489c"></a>

## create property — timeouts / 7f694c1dd699 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-10ce832519c442254f2d207d8ffc9b80d9a4bf8ffa1eacd8aa019049ee324dae"></a>

<a id="canonical-17b386595d192ffad16c69bf0bc242a63bc995f27c3d82e0ade0668b67704645"></a>

## delete property — timeouts / 7f694c1dd699 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-4e4d15d0b910c0a72fe5338c49eced81a71647d8d09334f92fa7f243fcc49435"></a>

<a id="canonical-7dd0db2e0b317784f3637b2f4f5165e12b38bf5bb6c4722bd22e26d49e519e49"></a>

## read property — timeouts / 7f694c1dd699 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-150bf9a4a6a491ef2e84072db520dc49c8e5cad39c4362472069c76c9121e126"></a>

<a id="canonical-e2bfbddebf39d1dec882d8487db2a316fb213444c36205fe9c0be717c48454ca"></a>

## update property — timeouts / 7f694c1dd699 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2c298cef9260b464b446e561114e8f1f64174d78fc0ec55c5be2f5c7481e6299"></a>

## Next pages — timeouts / 7f694c1dd699 / 8

- [Property reference](resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [xcsh_data_type](../resources/data_type.md#canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc)
