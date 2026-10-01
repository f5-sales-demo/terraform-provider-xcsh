---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-6d60227e1092a8650ed034cb1230843782873fdc914ed8c8a02e6ccde20f23b2"></a>

## proxy_config.https.tls_parameters.use_mtls.trusted_ca — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-ec620ed4e1ceb922a41bc128e01244aeb0fb1aa8623271798d2ef1a91d157fe0"></a>

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

<a id="canonical-068023f5ba3dccb835ec7ad4582c4dc66b20712a61541d012224d7affc1f07d0"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 3

<a id="canonical-7c8e38d2ac3632dc782a06491206e36a12ad9d448b120239ee43bfd9ea27365a"></a>

<a id="canonical-91735bf70ceeef14aba239c9f7747e959779dd4a81191c49a8c348ee8a631129"></a>

## name property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 4

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

<a id="canonical-cfc89534b27cc3ccb20932fef8ba9d1ab55f2b664ebd7521101229abe80df15b"></a>

<a id="canonical-eea48ddab793011e6357cc3468c18403171ebd16db8e7d93cd85830eadb88596"></a>

## namespace property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 5

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

<a id="canonical-ea69f670383837a7d08e818e549e30186184511ebc01fd55102a560dc6c57709"></a>

<a id="canonical-90ff0353df84084934cbaf2fa1c71e824fef9d79fbf4028e055a944301c85932"></a>

## tenant property — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 6

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

<a id="canonical-853e1a4d56f7cc6d5ea79aafb99c5dd48d02c02d8c45b769651a9289b1bbea40"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.trusted_ca / d88b480dbff6 / 7

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-99a7b6ea2820551a1c2746a1ee6bb38e89719aa6fb15655b46cd67a766a646ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e3144849522e62397725b27550979ef6a39d3806ce4e32ba9eb42f73ce4f3c"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_disabled — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / ba9c1bb86d1f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-c566a03e11427042823e090396dab37a064b816e7a48ccc8597dd71e5ec61286"></a>

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

<a id="canonical-f3d7367c42b6ce6a25b2284ceae84783cb8c250cbf1b63fa4e100aa75e23454e"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / ba9c1bb86d1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c28c4be5f3c1903344fb7fda7bd5c0da533172728393b558f9f086ad4d2e7c5c"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.xfcc_disabled / ba9c1bb86d1f / 4

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-917867436879ba25f4b16d939daa8aae74dc830bffd147d9ff4cccd7d94d1236"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c0fd9900109b1f02ccb41c2790e50b8d15f546651e9f36d1c034b5286832b0c"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_options — proxy_config.https.tls_parameters.use_mtls.xfcc_options / 628c461653bb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d405255f80adf91f3a4ad9bfd0e009ab1873464798756708f718054e47015079)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-12b1668564b57fed383102e0655e5c2fd223f895bd3f947a1500d20e88bd4b90)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- proxy_config.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-bcaaab5fecf7103bbbc411b4950be077e8ee04e5bce81f2015e202487039407c"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-9f40f70c9a83ad92f3db53c6476d42306642de08fa70945bfde416aedba2b9d9"></a>

## Direct properties — proxy_config.https.tls_parameters.use_mtls.xfcc_options / 628c461653bb / 3

<a id="canonical-a677d94a10e6dac851af0f440efedaa8060a47b153a4ce37b4ceec210709e9cc"></a>

<a id="canonical-d2f30ce226fb49c3d004c3a7617be85f007638c978de5b08bc9b0c438d0c41e6"></a>

## xfcc_header_elements property — proxy_config.https.tls_parameters.use_mtls.xfcc_options / 628c461653bb / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3143674bdc5abdfc4764e7807ab3fc9ec1bb434d8d1be48ef7f5210deea98225"></a>

## Next pages — proxy_config.https.tls_parameters.use_mtls.xfcc_options / 628c461653bb / 5

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-86372ca7fbc2e20f711797628e4eeb717449261b6371d6fc347c5aba93edf190)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7abf91d7abab288a34e0a758f4873d44f417115f8f485e20488bded5bd607474"></a>

## proxy_config.https_auto_cert — proxy_config.https_auto_cert / b962218677ec / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- proxy_config.https_auto_cert

<a id="canonical-3ef7c36b7533a3872aefaec38bda05195fab6b72e8121c8e1d6a0e72748ced09"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

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

<a id="canonical-45b413922d075c077f08fc7c6a7668e4da7823bb56afabba005952d6fcf4527b"></a>

## Direct properties — proxy_config.https_auto_cert / b962218677ec / 3

<a id="canonical-11ef11f244c032c128274e0d8b33e9c33f3e9165d4ec6e756a70710d8d226690"></a>

<a id="canonical-d64aa8c3c9832bb2dd8f0dc1fb0c5720c0088674d72d9aa5076e30a1def3b7b4"></a>

## add_hsts property — proxy_config.https_auto_cert / b962218677ec / 4

Type: `"bool"`. Computed.

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

<a id="canonical-63e6def0fd14e9812233c0c633eba736ba8f5dad4cd9c005c0f4411e1e0e01fc"></a>

