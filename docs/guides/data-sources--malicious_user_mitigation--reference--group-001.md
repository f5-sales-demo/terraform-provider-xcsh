---
page_title: "xcsh_malicious_user_mitigation reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation reference."
---

# xcsh_malicious_user_mitigation reference

<a id="canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- Property reference

<a id="canonical-2221233133002200-2223022101232130-1200313321030130-3231023100033313-3110231300001220-2122210330210203-3200110100011201-0111201112101320"></a>

### Direct properties for `xcsh_malicious_user_mitigation`

<a id="canonical-2301230003312131-3222223001300021-3011001021021121-0132330333003122-1010220321313113-3010232203232003-3011202122030011-2220132032323201"></a>

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

<a id="canonical-0211000301010000-2103100203212221-2011100021313111-3310313101113321-0312112001332003-2000320323303023-3300023330030103-2120212312222131"></a>

<a id="canonical-1003133013002022-0133212111223231-0033130022033230-1302310102320030-0233021223233010-1203231203033202-1231003221122002-1123130131330022"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the MaliciousUserMitigation.

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

<a id="canonical-2231230302023031-0330231223210312-3300200201101321-0210231310003002-2020131230001322-3200023303001030-0112320211322303-0211322300222213"></a>

<a id="canonical-2323112032303313-3100012201312311-1020313023332133-3002031300101010-3323100032020012-2102012302003113-3321001201301230-3123302211211001"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2332001110001203-1121100000203221-3103010311230121-1121220102113332-1211230130320310-0232003032111103-3110322103022303-3313003222211323"></a>

<a id="canonical-0303132210003122-0200330020001312-0210332303213021-2031301222033231-0113203031013200-3302000313313121-3302020100011001-3323203013032111"></a>

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

- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211): complete subsection reference.

<a id="canonical-3331111031223322-3110020303220223-2011330111212221-2301000112103130-2133333010233111-2312313332333132-1331010332020112-0322322030321302"></a>

<a id="canonical-0001203223011213-2223323303002233-0111133313002011-0312221300330312-3211130210310322-3311331112220030-0232111002322101-0330210231332123"></a>

#### `name` property

Type: `"string"`. Required.

Name of the MaliciousUserMitigation.

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

<a id="canonical-2020132201100322-2111110002123333-3011023113220011-0123212121111131-2300110200201133-1313213112001202-0300201010000112-2302331211011010"></a>

<a id="canonical-3010222110130202-3031333031121101-2330121320022030-2321201020312010-0230201133220103-2010332312013223-1020101313100030-3112212120011010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the MaliciousUserMitigation exists.

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

<a id="canonical-3322330332111200-1011022120300131-1111201032002303-2122002130313223-0132000002332102-3012022212322231-0001103332113210-0032311322303102"></a>

