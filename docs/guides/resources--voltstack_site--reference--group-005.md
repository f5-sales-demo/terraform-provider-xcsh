---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-9640ee246b162d6fc57ba4a4ce94250d8498f1af2ffbb0fc6d2f77e2d7b59360"></a>

## tenant property — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / c44c09d54213 / 6

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

<a id="canonical-9ce84fb3aba4e5cb93843f5eb554fed167694d07bf132332d4fd3bd7acde63d3"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / c44c09d54213 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-46eac30bc0c70c9de3a7ec7a931822791531ed76f7bf3fa2bb319ff4d046f011)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-98bc5a3ff6af9ad3923d6256482ee1a24488a99197c1ef373363351e35883a30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce14360d1deee7601b22421afe6117c4431d4fb4f6611e0bd2ee5a1373170513"></a>

## custom_network_config.no_forward_proxy — custom_network_config.no_forward_proxy / c1b5606b03d1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.no_forward_proxy

<a id="canonical-23383989bc68272f151feca2b6761dc033a3afc86a41d524d217d1a3c6ba2c48"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-724849809ba9daf2fda79774fa11d13d5477cf1f9ea529b82bdc98e292c47433"></a>

## Direct properties — custom_network_config.no_forward_proxy / c1b5606b03d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7bf4328855f86b8726db2198534d69686a253efaf50de288f68ef00eb8ea753"></a>

## Next pages — custom_network_config.no_forward_proxy / c1b5606b03d1 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-d4f86e83fa178d4f00eff02424ec5da13c6bdb3bc807f2f5912809d2e67b0359"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae3e8011f3e39d4c4131e7292db15bbf9bd2294db4d7e9eca515e8261701a0ce"></a>

## custom_network_config.no_global_network — custom_network_config.no_global_network / cedf8be2ca6f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.no_global_network

<a id="canonical-e489ffdf14a21bdcc3e7b3fbb866afb74564957793f187dced08d87c470a2cab"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-5c47cbbe4a6e8587ac7dfef9eb918374c965905013c9c4caca87c255efa06c6a"></a>

## Direct properties — custom_network_config.no_global_network / cedf8be2ca6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38a9eb89ff8abf6c4cbd3f80c8b705b4be30fb0ec59bcabb7c9e32c2d4858355"></a>

## Next pages — custom_network_config.no_global_network / cedf8be2ca6f / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-4d403ca6e8c9501d6b8716a1318e34ab9946407fcc2cbe306da83168aff71e27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-747928fc2692a4d6c6135be22b33f46766c1f3e5099fc7ee0af107b32c00350e"></a>

## custom_network_config.no_network_policy — custom_network_config.no_network_policy / 136b50a1aff3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.no_network_policy

<a id="canonical-603c47e654045ed57731097feade81990e26055ad56eca194e63b78f87ac5a0b"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-c8d104746b04d9b009bdf6cce91e32c328da6361abc3b241baa51dea6ffe6810"></a>

## Direct properties — custom_network_config.no_network_policy / 136b50a1aff3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f25d5ac68fd8ea137615977608e53eb152e7e9b14dc393797a28c1040a319f1"></a>

## Next pages — custom_network_config.no_network_policy / 136b50a1aff3 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6b28df73fcff01db85eded04b82d1b8addbcfe272c05d006e21fc022b53af5b"></a>

## custom_network_config.sli_config — custom_network_config.sli_config / 6f88683bce35 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.sli_config

<a id="canonical-a4ff480e6285fcfea431f548de1a201baf3bdbd37c4a88a26db09dfac80b5268"></a>

Type: `"object"`. single nested block, Optional.

Site local inside network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
sli_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0954ac2225d2c5f827d9cb3a6086762b9739e26b4a9b5958424185e6d212b74f"></a>

## Direct properties — custom_network_config.sli_config / 6f88683bce35 / 3

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-809abf6dd6d4973c108cabf0893fd364206229d34daa017ec7338da2b7539ea5): complete subsection reference.

- [no_v6_static_routes](resources--voltstack_site--reference--group-005.md#canonical-a6d1f7ecef13d5a9bddb10d41973aa9f487cdcf06f482f40b59f4d45dc88046c): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1): complete subsection reference.

- [static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623): complete subsection reference.

<a id="canonical-3106e18b9b4b621eb108e18f77f2f092241e220344eded9a628c6333894f676d"></a>

## Next pages — custom_network_config.sli_config / 6f88683bce35 / 4

- [custom_network_config.sli_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-809abf6dd6d4973c108cabf0893fd364206229d34daa017ec7338da2b7539ea5)
- [custom_network_config.sli_config.no_v6_static_routes](resources--voltstack_site--reference--group-005.md#canonical-a6d1f7ecef13d5a9bddb10d41973aa9f487cdcf06f482f40b59f4d45dc88046c)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-809abf6dd6d4973c108cabf0893fd364206229d34daa017ec7338da2b7539ea5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a98c4514e70e55c8d757ffd0465c7a82c477f0239ffbb271529510d07c1a5608"></a>

## custom_network_config.sli_config.no_static_routes — custom_network_config.sli_config.no_static_routes / 92a62d36b7ae / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-624a76dd627213e823383837fb4d755ae6e37a0662ae59a12174c6dfc32f04cd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-dabd94d8d414bb0e252766eb731808cf08f76b990e228c269d78b743390edbb1"></a>

## Direct properties — custom_network_config.sli_config.no_static_routes / 92a62d36b7ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7dfba7c4b5437860a0dddaa4caac28a6826550f5431f50ca1a5ce84fce61a34"></a>

## Next pages — custom_network_config.sli_config.no_static_routes / 92a62d36b7ae / 4

- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-a6d1f7ecef13d5a9bddb10d41973aa9f487cdcf06f482f40b59f4d45dc88046c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecc4e8008148fde457d814c77b7599dbabef585d8dfe95f8cdbac19bbb12575a"></a>

## custom_network_config.sli_config.no_v6_static_routes — custom_network_config.sli_config.no_v6_static_routes / b69d9ea55a2d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-20ccfb7ad48642d1e5eafddd3f11bd2716e78499470c8bcf4dd71a52d0dcb9b3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

<a id="canonical-9f7cf62788570670eea7b3e16f63c33a203241db14dde513832b8a95d6cce0d6"></a>

## Direct properties — custom_network_config.sli_config.no_v6_static_routes / b69d9ea55a2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16bfc318e340d41dd14d0144b1da6acb8545555187ab7f7fe1c1e7506299fd7c"></a>

## Next pages — custom_network_config.sli_config.no_v6_static_routes / b69d9ea55a2d / 4

- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed7cb6f60e60267fa59aa393f8300992a60f29e706736f54a2c554c187098a69"></a>

## custom_network_config.sli_config.static_routes — custom_network_config.sli_config.static_routes / 4d46a773ab08 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- custom_network_config.sli_config.static_routes

<a id="canonical-ed88f51e89ac89ce5fcdb15aff715b05c4700420676c649e0cf9bf6e1480680d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a9467ef0cb19e1b4c15a97ede78d9d23cd7fd61db2d7c6a0e2a1e4e98bcbbe9"></a>

## Direct properties — custom_network_config.sli_config.static_routes / 4d46a773ab08 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51): complete subsection reference.

<a id="canonical-fb3ee985fb9db0de957ad262f4b44781c179e31e7e12f0adcd85705230c9b0ea"></a>

## Next pages — custom_network_config.sli_config.static_routes / 4d46a773ab08 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c50665f57e6a47ec9ddb4da59f92eafbaea31b85776472fb2caacdd3579c36a1"></a>

## custom_network_config.sli_config.static_routes.static_routes — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-d9e5272daf05c698aff07b246b99ed89729dab067adbdf807f1cc733563f1290"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-957c158104d76f8f096478a70f67fce85fb0728b3d49616fc119dc9f1480470c"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 3

<a id="canonical-7203dbfa0800e57be3d68239197c5c466549f70838d69ba4eb64e5adb0e713f3"></a>

<a id="canonical-848dc011eda8e9d83855e3138c4b5366f0ba48b43c1f48a0b2a7f68e4108e32d"></a>

## attrs property — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-197cea805d8ca80b849f2c28923d3996e24846c24323050afc00836e0d746af2): complete subsection reference.

<a id="canonical-a0e2a00c9613edc700d7c7ee73ec4d61b2fcdebf3d38ead2747ae357732a3d35"></a>

<a id="canonical-d2d606d4e9ee2f457909bd3e318dc8066be1a2a4020bf75ca737e89c214424a3"></a>

## ip_address property — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-f8d77fbd9d49128b8e4d64df0536361a67c52dc9624f075bb03209c7b6e27386"></a>

<a id="canonical-141e4dbef0473b85ecefb392fc94c6aeca2d09edc8aa08942d0afe70a93bbcc2"></a>

## ip_prefixes property — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca): complete subsection reference.

<a id="canonical-3912ff759f72e4aa234c68cbe19d1d4812317ec2634cd0d41a7d4ebbe31d6702"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes / 633005f0a4ff / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-197cea805d8ca80b849f2c28923d3996e24846c24323050afc00836e0d746af2)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-197cea805d8ca80b849f2c28923d3996e24846c24323050afc00836e0d746af2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a62ab617f6982a669385c667494cfa5838330e598d980e3d363c3e609bcf44de"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — custom_network_config.sli_config.static_routes.static_routes.default_gateway / a6debfcf6245 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-89e659466a1756dbe7b6a78391fc8c5674d6ac12fe7a955966cb20b0f4453f21"></a>

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

<a id="canonical-99fa53baa22d09fac3371c54a8bb7e9d6b4a38accc7715a2e6a8dc06d38ec31e"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.default_gateway / a6debfcf6245 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-283890d60ceb2fcd68f918a57ad3ef1ddc3e61cb3cc9b861e58d5ed8681f3edd"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.default_gateway / a6debfcf6245 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-407b403341820e92937480484ea4f5681443547e6602e13d54d268c908f6e9d9"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — custom_network_config.sli_config.static_routes.static_routes.node_interface / 9405d8748819 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-401c8e80562a0b2c694f54497fbd715607a952a06442f380fe32d8a3c9f374ce"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2a236978d3af1510c80dce258935e4ecd59bbf038afa0d66d831713af03ee90c"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface / 9405d8748819 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-6e093de46046983c3c430cd59412f5673746856a633bed74e9489049c22e1555): complete subsection reference.

