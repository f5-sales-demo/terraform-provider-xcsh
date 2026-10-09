---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-1322332203033330-0231332230212320-0322300021003213-0213020100110312-2200133203300120-3111120322210310-0330302021221303-1032021110203012"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122012023031032-1300220320331320-3201021023331300-2213211012101111-1200021033102113-0110002322333320-0313110020232010-1321211302300001"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-0013330210300103-1203112322213232-2003211320101132-1020013012003221-1213232023331020-0120332103113030-1031122233121332-2230011301211022): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-1312232001131331-1121013032013220-1221201221330323-2303101011333113-2212002001133202-1121002332302122-3123122002223111-2122223031221311): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-0233102133110232-3030221023310103-3202123303203321-1020133201330112-0310200211221303-1121210230121212-1220020032112313-0213123130213311): complete subsection reference.

<a id="canonical-0013330210300103-1203112322213232-2003211320101132-1020013012003221-1213232023331020-0120332103113030-1031122233121332-2230011301211022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0001203132213032-3200033312031022-1331132322321232-1223232333003203-3030000101121333-2132320323102111-3103012013332021-2020031112020130"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312232001131331-1121013032013220-1221201221330323-2303101011333113-2212002001133202-1121002332302122-3123122002223111-2122223031221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-0211122001222230-2200031303311020-3201020033223112-2013033213300131-0110111003030110-0001232311323022-2031210012311113-0232103231103122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220301132132001-0033300123131001-0302301022220331-0003212213110323-3123103113112212-0010123012230002-3213130213020011-0120210221300330"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-2302231333110321-1313022022023321-0033113301312112-3203123100021100-0002100202330330-1331021203031322-3033310021010313-2300211311203120"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1013212031322133-2221102002100223-0001022201222132-3030011033220021-3233231301121213-0331110202220313-3223011111201112-2211331013313020"></a>

<a id="canonical-2221010130300122-0101302113022030-1303333320021131-1321103010302100-1213213112020122-1100203210100012-1010323303111111-1013222031112123"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1213011212022322-0212202321321311-1000011101001312-0030022233012322-1110201032221102-1112232001221020-3112010233010323-0313020023203023"></a>

<a id="canonical-3212003023230020-2012132231123131-0100100022323103-0131202232002313-0320113010311322-0120233332133101-0200030031212101-1322300223032033"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-2323130231301302-1100113331012030-0220323210331031-1113333310012132-1323203011110132-0103211012022031-1321332133032211-1101332033310212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211210322301010-3110210331133011-3311020001033302-0223230131113332-1330111131133122-2030000202101302-3030131003213202-1211201332203002"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-2101032102102012-2231222312231310-3011222200230110-3111322300022131-2222302331002003-1233003003003221-0123031231000010-2320231032333320"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2021022012011331-2322232011021132-2201231032221333-0213001321310230-1311000131331201-2313331010010233-3120100113301123-3322033101000322"></a>

<a id="canonical-2220202102201101-3002122003020201-2313130313320032-0000202301133222-1322331220102122-0321031210220032-1303021302222121-1211200002231131"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0233102133110232-3030221023310103-3202123303203321-1020133201330112-0310200211221303-1121210230121212-1220020032112313-0213123130213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0032033202321002-0210212031111202-1331033003033211-1330023302332010-2122010231101000-3122012130321010-2330201301231103-3210323231212033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310233331011210-3001233202021111-1110333313110322-2133131121312000-1220232121131202-2113200013021212-2210310233231233-3103320112303102"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-0113231110320131-3321302222322112-2301130133322203-2310333132113022-1310221012302111-1132232230223222-0330203312120333-2111010211233322"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0331202130013030-1000313032300001-3110110303210121-3130012311310003-3132312231203123-1312312030203211-0232000212023010-0321313131332121"></a>

<a id="canonical-2201111322121212-2211210100321031-0323331032030301-0210212013213232-1322010320000132-0012310332002232-1010323022333033-2210320213313002"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0213331132102113-3002210232300033-3301302100102103-3320121102012112-1321210131120203-1031222223210132-1233233211103111-0010201312312320"></a>

<a id="canonical-3131222201201022-2330112000303010-3000122003123233-2222033211110020-2002220211233222-2232002021113122-1323300020302103-0122332110330120"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Optional.

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

<a id="canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insertion_rules

<a id="canonical-1103211231213033-2223301031103132-1031330333130321-3301020300113332-1201332231120320-0133103202220311-3332022102021102-2023130221210131"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320113312102032-1321212300331211-1222322102001031-3003333112032213-2012030322213300-1321230330122001-3012120003031223-1111300322202211"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules`

- [exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-012.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031): complete subsection reference.

