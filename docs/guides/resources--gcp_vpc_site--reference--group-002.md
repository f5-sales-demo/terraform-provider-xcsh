---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-061eaa2d8b178a2c0c5cd13420c61140d88d163bba4169e79688be285c3798bd"></a>

## gcp_zone_names property — ingress_egress_gw / 4e4946f67bcb / 5

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
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

- [global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6): complete subsection reference.

- [inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9): complete subsection reference.

- [inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb): complete subsection reference.

- [inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8): complete subsection reference.

- [no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-2b666e8da5e36ab6b073b61b99b4634dffe53fbf82fa8bd25a37681ca0e4fded): complete subsection reference.

- [no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-59f800c76a03ffeb99934404b60cf40395a134561fb9b764358e72709fae3e36): complete subsection reference.

- [no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-d227621ac9c0988f24f88572c0db7090ca5032a95b2c1f8c34912aa5d344481f): complete subsection reference.

- [no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-214d17241f00f1246d31aa08f82071976305d5c0539291623d4fe93adc1b2b67): complete subsection reference.

- [no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-72745e6670a4b81ab090ff751718a4bdc631007f55118552475a38b827747c74): complete subsection reference.

- [no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-5a8f713e1bb71eb78b47b12e22c03b335b6c75f1aca76f789725413debc883a4): complete subsection reference.

<a id="canonical-9a4d3e975620c56fbfdf085d117ee6ac3ac0f551a9eae2e1c9b6e1d7eba8d9a6"></a>

<a id="canonical-f0c126a44a14734188151c4f445e64c24a4d4ceac717c14754ed4834ad9a6dd6"></a>

## node_number property — ingress_egress_gw / 4e4946f67bcb / 6

Type: `"number"`. Optional.

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

- [outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe): complete subsection reference.

- [outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78): complete subsection reference.

- [outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1): complete subsection reference.

- [performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d): complete subsection reference.

- [sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-dd54f6695362eb9dc9e39e78abd5a0cf5108081701cee3694a848ce757bb44ef): complete subsection reference.

- [sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-a23466e2ca253c8f2e844bb917b8d4e2850077c11c5e21634704806eed281cf5): complete subsection reference.

<a id="canonical-02f5a78e8e287ab66901f6e89e3c1dcb6e3ec6f9d41d27f959ff351250e6a4af"></a>

## Next pages — ingress_egress_gw / 4e4946f67bcb / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-7105b339fc23b53004862f942e57d3faabc1403632864c87f7a2d774b0def427)
- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-b5b4c33b332c2f3597d8f4e3bccfd55bad4b343082aa5b6cda3e4b4c53735cfb)
- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-261ff2a541c862c8da4a249ad3655fa60954efc402186110687a8752816c111e)
- [ingress_egress_gw.dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-250c9d377018d89fb39f874d1fdf70099ca624d2e166fa0b513f2cd7d81f3bb0)
- [ingress_egress_gw.dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-9379418bfa11e879e5b850a97ce6b3ba8abd41fdd9ad9ad2531c1ba333f0d44a)
- [ingress_egress_gw.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-85c804a1ee2f8c0833998b30c7199fdd3a3abd09c91542ce4a85b65b7d6c8a3f)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8)
- [ingress_egress_gw.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-2b666e8da5e36ab6b073b61b99b4634dffe53fbf82fa8bd25a37681ca0e4fded)
- [ingress_egress_gw.no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-59f800c76a03ffeb99934404b60cf40395a134561fb9b764358e72709fae3e36)
- [ingress_egress_gw.no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-d227621ac9c0988f24f88572c0db7090ca5032a95b2c1f8c34912aa5d344481f)
- [ingress_egress_gw.no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-214d17241f00f1246d31aa08f82071976305d5c0539291623d4fe93adc1b2b67)
- [ingress_egress_gw.no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-72745e6670a4b81ab090ff751718a4bdc631007f55118552475a38b827747c74)
- [ingress_egress_gw.no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-5a8f713e1bb71eb78b47b12e22c03b335b6c75f1aca76f789725413debc883a4)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [ingress_egress_gw.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-dd54f6695362eb9dc9e39e78abd5a0cf5108081701cee3694a848ce757bb44ef)
- [ingress_egress_gw.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-a23466e2ca253c8f2e844bb917b8d4e2850077c11c5e21634704806eed281cf5)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-7105b339fc23b53004862f942e57d3faabc1403632864c87f7a2d774b0def427"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff8b05d0a6eb4d0aa173dc0d09a00d1521c085580a89c3c78abb90f855291ca"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies / 8e2eaf42b9b2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-bfb1a0f2c82cda5cda05bf459016cdb72bf99d65da842121f625ba46aadcb916"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-35f9e30288ae126a76b401db53ab8ca99bcf89cd592ec8a86b52eac4be8169ac"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies / 8e2eaf42b9b2 / 3

- [enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-6300e73f0354c64a1358e6752b631e219c83fb5fe13fa16ef34d103267ff3d1e): complete subsection reference.

<a id="canonical-06969820889988e16040128ab4664842a61259918f86b2ba431703f0c3e9be4e"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies / 8e2eaf42b9b2 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-6300e73f0354c64a1358e6752b631e219c83fb5fe13fa16ef34d103267ff3d1e)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-6300e73f0354c64a1358e6752b631e219c83fb5fe13fa16ef34d103267ff3d1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7cec87c87fea7e534867ce422e04cf250a97c85b37192c6b683bf1073892089"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-7105b339fc23b53004862f942e57d3faabc1403632864c87f7a2d774b0def427)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-17555357c5c1378d1e8c88a4f2d69b8ccf60233c48e2be8a5dd42299f4aed466"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-26d4566c525b82d15b5c7798b95fa345575f16b3ee2e606109415514ab8bb4be"></a>

## Direct properties — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 3

<a id="canonical-2138d0f069cb0c87aec03b446f29aa3c29207ced44e73c50552bc58c2910a105"></a>

<a id="canonical-a1ecdb4fa41ed6752000a520a16c1f1a95700a1e896029ec1b0d8fb79e8e29bb"></a>

## name property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 4

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

<a id="canonical-466f0d98a341c4e6a362913dd457834276b642e683b62963e599d9a346d708f7"></a>

<a id="canonical-09e5a9af7a425051a28dc7b14c655b24604d7396dc01b7b9bebd696a694ea1bc"></a>

## namespace property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 5

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

<a id="canonical-5321823f1f15e918b22d1e870db957c2ffe1ccd58cbc3da3a796fbe332a60643"></a>

<a id="canonical-d5a89018e9c6f21019f4c7198c20fa9366cdf70477b45c4c1ac338820128922d"></a>

## tenant property — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 6

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

<a id="canonical-5258b810f9a8ff6f96923df5208bdccf2143d3b56163e394b76e92f8df21093c"></a>

## Next pages — ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies / f6cf4fd33076 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-7105b339fc23b53004862f942e57d3faabc1403632864c87f7a2d774b0def427)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-b5b4c33b332c2f3597d8f4e3bccfd55bad4b343082aa5b6cda3e4b4c53735cfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c093d2df2d600661fca57dd1c38eecc0a8e542c274e56689e9a06dddc4a11ca"></a>

## ingress_egress_gw.active_forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies / a729b0d2a675 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-3e18e060fb0bd220f2044233bf27775082af2e8dae587c3492352d512ac2a96c"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-407ed58e4e9edd8efafc3531767b5ebed99e0056dfd1a256748ce40051bace9a"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies / a729b0d2a675 / 3

- [forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-dfa8f661b3c89b139e72f55c26be002da2064718f67e3a6d8eea00e4a3e13a1f): complete subsection reference.

