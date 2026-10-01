---
page_title: "xcsh_data_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type reference."
---

# xcsh_data_type reference

<a id="canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-838da1dfa635fd45f76519a1b76bc45bad552eaa64990cfa2b4096218348180f"></a>

## Property reference — Property reference / e380725c82d9 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- Property reference

<a id="canonical-2a9f56ccb0e21883ed09acdf271a31a50f43da65f160e6369bdc54c600feb151"></a>

## Direct properties — Property reference / e380725c82d9 / 3

<a id="canonical-901a440d452a6f419ecc4290cfe9a450646e6d3119c1cfeaf5f77ebf82d3a6ff"></a>

<a id="canonical-7c6cbe60c8d9e55645e67250d48484433ff67b6e127b9e17ba1aa90396a7a446"></a>

## annotations property — Property reference / e380725c82d9 / 4

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

<a id="canonical-ad56eec22ff5822fe7116ff893f20abb4dc8c40b069a2b69a70aec4c2b499dcb"></a>

<a id="canonical-dc6d6d0ff2b897c6421f4ab94a61b59582ca62912fff7c6f33ca777f6a8975ae"></a>

## compliances property — Property reference / e380725c82d9 / 5

Type: `["list", "string"]`. Computed.

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

<a id="canonical-6cd2217fc3f92f432eb7c49c86a3a428fd130daaf13a9a3a7f8acb75850ea6fb"></a>

<a id="canonical-9e90924c079feb9211cba2259edf0507f75862b607978532c34c5f6fc76e137c"></a>

## description property — Property reference / e380725c82d9 / 6

Type: `"string"`. Computed.

Description of the DataType.

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

<a id="canonical-095c646d3966bd6ee0ca2b9407e0e4605649bb5bb91e49c86506298105f04f7d"></a>

<a id="canonical-350df725c626d116d9686df74dc31e1f2c514f7f4a94a6f27f23bd3143dc3989"></a>

## id property — Property reference / e380725c82d9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4273659dbab11f6e27d2f7250bba976bad7ea2b6916f4288732aa6e7859db45b"></a>

<a id="canonical-dfab626e79da5372dc5c23f82d337846fc2ddb56a54ac45fffad0437f7f97640"></a>

## is_pii property — Property reference / e380725c82d9 / 8

Type: `"bool"`. Computed.

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

<a id="canonical-5a5cb94734ae9df1f37bdc7c2ae0112495bea31a553cf157c23f9e0c1cc7caaf"></a>

<a id="canonical-5b5d6970e1aa6a9669c9da217b59695ebab45199a9bc1b9a8ad8e258f99f27aa"></a>

## is_sensitive_data property — Property reference / e380725c82d9 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-f54d0b47758825ba4b115d992f6b23f64cd090e60a0206d5259007061272a8a3"></a>

<a id="canonical-312e773986e03a66233723685f57e4fe931772f4abab107d28307e645f446347"></a>

## labels property — Property reference / e380725c82d9 / 10

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

<a id="canonical-e6a2bbfbf55a8df3ba8325f2afc869b801bfbac90d77436416aff84472ae27f0"></a>

<a id="canonical-558ae75664c0eb003604d8710494e39de7dcb9e60fb855288e9fca02e927e028"></a>

## name property — Property reference / e380725c82d9 / 11

Type: `"string"`. Required.

Name of the DataType.

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

<a id="canonical-30c10722a680c262d0d8bcbdf432ea3cbbc27119ad6a48944baf5cf2ddeca57f"></a>

<a id="canonical-7233e8c69bac6c19e3368460d96a4178a7b255251f110f0bc75f8a88c2c31dda"></a>

## namespace property — Property reference / e380725c82d9 / 12

Type: `"string"`. Required.

Namespace where the DataType exists.

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

- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5): complete subsection reference.

<a id="canonical-eeecbf67c778ab5fbce350c6a569eae26dc20c3a52dac6856b8dc9abea90b7f2"></a>

