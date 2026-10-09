---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1200131021300310-0221222201133210-2313033000233113-2332330201301121-3132310022111213-3310103230231121-3311013330100311-1001223232332301"></a>

## Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0202230313313231-3303023000001100-1120130021320133-2121131122021203-1320311133323112-0301123021111132-0113132101310021-0232100132012312"></a>

### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2021211022301021-0031031302031231-1302121022133222-2000032330131000-3130000102310323-1101311320212312-3010201130213020-0301220330223020"></a>

<a id="canonical-3120230230002222-1020000322101023-2010121103212333-2233112232220131-1220300033121120-3003121211131113-1032223332333201-1232322233220121"></a>

### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3310213112111003-1221110110220102-0011000232130121-1303023101231310-1132320302113301-3233212031102002-2203001331233203-1103132020233312"></a>

<a id="canonical-2101003011130021-0102100101222003-2330311321332121-1020223003131030-3320023123033002-1020232013010231-1200122230203313-1210330123232301"></a>

### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2132112002132230-1001001332012103-0232001312331023-3021210022233330-2313331132001312-0311010013230320-1310033023203233-0023310012010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-3312122210323110-2022033001322111-0210020213112221-3111210310130011-1323323323032313-3320331113331030-1103323323113210-3000312012021303)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3303103030013111-3323301002131111-1023112110111323-0220302100033332-0132213322113223-2210200023101230-2203203122210112-2111113133110122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002203312333122-2100001323023122-2203312022001130-0320021210301011-2330023312033010-2210223100033110-2132121022200020-1212012013230232"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3230033111221113-0030120130221232-0202112121111111-1220332001030000-3021201312230010-0113311022232032-1202121102122303-3330022210002021"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1312033023110120-3111200303202001-2230200333210302-2010122021300100-1022211121231323-0021210101012323-1013311302233120-1031020030202311"></a>

<a id="canonical-2221301322213201-0322130220312323-1311030223110133-0021003330330321-3302212023331313-3100230322331203-1223121111213101-1101033130112220"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0020003122102012-0323213211021111-3103300212000331-2133023113001032-3323011311000002-1300013002231113-3133302101120331-3121200113232223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

<a id="canonical-0222011222031303-3122210120003303-1311201020211311-0231101300021222-1302312110113021-2111222132100220-2011103023321103-1330310010303312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- https.tls_cert_options.tls_inline_params.tls_config

<a id="canonical-0330301320313233-2333312211331000-0112212102330032-3202123010201000-1013020000011200-2211202322111233-1210320332203223-1311333022311213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100102112010032-2130223110113203-3122332103302033-2030312210221230-2030022120121330-3100301112130313-3333121323102220-3032103233002213"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_config`

- [custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0303300130002033-2230033101322323-1020021222132312-3003111302100122-1203211001331203-2330002121233213-0303300111122303-0320012132120223): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-1310212110230011-2220113001102310-1212022310013003-0020122010321210-2320231321211113-1312011210201102-1232132322131320-2331310020020030): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-1223303130002132-2223301101131033-3230221212103213-1311011112113312-1330210031312201-0012301011331201-1212123120023000-0201110203120133): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3313101012232102-0323331033200120-3022330112300111-3320023003211112-1001033100223011-2212211110132331-0110011033131023-0001000233111332): complete subsection reference.

<a id="canonical-0303300130002033-2230033101322323-1020021222132312-3003111302100122-1203211001331203-2330002121233213-0303300111122303-0320012132120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233)
- https.tls_cert_options.tls_inline_params.tls_config.custom_security

<a id="canonical-3323321222210210-1322333011002000-2333122003302210-1302301331031232-1331323220200303-1103230123021330-0303321322301123-3302130230012133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230232220022232-0311021233220112-2103220130313200-1131302031220302-2321330001001031-0321132022021211-0000002030301012-0021012130331011"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_config.custom_security`

<a id="canonical-0031201132111323-0211303222030122-0211111011032301-1301003201030120-3232120123130210-3023312332223300-2212210113100331-1331021002302322"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202333033111031-1103002301021022-1130201112211230-0322313021001323-1132312112131013-2101322122112332-2222130031331320-2211312132012002"></a>

<a id="canonical-1031323210131313-3231222111330203-0103321210132100-3300232003221032-1210010320132033-3022201010103320-3330013301022321-2023111131332000"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0313021202322112-0212102022122300-3103132213000031-3223001100121012-3103323310122131-1201322222122003-2333230033220100-0131203232112013"></a>

<a id="canonical-1200033100031131-0310102022012303-1320200001133110-1013333200021133-0130110121300201-3203021102303232-1133330213222132-2103311220211123"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-1310212110230011-2220113001102310-1212022310013003-0020122010321210-2320231321211113-1312011210201102-1232132322131320-2331310020020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233)
- https.tls_cert_options.tls_inline_params.tls_config.default_security

<a id="canonical-1010112123103033-0232232321011330-1131031122033022-1322333333022200-1321222100332233-0132032032213001-0211031203032001-1221113210200311"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223303130002132-2223301101131033-3230221212103213-1311011112113312-1330210031312201-0012301011331201-1212123120023000-0201110203120133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233)
- https.tls_cert_options.tls_inline_params.tls_config.low_security

<a id="canonical-2312032313312021-0232010030221230-0011031323202133-2203101010031312-2332231130000013-3321203121332103-3210301231100123-1013011021233313"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313101012232102-0323331033200120-3022330112300111-3320023003211112-1001033100223011-2212211110132331-0110011033131023-0001000233111332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233)
- https.tls_cert_options.tls_inline_params.tls_config.medium_security

<a id="canonical-3223000000021031-0311203302302030-3310231133301311-1311303112000031-2310012223121130-3333132033333233-0033022222020333-2000131113322100"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- https.tls_cert_options.tls_inline_params.use_mtls

<a id="canonical-2001301101313230-2002032123202333-0131000020322320-3020010003100130-1002232211200021-2322101201302121-3131312033331133-1120323030213230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011221132331100-0223031101322211-2312031130221212-0111020033323221-3110110303131030-0232003020330000-1201330210011000-3110013102121112"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls`

<a id="canonical-1122212223332133-3330031232021310-1103332301123113-3323110320031002-0022333113131230-1103111200321121-0030323232200320-0103301331230201"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

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

- [crl](resources--cdn_loadbalancer--reference--group-012.md#canonical-1011320000232202-0013313233031300-0113032101302230-2122001202033002-0311332202310102-0333323210231011-0010123330301231-0330100200331123): complete subsection reference.

- [no_crl](resources--cdn_loadbalancer--reference--group-012.md#canonical-3000113323222102-0303013030231122-1033312330302232-2003132122111322-2210021201322231-2030022110320120-2221302231002302-0330001332322203): complete subsection reference.

- [trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-1021000122221113-3311120000201212-3310012230213033-3123320331112332-3232200020010011-3231230132220331-3110102100021110-2122001130221132): complete subsection reference.

<a id="canonical-3012323332103231-2101230212120110-1212333023232020-2100303210033123-1101010202213302-1021103013121211-0303100213032100-2320111010301213"></a>

<a id="canonical-1112231133232333-3333031012212113-1130132212321111-0201220302011203-0021103301103222-0022122003020202-0121000033210130-2333010023001011"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [xfcc_disabled](resources--cdn_loadbalancer--reference--group-012.md#canonical-2312202101001211-3001223212213001-2031032001023323-1231330331011211-1112001131301003-3230002301213110-2120131103013221-0103112100130010): complete subsection reference.

- [xfcc_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0131022002232322-0233102110002221-1102101133120210-2130230201023311-0233121301020101-0210012320300220-1212121331311020-2023013211221202): complete subsection reference.

<a id="canonical-1011320000232202-0013313233031300-0113032101302230-2122001202033002-0311332202310102-0333323210231011-0010123330301231-0330100200331123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211)
- https.tls_cert_options.tls_inline_params.use_mtls.crl

<a id="canonical-3223331312123203-2023113133213222-0120232113302312-3202232311120131-1230031303210033-3313013323223103-0213232133113312-2101112030023310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302002023332001-2103000000133332-0133210001020120-2233223201000311-3013212220222313-1221212122312111-3110000022132221-1111012013123133"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.crl`

