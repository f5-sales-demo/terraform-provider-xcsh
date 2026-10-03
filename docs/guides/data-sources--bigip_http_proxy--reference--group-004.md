---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-2032022210330232-0123102103100013-3200021033113231-0131230212030313-2331133031000031-0202230022200002-2231013300011212-0132323110211210"></a>

## Next pages — medium_security / 131002133222 / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312300303312231-0011202321001132-0113030331133100-0322131112123100-1323031302212001-0013210010213301-1220033310201312-2323312323331113"></a>

## proxy_config.https.tls_parameters.use_mtls — use_mtls / 302222000220 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
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

<a id="canonical-3330220201120031-0030022323031120-0322101302003210-3011222313331121-0033233023021320-3213011211330332-0213131202012300-0032003210000030"></a>

## Direct properties — use_mtls / 302222000220 / 3

<a id="canonical-1222221102032102-0301010230213203-2303011321120022-1101221113201102-2132230111113303-0001210203101302-3210222100021002-3320121231111223"></a>

<a id="canonical-1032123120332322-0001020023212121-1330113301233303-0031313331132300-2302322021120003-3230122322110222-0311310331333333-0232102331002030"></a>

## client_certificate_optional property — use_mtls / 302222000220 / 4

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103023323301323-0111023033110200-0300023121212110-2202323211103312-0231200012231003-3222031211023230-1333111020132000-0323223013202100): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3213022221230221-0333221301301032-2332313122021221-3130323103301132-3232120330210301-2232300001123210-0320111231330122-2311033032123311): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1113323230002302-0300323000011031-0011032132221030-0301103302313220-1010122110211100-0111002312201230-1311100321003330-0010010103000312): complete subsection reference.

<a id="canonical-0203331203230323-2311010210221301-1222010030032211-1020233111321232-3110103023100013-0030203010101120-2132331002022002-1232310222320301"></a>

<a id="canonical-1123300003320202-0323002322212101-0011333133012201-2322032003300102-0303201232221223-1010131202333101-2122231101031012-1122231031302202"></a>

## trusted_ca_url property — use_mtls / 302222000220 / 5

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2121221323123222-0220020011110122-0130021310122201-3232122323032032-2021130121222212-3323011112111123-1012303112132213-1212221210122231): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2101132012131003-1220132123220211-3310230112312103-2131222220222232-1310313020030023-3333310110133121-3333103030303113-3121103101020312): complete subsection reference.

<a id="canonical-3233022003010132-3021300332312030-1022011010030102-1022023011232013-1301313130113113-1132012031101012-3111200331131300-3231302212302123"></a>

## Next pages — use_mtls / 302222000220 / 6

- [proxy_config.https.tls_parameters.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103023323301323-0111023033110200-0300023121212110-2202323211103312-0231200012231003-3222031211023230-1333111020132000-0323223013202100)
- [proxy_config.https.tls_parameters.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3213022221230221-0333221301301032-2332313122021221-3130323103301132-3232120330210301-2232300001123210-0320111231330122-2311033032123311)
- [proxy_config.https.tls_parameters.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1113323230002302-0300323000011031-0011032132221030-0301103302313220-1010122110211100-0111002312201230-1311100321003330-0010010103000312)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2121221323123222-0220020011110122-0130021310122201-3232122323032032-2021130121222212-3323011112111123-1012303112132213-1212221210122231)
- [proxy_config.https.tls_parameters.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2101132012131003-1220132123220211-3310230112312103-2131222220222232-1310313020030023-3333310110133121-3333103030303113-3121103101020312)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3103023323301323-0111023033110200-0300023121212110-2202323211103312-0231200012231003-3222031211023230-1333111020132000-0323223013202100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201023122323122-1100322320232323-3030023020330322-1111302212220030-0030213212002311-2033322202323003-1323320212102201-1212120222233113"></a>

## proxy_config.https.tls_parameters.use_mtls.crl — crl / 212123232202 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.crl

<a id="canonical-2230022002021313-3300233322301303-2032200013000312-2330132033211031-0033221033030130-1020332003331232-0230001121330011-0231221000230332"></a>

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

<a id="canonical-1323021030010311-3322010212120323-1131203320010131-0103332012113030-1330313023202202-2230333000331311-2213133300321313-1120011101110112"></a>

## Direct properties — crl / 212123232202 / 3

<a id="canonical-0020203023211202-3220002003002120-2310022121033330-1313122133320201-1000012210210233-0213122230113233-3213321202132320-2330213301012000"></a>

<a id="canonical-1211300103101030-2011003323031122-1322223312200222-0033322103203100-2302032012022301-2311213333330220-2203102200101301-0223203003011033"></a>

## name property — crl / 212123232202 / 4

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

<a id="canonical-2201210323033331-0013132010202030-1231130302231132-2032200110212120-0201113012203230-1001111200131130-2231033031210103-0001303021202012"></a>

<a id="canonical-0111220211302312-3332020220132321-0203131100322030-3100331210310123-2201321030033323-2132211132012212-0103232300011203-0323320103323320"></a>

## namespace property — crl / 212123232202 / 5

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

<a id="canonical-2230223213021202-3301210310212002-1221122122103210-2321131130300302-0003131223033113-3211211012313213-2333023013303312-3001020322321301"></a>

<a id="canonical-0103011003213131-1031302110112332-0000000003300320-2001232322013120-3013131310230202-3231211300310002-0233311022333300-2300303332321302"></a>

## tenant property — crl / 212123232202 / 6

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

<a id="canonical-3201213203223330-0023232020202202-1020203031233332-2301032101020321-0012210333311233-0113201110302321-1303201223033322-3320102313020232"></a>

## Next pages — crl / 212123232202 / 7

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3213022221230221-0333221301301032-2332313122021221-3130323103301132-3232120330210301-2232300001123210-0320111231330122-2311033032123311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200023022012103-3321222320212221-3203101333202331-2301030302222212-2203002103010021-1012010213221121-3330110001022201-1200021002201313"></a>

## proxy_config.https.tls_parameters.use_mtls.no_crl — no_crl / 001002310303 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1231121101201303-3003131010302113-0032003033121321-1220003331002032-0032301230233303-2012103131020200-0011333213103311-3120222223101223"></a>

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

<a id="canonical-1122323000332100-2013121111121132-3201220221322333-0110113210302233-3302120000030322-2330003321302231-0233212220200220-2001132211323021"></a>

