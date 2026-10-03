---
page_title: "xcsh_bot_defense_app_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure reference."
---

# xcsh_bot_defense_app_infrastructure reference

<a id="canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121220123210201-3312312023001103-1123103220022000-3223212311321222-3020310221123101-3212001211321011-0002012100102110-0322122310332322"></a>

## Property reference — Property reference / 112202330230 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- Property reference

<a id="canonical-0031131210032212-1013323302102012-2013011230020230-3020233011000022-3312310303021202-3133211112300013-1333332021232323-3120311230020030"></a>

## Direct properties — Property reference / 112202330230 / 3

<a id="canonical-3230210120313132-1013203110120333-0101123013030223-3010231013120332-2301020031000302-1001321220213310-3232331301203122-0200220232202330"></a>

<a id="canonical-1200202110310210-0320001302202320-0220321232031311-2133003030123300-3310011323331321-2332212032020032-0131102011233212-1301200201032031"></a>

## annotations property — Property reference / 112202330230 / 4

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

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333): complete subsection reference.

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001): complete subsection reference.

<a id="canonical-2313121102020301-1212113211222200-3210220030323122-1201221101321103-0102301213221333-0322331131331312-1230111111220330-0021332031301220"></a>

<a id="canonical-3113010123032030-3232231331223111-3220103003030330-2311231101230132-0132023031303011-3130121021220303-2103021003321212-2030331020200130"></a>

## description property — Property reference / 112202330230 / 5

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

<a id="canonical-2003333100011202-1133021000203031-1032332032213303-2110112301300310-3312133211301003-2032131312113311-0301300322122210-3201303230303203"></a>

<a id="canonical-3023000201100200-3003112223330303-0313201300130331-0131233200101011-3002202000222101-3120201323302021-1223300332212331-3222031032321230"></a>

## disable property — Property reference / 112202330230 / 6

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

<a id="canonical-1332013132322302-1000101000300323-3201202201302320-0000113303010322-2000231001200111-1232122310210202-3230332302301200-3003112201110230"></a>

<a id="canonical-0211300211101223-1010321322010301-1003003002021311-3202213200323333-2120231333032102-3233221321220023-1022211121300103-3001113030200133"></a>

## environment_type property — Property reference / 112202330230 / 7

Type: `"string"`. Optional, Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PRODUCTION",
    "TESTING"),
}
```

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

<a id="canonical-1221121110330021-2003110230213211-2102002101200020-1321023120111332-0133233120100202-1021330202223011-2323302321132123-0330333030033323"></a>

<a id="canonical-3201031122213023-1100012112100203-0312320320011333-0203031123302233-2102103230300130-3331033132332300-2031201211212231-0111113303023223"></a>

## ID property — Property reference / 112202330230 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1122113221332022-0200332020220321-1322002032330132-1202320302212113-1021033030230233-1203000222231211-1033031123001020-0122321202111111"></a>

<a id="canonical-3200323332002113-3233330031210230-2003112201122310-2202133103300000-0132323112011312-3302000232122100-1100323121100120-1110130000133211"></a>

## labels property — Property reference / 112202330230 / 9

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

<a id="canonical-1033113032301201-1301101131222000-0221032313033312-3202332001333202-0103323032121132-1310310002233013-1321111022100020-0130313032313310"></a>

<a id="canonical-1130200230310012-1120310302022012-3221010022310101-2213203122111312-2100320232223223-2012330020300021-2122002023020002-1310323200200213"></a>

## name property — Property reference / 112202330230 / 10

Type: `"string"`. Required.

Name of the Bot Defense App Infrastructure. Must be unique within the namespace.

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

<a id="canonical-3001321331113223-0302231000133013-3221020302312003-0131212231000003-3223301223030100-3111221132211001-1013111221221301-1020220200033113"></a>

<a id="canonical-3020212201112121-1132021021022320-3011211203300203-2200213032001331-3113231013000300-0222020211131221-0101302032103320-1013133013002021"></a>

## namespace property — Property reference / 112202330230 / 11

Type: `"string"`. Required.

Namespace where the Bot Defense App Infrastructure is created.

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

- [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0310301100000323-2023023221030012-0113123123113310-2121330331310003-3300311220110211-3213012210211102-3003320223030322-0332013111312110): complete subsection reference.

<a id="canonical-3030332021033210-3322213232000320-1320222032332203-0030203030022032-2120202122301123-1233122201331300-2100100310023100-3210123110312313"></a>

<a id="canonical-2322003332001210-0113211332022221-3022030103211011-1323032113333300-1022111032320212-2023030133302123-0113022320010002-3223302113331222"></a>

## traffic_type property — Property reference / 112202330230 / 12

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("WEB",
    "MOBILE"),
}
```

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

