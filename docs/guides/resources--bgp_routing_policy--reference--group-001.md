---
page_title: "xcsh_bgp_routing_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy reference."
---

# xcsh_bgp_routing_policy reference

<a id="canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212022333021021-2232330201022233-1211233122031113-0020213030120202-2222021203001322-2321022303323120-0222133101333103-3332102201302330"></a>

## Property reference — Property reference / 221033101012 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- Property reference

<a id="canonical-2232012122103022-2312300332103103-1123301313011033-3100321331112300-0022330302020021-2223022132032020-1310201013322033-3300332223311212"></a>

## Direct properties — Property reference / 221033101012 / 3

<a id="canonical-1310111230303331-1202233023012321-2322000230000103-2131003002330100-2132020131202322-3110223213212111-0310332021030313-0101222020012010"></a>

<a id="canonical-0303131201223203-1111202131113210-3011131013100132-0302320220031031-3023311311031231-1113132311020002-2012331123223102-3100233223121331"></a>

## annotations property — Property reference / 221033101012 / 4

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

<a id="canonical-3003031213323222-3131201103101322-3221003100012231-1102233022121013-2233132000123212-2331221100321333-2232230303320230-2202232311013101"></a>

<a id="canonical-1122100230013310-2003332333301131-3222332233010202-1013213010132033-2133130001000130-0003123101123230-0001323232323333-2011310002320123"></a>

## description property — Property reference / 221033101012 / 5

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

<a id="canonical-2231222023020100-2231122011311213-2313121032101333-3321200113001021-0301222023212213-2301221103222213-1202232132210313-0230213203021113"></a>

<a id="canonical-1000032012030113-2232102000312122-2200312212303212-3010333023311310-0002302132221110-0300010311002203-1122210022020301-2122203202213311"></a>

## disable property — Property reference / 221033101012 / 6

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

<a id="canonical-1213030020032323-0300022002300122-2111031220233000-1301020003021330-1120003232111013-0311320103031313-1323233013120333-0311320112120331"></a>

<a id="canonical-0020203023010220-1220222031003132-2122300213021131-3201312200312201-0223133110010322-3232230212130223-3203020030032302-2330311120303222"></a>

## ID property — Property reference / 221033101012 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1020201021012223-1322303320001303-2103131331203022-3022030023211322-0330113131111122-1311132311223103-1102113030110112-0101131303011013"></a>

<a id="canonical-1030011031232221-0011032031213313-3131110312302310-2232103102321102-0103122021320021-1110311320213113-1100113121302032-0011032111301210"></a>

## labels property — Property reference / 221033101012 / 8

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

<a id="canonical-1100110332322330-3303210301303122-3230331323033331-3301223003021322-0012003231023333-2203320233023301-2013132011011101-0231200112032313"></a>

<a id="canonical-1230022103132303-3032120112010013-2112202312003200-3000223202322221-2110312331120112-1022011221100301-0032032133202332-3011331003231031"></a>

## name property — Property reference / 221033101012 / 9

Type: `"string"`. Required.

Name of the BGP Routing Policy. Must be unique within the namespace.

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

<a id="canonical-2130230111201201-2122120200100312-1013120110011133-2321220301122232-2132010303033121-0113200002320010-0131021330333010-2203032001221011"></a>

<a id="canonical-2230302132213020-2000100032232123-1123023133320231-3111322120130322-2203231113110213-3100131002002023-0331320221232110-0103113332113312"></a>

## namespace property — Property reference / 221033101012 / 10

Type: `"string"`. Required.

Namespace where the BGP Routing Policy is created.

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

- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321): complete subsection reference.

- [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-3021222012003333-1321333233313303-1021221211303323-1033132203013020-3320122303302230-2130300113321310-0211311213111000-0310230310021023): complete subsection reference.

<a id="canonical-1203003231110320-2313032013301212-1333012220131313-2312112120300301-1110031300123300-3112123201020002-0303122203311303-3331122012210210"></a>

