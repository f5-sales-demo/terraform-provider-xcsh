---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-7d98587b8173641a22ad9de4a792fcf95cd839e4fdec3ca1a70ad8e2092109a1"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 744d3115af70 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0689f0bfca7d26b23071a772b8d65c15fc64999db82a50290514029c682d8a3d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-0689f0bfca7d26b23071a772b8d65c15fc64999db82a50290514029c682d8a3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca20d0415b38641cf2245c18bf1d3293eb02150ce7a6e5cba855fa629ab35966"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-cc4443798a349f1e548b8c703e5bde762fa33032f232fdd59876298c2a345373"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-27de65e2dc859885887a8ae34dd522af9e52c9c0c9d7d5fafe7c50bbfe6559e0"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 3

<a id="canonical-a098bb8421eae256fdbe283076069ef8ee6df65a0abee7ba5a26cf91c2e34c21"></a>

<a id="canonical-3c6e933c5277d2c6814b470e8e86d36902015c4c7bef6dcdff5ffe75cf5963d2"></a>

## kind property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 4

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

<a id="canonical-24aaa55969e8714d13c0f9be262ea01c4f29bbaf2c1dd2c9b5017f134ff6e42b"></a>

<a id="canonical-665b1039c1d7d54cab1c7c6fd7882bce3811fd92f428a52f7cce53b440831719"></a>

## name property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 5

Type: `"string"`. Computed.

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

<a id="canonical-1357929b2f1a0a98c8b67066d7846497e8659247a94eeb044f404089f15e6b57"></a>

<a id="canonical-9bd825311a26e5aa1aec585cf69c9edc9f4fe37a489f5064ca61bfd9a370e33c"></a>

## namespace property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-dfa26edc92420f21404a4f51896695d0d1e6afcb6b70ed1a60a2ddc898036390"></a>

<a id="canonical-340343401d8f28737fcee6411b68b53664b13848dd7d2a2ca3c9633abe4ab0ad"></a>

## tenant property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 7

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

<a id="canonical-f0d25a826ca744a7fc279b7e6ea9d70cf4b0b51a24f33aa0ec7c5acaab43f942"></a>

<a id="canonical-540717bc60fbc92fa4f4f5301a5887972b81c3939c59684068a050bfe29683b7"></a>

## uid property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 8

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

<a id="canonical-4afc9863aafb34d76f2d5e89588b8a97ac24ab408a99e87c777182ca966b24f5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 157a63b86c4a / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-273193e9c0c207ac371927d068d13605caac4751d429ecc358442f30b134cee6"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ac9d965cf2e5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-98f63b7663675cc67c4095d6324b0b7a1bc7f08f359c905271e26b237cc929e0"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-545e46e021b39e6ced9fdc4fa4589ec4b7950614ca95000d8ff9716ae61665f5"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ac9d965cf2e5 / 3

- [dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604): complete subsection reference.

- [ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d148bd1433e60e3392ece674310a70603fea0ec3aa2b71ca5bdc0514bf3160ad): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1b12c82985b9367c25bed2cab761a3bb4fab9e057914e3761be5ece4a70e7b90): complete subsection reference.

<a id="canonical-4423ccf6851c0d7da9c411c4ab7a0d9c54bdb44e95b2593367bea632dc6c0b61"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ac9d965cf2e5 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d148bd1433e60e3392ece674310a70603fea0ec3aa2b71ca5bdc0514bf3160ad)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1b12c82985b9367c25bed2cab761a3bb4fab9e057914e3761be5ece4a70e7b90)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4be26ccd209ef32bad8fc866efe5d3c4271a38cd00129a42c19c7157f1dd6eb"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ced22cac7420 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-532479bb3b78e08dee2fd35f54bd4bc3486fcacc034bd8f5cf7933d04d6653f9"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

<a id="canonical-b9ef632ee03427f9617f1129f23817c4ec5b6f985c9fa8906092fe7855dec3d6"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ced22cac7420 / 3

- [ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-16b556f736a92391870ffcbd30ffc2724800c0f369389531c1e1d95b9b1c672a): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5a023b29c475b7433a2ec321d5a152b01cac8daa80418906472268d08a61b8f9): complete subsection reference.

<a id="canonical-8678a56efb868aba50f2fb328633d4f7175efe572208962a7b289e4a30e20451"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ced22cac7420 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-16b556f736a92391870ffcbd30ffc2724800c0f369389531c1e1d95b9b1c672a)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5a023b29c475b7433a2ec321d5a152b01cac8daa80418906472268d08a61b8f9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-16b556f736a92391870ffcbd30ffc2724800c0f369389531c1e1d95b9b1c672a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79c63531f9cfa7c1f180ab841855ecf7c602dd4d53188e526194ad6271b5692f"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d7f85bb0de32 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-e595ac628b1b066ed2fb2d02a1b2583258a6995392bbdd6ea5753c90cb4195b5"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-b6d69784998e727171a51a02102404f791f490d2df45aeb606443ac91071284b"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d7f85bb0de32 / 3

<a id="canonical-f1660ecaf116c77ac2163d30b07470880720bba037a729c386240de8124dd1e0"></a>

<a id="canonical-c80f6e345ac589b7ef67bebef95596516cf10c259a2f2a58a47a5093f112a4b1"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d7f85bb0de32 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-9449ccae2f67874a05daebcf0d985a5682581c8eeeb156b602686828be2a26c5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d7f85bb0de32 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-5a023b29c475b7433a2ec321d5a152b01cac8daa80418906472268d08a61b8f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cecc5af753071ee96f8bd5669f9150354db4d0bdaff43d5ec42b8a92c6e4c16c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0b6b9fbc7706 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-20c91dd3c0a30d5bca734fe8119f34af51a03583cdeec41e1806b35c72d7e84c"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-5ed21368858b2a6885b0dc655f78d7e0b6debf989566ac0664abd244c3094c92"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0b6b9fbc7706 / 3

<a id="canonical-5d49670ae98a5110c82b751219d3b85d876e82171c088053a6d0955fae9b3ad5"></a>

<a id="canonical-59c64586c92aaf27f08762ff83b59e322ab101a8116704b75e77935adfc62652"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0b6b9fbc7706 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-e19c16f9c11a5b5ee8339b3e9691512af45aa3f55f42a61b73091649a7b83704"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0b6b9fbc7706 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-90456d493c2fc9c13cec997f4db6ccb8e9d37a512f03a03b5e6c649ffff84604)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-d148bd1433e60e3392ece674310a70603fea0ec3aa2b71ca5bdc0514bf3160ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68d7a809644152ecdf621bb3f38ed018c7707bcbc4beed8ebef42cf8157f62f5"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 04f594810ead / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-400e187c21a9f3f8ca85da9e39bf433115d897c373b4d62b5060283087ea8f82"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-05c7929d98aa6b7ed52ce018103f27368ede3a4770e76ec56faa1341f1b2718e"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 04f594810ead / 3

<a id="canonical-1f3c45ba1d2c8ec8febeb9f0413ee126b2e969395caf87f80723bf49c75aebd4"></a>

<a id="canonical-3f9c67839a8e08559388e992b9d1c893335615b3ad496da39260cb431253f317"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 04f594810ead / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-542e2a914a659ae666e088ba6673f80b115c4388327b303b4e3640280fb35f52"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 04f594810ead / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-1b12c82985b9367c25bed2cab761a3bb4fab9e057914e3761be5ece4a70e7b90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb341079873f78fac34fcc307748bd57d7f56221396b9823ce4ca95644e94158"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d913dc77e6c5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3922a6ec75961968ffe13fb9fe91d437828cefe8426e9f795c8ca5ffbb999bfc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-482dfa21a591c0f15d9bcedd5fe8799b119464862fcca4ef7639fbe51b0ee417"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-b9239d4d33fc833b13d53486b14a87517e94c3ca2bbbb261a1bccc4740833009"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d913dc77e6c5 / 3

<a id="canonical-368a1944df5047a9f961020b4d56bca54434c762a130cacf5fd21ce5d72e411c"></a>

<a id="canonical-69a504b79f8bed2bd48d38ffb02949cada251ae5d69cfa31010b1aef42becc3d"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d913dc77e6c5 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-5e4f3149e21a02e2473904d0f8bc6992ea97aab0c1d6486d94a48f00bcad172f"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d913dc77e6c5 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-06ca668faf612c5e1f3cbcce76663568c6f21a1c7aa8919546019c9e9b042ce3)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-ab4960ec98ac4745f5a22535394b369c98b66af037cb31902f6244c74785c0e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f4853df26e5855174e3c7ea7f6b5049a95e09238106330281be3a6fc15da46b"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7c94c07c3deb / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-4b7abb8186e32a0e625819ae39950addc60585b04a250492dd5a1239de3d5a21"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-921f7de0c6e7195298aa2824f35b32d2acc701c16c48f099e26a2f36398873f7"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7c94c07c3deb / 3

