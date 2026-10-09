---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0110222220210032-0331331320002032-0330112030312222-1132120120111312-3113131202131300-0331013230113132-2112213320303132-3220320330033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0031303213213230-2003331031200220-2032201120012222-1300113230130211-3321122000021131-1133222112012330-2331302010011132-3122231122103321"></a>

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

<a id="canonical-2130120202011330-1111032013232013-0000213203333101-0312320220003112-2303331313110110-2130322012313213-2203113033203120-3223313321213212"></a>

### Direct properties for `proxy_config.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-2100012232300311-1022312321330200-0320113031110201-2331220210122300-1332223301233032-3220331312020101-3010122131101022-2001023212020100"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-2131110222323031-1310312332103231-0003203223032310-0211103112012003-0230231030321332-3330232003300312-3200302001020032-0032222000110130"></a>

<a id="canonical-0231200201321311-3121223232011303-2112320323130123-1113303302333112-0101333331320133-3233032012210033-1303330222222023-1001333111110331"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-2120223232223003-2030201231223003-0131120133022011-3333131302213011-0322333323333033-0320321203031003-1311202231002232-1212222322102011"></a>

<a id="canonical-2113102122333012-1212121122003001-1321000011210303-0332010100123313-1233000213131220-1302011203000321-1120232231023212-2011302012321233"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-0022301131111233-1321032023222011-3000222310010121-0203121211232323-0031002300112021-1032033202203200-0222013023000101-3113313232330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-3013330031011133-0031101232332000-3320013320321210-2022112121323010-2200333123112331-1312301032101133-0212012003303311-0031300122010323"></a>

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

<a id="canonical-1320200110221000-1031112223132333-3133121330331232-0223310122222011-3103231101200312-1213130230303300-2102201100233313-1121102332232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-2010121312320302-0101102001333221-0103032311000230-0221230332312300-0033111103133232-2123120310213121-2002231312021323-3220230230033103"></a>

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

<a id="canonical-3121031223203211-3002232101112320-0213210011001032-1232303220311202-0112030222012333-3003322220101030-3022330210023031-3232211130321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2322331100001201-2230223332010032-2110202221222210-3220010010231032-0201023110212123-3231021120303112-1132012210021202-1231003202000130"></a>

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

<a id="canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-2002303301031111-2132312201122033-3220010121322013-0333031200101322-0203103201332033-3030000210201212-3320203123132031-1220221212230001"></a>

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

<a id="canonical-1321330011113121-2111132000122133-1220013123303230-3023112133021320-3212102310103110-0002113111212113-0201213003223021-3222002011112003"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls`

<a id="canonical-3013302213222033-3023202323332200-2003131122120220-3031201022201333-1231332233300133-1310320201223312-0312031001032213-1303303333013133"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-2001103200102102-3123312201111013-0233221112320031-2110311020310222-3103322321302002-1300320230202232-3003301332010033-3032110010131020): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-2303332100000202-0220310301012200-0120330311213201-3112033222320313-2000121202023110-1030303133332302-3222301013103021-0102232123300300): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-2302133000120033-0002333202301303-0103122120311021-0223300032001222-3010031112020220-1001001000233201-2321230203001113-3133020333310113): complete subsection reference.

<a id="canonical-0003220231231303-3130033332303332-3322310321001232-3020013321023310-0203000303200202-3302312200201110-2210302301030312-0122303302220230"></a>

<a id="canonical-1133011000132210-0112022122311133-1021233300023111-0332100013122212-1101302300303222-3222012203033121-0330031220121031-0320220302003301"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-3211103101231003-3022300000021000-0232311030103131-0132120231222232-2022310023101321-1131323011110101-2110320321212130-1213311133112230): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221): complete subsection reference.

<a id="canonical-2001103200102102-3123312201111013-0233221112320031-2110311020310222-3103322321302002-1300320230202232-3003301332010033-3032110010131020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-1100022302103132-2021302201320123-0122000012202212-3221132210021322-1202201031101030-1302010230302232-0323200313222132-3200011223123233"></a>

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

<a id="canonical-2332032033012021-0011300113012000-3010133023011322-2010313213302122-0120130013111121-2331022211000130-1320213303123201-2303202102311030"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.crl`

<a id="canonical-3133102110030302-1230311303021030-0313313000311230-3013300130031221-0000223202221333-1310012302032123-3332202133031212-3022312213100312"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-0321032210002011-1013222001011202-1020121230310020-2311331030000210-3332213103312212-1120303300333012-2213031203030102-2312232330311130"></a>

<a id="canonical-2203310010003120-0023203103001332-3110311330302233-2112023233032031-0230113301213001-0331101123312202-0332212321200010-3002023121111130"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-2313032301310211-1030022122233213-3220202310322302-1300012333123332-3321222110230200-2212211102222123-1011013002020001-2333112131210302"></a>

<a id="canonical-0202023030300001-1312011132021002-0131022221232103-3100210230112131-1102331100301011-3223030113322022-2222231111110133-0133202213102323"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-2303332100000202-0220310301012200-0120330311213201-3112033222320313-2000121202023110-1030303133332302-3222301013103021-0102232123300300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3020011320103021-1330123313300030-0012020133130110-0332032322023001-0332300202022133-3201101230021303-0211202030211323-1231220220130221"></a>

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

<a id="canonical-2302133000120033-0002333202301303-0103122120311021-0223300032001222-3010031112020220-1001001000233201-2321230203001113-3133020333310113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-2000232132011120-2011013201322020-1031122031302333-1302021203011212-2020221032312302-1331000203010313-1000222011111201-0103233332312232"></a>

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

<a id="canonical-2010323121213301-2101122211030313-2213013121220221-1101033303333220-1300012323221211-3123222031012313-0301131001103220-3200113022223113"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-1003202213002313-0031002130223112-1013121001110312-3200311020323310-3102020112003321-2102022112310031-2132332310111321-0201233031302321"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-1320010130103133-0020223220303031-3231323210000230-3023233331131231-0010223232111330-2122110013223310-2223032320210023-2302223020010120"></a>

<a id="canonical-1322101232101103-2201101233222320-3300000202220210-0212100232102123-2121023202212223-0103303313213021-1111003233232120-1312203302233203"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-2122330320130303-0011132332121021-0203120133010320-1101300202120102-1102330010212331-2012033213121100-2032031012321023-3132131000303301"></a>