## All schema paths — Property reference / 221033101012 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp_routing_policy--reference--group-001.md#canonical-1310111230303331-1202233023012321-2322000230000103-2131003002330100-2132020131202322-3110223213212111-0310332021030313-0101222020012010) |
| `description` | [description](resources--bgp_routing_policy--reference--group-001.md#canonical-3003031213323222-3131201103101322-3221003100012231-1102233022121013-2233132000123212-2331221100321333-2232230303320230-2202232311013101) |
| `disable` | [disable](resources--bgp_routing_policy--reference--group-001.md#canonical-2231222023020100-2231122011311213-2313121032101333-3321200113001021-0301222023212213-2301221103222213-1202232132210313-0230213203021113) |
| `id` | [id](resources--bgp_routing_policy--reference--group-001.md#canonical-1213030020032323-0300022002300122-2111031220233000-1301020003021330-1120003232111013-0311320103031313-1323233013120333-0311320112120331) |
| `labels` | [labels](resources--bgp_routing_policy--reference--group-001.md#canonical-1020201021012223-1322303320001303-2103131331203022-3022030023211322-0330113131111122-1311132311223103-1102113030110112-0101131303011013) |
| `name` | [name](resources--bgp_routing_policy--reference--group-001.md#canonical-1100110332322330-3303210301303122-3230331323033331-3301223003021322-0012003231023333-2203320233023301-2013132011011101-0231200112032313) |
| `namespace` | [namespace](resources--bgp_routing_policy--reference--group-001.md#canonical-2130230111201201-2122120200100312-1013120110011133-2321220301122232-2132010303033121-0113200002320010-0131021330333010-2203032001221011) |
| `rules` | [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3221201233303220-3002032231113230-3131332001132313-3233210131223220-0302021122133132-0001022313221220-3222323232311210-1333333210101021) |
| `rules.action` | [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2320233101301321-0023000100032001-2323130002220132-2002310320222321-1333202211012331-0131211031203003-1220332130002110-3233312001030110) |
| `rules.action.allow` | [rules.action.allow](resources--bgp_routing_policy--reference--group-001.md#canonical-3302212133202011-0100232023023120-1113101110001213-0232111300200321-2021310001323300-2001133303333121-2102123312233303-2322313313211210) |
| `rules.action.as_path` | [rules.action.as_path](resources--bgp_routing_policy--reference--group-001.md#canonical-1312312010200311-3310310123310330-3021330030300123-0302302101023030-0031310211001321-0201332031203320-1310131113212203-0332313123130021) |
| `rules.action.community` | [rules.action.community](resources--bgp_routing_policy--reference--group-001.md#canonical-1022201123230222-0230310112001300-2313331123133212-0001223221220310-3023331300020013-1130121002220100-2132331201003302-3013310102203121) |
| `rules.action.community.community` | [rules.action.community.community](resources--bgp_routing_policy--reference--group-001.md#canonical-2032221120230120-3331332113300110-0231220110310211-2211103023212232-1031023302200113-1222301323321123-3303313032310101-2132023003221101) |
| `rules.action.deny` | [rules.action.deny](resources--bgp_routing_policy--reference--group-001.md#canonical-0112330323120002-0010301110002331-0310032321031121-1313033212232231-0330110223230011-3230002212100101-1132200130010333-1330131301010023) |
| `rules.action.local_preference` | [rules.action.local_preference](resources--bgp_routing_policy--reference--group-001.md#canonical-2301111020212001-1203202021201003-0302313000333130-0231111301202122-1303233022213333-2212132323300232-1311013032103321-0123003331112321) |
| `rules.action.metric` | [rules.action.metric](resources--bgp_routing_policy--reference--group-001.md#canonical-3323001033122030-1101011030031120-0231200113332300-3020133301203303-0231200103100013-1130021030320221-1132012120003032-1201101030310310) |
| `rules.match` | [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-0133233211011030-1032311120211210-0020130133013213-1012010300213033-3301002033300212-0011310202221131-3300133223212313-1001032211113012) |
| `rules.match.as_path` | [rules.match.as_path](resources--bgp_routing_policy--reference--group-001.md#canonical-3330231023302312-1123011113121030-1310031103030112-2200320000021011-1301131232130011-1333321211221121-0233321333222231-1112113311311330) |
| `rules.match.community` | [rules.match.community](resources--bgp_routing_policy--reference--group-001.md#canonical-1110202201122121-0102230022023130-3313331012223220-1313100300133202-2201010000202100-1320333012110000-3120133201320021-2031010010323310) |
| `rules.match.community.community` | [rules.match.community.community](resources--bgp_routing_policy--reference--group-001.md#canonical-2110133332333010-3030303103231311-2020032320113111-2111202113021331-1213222330122011-0003302310101030-2023311212120312-3123030331212313) |
| `rules.match.ip_prefixes` | [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-3313111002301113-1122101323200112-1200201003302020-3220313210120213-3103123233020102-1210031003100233-2001022213230322-0103123001330033) |
| `rules.match.ip_prefixes.prefixes` | [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1121103202332031-0211001331003133-0122330131232123-2012222132020202-1121002311020002-0232031123023021-3333322010330211-2001031321310010) |
| `rules.match.ip_prefixes.prefixes.equal_or_longer_than` | [rules.match.ip_prefixes.prefixes.equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-3003031032010333-2220202321030102-3222120322132130-0303130031130231-0121131223312203-2233030231321232-2233323232130301-1333332000110310) |
| `rules.match.ip_prefixes.prefixes.exact_match` | [rules.match.ip_prefixes.prefixes.exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-1132223111023103-0331313320333231-3131232220311302-0222310110210033-1313122321033000-0010212012110301-3311020100300213-2331211201333121) |
| `rules.match.ip_prefixes.prefixes.ip_prefixes` | [rules.match.ip_prefixes.prefixes.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0110001303301332-3132123213102232-2232332203330003-2012211112133030-2323011311122333-0231320213102000-0031200000000010-3300031030211121) |
| `rules.match.ip_prefixes.prefixes.longer_than` | [rules.match.ip_prefixes.prefixes.longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-1313003210011312-2101020122233003-2233123101323003-2111112332031323-0202202022231211-2030323101031223-2012110002111231-3230100022003132) |
| `timeouts` | [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-2012203122201330-0202232122021313-3203011322033312-2232211310120302-2333123321113211-2000202021203130-2330332022223331-0321310201303301) |
| `timeouts.create` | [timeouts.create](resources--bgp_routing_policy--reference--group-001.md#canonical-3023310131300223-3213333103321320-1232102303021221-2312333000121130-0132113032223331-1013322310320002-1111303210131023-2203221310331221) |
| `timeouts.delete` | [timeouts.delete](resources--bgp_routing_policy--reference--group-001.md#canonical-3011220002201003-3033032322130102-2033213212112011-2231301023212330-1211211220031001-1213210013102100-2020210121010302-3020020321211103) |
| `timeouts.read` | [timeouts.read](resources--bgp_routing_policy--reference--group-001.md#canonical-0321130012201023-1032310002320330-2121001202002021-1312103301222322-3003020303021133-0332233223331030-0333111131321103-0212120033223200) |
| `timeouts.update` | [timeouts.update](resources--bgp_routing_policy--reference--group-001.md#canonical-0232121133112313-0302222333332221-1111022013123002-0231023012021111-1211103310121202-0003032303120300-2211010210320210-2231323033302122) |

<a id="canonical-3033212333211332-0003112230320203-0220322223021012-1313323323200203-2032033130331303-0220013133320133-3132233210220232-2222221000211321"></a>

## Next pages — Property reference / 221033101012 / 12

- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-3021222012003333-1321333233313303-1021221211303323-1033132203013020-3320122303302230-2130300113321310-0211311213111000-0310230310021023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012010310312313-2031031112301013-0131321003121200-0112300133121332-3130323223010232-0222033320322121-1210302320002022-0202021110023000"></a>

## rules — rules / 210013302102 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- rules

<a id="canonical-3221201233303220-3002032231113230-3131332001132313-3233210131223220-0302021122133132-0001022313221220-3222323232311210-1333333210101021"></a>

Type: `"object"`. list nested block, Optional.

BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Upstream description:

A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002101100030011-1220031121000312-3200313000333100-2223032133033333-2230201033032022-2101010100300023-1101021333003023-1003203132233131"></a>

## Direct properties — rules / 210013302102 / 3

- [action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023): complete subsection reference.

- [match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111): complete subsection reference.

<a id="canonical-3322233133303231-3030320013122331-3101132132103102-0301001233313200-0321130031110022-0120131121321213-0303220002230230-0323020003203231"></a>

## Next pages — rules / 210013302102 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100012210313220-1100313113220110-1201200100202233-3121310121013013-1033113231211322-0200312022303003-3332220312133203-3230331223321112"></a>

## rules.action — action / 312132333030 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- rules.action

<a id="canonical-2320233101301321-0023000100032001-2323130002220132-2002310320222321-1333202211012331-0131211031203003-1220332130002110-3233312001030110"></a>

Type: `"object"`. single nested block, Optional.

Action to be enforced if the BGP route matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "as_path"),
  validators.ConflictingObjectAttributes("allow",
    "community"),
  validators.ConflictingObjectAttributes("allow",
    "deny"),
  validators.ConflictingObjectAttributes("allow",
    "local_preference"),
  validators.ConflictingObjectAttributes("allow",
    "metric"),
  validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "deny"),
  validators.ConflictingObjectAttributes("as_path",
    "local_preference"),
  validators.ConflictingObjectAttributes("as_path",
    "metric"),
  validators.ConflictingObjectAttributes("community",
    "deny"),
  validators.ConflictingObjectAttributes("community",
    "local_preference"),
  validators.ConflictingObjectAttributes("community",
    "metric"),
  validators.ConflictingObjectAttributes("deny",
    "local_preference"),
  validators.ConflictingObjectAttributes("deny",
    "metric"),
  validators.ConflictingObjectAttributes("local_preference",
    "metric")}
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
  "x-ves-oneof-field-action_type": "[\"allow\",\"as_path\",\"community\",\"deny\",\"local_preference\",\"metric\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020001001100031-2311233100233312-1102233332112203-1030122120321300-0102000222003103-2311313023030013-0110110121110121-3100331202022332"></a>

## Direct properties — action / 312132333030 / 3

- [allow](resources--bgp_routing_policy--reference--group-001.md#canonical-0021130232302232-0313010111011031-2212010021233102-0302133233321200-1212213303100312-3023331320222022-0220232323330300-3201112022203213): complete subsection reference.

<a id="canonical-1312312010200311-3310310123310330-3021330030300123-0302302101023030-0031310211001321-0201332031203320-1310131113212203-0332313123130021"></a>

<a id="canonical-2330033023233003-2130003303130022-3020003303233132-1121102031321022-2313213033031223-0211321130322111-0333233230222211-0130112231020213"></a>

## as_path property — action / 312132333030 / 4

Type: `"string"`. Optional.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Upstream description:

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [community](resources--bgp_routing_policy--reference--group-001.md#canonical-0333110122323122-3023303013023002-2202303131020310-0232322022301020-3300221203311200-1012200322021331-0220222201223103-1323110201313130): complete subsection reference.

- [deny](resources--bgp_routing_policy--reference--group-001.md#canonical-3103303312013323-2203323103311102-2100033321112303-2010031131010120-2103003332223013-1033333003131223-3303310010103103-1100311202101011): complete subsection reference.

<a id="canonical-2301111020212001-1203202021201003-0302313000333130-0231111301202122-1303233022213333-2212132323300232-1311013032103321-0123003331112321"></a>

<a id="canonical-2113120131011033-3130221112013331-0111333302001300-1120212102323023-3021020133032032-2032002202232311-2023111232232320-1312033013320213"></a>

## local_preference property — action / 312132333030 / 5

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Upstream description:

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

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

<a id="canonical-3323001033122030-1101011030031120-0231200113332300-3020133301203303-0231200103100013-1130021030320221-1132012120003032-1201101030310310"></a>

<a id="canonical-3130010031013011-1232202102332332-2311003123312323-2332322132321110-0200033210301123-3221321230310202-1212311113031002-2331100211313011"></a>

## metric property — action / 312132333030 / 6

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Upstream description:

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

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

<a id="canonical-0031002202000311-1130003123332001-1033130200303202-2323332232001320-0313102333123302-3200132301310322-2133032102003133-2000332112102333"></a>

## Next pages — action / 312132333030 / 7

- [rules.action.allow](resources--bgp_routing_policy--reference--group-001.md#canonical-0021130232302232-0313010111011031-2212010021233102-0302133233321200-1212213303100312-3023331320222022-0220232323330300-3201112022203213)
- [rules.action.community](resources--bgp_routing_policy--reference--group-001.md#canonical-0333110122323122-3023303013023002-2202303131020310-0232322022301020-3300221203311200-1012200322021331-0220222201223103-1323110201313130)
- [rules.action.deny](resources--bgp_routing_policy--reference--group-001.md#canonical-3103303312013323-2203323103311102-2100033321112303-2010031131010120-2103003332223013-1033333003131223-3303310010103103-1100311202101011)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-0021130232302232-0313010111011031-2212010021233102-0302133233321200-1212213303100312-3023331320222022-0220232323330300-3201112022203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130311113023322-0202201222313130-1313203321201233-1013310331310203-0113230133101001-3021021320000122-2333310200132332-1211230110133331"></a>

## rules.action.allow — allow / 021010000313 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- rules.action.allow

<a id="canonical-3302212133202011-0100232023023120-1113101110001213-0232111300200321-2021310001323300-2001133303333121-2102123312233303-2322313313211210"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
allow = {}
```

<a id="canonical-0110031330212220-1103021330013323-2121110022111201-1022120201102021-3122301123212201-3210320110131113-0330013000120122-3013201100332310"></a>

## Direct properties — allow / 021010000313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033020031211130-1310132020302032-3033301231131303-2303111031013133-2302131023120123-3023223130103212-3111030300121301-2321113022202020"></a>

## Next pages — allow / 021010000313 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-0333110122323122-3023303013023002-2202303131020310-0232322022301020-3300221203311200-1012200322021331-0220222201223103-1323110201313130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232231102101130-0200233211220102-0321303101311323-0102231001011003-1021221312033200-0131011321313232-3200330321223233-3130202231033001"></a>

## rules.action.community — community / 331131320031 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- rules.action.community

<a id="canonical-1022201123230222-0230310112001300-2313331123133212-0001223221220310-3023331300020013-1130121002220100-2132331201003302-3013310102203121"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
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
community {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231021322200302-1121310132013033-2030213221113332-3322111320200220-1033233101110232-1112203212111222-2312113013210033-3132131201203323"></a>

## Direct properties — community / 331131320031 / 3

<a id="canonical-2032221120230120-3331332113300110-0231220110310211-2211103023212232-1031023302200113-1222301323321123-3303313032310101-2132023003221101"></a>

<a id="canonical-1031230210332302-2012032212010212-2310120100101203-1211221213212100-3223001320033333-0200121023321200-0022333212333212-0030030021321301"></a>

## community property — community / 331131320031 / 4

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1001231103112210-2123311111301011-1201303012122330-1132311322020112-2030012020121233-1022313003220303-1131013313210211-0210202120122002"></a>

## Next pages — community / 331131320031 / 5

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3103303312013323-2203323103311102-2100033321112303-2010031131010120-2103003332223013-1033333003131223-3303310010103103-1100311202101011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132223032202021-3023201201330121-3323120120221000-1111110121322130-0232313301103310-3301323302000230-1231211001311332-3110133232120322"></a>

## rules.action.deny — deny / 212211021112 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- rules.action.deny

<a id="canonical-0112330323120002-0010301110002331-0310032321031121-1313033212232231-0330110223230011-3230002212100101-1132200130010333-1330131301010023"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
deny = {}
```

<a id="canonical-1331103020232313-3203310133103230-1031011313230213-0022233130003133-0223002201230321-0003322321021312-0300121220302203-0333103320021110"></a>

## Direct properties — deny / 212211021112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131332020023133-3331311231013130-3300211200201210-0131200310223330-0003302222332100-0202113333033312-2203330023012112-0020332222123102"></a>

## Next pages — deny / 212211021112 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312023232212231-2031102103333013-1110333312132111-0301103003303101-0333333001022232-1000322012210321-0233333321213330-0310102132000231"></a>

## rules.match — match / 130213232231 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- rules.match

<a id="canonical-0133233211011030-1032311120211210-0020130133013213-1012010300213033-3301002033300212-0011310202221131-3300133223212313-1001032211113012"></a>

Type: `"object"`. single nested block, Optional.

Predicates which have to match information in route for action to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "ip_prefixes"),
  validators.ConflictingObjectAttributes("community",
    "ip_prefixes")}
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
  "x-ves-oneof-field-type_of_match": "[\"as_path\",\"community\",\"ip_prefixes\"]"
}
```

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212110102031000-1220301031110223-2320101111311330-2030321220321101-1202022133230330-3231302200221300-1231110022321113-0232221103312322"></a>

## Direct properties — match / 130213232231 / 3

<a id="canonical-3330231023302312-1123011113121030-1310031103030112-2200320000021011-1301131232130011-1333321211221121-0233321333222231-1112113311311330"></a>

<a id="canonical-2023312120123301-3022301111123211-2130330003200021-0120103030212310-1013101220033221-3012321000323003-2030310002133233-3331220221322303"></a>

## as_path property — match / 130213232231 / 4

Type: `"string"`. Optional.

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Upstream description:

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [community](resources--bgp_routing_policy--reference--group-001.md#canonical-1233011320220011-0332311213110301-3123132110302233-0200233200131200-1320320000122200-3203032322131103-1332203002302102-2022320002030130): complete subsection reference.

- [ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020): complete subsection reference.

<a id="canonical-3131202210203202-2023222220331110-3332130132231333-3011301012230230-2211211321203211-3231231200103320-2130301133302323-1101221303320100"></a>

## Next pages — match / 130213232231 / 5

- [rules.match.community](resources--bgp_routing_policy--reference--group-001.md#canonical-1233011320220011-0332311213110301-3123132110302233-0200233200131200-1320320000122200-3203032322131103-1332203002302102-2022320002030130)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-1233011320220011-0332311213110301-3123132110302233-0200233200131200-1320320000122200-3203032322131103-1332203002302102-2022320002030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010302022310132-3020012221302021-1122130303320222-3323112002121311-2223011310102311-1212022122300323-3233120202021030-0231110312233310"></a>

## rules.match.community — community / 210023133311 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- rules.match.community

<a id="canonical-1110202201122121-0102230022023130-3313331012223220-1313100300133202-2201010000202100-1320333012110000-3120133201320021-2031010010323310"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
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
community {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032330332121233-0033323331030002-3310123002301221-1121230133100102-0120300210232332-1302201030322012-1113221311333002-1313132122211213"></a>

## Direct properties — community / 210023133311 / 3

<a id="canonical-2110133332333010-3030303103231311-2020032320113111-2111202113021331-1213222330122011-0003302310101030-2023311212120312-3123030331212313"></a>

<a id="canonical-3233020011012300-1030310230213202-3023313130211303-2310302331211113-2322120023003230-2320122131101031-0211220131310231-0000133223121011"></a>

## community property — community / 210023133311 / 4

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3212311232033020-1312231302121302-3102303021312302-1323303203001032-1203232310003120-2213211021210210-1210121120300310-1022012323001120"></a>

## Next pages — community / 210023133311 / 5

- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020000310231002-3210011000311211-2133300222300303-2010020123211221-2200333313223212-0031231302020021-3202113213223110-2110222102222232"></a>

## rules.match.ip_prefixes — ip_prefixes / 212330120001 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- rules.match.ip_prefixes

<a id="canonical-3313111002301113-1122101323200112-1200201003302020-3220313210120213-3103123233020102-1210031003100233-2001022213230322-0103123001330033"></a>

Type: `"object"`. single nested block, Optional.

List of IP prefix and prefix length range match condition.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefixes")}
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
ip_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133202023321003-3212330020113233-1023311212233103-3203122212302231-1121310023122113-2212212311312031-2223331330120131-3331310223103311"></a>

## Direct properties — ip_prefixes / 212330120001 / 3

- [prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023): complete subsection reference.

<a id="canonical-3200201002302323-3201323102002022-3313210320111102-0311230100322213-0303321022232331-3200303133130020-1100221301321010-0110220020200003"></a>

## Next pages — ip_prefixes / 212330120001 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120033033022333-3013330213010201-0132122330302332-3231121000021323-0011013310322211-2021122211333302-2010023103122002-2121210100200030"></a>

## rules.match.ip_prefixes.prefixes — prefixes / 301331220311 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- rules.match.ip_prefixes.prefixes

<a id="canonical-1121103202332031-0211001331003133-0122330131232123-2012222132020202-1121002311020002-0232031123023021-3333322010330211-2001031321310010"></a>

Type: `"object"`. list nested block, Optional.

Prefix list. List of IP prefix.

Upstream description:

List of IP prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "exact_match"),
  validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "longer_than"),
  validators.ConflictingListObjectAttributes("exact_match",
    "longer_than")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122233011031012-0033211031023030-1122230212211230-0233111202321001-2013103233131101-3030033310232132-2001111110222221-3002013233103233"></a>

## Direct properties — prefixes / 301331220311 / 3

- [equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-2230123001022111-2113010223032100-2100121322221021-0132123322121303-0330013012310332-0202131023321012-2220031212310130-3321021123002301): complete subsection reference.

- [exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-3201223222000301-0233200201202220-1031233013112031-3102210032203131-3303221311011320-3220133002310220-0032223032131020-2312303230012010): complete subsection reference.

<a id="canonical-0110001303301332-3132123213102232-2232332203330003-2012211112133030-2323011311122333-0231320213102000-0031200000000010-3300031030211121"></a>

<a id="canonical-3023111301311121-3021130313102023-0003132223332132-0232322202310303-3310010100002233-1003303321312020-3101123313121012-2313332013203231"></a>

## ip_prefixes property — prefixes / 301331220311 / 4

Type: `"string"`. Optional.

IP Prefix. IP prefix to match on BGP route.

Upstream description:

IP prefix to match on BGP route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-3012300123222110-1211201132023330-3133033023001110-1320330202210311-1202131330102203-0303231311330123-0323130113211230-3230100233033010): complete subsection reference.

<a id="canonical-3302013301111312-2120101110300310-2013202211032033-2011332032121013-2211300010223010-2301102131222221-2200210110120003-2002211210013103"></a>

## Next pages — prefixes / 301331220311 / 5

- [rules.match.ip_prefixes.prefixes.equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-2230123001022111-2113010223032100-2100121322221021-0132123322121303-0330013012310332-0202131023321012-2220031212310130-3321021123002301)
- [rules.match.ip_prefixes.prefixes.exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-3201223222000301-0233200201202220-1031233013112031-3102210032203131-3303221311011320-3220133002310220-0032223032131020-2312303230012010)
- [rules.match.ip_prefixes.prefixes.longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-3012300123222110-1211201132023330-3133033023001110-1320330202210311-1202131330102203-0303231311330123-0323130113211230-3230100233033010)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-2230123001022111-2113010223032100-2100121322221021-0132123322121303-0330013012310332-0202131023321012-2220031212310130-3321021123002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311232121011202-2331332123331333-3302310222130110-2201033211030002-2330200102130100-0213332012020113-2210231211033123-1101113033313301"></a>

## rules.match.ip_prefixes.prefixes.equal_or_longer_than — equal_or_longer_than / 100002023201 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="canonical-3003031032010333-2220202321030102-3222120322132130-0303130031130231-0121131223312203-2233030231321232-2233323232130301-1333332000110310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for equal or longer than.

Upstream description:

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
equal_or_longer_than = {}
```

<a id="canonical-3302130220310322-1103010233211112-3210302221023320-2033322020220100-1030311101310032-3101113323210133-2012001301101321-1032002303311123"></a>

## Direct properties — equal_or_longer_than / 100002023201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013021201310330-2103132231130120-3311213322133222-0221110123321222-1220302012010322-2311213312312000-0132322212112101-3001122100111322"></a>

## Next pages — equal_or_longer_than / 100002023201 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3201223222000301-0233200201202220-1031233013112031-3102210032203131-3303221311011320-3220133002310220-0032223032131020-2312303230012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112102322200020-2102233120230322-0020301320131203-3320200031101332-3300120211003233-3300212300322332-3300121301331032-0123331300331313"></a>

## rules.match.ip_prefixes.prefixes.exact_match — exact_match / 030110110010 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- rules.match.ip_prefixes.prefixes.exact_match

<a id="canonical-1132223111023103-0331313320333231-3131232220311302-0222310110210033-1313122321033000-0010212012110301-3311020100300213-2331211201333121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exact match.

Upstream description:

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
exact_match = {}
```

<a id="canonical-0211223020001123-3211031322202103-1000313023323002-1311210121233121-0032011111112011-3023323300013312-1312003132203101-2210320102111331"></a>

## Direct properties — exact_match / 030110110010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003132220130120-1331123310231303-3321313222120210-1210332003003211-2121203323200102-2131002030203033-2001012330231201-2302103223000133"></a>

## Next pages — exact_match / 030110110010 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3012300123222110-1211201132023330-3133033023001110-1320330202210311-1202131330102203-0303231311330123-0323130113211230-3230100233033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001310301132120-0023333210113122-1201011331121210-1012002320121333-1332123001212001-0233201110023010-0220330323232200-2210232122021130"></a>

## rules.match.ip_prefixes.prefixes.longer_than — longer_than / 020123010312 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- rules.match.ip_prefixes.prefixes.longer_than

<a id="canonical-1313003210011312-2101020122233003-2233123101323003-2111112332031323-0202202022231211-2030323101031223-2012110002111231-3230100022003132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for longer than.

Upstream description:

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
longer_than = {}
```

<a id="canonical-3012002330112020-1321230330310011-0230202220201133-2130233001011010-0011131233010112-0100013202221202-1020301003202122-3220121333212331"></a>

## Direct properties — longer_than / 020123010312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200020130023223-1323010002300312-2231013032123201-0310333032203003-2100012222021223-0213122302032100-1132023131300102-0310211301121003"></a>

## Next pages — longer_than / 020123010312 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)

<a id="canonical-3021222012003333-1321333233313303-1021221211303323-1033132203013020-3320122303302230-2130300113321310-0211311213111000-0310230310021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002231112333033-3030111123323023-2302323101232032-2213123321303031-1013320130232020-2023132201011121-2233022201111100-1213012230303332"></a>

## timeouts — timeouts / 033123123212 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- timeouts

<a id="canonical-2012203122201330-0202232122021313-3203011322033312-2232211310120302-2333123321113211-2000202021203130-2330332022223331-0321310201303301"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322013002223023-2331223203130013-3112201330220301-1331211101012030-1122013230302032-1232011101221303-3313220230221110-1311122102212313"></a>

## Direct properties — timeouts / 033123123212 / 3

<a id="canonical-3023310131300223-3213333103321320-1232102303021221-2312333000121130-0132113032223331-1013322310320002-1111303210131023-2203221310331221"></a>

<a id="canonical-0110122233300123-1011000332113022-0032311200232220-1103113132221302-1320002020002323-1010111320201232-2232203102021303-3012301112110112"></a>

## create property — timeouts / 033123123212 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3011220002201003-3033032322130102-2033213212112011-2231301023212330-1211211220031001-1213210013102100-2020210121010302-3020020321211103"></a>

<a id="canonical-3110112033100002-0223311101221211-1010303311120300-1332120310300332-0031030232122001-1232131321003332-1330101131030332-2120231330310031"></a>

## delete property — timeouts / 033123123212 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0321130012201023-1032310002320330-2121001202002021-1312103301222322-3003020303021133-0332233223331030-0333111131321103-0212120033223200"></a>

<a id="canonical-3102122213302302-2302201231303232-0110000231330022-3033323210311101-2013221332330203-0321222103220221-0003133112120212-3320321321003303"></a>

## read property — timeouts / 033123123212 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0232121133112313-0302222333332221-1111022013123002-0231023012021111-1211103310121202-0003032303120300-2211010210320210-2231323033302122"></a>

<a id="canonical-1101131302133232-1301122112222013-2003130000111320-3033331013212031-2232213030112133-0300031130002322-0202033211121132-3131101202322232"></a>

## update property — timeouts / 033123123212 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0021200233322022-2011131323312332-1033211331311001-1031121012210002-3320120010222020-0313222023332323-2233000330121013-0223210132113311"></a>

## Next pages — timeouts / 033123123212 / 8

- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
