---
page_title: "xcsh_srv6_network_slice reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice reference."
---

# xcsh_srv6_network_slice reference

<a id="canonical-2221210201331300-0011233131200112-1301201030112212-2211123330221021-3333331233302122-0320123310310202-3200030231001021-3012033330201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323021323302032-0213321211321131-1310323113211320-2022030300011301-0200312023013030-1322030231332103-0103201113223130-3212102021130121"></a>

## Property reference — Property reference / 221023103012 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)
- Property reference

<a id="canonical-0132302001230001-3322223201202220-0113333200120102-1030221103233330-1223323103120011-0130100133322212-1212031230220121-0130133303011223"></a>

## Direct properties — Property reference / 221023103012 / 3

<a id="canonical-3223123301302320-2220102122230030-2212031313122121-2232211021103121-1012010311013230-2132123223132013-3331002022123233-2211111331010022"></a>

<a id="canonical-0220321213212110-2130211023202201-1130232221112322-3313201013013132-0120313313002331-0332330221120132-1301333023322012-3331230020312331"></a>

## annotations property — Property reference / 221023103012 / 4

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

<a id="canonical-2322133310321032-3212210130310201-2010033232332202-3202010200011023-2122223032103311-1321123101003120-3231032110101330-0320230210231012"></a>

<a id="canonical-2211101311332001-0121301313102101-2320103233022031-0030133003233130-2000103313033012-1313330322011100-2222321101033221-2122030132331333"></a>

## connect_to_access_networks property — Property reference / 221023103012 / 5

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

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

<a id="canonical-0333010203300212-2211112212330302-2321121311220122-2200010131122233-3022200001321012-1233230003022332-1210312021201211-0333112110133101"></a>

<a id="canonical-3021023123333133-3330120331131010-3203200102201230-2212010000103032-1003230211330200-3102001333010231-0203213210310212-3313323231003011"></a>

## connect_to_enterprise_networks property — Property reference / 221023103012 / 6

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

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

<a id="canonical-1023020031310102-2113000013012312-3101322122002130-1230302322211312-3021312311001121-3231101201301033-3301331332311210-0320131310022210"></a>

<a id="canonical-2303331200310021-0103112332310320-0312003013022131-0322121023201011-0302120022113110-3201032111211310-0133012323221303-0110010100200101"></a>

## connect_to_internet property — Property reference / 221023103012 / 7

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

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

<a id="canonical-1303212022101203-3030021121332023-2021002331323030-1131311331002001-1300122301000311-1231010122330022-1331211320023120-2003301012211322"></a>

<a id="canonical-1220331312002212-2013021030011123-0120222323223233-0130212222100120-1112220211321333-1322301220202132-1123222323130110-2331220231220321"></a>

## description property — Property reference / 221023103012 / 8

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

<a id="canonical-2330322132023333-0032212332130313-1202102031321031-1021020322222313-0220010112130023-2330130120023012-0230232102002311-1312211103321122"></a>

<a id="canonical-3033230321000012-3131131312100211-3130031131101320-0303023311301021-2123120232000133-0032301221300000-3011332130020020-3133230231231110"></a>

## disable property — Property reference / 221023103012 / 9

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

<a id="canonical-3233021201223011-2122022303112100-2030330021020111-0111333203300231-0023112121012033-1032330133232230-2102333320230032-1020120101131012"></a>

<a id="canonical-3223133223132331-1131033210203012-2223123220223312-3031010211013323-1201123233010332-2121331132010112-0030031010120211-2120120202232112"></a>

## ID property — Property reference / 221023103012 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0003210122312333-3233203011221222-3200203310233201-0302110223201000-0232013002110111-1322213201123200-0313213133203322-2313100123031020"></a>

<a id="canonical-0223231020011102-3120130031213003-3233022103110133-3000122132313230-2100303320131233-2120223211333331-1302012012002201-2110021123112231"></a>

## labels property — Property reference / 221023103012 / 11

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

<a id="canonical-3301211003020201-0231323000031211-0321230113022330-1300331201302333-3332301023103011-1010012221010300-3112121232301232-2320000021223113"></a>

<a id="canonical-2230100331331210-2332010333311320-3323232133003222-3301223132321131-1020100013332112-3232302130100320-2221032312333333-0331032003311200"></a>

## name property — Property reference / 221023103012 / 12