<a id="canonical-2223003121221112-3231320322120102-2030232102112230-1012313000313131-1300320002101033-1310223103031122-3312012321002300-0231310023223213"></a>

## All schema paths — Property reference / 112202330230 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3230210120313132-1013203110120333-0101123013030223-3010231013120332-2301020031000302-1001321220213310-3232331301203122-0200220232202330) |
| `cloud_hosted` | [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3122220300331331-0211213032220023-0121013103101011-2331022130000020-1223301000223312-0223001122013012-1203203022320113-1232000312321010) |
| `cloud_hosted.egress` | [cloud_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1302320301122102-2003120333120312-0021310132131131-2122333231020323-3031223302132232-3131311120033212-0133123121233122-0032131111010323) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0301323332233102-2213300210101222-0121310230211222-0331131110330201-2123110302221012-2120121201011323-3131032133111311-0312112212332132) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3321011031201021-1200022331001130-2300100101101213-3102233100023230-3022013013222332-3330222123322133-1323312102331001-2201032320110020) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2032113331221231-2300232321210031-3220000310212021-1323021330131232-0000220321022233-0033312003212113-3310033232223023-3301000210212202) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3023022210213010-0211020111131031-3120013222130322-3311203231301131-2111101103012333-3320010303120303-0221022213122013-0100033001100232) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2021133020013022-1101332330123221-3133310010032012-2310033101011202-3231312330323301-2132022331301330-3002012103220303-0132202203111032) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3333222130233020-2023000103311323-1131012301312333-1233003302010101-2012300130132233-1321213032110110-3321021311021201-2000222103123012) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3310233032122103-1132112211200330-2323201110110111-3320303231321312-3333003231230300-1311023333320013-3001232201331122-0101212033112332) |
| `cloud_hosted.region` | [cloud_hosted.region](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1001100110300111-0222120103333302-1300220132120301-3101323332230102-1310322203323333-1003310010203133-3313113002112223-3132011321020113) |
| `data_center_hosted` | [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1222131013332232-3113201330122121-2223201233003213-0201023021200300-3121322321033303-1020232331032113-3021310230220221-3300021231022103) |
| `data_center_hosted.egress` | [data_center_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2103012321100302-3222202212122121-0211013023212331-1230100003110121-0220200321212132-0303312110112111-3223312310310001-2033133002020201) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0232221110312200-1310102001201333-2203131023301222-2301203111133231-0331300030322331-2322320201020133-0000021102201013-2122233211113210) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1210202100233303-3223210003121111-0300320000311333-3122222320121313-2202022202003020-2220131133123111-0103203010101102-1031321221131301) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1001011221002133-3312011011201001-3210031300312310-3203222113213212-3101031313033002-1121102023131221-3011000313222233-0013202013333030) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3221133031022310-1232013013100032-2021210031330322-0200333213013132-2233023020211232-0331221233212220-0322312221102211-1330301233022220) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2233130223302211-1032211303322102-1000102001230310-1311311011002310-1002121210011122-1212212312213003-3330132022322230-2330232031312321) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132222010322022-2032221302323133-2330110333233213-1011323322303331-2111200133321310-0003332232330133-0213030022220201-1323212331221203) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0222013333102122-3100300311211120-1003020101221023-2231213110110303-0011330203232223-0211102220333333-2012010002121231-1221022301130132) |
| `data_center_hosted.region` | [data_center_hosted.region](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132320112313332-3202311223112301-2111312233331101-2231023003112332-2000332013230210-0131130203212311-2013323002001101-1330201332222133) |
| `description` | [description](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2313121102020301-1212113211222200-3210220030323122-1201221101321103-0102301213221333-0322331131331312-1230111111220330-0021332031301220) |
| `disable` | [disable](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2003333100011202-1133021000203031-1032332032213303-2110112301300310-3312133211301003-2032131312113311-0301300322122210-3201303230303203) |
| `environment_type` | [environment_type](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1332013132322302-1000101000300323-3201202201302320-0000113303010322-2000231001200111-1232122310210202-3230332302301200-3003112201110230) |
| `id` | [ID](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1221121110330021-2003110230213211-2102002101200020-1321023120111332-0133233120100202-1021330202223011-2323302321132123-0330333030033323) |
| `labels` | [labels](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1122113221332022-0200332020220321-1322002032330132-1202320302212113-1021033030230233-1203000222231211-1033031123001020-0122321202111111) |
| `name` | [name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1033113032301201-1301101131222000-0221032313033312-3202332001333202-0103323032121132-1310310002233013-1321111022100020-0130313032313310) |
| `namespace` | [namespace](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3001321331113223-0302231000133013-3221020302312003-0131212231000003-3223301223030100-3111221132211001-1013111221221301-1020220200033113) |
| `timeouts` | [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2002110233130321-3101031103021232-3101303210101121-1130300012330312-2302033303023122-0130020223112231-3313000322333320-3010011022322133) |
| `timeouts.create` | [timeouts.create](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1122103200122111-3032102103033312-3201311001002202-3022332133302312-0233002223220103-3102231022320202-0123212233122222-3113231220110233) |
| `timeouts.delete` | [timeouts.delete](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2103022033033320-1031303100033100-0122112011013110-2110133222021313-2331030001130022-2120133113010212-1102031232223012-3000233300033113) |
| `timeouts.read` | [timeouts.read](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1213331211203030-3033203233203212-0020320230012212-1021303123010000-0131221333313320-3013022232033323-1233020103201102-2132022022232311) |
| `timeouts.update` | [timeouts.update](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2102110301111030-3102122020301120-0232130313320200-1021213133010130-3322001133023110-1121202023332332-1131102223030203-3220301102123023) |
| `traffic_type` | [traffic_type](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3030332021033210-3322213232000320-1320222032332203-0030203030022032-2120202122301123-1233122201331300-2100100310023100-3210123110312313) |

