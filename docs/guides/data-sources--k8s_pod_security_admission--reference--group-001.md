---
page_title: "xcsh_k8s_pod_security_admission reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission reference."
---

# xcsh_k8s_pod_security_admission reference

<a id="canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4baabc1c968436621fc59a8fc3c5bb97cca6ddf5a380c7b15dedfd8f50879d03"></a>

## Property reference — Property reference / 3917a2950cca / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- Property reference

<a id="canonical-33fb07ba142bcfe85e00eff6ac54baaa8a23307cc137c2b49cef6ea80a51b749"></a>

## Direct properties — Property reference / 3917a2950cca / 3

<a id="canonical-2b2a331540ffc6d264262c015da7b8c0393caaaa2f77c1f79faba6409067cfa3"></a>

<a id="canonical-bac803c803a3e5456a6939da8e902f509f7d3ffdb2a948c5ce82f63bd6213bf0"></a>

## annotations property — Property reference / 3917a2950cca / 4

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

<a id="canonical-eb93082416ffa1396618a0af6f4b6f71681de600bef4ae7ef1663231cfda81f8"></a>

<a id="canonical-54d4f8695ff7ac72aaa523ee3a37cc89d408ee4d575147fe499f0286eac27a2d"></a>

## description property — Property reference / 3917a2950cca / 5

Type: `"string"`. Computed.

Description of the K8SPodSecurityAdmission.

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

<a id="canonical-0d524fb33dafcead1161f8d6119cf203503ffaa2c7084486a207c8eb69f9a58f"></a>

<a id="canonical-be23cb8461bcd41aabcb467d99b7f8d305cbaefd828c82e447d5687292aca4fb"></a>

## id property — Property reference / 3917a2950cca / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ae695f9f0f2fafdbb0299fca562f9aae5f2a2b80b110ac32b8f293349bfce4f4"></a>

<a id="canonical-27d907d8b06b8069178aaf7ce3382b2666c905a2ef361c1ecb21cc8c658d4f8e"></a>

## labels property — Property reference / 3917a2950cca / 7

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

<a id="canonical-8d68fc3a98db4d77341c334f3f680d5132b4bfbf7d867191e963e4e89c929493"></a>

<a id="canonical-4562677493a65ef66c0d0520979bcc01a6d583f1c8fce870406da0656eeed87b"></a>

## name property — Property reference / 3917a2950cca / 8

Type: `"string"`. Required.

Name of the K8SPodSecurityAdmission.

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

<a id="canonical-b3f3680f644776016a879603276af53072662f4516fb563b876922111357594f"></a>

<a id="canonical-8b060121eaf2aefbdba386beb3bc87c027d3cccfb25467c45d405a71e581e8aa"></a>

## namespace property — Property reference / 3917a2950cca / 9

Type: `"string"`. Optional, Computed.

Namespace where the K8SPodSecurityAdmission exists.

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

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d): complete subsection reference.

<a id="canonical-434f418c305e40d22060f063b839d583c6f3d19791bad55bd161f127c3021bf0"></a>

## All schema paths — Property reference / 3917a2950cca / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2b2a331540ffc6d264262c015da7b8c0393caaaa2f77c1f79faba6409067cfa3) |
| `description` | [description](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-eb93082416ffa1396618a0af6f4b6f71681de600bef4ae7ef1663231cfda81f8) |
| `id` | [id](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0d524fb33dafcead1161f8d6119cf203503ffaa2c7084486a207c8eb69f9a58f) |
| `labels` | [labels](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-ae695f9f0f2fafdbb0299fca562f9aae5f2a2b80b110ac32b8f293349bfce4f4) |
| `name` | [name](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-8d68fc3a98db4d77341c334f3f680d5132b4bfbf7d867191e963e4e89c929493) |
| `namespace` | [namespace](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-b3f3680f644776016a879603276af53072662f4516fb563b876922111357594f) |
| `pod_security_admission_specs` | [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-250815a8de957ce889f52baa7b9dcbe72a66f16023569384cb8c692d91c01210) |
| `pod_security_admission_specs.audit` | [pod_security_admission_specs.audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-bf1a12557b30ebd4bb6672371f8d661de042ebab7697dfe5b94b8d53612597f7) |
| `pod_security_admission_specs.baseline` | [pod_security_admission_specs.baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-eb47b9c09057dd9bd1421060ab7cde32dae5e8a3fa89d1ccd31734a48d16dd62) |
| `pod_security_admission_specs.enforce` | [pod_security_admission_specs.enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-063d9b3427536c5507af5bc74c46666117b82f584d8947eacde4442a3025e99f) |
| `pod_security_admission_specs.privileged` | [pod_security_admission_specs.privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-c07fb0be8b4a2d4e03ea4ae1b8e5905d7f057d31e5a04a15851b1752f34652e3) |
| `pod_security_admission_specs.restricted` | [pod_security_admission_specs.restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0de418e3c8f4f16e91d449e4281260ae2c18eb1ec0146be31c207c2c2778b90b) |
| `pod_security_admission_specs.warn` | [pod_security_admission_specs.warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3ddad708ac6485830b46fca96c9c1e8baf1719a247355421df20cb0ec73fb406) |

