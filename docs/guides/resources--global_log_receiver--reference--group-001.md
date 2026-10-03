---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022211320020320-3023013020211011-0210332121032032-2231123232221101-0032003030021211-3113230220333223-0123031200331110-0203323300230032"></a>

## Property reference — Property reference / 320133321330 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- Property reference

<a id="canonical-2220031323223323-3003103200121011-0013112330303120-2310012013110221-1112131101120311-1312103132232302-3330311013233033-3321211223223213"></a>

## Direct properties — Property reference / 320133321330 / 3

<a id="canonical-2323300230200323-1032211212100112-1310213300330132-1211013321012233-1032000321313023-3221102011321301-2220312132123311-3010232111030010"></a>

<a id="canonical-3333101000301031-1020332010032103-3233210013032111-1311112301122111-3120010122231133-0223030201020013-3220230323301022-0201210123122123"></a>

## annotations property — Property reference / 320133321330 / 4

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

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-3231131323212011-3030323112222132-1230222222202231-3212031220330031-2231302201012313-3103320303103003-3102113033100033-1132112103231012): complete subsection reference.

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333): complete subsection reference.

- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312): complete subsection reference.

- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031): complete subsection reference.

- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133): complete subsection reference.

<a id="canonical-0101300011230302-3212212211213211-2133301122323202-3321331000012021-0031210130320122-1203122131301100-1132303212130220-0331103333001201"></a>

<a id="canonical-3031322302312002-1012321210003012-2313303021203123-1220311303030212-3212313211020030-0001102101210313-0030111132100333-0010221213032330"></a>

## description property — Property reference / 320133321330 / 5

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

<a id="canonical-3133030313301311-3311210003123002-3022300020323121-1123223312011012-0013110111211011-1032110123000001-0233333331113303-1202120133211020"></a>

<a id="canonical-3333110103213111-3110331112233113-0023323121030301-0020211102200013-2211302322223030-1313132322321000-0230102012032213-0103322311133330"></a>

## disable property — Property reference / 320133321330 / 6

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

- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-3211233331121300-1322331131001220-0102011110001201-1030212021001001-2311133020330123-0200201221122221-1011112033100211-0230123330102313): complete subsection reference.

- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031): complete subsection reference.

- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333): complete subsection reference.

<a id="canonical-0331121310220213-0022123302322203-3101123131332131-0020121312302123-2133211030310101-2013203031012120-1102101332200211-1221121031131200"></a>

<a id="canonical-1202201023011132-2022202320301012-3112030331112211-1002331121221333-0222012320332120-1101130322022022-3031211002102121-3020331310313302"></a>

## ID property — Property reference / 320133321330 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112): complete subsection reference.

<a id="canonical-2020223323023222-1233310223030001-1221012210003030-3333003331013000-0100331002102000-2102133111203331-2021001311001213-3103032303013212"></a>

<a id="canonical-2003222311123020-2101032101011021-0301322133331011-1102323123012231-3003230033130222-0302310311212222-2010130001002103-1123132311311231"></a>

## labels property — Property reference / 320133321330 / 8

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

<a id="canonical-2301112001212110-2010010232130022-2320111323321011-0332022032331220-1323002223213031-3111020303110012-2113003212313312-2102311101111302"></a>

<a id="canonical-3102122332022210-1131211311133100-2121112220202320-2102230122112023-1303123303300003-0131330100120201-1133100312000320-2333031021130303"></a>

## name property — Property reference / 320133321330 / 9

Type: `"string"`. Required.

Name of the Global Log Receiver. Must be unique within the namespace.

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

<a id="canonical-3312003003120321-1321311211320200-1021210321323030-2220000231310203-2001122302332012-3313023310113003-1011220223320110-0210131230313203"></a>

<a id="canonical-2230122030330300-2032230001120212-3310220310202000-0230000012212021-1032210122230113-3322001102322132-2023211310212302-3310002202133301"></a>

## namespace property — Property reference / 320133321330 / 10

Type: `"string"`. Required.

Namespace where the Global Log Receiver is created.

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

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020): complete subsection reference.

- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-2110223021133332-2001211101313323-3012032332230023-3123201031102303-3033233222010310-0322300301310121-1032103033320210-2132202001210101): complete subsection reference.

- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-1131102221232000-0300222020002322-0302133211001210-2101203323112122-0311121123102201-1122110013312213-1303121113303221-1111202320322233): complete subsection reference.

- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-1323013202223030-2213130333303312-1001111102133131-1001220102221312-3002003012031130-3113012231221100-0102230132130332-0202133321133203): complete subsection reference.

- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332): complete subsection reference.

- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211): complete subsection reference.

- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122): complete subsection reference.

- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-3211022021302113-1000032210211020-3113000022100301-1332120310011110-2113003230211310-2211212131301002-0201111232112322-1122230010302122): complete subsection reference.

- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130): complete subsection reference.

- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130): complete subsection reference.

