---
page_title: "xcsh_log_receiver reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver reference."
---

# xcsh_log_receiver reference

<a id="canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98e7366ab2e5a3d3044a8c5a9f7292a2ca14def3beab7516bd5d34d4ce092f1d"></a>

## Property reference — Property reference / 3245db645141 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- Property reference

<a id="canonical-5a06b49818fc7f9607487b0cbbc667cf793d7f5b501c838982e2ea3744d2d81a"></a>

## Direct properties — Property reference / 3245db645141 / 3

<a id="canonical-02726dfdc1bac398f2a132470ec266efbe23ad762bc12f0bf2a3f55a3472847e"></a>

<a id="canonical-0e9a56e5ae25a31208130240c7ffe78600e6ebaac5632de8bba36d9eddfffdbe"></a>

## annotations property — Property reference / 3245db645141 / 4

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

<a id="canonical-e7cb1931caeb31da99c35d3b48b4ba9edff6df9784f1aa259b85763a2c18928c"></a>

<a id="canonical-56d70f90232640d2946749ec49f646c78ca7c7ec1b71991dfb19b408915a6648"></a>

## description property — Property reference / 3245db645141 / 5

Type: `"string"`. Computed.

Description of the LogReceiver.

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

<a id="canonical-2f5a2f4e17bced3961c16fcbcbbc11e9499e5f4a5952a54fb0501ff5cd8c858f"></a>

<a id="canonical-67897ea472edfee91de0e197eed9fd7323fa11e6ee9d3ea8389028d5b943c44b"></a>

## id property — Property reference / 3245db645141 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-18aa1aafb6b1406fc4dfb796d62492de69ebb5f79ab26e61af7d8589ee343c0e"></a>

<a id="canonical-401e31663c5b190155f267279089d797b98978da1658ff6ababb2bb03316c0a4"></a>

## labels property — Property reference / 3245db645141 / 7

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

<a id="canonical-6d48babc92a294dda4a1c18244a8d998e372b85ec70c9152b0d4643bce4502a7"></a>

<a id="canonical-fe1ef8f44f6aa9d0ab89afa193778eae25b35a41d0d6a79dc85dbb87a49912a5"></a>

## name property — Property reference / 3245db645141 / 8

Type: `"string"`. Required.

Name of the LogReceiver.

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

<a id="canonical-910669adb3a7431c7d73de34bf0e0282e30f486f8a23ac2861cf2635ba954f32"></a>

<a id="canonical-6dddc29e5018b1c19c37733bdf022b1804df44237394cad9b5e46ad9ba79625b"></a>

## namespace property — Property reference / 3245db645141 / 9

Type: `"string"`. Required.

Namespace where the LogReceiver exists.

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

- [site_local](data-sources--log_receiver--reference--group-001.md#canonical-41de8992f397c6ba3c8aaddb0fcbc4ee00cb085c1b8148cf5b89c8b33ed502e3): complete subsection reference.

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c): complete subsection reference.

<a id="canonical-7942c99d9910f7d260f1c4ef267371223303d169ab02de7d99e562ca79ed9694"></a>

