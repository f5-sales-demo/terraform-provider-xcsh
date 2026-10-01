---
page_title: "xcsh_advertise_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy reference."
---

# xcsh_advertise_policy reference

<a id="canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-800fe521d2169bc62cb8486f8f1779a8eaf34a43e3855680a47a383aa9cd7066"></a>

## Property reference — Property reference / d7881966eb2d / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- Property reference

<a id="canonical-687b77d7ce7d84898de182763504cdeea94d8b65e2a77faa3463b32a204f42db"></a>

## Direct properties — Property reference / d7881966eb2d / 3

<a id="canonical-367aeb41e763503ef9b995d4f4fc06385c1e05fb50389a627321eab075ea047a"></a>

<a id="canonical-1523ec3a6eebd5a5fbb051cd3e1938ab7a54941b77aca1e57a80eee285b1f773"></a>

## address property — Property reference / d7881966eb2d / 4

Type: `"string"`. Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Upstream description:

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
"where" is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site object will be used based on the network type. If inside\_vip/outside\_vip is not configured in
the site object, system use interface IP in the respected networks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-cd60c457b67010b688a698063a6b2aa0a1f838a4520a4e9f4d57d8a1a2dbaf2d"></a>

<a id="canonical-ae8ed6b27bbf73150e8e6bfd68a24a251f350d1dd454e28f1ae09f61f8dc1f8f"></a>

## annotations property — Property reference / d7881966eb2d / 5

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

<a id="canonical-206b351b5bb70ac3e5c745d03b9759bd670cb057c840c199ad83b1b6cd64088c"></a>

<a id="canonical-dedd036cae3a18b7fd973e8e0a2b343eb0ce79e958bc409a78553ce34dcd9cf8"></a>

## description property — Property reference / d7881966eb2d / 6

Type: `"string"`. Computed.

Description of the AdvertisePolicy.

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

- [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-70e5e251b6b2abb434deef2cde645164c62985a016d5baed3af5165a95475371): complete subsection reference.

<a id="canonical-f1cbb7ce50347199f0663d98df5f000478765376321e760d22f09dc0a89a7a60"></a>

<a id="canonical-7e80bd947f3dda7f49d00b20d0135b515e25a51df1a103d9b56b1e9aa740a5d8"></a>

## id property — Property reference / d7881966eb2d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](data-sources--advertise_policy--reference--group-001.md#canonical-70b3026eb14610cb525915688fb07423d772d06989389e42195215fac731a357): complete subsection reference.

