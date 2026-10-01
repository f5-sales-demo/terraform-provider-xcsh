---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-6b8061054fc78b5fa2d4e011cbd7b5dbaf083bfbba3916d9c46d29be4b52ac6e"></a>

## Direct properties — origin_pool.use_tls.tls_config / 67321f3fc13c / 3

- [custom_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ff88ca073d2bb1b68604fae15b4cf7c20abb5089ff6ccd4cc82d4431c290637f): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-523d8af062da29a5d6a72f732261dcc136d1fdc6fea4267cb8ac633b9930ebe1): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-8eb79f6ac2e7a3e4f1cc8a0b3be69e084bf1ac85aa4b028d84d0a8b28a6e7185): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2f2b3a52eeec1eb07383083164d78abc957518a3c7dedb17556ebc3dce9fe46a): complete subsection reference.

<a id="canonical-f33fa7b7e0680d8d4a6f51b9c03ca7ca57420c98da5740c6eb72e2e04cbff124"></a>

## Next pages — origin_pool.use_tls.tls_config / 67321f3fc13c / 4

- [origin_pool.use_tls.tls_config.custom_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ff88ca073d2bb1b68604fae15b4cf7c20abb5089ff6ccd4cc82d4431c290637f)
- [origin_pool.use_tls.tls_config.default_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-523d8af062da29a5d6a72f732261dcc136d1fdc6fea4267cb8ac633b9930ebe1)
- [origin_pool.use_tls.tls_config.low_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-8eb79f6ac2e7a3e4f1cc8a0b3be69e084bf1ac85aa4b028d84d0a8b28a6e7185)
- [origin_pool.use_tls.tls_config.medium_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2f2b3a52eeec1eb07383083164d78abc957518a3c7dedb17556ebc3dce9fe46a)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ff88ca073d2bb1b68604fae15b4cf7c20abb5089ff6ccd4cc82d4431c290637f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e4a9ce6e0ec433883ca27e321535aba5203e6bb67afc599d49dbe5936c7673e"></a>

## origin_pool.use_tls.tls_config.custom_security — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- origin_pool.use_tls.tls_config.custom_security

<a id="canonical-8ecaf5c5dbbbea782a5d26c2c4564fc796d385dd3f85a95e92d00fabf5779ed5"></a>

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

<a id="canonical-1a636f0e3cff86887be102de9277636d3921c6b6719f2caea8a276a402bef448"></a>

## Direct properties — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 3

<a id="canonical-401595cec8a7d8921f903a2018237e74c3c9321ed71c151afd4ee039bfb86e10"></a>

<a id="canonical-b997c524433a8f1654243947ab5abcd2b414b3185b7b82d18548ad51a07ca728"></a>

## cipher_suites property — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 4

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

<a id="canonical-1306453c905a813dbf4c55fb54c96b3d1a47d95f1ab22fa7c56a66c187058cab"></a>

<a id="canonical-3683198b36b6921555b2f9f105de44cfdb3254ae4ec3e6ea0f2fd0e72514ca18"></a>

## max_version property — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 5

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

<a id="canonical-accea31e2e1aff9114ec91d1a648c9f8b94141f18600141531fbd9e72b6906ee"></a>

<a id="canonical-2a5c7eed94b784965b4fdaf3e9fb730992739f62941839e982f94d82c0fcab1f"></a>

## min_version property — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 6

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

<a id="canonical-64c65c65bafd48fa19e6a8c76bd947c2c75bac99df5736e506976763d0b8cf18"></a>

## Next pages — origin_pool.use_tls.tls_config.custom_security / 3327517fd9c7 / 7

- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-523d8af062da29a5d6a72f732261dcc136d1fdc6fea4267cb8ac633b9930ebe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34b7c13eb82bae504a4843a73450c9b2b15163289249c395952a68900fcf2237"></a>

## origin_pool.use_tls.tls_config.default_security — origin_pool.use_tls.tls_config.default_security / 8a1d301cdc06 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- origin_pool.use_tls.tls_config.default_security

<a id="canonical-f39afa7f2f81b260ac1becddda42e40d7a4de626b06d3342a14768693c4d71b7"></a>

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

<a id="canonical-071076193de6733c2d1d2e78d673257664d5361aa345554a4856cee6ffd98d5d"></a>

## Direct properties — origin_pool.use_tls.tls_config.default_security / 8a1d301cdc06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-427648822a73a665e647893c9f7e678d863a4a8764d3832f551d1a85eb8b85f2"></a>

## Next pages — origin_pool.use_tls.tls_config.default_security / 8a1d301cdc06 / 4

- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8eb79f6ac2e7a3e4f1cc8a0b3be69e084bf1ac85aa4b028d84d0a8b28a6e7185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b00bcbf0529e98035c525d202ee7446ff28600354cfcc88527f7a4d80b4dbb3"></a>

## origin_pool.use_tls.tls_config.low_security — origin_pool.use_tls.tls_config.low_security / 075456ea75d3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- origin_pool.use_tls.tls_config.low_security

<a id="canonical-b35f100773010222a3b0dfb2355852772a4c3ddd1842672c84add397a81197f6"></a>

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

<a id="canonical-8e4e54ff1f32f3007d299b534005ba2d5a1582ea0860ab8b9d3061b113c90fd9"></a>

## Direct properties — origin_pool.use_tls.tls_config.low_security / 075456ea75d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44f350df6d9160e9b94b48650459526dcc8b68736b9d7ca70a80f0d2cb4e4daa"></a>

## Next pages — origin_pool.use_tls.tls_config.low_security / 075456ea75d3 / 4

- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2f2b3a52eeec1eb07383083164d78abc957518a3c7dedb17556ebc3dce9fe46a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4b023d993d97bdda579cf6e72bccc9f4615d2091930320d2b3e33d3d8747a2c"></a>

## origin_pool.use_tls.tls_config.medium_security — origin_pool.use_tls.tls_config.medium_security / ee17460429c1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- origin_pool.use_tls.tls_config.medium_security

<a id="canonical-31c834bd3996b447a6cc128581de1cbcf366a85718c612aeea813d8df144e514"></a>

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

<a id="canonical-eb345e82e971d1b9a5b2ba4377344ef1981c45a3a0be081e3fecf0362cc9dc6f"></a>

## Direct properties — origin_pool.use_tls.tls_config.medium_security / ee17460429c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b48a504171e11a39517808c5af85e29038e25d8d2cbfdbc6ab226ca4b656db3d"></a>

## Next pages — origin_pool.use_tls.tls_config.medium_security / ee17460429c1 / 4

- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dd65c566b56ea97b3d0a4838833b3a6f009c6b9c5ef34d0ae1e7226350f7322f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fba5e18dc641057221de1e12b77a46502cbc52dbf0719ebffb46de4400f6f16"></a>

## origin_pool.use_tls.use_host_header_as_sni — origin_pool.use_tls.use_host_header_as_sni / 87342940687c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.use_host_header_as_sni

<a id="canonical-522ff71c320631639eafa553a0e42e08b26abdef3a873a688e9c4f429273b2e6"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-db90232e5dd3522494d80f58275bb8d4504bd45cf9f44ef84f7aae1f82385911"></a>

## Direct properties — origin_pool.use_tls.use_host_header_as_sni / 87342940687c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a482e7e03f95bb93b996be9d409a8eac56d857975ddda5be8bad7a2084d465f0"></a>

## Next pages — origin_pool.use_tls.use_host_header_as_sni / 87342940687c / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4ec6134fb8766de4e4403ab1fe23f827d0291d86cff1156bd46762220aec64f"></a>

## origin_pool.use_tls.use_mtls — origin_pool.use_tls.use_mtls / ef834e1ae7c7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.use_mtls

<a id="canonical-6e33d4a0ef7e4167db1561a75755709505f18d8e82c5b3c45a130cc1e2297b33"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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

<a id="canonical-0530e0a0c8c5acf19ddba96e4bc0b4c959886d135737176b78ce18ab3d312be7"></a>

## Direct properties — origin_pool.use_tls.use_mtls / ef834e1ae7c7 / 3

- [tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f): complete subsection reference.

<a id="canonical-5ba3496bf2ad1e32a11b9da06f5843f338d04ccd8690e488507f73e457a7ef48"></a>

## Next pages — origin_pool.use_tls.use_mtls / ef834e1ae7c7 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4cb9b1faf1f0f058b4fbb30a9a31b7bc577e2885f77b68f609e3fb977e732c6"></a>

## origin_pool.use_tls.use_mtls.tls_certificates — origin_pool.use_tls.use_mtls.tls_certificates / 5d91f93a4e16 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-f719c2ddb749ecf9c3cee3d15f237efd997fe4128659642f2abb7f27f54776f5"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-689ea81da56d73dbac9d06b6557f4fe9132f07b9766931f343f872871e31a5a9"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates / 5d91f93a4e16 / 3

<a id="canonical-86588c628381a8cd63ec7c8beaa37af49fa173a3d97b8f25dbc7cbbe2a57ee0d"></a>

<a id="canonical-55fc6f7a2e1f1dbd1a1e87e98fd5d2a35d285c9181e83bb385cf06b1934d04b5"></a>

## certificate_url property — origin_pool.use_tls.use_mtls.tls_certificates / 5d91f93a4e16 / 4

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

