---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-25d2c2ad93fb85a8c912f9aacc3c4974dbecfdfbaa8087f8142a3da0f1524764"></a>

## tenant property — advertise_on_public.public_ip / 13db4c8fec6b / 6

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

<a id="canonical-710f8b2681b1f2deb9ddb0944fda1ebd1eb207b38900930d28f3f58ad99e7697"></a>

## Next pages — advertise_on_public.public_ip / 13db4c8fec6b / 7

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5684e25de901e83fdd06b403aedef7b7952f266b34a2576f9c1eabbbabb5b10d)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-93c42f400565a50ba7899bf25a14840f61f30439c967fbdb821dfad3d26713f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f74ce78adac77272905c4bd634484c211d10c1e0069ad432b7c0a55ae90716c"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / a037b48bde72 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- advertise_on_public_default_vip

<a id="canonical-5bf02f2574ca18e1572f4a5e318efeb20e4f23b884c4cdd1a7053091eea130a7"></a>

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

<a id="canonical-ea35dba5d6e1fb9c9c44eff4d3147e18c1cc140a7a0adcc74a3e7637164a6383"></a>

## Direct properties — advertise_on_public_default_vip / a037b48bde72 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7c900a56079116c915b3e01aa6734c5a67a2879d742bb66c01aef1430c574af"></a>

## Next pages — advertise_on_public_default_vip / a037b48bde72 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-58d165c905a3aa921c6ac38d12138ce201d65967f9ffc5dc81b357df7ae71d92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01b87df0fdbdee559af23c8245f7966b332dc53b2b828dff2d6075af30c3a7c3"></a>

## default_lb_with_sni — default_lb_with_sni / aaa0d46f43f9 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- default_lb_with_sni

<a id="canonical-fb0684273e818fd43aab97fe6e37b100e665012e3122609c33a936ce78119679"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-fb0684273e818fd43aab97fe6e37b100e665012e3122609c33a936ce78119679)
- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-965fdafaedb81e2f72308d6ed003e2ac4e28a96e84f352bc1b99b736d8b2d726)
- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b6a2633f2cb6856d063cae523509c0ab9608250c264c8c8489e72187a9450702)

Select alternatives according to the provider validators above.

<a id="canonical-264ba7312ab63eb22360bf6d7780c44dc548113d0e0b937b0d72aae7cbd34a3a"></a>

## Direct properties — default_lb_with_sni / aaa0d46f43f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a3565c3c317aaa9ed4a97f75264f3d4103fd3ecf9f916086c14cd06529fadf78"></a>

## Next pages — default_lb_with_sni / aaa0d46f43f9 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-b1801b83aca2b6ca4df5d63ecc6355d3044c0feac6da71d837ff4f327a6aaa0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cc6f274948b78399456663925e4b6888bc8b2a048b39225542e7b0b478b8e78"></a>

## do_not_advertise — do_not_advertise / 90ee39e041e6 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- do_not_advertise

<a id="canonical-01cc0932a46a5e320283e534345f8c825a17d59bebd68c99ece71365db19d7c0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-e97db0d8f4469f0fb718444009fc9c3ce537a69071b92214012cff4f965fdfa4"></a>

## Direct properties — do_not_advertise / 90ee39e041e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc336c10dafc966fc04cac8c8fb603d8726c2b104c97bf88fa28dd9ca3d1bfc6"></a>

## Next pages — do_not_advertise / 90ee39e041e6 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-92a4b8fc91e5719ee6dabdf2bf0c8c07e6c9afd663ed699d176b12d95be766bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e8cfdc12b8ebf30978b6df22223f687f4f1e8db55985b0ff91b1c5d07d3bd9c"></a>

## do_not_retract_cluster — do_not_retract_cluster / 4e5c899f861a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- do_not_retract_cluster

<a id="canonical-091de0d4244a87b5f28a89d1b769fbba251f63159ff4955552ce45874ce8c26e"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-091de0d4244a87b5f28a89d1b769fbba251f63159ff4955552ce45874ce8c26e)
- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5d2efbc23bdf348e8a597d843d71f01b03efd513109e51807d41264372f8ab6f)

Select alternatives according to the provider validators above.

<a id="canonical-73610979bb2b42c5217aef73b45bfe709250dca97d3dd9af93e608033e9a0daa"></a>

## Direct properties — do_not_retract_cluster / 4e5c899f861a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f43184a3ac86cf8f755fda008addf063e3d74020afdd84e4b41c02626a6162ac"></a>

## Next pages — do_not_retract_cluster / 4e5c899f861a / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-6fefce460aa1c56801d63f400e477b2c03db10c5fade95cd00405f57b6474f06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-418ed87579365139a095cc5b1bca3a508745460beed2f1b503335a4b9b107fdd"></a>

## hash_policy_choice_least_active — hash_policy_choice_least_active / dba5a1fb4509 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- hash_policy_choice_least_active

<a id="canonical-276f5b12ebcd77b22ad7ddc5af64d67552eec4766bc436625a8e827c56e62a70"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

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

- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-276f5b12ebcd77b22ad7ddc5af64d67552eec4766bc436625a8e827c56e62a70)
- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e63ed6ce45d8bec7456f0a0b6d3e37296135057044b29fe677c537905a2d1ba6)
- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-070f057da974c9821fd7516399b7360b78161a5faec435deb93fbc65fd8b23eb)
- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-01dd2b7992898b9df7a374e9b65ecd82f63a3c952ffb253f77930b95754a7daa)

Select alternatives according to the provider validators above.

<a id="canonical-0b3e7529b3320204604b4e4a82b2a36d40783f6ba4e7726e35a73d31c56c5896"></a>

## Direct properties — hash_policy_choice_least_active / dba5a1fb4509 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5230e62b87e859e60ec510b73fa533aee4fc5ba24b1c92668a18d76a709e434f"></a>

## Next pages — hash_policy_choice_least_active / dba5a1fb4509 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-7b4e92e9b7e4a2ce124345a9c73070c6c23728a208081113b06d1f9771e991f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efa0d9398090e3826ec8f6a9b7013ee368e4cdefc4a57f813232ea331c21ff0e"></a>

## hash_policy_choice_random — hash_policy_choice_random / cf718eca7bc5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- hash_policy_choice_random

<a id="canonical-e63ed6ce45d8bec7456f0a0b6d3e37296135057044b29fe677c537905a2d1ba6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice random.

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

<a id="canonical-36aa30b3891b4354e2170c86aef643a90f71073d0e153d3a9f84c6245a75a495"></a>

## Direct properties — hash_policy_choice_random / cf718eca7bc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52e85f866842bdbd7e57cb2ad88efd282d2bedaf8db61eead03eb03a44e7f420"></a>

## Next pages — hash_policy_choice_random / cf718eca7bc5 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-81aa7887946841c82a19ae76a9603b0fd6014eb80c03c27bd3f913d0edec525a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a478401df30048038b56ffed96d4060de9079873e68beaa9101e68d68d2649c2"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / b7513f416c21 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- hash_policy_choice_round_robin

<a id="canonical-070f057da974c9821fd7516399b7360b78161a5faec435deb93fbc65fd8b23eb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-27ec6dfd5e118242a24e5f95717c27c7e262bbbb0484f09426b8f99d61c50b51"></a>

## Direct properties — hash_policy_choice_round_robin / b7513f416c21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59b055aa7ba410a4eea003c5b958f2eac0718429cf8b75157e3ba455c0785644"></a>

## Next pages — hash_policy_choice_round_robin / b7513f416c21 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-919c33641590e0126296cedb9bf8574ce4d8993d7c5af43da49b3145d89baa95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a2f9a48581e51cd8323595510369aa11d480fd9f9080a7230a7ceb2b8fe0428"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / d86f2f864878 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-01dd2b7992898b9df7a374e9b65ecd82f63a3c952ffb253f77930b95754a7daa"></a>

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

<a id="canonical-95625d9f7f11efbc6c6cc0f13f7672b1e6995ba7d1b837b9beb6a0d08cbe4b17"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / d86f2f864878 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cea249c65b6c0c086558d380bcfc819b0768250dd8def608441cdf700568511e"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / d86f2f864878 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-eecd9cf3f444d97ef8d07cd7dc143b76295c7f091307b4c0ec0be1395b0afd03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90091c96323f75895a2d35d2686de7ae32a17bd088e041fb92b949e268c159b9"></a>

## no_service_policies — no_service_policies / c531d73945aa / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- no_service_policies

<a id="canonical-1959f2c19f4e27dd4cfd683029d504e1fca822c6390b85c034c7f8cfa01575aa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-58a3e25ccf7528073354ba7de0fe915ddec2003e37b9d20e17605ba7303ab7b8"></a>

## Direct properties — no_service_policies / c531d73945aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81975dac22cdc7267a66d1d90234a5708b483380dcb214369f4da886524f9ea1"></a>

## Next pages — no_service_policies / c531d73945aa / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-fbd7a2a7cc37149eea9e418797c4c0076cdcf0920443426b82cf1264fe067cbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-897ed9b096e027dacc64df512ae0a07ffd129bf9f903c701bcb492da4826e0e9"></a>

## no_sni — no_sni / 6a2f9d6b5f78 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- no_sni

<a id="canonical-965fdafaedb81e2f72308d6ed003e2ac4e28a96e84f352bc1b99b736d8b2d726"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-c396244a28fc0d4c7bd78774bd3a87cd1c7c65f6446d6a4ca662622d08b427fa"></a>

