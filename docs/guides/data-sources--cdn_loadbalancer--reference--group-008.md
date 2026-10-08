---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3221022013311210-1303333312110100-3231202033012211-1212032101120010-3012033110122123-2112021112133103-3031121010220331-2300033211120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0230031210110011-3222000202231331-3302213031220022-0321100121311020-3001121012111331-2322101031112303-1210021212202123-2122213113020220"></a>

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

<a id="canonical-3302131133301200-2312333313230321-0231022102310322-2030333233012300-0330222233233023-1030020001221232-2130212320321310-1200012022223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-2303131020031210-3232303331031133-2302011203331021-1020210302102120-0121133000013113-3002100303220010-3223213022322111-1313111313021032"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-0121103122313132-0000030011002021-3301020123003101-0321130212210332-0121031323200132-1330321322220320-3321220220310031-1100311033033020"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-2330130332222131-0233312222212031-1031212322311131-1222323021103020-1122003310030200-3200223201233010-2131320230322032-1232133112200130"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1312102301220300-2222002113221030-0131230101120312-2110121201022230-0222311320023222-1311232111300233-2010232312302112-1332211303101332"></a>

<a id="canonical-3223312201103312-3213300201321111-2120023002000230-0303220310110302-0021230210122013-2323320330301130-0101020000300012-2322300210333312"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2120231013122203-2213000233213133-2101003010133323-3010303203110311-3010032222011322-3331031100022032-2102320013221212-0011000033200111"></a>

<a id="canonical-0201113113300031-2020200001023032-3130202300123002-2001021130320202-3333230220202023-0011131132210123-3232233031230122-1331021313312311"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3002001232030111-2330311122330111-1001301203133100-0311232213323021-1222202301002021-2123322010233023-1200033010220110-1013223312300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1020213201003302-0320011123322032-0310021233120132-3313331100310013-1133211311311220-1120011211102233-3221123110331023-3033333112312000"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313200112232121-3300122331020022-2221313200103131-2033020123232210-1010133010002320-0113210100300122-2112030301311323-1230020003003213"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-1002221332113000-1330332111212132-3010030203101212-0233220013000130-0303320202023003-0000022031222223-0221122120201123-2220213123032000"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3133331222312303-3231011103031112-3223222110102100-1022323020332021-2021121222131231-1003013122021211-0330113130330332-3101211202003113"></a>

<a id="canonical-3132331312022333-0302010301113002-2132230021302130-3321003112102100-2323131011021101-3011322011031012-1032022121221201-2201101320011230"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0303132031033301-0102301230131313-1233011001011330-3221110132312121-1303013211012200-1331301112222131-2300311201113312-3332201113113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-1332122200323000-2322212020132000-1321122111101230-1102002332113111-0231133121012212-1200321132221333-2312232233032122-1012320110230020"></a>

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

<a id="canonical-1332333020121221-0211223210202330-0313122230133030-1012103311221321-3332001201203113-3022133013031311-0312101033102203-1333123131320103"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-1311003112001103-2001211122330020-3030331231112223-3112020323320313-0131323013321133-1332303133132130-0330323130010332-0321221120012323"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2300023200332301-3323023220330033-2101320313310301-1210122103121300-1130112130110111-2302212101211131-2212032133133223-2210012002012302"></a>

<a id="canonical-3322311332102230-1203112132023102-2212133321300130-2013111010022310-1300221001100021-1031312311130100-0030132331100131-2303323223022211"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1202202321203121-3032300333220332-1021210303212100-1131010303333100-0122230332103210-3010300321230233-0101221230110202-3331213002323303"></a>

<a id="canonical-2213123013011303-3201322110000230-1103023332210120-0221103101100131-3100013121010200-2330031103121212-3233221303120202-2230033112310220"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

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

<a id="canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insertion_rules

<a id="canonical-2031020211001333-0101031112333100-0102331001111233-1321001020213023-2212000222310122-3201110003023322-0102231122202110-0100130012330320"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3301021302112330-2203011121200120-2211332130310032-0010333003131133-1232102102123102-1101001113312010-1222121111002103-1000133022223322"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules`

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203): complete subsection reference.

<a id="canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0321310300010320-2103013132212003-0013201333301131-1212222113312330-2313330311112001-1211110112003330-2131110133011201-2021103200001212"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0012132012203101-0200133012212313-1221320303212232-2103230121131300-0032330330302202-3121031211320102-3130013232313122-2001220313201223"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1303010300323100-3202022113212012-0312201120103020-2003303323331302-0213232232213122-3330022022312320-3121120030112303-1032311302113203): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1201303333113203-2122223100332002-3231202110121231-0211123222003013-1111122133133223-0132303023010211-0220122303201231-0301223330331032): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2010123033223331-2020322103213330-2313232233301230-2010032011231200-2320012023222313-0132222230311120-1002211213031112-2001221121200300): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2221131121310021-2121220312021222-2103003303232102-1023032032313022-0321311030010023-1000332302000010-0331223111332311-2101101322211010): complete subsection reference.

<a id="canonical-1303010300323100-3202022113212012-0312201120103020-2003303323331302-0213232232213122-3330022022312320-3121120030112303-1032311302113203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0103301030001003-1112103211031321-2110130211333231-1323211013211222-2122003323210331-3031133231123112-1112100301233233-2020032110232120"></a>

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