- [custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-f7ec7f3495fb7a0b0a0b7ef33cb4d70d2941e5414c0802d92a0ee6cd9ee3b0de): complete subsection reference.

<a id="canonical-650632abf4d2b8383656bfa9faebb01d9b8197f91ebc4f05f7a9d1d6c12826e4"></a>

<a id="canonical-aa806a764971b6c26d4d5fdcf1b7a2f278ddedec3258e4b194c7030b79d99f93"></a>

## description_spec property — origin_pool.use_tls.use_mtls.tls_certificates / 5d91f93a4e16 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d58681fb4920f43eba6594984a4568aaba3eefdddb2f81ef18c885d7571d27c0): complete subsection reference.

- [private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7): complete subsection reference.

- [use_system_defaults](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5564a47f7f07b64d2eb2df9b7c59ca7e011a4dfc66f90cd4f999de7bf0ead5eb): complete subsection reference.

<a id="canonical-26f4d28b8ba43fd4fd03a606c8b10507fcf7ddba48e346d2c6b963f58d557c0b"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates / 5d91f93a4e16 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-f7ec7f3495fb7a0b0a0b7ef33cb4d70d2941e5414c0802d92a0ee6cd9ee3b0de)
- [origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d58681fb4920f43eba6594984a4568aaba3eefdddb2f81ef18c885d7571d27c0)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7)
- [origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5564a47f7f07b64d2eb2df9b7c59ca7e011a4dfc66f90cd4f999de7bf0ead5eb)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f7ec7f3495fb7a0b0a0b7ef33cb4d70d2941e5414c0802d92a0ee6cd9ee3b0de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edbf8dda603204086acc2de2e589cc8e236068c5611d3027084f098d17a81a21"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / e93c0aaa3adf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-431c5b2ef342a471ecf5cfca90ad67702dae7c2cbc4de4a736fee96b68af61f9"></a>

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

<a id="canonical-16509ce0deb868881119587787050fb3a4077f2a32b990527669b9010eaa439d"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / e93c0aaa3adf / 3

<a id="canonical-33405b6fee38a8bf0fa325cbb8f37f51d5f05453f5579259f3cedff2aec922c5"></a>

<a id="canonical-c7db5600c77ab31330988085327e3854ea609dbde80044dde06cc54b7e43c3c4"></a>

## hash_algorithms property — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / e93c0aaa3adf / 4

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

<a id="canonical-c537c1554ab2940e5439b130c59fa2ba887838ac7cfc8ca813461faedf91c499"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / e93c0aaa3adf / 5

- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d58681fb4920f43eba6594984a4568aaba3eefdddb2f81ef18c885d7571d27c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-862b88b757432099e8244626ae12d0e89faf5b8b90ca41ee668a7b22791477bf"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / ebf4ba249552 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-5598ac38a1dee8c6114da5cd34e62b3c317b5200cc061969ca523e334fb307fa"></a>

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

<a id="canonical-bded5882a7008f6df498b38ce230820aedf94535c901a17922302c9561f65cdf"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / ebf4ba249552 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9583105303d789126063005d7b339959a10bebc1d74519527b9160acf4a38813"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / ebf4ba249552 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29aa29b7bfa299f8489d21662d546a16aa1f1321b97433aaf2e42219427b6a61"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 21b80110a417 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-8ed0263270cb223655bc3d86c088eda2095a9fc0b39bf552a9a22338098d46d5"></a>

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

<a id="canonical-f1194a7c9311a6d6560f560ebc659375d4a05f70428576f3201ccfd2e5b1f169"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 21b80110a417 / 3

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-91de9f339029643eb40b7d04e0dca49d1a00cf68918f132826224a2d6d68f340): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-21ceb61a79672b6328a1375d43b0c5ce3c8adb59bdc415db2f0fa60107b3f93c): complete subsection reference.

<a id="canonical-baaf36d06d04c5935e0b498b0696b81e3bcd06b72da9590a8cbd4474d4da7a45"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key / 21b80110a417 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-91de9f339029643eb40b7d04e0dca49d1a00cf68918f132826224a2d6d68f340)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-21ceb61a79672b6328a1375d43b0c5ce3c8adb59bdc415db2f0fa60107b3f93c)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-91de9f339029643eb40b7d04e0dca49d1a00cf68918f132826224a2d6d68f340"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6770b577a875844087b385542bcef56de92827d9780599879f3be998baa30619"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-b4f457ade74c043baeba540d40fbcd08a0cd3e3d03d889cd4423d38f88b13f9d"></a>

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

<a id="canonical-9ae4b74f50504a9517be0ef26e5a22f577d185ff698243b42299d8acf52cadb2"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 3

<a id="canonical-9caf30787a31509c5251a430f2acffcab6f05f0b17d781fc8d7032d3830df425"></a>

<a id="canonical-79bdaec8ec4d341fac419a4999dfea74c138a82e7836a4abb78e4c5b3a59f396"></a>

## decryption_provider property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 4

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

<a id="canonical-2155364c1b2e05f39279a985ecdc8b869040f27eb08906e2b82a7d5c5640226a"></a>

<a id="canonical-ac47d3cac788f691850e1097a4ada14af67012b60ea6c574d2f9de6188a1a6bd"></a>

## location property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 5

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

<a id="canonical-efb2883d01c33f8470dbffecb76220f2bdad4b716a3c56f16a1e505d25aa0b3c"></a>

<a id="canonical-3b731344b3cbed6545e9efed6350568aef15f9edb2a6e0ed74501d5dce7c7164"></a>

## store_provider property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 6

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

<a id="canonical-06323ff56d67f8a568fc320af56e53d4ad26a871d64051363f91c15ee0afd395"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 95d4376aba4a / 7

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-21ceb61a79672b6328a1375d43b0c5ce3c8adb59bdc415db2f0fa60107b3f93c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1be12ba7f708c58fb76f68791d6d3e69b93bd282d602371d60b22c81b1846ea7"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / dc3db5b322e0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-5c1e84381e1409b64b7fcc775408b7fdf3896c770ec1b04cba5adba36df28c42"></a>

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

<a id="canonical-bb6ec43a1a1fbe64cda6f652530973f2a31ffce125f08aa5dc229f3d033bc779"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / dc3db5b322e0 / 3

<a id="canonical-9ba98a6ade47c9611f1ff69a696305e6d2beb5797e2b848f0b275a989d3b596d"></a>

<a id="canonical-9d6a65282317d9b77195b39bf5067a4b00bb184b3aea6f25bc75b827a9836d9b"></a>

## provider_ref property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / dc3db5b322e0 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-cc0a0e4754b8046c949a9b41ba277273170f0c49176e6edee5bc387880e05b19"></a>

<a id="canonical-e858d376b52434a0a9a6794f311aec54e675e90eaaa428273a7a2b6d765c4cac"></a>

## url property — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / dc3db5b322e0 / 5

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

<a id="canonical-09df91e32aa4c559a7968e4fa0ca025a71561db238af8b40050cfc67140b9dd0"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / dc3db5b322e0 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-5dd38d1e7a179801a80de5c9e6d7678db437c3d53f706c3eedd585a11ab8fbc7)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5564a47f7f07b64d2eb2df9b7c59ca7e011a4dfc66f90cd4f999de7bf0ead5eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-277d79417e38214f7b6c9140c595e5cc034fa149117764ab23fbfb5f99ac8f54"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / 115053acef23 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-61b34a9eca65c58313f2bf8be6654d22a890271db114b514d5077722e4b57cd0"></a>

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

<a id="canonical-ff47a2cb5f531906eadcd6653381a541027d4d518a7e6aba8fbea32d8447acb6"></a>

## Direct properties — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / 115053acef23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e4aca316c5266ed37a8fbc41cdecf9d5ba0d10fccec90f57ff3e90cad8f4bdb"></a>

## Next pages — origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / 115053acef23 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3ef3f25a8435167a72580ea30e42ed234b4704fabd806d01a94898d4a073b14f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ee98a678519e4f1a9cda4c6aa3ba19b6f2dd8396c25a322c9fbec0cf854492f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9797837d416287b4f7935b9ce8fdb1745fa20b33abf6f4dabeb711bd41823613"></a>

## origin_pool.use_tls.use_mtls_obj — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.use_mtls_obj

<a id="canonical-a6a9165941c26c0fed6e51cf00c9d2a430f0115fcb0b344cd83acfd96f4f0e50"></a>

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

<a id="canonical-1ea003e1b08e8f821c8bedd6e45899022a0ca729ac4bbf356961b615e82f95bf"></a>

## Direct properties — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 3

<a id="canonical-92eff0771872cc6838582235807b8cc473c9803159654756bf2e3d814f3b33ef"></a>

<a id="canonical-cedd14d28b5e0d7464a8c5f03f710f34daab05227b4a9dea8ad1b3a2888f9acf"></a>

## name property — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 4

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

<a id="canonical-c0c24fea972449723b1a29bda41eab6740c59481262ea08d192658a2e1796543"></a>

<a id="canonical-26bd0aba3a223fd2c3eb3961d2865e55e16ecc2f8223a50c8a7cbf96c1935b23"></a>

## namespace property — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 5

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

<a id="canonical-db8bdb536fde13c2f4779ba72053c6af8be988a376b6b2e8c3fd2251de3239a2"></a>