<a id="canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-1013221201111120-3013220222301201-1221300111001323-0123032011012332-3320202313231113-2003333323131230-2220301213320130-1102313213200212"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121010220313011-0000230010132311-2020331301022201-0012333320302102-2112202132230113-0302011031011010-3032121213222232-1123231003132120"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-1221220223333201-0010122311110311-2231131121203202-1021031031023130-1203320002102002-3020112220130021-0302013010012232-2122212020133221): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-1301021021033303-3230122321212212-3201323201222312-1111132010310211-0023201303210130-3320213120010002-1222032302112321-2221201133123132): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-0222121320012233-3330310002200133-1301202100032031-0101210210020130-1211303233220213-0133131320100230-0013010000113030-1321000302230022): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-0113132333232330-3231013001310232-0233321111200331-0311220003200130-2233301330231221-0202003211332213-0300113013322033-3010203100112221): complete subsection reference.

<a id="canonical-1221220223333201-0010122311110311-2231131121203202-1021031031023130-1203320002102002-3020112220130021-0302013010012232-2122212020133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2100203121003221-0310010233213133-2022023120302212-2023023310221032-2112101311231122-3302211211033332-3230133101221012-2222121231323223"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301021021033303-3230122321212212-3201323201222312-1111132010310211-0023201303210130-3320213120010002-1222032302112321-2221201133123132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1130230202222103-1330030323201020-2102221223000111-0120101133022023-1003100232313020-2202301212213320-1013313010211021-2012033010303323"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222221120102201-3323100110111213-2101310221123120-2002311201233013-3300122333200322-1210312031222132-3332112213011020-1002101003223301"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-2013030002322001-1133223103010233-1120230323300103-2123133000003120-0032332112031010-1101300122033313-1213210003011012-2330313222221213"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-2222001330323212-2311111003300301-1120332011103320-2021322123201001-3022111323030221-2330322010102133-3332102330222020-0211010023201030"></a>

<a id="canonical-1011313101332222-2130232132021312-1310301003232330-1312313232123033-3001031100120220-2011122312130012-2012020103113223-1313333233333132"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0301101021021230-3030210210301132-3021320201031020-3321033110213323-2221130211122130-3303113221231313-3301101131231110-3121213321101002"></a>

<a id="canonical-1121002133010101-2300013221031300-3130321231102202-3221300330221212-0210232002222300-3210201222333021-1113203131220222-0223320033022002"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0222121320012233-3330310002200133-1301202100032031-0101210210020130-1211303233220213-0133131320100230-0013010000113030-1321000302230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-1222020333010313-2222102301120120-1331220131211023-0311131201333021-2230311030331233-2203333033302113-1002302323120010-1233121120213320"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321313012322030-3310030231322131-0221120333102323-1113211023231111-3110200001232113-1130131220212011-0032300000321332-3022302222221031"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1232002133113020-3320223033323113-0020013132332110-0201031330103211-0132231130222131-3123203121232011-1312302211201220-1331131113022013"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3330212231102011-0130323212313303-2331321131000131-3101111211301310-3330020320331212-1112132111310023-1113033103233200-0120211123300322"></a>

<a id="canonical-0022001221333320-1203120210300103-2212122223312020-1203000300301232-2032323020212313-0321103330212210-2320100323231011-0013132021323233"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0113132333232330-3231013001310232-0233321111200331-0311220003200130-2233301330231221-0202003211332213-0300113013322033-3010203100112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-3102010230332230-0112221300321110-0232130111120300-3001302322002123-1003131313311020-1231320130320001-0311032223120131-0230102222212230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130030231133231-2232120222023313-0221110032311023-2321023232122312-0130302121030131-3021033022033222-2200021012332020-3313003213112013"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-0032002131013102-1012022332010222-2301202022330220-3213330031311310-0100120322333132-3322331122022100-1001232013303231-2300031302311013"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0212102102330030-2103001110101202-1312232223300001-1232012033301323-2123202201121320-2302031110030002-3322011031203303-3112022103313211"></a>

<a id="canonical-1201130000210320-1001133030322001-1032031232333021-3313302132101010-3332001032031013-2321211300012031-3030203101103321-0120332213231132"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2303012013211102-2101322302223222-0032230302103111-2112203013100311-2033123303131121-1120033112330321-0330112002130011-3230233032010011"></a>

<a id="canonical-2212111030331303-0112301113022200-1122111122213221-2220332220123300-2301132330223302-1310100220030313-2011231132022233-1012011231321020"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Optional.

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