- [timeouts](resources--global_log_receiver--reference--group-005.md#canonical-2300030330132202-1012220213322303-1300301203131302-1312202110301333-1300333322313010-2203010102023132-0210122130033122-1033121302122132): complete subsection reference.

<a id="canonical-2011302010130300-3311022102123332-3230010332020323-2323023030301222-1101121101111231-3322020233230012-2130201312320131-1323133111220102"></a>

## All schema paths — Property reference / 320133321330 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--global_log_receiver--reference--group-001.md#canonical-2323300230200323-1032211212100112-1310213300330132-1211013321012233-1032000321313023-3221102011321301-2220312132123311-3010232111030010) |
| `audit_logs` | [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-0101233320301302-1213113233000311-2013000202300333-3020021232010212-0210010120233212-2021333333322130-0002331232212331-0023310101122033) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0010031130120012-1201320000231232-2331132333100100-0310012020113322-1120212131013223-3021312000023122-2333123112111130-0212122212030102) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-2232132021012300-3021231320112302-1130100210132221-0221100333103023-0232032321131310-2213230330320101-2131031030313323-1330203213300321) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](resources--global_log_receiver--reference--group-001.md#canonical-0321000013211132-3032323300011313-0032000232003201-2012220332313312-1122100112223130-2000012321312322-3100012301311330-0231002013302013) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](resources--global_log_receiver--reference--group-001.md#canonical-3123303032220202-1213021102311131-0022210003020222-0101213232232232-0130222202010310-1001233101212303-3301121132203230-0202312231212033) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](resources--global_log_receiver--reference--group-001.md#canonical-0001102000000220-0030203110132301-3130332103133221-2312030331130003-3000312231023132-0201330203223113-0022310130203233-0111013233011223) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](resources--global_log_receiver--reference--group-001.md#canonical-2013030030212011-0333122012030213-1031323111100201-2113221120013322-1003021033010003-1031301201222300-1002001133221020-3031213112100100) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-2203331202332200-3323021210002333-0330112200301333-2332303021101123-0012203200203203-2323032210202322-0222100330212003-1030222230222131) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-001.md#canonical-2332020313332001-1030112322013022-1021321300120211-3233121321011200-2312030030003122-2313002312302131-1213111211100011-1000213112013321) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-1033112113301211-1032213012201102-0211333120003113-2210200111121320-0320310132100302-3223133133213301-3310001102011210-0002233220212213) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](resources--global_log_receiver--reference--group-001.md#canonical-0023232300211033-2110230220102223-0301131012003002-1120313031203021-3201000220303011-3230323222101010-1323301022331201-2112100202113230) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-2202100312230333-2010121313112133-3112033310033320-0333133230233132-3111222320210103-0310130102020103-0003322222222102-1213001022121013) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-001.md#canonical-3222303132331023-0102203220312111-1220221123021133-1310101322322100-1322200232132132-3112102131101011-3302302122222201-0300321320310213) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-0212020112113131-0332213230302223-1213201300303333-0033112222022021-1210300222003233-3322020130003023-1032001223301103-2312333312213033) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-2231303203110303-0102302202220020-2001130321101003-2121001032030132-1210220300120202-2222003210201023-1001003101202133-3201020122122010) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](resources--global_log_receiver--reference--group-001.md#canonical-0300231032011203-2031222213220233-3023213132332012-0100000330023221-3323003303031302-2203132021020122-1012031111233121-3333313303212313) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-3230230023231122-2012223223233330-3022301121020102-2322320320022222-2321123001012203-3033030311211231-3022311230123023-3003111330222033) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](resources--global_log_receiver--reference--group-001.md#canonical-2231221310103120-3330211120232021-0010232200301210-3331033330123021-0301331032210103-2101113311122223-0010133132210320-0203003002031211) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](resources--global_log_receiver--reference--group-001.md#canonical-3322300301331022-0100200310011211-0011033212322032-2233312332001100-3120232201102013-1220110122233111-0200030010220322-2302330000203212) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](resources--global_log_receiver--reference--group-001.md#canonical-3303221103330231-3211220200132020-0320302121222312-0310332013122003-0203330110222310-1212211301110232-3231113221233201-0100013313230201) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-2320200132011030-1231332200110211-0002331211012320-2213001200201030-1011331311001011-1022302212331332-0103310102303100-2031212300120002) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-0303101323120132-0101132321103131-1021223210313011-1331001302303103-3300311021201320-3311000032213030-3233103323001111-1211023231111323) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-2001122110210031-1313230203003331-3223300000033231-1231121233011310-0203112231102020-3032011110203012-2231023012331032-0213010123330031) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-001.md#canonical-2231102311232100-3132200132012302-3332000101001130-3002021022130133-0200002003333102-1333211310203010-2232033122323123-1312023321030103) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](resources--global_log_receiver--reference--group-001.md#canonical-0220331113010130-3022003111122202-2130111233102323-0313230321013031-1030231113211301-3013112032103303-0221113020311023-0332323133310211) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-001.md#canonical-1030212301213033-0012021230303111-2201100231332123-3013331231211012-3233220003200001-1313021331221211-3100202301203110-3100311201232213) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-1133030010012111-0011322330102322-3003002322030020-1012030002321123-1023221122103000-2103212120320220-0103303013211310-2020230010031020) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-001.md#canonical-3313010020223102-3223111210323322-2020230233330112-0201121332132302-0010022132021332-2323210311123312-2313330200200120-2223211323101020) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](resources--global_log_receiver--reference--group-001.md#canonical-0313031303323300-1301312310300113-2302200332321100-0011211311302031-1112023023312322-1132320323203103-0202321133032212-1332231321221303) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](resources--global_log_receiver--reference--group-001.md#canonical-3030111010023323-2101332000131212-0000222320202313-3320301123031300-3320133003332330-2112330230210133-0112331221321220-0022020210002022) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](resources--global_log_receiver--reference--group-001.md#canonical-2012113103313300-3310113010310322-2333213210112300-2000233330302210-3221022001030331-2220110121020230-1012112022332032-1023012012032022) |
| `azure_receiver` | [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3013200231333020-1303012020320320-1001100130303121-3211010010010303-2320110321121200-1102123310122230-2100321311330302-1302111303311231) |
| `azure_receiver.batch` | [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0012020231111121-0203112023302002-3201223223301221-0110122201023320-0033311132200101-3121223302013232-1322223022210222-0003231110333030) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-002.md#canonical-3233300230003111-1002232321120212-1120131320330230-2230011200000220-0310011031311213-1032213212031012-3032301302222211-1033011200102003) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2201121010212221-3333130012202230-0310323003321221-3302103331010130-3001122022013112-1231031202112300-2300223100212021-0023330003303031) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](resources--global_log_receiver--reference--group-002.md#canonical-2023133201123030-2200212013213212-3122333300012111-2220100300033212-0031303300101121-3112302011102000-2312023330311320-0213010330111331) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2331113323131313-2212320102112212-2321332232232321-0011112310321133-2003233220210202-3312301323330221-1230002211112211-0222200111013310) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-002.md#canonical-1133200331220331-2322000033231301-2102112113223033-2122100231002113-3333303112202123-2212003103122112-3231213312012110-1020120211321321) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-3010231230212120-3233010321201322-3122222301100211-3332231100012323-3032031103303113-1220231231032311-1030333112211301-2210133100010011) |
| `azure_receiver.compression` | [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1333001023030232-3033021230020012-2200212023020012-1123203330001131-0311311201310213-2113311230120332-3223302210130103-2113032301112023) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2003030013011233-1011012032210321-1321013022131323-3302031212113022-2120311232021332-2230110333211000-2330031110223031-2302033233123002) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-1223021002132132-1111323330133130-0332023101121001-0031210213211222-2133100211013312-0120222320302030-3231203303212120-3113003032130130) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-0322011332110320-3020030301102330-0300111213313323-1020300133023301-0032102330303013-3302131221110333-2102130313301233-0131310100110012) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3310012233002333-0122120112212320-2011103110201122-3222023032132113-1220113331113220-3331120211330113-1010121013010132-2012031002301020) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2221332330013210-2212303002000303-3033121011303301-1030333213220221-2100000312303010-1003332212223322-1302031311322332-2312133133321220) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-1332113000102211-1030032212232202-3300011232333200-0130320230323323-3031023020221231-0012031222222101-0331030200231213-1010020213100013) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-0201102231220210-3230032021131022-3330011203002113-1323322331300101-0310122203112311-0033210302001121-1301030032121201-2331301330031031) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-3202200210213230-2110103323333303-1000121213312333-0111310303030101-1220221212211023-0301120322000213-2001300023233111-3233211313322320) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-1120321123300223-0331213203033300-3111301311121132-2022213223112030-1330231331001130-0131023233102013-3111201200010021-0130103023213220) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-3232212030133111-0201122001203111-2111331210031313-0320102000130232-3003002301231302-0011121031203110-1312203112132302-0320103301311212) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-2123113022133202-3203332330103031-0030232122232132-3303330012123310-1232102122001310-2233313111101102-2303211332022101-2201233322101303) |
| `azure_receiver.container_name` | [azure_receiver.container_name](resources--global_log_receiver--reference--group-001.md#canonical-2303103313121301-2313301303211130-3302300023031203-3212330330121000-1220101333330100-3331012323131121-3302233312132020-1302211023033332) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-0311023201301101-1132312022231232-2103120011132102-0102011221233030-2112000210012131-3100333013323222-1231123302313012-2032013220322111) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-002.md#canonical-2221130031202331-3220332121302101-2101010001123320-2231331201310210-2123033003333300-3211222002110222-2210233313303303-3023123223312330) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-3302233113221300-0012021231120122-3130301002220321-0310331301321303-2113232110123023-3313222301031023-3201212111032130-2122033121002020) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-3001030000033221-2033211000100320-1211012312212131-0023313313201022-0222112323203011-3223331011200320-0100120221321310-1002231001332121) |
| `datadog_receiver` | [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-0003033200133213-2021130011102033-2321233212202033-0010010332230103-2101001230301213-0313001301111103-3112321322103202-2021022012211332) |
| `datadog_receiver.batch` | [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-1102323220133001-3102020103031023-1210120111332011-2323100010331012-2221123031201102-0030023002100213-0003223000300332-0233103101122020) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-002.md#canonical-3311123322330313-3203103233122202-0111200331123310-0000122202122003-0322221030210001-2122121311220232-0303321330022213-3132020330330201) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-0111213032122002-0212231001232223-2311101322323113-0333031230023033-2130300030102110-1213002130212313-1221312130003332-1322223102223202) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](resources--global_log_receiver--reference--group-002.md#canonical-2101230131331303-3133020102030322-0112001230122010-3210033211212112-3021002212121000-1321132010010112-2023321231130001-3133111201133001) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2000031012002030-2103300321130003-0320323122013100-3233102330022232-2232203011122111-0003123212201030-0211222001012103-3211312131332030) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-002.md#canonical-1220230221130321-1122130233032220-3033322103020133-1001100103230232-2320332113302010-1000010321213133-0111030112303123-1200013203023223) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-1222112302032112-2323321012111203-0333311321103031-0120123201103310-3113220231311303-2201331013212030-3330131032032113-3221301032331210) |
| `datadog_receiver.compression` | [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-0100020323133330-2021232133002231-1033032121301031-2323111023313303-3312120101223112-0101032021001221-3231210211100323-2102303333030003) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-3200131330223011-2220311302123323-1310310200203210-1202121210122333-2202123113301122-0311002200102130-1032323132010021-2222331132203000) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-1012031012132023-1230031221212030-0022331120022302-2312301031323102-2332121203010200-0233210203320123-1123000230230201-2022023322012323) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-1101132011322000-2131323330200310-1221031313131223-1023312231010022-0223232111332313-1112112233301010-0122100211220311-0313300220312013) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-0311110001003200-1333323330303023-0213211030202113-3313133213111110-2321301113233211-3023202232300333-2210202333032301-0230233131223232) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3212121320233333-2011211200232111-3201111031232012-0323123300322202-2121300022211201-2131331210310113-3223210022202301-3021102311320203) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-2110220013220200-2313232110030021-3131110210130320-0333213010102133-0310030213320013-3331323211301003-2312321231123201-3123202131020320) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-3020131131212121-3221322033002313-0022311120231303-1223221133331101-2332021301332202-2331210031332311-0320021223320031-0200012020311320) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-2333021100030001-1112322312113220-0222110000313022-3012022220020203-3330120011111310-3333301100123323-2010212203113101-0002120122201013) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0212212011332201-2112331202222210-1000231021300131-0303021231130032-1201111020201302-0210013003232011-2121210330133000-2303320213131033) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-2212222301022122-1022000120020000-0131221223313001-0133122110201310-0302113233012032-1100322033001113-2102330113232033-3330031010310232) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-3210031323132101-3300021330003110-3300330110031113-3320003322211223-3210111230332033-2133211313002023-1031202301103121-0000213111331222) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](resources--global_log_receiver--reference--group-002.md#canonical-2121132022131301-1021220221122333-3102223320122020-2032212120002232-1301010303210303-2100300311231021-3031223321200231-1202100221321011) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](resources--global_log_receiver--reference--group-002.md#canonical-2323030200323112-3112102121002321-2003201100321203-2101112010003303-2120203103130300-2210110123231110-0211301022233112-1002102330021000) |
| `datadog_receiver.site` | [datadog_receiver.site](resources--global_log_receiver--reference--group-002.md#canonical-3232330011012121-3222103200321203-3211122001332130-0113313111103100-1020321222033022-0312113313210232-3230303121002033-3110311123303200) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-3220330103021032-0210311031103113-3030233033030230-1223010313131033-1210301320222101-0003300030031001-3031012332221313-3202210033130111) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-3113220300212223-2130021103010123-3201010130113032-0232222212220312-3012200023230331-0022302323000320-3023310310022032-0030313333003310) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-3231303032021303-2223300101331120-2113322302333323-2222331300323313-1321022203030120-2320203122202133-3003103003003021-0220202102101003) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-2031312120033231-1233001020332033-2310103133202101-2111102232201003-3222323301012101-0110000213010323-1313022322232203-2303120200321003) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-1330321111112203-1023303310212100-0322310022233212-0323002213203312-1303120331131033-1230310331020001-2120132322003300-3322112300031200) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2030320302211130-0232233012123113-2302113030332003-2211303123301032-1312130301011011-3211203022310330-2003012220113221-3210122033022133) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-1211123010030223-3110100232033321-2101231123131001-0320231210220130-2112033121130200-2132301303120200-0233121002301310-0302111111101112) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-002.md#canonical-3231113030112213-2330031211212303-2001113222102323-3313020332233332-1000123010220203-0001021303111332-3132030231001010-1331002001023111) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-0003011213120310-3322300103012000-1033230100230302-2312302122110102-3131222202110123-1010200330333233-3331101332331010-3110230120323000) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0020311202202033-1333123302213210-0123210123123121-2033120213200203-1333203123012332-2013222130010220-0323001111113313-1202030020000303) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-0330121030032012-2210321121010120-3331302222132201-2230020212113103-1003101123211031-1030122232330103-3201103230210020-0313100230110110) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-1230332202301332-3311210220113211-1230331110202003-3020120121220313-0103023131113223-1333220130223033-2100221322212320-0223201203111233) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-2313310000211230-1320300112032003-1311032323132230-3032202233323101-3032020021311322-0202210213011333-3303110102312313-3301021203001331) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3301013130203001-2002233222321222-3121333300113233-0222011310330120-1111230333122223-1221123210222133-2111122313202332-2211130112110021) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-3112303302301022-2200121233221131-0102100031113030-1113312322113220-1123300211002332-0003100332213113-1113221103022302-1002000123121331) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-3201230100332312-2330123312313133-3213022213112301-2001331201320320-3123213311213211-1120320331330001-2203110311012032-1213302331230032) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-002.md#canonical-1233332031103301-0122321320010033-2321000001023003-2210203232012321-2331102211302211-1320113213032022-3303311102020312-3300200211013201) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-002.md#canonical-3000201022320103-1233121233010210-3123222302301310-0000003213030331-2232230210030111-1230110221113100-3201102031202033-2301202322222320) |
| `description` | [description](resources--global_log_receiver--reference--group-001.md#canonical-0101300011230302-3212212211213211-2133301122323202-3321331000012021-0031210130320122-1203122131301100-1132303212130220-0331103333001201) |
| `disable` | [disable](resources--global_log_receiver--reference--group-001.md#canonical-3133030313301311-3311210003123002-3022300020323121-1123223312011012-0013110111211011-1032110123000001-0233333331113303-1202120133211020) |
| `dns_logs` | [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-2112120022003322-1202120213333133-0010302200210023-3232021030212312-0232123211210022-2203312323223132-3100211030221032-3112323313223031) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-1221011012130302-2001213102221021-3130230103112313-1211200120113032-0230233232020101-3331301323120221-0220020210111122-1233101312331111) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2011111023320021-3023300220300200-2222110221112221-0303300000202213-3331330003213113-2301013013130301-0333123232022331-2002203130101121) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-002.md#canonical-0223013312231200-2103020030302213-3223210031031330-0233012223010022-2110300212232331-1331323103120313-2322130213203320-3032113130130102) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-0302321120031321-1031302022223301-1032231111231213-3220103021120122-0010131102123321-0133213223200223-0133210312022003-0230121201202000) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](resources--global_log_receiver--reference--group-002.md#canonical-1030312123132312-1011002010010202-0122120223121122-2122123313211032-0320020131013102-0021211310301011-1110231000311100-0332013213010320) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-0200020232111312-1113302203331230-0011233020301100-1313033200032013-1000021222011203-2130000202032331-3000013121321102-1223313013230120) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-002.md#canonical-0302122321111103-2302130231010202-2100233130102131-2010232133300130-1313330131122010-1002303111230201-2122233030331210-2023121113032331) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-3133313201003133-3230022032311232-3020001002322113-1130301011331210-0231201033101200-1010302001020302-2212010300103121-2331111233312131) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](resources--global_log_receiver--reference--group-002.md#canonical-3130231201110312-2313032121312001-2133000130000033-3013022311101111-2023011002320122-0010112210322221-2300233020333130-2211200202033133) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1113100101333332-3031020200011212-0113011233222121-0210211132002130-1121312010223302-0002133022300102-2020003231021321-2001132123010002) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-0002023331200331-0103323020333300-3213001300131331-0213202213222012-1012033211301210-2233003110313020-0032003333233231-2321130131313003) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-3011211111323330-1101102103301301-3100000032322022-0311102321320103-2323020011231330-1332010300012032-0111031322323303-0111020000311210) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-3112321202331322-0322330221231222-1033312200020312-3330213213211201-1200023132213030-3231212201212232-0010113020120231-1103233312223112) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2030133222111302-3213332031002101-2200321321200110-2021201112112012-2003202312120122-0331021230330030-1112232032132023-3212200123310222) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-002.md#canonical-2233333211102303-2101002211033330-2221010210020112-1332021230321220-3303213300002100-0111330223213213-3220123130311201-0112233212311320) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-0212210220323110-2102232312331320-0231130123323122-2323111200103133-1323113103003030-1031232101303100-3101010020222032-0130111311301230) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-0002110330232213-1110120223121211-0311033212020332-3113023000323223-2301321102121102-3320203110320113-3100310200001333-0013121110133203) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](resources--global_log_receiver--reference--group-002.md#canonical-0230223301232110-1003332030021033-2301330313131021-1310333023303130-2201131303320202-2103101133122332-0133210321101200-3110101212223300) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](resources--global_log_receiver--reference--group-002.md#canonical-0030321011133023-3000312230010333-2100310031233221-0132303023002120-3312031320233102-0300333023300000-0202201020111110-3031321232120222) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](resources--global_log_receiver--reference--group-002.md#canonical-1211202031120213-1031033321001130-0300201111213320-3003332012222233-0113131330220011-3003210230020312-3322101023012001-3212223330332330) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](resources--global_log_receiver--reference--group-002.md#canonical-1000113133002230-3130231000312130-0030032132013100-0201120132023133-3323031302001200-0111303100330012-0221320103103220-3321311022322102) |
| `http_receiver` | [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-1220320110200100-0011203102210333-0202102011023021-3001200203203131-1331100111313102-1233023130201221-3030123213313322-1030001103211002) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-3212310332211222-3232202111020220-3031310313113232-3313212222012200-2003321211113132-1013230300231033-1231033202002022-3310212132331000) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-2230311112201102-3323031123201210-3131022111310113-0230111200302000-0102332121021221-3033302311313132-2103000101122122-0302230223110221) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2021002230020210-3213022213201321-3122200021310323-1300123213003330-1031230133320031-1113102101322022-3311320011101222-2201201031200102) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-0131101320302333-2033321030133302-2000122020012211-3203313102003303-0231332112302220-1202002000111002-0003132033201322-3123331203122103) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-1210110320220213-3323201211311201-2310103123313231-3212301312020313-2303121021122321-2223203320012320-0033111313020231-0023112302212023) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-1311131111222112-0111123320012221-1331021102010331-1023121223022232-3013023233223233-3102233033202330-3332010321333311-3213020003323220) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-1211313223030302-2022231333303013-0313323321330023-1133333313300103-1203022230012210-0333200030333101-1131012131020200-2020310322000222) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-2132111303200032-0332221123110312-0012220203030100-2310230010032203-0202311203302321-2132123331113100-2332320211112320-2311111032131001) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-1002101213231003-0333103100122231-0020001021011201-0303001003133230-3122132211032330-2132213113132321-3122221003313332-0211012210113013) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](resources--global_log_receiver--reference--group-002.md#canonical-3101333321222000-2322120102103323-1302130201121123-2133330211330123-3310302030031022-3110130122022333-0222111330301223-1310103230330213) |
| `http_receiver.auth_none` | [http_receiver.auth_none](resources--global_log_receiver--reference--group-002.md#canonical-0330311131133321-1321020021102231-3100320111330201-3120000200123211-2311302122121303-0123202330213223-1101203311333222-2323220133111231) |
| `http_receiver.auth_token` | [http_receiver.auth_token](resources--global_log_receiver--reference--group-003.md#canonical-2012332322000301-3321203200031301-3321100111121022-2020101301201213-1001121221231012-0220022312301333-3111310220310033-0133203012101113) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-0131230233221123-0210000220011330-2200211322003100-2010223233331112-1011123203021122-1330203220222010-3012113320120003-1100102030031312) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2112223021201212-1010230300212213-3020312332111121-1001112113012121-3103031331131232-3002303311132311-2023030133030010-2333122031033021) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-3123332030303322-0333332022221001-3010012223232020-1331112312031212-0223112013020021-0101011220302200-2212233222223002-1312100321120202) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-3310223222121231-1131132202102230-1310202233321101-0213310220323010-1113300313331220-1322311301012033-1002033302020012-2233033112133330) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-2001221103021232-0011021201002002-2031102023312131-1233123122033321-2031101111202322-3311200211111010-1312030003021200-1223013333302112) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-1230002020101230-3213130320231220-3023330210001232-0100311230200232-2330020313302103-2033101013203312-1320330133301222-3021022012232101) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-2130322330011222-0322330300210311-0010330132222131-0232303131001220-1233012313233320-3133201010232233-2321332302221110-0110010300322100) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-1202222213300333-0303331323101133-3220120201323023-1022130301020213-3001211023030301-2311210223121302-3323332221201313-3033201201032123) |
| `http_receiver.batch` | [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0212110113100101-1302102130100131-1303323102101031-3232110012311112-0332330121310131-0012223323323212-3000230100311332-1102312133320213) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-1011130200212131-1031002221122100-2333331203030202-1203003301221012-1210202213303201-0022220303001111-2133022002121120-0211030300113303) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0220301012131310-2301132300103330-2312200312101001-0122123022220320-1322302021332103-0231112011112032-3020211031113210-1030010001310003) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-2301003313322031-0102013323111303-0232301233031212-3021103130331010-0030121121110133-0212002122230202-3313322010330111-2031202222032001) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0310223123233113-1323110321230333-3030023020121023-1121322110323203-2222130230222300-0211131330110102-0231311120221212-1123102303202300) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-3100133332212011-2213223013320023-1130102000301002-1032212303000021-3230311100310022-0301330113000323-1130010203123120-0332202030320101) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-0003000020012230-3121131132031020-0330320302100101-2331132010303000-1013332000120020-1200233231002330-0013311030322021-1210220012111202) |
| `http_receiver.compression` | [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-2320032320122303-0303033220030200-0133322101222122-2332112103133323-1031320012123010-3021022303023200-3033100330132223-1013313201331232) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-0133101033000002-0330202101311321-3213323312101301-1222030100133200-0210101213000132-0213222333131120-2210132333020301-3033320000133230) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-3331323200213231-0120103201102010-3030010331020322-3301000120331132-1333300211330202-1000021222121300-0123220232212310-3321102011200311) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-0330111322330010-3200220201223002-0220333222222333-0033310230333203-1300333130231331-2320021113221330-3210100322310122-0231133303220313) |
| `http_receiver.no_tls` | [http_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-0021010300200121-0020302101130303-0001312133311002-3122123120132133-2221222132121130-0320300113233120-0302113221030121-1311233111321202) |
| `http_receiver.uri` | [http_receiver.uri](resources--global_log_receiver--reference--group-002.md#canonical-1033020103120002-0123213303012030-3303331000212223-1303000233331233-3000022223011201-2333130132001220-3313100121010130-2122211223322303) |
| `http_receiver.use_tls` | [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-3122200203231003-0022232232103320-0303132131122332-2323323332112310-2013232023300220-2100003022111130-0030001321112012-3333303112121312) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1230011120133333-2210331123002101-0120223303323213-3030111023010000-2302201202310220-0020102232202211-0302212032330112-3132101131322313) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0022121122120312-0122300132022022-2330301012310123-1221232131121101-1111300323132131-0023130331302132-2210012200321021-1021301323222102) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-3213311211122021-1320312311322031-2301211010222032-0102000201011000-3203133233332110-3302200313110033-0330323133000012-3303030102000300) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-2120111331123120-1002033032002023-0131233322331330-0313020103112031-2002032103321023-2001311012321123-3110020202323100-3331203102300232) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-3331313133020210-2300121120320213-2012301031302101-1322313203221300-3011302131322310-3212233300222201-3210302323022321-0031130012301300) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3231003303012101-2233300303223221-1121012330210330-2201112223110213-0003120200131103-0321211230222330-2303322010110330-0033120010203102) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-003.md#canonical-0130333200031221-0322220233213003-1133311112130313-0133212012212003-0320123202121101-3202300113203131-2310203202012121-0121112210310123) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203310121332211-2031210130122101-2203021230311110-1032303331233321-2231121333000000-3033203211001032-3110133020022220-0210222320222230) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0022203101211020-0200100002133300-0221010012112322-1011202322012233-0110212003132312-1121001113230230-1323301132022113-0111001211100333) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-0112232210120133-3213301232011013-3231220220331231-3030331132002303-0102012303313130-3131221023011121-3313203031023111-2300130230310220) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-3120011313320303-3003212131102012-1013001303032302-3122102110110310-0333022132313302-2000031312001100-0103303130323313-0021333300032101) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-1111000231221000-1311332101130212-1302033300010021-2001110120022230-3220032310323123-1020222213200311-3333103301022103-0020231211201231) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0021311210103002-0323230322231121-2120030013223311-1013221000330112-0332013301220101-2002023333032331-0132021210201233-2310230233212032) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-0031130110321102-2022101230002202-2233030121102322-3332312301231202-2221231021010223-1330311113100312-2113233123330231-1330220202100231) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-0212300023232033-2132302113320130-2303011312333112-2101330321230200-1031122030131022-1110202122133312-0333211003203013-1110123330022321) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-0320303010223313-3103100300103013-0020011220310123-0110033013323222-3233313331203132-2013021030122303-1332023323110030-1211131230331202) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-003.md#canonical-0322301231001023-0020030312030110-3333311010310111-2011233120031230-0312112212300303-1212120203222010-0210231301320032-1311300003232233) |
| `id` | [ID](resources--global_log_receiver--reference--group-001.md#canonical-0331121310220213-0022123302322203-3101123131332131-0020121312302123-2133211030310101-2013203031012120-1102101332200211-1221121031131200) |
| `kafka_receiver` | [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-0300211310300033-3303232313232332-0103302032102312-3331310202130303-2112110331221010-0120101202030011-0132333301330000-1003221223001012) |
| `kafka_receiver.batch` | [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2112032212232221-1030333002202212-3302323323000212-3223031130103302-0320112232021330-2031202301002133-2023013122212303-2203133212301110) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-0101133130201200-1121011031221020-0200210112230102-3110312333233130-2200301302311333-2003223010022113-0113132102123121-2021313123302121) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-1112221213102031-1313102331032021-1233312122213201-3130030222333303-2231321313213010-3311023220100012-2333313321220221-1132100101030221) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-0010331020112202-1032322032022223-1210221330203203-0001310323002311-3020201023311000-3122001301122012-0102223123101210-2213212330031211) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0222332301112112-1323203030202220-3103013022033203-2203320200333023-2002213231012111-0103303121230030-2030021311013231-1120011333102213) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-1013211011223101-2233003023220001-0321233112232011-3003130032212003-3032213022121203-2212323200133201-1313111212102331-1110232013333200) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-2133331331213111-0313110103103201-3312202113031130-3010311000320231-3303103330103232-0330022211321202-1033202123303312-3202000032132202) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](resources--global_log_receiver--reference--group-003.md#canonical-1221210133303301-3131320221033213-0113123330213102-3023200111131233-3121001032200013-3332111022231122-1212311222202203-3203321121301323) |
| `kafka_receiver.compression` | [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-2131322222331221-2110213002112031-2300320001322210-2100201113112232-3020311002230033-0331111311303331-2102230230213332-3232323222033100) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-1013110231132100-3313101233013302-1003030230003122-2320120111230223-3011321321222110-0231312100130222-1302320101103312-2200021321003113) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-2210200010322033-3031130130333300-1313021023220310-3230222301121132-2002200331333212-1212210012013101-2332333102132003-3131100231103133) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3300301221313200-1310321202302222-3132011210312001-2113201212123201-0301003020122321-0323123132021133-2103330202211001-2233301023303031) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](resources--global_log_receiver--reference--group-003.md#canonical-0202200230302332-0301131301312301-1223322023220123-3203110321321211-1011233122201132-3133111111300230-1213001210221303-2311112233131130) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-0003333312013332-3312332113032023-2100333310320231-2212012030331222-0332123013120001-3011300232011122-3313133201230202-1010223222000111) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-3121332102100013-3310211230003011-3310331311200112-1211202111323012-2222010132122201-3023133030110011-0231031202120021-1333213123313031) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1202023110110032-3221310313122100-0212010202123132-0122012331332122-0210030212120221-3320302320312132-1202313320013200-3223321213333002) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-2133213212331130-1121230221302202-2111100110221222-2200233210033212-1213321202112010-3132332201333301-2000101001220323-0023332022211023) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-3111112213202003-3033300013212303-1103310011310113-0122130302033030-0123111122013221-3332120120023322-1310202231133012-2121012330013210) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-1101213130230231-0301030231020130-1330011000310103-2302233213111010-0033311330330220-0201111213202032-3102003032230123-3123121232303120) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-1130120021222200-3202130213203011-0112112223132001-3002320033321222-2032311310211330-1010231033001120-0021320233030133-3111230222120010) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-1131130022323103-1203321331130201-1310103333213203-3002020102120011-2001002122211211-0233313013200330-0223023330323000-3113200013003032) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-003.md#canonical-0110231120301200-3031202311211211-3121103021011312-0233313210312232-0321211331201002-0332220300311323-0321010001203313-3311113003223112) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-3300203201300012-0012033001230122-2303010112222211-3230021003303320-1103330030000313-1203131002321323-1322103102223033-1302302333013333) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0330202331311013-2221000013002320-1023002333321333-2202331030103120-2001020120113123-3312311100010030-2331110012223112-3123311313311003) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-0303101220303220-1002313302300221-0110000011310300-0220202210101020-1012300221223320-2003112303021120-1022222001122222-3121130201121212) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-1333303221122020-0331310322311000-1110331223312311-3102003020312220-3302001201000022-3302230223123323-0032302320101210-3323200323130103) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-3111032233112222-2033302111212203-1313210103322003-3213313112231010-0221030013103322-0033110231011012-2010022022200131-2032322112010212) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2100030213311013-1230320030301120-2303323230120331-0012202212031132-1122110120111110-2210232122110320-1000112302220233-3213322001230320) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-0030112321202221-3202023221313301-2221201132220210-1322231101113203-0132223230302123-1010312200113020-2232122321322033-3302322233330002) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-0002000202030112-0110002223313130-1213033302323321-0033323211220122-3130212123322023-0011111130111330-1203231021000110-0033113220332122) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-0311121221322223-3110013021312110-2102323332312233-3012123011123310-1020331301231130-1200103120013310-0000233021213232-0122002222203331) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-003.md#canonical-1112112120233110-3303111231021013-2200102232013332-1032310111313300-1032112202230210-2011203310021331-1010323311201201-0203220001310113) |
| `labels` | [labels](resources--global_log_receiver--reference--group-001.md#canonical-2020223323023222-1233310223030001-1221012210003030-3333003331013000-0100331002102000-2102133111203331-2021001311001213-3103032303013212) |
| `name` | [name](resources--global_log_receiver--reference--group-001.md#canonical-2301112001212110-2010010232130022-2320111323321011-0332022032331220-1323002223213031-3111020303110012-2113003212313312-2102311101111302) |
| `namespace` | [namespace](resources--global_log_receiver--reference--group-001.md#canonical-3312003003120321-1321311211320200-1021210321323030-2220000231310203-2001122302332012-3313023310113003-1011220223320110-0210131230313203) |
| `new_relic_receiver` | [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1230012130102110-2311120330023112-0022233130211001-0010032110111023-1303133132230232-3310210232122022-0323333321323230-1021002003003111) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0223021210302310-2012203211031313-1322122223013330-3333232121232301-0213001011330131-1222123201101313-1133230002211110-1110322300302202) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2032120110010311-0032132333203322-1132112202331322-2220130001023332-3213032222031103-0223321221232213-0033211222213300-3303200123321320) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-1310031222201233-1102211023320132-1201213110213311-1102002120320332-1100303320321332-2003302203112310-1220323111202123-2022010231212230) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-2223200113202210-3130331131313233-0223133322100133-0200013231321032-0231232231203311-3230011230103011-1301013203120111-2321323001131203) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-1133112221202003-3322132333313232-0210001032313002-2010103110102120-1123113211023121-0020310303210310-0030131031110202-0232023301112113) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0030311021310330-0131113210013213-1232010222121301-1122320200202300-0312220001232330-3100302211131301-0113022101000101-3011220003031232) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-0000022200121113-0101123102113002-0030331121313122-1103103311133202-1010213101000300-1321321132131131-0212213131212033-3102120200303022) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-2131333100232312-0221013031230030-0311001311122100-1133113130310013-0212310311112031-1031131203320301-3101021331230132-3222333311200332) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](resources--global_log_receiver--reference--group-003.md#canonical-2102232032231100-0233300020131310-1112120231023302-1031010013300202-1330132220031312-3013320300021031-3302312320130021-2011103002202231) |
| `new_relic_receiver.us` | [new_relic_receiver.us](resources--global_log_receiver--reference--group-003.md#canonical-1003110133033023-0101310221011312-2323123212111201-1332232223123003-3003100133102013-3000223032322323-0211131211032032-3300012321103200) |
| `ns_all` | [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-3002133321321130-0230232203010223-1110130322321023-0302022123020231-1112103320022121-3221331330103213-2132031121012132-0211011322121200) |
| `ns_current` | [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-0301322210023112-2021310211001030-3312222220121313-1132211011001223-2002213301202033-3323203023001212-3233001223111202-3221030231123021) |
| `ns_list` | [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-0331201313123231-3032231333310001-3022103022000000-3122210132332000-1233123202200013-1012100010010213-1321022201103201-3120010330312100) |
| `ns_list.namespaces` | [ns_list.namespaces](resources--global_log_receiver--reference--group-003.md#canonical-3311133003311012-2112320033113111-1003020223221332-0023033020003022-1012111233103312-0120211012200222-3302322101001330-1112303001300201) |
| `qradar_receiver` | [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3323203113301331-2013010003331233-3133212102210323-2032121213300323-3201011213001320-1331212033033120-2133211210012323-2133022310331130) |
| `qradar_receiver.batch` | [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2021331223032232-2031131311130103-3301232121321201-1332130100002112-1021223231300331-1012221203001221-3200312233310110-2012312002332330) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-2231321230023331-1121002301100021-0322132002100120-2020330002111131-2330022213303333-2201030021011220-0112100013113211-0120220123220003) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2103021033123232-0320313103132311-2322133023112320-3121310223102001-1120102103223233-2210311211012223-3302003113103103-2211200031001120) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-0221001313121110-0023321013210211-2032322313000310-3123120023022101-1220312302130231-2132203230000030-1313000031021022-1322310123000023) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2133021212221123-3010200123200301-3203012132120100-2212111332331322-2122200012111002-0102223202230033-0113232031322223-1221333203320033) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-2331032123013232-3003003033001121-0103210000003232-0200310101230221-2210321100211313-3233002122232231-3123112222112103-3101231010010222) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-1232330321113311-1212002031333133-1322323221022321-3132001303132001-3221310210013310-0230013132300213-1221312100201230-0112022330011132) |
| `qradar_receiver.compression` | [qradar_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-0300123130230212-0210212003303202-3200011320133021-0102110230222223-3030323320223300-0030201030311113-0031112321032233-0032000210221331) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](resources--global_log_receiver--reference--group-004.md#canonical-2222302002300322-1300312202111210-1330022331023221-3332012210021132-0023331300303200-1221133232321121-1120213331010002-0002222101001232) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-3001210031003331-1120010333231111-2230100202002110-3233022230230201-2333303212222232-2133301011121123-0233131003303211-2202330010302200) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](resources--global_log_receiver--reference--group-004.md#canonical-0002020103233212-0203321120022222-1101230000013021-1210000210221231-0313110201001101-1330022333110020-0002023102313020-3011103322131212) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](resources--global_log_receiver--reference--group-004.md#canonical-2011101121131031-2021332220303111-3233201131210220-2333212131300233-1222222330201202-3101133121102322-2202021000020111-1232000333200131) |
| `qradar_receiver.uri` | [qradar_receiver.uri](resources--global_log_receiver--reference--group-003.md#canonical-0011323232033200-1111202030133200-2301112132001000-1121023022001212-2201102222023303-2000003311222021-2031022101301322-3113123212002213) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-0233303223131023-0121300303011220-0100201323130200-1300032001302221-3020012001133222-1210003303303120-0323113312302012-3111113130232112) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-3132221033311332-2323012221020000-2013321311102230-1220011322023111-0312312321031013-1223332101310200-1012201010031103-1120133111212213) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-1033212320030331-3201021332321121-2302020221233001-2310321212020230-0231003101110212-3330222021122123-2030102302232313-0222020110102031) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-3103021300113221-1201123113100233-2023110312001011-2132002123200300-1033111222002122-0213210100230131-3203313001132003-3123301001121300) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-3112203002022122-0301133021332332-0101300002121221-2010120311333130-3102203310230002-3113333021331332-3121230320032102-2220020032200330) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0230303311011320-0020100002030302-2112211030302123-3221232200303131-1231103303320133-3120132313120331-0222031031000303-1030022003031001) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-1033311010032220-2002033203232031-0231230101001123-1332321122011200-1133023113122322-2103000101313132-0300230022201012-0322001220210311) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-004.md#canonical-0220203200212013-1200132030213013-3210011303023001-3021010332011131-3312233122321001-1202120221113013-3101333210111132-0210203023120001) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-3131312021101220-3322023330132313-3213111223310020-3313223203311331-0012122022210103-3322313302210001-1211302031020211-3013233102201133) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-3201330322110011-1122120022000322-1011330312003332-2012313211322231-3331301001213102-2220121100311011-0000032013210220-2323002210202103) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-3303212203111202-2102121000011223-3002213231303111-0310213010112313-1312000112002030-2032232210223201-2222110331102231-2030033313031101) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-0232031330230232-0121032122111003-3113320202001120-0233202323123011-0111030210212103-3221230233230010-0201030132303121-0010111301131100) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-1333010102230020-2030223210133111-1311300201320301-0230203011222320-1321001013121312-3210233220110021-2323100203312031-1121301030110113) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2322131313322332-2332013213120223-2223120133323332-3003330232130120-2120031001013012-3320301311330030-3232321232033132-1023321231033102) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-3213321001322213-1132332323301232-2320230022013020-2232131032130001-3120202211300101-3023311022031013-0113100133310323-2012131233033123) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-1032221303020130-2323222203002113-2032002023121003-0030233233012332-3130223123210121-1210033102112101-1103020031221303-3203203221121121) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-004.md#canonical-1020301303303302-2101011112310022-3232101302031310-3232331233033113-2030101222310231-2233200312001313-1210000202203110-1322021030101110) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-004.md#canonical-3203122322221121-2201321003222222-1230012021220301-2120232002222320-0021310330311223-2012303332311020-3302203320221322-0211331323033011) |
| `request_logs` | [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-2112333332233110-1101303330101330-2203303321021123-0311130301133120-3113001330201322-1301233113213302-1011110010211003-1302210033220211) |
| `request_logs.sampled` | [request_logs.sampled](resources--global_log_receiver--reference--group-004.md#canonical-1021022020311001-3323233000000011-2103013131202320-1031110313222320-3233311331223111-1312102200230323-0012232110003332-0021002211001220) |
| `request_logs.unsampled` | [request_logs.unsampled](resources--global_log_receiver--reference--group-004.md#canonical-1103211013210211-1302101123022230-1332332100130233-2221202000223123-1301322302322023-3122210110131101-1233121232030020-0112222322202020) |
| `s3_receiver` | [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2120303131032111-1131001123333020-1133113103233030-3000301211202000-1233230031302020-3210203032233122-1203232331022303-3313332301220301) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](resources--global_log_receiver--reference--group-004.md#canonical-0312003122213131-1120331302323313-1001103331301332-0012113311200000-3132021331332000-2221313211321311-0002300202023010-0130102311133032) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](resources--global_log_receiver--reference--group-004.md#canonical-2320023220212200-2201030100111002-2032130021232223-2331120031212231-1201332210300230-3221022220021223-2212100301330210-1130130003133121) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](resources--global_log_receiver--reference--group-004.md#canonical-3121302031103131-2112223002112321-1102123021302003-0221033031223021-3032020203001311-2220122220302200-3121003331020211-3222203132331312) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](resources--global_log_receiver--reference--group-004.md#canonical-3001233221333213-2002323321311022-2023202233032312-0023122213320032-0321132120212133-1203203210002031-2231302112121222-3103113301220212) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](resources--global_log_receiver--reference--group-004.md#canonical-1231020123112000-3323111222222231-2320212313010223-2130131013320113-1300100130331210-3311203012031022-0232212100010122-0320003113120012) |
| `s3_receiver.batch` | [s3_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2331332031200122-2130322330312232-2000302331232330-2033223120331123-2320231200033210-3123231100333112-1012133110021323-2203300101210020) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-004.md#canonical-0113101013331121-3130111212012123-3321030210031320-0232113101331210-0000212130121331-1221222130101001-0002103012203232-1103033010021130) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-004.md#canonical-3023230222202322-1322032200312333-2113321022221013-3213100103023331-0013132023000210-1200222021223322-3202103312321032-1112002210102102) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](resources--global_log_receiver--reference--group-004.md#canonical-1000333231232102-1330102030003021-3031303123201210-2121021330221002-3333330332330001-2203213320022233-3223322302231130-3031120003130222) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0022222122101311-2200322223103313-1103312312023330-2101103001310133-2301313221301123-1131223330130003-2320331321213220-2222220110200323) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-004.md#canonical-3233213113101112-0321213132031321-0203110103311031-0021323212303032-2332222221202023-2001001100302322-2213023313003310-1121100323022223) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-004.md#canonical-2312332231221303-2130010031321002-1031200133223200-3203221022130201-0003003112002101-2033300200011313-1312323013311332-1132233113123111) |
| `s3_receiver.bucket` | [s3_receiver.bucket](resources--global_log_receiver--reference--group-004.md#canonical-2213211332001011-0210130101011103-3333113031120333-0330212030023231-3210021231310231-3333232301221232-0213000032320201-2122211111302231) |
| `s3_receiver.compression` | [s3_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-1011210030121103-0232020132001132-3211333102230330-2120031231032323-3321010332020310-0030301013332320-3202233221230300-1133201023001312) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](resources--global_log_receiver--reference--group-004.md#canonical-3013230132333031-3201003101131320-2303120200203300-0221213333002122-0120223301131203-3211012132000130-1220320230200203-0122111313020201) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-1032002013300213-0020223002132132-0010200331211310-0110131200213011-0132331330023201-0033201232311330-3003220302131100-3031003210330030) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](resources--global_log_receiver--reference--group-004.md#canonical-1311201200012203-2332233221333223-2211120102313330-0120213232231312-1010310201210221-2312131300100211-0303211223010312-2221112032002212) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](resources--global_log_receiver--reference--group-004.md#canonical-2012320332233303-2002230101231132-3231323102013020-1000333131330223-0213122213001223-1102203333033013-3000233320132320-3112333220233223) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-004.md#canonical-0022012322321123-2013112133121200-1200102111122211-2010020120103032-1320201012330210-3302130003302201-1030113203030323-0201202200100010) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-004.md#canonical-0320023300210110-1213003231020313-3333011202213100-3023023113201212-3320123010330031-1113012013333031-2122120132230211-3021130033212113) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-004.md#canonical-3233202022203000-3103123123221132-2210213323000210-1231320213311111-1001332223000103-0113002222201121-1000100101202133-1030030231230131) |
| `security_events` | [security_events](resources--global_log_receiver--reference--group-004.md#canonical-1211211103013320-0312220220033311-0221122101003311-3101230113233111-1103132232203013-3121322010213021-1202332113201032-2002003231210233) |
| `splunk_receiver` | [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-3323012321222032-0301123231022010-0310113310103303-0311323212010001-2321221011222303-3332101132321300-2330331123203201-1202330320122010) |
| `splunk_receiver.batch` | [splunk_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-3001101301220201-2313210130330123-2312220230201203-3201122231333113-1102330203330311-1330030132330233-2133011303111213-0300113201302131) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-004.md#canonical-3222330330231301-3221221032030102-2020202001013301-1102202331201002-3022023031220120-2320122212010111-3111221111123000-2300122203023230) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-004.md#canonical-1130103322320110-0232222220021222-3220233202300102-3322022001013211-1300301222323131-2123300303000123-1333312123300322-0112101120120200) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](resources--global_log_receiver--reference--group-004.md#canonical-1302310120203021-2101100020320032-3002102201223002-2013321101211302-0332300013103001-0011110212322213-3111203031030322-1222112123122023) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0211320000123013-2301203330032120-2120332301313232-1130032003210130-0113332012111011-1230223122020023-2221310223211030-1120321131123013) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-004.md#canonical-3102120222323323-0202220321211302-1222213031302320-1000312112303112-0213122023223023-0223323122203100-2121132320103203-2122100200332303) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-004.md#canonical-0002211212013112-2211313321222322-1230121100003302-0231133331021113-3212330132201133-2233113123301030-0203122320231110-1030302013211012) |
| `splunk_receiver.compression` | [splunk_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-0123102103323303-2301231012333233-0231133100032202-2002200312130003-2212032232220002-1133003113203100-3330120300202211-3113003222130030) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](resources--global_log_receiver--reference--group-004.md#canonical-3133220301122210-1231223220302230-2321321233020110-2303332111310033-1213303101230010-3312333221330220-2302132230020131-1212321120201311) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-3033320211220110-1032120223122323-0012230130023302-2102101100213211-1233130133323112-0032033320221002-2001212310001111-1231101001310201) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](resources--global_log_receiver--reference--group-004.md#canonical-3213231113033213-0202202113131310-2301300323330003-2322031300200200-1013031223033331-0022221302311003-0310013200310013-2321003231302330) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](resources--global_log_receiver--reference--group-004.md#canonical-3211201101231311-2211021131220000-2102311221031133-2202131000020002-1322322112232330-3300120300213333-0300233210202033-2031010000012211) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](resources--global_log_receiver--reference--group-004.md#canonical-2302200133232232-2210300012333311-3113310120313132-1110312113332100-0203130310212231-0103233233131011-0331032301103010-1123332021222332) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](resources--global_log_receiver--reference--group-004.md#canonical-0212200222200100-1022012001210000-3330020001330201-3302103121101113-2200210320313132-3332120320213030-0000022333333020-3232100021302021) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-1113101201321200-2033223030020321-3132233233232313-0203030302223023-3110223101103201-2230000312103203-1303003201311310-0023033032110222) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-2110003321223122-1112231210032310-2132111032000003-2333001331321132-3133030233233223-2033101201211231-3211131030101012-0032230212330020) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-2311012201002231-2131031110222313-3323300101311012-0013310213003303-2300220202023211-0312203303322121-0123101200301000-3102210102310122) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-0102013302202322-1300302033222221-2010201010212030-0232103230331101-3112032131121310-2132003303331311-0200012002200020-3323030100123112) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-3220033202123033-1022023311133120-2223110223210102-0132331212312202-1000233121200302-1230223321130002-3323221002110020-0230033301111320) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-1013033223000333-0302113230311233-1031321002001212-2210122102123120-1020102122323023-2013332010132023-1220111312203200-3230030302212011) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-2100101001013311-1221320300033021-2132121101011210-0303020111311322-3011231130023100-1122032101223011-0022012202300100-1231133300003032) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-1301012031221211-3010122330230213-0020032213300201-0103010210220100-0101322202133201-0003333032103312-2210320013011312-3033202212330333) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-1002332031302211-3102311302032101-2303012113321113-2102231201022022-2311203222212111-0020001213313311-2012122023223202-1321123112130300) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-0313022131210333-3330113001131031-0030123332113211-2213323010310123-3003211003213203-1110233111030203-1012331020233133-1013211213130231) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-2331101230323102-3313131111212102-0311123303021221-1333320233003012-1100012123331123-2300021121330121-3030030221212120-2032222323321011) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-3310323221111010-0220132232310030-3001113200031303-2303311001302000-3331320110012131-3331310201222030-3013120321333023-3203211210001003) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0020121012101032-1033120030310110-0003310031003332-3300030332230122-2001011322123200-3013211312112121-0012322001221201-2033213021100221) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-0210231320311301-3103331023000321-2322122012101022-1230331131201310-1123031222332120-2020003320030133-2033131212021012-2312030113113330) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-004.md#canonical-3323030330301310-1323021133013320-1313333010331331-0210113130231031-0203130132110211-3012003012012032-1320302312121320-2321321003031212) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-1001303023230033-3213231320333033-3302020010223302-2111223100000303-0300312101100032-0333310202203103-2131121213213221-1320222333230010) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-3021313132232111-1131211332213330-1211123012033123-0000121003220201-3321222303220113-0310213313131312-3000123020233310-3220221003230232) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-3003032223302033-2132120033330112-3132321121020120-1022332330001211-2100332310322003-0030221120131202-1133321211200300-2223201012220131) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-3330100222331232-2121301013312310-1331232232332331-2223220312021131-3123223220132232-1113122122221302-1033102222231112-1322133023122022) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-1110132013220002-2331113232102322-2110302222210111-0122012112202201-0110122013212231-2332301000233221-2130232202130303-0313021031312031) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2230332321023310-1001122100210111-0023120323020220-3232023222202111-0100200001012002-2113112111010221-2110112011012111-1230112012123333) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-1021311213001100-0111031210121221-3022023321320312-2102013332222230-3110111131130033-2311323230313322-0200222202232202-2230231220132100) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-2210032320313023-1113021232210033-0331330133221130-2321130001130101-1031113123231210-0201321211323310-3033322131130320-2131220231203221) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-004.md#canonical-0012022201312133-1133210220001032-1300231332103210-3321310303032110-3000222301233320-0000031010232001-3123132223320121-1002122313322021) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-004.md#canonical-0220011300222101-0330103123330110-0200321332123021-1310222103003100-1112011003023310-3322030003223030-1032211212021120-3312332020303021) |
| `sumo_logic_receiver` | [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1113220201232312-0211101312212201-3102312333303113-0100111113320322-0002010113322332-0120002221203122-2300101023133211-1222013031111322) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-004.md#canonical-0331212021213032-3020113200223320-1121232303120211-3120023213233311-0301110331132002-3201200333223301-0211122303000021-2020223330322213) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-3010203311220001-1002210211100002-1031031113111332-1222121030213131-2030031132322300-2101300312020313-2001313033001301-3322131120011220) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-1111130001013120-2332213311000231-3313231031102033-1013103302132130-3022133133120030-2000231033231130-2221032212311230-2232130230232233) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-3030231301113303-0230220031233230-0002003302121332-2122333310131331-1300100021302010-3100333011330020-3333112130122230-1121020023220022) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-005.md#canonical-3222013330321333-2212200112223103-3121022313300311-2110223113000330-0222232202303221-3001313023112003-0330031133001303-3031200232122321) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-3230133031132200-3013323300021221-0112022130230010-3202231331223233-1020232013320220-1302102100022000-3020202200110211-0120121310313130) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-005.md#canonical-2311220113323202-3332332033030300-1330231021010113-0320111133110010-0112013000310000-3010220131131323-1000313020310201-1200201322000011) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](resources--global_log_receiver--reference--group-005.md#canonical-2321122321002030-1210232310222233-0110123110130110-1302010200301031-1201130100021020-0300011301033012-2000332231330311-3123203312321002) |
| `timeouts` | [timeouts](resources--global_log_receiver--reference--group-005.md#canonical-2323101322201020-2201013121231321-3110201100222113-1023233112330212-2113123032321331-2031001100133233-0232113100210011-1212031323321012) |
| `timeouts.create` | [timeouts.create](resources--global_log_receiver--reference--group-005.md#canonical-2130100022000302-0233020112231100-1203103031203200-0221232133223120-1010202210130203-1301002211010230-0133223110303203-0330233303323013) |
| `timeouts.delete` | [timeouts.delete](resources--global_log_receiver--reference--group-005.md#canonical-3121121313100001-0210133021230333-2012103322333331-1311011223032023-3311113113131130-2203312331030031-3233013333221131-2103111212303002) |
| `timeouts.read` | [timeouts.read](resources--global_log_receiver--reference--group-005.md#canonical-1221022123333101-3003311211233102-1033203310301211-3111100313022012-2231100203331300-2033122101323220-1211312002313233-3001222231021020) |
| `timeouts.update` | [timeouts.update](resources--global_log_receiver--reference--group-005.md#canonical-0313011203010113-3312002121230013-0210100012300332-3111003220030121-3033102121222311-1321312032232010-3002210210020333-2102131202223102) |

<a id="canonical-3112112122213313-3011011300301011-0213111201100101-0022301022112102-0132022303003012-3323213210112123-0022123312013331-3312013313233031"></a>

## Next pages — Property reference / 320133321330 / 12

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-3231131323212011-3030323112222132-1230222222202231-3212031220330031-2231302201012313-3103320303103003-3102113033100033-1132112103231012)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-3211233331121300-1322331131001220-0102011110001201-1030212021001001-2311133020330123-0200201221122221-1011112033100211-0230123330102313)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-2110223021133332-2001211101313323-3012032332230023-3123201031102303-3033233222010310-0322300301310121-1032103033320210-2132202001210101)
- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-1131102221232000-0300222020002322-0302133211001210-2101203323112122-0311121123102201-1122110013312213-1303121113303221-1111202320322233)
- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-1323013202223030-2213130333303312-1001111102133131-1001220102221312-3002003012031130-3113012231221100-0102230132130332-0202133321133203)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-3211022021302113-1000032210211020-3113000022100301-1332120310011110-2113003230211310-2211212131301002-0201111232112322-1122230010302122)
- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130)
- [timeouts](resources--global_log_receiver--reference--group-005.md#canonical-2300030330132202-1012220213322303-1300301203131302-1312202110301333-1300333322313010-2203010102023132-0210122130033122-1033121302122132)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3231131323212011-3030323112222132-1230222222202231-3212031220330031-2231302201012313-3103320303103003-3102113033100033-1132112103231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123332100302233-0003220332132330-3112030202023300-3201203003122231-0111200010121210-0030323200303212-3201220310131200-3112002120032123"></a>

## audit_logs — audit_logs / 332120321222 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- audit_logs

<a id="canonical-0101233320301302-1213113233000311-2013000202300333-3020021232010212-0210010120233212-2021333333322130-0002331232212331-0023310101122033"></a>

Type: `["object", {}]`. Optional.

\[OneOf: audit\_logs, DNS\_logs, request\_logs, security\_events\] Enable this option

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

OneOf alternatives in this subsection:

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-0101233320301302-1213113233000311-2013000202300333-3020021232010212-0210010120233212-2021333333322130-0002331232212331-0023310101122033)
- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-2112120022003322-1202120213333133-0010302200210023-3232021030212312-0232123211210022-2203312323223132-3100211030221032-3112323313223031)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-2112333332233110-1101303330101330-2203303321021123-0311130301133120-3113001330201322-1301233113213302-1011110010211003-1302210033220211)
- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-1211211103013320-0312220220033311-0221122101003311-3101230113233111-1103132232203013-3121322010213021-1202332113201032-2002003231210233)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
audit_logs = {}
```

<a id="canonical-3032020200103013-3312130200230030-0102220221220111-3223333331320020-2133233311003213-1031112222322120-3330130311313221-2111321031131110"></a>

## Direct properties — audit_logs / 332120321222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023002021020213-0123321123210300-0323303203310021-0300303003322202-2212100133200333-0003103313112332-2101331032221102-0323323311333130"></a>

## Next pages — audit_logs / 332120321222 / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302220312112000-3012103121033011-3311332311023301-0202203123032300-3333222302010022-1222211001210020-3110321002132131-2232122222230111"></a>

## aws_cloud_watch_receiver — aws_cloud_watch_receiver / 003312112311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- aws_cloud_watch_receiver

<a id="canonical-0010031130120012-1201320000231232-2331132333100100-0310012020113322-1120212131013223-3021312000023122-2333123112111130-0212122212030102"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_cloud\_watch\_receiver, Azure\_event\_hubs\_receiver, Azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Upstream description:

AWS Cloudwatch Logs Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "group_name",
    "stream_name")}
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

OneOf alternatives in this subsection:

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0010031130120012-1201320000231232-2331132333100100-0310012020113322-1120212131013223-3021312000023122-2333123112111130-0212122212030102)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-2320200132011030-1231332200110211-0002331211012320-2213001200201030-1011331311001011-1022302212331332-0103310102303100-2031212300120002)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3013200231333020-1303012020320320-1001100130303121-3211010010010303-2320110321121200-1102123310122230-2100321311330302-1302111303311231)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-0003033200133213-2021130011102033-2321233212202033-0010010332230103-2101001230301213-0313001301111103-3112321322103202-2021022012211332)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-1221011012130302-2001213102221021-3130230103112313-1211200120113032-0230233232020101-3331301323120221-0220020210111122-1233101312331111)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-1220320110200100-0011203102210333-0202102011023021-3001200203203131-1331100111313102-1233023130201221-3030123213313322-1030001103211002)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-0300211310300033-3303232313232332-0103302032102312-3331310202130303-2112110331221010-0120101202030011-0132333301330000-1003221223001012)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1230012130102110-2311120330023112-0022233130211001-0010032110111023-1303133132230232-3310210232122022-0323333321323230-1021002003003111)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3323203113301331-2013010003331233-3133212102210323-2032121213300323-3201011213001320-1331212033033120-2133211210012323-2133022310331130)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2120303131032111-1131001123333020-1133113103233030-3000301211202000-1233230031302020-3210203032233122-1203232331022303-3313332301220301)
- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-3323012321222032-0301123231022010-0310113310103303-0311323212010001-2321221011222303-3332101132321300-2330331123203201-1202330320122010)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1113220201232312-0211101312212201-3102312333303113-0100111113320322-0002010113322332-0120002221203122-2300101023133211-1222013031111322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_cloud_watch_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231133213021023-2011120213022300-1103222201231122-3033013311000003-2332133231110012-0120100033222131-3012222301333333-1300201113313203"></a>

## Direct properties — aws_cloud_watch_receiver / 003312112311 / 3

- [aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-3121021212132202-3132213011203333-0002323133332313-3121330021323223-2303230220013323-1113000231013011-1122331000031030-1122033132230301): complete subsection reference.

<a id="canonical-2013030030212011-0333122012030213-1031323111100201-2113221120013322-1003021033010003-1031301201222300-1002001133221020-3031213112100100"></a>

<a id="canonical-2321030121013003-3133110123322223-0223012213131321-1111203120203000-1120333331303332-0021031002013230-0223111312312323-2203330131232111"></a>

## aws_region property — aws_cloud_watch_receiver / 003312112311 / 4

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201): complete subsection reference.

<a id="canonical-3322300301331022-0100200310011211-0011033212322032-2233312332001100-3120232201102013-1220110122233111-0200030010220322-2302330000203212"></a>

<a id="canonical-3022032013000310-2200220233230120-3030201112012231-1301130310102312-2220333023222323-3121112013231121-0230323102210301-0000031213122001"></a>

## group_name property — aws_cloud_watch_receiver / 003312112311 / 5

Type: `"string"`. Optional.

The group name of the target Cloudwatch Logs stream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  }
}
```

<a id="canonical-3303221103330231-3211220200132020-0320302121222312-0310332013122003-0203330110222310-1212211301110232-3231113221233201-0100013313230201"></a>

<a id="canonical-2013212011202131-3330231123123201-3023232113203023-3321121031232232-2121023310130300-2033333232103020-2310311232321130-3233300312030000"></a>

## stream_name property — aws_cloud_watch_receiver / 003312112311 / 6

Type: `"string"`. Optional.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Upstream description:

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[^:*]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  }
}
```

<a id="canonical-1000201310333221-3011200212321001-1013111313312000-0232130021322330-1321022201002221-2333100333102111-1001030120033101-2221100222023110"></a>

## Next pages — aws_cloud_watch_receiver / 003312112311 / 7

- [aws_cloud_watch_receiver.aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-3121021212132202-3132213011203333-0002323133332313-3121330021323223-2303230220013323-1113000231013011-1122331000031030-1122033132230301)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3121021212132202-3132213011203333-0002323133332313-3121330021323223-2303230220013323-1113000231013011-1122331000031030-1122033132230301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330001100002203-0033020222013111-3013311030032023-2310111230330300-2213021332201131-2321202332100022-0011300211101222-0313112212200122"></a>

## aws_cloud_watch_receiver.aws_cred — aws_cred / 230312131011 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- aws_cloud_watch_receiver.aws_cred

<a id="canonical-2232132021012300-3021231320112302-1130100210132221-0221100333103023-0232032321131310-2213230330320101-2131031030313323-1330203213300321"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2331312331310300-0323321032211320-3012212122212130-0010103210033101-3003010012332033-0002032213310222-0011033011022313-0013220213310120"></a>

## Direct properties — aws_cred / 230312131011 / 3

<a id="canonical-0321000013211132-3032323300011313-0032000232003201-2012220332313312-1122100112223130-2000012321312322-3100012301311330-0231002013302013"></a>

<a id="canonical-0311222021033203-3321001330032000-2102003221012013-1112301122201012-0023003332312002-0032223301233011-2200221120232102-2110201113103310"></a>

## name property — aws_cred / 230312131011 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3123303032220202-1213021102311131-0022210003020222-0101213232232232-0130222202010310-1001233101212303-3301121132203230-0202312231212033"></a>

<a id="canonical-1112010132221301-3120110330113221-1122012301331322-1003322133121233-0202331112120212-0223201201320032-3312222110133112-2131213022032033"></a>

## namespace property — aws_cred / 230312131011 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0001102000000220-0030203110132301-3130332103133221-2312030331130003-3000312231023132-0201330203223113-0022310130203233-0111013233011223"></a>

<a id="canonical-0112230121000000-0302102311110210-1203320020120300-0103200023020133-0223331303221003-3331222303012313-3310130000333111-0213321303023033"></a>

## tenant property — aws_cred / 230312131011 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2110321002121111-3030200320121132-3102020302001232-3002221002332010-2021312013022132-1100123000102121-0303131012000223-3300221013100301"></a>

## Next pages — aws_cred / 230312131011 / 7

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222132001323103-3231321303313213-1221000130002113-1202003021320021-0132311011232122-1221003022211211-0321020033213003-2312132301031211"></a>

## aws_cloud_watch_receiver.batch — batch / 213220310231 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- aws_cloud_watch_receiver.batch

<a id="canonical-2203331202332200-3323021210002333-0330112200301333-2332303021101123-0012203200203203-2323032210202322-0222100330212003-1030222230222131"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132020022310003-0301112211322233-2203332112232303-1313321131121332-2122313013330232-0003021133203203-3022201231223323-2131033031210301"></a>

## Direct properties — batch / 213220310231 / 3

<a id="canonical-2332020313332001-1030112322013022-1021321300120211-3233121321011200-2312030030003122-2313002312302131-1213111211100011-1000213112013321"></a>

<a id="canonical-0002031023323111-3030201330001003-2311200302120112-3010021110002110-0021221112202202-2212002321321233-3301122111311033-2031111323210013"></a>

## max_bytes property — batch / 213220310231 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-0112333202213111-3012310213332112-0303121102121312-2301013310132233-0220033011011113-0121022320133203-1202111022023010-0130330220113110): complete subsection reference.

<a id="canonical-0023232300211033-2110230220102223-0301131012003002-1120313031203021-3201000220303011-3230323222101010-1323301022331201-2112100202113230"></a>

<a id="canonical-1121320200030003-3303231103021232-1232300330132321-1320003220122013-3202210222312032-1001303123123113-2001321013333302-3312203031222211"></a>

## max_events property — batch / 213220310231 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-3323120002300033-0331303333011330-0031233010332312-1000320112313321-3021223123033012-3021100311320310-0001112313313220-0200133031102001): complete subsection reference.

<a id="canonical-3222303132331023-0102203220312111-1220221123021133-1310101322322100-1322200232132132-3112102131101011-3302302122222201-0300321320310213"></a>

<a id="canonical-1113030113320013-1211200020122330-1033130003011331-0100010333113123-1311112012010131-0221313002221010-1122212122302001-0213322110210023"></a>

## timeout_seconds property — batch / 213220310231 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-1330020102111032-2101030322010323-2021110331100011-1112300133111322-0222202132023103-1103223301120303-3113033321322132-2332022230303200): complete subsection reference.

<a id="canonical-1111233333130213-1122200023102221-1203111000230300-2213230230033230-2202002021310003-1323301222111223-1200221331011220-1233223113302103"></a>

## Next pages — batch / 213220310231 / 7

- [aws_cloud_watch_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-0112333202213111-3012310213332112-0303121102121312-2301013310132233-0220033011011113-0121022320133203-1202111022023010-0130330220113110)
- [aws_cloud_watch_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-3323120002300033-0331303333011330-0031233010332312-1000320112313321-3021223123033012-3021100311320310-0001112313313220-0200133031102001)
- [aws_cloud_watch_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-1330020102111032-2101030322010323-2021110331100011-1112300133111322-0222202132023103-1103223301120303-3113033321322132-2332022230303200)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0112333202213111-3012310213332112-0303121102121312-2301013310132233-0220033011011113-0121022320133203-1202111022023010-0130330220113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013030002020113-2313231013132010-1201022010101001-2000230320012002-2231222010023210-0012021030210233-0013331021230312-3032220321233203"></a>

## aws_cloud_watch_receiver.batch.max_bytes_disabled — max_bytes_disabled / 023111121100 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- aws_cloud_watch_receiver.batch.max_bytes_disabled

<a id="canonical-1033112113301211-1032213012201102-0211333120003113-2210200111121320-0320310132100302-3223133133213301-3310001102011210-0002233220212213"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-0102113012312120-1110302311320100-2113222020230312-1110020022320012-2231013221231122-0321212113202301-2331223021310120-0101303230202012"></a>

## Direct properties — max_bytes_disabled / 023111121100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122102020121110-1302323101101033-0222230332111032-1300100121311323-3222212133213322-2212120010030200-0000331301332220-2223132202122233"></a>

## Next pages — max_bytes_disabled / 023111121100 / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3323120002300033-0331303333011330-0031233010332312-1000320112313321-3021223123033012-3021100311320310-0001112313313220-0200133031102001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302200013002201-0132021303122100-3112013332213211-0311323112003212-0213000202213212-3133023111032213-3131031323011031-3303311320210322"></a>

## aws_cloud_watch_receiver.batch.max_events_disabled — max_events_disabled / 122120103222 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- aws_cloud_watch_receiver.batch.max_events_disabled

<a id="canonical-2202100312230333-2010121313112133-3112033310033320-0333133230233132-3111222320210103-0310130102020103-0003322222222102-1213001022121013"></a>

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
max_events_disabled = {}
```