<a id="canonical-f559c20d36bf1fa7013028e878806413cf00d70a776a7d63fca637f4dff05af3"></a>

## append_server_name property — proxy_config.https_auto_cert / b962218677ec / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c): complete subsection reference.

<a id="canonical-438ed26a3fdce77bf3bad266aec8219c634da7463fa0f7be32bf74c34504ed70"></a>

<a id="canonical-59135f97520d0e4ca1581583e48987eb1bb3f34c5ae8b31451367d8f5ac614c8"></a>

## connection_idle_timeout property — proxy_config.https_auto_cert / b962218677ec / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0cc05e311e08de210db740c90551f0ebb9a2db237bd63d7b2497adc2d40d132b): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-90d527dd13fe9f93c574518cdabf57de499051197201ff313613014f7113e9df): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-476409eff1dcdb96ae1b1ea2f052add63ecbf35fe7f60070576e18562442f021): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d79c90f25d981a4c02bbc8c0d91f54ca7cbfc9900338b306d22b1751d5ca0c91): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655): complete subsection reference.

<a id="canonical-e3dbd983b0f2a1880cbae2b1e5d4f6b0094c41b35df6244e57a046a238c1d0de"></a>

<a id="canonical-5617fd5639fe2c5f4a5d5eaefa703e99211ba6bcacec16d66b608827178e8e38"></a>

## http_redirect property — proxy_config.https_auto_cert / b962218677ec / 7

Type: `"bool"`. Computed.

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

- [no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ce45e91f1f6c9d5c3f42e470d48b9c9817f843d43f04992dd2420002e5c85eb2): complete subsection reference.

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-325da1dca17116c529debc17ba83bb9f3ed9033337561b2708f03b79ef1d0aaa): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-6a5ce9c3be3ed3a2fd17a783b3ac61e842dd3e304efba117a370e43d36af3207): complete subsection reference.

<a id="canonical-bd8f7704b75afbfb9f986bdf5d5e2ef972caddb74e304d86ba93cd79783dc760"></a>

<a id="canonical-c6153a4c0dd3b784ceaabd1722ee8f50f67def6e746214fbdf723766899ac2da"></a>

## port property — proxy_config.https_auto_cert / b962218677ec / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-e79a28b026201900ab970350b70111fefc855bd3863e689e7925d1be31e559cf"></a>

<a id="canonical-d8617e9e5c05ca493476adb3996c4554445570ba73978c16ca5496022fa1a115"></a>

## port_ranges property — proxy_config.https_auto_cert / b962218677ec / 9

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

<a id="canonical-c3fbc1da1922128477e527330361c0ae0384272843da89ee35fe592c0a401f05"></a>

<a id="canonical-c572aa073fb59d35666ff55e47e98d96b677421e08955592e0497e5f5d5b57f8"></a>

## server_name property — proxy_config.https_auto_cert / b962218677ec / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0): complete subsection reference.

<a id="canonical-d50f9eef9005ab15a16fb327dbbdb912cc74d22850de96129b10de8f78602fc1"></a>

## Next pages — proxy_config.https_auto_cert / b962218677ec / 11

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c)
- [proxy_config.https_auto_cert.default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0cc05e311e08de210db740c90551f0ebb9a2db237bd63d7b2497adc2d40d132b)
- [proxy_config.https_auto_cert.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-90d527dd13fe9f93c574518cdabf57de499051197201ff313613014f7113e9df)
- [proxy_config.https_auto_cert.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-476409eff1dcdb96ae1b1ea2f052add63ecbf35fe7f60070576e18562442f021)
- [proxy_config.https_auto_cert.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d79c90f25d981a4c02bbc8c0d91f54ca7cbfc9900338b306d22b1751d5ca0c91)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [proxy_config.https_auto_cert.no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ce45e91f1f6c9d5c3f42e470d48b9c9817f843d43f04992dd2420002e5c85eb2)
- [proxy_config.https_auto_cert.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-325da1dca17116c529debc17ba83bb9f3ed9033337561b2708f03b79ef1d0aaa)
- [proxy_config.https_auto_cert.pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-6a5ce9c3be3ed3a2fd17a783b3ac61e842dd3e304efba117a370e43d36af3207)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3573cc6384b2bddef79ffbce784caa69b3df26062248e6d35f2ab35917dde6fe"></a>

## proxy_config.https_auto_cert.coalescing_options — proxy_config.https_auto_cert.coalescing_options / 25dc01258062 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-bbb5cec6e4b3a7739d2082fc84ba3ea56bcbc0c177d061a1f55974bc76faf6ea"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-edbbe2fcd3378c4b72796894a89836a8803c1eabc294fda66cf3bfdc875d0d43"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options / 25dc01258062 / 3

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-61b4aeea36b9665a775fd04c17567f7d4367e94c43d51b6f112192c79232fcb7): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ad7dc44d90b4ca7b0c10f541936b2cb69a2bb16cb8ae49f0eb052d519f5f5043): complete subsection reference.

<a id="canonical-c7935d3093eb985b5c69b59d275c2695b69e352b83034335b9a635514d73ca54"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options / 25dc01258062 / 4

