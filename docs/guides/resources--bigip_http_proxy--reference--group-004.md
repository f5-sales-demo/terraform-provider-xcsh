---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0102330111321010-0302230200222201-2330010030013220-3111200113122110-1230333200303121-2013100213131101-3132201021232020-0003123123023313"></a>

## proxy_config.https.tls_parameters.tls_config — tls_config / 110031103020 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-2103323030031233-2211133023111033-2132023100322132-2331320220012233-0223223032201022-1301023100321302-1102000013211332-3302012003202223"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

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

<a id="canonical-3021331302002221-3320032130122210-1313233003222131-3020332000200011-2032003230022002-2131303110002223-2202132332213202-0220212310211110"></a>

## Direct properties — tls_config / 110031103020 / 3

- [custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1033102301120112-0111120023211201-2100211300330213-3010031033010330-1111301311201012-2102300132010312-0330332112033023-1120031100331013): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1001302030122213-2212331100230223-3101020120201233-0322202103201113-1332020112312000-0001323311112233-3320200112101330-3023120132020100): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-3330310303203023-2110013312120302-1231122110200313-3022003222230031-3032031332003231-2210220220231220-3121112121023001-2033132121322102): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1030031031103101-3101301320202011-3031130233302300-3220211202033301-0120321113202212-0123231033333202-3202233332313211-0123311333021001): complete subsection reference.

<a id="canonical-1210122103031120-3333000231033031-1121002311103022-2020233133111033-2300321313300131-2101320323321033-3313020110023133-3131103102212133"></a>

## Next pages — tls_config / 110031103020 / 4

- [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1033102301120112-0111120023211201-2100211300330213-3010031033010330-1111301311201012-2102300132010312-0330332112033023-1120031100331013)
- [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1001302030122213-2212331100230223-3101020120201233-0322202103201113-1332020112312000-0001323311112233-3320200112101330-3023120132020100)
- [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-3330310303203023-2110013312120302-1231122110200313-3022003222230031-3032031332003231-2210220220231220-3121112121023001-2033132121322102)
- [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1030031031103101-3101301320202011-3031130233302300-3220211202033301-0120321113202212-0123231033333202-3202233332313211-0123311333021001)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1033102301120112-0111120023211201-2100211300330213-3010031033010330-1111301311201012-2102300132010312-0330332112033023-1120031100331013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201122223103303-2122010320303321-1231311023231132-1001002113030011-2112132000131223-0202210331002230-1302130303132133-3131003022020010"></a>

## proxy_config.https.tls_parameters.tls_config.custom_security — custom_security / 123132102021 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-2210003101010303-3213302030201200-1223323223330133-1101110322332312-1322302203210230-1310233210130310-3032112312000001-1303200130311330"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0200030122301331-1002013321301033-3222020232112203-3013222301232310-3230222000202121-1020220313320101-3203230103321133-3100131023233131"></a>

## Direct properties — custom_security / 123132102021 / 3

<a id="canonical-3323033112233302-2102233202022010-2031312223330323-0322112102033321-1130013033033131-0200122333003312-2310031023302223-1101200201111303"></a>

<a id="canonical-0102010003003330-0322300223111213-1130301320032110-2232022112332331-0201111120310203-3223321301131030-3223133233202210-3121123222033031"></a>

## cipher_suites property — custom_security / 123132102021 / 4

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

<a id="canonical-2233022230022031-3022210133310300-3311001103120332-0311130301100121-0200020103102300-2210011200302101-1013030320212223-3320200201111120"></a>

<a id="canonical-1213000332310332-3331302000120233-2022023131130101-3111010311013122-2322203203121122-2322212200220021-3010133100122121-1130132011122001"></a>

## max_version property — custom_security / 123132102021 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-3011132302031223-0212332323110111-2012322223213002-2132311302322212-0000311211101300-0022213131202103-2310312000300310-1021210111230130"></a>

## min_version property — custom_security / 123132102021 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-0233212233131212-0223201002003102-0113103222331330-1213032002200023-1122113110300232-3032000031203131-3300122230121332-3220132000221013"></a>

## Next pages — custom_security / 123132102021 / 7

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1001302030122213-2212331100230223-3101020120201233-0322202103201113-1332020112312000-0001323311112233-3320200112101330-3023120132020100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311012300202133-3032111033232222-0132113132312102-3112211312123200-0012323021131200-0111003213230212-0000111013013222-0131320233210100"></a>

## proxy_config.https.tls_parameters.tls_config.default_security — default_security / 133223011203 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-1022333100332033-2202313233010321-0033022011010231-0011021211031123-3002112000013110-0333322103010110-2112121010233130-3100123132023301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-2123032123201121-3100310000030031-3020233113130102-3221311002200221-0211311211332300-0320033313230113-1203331021103221-2302111032220231"></a>

## Direct properties — default_security / 133223011203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033133032033303-2120030211002121-1320020220222301-1201113302223011-0133222333223322-2210022312333001-2233033223023332-3113032131320210"></a>

## Next pages — default_security / 133223011203 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3330310303203023-2110013312120302-1231122110200313-3022003222230031-3032031332003231-2210220220231220-3121112121023001-2033132121322102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312001122002133-0023333223122110-3013333231000033-1133320002000300-2132223031200123-2310023113210111-2213011130102113-0322203023010222"></a>

## proxy_config.https.tls_parameters.tls_config.low_security — low_security / 302101332300 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-0212100230111030-1120211311122313-0130021100300301-0133232230003123-3130133203032303-2111020323111011-0311231333021201-2012120331102000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-1013030101332222-3001111130220032-3330113003232030-3230110003003033-1322331122023230-0023100221113310-3102321303202012-0231213012221322"></a>

## Direct properties — low_security / 302101332300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302321312200011-1212231120220331-2330001333203330-2221113202223330-3001012000002231-2030022300211202-0232132231212110-1320221023203101"></a>

## Next pages — low_security / 302101332300 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1030031031103101-3101301320202011-3031130233302300-3220211202033301-0120321113202212-0123231033333202-3202233332313211-0123311333021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001113311032112-0022103113331020-3220100323200220-0321032310323313-1103102310223112-1230312022231103-1200333222220031-0313231323200110"></a>

## proxy_config.https.tls_parameters.tls_config.medium_security — medium_security / 023232032203 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-0323133222230223-3132333012313201-0310113302022122-1210022232321131-3031201022012310-3330300002021201-0013300232103323-1331110032131202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-1111110331123322-3310131100332033-0221020130021100-0031211020111000-0323332200200322-3211210030101010-2103330132132313-3202011322232223"></a>

## Direct properties — medium_security / 023232032203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102031301222222-0030011223221230-2131233330331302-3222031301221111-1202231221312020-0032011232203011-3121010200002123-2322221222010212"></a>

## Next pages — medium_security / 023232032203 / 4

- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002013303000031-2112303311032101-1131123210000010-0030302131311200-3233113221101032-1102003031122302-1313023323332021-3332123031330211"></a>

## proxy_config.https.tls_parameters.use_mtls — use_mtls / 220332133231 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.use_mtls

