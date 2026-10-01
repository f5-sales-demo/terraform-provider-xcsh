---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-aad3ce602761f5ea7b4097d2214f6b252d1a1398b78094d5f155797c584e8f36"></a>

## Next pages — enable_vm / cb864e4bb75f / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0fb68b89f5f7a81145ed1211b5db6d166121e37f3e7e50e5ac8e2e05951d60ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8567c8a05c928f6742d24201e2900340ea9fbdd93984522115dc5e0334abf2ee"></a>

## k8s_cluster — k8s_cluster / a870e59be562 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- k8s_cluster

<a id="canonical-24638a4207f3f67134d482d74f8983eec200e1629db60a6436d499364471f19b"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: k8s\_cluster, no\_k8s\_cluster; Default: no\_k8s\_cluster\] Type establishes a direct
reference from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [k8s_cluster](resources--voltstack_site--reference--group-009.md#canonical-24638a4207f3f67134d482d74f8983eec200e1629db60a6436d499364471f19b)
- [no_k8s_cluster](resources--voltstack_site--reference--group-009.md#canonical-6e7f968f315aed7150d071f19751134cfd3baf18e1132fbae8a5eedb20b43c85)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0bd859e6050292c0b4e8532aca853131439ad74b320e5309e888e92cd6f97c02"></a>

## Direct properties — k8s_cluster / a870e59be562 / 3

<a id="canonical-ee0a78cd024c62c4ae17def9dbf9c7890993a44caf03cfd3d4d2492e746d1ebe"></a>

<a id="canonical-8aa999b5e3f4801737ecc53bf5eac06660b8784c0a196cd7d5bcb33f32ffaa59"></a>

## name property — k8s_cluster / a870e59be562 / 4

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

<a id="canonical-536992c14f945fc1966340fc711006d423e61bf468993af1d5f7a51d8e9c71fc"></a>

<a id="canonical-756692ec0bdea9d4d42c4a511bdf066de9f094fc5065ab41d122618d77ad5178"></a>

## namespace property — k8s_cluster / a870e59be562 / 5

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

<a id="canonical-183c075f59686759d0fccdb7f8916785c0be06b468822e3056cd86b1176d667f"></a>

<a id="canonical-3ff21e721f5d0963c337048bc9f7166b18a36370d98b6b5497d50adb5c7c1d27"></a>

## tenant property — k8s_cluster / a870e59be562 / 6

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

<a id="canonical-2b1381705bf23b0133bb3047446c1322aab687b6d024f8ef73604ef367ccc685"></a>

## Next pages — k8s_cluster / a870e59be562 / 7

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4636974f98edbab43d6780f9440b15c8c9b0cbf6bfbb9d3b493f968fb4a4d1ea"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / ee8f119f0395 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- kubernetes_upgrade_drain

<a id="canonical-d9a16736c134e092c2851db4e36979b279a73ce9b90f271536dcf88cd0026134"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-8be4a829d368c7f2f26dce3090b1340f4bf75102a9ec712dbb4ea1fddb342404"></a>

## Direct properties — kubernetes_upgrade_drain / ee8f119f0395 / 3

- [disable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-f4a90a089272511c31c3fccfe51f3f066b23f41a179aad4de1a76d96828d6385): complete subsection reference.

- [enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b): complete subsection reference.

<a id="canonical-c4f6316b8c390418b880bde7d84f8b0b723eee7be9dc1ed9d5f7c0f8539aa77f"></a>

## Next pages — kubernetes_upgrade_drain / ee8f119f0395 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-f4a90a089272511c31c3fccfe51f3f066b23f41a179aad4de1a76d96828d6385)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f4a90a089272511c31c3fccfe51f3f066b23f41a179aad4de1a76d96828d6385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e095b1b314adfb1623cb35fe63458cdbf69dd7d5102a730a3e0548f6d01ce869"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / 03160e1ff4d7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-5c861daf240147d08cc55c7eff09b16dff1969de604286ef4765f87c3150c107"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-dfc32d1d2b89cbe7c1486a409d348d1f92a348ceafa8f30969d7790929fca54f"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / 03160e1ff4d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d660e5cfcd7b333fbf81ef0a14e6aa51cf107637b7bed6bb921103f39e732cc8"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / 03160e1ff4d7 / 4

- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af7fd47f8fa28711e720622ac9a7885586d356225a7b7c019a110eed55d79dfa"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-cb80585fb7510fa95f8dfdd4ab302405759a8e16c718bf180b13e1b20520ac4f"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-29afa30b57fc751505d5b41a4ff3018459eec8313780d0a1ed915e6d359b9f53"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 3

- [disable_vega_upgrade_mode](resources--voltstack_site--reference--group-009.md#canonical-abca31b4d90e8580eb6be83f5af357332afed49f7cdea30773e61a2e82510d73): complete subsection reference.

<a id="canonical-b7eedfee62f82db7131d433bc70f6bd2952520db3b2fdfd4b99adc6cdf954021"></a>

<a id="canonical-7b850bbb660d50343ff84d1c4cb88e58a25548aef54f748e7e5aeb42f2c4b979"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-cfac5cd75f26ceb54312fc20f2cd0968b4bb0943e160407cf7ceabdf0dbba74b"></a>

<a id="canonical-6b897682ad1343a92118e12a282cb992f30e405c8e78c57cf07a64442b7127e7"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-7ecd72af8c91ddc3c2d41a7eded0dcc32bf7af7fd542a44625b0bfe1046e047f"></a>

<a id="canonical-efe8fa44e0418648aaca3208c10d312a93992cfd9b368f70201b7bf691ff5416"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--voltstack_site--reference--group-009.md#canonical-39d1328953a05ec20e52593c17fd163b97c2ba31e219302f62d317bc45e0e939): complete subsection reference.

<a id="canonical-11fa5f27ebcf70db6cb3afe31dd0aa9dfb2f55c861e4859bca76b946a90a8545"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / cf1b0340800c / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--voltstack_site--reference--group-009.md#canonical-abca31b4d90e8580eb6be83f5af357332afed49f7cdea30773e61a2e82510d73)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--voltstack_site--reference--group-009.md#canonical-39d1328953a05ec20e52593c17fd163b97c2ba31e219302f62d317bc45e0e939)
- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-abca31b4d90e8580eb6be83f5af357332afed49f7cdea30773e61a2e82510d73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38ba238a4ae5f2891b725391de94651b3e27b0290ebcf36fdce06bb52b236e44"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 075427b7f17c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-e195b2edea1d71e715bd47c6e44cba8c074296a55f7e07d93ee9ba7d17baf11d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-1e25fd63e33d4b5c7a6e07ca4f0cc961a7f405a0e747ad0be6426892ec8018ab"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 075427b7f17c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af75d58bfdd401bd89f69e64e44874e28bc69faae8e9caa10bf5c63ed06a0c1e"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 075427b7f17c / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-39d1328953a05ec20e52593c17fd163b97c2ba31e219302f62d317bc45e0e939"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f65498870c823dc71b2b2e267614ec9d771a7e9711c31c940d34297d16e60d52"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 12191f399681 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-01ac514b5d66e595c6659307cbc98556c2c5dad82804f68116fda373ef38cc04"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-5e04905220737b42e8f2ccadfa1f7734deb138a0b7677172d15c0439b34a7aeb"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 12191f399681 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb616974eec616351ff80dfd0671962e4a4c6498f2d3dde29f758a6e063007bb"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 12191f399681 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-c4b52f4c9cee21646d57cc4a072ba84a706921f95f5c7ec0a80cbe0bab46682b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92a9203f62235a97e4d2c9c4e8cf112573275dc4991fd467260a1062ab344689"></a>

## local_control_plane — local_control_plane / 5a91014b9e09 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- local_control_plane

<a id="canonical-b6d955d7292b9947c5a27218e876a24edeed28bb1cf9542f45b47f8dda845c7a"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: local\_control\_plane, no\_local\_control\_plane; Default: no\_local\_control\_plane\]
Enable local control plane for L3VPN, SRV6, EVPN etc.

Upstream description:

Enable local control plane for L3VPN, SRV6, EVPN etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_vn",
    "outside_vn")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_vn\",\"outside_vn\"]"
}
```

OneOf alternatives in this subsection:

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-b6d955d7292b9947c5a27218e876a24edeed28bb1cf9542f45b47f8dda845c7a)
- [no_local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-2d1254303e2c0c57b6d3c235b3025a4f6eeba159f11dfac8ff74589854cd13f0)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
local_control_plane {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e727c7dbdbf27af9a421c290e176cf5985e34ef8f41ca34c3f2b407ca5af126"></a>

## Direct properties — local_control_plane / 5a91014b9e09 / 3

- [bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1): complete subsection reference.

- [inside_vn](resources--voltstack_site--reference--group-009.md#canonical-50d928cb80f8ad30ac90e2ade0d1ca6932b970347e988f606ea3f20a3d763bf6): complete subsection reference.

- [outside_vn](resources--voltstack_site--reference--group-009.md#canonical-2d2c4f8d2142811937ab5694e882f60d181ea6e49e98d6828068740bfa122839): complete subsection reference.

<a id="canonical-a3b6d8cca7bfdaf7d4be5488b9f36a8b8decfc5a80482f97f335186dfc6e9713"></a>

## Next pages — local_control_plane / 5a91014b9e09 / 4

- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.inside_vn](resources--voltstack_site--reference--group-009.md#canonical-50d928cb80f8ad30ac90e2ade0d1ca6932b970347e988f606ea3f20a3d763bf6)
- [local_control_plane.outside_vn](resources--voltstack_site--reference--group-009.md#canonical-2d2c4f8d2142811937ab5694e882f60d181ea6e49e98d6828068740bfa122839)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8e96bfca9e69b856a8c4a83c637c79d6f26a5aa4a8278ad5e2118790ec36fd5"></a>

## local_control_plane.bgp_config — local_control_plane.bgp_config / e1a81de93d13 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- local_control_plane.bgp_config

<a id="canonical-b73bbd72af8a91111620d61da9e66ed77d787d1654d073cfbdf28ba6a56d41fb"></a>

Type: `"object"`. single nested block, Optional.

BGP Configuration. BGP configuration parameters.

Upstream description:

BGP configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn")}
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
bgp_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-4e27bfbb2694ecdf8ea58a03f8818a760eb5c3a1efc562dadf35ec4d0e86e60c"></a>

## Direct properties — local_control_plane.bgp_config / e1a81de93d13 / 3

<a id="canonical-b816d4911bef0c9b3c470b54aa9952c3dd813bf90ac01267bab35215d79dbebe"></a>

<a id="canonical-4a08936433f7baa3c44325bb824053ee3ec3db874f26d0cee9b623f34737d287"></a>

## asn property — local_control_plane.bgp_config / e1a81de93d13 / 4

Type: `"number"`. Optional.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c): complete subsection reference.

<a id="canonical-4d7e0dab6a291f8fb11f0b194fc3a7fcce37c5a91961d9b025cac28041bd6197"></a>

## Next pages — local_control_plane.bgp_config / e1a81de93d13 / 5

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f338ecf36671ce6ef4b305b7083d17dc8c9ef73e846d113504b3461db5680475"></a>

## local_control_plane.bgp_config.peers — local_control_plane.bgp_config.peers / 5f17ab6db491 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- local_control_plane.bgp_config.peers

<a id="canonical-64315d856445b24eecddbe68dbbe2fe205f0a3f8a20200ef5fc14e1eea5b6875"></a>

Type: `"object"`. list nested block, Optional.

Peers. BGP parameters for peer.

Upstream description:

BGP parameters for peer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bfd_disabled",
    "bfd_enabled"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "routing_policies"),
  validators.ConflictingListObjectAttributes("ebgp_multihop_disabled",
    "ebgp_multihop_enabled"),
  validators.ConflictingListObjectAttributes("passive_mode_disabled",
    "passive_mode_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
peers {
  # Configure direct properties listed below.
}
```