## Direct properties — no_sni / 6a2f9d6b5f78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64fff811f374fe85f031d121630c3c0db6d2b1ece1d3b12756789d1d7f038f0b"></a>

## Next pages — no_sni / 6a2f9d6b5f78 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ffedea0a644dd10482f88d575c2347ed60d610cdd68c57d32afeaef8e0d3258"></a>

## origin_pools_weights — origin_pools_weights / bfd02d62054f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- origin_pools_weights

<a id="canonical-e08be0f4252fd01319d67849816b91515c2ebc2959d64debca59b0f010b1fdc0"></a>

Type: `"list"`. Computed.

Origin pools and weights used for this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-26ce850f0f8c26d65b54d38fbb1c9e3fe89d350f4b17de2f3d5c26acc0e2b9f2"></a>

## Direct properties — origin_pools_weights / bfd02d62054f / 3

- [cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-50ed89cd098cfa532e9335fe6e69efb88ad8fb0fe396a3aea34e0bb418662cec): complete subsection reference.

- [endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ef230482e6835193f3ace8273961f92674e36886ebab758e78bcb312c49ad489): complete subsection reference.

- [pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d7252b75c8f0df51e93a9e8d4ae93b335c1aa6a6da8bbcfbd74f1e11dc32ca4e): complete subsection reference.

<a id="canonical-13c562149f911e9cd8e5f5dc67346c376f3513b08b0596f30c4ec1f56b958e76"></a>

<a id="canonical-8c741b2cb6e4ed3e8bd93f43ee6b0c7ab9d938769e3cb6b7d9e4ff7f9451ba2b"></a>

## priority property — origin_pools_weights / bfd02d62054f / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-6d7839ecb65bb4e77339d4f5482ac2fe81046117e1303de79121b6281bd83476"></a>

<a id="canonical-f8527294f15da40f94a7c473124f7f8d7600236bf3f614fdc665d86efc867360"></a>

## weight property — origin_pools_weights / bfd02d62054f / 5

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2f865466ef4378ba8b4c7ef41b2633c8bb976b822839d422df52d7ba2ee10e5e"></a>

## Next pages — origin_pools_weights / bfd02d62054f / 6

- [origin_pools_weights.cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-50ed89cd098cfa532e9335fe6e69efb88ad8fb0fe396a3aea34e0bb418662cec)
- [origin_pools_weights.endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ef230482e6835193f3ace8273961f92674e36886ebab758e78bcb312c49ad489)
- [origin_pools_weights.pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d7252b75c8f0df51e93a9e8d4ae93b335c1aa6a6da8bbcfbd74f1e11dc32ca4e)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-50ed89cd098cfa532e9335fe6e69efb88ad8fb0fe396a3aea34e0bb418662cec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5330efa2269b641d46f5fa7c9d6a95160d106d975675f36472ff8cb68734088e"></a>

## origin_pools_weights.cluster — origin_pools_weights.cluster / e7b3964f80f7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- origin_pools_weights.cluster

<a id="canonical-5719014c81e92c4c5d1053cc770c6251141633150eb20035c3e3cca76047d415"></a>

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

<a id="canonical-8b265fca697a1865d7731917de44506bc0615a5318a624eb12e4d5ab602933b0"></a>

## Direct properties — origin_pools_weights.cluster / e7b3964f80f7 / 3

<a id="canonical-a61fd4e9cdf8fe0d646bb6fd4ddd7d34d0af5d79d85f5776d6e9e086e8f5c3b3"></a>

<a id="canonical-a1d4627618f15c7f721527ca11d6932b5dce45baa48db3be78cae6294c6f4919"></a>

## name property — origin_pools_weights.cluster / e7b3964f80f7 / 4

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

<a id="canonical-e0fbc76dd38f030ccdef27a4dea599786c318626032e71a46907d14b603be28f"></a>

<a id="canonical-7c540a85fef978bdd7144336fd07c5d90602e2e4abe17cc4de4f0911b6f9a3e6"></a>

## namespace property — origin_pools_weights.cluster / e7b3964f80f7 / 5

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

<a id="canonical-30f6dd08a9107495ca41e78af63ade06913f8a5e5bd21d936980f11aaad3ca3b"></a>

<a id="canonical-1b03bf3bcc5d3c217b2128cf39b881f6fe64ccc62cb2f6260491bc4232a6899d"></a>

## tenant property — origin_pools_weights.cluster / e7b3964f80f7 / 6

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

<a id="canonical-ff658f35843b46c6f13e69c240c17c1c1432c935e229ed51220c50648643beca"></a>

## Next pages — origin_pools_weights.cluster / e7b3964f80f7 / 7

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-ef230482e6835193f3ace8273961f92674e36886ebab758e78bcb312c49ad489"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d058accc914244b16390cd54959933b784f94bd11ea61d6d69d2564df413e885"></a>

## origin_pools_weights.endpoint_subsets — origin_pools_weights.endpoint_subsets / fe10d05a699a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- origin_pools_weights.endpoint_subsets

<a id="canonical-e86dee8129b14a97b1077a0d835455aa240f45752d6756efd48feb75507d2fd9"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-ef02810a7d0bc30f7d9dae1e8f49f25cda8286af69ad9bb7a7c7e789de631aa6"></a>

## Direct properties — origin_pools_weights.endpoint_subsets / fe10d05a699a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df2446018bbcb61abad495b7bacaec8dc582ddd078668d95f6601aac7f9e1955"></a>

## Next pages — origin_pools_weights.endpoint_subsets / fe10d05a699a / 4

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d7252b75c8f0df51e93a9e8d4ae93b335c1aa6a6da8bbcfbd74f1e11dc32ca4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-180cdc98a51eda511905188c9fe95b2846c5867dc0405ce353dcaeda31d07bae"></a>

## origin_pools_weights.pool — origin_pools_weights.pool / e4a1ed70d06f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- origin_pools_weights.pool

<a id="canonical-5e9df8be7a18347bfc2c08053dbdf17eaac2b6a03d73aa9badfdc174c6d7aeaf"></a>

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

<a id="canonical-0b0d2c6f824516211220f90c54b8e15e3658dfb92239806120823c83846e264f"></a>

## Direct properties — origin_pools_weights.pool / e4a1ed70d06f / 3

<a id="canonical-62b1cc7112b4aa495eda22fd3b3fd451a022be5b9f9fca8fd705f3cd5ceb2b18"></a>

<a id="canonical-a0299a93fea868f3a822240d3902a88b927903ed82ae9f51e3b71d93654c67e1"></a>

## name property — origin_pools_weights.pool / e4a1ed70d06f / 4

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

<a id="canonical-ef404c1aeca4b6205718974e08891675d0a3c700bf38e68e8fb38e6e28bd06b4"></a>

<a id="canonical-a6bf83ff2d42e8f490d05c07cdbaf99e1076cc8b800fe520ab6d3fe5e17d933d"></a>

## namespace property — origin_pools_weights.pool / e4a1ed70d06f / 5

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

<a id="canonical-f98bdacf88c4d534a6114ebda8c8b0e1c6542c54151eda4e75d3fa145bb3df92"></a>

<a id="canonical-a2ea30e0fa77e6795e4a26e490af5d96109f6dd157d72d43abceb385db0b11b3"></a>

## tenant property — origin_pools_weights.pool / e4a1ed70d06f / 6

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

<a id="canonical-5ef2539131e6264a3529d6ff26b4e84c34f08b2a3aa964a2eec4dbb771f9c3c6"></a>

## Next pages — origin_pools_weights.pool / e4a1ed70d06f / 7

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-cc1c3b2f483ba6234b58359b9208b3f5b95fbdc41442e47e1b1563ca0af0f2e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25e722be9d77838d280b298369844a698f44c5d194c5a37cab1e722cf5b1808d"></a>

## retract_cluster — retract_cluster / 86ba1f149b78 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- retract_cluster

<a id="canonical-5d2efbc23bdf348e8a597d843d71f01b03efd513109e51807d41264372f8ab6f"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2e5c13b0f6fe03e09526d66f53ae12414953551fc2655fa044ca94edf8c7b4ec"></a>

## Direct properties — retract_cluster / 86ba1f149b78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a2bcc36f1008c83080589fc62535b08f39e76251eace8917e119c0e1582ae0e"></a>

## Next pages — retract_cluster / 86ba1f149b78 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-1b22f5ba2b642ff390ebaac0f099fe50f84b8b44393fb2f0efbdb4c499f4b7d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17aeb073c3d4875b3554c72851e3b5160c36be289c3f310b40bcae7df08dfacb"></a>

## service_policies_from_namespace — service_policies_from_namespace / 56de0236702d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- service_policies_from_namespace

<a id="canonical-f13cd904bb7cda8dcac837323ed3c4fa516dcaf5b3266754719068b762ee4480"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-b6c56f4730de725d5be87e39122095195e0ef76802b5e6a8867bf0bdda6be90d"></a>

## Direct properties — service_policies_from_namespace / 56de0236702d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17b51447b78ec84d09aa4731f71ff86a7d83a92d542cf5013bf40657416f3c3c"></a>

## Next pages — service_policies_from_namespace / 56de0236702d / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-9fc01974e3956c45c697aa96fa01d9561f837ca5f51871023cd162da605ea22b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24887561a1b482413539745f867d3e9f4081a6e74210f5da4895c7930ca7b5a6"></a>