<a id="canonical-2202303331113300-3101033121011310-0211110012020211-3202333122112301-0023232301101203-0111110122200310-0230301102120103-3200322322032202"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
```

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

<a id="canonical-3023212313113322-0221022322032333-2300012030210113-3213313120313100-0231332032311200-2012302220113111-0110210323030132-1121310301022232"></a>

## Direct properties — use_mtls / 220332133231 / 3

<a id="canonical-1323201330232021-1212100312030110-0020020201022100-0020210302330120-1310203330332002-3122112003211233-2003210131011302-2300010113202331"></a>

<a id="canonical-1013033223323133-1023133321032212-0023110130012231-1020132312020132-1220212321032310-0303011330032221-2232211322333012-0222322312020011"></a>

## client_certificate_optional property — use_mtls / 220332133231 / 4

Type: `"bool"`. Optional.

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3301132013002133-3300022203123130-0133211010212100-2012303233130300-2022332111110131-0212021230201221-2030321322130002-3031203021201132): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1122000200022122-3121300310203023-1223122012113203-2233112113122201-2232013203110223-1212001333132023-0231003003331202-3323321002130101): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1010131132012330-3113302320201301-2211330221201003-0321031133033121-2021333131022213-2301330203332311-2112330113100011-0132212321211000): complete subsection reference.

<a id="canonical-1321220310111301-1003313120032121-1213103103013301-0223300031223212-1122300033032212-2321330100301100-1131023101100221-2310112221020200"></a>

<a id="canonical-3201020313200211-0322230321333133-3203030313111211-1201113322313313-3111332211322322-3330221330011120-0321310310000321-1310001100302101"></a>

## trusted_ca_url property — use_mtls / 220332133231 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-0111131122320313-3201233031331303-3100132001000021-0213312100201200-3120223313203300-0031101330100012-3210112310010210-3022011031010300): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1130033202003022-3030302000113122-1213132121121231-2013222200323313-3322021011131110-0232030100121233-2332332322103310-1330330112321030): complete subsection reference.

<a id="canonical-0011011233230310-1121003023322110-1023103012322000-2301010122032132-3033011300023133-0201322202121010-3031021103322123-1010112202112210"></a>

## Next pages — use_mtls / 220332133231 / 6

- [proxy_config.https.tls_parameters.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3301132013002133-3300022203123130-0133211010212100-2012303233130300-2022332111110131-0212021230201221-2030321322130002-3031203021201132)
- [proxy_config.https.tls_parameters.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1122000200022122-3121300310203023-1223122012113203-2233112113122201-2232013203110223-1212001333132023-0231003003331202-3323321002130101)
- [proxy_config.https.tls_parameters.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1010131132012330-3113302320201301-2211330221201003-0321031133033121-2021333131022213-2301330203332311-2112330113100011-0132212321211000)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-0111131122320313-3201233031331303-3100132001000021-0213312100201200-3120223313203300-0031101330100012-3210112310010210-3022011031010300)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1130033202003022-3030302000113122-1213132121121231-2013222200323313-3322021011131110-0232030100121233-2332332322103310-1330330112321030)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3301132013002133-3300022203123130-0133211010212100-2012303233130300-2022332111110131-0212021230201221-2030321322130002-3031203021201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102313303032331-2332302000300300-0311101112201103-3031203030011223-3203200103123101-0121013221131233-1033222203313023-2113233310331201"></a>

## proxy_config.https.tls_parameters.use_mtls.crl — crl / 002312311100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-0123233322023122-3310221222012200-2020023031002133-1123201333210201-2322013300203211-3321232331223333-0223313223201132-2332302211210222"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2122112220030132-2103331312033331-2121301202031301-0101200123331302-0013122220000110-0023112002023213-0203102011302001-1121001212212010"></a>

## Direct properties — crl / 002312311100 / 3

<a id="canonical-0012120003132232-3000121133223031-0022202203213133-2331133320002333-0100222012323013-1220302110223101-2123123130303213-0220030133111102"></a>

<a id="canonical-0013311123332021-1111201112212130-3102030100230321-0033311012313211-1231002233202231-1000313030020221-1203021021321131-1211100300001212"></a>

## name property — crl / 002312311100 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3122230101301020-1113132300300211-3302132230010101-3213312010323232-2302300133231223-0330121311303310-0232031121322301-3322330033313312"></a>

<a id="canonical-0111203022302103-2023021111322203-1133323003311102-0322100333311310-0001123213123122-0330032031311222-3121100313211202-2311000100100202"></a>

## namespace property — crl / 002312311100 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1303233313232301-2303300000020320-0133010030102333-3221033112102203-3013323231001321-3111033020213000-0110032201013323-1123020210220221"></a>

<a id="canonical-1230113113002212-1032111201012313-2201220020201320-0333223322222303-3333221210330231-0110223212101131-3101200320230233-2123210211013020"></a>

## tenant property — crl / 002312311100 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2310132113103321-2232033103102111-1031100312221122-0130102200122202-1333011331131311-3103133030330023-1132011130012123-3322310101302032"></a>

## Next pages — crl / 002312311100 / 7

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1122000200022122-3121300310203023-1223122012113203-2233112113122201-2232013203110223-1212001333132023-0231003003331202-3323321002130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221323100222121-3120302021131331-1333300020202033-1133323002012131-1303310022102111-2332011100030203-2022002333231032-1132232332222123"></a>

## proxy_config.https.tls_parameters.use_mtls.no_crl — no_crl / 230030112120 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-2101011211001301-0201311300323303-2201111013301002-2221031131030200-1010232101101003-1323003122033203-0300111203311003-3012012030231132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_crl = {}
```

<a id="canonical-2331102133021233-2330122133300232-3132200213103121-3211233211032002-1132312331212201-2133131213033223-3233310200233133-2203332231123231"></a>

## Direct properties — no_crl / 230030112120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202000200323322-0221112323123132-0101313131033012-2100211300331012-1313122111103012-1120221012010303-3322021200313301-3203002000313123"></a>

## Next pages — no_crl / 230030112120 / 4

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1010131132012330-3113302320201301-2211330221201003-0321031133033121-2021333131022213-2301330203332311-2112330113100011-0132212321211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132103222033011-3321111113031332-2001232120011311-0103131320301302-2202321002132200-3222312222233013-1310302232003302-1133120121300301"></a>

## proxy_config.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 312100010320 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1133110232210011-2201103002102002-3300133300101313-1300002131011010-1102313112002022-2010003113123022-2332030301132321-2301012231133302"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3231022320322120-0130203003230333-3022023002310323-1321203031302130-0132232202130211-0310232110212322-2223223301002102-3332233330032323"></a>

## Direct properties — trusted_ca / 312100010320 / 3

<a id="canonical-0221323322330333-0100032031320121-2203002301212221-1033030131321231-2121120111202113-1232301010000213-0122131201311030-1102311320321331"></a>

<a id="canonical-0010233323333031-1202001320130112-0120121112230232-0232323211200011-3333233301002232-0000031323312021-3022313331333012-2233113330012313"></a>

## name property — trusted_ca / 312100010320 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3103003221321002-1313113211002222-2220311203312212-2302000230201321-3123232102020031-2003300300103223-1121313111101031-1110302211033012"></a>

<a id="canonical-0000331332002321-0133031332200020-2303202122302131-3323121022321312-3221012231303033-0123221021223102-2030020111213233-3211331133220032"></a>

## namespace property — trusted_ca / 312100010320 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0022203003022302-1131002012121320-3330223230322021-0030111130112002-3210231032020032-0132313112100223-3020132132130123-3232032200333010"></a>

<a id="canonical-3133223010012311-2330111213203000-2200112303103303-2213101103302123-0312031010031333-1212100111031122-2030321101302000-2102321303332002"></a>