<a id="canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-2100002131133131-1031332133313121-0330303132011333-2313110111110102-2200220113202221-3332133322321313-3031232120130002-2000202122233221"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303300012201023-3231212131120232-1132332032220321-0111121102130303-2300003323320330-1103133110313200-3000223320222220-2320223113221233"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-1001002323002012-3110312013223130-2112013322032231-2121302330331200-3002011200120122-3200112210221223-0232130030100310-0100021123330310): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-2000321122000302-3110033333020130-1111013102331231-3013331112111232-3111201332210213-3130000122022333-2331111321011221-2030120200221031): complete subsection reference.

<a id="canonical-2203033303010310-2313012001212322-3003012123031133-1013023313022121-0203122132330103-3310003123021003-2022122302022211-3303101003332212"></a>

<a id="canonical-1322332230033112-1032033211220022-1211133000322100-0121301100011031-2313222100331313-3020310100120001-2030011121033013-2220301011211020"></a>

#### `bot_defense.policy.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Optional.

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

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-2333131102203011-1321321100300330-3323312330133033-3302313013310323-2122100130203302-3022020003311213-1011222232221301-1110303123111312): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-3111321310231223-2020000000111022-2010000011123120-1113323320130110-3012112311323323-1301302331103020-0003311010013012-0222003210220021): complete subsection reference.

<a id="canonical-1001002323002012-3110312013223130-2112013322032231-2121302330331200-3002011200120122-3200112210221223-0232130030100310-0100021123330310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2203330121121303-1323131203030300-0021322221303030-3233120220302120-2210101323300120-2021323211102330-3110103333101220-3123121103120320"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000321122000302-3110033333020130-1111013102331231-3013331112111232-3111201332210213-3130000122022333-2331111321011221-2030120200221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3223002013132311-2102320003330110-1001033122112321-2112121030002211-0130223321323133-2212131103311103-2000113202102330-3112013000032320"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201101313122123-1112131121130232-2023012112332223-2132301112012332-1332321131312103-3223201002321103-0331000001310320-1022301202033101"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-2202333231120313-3312103000320023-3122323221111222-2002000122112223-2011300001211202-1030322030010211-2111323331331033-3032010322120111"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1022313002003232-1101312221233013-3121230020220310-1102333301201300-1011120123300122-0003011213021032-1233213331323200-1230130102123033"></a>

<a id="canonical-3333121120033310-1100300303023100-2012111301300322-2013011211103102-2003013030323003-2123100012211032-2231232130333312-2323202323200201"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-3302311130231011-1232332122030021-0123112222210011-0211321322031113-1010220032200002-2211000320121130-0121002103002132-2233310231130111"></a>

<a id="canonical-2321330203031322-1220112032332132-3221112033113302-1031112010133303-2230332230001103-3212013111331213-2020031212022200-2233020331230302"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-2333131102203011-1321321100300330-3323312330133033-3302313013310323-2122100130203302-3022020003311213-1011222232221301-1110303123111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-2123300000320021-3221033021002013-0202122101023301-0121120121321003-2113102201302102-1112220033121231-0211112010001310-3132102010313031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033330013211231-2333210033001123-0021110233032230-0231211310310133-0001023300001313-3101003120301123-2110112132232003-3221123223022011"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-0302331203212100-0331003101331332-1000110022302032-0223021331222323-0030211230311000-0313303202330123-3322021200233200-2112232201321302"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2331230212213112-2020303232132133-3203003321111310-1231210022010031-1133200213121121-3231002131121020-0010030032203330-2001031231031222"></a>

<a id="canonical-3310011301113020-0322203223202211-3223003303231302-2031030133331101-2112310112202001-0130203202331313-2333222021020330-1010332221132100"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3111321310231223-2020000000111022-2010000011123120-1113323320130110-3012112311323323-1301302331103020-0003311010013012-0222003210220021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-2310110003003322-3312100300203213-2221110320303001-2210000000320011-1313211103033120-1121100132032112-1323203012321200-0103201113102203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213201032311310-2320000302211232-0031130323021012-2121022133122020-0223222313313332-2302113232302122-0223133012132332-1313300301233001"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-0313231013122002-3333202211130033-2111133222023012-0121233032130230-2131223002321202-1023200120023210-1002333211012133-2103031123112320"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2203330103300022-3031330130020100-1112122100133320-1010211210033320-2310302021031301-0021010123120010-2202032232222033-3232320300300123"></a>

<a id="canonical-2232120100231100-3102123321030200-0112130031120122-2231112230320323-1321303222300133-1111221113010221-0210102120213230-0032211312321300"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3311312210001230-0102231210322312-2132223013221300-2231211112032002-1301233100311323-1321112001203203-3211230113321230-3103322333022102"></a>

<a id="canonical-1013201113310213-2123010100220101-1103223212302312-0330130013102210-2013321321300212-0202330120133231-2033030302301321-1121232231221012"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Optional.

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

