---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-3323333331121130-0101330023022013-2001210012100000-2022000100121200-0310131122310202-0101031010201313-1203001013330301-3133120100203002"></a>

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

<a id="canonical-2033031112330320-1122111033102122-1121331301230123-2311011200032212-2311023320212331-3002312213112330-3222033123302311-1231210210120021"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls`

<a id="canonical-0223021201013221-2100222321101211-0212033301233020-3300102302231001-3020001133101233-2200013120303212-0212310202313002-1311213003002001"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310221313023321-0120200010103302-0013013300313200-2100011310201311-0213220213000213-2212033212203220-1331210103023001-0322000010301330): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0012222132331101-2312210321320022-1002303101213002-2320123030211333-3332021112023303-2203130332133011-1100332132002313-2332111000132322): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1323013112333303-2220322000200100-1313103200210002-1101023103223001-0130031033110100-1132301003102203-3222321102111210-2303322212012013): complete subsection reference.

<a id="canonical-3110001310323220-3130221312113321-0010222103233203-2221133032231331-2013020333121313-2033130232023133-2203102323220113-0211111022323322"></a>

<a id="canonical-3210203111020110-3001120023323002-0230000321033202-2321112213221303-1031030001112211-2202020303311220-1102111111333313-3113112101221201"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1121132130223121-1103202113102313-2312023321211323-3130331013222023-2133111021003202-1100011201021312-1200110012001001-0103033211131010): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0220112023033102-0311310031331012-1221203323231030-2202230310001320-0232220222120003-3233120103310003-3302301023113023-0203232112121210): complete subsection reference.

<a id="canonical-2310221313023321-0120200010103302-0013013300313200-2100011310201311-0213220213000213-2212033212203220-1331210103023001-0322000010301330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-2013230130231331-3111213220120122-1313330120031121-3122113321103211-3032213332212102-2211011113012213-2011330023021230-1231311103321013"></a>

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

<a id="canonical-1212302202210233-3011113031111031-0220310213203033-0310302300012232-1223301222202302-1123033113023333-3010121022111032-1302123301021213"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.crl`

<a id="canonical-1000220020200012-1003223023321320-0103022332311222-2032202232021003-3222030213011123-0020020213120103-3110002330030210-1302002311000023"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-2031100130301331-3300031331121033-2200220010131121-3122323222123120-1303201210332213-3121202201032201-2300200301030033-0021221312101001"></a>

<a id="canonical-2220210103332210-0333100332323313-2232112310302323-3103320303100222-3113133102011123-3330032223023332-0112302102021013-0033230322212322"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-3303213323110123-3232230130022133-0322003010110203-1213330313232020-2121013100320312-3313122301211103-1031003131311232-1003100213220001"></a>

<a id="canonical-1102110020021113-1110233321221213-1221302131122231-2310322322110312-2031101032010202-2211323321300123-0331030201103222-1320230022113113"></a>

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

<a id="canonical-0012222132331101-2312210321320022-1002303101213002-2320123030211333-3332021112023303-2203130332133011-1100332132002313-2332111000132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3220230133002012-3023132332222313-3220100102330013-2300120223232302-0020132312130331-1013233322312203-1313301012222031-2021220011013031"></a>

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

<a id="canonical-1323013112333303-2220322000200100-1313103200210002-1101023103223001-0130031033110100-1132301003102203-3222321102111210-2303322212012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1320322012231103-1013213030132322-0101012012332221-3102031120102301-0033003121300312-1021302301331323-2311302232320021-2320011202013231"></a>

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

<a id="canonical-3212312131222023-3200210102122010-1202123203321210-1300212201300103-2131113220133311-0331100020322220-2012001303320311-3003211331333311"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3321312102123103-1121310202312232-3032023022101320-3300103312022211-3112323132131321-0023200103112102-3121113030303321-1120101230302221"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-0133201231011233-1110333033023312-0221110313013023-0003201011032021-1112232010112123-1312221202310113-3213020213100002-0331201030101132"></a>

<a id="canonical-2323001011321313-3121220021030012-0333232230231310-3303100310112203-2021213303323032-3011232333022132-2031122011113030-0232230331323220"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1110030112223112-0310110120121200-3330030101000020-1130311213301203-0213323213001013-1131212121112123-1032231210320302-1023111232323032"></a>

<a id="canonical-0031002232233030-0310203302002333-3331202000121012-1303233123232310-1120300333202112-1121312301000013-2123303003101102-2221212200003030"></a>

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

<a id="canonical-1121132130223121-1103202113102313-2312023321211323-3130331013222023-2133111021003202-1100011201021312-1200110012001001-0103033211131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2330113221112113-1320323320130300-3313132313032302-3312003333210001-2220001302302011-3003022122212333-2121133122130332-0210201001032031"></a>

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

<a id="canonical-0220112023033102-0311310031331012-1221203323231030-2202230310001320-0232220222120003-3233120103310003-3302301023113023-0203232112121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0333001012300132-2333031203202200-2011203013312100-1320310020101130-3201013330100303-0311122213123113-3101332233200123-2223302310111222"></a>

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

<a id="canonical-1031331302121302-1333230232123000-1112120331213001-2001122121203233-3121132021012212-2331330331323233-1231132202212212-2331012230220222"></a>