<a id="canonical-3123223230303300-1111101301223110-2030333210033200-2000020232120313-2303322033310110-0112020101221133-0011322120313222-1232212123012221"></a>

## Next pages — Property reference / 112202330230 / 14

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001)
- [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0310301100000323-2023023221030012-0113123123113310-2121330331310003-3300311220110211-3213012210211102-3003320223030322-0332013111312110)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203103122213130-1223220130213231-2022110021022023-0222213131333321-1112300030103113-3132102122011000-1000111321032333-0201233300211100"></a>

## cloud_hosted — cloud_hosted / 112022210303 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- cloud_hosted

<a id="canonical-3122220300331331-0211213032220023-0121013103101011-2331022130000020-1223301000223312-0223001122013012-1203203022320113-1232000312321010"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("egress",
    "infra_host_name",
    "ingress")}
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

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3122220300331331-0211213032220023-0121013103101011-2331022130000020-1223301000223312-0223001122013012-1203203022320113-1232000312321010)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1222131013332232-3113201330122121-2223201233003213-0201023021200300-3121322321033303-1020232331032113-3021310230220221-3300021231022103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cloud_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110312011103032-1020331130103113-0130312302020121-1022232033332123-2220203023222032-0103030100303031-0213232013300223-3332200031202113"></a>

## Direct properties — cloud_hosted / 112022210303 / 3

