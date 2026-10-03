---
page_title: "xcsh_alert_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver reference."
---

# xcsh_alert_receiver reference

<a id="canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000312002030320-1332011133100233-2211112020231222-1132222321201033-0223322133013123-2200320100223310-2223030123021203-1123303213233332"></a>

## Property reference — Property reference / 332122202031 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- Property reference

<a id="canonical-0233303011310122-2230300201022222-3203231002212311-0321000310113010-2122023213032311-3000122100233200-2210110133101113-3112321331003000"></a>

## Direct properties — Property reference / 332122202031 / 3

<a id="canonical-3011321101221331-1222100333203333-1320322032023301-0120312230032001-0132210032300112-0002112032113211-1012213311020323-2303220120203221"></a>

<a id="canonical-0232010201033321-0011221310302220-0201111112001322-2013331033110102-3211102202203310-1311112021133032-2300303211022021-2211000101322221"></a>

## annotations property — Property reference / 332122202031 / 4

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

<a id="canonical-0000333320321132-3131211301022202-1130221110130003-0103200220122202-2132312021201203-0322012200311233-3233332211033321-3132311201011021"></a>

<a id="canonical-1012331202020133-0223122230211211-1303002222213220-0301121122101001-0201102130011033-2113300312100010-3133303132230002-0103321330103313"></a>

## description property — Property reference / 332122202031 / 5

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

<a id="canonical-2132301123110120-3210303101100221-2232311211111220-0021223100001103-3222002212203311-3302003322032001-1203331201211012-3120120332132310"></a>

<a id="canonical-1312313211120120-1000033303202000-3132002113120032-3321110010332332-0310301311021333-3102201302201120-1200011222300103-2212202112030223"></a>

## disable property — Property reference / 332122202031 / 6

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

- [email](resources--alert_receiver--reference--group-001.md#canonical-1323022032232303-1330320332312203-2011200000313323-0131203103332132-0201310302130131-1010310000301301-3032213102033322-2120230210032032): complete subsection reference.

<a id="canonical-3132132302010200-3233131203323122-2000001230313323-3003130301332303-0233313001111011-1323113100312202-1022031032310311-1111302313032232"></a>

<a id="canonical-1000112010301332-0210120103112223-0213012331321010-3201033123320000-1021232032312000-0211123332030120-1001220210031212-0222321303112002"></a>

## ID property — Property reference / 332122202031 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2131131301221231-2100223301231002-2023110210120010-0000201011022122-1121131011320122-2022122303222030-3201332320103121-2033201312322130"></a>

<a id="canonical-0110321220320313-1313220023221203-2211310302022311-3330131121033002-3102212321201121-2131231333001313-2133201322122332-0211303122323123"></a>

## labels property — Property reference / 332122202031 / 8

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

<a id="canonical-1300123313023212-1222212100232121-2121200221133330-3321101113111201-3103033323210322-3113030102001123-1211030333011020-1112333013121102"></a>

<a id="canonical-2101100133011131-2113321002011202-0211132023322102-3223101233321231-3031313112030300-1202000202012013-0001103311211011-3001020313221300"></a>

## name property — Property reference / 332122202031 / 9

Type: `"string"`. Required.

Name of the Alert Receiver. Must be unique within the namespace.

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

<a id="canonical-1100221002223301-2103100012330023-0111321020312321-0103021311011100-1301122221202213-1022022133202210-0011322010022212-1232031002101033"></a>

<a id="canonical-1132233010102101-1232221333023032-3301323313231230-2310131323123313-0010201302031021-1021021212223120-1223302302002001-0231330113232032"></a>

## namespace property — Property reference / 332122202031 / 10

Type: `"string"`. Required.

Namespace where the Alert Receiver is created.

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

- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121): complete subsection reference.

- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001): complete subsection reference.

- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011): complete subsection reference.

