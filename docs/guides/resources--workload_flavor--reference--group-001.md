---
page_title: "xcsh_workload_flavor reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor reference."
---

# xcsh_workload_flavor reference

<a id="canonical-2201230231331130-0233333320310023-2033232302300302-2210032310122123-0221311000031233-1020311213100203-1113130201012033-0300212122012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001)
- Property reference

<a id="canonical-0111033100312130-2220003020021202-1032302021202232-1131223011020033-3220323313310111-2210100033122310-2010310311133022-0112003110333322"></a>

### Direct properties for `xcsh_workload_flavor`

<a id="canonical-3101112201013003-3011221013123010-2211230103213223-2011022200311103-0230210102320310-1011021320130211-0022323113020101-0131122201220010"></a>

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

<a id="canonical-2230202302113302-2312323223232321-2012332023200332-2013312001132311-1213202201003222-1333001302020331-0333011131132310-1233301100112220"></a>

<a id="canonical-1303222121032313-1210212213033312-1323030132200220-2331012020312120-1203310311231020-3011203110211322-3022121312221133-3133312123320101"></a>

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

<a id="canonical-3012123303321311-0131100212111303-2003123013112302-0011101230333220-0123120031301323-2233100111003131-0011231122223111-2131301123223003"></a>

<a id="canonical-3301110032202032-0231031223110132-0112211012200023-2020122220003233-2310030031222120-2322223110201201-3010013312220033-2331323231210001"></a>

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

<a id="canonical-3021101212313132-2133300232221211-3133123220310310-2031211011131203-2003003010321120-1202023211131033-2303003000011031-2232213233222032"></a>

<a id="canonical-3030001013210112-3221100103110330-0230210113001320-3013231133112032-2122133300320032-1102102303301232-3100023201030032-3300201302300103"></a>

#### `ephemeral_storage` property

Type: `"string"`. Optional, Computed.

Ephemeral storage in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  }
}
```

<a id="canonical-3200203132323110-3003202021232031-0213131301110223-2110031313023022-0303001201233303-1100320323003111-2221110103233120-2103010213133230"></a>

<a id="canonical-3202020032033110-2000022012013130-0230120313010231-0313230323332003-2211031213011303-1013200031200220-3232013120121230-2233121011030312"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2101321210233303-2323212112123232-1321121232132122-0212201010023213-2031311311202010-2333133010013300-0221211211311330-0331320131123320"></a>

<a id="canonical-2301112212220010-1333220201102030-3123300003212332-2222322013002203-1320110001030232-3230112022313201-1213020031122313-3030320231130112"></a>

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

<a id="canonical-2222321003220310-1113312123000210-1232313200321212-1112013122311311-2331030031320012-1302313230311312-1213001323333013-3210132321112003"></a>

<a id="canonical-0101332000223203-0320012130320123-2102201212323211-1122313031232202-3022110001133123-0020030113202230-0301313123101332-3313102102020002"></a>

#### `memory` property

Type: `"string"`. Optional, Computed.

Memory in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  }
}
```

<a id="canonical-3101123100221310-3323002021232222-1010221220213312-1113213033031211-2213120303221002-0311120302222222-0113012310122203-0012212223001010"></a>

<a id="canonical-1110002121233232-2122321031022223-0331002322013303-0232022013112030-1302222130312010-3000221321230010-1323220212122021-0221122003221211"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Workload Flavor. Must be unique within the namespace.

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

<a id="canonical-2031202002100331-0011331010230212-2203121322302300-0101230011332300-0210001002122312-2000122222100112-0110031030112310-0333332313320331"></a>