<a id="canonical-b40c91e193f928b61daf6bcc28cf836e255b30537321cbe8d0fbdaefa898a443"></a>

## tenant property — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 6

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

<a id="canonical-14af0cd55a6796846295620b62ef8e08ff6f90d7f7018fbb644bdadf39e9e1c4"></a>

## Next pages — origin_pool.use_tls.use_mtls_obj / 30f12313dde7 / 7

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d6d3e059c3665f643b999fdf3812795ba2895e20bd5b1ae4ea731267fdd56184"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-badf860806c8188fa8b473c904d5d78521c2a8e728999937335e09b959c77d8d"></a>

## origin_pool.use_tls.use_server_verification — origin_pool.use_tls.use_server_verification / ec4c0bfcb9b6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.use_server_verification

<a id="canonical-3007140a3e9db3ba72408a290c87ec8b816c80e4c04ce0278e442e7c1db391f0"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

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

<a id="canonical-f1190c6f26a6079ebbb24c1aae11044faf0b95b4e26f91ff3d85bf2ac5f5d27b"></a>

## Direct properties — origin_pool.use_tls.use_server_verification / ec4c0bfcb9b6 / 3

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-bd099fd6fa93d6e3f391602e615888f4fee9ee0d7308e9b552ce18a647fe6234): complete subsection reference.

<a id="canonical-77bf0372a5d08b0745c8c20b3f7290430bcb5319725a9c42864d4c61af570db8"></a>

<a id="canonical-9e95d44f1e96dc4c3053131fd0979e2763eb0efdfd1bfff310e215bddb347c0e"></a>

## trusted_ca_url property — origin_pool.use_tls.use_server_verification / ec4c0bfcb9b6 / 4

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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

<a id="canonical-65de75d731babd5434e2a9816b8de9829c9757bad7414fcccdc26f055c9a1742"></a>

## Next pages — origin_pool.use_tls.use_server_verification / ec4c0bfcb9b6 / 5

- [origin_pool.use_tls.use_server_verification.trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-bd099fd6fa93d6e3f391602e615888f4fee9ee0d7308e9b552ce18a647fe6234)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-bd099fd6fa93d6e3f391602e615888f4fee9ee0d7308e9b552ce18a647fe6234"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d66d16c3e0cbe72c06bf352c0db01e40dacf193b888b343409c8e1486636bee6"></a>

## origin_pool.use_tls.use_server_verification.trusted_ca — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [origin_pool.use_tls.use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d6d3e059c3665f643b999fdf3812795ba2895e20bd5b1ae4ea731267fdd56184)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-a11aa11f0331f12833d19d948b3b7c79844215de1b2fc63472e895ca60638c40"></a>

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

<a id="canonical-539a409108e777941e2a664816789c2570db2b44e136d9fa55a423d4b468d5b4"></a>

## Direct properties — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 3

<a id="canonical-0fe6d5790969008f02dc610e8f04e199112b55377000b9c13f07e6c0053b3f70"></a>

<a id="canonical-e426dc9479b0e32c29da53473f488316e8dd72abee4f32227393068878f34ec6"></a>

## name property — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 4

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

<a id="canonical-18373484e23ef8b59035441cbf6b4c351fa0aa3f96183e35ca74835a7fcff015"></a>

<a id="canonical-3463f1eac3c050d49917a6bbdd4b4387f6e521d5f81041fca933bc998cbb3b39"></a>

## namespace property — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 5

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

<a id="canonical-67ec232da7bf01bee0d25539a8576c4e28027576158170621611a4a418a735fb"></a>

<a id="canonical-cae6a09b858d3504d24adfb13bc7090c87325897fe8d9da3b7dc499fa12c74d4"></a>

## tenant property — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 6

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

<a id="canonical-da9ad882b071908d8b5d399b10b6bee4661850aa4bd9886e6c8e35720f7e870a"></a>

## Next pages — origin_pool.use_tls.use_server_verification.trusted_ca / 6ddc19937f25 / 7

- [origin_pool.use_tls.use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d6d3e059c3665f643b999fdf3812795ba2895e20bd5b1ae4ea731267fdd56184)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-79df3deef533facec1df2592d14b280dd2b8986e8cc122c85987c485e8800cfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-738933c57083ee98dd933602c511e07ebfdd69cb7cb37575802d8f5d1707b2f5"></a>

## origin_pool.use_tls.volterra_trusted_ca — origin_pool.use_tls.volterra_trusted_ca / e4926d02c5e1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.volterra_trusted_ca

<a id="canonical-fef58dad8abbc915c19db2a5a4f26faed542e951f9725c009479ec4cb4e19765"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-314cf2870c91702aeaf42bc3e4d0618ef11df41a7599bd2d3c0971ad4e0ddc3b"></a>

## Direct properties — origin_pool.use_tls.volterra_trusted_ca / e4926d02c5e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bccac4b03a91e9512d28276e394a609c9349949fb4aa887dbaec06da725f4e2c"></a>

## Next pages — origin_pool.use_tls.volterra_trusted_ca / e4926d02c5e1 / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7443060311e4e40f56ad58d95b6a4fd6522f01acc5850e9e20db976b812f7cc"></a>

## other_settings — other_settings / d590cd66cb01 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- other_settings

<a id="canonical-79ce0146d0a4ee9405e13098cfc94c2f2879e36cf661b86da4d2561a7e4345ac"></a>

Type: `"single"`. Computed.

Configuration parameter for other settings.

Upstream description:

Other Settings.

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

<a id="canonical-8b0f3c56d09e2d30b36b80fa5da253d163bd644a822f8298885b2c756f745650"></a>

## Direct properties — other_settings / d590cd66cb01 / 3

<a id="canonical-555ebe021f6bcf6f9ff3a4f4e46863bf693902f9b47abff18e2819c550d47a4a"></a>

<a id="canonical-01db72aee53e69f9ad1ca1e0724279f74fe8f0de7882c21f809911b5b50b0daf"></a>

## add_location property — other_settings / d590cd66cb01 / 4

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.

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

- [header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d): complete subsection reference.

- [logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24): complete subsection reference.

<a id="canonical-44338bd8e2ec293d9afebd39a6e0864314fdc8a0a391240c2bb98d1a97c127cf"></a>

## Next pages — other_settings / d590cd66cb01 / 5

- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb6f08e074fcaa238c0d4ac2fec641de124424261b7069cffe87aab7a51bc6c8"></a>

## other_settings.header_options — other_settings.header_options / 6100513deceb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- other_settings.header_options

<a id="canonical-9acffc6f64d7492211f46c60844d9f58b6155bf90290d641140c5160291b5141"></a>

Type: `"single"`. Computed.

Defines various OPTIONS related to request/response headers.

Upstream description:

This defines various OPTIONS related to request/response headers.

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

<a id="canonical-8412d100298e8ffeeed728c983c90ff070d59847cc3815a32c4b92e05b8cc4b3"></a>

## Direct properties — other_settings.header_options / 6100513deceb / 3

- [request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d): complete subsection reference.

<a id="canonical-6988a313a538d82b9da381a423ca57ce1ae2383db7b79c9f5b317ab3678de469"></a>

<a id="canonical-f4fc65f9cbb86c2aa7c9093c06467f4f5df5c85d1b60597ee277d075f4f7f3c9"></a>

## request_headers_to_remove property — other_settings.header_options / 6100513deceb / 4

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1): complete subsection reference.

<a id="canonical-3c90f8a2f680af443c6c11b5976636e41b771017dc3991c91560d7db24a31ee7"></a>

<a id="canonical-8cb1ae059de10a047e55132fa6a7827df774dbaf7cd8d7b52c10101fe7700b90"></a>

## response_headers_to_remove property — other_settings.header_options / 6100513deceb / 5

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-362cb17705d62895544737b00275c42ba0d8166ad52181d9baab6ec90a368fa1"></a>

## Next pages — other_settings.header_options / 6100513deceb / 6

- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fac6dd863eb579f16a7c395dfc491f6416dc7aac57ec0418f792061f74220ec9"></a>

## other_settings.header_options.request_headers_to_add — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- other_settings.header_options.request_headers_to_add

<a id="canonical-64dc5f5be5c97b164d28d265fa6f4b496044f700db10102246590fd5016f40c6"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-1529468053d2e005f6a92e7b4cc8e3d0c3e4541bfb308b23b359781dbb8e7c6c"></a>

## Direct properties — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 3

<a id="canonical-cfff3bcdbbd79e676f417c086751da2fc793dfd519c951c8332e979b32e28ca7"></a>

<a id="canonical-b8db1222a427bf483e55507609438120a74f3458880033215668747d99622e1c"></a>

## append property — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-89c4c38d2549771db6a9ffd5497b60d0e95575c94fde97aa0e62c49514de7e85"></a>

<a id="canonical-80c25cc7f9bac4ddc0a614866996a7d0ed35bea7835b1714c7b9f89bf2b4d7a2"></a>

## name property — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b): complete subsection reference.

<a id="canonical-77de7794536f7a673b63108258a79435d5e58dd00b0b7c3233c84bffe06ce575"></a>

<a id="canonical-f7cfc512c704103182c68b2d7338b4ae4f4974ed7ef1946a058b13159da1870d"></a>

