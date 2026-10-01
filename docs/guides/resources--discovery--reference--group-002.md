---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-8252af1ebdc675382085b306ff562495560e584baf5d5b960b19589995f40615"></a>

## where.virtual_network — where.virtual_network / aac6d3597e77 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- where.virtual_network

<a id="canonical-2ece54f6fd934b2eafca7f5c13918d65008786bec0c2dfe98242bfddb552cc0a"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-82e2e15809ce592890ef0d80a5c762e3f4569294c26073e7bf14f6c9986184e1"></a>

## Direct properties — where.virtual_network / aac6d3597e77 / 3

- [ref](resources--discovery--reference--group-002.md#canonical-f6b5eccb37e429185038e7e4aa6af564016847bd13df8c8ac68ace3696455d41): complete subsection reference.

<a id="canonical-8343a6c49af331f6e51ae99acdef118b220707ceb5a664c052dcbe893fc0ae43"></a>

## Next pages — where.virtual_network / aac6d3597e77 / 4

- [where.virtual_network.ref](resources--discovery--reference--group-002.md#canonical-f6b5eccb37e429185038e7e4aa6af564016847bd13df8c8ac68ace3696455d41)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-f6b5eccb37e429185038e7e4aa6af564016847bd13df8c8ac68ace3696455d41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8585b94a9a1ebdcf22b90f3ffd39dfc9b6dd01086e465288c5260c15be96b5d2"></a>

## where.virtual_network.ref — where.virtual_network.ref / 1e53065906ab / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.virtual_network](resources--discovery--reference--group-001.md#canonical-5460d6ee537cfa70bf28623e96e3dcce1005fe56ad7380a98fe2ea1c0d95b937)
- where.virtual_network.ref

<a id="canonical-3730ae9c211925657b8eb24be93addcb953c14c15ebba06a55d6a6c63ec31033"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb2c47f28387ed360e59d635339c97abaea177c9307631f21e11e67cb6cfca87"></a>

## Direct properties — where.virtual_network.ref / 1e53065906ab / 3

<a id="canonical-b4ceb9b83950c4062638d50c5c7358996cd8e49b9d3ef3986f843c86a751f328"></a>

<a id="canonical-1cc59b3499a3ac0a1769516f466dcf28e2acd8ac78fee582283965d34779e20c"></a>

## kind property — where.virtual_network.ref / 1e53065906ab / 4

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

<a id="canonical-519d934c44d79474813ee05c115f865045009d6208eee3dfd09f13acfe798548"></a>

<a id="canonical-d2a2eae2a4bc608142f85783d39b7f6c75de6b4dab480bce850624abaffad4b2"></a>

## name property — where.virtual_network.ref / 1e53065906ab / 5

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

<a id="canonical-acf92773b2f19c108dfe9a82e9806556a304918796aa2c37903e9c68fad8fe31"></a>

<a id="canonical-64bf2cc3bb586f5867aef4eef52bc35d4e7b98fd01c7429ad88e18681abd21e3"></a>

## namespace property — where.virtual_network.ref / 1e53065906ab / 6

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

<a id="canonical-b3523db29b324793fc08601b37bc2a7e9025156297ed5df8353af5e71febe1ea"></a>

<a id="canonical-090d5206cc51e58edd2f7d6c6939c432e0c27d0ae465d9df49a65b636570d2e0"></a>

## tenant property — where.virtual_network.ref / 1e53065906ab / 7

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

<a id="canonical-7c3722cc4833137ed8a79c8098258d784f03eb515d2aa1415fe620d261359dfb"></a>

<a id="canonical-3715953be5ce15f17a17ad372d90c0c890ff331090df227ffd8a349c96d0aa7d"></a>

## uid property — where.virtual_network.ref / 1e53065906ab / 8

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

<a id="canonical-bb52c5a6c90ad5946f808083f1cca8c5d1eefdcc5f80df3b3b2b656d7afdc7e1"></a>

## Next pages — where.virtual_network.ref / 1e53065906ab / 9

- [where.virtual_network](resources--discovery--reference--group-001.md#canonical-5460d6ee537cfa70bf28623e96e3dcce1005fe56ad7380a98fe2ea1c0d95b937)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5787fded32d860a3d3bea4743f953a592d12dc5cd0bc5c6a66144f2347f4ff4"></a>

## where.virtual_site — where.virtual_site / 4d37c6ee2d68 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- where.virtual_site

<a id="canonical-3a7c0709bfda6a9508467b4b8cd2fa470834d012c85c9cc35844050e1408c77e"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-b333725bbc32b6012f68c7bf64f2b27f418510d2fb466b0f9e7d66148a5183ac"></a>

## Direct properties — where.virtual_site / 4d37c6ee2d68 / 3

- [disable_internet_vip](resources--discovery--reference--group-002.md#canonical-06002e26540ccfe1a07bbc550d0fa2be9dd1d41a4848c9e97a16798c26bdc040): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-002.md#canonical-307b45602219e490d04d299b374b0a7fd2f74e539948f0a095f1737b45944084): complete subsection reference.

<a id="canonical-8076af3efeec3f8ea9ec03bd235873ecdad1362c924da96beae5d6b31f140e37"></a>

<a id="canonical-484f7c9d5296752c2bf22cec484d2b9526bab9d73c852a9c1c9d3a5517af5546"></a>

## network_type property — where.virtual_site / 4d37c6ee2d68 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--discovery--reference--group-002.md#canonical-b74a19695bf84a9edc0955ff74b34c895ebc460c9cb08f5ffdc4e1c3f41260ba): complete subsection reference.

<a id="canonical-ed47b957b199d7144fbf7d79c441c0944029874a7789610940bf6df01d08b8f7"></a>

## Next pages — where.virtual_site / 4d37c6ee2d68 / 5

- [where.virtual_site.disable_internet_vip](resources--discovery--reference--group-002.md#canonical-06002e26540ccfe1a07bbc550d0fa2be9dd1d41a4848c9e97a16798c26bdc040)
- [where.virtual_site.enable_internet_vip](resources--discovery--reference--group-002.md#canonical-307b45602219e490d04d299b374b0a7fd2f74e539948f0a095f1737b45944084)
- [where.virtual_site.ref](resources--discovery--reference--group-002.md#canonical-b74a19695bf84a9edc0955ff74b34c895ebc460c9cb08f5ffdc4e1c3f41260ba)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-06002e26540ccfe1a07bbc550d0fa2be9dd1d41a4848c9e97a16798c26bdc040"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f006d158d58ef36bc859bb19336028c1c8ee787116deec7b10a67d92a0fa607"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / bf934f72f825 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- where.virtual_site.disable_internet_vip

<a id="canonical-b645ba918516140ea91be6f1c272ceb17ec8803f729f50f5ecf785b4bce03754"></a>

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
disable_internet_vip = {}
```

<a id="canonical-b84afd6c451b192411afea69c3a1cfb840d36cfac009b8a8eccb220c3f9a21d9"></a>

## Direct properties — where.virtual_site.disable_internet_vip / bf934f72f825 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8091d2020b7addf7551f7be41a6e0281a2d4a0c20aec4b8bcf8598966955a17"></a>

## Next pages — where.virtual_site.disable_internet_vip / bf934f72f825 / 4

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-307b45602219e490d04d299b374b0a7fd2f74e539948f0a095f1737b45944084"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61d4d664270fa7522eb565f3bd2810f46333ef86963e653195b1cbacf26645fd"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / a73fef2fe2ec / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- where.virtual_site.enable_internet_vip

<a id="canonical-c908a61b4a10e6895437f558ace8c3cab83cf2ffbc354952fd2c6bc757ba44e5"></a>

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
enable_internet_vip = {}
```

<a id="canonical-bea66d856a107ed3c056b2b366be5d543d4461dc61b6d90fe1017d9014c35c4c"></a>

## Direct properties — where.virtual_site.enable_internet_vip / a73fef2fe2ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90c961c6c098d7f3fe4d8b231bd5854127f3298478ede05695ffabbea4d3597d"></a>

## Next pages — where.virtual_site.enable_internet_vip / a73fef2fe2ec / 4

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-b74a19695bf84a9edc0955ff74b34c895ebc460c9cb08f5ffdc4e1c3f41260ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d1339a50b617b0cffb7909d9c44671f15b7dc5816c9856422a0909a70859565"></a>

## where.virtual_site.ref — where.virtual_site.ref / f453e44527b9 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- where.virtual_site.ref

<a id="canonical-ab0bb56e76fbffed5041e5a906ff321ad2f815596ecd3f69acf8b6c1cd53772b"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-8519eeaff905bce350123f335f4d96d8d94651b9b7eda2fc385a692b3d0a9eab"></a>

## Direct properties — where.virtual_site.ref / f453e44527b9 / 3

<a id="canonical-b4beb48ee0b624863e488aa2b99e3b4f1dc247a99ad2c812e72ec5d69bff8600"></a>

<a id="canonical-df758fc578bab085f974d7c9752380c5b571e33c720ed0e8cc5bb3db2c4a86c7"></a>

## kind property — where.virtual_site.ref / f453e44527b9 / 4

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

<a id="canonical-b9e4c4457ce1d2c73fe641ca401615f1703e6680e7eeee50248f1ec52a00d050"></a>

<a id="canonical-ad8dd47b025ccee290e3310189aacb1b038b7348e2d2124e449d76a6d2bb6699"></a>

## name property — where.virtual_site.ref / f453e44527b9 / 5

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

<a id="canonical-747ef68cf56ab4656d7c06e46f0c2b742e366891005620ae077be3e48285425f"></a>

<a id="canonical-e7ff054a1ca6dd23f88a1b12e7de3aac622edfd8d9eb3c80bae94b586308a1cf"></a>

## namespace property — where.virtual_site.ref / f453e44527b9 / 6

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

<a id="canonical-1c1be349c43bbb104ae6aad5cff1887f77ffbdfc1d40ee238e271d8c05b6f95e"></a>

<a id="canonical-3f97f432e7d5644f09d11f7f869ba7f0a23e07262d5ee9ca11f6dd1bec015da2"></a>

## tenant property — where.virtual_site.ref / f453e44527b9 / 7

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

<a id="canonical-e5ac516569afa27073cd3e0f532acc6af811c42d11daf6088c10659e58bd3d1f"></a>

<a id="canonical-ec7130db8812cef761ddf6d5742969356cbb90288dec2f034605b5cb272e811b"></a>

## uid property — where.virtual_site.ref / f453e44527b9 / 8

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

<a id="canonical-32eefceab29a51bc9b22468931c021893d3d96bd5c8281db149689b9d71020bd"></a>

## Next pages — where.virtual_site.ref / f453e44527b9 / 9

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