<a id="canonical-3022120023013131-3213132333222203-3033103112301020-3221332121312203-3312210102001133-2101202230201300-0333333210122202-1232233131130101"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1133000103322201-3020212111211211-3323033100022030-0030322101200330-3123202311222132-2300010220023333-0033221002111020-0220133031033003"></a>

<a id="canonical-1221012312300011-1313032111012232-3121001020232031-2230320001331222-2212001213122033-1210202231001203-1022111210313320-3023033310120102"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122213300103212-2300112313333301-1023220210232232-3220212302101132-2332201201131031-3103110113220202-3331103331211032-3031112113200333"></a>

<a id="canonical-3002333330030210-1210102100003321-1123301221310332-3032210111010310-2013311000333032-0121100022032031-3002233131333323-0011111323111122"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3000113323222102-0303013030231122-1033312330302232-2003132122111322-2210021201322231-2030022110320120-2221302231002302-0330001332322203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211)
- https.tls_cert_options.tls_inline_params.use_mtls.no_crl

<a id="canonical-0011030003001103-2130201003130300-3303211033020321-1130110301333121-0030022221301113-0101020010330123-1013020032332311-1132022130022203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021000122221113-3311120000201212-3310012230213033-3123320331112332-3232200020010011-3231230132220331-3110102100021110-2122001130221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211)
- https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca

<a id="canonical-3112131300111302-2121022121221202-0013020323201321-2101322123033333-0010003302322320-2232002313103202-1100311330311212-1211321300032222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213301032123330-0000000031023200-3111232322021123-0200300131001122-1202223022202021-3232112323101312-0203211023112332-3122330312211212"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca`

<a id="canonical-3210330002202032-2032020330303300-1202031322201333-1321101320133120-1302213320013113-3220111103311020-2312012131302130-2112312300131201"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2020311331221011-2133310002321202-2031320013213123-2331202203222030-0303212010213233-2111202031312012-1002033231101321-0130113320013311"></a>

<a id="canonical-1113000323010321-2032022131321133-0222212010012230-1230113031231010-0022002022031031-1200012213130023-0000120002011132-0103131030003302"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3003313201211200-2300132003222003-3332001100210101-2221333202031202-1323201002100002-0033211021230123-2301221302030002-1120112330212210"></a>

<a id="canonical-0010033111212313-0323101230001310-3113100232321031-1203223011100302-0022100101311303-3322120003201301-2320132213330000-0230131021311220"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2312202101001211-3001223212213001-2031032001023323-1231330331011211-1112001131301003-3230002301213110-2120131103013221-0103112100130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled

<a id="canonical-2021232200231022-1232020311312012-3120033302220032-3133001223332321-0321221122031113-0313112332302010-1203311112221130-1032010203220001"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131022002232322-0233102110002221-1102101133120210-2130230201023311-0233121301020101-0210012320300220-1212121331311020-2023013211221202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options

<a id="canonical-3232030101331202-2311221130021133-1030312330130132-2300032012123133-1330102323331303-1131023202010223-1021113021331013-2030001000312111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332302123011032-3133122311212202-3133200311202332-3132231110102331-2132302122312231-2001111002022130-1300303203003303-2232220012211332"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options`

<a id="canonical-2102120000021023-3133213330013131-0302330013013101-0210013010333300-0220020102213013-1013023210030101-2103313211310230-3011311331131113"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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

<a id="canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- https_auto_cert

<a id="canonical-1011001132000320-2032211231022312-1103022230201012-2202332311121101-0230203310001320-2210023110030310-2011102022223002-0133123021300313"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTPS CDN distribution with bring your own certificates.

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
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122013311032000-0131300212230012-2130331033230203-3332223113020330-2121003011003210-2032323333233123-3332100221300301-2022212012203302"></a>

### Direct properties for `https_auto_cert`

<a id="canonical-3123101300123103-0120003011000220-2213111212121301-3002012113120030-3123301221220220-2132120200120102-0333313332020323-3213033101033211"></a>

#### `https_auto_cert.add_hsts` property

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

<a id="canonical-3312123122302311-0133201021200302-1300122120001323-2032032003312111-2333030131300111-0301112112210223-1300032010031213-3300203013322210"></a>

<a id="canonical-1320233022301322-0311220303320230-2330300011330010-1320202300121311-1112321003303022-1032003120312030-2123323300333110-1223021221102021"></a>

#### `https_auto_cert.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3330112310003221-2210231203321133-0110232120313131-0321132311312322-0013000323120230-1231130303233303-0030111013100230-0032311220323110): complete subsection reference.

<a id="canonical-3330112310003221-2210231203321133-0110232120313131-0321132311312322-0013000323120230-1231130303233303-0030111013100230-0032311220323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-012.md#canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323)
- https_auto_cert.tls_config

<a id="canonical-3021303101202233-3032130211322033-0021202301130231-2333033030331220-0210033030221231-2312112223030110-2000323232333213-1322111231233323"></a>

Type: `"object"`. single nested block, Optional.

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
  "x-ves-oneof-field-choice": "[\"tls_11_plus\",\"tls_12_plus\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232223030212013-3322201002130302-0212023102322101-1023131201101020-1221121333001030-0303200130212321-3102003102231121-2011222000023003"></a>

### Direct properties for `https_auto_cert.tls_config`

