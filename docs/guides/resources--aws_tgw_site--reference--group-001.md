---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- Property reference

<a id="canonical-1203013201012103-2121201032000021-3100132200232130-1231232323302002-1301033130123101-2332303132223311-3012012110002012-0231020021022231"></a>

### Direct properties for `xcsh_aws_tgw_site`

<a id="canonical-0022302130100131-3313303302012100-3313302002203233-1010322210323200-3110222302320302-3022233123333333-0320301333221220-3322321033313310"></a>

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

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020): complete subsection reference.

- [block_all_services](resources--aws_tgw_site--reference--group-001.md#canonical-3303203233101111-1231211011103232-0133220301212023-0131322023033222-3032123223330012-1020121030201121-2102031230301211-3211000022222033): complete subsection reference.

- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101): complete subsection reference.

- [coordinates](resources--aws_tgw_site--reference--group-001.md#canonical-3330100032212111-2023231232130113-2033223012210111-0102030223332333-1003211333301322-2022000120201201-1202112101331102-3200002321100010): complete subsection reference.

- [custom_dns](resources--aws_tgw_site--reference--group-001.md#canonical-0310130113103331-1321231230201233-0112322012222330-2020103011100211-0311022022233022-2032112020200200-0003021332313331-1212231013022131): complete subsection reference.

- [default_blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-0232210132032302-2311321302133012-3210002210020112-0120013313103120-1000032132332201-1332320121032013-1330221230222303-1221113233030112): complete subsection reference.

<a id="canonical-3131303101010321-1301133302311012-0000303112322110-3232100110123112-0121020132202310-1010101131031233-2310003002131033-2020121220203010"></a>

<a id="canonical-2313120011123022-3032232201103301-0023000333233121-3000021222010131-2220233230020231-1232121113310032-1310221122020313-1331223101210210"></a>

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

- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-0300022120212210-3312311111132100-2030111122121212-3323031300112111-1131201020023012-1112013201302223-0312310103022301-3030112201131321): complete subsection reference.

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123): complete subsection reference.

<a id="canonical-3221112231220130-2220123003301123-1132332231000333-3223332232333212-3200211000020332-3030210130300030-0313323201320110-3102223321301333"></a>

<a id="canonical-2102332023332131-0032011301322003-3202221111130321-0133320131100212-3311220300013210-0311320001102202-3103003033012021-0221332202332013"></a>

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

<a id="canonical-1120133123333333-1033102312020012-0203332101200233-3211322330103221-2312021311312101-1330203132213321-1130223131210121-0122303120003121"></a>

<a id="canonical-3330132001222102-1231031203320112-3031303202313000-1110313333023022-0022021010102232-1013103321001200-2210203120232331-3223130312310232"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332): complete subsection reference.

<a id="canonical-2230020112100311-2000012201220330-3201122333302222-1332223120111201-0322330112331210-3102312301222323-3011111101130223-3203313003300023"></a>

<a id="canonical-3002221222022022-0123021220123021-0230313311122202-1102120212023132-1232230223001302-3022111031220222-0230113321320301-0200322031002023"></a>

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

- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-1103022023311031-0103303110232123-3012233320211332-2012113211332023-0001320020120012-1023313201031223-2223210013123113-1123000001332222): complete subsection reference.

- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-3331230112202213-2130303012010322-2331122233103201-0333302220133100-3033003202003300-0301310322330232-2200123022312220-2023201330131100): complete subsection reference.

<a id="canonical-2130021320201111-3023222110200310-0310032211130121-0313123012133000-1312213120333123-1110131212212123-1233212011331100-3221133231002200"></a>

<a id="canonical-0100020000131011-0332022221301330-0031011332032223-3203023012200100-1303120311202300-0333330211200123-0031301313200122-3312301201012211"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AWS TGW Site. Must be unique within the namespace.

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

<a id="canonical-0100323121210313-3033112313322113-1221121123313210-2120203300013102-1220100200230301-2031222210311322-3111001213130213-3230011331330301"></a>

<a id="canonical-3313010102002332-3002033120111000-0101201132013230-2313302030021031-3111021111002122-1333131023333213-1210211221330330-3323033213210100"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AWS TGW Site is created.

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

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323): complete subsection reference.

- [os](resources--aws_tgw_site--reference--group-002.md#canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030): complete subsection reference.

- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113): complete subsection reference.

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211): complete subsection reference.

- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231): complete subsection reference.

<a id="canonical-1112300100113320-0232112303200010-3113201323211213-3230120011220002-3110231103033023-1333121130310031-0110011122100202-3300031031233211"></a>

<a id="canonical-2002112313323301-1200301130313202-2112302331000110-1302113220033133-0010133330322003-1133013333313322-3132012000333212-1123312001110213"></a>

#### `tags` property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{
  validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030): complete subsection reference.

- [timeouts](resources--aws_tgw_site--reference--group-003.md#canonical-1311110330231022-3020003023311003-2213221230300113-1201311012233213-1111213120311303-1312303200031100-2033333131211220-2303001123013033): complete subsection reference.

- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031): complete subsection reference.

- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303): complete subsection reference.

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032): complete subsection reference.

<a id="canonical-3023213203030011-0231122312122203-3023332011021220-0312233233311132-0212102033102312-3300221022211023-2031133332002313-1310233220221210"></a>

