---
page_title: "xcsh_cdn_purge_command reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command reference."
---

# xcsh_cdn_purge_command reference

<a id="canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- Property reference

<a id="canonical-2333000210203223-3230012102131112-1021220303121220-3030331322022121-3020312113322220-3121102131122232-3312100003000210-3013113230323120"></a>

### Direct properties for `xcsh_cdn_purge_command`

<a id="canonical-1002331010331233-2331120200100110-3212121232023131-2232132210121312-3230330330010020-3221122131112001-0101003333121032-1211001122202220"></a>

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

<a id="canonical-2231021111330102-3332322133330003-2003212102301323-2103002102033300-2133102000233333-2012233021202112-1032200031132112-3212101132013030"></a>

<a id="canonical-3312321201212111-2123111002132302-0011320222200013-1223312231331220-0001012213033031-2203300210332333-0302231323031000-1022010303211022"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CDNPurgeCommand.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-1003131013121223-0332023003220032-3022202203233103-2003100031213200-2033002123001211-1213330131031020-1010213010312013-0201320013312023): complete subsection reference.

<a id="canonical-1203123000321113-0232220213223203-0110233121133033-0133233313120302-1030200102203221-2303003120223133-0121133113303323-2013201310200210"></a>

<a id="canonical-1301333111312001-3103330210333233-3331121121212022-2001113332020202-3313333313021002-3321012310000210-2032003231013233-2001133210320003"></a>

#### `hostname` property

Type: `"string"`. Computed.

\[OneOf: hostname, pattern, purge\_all, URL\_path\] Exclusive with \[pattern purge\_all URL\_path\]
Purge cached content by Hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

OneOf alternatives in this subsection:

- [hostname](data-sources--cdn_purge_command--reference--group-001.md#canonical-1203123000321113-0232220213223203-0110233121133033-0133233313120302-1030200102203221-2303003120223133-0121133113303323-2013201310200210)
- [pattern](data-sources--cdn_purge_command--reference--group-001.md#canonical-3320323231133112-1113213000012212-1312032311011003-2021202303222211-2302101221233313-2122031301130123-1002301132213112-0131123122021022)
- [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-0101013120130302-3112113231223113-3133223121013212-3321111032030012-3023012211103012-1211321300021000-2021020303122120-3000222031323002)
- [url_path](data-sources--cdn_purge_command--reference--group-001.md#canonical-3333000311312133-3101223313120011-1231311101102232-2103130330102331-2303212130230002-1120001333302121-2113332203112201-2023110110100200)

Select alternatives according to the provider validators above.

<a id="canonical-3103113021211320-1103120233103132-3231310212201011-0221302333323002-2302301312103213-3222132123302130-3111110003002111-0121302020222231"></a>

<a id="canonical-2022222233321222-1132333121100333-2033133013020230-1130332233200020-1300000322323322-1302100033133133-0333212120010302-3213301211320003"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3301020130001131-0300311310122112-0020103031230100-2023212323000131-2230303033032210-3203310013120001-2113301033110030-0312211010232102"></a>

<a id="canonical-0303000021020211-3302031200201121-3312230121220233-1022022202003320-3222211313331200-3122123132322120-2033233301011222-3020133031103202"></a>

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

<a id="canonical-1222030001200313-3311010312032022-2120011030130112-1123121300003012-3323233033333023-3030131031113200-1030332133022032-0331313020133302"></a>

<a id="canonical-0222122010301312-0323232000030131-0132010332112311-0302023321100232-3033210030121032-0121003100021133-2220023233301300-2031131300001120"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CDNPurgeCommand.

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

<a id="canonical-1301321012001333-1211100103010232-1313323232200032-0003033022213222-0322320113022021-3002312111030211-3300213322002230-2230333113002312"></a>

<a id="canonical-2020103311000011-0321103323132312-2322032223302122-2222213332222213-2300032022122322-0303223203130010-1303303200133110-1300110232320233"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CDNPurgeCommand exists.

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
  }
}
```

<a id="canonical-3320323231133112-1113213000012212-1312032311011003-2021202303222211-2302101221233313-2122031301130123-1002301132213112-0131123122021022"></a>

<a id="canonical-0203011300100232-0113203112132121-3021031120210211-2220312320202023-0031220012323323-2020232231331220-1020311222032333-2222332223312302"></a>

#### `pattern` property

Type: `"string"`. Computed.

Exclusive with \[hostname purge\_all URL\_path\] Purge cached content using PCRE 1 compliant regular
expression.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-1110200333022131-0022021320031113-3022010120211113-0030122203213222-3030300331111322-0000130302332111-3212101112130203-3020031303213210): complete subsection reference.

- [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-3323103323002133-2030122033201133-2130111320202220-3303103332223211-2122030310101013-3313333310133332-0100331122032233-2111001123100310): complete subsection reference.

<a id="canonical-3333000311312133-3101223313120011-1231311101102232-2103130330102331-2303212130230002-1120001333302121-2113332203112201-2023110110100200"></a>

<a id="canonical-3320313030131203-3203030001223312-3032103001011322-1200201111300313-1233232111121133-0013032223030301-1013022321110113-0210202111000111"></a>

#### `url_path` property

Type: `"string"`. Computed.

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [virtual_host](data-sources--cdn_purge_command--reference--group-001.md#canonical-2330132311010220-0012200333222002-0333103013030112-1123120311123102-1012331100303020-0321113100133132-3230121033002032-0013030210133122): complete subsection reference.

<a id="canonical-2311102222013213-0131211021322332-2133203010120322-2002200312000313-2022201012112211-1133333331232100-1021330100110002-1010212310303112"></a>

### All schema paths for `xcsh_cdn_purge_command`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cdn_purge_command--reference--group-001.md#canonical-1002331010331233-2331120200100110-3212121232023131-2232132210121312-3230330330010020-3221122131112001-0101003333121032-1211001122202220) |
| `description` | [description](data-sources--cdn_purge_command--reference--group-001.md#canonical-2231021111330102-3332322133330003-2003212102301323-2103002102033300-2133102000233333-2012233021202112-1032200031132112-3212101132013030) |
| `hard_purge` | [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-3013110031232130-0113321303320223-3010213111322203-3032002233131322-2222001301300212-3122303113002020-2023312003013320-0321022103003323) |
| `hostname` | [hostname](data-sources--cdn_purge_command--reference--group-001.md#canonical-1203123000321113-0232220213223203-0110233121133033-0133233313120302-1030200102203221-2303003120223133-0121133113303323-2013201310200210) |
| `id` | [ID](data-sources--cdn_purge_command--reference--group-001.md#canonical-3103113021211320-1103120233103132-3231310212201011-0221302333323002-2302301312103213-3222132123302130-3111110003002111-0121302020222231) |
| `labels` | [labels](data-sources--cdn_purge_command--reference--group-001.md#canonical-3301020130001131-0300311310122112-0020103031230100-2023212323000131-2230303033032210-3203310013120001-2113301033110030-0312211010232102) |
| `name` | [name](data-sources--cdn_purge_command--reference--group-001.md#canonical-1222030001200313-3311010312032022-2120011030130112-1123121300003012-3323233033333023-3030131031113200-1030332133022032-0331313020133302) |
| `namespace` | [namespace](data-sources--cdn_purge_command--reference--group-001.md#canonical-1301321012001333-1211100103010232-1313323232200032-0003033022213222-0322320113022021-3002312111030211-3300213322002230-2230333113002312) |
| `pattern` | [pattern](data-sources--cdn_purge_command--reference--group-001.md#canonical-3320323231133112-1113213000012212-1312032311011003-2021202303222211-2302101221233313-2122031301130123-1002301132213112-0131123122021022) |
| `purge_all` | [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-0101013120130302-3112113231223113-3133223121013212-3321111032030012-3023012211103012-1211321300021000-2021020303122120-3000222031323002) |
| `soft_purge` | [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-2311232102113122-2013322031012210-0330210313120131-3120200023300110-3133033033123022-3131100100232313-3213123213223113-0012303012000012) |
| `url_path` | [url_path](data-sources--cdn_purge_command--reference--group-001.md#canonical-3333000311312133-3101223313120011-1231311101102232-2103130330102331-2303212130230002-1120001333302121-2113332203112201-2023110110100200) |
| `virtual_host` | [virtual_host](data-sources--cdn_purge_command--reference--group-001.md#canonical-0202233113221302-1010302103120100-0222020223213110-2112123132032332-3312203310310321-0313122210001010-3333122231130222-1020311232201221) |
| `virtual_host.name` | [virtual_host.name](data-sources--cdn_purge_command--reference--group-001.md#canonical-3011301222232301-1213213132132332-3210000100321200-3303102032333030-2032222210021312-2002313102321002-3033132310323333-0023202222212303) |
| `virtual_host.namespace` | [virtual_host.namespace](data-sources--cdn_purge_command--reference--group-001.md#canonical-2221232132013111-2301010133111333-1202330210222220-2010010330122211-0031033302110300-3130210330102000-1312320030321130-2000111030222211) |
| `virtual_host.tenant` | [virtual_host.tenant](data-sources--cdn_purge_command--reference--group-001.md#canonical-2332221201213201-3031331232120100-3132122223110001-2212220131010200-1012133000333220-0013213022131321-0211312131003203-0113033311233231) |

<a id="canonical-1003131013121223-0332023003220032-3022202203233103-2003100031213200-2033002123001211-1213330131031020-1010213010312013-0201320013312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hard_purge` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032)
- hard_purge