- [ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-40f6e237cd9f12a49a8302a0ed70c3941f1bdc071b8a2e8a8ebeb47f63ad0506): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-9978278f407bc25501900930b98187bb96bc2ed0dbb5b3d25709a78d51a95722): complete subsection reference.

<a id="canonical-026960860530ee57fd12d76dc9d62a1f16f0712bba7e55792471075527aea009"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7c94c07c3deb / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-40f6e237cd9f12a49a8302a0ed70c3941f1bdc071b8a2e8a8ebeb47f63ad0506)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-9978278f407bc25501900930b98187bb96bc2ed0dbb5b3d25709a78d51a95722)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-40f6e237cd9f12a49a8302a0ed70c3941f1bdc071b8a2e8a8ebeb47f63ad0506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d19558265365f571f9241d23b5b754c1ee6c6eec9001cfdd2f891587a7f51c6"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / c8c92a252dd3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ab4960ec98ac4745f5a22535394b369c98b66af037cb31902f6244c74785c0e0)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-6408d232945f1227213dcc82c1e6dceb9394b830f7597d72ed81ab58fd58b89a"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="canonical-d2e1f7fcc8e2b9cc7c38e1a677b693438715412f5410bb08d274a86bdd8d87df"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / c8c92a252dd3 / 3

<a id="canonical-2b36ee5f56beba20d8e2731a89b40677e18ae8c79b67f9c9f9bb193d04a9d696"></a>

<a id="canonical-63f1291e6996d2d39ac11f3c4f8ec4906f28e8ded1cc9aeab580ec94117f3f84"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / c8c92a252dd3 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
    }
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

<a id="canonical-a5c51127f2d6a58c9522bb14379f2e825f5b2951b84ff49482a14ab1f5b6ef5a"></a>

<a id="canonical-c03a16492936832b89035af8cb82bd62faca4d800c246517b9d7d8bbe57b10b7"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / c8c92a252dd3 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-bfcf9da99bfc4f6ab88b933998dc481974216a455fc4ba09ecddc2aa802c4296"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / c8c92a252dd3 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ab4960ec98ac4745f5a22535394b369c98b66af037cb31902f6244c74785c0e0)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-9978278f407bc25501900930b98187bb96bc2ed0dbb5b3d25709a78d51a95722"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4895c8b535886efee49bec6c202c65288943604a87012e6ea1369c658e70c96"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9487f2d66611 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-73f32d9e59fb930eeec21e3a87badc4540cdb285187f64c7e62ce3a5c659d519)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fc8b3e49200aaa55f865f16ab0aa292d1e0d8ef155f323ce0f76d29f65b913c7)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ab4960ec98ac4745f5a22535394b369c98b66af037cb31902f6244c74785c0e0)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-0f3c5531a7991a5129f7691bbef199e59eff8db9a7e6ff24069659fd67d710c1"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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

<a id="canonical-d5a7de5f64488b0a41118c77d865db8cf6e576abc16362f56341cfa47afa6716"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9487f2d66611 / 3

<a id="canonical-b816fbcbc551dd01bc15e87094c9739570f695db9045116b1f89134ffa290ab1"></a>

<a id="canonical-f4b4a00a6d532019be88d4e3033538716228f98e07ba83bd7b83b087a30f75f2"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9487f2d66611 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-fa7ed146e7908f04ce5eb09648c789f9baf18fdaf49a4d95c28951803df814da"></a>

<a id="canonical-d5cc0a4a7be60f9ffe92114d21e85589e70bb864f5a8c3a94505405826b12091"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9487f2d66611 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-2312eb68fb17f02123dc54a972333a13ba8a1cb1a52581b26f25fb5d70963352"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9487f2d66611 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ab4960ec98ac4745f5a22535394b369c98b66af037cb31902f6244c74785c0e0)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-802f7efa9c5bb45831b5499a2b8a18b7fe390fbca636525dfcfc208d85614e79"></a>

## ingress_egress_gw.outside_subnet — ingress_egress_gw.outside_subnet / cc48b711b396 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- ingress_egress_gw.outside_subnet

<a id="canonical-f8cd88576dd543f6e6c4635eec228efab99d1fe1e00133dd6d9ac53cc015f164"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

<a id="canonical-39d8d22b38f5e5fb1e582793864d93b229f8344a3728163c45d86e893fd9a13b"></a>

## Direct properties — ingress_egress_gw.outside_subnet / cc48b711b396 / 3

- [existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-38585b8c4bfe8fa0167cb2c6bc5105af69484dc6c8633d15cbe8946cf4f43ff1): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3e36d2fc87e57204c3c15bf4335e7e5559c2934fcb85557e47662e40f19de700): complete subsection reference.

<a id="canonical-9053148d176fde34037bdbf168e6947fc19db83dbebb33b9cd620ee0be7f612a"></a>

## Next pages — ingress_egress_gw.outside_subnet / cc48b711b396 / 4

- [ingress_egress_gw.outside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-38585b8c4bfe8fa0167cb2c6bc5105af69484dc6c8633d15cbe8946cf4f43ff1)
- [ingress_egress_gw.outside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3e36d2fc87e57204c3c15bf4335e7e5559c2934fcb85557e47662e40f19de700)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-38585b8c4bfe8fa0167cb2c6bc5105af69484dc6c8633d15cbe8946cf4f43ff1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-729dda65ae89b41b1ae7aa8ca93c7e86609737b7eda52c11915163022779ed59"></a>

## ingress_egress_gw.outside_subnet.existing_subnet — ingress_egress_gw.outside_subnet.existing_subnet / dfaefbc63aa2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4)
- ingress_egress_gw.outside_subnet.existing_subnet

<a id="canonical-2e590eb4c0b9a99d53d427f12f43a990fdfaf2c317e74ef2a4f9bb5ec2fe439a"></a>

Type: `"single"`. Computed.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

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

<a id="canonical-11ace57f67104c3b5b64ed2b0390e906bcdfe83ad1b57a68ff9d6ea6067002c5"></a>

## Direct properties — ingress_egress_gw.outside_subnet.existing_subnet / dfaefbc63aa2 / 3

<a id="canonical-8575e574ac51d637ba5eafceb2953770f71c69f51f4c000f202ee8335950ae0c"></a>

<a id="canonical-ec7bfa21bb6693f2c964f985a50599c8c156bf9c5fc6f911906e293c3141e864"></a>

## subnet_name property — ingress_egress_gw.outside_subnet.existing_subnet / dfaefbc63aa2 / 4

Type: `"string"`. Computed.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-74539f011311194d35f65fae5d457e9a11fae846c7425888371e34b2a42f6f1b"></a>

## Next pages — ingress_egress_gw.outside_subnet.existing_subnet / dfaefbc63aa2 / 5

- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-3e36d2fc87e57204c3c15bf4335e7e5559c2934fcb85557e47662e40f19de700"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-925362b33fe40de53411b2f3c44fbe57ab38d0d3b048cf735a4e62654339c5c8"></a>

## ingress_egress_gw.outside_subnet.new_subnet — ingress_egress_gw.outside_subnet.new_subnet / 69bdb42b17ef / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4)
- ingress_egress_gw.outside_subnet.new_subnet

<a id="canonical-63c30f2a5257aa53764f388468b82d885bdef7b154dbfee063bdae593ad3b120"></a>

Type: `"single"`. Computed.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

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

<a id="canonical-d2f7fca6f9362b380cdbb13dc353a862e79933dc9d05d13ae7e9485078831005"></a>

## Direct properties — ingress_egress_gw.outside_subnet.new_subnet / 69bdb42b17ef / 3

<a id="canonical-51f8c2d20891b9172f8cec6ebb67d09a744570d48d806fe6a77d2bc6de65c8b0"></a>

<a id="canonical-d225a75269395f74b0a46c1a34595099a08bf1f89f1cb0f73e909dddd4536764"></a>

## primary_ipv4 property — ingress_egress_gw.outside_subnet.new_subnet / 69bdb42b17ef / 4

Type: `"string"`. Computed.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-00ee8bcfd545c7c7fb1016886ab0018cf89875a1ad9b32e58a9a46e27646e689"></a>

<a id="canonical-a0e1e0f7320c159b1525c6be954b26f4c78982863a6b9680caa6bdb7b0e26a4d"></a>

## subnet_name property — ingress_egress_gw.outside_subnet.new_subnet / 69bdb42b17ef / 5

Type: `"string"`. Computed.

Name of new VPC Subnet, will be autogenerated if empty.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-ee7c7b9abd774db56502049404302bd91a52cdbbcb403064d8a233d787a7a617"></a>

## Next pages — ingress_egress_gw.outside_subnet.new_subnet / 69bdb42b17ef / 6

- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-480eff86e53d2e159619eb644bb4394528311341725633046b67c75ec2b504fe"></a>

## ingress_egress_gw.performance_enhancement_mode — ingress_egress_gw.performance_enhancement_mode / e9bd3ddb67fa / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-e2dd85b81d5ea2aa318efca5076e763ee383a0d30b324d42cfc738f5a9dc90f4"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-dcc02ebf69254fa94c07934c99450c76190aa3feb6b0218b75782f2755cdf92d"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode / e9bd3ddb67fa / 3

- [perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95): complete subsection reference.

<a id="canonical-41716458a81d0f5120259bbfe5363d70f488f83fe3ceee398f71ca9e7fb04b05"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode / e9bd3ddb67fa / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-904226a3402866382f8b08209e8185f4a1c9a3c328f023c76793da22b50dd436"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3e7fc4965ba8 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-385cd85551365076c4afd174713bbe375dd44f887c52c7fa9c809e4da4d2bdfa"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-d0f8e46744f5b56a65c2d47f74a130739b890b98a158ed68b594e5d1ea8a9198"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3e7fc4965ba8 / 3

- [jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-00856f87f21319b55a44a2f3267faa21ef0d1c431769865882804eefacbc8abf): complete subsection reference.

- [no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-791d6d9263661043982623c53a0938cf4fb3a91708b8e25986160aebe6448fd2): complete subsection reference.

<a id="canonical-3ec79f9bccee4f7c4609342e8c947d3213c8053c0ac3ab3b4f9c3b237c739192"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3e7fc4965ba8 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-00856f87f21319b55a44a2f3267faa21ef0d1c431769865882804eefacbc8abf)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-791d6d9263661043982623c53a0938cf4fb3a91708b8e25986160aebe6448fd2)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-00856f87f21319b55a44a2f3267faa21ef0d1c431769865882804eefacbc8abf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4d108eff76911a9eeb7c2df4be3f27f7228a4cfa3c1b8345166461eb9374b9b"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 6b0481c81f06 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-96cfe6d2ff2be8a0192c64f01b122efabd421d9d02d17635aefafdacd7edfdaf"></a>

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

<a id="canonical-0c8a5ff23a67d317c0b0d06a6ed1f4e646f135ef7dd458c56da669031ecbe2fc"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 6b0481c81f06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc006adb9ea56ec98ae6a5f4ae37c98a55f4d0f6113453282ef8c35ea25ba718"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 6b0481c81f06 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-791d6d9263661043982623c53a0938cf4fb3a91708b8e25986160aebe6448fd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-621f22242a92c1c14aaff60dd98ba85e92f4ec8d49855332716316b86e1dd9e1"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4a6ea44ea3bf / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-e1626ec32003f9c77147036a69112cb126d027da9357a9e86f34c1c70b953002"></a>

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

<a id="canonical-1fd4d037ba77c5cdc169f9d5fa4e4eb2bcde483a33e2bf191609766f1c854b59"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4a6ea44ea3bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a6deaf00a419beca3fff2300e2ff53f00c53e6b1005350a16a9af721c04c5b56"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4a6ea44ea3bf / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-58c8da39dc2efab314c03fc508e8b5df3dbddb7f5d1fde74e9f10a23d755fe60)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f83703644cc428dfdd62059405544d46ff391ba13e87eede5280b749b42a1b7f"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 40c6754ec7f2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0df98a05d13cf6e16b6b1b290efb96e2a8c20ee6e83ab51ce5827fdee57f70e6"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-3a4d4758b533059c6dd6e8e5fbe6330fba2f7572c50c10afb8d9186cc4619d9e"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 40c6754ec7f2 / 3

- [jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c5731457fed0f431f1f6aedb3c6369d5a7b3b5363c750e24da2f51e496bf0767): complete subsection reference.

- [jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-348c97ec0b59efe790c78314fcba3c0117b67e393ce1eb30836d6cb3198bc265): complete subsection reference.

<a id="canonical-11f26febd6b8280f6e05afe5c2621aa9cbaba6d93becfa1f217fe8e7892c662f"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 40c6754ec7f2 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c5731457fed0f431f1f6aedb3c6369d5a7b3b5363c750e24da2f51e496bf0767)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-348c97ec0b59efe790c78314fcba3c0117b67e393ce1eb30836d6cb3198bc265)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-c5731457fed0f431f1f6aedb3c6369d5a7b3b5363c750e24da2f51e496bf0767"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9731aa75db1f14d675e0d2f8d34fdab98e0e5c34dbc18f83cd980ec2550b52a4"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 5383400994dc / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-793e6bbdc9a8f4dbfb53752d393b830adf9f637da98f63f57c2f9912a77b7dd8"></a>

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

<a id="canonical-e299d0d5b18615760d977e3946d991b3ab8db719941541308bbe75332d4b2bb8"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 5383400994dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dabf5f62dc820f927ffc20e62c2c5d7a36ec1ee10162ba90d0792b71bdf89fa8"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 5383400994dc / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-348c97ec0b59efe790c78314fcba3c0117b67e393ce1eb30836d6cb3198bc265"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56f9357001f2e7d07378f75b2b2d580d2dfe30df32699147119d087964c41dab"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 75cc28bdbb6c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1471b51cee8f9bd043a7157daa97c12b14e8cb9e32a1e37c6320e7123a9f2aef"></a>

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

<a id="canonical-273a85ffb2eeaace19d7f6b07db24ace73c8cf758a41a5c284c1778644f5b840"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 75cc28bdbb6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9308744a40397f63d4a7d7d57b9c8ebaa9aacbb27b06d2a52f174e0683aedee"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 75cc28bdbb6c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1a8557e3dc9924063a9a34595704af65f59af08122bbc2d25c9ddc282fe30f95)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a77b1430d70c66150cc8dd053e3b84d88666df8becf3e78c28282070d1905aaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6238cd6ea2b1aae1244d509fae64068747b62f485c29eb05793e80dd15bfdc9"></a>

## ingress_egress_gw.sm_connection_public_ip — ingress_egress_gw.sm_connection_public_ip / 8ad0b4d23c10 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-495f375688a6c5b31a9672585e65b2262cf4989e12c71c3f1891ed570e5fe520"></a>

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

<a id="canonical-a9a5dd6bf27eb4153c7de4274b91774995058e0d65c459acc6491c68f5027011"></a>

## Direct properties — ingress_egress_gw.sm_connection_public_ip / 8ad0b4d23c10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e98898b3a7092e197ebab2e6f9788e145042cd8bca739260958db8f0e65db207"></a>

## Next pages — ingress_egress_gw.sm_connection_public_ip / 8ad0b4d23c10 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-6c2d5083214b37795bbea7501148b070fa5c5a5a70763d4879ba4cf47cdad721"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d73a10b4712a09dde00369466232dbf158c9f515aecc49c317f58b359c4fdb73"></a>

## ingress_egress_gw.sm_connection_pvt_ip — ingress_egress_gw.sm_connection_pvt_ip / db6ed805aa3a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-c7e967a8c5731a941213c6e7b2b5406eab34c033fec9d9d3ecc151c328aeed15"></a>

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

<a id="canonical-c55e7edbe616d785f2e13bedef54789b5d6924ed14bf622fc1616035036b5af3"></a>

## Direct properties — ingress_egress_gw.sm_connection_pvt_ip / db6ed805aa3a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a8c8813a202de6ba96f43c2153d4f55374e0b4cc93f188afe42ce6b0f0b670c"></a>

## Next pages — ingress_egress_gw.sm_connection_pvt_ip / db6ed805aa3a / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96e4a70b2b0df1328384df4eb317bcd7b6ad23e2fa5217b3405aa9844fdd401c"></a>

## ingress_gw — ingress_gw / a2a3c0f33cd8 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- ingress_gw

<a id="canonical-aabb0996e5fa997bf2ff3a50da2ecedf568fefa97b3dfc132843c1eefe4bc0dc"></a>

Type: `"single"`. Computed.

GCP Ingress Gateway. Single interface GCP ingress site.

Upstream description:

Single interface GCP ingress site.

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

<a id="canonical-ce651eb3e602f0e5b27525b0400ec619b89df3932e941ffc47e457fefcb5e8bd"></a>

## Direct properties — ingress_gw / a2a3c0f33cd8 / 3

<a id="canonical-4ca378f5cc102c5d5eb3e2769f111e397f76449a2e4f07aff46258668df58cd4"></a>

<a id="canonical-d692f0e0c34472e58be00ad50c9c12bd1abeb077079a90cb70fbb1a94a1bcea6"></a>

## gcp_certified_hw property — ingress_gw / a2a3c0f33cd8 / 4

Type: `"string"`. Computed.