## tenant property — trusted_ca / 312100010320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-0131321113110301-0111120322220212-0032113032211030-0001213231130133-1020100321010031-2003303112021200-1010321313022103-1010210012010300"></a>

## Next pages — trusted_ca / 312100010320 / 7

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0111131122320313-3201233031331303-3100132001000021-0213312100201200-3120223313203300-0031101330100012-3210112310010210-3022011031010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222212100020133-2213031122202203-1311302000302022-1133322002203203-3021101121121102-0333133302112220-1220211301030321-1120333230023321"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 320022302130 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2010022220303330-3000033201310110-2101011330223312-2021222033112200-1131330023322120-3123331023020320-0033223001032133-2302012110010220"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

<a id="canonical-0322122230101312-0122103033011002-0303210122322111-2032030212232311-0120121011120222-1001230301330322-2102001312223213-3122131113130232"></a>

## Direct properties — xfcc_disabled / 320022302130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132212111011132-1313321333222002-0003021011210110-3132133213122221-0230220330321203-2022232201300302-3311012233111021-2031212303212220"></a>

## Next pages — xfcc_disabled / 320022302130 / 4

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1130033202003022-3030302000113122-1213132121121231-2013222200323313-3322021011131110-0232030100121233-2332332322103310-1330330112321030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120010001123120-1023012101323111-2123110213220331-1311202111213201-0100000022220132-1100321122001110-3233310003110011-0033222100100332"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 032310010300 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- proxy_config.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2012203301003332-0100201131320312-2113020323131013-3301311113322132-1131210332321203-1232030201233331-3232113310333210-3302121310212100"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1331111120130103-1013113023003201-3201013133332030-2301021002021300-3332133021230223-1000230333301302-3013313312202200-3303110110232301"></a>

## Direct properties — xfcc_options / 032310010300 / 3

<a id="canonical-0231031333310200-2120233223033001-3110213031133110-0303131230132000-0100311020312221-3323122132112010-2023202000232221-3210120322201203"></a>

<a id="canonical-3012103001220321-1221231100323323-1232212030212010-3201110100020323-1232131322002213-0133133213230113-1320222232003131-3000302202220301"></a>

## xfcc_header_elements property — xfcc_options / 032310010300 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3102331103120313-3030111212021213-1333113210031131-2031102001033120-2030100323211003-3133222130330200-1032132213210221-3213230321202322"></a>

## Next pages — xfcc_options / 032310010300 / 5

- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012013233010312-2232112322333030-0300110221130123-0023210013011221-3223303113121030-0020200122303331-0001020310301301-3000211313302200"></a>

## proxy_config.https_auto_cert — https_auto_cert / 031132201110 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- proxy_config.https_auto_cert

<a id="canonical-1232130333122220-1122323101230313-3221320113210210-0220120212020321-3021313123110330-3103112331000021-3123222113203022-3022132121330232"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
```

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

<a id="canonical-0320201311131003-2301032032103110-0330112110333033-0222210133012212-1023213111033333-3032131302303132-1220311003320331-1021230131123013"></a>

## Direct properties — https_auto_cert / 031132201110 / 3

<a id="canonical-1021302330123221-1213000312303202-0300312012313312-3332123011132003-2221001220123113-0301032210010130-3200333222220301-2231021323213320"></a>

<a id="canonical-2003012220202311-0011032312210232-1311113321212130-1223212103322132-2200021123310331-2232213003201212-3133001033000232-3310102311201112"></a>

## add_hsts property — https_auto_cert / 031132201110 / 4

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

<a id="canonical-0211131320231002-3122032113032213-2111113022322103-0210102310100322-0321100313313002-3013003010211010-1231120012303001-2231103000233003"></a>

## append_server_name property — https_auto_cert / 031132201110 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211): complete subsection reference.

<a id="canonical-0120223312103123-1032030101311003-0000120121302233-1030021320203203-3323202301120112-1300102022312220-0201130132113003-2222120310000023"></a>

<a id="canonical-1010003220122033-1032110002112332-3030120213132310-2031200123121013-1303302003030000-3303011220313223-3320131212220220-2130202300333220"></a>

## connection_idle_timeout property — https_auto_cert / 031132201110 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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

- [default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-3322322023021102-0330023200022121-3123222001202000-0321220120321102-1013300001032221-1101313203133033-2001230222102123-3302301022203300): complete subsection reference.

- [default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1322223303132201-3203203230103220-2203103203202010-1231131323001112-0002301130030320-3022300303222223-1322032020302311-3130100322312311): complete subsection reference.

- [disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-1300003311013210-1200230101211221-1120000300213310-2321300001022333-1103120100000210-3030332320232100-0221202203100303-2023221122230203): complete subsection reference.

- [enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-2301232330002123-1001022220103322-3220010023133233-0323231331233301-0211031322212011-2123223312310101-2332331333100330-2231131033301331): complete subsection reference.

- [http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021): complete subsection reference.

<a id="canonical-1222321203200333-3031011110100312-0111110013013200-0101110232001131-1011022313132201-2101012233030011-3030123020213222-0022013012320220"></a>

<a id="canonical-1121022200000231-1320222220023333-2233320333312320-1221012303230232-3312312201101303-2212221020330213-3200102221130233-0111212031122022"></a>

## http_redirect property — https_auto_cert / 031132201110 / 7

Type: `"bool"`. Optional.

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

- [no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2031001112333220-2323222300022033-3312203100203123-0312020013232333-3331202232200023-2221322031111020-1322133232110000-1001031112123320): complete subsection reference.

- [non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1032131103202223-1332201101321212-0121230133302101-1030012020120133-1103202333112220-2221232200013000-0113013330323023-0230322022203201): complete subsection reference.

- [pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-3310121301102211-2313201113000003-1230032122220332-2023020012232231-1012020102310102-3010221322312003-3121000111211112-2312300123110302): complete subsection reference.

<a id="canonical-2030023110120333-2011333201022013-3323133032330233-2010332011100312-0210033032233131-0203111033011300-0301312303313012-1033012210323011"></a>

<a id="canonical-0013333323311010-0030201333230030-2021123301110101-2222010033003222-1201211222012300-3333211303311203-1011323102022233-0032103031313023"></a>

## port property — https_auto_cert / 031132201110 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

<a id="canonical-0131123230120033-1310101200201030-2020332122330312-2313231221221103-2131110020012100-2132220220312212-0300121330110210-0203323011000213"></a>

<a id="canonical-1302032203021220-3200231322101131-2323113113032110-3203220301031223-3230113200303220-3030333212330322-2203333113232313-3210213321001023"></a>

## port_ranges property — https_auto_cert / 031132201110 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-3222011031201331-1011131323210012-2121202201131011-2321320301122232-0311200333112331-2031030113203221-2302013023323013-3012330012011031"></a>

<a id="canonical-1300011202322312-3001211300203002-1233210301130021-1220123231130001-2303201131002300-0122333123202022-3312331123102210-1210013303330220"></a>

## server_name property — https_auto_cert / 031132201110 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101): complete subsection reference.

<a id="canonical-2021003000233011-1131033013321011-0201223122020000-3322111032202200-2203130022230201-0202023130033021-1000002320100223-0320003232122212"></a>

## Next pages — https_auto_cert / 031132201110 / 11

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211)
- [proxy_config.https_auto_cert.default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-3322322023021102-0330023200022121-3123222001202000-0321220120321102-1013300001032221-1101313203133033-2001230222102123-3302301022203300)
- [proxy_config.https_auto_cert.default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1322223303132201-3203203230103220-2203103203202010-1231131323001112-0002301130030320-3022300303222223-1322032020302311-3130100322312311)
- [proxy_config.https_auto_cert.disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-1300003311013210-1200230101211221-1120000300213310-2321300001022333-1103120100000210-3030332320232100-0221202203100303-2023221122230203)
- [proxy_config.https_auto_cert.enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-2301232330002123-1001022220103322-3220010023133233-0323231331233301-0211031322212011-2123223312310101-2332331333100330-2231131033301331)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [proxy_config.https_auto_cert.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2031001112333220-2323222300022033-3312203100203123-0312020013232333-3331202232200023-2221322031111020-1322133232110000-1001031112123320)
- [proxy_config.https_auto_cert.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1032131103202223-1332201101321212-0121230133302101-1030012020120133-1103202333112220-2221232200013000-0113013330323023-0230322022203201)
- [proxy_config.https_auto_cert.pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-3310121301102211-2313201113000003-1230032122220332-2023020012232231-1012020102310102-3010221322312003-3121000111211112-2312300123110302)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013023302323000-3101101331221111-3332222100103001-2233210313112012-3131222311302210-1102311030213010-3012031000322010-1001302023012213"></a>

## proxy_config.https_auto_cert.coalescing_options — coalescing_options / 001223120200 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-2023122032222133-0320312030031012-0233013300033023-1221200320321120-3133332302333212-1123032100210030-0130231233322030-1120012213013033"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
```

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

