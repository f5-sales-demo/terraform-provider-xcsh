---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-415f53960a4d7f48e843b828393b4ef7534b4ad66cd8ab5360feaa0d37b7b814"></a>

## proxy_config.https.tls_parameters.tls_config.medium_security — proxy_config.https.tls_parameters.tls_config.medium_security / acc4c82ee3a3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-3b7eab2bdefc6de1345f229a642aee5dcd84a1b4fcc0226107c2e4fb7d50e762"></a>

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
medium_security = {}
```

<a id="canonical-5553d6faf4750f8f2921c2500d9485403bfa083ae590c44493f1e7b7e217abab"></a>

## Direct properties — proxy_config.https.tls_parameters.tls_config.medium_security / acc4c82ee3a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2371aaa0c16ba6c9dbfcf72ea371a5562b69d880e16e8c5d912009bbaa6a126"></a>

## Next pages — proxy_config.https.tls_parameters.tls_config.medium_security / acc4c82ee3a3 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-8d3d19560fe789477a691a271848845c6d01c63fd3ae309a100ca27dbb7ba220)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-021f300d96cf53915d6e40040cc9dd60ef5e944e520cd6b2772fbf89fe6cdf25"></a>

## proxy_config.https.tls_parameters.use_mtls — proxy_config.https.tls_parameters.use_mtls / d2d6b7a3e7ed / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- proxy_config.https.tls_parameters.use_mtls

<a id="canonical-a2cfd5f0d13d917425506225e2fda5b10bbb14631551a8342cc52613e0eba3a2"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb9b75fa292ba3bfb018c917e7dd8dd02df8ed6086ca85d51493b31e59d312ae"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls / d2d6b7a3e7ed / 3

<a id="canonical-7b87cb89664363140822129008932f18748fcf82da58396f8391d172b01178bd"></a>

<a id="canonical-473ebedf4b7f93a60b51c1ad487b621e689b93b43317c3a9ae97afc62aeb6205"></a>

## client_certificate_optional property — proxy_config.https.tls_parameters.use_mtls / d2d6b7a3e7ed / 4

Type: `"bool"`. Optional.

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-f178709ff02a36dc1f94499086cef7308af9551d2626c8698ce7a702cd8c985e): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-5a02029ad9c348cb6b6865e3af5976a1ae1e352b6607f78b2d0c3f62fbe42711): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-4475e1bcd7cb8871a5f298433935f3d989fdd2a7b1f23fb596f174051e9b9940): complete subsection reference.

<a id="canonical-79a3457143dd8399674d31f12bc0dae65ac0f3a6b9f10c505d2d1429b45a9220"></a>

<a id="canonical-e12378253ab39fdfe3337565615fadf7d5fa5ebafca7c15839d3403974050c91"></a>

## trusted_ca_url property — proxy_config.https.tls_parameters.use_mtls / d2d6b7a3e7ed / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-1575ae37e1bcdf73d078100927d90860d8af78f00d47c406e45b4124ca14d130): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-5c3e20caccc805da6779966d87aa0ef7fa2457542e31066fbefba4f47cf16e4c): complete subsection reference.

<a id="canonical-0516fb34590cbe944b4c6e80b111a39ecf1702df21ea2644cd253e9b445a25a4"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls / d2d6b7a3e7ed / 6

- [proxy_config.https.tls_parameters.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-f178709ff02a36dc1f94499086cef7308af9551d2626c8698ce7a702cd8c985e)
- [proxy_config.https.tls_parameters.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-5a02029ad9c348cb6b6865e3af5976a1ae1e352b6607f78b2d0c3f62fbe42711)
- [proxy_config.https.tls_parameters.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-4475e1bcd7cb8871a5f298433935f3d989fdd2a7b1f23fb596f174051e9b9940)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-1575ae37e1bcdf73d078100927d90860d8af78f00d47c406e45b4124ca14d130)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-5c3e20caccc805da6779966d87aa0ef7fa2457542e31066fbefba4f47cf16e4c)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-f178709ff02a36dc1f94499086cef7308af9551d2626c8698ce7a702cd8c985e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2df33bdbec80c3035456853cd8cc16be38136d1191e976f4faa3dcb97bf4f61"></a>

## proxy_config.https.tls_parameters.use_mtls.crl — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-1bbfa2daf4a6a1a0882cd09f5b87f921ba1f08e5f9bbdaff2bdeb85ebeca592a"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a5a831e93f763fd99c623711181bf72076a80140b5822e723485c8159066984"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 3

<a id="canonical-066037aec065facd0a8a39dfbd7f80bf10a86ec768c94ad19b6dcce72831f552"></a>

<a id="canonical-07d5bf895585699cd2310b390fd46de56d0af8ad40dcc22963249e5d65430066"></a>

## name property — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 4

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

<a id="canonical-dab11c48577b0c25f27ac111e7d84eeeb2c1fb6b3c675cf42e359eb1faf0fdf6"></a>

<a id="canonical-158cac938b255ea35fec3d523a43fd74016e76da3c38dd6ad9437962b5010422"></a>

## namespace property — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 5

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

<a id="canonical-73bf7bb1b3c002381f10c4bfe93d64a3c7eed079d53c89c0143a11fb5b224a29"></a>

<a id="canonical-6c5d70a64e5611b7a1a088783fafaab3ffa64f2d14ae645dd1838b2f9b9251c8"></a>

## tenant property — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 6

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

<a id="canonical-b47974f9ae3d34954d436a5a1c4a06a27f17d775d37ccf0b5e15c19bfad11c8e"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.crl / 1a0a780b6d50 / 7

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5a02029ad9c348cb6b6865e3af5976a1ae1e352b6607f78b2d0c3f62fbe42711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ed0a99d8c8977d7fc0888f5fec219d73d0a495be1503238a0bfb4e5ebbea9b"></a>

## proxy_config.https.tls_parameters.use_mtls.no_crl — proxy_config.https.tls_parameters.use_mtls.no_crl / 96ec05b0c598 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-9116507121d70ef3a1547c42a935d32044b914437b0da3e330563d43c618cb5e"></a>

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
no_crl = {}
```