\[Enum: gcp-byol-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-18b44f50ba201e42d866cf1ac629e82305911431b2017fc995aa3ed6019ce17a"></a>

<a id="canonical-f134dd338406134e64d73980a38a6a5e2206defcfb2d2147ce21105e32989b73"></a>

## gcp_zone_names property — ingress_gw / a2a3c0f33cd8 / 5

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4): complete subsection reference.

- [local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2): complete subsection reference.

<a id="canonical-c90b4e62337c696962c069c38f1a68cd5521dda9b2fd100b0e5fefd11b8c05f5"></a>

<a id="canonical-9f7dbe93e144443aa1377abc2474d9fa270654a8d06be01d1c681b8ad57eec8a"></a>

## node_number property — ingress_gw / a2a3c0f33cd8 / 6

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc): complete subsection reference.

<a id="canonical-2986e339e71f52d8a4d86cceecb75a7e2a1d608165dc261c7a795425e9dd1ad6"></a>

## Next pages — ingress_gw / a2a3c0f33cd8 / 7

- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7e26fe4149606ef3e2d53f5bd9439cef677ed59fbef5c3ece92f8a5101afb90"></a>

## ingress_gw.local_network — ingress_gw.local_network / 29bd8b298748 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- ingress_gw.local_network

<a id="canonical-b025cdbf474731cd9ba166b29cd99553722a0d7260bbe247c24e52a655f5929f"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

<a id="canonical-53007cba53818087fbf54822d21f4c41f4c6161c05dea32249e3e98e7fb1315d"></a>

## Direct properties — ingress_gw.local_network / 29bd8b298748 / 3

- [existing_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c795a0270005580bc6ace2dfd5bfa23a32a787ef58901a08531cc818d9eeb7b5): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f50697a12b48649b4b8d83fc589e9581f24c097a3b5c7140a1d441a80df64976): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-003.md#canonical-9ac8be679a1ef797daa899c8098a84d991e28dfc7650475b4eab3d13812080c4): complete subsection reference.

<a id="canonical-88ad98da53f227e399dee9a385f14feb29c6f75a718a62f001e69a6f47784091"></a>

## Next pages — ingress_gw.local_network / 29bd8b298748 / 4

- [ingress_gw.local_network.existing_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c795a0270005580bc6ace2dfd5bfa23a32a787ef58901a08531cc818d9eeb7b5)
- [ingress_gw.local_network.new_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f50697a12b48649b4b8d83fc589e9581f24c097a3b5c7140a1d441a80df64976)
- [ingress_gw.local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-003.md#canonical-9ac8be679a1ef797daa899c8098a84d991e28dfc7650475b4eab3d13812080c4)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-c795a0270005580bc6ace2dfd5bfa23a32a787ef58901a08531cc818d9eeb7b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b4f0ab864d7a04a06d39b96b915b8cc1539beba22d9e33399646badaedf72af"></a>

## ingress_gw.local_network.existing_network — ingress_gw.local_network.existing_network / 3ce6ba630002 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- ingress_gw.local_network.existing_network

<a id="canonical-4ab8c1fd4888c773bcf535b66259a1868d2d1bb72fad4420a514fb997fd2444a"></a>

Type: `"single"`. Computed.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

<a id="canonical-5aeaf229dd745e9d668c70e3d09bce2b1d06ce74926512c5689f0e230217ea7b"></a>

## Direct properties — ingress_gw.local_network.existing_network / 3ce6ba630002 / 3

<a id="canonical-534b9d23c7eb7db66df8d672533ab22911c9d19f7c236f28363d9cf2126c5b37"></a>

<a id="canonical-f0a3e6a2e4670843ba19bea496537185bb7bd72aa4fd15a7413c1878d8db8ab7"></a>

## name property — ingress_gw.local_network.existing_network / 3ce6ba630002 / 4

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2ade7b93bb8966f91224671b493322366a9f6523591fda1d6dee52e548d81b12"></a>

## Next pages — ingress_gw.local_network.existing_network / 3ce6ba630002 / 5

- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-f50697a12b48649b4b8d83fc589e9581f24c097a3b5c7140a1d441a80df64976"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29fd054d09455d9c5912757a7c6cf376fa5738ba37eb646bf420b603b093a1a0"></a>

## ingress_gw.local_network.new_network — ingress_gw.local_network.new_network / 84e5190a5b8c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- ingress_gw.local_network.new_network

<a id="canonical-4b7e1c8124d8ec4fa29192d47a25f4248266b7be3d24a2f59255c64d5ebaa241"></a>

Type: `"single"`. Computed.

Parameters to create a new GCP VPC Network.

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

<a id="canonical-b7451c016d07969d7fb973588e6bc66a8aad89f4c45cd05905911bb2300e3bf9"></a>

## Direct properties — ingress_gw.local_network.new_network / 84e5190a5b8c / 3

<a id="canonical-c5f131d442b44b489f95f2ee24e92c52233849511846b8f3f1388cc645af546a"></a>

<a id="canonical-0e7f1738df8530d44402fce17f4c6a5636ec4894e76fc66b1fa0f459fc033a1b"></a>

## name property — ingress_gw.local_network.new_network / 84e5190a5b8c / 4

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-e11036c68ef0a2af8015d911a01a35f6eff74a88ac71bc2e8edb8a880a276ddf"></a>

## Next pages — ingress_gw.local_network.new_network / 84e5190a5b8c / 5

- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-9ac8be679a1ef797daa899c8098a84d991e28dfc7650475b4eab3d13812080c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c6e4d159f728c1e6d1176292fad1281676c4985507bae0fa1e534628e1905fb"></a>

## ingress_gw.local_network.new_network_autogenerate — ingress_gw.local_network.new_network_autogenerate / 537304bbf789 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- ingress_gw.local_network.new_network_autogenerate

<a id="canonical-564b9cc8c81b6e2693d91afe2c7ce4f4a0aef152b0d28cf1c35326ba0e8cd97a"></a>

Type: `["object", {}]`. Computed.

Create a new GCP VPC Network with autogenerated name.

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

<a id="canonical-f6e406a3daf38ff539651332d10b3c4f2063c9e05197487be999318f0f8e1875"></a>

## Direct properties — ingress_gw.local_network.new_network_autogenerate / 537304bbf789 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a814512abde33328a3987d5829b9e7d9aa9b8cfc79a600c28a85afaa78a07bd3"></a>

## Next pages — ingress_gw.local_network.new_network_autogenerate / 537304bbf789 / 4

- [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-420f22b7892b11ded321aa8fd9a2dbe07774886158ffc1f960a59bd4b82ed3c4)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02cd5a7f70903809fc2a6fec84386576767be602a0a6ee604edd64fc460e458b"></a>

## ingress_gw.local_subnet — ingress_gw.local_subnet / 9a93b3df7883 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- ingress_gw.local_subnet

<a id="canonical-69fe675f863b2c9706dc2dc1e81a0c909db1bfceb5fb4771bbf3b23635a4022d"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

<a id="canonical-144546ecee2eda2da3d6534ee98d43b6ea164f747869b024dffb101151dade76"></a>

## Direct properties — ingress_gw.local_subnet / 9a93b3df7883 / 3

- [existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5a6a222bb1ca8cbdab09c1ea10f19893e4a08d99f2298f43297dc6af0c508cdc): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-681e3cdfded6bd7cec1ce264f352a5bb28bf63a437e7f3a4884ec061cf74b9e1): complete subsection reference.

<a id="canonical-64cdd0ac20ca876d91ca56fbea13c6cfc84870bce1d417d47b093a400c1b2ee0"></a>

## Next pages — ingress_gw.local_subnet / 9a93b3df7883 / 4

