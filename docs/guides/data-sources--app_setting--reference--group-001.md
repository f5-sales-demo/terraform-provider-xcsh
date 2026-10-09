---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- Property reference

<a id="canonical-3333102211121323-2130321101202033-2132311212332103-0032110210122122-0232013020133223-0023020230020222-2111132112022003-2232301330302030"></a>

### Direct properties for `xcsh_app_setting`

<a id="canonical-0011301213310312-0313010030012101-0331103131331321-2113111012130111-0013310013032232-0200331320030232-0300101123032131-3000011121231103"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331): complete subsection reference.

<a id="canonical-0111310310311000-3221312001100001-1301221011103032-1333303032130330-1301021221102333-0002032110133202-2230113121313322-2232222002023113"></a>

<a id="canonical-2313321123031301-3033121111021203-3031223110323100-1233232202003131-2321112123202330-3300220223002233-1303221033013200-3333120313212103"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AppSetting.

Additional upstream details:

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

<a id="canonical-0000321230202102-2133012323220202-0213032313110220-1031313203321211-0312200233322233-0230131320323023-3033310032210131-2210322333322030"></a>

<a id="canonical-1021122233102002-3331000103203233-1031230012312123-1221130022120032-2202312011200032-2200111020032102-3132133201323212-1010312223212332"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3021232120222113-3202000112021220-2211001010330303-3031010231320212-2231003123120203-3002121311233033-2303023202310322-0323302322331012"></a>

<a id="canonical-0331311111232013-2020022301220031-2323020220312231-3201202020023332-1233012030303103-0303332000311331-2032131021313221-1032321102020322"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-2011321202110313-1013022013201230-1321313023100032-2101311033321303-1033032301312011-2022012032100010-3223101130132313-3131201003211111"></a>

<a id="canonical-3333122123210331-3012202020112210-1200300121002331-0001313030302230-2032021223300211-0013311121313322-3321003130011221-3031301200032321"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AppSetting.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-2203012203323122-1302111302322232-3310200132100231-1313032222121022-2020223110212021-2300110011332221-1103002322112002-0002033003232112"></a>

<a id="canonical-3210302023103112-1120013310303213-1232002030003112-0031332201120033-2230131012202001-1202023330201000-1221202222123002-2312220301320203"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AppSetting exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-0030322313110002-2302131113211232-3133000103121121-3033300013332100-0323112113100233-3002302301233112-0031211020021310-1021100000222223"></a>