- [ipv6](data-sources--advertise_policy--reference--group-001.md#canonical-cce2439835b6749df7379367a63f32bbee88ace4601ae452b1d0601093449ee2): complete subsection reference.

<a id="canonical-a6962cf2001b3210bbaee657f08be0ec401b9835e324aa5bedccdb7f5136329d"></a>

<a id="canonical-5dfb14a22edbc93934b4a717b8b2d5e85d4a6a6bee720b1636ee099002998cb4"></a>

## labels property — Property reference / d7881966eb2d / 8

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

<a id="canonical-421339974fc6af59b6dc8f77d17c8b6b5ef0235bf16299a63f51f05d4fa20f01"></a>

<a id="canonical-d8f5e2c8a5fcaed887295049587bac12db49d12005538add41ce6162bd47e89d"></a>

## name property — Property reference / d7881966eb2d / 9

Type: `"string"`. Required.

Name of the AdvertisePolicy.

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

<a id="canonical-fd5b46c132724e92dd4359f140ed0152905e2ebef0cc41377af01e660bfd6d83"></a>

<a id="canonical-54326f8da8fe4cb498612bc8746ca5bd19b8c148cbaa1c7cee05ab533cca8e97"></a>

## namespace property — Property reference / d7881966eb2d / 10

Type: `"string"`. Required.

Namespace where the AdvertisePolicy exists.

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

<a id="canonical-b459e0de7a54375d0389c6fb342368288a22becdad2768763ae025c9ba70c0c4"></a>

<a id="canonical-77a454cb436a91bdde9b11bc825ab6f6536fa73d466172a393a756bc5fda00dd"></a>

## port property — Property reference / d7881966eb2d / 11

Type: `"number"`. Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Upstream description:

Exclusive with \[port\_ranges\] Port to advertise.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [port](data-sources--advertise_policy--reference--group-001.md#canonical-b459e0de7a54375d0389c6fb342368288a22becdad2768763ae025c9ba70c0c4)
- [port_ranges](data-sources--advertise_policy--reference--group-001.md#canonical-3c30e26fb3ef9bba07e9c17f7fffc8bab423476ef02302776451bc06f89276c0)

Select alternatives according to the provider validators above.

<a id="canonical-3c30e26fb3ef9bba07e9c17f7fffc8bab423476ef02302776451bc06f89276c0"></a>

<a id="canonical-1539a75ee8204a89aabf748cac40b0db65a16d26063aff03c3c5041b221d87b8"></a>

## port_ranges property — Property reference / d7881966eb2d / 12

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-db47acd0b9b9a2a7001a6a9ed394c6848cbafa303a70e9339213c4aeb758d818"></a>

<a id="canonical-eecc821311ada0c2f0186e96526ba2803076f3f25eaaff8ada6dbcd656e58120"></a>

## protocol property — Property reference / d7881966eb2d / 13

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Protocol to advertise.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](data-sources--advertise_policy--reference--group-001.md#canonical-607556708dc8302490124e253c5e1917a17eb5d99e0d18708651cd379ede6e9e): complete subsection reference.

<a id="canonical-af7ece424fad6efdc0307f6e78f861aee6b627205aa1313414f634d10439614c"></a>

<a id="canonical-7428a2968779777ada7502d9ad79d1664d4b80b9ef9acd4d4968c79a812378f1"></a>

## skip_xff_append property — Property reference / d7881966eb2d / 14

Type: `"bool"`. Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

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

- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3): complete subsection reference.

- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c): complete subsection reference.

<a id="canonical-5dd533ff2928fbb661e2c53c7592d105a798bb6363337c493518ca7a4b413337"></a>

## All schema paths — Property reference / d7881966eb2d / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--advertise_policy--reference--group-001.md#canonical-367aeb41e763503ef9b995d4f4fc06385c1e05fb50389a627321eab075ea047a) |
| `annotations` | [annotations](data-sources--advertise_policy--reference--group-001.md#canonical-cd60c457b67010b688a698063a6b2aa0a1f838a4520a4e9f4d57d8a1a2dbaf2d) |
| `description` | [description](data-sources--advertise_policy--reference--group-001.md#canonical-206b351b5bb70ac3e5c745d03b9759bd670cb057c840c199ad83b1b6cd64088c) |
| `dualstack` | [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-d4fbd8f76c3c2bed89791e5d5e47865e8d04be21e6792075c5d0cfaf04f19f42) |
| `id` | [id](data-sources--advertise_policy--reference--group-001.md#canonical-f1cbb7ce50347199f0663d98df5f000478765376321e760d22f09dc0a89a7a60) |
| `ipv4` | [ipv4](data-sources--advertise_policy--reference--group-001.md#canonical-35052be095be3ae22f2f70cdbf880f0de6d2d25804aa4c67eb5a6dfb51c2548f) |
| `ipv6` | [ipv6](data-sources--advertise_policy--reference--group-001.md#canonical-aba0356b9d4bcbdddd6fe53e72446973bdae8c055e788ff12f817d0b67442b5c) |
| `labels` | [labels](data-sources--advertise_policy--reference--group-001.md#canonical-a6962cf2001b3210bbaee657f08be0ec401b9835e324aa5bedccdb7f5136329d) |
| `name` | [name](data-sources--advertise_policy--reference--group-001.md#canonical-421339974fc6af59b6dc8f77d17c8b6b5ef0235bf16299a63f51f05d4fa20f01) |
| `namespace` | [namespace](data-sources--advertise_policy--reference--group-001.md#canonical-fd5b46c132724e92dd4359f140ed0152905e2ebef0cc41377af01e660bfd6d83) |
| `port` | [port](data-sources--advertise_policy--reference--group-001.md#canonical-b459e0de7a54375d0389c6fb342368288a22becdad2768763ae025c9ba70c0c4) |
| `port_ranges` | [port_ranges](data-sources--advertise_policy--reference--group-001.md#canonical-3c30e26fb3ef9bba07e9c17f7fffc8bab423476ef02302776451bc06f89276c0) |
| `protocol` | [protocol](data-sources--advertise_policy--reference--group-001.md#canonical-db47acd0b9b9a2a7001a6a9ed394c6848cbafa303a70e9339213c4aeb758d818) |
| `public_ip` | [public_ip](data-sources--advertise_policy--reference--group-001.md#canonical-1cee84d85c5c0ef0552b6c11cbac7db8157e2317adf49c519da601337f0b2e97) |
| `public_ip.kind` | [public_ip.kind](data-sources--advertise_policy--reference--group-001.md#canonical-9e5ee1925cc3552420cc7e4ce9495403d8420db71ec504fd413b78cc081c94c3) |
| `public_ip.name` | [public_ip.name](data-sources--advertise_policy--reference--group-001.md#canonical-091aa923c1c341c39926470920024182655a6dce046b0a33bb95cc872d37eb18) |
| `public_ip.namespace` | [public_ip.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-dbcb665244df678d182bb70712b8c4886ef1c3e406370f4f1bee5860cbda3c29) |
| `public_ip.tenant` | [public_ip.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-d7ebd74b483ba199811d99b6b1d5d7f6103f86cde0d920ae207ed04760704e61) |
| `public_ip.uid` | [public_ip.uid](data-sources--advertise_policy--reference--group-001.md#canonical-680944f9e813b7103a2c931f62a4553c2792f8a5e95891c9c6267356a03c17a7) |
| `skip_xff_append` | [skip_xff_append](data-sources--advertise_policy--reference--group-001.md#canonical-af7ece424fad6efdc0307f6e78f861aee6b627205aa1313414f634d10439614c) |
| `tls_parameters` | [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-8ccd747199751acd4f7b601bbfc46e7d4b4723fedb4141e9f2cd2b71798e92a6) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](data-sources--advertise_policy--reference--group-001.md#canonical-18bea2a4abf794de9345b9499d3f85c4e3ee7280520f67a08c2450492a552590) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](data-sources--advertise_policy--reference--group-001.md#canonical-8b7468f3c55393706daa51a1229e1fdfc4857e4bebfaccdec542f40b0cff2de7) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-8719d0cd252f769856be22f0e95034e18e21b4921a06e34c028ed5332bb9c6f2) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--advertise_policy--reference--group-001.md#canonical-62e6ca8a09bb5aecce3a473bd8736ba500752c43a58e9e2968773c434c8305fa) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--advertise_policy--reference--group-001.md#canonical-8b78d9ac31725ad486c6d7eca198bfcf7044e173c62a66599f36f8b5d3463a54) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--advertise_policy--reference--group-001.md#canonical-3d7fee966ee1fc0e09739e6914c66488d0af8aab297d36ed103eb89d09398b71) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-6a9dc83d5c5e993d2afe1d3fc7fe31f673c200c3e513ba6c995234a32e21105d) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--advertise_policy--reference--group-001.md#canonical-3af1432655eccb904dabf7748e35fe86469b3d07edb16f01ae91eb0aa9bd7869) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-0a72687f10e7e202183a1901492b2dbc98de4e8e5b4fe674989de1d20fc562a9) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-08b776fd21907b326b5b40b465c8d79591871df298a228b7fb2b0287b8c860a4) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--advertise_policy--reference--group-001.md#canonical-7f0c05b14f824707553442277eb3dda3d855f81c788a288ff8f357b7f581f221) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--advertise_policy--reference--group-001.md#canonical-ad7cf91eb8d41db1c8100b6613ada4e709dffb03af02b8ba8e1a9fe70309f0a0) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-5e9cbe26b84d305a87677cc9355ede328e7a2f1eebe997788fbe86c3947acea5) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-7a5d84f9443bd760655413a8f3413eb0e8edcb5b87150c2ad62663a78a5854f5) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--advertise_policy--reference--group-001.md#canonical-4538523189c6d7b3b0afda9741fe57d0f62cd701cf24f71a339796e54da1ac12) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--advertise_policy--reference--group-001.md#canonical-5e40fd9f56b7aceb8c34215d71416b2f9d5fc83cc0c1f2606f8e344e8e82ca41) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--advertise_policy--reference--group-001.md#canonical-f19a9daf69e43c8695d9551a4552e603f893ccdc5b2da14e07b89ead75e9a95a) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-e59f6ff8f6aee6006b7d1a6b9c46472f1548d3596a5f4b645161960012660491) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--advertise_policy--reference--group-001.md#canonical-119cf3f36da9f0f2dfc172b872e82d1271dfa31120ecab24585ee486589fd840) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--advertise_policy--reference--group-001.md#canonical-6d229d7b3d7116c2263af28a6906f37716873a4d78c83c1d5a3d1e09bb4abb2c) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--advertise_policy--reference--group-001.md#canonical-5e567aebab42776a6f9c28ef56e53ee0088915b9eee940004de1edb10cba88bd) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-2aff5616e6be69a47f02553bf6ea8af555da5d0038f96d3c1a5641869fcc6846) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--advertise_policy--reference--group-001.md#canonical-e30cb6b4a7c4e2f574c27cf57b5d8721e101212538a669aa345fdb14a3706872) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-611d38430ca0bf90f83a539304a4eb64177a2a21f3a169e1f931cfca6283aa5a) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--advertise_policy--reference--group-001.md#canonical-058458628e8a9b1aa849ab9b654ff02e03df543b5f850e2b460ed7e5c43b6021) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--advertise_policy--reference--group-001.md#canonical-7bc94f73eb115c4b6b2aa5dcf13feca06ef3c71e4c1b5ad9b683ad4daabd47d2) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--advertise_policy--reference--group-001.md#canonical-d79d2fbeae04a0f4f72c62005a92a78ea3667a4ab8619449403840a7593c96ae) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-1783553a8b6e1721e60072c1cb10cee7f192e51deaaeeeb0ed399e6b5ba4e12b) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-927593de9defd0982e8a391845554ca10a729d209aef30f891a19fa820a60223) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--advertise_policy--reference--group-001.md#canonical-6354a19c62bb8bbf197b04b2f9bbfff76631c75e9f912cd897dd1bd8df18ab07) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--advertise_policy--reference--group-001.md#canonical-cd9a72597337b2efe7bb02f832f5c888baccb155ee9c3f744249cf12ab17324d) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--advertise_policy--reference--group-001.md#canonical-b2e1cbe7a5021bfedc94da066a5f2b4c5ce2bd3d4fadfd4a82390821544e050d) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](data-sources--advertise_policy--reference--group-001.md#canonical-66bc2f249eadf0f8c133b96d6c31b10da09df8381d5fa2397ef2dc97ba83cd4d) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](data-sources--advertise_policy--reference--group-001.md#canonical-1e8e4660e5ca61a61b06ac9fe10c976ed0f6ca2a3ac31446b76ad56dc8fb2cba) |
| `where` | [where](data-sources--advertise_policy--reference--group-001.md#canonical-d3ccc0ec4a78fdd3de1a18498dc850f1a34cd11fa8d4be89d9d874536e32ec24) |
| `where.site` | [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-b52679d5eb09429f83d636afe7f01e55966f4bb9829a78f9e8ec1d7a58545241) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-d40579caac3d72f3b44515128156b7d77c6f94e6a8c1c965be3fb983e7e2ae0f) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-d9160472e7ab5b0e43f9328ddf67153432e9d7d636ef54f7d43b1dcdfbaa0beb) |
| `where.site.network_type` | [where.site.network_type](data-sources--advertise_policy--reference--group-001.md#canonical-c186fe173793eec7c0a3895cb6f745d975e38211bafb2a62c553ed069a2fe45c) |
| `where.site.ref` | [where.site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-2d833438f8472fd0502a1eb92b00d2320212655726e93185964257f6d0831848) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-1574cd68d27b4f2f96c2340a7e840d91758f038c829026e14195390d6e5c9f41) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-e298baf7acfb6c9a325a365a0acb14181803791c0f4c1fafbb41d90eebccf2f7) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-984c60899c475cda6e95e1a1dabf9285c7f2d235c748c511d5dc2d44cdab290b) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-b6b4087046e4e738c21082ff8539bec49a63bb84bede1b151d3e06c54098b6b0) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-65b961e1af8ebdfce86406830529ea7310b006474d1eda2d19c10618f870bacd) |
| `where.virtual_network` | [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-1f9161ed1b21c9918adebac9a8722bd8e92ac2d556cd0b38b57ce6cd3b56113c) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--advertise_policy--reference--group-001.md#canonical-27151b05a603e25357144b806238e653d3815b8df21818787706e47f32ad8454) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-2689d38c905159e09a15142128ec8ffa92a081af78c83697121a9c0bb557b627) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-88410dbb807e3d6d464aa8ebdf96c266c66e0e6148fddcb4e6e5a2b7ecfee376) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-3b51e2fea22e18126aa4578e43447fd4ec5545db2efd98ce4e4ca61957c8c623) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-7304c75386c43601296fc1e74b4791bfd01a46a5258b347461791a1113f6d948) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-253de281d1f2d2e380839305e22b2e17986c46bb870d226d7d5e986aa41aaee8) |
| `where.virtual_site` | [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-934e4370fb5316cca4a1bd8f8914c503978126085eb34d6f7b5076fe0f3f06d0) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-4a37d034e26fdf8244c168cf0639c9207bd43910aa5782d262667ec119a9a665) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-9230b30a51ff5b9de05b48bd0323b2c2eac94c15eb2af9fa8b887b482f0f6b51) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--advertise_policy--reference--group-001.md#canonical-1a89e7a26be89261838161c23ede592c76901f62789ee108129fe5c670dddef5) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-56d53bd38c202376eae5e843e2eadee6191f51d685d1ce6e66ec28807c75f630) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-9fa884cac5fd555f7203e6e7a36d93b9b976678c40f96dd9dac8d9f007fe2f20) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-c419db7b9fc0bd4daf7125f47244fd1be891d77ebfca927f9d05e83495952b37) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-cfb33a04497503311d6c6ea6107dd5d8036f1020c5e5bc0089109fd49cad1db8) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-433a4e12bbb82392962d91d6a29d288d10a43167df4b66ff0310cfecaa6e2d46) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-fd37c991f8dd804438ee3c386d1a6bd251ff87a2a9724a1e2e1a3253ce821097) |

<a id="canonical-6b02e4f007102725e0566084a65d86fe3c656597fc991a893c7e9c3e58696eb7"></a>

## Next pages — Property reference / d7881966eb2d / 16

- [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-70e5e251b6b2abb434deef2cde645164c62985a016d5baed3af5165a95475371)
- [ipv4](data-sources--advertise_policy--reference--group-001.md#canonical-70b3026eb14610cb525915688fb07423d772d06989389e42195215fac731a357)
- [ipv6](data-sources--advertise_policy--reference--group-001.md#canonical-cce2439835b6749df7379367a63f32bbee88ace4601ae452b1d0601093449ee2)
- [public_ip](data-sources--advertise_policy--reference--group-001.md#canonical-607556708dc8302490124e253c5e1917a17eb5d99e0d18708651cd379ede6e9e)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-70e5e251b6b2abb434deef2cde645164c62985a016d5baed3af5165a95475371"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b7b43a3589f3df80a6309eb424f51b14cd1ecfeb8ae8760e2d53ce28625a396"></a>

## dualstack — dualstack / 10005e04ddfb / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- dualstack

<a id="canonical-d4fbd8f76c3c2bed89791e5d5e47865e8d04be21e6792075c5d0cfaf04f19f42"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

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

- [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-d4fbd8f76c3c2bed89791e5d5e47865e8d04be21e6792075c5d0cfaf04f19f42)
- [ipv4](data-sources--advertise_policy--reference--group-001.md#canonical-35052be095be3ae22f2f70cdbf880f0de6d2d25804aa4c67eb5a6dfb51c2548f)
- [ipv6](data-sources--advertise_policy--reference--group-001.md#canonical-aba0356b9d4bcbdddd6fe53e72446973bdae8c055e788ff12f817d0b67442b5c)

Select alternatives according to the provider validators above.

<a id="canonical-e7f87ff14c5f233fc0017d7a06786ceaa57268dc19deeb26ef5051a6b8ff743b"></a>

## Direct properties — dualstack / 10005e04ddfb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8fa77a81cb14dca066f818a72d7b6f6b7601f7c48419216f143a98b0e206994"></a>

## Next pages — dualstack / 10005e04ddfb / 4

- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-70b3026eb14610cb525915688fb07423d772d06989389e42195215fac731a357"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bc0a8d8ffa522d64fcc2e23de1e9dd2a90823f555571864637cc82231c465ca"></a>

## ipv4 — ipv4 / 7da7ff6c3130 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- ipv4

<a id="canonical-35052be095be3ae22f2f70cdbf880f0de6d2d25804aa4c67eb5a6dfb51c2548f"></a>

Type: `["object", {}]`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

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

<a id="canonical-d8aa9c2215143562f7b95458b0892c6885547fdd4a5d0dca70806bbe23a02a3a"></a>

## Direct properties — ipv4 / 7da7ff6c3130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8f25b9f6442efb068dc70fa49149637d02866ec3dd5eee4e3d169146d4138b8"></a>

## Next pages — ipv4 / 7da7ff6c3130 / 4

- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-cce2439835b6749df7379367a63f32bbee88ace4601ae452b1d0601093449ee2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87f7279f20582ab66f85a7535bdd88d0a3799875b52dd9c4bad015c443e0a993"></a>

## ipv6 — ipv6 / d1332cabdab2 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- ipv6

<a id="canonical-aba0356b9d4bcbdddd6fe53e72446973bdae8c055e788ff12f817d0b67442b5c"></a>

Type: `["object", {}]`. Computed.

IPv6 address in colon-separated hexadecimal format.

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

<a id="canonical-b57f7e59726c2e4560f4c4beb0e19e58d7a76756a5e8e440b6292a6cddcf6a94"></a>

## Direct properties — ipv6 / d1332cabdab2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12cff60aad1c6550a019f24d1fa1bd1eb982c290d117184ad8c81be0ecba0c31"></a>

## Next pages — ipv6 / d1332cabdab2 / 4

- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-607556708dc8302490124e253c5e1917a17eb5d99e0d18708651cd379ede6e9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b542aefb647ce5450b8a2462694603704aa128bea9759f47bb20cab714170d90"></a>

## public_ip — public_ip / a919bfe9bd30 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- public_ip

<a id="canonical-1cee84d85c5c0ef0552b6c11cbac7db8157e2317adf49c519da601337f0b2e97"></a>

Type: `"list"`. Computed.

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

Upstream description:

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

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

<a id="canonical-c3e4fc6743e5b8e88d3922e269957dc9439a893d4a8a3a0def3433ed0bcd061c"></a>

## Direct properties — public_ip / a919bfe9bd30 / 3

<a id="canonical-9e5ee1925cc3552420cc7e4ce9495403d8420db71ec504fd413b78cc081c94c3"></a>

<a id="canonical-f0492e128f7a324c187441e2fba772cd43f9dc1d5280681ffe139d86cb33d805"></a>

## kind property — public_ip / a919bfe9bd30 / 4

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

<a id="canonical-091aa923c1c341c39926470920024182655a6dce046b0a33bb95cc872d37eb18"></a>

<a id="canonical-b8c3b4985054d7695ef4916ecbd514288c52d41aadf41f74ab75314b0b316373"></a>

## name property — public_ip / a919bfe9bd30 / 5

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

<a id="canonical-dbcb665244df678d182bb70712b8c4886ef1c3e406370f4f1bee5860cbda3c29"></a>

<a id="canonical-cbfd156b80ccb9d14c58ad03de1fdbde49d03a649d64d45b53c7650ab7f0a845"></a>

## namespace property — public_ip / a919bfe9bd30 / 6

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

<a id="canonical-d7ebd74b483ba199811d99b6b1d5d7f6103f86cde0d920ae207ed04760704e61"></a>

<a id="canonical-4dace8b9a3c0fd5f015b6e3a857d7282eec7adac6d658455469484ad95428c57"></a>

## tenant property — public_ip / a919bfe9bd30 / 7

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

<a id="canonical-680944f9e813b7103a2c931f62a4553c2792f8a5e95891c9c6267356a03c17a7"></a>

<a id="canonical-ec20ee853c91318d5da5f1b72293d3197f57e6b9b567eebc84b015b6632e6469"></a>

## uid property — public_ip / a919bfe9bd30 / 8

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

<a id="canonical-ff8c818c0d56a4d60094eedb151ae31192de64511db565b0c440f5ff86db5f69"></a>

## Next pages — public_ip / a919bfe9bd30 / 9

- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a86655dd673d0e86e553df9bd96bbd2d1196a9416c01c39c294e4207ac2a687"></a>

## tls_parameters — tls_parameters / 49205be770ba / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- tls_parameters

<a id="canonical-8ccd747199751acd4f7b601bbfc46e7d4b4723fedb4141e9f2cd2b71798e92a6"></a>

Type: `"single"`. Computed.

TLS configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

<a id="canonical-8fa8e5451f4f6cd472ed1fcd96b5b4584dc65c444df3a20f2525cf6e4c275cc6"></a>

## Direct properties — tls_parameters / 49205be770ba / 3

- [client_certificate_optional](data-sources--advertise_policy--reference--group-001.md#canonical-c8c421be48805b49c3ae3686f651f17ceea6e56574a1fc510974892f480dc7c7): complete subsection reference.

- [client_certificate_required](data-sources--advertise_policy--reference--group-001.md#canonical-afb4c0db2405c1b3f546f87ba41f171e695e8f55965cbb43e8a1bf7f31a8eeb2): complete subsection reference.

- [common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782): complete subsection reference.

- [no_client_certificate](data-sources--advertise_policy--reference--group-001.md#canonical-ab9a8555afb175a0d04e3051cc1a8cf3558f56a7bd2f1f20bd04557269873c55): complete subsection reference.

<a id="canonical-1e8e4660e5ca61a61b06ac9fe10c976ed0f6ca2a3ac31446b76ad56dc8fb2cba"></a>

<a id="canonical-7905554314f331c7ddc9a31d0f1e4bf344f05270995a2e7be739ea8120dcbf1a"></a>

## xfcc_header_elements property — tls_parameters / 49205be770ba / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bd8d7376731198fe77a4ef1de319e984ecbd006edc71affe9f1448fcc5d43964"></a>

## Next pages — tls_parameters / 49205be770ba / 5

- [tls_parameters.client_certificate_optional](data-sources--advertise_policy--reference--group-001.md#canonical-c8c421be48805b49c3ae3686f651f17ceea6e56574a1fc510974892f480dc7c7)
- [tls_parameters.client_certificate_required](data-sources--advertise_policy--reference--group-001.md#canonical-afb4c0db2405c1b3f546f87ba41f171e695e8f55965cbb43e8a1bf7f31a8eeb2)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.no_client_certificate](data-sources--advertise_policy--reference--group-001.md#canonical-ab9a8555afb175a0d04e3051cc1a8cf3558f56a7bd2f1f20bd04557269873c55)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-c8c421be48805b49c3ae3686f651f17ceea6e56574a1fc510974892f480dc7c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b7409bb738a6453b52161644dea585616928c93690a4f5e1f69fcd6f14a8f11"></a>

## tls_parameters.client_certificate_optional — tls_parameters.client_certificate_optional / 7d9ab3ed6341 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- tls_parameters.client_certificate_optional

<a id="canonical-18bea2a4abf794de9345b9499d3f85c4e3ee7280520f67a08c2450492a552590"></a>

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

<a id="canonical-c72f281c16e07c4145ee913fe749df5de7cd8cd9fc942fbb991463cecde195c1"></a>

## Direct properties — tls_parameters.client_certificate_optional / 7d9ab3ed6341 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f08aae2184b0fe8ed68e1bdee5fca50a80c72272906723c903541975e31b5c54"></a>

## Next pages — tls_parameters.client_certificate_optional / 7d9ab3ed6341 / 4

- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-afb4c0db2405c1b3f546f87ba41f171e695e8f55965cbb43e8a1bf7f31a8eeb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2695cbbf4e8eb47916a383ac1caabf200d7ccfc3dd3a87012cb65caf381c27c6"></a>

## tls_parameters.client_certificate_required — tls_parameters.client_certificate_required / a40125b9c7a4 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- tls_parameters.client_certificate_required

<a id="canonical-8b7468f3c55393706daa51a1229e1fdfc4857e4bebfaccdec542f40b0cff2de7"></a>

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

<a id="canonical-ca618f4a90866990339bb6adedd0de9b4c96808562bfea58b7b953d88abc7c06"></a>

## Direct properties — tls_parameters.client_certificate_required / a40125b9c7a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28feb3a7609307c2a62bfff29bde7110665e074ee1a88f38bd3f9f34d02b1dd6"></a>

## Next pages — tls_parameters.client_certificate_required / a40125b9c7a4 / 4

- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aca739edd626e1fd2d1ebeeaae5da98621f16bb0674254993fa87bd26b0ddc3b"></a>

## tls_parameters.common_params — tls_parameters.common_params / 58a8e69c014b / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- tls_parameters.common_params

<a id="canonical-8719d0cd252f769856be22f0e95034e18e21b4921a06e34c028ed5332bb9c6f2"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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

<a id="canonical-75892df5c63f233ef4ac27ee40ecbd73827dc1fdc969fee513df9a1063d3c8af"></a>

## Direct properties — tls_parameters.common_params / 58a8e69c014b / 3

<a id="canonical-62e6ca8a09bb5aecce3a473bd8736ba500752c43a58e9e2968773c434c8305fa"></a>

<a id="canonical-59036b9e26a4cd23bfb2956625a712c0ccf7ce84729cc363ccf94391a427a917"></a>

## cipher_suites property — tls_parameters.common_params / 58a8e69c014b / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8b78d9ac31725ad486c6d7eca198bfcf7044e173c62a66599f36f8b5d3463a54"></a>

<a id="canonical-c40337cae25a3be66daaeeefba332195e30d39814592d6989d637bf4d1c5847e"></a>

## maximum_protocol_version property — tls_parameters.common_params / 58a8e69c014b / 5

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

<a id="canonical-3d7fee966ee1fc0e09739e6914c66488d0af8aab297d36ed103eb89d09398b71"></a>

<a id="canonical-862f405e9d2cfaab849f3451a962860d246b5df4edd35f3eae2f84d46796809d"></a>

## minimum_protocol_version property — tls_parameters.common_params / 58a8e69c014b / 6

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

- [tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c): complete subsection reference.

- [validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868): complete subsection reference.

<a id="canonical-e0e5112df0c5124caa39048221c8996f1477cc7a493abcb2b5e554e9054cc5eb"></a>

## Next pages — tls_parameters.common_params / 58a8e69c014b / 7

- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2f22ce3b38c6615e776448425cd746e381b09c200f69a6c9f072a9ed3866fd1"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / de4c6dabb191 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- tls_parameters.common_params.tls_certificates

<a id="canonical-6a9dc83d5c5e993d2afe1d3fc7fe31f673c200c3e513ba6c995234a32e21105d"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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

<a id="canonical-c04e09c334ea0c5792cea7fab3a90953c99f292ab41ab57bbd100fd0453f058b"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / de4c6dabb191 / 3

<a id="canonical-3af1432655eccb904dabf7748e35fe86469b3d07edb16f01ae91eb0aa9bd7869"></a>

<a id="canonical-32f214afe3b3533beb346b03f68c4139f19fcf3d667e61967958bc0747e52d45"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / de4c6dabb191 / 4

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

- [custom_hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-399acec1b2b5a8d7d3d201b6e472f118abf4a96f1b359771b76746e21c5a31bf): complete subsection reference.

<a id="canonical-7f0c05b14f824707553442277eb3dda3d855f81c788a288ff8f357b7f581f221"></a>

<a id="canonical-8b7396740f64b4a35b518c61b90eb2446f087ddbb443f35073c27fa7892d39d3"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / de4c6dabb191 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--advertise_policy--reference--group-001.md#canonical-caf91004400138683b3279b991dd40ad8b70073393bbd680f13d5d493608a393): complete subsection reference.

- [private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73): complete subsection reference.

- [use_system_defaults](data-sources--advertise_policy--reference--group-001.md#canonical-e1d7c2a3fa6a254cd35f2a9623895d0e6426f7329fc0531f10f004ade3de13af): complete subsection reference.

<a id="canonical-682a65e7173e42c6eda73cb418aa5c1283b6f61abbb6d433bcfc5b6292dc9f0f"></a>

## Next pages — tls_parameters.common_params.tls_certificates / de4c6dabb191 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-399acec1b2b5a8d7d3d201b6e472f118abf4a96f1b359771b76746e21c5a31bf)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--advertise_policy--reference--group-001.md#canonical-caf91004400138683b3279b991dd40ad8b70073393bbd680f13d5d493608a393)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--advertise_policy--reference--group-001.md#canonical-e1d7c2a3fa6a254cd35f2a9623895d0e6426f7329fc0531f10f004ade3de13af)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-399acec1b2b5a8d7d3d201b6e472f118abf4a96f1b359771b76746e21c5a31bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1779064d876c202f598d73113e3baa38eaf469641a4e763eb21aed4d5be4601"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / aa0111d53c9d / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-0a72687f10e7e202183a1901492b2dbc98de4e8e5b4fe674989de1d20fc562a9"></a>

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

<a id="canonical-d7ee564947520604c830a8870f1534d12bc96e1f619bf2b9287fd91a9d0ee7f3"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / aa0111d53c9d / 3

<a id="canonical-08b776fd21907b326b5b40b465c8d79591871df298a228b7fb2b0287b8c860a4"></a>

<a id="canonical-a3bea38fb60454cff07bbb3928484c234ca7bba5cb872a0d6b5640ca807ca914"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / aa0111d53c9d / 4

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

<a id="canonical-22fd3991bcf8e0a2244ed5b082c07276ba67f64d475d16eee760ea8a39ba20ac"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / aa0111d53c9d / 5

- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-caf91004400138683b3279b991dd40ad8b70073393bbd680f13d5d493608a393"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1439e6e501e7357cfac4269f4918807eabaea5fd093b10fc85765b788f8477d"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 1667c306013e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-ad7cf91eb8d41db1c8100b6613ada4e709dffb03af02b8ba8e1a9fe70309f0a0"></a>

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

<a id="canonical-979f88252ad8c17dc7eac539c45d649641f5cd6159f448ff9a128535a4ea98ef"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 1667c306013e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d071e3a29b1371b26f5356f64014a756bdae76aad91c6d426a196bf7c7001cee"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 1667c306013e / 4

- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4619c4f9e6c6c6679cdbc29c9c30bd2d54b0b32842e295379a83eaccad48889f"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / 94c7f805f54b / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-5e9cbe26b84d305a87677cc9355ede328e7a2f1eebe997788fbe86c3947acea5"></a>

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

<a id="canonical-0e0fcb02e38ef4c4005a3401b8ee131c0aa4ad4cde17c6d6e6f5c776b7df17c6"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / 94c7f805f54b / 3

- [blindfold_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-8a36a8b8498ddd2a782d3ca9637b2eea6cb2b7d97a81b59b8b7c99920b9ddf7d): complete subsection reference.

- [clear_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-3e5b8891b67cc4851a5513ba101d67729eeabb37e6dbb961c063197d79b5bc1e): complete subsection reference.

<a id="canonical-a2ba28a64e231e8dcf0201393108683f063d7d0140d02e0b593afc6a1598c6ef"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / 94c7f805f54b / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-8a36a8b8498ddd2a782d3ca9637b2eea6cb2b7d97a81b59b8b7c99920b9ddf7d)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-3e5b8891b67cc4851a5513ba101d67729eeabb37e6dbb961c063197d79b5bc1e)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-8a36a8b8498ddd2a782d3ca9637b2eea6cb2b7d97a81b59b8b7c99920b9ddf7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68e463bc7bbfaaf7c8be966f4cd656a625e84852be679d8d9ea78d6cf60bf4c8"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-7a5d84f9443bd760655413a8f3413eb0e8edcb5b87150c2ad62663a78a5854f5"></a>

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

<a id="canonical-19832f1a03e7c8cc405cbcc1928488b0ab9c688bab1470d6605f8b5d3bf3b74f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 3

<a id="canonical-4538523189c6d7b3b0afda9741fe57d0f62cd701cf24f71a339796e54da1ac12"></a>

<a id="canonical-a8e30664b21f25d8001d68be6c2fce8e102eaa01ed6662bb20bc91861e2c7c2a"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 4

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

<a id="canonical-5e40fd9f56b7aceb8c34215d71416b2f9d5fc83cc0c1f2606f8e344e8e82ca41"></a>

<a id="canonical-b0ebd86f71b6b8d45581ccf7e14f57e7a48f892e142a824da442c6294313ee3b"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 5

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

<a id="canonical-f19a9daf69e43c8695d9551a4552e603f893ccdc5b2da14e07b89ead75e9a95a"></a>

<a id="canonical-e698a743bb2099876145f673793111532aa1b6ed4a4b298b27ea1bd85f083231"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 6

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

<a id="canonical-8d28a42e3b6fe89c602da29e2937538be649bf0715a8d9ec4d6a7a3986bf4888"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 47c8b2ed6b92 / 7

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-3e5b8891b67cc4851a5513ba101d67729eeabb37e6dbb961c063197d79b5bc1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a12724edac15ecd64cdc087f531cf29ae55061f20a7f49b1c60e2d0cb336ce4"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / b287331d9574 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-e59f6ff8f6aee6006b7d1a6b9c46472f1548d3596a5f4b645161960012660491"></a>

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

<a id="canonical-68a954ce581b9bb0a345898600dc818b6d93f056faa787053b789ff993f7942a"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / b287331d9574 / 3

<a id="canonical-119cf3f36da9f0f2dfc172b872e82d1271dfa31120ecab24585ee486589fd840"></a>

<a id="canonical-0ea4ec0f3400b1e86af49d8c79d67ad3b575ea279f0c2446d7bf1d06efdc8d16"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / b287331d9574 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6d229d7b3d7116c2263af28a6906f37716873a4d78c83c1d5a3d1e09bb4abb2c"></a>

<a id="canonical-33d78ef971d1db04a1aad3c0307ed3542f4fd988206ce0c22b5683f28c7a381b"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / b287331d9574 / 5

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

<a id="canonical-1fb957abcafbf5e88c8939843f4ba59c809769e3be39f772a9544a1d6273a0c8"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / b287331d9574 / 6

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-3b3373049341e8342fb910b303cb66305b44ab8a2efaea97409f89c3716d7c73)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-e1d7c2a3fa6a254cd35f2a9623895d0e6426f7329fc0531f10f004ade3de13af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fd391f09470002d04febc64a13572bd1730fa961e2cd1a78f59c007260bd540"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / a7a0ed85df01 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-5e567aebab42776a6f9c28ef56e53ee0088915b9eee940004de1edb10cba88bd"></a>

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

<a id="canonical-5f4322b537e0e073aed0b06cde590a0bf544fb64887459d9f4f3e12bae55a1d0"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / a7a0ed85df01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0d94e616c9976a83ac41879b95d85f69801a48e067de40b62adb5e1c4c50bc2"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / a7a0ed85df01 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-e6853428caffe6cb4dde23730a29447a8cb2a3c027a623a93b01f6f32a258a5c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecafd3c42c1b97d3fd19d9f045cb4d890e743a40307a02a87059504e9ed47803"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- tls_parameters.common_params.validation_params

<a id="canonical-2aff5616e6be69a47f02553bf6ea8af555da5d0038f96d3c1a5641869fcc6846"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-dda0a4256fe7da4c7d8f56481b99983e2eb6fd29e7f01cb81125b07fb909777d"></a>

## Direct properties — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 3

<a id="canonical-e30cb6b4a7c4e2f574c27cf57b5d8721e101212538a669aa345fdb14a3706872"></a>

<a id="canonical-212d8ef3d503ea8e6e51d7655abc0dc00631cd7bed54223211d1efe2c3e9b2a2"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-db1fb54ae99f7306d70e4b77c6287f03416274551959f1b65b1eb71c5c67cc0f): complete subsection reference.

<a id="canonical-cd9a72597337b2efe7bb02f832f5c888baccb155ee9c3f744249cf12ab17324d"></a>

<a id="canonical-a4c59fba1519d8daf0e0812679929382abd6a97d214fa96c85898b8aaf80284e"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-b2e1cbe7a5021bfedc94da066a5f2b4c5ce2bd3d4fadfd4a82390821544e050d"></a>

<a id="canonical-e57b99e80978398206a5965bee48f46014e502734b1921c3a52cd9f2de5ece2b"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-1f999215a6264265c647a7833307e78a626b89611a814b7480783c4b9a80406b"></a>

## Next pages — tls_parameters.common_params.validation_params / 2c32ed9f5117 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-db1fb54ae99f7306d70e4b77c6287f03416274551959f1b65b1eb71c5c67cc0f)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-db1fb54ae99f7306d70e4b77c6287f03416274551959f1b65b1eb71c5c67cc0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf2a3dc7559c50cff360b025b334af96de3b814ce420b85ccd8e4b9e4222b1c5"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / d9a951cb99c1 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-611d38430ca0bf90f83a539304a4eb64177a2a21f3a169e1f931cfca6283aa5a"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-51932eaa521c5f0a81ab315620a5368411286263c3a1d93351d5dfb18cfd0450"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / d9a951cb99c1 / 3

- [trusted_ca_list](data-sources--advertise_policy--reference--group-001.md#canonical-2e8ab692a659cea7cc4c3d91932efeaa168299cd13e7b488db95ec574dee6c26): complete subsection reference.

<a id="canonical-c60bef2d5e3f0e17cbbeee844d907f83a4b94ca179a91dabd08bc8b51ad66c61"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / d9a951cb99c1 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--advertise_policy--reference--group-001.md#canonical-2e8ab692a659cea7cc4c3d91932efeaa168299cd13e7b488db95ec574dee6c26)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-2e8ab692a659cea7cc4c3d91932efeaa168299cd13e7b488db95ec574dee6c26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-700c56269c62bab7e2db0a583839bcc8630e08d021ddd084d534af89793d4be4"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-f768a150501f3ce36ff75c06774cc564c91370cb1ccdd60f0a43777322710782)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-3a7cceb9b8b37570b3dd79052b33ad90bf4cce568aeabf10d15b7a8d33e42868)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-db1fb54ae99f7306d70e4b77c6287f03416274551959f1b65b1eb71c5c67cc0f)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-058458628e8a9b1aa849ab9b654ff02e03df543b5f850e2b460ed7e5c43b6021"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-c2ab60a4eeb5135547387abf96399edc38a2ae3076c28acd5146710dc0273b41"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 3

<a id="canonical-7bc94f73eb115c4b6b2aa5dcf13feca06ef3c71e4c1b5ad9b683ad4daabd47d2"></a>

<a id="canonical-bf8cdf7aa55ad7b88ea2b6ff05fb7b41620f7d133dd8f35cd2af9166d87e5e47"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 4

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

<a id="canonical-d79d2fbeae04a0f4f72c62005a92a78ea3667a4ab8619449403840a7593c96ae"></a>

<a id="canonical-736cf94e58565f286426dab107ef2256f2e4dff6b35f72232877baeb421eb068"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 5

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

<a id="canonical-1783553a8b6e1721e60072c1cb10cee7f192e51deaaeeeb0ed399e6b5ba4e12b"></a>

<a id="canonical-31b5bc7cedd5e27c00b4cae102f75cee6b1a4bc4e589ef000af8c3cdb207a101"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 6

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

<a id="canonical-927593de9defd0982e8a391845554ca10a729d209aef30f891a19fa820a60223"></a>

<a id="canonical-d71ccd7390f262b9f0a6aded1b245e80f4b0d17486d0c51920f8b2876f2f2e3b"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 7

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

<a id="canonical-6354a19c62bb8bbf197b04b2f9bbfff76631c75e9f912cd897dd1bd8df18ab07"></a>

<a id="canonical-cb8590e2c96eaa71efe884bdb3ca2c6839c77dd08a9057c2b2c969bb6fc02b8c"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 8

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

<a id="canonical-8d23f83bdcb687ca69d28a088f2bb6a171701c82dfb9f0d5926913fe1e1cff89"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 1ec83efa3ce0 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-db1fb54ae99f7306d70e4b77c6287f03416274551959f1b65b1eb71c5c67cc0f)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-ab9a8555afb175a0d04e3051cc1a8cf3558f56a7bd2f1f20bd04557269873c55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11b2ddcc72b5e575d9a6c5053490dfc3b899d0bf5a2ed24334bd338cdc941f93"></a>

## tls_parameters.no_client_certificate — tls_parameters.no_client_certificate / 7f887eca6040 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- tls_parameters.no_client_certificate

<a id="canonical-66bc2f249eadf0f8c133b96d6c31b10da09df8381d5fa2397ef2dc97ba83cd4d"></a>

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

<a id="canonical-20a301a18dcdbe705a79d1e8cab75d4e95d756ebd501ea1e01957f8d38ffeaf5"></a>

## Direct properties — tls_parameters.no_client_certificate / 7f887eca6040 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9f0184bcc63669d743f9c24a9417119756e457b0ef25c89ca226c76c95f67d7"></a>

## Next pages — tls_parameters.no_client_certificate / 7f887eca6040 / 4

- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-7f6fc7c2de8a1033a4cc1afb3b3677c7ce0e50671b79b916742f16c0b2c43fd3)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0006f1e7e7d223c77c744f249d4b1727099c334725735f374ec337e7c0b5a1ae"></a>

## where — where / d72ae4460574 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- where

<a id="canonical-d3ccc0ec4a78fdd3de1a18498dc850f1a34cd11fa8d4be89d9d874536e32ec24"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-3b167c82f7f5addb93fb1df1b31ecc3d03f07dd7c22874f8a0ffd35744647f8a"></a>

## Direct properties — where / d72ae4460574 / 3

- [site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a): complete subsection reference.

- [virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-ca675ca3d165af7e4c1f190c7c8867d6931a560ae2cded7aad807536edad4f5f): complete subsection reference.

- [virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4): complete subsection reference.

<a id="canonical-170b9c940a85bdb7d0daaab593f9182a19d3c3af462e1a2ad59ce0d0b055d0c3"></a>

## Next pages — where / d72ae4460574 / 4

- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-ca675ca3d165af7e4c1f190c7c8867d6931a560ae2cded7aad807536edad4f5f)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-378e5d1b01a1d7a9956b5ac75b747798781709708f87a5db5ff2959c4fe0337b"></a>

## where.site — where.site / a569a9b8fcc6 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- where.site

<a id="canonical-b52679d5eb09429f83d636afe7f01e55966f4bb9829a78f9e8ec1d7a58545241"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

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

<a id="canonical-e44383c5475cdce006f809761346861842cdbcd16b1267a0241a014841fc3a27"></a>

## Direct properties — where.site / a569a9b8fcc6 / 3

- [disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-491ca204f70467a190fda0b7592a3ce45edd9b802018dcbe1c2f2ff0d5f90bab): complete subsection reference.

- [enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-9f4373c2d4e6ac523aa385f8d8a174bd6b4a035e52d004455043f39bdbce07d6): complete subsection reference.

<a id="canonical-c186fe173793eec7c0a3895cb6f745d975e38211bafb2a62c553ed069a2fe45c"></a>

<a id="canonical-3b312b258b6e32869ab7a852137057ab453246423ecbedc235081043d19ddf91"></a>

## network_type property — where.site / a569a9b8fcc6 / 4

Type: `"string"`. Computed.

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

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-471e8a9443538f17aefe7ca51306e5ad88a34b7707659fef2c56cf918a1d828c): complete subsection reference.

<a id="canonical-0e2e1e2784e60184887a6f80e45326c5bd7b1e498f6da267715303a5b6a12784"></a>

## Next pages — where.site / a569a9b8fcc6 / 5

- [where.site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-491ca204f70467a190fda0b7592a3ce45edd9b802018dcbe1c2f2ff0d5f90bab)
- [where.site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-9f4373c2d4e6ac523aa385f8d8a174bd6b4a035e52d004455043f39bdbce07d6)
- [where.site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-471e8a9443538f17aefe7ca51306e5ad88a34b7707659fef2c56cf918a1d828c)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-491ca204f70467a190fda0b7592a3ce45edd9b802018dcbe1c2f2ff0d5f90bab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aad9e1efa977f7f931bb29c623a1e9af5b0c34ede5c466ff8a1d67ccc075455"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / c34ebd29ff76 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- where.site.disable_internet_vip

<a id="canonical-d40579caac3d72f3b44515128156b7d77c6f94e6a8c1c965be3fb983e7e2ae0f"></a>

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

<a id="canonical-5a86f3d49bbb63ef15464693c753f54df1219c039a42b8277ebec43577259456"></a>

## Direct properties — where.site.disable_internet_vip / c34ebd29ff76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c96ba85964fe0113b7064315b4e5ccb1d6d7b39f55ac367bc913e20919ff06d"></a>

## Next pages — where.site.disable_internet_vip / c34ebd29ff76 / 4

- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-9f4373c2d4e6ac523aa385f8d8a174bd6b4a035e52d004455043f39bdbce07d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-019291b3045b46f1f48e6b2536e2df07d09cfdd8d0d01fa3ae62c55d7a75d5b8"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 1da94b98c9a1 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- where.site.enable_internet_vip

<a id="canonical-d9160472e7ab5b0e43f9328ddf67153432e9d7d636ef54f7d43b1dcdfbaa0beb"></a>

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

<a id="canonical-277d06f0c6499b4aef2e92d673178e15260900829a85452dc7e5fe79210b5cf9"></a>

## Direct properties — where.site.enable_internet_vip / 1da94b98c9a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eefebdd7d07886619e69cb54bb95603cb274265ed9f8b986101bb4906076ade8"></a>

## Next pages — where.site.enable_internet_vip / 1da94b98c9a1 / 4

- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-471e8a9443538f17aefe7ca51306e5ad88a34b7707659fef2c56cf918a1d828c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-453ad0259209624308e36f5a4f5ff378fe03e10034e01cd4ee5f81bbe0f5fe7a"></a>

## where.site.ref — where.site.ref / 0a8a482cd11c / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- where.site.ref

<a id="canonical-2d833438f8472fd0502a1eb92b00d2320212655726e93185964257f6d0831848"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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

<a id="canonical-4fe537c054dd67f2db2ef0d10e0ac73e9c73559ea2608c949c1d052498f5bad7"></a>

## Direct properties — where.site.ref / 0a8a482cd11c / 3

<a id="canonical-1574cd68d27b4f2f96c2340a7e840d91758f038c829026e14195390d6e5c9f41"></a>

<a id="canonical-6bd87f79d93c33822f7af107d3b07531992dc5bed33b09ed5c710635465e5c31"></a>

## kind property — where.site.ref / 0a8a482cd11c / 4

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

<a id="canonical-e298baf7acfb6c9a325a365a0acb14181803791c0f4c1fafbb41d90eebccf2f7"></a>

<a id="canonical-260f534a98a78736ef13d6bf65d52b125605fa9225746cf3a31fa6a6de1e5556"></a>

## name property — where.site.ref / 0a8a482cd11c / 5

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

<a id="canonical-984c60899c475cda6e95e1a1dabf9285c7f2d235c748c511d5dc2d44cdab290b"></a>

<a id="canonical-c868ddb53dbd36840c727c44c99a5ed07aa1a470aa15d61a4a76b56166a7f8be"></a>

## namespace property — where.site.ref / 0a8a482cd11c / 6

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

<a id="canonical-b6b4087046e4e738c21082ff8539bec49a63bb84bede1b151d3e06c54098b6b0"></a>

<a id="canonical-c1ce3cc3546363d4120e2cf89fc024478b6e39dfee30b1fb54cad2bc93d03e66"></a>

## tenant property — where.site.ref / 0a8a482cd11c / 7

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

<a id="canonical-65b961e1af8ebdfce86406830529ea7310b006474d1eda2d19c10618f870bacd"></a>

<a id="canonical-34dabfe37bd0db2332b2ba72bb7d3e55b022f761310313b6024e01366c55a4de"></a>

## uid property — where.site.ref / 0a8a482cd11c / 8

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

<a id="canonical-abe310abf3b52aa45607f53fdb5f42a6ddc2978e14ba4d4348aeadb27753e747"></a>

## Next pages — where.site.ref / 0a8a482cd11c / 9

- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-96067c3581760eba33251ff8d5065664ba850ee2b0cba0cc3f9a0a5a621fce5a)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-ca675ca3d165af7e4c1f190c7c8867d6931a560ae2cded7aad807536edad4f5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7473f65148bc31fcf3c620a48fb6134c49ede20e6480f6837679a2a58deb918"></a>

## where.virtual_network — where.virtual_network / 1c9620510044 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- where.virtual_network

<a id="canonical-1f9161ed1b21c9918adebac9a8722bd8e92ac2d556cd0b38b57ce6cd3b56113c"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

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

<a id="canonical-ed0600f12e4677ebfad334dddc74bbfd4713c32a0ceb051765a76edcef6ada9e"></a>

## Direct properties — where.virtual_network / 1c9620510044 / 3

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-19f9bcfd38f5d3427f33a5ded73ab7309a75a6727205583a3a496e0d06d86544): complete subsection reference.

<a id="canonical-e3d00a4aca3d5a789f8126dd48cc36967b1baa4b4cce3e0affa6c80edd6770c7"></a>

## Next pages — where.virtual_network / 1c9620510044 / 4

- [where.virtual_network.ref](data-sources--advertise_policy--reference--group-001.md#canonical-19f9bcfd38f5d3427f33a5ded73ab7309a75a6727205583a3a496e0d06d86544)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-19f9bcfd38f5d3427f33a5ded73ab7309a75a6727205583a3a496e0d06d86544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaada16efd85cf53922dcbbed42c91e87094e4cbd52c8ca140bb3b9481845bbb"></a>

## where.virtual_network.ref — where.virtual_network.ref / 77bdf3f451e3 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-ca675ca3d165af7e4c1f190c7c8867d6931a560ae2cded7aad807536edad4f5f)
- where.virtual_network.ref

<a id="canonical-27151b05a603e25357144b806238e653d3815b8df21818787706e47f32ad8454"></a>

Type: `"list"`. Computed.

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

<a id="canonical-68741a9503e2b452c5660c855b05a5fae4b862e5a34eecdc6f9d39d0f1e6437a"></a>

## Direct properties — where.virtual_network.ref / 77bdf3f451e3 / 3

<a id="canonical-2689d38c905159e09a15142128ec8ffa92a081af78c83697121a9c0bb557b627"></a>

<a id="canonical-53f68885e270a289da1701b2c9815ed0e039284328389998cba1fbe8fc10239b"></a>

## kind property — where.virtual_network.ref / 77bdf3f451e3 / 4

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

<a id="canonical-88410dbb807e3d6d464aa8ebdf96c266c66e0e6148fddcb4e6e5a2b7ecfee376"></a>

<a id="canonical-3c1e3c4f839e58bbb3350e8da5b3605a07d7f1c8e0394a71a699cb838109d394"></a>

## name property — where.virtual_network.ref / 77bdf3f451e3 / 5

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

<a id="canonical-3b51e2fea22e18126aa4578e43447fd4ec5545db2efd98ce4e4ca61957c8c623"></a>

<a id="canonical-705e18005bda107a9f557b0629eb108acb6280473bc0618c67d5aad8db2b1b43"></a>

## namespace property — where.virtual_network.ref / 77bdf3f451e3 / 6

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

<a id="canonical-7304c75386c43601296fc1e74b4791bfd01a46a5258b347461791a1113f6d948"></a>

<a id="canonical-4a94c62f265cb9bf62ea1c2b9cf30e42676ffcb784e2d9953a45d8feb644ada6"></a>

## tenant property — where.virtual_network.ref / 77bdf3f451e3 / 7

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

<a id="canonical-253de281d1f2d2e380839305e22b2e17986c46bb870d226d7d5e986aa41aaee8"></a>

<a id="canonical-75040a1816d011d3f367ee741ea4241e48cffa599850b1f3a70afc7825c898de"></a>

## uid property — where.virtual_network.ref / 77bdf3f451e3 / 8

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

<a id="canonical-65ace576e83322ca2a431328a1c0d8db976d98a5c7a72804b0bbf278d44aabdf"></a>

## Next pages — where.virtual_network.ref / 77bdf3f451e3 / 9

- [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-ca675ca3d165af7e4c1f190c7c8867d6931a560ae2cded7aad807536edad4f5f)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-836a3a1e55671e48c018da1088d1269259d08ef45e164647cbe1b13da89cdf7e"></a>

## where.virtual_site — where.virtual_site / 66b13efe0064 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- where.virtual_site

<a id="canonical-934e4370fb5316cca4a1bd8f8914c503978126085eb34d6f7b5076fe0f3f06d0"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

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

<a id="canonical-dacd83ca20be2ca48f09726f5ab65ea5f5e96846190bacf927d4b122e35b9021"></a>

## Direct properties — where.virtual_site / 66b13efe0064 / 3

- [disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-74362324f0c96eafe1ce865b12cecf1afa710cca74fa37762b9369e208578ad2): complete subsection reference.

- [enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-fcaba6c260702f0bbde6e5f9e3d81261dbc61a8b086053c37d94f7fa11fa976c): complete subsection reference.

<a id="canonical-1a89e7a26be89261838161c23ede592c76901f62789ee108129fe5c670dddef5"></a>

<a id="canonical-7579bfe401d2f6b79bb0f0ad6447b658be72df84786dfa3fca94cb4a510bb5b9"></a>

## network_type property — where.virtual_site / 66b13efe0064 / 4

Type: `"string"`. Computed.

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

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-fffcbf7c67e3f36e0395325eacf7b235695255c218a39a5b049b111ac73b8df7): complete subsection reference.

<a id="canonical-a852e226089d33629865b3d9e75d38c3f52feca8150bdee7ab89c9556d8ddbe4"></a>

## Next pages — where.virtual_site / 66b13efe0064 / 5

- [where.virtual_site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-74362324f0c96eafe1ce865b12cecf1afa710cca74fa37762b9369e208578ad2)
- [where.virtual_site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-fcaba6c260702f0bbde6e5f9e3d81261dbc61a8b086053c37d94f7fa11fa976c)
- [where.virtual_site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-fffcbf7c67e3f36e0395325eacf7b235695255c218a39a5b049b111ac73b8df7)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-74362324f0c96eafe1ce865b12cecf1afa710cca74fa37762b9369e208578ad2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d40340fbdc97275f15c6796d2a41f685dc3afe1c172e3975b6659ad62aa521e"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 979ced083ee2 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- where.virtual_site.disable_internet_vip

<a id="canonical-4a37d034e26fdf8244c168cf0639c9207bd43910aa5782d262667ec119a9a665"></a>

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

<a id="canonical-22ec101aec3e63285c30a176b0f441796e5e4dd92ecf1b6f654577c582321b80"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 979ced083ee2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fac8dee7fa7098bdbce68210dd27154043d632a519a16bf83f8f1d231dc580b2"></a>

## Next pages — where.virtual_site.disable_internet_vip / 979ced083ee2 / 4

- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-fcaba6c260702f0bbde6e5f9e3d81261dbc61a8b086053c37d94f7fa11fa976c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d19b3340aa2657e04e9b38d4a9db2bbfdf4e78e22fc45d2344c18c9c91b62c34"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / 7a4f495875e7 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- where.virtual_site.enable_internet_vip

<a id="canonical-9230b30a51ff5b9de05b48bd0323b2c2eac94c15eb2af9fa8b887b482f0f6b51"></a>

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

<a id="canonical-0e04fe7fed17647b74dbc97f86e28c26a54de5e7dcab373727a4d2b78088ef1c"></a>

## Direct properties — where.virtual_site.enable_internet_vip / 7a4f495875e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe9b669f11276a2c9c69ec6cb992b8ca958a58af1288f44c6b76d29b27a7fc78"></a>

## Next pages — where.virtual_site.enable_internet_vip / 7a4f495875e7 / 4

- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-fffcbf7c67e3f36e0395325eacf7b235695255c218a39a5b049b111ac73b8df7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57d491f0f7e868c5fe6d5ab0a4894b089ee6650339e11f3bc6604ce5132b4b35"></a>

## where.virtual_site.ref — where.virtual_site.ref / 2f5288953fbc / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-cd636893f6ea17bde5bbcd57f3f1019b8bac3153b3656e94df442df23aafb56c)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- where.virtual_site.ref

<a id="canonical-56d53bd38c202376eae5e843e2eadee6191f51d685d1ce6e66ec28807c75f630"></a>

Type: `"list"`. Computed.

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

<a id="canonical-18c979e73b88c5c376cf6c5925aeb5524777d02dbd040f220a0fb41bf660864b"></a>

## Direct properties — where.virtual_site.ref / 2f5288953fbc / 3

<a id="canonical-9fa884cac5fd555f7203e6e7a36d93b9b976678c40f96dd9dac8d9f007fe2f20"></a>

<a id="canonical-a5c429a7b625a4fa9001d9a37c7197f535723bb3cb729aefa8fa80fb4b4f6917"></a>

## kind property — where.virtual_site.ref / 2f5288953fbc / 4

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

<a id="canonical-c419db7b9fc0bd4daf7125f47244fd1be891d77ebfca927f9d05e83495952b37"></a>

<a id="canonical-6a46c15b03db95e24f0a50f69b2b118011b02d4becc2de13405f227d2ba53c9a"></a>

## name property — where.virtual_site.ref / 2f5288953fbc / 5

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

<a id="canonical-cfb33a04497503311d6c6ea6107dd5d8036f1020c5e5bc0089109fd49cad1db8"></a>

<a id="canonical-ea8a6cc911449375af43ef19becf7fe9c2a89de3336a345b4a4908c45b5472df"></a>

## namespace property — where.virtual_site.ref / 2f5288953fbc / 6

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

<a id="canonical-433a4e12bbb82392962d91d6a29d288d10a43167df4b66ff0310cfecaa6e2d46"></a>

<a id="canonical-daeabd4053fd8231fe4ffa70a3d055c1b4fbfc3f190d46d9990cd3b13d6fbfe1"></a>

## tenant property — where.virtual_site.ref / 2f5288953fbc / 7

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

<a id="canonical-fd37c991f8dd804438ee3c386d1a6bd251ff87a2a9724a1e2e1a3253ce821097"></a>

<a id="canonical-06af932bfba5a1b7197485314d6b6e7f911616c1f1ffdfe9863111d43a3b4445"></a>

## uid property — where.virtual_site.ref / 2f5288953fbc / 8

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

<a id="canonical-1c3dd02059512d530dda33fa3f2a9df0e59bc90203bca64540cdd1ec4d11f00b"></a>

## Next pages — where.virtual_site.ref / 2f5288953fbc / 9

- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-5ac1128f4bbd53c0e06c86634d54eb0d9290eda425da2b11c9db2cc83c9625f4)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