<a id="canonical-bd49f26fbc69fc2ede8274d9e5be53825edbd9a19f7673ebefd20bdfa3fad6ed"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.no_crl / 96ec05b0c598 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62020efa295bb6de11ddd3c690970f46776954c658a46133fa260df1e3080ddb"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.no_crl / 96ec05b0c598 / 4

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4475e1bcd7cb8871a5f298433935f3d989fdd2a7b1f23fb596f174051e9b9940"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e4ea3c5f955737e81b9817513778c72a2e427a0eadaabc774cae0f25f619c31"></a>

## proxy_config.https.tls_parameters.use_mtls.trusted_ca — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-5f52e905a14c2482f07f04777009d14452dd608a840d76cabe3317b9b11ad7f2"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed2b8e981c8c3b3fca2c2d3b798cdc9c1eba272534b949baabaf1092febfc3bb"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 3

<a id="canonical-29efaf3f1038de19a30b19a94f31de6d996158976ec440271a761d4c52d78e7d"></a>

<a id="canonical-04bfbfcd6207871618656b2e2eee5805ffbf10ae0037bd89cadfdfc6af5fc1b7"></a>

## name property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 4

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

<a id="canonical-d30e9e42775e50aaa8d63da6b202c879dbb9220d83c304eb59dd544d54ca53c6"></a>

<a id="canonical-00f7e0b91f37e808b389ac9dfb64ae76e91adccf1ba49ad28c2159efe5f5fa0e"></a>

## namespace property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 5

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

<a id="canonical-0a8c32b25d086678fcaece890c55c582e4b4e20e1edd642bc879e71bee3a0fc4"></a>

<a id="canonical-dfac41b5bc5678c0a05b34f3a7453c9b3634437f6641535a8ce51c8092e73f82"></a>

## tenant property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 6

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