<a id="canonical-3210300133223301-1232113112111011-3232210030113001-1102011321233102-3011230101101323-2022231112113210-0330312110323321-3300101200330200"></a>

## Direct properties — coalescing_options / 001223120200 / 3

- [default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-1030310312101112-1223221122331332-0000000223201012-2300322333002113-0132121320133123-1200301103122010-3000210101203323-0302320303122033): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-2130220103102031-2021000303010132-3311001130323031-1113323123311220-2232203010201212-3030123313100330-0200230311313100-3223010002330131): complete subsection reference.

<a id="canonical-3132303232231003-3032332101221313-3002021222122122-1012231320030321-3013132133211321-2030223331211331-3230301132030012-2330133312100001"></a>

## Next pages — coalescing_options / 001223120200 / 4

- [proxy_config.https_auto_cert.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-1030310312101112-1223221122331332-0000000223201012-2300322333002113-0132121320133123-1200301103122010-3000210101203323-0302320303122033)
- [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-2130220103102031-2021000303010132-3311001130323031-1113323123311220-2232203010201212-3030123313100330-0200230311313100-3223010002330131)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1030310312101112-1223221122331332-0000000223201012-2300322333002113-0132121320133123-1200301103122010-3000210101203323-0302320303122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300123000130302-2233201101231000-0222330302022131-2131013110301033-1132122233210112-0333013323103022-1212102132002312-2133022123130311"></a>

## proxy_config.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 311131310221 / 2

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-0331013313130231-1120223000101100-2211220103103022-2121223021210220-0003320212113223-0110011130300322-3002232220300130-2331320123330102"></a>

## Direct properties — default_coalescing / 311131310221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301312212222230-1220001003310330-3102010112300201-1103103323000101-3012331201203001-0322323011033030-2123303203231320-3221221302201022"></a>

## Next pages — default_coalescing / 311131310221 / 4

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2130220103102031-2021000303010132-3311001130323031-1113323123311220-2232203010201212-3030123313100330-0200230311313100-3223010002330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123011222300201-1031013330302303-3000321200213232-3010301102133102-3120320132312132-1313322102230331-1321022330330303-2120123122022332"></a>

## proxy_config.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 211333112003 / 2

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-0030231222123201-1031311100031210-3000002030202013-2100010210202002-3313022302201212-1312210102312122-3321020031213213-3133212002030201"></a>

## Direct properties — strict_coalescing / 211333112003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231121002123323-1032101230120233-3000012132303312-3023331222132122-1223233221312322-2203201020230130-0303110102233111-2102111121001312"></a>

## Next pages — strict_coalescing / 211333112003 / 4

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1033131121123220-1021113011012202-0212310210313003-3013020123131030-0232300023313122-3012031110322210-1213232012220012-1211211003020211)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3322322023021102-0330023200022121-3123222001202000-0321220120321102-1013300001032221-1101313203133033-2001230222102123-3302301022203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001203302320233-2303212201233310-1020212300310121-3110331001012132-3133123123233300-1112201333020010-0132212200131001-1000330033021323"></a>

## proxy_config.https_auto_cert.default_header — default_header / 312002210313 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.default_header

<a id="canonical-3213133110002310-0013122112302013-3123322222020302-2330120220031131-0203132120210131-2333210023032210-3021313210022233-0320022322031101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-0301301220313330-2101233012333010-3123331003321013-2022123001031022-1313000212320021-0113211013332313-3331112330103233-3230120333202001"></a>

## Direct properties — default_header / 312002210313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122310110201023-2030320202101130-1232333210100122-2122310133301222-3310231301213031-2021110132213221-2201032210000121-1301122000310022"></a>

## Next pages — default_header / 312002210313 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1322223303132201-3203203230103220-2203103203202010-1231131323001112-0002301130030320-3022300303222223-1322032020302311-3130100322312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202000102232103-1201102111111331-2001213100321012-0202101012101203-2010223220000301-0223200212212023-1311330303333123-3233023110223300"></a>

## proxy_config.https_auto_cert.default_loadbalancer — default_loadbalancer / 221133112330 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-1012301213233321-3001220333032221-0220022203120031-0221020003130002-0233211221320112-3021201113020232-3213220112032232-2233231321200133"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-3110023321120311-3212322011213131-0033321212223112-0002321201321303-3211300013113300-3233220123300301-1333003102023002-3101122000113203"></a>

## Direct properties — default_loadbalancer / 221133112330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303233022202202-3123221123211330-2103302030032113-1023100021332311-2330303321301132-3202103333331030-1330312110013223-2110020311130132"></a>

## Next pages — default_loadbalancer / 221133112330 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1300003311013210-1200230101211221-1120000300213310-2321300001022333-1103120100000210-3030332320232100-0221202203100303-2023221122230203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331230222213032-2322333110121321-1331023102101031-3033131130121200-3302213033013123-2210330323133132-3010200231130220-3301222122332330"></a>

## proxy_config.https_auto_cert.disable_path_normalize — disable_path_normalize / 231121123211 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-3220013120110313-3313131231202332-3122133013011123-1203123220120330-2211332312100220-3332123303011222-0112032102312313-0033332013020122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_path_normalize = {}
```

<a id="canonical-2000023010232322-1102232200231230-0332011331232001-3033303100322100-3031213232100000-0020002313231223-0321301031131100-1222123111032123"></a>

## Direct properties — disable_path_normalize / 231121123211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210031113031331-1003201312112321-2022223013220320-0303302313101203-0333202000103311-0122123323311230-2022010021203333-3133101113302131"></a>

## Next pages — disable_path_normalize / 231121123211 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2301232330002123-1001022220103322-3220010023133233-0323231331233301-0211031322212011-2123223312310101-2332331333100330-2231131033301331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113023302303302-3220112011203101-2110320213320311-1003102212221313-1122200032030021-3321102331323322-1320321200120323-0322202222330301"></a>

## proxy_config.https_auto_cert.enable_path_normalize — enable_path_normalize / 030222331120 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-1220202012033120-1312202313313322-2123031100113031-3000313021200012-2202021021030323-0320321331001021-2021031122332232-1031220022013203"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_path_normalize = {}
```