## value property — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-bb9a476f1bc307479979b7eefeb7cb416308fcec2c7ca4aadf9be761ebd46ea0"></a>

## Next pages — other_settings.header_options.request_headers_to_add / 719c63e6db13 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76c33846ae67be5d62c73428c76635238161c2a63da4b871072a122afb06f255"></a>

## other_settings.header_options.request_headers_to_add.secret_value — other_settings.header_options.request_headers_to_add.secret_value / 86aa23c2554b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d)
- other_settings.header_options.request_headers_to_add.secret_value

<a id="canonical-bdcb52247bd9a7f4953d78594047e5176d0451105a3a0a30528247a73abb60fc"></a>

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

<a id="canonical-9f4b8917c9927d36f4cc4ff7d24bb59fedc5882af77d0c9d95a38a147cd1651f"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value / 86aa23c2554b / 3

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-67939ae1901578a677e5439964641b28a3f9958c9aaec14b43519a3d8b551638): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9ba14e66444ec26ac6043c6a68c9b68a5bc66911a0985ad0aa1ad8ef3bd1b26d): complete subsection reference.

<a id="canonical-db40b67f8e18cba4622d906dfc550e4c24bb4b684bb8c093054e07117bc6ce57"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value / 86aa23c2554b / 4

- [other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-67939ae1901578a677e5439964641b28a3f9958c9aaec14b43519a3d8b551638)
- [other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9ba14e66444ec26ac6043c6a68c9b68a5bc66911a0985ad0aa1ad8ef3bd1b26d)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-67939ae1901578a677e5439964641b28a3f9958c9aaec14b43519a3d8b551638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee6960b824f2813e998ccebd3882889512b27ab124812c89658c794753a6ff38"></a>

## other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d)
- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b)
- other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-052aacea3e4a70c51b5c9dda25bd19a97dd8f3928ca8c05321af5ea8fe91b73e"></a>

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

<a id="canonical-378def03437fcf638a74806ae596c9af71417a6a186df8f901d2b2f9edbf5dec"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 3

<a id="canonical-ac34151c54d352cb6d43332225be5945efc77d08f21a06e55f5927de164fed30"></a>

<a id="canonical-455cdeea6530e752ed3b1889cbc23e6e938d169d2026e94c2426d71c12563431"></a>

## decryption_provider property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 4

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

<a id="canonical-9664782d75802b7a5392ee93a928eb673f91a89395f0a83d56cf0c4b595f49ef"></a>

<a id="canonical-4a6b68c0a132ea878ca1411b6ae132df636937ca1926ad71c9e19a5fcfdea0d3"></a>

## location property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 5

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

<a id="canonical-a90905012c044bff73bbd9e49f1cbef20f944c03e01e9d29edaa59425a57cb3a"></a>

<a id="canonical-de8cab6bfbfd452778d5cd4308ab01a2d0d9ec90512c189f964d2b188ba2d48b"></a>

## store_provider property — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 6

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

<a id="canonical-2f49669d3a62806ca590e52a7713b2962234c44722b8eab4156e7d5c57df4a88"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value.blindfold_secr / 5289832a57f7 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9ba14e66444ec26ac6043c6a68c9b68a5bc66911a0985ad0aa1ad8ef3bd1b26d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f4e78fa0fa7b3c3a1df1d2b985149bb1c8359a953658cdba175f7182ab369d2"></a>

## other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 1a22b83b69f6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c19d4c23ef49bd4c8598fd27eee5f59e92aca24a897ac106d3bfc905d5d14d7d)
- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b)
- other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-cce10ef6f1741c6562ec3ba81e9d35a2e4aea763126d4f80930598a4291eb996"></a>

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

<a id="canonical-e7944e9996480c89f5484b60be45d120f17c83f29a54e881345dcbdb6a578ae6"></a>

## Direct properties — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 1a22b83b69f6 / 3

<a id="canonical-bb209212788ba7f38a45c53325960dedd17a482a4c2b171591aacbe493261ab3"></a>

<a id="canonical-21e7a05c7fa6a470ae600e518887a495d2d4e15feebcf5b0e84e94fb4ac984c2"></a>

## provider_ref property — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 1a22b83b69f6 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b79776a3b8e3aa8d49a2c4b87a2205ad58f771a02c725adb5688421080783c86"></a>

<a id="canonical-c32020081dd61300330ebcf302f6deb8b7560c1eddaa1fa62e9c836c4c019054"></a>

## url property — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 1a22b83b69f6 / 5

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

<a id="canonical-3d4e424b47a90bad2f11b932eb65618c92ff2342eee6eb113450de849613ae75"></a>

## Next pages — other_settings.header_options.request_headers_to_add.secret_value.clear_secret_i / 1a22b83b69f6 / 6

- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3c7d5c0b6e8d1201ef9facdad1ef0fa7ea1c8500c1b458a4a8287dc6463a9b4b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-411843a088cd640a5c9ef4272b896a78e271fb93588af4f199088fbb4e47f2d4"></a>

## other_settings.header_options.response_headers_to_add — other_settings.header_options.response_headers_to_add / a133aabe1169 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- other_settings.header_options.response_headers_to_add

<a id="canonical-76a2e00467f94770b8f429cb18289623c5e6ae279235d0de489a0cb2fd1bfbdb"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-9ce406bc70cb0a01afe60cf56473abcc5b184ef2860666b4de15af3ece43ade5"></a>

## Direct properties — other_settings.header_options.response_headers_to_add / a133aabe1169 / 3

<a id="canonical-87dc3b714516847f99ac35436f0dbe8e74bcb7852d004746530ee28a402a9e66"></a>

<a id="canonical-38ef208eb5e49ae3c6572830bb592d2b2e0cad2ab80c63972d5b5635d3ca0b93"></a>

## append property — other_settings.header_options.response_headers_to_add / a133aabe1169 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-1b549b0f4c323787c1a3e0732146c601a1e76d7d5b0649826cce14296bb6e55b"></a>

<a id="canonical-85891d43efb639f6541d4af30f46bb58cf955345e1728293fcf2b501371c9bcf"></a>

## name property — other_settings.header_options.response_headers_to_add / a133aabe1169 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6): complete subsection reference.

<a id="canonical-2aa641cea008e7807f355ebf7803b7100d8ad580494e250f20e4cf25f5e3a695"></a>

<a id="canonical-513008f61a32f7e53636d6c12a592798da31399060c9c3976bd53c9e29e54295"></a>

## value property — other_settings.header_options.response_headers_to_add / a133aabe1169 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-be38eb69f95b0b7c011c72917b1a1b621074482c21fdf90cd53c524203493a3a"></a>

## Next pages — other_settings.header_options.response_headers_to_add / a133aabe1169 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2741b55ca0dfd405d6bd18f99b9b4c13916cc35ca4a57735bd9c9e3a07fe1c70"></a>

## other_settings.header_options.response_headers_to_add.secret_value — other_settings.header_options.response_headers_to_add.secret_value / 2ce4c45a2836 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1)
- other_settings.header_options.response_headers_to_add.secret_value

<a id="canonical-fb27cc972d0a25850f645141863728b5da09a9a58b4c93184aac375e78114156"></a>

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

<a id="canonical-9132a56fb91cd7d7324fc1a8fe4b19c125b7a86c2854c6fa34ed6d446e797a31"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value / 2ce4c45a2836 / 3

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-b243298d3afe1d526c691bd7a556ff4d06cf85ca5f6ed775d2b6d0fafd8b968f): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9ba9df42e9b6a9fe8ade95c5c62a5293dd1f5927a49503f393a445afe80a6931): complete subsection reference.

<a id="canonical-15267d0b8f6aeafb7275da3d9aafc289e8084f19fa536f53fe2a0476c612d127"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value / 2ce4c45a2836 / 4

- [other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-b243298d3afe1d526c691bd7a556ff4d06cf85ca5f6ed775d2b6d0fafd8b968f)
- [other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9ba9df42e9b6a9fe8ade95c5c62a5293dd1f5927a49503f393a445afe80a6931)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b243298d3afe1d526c691bd7a556ff4d06cf85ca5f6ed775d2b6d0fafd8b968f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9530f2a41b3bfb1d231e24d025b5f83de0e551b39179d4edecf0775f5b970667"></a>

## other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1)
- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6)
- other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-32db4dce901edb6d90c1e8e581d90e49c40f8f0673fe1f1486057be0c15425d8"></a>

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

<a id="canonical-1a04e27e133f0a68aade5e58a914ba00497b00b8ec3f979523ec6721d360a71a"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 3

<a id="canonical-4254fb62eabb046c148551e48f34a8024635e17336dbf86ae1d8a1fa813b7cc1"></a>

<a id="canonical-e884d7760e10ccd9fe3fe8a21e21352e369e93049d7e6f622331b84861bf0fbd"></a>

## decryption_provider property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 4

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

<a id="canonical-133b71128cb5f245c1fef5906edaa99c201666862c5d048f51ef1ca351a23f23"></a>

<a id="canonical-8fe68be7d13ef9023e115eead6aa32f5bf79a1aee03ad0dfc613daf9609b8abc"></a>

## location property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 5

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

<a id="canonical-ec465c4774dc9fa23cf47cd59776efbd7d4a38ec2e16ce14b9e347355283ac68"></a>

<a id="canonical-5a6731a56ef2b7b1b671fe6fbdfbd79853929c323d1d9156929854430ae0e1be"></a>

