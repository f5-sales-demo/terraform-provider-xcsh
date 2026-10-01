---
page_title: "xcsh_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer reference."
---

# xcsh_policer reference

<a id="canonical-0211330131013020-2310303133111013-2220323023232330-2223132030123020-2011122031201112-2310200001103220-1221101020202030-3311103200011300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020232212210111-0111000003320131-2313123331013003-3203003013123202-0003121031112003-2230032010020322-2023301203133220-2021303133130220"></a>

## Property reference — Property reference / 113031000123 / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)
- Property reference

<a id="canonical-1133032103132020-3322111332301202-0030233333022133-3330133123030213-3221303100002222-1122330333332033-1121303022013100-0031110011133212"></a>

## Direct properties — Property reference / 113031000123 / 3

<a id="canonical-1211031001320321-3030103113211320-1333103233202113-1301111210010232-3213211020000220-1331010102220031-2232130021330203-2033121331133323"></a>

<a id="canonical-0000102310033121-1131231122000300-3201213210210020-3112331220312200-0320022332200332-0121111030022102-3221111312300303-0201312313002123"></a>

## annotations property — Property reference / 113031000123 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

<a id="canonical-0320032300221101-0202202313001110-2312023211011133-3220111000002303-2013213312311312-2300101230203012-3120202212203331-3103030203301301"></a>

<a id="canonical-1002000103032301-1023303033330011-0123230032000222-0210111022001130-1300303212032030-2300000303330030-0303110102002110-3212110200020132"></a>

## burst_size property — Property reference / 113031000123 / 5

Type: `"number"`. Required.

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Upstream description:

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

<a id="canonical-1123011010113210-0230323311102230-2310323112321230-0011213133112210-1100201112033312-2310001200210011-0330100022301100-2112302130212231"></a>

<a id="canonical-2212032301120123-3303222000321212-3110222002133102-2333113002021003-2010320333021012-2120232222201000-1310030121322202-1232122221230032"></a>

## committed_information_rate property — Property reference / 113031000123 / 6