<a id="canonical-3221232112013123-3303120221220100-0332203120210321-1111030030313313-3300210203032212-1303001203013031-1103102113233011-0020311311322232"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3211103101231003-3022300000021000-0232311030103131-0132120231222232-2022310023101321-1131323011110101-2110320321212130-1213311133112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0111033302023221-3310003303303131-1310332121313111-1013031100131132-0130010033111122-1113222120323210-3200300013123322-3203323313213131"></a>

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

<a id="canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0332212230300113-2113001211110101-1300133321321132-3203023313200121-3101231200331111-2013121011120303-2311113123200033-1003222322321120"></a>

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

<a id="canonical-3031221300210302-3320300311312002-2000232000103122-1031021101123002-1102133200211321-3221330001013030-2000202003311330-0211230310110202"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-2212233001102010-0302012322301022-1331112010300111-0102001013012231-2301313011311012-3012201000313230-0123301213020113-2001130100331302"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.tls_parameters

<a id="canonical-0120102120100020-1100011221033102-0221220220332001-1333311021222120-2023123132110312-0002031111321112-1022231022201102-3131022230333133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230032303200213-1000103312100211-3302031122311020-0000120222310223-1010230210213222-3232123300121003-0200131223211133-1123212010201221"></a>

### Direct properties for `proxy_config.https.tls_parameters`

- [no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2230003003322321-2030203023030130-0311103033230111-2312121003013022-0122112230330223-3233101232232332-3200133323000330-0231021223022121): complete subsection reference.

- [tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300): complete subsection reference.