<a id="canonical-acb88edd9a0f233a7f8b0b87366d150fba40f97baa7cea2a456d3e3d27fa8b7e"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies / a729b0d2a675 / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-dfa8f661b3c89b139e72f55c26be002da2064718f67e3a6d8eea00e4a3e13a1f)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-dfa8f661b3c89b139e72f55c26be002da2064718f67e3a6d8eea00e4a3e13a1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3019f14cacb296692694f9f60a5fa9ab4bea40bc67b666adb34465d5acf88de"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-b5b4c33b332c2f3597d8f4e3bccfd55bad4b343082aa5b6cda3e4b4c53735cfb)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-e01e25d7ac6bc1e84c85ce135320e9e06317e85a31f93f3ea2e241e8b2f8336e"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3271c03033baabef4748b446bfba937a142a9cd31831d1fdedaa94a1c16cc34"></a>

## Direct properties — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 3

<a id="canonical-3d5cdef532b3b658a72d04af05f7f5be6ed78cdceda3d96eb9427db96a49abb2"></a>

<a id="canonical-187438b3cfffe59fce496d1f12e530b0b206383b44f9c1886d7c710695017fba"></a>

## name property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 4

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

<a id="canonical-ee4e6c0e29432561b192b6100a18a01262d0c35f7374636731890505ec6c0430"></a>

<a id="canonical-42bc7b095cb73495d2c953911192e3d3933d41bf8300eb850bcd9f6b204870d0"></a>

## namespace property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 5

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

<a id="canonical-4948b76283667f872336c949f8a73f1848be9e03db0910afa220b7218e6437bf"></a>

<a id="canonical-40ed39c1fea472664b8309438ffbefc1ef4f207fe01296debb5fa735dc62bec1"></a>

## tenant property — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 6

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

<a id="canonical-6420e6c75b65e86aa4d1c3e4c9c7049b15bf7cb2c2a60b54b1d6cd52f5091688"></a>

## Next pages — ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies / 999e624a330b / 7

- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-b5b4c33b332c2f3597d8f4e3bccfd55bad4b343082aa5b6cda3e4b4c53735cfb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-261ff2a541c862c8da4a249ad3655fa60954efc402186110687a8752816c111e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b328c0e5446f3e6f8cea39ee7ac4ab75544fc45a6deff6667f0abd97c693692e"></a>

## ingress_egress_gw.active_network_policies — ingress_egress_gw.active_network_policies / d6f4413846e8 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.active_network_policies

<a id="canonical-a880a857c5d584ab1c88d4d44fe6d6dc875d71745de890c8feab5add45023087"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ea3ea248f56b570c84f95f4010434edbb1c8141cfbd17f341b40cb1034763e2"></a>

## Direct properties — ingress_egress_gw.active_network_policies / d6f4413846e8 / 3

- [network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0d56221751e43d539fe56e380f093a30435615e2d6e4c5fda2df67c5df218b02): complete subsection reference.

<a id="canonical-57bd4d06d2816b9326e6cbc627c618f3a021509e1b7f62fd4caa7f15d04d9afb"></a>

## Next pages — ingress_egress_gw.active_network_policies / d6f4413846e8 / 4

- [ingress_egress_gw.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0d56221751e43d539fe56e380f093a30435615e2d6e4c5fda2df67c5df218b02)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-0d56221751e43d539fe56e380f093a30435615e2d6e4c5fda2df67c5df218b02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f60b986f72fe91e6de2927dde044260c400dfce56ebd3d9e1436efe8bc381c5"></a>

## ingress_egress_gw.active_network_policies.network_policies — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-261ff2a541c862c8da4a249ad3655fa60954efc402186110687a8752816c111e)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-e58cebf376467eded302bf842364ed535f299c8f138ff6e30fcdb1b2698d95d5"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-fa3fa280c357e28f7501d49de08e5e37c2db7b10d2c0391bb42b84b29ce91f8a"></a>

## Direct properties — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 3

<a id="canonical-1529af2b1c1953451c85bb5ef26235293fb12ba3b36d96da203eef02ae038ad9"></a>

<a id="canonical-022cd52d78ac2fcffc5b0d5bb36d4ee91063d80d9d6795a9aca0a2ba38ff3dff"></a>

## name property — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 4

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

<a id="canonical-1a559c7cba3eae0d49bd1a966d7ff53ec651e1503f98b9b6e90ee7e1491c55d7"></a>

<a id="canonical-155acfa447d7978981f99f770040b64c1010de0cdd5bf68f8408b500af200caa"></a>

## namespace property — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 5

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

<a id="canonical-22c290070148b0a53f52309bc3f7c8097f38a991669279bcaf5494559a92fcd3"></a>

<a id="canonical-925a968576b8c8e08533e6ec2693884223692a626415b647733779c36a10e794"></a>

## tenant property — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 6

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

<a id="canonical-9aaae7126b2af6d7c09144e237ad7d30440fd9fd6493f557f1b376496af8f05b"></a>

## Next pages — ingress_egress_gw.active_network_policies.network_policies / 5bc5089c05a7 / 7

- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-261ff2a541c862c8da4a249ad3655fa60954efc402186110687a8752816c111e)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-250c9d377018d89fb39f874d1fdf70099ca624d2e166fa0b513f2cd7d81f3bb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd28f37fef99b2a06e31a897eeffabbe13479186424fb32af24be08543261d79"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-501f103b08f1f067b8ec1f5805a743df1528958fb42a771e09de7dde8b5c9ff3"></a>

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7d0641e6540e2831296fa48d7dfe25a96182b49d2260aa84cc1b6eaa65c3052"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 3

<a id="canonical-0046e8910fb8a8ecbf7dd3534ffead266b6610709d36ab5c76d74e3875a4f311"></a>

<a id="canonical-3632d7baf036033a49ec59c4ea54f26e9b090607318e7a09a743722cb7bf10ea"></a>

## name property — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 4

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

<a id="canonical-07557df8f45696b2a22e98b7b2fb12bc07f0cc940737353c6aaf721a0f9f21ca"></a>

<a id="canonical-1e17591dc9ba30ffce729fff57b3100d332c59a89a894f45f31760abf2804a9c"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 5

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

<a id="canonical-684b0be60f4ba4f4d48263d96c042b4d56bee5c8b0c8489ed227a128eb66d25d"></a>

<a id="canonical-3cf9f38ba06463a01195661d79eed88ef5b1fc55a6b1a55cdae3c6857f773594"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 6

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

<a id="canonical-e3c9866131facf295573fbb7911ad2d20bda531b971f2dceee243e96d73d1d65"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_inside_vn / 63b56b7b8273 / 7

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-9379418bfa11e879e5b850a97ce6b3ba8abd41fdd9ad9ad2531c1ba333f0d44a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aa8e2f5fe35a46528cfeb4830e8b2912f116b4362c9f8d6d2d2ccf042f50bf8"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-12cb66a01c69faea311cd5b34a6a4e1d7598b4a791f914c0a99f414a5f4e09ff"></a>

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-77415dd91586ed00eb9883b4f32928bf31b441c62003e2d3ccb05662d88d1516"></a>

## Direct properties — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 3

<a id="canonical-1d1d9e0d49f03b6179d04ee6b098e10c629c139f3983b0659e29a5f297ff918f"></a>

<a id="canonical-6d229ff6a56d986cb1c26632ca68753659823967f111a73bd12736985831976d"></a>

## name property — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 4

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

<a id="canonical-4409cf0daeff9527c2f810447d8f30f5f4379d960411db8c085105f96e387462"></a>

<a id="canonical-e617190ab2ff3fe4530a51ad2121521b08c39478b8d753ff83726e6523efb423"></a>

## namespace property — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 5

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

<a id="canonical-9c1a17a37b5159c4e6af233f3ea2d8be09f66bfc1a5498ff1d53db78b1aaa3ec"></a>

<a id="canonical-f606a44f8930fb2fe680f6bb57e142d4517d8bde57a38d472b9387d840f2bde1"></a>

## tenant property — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 6

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

<a id="canonical-742bf91d72128cc5c9c3380989ea6e14911c71ce43dfb6d34b2637b5f4ac86b8"></a>