<a id="canonical-1de575311563aa260e5ce94c019ed71f4843910d83cd626044e7729344906130"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.trusted_ca / 7d9a79d90138 / 7

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-1575ae37e1bcdf73d078100927d90860d8af78f00d47c406e45b4124ca14d130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a99021fa735a8a375c80c8a5fe828e3c94596523f7f25a86897133958fec2f9"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_disabled — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / 797f24e0ac9c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-842a8cfcc03e1d149117caf689a8f5a05df0be98dbf4b2380fac139fb2194128"></a>

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
xfcc_disabled = {}
```

<a id="canonical-3a6ac4761a4cf1423391ae958e326bb51864562a41b31f3a92076ae7da75772e"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / 797f24e0ac9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e99515e77e7fa8203245914de7e76a92ca3ce638aba1c32f51af5498d9b39a8"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / 797f24e0ac9c / 4

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5c3e20caccc805da6779966d87aa0ef7fa2457542e31066fbefba4f47cf16e4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d81016d84b191ed59b527a3d758959e11000aa1e50e5a054efd035050fa9043e"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_options — proxy_config.https.tls_parameters.use_mtls.xfcc_options / ffecee3b4130 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-e3664ed7913a0e6204cccd79e8b7ecde744c50efe983cbd84cf912e843de1655)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-1a2fff832c914deeb21855b2b6d870103e691b9c714316eef87fcff21a04a187)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- proxy_config.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-868f10fe1085de369723b747f1d57e9e5d93ee636e321bfdee5f4fe4f2674990"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-7d558713475cb0e1e11dff8cb1242270fe7c9b2b40b3fc72c7df68a0f3514bb1"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.xfcc_options / ffecee3b4130 / 3

<a id="canonical-2d37fd2098beb3c1d49cd7d43376c78010d48da9fb69e5848b880ba9e463a863"></a>

<a id="canonical-c64c1a3969b50efb6e98c984e151023b6e77a0a71f7e7b1778aae0ddc0ca2a31"></a>

## xfcc_header_elements property — proxy_config.https.tls_parameters.use_mtls.xfcc_options / ffecee3b4130 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-d2f53637cc5662677f5e435d8d4813d88c43b943dfa9cf204e7a7929e7b398ba"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.xfcc_options / ffecee3b4130 / 5

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7c66469fcd6b8d3805bd3f8b1b7b58d8a3f9725abb801b3be3c1a43c83ba8830)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-861ef136ae5bafcc3052971b0b907169ebcd764c0881acfd01234c71c0977ca0"></a>

## proxy_config.https_auto_cert — proxy_config.https_auto_cert / 60f67635e854 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- proxy_config.https_auto_cert

<a id="canonical-6e73f6a85aed1b37e9e1792428626239c9ddb53cd35bd009dba978caca799f2e"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-38875743b138e4d43c594fcf2a91f1a64b9d53ffce772cde68d43e3d49b1d6c7"></a>

## Direct properties — proxy_config.https_auto_cert / 60f67635e854 / 3

<a id="canonical-49cbc6e967036ce230d86df6fe6c5783a90686d7313a411ce0feaa31ad27b9f8"></a>

<a id="canonical-831a88b5053b692e755f999c6b993e9ea025bd3dae9c3866df04f02ef44b5856"></a>

## add_hsts property — proxy_config.https_auto_cert / 60f67635e854 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-87db4c1524c42c7dab588cd78f2bf6431b6a9fb2d47f9fcdd070dde6208ca351"></a>

<a id="canonical-25778b42da3973a7955cae93244b443a39437dc2c70c49446d606cc1ad4c0bc3"></a>

## append_server_name property — proxy_config.https_auto_cert / 60f67635e854 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225): complete subsection reference.

<a id="canonical-18af64db4e311d4300619caf4c2788e3fb8b16167048ada82171e5c3aa63400b"></a>

<a id="canonical-440e868f4e5025becc6277b48d81b64773c83300f3168debf8766a289c8b0fe8"></a>

## connection_idle_timeout property — proxy_config.https_auto_cert / 60f67635e854 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-fae8b2523c2e0299dba8188039a18e5247c013a951de37cf81b2a49bf2c4a8f0): complete subsection reference.

- [default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-7aaf37a1e38ec4e8a34e38846d77b05602c5c338cac33aab7a388cb5dc43adb5): complete subsection reference.

- [disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-700f51e460b11969580309f4b9c012bf53610024ccfb8b90298a34338ba5ab23): complete subsection reference.

- [enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-b1bbc09b412a84fae810b7ef3bb7dbf12537a9859baf6d11bef7f43cad74fc7d): complete subsection reference.

- [http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89): complete subsection reference.

<a id="canonical-6ae6383fcd154436155071e01152e05d452b77a1911af305cc6c89ea0a1c6e28"></a>

<a id="canonical-592a002d78aa82ffafe3fdb8691b3b2ef6da1473a6a48f27e04a972f1598d68a"></a>

## http_redirect property — proxy_config.https_auto_cert / 60f67635e854 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-8d056fe8bbab028ff68d08db36207bbffd8ae80ba9e8d5487a7ee500413566f8): complete subsection reference.

- [non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-4e7538ab7e851e6619b1fc914c18861f538bf5a8a9ba01c0171fcecb2ce8a8e1): complete subsection reference.

- [pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-f46714a5b78570036c39aa3e8b206bad46212d12c4a7ad83d9015956b6c1b532): complete subsection reference.

<a id="canonical-8c2d463f85fe1287fb7cef2f84f85436243cebdd2354f17031db3dc64f1a4ec5"></a>

<a id="canonical-07ffbd440c87fb0c896f1511aa10f0ea6196a1b0ff973d6345ed22af0e4cddcb"></a>

## port property — proxy_config.https_auto_cert / 60f67635e854 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-1d6ec60f7446084c88f9af36b7b69a539d5081909ea28da63067c52423ec5027"></a>

<a id="canonical-723a3268e0b7a45dbb5d7394e3a3136bec5e0ce8ccfe6f3aa3fd7bb7e49f904b"></a>

## port_ranges property — proxy_config.https_auto_cert / 60f67635e854 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-ea14d87d4577b906998a1745b9e316ae3583f5bd8d3178e9b21cbec7c6f0614d"></a>

<a id="canonical-70162eb6c19708c26f931709686ed701b385d0b01afdb88af6f5b4a4641f3f28"></a>

## server_name property — proxy_config.https_auto_cert / 60f67635e854 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1): complete subsection reference.

<a id="canonical-890c0bc55d3c7e4521ada200fa54e8a0a370ab21222dc3c9400b842b380ee6a6"></a>

## Next pages — proxy_config.https_auto_cert / 60f67635e854 / 11

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225)
- [proxy_config.https_auto_cert.default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-fae8b2523c2e0299dba8188039a18e5247c013a951de37cf81b2a49bf2c4a8f0)
- [proxy_config.https_auto_cert.default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-7aaf37a1e38ec4e8a34e38846d77b05602c5c338cac33aab7a388cb5dc43adb5)
- [proxy_config.https_auto_cert.disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-700f51e460b11969580309f4b9c012bf53610024ccfb8b90298a34338ba5ab23)
- [proxy_config.https_auto_cert.enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-b1bbc09b412a84fae810b7ef3bb7dbf12537a9859baf6d11bef7f43cad74fc7d)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [proxy_config.https_auto_cert.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-8d056fe8bbab028ff68d08db36207bbffd8ae80ba9e8d5487a7ee500413566f8)
- [proxy_config.https_auto_cert.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-4e7538ab7e851e6619b1fc914c18861f538bf5a8a9ba01c0171fcecb2ce8a8e1)
- [proxy_config.https_auto_cert.pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-f46714a5b78570036c39aa3e8b206bad46212d12c4a7ad83d9015956b6c1b532)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-072f2ec0d147da55fea904c1af937586ddab5ca452d4c9c4c6340e8441c8b1a7"></a>

## proxy_config.https_auto_cert.coalescing_options — proxy_config.https_auto_cert.coalescing_options / 52b02806b620 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-8b68ea9f38d8c3462f1f03cb69838e58dffb2fe65b39090c1cb6fe8c581a71cf"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-e4c1faf16e5d6545ee90c5c152179bd2c5b1147b8ab565e43cd94ef9f0460f20"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options / 52b02806b620 / 3

- [default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-4cd364566ba5af7e0002b846b0ebf0971e6787db60c53684c09118fb32e3368f): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-9ca1348d8903311ef505cecd57edbd68ae8c4866cc6f743c20b35dd0eb102f1d): complete subsection reference.

<a id="canonical-deceeb43cef91a77c226a69a46b78339c779f9798cafd97decc5e306bc7f6401"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options / 52b02806b620 / 4

- [proxy_config.https_auto_cert.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-4cd364566ba5af7e0002b846b0ebf0971e6787db60c53684c09118fb32e3368f)
- [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-9ca1348d8903311ef505cecd57edbd68ae8c4866cc6f743c20b35dd0eb102f1d)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4cd364566ba5af7e0002b846b0ebf0971e6787db60c53684c09118fb32e3368f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-706c0732af851b402af3229d9d1d4c4f5e6af9163f1fb4ca6649e0b69f29b735"></a>

## proxy_config.https_auto_cert.coalescing_options.default_coalescing — proxy_config.https_auto_cert.coalescing_options.default_coalescing / 1b89b6d5dd29 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225)
- proxy_config.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-55c5364e26157de294c5fc705de2ccb04094b64de7e0012563da77f9aff516b5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-3d1f772d58ac0450a5a134ca99ac992803e265eb1415cc3ac2ba8c1cbde1bf12"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options.default_coalescing / 1b89b6d5dd29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31da6aac68043d3cd2116c21534fb011c6f618c13aec53cc9bce3b78e9a7284a"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options.default_coalescing / 1b89b6d5dd29 / 4

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-9ca1348d8903311ef505cecd57edbd68ae8c4866cc6f743c20b35dd0eb102f1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db16ac214d1fccb3c0e609eec4c527d2d8e1ed9e77e92b3d792bcf33986da2be"></a>

## proxy_config.https_auto_cert.coalescing_options.strict_coalescing — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / caf75797f583 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225)
- proxy_config.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-c4da4f577e3775bf63626abfdd6e36cd2eeca53e25ce5c35d694a7751c5f80db"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-0cb6a6e14dd50364c008c88790124882f72b286676912d9af920d9e7df982321"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / caf75797f583 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed6426fb4e46c62fc019ecf6cbf6a79a6bbe9dbaa3848b1c33512bd592559076"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / caf75797f583 / 4

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-4f7596e8495c51a226d24dc3c721b74c2ec0bddac6354ea467b86a0665943225)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-fae8b2523c2e0299dba8188039a18e5247c013a951de37cf81b2a49bf2c4a8f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-818f2e2fb39a1bf4489b0d19d4f4119edf6dbbf05687f2041e9a074140f0f27b"></a>

## proxy_config.https_auto_cert.default_header — proxy_config.https_auto_cert.default_header / 9b6d3cd82937 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.default_header

<a id="canonical-e77d40b407696c87dbeaa232bc62835d2379891dbf90b3a4c9de42af382ba351"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-31c68dfc91bc6fc4dbf43e478a6c134a77026e0917947fb7fd5bc4efec63f881"></a>

## Direct properties — proxy_config.https_auto_cert.default_header / 9b6d3cd82937 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ad1484b8ce2245c6efe441a9ad1fc6af4b719cd8951e9e9a13a401971680d0a"></a>

## Next pages — proxy_config.https_auto_cert.default_header / 9b6d3cd82937 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7aaf37a1e38ec4e8a34e38846d77b05602c5c338cac33aab7a388cb5dc43adb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2012b936149557d819d0e462244646384ae80312b82698b75f33fdbef2d4af0"></a>

## proxy_config.https_auto_cert.default_loadbalancer — proxy_config.https_auto_cert.default_loadbalancer / 879b4aa5f5bc / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-46c67bf9c1a3f3a9282a360d292037022f969e16c985722ee7a163aeafb7981f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-d42f9635e6e859dd0fe66ad602e61e73e5c075f0efa1bc317f0d22c2d16805e3"></a>

## Direct properties — proxy_config.https_auto_cert.default_loadbalancer / 879b4aa5f5bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3bca8a2dba5b97c93c8c3974b409fb5bccf9c5ee24fff4c7cd941eb9423571e"></a>

## Next pages — proxy_config.https_auto_cert.default_loadbalancer / 879b4aa5f5bc / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-700f51e460b11969580309f4b9c012bf53610024ccfb8b90298a34338ba5ab23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7db2a9cebafd46797d2d244dcf75c660f29cf1dba4f3b7dec482d728f1a9afbc"></a>

## proxy_config.https_auto_cert.disable_path_normalize — proxy_config.https_auto_cert.disable_path_normalize / d06f74b596e5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-e81d8537f776d8beda7c715b636e863ca5fb6428fe6f316a16392db70ff8721a"></a>

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
disable_path_normalize = {}
```

