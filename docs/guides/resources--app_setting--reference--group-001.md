---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103001322310113-1222310332200031-2232122313002121-1011310333301302-2003001030202030-1330010222212113-1311232332201331-3300122302123023"></a>

## Property reference — Property reference / 001031130333 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- Property reference

<a id="canonical-3203222223210220-1321210330130323-0322231002332110-3212033230113232-2232003331013100-3110020110223012-2320232220211321-3213212133322003"></a>

## Direct properties — Property reference / 001031130333 / 3

<a id="canonical-0303201112010032-0130033111212301-3221333312200100-3002020101323021-3033202030313330-1021212023321032-2321120121000330-3122123300120133"></a>

<a id="canonical-3103010001321023-1123010112122120-3023121132322022-3021301332333212-3323102030313121-2332111320331211-1203012233321021-0002012221003313"></a>

## annotations property — Property reference / 001031130333 / 4

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

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112): complete subsection reference.

<a id="canonical-0022030103333011-1103131120023100-0013122200311011-2302021011212333-1031101203203122-0230131100333132-1312120111312121-0232323102122213"></a>

<a id="canonical-3333111001100012-1322322220101030-2202002222333120-2022231320031013-2122233320010033-1023322300223011-1223223033033311-1101123011320101"></a>

## description property — Property reference / 001031130333 / 5

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

<a id="canonical-3121001130200322-1000333101100110-2101112000012100-2202113011032002-0113302313303321-3320333030112312-1000222201012000-3113122102232210"></a>

<a id="canonical-2111211110003120-1331320032022331-1321011231312021-3203201321012333-3210320122321223-3331100212310012-2101221121100033-0201321303312031"></a>

## disable property — Property reference / 001031130333 / 6

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

<a id="canonical-0300101311300122-3300212201013103-1203120120233021-2321203302331002-2031123100321122-2120123110210001-2200202302333011-0102133132012313"></a>

<a id="canonical-1132331031323323-1301232301332312-2113203223203231-3303323213023011-2023132333021200-3333223111321100-3312211111120210-1101303023021003"></a>

## ID property — Property reference / 001031130333 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3120311110320202-3222133302121023-0103232023021121-1301003223003212-2232020133303310-1021120232212210-2232330101100102-2023332102003333"></a>

<a id="canonical-2022111313002112-0322121033221123-3120210100010330-0123030100130123-1112102123110013-2213021311302311-3321313230102023-3123002323122322"></a>

## labels property — Property reference / 001031130333 / 8

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

<a id="canonical-0000300203012300-1313132031011220-0031113023211212-1120333330001213-3133102320101320-2112033133230113-3321220023211300-2302111000130131"></a>

<a id="canonical-2120111000302022-3103232203120112-3303222121210312-2223231110313210-2030312220210201-3203333202100220-0312103013130032-2232310033220033"></a>

## name property — Property reference / 001031130333 / 9

Type: `"string"`. Required.

Name of the App Setting. Must be unique within the namespace.

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

<a id="canonical-2131211130112020-1032331010231300-1222100030130133-0301200230133320-1023303031133131-1303103120302020-1113103210123130-3203221202033311"></a>

<a id="canonical-3212232231032313-3021003123011231-0333322322213120-0302130213000222-3131221122010322-2333311212312010-0123203331221300-0102100231033012"></a>

## namespace property — Property reference / 001031130333 / 10

Type: `"string"`. Required.

Namespace where the App Setting is created.

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