- [tls_11_plus](resources--cdn_loadbalancer--reference--group-012.md#canonical-3232002102011221-1320211320300313-0021120200200321-1333202000211110-2232201022221220-3102123133222121-2303223013322003-0023132200321031): complete subsection reference.

- [tls_12_plus](resources--cdn_loadbalancer--reference--group-012.md#canonical-1103322121320103-2102011000203232-0233313222010210-0003211002221230-2320303121320331-1213013030231213-1232303021202221-0022320100112132): complete subsection reference.

<a id="canonical-3232002102011221-1320211320300313-0021120200200321-1333202000211110-2232201022221220-3102123133222121-2303223013322003-0023132200321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.tls_11_plus` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-012.md#canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323)
- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3330112310003221-2210231203321133-0110232120313131-0321132311312322-0013000323120230-1231130303233303-0030111013100230-0032311220323110)
- https_auto_cert.tls_config.tls_11_plus

<a id="canonical-1032101313320221-1012032013022332-1233020302100103-2033101201203030-3122201110132322-1223022212220102-3321013330221102-2002323322332221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls 11 plus.

Additional upstream details:

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
tls_11_plus = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103322121320103-2102011000203232-0233313222010210-0003211002221230-2320303121320331-1213013030231213-1232303021202221-0022320100112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.tls_12_plus` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-012.md#canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323)
- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3330112310003221-2210231203321133-0110232120313131-0321132311312322-0013000323120230-1231130303233303-0030111013100230-0032311220323110)
- https_auto_cert.tls_config.tls_12_plus

<a id="canonical-0222002202312133-2123303122131330-1000113220230021-0013011131002310-1031332202200012-3000021201231321-1333210200321222-3012103012102310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls 12 plus.

Additional upstream details:

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
tls_12_plus = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013121320212310-0321212022202332-3221102212101112-2133200201303301-2223113031321031-2230211002301113-1023222011212331-2310212110032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- js_challenge

<a id="canonical-1102301031312322-1122311201230301-2021213310112100-3011201130322112-3031202232133333-1230111121313021-2312100212122201-2010022231031120"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322033002021220-1021213311022212-1021232223303301-1312122033020013-1320033321231011-2010032332010320-1330301303301032-3332202120233313"></a>

### Direct properties for `js_challenge`

<a id="canonical-2310033200001103-0110103133022213-2130031333100123-1301031131223322-0133011132301313-3210000123221123-2133233223330013-3002230331033311"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021033310313010-3210230321331200-3133100322322120-2201203330232133-3112031223101101-1011002030321011-2001111020120220-0023100300001220"></a>

<a id="canonical-0211323120033020-3310233333302103-3223112021320320-3133302233132122-1232303102323222-3310300221200030-1021031320220113-0120120132231231"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2022321002033033-1222211220022031-2322312231321001-3313100203332232-1212331321111223-0023212323333213-2021122210101220-0203001032332202"></a>

<a id="canonical-2122102113233110-2200210010220310-3310202012002022-0012013110103212-1000113102012003-1221113122200323-3021133003110230-3120320230302111"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- jwt_validation

<a id="canonical-3010000120333003-3110233200033313-1221101022203110-0331331211223011-0323110201021110-2202203132122323-3111133202330220-0211121311033102"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121121221032111-0001213320322323-3313031223203231-2330021213231230-0022330130330123-0103113031311201-3112101123001032-0131202003232111"></a>

### Direct properties for `jwt_validation`

- [action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022): complete subsection reference.

- [authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121): complete subsection reference.

- [jwks_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3021130213331113-3212213232102020-2003001113221112-1322313333022211-2202133110100233-2211202102121021-2330203102211021-1230100211231210): complete subsection reference.

- [mandatory_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313103132331201-1302213302203123-3300032330210320-3332032111310013-2320320003110320-2120230132023221-2312130313220303-1223233010233300): complete subsection reference.

- [reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131): complete subsection reference.

- [target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323): complete subsection reference.

- [token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002): complete subsection reference.

<a id="canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.action

<a id="canonical-1323120013011333-2220201103302333-3210232213210113-0121022002301320-0203213312201121-2103031230220110-3333032030002111-2201231233211130"></a>

Type: `"object"`. single nested block, Optional.

Action

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

<a id="canonical-2221110232013323-1011301132310031-3103102023322231-0303202312101120-3102131233300110-3212132020132332-2221332223321211-1222332022002302"></a>

### Direct properties for `jwt_validation.action`

- [block](resources--cdn_loadbalancer--reference--group-012.md#canonical-2213210033112131-3323111210023330-2000100231101110-2022222000323321-1101102311333112-2231001313211111-0211223113331110-1101132113220023): complete subsection reference.

- [report](resources--cdn_loadbalancer--reference--group-012.md#canonical-1201112212022020-2033223003230220-3301230021022333-3031213113002113-1313121302333233-3113203121112312-0120303112102212-2123233223223130): complete subsection reference.

<a id="canonical-2213210033112131-3323111210023330-2000100231101110-2022222000323321-1101102311333112-2231001313211111-0211223113331110-1101132113220023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- jwt_validation.action.block

<a id="canonical-1300213012130111-2302002101000212-1231220231232232-1300023223132331-3011201002122010-1022200203332230-3223321210201311-1131330213121123"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201112212022020-2033223003230220-3301230021022333-3031213113002113-1313121302333233-3113203121112312-0120303112102212-2123233223223130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- jwt_validation.action.report

<a id="canonical-1103320213203001-0031310020200121-0331001003230312-1230013211231103-3313003113330333-0003013022011320-3331231303323322-2221332111211032"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.authorization_server

<a id="canonical-1301101223212301-3031321230202031-3010202033010311-0332103231001231-0030131122103310-2022002320010320-1213011120132312-3300113223033011"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

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

<a id="canonical-2021300113033112-1221332020212222-1310021233023022-1223003023103303-1131303223102103-0213012000333210-0132332212000212-2220132232331121"></a>

### Direct properties for `jwt_validation.authorization_server`

- [authorization_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2110120333003020-2233231003102223-0002322100233100-0302222022313331-2223332322031020-3233113122212222-0020231012121211-0202000100031201): complete subsection reference.

<a id="canonical-2110120333003020-2233231003102223-0002322100233100-0302222022313331-2223332322031020-3233113122212222-0020231012121211-0202000100031201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server.authorization_servers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-2303130133022010-2220212010022333-2111131301133022-2001203031101013-3001220030013323-2123021020200031-1022232020020223-0111230332101312"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

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

<a id="canonical-2212131033101002-0030221301122032-1132001300311220-1330133100013130-3011233231211011-2113313022020132-3022201313000123-1311322233313203"></a>

### Direct properties for `jwt_validation.authorization_server.authorization_servers`

<a id="canonical-2030332012113211-2011023201000222-0303022200231101-2010202020313312-2303203320220301-3320112103110313-2320002132133012-2310321301223022"></a>

#### `jwt_validation.authorization_server.authorization_servers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3333231222133331-0100223112121000-2110012013330130-1102230103300311-3100231012313103-1023233332310330-3031231302012232-2001013230113021"></a>

<a id="canonical-3230332332000011-2021230010232013-3210120301031211-3202102113231120-1121312133202213-1011031300302200-1322022000302112-1113130200201023"></a>

#### `jwt_validation.authorization_server.authorization_servers.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3302132101030320-3301312333111001-3022300232301112-0331103332121302-1011123023021100-0300210302332012-3202033022122000-0311032120213212"></a>

<a id="canonical-2300233313310002-0301031031330113-0033222031323320-3123101322000222-1203320223012120-3220212223323212-3321200123303012-3320020010213201"></a>

#### `jwt_validation.authorization_server.authorization_servers.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3021130213331113-3212213232102020-2003001113221112-1322313333022211-2202133110100233-2211202102121021-2330203102211021-1230100211231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.jwks_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.jwks_config

<a id="canonical-0220302331321313-1021303322302132-3111222311013312-0213023112300030-0210231211011210-1320130012203003-1110221030101232-0213233221022212"></a>

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

<a id="canonical-2020330023010223-3320220223000311-3223010312232101-2110011133333330-2231303002133321-1220333220222220-0301303232000011-1310320023232212"></a>

### Direct properties for `jwt_validation.jwks_config`

<a id="canonical-1000021132311230-3122301123332220-0022221032222012-1003333102101301-0000103112133102-1101131000010103-1302212301330101-3010320302311322"></a>

#### `jwt_validation.jwks_config.cleartext` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313103132331201-1302213302203123-3300032330210320-3332032111310013-2320320003110320-2120230132023221-2312130313220303-1223233010233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.mandatory_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.mandatory_claims

<a id="canonical-2333203332013032-3133013113210012-2023130021310200-2212013021031222-1012021102002132-1232023332033131-1212232122221212-2120320101030133"></a>

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

<a id="canonical-3333311322103310-2312200220121231-0020233033001231-1000132313103003-0313033330320132-3012122032131032-2221110032212031-1131221022132231"></a>

### Direct properties for `jwt_validation.mandatory_claims`

<a id="canonical-3102233101031012-2110001222130201-1010112222131103-1102112203002130-0112032200320113-2220211020122223-1330002222132203-1103313101220330"></a>

#### `jwt_validation.mandatory_claims.claim_names` property

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.reserved_claims

<a id="canonical-3212112200320031-2330221301133131-1130100101222321-3330113120121203-1030103110202033-1030220122031003-1213310023233001-1313120322230210"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

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

<a id="canonical-0113122311320110-1300300323321122-2332313120222013-1203220031210000-2230022113101021-0032102301311101-0232100301303002-2011333003210212"></a>

### Direct properties for `jwt_validation.reserved_claims`

- [audience](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100): complete subsection reference.

- [audience_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1213133021031020-1001310301233300-0201202323232203-0310130100002333-3311220123223122-1002211313133332-2203012101100123-2131133012121323): complete subsection reference.

<a id="canonical-2221103212000313-1333311201322302-1332221123230230-1122022312002103-0130300323111231-2102300300330032-2100313313132031-3300221201013303"></a>

<a id="canonical-3330311220211020-3320000030023000-0333112121002001-0212222201300032-2313222112122303-3213002113222202-2013022312332323-2110303322233322"></a>

#### `jwt_validation.reserved_claims.issuer` property

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [issuer_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321023311132211-3113003212232022-0020212230220103-1121320022303333-2101300230003012-3011321131221023-0011313331332012-2132133331131231): complete subsection reference.

- [validate_period_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1111232312303222-0032011222322301-1231332201023300-3300301202110211-1131213211130213-2012023033233123-1003121002331323-2310110002331131): complete subsection reference.

- [validate_period_enable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0112011011200030-1002002021103021-0023122323131033-3331303133200000-0233212311122232-1130013133032100-0300110311310223-1123210303212102): complete subsection reference.

<a id="canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.audience

<a id="canonical-0300301132322231-2230203222131233-2222323000023201-1111010310313331-1030331102020023-1231302222302330-2210300203322031-1313321010011102"></a>

Type: `"object"`. single nested block, Optional.

Audiences

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

<a id="canonical-1112121021333031-2132301303203301-1202120033031012-1310203002030313-0112332033330021-0111123000010102-2213323000010021-3321231201213111"></a>

### Direct properties for `jwt_validation.reserved_claims.audience`

<a id="canonical-1200301323211332-0310021131020323-1332211203312102-1223321201322200-1132031232203131-0210031302131300-2131310013210220-0223311102232322"></a>

#### `jwt_validation.reserved_claims.audience.audiences` property

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1213133021031020-1001310301233300-0201202323232203-0310130100002333-3311220123223122-1002211313133332-2203012101100123-2131133012121323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-2013213021201033-1130101000020020-3211030031203202-3233301032000201-2211322313321133-0101322100331300-2030230032031100-3102120111023300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321023311132211-3113003212232022-0020212230220103-1121320022303333-2101300230003012-3011321131221023-0011313331332012-2132133331131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.issuer_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-1122300310321001-2110220222230333-2210201211331003-3222102321100031-2231211122000121-1310033231312332-1002223113212002-2233232202212103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111232312303222-0032011222322301-1231332201023300-3300301202110211-1131213211130213-2012023033233123-1003121002331323-2310110002331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-0123313021103223-3321120103102331-3301213120202110-3210003033211110-3111003221121221-1313113332221121-1131211120330133-2112331131131211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112011011200030-1002002021103021-0023122323131033-3331303133200000-0233212311122232-1130013133032100-0300110311310223-1123210303212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_enable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-3102112110201103-1131231020102030-2100032030213210-1301002231200123-2210120222323333-0211221121020310-0210330113112200-1320102332133300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.target

<a id="canonical-2111203333133231-0102221210211213-1001311213213320-1331133320211012-3122321001013000-1001201133133203-2210230132302012-2330030300223013"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

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

<a id="canonical-1022331003133211-1021333221233320-1203003010113113-0203332120233213-1020313211030131-0301013212122321-3120203021332203-3222102232020132"></a>

### Direct properties for `jwt_validation.target`

- [all_endpoint](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-012.md#canonical-3131230110321221-3002023032100200-0013122221223303-2133331321332211-1313130012021222-3001303132131030-1303330223100100-0121010000020200): complete subsection reference.

- [base_paths](resources--cdn_loadbalancer--reference--group-012.md#canonical-3202113120131021-0320333000010332-0032112121222220-2203030233320222-3021322010321222-3123302302310302-1310121322300300-2023103213131203): complete subsection reference.

<a id="canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.all_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.all_endpoint

<a id="canonical-3002330221301310-1023102100110000-1221110323321031-1220202323033230-0012031322013320-3033131331232110-2112031330303302-0213302001003322"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131230110321221-3002023032100200-0013122221223303-2133331321332211-1313130012021222-3001303132131030-1303330223100100-0121010000020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.api_groups` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.api_groups

<a id="canonical-1003333003022021-0123320210213200-0231030011303132-2021100303201323-2320113103221000-1100320112303230-1201123000213201-0032123100223112"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

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

<a id="canonical-2220302200022001-3010321033110122-1333332221032233-0231230330220001-0002012100002113-3032333020001221-0121111101303133-2101001312030320"></a>

### Direct properties for `jwt_validation.target.api_groups`

<a id="canonical-2011010113312321-1232232300221232-3011100312223123-1101010332330213-0123211201111331-3110133232112310-2011232032012321-0212232232331332"></a>

#### `jwt_validation.target.api_groups.api_groups` property

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3202113120131021-0320333000010332-0032112121222220-2203030233320222-3021322010321222-3123302302310302-1310121322300300-2023103213131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.base_paths` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.base_paths

<a id="canonical-1002012233312203-2010132302013220-3110220302223123-0102330101233123-3011311321231011-0012323032331023-1230101330313111-1002111130200332"></a>

Type: `"object"`. single nested block, Optional.

Base Paths.

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

<a id="canonical-2012211122103301-0023033021332131-2320322120321302-3032010132310010-1331202212221201-0311020012030123-0000331122123131-2030333313322002"></a>

### Direct properties for `jwt_validation.target.base_paths`

<a id="canonical-1221103003030123-1102223231312223-0130123232021011-3021231210223021-1103210202101322-1101113303013303-2112111121103331-1023100211223010"></a>

#### `jwt_validation.target.base_paths.base_paths` property

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.token_location

<a id="canonical-3222122003223131-3300113130232013-1021013133211032-2220222012232313-0233132133230130-2230113120303112-2210333302201301-1333331330321023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Additional upstream details:

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

<a id="canonical-3002113311110313-2230213023031013-0020312221122333-2200031220200310-0132332130301301-1033121230220210-2133000330001121-2212233103010301"></a>

### Direct properties for `jwt_validation.token_location`

- [bearer_token](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313130223203023-3212313023211320-2030123032320301-2200113333021001-3001003000201002-0020210321111101-1211213033311112-1113332103231330): complete subsection reference.

<a id="canonical-2313130223203023-3212313023211320-2030123032320301-2200113333021001-3001003000201002-0020210321111101-1211213033311112-1113332103231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location.bearer_token` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002)
- jwt_validation.token_location.bearer_token

<a id="canonical-1122001322110203-3110301233333101-2203312111331323-3131233011211321-3010113321113202-3030232013002122-3120101210012133-1132333130101223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302023000110302-3120111302312101-1232211223311220-1031332223001000-2302333000033013-1001023131000303-3232033331020302-2320323022012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_block

<a id="canonical-0100201001123100-2323102022302103-2201020102221003-0300000001122200-2003030213312113-0010303131320223-2011133100020323-0102111133321110"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

Additional upstream details:

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

- [l7_ddos_action_block](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201001123100-2323102022302103-2201020102221003-0300000001122200-2003030213312113-0010303131320223-2011133100020323-0102111133321110)
- [l7_ddos_action_default](resources--cdn_loadbalancer--reference--group-012.md#canonical-3031030021303312-3321123103123321-1201032120110223-1301323023112110-1300131320013230-1110313213213133-2323001312232023-3233331200030031)
- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-2023312331022300-1320031101102101-2232303210212002-3310102232320022-3213121210230333-1311223111301323-2221121220030321-2331213011201131)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230330203013032-1133333031113100-0021112120003221-1230203101320213-0202130233200301-3000332131012213-3203322211203031-1131222233212002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_default

<a id="canonical-3031030021303312-3321123103123321-1201032120110223-1301323023112110-1300131320013230-1110313213213133-2323001312232023-3233331200030031"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230011220011130-3203232003333030-0210211220332212-1300220223001301-2310021232213132-0112010313031231-3000331122302003-2133202101030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_js_challenge

<a id="canonical-2023312331022300-1320031101102101-2232303210212002-3310102232320022-3213121210230333-1311223111301323-2221121220030321-2331213011201131"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-3300330313230010-2003311202121011-1111211300102330-2321021303110321-0133213010010230-1020101231332103-0202110033032203-2321000331230100"></a>

### Direct properties for `l7_ddos_action_js_challenge`

<a id="canonical-2301032201201232-3301200030001211-1322221010210101-3013220220013110-1132312132330130-1111100130002132-2003130003331012-1201133333132322"></a>

#### `l7_ddos_action_js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0313311313220003-1013310212321020-2203011303210033-0213110011012201-0030022202002003-0103312312110211-3130301323113230-3213020232311332"></a>

<a id="canonical-3001202123100011-3201021132003310-1233221322200121-3121312310312123-0030010211010021-3030023122312232-0023101231031010-3222000212131020"></a>

#### `l7_ddos_action_js_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0202131023303313-2103100021130301-1323121013100033-1200212103313230-2322122220333013-3003120120321030-0100220303233022-1001203030222301"></a>

<a id="canonical-3323303011103300-2302003003133313-2102330101322100-2310101232122122-1103120321023012-0022120303102131-3113001130332323-1200201202310220"></a>

#### `l7_ddos_action_js_challenge.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3312000011031203-2231300221302120-2103312210100030-2210320110231313-2210213211211221-2111110102013130-3032201012110003-2023110221100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- no_challenge

<a id="canonical-3122322231130001-0110033223312111-0211313102301220-1122203023011032-2300330122000333-2313023000331020-1121211003020210-1230333230203311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

Additional upstream details:

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
no_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123112331011010-0132220300333311-0001300132330131-1132331012130223-3311232302020312-2200131102201331-2233330000312303-0132221122110323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- no_service_policies

<a id="canonical-1013310300231013-2321212333203010-0332022031203123-3112002213301121-3010211320310130-3233202333011223-1131111100001030-2310100131120023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

Additional upstream details:

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
no_service_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- origin_pool

<a id="canonical-1333023300232232-3020120010223111-2210032333323302-3202332123100233-2202001011030210-3120233023331030-0011213003323111-2130022000131010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pool.

Additional upstream details:

Origin Pool for the CDN distribution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213100313032122-2332102200122112-1312110310120331-0121331312231011-1310313232103111-3310302132212331-1100033313033133-3311223111320003"></a>

### Direct properties for `origin_pool`

- [more_origin_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0213200101331231-2120013110121131-1210303023133011-1030133313113032-3211300011331111-3213130011023120-2021320133300332-2231122333022102): complete subsection reference.

- [no_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0031320223020112-3322011231321300-2133213030302010-3003303230121300-2012120113121303-3110323013020211-2002003232313303-3223222032032013): complete subsection reference.

<a id="canonical-3121320022332231-0330103022121111-3103320130101311-1321320232331332-2303131132013201-2201010220211332-2221332111321201-2200012313213033"></a>

<a id="canonical-2300011113110010-1033012302100221-0013230102033201-3311120231201100-0200011020113011-2301213220131331-1131010231331032-2030301212030103"></a>

#### `origin_pool.origin_request_timeout` property

Type: `"string"`. Optional.

Configures the time after which a request to the origin will time out waiting for a response.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002100232220132-2323022211233331-1220132030312302-2222333312210131-3203022333311220-1311221201332302-1221221333110130-1020210013322232): complete subsection reference.

- [use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200): complete subsection reference.

<a id="canonical-0213200101331231-2120013110121131-1210303023133011-1030133313113032-3211300011331111-3213130011023120-2021320133300332-2231122333022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.more_origin_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.more_origin_options

<a id="canonical-3302202302203103-1033021301213130-3302230113222131-0033223301201022-1303001133313213-2010300301030201-0321310022221231-0103023322011133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for more origin options.

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
more_origin_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100032232121320-0322202221312302-1213113110231203-0200020120211101-2130130211133010-2031230303132013-0221200331323310-1220102000000221"></a>

### Direct properties for `origin_pool.more_origin_options`

<a id="canonical-2221000300122130-2121103112003221-1221332002221213-1332000023111032-0302331003113013-0121121132200112-1002320311120123-0000030031312132"></a>

#### `origin_pool.more_origin_options.enable_byte_range_request` property

Type: `"bool"`. Optional.

Choice to enable/disable byte range requests towards origin.

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

<a id="canonical-1312220230113033-0103032332110020-2012223323133320-2201031211223023-0102111220021221-2021131211102002-3131111001322201-3312123330330233"></a>

<a id="canonical-0212211012001131-3223122310220033-1200120232022002-1022230331333301-0113112001301222-0113221231203021-1303231021230300-1100111321310113"></a>

#### `origin_pool.more_origin_options.websocket_proxy` property

Type: `"bool"`. Optional.

Option to enable proxying of websocket connections to the origin server.

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

<a id="canonical-0031320223020112-3322011231321300-2133213030302010-3003303230121300-2012120113121303-3110323013020211-2002003232313303-3223222032032013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.no_tls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.no_tls

<a id="canonical-1203001110313320-3220003320211313-0312022231002030-2200322111111013-3133321221023222-3022002203200312-0003011313033303-0020330230103222"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
no_tls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.origin_servers

<a id="canonical-2211120211003300-2321001131221201-3101131113113011-0233112112030003-3031011010231112-3120133123232213-1020310022000002-2213202033111130"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of original servers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111310013023320-0321112130211320-3233331021111003-0300111311212211-0131321003332021-3002020130323020-2120331320213012-0022102332132302"></a>

### Direct properties for `origin_pool.origin_servers`

<a id="canonical-2321022210203333-1310212000331103-3313231200030223-1231322102231330-0203301233103221-2011213300200212-2233231020001013-1302110010001000"></a>

#### `origin_pool.origin_servers.port` property

Type: `"number"`. Optional.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
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

- [public_ip](resources--cdn_loadbalancer--reference--group-012.md#canonical-3331223110233121-1002130131301031-3213012113300333-1110202110221311-3231013323001202-0130301030331302-0121001202011220-3002213210031303): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213): complete subsection reference.

<a id="canonical-3331223110233121-1002130131301031-3213012113300333-1110202110221311-3231013323001202-0130301030331302-0121001202011220-3002213210031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- origin_pool.origin_servers.public_ip

<a id="canonical-3200132003133110-1011032300320213-0303112220121030-0112013201112302-1310310021211202-3010023110132031-0032023313310131-3003232311201220"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012132133112101-0000320301302011-0321331122122302-2133012332212221-0103131202020233-2132101113112000-1020100032100101-0333220012231010"></a>

### Direct properties for `origin_pool.origin_servers.public_ip`

<a id="canonical-2233033111012310-2303103213300232-3030013003220222-1100023201021130-1231221231002311-3103210310031000-2301002020300300-0333030332102333"></a>

#### `origin_pool.origin_servers.public_ip.ip` property

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- origin_pool.origin_servers.public_name

<a id="canonical-1033012323010132-0020013200300212-1000233122311022-1122332310102333-3111012322230212-2321130323031023-1221323101003202-3121121330102003"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012131312020303-3200030002301302-1030102210133202-2010311232033011-1022131101331323-3313331000202121-1113223231102031-1020020213333113"></a>

### Direct properties for `origin_pool.origin_servers.public_name`

<a id="canonical-3333300330303300-1320021323010101-2131200311133003-0231132230223301-2000133010322202-0122110103001230-1003222311121130-0123300000013111"></a>

#### `origin_pool.origin_servers.public_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3132301011020133-2023031023112123-0111310122001310-0121200001201220-1130102113323220-3313301112031100-1322330033312233-1122032110011000"></a>

<a id="canonical-3332202232130231-3223130112003311-2323201232013221-1310012122311111-3113100333302302-0211112212121212-2011321003112013-0030132330013310"></a>

#### `origin_pool.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-1002100232220132-2323022211233331-1220132030312302-2222333312210131-3203022333311220-1311221201332302-1221221333110130-1020210013322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.public_name` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.public_name

<a id="canonical-0303202020030112-0203311230230001-1110330132311010-0233202000131121-2010001002333220-0020320310232100-2333112333000013-2133310230000213"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223310332033300-2330031111330133-2112233201033130-0013102030120030-0202031223222333-1112300211130013-2312120002211232-0102131020110131"></a>

### Direct properties for `origin_pool.public_name`

<a id="canonical-0223300300303210-2023312333120132-1003113331102232-3222121330301001-3233212221022201-2302023201311031-2200300121222323-1002322221321022"></a>

#### `origin_pool.public_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1210021301003202-3101010022320211-0101033203322133-3221313332200211-2030112103321313-2001100011322031-1023310133022220-1301110322031210"></a>

<a id="canonical-1201122012130013-2310301332321112-0112310020303122-2223103233200202-2130231313010111-1212132102023001-0313221011012132-3122033023030202"></a>

#### `origin_pool.public_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.use_tls

<a id="canonical-1031302012301011-3023111332123113-1303121212021210-1332102313013212-1130333010110110-2210222231231211-2332330323210030-0200023310323320"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332310312012312-0100310233222021-1211002122131220-2223022323322311-3331032312222311-3332202000023321-0030031113010131-3312322201223200"></a>

### Direct properties for `origin_pool.use_tls`

- [default_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-0302320330132121-2000211122231110-0233002332102313-0202233120013211-3231303120200033-3312032033120003-1320111203022003-2002321023302011): complete subsection reference.

- [disable_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-1122220121331121-3031112213221221-0221331232102133-0122100330030331-2133220122321030-1302320222123332-3123123132103002-3203320023030012): complete subsection reference.

- [disable_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002301103100330-2321023201300003-0320202012221111-2222033332010301-0123023001120110-2333020322010211-0211221333230303-2313203222223103): complete subsection reference.

<a id="canonical-2122000022123202-1321111232003120-3010213212130002-2331231321313322-2212122212211310-0230223333300030-1103120033112322-2023202302023021"></a>

<a id="canonical-3123032011112112-3303131010113033-1323321330233201-2203120102032011-0212103201113112-1223021202230312-0313023312323212-1103130211013100"></a>

#### `origin_pool.use_tls.max_session_keys` property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0303111300201320-3230001233023023-3331123310130210-0122201131011003-2000202333110003-2300331000111120-1033201112021023-2020010313130103): complete subsection reference.

- [skip_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231030211300212-1000300232012123-2103303232131013-0111201013200331-0200133221231133-0311012120210312-3200300300011023-3232131131321320): complete subsection reference.

<a id="canonical-1201121023312011-2222203223203102-3130222221000033-2311113110013133-1030003103020312-2002200101033232-2222031031132110-2130210131303323"></a>

<a id="canonical-1111302223222231-3220210233003200-2113120222110123-3023213000212321-2120331302022012-2203322031201323-0021031321010232-1202211131230203"></a>

#### `origin_pool.use_tls.sni` property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230): complete subsection reference.

- [use_host_header_as_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-3123303302303123-3202333112001312-3112222203202103-3033010031032020-2300210110111002-1031113222122101-1103113202002302-3331312230112131): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220): complete subsection reference.

- [use_mtls_obj](resources--cdn_loadbalancer--reference--group-012.md#canonical-2031020210010023-1201331221331233-3030210310323100-3110000032123022-2313130031331210-2233320022202102-2211203032032301-1102211332110031): complete subsection reference.

- [use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023): complete subsection reference.

- [volterra_trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031): complete subsection reference.

<a id="canonical-0302320330132121-2000211122231110-0233002332102313-0202233120013211-3231303120200033-3312032033120003-1320111203022003-2002321023302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.default_session_key_caching

<a id="canonical-3023201323332222-3231221130013103-3313312000011330-1233023102200133-1203232123200310-0233203202322033-1113023333222201-1130330231301333"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

Additional upstream details:

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
default_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122220121331121-3031112213221221-0221331232102133-0122100330030331-2133220122321030-1302320222123332-3123123132103002-3203320023030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.disable_session_key_caching

<a id="canonical-1110001133011330-3100221020030101-0011023211230030-0121010121122123-3230312110012112-1132313121031311-3110121122030111-3312011022101222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

Additional upstream details:

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
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002301103100330-2321023201300003-0320202012221111-2222033332010301-0123023001120110-2333020322010211-0211221333230303-2313203222223103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.disable_sni

<a id="canonical-0222101131313000-3012302320321113-1201221102302030-2222022000013302-2013021132013220-0233213033002322-3003221122221203-0222002221130131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

Additional upstream details:

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
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303111300201320-3230001233023023-3331123310130210-0122201131011003-2000202333110003-2300331000111120-1033201112021023-2020010313130103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.no_mtls

<a id="canonical-0222232101212200-0131011321333311-0211030210033330-2121103101122231-2232231103003013-3111320133021011-3321101113330122-2320000132331230"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231030211300212-1000300232012123-2103303232131013-0111201013200331-0200133221231133-0311012120210312-3200300300011023-3232131131321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.skip_server_verification

<a id="canonical-0012012332030002-0112333313030103-2300102222012323-1331300001232100-2120301300121212-0302223312332130-1313110233122330-1122203200223021"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
skip_server_verification = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.tls_config

<a id="canonical-2010031003200021-3000332113323202-1000031233322211-2113212223303132-3332222022033200-0322320111203222-0331122030202032-3110030101012311"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
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

<a id="canonical-0203301000023331-1233310011210300-0020111323222033-1313001300221001-2003233220231113-3333230213113310-3300300313103110-1311023111000203"></a>

### Direct properties for `origin_pool.use_tls.tls_config`

- [custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231112012220003-2322021233233212-2232322232013110-3013300211332211-0231000111131313-2101110200020213-1120133331232301-3312221110130122): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-2131312123312001-3112012002301233-1322130102031332-2232020023022201-3210013231210231-0132023010112132-3212133120110112-3302010133231023): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020300200001212-2220223302331012-2101011113033301-3002031232333231-0001001301131332-0303300111211310-3330303303031013-3202313313203201): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0333113000022310-0232221001030222-3020002200123203-2001313101112132-3203303110010030-1322300331321313-3011021302333320-0102101222312111): complete subsection reference.