## Direct properties — no_crl / 001002310303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112221103130112-3130231212003000-3201011012132132-3233323033012012-3211212022003031-0231203022203131-1311200100120301-1213033133102100"></a>

## Next pages — no_crl / 001002310303 / 4

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1113323230002302-0300323000011031-0011032132221030-0301103302313220-1010122110211100-0111002312201230-1311100321003330-0010010103000312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231120002021332-0100210222201211-0032310003103023-0102030020100313-2002201303333130-2101103231203020-2200023212303031-3202003302032302"></a>

## proxy_config.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 003123333312 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-3230120200323110-3201303223210202-2210012330010220-3200010210102232-2300332301222220-1202030213011321-2031023233012221-0131011113333200"></a>

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

<a id="canonical-0012200002033311-2322033130302320-0311323013223110-1120023010313012-1223020013010222-1201111001310001-0202021031132233-3330013300133100"></a>

## Direct properties — trusted_ca / 003123333312 / 3

<a id="canonical-1330203203203102-2230031203023130-1320022200121021-0102001232031222-0102223121311010-2023010200020321-3232100323333121-3222021303121122"></a>

<a id="canonical-2101130311233313-0030323232330110-2223220203213021-3313131013322111-2113132131311022-2001012101301021-2220300310203232-2022120301010221"></a>

## name property — trusted_ca / 003123333312 / 4

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

<a id="canonical-3033302021110310-2302133030033030-2302002103023332-3320232221310122-2311113302231212-1032233113110201-0100010202212223-3220003133011123"></a>

<a id="canonical-3232221020313122-2313210300010132-1203111330300310-1220300120100003-0113013223310112-3123203213312103-3031201120030032-2231232020112112"></a>

## namespace property — trusted_ca / 003123333312 / 5

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

<a id="canonical-3222122133121300-0320032003132213-3100203220012032-1110213203000120-1201201011010132-2330000133311111-0100022211120031-3012301113130021"></a>

<a id="canonical-2100333300031103-3133201000201021-0310302322330233-2201301301322002-1033323321311321-3323331000022032-0011112221101003-0001302011210302"></a>

## tenant property — trusted_ca / 003123333312 / 6

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

<a id="canonical-2011033201221031-1112331330301231-1132221321222233-2321213011313110-2031000230000231-2030101123131221-1211012221022021-2301232332221000"></a>

## Next pages — trusted_ca / 003123333312 / 7

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2121221323123222-0220020011110122-0130021310122201-3232122323032032-2021130121222212-3323011112111123-1012303112132213-1212221210122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201320301101020-1021110202321202-0321131302112302-1311110021132132-3312220321310320-0012303210320302-2322213223100233-1303303210330330"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 232012310133 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- proxy_config.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3011121222000332-0101100213001002-2002033200210003-2112312223031322-0012102320011232-1322102030303020-1121133131130132-1132301201022012"></a>

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

<a id="canonical-3303311303121330-1002231230321222-0211230202201030-3222322010132003-3023203002110030-2333012312033322-1032010000222213-1132020310111032"></a>

## Direct properties — xfcc_disabled / 232012310133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002203010233211-3303300121000303-1010332313333122-1323311130003122-1103030113021302-2003210323111120-3321330020122231-1031023213301130"></a>

## Next pages — xfcc_disabled / 232012310133 / 4

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2101132012131003-1220132123220211-3310230112312103-2131222220222232-1310313020030023-3333310110133121-3333103030303113-3121103101020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230003331212100-0001002123013300-0230302310013002-1321003211002320-3101113311101212-1101322133031231-0130000310231102-2012200302230030"></a>

## proxy_config.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 011211032323 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
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

<a id="canonical-2133100033130030-2122200322312102-3303312311033012-1013123110020300-1212100231320020-3322130021101123-3331321001122232-3123220223213121"></a>

## Direct properties — xfcc_options / 011211032323 / 3

<a id="canonical-2212131331211022-0100321231223020-1101223300331010-0032333231222220-0012002210132301-1103221030320313-2310303232300201-0013002132213030"></a>

<a id="canonical-3102330300303202-0212332310213003-3100001030032213-1201132332201133-0000131203203021-1320313211230020-2330212300301003-2031003010013212"></a>

## xfcc_header_elements property — xfcc_options / 011211032323 / 4

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

<a id="canonical-0301100312131023-3130112223313330-1013121032132000-1322230333302132-3001232310031031-2031012332102032-3313331102010031-3232222120020211"></a>

## Next pages — xfcc_options / 011211032323 / 5

- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322233321013113-2223222302202022-0310320022131120-3310201303311010-3310011301011133-2033102011320200-1020202331323111-2331120013101310"></a>

## proxy_config.https_auto_cert — https_auto_cert / 201213133230 / 2

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

<a id="canonical-1011231001032102-0231001311300013-1333002033301330-1222131212203210-3122132002032323-1112223322232322-0000112111023112-3330331011021323"></a>

## Direct properties — https_auto_cert / 201213133230 / 3

<a id="canonical-0101323301013302-1010300003023001-0220021310320031-2023030332213003-0333033221011211-3110323012321311-1222130013010031-2031020212122100"></a>

<a id="canonical-3112102222203003-3021200302232302-3131203300313001-3323003011130200-3000002020121310-3113023121222211-0013123203002201-3132330323132310"></a>

## add_hsts property — https_auto_cert / 201213133230 / 4

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

<a id="canonical-3311112130020031-0312233301332213-0001030002203220-1320200012100103-3033000031130022-1313122213311203-3330221203133310-3133330011223303"></a>

## append_server_name property — https_auto_cert / 201213133230 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030): complete subsection reference.

<a id="canonical-1003203231021222-0333313032131323-3303232231021212-2232302002012130-1203103122131012-0333220033132332-0302233313103003-1011001032311300"></a>

<a id="canonical-1121010311332113-1102003100321030-2201112001112003-3210202120133223-0123230333031030-1122322023030110-1101031213312033-1122301201103020"></a>

## connection_idle_timeout property — https_auto_cert / 201213133230 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

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

<a id="canonical-1112011333311112-0321333202301133-1022113111322232-3322130003322121-0201012322122330-2230323001123112-1223120020200213-0113203220320320"></a>