## store_provider property — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 6

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

<a id="canonical-3446de9b8e84b25a0e25ae9ab1346d73513c7016e844e9518c7ef6b47ef07a5c"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value.blindfold_sec / 057221e04b01 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9ba9df42e9b6a9fe8ade95c5c62a5293dd1f5927a49503f393a445afe80a6931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4e312d6bb6f96985690234534abc62c94057a3d9ec785d93ac110ff243268b5"></a>

## other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 75af374d93f0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-554f10e0dc7a3e84ca2f7e4aeca40f08c319ff598d11892d388c475b14cf296d)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-591a9db3e0bb4e1a0b4d1deeff8914267bf67ec1d1745b2b1b27b811febbbfa1)
- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1770414ccaef537d641ccc5961b76c9a7a32e3ddfe97a97e028265c0d469b378"></a>

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

<a id="canonical-edb73379dbabf3ab7a4fd37353a40271c5acd2874fcf2fe48403fefffaa811b9"></a>

## Direct properties — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 75af374d93f0 / 3

<a id="canonical-f21206fcc3879f4950a43b73a84cf9329f24b6c2bc88b41f12e3019dbfc765ee"></a>

<a id="canonical-b509a193b2fb99a373f3468dec068a85e6d3b88222c3c622a2e15aad8455bf3c"></a>

## provider_ref property — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 75af374d93f0 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d5e4de992e1f3ab7693540ceefcb657c1db658d0166028d7e9f258ccbe64311a"></a>

<a id="canonical-58f71a6c3e202a22265f6f3ba590bc7f15d8bf53a629170231f6020ba476839b"></a>

## url property — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 75af374d93f0 / 5

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

<a id="canonical-510be151742e124ce44cd45d33c13f37a9c0fc36b464c7799723fd707b3e23fb"></a>

## Next pages — other_settings.header_options.response_headers_to_add.secret_value.clear_secret_ / 75af374d93f0 / 6

- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d730b17a92fe8291492a053f6723092ead2a9e3bb280e63ed0b6bb56a91242b6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6f33ae2fb46e18e311fc9ab6007b806119d2abe7972520cafc60df42390cedd"></a>

## other_settings.logging_options — other_settings.logging_options / 703151f69bbb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- other_settings.logging_options

<a id="canonical-603a526776f9af25c06cb779ea2aa0eb47683ace85bb71388833d8ae08136be5"></a>

Type: `"single"`. Computed.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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

<a id="canonical-0d1ace18ababffdd0f53e46c5fc128531ddc8964f57dc53727110b4d88e68590"></a>

## Direct properties — other_settings.logging_options / 703151f69bbb / 3

- [client_log_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ce469d84fce86ed9fb3402428e7e6fb6c3909769f6b36b0012edc83f7893ee4a): complete subsection reference.

- [origin_log_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-742a3e885798056f43f818ecd180934e78897d8562a8e95da0f910bf1d4cbdb0): complete subsection reference.

<a id="canonical-d822b58fa8c8519abe5371c9dfe6e6d1e1a9f3f6b4f87fbc501b1337aa203d2e"></a>

## Next pages — other_settings.logging_options / 703151f69bbb / 4

- [other_settings.logging_options.client_log_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ce469d84fce86ed9fb3402428e7e6fb6c3909769f6b36b0012edc83f7893ee4a)
- [other_settings.logging_options.origin_log_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-742a3e885798056f43f818ecd180934e78897d8562a8e95da0f910bf1d4cbdb0)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ce469d84fce86ed9fb3402428e7e6fb6c3909769f6b36b0012edc83f7893ee4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6200118a3fa4eaff5a07c8fbaefe37345f5375a0601d7a306e9a89edbb81197"></a>

## other_settings.logging_options.client_log_options — other_settings.logging_options.client_log_options / 21ef5b4da8cc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24)
- other_settings.logging_options.client_log_options

<a id="canonical-13bb753a75360d773855332c9b1165c32fbb00a585409e0413659452f56b3d26"></a>

Type: `"single"`. Computed.

Headers to Log. List of headers to Log.

Upstream description:

List of headers to Log.

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

<a id="canonical-6e6cfbd6a6afc0d35a5b87b973d1120934e26e2c46cca671906f8bc4292f8b33"></a>

## Direct properties — other_settings.logging_options.client_log_options / 21ef5b4da8cc / 3

<a id="canonical-1dc185416e583555535a8386132567f1373efbaae16a062a0756bf748ad9e201"></a>

<a id="canonical-ada4dadb278720c4068add52a3649061eb9ccbc35363e4c558e76ad1a4370c65"></a>

## header_list property — other_settings.logging_options.client_log_options / 21ef5b4da8cc / 4

Type: `["list", "string"]`. Computed.

Headers. List of headers.

Upstream description:

List of headers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-be43eb8de7e621a80d7a1ea6f6283c0710480feaa904e86b35745843a71d8dc7"></a>

## Next pages — other_settings.logging_options.client_log_options / 21ef5b4da8cc / 5

- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-742a3e885798056f43f818ecd180934e78897d8562a8e95da0f910bf1d4cbdb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-976b7d9475d491845bf99fa719906edfb7906ee7b5568fd67a461bfd8a4ce5a3"></a>

## other_settings.logging_options.origin_log_options — other_settings.logging_options.origin_log_options / e5fae1521862 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24)
- other_settings.logging_options.origin_log_options

<a id="canonical-3311c5f3dd2f8865ec29e1cf1852358a30182fb59fa850b27f6446616198012a"></a>

Type: `"single"`. Computed.

Configuration parameter for origin log options.

Upstream description:

List of headers to Log.

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

<a id="canonical-bd8cbab0cd1a96800d4194a7fcd900b915d96488ccba945b033b1b5a9a55e766"></a>

## Direct properties — other_settings.logging_options.origin_log_options / e5fae1521862 / 3

<a id="canonical-127ce4ba149012d4686ba1cecbd285ab43789a09f0acecd6ebad9ede6a389290"></a>

<a id="canonical-85e5464dec775256d6eb9fd91c9c1e5c1e0e228a965e70100928e577bce2adb4"></a>

## header_list property — other_settings.logging_options.origin_log_options / e5fae1521862 / 4

Type: `["list", "string"]`. Computed.

Headers. List of headers.

Upstream description:

List of headers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c3103686263f4fdf08721d422174b85ed60112fb050d8b11b0013ef3e240f010"></a>

## Next pages — other_settings.logging_options.origin_log_options / e5fae1521862 / 5

- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ab4d149f7985ab013b41b60a849a7eed26089779bcf1ac7d21bea6dc05abaa24)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77ef377c2d5ececa7e738057f7e34f8494ba844c842a1021fba8f8176ce02c24"></a>

## policy_based_challenge — policy_based_challenge / dc9f825d120d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- policy_based_challenge

<a id="canonical-f8ea95ea7129b9553efc10dceba7dc6c0e8b0c27c69d7bc9c50528e5f64df06d"></a>

Type: `"single"`. Computed.

Specifies the settings for policy rule based challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

<a id="canonical-92595c328bab611f638a28f56577ea909bbb4966ea116841c977b910e0504399"></a>

## Direct properties — policy_based_challenge / dc9f825d120d / 3

- [always_enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-8c4015fc967205a2305f4037053e0f02818d855624562d18461764415e837ecf): complete subsection reference.

- [always_enable_js_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-22b427bde5af1757e3b8d0651ba9ef31c08964a2802d25059f93f8fa1c2b108d): complete subsection reference.

- [captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-6f2a0fbe3cdfe25d8793cf555d68b9e38690ffc2c87d62bcd0506bc2bf92f0ba): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3de583d7d1a460b6cae22bb8b415e88595b986fac4322e66b25f80006335eb78): complete subsection reference.

- [default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c833de02bbd9362a2fb3f11109842a365061b7f66bffc6ee3c919899b1c2a6f0): complete subsection reference.

- [default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-07d13c1d38493df1d4869c832147e1f114a7e89a20ac478214b68c812664f492): complete subsection reference.

- [default_temporary_blocking_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-80f831137eba2a92cb74a6c2d6501d72f10e21912f2dd9ed99eb13e1b722ae43): complete subsection reference.

- [js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-f6f8ff3492217a57d1439bcb56b30054a16ac6a2d2551fdc7ea2fbbe830bc0ca): complete subsection reference.

- [malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-7c0ac3f151925f0fea6d6d3e98218a147da77c8e2ff30337aab83b8db2087619): complete subsection reference.

- [no_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-b90cbb0501d1fb3679d5aadcf5c2e7146dc5e93b124d41c3debb5582abbe085a): complete subsection reference.

- [rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e): complete subsection reference.

- [temporary_user_blocking](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-43ab134c64429c6ced8f53097b7e5fc4a7396dc81a64fe89365c55c91247d744): complete subsection reference.

<a id="canonical-1546a701767ecdb2ca065bed9211ff72f495c013eadb829630f9dcda06eaa91c"></a>

## Next pages — policy_based_challenge / dc9f825d120d / 4