### Direct properties for `proxy_config.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0131321223303031-0133121131121120-2230320210110233-2202112121130211-3310323101220121-0231121311121001-0200301133333020-0122310320232103"></a>

#### `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.tls_parameters

<a id="canonical-1201123003111312-3320313130233200-2321331213113100-0003203220311322-1320311011310223-1033012213322111-1232231132122133-0030211221111110"></a>

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

<a id="canonical-1011102022310221-2002332311211202-3311000130113232-1311222201130112-2010322330000302-2313321131132132-3001122202230212-2331212011300303"></a>

### Direct properties for `proxy_config.https.tls_parameters`

- [no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103211131002331-2002013001220022-2320310201010130-1010003223300220-3103100013213202-3013320310020003-1001102312233132-3232203210132002): complete subsection reference.

- [tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100): complete subsection reference.

<a id="canonical-3103211131002331-2002013001220022-2320310201010130-1010003223300220-3103100013213202-3013320310020003-1001102312233132-3232203210132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-1100013033220213-2233310102033130-2110202321202203-0212132030023300-3002210011323221-1030212023013000-1233010010213232-3022110103121021"></a>

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

<a id="canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-3320323100113312-2220312130011303-0123030113323211-0031222103100213-1032213130131310-1222222303020021-3221203123122013-0031331013033103"></a>

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

<a id="canonical-0202021203001012-1203120300122312-3302233231321321-0301102130331111-3113211221322130-3213312230110103-0110321233200210-3132333101000303"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates`

<a id="canonical-0303030323131121-1333000132000120-1111313300303203-1033022133313022-3300212333011333-0323130003312211-1023202110212233-0133202001000132"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2011203310232100-2112322003133230-3222313123011110-0213232303232213-3010110112021031-2112233010302131-2222021030010301-1231111230333102): complete subsection reference.

<a id="canonical-0100230010000020-3010103313122030-3131230202203330-1131102001303213-1320021022133103-2232121013030232-0132212131231212-2100003332323300"></a>

<a id="canonical-3110131132313113-0222303310031113-1333302313130011-0332023130213211-0232333001220203-3323003212330311-3103010313303230-0021232002232230"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1330011020332031-1220021211113001-0032321202221101-1303313333203300-1233202302310313-2202123030110222-0202120123130330-3230310122210000): complete subsection reference.

- [private_key](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232): complete subsection reference.

- [use_system_defaults](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3132031010020200-0021303130021321-1210330310132010-1311331201300313-2111210101232023-3213113222211121-3303231303200301-1322122102032001): complete subsection reference.

<a id="canonical-2011203310232100-2112322003133230-3222313123011110-0213232303232213-3010110112021031-2112233010302131-2222021030010301-1231111230333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1121011121201310-3022001113102121-1313021322231002-0023013031111212-1122011302123011-1131302130233032-2222311120022111-2013003023101212"></a>

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

<a id="canonical-3103202013030330-1301231013112130-1300112132202122-2312220002012332-1002132201022010-2022232310000101-1302313222230022-2320312310233320"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-2200020112131100-2031330330022003-1030303013233020-1222233033031010-2232303233121010-1132220102013113-3023303112330023-2231011330123321"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-1330011020332031-1220021211113001-0032321202221101-1303313333203300-1233202302310313-2202123030110222-0202120123130330-3230310122210000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3112011131032311-0111312110200132-2202003321313111-2101210131233203-1213033313000333-2132032011010300-1131132021300230-1113230332332201"></a>

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

<a id="canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-0032223323211231-2113320120213000-0111230213022012-1220010202312311-0313123311113020-3013021103123033-1111221300031323-3232030331233301"></a>

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

<a id="canonical-2102200302110003-2231033003230300-2312332022003301-1130232003320332-3132131223233102-3123200320101113-3113310020132310-3233213001112320"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0131001320232323-3330311030331133-2031230303203103-2112210013223332-2300330322032011-2110201333113021-3131320322203220-1010222132020223): complete subsection reference.

- [clear_secret_info](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1312333132310200-0103210210300030-1222330122312223-0132012200110323-1131102112223003-2130201033122030-1120331213303233-1313223330302311): complete subsection reference.

<a id="canonical-0131001320232323-3330311030331133-2031230303203103-2112210013223332-2300330322032011-2110201333113021-3131320322203220-1010222132020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2200023321001131-3311220312221020-1111210312231103-2011132202220311-3120213113122313-2020312010101211-2223220022021331-1300312002130220"></a>

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

<a id="canonical-0000310201232212-3223102000221220-3023301221022222-0001230002122331-0331010102210002-2232301102103230-0220103212233130-0102223031231022"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3203210032121330-0230123312002310-3132321111223331-2132203301333223-0000322102001232-2333100313111131-2110320130312331-0230331103121012"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3133203300202213-2010110021031120-2310111103022013-1323330032110333-2212011002120023-3032333113111010-1331223321113011-1100132330230011"></a>

<a id="canonical-3122332330020232-0322032033332330-0100320011121113-1231331010001100-0133013223201121-0030333120031022-3103103333210032-2330221212101212"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-0003033313021112-1122233222001103-2123023031212310-2312121130212102-0112122323100230-3100021033101012-2022331230120120-2323320112301022"></a>

<a id="canonical-1311213222321313-3120201321333233-2220100101200231-1233110122102223-3310023020032020-3003210301033330-3030233032121301-0223322301103102"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-1312333132310200-0103210210300030-1222330122312223-0132012200110323-1131102112223003-2130201033122030-1120331213303233-1313223330302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-2203023003102211-1300330303212210-1122003002201130-3121000201231202-1130012303321003-0313301130211310-2232130102233210-3023203302023122"></a>

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

<a id="canonical-1010333303310130-3101130313203213-0330213100300301-2322031120331202-0220221123213133-0022320000213200-3323330333301101-0103023030021313"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0212003223201121-3213222022213233-3101233023311001-3230213033130132-2300211102020031-3130000231303130-1201202331123231-1131133303033012"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2133231200012200-0102101010102020-3221030200001102-2323110103000120-1010100031101002-0332010003130021-1220011022133310-0020200023312210"></a>

<a id="canonical-0131300131331032-2002231023123023-3011123212220302-0310201202001103-2330102232000222-2133133112001131-0300112312331021-0020223213011030"></a>

#### `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-3132031010020200-0021303130021321-1210330310132010-1311331201300313-2111210101232023-3213113222211121-3303231303200301-1322122102032001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2131012012330020-2201200032231133-2222102001211331-0131002102002102-3033021312133232-2013003331001100-3033000320101203-3030003312221022"></a>

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