<a id="canonical-0231112012220003-2322021233233212-2232322232013110-3013300211332211-0231000111131313-2101110200020213-1120133331232301-3312221110130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.custom_security

<a id="canonical-1200020030320103-1330121231321310-2303323203103013-3101121031112010-3000031332022203-0331322023302122-3113021323013323-0233203203132000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133110331103332-0003202212210010-0103022130323331-2010130133112012-0133022000202131-2311120003032111-2111010333031133-2233310002312323"></a>

### Direct properties for `origin_pool.use_tls.tls_config.custom_security`

<a id="canonical-0003001210322330-1120012102322313-1001222322000313-2033333202200232-1002210210111112-2123101131131222-0102002121123010-0101322301202301"></a>

#### `origin_pool.use_tls.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2002012333201002-0031001323111320-1331310111120020-1130333200112302-3033213322302311-1312321310333230-2123301330112211-0211013001200130"></a>

<a id="canonical-3202300121331021-1232111220021003-2131013010202101-1113200210301001-0023011022022322-2112102231211112-3131322211311000-2331101332322321"></a>

#### `origin_pool.use_tls.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-1122331132303012-3300222303232133-0223333121300212-1213111220030220-0223023320021021-0331333003211223-3120130300030013-3012020222313212"></a>

<a id="canonical-2010122010233020-2102200121313232-0210032301103010-0323333001232310-0220103312131112-2112001331202012-1110111322023011-2211113300120013"></a>