<a id="canonical-2333002312313203-2221232103030112-1021312220012000-1300031020311200-1213233123221111-1013332033022122-0133000201203223-3212210322110131"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Workload Flavor. The F5 XC API restricts this resource to the shared namespace; it
defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("shared")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["shared"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("shared"),
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

- [timeouts](resources--workload_flavor--reference--group-001.md#canonical-3232213233303102-3110020321220110-3223332200322301-0333130201321131-1331230010022023-3311301302101312-1221022003323311-1021200120010202): complete subsection reference.

<a id="canonical-2300132331222310-1031020230030230-2330231023023321-1311203001321013-1202311032323121-2212022322111022-2310132112331212-2213013331231121"></a>

<a id="canonical-0020123300102010-0323310003100203-2033032132133120-2222331201113322-1331331213130302-3010222100223111-2002120302133223-1211210030102201"></a>

#### `vcpus` property

Type: `"number"`. Optional, Computed.

Number of vCPUs allocated for the workload\_flavor. Each vCPU is a thread on a CPU core.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  }
}
```

<a id="canonical-1233223003303323-2231222300312111-1103331002221022-1023020111012103-1320032132020203-0330210200013231-2110012310300130-3231210012303131"></a>

### All schema paths for `xcsh_workload_flavor`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--workload_flavor--reference--group-001.md#canonical-3101112201013003-3011221013123010-2211230103213223-2011022200311103-0230210102320310-1011021320130211-0022323113020101-0131122201220010) |
| `description` | [description](resources--workload_flavor--reference--group-001.md#canonical-2230202302113302-2312323223232321-2012332023200332-2013312001132311-1213202201003222-1333001302020331-0333011131132310-1233301100112220) |
| `disable` | [disable](resources--workload_flavor--reference--group-001.md#canonical-3012123303321311-0131100212111303-2003123013112302-0011101230333220-0123120031301323-2233100111003131-0011231122223111-2131301123223003) |
| `ephemeral_storage` | [ephemeral_storage](resources--workload_flavor--reference--group-001.md#canonical-3021101212313132-2133300232221211-3133123220310310-2031211011131203-2003003010321120-1202023211131033-2303003000011031-2232213233222032) |
| `id` | [ID](resources--workload_flavor--reference--group-001.md#canonical-3200203132323110-3003202021232031-0213131301110223-2110031313023022-0303001201233303-1100320323003111-2221110103233120-2103010213133230) |
| `labels` | [labels](resources--workload_flavor--reference--group-001.md#canonical-2101321210233303-2323212112123232-1321121232132122-0212201010023213-2031311311202010-2333133010013300-0221211211311330-0331320131123320) |
| `memory` | [memory](resources--workload_flavor--reference--group-001.md#canonical-2222321003220310-1113312123000210-1232313200321212-1112013122311311-2331030031320012-1302313230311312-1213001323333013-3210132321112003) |
| `name` | [name](resources--workload_flavor--reference--group-001.md#canonical-3101123100221310-3323002021232222-1010221220213312-1113213033031211-2213120303221002-0311120302222222-0113012310122203-0012212223001010) |
| `namespace` | [namespace](resources--workload_flavor--reference--group-001.md#canonical-2031202002100331-0011331010230212-2203121322302300-0101230011332300-0210001002122312-2000122222100112-0110031030112310-0333332313320331) |
| `timeouts` | [timeouts](resources--workload_flavor--reference--group-001.md#canonical-3213131000233332-3201102313330330-0013000032200232-2022321033220322-0222103130313311-2011123111313010-0120210012233013-1020121020110333) |
| `timeouts.create` | [timeouts.create](resources--workload_flavor--reference--group-001.md#canonical-3300203112123320-3132320303302320-1131312110221232-0012302212022002-1210120021030330-3220123331212312-2302122002200321-2303023200321311) |
| `timeouts.delete` | [timeouts.delete](resources--workload_flavor--reference--group-001.md#canonical-3033323133301133-2113302301331310-2011331232002011-3131113133311323-1001320212001102-3303122113200332-0133001032010311-2311201330010002) |
| `timeouts.read` | [timeouts.read](resources--workload_flavor--reference--group-001.md#canonical-2023311020310033-3023210003321220-1231310022300200-1200213010330321-3331230320300331-0113332111311200-0011011212003123-2302112321221033) |
| `timeouts.update` | [timeouts.update](resources--workload_flavor--reference--group-001.md#canonical-2102210312020301-3102320330330113-3220312003233333-2021030320311310-3133021113012322-3332031333231120-0203321112321102-2130222332122113) |
| `vcpus` | [vcpus](resources--workload_flavor--reference--group-001.md#canonical-2300132331222310-1031020230030230-2330231023023321-1311203001321013-1202311032323121-2212022322111022-2310132112331212-2213013331231121) |

<a id="canonical-3232213233303102-3110020321220110-3223332200322301-0333130201321131-1331230010022023-3311301302101312-1221022003323311-1021200120010202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001)
- [Property reference](resources--workload_flavor--reference--group-001.md#canonical-2201230231331130-0233333320310023-2033232302300302-2210032310122123-0221311000031233-1020311213100203-1113130201012033-0300212122012213)
- timeouts

<a id="canonical-3213131000233332-3201102313330330-0013000032200232-2022321033220322-0222103130313311-2011123111313010-0120210012233013-1020121020110333"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200002101030330-0323313230131332-1320131031032223-2231023322120012-3123111003131301-1131023220111331-0023322012211213-3232303202132333"></a>

### Direct properties for `timeouts`

<a id="canonical-3300203112123320-3132320303302320-1131312110221232-0012302212022002-1210120021030330-3220123331212312-2302122002200321-2303023200321311"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3033323133301133-2113302301331310-2011331232002011-3131113133311323-1001320212001102-3303122113200332-0133001032010311-2311201330010002"></a>

<a id="canonical-0320323323131013-2012232303202121-1322030323332222-0030020032012330-1211322312002321-1033031230211213-0033330021301023-1022121030203322"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2023311020310033-3023210003321220-1231310022300200-1200213010330321-3331230320300331-0113332111311200-0011011212003123-2302112321221033"></a>

<a id="canonical-0010321303332110-0012120303110002-3133022022231103-1330000231302313-2030302311303031-2100100020220222-2000030311221032-2123311322313022"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2102210312020301-3102320330330113-3220312003233333-2021030320311310-3133021113012322-3332031333231120-0203321112321102-2130222332122113"></a>

<a id="canonical-0020000213322113-0303210222113002-2222330111222221-2213013231320132-1320202213223102-3101200312010110-3120130013013023-2130011303302230"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