## http_redirect property — https_auto_cert / 201213133230 / 7

Type: `"bool"`. Computed.

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

- [no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3032101132210133-0133123021311130-0333100232101300-3110202321302120-0113332010033110-0333001021210231-3102100200000002-3211302011322302): complete subsection reference.

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0302113122013130-2201130101123011-0221313223300113-2322200323232133-0332312100030303-0313111201230213-0020330003231321-3233013100222222): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1222113032213003-2332033231032202-3331011322132003-2303223012013220-1002313103320300-1032332322010113-2203130032100331-0312223303020013): complete subsection reference.

<a id="canonical-2331203313130010-2313112233233323-2133212012233133-1131113202323321-1302302231312313-1032030010312012-2322210330311321-1320033130131200"></a>

<a id="canonical-3012011103221030-0031310323132010-3032222223310113-0202323220331100-3312133132331232-1310120201103323-3133130203131212-2021212230023122"></a>

## port property — https_auto_cert / 201213133230 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3120120113322132-1130001130221021-0310131222312303-2121123010111110-1010111113002322-1303211320300112-3022111021120002-0233220122010111"></a>

## port_ranges property — https_auto_cert / 201213133230 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-3011130222220013-0333231121310311-1212123333111132-1013322120312112-2312131310020132-0020211111112102-3200102113321133-1131112311133320"></a>

## server_name property — https_auto_cert / 201213133230 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200): complete subsection reference.

<a id="canonical-3111003321323233-2100001122230111-2201123323030213-3123233123210102-3030131031020220-1100313221120102-2123010031322033-1320120002333001"></a>

## Next pages — https_auto_cert / 201213133230 / 11

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030)
- [proxy_config.https_auto_cert.default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0030300011320301-0132002031320201-0031231310003021-0011110133003223-2321220231230203-1323311203311323-0210211322313002-3110003101030223)
- [proxy_config.https_auto_cert.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2100311102133131-0103333221332103-3011131011012030-3122233311133132-1021210011010121-1302000133330301-0312010300011033-1301010332213133)
- [proxy_config.https_auto_cert.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1013121000213233-3301313031232112-2232012301322202-3300110222313112-0332302333031133-3213331200001300-1113123201201112-0210100233000201)
- [proxy_config.https_auto_cert.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3113213021003302-1131212001221030-0002232330203000-3121013311103022-1330233330212100-0003032023030012-3102022301131101-3111302200302101)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [proxy_config.https_auto_cert.no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3032101132210133-0133123021311130-0333100232101300-3110202321302120-0113332010033110-0333001021210231-3102100200000002-3211302011322302)
- [proxy_config.https_auto_cert.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0302113122013130-2201130101123011-0221313223300113-2322200323232133-0332312100030303-0313111201230213-0020330003231321-3233013100222222)
- [proxy_config.https_auto_cert.pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1222113032213003-2332033231032202-3331011322132003-2303223012013220-1002313103320300-1032332322010113-2203130032100331-0312223303020013)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311130330301203-2010230223313132-3313213333233032-1320103022221221-2303313302120012-0202102032123103-1133022223031121-0113313132123332"></a>

## proxy_config.https_auto_cert.coalescing_options — coalescing_options / 021120001202 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.coalescing_options

<a id="canonical-2323231130323012-3210230322131303-2131020020023330-2010232203322211-1223302330003001-1313310012012201-3311112113102330-1312332233123222"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-3231232332023330-3103031320301023-1302132112202110-2220212003122220-2000033001322223-3002211033312212-1230330323333130-2013113100311003"></a>

## Direct properties — coalescing_options / 021120001202 / 3

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1201231022323222-0312232112121122-1313113331001030-0113111213331331-1003121332211030-1003311101231233-0101020121023013-2102030233302313): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2231133130101031-2100231030221323-0030010033111001-2103122302302312-2122022323011230-2320223210213300-3223001102311101-2133113311001003): complete subsection reference.

<a id="canonical-3013210311310300-2103322321201123-1130122123112131-0213113002122111-2312213203110223-2003000310030311-2321221203111101-1031130330221110"></a>

## Next pages — coalescing_options / 021120001202 / 4

- [proxy_config.https_auto_cert.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1201231022323222-0312232112121122-1313113331001030-0113111213331331-1003121332211030-1003311101231233-0101020121023013-2102030233302313)
- [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2231133130101031-2100231030221323-0030010033111001-2103122302302312-2122022323011230-2320223210213300-3223001102311101-2133113311001003)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1201231022323222-0312232112121122-1313113331001030-0113111213331331-1003121332211030-1003311101231233-0101020121023013-2102030233302313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013031201301233-1030020022000012-2021102031022322-2132002103033333-1011021100111232-2213001203033120-2210120002031203-3323210111002103"></a>

## proxy_config.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 033010100310 / 2

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

<a id="canonical-1301313323132300-0113103003120301-1113103012103113-0303213200033212-1132223111200302-0210121201133332-1203102001200110-2222111133323301"></a>

## Direct properties — default_coalescing / 033010100310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021033020231100-1302203121311303-2222121301300003-1321220330320201-3100232223232100-1023013212223111-2022233003302011-3310021200023202"></a>

## Next pages — default_coalescing / 033010100310 / 4

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2231133130101031-2100231030221323-0030010033111001-2103122302302312-2122022323011230-2320223210213300-3223001102311101-2133113311001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210220101231103-3030301220023002-1221010030130222-0021121303230110-3222112113132121-1110213330022321-1010130313110111-0110002110220110"></a>

## proxy_config.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 023211122033 / 2

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

<a id="canonical-2003021300013133-1220323021303021-0222233221021111-3213033011313101-1322302030331223-1010223102033032-0330130211000113-2012033021203101"></a>

## Direct properties — strict_coalescing / 023211122033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111112021120200-0131330220201300-3231302013100103-2103121300010220-2002111132121310-3223331011133300-0232133020133230-1310010113021033"></a>

## Next pages — strict_coalescing / 023211122033 / 4

- [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2310302203213123-2333210123111120-2201211302012310-1133232330203102-0222022020232011-3022100330132330-3222121330202112-2031002203001030)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0030300011320301-0132002031320201-0031231310003021-0011110133003223-2321220231230203-1323311203311323-0210211322313002-3110003101030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123113232121220-1220131111011000-0322210000301230-0122023002113322-3310212031233131-2121122231210130-1011022220121313-2201000012202302"></a>