<a id="canonical-802c4bba52ba0b6c3e17db81cfcd0e90cd9ee400080b7b6b39c4d7506a6d539b"></a>

## Direct properties — proxy_config.https_auto_cert.disable_path_normalize / d06f74b596e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e435737d438765b98aac7a3833cb74633f8804f51a6fbd6c8a1098ffdf457c9d"></a>

## Next pages — proxy_config.https_auto_cert.disable_path_normalize / d06f74b596e5 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b1bbc09b412a84fae810b7ef3bb7dbf12537a9859baf6d11bef7f43cad74fc7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-572f2cf2e85858d194e27e35434a6a775a80e309f94bdefa78e6063b3a8aaf31"></a>

## proxy_config.https_auto_cert.enable_path_normalize — proxy_config.https_auto_cert.enable_path_normalize / e0f13832af58 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-688863d8768b7dfa9b3505cdc0dc9806a224933b38e7d0498935afae4da0a1e3"></a>

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
enable_path_normalize = {}
```

<a id="canonical-60a0d260e7d7614d041fcda39e4395ae6d87af3be549651e7ccb49af3d9d93c2"></a>

## Direct properties — proxy_config.https_auto_cert.enable_path_normalize / e0f13832af58 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f3f59e927b60e7da6fa42664b8c773fc1750b3d472d63a427c0d17b053acb7a"></a>

## Next pages — proxy_config.https_auto_cert.enable_path_normalize / e0f13832af58 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0786f5c228545b37484b081d14a7cdd0eac9594a990a68c058763d3b4ed13642"></a>

## proxy_config.https_auto_cert.http_protocol_options — proxy_config.https_auto_cert.http_protocol_options / 7f9b9691a7de / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.http_protocol_options

<a id="canonical-2ac09f80a108783afbe72e1b3ee56d7b09dacdb18dd8b3a943073eefba31b5b5"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-da7a8d9992b2d8e68a549a575137bae8c2a0caf0556d7bf6330f9cce068b0965"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options / 7f9b9691a7de / 3

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-51a39ee0c4cf5f8aa04d7d5b8519f2ed20cdf00ca327c2847e425d3765b77328): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-33e71d26217ce7f58a9c557aa1372b44ac203a0cf4383f63a3dac02144b16d84): complete subsection reference.

<a id="canonical-77da60e24d1c2e577363d0d150254ea84d2fa4c33a06736b22ed2541f620adda"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options / 7f9b9691a7de / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-51a39ee0c4cf5f8aa04d7d5b8519f2ed20cdf00ca327c2847e425d3765b77328)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-33e71d26217ce7f58a9c557aa1372b44ac203a0cf4383f63a3dac02144b16d84)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adcda3ddcf979ff681d3ca1bbce3231645e519f33b233f3d6f5a2a7defca5b93"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / a62abd823223 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-e15485a4cc03412614fe65912546121c655f56d561ef516748208d8437f42091"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-96badae8ff8de27db0da2df245d4f2837844fe2725c751bef2a4768ab59dac1f"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / a62abd823223 / 3

- [header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2): complete subsection reference.

<a id="canonical-b76f2f2198caa5e3b28af3f4c2e5a4ec13dcae26d48caf928f36fc6605e2ffac"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / a62abd823223 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d275d4410d3b4c3705c82d5bcb7352b5ec44c396c799965074895f700f9e61e0"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 67d434a99593 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-b63627d3566db26fc58addd032afc8304c05a9fa0a840063256548e99863e534"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd96808179aa020d2b385faa0f1988cda2d58f77351feefd5625cc995694bae9"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 67d434a99593 / 3

- [default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d0f58d776b7f51ce9da19e8087be0b8009252d7f06952c2a201250589752b3a9): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-f74ecaaacc13c81fc872de0fb004bb07ad8516ec6805b8e550da546817167402): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-b8f5503c01d08fe7cb424a482ad99a1084b637330eeb79b6499652237da71e41): complete subsection reference.

<a id="canonical-1b72ca59993995e03be2f4275d4f93471afb66e4a68743cf034b00c767bf3110"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 67d434a99593 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d0f58d776b7f51ce9da19e8087be0b8009252d7f06952c2a201250589752b3a9)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-f74ecaaacc13c81fc872de0fb004bb07ad8516ec6805b8e550da546817167402)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-b8f5503c01d08fe7cb424a482ad99a1084b637330eeb79b6499652237da71e41)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d0f58d776b7f51ce9da19e8087be0b8009252d7f06952c2a201250589752b3a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8873ed5053182d617f5dd369750c3b0cdbc8a4f4e70f9f0893fd69dcc1c39483"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 2d020689c015 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3497eddb36495381b3aca9797b6f5a5669a3a45ca54437185b9b1d9e8004d3db"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

<a id="canonical-06d7b2ac4356437ee67a41d021384df88a7c2636abbe58fd4bba319d39773751"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 2d020689c015 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64d304bdaafa81b0e943e0dd44d13140240da9d8002964f27483412527d57d9d"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 2d020689c015 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-f74ecaaacc13c81fc872de0fb004bb07ad8516ec6805b8e550da546817167402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3f16703d6b85f0e68eb0ef3dd182731ec3834e748a3a6c4b034fa34d2c43590"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 510eeeadd04e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-46bb6e5123ab3e50c6896eda15286455cd33a02c85555e6fbff75b16ec090370"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

<a id="canonical-302593a99f484ad408f0edbc67efa1bff34f789a779d14a703811babf25821b8"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 510eeeadd04e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e701c7800ace1d43ab3524bd69c499df40de77541033aabc0f41de4aa7a933c"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 510eeeadd04e / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b8f5503c01d08fe7cb424a482ad99a1084b637330eeb79b6499652237da71e41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a328f1da1e790b6f475f1e801a399d975f0556494e6806ddb1b62705a2dc2e51"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / e8e16c02f589 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-728435dafdc51a3b781946378eb4c9c1f32e2269ef2e84b636e02416d59dd418)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-8dfff61c77fd6dccd478839f98a907a8fc97b7aaa8cf87d47d113f7bd0b3c184"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

<a id="canonical-490e5b17a7da4d1ae7520b72fd31cf411c0b2917b726e0aeb77ee4b672c3ef88"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / e8e16c02f589 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07ecd429163fc8d08110d1ac197be5ed083ef8520e753588b6c4d6dee3a5c842"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / e8e16c02f589 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-d9dada07e40d4d5492da28de24dfd830ac1a8d3d4d6bd5e115f0eda971e47db2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-51a39ee0c4cf5f8aa04d7d5b8519f2ed20cdf00ca327c2847e425d3765b77328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c06e805bf3acb05589741e0afc503e4daef53f2bf659eae62b46a692835f872"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / add003c331a2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-1dde5b3889ba9e797e56f805cb221f365cbec507f6e276dd0efdab84e5432a1d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-97204afad3642ac573bd50f3aaf41706c9d7314de72483adb29dd1b972b3a68f"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / add003c331a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aebe6e370e18b5c58931d64636bd6b2151358682bf14d927a2b5880d7050d5a8"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / add003c331a2 / 4

- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-33e71d26217ce7f58a9c557aa1372b44ac203a0cf4383f63a3dac02144b16d84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5a0a389df3c37079f763fa4ffbe6bfc95499b625561745871f8d6c8d1addeef"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / dbcaacc3af56 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-e620f4ba679bb3982780ae873178e9239b29c8f66fcd51450d7be23923998a31"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-d68d16a5283c8f2d60cd0b8149ae4ac34bf4096bd4b1c748aa165662dea9ede0"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / dbcaacc3af56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d1c5880f89fa4ec01c5ce60abf5bbd374b8d69cc3bfe50a8cb995e6da8fe086"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / dbcaacc3af56 / 4

- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-615fd3d8ffdd3dfa4b7dc76ef6afb17d20323bbe3e5434fc535dacb8042e8f89)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-8d056fe8bbab028ff68d08db36207bbffd8ae80ba9e8d5487a7ee500413566f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c38d632bb00d25ba9ec18677d6f3111e1e78c62895c039c95e0ed95e6b343ed"></a>

## proxy_config.https_auto_cert.no_mtls — proxy_config.https_auto_cert.no_mtls / a39e75ba5b57 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-7af6a275121c6ecfc5cda3ba03ba4e6fa38f1f2bc7d8457b2e55749a75389b04"></a>

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
no_mtls = {}
```