<a id="canonical-1201303333113203-2122223100332002-3231202110121231-0211123222003013-1111122133133223-0132303023010211-0220122303201231-0301223330331032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1213101213202120-2021233320231001-1001120312131133-2003213301301100-2313123331101233-3012103102123123-2222323023302022-3313122201031202"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-0011202321233201-1221233130313122-1012223323101031-1011313112102200-2011102031323013-1112211232300233-3100212030010120-0201033203002202"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-3301312232013001-2221222121103313-3233313231322300-3010003222001123-3322211122320210-0112120300221200-0331332131131313-0213111332201111"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3010132213021321-3220000320322101-0112131330313232-1112133003213323-2021313221111313-0200202333333130-0122212003031110-2311323010030020"></a>

<a id="canonical-0212022211000233-1310130301312011-1112120210310212-1333003021021231-3033023122130112-1102013211101212-1012233121002313-3103021101122002"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2301102131022332-0313333012323321-2123203200003201-2223331112310132-3313022130031130-1111132133312201-0333102202333012-1320222330101110"></a>

<a id="canonical-3103003311132322-2323223101320332-2003101330212321-3112211220102330-3331033331033113-1231300032201022-1222000313210031-3003202221131310"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2010123033223331-2020322103213330-2313232233301230-2010032011231200-2320012023222313-0132222230311120-1002211213031112-2001221121200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-1102313311101201-1111102331010002-0123230102220200-1020021220322330-3002322120030022-1301313223302130-3333320223102100-3231012030022112"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012300231320010-1003021220133302-1022210213232231-1220021000323112-2200203002102031-3033332300013023-2323202310310211-1320002311322010"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-0202200130133231-0130303310313321-3001112132111303-1330102013011021-0120232030313000-1013331201111130-3313303302220203-3221112213002023"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1202001003131203-1231012233321011-2303223231233203-1100022110122321-1120310002201202-1332100202300113-2031203001022010-2111123033303130"></a>

<a id="canonical-2013312010310121-0222203132211203-0200231003110113-0333322133223130-2120201220321031-0230210120302012-2013310211213021-1033000221113303"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2221131121310021-2121220312021222-2103003303232102-1023032032313022-0321311030010023-1000332302000010-0331223111332311-2101101322211010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-3121112323033300-2102021020110010-0210100313213310-1131233213001001-2011111302303222-2101021132133010-2212120121102202-0300301002123021"></a>

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

<a id="canonical-2302003222002133-3200321003010131-2323122203323221-3002322100101323-3000300311332202-3132013233213100-2113300211210212-1112323022231230"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-0021230210033310-3120301232022223-1123202010312120-3133301300132231-3212210301012121-3003123121200112-3330321021233221-2120320222001200"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1021320132113112-0333100302210103-1331230301313020-2203233112232013-1201221032103112-3133132110122121-3312321313001202-3303022000233023"></a>

<a id="canonical-1100300121201323-1110123003023112-2102303110010331-1201022113010021-1033312203000223-3200110231233013-3211221101233303-3320322220213210"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2123013221310130-2322233231201033-0003032302233203-0002132330102021-0023132302212033-1001131212332123-1200013030232312-0321101310122021"></a>

<a id="canonical-0120130221230120-3331103222123320-2020303223033001-2012032322303233-2310230113030102-3012122233321333-0201011231312020-2130333233022032"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.regex` property

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

<a id="canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-0111001210212010-0103012321320001-0212033232231230-1230033323300103-0101000130333210-0031123002313201-2230213202001223-3220131203320300"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233132321100012-1211113131200131-3002303033220203-3231102302021131-0233101230022022-3311310112013031-3012311202100220-3232202201222223"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3033221120230023-2003032032110220-2121133021231222-2103312231132303-0300120302211022-3003202011203201-1011322302131033-3113331013013133): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3120310332302230-2022101100122101-1023223311311132-3231011323202100-2330133330200310-3322010020211011-0002322102330300-0322330010022223): complete subsection reference.

<a id="canonical-2232202132010011-0123311013333112-1320103001312331-3132220231110301-1310101310302320-3130303120202001-0030012002121100-1333102111201100"></a>

<a id="canonical-1002223202132020-0223132110033212-0230333311003322-1011201003131123-3223302201032310-2223022233133301-3013132023323201-0312221022001331"></a>

#### `bot_defense.policy.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1320002021202303-1300330312202300-2213101032300021-2330111323130110-3220102210313203-2033213200201201-0010313221302112-3030032012032012): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2001110231301133-3233002312320121-1323121031211213-1131301311230233-2220323210211320-0002232222312020-1120203313102001-2002212010321200): complete subsection reference.

<a id="canonical-3033221120230023-2003032032110220-2121133021231222-2103312231132303-0300120302211022-3003202011203201-1011322302131033-3113331013013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2131213112202030-3033113303202220-3212130232003330-0032013113022013-0022022221121321-2320332211333122-1210023322132203-3310030030331103"></a>

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