<a id="canonical-1d4e3dfc11d8f4a491e6380729a679b374f2989dbd1a98df31bfdb3432c2d0ff"></a>

## Next pages — Property reference / 3917a2950cca / 11

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-084260a3488950e97bec19f17f44fac02ee73f8aacbe2c32f15226a6159465fd"></a>

## pod_security_admission_specs — pod_security_admission_specs / 78ae448f9a23 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- pod_security_admission_specs

<a id="canonical-250815a8de957ce889f52baa7b9dcbe72a66f16023569384cb8c692d91c01210"></a>

Type: `"list"`. Computed.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3e18d29e6d4580d8132a2052a46a72f1031cece7c55cf93176e04c0e07198d5b"></a>

## Direct properties — pod_security_admission_specs / 78ae448f9a23 / 3

- [audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-d4c1328337fd82c0b5acdab818b68d83131b3c75f55af748138e50a473886b35): complete subsection reference.

- [baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-900ee9dc88027d6aecd70aa484ed3f6b40835883e82d66899959d0d9c53bffc1): complete subsection reference.

- [enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-6935b4a90582c7baf618aba4023cd777f29d65a755f24f8346296f6f87068357): complete subsection reference.

- [privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-b25f495c80a14ed5fa7b1b9046b4ae45c8ba89a949f1715ce9863527256348a4): complete subsection reference.

- [restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-49feda6a05bee11b5a41a05d8042cad1615ea7bab5aeb308a2a816fa7e37204d): complete subsection reference.

- [warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-72c6d4e282506ee8b866fcf6556c0eed1375fad843457c675b5aeaa1459ca52f): complete subsection reference.

<a id="canonical-187b9ebb0dcc4a1aa8b6100e72fc8137375c7a6d3d3ea781da38a4cd389f3084"></a>

## Next pages — pod_security_admission_specs / 78ae448f9a23 / 4

- [pod_security_admission_specs.audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-d4c1328337fd82c0b5acdab818b68d83131b3c75f55af748138e50a473886b35)
- [pod_security_admission_specs.baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-900ee9dc88027d6aecd70aa484ed3f6b40835883e82d66899959d0d9c53bffc1)
- [pod_security_admission_specs.enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-6935b4a90582c7baf618aba4023cd777f29d65a755f24f8346296f6f87068357)
- [pod_security_admission_specs.privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-b25f495c80a14ed5fa7b1b9046b4ae45c8ba89a949f1715ce9863527256348a4)
- [pod_security_admission_specs.restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-49feda6a05bee11b5a41a05d8042cad1615ea7bab5aeb308a2a816fa7e37204d)
- [pod_security_admission_specs.warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-72c6d4e282506ee8b866fcf6556c0eed1375fad843457c675b5aeaa1459ca52f)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-d4c1328337fd82c0b5acdab818b68d83131b3c75f55af748138e50a473886b35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd9292b50c9baf873954387f9165b9e657a2041122db7e4658aaa321869bcd21"></a>

## pod_security_admission_specs.audit — pod_security_admission_specs.audit / 603eb7da275a / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.audit

<a id="canonical-bf1a12557b30ebd4bb6672371f8d661de042ebab7697dfe5b94b8d53612597f7"></a>

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

<a id="canonical-29dcd6e11c7f727a58601f7be4678a34f985f8046a95f769ee521ac2c5380f78"></a>

## Direct properties — pod_security_admission_specs.audit / 603eb7da275a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fac15f16520ef2e893dd950d8da0ac6c9b997a373b7c57926be4871a2281dca"></a>

## Next pages — pod_security_admission_specs.audit / 603eb7da275a / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-900ee9dc88027d6aecd70aa484ed3f6b40835883e82d66899959d0d9c53bffc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a68dfc7ec96ca13a6ef53faad503445250c0daae9fd54dd7858075879716070"></a>