<a id="canonical-b37440d4b309abf9b3b2c93afd7436ae6ace424275d22185305dde7c270b31e9"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface / 9405d8748819 / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-6e093de46046983c3c430cd59412f5673746856a633bed74e9489049c22e1555)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-6e093de46046983c3c430cd59412f5673746856a633bed74e9489049c22e1555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4eaad0845373ccea08ef9312e89c72898cee777e7927b1a0bb56dd073c3a21eb"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 706b57e22dd1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-0d47c7cef8251bf78c14e191215b0b6a1030f3e277bb6dc2316dfda42dac946e"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a76a6b1a1b3aee5d7f352e297d0240d7dd42baf430d9c996f71fbfb2f7aafe07"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 706b57e22dd1 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-340a57fbc0e14f2b2a9a126045853ab78dfbf144f5a3f84d4b5c709891d5f0a1): complete subsection reference.

<a id="canonical-2a4bf94abc2cd38a2722a15bf92b1faf62e88c9ea17c1f25be63452845fd698a"></a>

<a id="canonical-61cdb36ea92d8807bf5ef637a5c7d615a78198451059dba4fc37d1fa6226aecc"></a>

## node property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 706b57e22dd1 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-f4441f0b80f241cc1fac1ad4171dac4b9e29ed1263c946d70bfee23980296b6f"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 706b57e22dd1 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-340a57fbc0e14f2b2a9a126045853ab78dfbf144f5a3f84d4b5c709891d5f0a1)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-340a57fbc0e14f2b2a9a126045853ab78dfbf144f5a3f84d4b5c709891d5f0a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec41383424ef432c47e1329fa614ea7c79c0fc8e820a19b0bde3bca8e499f0bd"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-c48797194375e292796f076575a8890764a6d851e5284e17a29c8a3fee7bdbe1)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-ddc9aed50af99cda4d259956d7cc86eaa614b423c4a07edd12e2920e16c29e51)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-bd7b21abb04917793d1581503f7eaf0aecddb936eb6a8b4f41bb4782b562adca)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-6e093de46046983c3c430cd59412f5673746856a633bed74e9489049c22e1555)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-ebaf595e7d54243f089e7e0b9111a84989eb56d97f4aabc9ab993d7032e7c485"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-68a3eedd1cf3e3f281bd0ee39ab5ba110384a82fe2acb453875557d847f82925"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 3

<a id="canonical-2e0dd958018067fb3385be0371d5bd3193db9d2256476ee390a1abbc4addb782"></a>

<a id="canonical-78a6b85b535773a8e90e2eacdf0209183500fdd608c51455d905b0b20d549793"></a>

## kind property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 4

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

<a id="canonical-987248bad02b84bc7ddaa2fb8f4113105c960cbee45885fd21104d44af5ef244"></a>

<a id="canonical-6d278e47676386f89f2df26c8a213856b16dacd67a825cb6842091d6dad1d637"></a>

## name property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 5

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

<a id="canonical-7d51b0c93ed725fcd539ee242368808dde884b0453583d9ba6f73106f506b4c6"></a>

<a id="canonical-ba71aff7b1c904cc3e7e8c137858709c10f4a8fb9e6c2d216594407d3fd06f86"></a>

## namespace property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 6

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

<a id="canonical-2d1e01c3ad6c43bc410d75c39d99d4a08a3430a8e968c790c0eb214981bd1ae9"></a>

<a id="canonical-427d8822d63725ed2c2f144ea1ae689b52d96722f784e8a8979cfc429af1525e"></a>

## tenant property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 7

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

<a id="canonical-dbe15467d27c234e62ae6cb3a10d8e30c9284af7b492d6b35aa9f0af47ada888"></a>

<a id="canonical-4271274bec1ba7d362b8f6bb1e8c62b7ae1266af0eb8519a2b75387d36d2177b"></a>

## uid property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 8

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

<a id="canonical-cc25f48b7eeffc633fbccf1f4289616be2c05ba89371dbf0678bbf7d630485c8"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / 36b58bde88cb / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-6e093de46046983c3c430cd59412f5673746856a633bed74e9489049c22e1555)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbae371d777d429c1fb846d3f1109b227f6d8884f5a34c139a1d3ee5fb5a38de"></a>

## custom_network_config.sli_config.static_v6_routes — custom_network_config.sli_config.static_v6_routes / 7aa96c2a37ae / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-5ca0d3033e0c7f2974959859e9a94b94b6b310ae6f142338303eea29aed2cdd8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-b88f3a7aea0783b35cec994c5479324309464b9d06bee4abcf9d0c1dbd500a6d"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes / 7aa96c2a37ae / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06): complete subsection reference.

<a id="canonical-6d0738bb51f00b88b4fa7a28086a137382df6410c398131b90864a465d5d6586"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes / 7aa96c2a37ae / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93e9123bc2e9dbc3c042dcc5b8a11b234d3d444737210bb0c9f92ee42729c3fb"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-55a139aacf6a5c34a5e59d5d0568f6e29a9123549b4553f2c1e4494929c3a17f"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a7cd6625041237e139393e3368b0a73b86fda5f1332e2317675c6b551f1ea39"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 3

<a id="canonical-1fc7baa8cee1694457922f1b3e1fccd02c9287c00469b0af6c38bfaa18e8be92"></a>

<a id="canonical-5f20d89e91bc1b008cd9b71156857ff92b15dc9409f205b5f45ba50dca7abfa0"></a>

## attrs property — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-08aa96c218ea720f000217a4448531bfd662ddc1107b526e9dd66d6cbb61eaef): complete subsection reference.

<a id="canonical-e3e00fbfe8e77b99cd312bf760514f0669ac36d8aa02347799e5dcf9eb602bb0"></a>