## sni — sni / 3abdf91b45e8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- sni

<a id="canonical-b6a2633f2cb6856d063cae523509c0ab9608250c264c8c8489e72187a9450702"></a>

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

<a id="canonical-9b17c4c3887846ee2a8a4e6c5165b5745f47b90981b5608a133b927f38138c14"></a>

## Direct properties — sni / 3abdf91b45e8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04df5ad890e25110697255c38477c3e1edc5e41f1a948badf0cc0eae68b04536"></a>

## Next pages — sni / 3abdf91b45e8 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-8e30515f89b566c0046c0b656f5f1cae28136176409095156357cbe5f3cab9fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-904f197f2288d4551b7182a2be0a70d69079e04178fc2219cc52fc842e11bdaa"></a>

## tcp — tcp / 5c173a293cc3 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- tcp

<a id="canonical-6fd58b64bda394756870c8b71293a539b1718404753bd67efcf8ce90155f8d87"></a>

Type: `["object", {}]`. Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6fd58b64bda394756870c8b71293a539b1718404753bd67efcf8ce90155f8d87)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6649da9a04b26b4b8eb2d282f1e67f5245a513f9bd8538217e65210664609bae)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-da1071ff0460dbe3ac648e09b3a5265ecc1c738a79ff6727775e870c422eff00)

Select alternatives according to the provider validators above.

<a id="canonical-23900180b4fa1e224c075c1ea450b12fab834dd35afa06ae77bb4e85addb7b8d"></a>

## Direct properties — tcp / 5c173a293cc3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0029bffc472cad5e886a4dbaa35434e8a8edaa42be53f5c111447d52a85878e2"></a>

## Next pages — tcp / 5c173a293cc3 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbd84a06026a84a09a04b93c78a762c9feddb5e22401b4f758e20eed8a2733ba"></a>

## tls_tcp — tls_tcp / 6a7e4b4cdb45 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- tls_tcp

<a id="canonical-6649da9a04b26b4b8eb2d282f1e67f5245a513f9bd8538217e65210664609bae"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-faae987be8f06ae1bb00856847cc6b939e662a8d39203f7097a389eb828ecb93"></a>

## Direct properties — tls_tcp / 6a7e4b4cdb45 / 3

- [tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db): complete subsection reference.

- [tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8): complete subsection reference.

<a id="canonical-98c28e6cb65f7fc394db2e300625fbe6f1ec87f435046e4a003cfb3da00318bf"></a>

## Next pages — tls_tcp / 6a7e4b4cdb45 / 4

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-254d057e459a0ce87ce75cdd8b2438c36e2c3698d1a9f1576423d572b66922fc"></a>

## tls_tcp.tls_cert_params — tls_tcp.tls_cert_params / e318a4a7a962 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- tls_tcp.tls_cert_params

<a id="canonical-dead4bd0781a288f5e40cb6f0f4b9060ae8d3c234f955ff93d80609436ccff58"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-95fd198977d102b291a037a8ddb40f4abeaf75cb650e7d0a563dfe25a0dd7c1a"></a>

## Direct properties — tls_tcp.tls_cert_params / e318a4a7a962 / 3

- [certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-56e0ac97df793576879828e3b8b4360c58b92df7c670728b7dc22a5888b79bca): complete subsection reference.

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-860c35fc1e5b71d94bfd20f95f34ef041db87f67f0fec0aca96c6a3cfcdb4bc6): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e): complete subsection reference.

<a id="canonical-3e26036e8bcd98542e9d4cb348272353bbe2d038400f950d6aafb2a84437152f"></a>

## Next pages — tls_tcp.tls_cert_params / e318a4a7a962 / 4

- [tls_tcp.tls_cert_params.certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-56e0ac97df793576879828e3b8b4360c58b92df7c670728b7dc22a5888b79bca)
- [tls_tcp.tls_cert_params.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-860c35fc1e5b71d94bfd20f95f34ef041db87f67f0fec0aca96c6a3cfcdb4bc6)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-56e0ac97df793576879828e3b8b4360c58b92df7c670728b7dc22a5888b79bca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-517591a54e264b504a36c2d3a7a0a530c209216053e7d10e733b3309c9b7a322"></a>

## tls_tcp.tls_cert_params.certificates — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-9d3635db982ed7e48de7a8fac8a1cb51eeec01c79c3293e7bfca77bb256cb36a"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a4aeb840753134b9aecd92ed6219e23e29f36319edd01dd219373e4b280c76b3"></a>

## Direct properties — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 3

<a id="canonical-3c9dc3f2b57e60ee6d162a3d21d06d4272ece514854d30d152420242bdf40fb5"></a>

<a id="canonical-d977b2798c32f03461ae4f35ef05e02133da79ea9c08bd0c72ecc5d74efc0fac"></a>

## name property — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 4

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

<a id="canonical-86a77c75d9131a8473c971dc3968ff11fbb8fd20e242c8656ca9d2fef3a91658"></a>

<a id="canonical-561bc8c31eacd8fc01687217b371710ddc5b06137f7792a93a3748fa75e884fa"></a>

## namespace property — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 5

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

<a id="canonical-8efce4f1748e1934a7a7034b7f5a68b5056b028e5707373e2ef0f6f5db4123f8"></a>

<a id="canonical-a2ca4232e50ac7e81082b22a7fd0f4f4c98eaba6848bd7ef8ba2a3f71ae4cb6f"></a>

## tenant property — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 6

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

<a id="canonical-1e642c1ce5d27bd21961322f1f85d7b077c0297eefd6ca4dc935785a7c455250"></a>

## Next pages — tls_tcp.tls_cert_params.certificates / 18090fc71cd8 / 7

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-860c35fc1e5b71d94bfd20f95f34ef041db87f67f0fec0aca96c6a3cfcdb4bc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6f628660acfd374311c8f0ef5c8b437defce592880546696153bfd9ea8a30d9"></a>

## tls_tcp.tls_cert_params.no_mtls — tls_tcp.tls_cert_params.no_mtls / d4fa12a7cbf5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-5dfbeaaaa53dacbf740dfa2357f228e5a254cffffd97a7c1dc398629eb5234f2"></a>

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

<a id="canonical-28f3452b45ef6440dbb964bbac922bf039e877b7321b6c38c25e859bdeb11f6d"></a>

## Direct properties — tls_tcp.tls_cert_params.no_mtls / d4fa12a7cbf5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5effe666198d5f3e2abe3b61672abc32fa1a3bf76d6a891ba47eef10221e103"></a>

## Next pages — tls_tcp.tls_cert_params.no_mtls / d4fa12a7cbf5 / 4

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd15a59eb6b7c7150c08f611fafbfd7c4a646b67f3240db6e93e50400128cde2"></a>

## tls_tcp.tls_cert_params.tls_config — tls_tcp.tls_cert_params.tls_config / f2ca5dffb6a1 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-59047fc953cd1e5a6d0fd9178d12912477a531eb55df1b222e35e95cf5c2bd0d"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-50d9e679d3bccdce7178deb148e80720d46f4a17f4361c5e09d30e9380768ee2"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config / f2ca5dffb6a1 / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f5fd970aa5a23b5e808fcf891683f1d14a5c1cfe47bee06f69ba803899bab990): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2da4bf96075705c979c83a3f36df1d07d2047b5aed26c233922704c459200199): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c79465d450e5481bebc3227e6a4a3f12c4c1d56d407b6990689d8069aa357c58): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a60b944d90c44da41e07a01c98fbb604b15f9aeeabd0d58c5542835599d98a60): complete subsection reference.

<a id="canonical-c419f3c39a785f76d3fa63b7b7e291510c8b60108dc814d5d319b318c1d52477"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config / f2ca5dffb6a1 / 4

- [tls_tcp.tls_cert_params.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f5fd970aa5a23b5e808fcf891683f1d14a5c1cfe47bee06f69ba803899bab990)
- [tls_tcp.tls_cert_params.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2da4bf96075705c979c83a3f36df1d07d2047b5aed26c233922704c459200199)
- [tls_tcp.tls_cert_params.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c79465d450e5481bebc3227e6a4a3f12c4c1d56d407b6990689d8069aa357c58)
- [tls_tcp.tls_cert_params.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a60b944d90c44da41e07a01c98fbb604b15f9aeeabd0d58c5542835599d98a60)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-f5fd970aa5a23b5e808fcf891683f1d14a5c1cfe47bee06f69ba803899bab990"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7516dcab7d28f320222a198ade9c5650e9ae761d99fbb3c0b1364adc01a1d39a"></a>

## tls_tcp.tls_cert_params.tls_config.custom_security — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-602b554d6947311fe0d865fb62d38cdca59b89fa749d9bf89c0936b8931e91f3"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-5c72df9511a4a178dcbe45a7255454482f6be0acb88ec2ca6156c23eaf239830"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 3

<a id="canonical-b61790c898143a9dac35a2a54a20a3f32a7713be7693deebc04666b0231f6a08"></a>

<a id="canonical-610550f1c46dacbc707006ad4843719824725028bbf1278464e3e81f56d241c7"></a>

## cipher_suites property — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8d647bb4a6145d29485b4d58f0ff5269bf63ed28d56841cfb95f0d4b3bbcd4a4"></a>

