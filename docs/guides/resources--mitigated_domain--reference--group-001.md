---
page_title: "xcsh_mitigated_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain reference."
---

# xcsh_mitigated_domain reference

<a id="canonical-3120221331031233-2101223210101003-1020302123022020-1212033230313332-1110320212031320-1030311212231013-0201130022100233-0012321303020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201203121112333-2133111212003033-1303220022123111-3332202300010202-2230302333212210-0031322311322031-2030332212130101-1230301110022223"></a>

## Property reference — Property reference / 331130333133 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102)
- Property reference

<a id="canonical-3310000302330331-0100032330103101-3033322303332012-3023010300131032-2100310120011012-3321021323003321-0220223313120321-1031033302313311"></a>

## Direct properties — Property reference / 331130333133 / 3

<a id="canonical-0211230112331030-1023131211112313-2221111033032301-3003303333123120-1300211001013130-3220212210200321-2200221310332230-0233023223213100"></a>

<a id="canonical-2310322213222122-0133323003003020-3101032211331132-2320110002322202-2022001030123201-1022131031320012-1322203213132310-0033102312113222"></a>

## annotations property — Property reference / 331130333133 / 4

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

<a id="canonical-1321323030021012-0330012000012032-3123220203320322-3000221331121023-1230310312301031-0320333001121231-1203213023020101-1201310201002002"></a>

<a id="canonical-0131002222122032-3023202313101102-0033020121220013-3113000111010233-3202220113030311-3112121131300300-0313131132122231-3131002332100230"></a>

## description property — Property reference / 331130333133 / 5

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

<a id="canonical-1211213200011213-1012311202122101-3223303212030233-2301020133021012-2112100312000110-2131010233131032-2023031331320323-0022301022300220"></a>

<a id="canonical-3131101012121132-0333123123213111-3020311020003333-3001102033000011-1120331110213303-2310301101003200-3312122331113002-3310131231130003"></a>

## disable property — Property reference / 331130333133 / 6

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

<a id="canonical-1020132231030223-2333122022001011-2020101333203100-2220223020123111-3202130000113100-0313032121333323-1233032022230231-2102313012231202"></a>

<a id="canonical-3111331332231332-0331100132323132-2300211300231312-1220212322300303-1310331302313012-2332033211303302-1131011200223103-0102231201001233"></a>

## ID property — Property reference / 331130333133 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3100100110313223-1303332311210003-0203333003322323-2030231331301002-2012320103310033-2220310133333000-2023232301213120-2233023313220330"></a>

<a id="canonical-2033210111200010-0111222022331230-2201020020012202-3223210211120301-3331111020200223-3220303213000211-0232232002301110-0221312222230132"></a>

## labels property — Property reference / 331130333133 / 8

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

<a id="canonical-0301230133223221-2112312300202030-1311010230331012-3132230132020311-2111321211121010-1010233131022133-0223120031122021-1202201112300221"></a>

<a id="canonical-3301021103101211-1300220203220310-0133120003323011-0000131310312310-0331332323333001-3120103222321223-0133220131131301-0100032221013303"></a>

## mitigated_domain property — Property reference / 331130333133 / 9

Type: `"string"`. Required.

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry.

Upstream description:

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry. Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

Provider validators and defaults (from schema source):