## proxy_config.https_auto_cert.default_header — default_header / 021023030313 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.default_header

<a id="canonical-1220221311011023-3020210333333021-2113131331233221-0203100122313022-2112230013331323-1002110332332330-3313130123112130-0030202222200202"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1101101233232323-1101011013021330-3101230213020302-2023232321030223-2131301222110033-1210102021320231-3322030223103133-1011133131133232"></a>

## Direct properties — default_header / 021023030313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110023323030213-2312210130100222-1333332311010013-1232122010012201-3310020130010020-2330103311233033-0111221333020233-1023122033322333"></a>

## Next pages — default_header / 021023030313 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2100311102133131-0103333221332103-3011131011012030-3122233311133132-1021210011010121-1302000133330301-0312010300011033-1301010332213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030301003130203-2021003100101011-0103303110201113-1122000313211120-2121221131012211-2303101103001231-2310313103112313-3132133122222210"></a>

## proxy_config.https_auto_cert.default_loadbalancer — default_loadbalancer / 322100203011 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.default_loadbalancer

<a id="canonical-2322230011133301-3212102133120320-0022102002213201-3200230223231301-2310001331003030-3110323131321221-1103103311313111-3030000310022331"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1221122223222202-0020300210320232-3231203022121201-1032212212202010-3332312113303003-1012010232002201-0001021212211213-3101321033023201"></a>

## Direct properties — default_loadbalancer / 322100203011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131332312000110-2321202122310230-3223033011122033-3021320032301220-0203022233300010-3223300322133013-2213203332001330-2102002021023130"></a>

## Next pages — default_loadbalancer / 322100203011 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1013121000213233-3301313031232112-2232012301322202-3300110222313112-0332302333031133-3213331200001300-1113123201201112-0210100233000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323102300202001-0313310011323133-3031302121000303-0112120101223223-3030110012101010-1010223220311032-3102300031003112-2311013311111030"></a>

## proxy_config.https_auto_cert.disable_path_normalize — disable_path_normalize / 133333000103 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.disable_path_normalize

<a id="canonical-3331312223213323-3320203303333303-1222130303122200-0331230031023230-0113231210131310-0120331202022003-0221233022310020-2210101021101200"></a>

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

<a id="canonical-2231220321032333-3120011302323133-0012032303122210-1101233313322111-3231333201223100-3232230221111120-3102122023101323-1222230133231001"></a>

## Direct properties — disable_path_normalize / 133333000103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013020000030333-1011013122213030-1203101311102300-3203232321023133-3330303302320002-0202033013133021-0302233102332121-2110032133332231"></a>

## Next pages — disable_path_normalize / 133333000103 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3113213021003302-1131212001221030-0002232330203000-3121013311103022-1330233330212100-0003032023030012-3102022301131101-3111302200302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010301333322302-0221000313023030-1332330301302122-1202333211032111-0102112112131122-1023212012111113-1333300012011030-0233222011132123"></a>

## proxy_config.https_auto_cert.enable_path_normalize — enable_path_normalize / 003132131012 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.enable_path_normalize

<a id="canonical-0313012103101132-2312223303001101-1000003001312033-3320213321030033-2231131202213121-2122120221210311-1332102323100322-3332033132301313"></a>

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

<a id="canonical-0201312111332212-3122130000223131-3013112202113010-0210122003331011-0321011001103033-2322223032030010-0012131202312022-1320102213021232"></a>

## Direct properties — enable_path_normalize / 003132131012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012021002210001-2102110013212202-0002102300302002-1031203220202232-2131202101231032-0101311302121122-0321112230230131-3300020000213111"></a>

## Next pages — enable_path_normalize / 003132131012 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123012132010302-1231130103031320-2112121200333320-2112310232310212-1013111110120302-1130201122301231-3302012202213320-0332221131033320"></a>

## proxy_config.https_auto_cert.http_protocol_options — http_protocol_options / 201133233201 / 2

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

<a id="canonical-2122022230231222-2110121202031212-3213001010131210-1303001101102100-0213311332311033-2102031120101122-0123302032201200-2010203111300332"></a>

## Direct properties — http_protocol_options / 201133233201 / 3

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3000323121202033-3300212300023110-0111212212200013-3313101003222032-0103303331131032-3122111233120203-1211302210132103-1131222202033033): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2110332203130223-1121033002100221-2100112000011200-3300022201202021-0320222331231110-0131110232302203-2223203300010021-2211113311233003): complete subsection reference.

<a id="canonical-0020333120012122-3220223303220130-0013210011203331-0303301300230322-2201000102300301-0201221200111121-0102022203133100-2332212012230031"></a>

## Next pages — http_protocol_options / 201133233201 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3000323121202033-3300212300023110-0111212212200013-3313101003222032-0103303331131032-3122111233120203-1211302210132103-1131222202033033)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2110332203130223-1121033002100221-2100112000011200-3300022201202021-0320222331231110-0131110232302203-2223203300010021-2211113311233003)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023333302031200-3331103111222100-3300102331220030-2230113122222322-3030230003103300-2003020312110202-2131210001203312-1230132031130123"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 002030123211 / 2

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

<a id="canonical-1003211220331030-3123113000333300-3033221211012001-3011333111003310-2101320013312101-1322033211130103-0212221310132210-0012322210123132"></a>

## Direct properties — http_protocol_enable_v1_only / 002030123211 / 3

- [header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200): complete subsection reference.

<a id="canonical-2012010013310103-3320012101133231-0021233203212000-2113113233231012-2020021332003302-2102313132123233-3031311223303233-2113230212120323"></a>

## Next pages — http_protocol_enable_v1_only / 002030123211 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301122012130022-3231121021302212-0100221001202000-2030131000323201-0130023311221100-1301102312333313-1103220332302320-1133102322231013"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 322210233211 / 2

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

<a id="canonical-0032220203200321-0231220333120332-1232211333331032-3320010312121012-0212133212312010-3001320211021220-2130121301323132-2023033110102032"></a>

## Direct properties — header_transformation / 322210233211 / 3

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2023020201220130-2001110312010022-0301020130011113-1220312330020230-0300320230002333-3030223210302222-1210022120330030-0101002200031313): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1223312000323310-2123213220002032-1132333311221100-2203010220321213-3122013120301331-2012011131213223-2230130131233220-2113302220101211): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2312321333331322-2330031223122300-0212202121223330-0300212321112131-3301013033030330-0120331011030202-3311122322321220-3121132032111133): complete subsection reference.

