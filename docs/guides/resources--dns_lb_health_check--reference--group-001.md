---
page_title: "xcsh_dns_lb_health_check reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check reference."
---

# xcsh_dns_lb_health_check reference

<a id="canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c23f5fc3efbf329d9dcce556166ceab4dfadd6a64b9e1683b521ad5a5e69a045"></a>

## Property reference — Property reference / 0287c1db4e34 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- Property reference

<a id="canonical-98b6a4624ff4484ee976df7c2f8a1264ee014a735a9097aeb62fa8f2f304fe1e"></a>

## Direct properties — Property reference / 0287c1db4e34 / 3

<a id="canonical-e769d16b2d0baaf5e4456fd681d5c8ad93ab6ce0637bb472ceb54eaefc96452f"></a>

<a id="canonical-84f0490598a896c9071ef96593b2e6379197ff56857ef774b2e60527282449a7"></a>

## annotations property — Property reference / 0287c1db4e34 / 4

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

<a id="canonical-294bc0a3a7da49ae0fcff546e2b04526b36515d7e55d56787e985d0ca19128a6"></a>

<a id="canonical-e97878eb5def5c097a815a5c36b055b1409814bce468984b95e3802de7d35666"></a>

## description property — Property reference / 0287c1db4e34 / 5

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

<a id="canonical-3f6ccf99ed55794a0ad8a080d7e543a63ab3fb356214c1c0c7ebbf2d6c13ed49"></a>

<a id="canonical-1b66e6935c46fbbcb0d2cae9bc320443b7785ba8767dcac8e8814d862388c68b"></a>

## disable property — Property reference / 0287c1db4e34 / 6

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

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276): complete subsection reference.

- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff): complete subsection reference.

- [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0a5cc03bbcab134b18406831bf50e7a703b390e38bdaa7337a769e841fcc3cd3): complete subsection reference.

<a id="canonical-0e67e11e1c6fc27f3484f4b51c2385a195f36368548174c79a9cbef8fc3b944b"></a>

<a id="canonical-c981939f0cc2ab3e39369fdfd96bafd991020b9ceafcbf168fd917cd812c2e9f"></a>

## id property — Property reference / 0287c1db4e34 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-64e62cf17d839612e9be7da57425432b704509c63b300256aa78a99c7da309ed"></a>

<a id="canonical-d88e40f9d13d85569151f0d40b518ab262181155764e26c02c162290257395e7"></a>

## labels property — Property reference / 0287c1db4e34 / 8

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

<a id="canonical-5ec36172a1c05fdd89937c8618cfe5661c8fdec99e6c32bc76f654a1cb49792b"></a>

<a id="canonical-926211818a49608c5c89859e11cc86fe59608162e1c930c7f1df890ec2f271eb"></a>

## name property — Property reference / 0287c1db4e34 / 9

Type: `"string"`. Required.

Name of the DNS LB Health Check. Must be unique within the namespace.

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

<a id="canonical-002808913dc26b85a21e8340e1925ca9ce66d9ba6585121109c7d89d8ee1337f"></a>

<a id="canonical-28ae69a9e6f67076f02cfb317de28321e9f21082b9e9259cb26aca7c0c1d76d9"></a>

## namespace property — Property reference / 0287c1db4e34 / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS LB Health Check. The F5 XC API restricts this resource to the system
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

- [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-62d6e7ac26e4c7c9cad22dddf6aa6a721fb74d4c331e6b386d2e77ed20bfeeaa): complete subsection reference.

- [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-f46afd433d2b9dc4eabb0947247d488e322f8b8d7fc0347a8cf0758386f83f8d): complete subsection reference.