## Next pages — ingress_egress_gw.dc_cluster_group_outside_vn / cdf8c3a7848e / 7

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-85c804a1ee2f8c0833998b30c7199fdd3a3abd09c91542ce4a85b65b7d6c8a3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58b86f89a9f061cd359bdd6d51bd8dcee9010155dbfa16808a191f8477eceb02"></a>

## ingress_egress_gw.forward_proxy_allow_all — ingress_egress_gw.forward_proxy_allow_all / 423d9022e300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-69db5b5870f8eed91524c2f8c7dd28d548b50e4b244e7571215cdb013945db75"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-1a97d971887e7ffe87607cc363d4b17509229dd6a984d5d8f6266e46efa29c12"></a>

## Direct properties — ingress_egress_gw.forward_proxy_allow_all / 423d9022e300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0570208c6e88e38a9feabb1a2f2627c4493b6288957d34c3c00766c22866810"></a>

## Next pages — ingress_egress_gw.forward_proxy_allow_all / 423d9022e300 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b60e92fb607761dcc34211192b2809f0a57a908f91753661a9204564d49c092"></a>

## ingress_egress_gw.global_network_list — ingress_egress_gw.global_network_list / 7b995d633a5a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.global_network_list

<a id="canonical-f6232a6befb2fa4743cb124e92f64ffcc0eb1b13d53e893227fe7489860a3033"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-748ef6a8ade445e7152c103561b04eaa6ddb2dfdafd11585730e8d09f408d93b"></a>

## Direct properties — ingress_egress_gw.global_network_list / 7b995d633a5a / 3

- [global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099): complete subsection reference.

<a id="canonical-4baf73bbff795a2a54259d9222b495a6d61b4002f8e8f0e59674a922a2d695ee"></a>

## Next pages — ingress_egress_gw.global_network_list / 7b995d633a5a / 4

- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11fd58e8edded2db48c87ec763306bf9f8d095306a1c6ec2d33873ccc661bac8"></a>

## ingress_egress_gw.global_network_list.global_network_connections — ingress_egress_gw.global_network_list.global_network_connections / 13db3a7a1ab3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-a25718402d768acebe3c128d555b3ce87d2cccc30a13c385545b08c567615079"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-48aa364909148f5e90d6b3d0064c28da60af7d626fe752bca17d49e394d4883b"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections / 13db3a7a1ab3 / 3

- [sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-14e728034e356179188c1c5e44872efac5b3517b5dce7da05b8d421ff07a7f74): complete subsection reference.

- [slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-ad947942b5948fabd7d69980a71aa3a59625e8d178601e170ab01bdf072e27f3): complete subsection reference.

<a id="canonical-2af4b5d1a139c45784b43d0aa89d93b6b31791903242d44d0d5d5c1565099fea"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections / 13db3a7a1ab3 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-14e728034e356179188c1c5e44872efac5b3517b5dce7da05b8d421ff07a7f74)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-ad947942b5948fabd7d69980a71aa3a59625e8d178601e170ab01bdf072e27f3)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-14e728034e356179188c1c5e44872efac5b3517b5dce7da05b8d421ff07a7f74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffdc85e1d709efa4d87443034fed7ea1d01ce1a24ba816a5c3df2c6ac521ee89"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 7209b4b9ec1c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-6faccf5f50b0fbfd41bc71df4318824e9a713db744c6741fd8a04b9156b4fc82"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-158a5b22c088a78642962dc079276f104b26be13b5bc074dd2adf903d268b414"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 7209b4b9ec1c / 3

- [global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-5df0489f66fc661eb3a3a447d2a2f398046c8c84aa97cbf8ec54d7d71f71b438): complete subsection reference.

<a id="canonical-507282e12f6290070768fcc5c7a18837c52804dc6df7a5f246565ddf2a13f294"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / 7209b4b9ec1c / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-5df0489f66fc661eb3a3a447d2a2f398046c8c84aa97cbf8ec54d7d71f71b438)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5df0489f66fc661eb3a3a447d2a2f398046c8c84aa97cbf8ec54d7d71f71b438"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11bac237069196aa358d6107158d15ca6ac93514715c0670282757009c1cfecd"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-14e728034e356179188c1c5e44872efac5b3517b5dce7da05b8d421ff07a7f74)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-30a7611e7a9939dce982e704f025ad3d6996d4856a24847a645cf9b4e0d2bd11"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0e9098e9d8ae782e8dfef17658501d8a870ccdbf02bc73808ffa00bcda1a09cf"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 3

<a id="canonical-f562899604a617e600624a445bcb3afae218efadab8c5520b4e1cb4ec06919a1"></a>

<a id="canonical-5d1600ab6ecd38921efb491af49d93dc2eaceb580e23ae019c9ae212d07d20c1"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 4

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

<a id="canonical-e2c5f80ea18cc8ad7b1af5f17561930e5ee4dd8e8a87588f1a147fd108cc1cfd"></a>

<a id="canonical-5490ddfa65dea875877337aaf91e5e96597ed52e706a26a42514a0a44798f8f2"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 5

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

<a id="canonical-13eec7d79efa48b34a29f3764caa4d82cfaca68331892b654b5336b7b4e0b757"></a>

<a id="canonical-af3b4739ac378fc71b88695a51b361ae68a7c003a9853d4c6dcf992034786fd6"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 6

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

<a id="canonical-1d261fdd2e4aa1fbfda674f869bf30debc93adc913c8963abda834053eb47ba6"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_d / e85b46e821be / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-14e728034e356179188c1c5e44872efac5b3517b5dce7da05b8d421ff07a7f74)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ad947942b5948fabd7d69980a71aa3a59625e8d178601e170ab01bdf072e27f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1976a92d2800410cd8dd012b581aa3cbcb46c191e0e07381ba6536557ad21a81"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 3bf4a72289ae / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-9551ba1c94e9b38f2b07473534a8df451870eb4e7f6e2bad50b772304dde5733"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ba5b93bb5998f53558f1acb04c654034e45fcba7d9c7ac655a536810ee57890"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 3bf4a72289ae / 3

- [global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-1e79acb40b57ff2599856314c4b84497ae21da9e1861f72ea83eb062fdcb064d): complete subsection reference.

<a id="canonical-2094c97e4ab10fec5239798eb7db5c0d326cb4229b834288674cd50b33a788d3"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 3bf4a72289ae / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-1e79acb40b57ff2599856314c4b84497ae21da9e1861f72ea83eb062fdcb064d)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1e79acb40b57ff2599856314c4b84497ae21da9e1861f72ea83eb062fdcb064d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4fff44a58b56b9b189b3ceac1512eab626bf840deb4bdd1f28a9b96ec6882c8"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-5d7d09791ff839dcc12d85109683700550ff43340444f22e23f39679c6b979a6)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-549a2328cdee6c652a530f6aaa777b1f5f92606865b3377b0622d32919c5c099)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-ad947942b5948fabd7d69980a71aa3a59625e8d178601e170ab01bdf072e27f3)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-bbcfaddae81b03d5ddc64a3099f6654bfd1f344b96a4d804599dc35aa2d6cd3d"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-43fb9a9450626e44e9a01c93448c4a157f101d2c94478aee26cc6f64aca412d8"></a>

## Direct properties — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 3

<a id="canonical-3214937b6989592445b33b0b90afe5700ecdd3ae8e7059b9949f18bcb0d8bcf4"></a>

<a id="canonical-c2dfa2fd37e3878b92df36ecf6890ab924f07bd0f924a97cec6061c06f0ceb6c"></a>

## name property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 4

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

<a id="canonical-ac388b74f7e2f5d60b6cceea6a9d589a5bb7a20b9d142bfabb2b79a4f8f5bd15"></a>

<a id="canonical-c2e1827533b155394f5ecab8bca4abfdf4b0185fa7527397b678ae9d19ffe6cc"></a>

## namespace property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 5

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

<a id="canonical-d9f871f76d9fab394140da3ab5975d0cea0be27244e3d81f8c2a3fd6fe958d86"></a>

<a id="canonical-ffb2834233572b64fa999ff87c8eae8ae60d81633a6539928237833e61fbb7b4"></a>

