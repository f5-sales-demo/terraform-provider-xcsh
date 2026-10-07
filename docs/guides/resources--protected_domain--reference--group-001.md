---
page_title: "xcsh_protected_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain reference."
---

# xcsh_protected_domain reference

<a id="canonical-0031202300233031-2101003333102132-1231000303030120-1020000033013310-3100323321110030-2313323021230021-3110000211021111-1302003213131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110)
- Property reference

<a id="canonical-3102103133201013-3001230200230111-2020330122221000-0311102321310112-2202213210232212-2021312313033330-3023031333312102-3112132210223022"></a>

### Direct properties for `xcsh_protected_domain`

<a id="canonical-2132331033221222-1100310212032230-2033313010002333-2200201321300012-2000032333223031-2233202310021100-1020210000322000-2003311320332110"></a>

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

<a id="canonical-0102231300233323-3203023301220122-3220330122232021-1012210120312233-0100312120100333-2131121230300320-1213321031003332-2222033301110302"></a>

<a id="canonical-2022212100120213-2030011330022131-2132230310033130-3203001210210202-3300112211133331-0310003001201102-3123002310303133-0220032230331011"></a>

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

<a id="canonical-2101102320230013-3030113031110013-3220001220222221-2010101113201103-0200313203220011-0110232333131233-1113010201201101-3031111300000103"></a>

<a id="canonical-3231021302101131-3233000230220321-0012010012202323-3232310322031031-2301022303233132-0323030322321113-1011203312112132-1003322013220031"></a>

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

<a id="canonical-0301032322322022-0002300210131000-0313120012122233-3220011313210212-2222330010122331-0201221233012222-2113333111030230-2032333220211321"></a>

<a id="canonical-3222302222020321-2331111001102002-0331103120000311-1120331102130300-1232113320322133-2111120121001222-1221111130310030-3123022312310210"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1223302003221302-0031301020011011-1113312223222213-1200200013130111-2201301213321210-0122101332231212-3033221110233123-3320120023021303"></a>

<a id="canonical-1103111220332212-1233203211101002-3331311001202331-1300212023123223-0102302211330010-1023210300132313-3212322120321222-1212002223300313"></a>

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

<a id="canonical-0133031232211231-2100232002031110-3120130321212011-1123320022231233-3210311322321320-2013330133121031-1102122303233033-1312020010201101"></a>

<a id="canonical-2302130002122002-0222013200301012-2102122123322302-3312300310120002-1312100113033302-1020223120012222-2223303001213032-2232121300310022"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Protected Domain. Must be unique within the namespace.

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

<a id="canonical-0330021223100033-1132110320011213-3103320200010333-1323222023111232-1112300021210213-3211323103003132-1133201033222302-0023101232313023"></a>

<a id="canonical-1103213200301021-0212132311011131-1213012022003102-2221233232022103-0012233020312012-1130313312221233-3131131323022233-1221030322202030"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Protected Domain is created.

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

<a id="canonical-1000233200233200-1311302123331321-0230032221203132-1203320012322011-2111321133101002-2103012323110123-3011223331100032-2131111010132311"></a>

<a id="canonical-1133320031200010-1210102311011030-1002130023211130-1300022203300302-3100230210222110-0131332001112221-1120112031122331-3133121002233032"></a>

#### `protected_domain` property