### All schema paths for `xcsh_aws_tgw_site`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--aws_tgw_site--reference--group-001.md#canonical-0022302130100131-3313303302012100-3313302002203233-1010322210323200-3110222302320302-3022233123333333-0320301333221220-3322321033313310) |
| `aws_parameters` | [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2320313303213230-0001313321120112-3320121311232221-0023233120021303-0201303023333210-0111222203013200-2232103003000221-2113211211012122) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-1013333212210132-2113020132003113-3202320333103133-0202000022220133-3231131331233301-2030030311231131-2122302220130202-1222231220012113) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-1022312320122023-2303103323112111-3332331301012102-3322102000221012-0313032120200113-3211103132333301-3113321231212311-2001233002202113) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](resources--aws_tgw_site--reference--group-001.md#canonical-0300121232211101-0030011201000131-3110000020032100-3332203303312312-0333211022030110-1220011230232021-0010323320323001-0030023300221333) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](resources--aws_tgw_site--reference--group-001.md#canonical-2100020320203101-2313033213323100-3200011320023303-1101132231222021-2102311200333123-1123122110322132-3130203212121023-1112211331110000) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](resources--aws_tgw_site--reference--group-001.md#canonical-3200100313323332-3133032132032201-0301330221221332-0333333012030203-3010320032313320-3000101010320311-2103333213220022-3211200213103110) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-2111222133222033-2011002113331213-2200100102213331-0212333033320232-2232003113111212-3301003312021202-1012131221111232-2213000332030202) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](resources--aws_tgw_site--reference--group-001.md#canonical-2321200220223022-2222212023120301-1202301123210223-2113223231032021-3112212330112320-0213203322023231-2131223022333330-2013201113033101) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](resources--aws_tgw_site--reference--group-001.md#canonical-2301213321132320-3323303330301102-2030200020231311-1233031302201123-2300313102333032-0013032232333230-3321220121300000-2011022231220333) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](resources--aws_tgw_site--reference--group-001.md#canonical-2233233113100032-0201022203031130-0122222132002002-3213331333100130-3023312203302013-3333111301123023-0023223202023032-0023302303222031) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](resources--aws_tgw_site--reference--group-001.md#canonical-2222002012302011-1122023323113011-2100122012230123-2201131200003031-1132320213321230-3311030201121102-0330101200320301-2022230321213012) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](resources--aws_tgw_site--reference--group-001.md#canonical-0310123300232330-0023232120232001-3212313113312203-3330100131302330-3103022010220121-2313033212221321-3012121112011020-2221133220210131) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](resources--aws_tgw_site--reference--group-001.md#canonical-3231333223323012-2211202102121002-1210332031001313-1303132133221310-0000323323203221-0222301110301211-2333033023310203-1321132030232101) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](resources--aws_tgw_site--reference--group-001.md#canonical-3312123123233201-0122110220031001-1203201230012111-2101033213131023-1033113331300200-1300210130302311-2210010012211010-0000133212111121) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-3323022333231001-3130003120022323-2010303312000303-1231022001333302-1000130113313000-0003001101331032-0021123000223011-3223022112323311) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](resources--aws_tgw_site--reference--group-001.md#canonical-1012030312111120-1030033021223300-2320231023100003-2132210111322123-1020002030201133-3000302200111001-2303022223122213-2102320313033222) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2321010031302032-1310031213133333-2020133002023220-0111020112013002-1033101111120332-3230030023231030-2201203322203012-2123002100200011) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-1031321212232331-0030201301102203-0313321120203311-1110100323302233-2131303220103030-2130031322132231-0001031113322002-2311001010323131) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-2113313100022212-2222103332300313-1302213332222023-1231032232011313-2323110002122022-2112022310203123-3212002002022220-3110310131110032) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-1301123000130032-0102101120213233-3011312201202200-2031321230020212-1110233031133112-2121002123320201-0111301023333021-1121203210201030) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-0010211200331010-0210310223003110-2310130030100100-2030333020022022-1011301300120103-0032332003302020-0100033101301220-2323303323212000) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-2300311121033220-0031011300302113-1123011322233222-2022021222200101-3101102112033211-1330211212122220-1101020023112320-2223302212312021) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-0113203113012330-3321013201201131-1310202333301132-1101030121010331-3203300122000202-3003220111330023-1303232022021112-0032233310301211) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-0111311133303133-0022332020223211-1101301230321002-2111300022021312-2000123120322003-3102023131212300-3023131311101102-3102310100002110) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2210332001131001-2100203001320031-3312321113332121-1321222330300203-3223100201312332-3203130132211020-0333030212330330-0131121220013233) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-3221221122222330-0131322203023032-2032122131312330-3231321301333113-1003103002131212-3201312012320032-3001212331322101-3202222102001000) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-1013121112232300-2312201321123322-3020202212032213-2320303123101001-3002233313033110-3010113110123212-0322011233303323-2020010232033310) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-0023011230113121-3123233010303310-1130021023100002-2230120322303120-1230211120333223-2332330200222101-1010021320220312-2311030002123210) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-0121030331313110-0121232222020130-1230322120310323-3030103212222233-2003011031312200-2201221210130130-1330002223100232-3031203111113302) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-3100211220330203-3212203213222111-1113200212122020-0333000030223231-2332102013230202-1322311003201321-0322221110231020-3000012212002033) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](resources--aws_tgw_site--reference--group-001.md#canonical-3032323303231200-2033121210030023-2132230201323200-1222322013031213-2022112313131333-3323121123233101-0001023120220303-2022220113201320) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](resources--aws_tgw_site--reference--group-001.md#canonical-1303300010221133-1230100003101203-1223222012112221-3202211320122133-3311130133323202-3012230302302020-1123303020120231-1100131212321110) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-0003200011233220-3022213033223110-2301231020101322-2133321023312300-3133332003010303-3111002321031312-2022113323110020-2032200330311003) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-2232103300330330-0122322131032012-2201103313200001-2321331212203320-2102101313210200-3111323313322330-0000113222133220-2331030032322012) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](resources--aws_tgw_site--reference--group-001.md#canonical-2130322022133132-2122321122303113-1030333013301221-2020330320233110-0031321322030232-3103222301220103-0321132210031322-3310232000023112) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-0032300331232011-3210330133311013-3320121000121113-2220320222003302-2311303210033011-0330302311231033-2313320001320031-2230101310230320) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](resources--aws_tgw_site--reference--group-001.md#canonical-1310331320212031-3332120112003010-1100222331220031-2132030330330311-0312232323323210-3031302101332111-1021303321021211-3131121300322330) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-2023232211120123-1202211130323331-0100012221323323-1031033132132321-3212002320202132-2231002310321321-1302112323330111-3133220022132033) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-3031331023112000-3331013023021131-3010021030110313-0313010021320131-2133000120112300-3033030312220010-0132000102230031-1222013210322310) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](resources--aws_tgw_site--reference--group-001.md#canonical-3203230123233002-0330200111233121-3310011033220133-2212301211311310-1322203112311031-1130333223100032-3131012222131233-1221210020113331) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](resources--aws_tgw_site--reference--group-001.md#canonical-1032231110133201-2111311220133120-2020200101323000-3300111133210312-2312023303211032-2120322022002011-3023203321030103-3301102112303103) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](resources--aws_tgw_site--reference--group-001.md#canonical-3011300101110232-3113333330101231-3230201332211132-1120132221133111-3012010310301011-1112132331120113-0201333031201132-1011030211322132) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-1110313013320312-0211222212133032-0212203000200113-3130100211330120-3132011322202330-3321020110130133-0331221323220210-0310011210122310) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](resources--aws_tgw_site--reference--group-001.md#canonical-3013100213030221-1031233222320201-1223333331303103-0033320002121223-3320110022330200-3120012203010101-3130300020033301-3030122012000112) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-1122113321021233-3231023030212103-0311112330131103-1111203231301132-3033030021313321-2001132330011333-0331332011313301-0020123121231130) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](resources--aws_tgw_site--reference--group-001.md#canonical-0111121102212333-0213012030310031-3030310222332333-3332300200303000-1113100212010122-1202123331013121-1321213111032012-3331003311210232) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](resources--aws_tgw_site--reference--group-001.md#canonical-0330001111021200-3103321232303210-3331011223323012-1300120130320102-1213002312130202-0013030121002232-2120012033130313-0131122121103022) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](resources--aws_tgw_site--reference--group-001.md#canonical-1011331022322102-3213010103332221-2130231022232013-3022211320330333-1133011200301013-3232230210123202-2111213323033120-0211033131113300) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](resources--aws_tgw_site--reference--group-001.md#canonical-1232102011212333-3331120221333323-0303033233121113-2011012001312002-3303000130223213-3021030313232013-3320201231033030-2303021202332112) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-1310323200030310-1133330202021202-1100131023133332-3020003203323020-3211303000210113-1221203012121332-1212110132010031-0133330001030123) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](resources--aws_tgw_site--reference--group-001.md#canonical-1030323220322200-1222101031101001-1101210132322011-2221103231200131-3130311101322132-2121213012112311-3231200030300203-1313220103131013) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](resources--aws_tgw_site--reference--group-001.md#canonical-1212320211313121-2000132201323233-1333203221233220-1003203212223101-3223110010000032-3110111030322022-2331113201002103-3023102001333020) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-0202210203032333-1011320030010133-3231013110333120-2010202203131303-2201012311021020-0101021301220021-0312020002100113-2022221030022230) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2330102331213101-2132233231211011-1133011110121003-1212332313202320-0100102233212210-2131112012221122-2232210213100033-2221323320303322) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](resources--aws_tgw_site--reference--group-001.md#canonical-1320001221222303-2010012230310233-1210302121012211-0012220101222101-0300122000320231-2021101112011331-0100202230002313-1103003303120311) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](resources--aws_tgw_site--reference--group-001.md#canonical-0201133133111122-2313023300020033-2110120023221213-2212010103221203-3101130033022023-3323222110101132-0211321201221120-0223101021232022) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](resources--aws_tgw_site--reference--group-001.md#canonical-2201103301200022-3230303231120232-1220020111021301-1023031323102101-2120110102111033-2223232220031130-0031001113031312-0320211011212321) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](resources--aws_tgw_site--reference--group-001.md#canonical-1000130131132022-3233322210220011-1103322332231323-1303320000223000-1022130001331122-0001121303233233-1212203130132022-0130300223310110) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-3133113100200201-0002021100330001-3110231111202001-0231002303223302-1223112123332130-2200212313123322-1302320221101220-3020012303322023) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-0330002120102212-0002102311111102-1313203223332120-1113112332001331-1130031321222310-0301331310133022-0123013221232131-2310211321231211) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](resources--aws_tgw_site--reference--group-001.md#canonical-1021010223033300-2001012023101212-1020300323133013-1012303322331033-0321321203033113-2003110022311200-2222002333031212-1222231210021213) |
| `block_all_services` | [block_all_services](resources--aws_tgw_site--reference--group-001.md#canonical-2102122013330111-1010023202001123-2202113001233212-1001103032131210-1101003121133022-1001323233130310-3210010332121103-0232332110003011) |
| `blocked_services` | [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3302233010132223-0120213300333133-0303022322031213-0031121322002022-3032132110113103-1023021232102221-0132300011023001-0112123320221203) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-001.md#canonical-3232323021210320-2030120321131233-2200003200231311-1122202131000103-2232212322013320-0200123033321333-1000122113222002-3310110030110322) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_tgw_site--reference--group-001.md#canonical-1202012320210102-2010111123230120-2210312110322221-1131322313023120-0220102223023011-2112231300121031-0320120011330122-0110230013010111) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_tgw_site--reference--group-001.md#canonical-1311331023131303-3013131202132123-1001203333303330-0133312133021122-0230003012302303-1221003112120102-3200103113322303-1100221323121223) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_tgw_site--reference--group-001.md#canonical-2300120013121203-0200230303333200-1130003323210323-3303110301202213-0232121120311210-2123313102331133-1322000313333232-2023333110333310) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_tgw_site--reference--group-001.md#canonical-0100000011322010-0000112302101201-1033121000101300-2230123232130323-3033123132202010-1221222212321021-0033022102211230-2200202113103033) |
| `coordinates` | [coordinates](resources--aws_tgw_site--reference--group-001.md#canonical-1322130112300230-2201202131210003-3110330212222003-3200131313331101-2031002300011121-0103333022000001-0101232323121033-3133102321032320) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_tgw_site--reference--group-001.md#canonical-2203232223332003-3121331313020231-2202132302103223-3102212232003033-1132213210200031-1130232113120103-1110031111032030-1331031331120302) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_tgw_site--reference--group-001.md#canonical-1203223312231130-1313333013331213-1203200202332120-1130002331110103-0030033320301123-1311332210123030-0132010103230102-1233033030311110) |
| `custom_dns` | [custom_dns](resources--aws_tgw_site--reference--group-001.md#canonical-3301212133232220-2000311102031332-2002222321000123-1123032130321021-1030202011310010-3222130302212312-3301212333220031-1330000221313110) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_tgw_site--reference--group-001.md#canonical-3202031013002113-1220210120210232-2212103110111213-2211022200231223-1212201101222113-3301323311312122-3031311333221313-0322303120033321) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_tgw_site--reference--group-001.md#canonical-1002100023101022-3013312000013012-2213111001322011-3231102000002022-1202021122303312-2120121203032222-0100231320212221-2133221001303032) |
| `default_blocked_services` | [default_blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-2000132300021100-3032231133121301-3320213021031121-3000122030011111-0212331022121323-3002011020221112-0222023233121020-3032212002101311) |
| `description` | [description](resources--aws_tgw_site--reference--group-001.md#canonical-3131303101010321-1301133302311012-0000303112322110-3232100110123112-0121020132202310-1010101131031233-2310003002131033-2020121220203010) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-3122301210323203-0212333020332103-3222231132103033-0120023200212032-1133103322211222-1022132010311032-1113200013022001-2103001020300111) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-0101210101031120-0031320100323200-0003123032213330-2111000120012310-3013021323212201-2113331121220033-1320011211103123-3030003300300023) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-3121232211122213-2201321102313323-3313232122120001-0203232112130021-3332110312311003-1313001210220231-2133312220123231-3212233223132130) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_tgw_site--reference--group-002.md#canonical-2313213201302220-3211201102221122-3230330010001211-2002302330213032-1210133120210103-2313330332120033-2311333011210102-3211303130330200) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-0101213200030233-1002131032330333-3322132312031301-2222130310133211-3021123001303033-0223032300213333-0031103331211302-0033111320332320) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-1302123302001321-3300232121100232-1102223012202033-0322220000222331-0323101112003120-3210001333212202-3310100023221212-3111023100031112) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_tgw_site--reference--group-002.md#canonical-2020113230220132-0320210213313323-2023031331130120-2003110020130301-3320231023031211-3003322212133100-1323132313203101-2100232310331321) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-1220203121310223-3032330333023223-3111211023101030-1020231331311101-3033032121300022-3002112230120120-2011301131131032-1302211102000210) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-1102332011103332-3003102021133331-2023221300213012-1220130210320033-3211010122103223-1330212212021120-0010123001122201-1313202111312220) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_tgw_site--reference--group-002.md#canonical-1230333100130313-3123210233300333-0320133111300223-1210131101221130-2211310203333230-2010002100311202-0112210212223033-1322133330111013) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-0203022322211121-1110020013101112-1111321310121311-3311113203210201-3233322013200211-2113003132031101-2123203202003201-1112231120003200) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_tgw_site--reference--group-002.md#canonical-3311133220111202-1201130302031100-3113022103321013-2011301233030231-1313013000002320-3302022213330223-1010331010121323-0123313021031111) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-3202231200200123-0213221330301310-1113000011111010-3012320231321102-2333300202232111-2133002311220111-3031312110232311-2110112112331202) |
| `disable` | [disable](resources--aws_tgw_site--reference--group-001.md#canonical-3221112231220130-2220123003301123-1132332231000333-3223332232333212-3200211000020332-3030210130300030-0313323201320110-3102223321301333) |
| `id` | [ID](resources--aws_tgw_site--reference--group-001.md#canonical-1120133123333333-1033102312020012-0203332101200233-3211322330103221-2312021311312101-1330203132213321-1130223131210121-0122303120003121) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1033200023221322-2013021020023300-1131032000223030-2302223233233033-1002030300321221-2333213323101132-1303230010021302-1311111233311103) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2321311030030323-3303130110223011-1313111023330230-2312101022101201-0311101131101120-2033221321112030-0131032131210332-0000200303113003) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1200232111122311-3110100123033033-2022302120121002-3131321030211202-1010212032220013-0312220211101223-0221011123231311-0033131231302222) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1123113300100201-1232102030333230-0001112321103330-2332133220230131-3103221113110133-3233222303032031-1322310131322323-1210200331301203) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_tgw_site--reference--group-002.md#canonical-0100032102010030-0211110231131030-3310031320313032-3232302113201301-3311300030323300-2133012231131203-0212231330100231-0033131133102122) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_tgw_site--reference--group-002.md#canonical-2132020032200131-0122312123002001-2133220323330323-3210113101210332-1030131322022210-2100303003010030-3201032310123020-3310211112132233) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_tgw_site--reference--group-002.md#canonical-2332333120231310-3130220320112211-3230010203111202-3233031313002012-2332311100321131-0130312002320023-0313313012301021-2322120333330213) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-2301300120132201-2230011223011021-0113233331330310-2131211201312221-1310021202212130-2031000122012101-0030323213113223-0020120003031212) |
| `labels` | [labels](resources--aws_tgw_site--reference--group-001.md#canonical-2230020112100311-2000012201220330-3201122333302222-1332223120111201-0322330112331210-3102312301222323-3011111101130223-3203313003300023) |
| `log_receiver` | [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-3200330213313000-2030032223302232-2201302110100212-2112331310122100-3222022123013021-3010123110322120-3312001030023232-2103001000200103) |
| `log_receiver.name` | [log_receiver.name](resources--aws_tgw_site--reference--group-002.md#canonical-0000202320001303-3323032101300221-3112021102003223-3023110313013100-2122130223322233-1132331232300323-2211131201212232-2301322230002022) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-3110200222111133-0222012113123120-0320320200220330-1212221000020202-2121301321333221-1323200011020300-1303023130312233-3203312231230000) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-1332012211222331-1202230002030213-1212133223320202-1310121022001131-0330213202212001-0313133221203212-3133100120003233-2032030000302123) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2002101032021020-0020310230203133-2120133233201122-2330201230310123-1230132103102232-3201030321232221-0102033001221130-1101311022321020) |
| `name` | [name](resources--aws_tgw_site--reference--group-001.md#canonical-2130021320201111-3023222110200310-0310032211130121-0313123012133000-1312213120333123-1110131212212123-1233212011331100-3221133231002200) |
| `namespace` | [namespace](resources--aws_tgw_site--reference--group-001.md#canonical-0100323121210313-3033112313322113-1221121123313210-2120203300013102-1220100200230301-2031222210311322-3111001213130213-3230011331330301) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3300322230023203-3003210213133023-2000020032102101-2302030102313121-0313002003300123-1030030132102223-0233311001201132-0322030200021100) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3102311111032021-2102131022103122-0323013130233001-2202323223033230-3033023120201323-3001221330312231-2120012003231112-1102323011130112) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1210333111300321-0312011223003131-1000332133031110-2021202311030122-1012131300311001-0121202101233210-1200301212300000-0222110110212313) |
| `os` | [os](resources--aws_tgw_site--reference--group-002.md#canonical-2022233012203233-1131120110103021-2231122321223033-1110132121110001-0110233111012332-1133131122312010-0111113200211020-0321130302300020) |
| `os.default_os_version` | [os.default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-1110320223313311-0102213232221200-0121123211033220-0220300200110212-3233132233230130-1133230212030031-1211301022233233-0020232011212100) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_tgw_site--reference--group-002.md#canonical-0230031123322331-1223213200113230-2012100100201131-0122120131121010-2010030203111133-1302212222112301-1131003022103310-1203310311011212) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0033131101311310-3332202301032222-1231111021311112-2103103213010223-3221322322311112-2231101023030002-0030233130121020-0013122120101122) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3031131233122022-3021100000201010-0100312021101123-0302201103003120-2222022303001011-0101132111330313-2102020030122230-3201322123332330) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-1133202232121332-2001123221223200-0200201001122002-1123021220313321-3102021220212320-2233032232213203-0311313033221012-0221122210331111) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-2102101130231123-3003330221202203-2100033210001220-1322113203212203-2002332203313302-1113321313213123-0232022331230110-3200112000003111) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3012111321231111-2101101120303013-3112002220233001-0103231321103101-1301102032000303-1121230021300303-0121131013013210-3202020210213011) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-3232113332032210-3033020110001332-3120303332123222-1313201303211311-1323232330320102-2210133300022211-2301131313332112-2031331001121001) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-1302132332003232-2231010300212103-3220223101330203-1030103232103133-2320221032333322-0113023013130103-0211213101221213-3021311201102112) |
| `private_connectivity` | [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0032101031203303-1213220331310310-0100112031003321-2233001013333112-0313130131201302-3200303332021210-2121212012311110-1202302003101320) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-3111321303121122-2312011022310303-3100232113220302-2131113333123132-0011231211122320-0333233213113021-0210213211331231-1220120202231003) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_tgw_site--reference--group-002.md#canonical-0232202221200200-0021331112120002-3210031312330320-2203122013111223-1320220211000013-3030111211001013-3022223310011123-2332320112133330) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-0013230331222100-1112002313210311-0011012202232320-0120012221100220-0303121330011123-1032322012213200-1013233013232111-3110320202202133) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-2300311023332321-0130022203130210-2103210303300010-1210211122232211-1233011212030010-1211030123132122-3132220123321130-3320330233122120) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_tgw_site--reference--group-002.md#canonical-3322222031113011-0022300130320002-0103000223333032-2002132022323331-3301233111322220-1021202102322010-3321201120003223-3300111221311300) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_tgw_site--reference--group-002.md#canonical-3010220230021001-3111020322112312-0012310223231331-1003222101220002-2123202202201312-3230213231033031-3321013301121021-2233222320322231) |
| `sw` | [sw](resources--aws_tgw_site--reference--group-002.md#canonical-1220203211332031-1012022012103003-2202213333310310-0003011022001030-2333301133330210-1321202102321312-3302111220323310-0220233021002132) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-2331321303310332-1300321330320310-0001030011123100-0231302303203300-1210200211212113-3101132200131013-2113203202302100-0120030010221232) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_tgw_site--reference--group-002.md#canonical-1321321203111033-1013033230200311-0123231331120201-3033201212120102-1213213211202313-0212000222010303-0123111021302231-1113023221232321) |
| `tags` | [tags](resources--aws_tgw_site--reference--group-001.md#canonical-1112300100113320-0232112303200010-3113201323211213-3230120011220002-3110231103033023-1333121130310031-0110011122100202-3300031031233211) |
| `tgw_security` | [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1312311100131121-3233120200020121-0002032311133320-3321112003303032-1030013003323100-3023320302103001-2101023301132013-3111100330221013) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1110123011202323-1301012013302030-3302320031320021-3130231311231300-3300033322110302-1120032122233231-0133230032331311-3103112210102221) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0032013332330212-1223121033100231-0211013010300312-3200310103111020-3330330210001323-3130022301111011-2131112133212203-2003323231002333) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-0132022012210012-0302103130320323-0112320303231331-0033212001222100-0032212010213300-0223332112130201-0121203032310123-3002310321102113) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-2323302332212223-3030033003231002-3111111213220311-3323122230102022-3023222010320221-2132001313322021-0111323133321322-1122130323333131) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-1203310231210323-1202132101011333-3103232112031203-1123230031333202-2302211020232013-3111332330030203-1231011322112122-1002220032102320) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0223021110012232-2102232232321313-1300331100001221-2121131131222332-0101133022121232-3133330011203213-0101312232330033-2031021121112033) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0133031213232322-2011100200011223-1013312231310002-1000222322200300-0003301312122011-2310201300213302-0232033102311022-2321113013221021) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-0002331331222220-3001220211102132-3123322023301321-0313002120232033-0131212313122033-0122112131201013-3332330000203003-1103220133003110) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-1320231312112002-3233203000120332-0120113030321030-0033213100013200-2133320232132132-3302313000202220-2031133313211033-0332320321032303) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-3330302211231010-1122033211133123-2313113302203230-3113310221323120-2030332003002113-0010320011300010-1000131022333123-2131221330313230) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0033133103231112-1322011022002333-1101200110203123-2211112332330230-2033111032003301-1032010021000210-3120103020011233-1312333132003201) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0120012202100031-1003213003133310-0010210302111203-2322111200130301-3101221011013332-1330111233102221-1110100230121323-0022103331111122) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-3320012022220330-2301110030300010-1010001121021323-2203031022222212-3032101222133130-1210020310230212-0222030200310211-0010301131321221) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-2001130010322213-1121312320023202-0213213032323233-1222113322301333-2111020000232333-0313311213231333-3033220000002221-3010101312111130) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-2330230101002121-0102021221322121-0313131233330030-2001221111323223-2001021222112020-0323212320312022-0222120303123231-0122011012311101) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1013300123232013-1133231232230201-1233133330301023-2223323200222110-3313312302312022-2032021311033220-1213111133330030-1003213120311330) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--reference--group-003.md#canonical-3322031111021233-2111122223311232-1123330223011130-1113203303012103-0101212220101211-1310231030032023-0030213132100231-3030310320013233) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](resources--aws_tgw_site--reference--group-003.md#canonical-2001010001020323-1231020032133321-1022223202113001-1222220202031303-2331032202300020-3032031333202302-3313111013102201-1013301303130000) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-3231103000203331-1032023231010333-1221201311023113-0232101103212323-2230311131320313-0121222330212100-0222223220231330-1121203023213010) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-0023021330101130-3110202032110331-3321013023003222-3012322110130030-3212033111010211-0032032303213112-3023330021223122-0033030101312023) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-003.md#canonical-2230001032112130-1221022223012331-1201031222002003-3213100210232013-3121010322112012-2203012101131333-1022333201302303-0001223010310023) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--reference--group-003.md#canonical-2333003131113210-2222122020212330-3010313002130002-0013111320230030-0110201233302222-0313233211302130-1300131313312001-0332311230012302) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](resources--aws_tgw_site--reference--group-003.md#canonical-0323023213030011-0012122230003031-3012313112310022-0233333100101312-2100201110023311-3003213002211131-3322021222131210-0123102200133322) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](resources--aws_tgw_site--reference--group-003.md#canonical-1001003212130100-3100301002123233-0301200133020200-3013032223303110-2003331220331303-0221130200210100-0201001200023223-3220230201321302) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](resources--aws_tgw_site--reference--group-003.md#canonical-0332311301232102-1212012333323331-3110001311320100-0022223012133333-1130203331132211-1212322332103010-2003131031010210-2110111302021300) |
| `timeouts` | [timeouts](resources--aws_tgw_site--reference--group-003.md#canonical-1001330203023232-3222112201313300-2201220002003002-2031102013333200-0003203021120232-2121313033133330-0223323010302310-0320102032032131) |
| `timeouts.create` | [timeouts.create](resources--aws_tgw_site--reference--group-003.md#canonical-2232310220122010-0111003333030323-3210311313211222-1321310321130222-0231112021112232-0330032323223330-3031031220300200-3000122012001020) |
| `timeouts.delete` | [timeouts.delete](resources--aws_tgw_site--reference--group-003.md#canonical-3033132130331010-2221120110131321-1113130133312200-1102122301300031-2111313121210311-2320303230120031-3211231121233212-3011320223123311) |
| `timeouts.read` | [timeouts.read](resources--aws_tgw_site--reference--group-003.md#canonical-0312201010021120-2010013212322031-2232301210030230-2030012131322313-3231312020103020-3133220302120113-0030003301220223-3321110002220020) |
| `timeouts.update` | [timeouts.update](resources--aws_tgw_site--reference--group-003.md#canonical-3211321230201002-1001303133330031-0212203131332103-2312020103030121-1320223003301333-3210312221102220-1111113022111301-3330121012111323) |
| `vn_config` | [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0231213112103323-1313232113301233-0102211023023121-0112211110031022-1021302323123233-2322022103231322-2020110003210002-1321013101021322) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-0322123320232130-2232021210332110-2103000301200010-0103232333132211-0211131130313121-3223302030132130-0131213131230001-3321330030210213) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-2210221201311230-0020103212311020-1201232031223033-2012213110222102-2100010302003231-2301320313230130-0033311322002012-2323311222101111) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](resources--aws_tgw_site--reference--group-003.md#canonical-0022311211011332-1101201323233122-3001203011212301-3023121231033102-1003012223301011-3310330233300331-0103122321330100-1110020323320022) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3022123030022230-1111210330310001-3030131321200131-0021233110033031-2000113010213130-2032303033303133-2232312323310201-2331102033131202) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-3022023121132231-0200213231322320-3111110023133120-2113012111320300-3131031232332220-2321230213113310-0211222300132123-0033120132331033) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-1201032212201131-2230120010032220-3310331222230100-1323032023111332-1003102031323213-2312021120231202-2002011130311222-2322032012310223) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-0230221320330033-3303110111330223-0322013311303111-3231022330300230-1033103210111002-2330010123001023-2021032322123203-0233121133223302) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-2231110102000131-3330313330010311-1320020303321130-3132133101103312-0003300210222323-1321131220303331-3302200102221003-0102302010212310) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-1023013121212221-0221013002310132-3010021331211023-0012313200022010-1123310102212122-3132010120123100-3323021203100122-0300211230011030) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_tgw_site--reference--group-003.md#canonical-0031310213331010-2320003023310333-0001303110022113-2021202113310321-0301201320031203-3220202220233102-2302011132322032-2112210002203021) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-0222111103301311-0312013300223131-0001001031201200-0223211220332231-3103003223003231-0200100303210002-0203200320101330-3330011112232022) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-1123103322221201-0211100120100323-3321321031330102-1203332012122133-0102320210211323-2001113322313331-2302223200313212-3211232231123100) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-2030102232313113-0013232300302121-0323300101032333-1221211002230201-3201203220213221-2300120033221023-1003231020323320-1312031132013222) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-3023230202320031-1121030022120030-3023201311311132-1221111102021132-0000011022131321-3201301023123102-1012130213112301-0302011111022201) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0302231131110100-3022103113313103-1000221313201011-0233001012322322-3201101310121333-0132223223110233-3020231002021101-2232103313211103) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-0303221312332333-0323020113123133-3320201111221322-2332011100033031-0302223320021032-0231111003310132-2201101311232033-2130212021031233) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-1122322202131232-3011131231101100-3230010311023232-0003210321102030-2301303312002211-1312211233333321-3133300021221311-3010030123121222) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-0023130323012222-2323331220221033-0011230002310123-2222132120001132-3232122220232310-3210020303303010-1012231120012023-1122211021033301) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0023211323100002-3332202102013120-0303203322221301-2130212010220330-0320321012323323-3101123001321320-1212023202000310-0033121230221100) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-1203012212120023-2112013000220013-3002212012123012-1001222232313212-0223131121111323-1330133032330320-1102110230221030-3220013202120303) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-0200200103332022-2230330011313300-2001322301231020-3311201012012132-3021200020021231-3101112313012211-2303030100213202-2000021131213030) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-3122323220210220-1022320132213100-2031003120021131-2332133013111010-2110103111330313-1110102000000133-3011000023333301-0330230303001232) |
| `vn_config.global_network_list` | [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-2123322011213233-2230010301230121-0221122302013301-1322333211123322-0232321210121111-2032312032112020-3212312201110210-0123212333122030) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-0313322200021003-0111210213133302-0230112130003021-1322201030210320-3301010222222020-3313301302102013-3102121030233202-3011232020122222) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-1011001112231030-1002323311310123-3231030021330013-3321021021303333-3201032232100321-1230013102100111-0031131030013311-0120021001221032) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-3211023022032301-1020022312131313-0211001301330310-1211001111013333-2230002030312200-2230332100032000-3123132230313003-3300222301023100) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-0203003332332300-3222233201333110-3113231210003211-3222211110221222-0133031211300222-0130121223301203-2102313002312330-2033223310112103) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-2201332000210120-2330130010120311-0313013222311230-1303122313003002-0100012012111322-2003133003003032-0111321013213300-0222201133120032) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-1212120011110023-0110110002300133-0120231320010312-3232011230113013-2310302003012023-3123302102122020-2131101023303223-0001211123022002) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-1023020200122112-3230113212012111-0130232200313200-2310102111211232-1212301311130011-3223112101003211-2312121210312223-3323230323110333) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0031233112311030-3010022211030331-1012202101332102-1123230312212320-3201001131012213-2333210020013020-0020033130123330-2121311330103102) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-0331003101300130-2132001111132313-0022012011123213-1210232000133200-3023010030020302-2112333113103333-3230000233303012-2011111002232021) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-3213031111202030-1023031333113110-1000301020303003-1100103010120212-1113011232123301-0302333333133301-3000221301001002-3330031213311121) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-0310121213312121-0030301312211122-1331111123232013-2100213121300311-1030233013022131-1110013232302202-1023333102221323-2000130113210023) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-3221311030211011-3101022211012001-0313003223332012-3001012023313120-3210331333233013-3101030101012320-0032202023111133-3022213232321331) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-2333012130203302-2023133122320133-1121312300121312-1022323101220013-1303010220222101-3312321010212210-1211111020010223-0012302200020132) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-0110010120123213-3312121231010122-0320121130113033-3000023310103322-0010321302220223-1012232011020103-0012101000233132-0002210110330201) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--reference--group-003.md#canonical-1113022302212132-2230120030303030-0210201103303131-0323122021022220-2322022211011130-2333123201211210-2310132201210132-2311323030313200) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-003.md#canonical-1201003121233210-3212202011021322-1013322223230221-2113120231021100-3210200320231001-2311201320312010-0222111033302032-2012123101021302) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0122110313201220-0223220301211310-1223321130332331-1211033012120212-2210301123221012-1322020233311302-1122131322113201-2133113120220220) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-003.md#canonical-3010100223133102-3001321222130321-2222002003212022-1101322312011221-2233203202030121-0133120012121222-2111323202300123-3113231231321011) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--reference--group-003.md#canonical-3030000113003331-3222000213101101-0200103021212133-1211101021031212-2222122221130200-3303033000320123-2313211323132203-2130113203332102) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--reference--group-003.md#canonical-1000213032310221-3202221331002231-1130032200332033-0012003333310311-0310320232220310-2103113133331321-2313002223123230-2021132001223121) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-0200022330112330-2223321001212030-3230203122312211-3201031103201021-1231200203310203-3320033130201313-2130220000221013-1113130210232213) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-3031102320111211-2330322010113103-3322302302113000-2011322333031323-1131101023030212-2300012322102301-0332100131020133-2312311303003301) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--reference--group-003.md#canonical-3221011100310021-2203333030011200-0312303303222022-1232101311203002-1011012100031020-2220313011212003-1302232003212301-1230230213002032) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-1233313132101102-0202022011132222-3320231111210133-2300221123301233-3220102201132032-3221121310200013-2130212330303202-2013330112020231) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-2300123323101133-0333233121212133-1323121113332300-0200012221230221-2132203130021110-1131112322211333-1322223333330132-1312300113203211) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-3322212033133201-1300120303002122-3331102013210213-2332102123212213-0223012011211222-2032331202213303-0312331020303102-1201230201312113) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--reference--group-003.md#canonical-2111002213230113-1020003211321131-2221221222303322-3330032322322121-1312211311223023-2123033301330323-1020322210123211-1210010222331331) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-1212131333032231-1121011031221110-3322011330203330-1330103321201032-1231203022030110-1220111102221123-0012132303100323-3223230301213221) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--reference--group-003.md#canonical-0203101322303211-1003112120331221-0000120110012032-0301011030221211-1222233102200311-2232330131100110-0110212032110312-3331320011111111) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-1023223131320110-3212102023100121-0122222113001321-0213311110012111-0101011000033303-2300310133222303-0011321011233213-0302113212031103) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--reference--group-004.md#canonical-3022002020030030-3223123333233231-0211003221310213-1320310113012102-1230020100132023-3110222000223300-0222301223230320-2220300132202021) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-1012312122322110-3131320000323011-0310211130210030-0233303323323212-1313122333231303-2200021000011202-1320201132211202-0302300130101022) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--reference--group-004.md#canonical-2231132120223200-2231102322020202-2022121320201212-2301302333012123-2121213022233211-2232033222200330-0101131233110022-3333200001122130) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--reference--group-003.md#canonical-1121332020110101-3001211000302033-2213211133203212-0200033203120133-2122010301121310-2020330021123123-3113022333120102-1010330222113103) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-3303000132001100-3220230222001113-3131003032113202-2032120200312223-2103031100133223-0320203032011222-0023212102301010-2211202123031011) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-0230122321233130-3113323101122330-1300201300113231-2221103020302302-1311320033131020-1022220033113303-2033201310321010-2121320033002131) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--reference--group-004.md#canonical-3322121301321211-3112333220233323-2303333033333120-1032322211112311-1210220300301222-0030120223213222-1323203320301301-0010323011022321) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-1330331223311222-0110100210213332-0132321013203101-0132002211333031-0202201021201023-2020121321223130-3011110232210132-2232123231033330) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-3120301021120021-0031000220112232-3130132120311122-2032222031220200-3133100232201231-3231322322331111-1023130212111221-0021223101001003) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--reference--group-004.md#canonical-3110021210020211-2003030310021312-1211213011312133-0202220030323210-2132200122100001-3000211233111012-3010231130321321-2120223220022301) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-1333013233113210-0101211333300001-2033031213212002-0131311230302331-0300312122302122-2223222112020110-2032103022022012-1011000333210310) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3321302013300302-0033030333200030-1101012320303102-2113301131010021-0032022331220031-1211121001203202-3213310232333132-3130111112321233) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](resources--aws_tgw_site--reference--group-004.md#canonical-2113010033121112-2303223233130301-3313211123220002-0002111010030112-3121030222220210-0021003120111210-1133030210223123-3000320312113202) |
| `vn_config.no_global_network` | [vn_config.no_global_network](resources--aws_tgw_site--reference--group-004.md#canonical-2221201312100210-0212332020333032-0301021012030333-1112120112312021-1332121002322313-1301312030213202-0321232322310130-0023132332110303) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-3203131130312231-2202032330220011-0212221102102311-0031200311002233-2032211232210211-1321133312003003-1331112322132223-2111211310002331) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-3002313302203313-0222330331130322-1322211332322222-0003010030200020-3232312220101331-0021113203201123-3021200203112210-3111221033321222) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2232131130132012-1110233200102200-1023231320300232-0102211112120011-2203120022123122-0100112130222200-2013320312201221-1302013232132330) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0211131030323112-0203000102023232-3021021030112021-1301323311222102-2232032113032012-1222001211303303-1131202022123103-3011103131213023) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3021211111133303-2232303013132321-1033221111103233-3300000111212030-3232002222011003-0121100300201210-1303330122212330-2302010303331010) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--reference--group-004.md#canonical-0003111212031130-3012322110322011-0022100122001201-0333331231032330-1332322223333021-3202213331001030-1213103031202031-1010322221233313) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-004.md#canonical-1332222301001012-3002002103331232-2022023010200300-3332311000301121-2000211010102231-3223322022211312-2020000311131021-3300100301222210) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-3321022100210320-1212022111321120-0230301231312311-0321200032023332-2322200100232033-0003233003212131-3000221212301130-2333322010112131) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-004.md#canonical-0121332302020121-2331131002030011-1303021001222310-0003221212303213-3211032122100200-2212003230000111-2212200030012310-0020120021130111) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--reference--group-004.md#canonical-3221002323313332-0123332021033331-1031030121231132-3212013301031102-2022100020231333-2332033001000201-3023023000030312-2100220031212033) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--reference--group-004.md#canonical-0030013322112311-2312010033031031-3033023303011122-0201001101332133-2213310312011132-3201210031020222-1212300310300211-0233010220110132) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--reference--group-004.md#canonical-0330022121012303-1030113301212300-0310322132332030-3313201123000222-0220000310011233-3010211221011333-2121033302032333-3230021133223032) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--reference--group-004.md#canonical-1002201013231301-1130112132333120-1032203321003211-0123320211111111-0003311010212020-0030103102220223-0031233301330332-1233013323103101) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--reference--group-004.md#canonical-0323201233133020-3310131130331131-2131113223113303-0111121121312302-2200232301220112-1101131320133330-0013312033113321-2023111023033232) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-3211232202201232-1031300221322332-3332100103100122-2313323303000230-3321130303320310-0121203121030022-1301200213222033-0120003101331032) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-2010303032211013-1130002131031332-0110231002123323-1121003031220200-0120130301132312-1220212303220102-1133330300211303-2312201313101110) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-0122332203100202-1323121013131003-1120331030123202-1033001023032020-1210212323133000-1031313230332330-3221130333020102-2303003332132000) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--reference--group-004.md#canonical-0132030300320113-0110132132331220-2022302221133112-3303312013130320-2332123312301202-0302212311213000-2303003230011013-3110233211133113) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-3103021333233110-1011210131322000-2123321202132102-1033132312123001-1332332020200203-3113112232100312-2202200210302311-0130210312320211) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--reference--group-004.md#canonical-3000222210112033-0313011023210122-2101122210233303-1102323223021113-3010010101022002-2302113201110111-3130020112301230-2020301312032030) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-0022102320330031-1301233021001322-0133021130031201-1313122222013012-1133113200001310-2320302220302023-3312133032223203-0233122221012212) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--reference--group-004.md#canonical-3012032120033230-1002320131101230-0302310110031322-0222123111033310-1130011333121033-2331211001320323-2021223223010332-1200133300013111) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-3110230121213301-2031211111333120-1223001300021321-3210231131001300-2222113113003332-2221101013021333-3132203201131120-2001231132120122) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--reference--group-004.md#canonical-3122323132320201-3302232002130000-0110111200113013-1022010232320320-2003321102313023-2131003022312220-0033023222122033-1221100021023201) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--reference--group-004.md#canonical-1122023323030231-0202321033332110-0233110120110231-3300111310130000-0312312010130310-0203123120222111-0323303021121100-2113212203322203) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-3033333132110230-0101321032301313-2200333332220220-3330322331100021-3221102322011032-2113022010001113-2220330301031031-3232030221323301) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-0323231030102312-2031232301232010-0220222200330232-0231223122131211-3213023310322223-0313313013310030-2301132112303002-1120303112133022) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--reference--group-004.md#canonical-1232010211121032-0232021230333032-1022300220202133-3323230101203020-3011111101320201-0201313213303000-1333310333033223-0310100330212221) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-3103022313302202-0002120301011210-3010302030010020-1012013123113010-0123300332121311-2300003102102120-3320212010212200-0013310312101123) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-1023232322123303-2120112113232100-0103223303310332-1313030210220223-0032302220033330-2330203030211331-3333100021200202-3112121010123032) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--reference--group-004.md#canonical-0132122202103201-0211030311300202-1132011203322003-2220101232103030-0110020121313133-2210002100223302-1203103120023210-0300123021322323) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-3111311323101023-3311301333110110-2330200200323033-2120232102212013-3333301203210010-0223001013321213-2111032000323001-0320313323013221) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3220220133332031-3332312230201232-1022333330022231-0022233011233103-0123030111003201-1120322231111012-2332113123233033-0200231110322021) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](resources--aws_tgw_site--reference--group-004.md#canonical-1202210310030101-1012103013201211-1002001200202022-3330223100321212-1113102330110213-1102313323310013-3212020313210310-2211230323011322) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](resources--aws_tgw_site--reference--group-004.md#canonical-3001222103030002-2113002333332213-3101301320231220-2101231001300030-3330022330113332-2130013113113300-2122332310131203-3002112231320021) |
| `vpc_attachments` | [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-2330320031031130-0223131233013102-0031030323103312-1323110303231202-0032131311202231-2300322232210321-2211310022213121-2021200310311132) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-1210321101230210-2200013012202212-0211221230322001-2121112201021310-3211233213301020-2002023033112131-3133321132020220-0222221111321220) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--reference--group-004.md#canonical-3210011200110111-1123303222023031-3121103133003212-3301201031101331-2302321202100113-1102100130033303-2110013002113232-1233213302200202) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](resources--aws_tgw_site--reference--group-004.md#canonical-3200001003301011-3223122010122023-1300221302312301-2031311121211021-1112132231112213-0300212231022323-3302203123022121-1010223002332223) |
| `waf_signatures` | [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-3222213203010330-0133312022232213-1321000200323331-1122003320120230-1011033320302333-3322120000321332-0231130203100112-2321011322233231) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_tgw_site--reference--group-004.md#canonical-0313330010000201-2032202213233101-2201101031123001-1023232213222320-2013230313222013-2101323121320132-3031123323200121-2030232013223213) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_tgw_site--reference--group-004.md#canonical-0012101302001113-2310030110220330-1112010323013232-0102111022123101-1032033001110320-1222021331023312-3302323013130322-0211021220212203) |