#### `origin_pool.use_tls.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2131312123312001-3112012002301233-1322130102031332-2232020023022201-3210013231210231-0132023010112132-3212133120110112-3302010133231023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.default_security

<a id="canonical-0133102230011213-3132310020000202-0200030231112211-1310332331020002-2221002323323232-3112323100031020-1331030201331111-2033210132300110"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020300200001212-2220223302331012-2101011113033301-3002031232333231-0001001301131332-0303300111211310-3330303303031013-3202313313203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.low_security

<a id="canonical-3130001032321102-3001021322030130-1132132220232300-0123112023233103-0022322023000220-1123010130100313-2220200322301130-3130211130013210"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333113000022310-0232221001030222-3020002200123203-2001313101112132-3203303110010030-1322300331321313-3011021302333320-0102101222312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.medium_security

<a id="canonical-0032002022100211-3011020022221031-1313211322013003-0300233120223012-0003232111313032-1032110130331010-3001201212122000-3111331121211010"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123303302303123-3202333112001312-3112222203202103-3033010031032020-2300210110111002-1031113222122101-1103113202002302-3331312230112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_host_header_as_sni

<a id="canonical-3333103210230101-0131313302121200-1333330203210022-0132011222100001-1021010232010111-0031211000301313-2101133122313100-2001313031120123"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

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
use_host_header_as_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_mtls