<a id="canonical-2230003003322321-2030203023030130-0311103033230111-2312121003013022-0122112230330223-3233101232232332-3200133323000330-0231021223022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-3310030003230002-1220330232102210-3310322322222322-3120220231131103-2013100022001031-0232220231320012-3222200121113102-2231011021001022"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-0220301032330333-1323013100331100-0012033233323112-2320332311233123-3212210220220222-0112021100023210-3201100302313332-0122102311013310"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213310222221312-2223232133310030-0212013001010220-1230322300322223-2333032121100131-2310112100011201-1203100100003211-3322223132031210"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates`

- [blindfold](resources--bigip_http_proxy--reference--group-004.md#canonical-2000123212320210-0120231210213311-1301001130331122-2131212100000131-3310210233011013-0222212231121303-0213331230010110-1233103013231331): complete subsection reference.

<a id="canonical-0032213100120331-0231021203102030-3023322201023331-1112010001322213-3132232010200100-0320222200322132-3211320213133323-0013323130333301"></a>

<a id="canonical-0022032132200233-1122220202011211-0002011021113113-2113200222003311-2131331021101020-0323300200113033-2023211232032321-2031000113020330"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](resources--bigip_http_proxy--reference--group-004.md#canonical-2003101000123301-3003321201003100-0032122001302231-1022322320210120-2213101123103231-2003102111322102-3223300103032332-2111300000213013): complete subsection reference.

<a id="canonical-3133310331011323-0312013011321130-3133332322303112-2232201032230200-0132020301023012-0013213132203202-2232223120002310-3221330123213032"></a>

<a id="canonical-2303002103013011-1333020333010222-0121201021030220-3121122010000121-3313000013311030-2020120213303230-1003313120223302-2332030003132123"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-004.md#canonical-3110232021220003-3003233101313123-2203211130303311-2032120100312000-1320013021303302-0203201310100032-0221200113301310-1100013033233030): complete subsection reference.

- [private_key](resources--bigip_http_proxy--reference--group-004.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302): complete subsection reference.

- [use_system_defaults](resources--bigip_http_proxy--reference--group-004.md#canonical-1032003223022220-2130323332313130-3221323000322303-3110310230102211-0210333013303123-0310301121230130-3112223130103113-0332222220001330): complete subsection reference.

<a id="canonical-2000123212320210-0120231210213311-1301001130331122-2131212100000131-3310210233011013-0222212231121303-0213331230010110-1233103013231331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-3020130232100132-2121311131231100-2303331223310233-2021330023311232-3311321321210301-0231321322301010-2010112002300123-3320233132301330"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-2120213133100100-0121323012131111-3212223133232300-3221322133001331-0030110232030221-0210030131103100-2313222000212331-3200203022210113"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-3201120310110031-0210023000223211-0230121313021133-1302202302033130-0000123320132212-0102101123212223-0300213320011110-0303222220220223"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2121012022202020-3121202310203022-1110301001231313-3033322212020230-1302102231133313-2013331132002102-1233023022012322-1330133013230330"></a>

<a id="canonical-3332333200330321-3031202010022032-0332213102303331-2020031103223030-2203012323001302-2110000130321231-3013333110313000-2023003020102232"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-3000012023331030-2212223333100033-0103231123233132-3230221131011311-1103103220232313-1230202112120303-3301033130121311-1312200013232000"></a>

<a id="canonical-2231213330120333-1021022033202230-2120121303332213-1312003302001120-2113300023311013-0302010121013032-0310321113132312-2033300023301211"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-2011101021220000-2010223311130321-2323032013321020-2221123232020030-2303113212200212-3111132130023133-0321321233133221-1330232332103211"></a>

<a id="canonical-1122101221210233-0022220032313222-2312311003313332-1333031011331322-1130201212323022-1330101013133133-3313201300031001-0103001023300023"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0302203312230112-2102323023333330-2132010331130110-3002113112111102-2233213001011111-2101022310031302-2332131221023020-3130110300311121"></a>

<a id="canonical-2021031130230131-2011222222122032-0220020302101001-1113111013212121-3120123223330321-3221302320012222-2100122110313302-2120102103233301"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-2111031303010321-1033033122012221-3102102221112302-3121310311111022-3321302231120210-3011223331200330-3112310230201010-0031122322332300"></a>

<a id="canonical-3022312032132320-2033300121133001-0121231020031013-3213033111102030-1002032312112120-0230311212101000-0202033013020202-3120301322000020"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-2022102211302332-0111223333333330-1233333100200110-3321322130333030-0201220120233002-1321010231131022-1332301312210120-2111010211202131"></a>

<a id="canonical-0022312000000332-0302120130032233-3301301232010113-1130301120332323-1223132311010303-3012301200230101-2213033320201300-1112103220010012"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1133311203100222-3211002132031303-0213332020010130-0312122013033000-0110120200312211-0030330312003130-3230120322210002-0002102303333030"></a>

<a id="canonical-1220033120133303-1232120023133032-1131002303302023-0121323221222301-3111131033201201-2132133033013112-3102031322333310-3311023222232312"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-2232233023303220-1230230212111022-1100131011223212-1202030302131231-0310110210310130-3323210100221231-0230002311020021-3020211331211223"></a>

<a id="canonical-2322303333100221-3220210301011323-3000130201302303-1102020011133232-1203220323310111-0213332103121003-3313313321011201-3332203202010121"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3330112010023003-3110333133231222-0320123320032231-1011111312213223-2032123032312033-3133200303332303-2001003121001122-2020332300202100"></a>

<a id="canonical-2010330031312303-2321013130031011-0013010233133212-1201332023211103-3330231210303033-3330330133132102-3321031010211031-3110133333100320"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-1001203031031013-2032122321233320-0133221333203310-2322033232323311-2213112130103202-2113022133200320-2311230332230100-1032330110033202"></a>

<a id="canonical-0332200023220200-2221132021223002-2322322231130120-0130112011001110-0111122133023323-1220121013102202-2201033002001311-3133212323010103"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-2103310312022200-3231212030113331-1011312312331321-0210123013003031-2013323311201111-1102212020313132-3200331120320001-3333032311210212"></a>

<a id="canonical-2210133010232232-0202103210212312-3010322111031311-0033000113121220-3302001331122321-1231331000301220-1100010030212331-0120321200331212"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2102101120032131-3300001022101311-1032123021221311-2210002231112203-1313202301212201-2001231120203213-2032110133330323-0321301003022221"></a>

<a id="canonical-1133313302131103-2320002123331200-0023131232011201-0001032102311310-2331123132300232-3310232221300120-0302012322311130-3230312102133300"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-3333303103313120-2022110023223220-1002023003032121-2102123202021123-0211002202111232-0033121312233332-0132302130011221-0132133032131322"></a>

<a id="canonical-2223302032212223-2212223122122210-2302032312001112-2311201110132320-1021133120222331-2003133133312021-3213120300122300-0302030002333223"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2320213201313033-3212132120032132-1320323123233022-3000101303330331-0333002133203032-1102103113320223-1123330112230121-1223111201013021"></a>

<a id="canonical-1220202000320200-2223000003120333-0013331230213001-3002002123220121-1333301312121102-1222303022323220-1230123021011101-0030030033131323"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-0330322030010001-3130112010000103-0001003132123111-1230103021123232-2313023210303130-1032332230113230-3101202213212011-1200131331102132"></a>

<a id="canonical-1003031231301032-0222033221333323-1201003331021022-0111222210013012-2200222300112202-1220101133122212-3311223232312033-3202230131300213"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0310213033121000-1320102213130013-1101311120123110-2331201220212003-0200003102211221-1333012311001310-3302132213233113-1033322223333111"></a>

<a id="canonical-0003320001021110-3133130111110101-0330223031020010-0300000131103123-1023330200101121-3312131113130121-0020222332201011-3320031010223000"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-3102113110221113-2131031113311312-3033220113102020-3300301211133112-0221323233033000-2332310020212021-2203333331123332-3012301323013311"></a>

<a id="canonical-3022300121112312-3231321013333212-3210233210122231-0332332202231121-2003012320032212-3112010002011013-0033201210210010-2210033001210130"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3222221130211213-2033210201313211-1002231321322230-3011112103323220-3001031222131131-3101032201010012-1222130100133320-2121002113022113"></a>

<a id="canonical-2321010003021332-2322122222103313-3113213111033001-3301331100122311-1213322022102033-3230001301022000-2333313111012112-0013110211112131"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-2003101000123301-3003321201003100-0032122001302231-1022322320210120-2213101123103231-2003102111322102-3223300103032332-2111300000213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2333232103133300-0103332332102100-3233300330020313-0131011021330123-1023112210330310-3321023123023210-2132112123212203-3313221233102030"></a>

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

<a id="canonical-0213111322030333-0311003223220322-0023323212031200-0330303313223031-1003330130211311-2223200310012311-2100333022302300-2110121312211210"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-1100222003230221-2033331123031030-3220021020133123-3012122121031213-3002223122312312-1011012231100202-2131011223020320-2331303222311231"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-3110232021220003-3003233101313123-2203211130303311-2032120100312000-1320013021303302-0203201310100032-0221200113301310-1100013033233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3233023212023113-3300203231011311-2232232201223103-2311313202200030-1313313102330213-3112331101121320-0330033210230300-0312131322102203"></a>

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

<a id="canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3013133203032233-3112011013132000-0201232133010111-3321300130121111-0211113323130333-0223111102210233-1321012113010101-2100001101012221"></a>

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

<a id="canonical-3121330001322003-0111231120333322-0201130230322330-1201010321111110-3233333232110133-3000111011320023-3131200222230021-1032020000230323"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--bigip_http_proxy--reference--group-004.md#canonical-0222132222003310-0221302202120120-1130032201222333-2022311310103320-1131322012101010-0120331202202020-2221321101320131-2022111211000000): complete subsection reference.

- [clear_secret_info](resources--bigip_http_proxy--reference--group-004.md#canonical-0003112113212000-3122200022201033-3113113103003111-1230203021201222-1323302110011213-0230320311200211-0301200022221103-0010221222311120): complete subsection reference.

<a id="canonical-0222132222003310-0221302202120120-1130032201222333-2022311310103320-1131322012101010-0120331202202020-2221321101320131-2022111211000000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-004.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3002230331121022-2011220313320330-2221223233211031-3212133202023210-0003333232030203-1221312322100203-2110002312132312-1303232011130120"></a>

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

<a id="canonical-1212031320112230-1021300130001110-3020002030120211-3111232331030010-1033321112301312-2311311321003313-1223210031332312-2010332011132233"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0212203223200200-2113132031320133-2132003302301211-2032130301011120-3210003130222203-0312202003201033-3223002230220022-0232200033113102"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0121123033113033-1022122331121000-2300103130203112-0213310122131231-2231020003332220-3003231111321113-3232112332320231-0211321310022301"></a>

<a id="canonical-2030212313321100-3210130113313332-3220310113102003-2211132332030010-2332011313232223-1002213021201002-1312210123223200-1332323123120130"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-0211002203112010-1101011032112213-0113213212112323-0013311030333312-2233033312212201-0211312122001100-2302320332300003-2100233231323110"></a>

<a id="canonical-2201202300213011-0222210030032311-1110130130031011-0123220222303320-0113312212013321-3331021210132010-2111222032001322-0200031312112323"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0003112113212000-3122200022201033-3113113103003111-1230203021201222-1323302110011213-0230320311200211-0301200022221103-0010221222311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-004.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3101321102031322-2010221022033033-3110011233010102-1132312122322000-1121203302023131-3301130212231331-0313000010132021-1200121313012321"></a>

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

<a id="canonical-3212020033303202-1031111322220303-3131132003230223-1303022000021300-0022233222123012-0002333231331230-2222101321020023-2312023222011230"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0110203200231022-2023013001221300-2203303231020311-2032020312220302-1203232200100002-2030101223302013-3011021003132223-2201332211032312"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2023202121220333-3310223010223223-3323320130200213-2022103011011220-3022133301303232-2130320133302202-0031110032320232-3230323032122111"></a>

<a id="canonical-0003111033323303-0103322012303131-3220302112002231-3101300332320013-1201330302201031-3220001211333220-0031012101130132-3300202333010200"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-1032003223022220-2130323332313130-3221323000322303-3110310230102211-0210333013303123-0310301121230130-3112223130103113-0332222220001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3300203102333201-1311112211011302-0031000102110232-2312210220211233-1002211132111022-1232202302023033-2232132100310132-0111012201021031"></a>

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

<a id="canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-2103323030031233-2211133023111033-2132023100322132-2331320220012233-0223223032201022-1301023100321302-1102000013211332-3302012003202223"></a>

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

<a id="canonical-0102330111321010-0302230200222201-2330010030013220-3111200113122110-1230333200303121-2013100213131101-3132201021232020-0003123123023313"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_config`