- [policy_based_challenge.always_enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-8c4015fc967205a2305f4037053e0f02818d855624562d18461764415e837ecf)
- [policy_based_challenge.always_enable_js_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-22b427bde5af1757e3b8d0651ba9ef31c08964a2802d25059f93f8fa1c2b108d)
- [policy_based_challenge.captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-6f2a0fbe3cdfe25d8793cf555d68b9e38690ffc2c87d62bcd0506bc2bf92f0ba)
- [policy_based_challenge.default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3de583d7d1a460b6cae22bb8b415e88595b986fac4322e66b25f80006335eb78)
- [policy_based_challenge.default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c833de02bbd9362a2fb3f11109842a365061b7f66bffc6ee3c919899b1c2a6f0)
- [policy_based_challenge.default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-07d13c1d38493df1d4869c832147e1f114a7e89a20ac478214b68c812664f492)
- [policy_based_challenge.default_temporary_blocking_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-80f831137eba2a92cb74a6c2d6501d72f10e21912f2dd9ed99eb13e1b722ae43)
- [policy_based_challenge.js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-f6f8ff3492217a57d1439bcb56b30054a16ac6a2d2551fdc7ea2fbbe830bc0ca)
- [policy_based_challenge.malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-7c0ac3f151925f0fea6d6d3e98218a147da77c8e2ff30337aab83b8db2087619)
- [policy_based_challenge.no_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-b90cbb0501d1fb3679d5aadcf5c2e7146dc5e93b124d41c3debb5582abbe085a)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.temporary_user_blocking](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-43ab134c64429c6ced8f53097b7e5fc4a7396dc81a64fe89365c55c91247d744)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8c4015fc967205a2305f4037053e0f02818d855624562d18461764415e837ecf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfb3b814a7eaf09c41091445999f87d720dbdc89ec70c885470ede488e951fdc"></a>

## policy_based_challenge.always_enable_captcha_challenge — policy_based_challenge.always_enable_captcha_challenge / 989364b9eb9e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-7470bf5bcff5eb228ff881b0a03d4a9c1324ead1c798194f6e876dc203c4e642"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable captcha challenge.

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

<a id="canonical-c7d89c10557c8ae7e6a41390f4b3c76ef4a664d7d07d8f8d4e41d5d90fcb8799"></a>

## Direct properties — policy_based_challenge.always_enable_captcha_challenge / 989364b9eb9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3cfe384d858c0f5db8863d30ef9ee37949532eab7f8cb24ade90baa1c371f834"></a>

## Next pages — policy_based_challenge.always_enable_captcha_challenge / 989364b9eb9e / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-22b427bde5af1757e3b8d0651ba9ef31c08964a2802d25059f93f8fa1c2b108d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95f7a5f13b5a2c28ded20f484ecc46245d527b9041c2b8732e00dac7ea7dfe8d"></a>

## policy_based_challenge.always_enable_js_challenge — policy_based_challenge.always_enable_js_challenge / 0c845783b893 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-338cadd5f9109053d92cc385312578a8860f7dcd235976c4752f629eef18734c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable js challenge.

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

<a id="canonical-bf3419757e968881b3a232526e0226a472b65de97f740d9b75648ae100272af6"></a>

## Direct properties — policy_based_challenge.always_enable_js_challenge / 0c845783b893 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5789af3c36cc79e1eb4eb7e504940d0aee79f1b2312967b024154b845323785a"></a>

## Next pages — policy_based_challenge.always_enable_js_challenge / 0c845783b893 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6f2a0fbe3cdfe25d8793cf555d68b9e38690ffc2c87d62bcd0506bc2bf92f0ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-608a501685f2298c8c146e8a050456928db65c9e1d623cd594520be0a348ecb6"></a>

## policy_based_challenge.captcha_challenge_parameters — policy_based_challenge.captcha_challenge_parameters / a9285f273b2a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-624be37b7ce39fc10d60c93af1c6a9fae5bf2f5770e90d387182f7c615c476e5"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2c62af91a0371659f20960762237bfd557918fcf40373e20c48b23b7db8cdaf9"></a>

## Direct properties — policy_based_challenge.captcha_challenge_parameters / a9285f273b2a / 3

<a id="canonical-f8d6f2f2fdb24dce2749e2f2837809cec1fcd4558878ab30159e9a23002cad8c"></a>

<a id="canonical-821f324694baef724763df31665a6aea07108ea1e5d9f282ba1ba3c993c1e125"></a>

## cookie_expiry property — policy_based_challenge.captcha_challenge_parameters / a9285f273b2a / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-26265d90502f59d7697492a3f3de05fb6e055eb2d1d44015737de65b81c0b416"></a>

<a id="canonical-381d125446e58351779f48e4a7ac159e2c9fec3d914b9dc5cdad378cb708825b"></a>

## custom_page property — policy_based_challenge.captcha_challenge_parameters / a9285f273b2a / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-1fe559a99ebe0346812ea00e1ba93ec9355d50d31e347a9102a213ef1d5113ff"></a>

## Next pages — policy_based_challenge.captcha_challenge_parameters / a9285f273b2a / 6

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3de583d7d1a460b6cae22bb8b415e88595b986fac4322e66b25f80006335eb78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d924930375016a18c23a7a2e1fa266d0c92e3d24871ba4871b8f8f156fccb653"></a>

## policy_based_challenge.default_captcha_challenge_parameters — policy_based_challenge.default_captcha_challenge_parameters / 539f8287e685 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-f10e028e8e012ca891583d4ff36c13670cf270f4a47a79f13a66d77d0f18f61e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-60b4494330caec49ab4f5139fb3d42696ddde1f23d8658f55f438d3152743abc"></a>

## Direct properties — policy_based_challenge.default_captcha_challenge_parameters / 539f8287e685 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d455a07dcd48955c735b255ebe26a65d1740c85692d74eaf705dabab85767895"></a>

## Next pages — policy_based_challenge.default_captcha_challenge_parameters / 539f8287e685 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c833de02bbd9362a2fb3f11109842a365061b7f66bffc6ee3c919899b1c2a6f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649b07a5d4f2ee9b8aeb094e4e55607b72e47c49dd7ecdd61a4d4446b29cd629"></a>

## policy_based_challenge.default_js_challenge_parameters — policy_based_challenge.default_js_challenge_parameters / 8b69be8915c0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-f1d28bfe096f6dc09487110b25e58e5c9a8f1ec006fc0c6c0101e444d4ec4144"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-912c5722928f6ca2d999400e51167785bf4b3a96c6c9729f8f80bce35fda492b"></a>

## Direct properties — policy_based_challenge.default_js_challenge_parameters / 8b69be8915c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82eb7aff07e359746525b10c9c17bf4f4ef2de349230f0fefa09c0800d360792"></a>

## Next pages — policy_based_challenge.default_js_challenge_parameters / 8b69be8915c0 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-07d13c1d38493df1d4869c832147e1f114a7e89a20ac478214b68c812664f492"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-065d93acd1d17c86558c146ff0da5930088842237fb9c85ea0004baaf47c0167"></a>

## policy_based_challenge.default_mitigation_settings — policy_based_challenge.default_mitigation_settings / 6658fb0a63a4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-31e39a6f9f29b189217f4a600127409ea5091513733b79e314fc30e86420a0c0"></a>

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

<a id="canonical-cabe8657fae424037cd567d8a0bd49e0f9b28c76d1fd309ed75ec043aad3dcfa"></a>

## Direct properties — policy_based_challenge.default_mitigation_settings / 6658fb0a63a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d52838b94a1d63395c6100dd4731d499ded1b0518ca787d0ec7e51ea2881c3c7"></a>

## Next pages — policy_based_challenge.default_mitigation_settings / 6658fb0a63a4 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-80f831137eba2a92cb74a6c2d6501d72f10e21912f2dd9ed99eb13e1b722ae43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0caeeefb6fb11702510a75004cceb527fcda75f679429c4859492ef02740b849"></a>

## policy_based_challenge.default_temporary_blocking_parameters — policy_based_challenge.default_temporary_blocking_parameters / f3e0ba549da9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-8e1341fb7b204f17f14386604b08e2fe5cd3bd51367a5d937fed150374282e68"></a>

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

<a id="canonical-291429272009505532fd14f92b146e6881449839600191ffb63bde6b10e46292"></a>

## Direct properties — policy_based_challenge.default_temporary_blocking_parameters / f3e0ba549da9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a19af5cade671866dfb39c5ed9a4f7745c079fba394fee2f9ed9a3605fc0d74c"></a>

## Next pages — policy_based_challenge.default_temporary_blocking_parameters / f3e0ba549da9 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f6f8ff3492217a57d1439bcb56b30054a16ac6a2d2551fdc7ea2fbbe830bc0ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afe92f1102498be579002bb7256cd116fc32bf78b994312768d6562f733194e3"></a>

## policy_based_challenge.js_challenge_parameters — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-d924ea071c7982a9b3cca53da35a701d90cf9ce5bbd8063cf24127b5f7618dcf"></a>

Type: `"single"`. Computed.

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

<a id="canonical-e7f554a8af346a950c19cf9c555308d307285ae8fbe2debfc79c5d4bde235749"></a>

## Direct properties — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 3

<a id="canonical-98071be61a3079a79135b77f7701700fefceb8edfac5bcc9d176541da81315d6"></a>

<a id="canonical-699a126058fbe316c0f5b35458e390d6b89113f214e21e4966d444781b964017"></a>