<a id="canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-2121230122113200-0212020330310131-0310333011030132-2120020003233100-2213122010011030-3120303122121221-3023023123230012-1233211002133223"></a>

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

<a id="canonical-1102013203132030-3312103200202123-2311312301302032-2211320101212302-2303312123200113-0221013221133210-3102331001121310-0032202113032210"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_config`

- [custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0330120033123122-2112211002012033-2303001101101321-3031102232122003-2132013101320230-1032121303230203-2320031322320132-0232011311021220): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0332311203233100-1201333122301330-3002021230212111-3303122103203312-1201220101000032-2202002330111331-3221302003313312-2111021121310130): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310021031211322-3303021001013103-2113023211023302-2321110031011133-2012021320132323-2321232102013100-0102200213333221-1200311321111231): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2311030031330313-2102313002123011-3302012221210203-1303232331330002-2102003202300333-2103311033100210-2032330201201000-3203222122103120): complete subsection reference.

<a id="canonical-0330120033123122-2112211002012033-2303001101101321-3031102232122003-2132013101320230-1032121303230203-2320031322320132-0232011311021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-3122231310302001-2030100011103000-3010312033223302-1100113123132312-1011022121100210-1221222110213221-3020301332233102-0130122203120331"></a>

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

<a id="canonical-2301312131031120-1223011232200301-2201122103331212-1320232133331032-0330322331001033-1203130312310123-0322002331221301-2311130233311330"></a>

### Direct properties for `proxy_config.https.tls_parameters.tls_config.custom_security`

<a id="canonical-3031001212333300-2013213133112030-0332131201101222-1333333211220120-3121010210303133-0003220000111020-0120322013113321-2302123221231201"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0020212002313020-2303323230131023-1331203132000210-0001322003331313-1320101023232123-3130033320301020-2332303322303022-3101112300302220"></a>

<a id="canonical-0033101120133222-3232002303323122-0123303011000201-0211103112013232-0022000002130022-3020101221002020-0201201001202100-1330011021323201"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-3321200203313332-3032003010003302-0322331010323230-3212201330231131-0131230201223001-3320021113230012-1202101021031310-2122132232102012"></a>

<a id="canonical-3022321102212332-3322032330002311-3030203012232300-1220002332110303-3200333112133012-2332313133133120-3032021131231001-1202132120322310"></a>

#### `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-0332311203233100-1201333122301330-3002021230212111-3303122103203312-1201220101000032-2202002330111331-3221302003313312-2111021121310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-2012022031013200-3131321000000311-0031322211310113-0010013021013330-3331320322210131-0320123231201133-3312120321001331-2301230233202320"></a>

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

<a id="canonical-0310021031211322-3303021001013103-2113023211023302-2321110031011133-2012021320132323-2321232102013100-0102200213333221-1200311321111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-1220231112001133-3313001100211310-1222032222212112-1310102202331033-0112231212113212-0212203020113330-0202202232232102-3110311223112322"></a>

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

<a id="canonical-2311030031330313-2102313002123011-3302012221210203-1303232331330002-2102003202300333-2103311033100210-2032330201201000-3203222122103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-3011212221001302-2002211233102332-3221321301030222-0320210030023333-3203313003301312-2111330333032013-0120230310133201-0230331011312221"></a>

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

<a id="canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.use_mtls

<a id="canonical-1210122230200233-0020102033331321-2100213011212200-1120102233222210-1232312111302212-3132231320130111-2233302211010130-3030021333331202"></a>

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