- [timeouts](resources--dns_lb_health_check--reference--group-001.md#canonical-d00e4fab4b2ead7258d2c702796d55c08f8490f1b0cb92769d0bba26733b6eb5): complete subsection reference.

- [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4f552f352436d0463501c8a9924874bb92e802b0cb87f8b4e656e7a64759d079): complete subsection reference.

<a id="canonical-ad94adbf1dadf64f5434308cbc1d218406c58a9295891303de0fd8ec7246e95c"></a>

## All schema paths — Property reference / 0287c1db4e34 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_lb_health_check--reference--group-001.md#canonical-e769d16b2d0baaf5e4456fd681d5c8ad93ab6ce0637bb472ceb54eaefc96452f) |
| `description` | [description](resources--dns_lb_health_check--reference--group-001.md#canonical-294bc0a3a7da49ae0fcff546e2b04526b36515d7e55d56787e985d0ca19128a6) |
| `disable` | [disable](resources--dns_lb_health_check--reference--group-001.md#canonical-3f6ccf99ed55794a0ad8a080d7e543a63ab3fb356214c1c0c7ebbf2d6c13ed49) |
| `http_health_check` | [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-69722ebd9ca6b61b656ea86daf4db5738e582708a815e2ec77bb16996d7f7d22) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-dcc441173256a435493ec992adc786373931c8af5b0ddbefdc382d74c1aa6f6a) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-8da00bcf900bd71e7aae5c2387f12233033b2f2490a4d9a5577627f0f717be1f) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-a7cb3cd172cb0f54260459107c81cc363a7449cc20b28529b343bc4bbe0d7aed) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-e8b6e9c86b14d25e8522b43f95f1a79937a18795758f8fcdc97e903dacfd8993) |
| `http_health_check.receive` | [http_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-560eaf912949ddac55bdcf49e71c663b23deb540ce0590da16a8d0b9ed88d500) |
| `http_health_check.send` | [http_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-96f1cc8fba4a52964fbdbf8d7956b38787998520704c4aae19a62aa0ef39024b) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-38e840e58b7f5d3213810ddca82712f733fd6c727a41b9333928867309b259b7) |
| `https_health_check` | [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-99ef8107d88c125a41eeda3c2d40a44b6d1502dd18652c11e677e1d0a84e4234) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-213b535424c01336c807ec46e3514853bded606b2ba4dc43cb1f55d8c5417e9b) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-a7969ed616de2645ef437a348c98f2a6ec2950cd2574cd760bb8c129cbd82e66) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-c8947622decd414a88c3262b7d01b49b0179c9d6baad9d0c4dbd5db0faef5033) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-2d0aec5ab7ea9e9cd51471c5d5064c87b90220dc442d704308ad40d894166a74) |
| `https_health_check.receive` | [https_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-fa25f9db970567c4917b17e8698656fcb28ce786f47a33e9d51dd8ba680c7702) |
| `https_health_check.send` | [https_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-bf4551218a84f61d8f1709f19fb211690fb8ec36f872fd25383b163cd4906459) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-17aadebc2cf67c40f25c32f745e8f01d7a1fbcff0b4265e2dfad03a3e8845d91) |
| `icmp_health_check` | [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-747b36a1772fbb22d67f5d3dff25fe646dc0889aff3f72250d317b8f8d5ab367) |
| `id` | [id](resources--dns_lb_health_check--reference--group-001.md#canonical-0e67e11e1c6fc27f3484f4b51c2385a195f36368548174c79a9cbef8fc3b944b) |
| `labels` | [labels](resources--dns_lb_health_check--reference--group-001.md#canonical-64e62cf17d839612e9be7da57425432b704509c63b300256aa78a99c7da309ed) |
| `name` | [name](resources--dns_lb_health_check--reference--group-001.md#canonical-5ec36172a1c05fdd89937c8618cfe5661c8fdec99e6c32bc76f654a1cb49792b) |
| `namespace` | [namespace](resources--dns_lb_health_check--reference--group-001.md#canonical-002808913dc26b85a21e8340e1925ca9ce66d9ba6585121109c7d89d8ee1337f) |
| `tcp_health_check` | [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-6fe38de1ccd572c5fe9a116099d16e6c3dc8039616288970a79ba26ce955768a) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-5b55e8702961597dceb0d3b4dd7d7d5bd2b79a23127c01682070324d5389dbff) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-3cf03b452b058b14a3d26cff70a35a38c2c15d6074cc4d72ebec025aff73aad2) |
| `tcp_health_check.receive` | [tcp_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-bfd942a8828e0928b841cd8952d44312b0a2199d7202728a1f5bd405e86ac5a0) |
| `tcp_health_check.send` | [tcp_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-cfec44ccc890e12f3246a387739a9bf94078cc359c671f140b895c25d0cad863) |
| `tcp_hex_health_check` | [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-032b80dc6805884e9c78af54ff89a0bae813100dfab171e7d6f8bfcd53767d30) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-0bb5d47b6baf9041b4ba42f7916cb080c76f9afed150d33b5e499d33e7d152e7) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-9d3faed5e56582ba707618b22d3e494d10bc2e0d5d6b6e03da37523db7edb940) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-7ea3d59aea5434be1e7668e3e1aad03657f515363b40a091711534a49b7039d1) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-c5126f13fd357cc49207f3dd058e239997dec8c2a91afdfdc44590e15ba4b99b) |
| `timeouts` | [timeouts](resources--dns_lb_health_check--reference--group-001.md#canonical-a4885e3ee79a2bfc828adf9cbb6c9e713d242fb5637b803673e2e6cea3cd8bf5) |
| `timeouts.create` | [timeouts.create](resources--dns_lb_health_check--reference--group-001.md#canonical-02602c0019c460ace6fbaab5596f63a42dd150f0e8ed407dd1028f5376254600) |
| `timeouts.delete` | [timeouts.delete](resources--dns_lb_health_check--reference--group-001.md#canonical-0213f25a8219b90596bde145628d75ea77c6b58d2b25421d36b3d11e0bbd9f31) |
| `timeouts.read` | [timeouts.read](resources--dns_lb_health_check--reference--group-001.md#canonical-14a99351b1bcc59b1c3cdfb5253d8fb8fc4033886173fe7f915c370a37540f61) |
| `timeouts.update` | [timeouts.update](resources--dns_lb_health_check--reference--group-001.md#canonical-4b3bc119f599d76c863087903140e607497e81af969c21235ae3afd0d46e839c) |
| `udp_health_check` | [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1cf4d048b2c5640530bf66154532c26036793a0d9357c0ecb81caa0bac457233) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-05c78fd9db3878c348ac59754bcad157ae66755318f9e1d88ad2b0017fb4be51) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-ae803a569c10ad6915e1028fde12e74553406890f55b8889c02e3ca8e5af423f) |
| `udp_health_check.receive` | [udp_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-234e5e508d31193c8176de8a7b849d8b0ac4b078a800f272e99e7bdd4d4a1c68) |
| `udp_health_check.send` | [udp_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-374eb56cb52f3f8aeea34b8815cffd2ed332ead04c6ede5f7fe16342cfc093b6) |

<a id="canonical-260fc4cb5c545d5ee9af268bf9fda0831f518c6cfdcd2314f56066787f2d3b4a"></a>

## Next pages — Property reference / 0287c1db4e34 / 12

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff)
- [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0a5cc03bbcab134b18406831bf50e7a703b390e38bdaa7337a769e841fcc3cd3)
- [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-62d6e7ac26e4c7c9cad22dddf6aa6a721fb74d4c331e6b386d2e77ed20bfeeaa)
- [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-f46afd433d2b9dc4eabb0947247d488e322f8b8d7fc0347a8cf0758386f83f8d)
- [timeouts](resources--dns_lb_health_check--reference--group-001.md#canonical-d00e4fab4b2ead7258d2c702796d55c08f8490f1b0cb92769d0bba26733b6eb5)
- [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4f552f352436d0463501c8a9924874bb92e802b0cb87f8b4e656e7a64759d079)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4896a1beab74f5947fd2e75fc31a60af569504491a0b7be9abb6366bc19a0dd9"></a>

## http_health_check — http_health_check / 772fda65ecfa / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- http_health_check

<a id="canonical-69722ebd9ca6b61b656ea86daf4db5738e582708a815e2ec77bb16996d7f7d22"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "inherit_load_balancer_fqdn"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "virtual_host"),
  validators.ConflictingObjectAttributes("inherit_load_balancer_fqdn",
    "virtual_host")}
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
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-69722ebd9ca6b61b656ea86daf4db5738e582708a815e2ec77bb16996d7f7d22)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-99ef8107d88c125a41eeda3c2d40a44b6d1502dd18652c11e677e1d0a84e4234)
- [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-747b36a1772fbb22d67f5d3dff25fe646dc0889aff3f72250d317b8f8d5ab367)
- [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-6fe38de1ccd572c5fe9a116099d16e6c3dc8039616288970a79ba26ce955768a)
- [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-032b80dc6805884e9c78af54ff89a0bae813100dfab171e7d6f8bfcd53767d30)
- [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1cf4d048b2c5640530bf66154532c26036793a0d9357c0ecb81caa0bac457233)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-32d7078ec56aa8c369b79f00052aab2730b4043b4378d682885b0a6afb964289"></a>