<a id="canonical-096b0eab54844c4a0cf68c79a9f92759d06c57d63626658e586886e4dce92121"></a>

## Direct properties — proxy_config.https_auto_cert.no_mtls / a39e75ba5b57 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fc8d7c50e2b52ab64e948207ab11e9ccbec99f8634a1027e8eebc2bd1635cd7"></a>

## Next pages — proxy_config.https_auto_cert.no_mtls / a39e75ba5b57 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4e7538ab7e851e6619b1fc914c18861f538bf5a8a9ba01c0171fcecb2ce8a8e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-550c87337b4639b7efc3aff6d970f527c6b5a424ba7198da6b7de025a2b82b16"></a>

## proxy_config.https_auto_cert.non_default_loadbalancer — proxy_config.https_auto_cert.non_default_loadbalancer / 80d0bac39fb8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-a80f9ce26fae60cd6a30b868716f56772e37eeb5b7f8f05f0c237acd14290caa"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-ce4a845beae451463335a563064c2204b9316666bd557a1d441318e408186129"></a>

## Direct properties — proxy_config.https_auto_cert.non_default_loadbalancer / 80d0bac39fb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa0a4837ff4e55afef8a526c9f747e1fba78c07618b44cd9fae7ce5ba80a0048"></a>

## Next pages — proxy_config.https_auto_cert.non_default_loadbalancer / 80d0bac39fb8 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-f46714a5b78570036c39aa3e8b206bad46212d12c4a7ad83d9015956b6c1b532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ceaceece5b11c0af6caf85bc43ddff832ef0b03ea04243fde36bd58f5b07dc0"></a>