<a id="canonical-244957664dfc49b00b8583464e3cc226e95785143c5ef860fe2cde7e4b825faf"></a>

## ip_address property — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-83f1285c7e8a7b13b6aaf54923422f4a7cb2e2f397c6bac61f4fc25bf5d91372"></a>

<a id="canonical-daefff0809cd59026f58e0146a92ac3778827759fe7f70f8ff1027a71366dcf2"></a>

## ip_prefixes property — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e): complete subsection reference.

<a id="canonical-1c9c45a8c65ae2adf89b647835c935bddcae4d1d8939109b75c56ee41255f85e"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes / 4d2cb21279ef / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-08aa96c218ea720f000217a4448531bfd662ddc1107b526e9dd66d6cbb61eaef)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-08aa96c218ea720f000217a4448531bfd662ddc1107b526e9dd66d6cbb61eaef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51c6839e243e714ba1ac820609203c9f398f61332f0c07adbce5d5786dd82de2"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 90148c0448ba / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-f6d904daf3c995b0a2809fddd65f31571cfa230ba1bb667357f345c7c51fd7cf"></a>

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

<a id="canonical-5a78957d414e2ea04a3cc0b21fa82d299d9c03f0878170c74d3752f6b933b0b8"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 90148c0448ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea0b34d6d8bee9199ab430666ec760d5a4e5fda109c85c43727c267cfe6c22ff"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 90148c0448ba / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-232cf10bacf7a9c93a5e599abbca54ac161635556129232e1abc10c78a0b3544"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / ff44c67381a8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-253fc8676373f1cd483f1f79f6cfb69194050167207e52c6b061d6ae87b28f62"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f279596d6825f93980b9bbbe0b8939216b3dcc0897ccd1c3152fd4c95413587"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / ff44c67381a8 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-61e90ba6be956618cd709a44a2e41cbb27165ad4362580b696a2927a87e09901): complete subsection reference.

<a id="canonical-d01ae33902f394f2a9dd5860e220f4df788ba3719da1749b5ca41bcc028ad596"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / ff44c67381a8 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-61e90ba6be956618cd709a44a2e41cbb27165ad4362580b696a2927a87e09901)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-61e90ba6be956618cd709a44a2e41cbb27165ad4362580b696a2927a87e09901"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-881a0b1c5867ed6da7987ec4f68796d592871e57b7485909ab1e6141fe765644"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / c168f15052ff / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-9028c21771d72d2061aba05a22b816e02395da0a163d88cbc96cd6ef5a54ee7b"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-331d9ca779786a85b45ef1aee02a4496e66e7dc89111e3f5400a050bcc300582"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / c168f15052ff / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-491b30d81de3eba2a81822215f4f280adcb06fa599548f8f581d32eeb0260160): complete subsection reference.

<a id="canonical-98bf61651493ff6a53d35bb914f35fa531b0d5928060a97c8ec2632f7e6b6b7f"></a>

<a id="canonical-df67ce5f2afdb31898a8c4d17636430b847b53f062873176facc93856eb38ee6"></a>

## node property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / c168f15052ff / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-d3d81a8b5afdb247d68e59489495ee35fc4b429303d42261b707db8969b89de7"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / c168f15052ff / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-491b30d81de3eba2a81822215f4f280adcb06fa599548f8f581d32eeb0260160)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-491b30d81de3eba2a81822215f4f280adcb06fa599548f8f581d32eeb0260160"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fb003244657d07117dc4135cc00ad09f475c8438f7122047100593bdb35020b"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-52c4bcf944cffb9290503355074177e6a9b567b4e6b7d7f4750c38f9d2544fb7)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f3891df5be4e55ad6b85231b7a431ff2bf9509ecaa69e32535a50edda05bf623)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-08fc4eead66092017102f59b51f7c7c83e0819d44c6eb0470fb6f14d89d43c06)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0924c56fe5ca6fec65642cd5877d7a3cf9c4cbd42f2343c935ff76eefac2a76e)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-61e90ba6be956618cd709a44a2e41cbb27165ad4362580b696a2927a87e09901)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-5d972b0888cf33b96b67777a13fce0df7f6ab9889e3392344c0ed8eb3e612e50"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-61deb5719c8390cfc1e0be640fa905363bdc57e30221129a46ad6fff72d67c8a"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 3

<a id="canonical-c5988d5b42239fb301a865be7201871997e68cdbb075c3993537fab9c82ac122"></a>

<a id="canonical-d09b25eba8701c4910a0b881bc52a9a21b8df624d763d1933bb13af42e2269f9"></a>

## kind property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 4

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

<a id="canonical-6b6606b48084751c197b178555179935e0d3338ba841644ef19dc1735b6af65a"></a>

<a id="canonical-3febaf2aabd021317fb3b30ca69b693bde36acc067811394b1617e5701c9034b"></a>

## name property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 5

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

<a id="canonical-7b2f599c57fa61657ebc8b4766015d5cdab9386a592f47bfced75126cf759daf"></a>

<a id="canonical-2f9b09c743ee1f1176cda98812f6417316a3119a5f6aab38e73e8079492098e2"></a>

## namespace property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 6

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

<a id="canonical-80b993ffdd7caa8d787a4df901c9d69855d5e764e996b6aed04db5a4514e73bf"></a>

<a id="canonical-b8811e61814442ceee0b361b38a21d7e356ed16913dab81499ffc9e3015983ca"></a>

## tenant property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 7

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

<a id="canonical-196bc73c8c5219e583e25662629d1c1bea99cf3cbca76d2acc6f2c30f335210a"></a>

<a id="canonical-51c7f07f1b6c80a318598eb6e9414608111a28b6165e80687806bfd3569d9ab0"></a>

## uid property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 8

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

<a id="canonical-7f4a878c18e15ec1547786c614aee2a37aea57e72bab967f2f085a0363994a43"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 64cc415d097f / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-61e90ba6be956618cd709a44a2e41cbb27165ad4362580b696a2927a87e09901)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d20755103000d21da1d27d84d5e90bcf18eb6c135278413bd8a64a23b22b4ef5"></a>

## custom_network_config.slo_config — custom_network_config.slo_config / 77643edbe768 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.slo_config

<a id="canonical-28f288b7be7c33770b55279a21e76f681398df2b58546a07fb6ea88678d9a490"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_static_v6_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3680226c3b7cf86453edca8463bd9d4783df529be7e4e4dc58301cf7f288f97"></a>

## Direct properties — custom_network_config.slo_config / 77643edbe768 / 3

- [dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-9dcb72b3a7654aff237e5950b5ea167656fcc55801c66ea988535d59dc23c132): complete subsection reference.

- [labels](resources--voltstack_site--reference--group-005.md#canonical-c45772491b62780d0b2f3eb3540e77d411908e2b17ba9c2a98f3a051b792e779): complete subsection reference.

- [no_dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-dc815ffc33530b0f4722045a09b8a35a909e4d7e0e91fcb4b2e1114a0dfdb2c5): complete subsection reference.

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-1c9d65435f57551a4b6394ddbae032c9a74bfaecc7b79c7dd79d625af23c08c6): complete subsection reference.

- [no_static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f400355fb47be5b7f3f36bc266ec32ddb48dc98a9e5a29a5279d3c264725ca79): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8): complete subsection reference.

- [static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db): complete subsection reference.

<a id="canonical-eeb85e901c4a8764dc5509d3fd01fd54d1910b088e302292dbcc31a8ee5007f8"></a>

## Next pages — custom_network_config.slo_config / 77643edbe768 / 4

