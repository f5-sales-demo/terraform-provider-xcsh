---
page_title: "xcsh_dns_lb_health_check reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check reference."
---

# xcsh_dns_lb_health_check reference

<a id="canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4585a065fe3e9fa91a5bae38eb4018028e4b31bb970a952ea064f586e46dfe7c"></a>

## Property reference — Property reference / 677197c49194 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- Property reference

<a id="canonical-79ef6e366c8190a99731f515e69f7b5a8113a52431d936c45baeef81a58a48e8"></a>

## Direct properties — Property reference / 677197c49194 / 3

<a id="canonical-9d34dc1b7ce21f7ab2865aaac6a4ef219c485b4eea7ebe2b2e17b90b58351a83"></a>

<a id="canonical-8a0a0e06b46e90e2202adfa1c05b843a09c7ad7833700a10da0c85ce549c6abf"></a>

## annotations property — Property reference / 677197c49194 / 4

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

<a id="canonical-60d77bbea7bab250f326c6b1db116a7dcf4c964531550e64ec3fefd6047cd898"></a>

<a id="canonical-74b2460e10519b5b1319f44d5304ff7754e443fe2625be2319ac82078374cc32"></a>

## description property — Property reference / 677197c49194 / 5

Type: `"string"`. Computed.

Description of the DNSLBHealthCheck.

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

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886): complete subsection reference.

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496): complete subsection reference.

- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-8d079515e1ebee028ec1536f7db33c9eb2c85a624380f5a5c8ff91dc69e65a4f): complete subsection reference.

<a id="canonical-6f5ca81ba2c0856024e50f0af7416330fc54f60c075de9b4e07d0fa5f47af062"></a>

<a id="canonical-12dca0a0b484b660a20c8e97868ebe947e73e258f9648ec3248a45c77e3bea28"></a>

## id property — Property reference / 677197c49194 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-15b4ab025a02bd2604fcb91d540cb9166de6c824a204ae654ef98d4ffae0c71a"></a>

<a id="canonical-6753231ecdbc4960f38378069a008fad34e0c542e9fc92d3226325f2fac622d7"></a>

## labels property — Property reference / 677197c49194 / 7

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

<a id="canonical-a59295600de7db9e5a1ac7f56a86f7464f5ea9c652a4c7564eed4559ccc41ab6"></a>

<a id="canonical-7291a6d5f056ff7fea4f6b7608f38f7ee02cb29e2894c06c9a01dee906c19b8b"></a>

## name property — Property reference / 677197c49194 / 8

Type: `"string"`. Required.

Name of the DNSLBHealthCheck.

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

<a id="canonical-3f23d8e6d2cde6a536a43f29290ef51d954f55dca402960a3b39132cb4ad6adc"></a>

<a id="canonical-99104314f8d8140b960bfdc2f1ddc24ca608060b5982cc135647a6365037afd5"></a>

## namespace property — Property reference / 677197c49194 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBHealthCheck exists.

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

- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-f6ebc80c6ee49cd5d14fd5e59781d78a5ad266fd7e36e4fc06c4bf49d3e2d37b): complete subsection reference.

- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-b37c57860ec3c66a1ce578d5df4f2e81fd285eecb0ca6dd0f33b6fb66bc7763d): complete subsection reference.

- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-4350c7f6bfdc26288aaa3248f59edf2895c799e100f59cd0269003ce43fdf936): complete subsection reference.

<a id="canonical-93835146f5e1044482044c59baef42113837deeae6a31f5cb6c1c61a51aeb312"></a>

