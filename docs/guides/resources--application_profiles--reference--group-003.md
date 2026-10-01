---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-0c2fd454d67df34a2cd865a1b6d12d74eeaf60b428b11a5d5894b2ab8e16606b"></a>

## virtual_server.http3.udp_server_profile — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-039be1a3ec0416bbb07feba4ff81a7ea46e0ddfcf8a1c6b56c629e6e8776f55e)
- virtual_server.http3.udp_server_profile

<a id="canonical-91ce9335ff70b3b99fa637571a1094c1c6e8c27b4961a023375409aa0e299f38"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2178f15fd964ec1656f1103dfe068458ce36def88a9f7b5b1cd8c3553cadc35"></a>

## Direct properties — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 3

<a id="canonical-00769b48749abda73d1de11475bf273bc7837e1deb61e19e586172c669bea849"></a>

<a id="canonical-2223ec3aaf5dd5e307552d0c84a7b5a1a170d4d94d9215d5825b04b639a4a623"></a>

## kind property — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 4

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

<a id="canonical-b0b3eb554762f4c6239e2f806e039b15a6d122e3812c0c65d88fccbb3c6a6f0f"></a>

<a id="canonical-55fef595edf534ed1b3f01132a37d22da15f5e582b9d1f344b95d8a95fea7619"></a>

## name property — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 5

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

<a id="canonical-0301e122e4dfa173b4ef80b43ac3444e8cd7fdd489e98c70e3d7001ac98fe5bb"></a>

<a id="canonical-8fc7d2c97ab9e7a2c43e4e28b338aad94968cfa8edfcc5d62dee1eae7c2cf6fa"></a>

## namespace property — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 6

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

<a id="canonical-aadec0fd8f3813c538fe0ff35c687e890897687780c5a61cdaa28fc486c86cb8"></a>

<a id="canonical-5e1eac2262afb905a31d3658c0dd6e230a63c76e25369ea83dd41ffda9fa666b"></a>

## tenant property — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 7

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

<a id="canonical-231443a49da8ace5e9a63eabecc3c84144c9925f146d4fc2ae015f4006b8ba60"></a>

<a id="canonical-58a2a03407733beb4362868c1879781c5ca825a71ccd514eee2ad0608e2d4b55"></a>

## uid property — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 8

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

<a id="canonical-28f7498a92c34d575b9e01f569faab181cf9e793ba15f47c67c1939b1d7b73b6"></a>

## Next pages — virtual_server.http3.udp_server_profile / bcbb238fcc96 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-039be1a3ec0416bbb07feba4ff81a7ea46e0ddfcf8a1c6b56c629e6e8776f55e)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a5a115c08449197d28663377fbd57b5210d7aa4b017a8a929548bc1e96b3a43"></a>

## virtual_server.https — virtual_server.https / 225d1fa8e8dd / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.https

<a id="canonical-c12f767912d5f44f0b3184d9bc70757f9b5d715c229a3b8a5bc245574b61efea"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

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
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-d4558eca2b91baf0af72704ba8dbb8bce7f2e23b77c0a7c31d2a5856b6b14284"></a>

## Direct properties — virtual_server.https / 225d1fa8e8dd / 3

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-1a803bf1f595f052d3bfe2792e02923d409eb9ace8ddb69cbe9a0638370414f1): complete subsection reference.

- [http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-23c558efe24e364204597f30d0c038d7e462219bd05e9602a0ac8a7325d3a71b): complete subsection reference.

- [http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-5ed53f1c6aaa140db7099d58d1920218c263e285b0b48c41368d1c6fb311d52d): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-003.md#canonical-e610d45b5d44a7a9a53e91460f45be69087631ffff9de53b2943178579e82317): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-003.md#canonical-5a33ad0d2dd5930f7283d9663fa2167464af94492a8bc5e8298925436fb3a8ab): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-ee91676db8298fa4fa6c3c499551b74c91c6dfc9aef6de63f9a88136b895d7e3): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-e44208a7cdf7713b1ecefd288eb29f665249fc759b244eaf6d69848dc7e3d4af): complete subsection reference.

