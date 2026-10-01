---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-461ca22f79505467bfdacf6280bf35f3445e377a61dccaba45ec3c871841a64d"></a>

## uid property — routes / 41209979738f / 8

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

<a id="canonical-f3681e298522e258423a4ed5aa15bdc58f56ad17b13597db626b09d369c02da8"></a>

## Next pages — routes / 41209979738f / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5570059452e7a42407d2b110e7733a91dd1e1471402407dd0b9bd2b2184c3e65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eabed2158a1781264607b8405ef05c7de836567b28405a066fe0d1cb5b6f6874"></a>

## sensitive_data_policy — sensitive_data_policy / b5b0973e949a / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- sensitive_data_policy

<a id="canonical-971a796701c3056e00203ca23b35d9ea84378d278b5024c02f0bf4a704cd4840"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

References to sensitive\_data\_policy objects.

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

<a id="canonical-e999b1ca909444c6c7dad321bdbca1d1c843b390d24819ac79891fd9fad62333"></a>

## Direct properties — sensitive_data_policy / b5b0973e949a / 3

<a id="canonical-5dbf72f59420fe0008d6be9d25abfeec0c60af0a64e7cce4ba2f579b49f0ad40"></a>

<a id="canonical-6f21908b29faedb68a13f37ffd5fe7d751bc9ec83e27886bff5537d8d8a5645a"></a>

## kind property — sensitive_data_policy / b5b0973e949a / 4

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

<a id="canonical-fb7b80c2084e04296b9476031b45d45d1257016e48038fe83a2956ee2ff70edd"></a>

<a id="canonical-af3ad2fc9480e91bcac50b0c0097967e029478fc3aae3561cb97f012ce3bbb05"></a>

## name property — sensitive_data_policy / b5b0973e949a / 5

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

<a id="canonical-6c45eff92d531b61f5edb5fa2df448def4235449878c4d8f30a163fa8025bc70"></a>

<a id="canonical-ea820578ca575a0bfffffbd5e406345c089fb8df7cee5c7887b5e55e004ce509"></a>

## namespace property — sensitive_data_policy / b5b0973e949a / 6

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

<a id="canonical-4b9a7fffafdd83bab5b7a77a97494f05785d7480cdcc6222aea2f49d5d33f6e5"></a>

<a id="canonical-ad36efc52d87a80819f8ee120bacc90c59536b3937653fa4560757778c78efbd"></a>

## tenant property — sensitive_data_policy / b5b0973e949a / 7

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

<a id="canonical-a688442abdee1715b7c57af8bbc89022588f56d7fcb64f984b7448c83cdbb706"></a>

<a id="canonical-496fbc5d0d7c97faea6e3c5180fdc3bbc3fba6336e5a3c04be0ccdaabdeb74df"></a>

## uid property — sensitive_data_policy / b5b0973e949a / 8

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

<a id="canonical-fe9a9b7c982681fae8155c9af42fad6b74d8e5a9b9a7a0cd6d449daf87f9b785"></a>

## Next pages — sensitive_data_policy / b5b0973e949a / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-88cb41571444c3a26fbfd8e039fa894468f209f6f9fb8b8397263ebd3c25f6e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1f5ded1bee7dc8d81cd3ed1d8382c8b87232a8c1db0b2bbc73f0598ee2e2ddf"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 9993a3a19bb6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- slow_ddos_mitigation

<a id="canonical-68f153655816b3c39b8aa77961cbd8d47fbeac2360184d417bbcdbbd10f9b2d0"></a>

Type: `"single"`. Computed.

'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

<a id="canonical-d8a6b1e47d3c21ffee1d3bfd48a34ce4c68fa2dd448fe72c4eeb708040a149fb"></a>

## Direct properties — slow_ddos_mitigation / 9993a3a19bb6 / 3

- [disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-bed5a0981222559bd129f2909e911b709b1d35a79b83aa9de28b631863d0df02): complete subsection reference.

<a id="canonical-a1280aa5af5593f1bd6d811e3b1e3b56922ba5d4b750cff40542466b24e0cbad"></a>

<a id="canonical-08c77802192ccff1743950cd9ad0f038b46c1be00c2b25ae90b1d4e17b6acda2"></a>

## request_headers_timeout property — slow_ddos_mitigation / 9993a3a19bb6 / 4

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-ecca1f5ac11fccb4eb6c52c6d18aceccd93f4783c179fbffc56453c135e70776"></a>

<a id="canonical-452b843aa9668a0dec11b190bfe740be4102596b5b66995e6ce5c8484af7f8a8"></a>

## request_timeout property — slow_ddos_mitigation / 9993a3a19bb6 / 5

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-55bc0515f9394e0727427335f11de5989f1477203d264f52df781386e41bdfb1"></a>

## Next pages — slow_ddos_mitigation / 9993a3a19bb6 / 6

- [slow_ddos_mitigation.disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-bed5a0981222559bd129f2909e911b709b1d35a79b83aa9de28b631863d0df02)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-bed5a0981222559bd129f2909e911b709b1d35a79b83aa9de28b631863d0df02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7386279390f71cf55d911f85e012c4567afeb54f20ef51304a1a64477b336e9"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / 86277c6ef297 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-88cb41571444c3a26fbfd8e039fa894468f209f6f9fb8b8397263ebd3c25f6e5)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-e25bd100b78830900ed44406d8e365fa994d0bb74482f537f6ec2f77bb1ef8da"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-d97f67f852072c98369fd59f3ec023852e2f358d3ab2a9221e38c69b55fc629c"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / 86277c6ef297 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5fd2600d9634de070f8a48542bbad23ea97a6689604a49c1383035f6e3d24f41"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / 86277c6ef297 / 4

- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-88cb41571444c3a26fbfd8e039fa894468f209f6f9fb8b8397263ebd3c25f6e5)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce8c8ba4d363a34d7b2cd67dc5f443fc6368cfb0cefd53afb1f25b8299d14c43"></a>

## tls_cert_params — tls_cert_params / b3f8b80cc3ff / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- tls_cert_params

<a id="canonical-ce341a9fe12b157fbb01cf6ce4594fc747e96462f75c5bc8f2287a8680c7daa4"></a>

Type: `"single"`. Computed.

\[OneOf: tls\_cert\_params, tls\_parameters\] Certificate Parameters for authentication, TLS
ciphers, and trust store.

Upstream description:

Certificate Parameters for authentication, TLS ciphers, and trust store.

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