<a id="canonical-1200220031021200-3213311312011031-0010013330312203-2132100321112232-1231201322330323-3211102112110132-1330302310212233-0331213121033002"></a>

## Direct properties — enable_path_normalize / 030222331120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133033311213221-0213231200321331-2212332210021212-1023203013130333-3001131100230331-1013023112032210-0213300031011323-0011032230231322"></a>

## Next pages — enable_path_normalize / 030222331120 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013201233113002-0220111011230313-1020102300200131-0110221330313100-3222302111211022-2121002212203000-1120131203310323-1032310103121002"></a>

## proxy_config.https_auto_cert.http_protocol_options — http_protocol_options / 210122133132 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.http_protocol_options

<a id="canonical-0222300021332000-2201002013200322-3323321302320123-0332321112311323-0021312230312301-2031312023032221-1003001303323233-2322030123112311"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
```

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

<a id="canonical-3122132220312121-2102230231203212-2022111021221113-1101031323223220-3002220030223300-1111123113233312-0303003321303032-0012202300211211"></a>

## Direct properties — http_protocol_options / 210122133132 / 3

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-1101220321323200-3010303311332022-2200103113311123-2011012133023231-0200303133000030-2203021330022010-1332100211310313-1211231313030220): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-0303321301310212-0201133032133311-2022213011111322-2201031302231010-2230020003220030-3310032003331203-2203312230000201-1010230112312010): complete subsection reference.

<a id="canonical-1313312212003202-1031013002321113-1303120331003101-1100021110322220-1031023322103003-0322001213031223-0202323102111001-3312020022313122"></a>

## Next pages — http_protocol_options / 210122133132 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-1101220321323200-3010303311332022-2200103113311123-2011012133023231-0200303133000030-2203021330022010-1332100211310313-1211231313030220)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-0303321301310212-0201133032133311-2022213011111322-2201031302231010-2230020003220030-3310032003331203-2203312230000201-1010230112312010)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231303122033131-3033211321333312-2001310330220123-2330320302030112-1011321101213303-0323020303330331-1233112202221331-3233302211232103"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 200203020203 / 2

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

<a id="canonical-2112232231223220-3333203132021331-2300312202313302-1011311033022003-1320101033320213-0211301311012332-3302221013122022-2311213122300133"></a>

## Direct properties — http_protocol_enable_v1_only / 200203020203 / 3

- [header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302): complete subsection reference.

<a id="canonical-2313123302330201-2120302222113203-2302202233033310-3002321122103230-0103313022320212-3110203022332102-2033031233301212-0011320233332230"></a>

## Next pages — http_protocol_enable_v1_only / 200203020203 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102131131101001-0031032310300313-0011302002311123-3023130311022311-3230101030032112-3013212121121100-1310202111331300-0033213212013200"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 222121112103 / 2

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
```

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

<a id="canonical-3031211220002001-1321222200020031-0223032011332222-0033012120203031-2202311120331313-0311013332323331-1112021130302121-1112211023223221"></a>

## Direct properties — header_transformation / 222121112103 / 3

- [default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3100331120311313-1223133311013032-2131220121322000-2013233200232000-0021021102311333-0012211102300222-0200010211001120-2113110223032221): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3313103230222222-3030010330200133-3020130231320033-2300001023230013-2231201101123230-1220001123203211-1100312211101220-0113011213100002): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-2320331111000330-0001310020333213-3023100210221020-0222312121220100-2010231203130303-0032322313212312-1021211211020203-1331221301321001): complete subsection reference.

<a id="canonical-0123130230221121-2121032121113200-0323320233100213-1131103321031013-0122332312123210-2212201310033033-0003102300003013-1213233303010100"></a>

## Next pages — header_transformation / 222121112103 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3100331120311313-1223133311013032-2131220121322000-2013233200232000-0021021102311333-0012211102300222-0200010211001120-2113110223032221)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3313103230222222-3030010330200133-3020130231320033-2300001023230013-2231201101123230-1220001123203211-1100312211101220-0113011213100002)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-2320331111000330-0001310020333213-3023100210221020-0222312121220100-2010231203130303-0032322313212312-1021211211020203-1331221301321001)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-1302201003113122-3331301101220323-1320012110120313-2032231030213001-3303023202021221-3233023220102312-0312320002100112-3111213131100120)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3100331120311313-1223133311013032-2131220121322000-2013233200232000-0021021102311333-0012211102300222-0200010211001120-2113110223032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020130332311100-1103012002311201-1333113131031221-1311003003230030-3123302022103310-3213003321330020-2103333112213130-3001300321102003"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 202130000111 / 2

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

<a id="canonical-0012311323022230-1003111210031332-3212132210013100-0201032010313320-2022133002120312-2223233211203331-1023232203012131-0321131303131101"></a>

## Direct properties — default_header_transformation / 202130000111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210310300102331-2222332220012300-3221100332003131-1010310103011000-0210003122213120-0000022112103302-1310200310010211-0213311113312131"></a>

## Next pages — default_header_transformation / 202130000111 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3313103230222222-3030010330200133-3020130231320033-2300001023230013-2231201101123230-1220001123203211-1100312211101220-0113011213100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103330112130003-3112232011330032-1220322300323303-3131012002130301-3230032003103213-1020220322123010-2300031033220310-3102301003112100"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 223131001032 / 2

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

<a id="canonical-0300021121032221-2133102010223110-0020330032312330-1213323322012333-3303103313202122-1313213101102213-0003200101232223-3302112002012320"></a>

## Direct properties — preserve_case_header_transformation / 223131001032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032130001301320-0000223032013110-0322230311021023-3112213010212131-3310003132131311-1001000303222223-3000331001313210-2222132221030330"></a>

## Next pages — preserve_case_header_transformation / 223131001032 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2320331111000330-0001310020333213-3023100210221020-0222312121220100-2010231203130303-0032322313212312-1021211211020203-1331221301321001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203022033013122-0132132100231233-1013113301322000-0122032121312113-1133001111121021-1032122000123131-2301231202130011-2202313002321101"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 000233112021 / 2

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

<a id="canonical-1021003211230113-2213312210310122-3213110200231302-3331030130331001-0130002302210113-2313021232002232-2313133232102312-1302300332332020"></a>

## Direct properties — proper_case_header_transformation / 000233112021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013323031100221-0112033330203100-2001010031012230-0121132332113231-0020033233201102-0032131103112020-2312301031123132-3203221130201002"></a>

## Next pages — proper_case_header_transformation / 000233112021 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3121312231220013-3210003110311110-2102312202203132-0210313331200300-2230012220310331-1031122331113201-0111330032312221-1301321013312302)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1101220321323200-3010303311332022-2200103113311123-2011012133023231-0200303133000030-2203021330022010-1332100211310313-1211231313030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130001232200011-2333032230230011-1120211310013200-2233301100033210-3122323311033302-2333121121322232-1202231012221221-0220031133201302"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 300303012202 / 2

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-2113020010223322-3103121002223011-1303233111003303-2222331001130012-3021311303011031-3213021020032231-2302213131012321-1302230322122033"></a>

## Direct properties — http_protocol_enable_v1_v2 / 300303012202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232233212320313-0032012023113011-2021030131121012-0312233112230201-1101031120122002-2333011031210213-2202231120200031-1300110031112220"></a>

## Next pages — http_protocol_enable_v1_v2 / 300303012202 / 4

- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0303321301310212-0201133032133311-2022213011111322-2201031302231010-2230020003220030-3310032003331203-2203312230000201-1010230112312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311220022032021-3133033003130013-2133131203332210-3333233212233330-2111102121231202-1111120113101120-1301332031123020-3101223131323233"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 300322331112 / 2

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

<a id="canonical-3112203101122211-0220033020330231-1200303100232001-1021223210223003-1023331000211223-3110230130131020-2222011211121202-3132222132313200"></a>

## Direct properties — http_protocol_enable_v2_only / 300322331112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331013011202000-3320213322103230-0001301130321200-2223331123233103-1310232031122130-3003233332110022-2030232121113212-3122203332002012"></a>

## Next pages — http_protocol_enable_v2_only / 300322331112 / 4

- [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1201113331033120-3333313103313322-1023133130131232-3312223323011331-0200030203232332-0332111003103330-1103113122302320-0010023220332021)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2031001112333220-2323222300022033-3312203100203123-0312020013232333-3331202232200023-2221322031111020-1322133232110000-1001031112123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230032031120302-2323000031021123-2221323001201213-1331123303010101-3201321320301202-2021113000032130-2111320032312111-3212230310033231"></a>

## proxy_config.https_auto_cert.no_mtls — no_mtls / 232211231113 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-1322331222021311-0102013012323033-3011303122032322-0003232210321233-2203203301330223-3013312010111323-0232111113102122-1311032021230010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_mtls = {}
```

<a id="canonical-0021122300322223-1110201010301022-0030331220301321-2221332102131121-3100123011133112-0312021212112032-1120122020123210-3130322102010201"></a>

## Direct properties — no_mtls / 232211231113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033302031133011-0032022311022223-1210322110200200-1322230101322130-3023323021213320-1203102201000213-3220323223300223-3101120311303113"></a>

## Next pages — no_mtls / 232211231113 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1032131103202223-1332201101321212-0121230133302101-1030012020120133-1103202333112220-2221232200013000-0113013330323023-0230322022203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111003020130303-1323101203212313-3233300322333312-3121130033110213-3012231122100210-2322130121203122-1223133132000211-2202232002230112"></a>

## proxy_config.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 300321332320 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-2220003321303202-1233223212003031-1222030023201220-1301123311121313-0232031332322311-2313332033001133-0030020313223031-0110022100302222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

<a id="canonical-3032102220101123-3222321011011012-0303031122111203-0012103002020010-2321030112121212-2331111113220131-1010010301203210-0020012012010221"></a>

## Direct properties — non_default_loadbalancer / 300321332320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222002210200313-3333103211112233-3233202211021230-2133131013320133-2322132030001312-0120231010303121-3322321330321123-2220002200001020"></a>

## Next pages — non_default_loadbalancer / 300321332320 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3310121301102211-2313201113000003-1230032122220332-2023020012232231-1012020102310102-3010221322312003-3121000111211112-2312300123110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330322230323230-3211230101300022-3312302233201123-3010033131333320-0302323300230003-3222001002100333-3132031223311120-3311230013313000"></a>

## proxy_config.https_auto_cert.pass_through — pass_through / 103002200331 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-1011133132220113-0100022133303333-0221231133011033-1002002123132211-3023320301231111-1031133031030322-3022032122233112-3222010210212021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

<a id="canonical-0200232112102011-1000103200223222-0310231202201303-3203021321302332-0331222330330300-0002330200013221-0032003222201232-0310302020203312"></a>

## Direct properties — pass_through / 103002200331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323322322013120-0031202033320022-3302330331230310-2312131231331331-0220232320032312-2301100032132110-1221220010000303-3211130021011201"></a>

## Next pages — pass_through / 103002200331 / 4

- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310030313032203-1122333002202131-0031202022333111-3301112023110000-2023232101032122-3200132103321330-1032132020010112-3221303021100132"></a>

## proxy_config.https_auto_cert.tls_config — tls_config / 232301332031 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-1101321313001121-1213000233030003-3313012310020000-2330012200331112-3220032121011033-1011213033101221-2203111301033030-2202212030010211"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

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

<a id="canonical-3231322130103212-3320231022102211-0231023223133032-2032122022210033-1322031103311113-1330113302111011-1130010232221023-2122201303003110"></a>

## Direct properties — tls_config / 232301332031 / 3

- [custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2121323333013302-3300320220203030-1100300210011003-0302122330311232-3212203231303303-3233111201002231-2303101220112202-3231032112113032): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1311033033013301-2311020033233032-1320110203121131-2102101000302002-1233021103033100-1233320001201321-2211032211310022-1001100333201010): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0232033033222320-3000201022300103-2010210223102312-0203003202222332-0132123032333021-0303110020130033-1031232122210011-0003212200030001): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0131330113200032-1131213213130232-2233120223102130-2001003233120011-2012001323201233-1113120312003221-2101133003021213-2013230111020212): complete subsection reference.

<a id="canonical-3103210211201012-0222022222302202-1022131330321213-1233001113000313-2022310213112202-2303302033210232-0201230022232101-0311310230330132"></a>

## Next pages — tls_config / 232301332031 / 4

- [proxy_config.https_auto_cert.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2121323333013302-3300320220203030-1100300210011003-0302122330311232-3212203231303303-3233111201002231-2303101220112202-3231032112113032)
- [proxy_config.https_auto_cert.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1311033033013301-2311020033233032-1320110203121131-2102101000302002-1233021103033100-1233320001201321-2211032211310022-1001100333201010)
- [proxy_config.https_auto_cert.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0232033033222320-3000201022300103-2010210223102312-0203003202222332-0132123032333021-0303110020130033-1031232122210011-0003212200030001)
- [proxy_config.https_auto_cert.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0131330113200032-1131213213130232-2233120223102130-2001003233120011-2012001323201233-1113120312003221-2101133003021213-2013230111020212)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2121323333013302-3300320220203030-1100300210011003-0302122330311232-3212203231303303-3233111201002231-2303101220112202-3231032112113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003230002303323-1133132222032233-2112103103132113-0013130103013202-2110022030222033-1130010310033002-3013113331220322-0031210120323013"></a>

## proxy_config.https_auto_cert.tls_config.custom_security — custom_security / 333331121332 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-2313203210100331-3033303200001332-0330020331233103-1003001013203200-2212030130112201-2003232021001031-1201133311320120-3103111121112223"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1101023100321312-3202232200000113-1000203312330132-2223213210022123-2102000211100120-0113231003123203-3033331300000231-1220131112321332"></a>

## Direct properties — custom_security / 333331121332 / 3

<a id="canonical-1220222203332022-1213001002122231-0212200110331121-1302000202210220-1131001301023120-1013310330320311-1212011223111221-1320312102230110"></a>

<a id="canonical-1132100232203332-1013323210231001-3132021222020031-2333302221322310-2323200121130030-2322311301332112-3212123011033230-3020312202131203"></a>

## cipher_suites property — custom_security / 333331121332 / 4

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

<a id="canonical-2012113233121211-0203031012013333-2300130132123111-0310032322101213-2131310101231030-2030233303010322-1321212012023123-1330202210210332"></a>

<a id="canonical-0132022331331222-0322101200212103-2230033112122201-1100320133301131-2200201001233310-3321101231031022-0300123221322032-3311131031031021"></a>

## max_version property — custom_security / 333331121332 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-0232121201323211-3302202332233300-3223112100212300-0213301001133230-2013201212210133-2102232001122111-2033323122320022-2113200320112032"></a>

## min_version property — custom_security / 333331121332 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-2103233231031231-3232032323032220-1231023013113121-2100201103200122-1102222122133133-0000322321202213-2230122020010211-0010223331003200"></a>

## Next pages — custom_security / 333331121332 / 7

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1311033033013301-2311020033233032-1320110203121131-2102101000302002-1233021103033100-1233320001201321-2211032211310022-1001100333201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212210211012333-0112200013222000-2220300122012013-1030322022313032-0202032200221021-2210111030000010-0032123303111313-1221121322101013"></a>

## proxy_config.https_auto_cert.tls_config.default_security — default_security / 132301231130 / 2

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-1033132100112332-2022101320211113-0123012320002032-3300203332210210-3300131332113211-0313210112331220-1210331322210102-2132100112111001"></a>

## Direct properties — default_security / 132301231130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301131301020132-3202132201312213-3013230123302022-3013120331030133-3122132333332003-0300331202000112-0311120222200320-2012223033313233"></a>

## Next pages — default_security / 132301231130 / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0232033033222320-3000201022300103-2010210223102312-0203003202222332-0132123032333021-0303110020130033-1031232122210011-0003212200030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002233333232313-3230322321232000-3310302220210132-2130222210201122-0122321201021231-0031210202011313-2111032202022300-1003330211303221"></a>

## proxy_config.https_auto_cert.tls_config.low_security — low_security / 313231323120 / 2

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-3133301111000010-2302111312202031-3001300320000003-2223221313221203-3021212013212320-2122311101233121-3102302330222111-2010003012303112"></a>

## Direct properties — low_security / 313231323120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120301313313023-3303322100203311-0103120111311013-0332223310101330-1232330220133000-3233120231012013-2110123000121103-1130303123312132"></a>

## Next pages — low_security / 313231323120 / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0131330113200032-1131213213130232-2233120223102130-2001003233120011-2012001323201233-1113120312003221-2101133003021213-2013230111020212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000113121202200-0021132120112011-1323312122213211-3331332200122000-1302111033213203-2010001331010122-2210020301312011-1030203003112323"></a>

## proxy_config.https_auto_cert.tls_config.medium_security — medium_security / 330331322000 / 2

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-2010210113210211-3201211203133303-2130230300221132-2100000032020312-0113333130021211-1113333113120230-3213132012333331-3020003010103301"></a>

## Direct properties — medium_security / 330331322000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110003233033203-2332100100010331-2101203202333100-0210032303111031-3131122011320220-0010131131322100-0102331331312313-0323233011000122"></a>

## Next pages — medium_security / 330331322000 / 4

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1030321110120213-3132102133300203-2330222303202003-3012330021212002-1310103221220021-2233322030010323-2312010230333301-0010123323312010)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011210010210111-3231221123220211-2303122013003320-1210201333033300-1331313003312231-2100313311120011-1031101210033301-1121223201121323"></a>

## proxy_config.https_auto_cert.use_mtls — use_mtls / 300022101132 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- proxy_config.https_auto_cert.use_mtls

<a id="canonical-0322202010222322-3213113223233213-1222022331200032-2110231101321221-1331110231102212-0223001113332212-1021331213120221-1013320011212211"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
```

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