- [proxy_config.https_auto_cert.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-61b4aeea36b9665a775fd04c17567f7d4367e94c43d51b6f112192c79232fcb7)
- [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ad7dc44d90b4ca7b0c10f541936b2cb69a2bb16cb8ae49f0eb052d519f5f5043)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-61b4aeea36b9665a775fd04c17567f7d4367e94c43d51b6f112192c79232fcb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7361c6f4c20a0068948d2ba9e0933ff4525056ea70633d8a4602363fb915093"></a>

## proxy_config.https_auto_cert.coalescing_options.default_coalescing — proxy_config.https_auto_cert.coalescing_options.default_coalescing / c5d6ad3c4434 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c)
- proxy_config.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-a071341990ceb7e43f3b40fbefdfd2e5a9a34ca0e3a4e409a8307f83e54753cf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-71dfb7b0174c3631574c64d7339e03e65ead5832246617fe63481814aa55fef1"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options.default_coalescing / c5d6ad3c4434 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-493c8b50728d9d73aa671c0379a3ce21d0babb904b1e6ad58abc3c85f42602e2"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options.default_coalescing / c5d6ad3c4434 / 4

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ad7dc44d90b4ca7b0c10f541936b2cb69a2bb16cb8ae49f0eb052d519f5f5043"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24a11b53ccc682c26910c72a09673b14ea597799549fc2b94473751514094a14"></a>

## proxy_config.https_auto_cert.coalescing_options.strict_coalescing — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / 82484d2e568f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c)
- proxy_config.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-d6602addce1efc876ec1ca52777748472aec7afde4bfff18c072d48c5654ed00"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-832701df68ec9cc92abe9255e73c5dd17ac8cf6b44ad23ce3c725017863c98d1"></a>

## Direct properties — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / 82484d2e568f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-555896201df28870edc87413936701288255e674ebf457f02e7c87ec7411724f"></a>

## Next pages — proxy_config.https_auto_cert.coalescing_options.strict_coalescing / 82484d2e568f / 4

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b4ca39dbbf91b558a19721b45fbbc8d22a288b85ca43c7bcea67c8968d0a304c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-0cc05e311e08de210db740c90551f0ebb9a2db237bd63d7b2497adc2d40d132b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b5ee668687551403a900c6c1a2c25faf498dbdd996ad91c452a8677a10068b2"></a>

## proxy_config.https_auto_cert.default_header — proxy_config.https_auto_cert.default_header / 35f88524b337 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.default_header

<a id="canonical-68a7514bc893ffc99777dbe92341adca96b07f7b4253efbcf771b59c0c8aa822"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5146fbbb5114727cd1b272328bbb932b9dc6a50f64489e2dfa32b4df457dd7ee"></a>

## Direct properties — proxy_config.https_auto_cert.default_header / 35f88524b337 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-942fb327b691c42a7ffb51076e6841a1f421c108bc4f5bcf15a7f22f4b68febf"></a>

## Next pages — proxy_config.https_auto_cert.default_header / 35f88524b337 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-90d527dd13fe9f93c574518cdabf57de499051197201ff313613014f7113e9df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cc43723890d044513cd48575a03795899a5d1a5b345306db4dd35b7de7daaa4"></a>

## proxy_config.https_auto_cert.default_loadbalancer — proxy_config.https_auto_cert.default_loadbalancer / 55aba3e908c5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-bab057f1e649f6380a4829e1e0b2bb71b407d0ccd4edde69534f5dd5cc0342bd"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-696abaa208c24e2eed8ca6614e9a6884fed97cc34612e0a101266967d1e4f2e1"></a>

## Direct properties — proxy_config.https_auto_cert.default_loadbalancer / 55aba3e908c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dfb6014b989ad2ceb3c568fc9e0ec68232afc04ebc3a7c7a78fe07c920892dc"></a>

## Next pages — proxy_config.https_auto_cert.default_loadbalancer / 55aba3e908c5 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-476409eff1dcdb96ae1b1ea2f052add63ecbf35fe7f60070576e18562442f021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb4b088137d05edfcdc9903316611aebcc50644444ae8d4ed2c0d0d6b51f554c"></a>

## proxy_config.https_auto_cert.disable_path_normalize — proxy_config.https_auto_cert.disable_path_normalize / 91a6587ff013 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-fddab9fbf88f3ff36a7336a03db0d2ec17b6477418f6228329bcad08a4449460"></a>

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

<a id="canonical-ada393bfd8172edf063b36a451bf7e95edfe1ad0eeb29558d268b47b6ab1fb41"></a>

## Direct properties — proxy_config.https_auto_cert.disable_path_normalize / 91a6587ff013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4720033f451da9cc634754b0e3bb92dffccf2e02223c77c932bd2f999439ffad"></a>

## Next pages — proxy_config.https_auto_cert.disable_path_normalize / 91a6587ff013 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d79c90f25d981a4c02bbc8c0d91f54ca7cbfc9900338b306d22b1751d5ca0c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44c7feb2290372cc7ef31c9a62fe53951259675a4b9865577fc0614c2fa8579b"></a>