## Direct properties — http_health_check / 772fda65ecfa / 3

- [disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-68dceb6242c319c0539e9efd3af541ee05e6ea2a56c88df839521a3c51b6853d): complete subsection reference.

<a id="canonical-8da00bcf900bd71e7aae5c2387f12233033b2f2490a4d9a5577627f0f717be1f"></a>

<a id="canonical-2dc449f2e1af052fd173391807054e506c1411f7f579c51bd02209ba455fb63e"></a>

## health_check_port property — http_health_check / 772fda65ecfa / 4

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a7cb3cd172cb0f54260459107c81cc363a7449cc20b28529b343bc4bbe0d7aed"></a>

<a id="canonical-18371e450080b64e87e9e35b7f66a877bb86fb4206f6a3f57e5bf37a6132f433"></a>

## health_check_secondary_port property — http_health_check / 772fda65ecfa / 5

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-f828ec3d191d255374eda124f8cbc7c67594b143db4f9b0231cfbc47091401ae): complete subsection reference.

<a id="canonical-560eaf912949ddac55bdcf49e71c663b23deb540ce0590da16a8d0b9ed88d500"></a>

<a id="canonical-c389fbe114d994a52bd6b3c8f34914352d7eacef02e13810333fe7a84cae4c33"></a>

## receive property — http_health_check / 772fda65ecfa / 6