- [custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1033102301120112-0111120023211201-2100211300330213-3010031033010330-1111301311201012-2102300132010312-0330332112033023-1120031100331013): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1001302030122213-2212331100230223-3101020120201233-0322202103201113-1332020112312000-0001323311112233-3320200112101330-3023120132020100): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-3330310303203023-2110013312120302-1231122110200313-3022003222230031-3032031332003231-2210220220231220-3121112121023001-2033132121322102): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1030031031103101-3101301320202011-3031130233302300-3220211202033301-0120321113202212-0123231033333202-3202233332313211-0123311333021001): complete subsection reference.

<a id="canonical-1033102301120112-0111120023211201-2100211300330213-3010031033010330-1111301311201012-2102300132010312-0330332112033023-1120031100331013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-2210003101010303-3213302030201200-1223323223330133-1101110322332312-1322302203210230-1310233210130310-3032112312000001-1303200130311330"></a>

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

<a id="canonical-2201122223103303-2122010320303321-1231311023231132-1001002113030011-2112132000131223-0202210331002230-1302130303132133-3131003022020010"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_config.custom_security`

<a id="canonical-3323033112233302-2102233202022010-2031312223330323-0322112102033321-1130013033033131-0200122333003312-2310031023302223-1101200201111303"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-2233022230022031-3022210133310300-3311001103120332-0311130301100121-0200020103102300-2210011200302101-1013030320212223-3320200201111120"></a>

<a id="canonical-0200030122301331-1002013321301033-3222020232112203-3013222301232310-3230222000202121-1020220313320101-3203230103321133-3100131023233131"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-0101101111032333-0230110103311002-1210133112302333-2000033231011000-0201021000220002-0303311022232002-0330031210020100-1210331230020100"></a>

<a id="canonical-0102010003003330-0322300223111213-1130301320032110-2232022112332331-0201111120310203-3223321301131030-3223133233202210-3121123222033031"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-1001302030122213-2212331100230223-3101020120201233-0322202103201113-1332020112312000-0001323311112233-3320200112101330-3023120132020100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-1022333100332033-2202313233010321-0033022011010231-0011021211031123-3002112000013110-0333322103010110-2112121010233130-3100123132023301"></a>

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

<a id="canonical-3330310303203023-2110013312120302-1231122110200313-3022003222230031-3032031332003231-2210220220231220-3121112121023001-2033132121322102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-0212100230111030-1120211311122313-0130021100300301-0133232230003123-3130133203032303-2111020323111011-0311231333021201-2012120331102000"></a>

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

<a id="canonical-1030031031103101-3101301320202011-3031130233302300-3220211202033301-0120321113202212-0123231033333202-3202233332313211-0123311333021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-0323133222230223-3132333012313201-0310113302022122-1210022232321131-3031201022012310-3330300002021201-0013300232103323-1331110032131202"></a>

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

<a id="canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.use_mtls

<a id="canonical-2202303331113300-3101033121011310-0211110012020211-3202333122112301-0023232301101203-0111110122200310-0230301102120103-3200322322032202"></a>

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

<a id="canonical-0002013303000031-2112303311032101-1131123210000010-0030302131311200-3233113221101032-1102003031122302-1313023323332021-3332123031330211"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls`

<a id="canonical-1323201330232021-1212100312030110-0020020201022100-0020210302330120-1310203330332002-3122112003211233-2003210131011302-2300010113202331"></a>

#### `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3301132013002133-3300022203123130-0133211010212100-2012303233130300-2022332111110131-0212021230201221-2030321322130002-3031203021201132): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1122000200022122-3121300310203023-1223122012113203-2233112113122201-2232013203110223-1212001333132023-0231003003331202-3323321002130101): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1010131132012330-3113302320201301-2211330221201003-0321031133033121-2021333131022213-2301330203332311-2112330113100011-0132212321211000): complete subsection reference.

<a id="canonical-1321220310111301-1003313120032121-1213103103013301-0223300031223212-1122300033032212-2321330100301100-1131023101100221-2310112221020200"></a>

<a id="canonical-3023212313113322-0221022322032333-2300012030210113-3213313120313100-0231332032311200-2012302220113111-0110210323030132-1121310301022232"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-0111131122320313-3201233031331303-3100132001000021-0213312100201200-3120223313203300-0031101330100012-3210112310010210-3022011031010300): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1130033202003022-3030302000113122-1213132121121231-2013222200323313-3322021011131110-0232030100121233-2332332322103310-1330330112321030): complete subsection reference.