<a id="canonical-dfc32585194163c4852f6437a6cb6f0a8c75f8a43e9132a453723b9a00ce0f7f"></a>

## max_version property — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4733de4998e56446ca118a72157a3eb6b578d056f75716098fbcf444b02b1ee5"></a>

<a id="canonical-95c1b32827083fdb8ea69aa9b4d01702f63115999e3e4a641f0b04f927becc5e"></a>

## min_version property — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-807413581f6155ae6a9ff964491fc95bd4be5427645321d6b77eaedbaa729cf6"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.custom_security / 201934de0a77 / 7

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-2da4bf96075705c979c83a3f36df1d07d2047b5aed26c233922704c459200199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce6bf6d5e84f20fd41f793d95fe09dff71a67329861bb655cac7fc2cfa98e50f"></a>

## tls_tcp.tls_cert_params.tls_config.default_security — tls_tcp.tls_cert_params.tls_config.default_security / 87cbe256a84e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-c0923fe93e888dc23a70b5729cea9b54ad69a21deeae01695cf0cae3a7ffe059"></a>

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

<a id="canonical-2d7b1946065a73ab9db466798e26ce0bb6115babc9e322a16de228d132240a6e"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.default_security / 87cbe256a84e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce7584e07343571a4d419790fbf95f8c07ff4ba03e36c78f413a061da1a7a99e"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.default_security / 87cbe256a84e / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-c79465d450e5481bebc3227e6a4a3f12c4c1d56d407b6990689d8069aa357c58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb181b4df00cfe0363eca161eff2dfa9e68b56979230aae06ba05b1b95a3198e"></a>

## tls_tcp.tls_cert_params.tls_config.low_security — tls_tcp.tls_cert_params.tls_config.low_security / 8da0c937a200 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-0c1df4b2d5f504ec2675bac19cb8431878255542c81874eb37e3466e858bca8b"></a>

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

<a id="canonical-1489c38c70aef687ff434132246f16135cd84161ad36ce21f2407b9a7131978b"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.low_security / 8da0c937a200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0bf200ececbf99bb60536be53022c1825a168d321929d1414c30ca7cb610433f"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.low_security / 8da0c937a200 / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a60b944d90c44da41e07a01c98fbb604b15f9aeeabd0d58c5542835599d98a60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5463e02cd72bfdc6d3a4c2ea5234580c2dc7d99743b6ec6bd0ed76df1090abbc"></a>

## tls_tcp.tls_cert_params.tls_config.medium_security — tls_tcp.tls_cert_params.tls_config.medium_security / 03e6d0f723e7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-46e00117bf332968b77ae72b9c1687ab8d5137d3a13b70065a853d7de99c6b9e"></a>

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

<a id="canonical-5df539d2ae5c60652aa351a60e91b08012c784efb71acdabbd38a5ffe6215cc1"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.medium_security / 03e6d0f723e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3535b5fd6bf37bdaed3cccac58ce6caf05ec2f30c4f705c5dd38a7e000876a4f"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.medium_security / 03e6d0f723e7 / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90cb608132eff480b8525d3669c8b65532504c708c6d31333f3aaaa88a24111e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc1b70dc3bc0b0dd12acaa02d0edae40dc8ed2cef9bb6fb0217973dba73b688a"></a>

## tls_tcp.tls_cert_params.use_mtls — tls_tcp.tls_cert_params.use_mtls / 06ac4a83420a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-c6e71f375c0198a4de0a767c9151f012a951ea1c1ca26b899f77477c69c055fe"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-21d7897f38dc41734399afb06defce20a0e2975495a26f04e2df3edfbf704a07"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls / 06ac4a83420a / 3

<a id="canonical-e3003c52d4e81cc51ea49dbdc7a3d96762a38a8c4454c26c54b62b6210c6406f"></a>

<a id="canonical-84ed77453ec934b2f8d9c7f1dc15991daaa001bdd611863df5324fd4f98740e5"></a>

## client_certificate_optional property — tls_tcp.tls_cert_params.use_mtls / 06ac4a83420a / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8fa33b70eb43d5d1d8a13ed123f626e3247e782bd20d632d5ae3094a7188aec4): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-37919fe9ece7020df270e123884139ba5f3533743bad738ae0ffcafa8cdf5e7b): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-766da8bbdf25f1dabb19ddc1ac66f79c8051c9a92cb4209fb0c76195048529fa): complete subsection reference.

<a id="canonical-6e2c33b765640ffd512f4d98aa597bde4c01a6a5028264cfd306623a3cb682bd"></a>

<a id="canonical-e780006d56337271ead6344553558a82dba68c26ea1d0d4732f88fd007160162"></a>

## trusted_ca_url property — tls_tcp.tls_cert_params.use_mtls / 06ac4a83420a / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6e0795dde0ae7aed0d36c3121bce595e67515e0364a9bde1ef3a359c2dd73a70): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ee2e31daf4240ea6f40f93c082a5a8520048812c7d8aedeb0651910658c61064): complete subsection reference.

<a id="canonical-0fead4c0c35e084549e3606d93f6997436155b0617b67f82178633c41c2dc6b1"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls / 06ac4a83420a / 6

- [tls_tcp.tls_cert_params.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8fa33b70eb43d5d1d8a13ed123f626e3247e782bd20d632d5ae3094a7188aec4)
- [tls_tcp.tls_cert_params.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-37919fe9ece7020df270e123884139ba5f3533743bad738ae0ffcafa8cdf5e7b)
- [tls_tcp.tls_cert_params.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-766da8bbdf25f1dabb19ddc1ac66f79c8051c9a92cb4209fb0c76195048529fa)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6e0795dde0ae7aed0d36c3121bce595e67515e0364a9bde1ef3a359c2dd73a70)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ee2e31daf4240ea6f40f93c082a5a8520048812c7d8aedeb0651910658c61064)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-8fa33b70eb43d5d1d8a13ed123f626e3247e782bd20d632d5ae3094a7188aec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f83b1991994a5aeaecb5615facac7416910a12b0a909e76a76f036e6ba84a80"></a>

## tls_tcp.tls_cert_params.use_mtls.crl — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-49274aa568c97a65279caf8c884c69a21ed44162c32c2669ce110a51ec7e3d0f"></a>

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

<a id="canonical-8f02981592a17a6554cdf96230ab64e64e4b7b72b844f4c9df1c3a103d1e0e07"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 3

<a id="canonical-07a40c09745ddac316ec4411a0d9ee838ff59459281dd08c51d1c7d114156e06"></a>

<a id="canonical-fbce528e51b3f30f74c93e7ed5c4b9e5f3fc54c8144ba951e1f9c64953e76af9"></a>

## name property — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 4

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

<a id="canonical-1726181a9f6166d895ace7f1ed3a937f527b71d847c12c4c6f1fd03a8d64257c"></a>

<a id="canonical-1750848342d959a54146734263edbcbf2d449fb8108173c792428835431154b2"></a>

## namespace property — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 5

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

<a id="canonical-f10da52298c49a72ae51ddf0abe74b04d7f2f2acc99177a49b03fb86ca352b17"></a>

<a id="canonical-82c533154c520fdddd6d18158f9c74e74e14e9602c7913a8fe33a7108ef7f6da"></a>

## tenant property — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 6

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

<a id="canonical-51aeb945684995a59579edaa8fffcbe443619a6a20935e9573f59c41164158f6"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.crl / 1ffcd3a5b774 / 7

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-37919fe9ece7020df270e123884139ba5f3533743bad738ae0ffcafa8cdf5e7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-955587e905e3fb0a5f30844bc9bb171aa2fd7d12d9aa6683ddb5bd550aacd19e"></a>

## tls_tcp.tls_cert_params.use_mtls.no_crl — tls_tcp.tls_cert_params.use_mtls.no_crl / 46f5b8a0ce89 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-77a7ee254580a4bc3a9f386a0bf1e2f226c6acd6d19d0f8341aa1dd7b622e0dc"></a>

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

<a id="canonical-0d98cfadf84451335b4199c321cb6374ca6cdfdafa4465781ebf3e0aa61579cc"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.no_crl / 46f5b8a0ce89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc3570e29d6e5b49dc16ada158bcb68e5281c6c72c730aa3a3e80b8cea685cdc"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.no_crl / 46f5b8a0ce89 / 4

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-766da8bbdf25f1dabb19ddc1ac66f79c8051c9a92cb4209fb0c76195048529fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7fe750cc4fca9f7de8e310201b0b3d6b449b3ec0247e686604fdf00bc369134"></a>

## tls_tcp.tls_cert_params.use_mtls.trusted_ca — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-d0da501f02421f8c36e068b0784e568bdaf840ddbee4c06a5a7d8f1fb714d62b"></a>

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

<a id="canonical-14e87ad3f0e6dbb3f32fb10de54a17d208b90fbd97f240c3a40b156a8712ee8f"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 3

<a id="canonical-cd32564b3cc8c04adfaf8bf51a237c3a8d92cef7a55df49422acb7d879f1f193"></a>

<a id="canonical-e14d6cf227d93d89b671780a0c955530033083d24df88b03a200c04ad04db451"></a>

## name property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 4

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

<a id="canonical-a6653bece57f378cf09c4e4b89958623956280598ce23e0479f02bc44b17a61c"></a>

<a id="canonical-87ae451ca4c5482354de7bec58f2d6e3a8e5d074c1d43767becfef3248992027"></a>

