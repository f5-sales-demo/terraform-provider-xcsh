---
page_title: "xcsh_user_identification reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification reference."
---

# xcsh_user_identification reference

<a id="canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- Property reference

<a id="canonical-1301233323211233-3021133310323003-0001021332112022-1233203030000330-3101330210211031-0300311303323033-3110000200321002-3202333033001323"></a>

### Direct properties for `xcsh_user_identification`

<a id="canonical-1322133101032122-2130332213031233-3122112232322221-2333323111213300-0313231110233303-1030122220011323-0201130002212120-3030333113011201"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

<a id="canonical-2313033000221113-1231202321330311-1122133133310330-3330010323100332-1010023301012002-3110303020113100-2301001220121022-0033210311101120"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the UserIdentification.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3012031133011012-0213203231021013-2330330222020233-3020231002102103-1303213202203330-3210201230213313-3020220111122332-1110020100333111"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1210213113310023-0113322230012100-0233123232033311-1032003020333211-2221131002002102-3300131031221331-2321131211030200-3033311202321233"></a>

<a id="canonical-3133233211002000-1223231303213202-2331232323033002-2332123330121031-0123311303232330-3200213121300312-0103003220222221-2100211112112112"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

<a id="canonical-2223020200003233-0130302323330221-2312221032102133-3212113100212012-2133020322120000-1130322232033233-0121222003010002-2322322021330231"></a>

#### `name` property

Type: `"string"`. Required.

Name of the UserIdentification.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0232012332201022-3022011310330101-3110333302003133-0001031323220012-0332303313321112-0320332002131232-0033320100003022-3100213122322233"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the UserIdentification exists.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2233233213223131-1121011330003320-2203303000231232-2030122020112323-1013011011201233-2000012120330101-0110132020232232-3033011123113331"></a>

### All schema paths for `xcsh_user_identification`

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

<a id="canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- rules

<a id="canonical-1300203312202103-2131020133021121-2120203232210211-0213111101021203-0102001021002121-1022222222123311-0302302021202333-3332303202221023"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3203323013321111-1233220221230100-1102231322030222-1133213121221333-0311022312223102-1213031002021313-1313103220113030-1221013330223301"></a>

### Direct properties for `rules`

