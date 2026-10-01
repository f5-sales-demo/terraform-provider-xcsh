---
page_title: "xcsh_k8s_pod_security_admission reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission reference."
---

# xcsh_k8s_pod_security_admission reference

<a id="canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb168c5f7989a15414b31df0a92d8faa634085082a54b1cc0306586769141ec3"></a>

## Property reference — Property reference / 31ec89ab334b / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- Property reference

<a id="canonical-d666aa0ece641de8a3519099ecd9f305a3d13b21e155cf019bf38eae67b111f2"></a>

## Direct properties — Property reference / 31ec89ab334b / 3

<a id="canonical-a6f5d30740eaaf7ae504e1112cf382e8554fa77558c9e1935a7af65afbac2f54"></a>

<a id="canonical-6cd1bf916c5db60d6ba43e8ea7b30f27f85ccb68773d12ea22434bc0e49ead5c"></a>

## annotations property — Property reference / 31ec89ab334b / 4

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

<a id="canonical-aaf37cc584341118c194780802ba56fc72ec803c97dda6b15194f8295d3794a8"></a>

<a id="canonical-64daf470a3466376ad56fcb1b3f37b18d481b7d331095fbf894600b0f5675729"></a>

## description property — Property reference / 31ec89ab334b / 5

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

<a id="canonical-8be60c5c8fa91dba29bfc08a298678a577eeee81ea6be2d12594d1e95f1cd216"></a>

<a id="canonical-c06a60621be2d36b90ac16928a57bffdb0602060ed1f0c59663f16422d9562af"></a>

## disable property — Property reference / 31ec89ab334b / 6

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

<a id="canonical-75739b8e8c3a686e16fddb7eca9a1a0e9ffbbe8ec020d37bbc5d8b0f35311203"></a>

<a id="canonical-267a2d5310c51de368d16a0bed4b291c0fe1e3db8bea54ffb3f01108a4bd2f89"></a>

## id property — Property reference / 31ec89ab334b / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-af8d9406dea0ece6837dfb8291f047fbf4039c6d7194fd0eaeeec6d1b3a993d1"></a>

<a id="canonical-07803b6835b608c5876d6617a462cb9ec2bd9e049afbc9d2d04ebda6b624e791"></a>

## labels property — Property reference / 31ec89ab334b / 8

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

<a id="canonical-7ca32020046484775e8967d4b705e51aaba8532c20a7cec83d944f06293ad3bf"></a>

<a id="canonical-838d295e6acb73e6c265f9e8c7e3cc56f771680844114aae24e1995a69c6db7e"></a>

## name property — Property reference / 31ec89ab334b / 9

Type: `"string"`. Required.

Name of the K8S Pod Security Admission. Must be unique within the namespace.

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

<a id="canonical-605d3fe33c36641f2b1d7c90f57ec3ff65843f8878ebfafca315106a32268790"></a>

<a id="canonical-d29d66e8e828a8532b8bc1a5cdc6af76d3f26d2a2119980b53698aa0be6e5c39"></a>

## namespace property — Property reference / 31ec89ab334b / 10

Type: `"string"`. Optional, Computed.

Namespace for the K8S Pod Security Admission. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e): complete subsection reference.

- [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-f7aa1facb841be5bf233bc09b9ec7d04217b43ce010a0a065bd1361d5cd646a8): complete subsection reference.

<a id="canonical-853c12911fc281eb68736a74908bd506f6bbcfe63cddc29f4e30b388cb81ccad"></a>