<a id="canonical-3312300303312231-0011202321001132-0113030331133100-0322131112123100-1323031302212001-0013210010213301-1220033310201312-2323312323331113"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls`

<a id="canonical-1222221102032102-0301010230213203-2303011321120022-1101221113201102-2132230111113303-0001210203101302-3210222100021002-3320121231111223"></a>

#### `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103023323301323-0111023033110200-0300023121212110-2202323211103312-0231200012231003-3222031211023230-1333111020132000-0323223013202100): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3213022221230221-0333221301301032-2332313122021221-3130323103301132-3232120330210301-2232300001123210-0320111231330122-2311033032123311): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1113323230002302-0300323000011031-0011032132221030-0301103302313220-1010122110211100-0111002312201230-1311100321003330-0010010103000312): complete subsection reference.

<a id="canonical-0203331203230323-2311010210221301-1222010030032211-1020233111321232-3110103023100013-0030203010101120-2132331002022002-1232310222320301"></a>

<a id="canonical-3330220201120031-0030022323031120-0322101302003210-3011222313331121-0033233023021320-3213011211330332-0213131202012300-0032003210000030"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2121221323123222-0220020011110122-0130021310122201-3232122323032032-2021130121222212-3323011112111123-1012303112132213-1212221210122231): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2101132012131003-1220132123220211-3310230112312103-2131222220222232-1310313020030023-3333310110133121-3333103030303113-3121103101020312): complete subsection reference.

<a id="canonical-3103023323301323-0111023033110200-0300023121212110-2202323211103312-0231200012231003-3222031211023230-1333111020132000-0323223013202100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-2230022002021313-3300233322301303-2032200013000312-2330132033211031-0033221033030130-1020332003331232-0230001121330011-0231221000230332"></a>

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

<a id="canonical-0201023122323122-1100322320232323-3030023020330322-1111302212220030-0030213212002311-2033322202323003-1323320212102201-1212120222233113"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.crl`

<a id="canonical-0020203023211202-3220002003002120-2310022121033330-1313122133320201-1000012210210233-0213122230113233-3213321202132320-2330213301012000"></a>

#### `proxy_config.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-2201210323033331-0013132010202030-1231130302231132-2032200110212120-0201113012203230-1001111200131130-2231033031210103-0001303021202012"></a>

<a id="canonical-1323021030010311-3322010212120323-1131203320010131-0103332012113030-1330313023202202-2230333000331311-2213133300321313-1120011101110112"></a>

#### `proxy_config.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-2230223213021202-3301210310212002-1221122122103210-2321131130300302-0003131223033113-3211211012313213-2333023013303312-3001020322321301"></a>

<a id="canonical-1211300103101030-2011003323031122-1322223312200222-0033322103203100-2302032012022301-2311213333330220-2203102200101301-0223203003011033"></a>

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

<a id="canonical-3213022221230221-0333221301301032-2332313122021221-3130323103301132-3232120330210301-2232300001123210-0320111231330122-2311033032123311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1231121101201303-3003131010302113-0032003033121321-1220003331002032-0032301230233303-2012103131020200-0011333213103311-3120222223101223"></a>

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

<a id="canonical-1113323230002302-0300323000011031-0011032132221030-0301103302313220-1010122110211100-0111002312201230-1311100321003330-0010010103000312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-3230120200323110-3201303223210202-2210012330010220-3200010210102232-2300332301222220-1202030213011321-2031023233012221-0131011113333200"></a>

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

<a id="canonical-1231120002021332-0100210222201211-0032310003103023-0102030020100313-2002201303333130-2101103231203020-2200023212303031-3202003302032302"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1330203203203102-2230031203023130-1320022200121021-0102001232031222-0102223121311010-2023010200020321-3232100323333121-3222021303121122"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-3033302021110310-2302133030033030-2302002103023332-3320232221310122-2311113302231212-1032233113110201-0100010202212223-3220003133011123"></a>

<a id="canonical-0012200002033311-2322033130302320-0311323013223110-1120023010313012-1223020013010222-1201111001310001-0202021031132233-3330013300133100"></a>

#### `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3222122133121300-0320032003132213-3100203220012032-1110213203000120-1201201011010132-2330000133311111-0100022211120031-3012301113130021"></a>

<a id="canonical-2101130311233313-0030323232330110-2223220203213021-3313131013322111-2113132131311022-2001012101301021-2220300310203232-2022120301010221"></a>

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

<a id="canonical-2121221323123222-0220020011110122-0130021310122201-3232122323032032-2021130121222212-3323011112111123-1012303112132213-1212221210122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3011121222000332-0101100213001002-2002033200210003-2112312223031322-0012102320011232-1322102030303020-1121133131130132-1132301201022012"></a>

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

<a id="canonical-2101132012131003-1220132123220211-3310230112312103-2131222220222232-1310313020030023-3333310110133121-3333103030303113-3121103101020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2330222222231133-3230331301000323-2323301001012310-2111002332001313-3220323200103211-2330322001330200-0111320200021020-1300032110001330"></a>

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

<a id="canonical-1230003331212100-0001002123013300-0230302310013002-1321003211002320-3101113311101212-1101322133031231-0130000310231102-2012200302230030"></a>

### Direct properties for `proxy_config.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-2212131331211022-0100321231223020-1101223300331010-0032333231222220-0012002210132301-1103221030320313-2310303232300201-0013002132213030"></a>