OneOf alternatives in this subsection:

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-ce341a9fe12b157fbb01cf6ce4594fc747e96462f75c5bc8f2287a8680c7daa4)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-0732653dbaddc2627d7093bf3f40a2dab1f66fb00e4e404d4c12f855ad744991)

Select alternatives according to the provider validators above.

<a id="canonical-c4e01eff2023327071ae5caeb28fb4d5d47456036a546e44d09aa42e83e09cc6"></a>

## Direct properties — tls_cert_params / b3f8b80cc3ff / 3

- [certificates](data-sources--virtual_host--reference--group-003.md#canonical-bcf0ea18174d862cb2d52b51b5d1ce98d41993afe1a7d568316e43cf91b03166): complete subsection reference.

<a id="canonical-f01eb6c194893e186067c2c7fe05e275286ea79ea3ced0473d41eef07de3a08f"></a>

<a id="canonical-1fc9fbe40b55177cd902d93b67729111c528fef61a04329d716064e10b0cd711"></a>

## cipher_suites property — tls_cert_params / b3f8b80cc3ff / 4

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

- [client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-9a05706a9ab0d94b2b9b94d4c0b762c783a8c1dc8713d832663d278288c177b6): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-01cf4ad55d872b9f14e0800a711810c120b8610663082ef60641a48de8355094): complete subsection reference.

<a id="canonical-5d53c1bda99569b9d4df1a6dbc8e27b0fbc08dbf9a3c4ac3d6a7849cbb35ca25"></a>

<a id="canonical-1208ecbcb0f79b2229a99d0c86e7b46fcff1f10092c4cb478cbc8c85739b290f"></a>

## maximum_protocol_version property — tls_cert_params / b3f8b80cc3ff / 5

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

<a id="canonical-11303134574ecc97ec5cd0995012082e83206e722b9466e24199120ee1a920c1"></a>

<a id="canonical-67e4976c35389d34c07732e2595e985426bf89d38318bc0ecf1133d0382ea634"></a>

## minimum_protocol_version property — tls_cert_params / b3f8b80cc3ff / 6

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

- [no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-3801a65b8630eb60d82f9e548f363d3edfb65d39a9c5f7ae6d4fbf5688c3db63): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-003.md#canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef): complete subsection reference.

<a id="canonical-f4c81200f6e84bf480e43b51fa87012ab0e583a6c6ce8949372a3ffb0126426d"></a>

<a id="canonical-ed9c29241ed63b877e1792645c744c77d06bde8c5b044914fd7be8173dffc51f"></a>

## xfcc_header_elements property — tls_cert_params / b3f8b80cc3ff / 7

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

<a id="canonical-17eda5f5e69a6d040e524a938d6e767e54a5b63ac441e43c4c12334ca3127b1d"></a>

## Next pages — tls_cert_params / b3f8b80cc3ff / 8

