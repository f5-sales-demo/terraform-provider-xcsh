---
page_title: "xcsh_filter_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set reference."
---

# xcsh_filter_set reference

<a id="canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- Property reference

<a id="canonical-2312221023330300-1201122200010321-3220000100333023-3330231021022313-0310132031001022-1021001022232011-1313020112333133-2110321122300230"></a>

### Direct properties for `xcsh_filter_set`

<a id="canonical-2322030102133300-1112231123012313-3103121321300103-1321101203323331-2233020022302302-3110330012222210-2120311320203132-2013333023233112"></a>

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

<a id="canonical-3213311211230000-0210103000201301-2130133330321100-3212011121223333-2000202010030201-0333230320000022-1221131113130133-2123221021032213"></a>

<a id="canonical-0000100221202122-3303113230122302-1030023322230330-2313213020201001-1011200320113101-2032220013031103-1132232220331323-3030310310233110"></a>

#### `context_key` property

Type: `"string"`. Computed.

Indexable context key that identifies a page or page type for which the FilterSet is applicable.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3233211021033022-2213223131010112-2113102333122003-1122311310101000-0010013002132220-3313031100112210-2233210103003002-0003212001123100"></a>

<a id="canonical-2232202201201111-0213223121021000-2322300322211123-0222331010332002-3331323220001210-3320020231302331-2301120331313311-0132131203032110"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the FilterSet.

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

- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331): complete subsection reference.

<a id="canonical-3222321211030003-2023312101120112-3030200101321131-1230120322120131-0331121210000302-0232113033003232-3021101320202322-0321103313233130"></a>

<a id="canonical-0231113003220201-0220312122031020-3200112212013211-1202131322031300-3311011203320101-2330211023020020-2002311013132302-0330022302111320"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1212021232000302-2210302112203331-3002021011210102-0201310201303233-1201113130000313-3121211333011033-1101031113120333-1323232122120113"></a>

<a id="canonical-1221031012213310-3203100210103203-3110003333122323-2003131131310023-2331302131333033-1000311132202302-2332103233012023-1232021323002230"></a>

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

<a id="canonical-3210003312023001-3220021302200001-3202330301112212-3322121120030001-1010031113210311-3233202310322213-2032131002103220-2310001013022031"></a>

<a id="canonical-0221310103023002-0021232310030322-0213111311000212-2101103201032311-2312010201211012-1213111302200011-2012233133232121-3112231310312033"></a>

#### `name` property

Type: `"string"`. Required.

Name of the FilterSet.

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

<a id="canonical-2020011101310013-1132023210012311-0221202221111203-0023133203133212-1212010222232210-3103230222302233-2231210333131320-0212301231132023"></a>

<a id="canonical-3310332202023310-2031013221303011-3003133221132312-3220100203022112-2301302203313203-1110303232102203-3210113303310310-2011311110012112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the FilterSet exists.

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

<a id="canonical-3131132310101332-1302312011221332-2303202311001332-2203301123223320-1111312002130223-2300101101012132-3300001001031012-3011303302101311"></a>