Type: `"string"`. Optional.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-96f1cc8fba4a52964fbdbf8d7956b38787998520704c4aae19a62aa0ef39024b"></a>

<a id="canonical-66f131a80357f5f421708538191714974b06d6e5b936d319f8877a964b9d8ff2"></a>

## send property — http_health_check / 772fda65ecfa / 7

Type: `"string"`. Optional.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-38e840e58b7f5d3213810ddca82712f733fd6c727a41b9333928867309b259b7"></a>

<a id="canonical-353f1b4a022b6cd2fbe445457edd2bc3242cf8dcfc002684bfef3d79d2893520"></a>

## virtual_host property — http_health_check / 772fda65ecfa / 8

Type: `"string"`. Optional.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-05ad9f0d80fecca5b384686ec91936f40094979048927c9debc33ce49e97c392"></a>

## Next pages — http_health_check / 772fda65ecfa / 9

- [http_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-68dceb6242c319c0539e9efd3af541ee05e6ea2a56c88df839521a3c51b6853d)
- [http_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-f828ec3d191d255374eda124f8cbc7c67594b143db4f9b0231cfbc47091401ae)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-68dceb6242c319c0539e9efd3af541ee05e6ea2a56c88df839521a3c51b6853d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e95f5fc5139f10e3493c7ad677da1690ba3b9b90e80d8d1eba141908a772d457"></a>

## http_health_check.disable_virtual_host — http_health_check.disable_virtual_host / 12e13b5a2abc / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276)
- http_health_check.disable_virtual_host

<a id="canonical-dcc441173256a435493ec992adc786373931c8af5b0ddbefdc382d74c1aa6f6a"></a>

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
disable_virtual_host = {}
```

<a id="canonical-e12b37f26e9cfd6f5f72bd0fe73cacdd4e895b58824b3eecd2f45731c16e203f"></a>

## Direct properties — http_health_check.disable_virtual_host / 12e13b5a2abc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e46f35c7052b286c17ae23d168b69b7b027479828fe02b12c044a1e1a472574d"></a>

## Next pages — http_health_check.disable_virtual_host / 12e13b5a2abc / 4

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-f828ec3d191d255374eda124f8cbc7c67594b143db4f9b0231cfbc47091401ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-189fe07c89c061c6af97860ee49d8acdfe3588bd971c7a04fee90963761f9f37"></a>

## http_health_check.inherit_load_balancer_fqdn — http_health_check.inherit_load_balancer_fqdn / 95583bb5f6ea / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276)
- http_health_check.inherit_load_balancer_fqdn

<a id="canonical-e8b6e9c86b14d25e8522b43f95f1a79937a18795758f8fcdc97e903dacfd8993"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

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
inherit_load_balancer_fqdn = {}
```