- [tls_cert_params.certificates](data-sources--virtual_host--reference--group-003.md#canonical-bcf0ea18174d862cb2d52b51b5d1ce98d41993afe1a7d568316e43cf91b03166)
- [tls_cert_params.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-9a05706a9ab0d94b2b9b94d4c0b762c783a8c1dc8713d832663d278288c177b6)
- [tls_cert_params.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-01cf4ad55d872b9f14e0800a711810c120b8610663082ef60641a48de8355094)
- [tls_cert_params.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-3801a65b8630eb60d82f9e548f363d3edfb65d39a9c5f7ae6d4fbf5688c3db63)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-bcf0ea18174d862cb2d52b51b5d1ce98d41993afe1a7d568316e43cf91b03166"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ea81f230b4dd494fddd7c9bbcb76b232f21101b1a61b176086fbca76fa99215"></a>

## tls_cert_params.certificates — tls_cert_params.certificates / 26befb528512 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- tls_cert_params.certificates

<a id="canonical-e9bced0e6666f9101a112e5d0f573ce569b961b131447e930552ca5527efdf1e"></a>

Type: `"list"`. Computed.

Certificates. Set of certificates.

Upstream description:

Set of certificates.

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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d4c23cc115efe52c54ff690df15de5639805a8c0b3abac7d7caa81bbe2756ba3"></a>

## Direct properties — tls_cert_params.certificates / 26befb528512 / 3

<a id="canonical-bddbabc1d90178c5af129f75ba0a98699412ee43afc3b279f2657e4e7f9c5770"></a>

<a id="canonical-49b5b0daff0f9fa6554593acd971148af866ca62c0ee09c61df12a1921b8705b"></a>

## kind property — tls_cert_params.certificates / 26befb528512 / 4

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

<a id="canonical-057e6ec799c269b85c842358b178bde456d97e5c5ef0efb747aabed5dbe2d5ab"></a>

<a id="canonical-2485646e47286e0110d408a297b5c60a0ece38bc3b312c29a7b649107e29d524"></a>

## name property — tls_cert_params.certificates / 26befb528512 / 5

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

<a id="canonical-84df9e8497e0537f22b3ee7036f1a7db43893f33c8d658274c736486cc5f38fb"></a>

<a id="canonical-d8ed48859f24b661182f4dbf48373d31e0ae9168ae623f3e411ec017427fdd33"></a>

## namespace property — tls_cert_params.certificates / 26befb528512 / 6

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

<a id="canonical-b2ee7bda0217ee35c0ed7371decbc76ab5be4d187fecc593d70464bfc30e247e"></a>

<a id="canonical-9b864d3487d9ac592be9a813cf07330c4c296635aefb898e0dba2638014976d0"></a>

## tenant property — tls_cert_params.certificates / 26befb528512 / 7

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

<a id="canonical-f567841f5d2cbf50768f705ed18625cadbb1beeac5d71827632a5964c39de1ab"></a>

<a id="canonical-f1e385f83344456ae215773845ecb177f1b99b52d77829ccd7f57017f8f0f655"></a>

## uid property — tls_cert_params.certificates / 26befb528512 / 8

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

<a id="canonical-9626d3fa0b42f01d33bdb95551ce4618b1338f5d4e8d9c7e8f34cdcfad2e3a44"></a>

## Next pages — tls_cert_params.certificates / 26befb528512 / 9

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-9a05706a9ab0d94b2b9b94d4c0b762c783a8c1dc8713d832663d278288c177b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d7fee42fdbb863f2f07a775a436ab37cacccab324727919d3a148976c3fc807"></a>

## tls_cert_params.client_certificate_optional — tls_cert_params.client_certificate_optional / 76dd082922a4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- tls_cert_params.client_certificate_optional

<a id="canonical-09220efd5109ca70e867a121cbbf3c0a07ae02dd29409e854f07d2b99d550fca"></a>

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

<a id="canonical-fcb47b4f0126fbe122d518dd1be84447dcfce40e70c290b22929d72d95e59357"></a>

## Direct properties — tls_cert_params.client_certificate_optional / 76dd082922a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36edc8a204a670976bf6ac709234f7db27b781c26523e242e89bb365a6519e35"></a>

## Next pages — tls_cert_params.client_certificate_optional / 76dd082922a4 / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-01cf4ad55d872b9f14e0800a711810c120b8610663082ef60641a48de8355094"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a3c7da61169ffecb3897f74b2299eb365388cf08c44eadc46bf27d8c2ad98d6"></a>

## tls_cert_params.client_certificate_required — tls_cert_params.client_certificate_required / 3fa26b4d0640 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- tls_cert_params.client_certificate_required

<a id="canonical-74f54396c29bfa3e059f074e601cb1e8e626bdf1326ad1517921371987b972ef"></a>

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

<a id="canonical-a75c628a317d0aae346f6ced302a08313498fafd3789936f3163964005a1be53"></a>

## Direct properties — tls_cert_params.client_certificate_required / 3fa26b4d0640 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-623e0323b3060b7b6abbbb3f268128bf848fa37c5fd91d44eed5fe135943d17f"></a>

## Next pages — tls_cert_params.client_certificate_required / 3fa26b4d0640 / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-3801a65b8630eb60d82f9e548f363d3edfb65d39a9c5f7ae6d4fbf5688c3db63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ff4dd19389f7d95354db361fad0e2276aa528334ecd9b798c6c340aaba2473e"></a>

## tls_cert_params.no_client_certificate — tls_cert_params.no_client_certificate / af3543fce20b / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- tls_cert_params.no_client_certificate

<a id="canonical-24556e8bb2bfac0fa5850456f26f05a60ad35a525662b21f6eaac9db86589fdf"></a>

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

<a id="canonical-e21d6ba2ed65f03d7f5628fe0d9dffd6555454ab969cc30e08057fa9534d13cc"></a>

## Direct properties — tls_cert_params.no_client_certificate / af3543fce20b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9dad476960c21ef7ace43663b28c67adf55e861e759b4fa8ac965df7a59d96c"></a>

## Next pages — tls_cert_params.no_client_certificate / af3543fce20b / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-959bd5fe2f934673beb97beb692bd354b747d264d9b5aa9aae7428b575c7ab23"></a>

## tls_cert_params.validation_params — tls_cert_params.validation_params / db0bc8e577ae / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- tls_cert_params.validation_params

<a id="canonical-b0caac7847a38a1a85587cbb0319f454a5aa20fac3b51ff179a6ebbc7c73545b"></a>

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

<a id="canonical-24d0ae804a4f00f9d29a9a69d4a4f744165c57cd8b81dcba565874030bb5b102"></a>

## Direct properties — tls_cert_params.validation_params / db0bc8e577ae / 3

<a id="canonical-f86d8236bcb22b09ea359ad60fb742c61947f15deb97deb7c3fe9c18688a9476"></a>

<a id="canonical-9129dc0ac839bd043e7bd4b9720d616fb28294b385a7abd14b1fdee8cbe571f7"></a>

## skip_hostname_verification property — tls_cert_params.validation_params / db0bc8e577ae / 4

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

- [trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-4bd4bbafe8450589428c3f7cff673ce8848a238dbd15a3e96e02fb2fa9b07dab): complete subsection reference.

<a id="canonical-d9563e38fceedcc5cd63427e960825ad0c365799f62e18e62f077d5d05d744c6"></a>

<a id="canonical-580d8dcf59d2331f8a80e6685b3c0af14d725261b5a7b8f6e544d6c36f0c9709"></a>

## trusted_ca_url property — tls_cert_params.validation_params / db0bc8e577ae / 5

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

<a id="canonical-22140d719d51b247de5b0fb3999008a269de993beb6c00379a1b97e1cce5cde6"></a>

<a id="canonical-d67ab53c4d5f56aa0f54436e27023f376af85923abe977866f55b68a3391e6e1"></a>

## verify_subject_alt_names property — tls_cert_params.validation_params / db0bc8e577ae / 6

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

<a id="canonical-dde2d89b60cd7a7442f1f4c122cf5c28b83ca8050c3e4d608f75883dcfae8e8f"></a>

## Next pages — tls_cert_params.validation_params / db0bc8e577ae / 7

- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-4bd4bbafe8450589428c3f7cff673ce8848a238dbd15a3e96e02fb2fa9b07dab)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4bd4bbafe8450589428c3f7cff673ce8848a238dbd15a3e96e02fb2fa9b07dab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-752cf8e73bf6ebe8b8216e2a397a1ed48e566789cc011190ab886086567eb74f"></a>

## tls_cert_params.validation_params.trusted_ca — tls_cert_params.validation_params.trusted_ca / 22c4f2a74fbe / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-98d9f9782710c8dba62e79fef1708cbeb04118ce820579e058be56441a5bb196"></a>

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

<a id="canonical-e0917141cda394b9aad18c433b655ee7bb708ec8cdbf174fb2f601f97002a93b"></a>

## Direct properties — tls_cert_params.validation_params.trusted_ca / 22c4f2a74fbe / 3

- [trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-db3c1eb5dcd3238bc9220a83c614b6d8e41dcb5ba0db5719e6e2cc7dd4a86212): complete subsection reference.

<a id="canonical-229108b9da1fe5bc5c6c177c11b807f617b6f26bcdfcbb882ed8b4bf182f0cdd"></a>

## Next pages — tls_cert_params.validation_params.trusted_ca / 22c4f2a74fbe / 4

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-db3c1eb5dcd3238bc9220a83c614b6d8e41dcb5ba0db5719e6e2cc7dd4a86212)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-db3c1eb5dcd3238bc9220a83c614b6d8e41dcb5ba0db5719e6e2cc7dd4a86212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4b682486aede6de80b94b6a6e4bda410d41cd0a001965585b37ef3f7d69ae55"></a>

## tls_cert_params.validation_params.trusted_ca.trusted_ca_list — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-e70fab2a75f6ec1f2a89bff1c08fac9ca410debbcb93fcf0ff83160be60ca8ef)
- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-4bd4bbafe8450589428c3f7cff673ce8848a238dbd15a3e96e02fb2fa9b07dab)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-ba86bf454a89d1f506dccf57615105110045d5ccbf5a8d09c499a26800c29c71"></a>

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

