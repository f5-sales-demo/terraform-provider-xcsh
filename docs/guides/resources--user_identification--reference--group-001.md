---
page_title: "xcsh_user_identification reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification reference."
---

# xcsh_user_identification reference

<a id="canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- Property reference

<a id="canonical-3001021212130013-1223312200203232-1331002130312311-2032320210133302-2123320210303102-1110113101032303-2223302023320231-0001032230120010"></a>

### Direct properties for `xcsh_user_identification`

<a id="canonical-2213312132012231-1203022113312313-0103323222122230-3202321213302210-3031201101021110-3213011201111132-0300331120322030-3222312030210330"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-1302320333320221-1312112322023131-0222221331000301-0313331101113031-2203133220032313-0223210121232132-0330000011113323-0011221331000130"></a>

<a id="canonical-0031010010220103-2321003232122322-3010123202210131-0003111203321320-2110202111000123-3111211130223110-0311230101023110-0321231312231103"></a>

#### `description` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1223330110200303-2032231011000031-0000111311023331-0233031000301002-0201122213132212-2010223122013033-0312210200312030-1333110111033030"></a>

<a id="canonical-3003322231331232-2232310000330111-3120032101020202-3031230201113110-2110213310231010-2220112132132201-2013301200233030-3232011301102332"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132010013302303-2231121100221123-2032212300023232-2230033221212112-2021130030212011-1302313133213232-1332303311200301-1000120021013203"></a>

<a id="canonical-1111202331013323-3102223020322013-3101003300320210-2122102210221233-0101302330212012-0012031020222003-0130222022003112-0120332131033300"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3022131132312310-0212133122233321-0323330013001300-2110210032211101-1210112200231231-3230210223033112-0032213110103222-3013100000022203"></a>

<a id="canonical-0332312312010111-3220211023221303-0110113320123021-2012022010000330-1232201233100202-2133000101213203-3110101103331220-3233211131113103"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-3331313001233112-2123201322332322-3311202313333102-1333330010023000-3312130223020000-2030112110121210-2231302020020110-3311000011021013"></a>

<a id="canonical-0003133223333102-3020111312232222-1020100203212032-1202012132300212-0332011213121200-1013223202012103-2000001233103222-0113230012312110"></a>

#### `name` property

Type: `"string"`. Required.

Name of the User Identification. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2330131113100023-1103300223311021-2310223200020022-3230202103101000-1111130012033121-0211213030103302-0211223233302230-1122221030210030"></a>

<a id="canonical-0123230320130130-2302133123112331-1023001021021322-3233322233022323-0212013023331011-2122100101022120-0103103123320203-0330321301133300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the User Identification is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203): complete subsection reference.

