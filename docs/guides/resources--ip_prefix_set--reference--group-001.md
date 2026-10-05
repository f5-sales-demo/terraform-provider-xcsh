---
page_title: "xcsh_ip_prefix_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set reference."
---

# xcsh_ip_prefix_set reference

<a id="canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031122031133101-2223223311203011-1310133333312033-0321101012221213-3102002033212233-0002013112320320-2232330013202023-1103013122230032"></a>

## Property reference — Property reference / 221201303012 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
- Property reference

<a id="canonical-2233131110330110-2033303311133333-0020100203301231-0022011111122222-1321003031233221-1030031231120011-0131221013121013-0032011112331000"></a>

## Direct properties — Property reference / 221201303012 / 3

<a id="canonical-3111011131133310-0100322211101011-3111000112210011-0030302113013233-0200213310333200-1020302101330200-1333310330112022-1121102321130011"></a>

<a id="canonical-3302102200203320-1323033000201202-2311210011303032-0302103310003130-1112131211333313-1322322033132311-0232233230030023-1133321000113000"></a>

## annotations property — Property reference / 221201303012 / 4

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

<a id="canonical-1103032012112102-1222103002311330-1022001322132101-3011313120300011-2110103233012202-0002332122112103-3212130303111032-2122132323210212"></a>

<a id="canonical-1211033103232033-2200012113203200-2333003321113013-2301202013003123-1130211110002333-3223030332203020-0032003312021333-2112310312021223"></a>

## description property — Property reference / 221201303012 / 5

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

<a id="canonical-0220022013232322-3031210221012301-1132010213101330-3012121012230023-1013311011202202-0012232133102311-2231200223233032-0001330213232302"></a>

<a id="canonical-0120012023031132-2201111032103032-3201003033330203-3223103210100231-0320111300133130-1131320130310311-1322011232112022-3200300211230131"></a>

## disable property — Property reference / 221201303012 / 6

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

<a id="canonical-2220202030020313-2201313212211210-3323322031203010-0313031003111000-1001133110202130-3121113310222113-1023330122323220-1313202122112230"></a>

<a id="canonical-3212112213033200-2101201200120130-0331230321021302-0033123331331001-0031013030222321-1031233331233100-0003212310000100-1100222120111001"></a>

## ID property — Property reference / 221201303012 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-2200333320220021-3021300112012033-2230031200012331-2110030133331300-3320332221003021-1020200323133031-0201022120222300-2112013311220230): complete subsection reference.

<a id="canonical-0100010122102111-1023011210002002-2212001000132102-3102132002212200-0133133203013332-3020223203121202-1331031301111200-0211232203103222"></a>

<a id="canonical-0002112330321332-0313100122200123-1312301232110133-0222000001133312-2302100331212312-1200220211202002-1200330303320131-2303011011111210"></a>

## labels property — Property reference / 221201303012 / 8

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

<a id="canonical-1012222103020021-2003330222010003-2031030122020021-0010332302032312-1033030233322110-0202023103112301-1301122030200013-0032222303022123"></a>

<a id="canonical-0121232200303313-2000021221230200-2312103201030102-2130202222223122-2122110110302131-3023230120120030-0220120323030000-3131221130133202"></a>

## name property — Property reference / 221201303012 / 9

Type: `"string"`. Required.

Name of the IP Prefix Set. Must be unique within the namespace.

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

<a id="canonical-1122103330311033-3012030021000213-3310230223323200-2021311202333202-0311320212200330-1331021221033302-2231023013120202-3302133232132122"></a>

<a id="canonical-0210002001022300-3323103131330003-3030201110133223-2022200213222133-3131312213132221-0203331310220033-3030203023301323-0123003110003213"></a>

## namespace property — Property reference / 221201303012 / 10

Type: `"string"`. Required.

Namespace where the IP Prefix Set is created.

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

- [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-2321201120230030-3001232220000330-1113131203001211-2300110003111212-3111020133113030-0201303201202002-0203102013002031-1103113300220213): complete subsection reference.

<a id="canonical-1201133200203301-1002321113021212-2011101210323011-0032011003202110-1220230202303122-3031201101213133-3030321322211003-1100110122102321"></a>