## proxy_config.https_auto_cert.enable_path_normalize — proxy_config.https_auto_cert.enable_path_normalize / c3651d0de746 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-3719345eb6af3051400c1d8ff89f930fad7629d99a6299357e4bb43afe3dec77"></a>

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

<a id="canonical-21d95fa6da700addc75a25c424683f45391414cfbaace30406762d8a784a726e"></a>

## Direct properties — proxy_config.https_auto_cert.enable_path_normalize / c3651d0de746 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c6242901925079a2024b0c824d8e88ae9d891b4e11d7265a395acb1df02009d5"></a>

## Next pages — proxy_config.https_auto_cert.enable_path_normalize / c3651d0de746 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db19e1326d71337896660ff896d2ed26475546325c85ac6df21a29f83ea5d3f8"></a>

## proxy_config.https_auto_cert.http_protocol_options — proxy_config.https_auto_cert.http_protocol_options / afd5b485fbe1 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.http_protocol_options

<a id="canonical-8c8e0d531f4bbe24219dfa3f08c4cf6803f4a90ba3520f6c615b0d4962d82cb4"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-9a2acb6a94662366e70447647305149027d7ed4f9235845a1bc8e860848d5c3e"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options / afd5b485fbe1 / 3

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c0ed988ff09b02d4159a6807f7443a8e13cfd74eda56f62365ca47935daa23cf): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-94fa372b593c242990580160f02a188938abdb541d52eca3ab8f0109a55f5bc3): complete subsection reference.

<a id="canonical-08fd819ae8af3a1c079058fd33c70b3aa1012c3121a60559122a37d0be986b0d"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options / afd5b485fbe1 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c0ed988ff09b02d4159a6807f7443a8e13cfd74eda56f62365ca47935daa23cf)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-94fa372b593c242990580160f02a188938abdb541d52eca3ab8f0109a55f5bc3)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bff2360fd4d5a90f04bda0cac5daabaccb034f0832365229d9018f66c78d71b"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 308f8608c6e5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-7dfa1a82f6ed530298162a343af328933659ca2870d7ecaa2a1d976da5718c44"></a>

Type: `"single"`. Computed.

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

<a id="canonical-43968f4cdb5c0ff0cfa65181c5fd50f491e07d917a3e571326a747a406ea46de"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 308f8608c6e5 / 3

- [header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0): complete subsection reference.

<a id="canonical-86107d13f81917ed09be3980975efb468827e0f292dde6efcdd6bcef97b2663b"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 308f8608c6e5 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3168670aed649ca610a418808c740ee11c2f5a50714b6ff753a3ecb85f4bab47"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 35c2c5ea4be5 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-03d42450565ad2264a99a380d6ebd2afb5eaa4a8fdf4152a8ba6b5bc99e8abc2"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

<a id="canonical-0ea238392da3f63e6e97ff4ef8136646267e6d84c1e252689c671ede8b3d448e"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 35c2c5ea4be5 / 3

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-8b221a1c8153610a3121c15768dbc22c30e2c0bfccae4caa64298f0c110a0377): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-6bd80ef49b9e808e5eff5a50a3128e67da1d8c7d8615d9ebac71dbe897ca8465): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b6e7ff7abc36b6b026899afc309b959df11cf33c18f45322f56bae68d978e55f): complete subsection reference.

<a id="canonical-6aae79184dad183fe5ac05a03cd19581a71675ec18c7a9383f1b3f1ad31c744c"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 35c2c5ea4be5 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-8b221a1c8153610a3121c15768dbc22c30e2c0bfccae4caa64298f0c110a0377)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-6bd80ef49b9e808e5eff5a50a3128e67da1d8c7d8615d9ebac71dbe897ca8465)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b6e7ff7abc36b6b026899afc309b959df11cf33c18f45322f56bae68d978e55f)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-8b221a1c8153610a3121c15768dbc22c30e2c0bfccae4caa64298f0c110a0377"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4fff5aef2517853b5f2241a75fe6a6dcffc5eab1b5cdaf88b5cbcc77b60fee2"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / aa8925e1d4b7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-67a1e7ca424f152f15a604cb9e9ac6c7ea7ea89c75b63d0415f761473f4bf0c8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-02986f28f25b9d5552a33b971be707c4bf372785620857cd2fc00d796df1222a"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / aa8925e1d4b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6637374f56a8428c3d60d80ae6229ec05a3b4820d5a893bf80c1100e6572f79f"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / aa8925e1d4b7 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-6bd80ef49b9e808e5eff5a50a3128e67da1d8c7d8615d9ebac71dbe897ca8465"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3caefb735ed05688951373356c5ab86e894da202eb47d19d30841bd40df1dec4"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 9ff63d5f187e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-a57b478ee4cb6bb39562163494a4d20203ec92acb232ac0a08eba91e33e8b9ce"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7ec8a483415b3453e37662bad20cd45c8418ed1a898cc7321ff48973855e6e18"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 9ff63d5f187e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-92d23855c4cfca2b2ea730520a8f90e856643a18d27aff0f9c4df22258a7bd06"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 9ff63d5f187e / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-b6e7ff7abc36b6b026899afc309b959df11cf33c18f45322f56bae68d978e55f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24c462dc80fcd1300d0259ccd23c0e05983e763a507009c4bfc49919ad874cbe"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 59c834d53fc4 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4b55828b235c3ca0d12445500395d76b0b488efca9e9f4de2166d4c050b4cc6e)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-334154287faaa9133a6dabb103451b6b887245fbb1659c288bec322aef5ff631"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d1abc746b414bcab6815196862c36a044ca79454807579d90a38683207667f47"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 59c834d53fc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc1e7cc7cba5a7d5fea4ff797859f60d585e83416b831cbabda1aec208df520e"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only. / 59c834d53fc4 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-34406f596c18d7ce56823d277605b626ff76b9692fd46fa3b252d395ab3318e0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c0ed988ff09b02d4159a6807f7443a8e13cfd74eda56f62365ca47935daa23cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffb2df2279babf4f87258094001d03ba110199a314faad6abe8d8aa3b8c06fc8"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 32d60a234ec9 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-4c49064506ddda758ded0110b3d9161a8e3073cf3d9bbe238dff246afc3e35f0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ac559a6ac507d9edddbeaa8c6e39c6c5a03ab77ee9ded88718de0e24d36da723"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 32d60a234ec9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef31c78981956f800ac710a660221795e6b293a6b926343181b8918a3fa0f835"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 32d60a234ec9 / 4

- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-94fa372b593c242990580160f02a188938abdb541d52eca3ab8f0109a55f5bc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abed13ca2d5003c97aebe7a9c57a82edce759da40cd5a7738c21bf0d47d2a41f"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / 356e9665759f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-f0e9f8b0bb40b58745ae1b62c6160ae352b8ec117eda19ba20c5668cf2f0bca1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d2780b5908bb91091834e36c9139b04fd0530ebbe5e63facb380bfd2d49df862"></a>

## Direct properties — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / 356e9665759f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6805f3c5e431af137b3c2f7890c5c900c9a90081ae19dd0b27fc80d21bec1120"></a>

## Next pages — proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / 356e9665759f / 4

- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e1eb0545e193b2a83ad413167480c4b20e2c6376315643e04ac0a318ef330655)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ce45e91f1f6c9d5c3f42e470d48b9c9817f843d43f04992dd2420002e5c85eb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c36e2ba5e10c78f29518caa92007e71104de938cce92e692890d9a6d118228c1"></a>

## proxy_config.https_auto_cert.no_mtls — proxy_config.https_auto_cert.no_mtls / a74187c68220 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-bc9f5cadd0b39139fac8ac5bb8a65857a83c71b32770725c722785da1fbb4d48"></a>

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

<a id="canonical-738e41325d681081ab331f49f851068ceb4a7a4202a96a6b48161a91435757dd"></a>

## Direct properties — proxy_config.https_auto_cert.no_mtls / a74187c68220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4cebed8f138a962c4f92d506ac7330b7d76c435a33837a134a80d3e5db4da4b"></a>

## Next pages — proxy_config.https_auto_cert.no_mtls / a74187c68220 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-325da1dca17116c529debc17ba83bb9f3ed9033337561b2708f03b79ef1d0aaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ba400fc9d8a3918dbf3ce3186bea214ae5ed553f3a86e7d440021b375b4d7da"></a>

## proxy_config.https_auto_cert.non_default_loadbalancer — proxy_config.https_auto_cert.non_default_loadbalancer / c0ea24bce70d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-b6525bfc55510fbb7cfccd920e0764fd5408eb2a2289f8c13db5fd69f34aa8f7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7fbdda56dcf39f443b06c25bc3c324a3f5429138a923aa56314cbec7fc530bf4"></a>

## Direct properties — proxy_config.https_auto_cert.non_default_loadbalancer / c0ea24bce70d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb4c91f021e28f42757b0ba924645a13ac16d79246845b5c61d018cce8bfb7c6"></a>

## Next pages — proxy_config.https_auto_cert.non_default_loadbalancer / c0ea24bce70d / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-6a5ce9c3be3ed3a2fd17a783b3ac61e842dd3e304efba117a370e43d36af3207"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8271ceb3e7d410c1b3ef469d5aea924b367c3b9947ac89030d2d165e9a23c7fc"></a>

## proxy_config.https_auto_cert.pass_through — proxy_config.https_auto_cert.pass_through / 7fccdf4128cc / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-c6e3563e7bc2d3575a8f74ceabae2d560aed7d9f4fa4eede30f5627af296bc9f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fced15bb391cc929947b86ed9bc2ca92fc625c7350d1e82f9a003c493785c050"></a>

## Direct properties — proxy_config.https_auto_cert.pass_through / 7fccdf4128cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c376a739941c9d8b3c3e38dfec80f94811ab0b1b10248e9c779e9d5c18fd8106"></a>

## Next pages — proxy_config.https_auto_cert.pass_through / 7fccdf4128cc / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0466f922bf76e794a2b5407ced1db6cd3f38ed95dc2ba5f41dc719effc45c7e2"></a>