## All schema paths — Property reference / 677197c49194 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_lb_health_check--reference--group-001.md#canonical-9d34dc1b7ce21f7ab2865aaac6a4ef219c485b4eea7ebe2b2e17b90b58351a83) |
| `description` | [description](data-sources--dns_lb_health_check--reference--group-001.md#canonical-60d77bbea7bab250f326c6b1db116a7dcf4c964531550e64ec3fefd6047cd898) |
| `http_health_check` | [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-fc0355eccd0aa615b6462a2e3f766b87c3b9752b4e917f3a4f7c8c2da436ea9d) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1e01b296d8eb98e015531884fa38d9084e5a7a899e760ceb23203f4a7c2a0345) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-eb44e716159ba5c17f23d6119020f4baceea2002638531cd576e3a0051827d84) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-da74b22a8d9b8845443a2a3cf3ac6c44fd7675c11dedbaaba0842b7deb45413b) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-322d5223f7b0f6378da52e63729ee7f7c32caf060ba40d6f07c1bdf249c5efeb) |
| `http_health_check.receive` | [http_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-147240bcdc7de32d010b6ef038d8b8694c23129ca0eafa24f0aa26029fe9f244) |
| `http_health_check.send` | [http_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d45a3a8c1814167ee06d93ea75a852460f98953ae50e30555fddf3aade84428f) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-9bd866e427f96814552f0976651ee285b2188a7e8940ae5ac01b3395151f83b1) |
| `https_health_check` | [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-620872d5c1c0ca39d9a598ff368b4290d4abacd8331dca49743b2449bcccd3c6) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-43a42e5f393732bbc8dc9738d86b0b8688669f37178a73acbe86ab31936721ee) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-c6620e204e68cdc9e6349ed48a69333e297ca097123013a448c26e320ce178e4) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1ca6a82997c2808b75260192f0a20b6b3f340afaf0b43f3e5e7bc11e93104f89) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-b4e5d93b59ae866a99d04abd958272c5bac7eaff461b25ea3ee5d129daf8e7bf) |
| `https_health_check.receive` | [https_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-bf8581d5b1029d87942aada10ebecd4077ef74d66a6101f6b3f0a8a4a1406c6e) |
| `https_health_check.send` | [https_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-c10436e758ab723abea01e04ddd9bce65f94d2dbbdf54b1e339f62bee0f30bcb) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-c0354958b85c0dbab258d9e09c3c1cc4ec9c0a669d78ae9e0ceaa35048cd7639) |
| `icmp_health_check` | [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-79b79c14af5195b5f99d7ff0e15ca4917d29f81d8133a26a3c41ae995b4dee13) |
| `id` | [id](data-sources--dns_lb_health_check--reference--group-001.md#canonical-6f5ca81ba2c0856024e50f0af7416330fc54f60c075de9b4e07d0fa5f47af062) |
| `labels` | [labels](data-sources--dns_lb_health_check--reference--group-001.md#canonical-15b4ab025a02bd2604fcb91d540cb9166de6c824a204ae654ef98d4ffae0c71a) |
| `name` | [name](data-sources--dns_lb_health_check--reference--group-001.md#canonical-a59295600de7db9e5a1ac7f56a86f7464f5ea9c652a4c7564eed4559ccc41ab6) |
| `namespace` | [namespace](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3f23d8e6d2cde6a536a43f29290ef51d954f55dca402960a3b39132cb4ad6adc) |
| `tcp_health_check` | [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2827883a91f2fcb6edf33a763428e3c6a48f8a21bc98fc4a0f59b019e18038a3) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-58d0664c31c0693d6c444e91a8fdbedfbd2f054ddc676839a251e135bce1809c) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-02487b41b6b4eb85e2677e4d4bb72413b5fbc9ca5adce9f4fab9465a90cdce53) |
| `tcp_health_check.receive` | [tcp_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3828f30000447cf39ddf5002132576a0915ab9c47bda463b7255eaf74ea0cd73) |
| `tcp_health_check.send` | [tcp_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-a204f03504e4fe0f02a310d78db091a521cec72f76482bfffcc638dcc2abe844) |
| `tcp_hex_health_check` | [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-199046536be55acd4bc1f1b45f6aba1ff44a4a821642cd1c2b8f4affe9017a19) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-211e79b65b2f4adb476edbaeb411a446dcbc705be6919e698de3f03a2536f7cc) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-ee9aa321e24e5ba4626fa82b8a4b05c01f0c7d82cae1790f8d3642eb71cf20c9) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-302955bbf60896b7829c1249f4d0faa49103099b9692c80791eae463f5190111) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-c2bb2a57fa1ed448ff75f82c27f18dabfa7a3faa06c3cd1c5675649b76847b56) |
| `udp_health_check` | [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3ecc01bc17e2f94797a61d881c12432b5829b2232409dfc2b7345ab05d702508) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-419afb4242ddc7102096a7247d02071f983297f493620d32f0d301becc3f366d) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-ff504e7c2d3ecb60252c331fd6dd5df7c00106915483204b35bca289a9362c77) |
| `udp_health_check.receive` | [udp_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2cb5dfe10f2b52ee6a95f3350d6d6425d17c49b55438322bff3758d0a2e9071f) |
| `udp_health_check.send` | [udp_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-c26c971da87f5e1a0e27ba7ea704e68d963fbe07cd27060d61c8e64049669ffd) |