- [custom_network_config.slo_config.dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-9dcb72b3a7654aff237e5950b5ea167656fcc55801c66ea988535d59dc23c132)
- [custom_network_config.slo_config.labels](resources--voltstack_site--reference--group-005.md#canonical-c45772491b62780d0b2f3eb3540e77d411908e2b17ba9c2a98f3a051b792e779)
- [custom_network_config.slo_config.no_dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-dc815ffc33530b0f4722045a09b8a35a909e4d7e0e91fcb4b2e1114a0dfdb2c5)
- [custom_network_config.slo_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-1c9d65435f57551a4b6394ddbae032c9a74bfaecc7b79c7dd79d625af23c08c6)
- [custom_network_config.slo_config.no_static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-f400355fb47be5b7f3f36bc266ec32ddb48dc98a9e5a29a5279d3c264725ca79)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9dcb72b3a7654aff237e5950b5ea167656fcc55801c66ea988535d59dc23c132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d7b6543f734476fc85b5ec695239b4b54faa0913de8d30cae19fd29f995fd7a"></a>

## custom_network_config.slo_config.dc_cluster_group — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-79af8ea5e432faadb2721b1a8192b432f7d4c7533c440d4e861280810bd58cf5"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-13804f29765ce40a689c4b33641e11c2552366af838fe86dabc73f7d2f6ccfab"></a>

## Direct properties — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 3

<a id="canonical-ff2276f8666cf716806f1cd69dd347d1428437544861159bd8185cff7c9fd37d"></a>

<a id="canonical-8bd2d43aa5c14db5c7c02899df40334ac77c0f9ccf6f0868aa31696d51f4bf5b"></a>

## name property — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 4

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

<a id="canonical-602823d08b88f0d7f21f1361c0a7370a3b29b0033947cc68227c93a45290439d"></a>

<a id="canonical-5e0f7707720749c0f434ae634478db2c891598d6641cc89456b5f5841b4c86e6"></a>

## namespace property — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 5

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

<a id="canonical-1628020c831b96fa825f4a79907ee5ff93601cd9f168e63a08fb36d566639ac2"></a>

<a id="canonical-b67645c5ae7ee70a8cb83782f2cab336e6df2bc5a3c9831d414c950fc348737a"></a>

## tenant property — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 6

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

<a id="canonical-f8db9942ddb22193abf16b9694542bda148be20484939bdd7a226a553386dd6e"></a>

## Next pages — custom_network_config.slo_config.dc_cluster_group / eb00243eb4f8 / 7

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c45772491b62780d0b2f3eb3540e77d411908e2b17ba9c2a98f3a051b792e779"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a33668223d30bd79d787c6d0c008ab9c109b1d1d69860708b238cd7fc8847222"></a>

## custom_network_config.slo_config.labels — custom_network_config.slo_config.labels / ade3c53a3450 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.labels

<a id="canonical-6065dee1647bc9e840f7b4a8f5aad74a1cde1f8b336db392d205c0cb3bd9d0d0"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this network, these labels can be used in firewall policy.

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
labels {}
```

<a id="canonical-1972b9f74582ca75e959c8d7841a74e4c2b16bc82f0f17ad957e65fd0621686b"></a>

## Direct properties — custom_network_config.slo_config.labels / ade3c53a3450 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-effbc6c5ba6ea563615d8939526c89f4b97ffe0c31787811a99d4b5dbfdd8a7b"></a>

## Next pages — custom_network_config.slo_config.labels / ade3c53a3450 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-dc815ffc33530b0f4722045a09b8a35a909e4d7e0e91fcb4b2e1114a0dfdb2c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f916e24f664ddeac2ea3f82afa12e4c57c2ce19b45a0adc136c4b5e299d923a"></a>

## custom_network_config.slo_config.no_dc_cluster_group — custom_network_config.slo_config.no_dc_cluster_group / 3168d20b593d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-c29935fa820d1e5a64f4c5b303b820ab5424cf72e1362fc13ea7a63ee476ba3e"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-9a66f6e2f6f1fe41c95773262206945b5d78fac562f48adaca4bef9264b78987"></a>

## Direct properties — custom_network_config.slo_config.no_dc_cluster_group / 3168d20b593d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07f1b9934bd92d6f26337bcd087a15a1f57733b0bcade15bdeca6125856e2999"></a>

## Next pages — custom_network_config.slo_config.no_dc_cluster_group / 3168d20b593d / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-1c9d65435f57551a4b6394ddbae032c9a74bfaecc7b79c7dd79d625af23c08c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10c8e5f5ebcd3ee207a3177c3fb8ae145e4abadef00cc2cea886f6325f944652"></a>

## custom_network_config.slo_config.no_static_routes — custom_network_config.slo_config.no_static_routes / 1827449db3f7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-6c010a8fab4a55fd7e6a4d2e6fe8b96305e5d56af17534848ebcad956b4b6c92"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-bb0799072ea2313adbf5ea99b74f966e533f0d538e52477edb1a006aedba24b0"></a>

## Direct properties — custom_network_config.slo_config.no_static_routes / 1827449db3f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5468f011c335bd781216bd1221b94b5ffccc8e7abe0758176087d7a8f6cbf9c7"></a>

## Next pages — custom_network_config.slo_config.no_static_routes / 1827449db3f7 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f400355fb47be5b7f3f36bc266ec32ddb48dc98a9e5a29a5279d3c264725ca79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e837d9274b37a50520ef5b1ebbb77c6e84e5740a31cd50bba86bfa47770495f7"></a>

## custom_network_config.slo_config.no_static_v6_routes — custom_network_config.slo_config.no_static_v6_routes / cdfada481a0c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.no_static_v6_routes

<a id="canonical-e0b07a7fcfe688245dbdcd1d05eaea633ac8fd836ad0ecda95899c5ee7328ee9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static v6 routes.

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
no_static_v6_routes = {}
```

<a id="canonical-492000017c114a7931e708fb57ee103248dbcde6c23ad5c14d81448d2a6ac1cf"></a>

## Direct properties — custom_network_config.slo_config.no_static_v6_routes / cdfada481a0c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6d4f8188524fa7a4ddbffdebdbad91919dee4123ea4b09706a5969f63025783"></a>

## Next pages — custom_network_config.slo_config.no_static_v6_routes / cdfada481a0c / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1ef778952af4ccf26d1f0b12e525b6948f1131315487c80b3acf06aab528fae"></a>

## custom_network_config.slo_config.static_routes — custom_network_config.slo_config.static_routes / a3a2a69a14e0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.static_routes

<a id="canonical-3cb538ce0ede163333092d92304ddecde5d22ad49daa8829f068de50f7587aab"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b4a565b50508fff7bb89b5482c18ea208061469c761c0c6df256830109bedd8"></a>

## Direct properties — custom_network_config.slo_config.static_routes / a3a2a69a14e0 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711): complete subsection reference.

<a id="canonical-a006dee0a45e3d183a2691d40988c50dec2138feee34be6e95ab6a56cfc62113"></a>

## Next pages — custom_network_config.slo_config.static_routes / a3a2a69a14e0 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caf782fe47de1b674fc8662dc98f8c62894886d001a46e93f1ec2e67182ddb06"></a>

## custom_network_config.slo_config.static_routes.static_routes — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-576c2b71f995d1b5f2216cfcff0067663e2a4157aa373bd670bbea8fa3ce2dd1"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf66fcbe3995486ffcc64b1e234fdcc946f3e6b3317be4c39643656fc90ba368"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 3

<a id="canonical-ab0b712fd4bb58afa00ff3910d05fea2d9c6fa7d2bd4d64bb0e99615f1b349ad"></a>

<a id="canonical-c8264b0025dc4564ca24528e3996ea8e19faea9303fc75da31a5fb60a4e8c958"></a>

## attrs property — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-fe4a08ed1cd84f14409ac0b08f30a576962e5d34d194ac94cc6ff7fd4565e7ef): complete subsection reference.

<a id="canonical-6d1333befaadbf194ef6c8de02bdd6126cd2ba16afbccbae13e9e7590bfc3da9"></a>

<a id="canonical-f340db1c63339cb05148439a3e59de2ddfd18b65bb72a4549d8ccae94accfd7f"></a>

## ip_address property — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-7d9ba98a591ec7aaa13505ac8a1756098a7b1a210d0bfa5721f1b531f7ccd1b7"></a>

<a id="canonical-657dcc72a4fc1db41233e79e004d931871e4efb2fb14e10dd18435e7bca8a9b6"></a>

## ip_prefixes property — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1): complete subsection reference.

<a id="canonical-5b9b8dbadaf8ce35493f83c5500581e2156269745aa6478adbff2cc560f76583"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes / 367ccac93993 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-fe4a08ed1cd84f14409ac0b08f30a576962e5d34d194ac94cc6ff7fd4565e7ef)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fe4a08ed1cd84f14409ac0b08f30a576962e5d34d194ac94cc6ff7fd4565e7ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-885db528414fd8a1dd37fba73a8960641e645a4d2b6502cddcb6aa0acb8119a6"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 0e212a7af896 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-1d1dfca8497a8ac65192b92de84aa890038eb7ef8cc1881fca79d84b6115ed83"></a>

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

