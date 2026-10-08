---
page_title: "xcsh_cdn_purge_command reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command reference."
---

# xcsh_cdn_purge_command reference

<a id="canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- Property reference

<a id="canonical-2021320002132210-2131310120303220-2303213113203303-0013003131102333-3301100113331322-0303222120131030-1230110022323330-3210112332330101"></a>

### Direct properties for `xcsh_cdn_purge_command`

<a id="canonical-2000013102312130-0020213120311202-3302020333112101-0001312033201102-3220011311303220-3033231310110022-1203030320211300-2023211132213101"></a>

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

<a id="canonical-0100231323221023-3131112021303023-0322321232202222-3120302003200232-3201200320003320-3130200022111122-0310110330011310-2301001021113323"></a>

<a id="canonical-0120122110110003-2312220000012212-1012131112220320-1230201021002010-3313123303302232-2233320033221011-0100221132000012-3013131221132203"></a>

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

<a id="canonical-2210200011022213-2312203301133012-2230310031212310-2303003200222131-1213210113203001-0310130102101211-1321002021110033-3201320321313202"></a>

<a id="canonical-3321303113021031-2210001213232201-2023220030223302-2033131302322203-0102332033201012-0333302021013033-0110013030000200-0301003310311331"></a>

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

- [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-1131033021013223-2101330132310303-1131303103232322-3010100202300132-2222302013012332-0020231111311021-0230211131321232-1221221210312303): complete subsection reference.

<a id="canonical-1131201031203333-1100022213212031-1020011103022001-0110220110032311-1033002202021121-1210120111302333-0110103110130310-0213332333210012"></a>

<a id="canonical-2103212332211000-2303102123113231-1311101322121032-0123003311223030-0031202001000100-3202130303023311-1123330123033023-2132321212120332"></a>

#### `hostname` property

Type: `"string"`. Optional, Computed.