## proxy_config.https_auto_cert.tls_config — proxy_config.https_auto_cert.tls_config / 4afc10679535 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-a35e8126cfda6aadcc344ef3307215e5672c81e249e1ddc35b2547d32a8b0599"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-7ad0da4383803da4f0d9ee935d86fafb32286f6e12f0c9c0a926783ec870e442"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config / 4afc10679535 / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-30a4b0df60e6f968fb31723c6659664e07a6c3dfc9602224f54a21573d566d53): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-467d6278dc4106ad25b0222f3abbb93c0646812d54b8d30e364a947de8aeea00): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d2f9d69085f676a076e2f72e66d9d31df1ca1ce6f6ca05acfad296e47e59a5d7): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c53265400dd77234b37519321fb26d035ca37a87567eeb86b9801fc980cb1d95): complete subsection reference.

<a id="canonical-8d9072e8239a1f87d9b1a2c452b3ed85bf10f7c7d5571e9d92039e95c875671e"></a>

## Next pages — proxy_config.https_auto_cert.tls_config / 4afc10679535 / 4

- [proxy_config.https_auto_cert.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-30a4b0df60e6f968fb31723c6659664e07a6c3dfc9602224f54a21573d566d53)
- [proxy_config.https_auto_cert.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-467d6278dc4106ad25b0222f3abbb93c0646812d54b8d30e364a947de8aeea00)
- [proxy_config.https_auto_cert.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d2f9d69085f676a076e2f72e66d9d31df1ca1ce6f6ca05acfad296e47e59a5d7)
- [proxy_config.https_auto_cert.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c53265400dd77234b37519321fb26d035ca37a87567eeb86b9801fc980cb1d95)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-30a4b0df60e6f968fb31723c6659664e07a6c3dfc9602224f54a21573d566d53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be03c7a4bb2c18420ab149f371f14729a16525a541fd89ad3bcd1df540d9c6db"></a>

## proxy_config.https_auto_cert.tls_config.custom_security — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-c1ff9624bea553594b16f98585cb5a011521d4b1c8272ed48883febcc1ea71f4"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-d8d8347ae8aada38f17b925f37fc93701c4674baa459f3cb1ce4413aac9d3955"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 3

<a id="canonical-6af76e7ff875a52d9f5334e5c4f894b9f784ea2c616095ee0474ea581e9a87c1"></a>

<a id="canonical-cf9fd33f577c755c5a0ddcc35bdf9615e21d217650327e25140393802e0371ab"></a>

## cipher_suites property — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-fe871a0a492246a95eaf3f9d3e612d72ff2e3682ec13c3f881444406c22456c8"></a>

<a id="canonical-da7ef9e667cbf837ff45be8e4a37241292ad790482888fcef2f329e56bf56e05"></a>

## max_version property — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 5

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

<a id="canonical-69fd529f4ad4e69d4cba3f58682852528acaf23f5c4b98e5b633a4dbe48b1bb4"></a>

<a id="canonical-d5f43f69a28e36540f2596fdd550a238b2293580315f596077de6d3cc3a293ae"></a>

## min_version property — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 6

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

<a id="canonical-c77395ef5fccabfe79b7ab7ed346b10a6fc4bbe5eaa92f0d187091bee96e8e26"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.custom_security / 8b431c283890 / 7

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-467d6278dc4106ad25b0222f3abbb93c0646812d54b8d30e364a947de8aeea00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2adcf17656666222d3953515c2b4659ded78baf994c8e964cef8ddbaa9bd795d"></a>

## proxy_config.https_auto_cert.tls_config.default_security — proxy_config.https_auto_cert.tls_config.default_security / 414849ae1285 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- proxy_config.https_auto_cert.tls_config.default_security

<a id="canonical-5cd7b6fff6c17f0d8945190fc6a51417e51cd6b1510dd999d3750e732ed93cc1"></a>

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

<a id="canonical-36cee9d50bce440f9c3b65cbd3e052874f42b2192e8e9805a2ba9480374f41f0"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.default_security / 414849ae1285 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe85d8b976d0eb6ac637500bd5fc35b49759d982899594200e933dffadca1629"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.default_security / 414849ae1285 / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-d2f9d69085f676a076e2f72e66d9d31df1ca1ce6f6ca05acfad296e47e59a5d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4c3be1970cb534544d7b9ca4696b8f1d702a9ba1ba6f669680febbfaab1b809"></a>

## proxy_config.https_auto_cert.tls_config.low_security — proxy_config.https_auto_cert.tls_config.low_security / 148f8ba129eb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- proxy_config.https_auto_cert.tls_config.low_security

<a id="canonical-51cd6c49c397113c7d9654f076a5b1285babef9a3c194293ebe1f99089939649"></a>

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

<a id="canonical-907aeca8f7969769deedf82c73680e1f24406031b1a00820c1df6a6f1d47b044"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.low_security / 148f8ba129eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1823bcc1b5e857fae9afebe8ad186164ca39947cd96b49b6800b67588a216adb"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.low_security / 148f8ba129eb / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c53265400dd77234b37519321fb26d035ca37a87567eeb86b9801fc980cb1d95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f770d843b51c387953dea702d4f52de010b76324abe34726de4289aa98dbc3c"></a>

