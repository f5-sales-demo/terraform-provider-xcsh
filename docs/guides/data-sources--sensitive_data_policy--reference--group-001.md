---
page_title: "xcsh_sensitive_data_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy reference."
---

# xcsh_sensitive_data_policy reference

<a id="canonical-86a29b3194d50ff67495d3075172b43382c2563aef377018db9add9b0f5caf8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f223013e2c4ac7f5c8fa38e6b9cb5666db9cb7e93d0e47e643e37bcae399e1a1"></a>

## Property reference — Property reference / bf3b7bb52959 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
- Property reference

<a id="canonical-f2a70a613ce32eb8f796cbdaba31e05419e1d0061e01d9890c7510a8b660a277"></a>

## Direct properties — Property reference / bf3b7bb52959 / 3

<a id="canonical-39311ab26048b056fa105a7b44b30f0412fbad238bfa5632ed1cba4a5febca11"></a>

<a id="canonical-37b0f4ec529fd1debd3f8e5daa0ce39eb76b2eedb2f08da700d083043a5b25ab"></a>

## annotations property — Property reference / bf3b7bb52959 / 4

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

<a id="canonical-41f5bbf268656c9b651daad22bc39389c4bbf4bc247c46c255274fc8e9701545"></a>

<a id="canonical-5aa0a8fc784a5a8a53c95019f4f2c38c2d6d321aeab74140296db0c5c354df3e"></a>

## compliances property — Property reference / bf3b7bb52959 / 5