<a id="canonical-3120310332302230-2022101100122101-1023223311311132-3231011323202100-2330133330200310-3322010020211011-0002322102330300-0322330010022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-2222111010130012-2031111222312132-2331320322310303-2222221312032231-0123311110322102-0220323321200300-0133211103311012-3101213321321121"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-0110130213111332-0100303203102302-3331200203003133-1103213233211131-3033003003221300-1001222003030100-1221333012102113-0221030332102303"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-3102120033213113-0320333033201222-0213213333112312-1121303020100020-1020100133211202-2300033310302003-2331033311220231-0231022123100211"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1131002121132012-1013011220121012-0311331032201331-0131300201320233-0221330113032003-1000211310120000-1000322110233212-1300033113210131"></a>

<a id="canonical-0300211321310311-3000000322212133-3103211233121033-3001013032202231-1231032101113303-2123310303302030-0222032201213303-0013310331301210"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2230120010113122-1122203301211221-3311132200333003-3331331011130300-1132032111202012-1211211203011023-1202120113313102-3312031320033012"></a>

<a id="canonical-2102201310010102-2002231232332132-0220013211002332-0002203201333110-2301303331201012-2123200102000120-1200102231002011-3120301021120123"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1320002021202303-1300330312202300-2213101032300021-2330111323130110-3220102210313203-2033213200201201-0010313221302112-3030032012032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0110323131220222-3131311131002312-1023102110320130-0200311112223030-1003032231212322-0110211130100320-2201033212112212-3022002211322023"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011220131300323-0000003332321333-2112130230131210-0111322033120030-0233311322023311-2223223223210333-3010310100020003-1120312302301113"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-0030100211300001-1203211203213023-0002023320331021-0000013102310002-1031133222333230-2223331303310311-0011022000302331-3113233112211210"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1112212231303301-0211132032221130-2313020212101300-0031131001110131-0312123023130111-2231000331202321-0112213223001020-3100102133112012"></a>

<a id="canonical-2332330212233223-2123112113220312-0031231212130200-1230232130010212-0122033022312211-0020302030022132-1123121203313010-1013110201230123"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2001110231301133-3233002312320121-1323121031211213-1131301311230233-2220323210211320-0002232222312020-1120203313102001-2002212010321200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-0002102211012300-3331010101003121-1130122200130232-3122323022301311-3322120033303003-1021213132302033-2311123311033122-2311003103323201"></a>

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

<a id="canonical-3013202321121332-3313110300333010-3000313122112000-0200322000220113-3020031130231332-2331011313223302-3110200101210303-3022020332020233"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-2212311222122110-2023220130103300-0120121303112102-1132002310332202-0100213231132023-0301200113330232-3020210210321103-2131302022320123"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2332030113213010-1203131210030002-1230211131211132-0210011302333032-3333032302023133-2100300312302011-2202132012213020-0222301211121230"></a>

<a id="canonical-2012233010311030-0332131300003112-2203200123331110-3232212112323212-2220112102131102-3032023120333302-1112201220030132-3031032211033003"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1131031100131223-3003113030230031-2303203021310011-3102132002102323-0220203101023000-1322002031013122-0323220022312102-3011111213122112"></a>

<a id="canonical-1323031001220102-1030331023030031-1231033310333032-0112120033201323-1233213121033000-1200200233311323-2120210230021222-3123310332102332"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.regex` property

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

<a id="canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-1021131320020210-2322012312332212-3333013322013133-2222101121013212-1013112132321302-3323333313320111-2322323012023301-1000022002301032"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103100210100322-2120202100021000-1123030032221211-3131012121132311-1220111120021331-1021023102000230-2133100103001231-3131233221332102"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config`

- [mobile_identifier](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122): complete subsection reference.

<a id="canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-3123030323130011-3322222301301003-2103012023010031-3211203133011230-1131303023201103-0321220000223321-2300103212113202-3131303001020202"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2201220011102232-1100023131002313-0213130022033222-2111312102133021-1330121013220113-3013011212210331-0210221122220013-0200132122233203"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier`

- [headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122): complete subsection reference.

<a id="canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1333103103300013-0031113132013020-3323202121320201-0020303022331312-3022003121323311-1000100010212002-3213032321103220-2312313001013022"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0310203211320121-0011232312102001-0012013123021321-1200232010103103-3211001132220022-2320011102310222-2300032201222112-1232201301021302"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2230123331100012-2112313012300233-1103321222232130-1120002222022332-0102112232013231-0120202013301132-0220121000032311-0000323100322130): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1013313220210201-2300232330310012-1332212123231230-2313022021101001-1110231331222100-2300002030323231-3330323110020300-1031332131200231): complete subsection reference.

- [item](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0210331112101033-0310233332310321-1313223210301212-2031111013102032-1302121221100301-0110011211032013-0012312330130013-0333021111103202): complete subsection reference.

<a id="canonical-3002010313302102-0103211221033123-3221011103033111-1302103321011232-3232102012100220-3101322110320112-3331022111221002-0013123223010120"></a>