Type: `"string"`. Required.

Name of the Srv6 Network Slice. Must be unique within the namespace.

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

<a id="canonical-3321023001110032-0011232322102033-3321100020012201-3201232001000001-1220031310112200-3020022011020031-0303120231022202-3123231303121011"></a>

<a id="canonical-3230123111332222-1031212320200111-2131002122322111-2231220002200100-3210303000301231-2322323002301313-3003102233003130-0133213301111033"></a>

## namespace property — Property reference / 221023103012 / 13

Type: `"string"`. Optional, Computed.

Namespace for the Srv6 Network Slice. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

<a id="canonical-2203103303021120-3021100002132130-3211110001321331-2031032030132232-1010132321023233-2210122122221120-1321212321113133-2031333113023103"></a>

<a id="canonical-0300212111332210-0112031120111203-2232212331201023-1012310300300202-2022123000303111-3302332220032220-3022313331313101-3201222023313231"></a>

## sid_prefixes property — Property reference / 221023103012 / 14

Type: `["list", "string"]`. Required.

SID Locator from the prefix is allocated automatically for each node in each site.

Upstream description:

A SID Locator from the prefix is allocated automatically for each node in each site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-2323201033303013-1003020221320103-3103230112112221-3003302021310302-1103223322032111-2220312113303330-1300313232100232-2123233333232113): complete subsection reference.

<a id="canonical-1020100020101021-0323120312113210-1322013023122231-0023130301033123-1332020001123120-0212320002123131-1001102200312130-1003033020121010"></a>