## proxy_config.https_auto_cert.tls_config.medium_security — proxy_config.https_auto_cert.tls_config.medium_security / b1c4c1f2cd44 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- proxy_config.https_auto_cert.tls_config.medium_security

<a id="canonical-9b7ce7bdd92207f2f268d8468ecde3f620d93bf316fd0e4c2a6d02e497e27c54"></a>

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

<a id="canonical-4dd4a3d2290eb8e13a7141cc5e37ba4ef8674016a56c355cc25a4c9110868486"></a>

## Direct properties — proxy_config.https_auto_cert.tls_config.medium_security / b1c4c1f2cd44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61f514c85896e8df596b2bdf7374cf6f24e27f1ef9ba4fb0bca990433d506952"></a>

## Next pages — proxy_config.https_auto_cert.tls_config.medium_security / b1c4c1f2cd44 / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0a431e4faf90f5096d315b8392392574ef33d7be13b407829d72100abba5c16e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0015789679519de64f7eda975a7109eeee317d669ececc85ebf433700566a83e"></a>

## proxy_config.https_auto_cert.use_mtls — proxy_config.https_auto_cert.use_mtls / f4fba70cbf57 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- proxy_config.https_auto_cert.use_mtls

<a id="canonical-d0594eddbafbb60eb094435280b88c5d9924a63e3c2f06196e1240709e017e3e"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-895006a62a3c6714d72a051a10f79aaaa39670a23ad990166d65af21af7ccd0a"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls / f4fba70cbf57 / 3

<a id="canonical-b336fd6da0d447195d74038f94238cd41e451ef8180fccde78458702174ae638"></a>

<a id="canonical-f10241244569619b8055b64bd6afeecd80a50371226dc258922727eb8b863579"></a>

## client_certificate_optional property — proxy_config.https_auto_cert.use_mtls / f4fba70cbf57 / 4

Type: `"bool"`. Computed.

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-fe68fe80b7e64bb6e030c0c62a9d2acda5309d636b10156c888c5f8fa13d9f20): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-5a8f49270da7361572ea94dac3c9abb0ac0ef12e23524d1d0ad16ffc9f475c4a): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-7255bfeec5767d303d7a3b7663af2e2d4ad4a2ca8d0e70161de4f1031ce87e9e): complete subsection reference.

<a id="canonical-4ce2812effc4b4261a9a4f34ddc76c24b6688284d43dafe808faf25b97050355"></a>

<a id="canonical-d24147bb05721ed7e9edc5c5a50b77c85a4ca5a6efff162e7f5ce56e5143e416"></a>

## trusted_ca_url property — proxy_config.https_auto_cert.use_mtls / f4fba70cbf57 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-79e4d205a46d55dc1c707eda6fbd37bb6fdde2d8460c2558453cccb694671fe0): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a895cbe2d15d347828fe411545571d9ae855a79b8306ce708820b7e2e353489d): complete subsection reference.

<a id="canonical-fc648d8894f84132a053e3afe159b3feeecab853700dc95e2b0cf1a7e307e12e"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls / f4fba70cbf57 / 6

- [proxy_config.https_auto_cert.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-fe68fe80b7e64bb6e030c0c62a9d2acda5309d636b10156c888c5f8fa13d9f20)
- [proxy_config.https_auto_cert.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-5a8f49270da7361572ea94dac3c9abb0ac0ef12e23524d1d0ad16ffc9f475c4a)
- [proxy_config.https_auto_cert.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-7255bfeec5767d303d7a3b7663af2e2d4ad4a2ca8d0e70161de4f1031ce87e9e)
- [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-79e4d205a46d55dc1c707eda6fbd37bb6fdde2d8460c2558453cccb694671fe0)
- [proxy_config.https_auto_cert.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a895cbe2d15d347828fe411545571d9ae855a79b8306ce708820b7e2e353489d)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-fe68fe80b7e64bb6e030c0c62a9d2acda5309d636b10156c888c5f8fa13d9f20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad2f5f4e9283340a3290b431c1af3d02db95bcbb9734c13813c59569bf88657e"></a>

## proxy_config.https_auto_cert.use_mtls.crl — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-aca153d795edc745a50e74983fb0ccc156f05684b0079826168b36cad00e8b57"></a>

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

<a id="canonical-4023216f4640d05d021cfb2fdd087dd5158c7fac0ac53c583151abbc4ef96e1d"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 3

<a id="canonical-30c01d252bcb923cdc8d468646a08faec63573709ced712fe5378bc856dc2a09"></a>

<a id="canonical-c58a19186ec78596c5c68ed14193e0233abbd381339000eb6b74e460e76f0e41"></a>

## name property — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 4

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

<a id="canonical-00a8f0bf60ef21ea28dcebb8633175d5dd09c30f3b822831251765966ea906a6"></a>

<a id="canonical-2346a1e2f79ea16cb454dd7f5bc9126b55ef1c82fac5c8dbde31d625812191b2"></a>

## namespace property — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 5

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

<a id="canonical-64736012219c8c50f99730a24c368b044b7347f2266e951ff3814d2c91387aae"></a>

<a id="canonical-557ba169543a9281e28883c7194f59f8ce4caff48a752a13216baddb6ef6d4d6"></a>

