---
page_title: "xcsh_app_api_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group reference."
---

# xcsh_app_api_group reference

<a id="canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- Property reference

<a id="canonical-3101200123322223-1301010222232230-1323231323312330-3223103032203013-2323033023302203-2132033001213210-2111300302302020-1300003022001011"></a>

### Direct properties for `xcsh_app_api_group`

<a id="canonical-0112333122022130-0303110133222100-1330230030032103-3310023303122132-2213023032103002-0132220213300103-2331113031120202-3130111221002003"></a>

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

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-1031222130211103-1021233123310101-3110112321222203-0031313230202331-2233320301312110-2022232123223302-2101022333132131-1321310320121211): complete subsection reference.

- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-2122310110220121-1311210010012320-2130100120201203-2012322123033101-1200100321200231-2131013122300131-2032300202032303-2103323121312333): complete subsection reference.

<a id="canonical-1232122230112320-1322121112101202-2311202212202220-2103033211322232-1302312313022132-1113022121113231-3033000332020202-3310001013230331"></a>

<a id="canonical-0023230113310002-2311312020112133-2302103201133232-1000300121111021-0003122020302033-1021220111311330-3112123011011022-3022310002232201"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031020331100301-0111012022323000-0122200101120013-1222323302003303-1131213122131312-1300031120231331-1113012111202201-0232013230313033"></a>

<a id="canonical-3131211011203101-0120312210103013-2003031330023230-3310133001200212-1211133003213322-0332021333133333-2132033020233022-1001103133030320"></a>

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

- [elements](resources--app_api_group--reference--group-001.md#canonical-1200112120223122-2201203333220320-0111101133003203-2030033111231322-2122100331112303-0132313203213213-0220302212111033-0103230323020001): complete subsection reference.

- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-0111131201032300-2301200102322231-2002333313023333-3002103130232211-0212123110011020-1030313101300031-3221101213030123-0032303202233032): complete subsection reference.

<a id="canonical-1220201101200333-3103132323330110-0221230210132112-0132212033131333-0233020032303311-3213330113110203-2002133202223002-0322001230131112"></a>

<a id="canonical-0203303010030312-2000201331012022-3321332313100223-0301301010120133-2030333332232121-2113103200233322-2021313131101333-2213112223303310"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3303331130310012-2330333011021310-0001211100101013-0020110113303011-1112101323000223-0222211320012030-1003312230012221-3133220200021112"></a>

<a id="canonical-2013130312123032-1012113033000121-2211210012101301-0113331200030302-0110022002311123-3320021321332302-0323130021103013-3303101103123210"></a>

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

<a id="canonical-0113221300231002-0020102213333200-2002330201102321-0323012123111233-0211122033323020-2332232112121131-1221102301113133-2321133132331320"></a>

<a id="canonical-3322321320200221-3000300210223233-1201330111202211-2330303313212011-1231123300232022-1301100123221313-1301121003201222-2300023122312213"></a>

#### `name` property

Type: `"string"`. Required.

Name of the App API Group. Must be unique within the namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031112021221230-1113223210113202-3333121310212133-0001223110213232-0323220212213012-3231300131210113-2212013303321103-3332312002120033"></a>

