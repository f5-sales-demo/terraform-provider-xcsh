---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- tls_tcp.tls_cert_params

<a id="canonical-3132223110233100-1320012202202033-1132100030231233-0033102321001200-2232203103300203-1033211111333321-0331200012002110-0312303033331120"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-0211103100111332-1011212200303220-1330321311303131-2023021003203003-1232023003122120-3101222133011113-1210020331111302-2312122102023330"></a>

### Direct properties for `tls_tcp.tls_cert_params`

- [certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1112320022302113-3133132103111312-2013212002203203-2320231003120030-1120232102313313-3012130013022023-1331300202221120-2020231321233022): complete subsection reference.

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2012003003113330-0132112313013121-1023333102003321-1133031032330010-0131232013331213-3300333230002230-2221123012220330-3330312310233012): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232): complete subsection reference.

<a id="canonical-1112320022302113-3133132103111312-2013212002203203-2320231003120030-1120232102313313-3012130013022023-1331300202221120-2020231321233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-2131031203113123-2120023231133210-2031321322203322-3020220130231101-3232323000013013-2130030221033213-2333302213132323-0211123023031222"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101131121012211-1032021210231100-1022031230023103-2213220022110300-3002002102011200-1103321331010032-1303032303030021-3021231322030202"></a>

### Direct properties for `tls_tcp.tls_cert_params.certificates`

<a id="canonical-0330213130033302-2311133212003232-1231011202220331-0201310012311002-1302323032110110-2011103103003101-1102100200021002-2331331000332311"></a>

#### `tls_tcp.tls_cert_params.certificates.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2012221313301311-3121010301222010-1303302113013130-0321122033330101-3323232033310200-3202100230201211-1230222131023332-3303222101121120"></a>

<a id="canonical-2210223223201000-1311030103102321-2232303121023231-1202012132020332-0221330312030121-3231310001313102-0121031303321023-0220003013122303"></a>

#### `tls_tcp.tls_cert_params.certificates.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2032333032103301-1310203201210310-2213221300031023-1333112212202311-0011122300022032-1113001303130332-0232330033123311-3123100102033320"></a>

<a id="canonical-3121131323021321-2030030233000310-1201223210330311-3233001132000201-0303312213213222-2130002023310030-1302323030113113-1032333000332230"></a>

#### `tls_tcp.tls_cert_params.certificates.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2012003003113330-0132112313013121-1023333102003321-1133031032330010-0131232013331213-3300333230002230-2221123012220330-3330312310233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-1131332332222222-2211033122302333-1310003133220203-1113330202203211-2202111030333333-3331211322133001-3130032120120221-3223110203103302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-1121001013333021-1103303101321122-1231003331210113-2031010221010210-1313221103013223-1111313301230202-0232031132211130-3311300223310031"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3031011122112132-2312231330130111-0030002033120101-3322332333311330-1022121012231213-3303021000312312-3221033211001000-0001022030313202"></a>

### Direct properties for `tls_tcp.tls_cert_params.tls_config`

- [custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3311333121130022-2211220203231132-2000203330332021-0112200333013101-1022113001303332-1013233232001233-1221232220000320-2121232223212100): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0231221023332112-0013111300113021-1321302003220333-0312313301310013-3102001013231122-3231021230020303-2102021300103010-1121020000012121): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3013211012113110-1100321110200123-3223300302021332-1222102203330102-3010300131111231-1000132312212100-1220213120001221-2222031113301120): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2212002321101031-2100301010312210-0132001322000130-2120332323120010-2301113321223232-2223310031112030-1111100220031111-2121312120221200): complete subsection reference.

<a id="canonical-3311333121130022-2211220203231132-2000203330332021-0112200333013101-1022113001303332-1013233232001233-1221232220000320-2121232223212100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-1200022311111031-1221101303010133-3200312012113323-1202310320303130-2211212320213322-1310213121233320-2130002103122320-2103013221013303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1311011231302223-1331022033030200-0202022201212022-3132213011121100-3221223213120131-2121332323033000-2301031210223130-0001220131032122"></a>

### Direct properties for `tls_tcp.tls_cert_params.tls_config.custom_security`

<a id="canonical-2312011321003020-2120011003222131-2230031122022211-1022020022033303-0222131301032332-1312210331323223-3000101212122300-0203013312220020"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2031121013232310-2212011011310221-1020112310311120-3300333311021221-2333120332310220-3111122010013033-2321113300311023-0323233031102210"></a>