- [stream_profile](resources--application_profiles--reference--group-003.md#canonical-e62b2c71d32a14baf9623ee98bf5f31a1413183bf9ce510f22a00de2ba352ffc): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-fd69099393ce2d56bb23c4a6c895f66ebf0ad87cd55e62f27c94d974961fe153): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0a902fc1d6abb507c83cf2af55634dfd549e7de37d7390e1d01e39925b6e38de): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-6aa7c874d19dcd00bf97d424922d0acf0c19697b8bbe76e9684f6f03470007dd): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-254fec42f5a19078a2e47f5217af6dd612a04c0b27d3e91fe5e117b12752dff5): complete subsection reference.

<a id="canonical-974f59b70236f0ab1f6494bddbf320c0ddeed9ed06f7db44199183b5021bed5b"></a>

## Next pages — virtual_server.https / 225d1fa8e8dd / 4

- [virtual_server.https.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-1a803bf1f595f052d3bfe2792e02923d409eb9ace8ddb69cbe9a0638370414f1)
- [virtual_server.https.http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-23c558efe24e364204597f30d0c038d7e462219bd05e9602a0ac8a7325d3a71b)
- [virtual_server.https.http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-5ed53f1c6aaa140db7099d58d1920218c263e285b0b48c41368d1c6fb311d52d)
- [virtual_server.https.http_client_profile](resources--application_profiles--reference--group-003.md#canonical-e610d45b5d44a7a9a53e91460f45be69087631ffff9de53b2943178579e82317)
- [virtual_server.https.http_server_profile](resources--application_profiles--reference--group-003.md#canonical-5a33ad0d2dd5930f7283d9663fa2167464af94492a8bc5e8298925436fb3a8ab)
- [virtual_server.https.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-ee91676db8298fa4fa6c3c499551b74c91c6dfc9aef6de63f9a88136b895d7e3)
- [virtual_server.https.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-e44208a7cdf7713b1ecefd288eb29f665249fc759b244eaf6d69848dc7e3d4af)
- [virtual_server.https.stream_profile](resources--application_profiles--reference--group-003.md#canonical-e62b2c71d32a14baf9623ee98bf5f31a1413183bf9ce510f22a00de2ba352ffc)
- [virtual_server.https.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-fd69099393ce2d56bb23c4a6c895f66ebf0ad87cd55e62f27c94d974961fe153)
- [virtual_server.https.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0a902fc1d6abb507c83cf2af55634dfd549e7de37d7390e1d01e39925b6e38de)
- [virtual_server.https.websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-6aa7c874d19dcd00bf97d424922d0acf0c19697b8bbe76e9684f6f03470007dd)
- [virtual_server.https.websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-254fec42f5a19078a2e47f5217af6dd612a04c0b27d3e91fe5e117b12752dff5)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1a803bf1f595f052d3bfe2792e02923d409eb9ace8ddb69cbe9a0638370414f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08216f3e5727afc421398c0426e00ba47e9279f14f2a452c0caa88883f0f41b9"></a>

## virtual_server.https.client_ssl_profile — virtual_server.https.client_ssl_profile / f5297d6db713 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.client_ssl_profile

<a id="canonical-346592f022e0465aac97e141e23a6da552e8014118356b040391f2bbfce06a2d"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-afaa8d1394131ff0a12096d8c9537568a312a3475c50c8462c9e281d33ff2bac"></a>

## Direct properties — virtual_server.https.client_ssl_profile / f5297d6db713 / 3

<a id="canonical-ceb5453b4fe4933a1bc7305cebf1c205e961e21e6ba4236ba413880ef71d6c44"></a>

<a id="canonical-372edc0c112cfe2c0df64e09776f5f0ed628afad0a19db450eaf9bf6bf2689b6"></a>

## kind property — virtual_server.https.client_ssl_profile / f5297d6db713 / 4

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

<a id="canonical-030f102a31424332ded3239128e3cc41c9a049bfddf95d4228e82d2bcea0b825"></a>

<a id="canonical-7762f0eed5e8a550c5f41c76a265cb9ab0198574c23c507196d57b752dccdd47"></a>

## name property — virtual_server.https.client_ssl_profile / f5297d6db713 / 5

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

<a id="canonical-75ee85239f6c9daf8b38db97e307d1441e73c4af51c11a3d4d8f17afe112b424"></a>

<a id="canonical-dd08d937a2f3b80e84a5ee1940ef1b6112c08c451879e5583dfd637493a74548"></a>

## namespace property — virtual_server.https.client_ssl_profile / f5297d6db713 / 6

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

<a id="canonical-8a8b8c9b5feabe366cf80e9bca586afaa6e7cd3850e80bbba957cf940950b18c"></a>

<a id="canonical-64c55f1eca790cd5a8ce79d6f0b3be6bc40f80f6816f2176b9a03dd8defb46d2"></a>

## tenant property — virtual_server.https.client_ssl_profile / f5297d6db713 / 7

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

<a id="canonical-abd120550a81fe102dc071773511f302e53a877b779f6b52fde282e6a4cae6ed"></a>

<a id="canonical-0fb911ae76c68498daf111f38af28b1494f5ccc6178daaddb3e611622eb66859"></a>

## uid property — virtual_server.https.client_ssl_profile / f5297d6db713 / 8

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

<a id="canonical-3e81d30b7385dba09ae8f6ace44f42658e9995247b48cc88322fffc93867a4d2"></a>

## Next pages — virtual_server.https.client_ssl_profile / f5297d6db713 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-23c558efe24e364204597f30d0c038d7e462219bd05e9602a0ac8a7325d3a71b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3a7a43663e206bd0ef5b006df2c83f950e317be5a1551403670a3d87a026e85"></a>

## virtual_server.https.http2_client_profile — virtual_server.https.http2_client_profile / a80abb1d8181 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.http2_client_profile

<a id="canonical-436f0da56064a67c4727b5d52442b880b80dd2b9aba4470690cb8e1a9fd56958"></a>

Type: `"object"`. list nested block, Optional.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-770c4e21ace14a7da0d76104f5249f4110166de07e560e80bd45cbb39a21a848"></a>

## Direct properties — virtual_server.https.http2_client_profile / a80abb1d8181 / 3

<a id="canonical-08939e498c5bea659814aaec35fc4cccbc6d16b0686d63ae5b28375f7d5e0f09"></a>

<a id="canonical-7d76622e31b544c63b75457c6a68f7b526c17983f7fc9b738985a368c79c1352"></a>

## kind property — virtual_server.https.http2_client_profile / a80abb1d8181 / 4

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

<a id="canonical-da093a1f1d2d42b1396b2c9703d6fd3ffe38f3c74e117dfc3a1b9b3430ab483d"></a>

<a id="canonical-fb024b75605e92563a550148ac2236fdfa7b1da89d59da3e8bb37e7114892404"></a>

## name property — virtual_server.https.http2_client_profile / a80abb1d8181 / 5

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

<a id="canonical-65f7ded224eb0bf5cf08649759b5581892a1855d69148cc69275c97c1be5e060"></a>

<a id="canonical-3d6289b5c5eb974ebe16e3c61a9c65dd777e175f79c97ff0e1a2228116151c38"></a>

## namespace property — virtual_server.https.http2_client_profile / a80abb1d8181 / 6

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

<a id="canonical-7ac3469a0378b3924fc576fb897fd1065d1e625abf2f051963cb192472f20e0f"></a>

<a id="canonical-b72aa4050572c0d82853f563fc0cdfa58287d5abfa8eb24748b153922afbd449"></a>

## tenant property — virtual_server.https.http2_client_profile / a80abb1d8181 / 7

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

<a id="canonical-53737348eb9297b48dbe4927a0b988fbc1ba6e9bdb944ef8720ccdb64123b6a1"></a>

<a id="canonical-c75a399624afde3fbb56d7db3568713a94fef7233adcb73a40fd47a9d2c1ec94"></a>

## uid property — virtual_server.https.http2_client_profile / a80abb1d8181 / 8

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

<a id="canonical-e18952bc41a54b192da40c6fb85f93f161496b69a41b26c0e02381e01b6f1b22"></a>

## Next pages — virtual_server.https.http2_client_profile / a80abb1d8181 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5ed53f1c6aaa140db7099d58d1920218c263e285b0b48c41368d1c6fb311d52d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd6c78f0caef63c4021b759d7d10eed2b713b21f4a4d48f32a7c71dc0287698"></a>

## virtual_server.https.http2_server_profile — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.http2_server_profile

<a id="canonical-71d4c299e04fca28e4bd687804cd9b6b37a97408d672141afa6c4ac842ced0a3"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-23c2983d5070f53706b0ee4e593b92ea4cf64e323f088b3e8eddc31aa837b17b"></a>

## Direct properties — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 3

<a id="canonical-5796dbce2add113cbd542b582e834486ec8395d9783d7cf5656c8cdc1b3dd01b"></a>

<a id="canonical-f296f8fa6a40c4abd7e39ad406aa5dc965d560b5a67f11989fa36c1a4a48658f"></a>

## kind property — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 4

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

<a id="canonical-cbbcf21d79d5c549bd19398b07c5f80da5906f0afac2024609b446aec9b421ef"></a>

<a id="canonical-d1d19e9b40aaddb6fd77a6e351c18d38e4fdbb93332010cc3a31b2aa0e85caec"></a>

## name property — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 5

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

<a id="canonical-b1e3e6c571dd8d7a009472f52b238df74234ab799fdca2e093492e1954d5b394"></a>

<a id="canonical-92239c44afe98724f146b9611f1028089dc9f18dfedf5c8a4ef5bcf811e78a28"></a>

## namespace property — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 6

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

<a id="canonical-38a478a384a7d503e37a0c95c8cf1bdce2b7b7018b1966ba08e148d648bb1503"></a>

<a id="canonical-f7f4e4b06a676f1165c726f0878814348845d10303939231017f36d8fbc38e17"></a>

## tenant property — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 7

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

<a id="canonical-13c530ac0dd187503d291ef1e5838038144392d8bf71783974a57473f0b987af"></a>

<a id="canonical-ae36fcbac9c734ea9288b7b9111493a8fd233b2d9eb11903fe0d2972b8d04923"></a>

## uid property — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 8

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

<a id="canonical-514262cc9690d433e1db161ab3cfa250707ff6c114e2191381559879c050ff57"></a>

## Next pages — virtual_server.https.http2_server_profile / 35d11bddd4d6 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-e610d45b5d44a7a9a53e91460f45be69087631ffff9de53b2943178579e82317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ee73d0a101babc9566877377a8ecd41785f6f510c54066ecaa9f90839ba9cd8"></a>

## virtual_server.https.http_client_profile — virtual_server.https.http_client_profile / d156670c0f4b / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.http_client_profile

<a id="canonical-0eef7685571edbc66585e9e5580c3b8957f4d2fe839d7660339cc5d329d91163"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-10129ac2e2f2e2c5fbca1e1cdb89f5cf816ac0be8c46ac1f501b73deb0b27f6f"></a>

## Direct properties — virtual_server.https.http_client_profile / d156670c0f4b / 3

<a id="canonical-8b5e4800205aa9d52d8dcc091d4c6c736f88ee59aae8661c7003b273939e9969"></a>

<a id="canonical-d0b7b6f9ca58b31e589f75a6f8f264c2c5ad3484d69836dae2fed915a9b50f55"></a>

## kind property — virtual_server.https.http_client_profile / d156670c0f4b / 4

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

<a id="canonical-33ee89c880a974ec2be065a4ec2436c4dd21417dc1a0c1b95d9713047db5fbaf"></a>

<a id="canonical-e2cd6c03432fb85fcda8e894228094d2a9650a479461865d4d6ae84802dd0e3e"></a>

## name property — virtual_server.https.http_client_profile / d156670c0f4b / 5

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

<a id="canonical-0abfd85db6ee4d3f6fc6ce86ede7a2b6d8403132c283b81328e465e931787f11"></a>

<a id="canonical-23f139e8755f86e51e44e6f041d85b81d501bc535a1f48e616b681e1e7d30d49"></a>

## namespace property — virtual_server.https.http_client_profile / d156670c0f4b / 6

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

<a id="canonical-89ddb871e37e5990782f46d225cac1303b188c289c7d3f4837e43aaf532838b8"></a>

<a id="canonical-069701aaca98678cb0deeb50412df671523cd921b0cc9d2cfa61dec48aed5318"></a>

## tenant property — virtual_server.https.http_client_profile / d156670c0f4b / 7

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

<a id="canonical-af3d4d3db3887e57d6c294ca2bb9fa411cef2d8387403fe453e4fe2e3ab041b9"></a>

<a id="canonical-9abc6375840e3386caa1c202de5ada32a9a1c1fd9bc9d37e0a80bced37005c0c"></a>

## uid property — virtual_server.https.http_client_profile / d156670c0f4b / 8

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

<a id="canonical-d8a22fa877d455fee88693faadfd0ababdafdb746add799630ff170c6e78c40e"></a>

## Next pages — virtual_server.https.http_client_profile / d156670c0f4b / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5a33ad0d2dd5930f7283d9663fa2167464af94492a8bc5e8298925436fb3a8ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e89eb55b087ff4e8df3a19b53bd28d1e679ec29e6bfeb6a5c8124e02649d92fa"></a>

## virtual_server.https.http_server_profile — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.http_server_profile

<a id="canonical-0a5c06e974d71222cfc2cda8d64f7623e86e342dafb61f1c879932ecb3483d86"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ab7043a23bd793de77fbda216c341ac62fa2d222f41dd003dbe16c2fcf51e51"></a>

## Direct properties — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 3

<a id="canonical-e3991582b273e401b6e8534bdb7cdd6e0dfdd36f52714e2f024b11eea274ac22"></a>

<a id="canonical-454cf712b339f12c84ecc2020a2ebf1b691c06eeb0fdccbbd90363d495620ad3"></a>

## kind property — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 4

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

<a id="canonical-23e0d69506efbe7c2e2637267f1ef732e985ec1296f5a57878e9a60d73873468"></a>

<a id="canonical-4c3c54b66c174f67a2adfd799b0ff96caac69f210bb2ad612cf89623ef96c329"></a>

## name property — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 5

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

<a id="canonical-31d4473e87ad4f43b2addf04407d1d9fcf192e23777de39db0258ea7ff8fb775"></a>

<a id="canonical-6983a0568bc1c0683c3d602938ea91c0a287bd593c3891a7224814264f77aa2c"></a>

## namespace property — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 6

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

<a id="canonical-4fe9cffff49b64326a653a12c5e7306f8499771324ccff92819f86c11845138e"></a>

<a id="canonical-11b3b2b07ef650336413c1549716d21cb1179f3d1b1e4b1ce3fa69f7e4b6361b"></a>

## tenant property — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 7

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

<a id="canonical-3e4adfa944cdcf115b07ef8fb2d5e9e5af4be2b4fea6cff44fef3bca0c56bb74"></a>

<a id="canonical-fa1f7b67acb7cf71126ed87eb6b78a405de01fe30d70ff358228c332cf339e8a"></a>

## uid property — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 8

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

<a id="canonical-90a11302fd349f42dee180bc8929ae471c9871deff54d7f6260eda71cc3c4c35"></a>

## Next pages — virtual_server.https.http_server_profile / 4ebce49f1ef8 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-ee91676db8298fa4fa6c3c499551b74c91c6dfc9aef6de63f9a88136b895d7e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd2f3a94a794a19ab0db7a64bab3fb5d1018a3b7a0c2694f84def682cc40dfd4"></a>

## virtual_server.https.ocsp_profile — virtual_server.https.ocsp_profile / f26eec94715a / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.ocsp_profile

<a id="canonical-5e18bda1a9fbe6751673f7300838c82ff71f6ce67c2a7316b75f38651366c0a3"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-921e08600c7d241023caa357c5031662b54b1f66a983c1279ef84ffe79b5335e"></a>

## Direct properties — virtual_server.https.ocsp_profile / f26eec94715a / 3

<a id="canonical-3dcf425ed83a0ac8ff11fe133fbee859bad0700d94ac8e5ca1f8312bb962d356"></a>

<a id="canonical-caf3eece46f0bec77e5f80d573bd83a9ef631f39f81e0a35b85aae5ac1c8c19c"></a>

## kind property — virtual_server.https.ocsp_profile / f26eec94715a / 4

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

<a id="canonical-ecafa25f3abf950b937457ae2c64302ecffb7fff4bfaef4ab59a0a07c3ae5a25"></a>

<a id="canonical-dff5676c4604d953f7ae5ddac1332e5641c47b32008a801615c28bc70638bb16"></a>

## name property — virtual_server.https.ocsp_profile / f26eec94715a / 5

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

<a id="canonical-200150ae1dd4248747080e86f10f1c8ca45014bd2e2965e57822a9123b42ee76"></a>

<a id="canonical-7564578780eda8982409ea49aa4e7d4c275098a354a1077559c8884a64a9c420"></a>

## namespace property — virtual_server.https.ocsp_profile / f26eec94715a / 6

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

<a id="canonical-8aefd7bc18f5394b2fe911443561a33aab3fa094827832828faf9da885d6a023"></a>

<a id="canonical-30609f8c71e76d3a712ca9e033a1cdfae88ae93d32f13df87c788242faf281f8"></a>

## tenant property — virtual_server.https.ocsp_profile / f26eec94715a / 7

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

<a id="canonical-2116f561c3af4ed03122b1602bfd66fa58e71f472a179060b4f0c2d339c5a1f2"></a>

<a id="canonical-e82f7256716508dbcd628d4276bde7328800c52987d1d17788457806710cd0a2"></a>

## uid property — virtual_server.https.ocsp_profile / f26eec94715a / 8

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

<a id="canonical-462cf8a83f3ae77df258361c0fe2bbb19cd3d4fcf05cdc57c44cfb677e3bbdf4"></a>

## Next pages — virtual_server.https.ocsp_profile / f26eec94715a / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-e44208a7cdf7713b1ecefd288eb29f665249fc759b244eaf6d69848dc7e3d4af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fc616c5664f4b51cdd5398bf41c4637bde3d530ed2289bfe3d6129553cc8158"></a>

## virtual_server.https.server_ssl_profile — virtual_server.https.server_ssl_profile / 614b540f053f / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.server_ssl_profile

<a id="canonical-7274ce87cb31a97f73e61974cba60f4bf10415aefae6a39cdd2881ce85c3e836"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-6be1a865e97a4f7e1ebc5ab6499d0cdf42e185ac22349fb2c0ab82498773d0d0"></a>

## Direct properties — virtual_server.https.server_ssl_profile / 614b540f053f / 3

<a id="canonical-19f1db1a9e8e8d4ea39a99a26df65bf96e28d7ca84e5ad78f10ba9bc974dc68c"></a>

<a id="canonical-d649ac0f13ab59c21395ed61b448f336bc656acc65b08bd4af3f0a35055c2ce4"></a>

## kind property — virtual_server.https.server_ssl_profile / 614b540f053f / 4

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

<a id="canonical-c530a3ac6e0778ca486c987d8a16814383a104f86acaa43df2472d8208c3601c"></a>

<a id="canonical-80d73a974eab02b5b30b03ba4f866b07ac70e821315935f4bdf39e70552fed6d"></a>

## name property — virtual_server.https.server_ssl_profile / 614b540f053f / 5

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

<a id="canonical-05a5e04656d7cf1e3a3722338295809e06ec3cebbb0c006365930ac60cfdfd58"></a>

<a id="canonical-f7cc98f965d0ca030e5ffc4aeadc1753413d44954da4a7f19a0ec1750efaa15f"></a>

## namespace property — virtual_server.https.server_ssl_profile / 614b540f053f / 6

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

<a id="canonical-34f1a30d11ab5e024c51ef7bb3ddca98d5445cfd19df1be1781507f247395c9e"></a>

<a id="canonical-411919406665ce1d1b8e4b1358c31401325abdfb31ce8881e4b9f62fce428dac"></a>

## tenant property — virtual_server.https.server_ssl_profile / 614b540f053f / 7

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

<a id="canonical-324613bde8ef8d8a046130fd3603dae3d1b4c7b4d7eb4dfaf4a9293f68ab32ae"></a>

<a id="canonical-2ec3c1d97b1e7d783aff9e3346feb1ef1ade420cc57b3f9058b0bbbd7c3ae722"></a>

## uid property — virtual_server.https.server_ssl_profile / 614b540f053f / 8

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

<a id="canonical-7f969999d1606a4d63f6deeee8b885e4690c2b7d4fb7114c38032d07af483ddc"></a>

## Next pages — virtual_server.https.server_ssl_profile / 614b540f053f / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-e62b2c71d32a14baf9623ee98bf5f31a1413183bf9ce510f22a00de2ba352ffc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3adbaff114a8b965b44648a3a972b859a047544633a9a3cac86732db5d6979c8"></a>

## virtual_server.https.stream_profile — virtual_server.https.stream_profile / dd12cb9e09e6 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.stream_profile

<a id="canonical-7b7e9b1165a094bd3556f71afd636c6afb817abc42116819b88593d54deea18a"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
stream_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-172ce8da4841cec8e4d8c0a744ba2c450bdb0cb4459bf6a5d6aa8ac4b1867093"></a>

## Direct properties — virtual_server.https.stream_profile / dd12cb9e09e6 / 3

<a id="canonical-5098cebd6f266741e34cfa9436d44ec0c33df4670e1ec93e854e8aa9f001bd7c"></a>

<a id="canonical-1aec96d0d3f79188887cbf7686a55ad05125080a478e43d79eb379469945ca67"></a>

## kind property — virtual_server.https.stream_profile / dd12cb9e09e6 / 4

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

<a id="canonical-e5f5393ba9fbfa2de14687c2c3f3c867fef9e21403b6c35985e51633ce835599"></a>

<a id="canonical-85ade32a7d748ea6063cf6f3770e9cbbc7927a4277ebe5ca993b197c31e4edbf"></a>

## name property — virtual_server.https.stream_profile / dd12cb9e09e6 / 5

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

<a id="canonical-d1771fd482dcebb3f940a55e6a13e4b112aab7293cf0638424f90ae04c73d25c"></a>

<a id="canonical-fa8d2922a90df2e8fe8c76a507fdba36df2f98f8d0527718257445d52a2e5af3"></a>

## namespace property — virtual_server.https.stream_profile / dd12cb9e09e6 / 6

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

<a id="canonical-b37ade79bbb8cf1a0ffcc042eddb412dd80a166f373cc0ebf3eacb3c15f8c72e"></a>

<a id="canonical-fa1a86fb34a6a5a87ece96668bbf5f206b3cfb87eb3c30d64812a042ff9f124c"></a>

## tenant property — virtual_server.https.stream_profile / dd12cb9e09e6 / 7

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

<a id="canonical-b094ac5875ec5b65be23438d8b8ff8c191cb79fc6a65b4f2b4105b83b16faf48"></a>

<a id="canonical-0c2312ed20cd7ed23a09875fcefff09fc4ad499ad0d042323e712d986083b65c"></a>

## uid property — virtual_server.https.stream_profile / dd12cb9e09e6 / 8

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

<a id="canonical-1e3ebd1bb144a6610ebf259c81f9011a2c47031a71ba63a79e62c668dd167e22"></a>

## Next pages — virtual_server.https.stream_profile / dd12cb9e09e6 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-fd69099393ce2d56bb23c4a6c895f66ebf0ad87cd55e62f27c94d974961fe153"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fcb285665ae78c1b251ca633d2f58ffb2318e0048d5725579cd9bc6a4a64481"></a>

## virtual_server.https.tcp_client_profile — virtual_server.https.tcp_client_profile / b880253bb818 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.tcp_client_profile

<a id="canonical-f3231e170758e2607aa8b03298743fc33977baedd5b4d2e7f4e91354f57c0007"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6f4989cea8990521dc95b6cbd6f23a99e64cc9d165dc850d463247eb7fd20ec"></a>

## Direct properties — virtual_server.https.tcp_client_profile / b880253bb818 / 3

<a id="canonical-cec26f11880ce06ab9beb3010da2f6264f9715d41cc68cfa022ccf4bee1ab35e"></a>

<a id="canonical-622d45e2cc072dff4e5018129ebbe1e0382ac70f511a6c6f81748524a5fa2ef6"></a>

## kind property — virtual_server.https.tcp_client_profile / b880253bb818 / 4

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

<a id="canonical-217c450d943ca852ce71e5fe397ec6c26564a697ac3eb2c98ed72d6cc4f79f13"></a>

<a id="canonical-045f408e8a0aca1f323d43e44e307d975b25aac0507978657bd6211eab6617ba"></a>

## name property — virtual_server.https.tcp_client_profile / b880253bb818 / 5

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

<a id="canonical-83a5c58c9be0974b41f3d532b51d1f348f2289de52812dac11658ee7b067211e"></a>

<a id="canonical-8bd1293e9909da921e3a815ef9c2e4bbc996ad37feeaab300a0c4bfdd3206adb"></a>

## namespace property — virtual_server.https.tcp_client_profile / b880253bb818 / 6

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

<a id="canonical-ca81222692db9467fe6a7630cc84f9f473b4d07bdb95680a13cf9ff9fd940670"></a>

<a id="canonical-5e6ee1f54feb04bb83cd6fb090b3002227b646f1a890aac3f75d52644f050395"></a>

## tenant property — virtual_server.https.tcp_client_profile / b880253bb818 / 7

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

<a id="canonical-5ef0790081802244a3299e17392289902ce36a60c888fcdfe869f54b0dc6f0cf"></a>

<a id="canonical-af49fc94ff1405253334f51126b032c07d78fb9abc33de59bfd9766fa0323ee5"></a>

## uid property — virtual_server.https.tcp_client_profile / b880253bb818 / 8

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

<a id="canonical-6e4cb47167a453cda2f3d990ecf393b6a491b1da86deca84a588dab3326d11fc"></a>

## Next pages — virtual_server.https.tcp_client_profile / b880253bb818 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-0a902fc1d6abb507c83cf2af55634dfd549e7de37d7390e1d01e39925b6e38de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7311043af8d0ce56928ea859384eaf3b9963ae048695c979315278c83db9a59c"></a>

## virtual_server.https.tcp_server_profile — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.tcp_server_profile

<a id="canonical-a4f8d051ce5f5ab415b4312e3faa5f4adca3cc353b6ff2e7dab0d73f3ac7c222"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-d6c51d0d314afdc9a8f74a0df86afceb82df38fb98110f267c599b3630fc7f7b"></a>

## Direct properties — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 3

<a id="canonical-9b4788bf6cdd5603f99f0bc33f248a771ac4aa4985cb553cb450cdbd62f25f2d"></a>

<a id="canonical-9621a4826d2c68812e4cf55b6e11652f71f96c9982fcdcfa313e84ba66a26d33"></a>

## kind property — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 4

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

<a id="canonical-659fe76bf938d6b572b196c9547b7b85c9e39f0f2c031169cb3a47e80f79ac12"></a>

<a id="canonical-b57474f5cecceb2ab9c02bb66eafa1656b3c230417e812ed5da3f0679511a9ea"></a>

## name property — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 5

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

<a id="canonical-2b11f28303d27b9c834ffced3c02517e7d50f82d9793ae4485046a78a0f3497a"></a>

<a id="canonical-c31db320065cd3e24fdc180fa4c6c30dbba222b713a9c85f8e8239472d6527ce"></a>

## namespace property — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 6

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

<a id="canonical-fec493f5811dd05474d3dedf3fdff55348b10bb4e72d96536f079f1c9a2d7f8f"></a>

<a id="canonical-a58b372989a5a1a6c6cc04d22730f777cabfe7079c6f35d124f0fbd4088ed7b0"></a>

## tenant property — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 7

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

<a id="canonical-b3d34e7e850d6c00ab6cfeac09bd7b225556766340217145708fb26f113c92ca"></a>

<a id="canonical-650ea04f82e0b1421b6e9db8e452cac4ba8417908ead912d6344f60f107fdec7"></a>

## uid property — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 8

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

<a id="canonical-e55212f17b6eb2f99ce13c1f1a57e3e2756e5412c4d670f252871963f5f0f281"></a>

## Next pages — virtual_server.https.tcp_server_profile / 4ad9f45db653 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-6aa7c874d19dcd00bf97d424922d0acf0c19697b8bbe76e9684f6f03470007dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9041159bc197a1cd86ebb68212cf1fab7e2b6dcaf63afda370b2fa8d53742852"></a>

## virtual_server.https.websocket_client_profile — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.websocket_client_profile

<a id="canonical-3979bc2e7735f3e2d8cc8e2793517d62d1c16c29b2b540fbd5fa60bc2714d176"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-084ceba253ab568913c176facebe05c6112ca8eb1c9303bda8dec75ed63b9df3"></a>

## Direct properties — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 3

<a id="canonical-6c4536a7a8e99e55d40511a2270fe2c901e616722d1dad28feb56102a371ad72"></a>

<a id="canonical-7a6a8c27d1d53b53a782eaa9d36df553c7811ba2132dc0c3e5431078210a9ea5"></a>

## kind property — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 4

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

<a id="canonical-cf4a4344a849f6881dd488b04ed0009d807995da7d50b6339ea2f2a436f54224"></a>

<a id="canonical-b6c121aaeab4cbe222c164bd8b425fda083128174ab04862d870cc6a505d27be"></a>

## name property — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 5

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

<a id="canonical-6bc98672809a8db34704b2757caec71c7059ac0348a2eb1cd4fd45c124d39421"></a>

<a id="canonical-a3d81852567931ed4439831a3e55ceb9b30514fe68c43358315d81ac782860be"></a>

## namespace property — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 6

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

<a id="canonical-0dbccd77ef194e199f2298cb39a3043cc01176576d7ebe03dde33ea4dfc27c90"></a>

<a id="canonical-baa1420b9c3e132f9ffcf6647beea49d1793dcc5aca3cf0b9523f4260aed444f"></a>

## tenant property — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 7

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

<a id="canonical-22427b8a9609c50e6fe35670f0d868feba5378682e493a2598dca4cf3204cc15"></a>

<a id="canonical-1271fb7b22c0da0a250b5d70d0c0ac93423255e2dcf0573e3db1fe143c39a64c"></a>

## uid property — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 8

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

<a id="canonical-3b6b76b27d936194ef6aacb54b02aa936167af3723e960c56e207ea5a581c49b"></a>

## Next pages — virtual_server.https.websocket_client_profile / d2c9bc12b753 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-254fec42f5a19078a2e47f5217af6dd612a04c0b27d3e91fe5e117b12752dff5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d3b3ee27d46d6e611a8ebf5de0d97a735e68e7813e179e7ee85da5cf3e80023"></a>

## virtual_server.https.websocket_server_profile — virtual_server.https.websocket_server_profile / 50091f85ea6f / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- virtual_server.https.websocket_server_profile

<a id="canonical-654f3287c560d8b6f61345373d9da715135d8e42102e4b8ae75c4734f39313ef"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-5271fc9f93a8acafcbe628f1ff43ea08a11f0fe2c913ddfcb9e7972a162c01fe"></a>

## Direct properties — virtual_server.https.websocket_server_profile / 50091f85ea6f / 3

<a id="canonical-8279ec2c21d1e98f7a58b976cb7034130eb0b1f2e273a7e10ecd6717e8940e2e"></a>

<a id="canonical-cadb2039b581d6d99155eab7acf7cea2104167008cdacd0d027cb883cf8139fa"></a>

## kind property — virtual_server.https.websocket_server_profile / 50091f85ea6f / 4

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

<a id="canonical-d346a6c528997cbd2339e786e76e516885bcfb909d2969805b84d146b3b83934"></a>

<a id="canonical-38c6bbf02fc27ecb5078c3c93856d90aa473b6b75ec36f29d99adff6aafb40ae"></a>

## name property — virtual_server.https.websocket_server_profile / 50091f85ea6f / 5

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

<a id="canonical-270afc9d0097ea62f1ac0ac1feac6f00f6093e6fcf689c31213b48f53e32289f"></a>

<a id="canonical-6f0a44e6f19b466af7740bc25641ca7c7c2cfded67d9514ba6f878e791ff2a2a"></a>

## namespace property — virtual_server.https.websocket_server_profile / 50091f85ea6f / 6

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

<a id="canonical-4b7405f33c3baf66e4965fad838bc643f10145eca86904a8bc86801dc5646706"></a>

<a id="canonical-6289d0012352377a176215e9197ff7436778c260314ae67f4aeae2c7f810866b"></a>

## tenant property — virtual_server.https.websocket_server_profile / 50091f85ea6f / 7

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

<a id="canonical-ddaf3e7f06a75149cd005ce515dc8efa4c214a0edfa19d0fada3f20ab6134382"></a>

<a id="canonical-548bee3969d588268aee9df8469b2f9966943530ad3539e5ffb74b43439c3c0c"></a>

## uid property — virtual_server.https.websocket_server_profile / 50091f85ea6f / 8

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

<a id="canonical-840d839ade39dac2b0e321dc56712d5e80383f7ac0a7ce07f1f5fea1e4de6750"></a>

## Next pages — virtual_server.https.websocket_server_profile / 50091f85ea6f / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee7e8201c107f9a83d9583ea24075422949e8c8539ec826283208befcf8ba4f1"></a>

## virtual_server.immediate_action_on_service_down — virtual_server.immediate_action_on_service_down / 44848abe01b3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.immediate_action_on_service_down

<a id="canonical-ec9cea6a07fdf90f6a787dd6ba4045377cddd5942c66d1842e3c1e9f3489e62d"></a>

Type: `"object"`. single nested block, Optional.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.

Upstream description:

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_none"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_reset"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_none",
    "immediate_action_on_service_down_reset")}
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
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

Terraform syntax:

```terraform
immediate_action_on_service_down {
  # Configure direct properties listed below.
}
```

<a id="canonical-62f898a0d96b6034a82a8d9ef3ec19bd9a52417ef297e68536cafbad58674980"></a>

## Direct properties — virtual_server.immediate_action_on_service_down / 44848abe01b3 / 3

- [immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-fd2136b89cae3e7d8a6a9993a1cd22f7855ffcc9445328a7fc85b99bb9d3a99a): complete subsection reference.

- [immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-65e1ddeb3031f4b358ce240f323d86622582a79ea96109613e9be052b8c8042a): complete subsection reference.

- [immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-48ec22b28d8067291c6c6a713308bd1a45cf772b27a0e12e28706ca2c8f191e1): complete subsection reference.

<a id="canonical-da6bb448830605d532de45ed159ce4df8d199c3ef2b94ee6687a06edff92109b"></a>

## Next pages — virtual_server.immediate_action_on_service_down / 44848abe01b3 / 4

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-fd2136b89cae3e7d8a6a9993a1cd22f7855ffcc9445328a7fc85b99bb9d3a99a)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-65e1ddeb3031f4b358ce240f323d86622582a79ea96109613e9be052b8c8042a)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-48ec22b28d8067291c6c6a713308bd1a45cf772b27a0e12e28706ca2c8f191e1)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-fd2136b89cae3e7d8a6a9993a1cd22f7855ffcc9445328a7fc85b99bb9d3a99a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2719d1b3cd22bcc63e13a05e069f3ae12139090c15084fa95914cd3a54d3676e"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / c93a70f70224 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-691f984d154b0ab1f511d9cc01826c7e202854dc6c33ac416448f2233412ebc6"></a>

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
immediate_action_on_service_down_drop = {}
```

<a id="canonical-dc8e720afe4a46fab7adb02fdba08ae470f7fabef809085c8cb44b1542323d08"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / c93a70f70224 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a251ec5eb2e58de5a93c5d096af30272ecd4c0baccc096be6bf46bee3268c1f"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / c93a70f70224 / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-65e1ddeb3031f4b358ce240f323d86622582a79ea96109613e9be052b8c8042a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b79e2bbf4e1fe3bb7b757af33ec84deb5e16768e7de275793ffeec03cf7a41"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 2b73cc500c9c / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-8374076fd0af922f614197f2e08f815516788b0188f7785601ae47d68d6bace3"></a>

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
immediate_action_on_service_down_none = {}
```