Type: `["list", "string"]`. Computed.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Select relevant compliance frameworks, such as GDPR, HIPAA, or PCI-DSS, to ensure monitoring under
your sensitive data discovery. Defaults to \`\[\]\`. Server applies default when omitted. Possible
values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`, \`APPI\`, \`HIPAA\`,
\`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`, \`ISO\_IEC\_27701\`,
\`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Upstream description:

Select relevant compliance frameworks, such as GDPR, HIPAA, or PCI-DSS, to ensure monitoring under
your sensitive data discovery.

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

- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-a7f7a7d237b7d31824e7c600e33a923f895c17509e2e63925b8d3eb86c184cd6): complete subsection reference.

<a id="canonical-855faf59417629a6465d94fa712456f85e2338cdc480a0a30d96160246db2934"></a>

<a id="canonical-26ea07a6f633d5346e9b0e283db4b93652543c3b53b06cddff4667968b23ec3b"></a>

## description property — Property reference / bf3b7bb52959 / 6

Type: `"string"`. Computed.

Description of the SensitiveDataPolicy.

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

<a id="canonical-7ad9ac2ad29ae55670c66abfffca2c54f484685231942fd476ae03702883dc3e"></a>

<a id="canonical-82bf5cbca45ae1bb6581ee83e35abbe085c2a63712be5dd96a00302b769f2dd6"></a>

## disabled_predefined_data_types property — Property reference / bf3b7bb52959 / 7

Type: `["list", "string"]`. Computed.

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f1534dc013fa4267876374b9b0d3cdd3fcfb60d82036caaf813a5176d64a544e"></a>

<a id="canonical-3e40d06efe0722fb870d38f31b300a8ca1758b0481766d3b955c18f99eefa65c"></a>

## id property — Property reference / bf3b7bb52959 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-12e94636b93871c2da9876330d82b922e82d76cf9ec650ffb2a3afdae28b0843"></a>

<a id="canonical-7cb5a344626f76a9fa98e73a859a6cf02797048fb3bbbe38e3554b3f7b6dfbec"></a>

## labels property — Property reference / bf3b7bb52959 / 9

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

<a id="canonical-cf358b91196570f9ca26cee6e0ba7728dc3f5c450d3959ed708ce39b72aebaa6"></a>

<a id="canonical-f957e16a84e100190c66d8fd59f7257e363996889208c2af990845920a8210d1"></a>

## name property — Property reference / bf3b7bb52959 / 10

Type: `"string"`. Required.

Name of the SensitiveDataPolicy.

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

<a id="canonical-2fd0b7e8b846cb5424c21ce6c9a07eba7b24ff56b5446ecf5332b94bf90e36ee"></a>

<a id="canonical-84137561aeed5b58c8138ff5f2723b3dfae3a982b520a44ca4b51453d4e86e1b"></a>

## namespace property — Property reference / bf3b7bb52959 / 11

Type: `"string"`. Required.

Namespace where the SensitiveDataPolicy exists.

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

<a id="canonical-081a673b4a5ee2ec5d03740d659d18fb8550e8d99e460754897f10f9ffe1c53c"></a>

## All schema paths — Property reference / bf3b7bb52959 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--sensitive_data_policy--reference--group-001.md#canonical-39311ab26048b056fa105a7b44b30f0412fbad238bfa5632ed1cba4a5febca11) |
| `compliances` | [compliances](data-sources--sensitive_data_policy--reference--group-001.md#canonical-41f5bbf268656c9b651daad22bc39389c4bbf4bc247c46c255274fc8e9701545) |
| `custom_data_types` | [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-c3b500dbddacac92dde89fa221da3d977dee2e78d6bc7e5c86573122cf51ebff) |
| `custom_data_types.custom_data_type_ref` | [custom_data_types.custom_data_type_ref](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2bf4843e3210173f122cd9624a1f766636b2325713b1c890f8aff1dbadc585b8) |
| `custom_data_types.custom_data_type_ref.name` | [custom_data_types.custom_data_type_ref.name](data-sources--sensitive_data_policy--reference--group-001.md#canonical-604c35873ca69246f981e27cccc623bdb7ae2804c0369723ee6f0060dbf7f4d6) |
| `custom_data_types.custom_data_type_ref.namespace` | [custom_data_types.custom_data_type_ref.namespace](data-sources--sensitive_data_policy--reference--group-001.md#canonical-9195f19eabd5feece66c512c8e46f9baef2e4415099f2c479be528ebd4e7da12) |
| `custom_data_types.custom_data_type_ref.tenant` | [custom_data_types.custom_data_type_ref.tenant](data-sources--sensitive_data_policy--reference--group-001.md#canonical-320971a6f4c657b59673e44d632a4fcd05d63da97e44d1a716fcdd4657d0ab7a) |
| `description` | [description](data-sources--sensitive_data_policy--reference--group-001.md#canonical-855faf59417629a6465d94fa712456f85e2338cdc480a0a30d96160246db2934) |
| `disabled_predefined_data_types` | [disabled_predefined_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-7ad9ac2ad29ae55670c66abfffca2c54f484685231942fd476ae03702883dc3e) |
| `id` | [id](data-sources--sensitive_data_policy--reference--group-001.md#canonical-f1534dc013fa4267876374b9b0d3cdd3fcfb60d82036caaf813a5176d64a544e) |
| `labels` | [labels](data-sources--sensitive_data_policy--reference--group-001.md#canonical-12e94636b93871c2da9876330d82b922e82d76cf9ec650ffb2a3afdae28b0843) |
| `name` | [name](data-sources--sensitive_data_policy--reference--group-001.md#canonical-cf358b91196570f9ca26cee6e0ba7728dc3f5c450d3959ed708ce39b72aebaa6) |
| `namespace` | [namespace](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2fd0b7e8b846cb5424c21ce6c9a07eba7b24ff56b5446ecf5332b94bf90e36ee) |

<a id="canonical-40e33c00695da0f347ce26b5a913fa2a8cbaddbb52c047d961b2ea62859cc695"></a>

## Next pages — Property reference / bf3b7bb52959 / 13

- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-a7f7a7d237b7d31824e7c600e33a923f895c17509e2e63925b8d3eb86c184cd6)
- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)

<a id="canonical-a7f7a7d237b7d31824e7c600e33a923f895c17509e2e63925b8d3eb86c184cd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d2d3550371dc056546528b1b1cf5206023a3fe111bba912824467defbc98217"></a>

## custom_data_types — custom_data_types / 42f8c7b87415 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
- [Property reference](data-sources--sensitive_data_policy--reference--group-001.md#canonical-86a29b3194d50ff67495d3075172b43382c2563aef377018db9add9b0f5caf8f)
- custom_data_types

<a id="canonical-c3b500dbddacac92dde89fa221da3d977dee2e78d6bc7e5c86573122cf51ebff"></a>

Type: `"list"`. Computed.

Select your custom data types to be monitored in the API discovery. Defaults to \`\[\]\`. Server
applies default when omitted.

Upstream description:

Select your custom data types to be monitored in the API discovery.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e345879dfac18391a879d890bcb286007cb83b4ddff0125d6ef3e619f5d50c07"></a>

## Direct properties — custom_data_types / 42f8c7b87415 / 3

- [custom_data_type_ref](data-sources--sensitive_data_policy--reference--group-001.md#canonical-7603a1a30ba822ecd2f0058e15b6cc221c07c15967101be30db90625f2545a10): complete subsection reference.

<a id="canonical-5e0938503f4df874aa0b4de71289705652ae3367be298d729cb60412d59016c1"></a>

## Next pages — custom_data_types / 42f8c7b87415 / 4

- [custom_data_types.custom_data_type_ref](data-sources--sensitive_data_policy--reference--group-001.md#canonical-7603a1a30ba822ecd2f0058e15b6cc221c07c15967101be30db90625f2545a10)
- [Property reference](data-sources--sensitive_data_policy--reference--group-001.md#canonical-86a29b3194d50ff67495d3075172b43382c2563aef377018db9add9b0f5caf8f)
- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)

<a id="canonical-7603a1a30ba822ecd2f0058e15b6cc221c07c15967101be30db90625f2545a10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef530977e51cf19734e4f852b6a3ba267cb24a427623c702bda533760f7758c7"></a>

## custom_data_types.custom_data_type_ref — custom_data_types.custom_data_type_ref / fe92342f70e9 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
- [Property reference](data-sources--sensitive_data_policy--reference--group-001.md#canonical-86a29b3194d50ff67495d3075172b43382c2563aef377018db9add9b0f5caf8f)
- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-a7f7a7d237b7d31824e7c600e33a923f895c17509e2e63925b8d3eb86c184cd6)
- custom_data_types.custom_data_type_ref

<a id="canonical-2bf4843e3210173f122cd9624a1f766636b2325713b1c890f8aff1dbadc585b8"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-89a4ffa94e8891d319df5ada919f0a840247e23d57b61a3059b1044fa0dc2aec"></a>

## Direct properties — custom_data_types.custom_data_type_ref / fe92342f70e9 / 3

<a id="canonical-604c35873ca69246f981e27cccc623bdb7ae2804c0369723ee6f0060dbf7f4d6"></a>

<a id="canonical-9b8f48973bd79abf59a6efe47295dda16b4b3e4cc3d5c67280a8e1daf5bc955d"></a>

## name property — custom_data_types.custom_data_type_ref / fe92342f70e9 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-9195f19eabd5feece66c512c8e46f9baef2e4415099f2c479be528ebd4e7da12"></a>

<a id="canonical-3d7131904d62c76b358dd586c73426cb581bce56e441dd1bf7468c8f1d682b03"></a>

## namespace property — custom_data_types.custom_data_type_ref / fe92342f70e9 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-320971a6f4c657b59673e44d632a4fcd05d63da97e44d1a716fcdd4657d0ab7a"></a>

<a id="canonical-65375dc104045134a0049b648d27fd0c596afddc2ef708749436fe7a7127a189"></a>

## tenant property — custom_data_types.custom_data_type_ref / fe92342f70e9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0c68144085a9ed7a832438cab64270c11fbf5f1735ee64d80bae5ddc26a9e431"></a>

## Next pages — custom_data_types.custom_data_type_ref / fe92342f70e9 / 7

- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-a7f7a7d237b7d31824e7c600e33a923f895c17509e2e63925b8d3eb86c184cd6)
- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
