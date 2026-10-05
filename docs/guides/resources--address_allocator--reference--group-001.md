---
page_title: "xcsh_address_allocator reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator reference."
---

# xcsh_address_allocator reference

<a id="canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201002103212200-3121030323310122-0121103233202320-2231221113011302-2113301120000300-3012023121332332-1031023320132222-1232330103130223"></a>

## Property reference — Property reference / 331112023300 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
- Property reference

<a id="canonical-2113300131320033-3223203010023102-2112033103021320-1332220211112221-1110131300231102-1220123012221030-1210003313213312-1331220113203100"></a>

## Direct properties — Property reference / 331112023300 / 3

- [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-3230300123333023-3021311210200313-1101220223020003-1031221103211013-2001322121201301-0211230330330103-2030331033011233-2322001313001032): complete subsection reference.

<a id="canonical-3332010132321031-0032111101231013-1031203301132210-0003013021221012-2330302012220200-1210100130221021-2333101320213012-3032033333031133"></a>

<a id="canonical-1322132122022222-3200233011333023-3300303231113320-0301310320012213-2120103301030032-0103323211122333-0302011213330323-0310331331032012"></a>

## address_pool property — Property reference / 331112023300 / 4

Type: `["list", "string"]`. Required.

Address pool from which the allocator carves out subnets or addresses to its clients.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2233321221301121-0202110131211233-1002223020213132-0021100132010312-3233020222000330-2222301230100323-3322302200303100-1221031013211011"></a>

<a id="canonical-0203311102133222-2103211202201320-2232302332113231-0121230010201121-1000121003213110-2121101010320212-1120130230132310-2332100211110031"></a>

## annotations property — Property reference / 331112023300 / 5

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

<a id="canonical-3300333301313212-1100210331233331-2232100112312300-1312213101332120-0010311231021211-1222322023031300-0130032113100000-3020031013222010"></a>

<a id="canonical-1002301301223301-0303010233322312-1211023323303103-3001220001122200-2312010120013300-0033321223002121-2212112213012201-2330031002030033"></a>

## description property — Property reference / 331112023300 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1211013200301211-2030311011110100-3211132203200023-3112010122113131-1013021022020213-2003031200233021-0233111323233103-2031103030031123"></a>

<a id="canonical-2302223132101200-2122003100110300-1222332133220232-0023132120111031-0302233131232032-1121203100023003-1131002120112200-3322132123332023"></a>

## disable property — Property reference / 331112023300 / 7

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

<a id="canonical-2301101130223003-2132011301100003-3001033001113311-2000130320112103-2232022202202103-1333033103313330-1330020210313112-2221231032022212"></a>

<a id="canonical-3210002323200213-2221032123232320-3232302211003121-1100202111000230-3030103121333310-0001112010131210-0011201210111022-0330201200132130"></a>

## ID property — Property reference / 331112023300 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1213202000020200-1121033322023202-3330321233131313-1103131032013231-1121031022313211-3121321212023211-1030310012330313-3011000333010010"></a>

<a id="canonical-1112100022332202-2003000332331230-3102030313123001-1231302323022012-1101132101113003-2011021222213223-3001201310202131-1011201303221333"></a>

## labels property — Property reference / 331112023300 / 9

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

<a id="canonical-1311003202012030-0023200231320012-0333200323221132-3100310230303022-2113131003023130-0003300313332233-0230211220213303-1020232030013130"></a>

<a id="canonical-2310121222202331-2232002120102022-3103311130110102-1322122022231103-1132101013230011-2311013313302231-1321102031230203-0302002213132203"></a>

## mode property — Property reference / 331112023300 / 10

Type: `"string"`. Optional, Computed.