- [client_asn](data-sources--user_identification--reference--group-001.md#canonical-2211301013311301-2211231001201130-0233313313010001-3203233201320132-2011102013302302-2313000010221331-1012332332212100-0212313001310222): complete subsection reference.

- [client_city](data-sources--user_identification--reference--group-001.md#canonical-0012131122003213-3031230220132323-3122132211021220-0122121212112201-1031011030010033-1212301010030133-1312221133323333-0222003300200223): complete subsection reference.

- [client_country](data-sources--user_identification--reference--group-001.md#canonical-1221202030311211-1033023013211233-0311231233132301-3100112000023022-2133031010303003-1133210213213120-0120101211300223-1011020033312221): complete subsection reference.

- [client_ip](data-sources--user_identification--reference--group-001.md#canonical-3111322221000120-1310110221322323-0311103020200202-1031032003210212-0321232301202302-0032133032011221-3320112303313213-2300032200131123): complete subsection reference.

- [client_region](data-sources--user_identification--reference--group-001.md#canonical-0333121032300332-3131002210130033-3001123000003111-3333003122031123-1332223031021112-3321203213311303-3121111000230301-3101121310232021): complete subsection reference.

<a id="canonical-0300103002302010-3020313100312122-1230311312000313-3300312231310212-2102102003132300-3211112301110321-0013323022111131-1203202202133310"></a>

<a id="canonical-2223202330011221-3200030320113120-2030211001013031-1221323331000232-3202231323002112-2313010312030123-2131121011032011-3330200300312311"></a>

#### `rules.cookie_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1332223132231122-0222031213313231-0302132332321133-1033121203113011-3122100220103300-1223013323010110-0203300101132221-2022001300022120"></a>

#### `rules.http_header_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3032123220333023-0122031103212032-0020023213133123-0102220311323211-3303223131203112-2222231123313300-1130130311323231-3102010321132200"></a>

#### `rules.ip_and_http_header_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2320030310333111-2000121122133101-2132013200223213-3303033131011003-2322013023021200-1312022011310230-3201113102320310-3011220330222001"></a>

#### `rules.jwt_claim_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3122022100133111-0223001120202022-3321330233030031-0003331320302310-2230122223021022-3320032100120320-0222002211003311-3022022202311233"></a>

#### `rules.query_param_key` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2211301013311301-2211231001201130-0233313313010001-3203233201320132-2011102013302302-2313000010221331-1012332332212100-0212313001310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_asn` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_asn

<a id="canonical-1102100311200013-2133022310231013-3121112103313213-3031202030102123-0221112022101020-3100311030002120-0200031101220113-3220120202111202"></a>

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

<a id="canonical-0012131122003213-3031230220132323-3122132211021220-0122121212112201-1031011030010033-1212301010030133-1312221133323333-0222003300200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_city` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_city

<a id="canonical-1210312222322313-0233300233032210-3001110000130302-2310033232130323-2322102000133001-3100200122112131-2202331330112123-0032322101201233"></a>

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

<a id="canonical-1221202030311211-1033023013211233-0311231233132301-3100112000023022-2133031010303003-1133210213213120-0120101211300223-1011020033312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_country` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_country

<a id="canonical-3221231333033303-2112201123230121-1000131020101103-0002123033320011-1111201021233223-1123303202023331-3322332323111100-3102300000331021"></a>

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

<a id="canonical-3111322221000120-1310110221322323-0311103020200202-1031032003210212-0321232301202302-0032133032011221-3320112303313213-2300032200131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_ip` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_ip

<a id="canonical-2120230013331123-2132011212300320-1033102012032113-2001221321023323-2311031311213213-2032302000332013-0310222121302113-0031003302111131"></a>

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

<a id="canonical-0333121032300332-3131002210130033-3001123000003111-3333003122031123-1332223031021112-3321203213311303-3121111000230301-3101121310232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_region` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.client_region

<a id="canonical-0103121020130301-1211123203310001-1311221010303300-2032102011103321-2212301333221133-0020133222120303-0003230012002222-0101212133112002"></a>

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

<a id="canonical-2320303013222332-3023202333130120-2011202212121031-3030110222232212-1011222021222313-1011300022120123-3133333112213213-3332112120303220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ip_and_ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ip_and_ja4_tls_fingerprint

<a id="canonical-3232002210320011-3300321110000230-0001231331132002-2130220203300031-0311211302220003-3120332300320012-2301322103001101-3210022131100311"></a>

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

<a id="canonical-1130330220233023-2022313333120113-3131003111020231-1001113322101212-0111331310100011-1030102213202123-2020101020101123-0123311013123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ip_and_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ip_and_tls_fingerprint

<a id="canonical-1132201322222321-3223010210211323-0303023321322311-1223100313021322-0331222121312121-0221210303132233-1030013102121033-0112230332320211"></a>

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

<a id="canonical-1333132302231133-1012010222321122-3100100122132320-2333322113103220-3113003000022031-2300110231011113-2222222311221022-3001230000222003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.ja4_tls_fingerprint

<a id="canonical-1330000131303001-1310322002222221-2122303312023132-2133323033102022-3231321201303211-3332300012133202-2300213203210010-2121002000030221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ja4 tls fingerprint.

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

<a id="canonical-0331101123301300-0213110012321311-0312102201213001-1323100301231310-0313331130310122-2330001001022200-0111131122322100-3313033121220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.none` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.none

<a id="canonical-1312312330001211-0212303120013200-0223310110221031-2211133112200223-0201301003210122-3210003132011200-0200133223012032-1113033002011202"></a>

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

<a id="canonical-3223201212222123-1320213332302222-3033301310213001-0232101010110322-1130011101033201-2231133010313101-0010013203310202-0332121310312003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233)
- rules.tls_fingerprint

<a id="canonical-1103021113222033-3123223220203112-2032322121203120-3300233333300231-2203311331103023-2201212001122221-3312120321333213-3010013322320311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls fingerprint.

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