<a id="canonical-3231202332301031-2232312203320113-3021001000032030-1100230112131223-1022032211313232-3031301103301110-1010311013221023-0230201130331122"></a>

## Direct properties — max_events_disabled / 122120103222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330221310100332-1010231023131012-1303130232111010-1123322310021013-2303223202002321-0120112100333112-2033231222030330-3023203032022023"></a>

## Next pages — max_events_disabled / 122120103222 / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1330020102111032-2101030322010323-2021110331100011-1112300133111322-0222202132023103-1103223301120303-3113033321322132-2332022230303200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300112020030010-3310310021000302-0331300320323201-2232030020023312-1100222101011321-1103130023232331-0133302300110202-3310022031201323"></a>

## aws_cloud_watch_receiver.batch.timeout_seconds_default — timeout_seconds_default / 123233123322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- aws_cloud_watch_receiver.batch.timeout_seconds_default

<a id="canonical-0212020112113131-0332213230302223-1213201300303333-0033112222022021-1210300222003233-3322020130003023-1032001223301103-2312333312213033"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-1211222110220132-3230121331321333-2221220110132301-2032222300320132-1011202203102330-0330020233210200-3033312201211113-1322103032202330"></a>

## Direct properties — timeout_seconds_default / 123233123322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212112221311002-2133321110101320-3300301333110331-1220112132032100-1301110011121312-2130123120312302-2032213221100233-3312310231313312"></a>

