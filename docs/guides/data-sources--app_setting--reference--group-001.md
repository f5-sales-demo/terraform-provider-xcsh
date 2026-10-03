---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333102211121323-2130321101202033-2132311212332103-0032110210122122-0232013020133223-0023020230020222-2111132112022003-2232301330302030"></a>

## Property reference — Property reference / 133212232213 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- Property reference

<a id="canonical-2313321123031301-3033121111021203-3031223110323100-1233232202003131-2321112123202330-3300220223002233-1303221033013200-3333120313212103"></a>

## Direct properties — Property reference / 133212232213 / 3

<a id="canonical-0011301213310312-0313010030012101-0331103131331321-2113111012130111-0013310013032232-0200331320030232-0300101123032131-3000011121231103"></a>

<a id="canonical-1021122233102002-3331000103203233-1031230012312123-1221130022120032-2202312011200032-2200111020032102-3132133201323212-1010312223212332"></a>

## annotations property — Property reference / 133212232213 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331): complete subsection reference.

<a id="canonical-0111310310311000-3221312001100001-1301221011103032-1333303032130330-1301021221102333-0002032110133202-2230113121313322-2232222002023113"></a>

<a id="canonical-0331311111232013-2020022301220031-2323020220312231-3201202020023332-1233012030303103-0303332000311331-2032131021313221-1032321102020322"></a>

## description property — Property reference / 133212232213 / 5

Type: `"string"`. Computed.

Description of the AppSetting.

Upstream description:

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

<a id="canonical-0000321230202102-2133012323220202-0213032313110220-1031313203321211-0312200233322233-0230131320323023-3033310032210131-2210322333322030"></a>

<a id="canonical-3333122123210331-3012202020112210-1200300121002331-0001313030302230-2032021223300211-0013311121313322-3321003130011221-3031301200032321"></a>

## ID property — Property reference / 133212232213 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3021232120222113-3202000112021220-2211001010330303-3031010231320212-2231003123120203-3002121311233033-2303023202310322-0323302322331012"></a>

<a id="canonical-3210302023103112-1120013310303213-1232002030003112-0031332201120033-2230131012202001-1202023330201000-1221202222123002-2312220301320203"></a>

## labels property — Property reference / 133212232213 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-2011321202110313-1013022013201230-1321313023100032-2101311033321303-1033032301312011-2022012032100010-3223101130132313-3131201003211111"></a>

<a id="canonical-0030322313110002-2302131113211232-3133000103121121-3033300013332100-0323112113100233-3002302301233112-0031211020021310-1021100000222223"></a>

## name property — Property reference / 133212232213 / 8

Type: `"string"`. Required.

Name of the AppSetting.

Upstream description:

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

<a id="canonical-2203012203323122-1302111302322232-3310200132100231-1313032222121022-2020223110212021-2300110011332221-1103002322112002-0002033003232112"></a>

<a id="canonical-0013313020031013-0323112313233000-0220230230232333-2233210003122303-2233331301211202-3332330330112102-2220320232331210-2333200102021000"></a>

## namespace property — Property reference / 133212232213 / 9

Type: `"string"`. Required.

Namespace where the AppSetting exists.

Upstream description:

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

<a id="canonical-2333203023313330-1311232011133310-2002212331312312-0221001011330201-1302333113232112-3020330330133222-3023322222302113-3030200210320230"></a>

## All schema paths — Property reference / 133212232213 / 10

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

<a id="canonical-3023011220121220-0302031011300103-1131131212333111-3123302310302330-2001300030001223-0002231020113211-2133213202110310-2000303311013113"></a>

## Next pages — Property reference / 133212232213 / 11

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032233320213311-0201003330032123-3212023002213200-0111232333223330-0201233323003113-3100202113230121-3201213032113333-0122313321010200"></a>

## app_type_settings — app_type_settings / 011331132110 / 2

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

<a id="canonical-2223011102113210-2113300021231320-2233103311120001-2233133310333313-0122130033131310-0231322010010333-0311221303011122-3313110303210123"></a>

## Direct properties — app_type_settings / 011331132110 / 3

- [app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-1100121030331120-1012103023301132-1133133300000023-2330112310202130-1131132332222210-0022123010222210-3030312210221031-3211311123303021): complete subsection reference.

- [business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323): complete subsection reference.

- [timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323): complete subsection reference.

- [user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002): complete subsection reference.

<a id="canonical-0210233023033211-3200212303031303-1333220212003320-3301310020022213-1103233133202212-0002013211200213-1301332102323302-2132330032000311"></a>

## Next pages — app_type_settings / 011331132110 / 4

- [app_type_settings.app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-1100121030331120-1012103023301132-1133133300000023-2330112310202130-1131132332222210-0022123010222210-3030312210221031-3211311123303021)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1100121030331120-1012103023301132-1133133300000023-2330112310202130-1131132332222210-0022123010222210-3030312210221031-3211311123303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333122021331233-0133200002113121-0230211303020221-1222231300103201-2031023312230230-0323131210213032-2221003312333321-1313301130001023"></a>

## app_type_settings.app_type_ref — app_type_ref / 010331001013 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.app_type_ref