<a id="canonical-1130130231332111-0101221022011320-3130233210112213-0211111011101020-0233122332002230-2320203230023022-1201111230020332-2233020321200300"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

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

<a id="canonical-1013030331321021-2120321112101012-3022010120221302-0111132203322312-2311132031001112-3313111301120021-2033233033101010-2300022301323211"></a>

<a id="canonical-1201001111003301-3010123122302330-1300130000122231-1020100313012120-0210130211000220-2323330102132010-1210320332200133-1112310210013013"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

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

<a id="canonical-0231221023332112-0013111300113021-1321302003220333-0312313301310013-3102001013231122-3231021230020303-2102021300103010-1121020000012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-3000210203333221-0332202020313002-0322130023111302-2130322221231110-2231122122020131-3232223200011221-1130330030223203-2213333332001121"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013211012113110-1100321110200123-3223300302021332-1222102203330102-3010300131111231-1000132312212100-1220213120001221-2222031113301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-0030013133102302-3111331100103230-0212131123223001-2130232010030120-1320021111111002-3020012013103223-0313320310121232-2011202330222023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212002321101031-2100301010312210-0132001322000130-2120332323120010-2301113321223232-2223310031112030-1111100220031111-2121312120221200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-1012320000010113-2333030302211220-2313132232130223-2130011220132223-2031110103133103-2201032313000012-1122201103311331-3221213012232132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-3012321301330313-1130000121202210-3132002213121330-2101110133000102-2221110132220130-0130220212232021-2133131310131330-1221300011113332"></a>

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

<a id="canonical-3330012313003130-0323300023003131-0102223022220002-3100323122321000-3130203231023032-3321232312332300-0201132113033123-2213032312202022"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls`

<a id="canonical-3203000003301102-3110322001303011-0132221021312331-3013220331211213-1202220320222030-1010111030021230-1110231202231202-0100301210001233"></a>

#### `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2033220303231300-3223100331113101-3120220103323101-0203331202123203-0210133213200223-3102003112030231-1122320300211022-1301202022323010): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0313210121333221-3230321300020031-3302130032010203-2020100103212322-1133031103031310-0323223113032022-3200333330223322-2030313311321323): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1312123122202323-3133021133013122-2323012131313001-2230121233132130-2000110130212221-0230231002002133-2300301312012111-0010201102213322): complete subsection reference.

<a id="canonical-1232023003032313-1211121000333331-1101023310312120-2222112113233132-1030000122122211-0002200212103033-3103001212020322-0330231220022331"></a>

<a id="canonical-0201311320211333-0320313010011303-1003212122332300-1231323330320200-2200320221131110-2111220212330010-3202313303323133-2333130010220013"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1232001321113131-3200223213223231-0031031230030102-0123303211211132-1213110111320003-1210222123313201-3233032203112130-0231311303221300): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3232023203013122-3310021000322212-3310003321033000-2002221122201102-0000102020010230-1331202232313223-0012110121010012-1120301201001210): complete subsection reference.

<a id="canonical-2033220303231300-3223100331113101-3120220103323101-0203331202123203-0210133213200223-3102003112030231-1122320300211022-1301202022323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-1021021310222211-1220302113221211-0213213022332030-2020103012212202-0132311010011202-3003023002121221-3032010100221101-3230133203310033"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0333200323012121-0121211022112232-2232302311120111-3322302230131001-1221010022010223-0022210021321312-2213123300031232-1223222010222000"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.crl`

<a id="canonical-0013221000300021-1310113131223003-0112323010100101-2200312132322003-2033331121101121-0220013131002030-1101310130133101-0110011112320012"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0113021201200122-2133120112123120-2111223032133301-3231032221031333-1102132313013120-1013300102301030-1233013331000322-2031121002111330"></a>

<a id="canonical-2033000221200111-2102220113221211-1110303133211202-0300222312103212-1032102313231302-2320101033103021-3133013003220100-0331013200320013"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3301003122110202-2120301021221302-2232110131313300-2223321310230010-3113330233022230-3021210113132210-2123000333232012-3022031102230113"></a>