- [sms](resources--alert_receiver--reference--group-001.md#canonical-2321322330122322-3323101322120202-0332000331221112-1333103013101001-3330200333023112-3123201202122121-1121033313113332-2312331010332202): complete subsection reference.

- [timeouts](resources--alert_receiver--reference--group-001.md#canonical-1031122120302100-1202022023002122-0331303300111330-2033331012213033-0230220022302311-0210013101031223-2231203211323301-1110212233302221): complete subsection reference.

- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103): complete subsection reference.

<a id="canonical-0221233303232303-0001233001132000-2033222330313223-0012203211321110-3223310312023030-2013000200013311-0133013102002003-1331000223300031"></a>

## All schema paths — Property reference / 332122202031 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--alert_receiver--reference--group-001.md#canonical-3011321101221331-1222100333203333-1320322032023301-0120312230032001-0132210032300112-0002112032113211-1012213311020323-2303220120203221) |
| `description` | [description](resources--alert_receiver--reference--group-001.md#canonical-0000333320321132-3131211301022202-1130221110130003-0103200220122202-2132312021201203-0322012200311233-3233332211033321-3132311201011021) |
| `disable` | [disable](resources--alert_receiver--reference--group-001.md#canonical-2132301123110120-3210303101100221-2232311211111220-0021223100001103-3222002212203311-3302003322032001-1203331201211012-3120120332132310) |
| `email` | [email](resources--alert_receiver--reference--group-001.md#canonical-0030002300201200-0302300301201222-3233111222122103-0021220120133133-2232122222022210-1201221103321010-2313012100310130-0033220222122103) |
| `email.email` | [email.email](resources--alert_receiver--reference--group-001.md#canonical-0120333330213020-1121213030200101-3221331011031223-1333102133133022-3211201112203103-3231130030200002-2322212302021022-3311201012133112) |
| `id` | [ID](resources--alert_receiver--reference--group-001.md#canonical-3132132302010200-3233131203323122-2000001230313323-3003130301332303-0233313001111011-1323113100312202-1022031032310311-1111302313032232) |
| `labels` | [labels](resources--alert_receiver--reference--group-001.md#canonical-2131131301221231-2100223301231002-2023110210120010-0000201011022122-1121131011320122-2022122303222030-3201332320103121-2033201312322130) |
| `name` | [name](resources--alert_receiver--reference--group-001.md#canonical-1300123313023212-1222212100232121-2121200221133330-3321101113111201-3103033323210322-3113030102001123-1211030333011020-1112333013121102) |
| `namespace` | [namespace](resources--alert_receiver--reference--group-001.md#canonical-1100221002223301-2103100012330023-0111321020312321-0103021311011100-1301122221202213-1022022133202210-0011322010022212-1232031002101033) |
| `opsgenie` | [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3111131130312030-2101320011011203-1131232002201131-1133333233121002-1030301213202211-1202213131112120-0203200110121030-1203330020201323) |
| `opsgenie.api_key` | [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-3121021220200000-0110120211113011-1131331111103211-3012022331110130-1131130233202110-3312301120132030-1202322210102130-1131203002300220) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0303103231203302-0222130313032331-2301003133001302-2333330312222113-0030222231113122-3301333002013132-3133313132331023-1000203231211021) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-0223231301000211-0003301012310100-2320112301210101-2213020221103213-1011312033320313-0212331100221010-3332010003103130-0132303023023313) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-2110311012200201-1300302020013101-2211033101032331-2201010200100313-1022203312323333-1022010002122303-1023213333211123-1300320231011022) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-2022111112312303-0233120210032002-0033213203233130-3011202322010333-1201130103102032-2131310303033023-0322013231002021-3132100001301202) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2130031232220211-2201320220311310-1201210231230212-3321322200021213-3131303313002022-0012112111132130-1310300023110131-2022031313003102) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-0223222230103111-3311320001202230-3122011213101031-2013200131200323-1231322113332221-1213101322133213-0131220221213123-1030300212232211) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-1102021221230330-2232033321231033-0322022032233200-3212101210021003-1122023001001010-1211120301012012-3202233302133311-2113311302003101) |
| `opsgenie.url` | [opsgenie.url](resources--alert_receiver--reference--group-001.md#canonical-1212231100313323-2212212130011322-1210332110223100-0201033001010320-0330200011121101-0310111130321311-2301322212123032-0103132321222333) |
| `pagerduty` | [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-3100300032012333-1312221311201212-0211101212301203-2211201103303122-0023221201123232-1112110011331323-3132331130220031-1321210321332303) |
| `pagerduty.routing_key` | [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0312120231223213-0230123013020021-0122130001032303-1300213211323232-1000230002231323-1321202002221233-0311202111112333-2000330023300301) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0022330003333020-0131311233302020-0132122032231301-2001221203113221-2233100212030023-3203100103320231-3000101232131122-2020101202201312) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-3033313132221101-3001223302301311-0301213203320000-2112301211330230-1200120320322020-3212320131100331-1021001133123333-1022212310231203) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-2221101221130102-3121000301312222-1121131323223230-0313032222013013-0322302030010102-1032202322302313-3001111202130102-0132110230331231) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-2320020130022231-2303111110303013-0133011033101222-1130020110130113-3231131302311213-1210022230132333-0323210310132301-2313332033131122) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-3233333303000033-3223011230312120-1011011310231112-3123022032331102-0310203000232033-0322030021111132-1123232002123312-1111200110202001) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-1331110231232202-0113010002233112-0012122112122020-0020012102301100-2322200132222223-0230013223202111-2100311130133100-2001022200303332) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-2220102011031033-0200100123022012-0333211302200220-1202330110330022-0222011313203333-3113023303101030-0312123223033122-0300133313031201) |
| `pagerduty.url` | [pagerduty.url](resources--alert_receiver--reference--group-001.md#canonical-0232002121323110-1333130123113103-1122223012121021-0221321101200332-1020010011031021-2233210320321333-1110001122232322-1030231113120132) |
| `slack` | [Slack](resources--alert_receiver--reference--group-001.md#canonical-1320300122010130-2232222012202300-0030121030003303-0303223233322220-1020231100330012-1011120111003022-0313132333312233-2100230022031022) |
| `slack.channel` | [slack.channel](resources--alert_receiver--reference--group-001.md#canonical-1013013103303122-0323323301321131-2221102231000133-1310302113102332-0001311203100302-0030023220311333-1000302210220132-0203131021032231) |
| `slack.url` | [slack.url](resources--alert_receiver--reference--group-001.md#canonical-1131033103000113-3211113123030231-2211200132333332-1333020310121202-3003300332011212-0021200302123203-0103200322310003-1212333002313211) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0202033323122230-0001133233200222-1200111001210013-0312302102212303-0111203030323123-3002230021133312-1213231322301223-3023100010133032) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-0122303310033313-0112230201133223-1023313121011123-3031121111203131-0331002313001331-3332121303013321-1100001300132310-1333301103231033) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-0011310002332301-3131023210101311-3013333112213232-3211202310030131-0323323122002100-1303133203121032-2311001201123000-2212200013020322) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-0103213032031303-1120003322120303-1310101032100120-2023102022322031-0012230023123103-2102013003132233-3123222013033322-2220102323311321) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0122121122133202-3023223032230111-1002133210303311-2231030110302210-3222321102322233-1131111102010310-1021313320302201-1203313231003200) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-1130330131022133-0033012323311302-0033021233013200-1013133103222132-2321100030120312-1113231210220222-1203101330333123-0130102212220302) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-3330023300002312-2032031313113132-0131103113212122-2133302122330030-1333211010122133-1101022312112210-1303323323323020-2310200010202301) |
| `sms` | [sms](resources--alert_receiver--reference--group-001.md#canonical-2001203113113030-1001301312011133-3330110123200020-0021112321022122-3322120000332333-0210220131011131-0000132230201220-0020031131323032) |
| `sms.contact_number` | [sms.contact_number](resources--alert_receiver--reference--group-001.md#canonical-0020223331303213-2110302020030302-3220031011122013-1022133102012032-0023221131203333-3012003011220023-2000102221322021-3033033010303110) |
| `timeouts` | [timeouts](resources--alert_receiver--reference--group-001.md#canonical-3111010001003100-0301322333220113-3030211110332331-1000313123023121-1000313201231203-1130011030102303-1001010311101022-2221102112022000) |
| `timeouts.create` | [timeouts.create](resources--alert_receiver--reference--group-001.md#canonical-2133111332111010-2210112113323221-0100322202313011-2322323122111202-0120113130111332-3313222313110120-2332203211203202-0102233012100221) |
| `timeouts.delete` | [timeouts.delete](resources--alert_receiver--reference--group-001.md#canonical-2003200003120322-1103010100130121-2102000301031033-3231101031011323-3003332113322302-0332033113013211-1101310011132230-1332231301303220) |
| `timeouts.read` | [timeouts.read](resources--alert_receiver--reference--group-001.md#canonical-3133323033210033-1011230031032132-1122203301201312-0231321110330203-2012221301122003-1012321213131200-3331311033300001-1332111100021321) |
| `timeouts.update` | [timeouts.update](resources--alert_receiver--reference--group-001.md#canonical-2222311230213212-3133310133210333-1113300232310332-0111031212033121-2231331323203023-0133210000203102-2323022010033202-3021302123102121) |
| `webhook` | [webhook](resources--alert_receiver--reference--group-001.md#canonical-1333331100130003-0313223303332030-1311330010131002-3012000121222322-2030131001020322-2202221030133121-2121100010101122-0302312312032313) |
| `webhook.http_config` | [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-2010120033333122-0132013111022130-2213201013110133-3123021230201131-2011203131020310-3323210123101120-1302012221122200-2100031130031223) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-1013233330332110-0103112021213021-3330121032203011-2211233211103101-2203313232100002-0033311213020303-2311220011031212-2220000103330222) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-2022301020102203-0021031310130013-1301312023221133-3013232233321231-2011331330022200-3122331002220112-2100321111331110-0123203231121300) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-3010103232121003-0001302021111313-0131213010222302-0030220210222323-0230211001322322-3213012032132131-1332200113301302-3101112302333231) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-0332222323211312-0301130000201202-2010012020210222-2133031333020133-0013023101310223-1332122312121211-1112333120133200-1212122113100321) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-3211321200203312-3232030023102222-3323311133112332-0033113230222330-3213101001001311-1030300032111113-2332333330023022-2100232210332123) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-0230230323303100-0100330123233012-1233123322000222-3231223110110011-0130132332131221-0022320121023222-0102213321023312-2011321020022012) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2211302323003230-1201211332033223-3332330103110033-0212131102321101-1011112122132220-2201310103033031-0000000123331220-0101320312121301) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-0200333222031131-2011123330332113-2110030311101123-2031033003332113-3110021303112011-2331031302200213-1232313112233031-3311123112030321) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-2203232232033133-2201001223001000-2310010011202120-2103133222130031-1120312213121021-3022200031311011-0103102220321203-1113021200320200) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2132321110302210-1123002011303202-3231322022201031-2001311220022330-2201011313133011-0332223303203101-2031003113031220-3121202032322123) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1131033123221300-1211101023322300-3300232110223023-1212303331321131-3101023313222120-0230211121322012-3012022320321232-0100033210302203) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-3230330233113323-2020323331212310-2202103232112220-3201201203020333-1102102033311321-3100310230131112-0322211302023123-1323333312311122) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-0002233002312011-1033312203200133-3030223212221023-0131031112231212-0220232132203033-3000102330223211-1013220302110002-3023000221321201) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-2310201012131032-2200200132021000-0330213313311211-3011002211233213-2320313223030000-1021130001321232-3331320213113011-2200110232210222) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-0202212111230222-3132300332032130-2000122310321222-0112233103022030-3022102103231111-0311333032333011-0121332030212221-2202100232100132) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1233230332311121-2131122010012120-0000210321203211-1113220003111020-2111223300100030-0101100303211030-2222120132001202-0121013000031000) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-0113322303011031-0211303011212212-0222003103303203-0231320030322203-2101223103013232-3022223112310310-3002013011211123-0322300031203230) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-3010002132000302-3101012333331012-0123220022312311-1223103001022020-0230221011022010-3333313222330010-2332201111213103-0000300212320033) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](resources--alert_receiver--reference--group-001.md#canonical-0312331210333010-3220133212302312-0213033321122133-3103211120003122-1320002300210120-3313132012210201-2301122222312001-1302030102021002) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-3113201222033133-1132310033201201-3021000213120200-3332320011221213-1212203322320012-1310023310112331-3020100022233300-1232102013330323) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-1210212231300131-1131030033131310-3011213022311031-1202332010322120-0211210322200010-3300101303221110-3023021201101021-2011210203233301) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](resources--alert_receiver--reference--group-001.md#canonical-2223000133011203-0031301211133122-0302231223021322-0310222010120012-0313203300213202-3303100311113100-2203232200313001-1110213211111332) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](resources--alert_receiver--reference--group-001.md#canonical-0121321000021301-2022201010202213-0032112333223121-1100111111131012-0021311000232323-3101002212113002-2320300233221201-0223003210021313) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](resources--alert_receiver--reference--group-001.md#canonical-1000131010012321-0020212230313211-1303123220330123-0022100123313232-1333202002220212-0023301301202202-2212232203131231-0133320300120331) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](resources--alert_receiver--reference--group-001.md#canonical-0002130200133230-3101211102330321-2101023211002230-2133010210323112-1002311332011012-0313013212033210-2121330123230012-0122102023323330) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](resources--alert_receiver--reference--group-001.md#canonical-0230231311113020-1233000310302012-1312320003131112-3201331102203123-2021330013322123-0130103002030121-2033223302032102-1003300023322302) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](resources--alert_receiver--reference--group-001.md#canonical-3221133213121102-0022033310020232-2113123311001000-1132323001012300-3300213022003112-2223100113030213-1120121331002103-1321100322210120) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](resources--alert_receiver--reference--group-001.md#canonical-2033000313233003-0013302320232002-1201212021332120-1332211113002111-0000020030031221-3220102013100313-2323312222203130-1002232303020203) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](resources--alert_receiver--reference--group-001.md#canonical-3122301322312032-3333300213203211-2223212030332133-2202120200232110-3331112021302122-1001000000133323-0032101201133113-3320020030132331) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](resources--alert_receiver--reference--group-001.md#canonical-3000033011022110-3033303331221103-0322320031203301-2130020022331303-3331222010203100-3201310220311120-1013132200232220-3122000020230003) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-3203003230133131-2132222112332302-3001002000033331-1031112202031122-3033022113222201-0200013300301033-2210100111023112-2231311203123310) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](resources--alert_receiver--reference--group-001.md#canonical-2233220301111311-3001232323312030-3132210102232110-2320110222330110-1333122332110011-2100231033111100-2313223021310111-2212302200002311) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](resources--alert_receiver--reference--group-001.md#canonical-2132311132110310-2322132300003323-2012300213020112-0002331323231231-3300001302101011-0000021210130132-3102311113101003-1031002113210221) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](resources--alert_receiver--reference--group-001.md#canonical-1111111120220200-0112133231120332-1031330312022122-2303132010131321-3303032033001220-1133322212103122-0001222030302203-3233310300023031) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](resources--alert_receiver--reference--group-001.md#canonical-1103321111032013-2101322120221000-0301003122112232-2011320103033223-1222323213003233-1303320213131310-0323131011223212-3212131002201331) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-1130011322013110-1310113121022120-3021113120331022-3132222200302332-3310103132331101-0323021100133232-1112333120203321-2031033213021231) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-2230233220111002-3110010312003033-3220312013033201-1200013233331230-2110230233110132-0310323011133002-3331333211322202-0001020111112222) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-1111020022202313-1203220000100031-2321212023122331-0230003311133320-3033300032220211-0312222130111311-3011023303221331-2022303030012001) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](resources--alert_receiver--reference--group-001.md#canonical-2213033213222313-2321020113320011-1313200312302122-0302031110012231-0120303023212002-1222031112203320-2101333121202220-3132130103120213) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](resources--alert_receiver--reference--group-001.md#canonical-1220111222223030-3233012313233102-2333122301023131-2322323102323330-0120203111202030-3230010303222213-1001223120211112-2132213200303023) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](resources--alert_receiver--reference--group-001.md#canonical-2333223211132103-1322011223100320-2103101002301120-2101233320013132-1122201210020221-0133322312300200-2010223120300201-0331010111013233) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](resources--alert_receiver--reference--group-001.md#canonical-1102322330022330-0301022310130002-1312320113201131-2322021310112210-2200202202323330-3103112122003230-1333102331103133-2323210231232321) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](resources--alert_receiver--reference--group-001.md#canonical-2133231320020020-1313202033000122-3023100301133311-1320302113333101-1202013131313331-3323100001021223-1333022011002023-3232222300232111) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-3210212011221031-3123201111011110-2323320112202230-3232101330200020-0313030200020301-0020011030133230-2232220132323011-1210030233210102) |
| `webhook.url` | [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-0200012202322320-1200012331001222-1102102031332331-3201002201303322-1022212222103122-1011232133101013-3021133201101032-2312022331230303) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0211232023101031-3102231010220121-1020020112322320-3332300311132022-3122223121023301-1332002033210130-2132000020031313-2131012103202301) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-2102022121212202-0133130203201103-3211120321033102-3322111031220313-2222212131013120-0120230133230202-3212033010312033-0303131310301220) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-2310122212201313-3132322132012323-0211302131233323-3120123320123023-2222121200231130-0031003211221020-2221113212012232-0002121020133221) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-3210310302302310-3011222033130203-3333131202233232-3032130313213032-1303102333321101-3121330022033312-2100001131102002-0032310210212323) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0321233303002210-0221201310111001-3120312013002332-3112322031123212-2233220312202211-2201003002022221-3331211311211031-0013321223220120) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-1212002112122303-3132030121012203-0131112131233303-0232223010132301-3332310133012120-2222302311230330-0230002020210332-2333332311201003) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-0232311102020321-2311203012102211-3001130213102003-0113202021233303-3012100013032102-0102031303203013-3231320222133002-2323111013222223) |

<a id="canonical-0131231122203103-1121031132230103-2223200022201320-0032310313311332-0321312323310211-1213001321330121-0002302202020311-0311110303101231"></a>

## Next pages — Property reference / 332122202031 / 12