<a id="canonical-e868acaf5d93796843df8088b909aeada22438144dd5672d5505d3e967a1e240"></a>

## Direct properties — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 3

<a id="canonical-ba7e1ae225c0883fd61397314bce4f8ddc8dd2a5df6e0ee2710ed0e593517e30"></a>

<a id="canonical-806a9fc39c3c972027fd6d802d76ced61a865d13e6fc434132d70b33c15f1bc0"></a>

## kind property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 4

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

<a id="canonical-2355f29906826f79ff26c65b860c4c7e4f1f192722f21c0a2dc93d29686124a8"></a>

<a id="canonical-94d0274ab861e87f8907c9c447a9f26ef67ca7a08e7a37a17d2d089758b1cd96"></a>

## name property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 5

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

<a id="canonical-0843e82997e48b07187255008aacdfa0d37ad74fb1f1b95f44c7e13048df583e"></a>

<a id="canonical-715342e7719db5307173d83d89689e08a68132d5db827e8e2f7a412f10a5ba40"></a>

## namespace property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 6

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

<a id="canonical-e4369d49371818866851a9d1b94abb972b38025f42c460c8f5b934330a12cf90"></a>

<a id="canonical-50154960abb543087f994cc21a6f7c715eea15041df301fbc420639dbb9de8f2"></a>

## tenant property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 7

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

<a id="canonical-f1215088f8189e7d1d1c5d9042a28ea72675b40f3807cae56c0ac85360bc4c10"></a>

<a id="canonical-788199e7ec0f82fb00cd7fde1a378ff5179992f7fdfb52e5c8641201ed5f3a2d"></a>

## uid property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 8

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

<a id="canonical-d077f727a0e25de517348db5e03499dc0d33320e7c112093d2bba312017e6ffe"></a>

## Next pages — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / cf4b5ab82bbb / 9

- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-4bd4bbafe8450589428c3f7cff673ce8848a238dbd15a3e96e02fb2fa9b07dab)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a3ae3ecd2c0ae66434bbc55d01bb1d8c8b9ea1eaebbec110c464b9f9b580320"></a>

## tls_parameters — tls_parameters / f2559f4bbf77 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- tls_parameters

<a id="canonical-0732653dbaddc2627d7093bf3f40a2dab1f66fb00e4e404d4c12f855ad744991"></a>

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

<a id="canonical-bf495eba3a563fe9624bff1a67349fa3dc0462c778548f4e8b5003301e19ffd5"></a>

## Direct properties — tls_parameters / f2559f4bbf77 / 3

- [client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-359dd9bdcf471c1c36b398664c672b788d4ea9519cb2ad547d9aeb8eee475058): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-d0d784ed1cdc3c269510b7af4214d9205b763d64c7f173a76449a9ffe1f894b6): complete subsection reference.

- [common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415): complete subsection reference.

- [no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-03ae3865667587073d9d3b0db40499708a6196253d0369c3ab4df10581ccacf7): complete subsection reference.

<a id="canonical-e0f94dab45d6986874c32e92e76bdfbc7cb82cff042f5a54e90c98d14cf3311a"></a>

<a id="canonical-ff1a6b87c328b4aaf1d2db3d83139ffbb085c2a55c2c63427c9129db4c640d1a"></a>

## xfcc_header_elements property — tls_parameters / f2559f4bbf77 / 4

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

<a id="canonical-cbccd3bcd54d0078138d530075308665c89ab9b2aadcff8e9fe4b2b1bc5a3d63"></a>

## Next pages — tls_parameters / f2559f4bbf77 / 5

- [tls_parameters.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-359dd9bdcf471c1c36b398664c672b788d4ea9519cb2ad547d9aeb8eee475058)
- [tls_parameters.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-d0d784ed1cdc3c269510b7af4214d9205b763d64c7f173a76449a9ffe1f894b6)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-03ae3865667587073d9d3b0db40499708a6196253d0369c3ab4df10581ccacf7)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-359dd9bdcf471c1c36b398664c672b788d4ea9519cb2ad547d9aeb8eee475058"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd0a2c5638be7a663c13bcab308b51f3da07499e47ebda3dd0d38c430f76a5d4"></a>

## tls_parameters.client_certificate_optional — tls_parameters.client_certificate_optional / 08e13b15ec99 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- tls_parameters.client_certificate_optional

<a id="canonical-e999250ea6fc4e9c54c4d9562fc16d0e685f91c0ebe654c249b098e5d0d45994"></a>

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

<a id="canonical-c00bb5e6e6f234e180e6a2ed207b8d7168b93045d83b4f622354a9d30c1a47f6"></a>

## Direct properties — tls_parameters.client_certificate_optional / 08e13b15ec99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ebc67bd2ba2de1d3880f0509eccc829d440ca86f4e24639c2092e2969c4bb4b"></a>

## Next pages — tls_parameters.client_certificate_optional / 08e13b15ec99 / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-d0d784ed1cdc3c269510b7af4214d9205b763d64c7f173a76449a9ffe1f894b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-868082f7d7fab8d25984522f9e9aec18380fa96217896e136afeb79e740e62db"></a>

## tls_parameters.client_certificate_required — tls_parameters.client_certificate_required / ab1571efdbbe / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- tls_parameters.client_certificate_required

<a id="canonical-7a64a241651dd0dfaf6aef17a3667c03defafc35a50df93736203606880b338f"></a>

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

<a id="canonical-f6a391c26c9643c0f3b6df081e0cd5adf1314988b0ef9e25f963ff4daa9d663e"></a>

## Direct properties — tls_parameters.client_certificate_required / ab1571efdbbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75db33574d864d15647ec42be5be6bbd8709175a73a9784d420d585752a06e80"></a>

## Next pages — tls_parameters.client_certificate_required / ab1571efdbbe / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4ca8f19d6bbb66c28250c09c14f53d6294e9f8b48fd0bb4f8159e3febe59346"></a>

## tls_parameters.common_params — tls_parameters.common_params / 874417648e95 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- tls_parameters.common_params

<a id="canonical-5abc749c896aca2bc2343b0cc8e7002a7a1225039ab72ee0a3a32dddbbd4180e"></a>

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

<a id="canonical-1ab7768277ff5c326b45b90d4550f87f7e04ea6823372d76f0210c8af6d5d611"></a>

## Direct properties — tls_parameters.common_params / 874417648e95 / 3

<a id="canonical-244de8fa402dab3c63864993b1419b1729ca0a57a6990fb3095bd059f4c79ce5"></a>

