---
page_title: "xcsh_bot_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure reference."
---

# xcsh_bot_infrastructure reference

<a id="canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111233101203003-3230301320110322-2033130312301131-1301102031202333-0330213101313201-3230021300303112-0122220012332310-0331003301101023"></a>

## Property reference — Property reference / 003200213331 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- Property reference

<a id="canonical-0100130112110130-0121320332210102-0312211231113022-0000233003103011-0233030013003220-0003112333120331-2232311113300310-2333223020111303"></a>

## Direct properties — Property reference / 003200213331 / 3

<a id="canonical-0031201330213311-2111013301332330-2000222302202330-1020131203220213-3131312301223321-2000301313220022-2331312233133033-3213331121103012"></a>

<a id="canonical-0130313101222100-1323300222323220-0120201213030001-3321033200121231-2002122230001021-1211113310110010-0323330003301202-3010111023112003"></a>

## annotations property — Property reference / 003200213331 / 4

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

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121): complete subsection reference.

<a id="canonical-1202313331002232-1113112120210022-3223000320102200-1120331101010100-3321202121013003-2323123013102133-2231122333231321-2112232011032003"></a>

<a id="canonical-1133331233203312-0313332010011013-0000222303133313-0213332021000321-2321332300020200-2302223012123011-2131122031322210-3333230101310122"></a>

## description property — Property reference / 003200213331 / 5

Type: `"string"`. Computed.

Description of the BotInfrastructure.

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

<a id="canonical-2323300122311222-1122132113202301-1211221123031210-1023330320133311-3120330331023031-3100023311123010-0102001312333330-2313013103213103"></a>

<a id="canonical-3112323001032020-1021322320011312-0322112020010010-1021230110132320-0013023333203200-2011331323102200-1031003132310103-0103221022213133"></a>

## ID property — Property reference / 003200213331 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0102103023233312-3322222032023103-3223010220111103-2321210123333233-1312030123123021-2202322023233113-0022321102030020-0010301311221010"></a>

<a id="canonical-2233223013112120-3013001011200233-2132220112201101-0000301223130222-3320331232033103-3032300210223232-0002221002230200-3003200032323230"></a>

## labels property — Property reference / 003200213331 / 7

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

<a id="canonical-2003220112120023-0212101022000001-2000113012012132-3202133132222013-0013223013330133-0212020011303102-3212113133010022-0113310300023003"></a>

<a id="canonical-0110310211133331-2130032110302212-1320211232220323-0012113333132113-1130202201010002-2200230122113003-0020312123100000-1010313101211021"></a>

## name property — Property reference / 003200213331 / 8

Type: `"string"`. Required.

Name of the BotInfrastructure.

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

<a id="canonical-3012303300032332-1231303012120303-3013210133101332-0110023020030101-0220012111031211-3301231233200003-0101031230310230-3121212012220013"></a>

<a id="canonical-1110230002100123-0000301031203202-2210310321221033-3113233222321221-3211201320321201-1322220222212001-0320012021320020-3230201031023213"></a>

## namespace property — Property reference / 003200213331 / 9

Type: `"string"`. Required.

Namespace where the BotInfrastructure exists.

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

<a id="canonical-0201300311030302-1333123012130320-1033330120223313-3200131010310113-3000023230020001-3203332232333230-1333100221101023-3032323233201110"></a>

<a id="canonical-1000311110111210-0330030012202010-3211013122100322-0002012001030113-1311320312012102-2110103223233321-1322333301032320-1133312200110322"></a>

## traffic_type property — Property reference / 003200213331 / 10

Type: `"string"`. Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

The type of traffic that is routed to and processed by this infrastructure (Web or Mobile).

Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot
Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are
routed through this Bot Defense infrastructure.

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

<a id="canonical-1100231130013330-0111203100113231-1100003032013032-2203000121031230-3320300322113002-2221003000323320-0202112320323321-2001102103110033"></a>