<a id="canonical-0123313132030132-1303213313122012-2102022203133203-2313202133333201-1102233330211333-2203123102101223-0012310223311211-3213212211123001"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2303203221302030-0223232223230000-2331113201023301-2120131312322330-0101131131302020-1203000003213131-3221200323301232-0331012311200023"></a>

## Direct properties — app_type_ref / 010331001013 / 3

<a id="canonical-3200121223320201-3202111100132112-0110200311011111-2313211132113020-0110023132100121-2220101132123012-0323031001303030-1133203113011321"></a>

<a id="canonical-0131220301232323-1230133000320121-2200013220011200-2123022203102011-0011321200300310-2021311030332220-0321021322011321-3123123013031233"></a>

## kind property — app_type_ref / 010331001013 / 4

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

<a id="canonical-3102211030001203-0232030032012002-2310302102201000-2233103110111221-3322000201231130-0130313012202322-1320100203211330-2321132011201122"></a>

<a id="canonical-0100000322022100-3300112230122312-2102113231122130-3122021023003031-0010003311120331-1112322100010313-3300132131033202-3121313022232030"></a>

## name property — app_type_ref / 010331001013 / 5

Type: `"string"`. Computed.

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

<a id="canonical-1300131001233133-1303230133023113-1130332220320003-1200122031210213-1303211013230003-1221232032030130-2000033321211221-2100331010120012"></a>

<a id="canonical-3223112001020023-0320010231103200-2122230232311230-2322100200011301-3312123313000030-2111331003121320-1120010013021030-2020120330212023"></a>

## namespace property — app_type_ref / 010331001013 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2313223120121203-2323230201203120-0120121331020101-1120013203023000-0221013213012020-0122322000231002-0133001010201011-1220121303121222"></a>

<a id="canonical-2223003323202303-2003210023133310-1021202132032030-3302013313320110-1102321013112030-3101311103112102-2331123313113112-2123002112211232"></a>

## tenant property — app_type_ref / 010331001013 / 7

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

<a id="canonical-3110033211211203-0131231232112302-3310030201203221-3131102012212000-0113320312132221-3232210131223320-1333010113223310-2203031200332221"></a>

<a id="canonical-3332032202301012-2220322202011130-1011303102011012-3131120302213120-3211113130133111-0131202302131332-2120021100032320-3112331300130221"></a>

## uid property — app_type_ref / 010331001013 / 8

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

<a id="canonical-0213123300300030-0221122202212131-2000121110320331-0131210021231230-0111231200222032-3012112222211132-0230030031333310-1103112302002232"></a>

## Next pages — app_type_ref / 010331001013 / 9

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133013013211030-3130230301213233-3332222021022031-2011030321300011-2120211003300111-2103023200113012-1001223111313102-3100200213230303"></a>

## app_type_settings.business_logic_markup_setting — business_logic_markup_setting / 131033103223 / 2

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

<a id="canonical-0103000201010212-3323021131201221-0322200010332023-3211112322321310-0110000122230212-3032101200033122-2133331222122030-3210102332133330"></a>

## Direct properties — business_logic_markup_setting / 131033103223 / 3