Type: `"string"`. Required.

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below. Example: if you are adding Client-Side Defense JS on checkout.example.com, you
should enter example.com here.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
  validators.ETLDPlusOneValidator(),
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [timeouts](resources--protected_domain--reference--group-001.md#canonical-1220210033121133-3313201301123333-3003302123330310-3300111130123120-2033231020300331-0332131030002313-0020302132202123-2132030000213013): complete subsection reference.

<a id="canonical-3121213112300110-3100310003321201-3321010331230132-1110011011302213-3331010222320132-2132233213202010-2312133300003013-0133130333210020"></a>

### All schema paths for `xcsh_protected_domain`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--protected_domain--reference--group-001.md#canonical-2132331033221222-1100310212032230-2033313010002333-2200201321300012-2000032333223031-2233202310021100-1020210000322000-2003311320332110) |
| `description` | [description](resources--protected_domain--reference--group-001.md#canonical-0102231300233323-3203023301220122-3220330122232021-1012210120312233-0100312120100333-2131121230300320-1213321031003332-2222033301110302) |
| `disable` | [disable](resources--protected_domain--reference--group-001.md#canonical-2101102320230013-3030113031110013-3220001220222221-2010101113201103-0200313203220011-0110232333131233-1113010201201101-3031111300000103) |
| `id` | [ID](resources--protected_domain--reference--group-001.md#canonical-0301032322322022-0002300210131000-0313120012122233-3220011313210212-2222330010122331-0201221233012222-2113333111030230-2032333220211321) |
| `labels` | [labels](resources--protected_domain--reference--group-001.md#canonical-1223302003221302-0031301020011011-1113312223222213-1200200013130111-2201301213321210-0122101332231212-3033221110233123-3320120023021303) |
| `name` | [name](resources--protected_domain--reference--group-001.md#canonical-0133031232211231-2100232002031110-3120130321212011-1123320022231233-3210311322321320-2013330133121031-1102122303233033-1312020010201101) |
| `namespace` | [namespace](resources--protected_domain--reference--group-001.md#canonical-0330021223100033-1132110320011213-3103320200010333-1323222023111232-1112300021210213-3211323103003132-1133201033222302-0023101232313023) |
| `protected_domain` | [protected_domain](resources--protected_domain--reference--group-001.md#canonical-1000233200233200-1311302123331321-0230032221203132-1203320012322011-2111321133101002-2103012323110123-3011223331100032-2131111010132311) |
| `timeouts` | [timeouts](resources--protected_domain--reference--group-001.md#canonical-2113233200011033-2002011303230013-3031110002001310-0203330021331311-0330101133033102-3030030221230013-1123201103122003-0213001112030312) |
| `timeouts.create` | [timeouts.create](resources--protected_domain--reference--group-001.md#canonical-1300331200002112-0013123103121012-1323001322130020-2231203131301113-3120013123010221-2230030330131202-2213123321333221-0232110122102230) |
| `timeouts.delete` | [timeouts.delete](resources--protected_domain--reference--group-001.md#canonical-1203112112000221-2022023112101332-2210020102112123-1113201300131003-0312222111233211-1322123020221203-0333210101031220-0021331220322123) |
| `timeouts.read` | [timeouts.read](resources--protected_domain--reference--group-001.md#canonical-2330210322003101-1120211100310331-0131110211101133-0023000230101303-3233330321103312-3002032321121313-2013201022213213-3331333001310002) |
| `timeouts.update` | [timeouts.update](resources--protected_domain--reference--group-001.md#canonical-3010312130323220-2121132330332332-2302013100120301-1333122302332333-2233312020120001-2100020230322000-3113132120111112-0120123312103230) |

<a id="canonical-1220210033121133-3313201301123333-3003302123330310-3300111130123120-2033231020300331-0332131030002313-0020302132202123-2132030000213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110)
- [Property reference](resources--protected_domain--reference--group-001.md#canonical-0031202300233031-2101003333102132-1231000303030120-1020000033013310-3100323321110030-2313323021230021-3110000211021111-1302003213131230)
- timeouts

<a id="canonical-2113233200011033-2002011303230013-3031110002001310-0203330021331311-0330101133033102-3030030221230013-1123201103122003-0213001112030312"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223210020301123-2132210031122201-0312121201322202-2110300020022002-2203213302310023-2120333323021231-2333331003121133-3012212303032013"></a>

### Direct properties for `timeouts`

<a id="canonical-1300331200002112-0013123103121012-1323001322130020-2231203131301113-3120013123010221-2230030330131202-2213123321333221-0232110122102230"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1203112112000221-2022023112101332-2210020102112123-1113201300131003-0312222111233211-1322123020221203-0333210101031220-0021331220322123"></a>

<a id="canonical-3111313000030112-1210132032200001-3320303112131222-1003300110310103-2310110232120332-3232021121023221-1203121001131111-0020221203321210"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2330210322003101-1120211100310331-0131110211101133-0023000230101303-3233330321103312-3002032321121313-2013201022213213-3331333001310002"></a>

<a id="canonical-0322131202332133-1133303221222023-1011103010200333-3010312102101301-0022033323020300-1321100000020031-3223311211031321-3331320013012303"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3010312130323220-2121132330332332-2302013100120301-1333122302332333-2233312020120001-2100020230322000-3113132120111112-0120123312103230"></a>

<a id="canonical-1311322303011230-0100323233112010-2100223013310301-0013000113213033-3032100313333030-2232213101021032-1221102012233010-3132302011212330"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
