---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-3312201232222031-0313110003202330-2120303303020110-1230101202211110-0231222001103032-1030031010301132-1233112011211021-1032230232331212"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options — xfcc_options / 001030313021 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-0212332210313031-0032001310031302-1113322132113321-0121200130233013-0333220210322331-2020333113033223-0320213233100031-1332233013120303"></a>

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

<a id="canonical-3103210003310101-3002020112333103-3131011020103023-0021123331003121-3013001331101102-0313320001032313-0321021323003313-0000301232011031"></a>

## Direct properties — xfcc_options / 001030313021 / 3

<a id="canonical-2011010030320310-0101110311022121-1203120331220103-0313110322223331-3021110313131103-0333231301121101-2203330103303000-3001113320313312"></a>

<a id="canonical-0223003200300033-2013103330131233-0301302303223322-0030002111311300-1003032030111310-3031110201112110-2300112321122121-3222313021202313"></a>

## xfcc_header_elements property — xfcc_options / 001030313021 / 4

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

<a id="canonical-2301022331331211-1030203200023133-3132310300311232-3231022111203122-3100220200232131-1310220200100010-1303201030232012-3302113203200101"></a>

## Next pages — xfcc_options / 001030313021 / 5

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233002110320322-0102010102333101-2003033013313110-2003303300203122-2103003221010122-1131001121123130-2123132120321010-1331020002100021"></a>

## https_management.advertise_on_slo_sli — advertise_on_slo_sli / 312113002121 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.advertise_on_slo_sli

<a id="canonical-3131211132010013-0223123101133212-0231212131233312-1000123022221333-1001103213230001-2301212201113200-2033231301231311-2103100233111011"></a>

Type: `"single"`. Computed.

Configuration parameter for advertise on slo sli.

Upstream description:

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

<a id="canonical-2322331321011000-2010230000000103-1101031133230121-2121312010312023-3011113002302322-2012201312012220-2322333203111230-2332230200010332"></a>

## Direct properties — advertise_on_slo_sli / 312113002121 / 3

- [no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2012022332300101-3323303201302013-1323301222200232-3100302230023130-1330212212330031-2013012103320020-3023022320131130-3221321301033110): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001): complete subsection reference.

<a id="canonical-1323032103022230-2322231011101011-3311111222002200-0012222112113212-0313231102232230-1220330222020123-3322311133020220-0132020013133202"></a>

## Next pages — advertise_on_slo_sli / 312113002121 / 4

- [https_management.advertise_on_slo_sli.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2012022332300101-3323303201302013-1323301222200232-3100302230023130-1330212212330031-2013012103320020-3023022320131130-3221321301033110)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2012022332300101-3323303201302013-1323301222200232-3100302230023130-1330212212330031-2013012103320020-3023022320131130-3221321301033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113001020212213-1212010020132101-3123333230013113-2012113221102121-1021032222112213-0121002112033302-1012222302131231-0212333112210001"></a>

## https_management.advertise_on_slo_sli.no_mtls — no_mtls / 000003330211 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-1321122231003130-1300322123110300-1020220201321121-2333132000103302-2013131323023332-2231220132132223-2320100200322300-0331220031323033"></a>

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

<a id="canonical-1131323033302302-3101312000100131-2233231220221130-0132300223211300-0031332301232201-1102221210112022-1113010203033022-1021020113333031"></a>

## Direct properties — no_mtls / 000003330211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113222000222313-3111111123210203-1001201003011201-3300121010202013-0323313201103333-1122212003033100-2231201022001302-1211011113300211"></a>

## Next pages — no_mtls / 000003330211 / 4

- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231222332022020-1310013002313132-1313233310330232-2022123210221023-0202011020020233-1320211103121223-1133232103312300-0333010322321220"></a>

## https_management.advertise_on_slo_sli.tls_certificates — tls_certificates / 011221311012 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-3222023331233302-0313302322333022-0221212310233102-3200222230312110-2123022333032330-3113301231331322-1332033212110313-3232022021012231"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3033323203113233-1001023220103313-3123020313231120-0231310302011100-0133031121001333-1020123101202100-1101001312132132-1213033110303132"></a>

## Direct properties — tls_certificates / 011221311012 / 3

<a id="canonical-0010332003012210-0331200330210321-2133011201123012-1330310203330220-1303223200031231-1101013022001201-2120122232212023-2103320003312233"></a>

<a id="canonical-0232333323210133-1003013110223322-0311200103222220-2031211321002112-1110220120333333-1032200011323300-0031330022221022-2113331332210223"></a>

## certificate_url property — tls_certificates / 011221311012 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2222100222232310-2013003031331332-0313213033020011-1133230222212303-2333223201103113-2013212312101133-3211310311133022-0331220301223231): complete subsection reference.

<a id="canonical-3233331210221210-1201302113031023-2011220312131102-0011113211323020-0012002222111310-3021113311101022-1300300231120113-3112021202213123"></a>

<a id="canonical-3110330300113012-1221100022210132-1131303113122333-2130012130320311-2232303120302122-1023111220333123-1223102203320302-2323021103323033"></a>

## description_spec property — tls_certificates / 011221311012 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-1212032211010213-1300113010212023-1211323101023211-2231123120212113-0303000101220112-0221113031322132-0303210023110130-3303113031110302): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-0133101210102013-0312301022020002-2113101030100112-2020212012103230-0020300230032231-3101002300120132-1333020300122310-3020301232020201): complete subsection reference.

<a id="canonical-3310201102320032-2111123200232213-1122001212223301-1012130330030001-0223023213133133-0210310011320201-2300132323233221-3200133331211033"></a>

## Next pages — tls_certificates / 011221311012 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2222100222232310-2013003031331332-0313213033020011-1133230222212303-2333223201103113-2013212312101133-3211310311133022-0331220301223231)
- [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-1212032211010213-1300113010212023-1211323101023211-2231123120212113-0303000101220112-0221113031322132-0303210023110130-3303113031110302)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-0133101210102013-0312301022020002-2113101030100112-2020212012103230-0020300230032231-3101002300120132-1333020300122310-3020301232020201)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2222100222232310-2013003031331332-0313213033020011-1133230222212303-2333223201103113-2013212312101133-3211310311133022-0331220301223231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210031130021202-1332133011220110-0031010023033312-3202231102231231-1222202301332111-3231203012123122-3032300011233021-0100120300113122"></a>

## https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 111300222213 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-1321301323022332-0030100002131133-0223033322200000-0122311100311112-3023232321332103-3001020333313011-2211002022332312-3202213321030212"></a>

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

<a id="canonical-3223013221300310-1101022333120003-2203312321200311-0022222321303033-2011211022011013-3130233301132013-1130023322212333-2101100321310011"></a>

## Direct properties — custom_hash_algorithms / 111300222213 / 3

<a id="canonical-1121330301113311-2130213323201030-0002200303001022-3222110230301233-0132221022001023-2023301220303132-0110002303232101-2333210313200310"></a>

<a id="canonical-1123023301113200-3123221003221222-3133131332110021-2032132130322333-3310122302212101-1131312111001312-3220332302302230-0130231311203112"></a>

## hash_algorithms property — custom_hash_algorithms / 111300222213 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3020010212122332-1001001200311320-3201213323102021-0003320112211323-3031320123202331-2032220302312303-0232011132320323-2102120110012113"></a>

## Next pages — custom_hash_algorithms / 111300222213 / 5

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1212032211010213-1300113010212023-1211323101023211-2231123120212113-0303000101220112-0221113031322132-0303210023110130-3303113031110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200030001330012-2231032023111103-1032103233011133-3100212212002000-3220113101303030-3010033221322300-3110232022000112-2001330202022031"></a>

## https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 200322113301 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-0021311210231103-2310002322101333-1130131301313130-1313223220220312-1001320000200121-1332133220111122-1210221332301210-0002003223230022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-2120322211313122-1132202333220130-1201112013132203-1012022032130002-2121330213122113-3002001033003120-2111311313233102-3022222320001113"></a>

## Direct properties — disable_ocsp_stapling / 200322113301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002322032012333-1110003222103130-2111123313211223-3333032201331220-2021032232310332-0321030032133212-2330002320100110-1020332003101003"></a>

## Next pages — disable_ocsp_stapling / 200322113301 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332200212032010-2022011032021310-3332222021200203-2101100323232332-2321023201001102-1332200221120012-2113311132222221-1212221131010023"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key — private_key / 203321303212 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-2010232010033330-2133223331322131-1002033103313321-1123230100033321-0122223001030033-2200202013012020-3323321202130202-1212030312122302"></a>

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

<a id="canonical-3113220023103233-1030022011013233-3033211331002232-2210010000023131-2023003232122222-3310311000300033-3122231100213111-1133110033102122"></a>

## Direct properties — private_key / 203321303212 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0102020033231312-3211201210113123-2303221302220301-1233233220321202-3103131303032200-1122313020233023-3132330211330300-2031230302213321): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022): complete subsection reference.

<a id="canonical-0320303332210223-0210222211112021-0220323112300311-0021001232300233-3201200210103230-2303332321031001-1200111320203200-2013313031032013"></a>

## Next pages — private_key / 203321303212 / 4

- [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0102020033231312-3211201210113123-2303221302220301-1233233220321202-3103131303032200-1122313020233023-3132330211330300-2031230302213321)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0102020033231312-3211201210113123-2303221302220301-1233233220321202-3103131303032200-1122313020233023-3132330211330300-2031230302213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002132300301031-1012033222100022-2101331020021121-2231100002132102-3303001220010120-2320231020202010-3231332223031201-1333213020322111"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 123031201101 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3013111330211010-1310323113231010-0020212013333233-0332213130210110-1223122312212233-2102300100220321-0022003310213020-2113022201003013"></a>

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

<a id="canonical-0002200221122002-2013223321232301-1031031113133011-0201011000221221-1113003333011131-3022110003001200-3323333213013112-3012321103103210"></a>