<a id="canonical-17afd0f637ce9092169b161dbce7c1b2266cf974e6e2d847a524f35856ab1f06"></a>

## Next pages — Property reference / 677197c49194 / 11

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496)
- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-8d079515e1ebee028ec1536f7db33c9eb2c85a624380f5a5c8ff91dc69e65a4f)
- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-f6ebc80c6ee49cd5d14fd5e59781d78a5ad266fd7e36e4fc06c4bf49d3e2d37b)
- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-b37c57860ec3c66a1ce578d5df4f2e81fd285eecb0ca6dd0f33b6fb66bc7763d)
- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-4350c7f6bfdc26288aaa3248f59edf2895c799e100f59cd0269003ce43fdf936)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-231cd7729ff9502a378060d274e18572ca72ef13ea9b26337598d41dfbf8fe99"></a>

## http_health_check — http_health_check / 883c5cca92e3 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- http_health_check

<a id="canonical-fc0355eccd0aa615b6462a2e3f766b87c3b9752b4e917f3a4f7c8c2da436ea9d"></a>

Type: `"single"`. Computed.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

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

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-fc0355eccd0aa615b6462a2e3f766b87c3b9752b4e917f3a4f7c8c2da436ea9d)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-620872d5c1c0ca39d9a598ff368b4290d4abacd8331dca49743b2449bcccd3c6)
- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-79b79c14af5195b5f99d7ff0e15ca4917d29f81d8133a26a3c41ae995b4dee13)
- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2827883a91f2fcb6edf33a763428e3c6a48f8a21bc98fc4a0f59b019e18038a3)
- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-199046536be55acd4bc1f1b45f6aba1ff44a4a821642cd1c2b8f4affe9017a19)
- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3ecc01bc17e2f94797a61d881c12432b5829b2232409dfc2b7345ab05d702508)

Select alternatives according to the provider validators above.

<a id="canonical-1683afa6daa376d1ccfcb1272cf957850737fcbc8651b5ee51e02d01cce97864"></a>

## Direct properties — http_health_check / 883c5cca92e3 / 3

- [disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-a8f7e1a415fe7ca014643c8a5bd207ed20a622b8ff1e5733d1645d8fa9ad7d96): complete subsection reference.

<a id="canonical-eb44e716159ba5c17f23d6119020f4baceea2002638531cd576e3a0051827d84"></a>

<a id="canonical-d8f8a8a5159e69003676615a22bb9bea845f0aa4c81abcb7fc9fb25e8759bf46"></a>

## health_check_port property — http_health_check / 883c5cca92e3 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="canonical-da74b22a8d9b8845443a2a3cf3ac6c44fd7675c11dedbaaba0842b7deb45413b"></a>

<a id="canonical-97e811f6f24413e05aa93e208a1e9d7d594d8ca6f3756d1ff37a01aeb20d20f9"></a>

## health_check_secondary_port property — http_health_check / 883c5cca92e3 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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

- [inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3cf9492c903a310f1bba851a12ffac5e64ed092f82913e3801f4dda5444a36dd): complete subsection reference.

<a id="canonical-147240bcdc7de32d010b6ef038d8b8694c23129ca0eafa24f0aa26029fe9f244"></a>

<a id="canonical-6ee292d47a050df5362438a6840cbad064a2431853bec5f43fe0d9d59e0124fa"></a>

## receive property — http_health_check / 883c5cca92e3 / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

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

<a id="canonical-d45a3a8c1814167ee06d93ea75a852460f98953ae50e30555fddf3aade84428f"></a>