### All schema paths for `xcsh_malicious_user_mitigation`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2301230003312131-3222223001300021-3011001021021121-0132330333003122-1010220321313113-3010232203232003-3011202122030011-2220132032323201) |
| `description` | [description](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0211000301010000-2103100203212221-2011100021313111-3310313101113321-0312112001332003-2000320323303023-3300023330030103-2120212312222131) |
| `id` | [ID](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2231230302023031-0330231223210312-3300200201101321-0210231310003002-2020131230001322-3200023303001030-0112320211322303-0211322300222213) |
| `labels` | [labels](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2332001110001203-1121100000203221-3103010311230121-1121220102113332-1211230130320310-0232003032111103-3110322103022303-3313003222211323) |
| `mitigation_type` | [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2003020011100221-3301001332120301-1010212211032330-3022131110031211-1213212111221033-3220223321322100-3201100112323112-2223222111231302) |
| `mitigation_type.rules` | [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2030030030202010-0222010322032020-2311023101123020-3212201011003200-0322012100131310-1130332121202132-2000001320121210-0033221220212113) |
| `mitigation_type.rules.mitigation_action` | [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1330230200230002-2103023032220020-1222212331232020-2220000323201232-0032030030031331-1203310213132332-2131233012002211-3111330232113321) |
| `mitigation_type.rules.mitigation_action.block_temporarily` | [mitigation_type.rules.mitigation_action.block_temporarily](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0211222323103221-3023110221211330-2300130001000200-1233312302001031-1333233131211312-2003121211022003-1120331322133201-2230022303001310) |
| `mitigation_type.rules.mitigation_action.captcha_challenge` | [mitigation_type.rules.mitigation_action.captcha_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2122130111111021-2011232201113302-2012221331001201-0023011122003102-0301333313313010-3131002221013330-2002133030000030-0321320302030320) |
| `mitigation_type.rules.mitigation_action.javascript_challenge` | [mitigation_type.rules.mitigation_action.javascript_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0133312211012132-3133330223113120-2123213030021201-3303233011201033-3230301203111013-0321001330313311-0313300123200201-1010230100201033) |
| `mitigation_type.rules.threat_level` | [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3200111022212213-1311102021132312-2223212032200113-2210101113132301-3033311232020223-0122031220223032-1012301032231001-1020202131202221) |
| `mitigation_type.rules.threat_level.high` | [mitigation_type.rules.threat_level.high](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0123110203132002-1001233110103222-0122001323323223-0310200101110223-1212320133230111-1010232201203113-1231320113100310-2231312000022321) |
| `mitigation_type.rules.threat_level.low` | [mitigation_type.rules.threat_level.low](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0112231112231102-2000313221222111-2303313333031112-0200031203223131-2320231012022120-1130233313330313-3220301100301331-3300031013200213) |
| `mitigation_type.rules.threat_level.medium` | [mitigation_type.rules.threat_level.medium](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1321003002120013-1301201012311030-3300230213230301-1200330131003030-1333201111011303-1310001011121323-0112002231011012-3213022213210133) |
| `name` | [name](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3331111031223322-3110020303220223-2011330111212221-2301000112103130-2133333010233111-2312313332333132-1331010332020112-0322322030321302) |
| `namespace` | [namespace](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2020132201100322-2111110002123333-3011023113220011-0123212121111131-2300110200201133-1313213112001202-0300201010000112-2302331211011010) |

<a id="canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- mitigation_type

<a id="canonical-2003020011100221-3301001332120301-1010212211032330-3022131110031211-1213212111221033-3220223321322100-3201100112323112-2223222111231302"></a>

Type: `"single"`. Computed.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Additional upstream details:

The settings defined in malicious user mitigation specify what mitigation actions to take for user
determined to be at different threat levels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2312332003132220-1133220113210201-2320211312001312-2102033332330132-1120102231222210-0003022222032331-1222102322120320-0201321113002122"></a>

### Direct properties for `mitigation_type`

- [rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121): complete subsection reference.

<a id="canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- mitigation_type.rules

<a id="canonical-2030030030202010-0222010322032020-2311023101123020-3212201011003200-0322012100131310-1130332121202132-2000001320121210-0033221220212113"></a>

Type: `"list"`. Computed.

Define the threat levels and the corresponding mitigation actions to be taken.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  }
}
```

<a id="canonical-0330302310320313-1123331110233201-0331331213011322-3202101131331333-1320020300302103-2210300100310321-2213303031232230-3232030003122202"></a>

### Direct properties for `mitigation_type.rules`

- [mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103): complete subsection reference.

- [threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022): complete subsection reference.

<a id="canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.mitigation_action` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- mitigation_type.rules.mitigation_action

<a id="canonical-1330230200230002-2103023032220020-1222212331232020-2220000323201232-0032030030031331-1203310213132332-2131233012002211-3111330232113321"></a>

Type: `"single"`. Computed.

Supported actions that can be taken to mitigate malicious activity from a user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

<a id="canonical-2231230023122032-2022221310013202-2232230210102221-0023210022303112-3131232333331312-3003310331103113-3320011011020210-1211013101022012"></a>

### Direct properties for `mitigation_type.rules.mitigation_action`

- [block_temporarily](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3300313312222210-0103300001101220-0021302330333312-2022230023210330-0222221100303330-3121033003333211-1112300103022230-0322103231322131): complete subsection reference.

- [captcha_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3311331130130021-0200123222302002-3130331331022113-2313113303132133-0032230322302021-1322021111313202-3202000011210031-1232230211311231): complete subsection reference.

- [javascript_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1321330123332210-2221320013020213-3200013021202301-2301120100133113-2202332031033212-0202111101110122-2311031300211313-1303223320032011): complete subsection reference.

