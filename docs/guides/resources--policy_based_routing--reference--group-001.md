---
page_title: "xcsh_policy_based_routing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing reference."
---

# xcsh_policy_based_routing reference

<a id="canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- Property reference

<a id="canonical-3022030000311320-3323001122001310-0000320302001110-0113012311201130-0001220101110221-1233133100023033-0022000033021202-3012033211011322"></a>

### Direct properties for `xcsh_policy_based_routing`

<a id="canonical-0210110301122112-2331301330200321-3221130013311020-3320313211033123-1032320232232102-2133201011202201-0003003223320203-2313133033000300"></a>

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

<a id="canonical-1101313001201302-1210102033233100-0033201032231113-2100213033223022-3320321310222021-0012231000210031-2030321232223011-2101133301313311"></a>

<a id="canonical-1101002012010300-1301321010112031-1020111310211322-0002330132233320-2220333230023200-0011000033030200-3223202121322221-2231233332310320"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0103101200031212-2003123023020312-1101232221203303-3121020001130332-0020313232201110-0200203033312110-2221113303313301-1132001001322033"></a>

<a id="canonical-2102311102103200-1003230020013222-1210212330103221-0010333331120032-2212020131021021-1132300032121031-1233022021331103-3100020013031312"></a>

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

- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220): complete subsection reference.

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-0121322222312032-0032020022000211-3121023130103220-3220223120230220-0100013220100200-1102310203201121-2102103313201303-2121203111213300): complete subsection reference.

<a id="canonical-2323220232112130-2212331000221102-3310112133002022-2211202232130111-0130333011020102-2122012322000313-3033103232230130-3310030133133103"></a>

<a id="canonical-3032230323130031-2012121020131220-1203311003210001-3100002022013020-1133313312120012-3013213212203113-1300111310310032-1001121001032100"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1212201201101202-1031031330232313-2213322003012201-1230222220313030-2232310121200023-0132021213303101-3212201312121013-3123303000332311"></a>

<a id="canonical-0323003003021201-1223223201230211-0010123001233122-3323202311201303-3022132230322033-3333333213003000-1300003301003120-2123203200021122"></a>

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

<a id="canonical-3212220223203313-0230223230130002-3013330323202313-0001003200133012-2130322212000130-1022101122211102-1231311113112021-3110013100133233"></a>

<a id="canonical-0302301230222301-0000021333012311-2311001131101221-1300001323311111-1321013201232331-0230103232031312-3310031102313000-2033013013212111"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Policy Based Routing. Must be unique within the namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0313032212231333-1111112033331200-0102311122300120-3102311233230130-1331032230323232-2222100302300320-2331013030103100-2102033002121020"></a>

<a id="canonical-3323310201222103-0323333220032331-1231013111131001-2232023323200302-2013212010222311-2112101213300220-3013131311330322-1102320322012021"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Policy Based Routing is created.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111): complete subsection reference.