## pod_security_admission_specs.baseline — pod_security_admission_specs.baseline / f28d60320619 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.baseline

<a id="canonical-eb47b9c09057dd9bd1421060ab7cde32dae5e8a3fa89d1ccd31734a48d16dd62"></a>

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

<a id="canonical-2aec99d9ddb89b5587c4f9972234455a2943cdadcaf0f5188ab8ceb48439d16e"></a>

## Direct properties — pod_security_admission_specs.baseline / f28d60320619 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e8a244730ce2dff04d797d1c6315aaf186c1abb2c4a4349a559b9c6f1b6ed15"></a>

## Next pages — pod_security_admission_specs.baseline / f28d60320619 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-6935b4a90582c7baf618aba4023cd777f29d65a755f24f8346296f6f87068357"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55832c5698f6e84676b3840d7e71f8121df7d9608f518f4f0811ef9dc67ecbfe"></a>

## pod_security_admission_specs.enforce — pod_security_admission_specs.enforce / 14bc9a887ad3 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.enforce

<a id="canonical-063d9b3427536c5507af5bc74c46666117b82f584d8947eacde4442a3025e99f"></a>

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

<a id="canonical-50a8501be8a3c4c2840eb36ac484aa4617da5b846f737cadb1a428371314b679"></a>

## Direct properties — pod_security_admission_specs.enforce / 14bc9a887ad3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18d4176310bac41f993d40ab506a6f24cbd0712cfc81999ca201715550aff69e"></a>

## Next pages — pod_security_admission_specs.enforce / 14bc9a887ad3 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-b25f495c80a14ed5fa7b1b9046b4ae45c8ba89a949f1715ce9863527256348a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5880eb0484571482bacb48523dd49334a5bd5e3d27e8aa915f4be9300125b40"></a>

## pod_security_admission_specs.privileged — pod_security_admission_specs.privileged / a9bb9ec17b88 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.privileged

<a id="canonical-c07fb0be8b4a2d4e03ea4ae1b8e5905d7f057d31e5a04a15851b1752f34652e3"></a>

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

<a id="canonical-3492786a648b7de218001124473bcda950562c9f3b957c1b18ed0238bedb12a5"></a>

## Direct properties — pod_security_admission_specs.privileged / a9bb9ec17b88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a10c9045553c74744c07fddcd0e6e6fb82e3dd49e4acd256bb4e3f2265709095"></a>

## Next pages — pod_security_admission_specs.privileged / a9bb9ec17b88 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-49feda6a05bee11b5a41a05d8042cad1615ea7bab5aeb308a2a816fa7e37204d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4702e6d9b08b70ac0c08e9b65488ad8e3824ca35667d2e99705386140037927"></a>

## pod_security_admission_specs.restricted — pod_security_admission_specs.restricted / 45704e373db5 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.restricted

<a id="canonical-0de418e3c8f4f16e91d449e4281260ae2c18eb1ec0146be31c207c2c2778b90b"></a>

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

<a id="canonical-82fbb6a176c3012ca1471bd00509bc037f509cd8c5f06081497a6cbd6779e148"></a>

## Direct properties — pod_security_admission_specs.restricted / 45704e373db5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06d19dc3b8f9405b9f891c80817baa39ccc5e2b8716e01037b243769515fa091"></a>

## Next pages — pod_security_admission_specs.restricted / 45704e373db5 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-72c6d4e282506ee8b866fcf6556c0eed1375fad843457c675b5aeaa1459ca52f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43bd4e5d8217dc8ceddedf03f39bf91f2b77ff69f982be2a73196933f97fcaed"></a>

## pod_security_admission_specs.warn — pod_security_admission_specs.warn / 4a693036d2fd / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- pod_security_admission_specs.warn

<a id="canonical-3ddad708ac6485830b46fca96c9c1e8baf1719a247355421df20cb0ec73fb406"></a>

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

<a id="canonical-c748d8b27927583c01e4223b60329ccb6a2d69fedcbb63d103887602eea246a8"></a>

## Direct properties — pod_security_admission_specs.warn / 4a693036d2fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5cc4dad11bf74e0bf95908dfdf9cd3812e19e1ff5951b2a0c8898e352fe949d3"></a>

## Next pages — pod_security_admission_specs.warn / 4a693036d2fd / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-655890b9a9c3b2fbe04bd5a1072f68134d1060b7ed9bbc6c4523ba7dd8fbf33d)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
