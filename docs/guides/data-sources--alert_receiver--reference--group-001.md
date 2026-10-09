---
page_title: "xcsh_alert_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver reference."
---

# xcsh_alert_receiver reference

<a id="canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- Property reference

<a id="canonical-2333232112120122-1320020233222231-3113302202003022-1203213320001033-3322133010122331-1302310200031213-1012022302200022-3233321113103230"></a>

### Direct properties for `xcsh_alert_receiver`

<a id="canonical-1002013232220131-2200103113213122-0020012321021003-3100333120213323-3203021223103322-2130321102111132-0221131310310031-0133203031210011"></a>

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

<a id="canonical-1013000032110200-3210203130131121-0201111131201220-1030331022012223-2311002102123101-2202201110022333-2132132032113030-0201200222103222"></a>

<a id="canonical-0231201232022303-1120303201031303-2032122111121220-3122232323222200-1000302112201012-2233030011331313-0302301232013130-2011230110030130"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AlertReceiver.

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

- [email](data-sources--alert_receiver--reference--group-001.md#canonical-0201220120220213-2102011031101202-1102230231331111-2220122101301312-3030220000230030-2233120122013213-3232011302301001-2102231023110000): complete subsection reference.

<a id="canonical-0131210231333210-0330030332210020-3103302030102132-2000302230322233-0223033230121000-3332332323320321-2133212320222103-3321221112122230"></a>

<a id="canonical-3012133313300020-0301321002031321-1120130120110001-0103203232321132-0211120302222032-3003103212201330-0323111032232022-2102222123301111"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3033301132101222-2312132313113030-3312111011221211-2331030233230111-2203000331231020-2221312311321301-0113133122131220-1220312101122130"></a>

<a id="canonical-2322221131200302-0322301221211313-2103233023323101-3332131030211031-0233332001321210-0131121200232330-1223212020300131-3013130103020111"></a>

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

<a id="canonical-2301323011120200-1002003121000213-1133223313122302-1320303022130313-1130012303101100-3300311012303003-3323111222022013-2022112330032111"></a>

<a id="canonical-1321200020110310-1233312001033320-2203212001202211-2132002112220010-0330001120302323-3030213303112000-3113330213100020-3011220320000010"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AlertReceiver.

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

<a id="canonical-3312231202103202-2012130011200113-2101220303311330-1012111311201102-1212102112022313-3332232323202100-3012200300102010-2133213321132232"></a>

<a id="canonical-2103321320303032-2332011312212321-0020333302013033-0031313313323221-2000301001131101-1212313100011112-3231332010300231-3221103021200012"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AlertReceiver exists.

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

- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-1211322113200310-1213220213023020-2232330133023322-0012212203103221-1123201022032210-1310021300010023-0123031121310201-0132303120000123): complete subsection reference.

- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2310320030111211-3032330222323111-2003331131121231-2021223211110101-3221203112203233-3022010211323123-2003032312303302-0111312030013211): complete subsection reference.

- [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323): complete subsection reference.

- [sms](data-sources--alert_receiver--reference--group-001.md#canonical-1102020323302231-3330031312223101-2332200311130100-0221003321130033-0310033330311110-2222220310130220-0311300330010121-3122100310111130): complete subsection reference.

- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303): complete subsection reference.

<a id="canonical-3322030022321123-1031202212103220-1321322112133223-2332323331032312-2222010213330120-0110012203102300-3011003312103000-2002323203033302"></a>