\[Enum: LOCAL|GLOBAL\_PER\_SITE\_NODE\] Mode of the address allocator Address allocator is for VERs
within the local cluster or site Allocation is per site and then per node. Possible values are
\`LOCAL\`, \`GLOBAL\_PER\_SITE\_NODE\`. Defaults to \`LOCAL\`.

Upstream description:

Mode of the address allocator

Address allocator is for VERs within the local cluster or site Allocation is per site and then per
node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["GLOBAL_PER_SITE_NODE","LOCAL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LOCAL",
    "GLOBAL_PER_SITE_NODE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL",
  "enum": [
    "LOCAL",
    "GLOBAL_PER_SITE_NODE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231210112102022-1130211230333123-3032303321010323-3320300132013133-2322213310203202-0011210301011013-2221100111332202-2110321021202010"></a>

<a id="canonical-3131011323122220-3023300033132232-2210032210121321-2121013222203132-2022221013100120-0010221313202332-2123131303133303-1133012333000213"></a>

## name property — Property reference / 331112023300 / 11

Type: `"string"`. Required.

Name of the Address Allocator. Must be unique within the namespace.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2322201302320003-2230011310311213-1121301001111211-0201032122021121-2321001310232320-1100332302121232-1030313311112212-0301301031201022"></a>

<a id="canonical-0210103332033220-3212333113022303-3033202011131333-0100123120000013-0110220111031011-0132123303210130-1320033302212013-2110313000132112"></a>

## namespace property — Property reference / 331112023300 / 12

Type: `"string"`. Required.

Namespace where the Address Allocator is created.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [timeouts](resources--address_allocator--reference--group-001.md#canonical-2320031132000030-2212003130031203-2230222130221323-1003101231321100-0022133332033021-1321101123002312-1001202113000200-0003221202311001): complete subsection reference.

<a id="canonical-2131323013103213-0110113212132301-2111122101123212-3121033230023111-0130322201033223-1001223102220012-3132300031312331-3121003322201211"></a>

## All schema paths — Property reference / 331112023300 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_allocation_scheme` | [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-0123330131320200-1301120233030303-3011331133130201-1102031020020020-1313111203101031-2030320002230001-0231230323010110-2010222302323030) |
| `address_allocation_scheme.allocation_unit` | [address_allocation_scheme.allocation_unit](resources--address_allocator--reference--group-001.md#canonical-1301012020012213-2333232322312013-1300102030011121-1330332202321030-2120003220003121-2113022013022233-3301213131113000-2012112103210123) |
| `address_allocation_scheme.local_interface_address_offset` | [address_allocation_scheme.local_interface_address_offset](resources--address_allocator--reference--group-001.md#canonical-2302332132013232-2113213003100303-0032011311101232-3132023101312321-2102132032033321-2011022101203023-1203101022103003-3011213222200021) |
| `address_allocation_scheme.local_interface_address_type` | [address_allocation_scheme.local_interface_address_type](resources--address_allocator--reference--group-001.md#canonical-0322312101330322-2111331011020210-1200231030221320-3332333111201330-2101322303201232-3111020320013333-2030301132130301-2100310321221132) |
| `address_pool` | [address_pool](resources--address_allocator--reference--group-001.md#canonical-3332010132321031-0032111101231013-1031203301132210-0003013021221012-2330302012220200-1210100130221021-2333101320213012-3032033333031133) |
| `annotations` | [annotations](resources--address_allocator--reference--group-001.md#canonical-2233321221301121-0202110131211233-1002223020213132-0021100132010312-3233020222000330-2222301230100323-3322302200303100-1221031013211011) |
| `description` | [description](resources--address_allocator--reference--group-001.md#canonical-3300333301313212-1100210331233331-2232100112312300-1312213101332120-0010311231021211-1222322023031300-0130032113100000-3020031013222010) |
| `disable` | [disable](resources--address_allocator--reference--group-001.md#canonical-1211013200301211-2030311011110100-3211132203200023-3112010122113131-1013021022020213-2003031200233021-0233111323233103-2031103030031123) |
| `id` | [ID](resources--address_allocator--reference--group-001.md#canonical-2301101130223003-2132011301100003-3001033001113311-2000130320112103-2232022202202103-1333033103313330-1330020210313112-2221231032022212) |
| `labels` | [labels](resources--address_allocator--reference--group-001.md#canonical-1213202000020200-1121033322023202-3330321233131313-1103131032013231-1121031022313211-3121321212023211-1030310012330313-3011000333010010) |
| `mode` | [mode](resources--address_allocator--reference--group-001.md#canonical-1311003202012030-0023200231320012-0333200323221132-3100310230303022-2113131003023130-0003300313332233-0230211220213303-1020232030013130) |
| `name` | [name](resources--address_allocator--reference--group-001.md#canonical-3231210112102022-1130211230333123-3032303321010323-3320300132013133-2322213310203202-0011210301011013-2221100111332202-2110321021202010) |
| `namespace` | [namespace](resources--address_allocator--reference--group-001.md#canonical-2322201302320003-2230011310311213-1121301001111211-0201032122021121-2321001310232320-1100332302121232-1030313311112212-0301301031201022) |
| `timeouts` | [timeouts](resources--address_allocator--reference--group-001.md#canonical-2101323213131112-0021020330100320-3003301021103130-0000232312230211-3300213021032011-0213020231111231-1203103311011130-1321331002201221) |
| `timeouts.create` | [timeouts.create](resources--address_allocator--reference--group-001.md#canonical-0200213311100032-2233220102301023-3100022102032002-2120331103321123-1202130302021121-3312132002230122-0332201121322033-2211332323031302) |
| `timeouts.delete` | [timeouts.delete](resources--address_allocator--reference--group-001.md#canonical-2000030022111023-2103021031023013-3121332321121201-1320030312332333-0031233202211101-0122232130300123-0213212303210113-3030321303023323) |
| `timeouts.read` | [timeouts.read](resources--address_allocator--reference--group-001.md#canonical-0033103101111032-2301030203131222-3111122321303210-2312321231020101-1201100030021211-1212102201130332-1120003331210011-1112323003201131) |
| `timeouts.update` | [timeouts.update](resources--address_allocator--reference--group-001.md#canonical-1230031200010133-2000110323002313-3213322213212013-3123310210332013-2303022123221111-2132200000023222-1123020023220001-1020023201223320) |

<a id="canonical-1122212323223102-2200231333112120-0220132311202213-0101012001312332-2312212023302130-3231020331220303-0210002022331002-1101323103032112"></a>

## Next pages — Property reference / 331112023300 / 14

- [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-3230300123333023-3021311210200313-1101220223020003-1031221103211013-2001322121201301-0211230330330103-2030331033011233-2322001313001032)
- [timeouts](resources--address_allocator--reference--group-001.md#canonical-2320031132000030-2212003130031203-2230222130221323-1003101231321100-0022133332033021-1321101123002312-1001202113000200-0003221202311001)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)

<a id="canonical-3230300123333023-3021311210200313-1101220223020003-1031221103211013-2001322121201301-0211230330330103-2030331033011233-2322001313001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111312210110331-0332203103233001-2121110323121300-1212223211300101-2230201020231023-0333113320113133-1311211021333113-3103320033011133"></a>

## address_allocation_scheme — address_allocation_scheme / 121203210110 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
- [Property reference](resources--address_allocator--reference--group-001.md#canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213)
- address_allocation_scheme

<a id="canonical-0123330131320200-1301120233030303-3011331133130201-1102031020020020-1313111203101031-2030320002230001-0231230323010110-2010222302323030"></a>

Type: `"object"`. single nested block, Optional.

Decides the scheme to be used to allocate addresses from the configured address pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("allocation_unit")}
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
address_allocation_scheme {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222212322220130-0300013231013113-1000033201031021-0132103210012110-0133223330121330-0233123233202223-0330231030323112-3230103111011033"></a>

## Direct properties — address_allocation_scheme / 121203210110 / 3

<a id="canonical-1301012020012213-2333232322312013-1300102030011121-1330332202321030-2120003220003121-2113022013022233-3301213131113000-2012112103210123"></a>

<a id="canonical-2212210321101012-2111033133122301-1021001220101110-0012223000301013-0313031202132001-3332023103221212-3121112311003013-2222222033001002"></a>

## allocation_unit property — address_allocation_scheme / 121203210110 / 4

Type: `"number"`. Optional.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Upstream description:

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2302332132013232-2113213003100303-0032011311101232-3132023101312321-2102132032033321-2011022101203023-1203101022103003-3011213222200021"></a>

<a id="canonical-3111223231030320-0023232131221203-2311203222012332-3322102231130002-0203021111020133-2312202120233322-1233331300311200-2000301133203332"></a>

## local_interface_address_offset property — address_allocation_scheme / 121203210110 / 5

Type: `"number"`. Optional.

Used to derive address for the local interface from the allocated subnet. If Local Interface Address
Type is set to 'Offset from beginning of Subnet', this offset value is added to the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24..

Upstream description:

This is used to derive address for the local interface from the allocated subnet.

If Local Interface Address Type is set to "Offset from beginning of Subnet", this offset value is
added to the allocated subnet and used as the local interface address. For example, if the allocated
subnet is 192.0.2.0/24 and offset is set to 2 with Local Interface Address Type set to "Offset from
beginning of Subnet", local interface address of 192.0.2.204 is used.

If Local Interface Address Type is set to "Offset from end of Subnet", this offset value is
subtracted from the end of the allocated subnet and used as the local interface address. For
example, if the allocated subnet is 192.0.2.0/24 and offset is set to 1 with Local Interface Address
Type set to "Offset from end of Subnet", local interface address of 192.0.2.204 is used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0322312101330322-2111331011020210-1200231030221320-3332333111201330-2101322303201232-3111020320013333-2030301132130301-2100310321221132"></a>

<a id="canonical-2311000233011313-1320010001123102-1203300112122212-1220013101112223-2332300203113222-3130321322220210-3110133232133200-1331033203132332"></a>

## local_interface_address_type property — address_allocation_scheme / 121203210110 / 6

Type: `"string"`. Optional.

\[Enum:
LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN|LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END|LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\]
Dictates how local interface address is derived from the allocated subnet Use Nth address of the
allocated subnet as the local interface address, N being the Local Interface Address Offset. For
example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and
Local.. Possible values are \`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`,
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END\`,
\`LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\`. Defaults to
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`.

Upstream description:

Dictates how local interface address is derived from the allocated subnet

Use Nth address of the allocated subnet as the local interface address, N being the Local Interface
Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset
is set to 2 and Local Interface Address Type is set to "Offset from beginning of Subnet", local
address of 192.0.2.204 is used.

Use Nth last address of the allocated subnet as the local interface address, N being the Local
Interface Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface
Address Offset is set to 1 and Local Interface Address Type is set to "Offset from end of Subnet",
local address of 192.0.2.204 is used.

This case is used for external\_connector.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOCAL_INTERFACE_ADDRESS_FROM_PREFIX","LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN","LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
  "enum": [
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000033303301331-2002100001123223-0300312323200231-0001201133031200-3200303301012111-3023112002102003-0313333212322223-1202101123010132"></a>

## Next pages — address_allocation_scheme / 121203210110 / 7

- [Property reference](resources--address_allocator--reference--group-001.md#canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)

<a id="canonical-2320031132000030-2212003130031203-2230222130221323-1003101231321100-0022133332033021-1321101123002312-1001202113000200-0003221202311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103023303233223-3230110301231031-0102221312130101-2301032121001323-2012122011300103-0230232003323023-2310032012322330-3332103002122021"></a>

## timeouts — timeouts / 021010133131 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
- [Property reference](resources--address_allocator--reference--group-001.md#canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213)
- timeouts

<a id="canonical-2101323213131112-0021020330100320-3003301021103130-0000232312230211-3300213021032011-0213020231111231-1203103311011130-1321331002201221"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112220321103102-0031233031222230-1322320300201121-1322332110103000-1211300200111230-0203330312300222-0211213302210000-1300000031312132"></a>

## Direct properties — timeouts / 021010133131 / 3

<a id="canonical-0200213311100032-2233220102301023-3100022102032002-2120331103321123-1202130302021121-3312132002230122-0332201121322033-2211332323031302"></a>

<a id="canonical-3211220320030203-0020001211300203-2320312020020230-1200222033013132-3233201111032113-2311112222222223-3031123122121232-0210110301330321"></a>

## create property — timeouts / 021010133131 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2000030022111023-2103021031023013-3121332321121201-1320030312332333-0031233202211101-0122232130300123-0213212303210113-3030321303023323"></a>

<a id="canonical-1110022032302002-2130021310222100-1221203313231113-1012303230032301-3212230031302230-0223111000112201-1233012321302122-0330130032022212"></a>

## delete property — timeouts / 021010133131 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0033103101111032-2301030203131222-3111122321303210-2312321231020101-1201100030021211-1212102201130332-1120003331210011-1112323003201131"></a>

<a id="canonical-2213211310030033-2010120023302330-1113210120003311-1121230211022022-0122110030013123-1012230022210203-3221130132311013-0232330103020233"></a>

## read property — timeouts / 021010133131 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1230031200010133-2000110323002313-3213322213212013-3123310210332013-2303022123221111-2132200000023222-1123020023220001-1020023201223320"></a>

<a id="canonical-2032203313031120-3102221130120201-0222222213330331-3030313302202210-2011210323031130-3121231213032213-0130323203003321-0301000032331211"></a>

## update property — timeouts / 021010133131 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0312001110201203-2203102231302330-1222103210122100-0322123321231330-3010123202130022-1112302203220301-1121021013311221-1032003110323322"></a>

## Next pages — timeouts / 021010133131 / 8

- [Property reference](resources--address_allocator--reference--group-001.md#canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