- [timeouts](resources--app_setting--reference--group-001.md#canonical-2312233300001223-3232022233011210-1010223101311011-0232002130332010-0102020202112232-0333003300011211-1233030122120233-1122012301112032): complete subsection reference.

<a id="canonical-1002230023302313-1103032220113112-0300031231320312-3310111022310122-1310221310213320-1113112210201220-1200122110121112-0312113131101003"></a>

## All schema paths — Property reference / 001031130333 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_setting--reference--group-001.md#canonical-0303201112010032-0130033111212301-3221333312200100-3002020101323021-3033202030313330-1021212023321032-2321120121000330-3122123300120133) |
| `app_type_settings` | [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1022101331322312-0121111332022031-3200100022221311-3321132320321103-0231120332110310-0013320223203033-1131021130121301-0101130102121203) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](resources--app_setting--reference--group-001.md#canonical-2302303212222310-3130303222330222-0301131232301121-2002220010203131-3211033122112120-0332203232212131-0113001311131112-1113220102220010) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](resources--app_setting--reference--group-001.md#canonical-0312010200032200-3301031213200101-1001223332313232-3022220123220200-3222302031302313-2320002201212203-0000322302123311-0313310120003133) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](resources--app_setting--reference--group-001.md#canonical-3123000302102330-1132203233302330-3103222131001220-3100203220103322-2303021222002210-0100211300131000-0233003310231203-2000322300003131) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](resources--app_setting--reference--group-001.md#canonical-3303112110313133-3232122110302011-0311103212300021-0211212112111102-1022012111012212-3003330211030213-1201221121011312-1001301111021300) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](resources--app_setting--reference--group-001.md#canonical-3000033103021023-3120203002120202-3112310221300320-1022211101323310-0002303130122310-2000021001321312-3022303131313110-3320131010303130) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](resources--app_setting--reference--group-001.md#canonical-0001313312012331-1130312012223303-2301021121122313-1121202223132103-1111301321221003-2121311210213020-1311033030320311-0202212010123123) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-2312103220210131-3112220313013230-3122121122221133-2121310031233321-1313130033011021-1123002100201213-1121233010220013-0313321311213301) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](resources--app_setting--reference--group-001.md#canonical-0320030202033003-2210003023203220-3330132003322022-1122010210213221-2131000211323120-0313033331302323-0310300112030211-3111033113001312) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](resources--app_setting--reference--group-001.md#canonical-0233302011231203-3201101221012200-1122323332202333-2233110033112202-2032210221331312-3132210022230322-1220100200301013-0222231111022201) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2211000132221302-2023211332112103-0020311311310330-0200223033312001-3110313133030220-2210111203322012-0203131313231102-1012302322313002) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--reference--group-001.md#canonical-2213012330333030-2113102310211123-0301121031311002-3121212210001103-2310012133330020-3103221301301300-1021010332313033-0131100211120310) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](resources--app_setting--reference--group-001.md#canonical-0301103010320221-3231132300122212-1021323003013113-0012130300202211-0330001230211321-3221212022321131-1021301133033323-2203130122113113) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](resources--app_setting--reference--group-001.md#canonical-0212031023333031-1200013211131000-0033223013123201-2003201021322303-1023310311120032-3122112021130002-3123312222022232-2023030131131132) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2221112321313133-2300211020210030-2133213110233103-1200320232213102-2003202312302213-0322313131310122-0111230010200230-0112212312003322) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](resources--app_setting--reference--group-001.md#canonical-3312103330121211-1121031103030030-2000011211201032-0023320120020213-1323323322202330-0213030033120213-1323101113322231-2321112333232133) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](resources--app_setting--reference--group-001.md#canonical-1130100232002131-3302203001022301-3332213003111132-3033030231100230-1130002221310033-1113320300100021-0202123010023321-2133032313203022) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1000003032102310-0030031303033221-0102011320212311-1331322222233031-1003023102220100-0303233313000200-1010222221003111-1311120203130201) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-1022032112020210-2121313211000312-0321022122131013-1230130312301312-0102331123223302-2323220111122312-0330130300213123-0103013021100012) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](resources--app_setting--reference--group-001.md#canonical-1201232013121300-1021302133211123-2013003331301221-0111022311222112-2233111120100011-3312033213221001-1331233020113113-2232233022303232) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-3103321212320100-3111111130100313-0320313221200310-1030222103231120-3330111133212002-1320122330223022-1203032231320001-2302233120300202) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-3333033130113320-3000223002133131-1322131313313031-2102013021013131-1010202332133221-2200112222203001-0021310102213132-3102232333312001) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-2322321123112232-3302022113132321-3011010210120122-2212320332320011-3012223311222011-1201110332333312-3130000020232321-0111033022033311) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-0011112222323200-0103323210133020-3122211102302213-1113302312330112-3022122312023023-1112220213000320-0021310323012102-1030210202132211) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-1210221211222122-1323233322023121-3221232111020113-1033030310112230-3202030223211020-0232302300231031-3210020132012233-0301003032233012) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-2021021302213211-1232033122113222-3201033120122030-3101130032023132-2101120101200112-1222031113233021-2100123012130000-0101322130313130) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-1330202131313302-1332312212230111-3133133011201210-2330202200201223-3331110221033313-2000312330120001-1001222203010230-1220230220000302) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-2332110131202222-2111002303111121-0000331013322322-2331312233121132-0203220101330130-0131131110001220-2330130022300320-1300012110120012) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-3132101101300012-0032311103102122-3331012111220112-3230203201030010-1110101100231001-3312300121303113-3332200203221032-3211332020233212) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-0122231033031000-0321131233311023-3320100221322303-0030331222301123-1312300233001030-0203300312103212-3031022103102323-0031001010112122) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](resources--app_setting--reference--group-001.md#canonical-3132201022212123-0000113031022311-2200133010021110-0321202112220331-1131113022002233-0330321003111110-2312030300100302-1030112102022300) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-3021212220223230-0312130011122113-3213100010112330-1211212323233213-3230220222331033-2333033021333301-2311221323120310-1222031222322030) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](resources--app_setting--reference--group-001.md#canonical-0103213103121333-0003123000103322-1322303021333232-0100212323200220-1201031320032010-1013232330223230-0231111000333332-2100010132010202) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-2032310031102112-0111021310033321-1212310022130031-1102121102133300-2222321102300110-3212210001210011-1201302031130203-1001000101313233) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-3130221113120330-1110301213110012-0333110110300300-1332231103201112-3301011013031303-0132321200020202-2102033333113233-3012012022002133) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](resources--app_setting--reference--group-001.md#canonical-2330332321200000-2221100213221313-3132230012323101-3232201120120030-0020221022102031-3023313202002303-2120213132123110-0331032133012012) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](resources--app_setting--reference--group-001.md#canonical-0023223322112312-3112123130112121-3232331322300131-3222220222322103-3023323032131101-2210000003000212-3120031001331131-2322111122302120) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](resources--app_setting--reference--group-001.md#canonical-2002022310120023-2121123232203133-2203131000102113-0113222223122313-2133133222002110-0230033021103010-1003131322030223-3232201003131301) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-2003330031313320-3320233122231100-0311222132030230-0232221021222203-2131230203013123-0103002121302002-2103202233322111-1210112202132101) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](resources--app_setting--reference--group-001.md#canonical-2003120012102003-0233113302210233-2130112031023303-3121103012110203-1232100212221123-0312011010313112-0012113332212323-0000010010133332) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](resources--app_setting--reference--group-001.md#canonical-2111030211300312-0011002233113012-3310210102232323-2110012230132312-2321302010331021-2310201121201100-3012211010011222-3303322330332313) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](resources--app_setting--reference--group-001.md#canonical-3020333331220203-2330321331212011-3202033030001221-3021122212032113-1102013222000132-1032301133220010-1122220111130200-0210033232233310) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](resources--app_setting--reference--group-001.md#canonical-3100131221210102-0300111111020311-2001110212222021-0310032201303003-0313022331322000-1212331032120301-0210310032131300-0123303011300103) |
| `description` | [description](resources--app_setting--reference--group-001.md#canonical-0022030103333011-1103131120023100-0013122200311011-2302021011212333-1031101203203122-0230131100333132-1312120111312121-0232323102122213) |
| `disable` | [disable](resources--app_setting--reference--group-001.md#canonical-3121001130200322-1000333101100110-2101112000012100-2202113011032002-0113302313303321-3320333030112312-1000222201012000-3113122102232210) |
| `id` | [ID](resources--app_setting--reference--group-001.md#canonical-0300101311300122-3300212201013103-1203120120233021-2321203302331002-2031123100321122-2120123110210001-2200202302333011-0102133132012313) |
| `labels` | [labels](resources--app_setting--reference--group-001.md#canonical-3120311110320202-3222133302121023-0103232023021121-1301003223003212-2232020133303310-1021120232212210-2232330101100102-2023332102003333) |
| `name` | [name](resources--app_setting--reference--group-001.md#canonical-0000300203012300-1313132031011220-0031113023211212-1120333330001213-3133102320101320-2112033133230113-3321220023211300-2302111000130131) |
| `namespace` | [namespace](resources--app_setting--reference--group-001.md#canonical-2131211130112020-1032331010231300-1222100030130133-0301200230133320-1023303031133131-1303103120302020-1113103210123130-3203221202033311) |
| `timeouts` | [timeouts](resources--app_setting--reference--group-001.md#canonical-3020132020223121-0222232030021110-2012231100012322-2230233130312132-1011310210200130-0230203330201102-0202012332312032-0323000103123330) |
| `timeouts.create` | [timeouts.create](resources--app_setting--reference--group-001.md#canonical-1301133022133031-3101113123011122-1213010231330230-0211012200300122-0210233020131200-0331022131113211-1123120323333221-3302331103030000) |
| `timeouts.delete` | [timeouts.delete](resources--app_setting--reference--group-001.md#canonical-2230112213303312-1033101212023132-3233131122103202-2102303300223023-3202301200202000-1020100121012210-3003131313001300-3320030033033113) |
| `timeouts.read` | [timeouts.read](resources--app_setting--reference--group-001.md#canonical-1333223232122332-0123230223121211-1201201303320312-1231032112013131-1302221300230011-2210233333131102-0202112001030301-1203000130213320) |
| `timeouts.update` | [timeouts.update](resources--app_setting--reference--group-001.md#canonical-0303033330121032-3221202130112132-2121223230000323-3220220223121131-2133110112333202-3130122101220321-3300101311112002-0120300221320202) |

<a id="canonical-1032333102102031-1001213311213122-3120013021001003-0203110031210023-3002303213103122-0122121223231020-3213201003303020-0211012213010020"></a>

## Next pages — Property reference / 001031130333 / 12

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [timeouts](resources--app_setting--reference--group-001.md#canonical-2312233300001223-3232022233011210-1010223101311011-0232002130332010-0102020202112232-0333003300011211-1233030122120233-1122012301112032)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023031003330210-3221021101101303-3321032301332020-2012111230011010-1122312221032233-3202211310313200-3331232120030102-3000233211220131"></a>

## app_type_settings — app_type_settings / 213300023220 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- app_type_settings

<a id="canonical-1022101331322312-0121111332022031-3200100022221311-3321132320321103-0231120332110310-0013320223203033-1131021130121301-0101130102121203"></a>

Type: `"object"`. list nested block, Optional.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("app_type_ref")}
```

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
app_type_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020103021003031-2232133210320313-3100110123201012-2330233110300202-0113013131310231-0311221312011203-3011222222211233-0330230231110222"></a>

## Direct properties — app_type_settings / 213300023220 / 3

- [app_type_ref](resources--app_setting--reference--group-001.md#canonical-0031103021102102-3301300233201013-3333223113310100-2203322101303330-0111303213312001-3011103102123213-2103130311000303-2223003002000211): complete subsection reference.

- [business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032): complete subsection reference.

- [timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020): complete subsection reference.

- [user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012): complete subsection reference.

<a id="canonical-2301212000301132-1321322122221000-0221313212231011-3223202220002202-1213132031330230-2202132332101000-2223223331133023-3211312003301102"></a>

## Next pages — app_type_settings / 213300023220 / 4

- [app_type_settings.app_type_ref](resources--app_setting--reference--group-001.md#canonical-0031103021102102-3301300233201013-3333223113310100-2203322101303330-0111303213312001-3011103102123213-2103130311000303-2223003002000211)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-0031103021102102-3301300233201013-3333223113310100-2203322101303330-0111303213312001-3011103102123213-2103130311000303-2223003002000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131120122013001-3300110220222101-1221111233100203-2333203333032120-0311012023200222-2223122023002330-3100211300302200-3122021021213020"></a>

## app_type_settings.app_type_ref — app_type_ref / 313321332211 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.app_type_ref

<a id="canonical-2302303212222310-3130303222330222-0301131232301121-2002220010203131-3211033122112120-0332203232212131-0113001311131112-1113220102220010"></a>

Type: `"object"`. list nested block, Optional.

The AppType of App instance in current Namespace. Associating an AppType reference, will enable
analysis on this instance's generated data.

Upstream description:

The AppType of App instance in current Namespace. Associating an AppType reference, will enable
analysis on this instance's generated data.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_type_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331332100303001-2100310003233123-1010130333322012-2211333122211330-1301210210200212-2200021011010331-1033231232332013-1220123022022001"></a>

## Direct properties — app_type_ref / 313321332211 / 3

<a id="canonical-0312010200032200-3301031213200101-1001223332313232-3022220123220200-3222302031302313-2320002201212203-0000322302123311-0313310120003133"></a>

<a id="canonical-2210231213313210-0330003033013332-0010131133130232-1110322222130323-0023211100113320-3331220132300010-1212333302211320-2132320210100130"></a>

## kind property — app_type_ref / 313321332211 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3123000302102330-1132203233302330-3103222131001220-3100203220103322-2303021222002210-0100211300131000-0233003310231203-2000322300003131"></a>

<a id="canonical-2130232031231103-2030000233203020-0211220233023323-3300322301102201-1222120233321130-1321111010032020-0310212123022112-3200301003002333"></a>

## name property — app_type_ref / 313321332211 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3303112110313133-3232122110302011-0311103212300021-0211212112111102-1022012111012212-3003330211030213-1201221121011312-1001301111021300"></a>

<a id="canonical-1132230301213330-2310130132202330-2223002112020112-3220310232022103-0322010100003010-2201022113103323-2133212302213133-1132232033111022"></a>

## namespace property — app_type_ref / 313321332211 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3000033103021023-3120203002120202-3112310221300320-1022211101323310-0002303130122310-2000021001321312-3022303131313110-3320131010303130"></a>

<a id="canonical-0023233310300323-2113220201330033-2100112011113333-0020320302223132-1111130101130113-2122003333301022-0213032222223221-1122203213310232"></a>

## tenant property — app_type_ref / 313321332211 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0001313312012331-1130312012223303-2301021121122313-1121202223132103-1111301321221003-2121311210213020-1311033030320311-0202212010123123"></a>

<a id="canonical-1310322020110112-0230103032003102-2110012021332320-0032320231321202-3232001030020032-1210012100233231-3201331131121012-2121323111330321"></a>

## uid property — app_type_ref / 313321332211 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2310220201320133-3110200321112021-1311102100322320-3112211300330310-3301332112023021-2032100020122222-1220030311313333-3330132133031230"></a>

## Next pages — app_type_ref / 313321332211 / 9

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213002223133031-2123102310333100-3001310221010121-1020120131212330-2203013212233000-0203101013333322-0130233132210103-1031232230101233"></a>

## app_type_settings.business_logic_markup_setting — business_logic_markup_setting / 301013301321 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.business_logic_markup_setting

<a id="canonical-2312103220210131-3112220313013230-3122121122221133-2121310031233321-1313130033011021-1123002100201213-1121233010220013-0313321311213301"></a>

Type: `"object"`. single nested block, Optional.

Settings specifying how API Discovery will be performed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-learn_from_namespace": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310311212031223-0001132103131302-1301000123113312-0223000302332133-0321113210103223-0012221221323310-3100321030000302-0322011200120313"></a>

## Direct properties — business_logic_markup_setting / 301013301321 / 3

- [disable_spec](resources--app_setting--reference--group-001.md#canonical-2103102313033100-2102232331010211-0001221033310303-2013223003002020-0202130220122002-1310121331003301-0232111331100202-0201012300110131): complete subsection reference.

- [enable](resources--app_setting--reference--group-001.md#canonical-1313220311133123-2232201103020022-2333020322311131-1220310013331221-2312103201311021-2110033002100310-2200313200113031-3131313102113203): complete subsection reference.

<a id="canonical-0303202022200103-0312031313130110-3323130002113203-2130313100332222-0030230031113211-1312001233211133-2203222212231033-2110000232222323"></a>

## Next pages — business_logic_markup_setting / 301013301321 / 4

- [app_type_settings.business_logic_markup_setting.disable_spec](resources--app_setting--reference--group-001.md#canonical-2103102313033100-2102232331010211-0001221033310303-2013223003002020-0202130220122002-1310121331003301-0232111331100202-0201012300110131)
- [app_type_settings.business_logic_markup_setting.enable](resources--app_setting--reference--group-001.md#canonical-1313220311133123-2232201103020022-2333020322311131-1220310013331221-2312103201311021-2110033002100310-2200313200113031-3131313102113203)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2103102313033100-2102232331010211-0001221033310303-2013223003002020-0202130220122002-1310121331003301-0232111331100202-0201012300110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323111030203302-2310221331312201-2102233200133001-0313021103122113-0311030113313122-2313122300333021-3311312221303100-0300012332121031"></a>

## app_type_settings.business_logic_markup_setting.disable_spec — disable_spec / 221333203322 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- app_type_settings.business_logic_markup_setting.disable_spec

<a id="canonical-0320030202033003-2210003023203220-3330132003322022-1122010210213221-2131000211323120-0313033331302323-0310300112030211-3111033113001312"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-2300020012230112-2022131222221022-0211200312323220-1303320103130203-2032221320130331-2030201321103212-0301131110111332-0111010011022323"></a>

## Direct properties — disable_spec / 221333203322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233031113100230-2313211203122001-3032221201122010-2313111023223231-0012012312111012-1210311021220220-3110201200113130-0113130232201013"></a>

## Next pages — disable_spec / 221333203322 / 4

- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1313220311133123-2232201103020022-2333020322311131-1220310013331221-2312103201311021-2110033002100310-2200313200113031-3131313102113203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323103021102121-2331310322000232-2333323131113001-2033100122133333-1302002022022122-2321033010211103-2201030131101100-3232021213231320"></a>

## app_type_settings.business_logic_markup_setting.enable — enable / 111310223020 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-0233302011231203-3201101221012200-1122323332202333-2233110033112202-2032210221331312-3132210022230322-1220100200301013-0222231111022201"></a>

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
enable = {}
```

<a id="canonical-1113103132001233-3300231230231222-0310221101220103-1031331033013101-3021220022020102-1200311233021020-2221321030111103-1233233213231012"></a>

## Direct properties — enable / 111310223020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112210233130310-2200110313321320-3311201121131103-1323321031121101-2100103323020001-0211023311333100-3203223111230303-1310000002011131"></a>

## Next pages — enable / 111310223020 / 4

- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032321322001131-1333030130002103-0121200032123203-0310033021211230-3102233001102131-2321123102012312-2333130301013113-0323202031331230"></a>

## app_type_settings.timeseries_analyses_setting — timeseries_analyses_setting / 001111222011 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-2211000132221302-2023211332112103-0020311311310330-0200223033312001-3110313133030220-2210111203322012-0203131313231102-1012302322313002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for timeseries analyses setting.

Upstream description:

Configuration for DDoS Detection.

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
timeseries_analyses_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303033321301111-0302123231103133-3220201012032201-1220323201211220-1321320320201102-0201300212001002-1132132022313020-3323323211013222"></a>

## Direct properties — timeseries_analyses_setting / 001111222011 / 3

- [metric_selectors](resources--app_setting--reference--group-001.md#canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123): complete subsection reference.

<a id="canonical-1232012323212002-2130110022321322-1231121320032212-1310201202022010-3133211322212101-3111211032221031-0323032331101332-3230002203100223"></a>

## Next pages — timeseries_analyses_setting / 001111222011 / 4

- [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--reference--group-001.md#canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221110120000300-3023200021203323-1020021333210311-1233003121323011-3333133011211020-3220130113330030-3230001302211303-0023203213111013"></a>

## app_type_settings.timeseries_analyses_setting.metric_selectors — metric_selectors / 332203002233 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020)
- app_type_settings.timeseries_analyses_setting.metric_selectors

<a id="canonical-2213012330333030-2113102310211123-0301121031311002-3121212210001103-2310012133330020-3103221301301300-1021010332313033-0131100211120310"></a>

Type: `"object"`. list nested block, Optional.

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

Upstream description:

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
metric_selectors {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300001010332022-1120020113033103-3311120303122233-1212011333203202-2112322323022211-0132303311021101-1232031202011122-0121203213301131"></a>

## Direct properties — metric_selectors / 332203002233 / 3

<a id="canonical-0301103010320221-3231132300122212-1021323003013113-0012130300202211-0330001230211321-3221212022321131-1021301133033323-2203130122113113"></a>

<a id="canonical-1200212121220213-0310232023330202-2132222000202121-0130210322203231-0221100120230201-1322023112332022-2111001132212321-1332203311310333"></a>

## metric property — metric_selectors / 332203002233 / 4

Type: `["list", "string"]`. Optional.

\[Enum: NO\_METRICS|REQUEST\_RATE|ERROR\_RATE|LATENCY|THROUGHPUT\] Choose one or more metrics to be
included in the detection logic. Possible values are \`NO\_METRICS\`, \`REQUEST\_RATE\`,
\`ERROR\_RATE\`, \`LATENCY\`, \`THROUGHPUT\`. Defaults to \`NO\_METRICS\`.

Upstream description:

Choose one or more metrics to be included in the detection logic.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0212031023333031-1200013211131000-0033223013123201-2003201021322303-1023310311120032-3122112021130002-3123312222022232-2023030131131132"></a>

<a id="canonical-2030201203023011-3222313210012112-2231333300030003-3130222201220330-3220000000201331-0223211200320011-1013320321222203-1211022322210023"></a>

## metrics_source property — metric_selectors / 332203002233 / 5

Type: `"string"`. Optional.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

Upstream description:

Supported sources from which Metrics can be analyzed

All edges in the service mesh graph. Metrics are analyzed separately between all source and
destination service combinations.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EDGES","NODES","NONE","VIRTUAL_HOSTS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122232003123023-0113033333120000-3001003222123332-2311002210112111-3012323011002323-2322033133100010-3022323112011102-1301212302321131"></a>

## Next pages — metric_selectors / 332203002233 / 6

- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011312130020313-2213000001131102-3010332310032220-0222212002213111-0210110133113130-2133100031302232-0231302323002121-0021003202231232"></a>

## app_type_settings.user_behavior_analysis_setting — user_behavior_analysis_setting / 021020311103 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.user_behavior_analysis_setting

<a id="canonical-2221112321313133-2300211020210030-2133213110233103-1200320232213102-2003202312302213-0322313131310122-0111230010200230-0112212312003322"></a>

Type: `"object"`. single nested block, Optional.

Configuration for user behavior analysis.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_detection",
    "enable_detection"),
  validators.ConflictingObjectAttributes("disable_learning",
    "enable_learning")}
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
  "x-ves-oneof-field-learn_from_namespace": "[\"disable_learning\",\"enable_learning\"]",
  "x-ves-oneof-field-malicious_user_detection": "[\"disable_detection\",\"enable_detection\"]"
}
```

Terraform syntax:

```terraform
user_behavior_analysis_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121032101032320-0101120011010100-2023322303031212-2010000012112123-3220110231221021-1120321011322133-2330003022111112-0233131033110101"></a>

## Direct properties — user_behavior_analysis_setting / 021020311103 / 3

- [disable_detection](resources--app_setting--reference--group-001.md#canonical-1113321302003221-2113300101202330-2030303232031111-2322310111302001-3300012220001130-3033333310133300-2322330130303300-0330210101132230): complete subsection reference.

- [disable_learning](resources--app_setting--reference--group-001.md#canonical-2133320120113332-2132323030231223-3331322212232313-3320132031030321-3122103030120331-1010200101323213-2102031323331202-0112111302022203): complete subsection reference.

- [enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013): complete subsection reference.

- [enable_learning](resources--app_setting--reference--group-001.md#canonical-3100010103133320-1031023303101233-1020203111023232-3321121123112321-2033103000311233-3202320323210321-2131232310323032-3003101030122022): complete subsection reference.

<a id="canonical-0221211203221220-2232320310022110-3220211202211100-1011201312110330-0333122021303223-1021120120301120-0011123030012210-2310323000333203"></a>

## Next pages — user_behavior_analysis_setting / 021020311103 / 4

- [app_type_settings.user_behavior_analysis_setting.disable_detection](resources--app_setting--reference--group-001.md#canonical-1113321302003221-2113300101202330-2030303232031111-2322310111302001-3300012220001130-3033333310133300-2322330130303300-0330210101132230)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](resources--app_setting--reference--group-001.md#canonical-2133320120113332-2132323030231223-3331322212232313-3320132031030321-3122103030120331-1010200101323213-2102031323331202-0112111302022203)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](resources--app_setting--reference--group-001.md#canonical-3100010103133320-1031023303101233-1020203111023232-3321121123112321-2033103000311233-3202320323210321-2131232310323032-3003101030122022)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1113321302003221-2113300101202330-2030303232031111-2322310111302001-3300012220001130-3033333310133300-2322330130303300-0330210101132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332023302201112-0230131331111320-3023030111022113-3231120222102322-1103112000311313-0132233122132111-2132021310332130-0113120110200033"></a>

## app_type_settings.user_behavior_analysis_setting.disable_detection — disable_detection / 332302300222 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-3312103330121211-1121031103030030-2000011211201032-0023320120020213-1323323322202330-0213030033120213-1323101113322231-2321112333232133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable detection.

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
disable_detection = {}
```

<a id="canonical-3013322121003230-3230220010003001-1322313010110200-2200303213322323-2112133122111330-3033102130223323-3233021123301031-1122222202010223"></a>

## Direct properties — disable_detection / 332302300222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302030310030023-1322123303230031-1112003010002333-1032010230030032-3303322123012011-0130312100331032-0122331020233031-0210222020310031"></a>

## Next pages — disable_detection / 332302300222 / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2133320120113332-2132323030231223-3331322212232313-3320132031030321-3122103030120331-1010200101323213-2102031323331202-0112111302022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211232123023003-1013310220102011-1201011312132110-0003203111323300-1003323202211213-1311302212023133-0232313001010033-2331013001000312"></a>

## app_type_settings.user_behavior_analysis_setting.disable_learning — disable_learning / 213331012301 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-1130100232002131-3302203001022301-3332213003111132-3033030231100230-1130002221310033-1113320300100021-0202123010023321-2133032313203022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learning.

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
disable_learning = {}
```

<a id="canonical-1101100000202032-2000033031223312-3213200103202120-0221323013222031-0002013120333223-2003212231310020-2023311213201031-0020111100030102"></a>

## Direct properties — disable_learning / 213331012301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011010123110202-1101133212133323-1302201031322303-3021322030232213-1111211123010023-1211021302210123-1132213223121123-1013100221230232"></a>

## Next pages — disable_learning / 213331012301 / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010132221123020-3213211130323302-0312211330200320-1112111311003231-3111100113201201-0322001103031320-0223203203212121-2101333110332213"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection — enable_detection / 301120030133 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.enable_detection

<a id="canonical-1000003032102310-0030031303033221-0102011320212311-1331322222233031-1003023102220100-0303233313000200-1010222221003111-1311120203130201"></a>

Type: `"object"`. single nested block, Optional.

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Upstream description:

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bola_detection_automatic",
    "exclude_bola_detection"),
  validators.ConflictingObjectAttributes("exclude_bot_defense_activity",
    "include_bot_defense_activity"),
  validators.ConflictingObjectAttributes("exclude_failed_login_activity",
    "include_failed_login_activity"),
  validators.ConflictingObjectAttributes("exclude_forbidden_activity",
    "include_forbidden_activity"),
  validators.ConflictingObjectAttributes("exclude_ip_reputation",
    "include_ip_reputation"),
  validators.ConflictingObjectAttributes("exclude_non_existent_url_activity",
    "include_non_existent_url_activity_automatic"),
  validators.ConflictingObjectAttributes("exclude_non_existent_url_activity",
    "include_non_existent_url_activity_custom"),
  validators.ConflictingObjectAttributes("exclude_rate_limit",
    "include_rate_limit"),
  validators.ConflictingObjectAttributes("exclude_waf_activity",
    "include_waf_activity"),
  validators.ConflictingObjectAttributes("include_non_existent_url_activity_automatic",
    "include_non_existent_url_activity_custom")}
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
  "x-ves-oneof-field-bola_activity_choice": "[\"bola_detection_automatic\",\"exclude_bola_detection\"]",
  "x-ves-oneof-field-bot_defense_activity_choice": "[\"exclude_bot_defense_activity\",\"include_bot_defense_activity\"]",
  "x-ves-oneof-field-cooling_off_period_setting": "[\"cooling_off_period\"]",
  "x-ves-oneof-field-failed_login_activity_choice": "[\"exclude_failed_login_activity\",\"include_failed_login_activity\"]",
  "x-ves-oneof-field-forbidden_activity_choice": "[\"exclude_forbidden_activity\",\"include_forbidden_activity\"]",
  "x-ves-oneof-field-ip_reputation_choice": "[\"exclude_ip_reputation\",\"include_ip_reputation\"]",
  "x-ves-oneof-field-non_existent_url_activity_choice": "[\"exclude_non_existent_url_activity\",\"include_non_existent_url_activity_automatic\",\"include_non_existent_url_activity_custom\"]",
  "x-ves-oneof-field-rate_limit_choice": "[\"exclude_rate_limit\",\"include_rate_limit\"]",
  "x-ves-oneof-field-waf_activity_choice": "[\"exclude_waf_activity\",\"include_waf_activity\"]"
}
```

Terraform syntax:

```terraform
enable_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031132130010133-0000103211013301-2100330223102033-3223332001231131-0230112231131222-0212320322223301-3302100203013332-1201010032122333"></a>

## Direct properties — enable_detection / 301120030133 / 3

- [bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-2030002320211330-3031023213120032-1210231322000111-0010301210211210-0111231213220102-1201113322232000-1102103101012232-0231212332003011): complete subsection reference.

<a id="canonical-1201232013121300-1021302133211123-2013003331301221-0111022311222112-2233111120100011-3312033213221001-1331233020113113-2232233022303232"></a>

<a id="canonical-1003323332203032-0332013120023122-0232100220333231-0001012300322031-3313022101231120-0210310330203301-1331231332131203-0123300301201032"></a>

## cooling_off_period property — enable_detection / 301120030133 / 4

Type: `"number"`. Optional.

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels..

Upstream description:

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels. This field specifies the time period, in minutes, used by the system to decay a user's
threat level from a high to medium or medium to low or low to none.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(5, 120),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 120,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  }
}
```

- [exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-1022101303111302-0031312232020020-2300201001201023-0003010330133112-2011023303111323-1113311031012231-0032111030101133-1202113310210203): complete subsection reference.

- [exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-3212222131011313-2012221300332321-0322233133000202-2023032223232022-1121000000020312-2010123002132123-3103031112003200-2023223020232002): complete subsection reference.

- [exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-3123100011311131-3012303320313111-0220332210310102-1120211213120131-2113301303233013-2132213223100023-3230213112221022-0300021132032031): complete subsection reference.

- [exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-1030233230101130-3101011302232020-1023200222133213-1103303220112213-2303221200220223-0020212202030203-0310322202322000-0312120123122311): complete subsection reference.

- [exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-3220213311333230-3023230111301321-0030232113111131-3020132103212311-0231232211220013-0200321110032020-1120301120313232-0300211321202303): complete subsection reference.

- [exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-3322023111302302-1201113202023331-1031223230133132-1301130022012201-0320031310212302-3213211122322031-1112223120130312-1220210332122010): complete subsection reference.

- [exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-2000121031220032-0232201010313321-0132310003122212-1122300023210111-3221210201223223-3013030300113001-3222211001332131-2003311122313032): complete subsection reference.

- [exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-1000302010301201-1202112323003110-0123033202332102-1100032003011131-1302333312310023-2112102133233020-2211010302312230-1000223033312130): complete subsection reference.

- [include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-2331333103301032-2111022012032100-2011300130000230-2303021301301220-1200303320322030-2110020223323003-3203203333000021-0320110220313323): complete subsection reference.

- [include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-1211100201310202-3023000003233301-0231313213132232-2333311031330110-1322020233020000-0313120013202310-2313302232013033-0020012232333123): complete subsection reference.

- [include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-1121232011330112-3032012310222130-3332310022303323-2003302020201303-0202221021100123-0200112032023313-3203111302100302-0011332203013210): complete subsection reference.

- [include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-0200212113013321-0010302132130032-2002230233010100-0013122033122022-1101230120020032-3232303113331303-3322212303122100-1022321321123323): complete subsection reference.

- [include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101): complete subsection reference.

- [include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-3033220110323013-2323013003303301-1033231123310221-0132003332000112-3223301221302121-3321102103000233-3133131310032101-0123333321232123): complete subsection reference.

- [include_rate_limit](resources--app_setting--reference--group-001.md#canonical-3202003232100302-3113132033010233-1321100232302013-2220213133221221-2310213332333030-2133321132322100-3331103022312200-0120222001013000): complete subsection reference.

- [include_waf_activity](resources--app_setting--reference--group-001.md#canonical-2321311012200003-3323332300100232-0110223212200132-0220232021222023-3133212033112223-3102230303211101-0031202030213333-1213200213303111): complete subsection reference.

<a id="canonical-1232110222230232-3210031003113123-3023030011333311-2103220310030232-2131310000201233-3031321212232201-3020003121232312-2201301201032002"></a>

## Next pages — enable_detection / 301120030133 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-2030002320211330-3031023213120032-1210231322000111-0010301210211210-0111231213220102-1201113322232000-1102103101012232-0231212332003011)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-1022101303111302-0031312232020020-2300201001201023-0003010330133112-2011023303111323-1113311031012231-0032111030101133-1202113310210203)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-3212222131011313-2012221300332321-0322233133000202-2023032223232022-1121000000020312-2010123002132123-3103031112003200-2023223020232002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-3123100011311131-3012303320313111-0220332210310102-1120211213120131-2113301303233013-2132213223100023-3230213112221022-0300021132032031)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-1030233230101130-3101011302232020-1023200222133213-1103303220112213-2303221200220223-0020212202030203-0310322202322000-0312120123122311)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-3220213311333230-3023230111301321-0030232113111131-3020132103212311-0231232211220013-0200321110032020-1120301120313232-0300211321202303)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-3322023111302302-1201113202023331-1031223230133132-1301130022012201-0320031310212302-3213211122322031-1112223120130312-1220210332122010)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-2000121031220032-0232201010313321-0132310003122212-1122300023210111-3221210201223223-3013030300113001-3222211001332131-2003311122313032)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-1000302010301201-1202112323003110-0123033202332102-1100032003011131-1302333312310023-2112102133233020-2211010302312230-1000223033312130)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-2331333103301032-2111022012032100-2011300130000230-2303021301301220-1200303320322030-2110020223323003-3203203333000021-0320110220313323)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-1211100201310202-3023000003233301-0231313213132232-2333311031330110-1322020233020000-0313120013202310-2313302232013033-0020012232333123)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-1121232011330112-3032012310222130-3332310022303323-2003302020201303-0202221021100123-0200112032023313-3203111302100302-0011332203013210)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-0200212113013321-0010302132130032-2002230233010100-0013122033122022-1101230120020032-3232303113331303-3322212303122100-1022321321123323)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-3033220110323013-2323013003303301-1033231123310221-0132003332000112-3223301221302121-3321102103000233-3133131310032101-0123333321232123)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](resources--app_setting--reference--group-001.md#canonical-3202003232100302-3113132033010233-1321100232302013-2220213133221221-2310213332333030-2133321132322100-3331103022312200-0120222001013000)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](resources--app_setting--reference--group-001.md#canonical-2321311012200003-3323332300100232-0110223212200132-0220232021222023-3133212033112223-3102230303211101-0031202030213333-1213200213303111)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2030002320211330-3031023213120032-1210231322000111-0010301210211210-0111231213220102-1201113322232000-1102103101012232-0231212332003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332132202200221-1110010300111322-1332321310022113-2012031221121300-3313111233231333-1220122230110330-2322222302002112-2320132232323213"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic — bola_detection_automatic / 301123022321 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic

<a id="canonical-1022032112020210-2121313211000312-0321022122131013-1230130312301312-0102331123223302-2323220111122312-0330130300213123-0103013021100012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for bola detection automatic.

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
bola_detection_automatic = {}
```

<a id="canonical-0331210030020213-1310222022333001-2031000110100022-1033230321020221-3300203311220333-2221002310313211-1021201322110321-1003211313223031"></a>

## Direct properties — bola_detection_automatic / 301123022321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121020131221102-2111330130321323-2331231223120321-2122022203130121-3331322120123003-0021231322313023-0220102102000002-3231013112120312"></a>

## Next pages — bola_detection_automatic / 301123022321 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1022101303111302-0031312232020020-2300201001201023-0003010330133112-2011023303111323-1113311031012231-0032111030101133-1202113310210203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132123113202111-2321220131110330-0132331220020312-2102013121002100-0212322112131332-1123013111103232-3112210313220211-0003233011031012"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection — exclude_bola_detection / 303120110120 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection

<a id="canonical-3103321212320100-3111111130100313-0320313221200310-1030222103231120-3330111133212002-1320122330223022-1203032231320001-2302233120300202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude bola detection.

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
exclude_bola_detection = {}
```

<a id="canonical-3222033210303311-2322030031010020-0000233221121020-0332030202131000-2200003003323123-2030310303332223-3010310222001100-1112001322012123"></a>

## Direct properties — exclude_bola_detection / 303120110120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022232332013333-2302232021132023-2300301102332211-1123102331002100-0201211231022320-3133100112003320-0022203200330232-2213300122322102"></a>

## Next pages — exclude_bola_detection / 303120110120 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3212222131011313-2012221300332321-0322233133000202-2023032223232022-1121000000020312-2010123002132123-3103031112003200-2023223020232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202133231011332-2031223333023320-2001103021210013-0120012203103302-0323223120010331-3230310210113330-2123222000113131-0222021033000011"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity — exclude_bot_defense_activity / 121133311111 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity

<a id="canonical-3333033130113320-3000223002133131-1322131313313031-2102013021013131-1010202332133221-2200112222203001-0021310102213132-3102232333312001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude bot defense activity.

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
exclude_bot_defense_activity = {}
```

<a id="canonical-1033200011130023-0310202202133122-2313301010103032-0102303103303213-1031113002021323-2332203023132002-0213213223332133-2312331003313002"></a>

## Direct properties — exclude_bot_defense_activity / 121133311111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131133012213310-2311011200130321-3030120212321120-2331311330332330-3013103310212312-0112033121112223-0211132130012130-1310101321031310"></a>

## Next pages — exclude_bot_defense_activity / 121133311111 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3123100011311131-3012303320313111-0220332210310102-1120211213120131-2113301303233013-2132213223100023-3230213112221022-0300021132032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222201211332020-3002102323122012-1002222220202113-0210200202100001-3011331032032233-3301212130320212-2230200000223311-0322112023201322"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity — exclude_failed_login_activity / 102031123000 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

<a id="canonical-2322321123112232-3302022113132321-3011010210120122-2212320332320011-3012223311222011-1201110332333312-3130000020232321-0111033022033311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude failed login activity.

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
exclude_failed_login_activity = {}
```

<a id="canonical-3123232332002323-2233313022112220-1203313002300103-3032133031233310-0013211310223213-0133233121201230-0013003031321000-0012121231110212"></a>

## Direct properties — exclude_failed_login_activity / 102031123000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200011220113330-1312012120331201-1022322100221123-0011221010113123-3123022111012323-0133112133323020-1231311123231131-1013202131003120"></a>

## Next pages — exclude_failed_login_activity / 102031123000 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1030233230101130-3101011302232020-1023200222133213-1103303220112213-2303221200220223-0020212202030203-0310322202322000-0312120123122311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002332020111233-2002321323003113-0320011011023110-1133230320113300-0311213202111320-1110203330302221-3222201022200012-3312333320202332"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity — exclude_forbidden_activity / 123000030023 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity

<a id="canonical-0011112222323200-0103323210133020-3122211102302213-1113302312330112-3022122312023023-1112220213000320-0021310323012102-1030210202132211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude forbidden activity.

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
exclude_forbidden_activity = {}
```

<a id="canonical-2222123233000131-2301311300230112-0132221122102030-2103203310221011-3331132011201120-0232120301101110-3012230030103021-3020310021012210"></a>

## Direct properties — exclude_forbidden_activity / 123000030023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313330201202311-3101130111131012-2012330210101210-0111101231333032-1231111220230230-2023111332332133-3312321101130000-0331033332133212"></a>

## Next pages — exclude_forbidden_activity / 123000030023 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3220213311333230-3023230111301321-0030232113111131-3020132103212311-0231232211220013-0200321110032020-1120301120313232-0300211321202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211132322222032-3120310312103323-2311010103313110-1213012310011302-0213022133022202-0131121210322003-2203110321030100-1010023023132101"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation — exclude_ip_reputation / 020000313122 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation

<a id="canonical-1210221211222122-1323233322023121-3221232111020113-1033030310112230-3202030223211020-0232302300231031-3210020132012233-0301003032233012"></a>

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
exclude_ip_reputation = {}
```

<a id="canonical-0123222203222302-0013033300100130-0213110203001322-0111013002201110-0323330030033230-0123113121130001-1203232002013002-1332220100133013"></a>

## Direct properties — exclude_ip_reputation / 020000313122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002232022032112-3223130200322012-3112203002222220-3333130120103110-0311112010201203-2103133231120020-3330133311313001-0231032200022021"></a>

## Next pages — exclude_ip_reputation / 020000313122 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3322023111302302-1201113202023331-1031223230133132-1301130022012201-0320031310212302-3213211122322031-1112223120130312-1220210332122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130011131101322-1223213322303322-3111233223013200-2133332332032213-2210033123330001-1021002100022321-2033232233113230-1211023101310310"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity — exclude_non_existent_url_activity / 203233303012 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity

<a id="canonical-2021021302213211-1232033122113222-3201033120122030-3101130032023132-2101120101200112-1222031113233021-2100123012130000-0101322130313130"></a>

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
exclude_non_existent_url_activity = {}
```

<a id="canonical-1322231301201302-1210221300211022-0120212130232002-0302013031203201-0231301210122010-2101131233331113-3332021100022230-1003232000222302"></a>

## Direct properties — exclude_non_existent_url_activity / 203233303012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132130200132030-2222013030023202-1222111122203033-3032102012322202-2103030302011313-1310300232310201-3200031311001212-2222300201310211"></a>

## Next pages — exclude_non_existent_url_activity / 203233303012 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2000121031220032-0232201010313321-0132310003122212-1122300023210111-3221210201223223-3013030300113001-3222211001332131-2003311122313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123312122011211-1113112223033132-1030012321011020-0213031320232111-1202121233231210-2132020011121210-0321123112200032-0003011023320010"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit — exclude_rate_limit / 212113223130 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit

<a id="canonical-1330202131313302-1332312212230111-3133133011201210-2330202200201223-3331110221033313-2000312330120001-1001222203010230-1220230220000302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude rate limit.

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
exclude_rate_limit = {}
```

<a id="canonical-2033220030120012-0021112003021202-1130103030201203-3021113020011301-1320310202202233-3111300311010221-1121231100322003-3123023101301120"></a>

## Direct properties — exclude_rate_limit / 212113223130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013122311211120-1322323320332101-0200202012110132-1322312030231101-2023310110113321-0320133323130303-2202121313212332-3212201131201033"></a>

## Next pages — exclude_rate_limit / 212113223130 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1000302010301201-1202112323003110-0123033202332102-1100032003011131-1302333312310023-2112102133233020-2211010302312230-1000223033312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000101220313223-0203031223023330-0010110113231012-0030333221310133-0023220210210211-1130311011032103-2331332020220330-3310211321111111"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity — exclude_waf_activity / 032332323002 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity

<a id="canonical-2332110131202222-2111002303111121-0000331013322322-2331312233121132-0203220101330130-0131131110001220-2330130022300320-1300012110120012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude waf activity.

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
exclude_waf_activity = {}
```

<a id="canonical-1323130302131220-0013001212221130-2003022231200032-3330331110210002-3131322121120300-0232301103032313-3310131332301323-2321010021103233"></a>

## Direct properties — exclude_waf_activity / 032332323002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011033111031211-3022113302020122-3013201022003000-2310213320101311-2022012333021121-3333223133013221-1213213130312012-0032222031122002"></a>

## Next pages — exclude_waf_activity / 032332323002 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2331333103301032-2111022012032100-2011300130000230-2303021301301220-1200303320322030-2110020223323003-3203203333000021-0320110220313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130103013021012-1203232123300030-3310231222301113-2301000000000000-2020222011323023-1213033233303312-0102110022000003-1131223330033311"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity — include_bot_defense_activity / 201203130132 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity

<a id="canonical-3132101101300012-0032311103102122-3331012111220112-3230203201030010-1110101100231001-3312300121303113-3332200203221032-3211332020233212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include bot defense activity.

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
include_bot_defense_activity = {}
```

<a id="canonical-3201131130101213-2021111121323233-1311302210131312-3220332201003311-1232033322233130-2303032301202311-0223333013220113-2003021210320021"></a>

## Direct properties — include_bot_defense_activity / 201203130132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102011012022321-2001300223332003-2322231213222101-3110110302132322-0322230130303020-0311011312030031-0112020101110201-0332323212110031"></a>

## Next pages — include_bot_defense_activity / 201203130132 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1211100201310202-3023000003233301-0231313213132232-2333311031330110-1322020233020000-0313120013202310-2313302232013033-0020012232333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132010203302213-1300031001311102-1321320133230232-3302220121120033-2131313322322023-3112033022133131-2033123333220001-3303323012011203"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity — include_failed_login_activity / 002110231332 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

<a id="canonical-0122231033031000-0321131233311023-3320100221322303-0030331222301123-1312300233001030-0203300312103212-3031022103102323-0031001010112122"></a>

Type: `"object"`. single nested block, Optional.

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

Upstream description:

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("login_failures_threshold")}
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
include_failed_login_activity {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031312300111201-1013321010103000-3111030223132210-1220123013033201-2220201020300203-2233002221333123-0003322203010132-0301331120301333"></a>

## Direct properties — include_failed_login_activity / 002110231332 / 3

<a id="canonical-3132201022212123-0000113031022311-2200133010021110-0321202112220331-1131113022002233-0330321003111110-2312030300100302-1030112102022300"></a>

<a id="canonical-2333103301300032-0132100323133010-3122231320333131-1013133211121100-0332230122223033-3120131322102013-2330103333233323-2130301323112322"></a>

## login_failures_threshold property — include_failed_login_activity / 002110231332 / 4

Type: `"number"`. Optional.

The number of failed logins beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-2231203322113300-0210110203221321-1013122032030000-3103302302310123-1020302212222303-3232131220302022-2310013200222312-1010231230311020"></a>

## Next pages — include_failed_login_activity / 002110231332 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1121232011330112-3032012310222130-3332310022303323-2003302020201303-0202221021100123-0200112032023313-3203111302100302-0011332203013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220331111200232-2102123232020200-3223203312320323-1002200302301301-3221232323001020-3301210321323212-2100020332132210-1223211112230221"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity — include_forbidden_activity / 023002101102 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

<a id="canonical-3021212220223230-0312130011122113-3213100010112330-1211212323233213-3230220222331033-2333033021333301-2311221323120310-1222031222322030"></a>

Type: `"object"`. single nested block, Optional.

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Upstream description:

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("forbidden_requests_threshold")}
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
include_forbidden_activity {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311021101203301-0102321013333112-1232200300210302-0220001321001121-1302022103033132-3021110313331200-1020300023123032-1231331101323021"></a>

## Direct properties — include_forbidden_activity / 023002101102 / 3

<a id="canonical-0103213103121333-0003123000103322-1322303021333232-0100212323200220-1201031320032010-1013232330223230-0231111000333332-2100010132010202"></a>

<a id="canonical-1023120312020310-1202211203211033-2111333202213131-1013223003310101-3010320100232221-0203333010100003-2233020002230121-2112310021103233"></a>

## forbidden_requests_threshold property — include_forbidden_activity / 023002101102 / 4

Type: `"number"`. Optional.

The number of forbidden requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-3332330021032233-3201221231123110-1011301110231003-3203113020320222-0220330021310131-2333232023020022-3020331300232110-3311103122330133"></a>

## Next pages — include_forbidden_activity / 023002101102 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-0200212113013321-0010302132130032-2002230233010100-0013122033122022-1101230120020032-3232303113331303-3322212303122100-1022321321123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323001311233330-0100310112101022-2021102113021232-0022003232203233-1212331331020211-1201310023222102-2032132131213010-3021233022212231"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation — include_ip_reputation / 003221323102 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation

<a id="canonical-2032310031102112-0111021310033321-1212310022130031-1102121102133300-2222321102300110-3212210001210011-1201302031130203-1001000101313233"></a>

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
include_ip_reputation = {}
```

<a id="canonical-3101123010010223-3112222001202001-0321122022130320-2132202001201003-3032122212123322-0312122332323322-1133331222031033-0201333011101112"></a>

## Direct properties — include_ip_reputation / 003221323102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230133020022321-2322210013121201-1011301322212100-0301310210130030-3232101312323310-3232120032320033-1032011100131302-3212113022312310"></a>

## Next pages — include_ip_reputation / 003221323102 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230200133201302-1331123133223010-1132333022210323-3221231322222213-0002233333122031-3111333000103322-3301110330101332-1323021213003313"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic — include_non_existent_url_activity_automatic / 331000110202 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic

<a id="canonical-3130221113120330-1110301213110012-0333110110300300-1332231103201112-3301011013031303-0132321200020202-2102033333113233-3012012022002133"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Automatic Activity Settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("high",
    "low"),
  validators.ConflictingObjectAttributes("high",
    "medium"),
  validators.ConflictingObjectAttributes("low",
    "medium")}
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
  "x-ves-oneof-field-sensitivity": "[\"high\",\"low\",\"medium\"]"
}
```

Terraform syntax:

```terraform
include_non_existent_url_activity_automatic {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112203203312112-3101320323102002-1001020233030001-2131023130203031-2021112120121312-2230322201211102-3001203332332210-2023210202122312"></a>

## Direct properties — include_non_existent_url_activity_automatic / 331000110202 / 3

- [high](resources--app_setting--reference--group-001.md#canonical-3303230010202332-0312123200203303-2123011123323103-3232002323102230-0132212302310312-0132113332301111-0113321323130121-2330111122032023): complete subsection reference.

- [low](resources--app_setting--reference--group-001.md#canonical-2103100312232330-1231013311132010-1112022211222323-3011103230101301-2021013112210103-3001211322213110-1102032311103232-0321331013321333): complete subsection reference.

- [medium](resources--app_setting--reference--group-001.md#canonical-1110100202102231-2233123321333200-1111123101031321-2331121112102030-1002021301203101-0122310321201220-3010000221301120-0323030132030200): complete subsection reference.

<a id="canonical-3013311302012230-3121211031222131-1123332221133201-0330033213233100-0322200300330312-0120211112121221-0200011221313201-2103320023211330"></a>

## Next pages — include_non_existent_url_activity_automatic / 331000110202 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](resources--app_setting--reference--group-001.md#canonical-3303230010202332-0312123200203303-2123011123323103-3232002323102230-0132212302310312-0132113332301111-0113321323130121-2330111122032023)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](resources--app_setting--reference--group-001.md#canonical-2103100312232330-1231013311132010-1112022211222323-3011103230101301-2021013112210103-3001211322213110-1102032311103232-0321331013321333)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](resources--app_setting--reference--group-001.md#canonical-1110100202102231-2233123321333200-1111123101031321-2331121112102030-1002021301203101-0122310321201220-3010000221301120-0323030132030200)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3303230010202332-0312123200203303-2123011123323103-3232002323102230-0132212302310312-0132113332301111-0113321323130121-2330111122032023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221210301133333-1203001013320211-3022131022102110-2022213330131303-1301332133103130-0010301230332323-2132323310301200-0012030113330230"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high — high / 002013310330 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high

<a id="canonical-2330332321200000-2221100213221313-3132230012323101-3232201120120030-0020221022102031-3023313202002303-2120213132123110-0331032133012012"></a>

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
high = {}
```

<a id="canonical-3231223313231203-3333323233110010-1212133302322101-1110000121223132-2031231131000100-2100221223003221-2222230333321022-1003322301111030"></a>

## Direct properties — high / 002013310330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100312201201112-3000020000202312-3230232123313110-1003011313330020-3022022001131210-2201011213311013-1321213010213301-1002213112302130"></a>

## Next pages — high / 002013310330 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2103100312232330-1231013311132010-1112022211222323-3011103230101301-2021013112210103-3001211322213110-1102032311103232-0321331013321333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103033331230300-0132333000030002-1230231200232212-3133322031132210-2102000020101001-0222020113322113-2232030032222311-3122022032030011"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low — low / 111301223322 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low

<a id="canonical-0023223322112312-3112123130112121-3232331322300131-3222220222322103-3023323032131101-2210000003000212-3120031001331131-2322111122302120"></a>

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
low = {}
```

<a id="canonical-2232003333022233-2232321033201333-3313010222132002-1202322221321223-3311031000031000-3221211011322130-1001103330311003-1011303032301121"></a>

## Direct properties — low / 111301223322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321102121211001-3200211120221211-1133023102001310-3110202333210330-0101202312032211-3332333203010130-0233223313303031-3133323120333222"></a>

## Next pages — low / 111301223322 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-1110100202102231-2233123321333200-1111123101031321-2331121112102030-1002021301203101-0122310321201220-3010000221301120-0323030132030200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130310333203032-2311012010201022-2133101110332301-2312030020032200-2100010110103002-2232100121112000-3111330311120201-1133201022130031"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium — medium / 303032312310 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium

<a id="canonical-2002022310120023-2121123232203133-2203131000102113-0113222223122313-2133133222002110-0230033021103010-1003131322030223-3232201003131301"></a>

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
medium = {}
```

<a id="canonical-0331003111331310-3213110312122101-3221101301122232-2310100010101013-2332321233103031-1310330030133201-0201123012201310-2101302213201031"></a>

## Direct properties — medium / 303032312310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033030121211220-3112123012323322-1233323110221033-1213103320002223-1333322012223030-1231232003311302-2022023122333302-2110003230221211"></a>

## Next pages — medium / 303032312310 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3033220110323013-2323013003303301-1033231123310221-0132003332000112-3223301221302121-3321102103000233-3133131310032101-0123333321232123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332000131101112-3123212121033320-1002232322112210-3110031121302233-3131311300200332-2102013233131001-0223010313203030-0200233102023223"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom — include_non_existent_url_activity_custom / 102012310113 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="canonical-2003330031313320-3320233122231100-0311222132030230-0232221021222203-2131230203013123-0103002121302002-2103202233322111-1210112202132101"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Custom Activity Setting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("nonexistent_requests_threshold")}
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
include_non_existent_url_activity_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333331303201033-0332303322011201-0211031312231112-3221030231121101-3003100230331210-0310320131212321-2300302023233011-1110231112320210"></a>

## Direct properties — include_non_existent_url_activity_custom / 102012310113 / 3

<a id="canonical-2003120012102003-0233113302210233-2130112031023303-3121103012110203-1232100212221123-0312011010313112-0012113332212323-0000010010133332"></a>

<a id="canonical-3331210023121003-3312313131321011-2120112103100121-1030102321001222-3303221132023023-0233333203133031-1211131221021131-1101030111012012"></a>

## nonexistent_requests_threshold property — include_non_existent_url_activity_custom / 102012310113 / 4

Type: `"number"`. Optional.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0232222312203030-2203312133323310-3101202120011103-3210120223022202-3231201231031023-3030113132322330-2320121123311112-1330022002001332"></a>

## Next pages — include_non_existent_url_activity_custom / 102012310113 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3202003232100302-3113132033010233-1321100232302013-2220213133221221-2310213332333030-2133321132322100-3331103022312200-0120222001013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130103003021231-2031232330002323-3210131103123100-0233110311012121-3200312231010102-1200323100331232-3100230330233112-2101003032010111"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit — include_rate_limit / 131030311030 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit

<a id="canonical-2111030211300312-0011002233113012-3310210102232323-2110012230132312-2321302010331021-2310201121201100-3012211010011222-3303322330332313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include rate limit.

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
include_rate_limit = {}
```

<a id="canonical-3221123010023123-0010312230333110-0003012232003322-0222211121302122-1120130032113132-2033202330101320-2031111033012033-3323330230130203"></a>

## Direct properties — include_rate_limit / 131030311030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133222220233003-1301121130312000-2010001132301220-3102001223111130-3323130012300301-1021212230110212-0203312113223321-1221303303210330"></a>

## Next pages — include_rate_limit / 131030311030 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2321311012200003-3323332300100232-0110223212200132-0220232021222023-3133212033112223-3102230303211101-0031202030213333-1213200213303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300101321101002-2031002133321221-2013220011323101-2232230122230221-0300303013020320-0303230033132303-0313003110112301-1000203321201032"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity — include_waf_activity / 222010202023 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

<a id="canonical-3020333331220203-2330321331212011-3202033030001221-3021122212032113-1102013222000132-1032301133220010-1122220111130200-0210033232233310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include waf activity.

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
include_waf_activity = {}
```

<a id="canonical-3031321001001031-3331211200003103-0113220311111000-1230031120023131-3110331303011101-1031333122310030-0300000332001212-0323223200113200"></a>

## Direct properties — include_waf_activity / 222010202023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322200121333113-0121000132203112-2331021301202232-3000123223322230-0122310301233320-1122021201322101-3101310032003103-1332303023113023"></a>

## Next pages — include_waf_activity / 222010202023 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-3100010103133320-1031023303101233-1020203111023232-3321121123112321-2033103000311233-3202320323210321-2131232310323032-3003101030122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310331310201110-1130013210222000-0120201300033333-2232322211221220-1022011201200102-2001013000133033-3111133033021010-2203022210113202"></a>

## app_type_settings.user_behavior_analysis_setting.enable_learning — enable_learning / 230330101230 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-3100131221210102-0300111111020311-2001110212222021-0310032201303003-0313022331322000-1212331032120301-0210310032131300-0123303011300103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learning.

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
enable_learning = {}
```

<a id="canonical-3203200101011213-2230122220012232-2212110020321003-1212023133211032-3310301323102230-3133233321332202-2120230021201313-1232123221110120"></a>

## Direct properties — enable_learning / 230330101230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222011302111200-0331133020031201-3023222231233333-2312032110222200-1312133100110032-3333122130023322-0211303203031310-3121223122330023"></a>

## Next pages — enable_learning / 230330101230 / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)

<a id="canonical-2312233300001223-3232022233011210-1010223101311011-0232002130332010-0102020202112232-0333003300011211-1233030122120233-1122012301112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120233123111202-1102013303013231-1203110233313311-1322302332321312-0031123021311002-1233133132330330-2213231131120232-3221001103130013"></a>

## timeouts — timeouts / 301011003201 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- timeouts

<a id="canonical-3020132020223121-0222232030021110-2012231100012322-2230233130312132-1011310210200130-0230203330201102-0202012332312032-0323000103123330"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331322211003310-1102123310123103-0230311302133012-2130130213213013-1120033202130020-3302000131023203-2311213020333012-0300233101023333"></a>

## Direct properties — timeouts / 301011003201 / 3

<a id="canonical-1301133022133031-3101113123011122-1213010231330230-0211012200300122-0210233020131200-0331022131113211-1123120323333221-3302331103030000"></a>

<a id="canonical-2102032310130113-2331230113302130-2222021131321021-0100111102210203-2320030301103223-3310330013033201-1103020103131223-1123001332202223"></a>

## create property — timeouts / 301011003201 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2230112213303312-1033101212023132-3233131122103202-2102303300223023-3202301200202000-1020100121012210-3003131313001300-3320030033033113"></a>

<a id="canonical-0111113312123301-0103212213200002-3021212222333301-3233002323311202-1000313213202321-3213010133120311-1103232133120001-2232330123310121"></a>

## delete property — timeouts / 301011003201 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1333223232122332-0123230223121211-1201201303320312-1231032112013131-1302221300230011-2210233333131102-0202112001030301-1203000130213320"></a>

<a id="canonical-0201312023130022-3121121100023303-0010132023022200-0030221023202122-3000122232122202-2022223231223102-2022122012030010-2121331321101203"></a>

## read property — timeouts / 301011003201 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0303033330121032-3221202130112132-2121223230000323-3220220223121131-2133110112333202-3130122101220321-3300101311112002-0120300221320202"></a>

<a id="canonical-1123303020330000-3202033101020212-3220023231131111-2021303200131222-1002331303133301-0222113121333010-1111122112112331-1011221111113303"></a>

## update property — timeouts / 301011003201 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3001312003030103-2322310203020012-1330323133212133-3301211102001103-1330203232132112-0012031000233232-0000210203222331-2301200100213011"></a>

## Next pages — timeouts / 301011003201 / 8

- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