- [email](resources--alert_receiver--reference--group-001.md#canonical-1323022032232303-1330320332312203-2011200000313323-0131203103332132-0201310302130131-1010310000301301-3032213102033322-2120230210032032)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011)
- [sms](resources--alert_receiver--reference--group-001.md#canonical-2321322330122322-3323101322120202-0332000331221112-1333103013101001-3330200333023112-3123201202122121-1121033313113332-2312331010332202)
- [timeouts](resources--alert_receiver--reference--group-001.md#canonical-1031122120302100-1202022023002122-0331303300111330-2033331012213033-0230220022302311-0210013101031223-2231203211323301-1110212233302221)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1323022032232303-1330320332312203-2011200000313323-0131203103332132-0201310302130131-1010310000301301-3032213102033322-2120230210032032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332222021000033-3001021110220313-0330300113031320-3313313203103122-1221203310330220-2222000032300031-0030302111220223-2310120221200110"></a>

## email — email / 210030103021 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- email

<a id="canonical-0030002300201200-0302300301201222-3233111222122103-0021220120133133-2232122222022210-1201221103321010-2313012100310130-0033220222122103"></a>

Type: `"object"`. single nested block, Optional.

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

- [email](resources--alert_receiver--reference--group-001.md#canonical-0030002300201200-0302300301201222-3233111222122103-0021220120133133-2232122222022210-1201221103321010-2313012100310130-0033220222122103)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3111131130312030-2101320011011203-1131232002201131-1133333233121002-1030301213202211-1202213131112120-0203200110121030-1203330020201323)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-3100300032012333-1312221311201212-0211101212301203-2211201103303122-0023221201123232-1112110011331323-3132331130220031-1321210321332303)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1320300122010130-2232222012202300-0030121030003303-0303223233322220-1020231100330012-1011120111003022-0313132333312233-2100230022031022)
- [sms](resources--alert_receiver--reference--group-001.md#canonical-2001203113113030-1001301312011133-3330110123200020-0021112321022122-3322120000332333-0210220131011131-0000132230201220-0020031131323032)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-1333331100130003-0313223303332030-1311330010131002-3012000121222322-2030131001020322-2202221030133121-2121100010101122-0302312312032313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
email {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221311022230101-0313301030020221-0032013320122121-1132202132113223-1311320002211302-0211000310011223-3233213222201030-1300021132300033"></a>

## Direct properties — email / 210030103021 / 3

<a id="canonical-0120333330213020-1121213030200101-3221331011031223-1333102133133022-3211201112203103-3231130030200002-2322212302021022-3311201012133112"></a>

<a id="canonical-1220311232313013-0131021201221023-3203013230030033-0011213313033303-2021031221202033-0313330101331110-0032211132021133-2021111200212122"></a>

## email property — email / 210030103021 / 4

Type: `"string"`. Optional.

Email. Email ID of the user.

Upstream description:

Email ID of the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101310200330322-0131031101211231-0131320311322230-2021033312222203-2101100022220232-0031031233202133-2220123302310332-3011102202112033"></a>

## Next pages — email / 210030103021 / 5

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003302311123021-2220313033300333-1130011011302301-2000102221130333-2012321120201302-3032211202333320-3013033001202221-3010033300323030"></a>

## opsgenie — opsgenie / 201311133220 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- opsgenie

<a id="canonical-3111131130312030-2101320011011203-1131232002201131-1133333233121002-1030301213202211-1202213131112120-0203200110121030-1203330020201323"></a>

Type: `"object"`. single nested block, Optional.

OpsGenie configuration to send alert notifications.

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
opsgenie {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010023030003232-2130303320220233-1002013023011013-0112020100012233-2000100010213123-0000220311130112-2030031013333203-3203300022330132"></a>

## Direct properties — opsgenie / 201311133220 / 3

- [api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102): complete subsection reference.

<a id="canonical-1212231100313323-2212212130011322-1210332110223100-0201033001010320-0330200011121101-0310111130321311-2301322212123032-0103132321222333"></a>

<a id="canonical-3332031321333111-0233203310010020-0222311122322100-3232302032010110-0223030322231113-0301222330322101-1302211210212111-1010032220020221"></a>

## URL property — opsgenie / 201311133220 / 4

Type: `"string"`. Optional.

API URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-1023122310022001-2202312112232000-3303322331302332-3001000102033010-0222001301023101-0213311112323101-2213210132003203-1222221001320220"></a>

## Next pages — opsgenie / 201311133220 / 5

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210003332310102-0102132311333132-3102232001210010-3311203313002300-1200321032213323-3031212033012233-1232202030001202-0222132002332131"></a>

## opsgenie.api_key — api_key / 002112200311 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121)
- opsgenie.api_key

<a id="canonical-3121021220200000-0110120211113011-1131331111103211-3012022331110130-1131130233202110-3312301120132030-1202322210102130-1131203002300220"></a>

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222331030123222-0221023202211102-3100312023122031-1303333023000120-0121210303012210-2203233003112033-1003222101201010-2233130033123230"></a>

## Direct properties — api_key / 002112200311 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1312110332011303-2130000010101013-3321313330120201-2303221113210313-1100323122021101-0132212311133112-3222001011303301-0012000311330301): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0012032323210232-2200101030211310-1231003302033321-0232120310321011-1231300320132303-0202010332002003-3203222333211321-1203121031000020): complete subsection reference.

<a id="canonical-0103132313200030-1211020233203100-3121200230123032-0311102103302121-1202123012333122-3333302330211012-0120033223030103-0320232002203001"></a>

## Next pages — api_key / 002112200311 / 4

- [opsgenie.api_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1312110332011303-2130000010101013-3321313330120201-2303221113210313-1100323122021101-0132212311133112-3222001011303301-0012000311330301)
- [opsgenie.api_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0012032323210232-2200101030211310-1231003302033321-0232120310321011-1231300320132303-0202010332002003-3203222333211321-1203121031000020)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1312110332011303-2130000010101013-3321313330120201-2303221113210313-1100323122021101-0132212311133112-3222001011303301-0012000311330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310003333231130-3223112223301133-1101333101122110-3023020311030312-0322032123210332-2313123003033132-1120020230322313-1032013003033201"></a>

## opsgenie.api_key.blindfold_secret_info — blindfold_secret_info / 113333200230 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121)
- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102)
- opsgenie.api_key.blindfold_secret_info

<a id="canonical-0303103231203302-0222130313032331-2301003133001302-2333330312222113-0030222231113122-3301333002013132-3133313132331023-1000203231211021"></a>

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

<a id="canonical-0320121222331331-1112203130113000-2002023213022301-3212321331011110-2313100001323031-3300301221123132-0133223111103311-1130230123222011"></a>

## Direct properties — blindfold_secret_info / 113333200230 / 3

<a id="canonical-0223231301000211-0003301012310100-2320112301210101-2213020221103213-1011312033320313-0212331100221010-3332010003103130-0132303023023313"></a>

<a id="canonical-1220120313330330-1222032020001133-1031120323031330-1231200002213301-3000213033201333-3121331220001203-2130211310301232-2223112130132020"></a>

## decryption_provider property — blindfold_secret_info / 113333200230 / 4

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

<a id="canonical-2110311012200201-1300302020013101-2211033101032331-2201010200100313-1022203312323333-1022010002122303-1023213333211123-1300320231011022"></a>

<a id="canonical-2033013313310201-2220221331103013-2231202330022111-3113223120232113-1200030123211323-2032203011032020-1013332203100103-3100203231021310"></a>

## location property — blindfold_secret_info / 113333200230 / 5

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

<a id="canonical-2022111112312303-0233120210032002-0033213203233130-3011202322010333-1201130103102032-2131310303033023-0322013231002021-3132100001301202"></a>

<a id="canonical-3032003031030230-2210100321301203-1103231200101223-0102023300201111-3302011303210211-0310222110330102-0303101121003133-0120123033213012"></a>

## store_provider property — blindfold_secret_info / 113333200230 / 6

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

<a id="canonical-3131002212023302-1313133022200303-1122030331312102-3213132303230222-3303201303300100-1312222220212132-3202230213220320-1213232020012032"></a>

## Next pages — blindfold_secret_info / 113333200230 / 7

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0012032323210232-2200101030211310-1231003302033321-0232120310321011-1231300320132303-0202010332002003-3203222333211321-1203121031000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131231032121121-3312202333030303-2233031232231130-0300203200111233-1010310233003300-1003301203301133-3302201333012101-1201112130113230"></a>

## opsgenie.api_key.clear_secret_info — clear_secret_info / 332130021223 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-3031023232312101-1202323002033211-2132032310321310-1311023022112333-0123230310221030-1201112223012231-2133110202321302-1233221122212121)
- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102)
- opsgenie.api_key.clear_secret_info

<a id="canonical-2130031232220211-2201320220311310-1201210231230212-3321322200021213-3131303313002022-0012112111132130-1310300023110131-2022031313003102"></a>

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

<a id="canonical-3111233001002002-3210212002333212-1323123332320222-0012131321331031-3232002212113013-1233002323131220-2012131333131020-2232230232331032"></a>

## Direct properties — clear_secret_info / 332130021223 / 3

<a id="canonical-0223222230103111-3311320001202230-3122011213101031-2013200131200323-1231322113332221-1213101322133213-0131220221213123-1030300212232211"></a>

<a id="canonical-2110312203032022-2223230120001231-1303231231002310-1311032021301302-0333220301221001-0233232123131020-3131033300031302-0300112123033101"></a>

## provider_ref property — clear_secret_info / 332130021223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1102021221230330-2232033321231033-0322022032233200-3212101210021003-1122023001001010-1211120301012012-3202233302133311-2113311302003101"></a>

<a id="canonical-0302223300231000-3131121211322300-1201012003100120-1212222303333021-1120222302201311-2212321003312002-3011231002200033-0010312300213021"></a>

## URL property — clear_secret_info / 332130021223 / 5

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

<a id="canonical-1021000310020032-2301312003212131-3003230120001022-0232202033223223-0112232002210100-0200110223312302-2030000303321133-3013223131210002"></a>

## Next pages — clear_secret_info / 332130021223 / 6

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-1122233210111011-1231210321232132-2120130122233230-1103301232332223-3023203111102110-3113122020001331-1123220110000111-1101331002113102)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231300012103001-2102103130132021-2123103110312020-1213221330112003-0311331011133232-1311120221002021-2120103302230232-3221220120112203"></a>

## pagerduty — pagerduty / 333100022033 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- pagerduty

<a id="canonical-3100300032012333-1312221311201212-0211101212301203-2211201103303122-0023221201123232-1112110011331323-3132331130220031-1321210321332303"></a>

Type: `"object"`. single nested block, Optional.

PagerDuty configuration to send alert notifications.

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
pagerduty {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320121123311032-3202312201012230-3232203211300210-0313030300010022-2122222333332003-3212110322012133-2323231332332113-1322202330203132"></a>

## Direct properties — pagerduty / 333100022033 / 3

- [routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111): complete subsection reference.

<a id="canonical-0232002121323110-1333130123113103-1122223012121021-0221321101200332-1020010011031021-2233210320321333-1110001122232322-1030231113120132"></a>

<a id="canonical-3311323311132102-0331001213332112-1333121321000220-0301103230220233-1230212200330003-3333030202033230-0300130122030113-3102310311211320"></a>

## URL property — pagerduty / 333100022033 / 4

Type: `"string"`. Optional.

Pager Duty URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-1212231312210002-0201110131222302-3033221233330012-0201203023010111-0233221102311200-2103300131113121-3122130120312223-0303301101331003"></a>

## Next pages — pagerduty / 333100022033 / 5

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031001123213120-1212121000303132-2132233033323301-1003230003321221-0012223220001113-2211321110012133-2030310201103101-1211022030011321"></a>

## pagerduty.routing_key — routing_key / 301110111302 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001)
- pagerduty.routing_key

<a id="canonical-0312120231223213-0230123013020021-0122130001032303-1300213211323232-1000230002231323-1321202002221233-0311202111112333-2000330023300301"></a>

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
routing_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323032202322103-1322213310332020-1200132021202121-1112233330321030-2223311113323012-3022321131230220-1332101012203231-0023323011002013"></a>

## Direct properties — routing_key / 301110111302 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1210102222120230-2221110000312013-2323320111222222-0133230130232200-2232333013112030-1320330220202132-3020112132103010-0030203220122132): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0012212331200102-0323302333033131-2013020201232212-0333302002022033-1012333223211122-2022323313133000-3011312131010021-3220213333012112): complete subsection reference.

<a id="canonical-0001013331113120-1020102000210020-3133123313033033-3001301202111311-2303301233321003-3101323201332003-0302302233233301-3300031230301133"></a>

## Next pages — routing_key / 301110111302 / 4

- [pagerduty.routing_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1210102222120230-2221110000312013-2323320111222222-0133230130232200-2232333013112030-1320330220202132-3020112132103010-0030203220122132)
- [pagerduty.routing_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0012212331200102-0323302333033131-2013020201232212-0333302002022033-1012333223211122-2022323313133000-3011312131010021-3220213333012112)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1210102222120230-2221110000312013-2323320111222222-0133230130232200-2232333013112030-1320330220202132-3020112132103010-0030203220122132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310211212120003-0011012020031331-1200210220322231-1021302032012113-3312130101313010-1203323201213312-2202302002130322-3213130101003300"></a>

## pagerduty.routing_key.blindfold_secret_info — blindfold_secret_info / 020011312223 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001)
- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111)
- pagerduty.routing_key.blindfold_secret_info

<a id="canonical-0022330003333020-0131311233302020-0132122032231301-2001221203113221-2233100212030023-3203100103320231-3000101232131122-2020101202201312"></a>

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

<a id="canonical-0121311012111302-1220111023033202-1010012110301232-3103120122002203-2201221012201022-1211211023201311-0013200323213211-1022201032300030"></a>

## Direct properties — blindfold_secret_info / 020011312223 / 3

<a id="canonical-3033313132221101-3001223302301311-0301213203320000-2112301211330230-1200120320322020-3212320131100331-1021001133123333-1022212310231203"></a>

<a id="canonical-1030111233102031-3300202220120212-0303133301220003-0131030230102122-0020321320023322-3113100132322022-2321033303221120-2030102002303311"></a>

## decryption_provider property — blindfold_secret_info / 020011312223 / 4

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

<a id="canonical-2221101221130102-3121000301312222-1121131323223230-0313032222013013-0322302030010102-1032202322302313-3001111202130102-0132110230331231"></a>

<a id="canonical-2230121201113020-3333000300111112-3313100002033000-0201010301211113-1112012313103221-3102302022330322-0201132113030020-2120000133121100"></a>

## location property — blindfold_secret_info / 020011312223 / 5

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

<a id="canonical-2320020130022231-2303111110303013-0133011033101222-1130020110130113-3231131302311213-1210022230132333-0323210310132301-2313332033131122"></a>

<a id="canonical-0100130210111312-3320133133023303-1031132103303111-2001231300320202-2032321213230201-3233300230021012-2031213132001233-2323300310321301"></a>

## store_provider property — blindfold_secret_info / 020011312223 / 6

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

<a id="canonical-1113111021201131-0111230322021201-3230222131232301-0332101333203002-3322010120122101-3101100300120220-2212102301002131-3220212332120102"></a>

## Next pages — blindfold_secret_info / 020011312223 / 7

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0012212331200102-0323302333033131-2013020201232212-0333302002022033-1012333223211122-2022323313133000-3011312131010021-3220213333012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120022121203031-2310100330300313-1021313131220031-0102312010100113-3321203110212212-1133233132123101-3203221000220203-1211222111001030"></a>

## pagerduty.routing_key.clear_secret_info — clear_secret_info / 322021030320 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001)
- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111)
- pagerduty.routing_key.clear_secret_info

<a id="canonical-3233333303000033-3223011230312120-1011011310231112-3123022032331102-0310203000232033-0322030021111132-1123232002123312-1111200110202001"></a>

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

<a id="canonical-3222123223210103-3203233101311311-3003212220323223-3010202003030002-3031013013321003-1003202000332230-2220022112113311-1311000203222210"></a>

## Direct properties — clear_secret_info / 322021030320 / 3

<a id="canonical-1331110231232202-0113010002233112-0012122112122020-0020012102301100-2322200132222223-0230013223202111-2100311130133100-2001022200303332"></a>

<a id="canonical-0330212322330222-1032130033231222-2001102200231223-0221222231010113-3333130322332001-1310133030030132-1123010012130233-2010100031123033"></a>

## provider_ref property — clear_secret_info / 322021030320 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2220102011031033-0200100123022012-0333211302200220-1202330110330022-0222011313203333-3113023303101030-0312123223033122-0300133313031201"></a>

<a id="canonical-1122113021321101-2131112301223000-1333111230032232-2101203112203230-0220312110200333-0103023203120201-0330220202102321-1301111022013322"></a>

## URL property — clear_secret_info / 322021030320 / 5

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

<a id="canonical-0121323200321323-3111001022110320-1201212232110120-3033120002133113-1023002010130211-3111233303132331-1002320131030022-1330301313121231"></a>

## Next pages — clear_secret_info / 322021030320 / 6

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202311301123132-3111023331121312-1232132211133031-1022332310002212-2033303122133330-1231231333333110-2112322130302020-0303133022121323"></a>

## Slack — Slack / 103331110310 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- Slack

<a id="canonical-1320300122010130-2232222012202300-0030121030003303-0303223233322220-1020231100330012-1011120111003022-0313132333312233-2100230022031022"></a>

Type: `"object"`. single nested block, Optional.

Slack configuration to send alert notifications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("channel")}
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
slack {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103103300111312-3000100103103330-3222110011010031-0022312323013220-2103102202200030-3022332233330010-3112222231100230-2000232210002021"></a>

## Direct properties — Slack / 103331110310 / 3

<a id="canonical-1013013103303122-0323323301321131-2221102231000133-1310302113102332-0001311203100302-0030023220311333-1000302210220132-0203131021032231"></a>

<a id="canonical-0213221213022321-1232012300110221-1302132131303000-2302103333112211-2112302233310120-3313021031020013-0313321022211022-3101002332230320"></a>

## channel property — Slack / 103331110310 / 4

Type: `"string"`. Optional.

Channel or user to send notifications to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [URL](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033): complete subsection reference.

<a id="canonical-2103021313230313-1221132000200220-0313302131311300-2031121130323121-1323331331010120-3212001300330001-2220202112031000-3200111013133200"></a>

## Next pages — Slack / 103331110310 / 5

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001131100112232-1223230101320331-2213233231212011-2312000103122132-1010022023302223-0223111023133323-3332012322031202-2233232110230011"></a>

## Slack.URL — URL / 122201110002 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011)
- Slack.URL

<a id="canonical-1131033103000113-3211113123030231-2211200132333332-1333020310121202-3003300332011212-0021200302123203-0103200322310003-1212333002313211"></a>

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
url {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102333110301301-3333021033321123-0210303030321301-1011300113133011-0022301310133211-0230123000212211-1130132102202222-3300312232310322"></a>

## Direct properties — URL / 122201110002 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0100310013112113-2121131132303202-1000200032221223-1333312333233303-2031203123202212-0000030003331022-1120233113112001-2210320221131223): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0231300212011020-3010300332120202-3201120313021321-0332130211012100-3002130021023002-3303233203301312-0112033233130321-2321232213333022): complete subsection reference.

<a id="canonical-1121122223313200-2233301101211221-2333101211210321-2200132332120112-0111030200012223-2220321300320130-0002320303013312-0103002212102130"></a>

## Next pages — URL / 122201110002 / 4

- [slack.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0100310013112113-2121131132303202-1000200032221223-1333312333233303-2031203123202212-0000030003331022-1120233113112001-2210320221131223)
- [slack.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0231300212011020-3010300332120202-3201120313021321-0332130211012100-3002130021023002-3303233203301312-0112033233130321-2321232213333022)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0100310013112113-2121131132303202-1000200032221223-1333312333233303-2031203123202212-0000030003331022-1120233113112001-2210320221131223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112333032303033-3000113213301332-0010213130303002-2322220102001122-0001120321121111-2112311130313203-1310032113121222-1222230030110010"></a>

## Slack.URL.blindfold_secret_info — blindfold_secret_info / 301330011221 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011)
- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033)
- Slack.URL.blindfold_secret_info

<a id="canonical-0202033323122230-0001133233200222-1200111001210013-0312302102212303-0111203030323123-3002230021133312-1213231322301223-3023100010133032"></a>

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

<a id="canonical-3032030103003102-0231233131211123-3231123333110312-0010323302111202-2011233010021013-0002100222301102-0332332011322311-2133323213102030"></a>

## Direct properties — blindfold_secret_info / 301330011221 / 3

<a id="canonical-0122303310033313-0112230201133223-1023313121011123-3031121111203131-0331002313001331-3332121303013321-1100001300132310-1333301103231033"></a>

<a id="canonical-2303223131311331-3023213203130322-0121000122323331-1310312333231300-1112321031213210-2213302333311202-3003013133133111-2123211230222330"></a>

## decryption_provider property — blindfold_secret_info / 301330011221 / 4

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

<a id="canonical-0011310002332301-3131023210101311-3013333112213232-3211202310030131-0323323122002100-1303133203121032-2311001201123000-2212200013020322"></a>

<a id="canonical-1022131321130022-0301233202222333-1101001001011110-1213223232323231-3320300032232333-3123033121033201-0302302200032212-2231313233310102"></a>

## location property — blindfold_secret_info / 301330011221 / 5

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

<a id="canonical-0103213032031303-1120003322120303-1310101032100120-2023102022322031-0012230023123103-2102013003132233-3123222013033322-2220102323311321"></a>

<a id="canonical-0201233031313300-2011221110030010-3211222011320210-0313202002301310-0233213332203303-1113302230113321-2311103302300010-3303120020022102"></a>

## store_provider property — blindfold_secret_info / 301330011221 / 6

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

<a id="canonical-3010301011320011-0013113313333110-2300133022103220-2210233311122122-1102302120100230-2123012111311112-2331201310211120-1312121013223320"></a>

## Next pages — blindfold_secret_info / 301330011221 / 7

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0231300212011020-3010300332120202-3201120313021321-0332130211012100-3002130021023002-3303233203301312-0112033233130321-2321232213333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323303312310320-1113002032131332-0131023230023132-0312003010112122-1233331313233121-2120113000112103-1032031200232012-2321003133100032"></a>

## Slack.URL.clear_secret_info — clear_secret_info / 210121133113 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [Slack](resources--alert_receiver--reference--group-001.md#canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011)
- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033)
- Slack.URL.clear_secret_info

<a id="canonical-0122121122133202-3023223032230111-1002133210303311-2231030110302210-3222321102322233-1131111102010310-1021313320302201-1203313231003200"></a>

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

<a id="canonical-3133313230123003-1120321010313323-0020210311233112-2021031013303303-3020320120001022-1123101221211110-2101332230010202-2333302332013223"></a>

## Direct properties — clear_secret_info / 210121133113 / 3

<a id="canonical-1130330131022133-0033012323311302-0033021233013200-1013133103222132-2321100030120312-1113231210220222-1203101330333123-0130102212220302"></a>

<a id="canonical-0000132200002322-1232233130123313-2211121031122000-1222031200122302-1012010022130123-0213021320122332-2300003132313301-0232101320032002"></a>

## provider_ref property — clear_secret_info / 210121133113 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3330023300002312-2032031313113132-0131103113212122-2133302122330030-1333211010122133-1101022312112210-1303323323323020-2310200010202301"></a>

<a id="canonical-2003231001012101-0231003131221012-0230211312133212-0002003011002111-2312301211223230-0122221232203122-0303121013010010-2130111122211111"></a>

## URL property — clear_secret_info / 210121133113 / 5

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

<a id="canonical-0302233312230013-0001211112001330-1311323130003222-3100302123320323-2303122103212122-2213010100223111-0223131210200323-0231302231203330"></a>

## Next pages — clear_secret_info / 210121133113 / 6

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-2303331113332101-1001223132303023-1131130102311020-2230022102203123-0012323302311031-1313311132332111-3212023302301011-3330133333000033)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2321322330122322-3323101322120202-0332000331221112-1333103013101001-3330200333023112-3123201202122121-1121033313113332-2312331010332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323020122223203-2012122310213222-3003020130301123-2203033210313203-0031212322132200-1231222211101333-1303002222131123-2201222233031203"></a>

## sms — sms / 321112332030 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- sms

<a id="canonical-2001203113113030-1001301312011133-3330110123200020-0021112321022122-3322120000332333-0210220131011131-0000132230201220-0020031131323032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010100121101113-1212100312310313-1303020122012202-2103203000313130-0222121112302220-2222102232312110-0121300212122022-2301310212010130"></a>

## Direct properties — sms / 321112332030 / 3

<a id="canonical-0020223331303213-2110302020030302-3220031011122013-1022133102012032-0023221131203333-3012003011220023-2000102221322021-3033033010303110"></a>

<a id="canonical-1032102311323213-2003202000032133-1211331020313201-1233102130011030-0110213231002310-3212102202003302-2331333300211113-3113123021200130"></a>

## contact_number property — sms / 321112332030 / 4

Type: `"string"`. Optional.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Upstream description:

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\]

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

