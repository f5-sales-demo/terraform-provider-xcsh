---
page_title: "xcsh_authorization_server reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server reference."
---

# xcsh_authorization_server reference

<a id="canonical-0211111022313012-0333312311133222-3201102131123002-1000213313220112-0313210110001033-0311120333133101-3203103013101012-3320032021010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331)
- Property reference

<a id="canonical-3200013103322110-3213010103011120-0001330221231310-3101000022223220-0232030113002112-1220232322230113-2211200223030101-1233020131333123"></a>

### Direct properties for `xcsh_authorization_server`

<a id="canonical-3100121002020232-3021233321202211-3221323321032213-2132302013001222-2303002022130020-2223123031210300-3003102330130300-3031130232203222"></a>

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

<a id="canonical-2133032233103312-1032201021032101-0030213130223103-0111220321023131-2232302333103111-3200230303303222-3320112103312301-0013123230200331"></a>

<a id="canonical-0322030120122212-2323322120013222-2020200102330102-3201322230222221-0222303131010022-1312102311311101-3212003331011212-1201313303231121"></a>

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

<a id="canonical-1003011110030110-1103221111101123-1331301020332332-1100312210223103-0123200332313002-2302112021303022-0103221022113111-0332303133103232"></a>

<a id="canonical-0220121233331113-2302223211222102-0103022311000001-1112000222102132-1113200302112303-2333012020231011-2301233311232303-0311132313222332"></a>

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

<a id="canonical-0221021322212000-0332002111222222-2131012211121110-1233331202200230-1223312003132122-1230023200000023-0012310310113231-3102310023201310"></a>

<a id="canonical-1330221312323102-2220011111023321-2203110002320312-2202323033331032-1020201302020232-3312120000001223-3311130200223020-0211011001301230"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2333322200331103-3301130023310110-3211223322310123-0311102322333132-3112133000211332-3112212211130023-0323232220020221-0301302032023333"></a>

<a id="canonical-0202113201101200-3300133211133103-0132132133012111-1113003312031333-0030303212213020-3202022033101031-3322230203320122-1212123320021002"></a>

#### `jwks_uri` property