<a id="canonical-3301132013002133-3300022203123130-0133211010212100-2012303233130300-2022332111110131-0212021230201221-2030321322130002-3031203021201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-0123233322023122-3310221222012200-2020023031002133-1123201333210201-2322013300203211-3321232331223333-0223313223201132-2332302211210222"></a>

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

<a id="canonical-3102313303032331-2332302000300300-0311101112201103-3031203030011223-3203200103123101-0121013221131233-1033222203313023-2113233310331201"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.crl`

<a id="canonical-0012120003132232-3000121133223031-0022202203213133-2331133320002333-0100222012323013-1220302110223101-2123123130303213-0220030133111102"></a>

#### `proxy_config.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-3122230101301020-1113132300300211-3302132230010101-3213312010323232-2302300133231223-0330121311303310-0232031121322301-3322330033313312"></a>

<a id="canonical-2122112220030132-2103331312033331-2121301202031301-0101200123331302-0013122220000110-0023112002023213-0203102011302001-1121001212212010"></a>

#### `proxy_config.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-1303233313232301-2303300000020320-0133010030102333-3221033112102203-3013323231001321-3111033020213000-0110032201013323-1123020210220221"></a>

<a id="canonical-0013311123332021-1111201112212130-3102030100230321-0033311012313211-1231002233202231-1000313030020221-1203021021321131-1211100300001212"></a>

#### `proxy_config.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-1122000200022122-3121300310203023-1223122012113203-2233112113122201-2232013203110223-1212001333132023-0231003003331202-3323321002130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-2101011211001301-0201311300323303-2201111013301002-2221031131030200-1010232101101003-1323003122033203-0300111203311003-3012012030231132"></a>

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

<a id="canonical-1010131132012330-3113302320201301-2211330221201003-0321031133033121-2021333131022213-2301330203332311-2112330113100011-0132212321211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1133110232210011-2201103002102002-3300133300101313-1300002131011010-1102313112002022-2010003113123022-2332030301132321-2301012231133302"></a>

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

<a id="canonical-0132103222033011-3321111113031332-2001232120011311-0103131320301302-2202321002132200-3222312222233013-1310302232003302-1133120121300301"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-0221323322330333-0100032031320121-2203002301212221-1033030131321231-2121120111202113-1232301010000213-0122131201311030-1102311320321331"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-3103003221321002-1313113211002222-2220311203312212-2302000230201321-3123232102020031-2003300300103223-1121313111101031-1110302211033012"></a>

<a id="canonical-3231022320322120-0130203003230333-3022023002310323-1321203031302130-0132232202130211-0310232110212322-2223223301002102-3332233330032323"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0022203003022302-1131002012121320-3330223230322021-0030111130112002-3210231032020032-0132313112100223-3020132132130123-3232032200333010"></a>

<a id="canonical-0010233323333031-1202001320130112-0120121112230232-0232323211200011-3333233301002232-0000031323312021-3022313331333012-2233113330012313"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-0111131122320313-3201233031331303-3100132001000021-0213312100201200-3120223313203300-0031101330100012-3210112310010210-3022011031010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2010022220303330-3000033201310110-2101011330223312-2021222033112200-1131330023322120-3123331023020320-0033223001032133-2302012110010220"></a>

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

<a id="canonical-1130033202003022-3030302000113122-1213132121121231-2013222200323313-3322021011131110-0232030100121233-2332332322103310-1330330112321030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2012203301003332-0100201131320312-2113020323131013-3301311113322132-1131210332321203-1232030201233331-3232113310333210-3302121310212100"></a>

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

<a id="canonical-3120010001123120-1023012101323111-2123110213220331-1311202111213201-0100000022220132-1100321122001110-3233310003110011-0033222100100332"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-0231031333310200-2120233223033001-3110213031133110-0303131230132000-0100311020312221-3323122132112010-2023202000232221-3210120322201203"></a>

#### `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- proxy_config.https_auto_cert

<a id="canonical-1232130333122220-1122323101230313-3221320113210210-0220120212020321-3021313123110330-3103112331000021-3123222113203022-3022132121330232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012013233010312-2232112322333030-0300110221130123-0023210013011221-3223303113121030-0020200122303331-0001020310301301-3000211313302200"></a>

### Direct properties for `proxy_config.https_auto_cert`

<a id="canonical-1021302330123221-1213000312303202-0300312012313312-3332123011132003-2221001220123113-0301032210010130-3200333222220301-2231021323213320"></a>

#### `proxy_config.https_auto_cert.add_hsts` property

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

<a id="canonical-2013312310300111-0210301002301331-2223112020303113-2033022333121003-0123122221332302-3110133321333031-3100130031313212-0200203022031101"></a>

<a id="canonical-0320201311131003-2301032032103110-0330112110333033-0222210133012212-1023213111033333-3032131302303132-1220311003320331-1021230131123013"></a>

#### `proxy_config.https_auto_cert.append_server_name` property

Type: `"string"`. Optional.

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

- [coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211): complete subsection reference.

<a id="canonical-0120223312103123-1032030101311003-0000120121302233-1030021320203203-3323202301120112-1300102022312220-0201130132113003-2222120310000023"></a>

<a id="canonical-2003012220202311-0011032312210232-1311113321212130-1223212103322132-2200021123310331-2232213003201212-3133001033000232-3310102311201112"></a>

#### `proxy_config.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional.

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

- [default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-3322322023021102-0330023200022121-3123222001202000-0321220120321102-1013300001032221-1101313203133033-2001230222102123-3302301022203300): complete subsection reference.

- [default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1322223303132201-3203203230103220-2203103203202010-1231131323001112-0002301130030320-3022300303222223-1322032020302311-3130100322312311): complete subsection reference.

- [disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-1300003311013210-1200230101211221-1120000300213310-2321300001022333-1103120100000210-3030332320232100-0221202203100303-2023221122230203): complete subsection reference.

- [enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-2301232330002123-1001022220103322-3220010023133233-0323231331233301-0211031322212011-2123223312310101-2332331333100330-2231131033301331): complete subsection reference.

- [http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021): complete subsection reference.

<a id="canonical-1222321203200333-3031011110100312-0111110013013200-0101110232001131-1011022313132201-2101012233030011-3030123020213222-0022013012320220"></a>