<a id="canonical-3323303211022032-1101230333030033-1310302103321332-3111301023213211-3303333011103020-0110102322211101-3201332130121021-1103321312223321"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0313210121333221-3230321300020031-3302130032010203-2020100103212322-1133031103031310-0323223113032022-3200333330223322-2030313311321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-1313221332320211-1011200022102330-0322213303201222-0023330132023302-0212301222303112-3101213100332003-1001222201313113-2312020232003130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312123122202323-3133021133013122-2323012131313001-2230121233132130-2000110130212221-0230231002002133-2300301312012111-0010201102213322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-3100312211000133-0002100201332030-0312320012202300-1320103211122023-3122332010003131-2332321030001222-1122133120330133-2313011031120223"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2213333213110030-3010333022213313-3132203203010002-0001230023033112-2310102123033230-0002101332122012-1200103331330000-2330031221010310"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3031030211121023-0330302030001022-3133223320233311-0122020313300322-2031210230323313-2211113133102110-0202223023133120-1321330133012103"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2212121103233230-3211133303132030-3300213010321023-2021211120120203-2111120220001121-2030320203320010-1321330002233010-1023011322120130"></a>

<a id="canonical-0110322013223103-3300321231232303-3303023323010031-3211102201133102-0020232100332331-2113330210003003-2210002301111222-2013010232322033"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2210122211323312-3012020130321013-0233130102203210-2030013133202311-2321230221121310-1321012003132212-2210132002133330-1300121323132302"></a>

<a id="canonical-3201103112303302-0213312103312021-2312130113200022-0030211111110300-0003030020033102-1031332020230003-2202000030001022-3100103123101101"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1232001321113131-3200223213223231-0031031230030102-0123303211211132-1213110111320003-1210222123313201-3233032203112130-0231311303221300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2012330212030210-1130211020301313-1030130013333121-0003003012210021-0023313032110211-3112133000211232-1030103033122100-2031121123003333"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232023203013122-3310021000322212-3310003321033000-2002221122201102-0000102020010230-1331202232313223-0012110121010012-1120301201001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3113333131301233-1012113033202310-1200303321122130-0311033223320133-0300131110102203-0133233000013023-0112032032022021-3132023202202111"></a>

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

<a id="canonical-0232233303011313-0322123132010213-0231302130203300-3131020203330203-0332010101001123-3221313310212301-0030211332303320-1100300331130212"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0223223130003202-3320131021110302-2323220000220122-1013112221003111-1100223112320022-1322233300031301-1132133221302111-3002032130322232"></a>

#### `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- tls_tcp.tls_parameters

<a id="canonical-2121021311101012-2230010312001123-2312121013311221-3322320113212022-1023322132223233-0233320112310313-2000233212123201-3101323233030001"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3011012320331123-2211132222210220-2320200331030011-3222323212112303-0200231220320133-3201032220311033-2313010020323132-3102212311230011"></a>

### Direct properties for `tls_tcp.tls_parameters`

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0233013330312022-1322221102311201-1202300132103012-0133230011223300-2131202330320131-1210323030103011-2000210010321102-1223113020012032): complete subsection reference.

- [tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333): complete subsection reference.

<a id="canonical-0233013330312022-1322221102311201-1202300132103012-0133230011223300-2131202330320131-1210323030103011-2000210010321102-1223113020012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-3001133123312002-1000320221230212-0222121000310303-1200233122223222-2230332032132333-3313210001031231-1313031322200131-1021111110011000"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-1012102230213233-3212321103023301-1031112200203333-0003130113203002-0113233211132130-2100101021100331-2303203302000310-2130011002231231"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3222312330033003-3131213030102020-2132232122212112-3121313222200001-3103233312323110-0332013132133100-1330201220112322-0331122331133131"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates`

<a id="canonical-1100213020121312-1102013313211022-2021013123233101-1321231012321020-0320132023320322-1303211332030232-2331203210331133-0332202232232121"></a>

#### `tls_tcp.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1312320021030200-1023033111010123-2231303213022200-2333113221031111-3331010212013101-1221030110303211-1311330102003211-1303112122010003): complete subsection reference.

<a id="canonical-2230333001233120-2021110211021230-1132332320103130-1112110302310001-2322212112121030-3310320120310020-1331013233132303-2210320312222033"></a>

<a id="canonical-3003033121322100-0122313312221331-2103202110020033-1132113203302332-2130320021130233-3011033223332131-1123221303322331-1132302200312121"></a>

#### `tls_tcp.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023): complete subsection reference.

- [private_key](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102): complete subsection reference.