### All schema paths for `xcsh_alert_receiver`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--alert_receiver--reference--group-001.md#canonical-1002013232220131-2200103113213122-0020012321021003-3100333120213323-3203021223103322-2130321102111132-0221131310310031-0133203031210011) |
| `description` | [description](data-sources--alert_receiver--reference--group-001.md#canonical-1013000032110200-3210203130131121-0201111131201220-1030331022012223-2311002102123101-2202201110022333-2132132032113030-0201200222103222) |
| `email` | [email](data-sources--alert_receiver--reference--group-001.md#canonical-0310220320302133-0001131301031010-1023022303031230-2330231011323101-2122102210132322-2101013330202312-2221201110320200-1131121332022223) |
| `email.email` | [email.email](data-sources--alert_receiver--reference--group-001.md#canonical-3023230221321012-0331013210111010-2300003202013213-2123100223201323-0122321010022211-2213320201133321-0113211100122212-0201332333001123) |
| `id` | [ID](data-sources--alert_receiver--reference--group-001.md#canonical-0131210231333210-0330030332210020-3103302030102132-2000302230322233-0223033230121000-3332332323320321-2133212320222103-3321221112122230) |
| `labels` | [labels](data-sources--alert_receiver--reference--group-001.md#canonical-3033301132101222-2312132313113030-3312111011221211-2331030233230111-2203000331231020-2221312311321301-0113133122131220-1220312101122130) |
| `name` | [name](data-sources--alert_receiver--reference--group-001.md#canonical-2301323011120200-1002003121000213-1133223313122302-1320303022130313-1130012303101100-3300311012303003-3323111222022013-2022112330032111) |
| `namespace` | [namespace](data-sources--alert_receiver--reference--group-001.md#canonical-3312231202103202-2012130011200113-2101220303311330-1012111311201102-1212102112022313-3332232323202100-3012200300102010-2133213321132232) |
| `opsgenie` | [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-2101202022133322-1011003311212231-2331103013233003-3233221112202013-3011111111133130-0201312203221011-0301320113212001-1100011211202213) |
| `opsgenie.api_key` | [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-1130032232020101-1012103323103230-3310320101201131-1223200223211120-0302131002331111-2110213011130102-1201221020032322-2131000213323313) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3300302123231322-3023320300223021-2332321012331301-0103313213011113-0102123222202110-0013011000310113-0002322221031222-2313002233320012) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-3321323223131331-1221223202300001-1013013133323333-2221103220001322-1310120223312231-1321301311203313-1120310330310131-0132033033200203) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-1132022131302210-0003120303300303-1020301020220211-2302221311113331-2121110103213313-1031021303101031-2313312313211330-1330113103312223) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-0333021111031133-1300300303020231-3120000102300311-2200011001201123-0032120002020311-0322002210011102-3230021302311020-2200232332233211) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2300211323303221-3332202033022320-3202301231201221-2312222121231013-3032020213000012-0111113223101222-2222122301212001-1122131122203131) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-1311021301210303-1230203312211203-2220131221122113-3001001312032102-1232331322300022-2201323111322221-2022212130000113-0210003100301333) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-1331001133111233-1101322213021223-1300333301323030-1033312012121212-1022302202223112-2300312303222203-0101010211330012-3000210033013310) |
| `opsgenie.url` | [opsgenie.url](data-sources--alert_receiver--reference--group-001.md#canonical-3033200133010202-2322203113111232-0222222301303113-3330022033320120-1312201103202233-0202213202333330-1300021231001221-2302002203210323) |
| `pagerduty` | [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2112113212110212-0001231112012103-1132210002201220-2122211111010302-3201022020221000-0231220220120221-3210331010011021-1020023102002231) |
| `pagerduty.routing_key` | [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-2031321110220332-1111032111131220-1101033200323233-1222130231301302-3122203111112000-3202131023021020-3120122101201130-1313313331101233) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2031132113330330-1031111112333310-2310323313032103-0022022312000122-3230330231321322-1122322320020300-3302310231200201-1231331331210010) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-2130130012022123-0322123002100222-2013032112010233-3233002120103312-0311311212201233-3011321002132032-3103002111101230-0113130111102300) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-3320213301002133-3011203010131130-2201313231202020-3320211020213110-3202111203232330-0333100323220232-1321020020303112-0233002331122130) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-3120201201112323-1332201331332001-3211120210321102-2332213003221021-2101311201021231-3011110310022122-2002132221202322-2311132331000132) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2312010200333332-3300213230121231-0022002001222213-0002201310112123-2220021212130020-0013333313321020-2232021022120002-2213311022302033) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-1210331333211320-3013302023022103-0231032332002013-0310321232222311-2222220303330232-2313221032133230-3020012011001032-1112211322213232) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-3201020320013223-1330232011300031-2332021032321020-0212303030320213-2021132310313023-0302212232200002-2110013020210001-3102010200301301) |
| `pagerduty.url` | [pagerduty.url](data-sources--alert_receiver--reference--group-001.md#canonical-2203330312013112-3310211323103220-3323113031213313-1201123330202133-0013221330310130-0123310333011230-1032222302122111-1213103200013203) |
| `slack` | [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3102222130123210-3312210300201322-1010302131033232-0301233120313102-3213211130200330-2303213000202201-2213301023213202-1332122033032312) |
| `slack.channel` | [slack.channel](data-sources--alert_receiver--reference--group-001.md#canonical-2233013223212231-3230301102032222-1100123200013030-1111132013010101-2213100020331310-2032230330032102-3310331003311022-2320322122211032) |
| `slack.url` | [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-2102301112232111-2323123321232310-2321203110010021-3323333130320021-3010200020102320-0311220021303210-2123033020202012-0212022200012102) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-1123321313130201-2101203110331012-1011103013331323-3232232331311123-3033103233123200-2103331122022303-3300213210110332-2033121332131011) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-2133110003020230-0212310023001331-3032111331102201-2123320332320031-0332203033332021-1312303112221102-0103112202323112-1201213002211311) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-1321100313113321-2230021102203030-0101212130111011-1322310322030002-3230012000213330-3231032233103220-0011311320132013-3102222231111103) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-1013122100332110-0221000222223211-2112320230222000-2131121200321223-2331132023012303-2101031032223201-0120131033203213-3000011111313213) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-1302032020021000-1310130220133030-3211301013020132-1100223223033101-1030220001110033-0120122113003102-1230203202003121-2301131200211000) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-1332110233003311-3110110110112202-3100021330133223-1211210312301011-3021311203230313-2221130001012022-0021000002201033-2103321302120103) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-3033031201231130-0031312132003021-3300220203312132-1313320230020323-3003002010211131-0003033310132111-0131310312331310-1330321320022132) |
| `sms` | [sms](data-sources--alert_receiver--reference--group-001.md#canonical-0202013131321302-2300223130301112-1302213220133131-3003322033201200-1103303102233312-2113230103122220-0233021112120030-2223320323002103) |
| `sms.contact_number` | [sms.contact_number](data-sources--alert_receiver--reference--group-001.md#canonical-2311302303011123-0122312302031032-1331000122010120-3121101110211031-0010230220031000-3031021302111332-1322310333033323-0321122003230213) |
| `webhook` | [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0022101022221202-2100222333201221-2130213303213223-1022012233102212-2200210201130012-1012110220331022-0013311310233032-3112330303331123) |
| `webhook.http_config` | [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-3232100101023232-2221112313023231-2210012122022332-3330310321021210-0221102210132213-2020302031000133-1312223213023020-1002201300211032) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-0133130101020210-1233130111101013-1010330230203003-0020123101203212-3212132322113230-2321211303231201-0123320302131220-1102022010102303) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-2010000121233200-1221020002023011-1312302112010212-0120301203302212-3011222320222203-1131333111001003-2101323222000123-0103011312102210) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-0131021031000112-2102131223201102-0103122230011010-3210311233300332-2312302013110202-0303302303122023-0011210030112310-2112313020311221) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-2103010331111022-0103020301023130-0001220322020021-3322221111022320-2233033212013311-3023032000022002-0023312211230102-3020231322313211) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-1201113100232331-3303012021000211-3313011313200020-2100310210222213-1112130121233301-2331022130132333-2332023010312133-2122030111023310) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-3132320333310211-3012220110230213-2221300200330122-0311021333302112-2130201022301030-0023123122132223-1311130002210200-0230033312331302) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3210220221122212-3000312032222023-0113202211302003-3103333101223220-2001310202320023-3302023120012122-0111010001000131-3223113213303233) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-2231100203132000-3033120001320020-0031303321100300-1013011121023120-0122210200131301-3313331333332133-1212220201233330-2100021013001132) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-1303312313331233-3132301021220102-0032131321333102-2112233013022022-1001321303302321-0111331022100121-2203121331131110-3011113102131130) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-1230301233221213-0112313201312310-1113102210231301-3301302300203112-0023123101111013-3132202212023021-2223010131221133-1130321122030212) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-0331011202133100-2013201120131010-3103101110323100-1001023232122030-2031311202102023-0113230010211323-1200213130200222-3232010032131121) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-0333122222310333-3333223101133211-1221000133222010-3120102332133021-0311303310201130-1022323203220000-3121010110311100-3113102223200003) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-1230110022211030-2311132202221330-2133203010310330-0030033202303330-1310321330121032-3101303222330223-0200130330023302-3030032300122232) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-2121321321112333-1201223202231001-1212200120003131-1002001003030121-3121231212232133-3003320101322023-1111023201300131-1003220133321123) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-0011020200021222-3013211322212301-2013112213201121-0212211312202230-1131002213121003-0011030221033222-0003323331030013-3110030103212313) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2131223031100210-0020330202231010-0332103133232201-2221300001321112-0111302121232330-1100131010132223-2200301203102221-1211102200003320) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-2123021030011013-0323330121322231-1122001330213000-1132233121220120-0112311310331011-1213110132232110-2101323210033000-1120112131102331) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-1020300120302312-2313312322333103-2312132310113101-0312220011302220-3300030320111331-3312301113113112-3022333031023013-3120101201311333) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](data-sources--alert_receiver--reference--group-001.md#canonical-0010123102232110-3012222320333000-0133221313212201-2312030021332103-3312133212323023-2011002222212220-1032003112203012-1213221122330311) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-1332233003013210-1101013202331111-1331213222131332-3322020130130023-1202000121311101-2021000232012302-1002212233232222-1123130313222030) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](data-sources--alert_receiver--reference--group-001.md#canonical-3320200102121123-0131002212230120-1220330000203323-0102001022210300-2300332331023322-2033313310323302-2201222123120131-3320310013103102) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](data-sources--alert_receiver--reference--group-001.md#canonical-3233123031321213-3203121321301302-1112002201110133-0233012201220202-1003320010000030-1013011123113033-3122030311121221-3033033023301131) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](data-sources--alert_receiver--reference--group-001.md#canonical-1010111201222233-0302322231020311-1321211323333333-0011132310233032-1000030102033022-0111101323200022-1233230212301020-0302310020233301) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](data-sources--alert_receiver--reference--group-001.md#canonical-2102132130301110-2111222310011011-0321011311330332-3303323300222202-3101221210100203-1101230221213222-1200020111130100-3302230020013233) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](data-sources--alert_receiver--reference--group-001.md#canonical-1332113001230301-3312132303101120-3010113320022222-1132012201031200-3201111110101330-0111021323030222-3303213222033323-3201221011320322) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](data-sources--alert_receiver--reference--group-001.md#canonical-3021131022303032-3332331323133200-1101302032131201-1122003122230200-3133210332013110-1321031100303021-1133031233230311-0112312113201111) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](data-sources--alert_receiver--reference--group-001.md#canonical-1112113123113022-2201302220003231-2300223303302212-1011113120222113-0032030131130203-1220003132100203-2233121121323013-3311323222322003) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](data-sources--alert_receiver--reference--group-001.md#canonical-2332232030131122-1131331211023110-1322303313312200-1022313320300323-1133032123031112-0211123001332212-2231002121100011-2213102012013021) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](data-sources--alert_receiver--reference--group-001.md#canonical-3002220313101000-3213102313331210-3221212333121321-2323100031230133-1003100300212122-1013330320330013-1311012321333101-0203020132012310) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](data-sources--alert_receiver--reference--group-001.md#canonical-1230111303133223-0332323220211132-0330012231001001-0202323011000310-3011001322332223-1232312203222220-2331313110322002-0031301010031201) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-1302212300130231-1000200320321130-1331031222130111-3021001313113020-2201312010002230-1232021321111310-0101111011011321-3221113233022302) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](data-sources--alert_receiver--reference--group-001.md#canonical-2121311033120101-0333012122210013-0020010103302211-3313233122031010-0333011201331302-3122032003332331-2102220123021230-1012103101333100) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](data-sources--alert_receiver--reference--group-001.md#canonical-1131103201012213-2333312212211000-3013033213002233-0102133111011230-0202220002032120-2233230032110111-1031302123121102-3310023232111010) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](data-sources--alert_receiver--reference--group-001.md#canonical-0221013303301002-3333310122021332-3022121121212121-3113222202133312-2010023322131221-1310232232002233-3322023321021222-2311232131320123) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](data-sources--alert_receiver--reference--group-001.md#canonical-2013010231132001-3211030001031232-1211032212301000-3023211210330003-3131122201320310-1111031013211123-2103011202201101-3323103333013032) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-1012112320111001-0323200132313201-3100103311300130-0303101013213210-1233130211033121-3330011300103212-2013233332300031-3112322321131101) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-3232202300213311-3303103223033101-2310122330220100-3213021311100033-3123110203221111-1131332021112301-0322113032123110-3003131320130002) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-3102310112213232-0123133231222011-2003003213122123-0111101322203301-3323231030332101-1002332211232021-2100122330022031-1331301312232101) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](data-sources--alert_receiver--reference--group-001.md#canonical-1020012301321201-2310322102003123-0130122301212111-3201203132111311-1112231201212310-3013012223332332-3313323212300121-0100031201102212) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](data-sources--alert_receiver--reference--group-001.md#canonical-2003222120033032-0031021213100230-1121003101311120-0132123223300010-3302011103311212-2300212021200302-3100320302132321-0013330320111020) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](data-sources--alert_receiver--reference--group-001.md#canonical-1103303231223220-3122223110030303-3312323231200122-2310301310003001-1100120220003230-2002010202023121-2213300001021231-2221113030022120) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](data-sources--alert_receiver--reference--group-001.md#canonical-1031230110220100-3223200220330221-3202231312233303-1033010310022202-0220310303201220-2332332321331331-2012011131120231-2231021132102012) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](data-sources--alert_receiver--reference--group-001.md#canonical-3231001000200131-2103203312120233-2323102000112001-3001232323122330-0203311130110310-2220100201200001-0330133300231322-3231213120233023) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-1303320001222232-3333211233233121-2202303013322000-0033003110021001-1313112031012312-2123121013323223-3301031122010300-0131133013000331) |
| `webhook.url` | [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-1330332210233223-0011020202322002-3020313120320320-1002102000023212-2013113020022323-2101003022331333-1023132123330121-2000012020220020) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2310133302133001-0112001310332232-0032232213210030-1301200320223212-3313230121121232-0321021210303212-3130232303201133-3000230323232222) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](data-sources--alert_receiver--reference--group-001.md#canonical-1200330032012003-2112302033130131-1001300002321303-1213023223201313-1220313111010032-0200212333321303-1123010133131221-0101301111120211) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](data-sources--alert_receiver--reference--group-001.md#canonical-3013203230122311-3113131123223110-1332323211231220-1211130333213330-1112300110322321-1003033320213301-0332122102010100-1300233123201313) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](data-sources--alert_receiver--reference--group-001.md#canonical-1113030110013222-1131031203303313-1202111113332010-0032131133123000-2011111121303110-1213230231322112-3303110020020333-3212002322130331) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3033003221333022-1001203331132121-3103133110030302-0121322002012022-1201012300322000-1332132331212121-3200223200213333-0333331330010332) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](data-sources--alert_receiver--reference--group-001.md#canonical-1232220112330221-2002132311021202-1011031223010010-2123131331220131-0010133311333331-0311302003031002-2313112133322310-2233113220303202) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](data-sources--alert_receiver--reference--group-001.md#canonical-2322320133303103-2012233112322011-1313030201222203-0133132022233001-3112312303113203-0310322133211032-3320213021012011-1323303021021132) |

<a id="canonical-0201220120220213-2102011031101202-1102230231331111-2220122101301312-3030220000230030-2233120122013213-3232011302301001-2102231023110000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `email` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- email

<a id="canonical-0310220320302133-0001131301031010-1023022303031230-2330231011323101-2122102210132322-2101013330202312-2221201110320200-1131121332022223"></a>

Type: `"single"`. Computed.

\[OneOf: email, opsgenie, pagerduty, Slack, sms, webhook\] Email Configuration.

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

- [email](data-sources--alert_receiver--reference--group-001.md#canonical-0310220320302133-0001131301031010-1023022303031230-2330231011323101-2122102210132322-2101013330202312-2221201110320200-1131121332022223)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-2101202022133322-1011003311212231-2331103013233003-3233221112202013-3011111111133130-0201312203221011-0301320113212001-1100011211202213)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2112113212110212-0001231112012103-1132210002201220-2122211111010302-3201022020221000-0231220220120221-3210331010011021-1020023102002231)
- [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3102222130123210-3312210300201322-1010302131033232-0301233120313102-3213211130200330-2303213000202201-2213301023213202-1332122033032312)
- [sms](data-sources--alert_receiver--reference--group-001.md#canonical-0202013131321302-2300223130301112-1302213220133131-3003322033201200-1103303102233312-2113230103122220-0233021112120030-2223320323002103)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0022101022221202-2100222333201221-2130213303213223-1022012233102212-2200210201130012-1012110220331022-0013311310233032-3112330303331123)

Select alternatives according to the provider validators above.

<a id="canonical-2110330112220023-1011320203100301-3101132103233133-3011333301022021-0101310220301013-2033200212020222-0100001001021232-1213201112121311"></a>

### Direct properties for `email`

<a id="canonical-3023230221321012-0331013210111010-2300003202013213-2123100223201323-0122321010022211-2213320201133321-0113211100122212-0201332333001123"></a>

#### `email.email` property

Type: `"string"`. Computed.

Email. Email ID of the user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-1211322113200310-1213220213023020-2232330133023322-0012212203103221-1123201022032210-1310021300010023-0123031121310201-0132303120000123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `opsgenie` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- opsgenie

<a id="canonical-2101202022133322-1011003311212231-2331103013233003-3233221112202013-3011111111133130-0201312203221011-0301320113212001-1100011211202213"></a>

Type: `"single"`. Computed.

OpsGenie configuration to send alert notifications.

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

<a id="canonical-2302321303133333-2320310032201121-3200223201231122-2331020330122230-1132213201111303-3102123310002223-3320113010310002-0101200222222210"></a>

### Direct properties for `opsgenie`

- [api_key](data-sources--alert_receiver--reference--group-001.md#canonical-1100333223230322-1011221303322030-1212310003333232-0110200322020230-1202200230202120-1000013220031131-2203131301322221-0232223313331322): complete subsection reference.

<a id="canonical-3033200133010202-2322203113111232-0222222301303113-3330022033320120-1312201103202233-0202213202333330-1300021231001221-2302002203210323"></a>

<a id="canonical-1333231123122302-2213123200132322-2112111130312103-3021100323103223-1320102320323233-3220120113113123-0031112033130223-1220213030022122"></a>

#### `opsgenie.url` property

Type: `"string"`. Computed.

API URL. URL to send API requests to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1100333223230322-1011221303322030-1212310003333232-0110200322020230-1202200230202120-1000013220031131-2203131301322221-0232223313331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `opsgenie.api_key` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-1211322113200310-1213220213023020-2232330133023322-0012212203103221-1123201022032210-1310021300010023-0123031121310201-0132303120000123)
- opsgenie.api_key

<a id="canonical-1130032232020101-1012103323103230-3310320101201131-1223200223211120-0302131002331111-2110213011130102-1201221020032322-2131000213323313"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-0000302021313021-3212320300201330-2230123030011122-3222202101030123-1203231030131033-3131120323223202-2002130033300122-2312221121300031"></a>

### Direct properties for `opsgenie.api_key`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2110110220202022-2030023222111301-1201112031323320-2101022232002203-3202133113300201-3001120332303032-1110031300310131-2202223212222330): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3313101032122200-1110112322323110-3210022103321033-1312021310303323-2113322000213123-0301321110030320-0231013312000332-0332333000310123): complete subsection reference.

<a id="canonical-2110110220202022-2030023222111301-1201112031323320-2101022232002203-3202133113300201-3001120332303032-1110031300310131-2202223212222330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `opsgenie.api_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-1211322113200310-1213220213023020-2232330133023322-0012212203103221-1123201022032210-1310021300010023-0123031121310201-0132303120000123)
- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-1100333223230322-1011221303322030-1212310003333232-0110200322020230-1202200230202120-1000013220031131-2203131301322221-0232223313331322)
- opsgenie.api_key.blindfold_secret_info

<a id="canonical-3300302123231322-3023320300223021-2332321012331301-0103313213011113-0102123222202110-0013011000310113-0002322221031222-2313002233320012"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3222010203003303-3020222211200210-0030010031021110-2302231133021323-3131300030200112-1210311131333211-2321103021302311-3221023301203230"></a>

### Direct properties for `opsgenie.api_key.blindfold_secret_info`

<a id="canonical-3321323223131331-1221223202300001-1013013133323333-2221103220001322-1310120223312231-1321301311203313-1120310330310131-0132033033200203"></a>

#### `opsgenie.api_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1132022131302210-0003120303300303-1020301020220211-2302221311113331-2121110103213313-1031021303101031-2313312313211330-1330113103312223"></a>

<a id="canonical-1001102213220001-1023032311313133-0311001022033001-3013202230112021-3233312220113233-0230123202120013-0312103122031232-3021233113210002"></a>

#### `opsgenie.api_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0333021111031133-1300300303020231-3120000102300311-2200011001201123-0032120002020311-0322002210011102-3230021302311020-2200232332233211"></a>

<a id="canonical-0132021111111033-3012132211112010-2122031100230000-1201100002001121-3210033332110020-0021203003100303-3011221323000331-3011203223201213"></a>

#### `opsgenie.api_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3313101032122200-1110112322323110-3210022103321033-1312021310303323-2113322000213123-0301321110030320-0231013312000332-0332333000310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `opsgenie.api_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [opsgenie](data-sources--alert_receiver--reference--group-001.md#canonical-1211322113200310-1213220213023020-2232330133023322-0012212203103221-1123201022032210-1310021300010023-0123031121310201-0132303120000123)
- [opsgenie.api_key](data-sources--alert_receiver--reference--group-001.md#canonical-1100333223230322-1011221303322030-1212310003333232-0110200322020230-1202200230202120-1000013220031131-2203131301322221-0232223313331322)
- opsgenie.api_key.clear_secret_info

<a id="canonical-2300211323303221-3332202033022320-3202301231201221-2312222121231013-3032020213000012-0111113223101222-2222122301212001-1122131122203131"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-0312201233111132-0123201311100113-0210312330221132-0320132021021132-2230120233102121-2213312131312031-0013011013021000-1203123120312133"></a>

### Direct properties for `opsgenie.api_key.clear_secret_info`

<a id="canonical-1311021301210303-1230203312211203-2220131221122113-3001001312032102-1232331322300022-2201323111322221-2022212130000113-0210003100301333"></a>

#### `opsgenie.api_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1331001133111233-1101322213021223-1300333301323030-1033312012121212-1022302202223112-2300312303222203-0101010211330012-3000210033013310"></a>

<a id="canonical-1202100120033202-1113011302112113-0031132200103133-2131013020211013-2133200031131010-0220132311112013-0122032213322223-3200130121113011"></a>

#### `opsgenie.api_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2310320030111211-3032330222323111-2003331131121231-2021223211110101-3221203112203233-3022010211323123-2003032312303302-0111312030013211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pagerduty` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- pagerduty

<a id="canonical-2112113212110212-0001231112012103-1132210002201220-2122211111010302-3201022020221000-0231220220120221-3210331010011021-1020023102002231"></a>

Type: `"single"`. Computed.

PagerDuty configuration to send alert notifications.

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

<a id="canonical-1211030011200301-3231201333230200-3002011222110032-1031322233213133-3102333320111130-0203211322320130-3002120323313322-1012002100301223"></a>

### Direct properties for `pagerduty`

- [routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-2012000210132323-1002231011103021-3222330123122310-0202132020221333-2003330333212213-2032020202301331-0220311201101323-1120103223322103): complete subsection reference.

<a id="canonical-2203330312013112-3310211323103220-3323113031213313-1201123330202133-0013221330310130-0123310333011230-1032222302122111-1213103200013203"></a>

<a id="canonical-1113323120230333-0012313322003321-3123002331213211-0311111111010000-0233322202313201-1223213302102203-1221001011031301-1001000222202322"></a>

#### `pagerduty.url` property

Type: `"string"`. Computed.

Pager Duty URL. URL to send API requests to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2012000210132323-1002231011103021-3222330123122310-0202132020221333-2003330333212213-2032020202301331-0220311201101323-1120103223322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pagerduty.routing_key` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2310320030111211-3032330222323111-2003331131121231-2021223211110101-3221203112203233-3022010211323123-2003032312303302-0111312030013211)
- pagerduty.routing_key

<a id="canonical-2031321110220332-1111032111131220-1101033200323233-1222130231301302-3122203111112000-3202131023021020-3120122101201130-1313313331101233"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1021212103211302-2011113031103203-2211332213031112-2100111331211312-3231012331003200-3203323213302121-0102001300200320-1303112121213102"></a>

### Direct properties for `pagerduty.routing_key`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2002231303320201-0110113232212030-2321102232213210-1210212202321012-2033031311033110-0313230111002023-1202031331030200-0011122012213300): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-1301122001012323-0010103120233331-2202312122023203-3312100220102302-3210031103230200-0020003211223320-0301220123301033-2130210123123230): complete subsection reference.

<a id="canonical-2002231303320201-0110113232212030-2321102232213210-1210212202321012-2033031311033110-0313230111002023-1202031331030200-0011122012213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pagerduty.routing_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2310320030111211-3032330222323111-2003331131121231-2021223211110101-3221203112203233-3022010211323123-2003032312303302-0111312030013211)
- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-2012000210132323-1002231011103021-3222330123122310-0202132020221333-2003330333212213-2032020202301331-0220311201101323-1120103223322103)
- pagerduty.routing_key.blindfold_secret_info

<a id="canonical-2031132113330330-1031111112333310-2310323313032103-0022022312000122-3230330231321322-1122322320020300-3302310231200201-1231331331210010"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-0130102310023132-1133000003232011-0122302310111210-2232200003320130-3020002312011122-1103312222301020-1310330321100021-1330211331120302"></a>

### Direct properties for `pagerduty.routing_key.blindfold_secret_info`

<a id="canonical-2130130012022123-0322123002100222-2013032112010233-3233002120103312-0311311212201233-3011321002132032-3103002111101230-0113130111102300"></a>

#### `pagerduty.routing_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3320213301002133-3011203010131130-2201313231202020-3320211020213110-3202111203232330-0333100323220232-1321020020303112-0233002331122130"></a>

<a id="canonical-2012120311233210-3000121030231120-2022210002133233-3121203031323230-0013233120000113-2032333110001312-1101000223020021-1323222231320003"></a>

#### `pagerduty.routing_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120201201112323-1332201331332001-3211120210321102-2332213003221021-2101311201021231-3011110310022122-2002132221202322-2311132331000132"></a>

<a id="canonical-0200213030213103-1101113220123232-1020220313202211-2333111231100032-3213302331332301-2231012111311212-2222230320332300-0210003300210200"></a>

#### `pagerduty.routing_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1301122001012323-0010103120233331-2202312122023203-3312100220102302-3210031103230200-0020003211223320-0301220123301033-2130210123123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pagerduty.routing_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [pagerduty](data-sources--alert_receiver--reference--group-001.md#canonical-2310320030111211-3032330222323111-2003331131121231-2021223211110101-3221203112203233-3022010211323123-2003032312303302-0111312030013211)
- [pagerduty.routing_key](data-sources--alert_receiver--reference--group-001.md#canonical-2012000210132323-1002231011103021-3222330123122310-0202132020221333-2003330333212213-2032020202301331-0220311201101323-1120103223322103)
- pagerduty.routing_key.clear_secret_info

<a id="canonical-2312010200333332-3300213230121231-0022002001222213-0002201310112123-2220021212130020-0013333313321020-2232021022120002-2213311022302033"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1000033330330223-3230323330031133-0020111112302212-2032201003002310-0100101031002111-0220333311002123-3223033313131332-2313000322301123"></a>

### Direct properties for `pagerduty.routing_key.clear_secret_info`

<a id="canonical-1210331333211320-3013302023022103-0231032332002013-0310321232222311-2222220303330232-2313221032133230-3020012011001032-1112211322213232"></a>

#### `pagerduty.routing_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3201020320013223-1330232011300031-2332021032321020-0212303030320213-2021132310313023-0302212232200002-2110013020210001-3102010200301301"></a>

<a id="canonical-0310001322321001-0123112100123123-0200303302001002-0223122300203202-1323330200231300-0023333210023203-3121130203111200-1321123103202302"></a>

#### `pagerduty.routing_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slack` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- Slack

<a id="canonical-3102222130123210-3312210300201322-1010302131033232-0301233120313102-3213211130200330-2303213000202201-2213301023213202-1332122033032312"></a>

Type: `"single"`. Computed.

Slack configuration to send alert notifications.

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

<a id="canonical-2010222020012202-2013010313211021-0133311333321122-0023111321300033-3232123030103000-2313110331333132-1231032220233223-2023131011311212"></a>

### Direct properties for `slack`

<a id="canonical-2233013223212231-3230301102032222-1100123200013030-1111132013010101-2213100020331310-2032230330032102-3310331003311022-2320322122211032"></a>

#### `slack.channel` property

Type: `"string"`. Computed.

Channel or user to send notifications to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  }
}
```

- [URL](data-sources--alert_receiver--reference--group-001.md#canonical-1010221222222122-1232032003021211-0132111211200300-1230203023100021-3333313011300332-1123102031101110-1303333212103120-3133200331231032): complete subsection reference.

<a id="canonical-1010221222222122-1232032003021211-0132111211200300-1230203023100021-3333313011300332-1123102031101110-1303333212103120-3133200331231032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slack.url` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323)
- Slack.URL

<a id="canonical-2102301112232111-2323123321232310-2321203110010021-3323333130320021-3010200020102320-0311220021303210-2123033020202012-0212022200012102"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1321120212231103-0001220002033112-0322132300202232-2322321222020212-2121020311000230-0233121311133320-0310221010223131-2300003331232312"></a>

### Direct properties for `slack.url`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3303111233202201-0311012222123131-2012031103033213-0230130010101100-0103002112132210-2312123023001102-3103130220030320-1333222232212232): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3132222201223311-0121122210333121-1232333233310101-3321103003323113-0210321133201130-1223101002333000-2320032111010300-3310202323321020): complete subsection reference.

<a id="canonical-3303111233202201-0311012222123131-2012031103033213-0230130010101100-0103002112132210-2312123023001102-3103130220030320-1333222232212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slack.url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323)
- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-1010221222222122-1232032003021211-0132111211200300-1230203023100021-3333313011300332-1123102031101110-1303333212103120-3133200331231032)
- Slack.URL.blindfold_secret_info