<a id="canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- aws_parameters

<a id="canonical-2320313303213230-0001313321120112-3320121311232221-0023233120021303-0201303023333210-0111222203013200-2232103003000221-2113211211012122"></a>

Type: `"object"`. single nested block, Optional.

Setup AWS services VPC, transit gateway and site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "az_nodes",
    "instance_type",
    "ssh_key"),
  validators.ConflictingObjectAttributes("custom_security_group",
    "f5xc_security_group"),
  validators.ConflictingObjectAttributes("disable_encryption",
    "enable_encryption"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip"),
  validators.ConflictingObjectAttributes("existing_tgw",
    "new_tgw"),
  validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "nodes_per_az"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "total_nodes"),
  validators.ConflictingObjectAttributes("nodes_per_az",
    "total_nodes"),
  validators.ConflictingObjectAttributes("reserved_tgw_cidr",
    "tgw_cidr")}
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
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

Terraform syntax:

```terraform
aws_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133110331021000-2333033222312321-1133311120212001-1032301212310212-1020211120322333-2220113313022023-0201133022210331-3232031223113002"></a>

### Direct properties for `aws_parameters`

- [admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-2020332001100331-1200322312312112-1111030131330011-1232231201300330-2111132210322012-3131233231123113-0310232032230311-0112212210101102): complete subsection reference.

- [aws_cred](resources--aws_tgw_site--reference--group-001.md#canonical-3221321200323212-3031211300100213-2003322103022002-2131331222303222-0033301223120033-1001200231213320-3111011333213013-0202201101013200): complete subsection reference.

<a id="canonical-3312123123233201-0122110220031001-1203201230012111-2101033213131023-1033113331300200-1300210130302311-2210010012211010-0000133212111121"></a>

<a id="canonical-1112330010112201-2002031323213023-3103102001011200-0131121103121001-3120212233020102-3331110213203120-0301223031013112-2110010330032020"></a>

#### `aws_parameters.aws_region` property

Type: `"string"`. Optional.

AWS Region of your services VPC, where F5XC site will be deployed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323): complete subsection reference.

- [custom_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-2202312333030012-2012020030300223-2200032120320021-2300233330130103-2112121312321113-3001103331323221-2210221023122203-1123333023131111): complete subsection reference.

- [disable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-1233103122222130-1233112230333200-1023322310130032-0221033101122032-3122020310332121-2120020301022301-2030030101022101-3332111110111101): complete subsection reference.

- [disable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-2110031003230100-2312131002310323-1001220020132002-0323031212331203-1300101103030020-0231313100020130-3301130312131130-3233311021111033): complete subsection reference.

<a id="canonical-2130322022133132-2122321122303113-1030333013301221-2020330320233110-0031321322030232-3103222301220103-0321132210031322-3310232000023112"></a>

<a id="canonical-0301331232033011-0000311332011311-2002321003111113-2210213103202213-2213223233022031-3112210211112300-1322303122221202-0302323203100103"></a>

#### `aws_parameters.disk_size` property

Type: `"number"`. Optional.

Node disk size for all node in the F5XC site. Unit is GiB.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-2133303310323133-0210232321223011-2220133231121310-2230230312210302-1301312210121223-3011313310300231-3011001123132333-1121203312013021): complete subsection reference.

- [enable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-0000333323231203-1023210121132232-1303000013322013-1033201232230022-1013222120023203-1003132210132021-0220203321221332-3110132333033313): complete subsection reference.

- [existing_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0103122020332202-0102330312022121-1102331320110123-1220223221303201-1131000122312112-0020333333323033-1100131200000221-0230320110302120): complete subsection reference.

- [f5xc_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-2303103032232331-1022112310033202-2002020123001231-3312132123222012-0010203331222313-1011322213112012-1213231333203120-1013033213033200): complete subsection reference.

<a id="canonical-3013100213030221-1031233222320201-1223333331303103-0033320002121223-3320110022330200-3120012203010101-3130300020033301-3030122012000112"></a>

<a id="canonical-3313112312311313-3210121021033003-1223011213233332-2020212222012210-3210223010322232-2013100202333132-2312133331301321-2223103011302111"></a>

#### `aws_parameters.instance_type` property

Type: `"string"`. Optional.

AWS Instance Type for Node. Instance size based on the performance.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102): complete subsection reference.

- [new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322): complete subsection reference.

- [no_worker_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2221202120101010-2021301212311032-1333222232212320-0013032102113003-0302100203333220-1333213231210112-0133010203022110-3321330223121101): complete subsection reference.

<a id="canonical-1320001221222303-2010012230310233-1210302121012211-0012220101222101-0300122000320231-2021101112011331-0100202230002313-1103003303120311"></a>

<a id="canonical-3322230101322010-1022313302221113-2212303210203033-3103120031203323-3111132200313102-2022310121110112-2331221002123002-3211102323333120"></a>

#### `aws_parameters.nodes_per_az` property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](resources--aws_tgw_site--reference--group-001.md#canonical-3121213320330121-3012221332212132-3211100230303323-0221132120130330-0212223211102331-2031212313332200-1211230111210012-3023322021210301): complete subsection reference.

<a id="canonical-2201103301200022-3230303231120232-1220020111021301-1023031323102101-2120110102111033-2223232220031130-0031001113031312-0320211011212321"></a>

<a id="canonical-1111202022332000-3123312212130320-0220002001223111-3211021001023113-2331210220230112-0312022032003331-2133002101300211-2013110213310010"></a>

#### `aws_parameters.ssh_key` property

Type: `"string"`. Optional.

Public SSH key for accessing nodes of the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](resources--aws_tgw_site--reference--group-001.md#canonical-3131221233111313-0201223031103232-3231330222212320-3200313300213211-3332131122121033-2101311133231111-0111222033131212-1030022202322120): complete subsection reference.

<a id="canonical-0330002120102212-0002102311111102-1313203223332120-1113112332001331-1130031321222310-0301331310133022-0123013221232131-2310211321231211"></a>

<a id="canonical-2003220203112010-0122231203103200-0200003021030222-1002010122301101-3013111030123023-1202302032101303-1220121203323221-1210030102023203"></a>

#### `aws_parameters.total_nodes` property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="canonical-1021010223033300-2001012023101212-1020300323133013-1012303322331033-0321321203033113-2003110022311200-2222002333031212-1222231210021213"></a>

<a id="canonical-2022221122113223-0212012120232300-1302122320123113-3033021002121221-1123103033130033-0023021131002111-1133322233030231-0002133103331033"></a>

#### `aws_parameters.vpc_id` property

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Existing VPC ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2020332001100331-1200322312312112-1111030131330011-1232231201300330-2111132210322012-3131233231123113-0310232032230311-0112212210101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.admin_password` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.admin_password