- [use_system_defaults](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3013202212033200-2301011303032213-1113001000221013-2321021112001011-2333001300303322-1121311110012323-1330012011103030-0230031010332022): complete subsection reference.

<a id="canonical-1312320021030200-1023033111010123-2231303213022200-2333113221031111-3331010212013101-1221030110303211-1311330102003211-1303112122010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3333312302220231-2222020031203330-1203200113121021-2201210323030201-3033223233113213-2302123310131210-1213103311101220-1000102233110122"></a>

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

<a id="canonical-3333223211123312-0001232000301323-2232001213333113-2102313300201003-1031312032133031-1002111220323132-1331101200130022-2233311302132300"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-0201303311112033-2123130100113320-3331110213002110-1030033210221302-0331221313203021-1322222013121220-3312222013132233-1022000102330003"></a>

#### `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1011133323230033-0011110010130221-1323003103012012-0103311200301030-1130220220222033-2103111221122030-1320001303310010-2021201132022322"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-1101331003131213-1231312032211322-2233011011231000-1002211131023110-1022021311003332-0223300202331220-3102031301321103-2220022111213132"></a>

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

<a id="canonical-2311121003122132-2330202003301322-0332021231020200-3013003200131133-3203010212330231-3231300312300221-2023211220132032-2032001023122333"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1313303320012120-1133213001033103-0320333003301302-0200033133331101-3312030010123122-1001133223300000-2213213322101003-1000112231200000): complete subsection reference.

- [clear_secret_info](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0210110131100012-3202013230230003-2233100231111021-3233201333010332-1312220310320112-1003101221321103-0102213130030101-0032322120313211): complete subsection reference.

<a id="canonical-1313303320012120-1133213001033103-0320333003301302-0200033133331101-3312030010123122-1001133223300000-2213213322101003-1000112231200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3320320232322331-0001011010110300-2033233300111122-0013003310303200-2323210133200202-3030100013323002-0021300202230233-1122010002233112"></a>

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

<a id="canonical-3222020203212300-0202102322021222-2000122230330332-3102233312031200-2220322121131333-2202332121112331-0113032011103231-0102211032023123"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3033010312023320-3132320200102133-2303231033212110-0012020102201103-0120213302220303-2211210332001202-1301232302221232-3221221323301100"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0133330130022222-0021302211130321-3023230232220112-0110000232322331-2201321323012013-3103103132211023-1021332230220130-0310013333331303"></a>

<a id="canonical-2231211120301231-3233333210210223-0130233100100032-2033300023103022-3233110120333232-3031222301030300-2002230030032322-1222332201001311"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2023322332202320-2213301030232110-0302133102330320-3230222011123313-0120333221212013-2003103222201320-1320131031202012-1031323022221231"></a>

<a id="canonical-3303322310010132-2330333121313010-1122010231012033-2113133200213132-3212212030010102-0211333102033321-2101232310012011-2100113131321233"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0210110131100012-3202013230230003-2233100231111021-3233201333010332-1312220310320112-1003101221321103-0102213130030101-0032322120313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1312210322322010-2311302023001232-0032233233302100-1121300131032223-0232313303000000-0132231331200033-2310111021202313-3011303222210213"></a>

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

<a id="canonical-1010002003302121-1220221213321123-0202102030201302-3120013213120022-3313132223023131-1320003032233023-0020122322230130-2310003310313001"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3003112000123322-1131112122120313-1113332232201322-2032211033000333-3313120132010033-3130233333012003-1113033130331320-1011201223003101"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1023030133133323-2302200000330013-3220112010000021-1221303222033010-2110120301010220-0131220221211213-0313101120011321-0003322102212211"></a>