- [ingress_gw.local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5a6a222bb1ca8cbdab09c1ea10f19893e4a08d99f2298f43297dc6af0c508cdc)
- [ingress_gw.local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-681e3cdfded6bd7cec1ce264f352a5bb28bf63a437e7f3a4884ec061cf74b9e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-5a6a222bb1ca8cbdab09c1ea10f19893e4a08d99f2298f43297dc6af0c508cdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebce5e4ed57b9cccae7655aec7990e94aed49f4418b0e0477257b56ae4e97674"></a>

## ingress_gw.local_subnet.existing_subnet — ingress_gw.local_subnet.existing_subnet / 818fa8721e52 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2)
- ingress_gw.local_subnet.existing_subnet

<a id="canonical-652bf7196b5762a019a27c826fda740cb24210d768ce384b4147511380e331b9"></a>

Type: `"single"`. Computed.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

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

<a id="canonical-9b8b050a6e4ec5fe45b05af7e5bd5a1c15962cdbae68c893580cce1a50a35376"></a>

## Direct properties — ingress_gw.local_subnet.existing_subnet / 818fa8721e52 / 3

<a id="canonical-c242c72be6af2f8fa94148a4ce554cbd8177ac694218e0f854cbf945572d7325"></a>

<a id="canonical-859f111ee2e6609882ee6631a827cf4f9bd923ab326b3d99d85dea1d2e85f7e4"></a>

## subnet_name property — ingress_gw.local_subnet.existing_subnet / 818fa8721e52 / 4

Type: `"string"`. Computed.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-5f4c1f7a55570bedc3fdd7997f6a7079caf612c811e888e9770b4cff49451c65"></a>

## Next pages — ingress_gw.local_subnet.existing_subnet / 818fa8721e52 / 5

- [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-681e3cdfded6bd7cec1ce264f352a5bb28bf63a437e7f3a4884ec061cf74b9e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5353650ef23d4bca6dc08ac0d1db01b096c7417aa461eaa386ce3c5d5356ab1"></a>

## ingress_gw.local_subnet.new_subnet — ingress_gw.local_subnet.new_subnet / ccda3f75073a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2)
- ingress_gw.local_subnet.new_subnet

<a id="canonical-1193621c01047293f2c50ad9eb6485619935f0423054050fc805f37cf9416faf"></a>

Type: `"single"`. Computed.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

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

<a id="canonical-42ef9b3fc2341cdda23810fdfc6c7e595a5373fc1c7766a567d7f0e94f95ea5c"></a>

## Direct properties — ingress_gw.local_subnet.new_subnet / ccda3f75073a / 3

<a id="canonical-da43d4368b225a9b835f1155f5e2ce0f251f1baa11faa130262bb86de4e4c86d"></a>

<a id="canonical-f3dfc66d5dd18d105208d84928bc3438884fb185bfb0c33f908f6b664f2da0d2"></a>

## primary_ipv4 property — ingress_gw.local_subnet.new_subnet / ccda3f75073a / 4

Type: `"string"`. Computed.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-8f7daf3074b1c812b866a26850c233a6fe2d137192dd37161cd0cd3d30b65b8a"></a>

<a id="canonical-562c56916bf7b10c94c226eb298648d72b9cbf395a09a71d48dc8418ebab072d"></a>

## subnet_name property — ingress_gw.local_subnet.new_subnet / ccda3f75073a / 5

Type: `"string"`. Computed.

Name of new VPC Subnet, will be autogenerated if empty.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f84c467c1742876315fc4a01252addfa21848bcab2e5d0bff5c2935528ef2680"></a>

## Next pages — ingress_gw.local_subnet.new_subnet / ccda3f75073a / 6

- [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b6634fec35e01f65cd187b85888e9c2bbf8ae656da49502478f441b3ab4384e2)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1dde984dce8acc7e193d3e9ee800d7a5f2f0bb6b92683265c4e9bbd18b6dfbf"></a>

## ingress_gw.performance_enhancement_mode — ingress_gw.performance_enhancement_mode / b6e6c6ba7c99 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- ingress_gw.performance_enhancement_mode

<a id="canonical-4784c4b989e0feec8b060049ea862eaf9d3383da8ec72eee744e1578082e6a91"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-52d2da7a6b1829e024d52f6b122107e9371cbe3bf54e6777fa6fc83dae9ff542"></a>

## Direct properties — ingress_gw.performance_enhancement_mode / b6e6c6ba7c99 / 3

- [perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12): complete subsection reference.

<a id="canonical-47bdff80a0fac46ac8e52ef8b973ecb6a1d75a1d46c12c7fdf0bdefafa8a83d5"></a>

## Next pages — ingress_gw.performance_enhancement_mode / b6e6c6ba7c99 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e168758d646e6f6c5686b9c65722eeacca21863446645b6e5f2a3462933860b"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 000a9d178bf2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-4066cdbc71961c565ee75a558ed7189ed2af27656ba89da06a089ca0d4cd067b"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-000546f00550205b431a2b50670a51c4f472160c2f506059a2258d0021e95a5e"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 000a9d178bf2 / 3

- [jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b502c992c20f33e74657f8d70e2730996bbefcedbc712fec592e8531ad8c4359): complete subsection reference.

- [no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fc65b93812548445cb8220625ec800e4997b31c8e7fd47ffe5f0cbef1a7d846a): complete subsection reference.

<a id="canonical-3517448bf4a27ff5c29fb4647e322164a24e0b080be28f97f9e9a25b7899ed05"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 000a9d178bf2 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b502c992c20f33e74657f8d70e2730996bbefcedbc712fec592e8531ad8c4359)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fc65b93812548445cb8220625ec800e4997b31c8e7fd47ffe5f0cbef1a7d846a)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-b502c992c20f33e74657f8d70e2730996bbefcedbc712fec592e8531ad8c4359"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f0ab9f33d8551d0315eaab385f419a7e84c841bc3f3d0fa5ea1cb2b3837e337"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 897bfda5ad8d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-c50a382c4cf8027d1cf5060ccb70c52f361f0e33e078d83d32dd390347e44ca1"></a>

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

<a id="canonical-64bbc38b469060c1ae25c021d457f080baf52b050490679a12fb7fcb446cf2c7"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 897bfda5ad8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4840b52c20bcc6b44706092c82572776ca72b6402b51e44aa1758cb915373a1"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 897bfda5ad8d / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-fc65b93812548445cb8220625ec800e4997b31c8e7fd47ffe5f0cbef1a7d846a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78eaee943136dfbb40c37de78917deb22b1a6678409db4a4eebf2cf6350728e7"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4c5e59754a1e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-924dd473405c42b6fa8415572ad42742f032dd5ab40802a34f5422c3a1815ed1"></a>

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

<a id="canonical-88d0789e90f032c107cb7fb7be0d72cfc4ae0cc6f5d41461a1be2c4de7b3dcc6"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4c5e59754a1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6479bbb2d6f8cf76f154755feea7640e6a8c5b338372062395ed7c21a2dd625"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 4c5e59754a1e / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a97b388ee2dadcd1ed9f0b8a29db44587b6ebb22e04dda719a0b08d3992cf4a0)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cb14e28ff937d9d16a05b4efc263d81ce5b8a0d187834b8eb5c6f641ccdb7ed"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 630c1af942c9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-3783daa7c50d7b9be1a08bb5fb6426a6b32781d544732f9fc6d48d98dd3cc0dc"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-803edfeeb985eedf8f692f3003957a5759739ba52943a1e075d3a1c7fe5c6ab4"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 630c1af942c9 / 3

- [jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-88e0a92713cdb3a7717c54ecb8f197ce1c7bee8ae0ebe111784a26979d033667): complete subsection reference.

- [jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a805381e28690ae9c2980a599dc4b7162440b087c7d4169eccf403474d8c3c97): complete subsection reference.

<a id="canonical-b6ea154bb38a5146f1092d9b5c2c3e76ff4d62794a187e2d48762bb1e31b625b"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 630c1af942c9 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-88e0a92713cdb3a7717c54ecb8f197ce1c7bee8ae0ebe111784a26979d033667)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a805381e28690ae9c2980a599dc4b7162440b087c7d4169eccf403474d8c3c97)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-88e0a92713cdb3a7717c54ecb8f197ce1c7bee8ae0ebe111784a26979d033667"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b5b6249998deee6458afdf2da6a545571217266968de969a4a460aaa6c2ddf6"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b676e2452931 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-c4482003092e40fc90ba0d777308ea875138be775fa8034e9bd965895d12cd61"></a>

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

<a id="canonical-caec9398e2ff7de8d4824d4f064c4ba4f70fb8cfa99c3f6b5daa688a009a2239"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b676e2452931 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86a62dfa63ca8fa0d33246da601b250f911b7fbbf20a1f7b03c79aec405a9748"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b676e2452931 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a805381e28690ae9c2980a599dc4b7162440b087c7d4169eccf403474d8c3c97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-680b44c2432caf7d2c872684c5bb73b2b906ace8d77529d2953ccc8d7ee571fa"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / aab0f8dd9f5b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ef7b69a8befd16bf9aeca9338d318aea803fbce9ada832c66d542e7209fc79cc)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-adc78edef91a1e59df0b398c03a709fe80ffaf17fc7d78bb3a3689e8a31fec4c"></a>

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

<a id="canonical-92f82cbdcd194ca901d309356637e0b0d683fa9f9f85bc8a64102ec66b9401d3"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / aab0f8dd9f5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e935292cadc76ea068d99b919be97715373017fc79888d2035cc229920fc0a0"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / aab0f8dd9f5b / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e4daf4928402078ccfe3da31aa9563cdec6c0d395910097b9621276174490c12)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6622cc5b27ae0dcfa269733806dfeb699d68d2f106adee64809316315d4fb62"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 75a8c4cbac04 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- kubernetes_upgrade_drain

<a id="canonical-0268e772a08896abca785d965a1fd0d7d6a1eb88536dcbe6988097065a964f34"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

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

<a id="canonical-27593b7010aaa74edf2330aed4ea60dbf90c1284d1a4863e50bc866a7ae4e633"></a>