- [timeouts](resources--user_identification--reference--group-001.md#canonical-0202011013130232-3012230102103331-1133010102132130-1220213303130113-1000213000221020-3331310122110313-0223033232101123-3122302222310323): complete subsection reference.

<a id="canonical-2222021213121301-0121330312100323-3110123232312301-0113332333030200-1130022122322300-0233333011100320-2232323002302220-3233300321221120"></a>

### All schema paths for `xcsh_user_identification`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--user_identification--reference--group-001.md#canonical-2213312132012231-1203022113312313-0103323222122230-3202321213302210-3031201101021110-3213011201111132-0300331120322030-3222312030210330) |
| `description` | [description](resources--user_identification--reference--group-001.md#canonical-1302320333320221-1312112322023131-0222221331000301-0313331101113031-2203133220032313-0223210121232132-0330000011113323-0011221331000130) |
| `disable` | [disable](resources--user_identification--reference--group-001.md#canonical-1223330110200303-2032231011000031-0000111311023331-0233031000301002-0201122213132212-2010223122013033-0312210200312030-1333110111033030) |
| `id` | [ID](resources--user_identification--reference--group-001.md#canonical-1132010013302303-2231121100221123-2032212300023232-2230033221212112-2021130030212011-1302313133213232-1332303311200301-1000120021013203) |
| `labels` | [labels](resources--user_identification--reference--group-001.md#canonical-3022131132312310-0212133122233321-0323330013001300-2110210032211101-1210112200231231-3230210223033112-0032213110103222-3013100000022203) |
| `name` | [name](resources--user_identification--reference--group-001.md#canonical-3331313001233112-2123201322332322-3311202313333102-1333330010023000-3312130223020000-2030112110121210-2231302020020110-3311000011021013) |
| `namespace` | [namespace](resources--user_identification--reference--group-001.md#canonical-2330131113100023-1103300223311021-2310223200020022-3230202103101000-1111130012033121-0211213030103302-0211223233302230-1122221030210030) |
| `rules` | [rules](resources--user_identification--reference--group-001.md#canonical-0320322210231110-0032210033213220-0331032221100313-0302320221212101-0310100101113232-0202002133133211-1130230002013202-1130203312331313) |
| `rules.client_asn` | [rules.client_asn](resources--user_identification--reference--group-001.md#canonical-3303311233301101-2121303001112332-2002021230333203-1133002201123332-3301100020201032-0220013312313031-1230020122300201-2322030300322223) |
| `rules.client_city` | [rules.client_city](resources--user_identification--reference--group-001.md#canonical-1132310103123002-1021132303232333-3112303130310111-3101100331010221-0332302233101203-3310222100112032-2001223211010003-3303123133111202) |
| `rules.client_country` | [rules.client_country](resources--user_identification--reference--group-001.md#canonical-2013020232231100-1122231103011202-2011220222133322-2021300120212133-3312220000323021-0120331213121113-0123130323300101-0210313000323230) |
| `rules.client_ip` | [rules.client_ip](resources--user_identification--reference--group-001.md#canonical-0302301130313330-0202232223220230-3132033110230122-2001021320201332-2300130112020121-2330013121300103-0130121003032113-1122220312110300) |
| `rules.client_region` | [rules.client_region](resources--user_identification--reference--group-001.md#canonical-0133200113331203-2002220212133110-3120101013020021-0021033002303322-0122103132131203-1232231110212022-1232223012102010-3313332331011201) |
| `rules.cookie_name` | [rules.cookie_name](resources--user_identification--reference--group-001.md#canonical-0100122313301122-2323120003320130-1130130120230113-3122130123201210-1020200313302203-2012013131312233-3221213332321013-1032031102131100) |
| `rules.http_header_name` | [rules.http_header_name](resources--user_identification--reference--group-001.md#canonical-3220102113031101-1310120322012202-2010121132121123-1101233012101003-3232000203313202-3113021230200112-2021020320123220-3020231222032203) |
| `rules.ip_and_http_header_name` | [rules.ip_and_http_header_name](resources--user_identification--reference--group-001.md#canonical-0030132200031013-0002331313112121-2213010333032021-0331132213221220-0212220101231312-1232330102101000-3212221130332131-2120010210232000) |
| `rules.ip_and_ja4_tls_fingerprint` | [rules.ip_and_ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-3300331320130013-3130200320003112-2231131111231232-2330213000012210-1102103013201001-1112131330333013-2102302121323130-3131301310210111) |
| `rules.ip_and_tls_fingerprint` | [rules.ip_and_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-1110331120031113-3112111230321113-2112221120300100-2233302331021231-2202000330111203-1111310112113211-3202230202013002-2131030002323333) |
| `rules.ja4_tls_fingerprint` | [rules.ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-0112232223031021-0130200231013332-1121022123232131-3220232333210123-3333021001131212-1031031230201320-1220210301233121-1113103121010301) |
| `rules.jwt_claim_name` | [rules.jwt_claim_name](resources--user_identification--reference--group-001.md#canonical-3120303000110232-1020213133130131-3313020032231321-3122220032312031-3310120331113022-0013220330221011-3112212233322300-3011033032102230) |
| `rules.none` | [rules.none](resources--user_identification--reference--group-001.md#canonical-2232300020320132-3003210322212102-3023213320203231-2101301021302010-1333332233003220-0321123313101210-1131103221133133-0233133121102203) |
| `rules.query_param_key` | [rules.query_param_key](resources--user_identification--reference--group-001.md#canonical-0123201101233231-3012003013030011-0013312331322313-2230113233300302-0203323010220110-3222130032122132-0201033122023323-0210211233101033) |
| `rules.tls_fingerprint` | [rules.tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-1012020323130213-3023202211311131-1111332210123011-0133320233132332-3130221201332221-0332203012103113-0033212021030133-1330030010333122) |
| `timeouts` | [timeouts](resources--user_identification--reference--group-001.md#canonical-2212312302333211-2232323231023232-2312221113101033-1200103003301023-0212220100200312-2311121222121000-1130023313123301-3111010201202010) |
| `timeouts.create` | [timeouts.create](resources--user_identification--reference--group-001.md#canonical-1230011233323110-0031330130313000-0132023221221213-2330132003110002-2331210020003211-1120022010100103-0303212302100031-3101230133001320) |
| `timeouts.delete` | [timeouts.delete](resources--user_identification--reference--group-001.md#canonical-3113331011210032-2310301222302133-2023311311103120-1213231330012031-3033022103232132-2320131211202102-3101213111002330-1012132301111002) |
| `timeouts.read` | [timeouts.read](resources--user_identification--reference--group-001.md#canonical-3111210110213010-1213203232112032-2122031232232131-1201331112032330-0233313021022233-1021313213221200-3000122312022321-3213333132001223) |
| `timeouts.update` | [timeouts.update](resources--user_identification--reference--group-001.md#canonical-2302202132313211-2013303211210312-1122332021323233-1322302002002120-1203101320232010-0201211203003221-3021102032013310-0022010103221010) |

<a id="canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- rules

<a id="canonical-0320322210231110-0032210033213220-0331032221100313-0302320221212101-0310100101113232-0202002133133211-1130230002013202-1130203312331313"></a>

Type: `"object"`. list nested block, Optional.

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("client_asn",
    "client_city"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_asn",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "none"),
  validators.ConflictingListObjectAttributes("client_asn",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_asn",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_city",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "none"),
  validators.ConflictingListObjectAttributes("client_city",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_city",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_country",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "none"),
  validators.ConflictingListObjectAttributes("client_country",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_country",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_ip",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "none"),
  validators.ConflictingListObjectAttributes("client_ip",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_ip",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "none"),
  validators.ConflictingListObjectAttributes("client_region",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_region",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "none"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "none"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("none",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("none",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("query_param_key",
    "tls_fingerprint")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330212130330032-1232222003121123-0332102311320311-1212313202110312-2220110133013133-3132120133301112-1001311013210230-0320333010002013"></a>

### Direct properties for `rules`

- [client_asn](resources--user_identification--reference--group-001.md#canonical-0203121323212012-2311001132210100-2020332220323101-1033131223323032-3112223013332222-0232021130213122-3110201220113003-2130331010123022): complete subsection reference.

- [client_city](resources--user_identification--reference--group-001.md#canonical-3311130321010331-3323021212302320-2030123030310033-1121201303323103-0032230333132000-3221110011303120-0001220231033131-0110311031223310): complete subsection reference.

- [client_country](resources--user_identification--reference--group-001.md#canonical-1212201131133223-2002120013132312-1311200302110022-2311131133021132-3301130022112302-2221032212230301-0302231231213312-2012011232100331): complete subsection reference.

- [client_ip](resources--user_identification--reference--group-001.md#canonical-0013220003330023-0211312132131031-1030230121120312-1332033131112200-2033020200311131-0311113012112203-3232121102130013-2003020331313302): complete subsection reference.

- [client_region](resources--user_identification--reference--group-001.md#canonical-1210103232230301-2313101002103213-2311021121221131-2110101300222101-2231032233110330-2233211103103110-0303121133311232-0200013230122301): complete subsection reference.

<a id="canonical-0100122313301122-2323120003320130-1130130120230113-3122130123201210-1020200313302203-2012013131312233-3221213332321013-1032031102131100"></a>

<a id="canonical-3212011113033312-0231322000312210-1000131111020221-0211202303232333-3333313122312032-2333001331302020-0222212033133303-1313210313123323"></a>

#### `rules.cookie_name` property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3220102113031101-1310120322012202-2010121132121123-1101233012101003-3232000203313202-3113021230200112-2021020320123220-3020231222032203"></a>

<a id="canonical-0032011131002320-3130022112322312-3010210110020122-3132201312030302-2210002213132003-0311122120110022-0030022220321230-3320120130220120"></a>

#### `rules.http_header_name` property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030132200031013-0002331313112121-2213010333032021-0331132213221220-0212220101231312-1232330102101000-3212221130332131-2120010210232000"></a>

<a id="canonical-1333031023001322-3213000323322330-2121222023010021-1121011300120230-3223210233011321-2111212023030133-1110000212203100-1031012022220203"></a>

#### `rules.ip_and_http_header_name` property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [ip_and_ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-1212200102212113-0032021310131113-1101123110001332-2101200310111123-2310211033311221-1231200331110221-2033102002210021-3003101302111213): complete subsection reference.

- [ip_and_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-2311120321201223-3201102323200101-2023020101303322-1232011222231021-1121223301223211-1231030121302331-3123211030232003-0323001332231000): complete subsection reference.

- [ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-1020301301331321-2022132200102101-1112313101233031-1111313001302023-2002122203113303-2322300000032220-3330112002331022-1233322313203200): complete subsection reference.

<a id="canonical-3120303000110232-1020213133130131-3313020032231321-3122220032312031-3310120331113022-0013220330221011-3112212233322300-3011033032102230"></a>

<a id="canonical-3133320202321222-3130313000023111-0212101032211132-0133121303123223-0311113123000122-3220023330003222-0312003313320011-0310000220310102"></a>

#### `rules.jwt_claim_name` property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [none](resources--user_identification--reference--group-001.md#canonical-3231112200120302-0330222132332310-3031312013011121-1021103230121323-2133311231003331-3313230031333230-0333003223023121-1022230322000202): complete subsection reference.

<a id="canonical-0123201101233231-3012003013030011-0013312331322313-2230113233300302-0203323010220110-3222130032122132-0201033122023323-0210211233101033"></a>

<a id="canonical-2102312111113011-2022120222312311-2223330221202323-0113311100233133-1031301033100111-3023320031323003-0013132203321230-0230330220330021"></a>

#### `rules.query_param_key` property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-1212020131212000-2200211121203101-1132000223113002-3100213130300001-0130131003213202-0231102210230100-3123133200212123-1112302111102200): complete subsection reference.

<a id="canonical-0203121323212012-2311001132210100-2020332220323101-1033131223323032-3112223013332222-0232021130213122-3110201220113003-2130331010123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_asn` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.client_asn

<a id="canonical-3303311233301101-2121303001112332-2002021230333203-1133002201123332-3301100020201032-0220013312313031-1230020122300201-2322030300322223"></a>

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
client_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311130321010331-3323021212302320-2030123030310033-1121201303323103-0032230333132000-3221110011303120-0001220231033131-0110311031223310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_city` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.client_city

<a id="canonical-1132310103123002-1021132303232333-3112303130310111-3101100331010221-0332302233101203-3310222100112032-2001223211010003-3303123133111202"></a>

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
client_city = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212201131133223-2002120013132312-1311200302110022-2311131133021132-3301130022112302-2221032212230301-0302231231213312-2012011232100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_country` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.client_country

<a id="canonical-2013020232231100-1122231103011202-2011220222133322-2021300120212133-3312220000323021-0120331213121113-0123130323300101-0210313000323230"></a>

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
client_country = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013220003330023-0211312132131031-1030230121120312-1332033131112200-2033020200311131-0311113012112203-3232121102130013-2003020331313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_ip` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.client_ip

<a id="canonical-0302301130313330-0202232223220230-3132033110230122-2001021320201332-2300130112020121-2330013121300103-0130121003032113-1122220312110300"></a>

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
client_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210103232230301-2313101002103213-2311021121221131-2110101300222101-2231032233110330-2233211103103110-0303121133311232-0200013230122301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.client_region` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.client_region

<a id="canonical-0133200113331203-2002220212133110-3120101013020021-0021033002303322-0122103132131203-1232231110212022-1232223012102010-3313332331011201"></a>

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
client_region = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212200102212113-0032021310131113-1101123110001332-2101200310111123-2310211033311221-1231200331110221-2033102002210021-3003101302111213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ip_and_ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.ip_and_ja4_tls_fingerprint

<a id="canonical-3300331320130013-3130200320003112-2231131111231232-2330213000012210-1102103013201001-1112131330333013-2102302121323130-3131301310210111"></a>

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
ip_and_ja4_tls_fingerprint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311120321201223-3201102323200101-2023020101303322-1232011222231021-1121223301223211-1231030121302331-3123211030232003-0323001332231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ip_and_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.ip_and_tls_fingerprint

<a id="canonical-1110331120031113-3112111230321113-2112221120300100-2233302331021231-2202000330111203-1111310112113211-3202230202013002-2131030002323333"></a>

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
ip_and_tls_fingerprint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020301301331321-2022132200102101-1112313101233031-1111313001302023-2002122203113303-2322300000032220-3330112002331022-1233322313203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.ja4_tls_fingerprint

<a id="canonical-0112232223031021-0130200231013332-1121022123232131-3220232333210123-3333021001131212-1031031230201320-1220210301233121-1113103121010301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ja4_tls_fingerprint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231112200120302-0330222132332310-3031312013011121-1021103230121323-2133311231003331-3313230031333230-0333003223023121-1022230322000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.none` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.none

<a id="canonical-2232300020320132-3003210322212102-3023213320203231-2101301021302010-1333332233003220-0321123313101210-1131103221133133-0233133121102203"></a>

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
none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212020131212000-2200211121203101-1132000223113002-3100213130300001-0130131003213202-0231102210230100-3123133200212123-1112302111102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.tls_fingerprint` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [rules](resources--user_identification--reference--group-001.md#canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203)
- rules.tls_fingerprint

<a id="canonical-1012020323130213-3023202211311131-1111332210123011-0133320233132332-3130221201332221-0332203012103113-0033212021030133-1330030010333122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
tls_fingerprint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202011013130232-3012230102103331-1133010102132130-1220213303130113-1000213000221020-3331310122110313-0223033232101123-3122302222310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- timeouts

<a id="canonical-2212312302333211-2232323231023232-2312221113101033-1200103003301023-0212220100200312-2311121222121000-1130023313123301-3111010201202010"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112233203311211-3011211210111330-3101100210133011-2222323202101212-2232312110231200-2111210120330012-2103313131021113-2330312301102213"></a>

### Direct properties for `timeouts`

<a id="canonical-1230011233323110-0031330130313000-0132023221221213-2330132003110002-2331210020003211-1120022010100103-0303212302100031-3101230133001320"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3113331011210032-2310301222302133-2023311311103120-1213231330012031-3033022103232132-2320131211202102-3101213111002330-1012132301111002"></a>

<a id="canonical-2103320132030321-0210331301220002-3333122021013103-1220011001231312-0112002103022332-3000033033330330-0323211220220131-0323011221022002"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3111210110213010-1213203232112032-2122031232232131-1201331112032330-0233313021022233-1021313213221200-3000122312022321-3213333132001223"></a>

<a id="canonical-1132001233322132-0302223111102012-2112333003132032-2330033023131201-2320133331332023-1002033100132021-2131012103200000-0331300101300103"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2302202132313211-2013303211210312-1122332021323233-1322302002002120-1203101320232010-0201211203003221-3021102032013310-0022010103221010"></a>

<a id="canonical-1011030321223220-0220132031010302-2310002110103132-0001210023330310-0023103322302000-2113103231202230-3211322122132211-3300100232032201"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