<a id="canonical-5f834810e369008550306e683f127b8b6c5b08a937c7056765755791f3e857f2"></a>

## send property — http_health_check / 883c5cca92e3 / 7

Type: `"string"`. Computed.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

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

<a id="canonical-9bd866e427f96814552f0976651ee285b2188a7e8940ae5ac01b3395151f83b1"></a>

<a id="canonical-7618af17d145407bbc7e73ce5b5360d7b5c65be6ab99dd865016fcfadb76f9cb"></a>

## virtual_host property — http_health_check / 883c5cca92e3 / 8

Type: `"string"`. Computed.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

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

<a id="canonical-559c55092c2e4690c9d600e618d6c2f2c34ab52bf1ae301042ff0fb0e09d25c0"></a>

## Next pages — http_health_check / 883c5cca92e3 / 9

- [http_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-a8f7e1a415fe7ca014643c8a5bd207ed20a622b8ff1e5733d1645d8fa9ad7d96)
- [http_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3cf9492c903a310f1bba851a12ffac5e64ed092f82913e3801f4dda5444a36dd)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-a8f7e1a415fe7ca014643c8a5bd207ed20a622b8ff1e5733d1645d8fa9ad7d96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d2160611d2b07b23c025fb957a79f676fa7541e25f3988a90233704b12f1adc"></a>

## http_health_check.disable_virtual_host — http_health_check.disable_virtual_host / e2ac06f5b2ca / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886)
- http_health_check.disable_virtual_host

<a id="canonical-1e01b296d8eb98e015531884fa38d9084e5a7a899e760ceb23203f4a7c2a0345"></a>

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

<a id="canonical-3f4045b5e5cbb5c25ef2ffb72b73a571c4762822a3b47f44afff9df94aae293e"></a>

## Direct properties — http_health_check.disable_virtual_host / e2ac06f5b2ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49fe145189cee2246f8859b1fb2560b3ed966a318a0fa794c0e9b2ee135bcbf2"></a>

## Next pages — http_health_check.disable_virtual_host / e2ac06f5b2ca / 4

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-3cf9492c903a310f1bba851a12ffac5e64ed092f82913e3801f4dda5444a36dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d9796ee634e3fe57e2d7611beddc0588b62a7789dc61d4552de13e47b8dccd0"></a>

## http_health_check.inherit_load_balancer_fqdn — http_health_check.inherit_load_balancer_fqdn / d4bf1140b6dd / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886)
- http_health_check.inherit_load_balancer_fqdn

<a id="canonical-322d5223f7b0f6378da52e63729ee7f7c32caf060ba40d6f07c1bdf249c5efeb"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-c65283d3936527d8e7862e76dd73f52710af327fe6989b6f7f0a222b18e7849d"></a>

## Direct properties — http_health_check.inherit_load_balancer_fqdn / d4bf1140b6dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-777609d7d3b4896260a2b5939bf95b0d86666421779114a28400241120cc089e"></a>

## Next pages — http_health_check.inherit_load_balancer_fqdn / d4bf1140b6dd / 4

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d88b291d878f603463763718184291ee4bda195deb73e2db663987f1fa71f886)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db7e1ecc5acc27d3366460903b37da4ef8a07bcd09e233e9dd39af150939eec6"></a>

## https_health_check — https_health_check / 862f91646b4b / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- https_health_check

<a id="canonical-620872d5c1c0ca39d9a598ff368b4290d4abacd8331dca49743b2449bcccd3c6"></a>

Type: `"single"`. Computed.

Configuration parameter for https health check.

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

<a id="canonical-b01ec6373c2c30cf86a7b597cef6e9bcf4cc3b9739016494010b11cad7883534"></a>

## Direct properties — https_health_check / 862f91646b4b / 3

- [disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-ce2ec2f6d7c2d9132cff1e7af9de081fc1bc991de9c9285ff7f46d32eeb8b4bc): complete subsection reference.

<a id="canonical-c6620e204e68cdc9e6349ed48a69333e297ca097123013a448c26e320ce178e4"></a>

<a id="canonical-dec2ed1b4c4515e38751deac3a9becf68ee907ddd00949cdfade4e8a5fcea2dc"></a>