## tenant property — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 6

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

<a id="canonical-76d4b664f3af12060df6281e874948a7a8cb815c165397a184d801f6153b4f12"></a>

## Next pages — ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_d / 23be308c45db / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-ad947942b5948fabd7d69980a71aa3a59625e8d178601e170ab01bdf072e27f3)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01ef41be3f3afecdf2a85c596a1a14e82c19810f037f6a4e5f2fb9565fc4af57"></a>

## ingress_egress_gw.inside_network — ingress_egress_gw.inside_network / ce9fae0fbc77 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.inside_network

<a id="canonical-688741c6be5b66e14589972f597ef9095ba7fc1d056fa3da4f68a1effec97cf5"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
inside_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1155d865ec56ebbc6b251e6706f2ffebdf1b14c418052107d21d7c309dce78a"></a>

## Direct properties — ingress_egress_gw.inside_network / ce9fae0fbc77 / 3

- [existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-f3c412c4467f4f12df7929800577aa2c72e2a4a217b1978ea8825db50a23ec50): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-a52cc9e1d20f4030ae1018b8f270db7aba3d35eea86fc8d7ed4349181f5ed6a3): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-d377415219ee85834beb9dc9cf4945b50e5966e817567b276eb0729ecd8cb470): complete subsection reference.

<a id="canonical-f29392ed3a9d03440d8911856faa4ece6d9319cb69bcb71044c66ecfe246a094"></a>

## Next pages — ingress_egress_gw.inside_network / ce9fae0fbc77 / 4

- [ingress_egress_gw.inside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-f3c412c4467f4f12df7929800577aa2c72e2a4a217b1978ea8825db50a23ec50)
- [ingress_egress_gw.inside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-a52cc9e1d20f4030ae1018b8f270db7aba3d35eea86fc8d7ed4349181f5ed6a3)
- [ingress_egress_gw.inside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-d377415219ee85834beb9dc9cf4945b50e5966e817567b276eb0729ecd8cb470)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f3c412c4467f4f12df7929800577aa2c72e2a4a217b1978ea8825db50a23ec50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-700391a06404777ec252826d38cdcac527179677adc71fa2b449b8cd51ee5353"></a>

## ingress_egress_gw.inside_network.existing_network — ingress_egress_gw.inside_network.existing_network / d362e56f5c89 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- ingress_egress_gw.inside_network.existing_network

<a id="canonical-23da4928402c09228ed3df4e2bbeba49ae51992a5745b903d61376abccfdf295"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2bf3a7abc9a3ed90cb25b92e6936d7f4e552d74312ea9ba8e926b25605351509"></a>

## Direct properties — ingress_egress_gw.inside_network.existing_network / d362e56f5c89 / 3

<a id="canonical-5226a3218c5b4c1de6651ad472a4300af22c4e9dc1088510f6f7bb0f99343fd8"></a>

<a id="canonical-478e0141d509610a37c05b7c13ba83929405a9fc95efa6bbf8800e49b5c728df"></a>

## name property — ingress_egress_gw.inside_network.existing_network / d362e56f5c89 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-e23db389e43807d914ad0410a553939344010385f3c8f7914394fe25c8280caa"></a>

## Next pages — ingress_egress_gw.inside_network.existing_network / d362e56f5c89 / 5

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a52cc9e1d20f4030ae1018b8f270db7aba3d35eea86fc8d7ed4349181f5ed6a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e7f668ed99b4bb7c2b03aac1abc40e25801a35b076d90201b869c3f560aa3c2"></a>

## ingress_egress_gw.inside_network.new_network — ingress_egress_gw.inside_network.new_network / a1602275f6ff / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- ingress_egress_gw.inside_network.new_network

<a id="canonical-86bfa1987f5c1fa62b8ec6f7546ab2f1936e7f8998a6d5be7aed077a098a6689"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-4776d006747f30b5729da5dfe13c6ac412796335354eba4d17396cacf9e72205"></a>

## Direct properties — ingress_egress_gw.inside_network.new_network / a1602275f6ff / 3

<a id="canonical-239b52592316f0c10e05080431c80e5a4e89c33257e2f3fde102c600a3d632a4"></a>

<a id="canonical-81bc35ccead0c76fccad053d80f1635f74ab52ce3ffa79c82bdc5ef15ea76de9"></a>

## name property — ingress_egress_gw.inside_network.new_network / a1602275f6ff / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-a337770a3e70f26534723a37368d386650f2ee112570ff8e64b9f3d5b986e49b"></a>

## Next pages — ingress_egress_gw.inside_network.new_network / a1602275f6ff / 5

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d377415219ee85834beb9dc9cf4945b50e5966e817567b276eb0729ecd8cb470"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1668598c22369bda3fd13929726f0b51e0d8e783200e3316223623e1ecd1c3ae"></a>

## ingress_egress_gw.inside_network.new_network_autogenerate — ingress_egress_gw.inside_network.new_network_autogenerate / e3b80d416cd3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- ingress_egress_gw.inside_network.new_network_autogenerate