## Direct properties — blindfold_secret_info / 123031201101 / 3

<a id="canonical-3112321331322210-0013020320113211-2033103220133202-1313110010121310-1213223023112010-2210202010233110-3301023021211121-3210200330003103"></a>

<a id="canonical-0223102113100200-0013231013012021-3130310003120313-2212123223303231-0313301001130233-0001332333112102-3333321030121320-1321203131113220"></a>

## decryption_provider property — blindfold_secret_info / 123031201101 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1122212111011122-0213031131013010-2201000331031110-1102010000012323-3222222100030012-0200132132102021-2331123203020011-3233131333033131"></a>

<a id="canonical-2112032132103010-3012301013221332-1133330003210000-2221220111220033-3303002033322210-0010132332330231-1330313112213123-3013113113222223"></a>

## location property — blindfold_secret_info / 123031201101 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0212023201302100-3310001221033300-0002330220110011-2300232031213101-0230110331230130-3202131313013112-1102303100303210-1000331211332323"></a>

<a id="canonical-3102133321020200-3002310132101003-3113132013233231-0322131322030102-1120023001220330-1231301122211131-0100033200221131-3201300110033321"></a>

## store_provider property — blindfold_secret_info / 123031201101 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1010323022210023-3213001033231323-2020301013333230-2212203232123312-3333331210321232-3121212302021232-2301001322001021-2023310100202031"></a>

## Next pages — blindfold_secret_info / 123031201101 / 7

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032011223222211-3030220210210232-2230213012313011-2001000013213211-3201121201111330-0013033133123020-3013131010310031-1132211312131232"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info — clear_secret_info / 202323210333 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-3102200103003321-1303323013231331-1122013110333002-1210301123122030-3033313101202132-1311301021302112-1301022033310023-1103302103201223"></a>

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

<a id="canonical-1322031330301101-2212322320101332-1121123003232131-1121312003101132-2222231113202332-2033313003120011-1111221113200102-1313033130000330"></a>

## Direct properties — clear_secret_info / 202323210333 / 3

<a id="canonical-2022120033333001-2123301212310101-3002111223030100-1322023022301333-2200123231203100-2320233223002123-1111110002313230-0332102320130020"></a>

<a id="canonical-0232331123101111-0012101222103002-3232021310130021-1212311231103201-3311211300121103-2112211011233303-1312103313012232-0001302102020301"></a>

## provider_ref property — clear_secret_info / 202323210333 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3221220200233112-0200333223112331-2211201121120200-1120321332203232-2131233202221200-3031212102000303-3023331030323300-3232112200333201"></a>

<a id="canonical-1311031000111313-3123112030222021-3132211003120012-2133003223320213-0200113010030203-2001113133212103-1011112333310100-2333200232223323"></a>

## URL property — clear_secret_info / 202323210333 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2003012013301231-0200333000011212-3331313122001133-0211130010010110-0311301020100222-3100000313230012-0231013202123311-1323221320310100"></a>

## Next pages — clear_secret_info / 202323210333 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0133101210102013-0312301022020002-2113101030100112-2020212012103230-0020300230032231-3101002300120132-1333020300122310-3020301232020201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021302313112013-1123110322121332-2010123022233123-1322330313000021-3311000312023011-2321113101032130-0332311030130203-2010003023010302"></a>

## https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults — use_system_defaults / 311132133201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-1011300131103331-1101222232230232-1112012020002203-2121322200203100-1331102312310223-0320112330212201-2223113220331301-1232301233020200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-3311131002132130-2100110003322112-3323302303002003-1322231101320031-2122311220103202-2321223232312311-2322221310010231-3121020001120030"></a>

## Direct properties — use_system_defaults / 311132133201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202132213022323-3233113033031200-3031011310220330-0031032012130322-2222333331023033-0313120303312002-2100312020033022-3330022222130322"></a>

## Next pages — use_system_defaults / 311132133201 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130110131113030-1020332123023310-2132312012301023-1013123111002212-3320302012301310-2022323000003102-1013121313113232-2300023022321102"></a>

## https_management.advertise_on_slo_sli.tls_config — tls_config / 113300222231 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-0203331200121210-1010331113302133-1223110132233332-3101312231211102-3203121012112221-2220302010011030-0222212022112201-2332333300213223"></a>

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

<a id="canonical-3303011120112320-0011231202203111-1001332020101022-3013203300032030-3110301100023003-1303332012202313-1011222013233020-2213101011320323"></a>

## Direct properties — tls_config / 113300222231 / 3

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-3100231302011200-3012332322212010-0201121023003011-0131101100100332-0110221002330031-0110202001231020-2111020113221310-1232221131020020): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-2230121303013002-1110032200323102-0102100310120112-0333321133002023-2221120220001303-1323100121112210-2123022030020020-1113320123022223): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-003.md#canonical-0321132112313203-1121123332132013-1111113012300033-0010030031332231-3221001031220130-0321330010100231-3322032122321013-1321202222300232): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-003.md#canonical-3210123103110023-2233011033122122-2221001102310010-2221001111300221-3211012001310232-0313110101211103-3120310322302100-3031310132123211): complete subsection reference.

<a id="canonical-1201323123033110-2321123010302232-1012022121201320-3001200013233331-1132303021012110-2103103011212321-3323333201111233-3210210032020011"></a>

## Next pages — tls_config / 113300222231 / 4

- [https_management.advertise_on_slo_sli.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-3100231302011200-3012332322212010-0201121023003011-0131101100100332-0110221002330031-0110202001231020-2111020113221310-1232221131020020)
- [https_management.advertise_on_slo_sli.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-2230121303013002-1110032200323102-0102100310120112-0333321133002023-2221120220001303-1323100121112210-2123022030020020-1113320123022223)
- [https_management.advertise_on_slo_sli.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-0321132112313203-1121123332132013-1111113012300033-0010030031332231-3221001031220130-0321330010100231-3322032122321013-1321202222300232)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-3210123103110023-2233011033122122-2221001102310010-2221001111300221-3211012001310232-0313110101211103-3120310322302100-3031310132123211)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3100231302011200-3012332322212010-0201121023003011-0131101100100332-0110221002330031-0110202001231020-2111020113221310-1232221131020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020312023123023-3311022031330023-3232123000223233-3330022221132000-1303101233033030-3223031211103132-1012031220302131-2103013301123312"></a>

## https_management.advertise_on_slo_sli.tls_config.custom_security — custom_security / 100000010323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-2003310123302022-3031022321010231-1333130130333103-2232003202202030-1310013131300333-1200303020223010-0013332312011010-1220033210010011"></a>

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

<a id="canonical-1010001131223323-3131330233120033-0021020013223223-1210011333133021-0010003302111020-2222031032101120-2230120001123321-3021012313222201"></a>

## Direct properties — custom_security / 100000010323 / 3

<a id="canonical-3131101130000202-1333202011001023-2213000332123230-3321210310321211-0120112321312202-1031011320133203-3210023220320021-3013131113001013"></a>

<a id="canonical-2031321231110333-1103331020021322-3333321112031203-0312233122200123-1220232330102323-2333013130120020-3002002230121133-1211320122020210"></a>

## cipher_suites property — custom_security / 100000010323 / 4

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

<a id="canonical-2011213120131012-3003130120232020-0301101201220211-3320111022223203-1231220321103033-0001210213001233-2202331012332113-2033122020311232"></a>

<a id="canonical-3312333003211302-1202120131102110-1222230102301323-1232101213031103-2213101322122212-2130201121212120-0120322311320001-1023233211133130"></a>

## max_version property — custom_security / 100000010323 / 5

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

<a id="canonical-3212120022230003-2200332320323032-0201133302320001-0101233322032113-2101231010130023-3212001032220013-2131311022131322-0213131101133122"></a>

<a id="canonical-2302332031331031-2213113300223301-3132310113313232-3020333123231313-1002110012133300-2132101112123113-1321220321101323-3302223120333013"></a>

## min_version property — custom_security / 100000010323 / 6

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

<a id="canonical-2023302313300122-3333332213031020-2220102301023202-0032030130012100-2220301110022330-3231303323302210-3022001330233332-2112103102212023"></a>

## Next pages — custom_security / 100000010323 / 7

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2230121303013002-1110032200323102-0102100310120112-0333321133002023-2221120220001303-1323100121112210-2123022030020020-1113320123022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120010230021113-1100332310000202-3013211332301211-1032322033201330-3320013013313221-1002132111231223-1111111203223002-0310030333222231"></a>

## https_management.advertise_on_slo_sli.tls_config.default_security — default_security / 220123230111 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-0121021131322000-2220310131010230-0212113211102310-3322020111010211-3103232102321111-1101211203003332-2000103012223301-0223133231001233"></a>

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

<a id="canonical-2221231301030332-1111132312213110-2201202210301031-0300011032301122-3310203330321011-3032302313123330-0020211323310000-2032103100223333"></a>

## Direct properties — default_security / 220123230111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023202033213022-3113210002230120-0103203302111133-2022203100111033-1201310213000230-1102113131322031-0123132210200213-3132021032022030"></a>

## Next pages — default_security / 220123230111 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0321132112313203-1121123332132013-1111113012300033-0010030031332231-3221001031220130-0321330010100231-3322032122321013-1321202222300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300231212113121-3131213100330100-1202301023132123-2323110323032213-0110221200211121-2322111300133203-2322010330002001-1032323323111003"></a>

## https_management.advertise_on_slo_sli.tls_config.low_security — low_security / 112032333001 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-3221103332010201-3130003221220300-0020020213203002-0303311312032012-2033100312303223-2220333303122320-3321012200331030-0330211212100212"></a>

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

<a id="canonical-0100221230013310-1122203133001332-2330231310213022-0001002323303110-2223031031012030-3302000211123123-2321132110131312-2132120111032230"></a>