<a id="canonical-3013110031232130-0113321303320223-3010213111322203-3032002233131322-2222001301300212-3122303113002020-2023312003013320-0321022103003323"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hard\_purge, soft\_purge\] Enable this option

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

OneOf alternatives in this subsection:

- [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-3013110031232130-0113321303320223-3010213111322203-3032002233131322-2222001301300212-3122303113002020-2023312003013320-0321022103003323)
- [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-2311232102113122-2013322031012210-0330210313120131-3120200023300110-3133033033123022-3131100100232313-3213123213223113-0012303012000012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110200333022131-0022021320031113-3022010120211113-0030122203213222-3030300331111322-0000130302332111-3212101112130203-3020031303213210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `purge_all` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032)
- purge_all

<a id="canonical-0101013120130302-3112113231223113-3133223121013212-3321111032030012-3023012211103012-1211321300021000-2021020303122120-3000222031323002"></a>

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

<a id="canonical-3323103323002133-2030122033201133-2130111320202220-3303103332223211-2122030310101013-3313333310133332-0100331122032233-2111001123100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `soft_purge` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032)
- soft_purge

<a id="canonical-2311232102113122-2013322031012210-0330210313120131-3120200023300110-3133033033123022-3131100100232313-3213123213223113-0012303012000012"></a>

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

<a id="canonical-2330132311010220-0012200333222002-0333103013030112-1123120311123102-1012331100303020-0321113100133132-3230121033002032-0013030210133122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_host` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032)
- virtual_host

<a id="canonical-0202233113221302-1010302103120100-0222020223213110-2112123132032332-3312203310310321-0313122210001010-3333122231130222-1020311232201221"></a>

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

<a id="canonical-2033233110223333-1320021100213033-1202022120013022-1101113221021033-1000202000300132-0022213332132103-0103231320103221-2312001111213233"></a>

### Direct properties for `virtual_host`

<a id="canonical-3011301222232301-1213213132132332-3210000100321200-3303102032333030-2032222210021312-2002313102321002-3033132310323333-0023202222212303"></a>

#### `virtual_host.name` property

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

<a id="canonical-2221232132013111-2301010133111333-1202330210222220-2010010330122211-0031033302110300-3130210330102000-1312320030321130-2000111030222211"></a>

<a id="canonical-2132322121200100-3131121123310010-0111133223100321-1022103023033102-0302233233322111-0113321312123131-2213000032100021-3203233131220303"></a>

#### `virtual_host.namespace` property

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

<a id="canonical-2332221201213201-3031331232120100-3132122223110001-2212220131010200-1012133000333220-0013213022131321-0211312131003203-0113033311233231"></a>

<a id="canonical-0333202130333023-1220313123113323-0302231330213312-1313232221302032-2113002101231322-1130112331133111-2133301211001102-2211121010311132"></a>

#### `virtual_host.tenant` property

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