<a id="canonical-0111313113322111-0002030212322122-3202221103032330-2022030321021211-1232131310300112-1101213211131132-0100222213111133-0132112111133133"></a>

## Next pages — sms / 321112332030 / 5

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1031122120302100-1202022023002122-0331303300111330-2033331012213033-0230220022302311-0210013101031223-2231203211323301-1110212233302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230301221121233-3132323323220213-1110121033221102-0200313230322102-1020220203323030-0123010321220331-1020223101332022-0332103032221131"></a>

## timeouts — timeouts / 221213330111 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- timeouts

<a id="canonical-3111010001003100-0301322333220113-3030211110332331-1000313123023121-1000313201231203-1130011030102303-1001010311101022-2221102112022000"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112202312300012-2211310222023211-2311232321103202-1133301030030013-0333001220020002-2111013132222021-0131223210330013-1213100203211321"></a>

## Direct properties — timeouts / 221213330111 / 3

<a id="canonical-2133111332111010-2210112113323221-0100322202313011-2322323122111202-0120113130111332-3313222313110120-2332203211203202-0102233012100221"></a>

<a id="canonical-2023301210032021-2313302030330033-3030112111103203-2202232323011301-3203121332223321-0021231331001023-1122233323100021-3102100120010031"></a>

## create property — timeouts / 221213330111 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2003200003120322-1103010100130121-2102000301031033-3231101031011323-3003332113322302-0332033113013211-1101310011132230-1332231301303220"></a>