## Direct properties — low_security / 112032333001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101022213102221-2313023110333111-2122031213002310-0030003212202313-1221002311010023-1020010312321003-1332123003220203-0110012031330230"></a>

## Next pages — low_security / 112032333001 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3210123103110023-2233011033122122-2221001102310010-2221001111300221-3211012001310232-0313110101211103-3120310322302100-3031310132123211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021012313001330-1131311323020231-0020001332133110-0331101023000023-3011320331003300-2131332301330032-0031111310012130-3211231100310110"></a>

## https_management.advertise_on_slo_sli.tls_config.medium_security — medium_security / 012123011200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-3300332300100133-3110222201010132-0030300131320200-0213033221101010-1133210000223322-2303213302330013-0021322331230331-2201021323230302"></a>

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

<a id="canonical-2201211020323020-3010300203022113-2312030103022110-0022333202212013-0231200232103300-0112031103002113-3310003122003203-1100030310203310"></a>

## Direct properties — medium_security / 012123011200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213320120232220-1212121033132223-0122100010211031-0312211232103231-2313101012331120-1310002200010022-0003300112020013-2022332131320333"></a>

## Next pages — medium_security / 012123011200 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332103300300300-2112322203220002-0122310130111002-1333313232230200-2131101223130232-0232203133220031-0102130301103033-2231110320031202"></a>

## https_management.advertise_on_slo_sli.use_mtls — use_mtls / 230003100231 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-0220301210013231-2031133121322131-2321211303313123-2321200230333121-2131103031021331-3332000122011320-2322321100210003-3102211223333212"></a>

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

<a id="canonical-0323201230233322-1120213020000101-0101311313202202-0121022123010011-2202200312300000-3123032003001211-0221201313300013-0012331300112122"></a>

## Direct properties — use_mtls / 230003100231 / 3

<a id="canonical-1123210220103213-0002301223201301-1303320332033302-0021312232213330-0233022103102312-1011211113003232-0302303323010010-3231032001331011"></a>

<a id="canonical-2231101112323111-3323131310101323-1123323133120030-0021233321000002-3100230221213122-2231212210312122-1322032203120100-2332300000002021"></a>

## client_certificate_optional property — use_mtls / 230003100231 / 4

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

- [crl](data-sources--nfv_service--reference--group-003.md#canonical-2231012030033122-3100121100332200-2111010231133322-1323303201000022-2010023120122201-2110211333230232-1123110131133233-1233012211033203): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1310000021221121-2032333013130313-0110330133211200-2332221132231321-0300111223331310-2220102221030320-0103112020122322-0020300031232320): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-1202012222213033-1111202132312301-2203231331030310-2322322120332330-1330031100231220-2220012132310323-2033322100100032-1331222212323332): complete subsection reference.

<a id="canonical-3332101211031221-3302201201320010-2300211121103020-3311210223133013-3013231103121313-3321233213233231-0003001010132331-1122330310010301"></a>

<a id="canonical-1212221222231202-3110121202332231-2120303100311132-1002133002221012-2103013022221021-2110313332233233-0100001000330130-0201303110310121"></a>

## trusted_ca_url property — use_mtls / 230003100231 / 5

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-2001120122302233-3113012010021131-1023130131302032-2330332233131220-1022100121110221-0003220211022012-2221120213331102-2132130111101002): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0033021311113333-3302323211310210-2332213113112032-2000223203331002-3313031020032003-0330322322203202-0012130312220323-0101013112213022): complete subsection reference.

<a id="canonical-1220111123001111-2033313030103023-2230132002233113-1020300231330102-3310301103130021-0102201030122031-3122122013023112-3112313300213332"></a>

## Next pages — use_mtls / 230003100231 / 6

- [https_management.advertise_on_slo_sli.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-2231012030033122-3100121100332200-2111010231133322-1323303201000022-2010023120122201-2110211333230232-1123110131133233-1233012211033203)
- [https_management.advertise_on_slo_sli.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1310000021221121-2032333013130313-0110330133211200-2332221132231321-0300111223331310-2220102221030320-0103112020122322-0020300031232320)
- [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-1202012222213033-1111202132312301-2203231331030310-2322322120332330-1330031100231220-2220012132310323-2033322100100032-1331222212323332)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-2001120122302233-3113012010021131-1023130131302032-2330332233131220-1022100121110221-0003220211022012-2221120213331102-2132130111101002)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0033021311113333-3302323211310210-2332213113112032-2000223203331002-3313031020032003-0330322322203202-0012130312220323-0101013112213022)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2231012030033122-3100121100332200-2111010231133322-1323303201000022-2010023120122201-2110211333230232-1123110131133233-1233012211033203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010333320223201-2311231033120222-2201223331110020-2220233013103323-0303111202333313-2303022122110122-1300122233130101-2133200111230022"></a>

## https_management.advertise_on_slo_sli.use_mtls.crl — crl / 111100023200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-1030211031122330-3300021113200032-1201113121300013-0320000010120221-0011312021003132-0321030112211320-2221313231201121-1022103031202123"></a>

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

<a id="canonical-0010001222323311-3310220220111201-1032011312313110-1330233311011020-1320013230221331-3230223113211001-0231223103020332-2230233033022322"></a>

## Direct properties — crl / 111100023200 / 3

<a id="canonical-3312103323122202-3101213120330233-1010033020031001-2331030331112122-2333320112020201-3112131333200211-0002132031223322-0110203220202131"></a>

<a id="canonical-0131201001220123-3133312332123110-3023111202020202-2003011331210301-2113132200200100-3311113002031101-3133021330102322-0122203103100123"></a>

## name property — crl / 111100023200 / 4

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

<a id="canonical-2113331133210320-1212322022132230-1310002023132313-1133202322220032-1103311031131212-3332031101013323-2030003000300101-1230022011122011"></a>

<a id="canonical-1012330302310023-1133131010233320-0310011001020322-0200301221210223-2030210302303312-3221201131213122-3030023300221023-1220001023122030"></a>

## namespace property — crl / 111100023200 / 5

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

<a id="canonical-2010303300211231-3130302310322031-0310302211202010-0300023323222232-2032232221323020-1121320202132131-1203330203132303-2011333123332000"></a>

<a id="canonical-0010332022331212-1031121001230112-1231331323102121-0132210300231301-2112203302312011-0303321031220112-0211132112311122-0231333211032121"></a>

## tenant property — crl / 111100023200 / 6

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

<a id="canonical-2023002213100202-0021331301203301-0010002210310003-1111303000300222-2031023302333220-0112211132200313-0231311232330011-2023000030333333"></a>

## Next pages — crl / 111100023200 / 7

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1310000021221121-2032333013130313-0110330133211200-2332221132231321-0300111223331310-2220102221030320-0103112020122322-0020300031232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313010122332231-2122222023102331-0110030211310022-1320013331013032-1021322230233201-0230120210022130-2113313022000003-3330201303320120"></a>

## https_management.advertise_on_slo_sli.use_mtls.no_crl — no_crl / 012220002123 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-1103022002012221-1011003023310022-3101111122130111-0223010003023020-1131320030220203-2213002111033012-3022033003111112-3210211330010030"></a>

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

<a id="canonical-2301233133113221-1232201203333000-0322222213230202-1133322313320021-3131330310031132-2233121121223011-0331310031201312-3213120133131103"></a>

## Direct properties — no_crl / 012220002123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211121200023331-3330201000103313-1023122331022123-2202102201020111-2113100112002202-1323211332231032-1231230302312100-1210320003113232"></a>

## Next pages — no_crl / 012220002123 / 4

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1202012222213033-1111202132312301-2203231331030310-2322322120332330-1330031100231220-2220012132310323-2033322100100032-1331222212323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311300100133123-1012222331300310-2213330303010312-1313330033110200-1011022032221001-1311333332211010-2213332123311232-2223002112102330"></a>

## https_management.advertise_on_slo_sli.use_mtls.trusted_ca — trusted_ca / 230221002212 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-2323320112132010-2030320110022023-1023110313032123-1031010112113332-0030011211211310-0202033011120201-3203030032303120-0011231232111112"></a>

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

<a id="canonical-0231021211133111-2302133011323213-2131122133313220-0112201001113330-2133210111331131-3220300230303221-3210311310221303-1223332301123210"></a>

## Direct properties — trusted_ca / 230221002212 / 3

<a id="canonical-1032102323222210-1021132002133311-2332331221320313-3112230213220211-0231211103021231-0221022210230000-0201223010231330-3023132223313022"></a>

<a id="canonical-0000221300011113-3111231032010002-3133331032102302-1133212231000130-2203033110213221-1003312303303302-2213003213210210-2322333310023231"></a>

## name property — trusted_ca / 230221002212 / 4

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

<a id="canonical-0331311121231131-3031113302131210-0002203003122023-3010010311021013-3232000030100113-2133120100001003-0021023230013000-2203213121133011"></a>

<a id="canonical-3133322121101330-1102202103333221-0201321223211321-3202021112222112-1233230332333131-3222322021131120-2030302110300030-2333100232023030"></a>

## namespace property — trusted_ca / 230221002212 / 5

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

<a id="canonical-3232133223312131-3330132031300022-1022212102011202-0112131330013300-2013020010023302-3112012212231022-1002213020032122-1033311111233100"></a>

<a id="canonical-2020020131122313-1221233301030000-1122020320311100-2023102201311111-2130120121120111-1130212013112111-1021022102110322-1322231233231213"></a>

## tenant property — trusted_ca / 230221002212 / 6

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

<a id="canonical-3220131003000023-0223013223102132-3130300001233323-2011123333321130-2310212212031221-1113101131023202-1201331112131330-2301131312323203"></a>

## Next pages — trusted_ca / 230221002212 / 7

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2001120122302233-3113012010021131-1023130131302032-2330332233131220-1022100121110221-0003220211022012-2221120213331102-2132130111101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031313311210030-3312030131331310-1102222101212111-1030103002233120-2003310210211322-1131202031001222-3321221323331300-1122332130323221"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled — xfcc_disabled / 301130310021 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-3103311232022002-2312013011333210-3003200010031311-0203032221101331-1233201002213213-0300130200123013-2132030112233220-0013212223310122"></a>

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