<a id="canonical-3300313312222210-0103300001101220-0021302330333312-2022230023210330-0222221100303330-3121033003333211-1112300103022230-0322103231322131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.mitigation_action.block_temporarily` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103)
- mitigation_type.rules.mitigation_action.block_temporarily

<a id="canonical-0211222323103221-3023110221211330-2300130001000200-1233312302001031-1333233131211312-2003121211022003-1120331322133201-2230022303001310"></a>

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

<a id="canonical-3311331130130021-0200123222302002-3130331331022113-2313113303132133-0032230322302021-1322021111313202-3202000011210031-1232230211311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.mitigation_action.captcha_challenge` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103)
- mitigation_type.rules.mitigation_action.captcha_challenge

<a id="canonical-2122130111111021-2011232201113302-2012221331001201-0023011122003102-0301333313313010-3131002221013330-2002133030000030-0321320302030320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for captcha challenge.

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

<a id="canonical-1321330123332210-2221320013020213-3200013021202301-2301120100133113-2202332031033212-0202111101110122-2311031300211313-1303223320032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.mitigation_action.javascript_challenge` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0030030303020100-0311203100121220-3331220233330221-0101300133002333-1121230110221131-1301220103000110-2210230100031301-1200132002320103)
- mitigation_type.rules.mitigation_action.javascript_challenge

<a id="canonical-0133312211012132-3133330223113120-2123213030021201-3303233011201033-3230301203111013-0321001330313311-0313300123200201-1010230100201033"></a>

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

<a id="canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.threat_level` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- mitigation_type.rules.threat_level

<a id="canonical-3200111022212213-1311102021132312-2223212032200113-2210101113132301-3033311232020223-0122031220223032-1012301032231001-1020202131202221"></a>

Type: `"single"`. Computed.

Threat level estimated for each user based on the user's activity and reputation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

<a id="canonical-2311100003002321-2003033312011131-1323022313100310-0130303100302000-0131300331203333-0203130311032122-2300232122021030-2220223300032333"></a>

### Direct properties for `mitigation_type.rules.threat_level`

- [high](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1133112220102123-0302302220213112-1030123023222003-3203102220302303-1230203102231102-2302110102001312-0333312333003230-2322123303123201): complete subsection reference.

- [low](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2212012332003110-2302132111100233-0331010313311133-2311111333231021-0300201210333333-2120132301220131-1003130023210133-0103203221030132): complete subsection reference.

- [medium](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2130131002102213-2201121031331333-2112302021013131-0033011032223000-3222100300012123-3331330213232120-2213303332211223-2110032202012013): complete subsection reference.

<a id="canonical-1133112220102123-0302302220213112-1030123023222003-3203102220302303-1230203102231102-2302110102001312-0333312333003230-2322123303123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.threat_level.high` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022)
- mitigation_type.rules.threat_level.high

<a id="canonical-0123110203132002-1001233110103222-0122001323323223-0310200101110223-1212320133230111-1010232201203113-1231320113100310-2231312000022321"></a>

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

<a id="canonical-2212012332003110-2302132111100233-0331010313311133-2311111333231021-0300201210333333-2120132301220131-1003130023210133-0103203221030132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.threat_level.low` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022)
- mitigation_type.rules.threat_level.low

<a id="canonical-0112231112231102-2000313221222111-2303313333031112-0200031203223131-2320231012022120-1130233313330313-3220301100301331-3300031013200213"></a>

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

<a id="canonical-2130131002102213-2201121031331333-2112302021013131-0033011032223000-3222100300012123-3331330213232120-2213303332211223-2110032202012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mitigation_type.rules.threat_level.medium` properties

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0032022102331111-1302130022121002-3103031203113323-2113033130202013-1101333123213223-3331103031001011-2302231301213012-3232220322101211)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0010202120012102-2133302201320223-0111010332312020-3232011323010010-3112032100230212-0103000312223210-3002211221233301-0302332311123121)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-2101121211102013-0231111031232020-2230030313301212-2333132212311123-2022133110130301-3110310310213022-3113323001223321-0022020313302022)
- mitigation_type.rules.threat_level.medium

<a id="canonical-1321003002120013-1301201012311030-3300230213230301-1200330131003030-1333201111011303-1310001011121323-0112002231011012-3213022213210133"></a>

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