- [timeouts](resources--policy_based_routing--reference--group-001.md#canonical-3212230023133031-3011210133323333-0032031130022121-1032331122221030-2200112213330131-2233103032101030-2303030130313323-2233130002020233): complete subsection reference.

<a id="canonical-3320312313230200-3103032312213212-3330230210332103-0323013311032021-1233113123112212-0010320202133122-2222302222112230-3111211021020332"></a>

### All schema paths for `xcsh_policy_based_routing`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--policy_based_routing--reference--group-001.md#canonical-0210110301122112-2331301330200321-3221130013311020-3320313211033123-1032320232232102-2133201011202201-0003003223320203-2313133033000300) |
| `description` | [description](resources--policy_based_routing--reference--group-001.md#canonical-1101313001201302-1210102033233100-0033201032231113-2100213033223022-3320321310222021-0012231000210031-2030321232223011-2101133301313311) |
| `disable` | [disable](resources--policy_based_routing--reference--group-001.md#canonical-0103101200031212-2003123023020312-1101232221203303-3121020001130332-0020313232201110-0200203033312110-2221113303313301-1132001001322033) |
| `forward_proxy_pbr` | [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3331302101033011-2302203312333020-0332110312133120-1313310312300030-1030211131003000-0211232132212220-0103230001312121-3013011203111301) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-2230322223231101-2123223031220121-2332310320223023-2113230001132322-0023023103011130-2212010222002220-2211101033211020-1122310023201212) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](resources--policy_based_routing--reference--group-001.md#canonical-1001330301112330-2011001132011320-0133210312011312-3021233112001101-0101221213111110-1202332210030103-2102002113300320-0310233123130220) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](resources--policy_based_routing--reference--group-001.md#canonical-2330021302102001-3121102303210020-1322121331222321-2211221213313330-3002333113331031-1202132211003010-0101111131212133-3112033120011302) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-0331203313102310-2310032211200030-1101333233032031-1211221103101320-0311210321212032-1211311033003023-0320133023020022-1123233310223300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-0120231221113212-2013003022103122-1033202330233002-1032212013330122-2320312022220323-2201311313311323-1210120031131220-3201131222031030) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-3000231110031120-3133001113130021-2103223013123332-3221133231203130-0032101002002111-2111302022333302-3320303321130330-3300111131101033) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-0101313002212321-3211220101013100-3012232001003230-0112233320201020-3230311201030201-3021201312223133-1302013000120021-2231031220201301) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-2003120231000131-1030120220203002-2102210021311301-2023120231230000-2100021031311203-2223333001000033-0101322020211320-1310203121102122) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-0232132210111101-0232300112202332-1021013201110203-3310103132300312-2321331310011131-2320200010201221-2313102213332333-1333113202212100) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](resources--policy_based_routing--reference--group-001.md#canonical-3231113233131323-0210232132231101-1203300002120123-0301303030103021-1132001211312101-0201222113020122-1033223033100032-2203102322112113) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](resources--policy_based_routing--reference--group-001.md#canonical-2010021330020310-0031010001310220-2213332301032213-0203212121233310-1023100001022131-3330132222101231-2230000313110220-1223131101220303) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](resources--policy_based_routing--reference--group-001.md#canonical-3001311131031202-3211330200011230-3212103031201010-1031303100120313-3021213300210010-0222300032313302-1300303001231101-0101132031100103) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](resources--policy_based_routing--reference--group-001.md#canonical-1110023311132003-1033032311322112-3132222102002120-2210330021233130-1010030230303010-3213301203202011-2331221103212230-2110132212303200) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](resources--policy_based_routing--reference--group-001.md#canonical-0022033101212331-2100210210332112-2320332030033132-0312101023331211-1130310301013300-3213322102310233-3012102112002023-3233212333021231) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](resources--policy_based_routing--reference--group-001.md#canonical-0021312103223020-1202130122322121-2202021032312232-0230031202221003-1033211032003111-2201322133013232-3303031311101011-2031232110323112) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](resources--policy_based_routing--reference--group-001.md#canonical-0030330123010232-2010221203222131-1023201021220222-1210332211222233-1303302313000101-3211022010303331-1333323133200110-2221013323322211) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-1020300330013213-0002123302010302-3230120120110001-2102200103111213-3120332213221301-0330221113201332-0032301231312233-1130122001210313) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](resources--policy_based_routing--reference--group-001.md#canonical-1203100033322221-0310223020000000-0220001123303022-0300123221212233-3110130123321023-0231130220210000-0002130121131112-3013202031000021) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](resources--policy_based_routing--reference--group-001.md#canonical-1203130013201202-2230012113303113-3002300100211202-1113203133220230-2233231233311010-1300221102021022-3231222032213103-2013232112301230) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](resources--policy_based_routing--reference--group-001.md#canonical-1030103211010201-1110200010100111-1211330202033312-3232032202331103-0003121300312310-3322323333131311-2303322313223100-2102301010121303) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-1031322011103101-3010203003322030-1222202303111330-2023023231201123-0111313322132031-0310332010022321-1213110303113113-0032323212022200) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](resources--policy_based_routing--reference--group-001.md#canonical-1102220031111231-3300202301301112-3100002010131223-2313222220302313-0102233033131100-2112010011312322-2213013333332130-3321013003133300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-0313011322303233-3001203330312213-1331111223230103-3301020021031320-3232222232100122-0130223323112300-1123331002001210-3300120101121002) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](resources--policy_based_routing--reference--group-001.md#canonical-2000003331000311-0022033133101333-2203323311322130-1130031301332102-0203130002100222-3331111210203031-2233112120001222-1012122301213322) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](resources--policy_based_routing--reference--group-001.md#canonical-2012000030323012-1220122330203300-2032220222000322-1310001113020012-2321311313030321-1113122333111100-0203002313013310-0210312220231300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-3012030113031220-0133021311232013-1332111000303012-1120212333002113-0310011332200222-3200021010302303-3100130322211331-0132021023321100) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-3301023211123033-1200220310231120-2110113021311010-1013202021232010-3121333330123023-1101222233111110-1012113011102203-2020002320122300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-1121202122312210-3213331120131323-0213102233300122-3310011311230320-0332021110032102-3321120100323232-0030203020312133-0100102023200131) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-0120130132113220-1001210000130002-1331311131021202-3013031200020321-0310202330102200-1121330030020010-1201323332301021-1001110330303123) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](resources--policy_based_routing--reference--group-001.md#canonical-3013230131313333-3221031111213210-0020223201211322-2333100331330010-3203131333203123-0121223030011003-2113131310210103-0122223210030000) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](resources--policy_based_routing--reference--group-001.md#canonical-0321021003333333-3132313323232012-3330222210033122-3103002230301202-1301021033130222-3123210202002302-1220312233230220-0333030100100120) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](resources--policy_based_routing--reference--group-001.md#canonical-1222213321230022-1113112311000002-2003313322330322-3202300320231113-0301010321310001-2022212313030101-1313312212023312-3133313023130031) |
| `forwarding_class_list` | [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-2010312311230233-2313302212132102-3320121333011110-1001111132111131-3331320310330031-0022030101033012-0123123013133113-2103301131221313) |
| `forwarding_class_list.name` | [forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-3023001131031322-0022021020303312-3112330131322122-2313102220100101-1223210131001320-3103031233013210-0012300211132330-1122223020030030) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-0132330113223130-1310212131221120-0011201212133331-0132001203123001-0012310222101132-2311313320101102-0131132211331111-2333300233032303) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-3110233123303222-0323201132031023-1232313122320213-2031013213320212-0023033300211123-0220030113311100-0102103011213021-0010102222321011) |
| `id` | [ID](resources--policy_based_routing--reference--group-001.md#canonical-2323220232112130-2212331000221102-3310112133002022-2211202232130111-0130333011020102-2122012322000313-3033103232230130-3310030133133103) |
| `labels` | [labels](resources--policy_based_routing--reference--group-001.md#canonical-1212201201101202-1031031330232313-2213322003012201-1230222220313030-2232310121200023-0132021213303101-3212201312121013-3123303000332311) |
| `name` | [name](resources--policy_based_routing--reference--group-001.md#canonical-3212220223203313-0230223230130002-3013330323202313-0001003200133012-2130322212000130-1022101122211102-1231311113112021-3110013100133233) |
| `namespace` | [namespace](resources--policy_based_routing--reference--group-001.md#canonical-0313032212231333-1111112033331200-0102311122300120-3102311233230130-1331032230323232-2222100302300320-2331013030103100-2102033002121020) |
| `network_pbr` | [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-1201321313330112-1013222113232013-2301123220303111-2020010313203101-0303032012233001-0131210322000023-0331221331113003-2001303131001213) |
| `network_pbr.any` | [network_pbr.any](resources--policy_based_routing--reference--group-001.md#canonical-3021123022132011-3133201213301121-1330302223111220-0221300101033101-1031201111023012-0011012311130022-3011331101001331-3010011011311100) |
| `network_pbr.label_selector` | [network_pbr.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-3012322032013233-1320020232010201-3021103322133323-1100033121000230-1111222200000121-2203231223203102-3121120200131220-1321203103011323) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](resources--policy_based_routing--reference--group-001.md#canonical-3113030233021113-3312213232200110-1233202002000301-3031021221330021-0113001010120001-0033210020313003-1013032002321000-2311022011013201) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-2021010103323311-2012131313000100-3213120032211112-3230031223200221-0300020013221123-3230211200012210-3113202331201101-0330312210113213) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-1233232112200320-0321022212222130-1320120031230330-0231013320321102-3321331111222323-3123002023233200-2110122010233310-0310211313203322) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](resources--policy_based_routing--reference--group-001.md#canonical-2211303222033333-2221232001321121-1213000200013032-0303123231312310-0030132301102202-0011332120322132-1230301020220111-1312013123330021) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-2301321012012101-2212331003200231-3211002230130232-1100132132203230-0122211323232231-2201222130112132-3011120223302132-1321210201101233) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](resources--policy_based_routing--reference--group-001.md#canonical-3210021123233332-3133332100003103-0002020102210112-0301213010322013-2100313301020232-2320020203133110-0231322022233133-1311331113213232) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](resources--policy_based_routing--reference--group-001.md#canonical-0032301130030321-2330322110223213-3233032323230131-3120200212213102-3120312333000232-1110030101001320-1311232003020211-1013033102303322) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](resources--policy_based_routing--reference--group-001.md#canonical-2301313231212311-0023102020032232-0233110020111221-3112031233313010-1231301121302000-1223003101102301-3110110103333300-2201102102112301) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](resources--policy_based_routing--reference--group-001.md#canonical-3022311002130013-1110301033222212-2321313131221321-2131110112101301-3020101120002003-0121113233332123-1102320012021033-2201323020030330) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-2110331301310330-2001220013123110-0021032122102210-0012000330302131-1112312202021001-0013310310230312-3333232020230230-0013011113310203) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-3301231112203101-1303303132303113-1232031130330002-2303220312212122-1002201222212221-2212200202003202-0122320123202022-1212212010133011) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-2230132121113003-0020230132201231-3120312102321201-1130323023310310-0222323033002011-2212103000201101-1331332031300132-2020200201013320) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-3102113031213102-0132022033133003-1223012003110320-3221033302020320-0221222310101213-2213211210020202-0331331301211020-2031232000210231) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-1013011103203120-2213113022000300-3133130032021210-0030020230120232-2122132031303322-0012131132320330-1200220331330100-0330332311301021) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](resources--policy_based_routing--reference--group-001.md#canonical-2323230001221023-2021123021021113-0221312333320131-3011330313312033-3221023121001111-2012233303032133-1021231030232131-3121302221002032) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](resources--policy_based_routing--reference--group-001.md#canonical-0332232301233222-0233332113310103-3112102221201011-1312211032111201-0303213012213002-1101021120232312-1132321210010131-3020032300000321) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](resources--policy_based_routing--reference--group-001.md#canonical-0132323030133133-3233011010032333-0302122333231310-0031303300322111-2230312131200103-1010223311300001-1021213320002311-2002012130231212) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](resources--policy_based_routing--reference--group-001.md#canonical-3312111202231322-2331031132232333-0302303013221020-1130222030330031-1122213211121311-3101100133233331-2033221021033313-1301321203230012) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](resources--policy_based_routing--reference--group-001.md#canonical-1202123023002222-2221212010312013-1212033130001132-3123020332130033-1023330213122311-2002122322323032-1003033033021112-2000332200312222) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](resources--policy_based_routing--reference--group-001.md#canonical-1022101131221211-0001221302320330-3303110230120201-1103031120313311-2203223210132032-1202111020212322-1312231201300112-3112122231331030) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-1121013003213003-3100303223311302-3033103310120021-3120011123301300-3112313110301100-2013033120113213-1112210313320132-2332231332102113) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](resources--policy_based_routing--reference--group-001.md#canonical-2202122021013322-1131031321323131-1230302223311210-2003123312203231-3003023321012111-1221302331011011-2301232130332300-3222322322300222) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](resources--policy_based_routing--reference--group-001.md#canonical-1202210222032100-2201003331313133-0020312232212211-0313312210132302-2301310220023130-0230330001310203-0331303010332021-1103033312120231) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-2120022210331131-2111331232323022-1002030300012032-1201123011310020-0213321111022200-0111223313002202-3213331230302030-3003223322000302) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-3011133030013332-1220133330102323-2000222121013310-1333131101033020-0100233020320230-0212223020121232-2303113132102132-2233013301201230) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](resources--policy_based_routing--reference--group-001.md#canonical-1311310232010101-0132000122012201-2132131232333210-2000121212332132-3300312312312002-2100230033132300-3220133202223302-3233030321332120) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](resources--policy_based_routing--reference--group-001.md#canonical-0203101301210303-1323012010203031-1122313103220030-2221103023203222-0212132130000020-0220201012112332-2233201012022222-0331211323301213) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](resources--policy_based_routing--reference--group-001.md#canonical-2212023311301233-1101022130000313-2102323100023323-2122031321233010-2113203001011221-2203201100310210-0103102010131000-3311031011302132) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-3102313222210323-1032112011031000-0203122000231030-2200113000000101-3220001212011000-2122112010031302-2302020100233122-2322033300031331) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-3231021223333330-3003311010103333-1021213221002132-3223130001330211-3111022110230303-0202213210200021-3100100123303221-2323203222201320) |
| `timeouts` | [timeouts](resources--policy_based_routing--reference--group-001.md#canonical-2333110033200120-2132131212120312-2001312123002000-1012310010312023-3301032000210313-3321212212313000-0211013003112220-3302212121332200) |
| `timeouts.create` | [timeouts.create](resources--policy_based_routing--reference--group-001.md#canonical-0101132111013122-3010220333112132-2221112100230331-3123132012213222-1223312333130011-0213022012323121-0321323102133210-3232302212210232) |
| `timeouts.delete` | [timeouts.delete](resources--policy_based_routing--reference--group-001.md#canonical-2130110223213101-1033321202020011-1333332202233102-1222010103012301-1031330201003203-3112331023122201-1312011230000012-3332312100132201) |
| `timeouts.read` | [timeouts.read](resources--policy_based_routing--reference--group-001.md#canonical-0121101002302203-0321203312021313-0200211002110010-3012300301212333-0310321032032231-1001312203020331-2220111103323303-1021022311110322) |
| `timeouts.update` | [timeouts.update](resources--policy_based_routing--reference--group-001.md#canonical-2021101020331302-1221300102111021-2332200101012103-2132201321322133-2202321132010231-1333213200020333-0202100131033300-3231022310100221) |

<a id="canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- forward_proxy_pbr

<a id="canonical-3331302101033011-2302203312333020-0332110312133120-1313310312300030-1030211131003000-0211232132212220-0103230001312121-3013011203111301"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Additional upstream details:

Network(L3/L4) routing policy rule.

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

- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3331302101033011-2302203312333020-0332110312133120-1313310312300030-1030211131003000-0211232132212220-0103230001312121-3013011203111301)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-1201321313330112-1013222113232013-2301123220303111-2020010313203101-0303032012233001-0131210322000023-0331221331113003-2001303131001213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
forward_proxy_pbr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133131200203001-3112022311301111-3120203332332010-1312003222021123-1231101113100130-2331120002011132-2302103301211233-0203131322112010"></a>

### Direct properties for `forward_proxy_pbr`

- [forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312): complete subsection reference.

<a id="canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="canonical-2230322223231101-2123223031220121-2332310320223023-2113230001132322-0023023103011130-2212010222002220-2211101033211020-1122310023201212"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_pbr_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120011313213330-3113001113010320-0101003221213321-0000303131033233-2120322201223003-1010120123303002-3330201001323221-1330332021300131"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules`

- [all_destinations](resources--policy_based_routing--reference--group-001.md#canonical-2003230233203132-0311032232211232-3030021301211033-1002333012032333-0012220020202310-0310320331113312-2330330311023031-0203232311300220): complete subsection reference.

- [all_sources](resources--policy_based_routing--reference--group-001.md#canonical-3020001330322011-0202021133322322-2032332001010210-2203130333333320-2032311102321230-0200201102031300-0010330323200001-2211330312301110): complete subsection reference.

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-3202003211132023-2320222111022101-0022010202011201-0201113120211303-3013230122122020-3011223313322322-1203130113102210-3212131012330210): complete subsection reference.

- [http_list](resources--policy_based_routing--reference--group-001.md#canonical-3210232000021131-3021110123332103-3230030102101223-0110213323123223-1303021013203230-1000303200111200-2313133032212200-1321001213133221): complete subsection reference.

- [ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-1122010131130233-0000032311210122-3032303222332223-2112213110300212-3300102312121220-0003012331311000-2133001020111032-3221222002023020): complete subsection reference.

- [label_selector](resources--policy_based_routing--reference--group-001.md#canonical-2331210003112101-2221132102203031-0012233033331012-2130122201230011-3112200201231302-1322331210200101-3110003210000012-0000112012210100): complete subsection reference.

- [metadata](resources--policy_based_routing--reference--group-001.md#canonical-1233323021302120-0000310013012200-0220111002112111-1002030100021322-1023330003210310-2221123233202032-2322113212000003-3103332011113330): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-2312210302013123-3211333123302303-0312222032021131-0221230031203013-1000222002101201-0330023121210122-2330220020030201-3203121022202021): complete subsection reference.

- [tls_list](resources--policy_based_routing--reference--group-001.md#canonical-2302000223132332-0120121212222231-3102303013320021-1233103202231113-2200123132102030-2110101311030132-0331111101023210-1311133032213030): complete subsection reference.

<a id="canonical-2003230233203132-0311032232211232-3030021301211033-1002333012032333-0012220020202310-0310320331113312-2330330311023031-0203232311300220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

<a id="canonical-1001330301112330-2011001132011320-0133210312011312-3021233112001101-0101221213111110-1202332210030103-2102002113300320-0310233123130220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

Additional upstream details:

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
all_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020001330322011-0202021133322322-2032332001010210-2203130333333320-2032311102321230-0200201102031300-0010330323200001-2211330312301110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_sources

<a id="canonical-2330021302102001-3121102303210020-1322121331222321-2211221213313330-3002333113331031-1202132211003010-0101111131212133-3112033120011302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

Additional upstream details:

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
all_sources = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202003211132023-2320222111022101-0022010202011201-0201113120211303-3013230122122020-3011223313322322-1203130113102210-3212131012330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list

<a id="canonical-0331203313102310-2310032211200030-1101333233032031-1211221103101320-0311210321212032-1211311033003023-0320133023020022-1123233310223300"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if no rule match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203112123310212-2013001322010220-2210312123002001-2012322122222210-0213230011330123-3001203010313211-0202233010123323-2003221031010101"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list`

<a id="canonical-0120231221113212-2013003022103122-1033202330233002-1032212013330122-2320312022220323-2201311313311323-1210120031131220-3201131222031030"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3000231110031120-3133001113130021-2103223013123332-3221133231203130-0032101002002111-2111302022333302-3320303321130330-3300111131101033"></a>

<a id="canonical-0111030102311300-3033110212313311-1201331032212310-0221100010333102-3002203212022030-1122102202330322-2333000210232202-1302210100100302"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0101313002212321-3211220101013100-3012232001003230-0112233320201020-3230311201030201-3021201312223133-1302013000120021-2231031220201301"></a>

<a id="canonical-2130031201332320-1122132312320210-3211222321011200-1300222200311200-3010031123310320-1013110122120313-1310112001132111-0310013200303110"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3210232000021131-3021110123332103-3230030102101223-0110213323123223-1303021013203230-1000303200111200-2313133032212200-1321001213133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list

<a id="canonical-2003120231000131-1030120220203002-2102210021311301-2023120231230000-2100021031311203-2223333001000033-0101322020211320-1310203121102122"></a>

Type: `"object"`. single nested block, Optional.

URLListType.

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231222013231110-3311103213112211-2201231201312002-0230010330212010-3012131303012000-0231230201303120-3033212302212323-1120102230310200"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.http_list`

- [http_list](resources--policy_based_routing--reference--group-001.md#canonical-1123213010023000-3323013000302211-1120100201010301-1301023110313303-1033031110300213-0200101203113303-0002303010330231-2333100020231111): complete subsection reference.

<a id="canonical-1123213010023000-3323013000302211-1120100201010301-1301023110313303-1033031110300213-0200101203113303-0002303010330231-2333100020231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-3210232000021131-3021110123332103-3230030102101223-0110213323123223-1303021013203230-1000303200111200-2313133032212200-1321001213133221)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list

<a id="canonical-0232132210111101-0232300112202332-1021013201110203-3310103132300312-2321331310011131-2320200010201221-2313102213332333-1333113202212100"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300000223012312-0211000200200330-2033200023013003-1111011220211132-1013013301030330-0102202313212022-2130110031201310-0003321313022031"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list`

- [any_path](resources--policy_based_routing--reference--group-001.md#canonical-2003113320031101-3310231322011303-0100312023231010-2331301323112112-1002302302101201-3310233101212020-0333311331332233-3233101101331020): complete subsection reference.

<a id="canonical-2010021330020310-0031010001310220-2213332301032213-0203212121233310-1023100001022131-3330132222101231-2230000313110220-1223131101220303"></a>

<a id="canonical-1330022233023220-1002021031112320-2212330231320303-2200033201001100-1123332033300112-1310003032320210-1320201010310222-2330300022213121"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3001311131031202-3211330200011230-3212103031201010-1031303100120313-3021213300210010-0222300032313302-1300303001231101-0101132031100103"></a>

<a id="canonical-1110333330333210-3122223021001103-1210000312110233-2113011223113012-3003330321011313-3212301323203331-0223333203023311-2030313112230013"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1110023311132003-1033032311322112-3132222102002120-2210330021233130-1010030230303010-3213301203202011-2331221103212230-2110132212303200"></a>

<a id="canonical-2112130312022133-3001102201002012-3201303331213302-3303030232322212-1100322113330013-2303023220003011-3120113112030003-3100120031133020"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Additional upstream details:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0022033101212331-2100210210332112-2320332030033132-0312101023331211-1130310301013300-3213322102310233-3012102112002023-3233212333021231"></a>

<a id="canonical-3103100323120333-0131200320233132-3321021020232322-1201100320311013-3132030012230101-0000002122233121-0030313020031010-2302303100121032"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0021312103223020-1202130122322121-2202021032312232-0230031202221003-1033211032003111-2201322133013232-3303031311101011-2031232110323112"></a>

<a id="canonical-3221220222123003-2320300110120012-3120122302102211-1022013323301302-1322111120021213-0033010033112122-1320022102321321-1212013211020012"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0030330123010232-2010221203222131-1023201021220222-1210332211222233-1303302313000101-3211022010303331-1333323133200110-2221013323322211"></a>

<a id="canonical-1233203232303332-0231312112202233-2120321023321132-2001210311110233-3010033211230133-2021211333303331-3211110322223230-0233231101321310"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2003113320031101-3310231322011303-0100312023231010-2331301323112112-1002302302101201-3310233101212020-0333311331332233-3233101101331020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-3210232000021131-3021110123332103-3230030102101223-0110213323123223-1303021013203230-1000303200111200-2313133032212200-1321001213133221)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-1123213010023000-3323013000302211-1120100201010301-1301023110313303-1033031110300213-0200101203113303-0002303010330231-2333100020231111)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path

<a id="canonical-3231113233131323-0210232132231101-1203300002120123-0301303030103021-1132001211312101-0201222113020122-1033223033100032-2203102322112113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122010131130233-0000032311210122-3032303222332223-2112213110300212-3300102312121220-0003012331311000-2133001020111032-3221222002023020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set

<a id="canonical-1020300330013213-0002123302010302-3230120120110001-2102200103111213-3120332213221301-0330221113201332-0032301231312233-1130122001210313"></a>

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331201332023130-0010331120320130-0332220120032231-3203310102133331-3113203013020013-3012110130330233-2303001022000211-0213103330120232"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set`

<a id="canonical-1203100033322221-0310223020000000-0220001123303022-0300123221212233-3110130123321023-0231130220210000-0002130121131112-3013202031000021"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1203130013201202-2230012113303113-3002300100211202-1113203133220230-2233231233311010-1300221102021022-3231222032213103-2013232112301230"></a>

<a id="canonical-2101332310033100-1021111032220213-1003223303122031-0122023101233033-2303123012031210-3203332233210331-3212112013223031-3023330332212203"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1030103211010201-1110200010100111-1211330202033312-3232032202331103-0003121300312310-3322323333131311-2303322313223100-2102301010121303"></a>

<a id="canonical-3001221313122330-1320003033011002-3320213320003100-2121030301012231-2310032021332223-3321120003323120-3020100012013002-1013230033333010"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2331210003112101-2221132102203031-0012233033331012-2130122201230011-3112200201231302-1322331210200101-3110003210000012-0000112012210100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.label_selector

<a id="canonical-1031322011103101-3010203003322030-1222202303111330-2023023231201123-0111313322132031-0310332010022321-1213110303113113-0032323212022200"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310310331102113-2102320301012012-2031301313332020-1100132002302100-1003302003330322-2102001323020111-1110231020203231-1021022333202033"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector`

<a id="canonical-1102220031111231-3300202301301112-3100002010131223-2313222220302313-0102233033131100-2112010011312322-2213013333332130-3321013003133300"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1233323021302120-0000310013012200-0220111002112111-1002030100021322-1023330003210310-2221123233202032-2322113212000003-3103332011113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.metadata

<a id="canonical-0313011322303233-3001203330312213-1331111223230103-3301020021031320-3232222232100122-0130223323112300-1123331002001210-3300120101121002"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231223000302300-1330121010211131-0330323033101223-3003200110321030-3130212323012201-2111331322032121-2201302201212310-3321031120010102"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.metadata`

<a id="canonical-2000003331000311-0022033133101333-2203323311322130-1130031301332102-0203130002100222-3331111210203031-2233112120001222-1012122301213322"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2012000030323012-1220122330203300-2032220222000322-1310001113020012-2321311313030321-1113122333111100-0203002313013310-0210312220231300"></a>

<a id="canonical-0100313313130123-2022122131013233-0000203333103000-0331220031333200-2330201303103333-1100120211312022-3102022110223031-2322022020111302"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2312210302013123-3211333123302303-0312222032021131-0221230031203013-1000222002101201-0330023121210122-2330220020030201-3203121022202021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list

<a id="canonical-3012030113031220-0133021311232013-1332111000303012-1120212333002113-0310011332200222-3200021010302303-3100130322211331-0132021023321100"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303311013303212-1103220022321333-3203113210002200-0102223112203313-3020320133330111-3303230323321201-3133012202320332-3233013330333001"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list`

<a id="canonical-3301023211123033-1200220310231120-2110113021311010-1013202021232010-3121333330123023-1101222233111110-1012113011102203-2020002320122300"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2302000223132332-0120121212222231-3102303013320021-1233103202231113-2200123132102030-2110101311030132-0331111101023210-1311133032213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

<a id="canonical-1121202122312210-3213331120131323-0213102233300122-3310011311230320-0332021110032102-3321120100323232-0030203020312133-0100102023200131"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002222201302211-1032121230011110-0220301111110333-1023022330213003-2123030201032003-0230222031222321-2113313202201001-1121023313211112"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list`

- [tls_list](resources--policy_based_routing--reference--group-001.md#canonical-3130302031320302-0222233202203000-3121302130000102-3320220210200113-3230000031210231-2223003213101331-0330103233123320-0030213233331322): complete subsection reference.

<a id="canonical-3130302031320302-0222233202203000-3121302130000102-3320220210200113-3230000031210231-2223003213101331-0330103233123320-0030213233331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1023022331001313-2101300112322121-1321113113010332-1211211323300300-1300220233000200-1322111330001221-2023000212123212-1222001323211312)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-2302000223132332-0120121212222231-3102303013320021-1233103202231113-2200123132102030-2110101311030132-0331111101023210-1311133032213030)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list

<a id="canonical-0120130132113220-1001210000130002-1331311131021202-3013031200020321-0310202330102200-1121330030020010-1201323332301021-1001110330303123"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302012333000210-1221312312032030-2000212331210020-3121233311102333-1002202300100030-2100023130321201-2001321331030220-0300210333313330"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list`

<a id="canonical-3013230131313333-3221031111213210-0020223201211322-2333100331330010-3203131333203123-0121223030011003-2113131310210103-0122223210030000"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0321021003333333-3132313323232012-3330222210033122-3103002230301202-1301021033130222-3123210202002302-1220312233230220-0333030100100120"></a>

<a id="canonical-2033220300102310-2212133131001333-1302132211313021-3110323230102110-0120202320032103-2220231131302102-1010220313223203-0131203001011010"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1222213321230022-1113112311000002-2003313322330322-3202300320231113-0301010321310001-2022212313030101-1313312212023312-3133313023130031"></a>

<a id="canonical-2130021023121211-2223110000230330-3131131021303221-1013111112301302-2012211110333330-3321012031103212-1133021003102020-3031301302020021"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0121322222312032-0032020022000211-3121023130103220-3220223120230220-0100013220100200-1102310203201121-2102103313201303-2121203111213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- forwarding_class_list

<a id="canonical-2010312311230233-2313302212132102-3320121333011110-1001111132111131-3331320310330031-0022030101033012-0123123013133113-2103301131221313"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if source application match and no rule match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300122210321331-0200231133330010-2331213230103010-3302213023133300-3001022003102000-0200111113103211-1110332310132231-1023032132233003"></a>

### Direct properties for `forwarding_class_list`

<a id="canonical-3023001131031322-0022021020303312-3112330131322122-2313102220100101-1223210131001320-3103031233013210-0012300211132330-1122223020030030"></a>

#### `forwarding_class_list.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0132330113223130-1310212131221120-0011201212133331-0132001203123001-0012310222101132-2311313320101102-0131132211331111-2333300233032303"></a>

<a id="canonical-2003100200003310-0333111101000103-2311233301311030-0322230222210103-3031211130330211-0010000021211310-0103122001222130-3003011223200320"></a>

#### `forwarding_class_list.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3110233123303222-0323201132031023-1232313122320213-2031013213320212-0023033300211123-0220030113311100-0102103011213021-0010102222321011"></a>

<a id="canonical-3020301032110320-3220020322002122-1133300112330312-3311131111133131-2211322112121313-3120333312211032-3231222001311302-3120332321131013"></a>

#### `forwarding_class_list.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- network_pbr

<a id="canonical-1201321313330112-1013222113232013-2301123220303111-2020010313203101-0303032012233001-0131210322000023-0331221331113003-2001303131001213"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for network pbr.

Additional upstream details:

Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list")}
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
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
network_pbr {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120332012102002-2022200221300212-1312221130030021-0100121312000322-1323032333201311-1113020032231322-0003131001030223-2030312301030013"></a>

### Direct properties for `network_pbr`

- [any](resources--policy_based_routing--reference--group-001.md#canonical-0110132210212022-0323320110302011-2130302030011302-3320110101133022-2323101112302122-3013121323013113-1223331032231330-1022011321212320): complete subsection reference.

- [label_selector](resources--policy_based_routing--reference--group-001.md#canonical-3203123000230311-1231313332103132-0222100333303120-1202121000321303-0320322321203012-2330330023202021-3213122101010221-0120201002310313): complete subsection reference.

- [network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-3101212232003021-2323211123110001-0012110333222331-1201011231230231-2102031123132001-1332130233302300-2230211020313030-0332232031120201): complete subsection reference.

<a id="canonical-0110132210212022-0323320110302011-2130302030011302-3320110101133022-2323101112302122-3013121323013113-1223331032231330-1022011321212320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.any` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- network_pbr.any

<a id="canonical-3021123022132011-3133201213301121-1330302223111220-0221300101033101-1031201111023012-0011012311130022-3011331101001331-3010011011311100"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203123000230311-1231313332103132-0222100333303120-1202121000321303-0320322321203012-2330330023202021-3213122101010221-0120201002310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.label_selector` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- network_pbr.label_selector

<a id="canonical-3012322032013233-1320020232010201-3021103322133323-1100033121000230-1111222200000121-2203231223203102-3121120200131220-1321203103011323"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011300333313030-2030300223031330-2331301322222030-0303003322232211-1010231112211101-3030230110300330-0101230000332033-3110330323033012"></a>

### Direct properties for `network_pbr.label_selector`

<a id="canonical-3113030233021113-3312213232200110-1233202002000301-3031021221330021-0113001010120001-0033210020313003-1013032002321000-2311022011013201"></a>

#### `network_pbr.label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- network_pbr.network_pbr_rules

<a id="canonical-2021010103323311-2012131313000100-3213120032211112-3230031223200221-0300020013221123-3230211200012210-3113202331201101-0330312210113213"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "dns_name"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("dns_name",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dns_name",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
network_pbr_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330320301131032-0121003011102303-1133131131331111-3303013021311031-0203022123313013-2231203232221113-3112222313011223-1211003003321322"></a>

### Direct properties for `network_pbr.network_pbr_rules`

- [all_tcp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-1121223032020001-2302301003103121-1220032100221132-0021211113110122-0020020121030330-0132100212321122-1130031010230300-3233123332003331): complete subsection reference.

- [all_traffic](resources--policy_based_routing--reference--group-001.md#canonical-3120011102123323-0223113231122221-3131333102122223-3203101123020210-1001112232031130-0303231312002011-3012321311210300-3131320030332232): complete subsection reference.

- [all_udp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-0302133032233321-0303030020102010-2202013030113130-0021003013011130-0111331022113303-1131120100111322-1030222202123113-1011311022021331): complete subsection reference.

- [any](resources--policy_based_routing--reference--group-001.md#canonical-0013202122230100-3113012323102133-3203233323330032-0033210202012300-0100032301331122-1300112013232330-1123113333121132-3010310120303101): complete subsection reference.

- [applications](resources--policy_based_routing--reference--group-001.md#canonical-0032303031332332-1013122323122030-1030322033202310-1221123210103022-0021100013323122-3233123102222330-0220321113111211-0113232201210330): complete subsection reference.

<a id="canonical-3022311002130013-1110301033222212-2321313131221321-2131110112101301-3020101120002003-0121113233332123-1102320012021033-2201323020030330"></a>

<a id="canonical-2031320121130110-0300303201312110-0320231333122232-1231300200103300-3030113020022111-2123001021220201-3303131102122121-3312003330201011"></a>

#### `network_pbr.network_pbr_rules.dns_name` property

Type: `"string"`. Optional.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-0002033331210132-3311210233002012-0123313002133202-1112203310233222-0101133301103221-2232210122333100-3030230033303012-2310312132330320): complete subsection reference.

- [ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-0220301111020322-0213132333011101-3030220100103223-3333332013320203-0210321310010013-2021213310222113-2311132013120122-0323322231211312): complete subsection reference.

- [metadata](resources--policy_based_routing--reference--group-001.md#canonical-0030210322012101-2001200331003010-3212012102213003-1311313130021123-1330110000311303-0023202003323311-1130003231302311-3030301303002031): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-2132132200110012-3022010332130221-1121221021313022-1102123031130313-2333131330021103-2302120021100123-0303000101203121-1130200322311332): complete subsection reference.

- [protocol_port_range](resources--policy_based_routing--reference--group-001.md#canonical-1120221331011201-0231320011213102-2103130230212201-0013232313233121-0300320213302033-2111221111300312-3312232303133101-2110332002232102): complete subsection reference.

<a id="canonical-1121223032020001-2302301003103121-1220032100221132-0021211113110122-0020020121030330-0132100212321122-1130031010230300-3233123332003331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="canonical-1233232112200320-0321022212222130-1320120031230330-0231013320321102-3321331111222323-3123002023233200-2110122010233310-0310211313203322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

Additional upstream details:

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
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120011102123323-0223113231122221-3131333102122223-3203101123020210-1001112232031130-0303231312002011-3012321311210300-3131320030332232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.all_traffic

<a id="canonical-2211303222033333-2221232001321121-1213000200013032-0303123231312310-0030132301102202-0011332120322132-1230301020220111-1312013123330021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

Additional upstream details:

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
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302133032233321-0303030020102010-2202013030113130-0021003013011130-0111331022113303-1131120100111322-1030222202123113-1011311022021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.all_udp_traffic

<a id="canonical-2301321012012101-2212331003200231-3211002230130232-1100132132203230-0122211323232231-2201222130112132-3011120223302132-1321210201101233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

Additional upstream details:

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
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013202122230100-3113012323102133-3203233323330032-0033210202012300-0100032301331122-1300112013232330-1123113333121132-3010310120303101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.any` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.any

<a id="canonical-3210021123233332-3133332100003103-0002020102210112-0301213010322013-2100313301020232-2320020203133110-0231322022233133-1311331113213232"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032303031332332-1013122323122030-1030322033202310-1221123210103022-0021100013323122-3233123102222330-0220321113111211-0113232201210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.applications` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.applications

<a id="canonical-0032301130030321-2330322110223213-3233032323230131-3120200212213102-3120312333000232-1110030101001320-1311232003020211-1013033102303322"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Additional upstream details:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231032020203121-1120102131313123-3230310212002003-2331203221130210-1103121320321113-0010202231010123-0132222112311022-0332011021110232"></a>

### Direct properties for `network_pbr.network_pbr_rules.applications`

<a id="canonical-2301313231212311-0023102020032232-0233110020111221-3112031233313010-1231301121302000-1223003101102301-3110110103333300-2201102102112301"></a>

#### `network_pbr.network_pbr_rules.applications.applications` property

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-0002033331210132-3311210233002012-0123313002133202-1112203310233222-0101133301103221-2232210122333100-3030230033303012-2310312132330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.forwarding_class_list

<a id="canonical-2110331301310330-2001220013123110-0021032122102210-0012000330302131-1112312202021001-0013310310230312-3333232020230230-0013011113310203"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if rule match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121302133323333-0222001201131200-2320032010120230-2013121010210112-3320001300130303-2103031202321031-0032333100023331-1111333203113230"></a>

### Direct properties for `network_pbr.network_pbr_rules.forwarding_class_list`

<a id="canonical-3301231112203101-1303303132303113-1232031130330002-2303220312212122-1002201222212221-2212200202003202-0122320123202022-1212212010133011"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230132121113003-0020230132201231-3120312102321201-1130323023310310-0222323033002011-2212103000201101-1331332031300132-2020200201013320"></a>

<a id="canonical-1330102302321311-2031101230032203-2202013122210231-1002132311123303-0101111132120220-3022220122233302-2233011022323213-3102032322232120"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3102113031213102-0132022033133003-1223012003110320-3221033302020320-0221222310101213-2213211210020202-0331331301211020-2031232000210231"></a>

<a id="canonical-2233002120132020-2101322230320110-1332312320113103-0213032001132113-2222103231022032-2031021231300222-3231020023131320-2110303301022301"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0220301111020322-0213132333011101-3030220100103223-3333332013320203-0210321310010013-2021213310222113-2311132013120122-0323322231211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="canonical-1013011103203120-2213113022000300-3133130032021210-0030020230120232-2122132031303322-0012131132320330-1200220331330100-0330332311301021"></a>

Type: `"object"`. single nested block, Optional.

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302200222330020-2310130000321113-1010323320002331-0001020321320021-3303333103230022-1220010320232100-2110120113320001-3303111232222313"></a>

### Direct properties for `network_pbr.network_pbr_rules.ip_prefix_set`

- [ref](resources--policy_based_routing--reference--group-001.md#canonical-2330213331133302-3320201021230013-1003001322110200-1323131110030130-1311201103120112-0322100201223123-1202130010001100-1302122333120122): complete subsection reference.

<a id="canonical-2330213331133302-3320201021230013-1003001322110200-1323131110030130-1311201103120112-0322100201223123-1202130010001100-1302122333120122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-0220301111020322-0213132333011101-3030220100103223-3333332013320203-0210321310010013-2021213310222113-2311132013120122-0323322231211312)
- network_pbr.network_pbr_rules.ip_prefix_set.ref

<a id="canonical-2323230001221023-2021123021021113-0221312333320131-3011330313312033-3221023121001111-2012233303032133-1021231030232131-3121302221002032"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100000010003120-1220220212101222-1323220033013100-1123212122313212-1310112132102110-3032012223303320-3121112222303020-0321321130132333"></a>

### Direct properties for `network_pbr.network_pbr_rules.ip_prefix_set.ref`

<a id="canonical-0332232301233222-0233332113310103-3112102221201011-1312211032111201-0303213012213002-1101021120232312-1132321210010131-3020032300000321"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0132323030133133-3233011010032333-0302122333231310-0031303300322111-2230312131200103-1010223311300001-1021213320002311-2002012130231212"></a>

<a id="canonical-0202331202230213-3100323020233330-2033020131100111-1021013302310012-2011113122222222-0212321021332022-3230110121222110-3130231320201233"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3312111202231322-2331031132232333-0302303013221020-1130222030330031-1122213211121311-3101100133233331-2033221021033313-1301321203230012"></a>

<a id="canonical-2220033322030233-3112112302011313-2311213001300032-3121320002212132-2132221032033012-3022011301030330-2231113310333120-1330012010203203"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1202123023002222-2221212010312013-1212033130001132-3123020332130033-1023330213122311-2002122322323032-1003033033021112-2000332200312222"></a>

<a id="canonical-0023123312300020-1230031003021000-3120032323231332-2030012110331010-1022220021233032-0203030213201322-3120103102233010-2012321120211001"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1022101131221211-0001221302320330-3303110230120201-1103031120313311-2203223210132032-1202111020212322-1312231201300112-3112122231331030"></a>

<a id="canonical-1320013332031032-2123021210221033-3112200222321212-1221230312130100-2130032002200312-1022030223333022-3033022012332311-3021332312303331"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0030210322012101-2001200331003010-3212012102213003-1311313130021123-1330110000311303-0023202003323311-1130003231302311-3030301303002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.metadata` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.metadata

<a id="canonical-1121013003213003-3100303223311302-3033103310120021-3120011123301300-3112313110301100-2013033120113213-1112210313320132-2332231332102113"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001203003102223-2030301011030331-2203123102311030-3002300223102231-3103122103023001-3103332100030320-0200302202223010-2030131021301121"></a>

### Direct properties for `network_pbr.network_pbr_rules.metadata`

<a id="canonical-2202122021013322-1131031321323131-1230302223311210-2003123312203231-3003023321012111-1221302331011011-2301232130332300-3222322322300222"></a>

#### `network_pbr.network_pbr_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1202210222032100-2201003331313133-0020312232212211-0313312210132302-2301310220023130-0230330001310203-0331303010332021-1103033312120231"></a>

<a id="canonical-0311113032322310-0101301030200321-1231031011021113-0320232323231102-3023033312010320-2020130300021023-0132022002332001-2103022201202010"></a>

#### `network_pbr.network_pbr_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2132132200110012-3022010332130221-1121221021313022-1102123031130313-2333131330021103-2302120021100123-0303000101203121-1130200322311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.prefix_list

<a id="canonical-2120022210331131-2111331232323022-1002030300012032-1201123011310020-0213321111022200-0111223313002202-3213331230302030-3003223322000302"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321110313102300-2021031210101211-3221301131202013-0303203133232212-2012131222031112-2321113333021333-0332000233012110-3002231110323010"></a>

### Direct properties for `network_pbr.network_pbr_rules.prefix_list`

<a id="canonical-3011133030013332-1220133330102323-2000222121013310-1333131101033020-0100233020320230-0212223020121232-2303113132102132-2233013301201230"></a>

#### `network_pbr.network_pbr_rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120221331011201-0231320011213102-2103130230212201-0013232313233121-0300320213302033-2111221111300312-3312232303133101-2110332002232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000)
- network_pbr.network_pbr_rules.protocol_port_range

<a id="canonical-1311310232010101-0132000122012201-2132131232333210-2000121212332132-3300312312312002-2100230033132300-3220133202223302-3233030321332120"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111103311110003-1300332310020221-3033010303022203-2303330200021003-2223300210311133-0223232220231303-3301000200220303-2303131320013210"></a>

### Direct properties for `network_pbr.network_pbr_rules.protocol_port_range`

<a id="canonical-0203101301210303-1323012010203031-1122313103220030-2221103023203222-0212132130000020-0220201012112332-2233201012022222-0331211323301213"></a>

#### `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-2212023311301233-1101022130000313-2102323100023323-2122031321233010-2113203001011221-2203201100310210-0103102010131000-3311031011302132"></a>

<a id="canonical-0133120001101230-0210212312033230-2133301103033012-2300031220101231-1030211201303132-3030322322001333-2302232002021123-1013320320220102"></a>

#### `network_pbr.network_pbr_rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALL","ICMP","TCP","UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-3101212232003021-2323211123110001-0012110333222331-1201011231230231-2102031123132001-1332130233302300-2230211020313030-0332232031120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-3301100031233122-1021230020121022-1211012201030321-3102231313022032-1011023232310033-0113000312313103-0102312232213111-0000213333033111)
- network_pbr.prefix_list

<a id="canonical-3102313222210323-1032112011031000-0203122000231030-2200113000000101-3220001212011000-2122112010031302-2302020100233122-2322033300031331"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022222222302220-0113111010303203-3130133012300222-2000211332313121-3201100103323100-0301133110210203-2221031232023330-3331311112302220"></a>

### Direct properties for `network_pbr.prefix_list`

<a id="canonical-3231021223333330-3003311010103333-1021213221002132-3223130001330211-3111022110230303-0202213210200021-3100100123303221-2323203222201320"></a>

#### `network_pbr.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3212230023133031-3011210133323333-0032031130022121-1032331122221030-2200112213330131-2233103032101030-2303030130313323-2233130002020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- timeouts

<a id="canonical-2333110033200120-2132131212120312-2001312123002000-1012310010312023-3301032000210313-3321212212313000-0211013003112220-3302212121332200"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323130233020332-3031110012330033-0210020213112103-1013130332213212-0211202103202331-0213123031312130-3333132132321212-2110232302313122"></a>

### Direct properties for `timeouts`

<a id="canonical-0101132111013122-3010220333112132-2221112100230331-3123132012213222-1223312333130011-0213022012323121-0321323102133210-3232302212210232"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2130110223213101-1033321202020011-1333332202233102-1222010103012301-1031330201003203-3112331023122201-1312011230000012-3332312100132201"></a>

<a id="canonical-3220233203312223-3132133303023220-0222011011111123-1200213102221010-1200200321023012-1201012212211320-1031012332030212-3310221023120332"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0121101002302203-0321203312021313-0200211002110010-3012300301212333-0310321032032231-1001312203020331-2220111103323303-1021022311110322"></a>

<a id="canonical-3133230233001321-3031333332313213-0021320033103032-1030300023331001-2000322123220200-1323212232232303-3113110031233122-1133323232122310"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2021101020331302-1221300102111021-2332200101012103-2132201321322133-2202321132010231-1333213200020333-0202100131033300-3231022310100221"></a>

<a id="canonical-0222331132132202-2112223320031330-0123032203301002-2233321313331123-3010230030011132-2033012203103012-0302121313210310-2230210120332021"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