<a id="canonical-0033323000302023-1300110313210120-2322302012133320-2110201120332311-2032300332023321-0330232313322033-0102122320202313-2201301001200003"></a>

## Direct properties — xfcc_disabled / 301130310021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002123210312121-1232130010021303-2302201311213112-1003303220331010-1100020132321021-2101302121100311-1020210210001033-3001030131102001"></a>

## Next pages — xfcc_disabled / 301130310021 / 4

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0033021311113333-3302323211310210-2332213113112032-2000223203331002-3313031020032003-0330322322203202-0012130312220323-0101013112213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323213312313121-1213030211320022-2032020110303030-3313101232323123-1321131202110311-3321102313011330-1032312000103301-3011130120323011"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_options — xfcc_options / 331011311210 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-1302302203011030-0000223002133131-3223123331103311-3202223313033231-1130130330313121-2311313031100212-2111122111101000-1001121202230323"></a>

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

<a id="canonical-3023123221210321-2001321123110301-0303331311332111-1032320223010120-3001231320213200-0012000100320000-1131323210302033-1200110110322301"></a>

## Direct properties — xfcc_options / 331011311210 / 3

<a id="canonical-3123223231030223-0000100000012111-0132202320002301-1322331223213123-3110221033012013-2122101312300132-2013202111000111-3012123101202332"></a>

<a id="canonical-1120311102220000-1201321112301113-2010321311120022-2033022331201331-2222113010101000-0203331021201313-1000202323313211-2123223310130003"></a>

## xfcc_header_elements property — xfcc_options / 331011311210 / 4

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

<a id="canonical-1223011113332013-1223101023013112-2300103012222232-0012322133323223-0121331320022121-3221123113022021-0103210332321231-0110320300210131"></a>

## Next pages — xfcc_options / 331011311210 / 5

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021032320223210-1213112002110302-0222321300032130-0232000211132113-0030001232231213-0201031301130312-3213021233200211-3210303113003120"></a>

## https_management.advertise_on_slo_vip — advertise_on_slo_vip / 122130321121 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.advertise_on_slo_vip

<a id="canonical-1201300311232210-3120333123332212-2021000011210110-3212320212312103-0200203032021102-1102331330302300-0230103102031032-3001223111313130"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

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

<a id="canonical-0221022030301113-1333030112112103-3130221033111011-3302000130231212-2103221013333210-3203113122133112-3123023110320231-3230230013131110"></a>

## Direct properties — advertise_on_slo_vip / 122130321121 / 3

- [no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-3220211322023012-0001232100122300-2230133210110122-3120020203121220-2031310222201120-1200032123320032-1300012333132002-0010320013000121): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300): complete subsection reference.

<a id="canonical-3102132300311213-2221333313021111-2302132312300320-1222100010002131-3203132010022133-1000313330213313-2303320200011110-1120102311330120"></a>

## Next pages — advertise_on_slo_vip / 122130321121 / 4

- [https_management.advertise_on_slo_vip.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-3220211322023012-0001232100122300-2230133210110122-3120020203121220-2031310222201120-1200032123320032-1300012333132002-0010320013000121)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3220211322023012-0001232100122300-2230133210110122-3120020203121220-2031310222201120-1200032123320032-1300012333132002-0010320013000121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131331021222212-0213003223303133-3100020230311210-0101121132033223-0033023111012100-2313022313212302-3112312210010123-1300201222232133"></a>

## https_management.advertise_on_slo_vip.no_mtls — no_mtls / 003302311130 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-1220200323331332-3032011231113212-1312000202230132-1011301100021333-0033311332023202-0111021031311213-1130311120121212-0300120223131223"></a>

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

<a id="canonical-0133120213220312-1032133013300013-3223010303231111-0331203321333022-2223012333022200-2101010231330113-0101123110221232-3331303230330121"></a>

## Direct properties — no_mtls / 003302311130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302211010330311-2203113320020021-2310111320303133-1131220301030300-3023103033212021-1333303131023022-0221210022300101-1123030112312021"></a>

## Next pages — no_mtls / 003302311130 / 4

- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130031130303032-0330231201200200-1033333333330223-2020122203131132-2030212103002022-3333010113020312-0003030330133123-0033123221030023"></a>

## https_management.advertise_on_slo_vip.tls_certificates — tls_certificates / 121331331323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-0121122103200222-1012011112210020-3110030201112113-1012023330221331-3211121123032211-3233211021301202-2123020032213103-2022312131000303"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3223321020301103-2303030113031311-0230300320232013-3202200303212300-2033213321112111-0213132322202323-0031021101000100-0200212113211310"></a>

## Direct properties — tls_certificates / 121331331323 / 3

<a id="canonical-0322323131220131-0002320112021220-3232331022302301-3103332210210312-2301321223012031-0010200132003211-0123213200211031-2003221121202331"></a>

<a id="canonical-1001131011020131-1021102021001021-0322110100100313-0130201111000103-2021202132010103-0301313210121030-0200103322302233-1320310310020033"></a>

## certificate_url property — tls_certificates / 121331331323 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-0221000212202130-3210313010022010-2232113210121100-0011002030223301-1232233133331213-1032132202102110-2321002220230203-3120032222320220): complete subsection reference.

<a id="canonical-3123200033312303-2110122033222221-2331313212332013-3221231103322023-0022203331020131-0132232310003130-0232203022021221-3312230311310233"></a>

<a id="canonical-0311022222303030-1320101021031101-0231301322130322-0112100110220333-3100313133212120-0313322203010110-0110031110311301-3012011330001221"></a>

## description_spec property — tls_certificates / 121331331323 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-3103033100313002-1001200330133000-3123021313000313-3232312133120101-0203013302111030-1103330102002202-3201320312131321-3221102010023233): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-3120001133113131-1121113132221211-2220223333332210-1133330221220002-2332133022211320-3110123332103223-1012132320111232-1232332131231111): complete subsection reference.

<a id="canonical-2133331231332112-2113333012310231-2011113033123131-3101333003223323-2012130003321011-3222303212121321-1020331031033301-2103121303000121"></a>

## Next pages — tls_certificates / 121331331323 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-0221000212202130-3210313010022010-2232113210121100-0011002030223301-1232233133331213-1032132202102110-2321002220230203-3120032222320220)
- [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-3103033100313002-1001200330133000-3123021313000313-3232312133120101-0203013302111030-1103330102002202-3201320312131321-3221102010023233)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-3120001133113131-1121113132221211-2220223333332210-1133330221220002-2332133022211320-3110123332103223-1012132320111232-1232332131231111)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0221000212202130-3210313010022010-2232113210121100-0011002030223301-1232233133331213-1032132202102110-2321002220230203-3120032222320220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300313233202033-3322020103102110-2300111331332112-3231020323212223-3123132333002222-2323012203130201-0020331111301320-2031332200221303"></a>

## https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 112030203311 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-1221210211321123-0213132013332310-0012011303330123-2232330212002132-3212320012133301-3022110111320211-1131031333300023-0210010213323010"></a>

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

<a id="canonical-1310302322110330-1323300122011101-0030123112221221-0321310320111011-1121231131113222-0023113323210201-0231023302033132-0111022222301102"></a>

## Direct properties — custom_hash_algorithms / 112030203311 / 3

<a id="canonical-2302211003133230-3221322033323131-3003230303233231-2003232320002003-2011002333101230-1212200322011312-0210003220303101-2321132113312333"></a>

<a id="canonical-0013100313131122-1011300221101031-2010022031221303-0113202232030302-2110302120320000-3000102231213030-2312032330211311-1120301320223123"></a>

## hash_algorithms property — custom_hash_algorithms / 112030203311 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320230020120131-1233132130131300-2032131102312013-3001303232120020-3030223010131122-1233122032011022-1031113102123123-1320230031212232"></a>

## Next pages — custom_hash_algorithms / 112030203311 / 5

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3103033100313002-1001200330133000-3123021313000313-3232312133120101-0203013302111030-1103330102002202-3201320312131321-3221102010023233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312332301331232-0300312202033302-0111202233212320-2132031112201222-3123113323210233-3210122112232330-3332030123320111-2302213331023003"></a>

## https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 203000012301 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-2023321032211332-0103323222331103-2030223023200111-0231130223102001-3230000021032000-3122211131230232-0111220220031111-2122010233012110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-0323200033332320-0033330222311032-3023212102030202-1221110002331201-0322200011221330-1211013032300000-3123222320100332-1120122131203000"></a>

## Direct properties — disable_ocsp_stapling / 203000012301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313031123222033-3020001032301030-3303033323310322-3011001002020001-1233020030222213-0133032121132231-1132222232213302-1220111010132033"></a>

## Next pages — disable_ocsp_stapling / 203000012301 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213223221202111-2021300010121130-0202321003022302-2310231133212032-0323220030133320-3221312230231323-3223000133220203-2032003102330331"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key — private_key / 131230331300 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.private_key

<a id="canonical-3123231300111133-3110002301033002-3123233302201310-0310111320110032-2002320013002103-1032022011231322-2010301212332331-3211212321122010"></a>

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

<a id="canonical-1300103112231231-3201210123331211-2122311013321202-1113003322323303-2001012130210200-2322221113313213-3330121101102311-2020031030103033"></a>

## Direct properties — private_key / 131230331300 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0031321200102323-3033032000210210-3120313102002310-0230031222323222-2233211332203121-3032011233023200-2323220132023030-0022310322222210): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1120022112132112-0123020212030331-1311122312300330-3123320100221133-1320031221221133-1010003212000020-0312322003022132-2003130122310131): complete subsection reference.