<a id="canonical-1123321313130201-2101203110331012-1011103013331323-3232232331311123-3033103233123200-2103331122022303-3300213210110332-2033121332131011"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1331023133130233-3330010030031321-2123113231123302-3012213030033033-1210001301013013-2303130301112210-2220021203013130-0232012303203103"></a>

### Direct properties for `slack.url.blindfold_secret_info`

<a id="canonical-2133110003020230-0212310023001331-3032111331102201-2123320332320031-0332203033332021-1312303112221102-0103112202323112-1201213002211311"></a>

#### `slack.url.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1321100313113321-2230021102203030-0101212130111011-1322310322030002-3230012000213330-3231032233103220-0011311320132013-3102222231111103"></a>

<a id="canonical-3120131320300010-3212333201312222-3002030130213310-0101020111310231-3003222220300103-1310200313132202-0111221313013003-3313210233230112"></a>

#### `slack.url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1013122100332110-0221000222223211-2112320230222000-2131121200321223-2331132023012303-2101031032223201-0120131033203213-3000011111313213"></a>

<a id="canonical-0022302203220031-3301020203133301-3010133120323320-3200231110213332-0131023023123112-1200202310312010-0212201332031032-2101000003031121"></a>

#### `slack.url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3132222201223311-0121122210333121-1232333233310101-3321103003323113-0210321133201130-1223101002333000-2320032111010300-3310202323321020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slack.url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [Slack](data-sources--alert_receiver--reference--group-001.md#canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323)
- [slack.url](data-sources--alert_receiver--reference--group-001.md#canonical-1010221222222122-1232032003021211-0132111211200300-1230203023100021-3333313011300332-1123102031101110-1303333212103120-3133200331231032)
- Slack.URL.clear_secret_info

<a id="canonical-1302032020021000-1310130220133030-3211301013020132-1100223223033101-1030220001110033-0120122113003102-1230203202003121-2301131200211000"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2321322011103013-2100223012101311-0310301332113102-0030001012211013-1232222021332100-0010133101222201-2032002221312110-1320021031333121"></a>

### Direct properties for `slack.url.clear_secret_info`

<a id="canonical-1332110233003311-3110110110112202-3100021330133223-1211210312301011-3021311203230313-2221130001012022-0021000002201033-2103321302120103"></a>

#### `slack.url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3033031201231130-0031312132003021-3300220203312132-1313320230020323-3003002010211131-0003033310132111-0131310312331310-1330321320022132"></a>

<a id="canonical-1203231022311211-3231323332320110-2122021011101002-0110130133112223-0220213121222132-3332312003112213-3211002000222311-0211011020312222"></a>

#### `slack.url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1102020323302231-3330031312223101-2332200311130100-0221003321130033-0310033330311110-2222220310130220-0311300330010121-3122100310111130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sms` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- sms

<a id="canonical-0202013131321302-2300223130301112-1302213220133131-3003322033201200-1103303102233312-2113230103122220-0233021112120030-2223320323002103"></a>

Type: `"single"`. Computed.

SMS Configuration.

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

<a id="canonical-3002210202310112-3013321313233102-0322023201200223-2210203213321220-0213111111201320-1300201202011331-2200133310312130-2021022100011220"></a>

### Direct properties for `sms`

<a id="canonical-2311302303011123-0122312302031032-1331000122010120-3121101110211031-0010230220031000-3031021302111332-1322310333033323-0321122003230213"></a>

#### `sms.contact_number` property

Type: `"string"`. Computed.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

<a id="canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- webhook

<a id="canonical-0022101022221202-2100222333201221-2130213303213223-1022012233102212-2200210201130012-1012110220331022-0013311310233032-3112330303331123"></a>

Type: `"single"`. Computed.

Webhook configuration to send alert notifications.

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

<a id="canonical-2332133303032020-0101123022300300-1331220310212000-2021302110303330-0202332000233123-1013013033001122-0032021101232110-2132121310121111"></a>

### Direct properties for `webhook`

- [http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012): complete subsection reference.

- [URL](data-sources--alert_receiver--reference--group-001.md#canonical-3320213310121301-2102313133202001-1222223320030002-0210110101002111-2033020200031033-0021031212201202-1201202021022300-3332013032320210): complete subsection reference.

<a id="canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- webhook.http_config

<a id="canonical-3232100101023232-2221112313023231-2210012122022332-3330310321021210-0221102210132213-2020302031000133-1312223213023020-1002201300211032"></a>

Type: `"single"`. Computed.

HTTP Configuration. Configuration for HTTP endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-3030203132303311-2310203100323321-0312213321201222-0122020022312022-0010213121203210-3303013211130222-0313310133320232-0133132313332032"></a>

### Direct properties for `webhook.http_config`

- [auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-3212332102323121-2310200122313131-0321011120122000-0323031312333200-0110033322131121-2020223023120022-0320002133020011-1331100101132113): complete subsection reference.

- [basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-2122123202033231-3121123322333211-2313033202231112-3112301330133030-0103133110302103-1002212221302333-0322001010122031-2013020210012000): complete subsection reference.

- [client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-2311000221310311-1210201200003322-2321232110300330-3023130122002031-3131003023220200-1203200302202301-1223013222002320-0211023111213222): complete subsection reference.

<a id="canonical-1112113123113022-2201302220003231-2300223303302212-1011113120222113-0032030131130203-1220003132100203-2233121121323013-3311323222322003"></a>

<a id="canonical-1112211233310221-3132322000323110-1021000000100113-1203200123012312-0232100231322010-2132313210022230-3130010133312002-0220102301130213"></a>

#### `webhook.http_config.enable_http2` property

Type: `"bool"`. Computed.

Enable HTTP2. Configure to use HTTP2 protocol.

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

<a id="canonical-2332232030131122-1131331211023110-1322303313312200-1022313320300323-1133032123031112-0211123001332212-2231002121100011-2213102012013021"></a>

<a id="canonical-0323112220200111-2211112300010212-0222031200232023-3032222322133311-1012211202102231-0333300232013133-3133032113333310-2111210203231332"></a>

#### `webhook.http_config.follow_redirects` property

Type: `"bool"`. Computed.

Configure whether HTTP requests follow HTTP 3xx redirects.

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

- [no_authorization](data-sources--alert_receiver--reference--group-001.md#canonical-2301030333321310-1212222023211233-3121330011012100-1032111111121332-3220121300231203-2203323220130213-3112302233213120-3202010233130312): complete subsection reference.

- [no_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3332231222110321-3020221133211032-0213331012313022-2010232320233221-0313201211310220-1003103320200313-1210032223122331-3300020121000113): complete subsection reference.

- [use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132): complete subsection reference.

<a id="canonical-3212332102323121-2310200122313131-0321011120122000-0323031312333200-0110033322131121-2020223023120022-0320002133020011-1331100101132113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.auth_token` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.auth_token

<a id="canonical-0133130101020210-1233130111101013-1010330230203003-0020123101203212-3212132322113230-2321211303231201-0123320302131220-1102022010102303"></a>

Type: `"single"`. Computed.

Access Token. Authentication Token for access.

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

<a id="canonical-0200333331202113-0210331031202301-2210332010010212-2012203320021300-3232123030322110-0100003002201203-0302002310011010-2100212120211300"></a>

### Direct properties for `webhook.http_config.auth_token`

- [token](data-sources--alert_receiver--reference--group-001.md#canonical-0203333133310200-1310220321311310-1001100300021021-1120110230200302-0302120302300331-2231221111311310-3123231223231011-1222130221231200): complete subsection reference.

<a id="canonical-0203333133310200-1310220321311310-1001100300021021-1120110230200302-0302120302300331-2231221111311310-3123231223231011-1222130221231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.auth_token.token` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-3212332102323121-2310200122313131-0321011120122000-0323031312333200-0110033322131121-2020223023120022-0320002133020011-1331100101132113)
- webhook.http_config.auth_token.token

<a id="canonical-2010000121233200-1221020002023011-1312302112010212-0120301203302212-3011222320222203-1131333111001003-2101323222000123-0103011312102210"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1030231222100202-3312323331330020-2232202232111202-2030000331110100-2112003013333333-1000030300113103-3013103311102332-0002201300330312"></a>

### Direct properties for `webhook.http_config.auth_token.token`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3033301320212301-1113100232231230-3010002230123220-0120222020322322-2233333200102320-3332010020113010-0323203333022212-1103322201123222): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-0210333320202002-1131331311103001-0321032100013033-0031201313121113-0323132233230130-0013302033200111-1213230301022031-1301032000033213): complete subsection reference.

<a id="canonical-3033301320212301-1113100232231230-3010002230123220-0120222020322322-2233333200102320-3332010020113010-0323203333022212-1103322201123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.auth_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-3212332102323121-2310200122313131-0321011120122000-0323031312333200-0110033322131121-2020223023120022-0320002133020011-1331100101132113)
- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-0203333133310200-1310220321311310-1001100300021021-1120110230200302-0302120302300331-2231221111311310-3123231223231011-1222130221231200)
- webhook.http_config.auth_token.token.blindfold_secret_info

<a id="canonical-0131021031000112-2102131223201102-0103122230011010-3210311233300332-2312302013110202-0303302303122023-0011210030112310-2112313020311221"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-2112230300021002-0100032021121232-2023003332110010-1011030221232132-0203202223002312-3202030233022330-2020313003123213-2211113300210023"></a>

### Direct properties for `webhook.http_config.auth_token.token.blindfold_secret_info`

<a id="canonical-2103010331111022-0103020301023130-0001220322020021-3322221111022320-2233033212013311-3023032000022002-0023312211230102-3020231322313211"></a>

#### `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1201113100232331-3303012021000211-3313011313200020-2100310210222213-1112130121233301-2331022130132333-2332023010312133-2122030111023310"></a>

<a id="canonical-1331232300310020-1200121230131122-0200233330200133-1013002221313302-2301100312122001-1122233030120023-0003032203311200-2331000121120110"></a>

#### `webhook.http_config.auth_token.token.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3132320333310211-3012220110230213-2221300200330122-0311021333302112-2130201022301030-0023123122132223-1311130002210200-0230033312331302"></a>

<a id="canonical-1123211102223003-0333012121201110-2010311232111202-0131231013213032-1333301203130121-2321331003012103-2303300010300310-2113032031203213"></a>

#### `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0210333320202002-1131331311103001-0321032100013033-0031201313121113-0323132233230130-0013302033200111-1213230301022031-1301032000033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.auth_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.auth_token](data-sources--alert_receiver--reference--group-001.md#canonical-3212332102323121-2310200122313131-0321011120122000-0323031312333200-0110033322131121-2020223023120022-0320002133020011-1331100101132113)
- [webhook.http_config.auth_token.token](data-sources--alert_receiver--reference--group-001.md#canonical-0203333133310200-1310220321311310-1001100300021021-1120110230200302-0302120302300331-2231221111311310-3123231223231011-1222130221231200)
- webhook.http_config.auth_token.token.clear_secret_info

<a id="canonical-3210220221122212-3000312032222023-0113202211302003-3103333101223220-2001310202320023-3302023120012122-0111010001000131-3223113213303233"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2203100201101021-0022322330223132-1121302113003310-3323003022012111-3023111211313232-0111223212300013-0101302220201300-1113303300210021"></a>

### Direct properties for `webhook.http_config.auth_token.token.clear_secret_info`

<a id="canonical-2231100203132000-3033120001320020-0031303321100300-1013011121023120-0122210200131301-3313331333332133-1212220201233330-2100021013001132"></a>

#### `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1303312313331233-3132301021220102-0032131321333102-2112233013022022-1001321303302321-0111331022100121-2203121331131110-3011113102131130"></a>

<a id="canonical-0331203021121130-2313320002033301-2130132223302210-2213033002132222-0100311233221331-1023030120212301-2110132331322233-0002222332003232"></a>

#### `webhook.http_config.auth_token.token.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2122123202033231-3121123322333211-2313033202231112-3112301330133030-0103133110302103-1002212221302333-0322001010122031-2013020210012000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.basic_auth` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.basic_auth

<a id="canonical-1230301233221213-0112313201312310-1113102210231301-3301302300203112-0023123101111013-3132202212023021-2223010131221133-1130321122030212"></a>

Type: `"single"`. Computed.

Authorization parameters to access HTPP alert Receiver Endpoint.

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

<a id="canonical-2121102312002333-0121330021213010-3221310031003132-0130320032303213-1111110111331132-1031111013301323-1220313021203303-2222111030011113"></a>

### Direct properties for `webhook.http_config.basic_auth`

- [password](data-sources--alert_receiver--reference--group-001.md#canonical-3011232020310330-3031112210033310-3032132010032302-1113001212233200-1311313132301132-3223231030322231-0003012022031330-1233321130101213): complete subsection reference.

<a id="canonical-0010123102232110-3012222320333000-0133221313212201-2312030021332103-3312133212323023-2011002222212220-1032003112203012-1213221122330311"></a>

<a id="canonical-3333330322220322-2213130311021321-2233230101012332-1210230300203323-3022100322233131-1120022030232212-2132211010021111-3210200031132320"></a>

#### `webhook.http_config.basic_auth.user_name` property

Type: `"string"`. Computed.

username. HTTP Basic Auth username.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3011232020310330-3031112210033310-3032132010032302-1113001212233200-1311313132301132-3223231030322231-0003012022031330-1233321130101213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.basic_auth.password` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-2122123202033231-3121123322333211-2313033202231112-3112301330133030-0103133110302103-1002212221302333-0322001010122031-2013020210012000)
- webhook.http_config.basic_auth.password

<a id="canonical-0331011202133100-2013201120131010-3103101110323100-1001023232122030-2031311202102023-0113230010211323-1200213130200222-3232010032131121"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-3102021210130202-2111212320211303-3201121223100203-1210121010103210-0003332212022203-0102203303112302-3010112311013323-0323323132021301"></a>

### Direct properties for `webhook.http_config.basic_auth.password`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-3320100313203311-2333110031110030-0311321112312232-1013323011022212-2110123032013300-1012101111023000-2310133232223003-3030031021031003): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-2032001030033110-2010003211313011-3202121202100213-2010023300203212-3000003210130211-2013322000200103-0020113220331020-3101333203133202): complete subsection reference.

<a id="canonical-3320100313203311-2333110031110030-0311321112312232-1013323011022212-2110123032013300-1012101111023000-2310133232223003-3030031021031003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.basic_auth.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-2122123202033231-3121123322333211-2313033202231112-3112301330133030-0103133110302103-1002212221302333-0322001010122031-2013020210012000)
- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-3011232020310330-3031112210033310-3032132010032302-1113001212233200-1311313132301132-3223231030322231-0003012022031330-1233321130101213)
- webhook.http_config.basic_auth.password.blindfold_secret_info

<a id="canonical-0333122222310333-3333223101133211-1221000133222010-3120102332133021-0311303310201130-1022323203220000-3121010110311100-3113102223200003"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-0130300103312332-0303032101033321-1220021310222311-2113100322130331-0311202131220333-0213002123000230-3333002102120333-3000120001111002"></a>

### Direct properties for `webhook.http_config.basic_auth.password.blindfold_secret_info`

<a id="canonical-1230110022211030-2311132202221330-2133203010310330-0030033202303330-1310321330121032-3101303222330223-0200130330023302-3030032300122232"></a>

#### `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2121321321112333-1201223202231001-1212200120003131-1002001003030121-3121231212232133-3003320101322023-1111023201300131-1003220133321123"></a>

<a id="canonical-1203010130311033-0000201331003112-0210122320130201-1233020330032222-2102311321210002-0032223031131021-2223313332320023-1030211233313201"></a>

#### `webhook.http_config.basic_auth.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0011020200021222-3013211322212301-2013112213201121-0212211312202230-1131002213121003-0011030221033222-0003323331030013-3110030103212313"></a>

<a id="canonical-0123013222131123-0111333023013011-1031302112010111-3033130302302000-3110132130231302-3330000222032230-0110003021123131-1330330321113223"></a>

#### `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2032001030033110-2010003211313011-3202121202100213-2010023300203212-3000003210130211-2013322000200103-0020113220331020-3101333203133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.basic_auth.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--reference--group-001.md#canonical-2122123202033231-3121123322333211-2313033202231112-3112301330133030-0103133110302103-1002212221302333-0322001010122031-2013020210012000)
- [webhook.http_config.basic_auth.password](data-sources--alert_receiver--reference--group-001.md#canonical-3011232020310330-3031112210033310-3032132010032302-1113001212233200-1311313132301132-3223231030322231-0003012022031330-1233321130101213)
- webhook.http_config.basic_auth.password.clear_secret_info

<a id="canonical-2131223031100210-0020330202231010-0332103133232201-2221300001321112-0111302121232330-1100131010132223-2200301203102221-1211102200003320"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1112320310333110-1112130113023220-1303302211133132-3033131131221301-2133003103113121-1013233130330012-3303100100210011-1001332312203302"></a>

### Direct properties for `webhook.http_config.basic_auth.password.clear_secret_info`

<a id="canonical-2123021030011013-0323330121322231-1122001330213000-1132233121220120-0112311310331011-1213110132232110-2101323210033000-1120112131102331"></a>

#### `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1020300120302312-2313312322333103-2312132310113101-0312220011302220-3300030320111331-3312301113113112-3022333031023013-3120101201311333"></a>

<a id="canonical-2122300311222110-2122300003032010-0213212231123333-2220130222121323-3313013213120102-0231022231302332-2313120323031332-0132222200322323"></a>

#### `webhook.http_config.basic_auth.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2311000221310311-1210201200003322-2321232110300330-3023130122002031-3131003023220200-1203200302202301-1223013222002320-0211023111213222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.client_cert_obj` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.client_cert_obj

<a id="canonical-1332233003013210-1101013202331111-1331213222131332-3322020130130023-1202000121311101-2021000232012302-1002212233232222-1123130313222030"></a>

Type: `"single"`. Computed.

Client Certificate Object. Configuration for client certificate.

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

<a id="canonical-3033102321201311-1111121001032312-1232033031121201-2303320132200001-1103101123011223-2131332330231103-0311333321113212-2013033333103010"></a>

### Direct properties for `webhook.http_config.client_cert_obj`

- [use_tls_obj](data-sources--alert_receiver--reference--group-001.md#canonical-3311330210202103-0223333111233212-2020120302010232-2321110312012021-2003123330031221-1313000200301110-0302102122013102-0000121020331300): complete subsection reference.

<a id="canonical-3311330210202103-0223333111233212-2020120302010232-2321110312012021-2003123330031221-1313000200301110-0302102122013102-0000121020331300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.client_cert_obj.use_tls_obj` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.client_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-2311000221310311-1210201200003322-2321232110300330-3023130122002031-3131003023220200-1203200302202301-1223013222002320-0211023111213222)
- webhook.http_config.client_cert_obj.use_tls_obj

<a id="canonical-3320200102121123-0131002212230120-1220330000203323-0102001022210300-2300332331023322-2033313310323302-2201222123120131-3320310013103102"></a>

Type: `"list"`. Computed.

Certificate Object. Reference to client certificate object.

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

<a id="canonical-1222031312311221-2003223001212310-0211313021000132-2111211322222211-3120010000332130-1103101233333322-2000023133323010-0331222103332222"></a>

### Direct properties for `webhook.http_config.client_cert_obj.use_tls_obj`

<a id="canonical-3233123031321213-3203121321301302-1112002201110133-0233012201220202-1003320010000030-1013011123113033-3122030311121221-3033033023301131"></a>

#### `webhook.http_config.client_cert_obj.use_tls_obj.kind` property

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

<a id="canonical-1010111201222233-0302322231020311-1321211323333333-0011132310233032-1000030102033022-0111101323200022-1233230212301020-0302310020233301"></a>

<a id="canonical-2133122121102232-0330012220133112-1020202122233300-1223302112300123-2231103200013233-0301131220302003-3013212331321130-2032110133120303"></a>

#### `webhook.http_config.client_cert_obj.use_tls_obj.name` property

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

<a id="canonical-2102132130301110-2111222310011011-0321011311330332-3303323300222202-3101221210100203-1101230221213222-1200020111130100-3302230020013233"></a>

<a id="canonical-2121322212100201-1312213303010220-0103001330033210-3312130000302302-0132123111113201-2331131100310011-2220211321320122-2003312122211303"></a>

#### `webhook.http_config.client_cert_obj.use_tls_obj.namespace` property

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

<a id="canonical-1332113001230301-3312132303101120-3010113320022222-1132012201031200-3201111110101330-0111021323030222-3303213222033323-3201221011320322"></a>

<a id="canonical-2101231301000102-2112302131202331-3331302202020023-2210003230311010-2312120301203320-1231212003100313-0121031230111012-2032001221202032"></a>

#### `webhook.http_config.client_cert_obj.use_tls_obj.tenant` property

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

<a id="canonical-3021131022303032-3332331323133200-1101302032131201-1122003122230200-3133210332013110-1321031100303021-1133031233230311-0112312113201111"></a>

<a id="canonical-2231013331310312-2102030012202321-2322030021012010-3121200231021301-2200200310310310-3021232000102102-0002122011101123-1332323123332112"></a>

#### `webhook.http_config.client_cert_obj.use_tls_obj.uid` property

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

<a id="canonical-2301030333321310-1212222023211233-3121330011012100-1032111111121332-3220121300231203-2203323220130213-3112302233213120-3202010233130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.no_authorization` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.no_authorization

<a id="canonical-3002220313101000-3213102313331210-3221212333121321-2323100031230133-1003100300212122-1013330320330013-1311012321333101-0203020132012310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authorization.

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

<a id="canonical-3332231222110321-3020221133211032-0213331012313022-2010232320233221-0313201211310220-1003103320200313-1210032223122331-3300020121000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.no_tls` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.no_tls

<a id="canonical-1230111303133223-0332323220211132-0330012231001001-0202323011000310-3011001322332223-1232312203222220-2331313110322002-0031301010031201"></a>

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

<a id="canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- webhook.http_config.use_tls

<a id="canonical-1302212300130231-1000200320321130-1331031222130111-3021001313113020-2201312010002230-1232021321111310-0101111011011321-3221113233022302"></a>

Type: `"single"`. Computed.

Configures the token request's TLS settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

<a id="canonical-2222012230322320-0220120333011120-0032111021213313-0121233330200222-0001133031110302-0123323031212121-3101213231022001-1222221201010202"></a>

### Direct properties for `webhook.http_config.use_tls`

- [disable_sni](data-sources--alert_receiver--reference--group-001.md#canonical-2212010002301233-3212332222032310-2230120222111233-2230130233131320-3330312030301221-3010301223313301-2211030210120301-3212201222321001): complete subsection reference.

<a id="canonical-1131103201012213-2333312212211000-3013033213002233-0102133111011230-0202220002032120-2233230032110111-1031302123121102-3310023232111010"></a>

<a id="canonical-1212011102000311-0130001202131011-2022232320121011-3232330201233032-3020002001131110-3031023311312223-1112230131031311-0300202111020202"></a>

#### `webhook.http_config.use_tls.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0221013303301002-3333310122021332-3022121121212121-3113222202133312-2010023322131221-1310232232002233-3322023321021222-2311232131320123"></a>

<a id="canonical-0213322303002033-0111123310232132-0123120022132003-3121100232201033-0100212332220301-2213131031012030-2312331123132223-1003311310021111"></a>

#### `webhook.http_config.use_tls.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2013010231132001-3211030001031232-1211032212301000-3023211210330003-3131122201320310-1111031013211123-2103011202201101-3323103333013032"></a>

<a id="canonical-0313303223321112-0230303220211321-1330102011233030-2203002222213211-2122222231103301-0233222120323123-0311001213232221-3001133320130112"></a>

#### `webhook.http_config.use_tls.sni` property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-0001031210332313-3310013011030213-1220232223110312-0030002100112211-2132023301203221-2030213200231232-2232011311010032-3131031103222113): complete subsection reference.

- [volterra_trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-2202022223012232-1130032013220201-2112021321230103-3201200133303301-1113332303020231-1002132332112122-0022210131330211-1121011132300300): complete subsection reference.

<a id="canonical-2212010002301233-3212332222032310-2230120222111233-2230130233131320-3330312030301221-3010301223313301-2211030210120301-3212201222321001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132)
- webhook.http_config.use_tls.disable_sni

<a id="canonical-2121311033120101-0333012122210013-0020010103302211-3313233122031010-0333011201331302-3122032003332331-2102220123021230-1012103101333100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-0001031210332313-3310013011030213-1220232223110312-0030002100112211-2132023301203221-2030213200231232-2232011311010032-3131031103222113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132)
- webhook.http_config.use_tls.use_server_verification

<a id="canonical-1012112320111001-0323200132313201-3100103311300130-0303101013213210-1233130211033121-3330011300103212-2013233332300031-3112322321131101"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

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

<a id="canonical-0231021121031230-1222030212302112-3313031100232031-1131210013011130-0312222132103320-1330102313121321-0232320132120313-2232210213101112"></a>

### Direct properties for `webhook.http_config.use_tls.use_server_verification`

- [ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-3213001032311223-1032133001311020-1320031222113033-3213112031013112-0210122320133311-3321123033023310-1230310300323333-0033010033103333): complete subsection reference.

<a id="canonical-3213001032311223-1032133001311020-1320031222113033-3213112031013112-0210122320133311-3321123033023310-1230310300323333-0033010033103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-0001031210332313-3310013011030213-1220232223110312-0030002100112211-2132023301203221-2030213200231232-2232011311010032-3131031103222113)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="canonical-3232202300213311-3303103223033101-2310122330220100-3213021311100033-3123110203221111-1131332021112301-0322113032123110-3003131320130002"></a>

Type: `"single"`. Computed.

Configuration parameter for ca cert obj.

Additional upstream details:

Configuration for CA certificate.

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

<a id="canonical-3010302332302200-3300100020132212-1203122111003032-1333032300032120-3302223210321330-3330101033130020-1230233203003310-1301303133132210"></a>

### Direct properties for `webhook.http_config.use_tls.use_server_verification.ca_cert_obj`

- [trusted_ca](data-sources--alert_receiver--reference--group-001.md#canonical-1320011322010103-3322012322231233-3322011132231330-0010202221213301-0320301300202011-3011033033311222-3000202001323023-3202111233302132): complete subsection reference.

<a id="canonical-1320011322010103-3322012322231233-3322011132231330-0010202221213301-0320301300202011-3011033033311222-3000202001323023-3202111233302132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--reference--group-001.md#canonical-0001031210332313-3310013011030213-1220232223110312-0030002100112211-2132023301203221-2030213200231232-2232011311010032-3131031103222113)
- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--reference--group-001.md#canonical-3213001032311223-1032133001311020-1320031222113033-3213112031013112-0210122320133311-3321123033023310-1230310300323333-0033010033103333)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca

<a id="canonical-3102310112213232-0123133231222011-2003003213122123-0111101322203301-3323231030332101-1002332211232021-2100122330022031-1331301312232101"></a>

Type: `"list"`. Computed.

Certificate Object. Reference to client certificate object.

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

<a id="canonical-0110111001223020-3122101313213022-2333201001300032-0001132113012000-3122113010012212-1302320113131122-1022320323123232-2300110010202122"></a>

### Direct properties for `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca`

<a id="canonical-1020012301321201-2310322102003123-0130122301212111-3201203132111311-1112231201212310-3013012223332332-3313323212300121-0100031201102212"></a>

#### `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` property

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

<a id="canonical-2003222120033032-0031021213100230-1121003101311120-0132123223300010-3302011103311212-2300212021200302-3100320302132321-0013330320111020"></a>

<a id="canonical-1020103012003110-3211213012101202-2113212133113301-0221300333203202-3002120233300332-0133100203312002-3122211233101100-0201332023132302"></a>

#### `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` property

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

<a id="canonical-1103303231223220-3122223110030303-3312323231200122-2310301310003001-1100120220003230-2002010202023121-2213300001021231-2221113030022120"></a>

<a id="canonical-3333313112223223-0233311220101133-3303231303023322-3100333002023211-0322203011130213-0031003220200331-1033122200110200-3101302233033330"></a>

#### `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` property

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

<a id="canonical-1031230110220100-3223200220330221-3202231312233303-1033010310022202-0220310303201220-2332332321331331-2012011131120231-2231021132102012"></a>

<a id="canonical-1321123012212322-1230212210330221-0000302311311332-3100130203230322-2000223320202232-1010110213233232-2000200011003200-1210223212101313"></a>

#### `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` property

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

<a id="canonical-3231001000200131-2103203312120233-2323102000112001-3001232323122330-0203311130110310-2220100201200001-0330133300231322-3231213120233023"></a>

<a id="canonical-0113203102213002-3332220032311003-0101032023132223-1233331110021210-0002102020202122-2303022013010223-3103220232112312-0203232222020122"></a>

#### `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` property

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

<a id="canonical-2202022223012232-1130032013220201-2112021321230103-3201200133303301-1113332303020231-1002132332112122-0022210131330211-1121011132300300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.http_config.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.http_config](data-sources--alert_receiver--reference--group-001.md#canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012)
- [webhook.http_config.use_tls](data-sources--alert_receiver--reference--group-001.md#canonical-3302303211231132-3333130311332333-2212310130112020-0232103121300113-0001310332333303-3122223032322012-1133302303231230-1112203023223132)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="canonical-1303320001222232-3333211233233121-2202303013322000-0033003110021001-1313112031012312-2123121013323223-3301031122010300-0131133013000331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

<a id="canonical-3320213310121301-2102313133202001-1222223320030002-0210110101002111-2033020200031033-0021031212201202-1201202021022300-3332013032320210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.url` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- webhook.URL

<a id="canonical-1330332210233223-0011020202322002-3020313120320320-1002102000023212-2013113020022323-2101003022331333-1023132123330121-2000012020220020"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-0300120003200312-2130023010311230-1203110233231301-1113112111212230-3133230211020312-3031121213330300-3032332301113310-0313232232132222"></a>

### Direct properties for `webhook.url`

- [blindfold_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-1111322000011222-3000003103231301-1211011322020203-1321321223312301-2203000021303233-0031121220030001-2113313021033300-2111101312100203): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--reference--group-001.md#canonical-0020032311303213-1010222031122012-2313002103132233-2302333110122033-0033002112212101-1300331323030213-1020010123323003-1100101100212311): complete subsection reference.

<a id="canonical-1111322000011222-3000003103231301-1211011322020203-1321321223312301-2203000021303233-0031121220030001-2113313021033300-2111101312100203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-3320213310121301-2102313133202001-1222223320030002-0210110101002111-2033020200031033-0021031212201202-1201202021022300-3332013032320210)
- webhook.URL.blindfold_secret_info

<a id="canonical-2310133302133001-0112001310332232-0032232213210030-1301200320223212-3313230121121232-0321021210303212-3130232303201133-3000230323232222"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1101312133132223-3110321101023130-1032222012131221-0312012100112123-3302030001021322-0232331201313001-0001320133122133-2011001010310202"></a>

### Direct properties for `webhook.url.blindfold_secret_info`

<a id="canonical-1200330032012003-2112302033130131-1001300002321303-1213023223201313-1220313111010032-0200212333321303-1123010133131221-0101301111120211"></a>

#### `webhook.url.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3013203230122311-3113131123223110-1332323211231220-1211130333213330-1112300110322321-1003033320213301-0332122102010100-1300233123201313"></a>

<a id="canonical-0021000333321321-0312003301320101-0011210110310200-0000200211010203-2031233221023021-0332313233000330-2020101233100001-3321202221010020"></a>

#### `webhook.url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1113030110013222-1131031203303313-1202111113332010-0032131133123000-2011111121303110-1213230231322112-3303110020020333-3212002322130331"></a>

<a id="canonical-2122010233233131-0121011200330032-3213311330030122-2110320211303303-0222003301131120-2213202023203013-3100032213032220-3203320120100030"></a>

#### `webhook.url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0020032311303213-1010222031122012-2313002103132233-2302333110122033-0033002112212101-1300331323030213-1020010123323003-1100101100212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `webhook.url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Property reference](data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [webhook](data-sources--alert_receiver--reference--group-001.md#canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303)
- [webhook.url](data-sources--alert_receiver--reference--group-001.md#canonical-3320213310121301-2102313133202001-1222223320030002-0210110101002111-2033020200031033-0021031212201202-1201202021022300-3332013032320210)
- webhook.URL.clear_secret_info

<a id="canonical-3033003221333022-1001203331132121-3103133110030302-0121322002012022-1201012300322000-1332132331212121-3200223200213333-0333331330010332"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-3232301002231012-1302032123203200-1032213030012010-1220221100211332-2010121220113121-2202222333102232-1211132311000200-3320333210021232"></a>

### Direct properties for `webhook.url.clear_secret_info`

<a id="canonical-1232220112330221-2002132311021202-1011031223010010-2123131331220131-0010133311333331-0311302003031002-2313112133322310-2233113220303202"></a>

#### `webhook.url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2322320133303103-2012233112322011-1313030201222203-0133132022233001-3112312303113203-0310322133211032-3320213021012011-1323303021021132"></a>

<a id="canonical-3311213203223031-2033200301132220-1021122131101223-3131012300212003-1330001202100001-3113013312001002-3221133011132313-0103032000002221"></a>

#### `webhook.url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