<a id="canonical-1013333212210132-2113020132003113-3202320333103133-0202000022220133-3231131331233301-2030030311231131-2122302220130202-1222231220012113"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003003211312322-0120030131233102-2002122201031231-0330110020333222-0132323021332122-1230230322223132-3013000120102022-2003202220311111"></a>

### Direct properties for `aws_parameters.admin_password`

- [blindfold_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-2022320003322221-3111020123130023-0102313210031210-2032301122130033-3311202300313330-2211031321330333-2300011121223303-2131212232111032): complete subsection reference.

- [clear_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-3033013030123300-1003021000123010-1213231311212112-2101033313031212-3231222220221221-2000201311032233-0132313003000021-1100233310212310): complete subsection reference.

<a id="canonical-2022320003322221-3111020123130023-0102313210031210-2032301122130033-3311202300313330-2211031321330333-2300011121223303-2131212232111032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.admin_password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-2020332001100331-1200322312312112-1111030131330011-1232231201300330-2111132210322012-3131233231123113-0310232032230311-0112212210101102)
- aws_parameters.admin_password.blindfold_secret_info

<a id="canonical-1022312320122023-2303103323112111-3332331301012102-3322102000221012-0313032120200113-3211103132333301-3113321231212311-2001233002202113"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333322201312301-3202211300331332-0133230023110231-1330321030311332-3121211212221223-3130332313103320-3232321211331131-3331320100311323"></a>