- [disable_spec](data-sources--app_setting--reference--group-001.md#canonical-0200312311112012-0100321113122321-0131123221110312-0310131020012232-1012222023202220-1320232010223003-1312211331013110-3113233102013202): complete subsection reference.

- [enable](data-sources--app_setting--reference--group-001.md#canonical-3213132122211021-3203223330310022-1022310133310001-0101301000002222-1223321221122312-3103012302223211-2101310212231020-3131331231203101): complete subsection reference.

<a id="canonical-1132322112313103-1301132033310333-0133101212102020-3000232312313022-3321330303303233-2033222202230121-0001213023303323-0303022121010330"></a>

## Next pages — business_logic_markup_setting / 131033103223 / 4

- [app_type_settings.business_logic_markup_setting.disable_spec](data-sources--app_setting--reference--group-001.md#canonical-0200312311112012-0100321113122321-0131123221110312-0310131020012232-1012222023202220-1320232010223003-1312211331013110-3113233102013202)
- [app_type_settings.business_logic_markup_setting.enable](data-sources--app_setting--reference--group-001.md#canonical-3213132122211021-3203223330310022-1022310133310001-0101301000002222-1223321221122312-3103012302223211-2101310212231020-3131331231203101)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0200312311112012-0100321113122321-0131123221110312-0310131020012232-1012222023202220-1320232010223003-1312211331013110-3113233102013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012203101333120-3213231203033022-2332211003202323-1201210231200300-0012220220213013-2230233311002333-3033033121331022-0303110003033001"></a>

## app_type_settings.business_logic_markup_setting.disable_spec — disable_spec / 002203000121 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- app_type_settings.business_logic_markup_setting.disable_spec

<a id="canonical-1303110210301321-1212020310121103-2032221102231120-1133230112213132-2332013330020001-1012133032122213-1230303131233033-3003231331030303"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0011212131213300-0311032210100201-3313230000210312-0203302230322021-0102202310011113-1122021110312310-3312020312030202-3321321231323313"></a>

## Direct properties — disable_spec / 002203000121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100133303021023-3223113032120130-2213112031313103-1011031331232111-2003322101100203-1201032220100203-0210211330130132-2332012301333100"></a>

## Next pages — disable_spec / 002203000121 / 4

- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3213132122211021-3203223330310022-1022310133310001-0101301000002222-1223321221122312-3103012302223211-2101310212231020-3131331231203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130233223132001-3232222223000220-1112220223022211-1330112201303133-3201123230213222-3202221102321120-3303121010002310-3131023200211200"></a>

## app_type_settings.business_logic_markup_setting.enable — enable / 012232033233 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-2230102112101031-0011323131130231-2132020203332122-1233121002130201-2022333311031102-3313210321021232-0221120203300133-0102112203100232"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2212130202033203-2130201013302223-1222021111230311-1202223332023302-3102223210022131-1133020000223200-1330031032113221-2131300032301200"></a>

## Direct properties — enable / 012232033233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300202203011332-1011303102201221-1203332201323311-3300230210111221-2001330222133302-0210020021102202-1321011332000233-1213102331323011"></a>

## Next pages — enable / 012232033233 / 4

- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000121300131331-1212313330302333-2301322103232331-0323102201123311-3012100102231312-0332102020023312-1133331210322101-2300220200221001"></a>

## app_type_settings.timeseries_analyses_setting — timeseries_analyses_setting / 121100202303 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-2130010012101131-3220202100123233-2030312031232330-2303302230132022-1322222312301232-0201031203022111-2113231330033210-0302103321220212"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1231003331233133-1203001331211322-2231223033202313-3010322302212003-3331300003233131-0133333002011000-0302001100032223-3320301201220320"></a>

## Direct properties — timeseries_analyses_setting / 121100202303 / 3

- [metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-1103312111103121-2030011023300232-2312013100020133-2120013122001110-0020112230003033-1110213302130202-1230331321100203-1130313110113230): complete subsection reference.

<a id="canonical-1232101330203000-2000303323120321-2013011110033311-3201322121022202-1300031323203323-1321303303111003-0111113131131130-1002312220202322"></a>

## Next pages — timeseries_analyses_setting / 121100202303 / 4

- [app_type_settings.timeseries_analyses_setting.metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-1103312111103121-2030011023300232-2312013100020133-2120013122001110-0020112230003033-1110213302130202-1230331321100203-1130313110113230)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1103312111103121-2030011023300232-2312013100020133-2120013122001110-0020112230003033-1110213302130202-1230331321100203-1130313110113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000230133320203-1230031102202331-3013122310332220-2002130100231132-2022010213311000-3031331133221111-2223311011320110-2020100130312311"></a>

## app_type_settings.timeseries_analyses_setting.metric_selectors — metric_selectors / 031112001200 / 2

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

<a id="canonical-1230300031210220-0301122202210330-3321000133203310-0010202123220013-3222133012210100-0102012112003202-0322200231231200-1303322122323001"></a>

## Direct properties — metric_selectors / 031112001200 / 3

<a id="canonical-3130233320322221-0332213112132230-0112331121210112-1311003222123221-2203130131021333-0030230220310310-1000102302103022-1101021110331122"></a>

<a id="canonical-0123032303120332-3022033202120301-3331223112213003-1000022102103310-3113312003103330-1100320021123232-1330320202112111-1102113110031322"></a>

## metric property — metric_selectors / 031112001200 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1321221311111031-3300011122330112-0000011303132333-1001220111013020-2033010132202101-1232013312132021-2002112302202201-1303021330212230"></a>

<a id="canonical-1202213010001302-1110331233030111-1013012031030210-0231231201100231-0233002032331131-0023002002333322-2321203000130133-2310200032030330"></a>

## metrics_source property — metric_selectors / 031112001200 / 5

Type: `"string"`. Computed.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

Upstream description:

Supported sources from which Metrics can be analyzed

All edges in the service mesh graph. Metrics are analyzed separately between all source and
destination service combinations.

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

<a id="canonical-2012003021313121-2120311203021113-0332132112322331-2211332101100123-0213100032201220-0111321110300200-3230011122222031-3222130301302022"></a>

## Next pages — metric_selectors / 031112001200 / 6

- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330330320232310-0313332220101322-3013013313200003-3002101131321002-3110311200000101-0302031301013303-1132110231112001-1120112332022022"></a>

## app_type_settings.user_behavior_analysis_setting — user_behavior_analysis_setting / 013123311021 / 2

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

<a id="canonical-1332211030133131-2303131121111101-3321231233011132-1103223202130012-0223330210112311-3112033020303100-2233020332303333-3330023131203002"></a>

## Direct properties — user_behavior_analysis_setting / 013123311021 / 3

- [disable_detection](data-sources--app_setting--reference--group-001.md#canonical-3001102003101201-3122033211102003-2030133123001012-2000322102011101-0211123203200300-0212301013122023-1031321130113002-3300121203232230): complete subsection reference.

- [disable_learning](data-sources--app_setting--reference--group-001.md#canonical-0002202122203233-3232101201122203-3200201220023003-2213000013012303-3030221023110200-3322121201321200-1222302020320220-3222001211202100): complete subsection reference.

- [enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301): complete subsection reference.

- [enable_learning](data-sources--app_setting--reference--group-001.md#canonical-3002320012003003-0030000001130112-0331132232021323-0313322303313030-0011130012333322-0032131111213310-0000323222133133-3301202030313010): complete subsection reference.

<a id="canonical-2233223232320203-1302233313110032-1200210030231121-1233213320113102-0123020303212111-1002112102220202-0031312212032212-1313100132113012"></a>

## Next pages — user_behavior_analysis_setting / 013123311021 / 4

- [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--reference--group-001.md#canonical-3001102003101201-3122033211102003-2030133123001012-2000322102011101-0211123203200300-0212301013122023-1031321130113002-3300121203232230)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--reference--group-001.md#canonical-0002202122203233-3232101201122203-3200201220023003-2213000013012303-3030221023110200-3322121201321200-1222302020320220-3222001211202100)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--reference--group-001.md#canonical-3002320012003003-0030000001130112-0331132232021323-0313322303313030-0011130012333322-0032131111213310-0000323222133133-3301202030313010)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3001102003101201-3122033211102003-2030133123001012-2000322102011101-0211123203200300-0212301013122023-1031321130113002-3300121203232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113002311203211-2330233011331033-3020332113232201-2330022323233233-2133312231132012-2100110132111000-3012320200022020-2203131233033322"></a>

## app_type_settings.user_behavior_analysis_setting.disable_detection — disable_detection / 210031112032 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-3333000030201032-0013220232210300-0201213031222310-1302120323110333-1020023332032020-3023212111030203-3112131030021312-1302331013211333"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0102021121311220-3120013022210203-3120312113100331-3300110112321211-3000323211323113-0333120333023032-3121123030300100-3021313120311232"></a>

## Direct properties — disable_detection / 210031112032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210023013312103-0313333120102203-2133102021121122-1121123200132120-0311320011220311-1112232331220211-0323132223310031-0201013322200323"></a>

## Next pages — disable_detection / 210031112032 / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0002202122203233-3232101201122203-3200201220023003-2213000013012303-3030221023110200-3322121201321200-1222302020320220-3222001211202100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120200003333132-1111302110031202-2303331333102321-1100013023022103-3011322303013132-2100212032132000-0321302123333212-2130231132112301"></a>

## app_type_settings.user_behavior_analysis_setting.disable_learning — disable_learning / 032223013033 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-2022033131001212-2212010100111222-1120303002210002-3301330013010320-0000133002101333-2330003333131232-2330333223033233-1300002021230330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3322102013223301-3220321122023221-1112101233233130-0001003212202333-0210323312310032-1201033232001313-0201230130012032-0100310111122330"></a>

## Direct properties — disable_learning / 032223013033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223001121000033-0220321101132001-3023210023202123-1233200121102112-2003100112232231-0312222011102312-3330222031302133-3201020332122020"></a>

## Next pages — disable_learning / 032223013033 / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023002031202311-1301120023310220-1102321133222032-3101322301032031-0300230201121030-2130231130022010-2023321220232212-2031022111211001"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection — enable_detection / 300333031323 / 2

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

Upstream description:

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

<a id="canonical-2111012211221210-1210101203222222-2333113301000131-3220021132323211-1023130131221311-0333332001201001-2100230031003322-3110322000110322"></a>

## Direct properties — enable_detection / 300333031323 / 3

- [bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-1003323012123303-0311001311133112-1122123013113233-1102002230232223-1220001032230233-1131121213111333-3330211230232331-3101212323032031): complete subsection reference.

<a id="canonical-0333323103032112-0331221230113013-2101120210021003-0331013222032032-2303031132230123-3031030000013112-2013322030332201-0032202323301212"></a>

<a id="canonical-0322013113002013-1202231123303330-0300300112111113-0221221232100133-0011131021103331-3032130100130021-1020321023223011-2130032130232233"></a>

## cooling_off_period property — enable_detection / 300333031323 / 4

Type: `"number"`. Computed.

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

<a id="canonical-0002221323232301-1122030113313130-2220303211212230-3301131021321233-3110301023133321-2131032220231222-1301313113303311-2232211111321132"></a>

## Next pages — enable_detection / 300333031323 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-1003323012123303-0311001311133112-1122123013113233-1102002230232223-1220001032230233-1131121213111333-3330211230232331-3101212323032031)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-1021120211131220-2110223313300012-0000231023223103-1133302231111003-3103230023022132-0331331202302111-1011013010002302-1131131322210011)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-1100033330102233-2012303200223112-0300232222222123-1010003023133013-0122010331113220-3233012333203310-0323120101110222-1202313200101221)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0002122213313103-3032133310131230-0122100301113022-2300033021010121-0001012313111230-3100213022023000-3012131110022113-3212223011111031)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-2331033302323322-2312213221000000-2131231022021231-1003123130113231-0000001120032300-1311333122201000-3213103233023231-1021021333100103)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-1113211312303312-0301111101121123-2333013203310323-1320120023120111-2130313200130232-1032301012010230-0131213131200223-3133013111131213)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-3210100300120111-3132233223112133-3002111231313013-0032033313111111-0120311100231323-0333031002122001-3033023230103131-2300303111032230)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-3030332233112232-1133033211233331-0120323000303113-3132311201220213-1112032230222313-0213010032232010-1203230332300231-1312211311211223)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-2213012102133002-3220010203131211-2220232111102213-1103023130033121-0012200212113030-1110300103330031-2222012032231032-0120033301310233)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-3320120103323121-3003201202010010-0212000202130023-2211113030101330-0232030132023313-0023121032331230-0222003103231230-3222022031322000)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-3303321120111303-3001302033300222-0222230201200300-2103032200130311-1300121221302320-1211211022102132-3333302201033221-2313031223000131)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-0110231030312011-0111202223331123-0100020033012112-2333120220203123-0311203000130322-0223202333322213-3201312101123100-1322022301223033)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-1113332013020121-3220203212302100-2032013132131110-2312000021021212-1112023022133003-2021332321011002-1321101132010211-3101302110213130)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-1210311223301311-3213130103103302-0223213231222233-0023213320210120-0211111122112300-3132201323223322-3300102011333100-0013202310211030)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-1230131110313032-1010110212231113-1120321331001111-3100010211120222-3021332030310333-3330110113022003-2111211320011033-1301123230303221)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1003323012123303-0311001311133112-1122123013113233-1102002230232223-1220001032230233-1131121213111333-3330211230232331-3101212323032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131010231332121-1131300302331331-0323211000101123-0200013323131212-0031101030200011-0310220100232200-3320020102003202-0013012200103221"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic — bola_detection_automatic / 021323112223 / 2

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

<a id="canonical-0200303232323010-2122321212111003-3333133332110203-2012032231213010-0002021231321023-3322200032113013-3133210310103112-3303220233231201"></a>

## Direct properties — bola_detection_automatic / 021323112223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032230221200222-3000010201131310-1110130312221022-2220021210101123-0133322302301021-1231310103212223-2303211303200131-1121121123020110"></a>

## Next pages — bola_detection_automatic / 021323112223 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1021120211131220-2110223313300012-0000231023223103-1133302231111003-3103230023022132-0331331202302111-1011013010002302-1131131322210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002232333102221-2210231023303113-1322300233223230-2303222000033213-1000332211330130-3012112033132113-2032321201330333-3303133101211300"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection — exclude_bola_detection / 201203332003 / 2

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

<a id="canonical-2233301302201123-3323313132101203-1133313120232110-2212111111221322-1323013101313203-2131233012020210-1312221320032010-2032320213201230"></a>

## Direct properties — exclude_bola_detection / 201203332003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313211011330133-0231210032223212-3202333110103301-2102213323310330-3331022131313102-2000102121230201-1223121130303121-3010202113322213"></a>

## Next pages — exclude_bola_detection / 201203332003 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1100033330102233-2012303200223112-0300232222222123-1010003023133013-0122010331113220-3233012333203310-0323120101110222-1202313200101221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022023001121111-1222021322212201-3332320333133323-3333121013110031-1102232231231132-1330132013012202-3133031131103230-3302210121013311"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity — exclude_bot_defense_activity / 322113102313 / 2

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

<a id="canonical-3020223032130130-0212202133032032-2230202201213133-1100231100010101-2320210002111012-1220233003301131-0332220023201300-2313222031103002"></a>

## Direct properties — exclude_bot_defense_activity / 322113102313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202130020211133-1131001312030210-2312123112121120-2022113023011201-0202022111320031-0103222112113111-0102102130322112-2200002103302021"></a>

## Next pages — exclude_bot_defense_activity / 322113102313 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0002122213313103-3032133310131230-0122100301113022-2300033021010121-0001012313111230-3100213022023000-3012131110022113-3212223011111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113320211121033-3031303113132131-0012001200331303-2322031330332232-0300302231013003-0121231223201201-0231200323022121-0231030010333021"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity — exclude_failed_login_activity / 212203133023 / 2

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

<a id="canonical-2011113001123302-3121022200022012-3330221320133112-3301321020203213-3013322103032233-3312223323212220-1031010112222111-3003021122113330"></a>

## Direct properties — exclude_failed_login_activity / 212203133023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102320032331001-1100320012212123-3330202102321200-2313032121013213-1012331231012010-3312333023212122-0113203220320302-1322020213003202"></a>

## Next pages — exclude_failed_login_activity / 212203133023 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-2331033302323322-2312213221000000-2131231022021231-1003123130113231-0000001120032300-1311333122201000-3213103233023231-1021021333100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310130213330301-3013133311233313-1033000003121202-1331011000212013-2133211303323302-0301030332303002-1232112212033001-2103113030312212"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity — exclude_forbidden_activity / 110212111223 / 2

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

<a id="canonical-3201033211131222-1331301133121031-0320331331131303-0133002320310322-1311032210003033-2011323012332020-0132333010200233-3332110113123032"></a>

## Direct properties — exclude_forbidden_activity / 110212111223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022302012110030-1002011330011223-0001320113112003-0030331303033000-2210332011313222-3303030021211300-1110002003212200-1113032230203101"></a>

## Next pages — exclude_forbidden_activity / 110212111223 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1113211312303312-0301111101121123-2333013203310323-1320120023120111-2130313200130232-1032301012010230-0131213131200223-3133013111131213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311112020130031-1231021333220231-0103200223322210-3331212310310201-1022032012021300-0101220203033030-0001200220010013-2313231303003012"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation — exclude_ip_reputation / 023333121101 / 2

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

<a id="canonical-2313213110210123-2132131033221232-3212200321310030-1212023032010003-1133211200321333-0133113132222120-3123103103330000-2311302102021333"></a>

## Direct properties — exclude_ip_reputation / 023333121101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033132230210001-1302102311111232-1310302223030313-3313200202210100-2301312333130332-1210333131110021-0330133000130223-3321322302030010"></a>

## Next pages — exclude_ip_reputation / 023333121101 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3210100300120111-3132233223112133-3002111231313013-0032033313111111-0120311100231323-0333031002122001-3033023230103131-2300303111032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331301121112302-0111131311120300-1221003111303011-2202222313101221-1211210030030101-3003112232313122-3100123303012302-2311320310110232"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity — exclude_non_existent_url_activity / 121213130210 / 2

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

<a id="canonical-1303210222230201-3201202301330200-3032303022221011-2101200112011303-3310321220120333-1103201122232133-3022232213222121-2313021013312311"></a>

## Direct properties — exclude_non_existent_url_activity / 121213130210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303212203010203-1010122230113131-1010101131230320-3023300221330320-1101133302230301-3332030121320321-0132331000113033-0112121113131200"></a>

## Next pages — exclude_non_existent_url_activity / 121213130210 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3030332233112232-1133033211233331-0120323000303113-3132311201220213-1112032230222313-0213010032232010-1203230332300231-1312211311211223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232220100010101-3002130323322323-2332103002120311-0033202212111212-0111201201202011-1221032133313321-3302200203302232-0000003332121112"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit — exclude_rate_limit / 121112332322 / 2

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

<a id="canonical-2013301012022312-0002112103100330-3131030321313321-3133111221233313-0303331221332012-3101210300230112-3233111301113300-0133311221110101"></a>

## Direct properties — exclude_rate_limit / 121112332322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212110210333001-1331333323201031-2333222321122123-3112323033022102-1210121323212201-3032003033020232-1112001331301100-0212222112211121"></a>

## Next pages — exclude_rate_limit / 121112332322 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-2213012102133002-3220010203131211-2220232111102213-1103023130033121-0012200212113030-1110300103330031-2222012032231032-0120033301310233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112132330130021-2200112112232320-2230122101321301-2322313330023202-3301302213312102-0100111220011333-3103220220311102-3133110213131212"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity — exclude_waf_activity / 201023320301 / 2

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

<a id="canonical-0221220031101201-3132122020301133-1012101303013231-3323301331232112-2203132233100211-1203102313021120-2010201101131122-0001233112311323"></a>

## Direct properties — exclude_waf_activity / 201023320301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312021030312130-1031233301121132-1331012032301332-0312032202132220-2100002102132010-0122112221023223-3021333122330330-1212312101223123"></a>

## Next pages — exclude_waf_activity / 201023320301 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3320120103323121-3003201202010010-0212000202130023-2211113030101330-0232030132023313-0023121032331230-0222003103231230-3222022031322000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231022211313202-2231120130211312-1000000030110210-0033122132310221-0003012332321001-3221331200022211-3210311030323321-1130232313320220"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity — include_bot_defense_activity / 033201300110 / 2

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

<a id="canonical-1123020031122323-2320000102220312-1032312221312112-2233313223201112-0112110102330023-3223031322133311-2003202201212101-1222231303311001"></a>

## Direct properties — include_bot_defense_activity / 033201300110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320122322231221-1023221011020220-3220031303030003-2222120033003031-3222023322032220-1002222320220102-2003330203313101-3323132203030032"></a>

## Next pages — include_bot_defense_activity / 033201300110 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302010211032202-2003233012313230-2311013103211333-2302230122212213-2120331023012231-3321333210310033-0123202102130020-0223202013233312"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity — include_failed_login_activity / 023230103003 / 2

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

Upstream description:

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

<a id="canonical-0100000302120112-3210112001320203-0112111113322202-3102223333222213-3031120231113103-3202023111033331-2001220022130122-2210122311130100"></a>

## Direct properties — include_failed_login_activity / 023230103003 / 3

<a id="canonical-2023230020332303-0111332110003232-1031233312100302-2022201323330131-1001022013010021-0102031113300320-3133003323120301-3101011022233032"></a>

<a id="canonical-1112000202103033-2031333332021030-3331010213112030-3302023310333202-0131021212231323-3312200211310220-3201233313013330-1003111021131322"></a>

## login_failures_threshold property — include_failed_login_activity / 023230103003 / 4

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

<a id="canonical-3130303011122121-0130210131101031-1302320312001023-0313300132211321-0022312002023030-1020322310212132-3302311101330010-1020020321030300"></a>

## Next pages — include_failed_login_activity / 023230103003 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3303321120111303-3001302033300222-0222230201200300-2103032200130311-1300121221302320-1211211022102132-3333302201033221-2313031223000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221213220321231-2212303302132113-3101002302321130-3020302133122011-1323201100002202-1213031131230233-1012122333033201-3200223233102011"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity — include_forbidden_activity / 330301311023 / 2

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

Upstream description:

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

<a id="canonical-0332223112122031-3101021111310322-1000123210322212-0110121311221112-2012102220322130-0102221002202211-1310312001011101-1332010221111320"></a>

## Direct properties — include_forbidden_activity / 330301311023 / 3

<a id="canonical-1332021322322200-2022200020213210-2023030010321300-2000130302320321-2011003213020101-3311030020133332-2112312311002101-0210213011032303"></a>

<a id="canonical-1131003213211210-1332303022032302-0113201322012020-1000332212323211-3222000123120201-0332222202221011-0331032300012113-0213202111003133"></a>

## forbidden_requests_threshold property — include_forbidden_activity / 330301311023 / 4

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

<a id="canonical-2303112311222322-2020231220201223-0301310302213000-1101121131111102-2122203311201233-2110332330301120-2103120030033100-0302233120323333"></a>

## Next pages — include_forbidden_activity / 330301311023 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-0110231030312011-0111202223331123-0100020033012112-2333120220203123-0311203000130322-0223202333322213-3201312101123100-1322022301223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203111132211030-2001320110123233-3330233020320310-3110022212221211-0111302213200200-0233211112323022-1310300300332020-0311202213332313"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation — include_ip_reputation / 101210002213 / 2

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

<a id="canonical-2310031310333111-0222212021203221-1120102032121112-1200222002110301-1313200220233202-2211133103133310-1210013322212231-2202012032231320"></a>

## Direct properties — include_ip_reputation / 101210002213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331221002313203-3021032022010013-2132231201100133-0213001111011332-1222222121111331-0030132033121202-1012311011211302-0111023101111333"></a>

## Next pages — include_ip_reputation / 101210002213 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230312332013000-2011312002201130-1112123203012221-2203210330112122-1232210202122201-1212120030011002-0331003230022000-0033210113311210"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic — include_non_existent_url_activity_automatic / 221311111011 / 2

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

<a id="canonical-0032312312230213-2220203031322132-0120012303322332-0311122021102221-3133323021203322-2223030002122112-0001000333102301-1111123321002233"></a>

## Direct properties — include_non_existent_url_activity_automatic / 221311111011 / 3

- [high](data-sources--app_setting--reference--group-001.md#canonical-1000120102030200-3100112023223130-3320313213330012-2100310331303111-0210320233220323-1123232311123002-0102123032122311-2221100103010223): complete subsection reference.

- [low](data-sources--app_setting--reference--group-001.md#canonical-2001232101233222-3220321222002323-2233332332323121-2102000031320330-0123222002211320-3102112211312122-3222221312312202-3232211131202212): complete subsection reference.

- [medium](data-sources--app_setting--reference--group-001.md#canonical-1120301301313110-0223320213030232-1132320332002332-2330232201221011-3211112312320333-2131311120010323-0130311302101303-1031032213301303): complete subsection reference.

<a id="canonical-2211123032233333-1021212220102232-2320232232231203-3122213200030213-0212332111111322-3100330200022210-0323131231133333-1022122301112010"></a>

## Next pages — include_non_existent_url_activity_automatic / 221311111011 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](data-sources--app_setting--reference--group-001.md#canonical-1000120102030200-3100112023223130-3320313213330012-2100310331303111-0210320233220323-1123232311123002-0102123032122311-2221100103010223)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](data-sources--app_setting--reference--group-001.md#canonical-2001232101233222-3220321222002323-2233332332323121-2102000031320330-0123222002211320-3102112211312122-3222221312312202-3232211131202212)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](data-sources--app_setting--reference--group-001.md#canonical-1120301301313110-0223320213030232-1132320332002332-2330232201221011-3211112312320333-2131311120010323-0130311302101303-1031032213301303)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1000120102030200-3100112023223130-3320313213330012-2100310331303111-0210320233220323-1123232311123002-0102123032122311-2221100103010223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131311330321321-3232012211011220-2002112302123302-3332201022013012-0110112203200232-2012211132213312-3232110011300202-0021012322301021"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high — high / 103221121201 / 2

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

<a id="canonical-0332213320130000-3003321220030303-1222201203033322-3222110313010320-1111313130210111-2001310101202230-1210122330013223-3000213301203211"></a>

## Direct properties — high / 103221121201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221321022221230-2001031323331230-0330230223122002-1322300212332011-0021011230221210-2212133001121212-3122113222202112-0021031211032220"></a>

## Next pages — high / 103221121201 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-2001232101233222-3220321222002323-2233332332323121-2102000031320330-0123222002211320-3102112211312122-3222221312312202-3232211131202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220301122220010-3312221020201003-2233333311122212-2302111012210110-2300032103012103-0212032031330311-1132201013223100-2031301031111021"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low — low / 313133300321 / 2

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

<a id="canonical-1032101322002332-1131033210303100-1312323120122030-0011202021322100-1330131003230220-2300013110002013-2032222211023023-2301302131121222"></a>

## Direct properties — low / 313133300321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220203322121013-1031230102203221-1122033302301011-2030333201232301-2200300131232013-3221001231100101-1103303333320121-1332012311112032"></a>

## Next pages — low / 313133300321 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1120301301313110-0223320213030232-1132320332002332-2330232201221011-3211112312320333-2131311120010323-0130311302101303-1031032213301303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010213033212220-3321213100201303-0023313222320332-0310010013211022-2303212222101003-0202331202201313-3110120203221013-2301321112113120"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium — medium / 110031231322 / 2

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

<a id="canonical-2113110103220310-2223223123312022-3101131302002321-1021313032230300-0303323223100221-2110303210203011-2201322003203022-1132110213013001"></a>

## Direct properties — medium / 110031231322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233010231013213-3330233121102012-3223223103013022-2102323123111233-0113102020113230-0311333210322211-2333300012230100-0211312113303101"></a>

## Next pages — medium / 110031231322 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-3101100101121320-2121113003103220-3220313211132001-3212013131231222-0023312033033132-2112301033330131-2020202012323122-0232110232132110)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1113332013020121-3220203212302100-2032013132131110-2312000021021212-1112023022133003-2021332321011002-1321101132010211-3101302110213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331010330112313-0103332212202301-3011321120102033-0130103023030220-0302221111000110-0030012200320203-2021111303113012-2301111310303003"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom — include_non_existent_url_activity_custom / 112213322210 / 2

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

<a id="canonical-2203120032003021-0121300001111311-1312230000111223-0323123301333312-1112203330030001-3121333031121333-1022022012013033-1130022110122002"></a>

## Direct properties — include_non_existent_url_activity_custom / 112213322210 / 3

<a id="canonical-3033101031332120-1332000112000112-0022131121300031-0010201121323302-3301311330131113-0020120320313121-0022321202222301-3203113302030010"></a>

<a id="canonical-3112222232000031-3010131013110203-1002323012320211-3320101213120001-0222322002033132-0111321120123300-0330132120121303-0200233200100303"></a>

## nonexistent_requests_threshold property — include_non_existent_url_activity_custom / 112213322210 / 4

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

<a id="canonical-1013212323222131-2223133111321200-2303011220302110-3032020233132213-2200320100233120-2012013221000130-0233232331003331-0102223022132103"></a>

## Next pages — include_non_existent_url_activity_custom / 112213322210 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1210311223301311-3213130103103302-0223213231222233-0023213320210120-0211111122112300-3132201323223322-3300102011333100-0013202310211030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203301120033230-2222022223332320-0233111022121000-3010031011003303-1222313330203222-3212213002303031-3133113322131123-2131223010213133"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit — include_rate_limit / 230033012311 / 2

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

<a id="canonical-3301230103133311-3031203110122321-0122311333133302-2321333311113331-2330332030303200-0001011313111133-1313322033013232-3011230200131132"></a>

## Direct properties — include_rate_limit / 230033012311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002322310210111-1101013132220111-2012033012023231-3231013300130031-0320113111201301-0210233013021211-1123122103132222-0131202210132012"></a>

## Next pages — include_rate_limit / 230033012311 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-1230131110313032-1010110212231113-1120321331001111-3100010211120222-3021332030310333-3330110113022003-2111211320011033-1301123230303221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111103021100123-2320113203020213-2033323321120231-3003100032212310-3311023033201331-2311131022010101-1322232330232201-0113033000232010"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity — include_waf_activity / 123203013312 / 2

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

<a id="canonical-0013230103302002-1211303332333133-1303322330120200-1110032033021131-3230322231111122-0231330120311100-2200201322001113-3332011022133322"></a>

## Direct properties — include_waf_activity / 123203013312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211102303330211-0303133023233221-2012112003132111-3011322330002023-1323113221221012-0023300333101333-0111321320300202-2111202322030022"></a>

## Next pages — include_waf_activity / 123203013312 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)

<a id="canonical-3002320012003003-0030000001130112-0331132232021323-0313322303313030-0011130012333322-0032131111213310-0000323222133133-3301202030313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123020011030223-2000100222301100-0103103120110123-3332113313232210-2131113133333021-1230122121130203-0103232013212303-1013002100032310"></a>

## app_type_settings.user_behavior_analysis_setting.enable_learning — enable_learning / 321020230022 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-2122100133300012-2300300302110303-2303313330010132-3232220103310210-2030202233010122-3230213112231202-0312010332321231-2202203321202133"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1102110000202213-3202302131222001-1331213322330131-1320100020303032-2321002321031213-0022110320013321-2322233213032322-3200303332010201"></a>

## Direct properties — enable_learning / 321020230022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133332132123031-2331003212323231-1222020231013133-0023021303223031-0221001332033133-2012332012033112-0030100230003212-2203102303322221"></a>

## Next pages — enable_learning / 321020230022 / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-3310333331000202-1113310222302103-1203203133133101-3201122201220130-2232112101121012-3331332113331030-3202212130003030-0012212001312002)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021)