<a id="canonical-0211131320231002-3122032113032213-2111113022322103-0210102310100322-0321100313313002-3013003010211010-1231120012303001-2231103000233003"></a>

#### `proxy_config.https_auto_cert.http_redirect` property

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

- [no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2031001112333220-2323222300022033-3312203100203123-0312020013232333-3331202232200023-2221322031111020-1322133232110000-1001031112123320): complete subsection reference.

- [non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1032131103202223-1332201101321212-0121230133302101-1030012020120133-1103202333112220-2221232200013000-0113013330323023-0230322022203201): complete subsection reference.

- [pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-3310121301102211-2313201113000003-1230032122220332-2023020012232231-1012020102310102-3010221322312003-3121000111211112-2312300123110302): complete subsection reference.

<a id="canonical-2030023110120333-2011333201022013-3323133032330233-2010332011100312-0210033032233131-0203111033011300-0301312303313012-1033012210323011"></a>

<a id="canonical-1010003220122033-1032110002112332-3030120213132310-2031200123121013-1303302003030000-3303011220313223-3320131212220220-2130202300333220"></a>

#### `proxy_config.https_auto_cert.port` property

Type: `"number"`. Optional.

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

<a id="canonical-0131123230120033-1310101200201030-2020332122330312-2313231221221103-2131110020012100-2132220220312212-0300121330110210-0203323011000213"></a>

<a id="canonical-1121022200000231-1320222220023333-2233320333312320-1221012303230232-3312312201101303-2212221020330213-3200102221130233-0111212031122022"></a>

#### `proxy_config.https_auto_cert.port_ranges` property

Type: `"string"`. Optional.

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

<a id="canonical-3222011031201331-1011131323210012-2121202201131011-2321320301122232-0311200333112331-2031030113203221-2302013023323013-3012330012011031"></a>

<a id="canonical-0013333323311010-0030201333230030-2021123301110101-2222010033003222-1201211222012300-3333211303311203-1011323102022233-0032103031313023"></a>

#### `proxy_config.https_auto_cert.server_name` property

Type: `"string"`. Optional.

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

- [tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101): complete subsection reference.

<a id="canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-2023122032222133-0320312030031012-0233013300033023-1221200320321120-3133332302333212-1123032100210030-0130231233322030-1120012213013033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013023302323000-3101101331221111-3332222100103001-2233210313112012-3131222311302210-1102311030213010-3012031000322010-1001302023012213"></a>

### Direct properties for `proxy_config.https_auto_cert.coalescing_options`

- [default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-1030310312101112-1223221122331332-0000000223201012-2300322333002113-0132121320133123-1200301103122010-3000210101203323-0302320303122033): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-2130220103102031-2021000303010132-3311001130323031-1113323123311220-2232203010201212-3030123313100330-0200230311313100-3223010002330131): complete subsection reference.

<a id="canonical-1030310312101112-1223221122331332-0000000223201012-2300322333002113-0132121320133123-1200301103122010-3000210101203323-0302320303122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211)
- proxy_config.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1111301103121032-0212011113313202-2110301133301300-1131320230302300-1000211023121031-3213320000010211-1203312213133321-2233331101122311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130220103102031-2021000303010132-3311001130323031-1113323123311220-2232203010201212-3030123313100330-0200230311313100-3223010002330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211)
- proxy_config.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3010312210331113-1332031313112333-1203120212222333-3131123203123031-0232323022110332-0211303211300311-3112211022131311-0130113320003123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322322023021102-0330023200022121-3123222001202000-0321220120321102-1013300001032221-1101313203133033-2001230222102123-3302301022203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.default_header

<a id="canonical-3213133110002310-0013122112302013-3123322222020302-2330120220031131-0203132120210131-2333210023032210-3021313210022233-0320022322031101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322223303132201-3203203230103220-2203103203202010-1231131323001112-0002301130030320-3022300303222223-1322032020302311-3130100322312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-1012301213233321-3001220333032221-0220022203120031-0221020003130002-0233211221320112-3021201113020232-3213220112032232-2233231321200133"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300003311013210-1200230101211221-1120000300213310-2321300001022333-1103120100000210-3030332320232100-0221202203100303-2023221122230203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-3220013120110313-3313131231202332-3122133013011123-1203123220120330-2211332312100220-3332123303011222-0112032102312313-0033332013020122"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301232330002123-1001022220103322-3220010023133233-0323231331233301-0211031322212011-2123223312310101-2332331333100330-2231131033301331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-1220202012033120-1312202313313322-2123031100113031-3000313021200012-2202021021030323-0320321331001021-2021031122332232-1031220022013203"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.http_protocol_options

<a id="canonical-0222300021332000-2201002013200322-3323321302320123-0332321112311323-0021312230312301-2031312023032221-1003001303323233-2322030123112311"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013201233113002-0220111011230313-1020102300200131-0110221330313100-3222302111211022-2121002212203000-1120131203310323-1032310103121002"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-1101220321323200-3010303311332022-2200103113311123-2011012133023231-0200303133000030-2203021330022010-1332100211310313-1211231313030220): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-0303321301310212-0201133032133311-2022213011111322-2201031302231010-2230020003220030-3310032003331203-2203312230000201-1010230112312010): complete subsection reference.

<a id="canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3201111020112210-3030000310010212-0110333212112101-0211101201020130-1211113311123111-1201323311011213-1020020020312010-0313331002002101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231303122033131-3033211321333312-2001310330220123-2330320302030112-1011321101213303-0323020303330331-1233112202221331-3233302211232103"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302): complete subsection reference.

<a id="canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2312031202133103-1112123123021233-3011202231313100-0302223330200300-1030001122213322-0022201000001203-0211121110203221-2120120332110310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102131131101001-0031032310300313-0011302002311123-3023130311022311-3230101030032112-3013212121121100-1310202111331300-0033213212013200"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3100331120311313-1223133311013032-2131220121322000-2013233200232000-0021021102311333-0012211102300222-0200010211001120-2113110223032221): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3313103230222222-3030010330200133-3020130231320033-2300001023230013-2231201101123230-1220001123203211-1100312211101220-0113011213100002): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-2320331111000330-0001310020333213-3023100210221020-0222312121220100-2010231203130303-0032322313212312-1021211211020203-1331221301321001): complete subsection reference.