### Direct properties for `aws_parameters.admin_password.blindfold_secret_info`

<a id="canonical-0300121232211101-0030011201000131-3110000020032100-3332203303312312-0333211022030110-1220011230232021-0010323320323001-0030023300221333"></a>

#### `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2100020320203101-2313033213323100-3200011320023303-1101132231222021-2102311200333123-1123122110322132-3130203212121023-1112211331110000"></a>

<a id="canonical-2203223223331321-3311221333330203-0303013211221131-1300302002102301-3231312023330120-3022333000323130-3001121202023133-2333220103210210"></a>

#### `aws_parameters.admin_password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3200100313323332-3133032132032201-0301330221221332-0333333012030203-3010320032313320-3000101010320311-2103333213220022-3211200213103110"></a>

<a id="canonical-2030101321033202-3011101312311231-2221030303001303-0003200333330113-2222213103220311-1001202222021330-1210322323121223-3000130322203020"></a>

#### `aws_parameters.admin_password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3033013030123300-1003021000123010-1213231311212112-2101033313031212-3231222220221221-2000201311032233-0132313003000021-1100233310212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.admin_password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-2020332001100331-1200322312312112-1111030131330011-1232231201300330-2111132210322012-3131233231123113-0310232032230311-0112212210101102)
- aws_parameters.admin_password.clear_secret_info