<a id="canonical-2e79359dc5e302d9a271867c00eed35d7820c0a5e60e68cf38ff0b1fadbdef73"></a>

## cipher_suites property — tls_parameters.common_params / 874417648e95 / 4

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

<a id="canonical-f50844da477b463c357255916fffb2e78bc30a539b92fc387d950ecb1be64526"></a>

<a id="canonical-1167f6a6b7b255e4ee68a154b8304451a73670ed6f9999ade7cb01365698c206"></a>

## maximum_protocol_version property — tls_parameters.common_params / 874417648e95 / 5

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

<a id="canonical-845e8423da4825cf72c81ccc638fd6af5ad5f01f4fd40401321595c03afc31ac"></a>

<a id="canonical-c91b6f05714207fb5a903577c45a7380bdcfc64b42f51958d2d4d4eb436dc9d4"></a>

## minimum_protocol_version property — tls_parameters.common_params / 874417648e95 / 6

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

- [tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-003.md#canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c): complete subsection reference.

<a id="canonical-42158830b13553710938bdb7d51cda816d208f02f8d9c9b7c661d92f8141dfcb"></a>

## Next pages — tls_parameters.common_params / 874417648e95 / 7

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9abf108db6c8fb9f37cdd9c0f9aeb73f8ba0fb3adbff2d52a36aef265ec9fa3"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / ead86deb8340 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- tls_parameters.common_params.tls_certificates

<a id="canonical-da96377d543ffd34477bb3798b91b095b9eb8017f9eb7b01ba2499b2390097c1"></a>

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

<a id="canonical-6296b1b62088f85730f4c36741e20d4e8d22eae492fdd574fecf8c484c11b1dc"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / ead86deb8340 / 3

<a id="canonical-1c760bddd207ebd1d5d41d9709fffbfbe17f13a8ddc6c80667c44949f06e5b64"></a>

<a id="canonical-7975a2934b6ad1115bda5a9bf24371ebcad3d406f3856cce166cd7446dc0ea6c"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / ead86deb8340 / 4

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

- [custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-87a7cc7120645a081d40c7ded5a10c873bc471020e2e768a57328f53f9ab23fe): complete subsection reference.

<a id="canonical-569e9f474f7c9152ec16013f5ca1c06161cff65d143891569e59f048296e4318"></a>

<a id="canonical-58e963c708e6286862671dd036af5a503ae77f89f6124defefacff12b7867646"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / ead86deb8340 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-64ac2c11541473bd76cd8468ba32e314c5f0500cca030d47fc698c0834853b66): complete subsection reference.

- [private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a): complete subsection reference.

- [use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-8e009bc53fd743493e728d85faab901b7ae9a7875dd34aa13b2d7a66ff82761c): complete subsection reference.

<a id="canonical-486370b79deea41849da42917e1612a8967b22ea92870299a500b444a7fb4616"></a>

## Next pages — tls_parameters.common_params.tls_certificates / ead86deb8340 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-87a7cc7120645a081d40c7ded5a10c873bc471020e2e768a57328f53f9ab23fe)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-64ac2c11541473bd76cd8468ba32e314c5f0500cca030d47fc698c0834853b66)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-8e009bc53fd743493e728d85faab901b7ae9a7875dd34aa13b2d7a66ff82761c)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-87a7cc7120645a081d40c7ded5a10c873bc471020e2e768a57328f53f9ab23fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45de3a51a9c2829e3a1422e3b45354ee8cee250ff6ad1fb6b910710dbd08d788"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 5ac8229a78d0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-47435d5adffad679050ca1070a72037a479d4f50728ce7e39e0c8283a3a98ae5"></a>

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

<a id="canonical-7a1aacab8f69209f8dbbffe6ffd806feb6e989b6843aa8e252041bc6df673bc3"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 5ac8229a78d0 / 3

<a id="canonical-3fef8b5dd6b64e4d518e814e4c5e43df0e2972e5f251f8734dfb008169c030b0"></a>

<a id="canonical-80c45327c7dec97eaa858f85350868c2e715f32c0acf8451fe28bfe81f617629"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 5ac8229a78d0 / 4

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

<a id="canonical-c94bbdf1ba8d88cef59806dac1ce787bba11b3d39e7e57da369e1d70d8b65385"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 5ac8229a78d0 / 5

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-64ac2c11541473bd76cd8468ba32e314c5f0500cca030d47fc698c0834853b66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bedfedd1a5445ad7002fb9f40a0fd95faaece807d7de7586cdb3e87bfcb6dc13"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 22f8037f1d47 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-4e008c4e680a333ac574d1ad89fc42ce7f86908298613b2aca2e1ea3a2a690b4"></a>

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

<a id="canonical-9847fe267fabbc61743ea0c30743e269c9913b857ea1cc593e62fda746764e61"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 22f8037f1d47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cb8d37b9afc29bb2543c920a85586a317b7176ecd321699e87db9f3876eaf4e"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 22f8037f1d47 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ed4f72317d71c7e9e51dd2123e9e2e32351ac9a419a02cc65e25583978e5c34"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / 49b77ff6d453 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-1e9fe1f0fc7938955b696183a8c9e1918ed5a6a5a62e0e8eb42e7dfff8c3af34"></a>

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

<a id="canonical-4e166b64b88b74d1d17b652812ebc03be0635690721a166e76a2c6aa9b3cdc5f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / 49b77ff6d453 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-26b5ddb1c410275aff0844e3058d8dc61f42cc9acf421e697db8422e81d15b5c): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-4cbad69cf661184fb611f67cf7060b38887432a06844e78037d8c7a3d1eb30fb): complete subsection reference.

<a id="canonical-f2d9c816c6741f0cf9b02cfd30663667b957c88bda46229a233b7e700a71cf4d"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / 49b77ff6d453 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-26b5ddb1c410275aff0844e3058d8dc61f42cc9acf421e697db8422e81d15b5c)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-4cbad69cf661184fb611f67cf7060b38887432a06844e78037d8c7a3d1eb30fb)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-26b5ddb1c410275aff0844e3058d8dc61f42cc9acf421e697db8422e81d15b5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70bfb9a07fe8bca83a8e5a2cdc1cb502e5aa3659f21487a80ce90a553cbac926"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-5aa778016f56a7bb92a2a3c9f1a175eeab16884c0849063a3d075af3f3892b9f"></a>

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

<a id="canonical-b8aefce300b7b289f03f11f636f1de6a8a2da9cedb2e03feb12ab9ba6f93908f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 3