<a id="canonical-1222223213210120-1031223101200333-3211223000112200-0330310121112001-2213011213113230-0120301322210320-0333012303330122-3103013013101030"></a>

## Next pages — header_transformation / 322210233211 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2023020201220130-2001110312010022-0301020130011113-1220312330020230-0300320230002333-3030223210302222-1210022120330030-0101002200031313)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1223312000323310-2123213220002032-1132333311221100-2203010220321213-3122013120301331-2012011131213223-2230130131233220-2113302220101211)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2312321333331322-2330031223122300-0212202121223330-0300212321112131-3301013033030330-0120331011030202-3311122322321220-3121132032111133)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1023111120022023-0203113003302200-3101021010111100-0003211131131223-0023102020323330-2221322133103132-0201121231103000-1100231030301232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2023020201220130-2001110312010022-0301020130011113-1220312330020230-0300320230002333-3030223210302222-1210022120330030-0101002200031313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110333333112232-3302110113201103-2311330202100122-1311333212221231-3033333011322223-0123113031223320-2023113023303013-1323120033323202"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 320131102313 / 2

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

<a id="canonical-0002212012330220-3302112321311111-1102220303232113-0123321300133010-2333031302132011-1202002011133031-0233300000311321-1231330102020222"></a>

## Direct properties — default_header_transformation / 320131102313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212031303131033-1112222010022030-0331120031200022-3212020221323000-1122032310200200-3111222021032333-2000300101000032-1211130233132133"></a>

## Next pages — default_header_transformation / 320131102313 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1223312000323310-2123213220002032-1132333311221100-2203010220321213-3122013120301331-2012011131213223-2230130131233220-2113302220101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330223233231303-1132310011122020-2111010313030311-1230112223201232-2021103122020002-3223101331012131-0300201001233110-0031330131323010"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 113301201332 / 2

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

<a id="canonical-1332302022102003-1001112303101103-3203131212022322-3102003031101130-2010012032310122-2021203030130302-0133331020211303-2011113212320120"></a>

## Direct properties — preserve_case_header_transformation / 113301201332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102310203201111-3010303330220223-0232221303001102-0022203321003220-1112121003220120-3102132233330033-2130103133020202-1120221323310012"></a>

## Next pages — preserve_case_header_transformation / 113301201332 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2312321333331322-2330031223122300-0212202121223330-0300212321112131-3301013033030330-0120331011030202-3311122322321220-3121132032111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210301012023130-2000333031010300-0031000211213030-3102033000320011-2120033213120322-1100130000213010-2333301021210121-2231201310302332"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 311103333010 / 2

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

<a id="canonical-3101222330131012-2310011023302223-1220011101211220-1202300312220010-1030221321101110-2000131113213121-0022032012200302-0013121213331013"></a>

## Direct properties — proper_case_header_transformation / 311103333010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330013213303013-3023221122133111-3332221033331321-1320112133120031-1120113220031001-1223200301302322-2331220122323002-0020313311020032"></a>

## Next pages — proper_case_header_transformation / 311103333010 / 4

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0310100012331121-1230012031133032-1112200203310213-1312001123120212-3333131223211221-0233311012332203-2302110231032111-2223030301203200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3000323121202033-3300212300023110-0111212212200013-3313101003222032-0103303331131032-3122111233120203-1211302210132103-1131222202033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333230231330202-1321232223331033-2013021120002110-0000013100032322-0101000121212203-0110332222311222-2332203120222203-2320300012333020"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 020310323021 / 2

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

<a id="canonical-2230111121221222-3011001331213231-3131233222222030-1232032130123011-2200032223131332-3221313231202013-0120313200320210-3103123122130203"></a>

## Direct properties — http_protocol_enable_v1_v2 / 020310323021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233030130132021-2001211112332000-0022301301002212-1200020201132111-3212230221032212-2321021203100301-2001232021012022-0333220033200311"></a>

## Next pages — http_protocol_enable_v1_v2 / 020310323021 / 4

- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2110332203130223-1121033002100221-2100112000011200-3300022201202021-0320222331231110-0131110232302203-2223203300010021-2211113311233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223323101033022-0231110000033021-1322322332132221-3011132220023231-3032131121312210-0030311122131303-2030020123330031-1013310222100133"></a>

## proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 121113112133 / 2

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

<a id="canonical-3102132000231121-0020232321010021-0120031032031230-2101032123001033-3100110300322323-3211321203332230-2303200023333102-3110213133201202"></a>

## Direct properties — http_protocol_enable_v2_only / 121113112133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220001133033011-3210030122330103-1323033002331320-2100301130210000-3021222100002001-2232012131310023-0213333020003102-0123323001010200"></a>

## Next pages — http_protocol_enable_v2_only / 121113112133 / 4

- [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3201322300111011-3201210323022220-0322311001030112-1310200030102302-0032023012031312-0301111210033200-1022300022030120-3233030300121111)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3032101132210133-0133123021311130-0333100232101300-3110202321302120-0113332010033110-0333001021210231-3102100200000002-3211302011322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003123202232211-3201003013203302-2111012030222221-0200001332130101-0010313221032030-3032210232122102-2021003121221231-0101200202203001"></a>

## proxy_config.https_auto_cert.no_mtls — no_mtls / 301220020200 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.no_mtls

<a id="canonical-2330213311302231-3100230321010321-3322302022301123-2320221211201113-2220033013012303-0213130013021130-1302021320113122-0133232310311020"></a>

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

<a id="canonical-1303203210010302-1131122001002001-2223030301331021-3320110100122030-3223102213221002-0002222112221223-1020011201222101-1003111311133131"></a>

## Direct properties — no_mtls / 301220020200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210303223323120-3301032022211202-3010332102311100-1222301303030023-1331131230100311-2203032003132201-0310222000310332-1131231031221023"></a>

## Next pages — no_mtls / 301220020200 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0302113122013130-2201130101123011-0221313223300113-2322200323232133-0332312100030303-0313111201230213-0020330003231321-3233013100222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323221000003330-2131202203210120-3123330330320301-2012233222020110-2232113231111103-3303222012321331-1010000002012303-1311231031133122"></a>