<a id="canonical-9400b8253d3292caace4da4daef4fe0b243dcc7d39b111ff374aa5a5bb66b9f9"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 0e212a7af896 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7eaed53c517e8ca3fd10147d9807752516746d10442263e66c0d5a1456e7cb0e"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 0e212a7af896 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c266f435f7381ba1bd74e9335834b9ae941fdb1d4027500d8aff8f3f101869cf"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — custom_network_config.slo_config.static_routes.static_routes.node_interface / 1b8c65849e82 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-cd125dcd0c1991171c7ed7f6f63da9a116f760c0a98c6f3048ffe11397377678"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-393e8e1e8b66d5ea36a2036df37b8b91ad62c899982dbd61aaaee13d7aa6b708"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface / 1b8c65849e82 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-5668c5ed12c75cd564c81c42737792d83568302ddcc235dbbee4be04c8e817ef): complete subsection reference.

<a id="canonical-6499c2eff0eff6321db575d623f485ee59aeace4d004c4bf6c75ea804892dda6"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface / 1b8c65849e82 / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-5668c5ed12c75cd564c81c42737792d83568302ddcc235dbbee4be04c8e817ef)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5668c5ed12c75cd564c81c42737792d83568302ddcc235dbbee4be04c8e817ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eab8d560d658e2be338aab21f2c1f3f57df4503e1a4d5144f3eaf92bc153d8fa"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / e97ae8ad819a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-ccdd4a6c5d5c857e719c3fc90288c657b56dfa1845284039d61891f5f3c0d071"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-5932ca13af378b3229eabf0046ea8d8debb3bc542d0d821581f353b66b03c810"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / e97ae8ad819a / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-c4477d9635c838040a8b3a50f40eab6336a4107bd7241508e26f464d2961d4cc): complete subsection reference.

<a id="canonical-81af40251f26a64b24c1146929a5b1f982de32a69a990a3235b452eefeb3728c"></a>

<a id="canonical-b78d566fd35b0e8723e727a04c558f1594bad2b6229a5ffa153aeac256bc1b93"></a>

## node property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / e97ae8ad819a / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-d812ec444903bdbd062f70780f84f161c0912920fae57fedcac775e066111a72"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / e97ae8ad819a / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-c4477d9635c838040a8b3a50f40eab6336a4107bd7241508e26f464d2961d4cc)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c4477d9635c838040a8b3a50f40eab6336a4107bd7241508e26f464d2961d4cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53a86d3b062a0466458b7648f5ccb9d0eeddb2c01072ad9bbb30a6db9478ba55"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-6abfcebe752a50d12c61f6b5cc2f19110939b3c423db66cbdc917ef5d0fda2c8)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-92a10b51afe6251cdd045abd6a528236e54929b61c354092ca613542d0ea7711)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-a0a74c86fa820fc4ccbb222ab23d17c122ac313feb27ed1af5f2d1aa5fce91a1)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-5668c5ed12c75cd564c81c42737792d83568302ddcc235dbbee4be04c8e817ef)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-39cb63c92c65c489d3ff6a650030f8a7e699e4e3e7bc65aa00d5c73c30fdffdf"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3e8f7cd9abcff53012ad6fe6edae029b069b0a19700f40bbf10daecbcf83c9e"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 3

<a id="canonical-a9b5e6f1c744758c1dec62710002b5a8dd1eeea3e399881f9b7944237a3b5b66"></a>

<a id="canonical-b7f915852aea1ab59a4be2b9118a5395a4814cd066e3eba6e60718777a7cd899"></a>

## kind property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 4

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

<a id="canonical-458388a0030bad860aa716d6fdcc71f617e6fec24a94d2b52720b7ad9fddf6cb"></a>

<a id="canonical-e8d402ed66f4e55baac50db86c7df3d83a518ce70b49caaa229a16751e9ef4e0"></a>

## name property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 5

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

<a id="canonical-4814bbd4d3b06679a779cb6efbabd991189aa96b5b3d782e373c3c988716a40d"></a>

<a id="canonical-01d738870287671f77d2ec33b834b04719e28eed3cd8bacb8152c1b4092b09d2"></a>

## namespace property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 6

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

<a id="canonical-a4fab20d6f9fcdf5906970376861eb79ab8a99861b526738f1de38299b82f97b"></a>

<a id="canonical-43b74696b054481a64b8b88b9b0071540ec7b5b77292defb16c55c30924f9563"></a>

## tenant property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 7

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

<a id="canonical-0a7cd58126913f8cad86e7403a5f987ca4968bfacb24ae1a78f82dd6f459e0f2"></a>

<a id="canonical-48c72a2b537b57cd7ae14d2e0cdb925dfa49f3e7e9a53f958c4157f1e5614959"></a>

## uid property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 8

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

<a id="canonical-a859b30deedb25137b639a370696e8137932bdeb0e8b528775bf74fe36f6606f"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / 2de1133840dd / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-5668c5ed12c75cd564c81c42737792d83568302ddcc235dbbee4be04c8e817ef)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d2f7236f0755b9fffa2d4c805bc2668311ec490b6726f5f2afe35fe26a7ea9c"></a>

## custom_network_config.slo_config.static_v6_routes — custom_network_config.slo_config.static_v6_routes / dc97594039ec / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-e1ca88e3eb08b5fe064456366979f23e8bb9cc670bd230b1e932c3d7f9d54342"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-82de8ad5d564bf24a8a44c60c385a3e5b974bbb0c61e7f023d0c79af90cbd928"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes / dc97594039ec / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44): complete subsection reference.

<a id="canonical-6a21d5c18570e62dceaa3d2b4b35f81f9507c5947247d734dcc552bc60d2586f"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes / dc97594039ec / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6cfffec943d1538df41ba05c3e0aabfb12d077c97fffc49c37f399844f68ade"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-846cb5d0a1e3af2b2f458c9b853ded947293f33b04a79ab635839bdc1b73494d"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-39a038477a75b6be25d58558db511c95be01c6e850b795b12321d6fb76f56d55"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 3

<a id="canonical-e70b6d2fb4317cc97197a8966ee450847598ca9551383bf3372b3255e1882304"></a>

<a id="canonical-b45708047b1003d09d06ef93ca470b6e01d80c2ddcbd8f8c3ccb81b2c4cc00de"></a>

## attrs property — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-32971433455f253d4678dd35214c16cd795b6b3de1df02080e65e3aa0f682731): complete subsection reference.

<a id="canonical-ab14c69d9226857e871d422042e4e3c4fe0ba5750fd87ca15786e4c254b246e2"></a>

<a id="canonical-74596e953c2da8b1454b4fada85dfbc1d651373f8ad9046084063b2be3420ab2"></a>

## ip_address property — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-7d4823fae492494125184893abbcc9868e4075fab130328492f668a4434cae9f"></a>

<a id="canonical-9eccddc7fef310c2c6ace4797a2a7b5b9af0d2213a11516914109431d9e927ed"></a>

## ip_prefixes property — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0): complete subsection reference.

<a id="canonical-c8ae32d9dbd55fc4d2042e808f810645d557286b6ca22aa37d6e035adddd8681"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes / 8a127be87ca6 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-32971433455f253d4678dd35214c16cd795b6b3de1df02080e65e3aa0f682731)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-32971433455f253d4678dd35214c16cd795b6b3de1df02080e65e3aa0f682731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-061e42d3e3166c5b6f6d6618ab819f42441edc36d25024a2106debca34095d51"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / a4e484c68863 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-29dff930e8ca3cf7b7ae2f362a76ce0f9f13c4123a0a1b3fa28da6d3ae76c103"></a>

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