<a id="canonical-3231132332030202-1302313021233103-3110220311231231-1310033132303201-0223233023130021-0030233032321321-3311222010120110-2002200310323300"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2230123331100012-2112313012300233-1103321222232130-1120002222022332-0102112232013231-0120202013301132-0220121000032311-0000323100322130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-3012330323131123-0002302131313200-3231100003212123-3011313110133103-0021200133121232-1112001210212110-0212301100133330-2032002211120221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1013313220210201-2300232330310012-1332212123231230-2313022021101001-1110231331222100-2300002030323231-3330323110020300-1031332131200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-3302032202220312-0003032200310220-0022030000202012-3033310031030011-2301210122331232-3321023000033030-1332230002100131-3330111200122021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-0210331112101033-0310233332310321-1313223210301212-2031111013102032-1302121221100301-0110011211032013-0012312330130013-0333021111103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-3021310220012122-3333313022231113-1312313131310010-2021232200302021-0313233131020203-1323212111022313-0313332332123020-1101210313231331"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103012303302002-0030222220123133-3302123301100213-1033113101032003-3221300310001311-0102120121322022-1021203002312230-1030311111303313"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item`

<a id="canonical-1122322110201123-3230231000303001-2113300313312313-1001032231021003-2332010320322313-2211230321232212-2312231120130132-1022313310332213"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320130131221311-1100331130003013-3213203220220331-3021002220111203-2300032133002212-2310201033120311-3310032223100131-2110020112311220"></a>

<a id="canonical-2032133013310010-1022030020010331-3101132110131002-1020211032123230-1332310003213303-2231003020321202-0133312231120003-1121312331030030"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2030110203302110-1120233303033012-0210121200223211-1131320302210330-1122333330202322-1320132030233211-3333130232001231-0213333113101132"></a>

<a id="canonical-3330121111013101-0213030222333211-0122111313210030-0212333222231112-0203220233103010-1220311320312333-0311100021320311-1232103132030121"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-3102110122102031-2222231231231322-0103210011120322-3023333223201112-1002103001131330-0302311333020020-2223103200323320-2302132220230322"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3102232130221100-1103303211323111-1322021233123211-3322020201321212-0031313301330211-2002301033302230-0330102220332103-3100210001001031"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints`

- [allow_good_bots](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1320332211131210-1021321312023222-3123210020310132-0013210002312322-2301331012131331-3131232120002211-1123011300223200-1201120201211202): complete subsection reference.

- [any_domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0323031113221202-0100021020320233-1112333032313203-0303232011320202-0200230311321232-2230123010331133-2203121232101213-0223011210332121): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2021123002221002-2000312022300103-0310231302200312-0302021212223332-2321010211102020-0222122100230102-1232133203021120-0013300000200020): complete subsection reference.

- [flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0131301311233211-1033013031310110-2112103222212000-3320331133030301-3233122330302113-0100023002323130-0221021031221210-0013231120031111): complete subsection reference.

<a id="canonical-3311121100201102-1101021211103221-2222330322221303-3321200102011322-3323220123031221-1120122330302200-2321300332311033-1122122231111303"></a>

<a id="canonical-0030011200130303-2112111333331131-0102300202022013-0300022200222221-3223303332110011-1121013330212003-2021330133121131-1010321332212333"></a>

#### `bot_defense.policy.protected_app_endpoints.http_methods` property

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0310323201120212-2211300010002233-0131322301103120-1313223022102130-3200302012100210-2312020222000331-1321211312313202-2130133321123302): complete subsection reference.

- [mitigate_good_bots](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3220231313001133-3232133123323110-2113200211332101-1220120031022303-3033110013001233-3303211003002222-1223032222003112-1321230220100320): complete subsection reference.

- [mitigation](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0200202233003202-0102330323232112-1032303100002320-0221321031020031-1300031023211332-3313011300222213-3311231221322001-1313131230321112): complete subsection reference.

- [mobile](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2131310300213221-3110023200313012-3033330132222210-2012123330230132-2313132313030233-2003133202020011-0032020300330200-2221013113321022): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3013311333023121-2031310033103232-2110030121212030-0223102001022133-1331321311100101-2101002012113311-3331321003101023-2231120301233202): complete subsection reference.

<a id="canonical-2130102213121103-1122310332212333-3011033310303101-2311102101011000-2010032023203332-1102230003100021-2310031032321313-0203322212131000"></a>

<a id="canonical-0203201133031220-1213112120033010-3023030120211133-2123010303022333-0311220231212011-0312002033223201-0321120231021310-1033220103211110"></a>

#### `bot_defense.policy.protected_app_endpoints.protocol` property

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200): complete subsection reference.

- [undefined_flow_label](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1003003211020113-3103312331332221-2103313103132131-1113322131300002-1031212331122122-3012103220120231-3032222232112132-2303111022032323): complete subsection reference.

- [web](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2222302100002000-1121300110111033-2303003223331311-0113322310023022-1123233232023213-3103232323330012-0002320230103130-1111220303320033): complete subsection reference.

- [web_mobile](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2212233111032332-2322130100100201-1312110231313203-3103222223203300-1323133131321310-3310021133322023-0121033021113312-1231221220103101): complete subsection reference.

<a id="canonical-1320332211131210-1021321312023222-3123210020310132-0013210002312322-2301331012131331-3131232120002211-1123011300223200-1201120201211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.allow_good_bots` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-1201122311212111-3303132203303022-0321322123023030-2230012131030002-2130002003101333-1033310303123033-3330133000311002-3210202132111112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow good bots.

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

