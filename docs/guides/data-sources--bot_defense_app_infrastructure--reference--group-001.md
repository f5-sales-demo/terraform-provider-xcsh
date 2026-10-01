---
page_title: "xcsh_bot_defense_app_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure reference."
---

# xcsh_bot_defense_app_infrastructure reference

<a id="canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323220130210013-2103101123033020-2131231222223201-2123122231231111-3003301012002302-0321302030031333-2202131311012132-3021221212032221"></a>

## Property reference — Property reference / 112000223202 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- Property reference

<a id="canonical-2312123301121212-3032313331333312-2021201112300200-1301212321112101-0023131332302221-0232303021100320-1010012101110221-3010032233031013"></a>

## Direct properties — Property reference / 112000223202 / 3

<a id="canonical-3333310213020102-2200212102110302-0221011012332123-1233031120101310-2113230012312322-1022031030100132-0103330011030322-1120211302013121"></a>

<a id="canonical-2301200212133011-1302301031211310-2123211222011032-1100011212201133-2131100322010313-1212003031310113-0312032122011111-0212030210223233"></a>

## annotations property — Property reference / 112000223202 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103): complete subsection reference.

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001): complete subsection reference.

<a id="canonical-2331200232311030-0110130001031232-1303201023000013-3131113012212313-0202200333111112-0333202111332300-3320221033012200-1100023013311112"></a>

<a id="canonical-3211310122102123-0211033111103303-3021231221012312-1330202110203131-2233330313212333-3331123201233023-1211213011102113-0222202302012003"></a>

## description property — Property reference / 112000223202 / 5

Type: `"string"`. Computed.

Description of the BotDefenseAppInfrastructure.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1002023303130003-0012003332213310-1001302032123031-3122320331211302-1212313031331113-0233211220303222-1000232222002031-1102320302030323"></a>

<a id="canonical-3133233000312001-2221302220033101-2320310033202220-2011002223202333-2112101013232111-2121113231323010-2113210021320102-1203032013030203"></a>

## environment_type property — Property reference / 112000223202 / 6

Type: `"string"`. Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Receipt-pinned upstream constraints:

```json
{
  "default": "PRODUCTION",
  "enum": [
    "PRODUCTION",
    "TESTING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020103023001310-1132033210312000-0333030110300312-2203121103330003-2103103301230221-2132200002200331-3002220320002323-2320030203201012"></a>

<a id="canonical-1323112012023123-2032222312330033-1203333013001220-0230110312322320-2020323321100301-1311010213233131-3303321333100121-1130101112333133"></a>

## ID property — Property reference / 112000223202 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1102010012203032-0210103221220120-0013331203321021-0033112302220011-2220331120102013-2010213132202213-1032221033201203-3322322133031132"></a>

<a id="canonical-3023012120211300-0302103102101103-3132002221231120-2000031311213001-0213120123221132-2030302231131131-3001010010211302-3330002110101222"></a>

## labels property — Property reference / 112000223202 / 8

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

<a id="canonical-1010232211120131-1011231123000111-3133033002131331-2131333211123211-1010010230212110-2220113113231030-1203330133123012-2300031331211213"></a>

<a id="canonical-0201120132222223-2231301323322330-2211133000131130-2201000330010000-0030322033013031-0002131211311011-0221222313020221-2102211012231333"></a>

## name property — Property reference / 112000223202 / 9

Type: `"string"`. Required.

Name of the BotDefenseAppInfrastructure.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3333132213011333-3310232031010231-3332232101102103-1201333301221221-3133002223131313-1101303131300210-0123133312003120-0333003110120301"></a>

<a id="canonical-0200303020320001-0030213111203213-3233300102023123-0020210101010311-3201030223313033-0023322211002010-3233301130013023-1322230121023120"></a>

## namespace property — Property reference / 112000223202 / 10

Type: `"string"`. Required.

Namespace where the BotDefenseAppInfrastructure exists.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2031133110331332-3031223333030013-3223131203001203-1113001230213101-0120212223220220-3301021113123121-1011312202133313-3022112011230020"></a>

<a id="canonical-2032203331200320-3100113332020312-1100200211133332-2003011301113012-1021320221203023-0130122233101023-0312302221113123-3122132000232101"></a>

## traffic_type property — Property reference / 112000223202 / 11

Type: `"string"`. Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100202001231100-1103032113000013-0320302023302020-0023012121233033-3001333232131312-2223310123222320-2302111211200103-1310220021131120"></a>

## All schema paths — Property reference / 112000223202 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3333310213020102-2200212102110302-0221011012332123-1233031120101310-2113230012312322-1022031030100132-0103330011030322-1120211302013121) |
| `cloud_hosted` | [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3033221101121322-2311011112133011-1002103332321123-3122011302333313-2201220301123311-0100310023210313-2031010221331112-3123313233233000) |
| `cloud_hosted.egress` | [cloud_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2330313331330033-3102333011120002-1332203222321032-0102321111100312-0012303132102223-3231302101132320-2033012032231303-2032220221012103) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0023303111201310-1000301211213013-2123312003213230-3021113003321330-2031033000311123-1223122321333223-0203023202113313-0123222313323022) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2120303202013113-2023310233123031-1122120121013221-3302310230021312-2000103103233123-1111131022033333-3221302213121303-0322312120132123) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2322123102033122-2320012120203221-2232322200210131-2013131002332123-0120132112231102-1013110113231303-3230331322301112-2021020230001303) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2012132032332012-2332303132123200-0121300132121223-2103130132102013-2231231331230333-3221303203210331-1012003122003230-1223020001222112) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1310111103113221-2102022200211302-3023023322222132-0132131202013320-3100103012123332-2231031001032212-2130201001033013-2213100110210233) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0200113333310313-3130201110123121-0103013020210220-0132033000030213-0031232020232132-1121020122102111-2200032030320033-2020201321232000) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1032310231133233-0320102213211131-1000000322233313-2303032231202110-0103223033201101-1303320021000222-0102232000030103-1100013101023330) |
| `cloud_hosted.region` | [cloud_hosted.region](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1232221232313133-3031302221303302-1232223131211323-1311123322322312-1112000331303210-1030010201010110-3222222202121333-1230022100220033) |
| `data_center_hosted` | [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0110112320212020-0313230003020320-3022313223012000-0230120330230313-3331333310012011-3100010121231231-2323330303230321-2331011232113100) |
| `data_center_hosted.egress` | [data_center_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1210310323023012-1221023332323010-3030200222103132-2212021010102123-1002210100210322-2103321223333100-1000030323321310-1310221322312102) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3222021030112303-2103213130233003-0000003222120231-0323123230311003-2133233032203202-2030300100303021-0203131221021033-2000213203202003) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2023313310121300-3012333120223110-3231130313320122-0010313213311213-2000130200030223-3203302231301220-1032331011210111-1123301133211002) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2333220302301210-3212203030231333-3010323131313120-1230231012322110-3013320303333133-0312122121313332-0120131223220121-1010330201202110) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3011112302210210-2012133032211232-0332003302212323-0133323130103110-1020320231310121-1200031032211010-1212323213001210-2000023312320001) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0102210212212200-3200232330000033-3122302301313303-3032231303212322-3003212232301023-1130131010022222-2113203020103220-2320323123121332) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0233022233332310-3320232130300313-2230110133003132-1310322013123021-2323031213030130-0011030302213321-3022203202103033-2231103202021003) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0033322133323333-0222331102233122-0333130333120331-2120102113131213-3321231000303003-2033010232100200-3322121311030230-3331312203311113) |
| `data_center_hosted.region` | [data_center_hosted.region](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2323331021233310-0122200320031311-3111221111311133-2213230303113100-0100322201123023-3030301222330003-1302220123010201-1221202220111010) |
| `description` | [description](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2331200232311030-0110130001031232-1303201023000013-3131113012212313-0202200333111112-0333202111332300-3320221033012200-1100023013311112) |
| `environment_type` | [environment_type](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1002023303130003-0012003332213310-1001302032123031-3122320331211302-1212313031331113-0233211220303222-1000232222002031-1102320302030323) |
| `id` | [ID](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0020103023001310-1132033210312000-0333030110300312-2203121103330003-2103103301230221-2132200002200331-3002220320002323-2320030203201012) |
| `labels` | [labels](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1102010012203032-0210103221220120-0013331203321021-0033112302220011-2220331120102013-2010213132202213-1032221033201203-3322322133031132) |
| `name` | [name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1010232211120131-1011231123000111-3133033002131331-2131333211123211-1010010230212110-2220113113231030-1203330133123012-2300031331211213) |
| `namespace` | [namespace](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3333132213011333-3310232031010231-3332232101102103-1201333301221221-3133002223131313-1101303131300210-0123133312003120-0333003110120301) |
| `traffic_type` | [traffic_type](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2031133110331332-3031223333030013-3223131203001203-1113001230213101-0120212223220220-3301021113123121-1011312202133313-3022112011230020) |

<a id="canonical-1130220312013113-0223330330001132-3230122112223301-0013203301200322-2313223222332113-3302210001012002-3200212013123233-0110310120323230"></a>

## Next pages — Property reference / 112000223202 / 13

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132300202221321-0203023022132012-1323201021300103-1321123120320130-2120021232220213-0302302222103332-0232213020212312-1032233311300322"></a>

## cloud_hosted — cloud_hosted / 121102112022 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- cloud_hosted

<a id="canonical-3033221101121322-2311011112133011-1002103332321123-3122011302333313-2201220301123311-0100310023210313-2031010221331112-3123313233233000"></a>

Type: `"single"`. Computed.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

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

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3033221101121322-2311011112133011-1002103332321123-3122011302333313-2201220301123311-0100310023210313-2031010221331112-3123313233233000)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0110112320212020-0313230003020320-3022313223012000-0230120330230313-3331333310012011-3100010121231231-2323330303230321-2331011232113100)

Select alternatives according to the provider validators above.

<a id="canonical-0220011211310313-2213201302212020-3130003132103102-3003033302121000-3212300221201220-0132020100113213-1233212131001320-3220010212221310"></a>

## Direct properties — cloud_hosted / 121102112022 / 3

- [egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0023312311332123-0221102122213030-3200211213031023-1103111110301130-1132001223310022-0300322013110331-1221003022130001-1033003112233011): complete subsection reference.

<a id="canonical-2322123102033122-2320012120203221-2232322200210131-2013131002332123-0120132112231102-1013110113231303-3230331322301112-2021020230001303"></a>

<a id="canonical-1123330100133010-2311311312013120-2300103013130303-2123120122311201-1022011032013213-3211230023010231-2111022110331123-2332212222233130"></a>

## infra_host_name property — cloud_hosted / 121102112022 / 4

Type: `"string"`. Computed.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3212133321032131-1102223321001133-3211232110000131-0020130030120300-0130101311022012-1203102203102102-1121123310203211-2001001103300201): complete subsection reference.

<a id="canonical-1232221232313133-3031302221303302-1232223131211323-1311123322322312-1112000331303210-1030010201010110-3222222202121333-1230022100220033"></a>

<a id="canonical-3003002332102300-3031000312001320-3222011230030120-2220332230101121-2001202000212303-1032032112232033-1113301111002100-0003300313220031"></a>

## region property — cloud_hosted / 121102112022 / 5

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331311302232020-0231223010323231-0303320320112000-3232020101212223-2132113133020202-3032132103221310-2031222232131010-1032202101333220"></a>

## Next pages — cloud_hosted / 121102112022 / 6

- [cloud_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0023312311332123-0221102122213030-3200211213031023-1103111110301130-1132001223310022-0300322013110331-1221003022130001-1033003112233011)
- [cloud_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3212133321032131-1102223321001133-3211232110000131-0020130030120300-0130101311022012-1203102203102102-1121123310203211-2001001103300201)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-0023312311332123-0221102122213030-3200211213031023-1103111110301130-1132001223310022-0300322013110331-1221003022130001-1033003112233011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030323301102100-2110201320121003-1320310003330313-1312223322012221-2233303310102203-2011303203132223-0102332332021131-0222310001023302"></a>

## cloud_hosted.egress — egress / 123231011113 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103)
- cloud_hosted.egress

<a id="canonical-2330313331330033-3102333011120002-1332203222321032-0102321111100312-0012303132102223-3231302101132320-2033012032231303-2032220221012103"></a>

Type: `"list"`. Computed.

Egress. Egress

Upstream description:

Egress

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2103033213332203-1322212013222100-1021312200130112-1022130201112301-0211113332023120-3012203112112031-3210002011221330-1011122110302223"></a>

## Direct properties — egress / 123231011113 / 3

<a id="canonical-0023303111201310-1000301211213013-2123312003213230-3021113003321330-2031033000311123-1223122321333223-0203023202113313-0123222313323022"></a>

<a id="canonical-0230013011120331-1211000111130320-1323132031130313-0231300210333020-2032101233211301-3323221003000131-3002033210200321-3211213122211011"></a>

## ip_address property — egress / 123231011113 / 4

Type: `"string"`. Computed.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2120303202013113-2023310233123031-1122120121013221-3302310230021312-2000103103233123-1111131022033333-3221302213121303-0322312120132123"></a>

<a id="canonical-0131213110222303-1310100103303110-3203121222223030-2121020003122033-0230232130031013-3033133233332312-1101331211031101-0212200320011200"></a>

## location property — egress / 123231011113 / 5

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3222303330210112-2331022101113030-0203020020200102-3321200120313032-2310122130320330-2012230003301211-0303211001011130-2300221223132013"></a>

## Next pages — egress / 123231011113 / 6

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-3212133321032131-1102223321001133-3211232110000131-0020130030120300-0130101311022012-1203102203102102-1121123310203211-2001001103300201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121102301012223-2032013120110300-2303221011230012-0111323221202121-3203001321120123-2210021221033011-0102222001300003-2123223212113220"></a>

## cloud_hosted.ingress — ingress / 212210223230 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103)
- cloud_hosted.ingress

<a id="canonical-2012132032332012-2332303132123200-0121300132121223-2103130132102013-2231231331230333-3221303203210331-1012003122003230-1223020001222112"></a>

Type: `"list"`. Computed.

Ingress. Ingress

Upstream description:

Ingress

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301022210212121-3111120202101330-0131110230132000-1020010000230323-0110010010111120-0223331230020212-1032213023100321-3100310323211101"></a>

## Direct properties — ingress / 212210223230 / 3

<a id="canonical-1310111103113221-2102022200211302-3023023322222132-0132131202013320-3100103012123332-2231031001032212-2130201001033013-2213100110210233"></a>

<a id="canonical-3313131212333023-0100030023133213-2203211033022032-0322010021131312-3102132001131021-0223322131113203-2322233123313212-0121001131312223"></a>

## host_name property — ingress / 212210223230 / 4

Type: `"string"`. Computed.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0200113333310313-3130201110123121-0103013020210220-0132033000030213-0031232020232132-1121020122102111-2200032030320033-2020201321232000"></a>

<a id="canonical-2110120012021230-0221312131303010-3321201013212130-3031110330022010-3223221323123010-1033123002221120-0312210323303233-0222002021322131"></a>

## ip_address property — ingress / 212210223230 / 5

Type: `"string"`. Computed.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1032310231133233-0320102213211131-1000000322233313-2303032231202110-0103223033201101-1303320021000222-0102232000030103-1100013101023330"></a>

<a id="canonical-2113221322333210-2011110322331132-3220123131101332-3101133012022100-3033212323001012-3320100311132221-1220220003003232-1123032201123232"></a>

## location property — ingress / 212210223230 / 6

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031020002133320-2233223233123201-3300031003013133-0002301231000000-0313322103210312-0022103131100201-0012311131233222-0323032033222232"></a>

## Next pages — ingress / 212210223230 / 7

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100201031320101-2201121012220030-3220103233112010-2222021333031100-2232000232311221-2132131210223023-2021222011121131-1112310113101220"></a>

## data_center_hosted — data_center_hosted / 000303302003 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- data_center_hosted

<a id="canonical-0110112320212020-0313230003020320-3022313223012000-0230120330230313-3331333310012011-3100010121231231-2323330303230321-2331011232113100"></a>

Type: `"single"`. Computed.

F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

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

<a id="canonical-2130002023301122-3220222031133200-2001230022302103-2103311300111132-1201113210001131-0123121010130221-2223232202230012-0203333023121321"></a>

## Direct properties — data_center_hosted / 000303302003 / 3

- [egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2011223320333230-0030030221131322-1221130213210311-2320221332111331-1021012312031222-0210333023301103-1221331222011002-1020033031010000): complete subsection reference.

<a id="canonical-2333220302301210-3212203030231333-3010323131313120-1230231012322110-3013320303333133-0312122121313332-0120131223220121-1010330201202110"></a>

<a id="canonical-3321300321100331-2331023203333230-1302010100333003-3113201211130110-2000112302300000-1000211332232021-0102323110312300-1232013012322021"></a>

## infra_host_name property — data_center_hosted / 000303302003 / 4

Type: `"string"`. Computed.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2332132211113123-0320021203313232-3302130030312112-3002010012133101-1132123033332312-1022033303130002-2000100333021131-2121111210222313): complete subsection reference.

<a id="canonical-2323331021233310-0122200320031311-3111221111311133-2213230303113100-0100322201123023-3030301222330003-1302220123010201-1221202220111010"></a>

<a id="canonical-2331111321120111-0310011003031311-3200111103313022-2232120212131211-0002103232031032-2331122231302313-0322020221220011-2010332101021100"></a>

## region property — data_center_hosted / 000303302003 / 5

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1131221212100321-0220232323320301-1332112303220111-3210232320003013-1022200102202202-2032330101331101-2102212213221111-2120131211333131"></a>

## Next pages — data_center_hosted / 000303302003 / 6

- [data_center_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2011223320333230-0030030221131322-1221130213210311-2320221332111331-1021012312031222-0210333023301103-1221331222011002-1020033031010000)
- [data_center_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2332132211113123-0320021203313232-3302130030312112-3002010012133101-1132123033332312-1022033303130002-2000100333021131-2121111210222313)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-2011223320333230-0030030221131322-1221130213210311-2320221332111331-1021012312031222-0210333023301103-1221331222011002-1020033031010000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221230221002222-1232331221112210-3313021123203003-0231133103000230-1221131302020011-3301022322312331-2100332212333211-2032123231030202"></a>

## data_center_hosted.egress — egress / 202013013303 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001)
- data_center_hosted.egress

<a id="canonical-1210310323023012-1221023332323010-3030200222103132-2212021010102123-1002210100210322-2103321223333100-1000030323321310-1310221322312102"></a>

Type: `"list"`. Computed.

Egress. Egress

Upstream description:

Egress

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0303102212031002-3030120212001321-1331023332102330-2121201201010202-3221123320112212-1131102102121231-3100033311311113-2020220330102303"></a>

## Direct properties — egress / 202013013303 / 3

<a id="canonical-3222021030112303-2103213130233003-0000003222120231-0323123230311003-2133233032203202-2030300100303021-0203131221021033-2000213203202003"></a>

<a id="canonical-3310113021313012-0312310101130021-3220302020233010-3032321023000102-3231010303000102-0223311121011200-3022311302012021-1200212100322221"></a>

## ip_address property — egress / 202013013303 / 4

Type: `"string"`. Computed.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2023313310121300-3012333120223110-3231130313320122-0010313213311213-2000130200030223-3203302231301220-1032331011210111-1123301133211002"></a>

<a id="canonical-0130121222330232-2030032303303311-0123132113130210-2322330322030210-1131202312110301-3322203130310202-2212133331033101-3300202303113211"></a>

## location property — egress / 202013013303 / 5

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2010100330300102-1302230030010321-2100011300323001-3301131311020122-2213200113111010-2211302233210113-2323210312303010-2001111311100312"></a>

## Next pages — egress / 202013013303 / 6

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-2332132211113123-0320021203313232-3302130030312112-3002010012133101-1132123033332312-1022033303130002-2000100333021131-2121111210222313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101310211100132-3302313011003203-0103310330003310-3210001003330101-1000122122333021-1003113303003120-0333031113022023-0201212301120320"></a>

## data_center_hosted.ingress — ingress / 323103222133 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001)
- data_center_hosted.ingress

<a id="canonical-3011112302210210-2012133032211232-0332003302212323-0133323130103110-1020320231310121-1200031032211010-1212323213001210-2000023312320001"></a>

Type: `"list"`. Computed.

Ingress. Ingress

Upstream description:

Ingress

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022010310211110-2320303031123110-2213131023320331-2000102113011100-3220101300200113-3002222110023010-0222012020222022-1102121112021133"></a>

## Direct properties — ingress / 323103222133 / 3

<a id="canonical-0102210212212200-3200232330000033-3122302301313303-3032231303212322-3003212232301023-1130131010022222-2113203020103220-2320323123121332"></a>

<a id="canonical-0030213020020211-1211330301103303-2002103231311123-1300321322003323-1332021102203011-3001320220321313-1113332331323211-1230030123112330"></a>

## host_name property — ingress / 323103222133 / 4

Type: `"string"`. Computed.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0233022233332310-3320232130300313-2230110133003132-1310322013123021-2323031213030130-0011030302213321-3022203202103033-2231103202021003"></a>

<a id="canonical-3311111200100320-3003030221113111-2123201130331233-2111332212022101-2022120120021103-3002110111110103-0230312300012203-3201021123201313"></a>

## ip_address property — ingress / 323103222133 / 5

Type: `"string"`. Computed.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0033322133323333-0222331102233122-0333130333120331-2120102113131213-3321231000303003-2033010232100200-3322121311030230-3331312203311113"></a>

<a id="canonical-1120111120221010-2313033220312112-2031011123232321-0221303010131020-0123332223011130-1302002203321210-0130322022330113-1122201100220233"></a>

## location property — ingress / 323103222133 / 6

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0200330231312033-3231220311011020-0201222010202312-2321231132220132-2120233101133321-3032122111231103-3110022110113121-2332133103133112"></a>

## Next pages — ingress / 323103222133 / 7

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