<a id="canonical-0211003120323323-0002300013203321-1310110313113101-3002211032232232-3131230022102002-0021331032102022-3111201120200211-2320101331222120"></a>

## Next pages — private_key / 131230331300 / 4

- [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0031321200102323-3033032000210210-3120313102002310-0230031222323222-2233211332203121-3032011233023200-2323220132023030-0022310322222210)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1120022112132112-0123020212030331-1311122312300330-3123320100221133-1320031221221133-1010003212000020-0312322003022132-2003130122310131)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0031321200102323-3033032000210210-3120313102002310-0230031222323222-2233211332203121-3032011233023200-2323220132023030-0022310322222210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333103311132320-1101310212330000-3023200003111123-0302000210130210-0213031332002002-3320201301301332-0120132112110331-0133223023113312"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 132113233330 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3011320311322121-1221001133201211-0330102213303330-1230130220032130-2221130002331221-1111201303222201-3302310010021133-0211023002331111"></a>

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

<a id="canonical-1323323131111112-2312211133303022-0330202230203130-1223213013002302-3313001221311223-1000010123200122-3022222302122202-1311000320013011"></a>

## Direct properties — blindfold_secret_info / 132113233330 / 3

<a id="canonical-3122211320123100-0333111203023112-2022331012133010-1030202301011023-2230330131321032-3211010302120120-3113213011231130-0212110001211212"></a>

<a id="canonical-0030323321301102-3303100333032022-1122023230000130-1133122003331321-3203213111130131-3020031330223012-1121201323230102-1212110110112123"></a>

## decryption_provider property — blindfold_secret_info / 132113233330 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2201302300211210-1001313133331333-0222223200112033-2030133101323223-0021101303313032-3211132022032203-2331022010123131-1030111310232011"></a>

<a id="canonical-2130101303110113-1210302000033113-0202322213322123-0331023333233130-2122223320222111-3030301010301103-2130012011332103-1213110302300310"></a>

## location property — blindfold_secret_info / 132113233330 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1332200202330211-2033111200022200-2121333221212013-3212321212221231-1202130312133021-2210032201203233-0032221212011220-1200133112331210"></a>

<a id="canonical-1231132221320212-2211100110101222-1111211033100123-1230003013312120-1001012013313213-2001101101322012-2120221330311011-2110312332321311"></a>

## store_provider property — blindfold_secret_info / 132113233330 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2322310223102232-1112213233021110-3032220103301111-0220222223323011-1323112132222331-1220133131220331-1023102221210200-2200310003130130"></a>

## Next pages — blindfold_secret_info / 132113233330 / 7

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1120022112132112-0123020212030331-1311122312300330-3123320100221133-1320031221221133-1010003212000020-0312322003022132-2003130122310131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210222110231320-2102112211230211-3203112311001113-1020321010223121-3203322123323301-1000200023033031-2103031311121212-2220301121312302"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info — clear_secret_info / 330231211203 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-2321112223130110-1300120031131131-1330232121012122-3330303311000200-0131011223223121-2130122313300320-2232112032103203-1322233313303312"></a>

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

<a id="canonical-2001001322130210-3111301001111320-0333202232233130-2232233320032023-0112001333122011-3033120000222322-3021201003033112-0121011010022311"></a>

## Direct properties — clear_secret_info / 330231211203 / 3

<a id="canonical-3202321232020211-1123033003113021-1333000330131303-3322112112132202-3303130122320112-0030300132300122-2112222303210023-0111021001211013"></a>

<a id="canonical-2121101230211102-2022201220013120-0201121102003123-1113202103002100-3132220233321111-3332021300102020-1301311233300233-2031131300010231"></a>

## provider_ref property — clear_secret_info / 330231211203 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0121211120313132-1121222130130003-3223311113021333-2221102312001112-0033300113003032-0210022111223333-1212111120011322-1331021011132120"></a>

<a id="canonical-1302201130001222-1320021222302102-0103332031300123-3023001220112113-3233032001121310-1333010230101330-3121132003000113-3312010020030320"></a>

## URL property — clear_secret_info / 330231211203 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3332223330013332-0222310122322003-3310030332312321-0212012213000333-2023033201301332-3233222313321003-0123220300022313-0131220111203201"></a>

## Next pages — clear_secret_info / 330231211203 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3120001133113131-1121113132221211-2220223333332210-1133330221220002-2332133022211320-3110123332103223-1012132320111232-1232332131231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110331012121232-3320313221032012-0111121003201300-2323302210301301-0030102231002023-1232010002210102-1220023121122302-3302012010001312"></a>

## https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults — use_system_defaults / 231023101221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults

<a id="canonical-2113020101012132-1303202132222230-1102010311001030-2033300013033310-2101111202131100-1232223233133230-2320013211022331-0123322321000103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-0222022110130202-2000313100231220-1230231103103332-0022221112203132-2100021222132130-1201003301303023-0330201132220132-3221320000020323"></a>

## Direct properties — use_system_defaults / 231023101221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100202010301022-1113030211032133-0200300301313232-2332210230311331-2323120220221212-2113031012302123-2300201021321213-1320202230232020"></a>

## Next pages — use_system_defaults / 231023101221 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002013220100213-0012200311122032-1120320321203213-0130022012003012-2113212211202300-1220211023023330-3330200031323103-3332100303210131"></a>

## https_management.advertise_on_slo_vip.tls_config — tls_config / 012123033011 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.tls_config

<a id="canonical-0331311322030123-3123133111123223-1321020012313323-3133130002311013-1233121302202011-0213001010011132-2011231203312111-1113211312013310"></a>

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

<a id="canonical-2032030013030022-0123231303320022-3030223301320132-2202331212231123-0032123313331201-0222313232130333-0213013002322322-0300120133123203"></a>

## Direct properties — tls_config / 012123033011 / 3

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-0103313302212330-0232203133213031-1003020033121220-3102223121103031-0132120121110003-0121003220222031-1310033332203001-1033222120220333): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-0233013023232122-2002020013231313-2022201012303003-0222102212221232-2310131321230300-0133032311222301-0000002200101300-0123132333110211): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-003.md#canonical-2200001211220003-2202201113331331-1020230032121033-3001123000131233-0311231211131321-2010221332323310-1332322311011030-2003003121012001): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-003.md#canonical-1212101031202310-0021203331222312-3012300120123120-0223232312230320-3102201023221033-0233031213111021-0010032010110131-3211202131303121): complete subsection reference.

<a id="canonical-1130300103232032-2200001313310022-1111100102121312-3130222222122013-3003110310221121-0320001313303233-2003002221210123-3030002101113210"></a>

## Next pages — tls_config / 012123033011 / 4

- [https_management.advertise_on_slo_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-0103313302212330-0232203133213031-1003020033121220-3102223121103031-0132120121110003-0121003220222031-1310033332203001-1033222120220333)
- [https_management.advertise_on_slo_vip.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-0233013023232122-2002020013231313-2022201012303003-0222102212221232-2310131321230300-0133032311222301-0000002200101300-0123132333110211)
- [https_management.advertise_on_slo_vip.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-2200001211220003-2202201113331331-1020230032121033-3001123000131233-0311231211131321-2010221332323310-1332322311011030-2003003121012001)
- [https_management.advertise_on_slo_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-1212101031202310-0021203331222312-3012300120123120-0223232312230320-3102201023221033-0233031213111021-0010032010110131-3211202131303121)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0103313302212330-0232203133213031-1003020033121220-3102223121103031-0132120121110003-0121003220222031-1310033332203001-1033222120220333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022330121321220-1231221330212111-3330130001333332-3220023121120213-1010323303203310-0102213132202003-2121001220112230-0120233120232302"></a>

## https_management.advertise_on_slo_vip.tls_config.custom_security — custom_security / 311300000021 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.custom_security

<a id="canonical-3302111021202022-3232013023000120-1321033123102103-3002320312322001-3003112130302021-3202310211000131-2011022332212301-0211232211010110"></a>

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

<a id="canonical-1022323312100021-0323001210000001-1301232333233313-3031123110302031-1100033311013111-0223223213030012-2330332020111023-2002302123313323"></a>

## Direct properties — custom_security / 311300000021 / 3

<a id="canonical-0120031320113330-1320031212132112-1303331130121003-3300022211100111-2000131300102320-3112301022100211-3111010033201332-2010123310330023"></a>

<a id="canonical-2313232021311030-1311320003100220-3333233023033300-2101013023003001-0210321123302000-0200302223213220-3222122203032032-0111113033230300"></a>

## cipher_suites property — custom_security / 311300000021 / 4

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

<a id="canonical-2311111223013322-2322021130101320-0210000303302121-0233011213331320-2020221233213003-3131111001022032-3210231113331221-0130312220310333"></a>

<a id="canonical-3020301330011131-2121011022133031-1323302022302223-2301012013320131-2022330122332332-2012201302330332-3022010203020102-0003210323310333"></a>

## max_version property — custom_security / 311300000021 / 5

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

<a id="canonical-2102113320230302-2212022233111301-1220300111302322-0221023133011023-2012033122312100-1330313131210110-2332323222011130-3111000130030010"></a>

<a id="canonical-2230301323312122-3003132210022313-1120233031011223-2202010221030023-3201003202200211-0221121133132122-0121301030210121-3002331223222303"></a>

## min_version property — custom_security / 311300000021 / 6

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

<a id="canonical-3022301231120131-3123221030332203-1330100030213202-3132211033031023-3323221230122313-0220031030033032-1312330223012233-2232223031120102"></a>

## Next pages — custom_security / 311300000021 / 7

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0233013023232122-2002020013231313-2022201012303003-0222102212221232-2310131321230300-0133032311222301-0000002200101300-0123132333110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232332021000131-1321001011031003-2321022203230003-2003223113331222-2003000021211201-3212333302102203-3310032330320101-0212131310032110"></a>

## https_management.advertise_on_slo_vip.tls_config.default_security — default_security / 003330310230 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.default_security

