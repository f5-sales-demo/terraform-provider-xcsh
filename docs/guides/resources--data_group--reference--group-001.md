---
page_title: "xcsh_data_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group reference."
---

# xcsh_data_group reference

<a id="canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- Property reference

<a id="canonical-0103330120211221-1033131220001102-2212222201210110-1332223322230132-2020322111113311-0301330101121023-1101000001132023-1300212300010232"></a>

### Direct properties for `xcsh_data_group`

- [address_records](resources--data_group--reference--group-001.md#canonical-1200023100312100-3032303031001130-0022031202332330-3312333132103031-1322222120203122-1213021011122011-0003021001311301-0012003310211130): complete subsection reference.

<a id="canonical-2102032000333023-2212322020123102-3300123300311010-3303001310100132-2022231230113031-0303230103120130-1130010131101122-3103313331113311"></a>

<a id="canonical-3020101333021310-3212012220302111-1031002111202000-2022203213230320-3213220022210300-0033100311203032-3000031231210322-3110123032003300"></a>

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

<a id="canonical-1331101231213312-1023112012213210-1021101302121032-2013323103220213-3000132022000123-0120120131303121-1212123013132111-2203122011102332"></a>

<a id="canonical-0222102232133002-2022221122332223-3022312112130220-3003001023013011-2020013313110210-0023102231021111-1110113220302301-3111131103011313"></a>

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

<a id="canonical-3021113320110133-3221010013101020-3232022102011321-1203023112323203-1301022321101012-0223330100332210-3100121322011012-2101122203033122"></a>

<a id="canonical-0333132201223302-3220033221231131-3023013313331331-2113032030002210-0333011233230200-2121101000332203-2330200320331231-1022221032131220"></a>

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

<a id="canonical-1030320333130213-0220232222102030-1000130321232100-3220202023222301-0323033101213002-2020202113211230-2120230221100311-1332111012020221"></a>

<a id="canonical-0022103231201010-1222022012121231-2003323031032013-1200102111301323-1203213231310230-0121233102023022-0323210012330210-2012233333112101"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [integer_records](resources--data_group--reference--group-001.md#canonical-2300211220223031-2002323232103033-1232110333132002-2103220023013000-3201132301003203-0033320000330020-1200101010000223-2113112000301222): complete subsection reference.

<a id="canonical-2001300231303320-2323020212212230-1210313300133220-0332303013320002-1300000300111010-0302202233120020-2013001210100130-3303113210303030"></a>

<a id="canonical-1221121222011303-2223131132311301-0130101122032031-0133223313332221-3221121122030201-0111013011212311-0331213102020222-3003111021013003"></a>

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

<a id="canonical-3121322213121000-3012133313101233-0312311023202120-2231221122103013-1031210313023102-3301100223111121-3331212302333333-0012030012132331"></a>

<a id="canonical-2322013130212010-0202202011110003-2132003213122011-0031012103121010-0303201332232130-3212210232022012-0221031200202331-2332111322121113"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Data Group. Must be unique within the namespace.

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

<a id="canonical-0023133120202312-2022120111322311-2110023012112012-1011033002312220-0210323123011011-2213232311120033-1331003021302230-2011020320212023"></a>

<a id="canonical-1123211001123123-0010231310011210-2300312031301202-0002023202111033-0322322220122123-3112312300323112-3232121111013000-3320301012133300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Data Group is created.

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

- [string_records](resources--data_group--reference--group-001.md#canonical-3003011112012003-3321223223000230-1021221032211011-2020033331312312-3103223302132312-1313130123231311-0323230132031232-1213121320302323): complete subsection reference.

- [timeouts](resources--data_group--reference--group-001.md#canonical-1132320220121301-1023013312102120-3031320011031120-0133201213031331-3120333203320122-2222120333032210-0113212033303110-1022102212001001): complete subsection reference.

<a id="canonical-3203310023113311-1202122013322130-0123330331211210-1000021120103231-2312100103031103-3003102032103332-0021122031322132-0121113133022021"></a>

### All schema paths for `xcsh_data_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_records` | [address_records](resources--data_group--reference--group-001.md#canonical-0311100131010001-2213310210100211-2112022211022133-3323123030022201-2020110011000302-2013112020323200-0233322010002330-0302233300113220) |
| `address_records.records` | [address_records.records](resources--data_group--reference--group-001.md#canonical-3033211301030320-1313201311210312-1002101213223203-1131000032322332-2332101210111032-1110102331213333-2223322122132333-1022313021303332) |
| `annotations` | [annotations](resources--data_group--reference--group-001.md#canonical-2102032000333023-2212322020123102-3300123300311010-3303001310100132-2022231230113031-0303230103120130-1130010131101122-3103313331113311) |
| `description` | [description](resources--data_group--reference--group-001.md#canonical-1331101231213312-1023112012213210-1021101302121032-2013323103220213-3000132022000123-0120120131303121-1212123013132111-2203122011102332) |
| `disable` | [disable](resources--data_group--reference--group-001.md#canonical-3021113320110133-3221010013101020-3232022102011321-1203023112323203-1301022321101012-0223330100332210-3100121322011012-2101122203033122) |
| `id` | [ID](resources--data_group--reference--group-001.md#canonical-1030320333130213-0220232222102030-1000130321232100-3220202023222301-0323033101213002-2020202113211230-2120230221100311-1332111012020221) |
| `integer_records` | [integer_records](resources--data_group--reference--group-001.md#canonical-0003110301233113-0003002000032330-1121320311301202-0331002200132030-0321000012202032-2320320233301102-2312113032201300-3000233301301230) |
| `integer_records.records` | [integer_records.records](resources--data_group--reference--group-001.md#canonical-1022102213303001-1021110030102132-1333031022213030-3200211032223223-2301212121012223-3023233213200123-1231000201001320-3001103101303021) |
| `labels` | [labels](resources--data_group--reference--group-001.md#canonical-2001300231303320-2323020212212230-1210313300133220-0332303013320002-1300000300111010-0302202233120020-2013001210100130-3303113210303030) |
| `name` | [name](resources--data_group--reference--group-001.md#canonical-3121322213121000-3012133313101233-0312311023202120-2231221122103013-1031210313023102-3301100223111121-3331212302333333-0012030012132331) |
| `namespace` | [namespace](resources--data_group--reference--group-001.md#canonical-0023133120202312-2022120111322311-2110023012112012-1011033002312220-0210323123011011-2213232311120033-1331003021302230-2011020320212023) |
| `string_records` | [string_records](resources--data_group--reference--group-001.md#canonical-0012202003110031-2003012111233112-3011331122202003-0302331321210101-3120321003223020-1031031213222202-0310101233203112-3302220101130300) |
| `string_records.records` | [string_records.records](resources--data_group--reference--group-001.md#canonical-2220333321202323-1003121310021030-0332010220123321-0031123320121011-3210122110010000-0000230302221310-1210223021323323-3300332013023102) |
| `timeouts` | [timeouts](resources--data_group--reference--group-001.md#canonical-1203302000011120-3222003322003302-3131323202110322-0322132022100010-3221202331010131-1020002213223130-0021213021113322-1100113221032102) |
| `timeouts.create` | [timeouts.create](resources--data_group--reference--group-001.md#canonical-0011030212110031-2200121022013021-2210112212323111-1300033003033012-2101011023311211-0113032103011223-3323311122112010-0211313331111232) |
| `timeouts.delete` | [timeouts.delete](resources--data_group--reference--group-001.md#canonical-3312001201112113-2010303033032322-1203320031023123-0223030003002311-2003021133033130-1220031231221313-2233121210222002-3103201213123232) |
| `timeouts.read` | [timeouts.read](resources--data_group--reference--group-001.md#canonical-0333113003011232-3231130332011100-1323321320320313-2313133123230012-0010111232010321-2001132100132000-2122123023330112-2202002201111022) |
| `timeouts.update` | [timeouts.update](resources--data_group--reference--group-001.md#canonical-1211311010003011-2100230212101230-2210100210121330-1030312031011030-0301313303132232-3333110202011320-1321300020002101-0312121303012033) |

<a id="canonical-1200023100312100-3032303031001130-0022031202332330-3312333132103031-1322222120203122-1213021011122011-0003021001311301-0012003310211130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `address_records` properties

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- [Property reference](resources--data_group--reference--group-001.md#canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132)
- address_records

<a id="canonical-0311100131010001-2213310210100211-2112022211022133-3323123030022201-2020110011000302-2013112020323200-0233322010002330-0302233300113220"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: address\_records, integer\_records, string\_records\] Address Record. Data group with
address record List.

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

- [address_records](resources--data_group--reference--group-001.md#canonical-0311100131010001-2213310210100211-2112022211022133-3323123030022201-2020110011000302-2013112020323200-0233322010002330-0302233300113220)
- [integer_records](resources--data_group--reference--group-001.md#canonical-0003110301233113-0003002000032330-1121320311301202-0331002200132030-0321000012202032-2320320233301102-2312113032201300-3000233301301230)
- [string_records](resources--data_group--reference--group-001.md#canonical-0012202003110031-2003012111233112-3011331122202003-0302331321210101-3120321003223020-1031031213222202-0310101233203112-3302220101130300)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
address_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112301330310230-0320320302320213-1021230111231133-3313013112132023-1100031203111312-0213300131130031-0033011301222233-1103002330222001"></a>

### Direct properties for `address_records`

<a id="canonical-3033211301030320-1313201311210312-1002101213223203-1131000032322332-2332101210111032-1110102331213333-2223322122132333-1022313021303332"></a>

#### `address_records.records` property

Type: `["map", "string"]`. Optional.

Address records. Configuration parameter for records

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":4096},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"format\":\"ip-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.ip\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"4096\",\"ves.io.schema.rules.map.values.string.pattern\":\"^[0-9a-zA-Z._-]*$\"},\"values\":{\"pattern\":\"^[0-9a-zA-Z._-]*$\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "format": "ip-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.ip": "true",
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
    },
    "values": {
      "pattern": "^[0-9a-zA-Z._-]*$",
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
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-2300211220223031-2002323232103033-1232110333132002-2103220023013000-3201132301003203-0033320000330020-1200101010000223-2113112000301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `integer_records` properties

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- [Property reference](resources--data_group--reference--group-001.md#canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132)
- integer_records

<a id="canonical-0003110301233113-0003002000032330-1121320311301202-0331002200132030-0321000012202032-2320320233301102-2312113032201300-3000233301301230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for integer records.

Additional upstream details:

Data group with integer record List.

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
integer_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301213223213200-3210121311023003-0030030230233321-1031223202021202-2301230202302333-3331000100211221-0210330203301220-2321233231221131"></a>

### Direct properties for `integer_records`

<a id="canonical-1022102213303001-1021110030102132-1333031022213030-3200211032223223-2301212121012223-3023233213200123-1231000201001320-3001103101303021"></a>

#### `integer_records.records` property

Type: `["map", "string"]`. Optional.

Integer records. Configuration parameter for records

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":4096},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"originalRules\":{\"ves.io.schema.rules.map.max_pairs\":\"4096\",\"ves.io.schema.rules.map.values.string.pattern\":\"^[a-zA-Z_][a-zA-Z0-9_]*$\"},\"values\":{\"pattern\":\"^[a-zA-Z_][a-zA-Z0-9_]*$\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
    },
    "values": {
      "pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  }
}
```

<a id="canonical-3003011112012003-3321223223000230-1021221032211011-2020033331312312-3103223302132312-1313130123231311-0323230132031232-1213121320302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `string_records` properties

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- [Property reference](resources--data_group--reference--group-001.md#canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132)
- string_records

<a id="canonical-0012202003110031-2003012111233112-3011331122202003-0302331321210101-3120321003223020-1031031213222202-0310101233203112-3302220101130300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for string records.

Additional upstream details:

Data group with strings record List.

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
string_records {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122301022112133-2111111231133031-3113311010120121-1303332011301001-1203020120110130-0112232110303333-2213112033222230-1231011332110010"></a>

### Direct properties for `string_records`

<a id="canonical-2220333321202323-1003121310021030-0332010220123321-0031123320121011-3210122110010000-0000230302221310-1210223021323323-3300332013023102"></a>

#### `string_records.records` property

Type: `["map", "string"]`. Optional.

String records. Configuration parameter for records

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":4096},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"originalRules\":{\"ves.io.schema.rules.map.max_pairs\":\"4096\",\"ves.io.schema.rules.map.values.string.pattern\":\"^[0-9a-zA-Z._-]*$\"},\"values\":{\"pattern\":\"^[0-9a-zA-Z._-]*$\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
    },
    "values": {
      "pattern": "^[0-9a-zA-Z._-]*$",
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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-1132320220121301-1023013312102120-3031320011031120-0133201213031331-3120333203320122-2222120333032210-0113212033303110-1022102212001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- [Property reference](resources--data_group--reference--group-001.md#canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132)
- timeouts

<a id="canonical-1203302000011120-3222003322003302-3131323202110322-0322132022100010-3221202331010131-1020002213223130-0021213021113322-1100113221032102"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130132221113130-0122302211203232-0233022101122030-1222033131332313-3212200223022100-0332213200301111-3102220233112321-3310113201131201"></a>

### Direct properties for `timeouts`

<a id="canonical-0011030212110031-2200121022013021-2210112212323111-1300033003033012-2101011023311211-0113032103011223-3323311122112010-0211313331111232"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3312001201112113-2010303033032322-1203320031023123-0223030003002311-2003021133033130-1220031231221313-2233121210222002-3103201213123232"></a>

<a id="canonical-1000312110100010-1201131012220120-1201011203320202-2112033230332320-0000112233020013-3033333112122033-3131031133001303-1133123103312132"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0333113003011232-3231130332011100-1323321320320313-2313133123230012-0010111232010321-2001132100132000-2122123023330112-2202002201111022"></a>

<a id="canonical-2223203333321211-0010113223213223-1001233330031130-1322303011323313-1232002203133010-0133323223001032-3101122102011023-2020301132330002"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1211311010003011-2100230212101230-2210100210121330-1030312031011030-0301313303132232-3333110202011320-1321300020002101-0312121303012033"></a>

<a id="canonical-1233100230211310-1313321003113223-0300130233122013-2222032123321232-2230101233130233-1000033323131013-0002133212303212-3232033320130310"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
