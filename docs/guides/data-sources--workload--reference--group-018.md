---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-2303001302113120-0033131212123132-1020133320301310-3022303221032210-1020210022101202-1121321321132000-3223010102122010-2100211021303300"></a>

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

<a id="canonical-1201210201120121-0133021323100211-3230211200113111-3230030011322301-2031213100031002-1311202112303321-2003001101300201-2013101022012002"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](data-sources--workload--reference--group-018.md#canonical-2130031201022220-0133221130201233-2102011123210221-2202110130121332-1220113211222100-1230021313011331-2022310200200003-0210210131100320): complete subsection reference.

- [default_security](data-sources--workload--reference--group-018.md#canonical-3123213332103212-3130303033213300-1110022100031103-1213231303323131-0322023002013011-3320301030322233-3101011003103123-2301000010323030): complete subsection reference.

- [low_security](data-sources--workload--reference--group-018.md#canonical-3022033303210332-3200132323122021-1033222301312132-3231011132121212-0033002031033030-0003010122311011-2310332103101321-3301300130002310): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-018.md#canonical-1012220221211101-2202223232222013-1313013232111320-2203112323012220-0210231312322300-3212313222311322-2213103031220123-0320011201221301): complete subsection reference.

<a id="canonical-2130031201022220-0133221130201233-2102011123210221-2202110130121332-1220113211222100-1230021313011331-2022310200200003-0210210131100320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-018.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-2213303212322020-1033020031020310-0012212100123333-0131220230130033-3000302100100312-3233003032030232-3031012011001002-1032003223303022"></a>

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

<a id="canonical-3002320121000233-3000020103102133-1000213220113020-2312120300033220-3233200200231031-3003120113023112-3020223321333320-0001333231320201"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-1020020021133101-2211033033300013-0001132230323032-3111021230001031-2023101121112323-3131333330011310-3332333002311101-3201012312012200"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0110303333111330-2330122323201020-0123003213110323-0131313123322112-0013011022011030-2313220033113031-0110320203133300-2210110332300133"></a>

<a id="canonical-1103112210321321-2013303313111210-0031332120132032-2102100010133330-1031003210131031-0011120223213123-3322103210131221-1030031113133011"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-3121121201013210-0100313320333113-2132211110133002-2313231123032131-0131303123001233-1221032202131231-3133133321221231-0003122121113331"></a>

<a id="canonical-3210320200233321-2320122020303123-0103131323320100-2131233122003132-3202333233303103-0311021011313320-0331212132302233-0320121102303210"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-3123213332103212-3130303033213300-1110022100031103-1213231303323131-0322023002013011-3320301030322233-3101011003103123-2301000010323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-018.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-1202132113333001-3331110011203111-1001203212131313-3033200222222121-2032100301121220-3221333111302313-2020312310300130-0032221311020003"></a>

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

<a id="canonical-3022033303210332-3200132323122021-1033222301312132-3231011132121212-0033002031033030-0003010122311011-2310332103101321-3301300130002310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-018.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-1200302303301020-0031210101130122-0310303330101321-0223323331301322-2231033330320322-2031133220021011-0212010212200132-2310033030032020"></a>

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

<a id="canonical-1012220221211101-2202223232222013-1313013232111320-2203112323012220-0210231312322300-3212313222311322-2213103031220123-0320011201221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-018.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0232123300032021-2013201002213311-2120113213322121-3211333223121333-0012100323002033-1232120020233013-1203031200330222-0211120301020100"></a>

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

<a id="canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0322332201323022-1122120232033203-1301032100113303-2022321320233020-3222311100101232-0333121220100013-1231201023021310-0120220103130101"></a>

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

<a id="canonical-1220311133000222-0121220202313110-3333200321323010-3000013210201002-3121310203102102-2102221330100011-0001203010303112-1010333213020011"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls`

<a id="canonical-3110311221311233-3022011032032203-3021120332322012-2121113031221022-0111333003010211-0321031023321113-2221132213133110-1332032111310330"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-018.md#canonical-3331013300223021-2021212101022320-0232331132021013-2201032323231321-3301211130331122-0031020301030023-0202000012031300-1201131000331000): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-018.md#canonical-3300203031320333-2202000101011002-2311012301013130-1223110233110110-2231020123010023-2132033123123233-0013011302010313-0103111122001032): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-018.md#canonical-3222203102011301-1122112132000001-3303011031321001-3021322030033131-3311101211302021-3322211310122003-0202222133220230-1221303023230112): complete subsection reference.

<a id="canonical-2220131330323220-3211003123233123-1133002033203110-1001012220310313-2302233202230231-2000201210303020-3021023001003111-2111022330323323"></a>

<a id="canonical-2012332122220321-0110110322320300-1213011032100211-3120011302221300-0222211333320221-3131310311110210-2333020312010331-3322302303230301"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--workload--reference--group-018.md#canonical-1032212120033232-0302320103321232-0212301010212301-3002200131312311-0300103001211311-1021322302233201-1210000122302100-3331221200122301): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-018.md#canonical-0233310023013011-2212030213101032-3320003003000333-0303120231113103-1120310323222012-0213031012132300-3303001333220021-0031330222131333): complete subsection reference.

<a id="canonical-3331013300223021-2021212101022320-0232331132021013-2201032323231321-3301211130331122-0031020301030023-0202000012031300-1201131000331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-018.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3223031202302333-0032130122133211-2233230331100122-2330313321101013-1132321130333033-3000230031213320-3330032123300111-0003233031133012"></a>

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

<a id="canonical-1223302330222120-2021323032313233-1133313123232313-2201003222133023-1200032200013112-0202302200122112-0002013303123320-1221203102222202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl`

<a id="canonical-0202012223101120-1321132210021023-2111303003220202-1012110110023230-0020113010332100-0300102203021132-1323320000220332-1213103121102203"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-3122213323122000-0111021110032233-0222212313322123-0333211000313122-0030220032212311-2311303320130111-0013322132110020-1222010320321103"></a>

<a id="canonical-0230310233021131-3313221000221120-1332002131031300-2121221221030310-3201130132030011-2320311231002101-3210123121323032-2220312201030223"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-0133000300010110-0103013022112212-3131200003211031-0013230311312012-3132023122230203-3120122013012202-2000220231100210-0000333220021112"></a>

<a id="canonical-1303231122210310-1002212020112022-0300231123122113-2300322232022012-3212313300322131-0131002131300302-0323102220223200-2221333130212230"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-3300203031320333-2202000101011002-2311012301013130-1223110233110110-2231020123010023-2132033123123233-0013011302010313-0103111122001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-018.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-3320132212121120-1113302022120121-1021110030200010-2002123232132003-3322010000311202-0333022213123323-0030023031012101-0130030123000100"></a>

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

<a id="canonical-3222203102011301-1122112132000001-3303011031321001-3021322030033131-3311101211302021-3322211310122003-0202222133220230-1221303023230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-018.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-3010300223202003-3311231102112000-2323001021002330-3100200112021113-0310332321130332-2210101122110230-0332212011323222-2131003331221213"></a>

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

<a id="canonical-0031303321303311-0110123231023220-3330200332211003-1012130222120111-1001220321123311-0203011313200231-2000111221203101-0302330012033222"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-3201311012310121-0302113122032100-2331231121303031-0331013032213312-1030101330102331-3003001100200130-3103333113320022-0120210022030000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-2230102101203221-0321322323221202-3333100333323101-3302021120022022-1300323121302110-3232212022123123-3210332010202212-0122022200330112"></a>

<a id="canonical-0320131123001332-2303222200003100-0232033200100321-0111303023220322-2132230113002233-1000300113333113-2033222121113120-3311333322231311"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1020301202011130-1320300130103101-0232312233030322-3102223200313303-1302333002312011-1132221102210233-1002031011320201-3110031333001311"></a>

<a id="canonical-3322123020121312-3302312022120010-3032120211202113-0331300220230310-3022033301321320-1223120221010020-2021230212031331-1320022231020232"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-1032212120033232-0302320103321232-0212301010212301-3002200131312311-0300103001211311-1021322302233201-1210000122302100-3331221200122301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-018.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-2313200121300223-3200111113003103-1322021303033010-0000123212220131-2202133022223221-3030322313100131-0102120311112210-2011112013020033"></a>

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

<a id="canonical-0233310023013011-2212030213101032-3320003003000333-0303120231113103-1120310323222012-0213031012132300-3303001333220021-0031330222131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-018.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-0203021111130232-2201123322313020-1112222330233233-0103032333012220-3201002303110030-0310103220310223-0131212203310223-0322002130031333"></a>

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

<a id="canonical-2023032233321322-0303132332033321-0333103121210021-0230201021202101-1000133032000131-3311331231002021-0131221102212303-1132033303312103"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-0013021210122023-1300012233221231-0023213113221012-2330120130110223-2032133100330132-2201213131100030-2123321303223032-0331123323323122"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-0012003120033331-3110121011013232-3100003112101033-0023300212031133-2120210210132213-1113102311112321-2103310331000210-1033221333331330"></a>

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

<a id="canonical-1300300232020012-1112333332103111-0312101113113013-3111023102322322-2111112320331133-2000132022001002-1011202312100220-0020111123130320"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes`

- [routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221): complete subsection reference.

<a id="canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-3133321110130321-0010132323232201-2220033300100302-3003312330013032-0110300021300013-3210213112022021-1203301133022020-0301021203032010"></a>

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0033322303211323-0120000003302210-2313311112131223-1133132301131303-1201230230303230-0133302110330201-2332023230311330-0320030321011312"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes`

- [custom_route_object](data-sources--workload--reference--group-018.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-019.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110): complete subsection reference.

<a id="canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-2330132123222013-0210003011200332-1321102010020303-1123312033030232-3023033112303110-0303312213321332-0012332313002011-1233332021111233"></a>

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

<a id="canonical-0001302120013112-3222210110100333-3033201322202103-3221321012213100-1110112222100202-2213211220213122-3331312002312213-3121101310333122"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object`

- [caching_disable](data-sources--workload--reference--group-018.md#canonical-0231321100133302-3222022021103000-3322113023221133-0112321031003210-1003000003133113-1320031223313021-1303000003313320-0103201223233223): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-018.md#canonical-3323201223312101-3102322021300131-0000221012003302-3002120133332002-3302032130033002-1203002022001011-0301203232330003-3132013023222201): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-018.md#canonical-3333110013011302-2000302332230032-2012313330021211-3211023012223321-1233010212213312-3133203011021211-3002000302202132-0313312131211201): complete subsection reference.

<a id="canonical-0231321100133302-3222022021103000-3322113023221133-0112321031003210-1003000003133113-1320031223313021-1303000003313320-0103201223233223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-018.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-0332323313303233-2321101330101103-2302113110322122-1100331100133132-3030133331313010-2322011133003302-1022100231222313-0330333012312031"></a>

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

<a id="canonical-3323201223312101-3102322021300131-0000221012003302-3002120133332002-3302032130033002-1203002022001011-0301203232330003-3132013023222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-018.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-0220223233220133-2010302020203330-3001210101202010-0011132320130211-3030312121313311-2112311133223022-0011021223023012-0110303021223222"></a>

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

<a id="canonical-3333110013011302-2000302332230032-2012313330021211-3211023012223321-1233010212213312-3133203011021211-3002000302202132-0313312131211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-018.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-0221031130030121-1320332312033121-0200303211303031-3233212313122310-3131133233102330-0131101202212112-1311101013332221-0203333232102221"></a>

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

<a id="canonical-2211012200023313-2001321320020032-0100131322103323-0332333330320113-1231021223111203-3231010121322111-1221313203300333-3010220020130110"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref`

<a id="canonical-1320331321200103-1003233201211330-3300220012101220-2311123032320310-2103201000221322-1312311330102313-3120232213221032-1330123112331110"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.name` property

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

<a id="canonical-0122011313200201-3233231121122302-0200120223303212-1012121132023211-3002321020222113-2201323130303100-0132232132100032-1120033203300233"></a>

<a id="canonical-2133130102332330-2113010123301003-1131222103303122-0220223000021321-3322103303032321-3322033233223321-2201133231133003-2310212000223210"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.namespace` property

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

<a id="canonical-2031302113123032-3121101230323030-1032023132212210-1200012201113223-1200212102323231-2012313131020003-3102211311210332-1122101001103313"></a>

<a id="canonical-1120300100030220-3123120210000213-1301112133333033-2300323111121022-1210330331303102-3112312211200133-1102110320320121-2003110111002011"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.tenant` property

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

<a id="canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-0332100320202223-1022122011230302-2023021203001302-0321212232223302-2213220013012332-0210002132323133-1200020102010101-2103121121031031"></a>

Type: `"single"`. Computed.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303011103032331-2233331331232023-3323301200022103-0001021330202303-3233022023102000-0011023323313321-1213030212223030-2023233202031111"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route`

- [headers](data-sources--workload--reference--group-018.md#canonical-0120001310102110-2113233120002232-1022213310310223-2222321113203010-3311103131311100-0032211221103113-3003221231201111-0233022211223111): complete subsection reference.

<a id="canonical-0311231323123223-2022121231022211-1310002023310013-2230123321111030-3110112101002202-0330123032020003-3232311201303031-2223212131133210"></a>

<a id="canonical-1021111003120121-2200100201211220-1122122121101002-1111210301230323-3110132222021111-3230331301010121-0112103021010000-3312101120030220"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--workload--reference--group-018.md#canonical-1010201103210000-2333011102020132-1011201330223111-2111102132002111-1233332002120031-3222003030012112-2122220113020120-1133302200212233): complete subsection reference.

- [path](data-sources--workload--reference--group-018.md#canonical-2321230331023311-0102000232321322-2322222002011101-2210001011001321-0202232100331122-2120132200210321-1211202301330022-1220231203122002): complete subsection reference.

- [route_direct_response](data-sources--workload--reference--group-018.md#canonical-2232320320230221-1213000311000001-3221313321031131-0001301213323033-3212323303132232-1302003031302322-0112232230022220-3110222331122322): complete subsection reference.

<a id="canonical-0120001310102110-2113233120002232-1022213310310223-2222321113203010-3311103131311100-0032211221103113-3003221231201111-0233022211223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-3133113322302130-0231002202102130-0131333100101130-3020111223103033-3103330010233003-2131221120303200-1321203031331130-1333300101023213"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 0,
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

<a id="canonical-2212310222303113-2313331320300222-1310231322130221-0212212213022222-3202012112023222-3020100132013100-0302323231030110-0203312121131112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers`

<a id="canonical-0031201011213211-2032312321230200-0222010210303312-3133200332233110-3223222022113322-0010303012311231-2011330212303012-0002323102310211"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3212332022300011-3133012221120332-1221213120302211-3210320010230002-3323202032121113-1321113112033331-0230003233312203-3301021002120322"></a>

<a id="canonical-3121221100000103-3033233023122111-0002022300231112-0102203012230120-0222123101210001-1030113003101321-2300222001213333-2031212333123320"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321221221111032-1123123111201311-3313111210131110-2002332213233332-0021323312332223-3213230023221321-1333131033303212-1330100130331022"></a>

<a id="canonical-0310232312101132-3311231130013302-3023013001311113-0202030013013300-1013132320223020-0033020002201022-2001220300021011-1013333221121100"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3111023313132001-0032303122211022-0102303032003030-1300120201113323-3021101121021232-1110212022330110-1320010331112212-0310223132233213"></a>

<a id="canonical-0130303032221303-1300321003210210-2211032112222300-1331130133331111-3200113210132320-3030003221013230-1032203212323201-3130321100310323"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1110333000302300-0100012202311302-2201313311032032-2303233122332302-1303133321110303-0101213233033002-2232333221302223-2021023102102322"></a>

<a id="canonical-3013301220201300-2311310300102212-0320202223022011-0032233210033312-2002212213220330-0130212212031012-3000213333200011-2232123103033210"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1010201103210000-2333011102020132-1011201330223111-2111102132002111-1233332002120031-3222003030012112-2122220113020120-1133302200212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-1011003001023332-1230220131112013-2021203312201132-1323233121011213-2100130311033102-0222232013323311-0200131312303312-0233233212013333"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-3022110132201003-3132320021222102-0310213000023111-2011022033030313-0122330203333310-1032213103130331-1330013122131330-0111323232301101"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-018.md#canonical-3022312110023233-1332320232000001-2103333323331001-2203230121231030-3033333232002003-3023131302103022-1002312222221103-0123032320030023): complete subsection reference.

<a id="canonical-1121220131221110-2100210233313313-0003032122101230-2112031233003112-3232032332123000-2302130200201330-2330112202211221-0230121331311201"></a>

<a id="canonical-3030120320230311-1130303313303022-3332030303100200-1302301212031313-3313101003000013-3121121112033110-3100012301303233-1012102010111330"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-2233012000111011-3111202202032232-0330202133223313-3001221320003222-3132021300320330-1233320231110202-0120131023331132-0322131230333010"></a>

<a id="canonical-0321220230231302-0311131131100130-2223031303100001-2232200032323310-0221213321233200-0223102123031313-3333110230232121-2131003201313333"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-3022312110023233-1332320232000001-2103333323331001-2203230121231030-3033333232002003-3023131302103022-1002312222221103-0123032320030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-018.md#canonical-1010201103210000-2333011102020132-1011201330223111-2111102132002111-1233332002120031-3222003030012112-2122220113020120-1133302200212233)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-3100303322020230-1020021031321300-0022133211313202-0230021023022112-2123222131221031-1102033133120220-2121020023311313-3120120232332023"></a>

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

<a id="canonical-2321230331023311-0102000232321322-2322222002011101-2210001011001321-0202232100331122-2120132200210321-1211202301330022-1220231203122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-3210302120221232-0020013000010333-3310332031303312-1220203131311303-0200020232202103-1010112003210310-3010313320111120-2101202221122232"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-3113231110313201-2112112113300223-3333303023121101-2330232302230103-2001332131113231-2100111121230233-3110001123021202-0200120021323111"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path`

<a id="canonical-0323210210111311-0323130020320333-1133223330131030-1031000022103313-2301000203010112-0330320032001210-2310020320022213-1212303223023301"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0202201111102213-3111031222223102-1222203210331321-3110321223002222-0120233100111022-2222212111030231-0300023002111312-0013333021222130"></a>

<a id="canonical-1211331331110301-0013201301011313-3200122231303202-2101221100303031-2013002332302013-1331322220133111-1121202301120333-0213331203233030"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3203010331121221-3301332100023201-3323223101220223-1103330232231032-3202101010313313-3210221301000021-2212122110013313-0122330311120022"></a>

<a id="canonical-2132203123323102-3233020021202021-0322231310311030-2210323212321113-0211113122223211-2223121031122012-0032330312100113-1222212301210013"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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

<a id="canonical-2232320320230221-1213000311000001-3221313321031131-0001301213323033-3212323303132232-1302003031302322-0112232230022220-3110222331122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-018.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-2333312300120132-3231032321033021-3333122300213330-1232132322233031-1303231002011122-0101020000013211-0303310123321010-2103031033200010"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312103221123032-0202211030111200-1100210023231201-2012202210113111-3012212230121212-1111120332322131-2023130122101023-0221312031023013"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response`

<a id="canonical-0230123112012033-2011310222103303-0320021101103302-0132020201031303-1121113133133223-1201120212233220-3321211102033210-3220213331000022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0102200133313031-2133030111202222-0103323011113213-3301210322231120-3021110313231033-0312332101010231-0010002130001210-2230201001310110"></a>

<a id="canonical-0123221033230110-2010023112200211-2310133013221200-3222131300002003-2213020010212122-2123131302211001-1003031311023013-1331330322130112"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Computed.

Response Code. Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-2321223312133133-3031122113013300-1233020223313102-2332021023210113-1220013332111331-3221133001222210-3122201301323100-1020111033323333"></a>

Type: `"single"`. Computed.

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131200031323131-1232032233100130-2011012313230011-3202202031032012-3300231123103223-1303320102121302-1223211023023320-1333221303031101"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route`

- [headers](data-sources--workload--reference--group-019.md#canonical-1333302203002320-1331201010032332-2111120231210121-1321233130101102-0201303200002231-1131113212132310-2002330002132100-1020133012202321): complete subsection reference.

<a id="canonical-2203302222300130-0203003001131120-3122121101011301-2313232032131210-3111300130001102-1031300112202111-3220310010022133-1213011112321010"></a>