## All schema paths — Property reference / 221023103012 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--srv6_network_slice--reference--group-001.md#canonical-3223123301302320-2220102122230030-2212031313122121-2232211021103121-1012010311013230-2132123223132013-3331002022123233-2211111331010022) |
| `connect_to_access_networks` | [connect_to_access_networks](resources--srv6_network_slice--reference--group-001.md#canonical-2322133310321032-3212210130310201-2010033232332202-3202010200011023-2122223032103311-1321123101003120-3231032110101330-0320230210231012) |
| `connect_to_enterprise_networks` | [connect_to_enterprise_networks](resources--srv6_network_slice--reference--group-001.md#canonical-0333010203300212-2211112212330302-2321121311220122-2200010131122233-3022200001321012-1233230003022332-1210312021201211-0333112110133101) |
| `connect_to_internet` | [connect_to_internet](resources--srv6_network_slice--reference--group-001.md#canonical-1023020031310102-2113000013012312-3101322122002130-1230302322211312-3021312311001121-3231101201301033-3301331332311210-0320131310022210) |
| `description` | [description](resources--srv6_network_slice--reference--group-001.md#canonical-1303212022101203-3030021121332023-2021002331323030-1131311331002001-1300122301000311-1231010122330022-1331211320023120-2003301012211322) |
| `disable` | [disable](resources--srv6_network_slice--reference--group-001.md#canonical-2330322132023333-0032212332130313-1202102031321031-1021020322222313-0220010112130023-2330130120023012-0230232102002311-1312211103321122) |
| `id` | [id](resources--srv6_network_slice--reference--group-001.md#canonical-3233021201223011-2122022303112100-2030330021020111-0111333203300231-0023112121012033-1032330133232230-2102333320230032-1020120101131012) |
| `labels` | [labels](resources--srv6_network_slice--reference--group-001.md#canonical-0003210122312333-3233203011221222-3200203310233201-0302110223201000-0232013002110111-1322213201123200-0313213133203322-2313100123031020) |
| `name` | [name](resources--srv6_network_slice--reference--group-001.md#canonical-3301211003020201-0231323000031211-0321230113022330-1300331201302333-3332301023103011-1010012221010300-3112121232301232-2320000021223113) |
| `namespace` | [namespace](resources--srv6_network_slice--reference--group-001.md#canonical-3321023001110032-0011232322102033-3321100020012201-3201232001000001-1220031310112200-3020022011020031-0303120231022202-3123231303121011) |
| `sid_prefixes` | [sid_prefixes](resources--srv6_network_slice--reference--group-001.md#canonical-2203103303021120-3021100002132130-3211110001321331-2031032030132232-1010132321023233-2210122122221120-1321212321113133-2031333113023103) |
| `timeouts` | [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-3313201222200111-2113300221011303-2121232030003100-1201221303231121-3003110211023223-2312232101032132-1111120133112212-2100231223030002) |
| `timeouts.create` | [timeouts.create](resources--srv6_network_slice--reference--group-001.md#canonical-0002011312103132-0203220021100021-3302200022200221-2323311132303210-0210030011300222-3033313203231120-3020131230213030-1103200022320313) |
| `timeouts.delete` | [timeouts.delete](resources--srv6_network_slice--reference--group-001.md#canonical-2110320003023201-2322110010313211-1012221311312223-1303021311321113-1131213032333131-0031233132020312-0300123000202010-3233131330221311) |
| `timeouts.read` | [timeouts.read](resources--srv6_network_slice--reference--group-001.md#canonical-1101220202311133-2130320023131212-1113203022123333-3122311011003233-3230122330312302-1121323323103322-1012030020320330-0100232111300102) |
| `timeouts.update` | [timeouts.update](resources--srv6_network_slice--reference--group-001.md#canonical-2023222030013101-0001022022012201-2330033020311111-1211012230023131-2300103212303112-3033130230233221-2132130300123000-2212102213321032) |

<a id="canonical-0300102210122121-2100121220303112-1231232333322230-2023313303103131-2010323332012032-0333221330333000-3021123223001101-3113031120132013"></a>

## Next pages — Property reference / 221023103012 / 16

- [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-2323201033303013-1003020221320103-3103230112112221-3003302021310302-1103223322032111-2220312113303330-1300313232100232-2123233333232113)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)

<a id="canonical-2323201033303013-1003020221320103-3103230112112221-3003302021310302-1103223322032111-2220312113303330-1300313232100232-2123233333232113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312120123301033-1000013111301212-0222130300020120-1021330121133112-2021002222202023-1311012112222112-3233111210313123-2033103012210212"></a>

## timeouts — timeouts / 233332300131 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)
- [Property reference](resources--srv6_network_slice--reference--group-001.md#canonical-2221210201331300-0011233131200112-1301201030112212-2211123330221021-3333331233302122-0320123310310202-3200030231001021-3012033330201300)
- timeouts

<a id="canonical-3313201222200111-2113300221011303-2121232030003100-1201221303231121-3003110211023223-2312232101032132-1111120133112212-2100231223030002"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201332223331223-1230323313032101-0200322301110212-3110122020131000-2300033320132031-3200121323002031-1200213110303320-0121012133212303"></a>

## Direct properties — timeouts / 233332300131 / 3

<a id="canonical-0002011312103132-0203220021100021-3302200022200221-2323311132303210-0210030011300222-3033313203231120-3020131230213030-1103200022320313"></a>

<a id="canonical-0011310213112020-2303313211033311-1210313013001120-3302131013121013-1111230113000300-2200133300011210-0200000103013222-3312220303000033"></a>

## create property — timeouts / 233332300131 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2110320003023201-2322110010313211-1012221311312223-1303021311321113-1131213032333131-0031233132020312-0300123000202010-3233131330221311"></a>

<a id="canonical-1310321213210323-2133213012100313-2132203003030122-1222331311123031-0221100002132010-2103201132103230-0320333212330122-3103032132133300"></a>

## delete property — timeouts / 233332300131 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1101220202311133-2130320023131212-1113203022123333-3122311011003233-3230122330312302-1121323323103322-1012030020320330-0100232111300102"></a>

<a id="canonical-1230031300202221-0313330020320033-2321100111030123-1132311231223222-3130023120323033-2012123000033032-2200011101001202-0333233010323222"></a>

## read property — timeouts / 233332300131 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2023222030013101-0001022022012201-2330033020311111-1211012230023131-2300103212303112-3033130230233221-2132130300123000-2212102213321032"></a>

<a id="canonical-2300131101132221-1323102101302232-1311132302332022-2103011330032300-3211312233103133-1023302330023032-1332121202111213-2231300322322110"></a>

## update property — timeouts / 233332300131 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1322332233030113-1222220131020123-0200010101100201-3221230333312330-1230322201132011-3122202031100122-0020200203221101-1322221232322021"></a>

## Next pages — timeouts / 233332300131 / 8

- [Property reference](resources--srv6_network_slice--reference--group-001.md#canonical-2221210201331300-0011233131200112-1301201030112212-2211123330221021-3333331233302122-0320123310310202-3200030231001021-3012033330201300)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)