<a id="canonical-dc161626c8b1a2eeea7c3c86d4b74ea19072cddb7bb53283a727e4909a8b36f9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
new_network_autogenerate = {}
```

<a id="canonical-eed8377c0083dd084ce62245aa8e925da3dd8b97baf98c1911aac29b1f703268"></a>

## Direct properties — ingress_egress_gw.inside_network.new_network_autogenerate / e3b80d416cd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21bdc03787c62f623ce6e3fabedac20ac25d03f1d0ab33b1c31c5e914bb62f6f"></a>

## Next pages — ingress_egress_gw.inside_network.new_network_autogenerate / e3b80d416cd3 / 4

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-9bc51a171bfb8a95228f309286df21fbc5c4a5db310d23f8bef26ed6175798a9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2b73e6b4062fe0032c5df668f4198fea2bc9ea90667f228ab2e107d0a26ff6b"></a>

## ingress_egress_gw.inside_static_routes — ingress_egress_gw.inside_static_routes / 505bc84eee36 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.inside_static_routes

<a id="canonical-1f5a42ecc9836166f7d05a8587619b67138976bbfe4c6567b7a0f67385202d0f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ce627d3a233d08b8131f0580c246a663d32f4b648b4b1fa17667907f2707897"></a>

## Direct properties — ingress_egress_gw.inside_static_routes / 505bc84eee36 / 3

- [static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55): complete subsection reference.

<a id="canonical-2634293576dba884b08cbe99ad4595638cdd51b7aa38b075be079ca1f32ff247"></a>

## Next pages — ingress_egress_gw.inside_static_routes / 505bc84eee36 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-415c383c421def5353a5740a14ed445e08c4f59d012beeff22ba14bd425a300f"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — ingress_egress_gw.inside_static_routes.static_route_list / b7435968f382 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-f1d24af0e71b3982c6a0eafc9b59fdd30603c18c919b3a50377571b9c8201c52"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d0ce70f1e1b1b76b28edaa84b1525b2ecf69bfe57f8940fecf4ff46b9bea49d"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list / b7435968f382 / 3

- [custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8): complete subsection reference.

<a id="canonical-86dbae8132f87431ef54afd137385dd386c5cc0f0769e9ff1ae5b98b9ce220b4"></a>

<a id="canonical-1b3d9951485e4c6b0474a2328d5520d6b77bb0f4327b2f4f0d824ec9f207145b"></a>

## simple_static_route property — ingress_egress_gw.inside_static_routes.static_route_list / b7435968f382 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-8b669deeb9097816e3fa4f617290fc18a73ab4352a344d3ea028f7c7f32b79cd"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list / b7435968f382 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ea86b0ac01d62b85f36f655f6d657518de1a7e3f8a3af17d4b4039bf02a7628"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / ef1c4ea9368b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-e059814a50ae3fee71312bc400f89142391155727633f789fae00a93c308e027"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-d803e6316225257cb22d23d44010bc0d616c25eab37af9d65cdc910d46e86eec"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / ef1c4ea9368b / 3

<a id="canonical-01a47ae11c2f491aa43d586952cadd7b5b69212fc57b8d2079a6d206882b25bf"></a>

<a id="canonical-c073aa7db04d3820d6ff3e16ac2615a73e3a94606716749bad7c11f10dbd00dc"></a>

## attrs property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / ef1c4ea9368b / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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

- [labels](resources--gcp_vpc_site--reference--group-002.md#canonical-65ff6b2ca7915d9798b8232079709c1a8fb70b7eb90830b411c5b70a620cd0d6): complete subsection reference.

- [nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc): complete subsection reference.

- [subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16): complete subsection reference.

<a id="canonical-cba17930e0b05959506faae0537e818f99797544df3d72a3ff397cc69d370b86"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / ef1c4ea9368b / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-002.md#canonical-65ff6b2ca7915d9798b8232079709c1a8fb70b7eb90830b411c5b70a620cd0d6)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-65ff6b2ca7915d9798b8232079709c1a8fb70b7eb90830b411c5b70a620cd0d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37018c877dfe770f2047920534b83279f9d1ea1a53d34011ae1b19390a423ea5"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 631cc24d0295 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-e8974fee05391eb3de2a4421d40c8c4756341a254b72b5729523a8b10c2ada9f"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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

<a id="canonical-33806992cf630d5635411a544a8a2c2f7204b01dc0e68b22164a071e0a5f8725"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 631cc24d0295 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af6a6f04e2ac07c119245757a2fc738017491e9ff3ca0fd339074affc385057d"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 631cc24d0295 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1618590b1a6f78948dae0f8e5e294efd18c2747a23c75e3fa8e24444afa747e8"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e3dc0d2978b3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-fe771a6ac685c7225f4e7de136e7381b1da2c54118d8147ab7485f0bd5df7d8d"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-a61dd44a519375998f87c26860dd74a032ac54e334d8497d11c2e7b0d1f586a4"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e3dc0d2978b3 / 3

- [interface](resources--gcp_vpc_site--reference--group-002.md#canonical-54965f8a79b833532209229514f2795af6e229548ed128dbd89d1b185d13db49): complete subsection reference.

- [nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d): complete subsection reference.

<a id="canonical-cac937c00045c8f407cd86fd838dcf3337477fe364052ccc3baa212d477a4788"></a>

<a id="canonical-c82a4c6c629e450cfe7f397295776a30cbb1c2987def2d1bf7198746ed1043eb"></a>

## type property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e3dc0d2978b3 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c3b33c836a0b4372db8d1bd5d5b386e394edfc0c1b60de987bb4ce5959decbcb"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e3dc0d2978b3 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-002.md#canonical-54965f8a79b833532209229514f2795af6e229548ed128dbd89d1b185d13db49)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-54965f8a79b833532209229514f2795af6e229548ed128dbd89d1b185d13db49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16c49fc2bb2b44be8799d0c073186cce1a34535898847d6f18c6b72eb4de47df"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-154670badfd020ea043c1418db56d8e4af30331d3df5d1dad3c5a6d88b6afb19"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b67b9237e9c83fbc777d2d32b280cf7c18b34119edf071762dfe568360afe00"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 3

<a id="canonical-c7933a8dd8bbdf32e96f77a83ec0f7c08afed97868c62975add2e59cbbf99ad3"></a>

<a id="canonical-e453a18e57cdc668c5c8392ec39fc51799e7cb58b15932d2502bfc60fb662b7f"></a>

## kind property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 4

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

<a id="canonical-3dd81986b40c36c9bf5cb4e39e19e4618a479689a0903dcf286aaaa76451b73a"></a>

<a id="canonical-fc1d87c42a94c3ee3120ec7e14698a1456c46685ce3bd96926df8b44d9848bfd"></a>

## name property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 5

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

<a id="canonical-6afc29e9c08a3f62baf2a17be7bc1a7f2d011638ea825fc96694850e5ea4c656"></a>

<a id="canonical-3163a8a16adcd52e7c7b15bfe7f59ddb719dc3d7d6f7f164a93bece3eca6903e"></a>

## namespace property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 6

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

<a id="canonical-a930da51c4b37d2133c21fa7eb182bfd1f3661218f5614132d030ea1f4508455"></a>

<a id="canonical-18a5609637e657d658a1726bb1e9dbdd50e671499afb7412b9e3b8e5d3ee62c3"></a>

## tenant property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 7

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

<a id="canonical-6c826cbfd0eedc39ce1207c4ca1de921f576f70b6e4f7c6a46aa49172d09c6d8"></a>

<a id="canonical-cc5b252aa69010e5742cc24a206b4772aacf32367c650b7755c61789042499b9"></a>

## uid property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 8

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

<a id="canonical-396b13dcb502d77ee07a2f6fa9bbb79ba2722d040a3c6de4acc3685682819ec7"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 1a1ab4ce7be9 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08bde69795fb2127d2ab9f430cc08f6c93c16de99b4112912e83a2180d41adda"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3b530296187e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-976f72460a7570d06775a6e472f2dced5e20792e802aa73a2268ca4be3f981e9"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-4e6de751b4ff7c02458c3072dcc53d179cc745cf2fa59f504e9395d2c03922fb"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3b530296187e / 3

- [dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198): complete subsection reference.

- [ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-c82250a39c300c24419e06cd5b3dfc0dc43019e57026298109ded278b2e64e8a): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-b0e4a58488dac446cd59bedfbd566fe902c3da85016e9aee362249537f640f36): complete subsection reference.

<a id="canonical-c444b560718df984433a6b1f02534a4673cdf8e5dd7e63efcfa9d43ddcf96102"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3b530296187e / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-c82250a39c300c24419e06cd5b3dfc0dc43019e57026298109ded278b2e64e8a)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-b0e4a58488dac446cd59bedfbd566fe902c3da85016e9aee362249537f640f36)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a283b89c5daf6eb84b2d77755cca78b284571b8172a379200d3cf796e770fdf0"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / cb3ff1d01761 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-c81ec834fce277b239413aa90089352e682f42f489bc2e74fe0f4471f399175e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5f7fce79ac241c3087d5048518cad617a964666865e77f7c9dfe1d92fcc37cd"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / cb3ff1d01761 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-5fba001dcaf672d49bc2bf65a6cccf5b9aa62f90d2d00f7864a75eb230d8720e): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-68dfc11d1af3945157e9778e585085989c39c3062d76819c5c986dae36282b55): complete subsection reference.

<a id="canonical-4270cf3717bdcc15704354955950e2935577a34c29d9e79ef7637f3e397257fe"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / cb3ff1d01761 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-5fba001dcaf672d49bc2bf65a6cccf5b9aa62f90d2d00f7864a75eb230d8720e)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-68dfc11d1af3945157e9778e585085989c39c3062d76819c5c986dae36282b55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5fba001dcaf672d49bc2bf65a6cccf5b9aa62f90d2d00f7864a75eb230d8720e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-265bfc64e5620377bd09ac43d04ad9ac9b59dd1a62b7d870f71cd8245223bbe1"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 77a00512e17b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-6b07e0650516894cc0385b8191e8d9617cbcd7f62185e355007f199f6e5c3c6e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-837ba7be4c1c2e70d0624710f0acc8a6ba79185bdb4aa5b068310be1db0916ec"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 77a00512e17b / 3

<a id="canonical-15bfbcfadc77b064ef8cd1e8893e284d0bddcc112325e4175269fadcce4a9b46"></a>

<a id="canonical-3216a7418d91f0977d3c86df844660e9ceeef9a36da7f1e31bdf60221104f49a"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 77a00512e17b / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-98f56141667f2c01b0207039e62aae8f1ce6e18161b113eaf67b61bea54c1759"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 77a00512e17b / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-68dfc11d1af3945157e9778e585085989c39c3062d76819c5c986dae36282b55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba296265cce20b4a6b4618ebe82f11308f618ceffc0e95ec06f1af81c2ca75d4"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 4fab2b3e5401 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-c16376571cd79ec1ed524108c28d07af69b5f63e6f578da4527be0f1bbea6f37"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ea0bc2840a6b235e68ad3dae84b86a0a2c49eed30d67c8c4799675d3cedde36"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 4fab2b3e5401 / 3

<a id="canonical-7489972db52e02ef63ce31ebb066f18a468bf198589578debe6dafffb68e3479"></a>

<a id="canonical-ed4fb585fbdc718bbdb5fc0a8372e257d51f64547d2146611d5b48e10f177775"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 4fab2b3e5401 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-658fc3669b9f0740a419dac0680bfb1b17daf0eba3579affd51abcd4369d9b95"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 4fab2b3e5401 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-68ff0be570a2b8b38ac33350fd58ac4854707b50ae78865a2945d317a740d198)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c82250a39c300c24419e06cd5b3dfc0dc43019e57026298109ded278b2e64e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27cff90fb33533259c0b6ae17a784d7ef4b276764768082bcb594fb8fdcac048"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 906fd01f8bf5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-488b42146d17984186295ce78d62460fd298a0da407db930394c6a207b152287"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b1f9b28c9b5d4b250480da1ebfee7dcfce6203c2c34a2d041c067c8af193e6a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 906fd01f8bf5 / 3

<a id="canonical-1633e3856623517327f69a21c753363aecd13454137d217194ff313742435d39"></a>

<a id="canonical-7a48d10c81ccfb99fac2f656bce3fa9d926851a1b3bf4a5d5404fbbdd4239c16"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 906fd01f8bf5 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-cf7f168268ee430bf93dbf61aee4685faa61c1226ad1de911c59ac6fd3b62f52"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 906fd01f8bf5 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-b0e4a58488dac446cd59bedfbd566fe902c3da85016e9aee362249537f640f36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f91a83bfcd42bbda4663eab432210aa1f02965e48b8ef61106c8e3a908515b5"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3c5559423606 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-f03e9ea3e9308d39302ff931e721de959be31e5a37cc5833d8d695acd39c56cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-4e69b3ee4801194b6a31086e6fd82f1f8db01b2e446081e8fb36195c88e9c6c9"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-25ab88b8e946b881ab05c8ae20a3d65f7eab803b1661418bedd3669b5f3694c0"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3c5559423606 / 3

<a id="canonical-26d7b0d8aa65542646e79e959d6f9e497a94680d9b3dca65c055580869bdb731"></a>

<a id="canonical-5a0d8e3e02c4601a164ac2e44bffef1eea6301c80c18f7a707a3484d70d7e277"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3c5559423606 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-0f7d3f8185b77f7eb4e24583fb229ffe5d6b95e0ca3522829caf7f3ad50c3bc0"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 3c5559423606 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-28d37a160bf33b9d974ca0cb28a7a29f52f9ee8e8b184d42150fa29c2e3ef04d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d716b0e2796d5aaa86d8e0c82ca0f053d270280f6fe6b02a26d49f1a7cbed8b"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 05fae3feb762 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-f5a3834b9c9286f60ea28cfea73da85b9bb4c1e4ac8829b2dce0e9f1fdb51038"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-ce1676b68de0fc48efe494c6145582bf9d8fbeb551510c271a4ede782118176f"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 05fae3feb762 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-5a46800804a55c69b79d21d7408c75eaa35b4f048d86b70c6a293c19610ee054): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-399e39945b800331abcc5f02bf974afd1b16776c443b7c5bb5125c72e8a1d0cf): complete subsection reference.

<a id="canonical-e0ac5673a59ea2a45e048e2f6db372af5c2fd95a8840a1338d52d680b3819b15"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 05fae3feb762 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-5a46800804a55c69b79d21d7408c75eaa35b4f048d86b70c6a293c19610ee054)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-399e39945b800331abcc5f02bf974afd1b16776c443b7c5bb5125c72e8a1d0cf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5a46800804a55c69b79d21d7408c75eaa35b4f048d86b70c6a293c19610ee054"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9011baaa80f36b1e22201fae3987913c9cd2b5f954362d4512ac27e778a3f390"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 03554ef91c2f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-2551ac5faccad616bc98aa48db994d68f935788c978bf767f4921033eb9b868e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c383eaae4e82b1c4221b91889ae61ae74a822841910ea1e70960f0a92c977d9"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 03554ef91c2f / 3

<a id="canonical-ef1c0aaa625a8e96944187b927c6a1ef99d38fa1675cbf4c88aa485eaae7e1a7"></a>

<a id="canonical-fbfcf2b19d9c3d7c4c6948e901470ad4ad00b812ae7b0fa048593b3273b66299"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 03554ef91c2f / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
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

<a id="canonical-2c14dc8eb59474be45b08f4485491b0bb2a7ca4d58c2241f647264e0d52bcf87"></a>

<a id="canonical-62abec18dc617674621ff2a41401439ef2ff6695c9d147c7d5e6cc12da08d034"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 03554ef91c2f / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-5d353ee8eebeb30a99216455edcea8d9e2918cf3666a45b49ef8e7fc0a2f5ca8"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 03554ef91c2f / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-399e39945b800331abcc5f02bf974afd1b16776c443b7c5bb5125c72e8a1d0cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80c48b16e95007007bde4f8e65bc06af4dbb12e8c1c61a3a4d269b661c4af810"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / dd077494ccb0 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-28e401d68cc3403ba2779032165566252b6a8bab0b7a0981ec9d58c57c6f31cb)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-d60b035a4f3a37dcf83e447247958610c9067c4d5dc3232e5bd805ed3d462e55)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-ff1525bb94a846230eacf02d16a5d3fdada9782fb1b8e07e3db637e40c4f82c8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-085035890a888e3762ef6c21539aa4bea318c5d0d7e025c825486c4b52a3299c"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-67a8562bbaf07f925df7ac324ad4503a587de3b01ec498ceb66a1214afcaa75a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / dd077494ccb0 / 3

<a id="canonical-c44ea89329a0f77dc3c87009a0474722b52d04491f2e2c9aca777a20336e12e6"></a>

<a id="canonical-f0651dde418003ed7669d61a3ebbf20724458aede513e6c195d6a645be490be3"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / dd077494ccb0 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

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

<a id="canonical-505929013fc2037925c7c76e5c0f12788af48dc034856846d6880ac9b45fce60"></a>

<a id="canonical-e1b5ee16d302a32d6178cab22fc466266fc0269faaed74353176d7477cd3846a"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / dd077494ccb0 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-5044920e29f87090718e58298266a7c9b5b696d4cddd996af5572cfa3572515a"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / dd077494ccb0 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-2fa0b78f18a6be3b1e68a82809aea935c8c6d5bed5974012b0c19df523960e16)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b21c08f393ecb916b0858008857fe599b72cd2b66d2c221a6302aea6984f047"></a>

## ingress_egress_gw.inside_subnet — ingress_egress_gw.inside_subnet / 9b893dc7475f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.inside_subnet

<a id="canonical-420aacc24ff8ab865eefbc5bc007b11a3c034e8fc757507f75cba043cdfef8d9"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-4d8b229e110227deac407d4ad68113d406a3dadbb50221ffb58ceee3115630d9"></a>

## Direct properties — ingress_egress_gw.inside_subnet / 9b893dc7475f / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-40e9396956161e043f87d3e8ad07b5953de7132546412db744b113decc219557): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-b16016c465c1f716315223193811c3da6b4f8ccd28bb531a3b95497300088d65): complete subsection reference.

<a id="canonical-b9fc66dee15a9c834d4995c91c3c09c8d26b1a23d43ea69e4b23163449735b1b"></a>

## Next pages — ingress_egress_gw.inside_subnet / 9b893dc7475f / 4

- [ingress_egress_gw.inside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-40e9396956161e043f87d3e8ad07b5953de7132546412db744b113decc219557)
- [ingress_egress_gw.inside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-b16016c465c1f716315223193811c3da6b4f8ccd28bb531a3b95497300088d65)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-40e9396956161e043f87d3e8ad07b5953de7132546412db744b113decc219557"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9eb016d54ee286d505758add56f8a11f6d113319dcd7667dacb7725a997b3711"></a>

## ingress_egress_gw.inside_subnet.existing_subnet — ingress_egress_gw.inside_subnet.existing_subnet / 5b67e17cef81 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8)
- ingress_egress_gw.inside_subnet.existing_subnet

<a id="canonical-b73998122c553819b61386fc7eabb4a5adc8ae7d4eb5794a192f502fc5319506"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0a728b60f2b96f20ffa0ad5d258e3b5c82404dd03a76781058c07ccfbe0708c"></a>

## Direct properties — ingress_egress_gw.inside_subnet.existing_subnet / 5b67e17cef81 / 3

<a id="canonical-11237b3012e6657b15c1889e9bc896f85a98eb5cf585866f70ca5027f3aabc9b"></a>

<a id="canonical-51b6f9e1f5656133895e0edc92b1be795d360bc92fbfc7d78e3bf841c508128e"></a>

## subnet_name property — ingress_egress_gw.inside_subnet.existing_subnet / 5b67e17cef81 / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-550554539fec0b64c6f13960f9b1018783da2c2d567efa2ef0247ed2bf8a87fc"></a>

## Next pages — ingress_egress_gw.inside_subnet.existing_subnet / 5b67e17cef81 / 5

- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-b16016c465c1f716315223193811c3da6b4f8ccd28bb531a3b95497300088d65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dca8c1b215105d77a9ee1eb7fc1ca75d247c83303ed6d0385d970f0ed6bf5415"></a>

## ingress_egress_gw.inside_subnet.new_subnet — ingress_egress_gw.inside_subnet.new_subnet / 32c33be09902 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8)
- ingress_egress_gw.inside_subnet.new_subnet

<a id="canonical-6c1ccaf5f55d40aea33b0eb4a895f64e2d12649a07752fb2405d4a9fdb386154"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc66c7cb6a3b3f962264c420e37c7529a03b8e488fd589cc56ae43b1879e2f01"></a>

## Direct properties — ingress_egress_gw.inside_subnet.new_subnet / 32c33be09902 / 3

<a id="canonical-ff733535f87f2758f3e9b630d1c12d40eaef8f66bddd943f19d5f94346907893"></a>

<a id="canonical-c054313c1971ad75ebdc8f9291dec8df87bcce82cba1f76f6e3afcf2813c5b0f"></a>

## primary_ipv4 property — ingress_egress_gw.inside_subnet.new_subnet / 32c33be09902 / 4

Type: `"string"`. Optional.

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

<a id="canonical-fe5422ddc18180b47f3b141198fc8fe06f170ee9786976f59fcc9f4de5f94cba"></a>

<a id="canonical-ef6f40d949514b7d01a9fc279cffadc0d236776d1e12787132abba84d5f18462"></a>

## subnet_name property — ingress_egress_gw.inside_subnet.new_subnet / 32c33be09902 / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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

<a id="canonical-c47f5c63980ba448c5695cc7262089baffa5d1528a015776175bb82b40091819"></a>

## Next pages — ingress_egress_gw.inside_subnet.new_subnet / 32c33be09902 / 6

- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-ea05e40660d83d98f2dbfbb5ae503d60722845c76a1a87aeba90007f8e51a3d8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2b666e8da5e36ab6b073b61b99b4634dffe53fbf82fa8bd25a37681ca0e4fded"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4332acb4bcfcda04826e1c951c0b70b74eb4b36560beef07df9058cd3938b8e3"></a>

## ingress_egress_gw.no_dc_cluster_group — ingress_egress_gw.no_dc_cluster_group / a176cb120f81 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-54cc3c33b03c3cb3a886df18b96c8c80ff6a9107745a8fb7ab5a112f6867ef2c"></a>

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

<a id="canonical-7523a65f68f0a0d896c88d96e8c0cd1ef387bb7d9ea804e6dc605af3b8f984fd"></a>

## Direct properties — ingress_egress_gw.no_dc_cluster_group / a176cb120f81 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-660e90a183d8f26c1df5503480baec049baa585aceee2b6a2d3933c114d52cd3"></a>

## Next pages — ingress_egress_gw.no_dc_cluster_group / a176cb120f81 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-59f800c76a03ffeb99934404b60cf40395a134561fb9b764358e72709fae3e36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85a8a1bb2503891283784946d59a6383a3d8188918ce5896def8fd5a7e9ed628"></a>

## ingress_egress_gw.no_forward_proxy — ingress_egress_gw.no_forward_proxy / 31be2002200e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-d20d7af51f7feb3f13e5c8b9f2c44f98d42561acbfda71e323032b8f06f0180b"></a>

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

<a id="canonical-834f97d751c9b4052d78dd60696a31a55bb472b45b524f459c6d53decc549d81"></a>

## Direct properties — ingress_egress_gw.no_forward_proxy / 31be2002200e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b34a8d9e7bc85c129e9fd83577847b438ff8877b9555410db0bd7f09619e2244"></a>

## Next pages — ingress_egress_gw.no_forward_proxy / 31be2002200e / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d227621ac9c0988f24f88572c0db7090ca5032a95b2c1f8c34912aa5d344481f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b69a0f6f8663b2d7ce596bc1158b5c39f5656ed7a2e09942e151eaa6069ae916"></a>

## ingress_egress_gw.no_global_network — ingress_egress_gw.no_global_network / 2b9370d0ebc8 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_global_network

<a id="canonical-c63bd61d1bc31d7dff51c34e0f1254e3a3f51a6a10207372c6ed162ea80b3d6f"></a>

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

<a id="canonical-488746ad285be53db2b79d110f87cc383856a105a73b7459a62a0f302774596d"></a>

## Direct properties — ingress_egress_gw.no_global_network / 2b9370d0ebc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c43999cad2e2fd5112ac0542bf9a10fa5eea40690a221e90d31d3cb1fb6caec5"></a>

## Next pages — ingress_egress_gw.no_global_network / 2b9370d0ebc8 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-214d17241f00f1246d31aa08f82071976305d5c0539291623d4fe93adc1b2b67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-599b01533bf462b907641293f5e26dc289cd7abedd76d2e1c9887a117bd3bf8c"></a>

## ingress_egress_gw.no_inside_static_routes — ingress_egress_gw.no_inside_static_routes / edac218896cf / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-01912b556a4c47b59abb1efea77b039733661938973940dc313f390c520e83ae"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-82033e5b68d413f364b25d4617f517eacef75cfe1f9f951a7b6519f91265f83e"></a>

## Direct properties — ingress_egress_gw.no_inside_static_routes / edac218896cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ea0d2bf6d0ecc04b1c0bafeba981edccda91713a5d00a83536981dcdc448c7e"></a>

## Next pages — ingress_egress_gw.no_inside_static_routes / edac218896cf / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-72745e6670a4b81ab090ff751718a4bdc631007f55118552475a38b827747c74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a92d997955afc07bcabe7475deb1660b4c4f90400b2c281bc13001f537c7eead"></a>

## ingress_egress_gw.no_network_policy — ingress_egress_gw.no_network_policy / 07803923b210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_network_policy

<a id="canonical-3846f730f8f54159e62ed7ce86d63a3104a78f3191c44abc18e1e4897965a180"></a>

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

<a id="canonical-f3220b06c73113db65323eb5c5099d38b4ca8f504b56b5b5b777485f7d12d24f"></a>

## Direct properties — ingress_egress_gw.no_network_policy / 07803923b210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8200a185c8c52d990046b1c193329f96ff08a9afd882c0a1c350f0c133b7e9b4"></a>

## Next pages — ingress_egress_gw.no_network_policy / 07803923b210 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5a8f713e1bb71eb78b47b12e22c03b335b6c75f1aca76f789725413debc883a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1358d2d397aeb8800063fa164cc462c77503980c6e59b2a5adc48c0fdec1ad4"></a>

## ingress_egress_gw.no_outside_static_routes — ingress_egress_gw.no_outside_static_routes / 3ec961b068b9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-970607dd56e9eb8b1822d5c1967c3f4022244aa9c572ff34ce09e06a9365432c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-d1a460a19f967bfa44238ee691a0de50cdd96221e775a96bf5be79bb25a2e2cd"></a>

## Direct properties — ingress_egress_gw.no_outside_static_routes / 3ec961b068b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1659a7f158cd05bf41b18993406004c4f4a2ac4ac481f2e45174db18ff1627d7"></a>

## Next pages — ingress_egress_gw.no_outside_static_routes / 3ec961b068b9 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eee66d42c3872ad370688eb6dd93620dd150f390a88f62ad7d8d409925bbfcc3"></a>

## ingress_egress_gw.outside_network — ingress_egress_gw.outside_network / 0b68bac6cdbc / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.outside_network

<a id="canonical-e3d9fa2dcaa5ba7f8b5427ae8e29854aa4ac2c600b9212abd686b5fd5f232469"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
outside_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-57ee852ede9e46066e989ea8a9e58a539a17db3a6859d5c42b514ba7ffc3aea0"></a>

## Direct properties — ingress_egress_gw.outside_network / 0b68bac6cdbc / 3

- [existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-ef101b42f20a02b6670520d9b4bb0edd8769d12d5ae96cf19709fb924259c6c3): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-bc0b965c3be2e5b82abb032d1b18809c7cf939959e7df9297929d86f79379e3a): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-746347571399d962e59d048eaabe2bb422eabccc5b3b8b429b4a513deb7a213c): complete subsection reference.

<a id="canonical-64682f9a4db742a1f81af8178ae750aff596bbf0d21353c491f2900c501054ec"></a>

## Next pages — ingress_egress_gw.outside_network / 0b68bac6cdbc / 4

- [ingress_egress_gw.outside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-ef101b42f20a02b6670520d9b4bb0edd8769d12d5ae96cf19709fb924259c6c3)
- [ingress_egress_gw.outside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-bc0b965c3be2e5b82abb032d1b18809c7cf939959e7df9297929d86f79379e3a)
- [ingress_egress_gw.outside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-746347571399d962e59d048eaabe2bb422eabccc5b3b8b429b4a513deb7a213c)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ef101b42f20a02b6670520d9b4bb0edd8769d12d5ae96cf19709fb924259c6c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-485753553ac3ac302cb283ae53ec16864fe5f7c6ccf282d57ca5981eb16a5e4a"></a>

## ingress_egress_gw.outside_network.existing_network — ingress_egress_gw.outside_network.existing_network / 9a2c138f3923 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- ingress_egress_gw.outside_network.existing_network

<a id="canonical-21de07053b9ffb592d5074f7ac1aea74522765b8848723a38141afd5de67a6a2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6f030b04b31d503ab3c6b434b8bb0f0e03c5d69e466f9e8210511678d03af4f"></a>

## Direct properties — ingress_egress_gw.outside_network.existing_network / 9a2c138f3923 / 3

<a id="canonical-22e6afbbc6bcd329d9bdead2f84010e02dcb60de5473667c516a5aa4e8d39f60"></a>

<a id="canonical-ba2c7a3e311aa67b853040bf32ded7bd561a78fca6f2e62b7b12bdf9aa375c6e"></a>

## name property — ingress_egress_gw.outside_network.existing_network / 9a2c138f3923 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-04dcec601949b6a8ae64a9c8138a8b49cfd41bb5a37da2e705d061a090919e3e"></a>

## Next pages — ingress_egress_gw.outside_network.existing_network / 9a2c138f3923 / 5

- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-bc0b965c3be2e5b82abb032d1b18809c7cf939959e7df9297929d86f79379e3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a97b0a5e7797353c3c4af53623926d4f17c00d00562dca2fd7664e528412125"></a>

## ingress_egress_gw.outside_network.new_network — ingress_egress_gw.outside_network.new_network / f6c120672a5d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- ingress_egress_gw.outside_network.new_network

<a id="canonical-a16eb8cef119f498fa43d09bb0dd4e8bcd06c353b5fe91aab8580abc3c38c358"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-c057d12cb780e87977fa1c5066b1fb9f7747a9b18e374300a3a4c1bb9a914531"></a>

## Direct properties — ingress_egress_gw.outside_network.new_network / f6c120672a5d / 3

<a id="canonical-cd60a18b997ca50b4d503e1ec53e27809c02b2f4634bfe72462f90a163dbecc8"></a>

<a id="canonical-8ade2330f11f1555a14c5ae0559f6515f3ddc08ea6b130139b5f46945947e151"></a>

## name property — ingress_egress_gw.outside_network.new_network / f6c120672a5d / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-075545da4070aeeb54f1dbe2f31faab01c9e71162b923a23742d520b860aa71b"></a>

## Next pages — ingress_egress_gw.outside_network.new_network / f6c120672a5d / 5

- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-746347571399d962e59d048eaabe2bb422eabccc5b3b8b429b4a513deb7a213c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a59d22eb9cbc0014cc4ce02b462b6231c366c1311043e74cf75cf83866bb011"></a>

## ingress_egress_gw.outside_network.new_network_autogenerate — ingress_egress_gw.outside_network.new_network_autogenerate / f947518c3e34 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- ingress_egress_gw.outside_network.new_network_autogenerate

<a id="canonical-7456193481ed477d7e767a9ab6a4d8e376cc46a97a6b4503109682ed8e3c7d31"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
new_network_autogenerate = {}
```