### All schema paths for `xcsh_filter_set`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--filter_set--reference--group-001.md#canonical-2322030102133300-1112231123012313-3103121321300103-1321101203323331-2233020022302302-3110330012222210-2120311320203132-2013333023233112) |
| `context_key` | [context_key](data-sources--filter_set--reference--group-001.md#canonical-3213311211230000-0210103000201301-2130133330321100-3212011121223333-2000202010030201-0333230320000022-1221131113130133-2123221021032213) |
| `description` | [description](data-sources--filter_set--reference--group-001.md#canonical-3233211021033022-2213223131010112-2113102333122003-1122311310101000-0010013002132220-3313031100112210-2233210103003002-0003212001123100) |
| `filter_fields` | [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-2031310302131332-3321301211222211-2112002133000000-1322133201032213-1130231300212221-2012123301112021-3231201332001122-3020030123101032) |
| `filter_fields.date_field` | [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-2331313123020103-0232010010303201-3022201121303132-0103022231003133-2331331133021100-3113120213310020-1212001032230030-1111012010103020) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](data-sources--filter_set--reference--group-001.md#canonical-1303003033301002-2102102313320212-2111311122010201-2122012302013023-3021323130022110-2132110030202321-2022000320332132-2231012311333001) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](data-sources--filter_set--reference--group-001.md#canonical-1223020333003132-2132011123300200-2232012231202100-1010301311221103-2211111132213111-1120022003203131-3032321223130103-0022330110321201) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](data-sources--filter_set--reference--group-001.md#canonical-1011213100010030-3210211323311032-2322001212203320-2322133323111131-1112010313321101-3110331333130330-3000210101001222-0020113203210003) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](data-sources--filter_set--reference--group-001.md#canonical-0321333121133233-2120202303111301-3331033323122211-3030000331303220-2133310313303022-0320012321101312-1131310211200021-2002312210332223) |
| `filter_fields.field_id` | [filter_fields.field_id](data-sources--filter_set--reference--group-001.md#canonical-1112223310111103-2331203312010123-3002232203200203-1201103003323230-3302210130023102-3033031310030233-0000222022121011-0001003231122033) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](data-sources--filter_set--reference--group-001.md#canonical-3011210010200333-0220232023313303-1211121302301310-3321103133201121-0122030210220133-0120123123011100-3230130202121121-1333022013320233) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](data-sources--filter_set--reference--group-001.md#canonical-0210231220201132-3133232031301213-0233232322312130-3333122011201231-1212023030122133-3320101231101023-2222113212021100-3033123021231313) |
| `filter_fields.string_field` | [filter_fields.string_field](data-sources--filter_set--reference--group-001.md#canonical-0031122100220330-3313303010013201-1032323322123010-0000233210220212-0033100202123323-1003201010021300-0302101230132001-0211013111021210) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](data-sources--filter_set--reference--group-001.md#canonical-0213021012310321-3330322111302133-3330301110122001-3033220210022303-2022213212303332-2300221023201122-2111110232111301-0333321301033100) |
| `id` | [ID](data-sources--filter_set--reference--group-001.md#canonical-3222321211030003-2023312101120112-3030200101321131-1230120322120131-0331121210000302-0232113033003232-3021101320202322-0321103313233130) |
| `labels` | [labels](data-sources--filter_set--reference--group-001.md#canonical-1212021232000302-2210302112203331-3002021011210102-0201310201303233-1201113130000313-3121211333011033-1101031113120333-1323232122120113) |
| `name` | [name](data-sources--filter_set--reference--group-001.md#canonical-3210003312023001-3220021302200001-3202330301112212-3322121120030001-1010031113210311-3233202310322213-2032131002103220-2310001013022031) |
| `namespace` | [namespace](data-sources--filter_set--reference--group-001.md#canonical-2020011101310013-1132023210012311-0221202221111203-0023133203133212-1212010222232210-3103230222302233-2231210333131320-0212301231132023) |

<a id="canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields` properties

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- filter_fields

<a id="canonical-2031310302131332-3321301211222211-2112002133000000-1322133201032213-1130231300212221-2012123301112021-3231201332001122-3020030123101032"></a>

Type: `"list"`. Computed.

List of fields and their values selected by the user.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3221330333133313-3033123301211121-0000320320202230-1230031101130213-0232010202121210-3313023011020313-0131120201112201-1131033022233032"></a>

### Direct properties for `filter_fields`

- [date_field](data-sources--filter_set--reference--group-001.md#canonical-3001312312012001-2312220023302031-3121301222330313-0000210131021322-0101000321101211-2213333232122011-3320213333031130-3331130100333303): complete subsection reference.

<a id="canonical-1112223310111103-2331203312010123-3002232203200203-1201103003323230-3302210130023102-3033031310030233-0000222022121011-0001003231122033"></a>

<a id="canonical-2223030200332011-0231130102213122-2003033112120213-1233220001230300-1220103210233110-0323232321301132-0310102033322200-3031211010213321"></a>

#### `filter_fields.field_id` property

Type: `"string"`. Computed.

An identifier for the field that maps to some UI filter component.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [filter_expression_field](data-sources--filter_set--reference--group-001.md#canonical-2012103130001230-2322121013103302-0100211001222230-3300002022000120-2102320103311120-2120113212331331-3032232312010321-1112122332010332): complete subsection reference.

- [string_field](data-sources--filter_set--reference--group-001.md#canonical-2031230320001103-3213231000330120-3211130312032311-2312033002212011-0103322011023222-3232120311013323-2122003333102322-3231311310100230): complete subsection reference.

<a id="canonical-3001312312012001-2312220023302031-3121301222330313-0000210131021322-0101000321101211-2213333232122011-3320213333031130-3331130100333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.date_field` properties

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331)
- filter_fields.date_field

<a id="canonical-2331313123020103-0232010010303201-3022201121303132-0103022231003133-2331331133021100-3113120213310020-1212001032230030-1111012010103020"></a>

Type: `"single"`. Computed.

Either an absolute time range or a relative time interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

<a id="canonical-2312103212122210-1330110131112323-3212231001133003-0223001023320021-3331003000330302-3013121000033010-3013300121303033-0003033123200010"></a>

### Direct properties for `filter_fields.date_field`

- [absolute](data-sources--filter_set--reference--group-001.md#canonical-1312332221023310-0222323110202001-3102110301221333-3222003322210123-1121222032000211-1032023313123311-1301110223332321-0303110201112022): complete subsection reference.

<a id="canonical-0321333121133233-2120202303111301-3331033323122211-3030000331303220-2133310313303022-0320012321101312-1131310211200021-2002312210332223"></a>

<a id="canonical-1322323122233300-1301122113330121-0220303303231030-0301111022011310-0220200210311232-2221313311023002-0222231301232210-0012131101012313"></a>

#### `filter_fields.date_field.relative` property

Type: `"string"`. Computed.

Exclusive with \[absolute\] relative time duration.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1312332221023310-0222323110202001-3102110301221333-3222003322210123-1121222032000211-1032023313123311-1301110223332321-0303110201112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.date_field.absolute` properties

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331)
- [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-3001312312012001-2312220023302031-3121301222330313-0000210131021322-0101000321101211-2213333232122011-3320213333031130-3331130100333303)
- filter_fields.date_field.absolute

<a id="canonical-1303003033301002-2102102313320212-2111311122010201-2122012302013023-3021323130022110-2132110030202321-2022000320332132-2231012311333001"></a>

Type: `"single"`. Computed.

Date range is for selecting a date range.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212220131310212-0331320110223200-3012111212033210-0302020220012231-1003223100200212-2030131220100313-3012210231020011-3213012101100122"></a>

### Direct properties for `filter_fields.date_field.absolute`

<a id="canonical-1223020333003132-2132011123300200-2232012231202100-1010301311221103-2211111132213111-1120022003203131-3032321223130103-0022330110321201"></a>

#### `filter_fields.date_field.absolute.end_date` property

Type: `"string"`. Computed.

End Date. Contains end date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1011213100010030-3210211323311032-2322001212203320-2322133323111131-1112010313321101-3110331333130330-3000210101001222-0020113203210003"></a>

<a id="canonical-1033311123323300-3123003012331003-2332313002020022-1032030000301002-0300120212222021-1212131211202022-2031000303213003-1121303031010332"></a>

#### `filter_fields.date_field.absolute.start_date` property

Type: `"string"`. Computed.

Start Date. Contains start date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2012103130001230-2322121013103302-0100211001222230-3300002022000120-2102320103311120-2120113212331331-3032232312010321-1112122332010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.filter_expression_field` properties

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331)
- filter_fields.filter_expression_field

<a id="canonical-3011210010200333-0220232023313303-1211121302301310-3321103133201121-0122030210220133-0120123123011100-3230130202121121-1333022013320233"></a>

Type: `"single"`. Computed.

Filter Expression Field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2323330003330120-3002310002020020-0103033232322230-2201120211112003-2311121113020012-3020230212202220-3013131132003331-2222222313011231"></a>

### Direct properties for `filter_fields.filter_expression_field`

<a id="canonical-0210231220201132-3133232031301213-0233232322312130-3333122011201231-1212023030122133-3320101231101023-2222113212021100-3033123021231313"></a>

#### `filter_fields.filter_expression_field.expression` property

Type: `"string"`. Computed.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2031230320001103-3213231000330120-3211130312032311-2312033002212011-0103322011023222-3232120311013323-2122003333102322-3231311310100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.string_field` properties

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331)
- filter_fields.string_field

<a id="canonical-0031122100220330-3313303010013201-1032323322123010-0000233210220212-0033100202123323-1003201010021300-0302101230132001-0211013111021210"></a>

Type: `"single"`. Computed.

Filter String Field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122300102213022-2302223202311212-1222003110110013-2130221230130100-2331211333113220-0023001313211110-3330000110013032-2030230200001021"></a>

### Direct properties for `filter_fields.string_field`

<a id="canonical-0213021012310321-3330322111302133-3330301110122001-3033220210022303-2022213212303332-2300221023201122-2111110232111301-0333321301033100"></a>

#### `filter_fields.string_field.field_values` property

Type: `["list", "string"]`. Computed.

String Value(s). Field specification or configuration

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