<a id="canonical-bc193035f1e242174c71222d70e893084312e98342fed174b5b47e857a13d6b4"></a>

<a id="canonical-9c7392d20413d2ae5c1fbf41b7bb130fee2bbea167a653935f875e88dcf1de9e"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 4

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

<a id="canonical-1bce9aec142dcb9f5adca8b979b33b5f9d1cf90e50c71a5957c05ea75d2f629b"></a>

<a id="canonical-16691ff7b3b4b30b6ba230eb5e765127fba6fefc3f8272b28e6f40cef5f60e05"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 5

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

<a id="canonical-ad935fb7562b59702096e1d4773be02965f6c50af6d62614ca8b443af787226c"></a>

<a id="canonical-921f20b9b2c4b20a2a2dff74878df459cfe56c253ab22640acefce7511420a51"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 6

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

<a id="canonical-76295ecd1699a3b82a52e6d7ed37ce68017a7e6ae259a11c443156d25aa3237b"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / cc56165333cd / 7

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4cbad69cf661184fb611f67cf7060b38887432a06844e78037d8c7a3d1eb30fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6049d5812a31c6ee936ae1ac608a817583101eb1ec2b8ff3d05fb34173ab0a42"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / ce7ede0d9d92 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-12b21048f3cba1abea5ec4182bb46b60338d94455dd3b24ecc3ff150ff00e002"></a>

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

<a id="canonical-6284300396bd9428939a4676fe03def68566900637144a727ec1d9952e02a5ab"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / ce7ede0d9d92 / 3

<a id="canonical-fe6becafb08a33b7176b9a069a60ea36939a6aa1816d0bcd6ea88e1e3dc13991"></a>

<a id="canonical-0f622432d210ad4278a7c278ff43f8e96d62a479832d4c772858c465d3ab2c5d"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / ce7ede0d9d92 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3bc9433f4766825b7165736fdd227b89d2d6f95cb596181e50378085e9fa2847"></a>

<a id="canonical-b9070073b4920452e875a6a7d9908ac8f65055fefa978c004b82072d62a43bd0"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / ce7ede0d9d92 / 5

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

<a id="canonical-aeb162f8bd086f50a55de03313ced168ef71fa0c3d2fd1054008688d3887deb0"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / ce7ede0d9d92 / 6

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0ed5dabdeb9e90a8dcceed9f37e453a0915582c69c11446035c29179d361946a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-8e009bc53fd743493e728d85faab901b7ae9a7875dd34aa13b2d7a66ff82761c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bdc4dde701b8059aa3194bade387b0c889af3ac95122e094fafc183e1ea54fb"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / 5907840717b3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-94cf26757c1211fd5fe547448f667a26c6089f3e1fd46b4aaf89dba4e3b3e9ba"></a>

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

<a id="canonical-3b23e953c16093b2d29cb3e00a6ebfeb70ebc4e0ab38cd447a418b99eaa4d958"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / 5907840717b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-486e7ecef61598a7465d3557cc5785bf54da1ca4a3b0cc54dbd922da46c070dc"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / 5907840717b3 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-5036d1e27a04b4c1efc928809662229744f132ef9d87c572848e6be5836feff6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c7e2d5b94ce8327a2257ee81489bbd8fda9c65bc4f4af8ab9fb1fe7902562ad"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- tls_parameters.common_params.validation_params

<a id="canonical-a61a6478d604eedb8f8444249f4c3cf7cf283966277563a3aff654948e3d80f4"></a>

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

<a id="canonical-879081bf4a5efa9a7deb639e68d44041eb2fe84ba28f4eb0702ffac57fdbe7b2"></a>

## Direct properties — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 3

<a id="canonical-8c5f8be7287f0945daa7d32d468870a460ac9cd2d39d96503b6e7f71226e0e54"></a>

<a id="canonical-6a26b5b63bf387b98b922503f5b4e0a8ed40593af560e092268bb54303c7b01d"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 4

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

- [trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-fff7fd5c1d1833a501310719e589365f8fbf12d8a17dde52c6fe62297a4ba686): complete subsection reference.

<a id="canonical-a989b1cf31aa6aac024fa1a56b910f1afe94b77430aee89c74d63b70b185eae4"></a>

<a id="canonical-3101a2ab5bdf748f2088bfd764ddef2e1256ebf88495aac37986349318823521"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 5

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

<a id="canonical-8cf871809f50f833b3396fa80d2eb7b18df4781ed9d5701cd77c2ec2ec1e88a0"></a>

<a id="canonical-556e1a9757037578e32d01f690b9731d3a31b35756a59141c8b1f5b33adecb3e"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 6

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

<a id="canonical-a5a89c73656b98bb92ef6369191bcdb4bc3285b3dceb574a27d2ac68336cc842"></a>

## Next pages — tls_parameters.common_params.validation_params / e5b56d35c9c6 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-fff7fd5c1d1833a501310719e589365f8fbf12d8a17dde52c6fe62297a4ba686)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-fff7fd5c1d1833a501310719e589365f8fbf12d8a17dde52c6fe62297a4ba686"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49cf2597399dd69d3d6b2eb4529fd6b26ec004d0fc072393e9b8af98f88fb1a8"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / 481c18b67345 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-14abafaaf994cf232a1e931cadbb6ed2633b74e233a3653c3c6ac8322c3d6a9a"></a>

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

<a id="canonical-2777f393aebd84ef9d50e9348556591686f5db6a0a4bda2d106c89f55e7698c5"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / 481c18b67345 / 3

- [trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-a66ebc1ac54341786958578f517eac0325e7209a5971993a3c024ed1f97beeb1): complete subsection reference.

<a id="canonical-57787de8f1318067d86b52171bcd4b1109d4fcad2e217c1843aa2834f5701f84"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / 481c18b67345 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-a66ebc1ac54341786958578f517eac0325e7209a5971993a3c024ed1f97beeb1)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-a66ebc1ac54341786958578f517eac0325e7209a5971993a3c024ed1f97beeb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdfca63f2bdf093a6c6aa286f9d0c1afddb14bc79fbee4d292d5906529a1a1f7"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-86047c93ba6e4d261f8d0bfd4176629190c2e6e9a56481d059f887f5c75e1415)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-4581c07d78c68681f4c451bf715e256b04a5008d63e428dd54cf1100925b276c)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-fff7fd5c1d1833a501310719e589365f8fbf12d8a17dde52c6fe62297a4ba686)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-4f008fa20b87948309e8790c19d1485520a1e1d7e500fe18541d8ab07942bb47"></a>

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