<a id="canonical-d859fb683772f3e3de5428e49d336033f8d833de0ef30f79f1eb9826d1a8dcb7"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 2b73cc500c9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60f9a70209031a9444c37ce31444f2f069714428271f4d85e11813bf07570066"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 2b73cc500c9c / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-48ec22b28d8067291c6c6a713308bd1a45cf772b27a0e12e28706ca2c8f191e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3193bdda7647869b09309b9d56770aa11cbb6ec4f98af3deda314a74c2dd175"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 4f87d861e805 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-f294896154e9c1112b7c80c75beba8c9f6ede37ce021ecb600943fd6cfb325a3"></a>

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
immediate_action_on_service_down_reset = {}
```

<a id="canonical-354b12d3caf0fd8f44d262f1578ad972636942aca64918c7500a1df741246529"></a>

## Direct properties — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 4f87d861e805 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73c7317a63814cc2e195d31f818e734175476f629c152d7b7b0a298d123946e7"></a>

## Next pages — virtual_server.immediate_action_on_service_down.immediate_action_on_service_down / 4f87d861e805 / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5fc75b0574c39ea49487a9c012d59d535f61f433ca326badc03a12230bc11b50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-452e15f79bc4d89d6be6991fc3910ef055a541519ff7327f75d4ac378d7d7e21"></a>

## virtual_server.last_hop_pool — virtual_server.last_hop_pool / ab701cd57e59 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.last_hop_pool

<a id="canonical-dc4369260886abbb12afc2c793664126983cbba3b6bfd4436061a7e5ee3bc0d4"></a>

Type: `"object"`. list nested block, Optional.

Directs reply traffic to the last hop router using the specified pool.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
last_hop_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-c03ce8760317895ef5383e8e4b699a1e24bddbee9a6435b42165211642437681"></a>

## Direct properties — virtual_server.last_hop_pool / ab701cd57e59 / 3

<a id="canonical-1fa3e5a3988fd2f3134e105008d323777a679a94abb39220bbd12f5c94e8445f"></a>

<a id="canonical-4433655af838f914467883a2f851eada48a1271b3301e861eb529daefe895f17"></a>

## kind property — virtual_server.last_hop_pool / ab701cd57e59 / 4

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

<a id="canonical-1d6c072281ff1cf4f47d7da6a5f7160dc9221fe64b0f4139d2ef411773cb936c"></a>

<a id="canonical-0cddeec1e8d4818d86ea422582a49804f614e61bbc3466975bb199afeb2621c5"></a>

## name property — virtual_server.last_hop_pool / ab701cd57e59 / 5

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

<a id="canonical-d4a04fff1e52e838bfc7b1a4460b4e0a64c88cdd3616ef52e38715eca713724c"></a>

<a id="canonical-0951ffb1e9d61797b1c35e0f9b2aabe8b937d6fd3228c41d19abbb28f78492dd"></a>

## namespace property — virtual_server.last_hop_pool / ab701cd57e59 / 6

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

<a id="canonical-48661e5539ab3a35cda180df4bbec286897faba5c271c7e4371d1b4d7462bacd"></a>

<a id="canonical-21a69c38178df7ba2bccb25779661e2d5c5e761ef0be104322ca88ff72150ef8"></a>

## tenant property — virtual_server.last_hop_pool / ab701cd57e59 / 7

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

<a id="canonical-83307141818367bcc6adacd693a9b05fe6f4937d6a1304bc1c4a0bdc10afafa7"></a>

<a id="canonical-1a54a548651b6431808e85d7b5abdcd501da4d7a359213c03a9177850140569d"></a>

## uid property — virtual_server.last_hop_pool / ab701cd57e59 / 8

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

<a id="canonical-f0798b796285c1a10bd24ef749b3fc3047dafebe6adf78afd3ec282d5f2e98d4"></a>

## Next pages — virtual_server.last_hop_pool / ab701cd57e59 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68fd244dc0826566a542918a64c9f7c9e56cef969352ecb232133d40e23f7e2f"></a>

## virtual_server.nat64 — virtual_server.nat64 / 30202827723f / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.nat64

<a id="canonical-4346f65ce486e2d6f472232980a36fdd53c585a2ae5144e5c908fc0eecaaae0c"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("nat64_disable",
    "nat64_enable")}
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
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

Terraform syntax:

```terraform
nat64 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d994daeee311293187ac12115159db312bfce2788d8f1f59b58cbc2941271cd"></a>