## cookie_expiry property — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-5b13f71411c1fedc4e845156a08e64c43f3888fe06c2317c4029c7e15a1a393b"></a>

<a id="canonical-cd5f6fa3d2efffa09d4cfd1ec190794a324701ebc194a43f223b806ffb980a0f"></a>

## custom_page property — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-7d1b471af2a6ae749a8b5055aa14b22520bcd4f8ca09c7676508c77cb245d68e"></a>

<a id="canonical-1163bc5399bb983257d9bbd02576a3d476f08b43a3986d35eb4c763be67c5ba4"></a>

## js_script_delay property — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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

<a id="canonical-9e1115abc34188bcde697eb0d7c14188ffc288b62217251507c6f9e30e265c9c"></a>

## Next pages — policy_based_challenge.js_challenge_parameters / c1f96ea95763 / 7

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7c0ac3f151925f0fea6d6d3e98218a147da77c8e2ff30337aab83b8db2087619"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a3005dbadeb5dd48aab3d25236cdeb3a35ae26e9db977e36c04a401edb04cdc"></a>

## policy_based_challenge.malicious_user_mitigation — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-cb7da78954aa8849498276616921f8cda961362d221b2321a84508266a1f37ae"></a>

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

<a id="canonical-a6050d0404c3874b13f08ade6ec71cff29d4411f3266d0c2fc931e2fbfc05df2"></a>

## Direct properties — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 3

<a id="canonical-3fb8b8b30bac24867f66010daf83aaa03511e1be4b10af7c1747cb3b81131aaf"></a>

<a id="canonical-b81870b3c65e8a8188df3f085d0b4daef4306b69ea4243812428ea1af4e44ff9"></a>

## name property — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 4

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

<a id="canonical-42c63050b4f91ce3a126c940adbdb9b0a295cbecaf42de52fea5c168c428d343"></a>

<a id="canonical-5998c357bd209a38a8e24e3cb31b22018b84bb6fed719a11c6a39206af6c6695"></a>

## namespace property — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 5

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

<a id="canonical-04d244dfa09ef173aab013c7396bbd2ca2415f03b9264efd6bda068764f8a905"></a>

<a id="canonical-492539501bcb6e4f87b2908a25ed5ff94f72150b617c0c5fadccbf03b467d2b2"></a>

## tenant property — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 6

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

<a id="canonical-139bc3297ae7ef6a10d2656a624e52181ac83e3a836b52b0d48988c89ce89a6f"></a>

## Next pages — policy_based_challenge.malicious_user_mitigation / 6803badbe58a / 7

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b90cbb0501d1fb3679d5aadcf5c2e7146dc5e93b124d41c3debb5582abbe085a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c4fc101ee01a65a853c580cee2817db045d9789b4d72fd049a2bb6a090da99e"></a>

## policy_based_challenge.no_challenge — policy_based_challenge.no_challenge / e4255c5cb788 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.no_challenge

<a id="canonical-ba242615b050f0b2a7af26eecc8b91e00291cc2dbc8827282474b2d63a209b87"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-efdf7b74c9040f22a15a7c82c7a5466fe36828c1ffd5b5fd318d68541f0507c6"></a>

## Direct properties — policy_based_challenge.no_challenge / e4255c5cb788 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c64ce4879269e2d438ec045a95d8948016b3602c98605d450347f98ac5a15f5"></a>

## Next pages — policy_based_challenge.no_challenge / e4255c5cb788 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a598857082d2baf80cda0fe3f2dca3a8317d2044831845f0db23e382f89b02b"></a>

## policy_based_challenge.rule_list — policy_based_challenge.rule_list / 72fec7f283d1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- policy_based_challenge.rule_list

<a id="canonical-eb2efc97bde70d45cc263388b02b1d2f39c58d0b6d7840f77f81dd7b92a9e1b4"></a>

Type: `"single"`. Computed.

List of challenge rules to be used in policy based challenge.

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

<a id="canonical-8a7aad1ad379b5e85137f988673719414de829beb173e2111038216d4a1ce70a"></a>

## Direct properties — policy_based_challenge.rule_list / 72fec7f283d1 / 3

- [rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732): complete subsection reference.

<a id="canonical-86393efa17bb96034134b0b51851d3e873ead490a13e60d2cd0c1fe1a08a3d1d"></a>

## Next pages — policy_based_challenge.rule_list / 72fec7f283d1 / 4

- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cebfeae3d6db8230347903ac69e4ef7c1d2d768b82aa772fd89bdd828ee1b2f"></a>

## policy_based_challenge.rule_list.rules — policy_based_challenge.rule_list.rules / 7ea216be1935 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- policy_based_challenge.rule_list.rules

<a id="canonical-8e8498e096b8c6d639af34c16bf9fe9815f56a47a731b69eee70de9c9b92aba2"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-d83e2f7efc7963564761438fb64eb298ac3184ed73446b6fade815517a8c8038"></a>

## Direct properties — policy_based_challenge.rule_list.rules / 7ea216be1935 / 3

- [metadata](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-a7fae19800b991ad62e434008b3d99fe76406881e0cde4f9e374589e0bb82a51): complete subsection reference.

- [spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec): complete subsection reference.

<a id="canonical-a3f263d097e4c1acc08d42367c48274c7e510860d5c40dbbf364594ed06c0943"></a>

## Next pages — policy_based_challenge.rule_list.rules / 7ea216be1935 / 4

- [policy_based_challenge.rule_list.rules.metadata](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-a7fae19800b991ad62e434008b3d99fe76406881e0cde4f9e374589e0bb82a51)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a7fae19800b991ad62e434008b3d99fe76406881e0cde4f9e374589e0bb82a51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51db7ff593e78f21ab4a14d16decf18c58e42c97a478fa705368a52a16aa03fe"></a>

## policy_based_challenge.rule_list.rules.metadata — policy_based_challenge.rule_list.rules.metadata / af0853de1660 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-4b215c5f4061739f0b07ce0bc071e6e7e85b6389201a61eac4841a5cffe81990"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-eb42f0adf843d764e2b90144c572785b7fc1af6cf49d5cda63460e8f982ef6a9"></a>

## Direct properties — policy_based_challenge.rule_list.rules.metadata / af0853de1660 / 3

<a id="canonical-f14cd58fad4156669adcf75a56ae8145ac5f718a9bc192b0353566507aa2f783"></a>

<a id="canonical-593e2520df586dce5286b1679b4368c0ba96bac6dc28cbbcb3a2df7053b4e535"></a>

## description_spec property — policy_based_challenge.rule_list.rules.metadata / af0853de1660 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-329109718102b4aeb8c1f96fd6c0bcb4c3a9cdad7d6cd5e47b10e322279c7c83"></a>

<a id="canonical-38821e2836ed4b22c32592fb918432a53b0278616d0408624e36e1dad36501b3"></a>

## name property — policy_based_challenge.rule_list.rules.metadata / af0853de1660 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-ac8fd85e840a68188c69a6f7d0a2da5bb3f59446138fe2d03e4824fa3139d674"></a>

## Next pages — policy_based_challenge.rule_list.rules.metadata / af0853de1660 / 6

- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b1dfa6653202aeba42e6db75faa9abbfac8b9227cba78900d444cf7494a772d"></a>

## policy_based_challenge.rule_list.rules.spec — policy_based_challenge.rule_list.rules.spec / 4d01751646e7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-4a9d33ab252b47151ab13de302e7e33fc0f82130a9c5f5fb914ebf0e5c774a98"></a>

Type: `"single"`. Computed.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-d2c45ef753b2967b9a291f4c2801b98b0c8c623ae8a7f0b9a2a1404d5cb245b2"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec / 4d01751646e7 / 3

- [any_asn](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-659fbffddbb74b8b12f96510855bf5d4649c6c512f6e3d7f2603e1f425c5a16e): complete subsection reference.

- [any_client](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c13a4676ae2784bda9736b373a3a1c4f3760fd08587e5361a130634b89e39595): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9e26d7169be73be2433136fde8a71b7138662bbad7c994d4da15bcffd7ea820a): complete subsection reference.

- [arg_matchers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3f14e22a4d8a6cfab286b746592122bbb117c66a7217fb70f6fc6fc65f5d4ffb): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-ccbf80da4527e029e324a1d29cf03e9a63cb90bbe5921577067aa34db412360a): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-dad26e2f0dc997224662f7063acf4f7ed43e95cab51888a99b0233b646d918a6): complete subsection reference.

- [body_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-67d498f2be45d84a5aa71a6ab083062c8a7b570910eef4ef407fc0855bf68695): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-d26150635677c7b89912fa46b84c8f8497ac9ea6b5fb06077500c96a9e9fb1a5): complete subsection reference.

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-fdd5e640b26f63c3a194b08e601ba1568ea2d5a9300d21d11ffbe728d8b52fc0): complete subsection reference.

- [disable_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-ff34603aca76f83e6a83eb8fa4609b7e4bbbadc21fad3665035296fb37f7d2ac): complete subsection reference.

- [domain_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-67df7fd347a3d3a7baea3c182eb8bab154fcf30598cafd9ad2142f4b921a74e8): complete subsection reference.

- [enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-c578c161a5d6e9f036cb1b99fcd7497e5bf899c26eb2870e6dccc4ac5edd987b): complete subsection reference.