<a id="canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-0130221312123033-0313332210323112-0033311110202223-0302312211013330-3002202223300203-3321101321022011-3330221212111300-0201013100213013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303132233331322-1003112010212222-0031212002301231-3221122310121012-3003030333102223-2120233300011133-0222103312331000-1300121331113111"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config`

- [mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213): complete subsection reference.

<a id="canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-1023033232003103-3320133232023332-3032002001222220-3300211133111232-2300021323011221-3121220221302032-2012021211110023-3011133122233013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213032122222103-1330020110030132-0023111022203221-1212021031021103-3003001320202011-2303020222320303-1023003213200113-1110120331030103"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier`

- [headers](resources--http_loadbalancer--reference--group-012.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321): complete subsection reference.

<a id="canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-0123212033121002-1302313111011001-0112020301121321-3230233301130003-0230212030303312-3320310101130032-2312020001100001-2013311023131311"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333200333203300-3111220031332232-2333211222112012-2302232031031122-1331203201311202-3211110332330203-1030122310103203-2323033300303003"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-0011101013202231-3320022332210222-0221111210232201-0200302133323223-3000010030130102-0013200013301212-1332301001223303-3331102021320221): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-012.md#canonical-1210333020301111-3030330011120122-3133013022002203-0203213232003310-2201220000020222-0021021010011320-1032212332122310-0011311100231301): complete subsection reference.