<a id="canonical-5b2ac0dbfb0dd0fdb9bf8306b0a90ef8892a0b137b66dede1336ac0dee08a42a"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / a4e484c68863 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb5222a5c1636f65178474c5831f9ffe9686f6d6adeceab15fa26c4d52b2d388"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / a4e484c68863 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-743fc92dd8fea0520b47a8c7c8eccab1ea2802edfda4206117ef6165c62c5d02"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / 8e57a95602c6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-ca81cbfb752554bb1dd08926115c7d5912142a473fd187b8ae546f8d48f1bdec"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-203e49784d1d7a4bb137f9766d6a856c533d619d114eb310570cf639d5b3ae2c"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / 8e57a95602c6 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-55d5b52aceafc209c35351164e88b7994055024062b237c058dc1c141688dda1): complete subsection reference.

<a id="canonical-e81b76e18bbab9fe28e014c7fcee58bc5ae818466feae4126ad38e0f61c72d5d"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / 8e57a95602c6 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-55d5b52aceafc209c35351164e88b7994055024062b237c058dc1c141688dda1)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-55d5b52aceafc209c35351164e88b7994055024062b237c058dc1c141688dda1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fabf4831bfe99e15fe34d0fc9af1de1ec84c893b348c475eb1628c5bbd16d8a8"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6be1649ea392 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-95a76cc93b551f2d75e0cefb92f8da59a4f5ee0d6e413e0ab665360c234f78ad"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-17970a789752fffc9ca1974c72c97642bbe9c26116f7d24484c390f54efabe52"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6be1649ea392 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-8860474bbe0d47024f585cb987022d972d479915cc4c47550b60d9aadba7d1ae): complete subsection reference.

<a id="canonical-40966f127294b8d2ca8adc559fedea20cf5d8cb0012296cc7bbea1db7a24c4b9"></a>

<a id="canonical-a08e8032f5b85451c0df1096bcebaa27442648a7ef2d031f41ba1d0cf930fd5a"></a>

## node property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6be1649ea392 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-c25a61ca6487ef446b7c129066d4ce30ba808cdd9c7976a0fc5c2a7ee650d669"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6be1649ea392 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-8860474bbe0d47024f585cb987022d972d479915cc4c47550b60d9aadba7d1ae)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8860474bbe0d47024f585cb987022d972d479915cc4c47550b60d9aadba7d1ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7150b517e8e352c16043f68ad2e28d52d3ee8da2aa5daa2133fd53b7691a62ff"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-9e1ad2299d79e738428247be342f25c14573ed8775efe498cb1abef5f2c4a67a)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-4359c1858179f211f728adbf73d0e23a4aca69e6988790974244d24d1160a5db)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-e3c97111a59faa3beff3b93f44542e8339558e8ade6b28ec9f4d82f4c9296b44)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-e01ac3daa1731c7c47dbbbb99fa5ae5e74110d51385bcf8a7d2aa1db52a05aa0)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-55d5b52aceafc209c35351164e88b7994055024062b237c058dc1c141688dda1)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-a61ceeeb2f44bbad40fe213d9ab93ae6bfb46dc84ee94a9a1179bd83b9dc09d4"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-d34a3074e90a389e1f4ad9d3cb1c580f63f19e7485321cd1eb87d6ea232d6594"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 3

<a id="canonical-47c4dad5980536a9ccf7bbce25668500f5eba9399a0923c226c83d5b4c7f2c55"></a>

<a id="canonical-4a65e062d84baa7a3083262293209e1701346308bd90a314020d3d5da83cb685"></a>

## kind property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 4

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

<a id="canonical-3c328a3fc27655132d5e2941c134ed799064ba21b8a0b9051658fb85f1041eaa"></a>

<a id="canonical-5ffe81754914353fb2219c0a1a0840da600e24812b7484562fe88a5d94135529"></a>

## name property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 5

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

<a id="canonical-297bb953561e2ae7ec7f8d2cd4f7261b874d7f9690fead5e6d53333baaacd536"></a>

<a id="canonical-4840231622cd2a76e2c605e40f05f5eed302a0db70c143b7c5661e81e2434bb4"></a>

## namespace property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 6

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

<a id="canonical-573acaa49dde8ec6ea7c9f0fff8b7cace1d019797caf3799e68a3fa9f3cc59a7"></a>

<a id="canonical-55f36fe7cfe90a5abd79de9f10900d3f29be46c58016d18ccc3faf651be20315"></a>

## tenant property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 7

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

<a id="canonical-4c7870005935a5d6c296893ef09b307fef01097f8226fb54e76f40a7da3f0a9d"></a>

<a id="canonical-5022225dfe73b7551732646bb7bf13a2ab99e9525f0a48639c9784c3d418303b"></a>

## uid property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 8

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

<a id="canonical-7fdb7d4cdb8acbd6c984a21dd5f47a156cbde19f51906f93852c30d66b040638"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 6d5fcbbedd75 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-55d5b52aceafc209c35351164e88b7994055024062b237c058dc1c141688dda1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-347d07f4aa76a18174d77fd12792b7fd6ddb0a0e131714c9e0cdc995d9dd2b04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8998d9d66fe012a8f49ee0ada4f41c8af901051612e37cd47e81918e0eb0528e"></a>

## custom_network_config.sm_connection_public_ip — custom_network_config.sm_connection_public_ip / 05d8e32d5e36 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.sm_connection_public_ip

<a id="canonical-af0956a23f3c13d523b03d3f4c4a0ce7e659badf9248340b3c710984a66ea292"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-44fc1215de57a3b3d518c4a7b9128adc19e031e69955cbebc26bc28603e85120"></a>

## Direct properties — custom_network_config.sm_connection_public_ip / 05d8e32d5e36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31343cbddfebc15778b6aa28ea3a480e17b72020e4550a31f1fc83484ecb1e14"></a>

## Next pages — custom_network_config.sm_connection_public_ip / 05d8e32d5e36 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-d9ec4aaebd7b2323da05050164acd7c7581009d53407445f2091efdb5fddae1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2584716d5b5676c13ea4961e27ff25c282d41851b191c00fb439b0952295f67a"></a>

## custom_network_config.sm_connection_pvt_ip — custom_network_config.sm_connection_pvt_ip / 9396035f7905 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-d34b457af4526ac22d41f25425ce99276e8dc79cbe9400c8cff46f2636d43837"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-67ae2c8ede45195cfa4864aba979097c9b06860cf708b680bd55c362ecae294a"></a>

## Direct properties — custom_network_config.sm_connection_pvt_ip / 9396035f7905 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-026861d8e388497c3c31357ee2cecf83c59e7acf2d68341dfcebe0d2827402f7"></a>

## Next pages — custom_network_config.sm_connection_pvt_ip / 9396035f7905 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4935fa902b7c2473c2a5da1ed93750f9602c8185f5d057680731f0a97bc1c03"></a>

## custom_storage_config — custom_storage_config / c9928723634f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- custom_storage_config

