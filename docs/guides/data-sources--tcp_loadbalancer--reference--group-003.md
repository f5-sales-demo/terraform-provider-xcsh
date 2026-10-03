---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-2231130202323123-0001112020010013-2013120212303231-0030000130233320-0121130203001230-1102320302011101-2311012302130120-3321300131120011"></a>

## Next pages — tls_config / 230320200133 / 4

- [tls_tcp_auto_cert.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120310133012112-0130123232123221-1322002121223002-2320133321120133-3001330312233013-1201132210000310-2010200313230300-3333102030011100)
- [tls_tcp_auto_cert.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1011222301233103-0101132330111200-3002231332032030-2203231131213312-1211210211020231-3200103321020100-1202011022023230-0032320221323302)
- [tls_tcp_auto_cert.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3223122001220111-2210032122330221-0020001102233131-3130011002220131-3311122323002312-1112210212322202-2312211202121331-3013210100112313)
- [tls_tcp_auto_cert.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110130132333313-3000221311101333-1312011230032310-0120312223201000-1101313331312023-2133033110100221-2320313021012030-1212212201231323)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3120310133012112-0130123232123221-1322002121223002-2320133321120133-3001330312233013-1201132210000310-2010200313230300-3333102030011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122010323003120-0130323320223002-0223021310200230-1112103123311013-0202201010320203-1201033232301121-3121101101230021-0310331002020230"></a>

## tls_tcp_auto_cert.tls_config.custom_security — custom_security / 112202120032 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-3113213220330012-3210103010123110-2322302223233001-1132212132131022-3131220221333211-1113012020101213-1202012113211110-0121032132032023"></a>

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

<a id="canonical-3302021011121223-0331220130322100-2011330331013001-1013310133331001-2102330330310332-1003113220330013-2103130031223112-0110002101230000"></a>

## Direct properties — custom_security / 112202120032 / 3

<a id="canonical-0023332303002130-3332030023122220-1331331200202023-2321231133130303-1201010012203210-3332123211201113-1333112202311322-1003113331202113"></a>

<a id="canonical-3011321010020200-0121220031320021-3010221113010130-3013300123101023-3203323010033310-0031210101033220-0323212131333310-0333021013222210"></a>

## cipher_suites property — custom_security / 112202120032 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2303021012331011-2313110201320012-0213001121013312-0203233333223222-1020121103331321-1320012003213211-2321333201303121-3313233203330102"></a>

## max_version property — custom_security / 112202120032 / 5

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

<a id="canonical-2133221010110323-1020302201323201-3310131311130010-0230001321100232-0111101000131201-3221131320303221-1012000032112322-0303123110202021"></a>

<a id="canonical-2002310303213121-1120131233200301-1320200102300112-3031123002003032-2213310302113212-1122230332222331-0333303010012202-1333301011310011"></a>

## min_version property — custom_security / 112202120032 / 6

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

<a id="canonical-3113303330110001-0000023231013120-2103021210211211-2312202022231032-3311332011032122-2130120131231120-1130312003212201-1233301121102211"></a>

## Next pages — custom_security / 112202120032 / 7

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1011222301233103-0101132330111200-3002231332032030-2203231131213312-1211210211020231-3200103321020100-1202011022023230-0032320221323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221321211200130-2223102123330113-0310203100320312-1310020332032110-0212120213300313-2022233021201200-3220223102322211-1313102212033201"></a>

## tls_tcp_auto_cert.tls_config.default_security — default_security / 322132003232 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-1231123211121221-3221333310333310-2312011201121020-2110201113100321-1300001133302122-1312233121120032-2301330030010010-3000123120323211"></a>

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

<a id="canonical-1332003012032333-1222132131021203-1033310302322032-2002200221010202-1110011202010303-0302202010020032-1003300220002111-1120211331020031"></a>

## Direct properties — default_security / 322132003232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023023312031211-3323301230211022-3221000102302220-1030302321122203-3022301301032102-1221011303023021-0132113221330012-1302232032111201"></a>

## Next pages — default_security / 322132003232 / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3223122001220111-2210032122330221-0020001102233131-3130011002220131-3311122323002312-1112210212322202-2312211202121331-3013210100112313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100300033010232-3013303213033312-1303131230100020-1101213202331200-3320110201203002-2223320112210033-1311012302013313-1030321323231221"></a>

## tls_tcp_auto_cert.tls_config.low_security — low_security / 212112111011 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-0111030230322333-3323310102022203-3323320111212030-3132222030022222-2333002112200212-3232012221322001-1113202311300031-2333231131011010"></a>

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

<a id="canonical-1321221222031031-3113100200132200-0221331230323030-2101213222003311-0131213000113001-1121122012310022-2210321220010301-1010303021222213"></a>

## Direct properties — low_security / 212112111011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130122010121230-2020113103120011-3303100312023301-2220300302310223-2021021221332302-2303231000103033-2212010212002313-0333033000101203"></a>

## Next pages — low_security / 212112111011 / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0110130132333313-3000221311101333-1312011230032310-0120312223201000-1101313331312023-2133033110100221-2320313021012030-1212212201231323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121210110000103-3110222313211303-1012233212211023-3211013030031002-3130000322330321-3031031031032202-0311002001013223-1123131333312321"></a>

