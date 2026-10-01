---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-a54ef0c8b726270ba6da9f948aa3e9492180e3e845924ed221cd446ae0c7f69f"></a>

## Next pages — jwt_validation / 70f0895252ab / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f)
- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-8a0d75570e09c49fa8719c001a88ab35b0956a9973ae355867082dc1752dd674)
- [jwt_validation.jwks_config](resources--http_loadbalancer--reference--group-020.md#canonical-207912bbcabda8e22d01bf8ca2b79480e8c4733e7149ff9876e43c77f8d44f29)
- [jwt_validation.mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-c990ea81e10e45c0f4822e07ef17a4b567a240fe6e9ff10ac5dedf426ee2474c)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-9073b7d6843748237d154caa9a53d8d56c8367c9123bbaf8de6c42345c346f98)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b87c1c3a73435128f6d15eac09dbe8131c7c840df5e015836c3f30e1a5895db7"></a>

## jwt_validation.action — jwt_validation.action / 47f14786927b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.action

<a id="canonical-76b152879442aad0625bf7a8d1e942f3247767a7b0dc2860317dc37cf19f4568"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-77562dda621059999d5566c9d1ab2f21d5fb7594d0fd12900c454f107393a428"></a>

## Direct properties — jwt_validation.action / 47f14786927b / 3