<a id="canonical-8605869d88c061c89b1499645865246d196168c6c00ec0bf498bdd97583c7f6f"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_storage\_config, default\_storage\_config; Default: default\_storage\_config\]
VssStorageConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_storage_class",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_storage_device",
    "storage_device_list"),
  validators.ConflictingObjectAttributes("no_storage_interfaces",
    "storage_interface_list")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage_class\",\"storage_class_list\"]",
  "x-ves-oneof-field-storage_device_choice": "[\"no_storage_device\",\"storage_device_list\"]",
  "x-ves-oneof-field-storage_interface_choice": "[\"no_storage_interfaces\",\"storage_interface_list\"]"
}
```

OneOf alternatives in this subsection:

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-8605869d88c061c89b1499645865246d196168c6c00ec0bf498bdd97583c7f6f)
- [default_storage_config](resources--voltstack_site--reference--group-008.md#canonical-c100f681d91edf2d377878aaf241bdb0647f59d432a2651a5b36c4da9c98f6d9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_storage_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2bbf46af98655edcc24481b63e8217da17726379f2ee9aa92cc7d821e75b75c"></a>

## Direct properties — custom_storage_config / c9928723634f / 3

- [default_storage_class](resources--voltstack_site--reference--group-005.md#canonical-73669716997241983e0ae046b0430041077057b06359181eccd356c3e3c6cfef): complete subsection reference.

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-b4d75e2ed369e19ae0cb55dbbc20925c2074c8084bfdfc121d2fa36a992ce8ec): complete subsection reference.

- [no_storage_device](resources--voltstack_site--reference--group-005.md#canonical-33e463a6608d72edd8e57ebe5f9a4b2c433a91662d0ab26f964c3813d8ce48a4): complete subsection reference.

- [no_storage_interfaces](resources--voltstack_site--reference--group-005.md#canonical-862cbb7f15193ae85399ddea4ba03b438a4bc96f88f55647611bdd19d802bcf9): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a): complete subsection reference.

- [storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace): complete subsection reference.

- [storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679): complete subsection reference.

- [storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f): complete subsection reference.

<a id="canonical-057f8b2d9992c2953f1f6e81182f131e928008457d20ac2a2b51301ff2633f2a"></a>

## Next pages — custom_storage_config / c9928723634f / 4

- [custom_storage_config.default_storage_class](resources--voltstack_site--reference--group-005.md#canonical-73669716997241983e0ae046b0430041077057b06359181eccd356c3e3c6cfef)
- [custom_storage_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-b4d75e2ed369e19ae0cb55dbbc20925c2074c8084bfdfc121d2fa36a992ce8ec)
- [custom_storage_config.no_storage_device](resources--voltstack_site--reference--group-005.md#canonical-33e463a6608d72edd8e57ebe5f9a4b2c433a91662d0ab26f964c3813d8ce48a4)
- [custom_storage_config.no_storage_interfaces](resources--voltstack_site--reference--group-005.md#canonical-862cbb7f15193ae85399ddea4ba03b438a4bc96f88f55647611bdd19d802bcf9)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-007.md#canonical-b08755996fc0dd1370f5b1a65cec228ba0784ab7c0614c94d52d969da31ee93f)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-73669716997241983e0ae046b0430041077057b06359181eccd356c3e3c6cfef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ef36fd3474891641cf71ff077d2e518278f4f57bb31f74edb19bcf1cb083737"></a>

## custom_storage_config.default_storage_class — custom_storage_config.default_storage_class / 32730285aebb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.default_storage_class

<a id="canonical-40c9adb8a5d57b2a6c068fedb0214dc6628fc02f9a5c67a29db5e4bddb98be19"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage class.

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
default_storage_class = {}
```

<a id="canonical-4fcf58ad1bd015d0f78158800a6dad6ffb18861ea32dcb5493fcc5b0499fd224"></a>

## Direct properties — custom_storage_config.default_storage_class / 32730285aebb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4c277f3a14260d3a710dfa57ad0b7b7767f4049eeaa7824149e4160a9651d5a"></a>

## Next pages — custom_storage_config.default_storage_class / 32730285aebb / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b4d75e2ed369e19ae0cb55dbbc20925c2074c8084bfdfc121d2fa36a992ce8ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c923ffbe4c1ebc6c96a30b7db84db6f1fb73f45ef2649eeaa86318e38935817c"></a>

## custom_storage_config.no_static_routes — custom_storage_config.no_static_routes / f61943244e28 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.no_static_routes

<a id="canonical-7bfca05b4960e7c60652dcd0f8f6ec1a19075e27984bb565a18583a3aad03f8b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-8242ec5878f0868f0c3eb0794c8c188d6945de5e149e0b7fef93d097b81e2cba"></a>

## Direct properties — custom_storage_config.no_static_routes / f61943244e28 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b47736fba9b4d38cca1316a4a6c405486be481e71751e7088a6e43eaccf7a062"></a>

## Next pages — custom_storage_config.no_static_routes / f61943244e28 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-33e463a6608d72edd8e57ebe5f9a4b2c433a91662d0ab26f964c3813d8ce48a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da244fbdaabec6352d6fa0380946b9672c1e2b9884abfeea80966d935c4ec8b4"></a>

## custom_storage_config.no_storage_device — custom_storage_config.no_storage_device / 2ef57c7f60b0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.no_storage_device

<a id="canonical-61b255e938b38fbfff86ecfc3b77995a833f4aa9a763ef152f95c7cc0cd58b1f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage device.

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
no_storage_device = {}
```

<a id="canonical-c433c67f490790be287e79ebcadac93ff2e461fe44ba8a304838e1ea565c0383"></a>

## Direct properties — custom_storage_config.no_storage_device / 2ef57c7f60b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b12acb2f45b4c5c02dcc1e52c531e427469c799f063e0a8c80ddaab71366fba4"></a>

## Next pages — custom_storage_config.no_storage_device / 2ef57c7f60b0 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-862cbb7f15193ae85399ddea4ba03b438a4bc96f88f55647611bdd19d802bcf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50623bce8d529b6a75894480c39bf77cddc8d604f156202cf8e46a7c5a8ce9e2"></a>

## custom_storage_config.no_storage_interfaces — custom_storage_config.no_storage_interfaces / c2bb372cb4b8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.no_storage_interfaces

<a id="canonical-a75ec8292bbe8e959a416e3df6111cd5a85b30e4225f280942deaadded67ebd1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage interfaces.

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
no_storage_interfaces = {}
```

<a id="canonical-d1cad161b331fb8dfbeeb3bae162f2df50bfa217eb1faf4df968f03c7d968aa6"></a>

## Direct properties — custom_storage_config.no_storage_interfaces / c2bb372cb4b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a58c9668546323d26cac67d81db3eebd216b3a7e33d44d2a6801cec715d34a00"></a>

## Next pages — custom_storage_config.no_storage_interfaces / c2bb372cb4b8 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e2f128d394ddab2e1f1d28d40dcdce349dc4734da1e7d46d836fbe40d2ca8ea"></a>

## custom_storage_config.static_routes — custom_storage_config.static_routes / f74f69b139ed / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.static_routes

<a id="canonical-a0d6c9a707813d207672d575548697125fec898bed14fa122b062a14701e4545"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-7387a34e4e7d6dfae72513df6380757ecd263e2b0b54dbff008a21e8790dbe07"></a>

## Direct properties — custom_storage_config.static_routes / f74f69b139ed / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac): complete subsection reference.

<a id="canonical-24ed0d287eb46d1f722bd12df61b464cb81035a099ded28d910f28c21da5aa00"></a>

## Next pages — custom_storage_config.static_routes / f74f69b139ed / 4

- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-695f5424f0da33197558d0c5b370fbe8e4361ed4313a3f512dd70089a0bda1fd"></a>

## custom_storage_config.static_routes.static_routes — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- custom_storage_config.static_routes.static_routes

<a id="canonical-d8b7b85a2c1989b0082941df2ad6dde35fbb27eafa9bd8c148ad3ff64418f812"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c8095a2a2510b2f681134f5d32b0c0894c39f1cd11b5f8b3592ddf5fc2e945f"></a>

## Direct properties — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 3

<a id="canonical-58df575444e367076f10a4d4004cad913b4b238fdc5aee9eaa507be2f68f5def"></a>

<a id="canonical-f8da7a41c11ca2234176768b4b370dc27ec622544e8311f1deea379a97cf30d9"></a>

## attrs property — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-285195e51b838d678fb6675711cc63c23d826fe8103218b423c2a3566677e98e): complete subsection reference.

<a id="canonical-80094a14b38a327a9e5fefe13209be788476b84e292b414473b8d8d5889d1fda"></a>

<a id="canonical-aae8446d1779d0cfe0fa21ac33e9d7b2bcf59f6ec11522d9267aeb0fb5315d29"></a>

## ip_address property — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-18783e58c858c8eeef1b67e32a81dadc10e8223dab1693279412df3ee5e9e3a8"></a>

<a id="canonical-3827e7b8fae91c8017e627369c67b55ab8db601f4e5b1bb8bc39ef21e34cbbcd"></a>

