---
page_title: "xcsh_sensitive_data_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy reference."
---

# xcsh_sensitive_data_policy reference

<a id="canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7195e542683289b421aa8d7ec0025b6dd117cc222f53d206f2002e44d225334"></a>

## Property reference — Property reference / 91d065e86ce5 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- Property reference

<a id="canonical-5d8f3ac5d9b236d08a0707ad95fdb342c6d3f0aecc4f7042d5d354b802941438"></a>

## Direct properties — Property reference / 91d065e86ce5 / 3

<a id="canonical-34ace15c44d23a6e412e5b4dfaf93c9bc885be24b37f396f29fda924ca1cdc3b"></a>

<a id="canonical-758ff3bfdc9c0b3e95405f8824f762ac8aa657726f57ece59a8aae8ce59b3075"></a>

## annotations property — Property reference / 91d065e86ce5 / 4

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

<a id="canonical-15c7bab1b2d0f0b8c1811de902c5700f3c8cd0b302db155c1a373f9afe4ff94c"></a>

<a id="canonical-f2ede3c21e0814ca7485316dba6ba54f93dd1ac313515656fc1b7b42c7f90de5"></a>

## compliances property — Property reference / 91d065e86ce5 / 5

Type: `["list", "string"]`. Optional, Computed.

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

- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-90fdffbb81caf46084230c201cd4e42f6f8c408ef20d6c5635cfd2fc7d4190ad): complete subsection reference.

<a id="canonical-e01068a3af25aa7c2ea5050fe793cfad381b0e170280225bd30121b269abb88f"></a>

<a id="canonical-3a55cab961d94c34a9f96ea9da4d1ff01354603c38e8bfb962b416f197195d46"></a>

## description property — Property reference / 91d065e86ce5 / 6

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

<a id="canonical-894877ece852b67e9073286d0f3d1ed1528eda929a0ffc54fa82ba7cbda58602"></a>

<a id="canonical-b15684157270d2884fa85bc302801d6b60cc1fab1db0d7ceaf1dc43972b3af02"></a>

## disable property — Property reference / 91d065e86ce5 / 7

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

<a id="canonical-9e44e09afe760ab22e737e1b49fd89e04744cbf68724b52fba95f746b5304de6"></a>

<a id="canonical-275c0701335f73c8c0bba606bec003770bfd2d884b20882284899e5fd1c05893"></a>

## disabled_predefined_data_types property — Property reference / 91d065e86ce5 / 8