<a id="canonical-0310203130323002-0122330203003310-0212302230113321-1021221003332331-2232312220023112-3002033301303323-3102022110032010-1002101211232223"></a>

## delete property — timeouts / 221213330111 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3133323033210033-1011230031032132-1122203301201312-0231321110330203-2012221301122003-1012321213131200-3331311033300001-1332111100021321"></a>

<a id="canonical-3312121201223000-1302210103013031-3112021200003121-2122121003030302-1323020130322320-2332222210231313-3233332011231023-1020230020233130"></a>

## read property — timeouts / 221213330111 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2222311230213212-3133310133210333-1113300232310332-0111031212033121-2231331323203023-0133210000203102-2323022010033202-3021302123102121"></a>

<a id="canonical-2312220011122110-2312231223101103-2302200000000122-2012311313003300-2210303101213203-3113010231100202-1230130032122223-3123332030103310"></a>

## update property — timeouts / 221213330111 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1323313200110121-0131131011131133-3332330110102013-0233102102211200-0003223000023031-2302002120320101-3032002130210101-3331033132213011"></a>

## Next pages — timeouts / 221213330111 / 8

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330010200321121-1212333022201001-2120201001121102-2323213320212133-1132033311202133-0300131330020203-3230130130032010-1211333330021300"></a>

## webhook — webhook / 100032011022 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- webhook

