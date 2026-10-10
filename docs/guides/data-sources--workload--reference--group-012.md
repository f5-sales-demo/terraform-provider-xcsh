---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3201323022110132-1020103003321113-2200013102213323-0003030110012101-1111121302033312-1221200110010332-1331211330113002-0201221323100200"></a>

## Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-3232110320212120-0012210230110221-0200320010121330-2001100133010131-0213321212102000-2010221121032100-3120103222002302-0113330301201333"></a>

### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1302122131030222-3023203333220323-0133230312200003-3311132103211220-3112211031300321-1220203303032111-3211213320133322-0122302220132203"></a>

<a id="canonical-3210301200010233-3222002330030112-3112300301121101-1303333100320221-0010111021003223-0220102132002220-2131323211313130-3231000230230322"></a>

### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-0110100232201010-2013300132302000-0032223003013033-3213211302300031-1323023003210003-3220200113302321-1221333012102030-3213010132011030"></a>

<a id="canonical-3231312220311110-3103032223232132-1230213013202123-2302201213020013-0211020213013032-1032022223023100-1231300211002131-2112121300320103"></a>

### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-0310003021212133-2232022023323231-0220300122102233-0220131213301323-1130220303332302-3331022122122203-1331113212203333-2303021231002312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-011.md#canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-3100122223323032-2213110232310313-0001123200333230-1223331012201132-0212131313310010-3123231033020020-3002131110100103-1210332101220032"></a>

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

<a id="canonical-2110311100331112-0321030323233302-0112133013123303-1131210031000131-0210223112121121-2202202312231230-2222320120303012-1033331300203323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-011.md#canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-2113000103000022-2332112232222131-1122231111110113-1003120330201020-1113211100002012-2123323130110200-3113102333003200-2031032022322122"></a>

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

<a id="canonical-3202201333202302-1021302212332222-3231031323110200-3212020311123023-3200033211033131-0010033322122113-1303211333311303-1320301012133332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-011.md#canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-1300000100113333-2121021203221322-1132300131021013-3233131230110000-1102111210132003-1310120102220113-2231310232211130-3102300222233012"></a>

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

<a id="canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3233332212230110-0031211011302301-2232201002231330-1312121112303023-0201013121331030-3102330021201101-3312313030120110-3303110010231332"></a>

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

<a id="canonical-1023002032200113-3122131323323132-1132313111030002-0011002022110112-2100212333002130-2211002301221213-1322113231101122-1113000001031122"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-0222320032001110-0010233033033211-0012200121033322-0313010032013230-2232112310230023-1223310230210103-3231310221313120-3002201300013313"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-012.md#canonical-2022211031221201-0023301203102020-0230210203223133-1010023212213223-0020113100301310-1032201223223312-2020122302100230-2131002233323231): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-012.md#canonical-0100013111001103-2220103023130123-1120331233112003-1203010302323303-3320132223332320-1003021303202120-3331030022331311-3331011030322010): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-012.md#canonical-0202232021233001-3013332132211100-1231213202330030-1321230222210232-1010210110231212-3231021120013112-1320122123120122-2221031132001210): complete subsection reference.

<a id="canonical-3001112332003103-2331001303200310-3002302200202311-3330313222311022-1231302333020200-3232300021122101-1103100202032210-3302101321320212"></a>

<a id="canonical-1221123201110202-0302203230321111-1000033031300111-1303211102301221-2013031103310013-0301011113000301-1230101233322313-1030030123301201"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-012.md#canonical-3311110220000033-3012110200200022-1223330320331123-1203123010201031-3213320222333103-1212310003313303-0132230221003200-2302303321310321): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-012.md#canonical-1122023310123130-1100100003210310-2220200010131212-3210112022221130-2013013333003003-2032332001210311-2133000213302333-0333023302133301): complete subsection reference.

<a id="canonical-2022211031221201-0023301203102020-0230210203223133-1010023212213223-0020113100301310-1032201223223312-2020122302100230-2131002233323231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-0002120303013031-1302301202030213-2231320000221130-2313313111111003-0330220120120121-2303010230002012-1322110110010213-3231132201320310"></a>

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

<a id="canonical-0021131332223022-1332030303222202-1212322202130231-3032103010203010-2212333323123313-0132003310210122-2210002202031332-3310013322312203"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-3330121131333233-3311103311331320-2311001232120321-1013300233130033-2321320022102201-2020130223322131-3301300310103111-0212012222223320"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033303213131102-2033301210113303-3012123333330201-1001003230102032-1111023333303322-2221330033321123-3123001203212231-3000330200032220"></a>