<a id="canonical-0002310232222320-2020321033312131-0123100033033300-1310121331013001-3230032320313202-3122110323101213-2111303213001022-0010110013321103"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3013202212033200-2301011303032213-1113001000221013-2321021112001011-2333001300303322-1121311110012323-1330012011103030-0230031010332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0133232121100203-3112131312222332-1210211201023203-0213031301113222-1100110303101231-1002321023220200-2130320023203200-2102003202313001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-2221303331133203-1130203323222012-3330132321032001-2221230001200013-3211320212313313-1330022310200131-2012220011310103-1123313133101101"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3330131231023313-2100231003001221-1310230332203001-3211131302002120-3000000013133010-1131222221132333-2320313213313321-1211222123033000"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_config`

- [custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2313330220103123-2201220320310123-2011113302301211-1321222013332010-1231123323023011-3003130012232220-1002222233203333-1011000312201213): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1202210311313111-3213122012121030-0112132213012202-0212010323332023-0110123212333122-0000212302331210-3133120103320012-3022311123003121): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3320233103032110-0222202323122111-3120211200010322-0120203132021002-3031133233021303-2320223210003312-3030020302103031-2011202203323302): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1331020023012022-3032213132000232-0121020020323021-3101333010132310-2120022302203221-2112321202330321-3221300002133332-1313023302320030): complete subsection reference.

<a id="canonical-2313330220103123-2201220320310123-2011113302301211-1321222013332010-1231123323023011-3003130012232220-1002222233203333-1011000312201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-3013101131313113-0201013023331222-3233331012300022-1301301333232231-3312223021100102-0312131001221132-2103123100122010-0102133012331132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0213320132210202-1223021321222203-2222310220001301-3211232232330330-3001231323133332-2300022021331332-2122013110021230-1312111031302330"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_config.custom_security`

<a id="canonical-3100112232003232-2300221211320311-2021232303000202-0300221100013232-0311030213023021-2013322320033113-3302320120122101-1310220133233300"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0300113000203223-2220033311213120-3310333103003013-0301032103312203-0011301303010303-2001211122103231-0111103021111010-2210022200322330"></a>

<a id="canonical-1222123001232102-0123101001213331-1132130210120010-2112220121323123-3122012122001310-0103322232001032-0211201113002120-3231113232330212"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

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

<a id="canonical-3331100231131101-2031232201123000-1032022131120020-2331330012113233-2331102112333131-2202232211011330-1202011102333313-1110203132310202"></a>

<a id="canonical-1302110131101121-2222301310001313-0003332230323331-3113102213022003-3302313000231000-3222020100123212-3311011130011030-1023133310110213"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

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

<a id="canonical-1202210311313111-3213122012121030-0112132213012202-0212010323332023-0110123212333122-0000212302331210-3133120103320012-3022311123003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-1332022222321223-2110230113333020-1021123232000223-3223023231122321-0301011020323100-3232031023123112-2223012332033230-2012203201302001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320233103032110-0222202323122111-3120211200010322-0120203132021002-3031133233021303-2320223210003312-3030020302103031-2011202203323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-1233020333030001-0023231023111002-2320211210113132-0200321030010330-0222301121001232-2031100331121301-0100303211321012-3311212012121130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331020023012022-3032213132000232-0121020020323021-3101333010132310-2120022302203221-2112321202330321-3221300002133332-1313023302320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-3001103131311223-3010020313031202-3200003032201111-0002112323322103-3123310233032000-3131111123113132-3102113330310000-3231122221311101"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-3223222113230022-2123203033302010-3201002232021200-2112303222332011-1013320230313022-2100203313201023-0320013033033023-2203133121123123"></a>

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

<a id="canonical-3303311113012301-0232130033301002-0013021321000203-2221030213002320-3003310201331100-3031121300213013-3131233330120113-1010020331121330"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls`

<a id="canonical-0310032321121022-1131300132301122-0101302232113203-1330002133211013-1022010011303221-1310002003321301-2013330120233011-2331230323113310"></a>

#### `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3113103210011021-0023200113021303-0030130323322102-0132102302111102-3320331113032113-1222212231001332-2320122112023030-0110011013320110): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0231000011022020-0131323132201333-0233202001110131-2112133122102301-2023333032022121-3100321310331130-1223300310301230-3133100232113130): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3203103003231323-0010301020132033-1312333201232130-3313333321111313-3203333013300103-2122333003332312-1112012110001020-2223201310331302): complete subsection reference.

<a id="canonical-3011301223002313-1302132132023313-0330120020322202-3032302202111112-0010320111013303-3200332330210301-2302112231221303-2112120130322120"></a>

<a id="canonical-1011213212223123-0102212011200300-2211021012122212-1101213331102123-2303003320001132-1303132211121112-3023102203230112-1300133222100223"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1310312103233030-2113330332320313-0130130222220301-2103021121021111-0123113011232102-1130221201123030-3232312231303131-2332100210012133): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0102203131003320-1313021231132023-3111101010120312-0213011120320200-0032102200003313-0133301101301332-1012333131320011-0012003103322233): complete subsection reference.