## namespace property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 5

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

<a id="canonical-a46a5ef6c621ce472f7128e48c1df8b5b9b29674791837a6a47827fc7067b7b2"></a>

<a id="canonical-614b97af2fac223af90b87f0421275d1cb823f67e7bcec526610f986e29b89f0"></a>

## tenant property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 6

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

<a id="canonical-f68d2eb79b192b3c6fc72f101fe9533caae1d4b09ebe4116f69acb3f781e822d"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.trusted_ca / c77ef211cf1e / 7

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-6e0795dde0ae7aed0d36c3121bce595e67515e0364a9bde1ef3a359c2dd73a70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41c86c3e4a30215e51f32cd827932216b746ebb0850bcc784fed04ed733d731e"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_disabled — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / 3fa49480adea / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-86f263245c948c774c707fd9030c69090bdce525d67c096e4c4cf6908d65b0ff"></a>

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

<a id="canonical-7dada9f17346d9e8d39e619ff7fe88da74b88a779f5970bdd727817594b92a94"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / 3fa49480adea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-398e8827b28fc72c593810b055d3f19b25f8e14986365a2b1e5de70de2ae1a45"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / 3fa49480adea / 4

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-ee2e31daf4240ea6f40f93c082a5a8520048812c7d8aedeb0651910658c61064"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ebf31773a6de1272dc9c8f0dd223f233e11105be9df49b10c97ecf850c3d726"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_options — tls_tcp.tls_cert_params.use_mtls.xfcc_options / 22769a2d4b58 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-79ca84b11b7763b5f9aaff244db094d36809bc1d87e7fd7d83a07f9dca0a17db)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-d7fddc6f465cf8b460cf969c353ebe1f307544a31fbc01cb1638e289de2e2895"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-6a19d9bd035babff54d68cb6e309b8bd49e37b38071d89ec7e4904c790e7d140"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.xfcc_options / 22769a2d4b58 / 3

<a id="canonical-2badc0e2f8749532bba00a1a475a90d550ad6e0a7abf03715e7e9c95c239ceae"></a>

<a id="canonical-9c3eab1c0db9f87cdef0cc4c2b656b5a777afa7fdd26d98397d87a8c77a18cc9"></a>

## xfcc_header_elements property — tls_tcp.tls_cert_params.use_mtls.xfcc_options / 22769a2d4b58 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-65e7b977b46963b7427f278ae3a1bd11b3e8a990b956b1e23c5a32f559fed767"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.xfcc_options / 22769a2d4b58 / 5

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-bc03b0b1fbc89e253f575ea749d7a7089d48d658f3a0b5ae00bf2b7063245b2e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c51b8f5ba57aa928b883d305eaee65b320b68e1fe13a8d4fb7108eded29b5b05"></a>

## tls_tcp.tls_parameters — tls_tcp.tls_parameters / 1c136c122ff4 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- tls_tcp.tls_parameters

<a id="canonical-99275446ac13605bb6647d69fae1798a4be9eaef2fe16d3780be66e1d1eef301"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-acd9f92aef7959a594d56b2c10a00c739528fc1be9ed20d18bdbab6c6c751673"></a>

## Direct properties — tls_tcp.tls_parameters / 1c136c122ff4 / 3

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2f1fcd8a7aa52d6162c1e4c61fb05af09d8bce1d64ecc4c580904e526b5c818e): complete subsection reference.

- [tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f): complete subsection reference.

<a id="canonical-7e7f810f6e1299fcc4fc8df6abf477dfad973034bf1a92e79a8c45abe306b2f0"></a>

## Next pages — tls_tcp.tls_parameters / 1c136c122ff4 / 4

- [tls_tcp.tls_parameters.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2f1fcd8a7aa52d6162c1e4c61fb05af09d8bce1d64ecc4c580904e526b5c818e)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-2f1fcd8a7aa52d6162c1e4c61fb05af09d8bce1d64ecc4c580904e526b5c818e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-892fd2da17d207e504733d999e77601294eeb68394d15dc5f240eb84be252786"></a>

## tls_tcp.tls_parameters.no_mtls — tls_tcp.tls_parameters.no_mtls / 36bfa7272493 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-c17dbd8240e29b262a640d3360bdaaeaacf8e7bff790136d7737a81d49554140"></a>

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

<a id="canonical-066a976da9f987240e6b92fc8d0b54d43f49d85263fe5ed989adec8629a1d4a7"></a>

## Direct properties — tls_tcp.tls_parameters.no_mtls / 36bfa7272493 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68d4523d2c6e131992fb79f526ca9435eaec035bb5007e756141712898752b46"></a>

## Next pages — tls_tcp.tls_parameters.no_mtls / 36bfa7272493 / 4

- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eadbc3c3dd9cc4889eb9a996d9dea801d3bf6ed43e1de7d07c8685ba3d6bd7dd"></a>

## tls_tcp.tls_parameters.tls_certificates — tls_tcp.tls_parameters.tls_certificates / 6cdb48ccc763 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-464ac9efe6e532f14d5a08ff037178c217be579c9044943db38f20349c142b6d"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-c33d9e901adf6a7d9389420f5e5e3cbe9ce0972fc53ebf9d5ba73ebd5eca0d99"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates / 6cdb48ccc763 / 3

<a id="canonical-509c8676521f794a891dbbd179b46e483878be3a7397e32ebd8e4f5f3e8aeb99"></a>

<a id="canonical-4fe59d284f61c9fdded25891c38a9623237631f52caf1c5362860f869a6a3374"></a>

## certificate_url property — tls_tcp.tls_parameters.tls_certificates / 6cdb48ccc763 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-76e093204b3d511badce72a0bf5e9355fd1261d169314ce575f120e57359a103): complete subsection reference.

<a id="canonical-acfc1bd88952526c5efb84dc56532d01ba99664cf4e18d087d1ef7b3a4e36a8f"></a>

<a id="canonical-06f427ab024d1417b96b43f74de35d4ad8a8152decbba18870a2432721c47908"></a>

## description_spec property — tls_tcp.tls_parameters.tls_certificates / 6cdb48ccc763 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a5d27d8e8b4486f5eba1d7b04734a947aba46a086d10ce3ff0dcf45d70b3bf0b): complete subsection reference.

- [private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952): complete subsection reference.

- [use_system_defaults](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c78a63e0b11733a757040a47b9256045bf070cfa59d541bb7c1854cc2c344f8a): complete subsection reference.

<a id="canonical-abed99095317e27f4739b89d1bd6f94dab2bbf813c930823a6b8dcf066462c76"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates / 6cdb48ccc763 / 6

- [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-76e093204b3d511badce72a0bf5e9355fd1261d169314ce575f120e57359a103)
- [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a5d27d8e8b4486f5eba1d7b04734a947aba46a086d10ce3ff0dcf45d70b3bf0b)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952)
- [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c78a63e0b11733a757040a47b9256045bf070cfa59d541bb7c1854cc2c344f8a)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-76e093204b3d511badce72a0bf5e9355fd1261d169314ce575f120e57359a103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffae56f601b80c7bae067fd792df08434dd8e7cd42568ede7d46070aafd727b0"></a>

## tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / df88de1820aa / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-ffdb2a2daa20d8fc63817649a193b321cfaef5e7b26f4764674f5468404af51a"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-0871e7eba677666c7980ffab4c84f727437510015e0abe509cbbbbcb3b88e02d"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / df88de1820aa / 3

<a id="canonical-21cf558f9b7105f8fd5270944c3e4a723da778c97aa87668f6a877af4a012f03"></a>

<a id="canonical-584e32d5438e800e83b4d75d2433c4563aef7ff93cbb4abc87bd6a2195ceb948"></a>

## hash_algorithms property — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / df88de1820aa / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f99f7d1a42ceaa43a7d3a4cc1bfbfd8224c11368086ef9312cb4c46458600d87"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / df88de1820aa / 5

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a5d27d8e8b4486f5eba1d7b04734a947aba46a086d10ce3ff0dcf45d70b3bf0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d582d00be4a3928e2df9d32f8ad58396351eb461405e7bda996b892666c16d91"></a>

## tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / ab28a8d6ff41 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-457fbb0f055047297b0d318613d60c4c5ca28a8f9356968c78073d048985e2ba"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-745b7afaedc2a743ba23bf90a6491e2f955682416174536f131fd27970ba8fe5"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / ab28a8d6ff41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a7f35a6756682fb6835f4fb2aebf8b19d96e49a10c5597beaf45813f858eb4d"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / ab28a8d6ff41 / 4

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b564369ebc883c7a3e26d220c70e075fe3126f2dedc36c298b96878e8e04b6bf"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key — tls_tcp.tls_parameters.tls_certificates.private_key / fe5e76a52115 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-51f437676dd8e97aaf145b404295d2d44a2750fe2bc22f68d2371e53a82959de"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-650299d7dae1cb324dccddfec8da77345e7716db1c56122bbfa5d21650fc2dd8"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key / fe5e76a52115 / 3

- [blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-77cf81985f9c13d338fc3c72203dff51f63046da417ebc00a79fa443405ad800): complete subsection reference.

- [clear_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2451d406e21ecb03af42d549ef87f13e76a34e1643469e53129dc3110ee98de5): complete subsection reference.