- [enable_javascript_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-b43b8fcc43943096a8339879b69bc4ddd259957713c4a5a7be6ffb4a2f62a7cb): complete subsection reference.

<a id="canonical-fa71808332207d2ea1692dc9ca33be8b077a05f1449b9feb51a28244847651fb"></a>

<a id="canonical-c2f4802f7d8cae075976d7131717da80d172572aadd666b82b9e699a701dd916"></a>

## expiration_timestamp property — policy_based_challenge.rule_list.rules.spec / 4d01751646e7 / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-9fa0d84cf8cb93f57b225d650ce3d81ca93100d8544bc25c8796dbb9ca1ba89b): complete subsection reference.

- [http_method](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-23683c8435c4dd6d576c46f06d8125a6151d9bfc125e3e53bb7df77114e5783b): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2132535d55d177815e6ba86f041bfaa881bf5b62e47665e6b2e52beb22ed9de9): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-263ca72f702934a3e9be53940a83efa7d9a9fb871ab45230f543798d30b87d42): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-7b803221075bc73e2cecaeb9add4881f76e28d2872a59e26930a9b03321399da): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3a3e671f577ddea6131c12599336045bc59478c780aeda95e9670229861e4c0b): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-b58304aba01552cf73942701a9f2b0839f60e5014cf74a9273b7478835d054d2): complete subsection reference.

<a id="canonical-b728ddd7114d30dcca2381d710f9ab6543055a1258d24676ee5ba51f8c2dfb50"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec / 4d01751646e7 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-659fbffddbb74b8b12f96510855bf5d4649c6c512f6e3d7f2603e1f425c5a16e)
- [policy_based_challenge.rule_list.rules.spec.any_client](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c13a4676ae2784bda9736b373a3a1c4f3760fd08587e5361a130634b89e39595)
- [policy_based_challenge.rule_list.rules.spec.any_ip](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9e26d7169be73be2433136fde8a71b7138662bbad7c994d4da15bcffd7ea820a)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3f14e22a4d8a6cfab286b746592122bbb117c66a7217fb70f6fc6fc65f5d4ffb)
- [policy_based_challenge.rule_list.rules.spec.asn_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-ccbf80da4527e029e324a1d29cf03e9a63cb90bbe5921577067aa34db412360a)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-dad26e2f0dc997224662f7063acf4f7ed43e95cab51888a99b0233b646d918a6)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-67d498f2be45d84a5aa71a6ab083062c8a7b570910eef4ef407fc0855bf68695)
- [policy_based_challenge.rule_list.rules.spec.client_selector](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-d26150635677c7b89912fa46b84c8f8497ac9ea6b5fb06077500c96a9e9fb1a5)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-fdd5e640b26f63c3a194b08e601ba1568ea2d5a9300d21d11ffbe728d8b52fc0)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-ff34603aca76f83e6a83eb8fa4609b7e4bbbadc21fad3665035296fb37f7d2ac)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-67df7fd347a3d3a7baea3c182eb8bab154fcf30598cafd9ad2142f4b921a74e8)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-c578c161a5d6e9f036cb1b99fcd7497e5bf899c26eb2870e6dccc4ac5edd987b)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-b43b8fcc43943096a8339879b69bc4ddd259957713c4a5a7be6ffb4a2f62a7cb)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-9fa0d84cf8cb93f57b225d650ce3d81ca93100d8544bc25c8796dbb9ca1ba89b)
- [policy_based_challenge.rule_list.rules.spec.http_method](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-23683c8435c4dd6d576c46f06d8125a6151d9bfc125e3e53bb7df77114e5783b)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2132535d55d177815e6ba86f041bfaa881bf5b62e47665e6b2e52beb22ed9de9)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-263ca72f702934a3e9be53940a83efa7d9a9fb871ab45230f543798d30b87d42)
- [policy_based_challenge.rule_list.rules.spec.path](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-7b803221075bc73e2cecaeb9add4881f76e28d2872a59e26930a9b03321399da)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3a3e671f577ddea6131c12599336045bc59478c780aeda95e9670229861e4c0b)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-b58304aba01552cf73942701a9f2b0839f60e5014cf74a9273b7478835d054d2)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-659fbffddbb74b8b12f96510855bf5d4649c6c512f6e3d7f2603e1f425c5a16e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8d1a5a519d7a314d51bd4e16ee006d37f091c0d44f2b08ec9d494545a441337"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — policy_based_challenge.rule_list.rules.spec.any_asn / 003191f099dd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-8ae61d418a6c85f2ea76d48975482fe75b7648748289b5eab44e6baae0395697"></a>

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

<a id="canonical-fa3b17e9d53d49b6aafb8f1facb90f8e9fbf80ce11810f61253de0ee37e6be33"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_asn / 003191f099dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-028555437777248ff248f73b118016056e3c0e3f37400774622705adb72b1b53"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_asn / 003191f099dd / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c13a4676ae2784bda9736b373a3a1c4f3760fd08587e5361a130634b89e39595"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-565e01bcfe5e781d2a31c1736bf9df1abb520318066a7fb4e9270b254baed20a"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — policy_based_challenge.rule_list.rules.spec.any_client / 8d7d5b43f3b3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-97ec9b88f95a795e4f0420bb1a4623af7f173e12d75ad452fddf980990c88ba2"></a>

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

<a id="canonical-4c02e909ea16b3f1fd63c918f24ffb8be70460364822f951c1901a049c931629"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_client / 8d7d5b43f3b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-305035f7842247715f00afd6f2fa0a8ebf8453b6865193818313714368fb4aeb"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_client / 8d7d5b43f3b3 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9e26d7169be73be2433136fde8a71b7138662bbad7c994d4da15bcffd7ea820a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9d21e5e4462d7e9ce0042b00b3abd00ab74281b21f43b436358c8af49d1db09"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — policy_based_challenge.rule_list.rules.spec.any_ip / 2690e2405ab5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-11d050a685badcfdaf6a52faf3aa9d8496978483894f8d8f269880e66c9e9033"></a>

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

<a id="canonical-de04c3ab33c76c6d1024968162d6bb8c89951a06a0a3c6c364be57573f8f80bd"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_ip / 2690e2405ab5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed22024e2a35e43e1409db9787905d4e9670b3d426e46b3a8898f314b8270ddf"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_ip / 2690e2405ab5 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3f14e22a4d8a6cfab286b746592122bbb117c66a7217fb70f6fc6fc65f5d4ffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e95f9537608c9e6100f0d92d57682c40013feb360464b2de1de3ffc645520006"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — policy_based_challenge.rule_list.rules.spec.arg_matchers / 3034f15bd5aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-7a69a52ad4cc778773f0c498c77af5f47b727aa02ce23e2abfcc00b69b700fcd"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-233d61aac3ff39096faa8cd5f85711a89e945dd6da9cbc05ac4e80f04eec92c6"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers / 3034f15bd5aa / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-05920539df6ad9f16a4b5b7d65e02dad0162bc2e075899ac52d9b6127f2a2449): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-5efe823137fffc51b3f6f0c4630578742b1493010c664690d263064a451e5be1): complete subsection reference.

<a id="canonical-a25dd7853a13d66274b1e1cb940f8ee7fbc8fe6be8274dd8490dacc1a293a02c"></a>

<a id="canonical-68ddaf683173ec01c21ba27e9e38c3c8e11f3a16282fc3e26d092b607dd10a6b"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 3034f15bd5aa / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-aae0a7433fb1180c410fe1528fe964722d270f9acd6eba1a3d714b7fcd95afd0): complete subsection reference.

<a id="canonical-ed6308debcb321b2d5c763e0bedfa05856eae6ee577cb7e7b368e1b27ee01835"></a>

<a id="canonical-70da4127289affbfe6509ce29c2df33d9e66169d097630cd9f51ba1154120ac1"></a>

## name property — policy_based_challenge.rule_list.rules.spec.arg_matchers / 3034f15bd5aa / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-aaa09051edc1b2421637e0be9bded816c806edce6d2e217da77c036cb29e065c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers / 3034f15bd5aa / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-05920539df6ad9f16a4b5b7d65e02dad0162bc2e075899ac52d9b6127f2a2449)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-5efe823137fffc51b3f6f0c4630578742b1493010c664690d263064a451e5be1)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-aae0a7433fb1180c410fe1528fe964722d270f9acd6eba1a3d714b7fcd95afd0)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-05920539df6ad9f16a4b5b7d65e02dad0162bc2e075899ac52d9b6127f2a2449"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92a1b43557e5575a99a809bf54daa25431fcffb3545b658e2753e391d1d70a37"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / cec35b7e9d1e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24cb06baf50135710092a5bce0754a0628c17c7b9ea8bcf61c623d2c344dbc4e)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-9afa011786b3806b32f8c345bd840547e73c5eeabe0535556575a9ad800c2732)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-24730b9bf242b4c49663640dd4bbfa30201c0613314980baea0b50777c4302ec)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3f14e22a4d8a6cfab286b746592122bbb117c66a7217fb70f6fc6fc65f5d4ffb)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-1a4ba9e4329011fceecc8067eb760695b41885c92cb01643cce0f23c155f0fe2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-52dc559d36bd42638c7ff222bff4b987ee0cf2fea1d9a72c47e4d513bb00e7f9"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / cec35b7e9d1e / 3

This is an empty object or choice marker. It has no direct properties.