## health_check_port property — https_health_check / 862f91646b4b / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="canonical-1ca6a82997c2808b75260192f0a20b6b3f340afaf0b43f3e5e7bc11e93104f89"></a>

<a id="canonical-03f24f55a2322dfc5c72818f306db0d2c4cdbde9825d48adc3b9f904347d9f27"></a>

## health_check_secondary_port property — https_health_check / 862f91646b4b / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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

- [inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-9d114dcafc5ae2da09a5c53f08688b4ded50bfb0eaa21007d9aa3bc97ae17ca9): complete subsection reference.

<a id="canonical-bf8581d5b1029d87942aada10ebecd4077ef74d66a6101f6b3f0a8a4a1406c6e"></a>

<a id="canonical-8595924825a420ad41cc07b1e6fa47ef9df623df18fb865285fcf1526e6ed44e"></a>

## receive property — https_health_check / 862f91646b4b / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

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

<a id="canonical-c10436e758ab723abea01e04ddd9bce65f94d2dbbdf54b1e339f62bee0f30bcb"></a>

<a id="canonical-820b33516217f03b02291e6b8b7ac01293b6ea1916737d0a322ee4f582642147"></a>

## send property — https_health_check / 862f91646b4b / 7

Type: `"string"`. Computed.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

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

<a id="canonical-c0354958b85c0dbab258d9e09c3c1cc4ec9c0a669d78ae9e0ceaa35048cd7639"></a>

<a id="canonical-bffc420233b6efc7f3479a0761ed4001ea9c723814b2f365f25c669ee2c8881a"></a>

## virtual_host property — https_health_check / 862f91646b4b / 8

Type: `"string"`. Computed.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

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

<a id="canonical-2ce3b45b6a0af99a4d302400947380a1f16d4c958c1da72f67400874d8ba473c"></a>

## Next pages — https_health_check / 862f91646b4b / 9

- [https_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-ce2ec2f6d7c2d9132cff1e7af9de081fc1bc991de9c9285ff7f46d32eeb8b4bc)
- [https_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-9d114dcafc5ae2da09a5c53f08688b4ded50bfb0eaa21007d9aa3bc97ae17ca9)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-ce2ec2f6d7c2d9132cff1e7af9de081fc1bc991de9c9285ff7f46d32eeb8b4bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fcc18253540e673c14ec123071a0636e02011fa1aa27642a09f93d25e23bc2d"></a>

## https_health_check.disable_virtual_host — https_health_check.disable_virtual_host / 94f7ef4eba7c / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496)
- https_health_check.disable_virtual_host

<a id="canonical-43a42e5f393732bbc8dc9738d86b0b8688669f37178a73acbe86ab31936721ee"></a>

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

<a id="canonical-5888ebd371f5e2cde4d10d28d8039d9305d6b3da829fde79fe3b2d7c7433d0f6"></a>

## Direct properties — https_health_check.disable_virtual_host / 94f7ef4eba7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29f96d8a3bfc4f53cc9aa7504cd8ea86cc91963dbdf62896314a0a9059002c2a"></a>

## Next pages — https_health_check.disable_virtual_host / 94f7ef4eba7c / 4

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-9d114dcafc5ae2da09a5c53f08688b4ded50bfb0eaa21007d9aa3bc97ae17ca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ba6e1211a8e1822bd3a988a654aef7954e4b0dd314ea2c083c09710dca03fe0"></a>

## https_health_check.inherit_load_balancer_fqdn — https_health_check.inherit_load_balancer_fqdn / 509fec4023cc / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496)
- https_health_check.inherit_load_balancer_fqdn

<a id="canonical-b4e5d93b59ae866a99d04abd958272c5bac7eaff461b25ea3ee5d129daf8e7bf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e7d7ad1da2b1e21ce4395d31f5dfe0970d26f7ad91c2d48a1dd32fa983bc6cf7"></a>

## Direct properties — https_health_check.inherit_load_balancer_fqdn / 509fec4023cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0f2756686782cb244eb62cee978b62481b03d50304236bbd5e3dfd01216062f"></a>

