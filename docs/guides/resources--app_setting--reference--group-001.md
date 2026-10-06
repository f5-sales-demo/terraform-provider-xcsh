---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- Property reference

<a id="canonical-0103001322310113-1222310332200031-2232122313002121-1011310333301302-2003001030202030-1330010222212113-1311232332201331-3300122302123023"></a>

### Direct properties for `xcsh_app_setting`

<a id="canonical-0303201112010032-0130033111212301-3221333312200100-3002020101323021-3033202030313330-1021212023321032-2321120121000330-3122123300120133"></a>

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

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112): complete subsection reference.

<a id="canonical-0022030103333011-1103131120023100-0013122200311011-2302021011212333-1031101203203122-0230131100333132-1312120111312121-0232323102122213"></a>

<a id="canonical-3203222223210220-1321210330130323-0322231002332110-3212033230113232-2232003331013100-3110020110223012-2320232220211321-3213212133322003"></a>

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

<a id="canonical-3121001130200322-1000333101100110-2101112000012100-2202113011032002-0113302313303321-3320333030112312-1000222201012000-3113122102232210"></a>

<a id="canonical-3103010001321023-1123010112122120-3023121132322022-3021301332333212-3323102030313121-2332111320331211-1203012233321021-0002012221003313"></a>

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

<a id="canonical-0300101311300122-3300212201013103-1203120120233021-2321203302331002-2031123100321122-2120123110210001-2200202302333011-0102133132012313"></a>

<a id="canonical-3333111001100012-1322322220101030-2202002222333120-2022231320031013-2122233320010033-1023322300223011-1223223033033311-1101123011320101"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3120311110320202-3222133302121023-0103232023021121-1301003223003212-2232020133303310-1021120232212210-2232330101100102-2023332102003333"></a>

<a id="canonical-2111211110003120-1331320032022331-1321011231312021-3203201321012333-3210320122321223-3331100212310012-2101221121100033-0201321303312031"></a>

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

<a id="canonical-0000300203012300-1313132031011220-0031113023211212-1120333330001213-3133102320101320-2112033133230113-3321220023211300-2302111000130131"></a>

<a id="canonical-1132331031323323-1301232301332312-2113203223203231-3303323213023011-2023132333021200-3333223111321100-3312211111120210-1101303023021003"></a>

#### `name` property

Type: `"string"`. Required.

Name of the App Setting. Must be unique within the namespace.

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

<a id="canonical-2131211130112020-1032331010231300-1222100030130133-0301200230133320-1023303031133131-1303103120302020-1113103210123130-3203221202033311"></a>

<a id="canonical-2022111313002112-0322121033221123-3120210100010330-0123030100130123-1112102123110013-2213021311302311-3321313230102023-3123002323122322"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the App Setting is created.

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