## proxy_config.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 233032130031 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.non_default_loadbalancer

<a id="canonical-2312110211233330-1111110100332323-1330333030312102-0032001312103331-1110002032230222-0202202133203001-0331231133311221-3303102222203313"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1333233131221112-3130330321331010-0323001230021123-3003300302102203-3311100221010320-2221020322221112-0301103023323013-3330110300233310"></a>

## Direct properties — non_default_loadbalancer / 233032130031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323103021013300-0201320220331002-1311132300232221-0210121011220103-2230011231132102-1012201011231130-1201310001203030-3220233323133012"></a>

## Next pages — non_default_loadbalancer / 233032130031 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1222113032213003-2332033231032202-3331011322132003-2303223012013220-1002313103320300-1032332322010113-2203130032100331-0312223303020013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002130130322303-3213311001003001-2303323310122131-1122322221021023-0312133003232121-1013223020210003-0031023101121132-2122020330133330"></a>

## proxy_config.https_auto_cert.pass_through — pass_through / 100102203030 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.pass_through

<a id="canonical-3012320311120332-1323300231031113-1122203313103032-2223223202311112-0022323113312133-1033221032323132-0300331112021322-3302211223302133"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3330323101112323-0321013030210221-2110132320123231-2123300230222102-3330120211301303-1100310132200233-2122000003301021-0313201130001100"></a>

## Direct properties — pass_through / 100102203030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003131222130321-2110013021312023-0330033203203133-3230200033211020-0101222300230123-0100021020322130-1313213221311130-0120333120010012"></a>

## Next pages — pass_through / 100102203030 / 4

- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010121233210202-2333131232132110-2202231110001330-3231013123123031-0333032032312111-3130022322113310-0131301301213233-3330101130133202"></a>

## proxy_config.https_auto_cert.tls_config — tls_config / 121321110311 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- proxy_config.https_auto_cert.tls_config

<a id="canonical-2203113220010212-3033312212222231-3030031010323303-0300130201113211-1213023020013202-1021320131313003-1123021110133103-0222202300112121"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

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

<a id="canonical-1322310031221003-2003200003312210-3300312132322103-1131201233223323-0302022012331232-0102330030213000-2221021213200332-3020130032101002"></a>

## Direct properties — tls_config / 121321110311 / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0300221023003133-1200321233211220-3323030113020330-1212112112121032-0013221230033133-3021120002020210-3311102202011113-0331111212311103): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1012133112021320-3130100100122231-0211230002020233-0322232323210330-0012101220010231-1110232031030032-0312102221101331-3220223232220000): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3102332131122100-2011331213122200-1312320233130232-1212312131030131-3301302201303212-3312302200112230-3322310221123210-1332112122113113): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3011030212111000-0031311313020310-2303131101210302-0133230212310003-1130220313222013-1112133232232012-2321200001333021-2000302301312111): complete subsection reference.

<a id="canonical-2031210013023220-0203212201332013-3121230122023010-1102230332312011-2333010033133013-3111111301322131-2102000321322111-3020131112130132"></a>

## Next pages — tls_config / 121321110311 / 4

- [proxy_config.https_auto_cert.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0300221023003133-1200321233211220-3323030113020330-1212112112121032-0013221230033133-3021120002020210-3311102202011113-0331111212311103)
- [proxy_config.https_auto_cert.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1012133112021320-3130100100122231-0211230002020233-0322232323210330-0012101220010231-1110232031030032-0312102221101331-3220223232220000)
- [proxy_config.https_auto_cert.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3102332131122100-2011331213122200-1312320233130232-1212312131030131-3301302201303212-3312302200112230-3322310221123210-1332112122113113)
- [proxy_config.https_auto_cert.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3011030212111000-0031311313020310-2303131101210302-0133230212310003-1130220313222013-1112133232232012-2321200001333021-2000302301312111)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0300221023003133-1200321233211220-3323030113020330-1212112112121032-0013221230033133-3021120002020210-3311102202011113-0331111212311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332000330132210-2323023001201002-0022230110213303-1301330110130221-2201121102112211-1001333120212231-0323303101313311-1000312130123123"></a>

## proxy_config.https_auto_cert.tls_config.custom_security — custom_security / 022003202100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- proxy_config.https_auto_cert.tls_config.custom_security

<a id="canonical-3001333321120210-2332221111031121-1023011233212011-2011302311220001-0111020131102301-3020021302323110-2020200333322330-3001322213013310"></a>

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

<a id="canonical-3120312003101322-3220222231220320-3301132321021133-0313333021031300-0130101213102322-2210112133033023-0130321010010322-2230213103211111"></a>

## Direct properties — custom_security / 022003202100 / 3

<a id="canonical-1222331312321333-3320131122110231-2133110303103211-3010332021102321-3313201032220230-1201120021113232-0010131032221120-0132212220133001"></a>

<a id="canonical-3033213331030333-1113133013111130-1122003131303003-1123313321120111-3202013102011312-1100030213320211-0110000321032000-0232000313012223"></a>

## cipher_suites property — custom_security / 022003202100 / 4

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

<a id="canonical-3332201301220022-1021020210122221-1132223303332131-0332120102311302-3333023203122002-3230010330033320-2001101010100012-3002021011123020"></a>

<a id="canonical-3122133233213212-1213302333200313-3333101123322032-1022031302100102-2102223113210010-2002202020333032-3302330302213211-1223331112320011"></a>

## max_version property — custom_security / 022003202100 / 5

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

<a id="canonical-1221333111022133-1022311032122131-1030232203331120-1220022011021102-2022302233020333-1130102321203211-2312030322103123-3210202301232310"></a>

<a id="canonical-3111331003331221-2202203203121110-0033021121123331-3111110022020320-2302022103112000-0301113311211200-1313313212310330-3003220221032232"></a>

## min_version property — custom_security / 022003202100 / 6

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

<a id="canonical-3013130321113233-1133303022233332-1321231322231332-3103101223010022-1233301023233211-3222222102330031-0120130021012332-3221123220320212"></a>

## Next pages — custom_security / 022003202100 / 7

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1012133112021320-3130100100122231-0211230002020233-0322232323210330-0012101220010231-1110232031030032-0312102221101331-3220223232220000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222313033011312-1112121212020202-3103211103110111-3002231012112131-3231132023223321-2110302032211210-3032332031312322-2221233113211131"></a>