<a id="canonical-dba32c9e1345b8618f768faae7bd7ab4a95dd0b8ccf62bb5115ac4dbd96aecc2"></a>

## Direct properties — http_health_check.inherit_load_balancer_fqdn / 95583bb5f6ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e8cb6ce9203a73e2ec244f3485e6231519db1930b68a6e505b08940fa22c94b"></a>

## Next pages — http_health_check.inherit_load_balancer_fqdn / 95583bb5f6ea / 4

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-dbc5f3c8c1d963d38b1087e666e57dc4b32d0ed00def0b620e85f7a18adeb276)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4705dc6ff7d942469df3571e821b428ea50328f9d8c467a2a3e29a1a11daa9b3"></a>

## https_health_check — https_health_check / d64ad9e312d7 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- https_health_check

<a id="canonical-99ef8107d88c125a41eeda3c2d40a44b6d1502dd18652c11e677e1d0a84e4234"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "inherit_load_balancer_fqdn"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "virtual_host"),
  validators.ConflictingObjectAttributes("inherit_load_balancer_fqdn",
    "virtual_host")}
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
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

Terraform syntax:

```terraform
https_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-166d6d92317e58c84ab393427d24924852585de808139cb4ae071d99b4461761"></a>

## Direct properties — https_health_check / d64ad9e312d7 / 3

- [disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-bd368643eea0003f0289c23bf26b9c34edb614ce06b2e16b693864e548a92922): complete subsection reference.

<a id="canonical-a7969ed616de2645ef437a348c98f2a6ec2950cd2574cd760bb8c129cbd82e66"></a>

<a id="canonical-36c615f5f78844875f7c2e180dc1bca38cddee92e70e6cc1ba9867f822fdea1b"></a>

## health_check_port property — https_health_check / d64ad9e312d7 / 4

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c8947622decd414a88c3262b7d01b49b0179c9d6baad9d0c4dbd5db0faef5033"></a>

<a id="canonical-5dc6eb9fc54000c1783e89d9bfe147e732262092410411a07c6a4b9f4f55a029"></a>

## health_check_secondary_port property — https_health_check / d64ad9e312d7 / 5

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-464be6788f4cbc0a610cd2aabcda2553eebaae7ad0f30e7b0cf1e7a3bf077c15): complete subsection reference.

<a id="canonical-fa25f9db970567c4917b17e8698656fcb28ce786f47a33e9d51dd8ba680c7702"></a>

<a id="canonical-3e05e27971fde11a0c28aa0322d71588af8127bd613bbe7a0d46bdf181b45ded"></a>

## receive property — https_health_check / d64ad9e312d7 / 6

Type: `"string"`. Optional.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-bf4551218a84f61d8f1709f19fb211690fb8ec36f872fd25383b163cd4906459"></a>

<a id="canonical-a30a9de69fa66e2dd1563bf56e03c2ec8d734051c4153af9ef011264de1b926e"></a>

## send property — https_health_check / d64ad9e312d7 / 7

Type: `"string"`. Optional.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-17aadebc2cf67c40f25c32f745e8f01d7a1fbcff0b4265e2dfad03a3e8845d91"></a>

<a id="canonical-8604e973f3c6cf7a82f2114560e2ce062adc65b91e01c82dced084fb712a70c3"></a>

## virtual_host property — https_health_check / d64ad9e312d7 / 8

Type: `"string"`. Optional.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-9916930dc8117f5bc5d4b32226cadbb83a68259d3d5a01114c6eeb4fb66e46b5"></a>

## Next pages — https_health_check / d64ad9e312d7 / 9

- [https_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-bd368643eea0003f0289c23bf26b9c34edb614ce06b2e16b693864e548a92922)
- [https_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-464be6788f4cbc0a610cd2aabcda2553eebaae7ad0f30e7b0cf1e7a3bf077c15)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-bd368643eea0003f0289c23bf26b9c34edb614ce06b2e16b693864e548a92922"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a576a89ac91ab16dea022550246ae65a8d1c3168d58fc54ce272894abc422f7e"></a>

## https_health_check.disable_virtual_host — https_health_check.disable_virtual_host / bc9ad344279a / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff)
- https_health_check.disable_virtual_host

<a id="canonical-213b535424c01336c807ec46e3514853bded606b2ba4dc43cb1f55d8c5417e9b"></a>

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
disable_virtual_host = {}
```