### All schema paths for `xcsh_app_setting`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_setting--reference--group-001.md#canonical-0011301213310312-0313010030012101-0331103131331321-2113111012130111-0013310013032232-0200331320030232-0300101123032131-3000011121231103) |
| `app_type_settings` | [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-0201003132331031-0333222331213332-0010020121103033-3001120001300321-2032100113301302-1012200230233223-0102030100201113-0112001030231313) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-0123313132030132-1303213313122012-2102022203133203-2313202133333201-1102233330211333-2203123102101223-0012310223311211-3213212211123001) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](data-sources--app_setting--reference--group-001.md#canonical-3200121223320201-3202111100132112-0110200311011111-2313211132113020-0110023132100121-2220101132123012-0323031001303030-1133203113011321) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](data-sources--app_setting--reference--group-001.md#canonical-3102211030001203-0232030032012002-2310302102201000-2233103110111221-3322000201231130-0130313012202322-1320100203211330-2321132011201122) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](data-sources--app_setting--reference--group-001.md#canonical-1300131001233133-1303230133023113-1130332220320003-1200122031210213-1303211013230003-1221232032030130-2000033321211221-2100331010120012) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](data-sources--app_setting--reference--group-001.md#canonical-2313223120121203-2323230201203120-0120121331020101-1120013203023000-0221013213012020-0122322000231002-0133001010201011-1220121303121222) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](data-sources--app_setting--reference--group-001.md#canonical-3110033211211203-0131231232112302-3310030201203221-3131102012212000-0113320312132221-3232210131223320-1333010113223310-2203031200332221) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-1233120313012120-0010201001020013-0011111312310320-2012133121200303-2112333220320333-1101310312033023-0311210213100102-1123232120132113) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](data-sources--app_setting--reference--group-001.md#canonical-1303110210301321-1212020310121103-2032221102231120-1133230112213132-2332013330020001-1012133032122213-1230303131233033-3003231331030303) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](data-sources--app_setting--reference--group-001.md#canonical-2230102112101031-0011323131130231-2132020203332122-1233121002130201-2022333311031102-3313210321021232-0221120203300133-0102112203100232) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-2130010012101131-3220202100123233-2030312031232330-2303302230132022-1322222312301232-0201031203022111-2113231330033210-0302103321220212) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-2333230011011023-2210210122022313-3312010033133220-1300302120020223-2130230103213302-0022300312333301-3003100210011323-2231230221231011) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](data-sources--app_setting--reference--group-001.md#canonical-3130233320322221-0332213112132230-0112331121210112-1311003222123221-2203130131021333-0030230220310310-1000102302103022-1101021110331122) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](data-sources--app_setting--reference--group-001.md#canonical-1321221311111031-3300011122330112-0000011303132333-1001220111013020-2033010132202101-1232013312132021-2002112302202201-1303021330212230) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3011003200121202-0330231100131123-2022320023333202-2203003020211301-1103023033020202-1010322330011313-2333310122113121-2022032300200100) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--reference--group-001.md#canonical-3333000030201032-0013220232210300-0201213031222310-1302120323110333-1020023332032020-3023212111030203-3112131030021312-1302331013211333) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--reference--group-001.md#canonical-2022033131001212-2212010100111222-1120303002210002-3301330013010320-0000133002101333-2330003333131232-2330333223033233-1300002021230330) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-2330201111111133-1222002002002330-2220031032111120-3312031202210231-1200330231222301-1303111212123103-1310201123002133-0300111123100013) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-1231231302022020-1312001121331303-0231211102103231-0133223010323321-0231302021330200-3201130233112303-1223231220032320-0300122011100233) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](data-sources--app_setting--reference--group-001.md#canonical-0333323103032112-0331221230113013-2101120210021003-0331013222032032-2303031132230123-3031030000013112-2013322030332201-0032202323301212) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-1320030311312313-1000210102301233-0232021221201201-1212311301321331-3011321333130220-1302211222002103-1331231131021330-2013123132031100) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-1321101123312312-3113110013333311-0233211322012000-2020000010022202-2312013013222032-0021320221332130-3323113233110321-2210323233232122) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0300210111120032-2022033202310331-1022132111113223-0211312023003002-1130012323333121-3010220221212323-2022203131001232-3123010021033023) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-2130203210130013-3303030000012231-2012321001033030-2020332131021020-1231312331030103-3100022332013322-3210200303120303-1032310101111113) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-0120313133123130-2300103102203321-0310103020231002-2031103133020221-2122302213330200-2030200212031002-3113332131302013-1031012001213030) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-0131130100030120-2030312203221312-1302332033113023-3003322013003300-2103331001030201-0033123121011233-1121031013021331-2031231332120022) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-2120100201313200-1010031232312122-1000223122123120-0311333310102031-3213031211013301-3113303201302103-0110010202210212-2300122023103121) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-3012013303233221-0123313210012110-0113020123220221-2303121130120300-3211123132111002-2301101010301023-2122000020312023-1103132010011022) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-0122010012032331-1103313232201201-0132330312301003-1231131120211020-2201301020200023-2130302232220103-1102302221011031-1331303110321320) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0322110122120100-1013203330101203-1031312012332230-2333223222212011-2202311212321300-1011213032101032-0122202131103121-1310103211000333) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](data-sources--app_setting--reference--group-001.md#canonical-2023230020332303-0111332110003232-1031233312100302-2022201323330131-1001022013010021-0102031113300320-3133003323120301-3101011022233032) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-3211320030130320-0102312000222212-0221010321132223-2011110200330320-2123303320112120-0120033101112230-3220223223201012-2201330000013303) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](data-sources--app_setting--reference--group-001.md#canonical-1332021322322200-2022200020213210-2023030010321300-2000130302320321-2011003213020101-3311030020133332-2112312311002101-0210213011032303) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-0012121123223102-0002320210311222-0120223313031033-0110010312211021-3021121132210313-2311022012030100-1200031113132203-1211333001312023) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-0112221023302321-3312013131232132-3221213030312102-0002030230003221-1101200321013001-2031211101301122-1201131300333101-1201202021232120) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](data-sources--app_setting--reference--group-001.md#canonical-3131301000112220-0232030201013022-1331210110220200-0121201131223110-0200011101012312-1112320331311330-2223023011221312-1033201112020123) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](data-sources--app_setting--reference--group-001.md#canonical-1001303222113313-0112321011022312-0201121312312231-3010113031110123-2313030033232003-3011310110003231-1332203213002201-1320220130022230) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](data-sources--app_setting--reference--group-001.md#canonical-2223131330032101-0033011323311102-0022230123211322-3202022001322230-1113031121211233-3030001200121302-3210230013032122-2230102020311322) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-3312331001230333-0233223011213312-2221310012203203-1013130331330232-2230122012013013-2331111132321113-0300233212333120-3230222300000131) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](data-sources--app_setting--reference--group-001.md#canonical-3033101031332120-1332000112000112-0022131121300031-0010201121323302-3301311330131113-0020120320313121-0022321202222301-3203113302030010) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-3120202012023003-0313302013110103-1323231030330302-2212333113130102-2131222331120023-2131301033133033-2111023133032210-2010302130100312) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-2303232201033313-0201021221123130-3021332322312022-2330130030312113-0302213223221231-2322200300011221-1212012101330012-3213313002132132) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--reference--group-001.md#canonical-2122100133300012-2300300302110303-2303313330010132-3232220103310210-2030202233010122-3230213112231202-0312010332321231-2202203321202133) |
| `description` | [description](data-sources--app_setting--reference--group-001.md#canonical-0111310310311000-3221312001100001-1301221011103032-1333303032130330-1301021221102333-0002032110133202-2230113121313322-2232222002023113) |
| `id` | [ID](data-sources--app_setting--reference--group-001.md#canonical-0000321230202102-2133012323220202-0213032313110220-1031313203321211-0312200233322233-0230131320323023-3033310032210131-2210322333322030) |
| `labels` | [labels](data-sources--app_setting--reference--group-001.md#canonical-3021232120222113-3202000112021220-2211001010330303-3031010231320212-2231003123120203-3002121311233033-2303023202310322-0323302322331012) |
| `name` | [name](data-sources--app_setting--reference--group-001.md#canonical-2011321202110313-1013022013201230-1321313023100032-2101311033321303-1033032301312011-2022012032100010-3223101130132313-3131201003211111) |
| `namespace` | [namespace](data-sources--app_setting--reference--group-001.md#canonical-2203012203323122-1302111302322232-3310200132100231-1313032222121022-2020223110212021-2300110011332221-1103002322112002-0002033003232112) |

<a id="canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- app_type_settings

<a id="canonical-0201003132331031-0333222331213332-0010020121103033-3001120001300321-2032100113301302-1012200230233223-0102030100201113-0112001030231313"></a>

Type: `"list"`. Computed.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2032233320213311-0201003330032123-3212023002213200-0111232333223330-0201233323003113-3100202113230121-3201213032113333-0122313321010200"></a>

### Direct properties for `app_type_settings`

- [app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-1100121030331120-1012103023301132-1133133300000023-2330112310202130-1131132332222210-0022123010222210-3030312210221031-3211311123303021): complete subsection reference.

- [business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323): complete subsection reference.

- [timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323): complete subsection reference.

- [user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002): complete subsection reference.

<a id="canonical-1100121030331120-1012103023301132-1133133300000023-2330112310202130-1131132332222210-0022123010222210-3030312210221031-3211311123303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.app_type_ref` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.app_type_ref

<a id="canonical-0123313132030132-1303213313122012-2102022203133203-2313202133333201-1102233330211333-2203123102101223-0012310223311211-3213212211123001"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333122021331233-0133200002113121-0230211303020221-1222231300103201-2031023312230230-0323131210213032-2221003312333321-1313301130001023"></a>

### Direct properties for `app_type_settings.app_type_ref`

<a id="canonical-3200121223320201-3202111100132112-0110200311011111-2313211132113020-0110023132100121-2220101132123012-0323031001303030-1133203113011321"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102211030001203-0232030032012002-2310302102201000-2233103110111221-3322000201231130-0130313012202322-1320100203211330-2321132011201122"></a>

<a id="canonical-2303203221302030-0223232223230000-2331113201023301-2120131312322330-0101131131302020-1203000003213131-3221200323301232-0331012311200023"></a>

#### `app_type_settings.app_type_ref.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1300131001233133-1303230133023113-1130332220320003-1200122031210213-1303211013230003-1221232032030130-2000033321211221-2100331010120012"></a>

<a id="canonical-0131220301232323-1230133000320121-2200013220011200-2123022203102011-0011321200300310-2021311030332220-0321021322011321-3123123013031233"></a>

#### `app_type_settings.app_type_ref.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2313223120121203-2323230201203120-0120121331020101-1120013203023000-0221013213012020-0122322000231002-0133001010201011-1220121303121222"></a>

<a id="canonical-0100000322022100-3300112230122312-2102113231122130-3122021023003031-0010003311120331-1112322100010313-3300132131033202-3121313022232030"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110033211211203-0131231232112302-3310030201203221-3131102012212000-0113320312132221-3232210131223320-1333010113223310-2203031200332221"></a>

<a id="canonical-3223112001020023-0320010231103200-2122230232311230-2322100200011301-3312123313000030-2111331003121320-1120010013021030-2020120330212023"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.business_logic_markup_setting

<a id="canonical-1233120313012120-0010201001020013-0011111312310320-2012133121200303-2112333220320333-1101310312033023-0311210213100102-1123232120132113"></a>

Type: `"single"`. Computed.

Settings specifying how API Discovery will be performed.

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

<a id="canonical-1133013013211030-3130230301213233-3332222021022031-2011030321300011-2120211003300111-2103023200113012-1001223111313102-3100200213230303"></a>

### Direct properties for `app_type_settings.business_logic_markup_setting`

- [disable_spec](data-sources--app_setting--reference--group-001.md#canonical-0200312311112012-0100321113122321-0131123221110312-0310131020012232-1012222023202220-1320232010223003-1312211331013110-3113233102013202): complete subsection reference.

- [enable](data-sources--app_setting--reference--group-001.md#canonical-3213132122211021-3203223330310022-1022310133310001-0101301000002222-1223321221122312-3103012302223211-2101310212231020-3131331231203101): complete subsection reference.

<a id="canonical-0200312311112012-0100321113122321-0131123221110312-0310131020012232-1012222023202220-1320232010223003-1312211331013110-3113233102013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting.disable_spec` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- app_type_settings.business_logic_markup_setting.disable_spec

<a id="canonical-1303110210301321-1212020310121103-2032221102231120-1133230112213132-2332013330020001-1012133032122213-1230303131233033-3003231331030303"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213132122211021-3203223330310022-1022310133310001-0101301000002222-1223321221122312-3103012302223211-2101310212231020-3131331231203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.business_logic_markup_setting.enable` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-2230102112101031-0011323131130231-2132020203332122-1233121002130201-2022333311031102-3313210321021232-0221120203300133-0102112203100232"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.timeseries_analyses_setting` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-2130010012101131-3220202100123233-2030312031232330-2303302230132022-1322222312301232-0201031203022111-2113231330033210-0302103321220212"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2000121300131331-1212313330302333-2301322103232331-0323102201123311-3012100102231312-0332102020023312-1133331210322101-2300220200221001"></a>

### Direct properties for `app_type_settings.timeseries_analyses_setting`

- [metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-1103312111103121-2030011023300232-2312013100020133-2120013122001110-0020112230003033-1110213302130202-1230331321100203-1130313110113230): complete subsection reference.

<a id="canonical-1103312111103121-2030011023300232-2312013100020133-2120013122001110-0020112230003033-1110213302130202-1230331321100203-1130313110113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.timeseries_analyses_setting.metric_selectors` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323)
- app_type_settings.timeseries_analyses_setting.metric_selectors

<a id="canonical-2333230011011023-2210210122022313-3312010033133220-1300302120020223-2130230103213302-0022300312333301-3003100210011323-2231230221231011"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0000230133320203-1230031102202331-3013122310332220-2002130100231132-2022010213311000-3031331133221111-2223311011320110-2020100130312311"></a>

### Direct properties for `app_type_settings.timeseries_analyses_setting.metric_selectors`

<a id="canonical-3130233320322221-0332213112132230-0112331121210112-1311003222123221-2203130131021333-0030230220310310-1000102302103022-1101021110331122"></a>

#### `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1321221311111031-3300011122330112-0000011303132333-1001220111013020-2033010132202101-1232013312132021-2002112302202201-1303021330212230"></a>

<a id="canonical-1230300031210220-0301122202210330-3321000133203310-0010202123220013-3222133012210100-0102012112003202-0322200231231200-1303322122323001"></a>

#### `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` property

Type: `"string"`. Computed.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

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

<a id="canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.user_behavior_analysis_setting

<a id="canonical-3011003200121202-0330231100131123-2022320023333202-2203003020211301-1103023033020202-1010322330011313-2333310122113121-2022032300200100"></a>

Type: `"single"`. Computed.

Configuration for user behavior analysis.

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

<a id="canonical-2330330320232310-0313332220101322-3013013313200003-3002101131321002-3110311200000101-0302031301013303-1132110231112001-1120112332022022"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting`

- [disable_detection](data-sources--app_setting--reference--group-001.md#canonical-3001102003101201-3122033211102003-2030133123001012-2000322102011101-0211123203200300-0212301013122023-1031321130113002-3300121203232230): complete subsection reference.

- [disable_learning](data-sources--app_setting--reference--group-001.md#canonical-0002202122203233-3232101201122203-3200201220023003-2213000013012303-3030221023110200-3322121201321200-1222302020320220-3222001211202100): complete subsection reference.

- [enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301): complete subsection reference.

- [enable_learning](data-sources--app_setting--reference--group-001.md#canonical-3002320012003003-0030000001130112-0331132232021323-0313322303313030-0011130012333322-0032131111213310-0000323222133133-3301202030313010): complete subsection reference.

<a id="canonical-3001102003101201-3122033211102003-2030133123001012-2000322102011101-0211123203200300-0212301013122023-1031321130113002-3300121203232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.disable_detection` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-3333000030201032-0013220232210300-0201213031222310-1302120323110333-1020023332032020-3023212111030203-3112131030021312-1302331013211333"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002202122203233-3232101201122203-3200201220023003-2213000013012303-3030221023110200-3322121201321200-1222302020320220-3222001211202100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.disable_learning` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-2022033131001212-2212010100111222-1120303002210002-3301330013010320-0000133002101333-2330003333131232-2330333223033233-1300002021230330"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.enable_detection

<a id="canonical-2330201111111133-1222002002002330-2220031032111120-3312031202210231-1200330231222301-1303111212123103-1310201123002133-0300111123100013"></a>

Type: `"single"`. Computed.

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

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

<a id="canonical-1023002031202311-1301120023310220-1102321133222032-3101322301032031-0300230201121030-2130231130022010-2023321220232212-2031022111211001"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection`

- [bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-1003323012123303-0311001311133112-1122123013113233-1102002230232223-1220001032230233-1131121213111333-3330211230232331-3101212323032031): complete subsection reference.

<a id="canonical-0333323103032112-0331221230113013-2101120210021003-0331013222032032-2303031132230123-3031030000013112-2013322030332201-0032202323301212"></a>

<a id="canonical-2111012211221210-1210101203222222-2333113301000131-3220021132323211-1023130131221311-0333332001201001-2100230031003322-3110322000110322"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` property

Type: `"number"`. Computed.

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels. This field specifies the time period, in minutes, used by the system to decay a user's
threat level from a high to medium or medium to low or low to none.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-1021120211131220-2110223313300012-0000231023223103-1133302231111003-3103230023022132-0331331202302111-1011013010002302-1131131322210011): complete subsection reference.

- [exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-1100033330102233-2012303200223112-0300232222222123-1010003023133013-0122010331113220-3233012333203310-0323120101110222-1202313200101221): complete subsection reference.

- [exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0002122213313103-3032133310131230-0122100301113022-2300033021010121-0001012313111230-3100213022023000-3012131110022113-3212223011111031): complete subsection reference.

- [exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-2331033302323322-2312213221000000-2131231022021231-1003123130113231-0000001120032300-1311333122201000-3213103233023231-1021021333100103): complete subsection reference.

- [exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-1113211312303312-0301111101121123-2333013203310323-1320120023120111-2130313200130232-1032301012010230-0131213131200223-3133013111131213): complete subsection reference.

- [exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-3210100300120111-3132233223112133-3002111231313013-0032033313111111-0120311100231323-0333031002122001-3033023230103131-2300303111032230): complete subsection reference.

- [exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-3030332233112232-1133033211233331-0120323000303113-3132311201220213-1112032230222313-0213010032232010-1203230332300231-1312211311211223): complete subsection reference.

- [exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-2213012102133002-3220010203131211-2220232111102213-1103023130033121-0012200212113030-1110300103330031-2222012032231032-0120033301310233): complete subsection reference.

- [include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-3320120103323121-3003201202010010-0212000202130023-2211113030101330-0232030132023313-0023121032331230-0222003103231230-3222022031322000): complete subsection reference.

- [include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013): complete subsection reference.

- [include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-3303321120111303-3001302033300222-0222230201200300-2103032200130311-1300121221302320-1211211022102132-3333302201033221-2313031223000131): complete subsection reference.

- [include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-0110231030312011-0111202223331123-0100020033012112-2333120220203123-0311203000130322-0223202333322213-3201312101123100-1322022301223033): complete subsection reference.

- [include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110): complete subsection reference.

- [include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-1113332013020121-3220203212302100-2032013132131110-2312000021021212-1112023022133003-2021332321011002-1321101132010211-3101302110213130): complete subsection reference.

- [include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-1210311223301311-3213130103103302-0223213231222233-0023213320210120-0211111122112300-3132201323223322-3300102011333100-0013202310211030): complete subsection reference.

- [include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-1230131110313032-1010110212231113-1120321331001111-3100010211120222-3021332030310333-3330110113022003-2111211320011033-1301123230303221): complete subsection reference.

<a id="canonical-1003323012123303-0311001311133112-1122123013113233-1102002230232223-1220001032230233-1131121213111333-3330211230232331-3101212323032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic

<a id="canonical-1231231302022020-1312001121331303-0231211102103231-0133223010323321-0231302021330200-3201130233112303-1223231220032320-0300122011100233"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021120211131220-2110223313300012-0000231023223103-1133302231111003-3103230023022132-0331331202302111-1011013010002302-1131131322210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection

<a id="canonical-1320030311312313-1000210102301233-0232021221201201-1212311301321331-3011321333130220-1302211222002103-1331231131021330-2013123132031100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100033330102233-2012303200223112-0300232222222123-1010003023133013-0122010331113220-3233012333203310-0323120101110222-1202313200101221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity

<a id="canonical-1321101123312312-3113110013333311-0233211322012000-2020000010022202-2312013013222032-0021320221332130-3323113233110321-2210323233232122"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002122213313103-3032133310131230-0122100301113022-2300033021010121-0001012313111230-3100213022023000-3012131110022113-3212223011111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

<a id="canonical-0300210111120032-2022033202310331-1022132111113223-0211312023003002-1130012323333121-3010220221212323-2022203131001232-3123010021033023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331033302323322-2312213221000000-2131231022021231-1003123130113231-0000001120032300-1311333122201000-3213103233023231-1021021333100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity

<a id="canonical-2130203210130013-3303030000012231-2012321001033030-2020332131021020-1231312331030103-3100022332013322-3210200303120303-1032310101111113"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113211312303312-0301111101121123-2333013203310323-1320120023120111-2130313200130232-1032301012010230-0131213131200223-3133013111131213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation

<a id="canonical-0120313133123130-2300103102203321-0310103020231002-2031103133020221-2122302213330200-2030200212031002-3113332131302013-1031012001213030"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210100300120111-3132233223112133-3002111231313013-0032033313111111-0120311100231323-0333031002122001-3033023230103131-2300303111032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity

<a id="canonical-0131130100030120-2030312203221312-1302332033113023-3003322013003300-2103331001030201-0033123121011233-1121031013021331-2031231332120022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030332233112232-1133033211233331-0120323000303113-3132311201220213-1112032230222313-0213010032232010-1203230332300231-1312211311211223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit

<a id="canonical-2120100201313200-1010031232312122-1000223122123120-0311333310102031-3213031211013301-3113303201302103-0110010202210212-2300122023103121"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213012102133002-3220010203131211-2220232111102213-1103023130033121-0012200212113030-1110300103330031-2222012032231032-0120033301310233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity

<a id="canonical-3012013303233221-0123313210012110-0113020123220221-2303121130120300-3211123132111002-2301101010301023-2122000020312023-1103132010011022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320120103323121-3003201202010010-0212000202130023-2211113030101330-0232030132023313-0023121032331230-0222003103231230-3222022031322000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity

<a id="canonical-0122010012032331-1103313232201201-0132330312301003-1231131120211020-2201301020200023-2130302232220103-1102302221011031-1331303110321320"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

<a id="canonical-0322110122120100-1013203330101203-1031312012332230-2333223222212011-2202311212321300-1011213032101032-0122202131103121-1310103211000333"></a>

Type: `"single"`. Computed.

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

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

<a id="canonical-0302010211032202-2003233012313230-2311013103211333-2302230122212213-2120331023012231-3321333210310033-0123202102130020-0223202013233312"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity`

<a id="canonical-2023230020332303-0111332110003232-1031233312100302-2022201323330131-1001022013010021-0102031113300320-3133003323120301-3101011022233032"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` property

Type: `"number"`. Computed.

The number of failed logins beyond which the system will flag this user as malicious.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3303321120111303-3001302033300222-0222230201200300-2103032200130311-1300121221302320-1211211022102132-3333302201033221-2313031223000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

<a id="canonical-3211320030130320-0102312000222212-0221010321132223-2011110200330320-2123303320112120-0120033101112230-3220223223201012-2201330000013303"></a>

Type: `"single"`. Computed.

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

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

<a id="canonical-3221213220321231-2212303302132113-3101002302321130-3020302133122011-1323201100002202-1213031131230233-1012122333033201-3200223233102011"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity`

<a id="canonical-1332021322322200-2022200020213210-2023030010321300-2000130302320321-2011003213020101-3311030020133332-2112312311002101-0210213011032303"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` property

Type: `"number"`. Computed.

The number of forbidden requests beyond which the system will flag this user as malicious.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110231030312011-0111202223331123-0100020033012112-2333120220203123-0311203000130322-0223202333322213-3201312101123100-1322022301223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation

<a id="canonical-0012121123223102-0002320210311222-0120223313031033-0110010312211021-3021121132210313-2311022012030100-1200031113132203-1211333001312023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic

<a id="canonical-0112221023302321-3312013131232132-3221213030312102-0002030230003221-1101200321013001-2031211101301122-1201131300333101-1201202021232120"></a>

Type: `"single"`. Computed.

Non-existent URL Automatic Activity Settings.

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

<a id="canonical-3230312332013000-2011312002201130-1112123203012221-2203210330112122-1232210202122201-1212120030011002-0331003230022000-0033210113311210"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic`

- [high](data-sources--app_setting--reference--group-001.md#canonical-1000120102030200-3100112023223130-3320313213330012-2100310331303111-0210320233220323-1123232311123002-0102123032122311-2221100103010223): complete subsection reference.

- [low](data-sources--app_setting--reference--group-001.md#canonical-2001232101233222-3220321222002323-2233332332323121-2102000031320330-0123222002211320-3102112211312122-3222221312312202-3232211131202212): complete subsection reference.

- [medium](data-sources--app_setting--reference--group-001.md#canonical-1120301301313110-0223320213030232-1132320332002332-2330232201221011-3211112312320333-2131311120010323-0130311302101303-1031032213301303): complete subsection reference.

<a id="canonical-1000120102030200-3100112023223130-3320313213330012-2100310331303111-0210320233220323-1123232311123002-0102123032122311-2221100103010223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high

<a id="canonical-3131301000112220-0232030201013022-1331210110220200-0121201131223110-0200011101012312-1112320331311330-2223023011221312-1033201112020123"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001232101233222-3220321222002323-2233332332323121-2102000031320330-0123222002211320-3102112211312122-3222221312312202-3232211131202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low

<a id="canonical-1001303222113313-0112321011022312-0201121312312231-3010113031110123-2313030033232003-3011310110003231-1332203213002201-1320220130022230"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120301301313110-0223320213030232-1132320332002332-2330232201221011-3211112312320333-2131311120010323-0130311302101303-1031032213301303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium

<a id="canonical-2223131330032101-0033011323311102-0022230123211322-3202022001322230-1113031121211233-3030001200121302-3210230013032122-2230102020311322"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113332013020121-3220203212302100-2032013132131110-2312000021021212-1112023022133003-2021332321011002-1321101132010211-3101302110213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="canonical-3312331001230333-0233223011213312-2221310012203203-1013130331330232-2230122012013013-2331111132321113-0300233212333120-3230222300000131"></a>

Type: `"single"`. Computed.

Non-existent URL Custom Activity Setting.

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

<a id="canonical-0331010330112313-0103332212202301-3011321120102033-0130103023030220-0302221111000110-0030012200320203-2021111303113012-2301111310303003"></a>

### Direct properties for `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom`

<a id="canonical-3033101031332120-1332000112000112-0022131121300031-0010201121323302-3301311330131113-0020120320313121-0022321202222301-3203113302030010"></a>

#### `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` property

Type: `"number"`. Computed.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1210311223301311-3213130103103302-0223213231222233-0023213320210120-0211111122112300-3132201323223322-3300102011333100-0013202310211030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit

<a id="canonical-3120202012023003-0313302013110103-1323231030330302-2212333113130102-2131222331120023-2131301033133033-2111023133032210-2010302130100312"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230131110313032-1010110212231113-1120321331001111-3100010211120222-3021332030310333-3330110113022003-2111211320011033-1301123230303221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

<a id="canonical-2303232201033313-0201021221123130-3021332322312022-2330130030312113-0302213223221231-2322200300011221-1212012101330012-3213313002132132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002320012003003-0030000001130112-0331132232021323-0313322303313030-0011130012333322-0032131111213310-0000323222133133-3301202030313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_type_settings.user_behavior_analysis_setting.enable_learning` properties

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-2122100133300012-2300300302110303-2303313330010132-3232220103310210-2030202233010122-3230213112231202-0312010332321231-2202203321202133"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