## proxy_config.https_auto_cert.tls_config.default_security — default_security / 223201022011 / 2

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

<a id="canonical-0312303232213111-0023303210100033-2130032312113023-3103320011022013-1033100223020121-0232203221200011-2202232221102000-0313103310013300"></a>

## Direct properties — default_security / 223201022011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332201131202321-1312310032231222-3012031311000023-3111333003112310-2113112131212002-2021211121100200-0032210303313333-2231302201120221"></a>

## Next pages — default_security / 223201022011 / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3102332131122100-2011331213122200-1312320233130232-1212312131030131-3301302201303212-3312302200112230-3322310221123210-1332112122113113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310300323320121-1300302311031011-1010311323213022-1012211223203301-3113000222212322-0123221233121221-1220003332232333-2222230123200021"></a>

## proxy_config.https_auto_cert.tls_config.low_security — low_security / 220102213223 / 2

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

<a id="canonical-2100132232302220-3313211221131221-3132323133200230-1303122000320133-0210100012000301-2301220000200200-3001313312221233-0131101323001010"></a>

## Direct properties — low_security / 220102213223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120020323303001-2311322011133322-3221223332233220-2231012012011210-3022032121101330-3121122310212312-2000002312131120-2022020112223123"></a>

## Next pages — low_security / 220102213223 / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3011030212111000-0031311313020310-2303131101210302-0133230212310003-1130220313222013-1112133232232012-2321200001333021-2000302301312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133131300312010-0323110130032013-2111033132221300-0231103311023132-0001002313120302-1022233203101302-1231321002202122-2221203123300330"></a>

## proxy_config.https_auto_cert.tls_config.medium_security — medium_security / 330230311010 / 2

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

<a id="canonical-1031311022033102-0221003223203201-0322130110013030-1132031323221032-3320121310000112-2211123003111130-3002112210302101-0100201220102012"></a>

## Direct properties — medium_security / 330230311010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201331101103020-1120211232203133-1121122302233133-1303131030331233-0210320213330132-3321232210332300-2330222121001003-0331110012211102"></a>

## Next pages — medium_security / 330230311010 / 4

- [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0022100301321033-2233210033110021-1231030111232003-2102032102111310-3233030331132332-0103231000132002-2131130201000022-2323221130011232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000011113202112-1321110121313212-1033133231222113-1122130100213232-3232030113311212-2132303230302011-3223331003031300-0011121222200332"></a>

## proxy_config.https_auto_cert.use_mtls — use_mtls / 003023331113 / 2

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

<a id="canonical-2021110000122212-0222033012130110-3113022200110122-0100331321222222-2203211213002202-0322312121000112-1231121122330201-2233133030310022"></a>

## Direct properties — use_mtls / 003023331113 / 3

<a id="canonical-2303031233311231-2200311010130121-1131131000032033-2110020320303110-0132101101323320-0120003330303132-1320101120130002-0113102232120320"></a>

<a id="canonical-3301000210010210-1011122112012123-2000111123121023-3112223332323031-2000221100031301-0202123130021120-2102021302133223-2023201203111321"></a>

## client_certificate_optional property — use_mtls / 003023331113 / 4

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

- [crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3332122033322000-2313321210232312-3200030030003012-0222213102223031-2211030021311203-1223010001111230-2020203011332033-2201033121330200): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1122203310210213-0031221303120111-1302322221103122-3003302122232300-2230003233010232-0203110210310131-0022310112333330-2133101311301022): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1302111123333232-3011131213310300-0331132203231312-1203223302320231-1022311022023022-2031003213000112-0131321033010003-0130322013322132): complete subsection reference.

<a id="canonical-1030320220010232-3333301023100212-0122212210330310-3131301312300210-2312122020022010-3110033122333220-0020332233021123-2113001100031111"></a>

<a id="canonical-3102100110132323-0011130201323113-3221323130113011-2211002313133020-1122103022112212-3233333301120232-1333113032111232-1101100332100112"></a>

## trusted_ca_url property — use_mtls / 003023331113 / 5

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

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1321321031020011-2210123111113130-0130130013323122-1233233103132323-1233313132023120-1012003002111120-1011033030302312-2110121301333200): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2220211130233202-3101113103101320-0220333210010111-1011111301312122-3220111122132123-2003001230321300-2020020023133202-3203110310202131): complete subsection reference.

<a id="canonical-3330121020312020-2110332010010302-2200110332032233-3201112123033332-3232302223201103-1300003130211132-0223003033012213-3203001332010232"></a>

## Next pages — use_mtls / 003023331113 / 6

- [proxy_config.https_auto_cert.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3332122033322000-2313321210232312-3200030030003012-0222213102223031-2211030021311203-1223010001111230-2020203011332033-2201033121330200)
- [proxy_config.https_auto_cert.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1122203310210213-0031221303120111-1302322221103122-3003302122232300-2230003233010232-0203110210310131-0022310112333330-2133101311301022)
- [proxy_config.https_auto_cert.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1302111123333232-3011131213310300-0331132203231312-1203223302320231-1022311022023022-2031003213000112-0131321033010003-0130322013322132)
- [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1321321031020011-2210123111113130-0130130013323122-1233233103132323-1233313132023120-1012003002111120-1011033030302312-2110121301333200)
- [proxy_config.https_auto_cert.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2220211130233202-3101113103101320-0220333210010111-1011111301312122-3220111122132123-2003001230321300-2020020023133202-3203110310202131)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3332122033322000-2313321210232312-3200030030003012-0222213102223031-2211030021311203-1223010001111230-2020203011332033-2201033121330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231023311331032-2102200303100022-0302210023100301-3001223303310002-3123211123302323-2113031030010320-0103301121111221-2333202012111332"></a>

## proxy_config.https_auto_cert.use_mtls.crl — crl / 313202101103 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.crl

<a id="canonical-2230220111033113-2111323130131011-2211003213102120-0333230030303001-1112330011122010-2300001321200212-0112202303123022-3100003220231113"></a>

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

<a id="canonical-1000020302011233-1012100031001131-0002013033230233-3131002013313111-0111203013332230-0022301103301120-0301110122232330-1032332112320131"></a>

## Direct properties — crl / 313202101103 / 3