<a id="canonical-0311031330111112-1010002223111221-2233002030001110-0333003232333203-2203302030300101-1322203130301121-2112200032222133-3031023131312303"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213301202213132-0133113322211111-0332011122210010-0033021302322130-1322033321002020-2001110211100203-1010113111121030-1110303032233331"></a>

### Direct properties for `origin_pool.use_tls.use_mtls`

- [tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020): complete subsection reference.

<a id="canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-1321030022322100-2222213021121020-1113320103000223-2021312132133301-1012210300010031-3000230100031320-2121311223310003-0032231323131013"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202331212332112-1121302111102120-1312222010031122-3131102303211123-0132130031232020-0212131220200223-3021021113032120-3013033110103212"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates`

- [blindfold](resources--cdn_loadbalancer--reference--group-012.md#canonical-2213232302132010-1003123212311122-3303103232103020-1213303132020302-0330220331302202-2013313313102111-3212030303000132-0121220232120333): complete subsection reference.

<a id="canonical-3131123230300102-2303310123021231-0222013212213331-3101001202211223-0200202333322331-0302311211220331-2320230321331301-3032301333321102"></a>

<a id="canonical-0203100331210131-3201002302321223-1030311302011012-2210210320210100-0212033002111220-1001011000021032-1110331001012022-2313110313230021"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-012.md#canonical-2012310320122332-0012302201101121-1032312111313220-1113220311013313-0100313221310121-1031230003022000-2200023202330130-2230131201302112): complete subsection reference.

<a id="canonical-2002002010323003-3323213311232112-2033020330003313-3223021311131111-1233121211003320-0320220011122013-0333312020310203-3000102121233300"></a>

<a id="canonical-3131310320101201-3113023300012120-0033111013221130-1132033310131112-3203320003112310-2011031120031112-1133301301130101-3100232131003313"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002): complete subsection reference.