<a id="canonical-2200010202331020-1200013310313322-2313030002130232-0121322330210222-2312122231202323-0233333320220120-3120103220121320-3200122331130231"></a>

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

<a id="canonical-1311110322011011-2222310130313021-1321222100323303-1210113113011222-0012311313031103-0220210223110320-2120231121213220-3113223021000232"></a>

## Direct properties — default_security / 003330310230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120321113131221-2203031111002213-3322233013230302-3100210321102122-3310230221000030-0232330300333011-0020030213121020-2300101010302032"></a>

## Next pages — default_security / 003330310230 / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2200001211220003-2202201113331331-1020230032121033-3001123000131233-0311231211131321-2010221332323310-1332322311011030-2003003121012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002012212230030-0032031113202211-0133232131131232-3313200302131022-2211213323300231-1112210031133230-0321220202231001-2030122110030100"></a>

## https_management.advertise_on_slo_vip.tls_config.low_security — low_security / 231011223201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.low_security

<a id="canonical-3133102031001330-1133232102133002-0012111012020211-0311202012223202-3000223300202331-3322112211133232-1311122002202211-3131322113203211"></a>

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

<a id="canonical-2321202331311231-2110232201110331-1121221001212010-2033230313101011-2311332013001313-2021121101212220-0130020212222020-1120010233301001"></a>

## Direct properties — low_security / 231011223201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102030113120311-3001333122301130-0323132312203131-1112103130313121-1113131223301331-1102302023303130-2213031203003333-3133233000020230"></a>

## Next pages — low_security / 231011223201 / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1212101031202310-0021203331222312-3012300120123120-0223232312230320-3102201023221033-0233031213111021-0010032010110131-3211202131303121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223210322132212-2232122130302132-2031223023331200-0233131002123030-1211302022303111-3010022230130333-1010323020332003-0303110330233012"></a>

## https_management.advertise_on_slo_vip.tls_config.medium_security — medium_security / 222112232321 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.medium_security

<a id="canonical-3113230120313131-0002323223202320-1213312022212221-1013101021213222-3230022302002201-0230310131221212-0133031022000120-2120230100313321"></a>

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

<a id="canonical-2120023302103032-1313013013013000-2210121332231302-2232230203100010-1023320120200332-1121310103211331-2230022121310001-0300031033202000"></a>

## Direct properties — medium_security / 222112232321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231003311133123-1013121113101233-0201120220231202-1332010332130311-0332222313330330-0213132023101322-2313110122210003-1211330211032312"></a>

## Next pages — medium_security / 222112232321 / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311333233103012-3231000122322213-3033300031002100-1002130133022020-3111312111302211-3323102033013013-0000020211210203-2233213000113101"></a>

## https_management.advertise_on_slo_vip.use_mtls — use_mtls / 220120020113 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.use_mtls

<a id="canonical-0000010113012233-0231013031313230-1002310210333220-2221312211012332-1102321002200313-1222223113011310-1202220320011330-1012321200301023"></a>

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

<a id="canonical-0201312131122011-3321310211003323-0120321323332122-0311323310303110-1222113301320332-0001132033011031-0331122010032200-0021030330311003"></a>

## Direct properties — use_mtls / 220120020113 / 3

<a id="canonical-1213101133011003-2002203131213231-2311122102010132-1203231100131202-3310233000023331-1201103020233031-1231332201313122-2332023302003331"></a>

<a id="canonical-1210312230102010-0203321322202120-2223121023113132-3103100023120102-1213011113222102-1010302123032212-1010100220333210-2311013133300001"></a>

## client_certificate_optional property — use_mtls / 220120020113 / 4

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

- [crl](data-sources--nfv_service--reference--group-003.md#canonical-2113002312310331-0123102300210333-0120133032031301-3111103211131002-2001322100330233-0200112310200302-1033131210301131-2203102021233030): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1102320333000223-2323321202313323-3002020122201023-2201113220320332-1323130101021310-0022210002102110-3322302333031110-1323300033013130): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-3103020122330331-2302202202131002-1221113311301121-0233321212321202-3203222121310300-1333223021112201-3123110130232313-0303311111120202): complete subsection reference.

<a id="canonical-1132330210130302-1233232301102303-1302303010333330-3223321332130332-3032023003302020-3330011023210031-3222310033302102-3003133031123101"></a>

<a id="canonical-3331323210333233-3010032321220110-3300300233111021-2221022223012012-1230111230330230-0233031202220120-0031331313101122-3012121102313020"></a>

## trusted_ca_url property — use_mtls / 220120020113 / 5

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-3022323332031311-0312200321202103-3301201020310102-1103133201303202-0213023022010333-3310312102303232-1100303100313103-3221111133333311): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-1331301201211322-0110003030313003-1201032010210100-1103231323131321-2033210121221331-0121031132210221-2102321211122321-2013230320030311): complete subsection reference.

<a id="canonical-3033033122001300-0021002123321331-1003132130332230-3112030333030201-0331222310220303-3300321012223020-2103123033133102-2022212230312313"></a>

## Next pages — use_mtls / 220120020113 / 6

- [https_management.advertise_on_slo_vip.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-2113002312310331-0123102300210333-0120133032031301-3111103211131002-2001322100330233-0200112310200302-1033131210301131-2203102021233030)
- [https_management.advertise_on_slo_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1102320333000223-2323321202313323-3002020122201023-2201113220320332-1323130101021310-0022210002102110-3322302333031110-1323300033013130)
- [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-3103020122330331-2302202202131002-1221113311301121-0233321212321202-3203222121310300-1333223021112201-3123110130232313-0303311111120202)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-3022323332031311-0312200321202103-3301201020310102-1103133201303202-0213023022010333-3310312102303232-1100303100313103-3221111133333311)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-1331301201211322-0110003030313003-1201032010210100-1103231323131321-2033210121221331-0121031132210221-2102321211122321-2013230320030311)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2113002312310331-0123102300210333-0120133032031301-3111103211131002-2001322100330233-0200112310200302-1033131210301131-2203102021233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020203030212000-2100220333133311-1230030323101232-2322122200302333-2302201310021213-1313121312102302-0023331232230223-3210230102220300"></a>

## https_management.advertise_on_slo_vip.use_mtls.crl — crl / 003012232332 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- https_management.advertise_on_slo_vip.use_mtls.crl

<a id="canonical-2121101103120003-2101323332022131-0213220111201201-2013003331111312-3230203201101023-1231330220300233-3231122302111103-3213000123010131"></a>

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

<a id="canonical-0030010002212123-3321332323221112-3032003223333111-2110201331132233-3300333221312203-1332311311211222-1112331012320102-1121212222221100"></a>

## Direct properties — crl / 003012232332 / 3

<a id="canonical-0320312210333311-1031213201120212-0220222031023031-3130311013332213-1320331230310131-1130301111013212-2332202020221213-2222030322202103"></a>

<a id="canonical-2011113010231111-0211313302200023-0102100023033123-0122102311322322-3032133332111220-0222322322313201-0323023023111111-1310030203120230"></a>

## name property — crl / 003012232332 / 4

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

<a id="canonical-3123021301121120-1133310312200012-2211000301302111-3131031121121123-2302022203202322-3210130100011230-3213312212032101-2113332001201010"></a>

<a id="canonical-0231222223203122-2010333223321210-2201103200130201-3320020122012303-2212032003312211-1333331211120113-1102322121333233-3102032323220002"></a>

## namespace property — crl / 003012232332 / 5

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

<a id="canonical-3020013210220320-3001003320133310-3133021231110233-0112300121232311-2313332301113311-0203330102001120-3012222110102230-1222332102231302"></a>

<a id="canonical-0200132112330203-3001223101113100-0110003002312211-3012320133103312-2211203303321300-3012203122133200-3031001113120103-0021120200120130"></a>

## tenant property — crl / 003012232332 / 6

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

<a id="canonical-0122022323212212-0102200210130123-3111303223300021-2023221313032132-2130020232300323-0101331123011320-2132021323113122-0312212010213203"></a>

## Next pages — crl / 003012232332 / 7

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1102320333000223-2323321202313323-3002020122201023-2201113220320332-1323130101021310-0022210002102110-3322302333031110-1323300033013130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230303222312001-0233113010110211-1203132220122300-1021112223232022-1210033000032311-2120320322300120-2200013102020101-1102320030301221"></a>

## https_management.advertise_on_slo_vip.use_mtls.no_crl — no_crl / 013303000213 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- https_management.advertise_on_slo_vip.use_mtls.no_crl

<a id="canonical-2011333123320013-1301332010323222-3103233230311201-0231020002332200-1111021120023211-0212121131213112-3131122111310211-3003203112131111"></a>

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

<a id="canonical-1013033223020131-1300222002002201-1011132220320230-2332320010130102-1310203113100203-2323030302110023-1100121331200111-0111300313333131"></a>

## Direct properties — no_crl / 013303000213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002000120022103-0210311111012211-0323323010003320-1322013313132122-2113121320310323-2230013321011320-3203320020020212-2113200331333212"></a>

## Next pages — no_crl / 013303000213 / 4

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3103020122330331-2302202202131002-1221113311301121-0233321212321202-3203222121310300-1333223021112201-3123110130232313-0303311111120202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100311201230231-3120023220322213-0110121021312110-0213122333100230-2020213233001232-2331112110101212-3202030331332210-0322030220313213"></a>

## https_management.advertise_on_slo_vip.use_mtls.trusted_ca — trusted_ca / 200230001131 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- https_management.advertise_on_slo_vip.use_mtls.trusted_ca

<a id="canonical-3122023333112123-2322332100022023-2000333101313100-2023211212023330-0230333331231000-0213323220300001-2202303131202211-3201231330033322"></a>

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

<a id="canonical-0221012000010133-0101300233022133-1313021022030123-3302013103032021-2120121211303110-1303213021231200-2122310012322111-2323021033300122"></a>