<a id="canonical-3113103210011021-0023200113021303-0030130323322102-0132102302111102-3320331113032113-1222212231001332-2320122112023030-0110011013320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-0100200111001303-3031030302320223-0210132001100020-0232033032122220-2023302130333033-3322212220301321-2230201112022322-1103221030212303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2033013301212213-2331223123032203-1320020123112203-0023233023233101-1333020111021200-2322133003200200-3221200203021311-1031012022123002"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.crl`

<a id="canonical-1112302322302231-0333230320023333-1313222210010222-3020100210320032-0212133121233013-1202111100012021-2121001322020311-0033131210003102"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0101102323230331-1033111200123212-0313133311121003-3302110203330022-2312030020033123-2013312213303022-1332232331231313-0231323130312312"></a>

<a id="canonical-2103232102200100-0221123302212213-2133000300030003-2200201213130322-3331313111223023-3322311132133313-2111200023011233-0012032121223221"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3121333021032331-0112221131203000-2230102322330110-1212202010203110-3312012231213313-2033322210322331-3323302330310322-1301003203021132"></a>

<a id="canonical-0230120330013222-1002300320201312-0300002110111323-0221120013332002-0033302003102311-0310203002030012-2302033013301302-0331232011333011"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0231000011022020-0131323132201333-0233202001110131-2112133122102301-2023333032022121-3100321310331130-1223300310301230-3133100232113130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-0332112123301220-2310300121111102-2203101012332220-2313031111331131-0001313301002011-1101331312010313-2130222002312123-1233012313332223"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203103003231323-0010301020132033-1312333201232130-3313333321111313-3203333013300103-2122333003332312-1112012110001020-2223201310331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1310022002220230-3301211110111003-0120122110301201-2213301110010223-1231112203210021-3301302011233230-3001200111102203-2023313200231030"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2131333012022210-1332031133013313-0010123010201002-0133111030310000-1302303102223211-0132011332133031-3003212210310213-3231211133132211"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-3001302011113030-0232101312023210-1101333020331021-0000310222332101-1030013002322101-3333301122301213-1011203202223312-1101221023101233"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3012232013130300-1211110120202010-1033320132320030-1221001200130331-2212033003210322-0333131030330203-3003113322003101-1301032022120211"></a>

<a id="canonical-2212013100210222-1033012320322332-2131323301302021-2111231232022121-0032213311033320-0331030321331122-2320332000130212-1222312200021233"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3102121312210012-3230223000230330-0002200132121220-3303012000200202-2231302212012211-3202203102303302-1112112331221013-2332100132203313"></a>

<a id="canonical-1301113201322301-0110111122333201-2301003120221020-0002233011000303-2332003200012131-1131221133331200-1013031231133312-0332210133030121"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1310312103233030-2113330332320313-0130130222220301-2103021121021111-0123113011232102-1130221201123030-3232312231303131-2332100210012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3120300001022020-3010321310302003-1230131212331113-2120233033020121-3202321332133232-0122331131222200-3313301220303133-2033330230300020"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102203131003320-1313021231132023-3111101010120312-0213011120320200-0032102200003313-0133301101301332-1012333131320011-0012003103322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1020130210233233-3331300223301103-0110222032000010-1030001032212101-3020221332323101-1230201030332102-1112303223213101-2100020221011223"></a>

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

<a id="canonical-2111312231323211-2032123133321333-1211222300200112-2320010003010310-1121022203321331-3112210112321103-2120110022313113-1003332223121003"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-2203003223223132-0103012231300230-2103111213131010-0223322310201003-0300102321200222-1122132202030210-0230003011110302-1112111023231330"></a>

#### `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tls_tcp_auto_cert

<a id="canonical-3122010013013333-0010120031233203-2230121020320021-2303221102121132-3030013013032022-1321333312130213-1313113220130030-1002023233330000"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3001221002101232-2110033203133201-1001101101021320-0332030212221112-0031202100200130-0321222010330123-0203020110223012-0313112331300222"></a>

### Direct properties for `tls_tcp_auto_cert`

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0020031100213200-1201323322020133-1102331111203000-3232323102121231-3320311130130232-1110312330122333-0221033220013000-0203122030013300): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122): complete subsection reference.

<a id="canonical-0020031100213200-1201323322020133-1102331111203000-3232323102121231-3320311130130232-1110312330122333-0221033220013000-0203122030013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-1312323322322321-1230033311202000-0021330230212001-2121300030011302-1310011212312321-2321022223213330-2022210331101330-2232332123301233"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- tls_tcp_auto_cert.tls_config