- [private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311): complete subsection reference.

- [use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321211001212221-2201311223311330-0203120023100113-3102023123032112-3132003033033032-1313231223221223-0311320211032220-1102210003202231): complete subsection reference.

<a id="canonical-2213232302132010-1003123212311122-3303103232103020-1213303132020302-0330220331302202-2013313313102111-3212030303000132-0121220232120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.blindfold

<a id="canonical-1323221132100230-3300303221100010-1330021321323210-2113231022123313-0021100312221001-2300303022323223-0001011212132220-0322122212303021"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-0302221033020333-3203200120102201-0212313131023120-1302302212003320-3022330131302013-3212233113120130-0010100001030013-0011122210130022"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.blindfold`

<a id="canonical-0203130301123032-3220102010103221-1213322103030113-1303132223131033-0102110031133300-0111310012000223-3033200330320113-1033312213100202"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-3303210213322000-1111331320332310-3010001230230302-0213201012230100-2121201313121203-2022332212001222-2023310321332303-1112320230311331"></a>

<a id="canonical-0212323221232101-3012203033003133-3113203023130310-2331210122301000-3132130330333020-2020201301120121-1220231111300231-1203332010131103"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-3301320330321323-1231331123321002-1023123321103132-3313123333301222-1010033310203002-0021022230100132-3023312321200233-0023300332132233"></a>

<a id="canonical-1002331331013012-2012333220312130-3113023200210221-2230123133330220-3111033232110020-2211111100202221-3321030112111132-1023001111120001"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-2103213123313303-2203322221023100-2332000322321033-2332032221333303-2300113001001310-2102313322313330-1331232303220321-1030113002121212"></a>

<a id="canonical-3210013301302131-3313102011013002-0120100010032221-3010210230231003-0201321000322301-0011130310112132-2223210112311003-2011203233232323"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-3221210332311221-3330302000231102-2213001122313220-3323101333212020-0203310201200133-1003133231302032-2002012311203033-2022211203322302"></a>

<a id="canonical-1000112120111313-3303022332023222-0202102301032102-3312302111133101-1221003312310213-1330221001301323-2223010330110300-0020231223021032"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-1220332010333113-0323031011213333-2011201122113333-1300122012302132-3221101321112330-2012201321110110-2102033103230002-1330303203031001"></a>

<a id="canonical-0321100101330130-3223221130110302-1203330333202110-0000030212331121-1123302202212302-1213321003230323-3020002130010302-1022322123111132"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-1011231120011311-3220102200013230-0030111332002030-3203222021111202-2013302201023210-1022300202201203-2222303130133102-0120011030232310"></a>

<a id="canonical-2110230233101230-0332122330200333-0220323113211133-2212103023102030-2301110111033113-1310113301012301-3121313002330232-3202302032210231"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1303012002220101-2000132300303000-1002322002221111-0321333202021110-0320120101333233-2032031132333222-0223313230000302-2201120330110213"></a>

<a id="canonical-2101001023101231-1331120103300113-0013210103122323-0020233323330323-0230021310010120-3332202120130321-0212203221032211-1312212123113133"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-3022033103211100-2213032002100012-1223321230013232-2310311110001121-2113120313233130-1202130301311212-2102313320320302-2322011232210321"></a>

<a id="canonical-1223313323012010-2010030122113220-1121101133320102-2013002101332333-0011111211112220-1212300222132300-0001302030230300-0001222120003210"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-0021203001110300-2032211330330301-3223313302133122-2233332023320000-3001300032003332-2213321230003221-1130233213332123-3012333111110021"></a>

<a id="canonical-0112332133213121-2111323232201102-3110112212331100-2313310031321213-2001231223010312-0130033113321131-2210003223131332-2000111311213120"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-2331100001022303-1212303131123331-3110310032312312-1121103211003303-2211221311100133-3200122302312313-3033111331222131-1203002200333030"></a>

<a id="canonical-0103202012232131-0322120133120312-0321002013223303-1012202003021122-2133031123200132-1301321123221112-1331210231201223-2012000111320230"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3323021222100230-0211000002022033-0303222010013012-2330002231311301-1020212231221311-2130102301312003-3330121113103131-3100022301331221"></a>

<a id="canonical-0012300130031131-1031103322313110-1123332131002110-0231013012132210-2202031323311112-3033331012220333-1010123032112203-3222312030310311"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2322222203021120-0020011233210011-1321333130131102-0012130301321321-2202201100210311-3211132313111022-2310122203321200-0202023013331222"></a>

<a id="canonical-1110300212222130-2203222022113232-0222010232311032-1010213300210332-2210033213123312-3303330033333221-1133011330111300-1323323220031210"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2223030300223300-3022200121023011-2211322003013132-2321223030023321-3230202013113200-1100112303212200-3332131002133133-3301103300323310"></a>

<a id="canonical-2332113233211100-2001000202200101-0020110300300300-2322223210103133-3120003310001310-3222210103133032-2213000130100121-0330331011332013"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1320030213012212-3010112222003331-1220120310200012-2023331320201303-3002113221000303-2321020120131023-3023222110200113-1112033211203111"></a>

<a id="canonical-0121310311010001-3001200011221110-0113111203210311-0000031332103120-3331101103220333-1121101111201320-3133132111031313-3233223023130213"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-2023331002121112-3000022201031330-3310321313111213-3233003102133311-0002133101330122-1312130231203011-0232022011223202-3123330310213220"></a>

<a id="canonical-1000311331303231-3101013211331312-2322201122202202-2302012133131211-2013100232201111-0001230031033113-3320210101322312-0202223000131133"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-1222322031310320-0330223222111001-0003213133011320-0330101021303213-3213333322213213-3301212023013002-1023022222331030-2303302223231322"></a>

<a id="canonical-3021132233131322-3232230330013030-1223110122303121-0233201321322201-0220120101221123-1311211000302013-2323312010123032-3031112030213120"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-2002302301223210-3001312211100232-3213312012321112-3300310211222020-3213223101012132-0333320111130321-3112331010222121-2110300001010001"></a>

<a id="canonical-0330200133303123-1102332101222100-2022022303230201-1011022133010111-0212303033022203-0100232000201121-0020023203121121-0032310101212222"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1031032311001031-0311320123113231-1212231333011320-3333332121313113-0302023331322230-0220001230112121-0010033131233300-1123020103222012"></a>

<a id="canonical-1201010221321003-2220113203222003-0222303031223012-3113201130300001-3122233011300312-0330120300322201-3002032030213233-3031011213313003"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-2012310320122332-0012302201101121-1032312111313220-1113220311013313-0100313221310121-1031230003022000-2200023202330130-2230131201302112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-3323203333100332-1322332222201202-1131302031000213-0110333302133323-3012100332232120-3111121023031111-2312213211030010-2223000000303123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321321221213113-1020212011101231-0120013032231021-1121101310010222-3323031123312110-1311221203213103-1011130011123130-3300302203112100"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-2010220211220301-2332312122112303-3223120000033312-2021003300002032-0012111222002132-1110301010011311-3303332021312011-0332232212131303"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1000321223320222-2000312331113232-0312211332033021-1110100001010033-1001111133001303-1212012123013001-0102132300010210-2200222302132303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-3333131020323100-0110303201002010-0333022132323301-1122333202020131-1123311321223013-0213122101233301-1101320312012102-2022233220130012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310022120033101-2200213330020203-1230301323323000-3213222332110311-0011023212220032-1123311131033311-3322232211310323-1120031310011223"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212223001330012-3221100332133323-3210032013321303-1203001200102211-3212212013121212-0313221102232330-2232002100101321-1232013222211102): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3010101223303031-2320230222331211-1100320210230202-3321110030031210-3102211011003001-2303331301312101-2022332212120033-1122212000123130): complete subsection reference.