- [egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1032211213322112-1123203302133331-0222032003202003-0300011103313101-2323133203233331-2022231221300201-1101332032013222-0023010001213102): complete subsection reference.

<a id="canonical-2032113331221231-2300232321210031-3220000310212021-1323021330131232-0000220321022233-0033312003212113-3310033232223023-3301000210212202"></a>

<a id="canonical-2312223032300123-0021211101122020-2213003133003002-3211111130032320-2211100220221030-1111033331201110-3130203202100303-0303313100302210"></a>

## infra_host_name property — cloud_hosted / 112022210303 / 4

Type: `"string"`. Optional.

Infra hostname. Infra hostname.

Upstream description:

Infra hostname.

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
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3323023203121020-1200022222333233-0223323033211020-2303112221023011-2210033112200111-3303132212020330-3310132103213223-3001033201233303): complete subsection reference.

<a id="canonical-1001100110300111-0222120103333302-1300220132120301-3101323332230102-1310322203323333-1003310010203133-3313113002112223-3132011321020113"></a>

<a id="canonical-2302102331330032-3033333223212101-2100333321231131-3320032031333213-0210232323203310-3020223312231310-3130100002120022-0031320301301310"></a>

## region property — cloud_hosted / 112022210303 / 5

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA"),
}
```

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

<a id="canonical-1223002302031030-0231133200123203-1111102002101101-0230022231020313-1210123323003103-1000320000012331-2023332211033330-2023331332032233"></a>

## Next pages — cloud_hosted / 112022210303 / 6

- [cloud_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1032211213322112-1123203302133331-0222032003202003-0300011103313101-2323133203233331-2022231221300201-1101332032013222-0023010001213102)
- [cloud_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3323023203121020-1200022222333233-0223323033211020-2303112221023011-2210033112200111-3303132212020330-3310132103213223-3001033201233303)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-1032211213322112-1123203302133331-0222032003202003-0300011103313101-2323133203233331-2022231221300201-1101332032013222-0023010001213102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102323332001303-3122200303311201-0313030210213220-3020223312020120-3100321313113302-0122021133303313-2302102133100100-1332311102303121"></a>

## cloud_hosted.egress — egress / 211302033010 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333)
- cloud_hosted.egress

<a id="canonical-1302320301122102-2003120333120312-0021310132131131-2122333231020323-3031223302132232-3131311120033212-0133123121233122-0032131111010323"></a>

Type: `"object"`. list nested block, Optional.

Egress. Egress

Upstream description:

Egress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_address")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
egress {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311031133103111-0012111320323012-3200230320320123-0320301233130320-3332302103332320-0023122101320322-0031220100302112-2200233120330003"></a>

## Direct properties — egress / 211302033010 / 3

<a id="canonical-0301323332233102-2213300210101222-0121310230211222-0331131110330201-2123110302221012-2120121201011323-3131032133111311-0312112212332132"></a>

<a id="canonical-0212331103033222-0111112031313320-3333012301322221-0311312011022201-3030210232122301-1103222231123001-3001322223001323-3033110023131203"></a>

## ip_address property — egress / 211302033010 / 4

Type: `"string"`. Optional.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPValidator(),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3321011031201021-1200022331001130-2300100101101213-3102233100023230-3022013013222332-3330222123322133-1323312102331001-2201032320110020"></a>

<a id="canonical-0120211220301302-0330010102310201-1032312313032303-2021302212321132-2032020101201123-0111133121020211-1110030120022311-0332130131003311"></a>

## location property — egress / 211302033010 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
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
    "GCP_US_WEST_2"),
}
```

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

<a id="canonical-3202310031102122-2121112223331101-2232332120010131-0101113130303312-1301023010330332-3333132332131201-3001031333132132-1201323123230202"></a>

## Next pages — egress / 211302033010 / 6

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-3323023203121020-1200022222333233-0223323033211020-2303112221023011-2210033112200111-3303132212020330-3310132103213223-3001033201233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121101013303300-0022103320320333-0202210133130221-2133102310213231-1102202300223112-0330002301031132-3023030201113320-3310002202101203"></a>