## All schema paths — Property reference / e380725c82d9 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--data_type--reference--group-001.md#canonical-901a440d452a6f419ecc4290cfe9a450646e6d3119c1cfeaf5f77ebf82d3a6ff) |
| `compliances` | [compliances](data-sources--data_type--reference--group-001.md#canonical-ad56eec22ff5822fe7116ff893f20abb4dc8c40b069a2b69a70aec4c2b499dcb) |
| `description` | [description](data-sources--data_type--reference--group-001.md#canonical-6cd2217fc3f92f432eb7c49c86a3a428fd130daaf13a9a3a7f8acb75850ea6fb) |
| `id` | [id](data-sources--data_type--reference--group-001.md#canonical-095c646d3966bd6ee0ca2b9407e0e4605649bb5bb91e49c86506298105f04f7d) |
| `is_pii` | [is_pii](data-sources--data_type--reference--group-001.md#canonical-4273659dbab11f6e27d2f7250bba976bad7ea2b6916f4288732aa6e7859db45b) |
| `is_sensitive_data` | [is_sensitive_data](data-sources--data_type--reference--group-001.md#canonical-5a5cb94734ae9df1f37bdc7c2ae0112495bea31a553cf157c23f9e0c1cc7caaf) |
| `labels` | [labels](data-sources--data_type--reference--group-001.md#canonical-f54d0b47758825ba4b115d992f6b23f64cd090e60a0206d5259007061272a8a3) |
| `name` | [name](data-sources--data_type--reference--group-001.md#canonical-e6a2bbfbf55a8df3ba8325f2afc869b801bfbac90d77436416aff84472ae27f0) |
| `namespace` | [namespace](data-sources--data_type--reference--group-001.md#canonical-30c10722a680c262d0d8bcbdf432ea3cbbc27119ad6a48944baf5cf2ddeca57f) |
| `rules` | [rules](data-sources--data_type--reference--group-001.md#canonical-522de0e87aa5f6bc3840d67a684d23f83bb51017873fb292b5bac6b871a8c7d2) |
| `rules.key_pattern` | [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-9780abe76feb0cbd77e78762b760e8c9e1c8bc1b4fdd383d8cacd2e2cdcce48d) |
| `rules.key_pattern.exact_values` | [rules.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-594d3a7380623830b28ba8f4c9af4a1af6fb22d1f4c1f013c007e61fe6112ab2) |
| `rules.key_pattern.exact_values.exact_values` | [rules.key_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-0fbb74bbf8a9e384bee4146dd305f67c27dc4c2af95f711ca7d7a26c8ad97ca6) |
| `rules.key_pattern.regex_value` | [rules.key_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-b0e9eeb3e37b00f645c0dc0cd7cbf8880f0a09bf2954cfbba6b12b7918dd15d8) |
| `rules.key_pattern.substring_value` | [rules.key_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-b4c76b0af6210838e40668618c4daaf269af7498dff61e63a71a6c0dcac715fc) |
| `rules.key_value_pattern` | [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-55ca3e98d6e2fae43f125a24262ffd56c54a21e1cb58b884ba6ffdd37faba01e) |
| `rules.key_value_pattern.key_pattern` | [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-a3648b8169dda02946c954de51f9f2b6b4e62cdc0aef6198c23bbaff0adb1ddb) |
| `rules.key_value_pattern.key_pattern.exact_values` | [rules.key_value_pattern.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-231ca90085e20378704cffe1b2fc73fa9407eff8fda4c5935ee5b7e6cc1f0da6) |
| `rules.key_value_pattern.key_pattern.exact_values.exact_values` | [rules.key_value_pattern.key_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-ea48ab28c1c23cbc93c893454121c0229c3789c896b993b8dc488e68afbd09ea) |
| `rules.key_value_pattern.key_pattern.regex_value` | [rules.key_value_pattern.key_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-36706bf9e7f74dca97ae15ceba73292ef1105a9257e0b2397dcf1fbb57d5bd4c) |
| `rules.key_value_pattern.key_pattern.substring_value` | [rules.key_value_pattern.key_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-2a7a53bfa8568592bacee36d443955fb88ca57e6a079351702d594b678901001) |
| `rules.key_value_pattern.value_pattern` | [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-8fbd5e01023381ba4eb53d44c71fcceba2bf1352b0e95f5a0a29f6319a002539) |
| `rules.key_value_pattern.value_pattern.exact_values` | [rules.key_value_pattern.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-5d5487d1dca1d5482a960c000147064e6db8ba22770ac3471975d5030861c4f2) |
| `rules.key_value_pattern.value_pattern.exact_values.exact_values` | [rules.key_value_pattern.value_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-96cebbb25838a31cad1c556a6dadb7fb4f6228b4d12ca8f2d180446510c7bf36) |
| `rules.key_value_pattern.value_pattern.regex_value` | [rules.key_value_pattern.value_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-0172faab36f167af82ef3a2e0638e4ddc1e95d78709672b04b5fb00451f97bbe) |
| `rules.key_value_pattern.value_pattern.substring_value` | [rules.key_value_pattern.value_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-e136f37fa077626ca147197c8ea9fcd4d7610ee85812328dc3d3832d8ee27b86) |
| `rules.value_pattern` | [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-d18eb0a5b1f018168696794882de186882b150e999b32ed69f9c738d4eff9a85) |
| `rules.value_pattern.exact_values` | [rules.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-9d3cb4014408e0ed4ea28356b303a98bcc3d5a5b30e7a557c7c3e2308c96307e) |
| `rules.value_pattern.exact_values.exact_values` | [rules.value_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-8c1ae8659e8c60c9e55e2703eefc66c7cf0f79fc55978b13bb71fd3cf848be92) |
| `rules.value_pattern.regex_value` | [rules.value_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-52553d49d58c3dba3cc5641c34a1e8efe907973ee36d56b97c3cae93651e8a04) |
| `rules.value_pattern.substring_value` | [rules.value_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-c403373323af4eff215c3998bacad0357dc2d6dc611b46f8d76efa056f3b3ee0) |

<a id="canonical-346400bba526d224582b872ca7e8778b3606b0a36ab854664650c37dbadb2101"></a>

## Next pages — Property reference / e380725c82d9 / 14

- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f7e1a30d60dfab07575acb6aeefb3a56f78fd04ce4521c89f4c76f997d95acf"></a>

## rules — rules / 2c7cd4e04133 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- rules

<a id="canonical-522de0e87aa5f6bc3840d67a684d23f83bb51017873fb292b5bac6b871a8c7d2"></a>

Type: `"list"`. Computed.

Configure key/value or regex match rules to enable the platform to detect this custom data type in
the API request or response.

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

<a id="canonical-3031e7e93eaac4b6ddc77911678883ab7e3634977b3da5f90a5ba7949771aa22"></a>

## Direct properties — rules / 2c7cd4e04133 / 3

- [key_pattern](data-sources--data_type--reference--group-001.md#canonical-19e07202991b74ebda8962965d807217d82918c45897af20acc279927162f727): complete subsection reference.

- [key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676): complete subsection reference.

- [value_pattern](data-sources--data_type--reference--group-001.md#canonical-b5846f4ed4a8b0094a1d0d31934effe8c3557fd2aabc1691b98f015eb1597faa): complete subsection reference.

<a id="canonical-3b9e4256633d5200dabf84ac189e324e7a2f1c7b5716ec7331adde312ca13d4a"></a>

## Next pages — rules / 2c7cd4e04133 / 4

- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-19e07202991b74ebda8962965d807217d82918c45897af20acc279927162f727)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-b5846f4ed4a8b0094a1d0d31934effe8c3557fd2aabc1691b98f015eb1597faa)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-19e07202991b74ebda8962965d807217d82918c45897af20acc279927162f727"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8bfeedef0ca41cf40f6d193769e0438b1f02fdb15cc12619a79880c707a7ced"></a>

## rules.key_pattern — rules.key_pattern / d8f0e6cc0495 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- rules.key_pattern

<a id="canonical-9780abe76feb0cbd77e78762b760e8c9e1c8bc1b4fdd383d8cacd2e2cdcce48d"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Upstream description:

Test

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

<a id="canonical-4a570724e21053405dea931725b5644710231bfe493aa962a5b07f61c8f0a9df"></a>

## Direct properties — rules.key_pattern / d8f0e6cc0495 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-5f264e901a36eaa85021940f0b6e06bf8748643f60753e9f4ccefb69932c2c5e): complete subsection reference.

<a id="canonical-b0e9eeb3e37b00f645c0dc0cd7cbf8880f0a09bf2954cfbba6b12b7918dd15d8"></a>

<a id="canonical-9bf59c0aea183e2bd155eb9efbf3f874314732502846848604fbd760fda5db42"></a>

## regex_value property — rules.key_pattern / d8f0e6cc0495 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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

<a id="canonical-b4c76b0af6210838e40668618c4daaf269af7498dff61e63a71a6c0dcac715fc"></a>

<a id="canonical-4c096c1c3fae1e77a637d665085a0d381f68052f6a8500926f87c4f2253c80c7"></a>

## substring_value property — rules.key_pattern / d8f0e6cc0495 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

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

<a id="canonical-ac23f37f1a06c20b272d6fd3c04868ac368ea569d97311fcd593aa28e8afb917"></a>

## Next pages — rules.key_pattern / d8f0e6cc0495 / 6

- [rules.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-5f264e901a36eaa85021940f0b6e06bf8748643f60753e9f4ccefb69932c2c5e)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-5f264e901a36eaa85021940f0b6e06bf8748643f60753e9f4ccefb69932c2c5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bae4b85fcb0d1b3cbc00868919822a0645750a4b7231a0598c353c82fd0d3260"></a>

## rules.key_pattern.exact_values — rules.key_pattern.exact_values / 7f252b466b5f / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-19e07202991b74ebda8962965d807217d82918c45897af20acc279927162f727)
- rules.key_pattern.exact_values

<a id="canonical-594d3a7380623830b28ba8f4c9af4a1af6fb22d1f4c1f013c007e61fe6112ab2"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-b42a7650eaa44d2b4897a1817108dc378f4793c358e3f6bc99fc8e9e5a7d5013"></a>

## Direct properties — rules.key_pattern.exact_values / 7f252b466b5f / 3

<a id="canonical-0fbb74bbf8a9e384bee4146dd305f67c27dc4c2af95f711ca7d7a26c8ad97ca6"></a>

<a id="canonical-db015aae374548921e3bb743adea3234dd6e2eaa7516500e3880cae5ea6b32c7"></a>

## exact_values property — rules.key_pattern.exact_values / 7f252b466b5f / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-a10337cdae292127e1a9233cf9bebe453cbab9f9fec8441b333c0ea89182db63"></a>

## Next pages — rules.key_pattern.exact_values / 7f252b466b5f / 5

- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-19e07202991b74ebda8962965d807217d82918c45897af20acc279927162f727)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86ed3bc7d35195e9c5cd859ce3c410d4bdda31d865dace551ef410db86809021"></a>

## rules.key_value_pattern — rules.key_value_pattern / 25a33db83f22 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- rules.key_value_pattern

<a id="canonical-55ca3e98d6e2fae43f125a24262ffd56c54a21e1cb58b884ba6ffdd37faba01e"></a>

Type: `"single"`. Computed.

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

<a id="canonical-392dd617b0237d9d593b1368b171968e23dfd6b4672eec63f2e9f51e0eddcd2f"></a>

## Direct properties — rules.key_value_pattern / 25a33db83f22 / 3

- [key_pattern](data-sources--data_type--reference--group-001.md#canonical-557652934a1e34cbca2bb46fd6092051c18aa123de029812c71c582c4095bf80): complete subsection reference.

- [value_pattern](data-sources--data_type--reference--group-001.md#canonical-ae5e013ba5d17e14cdb36fdaea134df4f0cb4d8b3af21af2a7d04ace699106a5): complete subsection reference.

<a id="canonical-ccf8e5b22e0a4210506dd786b9bfa937d5be0e56af7064ffaa2a913f1d17a606"></a>

## Next pages — rules.key_value_pattern / 25a33db83f22 / 4

- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-557652934a1e34cbca2bb46fd6092051c18aa123de029812c71c582c4095bf80)
- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-ae5e013ba5d17e14cdb36fdaea134df4f0cb4d8b3af21af2a7d04ace699106a5)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-557652934a1e34cbca2bb46fd6092051c18aa123de029812c71c582c4095bf80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdbfc1a246e2f6e43523d65597d461c8a1f67f42e891921956a5ab9ae4098cb9"></a>

## rules.key_value_pattern.key_pattern — rules.key_value_pattern.key_pattern / 53c9c2b512f5 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- rules.key_value_pattern.key_pattern

<a id="canonical-a3648b8169dda02946c954de51f9f2b6b4e62cdc0aef6198c23bbaff0adb1ddb"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Upstream description:

Test

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

<a id="canonical-edeee704650d9b17a06048d36a5186b46417d7b6df56e3db666ca8cf8d972c7e"></a>

## Direct properties — rules.key_value_pattern.key_pattern / 53c9c2b512f5 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-44eebcdb167fd2cfe4ea9d76d9174aba3ebd5db84f382060d7083ea71cfea505): complete subsection reference.

<a id="canonical-36706bf9e7f74dca97ae15ceba73292ef1105a9257e0b2397dcf1fbb57d5bd4c"></a>

<a id="canonical-3e1d52110bee8aa8630c176c2a20ab16409cce8e828869e2e143b5aca8c56e18"></a>

## regex_value property — rules.key_value_pattern.key_pattern / 53c9c2b512f5 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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

<a id="canonical-2a7a53bfa8568592bacee36d443955fb88ca57e6a079351702d594b678901001"></a>

<a id="canonical-ed17f667fcc18fbc336117475c5aacc4a8e6b00867d0ca52b42b3f7b40a19947"></a>

## substring_value property — rules.key_value_pattern.key_pattern / 53c9c2b512f5 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

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

<a id="canonical-25cd4c4212522799cff5055be34b75efdb97c7664c07ccb1711dc871badb6e30"></a>

## Next pages — rules.key_value_pattern.key_pattern / 53c9c2b512f5 / 6

- [rules.key_value_pattern.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-44eebcdb167fd2cfe4ea9d76d9174aba3ebd5db84f382060d7083ea71cfea505)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-44eebcdb167fd2cfe4ea9d76d9174aba3ebd5db84f382060d7083ea71cfea505"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67d1f1c6fecaef79094cf402fe91c01fda4dadba4c2522d5f1fe5a25b882ef50"></a>

## rules.key_value_pattern.key_pattern.exact_values — rules.key_value_pattern.key_pattern.exact_values / e7ccedb1ee43 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-557652934a1e34cbca2bb46fd6092051c18aa123de029812c71c582c4095bf80)
- rules.key_value_pattern.key_pattern.exact_values

<a id="canonical-231ca90085e20378704cffe1b2fc73fa9407eff8fda4c5935ee5b7e6cc1f0da6"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-813d306a480485ba539bb4d4e361338eb889b6487dc903769c32b15882a0be7f"></a>

## Direct properties — rules.key_value_pattern.key_pattern.exact_values / e7ccedb1ee43 / 3

<a id="canonical-ea48ab28c1c23cbc93c893454121c0229c3789c896b993b8dc488e68afbd09ea"></a>

<a id="canonical-6c2a8cb48dcc708da546afb470415057d50de6af93387c110ee494c3370b1cff"></a>

## exact_values property — rules.key_value_pattern.key_pattern.exact_values / e7ccedb1ee43 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-f13d2640f8f4c3427f7ddf593eb5bbb8da096aa50906a516edcea0373c0ff339"></a>

## Next pages — rules.key_value_pattern.key_pattern.exact_values / e7ccedb1ee43 / 5

- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-557652934a1e34cbca2bb46fd6092051c18aa123de029812c71c582c4095bf80)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-ae5e013ba5d17e14cdb36fdaea134df4f0cb4d8b3af21af2a7d04ace699106a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac9e6c63755dc1fbe7e4bfc356429c5c6187b2dc3cbf0f16347fcc2a1e3c4e18"></a>

## rules.key_value_pattern.value_pattern — rules.key_value_pattern.value_pattern / 45f4e1a976ba / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- rules.key_value_pattern.value_pattern

<a id="canonical-8fbd5e01023381ba4eb53d44c71fcceba2bf1352b0e95f5a0a29f6319a002539"></a>

Type: `"single"`. Computed.

Configuration parameter for value pattern.

Upstream description:

Test

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

<a id="canonical-71768f72d4b761d495d9369bcfc497df3b8b5d5322d10324f4d8b3a006a57e4b"></a>

## Direct properties — rules.key_value_pattern.value_pattern / 45f4e1a976ba / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-73a7d1f26cc7636ac3af300d94a888ccfe85cba0a37e7172b1fa3ba03fb2da46): complete subsection reference.

<a id="canonical-0172faab36f167af82ef3a2e0638e4ddc1e95d78709672b04b5fb00451f97bbe"></a>

<a id="canonical-0525f5c3f482208c82ff4bb7c62f761229150a62b38f8dea77bb965486baabb5"></a>

## regex_value property — rules.key_value_pattern.value_pattern / 45f4e1a976ba / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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

<a id="canonical-e136f37fa077626ca147197c8ea9fcd4d7610ee85812328dc3d3832d8ee27b86"></a>

<a id="canonical-9f0ab186f75db1f3d1864eaca423ab328833846de148b4d8b637c76f9eb3d860"></a>

## substring_value property — rules.key_value_pattern.value_pattern / 45f4e1a976ba / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

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

<a id="canonical-45422afee4ab4fe3a4519a83ae2154280c623005d431e8f3ce70d3172c3b2f37"></a>

## Next pages — rules.key_value_pattern.value_pattern / 45f4e1a976ba / 6

- [rules.key_value_pattern.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-73a7d1f26cc7636ac3af300d94a888ccfe85cba0a37e7172b1fa3ba03fb2da46)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-73a7d1f26cc7636ac3af300d94a888ccfe85cba0a37e7172b1fa3ba03fb2da46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d82463d5d11f6bdf6c7a5cfa68af00ce2838e7025359a0365675fa4a25b43cc"></a>

## rules.key_value_pattern.value_pattern.exact_values — rules.key_value_pattern.value_pattern.exact_values / b8a00a7b740c / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1275c1fa7e74ef8097b64c6dcb421f48b55d3fd990d29fcf7e3bf9566aade676)
- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-ae5e013ba5d17e14cdb36fdaea134df4f0cb4d8b3af21af2a7d04ace699106a5)
- rules.key_value_pattern.value_pattern.exact_values

<a id="canonical-5d5487d1dca1d5482a960c000147064e6db8ba22770ac3471975d5030861c4f2"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-3863ccda23f07c2c866c5ed2c704815479628c0003484f8ec46cc3d205d5bb9c"></a>

## Direct properties — rules.key_value_pattern.value_pattern.exact_values / b8a00a7b740c / 3

<a id="canonical-96cebbb25838a31cad1c556a6dadb7fb4f6228b4d12ca8f2d180446510c7bf36"></a>

<a id="canonical-551655f65700b5b36c38c33b8813550148c2315cdaf75c04bc4290aa8d0e5cb9"></a>

## exact_values property — rules.key_value_pattern.value_pattern.exact_values / b8a00a7b740c / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-f952d5d22ac38bf80a0843fc858c5a49062b84685a4f97f910ce5044b65586fb"></a>

## Next pages — rules.key_value_pattern.value_pattern.exact_values / b8a00a7b740c / 5

- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-ae5e013ba5d17e14cdb36fdaea134df4f0cb4d8b3af21af2a7d04ace699106a5)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-b5846f4ed4a8b0094a1d0d31934effe8c3557fd2aabc1691b98f015eb1597faa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dbf941b61931ed733eebc2d954878657c887df8682a2735553bdece8d8e1fed"></a>

## rules.value_pattern — rules.value_pattern / 0a699ec87f7e / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- rules.value_pattern

<a id="canonical-d18eb0a5b1f018168696794882de186882b150e999b32ed69f9c738d4eff9a85"></a>

Type: `"single"`. Computed.

Configuration parameter for value pattern.

Upstream description:

Test

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

<a id="canonical-f7307293d85accfdfdc15d50b6adabdf9ac9f7defe0a472bc4e0af64edbfb95d"></a>

## Direct properties — rules.value_pattern / 0a699ec87f7e / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-2aa857953cd49c5123d43981c12e10135a7f5e205c59569c334058b9fea7543c): complete subsection reference.

<a id="canonical-52553d49d58c3dba3cc5641c34a1e8efe907973ee36d56b97c3cae93651e8a04"></a>

<a id="canonical-9be9971ed0e15dfd0fd5846f15170fdf67312ed3caeda8b7ef7937043d24d555"></a>

## regex_value property — rules.value_pattern / 0a699ec87f7e / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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

<a id="canonical-c403373323af4eff215c3998bacad0357dc2d6dc611b46f8d76efa056f3b3ee0"></a>

<a id="canonical-17dd2fce5970b80926206c6a2c36a568b7d89456759ffa8c8e7d0cbdf8f86215"></a>

## substring_value property — rules.value_pattern / 0a699ec87f7e / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

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

<a id="canonical-afd557c7098cb1adc53768fde05239507601e1ac53bb2cbc9e1482959f74ae49"></a>

## Next pages — rules.value_pattern / 0a699ec87f7e / 6

- [rules.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-2aa857953cd49c5123d43981c12e10135a7f5e205c59569c334058b9fea7543c)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-2aa857953cd49c5123d43981c12e10135a7f5e205c59569c334058b9fea7543c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9e8572ec881106abe896ec8ecbdadbe45659a0366317d8e88367cec98f31f7f"></a>

## rules.value_pattern.exact_values — rules.value_pattern.exact_values / 0dbf028c52f6 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [rules](data-sources--data_type--reference--group-001.md#canonical-3ffd25fe5c8d59ec3a97501ada8781cf4e0b4558e779ccb9739b38994084a9f5)
- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-b5846f4ed4a8b0094a1d0d31934effe8c3557fd2aabc1691b98f015eb1597faa)
- rules.value_pattern.exact_values

<a id="canonical-9d3cb4014408e0ed4ea28356b303a98bcc3d5a5b30e7a557c7c3e2308c96307e"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-5451792eaa754f3b64718debf620fae0d32d0634e6b6999486090cb6745f8721"></a>

## Direct properties — rules.value_pattern.exact_values / 0dbf028c52f6 / 3

<a id="canonical-8c1ae8659e8c60c9e55e2703eefc66c7cf0f79fc55978b13bb71fd3cf848be92"></a>

<a id="canonical-28e685d7d1548a8344744159a90dfd44167967101c07ade0867087740c19a158"></a>

## exact_values property — rules.value_pattern.exact_values / 0dbf028c52f6 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-80f7218344a80fc68ad245be2e80901d77c7b386cd4d37e47826e07c2335a281"></a>

## Next pages — rules.value_pattern.exact_values / 0dbf028c52f6 / 5

- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-b5846f4ed4a8b0094a1d0d31934effe8c3557fd2aabc1691b98f015eb1597faa)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