Type: `["list", "string"]`. Optional, Computed.

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(100),
}
```

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

<a id="canonical-f666b46fe85961c2d0d03eb0ebfdecf8737cc44c49a72c9ce739ff13399bf66d"></a>

<a id="canonical-09f243fabc02644385d57facbb7edf726cd65797d5bf23bc0d627034bb460466"></a>

## id property — Property reference / 91d065e86ce5 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-37bd908456d8c67d8043955607721b743b3175baf95c9886b1bf77e2a34a5dca"></a>

<a id="canonical-ce925dff06ee37b25e8faba03edc3e500a634b4e810e6387f803e5bc81cd10b0"></a>

## labels property — Property reference / 91d065e86ce5 / 10

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

<a id="canonical-ac2b3e88c7a8ee666d417f022aa1b70df05a62a4efb417af19ea67c170f64d2a"></a>

<a id="canonical-bb72b95bc9814cc4db3aa7290a1558f6ef771cf4680f4586b46001aed7c7b832"></a>

## name property — Property reference / 91d065e86ce5 / 11

Type: `"string"`. Required.

Name of the Sensitive Data Policy. Must be unique within the namespace.

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

<a id="canonical-784d0c879db045a1ce7b5f6d9c9a215463f86a75bc26ef77feb3c1d521b4b67c"></a>

<a id="canonical-30b063befa793955c4918484b676bb3797e51ec6b3d5364f5b343062fe2d6aab"></a>

## namespace property — Property reference / 91d065e86ce5 / 12

Type: `"string"`. Required.

Namespace where the Sensitive Data Policy is created.

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

- [timeouts](resources--sensitive_data_policy--reference--group-001.md#canonical-697bad1912f10d4df831518e879c8fe69fbb329ae06873fa5faced9e5f334fc6): complete subsection reference.

<a id="canonical-ac99c961406c51495139829a9028bdaae48f560d607147104b7e73855f76261a"></a>

## All schema paths — Property reference / 91d065e86ce5 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--sensitive_data_policy--reference--group-001.md#canonical-34ace15c44d23a6e412e5b4dfaf93c9bc885be24b37f396f29fda924ca1cdc3b) |
| `compliances` | [compliances](resources--sensitive_data_policy--reference--group-001.md#canonical-15c7bab1b2d0f0b8c1811de902c5700f3c8cd0b302db155c1a373f9afe4ff94c) |
| `custom_data_types` | [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-b7afd304aa44f4ab9cf24b7d41153973ac3ffb789c3fe27cd73096ef5e9ac06c) |
| `custom_data_types.custom_data_type_ref` | [custom_data_types.custom_data_type_ref](resources--sensitive_data_policy--reference--group-001.md#canonical-90fa7a18c0fc683b9408356d3a1db8e3a5ed5a5b5916a4625dee95c5f46e9f89) |
| `custom_data_types.custom_data_type_ref.name` | [custom_data_types.custom_data_type_ref.name](resources--sensitive_data_policy--reference--group-001.md#canonical-ace8071f167c98efc76f7b536c674c2b9bebcff687fec2d596f81809979c8b6f) |
| `custom_data_types.custom_data_type_ref.namespace` | [custom_data_types.custom_data_type_ref.namespace](resources--sensitive_data_policy--reference--group-001.md#canonical-0dd34f4d1178c47ca0a6a2f8e53a2976d5471c4c2ae34d90725fadb91e4a95d5) |
| `custom_data_types.custom_data_type_ref.tenant` | [custom_data_types.custom_data_type_ref.tenant](resources--sensitive_data_policy--reference--group-001.md#canonical-12508c4a6f7b9667cb084e1a0636f683ead55762407f2f31b6e5ac5225dbaf72) |
| `description` | [description](resources--sensitive_data_policy--reference--group-001.md#canonical-e01068a3af25aa7c2ea5050fe793cfad381b0e170280225bd30121b269abb88f) |
| `disable` | [disable](resources--sensitive_data_policy--reference--group-001.md#canonical-894877ece852b67e9073286d0f3d1ed1528eda929a0ffc54fa82ba7cbda58602) |
| `disabled_predefined_data_types` | [disabled_predefined_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-9e44e09afe760ab22e737e1b49fd89e04744cbf68724b52fba95f746b5304de6) |
| `id` | [id](resources--sensitive_data_policy--reference--group-001.md#canonical-f666b46fe85961c2d0d03eb0ebfdecf8737cc44c49a72c9ce739ff13399bf66d) |
| `labels` | [labels](resources--sensitive_data_policy--reference--group-001.md#canonical-37bd908456d8c67d8043955607721b743b3175baf95c9886b1bf77e2a34a5dca) |
| `name` | [name](resources--sensitive_data_policy--reference--group-001.md#canonical-ac2b3e88c7a8ee666d417f022aa1b70df05a62a4efb417af19ea67c170f64d2a) |
| `namespace` | [namespace](resources--sensitive_data_policy--reference--group-001.md#canonical-784d0c879db045a1ce7b5f6d9c9a215463f86a75bc26ef77feb3c1d521b4b67c) |
| `timeouts` | [timeouts](resources--sensitive_data_policy--reference--group-001.md#canonical-f857c91e432396dea8fed5264c59a430a40d0cbf59d19768481f4e715f722156) |
| `timeouts.create` | [timeouts.create](resources--sensitive_data_policy--reference--group-001.md#canonical-56a1528ea6da0a79b116e34c153fa51770a770de2bfc122a928d1786cb8e0f38) |
| `timeouts.delete` | [timeouts.delete](resources--sensitive_data_policy--reference--group-001.md#canonical-a7b4394492ac8768f4bb485ae78405ef9f20198f15d03eec504e34fe2e4a64df) |
| `timeouts.read` | [timeouts.read](resources--sensitive_data_policy--reference--group-001.md#canonical-9a0571df4547a71b73c1c9dc17eb4d2a0c54f9d4aeccc30cfe13069ddc375511) |
| `timeouts.update` | [timeouts.update](resources--sensitive_data_policy--reference--group-001.md#canonical-689b7e964a39f980552acc9f6978421cbbc46219abe269ef9ffac5a5d4a805c1) |

<a id="canonical-da224e2b5f48439e96f2b230b623af7c91f185c742fd3dd11f2f118290f1109c"></a>

## Next pages — Property reference / 91d065e86ce5 / 14

- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-90fdffbb81caf46084230c201cd4e42f6f8c408ef20d6c5635cfd2fc7d4190ad)
- [timeouts](resources--sensitive_data_policy--reference--group-001.md#canonical-697bad1912f10d4df831518e879c8fe69fbb329ae06873fa5faced9e5f334fc6)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)

<a id="canonical-90fdffbb81caf46084230c201cd4e42f6f8c408ef20d6c5635cfd2fc7d4190ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c94f62e62e58af45edffd9212415092b48ae36e98db4204e5b4e657db79ff95"></a>

## custom_data_types — custom_data_types / 64943356f5d5 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- custom_data_types

<a id="canonical-b7afd304aa44f4ab9cf24b7d41153973ac3ffb789c3fe27cd73096ef5e9ac06c"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
custom_data_types {
  # Configure direct properties listed below.
}
```