- [timeouts](resources--app_setting--reference--group-001.md#canonical-2312233300001223-3232022233011210-1010223101311011-0232002130332010-0102020202112232-0333003300011211-1233030122120233-1122012301112032): complete subsection reference.

<a id="canonical-2120111000302022-3103232203120112-3303222121210312-2223231110313210-2030312220210201-3203333202100220-0312103013130032-2232310033220033"></a>

### All schema paths for `xcsh_app_setting`

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

<a id="canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings` properties

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

<a id="canonical-1023031003330210-3221021101101303-3321032301332020-2012111230011010-1122312221032233-3202211310313200-3331232120030102-3000233211220131"></a>

### Direct properties for `app_type_settings`

- [app_type_ref](resources--app_setting--reference--group-001.md#canonical-0031103021102102-3301300233201013-3333223113310100-2203322101303330-0111303213312001-3011103102123213-2103130311000303-2223003002000211): complete subsection reference.

- [business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032): complete subsection reference.

- [timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020): complete subsection reference.

- [user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012): complete subsection reference.

<a id="canonical-0031103021102102-3301300233201013-3333223113310100-2203322101303330-0111303213312001-3011103102123213-2103130311000303-2223003002000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.app_type_ref` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.app_type_ref

<a id="canonical-2302303212222310-3130303222330222-0301131232301121-2002220010203131-3211033122112120-0332203232212131-0113001311131112-1113220102220010"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-3131120122013001-3300110220222101-1221111233100203-2333203333032120-0311012023200222-2223122023002330-3100211300302200-3122021021213020"></a>

### Direct properties for `app_type_settings.app_type_ref`

<a id="canonical-0312010200032200-3301031213200101-1001223332313232-3022220123220200-3222302031302313-2320002201212203-0000322302123311-0313310120003133"></a>

#### `app_type_settings.app_type_ref.kind` property

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

<a id="canonical-3123000302102330-1132203233302330-3103222131001220-3100203220103322-2303021222002210-0100211300131000-0233003310231203-2000322300003131"></a>

<a id="canonical-3331332100303001-2100310003233123-1010130333322012-2211333122211330-1301210210200212-2200021011010331-1033231232332013-1220123022022001"></a>

#### `app_type_settings.app_type_ref.name` property

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

<a id="canonical-3303112110313133-3232122110302011-0311103212300021-0211212112111102-1022012111012212-3003330211030213-1201221121011312-1001301111021300"></a>

<a id="canonical-2210231213313210-0330003033013332-0010131133130232-1110322222130323-0023211100113320-3331220132300010-1212333302211320-2132320210100130"></a>

#### `app_type_settings.app_type_ref.namespace` property

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

<a id="canonical-3000033103021023-3120203002120202-3112310221300320-1022211101323310-0002303130122310-2000021001321312-3022303131313110-3320131010303130"></a>

<a id="canonical-2130232031231103-2030000233203020-0211220233023323-3300322301102201-1222120233321130-1321111010032020-0310212123022112-3200301003002333"></a>

#### `app_type_settings.app_type_ref.tenant` property

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

<a id="canonical-0001313312012331-1130312012223303-2301021121122313-1121202223132103-1111301321221003-2121311210213020-1311033030320311-0202212010123123"></a>

<a id="canonical-1132230301213330-2310130132202330-2223002112020112-3220310232022103-0322010100003010-2201022113103323-2133212302213133-1132232033111022"></a>

#### `app_type_settings.app_type_ref.uid` property

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

<a id="canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting` properties

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

<a id="canonical-1213002223133031-2123102310333100-3001310221010121-1020120131212330-2203013212233000-0203101013333322-0130233132210103-1031232230101233"></a>

### Direct properties for `app_type_settings.business_logic_markup_setting`

- [disable_spec](resources--app_setting--reference--group-001.md#canonical-2103102313033100-2102232331010211-0001221033310303-2013223003002020-0202130220122002-1310121331003301-0232111331100202-0201012300110131): complete subsection reference.

- [enable](resources--app_setting--reference--group-001.md#canonical-1313220311133123-2232201103020022-2333020322311131-1220310013331221-2312103201311021-2110033002100310-2200313200113031-3131313102113203): complete subsection reference.

<a id="canonical-2103102313033100-2102232331010211-0001221033310303-2013223003002020-0202130220122002-1310121331003301-0232111331100202-0201012300110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting.disable_spec` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313220311133123-2232201103020022-2333020322311131-1220310013331221-2312103201311021-2110033002100310-2200313200113031-3131313102113203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting.enable` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-0233302011231203-3201101221012200-1122323332202333-2233110033112202-2032210221331312-3132210022230322-1220100200301013-0222231111022201"></a>

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
enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.timeseries_analyses_setting` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-2211000132221302-2023211332112103-0020311311310330-0200223033312001-3110313133030220-2210111203322012-0203131313231102-1012302322313002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for timeseries analyses setting.

Additional upstream details:

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

<a id="canonical-3032321322001131-1333030130002103-0121200032123203-0310033021211230-3102233001102131-2321123102012312-2333130301013113-0323202031331230"></a>

### Direct properties for `app_type_settings.timeseries_analyses_setting`

- [metric_selectors](resources--app_setting--reference--group-001.md#canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123): complete subsection reference.

<a id="canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.timeseries_analyses_setting.metric_selectors` properties

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

<a id="canonical-0221110120000300-3023200021203323-1020021333210311-1233003121323011-3333133011211020-3220130113330030-3230001302211303-0023203213111013"></a>

### Direct properties for `app_type_settings.timeseries_analyses_setting.metric_selectors`

<a id="canonical-0301103010320221-3231132300122212-1021323003013113-0012130300202211-0330001230211321-3221212022321131-1021301133033323-2203130122113113"></a>

#### `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` property

Type: `["list", "string"]`. Optional.

\[Enum: NO\_METRICS|REQUEST\_RATE|ERROR\_RATE|LATENCY|THROUGHPUT\] Choose one or more metrics to be
included in the detection logic. Possible values are \`NO\_METRICS\`, \`REQUEST\_RATE\`,
\`ERROR\_RATE\`, \`LATENCY\`, \`THROUGHPUT\`. Defaults to \`NO\_METRICS\`.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0212031023333031-1200013211131000-0033223013123201-2003201021322303-1023310311120032-3122112021130002-3123312222022232-2023030131131132"></a>

<a id="canonical-2300001010332022-1120020113033103-3311120303122233-1212011333203202-2112322323022211-0132303311021101-1232031202011122-0121203213301131"></a>

#### `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` property

Type: `"string"`. Optional.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

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

<a id="canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting` properties

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

<a id="canonical-3011312130020313-2213000001131102-3010332310032220-0222212002213111-0210110133113130-2133100031302232-0231302323002121-0021003202231232"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting`

- [disable_detection](resources--app_setting--reference--group-001.md#canonical-1113321302003221-2113300101202330-2030303232031111-2322310111302001-3300012220001130-3033333310133300-2322330130303300-0330210101132230): complete subsection reference.

- [disable_learning](resources--app_setting--reference--group-001.md#canonical-2133320120113332-2132323030231223-3331322212232313-3320132031030321-3122103030120331-1010200101323213-2102031323331202-0112111302022203): complete subsection reference.

- [enable_detection](resources--app_setting--reference--group-001.md#canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013): complete subsection reference.

- [enable_learning](resources--app_setting--reference--group-001.md#canonical-3100010103133320-1031023303101233-1020203111023232-3321121123112321-2033103000311233-3202320323210321-2131232310323032-3003101030122022): complete subsection reference.

<a id="canonical-1113321302003221-2113300101202330-2030303232031111-2322310111302001-3300012220001130-3033333310133300-2322330130303300-0330210101132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.disable_detection` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-3312103330121211-1121031103030030-2000011211201032-0023320120020213-1323323322202330-0213030033120213-1323101113322231-2321112333232133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable detection.

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
disable_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133320120113332-2132323030231223-3331322212232313-3320132031030321-3122103030120331-1010200101323213-2102031323331202-0112111302022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.disable_learning` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-1130100232002131-3302203001022301-3332213003111132-3033030231100230-1130002221310033-1113320300100021-0202123010023321-2133032313203022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learning.

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
disable_learning = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301233001121111-0330003101300113-3221313100331301-0133300032232330-2311130303230113-0221001323313120-2333312131001301-1102131003303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection` properties

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

<a id="canonical-3010132221123020-3213211130323302-0312211330200320-1112111311003231-3111100113201201-0322001103031320-0223203203212121-2101333110332213"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection`

- [bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-2030002320211330-3031023213120032-1210231322000111-0010301210211210-0111231213220102-1201113322232000-1102103101012232-0231212332003011): complete subsection reference.

<a id="canonical-1201232013121300-1021302133211123-2013003331301221-0111022311222112-2233111120100011-3312033213221001-1331233020113113-2232233022303232"></a>

<a id="canonical-2031132130010133-0000103211013301-2100330223102033-3223332001231131-0230112231131222-0212320322223301-3302100203013332-1201010032122333"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2030002320211330-3031023213120032-1210231322000111-0010301210211210-0111231213220102-1201113322232000-1102103101012232-0231212332003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` properties

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
bola_detection_automatic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022101303111302-0031312232020020-2300201001201023-0003010330133112-2011023303111323-1113311031012231-0032111030101133-1202113310210203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` properties

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
exclude_bola_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212222131011313-2012221300332321-0322233133000202-2023032223232022-1121000000020312-2010123002132123-3103031112003200-2023223020232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` properties

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
exclude_bot_defense_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123100011311131-3012303320313111-0220332210310102-1120211213120131-2113301303233013-2132213223100023-3230213112221022-0300021132032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` properties

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
exclude_failed_login_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030233230101130-3101011302232020-1023200222133213-1103303220112213-2303221200220223-0020212202030203-0310322202322000-0312120123122311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` properties

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
exclude_forbidden_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220213311333230-3023230111301321-0030232113111131-3020132103212311-0231232211220013-0200321110032020-1120301120313232-0300211321202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` properties

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
exclude_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322023111302302-1201113202023331-1031223230133132-1301130022012201-0320031310212302-3213211122322031-1112223120130312-1220210332122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` properties

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
exclude_non_existent_url_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000121031220032-0232201010313321-0132310003122212-1122300023210111-3221210201223223-3013030300113001-3222211001332131-2003311122313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` properties

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
exclude_rate_limit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000302010301201-1202112323003110-0123033202332102-1100032003011131-1302333312310023-2112102133233020-2211010302312230-1000223033312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` properties

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
exclude_waf_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331333103301032-2111022012032100-2011300130000230-2303021301301220-1200303320322030-2110020223323003-3203203333000021-0320110220313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` properties

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
include_bot_defense_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211100201310202-3023000003233301-0231313213132232-2333311031330110-1322020233020000-0313120013202310-2313302232013033-0020012232333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` properties

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

<a id="canonical-1132010203302213-1300031001311102-1321320133230232-3302220121120033-2131313322322023-3112033022133131-2033123333220001-3303323012011203"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity`

<a id="canonical-3132201022212123-0000113031022311-2200133010021110-0321202112220331-1131113022002233-0330321003111110-2312030300100302-1030112102022300"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1121232011330112-3032012310222130-3332310022303323-2003302020201303-0202221021100123-0200112032023313-3203111302100302-0011332203013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` properties

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

<a id="canonical-3220331111200232-2102123232020200-3223203312320323-1002200302301301-3221232323001020-3301210321323212-2100020332132210-1223211112230221"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity`

<a id="canonical-0103213103121333-0003123000103322-1322303021333232-0100212323200220-1201031320032010-1013232330223230-0231111000333332-2100010132010202"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0200212113013321-0010302132130032-2002230233010100-0013122033122022-1101230120020032-3232303113331303-3322212303122100-1022321321123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` properties

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
include_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332112223302200-3031023300100222-3321112203012230-0121103022300030-1210131113223010-2333023001202322-0330001213332103-2232033223123101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` properties

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

<a id="canonical-3230200133201302-1331123133223010-1132333022210323-3221231322222213-0002233333122031-3111333000103322-3301110330101332-1323021213003313"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic`

- [high](resources--app_setting--reference--group-001.md#canonical-3303230010202332-0312123200203303-2123011123323103-3232002323102230-0132212302310312-0132113332301111-0113321323130121-2330111122032023): complete subsection reference.

- [low](resources--app_setting--reference--group-001.md#canonical-2103100312232330-1231013311132010-1112022211222323-3011103230101301-2021013112210103-3001211322213110-1102032311103232-0321331013321333): complete subsection reference.

- [medium](resources--app_setting--reference--group-001.md#canonical-1110100202102231-2233123321333200-1111123101031321-2331121112102030-1002021301203101-0122310321201220-3010000221301120-0323030132030200): complete subsection reference.

<a id="canonical-3303230010202332-0312123200203303-2123011123323103-3232002323102230-0132212302310312-0132113332301111-0113321323130121-2330111122032023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` properties

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
high = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103100312232330-1231013311132010-1112022211222323-3011103230101301-2021013112210103-3001211322213110-1102032311103232-0321331013321333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` properties

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
low = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110100202102231-2233123321333200-1111123101031321-2331121112102030-1002021301203101-0122310321201220-3010000221301120-0323030132030200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` properties

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
medium = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033220110323013-2323013003303301-1033231123310221-0132003332000112-3223301221302121-3321102103000233-3133131310032101-0123333321232123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` properties

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

<a id="canonical-1332000131101112-3123212121033320-1002232322112210-3110031121302233-3131311300200332-2102013233131001-0223010313203030-0200233102023223"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom`

<a id="canonical-2003120012102003-0233113302210233-2130112031023303-3121103012110203-1232100212221123-0312011010313112-0012113332212323-0000010010133332"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3202003232100302-3113132033010233-1321100232302013-2220213133221221-2310213332333030-2133321132322100-3331103022312200-0120222001013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` properties

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
include_rate_limit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321311012200003-3323332300100232-0110223212200132-0220232021222023-3133212033112223-3102230303211101-0031202030213333-1213200213303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` properties

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
include_waf_activity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100010103133320-1031023303101233-1020203111023232-3321121123112321-2033103000311233-3202320323210321-2131232310323032-3003101030122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_learning` properties

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-2200230032231032-1201320201203010-3023331021011200-3013113323002321-0231221002331231-0201131101133011-2220013233031230-0123033010210012)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-3100131221210102-0300111111020311-2001110212222021-0310032201303003-0313022331322000-1212331032120301-0210310032131300-0123303011300103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learning.

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
enable_learning = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312233300001223-3232022233011210-1010223101311011-0232002130332010-0102020202112232-0333003300011211-1233030122120233-1122012301112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-1120233123111202-1102013303013231-1203110233313311-1322302332321312-0031123021311002-1233133132330330-2213231131120232-3221001103130013"></a>

### Direct properties for `timeouts`

<a id="canonical-1301133022133031-3101113123011122-1213010231330230-0211012200300122-0210233020131200-0331022131113211-1123120323333221-3302331103030000"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2230112213303312-1033101212023132-3233131122103202-2102303300223023-3202301200202000-1020100121012210-3003131313001300-3320030033033113"></a>

<a id="canonical-2331322211003310-1102123310123103-0230311302133012-2130130213213013-1120033202130020-3302000131023203-2311213020333012-0300233101023333"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1333223232122332-0123230223121211-1201201303320312-1231032112013131-1302221300230011-2210233333131102-0202112001030301-1203000130213320"></a>

<a id="canonical-2102032310130113-2331230113302130-2222021131321021-0100111102210203-2320030301103223-3310330013033201-1103020103131223-1123001332202223"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0303033330121032-3221202130112132-2121223230000323-3220220223121131-2133110112333202-3130122101220321-3300101311112002-0120300221320202"></a>

<a id="canonical-0111113312123301-0103212213200002-3021212222333301-3233002323311202-1000313213202321-3213010133120311-1103232133120001-2232330123310121"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