## All schema paths — Property reference / 31ec89ab334b / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_pod_security_admission--reference--group-001.md#canonical-a6f5d30740eaaf7ae504e1112cf382e8554fa77558c9e1935a7af65afbac2f54) |
| `description` | [description](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaf37cc584341118c194780802ba56fc72ec803c97dda6b15194f8295d3794a8) |
| `disable` | [disable](resources--k8s_pod_security_admission--reference--group-001.md#canonical-8be60c5c8fa91dba29bfc08a298678a577eeee81ea6be2d12594d1e95f1cd216) |
| `id` | [id](resources--k8s_pod_security_admission--reference--group-001.md#canonical-75739b8e8c3a686e16fddb7eca9a1a0e9ffbbe8ec020d37bbc5d8b0f35311203) |
| `labels` | [labels](resources--k8s_pod_security_admission--reference--group-001.md#canonical-af8d9406dea0ece6837dfb8291f047fbf4039c6d7194fd0eaeeec6d1b3a993d1) |
| `name` | [name](resources--k8s_pod_security_admission--reference--group-001.md#canonical-7ca32020046484775e8967d4b705e51aaba8532c20a7cec83d944f06293ad3bf) |
| `namespace` | [namespace](resources--k8s_pod_security_admission--reference--group-001.md#canonical-605d3fe33c36641f2b1d7c90f57ec3ff65843f8878ebfafca315106a32268790) |
| `pod_security_admission_specs` | [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-c9e5711f84cbee9e46ecd1d3857f9504b2fd2b60a207813439ad22aa2cddd9ed) |
| `pod_security_admission_specs.audit` | [pod_security_admission_specs.audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1cca172d04f63c669192355148d79c9b23408ccb78295c7d04f6172138c26ead) |
| `pod_security_admission_specs.baseline` | [pod_security_admission_specs.baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-14cb9dd55507fe120c8c09ca701025ba5a80ddca9fa86abc8dbf6cf5f2409eab) |
| `pod_security_admission_specs.enforce` | [pod_security_admission_specs.enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-db94b717e4b471f9f0866722f9c4ec52be47248f62ae99baf8c0ad155221c8e4) |
| `pod_security_admission_specs.privileged` | [pod_security_admission_specs.privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-cc1c91e265bc0cc041ad39d8c6ca952bae1eb7567cbe699181a6891fdb89a8c9) |
| `pod_security_admission_specs.restricted` | [pod_security_admission_specs.restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-98dbcac3306403677e1d16fca1c43a01c0c1a580067c60f121f1f597b3d905ef) |
| `pod_security_admission_specs.warn` | [pod_security_admission_specs.warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-768d9880d6fb8a8948a0ec06d868acd3c00c86aeef2fd116b2826f4d2ce6fc0f) |
| `timeouts` | [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-5c760c33de749493ff348648e1893fe87bec52445660657ceb8d1d0efe9c4560) |
| `timeouts.create` | [timeouts.create](resources--k8s_pod_security_admission--reference--group-001.md#canonical-c1f55c977f33a88b2b3a96dd8db05100f75d2c3a81f2d20b3a7d843bf3b636cb) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_pod_security_admission--reference--group-001.md#canonical-7388d9afaa715883366b7cd637d3cdd256d7cd0cfe9a9d89894086924d0b418e) |
| `timeouts.read` | [timeouts.read](resources--k8s_pod_security_admission--reference--group-001.md#canonical-21e2f797913af7d5dcac524551c0ad6f5692c25b4387ba5a13d86413c3963efd) |
| `timeouts.update` | [timeouts.update](resources--k8s_pod_security_admission--reference--group-001.md#canonical-533232fbdec79c51a19629a3c9b8985da9018aa6b4683545886c408d18dde969) |

<a id="canonical-65967d8b157e68cfc5f2f18632773b35ec69b0849e81a19cd683eac181b5c088"></a>

## Next pages — Property reference / 31ec89ab334b / 12

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-f7aa1facb841be5bf233bc09b9ec7d04217b43ce010a0a065bd1361d5cd646a8)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28493261e863f5fd80dc29e2ccd93bedcd9524175cb3f9e67ab3cb89c0ee1e0c"></a>

## pod_security_admission_specs — pod_security_admission_specs / f365ecfc12ff / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- pod_security_admission_specs

<a id="canonical-c9e5711f84cbee9e46ecd1d3857f9504b2fd2b60a207813439ad22aa2cddd9ed"></a>

Type: `"object"`. list nested block, Optional.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("audit",
    "enforce"),
  validators.ConflictingListObjectAttributes("audit",
    "warn"),
  validators.ConflictingListObjectAttributes("baseline",
    "privileged"),
  validators.ConflictingListObjectAttributes("baseline",
    "restricted"),
  validators.ConflictingListObjectAttributes("enforce",
    "warn"),
  validators.ConflictingListObjectAttributes("privileged",
    "restricted")}
```

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

Terraform syntax:

```terraform
pod_security_admission_specs {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c38de8990e80a91d05ced9d9b41c9b8e61036f6c88e86e31aa067f81fff54bd"></a>

## Direct properties — pod_security_admission_specs / f365ecfc12ff / 3

- [audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1591724e5ece9b7a1d5ef011347e16e1f3c5a91bf67cccf8ba0434ee749f9d3d): complete subsection reference.

- [baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-9f611a7c471cdfa615249f261a58677e836ea905f28e4d99027601c45988d53e): complete subsection reference.

- [enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-34feeddc884020fed68d84a3b3ed2a54582e51c04325791d242a3be0ff0c43e6): complete subsection reference.

- [privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-a984f5d8d2a4c12dc7665d4ff47e73aebf1fa1d8504c92f6aac345de40e9edff): complete subsection reference.

- [restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-5244bbbbf078727c3be3855de58d149690fbe7ffaeab51a39664641e84b8a55f): complete subsection reference.

- [warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-bbb378dd83284cbcfb899cc2767cebc970b9b60f3ad36e75bd04c3dbe9cf8301): complete subsection reference.

<a id="canonical-95b0d7f067c6ad1fd9f4def1395b0ea8fa64977c0ab62dd10fbc566ea47de3a6"></a>

## Next pages — pod_security_admission_specs / f365ecfc12ff / 4

- [pod_security_admission_specs.audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1591724e5ece9b7a1d5ef011347e16e1f3c5a91bf67cccf8ba0434ee749f9d3d)
- [pod_security_admission_specs.baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-9f611a7c471cdfa615249f261a58677e836ea905f28e4d99027601c45988d53e)
- [pod_security_admission_specs.enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-34feeddc884020fed68d84a3b3ed2a54582e51c04325791d242a3be0ff0c43e6)
- [pod_security_admission_specs.privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-a984f5d8d2a4c12dc7665d4ff47e73aebf1fa1d8504c92f6aac345de40e9edff)
- [pod_security_admission_specs.restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-5244bbbbf078727c3be3855de58d149690fbe7ffaeab51a39664641e84b8a55f)
- [pod_security_admission_specs.warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-bbb378dd83284cbcfb899cc2767cebc970b9b60f3ad36e75bd04c3dbe9cf8301)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-1591724e5ece9b7a1d5ef011347e16e1f3c5a91bf67cccf8ba0434ee749f9d3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fb46d7eb7a309a215166d650c76d0df327b8e769c6a78a7804fef7c9b214a68"></a>

## pod_security_admission_specs.audit — pod_security_admission_specs.audit / a627796e3d27 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.audit

<a id="canonical-1cca172d04f63c669192355148d79c9b23408ccb78295c7d04f6172138c26ead"></a>

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
audit = {}
```

<a id="canonical-9a1bd8523c9caec863ca2272ffe45e2a6333b32bb1bdd74bce61c58ea733f80c"></a>

## Direct properties — pod_security_admission_specs.audit / a627796e3d27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffd5b5c35c64734f55fdfa49584076f84f6b9d12e25371b4e400905fb9cad07d"></a>

## Next pages — pod_security_admission_specs.audit / a627796e3d27 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-9f611a7c471cdfa615249f261a58677e836ea905f28e4d99027601c45988d53e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-814d34415215bbb3715bedb1974bb302e39d925e70d21de4e102a7e984527f81"></a>

## pod_security_admission_specs.baseline — pod_security_admission_specs.baseline / 096bd22a3404 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.baseline

<a id="canonical-14cb9dd55507fe120c8c09ca701025ba5a80ddca9fa86abc8dbf6cf5f2409eab"></a>

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
baseline = {}
```

<a id="canonical-153aaa308a2c4554e42b367c0ebc6a890fcf669f61a373dd71a63959e54aa8be"></a>

## Direct properties — pod_security_admission_specs.baseline / 096bd22a3404 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db67282aa1488a4f1c92b59c22db54dc638c902632839ae29bbb2ac36197cb4c"></a>

## Next pages — pod_security_admission_specs.baseline / 096bd22a3404 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-34feeddc884020fed68d84a3b3ed2a54582e51c04325791d242a3be0ff0c43e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b04e8f15a9d12010d4bf560dc73bb42c9fe2dae94ecd42eb21d51207411ffe65"></a>

## pod_security_admission_specs.enforce — pod_security_admission_specs.enforce / 4962684235c7 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.enforce

<a id="canonical-db94b717e4b471f9f0866722f9c4ec52be47248f62ae99baf8c0ad155221c8e4"></a>

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
enforce = {}
```

<a id="canonical-f7e46b2a1277d22a048d8a4c29c5be52f940ec804a69564f25e4bd5d96df438e"></a>

## Direct properties — pod_security_admission_specs.enforce / 4962684235c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5708eb597706e16b987be1dc0930be727c57160135b1a9ce3899213b820fe453"></a>

## Next pages — pod_security_admission_specs.enforce / 4962684235c7 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-a984f5d8d2a4c12dc7665d4ff47e73aebf1fa1d8504c92f6aac345de40e9edff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c7bbd8787523063d208d6bfd3e8fbd6b913f2b5c3ee079f2f03b0dbc121e456"></a>

## pod_security_admission_specs.privileged — pod_security_admission_specs.privileged / d7b79acb9fc5 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.privileged

<a id="canonical-cc1c91e265bc0cc041ad39d8c6ca952bae1eb7567cbe699181a6891fdb89a8c9"></a>

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
privileged = {}
```

<a id="canonical-9feb0eff9a40791d4e2ca160204c95d7a47b35fd1db30f059a07b5eb91afebee"></a>

## Direct properties — pod_security_admission_specs.privileged / d7b79acb9fc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b400354f4547138ea9fca4fcaa8c8f13d30de92f79598a48808b8cb920b6d17"></a>

## Next pages — pod_security_admission_specs.privileged / d7b79acb9fc5 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-5244bbbbf078727c3be3855de58d149690fbe7ffaeab51a39664641e84b8a55f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15a8e5df38d4a47ac0ef0f3eb10c590fac037694157d53ba90cf93284229a105"></a>

## pod_security_admission_specs.restricted — pod_security_admission_specs.restricted / ff4c4aa1cfc4 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.restricted

<a id="canonical-98dbcac3306403677e1d16fca1c43a01c0c1a580067c60f121f1f597b3d905ef"></a>

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
restricted = {}
```

<a id="canonical-220bdd5c45b23c868eee6f83d7717bd5e975c780457cc7e60b739494ae154c9d"></a>

## Direct properties — pod_security_admission_specs.restricted / ff4c4aa1cfc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2682169e8d363d95dcc26a583186b929e2f3ec80e93f66812c2aa62074bf1c7"></a>

## Next pages — pod_security_admission_specs.restricted / ff4c4aa1cfc4 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-bbb378dd83284cbcfb899cc2767cebc970b9b60f3ad36e75bd04c3dbe9cf8301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a09b0752be3f046cf0b86955a0f383986ad42a92be4282e8b105a1e872545a7"></a>

## pod_security_admission_specs.warn — pod_security_admission_specs.warn / 027dce8b8c15 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- pod_security_admission_specs.warn

<a id="canonical-768d9880d6fb8a8948a0ec06d868acd3c00c86aeef2fd116b2826f4d2ce6fc0f"></a>

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
warn = {}
```

<a id="canonical-781c7cc456e4d2c1883fd3ddb118fd43affd044f25e4cbb4e78a624ab7da945a"></a>

## Direct properties — pod_security_admission_specs.warn / 027dce8b8c15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4b2f2e9567ce7d3ee412d4c9d1e44ea0b0013f3470090b1b9b04afa8041745e"></a>

## Next pages — pod_security_admission_specs.warn / 027dce8b8c15 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-4032c5a06df71454739961020a3b3c1f341c5524462000e18737356083a85a0e)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)

<a id="canonical-f7aa1facb841be5bf233bc09b9ec7d04217b43ce010a0a065bd1361d5cd646a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ad0f36104d8034fb41b0f628a4b8db898a72d679926548b3100b2581f246ccc"></a>

## timeouts — timeouts / 0a700297d5c2 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- timeouts

<a id="canonical-5c760c33de749493ff348648e1893fe87bec52445660657ceb8d1d0efe9c4560"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-93f6860831d572bcfcfc692795ef08c318b8e94f5deeedbfac034320da0d3e12"></a>

## Direct properties — timeouts / 0a700297d5c2 / 3

<a id="canonical-c1f55c977f33a88b2b3a96dd8db05100f75d2c3a81f2d20b3a7d843bf3b636cb"></a>

<a id="canonical-c69543a3f4c2044a8d51abd06009960ffaf99db4476b8a2b5006097ae621c442"></a>

## create property — timeouts / 0a700297d5c2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7388d9afaa715883366b7cd637d3cdd256d7cd0cfe9a9d89894086924d0b418e"></a>

<a id="canonical-c48e7aa233f5de00d409774b449999972b3a41ffffb2dc0138bf4d36fdf033cf"></a>

## delete property — timeouts / 0a700297d5c2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-21e2f797913af7d5dcac524551c0ad6f5692c25b4387ba5a13d86413c3963efd"></a>

<a id="canonical-3ce3c8c9bb09d6916291d5eb42b9317e2cda6ee20d9b59a6c6fabacb0bea1e08"></a>

## read property — timeouts / 0a700297d5c2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-533232fbdec79c51a19629a3c9b8985da9018aa6b4683545886c408d18dde969"></a>

<a id="canonical-370ab0bd8e019c5a73728104c3997b34e5daa23e4cfb0436a2b03735bac8c265"></a>

## update property — timeouts / 0a700297d5c2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c8b3bb41d80e16eb6f761ec09f5f8bdf6b2992d37651a3c4652f88bbce1f73d6"></a>

## Next pages — timeouts / 0a700297d5c2 / 8

- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-aaff46c769c7779ed88f3eb28bd97bdfe97115cc0982f7d385c2509e2607313c)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-37de369e9fe19aaf932b239f810d73047dfafd2b1a01d6804c3b7eb4d46f1b8b)
