---
page_title: "xcsh_api_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery reference."
---

# xcsh_api_discovery reference

<a id="canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- Property reference

<a id="canonical-2220032303211100-2331103030112112-2302000100231130-2302001313111320-0021020131301103-1313303301010311-3012020223132211-0010011021000201"></a>

### Direct properties for `xcsh_api_discovery`

<a id="canonical-2323022111020301-1130100200212231-3222000201032231-2223112112000231-1111313023102311-1333311302201200-3330110103312123-0000020021112131"></a>

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

- [custom_auth_types](resources--api_discovery--reference--group-001.md#canonical-2212121133030123-3313103211030133-1332101312033102-0023310100201330-1131130101331031-0031112010301020-0220030121333230-3300021031132021): complete subsection reference.

<a id="canonical-2201033011130210-0030020110122023-1232223301000311-2133031303323230-3010031221301303-1330220133212203-1201321100203232-3301232231232301"></a>

<a id="canonical-3102011311021011-1322000311111201-0102300000011232-3231020223211202-2301213201100120-1133132021121232-1332330221300213-1300133203012332"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2113321033211313-0302322101211013-2121112302232320-3132333301103123-1012101012123313-3103222313020131-2300231201221310-3331030310012010"></a>

<a id="canonical-3330323333133310-1033310120331223-3220200210010303-1133302010132200-3122311200213331-0333233110120300-2201223033002021-2310232302231100"></a>

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

<a id="canonical-0223101103121130-1000303133022330-2032113202332131-3213132033011200-3300100333202000-3100303100022110-0023312301021032-1301331201133203"></a>

<a id="canonical-1002322003310331-1203313201222210-3333123003030123-2002001102200133-1210022021202313-0031131201212210-1231202013033123-3122231323132103"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2301220123231220-0321032332312221-0203231032312313-3112223033220330-1001333203001311-1312022203030311-3031303120232103-1212021302131031"></a>

<a id="canonical-1012332212120011-3310021130013230-3200332203211200-1330023332222122-1323302003322120-1203221331000223-3333203203320022-0010013022120212"></a>

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

<a id="canonical-3112323020113233-3003113311120032-1001130113302133-1022002021331233-0101321003203303-2302131323233200-0001231311230102-1321330013010010"></a>

<a id="canonical-0230110033212033-2013132012200322-3220001103001312-3300332022020201-2322311301120013-1210303330130130-1233322323103003-3201121310201312"></a>

#### `name` property

Type: `"string"`. Required.

Name of the API Discovery. Must be unique within the namespace.

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

<a id="canonical-3122132311311221-1130300031011330-2120303310331302-2313030202313311-1033003023303131-3131102130132211-3302332103102303-2330213231322010"></a>

<a id="canonical-2133013001013233-2303303000312100-2220220122303233-2230331221331113-1231020323031312-1202223230223032-3020222002303033-1200023210123313"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the API Discovery is created.

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

- [timeouts](resources--api_discovery--reference--group-001.md#canonical-0330002233322121-1233033110301003-0230030201120323-0111002303003220-3301031130232123-2322330011312000-2312101022131231-3301102100120201): complete subsection reference.

- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122): complete subsection reference.

<a id="canonical-1131320021320000-2333012230121302-3011331321100001-0300013130023212-0332200230132323-2232021211132303-1300200310100120-0231203112213220"></a>

### All schema paths for `xcsh_api_discovery`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_discovery--reference--group-001.md#canonical-2323022111020301-1130100200212231-3222000201032231-2223112112000231-1111313023102311-1333311302201200-3330110103312123-0000020021112131) |
| `custom_auth_types` | [custom_auth_types](resources--api_discovery--reference--group-001.md#canonical-3320211132230310-2323120333312332-3223131100221213-3032122322223133-1122201100032012-3100100100313020-0332212122311231-2033033132030223) |
| `custom_auth_types.parameter_name` | [custom_auth_types.parameter_name](resources--api_discovery--reference--group-001.md#canonical-1311030212221033-3231310211331320-2303312021123213-0011102233101202-3200111011220023-0300021102130303-2102333131222311-0002013012221112) |
| `custom_auth_types.parameter_type` | [custom_auth_types.parameter_type](resources--api_discovery--reference--group-001.md#canonical-0230331323213212-1010102133233110-3131013302101333-1011230001232002-2022230222033120-0010022132111113-2122222032201312-3233323112301102) |
| `description` | [description](resources--api_discovery--reference--group-001.md#canonical-2201033011130210-0030020110122023-1232223301000311-2133031303323230-3010031221301303-1330220133212203-1201321100203232-3301232231232301) |
| `disable` | [disable](resources--api_discovery--reference--group-001.md#canonical-2113321033211313-0302322101211013-2121112302232320-3132333301103123-1012101012123313-3103222313020131-2300231201221310-3331030310012010) |
| `id` | [ID](resources--api_discovery--reference--group-001.md#canonical-0223101103121130-1000303133022330-2032113202332131-3213132033011200-3300100333202000-3100303100022110-0023312301021032-1301331201133203) |
| `labels` | [labels](resources--api_discovery--reference--group-001.md#canonical-2301220123231220-0321032332312221-0203231032312313-3112223033220330-1001333203001311-1312022203030311-3031303120232103-1212021302131031) |
| `name` | [name](resources--api_discovery--reference--group-001.md#canonical-3112323020113233-3003113311120032-1001130113302133-1022002021331233-0101321003203303-2302131323233200-0001231311230102-1321330013010010) |
| `namespace` | [namespace](resources--api_discovery--reference--group-001.md#canonical-3122132311311221-1130300031011330-2120303310331302-2313030202313311-1033003023303131-3131102130132211-3302332103102303-2330213231322010) |
| `timeouts` | [timeouts](resources--api_discovery--reference--group-001.md#canonical-1101210230113133-0233200311313013-3331012331021232-0300031132020003-3200232312222302-3203013303021203-1202223320003011-2301102001301002) |
| `timeouts.create` | [timeouts.create](resources--api_discovery--reference--group-001.md#canonical-3131130232113302-2103131332130212-2101201320021000-3002302103032202-2202321132100022-2112102321221022-0200332301001200-0201322000333333) |
| `timeouts.delete` | [timeouts.delete](resources--api_discovery--reference--group-001.md#canonical-3321120332002103-3231212001321301-3000033121313121-0210012322113233-3133112033322312-0223032101102333-1120202032000201-0122221113032122) |
| `timeouts.read` | [timeouts.read](resources--api_discovery--reference--group-001.md#canonical-0013232111110022-2310113121303312-3131133331011101-2000123031320233-2122030012023231-0030213311221232-2023322122222111-0032121123321032) |
| `timeouts.update` | [timeouts.update](resources--api_discovery--reference--group-001.md#canonical-0123123210321020-1031001301233022-0322310330230112-3033123210222001-1010230223222113-2213121120230201-0230330300203311-0000111203033030) |
| `user_defined_api_discovery_policy` | [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-0301100300331123-0311033301321010-3313113103103211-0002022220321230-1011013121032330-1101013133121333-1022132320220100-0103331312210030) |
| `user_defined_api_discovery_policy.discovery_rules` | [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-2323202303201121-1200021030321022-3311031320133103-0103211220220032-0203103201201232-2322020102203122-3001130313023302-0320333331200000) |
| `user_defined_api_discovery_policy.discovery_rules.labels` | [user_defined_api_discovery_policy.discovery_rules.labels](resources--api_discovery--reference--group-001.md#canonical-1132222011202032-0021010030101312-0102311300023211-1220112031332101-1333321232012220-2100313302100213-1023321220122113-3211212111210102) |
| `user_defined_api_discovery_policy.discovery_rules.metadata` | [user_defined_api_discovery_policy.discovery_rules.metadata](resources--api_discovery--reference--group-001.md#canonical-2013313031212331-2030211331012102-2012313002020103-3022021001003301-1032130020120131-2012213203212232-0100101122003222-2331102212132000) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` | [user_defined_api_discovery_policy.discovery_rules.metadata.description_spec](resources--api_discovery--reference--group-001.md#canonical-2302321003120133-2020003221022132-0001101010012032-3101203033033021-0323312231303202-3213103120111030-3301323122112222-0130132113201023) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.name` | [user_defined_api_discovery_policy.discovery_rules.metadata.name](resources--api_discovery--reference--group-001.md#canonical-1232032200113100-2232032000232331-0111321033121133-3202001310021023-1223133212213101-1101221330020303-1223021032210133-0121313302013212) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties` | [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-3322231013332101-0110012323031223-1011303012032000-2231331230303012-2020032320011300-2011113101022232-1213222233110103-0013211201113010) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-0022312311210230-3303331221321033-2013230022120110-0232201002000002-3021002133213113-0130332312212212-2112130311021011-1233212202221222) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](resources--api_discovery--reference--group-001.md#canonical-0110102112021331-3212102321312100-0331132020300102-1322101230003113-2202031200330010-2331113032301310-2023320303122120-0002002122130110) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](resources--api_discovery--reference--group-001.md#canonical-1113211001030210-1233111132030121-2202322002231011-2020100002130012-2122301232232232-2021230333323001-1010300122302000-2133211202202011) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](resources--api_discovery--reference--group-001.md#canonical-3233020221003233-2311232011233100-0333212223323203-3032321212330002-3203010233030222-0313331032230100-1030302220002020-2013013223313032) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name](resources--api_discovery--reference--group-001.md#canonical-3233010220121223-0212322203200000-1310331032221032-3120031310223210-3212030131030010-0032110321032123-0213011023012303-1311122223121132) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location](resources--api_discovery--reference--group-001.md#canonical-3212231212020123-3002132013132121-2010212112033211-1312100203102112-3020133303012123-3100022212312220-2300313231123110-2230133310301133) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type](resources--api_discovery--reference--group-001.md#canonical-0102113012130122-0331131112231020-1122331300330131-3111330333202323-0313223011332311-2311212312303213-0001113120310023-1011121012120322) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value](resources--api_discovery--reference--group-001.md#canonical-2021233312131120-1002313211323102-0222001031330232-0102111031013213-0211320022201210-3213330031310031-2213001003101020-0013203210132322) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](resources--api_discovery--reference--group-001.md#canonical-3230222122220021-2223303103032302-0130101311113211-2032012121202222-2312303131321130-0222221030001022-0330011222033120-3123222130023202) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern](resources--api_discovery--reference--group-001.md#canonical-3121231223010211-0221232300123203-3033121321213100-2210111101330122-0111033321102101-2031202203331330-0012103011222103-2210010212200221) |
| `user_defined_api_discovery_policy.exclusive` | [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-3023323013213321-0223121132223102-3330031233032110-2331103033222203-2130200321223321-3311121300302313-2222012200013112-0223302030010221) |
| `user_defined_api_discovery_policy.exclusive.archive` | [user_defined_api_discovery_policy.exclusive.archive](resources--api_discovery--reference--group-001.md#canonical-0221230021100333-1211310301022102-1221023233031103-1213322030031021-1302230011033203-3210311100033022-3213102023130210-1121312131110031) |
| `user_defined_api_discovery_policy.exclusive.ignore` | [user_defined_api_discovery_policy.exclusive.ignore](resources--api_discovery--reference--group-001.md#canonical-0320313322333232-1001100002302003-0122221131003331-1011110320002231-2333220302310213-0020323203030032-3323213210122130-0030031230310020) |
| `user_defined_api_discovery_policy.inclusive` | [user_defined_api_discovery_policy.inclusive](resources--api_discovery--reference--group-001.md#canonical-1102003011112221-1331320221000023-0013033221101232-3112003033011023-0230311112212312-3012222311223321-3333200202130102-1323113233133231) |

<a id="canonical-2212121133030123-3313103211030133-1332101312033102-0023310100201330-1131130101331031-0031112010301020-0220030121333230-3300021031132021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_auth_types` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- custom_auth_types

<a id="canonical-3320211132230310-2323120333312332-3223131100221213-3032122322223133-1122201100032012-3100100100313020-0332212122311231-2033033132030223"></a>

Type: `"object"`. list nested block, Optional.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("parameter_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  }
}
```

Terraform syntax:

```terraform
custom_auth_types {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013001321213112-2003200220212102-2321310202123033-3113101110203110-1300001231330103-1000332200230322-0023222333110233-1300030113022133"></a>

### Direct properties for `custom_auth_types`

<a id="canonical-1311030212221033-3231310211331320-2303312021123213-0011102233101202-3200111011220023-0300021102130303-2102333131222311-0002013012221112"></a>

#### `custom_auth_types.parameter_name` property

Type: `"string"`. Optional.

Parameter Name. The authentication parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  }
}
```

<a id="canonical-0230331323213212-1010102133233110-3131013302101333-1011230001232002-2022230222033120-0010022132111113-2122222032201312-3233323112301102"></a>

<a id="canonical-2012010012100033-1011110311103132-2332020120220323-2013221202112202-3232033001300232-2301031130022000-2111010301020021-3121100113313032"></a>

#### `custom_auth_types.parameter_type` property

Type: `"string"`. Optional.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["COOKIE","HEADER","QUERY_PARAMETER"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("QUERY_PARAMETER",
    "HEADER",
    "COOKIE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "QUERY_PARAMETER",
  "enum": [
    "QUERY_PARAMETER",
    "HEADER",
    "COOKIE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0330002233322121-1233033110301003-0230030201120323-0111002303003220-3301031130232123-2322330011312000-2312101022131231-3301102100120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- timeouts

<a id="canonical-1101210230113133-0233200311313013-3331012331021232-0300031132020003-3200232312222302-3203013303021203-1202223320003011-2301102001301002"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322020331213222-2100102331111032-2010113110113233-1231133232123233-2100211012013020-1301313232230310-2020201321100110-3001213220312103"></a>

### Direct properties for `timeouts`

<a id="canonical-3131130232113302-2103131332130212-2101201320021000-3002302103032202-2202321132100022-2112102321221022-0200332301001200-0201322000333333"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3321120332002103-3231212001321301-3000033121313121-0210012322113233-3133112033322312-0223032101102333-1120202032000201-0122221113032122"></a>

<a id="canonical-1132123231332002-2302332123301000-2201122002212310-0003023103321000-1120120211303133-2311313001331132-3313010333032033-1133023230111200"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0013232111110022-2310113121303312-3131133331011101-2000123031320233-2122030012023231-0030213311221232-2023322122222111-0032121123321032"></a>

<a id="canonical-3231111220220231-3013133300202310-3122132020031123-0200312033030330-1310131303101231-3110320200130200-2101020212202201-1313011012131131"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0123123210321020-1031001301233022-0322310330230112-3033123210222001-1010230223222113-2213121120230201-0230330300203311-0000111203033030"></a>

<a id="canonical-1302231203332112-2020010332113000-1011122000322030-1032331300110023-2310330113231032-1100320013103223-2203313230212301-2021131023031300"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- user_defined_api_discovery_policy

<a id="canonical-0301100300331123-0311033301321010-3313113103103211-0002022220321230-1011013121032330-1101013133121333-1022132320220100-0103331312210030"></a>

Type: `"object"`. single nested block, Optional.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusive",
    "inclusive")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

Terraform syntax:

```terraform
user_defined_api_discovery_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200333320123310-0210003010112113-0103223000111231-3102302323003313-1132013011130311-3303302113320031-0333130310233312-1121120321331313"></a>

### Direct properties for `user_defined_api_discovery_policy`

- [discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313): complete subsection reference.

- [exclusive](resources--api_discovery--reference--group-001.md#canonical-1323111022121002-1132133223122101-2223013020320022-0231220000301312-3210230203331032-2003110131131101-2120100303301100-3201211002333002): complete subsection reference.

- [inclusive](resources--api_discovery--reference--group-001.md#canonical-2213032120131311-3200112332220110-3021221201133002-1011013213312212-0033003103311013-0320112210022330-0300213232332322-2210313210122333): complete subsection reference.

<a id="canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- user_defined_api_discovery_policy.discovery_rules

<a id="canonical-2323202303201121-1200021030321022-3311031320133103-0103211220220032-0203103201201232-2322020102203122-3001130313023302-0320333331200000"></a>

Type: `"object"`. list nested block, Optional.

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action. Defaults to \`\[\]\`. Server applies default when
omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
discovery_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102123110233113-2010332333121313-0010120223102322-2312302301110013-0201311001013310-2123121130313210-3013022131233211-3103102201220300"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules`

- [labels](resources--api_discovery--reference--group-001.md#canonical-2103202000312133-0033230123010023-0133233023201233-1032220321002321-3331200022213030-2330211102330330-0310123121003111-0122121012220222): complete subsection reference.

- [metadata](resources--api_discovery--reference--group-001.md#canonical-2100110203233333-0232101103313212-0310312020103302-2113321311111133-1112000130310203-3231100122300011-2313210003020323-3122031220000213): complete subsection reference.

- [rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032): complete subsection reference.

<a id="canonical-2103202000312133-0033230123010023-0133233023201233-1032220321002321-3331200022213030-2330211102330330-0310123121003111-0122121012220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.labels` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- user_defined_api_discovery_policy.discovery_rules.labels

<a id="canonical-1132222011202032-0021010030101312-0102311300023211-1220112031332101-1333321232012220-2100313302100213-1023321220122113-3211212111210102"></a>

Type: `"object"`. single nested block, Optional.

Map of string keys and values that can be used to organize and categorize the rule.

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
labels {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100110203233333-0232101103313212-0310312020103302-2113321311111133-1112000130310203-3231100122300011-2313210003020323-3122031220000213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.metadata` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- user_defined_api_discovery_policy.discovery_rules.metadata

<a id="canonical-2013313031212331-2030211331012102-2012313002020103-3022021001003301-1032130020120131-2012213203212232-0100101122003222-2331102212132000"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233231232023030-1111303231312100-2022211132221122-2310111132013013-1322110001313203-1303333021000021-3222300201332231-3112101221311313"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.metadata`

<a id="canonical-2302321003120133-2020003221022132-0001101010012032-3101203033033021-0323312231303202-3213103120111030-3301323122112222-0130132113201023"></a>

#### `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1232032200113100-2232032000232331-0111321033121133-3202001310021023-1223133212213101-1101221330020303-1223021032210133-0121313302013212"></a>

<a id="canonical-0123330221223200-1000202202230001-0223221000011232-1303103211103031-0130203031302131-3002013313023003-0220023113022032-3122121020012211"></a>

#### `user_defined_api_discovery_policy.discovery_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="canonical-3322231013332101-0110012323031223-1011303012032000-2231331230303012-2020032320011300-2011113101022232-1213222233110103-0013211201113010"></a>

Type: `"object"`. single nested block, Optional.

Determines whether matching endpoints are included in API Discovery or excluded.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusion",
    "inclusion"),
  validators.ConflictingObjectAttributes("http_header_criteria",
    "pattern")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-criteria": "[\"http_header_criteria\",\"pattern\"]",
  "x-ves-oneof-field-rule_type_choice": "[\"exclusion\",\"inclusion\"]"
}
```

Terraform syntax:

```terraform
rule_properties {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003313101103321-2331220133031303-3022000130112320-1110321123000023-0200011320332012-3110131100322031-0220312230020210-2211023131123231"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties`

- [exclusion](resources--api_discovery--reference--group-001.md#canonical-2112010001030322-2131010130031232-2113133011332313-2002111202001233-1210202322303003-2313003003002331-1212321132323220-3201032221133202): complete subsection reference.

- [http_header_criteria](resources--api_discovery--reference--group-001.md#canonical-0113333110113320-2313002210321300-0203020312330021-0031000121231322-1001002023302203-3213112013303320-1220333321311302-1111010301331001): complete subsection reference.

- [inclusion](resources--api_discovery--reference--group-001.md#canonical-2313332131203231-2330302133122312-1003003322213033-1303131233123131-3223110212011313-2033112100020210-3211013032213303-1013221123200221): complete subsection reference.

<a id="canonical-3121231223010211-0221232300123203-3033121321213100-2210111101330122-0111033321102101-2031202203331330-0012103011222103-2210010212200221"></a>

<a id="canonical-2300102113322333-1220232132221321-2002321233303103-1103111313220311-2012113231312012-2320023310222310-0022200010021030-3100223321213232"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` property

Type: `"string"`. Optional.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2112010001030322-2131010130031232-2113133011332313-2002111202001233-1210202322303003-2313003003002331-1212321132323220-3201032221133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

<a id="canonical-0022312311210230-3303331221321033-2013230022120110-0232201002000002-3021002133213113-0130332312212212-2112130311021011-1233212202221222"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

Terraform syntax:

```terraform
exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011330211303310-3130001223101011-3002210221313332-1020000300220313-0030311013213332-1330020303131331-3102230121222233-1221012310001022"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion`

- [archive](resources--api_discovery--reference--group-001.md#canonical-3031132210200013-3020010332022302-0231300321303123-3210013133200112-1030220313013211-0301331031030213-0132333310033210-1201112102213130): complete subsection reference.

- [ignore](resources--api_discovery--reference--group-001.md#canonical-3020201301103131-1133213211120333-0210113023002202-2113323003121033-2332300301123221-2322003200012110-1301210002131210-3230220020201132): complete subsection reference.

<a id="canonical-3031132210200013-3020010332022302-0231300321303123-3210013133200112-1030220313013211-0301331031030213-0132333310033210-1201112102213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-2112010001030322-2131010130031232-2113133011332313-2002111202001233-1210202322303003-2313003003002331-1212321132323220-3201032221133202)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive

<a id="canonical-0110102112021331-3212102321312100-0331132020300102-1322101230003113-2202031200330010-2331113032301310-2023320303122120-0002002122130110"></a>

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
archive = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020201301103131-1133213211120333-0210113023002202-2113323003121033-2332300301123221-2322003200012110-1301210002131210-3230220020201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](resources--api_discovery--reference--group-001.md#canonical-2112010001030322-2131010130031232-2113133011332313-2002111202001233-1210202322303003-2313003003002331-1212321132323220-3201032221133202)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore

<a id="canonical-1113211001030210-1233111132030121-2202322002231011-2020100002130012-2122301232232232-2021230333323001-1010300122302000-2133211202202011"></a>

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
ignore = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113333110113320-2313002210321300-0203020312330021-0031000121231322-1001002023302203-3213112013303320-1220333321311302-1111010301331001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

<a id="canonical-3233020221003233-2311232011233100-0333212223323203-3032321212330002-3203010233030222-0313331032230100-1030302220002020-2013013223313032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header criteria.

Additional upstream details:

Criteria for matching HTTP headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("field_name",
    "value")}
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
http_header_criteria {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003231003121121-2030133131300120-0303321122021302-2011033321212113-3132332300330113-0201030213201022-0321332202113112-0201231102122010"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria`

<a id="canonical-3233010220121223-0212322203200000-1310331032221032-3120031310223210-3212030131030010-0032110321032123-0213011023012303-1311122223121132"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` property

Type: `"string"`. Optional.

HTTP Header Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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

<a id="canonical-3212231212020123-3002132013132121-2010212112033211-1312100203102112-3020133303012123-3100022212312220-2300313231123110-2230133310301133"></a>

<a id="canonical-1232322333300311-3121303200001011-2322322211321133-1133203001212301-0031013222210103-1100022011211233-2220103101102233-1233012210212110"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` property

Type: `"string"`. Optional.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["REQUEST","RESPONSE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("REQUEST",
    "RESPONSE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "REQUEST",
  "enum": [
    "REQUEST",
    "RESPONSE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0102113012130122-0331131112231020-1122331300330131-3111330333202323-0313223011332311-2311212312303213-0001113120310023-1011121012120322"></a>

<a id="canonical-0123322231232330-0011000130002102-3003222312332320-1210311031113003-2320030021302302-2103220203333110-2001330330233002-0311110102020213"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` property

Type: `"string"`. Optional.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXACT_MATCH","REGEX","SUBSTRING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EXACT_MATCH",
    "SUBSTRING",
    "REGEX"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EXACT_MATCH",
  "enum": [
    "EXACT_MATCH",
    "SUBSTRING",
    "REGEX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021233312131120-1002313211323102-0222001031330232-0102111031013213-0211320022201210-3213330031310031-2213001003101020-0013203210132322"></a>

<a id="canonical-0313101032101012-2230133110123303-2003230011333230-3033200033221232-0222003022333100-3320113031100033-0222013213230202-1232121100312213"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` property

Type: `"string"`. Optional.

Value. Configuration parameter for value

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2313332131203231-2330302133122312-1003003322213033-1303131233123131-3223110212011313-2033112100020210-3211013032213303-1013221123200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--reference--group-001.md#canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--reference--group-001.md#canonical-2321331002231121-0301012202233131-0210222002121301-2120102001133110-1102230333003131-1112200220303231-0132333122130331-3233003331023032)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion

<a id="canonical-3230222122220021-2223303103032302-0130101311113211-2032012121202222-2312303131321130-0222221030001022-0330011222033120-3123222130023202"></a>

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
inclusion = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323111022121002-1132133223122101-2223013020320022-0231220000301312-3210230203331032-2003110131131101-2120100303301100-3201211002333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- user_defined_api_discovery_policy.exclusive

<a id="canonical-3023323013213321-0223121132223102-3330031233032110-2331103033222203-2130200321223321-3311121300302313-2222012200013112-0223302030010221"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

Terraform syntax:

```terraform
exclusive {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013333323112132-2002333011120013-2003003312210000-3111210301013031-2023113201002221-1023323130210233-1222220321322301-0133002120213003"></a>

### Direct properties for `user_defined_api_discovery_policy.exclusive`

- [archive](resources--api_discovery--reference--group-001.md#canonical-0233112100020222-1012001102232121-0310231030201013-3222211101010220-0020300230130203-0031332231330033-0212102003020122-1202031322310012): complete subsection reference.

- [ignore](resources--api_discovery--reference--group-001.md#canonical-3031133300212101-0001123213300321-1333301200111021-3103100331030121-2321100311233312-1020100101032310-3110232112303200-2222232323111002): complete subsection reference.

<a id="canonical-0233112100020222-1012001102232121-0310231030201013-3222211101010220-0020300230130203-0031332231330033-0212102003020122-1202031322310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive.archive` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-1323111022121002-1132133223122101-2223013020320022-0231220000301312-3210230203331032-2003110131131101-2120100303301100-3201211002333002)
- user_defined_api_discovery_policy.exclusive.archive

<a id="canonical-0221230021100333-1211310301022102-1221023233031103-1213322030031021-1302230011033203-3210311100033022-3213102023130210-1121312131110031"></a>

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
archive = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031133300212101-0001123213300321-1333301200111021-3103100331030121-2321100311233312-1020100101032310-3110232112303200-2222232323111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive.ignore` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--reference--group-001.md#canonical-1323111022121002-1132133223122101-2223013020320022-0231220000301312-3210230203331032-2003110131131101-2120100303301100-3201211002333002)
- user_defined_api_discovery_policy.exclusive.ignore

<a id="canonical-0320313322333232-1001100002302003-0122221131003331-1011110320002231-2333220302310213-0020323203030032-3323213210122130-0030031230310020"></a>

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
ignore = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213032120131311-3200112332220110-3021221201133002-1011013213312212-0033003103311013-0320112210022330-0300213232332322-2210313210122333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.inclusive` properties

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Property reference](resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [user_defined_api_discovery_policy](resources--api_discovery--reference--group-001.md#canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122)
- user_defined_api_discovery_policy.inclusive

<a id="canonical-1102003011112221-1331320221000023-0013033221101232-3112003033011023-0230311112212312-3012222311223321-3333200202130102-1323113233133231"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
inclusive = {}
```

This is an empty object or choice marker. It has no direct properties.