<a id="canonical-2111222133222033-2011002113331213-2200100102213331-0212333033320232-2232003113111212-3301003312021202-1012131221111232-2213000332030202"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311110111101021-3330023102022003-1332312132231112-2103002131030301-2321102203113322-3200022303002230-3201100302310103-0102300222201330"></a>

### Direct properties for `aws_parameters.admin_password.clear_secret_info`

<a id="canonical-2321200220223022-2222212023120301-1202301123210223-2113223231032021-3112212330112320-0213203322023231-2131223022333330-2013201113033101"></a>

#### `aws_parameters.admin_password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301213321132320-3323303330301102-2030200020231311-1233031302201123-2300313102333032-0013032232333230-3321220121300000-2011022231220333"></a>

<a id="canonical-3233222123322333-0003033301002032-2313111132100121-3033323111323230-2030203310031201-2121103333110130-2201121210122012-1111100022312123"></a>

#### `aws_parameters.admin_password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3221321200323212-3031211300100213-2003322103022002-2131331222303222-0033301223120033-1001200231213320-3111011333213013-0202201101013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.aws_cred` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.aws_cred

<a id="canonical-2233233113100032-0201022203031130-0122222132002002-3213331333100130-3023312203302013-3333111301123023-0023223202023032-0023302303222031"></a>

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213220000032201-2031021223100002-0030113103000302-0312112020012201-3120010331311201-2121103313331330-1002100123320122-0000222233123302"></a>

### Direct properties for `aws_parameters.aws_cred`

<a id="canonical-2222002012302011-1122023323113011-2100122012230123-2201131200003031-1132320213321230-3311030201121102-0330101200320301-2022230321213012"></a>

#### `aws_parameters.aws_cred.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0310123300232330-0023232120232001-3212313113312203-3330100131302330-3103022010220121-2313033212221321-3012121112011020-2221133220210131"></a>

<a id="canonical-1320312231132311-1111312322120123-2123122302112213-3211303000000100-2200133311011010-3233132233312022-1322131031203213-3010132101120313"></a>

#### `aws_parameters.aws_cred.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3231333223323012-2211202102121002-1210332031001313-1303132133221310-0000323323203221-0222301110301211-2333033023310203-1321132030232101"></a>

<a id="canonical-0120313122333213-2213212203130213-0203121231103311-0212210311301202-3231030100321132-2312202222100232-1020210221131223-1120031031331100"></a>

#### `aws_parameters.aws_cred.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.az_nodes

<a id="canonical-3323022333231001-3130003120022323-2010303312000303-1231022001333302-1000130113313000-0003001101331032-0021123000223011-3223022112323311"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name"),
  validators.ConflictingListObjectAttributes("inside_subnet",
    "reserved_inside_subnet")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032030300330232-1211223311322231-3002130312000223-1132303220220013-0211023133031302-0212111120012331-2312212200020303-0303021121220023"></a>

### Direct properties for `aws_parameters.az_nodes`

<a id="canonical-1012030312111120-1030033021223300-2320231023100003-2132210111322123-1020002030201133-3000302200111001-2303022223122213-2102320313033222"></a>

#### `aws_parameters.az_nodes.aws_az_name` property

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-1121212130231203-0012123310011212-3233020203313300-1121323022303121-0231122031013021-3011332303233301-2103213101213101-2112001210321231): complete subsection reference.

- [outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-0113001211031102-3102103332211013-3221230222101122-1322113203000301-1320331302211332-3023313203312201-3102121331112020-1033102223203322): complete subsection reference.

- [reserved_inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-0211031222203202-1222213003223302-2032321201023112-1330000102013312-3001203200131322-3012021210301312-2221022210333212-0323112203110011): complete subsection reference.

- [workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2132211212002212-0302322022233330-1130233101011230-1202020012230320-3032003133223222-3203100030103022-0312210120011222-3101312020210003): complete subsection reference.

<a id="canonical-1121212130231203-0012123310011212-3233020203313300-1121323022303121-0231122031013021-3011332303233301-2103213101213101-2112001210321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.inside_subnet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- aws_parameters.az_nodes.inside_subnet

<a id="canonical-2321010031302032-1310031213133333-2020133002023220-0111020112013002-1033101111120332-3230030023231030-2201203322203012-2123002100200011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Additional upstream details:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102012121121220-0312210030231333-1311023330100122-0010222022122302-1122223002301312-3033001323020221-3112312332023333-2000223023020323"></a>

### Direct properties for `aws_parameters.az_nodes.inside_subnet`

<a id="canonical-1031321212232331-0030201301102203-0313321120203311-1110100323302233-2131303220103030-2130031322132231-0001031113322002-2311001010323131"></a>

#### `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-0212000320232300-0311222123313132-1320220103213130-1133122300233200-1321211111021310-2131223333130302-0203132033113202-3122213010132131): complete subsection reference.

<a id="canonical-0212000320232300-0311222123313132-1320220103213130-1133122300233200-1321211111021310-2131223333130302-0203132033113202-3122213010132131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.inside_subnet.subnet_param` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-1121212130231203-0012123310011212-3233020203313300-1121323022303121-0231122031013021-3011332303233301-2103213101213101-2112001210321231)
- aws_parameters.az_nodes.inside_subnet.subnet_param

<a id="canonical-2113313100022212-2222103332300313-1302213332222023-1231032232011313-2323110002122022-2112022310203123-3212002002022220-3110310131110032"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100211101322110-1332230110300011-3301002311001301-2031300203132002-3011322200030223-1123130220232111-2030010010020121-2130231231231231"></a>

### Direct properties for `aws_parameters.az_nodes.inside_subnet.subnet_param`

<a id="canonical-1301123000130032-0102101120213233-3011312201202200-2031321230020212-1110233031133112-2121002123320201-0111301023333021-1121203210201030"></a>

#### `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` property

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0113001211031102-3102103332211013-3221230222101122-1322113203000301-1320331302211332-3023313203312201-3102121331112020-1033102223203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.outside_subnet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- aws_parameters.az_nodes.outside_subnet

<a id="canonical-0010211200331010-0210310223003110-2310130030100100-2030333020022022-1011301300120103-0032332003302020-0100033101301220-2323303323212000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

