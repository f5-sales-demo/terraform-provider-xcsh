---
page_title: "xcsh_user_identification reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification reference."
---

# xcsh_user_identification reference

<a id="canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301233323211233-3021133310323003-0001021332112022-1233203030000330-3101330210211031-0300311303323033-3110000200321002-3202333033001323"></a>

## Property reference — Property reference / 033000122132 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- Property reference

<a id="canonical-2313033000221113-1231202321330311-1122133133310330-3330010323100332-1010023301012002-3110303020113100-2301001220121022-0033210311101120"></a>

## Direct properties — Property reference / 033000122132 / 3

<a id="canonical-1322133101032122-2130332213031233-3122112232322221-2333323111213300-0313231110233303-1030122220011323-0201130002212120-3030333113011201"></a>

<a id="canonical-3012031133011012-0213203231021013-2330330222020233-3020231002102103-1303213202203330-3210201230213313-3020220111122332-1110020100333111"></a>

## annotations property — Property reference / 033000122132 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3313132030200100-3032103011033200-1102303133001102-1000110103103201-1231313223013231-2323131122131122-2231320331322223-1322330331200212"></a>

<a id="canonical-3133233211002000-1223231303213202-2331232323033002-2332123330121031-0123311303232330-3200213121300312-0103003220222221-2100211112112112"></a>

## description property — Property reference / 033000122132 / 5

Type: `"string"`. Computed.

Description of the UserIdentification.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-2212100232220301-1301200330230223-1232202330233222-1331230303111020-3101311103130313-2202223302010101-3312132331131212-0302131100200210"></a>

<a id="canonical-2223020200003233-0130302323330221-2312221032102133-3212113100212012-2133020322120000-1130322232033233-0121222003010002-2322322021330231"></a>

## ID property — Property reference / 033000122132 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1210213113310023-0113322230012100-0233123232033311-1032003020333211-2221131002002102-3300131031221331-2321131211030200-3033311202321233"></a>

<a id="canonical-0232012332201022-3022011310330101-3110333302003133-0001031323220012-0332303313321112-0320332002131232-0033320100003022-3100213122322233"></a>

## labels property — Property reference / 033000122132 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102021213210231-3201220333011133-0133110031130202-0002113100101112-1320123320331110-3102233102023121-0131002103330212-3031202301231103"></a>

<a id="canonical-2233233213223131-1121011330003320-2203303000231232-2030122020112323-1013011011201233-2000012120330101-0110132020232232-3033011123113331"></a>

## name property — Property reference / 033000122132 / 8

Type: `"string"`. Required.

Name of the UserIdentification.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0221110310311033-3212130321012300-0210022031223021-1222320323313313-1201030001331202-3121130220210100-0310322031330303-0112211333011203"></a>

<a id="canonical-3313310100123323-3010013222100112-3212232100121222-0003022310322023-3301223020120233-0010011233100333-0301010101300212-2303311301032210"></a>

## namespace property — Property reference / 033000122132 / 9

Type: `"string"`. Required.

Namespace where the UserIdentification exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233): complete subsection reference.

<a id="canonical-2111300130131322-0022002010332212-0220103200203321-1103002032102131-0220133123013300-2032322023131210-0013203322323132-1022200000022321"></a>