## All schema paths — Property reference / 003200213331 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bot_infrastructure--reference--group-001.md#canonical-0031201330213311-2111013301332330-2000222302202330-1020131203220213-3131312301223321-2000301313220022-2331312233133033-3213331121103012) |
| `create_cloud_hosted` | [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-0201212312221030-3222321211221200-2103003203213231-1223302020023313-2322112101200212-2321231003211133-2010230000321100-3113212313230133) |
| `create_cloud_hosted.ip_addresses` | [create_cloud_hosted.ip_addresses](data-sources--bot_infrastructure--reference--group-001.md#canonical-2231131101132331-2033310132120323-0203111221031032-3302311003002031-3323120222013221-0210221321001023-0023322131331021-0110021311301201) |
| `create_cloud_hosted.production` | [create_cloud_hosted.production](data-sources--bot_infrastructure--reference--group-001.md#canonical-2102331331033031-1021112220120301-2333333112110302-3113331101032113-3011103312120213-0323323012122110-3302113031202010-0113332102002312) |
| `create_cloud_hosted.production.region_1` | [create_cloud_hosted.production.region_1](data-sources--bot_infrastructure--reference--group-001.md#canonical-1332101003303330-0212231012230213-2031121330211113-1303322301012333-1030223230131322-0230301300202120-0102013122102330-2202232003323100) |
| `create_cloud_hosted.production.region_2` | [create_cloud_hosted.production.region_2](data-sources--bot_infrastructure--reference--group-001.md#canonical-0232022112303221-0300132211110012-2030012100000320-3301011232023023-2230020112313320-2033010310231110-1112102131202123-0101210130321232) |
| `create_cloud_hosted.testing` | [create_cloud_hosted.testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-0221103230003012-1121311101231203-2032231200222222-3123312011101212-0323321100202301-0001322221131321-0333113332220233-0330011020332300) |
| `create_cloud_hosted.testing.region_1` | [create_cloud_hosted.testing.region_1](data-sources--bot_infrastructure--reference--group-001.md#canonical-0232103210331230-2230033132230332-2113100120122203-1312022130321100-1012021203031233-1100301010233020-2323131331020012-2011012202112031) |
| `description` | [description](data-sources--bot_infrastructure--reference--group-001.md#canonical-1202313331002232-1113112120210022-3223000320102200-1120331101010100-3321202121013003-2323123013102133-2231122333231321-2112232011032003) |
| `id` | [ID](data-sources--bot_infrastructure--reference--group-001.md#canonical-2323300122311222-1122132113202301-1211221123031210-1023330320133311-3120330331023031-3100023311123010-0102001312333330-2313013103213103) |
| `labels` | [labels](data-sources--bot_infrastructure--reference--group-001.md#canonical-0102103023233312-3322222032023103-3223010220111103-2321210123333233-1312030123123021-2202322023233113-0022321102030020-0010301311221010) |
| `name` | [name](data-sources--bot_infrastructure--reference--group-001.md#canonical-2003220112120023-0212101022000001-2000113012012132-3202133132222013-0013223013330133-0212020011303102-3212113133010022-0113310300023003) |
| `namespace` | [namespace](data-sources--bot_infrastructure--reference--group-001.md#canonical-3012303300032332-1231303012120303-3013210133101332-0110023020030101-0220012111031211-3301231233200003-0101031230310230-3121212012220013) |
| `traffic_type` | [traffic_type](data-sources--bot_infrastructure--reference--group-001.md#canonical-0201300311030302-1333123012130320-1033330120223313-3200131010310113-3000023230020001-3203332232333230-1333100221101023-3032323233201110) |

<a id="canonical-3213230132030112-1000201022003200-1212311321310231-2023122300332102-3013021102022032-1002333110210212-3001001012001130-2302213212203131"></a>

## Next pages — Property reference / 003200213331 / 12

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)

<a id="canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003123322113212-1003133232223213-3211033020002121-0330030331033102-1030103111003232-3100022232231213-0103013201023132-0232322123030230"></a>

## create_cloud_hosted — create_cloud_hosted / 020332022333 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232)
- create_cloud_hosted

<a id="canonical-0201212312221030-3222321211221200-2103003203213231-1223302020023313-2322112101200212-2321231003211133-2010230000321100-3113212313230133"></a>

Type: `"single"`. Computed.

F5 Cloud Hosted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

<a id="canonical-0010201110323201-0003223122130121-3012301332303223-1220322011012222-3121232100103023-2213020003332131-3210033103310221-1013212230003102"></a>

## Direct properties — create_cloud_hosted / 020332022333 / 3

<a id="canonical-2231131101132331-2033310132120323-0203111221031032-3302311003002031-3323120222013221-0210221321001023-0023322131331021-0110021311301201"></a>

<a id="canonical-3320211202012111-3300113331223112-1132201200130202-2103120332323101-3320121330322100-2103001023110111-0021021113132333-3022212203101203"></a>

## ip_addresses property — create_cloud_hosted / 020332022333 / 4

Type: `["list", "string"]`. Computed.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](data-sources--bot_infrastructure--reference--group-001.md#canonical-3103113012102212-1110021012012323-2311131123000022-0031100313022110-2213303012213310-2101002231133011-2020002333012001-0333003212110131): complete subsection reference.

- [testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-0213311311211033-0120002121021033-1011203122121200-0221322330031300-3112200022130021-1000120033110023-0220210100330020-0320231303320302): complete subsection reference.

<a id="canonical-1121200122231111-3102132302202212-2011230331330202-0300100333321212-3300310011111022-2310323201001203-3001003223001221-2211001132300213"></a>

## Next pages — create_cloud_hosted / 020332022333 / 5

- [create_cloud_hosted.production](data-sources--bot_infrastructure--reference--group-001.md#canonical-3103113012102212-1110021012012323-2311131123000022-0031100313022110-2213303012213310-2101002231133011-2020002333012001-0333003212110131)
- [create_cloud_hosted.testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-0213311311211033-0120002121021033-1011203122121200-0221322330031300-3112200022130021-1000120033110023-0220210100330020-0320231303320302)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)

<a id="canonical-3103113012102212-1110021012012323-2311131123000022-0031100313022110-2213303012213310-2101002231133011-2020002333012001-0333003212110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223223231113200-0111200310102011-3333102011211003-3122112000301330-1313020100113223-0133222102301032-2113313323030112-1023323033232332"></a>

## create_cloud_hosted.production — production / 210320230210 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232)
- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121)
- create_cloud_hosted.production

<a id="canonical-2102331331033031-1021112220120301-2333333112110302-3113331101032113-3011103312120213-0323323012122110-3302113031202010-0113332102002312"></a>

Type: `"single"`. Computed.

Production.

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

<a id="canonical-0310201310101101-3323200000303132-3311313132123210-0310322101301103-3303200033000113-2023200233013203-3131310230032303-3132111022323213"></a>

## Direct properties — production / 210320230210 / 3

<a id="canonical-1332101003303330-0212231012230213-2031121330211113-1303322301012333-1030223230131322-0230301300202120-0102013122102330-2202232003323100"></a>

<a id="canonical-2312123333230010-3312133222222031-0100002103321321-2131111113301320-3031230213130200-3100003120021022-1213200023100303-2321213013222130"></a>

## region_1 property — production / 210320230210 / 4

Type: `"string"`. Computed.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-0232022112303221-0300132211110012-2030012100000320-3301011232023023-2230020112313320-2033010310231110-1112102131202123-0101210130321232"></a>

<a id="canonical-3321202031301013-0133003201002311-1310322210033130-3332101021313331-3201120032321110-0220111100121220-0011121013110320-1231031301300000"></a>

## region_2 property — production / 210320230210 / 5

Type: `"string"`. Computed.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-0012302031012313-3231232132101310-0110032003211313-2213120321200112-3211320320031000-2102321333133023-0103213032311301-3012322113221031"></a>

## Next pages — production / 210320230210 / 6

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)

<a id="canonical-0213311311211033-0120002121021033-1011203122121200-0221322330031300-3112200022130021-1000120033110023-0220210100330020-0320231303320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033133310202230-0131321333102202-0113130121010031-0130120112332303-0101201221233112-3313022203213031-1223213310122300-2311020330033321"></a>

## create_cloud_hosted.testing — testing / 102303011122 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232)
- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121)
- create_cloud_hosted.testing

<a id="canonical-0221103230003012-1121311101231203-2032231200222222-3123312011101212-0323321100202301-0001322221131321-0333113332220233-0330011020332300"></a>

Type: `"single"`. Computed.

Testing

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

<a id="canonical-2020100332030110-1113233123113203-1313022133111010-1012202110111203-0333302303221220-0232033103023302-3111110123220021-1301222211300322"></a>

## Direct properties — testing / 102303011122 / 3

<a id="canonical-0232103210331230-2230033132230332-2113100120122203-1312022130321100-1012021203031233-1100301010233020-2323131331020012-2011012202112031"></a>

<a id="canonical-3213320311031303-2322022023001031-2310001122311031-0102010233323302-2211213222220320-2300332001303322-3030313003031323-3130011111132301"></a>

## region_1 property — testing / 102303011122 / 4

Type: `"string"`. Computed.

Active-Passive Infrastructure configuration where traffic is routed to a single region.

Upstream description:

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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

<a id="canonical-0022131013333213-3112122031020120-2012211311131200-0320220032310302-1222302230201312-1022303000201110-3231310321011302-2002302211332003"></a>

## Next pages — testing / 102303011122 / 5

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