## Next pages — https_health_check.inherit_load_balancer_fqdn / 509fec4023cc / 4

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-dddfc77ece5333ead2a04bbbef88eb179e68700718283190f091a451f94cd496)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-8d079515e1ebee028ec1536f7db33c9eb2c85a624380f5a5c8ff91dc69e65a4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ff0ccd7daf2b9977961a7d4b9de15ada7a5e32edc162507b25561e38889c67d"></a>

## icmp_health_check — icmp_health_check / 9a70ca982599 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- icmp_health_check

<a id="canonical-79b79c14af5195b5f99d7ff0e15ca4917d29f81d8133a26a3c41ae995b4dee13"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7d21e573a02031cf4aafcb584ed096a85d8e43e8e04c329617bceb39aea93b9a"></a>

## Direct properties — icmp_health_check / 9a70ca982599 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13c127786563af22c142af540fd9a2d3d734f0331a7d1fe167ed40633613e05b"></a>

## Next pages — icmp_health_check / 9a70ca982599 / 4

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-f6ebc80c6ee49cd5d14fd5e59781d78a5ad266fd7e36e4fc06c4bf49d3e2d37b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b23f985a13bcbed68cc73e989a0d4c0575472d392fc73fb45a85550f3908094b"></a>

## tcp_health_check — tcp_health_check / 22ced8fda02c / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- tcp_health_check

<a id="canonical-2827883a91f2fcb6edf33a763428e3c6a48f8a21bc98fc4a0f59b019e18038a3"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp health check.

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

<a id="canonical-a7de74e42b0b4c1792badd7d1afcaa7849d3ca443f63d1ceefa9ccdb39088499"></a>

## Direct properties — tcp_health_check / 22ced8fda02c / 3

<a id="canonical-58d0664c31c0693d6c444e91a8fdbedfbd2f054ddc676839a251e135bce1809c"></a>

<a id="canonical-d90cd7920c8e87e01de0909fbf2969c527dd2e368b1ea1ec5fdd57afa2f91dd9"></a>

## health_check_port property — tcp_health_check / 22ced8fda02c / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="canonical-02487b41b6b4eb85e2677e4d4bb72413b5fbc9ca5adce9f4fab9465a90cdce53"></a>

<a id="canonical-fd1a057054f14795797180ed5a0014b55d36acd7b05f3980a5e714b23717be11"></a>

## health_check_secondary_port property — tcp_health_check / 22ced8fda02c / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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

<a id="canonical-3828f30000447cf39ddf5002132576a0915ab9c47bda463b7255eaf74ea0cd73"></a>

<a id="canonical-ffae582f5307ebf0a09655f14acf1e01c38616f8f8334d2e4a0c853707d65c33"></a>

## receive property — tcp_health_check / 22ced8fda02c / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the monitor's request. Mark node up upon
receipt of a successful regular expression match. Uses re2 regular expression syntax.

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

<a id="canonical-a204f03504e4fe0f02a310d78db091a521cec72f76482bfffcc638dcc2abe844"></a>

<a id="canonical-43e0dbfb49690671d175a11f9ba90a0f52b7a445a140239542d367c4af841f57"></a>

## send property — tcp_health_check / 22ced8fda02c / 7

Type: `"string"`. Computed.

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS).

Upstream description:

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS)

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

<a id="canonical-a1e09ab523671653c28e95a25e080c1f10fabbd340710dbef808311792bfea93"></a>

## Next pages — tcp_health_check / 22ced8fda02c / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-b37c57860ec3c66a1ce578d5df4f2e81fd285eecb0ca6dd0f33b6fb66bc7763d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7fe1b38a4563830ef0aa15ebb334cdb644ce4052667ffbc1b87531c9cbf4587"></a>

## tcp_hex_health_check — tcp_hex_health_check / 5b74f2979271 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- tcp_hex_health_check

<a id="canonical-199046536be55acd4bc1f1b45f6aba1ff44a4a821642cd1c2b8f4affe9017a19"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp hex health check.

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

<a id="canonical-37fc07bf0eb6d1a553f14c0c21159485c4dd695d3e0a3aa02a2a982944fa81a4"></a>

## Direct properties — tcp_hex_health_check / 5b74f2979271 / 3