## cloud_hosted.ingress — ingress / 302303201220 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333)
- cloud_hosted.ingress

<a id="canonical-3023022210213010-0211020111131031-3120013222130322-3311203231301131-2111101103012333-3320010303120303-0221022213122013-0100033001100232"></a>

Type: `"object"`. list nested block, Optional.

Ingress. Ingress

Upstream description:

Ingress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("host_name",
    "ip_address")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
ingress {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323030121322322-1020122213213100-0323313201303010-2332002112011013-2303121210222302-0020331020212323-0313311033213122-0313001213103031"></a>

## Direct properties — ingress / 302303201220 / 3

<a id="canonical-2021133020013022-1101332330123221-3133310010032012-2310033101011202-3231312330323301-2132022331301330-3002012103220303-0132202203111032"></a>

<a id="canonical-3000321110220310-3301132311102031-2221031312333101-3100233212212212-3121021032133023-2202331010123221-2210321110103120-2221313213313331"></a>

## host_name property — ingress / 302303201220 / 4

Type: `"string"`. Optional.

Exclusive with \[ip\_address\] Ingress hostname.

Upstream description:

Exclusive with \[ip\_address\] Ingress hostname.

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
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3333222130233020-2023000103311323-1131012301312333-1233003302010101-2012300130132233-1321213032110110-3321021311021201-2000222103123012"></a>

<a id="canonical-1221131132213332-0021322300321310-1111302020010101-0030232123330203-3213222102323201-1103022212033033-1013303111230301-0100101330221113"></a>

## ip_address property — ingress / 302303201220 / 5

Type: `"string"`. Optional.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPValidator(),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3310233032122103-1132112211200330-2323201110110111-3320303231321312-3333003231230300-1311023333320013-3001232201331122-0101212033112332"></a>

<a id="canonical-3103000202313220-2230120011000203-1131333323121211-0213200210330102-2120330000202131-1213000030121233-0230010221112230-0001123311312131"></a>

## location property — ingress / 302303201220 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
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
    "GCP_US_WEST_2"),
}
```

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

<a id="canonical-3232221110003333-3321021300313201-2022131211211010-3233101321212110-2033221213012300-0221023111231332-1332202132002301-2222123200111330"></a>

## Next pages — ingress / 302303201220 / 7

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113321312311300-3320122203221210-3222010333300320-3111113103331113-3101002201311223-2233223020121210-0121023123131310-1231333103323313"></a>

## data_center_hosted — data_center_hosted / 032312203010 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- data_center_hosted

<a id="canonical-1222131013332232-3113201330122121-2223201233003213-0201023021200300-3121322321033303-1020232331032113-3021310230220221-3300021231022103"></a>

Type: `"object"`. single nested block, Optional.

F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("egress",
    "infra_host_name",
    "ingress")}
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
data_center_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220123100133320-3002033120121210-3130322333313133-1310332132212213-2033331022010303-3011230313102203-2032132302103231-1333112210330321"></a>

## Direct properties — data_center_hosted / 032312203010 / 3

- [egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0332101112212130-0312320021033012-1033103101113012-0010013311221313-1002002033323323-3130233333012311-2311232022033230-0002033121020211): complete subsection reference.

<a id="canonical-1001011221002133-3312011011201001-3210031300312310-3203222113213212-3101031313033002-1121102023131221-3011000313222233-0013202013333030"></a>

<a id="canonical-1200310330221210-2220232003220222-0113000223221301-3021220122110332-3003130310322302-3333132133221301-3301202333322131-1212301231323203"></a>

## infra_host_name property — data_center_hosted / 032312203010 / 4

Type: `"string"`. Optional.

Infra hostname. Infra hostname.

Upstream description:

Infra hostname.

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
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2201323113322321-0111313113001200-3002101003101212-3023123033010000-0213232310310000-2320212310002211-3102121222313331-0303213032101330): complete subsection reference.

<a id="canonical-2132320112313332-3202311223112301-2111312233331101-2231023003112332-2000332013230210-0131130203212311-2013323002001101-1330201332222133"></a>

<a id="canonical-3122322031230103-0010002120221013-2121330023332300-3020300022033113-0233022323002230-1332103202201101-2130213011013210-0323201103003112"></a>

## region property — data_center_hosted / 032312203010 / 5

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA"),
}
```

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