<a id="canonical-0323031113221202-0100021020320233-1112333032313203-0303232011320202-0200230311321232-2230123010331133-2203121232101213-0223011210332121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-0221230232002202-0121031113223000-3102332331012212-1033123321310121-2002101201021210-2202331030022011-1312302102233210-3102020010323232"></a>

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

<a id="canonical-2021123002221002-2000312022300103-0310231302200312-0302021212223332-2321010211102020-0222122100230102-1232133203021120-0013300000200020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-2200311010131130-1300310021220000-1302331023102310-3332221033010130-0200133033323112-1301111213311003-1103330211230033-1023203301021313"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-2210021200132202-2123123022111231-0330100131110100-2103032211231021-0110132021002030-0311121002222020-0302020132213312-3233333222232233"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.domain`

<a id="canonical-0330002131302100-2032313030211330-2130300322210021-0020013022310123-1111233030033322-1011011223323031-1300221332221331-3321210201010303"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2210222301110113-0032233102000031-3333233230321003-1033222311023321-0031111301232100-1333200120200313-2310231101312021-3010032221021230"></a>

<a id="canonical-0003203223230120-3332201211331013-1313131300232321-0133203130030022-1220123211112202-3221111022131021-0331302223111212-3032010303202210"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2231211101131323-0132313313023330-2321222303130013-0233320202032321-0233212330333130-0112322021323011-1322012000300133-3312123032121301"></a>

<a id="canonical-2311303131120133-0311231312002320-3013002233301003-0033302013100321-0020121332202000-1012331001023102-1000021300110331-3220030112301022"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-3112330311201102-0010022301203030-0210332022033003-0132121311113131-1302032033231031-1100032320201032-0313110133233223-3221211123233320"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-3011332132221230-1202010223200332-1212311012233300-3210023102233030-0201223230233301-1121131300131001-3001102103033302-3332201323301201"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label`

- [account_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203): complete subsection reference.

- [authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023): complete subsection reference.

- [financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133): complete subsection reference.

- [flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1121213000022101-2113233123203300-3203013121221330-1213013330213330-1201312301011221-0132003011203223-2333122213232121-2110111331301203): complete subsection reference.

- [profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310): complete subsection reference.

- [search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022): complete subsection reference.

- [shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131): complete subsection reference.

<a id="canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-2020100332200013-3200211313312301-3003302010223211-2111302320002301-0213000002022131-0032010230322012-0230103310222210-1223022213222230"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-1123220302101031-3321212322300123-0132320313122230-2122202030101331-1210221133303031-2011003330221211-2003022311213120-2020323131100222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.account_management`

- [create](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0031033110322310-2330323301202210-3332222220211301-1120030021320203-0013210222300311-2103011222332200-3312123300323020-2212110211202103): complete subsection reference.

- [password_reset](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3232030013331112-3001310332222322-0231131002323203-2220203001330222-1031300122223013-3100321312111310-3033321323222033-2310011333132222): complete subsection reference.

<a id="canonical-0031033110322310-2330323301202210-3332222220211301-1120030021320203-0013210222300311-2103011222332200-3312123300323020-2212110211202103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-2221002311112031-3233002303331230-3121013020303122-1102100331220332-1031132131020130-0033201101100323-2232020020102131-0201221211123312"></a>

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

<a id="canonical-3232030013331112-3001310332222322-0231131002323203-2220203001330222-1031300122223013-3100321312111310-3033321323222033-2310011333132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-0003201131100133-1012131001000121-2132233001202310-2231321211310210-0231131320220233-3111011111023101-1331311021110112-0330211302031220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-0201123103200010-2210202301212311-2312331101322033-3330302203212033-2133313013022001-0231232032323011-3110230211112131-3300000023120232"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-0331332333220113-0231212112021133-2321232132230031-0303230210130230-0000321012032201-1332213021221211-2103310211222003-3012331201103132"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication`

- [login](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220): complete subsection reference.

- [login_mfa](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2322222331301102-3001311123013310-0233301000321302-3301322311311123-1013323220233020-3333330021301001-1013130121302231-0220002321210223): complete subsection reference.

- [login_partner](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3321320212303103-2130110313220111-0323210221221213-0003221223112033-2321222101333231-3032101012310230-3232110132233212-3010013102123021): complete subsection reference.

- [logout](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2313232313010221-1120111022102132-2122233321130312-0110030020330301-1332311320000223-3311021210322103-1000112101020100-0220212132313020): complete subsection reference.

- [token_refresh](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2312332120303010-0220113223322203-1233122231300322-0012222223302112-1330302230132121-0031110300330302-0231322311230003-3020031230001103): complete subsection reference.

<a id="canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-0011033201321023-0313020223032020-3023331123121022-2320311111200100-2213213111021210-3031011023331032-2202331101320122-2302333013322203"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-1320122321101111-3313113333231321-3123102032013100-2332031002023302-1023120022222200-2321021322203211-3300102331013211-2112330021212132"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`

- [disable_transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3233102101112130-2221002133001201-3120023200110232-2310201331023211-0131321332222323-2221033123311123-3122332120132302-0011220111132031): complete subsection reference.

- [transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313): complete subsection reference.