#### `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- proxy_config.https_auto_cert

<a id="canonical-0332331330031223-1311030322032013-0222323322323003-2023312200110121-1133222312231302-3220010201302032-0131122200321302-1310203032310021"></a>

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

<a id="canonical-1322233321013113-2223222302202022-0310320022131120-3310201303311010-3310011301011133-2033102011320200-1020202331323111-2331120013101310"></a>

### Direct properties for `proxy_config.https_auto_cert`

<a id="canonical-0101323301013302-1010300003023001-0220021310320031-2023030332213003-0333033221011211-3110323012321311-1222130013010031-2031020212122100"></a>

#### `proxy_config.https_auto_cert.add_hsts` property

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

<a id="canonical-1203321231323300-3331011032212001-0202030330003012-0303322322130312-2322203311312231-1030312130000011-3000331010010132-0132003200013330"></a>

<a id="canonical-1011231001032102-0231001311300013-1333002033301330-1222131212203210-3122132002032323-1112223322232322-0000112111023112-3330331011021323"></a>

#### `proxy_config.https_auto_cert.append_server_name` property

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

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030): complete subsection reference.

<a id="canonical-1003203231021222-0333313032131323-3303232231021212-2232302002012130-1203103122131012-0333220033132332-0302233313103003-1011001032311300"></a>

<a id="canonical-3112102222203003-3021200302232302-3131203300313001-3323003011130200-3000002020121310-3113023121222211-0013123203002201-3132330323132310"></a>

#### `proxy_config.https_auto_cert.connection_idle_timeout` property

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

- [default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0030300011320301-0132002031320201-0031231310003021-0011110133003223-2321220231230203-1323311203311323-0210211322313002-3110003101030223): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2100311102133131-0103333221332103-3011131011012030-3122233311133132-1021210011010121-1302000133330301-0312010300011033-1301010332213133): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1013121000213233-3301313031232112-2232012301322202-3300110222313112-0332302333031133-3213331200001300-1113123201201112-0210100233000201): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3113213021003302-1131212001221030-0002232330203000-3121013311103022-1330233330212100-0003032023030012-3102022301131101-3111302200302101): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111): complete subsection reference.

<a id="canonical-3203312331212003-2300330222012020-0030232232022301-3211311033122300-0021103010012303-1131331202101032-1113220010122202-0320300131003132"></a>

<a id="canonical-3311112130020031-0312233301332213-0001030002203220-1320200012100103-3033000031130022-1313122213311203-3330221203133310-3133330011223303"></a>

#### `proxy_config.https_auto_cert.http_redirect` property

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

- [no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3032101132210133-0133123021311130-0333100232101300-3110202321302120-0113332010033110-0333001021210231-3102100200000002-3211302011322302): complete subsection reference.

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0302113122013130-2201130101123011-0221313223300113-2322200323232133-0332312100030303-0313111201230213-0020330003231321-3233013100222222): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1222113032213003-2332033231032202-3331011322132003-2303223012013220-1002313103320300-1032332322010113-2203130032100331-0312223303020013): complete subsection reference.

<a id="canonical-2331203313130010-2313112233233323-2133212012233133-1131113202323321-1302302231312313-1032030010312012-2322210330311321-1320033130131200"></a>

<a id="canonical-1121010311332113-1102003100321030-2201112001112003-3210202120133223-0123230333031030-1122322023030110-1101031213312033-1122301201103020"></a>

#### `proxy_config.https_auto_cert.port` property

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

<a id="canonical-3213212202202300-0212020001210000-2223211300031100-2313000101013332-3330201111233103-2012033212202132-1321021131012332-0301321111213033"></a>

<a id="canonical-1112011333311112-0321333202301133-1022113111322232-3322130003322121-0201012322122330-2230323001123112-1223120020200213-0113203220320320"></a>

#### `proxy_config.https_auto_cert.port_ranges` property

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

<a id="canonical-3003332330013122-0121020201022010-1313321102130303-0003120130002232-0003201002130220-1003312220213232-0311333211210230-0022100001330011"></a>

<a id="canonical-3012011103221030-0031310323132010-3032222223310113-0202323220331100-3312133132331232-1310120201103323-3133130203131212-2021212230023122"></a>

#### `proxy_config.https_auto_cert.server_name` property

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

- [tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200): complete subsection reference.

<a id="canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-2323231130323012-3210230322131303-2131020020023330-2010232203322211-1223302330003001-1313310012012201-3311112113102330-1312332233123222"></a>

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

<a id="canonical-0311130330301203-2010230223313132-3313213333233032-1320103022221221-2303313302120012-0202102032123103-1133022223031121-0113313132123332"></a>

### Direct properties for `proxy_config.https_auto_cert.coalescing_options`

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1201231022323222-0312232112121122-1313113331001030-0113111213331331-1003121332211030-1003311101231233-0101020121023013-2102030233302313): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2231133130101031-2100231030221323-0030010033111001-2103122302302312-2122022323011230-2320223210213300-3223001102311101-2133113311001003): complete subsection reference.