<a id="canonical-3031211122123300-1201123031323001-3010020332313000-0200201300311101-2300011202002130-1032022133200133-2322111032001201-2133111123023113"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3310122023221011-1130133210202130-2130120222122223-0131012321331220-1223013022211130-0301300022112220-3331230113221312-3113002100313122"></a>

<a id="canonical-3211200000110332-0120012230331332-1333002131130112-3220230030022330-3123132013303000-0030210210332033-3021201233110211-2112221012003112"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0100013111001103-2220103023130123-1120331233112003-1203010302323303-3320132223332320-1003021303202120-3331030022331311-3331011030322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3103203323123110-1200232033112211-1030332101112232-2022013033213011-2322203221321232-2002013203132123-3121110011232131-2200002233210331"></a>

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

<a id="canonical-0202232021233001-3013332132211100-1231213202330030-1321230222210232-1010210110231212-3231021120013112-1320122123120122-2221031132001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0111101233313313-0311321022101131-3203321230003122-0331010000333322-2123201333122021-2223213112213022-3200300031231311-2211213211133110"></a>

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

<a id="canonical-0300121233322033-2302220333213033-3221031033202212-2112301033013231-0312113312233212-1231233011222222-1021223130111023-0100120033322222"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-0323033332201133-3012033332320033-0223003123002332-2132032332223212-3222102100302223-0122231000011330-3103221101113102-1230210332221231"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213233333233111-3301223330331323-0000030210313132-1112201100101232-2313011130330300-0012301203201131-3112210323021031-1330121332122003"></a>

<a id="canonical-1210103012233131-0112013132100201-0012133112110121-3100330211130021-0111223223303102-3001232202303120-3030330000100322-3201223220231211"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213132110021310-0123122033310010-1031232103333332-1021111223112233-0230300233211013-2102301201313321-0312333210123100-0233201220020201"></a>

<a id="canonical-3301113203112000-3301221223100203-1131032132033222-0220311212331231-0223331232030121-2001200201102200-2001122122121200-0220031310013003"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3311110220000033-3012110200200022-1223330320331123-1203123010201031-3213320222333103-1212310003313303-0132230221003200-2302303321310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2300102212033231-1332112301130123-3102031001130112-2020020011200031-3121200323123302-1213103001301313-2331131103232032-2123121031013033"></a>

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

<a id="canonical-1122023310123130-1100100003210310-2220200010131212-3210112022221130-2013013333003003-2032332001210311-2133000213302333-0333023302133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0120120110200323-0332331210110030-2310220012313032-2131222212031223-1222102130113322-3012112232201331-0003203020200123-2312232100030013"></a>

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

<a id="canonical-0222301122023221-2232021101313322-2220211333301101-3312230322002031-0323313101313330-2211202213320023-2312320321310202-1222313111333331"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0332232332211310-0223330211020231-1313121033312102-3021132120310321-0003011220032000-3030111110323032-1301023322103132-0120101232222001"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-3132023121203121-1303132122013210-0001232233031132-1012120322031123-1232332323000333-3323030012221101-2333133111003031-0313222011323132"></a>

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

<a id="canonical-0103221321331201-2230102332123310-0001302202310322-1122320132121202-1221211233032230-1313112002201011-0013111223222001-3130323000310102"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters`

- [no_mtls](data-sources--workload--reference--group-012.md#canonical-1312332321333033-1202033303330303-2031331023323102-2131002312223122-0333013101303203-1123012210313021-3103031022311103-0300312113330330): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-012.md#canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022): complete subsection reference.

<a id="canonical-1312332321333033-1202033303330303-2031331023323102-2131002312223122-0333013101303203-1123012210313021-3103031022311103-0300312113330330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-2310302233221320-0223323300210001-3010213320110003-3202232203211323-0223000312321100-0223111301333210-2003111131030011-3111121302303111"></a>

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

<a id="canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-2331230320303131-1023332001130012-1212211203013113-2310110113220000-1003221211032222-1312211323321021-1002013020332003-2332333233322332"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0332303300301312-3121020021003111-1221133302110131-2000033110332211-0230320222312011-0032023031222202-0300212232100030-3331131322000133"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates`