<a id="canonical-3233102101112130-2221002133001201-3120023200110232-2310201331023211-0131321332222323-2221033123311123-3122332120132302-0011220111132031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-2301221232331022-1032321131133122-1201313102222021-3230310002232130-0310231113231321-0021131120010133-1123302013311201-1210301323002020"></a>

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

<a id="canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-1300033212102121-0320222320023113-2312313001111333-0023332233201022-3002201113133232-0331232031333011-0131103100111303-0311312113212023"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100312101101110-2102002002001131-2210312303321313-3133131331311113-1100203102123032-2322302233030301-0033310003322031-2133323211103033"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3333332313233003-1111101303223020-2302013312223012-3323311321103223-2211031101002320-1103031310232110-0111001113013331-3322322310203002): complete subsection reference.

- [success_conditions](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3322302123133120-2230032022101000-2222230033322233-0312311010322102-0130203312130313-3113030123122212-3132201033001221-1322200020202200): complete subsection reference.

<a id="canonical-3333332313233003-1111101303223020-2302013312223012-3323311321103223-2211031101002320-1103031310232110-0111001113013331-3322322310203002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-3233113002133131-1322221021333131-1021021131211333-1131100333330220-3011133230322112-0231202030211310-0211030130131320-0323221020013032"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1221131232100031-1322213311030133-1111311130303301-3203233211312022-1200200230230023-0201231002122110-1323220013000221-3233201112222102"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-2100220012222213-1330232201310010-1023313220302221-1300302310132210-0013132303331311-2201233120310112-1031212120121330-1022103001323111"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1303021133201303-2102220012200122-3233112221223212-2323212212312000-1331101023123201-2233333110112111-1230031312011210-1000010000001011"></a>

<a id="canonical-3130211333320032-2230101213222221-3113320322130021-1233122303221330-2313102302330301-1322121200010131-0022110202101032-3003100213320110"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222220333022311-0033100020323031-1220100023332301-3003010010101121-0323332212312132-1030022100100323-3123211230013222-2012330232321011"></a>

<a id="canonical-2201120131323133-0310131310332303-2213121222333231-2333132023013103-2032001231301111-2110200133211222-2000322130001132-0303331001332232"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322302123133120-2230032022101000-2222230033322233-0312311010322102-0130203312130313-3113030123122212-3132201033001221-1322200020202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-1000213310222032-1113110312133123-2003100112323233-0010301130303111-0012001201203021-2103220033112302-3223323212121030-1200201232101103"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3010321130230033-3311121301323122-3211221230002030-2013210120031010-2113020222312321-1130111312310101-2202213000131313-0113102322113120"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-3130130322313111-3010213101301330-2103212232031223-0003101300300113-2101230022000311-1101100030132020-1332323223321112-0122210030313302"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2233001211000003-1002111120302332-2133131302111000-1220123003232001-0002331313010223-1021000101222233-3122232332331013-3020232302022131"></a>

<a id="canonical-3031310202312321-2212003201200330-2320101130122203-0010222022333001-2223332113020233-2033320233203001-1311003221101012-3010330210201303"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0001112310222230-2200202100031133-3133033330223211-0111233032130333-2332020213311220-2012313323103001-3122232231303331-1023110200332033"></a>

<a id="canonical-2212001121230112-2130221310320323-0203110311102013-1001233223113000-3031321301113203-1021322212321132-1110000133011230-3333020323021121"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322222331301102-3001311123013310-0233301000321302-3301322311311123-1013323220233020-3333330021301001-1013130121302231-0220002321210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-2033303322200220-3312001223223213-2101011302020303-2301030230311332-1231330130120323-3300023300022330-3022032130203231-3021202222200321"></a>

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

<a id="canonical-3321320212303103-2130110313220111-0323210221221213-0003221223112033-2321222101333231-3032101012310230-3232110132233212-3010013102123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-1221103111203222-1220330223023323-0212302133333201-1320202233112010-3302331112331102-2223122231102113-2223222003213230-2131120020122011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-2313232313010221-1120111022102132-2122233321130312-0110030020330301-1332311320000223-3311021210322103-1000112101020100-0220212132313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-1123033030230333-1332321311321323-1021200322023313-3230311132000321-1011121010032013-1120021033101231-1201311300120310-3102200110123031"></a>

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

<a id="canonical-2312332120303010-0220113223322203-1233122231300322-0012222223302112-1330302230132121-0031110300330302-0231322311230003-3020031230001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-1230110210031333-3221322231123132-2110110233032000-0123321203031111-2201013100112200-1010233303333313-0022121200031101-3330000223011222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-3322112033003213-1121003022131232-0030332130111001-0201103033002320-0223322202302010-3301201013202322-1332223232030030-1210023020320203"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-2233101123102202-2012200231313203-3030300330100233-1121322332011102-1201313300323120-2203123233001313-1131203032133230-3031111003000030"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.financial_services`

- [apply](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2210031130133030-1102202311112310-1020013330230101-2020212132213120-2112032131212323-2130323321332211-3311020121002030-1102121020210200): complete subsection reference.

- [money_transfer](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1223030100312031-1211010103001231-1212320213120113-1013100003130010-1110233133130100-0012213002033010-3330003302113000-2332023100202003): complete subsection reference.