<a id="canonical-2a85ce726508b59e0e754159345b2e4d15ee2166ee847655d3d0ff075e8be09a"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 3

<a id="canonical-e74a68c645ec9ef69fee63a65468b890e275661be8b31f39cfd9eb33df0a6078"></a>

<a id="canonical-e584c5597ff6faaeee08109468c61ddc5bfe0bb6fd22fda92b0d631e01e4cd83"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 4

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

<a id="canonical-f9334ddc1b34097cc5748915163ac21c179b00ef6320935a518b3569ca6bad79"></a>

<a id="canonical-44818d4e57766dac5ff6aea0ce3203388bbf4a247f44d5696c2e61a1b12a941b"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 5

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

<a id="canonical-fd6f36998d075f73ae18de5b4645ed02e182216f61553c6a0958cce07fe341e6"></a>

<a id="canonical-ad703ab415cdf8925851d92f796f50cb129e1c98827e9757510e1dd2bf9dd767"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 6

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

<a id="canonical-986ba2b6bc1abba4391175b663f694aa293880f11e31fe816e00be1aa2d81702"></a>

<a id="canonical-fb64d8d745b074db600e9ee416858c0ce52946fd35f07b8267b5eb0a98e262cb"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 7

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

<a id="canonical-7a717aafe94881c2b5c5dbaaad166d169f2939c80aadc63948ac4d38273ff81b"></a>

<a id="canonical-d4f15e8b91dc90ae1f8b871cf8a9c6e0a5bd8ac355cc758e702a4f11a324ad5f"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 8

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

<a id="canonical-27c2edc4af44cbc37c48fd5c7d3fc632f5543b62ed5de69b0fc24b4ec3780f85"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / a80286c228ee / 9

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-fff7fd5c1d1833a501310719e589365f8fbf12d8a17dde52c6fe62297a4ba686)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-03ae3865667587073d9d3b0db40499708a6196253d0369c3ab4df10581ccacf7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b7e34ac7f849be36e229d63f1bd84c99c3b980545e46b72968091991194ef83"></a>

## tls_parameters.no_client_certificate — tls_parameters.no_client_certificate / 8dfa1de6df4b / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- tls_parameters.no_client_certificate

<a id="canonical-8367fcf195d958ca78e1cbf5cf0281bf212b1d0f766f063ce1b6d564e7181160"></a>

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

<a id="canonical-9c43e8355dba38217c5833acb45f15712a4a0622bf5a307dcf24322aa0aa8707"></a>

## Direct properties — tls_parameters.no_client_certificate / 8dfa1de6df4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8be4a4195913d06abe30495937fc5d7d0ecefa8a3a8ee2372554b583c53736cc"></a>

## Next pages — tls_parameters.no_client_certificate / 8dfa1de6df4b / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-42caf6807b285559de11020a7412783936bce26b48dc1c3b083cd332528d2cb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d5524c0a35bf0293a3d562661a47a7eb4cca366857586375ea26b0c1709815a"></a>

## user_identification — user_identification / 11dc42dedbf3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- user_identification

<a id="canonical-fed06eb3b1e4290fe0d4c0a7d0f9d0b0c96a25617fadf6321587fcb3903c8054"></a>

Type: `"list"`. Computed.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited.

Upstream description:

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

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

<a id="canonical-a1f1d58dcd3dc023d12a53e60674501e6e3ffb1f7a9549da6fe66a0c0e72ceb9"></a>

## Direct properties — user_identification / 11dc42dedbf3 / 3

<a id="canonical-653bfa9c573b55bd7a58b82fa9912c8e157a2b2f37d515d9ad591507946273b0"></a>

<a id="canonical-4bbd50c3bf7bdb991200ebc07e3ff90f4abe9275d7082c307b2d0b6118ce6381"></a>

## kind property — user_identification / 11dc42dedbf3 / 4

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

<a id="canonical-722c06a06f8279863408879de60e03f5538c29a906b28e77114a0bab9547b473"></a>

<a id="canonical-d5fbb56c6e23cb029bf1951c803b2d7d77a4802a252a2cc91fe6d24be5109340"></a>

## name property — user_identification / 11dc42dedbf3 / 5

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

<a id="canonical-6adfb7d5868d541895c97a73dca6d9c60f0067483160410479f0cb1f28ff825a"></a>

<a id="canonical-01b28131749be7335b7fb254c3eca2d12df76a2f6ce0f4f56b39e141068dcec4"></a>

## namespace property — user_identification / 11dc42dedbf3 / 6

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

<a id="canonical-5d35978ae0af9718d1ec5e98682cffefcc75397501b163072365cd49b5c32948"></a>

<a id="canonical-6472f1aa2a614562098f56b1f231a62462d14a4b8a7f9c2141439a6c87a104d0"></a>

## tenant property — user_identification / 11dc42dedbf3 / 7

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

<a id="canonical-2fd2dccda99d135b0fd73c524bf0ceec86dc7725ba1ccd25be6823c8a809f05d"></a>

<a id="canonical-ba9e4dccece029048df8aea8f4fdee0afbe74bde8c7c7419a138334b13b1a3cf"></a>

## uid property — user_identification / 11dc42dedbf3 / 8

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

<a id="canonical-9d25f1589dff8fb1a9d80c1a8bc4c048c1d2c417a364373c98a297770f53488c"></a>

## Next pages — user_identification / 11dc42dedbf3 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bda151c9403c57e60009cfc6a047b4ac66bcfaa8e64496ec85dca3e02e5516bd"></a>

## waf_type — waf_type / 97fef587c74f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- waf_type

<a id="canonical-267765b8c41d2fdf4390bb490a14e7eb3457f4303e4a407fd7b594c3070758f3"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

<a id="canonical-80ba00b1b5896fa61a61250cf7afbaa1470f2bca7fec43aab26014fdec57869f"></a>

## Direct properties — waf_type / 97fef587c74f / 3

- [app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-7fc3aed76e75100daf9058eefa2f0b0f6f7cbd296f04a63d83585bfcb7220007): complete subsection reference.

- [disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-ac13ccb4767bbc10cbe24a4b8c804140741ee3d51f4fbcaa49d821a8933400c7): complete subsection reference.

- [inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-d5e382facff0f1688da6a540e56a2229b34d9352273ffcea39eb77328a77f913): complete subsection reference.