```go
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2231110102202303-0133133330120313-1032232132221200-2011020033013132-0303212000230200-2000120313130320-3130322321132011-2120323212103012"></a>

<a id="canonical-1322132213023122-1100021130201011-1333010303123032-1230212232311110-0202000032202211-3302213331332012-1213022233311031-0302331001130021"></a>

## name property — Property reference / 331130333133 / 10

Type: `"string"`. Required.

Name of the Mitigated Domain. Must be unique within the namespace.

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

<a id="canonical-0102323132333022-3102012100102100-2013101101303231-3310112233221200-1203013301310310-0133121020030233-1331323222320301-1222222003111122"></a>

<a id="canonical-2201001000333201-1032220223103213-2231022331333231-1232313201330120-1311301002301033-2120322220233013-2021000223023131-3131333331322003"></a>

## namespace property — Property reference / 331130333133 / 11

Type: `"string"`. Required.

Namespace where the Mitigated Domain is created.

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

- [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-2200113210300321-2223123023101333-2312000202010111-3123233132121012-3132013113200221-3202220130211003-3023230020111021-2001302031332101): complete subsection reference.

<a id="canonical-3121130232113221-3302011330322033-2131121123003211-0232103221313320-2111320200021221-3021301003112012-1112132130213330-1231013031221302"></a>

## All schema paths — Property reference / 331130333133 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--mitigated_domain--reference--group-001.md#canonical-0211230112331030-1023131211112313-2221111033032301-3003303333123120-1300211001013130-3220212210200321-2200221310332230-0233023223213100) |
| `description` | [description](resources--mitigated_domain--reference--group-001.md#canonical-1321323030021012-0330012000012032-3123220203320322-3000221331121023-1230310312301031-0320333001121231-1203213023020101-1201310201002002) |
| `disable` | [disable](resources--mitigated_domain--reference--group-001.md#canonical-1211213200011213-1012311202122101-3223303212030233-2301020133021012-2112100312000110-2131010233131032-2023031331320323-0022301022300220) |
| `id` | [id](resources--mitigated_domain--reference--group-001.md#canonical-1020132231030223-2333122022001011-2020101333203100-2220223020123111-3202130000113100-0313032121333323-1233032022230231-2102313012231202) |
| `labels` | [labels](resources--mitigated_domain--reference--group-001.md#canonical-3100100110313223-1303332311210003-0203333003322323-2030231331301002-2012320103310033-2220310133333000-2023232301213120-2233023313220330) |
| `mitigated_domain` | [mitigated_domain](resources--mitigated_domain--reference--group-001.md#canonical-0301230133223221-2112312300202030-1311010230331012-3132230132020311-2111321211121010-1010233131022133-0223120031122021-1202201112300221) |
| `name` | [name](resources--mitigated_domain--reference--group-001.md#canonical-2231110102202303-0133133330120313-1032232132221200-2011020033013132-0303212000230200-2000120313130320-3130322321132011-2120323212103012) |
| `namespace` | [namespace](resources--mitigated_domain--reference--group-001.md#canonical-0102323132333022-3102012100102100-2013101101303231-3310112233221200-1203013301310310-0133121020030233-1331323222320301-1222222003111122) |
| `timeouts` | [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-1330111212223300-2122011303221311-0231002010033011-0122001003120020-1131232212003112-0031132300132231-3001313323112011-2231200222303233) |
| `timeouts.create` | [timeouts.create](resources--mitigated_domain--reference--group-001.md#canonical-3013232221013122-2232202323112022-1002311320300303-0100122231011302-1320330101112312-1300102213331122-1022330332312212-3032320131113001) |
| `timeouts.delete` | [timeouts.delete](resources--mitigated_domain--reference--group-001.md#canonical-0300200103212333-2132213032230223-2101010202002320-2300220303031032-3133200003300122-2310313100132012-2113332111112012-0101210333321202) |
| `timeouts.read` | [timeouts.read](resources--mitigated_domain--reference--group-001.md#canonical-2301220232212031-3130100322312030-2010013013122102-2311211122221330-3123300331233302-0102312010032112-3222113310312201-1301112022331110) |
| `timeouts.update` | [timeouts.update](resources--mitigated_domain--reference--group-001.md#canonical-0331222030101331-1032000201002100-3021120033333112-1313030020301010-2000021320211321-0123003012012121-2031211302021303-2123021210131001) |

<a id="canonical-2332022120202132-0120033022333321-0033201002331011-3223310013000001-1331300213011013-0331302200001311-3012232112113201-2031132112002213"></a>

## Next pages — Property reference / 331130333133 / 13

- [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-2200113210300321-2223123023101333-2312000202010111-3123233132121012-3132013113200221-3202220130211003-3023230020111021-2001302031332101)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102)

<a id="canonical-2200113210300321-2223123023101333-2312000202010111-3123233132121012-3132013113200221-3202220130211003-3023230020111021-2001302031332101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221211212210333-2211233121113302-3223213323000121-0330010200132323-3223031132313101-0311212231333322-2312233120221322-3110123102023022"></a>

## timeouts — timeouts / 132220013330 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102)
- [Property reference](resources--mitigated_domain--reference--group-001.md#canonical-3120221331031233-2101223210101003-1020302123022020-1212033230313332-1110320212031320-1030311212231013-0201130022100233-0012321303020210)
- timeouts

<a id="canonical-1330111212223300-2122011303221311-0231002010033011-0122001003120020-1131232212003112-0031132300132231-3001313323112011-2231200222303233"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023300022330223-0220322223113022-0030300203012233-0220203323203103-3223000100131132-1110031211132310-0211012020123312-1033213313103021"></a>

## Direct properties — timeouts / 132220013330 / 3

<a id="canonical-3013232221013122-2232202323112022-1002311320300303-0100122231011302-1320330101112312-1300102213331122-1022330332312212-3032320131113001"></a>

<a id="canonical-2123003300130300-2320013232110212-2003223032313112-3322220310300230-2022311332213033-3122002231220023-3211110132301120-3211101321210312"></a>

## create property — timeouts / 132220013330 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0300200103212333-2132213032230223-2101010202002320-2300220303031032-3133200003300122-2310313100132012-2113332111112012-0101210333321202"></a>

<a id="canonical-1322113131220023-1330121301113200-2333122210101310-3303222223331102-0212230032133133-1210013023110100-0320032300023331-2120021021003201"></a>

## delete property — timeouts / 132220013330 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2301220232212031-3130100322312030-2010013013122102-2311211122221330-3123300331233302-0102312010032112-3222113310312201-1301112022331110"></a>

<a id="canonical-3110000323011321-1002121333300222-2213303201302023-1320333321001110-0333111212030303-0120202122110213-2132100131100310-3331332320022103"></a>

## read property — timeouts / 132220013330 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0331222030101331-1032000201002100-3021120033333112-1313030020301010-2000021320211321-0123003012012121-2031211302021303-2123021210131001"></a>

<a id="canonical-0020223202032001-3211011020310231-1111213211231220-1330222101032222-1023012110302210-3020232113121313-3030301212203113-2032031100312320"></a>

## update property — timeouts / 132220013330 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0123213111112010-3113002203231233-1101103212312120-0110302210220013-3102122212311233-1300302210203311-3012101323101222-0210233112201303"></a>

## Next pages — timeouts / 132220013330 / 8

- [Property reference](resources--mitigated_domain--reference--group-001.md#canonical-3120221331031233-2101223210101003-1020302123022020-1212033230313332-1110320212031320-1030311212231013-0201130022100233-0012321303020210)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102)