## Next pages — timeout_seconds_default / 123233123322 / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-1013223231021011-2000200111231230-0331301013230332-3322113300333102-0001031232123131-0123031031013323-2322322303022210-3330230131120002)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013231323230301-3233002200322223-0000023000321300-3312023112333032-2312301030022212-3032003231010100-2212331221133321-0322200203301110"></a>

## aws_cloud_watch_receiver.compression — compression / 123031110033 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- aws_cloud_watch_receiver.compression

<a id="canonical-2231303203110303-0102302202220020-2001130321101003-2121001032030132-1210220300120202-2222003210201023-1001003101202133-3201020122122010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021101200231331-2122310200013321-0120333313322102-1300320333132310-0101302030313120-2211212003211230-1233212033320022-1311033302031313"></a>

## Direct properties — compression / 123031110033 / 3

- [compression_default](resources--global_log_receiver--reference--group-001.md#canonical-0332320020311111-2100103000112223-3331303013333122-2113113020200032-0322302011303333-0121032321310211-3111331110100031-3312001330230031): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-1202013120122012-1232200123000031-0002203332322020-3002023123200223-2033000211102112-0313013201031211-3020031022302222-2323312121112022): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-001.md#canonical-2003022313310032-2132133201023322-2002020220303212-0033220122201002-3022210021302300-3201123030021132-0311200300303012-2123223031212303): complete subsection reference.

<a id="canonical-0321033222220303-2111113020130121-3112313230332332-1310103031012120-2132112200013300-0013231031012012-0233102233011110-0113330031120322"></a>

## Next pages — compression / 123031110033 / 4

- [aws_cloud_watch_receiver.compression.compression_default](resources--global_log_receiver--reference--group-001.md#canonical-0332320020311111-2100103000112223-3331303013333122-2113113020200032-0322302011303333-0121032321310211-3111331110100031-3312001330230031)
- [aws_cloud_watch_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-1202013120122012-1232200123000031-0002203332322020-3002023123200223-2033000211102112-0313013201031211-3020031022302222-2323312121112022)
- [aws_cloud_watch_receiver.compression.compression_none](resources--global_log_receiver--reference--group-001.md#canonical-2003022313310032-2132133201023322-2002020220303212-0033220122201002-3022210021302300-3201123030021132-0311200300303012-2123223031212303)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0332320020311111-2100103000112223-3331303013333122-2113113020200032-0322302011303333-0121032321310211-3111331110100031-3312001330230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221102221202320-0321211320212211-3300233120203022-1320320132000031-3032223221020003-1223331300322210-3221003230002011-3023011210103101"></a>

## aws_cloud_watch_receiver.compression.compression_default — compression_default / 131032112203 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- aws_cloud_watch_receiver.compression.compression_default

<a id="canonical-0300231032011203-2031222213220233-3023213132332012-0100000330023221-3323003303031302-2203132021020122-1012031111233121-3333313303212313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-1120100203131330-0000312302132221-1103210222223322-0122230321110023-1020210311203121-3102030230310202-2312220203103203-1320200213231331"></a>

## Direct properties — compression_default / 131032112203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121103331330210-1110221210222322-1220230223332313-1302032210131002-1312311212232332-1123221200022032-2102301320011311-2123310000010233"></a>

## Next pages — compression_default / 131032112203 / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1202013120122012-1232200123000031-0002203332322020-3002023123200223-2033000211102112-0313013201031211-3020031022302222-2323312121112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101230330002001-1211211203322000-2021020200201102-0211312331331021-2330021002120330-1312321002302120-2202122213233320-0111200010212313"></a>

## aws_cloud_watch_receiver.compression.compression_gzip — compression_gzip / 300103013100 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- aws_cloud_watch_receiver.compression.compression_gzip

<a id="canonical-3230230023231122-2012223223233330-3022301121020102-2322320320022222-2321123001012203-3033030311211231-3022311230123023-3003111330222033"></a>

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
compression_gzip = {}
```

<a id="canonical-2101233222103003-0302230001011121-3122231021311212-2131200113001102-1103130221112133-1001201333033323-0311003232032030-0210233233231110"></a>

## Direct properties — compression_gzip / 300103013100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003233001123121-3201101121011300-0120231010103202-3112310310011002-2330303001220210-3332301221333103-1221133300103111-0020210202202113"></a>

## Next pages — compression_gzip / 300103013100 / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2003022313310032-2132133201023322-2002020220303212-0033220122201002-3022210021302300-3201123030021132-0311200300303012-2123223031212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012121103133031-2331310200333211-1001311220130222-0000300030103133-3213130200020131-3333321320302120-3031010101002010-0232002023123030"></a>

## aws_cloud_watch_receiver.compression.compression_none — compression_none / 130313133231 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- aws_cloud_watch_receiver.compression.compression_none

<a id="canonical-2231221310103120-3330211120232021-0010232200301210-3331033330123021-0301331032210103-2101113311122223-0010133132210320-0203003002031211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-2000103023131320-2000322201022220-0122032322131200-2101131023222212-0103302033333103-2302301221333222-2332231332101302-1222121233333002"></a>

## Direct properties — compression_none / 130313133231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331323231322231-3013213010131033-1203033123330210-3330011002101303-3022011030312001-1023102111123102-2310310031102333-1320302310230213"></a>

## Next pages — compression_none / 130313133231 / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-1002330123100331-3101023131011333-1032333111233033-3220012130133212-0021220022313110-2321230020210011-3330323231320002-2112231300113201)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133022303230200-1013033332132010-2003010112121031-0032212210333232-3013120002003013-1122032300112200-3010100310111222-0332013021310231"></a>

## azure_event_hubs_receiver — azure_event_hubs_receiver / 213310020012 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- azure_event_hubs_receiver

<a id="canonical-2320200132011030-1231332200110211-0002331211012320-2213001200201030-1011331311001011-1022302212331332-0103310102303100-2031212300120002"></a>

Type: `"object"`. single nested block, Optional.

Azure Event Hubs Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("instance",
    "namespace")}
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
azure_event_hubs_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021210033020122-0333310002110220-1110201213200212-3333120033331233-2201103103030331-1103103331202210-3320110200331313-2300301032311000"></a>

## Direct properties — azure_event_hubs_receiver / 213310020012 / 3

- [connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122): complete subsection reference.

<a id="canonical-3030111010023323-2101332000131212-0000222320202313-3320301123031300-3320133003332330-2112330230210133-0112331221321220-0022020210002022"></a>

<a id="canonical-0331213303103222-3201130132320213-2013332200101201-1100210302032310-2320013022130011-1223121221203112-0232013333210031-0111032200323322"></a>

## instance property — azure_event_hubs_receiver / 213310020012 / 4

Type: `"string"`. Optional.

Event Hubs Instance name into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-2012113103313300-3310113010310322-2333213210112300-2000233330302210-3221022001030331-2220110121020230-1012112022332032-1023012012032022"></a>

<a id="canonical-1300310122323121-1232233303333131-0302211113102202-3120012002332332-2002201300003021-0202033213013032-0211203013200013-2123101120133320"></a>

## namespace property — azure_event_hubs_receiver / 213310020012 / 5

Type: `"string"`. Optional, Computed.

Event Hubs Namespace is namespace with instance into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$",
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-1222312202112213-1000320301220323-1112232211022021-1101321103330102-1112203131033110-0313133001220321-2312201020031102-3312103320321223"></a>

## Next pages — azure_event_hubs_receiver / 213310020012 / 6

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322032203010330-3022032122330031-1333122330111201-0102100000233031-2111100323201322-1123310221221121-3033001222320021-1023302010321010"></a>

## azure_event_hubs_receiver.connection_string — connection_string / 033002102310 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312)
- azure_event_hubs_receiver.connection_string

<a id="canonical-0303101323120132-0101132321103131-1021223210313011-1331001302303103-3300311021201320-3311000032213030-3233103323001111-1211023231111323"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
connection_string {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022331321302302-0010013321112232-2310220323212000-2030102302220312-0310232021113112-3200330030033000-0130032130232301-0311102303033032"></a>

## Direct properties — connection_string / 033002102310 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-0020232312302013-1221212321022233-3312001230202132-1101002101322121-1301131003110231-0033022212102123-1010103121111211-1311333220233203): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-1300113323032120-0333312131102112-1010332031011032-3013023112020213-0002021302333010-1120132103200100-3113030100211013-2333212022001323): complete subsection reference.

<a id="canonical-1330212202201010-1203322132101021-0033220000232233-0303022322221033-3023331322132120-0203322102012202-2003023133122230-0310300223023112"></a>

## Next pages — connection_string / 033002102310 / 4

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-0020232312302013-1221212321022233-3312001230202132-1101002101322121-1301131003110231-0033022212102123-1010103121111211-1311333220233203)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-1300113323032120-0333312131102112-1010332031011032-3013023112020213-0002021302333010-1120132103200100-3113030100211013-2333212022001323)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0020232312302013-1221212321022233-3312001230202132-1101002101322121-1301131003110231-0033022212102123-1010103121111211-1311333220233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012122220310003-0023103012221112-3112110221103200-1302003331001233-3100012301103021-3033332110023113-3001223120113332-2233211131220202"></a>

## azure_event_hubs_receiver.connection_string.blindfold_secret_info — blindfold_secret_info / 020132020200 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312)
- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122)
- azure_event_hubs_receiver.connection_string.blindfold_secret_info

<a id="canonical-2001122110210031-1313230203003331-3223300000033231-1231121233011310-0203112231102020-3032011110203012-2231023012331032-0213010123330031"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1002032331120010-3320130021331033-3002113122222312-3133211310020003-3200213033030232-1333212100322133-2321131101301210-1231131033111123"></a>

## Direct properties — blindfold_secret_info / 020132020200 / 3

<a id="canonical-2231102311232100-3132200132012302-3332000101001130-3002021022130133-0200002003333102-1333211310203010-2232033122323123-1312023321030103"></a>

<a id="canonical-1010011203323212-2112030310101133-3303020210013132-3222230132032130-0333232303322123-3200203311220030-3331231103313110-1110001030223223"></a>

## decryption_provider property — blindfold_secret_info / 020132020200 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0220331113010130-3022003111122202-2130111233102323-0313230321013031-1030231113211301-3013112032103303-0221113020311023-0332323133310211"></a>

<a id="canonical-0030111011223231-3233133233001103-3321302230130333-0022220231022133-2321210200033212-1231103010103121-0323033022012333-1001311311220113"></a>

## location property — blindfold_secret_info / 020132020200 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1030212301213033-0012021230303111-2201100231332123-3013331231211012-3233220003200001-1313021331221211-3100202301203110-3100311201232213"></a>

<a id="canonical-2210303303133110-1030031231230002-0023312300100031-1013323133021023-1233223321311232-2223313301203003-1131333103120210-0002013311233010"></a>

## store_provider property — blindfold_secret_info / 020132020200 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0221003221221311-1010110200122302-0131330212110132-3111320002110332-1313110012233103-3110132330232223-1312100020030131-3233100012220023"></a>

## Next pages — blindfold_secret_info / 020132020200 / 7

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1300113323032120-0333312131102112-1010332031011032-3013023112020213-0002021302333010-1120132103200100-3113030100211013-2333212022001323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302122113022330-3200001312323313-2320002103200120-1000132332311121-3031303031002031-0323230330311210-2023123000222310-2213011323320012"></a>

## azure_event_hubs_receiver.connection_string.clear_secret_info — clear_secret_info / 321013211123 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0301232302221300-2020110021320100-3011312123300033-3310111120221330-3100213022012100-1133331002003131-3103013230201023-3223221130133312)
- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122)
- azure_event_hubs_receiver.connection_string.clear_secret_info

<a id="canonical-1133030010012111-0011322330102322-3003002322030020-1012030002321123-1023221122103000-2103212120320220-0103303013211310-2020230010031020"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1033211333101323-0010020003221120-2201110121303302-2021023010103001-0220232222000201-0100313130210120-3303011001213332-0311202223013313"></a>

## Direct properties — clear_secret_info / 321013211123 / 3

<a id="canonical-3313010020223102-3223111210323322-2020230233330112-0201121332132302-0010022132021332-2323210311123312-2313330200200120-2223211323101020"></a>

<a id="canonical-0211310303230301-0332112233310001-0203131002102022-1102302302002211-0100303332001330-2102111001130131-1022120011101033-1013030033001303"></a>

## provider_ref property — clear_secret_info / 321013211123 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0313031303323300-1301312310300113-2302200332321100-0011211311302031-1112023023312322-1132320323203103-0202321133032212-1332231321221303"></a>

<a id="canonical-2120022110332331-2003333020020012-3012013311032020-1110323303203113-0022202001112123-1002020333300101-0310230122033221-1313030012310211"></a>

## URL property — clear_secret_info / 321013211123 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1111131220220221-2110022230222200-3310130201123112-0101223210332222-0013202033202313-0132320010330231-0211012300313212-0111210203101203"></a>

## Next pages — clear_secret_info / 321013211123 / 6

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131132213023332-1112022230033000-0202311212230201-3000121030002030-2321323122020210-2203033222103311-3123320002133212-1201112103011102"></a>

## azure_receiver — azure_receiver / 213112032233 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- azure_receiver

<a id="canonical-3013200231333020-1303012020320320-1001100130303121-3211010010010303-2320110321121200-1102123310122230-2100321311330302-1302111303311231"></a>

Type: `"object"`. single nested block, Optional.

Azure Blob Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("container_name")}
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
azure_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003003121132220-0212322200330233-1033030020202202-2201231021111010-1132203332313132-2110203112101001-0032002212312020-1130210121213302"></a>

## Direct properties — azure_receiver / 213112032233 / 3

- [batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210): complete subsection reference.

- [connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131): complete subsection reference.

<a id="canonical-2303103313121301-2313301303211130-3302300023031203-3212330330121000-1220101333330100-3331012323131121-3302233312132020-1302211023033332"></a>

<a id="canonical-1021233210310331-3102312010112113-1132032303000223-3121000112203301-3033133002330320-1022320102103110-3302211211010213-3033103312200003"></a>

## container_name property — azure_receiver / 213112032233 / 4

Type: `"string"`. Optional.

Container Name is the name of the container into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

- [filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001): complete subsection reference.

<a id="canonical-3020321100022132-0022030320223332-0101101122033003-1312100232030003-3231130330323133-2211030301020112-0130001200020102-1311122111202021"></a>

## Next pages — azure_receiver / 213112032233 / 5

- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110132201000022-3222220222121330-1003201202211201-1212111121111213-0031221223010333-0312300010121212-3020221002320020-2023013030130131"></a>

## azure_receiver.batch — batch / 322202131020 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- azure_receiver.batch

<a id="canonical-0012020231111121-0203112023302002-3201223223301221-0110122201023320-0033311132200101-3121223302013232-1322223022210222-0003231110333030"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```