<a id="canonical-1c0629a59f070a90bd41b43721ca554cd458471046ca974e91f9782e6b2a0e60"></a>

## Direct properties — https_health_check.disable_virtual_host / bc9ad344279a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7195e78a3dea6e8305b387650e40dd258cc2a251721a1bcfe5aa36637d325de7"></a>

## Next pages — https_health_check.disable_virtual_host / bc9ad344279a / 4

- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-464be6788f4cbc0a610cd2aabcda2553eebaae7ad0f30e7b0cf1e7a3bf077c15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94dfbbfd21360163d666ade1f36420f404f2ab1ac731c0a347da710bf744d1ac"></a>

## https_health_check.inherit_load_balancer_fqdn — https_health_check.inherit_load_balancer_fqdn / f38912e16584 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff)
- https_health_check.inherit_load_balancer_fqdn

<a id="canonical-2d0aec5ab7ea9e9cd51471c5d5064c87b90220dc442d704308ad40d894166a74"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

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
inherit_load_balancer_fqdn = {}
```

<a id="canonical-756a1035fa1b6f836207f1d427df3d7ba6f634667d8fc473adc0057507e1f280"></a>

## Direct properties — https_health_check.inherit_load_balancer_fqdn / f38912e16584 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ad05a8396be8057cd13e39f4c8b83c8fdcd9735015fa66370107aa398f43697"></a>

## Next pages — https_health_check.inherit_load_balancer_fqdn / f38912e16584 / 4

- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-4764ffc1b598d0c9da6007f63538db0a0df0557cdc7e01a512f8aa1f906145ff)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-0a5cc03bbcab134b18406831bf50e7a703b390e38bdaa7337a769e841fcc3cd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7604773a156134df741485e03635fb09e5392a80469890ed67a621a64feec971"></a>

## icmp_health_check — icmp_health_check / 18470751d7b6 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- icmp_health_check

<a id="canonical-747b36a1772fbb22d67f5d3dff25fe646dc0889aff3f72250d317b8f8d5ab367"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

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
icmp_health_check = {}
```

<a id="canonical-8dbf487c99abbef51affed9cbd630c2f64841996369daa4082c20e56c5e44a8f"></a>

## Direct properties — icmp_health_check / 18470751d7b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d36bfc50c9b21b90ebfd75aec97d32cfb0a6c251f38c67c8bd9682ddf29d099a"></a>

## Next pages — icmp_health_check / 18470751d7b6 / 4

- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-62d6e7ac26e4c7c9cad22dddf6aa6a721fb74d4c331e6b386d2e77ed20bfeeaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a024bbcf278e156dd68f5acf36d87b940a0b10a579b8f5302820fdfc7bc1ed72"></a>

## tcp_health_check — tcp_health_check / 2db91436d14b / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- tcp_health_check

<a id="canonical-6fe38de1ccd572c5fe9a116099d16e6c3dc8039616288970a79ba26ce955768a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port")}
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-5cb106ed71d4dbcdc5b43aba55859fbefa07c54e3dc8892dbb70807365ae792d"></a>

## Direct properties — tcp_health_check / 2db91436d14b / 3

<a id="canonical-5b55e8702961597dceb0d3b4dd7d7d5bd2b79a23127c01682070324d5389dbff"></a>

<a id="canonical-839b9a3d8dd6db888ca0962e06e8addcabdf5576036172c3c311afe1650d8429"></a>

## health_check_port property — tcp_health_check / 2db91436d14b / 4

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3cf03b452b058b14a3d26cff70a35a38c2c15d6074cc4d72ebec025aff73aad2"></a>