<a id="canonical-6931ab703e2ed3572156cab279550f30be879d89d65a928d7e0bd57550095111"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key / fe5e76a52115 / 4

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-77cf81985f9c13d338fc3c72203dff51f63046da417ebc00a79fa443405ad800)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2451d406e21ecb03af42d549ef87f13e76a34e1643469e53129dc3110ee98de5)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-77cf81985f9c13d338fc3c72203dff51f63046da417ebc00a79fa443405ad800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea2239b0224ba26a806acf3ed2bf6360a8e9977fa2f995bd173854ed1294e2db"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-f8e2eebd011445308fbf055a070f4ce0bb91f822cc407ec209c22b2f5a102bd6"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-ad958c6deffe492b1cbd040e8fc0b4caef518feecdab133082b0c3ba6afa1075"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 3

<a id="canonical-cf1362f8dee2049fb3b4f99406212853189f2a33a593e06271bb2a6ee9a7bc50"></a>

<a id="canonical-f3eb411ebcfd9dc45a12d18f977e09dee698c11225fd23f991bb4185905dde6f"></a>

## decryption_provider property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-1ff1c2aa09ca5739cbb2ea161402eebda1e7b187d34de94b49faca1c341fff73"></a>

<a id="canonical-cf75a4abdc44e8d0750b640f41d2c38ba8cc9a4443e7ef273fafc262295b31e4"></a>

## location property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-8bebe8b8a7c4cb94327d2f38eca856f718fe9987834ea8787874d8864decaa6d"></a>

<a id="canonical-06a0ff74af86cd5edca6f21ec773705d98b37e13d9a701f6a4f97d42422e4e76"></a>

## store_provider property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-caec9cd730bb03bc126a46ad10eed7fd49266a3573d77ddb053ec47cae461b79"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 4d476d4eb468 / 7

- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-2451d406e21ecb03af42d549ef87f13e76a34e1643469e53129dc3110ee98de5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44083c9968a67e5b2248c872d81e760af77ab2dd780cebcb086bab1cb40f4dc1"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 9363df6646b0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-7693ae84b5c8b06e0ebefc9059c1d3ab2edf30001eb7d80fb45498b7c5cea927"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-02d2eab888e4fd9d1b40f3f07467d1c1ec3b8de2da53b46795ce704a04507e53"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 9363df6646b0 / 3

<a id="canonical-c35806fa5d59a63757fae87a8e94f03ff761e10fdcbff183573dcf784586b0d1"></a>

<a id="canonical-dc070384822c21e159bf1c490d113634d036cbd7abbc997b259badf62e4a4e04"></a>

## provider_ref property — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 9363df6646b0 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4b31f7fbb2800f07e858400969cea3c4946311281da299673745817903e929a5"></a>

<a id="canonical-1ecc28f3e1ce6f1490ac17c7d1e5f23a5450b09583d8287dce73c81868cb6d67"></a>

## url property — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 9363df6646b0 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3e97faa538c81b698aaf90a2e218ef9a0238c3a043df8e73b6180e358aaf6e30"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 9363df6646b0 / 6

- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-14ad7601dfa2dcdc7342761c09b2e469d4410370780918135a1bd6681ca16952)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-c78a63e0b11733a757040a47b9256045bf070cfa59d541bb7c1854cc2c344f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53f46f3d41b969f41a3c2615b103fa0effd63634d480cd3f0dba21d5e54c4249"></a>

## tls_tcp.tls_parameters.tls_certificates.use_system_defaults — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / aab3399453fe / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1fb99423d6776abe649612e3273715ea5053346d42e4ba209ce0b8e0920e2dc1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-9580ef048a1215655c2f66df655caa73ee095119550e73773aa19a72706614cf"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / aab3399453fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66fb8ede243753fa4836c7155c0b85dea90eaa15973aead7838823ab921561c6"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / aab3399453fe / 4

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-47a48a2a4c0a6c11bb6012c94a8c7242b4425b5c80988fe32ff0da6e6a43719e)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc76d2f790b4306974b3e8c1e5772098c00077c45daa97bfb8de7df965a9b3c0"></a>

## tls_tcp.tls_parameters.tls_config — tls_tcp.tls_parameters.tls_config / e9c83e486e9b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-a9cfd7e35c8fba86fc7b9381a9b01807e5e26df77c2b481d86a05d135bddf451"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-cb91d53b6146eae1b6fe83e1add54fa9e2cc5e8555762e79ed40abb3eb4b3a3f"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config / e9c83e486e9b / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b7f284dba1a38d1b855f2c6579a87f846d6fb2c5c3706ba842aaf8ff45036867): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-62935dd5e768664c167a71a22613bf8b146e6fda009b2f64df613e06cad5b0d9): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f8bd33942a8bb695d896013a188de242cd7ef273b8ae40f6cc2324cd858a3ef2): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7d20b18ace9de02e19208ec9d1fc47b4982b28e996e62f39e9c027fe772f2e0c): complete subsection reference.

<a id="canonical-c821a89e0b84344240553348946813a189b94cc821516eac8267411b408c8e04"></a>

## Next pages — tls_tcp.tls_parameters.tls_config / e9c83e486e9b / 4

- [tls_tcp.tls_parameters.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b7f284dba1a38d1b855f2c6579a87f846d6fb2c5c3706ba842aaf8ff45036867)
- [tls_tcp.tls_parameters.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-62935dd5e768664c167a71a22613bf8b146e6fda009b2f64df613e06cad5b0d9)
- [tls_tcp.tls_parameters.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f8bd33942a8bb695d896013a188de242cd7ef273b8ae40f6cc2324cd858a3ef2)
- [tls_tcp.tls_parameters.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7d20b18ace9de02e19208ec9d1fc47b4982b28e996e62f39e9c027fe772f2e0c)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-b7f284dba1a38d1b855f2c6579a87f846d6fb2c5c3706ba842aaf8ff45036867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27e1e9226b279aa3aad28071e5baef3cc1b7b7feb0289f7e9a1d426c7654dcbc"></a>

## tls_tcp.tls_parameters.tls_config.custom_security — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-c745ddd7211cbf6aeff46c0a71c7fbadf6ac941236741a5e936d0684127c6f5e"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-6a6c1b921b4419fd5e72460496a19edbda19a07413eae04e25857098ed5eef26"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 3

<a id="canonical-d05ae0eeb0a65e3589bb302230a501ee353272c987eb83d7f2e1869174a1fbf0"></a>

<a id="canonical-7251d459aac7407703facefdd74a7283f2dc0b40ea2106e6f515c14c4b7f4527"></a>

## cipher_suites property — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-305c08eba83f59d8f4fd30c731393da305c731338195a4ed154c9544a42a0ebc"></a>

<a id="canonical-5be1ea8b8b55e0c058b9d28e6cdd700f16bff7c03c6f78b7025880bd62e419e3"></a>

## max_version property — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fd42d7518dba16c04e29d608bdf065efbd496fdda2ba517c62152ff7548ded22"></a>

<a id="canonical-063ce6fea6e52c4eff56d0103625d093659846d55c35bd6ad618cf72a9af107d"></a>

## min_version property — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-016c566675e008a11a28d0093aa967562fe4478c400329a3bb9cc895d22489ed"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.custom_security / 2fae5c6bf4d1 / 7

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-62935dd5e768664c167a71a22613bf8b146e6fda009b2f64df613e06cad5b0d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-749c9dcbbf917430b42d3eed18750a7a1264c4486aee5b83960df4292b609e3e"></a>

## tls_tcp.tls_parameters.tls_config.default_security — tls_tcp.tls_parameters.tls_config.default_security / c0401f764b93 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-7e2aae6b94b17fc8496ee02beb2ed6b931148ed0ee34b6d6ab1be3ec868e1c81"></a>

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

<a id="canonical-af54e09e5b07d99f32f3707abb1b0aa70a8a55f84984bc38454db0b49ad2d479"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.default_security / c0401f764b93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-acbeddd565abaacd499012aae8ef8cf92ef2fb1717ac1a9d7b5a62a68f14b34c"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.default_security / c0401f764b93 / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-f8bd33942a8bb695d896013a188de242cd7ef273b8ae40f6cc2324cd858a3ef2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3f28500104fde72bec8862c3996caf7877c37a6549a32a4e3a6c38ba43dcc80"></a>

## tls_tcp.tls_parameters.tls_config.low_security — tls_tcp.tls_parameters.tls_config.low_security / b277e7ff486d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-6f23f3010bb4b542b89645de20e4c13c2ac5906e8d43d67110ce5e46f598665c"></a>

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

<a id="canonical-71093852c9110053ecb0309e8fb8c196ef794b742339a89aa5db631ed2d61cdf"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.low_security / b277e7ff486d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a109f0e2f1ebba81ef878467474913e279da2e57ef16c94e732bf6c9190bad7"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.low_security / b277e7ff486d / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-7d20b18ace9de02e19208ec9d1fc47b4982b28e996e62f39e9c027fe772f2e0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbcb7c1660ec337badad2a79d5a0a2f53f348fed8ba4330fca3c83d8855ee33b"></a>

## tls_tcp.tls_parameters.tls_config.medium_security — tls_tcp.tls_parameters.tls_config.medium_security / 8164217ee539 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-c14ddd6bc4237362e00ce855025bbe93dbd2f380dd55b5ded25fcd00ed6a9d51"></a>

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