- [item](resources--http_loadbalancer--reference--group-012.md#canonical-2222322013313020-0312211232300320-3103122221102233-2230133301322103-0200120220011002-2013021033213100-0233101301323221-3031123110012032): complete subsection reference.

<a id="canonical-3223301332113330-1321021313210130-1223332130323320-0300331310011013-1130120200131013-1231310320013122-0332022223322323-2111112301211012"></a>

<a id="canonical-0220112223021333-2212230203310030-3133331122223003-2013002312213032-1010110101030010-0032213021300113-2011231302312110-1213001200232213"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0011101013202231-3320022332210222-0221111210232201-0200302133323223-3000010030130102-0013200013301212-1332301001223303-3331102021320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-1133112122223310-2021103221121312-2310030100013333-0010102233301031-3201133022220000-2202113312311022-1220030022122021-0203210001101323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210333020301111-3030330011120122-3133013022002203-0203213232003310-2201220000020222-0021021010011320-1032212332122310-0011311100231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-3103001233332302-2321000103032021-3332230123301020-1330022220011232-0303313000213321-3111231033200112-1011023132200230-1320331102223013"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222322013313020-0312211232300320-3103122221102233-2230133301322103-0200120220011002-2013021033213100-0233101301323221-3031123110012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-0202122300121202-2230201332310123-1321313033301203-2032011323231221-3301132111232211-0213010231202102-1201131201013023-0123100320110311"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220300033123133-0330101102103003-1333221122002000-0021300201130210-3201320012102131-0322030122133003-2301132021223310-2232310201331230"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item`

<a id="canonical-2331011311311110-3123111222121010-2323332122032313-2000303013003033-1032132031322223-1200213212103330-2221202133110022-1300030022000312"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202302110321130-0321023030111230-2030331220003212-1103203330000021-1332010020333332-2113010332012111-1010310323310130-2011121200130230"></a>

<a id="canonical-3231323133201120-2322000212002332-2030333111100103-3223123202213310-1211120123010011-2132113320133220-1102322100110221-3332113321213232"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3213123331200210-1211001201203303-1233210001230103-1122321120030222-3103310112232133-2232221031131222-0302000123030211-3110113231332020"></a>

<a id="canonical-2033100022231322-1323023113132030-1120211103032220-3322033123210300-3133132222321300-1300110203120031-1133110323202012-3232133120223211"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-3122012111022201-0013313132031322-3323201303202211-2210110202222223-3112213100302320-3032210100003330-0322020232303102-1230001230323133"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322331020002122-0012221131111321-2322122020220311-3321033200310221-1232000120330033-2112321312131120-0120322222220103-1011333121310132"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints`

- [allow_good_bots](resources--http_loadbalancer--reference--group-012.md#canonical-1020312120302132-3022030103322213-3100322110233213-2122003033310021-0012313201123030-3033300203101301-0123132101123101-0221013130110003): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-1201322113311202-0002112110020122-2313113021202221-0023233220231122-2331333202131322-1121122310223233-1013231131113203-1200322213302013): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-1312130211202132-3212312301322212-1033322302133003-2031220203131211-0303200310203120-1302331210300113-1020132221310131-0012203132213203): complete subsection reference.

- [flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-013.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203): complete subsection reference.

<a id="canonical-1033110211130302-1303120301301220-1021310301121301-1222003332033113-0322130203010112-3030331103333221-3012222020213000-1213333131010303"></a>

<a id="canonical-2301332111013302-0332011203331013-0303220311123010-3313100002020210-0300002122200131-1120031032020032-2311230333003102-1212210021232000"></a>

#### `bot_defense.policy.protected_app_endpoints.http_methods` property

Type: `["list", "string"]`. Optional.

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

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-2023212323122011-1101323330132021-1310011202223312-0131122130110010-0101330133311212-3003232332003321-2322011030222200-1132132132002122): complete subsection reference.

- [mitigate_good_bots](resources--http_loadbalancer--reference--group-013.md#canonical-3200332103111311-3331103321320021-3300022321112003-2220210032301110-3121012100030020-3110320100033110-2213130032312123-2312232022333022): complete subsection reference.

- [mitigation](resources--http_loadbalancer--reference--group-013.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331): complete subsection reference.

- [mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0200310013012011-1313310020131023-0203321100301133-2101212302102010-3320231321202232-3031131303020130-0120030303011112-1220331131110320): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-3132020021201300-2213301200112303-1020320300322230-3032200333333311-3311101002332123-1010201023012022-0122220101130023-1320121003331033): complete subsection reference.

<a id="canonical-2231221210102333-1021330211133211-3120201111332003-0102000303303320-2110311002112010-3321031001032120-3312122010123203-2011311323123012"></a>

<a id="canonical-0021222313303012-2120021301130330-0212033210033231-2110332110302111-0232333112003102-3033221121103011-2311323201312221-0030133333100230"></a>

#### `bot_defense.policy.protected_app_endpoints.protocol` property

Type: `"string"`. Optional.

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

- [query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132): complete subsection reference.

- [undefined_flow_label](resources--http_loadbalancer--reference--group-013.md#canonical-1203210321222120-0033120203013313-2131212022012032-1233321301313330-3120311213033232-3302302000221200-3023302011011311-0232031130132223): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-013.md#canonical-0130020123113200-0230312201110222-1333301202301130-3302211302301012-2231221131222131-3321323201300333-0101100233311213-2020002023103320): complete subsection reference.

- [web_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-3101313121221202-1301032013003023-1220301222330001-2201122130031032-2313210123222031-3023312033232202-2101213232211113-0202330221311302): complete subsection reference.

<a id="canonical-1020312120302132-3022030103322213-3100322110233213-2122003033310021-0012313201123030-3033300203101301-0123132101123101-0221013130110003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.allow_good_bots` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-1310120223301222-2132112010203210-0000022322003002-0310121100202123-1111201131203120-3113123023320233-2312130121123322-0000331121003232"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
allow_good_bots = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201322113311202-0002112110020122-2313113021202221-0023233220231122-2331333202131322-1121122310223233-1013231131113203-1200322213302013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-1033312211011131-1122032112000010-3132233332010133-3030303020313022-1220331203222332-0332320131230312-3123031001300201-2011121232302013"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312130211202132-3212312301322212-1033322302133003-2031220203131211-0303200310203120-1302331210300113-1020132221310131-0012203132213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-3312320033120230-2330033230020223-2121210233101233-2310310102311203-2201022002103030-0301022010121302-2000201232233320-1130112003003103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232103010301233-3023103220102332-2310012112020233-3123003131022101-2221113300001321-0212322003132003-2302023123220212-0002003001300202"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.domain`

<a id="canonical-1123023123133302-3202220202120302-3102221320131023-2232333303030202-2010310220331212-3031312112133321-2212011202210121-1202000011202130"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-2302131001103212-1203330233113222-3313011312212010-0333213200221000-2103130211102303-0112301003301213-1330303131313212-1312122232002031"></a>

<a id="canonical-2111322232310321-1122331322321221-0002332130332210-3031300313132101-3103012000102021-0102110030201231-0131231102320311-3010203230210310"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1323003112100211-2323201222310023-3321021103320112-2230303331003112-0221022032312330-0230203323001012-3012012312310202-3011110112301202"></a>

<a id="canonical-0012210022300121-1221232331333031-0021130021010312-3122101201203023-0201132010121030-0330020022012112-2202022022303231-1320310122102201"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-0323032312113133-2022032203013122-0130102223023310-0022103013200001-0110122332201023-1131022203203023-2031222113223310-1213120003112300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231223323231111-3233130003111301-1233230202102022-2011323113011213-2221000200232232-0312130022203200-2311001023232122-0233031130322001"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label`

- [account_management](resources--http_loadbalancer--reference--group-012.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332): complete subsection reference.

- [authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302): complete subsection reference.

- [financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112): complete subsection reference.

- [flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132): complete subsection reference.

- [profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021): complete subsection reference.

- [search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121): complete subsection reference.

- [shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220): complete subsection reference.

<a id="canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-1001303332202133-1301130232113330-1332122011133222-0020212203111102-3110121221112122-3312131231002233-3131302122133233-1002300200303310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222020202133123-3211221132300023-3102311023333002-2012033032311112-0010321030330030-2330003132201012-1030312122123001-1122210031230003"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.account_management`

- [create](resources--http_loadbalancer--reference--group-012.md#canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223): complete subsection reference.

- [password_reset](resources--http_loadbalancer--reference--group-012.md#canonical-0122001002033033-0220122331200103-1032031030030033-3120113132220222-1000323220323300-2011023031301020-2123232322222330-0101231220210021): complete subsection reference.

<a id="canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-012.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-0030323002010212-1132010201210130-2233110020212023-0220032302023122-0332131131021323-2011203123312301-0223121033032311-0002230131110331"></a>

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122001002033033-0220122331200103-1032031030030033-3120113132220222-1000323220323300-2011023031301020-2123232322222330-0101231220210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-012.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-1011032000313301-1322303132013110-1232332133333013-3220132322302132-3013132332233011-2123113222130201-3330130003021213-1212030120130021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
password_reset = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-1133332031311322-1311020313030110-0010112310320222-2322320121033113-2020211130333300-0003001023102332-3023200030110310-0311003112130133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010100213022302-3302232130322211-0220320222312200-1302220032103031-1212200332331000-3232222111001123-3303032231011111-0112031200212220"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication`

- [login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133): complete subsection reference.

- [login_mfa](resources--http_loadbalancer--reference--group-012.md#canonical-2310113000131023-2200111020211101-2100210032102203-0332300332120322-3000122313320123-0202223030203021-1120231203232311-3003102313121100): complete subsection reference.

- [login_partner](resources--http_loadbalancer--reference--group-012.md#canonical-3220333122100331-3330111221312311-0133330123031210-1101111222312112-2011112023103020-1013310232003130-0323010222132310-2103111003002202): complete subsection reference.

- [logout](resources--http_loadbalancer--reference--group-012.md#canonical-1021233303223023-3111113131122032-2132210010322203-0320212131213331-3213030221323113-2211111022330123-1301300102012333-0132323333331321): complete subsection reference.

- [token_refresh](resources--http_loadbalancer--reference--group-012.md#canonical-2213022330021103-0212013331310113-2202010210023233-0020032032230122-0331132032133310-1113321111003033-2200321032022301-0113131200100033): complete subsection reference.

<a id="canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-2000023133232122-0031031300230132-1003333032112100-3223101333310311-1022032123222101-2301102323203112-3231311003322220-1123131312302120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121002102333331-0322220003121113-3210313310003311-2013103231022211-2213021310132010-0201332202122300-3311110211120031-0103311301222212"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`

- [disable_transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2002233110222300-3023110131333200-3002232311133002-0300020211021100-0310313001231321-2012123332113020-2201223333213103-1123011202211201): complete subsection reference.

- [transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100): complete subsection reference.

<a id="canonical-2002233110222300-3023110131333200-3002232311133002-0300020211021100-0310313001231321-2012123332113020-2201223333213103-1123011202211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-1230212012331220-3201000120232113-2323032010302323-2331112010203031-1233201333111210-2323333012222311-0223202122112221-0230302011033023"></a>

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
disable_transaction_result = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-2213203032213033-0033231122310222-2100321111000002-0102121120010033-1323313131312103-1220332301321210-1002001323031123-3211001322003003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211230112033101-1010202103333120-3023200033221031-3031332312211000-2010033212313321-2133021233113120-0221220113321310-2222031010003212"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-3232211032320313-1130031322033023-2221012313221033-0201130023033302-3311022132223311-2132101303331210-0033233300233213-3222103313032301): complete subsection reference.

- [success_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-2121212102313303-2313010110313312-2123211221312000-0233121300030033-0020323300200312-1232031220320310-0003323033212203-3233233103331030): complete subsection reference.

<a id="canonical-3232211032320313-1130031322033023-2221012313221033-0201130023033302-3311022132223311-2132101303331210-0033233300233213-3222103313032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-1220001002212120-0013022033000122-2312323233202321-1210300112101311-0202110302312002-2013003102120001-1032001010030201-2023210011010201"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203301111220322-1201011212213333-3331320133222130-3202321303032302-0210010220011103-1233212232223031-1230222012001310-3323031132331213"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-0100221213133020-0003123221310330-3003220010231321-2030303132200303-2101010022030331-0100110213311121-2213332110213110-2310320123212132"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0331100330021000-3121231121202220-0320210110301311-0021312200300302-3021010213012001-3011133233022212-3201233111330031-3202032032032102"></a>

<a id="canonical-2120000011033203-3311023103302232-2020121331302211-2012201321212021-1333103311100023-1130033100031333-0131130131212013-3032231222112132"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2112121011212212-1022111322200100-2033312113132220-3003001201132310-0000222331220131-0213322303310231-3000000010123020-1232333033103320"></a>

<a id="canonical-2021011232031023-2111101003032300-0012111302000223-3113112231112223-1222003100130321-1320321113313333-3122012033213120-3320202231132133"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

Type: `"string"`. Optional.

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

<a id="canonical-2121212102313303-2313010110313312-2123211221312000-0233121300030033-0020323300200312-1232031220320310-0003323033212203-3233233103331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-0132123133133003-0222312331000231-3002010302113221-1330132003002113-1233133100122023-1212330102211221-3011220030321000-3230113020001222"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322001201211010-0123321200330331-1001112030000230-0100102011012212-0102230301012302-0101010313201022-1320233100313203-3300311313102300"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-1310212002110102-3021310131312302-1012002302002301-2210300000001200-2132111312103132-0110311323332122-2010333033013023-3201133121201123"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1300212111021120-0332201233233100-3323202310101322-2030002220223011-1213320133323032-1113032110201211-0300330331313220-1122032110331233"></a>

<a id="canonical-0022331032211202-2302221033023121-1322022111122312-0033202203211311-2132330121112033-3311201132302202-1110003303233201-2102310220212022"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0013322233020021-3103302023212010-0331201000111122-1122313301032233-3212333311300311-1012203302131032-1103133101210012-3130023331211302"></a>

<a id="canonical-3021123211013230-0213212112302233-1022321222201132-0203003333320132-1202022301301212-2212121333223120-1031032100311003-3231212201111032"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

Type: `"string"`. Optional.

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

<a id="canonical-2310113000131023-2200111020211101-2100210032102203-0332300332120322-3000122313320123-0202223030203021-1120231203232311-3003102313121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3231230111201131-1132221302212231-3013122133130222-2021033100001001-3132123211211133-0221030103002112-1002330031020213-3011233321323332"></a>

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
login_mfa = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220333122100331-3330111221312311-0133330123031210-1101111222312112-2011112023103020-1013310232003130-0323010222132310-2103111003002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-1322110233302311-3102112301010030-3123313110123113-1011300232231331-0321131211233230-0100230211001200-1120330232033210-0020103010213122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
login_partner = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021233303223023-3111113131122032-2132210010322203-0320212131213331-3213030221323113-2211111022330123-1301300102012333-0132323333331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-1112113202210331-0311013230101212-2213011310031302-3332103201331220-1211320202303322-2130112301331202-2033310322221332-2133113110032012"></a>

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
logout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213022330021103-0212013331310113-2202010210023233-0020032032230122-0331132032133310-1113321111003033-2200321032022301-0113131200100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-012.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-1320323000202010-2121100233003322-3013121322021303-2012311211101301-2103322111131212-2123201002102303-1021233020233032-1311130001301112"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
token_refresh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-0023313203202312-3331201300113301-0020121131213012-1113121221002220-0122011123130321-3311100231132101-0223121301330300-3303332233222001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222203100322230-2011021211333321-0113320101011311-3023013310223230-1321121203233020-0031233321311012-2333130212000210-2230230322011222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.financial_services`

- [apply](resources--http_loadbalancer--reference--group-012.md#canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321): complete subsection reference.

- [money_transfer](resources--http_loadbalancer--reference--group-012.md#canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001): complete subsection reference.

<a id="canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-0303023201021203-1213331020131330-2121232130223013-0023102213212322-1232323320320101-1003212211102112-2333233020020303-2110111300023020"></a>

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
apply = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0000000033330222-2133212001221133-3120230302002120-0112000331223313-0220013330232103-0322031210302120-2232222330001331-1021230033132212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
money_transfer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-0112222123130203-0300003200212330-0222031122230301-0303100303302031-3032310120131132-2220033133001120-2002311232131032-3303122222032103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000131123232020-3331200202002101-1130022221023032-0202201221021201-3012120000020113-0321220233001021-0130012011003102-2320212322220222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.flight`

- [checkin](resources--http_loadbalancer--reference--group-012.md#canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200): complete subsection reference.

<a id="canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-0022211211033231-3330201033232230-2302333012201023-1310113121121123-3122322102130112-1210331131233323-0201011221210222-0321231200003220"></a>

Type: `"object"`. single nested block, Optional.

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
checkin {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-1123331100331221-3233022320303222-3010002023010313-0333302033211002-2233022023130030-0131222133203300-3331212323101213-3222122011321301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003213132023302-1013011011302233-2112233000231303-0232021301321330-2133011033130113-0223233201222310-0313310213113032-3132301330331311"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.profile_management`

- [create](resources--http_loadbalancer--reference--group-012.md#canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131): complete subsection reference.

- [update](resources--http_loadbalancer--reference--group-012.md#canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201): complete subsection reference.

- [view](resources--http_loadbalancer--reference--group-012.md#canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203): complete subsection reference.

<a id="canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-3112132212320123-0132131233323130-3000333221101133-2230201003300011-0203300023001230-0100233203213230-2221202331202232-1223021133010323"></a>

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-2211033010302312-2231010230330030-1012013003110302-1232021101132131-2121300210021132-2000031213012223-3022322220003310-3001120203003110"></a>

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
update = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-3101203232122311-2011111033011310-1013100001021311-0231131032121003-3221320013200030-2323122132311302-0222110233031330-3101230201111132"></a>

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
view = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-0230313123323323-0220201122122023-0301030122333211-0020003122122023-0210023231211332-3122202303211332-0132113310210201-1012011012210022"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212312203033233-0113311320023313-1211232222033201-3101020100333320-2230330312330213-3002320100020013-0210103301010022-0201330203031131"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.search`

- [flight_search](resources--http_loadbalancer--reference--group-012.md#canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032): complete subsection reference.

- [product_search](resources--http_loadbalancer--reference--group-012.md#canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102): complete subsection reference.

- [reservation_search](resources--http_loadbalancer--reference--group-012.md#canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212): complete subsection reference.

- [room_search](resources--http_loadbalancer--reference--group-012.md#canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133): complete subsection reference.

<a id="canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-3020011233310330-0323311030233213-0020003113013100-3330002031312103-0032303312101203-1012002113320330-1103011123221302-0032333222100100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
flight_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.product_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-1031211310300013-1000320303030231-2202322331203212-1021021131030310-0202331223320030-3003320022333312-3111120232010300-0213110120223330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
product_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-1010113203001113-3112203002030112-1222020112332232-0003232302223322-3221330112221021-1213302000313331-3202330032001110-0220022103230212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
reservation_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.room_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-1103211223102301-2002230123113121-2321131010301210-1223233300302213-1130110321131321-1200223331303033-2213333212332130-2232322012100101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
room_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-2312330022202233-0031211231120002-3010130013100312-2331212132230123-2100132331223010-2331213111201011-1210202211111330-1131320203012003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132001231321303-2320210012021323-1222331232123013-3300230013333333-0121123111031302-0211323012312230-2331121330101313-2330221313112232"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards`

- [gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103): complete subsection reference.

- [gift_card_validation](resources--http_loadbalancer--reference--group-012.md#canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223): complete subsection reference.

- [shop_add_to_cart](resources--http_loadbalancer--reference--group-012.md#canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203): complete subsection reference.

- [shop_checkout](resources--http_loadbalancer--reference--group-012.md#canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320): complete subsection reference.

- [shop_choose_seat](resources--http_loadbalancer--reference--group-013.md#canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300): complete subsection reference.

- [shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-013.md#canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301): complete subsection reference.

- [shop_make_payment](resources--http_loadbalancer--reference--group-013.md#canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012): complete subsection reference.

- [shop_order](resources--http_loadbalancer--reference--group-013.md#canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222): complete subsection reference.

- [shop_price_inquiry](resources--http_loadbalancer--reference--group-013.md#canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101): complete subsection reference.

- [shop_promo_code_validation](resources--http_loadbalancer--reference--group-013.md#canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211): complete subsection reference.

- [shop_purchase_gift_card](resources--http_loadbalancer--reference--group-013.md#canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123): complete subsection reference.

- [shop_update_quantity](resources--http_loadbalancer--reference--group-013.md#canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231): complete subsection reference.

<a id="canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0203201032022321-1101122012013302-1123023201002231-2333310320323031-1131211011333011-1203222223312020-0010021132303133-0023201002032122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
gift_card_make_purchase_with_gift_card = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2302300101313113-0221230010311211-3330323220321021-3211313301233112-3232210001312323-3120120320020131-0000031002120312-3122221203333311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
gift_card_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1312213101332122-3222210331132012-2233003203030201-2021000320011213-2321010022312103-0120032001201033-1203210212320020-3002200233222111"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
shop_add_to_cart = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0331000022302323-0303022031133110-0230123001303113-0120010230233031-3111222330131011-3132200321233022-0122230102102132-2011013333313201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
shop_checkout = {}
```

This is an empty object or choice marker. It has no direct properties.