<a id="canonical-bccdf88892aed5daba4f531c5636bc389cd7057dd95eb0e5eecbcbc38411fc49"></a>

## health_check_secondary_port property — tcp_health_check / 2db91436d14b / 5

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-bfd942a8828e0928b841cd8952d44312b0a2199d7202728a1f5bd405e86ac5a0"></a>

<a id="canonical-9681be632aad651c0af2d5b0436203d0acfe18326cad0f938d5398b92b0086fd"></a>

## receive property — tcp_health_check / 2db91436d14b / 6

Type: `"string"`. Optional.

Regular expression used to match against the response to the monitor's request. Mark node up upon
receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-cfec44ccc890e12f3246a387739a9bf94078cc359c671f140b895c25d0cad863"></a>

<a id="canonical-d9bf91bb942217d640e0d56f36850896db478972f669e84fd4fb868415b7a12c"></a>

## send property — tcp_health_check / 2db91436d14b / 7

Type: `"string"`. Optional.

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS).

Upstream description:

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3b792f8e830945901f07790a9ad13001abbcaf83966d0be350bcbeb4573dd325"></a>

## Next pages — tcp_health_check / 2db91436d14b / 8

- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-f46afd433d2b9dc4eabb0947247d488e322f8b8d7fc0347a8cf0758386f83f8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4486000bc995e47e36a0f1b69d92ab2639e60e742b1dc5e371ed9dad592067f"></a>

## tcp_hex_health_check — tcp_hex_health_check / 1f34c4f2e120 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- tcp_hex_health_check

<a id="canonical-032b80dc6805884e9c78af54ff89a0bae813100dfab171e7d6f8bfcd53767d30"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp hex health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port")}
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
tcp_hex_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-6f361766b77f300ddef117c24e12136a77f6224e00d2a422e7e1779e226323d0"></a>

## Direct properties — tcp_hex_health_check / 1f34c4f2e120 / 3

<a id="canonical-0bb5d47b6baf9041b4ba42f7916cb080c76f9afed150d33b5e499d33e7d152e7"></a>

<a id="canonical-03ac4f28abc612e21653cd5fe27157fe73589f5876454c4da940ce7c360c23f3"></a>

## health_check_port property — tcp_hex_health_check / 1f34c4f2e120 / 4

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-9d3faed5e56582ba707618b22d3e494d10bc2e0d5d6b6e03da37523db7edb940"></a>

<a id="canonical-783e6be97e652bb7cd25e7526244e0437563615e984811599025295c7ce6b1bf"></a>

## health_check_secondary_port property — tcp_hex_health_check / 1f34c4f2e120 / 5

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7ea3d59aea5434be1e7668e3e1aad03657f515363b40a091711534a49b7039d1"></a>

<a id="canonical-69e8b434f166331f7a79edad225dc2e6bca13bf19e765ff0854895e53ad13ac4"></a>

## receive property — tcp_hex_health_check / 1f34c4f2e120 / 6

Type: `"string"`. Optional.

Hex encoded raw bytes expected in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-c5126f13fd357cc49207f3dd058e239997dec8c2a91afdfdc44590e15ba4b99b"></a>

<a id="canonical-977c369602682467fc93e2ad6d4354319ded1b4c4c179227374b014508a99353"></a>

## send property — tcp_hex_health_check / 1f34c4f2e120 / 7

Type: `"string"`. Optional.

Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-a29eb09fe63c9e375d1e90255b316230c63539dc54d23e407180a6dea1466abe"></a>

## Next pages — tcp_hex_health_check / 1f34c4f2e120 / 8

- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-d00e4fab4b2ead7258d2c702796d55c08f8490f1b0cb92769d0bba26733b6eb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e64041c83749b339e3681b1c9fa046b9210eab9c9af517879061f588e80fffb3"></a>

## timeouts — timeouts / 7021a76dfbfb / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- timeouts

<a id="canonical-a4885e3ee79a2bfc828adf9cbb6c9e713d242fb5637b803673e2e6cea3cd8bf5"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-42ba42ea7606ebde7fd88bd4bfa788e89818cd99009d041d082f14ecc2dc9df0"></a>