- [block](resources--http_loadbalancer--reference--group-020.md#canonical-859767479caceecb5267ad13e9b0937ff703c9a2d08f6ec42d54a8bfc48b0fe8): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-020.md#canonical-df926452c66b807a23a88834ff0e510e653744aef77263818f859b51f610da31): complete subsection reference.

<a id="canonical-d0d416f80659079ae41c28f9d92098ba1f15bb89728ceadf87285503b62ab894"></a>

## Next pages — jwt_validation.action / 47f14786927b / 4

- [jwt_validation.action.block](resources--http_loadbalancer--reference--group-020.md#canonical-859767479caceecb5267ad13e9b0937ff703c9a2d08f6ec42d54a8bfc48b0fe8)
- [jwt_validation.action.report](resources--http_loadbalancer--reference--group-020.md#canonical-df926452c66b807a23a88834ff0e510e653744aef77263818f859b51f610da31)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-859767479caceecb5267ad13e9b0937ff703c9a2d08f6ec42d54a8bfc48b0fe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-014bba147c9e7a28ed1b6b118bf7cc1afb8ad5f254bddc036c77a7c9c9175b53"></a>

## jwt_validation.action.block — jwt_validation.action.block / e4eccc694aea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f)
- jwt_validation.action.block

<a id="canonical-769326d4bb90db8e860151702df46751bda0986c30877a3f5ee805aef0912538"></a>

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
block = {}
```

<a id="canonical-acd42fed42085e8e91a5bf9db06e34da1fa49ff22cc92ef340a9e7c3d3dd241f"></a>

## Direct properties — jwt_validation.action.block / e4eccc694aea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cda3632ae7508e54a1f2be48619fcb0e15f46bc2a811c7404e1b0c5904ccafe"></a>

## Next pages — jwt_validation.action.block / e4eccc694aea / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-df926452c66b807a23a88834ff0e510e653744aef77263818f859b51f610da31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a77d71e32478054870da16acc10264246f5c5ceea069417d7e35e2cf683167f"></a>

## jwt_validation.action.report — jwt_validation.action.report / 6bc547027901 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f)
- jwt_validation.action.report

<a id="canonical-54ff866e8581f9f89ec95e6356579d71dbc2d1f573a3e910f53f0e822aff9cfd"></a>

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
report = {}
```

<a id="canonical-235fef2329735a01f1d17d44ee03d63a1325c51b3797ca6d2515f084d30518c4"></a>

## Direct properties — jwt_validation.action.report / 6bc547027901 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9f32a19daa82c21c77d0da412b2a286190b9293686e8d792f4fa3ec7edcbb55"></a>

## Next pages — jwt_validation.action.report / 6bc547027901 / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8a0d75570e09c49fa8719c001a88ab35b0956a9973ae355867082dc1752dd674"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-195bb126d481090c53c782a87474c0846adc596b331afef2002fec35d8401a3c"></a>

## jwt_validation.authorization_server — jwt_validation.authorization_server / ca848c13985c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.authorization_server

<a id="canonical-b674f7e042529a2b383c0fa413e2487575ecaead3db92d99cf522d1b77361c48"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
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
authorization_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-4dd73ffa9eba9848d1384fb0013e7d99b3f171bc427c4f51ac7c151840af25d6"></a>

## Direct properties — jwt_validation.authorization_server / ca848c13985c / 3

- [authorization_servers](resources--http_loadbalancer--reference--group-020.md#canonical-b6dcbc2b2f9b141d01bb725dccd021197122d8eb79845d3b97971b9f512ddf1a): complete subsection reference.

<a id="canonical-02a4e5061fbe9fd8d9d24f64444bccb9732933f172b1499dff79ba6e2a13c55f"></a>

## Next pages — jwt_validation.authorization_server / ca848c13985c / 4

- [jwt_validation.authorization_server.authorization_servers](resources--http_loadbalancer--reference--group-020.md#canonical-b6dcbc2b2f9b141d01bb725dccd021197122d8eb79845d3b97971b9f512ddf1a)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b6dcbc2b2f9b141d01bb725dccd021197122d8eb79845d3b97971b9f512ddf1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22d838974c63bbf9bd74758ab68ed201e1b1d6ff688299aaeb4ebcc1978b6622"></a>

## jwt_validation.authorization_server.authorization_servers — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-8a0d75570e09c49fa8719c001a88ab35b0956a9973ae355867082dc1752dd674)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-5ddfdb465601431678a312af242c5a187ee2326332b5e374bffae397dbd0e033"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Upstream description:

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
authorization_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-15932d4203711a17514cdcf64e9520f48634aeac9880316099baa402e9a2888d"></a>

## Direct properties — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 3

<a id="canonical-3dd82b7866f52a801543f502e6b6b29aab30bb886d90e09ee0a338970ed6f2c5"></a>

<a id="canonical-c7e548890e83419c5eb028eb13db3fb1ca2a0afd577fc5d4c03b2b6c0bd46a27"></a>

## name property — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 4

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

<a id="canonical-f07032398901481a5b277139f85b3585c62b0334bbeed6e585781e2580f76a18"></a>

<a id="canonical-5de3aef602dc4186589baf7494150390280e4737645031bb6a5ce0df65b1c8f2"></a>

## namespace property — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 5

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

<a id="canonical-157d9ce561688644e432ad402e265a8b3f3b7fff786606f23549ebcf592c762e"></a>

<a id="canonical-74aad74731930d592f77a82162c27fbb5f366c2f3b9186c51a47f9b61575ce5f"></a>

## tenant property — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 6

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

<a id="canonical-5af3aebf42ac88d54782642cf2ddaa3fc64ff664cec48dde2788d080d5d787cb"></a>

## Next pages — jwt_validation.authorization_server.authorization_servers / 3646f3607031 / 7

- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-8a0d75570e09c49fa8719c001a88ab35b0956a9973ae355867082dc1752dd674)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-207912bbcabda8e22d01bf8ca2b79480e8c4733e7149ff9876e43c77f8d44f29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3c582e873b6f5927be0affa185ba37756e6236b62877a33989f0ed884fd27f0"></a>

## jwt_validation.jwks_config — jwt_validation.jwks_config / e26f79d47553 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.jwks_config

<a id="canonical-2686ccad24c04384fac0859d7739f1495955dc2503dceb743514d788230dd283"></a>

Type: `"object"`. single nested block, Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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
jwks_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-a88cdb06aef0f9bd936caf2ce1c360b1d1fa03217aa028a66b1fb7252a297cb5"></a>

## Direct properties — jwt_validation.jwks_config / e26f79d47553 / 3

<a id="canonical-bb9d1910a0c41c14d5acb38377f3b071f37464ef8c2698ab7694c3e16fc95b9c"></a>

<a id="canonical-02e375a65974550a091f6f3d347cf199e632d28820df0b3fcac9cace00895acf"></a>

## cleartext property — jwt_validation.jwks_config / e26f79d47553 / 4

Type: `"string"`. Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="canonical-321509324e24f4265eacb3e6d1844c10eefd20c4c5f3e32420d92eb7baad2993"></a>

## Next pages — jwt_validation.jwks_config / e26f79d47553 / 5

- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c990ea81e10e45c0f4822e07ef17a4b567a240fe6e9ff10ac5dedf426ee2474c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-009ed1a9d0ad2511564e5063c49ce3d2b694e8321bdc59b35d4b13c90a75eb62"></a>

## jwt_validation.mandatory_claims — jwt_validation.mandatory_claims / 1bf225b4e6d5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.mandatory_claims

<a id="canonical-5b53ccbdf9c89f0b460da85121fecec9ab0bac3dc6ed791a63b20a4f9023fd9c"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of mandatory Claims.

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
mandatory_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-c2331dcf59e2466203718076a8a237b029f0ecf070e7c733e8043dc94e52b689"></a>

## Direct properties — jwt_validation.mandatory_claims / 1bf225b4e6d5 / 3

<a id="canonical-39b76f22a64435508cfa199338ad18ceb5279be19429673dbe7fe7d2b8e2b728"></a>

<a id="canonical-00d470770b61936171f86ea596028874f1e58400620db9f0f248e7d792710a80"></a>

## claim_names property — jwt_validation.mandatory_claims / 1bf225b4e6d5 / 4

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-8a6c08e783e6f8d15551318e2cf8ba6be82399708c9f8859db19818ef2cdc03d"></a>

## Next pages — jwt_validation.mandatory_claims / 1bf225b4e6d5 / 5

- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e2cee668637b962c57c3587375e520555c0e2061fc259818a60baeabfd3ad94"></a>

## jwt_validation.reserved_claims — jwt_validation.reserved_claims / a69dd371986f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.reserved_claims

<a id="canonical-863519e2dd4c39628dfc08e84689c5a328e9e0e6fe9a272d3510ffd999e8884e"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("audience",
    "audience_disable"),
  validators.ConflictingObjectAttributes("issuer",
    "issuer_disable"),
  validators.ConflictingObjectAttributes("validate_period_disable",
    "validate_period_enable")}
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
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

Terraform syntax:

```terraform
reserved_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-4be44a43474295f916e4b96b23aa9fc934e7a4a3eb9cab5a17c6c51f0525f777"></a>

## Direct properties — jwt_validation.reserved_claims / a69dd371986f / 3

- [audience](resources--http_loadbalancer--reference--group-020.md#canonical-b0cfd01ad9463d3d511021c5e6d358ac58dba7fe130945288541a39eaa27e274): complete subsection reference.

- [audience_disable](resources--http_loadbalancer--reference--group-020.md#canonical-ab9a81ecd5e53de82a698f8f04180c04ae4e1f062680924fa7a220fc04b2712a): complete subsection reference.

<a id="canonical-1fc080b416e8f7b4fb62dd7773210373068d6e3482fb785d704101986445b8f4"></a>

<a id="canonical-e5764d36244b2cb7d0795df4211491186ad6919c9411b97abc15d8e6c25d9018"></a>

## issuer property — jwt_validation.reserved_claims / a69dd371986f / 4

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

Upstream description:

Exclusive with \[issuer\_disable\]

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

- [issuer_disable](resources--http_loadbalancer--reference--group-020.md#canonical-08e745012d9bae4ba4194513a8864b3741f392f71b24defd1df0aa1306755002): complete subsection reference.

- [validate_period_disable](resources--http_loadbalancer--reference--group-020.md#canonical-66d258fa9e5f89e40e4eea800359eae6d517672b5a9940c04b148d680eff1168): complete subsection reference.

- [validate_period_enable](resources--http_loadbalancer--reference--group-020.md#canonical-4bc23067ba6a0ef07c6f36b7100e0714f0e5084e9b4cf6533a6a73d2660c732f): complete subsection reference.

<a id="canonical-fec2fea2f729a4285212ee07466b08a8d3cf39d9a253d6279d2e23faa1d7f874"></a>

## Next pages — jwt_validation.reserved_claims / a69dd371986f / 5

- [jwt_validation.reserved_claims.audience](resources--http_loadbalancer--reference--group-020.md#canonical-b0cfd01ad9463d3d511021c5e6d358ac58dba7fe130945288541a39eaa27e274)
- [jwt_validation.reserved_claims.audience_disable](resources--http_loadbalancer--reference--group-020.md#canonical-ab9a81ecd5e53de82a698f8f04180c04ae4e1f062680924fa7a220fc04b2712a)
- [jwt_validation.reserved_claims.issuer_disable](resources--http_loadbalancer--reference--group-020.md#canonical-08e745012d9bae4ba4194513a8864b3741f392f71b24defd1df0aa1306755002)
- [jwt_validation.reserved_claims.validate_period_disable](resources--http_loadbalancer--reference--group-020.md#canonical-66d258fa9e5f89e40e4eea800359eae6d517672b5a9940c04b148d680eff1168)
- [jwt_validation.reserved_claims.validate_period_enable](resources--http_loadbalancer--reference--group-020.md#canonical-4bc23067ba6a0ef07c6f36b7100e0714f0e5084e9b4cf6533a6a73d2660c732f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b0cfd01ad9463d3d511021c5e6d358ac58dba7fe130945288541a39eaa27e274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32cb9804189dc420a43042a1f07da8035ee898f67acfef1db8a7dd8f54e459eb"></a>

## jwt_validation.reserved_claims.audience — jwt_validation.reserved_claims.audience / de3dfb664b93 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- jwt_validation.reserved_claims.audience

<a id="canonical-33a5b99179da493ca60dc2f96405a402bb2e3f6b8d5a427abd779707e5478dfb"></a>

Type: `"object"`. single nested block, Optional.

Audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("audiences")}
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
audience {
  # Configure direct properties listed below.
}
```

<a id="canonical-471a1a0589da2134cb18a636c711c83b045d1890485be41dc4ce5c89c505803f"></a>

## Direct properties — jwt_validation.reserved_claims.audience / de3dfb664b93 / 3

<a id="canonical-8a35832f44de81c6d41d126d4dfc397f37bab88f59b1e19d2630558e14715081"></a>

<a id="canonical-2a62e3a92b0d90889d7c43e415238dcb08de57d64289502a219f59c00fb0b7f5"></a>

## audiences property — jwt_validation.reserved_claims.audience / de3dfb664b93 / 4

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-30335ebcc416a5af4f5bcdb4dd306dd1be12f139c428a13fc02da7293a7ecb50"></a>

## Next pages — jwt_validation.reserved_claims.audience / de3dfb664b93 / 5

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ab9a81ecd5e53de82a698f8f04180c04ae4e1f062680924fa7a220fc04b2712a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67072ec0b3e8b44afce70d89fc358ae649b5479bcd63ebc38de18ec56c4704c0"></a>

## jwt_validation.reserved_claims.audience_disable — jwt_validation.reserved_claims.audience_disable / 0c4edb05cf7b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-e69359e11a7158f6a2c89396b72ace14fff1d6d522ed5834b05e7c7c9ef982c1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

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
audience_disable = {}
```

<a id="canonical-624210c738d6adfc5980e379a4eaaab67dc67bcedf21fa012100503eeaeb42a3"></a>

## Direct properties — jwt_validation.reserved_claims.audience_disable / 0c4edb05cf7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac21a30f8bfe8f0704aa1f9c6d82bbf9d6b304f495740919ab0eaf6b1a06b588"></a>

## Next pages — jwt_validation.reserved_claims.audience_disable / 0c4edb05cf7b / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-08e745012d9bae4ba4194513a8864b3741f392f71b24defd1df0aa1306755002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b98d7101c542e8523f1abd59db1d5fd4f1017408c558c17bdf5acb9709cdeabe"></a>

## jwt_validation.reserved_claims.issuer_disable — jwt_validation.reserved_claims.issuer_disable / ad39421f22fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-64ea38d4bc067c0d217b26aea9ea296f3028cb1d8bb03d9e06ec3206194e6f2b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

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
issuer_disable = {}
```

<a id="canonical-b7b9f16c8f60b706d1cf6e2c56f8b118fd0cdaa106a530f6f402c8be6048de84"></a>

## Direct properties — jwt_validation.reserved_claims.issuer_disable / ad39421f22fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8375839d360049fa8e810c86a896eb615af4c392057e4df61865b13d3befa0b"></a>

## Next pages — jwt_validation.reserved_claims.issuer_disable / ad39421f22fd / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-66d258fa9e5f89e40e4eea800359eae6d517672b5a9940c04b148d680eff1168"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6917a181f9e3d1c8aa2d102655964bef8a5b1731356b5f8e42b38dd8b8e7cbfd"></a>

## jwt_validation.reserved_claims.validate_period_disable — jwt_validation.reserved_claims.validate_period_disable / 334f0c13da95 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-17ba75d443bbcdf6d3434a13c5bdd10fb3bd12107960e1f7c702254f2e23afc6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

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
validate_period_disable = {}
```

<a id="canonical-c6ab85ce52dc4f2701c5e70e66843bdb9086306e010b1fdf76c3cf21902913c1"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_disable / 334f0c13da95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9bf0ea2bdbf581b44096c736dc04f9a72aaa13c69638444d9afafb25959ef3c7"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_disable / 334f0c13da95 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4bc23067ba6a0ef07c6f36b7100e0714f0e5084e9b4cf6533a6a73d2660c732f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bbbd8108f453af135ab88b77dd8660e5f083a456081b154b349bdd673524def"></a>

## jwt_validation.reserved_claims.validate_period_enable — jwt_validation.reserved_claims.validate_period_enable / 064abcf51b27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-09213e9709babefbee76794b1d373df97226c819f7c38fea20562f2bbee80bb4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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
validate_period_enable = {}
```

<a id="canonical-4e8b06b0282b317d9601a963f6e2a7858370ea827a63b7db5b6e17282c389784"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_enable / 064abcf51b27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-919dbd7dea75d0f9c02a8b9ade2c6801d334b7d41ad45e9f4a512693aba960ee"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_enable / 064abcf51b27 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0097f055d3a25eb73c453f356b436557a2f0b1620daf5c0a8c64a9ec106517be"></a>

## jwt_validation.target — jwt_validation.target / 78654ef9b696 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.target

<a id="canonical-fc17188d07f8657b9213ecc5b7779d729d716a395e118d574a1c27b532b3139e"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_endpoint",
    "api_groups"),
  validators.ConflictingObjectAttributes("all_endpoint",
    "base_paths"),
  validators.ConflictingObjectAttributes("api_groups",
    "base_paths")}
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
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

Terraform syntax:

```terraform
target {
  # Configure direct properties listed below.
}
```

<a id="canonical-87a7bd77ff2477473e8e860db4e72452579886f682f75d51bb7ef8bc8a7eb5eb"></a>

## Direct properties — jwt_validation.target / 78654ef9b696 / 3

- [all_endpoint](resources--http_loadbalancer--reference--group-020.md#canonical-48e0d2f9388b762348fdc34f2d60f34ac1e851de4d77de61bf66feaccfe87865): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-020.md#canonical-6a58f5d326ad3072fa29f0ebbb8a9d1c9f2bf575b8db3c048b705bd1df996452): complete subsection reference.

- [base_paths](resources--http_loadbalancer--reference--group-020.md#canonical-658379688c2a42d83c8c1f9928d5b07b96dfbd6f879071ece205e09a6fba6e94): complete subsection reference.

<a id="canonical-377ef9b029c93d2e883fa592c8fdf6e4ed579b4e0ab9d7ca95699e02ab22d4fd"></a>

## Next pages — jwt_validation.target / 78654ef9b696 / 4

- [jwt_validation.target.all_endpoint](resources--http_loadbalancer--reference--group-020.md#canonical-48e0d2f9388b762348fdc34f2d60f34ac1e851de4d77de61bf66feaccfe87865)
- [jwt_validation.target.api_groups](resources--http_loadbalancer--reference--group-020.md#canonical-6a58f5d326ad3072fa29f0ebbb8a9d1c9f2bf575b8db3c048b705bd1df996452)
- [jwt_validation.target.base_paths](resources--http_loadbalancer--reference--group-020.md#canonical-658379688c2a42d83c8c1f9928d5b07b96dfbd6f879071ece205e09a6fba6e94)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-48e0d2f9388b762348fdc34f2d60f34ac1e851de4d77de61bf66feaccfe87865"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c472f778032cd53906c8cb703e900065085ac1eaa57ab57f711aa4fbcb3fa89a"></a>

## jwt_validation.target.all_endpoint — jwt_validation.target.all_endpoint / a10fca17b9f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- jwt_validation.target.all_endpoint

<a id="canonical-c960b6b9279057bc0a823ab1ff51923544761b8cd5cb91f2f60eb2dd387858bd"></a>

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
all_endpoint = {}
```

<a id="canonical-8195c5b4aa6e8a70e17515faf1f283da15519ef132c4ca3183146a6c24dc78c1"></a>

## Direct properties — jwt_validation.target.all_endpoint / a10fca17b9f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b315312b5c5b555778c0a96295a58d182c2a47bc62ccb5d01f75d10e35fb2fd"></a>

## Next pages — jwt_validation.target.all_endpoint / a10fca17b9f7 / 4

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6a58f5d326ad3072fa29f0ebbb8a9d1c9f2bf575b8db3c048b705bd1df996452"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f828de1b89d767f4a281fdc024137960f0f6279af8443d8631daa5bfb97d21e"></a>

## jwt_validation.target.api_groups — jwt_validation.target.api_groups / 69d7b8e2fd30 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- jwt_validation.target.api_groups

<a id="canonical-f5bd1239c2ad033769d7ea271c6a832c0f67907343c2e9eb88cc836e9d5abced"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331f550c8700af9c37853dcd84b4b4cd592d3479906f04e9979394a2bacec58"></a>

## Direct properties — jwt_validation.target.api_groups / 69d7b8e2fd30 / 3

<a id="canonical-213276d227d8e6e290650f91812260dfdc5f77bec37be8b274e44e160999f051"></a>

<a id="canonical-d376f2269b2d10e296104cb52e3fccd4ec81eeb9f1411cc2e92118b1f1efe8a8"></a>

## api_groups property — jwt_validation.target.api_groups / 69d7b8e2fd30 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ae48f48a0e20680e74b09eafd22ebe39f46d6e0a658ef486c417a12aca6af26a"></a>

## Next pages — jwt_validation.target.api_groups / 69d7b8e2fd30 / 5

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-658379688c2a42d83c8c1f9928d5b07b96dfbd6f879071ece205e09a6fba6e94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-696b2dda6ac747c2b3ea71ba71381f007098f6ebdcf9b9009a90161182026b64"></a>

## jwt_validation.target.base_paths — jwt_validation.target.base_paths / 289995bf33e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- jwt_validation.target.base_paths

<a id="canonical-c071a486e466a2cba085d8af069f487e32e786509394087e17a5f948cf757c6f"></a>

Type: `"object"`. single nested block, Optional.

Base Paths.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("base_paths")}
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
base_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-201d4172d0ff2b998bd0edc6d019477ba5acd40d41bed8d70eda27d636151a4a"></a>

## Direct properties — jwt_validation.target.base_paths / 289995bf33e2 / 3

<a id="canonical-eb963cc4046f7dcfffa41455016936f4af48b75d25ff511facc522c94ad74cb1"></a>

<a id="canonical-1edc74c0ccc45d290b317ee321a9e524cd49a5716234f0058a9a48f5e22d44eb"></a>

## base_paths property — jwt_validation.target.base_paths / 289995bf33e2 / 4

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-be1cd03dae7e149d2403f52d4499d9dd6652c43a687c1183b638f067166dbb63"></a>

## Next pages — jwt_validation.target.base_paths / 289995bf33e2 / 5

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9073b7d6843748237d154caa9a53d8d56c8367c9123bbaf8de6c42345c346f98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a00c0935434911f4e1dc5a184e0d31238bcd5ac80f126ec350d20daaa54bd4"></a>

## jwt_validation.token_location — jwt_validation.token_location / a4603c838162 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- jwt_validation.token_location

<a id="canonical-14694448a72a7468c125ee6151405533d42e8994f2bcde0c1f802ae3ad78bc2c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a21d94818f2bb80bd709fc0b59bbee193920f301c1c54ce23370fb9a39dea53"></a>

## Direct properties — jwt_validation.token_location / a4603c838162 / 3

- [bearer_token](resources--http_loadbalancer--reference--group-020.md#canonical-f961034ffb9a09a01b2cdf0363b4775a85c4042ddd97f101c4080f67d4530f72): complete subsection reference.

<a id="canonical-ad8c98765198ca338f4909a4f08f8767916f79d2af45a9bce46061dcb2bf36d9"></a>

## Next pages — jwt_validation.token_location / a4603c838162 / 4

- [jwt_validation.token_location.bearer_token](resources--http_loadbalancer--reference--group-020.md#canonical-f961034ffb9a09a01b2cdf0363b4775a85c4042ddd97f101c4080f67d4530f72)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f961034ffb9a09a01b2cdf0363b4775a85c4042ddd97f101c4080f67d4530f72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6817743b87ff05dedb022c9334c6b1128cc550932c4f587b5c30082a9c02663"></a>

## jwt_validation.token_location.bearer_token — jwt_validation.token_location.bearer_token / 417b67f6de8c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff)
- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-9073b7d6843748237d154caa9a53d8d56c8367c9123bbaf8de6c42345c346f98)
- jwt_validation.token_location.bearer_token

<a id="canonical-25a633b966542d0f41a142a541f2c750884d8ed748d70823e702caaa4cf77b46"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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
bearer_token {}
```

<a id="canonical-2bf138e6f8afdb6357aba05a042cd24773e4e48800878a5beb2f2eb030bae271"></a>

## Direct properties — jwt_validation.token_location.bearer_token / 417b67f6de8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01bb20054d46fffaae9f4c882b6cf09660c2ce79f8515d2a93e8384184ec9954"></a>

## Next pages — jwt_validation.token_location.bearer_token / 417b67f6de8c / 4

- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-9073b7d6843748237d154caa9a53d8d56c8367c9123bbaf8de6c42345c346f98)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7851e2010afcdabe3e06b86c00ca2119a78dc076e9788767f1d31b1ba8a5e3f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-533c5f984ae4359a563117ab5a0bdfc153cb4b1fd4ea36e2fd1403f7ea9e7785"></a>

## l7_ddos_action_block — l7_ddos_action_block / ae5d57a194e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- l7_ddos_action_block

<a id="canonical-0fb9da1cfac3002b7b44eab295e41ab1369b30dcc2609ec5b800d659dc7f5722"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](resources--http_loadbalancer--reference--group-020.md#canonical-0fb9da1cfac3002b7b44eab295e41ab1369b30dcc2609ec5b800d659dc7f5722)
- [l7_ddos_action_default](resources--http_loadbalancer--reference--group-020.md#canonical-34aa9a380c9c3ac60358d0d0c0c362d5450d31b5eda2d9e9a8b9cd1a184ce358)
- [l7_ddos_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-e19c33bff160d3281df7efee5fd1c68742257df00a7e0f66ae1c26c861df783a)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

<a id="canonical-ef1ae7a2ad9c27478bd133a69726c89426bbe92281263ce8401ffb4ffc91e378"></a>

## Direct properties — l7_ddos_action_block / ae5d57a194e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0c9d9021dc043164dc8acf2d52491828b824447b4cee01f5248c2f9dd93887e"></a>

## Next pages — l7_ddos_action_block / ae5d57a194e6 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dacd16fe43e99c16d648b347417ec881594427f7192ce07760ccb0a434933ebf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a4a8074f2091d5db84b7c83165c05a9999df2a1b34b16116ea7a1cb5711175d"></a>

## l7_ddos_action_default — l7_ddos_action_default / d3642b883606 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- l7_ddos_action_default

<a id="canonical-34aa9a380c9c3ac60358d0d0c0c362d5450d31b5eda2d9e9a8b9cd1a184ce358"></a>

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
l7_ddos_action_default = {}
```

<a id="canonical-d792f4e7c08b0c44de12323664addeb5273eba5fb689ae06266a866bc61385a8"></a>

## Direct properties — l7_ddos_action_default / d3642b883606 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e58f8f40838f84825b8442f4bc479e633f94204d8e3d08dbe8231485c72e67a2"></a>

## Next pages — l7_ddos_action_default / d3642b883606 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5c2a0de006a63bfb8d7a46ca1a9ddaeb0faf0bc4cc2c75baf932466b1c4b9a8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b76853b97ed2928a3cca65b221fbc5cdae7f732b4c5c320668fa21625b39a1b3"></a>

## l7_ddos_action_js_challenge — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- l7_ddos_action_js_challenge

<a id="canonical-e19c33bff160d3281df7efee5fd1c68742257df00a7e0f66ae1c26c861df783a"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
l7_ddos_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-29d14dba1dc826d43974ad3636f64b89e72b401b6a48f3502e7ea2b412b66ce8"></a>

## Direct properties — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 3

<a id="canonical-dd2cc109a4cf2ff30d3455b103963080e13634548973490294aa344a2e9819d4"></a>

<a id="canonical-d1638f2362e9e39b1b77396c0859735e9c8ef06889444920cc43d225c51e719b"></a>

## cookie_expiry property — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-afb0cd994786e7df16763d42cc8240117d1ffb83cba34e1671f999f5772475a7"></a>

<a id="canonical-292e4d13071315fbd80db6d1420409c5f0f45fcf626cc0bf496b2a776d2ce8f1"></a>

## custom_page property — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4729467e8595811bd33a32bc22bc03d67b33004f4a12d7fcc49b52eed69bbcd4"></a>

<a id="canonical-62e56d354440dc4338cb74c515fa3bb07e9fda2a39be59b0e90762827430fc49"></a>

## js_script_delay property — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-fdb7e88fe2bd6b0ded4adcb56f5a2598628fa6d7cb379791415f814cc4ac0be8"></a>

## Next pages — l7_ddos_action_js_challenge / 9e2d8f05bdb4 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66ac7bce0842dc1ae89fddf301162d418688bd22bcfe89743e607ab026f23f9b"></a>

## l7_ddos_protection — l7_ddos_protection / 36ba8d68fa2c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- l7_ddos_protection

<a id="canonical-d4c9b4fd520f51592abae80c6e87b4ba80c3ab5eb3665e5081eae6c9b6ecf1b4"></a>

Type: `"object"`. single nested block, Optional.

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack. Defaults to \`map\[\]\`. Server
applies default when omitted.

Upstream description:

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_js_challenge"),
  validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("clientside_action_js_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("ddos_policy_custom",
    "ddos_policy_none"),
  validators.ConflictingObjectAttributes("default_rps_threshold",
    "rps_threshold"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_captcha_challenge"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_js_challenge"),
  validators.ConflictingObjectAttributes("mitigation_captcha_challenge",
    "mitigation_js_challenge")}
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
  "x-ves-oneof-field-clientside_action_choice": "[\"clientside_action_captcha_challenge\",\"clientside_action_js_challenge\",\"clientside_action_none\"]",
  "x-ves-oneof-field-ddos_policy_choice": "[\"ddos_policy_custom\",\"ddos_policy_none\"]",
  "x-ves-oneof-field-mitigation_action_choice": "[\"mitigation_block\",\"mitigation_captcha_challenge\",\"mitigation_js_challenge\"]",
  "x-ves-oneof-field-rps_threshold_choice": "[\"default_rps_threshold\",\"rps_threshold\"]"
}
```

Terraform syntax:

```terraform
l7_ddos_protection {
  # Configure direct properties listed below.
}
```

<a id="canonical-a624d1b7f79ebe336117509ec396a2a163d7fe7f122b067c9fe94ae5c163dac0"></a>

## Direct properties — l7_ddos_protection / 36ba8d68fa2c / 3

- [clientside_action_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-9ee9c79d290a45623691a814895290dc60f0dc84e932555d7fc5987e5a66295e): complete subsection reference.

- [clientside_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-3f63200e95108a27c1a90262f355a32bf11afb6b8bcf03137f4dff1087b4209b): complete subsection reference.

- [clientside_action_none](resources--http_loadbalancer--reference--group-020.md#canonical-2e20b9e50419f59fe1b9e4551ad2e988bed907e1c218066bbfd758e4a010df2e): complete subsection reference.

- [ddos_policy_custom](resources--http_loadbalancer--reference--group-020.md#canonical-8ecbd7a40b8da2bacef3ccc188ae23f08c57af013683a5e5c03f15e3f3099e52): complete subsection reference.

- [ddos_policy_none](resources--http_loadbalancer--reference--group-020.md#canonical-b011ef3a930c0e82c9d9621db300e9cfc7f669a2278d05c3a15bd21e309e0cb9): complete subsection reference.

- [default_rps_threshold](resources--http_loadbalancer--reference--group-020.md#canonical-31384f2f8e1e4e1ce447baae06278cbc2bf6d1b89de3e950814e0e677cc82492): complete subsection reference.

- [mitigation_block](resources--http_loadbalancer--reference--group-020.md#canonical-d2839e5778a0ac3e498951739f5e9f5642bf5308c0ccdc325a91d58755126790): complete subsection reference.

- [mitigation_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-b5d93240ad177bf7003b1ca054b369dbcc5c10575302c6b1d2a0c7cb5318e72e): complete subsection reference.

- [mitigation_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-50f1da05abef36adc0c6409a02ef2fe0aaa5720f36103b73042206ccdc3dea66): complete subsection reference.

<a id="canonical-82f26ee7105a7b45db745cf7b436c20b055fd3ba16a4dc6043724b363239434f"></a>

<a id="canonical-35794e73a294c7e0259ff9067df8c9d51e55b43f579f064ad749e13185ba89b8"></a>

## rps_threshold property — l7_ddos_protection / 36ba8d68fa2c / 4

Type: `"number"`. Optional.

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Upstream description:

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 50000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 50000,
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
    "ves.io.schema.rules.uint32.lte": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  }
}
```

<a id="canonical-a70669c5f5198e41c76cfc29261538e6b366f083680993b86d52a3bef142f924"></a>

## Next pages — l7_ddos_protection / 36ba8d68fa2c / 5

- [l7_ddos_protection.clientside_action_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-9ee9c79d290a45623691a814895290dc60f0dc84e932555d7fc5987e5a66295e)
- [l7_ddos_protection.clientside_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-3f63200e95108a27c1a90262f355a32bf11afb6b8bcf03137f4dff1087b4209b)
- [l7_ddos_protection.clientside_action_none](resources--http_loadbalancer--reference--group-020.md#canonical-2e20b9e50419f59fe1b9e4551ad2e988bed907e1c218066bbfd758e4a010df2e)
- [l7_ddos_protection.ddos_policy_custom](resources--http_loadbalancer--reference--group-020.md#canonical-8ecbd7a40b8da2bacef3ccc188ae23f08c57af013683a5e5c03f15e3f3099e52)
- [l7_ddos_protection.ddos_policy_none](resources--http_loadbalancer--reference--group-020.md#canonical-b011ef3a930c0e82c9d9621db300e9cfc7f669a2278d05c3a15bd21e309e0cb9)
- [l7_ddos_protection.default_rps_threshold](resources--http_loadbalancer--reference--group-020.md#canonical-31384f2f8e1e4e1ce447baae06278cbc2bf6d1b89de3e950814e0e677cc82492)
- [l7_ddos_protection.mitigation_block](resources--http_loadbalancer--reference--group-020.md#canonical-d2839e5778a0ac3e498951739f5e9f5642bf5308c0ccdc325a91d58755126790)
- [l7_ddos_protection.mitigation_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-b5d93240ad177bf7003b1ca054b369dbcc5c10575302c6b1d2a0c7cb5318e72e)
- [l7_ddos_protection.mitigation_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-50f1da05abef36adc0c6409a02ef2fe0aaa5720f36103b73042206ccdc3dea66)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9ee9c79d290a45623691a814895290dc60f0dc84e932555d7fc5987e5a66295e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f0ec3f78246fca08dca924cf51db9e0f8400ec2a55c1d58aa5506ead583ab12"></a>

## l7_ddos_protection.clientside_action_captcha_challenge — l7_ddos_protection.clientside_action_captcha_challenge / 251a79e6ce20 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.clientside_action_captcha_challenge

<a id="canonical-8b913dede49e8eb345187411006cb001af0d7c559fc0c8b1847f66d482240b7d"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
clientside_action_captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-e911d3d097dd93245d5595f448bd2d67d456706d1e425f5969eb68baf0d73f92"></a>

## Direct properties — l7_ddos_protection.clientside_action_captcha_challenge / 251a79e6ce20 / 3

<a id="canonical-3ec064d33a67bffb2e39a0c1a25d8df1f3e94d83e9f3037c35a87fac2e60bd0e"></a>

<a id="canonical-36f2a436a57ebf126fb969c464b1cab67fdf62ef4320e02d3da28a7ec8dd49bd"></a>

## cookie_expiry property — l7_ddos_protection.clientside_action_captcha_challenge / 251a79e6ce20 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-f119e2c28dddf2b0053af2ca8afb423f3f0ba38e1ee79477ea80080b7f2d1406"></a>

<a id="canonical-4b350cb447bda072b2a3c18a69739982c63b8b1e5985d8a7741109cabaa71de9"></a>

## custom_page property — l7_ddos_protection.clientside_action_captcha_challenge / 251a79e6ce20 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-9589106790ca395c17f8e69ff8813e53b92af0dbc567a65756352da538e0a242"></a>

## Next pages — l7_ddos_protection.clientside_action_captcha_challenge / 251a79e6ce20 / 6

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3f63200e95108a27c1a90262f355a32bf11afb6b8bcf03137f4dff1087b4209b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8892c39c0d1bdd757c28da4d864fe53e740665fc5bc3843e3c188a35bb5303f9"></a>

## l7_ddos_protection.clientside_action_js_challenge — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.clientside_action_js_challenge

<a id="canonical-426fd17e18bdbf9c5ec5067c3aee9f7db8c89c27356d266b48b965a7757930fb"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
clientside_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-eabec3fe52a6d1f5d875cfe226d1f35d0ba825d44bdc4ae88d093de038b7df69"></a>

## Direct properties — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 3

<a id="canonical-5cf30dc0aa6ba14930292dabfb432333a71a8098d8c989a74cc91da893dd6601"></a>

<a id="canonical-f0c1b0339557fc2e93778529140c27cad69ade789af3e99a232bdb1bbaa7b430"></a>

## cookie_expiry property — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-29a1917042dd1c19254aae6721b8ce4041d4604c678f98eacfb0f0c6b5f5af75"></a>

<a id="canonical-2e8335148decf2e878344dab3dae2179dc199356f845e7e59bbc70e6ea982081"></a>

## custom_page property — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-d87fddb1c4f07f1f9b7162fd28355b523b383c1ccae19839a043310b8465bb66"></a>

<a id="canonical-f90cb79faffb144e8475c7e2a970a916c65868133b4526f5a7c87eb814d8d2b9"></a>

## js_script_delay property — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-856f79b128110ef86d03ed0d9388b638debe38df5d983ab3f45d65661afaa24d"></a>

## Next pages — l7_ddos_protection.clientside_action_js_challenge / 8d9b3637303a / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2e20b9e50419f59fe1b9e4551ad2e988bed907e1c218066bbfd758e4a010df2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-783b5a365fa1a7791c87997043c495b6639b198a1451f74cd72981d81a21076c"></a>

## l7_ddos_protection.clientside_action_none — l7_ddos_protection.clientside_action_none / 64271029ea80 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.clientside_action_none

<a id="canonical-aecf4a2b990b0a48eb539ad988d5cd1437960badee7f85e6f6b638439de5eabf"></a>

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
clientside_action_none = {}
```

<a id="canonical-86338ddad59de54fd5a6a95c7128c466eb31695b6809dc71a3ba1022990f2142"></a>

## Direct properties — l7_ddos_protection.clientside_action_none / 64271029ea80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97c37aae067e11e833067f3faec86aa2d129804a3907e823ccbedf27dca74c39"></a>

## Next pages — l7_ddos_protection.clientside_action_none / 64271029ea80 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8ecbd7a40b8da2bacef3ccc188ae23f08c57af013683a5e5c03f15e3f3099e52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-834e54494f32a70b7e62a188153ef45223c944d95e6dac4b1bf53aa1a3d210b5"></a>

## l7_ddos_protection.ddos_policy_custom — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.ddos_policy_custom

<a id="canonical-8b5bfe66b69828e4fceaf90b15c2ab27dcfc2712aee8bcdfb75addd734a707b7"></a>

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
ddos_policy_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-60a0ff2b2382894e91a76abdc0aad59fc153c938d67c384c52613e25e552f3d0"></a>

## Direct properties — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 3

<a id="canonical-3c35f5b2c0a204c0d4b76c0c481809d2111effd771ea4b2f8e3c178925b19df4"></a>

<a id="canonical-0d0509d87201fdc72811508ae13c90296dd3a767d0affc3b55faf51d63ba481a"></a>

## name property — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 4

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

<a id="canonical-930ad6dc67b1100f02716a669406e950fff95083a2114f7ee3bd6bd716d89b0c"></a>

<a id="canonical-33fe21b483e51f98a7bb1a857fbe3031abfab68470160c2cb8f6a040318fa1d7"></a>

## namespace property — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 5

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

<a id="canonical-b7b4e392eff6ede3ea8f9b768a8c54667bf81eb45c136cb5ad05aa36246c15c4"></a>

<a id="canonical-d5b64ef4e5dda896b9b1e8a3cb0d359760e10f00caa6dda57947555a3d28fb0f"></a>

## tenant property — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 6

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

<a id="canonical-cee58135a758bb8591c7e106de99755bc69fc5f5a271037c6b50a1ab8a5a2d0f"></a>

## Next pages — l7_ddos_protection.ddos_policy_custom / 70e0959c5b2a / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b011ef3a930c0e82c9d9621db300e9cfc7f669a2278d05c3a15bd21e309e0cb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fee2e5b48a03c8ea8864d317f32d7c626af9a44af9b72c8d426a23e868a82ed"></a>

## l7_ddos_protection.ddos_policy_none — l7_ddos_protection.ddos_policy_none / baa0f85d3f3b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.ddos_policy_none

<a id="canonical-4c96870923d42ef3ee2025be116f994601801ccecf89a4e5addc6e8846aae9c9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ddos policy none.

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
ddos_policy_none = {}
```

<a id="canonical-3b88459a75484ec7f0ff7a64521a371f976b0a5cf4016d94959609659995a6e8"></a>

## Direct properties — l7_ddos_protection.ddos_policy_none / baa0f85d3f3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd2d9080e765fa622cfdab08f86774be28e104b51a524e81563a7dd0de8caf02"></a>

## Next pages — l7_ddos_protection.ddos_policy_none / baa0f85d3f3b / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-31384f2f8e1e4e1ce447baae06278cbc2bf6d1b89de3e950814e0e677cc82492"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35cb3656947e67d2c85a656911e4c3f8221057fe715d0359dffb6ea14e34a54b"></a>

## l7_ddos_protection.default_rps_threshold — l7_ddos_protection.default_rps_threshold / cf9e15737845 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.default_rps_threshold

<a id="canonical-6abdb732f6dd1443ceea7af0f70adb7505cf09f3d8173e5dede89f087b68b075"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default rps threshold.

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
default_rps_threshold = {}
```

<a id="canonical-cd5102a4736af7ab2c32fe9252e28cc0eb50f817651c2736a1ba5b59388f5f90"></a>

## Direct properties — l7_ddos_protection.default_rps_threshold / cf9e15737845 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef83f5b65896266f05595b02832ddb094fe4259fc31094154d9cef9e6442dd12"></a>

## Next pages — l7_ddos_protection.default_rps_threshold / cf9e15737845 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d2839e5778a0ac3e498951739f5e9f5642bf5308c0ccdc325a91d58755126790"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ecd474d6c914284dda899589de4aefc72a5ffaf0f5445f9462b468f23202303"></a>

## l7_ddos_protection.mitigation_block — l7_ddos_protection.mitigation_block / 4b79995f1368 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.mitigation_block

<a id="canonical-72bb4498f9241aedb000d71be2cddb9f84f8541c27d72e693bb3bde6936edb1f"></a>

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
mitigation_block = {}
```

<a id="canonical-6224239f74a93808882e6e02f635ad7c749ab8a2812c17d40f3139a452bece55"></a>

## Direct properties — l7_ddos_protection.mitigation_block / 4b79995f1368 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f0b8c6ac16cb9eb166c72a09cbd600adace8d7fb38984f7961c6ccc07743dc4"></a>

## Next pages — l7_ddos_protection.mitigation_block / 4b79995f1368 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b5d93240ad177bf7003b1ca054b369dbcc5c10575302c6b1d2a0c7cb5318e72e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d30d225050cec49120a6c6954d93d1e40d07ddb1dacc8142577d9fa95a235d64"></a>

## l7_ddos_protection.mitigation_captcha_challenge — l7_ddos_protection.mitigation_captcha_challenge / 8188b66d9681 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.mitigation_captcha_challenge

<a id="canonical-93a4e65b0a0470c88ae051aa104c66db4a8a18ba98adf45c5a036ccb15105236"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
mitigation_captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-76e9fed13e9cdce405439d9f6e705a29ceb818c8333f0c2edfb025fb53f0b8bb"></a>

## Direct properties — l7_ddos_protection.mitigation_captcha_challenge / 8188b66d9681 / 3

<a id="canonical-e38353bb81c08422b548c32d7bae2fcf145f1382a98402b5a28ab4a5c8dbf0e7"></a>

<a id="canonical-79345afa091070cd0e82b1f43861dd66628fa4c3689eda6b6bc954709e8357f9"></a>

## cookie_expiry property — l7_ddos_protection.mitigation_captcha_challenge / 8188b66d9681 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0ec8731c403d7a1922d87ac1c8dac76739db666cda77b3b0b80bc59953aec67f"></a>

<a id="canonical-87adf78fec11704278cfb484fa57696088bd78cda9fb0a487edf73f461a23138"></a>

## custom_page property — l7_ddos_protection.mitigation_captcha_challenge / 8188b66d9681 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4e06fe600a67b1403958ac32d304af0b0f74ae0dad5323d3e6df86beee65615c"></a>

## Next pages — l7_ddos_protection.mitigation_captcha_challenge / 8188b66d9681 / 6

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-50f1da05abef36adc0c6409a02ef2fe0aaa5720f36103b73042206ccdc3dea66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ce30deb42620e279323cef5ecae7e65f134c3b2668f078ec368266840bb4762"></a>

## l7_ddos_protection.mitigation_js_challenge — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- l7_ddos_protection.mitigation_js_challenge

<a id="canonical-501b1616a23965e7f3a3e2ff453ba74d3fe966ae01997eafeb99b2a9448df4cd"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
mitigation_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-23f39689812be22fccfd29ef501d9bbd647e0b16844acbda82fb9fee10070c65"></a>

## Direct properties — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 3

<a id="canonical-d8894b6336605f3cfbcdbab28a456097dac8890021a3e4fccd368c99d0cc6b1c"></a>

<a id="canonical-4f6aead620685dd5d3821557eb5b1b49d30f4766e53556159910503df9823d32"></a>

## cookie_expiry property — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-dc1648dc6d81b7f5d64426386517bd88956c4f90d8b92ee00759e53736f0c0f8"></a>

<a id="canonical-3ec7e43a88b4d92bfa44b0c0036598f0a0d5c79f26be0f07937735fba34abaf0"></a>

## custom_page property — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-6fc212e1e86b8362c46efc70ca9398e66fa871a2c8bff06c23fddd173661e203"></a>

<a id="canonical-144c1371c7adc795719b3250357c5c8bffe95afedf9ceaedebd5e2983a808d4a"></a>

## js_script_delay property — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-5b54f391e57b1f503b75901b2a436d85007225328214a166ca1f9b10b6f8298d"></a>

## Next pages — l7_ddos_protection.mitigation_js_challenge / ba58260586ad / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d01e18784c2ea24048b46a4e90708b3edcd19fab94a6a493e19cf7822cb11fb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-423de9aca4744caa248c5808584ff8a2bdbed4b5cd7ef2fa58ef578ee391ba68"></a>

## least_active — least_active / a7e4a7e22a31 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- least_active

<a id="canonical-063684637d11748256fee74bca48a125d8741c4075b2785483e6fc0b832132a5"></a>

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
least_active = {}
```

<a id="canonical-ab35cd29be1b09620d6b89b7b56c2b8b12a1bf4d65d80fb646e388ece970bde1"></a>

## Direct properties — least_active / a7e4a7e22a31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60d9edec28718f9ad389f05c8ef20428b07f18632783134865f8ad29ea61a6d8"></a>

## Next pages — least_active / a7e4a7e22a31 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3a9c0c6082edf3f8ab266f284730885ba0f4caccca0d568146e0f16986b03e3"></a>

## malware_protection_settings — malware_protection_settings / 185ae8da8241 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- malware_protection_settings

<a id="canonical-9d794a0e2d06dac24c5c9a6b68a3b7bc38c90d0c4ab591c9432720f0abd30c29"></a>

Type: `"object"`. single nested block, Optional.

Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in
real-time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("malware_protection_rules")}
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
malware_protection_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-c2fae662e3aae4ab99957bb4c2a4077b3c5cb042360b03d556060fb62dc02501"></a>

## Direct properties — malware_protection_settings / 185ae8da8241 / 3

- [malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2): complete subsection reference.

<a id="canonical-824b99bb733b2a876d4dcdf1770dd8de46b3df8b56788dde10d086be2263f9ea"></a>

## Next pages — malware_protection_settings / 185ae8da8241 / 4

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f86bbe723821dc18e6f7c0449918ccddf38686cfc187daf02379d898d3f7fdd1"></a>

## malware_protection_settings.malware_protection_rules — malware_protection_settings.malware_protection_rules / 54a89d1e626f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- malware_protection_settings.malware_protection_rules

<a id="canonical-83dfb9dce01abfe79b2e17e249deecf67bb03315e567a3b5a778ad8443fdb98c"></a>

Type: `"object"`. list nested block, Optional.

Configure the match criteria to trigger Malware Protection Scan.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
malware_protection_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee4221281beda3f2bcff984c969a1c7758ab90328818814503d4e7c1a70fd8fa"></a>

## Direct properties — malware_protection_settings.malware_protection_rules / 54a89d1e626f / 3

- [action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2): complete subsection reference.

<a id="canonical-5b96a029fc524f49d0c951fc781b7c2881782030c9015c426b1e3148b607c8a1"></a>

<a id="canonical-85c35a2fa829928b45dd7ec2bb0ed9794715c8c1f9c9c6e61a3b980f91d31b62"></a>

## http_methods property — malware_protection_settings.malware_protection_rules / 54a89d1e626f / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] HTTP Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

- [metadata](resources--http_loadbalancer--reference--group-020.md#canonical-2be85bff8b11cb7c04e4203b1a155a0ac3fb329e94fa9e8dcc5a55353e79796f): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-020.md#canonical-fa0cfaa6c7a9afaebc87c9a35d4105a8d59a2fd8874c56f2e63cfccf5749ac22): complete subsection reference.

<a id="canonical-a6b153a86efd7bdec13808c9676aa3195e2ce0274302b09cea945b93aa61a241"></a>

## Next pages — malware_protection_settings.malware_protection_rules / 54a89d1e626f / 5

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2)
- [malware_protection_settings.malware_protection_rules.metadata](resources--http_loadbalancer--reference--group-020.md#canonical-2be85bff8b11cb7c04e4203b1a155a0ac3fb329e94fa9e8dcc5a55353e79796f)
- [malware_protection_settings.malware_protection_rules.path](resources--http_loadbalancer--reference--group-020.md#canonical-fa0cfaa6c7a9afaebc87c9a35d4105a8d59a2fd8874c56f2e63cfccf5749ac22)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98d9250896ce37e4e23b9da29d5c2c2defb7c56b9b6dc46ad1f3d34daf97c0d3"></a>

## malware_protection_settings.malware_protection_rules.action — malware_protection_settings.malware_protection_rules.action / 533760ab870f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- malware_protection_settings.malware_protection_rules.action

<a id="canonical-a49b9a780139936c6823d007b8d831d829a904b4fd05c9ebba91c17472743a89"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-f168eff5b4c7a365b93f071878d35b3ca0a56f3eff26553f08bf6c6b247456a6"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.action / 533760ab870f / 3

- [block](resources--http_loadbalancer--reference--group-020.md#canonical-16366c83df2d4c5514520cb737dffea15a3b689e6a20abb568091a1409318b2f): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-020.md#canonical-d6f001fe5e21f7575ce107d27814cc1c61f19bad7dbc0702b17216970b1f3e12): complete subsection reference.

<a id="canonical-8a3df1402eec6ddab81a24288391805f6c1ec69f7e7616cc8695fb07a59a7885"></a>

## Next pages — malware_protection_settings.malware_protection_rules.action / 533760ab870f / 4

- [malware_protection_settings.malware_protection_rules.action.block](resources--http_loadbalancer--reference--group-020.md#canonical-16366c83df2d4c5514520cb737dffea15a3b689e6a20abb568091a1409318b2f)
- [malware_protection_settings.malware_protection_rules.action.report](resources--http_loadbalancer--reference--group-020.md#canonical-d6f001fe5e21f7575ce107d27814cc1c61f19bad7dbc0702b17216970b1f3e12)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-16366c83df2d4c5514520cb737dffea15a3b689e6a20abb568091a1409318b2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0a0d06b35e1453e7473662f8b31769ede61c45f294fcc274385d3a0986d0be1"></a>

## malware_protection_settings.malware_protection_rules.action.block — malware_protection_settings.malware_protection_rules.action.block / 5a39321cb505 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd)
- malware_protection_settings.malware_protection_rules.action.block

<a id="canonical-a12319305dba010f29158bb202b2033a813402f11fd719cb78ab16373f85bf55"></a>

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
block = {}
```

<a id="canonical-2457f009a78a570007e4d834cdd6e2e96d6b205712ae31d756f0c616feee8893"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.action.block / 5a39321cb505 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26dd565f8e8d9a98a975755f59f4194a75fc1e63b001ce05493ba11a12440926"></a>

## Next pages — malware_protection_settings.malware_protection_rules.action.block / 5a39321cb505 / 4

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d6f001fe5e21f7575ce107d27814cc1c61f19bad7dbc0702b17216970b1f3e12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6e4eba9750035143e50dcdfd6c5827eb19b35d753203f81fe40d7a682b8692d"></a>

## malware_protection_settings.malware_protection_rules.action.report — malware_protection_settings.malware_protection_rules.action.report / 422ba8050c6d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd)
- malware_protection_settings.malware_protection_rules.action.report

<a id="canonical-a6fa68e366e129f9751ee71f142a13a6478f9bb7b8a0bc6dae38a09fd57f04e0"></a>

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
report = {}
```

<a id="canonical-ac09d8597be9e25dcfda7c52c9870c0fa6643e4f8d7a6e0287092eb928d27eca"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.action.report / 422ba8050c6d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca33a2196a79be6ebf516129be7ecdaa25b676fdd5eac1b2fbaaa9ee11326eca"></a>

## Next pages — malware_protection_settings.malware_protection_rules.action.report / 422ba8050c6d / 4

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-53764377f5b493140e4c9dccd5d491d28cb1db2d3651095e29fc0cab278772fd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38b1534f0970bdf9ad7b9662289815e2d315239784b70d3cb34750dd9ca7460a"></a>

## malware_protection_settings.malware_protection_rules.domain — malware_protection_settings.malware_protection_rules.domain / b8e53490471a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- malware_protection_settings.malware_protection_rules.domain

<a id="canonical-43f202c1a40ff6da1a1e7cec16d6154f7391752b65aa590e00fb2d5631df4e50"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domain to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_domain",
    "domain")}
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
  "x-ves-oneof-field-domain_matcher": "[\"any_domain\",\"domain\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0f06dd2185977feee58fe8a6aa5c4fbf0ca77c63c4823fc21941197fabe53f70"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.domain / b8e53490471a / 3

- [any_domain](resources--http_loadbalancer--reference--group-020.md#canonical-9d48fef097862c09fcc725a96a6da91ba4d4d0be2614e2ea7ea13750d8d77adf): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-020.md#canonical-d1a417c56af6356de49d31a833809dc750ad928f9d0a4b288564c489f1beaef0): complete subsection reference.

<a id="canonical-fc6caeef9671962b7459b3efb7624756b7bf9012c273b2cdfe165ffd059c08b9"></a>

## Next pages — malware_protection_settings.malware_protection_rules.domain / b8e53490471a / 4

- [malware_protection_settings.malware_protection_rules.domain.any_domain](resources--http_loadbalancer--reference--group-020.md#canonical-9d48fef097862c09fcc725a96a6da91ba4d4d0be2614e2ea7ea13750d8d77adf)
- [malware_protection_settings.malware_protection_rules.domain.domain](resources--http_loadbalancer--reference--group-020.md#canonical-d1a417c56af6356de49d31a833809dc750ad928f9d0a4b288564c489f1beaef0)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9d48fef097862c09fcc725a96a6da91ba4d4d0be2614e2ea7ea13750d8d77adf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e59d31f6b8b716c94d4906f4832faeda3591e2c880ea557bb6133918e95cf4f2"></a>

## malware_protection_settings.malware_protection_rules.domain.any_domain — malware_protection_settings.malware_protection_rules.domain.any_domain / 899b0c6608aa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2)
- malware_protection_settings.malware_protection_rules.domain.any_domain

<a id="canonical-3a358b041eb6f11d5cacce1c061d5d8d92c2e9f96bfaf90acafce017278b96af"></a>

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
any_domain = {}
```

<a id="canonical-20becfb8828dd5d949b1bbc6f7cfe1d3e9dfbf5cefe16715ee4616fbc1f659f0"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.domain.any_domain / 899b0c6608aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-05186f2e4b5bd3d6aa83146ea376443e631bd5836e367fd794ee77822e7a53f4"></a>

## Next pages — malware_protection_settings.malware_protection_rules.domain.any_domain / 899b0c6608aa / 4

- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d1a417c56af6356de49d31a833809dc750ad928f9d0a4b288564c489f1beaef0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c00b7d9b62e0b20013053e38c2b15f425997677c88af73d45482583d885e1a99"></a>

## malware_protection_settings.malware_protection_rules.domain.domain — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2)
- malware_protection_settings.malware_protection_rules.domain.domain

<a id="canonical-c792c36c1fea63965d3e84747dca117de13f56d749d39cb02b653e760415b227"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf6470294806454b6215a9669576b934990d964cd039ab90b26ad2f2dec93020"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 3

<a id="canonical-2590eccee5e46c5e0ee917b2b87f7fa2ce980240315e1a3e93d590d869e4b010"></a>

<a id="canonical-321894521df801d10b58674c7c5ac50c35c31c5231826706d1e87b2b200cfd2f"></a>

## exact_value property — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0a029b10128ee0f173287e23bd282e435f52e88f7bc3575aa3931e1563e603d8"></a>

<a id="canonical-c98325183dde2c8223e68cb5501f0e5e3d2ba9fe787f5be080ae05864ab977eb"></a>

## regex_value property — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0410c4f18b5c3af8b1eb59dee20175a00fd039e9fc8046eadbec270ebb07cbc9"></a>

<a id="canonical-55ad4970f02d1b12b295592f0fbf70fa65b2187bfa2dfba6448f9eb5c6866294"></a>

## suffix_value property — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-ab8884f975abe6a7a37066279bfcb2df4590811a53d29fd5d5e8391519def309"></a>

## Next pages — malware_protection_settings.malware_protection_rules.domain.domain / 365680f66009 / 7

- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-55c75970893824db8c1375e83ad824d77be3f38df5c1cbae98461c19988a0aa2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2be85bff8b11cb7c04e4203b1a155a0ac3fb329e94fa9e8dcc5a55353e79796f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b07cdf89e88062d53703d275ccf53a3d4192f0b5d1b9f2298adc513ccf2d9a3"></a>

## malware_protection_settings.malware_protection_rules.metadata — malware_protection_settings.malware_protection_rules.metadata / c6af386d24e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- malware_protection_settings.malware_protection_rules.metadata

<a id="canonical-c50aae884c8c59fb067bd704f0bd883c5fc0a747460119d65bcec4bab50318c7"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-8bc061bbc64ebaeb3680f7696e3d35abb2e60fc4e35f4e415b883bb3b66170a0"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.metadata / c6af386d24e6 / 3

<a id="canonical-d876ed74cd2a6055313092baad579bfd3277c966af4cb8ae8b19a48ca7c60d2c"></a>

<a id="canonical-be985fbc0ce28830f6d5be4422dfc461a4cb19cc46a7241cff75b737c56c068e"></a>

## description_spec property — malware_protection_settings.malware_protection_rules.metadata / c6af386d24e6 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-5618946a27650d63b8ed3fc133fd481516bae4609461a37d4e1a93fe1bc80271"></a>

<a id="canonical-dad1d8251c44fde253f622e04b2ec17fbb8eca575c2d3d3cb6fad59637456918"></a>

## name property — malware_protection_settings.malware_protection_rules.metadata / c6af386d24e6 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-8bfc93db22b78f385a6e16d795d6b23dd4b28702495311bed95738336399d512"></a>

## Next pages — malware_protection_settings.malware_protection_rules.metadata / c6af386d24e6 / 6

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fa0cfaa6c7a9afaebc87c9a35d4105a8d59a2fd8874c56f2e63cfccf5749ac22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ba8d37b01b982b403a02bb5c23305b86773c7e8fb861dc17b4265c42d73b75e"></a>

## malware_protection_settings.malware_protection_rules.path — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- malware_protection_settings.malware_protection_rules.path

<a id="canonical-e2d5dd25c5de961786562230e818520ef1c489b93fb59d4852babfaa5da50b02"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-34206dde127ff2076ab18007b0d0d471e30478b197139517888adb36dcb16318"></a>

## Direct properties — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 3

<a id="canonical-1f1b8e7400f62702fe91785cc3ba92f87cebfe79093635f87f9824962b916d00"></a>

<a id="canonical-22fff3faaf5a99559516959b2c00d0b5a3b6e454224d012ae2e2ccfb8b12dc08"></a>

## path property — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0aa3920dfd7ce4c713803e85bb053df158225b2711439788564545fe733338f5"></a>

<a id="canonical-fb9a663fea01251350bbe4571ae41de677fe5abf1c6d45c1568307bd83c86999"></a>

## prefix property — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-ab808708ca838ef1c677b49073928649b8bfecfcd8a9f39244a93f2bafa62e9e"></a>

<a id="canonical-46dad5f6d8d6f0665d4ee7eec3e60768c312569c51f3fa64f9dcddf567f44ffb"></a>

## regex property — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2bc7cf4b7eacdf196fa306478ed1e5ac80b4d946cf05677ababb7f6d0eb1f396"></a>

## Next pages — malware_protection_settings.malware_protection_rules.path / b885f471e381 / 7

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0fcf064186129edec204085b27da7f1a32fc5a143b6e283a57d81eafbeb3a2c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22a222cb5be58e5fb9bd2256262991a4c24e2e9cfdb44ca18ea0111ea48c3822"></a>

## more_option — more_option / eda8f0687e86 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- more_option

<a id="canonical-31658663ea0e74591753fe703195dea444c09eaa78e56e877f37b0399cc41673"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-bb4bb2c47a00e8364cff20ec2dff5000b001c06846c9d9c021876516e1e51d6b"></a>

## Direct properties — more_option / eda8f0687e86 / 3

- [buffer_policy](resources--http_loadbalancer--reference--group-020.md#canonical-c1b19764b857e1c6951a995b210198d3ad7e57037fafd23cc255fa80bf850258): complete subsection reference.

- [compression_params](resources--http_loadbalancer--reference--group-020.md#canonical-a4080066048e2d3e44664b71e50148be0ab08478767ee3e4689c66b92b2e23c1): complete subsection reference.

<a id="canonical-6ed268cd450d6b75311899cb228f2146d43489ff29e86d723d7b43c01f706313"></a>

<a id="canonical-849d5e292c0121eb125bdba8f86b09031154a2cf2b86ab48a197079e4354ce27"></a>

## custom_errors property — more_option / eda8f0687e86 / 4

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-0310803dca748cc80c69668e6dab4f0dbdc92e6cf3b4042a3a2772db23f3928e"></a>

<a id="canonical-898057a21fc80d13617b582b1beeb74fe65cc49a8d1aee565b92e975fd779a83"></a>

## disable_default_error_pages property — more_option / eda8f0687e86 / 5

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-4f7819dc81eed733f27d18b086e2b565972b769bed3797a00aff23a441606e61): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-204bbf0f9b2c7bdae8e706e6c8501743b7d86f6d4c22e1a8f2915df7fffa24a3): complete subsection reference.

<a id="canonical-df7dc4f4a95e7e7f90f31b38d9f3c27246c8dd2f669ec038b72ee37a601fd850"></a>

<a id="canonical-99acdd0c6a7d1a3ef5ed142f38c6e9ea187c2e93250abbf434aab085d3de354f"></a>

## idle_timeout property — more_option / eda8f0687e86 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-b8859d47ea5f4f688a36db909b81debd0a41944787fd4c0e84f402eb3049ec7c"></a>

<a id="canonical-fe0bff335a55d76a71e3bc736136fa847e491a253247ad41984dc00188b15043"></a>

## max_request_header_size property — more_option / eda8f0687e86 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-e4fecc135493fdc38eea01a46f3298446fdfb3f9c535805db12b33dcb3667c87"></a>

<a id="canonical-bdf7698ad144fbb1c1833ecd368fe27e2b872e1e5b30d8439698bd07d5852d38"></a>

## max_requests_per_connection property — more_option / eda8f0687e86 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](resources--http_loadbalancer--reference--group-020.md#canonical-4b48432cbeadf2aedb5cf1e4ed353df8d271f7a86fd0baaaf4e1296e6e951ed3): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39): complete subsection reference.

<a id="canonical-018cb243a15d28b6611d71780b9415f9fbf1c8966d8568a309f793b975876de8"></a>

<a id="canonical-4ed84615b13ed7d34667259e316c662b3ed2e452593113ffedde4ebb8f84d335"></a>

## request_cookies_to_remove property — more_option / eda8f0687e86 / 9

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8): complete subsection reference.

<a id="canonical-80040f7fd296a5feaa4dc5a1b7a4a4a4ba7ff2e4abe3a9d0f0c4e3a33321d8bf"></a>

<a id="canonical-63a1b9563f9d1265ceebac7f5b09975a5e246c360f44e93f038e331bb3865835"></a>

## request_headers_to_remove property — more_option / eda8f0687e86 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa): complete subsection reference.

<a id="canonical-54be1132ce3f31c1835ef57e9b59a5c7543eb3dd5bfaebd66d1f0c5eb9b7f1bb"></a>

<a id="canonical-15b7b1353a55fe2c1963ec5f2348a0ce6518dd83cda4653bba1531434fad81fc"></a>

## response_cookies_to_remove property — more_option / eda8f0687e86 / 11

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926): complete subsection reference.

<a id="canonical-91b5ebcfd0b60a2529d60cf59913849ad40066122a5f10a82355f9d50f5efa2a"></a>

<a id="canonical-9a9551ca6badda220c5eff8b70c5e1df95d4d8b2c8341e1206f5520e069f2803"></a>

## response_headers_to_remove property — more_option / eda8f0687e86 / 12

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2db413e47e2359ff0241da296055c98573876f91f9868431edb619a894a9b9cc"></a>

## Next pages — more_option / eda8f0687e86 / 13

- [more_option.buffer_policy](resources--http_loadbalancer--reference--group-020.md#canonical-c1b19764b857e1c6951a995b210198d3ad7e57037fafd23cc255fa80bf850258)
- [more_option.compression_params](resources--http_loadbalancer--reference--group-020.md#canonical-a4080066048e2d3e44664b71e50148be0ab08478767ee3e4689c66b92b2e23c1)
- [more_option.disable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-4f7819dc81eed733f27d18b086e2b565972b769bed3797a00aff23a441606e61)
- [more_option.enable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-204bbf0f9b2c7bdae8e706e6c8501743b7d86f6d4c22e1a8f2915df7fffa24a3)
- [more_option.no_request_limit_per_connection](resources--http_loadbalancer--reference--group-020.md#canonical-4b48432cbeadf2aedb5cf1e4ed353df8d271f7a86fd0baaaf4e1296e6e951ed3)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c1b19764b857e1c6951a995b210198d3ad7e57037fafd23cc255fa80bf850258"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0df42277c7324889fb3adc16f83cd68c6b4935ddeccc0624cf45a02c22c30f36"></a>

## more_option.buffer_policy — more_option.buffer_policy / 30aacd53b465 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.buffer_policy

<a id="canonical-3b5d089e438dd7d2dd179882f5fb51e950ea25ac0ed6f128c98ee0db0f5f76c7"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-5fd38e5d1cd341e30ba9ff76eeaaeed01bb05f60c3ebb238a75e5fa8ee7ea599"></a>

## Direct properties — more_option.buffer_policy / 30aacd53b465 / 3

<a id="canonical-4227217bd517b0a4910a23b888e0e9432a6efc8b255ddbb72ddfb072810e1950"></a>

<a id="canonical-af717947e85ae2f6358dd5e8bfd9a1666cd711b9e1cdc1311e92759537a666b3"></a>

## disabled property — more_option.buffer_policy / 30aacd53b465 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-d172ad90cee3237640db41ee64f08a4ac67ecb7cf635381e610104ab85f82324"></a>

<a id="canonical-461e5d2ae7f2bd9cd57778bfbafc92dc36367533a2ce9937883e7f5858536664"></a>

## max_request_bytes property — more_option.buffer_policy / 30aacd53b465 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-cc9def6decfb4a07c6458d228ca7dfc06ee94b5b368b09d4df1c8c1efcdab7ea"></a>

## Next pages — more_option.buffer_policy / 30aacd53b465 / 6

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a4080066048e2d3e44664b71e50148be0ab08478767ee3e4689c66b92b2e23c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34461bcac668a930af1ad3bb2506dfdcab288b38fae7bf4d9f86e6fb890e9f65"></a>

## more_option.compression_params — more_option.compression_params / 9c3d25f26bf2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.compression_params

<a id="canonical-59746c3e89eee5e537e033073bf3faebbe45ca194fe63ff80dddd31812d8cf44"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/javascript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-81e5462d4628db649b2949137d5ee2e1f5b3f754f3f5c524a0e409f46ed72140"></a>

## Direct properties — more_option.compression_params / 9c3d25f26bf2 / 3

<a id="canonical-aea9660a64be17b085b8be8c7f385c70a82ca07ebcacd47ac45d9195e2a0bbda"></a>

<a id="canonical-0a43ab182ac4543bc8bc1437d15f4df0fc9459fd2ad20f94037f64a36f1c3d13"></a>

## content_length property — more_option.compression_params / 9c3d25f26bf2 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-6a65de5234a7e402965cf595b2a53d995850e282bbfe091d9b5053fae4f2f1d5"></a>

<a id="canonical-d297c2cf13f1912cd3756c5e3ae0791c554c1fd3261d8c1882e490b47e569d10"></a>

## content_type property — more_option.compression_params / 9c3d25f26bf2 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/javascript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/javascript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fa7abd4347c4e543418fbc44f714951e16b62f92298f0b691158ab7d4f2fdaf3"></a>

<a id="canonical-b0d8ba564fcfa1b1d94afcabcb55b0b9ed1b76714b4269644b81d38dd0a5b9cc"></a>

## disable_on_etag_header property — more_option.compression_params / 9c3d25f26bf2 / 6

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-c3f3129735effbb21c9e86ef8f86767942c21eeaf87d7ae66b8202d2fc855a02"></a>

<a id="canonical-24c75aca6dc785a39146666d440e79484745d78fcb1ed863642101456e010805"></a>

## remove_accept_encoding_header property — more_option.compression_params / 9c3d25f26bf2 / 7

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-a3bfc8e376493466dfb6e72d040832939fdd42107baf3b212f6cd6f0c0dbd9d4"></a>

## Next pages — more_option.compression_params / 9c3d25f26bf2 / 8

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4f7819dc81eed733f27d18b086e2b565972b769bed3797a00aff23a441606e61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7814cf835291c7609aafd12d0735055276d038ad9e188c6c42f0d654fbd69918"></a>

## more_option.disable_path_normalize — more_option.disable_path_normalize / 35493bd316a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.disable_path_normalize

<a id="canonical-0ddb8788c6e42be6863e8d6228e3da211fdb43c4c1c5a0197ce69b0ae4c0fe2e"></a>

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

<a id="canonical-674fa9d82ff62f44aebba2deb7c1349a08a50194b9e170a265bc8034f3b8e6a1"></a>

## Direct properties — more_option.disable_path_normalize / 35493bd316a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f287b1b43c71af8f7f63cb6bd8cd5ec9b17ee1746e15a8efade7ffb0deb1428c"></a>

## Next pages — more_option.disable_path_normalize / 35493bd316a5 / 4

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-204bbf0f9b2c7bdae8e706e6c8501743b7d86f6d4c22e1a8f2915df7fffa24a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bc7f4e45859a2df5b5e242375206da6b00d3f5e7894db4d69069a028e544e89"></a>

## more_option.enable_path_normalize — more_option.enable_path_normalize / 5e749c2d3d64 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.enable_path_normalize

<a id="canonical-ab78d121c107c4b350c9c83c5d7f61c559be074c6fc472fd9958544b52752a5e"></a>

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

<a id="canonical-dbb5b5bedc05451185befe24e80f30dd177fdd3317d4876ce73a44be115fd662"></a>

## Direct properties — more_option.enable_path_normalize / 5e749c2d3d64 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fca541647c25d331e9003489d64f93ac3674818fb4f0c3c2834859684eb1a80c"></a>

## Next pages — more_option.enable_path_normalize / 5e749c2d3d64 / 4

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4b48432cbeadf2aedb5cf1e4ed353df8d271f7a86fd0baaaf4e1296e6e951ed3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72237f49401be551e74c6590632a2c1d31c1fdb29e174cbb980f835133eb132d"></a>

## more_option.no_request_limit_per_connection — more_option.no_request_limit_per_connection / e3019c570736 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.no_request_limit_per_connection

<a id="canonical-551a29bdd1ac2c4a557d4ca7301ddc51f9d17ae8d403c7660cfdcceb2fb780d5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-7f8b54018e100a86d2a87b1911aeb165eb9a6f0bf026c381d7e10b11a827047f"></a>

## Direct properties — more_option.no_request_limit_per_connection / e3019c570736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a2b291a3c9c8a1305d7a6fa6be155dba5af882b659ac977dbc31e7e486750c1"></a>

## Next pages — more_option.no_request_limit_per_connection / e3019c570736 / 4

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2765a397d196b468dd28af49207079193a172ed5022c4f0558e8944ad3915f28"></a>

## more_option.request_cookies_to_add — more_option.request_cookies_to_add / 2dc03a21111a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.request_cookies_to_add

<a id="canonical-ceedad14d931b5ddcbe3ef81e2b4e9fe58c5f373a77deab86cebe836d99b6528"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-e28ea597aa4038a919aea9aff30c177f259cbb36427459959952464abbbb38fc"></a>

## Direct properties — more_option.request_cookies_to_add / 2dc03a21111a / 3

<a id="canonical-eb36c737921e7e272cdcc62a84e4800e6ed6a4d5147bd17e321be7140a1ddac4"></a>

<a id="canonical-0de79498356ec15377ddf76912225aa1c375239de762e3b4834928bab2ebfca3"></a>

## name property — more_option.request_cookies_to_add / 2dc03a21111a / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3c04ad7f56693b838efb49e0839e4acac8c3556856f7df5c969f3af51c5dbf82"></a>

<a id="canonical-c02d9e97ddc3cf5ad25d3d212868090beb37a48438c7458a309162276d470d95"></a>

## overwrite property — more_option.request_cookies_to_add / 2dc03a21111a / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2): complete subsection reference.

<a id="canonical-27f6e84f77d1ef8c7a1e1f9081fccaa4ea9a6deb7efd7b187a1665079697847e"></a>

<a id="canonical-c7e0844c6675b83408b918f94baa34e6d5ee880ff66965ac34c4e91d286f6863"></a>

## value property — more_option.request_cookies_to_add / 2dc03a21111a / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

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

<a id="canonical-d6a4c489f3093ea5dd88c620f7fd857dbcb2755fbf997ea1fcf7f1d3268bce75"></a>

## Next pages — more_option.request_cookies_to_add / 2dc03a21111a / 7

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acf0924a7e8f7610df698126725a7635902bd42e13c7c8e81633b122bc5e184a"></a>

## more_option.request_cookies_to_add.secret_value — more_option.request_cookies_to_add.secret_value / 67e8ab903b58 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39)
- more_option.request_cookies_to_add.secret_value

<a id="canonical-2967110dd7e0a14f9ea3aa3ef1f9ce5860b02284aca3f50cf10f71a99f387299"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-86a35c296a0e8f38efcec97625566373a194664973d8b49fe1d2133b38d16c7e"></a>

## Direct properties — more_option.request_cookies_to_add.secret_value / 67e8ab903b58 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-e6ac5d6c514495990007b459d913b1a95a3cf8bf17ea6daa1c47157346279053): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-5cb078aaf9cf25d40d29c8e2c901e7656e569603d570dd88807a203d674023bb): complete subsection reference.

<a id="canonical-c996c9ca976d6a4f3f32f4afc006135f07e4597f5a41990d92fd111b11ea1a5d"></a>

## Next pages — more_option.request_cookies_to_add.secret_value / 67e8ab903b58 / 4

- [more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-e6ac5d6c514495990007b459d913b1a95a3cf8bf17ea6daa1c47157346279053)
- [more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-5cb078aaf9cf25d40d29c8e2c901e7656e569603d570dd88807a203d674023bb)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e6ac5d6c514495990007b459d913b1a95a3cf8bf17ea6daa1c47157346279053"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98f9a12299af59c7475a29b4ea4b16d313f0118376bd06466ceb5c3fdb3d93a8"></a>

## more_option.request_cookies_to_add.secret_value.blindfold_secret_info — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2)
- more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-cac71f27dcae1586e91f224d3a4627239e78a15f5d8ef7ab3f980a729569a3d9"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-aac4ef503fc0e987a55fbb8455a25b77e51ec7dc15a0e9081a77b0cd3cbdc181"></a>

## Direct properties — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 3

<a id="canonical-d5909a4d7f1cca05a3bf3f0ca97c40727b9ebd6f53d2ea77747735fbe6e07a72"></a>

<a id="canonical-d79eec6381daa59eafce7c44f566f86bf109acab046aacb061d4205cc2de2c1c"></a>

## decryption_provider property — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 4

Type: `"string"`. Optional.

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

<a id="canonical-ad7cc0d47be6ae4d8d3fc4c695db7770639e83b2a6b0b86e6f5e48388738a3c0"></a>

<a id="canonical-baa4a3cbcda9c7176eb43e2edee4ae32195935131871ee575dfcec2873b3be87"></a>

## location property — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-74ba74a4eb355e6f084d8c737f0641abbbe3bd703f0cb2abdf24e5e381b36916"></a>

<a id="canonical-bc08ce709bb99fa23f1a08cf6cb9037a81c463a79d9b88402bdbbf2c6e12c770"></a>

## store_provider property — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 6

Type: `"string"`. Optional.

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

<a id="canonical-929f35f35c8f354beddf159662df0133db59c0f8ee99ff273c42918051a9fd2e"></a>

## Next pages — more_option.request_cookies_to_add.secret_value.blindfold_secret_info / aa09438e5631 / 7

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5cb078aaf9cf25d40d29c8e2c901e7656e569603d570dd88807a203d674023bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07fb0f82fabfb0edfd2f970c85ab4426c961d6ddf5e37c55c336799c18287198"></a>

## more_option.request_cookies_to_add.secret_value.clear_secret_info — more_option.request_cookies_to_add.secret_value.clear_secret_info / 62a7d7b1bb6a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-4225e8165882984a2897116811290e9fa2b0c59ee3e5b0e2e1a5d248a18b8c39)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2)
- more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-99eadaea4468fc94ca5c46b9c4ee2bcd41f16af7a1c2c0feb272fb5fb4f1cb2e"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-84e8f641fe9e43244b870ac4fe79b8241c789fba1a19ea097cde9c53b0c74c53"></a>

## Direct properties — more_option.request_cookies_to_add.secret_value.clear_secret_info / 62a7d7b1bb6a / 3

<a id="canonical-ee245e50586f344b9448676aeadbb9d06842e78990d1a2797bdd7c382ce7d01b"></a>

<a id="canonical-7b3af4de1e9878fbaddd015462984e88f7d31126f4febf4c4946884c982b1c0a"></a>

## provider_ref property — more_option.request_cookies_to_add.secret_value.clear_secret_info / 62a7d7b1bb6a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9462c7e75bddc559d7cbaee8c21d058dc2e0dd01c1fae3aeb4db1ea367426f74"></a>

<a id="canonical-87c440ba0e96bc1018946aaf149057628f0b184977ff1572eca87191446735f1"></a>

## url property — more_option.request_cookies_to_add.secret_value.clear_secret_info / 62a7d7b1bb6a / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-5088c12bae8c8a175e3d72243317f81cf8b765c5e2a588c1498baab2a4d50355"></a>

## Next pages — more_option.request_cookies_to_add.secret_value.clear_secret_info / 62a7d7b1bb6a / 6

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-60b6241dda4d2dacf9ed9c7238c8ad299ca1dec3c78f9a8e27e48bf812e710c2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0e6583ddf4de7054f81027d17c4b31225c8ca36615e64a6651d07f327834889"></a>

## more_option.request_headers_to_add — more_option.request_headers_to_add / 74ec808a2d32 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.request_headers_to_add

<a id="canonical-9c8c96133bca307822084b363751af21853edbd9c0e5cbb05f3a72a8dc8e8d33"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-6637cd027b3bb2e696fc53e9bbb80dabecd768b3e00b239e67b21f93564e3adc"></a>

## Direct properties — more_option.request_headers_to_add / 74ec808a2d32 / 3

<a id="canonical-089b89eb3890402357cadc83e1743697941f7406b4886a9dd6869a5c642f1e17"></a>

<a id="canonical-6525fc4763fe7546219e0fa018d1c2175ccf3d022d325d894385ca180887fbbe"></a>

## append property — more_option.request_headers_to_add / 74ec808a2d32 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-6b50a3dba6059800782aacb42407dc8c88f869bd924d37fe76d028583717aeee"></a>

<a id="canonical-270d077bd4f84f1ccb4347abbbe549d4fb67fc994b861a25af1fe018caccafdb"></a>

## name property — more_option.request_headers_to_add / 74ec808a2d32 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e): complete subsection reference.

<a id="canonical-54da54aef48a5a915dd32e58a46c68e742e9287047412f621f3c798fb4022ee3"></a>

<a id="canonical-b802bdae24963713303e4f968c5ca47c90d6ab638d88af9f9501201aa269b0fb"></a>

## value property — more_option.request_headers_to_add / 74ec808a2d32 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-39b9838cdd224590d383705b0dd00e176de6a72f489c33923c84b07e5ccd8ebc"></a>

## Next pages — more_option.request_headers_to_add / 74ec808a2d32 / 7

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b093499bafa88fcdc5b6cda118a7b7042457ad95c62dedf25d730adc0bb9a683"></a>

## more_option.request_headers_to_add.secret_value — more_option.request_headers_to_add.secret_value / d31dc61c08be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8)
- more_option.request_headers_to_add.secret_value

<a id="canonical-c90ff9b4264dd67e343b88deb0dd55749db2fad2b1151de27758ebaf20cedcf0"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-814b3ec311c04caecbd4b218473c73f7aaaf2dcdaf7c8bd6b4f839a8b0b0aac9"></a>

## Direct properties — more_option.request_headers_to_add.secret_value / d31dc61c08be / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-844de91a96e3e131b65e35e888892214ec546892873a81b191bbcdf113022e79): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-dee87c04b906ebacb9bae3622b9fab7344f049f7fa816565b4de3565534a6161): complete subsection reference.

<a id="canonical-9494c4c3afd3d07b09404932072ab43fc69c7cd28eff878fd2e205706b5c953c"></a>

## Next pages — more_option.request_headers_to_add.secret_value / d31dc61c08be / 4

- [more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-020.md#canonical-844de91a96e3e131b65e35e888892214ec546892873a81b191bbcdf113022e79)
- [more_option.request_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-dee87c04b906ebacb9bae3622b9fab7344f049f7fa816565b4de3565534a6161)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-844de91a96e3e131b65e35e888892214ec546892873a81b191bbcdf113022e79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