## All schema paths — Property reference / 221201303012 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ip_prefix_set--reference--group-001.md#canonical-3111011131133310-0100322211101011-3111000112210011-0030302113013233-0200213310333200-1020302101330200-1333310330112022-1121102321130011) |
| `description` | [description](resources--ip_prefix_set--reference--group-001.md#canonical-1103032012112102-1222103002311330-1022001322132101-3011313120300011-2110103233012202-0002332122112103-3212130303111032-2122132323210212) |
| `disable` | [disable](resources--ip_prefix_set--reference--group-001.md#canonical-0220022013232322-3031210221012301-1132010213101330-3012121012230023-1013311011202202-0012232133102311-2231200223233032-0001330213232302) |
| `id` | [ID](resources--ip_prefix_set--reference--group-001.md#canonical-2220202030020313-2201313212211210-3323322031203010-0313031003111000-1001133110202130-3121113310222113-1023330122323220-1313202122112230) |
| `ipv4_prefixes` | [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-0132122021001312-0033022200220203-2130231033302112-0300013033301112-2103232312033331-1031320310202330-1121023202220033-1303102011100310) |
| `ipv4_prefixes.description_spec` | [ipv4_prefixes.description_spec](resources--ip_prefix_set--reference--group-001.md#canonical-3112120332021330-3123231002003303-0331313130001211-1111022021030112-2301303020212321-1012133130210311-2133022101312001-2211220301103131) |
| `ipv4_prefixes.ipv4_prefix` | [ipv4_prefixes.ipv4_prefix](resources--ip_prefix_set--reference--group-001.md#canonical-1201113001132101-0230301320311331-2003203100032013-2321123200311010-2111000301012122-3033033330201030-1333230103330100-3211102023333212) |
| `labels` | [labels](resources--ip_prefix_set--reference--group-001.md#canonical-0100010122102111-1023011210002002-2212001000132102-3102132002212200-0133133203013332-3020223203121202-1331031301111200-0211232203103222) |
| `name` | [name](resources--ip_prefix_set--reference--group-001.md#canonical-1012222103020021-2003330222010003-2031030122020021-0010332302032312-1033030233322110-0202023103112301-1301122030200013-0032222303022123) |
| `namespace` | [namespace](resources--ip_prefix_set--reference--group-001.md#canonical-1122103330311033-3012030021000213-3310230223323200-2021311202333202-0311320212200330-1331021221033302-2231023013120202-3302133232132122) |
| `timeouts` | [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-0303300110230132-3103032210220230-3212000130013102-2033320131311112-0010101313313231-0102013132300101-2313002202021103-0322103032330020) |
| `timeouts.create` | [timeouts.create](resources--ip_prefix_set--reference--group-001.md#canonical-0011111132121222-3032112231020113-1031313120101021-0232212020321031-0002013132320011-1031033013131331-1111033302202111-2233313111123010) |
| `timeouts.delete` | [timeouts.delete](resources--ip_prefix_set--reference--group-001.md#canonical-0110231320102101-2022012112200332-1230032131112313-1221203330121132-1102010212123130-3221230101020203-2012212212320200-1233222100022300) |
| `timeouts.read` | [timeouts.read](resources--ip_prefix_set--reference--group-001.md#canonical-0331222110111123-2120103212000201-3123000201110030-1212120230122210-2132020101111101-1330122220011012-0213110122310100-2111303002120011) |
| `timeouts.update` | [timeouts.update](resources--ip_prefix_set--reference--group-001.md#canonical-1230120122210222-1313133021310102-3112000221031320-0101123300223210-1330200110222200-1031223332213030-1330321002330201-0312022002310103) |

<a id="canonical-0332110032121132-3133013232332122-2302102030002312-2103011010033122-0101302132012222-0200222000020223-2020221101322331-0223113003210000"></a>

## Next pages — Property reference / 221201303012 / 12

- [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-2200333320220021-3021300112012033-2230031200012331-2110030133331300-3320332221003021-1020200323133031-0201022120222300-2112013311220230)
- [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-2321201120230030-3001232220000330-1113131203001211-2300110003111212-3111020133113030-0201303201202002-0203102013002031-1103113300220213)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)

<a id="canonical-2200333320220021-3021300112012033-2230031200012331-2110030133331300-3320332221003021-1020200323133031-0201022120222300-2112013311220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020101312210110-3223112102313032-3313210021003220-2010112100320030-0133302301223000-0212300031033011-3133032102330230-2231302130312200"></a>

## ipv4_prefixes — ipv4_prefixes / 001003010220 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221)
- ipv4_prefixes

<a id="canonical-0132122021001312-0033022200220203-2130231033302112-0300013033301112-2103232312033331-1031320310202330-1121023202220033-1303102011100310"></a>

Type: `"object"`. list nested block, Optional.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ipv4_prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ipv4_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311113132333221-0021002212211032-1103003213313301-0133320022022203-2320320321311302-2233111210232131-2001100032021333-3210120233023000"></a>

## Direct properties — ipv4_prefixes / 001003010220 / 3

<a id="canonical-3112120332021330-3123231002003303-0331313130001211-1111022021030112-2301303020212321-1012133130210311-2133022101312001-2211220301103131"></a>

<a id="canonical-0121333212132230-0132113103323103-2310102033010031-0200331223102230-3003321123222331-0222111121022033-0320010220130031-3301332300300131"></a>

## description_spec property — ipv4_prefixes / 001003010220 / 4

Type: `"string"`. Optional.

Description. Human-readable description text

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-1201113001132101-0230301320311331-2003203100032013-2321123200311010-2111000301012122-3033033330201030-1333230103330100-3211102023333212"></a>

<a id="canonical-1313320321003230-1003032032112210-3010321132301303-2300220211211133-2300021222023313-2100212221030210-1032121303232100-2021231300203010"></a>

## ipv4_prefix property — ipv4_prefixes / 001003010220 / 5

Type: `"string"`. Optional.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2032021102003311-0200211013022300-1201132100331233-2203000323230202-3232200102323121-2113013300122313-1301310010331213-2132011130203223"></a>

## Next pages — ipv4_prefixes / 001003010220 / 6

- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)

<a id="canonical-2321201120230030-3001232220000330-1113131203001211-2300110003111212-3111020133113030-0201303201202002-0203102013002031-1103113300220213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000023031311100-0210211103110133-2201310133213332-2133300230100221-0301200330321010-3100200331012110-2232133132013231-3020300323001311"></a>

## timeouts — timeouts / 321123002101 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221)
- timeouts

<a id="canonical-0303300110230132-3103032210220230-3212000130013102-2033320131311112-0010101313313231-0102013132300101-2313002202021103-0322103032330020"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103233312301310-0111202213010330-1312003130233023-1110100110333322-3033231233031301-1020230332320302-3103023232311021-0210120102123222"></a>

## Direct properties — timeouts / 321123002101 / 3

<a id="canonical-0011111132121222-3032112231020113-1031313120101021-0232212020321031-0002013132320011-1031033013131331-1111033302202111-2233313111123010"></a>

<a id="canonical-3103330200002312-1003222332103232-1113112310223231-2233313111222211-2333210123100012-0010000223333313-3312231201202321-3221333012031222"></a>

## create property — timeouts / 321123002101 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0110231320102101-2022012112200332-1230032131112313-1221203330121132-1102010212123130-3221230101020203-2012212212320200-1233222100022300"></a>

<a id="canonical-0101313023031110-0210313113221321-3012323020323333-1233001001131312-0023201221231121-1232232020321212-1121023100303103-0212123300131223"></a>

## delete property — timeouts / 321123002101 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0331222110111123-2120103212000201-3123000201110030-1212120230122210-2132020101111101-1330122220011012-0213110122310100-2111303002120011"></a>

<a id="canonical-1233033030112012-3221212000230032-3302012301033302-0121220233000301-2202210011012102-2111322030300310-2311022013132132-2220311002223100"></a>

## read property — timeouts / 321123002101 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1230120122210222-1313133021310102-3112000221031320-0101123300223210-1330200110222200-1031223332213030-1330321002330201-0312022002310103"></a>

<a id="canonical-3012130011332123-3212322131133032-0013111010331111-0111302123232321-1100333223111022-0230022223032013-3100310113100120-0122003100011000"></a>

## update property — timeouts / 321123002101 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3312003003113101-0001101331330210-0333102320101231-0032011132110030-0213320202012300-0202100212333200-2301222131332310-1302012131230231"></a>

## Next pages — timeouts / 321123002101 / 8

- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