<a id="canonical-0200011221321300-3033310132200333-3321331202223010-2313232020000033-1230301100313121-2022302103113033-1220302001222212-0013310123231113"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the App API Group is created.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [timeouts](resources--app_api_group--reference--group-001.md#canonical-0303001000233311-1211232313332220-1300310023210222-1101303201302202-3102211230212102-3130301013333030-0202112031022103-3210110033131113): complete subsection reference.

<a id="canonical-1302033333131332-3002001200211221-1313132333313313-0222210013100121-2220001332030211-0122030300121313-3232303212321001-2002332233301212"></a>

### All schema paths for `xcsh_app_api_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_api_group--reference--group-001.md#canonical-0112333122022130-0303110133222100-1330230030032103-3310023303122132-2213023032103002-0132220213300103-2331113031120202-3130111221002003) |
| `bigip_virtual_server` | [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-3233222030212230-0111003201311121-3121030301002113-0121232111310212-1313201012302200-1002311321231131-0103110020202322-2013111103032232) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-1000131233101021-1011002303322100-2033121102030033-0030123110110231-0213101001032220-1300321032013002-2011312103203032-2201111031122220) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](resources--app_api_group--reference--group-001.md#canonical-2200302320002031-2002221023030222-2233132101101123-3232303331131002-1100111112233222-3212201020230011-1122123000120110-3033130012333301) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](resources--app_api_group--reference--group-001.md#canonical-0323023111313220-3132023003203021-2003001300301030-1300130001312220-0030123300213121-1331230233102001-3321320022223230-1022322123333030) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](resources--app_api_group--reference--group-001.md#canonical-0030112010301301-0130310123322110-0132112111210112-0123132202301312-2230322213332121-1103233011120332-0223320312303031-1221322213010310) |
| `cdn_loadbalancer` | [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-2033220111203201-2221210101111033-1102032121201021-0311221002022223-3213002321331220-0112333020123120-0021302013101313-1233101223100330) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-2233123331302212-0133112330222303-1110220313131220-1033000300211032-0001203100221233-0113223323132122-2332333320123302-2223130211311312) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](resources--app_api_group--reference--group-001.md#canonical-0231023333113013-1111131130321030-2203210131220333-2102123023233021-3133112212222203-1322120330312323-2222300020330323-3303001101100110) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](resources--app_api_group--reference--group-001.md#canonical-2103321213132200-0133313030133212-1310022201322033-0210312122020303-3213302012130223-1333330312232020-1003323233332020-3122202210132310) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](resources--app_api_group--reference--group-001.md#canonical-2100201203032313-0203300212332221-2312310100210131-0121013333020302-3233012220021300-2123322122230021-0210202001312312-2030203030003100) |
| `description` | [description](resources--app_api_group--reference--group-001.md#canonical-1232122230112320-1322121112101202-2311202212202220-2103033211322232-1302312313022132-1113022121113231-3033000332020202-3310001013230331) |
| `disable` | [disable](resources--app_api_group--reference--group-001.md#canonical-0031020331100301-0111012022323000-0122200101120013-1222323302003303-1131213122131312-1300031120231331-1113012111202201-0232013230313033) |
| `elements` | [elements](resources--app_api_group--reference--group-001.md#canonical-1030033312022321-2130232211231022-2211002121313032-2133101002313333-0230333010112310-1212011012123300-2133332321300001-3210310103223221) |
| `elements.methods` | [elements.methods](resources--app_api_group--reference--group-001.md#canonical-1132311202301021-2132211312212022-0020221022322300-3313303303033230-2010210330131332-1112333001210333-3302103111302033-2022032023033330) |
| `elements.path_regex` | [elements.path_regex](resources--app_api_group--reference--group-001.md#canonical-1111113221313133-3030220111130112-0210133231210113-1303110233210123-0132333132310232-3032303122021001-2020122201331230-2321121032103222) |
| `http_loadbalancer` | [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-0313313132110223-1111300010313110-2210000310030030-0311020102232310-2130300312032213-2102213232220303-0033232000103230-0232011301230112) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-1331322022213120-2022133013203030-2030120022312203-1331133120102101-2110010300003331-2231011100031123-2002221002330103-0000303302221011) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](resources--app_api_group--reference--group-001.md#canonical-1100232331320303-1033012332223001-3221230231201110-0200102313300313-3000032221201332-1303232303002021-0122303012230211-2113223223300000) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](resources--app_api_group--reference--group-001.md#canonical-1222033033011112-2223021030003113-1303213312031300-2222233011011011-2201213202132210-1011320020320001-1212120220101001-2001032112223122) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](resources--app_api_group--reference--group-001.md#canonical-1331320210323002-2002212122323302-0320331200121230-1211232303302300-0132331311201313-2312003111310311-2032321003100103-3012211300100132) |
| `id` | [ID](resources--app_api_group--reference--group-001.md#canonical-1220201101200333-3103132323330110-0221230210132112-0132212033131333-0233020032303311-3213330113110203-2002133202223002-0322001230131112) |
| `labels` | [labels](resources--app_api_group--reference--group-001.md#canonical-3303331130310012-2330333011021310-0001211100101013-0020110113303011-1112101323000223-0222211320012030-1003312230012221-3133220200021112) |
| `name` | [name](resources--app_api_group--reference--group-001.md#canonical-0113221300231002-0020102213333200-2002330201102321-0323012123111233-0211122033323020-2332232112121131-1221102301113133-2321133132331320) |
| `namespace` | [namespace](resources--app_api_group--reference--group-001.md#canonical-0031112021221230-1113223210113202-3333121310212133-0001223110213232-0323220212213012-3231300131210113-2212013303321103-3332312002120033) |
| `timeouts` | [timeouts](resources--app_api_group--reference--group-001.md#canonical-3122133313331122-3232332110012210-0300333210201212-2123313302313121-2332311032123211-3311133330003121-3133000021111011-2102113331233122) |
| `timeouts.create` | [timeouts.create](resources--app_api_group--reference--group-001.md#canonical-0002332110002003-0233132221321130-3120332012111000-3330112112201331-2112110003231202-3302021213211212-3010131113332030-0332100212130000) |
| `timeouts.delete` | [timeouts.delete](resources--app_api_group--reference--group-001.md#canonical-1021311121203010-3030210222211222-0333321110033111-3131001111123022-3300120110200101-0232222132323122-3003222222321231-1330122100202333) |
| `timeouts.read` | [timeouts.read](resources--app_api_group--reference--group-001.md#canonical-3321113130112212-3332132012323020-1103021303022122-0321112323100213-0113313122113311-1012311012120103-3013312330213012-3222111201120000) |
| `timeouts.update` | [timeouts.update](resources--app_api_group--reference--group-001.md#canonical-1330333300231321-0313333213223332-1313102133203110-0032332023322113-1001223022000013-2213321121230322-0131013132220010-3303223202332200) |

<a id="canonical-1031222130211103-1021233123310101-3110112321222203-0031313230202331-2233320301312110-2022232123223302-2101022333132131-1321310320121211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bigip_virtual_server` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- bigip_virtual_server

<a id="canonical-3233222030212230-0111003201311121-3121030301002113-0121232111310212-1313201012302200-1002311321231131-0103110020202322-2013111103032232"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

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

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-3233222030212230-0111003201311121-3121030301002113-0121232111310212-1313201012302200-1002311321231131-0103110020202322-2013111103032232)
- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-2033220111203201-2221210101111033-1102032121201021-0311221002022223-3213002321331220-0112333020123120-0021302013101313-1233101223100330)
- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-0313313132110223-1111300010313110-2210000310030030-0311020102232310-2130300312032213-2102213232220303-0033232000103230-0232011301230112)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bigip_virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320322100002022-2212311003201123-2030221311132120-3303101120132012-1303323010202030-1230011333102233-1220022133121133-0011201000021202"></a>

### Direct properties for `bigip_virtual_server`

- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-0103200330031331-1120100022232112-0101011011220323-0223231301021130-0222002131113133-2101331021211323-2031101010310020-3332002130101111): complete subsection reference.

<a id="canonical-0103200330031331-1120100022232112-0101011011220323-0223231301021130-0222002131113133-2101331021211323-2031101010310020-3332002130101111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bigip_virtual_server.bigip_virtual_server` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- [bigip_virtual_server](resources--app_api_group--reference--group-001.md#canonical-1031222130211103-1021233123310101-3110112321222203-0031313230202331-2233320301312110-2022232123223302-2101022333132131-1321310320121211)
- bigip_virtual_server.bigip_virtual_server

<a id="canonical-1000131233101021-1011002303322100-2033121102030033-0030123110110231-0213101001032220-1300321032013002-2011312103203032-2201111031122220"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
bigip_virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201331102312221-0011131121132200-1112323323111220-0130023000001010-1113120103101313-2021200030123201-0010202010302203-3221011333233200"></a>

### Direct properties for `bigip_virtual_server.bigip_virtual_server`

<a id="canonical-2200302320002031-2002221023030222-2233132101101123-3232303331131002-1100111112233222-3212201020230011-1122123000120110-3033130012333301"></a>

#### `bigip_virtual_server.bigip_virtual_server.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0323023111313220-3132023003203021-2003001300301030-1300130001312220-0030123300213121-1331230233102001-3321320022223230-1022322123333030"></a>

<a id="canonical-0112332223213113-3233323131220230-2210202003210132-1202001010330313-1322210233221010-0303202022322001-0020121010112330-3100133010001102"></a>

#### `bigip_virtual_server.bigip_virtual_server.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0030112010301301-0130310123322110-0132112111210112-0123132202301312-2230322213332121-1103233011120332-0223320312303031-1221322213010310"></a>

<a id="canonical-3300302020320010-2230033211322103-1221033113313233-0030133311333230-2122322210132130-2010103313303103-1200122011302232-0233221212331112"></a>

#### `bigip_virtual_server.bigip_virtual_server.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2122310110220121-1311210010012320-2130100120201203-2012322123033101-1200100321200231-2131013122300131-2032300202032303-2103323121312333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cdn_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- cdn_loadbalancer

<a id="canonical-2033220111203201-2221210101111033-1102032121201021-0311221002022223-3213002321331220-0112333020123120-0021302013101313-1233101223100330"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific CDN Loadbalancer.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101320230213133-3323302313033300-2331312321201310-2333300311103011-2232011000130320-3131102230020003-0333201213312100-1033012103301111"></a>

### Direct properties for `cdn_loadbalancer`

- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-1002300203112122-1122202021312003-2111302210111103-0033203310200313-1233301100110021-0101120133201330-1320332131310321-1233302213231033): complete subsection reference.

<a id="canonical-1002300203112122-1122202021312003-2111302210111103-0033203310200313-1233301100110021-0101120133201330-1320332131310321-1233302213231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cdn_loadbalancer.cdn_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- [cdn_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-2122310110220121-1311210010012320-2130100120201203-2012322123033101-1200100321200231-2131013122300131-2032300202032303-2103323121312333)
- cdn_loadbalancer.cdn_loadbalancer

<a id="canonical-2233123331302212-0133112330222303-1110220313131220-1033000300211032-0001203100221233-0113223323132122-2332333320123302-2223130211311312"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132320202000302-2033300312131100-3023202001103213-0012310100133133-1003303033212122-1121120101221211-3201203333021120-2233200320010130"></a>

### Direct properties for `cdn_loadbalancer.cdn_loadbalancer`

<a id="canonical-0231023333113013-1111131130321030-2203210131220333-2102123023233021-3133112212222203-1322120330312323-2222300020330323-3303001101100110"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2103321213132200-0133313030133212-1310022201322033-0210312122020303-3213302012130223-1333330312232020-1003323233332020-3122202210132310"></a>

<a id="canonical-1332331133312320-1130333213122100-2323030213332020-3131312100311113-1302232033210231-1210102330121131-2023011121233100-3200320313011112"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2100201203032313-0203300212332221-2312310100210131-0121013333020302-3233012220021300-2123322122230021-0210202001312312-2030203030003100"></a>

<a id="canonical-3202022020120301-0021133021222100-3230102031302122-3331011010023301-2200001030120310-0333033303300322-1012023030220023-3031212120230112"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1200112120223122-2201203333220320-0111101133003203-2030033111231322-2122100331112303-0132313203213213-0220302212111033-0103230323020001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `elements` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- elements

<a id="canonical-1030033312022321-2130232211231022-2211002121313032-2133101002313333-0230333010112310-1212011012123300-2133332321300001-3210310103223221"></a>

Type: `"object"`. list nested block, Optional.

List of API group elements with methods and path regular expression for matching requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("methods",
    "path_regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  }
}
```

Terraform syntax:

```terraform
elements {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001113200233302-3022332003001112-0302123023321122-1002032022101220-2123013022131233-3321110131123222-0202120223323023-0322321322001003"></a>

### Direct properties for `elements`

<a id="canonical-1132311202301021-2132211312212022-0020221022322300-3313303303033230-2010210330131332-1112333001210333-3302103111302033-2022032023033330"></a>

#### `elements.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of method values to
match the input request API method against. The match is considered to succeed if the input request
API method is a member of the list. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`,
\`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1111113221313133-3030220111130112-0210133231210113-1303110233210123-0132333132310232-3032303122021001-2020122201331230-2321121032103222"></a>

<a id="canonical-3102102313031113-1311132232302221-0211121312321320-2110320021102032-2023023011012013-2202211020121031-1312203101322121-1131030202202132"></a>

#### `elements.path_regex` property

Type: `"string"`. Optional.

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regular expression.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0111131201032300-2301200102322231-2002333313023333-3002103130232211-0212123110011020-1030313101300031-3221101213030123-0032303202233032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- http_loadbalancer

<a id="canonical-0313313132110223-1111300010313110-2210000310030030-0311020102232310-2130300312032213-2102213232220303-0033232000103230-0232011301230112"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132310231321233-3321210323320132-3302321033220031-3000111213033202-3231020212023131-1221132203301323-3130321211323221-1333203231332302"></a>

### Direct properties for `http_loadbalancer`

- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-0203322012023000-2321132233303231-2032312330112331-1331330333031203-0232132011300023-0023010311322132-0202210033233130-3312321230200022): complete subsection reference.

<a id="canonical-0203322012023000-2321132233303231-2032312330112331-1331330333031203-0232132011300023-0023010311322132-0202210033233130-3312321230200022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_loadbalancer.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- [http_loadbalancer](resources--app_api_group--reference--group-001.md#canonical-0111131201032300-2301200102322231-2002333313023333-3002103130232211-0212123110011020-1030313101300031-3221101213030123-0032303202233032)
- http_loadbalancer.http_loadbalancer

<a id="canonical-1331322022213120-2022133013203030-2030120022312203-1331133120102101-2110010300003331-2231011100031123-2002221002330103-0000303302221011"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001202211333012-0213033131221003-0131012020221331-2310131222301013-0322210030330222-3131021230301311-2321210230322100-3031112022102122"></a>

### Direct properties for `http_loadbalancer.http_loadbalancer`

<a id="canonical-1100232331320303-1033012332223001-3221230231201110-0200102313300313-3000032221201332-1303232303002021-0122303012230211-2113223223300000"></a>

#### `http_loadbalancer.http_loadbalancer.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1222033033011112-2223021030003113-1303213312031300-2222233011011011-2201213202132210-1011320020320001-1212120220101001-2001032112223122"></a>

<a id="canonical-3232010312323001-0121330312032110-0112101003133200-0031130123230201-0101002011132013-3030131202032332-2032110030101110-1230330212323233"></a>

#### `http_loadbalancer.http_loadbalancer.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1331320210323002-2002212122323302-0320331200121230-1211232303302300-0132331311201313-2312003111310311-2032321003100103-3012211300100132"></a>

<a id="canonical-1320100213313333-2111013102110102-1033331202332313-1021130103332210-1220122030231032-3123220031312300-3233323002323311-2123302030012302"></a>

#### `http_loadbalancer.http_loadbalancer.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0303001000233311-1211232313332220-1300310023210222-1101303201302202-3102211230212102-3130301013333030-0202112031022103-3210110033131113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Property reference](resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- timeouts

<a id="canonical-3122133313331122-3232332110012210-0300333210201212-2123313302313121-2332311032123211-3311133330003121-3133000021111011-2102113331233122"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021330132120210-1301312110030302-3203310211213201-3110102322113333-0111331030001133-2312112313112102-3202122301103021-2120011133311001"></a>

### Direct properties for `timeouts`

<a id="canonical-0002332110002003-0233132221321130-3120332012111000-3330112112201331-2112110003231202-3302021213211212-3010131113332030-0332100212130000"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1021311121203010-3030210222211222-0333321110033111-3131001111123022-3300120110200101-0232222132323122-3003222222321231-1330122100202333"></a>

<a id="canonical-3123111202310313-0103110131130131-3332131211132113-2302321312230332-1032023222112232-3110210221001323-2201011323132332-1213131033212110"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3321113130112212-3332132012323020-1103021303022122-0321112323100213-0113313122113311-1012311012120103-3013312330213012-3222111201120000"></a>

<a id="canonical-3010332111201323-2231223130110000-0030120022221302-0130231103132203-3232112210311222-0322332203300100-3032211231033201-2332001300031102"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1330333300231321-0313333213223332-1313102133203110-0032332023322113-1001223022000013-2213321121230322-0131013132220010-3303223202332200"></a>

<a id="canonical-3311011300101332-2023332111301232-3230200123123222-2333323210111303-2111230100310003-3110301021113331-1211001223013210-3111000222022301"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