Type: `"string"`. Required.

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  }
}
```

<a id="canonical-1230020131312101-3100101101231123-1331032112111010-0333313231331332-0220213202000220-3310320301022122-3110310011321300-1213320313311302"></a>

<a id="canonical-0233103101130030-3302313112002302-3230231312023103-3212022021202013-3331100311030113-0020222222120232-0133013300033111-1000121330133020"></a>

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

<a id="canonical-3312032230200222-2111201032220111-1103032021221302-1331121201032233-1202032200113222-2321321023132000-1001201101313103-0111122002321203"></a>

<a id="canonical-3213021030032133-1213032332100032-2222133211110122-1011211330300232-0020100030123030-2322033213331131-1200122231221320-0113023101122120"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Authorization Server. Must be unique within the namespace.

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

<a id="canonical-1113032210123002-0003320131103230-0332313013001013-2203013033110223-1130331200233312-1203122231301120-0211031212002232-0233331203321021"></a>

<a id="canonical-0122101202033330-3220000302311232-0111301113230123-3022311321111102-1331003002123102-3011302302112102-0321321213010011-2331000300011110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Authorization Server is created.

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

- [timeouts](resources--authorization_server--reference--group-001.md#canonical-3320232100332230-2102322332023311-1310321102020102-3012303101230013-0130201221220123-3131002310333103-3201101011230113-0220112012211300): complete subsection reference.

<a id="canonical-3330000032311101-3100100132123201-1321112231112123-1023123123302223-1020320230122301-1112231202220220-1112200133102121-0122332210221002"></a>

### All schema paths for `xcsh_authorization_server`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--authorization_server--reference--group-001.md#canonical-3100121002020232-3021233321202211-3221323321032213-2132302013001222-2303002022130020-2223123031210300-3003102330130300-3031130232203222) |
| `description` | [description](resources--authorization_server--reference--group-001.md#canonical-2133032233103312-1032201021032101-0030213130223103-0111220321023131-2232302333103111-3200230303303222-3320112103312301-0013123230200331) |
| `disable` | [disable](resources--authorization_server--reference--group-001.md#canonical-1003011110030110-1103221111101123-1331301020332332-1100312210223103-0123200332313002-2302112021303022-0103221022113111-0332303133103232) |
| `id` | [ID](resources--authorization_server--reference--group-001.md#canonical-0221021322212000-0332002111222222-2131012211121110-1233331202200230-1223312003132122-1230023200000023-0012310310113231-3102310023201310) |
| `jwks_uri` | [jwks_uri](resources--authorization_server--reference--group-001.md#canonical-2333322200331103-3301130023310110-3211223322310123-0311102322333132-3112133000211332-3112212211130023-0323232220020221-0301302032023333) |
| `labels` | [labels](resources--authorization_server--reference--group-001.md#canonical-1230020131312101-3100101101231123-1331032112111010-0333313231331332-0220213202000220-3310320301022122-3110310011321300-1213320313311302) |
| `name` | [name](resources--authorization_server--reference--group-001.md#canonical-3312032230200222-2111201032220111-1103032021221302-1331121201032233-1202032200113222-2321321023132000-1001201101313103-0111122002321203) |
| `namespace` | [namespace](resources--authorization_server--reference--group-001.md#canonical-1113032210123002-0003320131103230-0332313013001013-2203013033110223-1130331200233312-1203122231301120-0211031212002232-0233331203321021) |
| `timeouts` | [timeouts](resources--authorization_server--reference--group-001.md#canonical-2000012033031322-2212302001211310-1222212222203003-0220302022013013-3030322200033230-3023312201122330-0200301330013231-1212202201131022) |
| `timeouts.create` | [timeouts.create](resources--authorization_server--reference--group-001.md#canonical-1100213213023012-3201011232331132-3321011311133012-1231203312001310-1112233132202011-1121100232221232-0022321122333313-0002231122003332) |
| `timeouts.delete` | [timeouts.delete](resources--authorization_server--reference--group-001.md#canonical-3333011031301010-0203213321003230-0312120101110213-2301202203321111-2311212030121100-3210003300003031-3311032101001030-1120312110103202) |
| `timeouts.read` | [timeouts.read](resources--authorization_server--reference--group-001.md#canonical-3021133102311220-3023022000020121-0210313211023321-1332123113233322-3322130001111330-3310332110013122-1333011202323320-2121223320223321) |
| `timeouts.update` | [timeouts.update](resources--authorization_server--reference--group-001.md#canonical-0310000032303113-2203232012313213-3303212320101103-3333123121013200-3112222230302231-0213200011332102-3022311223230221-0320230111113230) |

<a id="canonical-3320232100332230-2102322332023311-1310321102020102-3012303101230013-0130201221220123-3131002310333103-3201101011230113-0220112012211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331)
- [Property reference](resources--authorization_server--reference--group-001.md#canonical-0211111022313012-0333312311133222-3201102131123002-1000213313220112-0313210110001033-0311120333133101-3203103013101012-3320032021010331)
- timeouts

<a id="canonical-2000012033031322-2212302001211310-1222212222203003-0220302022013013-3030322200033230-3023312201122330-0200301330013231-1212202201131022"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333213211200303-1032311112311320-2200331113232031-1022201311310310-0120213232330020-1033332312111230-0011300312210001-0020230331131100"></a>

### Direct properties for `timeouts`

<a id="canonical-1100213213023012-3201011232331132-3321011311133012-1231203312001310-1112233132202011-1121100232221232-0022321122333313-0002231122003332"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3333011031301010-0203213321003230-0312120101110213-2301202203321111-2311212030121100-3210003300003031-3311032101001030-1120312110103202"></a>

<a id="canonical-0022301211112121-3300002132312033-2302103133303321-0101321131020130-2223013220213223-1131001123113111-1201001323310021-1211001013201200"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3021133102311220-3023022000020121-0210313211023321-1332123113233322-3322130001111330-3310332110013122-1333011202323320-2121223320223321"></a>

<a id="canonical-3312130301232231-1203333011000033-1313323332203222-2120330210303300-2222311330003103-0331100222212333-2120323133112020-2320301133012302"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0310000032303113-2203232012313213-3303212320101103-3333123121013200-3112222230302231-0213200011332102-3022311223230221-0320230111113230"></a>

<a id="canonical-3332230031113002-1202213230112333-0013201133112320-1201031122223121-0213300100011032-3232002012300120-2203212300031303-1103023323300210"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