<a id="canonical-3100331120311313-1223133311013032-2131220121322000-2013233200232000-0021021102311333-0012211102300222-0200010211001120-2113110223032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0310211332313123-0312102111032001-2303223022211321-1323123311221112-1221220322101130-2211101003130120-1123212301312132-2000001031033123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313103230222222-3030010330200133-3020130231320033-2300001023230013-2231201101123230-1220001123203211-1100312211101220-0113011213100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1012232312321101-0203222303321100-3012202112323122-0111022012101111-3031030322000230-2011111111321233-2333331311230112-3230002100031300"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
preserve_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320331111000330-0001310020333213-3023100210221020-0222312121220100-2010231203130303-0032322313212312-1021211211020203-1331221301321001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2031333333120130-1313333112313030-3110132020032133-2120222100132220-3330211323132222-2220303320133110-1331010103331323-3100230330012010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
proper_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101220321323200-3010303311332022-2200103113311123-2011012133023231-0200303133000030-2203021330022010-1332100211310313-1211231313030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0131313211230320-2021232221321321-1332111233200011-3023020201330312-1130233230110013-3312320213123131-0032333122232010-3211100302220131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303321301310212-0201133032133311-2022213011111322-2201031302231010-2230020003220030-3310032003331203-2203312230000201-1010230112312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3212020033102322-1213212323032120-0213200022322013-0301132032210203-2123022130203312-1233303111011011-0031132332020321-0203212120220301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031001112333220-2323222300022033-3312203100203123-0312020013232333-3331202232200023-2221322031111020-1322133232110000-1001031112123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-1322331222021311-0102013012323033-3011303122032322-0003232210321233-2203203301330223-3013312010111323-0232111113102122-1311032021230010"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032131103202223-1332201101321212-0121230133302101-1030012020120133-1103202333112220-2221232200013000-0113013330323023-0230322022203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-2220003321303202-1233223212003031-1222030023201220-1301123311121313-0232031332322311-2313332033001133-0030020313223031-0110022100302222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310121301102211-2313201113000003-1230032122220332-2023020012232231-1012020102310102-3010221322312003-3121000111211112-2312300123110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-1011133132220113-0100022133303333-0221231133011033-1002002123132211-3023320301231111-1031133031030322-3022032122233112-3222010210212021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-1101321313001121-1213000233030003-3313012310020000-2330012200331112-3220032121011033-1011213033101221-2203111301033030-2202212030010211"></a>

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

<a id="canonical-1310030313032203-1122333002202131-0031202022333111-3301112023110000-2023232101032122-3200132103321330-1032132020010112-3221303021100132"></a>

### Direct properties for `proxy_config.https_auto_cert.tls_config`

- [custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2121323333013302-3300320220203030-1100300210011003-0302122330311232-3212203231303303-3233111201002231-2303101220112202-3231032112113032): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1311033033013301-2311020033233032-1320110203121131-2102101000302002-1233021103033100-1233320001201321-2211032211310022-1001100333201010): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0232033033222320-3000201022300103-2010210223102312-0203003202222332-0132123032333021-0303110020130033-1031232122210011-0003212200030001): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0131330113200032-1131213213130232-2233120223102130-2001003233120011-2012001323201233-1113120312003221-2101133003021213-2013230111020212): complete subsection reference.

<a id="canonical-2121323333013302-3300320220203030-1100300210011003-0302122330311232-3212203231303303-3233111201002231-2303101220112202-3231032112113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-2313203210100331-3033303200001332-0330020331233103-1003001013203200-2212030130112201-2003232021001031-1201133311320120-3103111121112223"></a>

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

<a id="canonical-3003230002303323-1133132222032233-2112103103132113-0013130103013202-2110022030222033-1130010310033002-3013113331220322-0031210120323013"></a>

### Direct properties for `proxy_config.https_auto_cert.tls_config.custom_security`

<a id="canonical-1220222203332022-1213001002122231-0212200110331121-1302000202210220-1131001301023120-1013310330320311-1212011223111221-1320312102230110"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-2012113233121211-0203031012013333-2300130132123111-0310032322101213-2131310101231030-2030233303010322-1321212012023123-1330202210210332"></a>

<a id="canonical-1101023100321312-3202232200000113-1000203312330132-2223213210022123-2102000211100120-0113231003123203-3033331300000231-1220131112321332"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-3031111101301111-0001231222303132-3023130202022013-1223313113130120-3230032032201313-1221333103332120-2200223022231013-1200112131203032"></a>

<a id="canonical-1132100232203332-1013323210231001-3132021222020031-2333302221322310-2323200121130030-2322311301332112-3212123011033230-3020312202131203"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-1311033033013301-2311020033233032-1320110203121131-2102101000302002-1233021103033100-1233320001201321-2211032211310022-1001100333201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- proxy_config.https_auto_cert.tls_config.default_security

<a id="canonical-2202321002001131-2333010123133101-0011030321323120-1232313310233022-2120120021321132-3313011331301323-2222102103113020-3321002131003302"></a>

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

<a id="canonical-0232033033222320-3000201022300103-2010210223102312-0203003202222332-0132123032333021-0303110020130033-1031232122210011-0003212200030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- proxy_config.https_auto_cert.tls_config.low_security

<a id="canonical-0131122301123112-2202202013133113-1221311312302230-1300033220330312-3310331212210203-0320001011231020-0202333100023021-3132013101020331"></a>

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

<a id="canonical-0131330113200032-1131213213130232-2233120223102130-2001003233120011-2012001323201233-1113120312003221-2101133003021213-2013230111020212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- proxy_config.https_auto_cert.tls_config.medium_security

<a id="canonical-1321233002030111-3132110222313321-3111101310030312-2031121301122300-0211120233323123-3031112020210313-2320301013101111-1032300113311130"></a>

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

<a id="canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.use_mtls

<a id="canonical-0322202010222322-3213113223233213-1222022331200032-2110231101321221-1331110231102212-0223001113332212-1021331213120221-1013320011212211"></a>

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