<a id="canonical-b293af636c7396326e53d042d1a0b3dcc9779ba04c8307d206fd656e0e7651ac"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.medium_security / 8164217ee539 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3967474b800e0924041ceb484863cef36e8d1882ba83cd7beb75f125921f314"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.medium_security / 8164217ee539 / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-57685562f25725e3cd0b756b7b0a3424e0538bfca5a3dc65f76ee02449eabb05)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3d571b12e70fc4207279023a93270b8c3d21f50cd6709c7ddbfc6174423d67c"></a>

## tls_tcp.tls_parameters.use_mtls — tls_tcp.tls_parameters.use_mtls / 9fcaaa729b1f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-eba97b0a9b8cfc84e10ae26096ceaf8547e2cdca908f784b381cf3cba37d96db"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-459e6adb12985830a52466a6519fd49bb30f805e737a5656cb4a3b16707ea42b"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls / 9fcaaa729b1f / 3

<a id="canonical-343b964a5dc1ec5a11cae5e37c09f9474a105ce974083e7187f18bc5bdb3b5f4"></a>

<a id="canonical-a541962f8c5a2c812ea874f5ffb1d078f01e0f84c558539a5fc7544f3924d658"></a>

## client_certificate_optional property — tls_tcp.tls_parameters.use_mtls / 9fcaaa729b1f / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d74e41490b8172730c73be921e4b2552f8f573976a9ad07eb86962cc14147e14): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2d0052881dede87f2f88151d967da4b18bfce299d0e74f5c6bc34c6cdf42e5dc): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e34c3b7b04c4878f76fe1b9cf7ff9577e3fc7c139afc3fb656194048ab874f72): complete subsection reference.

<a id="canonical-c5c6b0b77279e2f73c608ea2ceca255604e151f3e0fbc931b25ada739661ce98"></a>

<a id="canonical-4c50049c2708010298cfd5f5f4791f00e8fe5ecc59b6ef19e9e6a1a884f6f8af"></a>

## trusted_ca_url property — tls_tcp.tls_parameters.use_mtls / 9fcaaa729b1f / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-74d93bcc97f3ee371c72aa31932592551b5c5b925ca616cceedadcddbe42419f): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-128dd0f87726d78bd544463627158e200e4a00f71fc51c7e46fdde05060d3eaf): complete subsection reference.

<a id="canonical-51ccba3027358302c5d13c8020ee74d817d324716373968d7ae1322a82c8252f"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls / 9fcaaa729b1f / 6

- [tls_tcp.tls_parameters.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d74e41490b8172730c73be921e4b2552f8f573976a9ad07eb86962cc14147e14)
- [tls_tcp.tls_parameters.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2d0052881dede87f2f88151d967da4b18bfce299d0e74f5c6bc34c6cdf42e5dc)
- [tls_tcp.tls_parameters.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e34c3b7b04c4878f76fe1b9cf7ff9577e3fc7c139afc3fb656194048ab874f72)
- [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-74d93bcc97f3ee371c72aa31932592551b5c5b925ca616cceedadcddbe42419f)
- [tls_tcp.tls_parameters.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-128dd0f87726d78bd544463627158e200e4a00f71fc51c7e46fdde05060d3eaf)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d74e41490b8172730c73be921e4b2552f8f573976a9ad07eb86962cc14147e14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f1f19a7bdadb3a37821b5a30bbcbbd17f215260ba7c3820e98232754d18a6c2"></a>

## tls_tcp.tls_parameters.use_mtls.crl — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-10815073cd332e2b247814082e3ce6a88bc9cfcffa9a8c79ac8562ba53a4c9b3"></a>

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

<a id="canonical-93b92810296f29a79f030303a086773afddd5acbfad5e7f79580b16f06399ae9"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 3

<a id="canonical-56cbacad3fb382ff77aa412ac8424e0e267d9bc7625501899907a2350f7640d2"></a>

<a id="canonical-2c63c1ea42c388763009457b29607f820fc834b5348c2306b23c7c723db85fc5"></a>

## name property — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 4

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

<a id="canonical-114bbb3d4f5606e6377f5643f2523f0ab63083db87da7cca7ebbdb772dedcdb6"></a>

<a id="canonical-03562de6e9e48fc065fde5132e9166ac3240246a96126496ee8d7460745e429c"></a>

## namespace property — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 5

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

<a id="canonical-d9fc93bd16a5d8c0ac4baf14668848d4f61ad9f78fea4ebdfbcbcd3a710e325e"></a>

<a id="canonical-74ee7e53eaf739d436792bf48adec86a077b94031f88c7f2ae5fd1bb67112477"></a>

## tenant property — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 6

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

<a id="canonical-b7ce868b3b5d6e5cd76bf9465b14f3425777efe6991a4d1525248ebdd87e7bd2"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.crl / 31767f222688 / 7

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-2d0052881dede87f2f88151d967da4b18bfce299d0e74f5c6bc34c6cdf42e5dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be52b6964dd40baab9cfc66b4c6c0c0911470d17d188d12cffc038a0b8e4aa31"></a>

## tls_tcp.tls_parameters.use_mtls.no_crl — tls_tcp.tls_parameters.use_mtls.no_crl / 67a3476d5ca1 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-3e59bc68b4c19552a3446fa8b7355f5d01df108551f761379ca82d9b6f1b7fab"></a>

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

<a id="canonical-63c7a736ec0e1f8e6b35808251cd33e04704d4c9bd9c595e28d8c62dce094206"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.no_crl / 67a3476d5ca1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1bbd71e10f46ac69d12563443894191a26d2a56897ef11b0b067b504e465686"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.no_crl / 67a3476d5ca1 / 4

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-e34c3b7b04c4878f76fe1b9cf7ff9577e3fc7c139afc3fb656194048ab874f72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dfc62a47e35f1f7046c48421f54cd0072cd2ae51e17e7cdc39a4d27ed95f7a5"></a>

## tls_tcp.tls_parameters.use_mtls.trusted_ca — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-74282a2cf195454318694c61a7c5412b6d5a3909f1c85becc18154a38bde0b4c"></a>

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

<a id="canonical-a61d092a4f1b8ebe9def1c8995b6e2990e9f53f83d339f5ab8f807266ada026f"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 3

<a id="canonical-c1c855cc2e4762e451fc8f4900d2af914c1c2e91ffc5ac67458e2af651a4b46f"></a>

<a id="canonical-715e1eb11455afe1b10d8a4802bc5033be0e019d5da5ff604736d7f63e91f319"></a>

## name property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 4

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

<a id="canonical-c6b87730655188844fe1ee0c6906073da63c393a3f74cf23c35fa0d17138a625"></a>

<a id="canonical-3d5cf9373e16d62d57af16fe86c4945343d79692d4d1fd6f40f29a5e7dc0c427"></a>

## namespace property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 5

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

<a id="canonical-d2676906ecac0b3c0281e668f3180822adca61a5e28d2cf2565bda47be41e8f7"></a>

<a id="canonical-03a9ee03b20eb4b030483438bd17e12f3dc82df9bbf5fae3bfa97ff3a0aec906"></a>

## tenant property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 6

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

<a id="canonical-f42af04cc56d7d08256c49f5041e57486985e71a75ae248e379fd64b0024a3c6"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.trusted_ca / 70c1a169b70e / 7

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-74d93bcc97f3ee371c72aa31932592551b5c5b925ca616cceedadcddbe42419f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d22d14db36cadf535ba82b31a80a8344aeb3a284996eb2a7aded6e344ab1f714"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_disabled — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 8874d81140e4 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-d8c01288c4e74c836c766f5798bcf219e2e7e7ee1af5daa0f7c68cdf8ff2cc08"></a>

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

<a id="canonical-06188f0ef3d8d758cd33b39525825f835cab314a7353932d9a29f79e5e264416"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 8874d81140e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8d792094fc3486a8d9f562f4c575297095e81480951d38d66edb618ca47a08f"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 8874d81140e4 / 4

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-128dd0f87726d78bd544463627158e200e4a00f71fc51c7e46fdde05060d3eaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95dadee58e6dfe7f65ab0816b8103134592a3e7dd6916e539850add743fab643"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_options — tls_tcp.tls_parameters.use_mtls.xfcc_options / 8a73a0139d2a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-69d81c36e00bd08457808dc400670e3c6face532aab0dc7ea261105cbc37f5d8)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-48724beffdc2bc5314a8e0044c04e991c8a7eed16c84cf9256ceb9d19022916b"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-714752bf2b02460d40fde632c9124e8a4268d5b114a3418042016729a99d2571"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.xfcc_options / 8a73a0139d2a / 3

<a id="canonical-a30ebade131adc2c935677442beb4843304b982a5a7a23242c0c55325654bb7c"></a>

<a id="canonical-a721e6085d6ef6cd204719296840be9be5b6f3acf5a3788766355b92e6b4339b"></a>

## xfcc_header_elements property — tls_tcp.tls_parameters.use_mtls.xfcc_options / 8a73a0139d2a / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-f93a947dae475ca6d7607736ec4067ca4e60db23059eea6aa7333d429633f6ca"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.xfcc_options / 8a73a0139d2a / 5

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9488c678de5acd7009df06424bac3f089f9ee002a69ef91231125aa229bd17f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1a4246e943e37e1414512783e326a560d89081c39a84f1b23214ac6375bdc2a"></a>