Additional upstream details:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010001312121020-2120031020013213-3033200332110330-1330131120111210-3230132123033031-0003010212221310-0312321210123320-0013101033110120"></a>

### Direct properties for `aws_parameters.az_nodes.outside_subnet`

<a id="canonical-2300311121033220-0031011300302113-1123011322233222-2022021222200101-3101102112033211-1330211212122220-1101020023112320-2223302212312021"></a>

#### `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-2022030222003201-0131202133333200-3033302003121300-2231311200313022-3032332302123211-0120111311001132-3002223233131212-3310302212131303): complete subsection reference.

<a id="canonical-2022030222003201-0131202133333200-3033302003121300-2231311200313022-3032332302123211-0120111311001132-3002223233131212-3310302212131303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.outside_subnet.subnet_param` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-0113001211031102-3102103332211013-3221230222101122-1322113203000301-1320331302211332-3023313203312201-3102121331112020-1033102223203322)
- aws_parameters.az_nodes.outside_subnet.subnet_param

<a id="canonical-0113203113012330-3321013201201131-1310202333301132-1101030121010331-3203300122000202-3003220111330023-1303232022021112-0032233310301211"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110232213103232-1130100313100022-1220302123121133-1221031112220033-0211133211011303-1332220213200122-3210301120310211-2003012000022013"></a>

### Direct properties for `aws_parameters.az_nodes.outside_subnet.subnet_param`

<a id="canonical-0111311133303133-0022332020223211-1101301230321002-2111300022021312-2000123120322003-3102023131212300-3023131311101102-3102310100002110"></a>

#### `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` property

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0211031222203202-1222213003223302-2032321201023112-1330000102013312-3001203200131322-3012021210301312-2221022210333212-0323112203110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.reserved_inside_subnet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- aws_parameters.az_nodes.reserved_inside_subnet

<a id="canonical-2210332001131001-2100203001320031-3312321113332121-1321222330300203-3223100201312332-3203130132211020-0333030212330330-0131121220013233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved inside subnet.

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
reserved_inside_subnet = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132211212002212-0302322022233330-1130233101011230-1202020012230320-3032003133223222-3203100030103022-0312210120011222-3101312020210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.workload_subnet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- aws_parameters.az_nodes.workload_subnet

<a id="canonical-3221221122222330-0131322203023032-2032122131312330-3231321301333113-1003103002131212-3201312012320032-3001212331322101-3202222102001000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for workload subnet.

Additional upstream details:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
workload_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013333001002012-1233300113223021-1232232011303033-3130320122202133-1000121331330202-0231300000000202-2022331311121230-3100231132310020"></a>

### Direct properties for `aws_parameters.az_nodes.workload_subnet`

<a id="canonical-1013121112232300-2312201321123322-3020202212032213-2320303123101001-3002233313033110-3010113110123212-0322011233303323-2020010232033310"></a>

#### `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-0211322022321213-2202213023101323-1302003100321211-0131222011220220-0101122000120021-3320311231322032-1021330300031332-0123031103110200): complete subsection reference.

<a id="canonical-0211322022321213-2202213023101323-1302003100321211-0131222011220220-0101122000120021-3320311231322032-1021330300031332-0123031103110200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.az_nodes.workload_subnet.subnet_param` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323)
- [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2132211212002212-0302322022233330-1130233101011230-1202020012230320-3032003133223222-3203100030103022-0312210120011222-3101312020210003)
- aws_parameters.az_nodes.workload_subnet.subnet_param

<a id="canonical-0023011230113121-3123233010303310-1130021023100002-2230120322303120-1230211120333223-2332330200222101-1010021320220312-2311030002123210"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323121110212330-2232331210222220-1331023133000010-1210021231223323-0102131203321000-0112030202122221-0231111113021032-2013013101221321"></a>

### Direct properties for `aws_parameters.az_nodes.workload_subnet.subnet_param`

<a id="canonical-0121030331313110-0121232222020130-1230322120310323-3030103212222233-2003011031312200-2201221210130130-1330002223100232-3031203111113302"></a>

#### `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` property

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-2202312333030012-2012020030300223-2200032120320021-2300233330130103-2112121312321113-3001103331323221-2210221023122203-1123333023131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.custom_security_group` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.custom_security_group

<a id="canonical-3100211220330203-3212203213222111-1113200212122020-0333000030223231-2332102013230202-1322311003201321-0322221110231020-3000012212002033"></a>

Type: `"object"`. single nested block, Optional.

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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
custom_security_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033323333123010-3312332230231031-2203113313121322-3221321013213132-2010322021012100-2201230031003232-3330102113311120-2301232123113012"></a>

### Direct properties for `aws_parameters.custom_security_group`

<a id="canonical-3032323303231200-2033121210030023-2132230201323200-1222322013031213-2022112313131333-3323121123233101-0001023120220303-2022220113201320"></a>

#### `aws_parameters.custom_security_group.inside_security_group_id` property

Type: `"string"`. Optional.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-1303300010221133-1230100003101203-1223222012112221-3202211320122133-3311130133323202-3012230302302020-1123303020120231-1100131212321110"></a>

<a id="canonical-2133132332320021-2132203102230223-2012212133321232-1332322013000202-1010310131020322-3010202012300302-1302102110233231-1033110202120221"></a>

#### `aws_parameters.custom_security_group.outside_security_group_id` property

Type: `"string"`. Optional.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-1233103122222130-1233112230333200-1023322310130032-0221033101122032-3122020310332121-2120020301022301-2030030101022101-3332111110111101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.disable_encryption` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.disable_encryption

<a id="canonical-0003200011233220-3022213033223110-2301231020101322-2133321023312300-3133332003010303-3111002321031312-2022113323110020-2032200330311003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable encryption.

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
disable_encryption = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110031003230100-2312131002310323-1001220020132002-0323031212331203-1300101103030020-0231313100020130-3301130312131130-3233311021111033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.disable_internet_vip

<a id="canonical-2232103300330330-0122322131032012-2201103313200001-2321331212203320-2102101313210200-3111323313322330-0000113222133220-2331030032322012"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133303310323133-0210232321223011-2220133231121310-2230230312210302-1301312210121223-3011313310300231-3011001123132333-1121203312013021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.enable_encryption` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.enable_encryption

<a id="canonical-0032300331232011-3210330133311013-3320121000121113-2220320222003302-2311303210033011-0330302311231033-2313320001320031-2230101310230320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Additional upstream details:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110133310010201-0222123130120121-2003121312220201-1120021230032011-3303231301020021-3330211123111311-2011303003132103-1323322231221033"></a>

### Direct properties for `aws_parameters.enable_encryption`

<a id="canonical-1310331320212031-3332120112003010-1100222331220031-2132030330330311-0312232323323210-3031302101332111-1021303321021211-3131121300322330"></a>

#### `aws_parameters.enable_encryption.kms_key_id` property

Type: `"string"`. Optional.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0000333323231203-1023210121132232-1303000013322013-1033201232230022-1013222120023203-1003132210132021-0220203321221332-3110132333033313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.enable_internet_vip

<a id="canonical-2023232211120123-1202211130323331-0100012221323323-1031033132132321-3212002320202132-2231002310321321-1302112323330111-3133220022132033"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103122020332202-0102330312022121-1102331320110123-1220223221303201-1131000122312112-0020333333323033-1100131200000221-0230320110302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.existing_tgw` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.existing_tgw

<a id="canonical-3031331023112000-3331013023021131-3010021030110313-0313010021320131-2133000120112300-3033030312220010-0132000102230031-1222013210322310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing tgw.

Additional upstream details:

Information needed for existing TGW.

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
existing_tgw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333032332332102-2213201202223113-1013311320321211-3132333231203002-2322321130302033-3230322330011101-2302130211201121-0302302102133021"></a>

### Direct properties for `aws_parameters.existing_tgw`

<a id="canonical-3203230123233002-0330200111233121-3310011033220133-2212301211311310-1322203112311031-1130333223100032-3131012222131233-1221210020113331"></a>

#### `aws_parameters.existing_tgw.tgw_asn` property

Type: `"number"`. Optional.

Enter TGW ASN. TGW ASN.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1032231110133201-2111311220133120-2020200101323000-3300111133210312-2312023303211032-2120322022002011-3023203321030103-3301102112303103"></a>

<a id="canonical-2220332020131122-2110312112131000-0101333101222203-0011112333033232-2002011333230202-2212230213003302-1201322111030321-0331010131302031"></a>

#### `aws_parameters.existing_tgw.tgw_id` property

Type: `"string"`. Optional.

Existing TGW ID. Existing TGW ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3011300101110232-3113333330101231-3230201332211132-1120132221133111-3012010310301011-1112132331120113-0201333031201132-1011030211322132"></a>

<a id="canonical-1321232333333112-3001010101200312-1303330101323100-3123133101122031-3021311312300212-3322001003311232-0232231020210123-2233020133133210"></a>

#### `aws_parameters.existing_tgw.volterra_site_asn` property

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2303103032232331-1022112310033202-2002020123001231-3312132123222012-0010203331222313-1011322213112012-1213231333203120-1013033213033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.f5xc_security_group` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.f5xc_security_group

<a id="canonical-1110313013320312-0211222212133032-0212203000200113-3130100211330120-3132011322202330-3321020110130133-0331221323220210-0310011210122310"></a>

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
f5xc_security_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.new_tgw` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.new_tgw

<a id="canonical-1122113321021233-3231023030212103-0311112330131103-1111203231301132-3033030021313321-2001132330011333-0331332011313301-0020123121231130"></a>

Type: `"object"`. single nested block, Optional.

TGWParamsType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("system_generated",
    "user_assigned")}
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
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

Terraform syntax:

```terraform
new_tgw {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032221201332311-0101023333132033-1202131300023313-1132023112223332-0001102333120003-1211013130312102-1330011131301211-0210010002130103"></a>

### Direct properties for `aws_parameters.new_tgw`

- [system_generated](resources--aws_tgw_site--reference--group-001.md#canonical-0103221113020220-1023032300110200-1122320003021100-3301223110103000-1232122010031200-2123003310013320-0101220201312303-1332330322113022): complete subsection reference.

- [user_assigned](resources--aws_tgw_site--reference--group-001.md#canonical-1012321201321212-3333210011321220-1300013121113023-2010222022103120-1331133131000231-2010332022011311-0221120300323130-1130322133020332): complete subsection reference.

<a id="canonical-0103221113020220-1023032300110200-1122320003021100-3301223110103000-1232122010031200-2123003310013320-0101220201312303-1332330322113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.new_tgw.system_generated` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- aws_parameters.new_tgw.system_generated