## All schema paths — Property reference / 033000122132 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--user_identification--reference--group-001.md#canonical-1322133101032122-2130332213031233-3122112232322221-2333323111213300-0313231110233303-1030122220011323-0201130002212120-3030333113011201) |
| `description` | [description](data-sources--user_identification--reference--group-001.md#canonical-3313132030200100-3032103011033200-1102303133001102-1000110103103201-1231313223013231-2323131122131122-2231320331322223-1322330331200212) |
| `id` | [ID](data-sources--user_identification--reference--group-001.md#canonical-2212100232220301-1301200330230223-1232202330233222-1331230303111020-3101311103130313-2202223302010101-3312132331131212-0302131100200210) |
| `labels` | [labels](data-sources--user_identification--reference--group-001.md#canonical-1210213113310023-0113322230012100-0233123232033311-1032003020333211-2221131002002102-3300131031221331-2321131211030200-3033311202321233) |
| `name` | [name](data-sources--user_identification--reference--group-001.md#canonical-3102021213210231-3201220333011133-0133110031130202-0002113100101112-1320123320331110-3102233102023121-0131002103330212-3031202301231103) |
| `namespace` | [namespace](data-sources--user_identification--reference--group-001.md#canonical-0221110310311033-3212130321012300-0210022031223021-1222320323313313-1201030001331202-3121130220210100-0310322031330303-0112211333011203) |
| `rules` | [rules](data-sources--user_identification--reference--group-001.md#canonical-1300203312202103-2131020133021121-2120203232210211-0213111101021203-0102001021002121-1022222222123311-0302302021202333-3332303202221023) |
| `rules.client_asn` | [rules.client_asn](data-sources--user_identification--reference--group-001.md#canonical-1102100311200013-2133022310231013-3121112103313213-3031202030102123-0221112022101020-3100311030002120-0200031101220113-3220120202111202) |
| `rules.client_city` | [rules.client_city](data-sources--user_identification--reference--group-001.md#canonical-1210312222322313-0233300233032210-3001110000130302-2310033232130323-2322102000133001-3100200122112131-2202331330112123-0032322101201233) |
| `rules.client_country` | [rules.client_country](data-sources--user_identification--reference--group-001.md#canonical-3221231333033303-2112201123230121-1000131020101103-0002123033320011-1111201021233223-1123303202023331-3322332323111100-3102300000331021) |
| `rules.client_ip` | [rules.client_ip](data-sources--user_identification--reference--group-001.md#canonical-2120230013331123-2132011212300320-1033102012032113-2001221321023323-2311031311213213-2032302000332013-0310222121302113-0031003302111131) |
| `rules.client_region` | [rules.client_region](data-sources--user_identification--reference--group-001.md#canonical-0103121020130301-1211123203310001-1311221010303300-2032102011103321-2212301333221133-0020133222120303-0003230012002222-0101212133112002) |
| `rules.cookie_name` | [rules.cookie_name](data-sources--user_identification--reference--group-001.md#canonical-0300103002302010-3020313100312122-1230311312000313-3300312231310212-2102102003132300-3211112301110321-0013323022111131-1203202202133310) |
| `rules.http_header_name` | [rules.http_header_name](data-sources--user_identification--reference--group-001.md#canonical-3033310000103001-1232311230313301-2133003331031131-1300132010101303-2233030010121101-3223033322133102-3320111111021111-2323301102111302) |
| `rules.ip_and_http_header_name` | [rules.ip_and_http_header_name](data-sources--user_identification--reference--group-001.md#canonical-3131202031210102-0220222321233322-0330330103113132-3331022221212202-3000222310031121-1312101030130223-3313112111211221-0333100313312212) |
| `rules.ip_and_ja4_tls_fingerprint` | [rules.ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-3232002210320011-3300321110000230-0001231331132002-2130220203300031-0311211302220003-3120332300320012-2301322103001101-3210022131100311) |
| `rules.ip_and_tls_fingerprint` | [rules.ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1132201322222321-3223010210211323-0303023321322311-1223100313021322-0331222121312121-0221210303132233-1030013102121033-0112230332320211) |
| `rules.ja4_tls_fingerprint` | [rules.ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1330000131303001-1310322002222221-2122303312023132-2133323033102022-3231321201303211-3332300012133202-2300213203210010-2121002000030221) |
| `rules.jwt_claim_name` | [rules.jwt_claim_name](data-sources--user_identification--reference--group-001.md#canonical-3220232311010030-1303133302011323-1111002001112022-2331113230222211-1113103132231223-0220112223323011-1303211302220003-0020211320212333) |
| `rules.none` | [rules.none](data-sources--user_identification--reference--group-001.md#canonical-1312312330001211-0212303120013200-0223310110221031-2211133112200223-0201301003210122-3210003132011200-0200133223012032-1113033002011202) |
| `rules.query_param_key` | [rules.query_param_key](data-sources--user_identification--reference--group-001.md#canonical-0332221201102220-3002000332201201-0100033201332302-1303233321330033-1023010332323101-3130122122112212-2310333233311331-0031031032332013) |
| `rules.tls_fingerprint` | [rules.tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1103021113222033-3123223220203112-2032322121203120-3300233333300231-2203311331103023-2201212001122221-3312120321333213-3010013322320311) |

<a id="canonical-2221112012120012-3013223322123000-1130221333330202-3120211331303131-0223113030031130-2110132011201103-3132132201203222-0031121222133210"></a>

## Next pages — Property reference / 033000122132 / 11

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203323013321111-1233220221230100-1102231322030222-1133213121221333-0311022312223102-1213031002021313-1313103220113030-1221013330223301"></a>

## rules — rules / 012010002320 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- rules

<a id="canonical-1300203312202103-2131020133021121-2120203232210211-0213111101021203-0102001021002121-1022222222123311-0302302021202333-3332303202221023"></a>

Type: `"list"`. Computed.

Ordered list of rules that are evaluated sequentially against the input fields extracted from an API
request in order to determine a user identifier. Evaluation of the rules is terminated once a user
identifier has been extracted.

Upstream description:

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

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

<a id="canonical-2223202330011221-3200030320113120-2030211001013031-1221323331000232-3202231323002112-2313010312030123-2131121011032011-3330200300312311"></a>

## Direct properties — rules / 012010002320 / 3

- [client_asn](data-sources--user_identification--reference--group-001.md#canonical-2211301013311301-2211231001201130-0233313313010001-3203233201320132-2011102013302302-2313000010221331-1012332332212100-0212313001310222): complete subsection reference.

- [client_city](data-sources--user_identification--reference--group-001.md#canonical-0012131122003213-3031230220132323-3122132211021220-0122121212112201-1031011030010033-1212301010030133-1312221133323333-0222003300200223): complete subsection reference.

- [client_country](data-sources--user_identification--reference--group-001.md#canonical-1221202030311211-1033023013211233-0311231233132301-3100112000023022-2133031010303003-1133210213213120-0120101211300223-1011020033312221): complete subsection reference.

- [client_ip](data-sources--user_identification--reference--group-001.md#canonical-3111322221000120-1310110221322323-0311103020200202-1031032003210212-0321232301202302-0032133032011221-3320112303313213-2300032200131123): complete subsection reference.

- [client_region](data-sources--user_identification--reference--group-001.md#canonical-0333121032300332-3131002210130033-3001123000003111-3333003122031123-1332223031021112-3321203213311303-3121111000230301-3101121310232021): complete subsection reference.

<a id="canonical-0300103002302010-3020313100312122-1230311312000313-3300312231310212-2102102003132300-3211112301110321-0013323022111131-1203202202133310"></a>

<a id="canonical-1332223132231122-0222031213313231-0302132332321133-1033121203113011-3122100220103300-1223013323010110-0203300101132221-2022001300022120"></a>

## cookie_name property — rules / 012010002320 / 4

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3033310000103001-1232311230313301-2133003331031131-1300132010101303-2233030010121101-3223033322133102-3320111111021111-2323301102111302"></a>

<a id="canonical-3032123220333023-0122031103212032-0020023213133123-0102220311323211-3303223131203112-2222231123313300-1130130311323231-3102010321132200"></a>

## http_header_name property — rules / 012010002320 / 5

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3131202031210102-0220222321233322-0330330103113132-3331022221212202-3000222310031121-1312101030130223-3313112111211221-0333100313312212"></a>

<a id="canonical-2320030310333111-2000121122133101-2132013200223213-3303033131011003-2322013023021200-1312022011310230-3201113102320310-3011220330222001"></a>

## ip_and_http_header_name property — rules / 012010002320 / 6

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-2320303013222332-3023202333130120-2011202212121031-3030110222232212-1011222021222313-1011300022120123-3133333112213213-3332112120303220): complete subsection reference.

- [ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1130330220233023-2022313333120113-3131003111020231-1001113322101212-0111331310100011-1030102213202123-2020101020101123-0123311013123000): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1333132302231133-1012010222321122-3100100122132320-2333322113103220-3113003000022031-2300110231011113-2222222311221022-3001230000222003): complete subsection reference.

<a id="canonical-3220232311010030-1303133302011323-1111002001112022-2331113230222211-1113103132231223-0220112223323011-1303211302220003-0020211320212333"></a>

<a id="canonical-3122022100133111-0223001120202022-3321330233030031-0003331320302310-2230122223021022-3320032100120320-0222002211003311-3022022202311233"></a>

## jwt_claim_name property — rules / 012010002320 / 7

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [none](data-sources--user_identification--reference--group-001.md#canonical-0331101123301300-0213110012321311-0312102201213001-1323100301231310-0313331130310122-2330001001022200-0111131122322100-3313033121220222): complete subsection reference.

<a id="canonical-0332221201102220-3002000332201201-0100033201332302-1303233321330033-1023010332323101-3130122122112212-2310333233311331-0031031032332013"></a>

<a id="canonical-2031202220020222-0022232220323311-2023022112111203-2222213020133222-3013033123033110-0113212331231333-2033333012330031-1102001303110203"></a>

## query_param_key property — rules / 012010002320 / 8

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-3223201212222123-1320213332302222-3033301310213001-0232101010110322-1130011101033201-2231133010313101-0010013203310202-0332121310312003): complete subsection reference.

<a id="canonical-1230012220200300-0010103112023331-0020301223333201-1302021223333301-3113110201210230-1310021033031213-1013001302012120-1230022101021100"></a>

## Next pages — rules / 012010002320 / 9

- [rules.client_asn](data-sources--user_identification--reference--group-001.md#canonical-2211301013311301-2211231001201130-0233313313010001-3203233201320132-2011102013302302-2313000010221331-1012332332212100-0212313001310222)
- [rules.client_city](data-sources--user_identification--reference--group-001.md#canonical-0012131122003213-3031230220132323-3122132211021220-0122121212112201-1031011030010033-1212301010030133-1312221133323333-0222003300200223)
- [rules.client_country](data-sources--user_identification--reference--group-001.md#canonical-1221202030311211-1033023013211233-0311231233132301-3100112000023022-2133031010303003-1133210213213120-0120101211300223-1011020033312221)
- [rules.client_ip](data-sources--user_identification--reference--group-001.md#canonical-3111322221000120-1310110221322323-0311103020200202-1031032003210212-0321232301202302-0032133032011221-3320112303313213-2300032200131123)
- [rules.client_region](data-sources--user_identification--reference--group-001.md#canonical-0333121032300332-3131002210130033-3001123000003111-3333003122031123-1332223031021112-3321203213311303-3121111000230301-3101121310232021)
- [rules.ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-2320303013222332-3023202333130120-2011202212121031-3030110222232212-1011222021222313-1011300022120123-3133333112213213-3332112120303220)
- [rules.ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1130330220233023-2022313333120113-3131003111020231-1001113322101212-0111331310100011-1030102213202123-2020101020101123-0123311013123000)
- [rules.ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-1333132302231133-1012010222321122-3100100122132320-2333322113103220-3113003000022031-2300110231011113-2222222311221022-3001230000222003)
- [rules.none](data-sources--user_identification--reference--group-001.md#canonical-0331101123301300-0213110012321311-0312102201213001-1323100301231310-0313331130310122-2330001001022200-0111131122322100-3313033121220222)
- [rules.tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-3223201212222123-1320213332302222-3033301310213001-0232101010110322-1130011101033201-2231133010313101-0010013203310202-0332121310312003)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-2211301013311301-2211231001201130-0233313313010001-3203233201320132-2011102013302302-2313000010221331-1012332332212100-0212313001310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111211320213131-0310013102120012-3001230332210022-0102220313332121-2332321000201313-1311313233002300-2111311132331312-0322023133121231"></a>

## rules.client_asn — client_asn / 011103102033 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_asn

<a id="canonical-1102100311200013-2133022310231013-3121112103313213-3031202030102123-0221112022101020-3100311030002120-0200031101220113-3220120202111202"></a>

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

<a id="canonical-0022230031200202-3301003131121212-2322220320320010-0110223233211232-3030113030302012-2203011012010303-0202203001212031-3212100302130302"></a>

## Direct properties — client_asn / 011103102033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323232033020020-1313200332220100-2023220221212103-3020301310302300-1201103301120102-3120321213212131-3333020312022023-3303330303101132"></a>

## Next pages — client_asn / 011103102033 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-0012131122003213-3031230220132323-3122132211021220-0122121212112201-1031011030010033-1212301010030133-1312221133323333-0222003300200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313322022003002-1122031002213230-2301012012102032-0032100211122233-3033230200323101-0212121033112301-3000321022301032-1220223221333113"></a>

## rules.client_city — client_city / 331103112222 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_city

<a id="canonical-1210312222322313-0233300233032210-3001110000130302-2310033232130323-2322102000133001-3100200122112131-2202331330112123-0032322101201233"></a>

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

<a id="canonical-3020300032011301-3322311332110102-0012111331302121-0333331202222121-3210121222320011-2103300033202130-2201131201313321-3013002123131003"></a>

## Direct properties — client_city / 331103112222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110303123010332-2130312122131311-3123220132212331-0120033212101122-0020233322312302-2133331012202320-2312313110320323-0320112111322311"></a>

## Next pages — client_city / 331103112222 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-1221202030311211-1033023013211233-0311231233132301-3100112000023022-2133031010303003-1133210213213120-0120101211300223-1011020033312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101023002313122-3303011211000021-0110122133021023-1022201021021311-3200330011222031-0331121001233100-3310322311133000-3120021000211323"></a>

## rules.client_country — client_country / 202312212123 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_country

<a id="canonical-3221231333033303-2112201123230121-1000131020101103-0002123033320011-1111201021233223-1123303202023331-3322332323111100-3102300000331021"></a>

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

<a id="canonical-2012131202233010-2111323330110101-1131303012310033-2130303223232321-2212000010332331-1123330230011010-1120311122013332-3221102013320022"></a>

## Direct properties — client_country / 202312212123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310030100113031-1212001310103132-0121133321313001-0221001330101221-2130113233120002-2012313323203201-3012200032020003-3320102030231122"></a>

## Next pages — client_country / 202312212123 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-3111322221000120-1310110221322323-0311103020200202-1031032003210212-0321232301202302-0032133032011221-3320112303313213-2300032200131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310121133122110-2103202221213332-1212210023200323-1320223332220111-3002221111323210-0101303222223032-2031333021321301-2222200300320130"></a>

## rules.client_ip — client_ip / 233303203331 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_ip

<a id="canonical-2120230013331123-2132011212300320-1033102012032113-2001221321023323-2311031311213213-2032302000332013-0310222121302113-0031003302111131"></a>

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

<a id="canonical-3001033022012322-1112323100030201-1012110002231033-3012212021220313-3302210112023000-1333311120002111-0022330122133221-1313210100330322"></a>

## Direct properties — client_ip / 233303203331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003002321321202-3031221013202123-2320111131200233-3002012213033021-0011012313321023-1321213122211300-0330122301031021-3011121202031232"></a>

## Next pages — client_ip / 233303203331 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-0333121032300332-3131002210130033-3001123000003111-3333003122031123-1332223031021112-3321203213311303-3121111000230301-3101121310232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000120122333201-2100212000333133-1132212130332321-3111200002132033-2231120220322303-1233200113220123-0323112112031200-0103013100312231"></a>

## rules.client_region — client_region / 233211313030 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_region

<a id="canonical-0103121020130301-1211123203310001-1311221010303300-2032102011103321-2212301333221133-0020133222120303-0003230012002222-0101212133112002"></a>

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

<a id="canonical-1030112331023330-3313032013301013-3232230132001010-2333002013301013-2121013212310323-3210221313011010-2300222220030203-0331303221112022"></a>

## Direct properties — client_region / 233211313030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123102312033000-3111100122300332-3000122010001230-1131212110322001-2302303322202321-2013102001023131-0102030323023110-0022020011131031"></a>

## Next pages — client_region / 233211313030 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-2320303013222332-3023202333130120-2011202212121031-3030110222232212-1011222021222313-1011300022120123-3133333112213213-3332112120303220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002033221020233-1112202000322123-0230201223212310-3110131121111133-0331012030233200-2323012003032130-0000330133303023-1111312101122132"></a>

## rules.ip_and_ja4_tls_fingerprint — ip_and_ja4_tls_fingerprint / 030310031123 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ip_and_ja4_tls_fingerprint

<a id="canonical-3232002210320011-3300321110000230-0001231331132002-2130220203300031-0311211302220003-3120332300320012-2301322103001101-3210022131100311"></a>

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

<a id="canonical-1001020121212101-1031131013111200-3022122112313331-2222101100330310-3021020201123231-3331022320300232-0030020010012131-3020032020020303"></a>

## Direct properties — ip_and_ja4_tls_fingerprint / 030310031123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320020111101010-0003010311302111-1320312303020033-2322300220201310-0330232210210113-3200002130000102-1011030212003321-1122123222020130"></a>

## Next pages — ip_and_ja4_tls_fingerprint / 030310031123 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-1130330220233023-2022313333120113-3131003111020231-1001113322101212-0111331310100011-1030102213202123-2020101020101123-0123311013123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120210230000323-1322032212332313-2000112023130312-3113100033102021-0121113013031032-2321002323312010-1222332301111031-1322301011331311"></a>

## rules.ip_and_tls_fingerprint — ip_and_tls_fingerprint / 121301213123 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ip_and_tls_fingerprint

<a id="canonical-1132201322222321-3223010210211323-0303023321322311-1223100313021322-0331222121312121-0221210303132233-1030013102121033-0112230332320211"></a>

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

<a id="canonical-1121013221102320-3032112020320203-3320101112030030-2130320302211030-2031120122112103-1232023122201123-3012233112311111-2020013102303033"></a>

## Direct properties — ip_and_tls_fingerprint / 121301213123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222100210203100-2020212130132030-1203222020121020-3203130322022122-2022312033021313-3120101101032101-1231211110331100-2131202322211020"></a>

## Next pages — ip_and_tls_fingerprint / 121301213123 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-1333132302231133-1012010222321122-3100100122132320-2333322113103220-3113003000022031-2300110231011113-2222222311221022-3001230000222003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212031130020310-2200312031311202-2333121233130231-2333233203102331-3011013011332011-0230323033320111-1023232010031211-1233102210313111"></a>

## rules.ja4_tls_fingerprint — ja4_tls_fingerprint / 323010112210 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ja4_tls_fingerprint

<a id="canonical-1330000131303001-1310322002222221-2122303312023132-2133323033102022-3231321201303211-3332300012133202-2300213203210010-2121002000030221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ja4 tls fingerprint.

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

<a id="canonical-1333103220333033-1012103323133221-1112301211030311-1200301311021023-3221002120022320-0210302331033322-0322031111102030-1203110023101123"></a>

## Direct properties — ja4_tls_fingerprint / 323010112210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232200223131130-2200200031030221-2303302112302220-0231233321210212-1130112202300102-0122332310103323-2311023302200321-2133022002203002"></a>

## Next pages — ja4_tls_fingerprint / 323010112210 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-0331101123301300-0213110012321311-0312102201213001-1323100301231310-0313331130310122-2330001001022200-0111131122322100-3313033121220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313201010103101-0010103120303020-2202122111003000-1101310123221110-1221011220103011-2123322203211120-1110112302320031-0323213212131131"></a>

## rules.none — none / 023122002011 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.none

<a id="canonical-1312312330001211-0212303120013200-0223310110221031-2211133112200223-0201301003210122-3210003132011200-0200133223012032-1113033002011202"></a>

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

<a id="canonical-1130313033031130-3231122032002332-0112220213100013-2013002011332202-1211331320210311-0013200112322021-0221121310002200-0130121132123033"></a>

## Direct properties — none / 023122002011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332231323010130-2330303131301223-3133220012121220-2331132020233003-3311010211312123-3223020033020232-3231001032013011-3231111021323311"></a>

## Next pages — none / 023122002011 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)

<a id="canonical-3223201212222123-1320213332302222-3033301310213001-0232101010110322-1130011101033201-2231133010313101-0010013203310202-0332121310312003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311220022331320-2121021002001020-1011322312133011-2022132233200010-0321211202110203-0113101211311211-2032203101210231-3212020030000032"></a>

## rules.tls_fingerprint — tls_fingerprint / 313120200112 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.tls_fingerprint

<a id="canonical-1103021113222033-3123223220203112-2032322121203120-3300233333300231-2203311331103023-2201212001122221-3312120321333213-3010013322320311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls fingerprint.

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

<a id="canonical-3022111023333223-3312013322310221-3123212112312201-2322001230123311-1121032332311012-1130133010100022-2230212112322131-3333213211221100"></a>

## Direct properties — tls_fingerprint / 313120200112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202113133231201-1212110120131310-3010013210103123-3011011222221030-2233323233022232-2113302110120323-2111020221100302-3333131220130220"></a>

## Next pages — tls_fingerprint / 313120200112 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