## proxy_config.https_auto_cert.pass_through — proxy_config.https_auto_cert.pass_through / 4627024c283d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-457dea171029fcff29b5f14f4209b7a5cbe31b554d7cd33aca39abd6ea124989"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-20b96485404e0aea34b62873e3279cbe3dabcf3002f201e90e0ea86e34c888f6"></a>

## Direct properties — proxy_config.https_auto_cert.pass_through / 4627024c283d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bbeba1d80d88fe0af2f3db34b676df7d28bb83b6b140e79469a04033e5709161"></a>

## Next pages — proxy_config.https_auto_cert.pass_through / 4627024c283d / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-743373a35afc289d0d88afd5f158b5008bb9139ae0793e7c4e788116e9cc941e"></a>

## proxy_config.https_auto_cert.tls_config — proxy_config.https_auto_cert.tls_config / 46d89ebb1f8d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-51e770596702f303f71b4200bc1a0f56e839914f459cf469a35713cca298c125"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-ede9c4e6f8b4a4a52d2eb7ce8e68a90f7a353d577c5f25455c12ea4b9a8730d4"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config / 46d89ebb1f8d / 3

- [custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-99eff1f2f0e288cc50c24143326bcd6ee68edcf3ef5610adb34685a2ed3965ce): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-753cf1f1b520fbce7852365d92440c826f2533d06fe01879a53a5d0a4143f844): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2e3cfab8c084ac138492b4b6230e2abe1e6cefc93350870f4db9a905039a0301): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1df1780e5d9e772eaf62b49c810ef6058607b86f576360e9917c326787b15226): complete subsection reference.

<a id="canonical-d39258462a2aaca24a77ce676f0570378ad275a2b3c8f92e21b0ab9135d2cf1e"></a>

## Next pages — proxy_config.https_auto_cert.tls_config / 46d89ebb1f8d / 4

- [proxy_config.https_auto_cert.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-99eff1f2f0e288cc50c24143326bcd6ee68edcf3ef5610adb34685a2ed3965ce)
- [proxy_config.https_auto_cert.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-753cf1f1b520fbce7852365d92440c826f2533d06fe01879a53a5d0a4143f844)
- [proxy_config.https_auto_cert.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2e3cfab8c084ac138492b4b6230e2abe1e6cefc93350870f4db9a905039a0301)
- [proxy_config.https_auto_cert.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1df1780e5d9e772eaf62b49c810ef6058607b86f576360e9917c326787b15226)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-99eff1f2f0e288cc50c24143326bcd6ee68edcf3ef5610adb34685a2ed3965ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3b02cfb5f7aa3af964d3797077131e29428ca8f5c1343c2c75fda3a0d918ec7"></a>

## proxy_config.https_auto_cert.tls_config.custom_security — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-b78e443dcfce007e3c23dbd3430478e0a631c5a183b8904d617f5e18d35595ab"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-512d0e76e2ba0017408f6f1eab9e429b9202541817b436e3cff7002d68756e7e"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 3

<a id="canonical-68aa3f8a670426ad26814f59720229285d0712d847d3ce356616b56978d92b14"></a>

<a id="canonical-5e42e8fe47ee4b41de26a20dbfca9eb4bb81970cbad71f96e66c53ecc8da2763"></a>

## cipher_suites property — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-865ef665233461ffb071e6d5343ba4679dd11b4c8cbf313a799862db7c8a493e"></a>

<a id="canonical-1e2bdf6a3a460993ac3d66a150e1fc5da0841bf4f946d34a306e9e8ef574d349"></a>

## max_version property — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-cd551c5501b6acdecb7222876bdd7718ec38e87769fd3f98a0acab476059d8ce"></a>

<a id="canonical-2e661ee5f28bebf0eb5909b027c417ec8786691f92b816958fedae0a9783858e"></a>

## min_version property — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-93bed36dee3bb3a86d2c75d99085381a52a9a7df00eb98a7ac68812504afd0e0"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.custom_security / 856226ffd67e / 7

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-753cf1f1b520fbce7852365d92440c826f2533d06fe01879a53a5d0a4143f844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-269251bf16807a80a8c1a1874ce8adce223a0a49a454c0040e6f35776967a447"></a>

## proxy_config.https_auto_cert.tls_config.default_security — proxy_config.https_auto_cert.tls_config.default_security / 01f8857b1b5c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- proxy_config.https_auto_cert.tls_config.default_security

<a id="canonical-a2e4205dbf11b7d105339ed86edf4bca98609e5ef717dc7baa4935c8f909d0f2"></a>

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
default_security = {}
```

<a id="canonical-4f7905be8a4789571b1b808ef08fe924f077e5e537916f6864f7a9129e416541"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.default_security / 01f8857b1b5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b177121ee27a1da7c7b1bc8ac763d31fda7bff8330f620163562a83886acfdef"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.default_security / 01f8857b1b5c / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-2e3cfab8c084ac138492b4b6230e2abe1e6cefc93350870f4db9a905039a0301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82bffbb7eceb9b80f4ca891e9caa485a1ae6126d0d922177953a22b043f25ce9"></a>

## proxy_config.https_auto_cert.tls_config.low_security — proxy_config.https_auto_cert.tls_config.low_security / 90cf65deded8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- proxy_config.https_auto_cert.tls_config.low_security

<a id="canonical-1d6b16d6a28877d769d76cac703e8f36f4f6692338045b4822fd02c9de1d123d"></a>

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
low_security = {}
```