<a id="canonical-2100223333301213-2002332220230023-1111133020100000-1303221120200120-2202022230303211-1102203132031311-1102301231123233-0210203300120303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1110320230121221-3212210332303333-2110311133133323-0032301103032232-2301312130102321-3232012313200212-0023110123213121-2112311123102220"></a>

### Direct properties for `tls_tcp_auto_cert.tls_config`

- [custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120310133012112-0130123232123221-1322002121223002-2320133321120133-3001330312233013-1201132210000310-2010200313230300-3333102030011100): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1011222301233103-0101132330111200-3002231332032030-2203231131213312-1211210211020231-3200103321020100-1202011022023230-0032320221323302): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3223122001220111-2210032122330221-0020001102233131-3130011002220131-3311122323002312-1112210212322202-2312211202121331-3013210100112313): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110130132333313-3000221311101333-1312011230032310-0120312223201000-1101313331312023-2133033110100221-2320313021012030-1212212201231323): complete subsection reference.

<a id="canonical-3120310133012112-0130123232123221-1322002121223002-2320133321120133-3001330312233013-1201132210000310-2010200313230300-3333102030011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-3113213220330012-3210103010123110-2322302223233001-1132212132131022-3131220221333211-1113012020101213-1202012113211110-0121032132032023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1122010323003120-0130323320223002-0223021310200230-1112103123311013-0202201010320203-1201033232301121-3121101101230021-0310331002020230"></a>

### Direct properties for `tls_tcp_auto_cert.tls_config.custom_security`

<a id="canonical-0023332303002130-3332030023122220-1331331200202023-2321231133130303-1201010012203210-3332123211201113-1333112202311322-1003113331202113"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1310130000230122-0200213222323231-1211023301310201-1022132332202000-0020112121321330-1012110323020233-0121212220030322-0000310203212233"></a>

<a id="canonical-3302021011121223-0331220130322100-2011330331013001-1013310133331001-2102330330310332-1003113220330013-2103130031223112-0110002101230000"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

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

<a id="canonical-2133221010110323-1020302201323201-3310131311130010-0230001321100232-0111101000131201-3221131320303221-1012000032112322-0303123110202021"></a>

<a id="canonical-3011321010020200-0121220031320021-3010221113010130-3013300123101023-3203323010033310-0031210101033220-0323212131333310-0333021013222210"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

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

<a id="canonical-1011222301233103-0101132330111200-3002231332032030-2203231131213312-1211210211020231-3200103321020100-1202011022023230-0032320221323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-1231123211121221-3221333310333310-2312011201121020-2110201113100321-1300001133302122-1312233121120032-2301330030010010-3000123120323211"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223122001220111-2210032122330221-0020001102233131-3130011002220131-3311122323002312-1112210212322202-2312211202121331-3013210100112313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-0111030230322333-3323310102022203-3323320111212030-3132222030022222-2333002112200212-3232012221322001-1113202311300031-2333231131011010"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110130132333313-3000221311101333-1312011230032310-0120312223201000-1101313331312023-2133033110100221-2320313021012030-1212212201231323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-1303003213322311-0201122021101221-1302030000202223-1013001323232310-2300223002320112-1030311322130322-0320202103323000-0132020320332123"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- tls_tcp_auto_cert.use_mtls

<a id="canonical-2000030223230202-2220210002210222-3002220323331312-3303332212110220-2032210011113200-0101313001313133-0222032001300033-0012223200133211"></a>

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

<a id="canonical-0123303022132121-3223111020100103-0003110020011302-2111302220131110-1103320201133011-3322322223202312-0222021131122131-2232210102102113"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls`

<a id="canonical-1310222003001333-0201010230220131-3213212222221213-0232301131302131-2213002223313320-3130200000013213-0222123101011210-1102010112200110"></a>

#### `tls_tcp_auto_cert.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0010333131321101-3231033121330022-0112020133320232-2110003013021123-0231013102001233-2321330023120000-2303303331230000-2332321222133221): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3133312230033302-0132212002130223-3222221103032001-0301212130000132-2200221310303233-2101033213012102-1213300302100120-0120031303133322): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0330323000330322-1300012003301320-1203000313012021-2203010122221221-1220002230210202-3322233303122111-2023202303003331-3223022223022002): complete subsection reference.