<a id="canonical-176a18059ab30509460f5221f9293f44954f060ad4c4770261a54a996eee0ad4"></a>

## Direct properties — custom_data_types / 64943356f5d5 / 3

- [custom_data_type_ref](resources--sensitive_data_policy--reference--group-001.md#canonical-411b729972bc037abc7455368b0abbbe753244ac6ed185ead73c39ffa8d18c80): complete subsection reference.

<a id="canonical-de3377ce4a82e5ddbbe1552cc08c94ce51ec37ee2bb0c9c06034b1aba0d5232c"></a>

## Next pages — custom_data_types / 64943356f5d5 / 4

- [custom_data_types.custom_data_type_ref](resources--sensitive_data_policy--reference--group-001.md#canonical-411b729972bc037abc7455368b0abbbe753244ac6ed185ead73c39ffa8d18c80)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)

<a id="canonical-411b729972bc037abc7455368b0abbbe753244ac6ed185ead73c39ffa8d18c80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2771e7f80762eaf26f4a3891670ad6504efbf3083434f7c0b894aeaa3b2ed13"></a>

## custom_data_types.custom_data_type_ref — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-90fdffbb81caf46084230c201cd4e42f6f8c408ef20d6c5635cfd2fc7d4190ad)
- custom_data_types.custom_data_type_ref

<a id="canonical-90fa7a18c0fc683b9408356d3a1db8e3a5ed5a5b5916a4625dee95c5f46e9f89"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
custom_data_type_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c6bc9c76c7369e04bf891a774a51192cc45c011525cacffed96ab85bbc8c36d"></a>

## Direct properties — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 3

<a id="canonical-ace8071f167c98efc76f7b536c674c2b9bebcff687fec2d596f81809979c8b6f"></a>

<a id="canonical-27463f736154ff65291de13c97231fc6797e59063dd2b5e2e71f9ed17a7faa2e"></a>

## name property — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0dd34f4d1178c47ca0a6a2f8e53a2976d5471c4c2ae34d90725fadb91e4a95d5"></a>

<a id="canonical-bfbc096e2daf3d4768492de934a10baa262cb4eecc15e7985704f030399423c0"></a>

## namespace property — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 5

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
}
```

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

<a id="canonical-12508c4a6f7b9667cb084e1a0636f683ead55762407f2f31b6e5ac5225dbaf72"></a>

<a id="canonical-7b9b142e6364d8ef8e4aead949e3bcc797f11328105ca96b078c70beeb435623"></a>

## tenant property — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-aee52adb169b4da3e52505a2557724a82fe13d71808a210074e6622f94a3a1f3"></a>

## Next pages — custom_data_types.custom_data_type_ref / cd9b9ec4206c / 7

- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-90fdffbb81caf46084230c201cd4e42f6f8c408ef20d6c5635cfd2fc7d4190ad)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)

<a id="canonical-697bad1912f10d4df831518e879c8fe69fbb329ae06873fa5faced9e5f334fc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38bcffc0c13942c2afd29b5e4f11cc1c7af61eda14516dbaf402c6b96855a355"></a>

## timeouts — timeouts / a5aaf1f1697a / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- timeouts

<a id="canonical-f857c91e432396dea8fed5264c59a430a40d0cbf59d19768481f4e715f722156"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2f956aec3800d1fdb2f04df9adac0af911c23c832fc8e44614a7b6ea10f5ad7d"></a>

## Direct properties — timeouts / a5aaf1f1697a / 3

<a id="canonical-56a1528ea6da0a79b116e34c153fa51770a770de2bfc122a928d1786cb8e0f38"></a>

<a id="canonical-10e39bff16aa441bc85aff1fb0c8c0745056eb4d7f058abcc36ebe8bd3b07ac0"></a>

## create property — timeouts / a5aaf1f1697a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a7b4394492ac8768f4bb485ae78405ef9f20198f15d03eec504e34fe2e4a64df"></a>

<a id="canonical-95f92e2a5e7067ef809aaa9e5b58e9843b5c42199c3ea0e6432035e6ae92803c"></a>

## delete property — timeouts / a5aaf1f1697a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-9a0571df4547a71b73c1c9dc17eb4d2a0c54f9d4aeccc30cfe13069ddc375511"></a>

<a id="canonical-2e80bb414208235c0e397e5cfc87d8f0bc88a44a0c7ce6d75348e0f857fca6dd"></a>

## read property — timeouts / a5aaf1f1697a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-689b7e964a39f980552acc9f6978421cbbc46219abe269ef9ffac5a5d4a805c1"></a>

<a id="canonical-fb858df1bbd4684a772130ba42e22d2add362f18e72a6d3ffac67416326d2c08"></a>

## update property — timeouts / a5aaf1f1697a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-caa283b47fc7740a4d16d8597b1ffd38bfa789030e1867b24d72f7dca4e01a6c"></a>

## Next pages — timeouts / a5aaf1f1697a / 8

- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