<a id="canonical-88f0a4d69de9ebd5cf7e561e9bfa08659324092850c04d5ee5055680c8ff2d1d"></a>

## Direct properties — ingress_egress_gw.outside_network.new_network_autogenerate / f947518c3e34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-137a4b2c7bfe9ae2ec6f162673ffd466de585baf26840706c158b2394efbfd0b"></a>

## Next pages — ingress_egress_gw.outside_network.new_network_autogenerate / f947518c3e34 / 4

- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-656faf2ff6fc608d46979780060dec6b81ddef72cf1c9f4620732b4143eaffbe)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68202c1f62fdca18bac0ff7140aa4cabef9056bcaf5dc839826430b88e617b02"></a>

## ingress_egress_gw.outside_static_routes — ingress_egress_gw.outside_static_routes / aa85d456b769 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.outside_static_routes

<a id="canonical-5aeb2c18af6df8c821efc8419b22d8f17025852aab1a8a6a917ebb9288f88401"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ab6ae8754139230375c18dc5442b468b7cc0f1018912db936c4c037dc74455c"></a>

## Direct properties — ingress_egress_gw.outside_static_routes / aa85d456b769 / 3

- [static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43): complete subsection reference.

<a id="canonical-efdff6bf0e530b3ce7c4adfc4b28403b226aee5fffdeb363552019d580db7c59"></a>

## Next pages — ingress_egress_gw.outside_static_routes / aa85d456b769 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8991658a5f0fa90f7e4e1521cb379d2b2d22ceced3b9f5a1183f4dc4ed003cea"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — ingress_egress_gw.outside_static_routes.static_route_list / 8def865dc2f4 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-b92a885e9021fda1f44f828361d0f61d0c347e5b87f6a5502e614341ebafda4c"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-471bd610459dcc1f0c6f77db7eb768cd4ba94871d9cea22849e980a9d0d344e5"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list / 8def865dc2f4 / 3

- [custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d): complete subsection reference.

<a id="canonical-51c19c7a33174067acaef98b42f4eb5d04882a4824e200a94ea3d8ba0098cd1a"></a>

<a id="canonical-6a7e0e8b47a28b1bb55b7a031e214128483bc53c2ed530689307138f98810227"></a>

## simple_static_route property — ingress_egress_gw.outside_static_routes.static_route_list / 8def865dc2f4 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-9032cb7d93ed2dc8c9a5dc7341ddf1f75c13ea9a413aaff7183929a01da74ee1"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list / 8def865dc2f4 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