## Direct properties — kubernetes_upgrade_drain / 75a8c4cbac04 / 3

- [disable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f5da688f8c746db3ec3887a51035c617b867dc3b0f45e17e6e55d7ee409ea040): complete subsection reference.

- [enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5): complete subsection reference.

<a id="canonical-5debb1c8299e42a0a4c4975622f3d89a2d9d6d3b60e1f1c8c14455a1d73e567f"></a>

## Next pages — kubernetes_upgrade_drain / 75a8c4cbac04 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f5da688f8c746db3ec3887a51035c617b867dc3b0f45e17e6e55d7ee409ea040)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-f5da688f8c746db3ec3887a51035c617b867dc3b0f45e17e6e55d7ee409ea040"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37da5e84979d2bea1263b0bdd503f8bb9f6d4aa761ae4ed8632d6c000d3dc984"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / b2e6175e2706 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-f159937ac70275d80158bfe5807991339f3ae51fc4a495674bcd50beb14f7bab"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fc5893a21ea7c971cbdb77fc49384d09385534f9f988629702636e139ef08917"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / b2e6175e2706 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-786e34ce1e76f223ee285a6e1833b612ecaef464455ba9c93d08ccef6c68eaaf"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / b2e6175e2706 / 4

- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e42527ec4ea7628684a14a0c1f328559bf5122dc7b0a5c666795c84e6bd5bfd"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-342fe8ee401299b451d973348f86b8b1ec97ab448830d5ff62d302029b356a85"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

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

<a id="canonical-d2df83c0f58ea70880d4e4689ad279fb708699609ed9f09ae0224cb27dce9118"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 3

- [disable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4eced3cff3d6c5eea1b3679490b858aa985f47d80542be5bf3d048bfb655014a): complete subsection reference.

<a id="canonical-9a9461b1a3fb747c8ec94a179e6a2c9d3958edf9d39b9d25767bd2d1168f6a81"></a>

<a id="canonical-cfeb811f06f3b5dc1d3aae4d6962e54841a91382ea20f687bc1adac2f46922c9"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 4

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

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

<a id="canonical-96e7b67e87ab5ea93261a30ca2d2a4a4eff9d87449c9e56e0a53bfcbf64eab63"></a>

<a id="canonical-377cba82d8183988778cb904b1c63b23933ed2c383fa63ac48825a4cf8528e34"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-40d830e91e3e58daf4b3b66d90363bcfabf4bcd28441f235b1017c29915d6eb6"></a>

<a id="canonical-9b4e8f24b6ab269bab769fff35814bfc834c1ed1cb650005ea34edb40339c5a6"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 6

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

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

- [enable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d290e1f1bf1a659a2a0621c4341226defb5f37d17e4becde5d0d4a0ad636f5f0): complete subsection reference.

<a id="canonical-d8f9cf1a2ca24455f5185c7f8ddb6699712307c2c5abd5961ff6ea63e1017421"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / d3a4e3662642 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4eced3cff3d6c5eea1b3679490b858aa985f47d80542be5bf3d048bfb655014a)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d290e1f1bf1a659a2a0621c4341226defb5f37d17e4becde5d0d4a0ad636f5f0)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-4eced3cff3d6c5eea1b3679490b858aa985f47d80542be5bf3d048bfb655014a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10a691c517eaa0568fd6f9414f16316b6068c2ac1fee456bb3070895e9af519b"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 94c0e045713c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-5d4a035171a30a1c5dfc3d58dd7ec5d5c66ac73cdb6fe89f55edf7f1c0f0be8b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-70facbc9e2db7e80d13b0a606743c1468f40cbcdb5244c51628da600cfbe4213"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 94c0e045713c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1902f5e0a07c671ae2f7db05020180e7537a262c28f16e8468570938bec00e74"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 94c0e045713c / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-d290e1f1bf1a659a2a0621c4341226defb5f37d17e4becde5d0d4a0ad636f5f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3e1ed330171c8f86136a0bc0959a2732b428445a29f19f45e5da17852ea004d"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 9cf12a3c4d4c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-ee928bd26fbe59cf2300eeff39543e8469959772a8c5a889d19ec7c2f4426283"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ea906bc4dee8991024816d5694f19ef9aba3294193870ad6d97c5a42f3a211d1"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 9cf12a3c4d4c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-56809d3d5e6656b1f1094886325adcd041439550301d41cc7cf79297e2c94762"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 9cf12a3c4d4c / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fb99eb03b6ce4b5b5d7f58921240b398ee1eec5a6624cbc9851c7b98237724d5)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-c62c46141ea3c05a0ddfe9ee1c3e3f5236aacb95d035b7d2379729dc9adaad66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-112cb07c80162ef6d125f4c3115bdb40bc5733aecb4dea9790bd0c88598fec4e"></a>

## log_receiver — log_receiver / 5bb832839f5b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- log_receiver

<a id="canonical-fdee6744d558d4e71fe3e34712018c2f7c35dcc1f9574e2c207e17cd9ef7943b"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fdee6744d558d4e71fe3e34712018c2f7c35dcc1f9574e2c207e17cd9ef7943b)
- [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1548d29124796eb96edf7e8c04c533c717c5f95b4cbf3898d0d80d3ba00c0cc1)

Select alternatives according to the provider validators above.

<a id="canonical-ebb6ddeeac4911d42b94c1bdb486a63e8444bacfe791865c5861ca5e4dc238d7"></a>

## Direct properties — log_receiver / 5bb832839f5b / 3

<a id="canonical-ec1c479950b739d137fe786a1fa3e6f6a9336e624d2723574e5d7efad1369c70"></a>

<a id="canonical-b0fbc123f81968b7615567129952e7e74b5ece6fef80a1265957763109e7c488"></a>

## name property — log_receiver / 5bb832839f5b / 4

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

<a id="canonical-6960d7beb0e2596a51c671730654f8d789faac138d9d0bc913fe06e79c23623e"></a>

<a id="canonical-2533980ac5980c0bb2d108b89da59bce1cfd557180c5a3dfece39b3cc22f7dc8"></a>

## namespace property — log_receiver / 5bb832839f5b / 5

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

<a id="canonical-d116bc4137e040463ae11fb35d3251c01b7a13fe9b0caa1e15d7fda06488a27a"></a>

<a id="canonical-1323ca64e9ad156a71d9e8045e505b081eb2508dc08cc11fb5bf0d10f155ea67"></a>

## tenant property — log_receiver / 5bb832839f5b / 6

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

<a id="canonical-7b60d7fbcd78322a5e0d30d978b27ffd8d9af7f89dbcc0c29884bcde2afec719"></a>

## Next pages — log_receiver / 5bb832839f5b / 7

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a22274d6f2b4e4aa01fd0485ae5fec471b2949c4b413dae12f16e8d498728f73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6e8f606df26c1b35842ace69eb416d8a2d46a886c79c493b3df2a7990e0c319"></a>

## logs_streaming_disabled — logs_streaming_disabled / 2fb966b20718 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- logs_streaming_disabled

<a id="canonical-1548d29124796eb96edf7e8c04c533c717c5f95b4cbf3898d0d80d3ba00c0cc1"></a>

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

<a id="canonical-c62ecbff67453165cb066a433a8e9e07cd15ed01827b2cb1fe99a3f99a444606"></a>

## Direct properties — logs_streaming_disabled / 2fb966b20718 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9616f2edd784f09c550a971d6a979d0699230a1bfa7427f9da3e611b33ede8a1"></a>

## Next pages — logs_streaming_disabled / 2fb966b20718 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2466c4656ffbf9b05d1cbd1f3c430062985bd198db5d54820bcbfefb1088145"></a>

## offline_survivability_mode — offline_survivability_mode / 51b6a60541ba / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- offline_survivability_mode

<a id="canonical-6de2af680f24fe7199319dfc3a733cc421376ba82a48b7fa0473c5758b4ea04d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-56665aa01f28dfcdad9b5be8cdc1b3cf35ea49509f539a744cdbaf37016a9ca3"></a>

## Direct properties — offline_survivability_mode / 51b6a60541ba / 3

- [enable_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c1bdf0d02ea8cb02e525c103ea3040b46c5fb56d66e9297263539d11b55bee84): complete subsection reference.

- [no_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-8825702bab6e620ac64b374a3fc5c799e935c6c6c57fe3e4b1a6e087c40207d0): complete subsection reference.

<a id="canonical-3227c01557e21d986494f5143396a24a9c3a57798f38167a5f51fdfc90b5fb4b"></a>