<a id="canonical-211e79b65b2f4adb476edbaeb411a446dcbc705be6919e698de3f03a2536f7cc"></a>

<a id="canonical-45b74cb2af7fcdb5b257debd233771736762f54c79bcbc5fc71d660717b7174d"></a>

## health_check_port property — tcp_hex_health_check / 5b74f2979271 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="canonical-ee9aa321e24e5ba4626fa82b8a4b05c01f0c7d82cae1790f8d3642eb71cf20c9"></a>

<a id="canonical-beb5d55c286ef79b3d6d8d17b1a39460a2924088f08c72807c4704a3b73c727a"></a>

## health_check_secondary_port property — tcp_hex_health_check / 5b74f2979271 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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

<a id="canonical-302955bbf60896b7829c1249f4d0faa49103099b9692c80791eae463f5190111"></a>

<a id="canonical-b47366a4e891a8e4f7d76a8b43118ba307890fc4f8db19b84f7a74ca8f445f00"></a>

## receive property — tcp_hex_health_check / 5b74f2979271 / 6

Type: `"string"`. Computed.

Hex encoded raw bytes expected in the response.

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

<a id="canonical-c2bb2a57fa1ed448ff75f82c27f18dabfa7a3faa06c3cd1c5675649b76847b56"></a>

<a id="canonical-8efa12d95c85210124efc406b53da5fb54ebaa5a019f34dc5cf1ea2548ce63ef"></a>

## send property — tcp_hex_health_check / 5b74f2979271 / 7

Type: `"string"`. Computed.

Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.

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

<a id="canonical-e3a0cd07ef8607712f06cc37af080c1673dd82ee1e79d23f668d9ef68c5594bc"></a>

## Next pages — tcp_hex_health_check / 5b74f2979271 / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-4350c7f6bfdc26288aaa3248f59edf2895c799e100f59cd0269003ce43fdf936"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86163c8f56bd32d1ff91a3917e0404c090db6fa382d5f1edeb7eb76192cca796"></a>

## udp_health_check — udp_health_check / 05987d41a078 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- udp_health_check

<a id="canonical-3ecc01bc17e2f94797a61d881c12432b5829b2232409dfc2b7345ab05d702508"></a>

Type: `"single"`. Computed.

Configuration parameter for udp health check.

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

<a id="canonical-3470df7def6b3cee04b7d4de4f19cdf994ede6d62730eac42428670c91f59a7b"></a>

## Direct properties — udp_health_check / 05987d41a078 / 3

<a id="canonical-419afb4242ddc7102096a7247d02071f983297f493620d32f0d301becc3f366d"></a>

<a id="canonical-5fb91b152462595d493bce4dd2a3356708a6259bfd69f708cc8c550780778204"></a>

## health_check_port property — udp_health_check / 05987d41a078 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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

<a id="canonical-ff504e7c2d3ecb60252c331fd6dd5df7c00106915483204b35bca289a9362c77"></a>

<a id="canonical-c561b626df78e6063012ed9c7fda1ccf6899a10f218757c1aea595e3430e3915"></a>

## health_check_secondary_port property — udp_health_check / 05987d41a078 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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

<a id="canonical-2cb5dfe10f2b52ee6a95f3350d6d6425d17c49b55438322bff3758d0a2e9071f"></a>

<a id="canonical-ec8439d4358525953ba434f7365f53045e15963082ce1f3b620a7b806cd837dc"></a>

## receive property — udp_health_check / 05987d41a078 / 6

Type: `"string"`. Computed.

UDP response to be matched. It can be a regex.

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

<a id="canonical-c26c971da87f5e1a0e27ba7ea704e68d963fbe07cd27060d61c8e64049669ffd"></a>

<a id="canonical-0be5b98056e1a16eb6291cc93562f7a8a0cbdaf1c0dc91e47af874d77efc0fce"></a>

## send property — udp_health_check / 05987d41a078 / 7

Type: `"string"`. Computed.

Send String. UDP payload.

Upstream description:

UDP payload.

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

<a id="canonical-49e4c3261f449d8432f4cb42c6af76580d9e55f047c0b74cfe39887367670bb3"></a>

## Next pages — udp_health_check / 05987d41a078 / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