<a id="canonical-1020303031312231-1102021103020033-0332101102131311-2103231011101230-1333211031312111-1323233030031010-0303120302331131-0231231123003123"></a>

## Next pages — data_center_hosted / 032312203010 / 6

- [data_center_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0332101112212130-0312320021033012-1033103101113012-0010013311221313-1002002033323323-3130233333012311-2311232022033230-0002033121020211)
- [data_center_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2201323113322321-0111313113001200-3002101003101212-3023123033010000-0213232310310000-2320212310002211-3102121222313331-0303213032101330)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-0332101112212130-0312320021033012-1033103101113012-0010013311221313-1002002033323323-3130233333012311-2311232022033230-0002033121020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321100103333331-3313232320233111-1103011130011022-1033013010111123-0111321212031002-0200322113223310-1300021130212202-1222330310201103"></a>

## data_center_hosted.egress — egress / 112031313211 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001)
- data_center_hosted.egress

<a id="canonical-2103012321100302-3222202212122121-0211013023212331-1230100003110121-0220200321212132-0303312110112111-3223312310310001-2033133002020201"></a>

Type: `"object"`. list nested block, Optional.

Egress. Egress

Upstream description:

Egress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_address")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
egress {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022312332021222-0313233001202120-0030310221301100-1233111233311312-3330133121013000-2222300322303111-3300210313203330-0310000232302220"></a>

## Direct properties — egress / 112031313211 / 3

<a id="canonical-0232221110312200-1310102001201333-2203131023301222-2301203111133231-0331300030322331-2322320201020133-0000021102201013-2122233211113210"></a>

<a id="canonical-2110032011322112-3300312112010003-0110000203331021-3221203032033212-0330200333010132-1102102323033100-0121013002033231-3301033232231220"></a>

## ip_address property — egress / 112031313211 / 4

Type: `"string"`. Optional.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPValidator(),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1210202100233303-3223210003121111-0300320000311333-3122222320121313-2202022202003020-2220131133123111-0103203010101102-1031321221131301"></a>

<a id="canonical-3132212300311133-2111130000132131-0013110103200131-2112030202232030-1112211310221001-3111030203002013-2132012132231030-1012311120003312"></a>

## location property — egress / 112031313211 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
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
    "GCP_US_WEST_2"),
}
```

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

<a id="canonical-1320010303331121-0231123211203213-3031121022320021-2230100201102123-3101203122212200-3123300001000230-0312030312101122-3223032330221213"></a>

## Next pages — egress / 112031313211 / 6

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-2201323113322321-0111313113001200-3002101003101212-3023123033010000-0213232310310000-2320212310002211-3102121222313331-0303213032101330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203210323311203-0310012102200220-3220301301333200-0323310013003310-1121010333202011-2011011310211101-0003121211023232-3203302111302000"></a>

## data_center_hosted.ingress — ingress / 333322203223 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001)
- data_center_hosted.ingress

<a id="canonical-3221133031022310-1232013013100032-2021210031330322-0200333213013132-2233023020211232-0331221233212220-0322312221102211-1330301233022220"></a>

Type: `"object"`. list nested block, Optional.

Ingress. Ingress

Upstream description:

Ingress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("host_name",
    "ip_address")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
ingress {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230032311221212-1023111313332131-1002031322331023-1021333323323001-0003300101011121-0133123230013102-1120101201222311-2102132312232001"></a>

## Direct properties — ingress / 333322203223 / 3

<a id="canonical-2233130223302211-1032211303322102-1000102001230310-1311311011002310-1002121210011122-1212212312213003-3330132022322230-2330232031312321"></a>

<a id="canonical-3010111300201202-1032101103123021-0222230133003200-2303221311332000-3013211311213033-3130111112323100-3120313213312113-0222230313012020"></a>

## host_name property — ingress / 333322203223 / 4

Type: `"string"`. Optional.

Exclusive with \[ip\_address\] Ingress hostname.

Upstream description:

Exclusive with \[ip\_address\] Ingress hostname.

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
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2132222010322022-2032221302323133-2330110333233213-1011323322303331-2111200133321310-0003332232330133-0213030022220201-1323212331221203"></a>

<a id="canonical-3321323100300211-0002120230313003-3121100220030012-1013122330133120-3132301333321003-3223110002323322-3111320310010200-2102103003131000"></a>

## ip_address property — ingress / 333322203223 / 5

Type: `"string"`. Optional.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPValidator(),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0222013333102122-3100300311211120-1003020101221023-2231213110110303-0011330203232223-0211102220333333-2012010002121231-1221022301130132"></a>

<a id="canonical-1012230100302210-0022011213213230-1202022000202303-1222002331313212-2302102210031222-2000023301002301-0211331012303133-0101332111310211"></a>

## location property — ingress / 333322203223 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
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
    "GCP_US_WEST_2"),
}
```

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