<a id="canonical-2210031130133030-1102202311112310-1020013330230101-2020212132213120-2112032131212323-2130323321332211-3311020121002030-1102121020210200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-1111021000022000-1203123010133000-0022000231030113-1112300123222231-0221201132300233-3132131303022123-2023230032013002-0212331311210120"></a>

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

<a id="canonical-1223030100312031-1211010103001231-1212320213120113-1013100003130010-1110233133130100-0012213002033010-3330003302113000-2332023100202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-2220013221102222-0332032333030033-0132031230131030-0322033223121120-1002333133111220-2230011032200321-2111302033030210-3021111103020011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-1121213000022101-2113233123203300-3203013121221330-1213013330213330-1201312301011221-0132003011203223-2333122213232121-2110111331301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-0101322303222233-0232121120211123-2122202133313330-3100232201332212-1132330031200133-2120231110012012-3113103113121131-0001102233101103"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-2120103023001201-0321000231223021-2203322300323333-1013002323131201-1233033111010122-1310232220201012-3022330323212321-3113032030030111"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.flight`

- [checkin](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2200130022212213-1231001032210331-0311100110230110-0032303302233330-1332231111013001-0213230111233030-0202020121112121-3313200212021130): complete subsection reference.

<a id="canonical-2200130022212213-1231001032210331-0311100110230110-0032303302233330-1332231111013001-0213230111233030-0202020121112121-3313200212021130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1121213000022101-2113233123203300-3203013121221330-1213013330213330-1201312301011221-0132003011203223-2333122213232121-2110111331301203)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-1231132332211113-2321010010332300-0111112210023320-1312113012011002-0121320102002302-2010330320203220-3210211000313211-3233123013132210"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-0311110313120211-2231010221313120-2323310120300222-0200010021300303-0312223233231101-1032322211300010-2133113200012233-3312332320020231"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-0112120012310301-2213012010030000-3023103111320222-2120123131300122-2221201320310230-1130213111101332-2321202301113011-1121223210000133"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.profile_management`

- [create](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3123312113123303-1110320203220213-1202121012233211-1301033010012111-1013000000300003-2301222130101330-0300111132233300-2202323302032311): complete subsection reference.

- [update](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2233201122132120-3313012102032030-1323120230010212-0210123331231333-0312012212000332-3300233001210000-1100230123232031-2113321222000033): complete subsection reference.

- [view](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2031023010213011-3021212102022023-0221323013201023-0020020032001112-1301301310232133-1331122211221331-1002231132132303-2120012001220000): complete subsection reference.

<a id="canonical-3123312113123303-1110320203220213-1202121012233211-1301033010012111-1013000000300003-2301222130101330-0300111132233300-2202323302032311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-1003233203010010-3002121300303301-0231301020022331-0132321103313100-1002111121023020-1121300102011201-1133133321231030-3003303133100333"></a>

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

<a id="canonical-2233201122132120-3313012102032030-1323120230010212-0210123331231333-0312012212000332-3300233001210000-1100230123232031-2113321222000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-2102112003221310-3112321201111320-0212102312131301-0201130302210003-3233031030101110-0113013013312133-1301030112230331-0312223312302332"></a>

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

<a id="canonical-2031023010213011-3021212102022023-0221323013201023-0020020032001112-1301301310232133-1331122211221331-1002231132132303-2120012001220000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-1022313313001123-2202222011302330-0201310001031203-2320203333213331-3122213322230122-2112310011012211-2021022321200233-3213213203211211"></a>

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

<a id="canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-2302013212002313-2212101003021121-0013031123103122-0321020300303113-1310211302221120-2230230213200331-2300302230211030-2333033112130021"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

<a id="canonical-2010030321300100-0031020103033112-1332030133011132-3302003212020311-1333202233131020-0132112222330231-3210010330022333-3112001201113031"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.search`

- [flight_search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0133231201211013-0032130313222121-1013231310120221-0302231123201233-2002032202322121-1130331013310223-3332312302123103-1110020322321211): complete subsection reference.

- [product_search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0303302233001032-0122200233122303-2002223333223310-2122211210222303-0103233020112122-0313131212102310-0230303200312112-1013323201331311): complete subsection reference.

- [reservation_search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3012031101212101-0320033323022310-3300003032311333-0232210102212222-2301010031131033-2011123030300212-0213122213301301-1033223331223212): complete subsection reference.

- [room_search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1311231113313030-1131033101313111-3203110202032232-3321023310031332-0202231130012331-0321012021113302-3023102101301033-1323122102033300): complete subsection reference.

<a id="canonical-0133231201211013-0032130313222121-1013231310120221-0302231123201233-2002032202322121-1130331013310223-3332312302123103-1110020322321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-0011021100300311-0233131210100221-1130232323013112-1113330331122011-1311102202232030-2011331311002112-2310303132133303-3110312210331201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for flight search.

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

<a id="canonical-0303302233001032-0122200233122303-2002223333223310-2122211210222303-0103233020112122-0313131212102310-0230303200312112-1013323201331311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.product_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-2320322130220123-0300111322012111-2113030123313133-0230313012201300-1032130130013211-0322320012002212-0023102220221230-3332203223323100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for product search.

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