## Next pages — offline_survivability_mode / 51b6a60541ba / 4

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c1bdf0d02ea8cb02e525c103ea3040b46c5fb56d66e9297263539d11b55bee84)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-8825702bab6e620ac64b374a3fc5c799e935c6c6c57fe3e4b1a6e087c40207d0)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-c1bdf0d02ea8cb02e525c103ea3040b46c5fb56d66e9297263539d11b55bee84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e882f6407c849532436c992c5e3c419381543278bdf6efa616b58e85bfd9af8c"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / c3217cdccba6 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-efb7a9580421794aa919fe6bb3975a662c68b6267a222abcec69762734158b70"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f3879a94bb0b076f1ded1736b8ec3aea38aed1262277c429c1e00c1cf468ade3"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / c3217cdccba6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae03d90dbd9bd4a462bc8e872cede2d71f00647d7ca6117b06acf46cedf2dd0e"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / c3217cdccba6 / 4

- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-8825702bab6e620ac64b374a3fc5c799e935c6c6c57fe3e4b1a6e087c40207d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7395a77a816bf38db94b59091355d5720f0c72359e95067c3250321211c7cda7"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / b627782c7244 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-e9781e3460c28bd8d66fac4c8af31c886de8825d6bcbe2f96b9173b505c6949d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no offline survivability mode.

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

<a id="canonical-865c3867c52d5cd28d2f6106c3f4d494ef326697a5ae5cf741715c6319aadc94"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / b627782c7244 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a1b78462dd7a6813e24e869353d134e29304a8f2bbf3085e0e752329de5ee44"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / b627782c7244 / 4

- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-b663f6aa7a1c6c99cd1a7670e652cfc36d3df002f9b7a3b57049db2b9312d25e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b97158bb3d015657f76f205cc95a4c2b30f7ec68d7235bf35bc582778a45fb89"></a>

## os — os / 4ef7b0653eb8 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- os

<a id="canonical-022c88b9af45074655f9cbda2b60a4e804068f95b48ae7f759da2a1bb98cfd1e"></a>

Type: `"single"`. Computed.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

<a id="canonical-adef09e6102f8b3232f8e14938f91a25eef744888063b1687f65bb5d23c3b17a"></a>

## Direct properties — os / 4ef7b0653eb8 / 3

- [default_os_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a3497cf851d2c66dac13943989cbe796541923cd5864d0224fd123a3b7388571): complete subsection reference.

<a id="canonical-73dc8ae8ceb28857e0a535c823d66b28ae7fba5a464e78ca932c37c0e3a05b1a"></a>

<a id="canonical-fa7d9ba045954aed810ffd7ac588875b038b98cb251dc7b22816cd4cc4c64c95"></a>

## operating_system_version property — os / 4ef7b0653eb8 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-885cedc9fe216d6ec2c05ce9a4938ca47eb739f365725ed02f6a72a0460e500e"></a>

## Next pages — os / 4ef7b0653eb8 / 5

- [os.default_os_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a3497cf851d2c66dac13943989cbe796541923cd5864d0224fd123a3b7388571)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a3497cf851d2c66dac13943989cbe796541923cd5864d0224fd123a3b7388571"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab1483fa0219fd0dcf77e998b1ac3accea8dee4ab64944b3ee8bf915789d9612"></a>

## os.default_os_version — os.default_os_version / d689bd1636c6 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b663f6aa7a1c6c99cd1a7670e652cfc36d3df002f9b7a3b57049db2b9312d25e)
- os.default_os_version

<a id="canonical-4448d6b432dec4648caeafd7fb34b2b265aba59658bf1a840984ad29efe4e601"></a>

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

<a id="canonical-4554d90c59fcd96e0c58e6888d243f21959fc6c3d4033b9e312f1dfe8e66b188"></a>

## Direct properties — os.default_os_version / d689bd1636c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b49d61c3dc15a80f87fab06af14997ff78a77216c88a340be9d90cae66e6a170"></a>

## Next pages — os.default_os_version / d689bd1636c6 / 4

- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b663f6aa7a1c6c99cd1a7670e652cfc36d3df002f9b7a3b57049db2b9312d25e)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a78509fd22231425759d40c893af69315775e0ffd8a5bc1c8ec9250f2acb496b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a6f3987dfd919a7d6a2502c888fe735b8c2aa1d94a0fc8e6836bb492417bf6"></a>

## private_connect_disabled — private_connect_disabled / ad97aa954859 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- private_connect_disabled

<a id="canonical-ed8be2ab70fa1176e184fddc1a39b746916aae75e68558a2ec04589ad34ee73f"></a>

Type: `["object", {}]`. Computed.

\[OneOf: private\_connect\_disabled, private\_connectivity\] Enable this option

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

- [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ed8be2ab70fa1176e184fddc1a39b746916aae75e68558a2ec04589ad34ee73f)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d14b31de9b35e994d69a640b2e408bbd04ccafcf3ce9143cf4858355c9ae3f9f)

Select alternatives according to the provider validators above.

<a id="canonical-5b91d0895f7248c078e9a0524329175e925c6a8adfd894fbc898247c629530dd"></a>

## Direct properties — private_connect_disabled / ad97aa954859 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-295c84579e8152ccaf3560d47ccd7e012660d9c70592c46543da7db65f2725c6"></a>

## Next pages — private_connect_disabled / ad97aa954859 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ade9a78512a79b334cdcf9b77855c7c759febe180e7318793e1e0f5a051e9250"></a>

## private_connectivity — private_connectivity / 17e57c4f804b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- private_connectivity

<a id="canonical-d14b31de9b35e994d69a640b2e408bbd04ccafcf3ce9143cf4858355c9ae3f9f"></a>

Type: `"single"`. Computed.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

<a id="canonical-50710ad37f7a2e51caa55098ecde2b96fe29b7760f3c4c68a8044399d0e4a200"></a>

## Direct properties — private_connectivity / 17e57c4f804b / 3

- [cloud_link](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ed8c19e5fe87497dfdd40b0393faf12a1a893822ab90dbe4691f75c618c8aeaf): complete subsection reference.

- [inside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-7e3f8ebb35f3095b7952cce03dcf08cdb064ab58957110d40029b6d91bd7eae1): complete subsection reference.

- [outside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-980378d48bd4b6639983c8cd0e280fb9427889cd06fc474147437ef7ad1e5c24): complete subsection reference.

<a id="canonical-d48cf6d564635f78ff901c4304422a9812954b7001892c7adf782c8b4ad39e3d"></a>

## Next pages — private_connectivity / 17e57c4f804b / 4

- [private_connectivity.cloud_link](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ed8c19e5fe87497dfdd40b0393faf12a1a893822ab90dbe4691f75c618c8aeaf)
- [private_connectivity.inside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-7e3f8ebb35f3095b7952cce03dcf08cdb064ab58957110d40029b6d91bd7eae1)
- [private_connectivity.outside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-980378d48bd4b6639983c8cd0e280fb9427889cd06fc474147437ef7ad1e5c24)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-ed8c19e5fe87497dfdd40b0393faf12a1a893822ab90dbe4691f75c618c8aeaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72c61d422f8c1e8bd07340aa7b26dcddac52cd63a8035b416f7451649c41cdb0"></a>

## private_connectivity.cloud_link — private_connectivity.cloud_link / 939631d8fc1a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- private_connectivity.cloud_link

<a id="canonical-6b1adc8c43e25f15d237b16aeafecf46e49104839df0d9f760ee13048c56be42"></a>

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

<a id="canonical-4f5a21412923cd9284868bffa40fc1f18e669d76e4691707b18298fcd090497b"></a>

## Direct properties — private_connectivity.cloud_link / 939631d8fc1a / 3

<a id="canonical-09ec659c7daa34d7722eb8467c2396b139751456e5754b8ec18f069b5d9b9851"></a>

<a id="canonical-9175c6dc4f51f1162fd0169772579039be627d4152b062fa5b180a2cdc8cb2a1"></a>

## name property — private_connectivity.cloud_link / 939631d8fc1a / 4

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

<a id="canonical-e8fbd459fbaad54b4c21c5de821b0a6525831c7abb60c3ba28e2db289fc05e61"></a>

<a id="canonical-c51fb1a27ffcb72b4810b143e2e7e76b2ff62fe8810ad70295538ec13c8e5fe4"></a>

## namespace property — private_connectivity.cloud_link / 939631d8fc1a / 5

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

<a id="canonical-1d802fe6c857c31fb3e3d4f6f2f62dc387ea260e362af02621cea6399b7f6fa1"></a>

<a id="canonical-125b132509f1bc904335fddd189677061788cdab1182f74c2512a5b9660464cf"></a>

## tenant property — private_connectivity.cloud_link / 939631d8fc1a / 6

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

<a id="canonical-464fea64d1eda1db76f36af0696752c352eaa838fa647e635beeae674b9fd2a7"></a>

## Next pages — private_connectivity.cloud_link / 939631d8fc1a / 7

- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-7e3f8ebb35f3095b7952cce03dcf08cdb064ab58957110d40029b6d91bd7eae1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51b0a310329002c3c1b4f7d30a9b711fd3debc98e8c254de336d09e68a20af26"></a>

## private_connectivity.inside — private_connectivity.inside / 72d2fdec41a3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- private_connectivity.inside

<a id="canonical-c1901975a1a05b8599634952b603530ba72302db7fe8f7c3306f47b66664e78c"></a>

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

<a id="canonical-286170785ef58982aff0d7362427afba18ac3b28013cd39f793630adac400378"></a>

## Direct properties — private_connectivity.inside / 72d2fdec41a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1fa2f3dc7fa702d4351be613d1c480c88a27fdeed50c20860a40e2297303d54e"></a>

## Next pages — private_connectivity.inside / 72d2fdec41a3 / 4

- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-980378d48bd4b6639983c8cd0e280fb9427889cd06fc474147437ef7ad1e5c24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6c21f3d54ffecb36dbcab7ae93f925102f59cf71ded44c5062882e6f7218483"></a>

## private_connectivity.outside — private_connectivity.outside / 5e6993fe59a4 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- private_connectivity.outside

<a id="canonical-e038781147b5823e39e40b0db7dd6817328b1ede39d224923b0a8f3b23e49b86"></a>

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

<a id="canonical-9876883f3b050634d39758d90634724f95343203807be1cf0194479d8dbd4978"></a>

## Direct properties — private_connectivity.outside / 5e6993fe59a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2cd6490982fd38e5060e7f04fa6ac99443181a98fe00ef19da0bc03b946eb64a"></a>

## Next pages — private_connectivity.outside / 5e6993fe59a4 / 4

- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-34f2626bf35417bf94a0af238c2fe36603d4617d24af28174e43297681a6aa08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21e0060395f92d791b92e48342c45867e6ea5e4edbd7e491ab3161b98af75df7"></a>

## sw — sw / 3348a910c3aa / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- sw

<a id="canonical-2d11c7b1ecae586c5d9a973c05d5a34a18633255a86052b4454efe063ceb9f65"></a>

Type: `"single"`. Computed.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

<a id="canonical-92bc71767fb023624eeca726285dea479f3c5672bcbf875e50a5fc7b5fdf5879"></a>

## Direct properties — sw / 3348a910c3aa / 3

- [default_sw_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-7a14d4b035d2ded40345937a6e6219370299442439ea7629125e4b162e26e6c8): complete subsection reference.

<a id="canonical-a3b6eb8b63e4cd0cd39084c8fa2a6f17ab89a426ad486386231507414a836a82"></a>

<a id="canonical-c355e4908399450809df562eb695d6cbb35640f5dee53d1026d72554561f73b4"></a>

## volterra_software_version property — sw / 3348a910c3aa / 4

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-76bd0b718619f1e3e488561bf0ab4e658459afbbc46a851f9bf82d2b95f8cdf8"></a>

## Next pages — sw / 3348a910c3aa / 5

- [sw.default_sw_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-7a14d4b035d2ded40345937a6e6219370299442439ea7629125e4b162e26e6c8)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-7a14d4b035d2ded40345937a6e6219370299442439ea7629125e4b162e26e6c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17707d47d578fed8b742d71979fc853e6b0843cc4ae70585b2e39dcd2c50fd46"></a>

## sw.default_sw_version — sw.default_sw_version / 0f05045ba4b2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-34f2626bf35417bf94a0af238c2fe36603d4617d24af28174e43297681a6aa08)
- sw.default_sw_version

<a id="canonical-1ebcfee274f7425599b968996267fae3813db4272088c7cbd192cf0f3f9b5d32"></a>

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

<a id="canonical-a98931f7534251e9990bd861d1d02492555d665e0e3cf1327eea84ddc8539b5a"></a>

## Direct properties — sw.default_sw_version / 0f05045ba4b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-082ab3b13ed57b90f3696437d50d941441845d8f12872bf8c0b052e2883ba361"></a>

## Next pages — sw.default_sw_version / 0f05045ba4b2 / 4

- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-34f2626bf35417bf94a0af238c2fe36603d4617d24af28174e43297681a6aa08)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a7dfede66c6f88c6d8649f57e4468a404b29ff1dcbfc30fbb50f7fbfa3cd54a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c61ee103aa46578fa173eccfcd063a440b49b191a19cd59e8adc02f6099985a"></a>

## voltstack_cluster — voltstack_cluster / af58acb1da61 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- voltstack_cluster

<a id="canonical-237e3ceae77303258b3ff82422e1aa07d05f60ce403476754ef00c324a4d28d1"></a>

Type: `"single"`. Computed.

App Stack cluster of single interface GCP site.

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
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

<a id="canonical-4903dca83f2a123e9820df3380a14ebe9780b44c878a50260211cb0aed0b3e6f"></a>

## Direct properties — voltstack_cluster / af58acb1da61 / 3

- [active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2e2a2fc3a23f521a071cb68f31a2caa3c7015218017561a49ddf8ab71c68a4ba): complete subsection reference.

- [active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c07fa8d96343829cf25c447a8a99570a3d2395e456df02bd1bb3c9edec790c08): complete subsection reference.

- [active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-434b52ac6cd13c93dd42bdf830d07e7d9a283e45722b52e43fdf14c6bbaa44bf): complete subsection reference.

- [dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9a6d9df3421ae74060960c2c46dc4e936eb28a12afd22aed02ab5fff080f47de): complete subsection reference.

- [default_storage](data-sources--gcp_vpc_site--reference--group-004.md#canonical-d8aa67690095e45ee0fad45bf634ba3ba54c6b93a27c47e33e0a29c555b28e1d): complete subsection reference.

- [forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-004.md#canonical-25f646bc37236b792198e5a6410d9441072f41ce6dd08bb6aa962f5eb9e6f81c): complete subsection reference.

<a id="canonical-fcb47aaef9dec15c0ee7ef6d678d2a101157b556ca6f197bc2d83a2d01559c87"></a>

<a id="canonical-20d5d093422134cca3723639fbca9ac9a64e84fe8ee10e97915748423be15f11"></a>

## gcp_certified_hw property — voltstack_cluster / af58acb1da61 / 4

Type: `"string"`. Computed.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-126757094d19ebc76dc55516f6711b7a76e35af16a3ab8f722ed64ecc88f6c49"></a>

<a id="canonical-d11437e56b1644cead7753d70144a0ccaf3d0705cd0566cf773903373309d26e"></a>

## gcp_zone_names property — voltstack_cluster / af58acb1da61 / 5

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-b6d154e387004883ccbebea0768e6bf15fbd386a2a8efc36e443e7d3b3e55db9): complete subsection reference.

- [k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c494cbfca3211f542e9b26a310d6bb123d5d75c8913eaaa35fe1727f43743e75): complete subsection reference.

- [no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3df67e3c4c257873f977842cd2e829b7bdd23f0d7796590eff0535af0be9433e): complete subsection reference.

- [no_forward_proxy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-4bc0ffded5b81346d40ae76755e0522e8a3a0f6487afd70512a1e7edd9e0cbd9): complete subsection reference.

- [no_global_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9996d6690a71b720d13c2dd7153631e84a5ac882b0ff93c7817d9cd0d5e30a7d): complete subsection reference.

- [no_k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-7025a45fea4ebef0ac0c0ef17ef80054872f14c99dc7db728104cc8c11ca7c95): complete subsection reference.

- [no_network_policy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f0b692c3f372c61126b14efc86403e6406091305976a4705cae210b86e660a0f): complete subsection reference.

- [no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-70fae7b81864b8e77424d15b1ced4ff426e25ca7ddc8e45a8869d1da337eac8a): complete subsection reference.

<a id="canonical-d70324685a22392dd31d1ba2ad760eb91ee8da1e465ed73b3ebd834168d4d687"></a>

<a id="canonical-5450a145dce51aa04df399614ace1a954e68cceb36f45b6ee5578e9994f8d99d"></a>

## node_number property — voltstack_cluster / af58acb1da61 / 6

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f2d9003209dd614de6d91c14efe189b600687dfd3659c8f4a05789934b75eefc): complete subsection reference.

- [site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-d552fcd8c86e0927d7169d098e4b066d01031fa46baa668f93ed9ce49a703926): complete subsection reference.

- [site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0e6bd161a6ef6dc57a0ed216b21bef8fb4661dfec32465fbcaef0329a8fab4bc): complete subsection reference.

- [sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-625de50c3d04f06b78954d05cfb89428d4fe2c0edb1ae146e77385458ed0bf4f): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0045f289c8061302635b31b241c1925ae3549fae6f5d30298920acc33ec764af): complete subsection reference.

- [storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-4647ce26177cf0ab24c4ad65ee013f356f3968eebc85f179c90cbc11320a1122): complete subsection reference.