<a id="canonical-dfc55004b257688dc1c38003aba77a63c99879b89ad51bd9d2cbca95840c6cd6"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.low_security / 90cf65deded8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58c77dcbf3e908f513615d473eaf447c6ef287c0ef62d187946c06535ccdbd9e"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.low_security / 90cf65deded8 / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-1df1780e5d9e772eaf62b49c810ef6058607b86f576360e9917c326787b15226"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-405d98a0097985857bd9a9e5fdfa06807254f9e38407d11aa4231d854c8c35bb"></a>

## proxy_config.https_auto_cert.tls_config.medium_security — proxy_config.https_auto_cert.tls_config.medium_security / cfd1f9f3de80 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- proxy_config.https_auto_cert.tls_config.medium_security

<a id="canonical-79bc2315de52adf9d54743368d6716b02562fedbcd588937b8c474554ec17d5c"></a>

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
medium_security = {}
```

<a id="canonical-84917925e19637f39cb30a5e9000e23617fdc26557fd762ce7786ffdc80c44f1"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.medium_security / cfd1f9f3de80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-940ef3e3be41013d918e2fd0243b354ddd685e280475de9012f7ddb73bbc501a"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.medium_security / cfd1f9f3de80 / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-4ce54627de49fc23bcab3883c6f09982744e9a09afe8c13bb612cff1046fbd84)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05904915eda5ba25b36870f86487f3f07ddc3dad90df56054d4643f159ae167b"></a>

## proxy_config.https_auto_cert.use_mtls — proxy_config.https_auto_cert.use_mtls / 044573c0a45e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- proxy_config.https_auto_cert.use_mtls

<a id="canonical-3a884abae75ebbe76a2bd80e94b51e697d52d4a62b057fa649f6762947e059a5"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-756cc95a81c373a3d177a7cd52bb0801d32532a71edb562776f661270a1c0836"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls / 044573c0a45e / 3

<a id="canonical-9b703272f479965f87071e224f936a4abb6f6c15d28fa76c26c11d5d8453e36d"></a>

<a id="canonical-083ba97c764bcc7e37853a6a6412ef0cc9a253ab3857fb07043d9f71b4e2ae94"></a>

## client_certificate_optional property — proxy_config.https_auto_cert.use_mtls / 044573c0a45e / 4

Type: `"bool"`. Optional.

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-f7a9038505b12da73bf52861015af7c9861dac8e6e5a90787e83b28d4129cfbc): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-5eae543838ae963c921fca09c8ecd5f3ac107871a6f41b335a4063e0fbcaba80): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-47109dad094b5605c4f4c330e5c7982617871d2fcae041a21de3de23b4581960): complete subsection reference.

<a id="canonical-04d3e74b85e8b941530ae926514a7ff91a4327ed077a8d97c271d48f52fd31b2"></a>

<a id="canonical-0b3ccb9eab30dd3a3d56423239bbd78246b051c390cbc33b84c1b01367ada1da"></a>

## trusted_ca_url property — proxy_config.https_auto_cert.use_mtls / 044573c0a45e / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-e587d60ed9ad0a02ba79d01f0a28fa2d579314aad4cf5dc7b930aa9ba9320e82): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-3dd3a10a9be9133b4f76ba23f556566555847e4e4b525ae1f39cbb4d00170f80): complete subsection reference.

<a id="canonical-3b036b5aedef5b0dcc12c36fa9fbd0be5231c3f769e772953598307b7a4dda63"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls / 044573c0a45e / 6

- [proxy_config.https_auto_cert.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-f7a9038505b12da73bf52861015af7c9861dac8e6e5a90787e83b28d4129cfbc)
- [proxy_config.https_auto_cert.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-5eae543838ae963c921fca09c8ecd5f3ac107871a6f41b335a4063e0fbcaba80)
- [proxy_config.https_auto_cert.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-47109dad094b5605c4f4c330e5c7982617871d2fcae041a21de3de23b4581960)
- [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-e587d60ed9ad0a02ba79d01f0a28fa2d579314aad4cf5dc7b930aa9ba9320e82)
- [proxy_config.https_auto_cert.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-3dd3a10a9be9133b4f76ba23f556566555847e4e4b525ae1f39cbb4d00170f80)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-f7a9038505b12da73bf52861015af7c9861dac8e6e5a90787e83b28d4129cfbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d0e3e572722c609f9b46899180b46a4eb53d7d2b13220f3f29db18a145446f"></a>

## proxy_config.https_auto_cert.use_mtls.crl — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-da41602f3563f511e64796eda211e1bd0f44a74cf086920c5bf3c8deba2941fd"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0331e7993b920433c42a40023d5398e33d4bd57d6d2e0020ce6ca3b706540bd"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 3

<a id="canonical-8880989e586a671c86b9e14a2c3ab7ea1c43dba24a9349314fe84c16a3748876"></a>

<a id="canonical-8389ce0cfe4010b919cc4d7a7c46f96d71e721dac0283321340c8a1a9d04fadb"></a>

## name property — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 4

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

<a id="canonical-d7ae64d7bbb290997bbc1157e5dfc9a21f1054ad671c913d4e373bf4144cca7e"></a>

<a id="canonical-6b76a9080317b93aada4b0963af3dca65dc773c0040a8929ef497d0b5e248466"></a>

## namespace property — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 5

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

<a id="canonical-56a124e77db3dbadf4a3c89b3af013d6f3b958f825e23afd87216ba72446e021"></a>

<a id="canonical-f774402395cb92dd2d174a07d824a05c7d25e19946f92f4e595f5a337e543089"></a>

## tenant property — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 6

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

<a id="canonical-e6d63c067eeeb822ce9e749a752289ecb976f1691e15859fcfac7de86fc0f108"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.crl / 933b841e922c / 7

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-5eae543838ae963c921fca09c8ecd5f3ac107871a6f41b335a4063e0fbcaba80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d918f831fed648960d47239b61e6126e51373e46413b5220c3a0086c1542f64"></a>

## proxy_config.https_auto_cert.use_mtls.no_crl — proxy_config.https_auto_cert.use_mtls.no_crl / 28f9d8e99bd8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- proxy_config.https_auto_cert.use_mtls.no_crl

<a id="canonical-3d3f41cc0a2d56b4788397890278edb02ce74b27b17445dec91d62f470aebd36"></a>

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
no_crl = {}
```