<a id="canonical-9616417ebd6f1b6277865318c09e44eaff57bd4ac3bf755790f591d3e4bb3c1c"></a>

## Direct properties — local_control_plane.bgp_config.peers / 5f17ab6db491 / 3

- [bfd_disabled](resources--voltstack_site--reference--group-009.md#canonical-035691489f1bfeb809ec389ac909de23532203476bdef79ffd68544d387fd224): complete subsection reference.

- [bfd_enabled](resources--voltstack_site--reference--group-009.md#canonical-48f8f02e001750028b922501f17456fe447ae8290ba92ed993c271eae6ed6ef9): complete subsection reference.

- [disable_spec](resources--voltstack_site--reference--group-009.md#canonical-03d404c2fa126e5d4b7e0177a2db328004f6d78680d38e245650a54744aa7bb8): complete subsection reference.

- [ebgp_multihop_disabled](resources--voltstack_site--reference--group-009.md#canonical-56b294de78a86967de895de98b4e5d7d82a13dcb0b09d65486ed37e19e56c5d0): complete subsection reference.

- [ebgp_multihop_enabled](resources--voltstack_site--reference--group-009.md#canonical-5f067fbf8bbe6a0b0f53ae413efa06923ebb4d5bdbb3cef52c876228b3ce7a84): complete subsection reference.

- [external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0): complete subsection reference.

<a id="canonical-b3186db401683e19d1cad1f691b1d1a6d2e0acb8d9c1e1dcf4e499a211959240"></a>

<a id="canonical-9ca9c1e397afc2e558894002b7b51a63da097e3320a6eaa89d742614d068542e"></a>

## label property — local_control_plane.bgp_config.peers / 5f17ab6db491 / 4

Type: `"string"`. Optional.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--voltstack_site--reference--group-009.md#canonical-f17b01f90720991047a1029f90e2d8403a2b7cad7d3fd44b0c034806f0c266f6): complete subsection reference.

- [passive_mode_disabled](resources--voltstack_site--reference--group-009.md#canonical-1dc593ec427c6d0406f66e5bcdfd5559bc40783be3e95e9ba3798c11cdb9e0d5): complete subsection reference.

- [passive_mode_enabled](resources--voltstack_site--reference--group-009.md#canonical-da39b18bfae933a1126eba6844a5430a7831511873f28be9943b90c9bd5a7b8b): complete subsection reference.

- [routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262): complete subsection reference.

<a id="canonical-428a92cd0e422c0e8b94391d6cd65073b3e94a68373f42c142f30647c2247fe2"></a>

## Next pages — local_control_plane.bgp_config.peers / 5f17ab6db491 / 5

- [local_control_plane.bgp_config.peers.bfd_disabled](resources--voltstack_site--reference--group-009.md#canonical-035691489f1bfeb809ec389ac909de23532203476bdef79ffd68544d387fd224)
- [local_control_plane.bgp_config.peers.bfd_enabled](resources--voltstack_site--reference--group-009.md#canonical-48f8f02e001750028b922501f17456fe447ae8290ba92ed993c271eae6ed6ef9)
- [local_control_plane.bgp_config.peers.disable_spec](resources--voltstack_site--reference--group-009.md#canonical-03d404c2fa126e5d4b7e0177a2db328004f6d78680d38e245650a54744aa7bb8)
- [local_control_plane.bgp_config.peers.ebgp_multihop_disabled](resources--voltstack_site--reference--group-009.md#canonical-56b294de78a86967de895de98b4e5d7d82a13dcb0b09d65486ed37e19e56c5d0)
- [local_control_plane.bgp_config.peers.ebgp_multihop_enabled](resources--voltstack_site--reference--group-009.md#canonical-5f067fbf8bbe6a0b0f53ae413efa06923ebb4d5bdbb3cef52c876228b3ce7a84)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.metadata](resources--voltstack_site--reference--group-009.md#canonical-f17b01f90720991047a1029f90e2d8403a2b7cad7d3fd44b0c034806f0c266f6)
- [local_control_plane.bgp_config.peers.passive_mode_disabled](resources--voltstack_site--reference--group-009.md#canonical-1dc593ec427c6d0406f66e5bcdfd5559bc40783be3e95e9ba3798c11cdb9e0d5)
- [local_control_plane.bgp_config.peers.passive_mode_enabled](resources--voltstack_site--reference--group-009.md#canonical-da39b18bfae933a1126eba6844a5430a7831511873f28be9943b90c9bd5a7b8b)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-035691489f1bfeb809ec389ac909de23532203476bdef79ffd68544d387fd224"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-178fcded7e9081bf27a4a67a2efb141a9165a1c64e94550e6bd2aad3075782b0"></a>

## local_control_plane.bgp_config.peers.bfd_disabled — local_control_plane.bgp_config.peers.bfd_disabled / b47d3964ff01 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.bfd_disabled

<a id="canonical-f7d6c3805b1c7405af88c652a650ed6d7f4348a941e600b284e8124e4d2564e9"></a>

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
bfd_disabled = {}
```

<a id="canonical-a5810794a369ff2f86ff4dbebd21f542d5b0a3136f905f69ac48018af3efb06a"></a>

## Direct properties — local_control_plane.bgp_config.peers.bfd_disabled / b47d3964ff01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc7542e213c0979a5d3a3873cbb1f8594a799135a5756b995ee04db95655476b"></a>

## Next pages — local_control_plane.bgp_config.peers.bfd_disabled / b47d3964ff01 / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-48f8f02e001750028b922501f17456fe447ae8290ba92ed993c271eae6ed6ef9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32c37f5dfab816d5ab59d1655aae9a0763bc1ad6f469bbd0f4bea9e8eccad88c"></a>

## local_control_plane.bgp_config.peers.bfd_enabled — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.bfd_enabled

<a id="canonical-8316b014612d3e1d3e13171bd471c81120576fde76c512cae5924f248d0556a2"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Upstream description:

BFD parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
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
bfd_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4896c9432fb50fff0ec7efc6ba09ca776387fd05c9a4ab709189d00149d26a0"></a>

## Direct properties — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 3

<a id="canonical-0264909b5a52b8cf3c2bd930fd43a904bebfb28bf9e5e3756a942127e4b34e2e"></a>

<a id="canonical-9e165d623ed494d9fb1d5fd9a6114a6990021a6cea5d2d908a7a9d63bc2c69c0"></a>

## multiplier property — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 4

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-fe313cddc236bea279dbe79ddcceb5611b13a81e07a01c4a93161b639dfc81f5"></a>

<a id="canonical-9a6b0d3dac2f6a1f148cd7a030e4153ba08b7a9a4442b397c5ca007d8b08e912"></a>

## receive_interval_milliseconds property — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 5

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-f1610e8a78b41bcc998656a0e5974fd4c5c73499bbf1e0f2650d7c716d7fd7ea"></a>

<a id="canonical-7f876fbea429c5879731db9cbc45011060498029cdc166bc2441a53aeac7e49d"></a>

## transmit_interval_milliseconds property — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 6

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-956667412cc0d5e0aa76bf16135cced0e9f38a48688395df510690217c7a14f8"></a>

## Next pages — local_control_plane.bgp_config.peers.bfd_enabled / be7cda362d2d / 7

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-03d404c2fa126e5d4b7e0177a2db328004f6d78680d38e245650a54744aa7bb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9b31e080bb9c8a4632921b89970f260848e43d5774844f0c8a0928eb1a99f32"></a>

## local_control_plane.bgp_config.peers.disable_spec — local_control_plane.bgp_config.peers.disable_spec / 57a9b1e90390 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.disable_spec

<a id="canonical-c82fb6ac4e0c50f2e118db97747665113d28b88b3ad7ae77cef241faa37afcfb"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-c8ae5b376248c3e3ddd5a48aab9128041475880e7b3c77e64db8602dab408ec7"></a>

## Direct properties — local_control_plane.bgp_config.peers.disable_spec / 57a9b1e90390 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19786b06e60ea1710d524e3b932fdca1aa25fac22cc0706a327a8d0d6ef1c465"></a>

## Next pages — local_control_plane.bgp_config.peers.disable_spec / 57a9b1e90390 / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-56b294de78a86967de895de98b4e5d7d82a13dcb0b09d65486ed37e19e56c5d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91c9bdc88d00201e6054192d10d4865724839db3eaf616c723ae0a8ac848e1d5"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_disabled — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 311f47c8107f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.ebgp_multihop_disabled

<a id="canonical-c4ef07c627e5dc6339b108de82f653c40d11b7b7895084e6962d32453545d6f1"></a>

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
ebgp_multihop_disabled = {}
```

<a id="canonical-4b8a64f9f0a25554165e2e51822b01606bbc0b66a3742d8a09720461c1b103af"></a>

## Direct properties — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 311f47c8107f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98c5d000f2bacca7ccde611bc07027bef127e92e9a82a728f70d52af8e17f5ea"></a>

## Next pages — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 311f47c8107f / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5f067fbf8bbe6a0b0f53ae413efa06923ebb4d5bdbb3cef52c876228b3ce7a84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a7bc0c97b67030ef5c8c042990862e0eef2ae9f9e8289a35574534fde31687e"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_enabled — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 7b3e298270ea / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.ebgp_multihop_enabled

<a id="canonical-a36ea99eb1628fe54f815887ad0210754d5be43437c6afab2b2ad7505fbb91b3"></a>

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
ebgp_multihop_enabled = {}
```

<a id="canonical-16019fc36eb4c6effadceab07f061437bbe6fa40199f4fe9b39e8066d746c6b1"></a>

## Direct properties — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 7b3e298270ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-430c9e59da6ebfc24ee7cd65cd0eef5b370c81888c15383317ed7d69ca71d1ba"></a>

## Next pages — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 7b3e298270ea / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-303c2838d38065d1061072bad1f7fab7dc1189e3a96a793f37b7c3b0ff9455e4"></a>

## local_control_plane.bgp_config.peers.external — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.external

<a id="canonical-df7ecc478caf3ad3bc933612346674ed3c9a98c5d00a92baf5523ae18d46f153"></a>

Type: `"object"`. single nested block, Optional.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn",
    "port"),
  validators.ConflictingObjectAttributes("address",
    "default_gateway"),
  validators.ConflictingObjectAttributes("address",
    "disable_spec"),
  validators.ConflictingObjectAttributes("address",
    "external_connector"),
  validators.ConflictingObjectAttributes("address",
    "from_site"),
  validators.ConflictingObjectAttributes("address",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("address",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "default_gateway_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_gateway",
    "external_connector"),
  validators.ConflictingObjectAttributes("default_gateway",
    "from_site"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("disable_spec",
    "external_connector"),
  validators.ConflictingObjectAttributes("disable_spec",
    "from_site"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("disable_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("external_connector",
    "from_site"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("interface",
    "interface_list"),
  validators.ConflictingObjectAttributes("md5_auth_key",
    "no_authentication"),
  validators.ConflictingObjectAttributes("subnet_begin_offset",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("subnet_begin_offset_v6",
    "subnet_end_offset_v6")}
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
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

Terraform syntax:

```terraform
external {
  # Configure direct properties listed below.
}
```

<a id="canonical-0dcf6f272d7d9ca0a828937f0d1b4a61a61a66650a213fd7285ab666eeed1cfa"></a>

## Direct properties — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 3

<a id="canonical-af577078fa95e7e2e10cb1f4fec61bd68e67a20c7f4f8f59206f08bd5ba53e23"></a>

<a id="canonical-9a940c0f3c9a32027dff20b6d5f3651eb103cd96ef34195e5e80e85fab57c52c"></a>

## address property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-f616eecba9e285fa2f6d758c6cae031a863e8b976a9328d53a65d717cbd910f3"></a>

<a id="canonical-bc3153f3c8387222279b01619231e8bdb16103ccda5c3422d75b9d5a024d293b"></a>

## address_ipv6 property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-209ae36d67a7e4aa29db623090e7767c212f3817cc2e166625c0c83b5c06b297"></a>

<a id="canonical-3137e5fafa284613ede41ea11dada1a12098dd0b624d7c9ba96611edab82fc02"></a>

## asn property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 6

Type: `"number"`. Optional.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-009.md#canonical-978604a54997867f769fb29c1291d679042882d30ec7802343df0f764813b9ce): complete subsection reference.

- [default_gateway_v6](resources--voltstack_site--reference--group-009.md#canonical-ca3da680871bfb44b3009cc6f30556ca5206960c69efc97a3bde613eee13d6a4): complete subsection reference.

- [disable_spec](resources--voltstack_site--reference--group-009.md#canonical-7544c2f9e2ce342aabd5180eb188921f9e81cd7fa7ef27a624bd4cd991c19778): complete subsection reference.

- [disable_v6](resources--voltstack_site--reference--group-009.md#canonical-ac3e20869fe55c01fa0ae06ece98cdd785da53dd5b977a56496589fd1f081192): complete subsection reference.

- [external_connector](resources--voltstack_site--reference--group-009.md#canonical-8ede8b9b97e72cbfadab4e20f646e7d95aa6737824ec71bbe99ca9545fc2a4de): complete subsection reference.

- [family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5): complete subsection reference.

- [from_site](resources--voltstack_site--reference--group-009.md#canonical-dc9a15370a8564e4fcd7900e2929ee301538b3aa2cd766bd615973845540edef): complete subsection reference.

- [from_site_v6](resources--voltstack_site--reference--group-009.md#canonical-ba2b8f9374eecd8a72ab80f74384fd4d2ec827274cfbcfd7b932362fcccf4e2a): complete subsection reference.

- [interface](resources--voltstack_site--reference--group-009.md#canonical-9e67b24f71c2cd22f4c18901ef398b83e25e23bce9af799e0d1429f6249581b2): complete subsection reference.

- [interface_list](resources--voltstack_site--reference--group-009.md#canonical-2f23201d864e821b2dac5be261e8da687697694f4694387e2229c1b732929ab4): complete subsection reference.

<a id="canonical-e7a18443d874366833e779ff9e9e828d45221fd2cefd00cc4fa92f5aa77437fd"></a>

<a id="canonical-3dcd1d0c6a3b1da27b3c7ee9e50e87c952e7394943e626166ae304b0e6437736"></a>

## md5_auth_key property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 7

Type: `"string"`. Optional.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](resources--voltstack_site--reference--group-009.md#canonical-bf787f3a568ad5e358e88f3374f9642ecfbda707330d6a86e313a024a6c89671): complete subsection reference.

<a id="canonical-81733d040e9faebf4dd05ebe770b945a7931f6a47704dbef3054acc1208abfdc"></a>

<a id="canonical-5ffe8ddbaf9e464480fde0e164611feecf6b2c18b9a29098d6a9686530016fd3"></a>

## port property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 8

Type: `"number"`. Optional.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b1c08fb0f5265600fe54b9becbb3110fa21de21fd355164390916d329051f196"></a>

<a id="canonical-f4fc091250bb1ebb3df90ceac929b9dae7497fc96857dd3408def53ecb65b39e"></a>

## subnet_begin_offset property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 9

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-6a5e5de5c8241f28bfbc0f6b08bdc0c44b6719a5251ea8a557de1b26cc9eb424"></a>

<a id="canonical-a15f467345e94ba10abe5707f9dd03e45df1e3c860a282219640e8e5719a57c8"></a>

## subnet_begin_offset_v6 property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 10

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-772ead3cd170aa581d86a8539c4e95a2b2f01349069b0486b64b44b781babd83"></a>

<a id="canonical-a3dc88ef685852d120f025adcc70df637997ce840d8015a32746922334167195"></a>

## subnet_end_offset property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 11

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-c594445b11bea9cc4c23c475362d054f0693183c8dbc22ea9dc46decc8d46add"></a>

<a id="canonical-0ebe1b7e628a89cd156b6f6788e24b493d6c26645b4c811b75b59725de98bf98"></a>

## subnet_end_offset_v6 property — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 12

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-c019e68a5e85d6d5c0b9764a0aaee20a9a4a375244f632519cff0c9bed10d460"></a>

## Next pages — local_control_plane.bgp_config.peers.external / 4a9aff1a0011 / 13

- [local_control_plane.bgp_config.peers.external.default_gateway](resources--voltstack_site--reference--group-009.md#canonical-978604a54997867f769fb29c1291d679042882d30ec7802343df0f764813b9ce)
- [local_control_plane.bgp_config.peers.external.default_gateway_v6](resources--voltstack_site--reference--group-009.md#canonical-ca3da680871bfb44b3009cc6f30556ca5206960c69efc97a3bde613eee13d6a4)
- [local_control_plane.bgp_config.peers.external.disable_spec](resources--voltstack_site--reference--group-009.md#canonical-7544c2f9e2ce342aabd5180eb188921f9e81cd7fa7ef27a624bd4cd991c19778)
- [local_control_plane.bgp_config.peers.external.disable_v6](resources--voltstack_site--reference--group-009.md#canonical-ac3e20869fe55c01fa0ae06ece98cdd785da53dd5b977a56496589fd1f081192)
- [local_control_plane.bgp_config.peers.external.external_connector](resources--voltstack_site--reference--group-009.md#canonical-8ede8b9b97e72cbfadab4e20f646e7d95aa6737824ec71bbe99ca9545fc2a4de)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [local_control_plane.bgp_config.peers.external.from_site](resources--voltstack_site--reference--group-009.md#canonical-dc9a15370a8564e4fcd7900e2929ee301538b3aa2cd766bd615973845540edef)
- [local_control_plane.bgp_config.peers.external.from_site_v6](resources--voltstack_site--reference--group-009.md#canonical-ba2b8f9374eecd8a72ab80f74384fd4d2ec827274cfbcfd7b932362fcccf4e2a)
- [local_control_plane.bgp_config.peers.external.interface](resources--voltstack_site--reference--group-009.md#canonical-9e67b24f71c2cd22f4c18901ef398b83e25e23bce9af799e0d1429f6249581b2)
- [local_control_plane.bgp_config.peers.external.interface_list](resources--voltstack_site--reference--group-009.md#canonical-2f23201d864e821b2dac5be261e8da687697694f4694387e2229c1b732929ab4)
- [local_control_plane.bgp_config.peers.external.no_authentication](resources--voltstack_site--reference--group-009.md#canonical-bf787f3a568ad5e358e88f3374f9642ecfbda707330d6a86e313a024a6c89671)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-978604a54997867f769fb29c1291d679042882d30ec7802343df0f764813b9ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05aa505bd89085d6f8236e2448c7aab6a965c1f74a138e6fcd9489f4e565fc7c"></a>

## local_control_plane.bgp_config.peers.external.default_gateway — local_control_plane.bgp_config.peers.external.default_gateway / 5173601f4064 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.default_gateway

<a id="canonical-c223cbb50e156717382cbcde74fb5aa01d43503fae2672a2560d2508858db93b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-198bc4306fbec49d6c0cfb4bc77a4ccf5eb182edacf3e89c6ec913e374fcc6ac"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.default_gateway / 5173601f4064 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8acd3dedad94778c5045ae9b238b632e97be0f237233f0fa1717e761141eb992"></a>

## Next pages — local_control_plane.bgp_config.peers.external.default_gateway / 5173601f4064 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ca3da680871bfb44b3009cc6f30556ca5206960c69efc97a3bde613eee13d6a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18099103b3e5e1ae6154f908ad86ef852f6eddc25ae45b5fb673f491284aec43"></a>

## local_control_plane.bgp_config.peers.external.default_gateway_v6 — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 9534dc8be64f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.default_gateway_v6

<a id="canonical-2b4d7dbc1c595e22b54ad67ba82377f3b2fb4dde8502cfbcfd3caddcff24a76c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway v6.

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
default_gateway_v6 = {}
```

<a id="canonical-f4559068ff7b44960a792fece018e0a76bf2d95b7c90442b86923238d651c4ac"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 9534dc8be64f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-828a632ff4882163d416c6430d6b31059c916246f285cd0ef78f439b3eb488de"></a>

## Next pages — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 9534dc8be64f / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-7544c2f9e2ce342aabd5180eb188921f9e81cd7fa7ef27a624bd4cd991c19778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-594bf32284a078bc0c069d5914aefd9f6ebe66af18b26234b66f10e49d34d712"></a>

## local_control_plane.bgp_config.peers.external.disable_spec — local_control_plane.bgp_config.peers.external.disable_spec / e6fc0d319332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.disable_spec

<a id="canonical-cbc6043b483632b6e7c0ca27b1bef65abab1c59ca4253adc0ec83038e7c4f88e"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-fa800f322c0ba1f279c347e4135a562a505a29e674a5a6870485e1d736b12764"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.disable_spec / e6fc0d319332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5896c45a6ee2a38fe710acd6e68e840eeb8444d473d15e046772e0658aec6e4"></a>

## Next pages — local_control_plane.bgp_config.peers.external.disable_spec / e6fc0d319332 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ac3e20869fe55c01fa0ae06ece98cdd785da53dd5b977a56496589fd1f081192"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1004565e8745e36a3fae0050c083987e5ffeb82301533296bf4baca94d28f971"></a>

## local_control_plane.bgp_config.peers.external.disable_v6 — local_control_plane.bgp_config.peers.external.disable_v6 / 094dbf666231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.disable_v6

<a id="canonical-ac60234a2512b6e0cf04ceb697a8304233ce940e56568f8096e6fe0d664875f3"></a>

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
disable_v6 = {}
```

<a id="canonical-620e0e996ec09d65667c705ff5a3aa0c1f0cf4919f9c243a8d979e575981166c"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.disable_v6 / 094dbf666231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3918dd7faac54fb7ce83a1803b8a9add26e6a748f94e55f63c14829c52e9fdb"></a>

## Next pages — local_control_plane.bgp_config.peers.external.disable_v6 / 094dbf666231 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8ede8b9b97e72cbfadab4e20f646e7d95aa6737824ec71bbe99ca9545fc2a4de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-329405590cf507f6917ad4340662002ffa4713f572e22b979d083a45543c778a"></a>

## local_control_plane.bgp_config.peers.external.external_connector — local_control_plane.bgp_config.peers.external.external_connector / d7b686ef9182 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.external_connector

<a id="canonical-3d69573f84275fb92ed8827d60e66e3fc83d92756507e30fda22ebcc00ed998a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external connector.

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
external_connector = {}
```

<a id="canonical-911b1ce353ce33248e31617fc226fccb54fdc7c79b6e907335c032faadf15159"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.external_connector / d7b686ef9182 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a14315ce00649b8d5e31c5211dbd27f9c415e8c84e9347834f3ccfa129868947"></a>

## Next pages — local_control_plane.bgp_config.peers.external.external_connector / d7b686ef9182 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7769fa4685f8a4746acea9c67230020fb59fcf45d72b9939770ddd3ca5b1ab4"></a>

## local_control_plane.bgp_config.peers.external.family_inet — local_control_plane.bgp_config.peers.external.family_inet / 5af58ca3ecbb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.family_inet

<a id="canonical-7c4169704f33ffbe97c6efbb320041b534b4b3c6f515b6abcd6f2ae36b235d59"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

<a id="canonical-83e44e476f4d502dd2d1a7704a4f3fcbbc4a0e678d3def2c91cf326157a218fa"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet / 5af58ca3ecbb / 3

- [disable_spec](resources--voltstack_site--reference--group-009.md#canonical-0b06efc6f2b37b140e17943589bc9f5447a031052c11b7155ae870908fdcccd4): complete subsection reference.

- [enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2): complete subsection reference.

<a id="canonical-05c52b31b140a230643de58ada3bd027df7d65d8ac20b130a33bb0ba44f67507"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet / 5af58ca3ecbb / 4

- [local_control_plane.bgp_config.peers.external.family_inet.disable_spec](resources--voltstack_site--reference--group-009.md#canonical-0b06efc6f2b37b140e17943589bc9f5447a031052c11b7155ae870908fdcccd4)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0b06efc6f2b37b140e17943589bc9f5447a031052c11b7155ae870908fdcccd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25619e59a712f2e2e785ce1bd9717e2dc19561226c3d9ecbc7fdde26a3dfe021"></a>

## local_control_plane.bgp_config.peers.external.family_inet.disable_spec — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / 138e017aa44c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- local_control_plane.bgp_config.peers.external.family_inet.disable_spec

<a id="canonical-99b0096dacff20630c5186f63f0c8aa2d3bc3e835b00dba1d4579e4d4c78fb1c"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-8499d1e3b4705dd9a7a3091864d5679acca7001cd6066f9a8c3a220d0f275af2"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / 138e017aa44c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d0076e65d50e6ffb4d5e8f93aa41947413d4c9455b66d83a59811334ca37124"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / 138e017aa44c / 4

- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-081ead74ac442195e0cfd740a0592283dd935f73c5c93ceb109d7dc8c95516c5"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable — local_control_plane.bgp_config.peers.external.family_inet.enable / 2ee578dc4818 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- local_control_plane.bgp_config.peers.external.family_inet.enable

<a id="canonical-4f9eeb7ed319be688d972f4eee98d82628b145b63204113e2ec2aba5fe6e595c"></a>

Type: `"object"`. single nested block, Optional.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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
enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-95045a6b2faf042a8dc960aad25316cea6accca6c419a83566a5f4b7a75e57c4"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable / 2ee578dc4818 / 3

- [aggregation](resources--voltstack_site--reference--group-009.md#canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb): complete subsection reference.

<a id="canonical-4f3931a86f6d4ec4ed38be8696113f9b91c2c7234056be1820200f76a92d4ad5"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable / 2ee578dc4818 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--reference--group-009.md#canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e35ef0d95151fba11757f6cca4d4a8c4428e9b48e1050b63238b861fdffcd01"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / dabc15eb5ab8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

<a id="canonical-0ff7b50428339c58879188aaf332f0528d4049811df01e7e296d8f56958406c9"></a>

Type: `"object"`. list nested block, Optional.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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
aggregation {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e5ca41d9081a2b9db8b03ac3de404556a27780ac7860358e05632a0f2d599dd"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / dabc15eb5ab8 / 3

<a id="canonical-c2c5058bc333ed5adce3b9f1f78dd4b951949e77b81149dd8fbe33625251817e"></a>

<a id="canonical-3aa1683aba7b3614ed4c4af404776195723e2a4a8d469166a99a1ec6f309b30c"></a>

## ip_prefix property — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / dabc15eb5ab8 / 4

Type: `"string"`. Optional.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

- [options](resources--voltstack_site--reference--group-009.md#canonical-777ca09f63e8b16a12447a46dc2eeb7abdd5cf3897b23908cb924944f0057df8): complete subsection reference.

<a id="canonical-0915c450bdcb2360fa0c1ba26ed5344ed5f12ada6eae1b56bbe7d5cc402eec77"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / dabc15eb5ab8 / 5

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](resources--voltstack_site--reference--group-009.md#canonical-777ca09f63e8b16a12447a46dc2eeb7abdd5cf3897b23908cb924944f0057df8)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-777ca09f63e8b16a12447a46dc2eeb7abdd5cf3897b23908cb924944f0057df8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-578a6af881fac4ad6ef8ce51c5b45aafdfa432cac0c9764df2cc5bf3f0475cd8"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / a1bf388bafbd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--reference--group-009.md#canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

<a id="canonical-9eb30233c3128e265b3922493746a4eba9821ed475af387a04113a6317e13357"></a>

Type: `"object"`. list nested block, Optional.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

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
options {
  # Configure direct properties listed below.
}
```

<a id="canonical-9e912c13f17bfb558518a104da0cba377850e5dc5ed0ce3d9db74ff5fa447601"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / a1bf388bafbd / 3

- [summary_only](resources--voltstack_site--reference--group-009.md#canonical-9d55e5c1bc70f2323defa665ca6fb80159aa734205531ddf70f1532afe50afbc): complete subsection reference.

<a id="canonical-0b91aaf042ef8fa85ce82492a09c6705aacdea0a216e690bba925b38fb0086af"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / a1bf388bafbd / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only](resources--voltstack_site--reference--group-009.md#canonical-9d55e5c1bc70f2323defa665ca6fb80159aa734205531ddf70f1532afe50afbc)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--reference--group-009.md#canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9d55e5c1bc70f2323defa665ca6fb80159aa734205531ddf70f1532afe50afbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-497df003d2bea1cb348920c8200953a6af6161e96418c9f436cb825fee0ffa6d"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 39a6870366e1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--reference--group-009.md#canonical-1d627372d0a598519a94522340dac4c2b1ca80e04ba9a3620401767410e6b1c5)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--reference--group-009.md#canonical-20477c64c309dc81d11a54eefaafe64742e880b3970bbd489fbce6f22407b8b2)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--reference--group-009.md#canonical-b6d17049c78f404ffb1ab20675682be93e91563ba9357406cd1ddd507c9b81cb)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](resources--voltstack_site--reference--group-009.md#canonical-777ca09f63e8b16a12447a46dc2eeb7abdd5cf3897b23908cb924944f0057df8)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-ca5337690c8c8560b50bf7003bbbce370f161344a9cc04fd46d9d04bd4d8de0a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for summary only.

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
summary_only {}
```

<a id="canonical-cdd652c6e522f4fdc87b2d1c0421d2432f2976f68221774f0273e1b32e9877d9"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 39a6870366e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3795e36b4ed9b0a97bbededf438ba7cab17ed6ce3a1243ac3a1eb0ac7893b71"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 39a6870366e1 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](resources--voltstack_site--reference--group-009.md#canonical-777ca09f63e8b16a12447a46dc2eeb7abdd5cf3897b23908cb924944f0057df8)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-dc9a15370a8564e4fcd7900e2929ee301538b3aa2cd766bd615973845540edef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ab1f1d486ab8982770dc759406c2804bfed17c866b7fada3c86091c06376620"></a>

## local_control_plane.bgp_config.peers.external.from_site — local_control_plane.bgp_config.peers.external.from_site / a811c9fcdb42 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.from_site

<a id="canonical-3f3f896a59b6f647e410da83755305dae046a97a6e8ca2313c541e16b0077aae"></a>

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
from_site = {}
```

<a id="canonical-179b504b0a08fd2fd24319f258213ea63597ecd33ee87952569752a9b8675063"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.from_site / a811c9fcdb42 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-047ca71568266af23d0c0ad9428cdb03d0f952c293bc49affc4f8105a51920e3"></a>

## Next pages — local_control_plane.bgp_config.peers.external.from_site / a811c9fcdb42 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ba2b8f9374eecd8a72ab80f74384fd4d2ec827274cfbcfd7b932362fcccf4e2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-376898b52fdee3c4859de21a9d7343e9fb58d3ecfa931ade8a2e91bdf37291a2"></a>

## local_control_plane.bgp_config.peers.external.from_site_v6 — local_control_plane.bgp_config.peers.external.from_site_v6 / 152aa0e80023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.from_site_v6

<a id="canonical-bfd75dc0bb2276424dc3242f8d70b3ff795ebebfa2c9dcc00b32d26850f13e41"></a>

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
from_site_v6 = {}
```

<a id="canonical-d6593397acc76f7eb02b2cc5968b5c182b97cf85a5709402fc494b49b60c0b44"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.from_site_v6 / 152aa0e80023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e50cdea9b5915fee0923d0ad97f272a8f1d697bb5719ed0eebdf76af00e5d46"></a>

## Next pages — local_control_plane.bgp_config.peers.external.from_site_v6 / 152aa0e80023 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9e67b24f71c2cd22f4c18901ef398b83e25e23bce9af799e0d1429f6249581b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e50e66acf663e1db2153f12d6f62175e437e9c5a688449116384914ee04a7f70"></a>

## local_control_plane.bgp_config.peers.external.interface — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.interface

<a id="canonical-99258eb150e5d7c8618b3114b281574db0391a705bc1a08cf62ee51d86060406"></a>

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-a80b1188c10aae09f1bc23eaaef63c8a6f8e1c9c8b6a1e6bb87eb47a3605ab51"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 3

<a id="canonical-bc98308cf9f49565fc848626dd5d5713841dd404a9564e235b3e1c5e4b0122c9"></a>

<a id="canonical-f6aa8873a12a04355d929ed628e67a37c49beeeb69f61075b81738ea7a4cce8f"></a>

## name property — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 4

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

<a id="canonical-684af975058f6882d4c3a2e65378cd3ca821b45d09940dd7d5a146658b3c5793"></a>

<a id="canonical-2f49824f49d4e9e6be576cdecae85d3818c8069d0f24615f872ae917fb3be6dd"></a>

## namespace property — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 5

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

<a id="canonical-714e6c268673c860ed6cc5402d754a75039e0a66312a22effb51ff780b628dd3"></a>

<a id="canonical-984fc9540b9e0e4236e57e69d584a13f9d3ec365c76f0113c839390933ad71b0"></a>

## tenant property — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 6

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

<a id="canonical-00511481a8f9f5a23901b5b7d10bc0aa8a5a0f14afdc9dcf8464694cdc2e31b3"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface / 272533d05f9c / 7

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-2f23201d864e821b2dac5be261e8da687697694f4694387e2229c1b732929ab4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a48ee85775dd9341b4dc0e18c4540d22765c95a41f6fb6ef0847972ed3a8bb6"></a>

## local_control_plane.bgp_config.peers.external.interface_list — local_control_plane.bgp_config.peers.external.interface_list / 81db67f52416 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.interface_list

<a id="canonical-93d9bc9393250d41446034a44666bdef6083d2fe98ce5db6f62c2f72cd60b0cb"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-b7d68552e11eca119dc709cdb0eaad6358f1fb8eeaeaa3448525438f894675a3"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface_list / 81db67f52416 / 3

- [interfaces](resources--voltstack_site--reference--group-009.md#canonical-a379919a876c9d2b740d5634bd0f55ea03df1910d1367d1a502394c977e57d52): complete subsection reference.

<a id="canonical-b91853db5db8eaefe42b28fa8ea4230210d7d55611f90d8f37b894f111722b26"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface_list / 81db67f52416 / 4

- [local_control_plane.bgp_config.peers.external.interface_list.interfaces](resources--voltstack_site--reference--group-009.md#canonical-a379919a876c9d2b740d5634bd0f55ea03df1910d1367d1a502394c977e57d52)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-a379919a876c9d2b740d5634bd0f55ea03df1910d1367d1a502394c977e57d52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95caece702ecba625daa785c93d93f8e9314cc8a6a849df4d9c32fed180182c8"></a>

## local_control_plane.bgp_config.peers.external.interface_list.interfaces — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [local_control_plane.bgp_config.peers.external.interface_list](resources--voltstack_site--reference--group-009.md#canonical-2f23201d864e821b2dac5be261e8da687697694f4694387e2229c1b732929ab4)
- local_control_plane.bgp_config.peers.external.interface_list.interfaces

<a id="canonical-f7a1c317e63ad1d56a5748ec11829d27319e09013d36f07eb74ea467803e9307"></a>

Type: `"object"`. list nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-e269686e6c90cfc8af83d89c1440ccf1ff75b12f9f79c26f14edceac58e8c0fa"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 3

<a id="canonical-2513d3ab679fe7f7be6767a097ecc3c67666f96b23b04b6362cdb3e05cea2873"></a>

<a id="canonical-148b3e58b7e2a80c59fd9dac305334b190c6339bd0e05264da9e68061df9b4bb"></a>

## name property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 4

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

<a id="canonical-b02f1c9c2f4bb945f89fce4b30d7e7ed34baada65e1266a8e33cefb84f913d73"></a>

<a id="canonical-32c1eab0d7c4ac07f7a923c3d67b82e7081425592b56b9454c963bdee9b05bd4"></a>

## namespace property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 5

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

<a id="canonical-a6b270de5504e430c526914e66bcb4dcbb2b932bd4dd66c612faf910edb03acb"></a>

<a id="canonical-328fe4558b30b4e4d0078ab57b42ccdeb7a0475f0fa22afaea50f5c51547ee0d"></a>

## tenant property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 6

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

<a id="canonical-05ff76b4cbf5c9dd0348bc3de3430ac6fe252d05ad1fc87856be5630d47a76ec"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface_list.interfaces / 6a42c4a573f3 / 7

- [local_control_plane.bgp_config.peers.external.interface_list](resources--voltstack_site--reference--group-009.md#canonical-2f23201d864e821b2dac5be261e8da687697694f4694387e2229c1b732929ab4)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bf787f3a568ad5e358e88f3374f9642ecfbda707330d6a86e313a024a6c89671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab0a284a9909608bccf01782226e2a4c7661b6e2b670e8be0e3c03d2676439ab"></a>

## local_control_plane.bgp_config.peers.external.no_authentication — local_control_plane.bgp_config.peers.external.no_authentication / 1c873e68d3b6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- local_control_plane.bgp_config.peers.external.no_authentication

<a id="canonical-2c0c0f851af7776fc8b5927cf5775261f57d352516bd2cc9b63bfd52a7b35783"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

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
no_authentication = {}
```

<a id="canonical-9db04d303d3d32cc10e15f09792a1c4063a0bb4a2f3f0bd8f305a98d16f714eb"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.no_authentication / 1c873e68d3b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-41ea60d259bd2c652adafb6880fc5669646aea87aa50925fd1b184258a35a629"></a>

## Next pages — local_control_plane.bgp_config.peers.external.no_authentication / 1c873e68d3b6 / 4

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--reference--group-009.md#canonical-5366030978aae805c1eeca11365d95ad08451ed9c81efb5835558dfebf171fc0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f17b01f90720991047a1029f90e2d8403a2b7cad7d3fd44b0c034806f0c266f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a90f96dd89200863dc57a54b50d2eba3f17e0a04db6af77ef32c78751c2a04ab"></a>

## local_control_plane.bgp_config.peers.metadata — local_control_plane.bgp_config.peers.metadata / 9c4eee4ba561 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.metadata

<a id="canonical-b22f820a67380019230a9da1e7d85f06722aed00361184793c9533c52cb5253e"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-95967d060869283f20395630cb77a12e76ce9bba8cedc773bece041030321182"></a>

## Direct properties — local_control_plane.bgp_config.peers.metadata / 9c4eee4ba561 / 3

<a id="canonical-3d28a7fd196812447677905c46b728f8c92a253f95241a983f408f364cf66289"></a>

<a id="canonical-491169b9c6fa4f7b03906b1e568e18c5c54f8a28e6682b49b0dd3b16e1a16598"></a>

## description_spec property — local_control_plane.bgp_config.peers.metadata / 9c4eee4ba561 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-ca67a5ca7d0d592e206381e10ce6c2b62d46c11c7062e16b3e22bd0f2b331a83"></a>

<a id="canonical-a209b4a9f961796c0f500cc7945580c3018fd566237a01215d033eed5a931c1b"></a>

## name property — local_control_plane.bgp_config.peers.metadata / 9c4eee4ba561 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-f378424c1ce9e5da4104e5640f3f479aba0771c2cc045f86c3d03a7440d16ea3"></a>

## Next pages — local_control_plane.bgp_config.peers.metadata / 9c4eee4ba561 / 6

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-1dc593ec427c6d0406f66e5bcdfd5559bc40783be3e95e9ba3798c11cdb9e0d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-039f52f82ddbca4d633a5a6761bebc12b334cc799ef06ae4b1d56b0c7123363e"></a>

## local_control_plane.bgp_config.peers.passive_mode_disabled — local_control_plane.bgp_config.peers.passive_mode_disabled / 9edb384c6e62 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.passive_mode_disabled

<a id="canonical-88fc8ad6f99ab9b02bbf4ef5bae1942f6b0ba26ddbfbf0c39a2f4e589d973f1c"></a>

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
passive_mode_disabled = {}
```

<a id="canonical-892266cc62244ee869906ae4356346e3b2d830c830f5c1cf776e1a61b1b46613"></a>

## Direct properties — local_control_plane.bgp_config.peers.passive_mode_disabled / 9edb384c6e62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a53fc4e8dda3f32bfb4bb7946142ec754aeee1ca6fa687262c4ab8d56bbca6d2"></a>

## Next pages — local_control_plane.bgp_config.peers.passive_mode_disabled / 9edb384c6e62 / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-da39b18bfae933a1126eba6844a5430a7831511873f28be9943b90c9bd5a7b8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e631e8ab6a419bf30ba41e7c4d7b3b382b86803824904e9564721f941047591"></a>

## local_control_plane.bgp_config.peers.passive_mode_enabled — local_control_plane.bgp_config.peers.passive_mode_enabled / 9253b5f02eea / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.passive_mode_enabled

<a id="canonical-34b23d143acca53b3c68170b82d2ac73dd982b04229498e8fd0d7b53be856d50"></a>

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
passive_mode_enabled = {}
```

<a id="canonical-f8e5661fcf84e4b53947886a5c7cd066b37a079ea7ce84637db1abe12fc7d3b6"></a>

## Direct properties — local_control_plane.bgp_config.peers.passive_mode_enabled / 9253b5f02eea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-848775fd06103712788be51650832170a90774aee0db5c1b504317b82b7bc652"></a>

## Next pages — local_control_plane.bgp_config.peers.passive_mode_enabled / 9253b5f02eea / 4

- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b3ec79a2fde1cce623fd97885074907de11e5046440e73a81dcaca95a9a7721"></a>

## local_control_plane.bgp_config.peers.routing_policies — local_control_plane.bgp_config.peers.routing_policies / a9d165c3cc2d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- local_control_plane.bgp_config.peers.routing_policies

<a id="canonical-0665f2cbc6db4bf5eb71a5d48fc7b9c0e5e36985077e367bdd9727cba6a290bf"></a>

Type: `"object"`. single nested block, Optional.

List of rules which can be applied on all or particular nodes.

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
routing_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-923aed0e2de00f67f5c85b6dfd181a46643ec3b1d1ce92f21f180677aaa7d42b"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies / a9d165c3cc2d / 3

- [route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962): complete subsection reference.

<a id="canonical-e7c3a35b76fd811375f0b3c027bf3b5130b6837242590c7e994975f2719b829a"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies / a9d165c3cc2d / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-743a7c8106e5b743db80743a0f037cc07ff7b3f3c2ddd86d490ca5467a3dc6d1"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy — local_control_plane.bgp_config.peers.routing_policies.route_policy / 1e261732e13d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="canonical-335cbcccb5706110bc0010ddadd026aa57d8983ef1518d73f43bcf5d36178741"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("object_refs"),
  validators.ConflictingListObjectAttributes("all_nodes",
    "node_name"),
  validators.ConflictingListObjectAttributes("inbound",
    "outbound")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
route_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-c78175d2d49df1ae88022472b339c1da0ec15217891a5b122d1499c094ff04d4"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy / 1e261732e13d / 3

- [all_nodes](resources--voltstack_site--reference--group-009.md#canonical-bed07a48a3e67551f3634c7230d9ac81c63dfd83468c91efdcc7fca61542841a): complete subsection reference.

- [inbound](resources--voltstack_site--reference--group-009.md#canonical-8f92cb5a9b40e73957e0fa5b0ff17bc75a228ff9d2586270ea32298d7e51d272): complete subsection reference.

- [node_name](resources--voltstack_site--reference--group-009.md#canonical-2bde4a16f8bc6ee1291f50655cc5767f4b4f6a61305b4444f1450bbd375c1000): complete subsection reference.

- [object_refs](resources--voltstack_site--reference--group-009.md#canonical-b79452ca032da7792b0967e479a3f03e51c6999ccf24268fb274cec15bc3a680): complete subsection reference.

- [outbound](resources--voltstack_site--reference--group-009.md#canonical-66e375a2905d36dd0cf9ed3532fda12f5dd4f516c68cf018d535008fbfe011f7): complete subsection reference.

<a id="canonical-4e09b3117b99fc134e9238f1e38ef698738547bcd8c18edd51c5f2cd3a23bcc5"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy / 1e261732e13d / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](resources--voltstack_site--reference--group-009.md#canonical-bed07a48a3e67551f3634c7230d9ac81c63dfd83468c91efdcc7fca61542841a)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](resources--voltstack_site--reference--group-009.md#canonical-8f92cb5a9b40e73957e0fa5b0ff17bc75a228ff9d2586270ea32298d7e51d272)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](resources--voltstack_site--reference--group-009.md#canonical-2bde4a16f8bc6ee1291f50655cc5767f4b4f6a61305b4444f1450bbd375c1000)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](resources--voltstack_site--reference--group-009.md#canonical-b79452ca032da7792b0967e479a3f03e51c6999ccf24268fb274cec15bc3a680)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](resources--voltstack_site--reference--group-009.md#canonical-66e375a2905d36dd0cf9ed3532fda12f5dd4f516c68cf018d535008fbfe011f7)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bed07a48a3e67551f3634c7230d9ac81c63dfd83468c91efdcc7fca61542841a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72babf4d3a000ea5b550a31b85d42169f53d90d0e7af37aa3be77d9e9026e833"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / df73f341f4df / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes

<a id="canonical-83192ec6a36ce1a46068a312ed9e4ec3731ce59cac56074b7d5a3499eccd28d7"></a>

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
all_nodes = {}
```

<a id="canonical-fadcf901c7b00516118dd7695ff497a1c9a838af447f0e575f2d7454a0b96c4d"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / df73f341f4df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6243a9b1761b6d46e31d74fc1fcc0d2a3ca3301696517c9c5f08b06648595b60"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / df73f341f4df / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8f92cb5a9b40e73957e0fa5b0ff17bc75a228ff9d2586270ea32298d7e51d272"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e712e3dbd4d0a6dd33fad6fa0557a751db6fbaf204dac61f84679150d861c82"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / c10b0644afdb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound

<a id="canonical-9e7daa958acc65bfd2a8a5506e4f0658338a6b947e70400fcc60e0dbb65a9a47"></a>

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
inbound = {}
```

<a id="canonical-e26d607f606dd0cbd367bf35c7e18e0a1a2c689212574fe90d502ba3c6e614e7"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / c10b0644afdb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57fbca82d3d183bb51d702a49f53079cccdf8f94b29317adcb33c30df57cb0f5"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / c10b0644afdb / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-2bde4a16f8bc6ee1291f50655cc5767f4b4f6a61305b4444f1450bbd375c1000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ec3885e4b939faeaad282eff6ab438276d17975f846725db62f719af1a540a4"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 04ce09fd0444 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

<a id="canonical-ca21878f49d30142cdd476551d1f286d509b712490df10767f35683c8649c2a3"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

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
node_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c7a2ff08eb5154606345d7e715cddec35bd0fbe4f4a493f9e6518937871f119"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 04ce09fd0444 / 3

<a id="canonical-e22d38aaa11d6911ae5d1c925b64dc09fb4fa00308edde05176d6356ff67a081"></a>

<a id="canonical-7ea40ece13a2f5d6f563119afe51fc69c1e9dfbce37d6f7d4627afbfee38ff24"></a>

## node property — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 04ce09fd0444 / 4

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

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

<a id="canonical-ffc0e70bafd6762a35f4f06b4d7eb2bfabec616971fd41f05ff35ef0520c92e6"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 04ce09fd0444 / 5

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b79452ca032da7792b0967e479a3f03e51c6999ccf24268fb274cec15bc3a680"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e69c018c0132bb75b9155d2113c7969861e22b4bd8fe55b7836eeb433bf64ea"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs

<a id="canonical-aa59af8cea6465ca662aaed841c83af13b68521487edf8c27b5937eee4acf544"></a>

Type: `"object"`. list nested block, Optional.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
object_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-a30d024f5d64b200af20c07281bb860cd41eecdce27e856eaef7fd66b0651a96"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 3

<a id="canonical-5171e2e321086be8606e82072028cf541f3c03c8fcbcab1f2962ff3fd4cdc061"></a>

<a id="canonical-f795c7dc852842a04fe75e2285f79b5f8666c01984264dc432c728bd99f53038"></a>

## kind property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 4

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

<a id="canonical-d1342b86d4998da9f7aadaba748d7e1f39d940d4ba087abdfb440e736ae04656"></a>

<a id="canonical-2be1fcdfddbf83d1624c470b4580c41b3a96175386eeeb2a99f1fca77ed67a4c"></a>

## name property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 5

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

<a id="canonical-f301109e1b1362ad0d48760ace15a6828dd6f03e5cf3b2d83589b0dfaac2694f"></a>

<a id="canonical-a2018d892fbc64547bac923191404a7f4eb8383f95ce7e69e68c8b0e1a62ed01"></a>

## namespace property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 6

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

<a id="canonical-290d50257a496fb695db5bf82857bd51d5ad51535ecac0342fd2af7328d7400c"></a>

<a id="canonical-8b8fc22c222ed3ebea7ab536cf70100c34d5bec7cf9586d9a34cc0f89825c19e"></a>

## tenant property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 7

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

<a id="canonical-52233453aae054705f34f0e65135bad7a550c426686092ff7f421a19505802f3"></a>

<a id="canonical-d631b07c0fb9b901766f33b920ff26aa56a5b9d903e0aa3e5c4ee0e8e3323e5a"></a>

## uid property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 8

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

<a id="canonical-26207105907fea6e98c2a8b273d6b85b71430ef1332a63e5ab68b77ee2c39f95"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 4a6ce04a844a / 9

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-66e375a2905d36dd0cf9ed3532fda12f5dd4f516c68cf018d535008fbfe011f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a35f002f75f5232a9c75ae1e2f672f1ac77c08b54ee7d553fd37a628c174a890"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / 2fdb092eea86 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-32519311fe5ffa5d51f85f730b7c0b55cedeacd4f09e8bd90c1e74875fbffbe1)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-9af925f890c0639bc5e620cd72135623548a3b2c1a4388e96a655de0e685662c)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-dba7757383bb50ac8a1f642e46941db414865c64444eb03a9ddadf3baed99262)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound

<a id="canonical-f9c6b3c3328c09656eec8a65022fab39eb5bb913a2987b6b5ff080efaceef839"></a>

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
outbound = {}
```

<a id="canonical-d11ca8944ffc65ec146e1a99f377e9dcf94c8709981d67bc7ee830b16bcf61ed"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / 2fdb092eea86 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b338f40697bd17e502de205cb55c8b506f56c2d1f5f966ce79b5caf3129c531"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / 2fdb092eea86 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-152c634f54e8e1746baf55be3bb50cfe7cd0278875cfe737101341547c48d962)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-50d928cb80f8ad30ac90e2ade0d1ca6932b970347e988f606ea3f20a3d763bf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eed812724b999a56d96128b85cdb88a2202c597465621d4b3a8927379af84063"></a>

## local_control_plane.inside_vn — local_control_plane.inside_vn / 1155c57c2443 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- local_control_plane.inside_vn

<a id="canonical-c839052e749f435828d97fd2cd9eaad4fc7df7920f27351fd918c741eeeaae8c"></a>

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
inside_vn = {}
```

<a id="canonical-24bdebb49275277f1a50cd1bfa4fb47e19df1135036d8adaa80dbb04f2ebbc31"></a>

## Direct properties — local_control_plane.inside_vn / 1155c57c2443 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9fb98a76e19c9f79534430b0a490d20c83e2af5e8417fd48b33caff7ed9a7d8"></a>

## Next pages — local_control_plane.inside_vn / 1155c57c2443 / 4

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-2d2c4f8d2142811937ab5694e882f60d181ea6e49e98d6828068740bfa122839"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87946144d42dbd2bf94f9a4278d8c90de304726c8cba1d4bf6ba6d16a50d1ef4"></a>

## local_control_plane.outside_vn — local_control_plane.outside_vn / 135b7c2bfc84 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- local_control_plane.outside_vn

<a id="canonical-736f5f3c08b2edd2e4b5bbb79ad321f1f62de6c9bd707473e676ea43349b7957"></a>

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
outside_vn = {}
```

<a id="canonical-41d786da784df280b4d8ece8530c6deb64903f523fea3531bc0d0ceec2e8f7dd"></a>

## Direct properties — local_control_plane.outside_vn / 135b7c2bfc84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ca8454a4543e7930b181754c5d9dab26e17ce2f65df58e4adb5f9adeb09a054"></a>

## Next pages — local_control_plane.outside_vn / 135b7c2bfc84 / 4

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f427be209dfb9b1b092fec98b6fc07c003c09f3ccd6188933e00188986bba470"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c7d8406fc2c257e54e1bae2f85c5c218cf9241071a8dfb22686997ebc9fa0d5"></a>

## log_receiver — log_receiver / 5299b0c8843d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- log_receiver

<a id="canonical-4085a3fb200cb27243b3a79f489362fb84c15b0c5a67499f76252129914aae56"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](resources--voltstack_site--reference--group-009.md#canonical-4085a3fb200cb27243b3a79f489362fb84c15b0c5a67499f76252129914aae56)
- [logs_streaming_disabled](resources--voltstack_site--reference--group-009.md#canonical-55499b6655fb5685eaca9e53c91604d83c8985ac79fafa8b153e1356caed02c2)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac5f78729302ee66fa010ecbc593334149a0950ce64256e2553015fc33286d26"></a>

## Direct properties — log_receiver / 5299b0c8843d / 3

<a id="canonical-118ced4c49a5d7a186d206b4bf49346c58b2ded4dfda47906ca47480d7f89c35"></a>

<a id="canonical-1ac896a70e4272ce2522d821de9033bb5085bbe142e6d424513031726ff47aea"></a>

## name property — log_receiver / 5299b0c8843d / 4

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

<a id="canonical-6e44157ea0ccd32e9ebd98135d683dd13bad8b4855d47b83f5a6e29d35a01ab8"></a>

<a id="canonical-d9712071ba16a7fd16d853d1a4e23b12cef4105c6690b03a214078c5b214c4f7"></a>

## namespace property — log_receiver / 5299b0c8843d / 5

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

<a id="canonical-323a8d47a8e85b3f350ca638f0486853e9679cb68566e1f1058879bf95ed2855"></a>

<a id="canonical-5a9bfd47ada45746604372aba4359f48008a7fb6274c07b2d1af34cf687e3e62"></a>

## tenant property — log_receiver / 5299b0c8843d / 6

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

<a id="canonical-d8c066fcce3ca3ae1316e0446b93fd9b2df5854804f1531ddeec3e665e0c1af2"></a>

## Next pages — log_receiver / 5299b0c8843d / 7

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-870e433337c76cd152ff0fdeb29b4d50767116fb7e6727361a766d0987cd9cee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-065a48bee67b75a974f7e5789d903542967e4f0511d171689e86f3fbd7206f1e"></a>

## logs_streaming_disabled — logs_streaming_disabled / 886a14bab3d0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- logs_streaming_disabled

<a id="canonical-55499b6655fb5685eaca9e53c91604d83c8985ac79fafa8b153e1356caed02c2"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-9698338f675022b97783de69e512e60aecf82bb6c552ca3ccd7f84838bd47899"></a>

## Direct properties — logs_streaming_disabled / 886a14bab3d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6cb23dcd68c4879ab877d0f39ec4c33d9597577fc2ab1c41e8a04afad74b3d4"></a>

## Next pages — logs_streaming_disabled / 886a14bab3d0 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ad1c28b9743413718a3b177c526c28dd11e9734d12ec0a1899837c724ececc4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c103803a3c2c62b478017854be6a11e2729c2f755541c889613351fc29295382"></a>

## master_node_configuration — master_node_configuration / 5907e745e94e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- master_node_configuration

<a id="canonical-cb35759b1a8c969d5e610c66d66f847924e0e9cb958320d20eb3c9e1d3b0c980"></a>

Type: `"object"`. list nested block, Optional.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  }
}
```

Terraform syntax:

```terraform
master_node_configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f67f012f61489265f2c59fb14247b18577e23165a2f896a7acd73a1b81d8a61"></a>

## Direct properties — master_node_configuration / 5907e745e94e / 3

<a id="canonical-5de30f1434c8843a7b63fb56e0c4d150253d3a7d3a78585bbd2b770d771a53cd"></a>

<a id="canonical-0ab6203e1b23c6cdb6e58c6737fc8ed75a3d793257ffd6307e7f68a4f01a55a5"></a>

## name property — master_node_configuration / 5907e745e94e / 4

Type: `"string"`. Optional.

Name. Names of master node.

Upstream description:

Names of master node.

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

<a id="canonical-2ebd275dc4f13b204cc0d92610b5f4c39f5dd7539dde03779ea5c492665e7219"></a>

<a id="canonical-9526d9a28c3ed04d5002c2925e1c8666c71c9ce676a7ec3b4605a1fb19e5aefa"></a>

## public_ip property — master_node_configuration / 5907e745e94e / 5

Type: `"string"`. Optional.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-13c0d75ba858b52ba781e7c4c46e1bfc29a41d9339999d0b19c961c93a738011"></a>

## Next pages — master_node_configuration / 5907e745e94e / 6

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-15682f23e72c547d6211938bb42701baedc08ac036e147a1efa0950d69f09bbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-026c5d43db1b0a9e0038979d6484db9bc46805280a9a4b1b608869fcb0e3c5ae"></a>

## no_bond_devices — no_bond_devices / 131cdfc60c9a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- no_bond_devices

<a id="canonical-840ebc49693dfa072b1c47b97804191267324348e01de42e27b923a9bd2bcec4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no bond devices.

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
no_bond_devices = {}
```

<a id="canonical-dac928505671bb2ad38c877e578096645b2257bfd752263334fb6cdc86c87d9a"></a>

## Direct properties — no_bond_devices / 131cdfc60c9a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40f46a3229629d243b46b2b46a20130e72982eeefeb6ae2aa169e31dcf6922d2"></a>

## Next pages — no_bond_devices / 131cdfc60c9a / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-54d6d51504ff10e4b354b972db67901b7e303b239c786fda5988906e6b72c9eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c91b8d0480b49dd0d963e32bee867e652e5d41aab7ad9e7d6a842385c3614bc1"></a>

## no_k8s_cluster — no_k8s_cluster / 57f0f07efec0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- no_k8s_cluster

<a id="canonical-6e7f968f315aed7150d071f19751134cfd3baf18e1132fbae8a5eedb20b43c85"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-fa705810fcf87a02d35bff6dcc24d9e962cdf349ed4e6b115b75b634c79cdb82"></a>

## Direct properties — no_k8s_cluster / 57f0f07efec0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-077535234bd48b06e9af2614019276d597788ff81cc5da376999c9327a0279ef"></a>

## Next pages — no_k8s_cluster / 57f0f07efec0 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-6671aef9498eb983a7034a3bc3ed9724f6fe735664ca5250753bef797a70f967"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e34fde3021aaca0b603ce353b128f5f4ddaf2917656dfb9ff770f78d3e608a9c"></a>

## no_local_control_plane — no_local_control_plane / d05a1e349be7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- no_local_control_plane

<a id="canonical-2d1254303e2c0c57b6d3c235b3025a4f6eeba159f11dfac8ff74589854cd13f0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no local control plane.

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
no_local_control_plane = {}
```

<a id="canonical-b905021f6e82ed16f820efbccf39d3b0605f51d32290390992c48e6a1afa0973"></a>

## Direct properties — no_local_control_plane / d05a1e349be7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f53c50eb813897e65a0d6722db92457684f1e43c2014faa364578fd91b29315"></a>

## Next pages — no_local_control_plane / d05a1e349be7 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a46a5bf147881ebfb864776d173dc1b9f457fcac10bf86028b80661d8fd8cfa2"></a>

## offline_survivability_mode — offline_survivability_mode / 5a3b2f468caf / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- offline_survivability_mode

<a id="canonical-73da32133a89af9e4faa038381041d159cbb71a4c879777e4f58cb8b21c9fc2e"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-e70eb83674fb025acf04dbf873517264aa30a6abf8358da34b49424aa2d84a56"></a>

## Direct properties — offline_survivability_mode / 5a3b2f468caf / 3

- [enable_offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-bc68728488f83bd04bdd2e492e0f4b20fc715d093d930cdf0e25f6994798ab42): complete subsection reference.

- [no_offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-a48450ac672f696afc7b9a77ab6013ee7b6e0f710cf6629bfa6cbc620602506f): complete subsection reference.

<a id="canonical-2ba141b866e17a51db1fd11a13b63040843f7c2cc500a0f0860e34657e3f478c"></a>

## Next pages — offline_survivability_mode / 5a3b2f468caf / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-bc68728488f83bd04bdd2e492e0f4b20fc715d093d930cdf0e25f6994798ab42)
- [offline_survivability_mode.no_offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-a48450ac672f696afc7b9a77ab6013ee7b6e0f710cf6629bfa6cbc620602506f)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bc68728488f83bd04bdd2e492e0f4b20fc715d093d930cdf0e25f6994798ab42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9be06fe2f7a8061ef235609cadcf212da845178f300ac40ea61b10eed3f13155"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / 3a3a5b515462 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-9d2335d5ad2186a03ad3a61d4f5c6908218193518f83cc9cc59d8419738433a5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-96be7e2bd5923bd4aedffe1cc3988ccb7b6f4e33d3d71dd2f13390e3d932bbe6"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / 3a3a5b515462 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba2d49a6a1698754be6376c2e4fab13b170955880b0b85cacbf9be736cb2e5ef"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / 3a3a5b515462 / 4

- [offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-a48450ac672f696afc7b9a77ab6013ee7b6e0f710cf6629bfa6cbc620602506f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