<a id="canonical-0300300001310211-0223302321020330-3130203110122012-1012220020332232-3012031113031300-2130323113010233-3211031320233020-1112313002220021"></a>

<a id="canonical-3011202201210120-1232301320112112-3011301220323101-1001210332000203-0322232331032001-0303210000003223-1223131032101200-3213123300321001"></a>

## name property — crl / 313202101103 / 4

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

<a id="canonical-0000222033002333-1200323302013222-0220313032232320-1203030113113111-3131002130030033-0323200202200301-0211011312112112-1232222100122212"></a>

<a id="canonical-0203101222013202-3313213222011230-2310111031311333-1123302101021223-1111323301302002-3322301130203123-3132030131120211-2001020121012302"></a>

## namespace property — crl / 313202101103 / 5

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

<a id="canonical-1210130312000102-0201213020301100-3321211303002202-1030031220230010-1023130310133302-0212123221110133-3303200110310230-2101032013222232"></a>

<a id="canonical-1111132322011221-1110032221022001-3202202020033013-0121103311213320-3032103022333310-2022131102220103-0201122322313123-1232331231103112"></a>

## tenant property — crl / 313202101103 / 6

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

<a id="canonical-2001000023123132-1201211012011203-3111311300030030-1313003311303103-2100313121232200-2001201111121132-1313230112330323-1132211122303003"></a>

## Next pages — crl / 313202101103 / 7

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1122203310210213-0031221303120111-1302322221103122-3003302122232300-2230003233010232-0203110210310131-0022310112333330-2133101311301022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323033030301332-0030023200201221-3233301122222101-1011011332233033-0313313132000322-2001302302121330-2200123031013101-3210201212022213"></a>

## proxy_config.https_auto_cert.use_mtls.no_crl — no_crl / 012121321020 / 2

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

<a id="canonical-2103012220023122-0333022011320212-1202021102121223-3022323313022130-2213031020320210-2101322222333011-0100231001012330-2222301033222230"></a>

## Direct properties — no_crl / 012121321020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121022011002100-3033121321331203-0010220013223020-3211132220033200-1131201123031313-3122122013003233-3221331331311201-3133121201313332"></a>

## Next pages — no_crl / 012121321020 / 4

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1302111123333232-3011131213310300-0331132203231312-1203223302320231-1022311022023022-2031003213000112-0131321033010003-0130322013322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102301030110333-3323320001332222-2121302120030102-1232333323130021-2013002001313322-0121321130012230-0301121030022330-1112230123301133"></a>

## proxy_config.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 120002203100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- proxy_config.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-2000202221311312-0023123202220111-2330023030332203-0211131310322113-1202213100003130-2003103030033120-0112333330023023-0002231230020130"></a>

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

<a id="canonical-2322000100321232-2122100123003111-1133031031112023-1120202010020222-0021303201220302-2211303110030321-3113110213320311-1303011010021011"></a>

## Direct properties — trusted_ca / 120002203100 / 3

<a id="canonical-2100012013213033-1123120032300232-3000123112232012-1311122121223213-1120010012000011-3210113103112332-0021002101312021-0123203231220201"></a>

<a id="canonical-2211101120230021-3011132133201112-2013201233022301-1221122012112313-1201022030010110-1213000302123121-3322010300332022-2320002321031200"></a>

## name property — trusted_ca / 120002203100 / 4

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

<a id="canonical-3323201320002212-2101211001123202-2222112111103102-3211100102100001-3201133102231032-3122222101121022-3313011032102321-3111223023321112"></a>

<a id="canonical-2321303302312322-1303321201202221-0011212103330031-2102013201301022-0100023113320233-2022133212222131-1131013132101312-2020032333320110"></a>

## namespace property — trusted_ca / 120002203100 / 5

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

<a id="canonical-0000300110313023-3110122232233002-3031222220122221-0212110330101222-3203032223223230-2001320000312010-0332200222322123-1222213231010101"></a>

<a id="canonical-1213112212111111-2012222112200120-2120113311020203-2211213332121321-3010203202120010-3232012131010322-3130031331232110-2121231301321302"></a>

## tenant property — trusted_ca / 120002203100 / 6

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

<a id="canonical-0131121033232011-3310010021322111-1303333230120332-3022122100210013-3010130001232233-3321023010223221-3310032001302311-1002312033001231"></a>

## Next pages — trusted_ca / 120002203100 / 7

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1321321031020011-2210123111113130-0130130013323122-1233233103132323-1233313132023120-1012003002111120-1011033030302312-2110121301333200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203111111312330-0010113312110201-1330013002300200-1021102022020312-2132212133230000-0232122103212031-2200232020133110-0212222320321132"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 121321101300 / 2

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

<a id="canonical-3030331311110102-1301221032323332-2202130203332332-3030232200310232-0110302231120132-0133323010102311-0211332100103331-2211100030103300"></a>

## Direct properties — xfcc_disabled / 121321101300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312111101233110-0000012122213210-2230012232222220-2200323231320311-1301320320233210-3130230223110311-1012210022131301-0021021102313230"></a>

## Next pages — xfcc_disabled / 121321101300 / 4

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2220211130233202-3101113103101320-0220333210010111-1011111301312122-3220111122132123-2003001230321300-2020020023133202-3203110310202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231021311211130-2111221113033111-0332210121011101-2122230310303120-1230303021332210-0031330213001012-1300102223330331-3010011331221200"></a>

## proxy_config.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 031131330103 / 2

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

<a id="canonical-3013210031211212-1223302320031030-3130300213023022-3101302123213200-3212121001313010-1123132021331331-2203333310133323-3210222210010211"></a>

## Direct properties — xfcc_options / 031131330103 / 3

<a id="canonical-3102012303001312-1122213203200321-1121003003312322-1310032330230232-0303010320133131-0320110222223321-0131031213323210-1320301203200221"></a>

<a id="canonical-3111020332332032-0233211312023130-1200123220312222-2323220010030332-0101021130113333-2022122121313211-2223013230230120-3303030233132030"></a>

## xfcc_header_elements property — xfcc_options / 031131330103 / 4

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

<a id="canonical-1133021103232122-2330033023123121-2000322111321302-3231203321330210-1312321103330031-2323322330002011-1223012010213133-2311133331110330"></a>

## Next pages — xfcc_options / 031131330103 / 5

- [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