<a id="canonical-3012031101212101-0320033323022310-3300003032311333-0232210102212222-2301010031131033-2011123030300212-0213122213301301-1033223331223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-2133310102130222-1312112032120201-1333231020030210-0111100321021033-1303131312213333-0000221120332033-3133201013121311-1211200100231303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reservation search.

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

<a id="canonical-1311231113313030-1131033101313111-3203110202032232-3321023310031332-0202231130012331-0321012021113302-3023102101301033-1323122102033300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.room_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-3233120110233001-2213123212200302-3123010312301321-0202122210002313-0231202200030200-2210322010022312-3300222313122331-2311020000322103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for room search.

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

<a id="canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-2311110212013223-3000323320302331-1113230320301221-3112200122230210-0113011003221220-1321213212111013-1032332012330021-2103020202200022"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

<a id="canonical-0310013101102012-0121003121222033-1332031113231230-1200103330301112-3201112331011331-1203322201332322-1321011233013002-0322002020121231"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards`

- [gift_card_make_purchase_with_gift_card](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1231302011030210-2200311320311332-1011122003300023-3220022003302011-0221320002130102-1302100113221321-1310301021330111-1303021220022332): complete subsection reference.

- [gift_card_validation](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2230000123322101-3020310121122211-0003012133320112-0023023021220213-1220023302131232-2023311120201222-1130222121122332-1130332332110332): complete subsection reference.

- [shop_add_to_cart](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0330300120303221-3303123210010133-1330102002102200-1120001033210222-3132121113003222-3031233230122303-0121212031331333-1210002302331120): complete subsection reference.

- [shop_checkout](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3021131131023223-0112133111021310-1221123123233001-0110331331333220-3221220223323211-2003010210302302-1223021112211000-3002010333130310): complete subsection reference.

- [shop_choose_seat](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0312011333231001-2213011023012033-0233012120130123-3013112232112333-2013022001310321-2231232201113110-2130131233222303-2330112010110303): complete subsection reference.

- [shop_enter_drawing_submission](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1100220031013230-1320312101320021-1121023312201131-2201313103323133-3331012313311321-2023223100033033-2010001203313301-2201300131212333): complete subsection reference.

- [shop_make_payment](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2210030003013001-0310102122101321-1112233123133232-3103103300321130-0133203113331010-2130220020223312-2330221311133131-1001130211113103): complete subsection reference.

- [shop_order](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1121132221023033-1230202330331000-3103113133201121-3020311211022000-2321101221011000-2322002020003113-2110033332023000-3033103222023301): complete subsection reference.

- [shop_price_inquiry](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2323221200231033-2120102311133000-1113022103213223-0003013000120133-1201123201330120-1232002112301322-3001331111322113-0321220112311021): complete subsection reference.

- [shop_promo_code_validation](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1133231113010010-1130002023131223-2200233020011231-0113130131310221-2120202000210331-1213000000002011-0303120022211223-2132103202313002): complete subsection reference.

- [shop_purchase_gift_card](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1111223020202113-2101223333133002-3133012211000120-2320022133320203-2112123030333001-2030232322323330-0231130213201113-3303002320011233): complete subsection reference.

- [shop_update_quantity](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1130110022301002-2213311013213310-1333332113322030-1312320101331123-2221100120231030-0113033313213102-2121022301011333-1111202301200111): complete subsection reference.

<a id="canonical-1231302011030210-2200311320311332-1011122003300023-3220022003302011-0221320002130102-1302100113221321-1310301021330111-1303021220022332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-2003030313111332-1322321201021033-1002121311002110-3013200200202312-2332113003013323-0310232300300223-2120011121201120-3310310303230301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card make purchase with gift card.

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

<a id="canonical-2230000123322101-3020310121122211-0003012133320112-0023023021220213-1220023302131232-2023311120201222-1130222121122332-1130332332110332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-3223330313000002-3030330220120221-1101303231220030-3201120101330010-1331221130331010-3313012232001133-2310222303311121-0333312301131002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card validation.

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

<a id="canonical-0330300120303221-3303123210010133-1330102002102200-1120001033210222-3132121113003222-3031233230122303-0121212031331333-1210002302331120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-2120112322230230-1301332201321120-2230010201123031-2110003220111210-0121022230302002-3103013003211031-2120203011133020-3223011000223313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop add to cart.

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

<a id="canonical-3021131131023223-0112133111021310-1221123123233001-0110331331333220-3221220223323211-2003010210302302-1223021112211000-3002010333130310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0231131331033230-3113322322323030-2210212312222303-3312203223001320-1311030313003122-1321120221212301-0313113311020322-3323302210012320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop checkout.

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

<a id="canonical-0312011333231001-2213011023012033-0233012120130123-3013112232112333-2013022001310321-2231232201113110-2130131233222303-2330112010110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-0132211311020302-2211130232003311-2332331022123020-2001020020302323-1301001230311221-0201231101021100-1330120103032021-3101330330133003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop choose seat.

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

<a id="canonical-1100220031013230-1320312101320021-1121023312201131-2201313103323133-3331012313311321-2023223100033033-2010001203313301-2201300131212333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-0130013222311322-1101010302332330-1313030030120212-0100302320222323-1033001023123320-0213110330211133-0333211303213101-3202011313101122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop enter drawing submission.

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