<a id="canonical-0011210010210111-3231221123220211-2303122013003320-1210201333033300-1331313003312231-2100313311120011-1031101210033301-1121223201121323"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls`

<a id="canonical-2123130003021302-3310132121121133-2013001301320202-1033210312221022-2323123312300111-3102203322131230-0212300101311131-2010110332031231"></a>

#### `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3313222100032011-0011230102312213-0323331102201201-0001112233133021-2012013122302032-1232112221001320-1332200323022031-1001022130332330): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1132223211100320-0320223221120330-2102013330220021-3020323031113303-2230010013201301-2212331001230303-1122100012033200-3323302223222000): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1013010021312231-0021102311120011-3010331030030300-3211301321200212-0113201301310233-3022320010012202-0131320331320203-2310112001211200): complete subsection reference.

<a id="canonical-0010310332131023-2011322023211001-1103002232210212-1101102213333321-0122100302133231-0013132220312113-3002130131102033-1102333103012302"></a>

<a id="canonical-1311123030211122-2001300313032203-3101131322133031-1102232300200001-3103021103022213-0132312311120213-1312331212010213-0022013000200312"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-0331310322010022-2123322101030323-1033131223220203-3311111211121211-1111201013321032-1023110211223201-3303213023231031-0000011300332000): complete subsection reference.

<a id="canonical-3313222100032011-0011230102312213-0323331102201201-0001112233133021-2012013122302032-1232112221001320-1332200323022031-1001022130332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-3122100112000233-0311120333110101-3212101321123231-2202010132012331-0033101022131030-3300201221020030-1123330330203132-2322022110013331"></a>

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

<a id="canonical-3000310032033211-1302130202301200-2133212310122021-2101200023101222-1032231103311331-0223010302020033-0333022131230120-2201101110101233"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.crl`

<a id="canonical-2020200021202132-1120122212130130-2012232132011022-0230032223133222-0130100331232202-1022210310210301-1033322010300112-2203131020201312"></a>

#### `proxy_config.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-3113223212103113-2323230221002121-1323233001011113-3211313330212202-0133010011102231-1213013021010331-1032031303233310-0110103030221332"></a>

<a id="canonical-2200030301321321-2103232102001003-0330100222100000-0203311103212032-0303311023311113-3112310232000002-0030321230220323-1300121110002331"></a>

#### `proxy_config.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-1112220102103213-1331230331232231-3310220330202123-0322330001033112-3303232111203320-0211320203223331-2013020112232213-0210101232000201"></a>

<a id="canonical-2003202130320030-3332100001002321-0121303010311322-1330101233211231-1301321302013122-3000022003030201-0310003020220122-2131001033223123"></a>

#### `proxy_config.https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-1132223211100320-0320223221120330-2102013330220021-3020323031113303-2230010013201301-2212331001230303-1122100012033200-3323302223222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.no_crl

<a id="canonical-0331033310013030-0022023111122310-1320200321132021-0002132032312300-0230321310230213-2301131010113132-3021013112023310-1300223223310312"></a>

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

<a id="canonical-1013010021312231-0021102311120011-3010331030030300-3211301321200212-0113201301310233-3022320010012202-0131320331320203-2310112001211200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1022212220320001-1323001123311200-0203031020331131-3001031320123120-0300323201103002-0313013012200332-1010120203111302-0133232033220103"></a>

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

<a id="canonical-3030331223002021-0002122313122120-3203202301213001-1323101220013013-0321020001113311-3302330303111312-3300200333001322-1231222010311010"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-3230313211001333-0030221213223130-3022110110310110-1221223302200311-1200131131202011-3131321031003321-2212000032123232-3003330200311201"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-1133311202323131-1203002110031302-1120023012122321-1023300201222322-3133101132003320-0033322023232100-0301032112323032-1003120013111111"></a>

<a id="canonical-0211030231320130-0123220130200330-1013203303122032-1100200333133122-1312200222212310-1333321212120302-1330210323002023-2130132130130311"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0123122023022310-2233012221302221-2221032103100031-2103332233020111-1020110332013100-3233312023301213-3211333032133221-0123333033000232"></a>

<a id="canonical-0030323202011023-3213101303122312-0210002133312300-2330303000330210-1132321323230122-0031232321131313-2013231121221233-1222010233300001"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0332021101303111-2321212313232113-2311232212001311-2212203011011021-1221333101001300-2013200220200333-1223323310232110-0033131013000023"></a>

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

<a id="canonical-0331310322010022-2123322101030323-1033131223220203-3311111211121211-1111201013321032-1023110211223201-3303213023231031-0000011300332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-1110022331333333-1303013102221111-1213333230130330-1030133231302313-1112220300122311-1003103212200100-3301013200031011-2122101130311110"></a>

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

<a id="canonical-0130331232101302-2132022330102333-2321022102230012-3233311222321211-0320211133000232-3330310201220310-3300230312211102-2223312123022302"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-1112003113111331-2123221101023333-3233023323222010-3132322001133120-1111001230300223-2323310313221100-1201321031012002-1310113030111332"></a>

#### `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1231333103033012-2112201220312122-2222023320130000-2122221200023233-1122012332200101-1313233133310311-3230012111033123-0101131100232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- timeouts

<a id="canonical-1120110110032221-2003012033023222-3122333201020321-3201331110133123-0233100003200311-0001322102030133-3011131312323321-3122331223231202"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003132121122130-2301011030132302-3132230021331210-1230222331113331-2322210123233331-3303033322002122-2110300321112111-0323202100210321"></a>

### Direct properties for `timeouts`

<a id="canonical-3021002223001012-0202020332321100-1312202302020313-3011210111231030-3213131023322321-2103102303123301-0301001032030221-3230333202203011"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1221203203020013-1111003330301122-3222020003321200-2021212331111321-2100010321032303-2233033222023113-0222023310011022-0331232212311112"></a>

<a id="canonical-2023223120322012-3233222302030001-3303110020023231-3231103222333133-0210111011030031-3221102221203221-0311230020021221-2112211312220031"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3111202213100121-2023201200020303-2203123210101010-0311230012132300-2031113223221113-2311030330020320-1102202032333323-3032023221011013"></a>

<a id="canonical-0203210130031300-0000201323002321-1012000130022133-0121021330012131-2011122322322023-2111311023100100-0300301323001113-0230113300332221"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1102310223003003-3233021331311302-3201131102023300-0121100132013020-1203033102020030-0001010201201022-1132121332021222-3103333330003310"></a>

<a id="canonical-0220113203011311-2212303112323003-0000220011200222-2130230332302113-0112003333312121-0103100221220331-1000102222203331-3313111330023223"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