<a id="canonical-b436f5da13c9b2a5ec1cfe68dfe33ad2d65bdf46bddf9e2176a7fde3d6ebfa6c"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.no_crl / 28f9d8e99bd8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03781212a8928243cc0565c41d38cb09780a5efc9c9296f397a2cdbb44a343c0"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.no_crl / 28f9d8e99bd8 / 4

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-47109dad094b5605c4f4c330e5c7982617871d2fcae041a21de3de23b4581960"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccf6b089026b7698e38b19c17b4681c7392015f5f2f33576f083f07a6da84d44"></a>

## proxy_config.https_auto_cert.use_mtls.trusted_ca — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-4a9a8e017b05bd6023348f5dc13786d830ee14c2371c683e446235721fb8fa13"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2532de1c1ba1c83c478f368e5083f7da7682a9b47fe666327c93b08b9c79c735"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 3

<a id="canonical-ecde507f0ca67adcca514d1469af28356075d885dde4d0f9a600e6eec3f20d61"></a>

<a id="canonical-0cee214be74736b62409fdb0bccc0f245ee7bb1a0dbb977787b59a6f6a12fc01"></a>

## name property — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 4

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

<a id="canonical-5fd62edd63094372582c66b94bc21abadf45e0f80fe8bb9031396ece43607555"></a>

<a id="canonical-7887620ecb165b5c9ff3af044b0cc687c70ef758af6a441ca40b88c4a3ed5c2c"></a>

## namespace property — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 5

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

<a id="canonical-1b68b2b4af1a9ca9a939340d93faf2154853e1d0efd8bc67e5fce7e91bfcf02e"></a>

<a id="canonical-5e6c5921447016b9936b5daa5efe5cfd7c8ac81c2c66f0821579d24c64ca2596"></a>

## tenant property — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 6

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

<a id="canonical-94d632ec2ad9e962f81ee96db16d9db1c6dad73001e86252f7927f1682aaf777"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.trusted_ca / afc144a3ff03 / 7

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e587d60ed9ad0a02ba79d01f0a28fa2d579314aad4cf5dc7b930aa9ba9320e82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba60d3cf829df4b714cd09f7793c708415b463c30558054853ca1fa9e64fa169"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_disabled — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / c2d67fddf7ac / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- proxy_config.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3e251cd5b99b7b97b5ba6075a68c514969fd10708782883f6bef4b940f74700b"></a>

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
xfcc_disabled = {}
```

<a id="canonical-10fe81543a0b2b08f40bf2a404fea1b471484e388018ccdff7fd5ae5887f1654"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / c2d67fddf7ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03502d2e114971dc2fe7c48fc6496cd25f554b353e455de64efdb69f741d1419"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / c2d67fddf7ac / 4

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-3dd3a10a9be9133b4f76ba23f556566555847e4e4b525ae1f39cbb4d00170f80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cf6e4729e2bc4bfb9292b06efd6ae653895f02efcd21a34f0b36952abd9b2b2"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_options — proxy_config.https_auto_cert.use_mtls.xfcc_options / c8938edc0305 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-a89b29ce5d64616933b19dfd941b49ab97fb290e7c1567333a820ab327fa5df6)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- proxy_config.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-542bdfff731d2a5567fec73c4c7edcb756a306b5434e6810f11e03459a45cd54"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-a26d410834812517d8185d987baf401845023efd8306c8df318d67a7074d5f6c"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.xfcc_options / c8938edc0305 / 3

<a id="canonical-560d757d9ba512ffef2fba84dee817d85506cc2bbbd37a5061e4d182745cc57e"></a>

<a id="canonical-84ff6d7a9ed5e98df37ab5798cc2f669e046e5960a641671589c182316938a1d"></a>

## xfcc_header_elements property — proxy_config.https_auto_cert.use_mtls.xfcc_options / c8938edc0305 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-d50ee5a727ecd0d785acdb1a6d3b7af1922693ba7416a6d486de5311d6b8aa6c"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.xfcc_options / c8938edc0305 / 5

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-9e17e562a25dda4678b1a8f2569c71871e06ea268a0077037f271ef6b95e7ad1)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-6dfd33c696868d9aaa2f87009aa602ef5a1be81177bdfd35ec1953db11750b92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c379969cb114c7b2deb09f646cabd5fdba91bbfdf33fa09a94c395953b890939"></a>

## timeouts — timeouts / b7f24cce38c3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- timeouts

<a id="canonical-585143a98318f2eadafe1239e1f547db2f40383501e9231fc5776ef9daf6bb62"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-8bad8e86efab2301f35082eded4eafdf2454530de94a98e935b0826996976a0d"></a>

## Direct properties — timeouts / b7f24cce38c3 / 3

<a id="canonical-c90ab0462223ee50768b2237c5915b4ce774beb9934b36f13104e329ecfe28c5"></a>

<a id="canonical-2391c3700087b0b94601c29f1927c19d856bae8b95d4b41030c7b0572c5f0fa9"></a>

## create property — timeouts / b7f24cce38c3 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-698e3207550fcc5aea203e60899bd579901393b3af3ea2d72a2f414a3dba6d56"></a>

<a id="canonical-285e3175a6cd6ec300a0582a9cb3ec97160ffd9913429a3d404aa8fdf757c2eb"></a>

## delete property — timeouts / b7f24cce38c3 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-d58a74198b860233a36e444435b067b08d5eba57b533c2385288effbce2e9147"></a>

<a id="canonical-7d527122e3b65e115ec5dda78c57c52b3fad7e0e51b9fa8b18db738ecf02d57b"></a>

## read property — timeouts / b7f24cce38c3 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-52d2b0c3ef27dd72e17522f01941e1c8633d220c0112184a5e67e26ad3ffc0f4"></a>

<a id="canonical-581d47c4bd899c5b1285f4fabb9d7e35e9043138c510f48910dd6a86a84dc05b"></a>

## update property — timeouts / b7f24cce38c3 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9b3159110cd126e919f3019b32c606a33d4522efd7227c34e863c75c4f62ecab"></a>

## Next pages — timeouts / b7f24cce38c3 / 8

- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