<a id="canonical-1201231022323222-0312232112121122-1313113331001030-0113111213331331-1003121332211030-1003311101231233-0101020121023013-2102030233302313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030)
- proxy_config.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-2200130103100121-2100303223133210-0333032310003323-3233313331023211-2221220310302200-3203221032100021-2220030013332003-3211101311033033"></a>

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

<a id="canonical-2231133130101031-2100231030221323-0030010033111001-2103122302302312-2122022323011230-2320223210213300-3223001102311101-2133113311001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030)
- proxy_config.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3112120002223131-3032013233302013-1232300130221102-1313131310201013-0222323013223331-3210233333330120-3000130231102030-1112111032310000"></a>

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

<a id="canonical-0030300011320301-0132002031320201-0031231310003021-0011110133003223-2321220231230203-1323311203311323-0210211322313002-3110003101030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.default_header

<a id="canonical-1220221311011023-3020210333333021-2113131331233221-0203100122313022-2112230013331323-1002110332332330-3313130123112130-0030202222200202"></a>

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

<a id="canonical-2100311102133131-0103333221332103-3011131011012030-3122233311133132-1021210011010121-1302000133330301-0312010300011033-1301010332213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-2322230011133301-3212102133120320-0022102002213201-3200230223231301-2310001331003030-3110323131321221-1103103311313111-3030000310022331"></a>

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

<a id="canonical-1013121000213233-3301313031232112-2232012301322202-3300110222313112-0332302333031133-3213331200001300-1113123201201112-0210100233000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-3331312223213323-3320203303333303-1222130303122200-0331230031023230-0113231210131310-0120331202022003-0221233022310020-2210101021101200"></a>

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

<a id="canonical-3113213021003302-1131212001221030-0002232330203000-3121013311103022-1330233330212100-0003032023030012-3102022301131101-3111302200302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-0313012103101132-2312223303001101-1000003001312033-3320213321030033-2231131202213121-2122120221210311-1332102323100322-3332033132301313"></a>

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

<a id="canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.http_protocol_options

<a id="canonical-2030203200311103-0133102323320210-0201213133220333-0020301030331220-0003331022210023-2203110200331230-1201112300311021-1202312002302310"></a>

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

<a id="canonical-3123012132010302-1231130103031320-2112121200333320-2112310232310212-1013111110120302-1130201122301231-3302012202213320-0332221131033320"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3000323121202033-3300212300023110-0111212212200013-3313101003222032-0103303331131032-3122111233120203-1211302210132103-1131222202033033): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2110332203130223-1121033002100221-2100112000011200-3300022201202021-0320222331231110-0131110232302203-2223203300010021-2211113311233003): complete subsection reference.

<a id="canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1331332201222002-3312323111030002-2120011202220310-0322330302202103-0312112130220220-1300311332302222-0222013121131231-2211130120301010"></a>

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

<a id="canonical-1023333302031200-3331103111222100-3300102331220030-2230113122222322-3030230003103300-2003020312110202-2131210001203312-1230132031130123"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200): complete subsection reference.

<a id="canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0003311002101100-1112112231020212-1022212122032000-3112322331022233-2311322222102220-3331331001110222-2023221223112330-2121322022233002"></a>

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

<a id="canonical-0301122012130022-3231121021302212-0100221001202000-2030131000323201-0130023311221100-1301102312333313-1103220332302320-1133102322231013"></a>

### Direct properties for `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2023020201220130-2001110312010022-0301020130011113-1220312330020230-0300320230002333-3030223210302222-1210022120330030-0101002200031313): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1223312000323310-2123213220002032-1132333311221100-2203010220321213-3122013120301331-2012011131213223-2230130131233220-2113302220101211): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2312321333331322-2330031223122300-0212202121223330-0300212321112131-3301013033030330-0120331011030202-3311122322321220-3121132032111133): complete subsection reference.

<a id="canonical-2023020201220130-2001110312010022-0301020130011113-1220312330020230-0300320230002333-3030223210302222-1210022120330030-0101002200031313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1213220132133022-1002103301110233-0111221200103023-2132212230123013-3222133222202130-1311231203310010-0111331312011013-0333102333003020"></a>

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

<a id="canonical-1223312000323310-2123213220002032-1132333311221100-2203010220321213-3122013120301331-2012011131213223-2230130131233220-2113302220101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2211132310132032-3210302312232303-2111120201120310-2110221031020002-0003323021022230-2302030222300022-0020322322210132-0303322023213032"></a>

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

<a id="canonical-2312321333331322-2330031223122300-0212202121223330-0300212321112131-3301013033030330-0120331011030202-3311122322321220-3121132032111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0303100111100220-1333222222210103-0322123122232301-0003101101231223-2020130210113323-2301121121300220-2023323003020222-3233113333120301"></a>

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

<a id="canonical-3000323121202033-3300212300023110-0111212212200013-3313101003222032-0103303331131032-3122111233120203-1211302210132103-1131222202033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-1030102100121011-0012313131221311-2031323100010100-2303312101120122-2032030013033033-0331212323320203-2031333302101222-3330033203113300"></a>

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

<a id="canonical-2110332203130223-1121033002100221-2100112000011200-3300022201202021-0320222331231110-0131110232302203-2223203300010021-2211113311233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3300322133202300-2323100023112013-1011223201231202-3012011200223203-1102232032300101-1332312201212322-0200301112122030-3302330023302201"></a>

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

