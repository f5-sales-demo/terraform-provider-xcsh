---
page_title: "xcsh_virtual_k8s reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s reference."
---

# xcsh_virtual_k8s reference

<a id="canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf9d178d417c50c1d7064853e8e68509d0145f3923983aa52648fb5c04318434"></a>

## Property reference — Property reference / 909d57b045b4 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- Property reference

<a id="canonical-893ea6c27ed4b06cb64cccf535802af2016c2e8d8c9b4dbb2f612d219b6f25a2"></a>

## Direct properties — Property reference / 909d57b045b4 / 3

<a id="canonical-ebb3569b05ec271ef839bb1e28f41ce4b985da2843cd5411f4f6c17bb4994215"></a>

<a id="canonical-37ca1203b8bd218d0d4c00cc4661fc3245fee4a287d2b5a3d49e45bc7f40019c"></a>

## annotations property — Property reference / 909d57b045b4 / 4

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

- [default_flavor_ref](resources--virtual_k8s--reference--group-001.md#canonical-0f54a9b86bfa22b142e78745108f94440ea9378e62a1fb7c41a8dc85d2d2d9e0): complete subsection reference.

<a id="canonical-cfe27972eda0ff4879b7bc7fbe04b41ea91fc0f6fee0fd78181e013c506cffe0"></a>

<a id="canonical-c816d513cb516d549a67dae195e9e6efbeac3fc9439abfa5cb8bad416020b1bb"></a>

## description property — Property reference / 909d57b045b4 / 5

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

<a id="canonical-22c965379409db23a824d128cd80d256f273007de6251fbd1b0139841d3efafe"></a>

<a id="canonical-81f8d00b953b93dae0aaf8b0642b7468e997794007c3d50f1166958a739a26be"></a>

## disable property — Property reference / 909d57b045b4 / 6

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

- [disabled](resources--virtual_k8s--reference--group-001.md#canonical-4052571cad83073706a8fb89f2d4138b8c2af77c88cd78371c390897b80e9fcf): complete subsection reference.

<a id="canonical-7f736006a6818c2241689c1eead27ac0e6b7f6b1b178f61aca208109e8a76fa6"></a>

<a id="canonical-8bb3e1918f30125d90b18c608162d2f9947f0957fb5b47c262ab1e45611021b0"></a>

## id property — Property reference / 909d57b045b4 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](resources--virtual_k8s--reference--group-001.md#canonical-2d46609cf2240b060a6e693c4f0aa7f12184c8e2f9d74abb3bf104f322f18a55): complete subsection reference.

<a id="canonical-72f9c27aa75f5eb900f5b1c948904a4167db626ffef5908516ebaa91897d3995"></a>

<a id="canonical-622efbff33e86d7f6292a5bead704a82c6db42c71a36a5711d1639e071706518"></a>

## labels property — Property reference / 909d57b045b4 / 8

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

<a id="canonical-b4cf2b6047dfb1a14d5e807038a0bea3e639c41caab751ffa933662804518983"></a>

<a id="canonical-b21a152b361656946e44b0d2e73c94a6ab58a9ac292c7d3ee993cb3445851e21"></a>

## name property — Property reference / 909d57b045b4 / 9

Type: `"string"`. Required.

Name of the Virtual K8S. Must be unique within the namespace.

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

<a id="canonical-9698d070b8d254dd721c79c5deb6efcdad9c2a29dd8434faf854ca562da845c8"></a>

<a id="canonical-4c8a821bc0530a6abeeda4aea88a0c3838aac14aec6abf52949071de005b6434"></a>

## namespace property — Property reference / 909d57b045b4 / 10

Type: `"string"`. Required.

Namespace where the Virtual K8S is created.

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

- [timeouts](resources--virtual_k8s--reference--group-001.md#canonical-7cbca427f616f67b6467a1c59473022dc1de8bb3efc25a376920639794ceea2d): complete subsection reference.

- [vsite_refs](resources--virtual_k8s--reference--group-001.md#canonical-aa124a377812667d48addd7e5cc3512af59946547c02528b29fddd94c4ab218d): complete subsection reference.

<a id="canonical-37dfffb6fdb0fd717767e898357b023b651fac697070d8365c7c3a2fc8560104"></a>

## All schema paths — Property reference / 909d57b045b4 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_k8s--reference--group-001.md#canonical-ebb3569b05ec271ef839bb1e28f41ce4b985da2843cd5411f4f6c17bb4994215) |
| `default_flavor_ref` | [default_flavor_ref](resources--virtual_k8s--reference--group-001.md#canonical-93f885a218075339bffc3ff07ef9f712bfd8ea99f950edec50ad72c6c402872c) |
| `default_flavor_ref.name` | [default_flavor_ref.name](resources--virtual_k8s--reference--group-001.md#canonical-d93475cc81963b028745e54e01a611abadb9f4f8e1c456ce2adcc395374c5365) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](resources--virtual_k8s--reference--group-001.md#canonical-b30488adc2736385402cd8b13c7ca0b0c44ec127011db7ca2e3e13c0f502eacc) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](resources--virtual_k8s--reference--group-001.md#canonical-2a8f6dd4bafa3d6944b9e99b8fc87005e956c315d3141dedbc0f393532e3d78c) |
| `description` | [description](resources--virtual_k8s--reference--group-001.md#canonical-cfe27972eda0ff4879b7bc7fbe04b41ea91fc0f6fee0fd78181e013c506cffe0) |
| `disable` | [disable](resources--virtual_k8s--reference--group-001.md#canonical-22c965379409db23a824d128cd80d256f273007de6251fbd1b0139841d3efafe) |
| `disabled` | [disabled](resources--virtual_k8s--reference--group-001.md#canonical-efbb5e2fe66a1a5523b753f26205e8008885f6420377c2d1986b645b9f87969b) |
| `id` | [id](resources--virtual_k8s--reference--group-001.md#canonical-7f736006a6818c2241689c1eead27ac0e6b7f6b1b178f61aca208109e8a76fa6) |
| `isolated` | [isolated](resources--virtual_k8s--reference--group-001.md#canonical-8359ab4237d5f38583f0fa323a7c716b06ef98295cb085826c17d0ca71b8dd8e) |
| `labels` | [labels](resources--virtual_k8s--reference--group-001.md#canonical-72f9c27aa75f5eb900f5b1c948904a4167db626ffef5908516ebaa91897d3995) |
| `name` | [name](resources--virtual_k8s--reference--group-001.md#canonical-b4cf2b6047dfb1a14d5e807038a0bea3e639c41caab751ffa933662804518983) |
| `namespace` | [namespace](resources--virtual_k8s--reference--group-001.md#canonical-9698d070b8d254dd721c79c5deb6efcdad9c2a29dd8434faf854ca562da845c8) |
| `timeouts` | [timeouts](resources--virtual_k8s--reference--group-001.md#canonical-09a6c6fdb4c4765341ea198c27fc40f100db2c08099e52713f8a0b530f1e9fa6) |
| `timeouts.create` | [timeouts.create](resources--virtual_k8s--reference--group-001.md#canonical-7bc003b8d9aee82135f1b5750d2bc968c6dd3a4e0d82f77f8254f544b50dce3a) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_k8s--reference--group-001.md#canonical-aaf0aae31d8fc2a986aec629ec1bb3ab6f017d6c34fb44513efbdb70246adb74) |
| `timeouts.read` | [timeouts.read](resources--virtual_k8s--reference--group-001.md#canonical-6fea8b3345ab414b1fc32589a70e5df4410be019774a92be61fcd158de051fab) |
| `timeouts.update` | [timeouts.update](resources--virtual_k8s--reference--group-001.md#canonical-a61c07f56c82f19eb60d04a90dfc71b1486cf9767222f22b22e141649fa81ec4) |
| `vsite_refs` | [vsite_refs](resources--virtual_k8s--reference--group-001.md#canonical-bd24655f8c5bd0b00af6c60238f855bf35ee55ec67eb8d63999dfeae5e2c8647) |
| `vsite_refs.kind` | [vsite_refs.kind](resources--virtual_k8s--reference--group-001.md#canonical-5b9781a2f972ea6d2095036bd9aa3b0c7cd24e64982ba2b0ad35f8c7b560e851) |
| `vsite_refs.name` | [vsite_refs.name](resources--virtual_k8s--reference--group-001.md#canonical-d231028ec731c0a93979d7170b9ff4c96cfc4a9a44c1f09731c202c7638a05e2) |
| `vsite_refs.namespace` | [vsite_refs.namespace](resources--virtual_k8s--reference--group-001.md#canonical-f2f959b95fa3da3738561df27b8a4321cf06fede8bc48ed78da68ea86153573a) |
| `vsite_refs.tenant` | [vsite_refs.tenant](resources--virtual_k8s--reference--group-001.md#canonical-3af01b71e2152e39b683354a0af823df26a8b8f8772e3b0390ba1bdc690ee6f3) |
| `vsite_refs.uid` | [vsite_refs.uid](resources--virtual_k8s--reference--group-001.md#canonical-86b3d1c4752c92b7a5c7046a0a7eccbd0970b4edc1f97cf86cb35a7ec8e5da15) |

<a id="canonical-0ad9cbcfbb788eff691bbaaac40a03241b96e841acec6187f18afac313cedc4c"></a>

## Next pages — Property reference / 909d57b045b4 / 12

- [default_flavor_ref](resources--virtual_k8s--reference--group-001.md#canonical-0f54a9b86bfa22b142e78745108f94440ea9378e62a1fb7c41a8dc85d2d2d9e0)
- [disabled](resources--virtual_k8s--reference--group-001.md#canonical-4052571cad83073706a8fb89f2d4138b8c2af77c88cd78371c390897b80e9fcf)
- [isolated](resources--virtual_k8s--reference--group-001.md#canonical-2d46609cf2240b060a6e693c4f0aa7f12184c8e2f9d74abb3bf104f322f18a55)
- [timeouts](resources--virtual_k8s--reference--group-001.md#canonical-7cbca427f616f67b6467a1c59473022dc1de8bb3efc25a376920639794ceea2d)
- [vsite_refs](resources--virtual_k8s--reference--group-001.md#canonical-aa124a377812667d48addd7e5cc3512af59946547c02528b29fddd94c4ab218d)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-0f54a9b86bfa22b142e78745108f94440ea9378e62a1fb7c41a8dc85d2d2d9e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d35bea5663a0872fa10ce78f88c6f0f5575f278975720888bc9f7c034adde1ef"></a>

## default_flavor_ref — default_flavor_ref / 90a94a5e4283 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- default_flavor_ref

<a id="canonical-93f885a218075339bffc3ff07ef9f712bfd8ea99f950edec50ad72c6c402872c"></a>

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
default_flavor_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-b37ff6b6f9a8ec9e1171ede4c89fd64f06eda6139d2f1d303377fc77cb7a1f94"></a>

## Direct properties — default_flavor_ref / 90a94a5e4283 / 3

<a id="canonical-d93475cc81963b028745e54e01a611abadb9f4f8e1c456ce2adcc395374c5365"></a>

<a id="canonical-efb8f7afbd901c111bc1ed6086aad560cf1692777e303324685c0839e9efe945"></a>

## name property — default_flavor_ref / 90a94a5e4283 / 4

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

<a id="canonical-b30488adc2736385402cd8b13c7ca0b0c44ec127011db7ca2e3e13c0f502eacc"></a>

<a id="canonical-15c78b52eab7d484074f0164c488d9b391a36add0d6f5f69e4665bb97d2c0281"></a>

## namespace property — default_flavor_ref / 90a94a5e4283 / 5

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

<a id="canonical-2a8f6dd4bafa3d6944b9e99b8fc87005e956c315d3141dedbc0f393532e3d78c"></a>

<a id="canonical-4eb79e5d92d7d0347dabced36f6ac4542248daabec8d8ade35fcf97b5a13eac2"></a>

## tenant property — default_flavor_ref / 90a94a5e4283 / 6

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

<a id="canonical-47337533061bf9e03b56ca3fa63f455965b49d5f5cb90f80d5d9c37b6585875b"></a>

## Next pages — default_flavor_ref / 90a94a5e4283 / 7

- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-4052571cad83073706a8fb89f2d4138b8c2af77c88cd78371c390897b80e9fcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c906356eecf1c89eee690932faeb93cf737abb83380fa0d7648760541f000a59"></a>

## disabled — disabled / 665efe66dcc8 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- disabled

<a id="canonical-efbb5e2fe66a1a5523b753f26205e8008885f6420377c2d1986b645b9f87969b"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, isolated\] Enable this option

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

- [disabled](resources--virtual_k8s--reference--group-001.md#canonical-efbb5e2fe66a1a5523b753f26205e8008885f6420377c2d1986b645b9f87969b)
- [isolated](resources--virtual_k8s--reference--group-001.md#canonical-8359ab4237d5f38583f0fa323a7c716b06ef98295cb085826c17d0ca71b8dd8e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

<a id="canonical-e5f4560c4244802f8d8bf02036b6099c0a75a05b6dd927d5dc0a0d0c5c44e429"></a>

## Direct properties — disabled / 665efe66dcc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-291c676ffe9565f6efa840c7c57d09a42100d4524c9525526ae1bc613ea17f03"></a>

## Next pages — disabled / 665efe66dcc8 / 4

- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-2d46609cf2240b060a6e693c4f0aa7f12184c8e2f9d74abb3bf104f322f18a55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63d20f5aab3892edc609e54e428f5870ff7c7217f6cb7830926d77539aff1152"></a>

## isolated — isolated / 966899434a89 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- isolated

<a id="canonical-8359ab4237d5f38583f0fa323a7c716b06ef98295cb085826c17d0ca71b8dd8e"></a>

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
isolated = {}
```

<a id="canonical-66d5f57a8f850db8479527b8da12eae9b7524214b06355f8b958271092b6cd24"></a>

## Direct properties — isolated / 966899434a89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69f31b451706e0fcc8dcafa8f20aa54a0f2cc356f1765b60f0741c332f33e80b"></a>

## Next pages — isolated / 966899434a89 / 4

- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-7cbca427f616f67b6467a1c59473022dc1de8bb3efc25a376920639794ceea2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef47ad3824f8b812c54cca60bf61954dcc9a4293b1b7fbd071cd8081e7838394"></a>

## timeouts — timeouts / f5bd9243ef5a / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- timeouts

<a id="canonical-09a6c6fdb4c4765341ea198c27fc40f100db2c08099e52713f8a0b530f1e9fa6"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-8b1e3615717eb885dbd687a1a83a1c95b165c5a401cb4e7298013d40c7454972"></a>

## Direct properties — timeouts / f5bd9243ef5a / 3

<a id="canonical-7bc003b8d9aee82135f1b5750d2bc968c6dd3a4e0d82f77f8254f544b50dce3a"></a>

<a id="canonical-e5882d1a69c9710f440268c77a082aa17e4bf703e1c2567f76e8b86f4c7f2f6b"></a>

## create property — timeouts / f5bd9243ef5a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-aaf0aae31d8fc2a986aec629ec1bb3ab6f017d6c34fb44513efbdb70246adb74"></a>

<a id="canonical-c19fead074d2acf49553b5f223e32b9c1527cef3bf3ecaaa9ad052211dacbea9"></a>

## delete property — timeouts / f5bd9243ef5a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-6fea8b3345ab414b1fc32589a70e5df4410be019774a92be61fcd158de051fab"></a>

<a id="canonical-26f63c1f1a8f4a837e904b8e22d27844e7a06b2f8ddc7eefdb13d22e70ccdeb1"></a>

## read property — timeouts / f5bd9243ef5a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a61c07f56c82f19eb60d04a90dfc71b1486cf9767222f22b22e141649fa81ec4"></a>

<a id="canonical-7a082dd519a5ca783c02ba80691b2bafbfe6ae974d26c5bc884c4c40d5b2b004"></a>

## update property — timeouts / f5bd9243ef5a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3426398ff99b7c99aa7e2472bbf547a975408f64e083b0c465d5f631e65b599c"></a>

## Next pages — timeouts / f5bd9243ef5a / 8

- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-aa124a377812667d48addd7e5cc3512af59946547c02528b29fddd94c4ab218d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4172878356fe071fb56fbd44257fc144f98f96ea8e7723fb2ba792bc77e3bee"></a>

## vsite_refs — vsite_refs / 884e501f962f / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- vsite_refs

<a id="canonical-bd24655f8c5bd0b00af6c60238f855bf35ee55ec67eb8d63999dfeae5e2c8647"></a>

Type: `"object"`. list nested block, Optional.

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Upstream description:

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
vsite_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-bdc46a37b528a2b4cb90fb967f34a5b44170db48f44e3fa0ed17b5fbd6d11c3e"></a>

## Direct properties — vsite_refs / 884e501f962f / 3

<a id="canonical-5b9781a2f972ea6d2095036bd9aa3b0c7cd24e64982ba2b0ad35f8c7b560e851"></a>

<a id="canonical-bfe8be435a3c7e18955a5eb394e716b32399e8b3243771af2b4337c53313b57f"></a>

## kind property — vsite_refs / 884e501f962f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-d231028ec731c0a93979d7170b9ff4c96cfc4a9a44c1f09731c202c7638a05e2"></a>

<a id="canonical-c4803a7c179f510f0dff4276e9e541208d5a5e865935c077c3d582817a5a9409"></a>

## name property — vsite_refs / 884e501f962f / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-f2f959b95fa3da3738561df27b8a4321cf06fede8bc48ed78da68ea86153573a"></a>

<a id="canonical-367a5818b04c3b78878de24fe80ccf6d6c3b283b23f8d3b2b21cf07049024569"></a>

## namespace property — vsite_refs / 884e501f962f / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-3af01b71e2152e39b683354a0af823df26a8b8f8772e3b0390ba1bdc690ee6f3"></a>

<a id="canonical-59383d89ebf6b8e4078c523de75fb85414e44210d75ad9265a3f42154e5b51ec"></a>

## tenant property — vsite_refs / 884e501f962f / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-86b3d1c4752c92b7a5c7046a0a7eccbd0970b4edc1f97cf86cb35a7ec8e5da15"></a>

<a id="canonical-597813ca12489164297b75af34de79971095f3419fdcc4d50698d29fa7b3d104"></a>

## uid property — vsite_refs / 884e501f962f / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-ba5b917383983416e056038634ad10765d86ebdc4b9cdc446ece9325ec795fbd"></a>

## Next pages — vsite_refs / 884e501f962f / 9

- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