<a id="canonical-e7e1ff3ac7364320da194b5bba1e97298515e61ed275f7300fff2757fafe8c19"></a>

## Next pages — waf_type / 97fef587c74f / 4

- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-7fc3aed76e75100daf9058eefa2f0b0f6f7cbd296f04a63d83585bfcb7220007)
- [waf_type.disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-ac13ccb4767bbc10cbe24a4b8c804140741ee3d51f4fbcaa49d821a8933400c7)
- [waf_type.inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-d5e382facff0f1688da6a540e56a2229b34d9352273ffcea39eb77328a77f913)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-7fc3aed76e75100daf9058eefa2f0b0f6f7cbd296f04a63d83585bfcb7220007"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a286492f472ca435372c4ab2391ccecedc4c4545a2cb9fb5cd488e3536297d1"></a>

## waf_type.app_firewall — waf_type.app_firewall / ef4db9ff8292 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- waf_type.app_firewall

<a id="canonical-c0aa95543043bb54e77148cad240e23472fa45b6cd36dabf241432591dde6aa3"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

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

<a id="canonical-4f3e39b783e4dc751f0ca28c78d5aca530ee54e7b4ffe86e466176eee17734f9"></a>

## Direct properties — waf_type.app_firewall / ef4db9ff8292 / 3

- [app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-f09608cd1232894e91f87e4a88bc7177d519acaff9e9585af1b92d6342a8bab2): complete subsection reference.

<a id="canonical-7c7dd6bee560b9023c641ab380a71a0c56d968d230f297301f3b7ade57c83a74"></a>

## Next pages — waf_type.app_firewall / ef4db9ff8292 / 4

- [waf_type.app_firewall.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-f09608cd1232894e91f87e4a88bc7177d519acaff9e9585af1b92d6342a8bab2)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f09608cd1232894e91f87e4a88bc7177d519acaff9e9585af1b92d6342a8bab2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7d816d326d8cb5f5fe6ae8eb1de08fd00777f4aab714d537016eecf9b7f44d4"></a>

## waf_type.app_firewall.app_firewall — waf_type.app_firewall.app_firewall / 3f518c610c38 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-7fc3aed76e75100daf9058eefa2f0b0f6f7cbd296f04a63d83585bfcb7220007)
- waf_type.app_firewall.app_firewall

<a id="canonical-cde32b93374d2c59f4d3a3287daaff7afbd96c181ef33ad0ec3d34d97271f010"></a>

Type: `"list"`. Computed.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

<a id="canonical-931bd98d7f711c1a46d172ba21a43c75c503e9ac4176eeebcf46dcbe66ba7870"></a>

## Direct properties — waf_type.app_firewall.app_firewall / 3f518c610c38 / 3

<a id="canonical-9ac7701ae769e8bd33e3767447bfab5a8f18a07446d6bb455ae1c0a50b1978f0"></a>

<a id="canonical-bc9e7c96ab143b95791a72c8a68d78029a0b8f97626e32da9ec03dcd3d17879c"></a>

## kind property — waf_type.app_firewall.app_firewall / 3f518c610c38 / 4

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

<a id="canonical-a424a221510dd1334f1750cff75b57c8a6059949509c13b5ef17131856de9d57"></a>

<a id="canonical-7473fe5f2939201b7e388b8fb76d6547adbbbddaface72e2a7efce82efb24191"></a>

## name property — waf_type.app_firewall.app_firewall / 3f518c610c38 / 5

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

<a id="canonical-cf0b77a9ae94dcc6d45f2cffc9589b14f9ed79132ba6507b3241f9e649da68df"></a>

<a id="canonical-1719d1b98b05866354c10635a7ccd5dfc86d0ba40a3561e267981b33a8fad21c"></a>

## namespace property — waf_type.app_firewall.app_firewall / 3f518c610c38 / 6

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

<a id="canonical-fd74eca47b642fe16cd45ad3197deaeb1cd55f21313b2f5df3e135e0357a95df"></a>

<a id="canonical-2c9e02337d5b9dc382bfaaf47c39f72703965c71dbcced0f65ed97c2aba586d6"></a>

## tenant property — waf_type.app_firewall.app_firewall / 3f518c610c38 / 7

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

<a id="canonical-479481c905e506fa2daa7dfd8961cfd5e5cd29b99fb0519e2218a4d046d30118"></a>

<a id="canonical-74a2a5eee3921d5f19e3a0d1dc7153e0d67287a419fe49d1be5b6a3ba06b85d5"></a>

## uid property — waf_type.app_firewall.app_firewall / 3f518c610c38 / 8

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

<a id="canonical-556cfb2cc918665017231877f177b6e93c778e91b30b93f51668170a37fae856"></a>

## Next pages — waf_type.app_firewall.app_firewall / 3f518c610c38 / 9

- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-7fc3aed76e75100daf9058eefa2f0b0f6f7cbd296f04a63d83585bfcb7220007)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-ac13ccb4767bbc10cbe24a4b8c804140741ee3d51f4fbcaa49d821a8933400c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cddf22774b26766879c4771d2ae79597f98f79bb1a6f208ec1cc58f4614fbefe"></a>

## waf_type.disable_waf — waf_type.disable_waf / bb536a9454af / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- waf_type.disable_waf

<a id="canonical-0761c85d42142623abe0fe0be2662756e0a9533ddd0030a1bf5db82bb9c58375"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-1255ba7d90bb5d1ed23b4ddd063a173605b27fcd00b1be934e229ed023710416"></a>

## Direct properties — waf_type.disable_waf / bb536a9454af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ac62deef63ec0667cfc497775619016ea469882a5cf7e80de4309fa73eb0ff3"></a>

## Next pages — waf_type.disable_waf / bb536a9454af / 4

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-d5e382facff0f1688da6a540e56a2229b34d9352273ffcea39eb77328a77f913"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c45400da1f7e43d4f0ddcf9982a35d431b2d146a8b0ad05730749db4438d3fab"></a>

## waf_type.inherit_waf — waf_type.inherit_waf / dd28525043a1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- waf_type.inherit_waf

<a id="canonical-c4c69772973a07336ba269301194566088878696c7e14de3a7a67fcf98280702"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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

<a id="canonical-61eb57f23e0e4a7bdabd92edc9226939136d993f111b20d759f00896759ce61c"></a>

## Direct properties — waf_type.inherit_waf / dd28525043a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8163b24371d6c363055e17c8463fd29805b1091cc1fe47985da1186a389fa173"></a>

## Next pages — waf_type.inherit_waf / dd28525043a1 / 4

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