## All schema paths — Property reference / 3245db645141 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--log_receiver--reference--group-001.md#canonical-02726dfdc1bac398f2a132470ec266efbe23ad762bc12f0bf2a3f55a3472847e) |
| `description` | [description](data-sources--log_receiver--reference--group-001.md#canonical-e7cb1931caeb31da99c35d3b48b4ba9edff6df9784f1aa259b85763a2c18928c) |
| `id` | [id](data-sources--log_receiver--reference--group-001.md#canonical-2f5a2f4e17bced3961c16fcbcbbc11e9499e5f4a5952a54fb0501ff5cd8c858f) |
| `labels` | [labels](data-sources--log_receiver--reference--group-001.md#canonical-18aa1aafb6b1406fc4dfb796d62492de69ebb5f79ab26e61af7d8589ee343c0e) |
| `name` | [name](data-sources--log_receiver--reference--group-001.md#canonical-6d48babc92a294dda4a1c18244a8d998e372b85ec70c9152b0d4643bce4502a7) |
| `namespace` | [namespace](data-sources--log_receiver--reference--group-001.md#canonical-910669adb3a7431c7d73de34bf0e0282e30f486f8a23ac2861cf2635ba954f32) |
| `site_local` | [site_local](data-sources--log_receiver--reference--group-001.md#canonical-236a21f0a33913fd3ea3159f961ac14e418ab85d90d2b99ce5943d1817500d12) |
| `syslog` | [syslog](data-sources--log_receiver--reference--group-001.md#canonical-bff263eba5710055aa83f1d3dcab925cde2f92e5ccaab7a3ee5a1e0a44390f37) |
| `syslog.syslog_rfc5424` | [syslog.syslog_rfc5424](data-sources--log_receiver--reference--group-001.md#canonical-126a944329d27421170ffb0c77783b366b0b4d31895061b1da120c5debfe7a75) |
| `syslog.tcp_server` | [syslog.tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-d1ef828c022f7bbe95ef721f6b4089df8007e4d77b106f75049c2022e4a34d5a) |
| `syslog.tcp_server.port` | [syslog.tcp_server.port](data-sources--log_receiver--reference--group-001.md#canonical-29cab7fec224f61289fd4f0934619653787cb0fbc3d4b8d6d5f58074a713c4a8) |
| `syslog.tcp_server.server_name` | [syslog.tcp_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-95eb972e2aa0a80e1ed0fbbfd22df573b7eb3fa706b47575ab718065c74cc2b7) |
| `syslog.tls_server` | [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-b9a7bd2a1a194179caa2c865db074c6896ef07d0fce1d2f25ba81826f83c4579) |
| `syslog.tls_server.default_https_port` | [syslog.tls_server.default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-78782507dcbf075d2059596bc58fc85cf9e7a1e394751e3b77df6734fd376c6c) |
| `syslog.tls_server.default_syslog_tls_port` | [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-4737c459ba38398436871b72144f2c3fa9fc9c15d4be3f2a5a6afb36f2a50b6c) |
| `syslog.tls_server.mtls_disabled` | [syslog.tls_server.mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-1bf305c83aabc3e79f183d7f98504d101f58ea6d5a1d8ed41924236886b46dfc) |
| `syslog.tls_server.mtls_enable` | [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-414958da61e7cbe7ae2248370675c7c3d66cb7f42a73caed0455529e5b5656b2) |
| `syslog.tls_server.mtls_enable.certificate` | [syslog.tls_server.mtls_enable.certificate](data-sources--log_receiver--reference--group-001.md#canonical-f1d3649e2e51151abab3d5fc464e521cf3827524d867b6527c2b55310fc98952) |
| `syslog.tls_server.mtls_enable.key_url` | [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-b309e4a206250b4bf60eb52d2b12d964116d3b7533239a4f79a975ce0fbc493d) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-140ee3e493aa3b8d89d2aeae61d362e4c38486132f4edcc0cbd16089fa8eda9f) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--log_receiver--reference--group-001.md#canonical-e37d83d7f13ef295a976fd457b8f180976c613be86b4f6087f331c009d5f137c) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location](data-sources--log_receiver--reference--group-001.md#canonical-5215cb6f83c3e0b0a9070ae8baa6e3df9fb732d268768f1fee20d85fd23bd011) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--log_receiver--reference--group-001.md#canonical-5e95412dcdc8acef5e25400807165f233efc4b57c4825ce008918bb912b0e64e) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-178c2ee30d939379eb9a9dd9d5f577222dd8f87f03d595fef74667186a98d855) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--log_receiver--reference--group-001.md#canonical-1fd9037a644a277b3f09ed2fafa6708530d382d65f86ff391a2c7057ca30f09a) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.url` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.url](data-sources--log_receiver--reference--group-001.md#canonical-ef1eb26cf06a2e45269c71e7c3c4dd5078fd2fdbb3aa829403e84100980199de) |
| `syslog.tls_server.port` | [syslog.tls_server.port](data-sources--log_receiver--reference--group-001.md#canonical-77f842488e5878705e8eb70cc8e8cc542af1ab2e9be4738ca4a8ebe2a77695f7) |
| `syslog.tls_server.server_name` | [syslog.tls_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-92dc8245f062be1921447a734dbf1d5d66026834445064447057895c0086e851) |
| `syslog.tls_server.trusted_ca_url` | [syslog.tls_server.trusted_ca_url](data-sources--log_receiver--reference--group-001.md#canonical-454ed33639258618c02da010ac6fa2f47523783503c186d22c4903e0fac6eb9c) |
| `syslog.tls_server.volterra_ca` | [syslog.tls_server.volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-7d7155709815e214bcb10aeef7f13984102331434c0f18642a99f08f945fb81e) |
| `syslog.udp_server` | [syslog.udp_server](data-sources--log_receiver--reference--group-001.md#canonical-a2f383f51ee92bc744b6095257a150f8b2c2cf0c35db72b543e4a4cd0d371281) |
| `syslog.udp_server.port` | [syslog.udp_server.port](data-sources--log_receiver--reference--group-001.md#canonical-4bb74d27962f2d500bc4ddc5345f829ce77b295047135a5c9756af2a03940b18) |
| `syslog.udp_server.server_name` | [syslog.udp_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-a088dfb8aeabb40aa586a38ecbd597839686a20fc194eb20a5103a5dcd68ce42) |

<a id="canonical-a7326a85a0719c3356237b25e199c2ba451614ecb454fea67885be842b44de0b"></a>

## Next pages — Property reference / 3245db645141 / 11

- [site_local](data-sources--log_receiver--reference--group-001.md#canonical-41de8992f397c6ba3c8aaddb0fcbc4ee00cb085c1b8148cf5b89c8b33ed502e3)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-41de8992f397c6ba3c8aaddb0fcbc4ee00cb085c1b8148cf5b89c8b33ed502e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c96f710ce6816b2546dbe8c0c01ab40f63491a1c329882cc1c00a91d695b9a7"></a>

## site_local — site_local / 24eb288562c0 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- site_local

<a id="canonical-236a21f0a33913fd3ea3159f961ac14e418ab85d90d2b99ce5943d1817500d12"></a>

Type: `"single"`. Computed.

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

<a id="canonical-b00ca28f708849ea5c142133e2b4a6eedcbb2e260b9e850d7df99fe6949736f9"></a>

## Direct properties — site_local / 24eb288562c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc9d449a4c3bcc6002cf90ee22e4b3b3e170874515dccf2a1fe80c89797cf916"></a>

## Next pages — site_local / 24eb288562c0 / 4

- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19baf01ee8f504fae2be5c54c0c3aa77a8001c797fb7eb677b9408f237491625"></a>

## syslog — syslog / 889b384160a1 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- syslog

<a id="canonical-bff263eba5710055aa83f1d3dcab925cde2f92e5ccaab7a3ee5a1e0a44390f37"></a>

Type: `"single"`. Computed.

Syslog Server Configuration. Configuration for syslog server.

Upstream description:

Configuration for syslog server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

<a id="canonical-3da043aa0a5125210b260bf8d1c43d8e50f077e1a8843478b08e333274c4b480"></a>

## Direct properties — syslog / 889b384160a1 / 3

<a id="canonical-126a944329d27421170ffb0c77783b366b0b4d31895061b1da120c5debfe7a75"></a>

<a id="canonical-86832a57fa4f936d6fb613641f434166afd0a8f16e2a382238458fb92044d007"></a>

## syslog_rfc5424 property — syslog / 889b384160a1 / 4

Type: `"number"`. Computed.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Upstream description:

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 268435456,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 408
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  }
}
```

- [tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-77d5a3b7990c7261a62ebd17940158cb792a0192ccefa8c05a24f365ef75d326): complete subsection reference.

- [tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e): complete subsection reference.

- [udp_server](data-sources--log_receiver--reference--group-001.md#canonical-21eceb7c421b14d2eba798d112127a36ea7ffa72b094ee94b552236cc48fd498): complete subsection reference.

<a id="canonical-cab94f2a1427b62a517f862210ab237b915f57407e19c78688346f96b4bae677"></a>

## Next pages — syslog / 889b384160a1 / 5

- [syslog.tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-77d5a3b7990c7261a62ebd17940158cb792a0192ccefa8c05a24f365ef75d326)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [syslog.udp_server](data-sources--log_receiver--reference--group-001.md#canonical-21eceb7c421b14d2eba798d112127a36ea7ffa72b094ee94b552236cc48fd498)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-77d5a3b7990c7261a62ebd17940158cb792a0192ccefa8c05a24f365ef75d326"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b70542c66de54d5a8a8d76531c4ce4f39ec52fad0c0d6dd52175f4042256bdb"></a>

## syslog.tcp_server — syslog.tcp_server / 2fe8dc0e26ff / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- syslog.tcp_server

<a id="canonical-d1ef828c022f7bbe95ef721f6b4089df8007e4d77b106f75049c2022e4a34d5a"></a>

Type: `"single"`. Computed.

TCP Server name and Port Number. Name and port number for a TCP server.

Upstream description:

Name and port number for a TCP server.

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

<a id="canonical-8e8691309610ac391e481a8b367d4cf1511c201a3f9cfe32c370da9aeeb59a48"></a>

## Direct properties — syslog.tcp_server / 2fe8dc0e26ff / 3

<a id="canonical-29cab7fec224f61289fd4f0934619653787cb0fbc3d4b8d6d5f58074a713c4a8"></a>

<a id="canonical-0ce339918925f217f317201748b5edfbc6da97e1a47e3f60c26cf594e39b5add"></a>

## port property — syslog.tcp_server / 2fe8dc0e26ff / 4

Type: `"number"`. Computed.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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

<a id="canonical-95eb972e2aa0a80e1ed0fbbfd22df573b7eb3fa706b47575ab718065c74cc2b7"></a>

<a id="canonical-c6a51e907b58da0b76912ee31ae9ab10012eff2c30919285ee9457f3d3d7e819"></a>

## server_name property — syslog.tcp_server / 2fe8dc0e26ff / 5

Type: `"string"`. Computed.

Server name is fully qualified domain name or IP address of the server.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-fedbae9b460b5d3af8a3aaa124c4129d0a07e031b124964f1cb83d7a4c4a9b9c"></a>

## Next pages — syslog.tcp_server / 2fe8dc0e26ff / 6

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0e7c9bfa57a0644ec7a1b86cf066fbd4e737343de8fdd21c6cafbb8e7f3f124"></a>

## syslog.tls_server — syslog.tls_server / 707e61f6a740 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- syslog.tls_server

<a id="canonical-b9a7bd2a1a194179caa2c865db074c6896ef07d0fce1d2f25ba81826f83c4579"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

<a id="canonical-f5e3e366717be2c3431fd89ba9a50d82b11c23bb8bb11e90a629c300e1385cd0"></a>

## Direct properties — syslog.tls_server / 707e61f6a740 / 3

- [default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-53b37b184d1f59953869cb72cfac059b5d4f018d77226c19aae24b0759469016): complete subsection reference.

- [default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-fc2045ba7407723a8359691fa0e24e7c569ec82f3788db3030482d8a3e07377a): complete subsection reference.

- [mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-170d23314086bcb9e696175561264f28e50d0507f3a3632bc28b9ef5a26e5cf1): complete subsection reference.

- [mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427): complete subsection reference.

<a id="canonical-77f842488e5878705e8eb70cc8e8cc542af1ab2e9be4738ca4a8ebe2a77695f7"></a>

<a id="canonical-3c035b9c675ecaac37682f6699ba3b71d91eadb9764fe43010cf2c804b91a4df"></a>

## port property — syslog.tls_server / 707e61f6a740 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-92dc8245f062be1921447a734dbf1d5d66026834445064447057895c0086e851"></a>

<a id="canonical-0bb8b427975b79df0d2cc6a02c5e2651cba2f81f775aad7c4a1c8f0f620f9b23"></a>

## server_name property — syslog.tls_server / 707e61f6a740 / 5

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-454ed33639258618c02da010ac6fa2f47523783503c186d22c4903e0fac6eb9c"></a>

<a id="canonical-c778cfab01add3d0ea330de16cb6dc322c1d84a1b98132780310169dbdeee072"></a>

## trusted_ca_url property — syslog.tls_server / 707e61f6a740 / 6

Type: `"string"`. Computed.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-3e63a05b5136467cf96452992874971c6aa1da0800d26bc5d1d535b332a8baa3): complete subsection reference.

<a id="canonical-bfc4483e310c4a68a18ef4c9738d2e06d2955220f4ccf24af917de308d0ee30e"></a>

## Next pages — syslog.tls_server / 707e61f6a740 / 7

- [syslog.tls_server.default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-53b37b184d1f59953869cb72cfac059b5d4f018d77226c19aae24b0759469016)
- [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-fc2045ba7407723a8359691fa0e24e7c569ec82f3788db3030482d8a3e07377a)
- [syslog.tls_server.mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-170d23314086bcb9e696175561264f28e50d0507f3a3632bc28b9ef5a26e5cf1)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427)
- [syslog.tls_server.volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-3e63a05b5136467cf96452992874971c6aa1da0800d26bc5d1d535b332a8baa3)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-53b37b184d1f59953869cb72cfac059b5d4f018d77226c19aae24b0759469016"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd1771b63aa076a1853474e62b660a30d847328f752791cfc7b3b543d83ab568"></a>

## syslog.tls_server.default_https_port — syslog.tls_server.default_https_port / 066de0026085 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- syslog.tls_server.default_https_port

<a id="canonical-78782507dcbf075d2059596bc58fc85cf9e7a1e394751e3b77df6734fd376c6c"></a>

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

<a id="canonical-74e44b29cd7a981e2a6e75401325bbdb0986738c15ab0417aa737cb7786e4d5a"></a>

## Direct properties — syslog.tls_server.default_https_port / 066de0026085 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b06afd77e7690a27976678c2e45608f4658a92b5ea3c6ad7e14ff42a578cccae"></a>

## Next pages — syslog.tls_server.default_https_port / 066de0026085 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-fc2045ba7407723a8359691fa0e24e7c569ec82f3788db3030482d8a3e07377a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b0dd0615a5f20e4caf8ce755ef7aa964ed809726e246270f1e6e3b1975f3fc0"></a>

## syslog.tls_server.default_syslog_tls_port — syslog.tls_server.default_syslog_tls_port / 347efd16043e / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- syslog.tls_server.default_syslog_tls_port

<a id="canonical-4737c459ba38398436871b72144f2c3fa9fc9c15d4be3f2a5a6afb36f2a50b6c"></a>

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

<a id="canonical-c36ca640c9260474a6a77b49c9b6d8a7096ac4d28754b76e2c016d048540216b"></a>

## Direct properties — syslog.tls_server.default_syslog_tls_port / 347efd16043e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44deee2796b2beef9e190a55626690d6d22ed11d8bfb2b68d3b137c9c303599c"></a>

## Next pages — syslog.tls_server.default_syslog_tls_port / 347efd16043e / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-170d23314086bcb9e696175561264f28e50d0507f3a3632bc28b9ef5a26e5cf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e18525d23f6b75a3ebdfae17f2c0c3d4cf6c948531b50e65a50a2739e7bc32fb"></a>

## syslog.tls_server.mtls_disabled — syslog.tls_server.mtls_disabled / 9c7508c2333b / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- syslog.tls_server.mtls_disabled

<a id="canonical-1bf305c83aabc3e79f183d7f98504d101f58ea6d5a1d8ed41924236886b46dfc"></a>

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

<a id="canonical-a6edbcdd6becde2762521c6456e3a802455b9f9e70146a0ef85eaf2adb3b81ad"></a>

## Direct properties — syslog.tls_server.mtls_disabled / 9c7508c2333b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ef6a2cd1e119eba33df4a3d17140fd82f9fbca969f0d790e541f51e350aa34b"></a>

## Next pages — syslog.tls_server.mtls_disabled / 9c7508c2333b / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ab052023e6bf07d92ceee6dc32a0ba15d8680196bf26c6f6aba0fdf2e40ad07"></a>

## syslog.tls_server.mtls_enable — syslog.tls_server.mtls_enable / 0381481c20c5 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- syslog.tls_server.mtls_enable

<a id="canonical-414958da61e7cbe7ae2248370675c7c3d66cb7f42a73caed0455529e5b5656b2"></a>

Type: `"single"`. Computed.

Configuration parameter for mtls enable.

Upstream description:

TLS config for client.

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

<a id="canonical-9297afabfe6eda8c7b615a23ecace551e884d3171cd899973ee9b4fd7786b6fb"></a>

## Direct properties — syslog.tls_server.mtls_enable / 0381481c20c5 / 3

<a id="canonical-f1d3649e2e51151abab3d5fc464e521cf3827524d867b6527c2b55310fc98952"></a>

<a id="canonical-4f1e9e5584b660c97cc7166bb510fec6514fd3af27138da379a8f533ec381d2c"></a>

## certificate property — syslog.tls_server.mtls_enable / 0381481c20c5 / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543): complete subsection reference.

<a id="canonical-10b16af3ba1ca4dbeaa35de1aca1613bc9d015a951cf3c313f92c64ad9d7bb69"></a>

## Next pages — syslog.tls_server.mtls_enable / 0381481c20c5 / 5

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a529b9d29255abca326536a129233a1d1d1442adb90ab05ff335e36a0b10f3c5"></a>

## syslog.tls_server.mtls_enable.key_url — syslog.tls_server.mtls_enable.key_url / 444fac660529 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427)
- syslog.tls_server.mtls_enable.key_url

<a id="canonical-b309e4a206250b4bf60eb52d2b12d964116d3b7533239a4f79a975ce0fbc493d"></a>

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

<a id="canonical-31adb78f1de567015915fad579222475753ba726c5350726e073601b69f23d1e"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url / 444fac660529 / 3

- [blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-382c310005d1d260ef7f5b5a31f900460ca67aed6e65e91ad8c5e8d8e53ed1d1): complete subsection reference.

- [clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-62b8875a06b25095530117cf86497e5a46107ac356f12779b3d807df348efb61): complete subsection reference.

<a id="canonical-3f5c6d126a57439d47662ccece110112f1b2706bef029ea3089026214fa43b7c"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url / 444fac660529 / 4

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-382c310005d1d260ef7f5b5a31f900460ca67aed6e65e91ad8c5e8d8e53ed1d1)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-62b8875a06b25095530117cf86497e5a46107ac356f12779b3d807df348efb61)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-382c310005d1d260ef7f5b5a31f900460ca67aed6e65e91ad8c5e8d8e53ed1d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f32dc4c65ae4106350e3302584ef1ee6cf6b045b1a6c61374fd77a897b527b8"></a>

## syslog.tls_server.mtls_enable.key_url.blindfold_secret_info — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427)
- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543)
- syslog.tls_server.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-140ee3e493aa3b8d89d2aeae61d362e4c38486132f4edcc0cbd16089fa8eda9f"></a>

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

<a id="canonical-cde0215c8e8df63a166aabc092d40e214835ca8d2a45cc02955776ef154c1d5f"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 3

<a id="canonical-e37d83d7f13ef295a976fd457b8f180976c613be86b4f6087f331c009d5f137c"></a>

<a id="canonical-510424e8d79d5b1f5a2befdab753dd5b140821f888743014758f8f66351acc3e"></a>

## decryption_provider property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 4

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

<a id="canonical-5215cb6f83c3e0b0a9070ae8baa6e3df9fb732d268768f1fee20d85fd23bd011"></a>

<a id="canonical-b8102a209d2ad627842dc2c62c00bd9e81c7be42b4f07cbe1da315ef9007cbe5"></a>

## location property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 5

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

<a id="canonical-5e95412dcdc8acef5e25400807165f233efc4b57c4825ce008918bb912b0e64e"></a>

<a id="canonical-3c02f65533572a028b35baa35325144c2f1a1e5149a77391c9eb5fcac2509fc7"></a>

## store_provider property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 6

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

<a id="canonical-bdbacd468f832670f400195063280976318e1d229fb5a146dae697f61b828440"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / 52bb4662576e / 7

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-62b8875a06b25095530117cf86497e5a46107ac356f12779b3d807df348efb61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a74c7d0ebd451eef2464b6db68c85efdd02a2cad986a39874de35af47306f26a"></a>

## syslog.tls_server.mtls_enable.key_url.clear_secret_info — syslog.tls_server.mtls_enable.key_url.clear_secret_info / 0d27b01e4c5d / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-09350d6920d3ad24437b6245b02371f566be2a4739ac6b27a0f34d13a623e427)
- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543)
- syslog.tls_server.mtls_enable.key_url.clear_secret_info

<a id="canonical-178c2ee30d939379eb9a9dd9d5f577222dd8f87f03d595fef74667186a98d855"></a>

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

<a id="canonical-49da3f65be88d48c22533299b08c5af9df54202280990386971f2149e5276552"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url.clear_secret_info / 0d27b01e4c5d / 3

<a id="canonical-1fd9037a644a277b3f09ed2fafa6708530d382d65f86ff391a2c7057ca30f09a"></a>

<a id="canonical-f563de2346f9523be4c31aa996ca3c6d3eb7d9d195a9ff633468cfeb66d8bf68"></a>

## provider_ref property — syslog.tls_server.mtls_enable.key_url.clear_secret_info / 0d27b01e4c5d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ef1eb26cf06a2e45269c71e7c3c4dd5078fd2fdbb3aa829403e84100980199de"></a>

<a id="canonical-524267c3ae2f90b9c72fbb80ae17aebb19af33adb7a365f1adb4312e87bb2015"></a>

## url property — syslog.tls_server.mtls_enable.key_url.clear_secret_info / 0d27b01e4c5d / 5

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

<a id="canonical-850d7a92f807a1ba42e42fa63dc37699efeb842dbbd6e07e42c79ac06c4af2a6"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url.clear_secret_info / 0d27b01e4c5d / 6

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-606dd31028ef8bda6b3f2aa64d31eaa5997fcff924ac0509479ff4308e4a5543)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-3e63a05b5136467cf96452992874971c6aa1da0800d26bc5d1d535b332a8baa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4289ce0596d1f0dfe8a4f0e4b01275093031c27b14f264bcff4e0d91574d26b"></a>

## syslog.tls_server.volterra_ca — syslog.tls_server.volterra_ca / f2d67dadd6c2 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- syslog.tls_server.volterra_ca

<a id="canonical-7d7155709815e214bcb10aeef7f13984102331434c0f18642a99f08f945fb81e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra ca.

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

<a id="canonical-bb166e2736129f539e294e86087a7dd93c7830981decd9cbf918201dd69c4957"></a>

## Direct properties — syslog.tls_server.volterra_ca / f2d67dadd6c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33ab978a6665cc8c811e39c5824c66255977891339991a66b68336184ac6a029"></a>

## Next pages — syslog.tls_server.volterra_ca / f2d67dadd6c2 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-95cc860a770350e5413f7907755c3675b9d94763f385e920120568c436672c0e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-21eceb7c421b14d2eba798d112127a36ea7ffa72b094ee94b552236cc48fd498"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22eda29e01cc01a16657558f7d279ef20d7362bc10ed294f92bf36f4a992e177"></a>

## syslog.udp_server — syslog.udp_server / 3a8dab66e572 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- syslog.udp_server

<a id="canonical-a2f383f51ee92bc744b6095257a150f8b2c2cf0c35db72b543e4a4cd0d371281"></a>

Type: `"single"`. Computed.

UDP Server Name and Port Number. Name and port number for a UDP server.

Upstream description:

Name and port number for a UDP server.

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

<a id="canonical-ce2fae5993f878b9b3c72daacb801c8a71da0286c6ea3a2724320b2ca669f0d8"></a>

## Direct properties — syslog.udp_server / 3a8dab66e572 / 3

<a id="canonical-4bb74d27962f2d500bc4ddc5345f829ce77b295047135a5c9756af2a03940b18"></a>

<a id="canonical-867a2a678e6d9ca147e43e3a2179fa069418c60aadd9add040801d604d33a921"></a>

## port property — syslog.udp_server / 3a8dab66e572 / 4

Type: `"number"`. Computed.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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

<a id="canonical-a088dfb8aeabb40aa586a38ecbd597839686a20fc194eb20a5103a5dcd68ce42"></a>

<a id="canonical-08c8c825030928aa7db8a9f965a76a7dd32b49de9a091cf35a8cca33193aa213"></a>

## server_name property — syslog.udp_server / 3a8dab66e572 / 5

Type: `"string"`. Computed.

Server name is fully qualified domain name or IP address of the server.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-d6a0f7dda061594ec483c31ba9a1011c0f7de95095882a4db25411f4c5c80851"></a>

## Next pages — syslog.udp_server / 3a8dab66e572 / 6

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-afcc9e39b7ac4a917063dccd356d3bbee1f095fbdd1822942f04280ce8e3627c)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