## tenant property — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 6

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

<a id="canonical-8100b6de61946163d5d7030c770f5cd390dd9ba08185565e77b16f3b5e95acc3"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.crl / e79daade2453 / 7

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-5a8f49270da7361572ea94dac3c9abb0ac0ef12e23524d1d0ad16ffc9f475c4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb3ccc7e0c2e0869efc5aa914517ebcf37dde03a81cb267ca06cd1d1e48662a7"></a>

## proxy_config.https_auto_cert.use_mtls.no_crl — proxy_config.https_auto_cert.use_mtls.no_crl / 914240199e48 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- proxy_config.https_auto_cert.use_mtls.no_crl

<a id="canonical-c1e718b10e3ecded681ca3c1d8fa663555fa1e95668b2b8f6f5f96959e2f427a"></a>

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

<a id="canonical-931a82da3f285e266225266bcaef729ca7348e2491eaafc510b411bcaac4faac"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.no_crl / 914240199e48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59285090cf679f6304a07ac8e57a83e05d85b377da6870efe9f7dd61df661dfe"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.no_crl / 914240199e48 / 4

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-7255bfeec5767d303d7a3b7663af2e2d4ad4a2ca8d0e70161de4f1031ce87e9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52c4c53ffbe01faa99c983126effb70987081dfa19e5c1ac3164c2bc56b1bc5f"></a>

## proxy_config.https_auto_cert.use_mtls.trusted_ca — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-808a9d760b6e2a15bc2ccfa325774e97629d00dc834cc3d816ffc2cb02b6c21c"></a>

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

<a id="canonical-ba010e6e9a41b0d55f34d58b5888422a09ce1a32a5cd4339d7527e3573144245"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 3

<a id="canonical-901879cf5b60ec2ec06d6b8675699ae758106005e45d35be09091d891b8eda21"></a>

<a id="canonical-a5458b09c579f8568786f2b1696865b76128c114670326d9fa130f8ab80b9360"></a>

## name property — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 4

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

<a id="canonical-fb8780a6919416e2aa5954d2e5412401e17d2b4edaa9164af714e4b9d5acbe56"></a>

<a id="canonical-b9cf2dba73e618a905993f0d921e1c4a102d7e2f8a7e6a9d5d1de476883bfe14"></a>

## namespace property — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 5

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

<a id="canonical-00c14dcbd46aebc2cdaa86a92653c46ae33abaec81e00d843e82ae9b6a9ed111"></a>

<a id="canonical-675a655586a96818985f5223a59fe679c48e2604ee19d13adc37db9499b71e72"></a>

## tenant property — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 6

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

<a id="canonical-1d64fb85f4109e9573fec63eca690907c4701baff92c4ae9f4381cb542d8f06d"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.trusted_ca / c691fa6028d0 / 7

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-79e4d205a46d55dc1c707eda6fbd37bb6fdde2d8460c2558453cccb694671fe0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63555dbc045f65217c1c2c204948a2369e99fb002e69398da0b887d426ab8e5e"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_disabled — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / 886dd7679470 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- proxy_config.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-f0a43bbde557cb445f489e69d7bd8a58ff91827490fe8607c73dbd694c4ebd0f"></a>

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

<a id="canonical-ccf7551271a4eefea2723fbeccba0d2e14cad61e1fec44b525f904fda540c4f0"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / 886dd7679470 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76551bd40019a9e4ac1aeaa8a0eede3571e38be4dcb2b5354690a77109252dec"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.xfcc_disabled / 886dd7679470 / 4

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-a895cbe2d15d347828fe411545571d9ae855a79b8306ce708820b7e2e353489d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d27595c95a573d53e9191519ab34cd86ccc9fa40df27046704abf3dc417da60"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_options — proxy_config.https_auto_cert.use_mtls.xfcc_options / e3e5aa35df13 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d31637bdd4b16cbe32ef66cceeb81a16f39e53c629c297d234c78d6c5b502e7c)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- proxy_config.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-8fa2d528f11ea7f6af06a9e047eebad3e86971d16755f9a1887a5a7b6db64afa"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-c790d9666bcb834cdcc272cad1c9b9e0e6641dc45b789f7da3ff47fbe4aa4125"></a>

## Direct properties — proxy_config.https_auto_cert.use_mtls.xfcc_options / e3e5aa35df13 / 3

<a id="canonical-d21b30765a9e3839590c3dba743bcb2e331387dd3852aaf91d367ee478c63829"></a>

<a id="canonical-d523ef8e2f9762dc606e8daabba0433e1125c5ff8a699de5ab1ecb18f332f78c"></a>

## xfcc_header_elements property — proxy_config.https_auto_cert.use_mtls.xfcc_options / e3e5aa35df13 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-5f253b9abc3cb6d980e95e72ed8f9f2476e53f0dbbebc0856b1849dfb57fd53c"></a>

## Next pages — proxy_config.https_auto_cert.use_mtls.xfcc_options / e3e5aa35df13 / 5

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-56df24338ec3d07883a6d5866593e988267c68f112a03c91ca654a8cc070fee0)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
