---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1310323011103130-0300322132203323-1233131031023013-2002130323022100-2313013311123210-1113332311020102-0132203101013133-3121120333301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0323111013033110-2031211121210030-2101311220010202-0310323211113101-1020203300200302-1003322210303021-0132312032230211-1332101023333033"></a>

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

<a id="canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-3331331200113210-2332312231210213-1100113302011121-1102000322301213-1220121123100311-2122223202100202-2233232233220110-3301322021202210"></a>

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

<a id="canonical-2032332133101301-0020320112213300-2112222003313222-1132211133311322-0200122210110202-3320120301013100-2203002210233132-3122300222320132"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](data-sources--workload--reference--group-007.md#canonical-1233211231312010-2021020331101123-0122211212012011-1132032102220323-3010013210100112-1320112222033332-2320023121313322-0323233333311030): complete subsection reference.

- [default_security](data-sources--workload--reference--group-007.md#canonical-1021030310301020-0111113212130332-0211021220012000-3230102012023112-3111100233321202-3221233303032023-2232211220113012-0303103031231110): complete subsection reference.

- [low_security](data-sources--workload--reference--group-007.md#canonical-2213232010122122-2322113000101322-2110312223212010-3020211330323212-3313103201202210-2111031313322102-3110011031122003-0103333223301200): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-007.md#canonical-2023210311000132-0203302312323303-0033103213323013-1111123223331321-1201002220013321-1130102233302231-0231023313110301-0130331123031202): complete subsection reference.

<a id="canonical-1233211231312010-2021020331101123-0122211212012011-1132032102220323-3010013210100112-1320112222033332-2320023121313322-0323233333311030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-007.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0233332331101022-3032120031023133-2001030321110001-0231131001110121-0023120311013011-3100132100321133-1021203321321111-0230012033002300"></a>

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

<a id="canonical-3101122321233203-2000300303013101-3320013033300303-1100022211213300-3112231211102023-2111123231230310-0221130211112232-1301012200221012"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-0331031132331110-3102233020223132-1033120020130131-0313011321112110-2320231330200101-2311311123121012-2230202203101022-1323000230111100"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-1321203031222102-2220333310011313-1111210211133313-0131322302131021-2031011020300130-3130030022033331-1220013133011113-3122133233301001"></a>

<a id="canonical-3210213032213303-3221112322103332-3312303031103031-2330100230122330-2200202112213013-1133120111012303-2201010213111113-2031130210311033"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-0312023223223003-0331000213110003-1020301300231232-0030212311020301-3123130303221331-0132001312320011-3201322030033233-1222012202000321"></a>

<a id="canonical-0323122310112300-3111200300231123-3203213312133330-3022000233321011-1202333200100201-3232131231032102-0111100200101201-2231101203101021"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-1021030310301020-0111113212130332-0211021220012000-3230102012023112-3111100233321202-3221233303032023-2232211220113012-0303103031231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-007.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3110332212130123-1230020231320233-1133132032123331-0212331031030132-2123302223233303-1011333120222110-0110023212103102-2112022012313232"></a>

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

<a id="canonical-2213232010122122-2322113000101322-2110312223212010-3020211330323212-3313103201202210-2111031313322102-3110011031122003-0103333223301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-007.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-3223232123030130-1011203330131131-2311300332200031-0032333121133323-3331103131210013-3000110301203222-0110221112100000-2102131012013030"></a>

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

<a id="canonical-2023210311000132-0203302312323303-0033103213323013-1111123223331321-1201002220013321-1130102233302231-0231023313110301-0130331123031202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-007.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-1202301331131320-3311110030210230-3022332221323203-3231021122010321-1311312113122121-0313301001211111-3233103112122333-0312132010130231"></a>

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

<a id="canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1000232002133033-0032022332110212-3012121333101300-2231332121203233-2213310020010300-0000322020111321-2210110213222133-0102233012020030"></a>

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

<a id="canonical-1323301221201030-0010303110200120-0022210320130010-1113221221010233-0231130332101123-3200020021001222-3002333020132220-3300222201320132"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-3330113200031031-3011210131233302-2202003032110002-3323330023022220-0102332111033321-2211203101232000-0131201211030033-1012332301133020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-007.md#canonical-1332100112202202-2011021330011322-3023133201101030-2323233132233002-3221213121123003-1120301313032121-0130131330220203-1231100000310101): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-007.md#canonical-2111030321113021-1003300000221331-3201230022120022-0131021210322123-0210333121221310-3023232312333102-2313131303210330-2333303201202222): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-007.md#canonical-2211100001003313-1000000131033111-2232310310032121-2301221130001132-1103011320133122-1310020231221213-2101020133103200-3312120010113331): complete subsection reference.

<a id="canonical-0102131113130113-3311111100212022-3002212133021311-2021100331233110-3303331222302332-2221102212103002-3321003221313123-2023202100210332"></a>

<a id="canonical-2220022133122023-1311121101130301-2121103021313132-0003112012001223-1201320130313122-2321003210220331-3222131120003010-3303222132233223"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--workload--reference--group-007.md#canonical-0023021220022003-3131020323102102-3112121300133102-1023332113131301-0101102132032032-3112321203012311-3112312013132000-1221330032222312): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-007.md#canonical-2332032203103031-1210222110013121-1232330112000002-1213101303310103-1233133133010302-2131132110203032-2010311300113111-1113002221101333): complete subsection reference.

<a id="canonical-1332100112202202-2011021330011322-3023133201101030-2323233132233002-3221213121123003-1120301313032121-0130131330220203-1231100000310101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-007.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-1312132000300310-1131312223322312-2111032233220232-2232103212322213-2230230121301012-2212210021112311-3101031033120221-1000311102131112"></a>

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

<a id="canonical-1222112033310022-1122030301230330-0003033103100000-1203212111012333-0101330003012212-3323202311320133-0011102023210212-2013300300201202"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-3132032210111301-2023023322021110-1211331200323032-0100120233010330-0201131313201012-3201213011303200-1302003121303222-1033313111212333"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-0132111212110330-2332231320130013-0130111223120103-3213233020303223-0120100122313220-1022332230000333-3003332211321203-2201013230320223"></a>

<a id="canonical-0323020210310302-2232202332032212-0300220002222230-1030231211201233-2000012323001311-0332322003201123-2321222231111322-0001120222003020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-1231032003322121-1223211003102111-1323011313222111-3233102320331121-0111331001112300-3303320001132012-0220200021233311-2130331200103110"></a>

<a id="canonical-1123302321103222-2202222232001001-3320302020302030-2010001233020231-1223213110301123-2223313123223230-2013002112232031-2123223130311300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-2111030321113021-1003300000221331-3201230022120022-0131021210322123-0210333121221310-3023232312333102-2313131303210330-2333303201202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-007.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1100323100133312-3220010023001113-3302033313021320-3332200220221030-3332321111201123-3212223022223010-2113313332001320-2330020031230312"></a>

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

<a id="canonical-2211100001003313-1000000131033111-2232310310032121-2301221130001132-1103011320133122-1310020231221213-2101020133103200-3312120010113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-007.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1012122003023302-1203000110001031-3122301211110322-1020203010113311-0222201232001303-3230110131221011-1123222113110121-2023323313031210"></a>

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

<a id="canonical-0331330020331110-0133210021310200-1000130223123321-1221010033211202-3203130330202300-1021313120130322-0033211332230012-0012312330000322"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1200003133310013-1110220213301201-0101131120111232-0020302121222130-1331001300022131-3313302130133323-0032130101022001-2013112102101011"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-1130132311322301-1011222303231313-0211223300103201-1212310133121022-1310022232111003-1212030201131030-2333000312313003-2130301230020011"></a>

<a id="canonical-1302323232332001-2002130131222100-2103000212120322-1030001131231301-2112132132223313-2031023100321120-0202221211032232-0201133122001133"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1123120133110112-1223322330203331-1002122022131010-3312311132110012-2322031233220010-0111322303102000-3120211221223331-0132311120023300"></a>

<a id="canonical-2222022331331023-0233003302033333-3300300012211321-0301202122022002-1331002131212211-3322021010211312-2233033111002100-0023001323301031"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-0023021220022003-3131020323102102-3112121300133102-1023332113131301-0101102132032032-3112321203012311-3112312013132000-1221330032222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-007.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1212103200112221-3020132301202020-0102333111013311-2330221003112022-0202300130213010-3010103311001211-0011202312311111-0302301001223003"></a>

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

<a id="canonical-2332032203103031-1210222110013121-1232330112000002-1213101303310103-1233133133010302-2131132110203032-2010311300113111-1113002221101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-006.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-007.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-0212220210020302-1330202132101020-0331233011311200-1031000021333330-2220132130130100-1032303001213311-3133032320021101-1132111203212130"></a>

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

<a id="canonical-0010223332310133-2133021300132201-1301102233032213-3330013001113010-1013222020200121-3202003032221303-3111113223103300-1212201013002102"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-2021120030001003-3112010130121203-0312032133331103-2022331031330320-0003130330302111-0220102313230203-0121113123210322-0232213002101212"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-0202020012201012-2002032101110031-2212023300112033-2330231221200330-3031321223223123-2203303032111012-3203032210333223-3133121102000333"></a>

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

<a id="canonical-3330011012022230-0230102102000123-2131101132000022-2313101101331020-2312210330020301-3012032203103202-1101330322312232-3021312120022301"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert`

<a id="canonical-3121023230010300-0102012210121031-0232001221321133-3021222002020220-2321130222211232-1003012131333131-1120110211011020-2120232312011303"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.add_hsts` property

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

<a id="canonical-0131231133130102-0222323300032000-1022233013023223-2112123000130330-0111330211121013-2233130012220032-3120131332321322-0121120023131323"></a>

<a id="canonical-1102101123321100-0213101222301032-0302310100031203-0002331231131011-3220232031211310-3023113322231002-3323202332320121-0031033230032333"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.append_server_name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103): complete subsection reference.

<a id="canonical-2011231330120010-3321200003121110-3121110230233232-0133103023010112-0302112132023120-2212302000013012-3222331113333230-1113101222100103"></a>

<a id="canonical-3103111000311001-2123100122110201-1031202311021322-1220202000211210-3310020232300023-2310003222321322-0113032232110210-3003312100312112"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-007.md#canonical-0032013200101331-3121101033132032-0201111023102312-3310122312023202-1011012000331210-3022033220220302-0200212023113213-0311301132011033): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-2303033023003313-0320312301010320-1232130123120130-1031233120230310-3121203332301102-3100303300110230-2031223030231013-2131312113012223): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2202303122103033-0101112331121031-2302302322233103-1211312111211132-3001231303232133-2132213010032312-3012203131322300-1130300102322033): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2100001101020213-1220100210103220-1202213122130010-3121202002331001-1122131123302303-1211021303321321-1020132302133130-1012012113002110): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133): complete subsection reference.

<a id="canonical-0132022311223331-1310331213233231-1220213000013313-1233231110123300-3302023002211130-3221122120331012-0110220330122113-0023110331213323"></a>

<a id="canonical-2101113203013002-3031333201203202-2323233303311321-3311323133131320-2033332003212001-1333230300232133-1302232211022103-2301121302030023"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_redirect` property

Type: `"bool"`. Computed.

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

- [no_mtls](data-sources--workload--reference--group-007.md#canonical-0231223001232221-2201332131123020-1313300302020312-0221201321000320-0033233210023223-2110213013103111-0102220330311200-2331112323002321): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-3031130303320312-3011210220300230-1112101210130120-3233211133223031-3232211103310222-3322323003212131-0300113203000232-1001321332302222): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-007.md#canonical-2213213201020321-0303321203211020-2202012103010022-2302121320212312-1023313101332232-3230002111032132-2311201002200103-3213110102330302): complete subsection reference.

<a id="canonical-3003311001113331-2012032101133002-0102020123011321-2022101023023322-1023221232111333-2210202033020200-1222333022103211-2232130100012113"></a>

<a id="canonical-0131312313230013-1023200233023100-1303233211201222-2120000313322110-2303020220203032-0233022031300012-0100011110020111-2322011311322110"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3331110120222322-0210303322322302-1331132213031331-1211213210103330-0113000021331103-1102221021100110-2231120333320010-1222222330021303"></a>

<a id="canonical-1320321221210011-2112100313022100-3020301133310011-1100120132322202-1013233230322201-2211331010011122-1302101033323331-2321033213233221"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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

<a id="canonical-1022320132333323-0312210223033001-1132123011011111-0021111200112312-3333033311201122-1101302330312033-2101221011212010-2133103302003133"></a>

<a id="canonical-3300121032313320-2332221120232000-2132320132010200-1320202202110131-2021332211202122-0302203331333103-0132000220211333-1103203121213220"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.server_name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221): complete subsection reference.

<a id="canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-3130013303313213-3302322223320022-1323001030121102-3212321332123211-1332020301003102-0311030012332233-2331130302332201-3121301313323121"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-0333030101022212-0012103210331333-3000322301120210-1001323232301121-0003012002332003-3100022331332200-0213233213013102-2132310123130211"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-007.md#canonical-2020120202212302-0203103023210320-2111111222030002-2212211322121022-0202020233223311-1201103112121023-1000233112231110-1313230233220121): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-007.md#canonical-0302313111113032-1332101201121110-2132322021011032-2102211001003303-3011233030221220-0130010201300210-3201230000213021-3310320223021033): complete subsection reference.

<a id="canonical-2020120202212302-0203103023210320-2111111222030002-2212211322121022-0202020233223311-1201103112121023-1000233112231110-1313230233220121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-2100003222032312-0331332301000102-0232023011031201-3331333213122110-3332022333033130-3001203123221203-3031333123332320-2333100031022030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-0302313111113032-1332101201121110-2132322021011032-2102211001003303-3011233030221220-0130010201300210-3201230000213021-3310320223021033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-0200021310311221-3022223301122020-3000020202032313-3102312213202312-1301112111030302-0120023323231022-1020100300111111-2101322220220302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-0032013200101331-3121101033132032-0201111023102312-3310122312023202-1011012000331210-3022033220220302-0200212023113213-0311301132011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1213203211230211-3222310303202300-0112201021312230-2320000013033323-2331203131011000-1021020220302100-0310302112303031-0333030232221221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-2303033023003313-0320312301010320-1232130123120130-1031233120230310-3121203332301102-3100303300110230-2031223030231013-2131312113012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2032322313312233-0322013121000201-1233132312212212-3332212322230231-2123001121122203-2102200233310330-2210031033303233-2232330233333030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-2202303122103033-0101112331121031-2302302322233103-1211312111211132-3001231303232133-2132213010032312-3012203131322300-1130300102322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-0210303030313003-2030132311201230-2230021321113113-3332022102001000-3010230113223023-1323320100020013-3300300201202232-1031011333231101"></a>

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

<a id="canonical-2100001101020213-1220100210103220-1202213122130010-3121202002331001-1122131123302303-1211021303321321-1020132302133130-1012012113002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-1313031321203132-3333222313302030-1312113032133013-0332202033121103-3201210233021231-1131011003000212-3012200112021213-3330112110300111"></a>

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

<a id="canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2100111033000101-0122211330030000-2110030330103100-3111031213320112-0011021232231013-0130302232100032-3133323120012003-1002122332312131"></a>

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

<a id="canonical-1231222212133323-1331123303231210-0002033313110232-2332213323232032-0201100300021213-2102313303332333-2221131222203302-0230211012102001"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-007.md#canonical-1031212221332330-2102230133301102-3232001300110013-1222113100011112-1303112001102330-2000111220132012-2103131121223001-3302311113122332): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-007.md#canonical-3320031010211320-3011101111303120-1233002033320120-2011222132200301-3123232201023323-2123111102131312-2033100130113132-2003203113012202): complete subsection reference.

<a id="canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0130322020100113-0021302023333123-0211000313302110-2131232200122001-2133330210210033-1023223212030210-3222233332231000-1033330302210202"></a>

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

<a id="canonical-1120022313330023-3333303313032332-0301013313330313-2011033300212313-2132322301303212-1033103313113021-3323201212321000-3213012102100321"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321): complete subsection reference.

<a id="canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0022211303323132-1222002133111232-2331320232120231-2031010122110020-1333012221333220-1022133210130300-2120320332211100-1123013131302210"></a>

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

<a id="canonical-2133233322230103-3110233003311311-3303301323230232-3111233200223203-1320310221111100-2233210222102230-3231032012102312-0103301201312112"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-007.md#canonical-2323123103323022-0232220111220332-3100302222001000-3230202321222102-1202302031020232-1332022102101223-2121112001332111-2202132232200332): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-0012200202002333-2311131131332011-1312223033312002-3301223023332033-2120233112302231-2130320201012022-2221221311130003-0023013023223032): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-2213131103121312-1022113221103330-3321130300100332-1322010130302300-1310020232332330-1213122321110011-2332022013001021-0212121311112212): complete subsection reference.

<a id="canonical-2323123103323022-0232220111220332-3100302222001000-3230202321222102-1202302031020232-1332022102101223-2121112001332111-2202132232200332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1301311113011003-2133320121330213-2013320322010013-1031320110132103-3222022122213010-2133120032003013-2311333100001310-2200310031231122"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012200202002333-2311131131332011-1312223033312002-3301223023332033-2120233112302231-2130320201012022-2221221311130003-0023013023223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1123211312233133-3112223321202203-1002200332210223-3001032210333203-1231320133220030-2103300210110003-3210130132123033-3322310103010131"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213131103121312-1022113221103330-3321130300100332-1322010130302300-1310020232332330-1213122321110011-2332022013001021-0212121311112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1230333133100031-3312012213121101-3110101013021003-1110122131132301-3123223301101013-3202203232022221-0001333001211031-0212313331112310"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031212221332330-2102230133301102-3232001300110013-1222113100011112-1303112001102330-2000111220132012-2103131121223001-3302311113122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3113101330200330-1232203333212313-3122012131132003-3331321230201213-2100313111112201-2012303122311001-2001332130312202-2221100010202331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-3320031010211320-3011101111303120-1233002033320120-2011222132200301-3123232201023323-2123111102131312-2033100130113132-2003203113012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3023231221013111-2213111001203120-0232010001122112-3302031212331200-2311200310112013-1232120030023321-1232323133110221-3030321203001213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-0231223001232221-2201332131123020-1313300302020312-0221201321000320-0033233210023223-2110213013103111-0102220330311200-2331112323002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-3311322301021120-0310223201221103-1311202122202221-0021301331302202-2130332110222220-2330302012123201-1000300332231130-1320322210332313"></a>

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

<a id="canonical-3031130303320312-3011210220300230-1112101210130120-3233211133223031-3232211103310222-3322323003212131-0300113203000232-1001321332302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-0233011301010011-2201203111202133-3102020022301303-2321333002102212-3322311002321213-3312020232211001-1200013330210330-0200313302223100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-2213213201020321-0303321203211020-2202012103010022-2302121320212312-1023313101332232-3230002111032132-2311201002200103-3213110102330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-2010212131111300-0112030223010112-0203001303303010-2301030230002003-2103101123121222-3322330120132303-0131021022322110-0010303232022010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0031221303330331-2120132321333323-2130101022302100-2332130101101010-3331212021120222-2300232100211100-3000321103002032-1302110202101203"></a>

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

<a id="canonical-1033130131330030-1310110020301231-3100021031212211-2303322021002013-2022013301011022-2203200023212232-2112302203311212-0222131110310312"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](data-sources--workload--reference--group-007.md#canonical-1122123112031113-1311322231112010-1311313021221333-1332011303002300-0033333233222000-2220321100231003-0001301112320010-1330313001100032): complete subsection reference.

- [default_security](data-sources--workload--reference--group-007.md#canonical-3110301022202322-3120003223122202-0210213132332120-1122122112330322-1022112221330300-1323121311122332-1113213001001211-3330333200000031): complete subsection reference.

- [low_security](data-sources--workload--reference--group-007.md#canonical-1222030000033211-0230203211002333-3000321123130130-0120202320022023-1210033033100201-2331311023003130-0133111221033110-0123231022112023): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-007.md#canonical-3010321212113333-1303302230102330-1213023330002121-1130231200013033-3300232122111101-0221100211110312-2123031031021310-3223010111221222): complete subsection reference.

<a id="canonical-1122123112031113-1311322231112010-1311313021221333-1332011303002300-0033333233222000-2220321100231003-0001301112320010-1330313001100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-3120031211203232-2010320231212022-3010000013013002-1312133223332010-2333213003303010-2230201101322300-2232110100021032-1031033222213102"></a>

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

<a id="canonical-2200220223323020-3121303200323332-1030231023010013-0113003302023021-1333100023320132-2133202113133230-1331201002121322-0000210030112121"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-2223130123313210-0020200321033032-3120130223320313-2223322322103221-2223010202011231-0003211203132130-0332103032121011-0133313300312102"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-1110030212013200-2022203023311131-2103110311023232-1031330201231320-3033233302101301-2203210200312012-3300211322311231-3021301130023121"></a>

<a id="canonical-1110132100023312-1322112120212110-2030110132313323-3310300203312233-2230203021130212-2011231133210222-3132110100012021-1303000113122322"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-3303222103022001-3031123032202330-0311001320232300-1133230231211113-0010221323211313-2112013013130311-1232203121322232-1010232033030200"></a>

<a id="canonical-0202100210113211-0130310212211003-2232122010101110-2331330213130020-3221123012011123-0203000122200030-1031211323212201-0222030311221200"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-3110301022202322-3120003223122202-0210213132332120-1122122112330322-1022112221330300-1323121311122332-1113213001001211-3330333200000031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-3112103320132031-0103133202203211-0312303320022031-3222103210202111-2321333032023032-1233233103002303-2111221011032201-2122022231321210"></a>

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

<a id="canonical-1222030000033211-0230203211002333-3000321123130130-0120202320022023-1210033033100201-2331311023003130-0133111221033110-0123231022112023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-0010130220213121-1223133030201130-2212330102010101-3303013003331210-2231321230103200-2311331210220012-0031213200012101-2202201122002333"></a>

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

<a id="canonical-3010321212113333-1303302230102330-1213023330002121-1130231200013033-3300232122111101-0221100211110312-2123031031021310-3223010111221222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0222322111133123-0032222223222303-2321032220213233-3332033131123310-3232213231022020-0013230002233123-0030112320033302-1233103131012013"></a>

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

<a id="canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-3210223312013303-3213002100203233-1310321002102031-1132230203011023-3013032310122030-2011313022130220-1122113212333133-1100223022213203"></a>

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

<a id="canonical-1030032232300210-3003222221322131-0313120101011123-3132011320220130-3101302001000210-0030020133312213-0321231133202110-1001103331110302"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls`

<a id="canonical-1132110321310320-0011220103211020-1313200231233020-0301223201321320-3332123233001110-1320211221312001-3223103221103020-2320313333211300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-007.md#canonical-2321000000031102-3203222333122132-2203200210333132-2000101211323221-1100101121002222-1303302132031121-2201001000232031-1231122122132000): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-007.md#canonical-1213223112013210-0233221103323322-2101133230011202-1111203101300231-2021331221302330-3133100001330003-1032200101012223-2212313120122303): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-007.md#canonical-3322301003313000-0331132313112123-3301002020332202-2022220001000011-3032101221331121-2000011120023012-0023023311320330-2123103003232213): complete subsection reference.

<a id="canonical-0230330131022200-1201133323101323-1230132210120130-1123020223001322-1321002232231301-2003231103230220-2332133010021110-3210110122310310"></a>

<a id="canonical-3323333210121221-1000131310102302-3301202030300011-3011132132221332-2111123122301223-2022333133000031-0101101222111102-0010112100102131"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--workload--reference--group-007.md#canonical-3331323222331113-1111323202020100-1330010100013220-0123312320332130-2321310322332231-0311101220122021-3300002302031313-0021132000310301): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-007.md#canonical-3311120233030132-3103210132011111-3331020301021232-0201010233332213-1111233322120100-0311320320303323-1230032102001101-0121220011121313): complete subsection reference.

<a id="canonical-2321000000031102-3203222333122132-2203200210333132-2000101211323221-1100101121002222-1303302132031121-2201001000232031-1231122122132000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-0122002311033323-1000220103123122-0200033210200033-1132211320232333-1310012221002011-2032230223332100-3222312133032101-3002001033022003"></a>

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

<a id="canonical-1200223103133200-1301013223123103-3121323111210021-3111023333221101-1310013322130011-0301021122131001-1100210320110330-1002233213022202"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl`

<a id="canonical-0332032223322023-1211100232002032-0231301302120200-2301213111112003-2220100101121122-0123032211030030-0311122012022122-3112033121201122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-2022200122011223-2302012303232012-1313003122301330-2300323220212021-1110332320013212-3122203320200321-3111222331132302-1200331200213231"></a>

<a id="canonical-3033201121203100-3331202212313021-3112302323332122-3230300111111312-2332033230000232-2333121133131001-2312022100122130-0103210210230030"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-0222112222103212-0302300133022021-2200030031213101-3031013310222203-2031230312113313-3230210322022312-2202223200122030-2331012121020321"></a>

<a id="canonical-3312020313202322-1133212303100101-2203112122112020-3111311133333233-0020131130012323-2332010010312322-0310312132320321-0223102232301002"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-1213223112013210-0233221103323322-2101133230011202-1111203101300231-2021331221302330-3133100001330003-1032200101012223-2212313120122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-2300312203223212-1300103231113020-3302331110211031-2010123031031032-1121003300321221-3331130202220113-0212330321230231-0123132103331101"></a>

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

<a id="canonical-3322301003313000-0331132313112123-3301002020332202-2022220001000011-3032101221331121-2000011120023012-0023023311320330-2123103003232213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1303000121113201-3321230212202302-0003200100103032-0302321032312003-3220332302120333-2011300202200002-0220030311030222-2322321113222321"></a>

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

<a id="canonical-0331113121022200-3300030300122330-0310022303132012-1002200000030102-0112122002113230-2110000122203330-0031322022321003-3203222230013332"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-2333312313332133-0330310100202003-2032231213011123-1231301131213312-2023212302011013-2222203333113022-2212301202011022-2120223302221110"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-2301121132132113-3223102133302000-1110301330010013-1100122223303001-1320321032320201-2032213213301211-0021012112320132-1203030102211220"></a>

<a id="canonical-1230303102011312-0213212203102010-1231101010201123-0012232033300222-2132001322210230-3203003302113023-0310000131301200-1331321023121112"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3000133322120022-1032221210032321-0210122302210130-3020000323013213-2022120213301130-2300112200113123-2121100200222110-2212331103223322"></a>

<a id="canonical-1123102111010311-1113011110300000-0330002030011133-1212231122323130-3300022102133010-3021101203120011-1130223231321121-1003202033322320"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3331323222331113-1111323202020100-1330010100013220-0123312320332130-2321310322332231-0311101220122021-3300002302031313-0021132000310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-1000322312111210-0211323022022332-2100331110230333-3011023101332121-2130233013032312-1133213312112130-1130203110220202-3112011232212122"></a>

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

<a id="canonical-3311120233030132-3103210132011111-3331020301021232-0201010233332213-1111233322120100-0311320320303323-1230032102001101-0121220011121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-007.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-1131232031011203-0202313231003002-2321322231031312-0033233322202133-3012322003313330-2020020233022023-0121020330030323-0220100021110220"></a>

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

<a id="canonical-1011001020002102-1202020302220233-1323212220022310-3230202132002223-3003111200320001-1232122100002300-3022030212121230-1033230313322322"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-3320101012301111-2023111220123312-3210002000101200-0313322320121201-0313130032223113-0110001303010022-1011232211013322-0201333110212313"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-3300032310020202-3222222111323200-3232230300021332-0020030313020100-1032203000002030-0121223123121011-2330233221121002-0001000011032100"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320333002230013-1302101012111221-2100023202313133-0303233300303232-1300321230011023-1223232211223303-1101031000211103-3002332330020200"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes`

- [routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212): complete subsection reference.

<a id="canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-1103310133211320-1200303022033030-0301101133212303-1330231003331120-2112110012020222-1003013320302212-0011021212332320-2111120003303031"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0103100202223123-1313220120013203-2232230001303231-0110030001233202-0203221211032302-1302310113233203-2120313312312033-1233030222320301"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes`

- [custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221): complete subsection reference.

<a id="canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-0303230120203313-1101131022103103-1212010331323221-3123232013113302-3130311122132120-0200230113031202-2102022023323212-0233301313012113"></a>

Type: `"single"`. Computed.

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

<a id="canonical-3311320212020121-1132313231212001-0103203232203331-1331131022021021-2103110322201332-0032333133321310-0212033201102102-1031330230101302"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object`

- [caching_disable](data-sources--workload--reference--group-007.md#canonical-3033221110010120-0000213123333022-2212203332220022-0212200211130030-3033123113222311-1302020323102320-3320211300022231-0212023321202201): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-007.md#canonical-1310213032101203-3202230001230030-2333033331321112-1100211203120120-0221021123003333-2112112210132320-2332012020103010-1003032032131313): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-007.md#canonical-1111320223332200-0311233203030221-1010332113310300-2311002021001232-0212212120201020-0120103313011213-2322330113202223-1222023101231301): complete subsection reference.

<a id="canonical-3033221110010120-0000213123333022-2212203332220022-0212200211130030-3033123113222311-1302020323102320-3320211300022231-0212023321202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2022123120232312-1133301120201131-0321323020203122-1011033113021100-0232012100021021-0311120030030132-2220113231201300-3221012210121230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-1310213032101203-3202230001230030-2333033331321112-1100211203120120-0221021123003333-2112112210132320-2332012020103010-1003032032131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-3323113122202130-2133000001032120-1012113310331022-3303011012001030-2230111030331330-0001313302230300-1320330130103130-1101123303101132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-1111320223332200-0311233203030221-1010332113310300-2311002021001232-0212212120201020-0120103313011213-2322330113202223-1222023101231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-1110100212223133-0000033001301201-3232300020200233-2101030222033010-1022323203003103-3103310102131113-2013322321110323-1200332011001002"></a>

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

<a id="canonical-1111132111113320-0023012301230303-1031023113230022-1221120233201311-1313000101232031-1221202231322333-3003113033033231-2012133033032000"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref`

<a id="canonical-0322122002211100-0013031333310302-2101002112002130-2100023331203200-2202220310310222-2030111020101012-0123133130133122-1310332002322123"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.name` property

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

<a id="canonical-1123120321121132-1130100030113310-1002032331013330-0223033303223230-2030323001210321-0311113101000020-0130033231122002-2101210102031120"></a>

<a id="canonical-2123220312313123-3323200010110333-2331102311330111-1133013022020232-2030021313311303-0322220023321101-2030200103132003-3113003312031130"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.namespace` property

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

<a id="canonical-1233200031213312-0212202001202101-2023330021321121-1032002221311333-1100323303331101-3321303212210223-0021022132123322-2102111330011223"></a>

<a id="canonical-1211110301203032-3020111332121003-1233012211130213-3121333213323130-3220310121103131-0110033211312320-1220213110102000-2123032032021010"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.tenant` property

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