<a id="canonical-0331130301032010-1010022212003202-0033103213132112-2213222111022202-0021332303010300-1232200212233131-0020303320022133-1200122221011120"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_hash_algorithms](data-sources--workload--reference--group-012.md#canonical-0332201131011010-3001210001210131-0020302010100100-2110233222221233-1322122023102231-3122131311033100-0333222103011103-1102133231131033): complete subsection reference.

<a id="canonical-3310132102233232-3033320331130210-0001300102132233-0123203220323002-2210302323123300-3133210223103210-0032101233131232-3300221103012301"></a>

<a id="canonical-2130010103202113-2303230100320002-3303033312210100-0100013230001222-1202221231313313-2321331222301130-1333232220301103-1023111130003110"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-012.md#canonical-1022331031103033-0033310100101330-3133321223011301-2102033022232030-2323112011323212-0311000230011012-0020033313102233-3123300221120202): complete subsection reference.

- [private_key](data-sources--workload--reference--group-012.md#canonical-0320112122002300-1103300133202003-3122320321011121-2101030122200300-1303000120232222-1120000303030013-2201001102232031-3220132123133022): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-012.md#canonical-2101130100122133-3210212023013221-3331013032133322-2202131302103100-0033002202010133-3112130321103032-2231122230223002-1301031110311202): complete subsection reference.

<a id="canonical-0332201131011010-3001210001210131-0020302010100100-2110233222221233-1322122023102231-3122131311033100-0333222103011103-1102133231131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-0031032331122212-1100100002323100-2022100011322202-2313201012102111-0100330222222220-3113033313233123-0013210310121121-3112330212003022"></a>

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

<a id="canonical-1102320030303312-3232210112320333-2023023032312311-1211030203303112-0230121013232112-0022100033103012-3100120221130130-2332203231222022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-2300230222023113-1210023202123232-0110023331312003-1332033301120001-0021120201211112-1331011110110201-1303013132210133-3110310013100000"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1022331031103033-0033310100101330-3133321223011301-2102033022232030-2323112011323212-0311000230011012-0020033313102233-3123300221120202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1321200211232332-0031230031303310-2333221222023032-1322211312121133-1233202301320223-0112033231322201-0212332131023200-1333222212212230"></a>

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

<a id="canonical-0320112122002300-1103300133202003-3122320321011121-2101030122200300-1303000120232222-1120000303030013-2201001102232031-3220132123133022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-1231111312331132-3233033031011221-1101331333232002-1000120001232112-2201001221130211-2101110020313312-3032301302301312-0233030222211011"></a>

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

<a id="canonical-2223221200011032-2200210213320211-2133123321302101-2302011011133320-0313130220230233-1300111200231023-3120003200023312-2130110030213222"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--workload--reference--group-012.md#canonical-3000210222210103-3113133231033030-2000121233230130-2130033033023220-0301232030133323-3112012021111122-1200213013222103-1100223320313131): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-012.md#canonical-3103203000233103-0220201302123013-0213031030323213-2021013300000102-2120311133231123-3003023331211133-3203123301202133-2223133220031310): complete subsection reference.

<a id="canonical-3000210222210103-3113133231033030-2000121233230130-2130033033023220-0301232030133323-3112012021111122-1200213013222103-1100223320313131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-012.md#canonical-0320112122002300-1103300133202003-3122320321011121-2101030122200300-1303000120232222-1120000303030013-2201001102232031-3220132123133022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2201333230300302-3031322003123221-1231202000211022-2012203210131121-2011100013022112-0102320230213302-2022123332310033-1113322101312210"></a>

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

<a id="canonical-3101200120321233-0101103313133123-0310311232301313-3313133113230312-3333233100112000-0332210130133223-2003032322323130-0231331110020032"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3222223033032023-0103011321101030-3322232201300332-1022201311103222-0022332003113020-2131200332011020-0120002111300313-2121112132122230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1030131002323112-3031030100320221-1323001230223113-2101102222123131-3131101311001132-1010313111120322-1123201003201023-1031303310113003"></a>

<a id="canonical-1121102313132321-1103111300231222-1220113021221001-0030120230001031-2001031310230223-2113212333002210-3120323021222200-0331000002331223"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123033322201302-0101301310023121-2022333123013322-3110113212022031-3033030311222312-0132303013211201-3121023311222231-3130330022020132"></a>

<a id="canonical-3200022320311230-1023213302101332-1222030002133101-0031101220030130-0132130303011333-0300300021222310-2222010233130011-2332212231100102"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3103203000233103-0220201302123013-0213031030323213-2021013300000102-2120311133231123-3003023331211133-3203123301202133-2223133220031310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-012.md#canonical-0320112122002300-1103300133202003-3122320321011121-2101030122200300-1303000120232222-1120000303030013-2201001102232031-3220132123133022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1231111103321223-1321012131003010-0322000211113312-1102213332101101-3130213023021230-3332220232333320-2020211031332000-3110123031101302"></a>

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

<a id="canonical-0330121310332212-0312212033233113-1222331311300023-1202030310223122-1120223222122233-3010030322302123-3113221323033000-2211113133111203"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-1010321131102010-0300120033013110-1312212021312010-3212323222030201-2301100132300312-0311123123113131-0322130332322320-1002022113333303"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3130201230200222-0231331302100210-2303032302211321-3100032011003031-0100001002121012-1310322032331203-3202303101310031-2132011320130132"></a>

<a id="canonical-1200233022322200-2120310231323323-1203110311303112-1303310201120233-1110021101113023-0310013100113230-1033032333211002-3110210112313310"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2101130100122133-3210212023013221-3331013032133322-2202131302103100-0033002202010133-3112130321103032-2231122230223002-1301031110311202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-012.md#canonical-0100131231002011-1100333112012031-0233223013020130-3010103111231100-0031112123233123-2230220022231110-0031102233112213-0010020301110313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2301233232231330-2001220100121233-0302332000130012-2102021032031210-2221100212222330-0030011200023022-2310213010310021-0322210212022131"></a>

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

<a id="canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-2312020123012230-2020012022023002-2211220133121023-0202102320331121-2003311120130202-3031001103021313-2233321212123012-1222033310003121"></a>

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

<a id="canonical-3033120012121132-1000021233200112-0111230332302120-2211231210230203-2101322300213312-0302332320033202-3231301022311310-3131111221022102"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](data-sources--workload--reference--group-012.md#canonical-0021110020210010-2033132112102233-2332322312113000-0020203031333213-3313232123001203-0030303133013030-0321233332202222-0322033330203123): complete subsection reference.

- [default_security](data-sources--workload--reference--group-012.md#canonical-1003312021013123-2100202003303130-1313210220302011-2221320222213123-0301021032210120-2033033012002001-2202003002220233-0020033222221202): complete subsection reference.

- [low_security](data-sources--workload--reference--group-012.md#canonical-2031122312202031-0021110300213222-2313032110130323-1131011010012220-0332031333210033-0123020221232303-1101201133332311-0221201021221112): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-012.md#canonical-1000021223131120-2320310203111131-2102031101000030-0311130132022032-0323301011303033-1212230103112131-1000333233302120-1223320133122130): complete subsection reference.

<a id="canonical-0021110020210010-2033132112102233-2332322312113000-0020203031333213-3313232123001203-0030303133013030-0321233332202222-0322033330203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-012.md#canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1101220320211220-0010302131013030-2013200310211211-0222200012110022-1101211302131013-1110010313222103-2330021331221000-3032012032321210"></a>

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

<a id="canonical-1322200013113000-2001301010232132-2031121302211302-0033010210211310-3233333313303333-0130232332332033-2220033211122210-1330201030030110"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-1031331202132032-1200011331211200-3310221312130220-2331101202113223-0122210233222302-1300100021222111-3031122332103201-3232331123032021"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3013313223201213-0002233231003201-1320223012101232-2302210320332231-1122132101321022-1030021323303211-3310011313333313-1032330121312303"></a>

<a id="canonical-3203222131301312-0332201322221220-1210111332001333-0233231103202132-3011233313121312-3112232223223022-3301013223103320-3011212003133132"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-0120031111110333-1221310023232330-1323033030203110-1130311102302122-2023331102121332-2321133232233302-0012033010200001-1133030302013312"></a>

<a id="canonical-1201132122201223-2013132331311002-2133030120022210-0033103032310221-0001223101032331-2121003320130111-1313323320032033-3012233223000001"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-1003312021013123-2100202003303130-1313210220302011-2221320222213123-0301021032210120-2033033012002001-2202003002220233-0020033222221202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-012.md#canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-1003332313323003-2200113322122300-0033221231302212-1232211221230212-1132212133010130-2223122301223221-3211000201131231-2033210202202112"></a>

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

<a id="canonical-2031122312202031-0021110300213222-2313032110130323-1131011010012220-0332031333210033-0123020221232303-1101201133332311-0221201021221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-012.md#canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-3200110122202202-0233321012220212-0033221122121222-3013002210202002-0012122322211112-0001333230021230-2212232210133213-2301230001112012"></a>

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

<a id="canonical-1000021223131120-2320310203111131-2102031101000030-0311130132022032-0323301011303033-1212230103112131-1000333233302120-1223320133122130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-012.md#canonical-1233123132311010-3302222011301201-0303230210001232-1213133310123103-3332311322230100-3200110220022303-0133013303011102-1301312020230131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-3020021001333130-3133022120103203-0331300302221333-3301330231020320-0303310213322020-1213130323313021-0123323300320232-3321302031121232"></a>

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

<a id="canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-0212220133320322-0031201220012012-3101333220021032-2320200122211310-1131213332011000-0011100231213300-3133321101002020-0131032330120323"></a>

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

<a id="canonical-0320133121121132-0222102113323022-1101001023120300-1102232322332023-2111021121020311-3201103312011112-2212230112321322-0310011112112201"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-2222302231033121-2112022232022321-2320222003121100-2123230000320332-3313130113212223-3211100302332113-3002223200101232-2113320103221221"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-012.md#canonical-2230202213201113-2320213322201131-1003230101200020-1013302322101001-2331222230232210-3230302122210023-1331213020332131-1213111113110033): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-012.md#canonical-0222001033223210-0212132002210130-1030213331211101-3110023220212123-1232113203103203-2300031301033120-1232112123203303-0230331011130123): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-012.md#canonical-2032133233021011-2002200332213103-3120300100123112-2330000133320100-3130002321312211-1312011101211322-0100332301230323-1323010230013323): complete subsection reference.

<a id="canonical-0000223333203112-3303332111122122-2333230131222122-0231313230001030-0320220312111321-3210310313001302-0311200120022210-2013231233330323"></a>

<a id="canonical-1312301020210220-3331231032101103-3032000030033232-2213010002233301-0033130030102020-2020222222300103-3303022013020120-1103302221101310"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-012.md#canonical-2132010001001022-3011000133201002-1323201023102113-0013101122020300-1323223312122302-1331210311123223-1003130102222132-2303000033001132): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-012.md#canonical-2232000321032301-3122011303100322-0112232013233013-0331203120122332-3121103133103112-0021123312313102-0233122202022012-2133031031213130): complete subsection reference.

<a id="canonical-2230202213201113-2320213322201131-1003230101200020-1013302322101001-2331222230232210-3230302122210023-1331213020332131-1213111113110033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-1223310231230100-3212221011122221-1333211010303103-1201031332111300-1023201222221101-2203330213032201-3132123012201113-1223213100033123"></a>

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

<a id="canonical-2013222131321130-1300333103033203-1002222212211033-3030111102312002-2001121020122303-3322011032203121-3200331200302000-0320003211311001"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-0313121310331212-2302311313132231-1330013220122110-3013300120121210-1222310003221012-3100212220030133-0303012123102202-0121003130213020"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0210303233110230-0101001232012301-3102120130133320-1201002123121023-0113202213022232-3300231122322220-3322030101122112-0010021110323232"></a>

<a id="canonical-3132203213210030-2122120230203310-1022113010030221-2313110112033022-2121020002031330-3111101222123133-1223322103203123-0021333130220222"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0232311330310200-1223211313110120-0111013310321313-0013301301310010-2103310202211223-0201123321231102-2130232321113021-1213201230030003"></a>

<a id="canonical-2122202120223132-2101133122002130-2301023102303023-0320322103022033-2331002033131001-3012313111102331-3001120032210010-2203101321020231"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0222001033223210-0212132002210130-1030213331211101-3110023220212123-1232113203103203-2300031301033120-1232112123203303-0230331011130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-0010120131020103-3320020231232301-1332332222113033-0112233130330022-1032013332331323-0013211220233120-0322223200100233-3020311033303102"></a>

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

<a id="canonical-2032133233021011-2002200332213103-3120300100123112-2330000133320100-3130002321312211-1312011101211322-0100332301230323-1323010230013323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-3020232123133233-3001203321230233-1001302201001030-0233011002311210-2211101102112101-1233200310003130-1120333000031332-2110233103123112"></a>

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

<a id="canonical-1103122233122001-3322010230023222-1001132231221313-2121313120232110-1212311132130203-0312302002323233-3013122110131231-3101333100313112"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-2233201032212212-2133010122000000-1010132123302131-0330310322012323-1322301212220101-2013131110223122-0203232213130131-2021012010323203"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2100011022100302-1311021102030203-0320012313201103-2212310311102130-0013113100320001-3322033030301200-1200233323311001-0111221230202321"></a>

<a id="canonical-0011220300333313-3130211102112013-0213033130210103-1230310113013200-3222301022303213-3031123200001113-3313121130110201-0102311303321103"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3322223123012333-0021311223211332-0331303300031000-2210022233022213-3231121001321201-0323202021113323-1001023213021103-3320201032132220"></a>

<a id="canonical-2122310313330121-0023210023213132-2012121213010221-3032311320303113-0011102221202310-2202200110212021-0101212322022133-3022200201100220"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2132010001001022-3011000133201002-1323201023102113-0013101122020300-1323223312122302-1331210311123223-1003130102222132-2303000033001132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3113310312103031-0230222101301202-0103201211001021-2013111232003013-0221210231202311-1101032321013333-3012211321222321-0132333211300220"></a>

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

<a id="canonical-2232000321032301-3122011303100322-0112232013233013-0331203120122332-3121103133103112-0021123312313102-0233122202022012-2133031031213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-012.md#canonical-3302330023110201-1230212033321231-3331302310111320-2012310110300301-1321330032133121-1032223101110333-2300130130212232-2020000012122022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2120223103220022-2110012121131022-3103302332231120-1031110102001310-3130200022120231-1002211003310312-3313030010223221-0330201232233113"></a>

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

<a id="canonical-1112223033220321-1233111020013313-1022003101030113-2301101302001200-1111002301102110-1322031000213210-2013032120320113-3233111022201200"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-1311020233223302-0030320021301100-0333320331223230-1030201120103001-2012322000201201-3033311133202333-3032213220133232-2123321201112221"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-3310020223012210-3221300313011020-2133303032230000-0211033310112122-3011030022312320-3132100202212132-1002310023232201-0303102032033012"></a>

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

<a id="canonical-3013321111111320-1122212013302003-0302031023210200-2321213023030020-2320312130322001-1002012230001313-3313111221212231-1231230133031300"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert`

<a id="canonical-2322321031121123-3100301200331012-2232022020023032-2023002321232031-2311103202332303-0332112230103020-1102012321131020-0020113000002322"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.add_hsts` property

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

<a id="canonical-2333003013102211-1210223103113112-1013132232222131-3331010212012320-1321132222330301-3202101310232120-0002031112321312-0011111203302012"></a>

<a id="canonical-1021212313303022-0101122122012323-1021330321330232-0002001333101323-0223101112203311-0121231220033033-3232130311302331-0321301120233133"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.append_server_name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [coalescing_options](data-sources--workload--reference--group-012.md#canonical-1001303023332201-2003320231033010-3211130131310030-0201230021000101-1111321320133220-0130001300112113-3120311101313010-2132230033231303): complete subsection reference.

<a id="canonical-2232000132322210-2310211122120222-0203210333131131-3103323123013011-2021132031020023-1332330030333332-3332302201333220-0031012132121210"></a>

<a id="canonical-0300220002013311-3313202020203011-3211013012033123-3010033223120112-3010020331001333-0302020201211320-2222010122133032-1122212322310023"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [default_header](data-sources--workload--reference--group-012.md#canonical-0032210220203231-1313100300121023-1332120320123201-0103300221322211-0310112133102311-0212323330123322-1331222111331131-0232311101332121): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-012.md#canonical-0203033233230321-0022211303300223-1013233113300233-2310130030332310-2111032023021200-1321102213200001-0002232211130130-1312302120101001): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-012.md#canonical-0031212213223302-1132300301101023-0101022202110231-2023200023220131-2202100113222010-3332200222030031-2101003112111030-0001031020020303): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-012.md#canonical-2321032231322301-1111200012223113-1220100123020032-0211011211313320-2022131333032333-3330220102230000-0010302100223223-1112323103130102): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131): complete subsection reference.

<a id="canonical-1221123012223303-0023231103312001-2311030333303223-2311132003112112-3320020121120233-0003201121100032-0210200212322122-3330303012322100"></a>

<a id="canonical-2223110110022020-3201002313102131-0131300030102133-2201020102312100-3310031123132102-2223000012201312-0310000311121313-1110321231102121"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_redirect` property

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

- [no_mtls](data-sources--workload--reference--group-013.md#canonical-0201100323213210-3121011203320200-0212020201102002-2132200113002000-1203231120101200-0310010130002222-2033203321211332-3120001003110021): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-013.md#canonical-1122012103333122-0212322212331233-3000313111233001-1003213120030012-1010031023312333-3010310122303113-2033023223230000-1030130010211330): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-013.md#canonical-3330202110233223-0000220122333010-1212310020333112-1133022312113120-1301022301310333-1121332022110011-2333302321003312-2102333003300331): complete subsection reference.

<a id="canonical-0312210112302223-0131311131323212-0133201302300333-0033202011212030-0123020133313123-3331021222011332-2321220133311312-2203001223232222"></a>

<a id="canonical-1031332023212300-1121003013311303-3331020203133001-1231210312110022-1102213100013132-0210111233301302-1332133032230322-2012203201323230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3022233023132232-0302113003223203-1111112133023301-1031210000010322-1032313231033032-3011001233310100-2011001001232112-0002231033111202"></a>

<a id="canonical-3302113300023301-0210133132300022-1322102303222210-0221122031233203-3210222030130003-1330131202311300-3120101103211031-2132302033300120"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port_ranges` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123213322302231-2313232130112113-3321320212123122-1013101122203203-0132201110002001-3202313123131312-2323022021120001-0112103323310201"></a>

<a id="canonical-3131011310032231-3033103031220321-2123322013030120-0013122112202112-1313013113222212-2233222300112232-1330301232132102-0101100131103313"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.server_name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [tls_config](data-sources--workload--reference--group-013.md#canonical-1000301010323030-2230211020313020-2313002312130000-2232100021000310-2123103333020211-0103021312321023-0011031203020202-0313212132232313): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-013.md#canonical-2003002322222321-0102111232120301-3023013123301100-1021001100031113-1100221032032303-0033012220321031-2232320310132322-3211320332132302): complete subsection reference.

<a id="canonical-1001303023332201-2003320231033010-3211130131310030-0201230021000101-1111321320133220-0130001300112113-3120311101313010-2132230033231303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-1022200313202110-2302032232331213-0110321221133223-3233123012203113-3012102003011232-2030323221300203-2302220300110330-2212131230213122"></a>

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

<a id="canonical-1321212213021032-3332030310022122-1121100133112002-3202313111311231-0122232222011032-1102321220201310-1131312001032020-0111030012020103"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-012.md#canonical-0013121022022133-1212000002303022-0221123332013012-2101113323311212-2332123010120012-1312300000210302-0233300203112311-1113300323021012): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-012.md#canonical-0212303332012031-3030010102020201-0330203122113102-0321333011103131-3312020221001330-1213200132320301-1222230301132132-3303130020100003): complete subsection reference.

<a id="canonical-0013121022022133-1212000002303022-0221123332013012-2101113323311212-2332123010120012-1312300000210302-0233300203112311-1113300323021012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-012.md#canonical-1001303023332201-2003320231033010-3211130131310030-0201230021000101-1111321320133220-0130001300112113-3120311101313010-2132230033231303)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-0101102132212032-3121012201223211-1201313001333203-2001031330301022-1202113210310302-3120231111330023-2331003000110310-0101003022113201"></a>

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

<a id="canonical-0212303332012031-3030010102020201-0330203122113102-0321333011103131-3312020221001330-1213200132320301-1222230301132132-3303130020100003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-012.md#canonical-1001303023332201-2003320231033010-3211130131310030-0201230021000101-1111321320133220-0130001300112113-3120311101313010-2132230033231303)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3031230113232031-2320030101332212-0302321221030130-2313012032121032-0222130122130113-0123112320312000-1100220113323221-3023200332131233"></a>

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

<a id="canonical-0032210220203231-1313100300121023-1332120320123201-0103300221322211-0310112133102311-0212323330123322-1331222111331131-0232311101332121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-0312201023322003-0121332022033211-1130223032100121-1133001322220010-0022021112123332-0111001220332131-1010102120203233-2212011303300011"></a>

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

<a id="canonical-0203033233230321-0022211303300223-1013233113300233-2310130030332310-2111032023021200-1321102213200001-0002232211130130-1312302120101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2301322002013301-0201020110220131-2123031200212332-0311323211213310-0223103102133202-0003031220213201-3333203001310121-0101320301112032"></a>

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

<a id="canonical-0031212213223302-1132300301101023-0101022202110231-2023200023220131-2202100113222010-3332200222030031-2101003112111030-0001031020020303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-2201223311303003-3323030013022031-3222022110303102-3013333033130001-3031330122021003-3320020023331312-1032012131011112-3233330313113210"></a>

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

<a id="canonical-2321032231322301-1111200012223113-1220100123020032-0211011211313320-2022131333032333-3330220102230000-0010302100223223-1112323103130102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2100112012133211-0200302203221222-0122300001312023-0220330212003310-0203102021030312-1231310011221222-3011201332023120-0021232103331322"></a>

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

<a id="canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-3332111203303200-2331222033213212-0203210110210301-0310301101212232-1223002120330331-0213303213121103-2201231102031012-3221312032011323"></a>

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

<a id="canonical-0123211200022221-0231323003333331-3321300123122120-3020120032021110-2232033200133230-1303322020132123-2113133310130332-0021200233030220"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-012.md#canonical-3003113120030330-3231103231332320-2322220100101003-0000031010213210-0122102033313333-1332301030300133-1311120001331102-3023121001121331): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-012.md#canonical-3022313130100101-3210231221220030-1332210213220101-0302002120003300-1031313302313023-3212110210302232-3312101103101003-1312011033112102): complete subsection reference.

<a id="canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1220312110212130-0301010221213120-0020323201000121-3231202233201230-1112322300222013-0103100210123210-3022031020232313-2003022203203021"></a>

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

<a id="canonical-0133203212003320-1213122200203122-1220213322222030-0013012032023200-3120031102130233-0232223220303113-2331110032310132-0103311233231131"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-012.md#canonical-1113233033130231-1222101333112301-2221130231233023-0033300232210012-3121033210121212-1111332322313323-2102112313311113-0323323010303133): complete subsection reference.

<a id="canonical-1113233033130231-1222101333112301-2221130231233023-0033300232210012-3121033210121212-1111332322313323-2102112313311113-0323323010303133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0303232212030222-2100232313310230-3030113011213130-0131101213001300-0320333101232213-3211132121021120-2032130120231101-1102112303111103"></a>

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

<a id="canonical-2010033010133002-2313011202221033-0330013131103001-1213110013130210-2203110100320132-1022221021220113-0123303333233000-1323220300333000"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-012.md#canonical-0112100302131131-1021303112121103-3213210001032130-0303322233113223-1203312002033230-3110022033232201-2030133302200132-1010000033312122): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-3221322033233221-1011210222332101-3321021233030002-3230001122320131-1330123122232320-2323023102303221-0301010032122120-0311333222203120): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-012.md#canonical-1221033122110031-0001110021331223-0300232231131202-0332303220100131-3311201313330230-0331201211022211-1020120021001022-3110300201233213): complete subsection reference.

<a id="canonical-0112100302131131-1021303112121103-3213210001032130-0303322233113223-1203312002033230-3110022033232201-2030133302200132-1010000033312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-1113233033130231-1222101333112301-2221130231233023-0033300232210012-3121033210121212-1111332322313323-2102112313311113-0323323010303133)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2220020122230210-0133322021232133-2123312220200332-1203210020233031-2101311300331021-0103332321312223-0031033322022033-1002101201311322"></a>

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

<a id="canonical-3221322033233221-1011210222332101-3321021233030002-3230001122320131-1330123122232320-2323023102303221-0301010032122120-0311333222203120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-1113233033130231-1222101333112301-2221130231233023-0033300232210012-3121033210121212-1111332322313323-2102112313311113-0323323010303133)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2230102330133222-0201021130133021-0011202320032130-3000012313023331-0100020311210102-0133103103113022-1231212000232201-2230133101320301"></a>

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

<a id="canonical-1221033122110031-0001110021331223-0300232231131202-0332303220100131-3311201313330230-0331201211022211-1020120021001022-3110300201233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-012.md#canonical-1101233102223221-3320311323330113-1000003021321312-3310321312211300-3203031123201221-1222220202131320-3101033301010023-2111202211023311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-012.md#canonical-1113233033130231-1222101333112301-2221130231233023-0033300232210012-3121033210121212-1111332322313323-2102112313311113-0323323010303133)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0102111330332003-3201203101131132-3310121020323200-3210011001132301-1021213211301001-0113223322211323-1221202010010022-1033020101021233"></a>

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

<a id="canonical-3003113120030330-3231103231332320-2322220100101003-0000031010213210-0122102033313333-1332301030300133-1311120001331102-3023121001121331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2332201222301110-1122332221002331-1231123303133203-1323013222101003-3333022011112311-2302030123032320-3022201320123001-0333020130103320"></a>

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

<a id="canonical-3022313130100101-3210231221220030-1332210213220101-0302002120003300-1031313302313023-3212110210302232-3312101103101003-1312011033112102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-012.md#canonical-0313210132232031-3113221121100223-3321013030003203-0222023000121030-2122313000303123-1220232120122202-0301132112302211-3232100311330131)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3120311130222101-3200201000322033-0020221021123133-1023022000223312-1011302221000222-1101122000131232-1021222000132301-0303303031313303"></a>

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