<a id="canonical-3032101132210133-0133123021311130-0333100232101300-3110202321302120-0113332010033110-0333001021210231-3102100200000002-3211302011322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-2330213311302231-3100230321010321-3322302022301123-2320221211201113-2220033013012303-0213130013021130-1302021320113122-0133232310311020"></a>

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

<a id="canonical-0302113122013130-2201130101123011-0221313223300113-2322200323232133-0332312100030303-0313111201230213-0020330003231321-3233013100222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-2312110211233330-1111110100332323-1330333030312102-0032001312103331-1110002032230222-0202202133203001-0331231133311221-3303102222203313"></a>

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

<a id="canonical-1222113032213003-2332033231032202-3331011322132003-2303223012013220-1002313103320300-1032332322010113-2203130032100331-0312223303020013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-3012320311120332-1323300231031113-1122203313103032-2223223202311112-0022323113312133-1033221032323132-0300331112021322-3302211223302133"></a>

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

<a id="canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-2203113220010212-3033312212222231-3030031010323303-0300130201113211-1213023020013202-1021320131313003-1123021110133103-0222202300112121"></a>

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

<a id="canonical-0010121233210202-2333131232132110-2202231110001330-3231013123123031-0333032032312111-3130022322113310-0131301301213233-3330101130133202"></a>

### Direct properties for `proxy_config.https_auto_cert.tls_config`

- [custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0300221023003133-1200321233211220-3323030113020330-1212112112121032-0013221230033133-3021120002020210-3311102202011113-0331111212311103): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1012133112021320-3130100100122231-0211230002020233-0322232323210330-0012101220010231-1110232031030032-0312102221101331-3220223232220000): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3102332131122100-2011331213122200-1312320233130232-1212312131030131-3301302201303212-3312302200112230-3322310221123210-1332112122113113): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3011030212111000-0031311313020310-2303131101210302-0133230212310003-1130220313222013-1112133232232012-2321200001333021-2000302301312111): complete subsection reference.

<a id="canonical-0300221023003133-1200321233211220-3323030113020330-1212112112121032-0013221230033133-3021120002020210-3311102202011113-0331111212311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-3001333321120210-2332221111031121-1023011233212011-2011302311220001-0111020131102301-3020021302323110-2020200333322330-3001322213013310"></a>

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

<a id="canonical-2332000330132210-2323023001201002-0022230110213303-1301330110130221-2201121102112211-1001333120212231-0323303101313311-1000312130123123"></a>

### Direct properties for `proxy_config.https_auto_cert.tls_config.custom_security`

<a id="canonical-1222331312321333-3320131122110231-2133110303103211-3010332021102321-3313201032220230-1201120021113232-0010131032221120-0132212220133001"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-3332201301220022-1021020210122221-1132223303332131-0332120102311302-3333023203122002-3230010330033320-2001101010100012-3002021011123020"></a>

<a id="canonical-3120312003101322-3220222231220320-3301132321021133-0313333021031300-0130101213102322-2210112133033023-0130321010010322-2230213103211111"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-1221333111022133-1022311032122131-1030232203331120-1220022011021102-2022302233020333-1130102321203211-2312030322103123-3210202301232310"></a>

<a id="canonical-3033213331030333-1113133013111130-1122003131303003-1123313321120111-3202013102011312-1100030213320211-0110000321032000-0232000313012223"></a>

#### `proxy_config.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-1012133112021320-3130100100122231-0211230002020233-0322232323210330-0012101220010231-1110232031030032-0312102221101331-3220223232220000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- proxy_config.https_auto_cert.tls_config.default_security

<a id="canonical-1130311323123333-3312300113330031-2021101101210033-3012221101100113-3211013031122301-1101003131212121-3103131100321303-0232312103303001"></a>

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

<a id="canonical-3102332131122100-2011331213122200-1312320233130232-1212312131030131-3301302201303212-3312302200112230-3322310221123210-1332112122113113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- proxy_config.https_auto_cert.tls_config.low_security

<a id="canonical-1101303112301021-3003211301010330-1331211211103300-1312221123010220-1123222332332122-0330012110022103-3223320133212100-2021210321121021"></a>

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

<a id="canonical-3011030212111000-0031311313020310-2303131101210302-0133230212310003-1130220313222013-1112133232232012-2321200001333021-2000302301312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- proxy_config.https_auto_cert.tls_config.medium_security

<a id="canonical-2123133032132331-3121020200133302-3302122031201012-2032303132033312-0200312103233303-0112333100321030-0222123100023210-2113320213301110"></a>

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

<a id="canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.use_mtls

<a id="canonical-3100112110323131-2322332323120032-2300211010031102-2000232020301131-2121021022120332-0330023300120121-1232010210001300-2132000113320332"></a>

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