<a id="canonical-1311123030211122-2001300313032203-3101131322133031-1102232300200001-3103021103022213-0132312311120213-1312331212010213-0022013000200312"></a>

## Direct properties — use_mtls / 300022101132 / 3

<a id="canonical-2123130003021302-3310132121121133-2013001301320202-1033210312221022-2323123312300111-3102203322131230-0212300101311131-2010110332031231"></a>

<a id="canonical-0020032322211330-1312102330301332-0313201103221222-1210010232330030-3021220211032223-0320111333230013-0010033121331301-2310320222322110"></a>

## client_certificate_optional property — use_mtls / 300022101132 / 4

Type: `"bool"`. Optional.

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

- [crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3313222100032011-0011230102312213-0323331102201201-0001112233133021-2012013122302032-1232112221001320-1332200323022031-1001022130332330): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1132223211100320-0320223221120330-2102013330220021-3020323031113303-2230010013201301-2212331001230303-1122100012033200-3323302223222000): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1013010021312231-0021102311120011-3010331030030300-3211301321200212-0113201301310233-3022320010012202-0131320331320203-2310112001211200): complete subsection reference.

<a id="canonical-0010310332131023-2011322023211001-1103002232210212-1101102213333321-0122100302133231-0013132220312113-3002130131102033-1102333103012302"></a>

<a id="canonical-0023033030232132-2223030031310322-0331111210020302-0321232331132002-1012230011013003-2100302330030323-2010300123000103-1213223122013122"></a>

## trusted_ca_url property — use_mtls / 300022101132 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-0331310322010022-2123322101030323-1033131223220203-3311111211121211-1111201013321032-1023110211223201-3303213023231031-0000011300332000): complete subsection reference.

<a id="canonical-0323000312231122-3231323311230031-3030010230031233-2221332331002332-1102030130033313-1221321313022111-0311212003001323-1322103131221203"></a>

## Next pages — use_mtls / 300022101132 / 6

- [proxy_config.https_auto_cert.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3313222100032011-0011230102312213-0323331102201201-0001112233133021-2012013122302032-1232112221001320-1332200323022031-1001022130332330)
- [proxy_config.https_auto_cert.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1132223211100320-0320223221120330-2102013330220021-3020323031113303-2230010013201301-2212331001230303-1122100012033200-3323302223222000)
- [proxy_config.https_auto_cert.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1013010021312231-0021102311120011-3010331030030300-3211301321200212-0113201301310233-3022320010012202-0131320331320203-2310112001211200)
- [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002)
- [proxy_config.https_auto_cert.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-0331310322010022-2123322101030323-1033131223220203-3311111211121211-1111201013321032-1023110211223201-3303213023231031-0000011300332000)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3313222100032011-0011230102312213-0323331102201201-0001112233133021-2012013122302032-1232112221001320-1332200323022031-1001022130332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000310032033211-1302130202301200-2133212310122021-2101200023101222-1032231103311331-0223010302020033-0333022131230120-2201101110101233"></a>

## proxy_config.https_auto_cert.use_mtls.crl — crl / 013221020230 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-3122100112000233-0311120333110101-3212101321123231-2202010132012331-0033101022131030-3300201221020030-1123330330203132-2322022110013331"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2200030301321321-2103232102001003-0330100222100000-0203311103212032-0303311023311113-3112310232000002-0030321230220323-1300121110002331"></a>

## Direct properties — crl / 013221020230 / 3

<a id="canonical-2020200021202132-1120122212130130-2012232132011022-0230032223133222-0130100331232202-1022210310210301-1033322010300112-2203131020201312"></a>

<a id="canonical-2003202130320030-3332100001002321-0121303010311322-1330101233211231-1301321302013122-3000022003030201-0310003020220122-2131001033223123"></a>

## name property — crl / 013221020230 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3113223212103113-2323230221002121-1323233001011113-3211313330212202-0133010011102231-1213013021010331-1032031303233310-0110103030221332"></a>

<a id="canonical-1223131222210020-0003011323210322-2231221023002112-0322330331302212-1131301313033000-0010002220210221-3233102113310023-1132021020101212"></a>

## namespace property — crl / 013221020230 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1112220102103213-1331230331232231-3310220330202123-0322330001033112-3303232111203320-0211320203223331-2013020112232213-0210101232000201"></a>