Type: `"number"`. Required.

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Upstream description:

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 10000000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10000000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  }
}
```

<a id="canonical-2211330112230121-0102210021031220-3132311021311000-0130320123023200-1313221100300220-0030101000020220-3110131312103010-1311303100020202"></a>

<a id="canonical-3111231333001120-2330010101010020-0321022032201121-3103103121032203-1111102121101310-2200333333011302-1120001102000331-3001323012313023"></a>

## description property — Property reference / 113031000123 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2202313012130320-3033221132101332-3220001310232132-2213313133012230-2313002203133131-2013003300223313-1032032302300013-2221122010121011"></a>

<a id="canonical-0113211210130030-1132221023221322-1211003333133202-0212133333031101-3122301120033323-3213000331303112-1110221201133121-3303132112200200"></a>

## disable property — Property reference / 113031000123 / 8

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

<a id="canonical-0112012300232013-2321220130121020-2303333132031331-3030103303003200-3201220331203133-1030120212131321-1032333230121212-1201010100320103"></a>

<a id="canonical-1223322333100020-2121223112330323-1012131313200112-2102120313331301-1212122333020222-1010121101323203-0020323202212201-1231030030120233"></a>

## ID property — Property reference / 113031000123 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1113321021100121-1103203011201231-2012130201222323-2123023113231113-0013031301221230-1313031312213202-2213332122110310-1110333312103030"></a>

<a id="canonical-3311331333231232-2321331023322221-2213312010210333-2202033333011110-3030103131331000-0002010200123012-3310013221322300-3131012332302320"></a>

## labels property — Property reference / 113031000123 / 10

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Upstream description:

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

<a id="canonical-0030133120033221-0323210222203212-3103302100131020-3322110032221220-0032330322203303-3230313300100020-0132220110020330-3202321231122003"></a>

<a id="canonical-1102333122333002-3101332112111200-3223220211231323-3323333201003121-2203132022003311-0101231030321032-2333121122222012-3333303020003121"></a>

## name property — Property reference / 113031000123 / 11

Type: `"string"`. Required.

Name of the Policer. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1202321012112223-1031102102211211-1303101323313321-1331012031030123-2310311030022302-3112120212212000-1010132330113203-3200132303310213"></a>

<a id="canonical-2333322312300331-2130213302012301-2212232332122303-3330202201230312-3123101100331321-0231102011002003-1033032332333100-0130023332330021"></a>

## namespace property — Property reference / 113031000123 / 12

Type: `"string"`. Required.

Namespace where the Policer is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1103030223023111-0231123311100330-0320323120121112-0203300131133133-1330112303210000-3011232131202220-2100001332100200-2022212033301003"></a>

<a id="canonical-2111200200200300-2010123303033301-3022323130330033-3131020002202030-3220000231012010-3201010223132032-0110010113203212-2321302033103010"></a>

## policer_mode property — Property reference / 113031000123 / 13

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_MODE\_NOT\_SHARED|POLICER\_MODE\_SHARED\] - POLICER\_MODE\_NOT\_SHARED: Not Shared
A separate policer instance is created for each reference to the policer - POLICER\_MODE\_SHARED:
Shared A common policer instance is used for for all references to the policer. Possible values are
\`POLICER\_MODE\_NOT\_SHARED\`, \`POLICER\_MODE\_SHARED\`. Defaults to
\`POLICER\_MODE\_NOT\_SHARED\`. Server applies default when omitted.

Upstream description:

&#8203;- POLICER\_MODE\_NOT\_SHARED: Not Shared

A separate policer instance is created for each reference to the policer &#8203;-
POLICER\_MODE\_SHARED: Shared

A common policer instance is used for for all references to the policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_MODE_NOT_SHARED",
  "enum": [
    "POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322021012230232-0333313212331122-2332331221102002-2033300130233332-0333103203120313-3213202020321223-0211331332203132-2002023001223123"></a>

<a id="canonical-0013330212220032-1131032030033121-3032301103302000-0132333132201102-2311023112330101-3331122000230131-1112323331211313-3233010111111211"></a>

## policer_type property — Property reference / 113031000123 / 14

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_SINGLE\_RATE\_TWO\_COLOR\] Specifies the type of Policer Basic Single-Rate
Two-Color Policer. The only possible value is \`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Defaults to
\`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Server applies default when omitted.

Upstream description:

Specifies the type of Policer

Basic Single-Rate Two-Color Policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_SINGLE_RATE_TWO_COLOR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_SINGLE_RATE_TWO_COLOR",
  "enum": [
    "POLICER_SINGLE_RATE_TWO_COLOR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--policer--reference--group-001.md#canonical-1123323232130322-2233220022020310-3310111323013231-2100101120012233-2210131012331022-3322231111021120-1130013003313101-2311231312201013): complete subsection reference.

<a id="canonical-0303011030133031-2113100120222100-2023231020011103-1321112011210211-3002203202311130-3123210101310310-2332021102132323-2113101202003330"></a>

## All schema paths — Property reference / 113031000123 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--policer--reference--group-001.md#canonical-1211031001320321-3030103113211320-1333103233202113-1301111210010232-3213211020000220-1331010102220031-2232130021330203-2033121331133323) |
| `burst_size` | [burst_size](resources--policer--reference--group-001.md#canonical-0320032300221101-0202202313001110-2312023211011133-3220111000002303-2013213312311312-2300101230203012-3120202212203331-3103030203301301) |
| `committed_information_rate` | [committed_information_rate](resources--policer--reference--group-001.md#canonical-1123011010113210-0230323311102230-2310323112321230-0011213133112210-1100201112033312-2310001200210011-0330100022301100-2112302130212231) |
| `description` | [description](resources--policer--reference--group-001.md#canonical-2211330112230121-0102210021031220-3132311021311000-0130320123023200-1313221100300220-0030101000020220-3110131312103010-1311303100020202) |
| `disable` | [disable](resources--policer--reference--group-001.md#canonical-2202313012130320-3033221132101332-3220001310232132-2213313133012230-2313002203133131-2013003300223313-1032032302300013-2221122010121011) |
| `id` | [id](resources--policer--reference--group-001.md#canonical-0112012300232013-2321220130121020-2303333132031331-3030103303003200-3201220331203133-1030120212131321-1032333230121212-1201010100320103) |
| `labels` | [labels](resources--policer--reference--group-001.md#canonical-1113321021100121-1103203011201231-2012130201222323-2123023113231113-0013031301221230-1313031312213202-2213332122110310-1110333312103030) |
| `name` | [name](resources--policer--reference--group-001.md#canonical-0030133120033221-0323210222203212-3103302100131020-3322110032221220-0032330322203303-3230313300100020-0132220110020330-3202321231122003) |
| `namespace` | [namespace](resources--policer--reference--group-001.md#canonical-1202321012112223-1031102102211211-1303101323313321-1331012031030123-2310311030022302-3112120212212000-1010132330113203-3200132303310213) |
| `policer_mode` | [policer_mode](resources--policer--reference--group-001.md#canonical-1103030223023111-0231123311100330-0320323120121112-0203300131133133-1330112303210000-3011232131202220-2100001332100200-2022212033301003) |
| `policer_type` | [policer_type](resources--policer--reference--group-001.md#canonical-2322021012230232-0333313212331122-2332331221102002-2033300130233332-0333103203120313-3213202020321223-0211331332203132-2002023001223123) |
| `timeouts` | [timeouts](resources--policer--reference--group-001.md#canonical-2202230223300322-0211333111300132-2023110213102022-0300202002021323-0231002022130101-1011302320300231-2211103000333133-1301100123113110) |
| `timeouts.create` | [timeouts.create](resources--policer--reference--group-001.md#canonical-0111331033121111-0320320002000233-0113101122010111-1302013203323332-1231131013330102-3201013213110332-3300023013121231-1300202012012012) |
| `timeouts.delete` | [timeouts.delete](resources--policer--reference--group-001.md#canonical-1330231220313321-1323321030321323-0022020331010311-0212133022303220-2030111323210111-3230201323320302-2223122123311110-3000023002023021) |
| `timeouts.read` | [timeouts.read](resources--policer--reference--group-001.md#canonical-1331020312031002-0010200211021001-1213030123020332-3310121032102113-0231200002112223-0233121202212212-2201111130102313-0030213301322320) |
| `timeouts.update` | [timeouts.update](resources--policer--reference--group-001.md#canonical-2332212221032331-1233123033033331-3123002103203013-2020120213131021-2200131231203112-1130311030330330-1212002132030200-0232213303322332) |

<a id="canonical-2300132131303113-2321130202002001-2301110232122033-2013120231232202-0030103100320223-3322330330312201-2031320013023303-1200211012301202"></a>

## Next pages — Property reference / 113031000123 / 16

- [timeouts](resources--policer--reference--group-001.md#canonical-1123323232130322-2233220022020310-3310111323013231-2100101120012233-2210131012331022-3322231111021120-1130013003313101-2311231312201013)
- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)

<a id="canonical-1123323232130322-2233220022020310-3310111323013231-2100101120012233-2210131012331022-3322231111021120-1130013003313101-2311231312201013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103133232123110-1122101112002012-0223311131000002-0033332102123323-2223121023033123-2101000112310201-2332321101031033-3330301020110030"></a>

## timeouts — timeouts / 131312013202 / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)
- [Property reference](resources--policer--reference--group-001.md#canonical-0211330131013020-2310303133111013-2220323023232330-2223132030123020-2011122031201112-2310200001103220-1221101020202030-3311103200011300)
- timeouts

<a id="canonical-2202230223300322-0211333111300132-2023110213102022-0300202002021323-0231002022130101-1011302320300231-2211103000333133-1301100123113110"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210201231310210-0200023323133213-3103320212121102-1032000201112210-3211031001110122-1201230001202001-2023220302032303-3133101312301033"></a>

## Direct properties — timeouts / 131312013202 / 3

<a id="canonical-0111331033121111-0320320002000233-0113101122010111-1302013203323332-1231131013330102-3201013213110332-3300023013121231-1300202012012012"></a>

<a id="canonical-0230311030022202-1211010310322201-1023032301202032-3230203331101012-1223332020133201-2210001023111310-1013032332032033-1022133113230303"></a>

## create property — timeouts / 131312013202 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1330231220313321-1323321030321323-0022020331010311-0212133022303220-2030111323210111-3230201323320302-2223122123311110-3000023002023021"></a>

<a id="canonical-1000031101013011-2133033202220332-0132010300222202-3300131213323212-0003303212222120-0331223001112132-0202012100022202-2022012201212323"></a>

## delete property — timeouts / 131312013202 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1331020312031002-0010200211021001-1213030123020332-3310121032102113-0231200002112223-0233121202212212-2201111130102313-0030213301322320"></a>

<a id="canonical-2100003323000122-1021030000020201-3001030021321123-1330310031221122-3033133102031322-3003123110212110-0131133010202123-3031002313203313"></a>

## read property — timeouts / 131312013202 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2332212221032331-1233123033033331-3123002103203013-2020120213131021-2200131231203112-1130311030330330-1212002132030200-0232213303322332"></a>

<a id="canonical-3300230223113023-0312311120020102-0003232120003012-1200122310132122-0322112220233331-0032333030312101-2010212220310203-0100211101332303"></a>

## update property — timeouts / 131312013202 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3302212020133033-2023010301300332-2322232120113111-1012320311002031-3032333021300323-0331332231001300-3023222030320222-1330231100112301"></a>

## Next pages — timeouts / 131312013202 / 8

- [Property reference](resources--policer--reference--group-001.md#canonical-0211330131013020-2310303133111013-2220323023232330-2223132030123020-2011122031201112-2310200001103220-1221101020202030-3311103200011300)
- [xcsh_policer](../resources/policer.md#canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113)