<a id="canonical-1333331100130003-0313223303332030-1311330010131002-3012000121222322-2030131001020322-2202221030133121-2121100010101122-0302312312032313"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
webhook {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303121111023321-3323220001332212-0203333131300132-2203001013322310-1032332133320031-3111032212000302-3200102113013222-3211202113131021"></a>

## Direct properties — webhook / 100032011022 / 3

- [http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022): complete subsection reference.

- [URL](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110): complete subsection reference.

<a id="canonical-1101211001323003-2013332323131202-2013020113003320-0102211002311230-0031030032101332-1001212223210003-2320010321122331-3300112012232020"></a>

## Next pages — webhook / 100032011022 / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011021212132210-1233102133023131-1213021213212100-3303003131101103-0130000132223321-3233022310213102-1211331110002031-1111311100320022"></a>

## webhook.http_config — http_config / 010002020210 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- webhook.http_config

<a id="canonical-2010120033333122-0132013111022130-2213201013110133-3123021230201131-2011203131020310-3323210123101120-1302012221122200-2100031130031223"></a>

Type: `"object"`. single nested block, Optional.

HTTP Configuration. Configuration for HTTP endpoint.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_token",
    "basic_auth"),
  validators.ConflictingObjectAttributes("auth_token",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("auth_token",
    "no_authorization"),
  validators.ConflictingObjectAttributes("basic_auth",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("basic_auth",
    "no_authorization"),
  validators.ConflictingObjectAttributes("client_cert_obj",
    "no_authorization"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112013103012120-0001300320202222-0221001132123300-0332001233231131-1111131123131230-3001010111022313-2303303030221113-1220201212312033"></a>

## Direct properties — http_config / 010002020210 / 3

- [auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312): complete subsection reference.

- [basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233): complete subsection reference.

- [client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100): complete subsection reference.

<a id="canonical-3221133213121102-0022033310020232-2113123311001000-1132323001012300-3300213022003112-2223100113030213-1120121331002103-1321100322210120"></a>

<a id="canonical-1120212021301202-3233001231313023-2210120030320232-1312012332222120-1312100122113021-0022110011102022-2133033030323211-3301212101310300"></a>

## enable_http2 property — http_config / 010002020210 / 4

Type: `"bool"`. Optional.

Enable HTTP2. Configure to use HTTP2 protocol.

Upstream description:

Configure to use HTTP2 protocol.

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

<a id="canonical-2033000313233003-0013302320232002-1201212021332120-1332211113002111-0000020030031221-3220102013100313-2323312222203130-1002232303020203"></a>

<a id="canonical-2201030312210311-1310101011222032-0111301011331303-3312221332212100-3113110110232312-2111131221033113-2332000322223022-3321101011020323"></a>

## follow_redirects property — http_config / 010002020210 / 5

Type: `"bool"`. Optional.

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

- [no_authorization](resources--alert_receiver--reference--group-001.md#canonical-3001330103012311-2320121223330230-2011300030332132-3232110222030120-1130232011331212-1213030222332331-3231203321321130-2031122012102232): complete subsection reference.

- [no_tls](resources--alert_receiver--reference--group-001.md#canonical-3322110000220223-2002013220102223-2301310132030110-1131321030322130-3023232113100010-3013022200212212-0103033022220323-2003310000213130): complete subsection reference.

- [use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213): complete subsection reference.

<a id="canonical-1120000002123332-1333310030022210-1120102233213330-2210231131111202-3310023012102001-3120100102112311-3221213001232123-2230331113012223"></a>

## Next pages — http_config / 010002020210 / 6

- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233)
- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100)
- [webhook.http_config.no_authorization](resources--alert_receiver--reference--group-001.md#canonical-3001330103012311-2320121223330230-2011300030332132-3232110222030120-1130232011331212-1213030222332331-3231203321321130-2031122012102232)
- [webhook.http_config.no_tls](resources--alert_receiver--reference--group-001.md#canonical-3322110000220223-2002013220102223-2301310132030110-1131321030322130-3023232113100010-3013022200212212-0103033022220323-2003310000213130)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230101020113310-0102302003110112-2123210101232111-0310311123111110-3100311331202333-2322303223313230-3300030303130201-0203320211301301"></a>

## webhook.http_config.auth_token — auth_token / 213012022203 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.auth_token

<a id="canonical-1013233330332110-0103112021213021-3330121032203011-2211233211103101-2203313232100002-0033311213020303-2311220011031212-2220000103330222"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

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
auth_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301032223032323-3122123010113313-3301210200000201-2000322223032002-0223012332001303-2313101121303201-1113331002033021-3003023301000303"></a>

## Direct properties — auth_token / 213012022203 / 3

- [token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022): complete subsection reference.

<a id="canonical-3032003310133021-2331211133031000-2120322111311223-3210330131131310-1221023130131221-2301230103033303-3303133201302123-2103300001310002"></a>

## Next pages — auth_token / 213012022203 / 4

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233112011032300-2300103302130102-3230213232330102-2221031301221210-2023103132232210-2023002103300013-1300023223300022-3302101320022221"></a>

## webhook.http_config.auth_token.token — token / 331110213013 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312)
- webhook.http_config.auth_token.token

<a id="canonical-2022301020102203-0021031310130013-1301312023221133-3013232233321231-2011331330022200-3122331002220112-2100321111331110-0123203231121300"></a>

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
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113320133021311-3132021220000112-3323301301113213-2213223310212000-0120113123220033-0103200103203001-1001002021201223-2030113020002000"></a>

## Direct properties — token / 331110213013 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1203313303020333-3001003110100123-1101323310203210-2001201300303003-0201111223211302-3021203133302230-3221132201311313-2303122321031110): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2330202311221103-2213331010031213-2220200231020121-0112131120111222-1101022113203323-2203003112222322-1230133021230201-1003202221231230): complete subsection reference.

<a id="canonical-2010211203221301-0030213022330302-1201323230331232-2331120013223300-1210113320323112-3332312331001131-3200302021203122-2023233023132301"></a>

## Next pages — token / 331110213013 / 4

- [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1203313303020333-3001003110100123-1101323310203210-2001201300303003-0201111223211302-3021203133302230-3221132201311313-2303122321031110)
- [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2330202311221103-2213331010031213-2220200231020121-0112131120111222-1101022113203323-2203003112222322-1230133021230201-1003202221231230)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1203313303020333-3001003110100123-1101323310203210-2001201300303003-0201111223211302-3021203133302230-3221132201311313-2303122321031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323002110013220-0123102300100233-2001130132031123-0131012113033332-1303121123021100-2331323212210020-0030230031330020-3200133001002021"></a>

## webhook.http_config.auth_token.token.blindfold_secret_info — blindfold_secret_info / 321110101213 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312)
- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022)
- webhook.http_config.auth_token.token.blindfold_secret_info

<a id="canonical-3010103232121003-0001302021111313-0131213010222302-0030220210222323-0230211001322322-3213012032132131-1332200113301302-3101112302333231"></a>

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

<a id="canonical-3310013300221223-3130110323103022-2130212131320210-0033002113113302-0111133233123111-3033211303201032-0101113130230130-2331131221013213"></a>

## Direct properties — blindfold_secret_info / 321110101213 / 3

<a id="canonical-0332222323211312-0301130000201202-2010012020210222-2133031333020133-0013023101310223-1332122312121211-1112333120133200-1212122113100321"></a>

<a id="canonical-3120031021022332-3200102230212023-2322102222332333-1330321331133013-2310101023201013-1102120001201202-0331001210233233-0332131021111332"></a>

## decryption_provider property — blindfold_secret_info / 321110101213 / 4

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

<a id="canonical-3211321200203312-3232030023102222-3323311133112332-0033113230222330-3213101001001311-1030300032111113-2332333330023022-2100232210332123"></a>

<a id="canonical-0103131010330022-3231113202320232-2032203320332110-2232201313201110-1011211221313030-1110311031110302-1131111002121003-1112122023012322"></a>

## location property — blindfold_secret_info / 321110101213 / 5

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

<a id="canonical-0230230323303100-0100330123233012-1233123322000222-3231223110110011-0130132332131221-0022320121023222-0102213321023312-2011321020022012"></a>

<a id="canonical-0221300303222311-1110330100113210-0310031331122333-2033310133200131-2020112231222020-3322020121102022-0302110302010200-1321312221133101"></a>

## store_provider property — blindfold_secret_info / 321110101213 / 6

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

<a id="canonical-3131330212101111-0003013111313320-0210333302111202-3111133220012223-3223112113212211-0003013232131023-2030231230023213-3132111030330120"></a>

## Next pages — blindfold_secret_info / 321110101213 / 7

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2330202311221103-2213331010031213-2220200231020121-0112131120111222-1101022113203323-2203003112222322-1230133021230201-1003202221231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102223212233003-0130222133220320-0322003213233013-2113122121020213-2330100030111201-0310330321210323-0211133212102223-1230333230232130"></a>

## webhook.http_config.auth_token.token.clear_secret_info — clear_secret_info / 223212333223 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-3020311131213031-1020223133201012-3210330203321111-2302203020010020-1111303003131230-2331200011011211-2313311130212123-0220112010033312)
- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022)
- webhook.http_config.auth_token.token.clear_secret_info

<a id="canonical-2211302323003230-1201211332033223-3332330103110033-0212131102321101-1011112122132220-2201310103033031-0000000123331220-0101320312121301"></a>

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

<a id="canonical-3013110122111321-3213210201011102-3332123333301133-3112013031302333-1023133011221113-1003000001100330-0103011233231012-2302221130202002"></a>

## Direct properties — clear_secret_info / 223212333223 / 3

<a id="canonical-0200333222031131-2011123330332113-2110030311101123-2031033003332113-3110021303112011-2331031302200213-1232313112233031-3311123112030321"></a>

<a id="canonical-1013013133022030-2010032111112010-1122222011322012-1203012023132223-3022021112313132-1100230030303100-1133211031232030-1222233020321132"></a>

## provider_ref property — clear_secret_info / 223212333223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2203232232033133-2201001223001000-2310010011202120-2103133222130031-1120312213121021-3022200031311011-0103102220321203-1113021200320200"></a>

<a id="canonical-1313121221230000-2002012003120330-3100122133323030-3221201130320011-2123312003120200-2303130122303013-2202013313122032-2301032220030213"></a>

## URL property — clear_secret_info / 223212333223 / 5

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

<a id="canonical-2123211011202023-2320201030103113-0003323000121030-1021230331011033-2121003220132323-2020220122301303-0113320233321332-2231120302103232"></a>

## Next pages — clear_secret_info / 223212333223 / 6

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201002100220222-2012321132210020-0310231300233220-0223313323213302-3003313221300213-3331223211303203-0120322312131311-0031003200210233"></a>

## webhook.http_config.basic_auth — basic_auth / 200302012212 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.basic_auth

<a id="canonical-2132321110302210-1123002011303202-3231322022201031-2001311220022330-2201011313133011-0332223303203101-2031003113031220-3121202032322123"></a>

Type: `"object"`. single nested block, Optional.

Authorization parameters to access HTPP alert Receiver Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user_name")}
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
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003100112131001-2202031100321111-2221103220312201-3003023311320332-3112300021320220-0113311231102011-2321002110203231-0222210011120000"></a>

## Direct properties — basic_auth / 200302012212 / 3

- [password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233): complete subsection reference.

<a id="canonical-0312331210333010-3220133212302312-0213033321122133-3103211120003122-1320002300210120-3313132012210201-2301122222312001-1302030102021002"></a>

<a id="canonical-2013102003000212-0030202001210202-2313111133101001-0102031212023133-3203322120100003-1310020313330123-0131122120313200-3110311331101322"></a>

## user_name property — basic_auth / 200302012212 / 4

Type: `"string"`. Optional.

username. HTTP Basic Auth username.

Upstream description:

HTTP Basic Auth username.

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

<a id="canonical-3032130230013023-3131313123301022-2232010210122100-1212330102033331-2131032211331020-1333230012001222-0020012310323202-3220330311312103"></a>

## Next pages — basic_auth / 200302012212 / 5

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102111311321211-0133013123313210-0303101000222323-3033013120003221-3230231122212020-3022100033102121-1210332301233230-3310102031323113"></a>

## webhook.http_config.basic_auth.password — password / 100310023221 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233)
- webhook.http_config.basic_auth.password

<a id="canonical-1131033123221300-1211101023322300-3300232110223023-1212303331321131-3101023313222120-0230211121322012-3012022320321232-0100033210302203"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121131302101132-0331332123333220-0313102103321310-2203213033133113-0001033201202121-0002102311313103-3002010101131202-2231113102321202"></a>

## Direct properties — password / 100310023221 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0033011200113200-0110002003122001-3303022131132312-0021322211223031-1200103310302111-1323122301311013-2303013322321213-0033313302211020): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0232010013233031-2021210013313212-2002001021331113-3031213132230301-0022330230330000-0110313103310303-3132330312130203-3323322303012303): complete subsection reference.

<a id="canonical-3123020222023113-0110231212033032-3023100313103201-1132031300312121-2101331020231223-1201303322220032-1201232020030300-3113223112222032"></a>

## Next pages — password / 100310023221 / 4

- [webhook.http_config.basic_auth.password.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0033011200113200-0110002003122001-3303022131132312-0021322211223031-1200103310302111-1323122301311013-2303013322321213-0033313302211020)
- [webhook.http_config.basic_auth.password.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0232010013233031-2021210013313212-2002001021331113-3031213132230301-0022330230330000-0110313103310303-3132330312130203-3323322303012303)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0033011200113200-0110002003122001-3303022131132312-0021322211223031-1200103310302111-1323122301311013-2303013322321213-0033313302211020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012012330332302-3232323022331212-0330123320033302-2123301202210111-3021021221201211-0103223031332233-2302310300311332-3103010100001320"></a>

## webhook.http_config.basic_auth.password.blindfold_secret_info — blindfold_secret_info / 122130103300 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233)
- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233)
- webhook.http_config.basic_auth.password.blindfold_secret_info

<a id="canonical-3230330233113323-2020323331212310-2202103232112220-3201201203020333-1102102033311321-3100310230131112-0322211302023123-1323333312311122"></a>

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

<a id="canonical-3010303120202203-3303213213323203-1310321003301121-1302102311002132-2322323201020122-2321203220213133-2030031030031222-3132312203021300"></a>

## Direct properties — blindfold_secret_info / 122130103300 / 3

<a id="canonical-0002233002312011-1033312203200133-3030223212221023-0131031112231212-0220232132203033-3000102330223211-1013220302110002-3023000221321201"></a>

<a id="canonical-0211310033220131-3101031223230010-0130322111300023-1221332223111020-1230313330130312-0223122121111230-1231001312221023-1002201021012322"></a>

## decryption_provider property — blindfold_secret_info / 122130103300 / 4

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

<a id="canonical-2310201012131032-2200200132021000-0330213313311211-3011002211233213-2320313223030000-1021130001321232-3331320213113011-2200110232210222"></a>

<a id="canonical-2213330223123130-3230011200133321-3213222223131231-2012302033330001-3300003131301313-2110032223210103-1123113130130101-0131133321030033"></a>

## location property — blindfold_secret_info / 122130103300 / 5

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

<a id="canonical-0202212111230222-3132300332032130-2000122310321222-0112233103022030-3022102103231111-0311333032333011-0121332030212221-2202100232100132"></a>

<a id="canonical-2210113133223311-0303202120322213-2323203231313003-3302112022122001-1001011232110032-1201220211001200-0220033013002201-2323300221022102"></a>

## store_provider property — blindfold_secret_info / 122130103300 / 6

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

<a id="canonical-1110031133212202-3221312331331123-1211022231303132-0211130210303231-1020022332223223-3232022011232003-2033230122110313-3223203023133231"></a>

## Next pages — blindfold_secret_info / 122130103300 / 7

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0232010013233031-2021210013313212-2002001021331113-3031213132230301-0022330230330000-0110313103310303-3132330312130203-3323322303012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011211302330203-2023100311112203-0213213211333223-1313220122301110-3322121133121313-1111231332221322-1032222221112222-3232003111212113"></a>

## webhook.http_config.basic_auth.password.clear_secret_info — clear_secret_info / 122313021322 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233)
- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233)
- webhook.http_config.basic_auth.password.clear_secret_info

<a id="canonical-1233230332311121-2131122010012120-0000210321203211-1113220003111020-2111223300100030-0101100303211030-2222120132001202-0121013000031000"></a>

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

<a id="canonical-3233121010203012-0323220031323130-0021031130102220-1312302202133232-2302032000102101-3130230201223302-0120201022123002-0300121122100131"></a>

## Direct properties — clear_secret_info / 122313021322 / 3

<a id="canonical-0113322303011031-0211303011212212-0222003103303203-0231320030322203-2101223103013232-3022223112310310-3002013011211123-0322300031203230"></a>

<a id="canonical-2033210321033020-3011032303332222-3033200013330022-2221032030230330-0023221302023111-2200200031022112-1121312002032203-1121311013220202"></a>

## provider_ref property — clear_secret_info / 122313021322 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3010002132000302-3101012333331012-0123220022312311-1223103001022020-0230221011022010-3333313222330010-2332201111213103-0000300212320033"></a>

<a id="canonical-0131101132122212-3030211121202332-0221132231023333-2101332211100321-2221110211231121-2011100221202303-1330123330010132-2203333310013222"></a>

## URL property — clear_secret_info / 122313021322 / 5

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

<a id="canonical-3021030330122232-1310212102102230-1203302322331112-0310023122022000-0003100003230000-0323331302011201-2232223032203211-2032020222121211"></a>

## Next pages — clear_secret_info / 122313021322 / 6

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303333232002221-1001133223213033-0113022301222032-2133110002222331-3230010110132301-3302021021031000-2213200233232232-3001001122003203"></a>

## webhook.http_config.client_cert_obj — client_cert_obj / 023003012233 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.client_cert_obj

<a id="canonical-3113201222033133-1132310033201201-3021000213120200-3332320011221213-1212203322320012-1310023310112331-3020100022233300-1232102013330323"></a>

Type: `"object"`. single nested block, Optional.

Client Certificate Object. Configuration for client certificate.

Upstream description:

Configuration for client certificate.

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
client_cert_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123331113031303-2302232302020222-2023210312030322-2213300320023133-0020123022332130-0230023212100222-1221033102131300-2011232223203021"></a>

## Direct properties — client_cert_obj / 023003012233 / 3

- [use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-3313133313133231-3002013323102012-0102323202333012-2000222010000012-1123202321221331-2310320010313032-3320121212101102-3223023133132310): complete subsection reference.

<a id="canonical-1210301031132102-2100130013122332-1211331202110001-0300203200211013-0100313222121132-2210010011201133-2000301222323230-3330203210013222"></a>

## Next pages — client_cert_obj / 023003012233 / 4

- [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-3313133313133231-3002013323102012-0102323202333012-2000222010000012-1123202321221331-2310320010313032-3320121212101102-3223023133132310)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3313133313133231-3002013323102012-0102323202333012-2000222010000012-1123202321221331-2310320010313032-3320121212101102-3223023133132310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013033121111220-3113000220300002-1002201311220011-1333323122231302-3212121230300001-3002013213122013-2303301300111203-1230133012013211"></a>

## webhook.http_config.client_cert_obj.use_tls_obj — use_tls_obj / 013302313200 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100)
- webhook.http_config.client_cert_obj.use_tls_obj

<a id="canonical-1210212231300131-1131030033131310-3011213022311031-1202332010322120-0211210322200010-3300101303221110-3023021201101021-2011210203233301"></a>

Type: `"object"`. list nested block, Optional.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
use_tls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000332222002223-1203101311322200-1003322332232212-1331013230032213-1322001013000302-1313221320031002-3023201211022033-3101100103011211"></a>

## Direct properties — use_tls_obj / 013302313200 / 3

<a id="canonical-2223000133011203-0031301211133122-0302231223021322-0310222010120012-0313203300213202-3303100311113100-2203232200313001-1110213211111332"></a>

<a id="canonical-1103123321212103-1013033301020300-2001322212002323-0012002132021202-0123201213031202-1130000002222323-2313010103302322-2212303230130030"></a>

## kind property — use_tls_obj / 013302313200 / 4

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

<a id="canonical-0121321000021301-2022201010202213-0032112333223121-1100111111131012-0021311000232323-3101002212113002-2320300233221201-0223003210021313"></a>

<a id="canonical-2003331230301103-3001303203102101-2032011200103301-0131330203031123-0102300131131313-3313102022223003-0213003210321112-0121113321023010"></a>

## name property — use_tls_obj / 013302313200 / 5

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

<a id="canonical-1000131010012321-0020212230313211-1303123220330123-0022100123313232-1333202002220212-0023301301202202-2212232203131231-0133320300120331"></a>

<a id="canonical-2022301331101010-2021103311303221-0310011312020232-0312113303033333-3111112202202021-2122131021113021-3010100133130111-3011010332233320"></a>

## namespace property — use_tls_obj / 013302313200 / 6

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

<a id="canonical-0002130200133230-3101211102330321-2101023211002230-2133010210323112-1002311332011012-0313013212033210-2121330123230012-0122102023323330"></a>

<a id="canonical-3003313120110011-3022320013122310-3301022302031333-1023221003233210-1330323223112133-3303130133221300-3101212320000313-0202131132000020"></a>

## tenant property — use_tls_obj / 013302313200 / 7

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

<a id="canonical-0230231311113020-1233000310302012-1312320003131112-3201331102203123-2021330013322123-0130103002030121-2033223302032102-1003300023322302"></a>

<a id="canonical-3322111223302032-3012113312330231-3130012223123003-3022103023233112-3032112130310010-2133110113320100-3312030301310303-1223122302223221"></a>

## uid property — use_tls_obj / 013302313200 / 8

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

<a id="canonical-0130020010300301-1003231223132200-0323211322030202-1122020202232103-3030003321010032-2130303210223203-3021133120210302-3011200001221111"></a>

## Next pages — use_tls_obj / 013302313200 / 9

- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3001330103012311-2320121223330230-2011300030332132-3232110222030120-1130232011331212-1213030222332331-3231203321321130-2031122012102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000011113233323-2221333032101300-1302223132111302-1021313303221011-2332022130201330-0231221322133021-0221012220022100-2200011333133212"></a>

## webhook.http_config.no_authorization — no_authorization / 311110300211 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.no_authorization

<a id="canonical-3122301322312032-3333300213203211-2223212030332133-2202120200232110-3331112021302122-1001000000133323-0032101201133113-3320020030132331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authorization.

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
no_authorization = {}
```

<a id="canonical-1213100101303131-0203303203301123-3033103112111230-2301132230010011-1012111333221200-0020022220101322-2213321020303333-1001202121333311"></a>

## Direct properties — no_authorization / 311110300211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200210013200110-1131303212012212-0310020230033321-3222133230321122-1301012113023110-1313011023323320-3122310323130120-1220303202321310"></a>

## Next pages — no_authorization / 311110300211 / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3322110000220223-2002013220102223-2301310132030110-1131321030322130-3023232113100010-3013022200212212-0103033022220323-2003310000213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211121211122332-1110333010132233-3212120022331320-0323132210321031-3130121013022213-2332002203113333-2011310031131001-3322113122222230"></a>

## webhook.http_config.no_tls — no_tls / 233020123111 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.no_tls

<a id="canonical-3000033011022110-3033303331221103-0322320031203301-2130020022331303-3331222010203100-3201310220311120-1013132200232220-3122000020230003"></a>

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
no_tls = {}
```

<a id="canonical-1330210002021332-1122033330132121-3001103111013300-2320222200232211-0030323313332111-1032133332232121-2300230300023300-2100222112323233"></a>

## Direct properties — no_tls / 233020123111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211313121032120-1032111001133303-3230331122021132-1113331013131200-1301100302022131-2022113031203011-2003333300002030-1001320100010222"></a>

## Next pages — no_tls / 233020123111 / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330032010302321-2033202310312201-3021102111310031-2311102320130100-3313320030103203-3300011012121312-3110030130311101-0013201331031211"></a>

## webhook.http_config.use_tls — use_tls / 003130023023 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- webhook.http_config.use_tls

<a id="canonical-3203003230133131-2132222112332302-3001002000033331-1031112202031122-3033022113222201-0200013300301033-2210100111023112-2231311203123310"></a>

Type: `"object"`. single nested block, Optional.

Configures the token request's TLS settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000233010011033-2021222311301021-1310320323321130-1101203013113330-2123120011100332-2110212300022022-1232002232200000-1302010220001203"></a>

## Direct properties — use_tls / 003130023023 / 3

- [disable_sni](resources--alert_receiver--reference--group-001.md#canonical-0320030233030310-1013031020112330-1322120301100231-2033031312001131-2123300130000202-3123120132213321-0122310003223200-3102022222220203): complete subsection reference.

<a id="canonical-2132311132110310-2322132300003323-2012300213020112-0002331323231231-3300001302101011-0000021210130132-3102311113101003-1031002113210221"></a>

<a id="canonical-2120311103032233-1010003201231033-1320130323231023-3331321033200210-2131110003012200-1331212222302202-2213301113110133-2331231023212002"></a>

## max_version property — use_tls / 003130023023 / 4

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-1111111120220200-0112133231120332-1031330312022122-2303132010131321-3303032033001220-1133322212103122-0001222030302203-3233310300023031"></a>

<a id="canonical-3312213001201012-0101030022201231-1333031102121030-0221132031101031-3021203221323133-3012131021100320-0220331300333331-1220310002233231"></a>

## min_version property — use_tls / 003130023023 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-1103321111032013-2101322120221000-0301003122112232-2011320103033223-1222323213003233-1303320213131310-0323131011223212-3212131002201331"></a>

<a id="canonical-0032013021133103-2022013030002013-0203212220133232-3333212200231032-2211231123002211-1100223003022323-3113032110022303-3320023130102323"></a>

## sni property — use_tls / 003130023023 / 6

Type: `"string"`. Optional.

Exclusive with \[disable\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133): complete subsection reference.

- [volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-3231300230002132-1122230313302133-3111223223032033-1022113031333210-0310103333120312-2101221001200012-0331102202123022-3030211012211102): complete subsection reference.

<a id="canonical-0312111020200231-3002112221033000-0320100233212200-1023201111300231-1302313210113033-3333100000311032-1320312132200110-3131111232313123"></a>

## Next pages — use_tls / 003130023023 / 7

- [webhook.http_config.use_tls.disable_sni](resources--alert_receiver--reference--group-001.md#canonical-0320030233030310-1013031020112330-1322120301100231-2033031312001131-2123300130000202-3123120132213321-0122310003223200-3102022222220203)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133)
- [webhook.http_config.use_tls.volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-3231300230002132-1122230313302133-3111223223032033-1022113031333210-0310103333120312-2101221001200012-0331102202123022-3030211012211102)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0320030233030310-1013031020112330-1322120301100231-2033031312001131-2123300130000202-3123120132213321-0122310003223200-3102022222220203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030300122121123-0330213200212020-1031012110310210-0120011332123300-0321110231222122-3023030012313232-1122100011101301-0120011221310333"></a>

## webhook.http_config.use_tls.disable_sni — disable_sni / 302220310230 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- webhook.http_config.use_tls.disable_sni

<a id="canonical-2233220301111311-3001232323312030-3132210102232110-2320110222330110-1333122332110011-2100231033111100-2313223021310111-2212302200002311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-2123233110213101-2123233320023012-1202301033013022-3212021313121331-0030120322013201-3101220001122010-0313332211033332-3131032200323123"></a>

## Direct properties — disable_sni / 302220310230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101000210131022-2132323320210000-2132032222322202-1222211212021012-2203032123021311-0002211203120300-1102302203211000-2121020023131321"></a>

## Next pages — disable_sni / 302220310230 / 4

- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010210120322221-2131120023031121-0303123323301022-2222101003310102-2212332001330223-2002010232100102-2222200300210002-1110201123012103"></a>

## webhook.http_config.use_tls.use_server_verification — use_server_verification / 332331303132 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- webhook.http_config.use_tls.use_server_verification

<a id="canonical-1130011322013110-1310113121022120-3021113120331022-3132222200302332-3310103132331101-0323021100133232-1112333120203321-2031033213021231"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

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

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032231013110102-0001222103333332-0013001210131002-3231020003001101-2033301031002033-3130231312113021-1223202213233003-0001102101113033"></a>

## Direct properties — use_server_verification / 332331303132 / 3

- [ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-1003300023330312-1101230210120132-3233321202232310-3001331330131032-3323000222031023-3020031003212202-1022202301130111-0030112133330021): complete subsection reference.

<a id="canonical-1321320131023302-1200213112113022-2231311230200311-3102000212233310-1101000133002210-2200301210321001-1102230032132030-3322120100313132"></a>

## Next pages — use_server_verification / 332331303132 / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-1003300023330312-1101230210120132-3233321202232310-3001331330131032-3323000222031023-3020031003212202-1022202301130111-0030112133330021)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1003300023330312-1101230210120132-3233321202232310-3001331330131032-3323000222031023-3020031003212202-1022202301130111-0030112133330021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100131013001221-3201000303031310-3000112033200312-1023321130201222-2333001331112301-3100311001011022-3330311020220013-1010231110021001"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj — ca_cert_obj / 300302021211 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="canonical-2230233220111002-3110010312003033-3220312013033201-1200013233331230-2110230233110132-0310323011133002-3331333211322202-0001020111112222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ca cert obj.

Upstream description:

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

Terraform syntax:

```terraform
ca_cert_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313311213110220-0300213311300301-0101203111110121-1030231020300012-1200031132030231-2213310003121311-3110012132213232-2100020021301133"></a>

## Direct properties — ca_cert_obj / 300302021211 / 3

- [trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-1331201222320013-0332232332222031-2122022122230201-0300201112310003-3313102001021303-3021222223030201-1211211102100201-2000120311003122): complete subsection reference.

<a id="canonical-1000200203102000-3100013211230011-2103033322230101-3012200210201333-0311332003002321-0323113300001130-2130202000130023-0322331011110020"></a>

## Next pages — ca_cert_obj / 300302021211 / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-1331201222320013-0332232332222031-2122022122230201-0300201112310003-3313102001021303-3021222223030201-1211211102100201-2000120311003122)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-1331201222320013-0332232332222031-2122022122230201-0300201112310003-3313102001021303-3021222223030201-1211211102100201-2000120311003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020313231233122-1032112122203121-1213102232311013-0333220023010021-2012333101103332-2213223311322023-1011022033202320-3033213213013311"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca — trusted_ca / 212132200211 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-2131220010301300-3123321232201323-2213202011132023-1013020222133000-1021033331010320-2233102323120313-1301211312133201-3212230200301133)
- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-1003300023330312-1101230210120132-3233321202232310-3001331330131032-3323000222031023-3020031003212202-1022202301130111-0030112133330021)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca

<a id="canonical-1111020022202313-1203220000100031-2321212023122331-0230003311133320-3033300032220211-0312222130111311-3011023303221331-2022303030012001"></a>

Type: `"object"`. list nested block, Optional.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213033121332200-0222111312033120-3331300200303223-3223021110301331-3031311130020000-1033013301230232-3102212111302130-1321012031232122"></a>

## Direct properties — trusted_ca / 212132200211 / 3

<a id="canonical-2213033213222313-2321020113320011-1313200312302122-0302031110012231-0120303023212002-1222031112203320-2101333121202220-3132130103120213"></a>

<a id="canonical-3223210303313020-2300330002121201-2021320031121013-1322220110231323-3232332021211303-1113311111102330-0203021303110110-2310231233120001"></a>

## kind property — trusted_ca / 212132200211 / 4

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

<a id="canonical-1220111222223030-3233012313233102-2333122301023131-2322323102323330-0120203111202030-3230010303222213-1001223120211112-2132213200303023"></a>

<a id="canonical-1310211213022212-0221312202020310-0013121020002103-2003202222201002-0232111231203203-0311200220113302-0311330032002202-3201203222333313"></a>

## name property — trusted_ca / 212132200211 / 5

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

<a id="canonical-2333223211132103-1322011223100320-2103101002301120-2101233320013132-1122201210020221-0133322312300200-2010223120300201-0331010111013233"></a>

<a id="canonical-0133031303233312-1002113330300332-1322131002111101-0023213210102213-0133113303200333-2331300310202312-3013220111300123-3131230231102022"></a>

## namespace property — trusted_ca / 212132200211 / 6

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

<a id="canonical-1102322330022330-0301022310130002-1312320113201131-2322021310112210-2200202202323330-3103112122003230-1333102331103133-2323210231232321"></a>

<a id="canonical-3233223000310100-2032003122032032-1120201201302110-0323022002023003-1332002320323113-1120030011102120-3031033000131132-3310003030131012"></a>

## tenant property — trusted_ca / 212132200211 / 7

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

<a id="canonical-2133231320020020-1313202033000122-3023100301133311-1320302113333101-1202013131313331-3323100001021223-1333022011002023-3232222300232111"></a>

<a id="canonical-3323122321202021-1220221200220311-0310011132122330-3000332303103312-3213033013032332-1323310210101031-1130221022322103-0312131201223222"></a>

## uid property — trusted_ca / 212132200211 / 8

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

<a id="canonical-3201330002120210-0322330200230031-0020310313301201-1100132103200222-3231233100023322-2321232033013220-0313300122111332-2001000201131110"></a>

## Next pages — trusted_ca / 212132200211 / 9

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-1003300023330312-1101230210120132-3233321202232310-3001331330131032-3323000222031023-3020031003212202-1022202301130111-0030112133330021)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-3231300230002132-1122230313302133-3111223223032033-1022113031333210-0310103333120312-2101221001200012-0331102202123022-3030211012211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312023110211210-2213133031322101-0130003120120331-3110302110110003-3313322302311002-1111100032001220-2310200230211212-0112230001322011"></a>

## webhook.http_config.use_tls.volterra_trusted_ca — volterra_trusted_ca / 000322020112 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="canonical-3210212011221031-3123201111011110-2323320112202230-3232101330200020-0313030200020301-0020011030133230-2232220132323011-1210030233210102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-3302000110012323-2133302023323230-1101021300033032-1130020013002221-1102322032130023-2203321020211132-1200201132323021-3321312331213212"></a>

## Direct properties — volterra_trusted_ca / 000322020112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231112123203301-1023120232302312-3322021321130111-2230101023223011-2001211013201132-3323212020132113-2113000213010102-2131032330103222"></a>

## Next pages — volterra_trusted_ca / 000322020112 / 4

- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322011301030011-1233212330100300-0233022203231112-2012012232010013-1211030232103110-1021113121031303-1321230023303220-3310012301331001"></a>

## webhook.URL — URL / 302120311100 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- webhook.URL

<a id="canonical-0200012202322320-1200012331001222-1102102031332331-3201002201303322-1022212222103122-1011232133101013-3021133201101032-2312022331230303"></a>

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
url {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213031133200021-3030033011311012-2011222331323010-1220023323301330-0000121101000120-0313222331323311-3100202212132121-1201231202321110"></a>

## Direct properties — URL / 302120311100 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0321102330022111-2303202000212013-1132311133333221-0000022112122022-0110110231011302-0112131101131012-0032130320322021-1003021112033003): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2020321110121221-3013300001320101-0323112233200321-3300220232300012-2213330010233320-1202322200102202-2212132013111121-0030113233021210): complete subsection reference.

<a id="canonical-0201012001032301-3022133102333201-1200211330320321-3133033331202021-0322320322312100-2020012021320223-1011221011001322-3211300320330212"></a>

## Next pages — URL / 302120311100 / 4

- [webhook.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0321102330022111-2303202000212013-1132311133333221-0000022112122022-0110110231011302-0112131101131012-0032130320322021-1003021112033003)
- [webhook.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2020321110121221-3013300001320101-0323112233200321-3300220232300012-2213330010233320-1202322200102202-2212132013111121-0030113233021210)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-0321102330022111-2303202000212013-1132311133333221-0000022112122022-0110110231011302-0112131101131012-0032130320322021-1003021112033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120003121033211-1123113132212030-2102120202313003-0233130100330123-1233200320112130-0033110021113130-3111212311123301-2122303301221220"></a>

## webhook.URL.blindfold_secret_info — blindfold_secret_info / 203122001010 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110)
- webhook.URL.blindfold_secret_info

<a id="canonical-0211232023101031-3102231010220121-1020020112322320-3332300311132022-3122223121023301-1332002033210130-2132000020031313-2131012103202301"></a>

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

<a id="canonical-0213120310333120-1300120020002122-2200013110122232-2223121123010110-0303311230201310-0110231131010011-3001312011112300-2020011320223223"></a>

## Direct properties — blindfold_secret_info / 203122001010 / 3

<a id="canonical-2102022121212202-0133130203201103-3211120321033102-3322111031220313-2222212131013120-0120230133230202-3212033010312033-0303131310301220"></a>

<a id="canonical-3220113330030030-2231302302012320-0321300131122131-1112323223131131-0031003220123231-2333121312230130-3111203312110113-2201112233033000"></a>

## decryption_provider property — blindfold_secret_info / 203122001010 / 4

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

<a id="canonical-2310122212201313-3132322132012323-0211302131233323-3120123320123023-2222121200231130-0031003211221020-2221113212012232-0002121020133221"></a>

<a id="canonical-2322333132103120-1331201122030110-0332002013330120-0320122323021001-3111030232122101-1201301331001000-1111020013102332-0123333221220001"></a>

## location property — blindfold_secret_info / 203122001010 / 5

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

<a id="canonical-3210310302302310-3011222033130203-3333131202233232-3032130313213032-1303102333321101-3121330022033312-2100001131102002-0032310210212323"></a>

<a id="canonical-3011322312311321-2202323223030300-3330031111122033-2323001132130320-1111220331303002-1131310113312111-1132200220113002-0103133121110330"></a>

## store_provider property — blindfold_secret_info / 203122001010 / 6

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

<a id="canonical-1321032212031232-1030031020030211-0011312331221313-1120313020003133-3013321003321020-3132222030113122-3312233212302200-1202323132212231"></a>

## Next pages — blindfold_secret_info / 203122001010 / 7

- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)

<a id="canonical-2020321110121221-3013300001320101-0323112233200321-3300220232300012-2213330010233320-1202322200102202-2212132013111121-0030113233021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111113302313233-0023101112122210-2221010000303101-0320212230013300-2200333203021310-1000121232310120-3302233032010201-0310003322101121"></a>

## webhook.URL.clear_secret_info — clear_secret_info / 200210221110 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-0203003021003112-2331200103023302-0122130211113012-1101331111211212-3333322200311030-1220011302033023-0300320331122021-0001122031122103)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110)
- webhook.URL.clear_secret_info

<a id="canonical-0321233303002210-0221201310111001-3120312013002332-3112322031123212-2233220312202211-2201003002022221-3331211311211031-0013321223220120"></a>

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

<a id="canonical-0001312020023233-2102202231022311-1232211132201202-2302100200313023-2011303012323123-3233223213123013-3033003333313212-0230222303203032"></a>

## Direct properties — clear_secret_info / 200210221110 / 3

<a id="canonical-1212002112122303-3132030121012203-0131112131233303-0232223010132301-3332310133012120-2222302311230330-0230002020210332-2333332311201003"></a>

<a id="canonical-2010132130212000-2311133032122022-3130310100010023-1010301200011130-1333221133000013-1131023012113120-3201332010110000-3200221323032013"></a>

## provider_ref property — clear_secret_info / 200210221110 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232311102020321-2311203012102211-3001130213102003-0113202021233303-3012100013032102-0102031303203013-3231320222133002-2323111013222223"></a>

<a id="canonical-3110310111032202-1312100231113102-1313022131112331-0302322311022013-2102312312023212-1033223010202101-3013002321223131-2330103212121212"></a>

## URL property — clear_secret_info / 200210221110 / 5

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

<a id="canonical-3033303130112103-0003113102320203-3301113212301030-1002102322301013-3200101331130330-1313121232020213-3113123213301102-0100101132311001"></a>

## Next pages — clear_secret_info / 200210221110 / 6

- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-2033110323033103-3132333200000003-0120310200220223-3323203323303332-3101321030211112-3121123120032101-0322231230021112-2031012010103110)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