<a id="canonical-0000011113202112-1321110121313212-1033133231222113-1122130100213232-3232030113311212-2132303230302011-3223331003031300-0011121222200332"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls`

<a id="canonical-2303031233311231-2200311010130121-1131131000032033-2110020320303110-0132101101323320-0120003330303132-1320101120130002-0113102232120320"></a>

#### `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3332122033322000-2313321210232312-3200030030003012-0222213102223031-2211030021311203-1223010001111230-2020203011332033-2201033121330200): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1122203310210213-0031221303120111-1302322221103122-3003302122232300-2230003233010232-0203110210310131-0022310112333330-2133101311301022): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1302111123333232-3011131213310300-0331132203231312-1203223302320231-1022311022023022-2031003213000112-0131321033010003-0130322013322132): complete subsection reference.

<a id="canonical-1030320220010232-3333301023100212-0122212210330310-3131301312300210-2312122020022010-3110033122333220-0020332233021123-2113001100031111"></a>

<a id="canonical-2021110000122212-0222033012130110-3113022200110122-0100331321222222-2203211213002202-0322312121000112-1231121122330201-2233133030310022"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1321321031020011-2210123111113130-0130130013323122-1233233103132323-1233313132023120-1012003002111120-1011033030302312-2110121301333200): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2220211130233202-3101113103101320-0220333210010111-1011111301312122-3220111122132123-2003001230321300-2020020023133202-3203110310202131): complete subsection reference.

<a id="canonical-3332122033322000-2313321210232312-3200030030003012-0222213102223031-2211030021311203-1223010001111230-2020203011332033-2201033121330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-2230220111033113-2111323130131011-2211003213102120-0333230030303001-1112330011122010-2300001321200212-0112202303123022-3100003220231113"></a>

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

<a id="canonical-2231023311331032-2102200303100022-0302210023100301-3001223303310002-3123211123302323-2113031030010320-0103301121111221-2333202012111332"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.crl`

<a id="canonical-0300300001310211-0223302321020330-3130203110122012-1012220020332232-3012031113031300-2130323113010233-3211031320233020-1112313002220021"></a>

#### `proxy_config.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-0000222033002333-1200323302013222-0220313032232320-1203030113113111-3131002130030033-0323200202200301-0211011312112112-1232222100122212"></a>

<a id="canonical-1000020302011233-1012100031001131-0002013033230233-3131002013313111-0111203013332230-0022301103301120-0301110122232330-1032332112320131"></a>

#### `proxy_config.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-1210130312000102-0201213020301100-3321211303002202-1030031220230010-1023130310133302-0212123221110133-3303200110310230-2101032013222232"></a>

<a id="canonical-3011202201210120-1232301320112112-3011301220323101-1001210332000203-0322232331032001-0303210000003223-1223131032101200-3213123300321001"></a>

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

<a id="canonical-1122203310210213-0031221303120111-1302322221103122-3003302122232300-2230003233010232-0203110210310131-0022310112333330-2133101311301022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.no_crl

<a id="canonical-3001321301202301-0032033230313231-1220013022033001-3120332212120311-1111332201322111-1212202302232033-1233113321122111-2132023310021322"></a>

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

<a id="canonical-1302111123333232-3011131213310300-0331132203231312-1203223302320231-1022311022023022-2031003213000112-0131321033010003-0130322013322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-2000202221311312-0023123202220111-2330023030332203-0211131310322113-1202213100003130-2003103030033120-0112333330023023-0002231230020130"></a>

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

<a id="canonical-1102301030110333-3323320001332222-2121302120030102-1232333323130021-2013002001313322-0121321130012230-0301121030022330-1112230123301133"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-2100012013213033-1123120032300232-3000123112232012-1311122121223213-1120010012000011-3210113103112332-0021002101312021-0123203231220201"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-3323201320002212-2101211001123202-2222112111103102-3211100102100001-3201133102231032-3122222101121022-3313011032102321-3111223023321112"></a>

<a id="canonical-2322000100321232-2122100123003111-1133031031112023-1120202010020222-0021303201220302-2211303110030321-3113110213320311-1303011010021011"></a>

#### `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0000300110313023-3110122232233002-3031222220122221-0212110330101222-3203032223223230-2001320000312010-0332200222322123-1222213231010101"></a>

<a id="canonical-2211101120230021-3011132133201112-2013201233022301-1221122012112313-1201022030010110-1213000302123121-3322010300332022-2320002321031200"></a>

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

<a id="canonical-1321321031020011-2210123111113130-0130130013323122-1233233103132323-1233313132023120-1012003002111120-1011033030302312-2110121301333200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3300221003232331-3211111330231010-1133102021321221-3113233120221120-3333210120021310-2100333220120013-3013033123311221-1030103223310033"></a>

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

<a id="canonical-2220211130233202-3101113103101320-0220333210010111-1011111301312122-3220111122132123-2003001230321300-2020020023133202-3203110310202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-2033220231110220-3301013222133312-2233001222213200-1013323223223103-3220122113013101-1213111133212201-2020132211221323-1231231210223322"></a>

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

<a id="canonical-1231021311211130-2111221113033111-0332210121011101-2122230310303120-1230303021332210-0031330213001012-1300102223330331-3010011331221200"></a>

### Direct properties for `proxy_config.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-3102012303001312-1122213203200321-1121003003312322-1310032330230232-0303010320133131-0320110222223321-0131031213323210-1320301203200221"></a>

#### `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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