## ip_prefixes property — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f): complete subsection reference.

<a id="canonical-230a4c46d746e16d2013cd106f36fe53d83bd9e4a93fb2312c5f7bd1597e1a54"></a>

## Next pages — custom_storage_config.static_routes.static_routes / 6478da75a4fe / 7

- [custom_storage_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-285195e51b838d678fb6675711cc63c23d826fe8103218b423c2a3566677e98e)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-285195e51b838d678fb6675711cc63c23d826fe8103218b423c2a3566677e98e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2c0b0b5cbdaa24b0726e03d324ebfef3ea04c22f3745a8713c8a81f099afd62"></a>

## custom_storage_config.static_routes.static_routes.default_gateway — custom_storage_config.static_routes.static_routes.default_gateway / 513b8d688fd0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- custom_storage_config.static_routes.static_routes.default_gateway

<a id="canonical-c3513bf8725fc5a1ff9c66e3971e9fbb98a4617b2b9c3820cbeec0e5cd36427a"></a>

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

<a id="canonical-b3b712eefc631165fe203d233c46f0f9ed7b6e11299f944b3c91353f2486e95e"></a>

## Direct properties — custom_storage_config.static_routes.static_routes.default_gateway / 513b8d688fd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5911c588d613d168fc8139b91ca676bdfa7b0ead138f4fca8967532018fceec"></a>

## Next pages — custom_storage_config.static_routes.static_routes.default_gateway / 513b8d688fd0 / 4

- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-139a3e658567462d83bcef3ab512757a024a3c7f1d7ec64ec51c1c868fa365f3"></a>

## custom_storage_config.static_routes.static_routes.node_interface — custom_storage_config.static_routes.static_routes.node_interface / 9fcd4dc321b4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- custom_storage_config.static_routes.static_routes.node_interface

<a id="canonical-c661046290790ac25dfaad560c006ab96588f348251b3c3acb671d8a1b6bc732"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6865c25c1fac17bc052203043cd9b174304ab9940e458ca1e9421d08a4e2284"></a>

## Direct properties — custom_storage_config.static_routes.static_routes.node_interface / 9fcd4dc321b4 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-be8b4d161b24a345969011d3e255e7c87b1ba7adc86dcc90ccd62fc0c1147c39): complete subsection reference.

<a id="canonical-fd6224e15927341af40ef4e5d9787c45856031cebb175cd4f7a1b585a2012ca5"></a>

## Next pages — custom_storage_config.static_routes.static_routes.node_interface / 9fcd4dc321b4 / 4

- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-be8b4d161b24a345969011d3e255e7c87b1ba7adc86dcc90ccd62fc0c1147c39)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-be8b4d161b24a345969011d3e255e7c87b1ba7adc86dcc90ccd62fc0c1147c39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed1f8b98dde89109dbd4ba898b6c362be8fa5d0d5169267fb8c41ed4fa9cbf91"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list — custom_storage_config.static_routes.static_routes.node_interface.list / 2d53f28d6a67 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f)
- custom_storage_config.static_routes.static_routes.node_interface.list

<a id="canonical-8fd3f036fb8923f67cbb7305df449f93cb461f52515ab67ba2df403df94675be"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-931400735ba50a502df54d07a26b39645ddc0b29092cd78abd0651c0da85a75d"></a>

## Direct properties — custom_storage_config.static_routes.static_routes.node_interface.list / 2d53f28d6a67 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-5b9dfe139a63ef9b2a25f424d839d0e8b80b4f47d8475c01efc54fdc770eefed): complete subsection reference.

<a id="canonical-7342e102ad58dd748eb879d7291f5585489fcd8f787a0d7a7856ddaed9c8936e"></a>

<a id="canonical-9b74daf3a460bf45a2358dacfe0526b61fcdc39c0aadc1c3d271a5c8a060701b"></a>

## node property — custom_storage_config.static_routes.static_routes.node_interface.list / 2d53f28d6a67 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-5574a7d3da529b1c9d4ed2ef50bbcef5c29f1e4f0b9ba490af70c9a690d131c1"></a>

## Next pages — custom_storage_config.static_routes.static_routes.node_interface.list / 2d53f28d6a67 / 5

- [custom_storage_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-5b9dfe139a63ef9b2a25f424d839d0e8b80b4f47d8475c01efc54fdc770eefed)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5b9dfe139a63ef9b2a25f424d839d0e8b80b4f47d8475c01efc54fdc770eefed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-527cacf9a405feb9f43ecd66141ba0a6985ad6ea6316027ab7ecf7d9f96117f4"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list.interface — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-041fe8c1177b48214534dec8dc0bec1bd46a48157c395c9b90a8c06be735a59a)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-b0d140fa20189a1c9b8c9b2cb32bf52fb7db59252e0d0d94ac76305ea30916ac)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-93101fb9c0c5a17d937d526c31ba3496ff6a6c78046f16fb4d3e1ffe8dce7e0f)
- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-be8b4d161b24a345969011d3e255e7c87b1ba7adc86dcc90ccd62fc0c1147c39)
- custom_storage_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-b1a4eeb9ec3d936b1cdf8f92b3817dedbce1e2181d8cda7d9557ef3378fe9df9"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe31091f26ac2db1c361d6c0a8ade2c4066906fe4d3aeeadb7c95c8250764676"></a>

## Direct properties — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 3

<a id="canonical-8b5c5bffe25e9b2dbda775dc2a702c21728cb6c8a20b1425e97d1a46f2af804e"></a>

<a id="canonical-83ed99201e73b25549d4fc1cc4a0289b5ef7e0600e1bde4d2cd824cfc450c0e6"></a>

## kind property — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 4

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

<a id="canonical-04789d0811b2ac5100ec41e007b9b2fe59994c7ed84ab030b6bac9268a40fed2"></a>

<a id="canonical-f651f37830a2c7b96a60734300ab0a4b1c38d3cf445c9c2875a884bdc43bc728"></a>

## name property — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 5

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

<a id="canonical-07066e77adf8dd95bb6a682030ed47e1ecfec7195d7a9cdc309e471801222bbe"></a>

<a id="canonical-049489b1bb8bdd4b5d24b23c590be62e1b885f600697ca82a8b496d6ece3a42e"></a>

## namespace property — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 6

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

<a id="canonical-6e521464accebeef5849dc0ef17e7686913c9f20c2d03a868de417ffbfb02d85"></a>

<a id="canonical-1451521d18f02a7469faaadff51ed548175645ca12d082d63dc6efba572c3ee6"></a>

## tenant property — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 7

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

<a id="canonical-9621f663253d547ba3b351357e90bbcfa1242f9684554ed3c36b6720d51d9b98"></a>

<a id="canonical-b869297675b58044f24668f25de42d0d0b0cfd1c934e4b8281bbd618f4f4ba86"></a>

## uid property — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 8

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

<a id="canonical-ae7294814eb5977534726ee1213c55348f2714ec4158f2c452be8da4df50826e"></a>

## Next pages — custom_storage_config.static_routes.static_routes.node_interface.list.interface / 0f6e26f2317e / 9

- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-be8b4d161b24a345969011d3e255e7c87b1ba7adc86dcc90ccd62fc0c1147c39)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70e71ee20cf5e314e4badf961e4ee16a149bf13b59cf25f6e4cddd6570ff44d6"></a>

## custom_storage_config.storage_class_list — custom_storage_config.storage_class_list / 194d7c7e8cd8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.storage_class_list

<a id="canonical-971723a60e45a5b5425fc91e425c50e16b8216dbf432532c60a83f4c7dc42863"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b15c108838760691a4e9ebc79cc02cb5d42949d5919cf6de531103d0139141f"></a>

## Direct properties — custom_storage_config.storage_class_list / 194d7c7e8cd8 / 3

- [storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f): complete subsection reference.

<a id="canonical-ae88a70d1fb891cdd2640317c416d93382e5a2e7072333164c08e5aaa9d385ab"></a>

## Next pages — custom_storage_config.storage_class_list / 194d7c7e8cd8 / 4

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