<a id="canonical-3213133123022310-2300013011102123-1103210132220113-3212121330103200-2321021231322121-2323011101302212-0231103102213012-1023210111133332"></a>

<a id="canonical-2113232312122103-2212102311002221-0231131312202212-0103223330201013-1023232020232101-1312011002301123-3202333333310120-1122203112011312"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120203032230201-1103300312312203-0023001222012311-3233112322211202-2011132033013001-0220212213223003-1220202202030012-3103212130313203): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1313300112223011-3201011300103321-3000222000110100-2200211103132331-0202331320000020-0023221203123333-3100213133300132-0220111231301333): complete subsection reference.

<a id="canonical-0010333131321101-3231033121330022-0112020133320232-2110003013021123-0231013102001233-2321330023120000-2303303331230000-2332321222133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-2023121130131103-1130230220332001-1103011133232103-3102131201233302-3312020120132220-0000322010023032-2121311233322131-2231233333002110"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2302232233111222-3320302002211210-1010022030010003-0001203213302020-3330310221023103-0110201030233223-1011221120103011-2000313233232203"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.crl`

<a id="canonical-1121101213322130-1221310330020213-2101321311010310-3200322332131121-3311132212223030-0031122102303322-2222331010330031-1120332330003221"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2033021222131122-1230303230120112-3202202133121122-2330102122020101-1231232110303210-0030000310311322-1101103122110011-2201000011321122"></a>

<a id="canonical-0111233130033222-2223213213202311-2333002303100200-2112320032000110-3212031300222301-3221131130102012-1332121013322031-0031112012021110"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1020322103321221-3220301113003332-0333032132321002-3132112101201021-3312101231121322-1110002012321323-3030110131323203-3233131331032123"></a>

<a id="canonical-0302322320001221-0021203033022022-2032300203212113-0022332220200132-2131010212313111-1231112120202111-0232202103022300-1320122323220101"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3133312230033302-0132212002130223-3222221103032001-0301212130000132-2200221310303233-2101033213012102-1213300302100120-0120031303133322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-2310311132122032-0233210101310000-0222220321230000-3113133333012201-0102321222031030-0000100020332131-0323131101032031-2331003112313103"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330323000330322-1300012003301320-1203000313012021-2203010122221221-1220002230210202-3322233303122111-2023202303003331-3223022223022002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-0212132103311331-2230222302111013-3331301031021130-0230332012032302-0130203300223022-3220230331220113-3211323102100202-2002312331003211"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3103121013120133-2321101013302230-3012211102132332-1331232203033323-3111311300333333-3021223131213231-1201022333322211-0011320211023320"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.trusted_ca`

<a id="canonical-2302212322000212-3323020331121211-2230020322321131-2133000212233221-3133101011132222-3113021123013021-2332012010300321-0100121033133033"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2022003210030130-0201311121121312-1030310203231031-1332310201211221-3213131001311212-1231002213110031-1013330311221010-0103033010032120"></a>

<a id="canonical-0232332302020103-3030002100012133-0101003000200111-0310310111113302-0032330220110113-2032110231203331-0223130001212221-3021021000002001"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1313330103302210-0333112133213022-1220221203012210-3311311021330130-3233301101022123-3133211211001111-1010211323300311-3001102323013123"></a>

<a id="canonical-3020002331132032-0213111303133102-2002203322211033-1222113310000333-0021023113020332-2102130030010120-0222201012330230-1030033301002231"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3120203032230201-1103300312312203-0023001222012311-3233112322211202-2011132033013001-0220212213223003-1220202202030012-3103212130313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3232232030103320-3130121211101303-1210132231322022-3013321130230021-0300032021203111-1332012233010021-2033122231021012-1201312103222301"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313300112223011-3201011300103321-3000222000110100-2200211103132331-0202331320000020-0023221203123333-3100213133300132-0220111231301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.xfcc_options

<a id="canonical-2213310110203313-1232313033031131-1332121213033031-0213112212010211-3101012111032012-0312211230202300-2203233212200002-3212230100213220"></a>

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

<a id="canonical-1122222310010203-2302303231320023-3112130130220303-0110200312331023-0312111012211213-3300311311231022-0130223123202303-1331313220221210"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.xfcc_options`

<a id="canonical-1033233113222132-2323230213321012-2100311332211000-0303232301222200-2303311211222023-1121333223220012-1122311023322032-1232322210203201"></a>

#### `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

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