\[OneOf: hostname, pattern, purge\_all, URL\_path\] Exclusive with \[pattern purge\_all URL\_path\]
Purge cached content by Hostname.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [hostname](resources--cdn_purge_command--reference--group-001.md#canonical-1131201031203333-1100022213212031-1020011103022001-0110220110032311-1033002202021121-1210120111302333-0110103110130310-0213332333210012)
- [pattern](resources--cdn_purge_command--reference--group-001.md#canonical-0103203032111020-2032302310123332-2020133202101001-0010101011231130-1302113012102022-2213000300002211-2000223312220333-3013233203321301)
- [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-1102001032321011-3002220122233001-2002112122233131-3013102201012200-1001011002020201-3132031020120213-0302012321232103-3333033112002132)
- [url_path](resources--cdn_purge_command--reference--group-001.md#canonical-1220323320123013-0101020301321230-1021110320003001-0330120212133331-1230011233031013-2233121210202311-3100033321332332-3332101322211032)

Select alternatives according to the provider validators above.

<a id="canonical-1330111122311132-1333311022202333-1321022102312210-3331332213300020-0211032003212231-1123313232233100-3031332233130202-3210122201033031"></a>

<a id="canonical-2220123313122001-1210233330121121-1101323030003231-1200030202202331-3123232312122330-2120321321030030-2301302213112200-3231003000330331"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3010223310010200-1010323200130110-0103022003002121-1103320021003300-1331200002222313-1233202333220101-2012022203221013-1332023132122021"></a>

<a id="canonical-2111230122023322-1102001011233023-0230100003010013-3000120323223213-0101311021221221-2010212312321013-2100232302220333-2000320122300313"></a>

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

<a id="canonical-1030111332320120-1321322332003022-3313101301030013-2301102030110033-2300331010302121-3312331031221032-3221201202021201-1301032231103323"></a>

<a id="canonical-2302030022130030-0013010022202331-2102313120121121-3301110302100221-0313221130131232-1221113023103130-2101101313002110-0330220213011332"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CDN Purge Command. Must be unique within the namespace.

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

<a id="canonical-2323322132310023-2331223113223022-2111112112132311-0232022020020210-1131302013202130-3332220121022220-2101012323301322-3000121001003131"></a>

<a id="canonical-3032213133312230-3312303030312101-0231032112003222-0333312130120203-1113021113022122-3120111102102110-0013303003101023-2311300301123032"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CDN Purge Command is created.

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

<a id="canonical-0103203032111020-2032302310123332-2020133202101001-0010101011231130-1302113012102022-2213000300002211-2000223312220333-3013233203321301"></a>

<a id="canonical-3100332231113110-3213030200022201-3230130210113110-3303021102021333-2203031301221231-2122112301302220-1302301222313003-2313102123211022"></a>

#### `pattern` property

Type: `"string"`. Optional, Computed.

Exclusive with \[hostname purge\_all URL\_path\] Purge cached content using PCRE 1 compliant regular
expression.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-1123320230111100-2132200320220211-2020231302000131-1312020213123222-0321132332223200-2232110213232300-0230310022012131-3223130101212223): complete subsection reference.

- [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-3103210211233000-3133311121313023-1110013322312202-1100101203003031-0300110300130301-1000112300111030-1012220312010020-2031321020131012): complete subsection reference.

- [timeouts](resources--cdn_purge_command--reference--group-001.md#canonical-3232331031333302-1112023301312211-1132010221321321-3132132013111311-2121000023123300-0131103100300301-0233031113121001-1101112133120202): complete subsection reference.

<a id="canonical-1220323320123013-0101020301321230-1021110320003001-0330120212133331-1230011233031013-2233121210202311-3100033321332332-3332101322211032"></a>

<a id="canonical-2323213130221033-2320303103213021-1221322230213200-0131333213012210-2030210233011300-1103100101103222-3310201000230003-3322003133230020"></a>

#### `url_path` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [virtual_host](resources--cdn_purge_command--reference--group-001.md#canonical-2102032030131310-1103132322002331-1331100112012100-1031113102213333-0032231331203322-3020321120331113-2031202100120212-0032011030323102): complete subsection reference.

<a id="canonical-1030102210331211-1101123212123231-2233110100131103-0020113220203030-2222231000112233-1123221033222232-0032301303311010-1020223003000112"></a>

### All schema paths for `xcsh_cdn_purge_command`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_purge_command--reference--group-001.md#canonical-2000013102312130-0020213120311202-3302020333112101-0001312033201102-3220011311303220-3033231310110022-1203030320211300-2023211132213101) |
| `description` | [description](resources--cdn_purge_command--reference--group-001.md#canonical-0100231323221023-3131112021303023-0322321232202222-3120302003200232-3201200320003320-3130200022111122-0310110330011310-2301001021113323) |
| `disable` | [disable](resources--cdn_purge_command--reference--group-001.md#canonical-2210200011022213-2312203301133012-2230310031212310-2303003200222131-1213210113203001-0310130102101211-1321002021110033-3201320321313202) |
| `hard_purge` | [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-2021302013113120-0212320231331120-0220130003011203-0212112313021232-0021222123132320-3300002200021100-0210221110000101-2112330112000302) |
| `hostname` | [hostname](resources--cdn_purge_command--reference--group-001.md#canonical-1131201031203333-1100022213212031-1020011103022001-0110220110032311-1033002202021121-1210120111302333-0110103110130310-0213332333210012) |
| `id` | [ID](resources--cdn_purge_command--reference--group-001.md#canonical-1330111122311132-1333311022202333-1321022102312210-3331332213300020-0211032003212231-1123313232233100-3031332233130202-3210122201033031) |
| `labels` | [labels](resources--cdn_purge_command--reference--group-001.md#canonical-3010223310010200-1010323200130110-0103022003002121-1103320021003300-1331200002222313-1233202333220101-2012022203221013-1332023132122021) |
| `name` | [name](resources--cdn_purge_command--reference--group-001.md#canonical-1030111332320120-1321322332003022-3313101301030013-2301102030110033-2300331010302121-3312331031221032-3221201202021201-1301032231103323) |
| `namespace` | [namespace](resources--cdn_purge_command--reference--group-001.md#canonical-2323322132310023-2331223113223022-2111112112132311-0232022020020210-1131302013202130-3332220121022220-2101012323301322-3000121001003131) |
| `pattern` | [pattern](resources--cdn_purge_command--reference--group-001.md#canonical-0103203032111020-2032302310123332-2020133202101001-0010101011231130-1302113012102022-2213000300002211-2000223312220333-3013233203321301) |
| `purge_all` | [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-1102001032321011-3002220122233001-2002112122233131-3013102201012200-1001011002020201-3132031020120213-0302012321232103-3333033112002132) |
| `soft_purge` | [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-3213202230321301-2330002312311113-0300102322223123-0313331020231120-3001322223022313-3313311021033113-1201011221303313-2023101013113110) |
| `timeouts` | [timeouts](resources--cdn_purge_command--reference--group-001.md#canonical-2321210331212222-1311111211212101-0301201321111021-1322201230123121-1332101203013031-0022213221023203-0313002330320212-1130211313011023) |
| `timeouts.create` | [timeouts.create](resources--cdn_purge_command--reference--group-001.md#canonical-3332231122131011-3223310310031122-3213130211101312-0022210003002133-3323010210232313-1102020202121010-3121222110130223-3230110333011302) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_purge_command--reference--group-001.md#canonical-2323312020012122-0300231203132102-3213031131022113-1130211320120330-0213011311031222-0013211321310320-2303220223012010-3223303302303111) |
| `timeouts.read` | [timeouts.read](resources--cdn_purge_command--reference--group-001.md#canonical-3223123112023323-0303310223312122-2022102031101200-0110120313220030-3002002120100201-2221211210101032-1021002013133221-1231013102301202) |
| `timeouts.update` | [timeouts.update](resources--cdn_purge_command--reference--group-001.md#canonical-1320110123310302-1010132313111000-0200001103122022-0213230200233101-3122002321030002-1120111312020020-1122301313201320-2130322321212230) |
| `url_path` | [url_path](resources--cdn_purge_command--reference--group-001.md#canonical-1220323320123013-0101020301321230-1021110320003001-0330120212133331-1230011233031013-2233121210202311-3100033321332332-3332101322211032) |
| `virtual_host` | [virtual_host](resources--cdn_purge_command--reference--group-001.md#canonical-0111321331011120-0330323012203221-1223312133032132-0301200320213200-1230302222231010-3030303111212302-3122201101122333-2032310021133131) |
| `virtual_host.name` | [virtual_host.name](resources--cdn_purge_command--reference--group-001.md#canonical-1222030212023030-1131123200022322-3023313123233202-1302323032103200-0301030011202333-2212213313023213-3022133221313030-0020133121221001) |
| `virtual_host.namespace` | [virtual_host.namespace](resources--cdn_purge_command--reference--group-001.md#canonical-2130131310300303-2202200233301102-2331000223121102-0000033212201203-0333133022313032-0312221122212103-3323202011003320-2301002113320003) |
| `virtual_host.tenant` | [virtual_host.tenant](resources--cdn_purge_command--reference--group-001.md#canonical-3222010202303313-0332132220132122-2120123230120212-0233101232200301-0120311302002313-0233012320212010-2220003100023012-1111103011022212) |

<a id="canonical-1131033021013223-2101330132310303-1131303103232322-3010100202300132-2222302013012332-0020231111311021-0230211131321232-1221221210312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hard_purge` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- hard_purge

<a id="canonical-2021302013113120-0212320231331120-0220130003011203-0212112313021232-0021222123132320-3300002200021100-0210221110000101-2112330112000302"></a>

Type: `["object", {}]`. Optional.

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

- [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-2021302013113120-0212320231331120-0220130003011203-0212112313021232-0021222123132320-3300002200021100-0210221110000101-2112330112000302)
- [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-3213202230321301-2330002312311113-0300102322223123-0313331020231120-3001322223022313-3313311021033113-1201011221303313-2023101013113110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hard_purge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123320230111100-2132200320220211-2020231302000131-1312020213123222-0321132332223200-2232110213232300-0230310022012131-3223130101212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `purge_all` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- purge_all

<a id="canonical-1102001032321011-3002220122233001-2002112122233131-3013102201012200-1001011002020201-3132031020120213-0302012321232103-3333033112002132"></a>

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
purge_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103210211233000-3133311121313023-1110013322312202-1100101203003031-0300110300130301-1000112300111030-1012220312010020-2031321020131012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `soft_purge` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- soft_purge

<a id="canonical-3213202230321301-2330002312311113-0300102322223123-0313331020231120-3001322223022313-3313311021033113-1201011221303313-2023101013113110"></a>

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
soft_purge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232331031333302-1112023301312211-1132010221321321-3132132013111311-2121000023123300-0131103100300301-0233031113121001-1101112133120202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- timeouts

<a id="canonical-2321210331212222-1311111211212101-0301201321111021-1322201230123121-1332101203013031-0022213221023203-0313002330320212-1130211313011023"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300320211003103-3022011121231120-1212320300220323-0023103121003021-2310223302100203-2320020202311330-0321132003112332-0203020121321310"></a>

### Direct properties for `timeouts`

<a id="canonical-3332231122131011-3223310310031122-3213130211101312-0022210003002133-3323010210232313-1102020202121010-3121222110130223-3230110333011302"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2323312020012122-0300231203132102-3213031131022113-1130211320120330-0213011311031222-0013211321310320-2303220223012010-3223303302303111"></a>

<a id="canonical-0100002130131123-2020303303113003-3201102121132220-2332001123033033-3223023121332300-2320103113131010-0120101213000002-3112303111301200"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3223123112023323-0303310223312122-2022102031101200-0110120313220030-3002002120100201-2221211210101032-1021002013133221-1231013102301202"></a>

<a id="canonical-2121030113003320-1330121030121220-3121301100331213-0222013103121103-2031130301012132-2120223232122120-0313213222332023-2203210012002202"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1320110123310302-1010132313111000-0200001103122022-0213230200233101-3122002321030002-1120111312020020-1122301313201320-2130322321212230"></a>

<a id="canonical-0022022232000133-1231102100011022-1221112320320310-3213010300010221-1221113013031010-1013112222121102-1333332012320201-1203302320120231"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2102032030131310-1103132322002331-1331100112012100-1031113102213333-0032231331203322-3020321120331113-2031202100120212-0032011030323102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_host` properties

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- virtual_host

<a id="canonical-0111321331011120-0330323012203221-1223312133032132-0301200320213200-1230302222231010-3030303111212302-3122201101122333-2032310021133131"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
virtual_host {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230110000323111-3023023220322333-3322322013133230-3021203233200203-0233012331020113-3213212220132201-1323011010310103-1312200001013112"></a>

### Direct properties for `virtual_host`

<a id="canonical-1222030212023030-1131123200022322-3023313123233202-1302323032103200-0301030011202333-2212213313023213-3022133221313030-0020133121221001"></a>

#### `virtual_host.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2130131310300303-2202200233301102-2331000223121102-0000033212201203-0333133022313032-0312221122212103-3323202011003320-2301002113320003"></a>

<a id="canonical-2332313010013112-1013330230233021-2331011131100111-3003302300001031-2333310312103121-1220030210030022-3130001321200312-0110121330032022"></a>

#### `virtual_host.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3222010202303313-0332132220132122-2120123230120212-0233101232200301-0120311302002313-0233012320212010-2220003100023012-1111103011022212"></a>

<a id="canonical-2000000300010211-1220213023133131-1213200023030331-0313233002123122-2021133130231330-3212303330203013-0213323200110030-1233113103231213"></a>

#### `virtual_host.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```