## tls_tcp_auto_cert — tls_tcp_auto_cert / 20aff4ce9f1c / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- tls_tcp_auto_cert

<a id="canonical-da1071ff0460dbe3ac648e09b3a5265ecc1c738a79ff6727775e870c422eff00"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-456ccc0f6dc0a1a44405347482d32e001480f082e73682faef53062e28df950b"></a>

## Direct properties — tls_tcp_auto_cert / 20aff4ce9f1c / 3

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-083509e061efa21f52f558c0eeed266df8d5c72e54dbc6bf293e81c02368c1f0): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a): complete subsection reference.

<a id="canonical-b68ab50e28ee9c4ef8bc3438077014f4bf1366d388f727298c0efe89097150c4"></a>

## Next pages — tls_tcp_auto_cert / 20aff4ce9f1c / 4

- [tls_tcp_auto_cert.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-083509e061efa21f52f558c0eeed266df8d5c72e54dbc6bf293e81c02368c1f0)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-083509e061efa21f52f558c0eeed266df8d5c72e54dbc6bf293e81c02368c1f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9862ee86b17e86b12df4b7f3a903a1b81473781a845c197bcefcc2ee11085418"></a>

## tls_tcp_auto_cert.no_mtls — tls_tcp_auto_cert.no_mtls / 78849dda4ffe / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-76efaeb96c3f588009f2c98199c0c17274166db9b92ab9fc8a93d47caef9bc6f"></a>

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

<a id="canonical-c4e4fbd13a8bcbacc3d791c7a1cba34f28a0780be4991cc2a7af1f3e9ffe4060"></a>

## Direct properties — tls_tcp_auto_cert.no_mtls / 78849dda4ffe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30b49e0f444c991cf639ce7a359dafa584300a7b1a0919a2975045fc013137ba"></a>

## Next pages — tls_tcp_auto_cert.no_mtls / 78849dda4ffe / 4

- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e2c669e693ecff94d5f7fb0ec533aeb1d9c4b9ee1b78260b51b9d996d5b4a8"></a>

## tls_tcp_auto_cert.tls_config — tls_tcp_auto_cert.tls_config / 615b27b3881f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- tls_tcp_auto_cert.tls_config

<a id="canonical-90affc6782fa8b0b557c840073a58818a22acce5528de37552c6d6ef248f0633"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-8cd05bae16834383e145c4113f6eceb2270d1a34f5b3ed84327780b07b9e38dd"></a>

## Direct properties — tls_tcp_auto_cert.tls_config / 615b27b3881f / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d8d1f1961c6ee6e97a099ac2b87f961fc1f36bc7617a403484837b30ff48c150): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-45ab1bd3117bc560c2b7e38ca3b5d9f66592522de04f92106214a2ec0ee29ef2): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-eb681a15a439af2908052bdddc142a1df56bb0b656926ea2b696267dc79105b7): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1471eff7c0a7547f7616c3b418dab84051dfdd8b9f3d4429b8dc918c669a1b7b): complete subsection reference.

<a id="canonical-ad722edb0158810787626ced0c01cbf81972306c52e32151b51b2718f9c1d605"></a>

## Next pages — tls_tcp_auto_cert.tls_config / 615b27b3881f / 4

- [tls_tcp_auto_cert.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d8d1f1961c6ee6e97a099ac2b87f961fc1f36bc7617a403484837b30ff48c150)
- [tls_tcp_auto_cert.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-45ab1bd3117bc560c2b7e38ca3b5d9f66592522de04f92106214a2ec0ee29ef2)
- [tls_tcp_auto_cert.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-eb681a15a439af2908052bdddc142a1df56bb0b656926ea2b696267dc79105b7)
- [tls_tcp_auto_cert.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1471eff7c0a7547f7616c3b418dab84051dfdd8b9f3d4429b8dc918c669a1b7b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d8d1f1961c6ee6e97a099ac2b87f961fc1f36bc7617a403484837b30ff48c150"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a13b0d81cef8ac22b27482c564dbd4722844e23613eec59d9451b0934f4222c"></a>

## tls_tcp_auto_cert.tls_config.custom_security — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-d79e8f06e44c46d4bacabbc15e99e74adda29fe557188467621979541939e38b"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-f224566b3da1ce9085f3d1c147d1ff4192f3cd3e435e8f079370dad614091b00"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 3

<a id="canonical-0bfb309cfe30b6a87df6088bb9b5f733611068e4fe6e58577f5a2d7a435fd897"></a>

<a id="canonical-c5e4422019a0de09c4a5711cc7c1b44be3ec43f40d9113e83b99dff43f247aa4"></a>

## cipher_suites property — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-74700b1a209eaeed652f1d214a7be88008599e7c4653b22f199a833a00d239af"></a>

<a id="canonical-b3246f45b7521e06270591f623bffaea48653f79781839e5b9fe1cd9f7be3f12"></a>

## max_version property — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9fa4453b48ca1ee1f47757042c07942e15440761e9778ce94600e5ba336d4889"></a>

<a id="canonical-82d339d95876f83178812c16cd6c20cea7d325e65ab3eabd3fcc41a27fc45d05"></a>

## min_version property — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d7cfc501002ed1d893264965b688ab4ef5f8539a9c61db585cd839a16fc594a5"></a>

## Next pages — tls_tcp_auto_cert.tls_config.custom_security / 59950e5a260e / 7

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-45ab1bd3117bc560c2b7e38ca3b5d9f66592522de04f92106214a2ec0ee29ef2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29e6581cab49bf17348d0e367423e39426627c378abc9860e8ad2ea5774a63e1"></a>

## tls_tcp_auto_cert.tls_config.default_security — tls_tcp_auto_cert.tls_config.default_security / c72589e9e0ee / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-6d6e5669e9ff4ff4b6161648948574397005fc9a76bd960eb1f0c104c06d8ee5"></a>

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

<a id="canonical-7e0c63bf6a79d2634fd32e8e82829122541621333288420e43c280955897d20d"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.default_security / c72589e9e0ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b2f6365fbc6c94ae9012ca84ccb96a3cac71392691732c91e5e9f0672b8e561"></a>

## Next pages — tls_tcp_auto_cert.tls_config.default_security / c72589e9e0ee / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-eb681a15a439af2908052bdddc142a1df56bb0b656926ea2b696267dc79105b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c0f12ec7ce73f67376c408519e2f60f85218c2abe1690f751b21f74ce7bb69"></a>

## tls_tcp_auto_cert.tls_config.low_security — tls_tcp_auto_cert.tls_config.low_security / 211f9b996545 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-1532cebffbd122a3fbe1598cdea8c2aabf096826ee1a9e81578b5c0dbfb5d144"></a>

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

<a id="canonical-79a6a34dd74207a029f6cecc919ea0f51d9c05c159686d0aa4e6813144cc9aa7"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.low_security / 211f9b996545 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c68466c885d3605f34362f1a8c32d2b89269fb2b3b404cfa61260b73f3c0463"></a>

## Next pages — tls_tcp_auto_cert.tls_config.low_security / 211f9b996545 / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-1471eff7c0a7547f7616c3b418dab84051dfdd8b9f3d4429b8dc918c669a1b7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99914013d4ab797346be694be51cc342dc03af39cd34d3a2350811eb5b77fdb9"></a>

## tls_tcp_auto_cert.tls_config.medium_security — tls_tcp_auto_cert.tls_config.medium_security / 01e2af2c5bc0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-730e7eb521689469723008ab4707bbb4b0ac2e164cd7a73a38893ec01e238f9b"></a>

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

<a id="canonical-dc5460b4bd2974146e1565271649100caea6efc02d0e53635d08f54b48e73064"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.medium_security / 01e2af2c5bc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20704df775ed30b17ec5a9b7be486421ba600146984f2d37b266209300872a82"></a>

## Next pages — tls_tcp_auto_cert.tls_config.medium_security / 01e2af2c5bc0 / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8c15c86409b269cf7a55242cc1d4d851beddf32ed88cb5f4c48ee5ecc573b0ab)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bcca799eb5484130350817295ca875453e217c5faeab8b62a25d69dae912497"></a>

## tls_tcp_auto_cert.use_mtls — tls_tcp_auto_cert.use_mtls / d46328e2298e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- tls_tcp_auto_cert.use_mtls

<a id="canonical-8032bb22a890292ac2a3bf76f3fa65288e9055e011dc1ddf2a381c0f06ae07e5"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-97bb6693a64b50a92d7768a613afc8474bb88b9176142c5be2fffd185a8d6176"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls / d46328e2298e / 3

<a id="canonical-74a8307f2112ca1de79aaa672ec5dc9da70abdf8dc8001e72a6d116452116814"></a>

<a id="canonical-d4a8a565a37f8b44add63d39f409fb1d61a517fde37e03c0bf1c98952438d664"></a>

## client_certificate_optional property — tls_tcp_auto_cert.use_mtls / d46328e2298e / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-04fdde51ed3d9f0a1621fe2e940c725b2d1d206fb9f0b600b3cfdb00bee6a7e9): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-dfdac3f21e98272beaa533813199c01ea0a74cef913e719267c32418183737fa): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3cec0f3a70183c7863037189a311aa69680ac922fabf36958b8b30fdeb2ab282): complete subsection reference.

<a id="canonical-e77db2b4b01c549b5391ea17e667c4e0b926de99bb151ca62d4d29c64b9157fe"></a>