## tls_tcp_auto_cert.tls_config.medium_security — medium_security / 023011233000 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-1303003213322311-0201122021101221-1302030000202223-1013001323232310-2300223002320112-1030311322130322-0320202103323000-0132020320332123"></a>

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

<a id="canonical-3130111012002310-2331022113100110-1232011112110213-0112102101000030-2232221232333000-0231003211031203-1131002033111023-1020321303001210"></a>

## Direct properties — medium_security / 023011233000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200130010313313-1311323103002301-1332301122212313-2332102012100201-2322120000011012-2120103302310313-2302121202002103-0000201302222002"></a>

## Next pages — medium_security / 023011233000 / 4

- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123303022132121-3223111020100103-0003110020011302-2111302220131110-1103320201133011-3322322223202312-0222021131122131-2232210102102113"></a>

## tls_tcp_auto_cert.use_mtls — use_mtls / 320202212032 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
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

<a id="canonical-2113232312122103-2212102311002221-0231131312202212-0103223330201013-1023232020232101-1312011002301123-3202333333310120-1122203112011312"></a>

## Direct properties — use_mtls / 320202212032 / 3

<a id="canonical-1310222003001333-0201010230220131-3213212222221213-0232301131302131-2213002223313320-3130200000013213-0222123101011210-1102010112200110"></a>

<a id="canonical-3110222022111211-2203133320231010-2231311203310321-3310002133230131-1201221101133331-3203133200033000-2333013021202111-0210032031121210"></a>

## client_certificate_optional property — use_mtls / 320202212032 / 4

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

- [crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0010333131321101-3231033121330022-0112020133320232-2110003013021123-0231013102001233-2321330023120000-2303303331230000-2332321222133221): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3133312230033302-0132212002130223-3222221103032001-0301212130000132-2200221310303233-2101033213012102-1213300302100120-0120031303133322): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0330323000330322-1300012003301320-1203000313012021-2203010122221221-1220002230210202-3322233303122111-2023202303003331-3223022223022002): complete subsection reference.

<a id="canonical-3213133123022310-2300013011102123-1103210132220113-3212121330103200-2321021231322121-2323011101302212-0231103102213012-1023210111133332"></a>

<a id="canonical-3301323010133020-1221023032220110-0230300233111301-3203332201023301-3110331303230001-1331201313320303-0303133031133111-3233031002022100"></a>

## trusted_ca_url property — use_mtls / 320202212032 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1112213133232221-3301331011033302-1031302121213100-3301322220000303-0020013000321331-3110203302003300-3111222210200200-3212121130012102"></a>

## Next pages — use_mtls / 320202212032 / 6

- [tls_tcp_auto_cert.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0010333131321101-3231033121330022-0112020133320232-2110003013021123-0231013102001233-2321330023120000-2303303331230000-2332321222133221)
- [tls_tcp_auto_cert.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3133312230033302-0132212002130223-3222221103032001-0301212130000132-2200221310303233-2101033213012102-1213300302100120-0120031303133322)
- [tls_tcp_auto_cert.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0330323000330322-1300012003301320-1203000313012021-2203010122221221-1220002230210202-3322233303122111-2023202303003331-3223022223022002)
- [tls_tcp_auto_cert.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120203032230201-1103300312312203-0023001222012311-3233112322211202-2011132033013001-0220212213223003-1220202202030012-3103212130313203)
- [tls_tcp_auto_cert.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1313300112223011-3201011300103321-3000222000110100-2200211103132331-0202331320000020-0023221203123333-3100213133300132-0220111231301333)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0010333131321101-3231033121330022-0112020133320232-2110003013021123-0231013102001233-2321330023120000-2303303331230000-2332321222133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302232233111222-3320302002211210-1010022030010003-0001203213302020-3330310221023103-0110201030233223-1011221120103011-2000313233232203"></a>

## tls_tcp_auto_cert.use_mtls.crl — crl / 032130311223 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-2023121130131103-1130230220332001-1103011133232103-3102131201233302-3312020120132220-0000322010023032-2121311233322131-2231233333002110"></a>

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

<a id="canonical-0111233130033222-2223213213202311-2333002303100200-2112320032000110-3212031300222301-3221131130102012-1332121013322031-0031112012021110"></a>

## Direct properties — crl / 032130311223 / 3

<a id="canonical-1121101213322130-1221310330020213-2101321311010310-3200322332131121-3311132212223030-0031122102303322-2222331010330031-1120332330003221"></a>

<a id="canonical-0302322320001221-0021203033022022-2032300203212113-0022332220200132-2131010212313111-1231112120202111-0232202103022300-1320122323220101"></a>

## name property — crl / 032130311223 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0302121230200103-3232120101130130-1303332233132112-1130201333001001-3233103023030323-2010213223300130-3110301002130221-0102223123332211"></a>

## namespace property — crl / 032130311223 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2232212111331122-3013010121002233-0200112302202032-0133203321100201-3132112312011112-3203030332330022-1313211013310031-0223202033220203"></a>