<a id="canonical-3313131010000203-2111302321023131-0231011310220013-3120021022001130-1331021132012121-1012332102331032-1121113311220303-1332111003002021"></a>

## tenant property — crl / 013221020230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3212311203300012-1332323223200202-3032213213102122-1311020220213230-2321131233011221-0132011120112133-3033223013313220-1233300033010020"></a>

## Next pages — crl / 013221020230 / 7

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1132223211100320-0320223221120330-2102013330220021-3020323031113303-2230010013201301-2212331001230303-1122100012033200-3323302223222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031210120332003-0133323112102021-1200311013020321-2312013212010212-3211010313033210-1210010323110202-0030032200002012-3001111002331210"></a>

## proxy_config.https_auto_cert.use_mtls.no_crl — no_crl / 322121233120 / 2

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

Terraform syntax:

```terraform
no_crl = {}
```

<a id="canonical-2310031233113122-0103302123022211-3230013033321220-3133320303223102-3112112331331012-2331313321320201-1312221333313203-3112322333221230"></a>

## Direct properties — no_crl / 322121233120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003132001020102-2220210220021003-3030001112113010-0131032030230021-1320002211323330-2130210221123303-2113220230312323-1010220310033000"></a>

## Next pages — no_crl / 322121233120 / 4

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1013010021312231-0021102311120011-3010331030030300-3211301321200212-0113201301310233-3022320010012202-0131320331320203-2310112001211200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030331223002021-0002122313122120-3203202301213001-1323101220013013-0321020001113311-3302330303111312-3300200333001322-1231222010311010"></a>

## proxy_config.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 220333330003 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1022212220320001-1323001123311200-0203031020331131-3001031320123120-0300323201103002-0313013012200332-1010120203111302-0133232033220103"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0211030231320130-0123220130200330-1013203303122032-1100200333133122-1312200222212310-1333321212120302-1330210323002023-2130132130130311"></a>

## Direct properties — trusted_ca / 220333330003 / 3

<a id="canonical-3230313211001333-0030221213223130-3022110110310110-1221223302200311-1200131131202011-3131321031003321-2212000032123232-3003330200311201"></a>

<a id="canonical-0030323202011023-3213101303122312-0210002133312300-2330303000330210-1132321323230122-0031232321131313-2013231121221233-1222010233300001"></a>

## name property — trusted_ca / 220333330003 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1133311202323131-1203002110031302-1120023012122321-1023300201222322-3133101132003320-0033322023232100-0301032112323032-1003120013111111"></a>

<a id="canonical-1320201312020032-3023011211231130-2133330322330010-1023003030122013-3013003233131120-2233122210100130-2210002320203010-2203323111300230"></a>

## namespace property — trusted_ca / 220333330003 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0123122023022310-2233012221302221-2221032103100031-2103332233020111-1020110332013100-3233312023301213-3211333032133221-0123333033000232"></a>

<a id="canonical-1132123011210201-1010130001122321-2103122311312222-1132333211303331-1330202230200130-0230121233002002-0111132131021030-1210302202112112"></a>

## tenant property — trusted_ca / 220333330003 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2110311203023230-0222312132211202-3320013232211231-2301123121312301-3012312231130300-0001322012021102-3313210213330112-2002222233131313"></a>

## Next pages — trusted_ca / 220333330003 / 7

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322120031033033-2002213133102313-0110303100213313-1321033013002010-0111231012033003-0011112000111020-1103302201332221-3212103322011221"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 313133132230 / 2

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

<a id="canonical-0100333220011110-0322002302230020-3310002333022210-0010333222012310-1301102010320320-2000012030303133-3313333111223211-2020133301121110"></a>

## Direct properties — xfcc_disabled / 313133132230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003110002310232-0101102113013130-0233321330102033-3012102112303102-1133111110230311-0332101111313212-1032333123122133-1310013101100121"></a>

## Next pages — xfcc_disabled / 313133132230 / 4

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0331310322010022-2123322101030323-1033131223220203-3311111211121211-1111201013321032-1023110211223201-3303213023231031-0000011300332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130331232101302-2132022330102333-2321022102230012-3233311222321211-0320211133000232-3330310201220310-3300230312211102-2223312123022302"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 313000030011 / 2

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2202123110010020-0310200102110113-3120012011312120-1323223310000120-1011000203323331-2003001230203133-0301203112132213-0013103111331230"></a>

## Direct properties — xfcc_options / 313000030011 / 3

<a id="canonical-1112003113111331-2123221101023333-3233023323222010-3132322001133120-1111001230300223-2323310313221100-1201321031012002-1310113030111332"></a>

<a id="canonical-2010333312311322-2132311132212031-3303132223111321-2030300233121221-3200101232112112-0022121001121301-1120213001200203-0112210320220131"></a>

## xfcc_header_elements property — xfcc_options / 313000030011 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3111003232112213-0213323031003113-2011223031230122-1231032313223301-2102021221032322-1310011222123110-2012313211030101-3112232022221230"></a>

## Next pages — xfcc_options / 313000030011 / 5

- [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2132011332111202-2202113131221012-1320230122203302-1112213013012013-0132001232220212-2022000013130003-1333021301323312-2321113213223101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1231333103033012-2112201220312122-2222023320130000-2122221200023233-1122012332200101-1313233133310311-3230012111033123-0101131100232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003132121122130-2301011030132302-3132230021331210-1230222331113331-2322210123233331-3303033322002122-2110300321112111-0323202100210321"></a>

## timeouts — timeouts / 303203203003 / 2

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

<a id="canonical-2023223120322012-3233222302030001-3303110020023231-3231103222333133-0210111011030031-3221102221203221-0311230020021221-2112211312220031"></a>

## Direct properties — timeouts / 303203203003 / 3

<a id="canonical-3021002223001012-0202020332321100-1312202302020313-3011210111231030-3213131023322321-2103102303123301-0301001032030221-3230333202203011"></a>

<a id="canonical-0203210130031300-0000201323002321-1012000130022133-0121021330012131-2011122322322023-2111311023100100-0300301323001113-0230113300332221"></a>

## create property — timeouts / 303203203003 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1221203203020013-1111003330301122-3222020003321200-2021212331111321-2100010321032303-2233033222023113-0222023310011022-0331232212311112"></a>

<a id="canonical-0220113203011311-2212303112323003-0000220011200222-2130230332302113-0112003333312121-0103100221220331-1000102222203331-3313111330023223"></a>

## delete property — timeouts / 303203203003 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3111202213100121-2023201200020303-2203123210101010-0311230012132300-2031113223221113-2311030330020320-1102202032333323-3032023221011013"></a>

<a id="canonical-1331110213010202-3203231211320101-1132301131312213-2030111330110223-0333223113320032-1101232133222023-0120312313032032-3033000231111323"></a>

## read property — timeouts / 303203203003 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1102310223003003-3233021331311302-3201131102023300-0121100132013020-1203033102020030-0001010201201022-1132121332021222-3103333330003310"></a>

<a id="canonical-1120013110133010-2331202121301123-0102201133103322-2323213113320311-3221001003010320-3011010033102021-0100313112222012-2220103130001123"></a>

## update property — timeouts / 303203203003 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2123030111210101-0030310102123221-0121330300012123-0302301200122203-0331101102023233-3113020213300310-3220120330131130-1033120232302223"></a>

## Next pages — timeouts / 303203203003 / 8

- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
