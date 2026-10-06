---
page_title: "xcsh_irule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule reference."
---

# xcsh_irule reference

<a id="canonical-1022201003222120-3100001332021300-1113323102020013-3212101111010201-3331213101333303-0013033103232213-3130321213220101-2011220222210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203)
- Property reference

<a id="canonical-1031201221212013-0013131332001232-1030001032310301-3030212323103202-3123023203311002-1132321321220122-3301021200302113-0101132232213022"></a>

### Direct properties for `xcsh_irule`

<a id="canonical-3012130310201232-0221202323003313-3031233100131032-3210010033331300-2033232222030310-1322213102323000-2320001321031200-2330302220303230"></a>

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

<a id="canonical-2320231200002313-3202013331203022-0023213023003322-1132131030333002-0200210222220202-3323202021323212-2100120223231303-3213001020231200"></a>

<a id="canonical-0212222112233013-3012123020332312-3332323121310021-2300330331311022-0002121230302120-2213212110033101-3001101332121230-1222322331211310"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Additional upstream details:

Specify Description for iRule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 0
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

<a id="canonical-2101031333000213-0022133001122002-1011200121332100-0332200022201030-0312313230020333-0000111202100200-1111201133201031-3232120101321323"></a>

<a id="canonical-1213112210230012-1201132001003021-2020132323302202-2333122012320210-1010212222211021-1000122321220311-0300003020101333-0323023131323232"></a>

#### `description_spec` property

Type: `"string"`. Required.

Description for iRule. Specify Description for iRule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0303301122201130-0000333322032110-2113311002221012-3213020302100033-3222203123313220-3332211322030322-1322021213331232-2323233213300123"></a>

<a id="canonical-0213201333102000-2200003122110302-0300032011222233-1312123010100122-0020230121010121-1001321311332321-1330021110222311-2323012323231333"></a>

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

<a id="canonical-0211013103020302-2013021100200321-1212312101033111-0211320033200133-3220211102323023-1130201130033322-3131302233132211-1021102001332233"></a>

<a id="canonical-2001022031301312-2010113310213031-3002331203002111-3222203311130330-3100122023030321-3032233223102200-3131320210222213-1023100202202133"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3201220032331202-3131312131121010-0101200213231201-1222002212300333-0031201030333332-2013303201003323-0221100102020323-2212320100320332"></a>

<a id="canonical-2232202121021310-2031203132012223-0203300223212220-3033010121213210-1132333300111323-1023300111020033-2333313102012010-2013000031202102"></a>

#### `irule` property

Type: `"string"`. Required.

www&#46;internal.example.f5.com')\} DNS::drop\} irule content.

Additional upstream details:

www&#46;internal.example.f5.com")\} DNS::drop\} irule content.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(24576),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 24576,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 24576,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "24576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "24576"
  }
}
```

<a id="canonical-1112312323113210-1122122303212101-3203103123023232-1010202212323301-3301210021100320-1023102213300032-3333101022111312-3310333321023321"></a>

<a id="canonical-2210310310313032-1303311302223330-3311212200302132-0322212201202330-1122231103213300-1133210220002030-3003321002313303-2223231332320130"></a>

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

<a id="canonical-3113212233313332-2332112101312120-3230221222110202-2310232213332100-0003110013033121-3012220103220213-2230002033031312-2123232321033210"></a>

<a id="canonical-1022112312012103-3101111032203212-1231313332130132-2221030223202321-1013001302322321-3312310330123211-3322210102030103-3223121110213312"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Irule. Must be unique within the namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0313033311113110-3320212022002000-2103133103311313-1312023111013132-3223122122010300-0213022221212132-1122112002310222-3303033232312313"></a>

<a id="canonical-1031111033001013-3210300331213320-2030212313300013-3232211231223113-3110121131331012-3321331333133132-3032103132103331-3120030223331103"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Irule is created.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [timeouts](resources--irule--reference--group-001.md#canonical-0031111210001221-1233022110020210-2022201202303023-3321313323310310-1323310012031320-0122202302102210-0000030321013110-1110200301110212): complete subsection reference.

<a id="canonical-2223203210301320-3300022010022321-0023320022102030-3322033211033223-3131122122130220-0013131203013113-0110102231111112-0031020013122331"></a>

### All schema paths for `xcsh_irule`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--irule--reference--group-001.md#canonical-3012130310201232-0221202323003313-3031233100131032-3210010033331300-2033232222030310-1322213102323000-2320001321031200-2330302220303230) |
| `description` | [description](resources--irule--reference--group-001.md#canonical-2320231200002313-3202013331203022-0023213023003322-1132131030333002-0200210222220202-3323202021323212-2100120223231303-3213001020231200) |
| `description_spec` | [description_spec](resources--irule--reference--group-001.md#canonical-2101031333000213-0022133001122002-1011200121332100-0332200022201030-0312313230020333-0000111202100200-1111201133201031-3232120101321323) |
| `disable` | [disable](resources--irule--reference--group-001.md#canonical-0303301122201130-0000333322032110-2113311002221012-3213020302100033-3222203123313220-3332211322030322-1322021213331232-2323233213300123) |
| `id` | [ID](resources--irule--reference--group-001.md#canonical-0211013103020302-2013021100200321-1212312101033111-0211320033200133-3220211102323023-1130201130033322-3131302233132211-1021102001332233) |
| `irule` | [irule](resources--irule--reference--group-001.md#canonical-3201220032331202-3131312131121010-0101200213231201-1222002212300333-0031201030333332-2013303201003323-0221100102020323-2212320100320332) |
| `labels` | [labels](resources--irule--reference--group-001.md#canonical-1112312323113210-1122122303212101-3203103123023232-1010202212323301-3301210021100320-1023102213300032-3333101022111312-3310333321023321) |
| `name` | [name](resources--irule--reference--group-001.md#canonical-3113212233313332-2332112101312120-3230221222110202-2310232213332100-0003110013033121-3012220103220213-2230002033031312-2123232321033210) |
| `namespace` | [namespace](resources--irule--reference--group-001.md#canonical-0313033311113110-3320212022002000-2103133103311313-1312023111013132-3223122122010300-0213022221212132-1122112002310222-3303033232312313) |
| `timeouts` | [timeouts](resources--irule--reference--group-001.md#canonical-3100310312301302-1233220001332200-1310231331131110-2201111320330111-0110201102020202-1232301133032210-3301011222313003-2301121011200010) |
| `timeouts.create` | [timeouts.create](resources--irule--reference--group-001.md#canonical-2111321111120020-3201300110133230-0011120321213313-2200311123102223-0333033300100323-1030021211013122-2102010122033310-3202231132102331) |
| `timeouts.delete` | [timeouts.delete](resources--irule--reference--group-001.md#canonical-2033032123203300-3020203032100322-3133102232012100-1222312330311103-2231032302231231-2033311323331112-0031113330001230-2123130130011223) |
| `timeouts.read` | [timeouts.read](resources--irule--reference--group-001.md#canonical-1002230220231210-0130300031221101-1301102232002232-1000002202313031-1101213103301113-3323121030121311-2222201201031232-0333121213111223) |
| `timeouts.update` | [timeouts.update](resources--irule--reference--group-001.md#canonical-1020110203010033-0131310011320003-0300001310200113-1131101132323103-0221001202032012-0222333011233211-0001122023330001-1000312330102123) |

<a id="canonical-0031111210001221-1233022110020210-2022201202303023-3321313323310310-1323310012031320-0122202302102210-0000030321013110-1110200301110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203)
- [Property reference](resources--irule--reference--group-001.md#canonical-1022201003222120-3100001332021300-1113323102020013-3212101111010201-3331213101333303-0013033103232213-3130321213220101-2011220222210011)
- timeouts

<a id="canonical-3100310312301302-1233220001332200-1310231331131110-2201111320330111-0110201102020202-1232301133032210-3301011222313003-2301121011200010"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303301323300022-0122132123332203-0022132322223033-2312211222031202-2203303310203122-3031023121200301-2330233313232311-1001033310122300"></a>

### Direct properties for `timeouts`

<a id="canonical-2111321111120020-3201300110133230-0011120321213313-2200311123102223-0333033300100323-1030021211013122-2102010122033310-3202231132102331"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2033032123203300-3020203032100322-3133102232012100-1222312330311103-2231032302231231-2033311323331112-0031113330001230-2123130130011223"></a>

<a id="canonical-2203323200230030-0013001022103011-0103311033032201-1113212020332331-0322112113212221-2232102010122012-3302201033101013-0231333013031221"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1002230220231210-0130300031221101-1301102232002232-1000002202313031-1101213103301113-3323121030121311-2222201201031232-0333121213111223"></a>

<a id="canonical-3231301213100320-2300233230333213-3221222022212032-0202313231120223-1021221112222031-0100103001110032-0132220311131032-3221121130323101"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1020110203010033-0131310011320003-0300001310200113-1131101132323103-0221001202032012-0222333011233211-0001122023330001-1000312330102123"></a>

<a id="canonical-3132120121201022-2121303202132203-3113020100020122-0022030020210303-3112101332233001-3120221200032031-3012223232302312-3212120111002332"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