## tenant property — crl / 032130311223 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3133000121002213-0331013320313331-1113222011131120-3233203011122311-0133210220121330-1211121110010212-1030310231023220-3232212100103132"></a>

## Next pages — crl / 032130311223 / 7

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3133312230033302-0132212002130223-3222221103032001-0301212130000132-2200221310303233-2101033213012102-1213300302100120-0120031303133322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211332030123231-1022311310111002-0323330321131310-2213011102132122-0102330231303121-3102331011111033-3001311201100110-0231330001212010"></a>

## tls_tcp_auto_cert.use_mtls.no_crl — no_crl / 023301322010 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-2310311132122032-0233210101310000-0222220321230000-3113133333012201-0102321222031030-0000100020332131-0323131101032031-2331003112313103"></a>

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

<a id="canonical-2233301022103112-2122330331021030-1230102221123121-1013311000103300-1033300110301211-2113011333120122-3022232320132132-1133201003002112"></a>

## Direct properties — no_crl / 023301322010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230301311212123-2202020333212231-1300101310311320-1011212010101133-0002222031312210-3200300032003010-3332303230132313-3131321202003102"></a>

## Next pages — no_crl / 023301322010 / 4

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0330323000330322-1300012003301320-1203000313012021-2203010122221221-1220002230210202-3322233303122111-2023202303003331-3223022223022002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103121013120133-2321101013302230-3012211102132332-1331232203033323-3111311300333333-3021223131213231-1201022333322211-0011320211023320"></a>

## tls_tcp_auto_cert.use_mtls.trusted_ca — trusted_ca / 223133031010 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-0212132103311331-2230222302111013-3331301031021130-0230332012032302-0130203300223022-3220230331220113-3211323102100202-2002312331003211"></a>

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

<a id="canonical-0232332302020103-3030002100012133-0101003000200111-0310310111113302-0032330220110113-2032110231203331-0223130001212221-3021021000002001"></a>

## Direct properties — trusted_ca / 223133031010 / 3

<a id="canonical-2302212322000212-3323020331121211-2230020322321131-2133000212233221-3133101011132222-3113021123013021-2332012010300321-0100121033133033"></a>

<a id="canonical-3020002331132032-0213111303133102-2002203322211033-1222113310000333-0021023113020332-2102130030010120-0222201012330230-1030033301002231"></a>

## name property — trusted_ca / 223133031010 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2032002202213012-2110220110013011-0203330103201122-3023231203313001-2103100333021020-2302231313332112-1011323123311122-3120121101302011"></a>

## namespace property — trusted_ca / 223133031010 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3301320311213211-0211111130200320-1113201122203211-2010301130121232-3200202230321213-0101032211331110-3000321102202103-0003220302133221"></a>

## tenant property — trusted_ca / 223133031010 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0100323023332301-2202011200113200-3303312313022113-0023120002311210-0013213213333023-3211021222131010-2303232211102121-3222101130121312"></a>

## Next pages — trusted_ca / 223133031010 / 7

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3120203032230201-1103300312312203-0023001222012311-3233112322211202-2011132033013001-0220212213223003-1220202202030012-3103212130313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211230221000222-3021203302210113-1110312011221131-1211313112021111-2310232000300213-1222300111002121-0303103212113132-0022011233323223"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 201311123113 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3232232030103320-3130121211101303-1210132231322022-3013321130230021-0300032021203111-1332012233010021-2033122231021012-1201312103222301"></a>

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

<a id="canonical-3110233302202200-0013000110001302-1102002222123122-2311010230011320-3331211311132200-3211301130220131-2122122100333333-2131030123210320"></a>

## Direct properties — xfcc_disabled / 201311123113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122210201031031-1102302331233133-0132131301021032-1223101002113320-1223100003133312-0002211301033021-3201202211221012-0003121213011220"></a>

## Next pages — xfcc_disabled / 201311123113 / 4

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1313300112223011-3201011300103321-3000222000110100-2200211103132331-0202331320000020-0023221203123333-3100213133300132-0220111231301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122222310010203-2302303231320023-3112130130220303-0110200312331023-0312111012211213-3300311311231022-0130223123202303-1331313220221210"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_options — xfcc_options / 333321121130 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
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

<a id="canonical-2221010113001033-1321223003211331-1213302131300332-3323301320120100-1031232211122322-0001300331101221-3232233033131102-2012333303030111"></a>

## Direct properties — xfcc_options / 333321121130 / 3

<a id="canonical-1033233113222132-2323230213321012-2100311332211000-0303232301222200-2303311211222023-1121333223220012-1122311023322032-1232322210203201"></a>

<a id="canonical-2131123300323111-0330321313303113-1131000120131211-0300213021012232-3121002333030021-0110011101211322-0310313020000012-3301110003012013"></a>

## xfcc_header_elements property — xfcc_options / 333321121130 / 4

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

<a id="canonical-1210120301023110-1112013202232312-2233120332032232-1210312330330310-1130310213213132-3121133200123231-3101133211311231-3330312333032112"></a>

## Next pages — xfcc_options / 333321121130 / 5

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