<a id="canonical-2212103211212101-1032223201121231-2133213031003120-3110200200112323-2021232321302301-1120020312133223-0221213023012231-2032003030010111"></a>

## Next pages — ingress / 333322203223 / 7

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)

<a id="canonical-0310301100000323-2023023221030012-0113123123113310-2121330331310003-3300311220110211-3213012210211102-3003320223030322-0332013111312110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121022103010232-3310212311302231-3102321321223333-0202233032031111-1112012110221013-3212131202301121-3002012221200330-0020030213022012"></a>

## timeouts — timeouts / 123232113102 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- timeouts

<a id="canonical-2002110233130321-3101031103021232-3101303210101121-1130300012330312-2302033303023122-0130020223112231-3313000322333320-3010011022322133"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300322122111032-1133113010332032-1033303203210122-3222302031133031-1131310032330000-3113113211203201-3332013223122111-0200011231330032"></a>

## Direct properties — timeouts / 123232113102 / 3

<a id="canonical-1122103200122111-3032102103033312-3201311001002202-3022332133302312-0233002223220103-3102231022320202-0123212233122222-3113231220110233"></a>

<a id="canonical-0201200130202222-0231003230310232-3032020210120322-0201313123002103-2332130321232111-3020323123132121-0000023121103111-0323300123103200"></a>

## create property — timeouts / 123232113102 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2103022033033320-1031303100033100-0122112011013110-2110133222021313-2331030001130022-2120133113010212-1102031232223012-3000233300033113"></a>

<a id="canonical-2132103212112121-2013002033311131-0100311112102212-3101012101001202-1321322001003132-1020121111013333-1012110031332032-2333130231300000"></a>

## delete property — timeouts / 123232113102 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1213331211203030-3033203233203212-0020320230012212-1021303123010000-0131221333313320-3013022232033323-1233020103201102-2132022022232311"></a>

<a id="canonical-0002030100023202-2331113021310223-2232112330312200-1123212310202022-2022330311112002-1222211121220110-2020000210213212-2203233012322010"></a>

## read property — timeouts / 123232113102 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2102110301111030-3102122020301120-0232130313320200-1021213133010130-3322001133023110-1121202023332332-1131102223030203-3220301102123023"></a>

<a id="canonical-1212232200232203-0221301302233223-3202131222200311-3330321301330211-2120231132311232-3301113022133202-3011100031002202-3111113231202300"></a>

## update property — timeouts / 123232113102 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1020032003320332-3122112323001321-2130331231033000-1223020131123213-3310233100112103-0320130012221223-1113012301301230-2321021302033233"></a>

## Next pages — timeouts / 123232113102 / 8

- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