<a id="canonical-3212223001330012-3221100332133323-3210032013321303-1203001200102211-3212212013121212-0313221102232330-2232002100101321-1232013222211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2331131303101201-2232012000203022-2313312210031022-2303030003303012-0311030012301231-1133233022303210-3232003120121210-3230303023322200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232300332131102-2332333033031312-3311011130100211-2310013023010000-3333201300121211-2233312301121030-1033213133121212-3031301133202033"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0002233003331020-0300323231333131-3321231212212202-3100023102001002-0000233201210321-1033132302213033-1102202223311220-0331110301030032"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0012122322121213-1103100320122022-1311233023211211-0220003232103333-3221310303113101-0332133220202330-1310033012113021-3201201302113033"></a>

<a id="canonical-1112232332121132-2223331002220001-2323211100321330-1312201002221112-3203220102302300-0102122211312332-0332213133322303-3221302133101133"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2301100211103023-0310131022111003-0221033310210211-1333003332110013-1022312012300102-0211103022131222-0213322202022002-3033310223313031"></a>

<a id="canonical-0201213313232330-2233212330001232-2132210111010320-0020230122202022-1200033313113312-0330202102012223-1112130231130112-0311210232323211"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3010101223303031-2320230222331211-1100320210230202-3321110030031210-3102211011003001-2303331301312101-2022332212120033-1122212000123130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-0223231001120000-2220203332332221-0012000213303032-0022023030203310-0013321022320111-0031111130002023-3033302131333202-1300233320330220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133302232123130-1321012322123310-0010110013131213-3223321123220211-3033003122212331-3132133230232030-3013233321223230-2023132030231223"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2032111223032301-0111330313010311-0220023123130323-1130110212231010-3002113111121310-1320321333011322-3321121231123202-0111302101203211"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1022211131012012-2000233202323231-3000221013011201-2102332130120101-3032121121110023-2133320331312303-0113010320120320-3213213220233021"></a>

<a id="canonical-3313223220323300-1300210330301103-3212300003231221-2320032302302232-2122202213313102-1222233310233313-3221322202221230-3231323012031020"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0321211001212221-2201311223311330-0203120023100113-3102023123032112-3132003033033032-1313231223221223-0311320211032220-1102210003202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-1211300023330002-0331112112113031-0003210101323130-1222232323310333-0110303302001103-2131331233202030-0103011113331300-1020102212322012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031020210010023-1201331221331233-3030210310323100-3110000032123022-2313130031331210-2233320022202102-2211203032032301-1102211332110031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_mtls_obj

<a id="canonical-1210331023213212-0010002012301121-0013301313330320-3012011213022320-2331322321312331-0223112022133322-0023012302222300-3312132331203213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313312000321331-2113331002000120-2232211033301212-2100110031333003-3200333133133021-2100331000212010-3231221210011212-2302300103232233"></a>

### Direct properties for `origin_pool.use_tls.use_mtls_obj`

<a id="canonical-1133012123113013-2121323023101021-1013310101222200-3220021011331210-1230130131211131-1103331000010003-0123213202001303-2222130102302232"></a>

#### `origin_pool.use_tls.use_mtls_obj.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0132321313033030-0111211332300022-3320303032002212-2322331213122001-3023213031321220-2222331103213211-2121233303332232-2223130000203133"></a>

<a id="canonical-1000303110012033-1002211001330201-2222222100313032-3112113331103231-3010321232030231-2201110023233302-2121332123333002-2003130300301330"></a>

#### `origin_pool.use_tls.use_mtls_obj.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1201331222011300-2311211121323033-2103101321101021-0331001011120133-2331332222230133-3030030211120032-0203100321020211-0200030300331221"></a>

<a id="canonical-2122231200023221-2201313030210102-2332013113020000-3230221110203010-0330100111302111-3311203231112031-2323102121312222-1011212133202311"></a>

#### `origin_pool.use_tls.use_mtls_obj.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_server_verification

<a id="canonical-2220130033132213-0011021310023330-1102010203300320-2123111000023020-3032130011203121-0102210310033012-2333032132133221-0111100323132103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Additional upstream details:

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

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310231212322113-0220122301031302-1222131010132201-1311130201302320-1203032330312020-0011330330120233-0233303031200001-1131323113120032"></a>

### Direct properties for `origin_pool.use_tls.use_server_verification`

- [trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300): complete subsection reference.

<a id="canonical-3012133003230110-3122032221312322-1201031021010123-1202223303033323-1012002131300311-2233220201012103-3232230202322233-0310300302333023"></a>

<a id="canonical-2302200312001023-1210122322331133-3303211132030003-1102202111132113-0313222013323100-0113231000212312-2000022321230331-3033003322010013"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-2303233022001300-2311022331230210-3111001002332233-3330113303203101-0213231122203020-1012332002020133-0112102000002220-1223010221311102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210212023002231-2321210002112023-2100330303113023-2331032200323100-0321120120023213-2133303233002201-3012013000312031-0323110300303013"></a>

### Direct properties for `origin_pool.use_tls.use_server_verification.trusted_ca`

<a id="canonical-0031000020110300-2132200001212233-2033132003332201-1021121313123231-3112130210331312-1310310203001301-1301011101311211-0022010203001200"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313133021321132-2223122331210320-3311110013020013-0323221300220122-2131100131021200-1131102002201022-1231212122233001-1112001333112301"></a>

<a id="canonical-2123300223303230-0322230021202002-3203031100200020-3022330102333203-0323221210101120-0021100213012310-1002101121301032-1031312120333321"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2213330020231010-2112030003303333-2313001320031132-1102331220232123-0313002213032320-2122220101103210-2131303002311300-3112303311222012"></a>

<a id="canonical-2200311131221330-1101121333333201-1023020330001332-3221323002101333-0121303011300003-1132213121032223-3320031222023333-2113102323112231"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.volterra_trusted_ca

<a id="canonical-2111132302100021-0110132102031000-1113022233110011-1303020322030321-1112330132011311-2033021022123320-2301131231223000-3322101003221033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

Additional upstream details:

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- other_settings

<a id="canonical-1222331321330011-0201120000010222-1203121020212220-2233310323010120-2321122200001333-2201233222312132-1302003321013013-1130233122001302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for other settings.

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
other_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213113032333122-3322300230013333-0100130030330002-2110331310121031-1313031103202100-3200322103212102-1212210102321000-1030103020003313"></a>

### Direct properties for `other_settings`

<a id="canonical-3001102222002010-2330013221301331-1301220323030223-3313000213011022-0330322332223111-2320332113112003-0021202120111201-2023032332032011"></a>

#### `other_settings.add_location` property

Type: `"bool"`. Optional.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

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

- [header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123): complete subsection reference.

- [logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130): complete subsection reference.

<a id="canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- other_settings.header_options

<a id="canonical-0120012201013020-1223101101033003-2320230302032330-2210121213203320-2100013023120233-3320230320003211-3333321233103120-1232221011330033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
header_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233132100212301-0112010323313011-0021013012113331-3112311323223230-1022002220033003-2311101332202020-1221231033233120-3132230200313020"></a>

### Direct properties for `other_settings.header_options`

- [request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202): complete subsection reference.

<a id="canonical-3313230232133303-2111301013112013-2230131011103332-3001113231030203-2030303002103210-3311333320200133-0310331111013021-3111023003321302"></a>

<a id="canonical-3302322230001203-2021301122103221-0021113033000021-2213230102021223-1232002303103131-1202310201121232-1021113002323103-3233130003230322"></a>

#### `other_settings.header_options.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322): complete subsection reference.

<a id="canonical-2300220121021030-1030311131100310-3213213031130331-1220300302213232-1120310220232110-2100312103312302-0231131211302231-0220121030233031"></a>

<a id="canonical-1123113231020222-3012300210233322-3223232121002122-2202221022002320-2131211131110221-1110030112010110-0022202330112210-2301003322112131"></a>

#### `other_settings.header_options.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- other_settings.header_options.request_headers_to_add

<a id="canonical-2213300010232000-0110103232330030-0312122330132333-0122333201121313-1002000301131211-3210033201310312-1020111111310122-1303123112022020"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1103311032222101-1220030010122210-1332310003303101-3112020123330333-3122300223313303-0013212323212133-3002120021323321-0010333220230122"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add`

<a id="canonical-3103132120221331-1123001131211212-2032221232013200-1312100110113210-0113033311300201-3131310220102110-3301031001103200-2333212021003221"></a>

#### `other_settings.header_options.request_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

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

<a id="canonical-0232233301120101-2212210120111112-0133001332121232-3221033110033213-0322100013302322-1122212212203113-2322222233130310-0220123333033031"></a>