<a id="canonical-0111121102212333-0213012030310031-3030310222332333-3332300200303000-1113100212010122-1202123331013121-1321213111032012-3331003311210232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system generated.

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
system_generated = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012321201321212-3333210011321220-1300013121113023-2010222022103120-1331133131000231-2010332022011311-0221120300323130-1130322133020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.new_tgw.user_assigned` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- aws_parameters.new_tgw.user_assigned

<a id="canonical-0330001111021200-3103321232303210-3331011223323012-1300120130320102-1213002312130202-0013030121002232-2120012033130313-0131122121103022"></a>

Type: `"object"`. single nested block, Optional.

Information needed when ASNs are assigned by the user.

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
user_assigned {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211323002213233-2021122023223313-1333013331102100-0323333032011132-3321013231320010-3313122133232232-2230011013311123-0203013103031201"></a>

### Direct properties for `aws_parameters.new_tgw.user_assigned`

<a id="canonical-1011331022322102-3213010103332221-2130231022232013-3022211320330333-1133011200301013-3232230210123202-2111213323033120-0211033131113300"></a>

#### `aws_parameters.new_tgw.user_assigned.tgw_asn` property

Type: `"number"`. Optional.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(64513, 65534),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="canonical-1232102011212333-3331120221333323-0303033233121113-2011012001312002-3303000130223213-3021030313232013-3320201231033030-2303021202332112"></a>

<a id="canonical-3302023211031203-1202103230002200-3031112303222122-3313101331212203-0333020002020002-0231101320033331-3313102301313302-3100222202031330"></a>

#### `aws_parameters.new_tgw.user_assigned.volterra_site_asn` property

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.new_vpc` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.new_vpc

<a id="canonical-1310323200030310-1133330202021202-1100131023133332-3020003203323020-3211303000210113-1221203012121332-1212110132010031-0133330001030123"></a>

Type: `"object"`. single nested block, Optional.

AWS VPC Parameters. Parameters to create new AWS VPC.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name_tag")}
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
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

Terraform syntax:

```terraform
new_vpc {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230120221121102-0311200001031212-3123101033013220-3013023131111203-2022303231230203-1330332323202211-2021110103201112-0002013322210111"></a>

### Direct properties for `aws_parameters.new_vpc`

- [autogenerate](resources--aws_tgw_site--reference--group-001.md#canonical-1202212102332103-2130320110010101-3030031132133003-0203112230312220-1031032300311121-2022201222313202-3111232110333201-1332220130120323): complete subsection reference.

<a id="canonical-1212320211313121-2000132201323233-1333203221233220-1003203212223101-3223110010000032-3110111030322022-2331113201002103-3023102001333020"></a>

<a id="canonical-3311133220201132-1011110110230122-1002120232332331-2303310123123201-2321301013323132-3211000232111230-0022113003120023-0212123123222303"></a>

#### `aws_parameters.new_vpc.name_tag` property

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VPC Name.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0202210203032333-1011320030010133-3231013110333120-2010202203131303-2201012311021020-0101021301220021-0312020002100113-2022221030022230"></a>

<a id="canonical-2220030030012112-2013012112001201-1012231030011123-2332012001311320-1303132112200220-1231030123122003-3010120223030030-0123131003203323"></a>

#### `aws_parameters.new_vpc.primary_ipv4` property

Type: `"string"`. Optional.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-1202212102332103-2130320110010101-3030031132133003-0203112230312220-1031032300311121-2022201222313202-3111232110333201-1332220130120323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.new_vpc.autogenerate` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322)
- aws_parameters.new_vpc.autogenerate

<a id="canonical-1030323220322200-1222101031101001-1101210132322011-2221103231200131-3130311101322132-2121213012112311-3231200030300203-1313220103131013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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
autogenerate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221202120101010-2021301212311032-1333222232212320-0013032102113003-0302100203333220-1333213231210112-0133010203022110-3321330223121101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.no_worker_nodes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.no_worker_nodes

<a id="canonical-2330102331213101-2132233231211011-1133011110121003-1212332313202320-0100102233212210-2131112012221122-2232210213100033-2221323320303322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no worker nodes.

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
no_worker_nodes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121213320330121-3012221332212132-3211100230303323-0221132120130330-0212223211102331-2031212313332200-1211230111210012-3023322021210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.reserved_tgw_cidr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.reserved_tgw_cidr

<a id="canonical-0201133133111122-2313023300020033-2110120023221213-2212010103221203-3101130033022023-3323222110101132-0211321201221120-0223101021232022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved tgw cidr.

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
reserved_tgw_cidr = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131221233111313-0201223031103232-3231330222212320-3200313300213211-3332131122121033-2101311133231111-0111222033131212-1030022202322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_parameters.tgw_cidr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.tgw_cidr

<a id="canonical-1000130131132022-3233322210220011-1103322332231323-1303320000223000-1022130001331122-0001121303233233-1212203130132022-0130300223310110"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
tgw_cidr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220110012320321-3101303103031133-0003122021330302-1031031032333101-1113322233110230-2103023110110211-1311102011202130-2313012323111002"></a>

### Direct properties for `aws_parameters.tgw_cidr`

<a id="canonical-3133113100200201-0002021100330001-3110231111202001-0231002303223302-1223112123332130-2200212313123322-1302320221101220-3020012303322023"></a>

#### `aws_parameters.tgw_cidr.ipv4` property

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3303203233101111-1231211011103232-0133220301212023-0131322023033222-3032123223330012-1020121030201121-2102031230301211-3211000022222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `block_all_services` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- block_all_services

<a id="canonical-2102122013330111-1010023202001123-2202113001233212-1001103032131210-1101003121133022-1001323233130310-3210010332121103-0232332110003011"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](resources--aws_tgw_site--reference--group-001.md#canonical-2102122013330111-1010023202001123-2202113001233212-1001103032131210-1101003121133022-1001323233130310-3210010332121103-0232332110003011)
- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3302233010132223-0120213300333133-0303022322031213-0031121322002022-3032132110113103-1023021232102221-0132300011023001-0112123320221203)
- [default_blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-2000132300021100-3032231133121301-3320213021031121-3000122030011111-0212331022121323-3002011020221112-0222023233121020-3032212002101311)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- blocked_services

<a id="canonical-3302233010132223-0120213300333133-0303022322031213-0031121322002022-3032132110113103-1023021232102221-0132300011023001-0112123320221203"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102230201223002-2133000100311032-1202302311220121-1313223333211332-1133211102123100-3031013131223203-1020102322020110-1013000001310213"></a>

### Direct properties for `blocked_services`

- [blocked_service](resources--aws_tgw_site--reference--group-001.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310): complete subsection reference.

<a id="canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- blocked_services.blocked_service

<a id="canonical-3232323021210320-2030120321131233-2200003200231311-1122202131000103-2232212322013320-0200123033321333-1000122113222002-3310110030110322"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303231220132022-2020313312313232-1120020210002323-0301222133122122-2003031111111312-0102322113200300-2030123323000312-1030020300013320"></a>

### Direct properties for `blocked_services.blocked_service`

- [DNS](resources--aws_tgw_site--reference--group-001.md#canonical-3020320313222311-1001121100203321-0103302101312103-1110011032310223-1311003322332201-1311110320103121-0221331030321101-3303203312111111): complete subsection reference.

<a id="canonical-1311331023131303-3013131202132123-1001203333303330-0133312133021122-0230003012302303-1221003112120102-3200103113322303-1100221323121223"></a>

<a id="canonical-3303100302131330-1023103221203212-0123111132331321-0233312310123230-3120002300210230-3010013222230310-1333201102332311-3310111130311011"></a>

#### `blocked_services.blocked_service.network_type` property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](resources--aws_tgw_site--reference--group-001.md#canonical-1213232301302012-0033110021321203-0001210010031131-0213213120033203-2020321313002311-0230123221101221-0133132203122332-1121203001001130): complete subsection reference.

- [web_user_interface](resources--aws_tgw_site--reference--group-001.md#canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310): complete subsection reference.

<a id="canonical-3020320313222311-1001121100203321-0103302101312103-1110011032310223-1311003322332201-1311110320103121-0221331030321101-3303203312111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.dns` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-001.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.DNS

<a id="canonical-1202012320210102-2010111123230120-2210312110322221-1131322313023120-0220102223023011-2112231300121031-0320120011330122-0110230013010111"></a>

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
dns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213232301302012-0033110021321203-0001210010031131-0213213120033203-2020321313002311-0230123221101221-0133132203122332-1121203001001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.ssh` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-001.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.SSH

<a id="canonical-2300120013121203-0200230303333200-1130003323210323-3303110301202213-0232121120311210-2123313102331133-1322000313333232-2023333110333310"></a>

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
ssh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.web_user_interface` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-001.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-001.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0100000011322010-0000112302101201-1033121000101300-2230123232130323-3033123132202010-1221222212321021-0033022102211230-2200202113103033"></a>

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
web_user_interface = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330100032212111-2023231232130113-2033223012210111-0102030223332333-1003211333301322-2022000120201201-1202112101331102-3200002321100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `coordinates` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- coordinates

<a id="canonical-1322130112300230-2201202131210003-3110330212222003-3200131313331101-2031002300011121-0103333022000001-0101232323121033-3133102321032320"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212210012033210-0133203223202020-0301112213310031-0022023030300122-3230102100011102-3112331133201133-1302101312030322-3202231301032213"></a>

### Direct properties for `coordinates`

<a id="canonical-2203232223332003-3121331313020231-2202132302103223-3102212232003033-1132213210200031-1130232113120103-1110031111032030-1331031331120302"></a>

#### `coordinates.latitude` property

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-1203223312231130-1313333013331213-1203200202332120-1130002331110103-0030033320301123-1311332210123030-0132010103230102-1233033030311110"></a>

<a id="canonical-1011323303132013-3321222230023221-3020002020122221-1223011202320302-3031021303233122-2320013322220001-0221232321110002-2100230322220111"></a>

#### `coordinates.longitude` property

Type: `"number"`. Optional.

Longitude. Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-0310130113103331-1321231230201233-0112322012222330-2020103011100211-0311022022233022-2032112020200200-0003021332313331-1212231013022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_dns` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- custom_dns

<a id="canonical-3301212133232220-2000311102031332-2002222321000123-1123032130321021-1030202011310010-3222130302212312-3301212333220031-1330000221313110"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002031113203111-2031022232001123-2101222010011123-1023313002103231-1031213321330130-0313001132133321-1220220223121223-3321011223333121"></a>

### Direct properties for `custom_dns`

<a id="canonical-3202031013002113-1220210120210232-2212103110111213-2211022200231223-1212201101222113-3301323311312122-3031311333221313-0322303120033321"></a>

#### `custom_dns.inside_nameserver` property

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1002100023101022-3013312000013012-2213111001322011-3231102000002022-1202021122303312-2120121203032222-0100231320212221-2133221001303032"></a>

<a id="canonical-2213320010133232-1010111302302101-0131212232110101-3211221212320113-0013332322110222-0320003301030030-3311112122222020-0100122221033020"></a>

#### `custom_dns.outside_nameserver` property

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0232210132032302-2311321302133012-3210002210020112-0120013313103120-1000032132332201-1332320121032013-1330221230222303-1221113233030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_blocked_services` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- default_blocked_services

<a id="canonical-2000132300021100-3032231133121301-3320213021031121-3000122030011111-0212331022121323-3002011020221112-0222023233121020-3032212002101311"></a>

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
default_blocked_services = {}
```

This is an empty object or choice marker. It has no direct properties.