## Direct properties — timeouts / 7021a76dfbfb / 3

<a id="canonical-02602c0019c460ace6fbaab5596f63a42dd150f0e8ed407dd1028f5376254600"></a>

<a id="canonical-33570d94fc6ecafe6abd7997af7e53ded37e8fa7cfb0726c2c9a76010d81f604"></a>

## create property — timeouts / 7021a76dfbfb / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0213f25a8219b90596bde145628d75ea77c6b58d2b25421d36b3d11e0bbd9f31"></a>

<a id="canonical-be2df12e17ac05dcdf83bc09f095dd3eb4534238f6a1ee7945eee65d444001dd"></a>

## delete property — timeouts / 7021a76dfbfb / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-14a99351b1bcc59b1c3cdfb5253d8fb8fc4033886173fe7f915c370a37540f61"></a>

<a id="canonical-582324b0d459d4b950e591315947e41cc7d59fab6f03d2e9efe7340f01a91c36"></a>

## read property — timeouts / 7021a76dfbfb / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-4b3bc119f599d76c863087903140e607497e81af969c21235ae3afd0d46e839c"></a>

<a id="canonical-e14b18b571724031a6fe7cf506f769dcc4afd495e033f11dd7bf7d06f5825479"></a>

## update property — timeouts / 7021a76dfbfb / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d74c78cd7fc58227c7c01c09e8ad1f8cdf9e036d492de3e586b654c8d1586520"></a>

## Next pages — timeouts / 7021a76dfbfb / 8

- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)

<a id="canonical-4f552f352436d0463501c8a9924874bb92e802b0cb87f8b4e656e7a64759d079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d6ceb443f3de091bec6115a2a3abe1cbaa683ace5fe607dbd099c1302933877"></a>

## udp_health_check — udp_health_check / dbad9aa20614 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- udp_health_check

<a id="canonical-1cf4d048b2c5640530bf66154532c26036793a0d9357c0ecb81caa0bac457233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for udp health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port",
    "receive",
    "send")}
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
udp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-40765dd3c93ee74c191bbf57dc2fca7cd16fcbdbb3346f9c95ea32ac8fdd9f6f"></a>

## Direct properties — udp_health_check / dbad9aa20614 / 3

<a id="canonical-05c78fd9db3878c348ac59754bcad157ae66755318f9e1d88ad2b0017fb4be51"></a>

<a id="canonical-3f076010fec57ebaae402c3f08c4ceb6fca0a71682695818d6ff221e78de7ee8"></a>

## health_check_port property — udp_health_check / dbad9aa20614 / 4

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ae803a569c10ad6915e1028fde12e74553406890f55b8889c02e3ca8e5af423f"></a>

<a id="canonical-d23b9a88500da097b5bcc1e9832182ea2771f27e71898e4c10bdc0c0eef05c23"></a>

## health_check_secondary_port property — udp_health_check / dbad9aa20614 / 5

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-234e5e508d31193c8176de8a7b849d8b0ac4b078a800f272e99e7bdd4d4a1c68"></a>

<a id="canonical-2171d4502609677ba0dbb11ecfa83daa023e0d6817ad91aa527714ce60c44b3f"></a>

## receive property — udp_health_check / dbad9aa20614 / 6

Type: `"string"`. Optional.

UDP response to be matched. It can be a regex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-374eb56cb52f3f8aeea34b8815cffd2ed332ead04c6ede5f7fe16342cfc093b6"></a>

<a id="canonical-e914a081516ce67f6bc6579db8ca49fd589578dde08b33c4d5ac8e5e7e7c08a8"></a>

## send property — udp_health_check / dbad9aa20614 / 7

Type: `"string"`. Optional.

Send String. UDP payload.

Upstream description:

UDP payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-901251767a9458b4ab06f89cd5dddbcbe9c4c96e59ef35555f02f75711a17f70"></a>

## Next pages — udp_health_check / dbad9aa20614 / 8

- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-5b05d1829252c107baca6b4d4a3a4f17cdca35a4df305b72cc1d45e89be3d8d0)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-3201644db8dffc1140cc9cf2b5f50daeb995ff6e2c8f28ac0c33e6260193837f)