## Direct properties — virtual_server.nat64 / 30202827723f / 3

- [nat64_disable](resources--application_profiles--reference--group-003.md#canonical-5177ee374ccc9cd9e14c4f16207390e3efec9abb09ce155c00dced3f167e7c77): complete subsection reference.

- [nat64_enable](resources--application_profiles--reference--group-003.md#canonical-1e42d028ef5287b822b2e0693100b333ede1b2a1bb1276c81e0913cdd8cdb91c): complete subsection reference.

<a id="canonical-c4c37c09fd35163291f8e8fcb276a9cf91921c084874bbc6933eba6d77afcb3e"></a>

## Next pages — virtual_server.nat64 / 30202827723f / 4

- [virtual_server.nat64.nat64_disable](resources--application_profiles--reference--group-003.md#canonical-5177ee374ccc9cd9e14c4f16207390e3efec9abb09ce155c00dced3f167e7c77)
- [virtual_server.nat64.nat64_enable](resources--application_profiles--reference--group-003.md#canonical-1e42d028ef5287b822b2e0693100b333ede1b2a1bb1276c81e0913cdd8cdb91c)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5177ee374ccc9cd9e14c4f16207390e3efec9abb09ce155c00dced3f167e7c77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5aca1e0004a7e824fe5767f457b080ccaeebbe4898cae03a010b080c7541a81"></a>

## virtual_server.nat64.nat64_disable — virtual_server.nat64.nat64_disable / d52eec717cfa / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42)
- virtual_server.nat64.nat64_disable

<a id="canonical-0e82a5dbebd31e94d2ef4c793be3161ee7c6abbf73ee30316c0b9c5720ad0793"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 disable.

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
nat64_disable = {}
```

<a id="canonical-661be2f7fa5cd03da5a7b13f76ef75628f672d02177ee363d2e2c753488fe9f3"></a>

## Direct properties — virtual_server.nat64.nat64_disable / d52eec717cfa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a63abdd17deed905abdb0d6c7d4897a0a06b9d1ca5d1ece517d4bc1b35aac486"></a>

## Next pages — virtual_server.nat64.nat64_disable / d52eec717cfa / 4

- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1e42d028ef5287b822b2e0693100b333ede1b2a1bb1276c81e0913cdd8cdb91c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5bc72e00239a705cce7b2c389d658fc7bd42e825000bad5bc16971a48b03387"></a>

## virtual_server.nat64.nat64_enable — virtual_server.nat64.nat64_enable / 794061f4727e / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42)
- virtual_server.nat64.nat64_enable

<a id="canonical-e7d81cb31bc12f4185be25e6cb7852352a6cc6d54beb514343c78efe6d81aefc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 enable.

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
nat64_enable = {}
```

<a id="canonical-9c42db89134719e082e9bbc3be283d0f4e0c518d711810d7b4bc2648ad505cef"></a>

## Direct properties — virtual_server.nat64.nat64_enable / 794061f4727e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec21223f6648857721aa6ff854f92548bee551582fe6e355da5fdafe1f8af85c"></a>

## Next pages — virtual_server.nat64.nat64_enable / 794061f4727e / 4

- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1450d28ba7ea514796b0ba852c55641715eb68bc99debe45cdc65ba4e29225fc"></a>

## virtual_server.port_translation — virtual_server.port_translation / 3655a6532768 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.port_translation

<a id="canonical-f095dcca27f2dc833b94bad6297a8a3659fd99e49ff6e6be04c31bb61a640999"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
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
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

<a id="canonical-e575cf4a7fb013e75d4cc98c3013019806bdd840930e2bb1b2f8670d009020ef"></a>

## Direct properties — virtual_server.port_translation / 3655a6532768 / 3

- [port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-ba68e06238ffb5d37854a42d5aa09be31ef5ef452adc3bf635d2fa5a30bc2c04): complete subsection reference.

- [port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-bfd7f25f120c9ee8b8f0572232b85c051f253c334c7728b1992ad78a3f18f64b): complete subsection reference.

<a id="canonical-22dac7ee04587ab07f2e76a11ff590853d928bfe69f9371683292f98ba3ac1b5"></a>

## Next pages — virtual_server.port_translation / 3655a6532768 / 4

- [virtual_server.port_translation.port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-ba68e06238ffb5d37854a42d5aa09be31ef5ef452adc3bf635d2fa5a30bc2c04)
- [virtual_server.port_translation.port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-bfd7f25f120c9ee8b8f0572232b85c051f253c334c7728b1992ad78a3f18f64b)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-ba68e06238ffb5d37854a42d5aa09be31ef5ef452adc3bf635d2fa5a30bc2c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a73490a68cd303ed3a9209e974a4b8d0baab64f021b32e393477d3e9962fa860"></a>

## virtual_server.port_translation.port_translation_disable — virtual_server.port_translation.port_translation_disable / c0819d9a7dbb / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-3d87c5a8fa021cabf2ee51f4452afc39128f1d04315128e02a7698771c256c2c"></a>

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
port_translation_disable = {}
```

<a id="canonical-0e6981d0658aef759ac51d77bf4929a73abd8e2959b681397a516972179f51b9"></a>

## Direct properties — virtual_server.port_translation.port_translation_disable / c0819d9a7dbb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67d550f23a017b5429453adc3d8e44aa0baf72280a207658badb5a63eb2fcc39"></a>

## Next pages — virtual_server.port_translation.port_translation_disable / c0819d9a7dbb / 4

- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-bfd7f25f120c9ee8b8f0572232b85c051f253c334c7728b1992ad78a3f18f64b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e6cfd4680b7f083653f494b93f9b631341131abe03460a92a2714ce3a0db476"></a>

## virtual_server.port_translation.port_translation_enable — virtual_server.port_translation.port_translation_enable / b06e61991e8a / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-9493a6c550b1ee2511d94f2104de2634b2c317adf3804f14486107eb3c16fb49"></a>

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
port_translation_enable = {}
```

<a id="canonical-b48e78d3e8f4d2f4f8b0e497466d4a38619012afa6c46251f064dbb0cf87ee6a"></a>

## Direct properties — virtual_server.port_translation.port_translation_enable / b06e61991e8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-423430208e325f740cc86f3477c2ebaa39f4d762390beb4549eca83c89378bc0"></a>

## Next pages — virtual_server.port_translation.port_translation_enable / b06e61991e8a / 4

- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-e48b0df1348ee94f794c3241260b6d7e56f733fdfb00cdf93002db17e8b5c4c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-686fdd876f9133c2209047664c49adaff700c2fa738f4271cf3c4eb12e563e69"></a>

## virtual_server.request_logging_profile — virtual_server.request_logging_profile / b1ccceff8991 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.request_logging_profile

<a id="canonical-c94395b16fc46d6036f0b56081c48cf5afc9c78d972d11d51f26f9fbeedbc9cb"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for request logging profile.

Upstream description:

Configuration parameter for request logging profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_logging_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7afe422b98b14dbd5653e4187da641446f18fe584fe28996a32600181b4ed82d"></a>

## Direct properties — virtual_server.request_logging_profile / b1ccceff8991 / 3

<a id="canonical-afb8c2f7df8bb51dddeae68a486f983fb995effc3c91a880aebd9236597f6df9"></a>

<a id="canonical-da529a0060046e6cc6748df116c38ae39ecea5b9f6737a96262a0b4fd61360c5"></a>

## kind property — virtual_server.request_logging_profile / b1ccceff8991 / 4

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

<a id="canonical-bc5da30fffd28d7e2b2519f37b4f9a59e197667f0f391774be2ac7cbfaade0e8"></a>

<a id="canonical-f43391ee39ddd38ff9a2d14ecb43988aefffc9cdf69c8749280dd4d16830de46"></a>

## name property — virtual_server.request_logging_profile / b1ccceff8991 / 5

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

<a id="canonical-2f6134327585451158f6e06400669c19981cd0d127f6f9ba04eac6bfdb6fe1a4"></a>

<a id="canonical-6c3ebc24efd278f3030e73b6aec3fc2e6c691935c184933f677a5e028ab6ab6a"></a>

## namespace property — virtual_server.request_logging_profile / b1ccceff8991 / 6

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

<a id="canonical-234a5addd3487780dc6e4cf590d07e8b83e708b211079bb22e8f9400b41d4f01"></a>

<a id="canonical-29769e3d045ebcceb3f5e9c1764835ee1a6a7abae2c562d422f10fa7464cb6f9"></a>

## tenant property — virtual_server.request_logging_profile / b1ccceff8991 / 7

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

<a id="canonical-4d23b12874549cdd6573161418cf2f36318acf941a955208bb4977fc289cf9ad"></a>

<a id="canonical-579fbf47ef2205ed1190fe95b10f58d3bdd0a7443eb969a9edb64b6e92366bc4"></a>

## uid property — virtual_server.request_logging_profile / b1ccceff8991 / 8

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

<a id="canonical-e3237260fd8cdb54190530f5dd1cfaefbcff65b17e92da4ae17464b550af1d58"></a>

## Next pages — virtual_server.request_logging_profile / b1ccceff8991 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef75efb620e41175c0d2a1bdae4991a82b3488a1ed001ba22c4fef309c9b0040"></a>

## virtual_server.source_port — virtual_server.source_port / 3443ef02f9af / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.source_port

<a id="canonical-a613eadd95ee70e9da9f32f488fddebb138504424812e277158e224aab781e27"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve"),
  validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve_strict"),
  validators.ConflictingObjectAttributes("source_port_preserve",
    "source_port_preserve_strict")}
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
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-a78dd1e6c13d8722b31f4e5a5cc3e735e156cce3dbef625614b78135092cb820"></a>

## Direct properties — virtual_server.source_port / 3443ef02f9af / 3

- [source_port_change](resources--application_profiles--reference--group-003.md#canonical-01b9e3892319b95f38575fd3ffcfa457dcd1bc58f4c2e35e174fa446c9c4fd0a): complete subsection reference.

- [source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-1351080d981ac798e9d1b0ea1214dbc4fbe40562e438bc14641a181a867ffe73): complete subsection reference.

- [source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-940bfab690fbccea000da52e14785d6c5b2c3cef9864e8e2e56e6290ccbf2d0d): complete subsection reference.

<a id="canonical-2e5242cabc2e9e02316f231ae28ebf42434588ff6d74896e2e46a7364e779e41"></a>

## Next pages — virtual_server.source_port / 3443ef02f9af / 4

- [virtual_server.source_port.source_port_change](resources--application_profiles--reference--group-003.md#canonical-01b9e3892319b95f38575fd3ffcfa457dcd1bc58f4c2e35e174fa446c9c4fd0a)
- [virtual_server.source_port.source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-1351080d981ac798e9d1b0ea1214dbc4fbe40562e438bc14641a181a867ffe73)
- [virtual_server.source_port.source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-940bfab690fbccea000da52e14785d6c5b2c3cef9864e8e2e56e6290ccbf2d0d)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-01b9e3892319b95f38575fd3ffcfa457dcd1bc58f4c2e35e174fa446c9c4fd0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1db1e9a7c2f14b4daa3d855b87073ed122373cf1b00f96b64dc52e9e4ad9eec"></a>

## virtual_server.source_port.source_port_change — virtual_server.source_port.source_port_change / 4e2f2e48c296 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- virtual_server.source_port.source_port_change

<a id="canonical-75f3b24ea41ffff27dc4372303fe8840ddf46180873b274cb3897683fb5e15b8"></a>

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
source_port_change = {}
```

<a id="canonical-c24f51cca4336b5f90a010c97a17a86798996be05a1d93fb84c53d8543d4cb35"></a>

## Direct properties — virtual_server.source_port.source_port_change / 4e2f2e48c296 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1d35262d391111e436409a971ff734366e9604a2c18bc5b8a621cdcdff74805"></a>

## Next pages — virtual_server.source_port.source_port_change / 4e2f2e48c296 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1351080d981ac798e9d1b0ea1214dbc4fbe40562e438bc14641a181a867ffe73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d4097bda118d319a397152996b727d6d8f96e1bd3dfdd8d4a122b83634f6ac0"></a>

## virtual_server.source_port.source_port_preserve — virtual_server.source_port.source_port_preserve / c7c73a7042d3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- virtual_server.source_port.source_port_preserve

<a id="canonical-4adc65163883520b628ec14c48429bb07a04b018dbdeb1a633acca5b5265e1fa"></a>

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
source_port_preserve = {}
```

<a id="canonical-4844b786f743255a1fc596ac86bacb7b8dda9da28ec1189e918e8d686ac6ffd3"></a>

## Direct properties — virtual_server.source_port.source_port_preserve / c7c73a7042d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed24a88f96ebd82daa3ac1037780d96b991d9a20bf22dda1d7a93bf5a997df40"></a>

## Next pages — virtual_server.source_port.source_port_preserve / c7c73a7042d3 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-940bfab690fbccea000da52e14785d6c5b2c3cef9864e8e2e56e6290ccbf2d0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0da9f9e3190399b29142b7fbb2cb098741990f60ea14f7d590f15f5c8564409"></a>

## virtual_server.source_port.source_port_preserve_strict — virtual_server.source_port.source_port_preserve_strict / 72994dc74b48 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-a654341548db2fd370475774355a37928e746b675b49ad0db436c1311c08d1b0"></a>

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
source_port_preserve_strict = {}
```

<a id="canonical-f828234cdd91a1520879237e3635320a368972d1912108cdea70215b1a3620f0"></a>

## Direct properties — virtual_server.source_port.source_port_preserve_strict / 72994dc74b48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9758c3ccfad55d5e82cae3d3b9afb041a3773caeef965056cefe4b0c5e230ae0"></a>

## Next pages — virtual_server.source_port.source_port_preserve_strict / 72994dc74b48 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-acf84624ce84de1878e7e0416c9389ebf7710cc4cca4e71825ff53f6c74fa497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec64ca820f3435886a18aa2b8b11b8bf58d23a3927cfec5a35016822de9a6a9b"></a>

## virtual_server.statistics_profile — virtual_server.statistics_profile / 70b2e6146fd2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.statistics_profile

<a id="canonical-8f35735e0b193c1f57d5ef898184c4aed550735b7acf117ec911e8dad9b30fc0"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for statistics profile.

Upstream description:

Configuration parameter for statistics profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
statistics_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-dbda11efc63c86538d212039352cd68bd503ea051b6489f8e74bc24db4fa2768"></a>

## Direct properties — virtual_server.statistics_profile / 70b2e6146fd2 / 3

<a id="canonical-432884df9d55f0fd3c3d8833a132033ab3e420a5228d387ab4f582f08a8a48cd"></a>

<a id="canonical-12300c86d1fed53ea9cd086836382adbbf7e875b6e69ed4ca4716faf95704c67"></a>

## kind property — virtual_server.statistics_profile / 70b2e6146fd2 / 4

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

<a id="canonical-70f6775064b6be64084d9ff6d172a12ede8a5d51c41757da0d47436772347dfb"></a>

<a id="canonical-bba582b0fd672285805b4efeba91dfd9a92a5c39514c6f6f1baf01d852caf055"></a>

## name property — virtual_server.statistics_profile / 70b2e6146fd2 / 5

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

<a id="canonical-ebef21e9a6a78d5710d8a0913ce3e46110be1e894cd96d2e05294dc9f02e582b"></a>

<a id="canonical-6d21d90ff9bdfc3b80acae8dbc0097105cb823edabaf7f16411ce4f88b38454d"></a>

## namespace property — virtual_server.statistics_profile / 70b2e6146fd2 / 6

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

<a id="canonical-a97d56edd44712d6ed0c70069ab11533e61eca0caeb8f1a86bbd620d0233fd59"></a>

<a id="canonical-dd386001bd50fdf48f4b726240a3dbb9c6a04640c6d9c649e8ac4b79f3e2412d"></a>

## tenant property — virtual_server.statistics_profile / 70b2e6146fd2 / 7

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

<a id="canonical-1ee1a0e52db5afbd43137d279bc5ac6d3afa9d0a236e56e86a3727e960f5a2d4"></a>

<a id="canonical-8779172c73835c3a80e4e7f60c47dc5f8a453c303d79c20b231747c3f5d29a10"></a>

## uid property — virtual_server.statistics_profile / 70b2e6146fd2 / 8

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

<a id="canonical-9e95db4b080fc33b17442e02867a8b0e36162102d64dace30b769c81ee3723a2"></a>

## Next pages — virtual_server.statistics_profile / 70b2e6146fd2 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d422d8aa33496285107b9105153f38ee9ae95208b94197546bccba6fec79a864"></a>

## virtual_server.tcp — virtual_server.tcp / d04dcd4a8203 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.tcp

<a id="canonical-7491386678c5e86f64807b6b51b0179bf23ee15ebee76a46bb7a9f661ceeebf1"></a>

Type: `"object"`. single nested block, Optional.

TCP profiles.

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
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2d466218b28d33037bd8cf8f7c6b1105d41d929f7adc6c8cb31ba5363ff9ed8"></a>

## Direct properties — virtual_server.tcp / d04dcd4a8203 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-9eaabcdc4a72337ff475a4d426a22c77f6e2e3bbfdac45fd7abdd37df7ec83ea): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-412d6cddbdee3fe46e791880afd424add4d9fe9afa40ff48f292d86eb74e9066): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-19e355c6d28e8e519d31da597674c7a38972fed44bbed952ed34f5f7e906613b): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-56c590380ecbdb2adb9faa53b8d2d078ae8a82d83ed44592f310ed6181eec241): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0ee57a2ca4a7fcef1bebc04d25136cb1155693eab493d967c0e8a6d6ad935696): complete subsection reference.

<a id="canonical-1f73e45cf746190f081804f2abb0f2eeaa81b6bfe70f28587836866613519ae0"></a>

## Next pages — virtual_server.tcp / d04dcd4a8203 / 4

- [virtual_server.tcp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-9eaabcdc4a72337ff475a4d426a22c77f6e2e3bbfdac45fd7abdd37df7ec83ea)
- [virtual_server.tcp.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-412d6cddbdee3fe46e791880afd424add4d9fe9afa40ff48f292d86eb74e9066)
- [virtual_server.tcp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-19e355c6d28e8e519d31da597674c7a38972fed44bbed952ed34f5f7e906613b)
- [virtual_server.tcp.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-56c590380ecbdb2adb9faa53b8d2d078ae8a82d83ed44592f310ed6181eec241)
- [virtual_server.tcp.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0ee57a2ca4a7fcef1bebc04d25136cb1155693eab493d967c0e8a6d6ad935696)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-9eaabcdc4a72337ff475a4d426a22c77f6e2e3bbfdac45fd7abdd37df7ec83ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f763e6c8c2d45fc1f0a34a2200bad2a7eeb967fafbea62d4899c974650264b27"></a>

## virtual_server.tcp.client_ssl_profile — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-46f674b5b5a3857a1beefa61338dc0a28019c5051d7c8de2956b94f67ed194f7"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-45ea2bf233ddb0ad48f727dff0fe8cb81402f15c59962ff7d7389753e0169072"></a>

## Direct properties — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 3

<a id="canonical-6f1b3f4271c2b70864494fc4614ce3ec4cb21b7dbcb39bea9bde0d32c481cf61"></a>

<a id="canonical-80335da06e63853a8c79935997c7760d97572fe46976b4e520cdbe42d9a197a4"></a>

## kind property — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 4

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

<a id="canonical-f627ea79a430a4716b4af5d374a19886224d3cba388efd26ba9abe8f6f33b339"></a>

<a id="canonical-500ea95aab22e457b5c871ed7fe248314cb639fce5302386978763f4953e6ccb"></a>

## name property — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 5

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

<a id="canonical-ad2582b2cd63a1a8db8ac2e261bcac46d666015bfdadd2e26cd3bb2ee8134955"></a>

<a id="canonical-01579f3992f7966d00de5ca38374a38d72735899fd767085420f34490a1046b9"></a>

## namespace property — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 6

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

<a id="canonical-d898654e474efa846303d9945f48dbb5d9aeedff2fa73b9f3bb7d52861ac9027"></a>

<a id="canonical-cfe64bd6d66f77b07b232858b5ffd96e585372f2532418567015e9f010c49c33"></a>

## tenant property — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 7

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

<a id="canonical-59bc497dcda88b59b601a1d468a7242ce0e47a36fd83c786c1d02d77fa6bb78f"></a>

<a id="canonical-59234f8f2e0311a98217cdd090265b5107bdb8d0a310167a4cf20dcf15edde41"></a>

## uid property — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 8

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

<a id="canonical-e4d6e3a5048c518f0344fb7f7fe23183c85ac8010e8a48d3b238e0705824b7fa"></a>

## Next pages — virtual_server.tcp.client_ssl_profile / fe90f9fae1f3 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-412d6cddbdee3fe46e791880afd424add4d9fe9afa40ff48f292d86eb74e9066"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5aa1fcc9cf43b59dc70899b4ed5fbee9cd2cf7993256030110e8595b11f5bad"></a>

## virtual_server.tcp.ocsp_profile — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- virtual_server.tcp.ocsp_profile

<a id="canonical-a1f12e75d7498d531e030c1baef6d5890cd6e3fc4cfe8f5faf9a98c0a8247e9c"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-d66d29b4efa957ecb68eee22eb46e25a2527a9b3adcc6d7079c00cf0dc1a0fa6"></a>

## Direct properties — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 3

<a id="canonical-6149962a4aab398e3a8f0be1b968ef657635b8b0db7490510ddb6129cedf0702"></a>

<a id="canonical-2a96ceb26e901c678dff2407ddd9b1650925dde06d02197afb3458141f8836a0"></a>

## kind property — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 4

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

<a id="canonical-e80226f19fd88d9aa705d679b06b98f57f72f867632d151f35708ca09fa23494"></a>

<a id="canonical-be7bcc96bda760a9e0f41711adef4dfef2e481bbccd182c8730ef51e877796a9"></a>

## name property — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 5

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

<a id="canonical-4cf1d22dc6ca64b21daafa29b7a68f2457fec62beb64ece4caa06a47c9be6d75"></a>

<a id="canonical-e4e8e1b27abdf248efc02a015df18f7d179edadda382833847d7e4d4d7dbb423"></a>

## namespace property — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 6

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

<a id="canonical-0eb03436592e5cadd8358ae0a6296ff798721d215e4ab5bc33c7e0b50188b0fb"></a>

<a id="canonical-9896f03aa2eec898f1008d7a9b3d59fc1d350a7d5f22e5cfa8b343d538a0c848"></a>

## tenant property — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 7

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

<a id="canonical-a92d86e4c73d88906a8edb8524bb4b23f980c7e0902bad20d0c68644582876f5"></a>

<a id="canonical-86e54f1df1fc687abeac688a50e936c52b97942be370658b823fd72512be27aa"></a>

## uid property — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 8

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

<a id="canonical-81c8cde89dbff57bc6d53a092757d1b986508b57055b59a50d9392b4636b3544"></a>

## Next pages — virtual_server.tcp.ocsp_profile / a34c88cacdeb / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-19e355c6d28e8e519d31da597674c7a38972fed44bbed952ed34f5f7e906613b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf558f7139e18cc9fe9c9bd3869f4f620ba6f71e29dda51da15ecb9e74ad173f"></a>

## virtual_server.tcp.server_ssl_profile — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-37f95f02b022cd2dad04acf92ff97a23ed1d35099dddfbd25f50738edca4daf2"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-05b2653f8f6e80b48d9ce1f7c79b42f740fd1b02f517d3ca70ee8935ea1a9557"></a>

## Direct properties — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 3

<a id="canonical-c0c673082eaf892b508cb555a06fcb1b285d76880feb26a0488d774d8fef8a5b"></a>

<a id="canonical-2f7f5a02e32c1fff3fe1d360bb12a072932a89ae0e3dca68bdd47718eb73afc8"></a>

## kind property — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 4

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

<a id="canonical-3d175c13b323714b3bd5d07eae19bfa66746c1d4a3b86953365892a49d8bbf56"></a>

<a id="canonical-285fb36deeaf79dc3470e292061a790acc995badf20de65de63706234c7117e3"></a>

## name property — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 5

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

<a id="canonical-a5de5ab96071fd4179ae36f6ba1c60cc91e0d3afc975c154f8b6350f9479c90f"></a>

<a id="canonical-d69ce7bb3dc21bec238140530c6085b29a230ea0a3464d89b4d3cff47e777fb4"></a>

## namespace property — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 6

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

<a id="canonical-38cff4667005b8e658302ecdce3f85c2ae37f7834f2963fa701ddd8b64091fa3"></a>

<a id="canonical-4e6a1765cab65b08cf1aa9465286b9d9b0886f94eea520ca4afb72158f729c39"></a>

## tenant property — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 7

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

<a id="canonical-b1a0b3fc8b7789c39db27e3570a9795d56f5a65ca532833f96cba851e1c6bfe9"></a>

<a id="canonical-f87f78064f73f8b6a8572df409903d0afd12da057b802ea6008a08e9cf710b91"></a>

## uid property — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 8

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

<a id="canonical-f60f30d4daf9b5afca702e23e945698e9665fab97921dd2215cd81e9ec1d1ce6"></a>

## Next pages — virtual_server.tcp.server_ssl_profile / d079f23d2694 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-56c590380ecbdb2adb9faa53b8d2d078ae8a82d83ed44592f310ed6181eec241"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4be52869dc662920f79cf753b2d815e1691dce2e52fb300c2f92745b1e526157"></a>

## virtual_server.tcp.tcp_client_profile — virtual_server.tcp.tcp_client_profile / 921045824806 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-bb16b3cac76eed33a02eff1c9418eeb8a503a0fb5ef1730a64511eb885792ab0"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-22bf53805ecb7b4785f3b7baf9e33cecdd2cbdeaf3d47bcd0af0c3bd01083aba"></a>

## Direct properties — virtual_server.tcp.tcp_client_profile / 921045824806 / 3

<a id="canonical-62fdf1eb39c2e7b0fafbeafaf4424b80b701bdca1f34634c153c15cb7f2df59f"></a>

<a id="canonical-dcc6fbe852e9d7a5c8ccf00e3272a129fecc93fa10f756de0e01df6a6236c6ed"></a>

## kind property — virtual_server.tcp.tcp_client_profile / 921045824806 / 4

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

<a id="canonical-b7e62883a54f8355789f45c16d8766d61f3c75e7c1dfc75fa00431e3afdbded8"></a>

<a id="canonical-d29d5b5203d06f5d710ff12c55ac80a8c1dd4132dbc3cdbf8c5897bb6710dd7c"></a>

## name property — virtual_server.tcp.tcp_client_profile / 921045824806 / 5

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

<a id="canonical-95618de0ea44a30471207585ad63b4d884e09cfcca5f7f912cce37ddadee2136"></a>

<a id="canonical-8360179c235c1ff94dbed1934094e7b049a35ad4b5734093338ea41b42d00a87"></a>

## namespace property — virtual_server.tcp.tcp_client_profile / 921045824806 / 6

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

<a id="canonical-4d5787a320183948c94ff7d3a49b84474013b2f3801154512c57d89bed6063c7"></a>

<a id="canonical-1a3850bb3af8718c32bb873865cbf1640fb467d0f5a28c1bb9203aac91a51889"></a>

## tenant property — virtual_server.tcp.tcp_client_profile / 921045824806 / 7

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

<a id="canonical-8d1fc82b1f3a8c9feb92bba254f6864927a41e4f16225623eef7482858be6e3d"></a>

<a id="canonical-9f77ce006c1544452450219b870253f6881750fbc9cdc466c1835958aa225f42"></a>

## uid property — virtual_server.tcp.tcp_client_profile / 921045824806 / 8

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

<a id="canonical-7810c8f44bfe81ab742563db17f6b456195554445f6c75eb58ae89c5a00a2755"></a>

## Next pages — virtual_server.tcp.tcp_client_profile / 921045824806 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-0ee57a2ca4a7fcef1bebc04d25136cb1155693eab493d967c0e8a6d6ad935696"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc5581f43773c896e5d0880e3841f0b17e28c1b895132a264d1f77bc17394b20"></a>

## virtual_server.tcp.tcp_server_profile — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-da4955a924121eed898094164813bd0e45351a385500691b04f1cdab6dba869f"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-efa7a78f82c1eecb342761fb8b6082427f12afa2da4ec4b3759b0e1d0d7be1c6"></a>

## Direct properties — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 3

<a id="canonical-73031e72fb2d73419d5458885717ae51d1db154590308370a9ab51020c6e79f2"></a>

<a id="canonical-9bc9bc5b1b9a879eaa50e5eb94e5ec7bfaa7dfeee4e931d9b901bdd14045cff9"></a>

## kind property — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 4

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

<a id="canonical-440a9eada095842aebeea3900c00f6e361196375619c513f2a3f0e388e610e7c"></a>

<a id="canonical-69ca3d7ea265687b53b496eda39a8f42c5da75c7d138466d8ee96d78c90be5ae"></a>

## name property — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 5

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

<a id="canonical-987a403b9295cda36b8894219d7c2cf9e36d7a46a7dde05f527a1b824938852c"></a>

<a id="canonical-52a900215114cd22999356cf8ad0512bdb5d2d0b3690e10827dfe4ca3fad3751"></a>

## namespace property — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 6

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

<a id="canonical-8c730f605f4c9b733610fbd87ee71b308252c5a6842b20361af48122cd38415f"></a>

<a id="canonical-139d8b92f75a94d5a172d153fbbde3ec4a9db7950fa52880416beab7ac41169f"></a>

## tenant property — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 7

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

<a id="canonical-eaab5ff199a86544424a598eb93b5c8bb67b1269872d36838bf2fde9e2dbdc65"></a>

<a id="canonical-84006216cd06306c2f5b2eae570a37e95a17711245a680bcf3cf3fa102867496"></a>

## uid property — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 8

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

<a id="canonical-8679e050d7fc843aae1999ea48c4555e9bc8d3da30920c0def3be039b0a4a206"></a>

## Next pages — virtual_server.tcp.tcp_server_profile / f6ced22e05a1 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f4486b0f934bb935595aee121b3cf83f0234f391a357e095a6f0463ee6f252"></a>

## virtual_server.udp — virtual_server.udp / c2e9037b1192 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.udp

<a id="canonical-22d0d88487b3c3e8e360cc38efd40f86bf000b27c5b177b9b53f087cb8ef72f3"></a>

Type: `"object"`. single nested block, Optional.

UDP profiles.

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
udp {
  # Configure direct properties listed below.
}
```

<a id="canonical-03bff815d589e96bfc882ee770d62530dddf6c49ef1b1d486f46404c09403490"></a>

## Direct properties — virtual_server.udp / c2e9037b1192 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-5a7242abada2aed97b9626ffccbe10ebfb3a4a3d61d7ddb94e56bb7cace9f5db): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-89e8ad280ef58fc47b9f7d948b455aa0bc5d1e6dbaa07b2890f5967080cb0a43): complete subsection reference.

- [udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-33e5527ec99dc25fa429d3665441749d79c053b3223a1703404360aa132266d8): complete subsection reference.

- [udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-5e3b83980cd726d4cfab529bb0c0403d3815769386b584aa442cba6d53fc9486): complete subsection reference.

<a id="canonical-89123aa7c519e24734b75fdfcbb2e34906ddb884d9ac7ad44647fc96c2a0087f"></a>

## Next pages — virtual_server.udp / c2e9037b1192 / 4

- [virtual_server.udp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-5a7242abada2aed97b9626ffccbe10ebfb3a4a3d61d7ddb94e56bb7cace9f5db)
- [virtual_server.udp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-89e8ad280ef58fc47b9f7d948b455aa0bc5d1e6dbaa07b2890f5967080cb0a43)
- [virtual_server.udp.udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-33e5527ec99dc25fa429d3665441749d79c053b3223a1703404360aa132266d8)
- [virtual_server.udp.udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-5e3b83980cd726d4cfab529bb0c0403d3815769386b584aa442cba6d53fc9486)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5a7242abada2aed97b9626ffccbe10ebfb3a4a3d61d7ddb94e56bb7cace9f5db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b08617e4adea397d92aa662842d06571b3c01c8b5e0ffb16f000b74b0c087b97"></a>

## virtual_server.udp.client_ssl_profile — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- virtual_server.udp.client_ssl_profile

<a id="canonical-cf36b72ed7adc83c897fa39d9928abef14c1d9fe699c42548fb0f44f710d390e"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b7dd2c3e0686d22031bd1a20a9ba570a97e0d1ea6bf78eb5cb047ba8d3aaa8d"></a>

## Direct properties — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 3

<a id="canonical-16a7cf0f1b42af170076b3edda04fac7b7767653cebfe5251609fcb506f9a666"></a>

<a id="canonical-19a376fe22e9a8e22b103e6296391982915b72a7e0baf8d9c27a72cc3ddd51b9"></a>

## kind property — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 4

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

<a id="canonical-f40d5d851453282c74de895578529bbe6e7f61e24ee6645def081fec033a002c"></a>

<a id="canonical-d4b456e5dd25bbd96c5f377068068cc69fc0716b1605195d8d98dda531dcb5b7"></a>

## name property — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 5

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

<a id="canonical-a324fee4dd00c609fee1aa074ebfb80d5be9e727ee4e49883656aac1d3155444"></a>

<a id="canonical-177af0df3d028ab94f6c09483c95c86644581e0b7000b2c425d35a5b6442439f"></a>

## namespace property — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 6

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

<a id="canonical-f3b30e0bb6944a1cd826fa36b54e7e8138b77c97e05234526ca1179ed583bf02"></a>

<a id="canonical-885a7421941f0669f4bf58096f25c9f3acf7b923445d26682fe0ef1ead4b77fe"></a>

## tenant property — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 7

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

<a id="canonical-49b36101fbabc1eb3687c7e634c34f3b5c686a18c75ded5ab8032a084ae2e883"></a>

<a id="canonical-58812dbbbadd0b977fa404fc87e1ce5b20a2881e7edc4e359f5390628db68250"></a>

## uid property — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 8

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

<a id="canonical-490452f7ec0338749fc8394ba8b42b0eb677a73d79d674e17fefd64ae6c00c19"></a>

## Next pages — virtual_server.udp.client_ssl_profile / a00c0e9c8f93 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-89e8ad280ef58fc47b9f7d948b455aa0bc5d1e6dbaa07b2890f5967080cb0a43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d2c52c5335d3a308607a270d45ddbf303e75db9877c3b2d1977bc87b9f33e4a"></a>

## virtual_server.udp.server_ssl_profile — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- virtual_server.udp.server_ssl_profile

<a id="canonical-c30e16a22b12eecc18145b78a48be5565b3152522d619f01634c8f5b030f3eb2"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb3f93613d641a4dcd35db0cfc321a30565b4c6e1dc808b8fe2ab1541fdb8f11"></a>

## Direct properties — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 3

<a id="canonical-fc151e7f926a7fbafd6e3d067861f93c9d5e43df09dfa70c89a0080aa5268393"></a>

<a id="canonical-b8696985a3a4840edb9a5789f08f4fb6981486d07491e95376dbe2e6af2152ad"></a>

## kind property — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 4

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

<a id="canonical-77e98811014058dcaa2cdf867794552c90e8aeeeb0d4e20addc5939ff29d8705"></a>

<a id="canonical-bcde078b385c3887f7203d5873591a94460aefbef67fa17578853e3d15ba21d9"></a>

## name property — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 5

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

<a id="canonical-690d8ab3976fbf1658b8a3dba50eae296060ac558b2469597a2c51e10adf1117"></a>

<a id="canonical-544c6e9c8ca73bae48b4ec696de122ff7eb7f787b8d1fd5adf8ab276c651d668"></a>

## namespace property — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 6

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

<a id="canonical-2a64cae754748a3efbea5e02dd3ff2951eb2a7c893ccbea8ca242b5af216485d"></a>

<a id="canonical-bc2c23171a50239c795e1b96827b029d96aa7c340297047d4a1738a28747ab11"></a>

## tenant property — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 7

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

<a id="canonical-c77d25a478656bc55f734b9f4ee2524bf024af0be94e911b521d71f3c3806c28"></a>

<a id="canonical-44a06e2435e16299eeb6c9dea237b00e248df79f574cc03ab703dde76509593a"></a>

## uid property — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 8

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

<a id="canonical-9460a10311e698c4333191a42c85fa9421522eb78e89f2a51f3bbe7398b514ca"></a>

## Next pages — virtual_server.udp.server_ssl_profile / 7f9ffc9526c7 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-33e5527ec99dc25fa429d3665441749d79c053b3223a1703404360aa132266d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9443e8facf3049f87dff6f6b54aadfc8592a947c4a40ee3b74af59b3aee5562d"></a>

## virtual_server.udp.udp_client_profile — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- virtual_server.udp.udp_client_profile

<a id="canonical-d2baaecdbb6d1a58ac2ce6534560657553e415aa317f5bd427436a891a1f557c"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-602e711aa41ad6f2e077622634eda3fae7f93232e1754af52ab55af69a72d862"></a>

## Direct properties — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 3

<a id="canonical-8c332d1ae54035a06b8d61cc57d8fe139bf3df332541833d2d96499d57b14043"></a>

<a id="canonical-2483f936d82de3ef111fefa9078317b4589505895aaebe31dd9bf5ee5a90ddfe"></a>

## kind property — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 4

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

<a id="canonical-d27a75fc5003afbe27b487c907de84a3ff6e2ccb03754221536ed611c657fde5"></a>

<a id="canonical-92fd7c541e134921e43b748e9a5ff126603c11282d7bdf8fd37627a06f447133"></a>

## name property — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 5

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

<a id="canonical-b9549f576e68f1d0ff83fa2401e929207dfd56dc8f18230d784a28a33b205d86"></a>

<a id="canonical-ad5f851a7dfbfdc855255cb388ef1dcbba72fb24097bf33b457c5681f92e4681"></a>

## namespace property — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 6

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

<a id="canonical-0e43f8532805921fa9e0fb36a483b58b8161c32732a970c50bd3162cf933c64a"></a>

<a id="canonical-7f04fe47573f01714aff000be82fbed5da44fa16171fe2ddbf41a69fc377df26"></a>

## tenant property — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 7

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

<a id="canonical-7019f2e20d4ec4d7886ba091d0408788af586f3321fa03e29c55a6f1312d4a13"></a>

<a id="canonical-0f326d37ba64fac4d437c8df87c6c324374de55f22578331c699c286ba57c0a6"></a>

## uid property — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 8

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

<a id="canonical-c30b2ac1860d4c174e73c5e104afcb56dabecd84a5c2d26dd14ed860dd78936f"></a>

## Next pages — virtual_server.udp.udp_client_profile / 78f33ea58e53 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-5e3b83980cd726d4cfab529bb0c0403d3815769386b584aa442cba6d53fc9486"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6732eae60df672ee764385e4fce3eba02c3b0dc727dfd756ce99bebc139012ce"></a>

## virtual_server.udp.udp_server_profile — virtual_server.udp.udp_server_profile / a54fd26634c8 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- virtual_server.udp.udp_server_profile

<a id="canonical-2c6e3a7a3c1a9b7639513be87575dccd26e8ce6e41089edb7fedcef1339c6ec5"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-096da47cbc8d62ca0b211b15b3dc5f22d29a227007a2a82a438e64845b964258"></a>

## Direct properties — virtual_server.udp.udp_server_profile / a54fd26634c8 / 3

<a id="canonical-ffacc8f4ca1e60d368672b248bdc71b616122230320839f9b5607d1acbd82168"></a>

<a id="canonical-42d2ead22c0faf739475882dbe79d341dff6c04a8676962deb8e3292defbf6aa"></a>

## kind property — virtual_server.udp.udp_server_profile / a54fd26634c8 / 4

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

<a id="canonical-980d55c1edcb03c796936c7accacf2ef901d89cad5b3d1bd8a2021970ff6ed13"></a>

<a id="canonical-765cb63db517364a562e37bfb85735d96905994b381b213b325a0ecbc8e24315"></a>

## name property — virtual_server.udp.udp_server_profile / a54fd26634c8 / 5

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

<a id="canonical-012c65d41fc10bf5a36ac1ab00a095a4c2588dcdba93f935b15d0ff763c41343"></a>

<a id="canonical-eb3c58b61d848f19b42cf016b4e87abb100c308c2fbaa5785342e9b6c906b925"></a>

## namespace property — virtual_server.udp.udp_server_profile / a54fd26634c8 / 6

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

<a id="canonical-2dd5d48f60ae83deb1b0095c79ba0fa20cb3423b86379dc2d7d028688a1f8d0a"></a>

<a id="canonical-be9347edeffc3f4df785314120493da6442801a72e889435e063ad966d1bda29"></a>

## tenant property — virtual_server.udp.udp_server_profile / a54fd26634c8 / 7

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

<a id="canonical-10cff189a038bbf899d9907a0ddcb9ac6216d503f13f5caef8fe8deb86ea117e"></a>

<a id="canonical-6df91e1931a68148cde5092766d6d3fa413256e2eb6f1f80619a7cd180963fbd"></a>

## uid property — virtual_server.udp.udp_server_profile / a54fd26634c8 / 8

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

<a id="canonical-c82b79c49985e470bf8fdeb65b31799a19a79eb929e5e5f9227bbb2047687fb5"></a>

## Next pages — virtual_server.udp.udp_server_profile / a54fd26634c8 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-3b26eca3fc9a3a7609a95519ac21cca375b3782df408a853bdc25fcec9eabc05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c1444890d60e16239ac1ed4ec8bafdadfbe3672f4efdbb49cfae78f29afad29"></a>

## virtual_server.virtual_server_state — virtual_server.virtual_server_state / f2e2dccd6919 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.virtual_server_state

<a id="canonical-1acda12d7935f75c62c8c4fdd53ef74bd1232373ca5bb82f83205675a70218fa"></a>

Type: `"object"`. single nested block, Optional.

Displays the current state on the object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("state_disabled",
    "state_enabled")}
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
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

Terraform syntax:

```terraform
virtual_server_state {
  # Configure direct properties listed below.
}
```

<a id="canonical-4044435b3ab3f4c589fdaa3bdc0c745085ff374dff9b829abec825be88eaca63"></a>

## Direct properties — virtual_server.virtual_server_state / f2e2dccd6919 / 3

- [state_disabled](resources--application_profiles--reference--group-004.md#canonical-19f407a877ac42557068991f9912dc12fcc35ce60d807ebf5b6fafbc849b4ffc): complete subsection reference.

- [state_enabled](resources--application_profiles--reference--group-004.md#canonical-abb3343f12d48499e265d22e86c6b52fd59799174f31f02b8f7483e16a95ae05): complete subsection reference.