## Direct properties — trusted_ca / 200230001131 / 3

<a id="canonical-0221201022212102-3001301113332211-0202200213223121-2233300031221033-1333010302032022-3113000100333201-2120200131321332-2301023010122311"></a>

<a id="canonical-2032131203112101-3133101133110311-1023003113223002-2221310120100333-2120111210010012-1221230031011200-0300122302302021-1113312002213130"></a>

## name property — trusted_ca / 200230001131 / 4

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

<a id="canonical-2022032030323020-2130120113122112-2211203120220021-2323321120211313-2010111312321102-1130000120103000-1111110033021002-3010333333323222"></a>

<a id="canonical-3121001223021331-3331131132030100-2201030211200202-3022120210131321-0010302031021330-2230213010023033-1010230230303333-1322031110131123"></a>

## namespace property — trusted_ca / 200230001131 / 5

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

<a id="canonical-2133012010032133-0310333111101020-1120111123013212-1203210023103203-0110312332231102-1133023233311000-0313202223100221-0303323302313313"></a>

<a id="canonical-0101112322210212-0013003231330033-2213120310300011-2023120112221230-0032132002211010-2312100012230112-0223022321300032-2130030013210013"></a>

## tenant property — trusted_ca / 200230001131 / 6

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

<a id="canonical-1100203221000221-1201300030201012-3303020121203201-0032002233203111-3121232312203021-3332210021330330-2113331322011222-3032213013111132"></a>

## Next pages — trusted_ca / 200230001131 / 7

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3022323332031311-0312200321202103-3301201020310102-1103133201303202-0213023022010333-3310312102303232-1100303100313103-3221111133333311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310231322203230-2113021002110110-1101312210000302-0021103301122310-2010112232233223-1213022222031131-1232331212333200-2201120101123232"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled — xfcc_disabled / 131112120323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled

<a id="canonical-2300103102230012-0003231020030332-0303122002220110-3123230023101003-2211231120002021-1233321033332121-3302232022232311-1100032232321302"></a>

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

<a id="canonical-3300123023001021-2233110102111000-1333001020311321-1023113120030030-1123133210322010-1202321001302003-2122310213202100-1303122203110202"></a>

## Direct properties — xfcc_disabled / 131112120323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222212233002232-2203112231223330-1220330223233312-3301222333113110-3201213321311123-2012300312313202-2032231111103232-1013333031302323"></a>

## Next pages — xfcc_disabled / 131112120323 / 4

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1331301201211322-0110003030313003-1201032010210100-1103231323131321-2033210121221331-0121031132210221-2102321211122321-2013230320030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021113130203212-2322202223212200-2030130212113223-1212313311021201-2000213332221022-0221233200233123-0232131123103213-1313001103120313"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_options — xfcc_options / 213310223223 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_options

<a id="canonical-2230321013321121-2230030032332200-0020111102220121-3230221201200123-2031323031111213-0313220102332003-2300330110320103-0310032301321230"></a>

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

<a id="canonical-2102010320020312-1213032321210032-2213023011200102-1201212031210203-0211302300021123-1230320300221323-2330012200131030-2332303302212303"></a>

## Direct properties — xfcc_options / 213310223223 / 3

<a id="canonical-3301020001203301-1120002212030212-2212021121102322-2223011023112321-0120333223221333-0022123233322032-2101223223231220-2233230321002032"></a>

<a id="canonical-1013022203320013-1331031032311112-1121332123232322-3133311010031333-0001011013032132-2221121222202200-2300213010130200-3112233021321110"></a>

## xfcc_header_elements property — xfcc_options / 213310223223 / 4

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

<a id="canonical-1013202113133100-0112102231001133-0231121021232302-3222123213013130-1310122112123002-1220022113233000-2021011113332113-0220332110301312"></a>

## Next pages — xfcc_options / 213310223223 / 5

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0132221201130030-0232000021033212-1032332030010130-3332223123302221-0203232211320032-3130232202230111-3323311200123201-2222003003133213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223312120031020-0023031213203113-2232321322331110-0303003301020000-2120122132302013-3322011003233113-2032012110210120-1011030011221311"></a>

## https_management.default_https_port — default_https_port / 301130100323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.default_https_port

<a id="canonical-1031231313033223-2323221332031020-3012130221213012-2133111033212130-2221030321332220-1012301313121210-2203113230230000-3112120101323131"></a>

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

<a id="canonical-3333323032300101-3011000331303113-0230022310223330-0301031201102312-2121110300030312-1203002211331302-1320020321200230-1323313113213213"></a>

## Direct properties — default_https_port / 301130100323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201031030322133-3123313000123133-0221332020031301-1331322103012211-0233323230031303-3022002221102000-1331302202021233-0101111001220223"></a>

## Next pages — default_https_port / 301130100323 / 4

- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322100303321100-3033213310021230-3330000313012101-0333113303302030-0311313332003113-1321101113222121-3023221013003030-0233103011223112"></a>

## palo_alto_fw_service — palo_alto_fw_service / 003201212233 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- palo_alto_fw_service

<a id="canonical-2212323110330013-3101222023103001-2303100101011111-2301211332231032-2320133322031322-3100212302103322-2302012130013323-3313030213120310"></a>

Type: `"single"`. Computed.

Palo Alto Networks VM-Series next-generation firewall configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ami_choice": "[\"pan_ami_bundle1\",\"pan_ami_bundle2\"]",
  "x-ves-oneof-field-panaroma_connection": "[\"disable_panaroma\",\"panorama_server\"]",
  "x-ves-oneof-field-setup_options": "[\"auto_setup\",\"ssh_key\"]"
}
```

<a id="canonical-3322031301232022-0200003133130212-1223113102003333-0322300001020232-2033113002322031-2331103330101023-3031233300022100-1210310113120120"></a>

## Direct properties — palo_alto_fw_service / 003201212233 / 3

- [auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311): complete subsection reference.

- [aws_tgw_site](data-sources--nfv_service--reference--group-004.md#canonical-3033333023320323-0001020103321300-0132300322323200-3332031131310232-2330121300210232-2000130003110030-3321123030223130-0200133200122303): complete subsection reference.

- [disable_panaroma](data-sources--nfv_service--reference--group-004.md#canonical-1203222020023202-0332023100202302-3112302323202100-2031120110113020-2003012123121220-2321212122231220-3201112333213233-1120321103332100): complete subsection reference.

<a id="canonical-3130120003123322-3320211312112110-1130232133011123-2031302212033032-0001210331320200-0110120132020313-2300231321333023-1013120201330312"></a>

<a id="canonical-0303333321003003-2332321310200232-2003130222003320-2000033210000010-2132010302220333-3112300312122030-3013001332332323-3031222121331301"></a>

## instance_type property — palo_alto_fw_service / 003201212233 / 4

Type: `"string"`. Computed.

\[Enum:
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\]
&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge .. Possible values are
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\`. Defaults to
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`.

Upstream description:

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE: m5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE: m5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE: m5.12xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE: m5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE: m5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE: m5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE: m5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE: c4.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE: c4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE: c4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE: c4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE: c4.8xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE: c5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE: c5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE: c5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE: c5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE: c5.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE: c5.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE: c5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE: c5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE: c5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE: c5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE: c5n.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE: c5n.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE: r5.2xlarge.

Receipt-pinned upstream constraints:

```json
{
  "default": "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
  "enum": [
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pan_ami_bundle1](data-sources--nfv_service--reference--group-004.md#canonical-2110223121000233-3221320303312122-2000110302020133-0131111331123033-3303013000311223-3302120201222211-2220203233302110-3213113221022320): complete subsection reference.

- [pan_ami_bundle2](data-sources--nfv_service--reference--group-004.md#canonical-0123023022211213-3102321300301221-2221111210200132-2001212120303313-0301133303031123-2030312313201113-0210201002303230-0302312003233303): complete subsection reference.

- [panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200): complete subsection reference.

- [service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222): complete subsection reference.

<a id="canonical-1213322231133131-0331010312322200-0210312021033021-0231030210232123-3233001223133231-1103321221211111-0031203121211000-3233120123132012"></a>

<a id="canonical-3122223210303213-1112111303130331-2033000222131331-3110010120222012-1330322332121222-3111033233310113-1002023203212301-3333333233211230"></a>

## ssh_key property — palo_alto_fw_service / 003201212233 / 5

Type: `"string"`. Computed.

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Upstream description:

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2231312321311130-0301220130301031-3110131112331020-0310122220330231-0003132130320123-0232021132311022-1003101231203231-2133331332011100"></a>

<a id="canonical-1300003310322003-1021130110332201-0212102233311212-1030213312330223-3203013332001201-0030030120313312-0212211100001001-0100013003220020"></a>

## tags property — palo_alto_fw_service / 003201212233 / 6

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-2231020202202023-2220020331032032-1331002231230321-3222023221302020-3130033201111120-2001012300233302-2310113313122230-1201000110101032"></a>

<a id="canonical-2320122002223120-2321133132132001-0000213331031133-2320011112003233-3112331030221022-0321031010213111-3102231331321021-2102002220333000"></a>

## version property — palo_alto_fw_service / 003201212233 / 7

Type: `"string"`. Computed.

\[Enum: 11.0.0\] PAN VM-Series version. PAN-OS version. The only possible value is \`11.0.0\`.

Upstream description:

PAN-OS version.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "11.0.0"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  }
}
```

<a id="canonical-3232303011101033-3303132110002201-2100201331133133-3323330323011120-3112020200302203-2300223122032320-2300232101200122-1312332101213200"></a>

## Next pages — palo_alto_fw_service / 003201212233 / 8

- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.aws_tgw_site](data-sources--nfv_service--reference--group-004.md#canonical-3033333023320323-0001020103321300-0132300322323200-3332031131310232-2330121300210232-2000130003110030-3321123030223130-0200133200122303)
- [palo_alto_fw_service.disable_panaroma](data-sources--nfv_service--reference--group-004.md#canonical-1203222020023202-0332023100202302-3112302323202100-2031120110113020-2003012123121220-2321212122231220-3201112333213233-1120321103332100)
- [palo_alto_fw_service.pan_ami_bundle1](data-sources--nfv_service--reference--group-004.md#canonical-2110223121000233-3221320303312122-2000110302020133-0131111331123033-3303013000311223-3302120201222211-2220203233302110-3213113221022320)
- [palo_alto_fw_service.pan_ami_bundle2](data-sources--nfv_service--reference--group-004.md#canonical-0123023022211213-3102321300301221-2221111210200132-2001212120303313-0301133303031123-2030312313201113-0210201002303230-0302312003233303)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323120133301013-2300103313311301-1332103113320030-1130110033003322-0201212102131301-2221112323322201-2120333212232303-2333122001110202"></a>

## palo_alto_fw_service.auto_setup — auto_setup / 133312111303 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.auto_setup

<a id="canonical-2011011022311021-3012212223133002-0201233101202032-0133213022203213-3200021112003023-1223223003210011-0133222313233113-3212333102110333"></a>

Type: `"single"`. Computed.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ssh_keys_choice": "[\"manual_ssh_keys\"]"
}
```

<a id="canonical-1223202011100110-0332121111331003-2033010110303322-1023313330112322-2012010332033301-0102121310323203-0231330121123110-0233120130113011"></a>

## Direct properties — auto_setup / 133312111303 / 3

- [admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223): complete subsection reference.

<a id="canonical-0321301213031111-2122012313210232-1031100320200110-3023223100301103-0010332230312203-1002023101211303-0132003221312221-1022232022032132"></a>

<a id="canonical-0110322202130012-3123133013111233-0010222312320302-1320302323120332-0313201200033012-2022132023233010-2211331111313230-1103331211331031"></a>

## admin_username property — auto_setup / 133312111303 / 4

Type: `"string"`. Computed.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121): complete subsection reference.

<a id="canonical-2021111203321321-2033312303302301-0321330333003133-3311201111330001-3113331133213301-1233031200001211-2012023003200021-0010033230130000"></a>

## Next pages — auto_setup / 133312111303 / 5

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320023102021222-2201010100012231-2322203013331302-2121321011230122-3323112200000100-0203312303312201-0123202302013332-2202020111312131"></a>

## palo_alto_fw_service.auto_setup.admin_password — admin_password / 103210110102 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- palo_alto_fw_service.auto_setup.admin_password

<a id="canonical-0020233003201323-0221001312103003-1232123332133003-1202032301221022-1302202211220011-1310310203301000-3333323133313032-2122012111230330"></a>

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

<a id="canonical-0212103303201210-1101332201303213-2332311302223312-2200301033321011-2131312113330003-2022200332130203-1112103300133200-0320233112202130"></a>

## Direct properties — admin_password / 103210110102 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3311233011020021-3201330213011232-2332012201013010-1220232120302113-0303230013102200-0032002111002001-3003210300322211-2231200033220223): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3230112023323222-3000022110310023-1311232103220123-3233021210213130-0310301230001302-3101212122023013-0112033331223202-2302302013033302): complete subsection reference.

<a id="canonical-2122331230220320-1130300110000213-1203211322111223-3313302203202300-2012202023002221-3332232322321020-2223032012120303-3121002231131020"></a>

## Next pages — admin_password / 103210110102 / 4

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3311233011020021-3201330213011232-2332012201013010-1220232120302113-0303230013102200-0032002111002001-3003210300322211-2231200033220223)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3230112023323222-3000022110310023-1311232103220123-3233021210213130-0310301230001302-3101212122023013-0112033331223202-2302302013033302)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3311233011020021-3201330213011232-2332012201013010-1220232120302113-0303230013102200-0032002111002001-3003210300322211-2231200033220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003102012132130-2321102311021330-3310001130031300-0023311030212131-2110320300021213-2213332000131103-0322311112030132-2032112222311201"></a>

## palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info — blindfold_secret_info / 022330333111 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223)
- palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info

<a id="canonical-1023123103310310-3310102330023203-2331103030230201-2003313021311303-0101313032120220-2100011011322303-2111231332332333-1102113111103020"></a>

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

<a id="canonical-0010131220210112-1012010320013213-2312010002100310-0132002221320201-1333220011101113-2131211300000203-0121212303003032-0121032120321223"></a>

## Direct properties — blindfold_secret_info / 022330333111 / 3

<a id="canonical-1220031212200222-2321111001232312-1103003232023320-0133131202011220-2321022210022010-3221110302030301-2130012332220112-2321122303001332"></a>

<a id="canonical-3300333311130120-0333331003312332-1132123312023130-3122020202130221-0331202330000203-3131022122212013-2030012230333202-3000213211011321"></a>

## decryption_provider property — blindfold_secret_info / 022330333111 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2210220122110110-0123032113200212-0132131003312020-2333333022032220-3112102230001302-2332121302130301-0103113130003203-3311300202210130"></a>

<a id="canonical-0030312301021131-2011200020123031-0012323132033331-2331231211301032-3131110323210031-3123301102302133-0011213320123123-0111211102011022"></a>

## location property — blindfold_secret_info / 022330333111 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2111033033022230-2113120103311230-0121100111012211-0001123023210031-0322112221220221-0131312032303112-3233100220310111-1121330320323133"></a>

<a id="canonical-1200202300102200-1021021321301322-2312320203313311-0100133220333233-2221132033203010-0020200030320200-2103331100313222-3202101210120230"></a>

## store_provider property — blindfold_secret_info / 022330333111 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0123001311201331-3232000101332111-3133111211110103-2300131123321301-1200020221221010-0001223021213300-0020000133123310-2220211101011213"></a>

## Next pages — blindfold_secret_info / 022330333111 / 7

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3230112023323222-3000022110310023-1311232103220123-3233021210213130-0310301230001302-3101212122023013-0112033331223202-2302302013033302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133210003302101-2021300121021211-3301123323102112-2323322333302222-2111002013122201-0330200230221131-1102023232213022-0101230003133213"></a>

## palo_alto_fw_service.auto_setup.admin_password.clear_secret_info — clear_secret_info / 130121330122 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223)
- palo_alto_fw_service.auto_setup.admin_password.clear_secret_info

<a id="canonical-3131022021011233-1103132133133231-0233232102303312-3213223223010013-1102133330023213-3221132332133020-0323122330222331-2212031133131011"></a>

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

<a id="canonical-1323011323023131-3123213001031022-2121313120010131-3010131310112313-0223122033101131-1123121300000122-0031221333030101-3102111332301032"></a>

## Direct properties — clear_secret_info / 130121330122 / 3

<a id="canonical-1200201212101203-1100333121212101-3022212012032230-2012131211212231-2311210331001030-2132001001003130-0132301003212112-1232123202321211"></a>

<a id="canonical-1233201011100212-0122112311332222-3313301223232111-0010213321301121-3313221332110123-0330331010202211-2310301311302101-0210012211320212"></a>

## provider_ref property — clear_secret_info / 130121330122 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3232013331012300-0232221033113323-0332232322300302-1032232122320113-0210323221022311-3322000002332333-2320011103123010-1022310311112010"></a>

<a id="canonical-0201312113013131-1312101102120221-2232200133000013-1132231300221120-0201012021303230-2211203101221021-0001023331201022-1201110323230030"></a>

## URL property — clear_secret_info / 130121330122 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1212113330322133-3011202212132031-3312200022000210-2323322010020111-1103121230323311-1122001102133122-0230120331130112-0331033310020212"></a>

## Next pages — clear_secret_info / 130121330122 / 6

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232211303133101-3303000030033211-2202000212110013-3012323211132223-1013032013203301-1000301211121131-3000213222131213-1010131202210333"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys — manual_ssh_keys / 303323200132 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="canonical-2031010030111102-2301013001331211-0320133301320000-3222112031320120-0013332010023210-0121210123321000-3103002222232231-0000130112133013"></a>

Type: `"single"`. Computed.

SSH Key includes both public and private key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210223231230232-2112101311012211-2103003023310230-0101331012113102-3200010132102322-2210231331323013-1232330100223331-3212021131132103"></a>

## Direct properties — manual_ssh_keys / 303323200132 / 3

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131): complete subsection reference.

<a id="canonical-3313303030201130-3002023313231213-1103122331213121-3300133203130131-1301112320031300-2303221220012213-2003203232010121-0300302103203201"></a>

<a id="canonical-2110100033122021-2011132210203131-1001111223023121-3200132220002221-1202232122330112-0213120231330300-3132322300012021-2131202221002221"></a>

## public_key property — manual_ssh_keys / 303323200132 / 4

Type: `"string"`. Computed.

Authorized Public SSH key which will be programmed on the node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----\\n.*\\n-----END (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2002211232302103-2333221220120203-3322032231201313-3002202233223113-0200122103003021-2220110200110133-3120333020123030-3101100313333023"></a>

## Next pages — manual_ssh_keys / 303323200132 / 5

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331333132132211-2302112113011022-1231130232023221-3212003213330033-0203320313000123-2313222031012232-3330002121121323-1232213331112331"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key — private_key / 011213022110 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key

<a id="canonical-2130112210210233-0210012001001312-2121021123301120-2031031113112221-3003031101212022-0220302223230133-3223332310233132-1303231021210000"></a>

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

<a id="canonical-1011213113222321-0200333200212001-1202112233013101-3132101332211123-3331203310332222-3013213302321320-0130221301222223-3113332130031231"></a>

## Direct properties — private_key / 011213022110 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-2002003200130322-3223321312321232-3113023310323022-2012131022123111-0022311101201101-2312033012223223-1213302232112300-0201320233032120): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1102110202003220-1120303130223002-1021013031233113-1132102330322333-0032331121300030-3030020011000300-0033301323103302-1223110212033221): complete subsection reference.
