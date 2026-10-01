---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032123332200023-2313100111011303-2021330020221231-1033302030202313-3320312211132332-2211231003111011-2211132120031302-1213013303101322"></a>

## Property reference — Property reference / 103230111112 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- Property reference

<a id="canonical-2122231201112112-3230221020332331-1031110313332330-0220133020003231-0313210033133313-0021020200103033-0230302001320203-0130123030210102"></a>

## Direct properties — Property reference / 103230111112 / 3

<a id="canonical-2223333021213212-1212133123210320-3202232030231332-1200222303222101-3003021331110222-0201213301333002-3012111010303321-2332203013113021"></a>

<a id="canonical-0200233330332331-1131300133302303-3330201333232300-3133031101322003-2110001311110203-3113332120020031-2203222333303002-2331230112300322"></a>

## address property — Property reference / 103230111112 / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2213101020002111-0230033231020231-2220312213311333-0131120203021321-3313002000011300-0121230103011222-3320010203232110-0100330330302231"></a>

<a id="canonical-3303002303102211-3011331222112020-3230212021032212-1230231111031031-0013232233331202-3113300210123223-1103303203321133-1222331032101011"></a>

## annotations property — Property reference / 103230111112 / 5

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

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231): complete subsection reference.

- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232): complete subsection reference.

- [coordinates](resources--securemesh_site--reference--group-001.md#canonical-1222331220131320-1233331310213200-1221203223323001-1313113213210333-2110020133102120-2131211230222000-2022131203310321-3311113111312312): complete subsection reference.

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230): complete subsection reference.

- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-2130322011122113-2023213101220320-0310300323120322-1213333313311002-3123113231010313-0110102302312320-1200031302010203-0311320121030301): complete subsection reference.

- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-1211212023033021-3013201123022023-0023321103121133-1222322023300323-0131132032303333-3332330303301122-2032220120223122-2021102112323110): complete subsection reference.

<a id="canonical-2023211221311000-0102123210113130-3133323112001132-2112202203121011-2112032223200003-3020320020111101-1313221020001302-1320330333023012"></a>

<a id="canonical-2303220302031201-0131001332332030-2012210030123133-3301332333100032-2012022013221032-2030010302233221-2110132223331202-0221313212231100"></a>

## description property — Property reference / 103230111112 / 6

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

<a id="canonical-1302322210013111-0133321020313031-0310010110032011-0223103020310201-0022330021003031-1210210312030022-1010011323302122-1201032113130331"></a>

<a id="canonical-1013020030132310-3201022031121031-3332113313111310-1013003002232021-0022230002310122-2133302323300110-3320011200312112-2113032023001232"></a>

## disable property — Property reference / 103230111112 / 7

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

<a id="canonical-1102213121023221-2032300332101100-3023020011222000-3101023333032112-0032312031213211-0322033132022033-0301202101322333-1212003202013113"></a>

<a id="canonical-2221020222201322-2033213213322021-2102002233133112-2303002321033332-2223131013220012-3032321031220203-0303213001200330-0310023100133132"></a>

## ID property — Property reference / 103230111112 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221): complete subsection reference.

<a id="canonical-3203231321003101-1323220303111222-3013200222311121-0320213132100030-1013312222312220-3033230132001201-0033302331211031-0113220312012301"></a>

<a id="canonical-0211213102213101-3023021103212223-3002312311023021-1221011310223320-0311320130332011-2203323013310211-3113311111003201-3110121222313002"></a>

## labels property — Property reference / 103230111112 / 9

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

- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-1210231003210032-0210302031033031-1321120333323332-0113030213300300-3003201313113111-2321333000300322-1312122112223131-2303302010130323): complete subsection reference.

- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-0301322311211110-3313100231331120-3020131300021210-0031233031211302-0223101202102010-3322120202121023-2112032200010220-2203211233300002): complete subsection reference.

- [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-2030021313103121-0101111110133002-2130111211310100-2020102210122203-3100100030322221-2033211312031111-3022033203312230-3201003223312122): complete subsection reference.

<a id="canonical-3213003332010010-3133230111231132-0023221012011322-2100320131313120-2211123100012013-3333313112110202-1302000232223112-2033230310012331"></a>

<a id="canonical-3101032332322230-0232210333122030-3202031030102303-2322030211021030-2300031302203200-3211330102122031-0331310330122002-3211221321232202"></a>

## name property — Property reference / 103230111112 / 10

Type: `"string"`. Required.

Name of the Securemesh Site. Must be unique within the namespace.

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

<a id="canonical-0010133111201002-2032112111030133-1122331020213130-2302321302030200-0032010333320223-3011230320021231-2220003230330223-1112111021320011"></a>

<a id="canonical-3112123220321011-2111102231203210-1223323203002221-2301013122123231-1012202323021132-3130103310221020-0101231003210210-2013033112220002"></a>

## namespace property — Property reference / 103230111112 / 11

Type: `"string"`. Required.

Namespace where the Securemesh Site is created.

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

- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-0333112101031113-3121130300222310-2130011201332013-2012000233233023-2231013313212021-1321312222230212-0323220201313111-3200133332031330): complete subsection reference.

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100): complete subsection reference.

- [os](resources--securemesh_site--reference--group-004.md#canonical-0102133232320131-2231010312223023-0203311201032131-2210230021210322-0310110201031312-2321113100212222-0201202312033330-3000120030223322): complete subsection reference.

- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033): complete subsection reference.

- [sw](resources--securemesh_site--reference--group-004.md#canonical-1000012000123120-2210011233010203-3333022101301022-1133130210233321-1101033113332231-1333113032202102-2032212300223010-2110321233023213): complete subsection reference.

- [timeouts](resources--securemesh_site--reference--group-004.md#canonical-3231302133333333-3210323010223133-3130032303201002-1121332221230320-2013212123011031-2102312110221123-1333231321102102-3213311103101232): complete subsection reference.

<a id="canonical-2201331321112131-2223032000131211-0031230323101122-2001022033121020-2021112031300120-3202221013313332-1131112323020110-3103110303311032"></a>

<a id="canonical-2121301122233201-1122002110330222-3321322300003033-0310101311112302-1102232313332200-0300002332023011-2323120113233020-3303211001120222"></a>

## volterra_certified_hw property — Property reference / 103230111112 / 12

Type: `"string"`. Required.

Name for generic server certified hardware to form this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200): complete subsection reference.

<a id="canonical-1220330132001101-3301000301213310-3232121012110020-1321011320000203-0121300200232311-0323213212203122-2320111301102230-0130130310023122"></a>

<a id="canonical-0012020303220102-1213111011021030-3121111100311102-1231203301013003-3230333032030131-2000012030302202-0313311131231320-3020010002203301"></a>

## worker_nodes property — Property reference / 103230111112 / 13

Type: `["list", "string"]`. Optional.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333103201123311-1033210202223031-2033212113200201-2111331302113122-2030122330022110-3012210212033230-1012021101221123-3223313121221030"></a>

## All schema paths — Property reference / 103230111112 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--securemesh_site--reference--group-001.md#canonical-2223333021213212-1212133123210320-3202232030231332-1200222303222101-3003021331110222-0201213301333002-3012111010303321-2332203013113021) |
| `annotations` | [annotations](resources--securemesh_site--reference--group-001.md#canonical-2213101020002111-0230033231020231-2220312213311333-0131120203021321-3313002000011300-0121230103011222-3320010203232110-0100330330302231) |
| `blocked_services` | [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-1323331333032130-1212021222032303-1231200023110010-1223330121132013-3332113300121023-0203313201001010-3121311032130332-3100313123123303) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-0330011031130302-3233323031112113-3102033333121013-3010102123102302-2132201312331133-1222021231133002-3022210223032100-1320010323312032) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--securemesh_site--reference--group-001.md#canonical-0113110303022303-2203030323221213-2003233121310223-0023030001010010-1203010023010333-2023110031100332-3023112021213231-1321022200123112) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--securemesh_site--reference--group-001.md#canonical-2302100023010120-3130333210131211-1233300103321133-0213002231030120-1103121013220010-1122223010030300-2113101100202231-2120221312111030) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--securemesh_site--reference--group-001.md#canonical-0021010212112212-3331101013231022-0123033331001212-0032101223003132-3201302230103303-1132310221210121-3031212201103030-2022320220213122) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-0330033130303123-1003100130333232-2221022303222022-1022011211331033-2332222022032021-1213212200210103-2322220321132011-3233131222211210) |
| `bond_device_list` | [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2200121111110131-1232233011211322-2200022100310221-3010233332001312-0230321231201201-2003033310002110-1113333310122311-0211303103120030) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-3002022202303211-1000130111031023-1233132202111210-1330202322233210-3212202113332210-3011200111211131-3221212100103303-3301232132320203) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](resources--securemesh_site--reference--group-001.md#canonical-1013010322301133-2213221213300231-0020123112203103-3322303112321123-3022203211023020-0021211121122113-1300322032112302-3102312102300323) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](resources--securemesh_site--reference--group-001.md#canonical-3300202223202001-1122012313201200-3031322000321111-1201322203221201-1101012132103113-2200223011110202-3222211312103012-2033032000112233) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](resources--securemesh_site--reference--group-001.md#canonical-1311311100112200-0010303010312131-2013331001231102-3223321001222031-2022230211320212-2300013023323232-2130113120203302-2011300123222230) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](resources--securemesh_site--reference--group-001.md#canonical-1113311310223120-3223002013233131-3313101030033211-1323311000233021-0022301313220032-2030100132310022-1303011013120213-2330031222232322) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](resources--securemesh_site--reference--group-001.md#canonical-3110133020122311-1211301023211121-3331022120310121-2200222132211230-0012312212103130-0333313132300012-1013121102100321-2022110001333200) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](resources--securemesh_site--reference--group-001.md#canonical-1223222333323002-0013021331201222-0133311101102332-0021322302110312-0130310203003101-0232211003203310-0121032221333312-3020233111113132) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](resources--securemesh_site--reference--group-001.md#canonical-3132211300333023-1313000201333032-0211320231131102-2222013320100203-3300322011133100-2110003113030223-3031320202200010-0110220100130300) |
| `coordinates` | [coordinates](resources--securemesh_site--reference--group-001.md#canonical-2322032123200012-1300021322020212-0323232000212132-1133123020330113-2020310330010010-3221200210201201-1330002012213220-2331312022020113) |
| `coordinates.latitude` | [coordinates.latitude](resources--securemesh_site--reference--group-001.md#canonical-3132202011221022-0322010002101201-2113123100300221-2332122111220333-3201301132202201-1031000131302300-1003100203002111-1223021113132330) |
| `coordinates.longitude` | [coordinates.longitude](resources--securemesh_site--reference--group-001.md#canonical-1012301111020231-3113010030110020-1033023300103230-3030032323030001-2221221222030330-2323313011200331-0300011230110323-3331311123022013) |
| `custom_network_config` | [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-1221311101112101-3010111123020230-0110121233010111-1231223120031232-0232321112013302-3020110211302320-2121031120023210-0031003232203110) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-0021032211102020-2131030200222300-2020022311303020-2203101233312212-3110131210011221-1323320112220321-3211130221012013-2230212111031120) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-0001313102020122-0023222332021222-2003202332110011-3111331030223330-3330213200110222-3003200233021200-0221310131032112-1330130222321123) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--securemesh_site--reference--group-001.md#canonical-3330031332202000-2113030031231101-0003021112233121-2331113233213023-0012001003203223-2223122111233031-0103323200133203-2223003101103102) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--securemesh_site--reference--group-001.md#canonical-1012222121232311-1303303010332003-3021133202331130-2221301011033233-2112220300221201-1120320201231112-2333033123300232-2100222111122011) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--securemesh_site--reference--group-001.md#canonical-2030200023320111-2301020011221220-1110232212211103-1130333102001030-3113031310321223-1310321122000323-2101232101330220-2113030022321220) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-2003211123303112-0103300310123022-1232130213230100-0312310200332100-0030033332012200-2202013010302011-0323201110002202-0101000020123103) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-3113122231312030-1201021123121232-2302013330011002-2313112202121031-0231023022321111-1100330330130103-1310300032101111-3310032213002000) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](resources--securemesh_site--reference--group-001.md#canonical-0230030102013123-1202212132220101-1221130312222331-0020132312200112-2311111001112200-0322003220122033-1302302302221110-2000312122110323) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--securemesh_site--reference--group-001.md#canonical-2231020310002301-0231122122210110-0232030303321203-2223100200300023-1012102003031122-1231001122002202-2221102003211032-1310111013010002) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--securemesh_site--reference--group-001.md#canonical-2303123212221233-3012220231323300-0131321333131321-0010223312113130-0030210301203232-2221203103030101-2003210013301212-1121103103213130) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-2113313301002330-1121203111303013-0133313302021112-0110023221130023-2231202231230210-2121212331221122-2320321132302111-0230302321230112) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](resources--securemesh_site--reference--group-001.md#canonical-1012030130033103-1113321221300112-3001312323123133-3331031201312333-0002210332002031-2131020333220223-3321001203321003-1033302003132012) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](resources--securemesh_site--reference--group-001.md#canonical-3021003330011311-1333303212001111-2212311033210032-0322000022020120-2210331323310132-2022230310111201-3030002101031112-3133231022132211) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](resources--securemesh_site--reference--group-002.md#canonical-0303201310031031-2323102001221311-3002030221031101-3110221232010120-3011303132222200-2031020103202213-3000233313230313-3320102323310231) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](resources--securemesh_site--reference--group-002.md#canonical-2010102330302122-2133321020013333-0023311313301131-3133013012022310-2313000203231010-2223102121032133-3121213101002113-3123322303113011) |
| `custom_network_config.default_config` | [custom_network_config.default_config](resources--securemesh_site--reference--group-002.md#canonical-0132110310222130-3311203300231032-3201323020201333-1202321122103233-1203330130233232-2100130003210300-1323110131102001-3332312212033011) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-3213011112013032-3310111303321210-3011002102311303-2222032110200133-2123000302231211-3311013101000330-3032331031210203-1200101200200202) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-3120010312003000-3021112032333311-3123232332031133-3020013323000023-0320020211133203-3323300311012121-3030230110232111-3203110000221320) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-0033113133211033-2200312302301331-2232032200220213-0332331032231221-3111013102222010-3120021202010333-1213203322322230-3313122332332220) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-3102110030132111-0123331011120312-2233011232332200-2320231102310003-0010002303131303-1233100000031103-3311120310001303-2233002123001201) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-2200020230022100-2032000100333202-3320322010131102-3211113300331333-1310100010211333-3113133121231212-2313300133013330-3312112110110122) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-3132131320231210-0103103230123002-0211023021122210-2230200101033332-1130101100203222-2330200133100301-0010301132000012-3333203032000132) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-1110133321120000-0102112023330021-1223302002203030-1013223112331232-2022131330233321-2223200300033221-1123220101113232-2211303233010330) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--securemesh_site--reference--group-002.md#canonical-1322013322330210-0312000331223311-2021203131013032-1302303320330220-1231222022101022-0112201023322333-1212131201123011-0013202100223223) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--securemesh_site--reference--group-002.md#canonical-2113311123223020-2120232210331312-3000310211123223-3113102232223020-1032333333100001-2200101131130323-1133333103120212-2110223222313202) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--securemesh_site--reference--group-002.md#canonical-2111003011223233-2122022213112300-1111001113103031-0032223131013311-3223311132311122-2101331122130320-1132123333032323-1230133123202131) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-1130122202111221-2101311313322203-2232022112210300-2020310233312121-2203201211230022-2311232323103210-2311131102111300-0210220020320221) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-0123230103131231-0220320002332001-1101322321131320-3103021202022210-3230202301002010-2023303012333122-1210000221002102-0333002012020003) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--securemesh_site--reference--group-002.md#canonical-0032211002133100-0321320030223013-2023021220302212-3213021312332310-1020201033332330-2320103221200001-0313231301200312-2211233102323133) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--securemesh_site--reference--group-002.md#canonical-0101030131012133-3322101031123203-3000112230332212-3133130323202300-0012121211013332-0101101323023211-1131022111023233-2001230021110020) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--securemesh_site--reference--group-002.md#canonical-3131012021301011-3032322101212033-2122311232212130-0003320331131133-1202131231131101-2120321310232031-0301313303023022-3322213330000333) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-3031231313103301-0111321211211020-0310322321322320-2210002113230301-1121212122023203-0000032233330210-1302203201113233-1103020210020100) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2032133103223310-2032211132321000-0022021112112130-1312323203332222-2121331223311332-0031212310001033-1131210001302312-1111112312223201) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](resources--securemesh_site--reference--group-002.md#canonical-1121023011212102-2011312132112320-2033210202101033-3321033133231032-1221021013322203-3000030003021222-2002010213111110-1313030100311100) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](resources--securemesh_site--reference--group-002.md#canonical-3122003012100110-1100012112003233-2313330013223110-3222102000203102-0030122200232132-3313030122120121-0210133223110033-3101311123113332) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-2013203101311331-1211130221323023-1000302320233223-3231100023000130-2033113230300323-0231301301133133-0032103303000111-1000110232323300) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-3100332220122303-1310013311312013-0111303013330300-0033331031131003-0121001122310111-0313322021301320-1331211300201110-3002102213110320) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](resources--securemesh_site--reference--group-002.md#canonical-1111323103100302-3133203303312120-1231222003113333-1213310022230320-0032313320210301-3032031310312112-2310101131113112-2031231010111011) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](resources--securemesh_site--reference--group-002.md#canonical-3000301222323323-1302031123031020-3202010120303012-1120332201002322-3102012033101210-1023122003300332-1323113032310301-0021133212123031) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](resources--securemesh_site--reference--group-002.md#canonical-3312021020220211-3330222320233310-1100110221020301-3231202201213013-2232322120232200-2031011131313021-3003302332331033-3312301123113211) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](resources--securemesh_site--reference--group-002.md#canonical-2012110200103223-0030010333123323-1302023313130023-1202010030122202-2013011111223013-3233210323132111-0230321312111310-2201113122112221) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-3032012030202202-0311003023312100-3301331332313131-2111113333212212-3231132313200323-0133322320132131-1213331001011333-0210102013231113) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](resources--securemesh_site--reference--group-002.md#canonical-3101120222122310-0231112312301010-0101211103121030-2000231100210211-0223022330130020-2221021013233132-2312203110021121-2200023310120130) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](resources--securemesh_site--reference--group-002.md#canonical-3321200320221230-3303003003200202-1231122002003020-2031213320130231-2222021223320103-0002103311112110-3030201310213131-1300132203101213) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](resources--securemesh_site--reference--group-002.md#canonical-0222101323122002-1110012223332123-2311210212311021-2100120111001202-2132010323001221-3201022120122021-3132300202222012-1303130112023122) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-3301323030000332-1201131333002310-3330220203332131-3010200300123023-0022001030202301-3131001332213323-0012302101012310-0003100202012022) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-2323233313312111-2332001212000031-2012033202230313-2131121233032213-0031121330222221-0001210011301103-3232021203230212-1130232213121032) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](resources--securemesh_site--reference--group-002.md#canonical-0323302323210231-2102222131112033-0101320021232012-0000312132123131-3302113331303320-0300013101231012-0123133333312011-3100203321003000) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-3121031200102111-1030300130031101-1312101303213013-0233332230103222-3011202233110013-1121111202013110-0002002013120022-3233122221031310) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](resources--securemesh_site--reference--group-002.md#canonical-1023321210203333-1332333231223211-2302110302200230-0231220303030103-0221123012021310-0330332230022030-3333202320123030-0232101210013233) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](resources--securemesh_site--reference--group-002.md#canonical-3230221312230332-0200232300013121-0332122200300321-3322232102032101-0010001022030331-2011120130203321-1223020003103323-2003200213120330) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1031003210123321-0013200222022231-3031103030300110-0213213302101332-3100112313233323-3032200010011031-0023010122201203-3032211002201023) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-0132131003210032-0322301220223313-3200103312230232-2123002313310202-0330233133122203-3201302021233122-3211013320110323-1300301000023000) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](resources--securemesh_site--reference--group-002.md#canonical-0122132122330103-2033200201213032-2232331020001230-3133202001133322-2213110212111022-1003320222203133-0111101130331332-3020223102212102) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](resources--securemesh_site--reference--group-002.md#canonical-3200333213221023-0000220311321133-2123030023232122-1230130203320230-1301323031231033-2320332232011200-3331132133122331-2223232301130023) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1122322021331212-1031311233110231-3032010032220012-0102103312001022-0211301022132232-0031232333002212-2212103333202123-3030130313201111) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--securemesh_site--reference--group-002.md#canonical-2101233300302011-3310303230221022-0120023323222221-0122100332120312-3010130311321322-2332011000011230-0213113121021312-3223010031200132) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--securemesh_site--reference--group-002.md#canonical-0102121110330323-1303023131033002-0310131303122233-2111100023220021-2100113011232221-0113120330200131-3121222103031201-1130022120021313) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-1130333220231003-1332233031213033-3330113001131022-1022021213210302-1100212312223220-2330121201331322-0313013010321132-0300222002111023) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](resources--securemesh_site--reference--group-002.md#canonical-0320332111002330-0233001103313032-0122013332322211-2301333003321210-0233111000220033-0103320013202212-1210300102200131-0223130313102111) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](resources--securemesh_site--reference--group-002.md#canonical-3302203233101223-2023100312030332-0131023032202123-0201200100213010-3312232311100012-0001221311311130-0223303011132333-1031202302101110) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--securemesh_site--reference--group-002.md#canonical-0131230321121231-2002302311312131-3203111110031303-0103203102201111-0310021211301012-1120301200201031-1020123122012321-1303112202220012) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--securemesh_site--reference--group-002.md#canonical-0202132222010003-0313200331313120-0111130221120300-3111112200111131-3233223011320032-1023223220223213-1131202102020223-0331133010222312) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](resources--securemesh_site--reference--group-002.md#canonical-0020013120211222-0213300222232123-1330233332323112-0001332112323310-0023130211233020-2101130011200212-0103121332333031-1301330103300321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](resources--securemesh_site--reference--group-002.md#canonical-0131320311000331-1323233300232312-0220101030320001-1203232100001132-1102001301101230-1100323303230323-2301132220110032-0301033003113321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](resources--securemesh_site--reference--group-002.md#canonical-3030213331123103-2012002003222313-0230110003333311-0132302111211031-3102220301310333-3303310213121131-2113321110233000-0132130112013022) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](resources--securemesh_site--reference--group-002.md#canonical-0312122120030101-1212003330220130-0322013130300321-1301310332332103-3032102033322303-2011302031013112-0023321331013330-0202123020332120) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](resources--securemesh_site--reference--group-002.md#canonical-3131012130201310-2312333213231121-3313020313212321-1102021331223100-1001220112332202-3320210103032303-0330221333312303-2303231130213232) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](resources--securemesh_site--reference--group-002.md#canonical-1011011133313003-1022010010012123-3120022011101101-2102021133122001-1112320121120322-0031322331231001-2132212123202000-0030331110200100) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site--reference--group-002.md#canonical-1221211020131221-1110022301031000-0222020001303110-0203110120210121-0302111303211102-0023333131333032-1102102123232301-1033111332302333) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](resources--securemesh_site--reference--group-002.md#canonical-0011031333313230-0322311121113301-0301102132032303-3013202031121303-0123321110103322-3302333013112101-0030303201023322-2123201013232232) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](resources--securemesh_site--reference--group-002.md#canonical-0311333122132032-3303033003103102-1231301200030011-3323103133013002-3222222103321313-2031212012120323-1210321223231210-2013030310220321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-0212233001123230-2121202202310120-2032300302011100-3203111322322113-3220211131221333-2201330302103311-3200333110311303-3012213330232232) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-3111021021313321-3213113012001222-3030111210111002-1211011121203120-3122220201330100-1110030133213022-3201120230231132-3122001300030320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3321203313230201-3120112203113011-0300311020111331-3301233232103322-1331013222333001-3312123211003103-3203111221112002-3022213033332123) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](resources--securemesh_site--reference--group-002.md#canonical-3322011220321130-1131210333103210-1232113033331130-2223231200300300-1131002032303022-1220000202303002-0000013211322300-2231232012121332) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-1233102012103223-0130110331103211-0123202023003203-3332022303122000-0122033020133101-3210212301133021-3131313222002200-1321032120222001) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3021332123121111-3022312301011110-2301221001131200-0201313010233203-3121212331312120-1313303320322310-0333000222221323-3222123211203100) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site--reference--group-002.md#canonical-3121033201310302-1032032112201122-0230231203131010-0121200001200322-1031132220223233-2300130021333132-1210301222323212-3220300303221232) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](resources--securemesh_site--reference--group-002.md#canonical-2221202212202133-0230123100320113-1201123103130222-1212002203111210-1231301321312230-1223320320331003-0321120002211200-2332232123033011) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-2212211130133100-0212200111310312-1132221210322010-0232013230311020-1001202001321320-1330012130000313-3221030213203231-2310021111312311) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](resources--securemesh_site--reference--group-002.md#canonical-2000312132020223-1003310323222320-1030013211111330-1230113101130232-1010001312113021-0002313021312313-0231003103021331-0313000031322133) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site--reference--group-002.md#canonical-2110230330110231-2233131100101202-2213300101322102-1322330100203230-3310330111113000-0033031330131232-3332332022201302-0132102002220321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site--reference--group-002.md#canonical-0321031132221123-0132231221330120-0030101322320221-2003031330300131-1222030003211131-2121231300013231-3311021333013221-1313300031121320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](resources--securemesh_site--reference--group-002.md#canonical-3010001202013200-0231331333020133-1310210233202112-3102230101011033-1311131321130310-2212332011300302-2022301320030212-0333211312233200) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-0112113122003223-3232120223310300-2122331202030123-1120110020323212-0210202301003021-0203212201010232-1132012121131020-3103020200332320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site--reference--group-003.md#canonical-3200323330022333-0333323130310022-3111203112120220-3333032330231333-3323001330033113-3002332321203333-3302113301131223-3223211032320122) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site--reference--group-003.md#canonical-1330031220101333-2222133233022330-2200111110313313-3030123302230321-2130332310030320-1122312131231321-1031112212311003-3031120322123302) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-0032213200201323-2300121003320303-1121020321323331-1121103131132230-1233321030201322-3123021113102110-2332101202010331-0111001120010113) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](resources--securemesh_site--reference--group-003.md#canonical-3011232122222000-2002001310012220-0322322020133012-1010203223022132-0002333312021130-2201131001312332-1210020132310333-1113013120321312) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](resources--securemesh_site--reference--group-003.md#canonical-1000000031023302-3023023210003232-2131322121112223-1201320213320311-1111000212311331-3133200003213103-0000211230332012-0121100102122031) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site--reference--group-003.md#canonical-0303212002002121-0202203000033112-2002020310201020-3011232210002100-3021130032120313-3123112110302003-1312230010203300-1201322220102323) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](resources--securemesh_site--reference--group-003.md#canonical-2012220332210200-2023011003021210-2212222202323222-3103101003032312-2131122013230010-3133310032022132-2322232032030022-2120300302210231) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](resources--securemesh_site--reference--group-003.md#canonical-2313230000013231-3033220122310200-3302330211221123-1013131012331103-0212203312022222-2312333320320212-0012122030023210-3300030111203220) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](resources--securemesh_site--reference--group-002.md#canonical-1230012333033112-2100220013102233-1113332020112013-0321120021202311-3212310112220233-0212332112122103-3211331122121102-1322103113030331) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-0203120023130303-3312323333130300-2203102130201222-1122212300323200-3130210213212111-3121322302110313-1123131203031122-1332000120030310) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-0023222100212021-0331121000300002-0031320002022232-2132301103001120-2201001100112130-2230210000211010-2232321021202211-0101211331301321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](resources--securemesh_site--reference--group-003.md#canonical-2130312313210321-2003330020330030-2302021031033201-3030300131230031-2112321023212221-2021210132213003-1131320132000133-0112203113233313) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](resources--securemesh_site--reference--group-003.md#canonical-1311011203132213-0000322233213311-2322120202302031-3303121320200132-3101012000231230-3002200102010003-3303102202201211-2232033233221303) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](resources--securemesh_site--reference--group-003.md#canonical-3323121121232201-1220323001031011-2213032033203313-0130003210101002-3022110002031211-1201021202131011-2221111202333221-1231110010230122) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-2101122022313310-2313110020003331-3212032311103300-0313212332013320-3021111221223031-3111311123201123-3322323012121211-0001112210200303) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1103211331021131-2302122210201333-1001111200322120-1333230332133002-0200321032130113-2211331012030300-2011232303122102-3013320103111333) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](resources--securemesh_site--reference--group-002.md#canonical-0221303232100311-3011011022203013-1322310332312031-0022201322033303-0122112023312000-0033320202311301-2211133133013000-2011202031032320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](resources--securemesh_site--reference--group-003.md#canonical-2113110110022022-3022102102322020-3333021231121321-3302330231223133-2210312330231012-2012331122012001-0000231220212130-3122103312111323) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](resources--securemesh_site--reference--group-002.md#canonical-0031023222123301-3102222211030221-3200323020333021-1111131110221323-3220333012303012-2103222321223021-1332013230200233-3132331101213020) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](resources--securemesh_site--reference--group-003.md#canonical-3121201221101101-1102012011000201-0031201123233212-3222133320201110-0011021233121002-1033103023313113-2301312300122021-1131013300002010) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](resources--securemesh_site--reference--group-003.md#canonical-2020332031032130-2233331023323201-0003021131220330-2322331021130311-1223012120023311-1121301112121202-0322130200320333-2110212012112000) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-2013333031130003-3021120002010000-2002131103002122-3013021002201031-3102030111222301-2211310201203130-2321000311032332-3311320321332310) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-0301302130000132-3013102232032033-0330031212211002-3303300112102220-2102301310021023-3132233120023212-2110310233000320-2202013333210100) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-2001110222202210-1230112311110022-1033033203123100-0032033100133322-1332233000002102-3113310310111112-0301301200321201-3032330330212213) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-3222122212332010-0011102003003330-1031012320000321-1210012002123112-0112321300000033-1300000103100300-2200102013231001-0320020202031102) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](resources--securemesh_site--reference--group-003.md#canonical-2211212103302200-3121210330020123-2330213232123312-0022132333232332-2201130221122333-1103320110232022-1013213102022211-0320112203330233) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](resources--securemesh_site--reference--group-003.md#canonical-3103002230223231-1132233312000132-2202100020303121-3320030033301032-3221221201130212-0211101333131122-2332123131013233-0303220130322123) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](resources--securemesh_site--reference--group-003.md#canonical-3120000330103202-3031223031200233-0322222012200211-0300002232121220-2232031320311231-1133123121200202-1032001113322030-2223013001130201) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-2131012233122132-1011222013320332-1013210031230223-2302101132021300-2210002122111020-3102221312111202-1133033212301321-2010032001311320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-1002331233313332-2032032030323222-1301313303300300-1223101131001030-0133030300232031-3300201122222112-3223320033212323-0102210201220010) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-3100231000120120-3033132333111003-1031201122223233-2232122322301212-0332320302323211-1223001022030331-0311121200013133-1300132131310121) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-2000323231121333-2133101122021032-1112021021221023-0011302130111122-3112323123231020-1233110120111011-0333001223023110-3311233000210030) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](resources--securemesh_site--reference--group-003.md#canonical-0213333312320111-1323033120202323-2302031323312002-3021132221013222-0002102132311323-0030232320233010-0313100230023311-2220300321132332) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](resources--securemesh_site--reference--group-003.md#canonical-2102231112022020-3003120032101303-3023331210301133-3213110320003211-0020130120112212-0001020131011010-0300311332001320-2323103131301202) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](resources--securemesh_site--reference--group-003.md#canonical-2320020213110321-0111112211131131-0313000320301102-2233020012232012-2011312101220312-2200223303331202-1332200303321012-2110211102221020) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](resources--securemesh_site--reference--group-003.md#canonical-2103230310230032-2233122030312032-1002103230222232-2111302120331210-0311222303311223-3230112232130231-1223212231103212-2221203013201230) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](resources--securemesh_site--reference--group-003.md#canonical-0301331133230213-2011023212033110-3133230220223000-3012300230203001-2230212013333103-0330320023121002-0103023110300131-1322102302312113) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](resources--securemesh_site--reference--group-002.md#canonical-3021202021101201-3231211133003011-1330012200031233-2300231001300223-0212020002022033-2132101113220232-2321021122022001-2132031131022030) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](resources--securemesh_site--reference--group-002.md#canonical-0122001301203001-2120331302231300-3120313223320112-2330311000031311-3210313132011131-3230300130303122-1030303300330232-3012103022021211) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-3033333101130130-0311113313032123-1203213102322021-1301013300310013-3231001310203012-0300023323321123-3000300321231211-0012022302331111) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](resources--securemesh_site--reference--group-003.md#canonical-3103312011210112-2230132221023122-2103133133233313-2022123012132200-2030110121212010-0323111332333001-0202122230121200-3131302330223210) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-1322312130131120-1113130023002232-0023120023303013-1130313200313031-2012303012322101-2102222221031010-0312010333033230-1213130233322001) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-1221312312123320-0103030110021301-3310022001112203-0213122002123300-3003333212130302-1323130020133011-0030011223030203-2133122033013013) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-1000020133210023-1302033220112300-0103002003221230-2033331113332102-0322232322010032-2203103210112221-2230122322102021-0320321010031232) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](resources--securemesh_site--reference--group-003.md#canonical-0210302000131311-2202133120332002-2030122003133131-3010032222322023-0101012212200013-3310102132222110-3112300322311103-1301203310202011) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](resources--securemesh_site--reference--group-003.md#canonical-0330022212100021-1000012003222220-2213301301222003-1212132012101002-3300023100110332-3232212332211101-2131133332010132-3210111312331221) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](resources--securemesh_site--reference--group-003.md#canonical-2312103002331133-0110122032323001-1132103311203222-0222232132220010-0332203202010111-2110301320202102-1121013232002031-3123221003130301) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](resources--securemesh_site--reference--group-003.md#canonical-1333212231110233-3313013302112231-3200223200130001-0311221332331131-1330311202013020-1230321123133301-2103230103331213-3203232212032213) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](resources--securemesh_site--reference--group-003.md#canonical-0122103111333233-1102333232122230-0202321003323120-0020300102132000-1013332021203321-3303210330311013-3312113213103302-3100033311022030) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-1112312322221032-3202020333111101-2001120331330202-2222100001010300-2223231113120321-3232031321120000-1001303121302012-1231130232320312) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-0301321111133210-3223032221022100-0032211013121203-0221122203110202-2102211201202121-2122120301100303-3230002131213112-1322112320211113) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-2033231302311222-1322210121332311-1233320111113323-1033220310112231-2320303020003310-3111131033111220-0332121223101011-1121332320002321) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0233103320310211-3131202121000330-0113032000300230-3201230232302003-3333213202023203-3323210310021120-2011211310200233-1003100220213111) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1230123031130332-1000101310130020-2202103020330220-0212120312203113-0121032132101130-0322303122230312-2031201100120211-0110331202133230) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-0323001232320220-2221322002021321-1113333003203310-1101110112321202-2012220121202231-2132203321330121-3310022013103110-1222023323330031) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-0100321213212033-1020110131211103-1021102022121303-1333023232223302-0232020210221123-1233123133333123-3232111022323321-3321313100220213) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-3333022031022220-3232010121321333-0100000223132022-0231221213221100-2232310023331202-3033131233110113-2231311310313122-1232103123100221) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-3200112122221113-0113101101321303-0122022313113102-1111120330111110-2023330131320223-2211301231010001-3100330022130232-2203303101000231) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-2131303320333211-0213221212013210-1111031302001201-2132303303222323-3011303231030021-3133102130231231-0030010132010012-3203223203023023) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-0331000312013131-1102133311322311-2200201223322211-3311202331213002-2100131132003302-3311123333103312-1312033022130122-0122321022201200) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-2030121333102222-3101030030113322-0111121230032201-3121123213313030-3202201020200213-3132222010011121-1213122300123133-3000012211323030) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-003.md#canonical-3220213303133332-0221033120110203-3302313011002210-0021331201301322-2231302230023231-2200003130133322-0123200113210333-3330331310222230) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-003.md#canonical-1221111002030132-0322311031033321-2023002333230023-1232313301221001-0330101320230312-2200101111130201-2021203303323001-0311031110310123) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-003.md#canonical-3033102300022332-0000103232012022-2101201021223201-1021011111230200-1310112131020111-1212320313312122-1101021330021020-2300120131300001) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-003.md#canonical-3122301033122312-1223122131330202-1100211030203300-2020230002120211-0200201232013110-1013322302303333-0200012313321012-0321112100111331) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-003.md#canonical-2113013230021120-1002020021110233-3313212231030212-1103300123332310-1330223311311100-1303122000032022-2230120301222322-2330201223313103) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-003.md#canonical-1323101321132201-3323211213212121-0321213101010131-0000223312301303-0202113313020201-1231321233220232-3023222320300113-3203010223103021) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2111310301211013-2301313322200033-0201321230211221-3112133120233000-0231010013111030-3301212301000101-1332022312031021-0033322003221231) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0030201022032223-3201032102310001-0212110223320130-2120000211111200-1031312333310211-3301302112310113-3133100230131210-2222213232211301) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-1012100133230021-3001121021232310-2321320320021220-1212321020103012-0211132133000300-2301013333031013-3020023320123322-0332213002310230) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-1333203222203013-3220233003321210-3121210232021000-1322013133222231-0233203323221122-3130213222232023-0223321131101330-0330332111133102) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-2000213100231103-2331000012220230-0223020032233100-1203131131002320-2233031301030022-1233130320201020-0022323213031222-3010110130303100) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-1120112320333000-1021311330200301-1221213011011231-2223111102100202-2321130200032003-0123210120300232-0032112011331233-0330110200312203) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-3200300132322113-0230030202321201-1333030103301123-3113200001102123-0133231110111033-3331312232332010-1131011102312310-1020032000222322) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-1120230030012303-2001212101220221-0022123001100111-3132132311223211-3222110120120231-0013233013000012-0000033121232121-3211001200022322) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-1222000303032110-1230102100323001-3323220213233031-3332131003012001-0133303102030231-0011122313323000-0120100200101122-2233032311100031) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-003.md#canonical-0232111311232131-1300233211032011-0001323132102120-0330323312222002-1331312111200103-0110020320112223-3232203231123331-0213111110322300) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-003.md#canonical-1201123032300310-3213123010022301-2111102020323221-3201202101231003-1332022331222131-1123131323130202-0331211102232321-3010232313230332) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-003.md#canonical-3323223202103332-0330003223130202-3010313023210221-0110231230330203-1012032311222002-2121123233310321-3331220200200232-1201120023011211) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-003.md#canonical-2031312121221103-0333201310012122-2320320111132023-3023100332313332-0112213012221002-2030300021322130-2333001010233330-0231302213120332) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-003.md#canonical-3303232002023010-2123023103131021-0311331022100113-3010212210032000-2011332313303322-2313200012203022-0211312323103012-3320011212100103) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-003.md#canonical-3030330301102000-1210020333002031-3310030110033021-3302001033031020-3330102220101102-3312212323021300-2131133000312313-3313230302231211) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](resources--securemesh_site--reference--group-003.md#canonical-2301202111113222-3001320213300220-3321320200022310-0110323321022033-1330223033101210-1212211122012222-3032222321223200-0332321103203313) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1332331002001001-1013113132220202-2132200030301021-3332011103110232-1023310301321202-0301033032212200-3121302222131300-2122311312010322) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2221310222200112-1031302213311311-2012033211320022-3010233122031103-0303332213312223-2322032330032120-2312322202310133-2301120211132230) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](resources--securemesh_site--reference--group-003.md#canonical-1121122111300133-1033201121103011-1222003331222322-0020011131010223-1013331312202110-3331222011112033-3331133320203220-0023010321211300) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](resources--securemesh_site--reference--group-003.md#canonical-0001303230121110-2131013212222232-2003300030020111-1311011023003023-0330012021101303-2302200033201332-1231120012133111-1002312032333210) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](resources--securemesh_site--reference--group-003.md#canonical-1212303002202220-1003133231301110-3300222112111121-2003230121023320-0122112031112003-2212221313112013-1022222232231002-0320213102202220) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](resources--securemesh_site--reference--group-003.md#canonical-2303212331122012-0202220223320012-1323101023222012-1312010322220332-2321113301321011-1110132312223120-3320310110211002-2222301333022022) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](resources--securemesh_site--reference--group-003.md#canonical-3110022030203000-1120223010033323-3203120022122312-0020320312132320-1220033310220000-0332023202120013-3000111123222121-1001101101100200) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2300300123211131-2002221213030310-1011233130323300-3033300112312011-2132030022331301-1131203003221111-1003301230202213-2300333313320303) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-0312222212133232-2123321103323020-3330121001122021-0001102032203021-0123321103120321-0011301312220203-0203201202110102-3303323202223010) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-2001020220310202-0332033311232313-1111003110231220-0021132100301201-0022310320121002-1011101213332111-0033023031220323-1022313031010001) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1033003123123002-1102020233203110-1313031031331002-0322323003333131-1300332303130320-3012023130030333-3010021032232231-0213030130031233) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3300301301210202-2310233101311223-0321103200232031-2133022300120131-1313020203311123-2330323022113303-0202210033020100-1310203330202000) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-3330300221000233-2333212101202312-3113230120313133-2201323121322102-1202111300302303-1322231223132103-0102103010331210-0332010120012230) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-2331333102320030-0331112222021011-3102023033130030-2302321130201123-3100112002032322-2200002211300233-3311102321311121-1221021023230111) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-2133113301212101-2020223000330200-1202223100301122-0012001310023033-3000123002220212-1320220330131303-0323233132030133-2323322300222110) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-3331020132133122-1201131023321200-2020212311022112-2222130033331311-3032311003010031-0123210000103101-0323033123000032-3301101231032113) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-1031030230221131-0110111022003123-2030003113212320-1221232301102330-1321300020332223-2303321301033310-1122201010000200-3220102032323032) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-3131312112320111-3111310011003123-3313010322303233-3323213012031320-0013232100233002-0112311010000021-3031311100202312-0012222011312211) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-3233113213331302-3123002131133100-3121322212300132-3010033321301012-2202222000221102-2201110031333302-3323013303010331-1213230223133322) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-004.md#canonical-2311232103110012-1132133330230023-3021032001122031-3021132330332333-0202012213312220-2011232122023311-2302221312211210-0101001312312132) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-004.md#canonical-0132111203001133-3001303200233332-1013123332330131-2132221003010131-3131302012001320-1322131031030321-2003312301033112-2113020131033010) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-004.md#canonical-3223303102111101-1001103211211322-3001212023223022-0010312211233132-0030223033230002-1132020010321322-0032101212200111-3010100030200220) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-004.md#canonical-0002331113211211-1111023300231012-1311322301122212-1223231001331120-2202210023302002-0300222003333330-3232322103311010-1013220333002012) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-004.md#canonical-3303233311313003-2322113120232131-1113021101213031-2011000133321210-1303210333210121-0211231222120122-2033122312032222-1001101123131200) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-004.md#canonical-3000010130330231-2321231311010322-0001200302211231-3033023321010300-3333012310331010-1033321311310031-2132113220022130-1011010033320111) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2221001023013331-3121023202313101-0230313310210103-0332123220003022-1122202212130132-3232121131333032-3202213213333123-1302112330322131) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-0332000210122313-1210231033230200-1002002221010202-0201030111300330-2011122231231331-3221301113120011-1020103012230013-1131233223310031) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--reference--group-004.md#canonical-3210222332312121-3033222321320301-1121021022310133-2033132211102121-3000321003200032-1301030222032130-3233302100210022-0000100221003131) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-1300233333331023-2210303110323123-2221120202302012-1111021123103212-3002231100123113-2322203333130032-1333201220212022-2131232330022133) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--reference--group-004.md#canonical-2230230033300310-2332321211203211-0010123200311132-0212110033010001-2023333232231110-1220113310303322-3120133302131331-1111111301112010) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-004.md#canonical-2310323310302111-1003030223122331-2120200111211303-0313223010231310-3101113210223122-2031300323220331-3030301222303033-2321301021021002) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-0301013031230013-0221300011223232-3131212002000132-1201023100313300-0311132222300213-0211311122223032-2033333201012032-1333032123002020) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-2213002010113230-1203011133313113-3212103123002003-3120322210230110-1122122030233120-3322123101330023-1300130231131031-3012201101330122) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-2212130010022032-3313120010132213-2223021220132202-1201332332112202-1132321302211331-3100103122313131-0110120330120131-0133311211303100) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-004.md#canonical-3331213020100011-2302212000003022-2102331103211032-0113113300031202-2321121131132210-2330133130012103-0100200032012210-1210000202221123) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-004.md#canonical-3103231232113200-2231100200313311-1030123010011203-3013033032010312-1011131023102102-2300332130333301-1310100010233111-0320130311003021) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-004.md#canonical-1223121103312310-3201223100320102-0201321113133330-2232133212202221-3330030211330202-1301121132032322-3223111112030320-0010110033132020) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-004.md#canonical-2102212033103322-0323312321333333-0321222313010330-2321203220101322-2021310111300133-3110301131003210-2300320110013213-3120223220220221) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-004.md#canonical-3123222123313223-2302021031111220-0320213320230332-2013101211220311-3000200120312220-2330320101201033-1211303333130033-0110121213300333) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-004.md#canonical-1023332213033312-0300121113102002-3013001331000132-2221101312331022-3201310200223033-3013203021120202-1131121022202033-2222101132210113) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](resources--securemesh_site--reference--group-003.md#canonical-1121110122012012-0222302320000323-1032302212033230-1011321213010311-1133323232222211-0023312311222312-2111103022301331-3221000013310211) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-3300333211213131-2330032100010033-3120000303101223-0012220031333100-1222320113113213-2332001101001011-1231321201333110-2303211031333323) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-3121010132212233-2123300123202210-1313310001122332-2113313002320223-0013312211213301-3001230112122220-0033233133123200-3303203012202310) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](resources--securemesh_site--reference--group-001.md#canonical-2001331121220031-0332022000321121-1330022013221020-3200011231302211-2332002223230301-0231232221111210-0122323013200300-3033031121020133) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](resources--securemesh_site--reference--group-001.md#canonical-2223011101113110-2131120301113221-3102020332201023-1311123311211303-1131013012130033-2032122101010313-0232110320010013-3103002203131113) |
| `default_blocked_services` | [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-1303223332233202-3102111200020010-1101221113320021-3021211312123002-0022331320030330-2201001130120012-2332202001033010-0222100120332130) |
| `default_network_config` | [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-2013013131311232-2133302221020332-3112133011223112-1223202233333032-3300330102101202-3320111211322111-0320131312220332-1201332032332311) |
| `description` | [description](resources--securemesh_site--reference--group-001.md#canonical-2023211221311000-0102123210113130-3133323112001132-2112202203121011-2112032223200003-3020320020111101-1313221020001302-1320330333023012) |
| `disable` | [disable](resources--securemesh_site--reference--group-001.md#canonical-1302322210013111-0133321020313031-0310010110032011-0223103020310201-0022330021003031-1210210312030022-1010011323302122-1201032113130331) |
| `id` | [ID](resources--securemesh_site--reference--group-001.md#canonical-1102213121023221-2032300332101100-3023020011222000-3101023333032112-0032312031213211-0322033132022033-0301202101322333-1212003202013113) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-1200220010220323-1103231332321032-2113132203330200-2203233303031201-0212130320011230-3320320203323220-0033002002202010-1022113332333022) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-1101033033232300-1323000102331313-3210313112302012-2210230210313321-0003212132322001-0100103231302000-3313213000111020-0001012032112330) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-0002223213033113-2113031212023303-1000232021300001-3033001100132300-1023231213231101-3032301301132321-2100020323313013-3003122211120320) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-1222031100203321-1211301220131003-1211320313003011-3323013202223020-3230120022121111-1012320010222110-1002222202110001-0320332233103013) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--securemesh_site--reference--group-004.md#canonical-3330311032310010-0003122302023322-1121030031100022-1210123012203332-1232211103320233-0200221000202003-2233010131213101-2121303023111211) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--securemesh_site--reference--group-004.md#canonical-2203313233203013-1211333112112221-2322002102210130-1210312123313111-2132232130231330-3021023302311212-1232022301122211-0221310303332212) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--securemesh_site--reference--group-004.md#canonical-0112333212002022-0213230212102002-3221303102222032-3221131330300033-0132021033201011-2021011001230002-1202100321311101-2211213321321121) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-3131300113101223-2112212220111302-3203030302233303-1210310113233133-3130320210223010-2031030323201331-3203103131130002-2102032011100120) |
| `labels` | [labels](resources--securemesh_site--reference--group-001.md#canonical-3203231321003101-1323220303111222-3013200222311121-0320213132100030-1013312222312220-3033230132001201-0033302331211031-0113220312012301) |
| `log_receiver` | [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-3003100221222100-0203310301302221-0311123203310230-3223013221302213-0203133020303031-3121000331221231-2332200211020200-1130233321233032) |
| `log_receiver.name` | [log_receiver.name](resources--securemesh_site--reference--group-004.md#canonical-2020000112203102-0130023310330131-0232221103030332-3023013033321023-0330112220221112-3313113320300022-1321023100012313-1010001300312030) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--securemesh_site--reference--group-004.md#canonical-1011010000311020-0033222332122130-3022311221301212-1111023310021210-0113033112303020-2123210201122313-1031031110002003-1133101031133021) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--securemesh_site--reference--group-004.md#canonical-1022103111022303-0302310103121232-1321323030210222-2323122101320211-1301323332132300-1200231110231101-2122202122010033-0121023203031322) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-0122012011201212-3003010312303213-0323233010133011-1321302030301332-3321311103110011-3210001001022311-3203201331001001-1120311203020313) |
| `master_node_configuration` | [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-2130211113130203-0203010311303120-2213000303013002-0222132121111333-0330131130030031-0122010020121120-1320203210233323-3332233211022331) |
| `master_node_configuration.name` | [master_node_configuration.name](resources--securemesh_site--reference--group-004.md#canonical-2012133132133233-3122030223223023-0102301031103232-0322112023000000-2030130302312012-1210310331312001-0102132132121132-2100012102112033) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](resources--securemesh_site--reference--group-004.md#canonical-0031102331331311-1112021130321120-0021313130110032-2310232122032313-3023123210013131-1331330000123101-2023033001320320-1010320100011221) |
| `name` | [name](resources--securemesh_site--reference--group-001.md#canonical-3213003332010010-3133230111231132-0023221012011322-2100320131313120-2211123100012013-3333313112110202-1302000232223112-2033230310012331) |
| `namespace` | [namespace](resources--securemesh_site--reference--group-001.md#canonical-0010133111201002-2032112111030133-1122331020213130-2302321302030200-0032010333320223-3011230320021231-2220003230330223-1112111021320011) |
| `no_bond_devices` | [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-0113102101220123-2031110131133030-1302213332233013-2303211033233232-3233310230201330-3110102232321022-3310032031233310-1231321301312313) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-0122021021232023-1111302211230021-2111220301121202-0200123220133110-0110232302312300-0020212332331122-3103200101012013-0132233322012210) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-3100103200221120-3001022213312012-1313120222111213-1230322112023313-3131220211132200-0011330321020331-2111222130201330-0122100320013123) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-3113330332210203-0122313200031100-1111113332223312-2211000201133312-0032212012332210-3102323331211322-1031103331301131-0113122023311200) |
| `os` | [os](resources--securemesh_site--reference--group-004.md#canonical-1322303212201302-3130212320301231-0113001101122220-1003031021203100-2220033322332223-0023132320030010-1211131201210110-2102132300211012) |
| `os.default_os_version` | [os.default_os_version](resources--securemesh_site--reference--group-004.md#canonical-3031210222033031-0112020002333111-2102132210011321-1220213331332311-0212103212330001-1003010210113123-2222302331010130-2231202223022033) |
| `os.operating_system_version` | [os.operating_system_version](resources--securemesh_site--reference--group-004.md#canonical-0232310003231232-1103300210021021-3022201321301113-1232230302313321-1232231323221032-1212330223331200-0310200103133112-1220110103122232) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0200201230210303-1312310131231003-0131311210012310-3223120003223203-3112102211300332-3000211023302211-1122020102231101-1003133121023302) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-1331133330203133-2230131130221321-1323000002212112-3121331301030301-1330111123020010-2031013122233313-3332332233000301-0111121212000332) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site--reference--group-004.md#canonical-3210230031131113-3223233212002002-0103311323111111-1130010001130231-2120132033020332-3123113231321031-0120000301103112-2201133010201100) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-0031023110120211-0200100020310201-2221010121133010-1222022123110330-3112130133311102-1211320013210022-2021303102112013-1213300203002310) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-1013301130032023-0222200031001211-0222303310113010-2201110033331202-1212330221230000-1130033032220321-2000030111301000-2232332232213022) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-3222310222320011-3202320000202203-3312332010020321-1122202301122133-0210102322300222-1100030122101011-0000002113202203-3011012012123211) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-0132111022322021-2002211322301322-2103112330121223-3031103313110011-3330122200313031-0321233222130212-1131013031033011-1033011220320001) |
| `sw` | [sw](resources--securemesh_site--reference--group-004.md#canonical-3223310122101302-0103131020331231-0201231112030130-3130013311322202-2000121210232312-1133210323321010-2320311023031101-1030210223333112) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-1102033032300020-1120013312333021-3103231023002121-1322213122033321-2221320122111032-0301003011130011-1202023301003211-3101023322211312) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--securemesh_site--reference--group-004.md#canonical-1332123030321022-1111110220313130-1122312000200103-0233313210033012-0001013301111010-2200300232011210-1013032122123302-2201313130121013) |
| `timeouts` | [timeouts](resources--securemesh_site--reference--group-004.md#canonical-0330122001113023-1322112322032123-0131023302203121-2210131202331232-0020311311332023-2301222310022112-1202303321003130-3003031113212310) |
| `timeouts.create` | [timeouts.create](resources--securemesh_site--reference--group-004.md#canonical-0233102122122003-3220203011011022-2111033110002010-3130213021103221-0322213102233011-1100112320322203-0103132201001020-0132223032310131) |
| `timeouts.delete` | [timeouts.delete](resources--securemesh_site--reference--group-004.md#canonical-2101112113221302-2312230030133002-2211211032002113-3013200021111032-2031121200033303-1120113330122001-1330023121112100-3223100133301000) |
| `timeouts.read` | [timeouts.read](resources--securemesh_site--reference--group-004.md#canonical-2213310131322121-1231022303222220-3010003323321103-0010303112122202-0322021321203331-3310102012331100-2331313232212013-1303110131323122) |
| `timeouts.update` | [timeouts.update](resources--securemesh_site--reference--group-004.md#canonical-3210202220122102-3120101303101210-0000101123102300-1220000321321331-2021000200300200-3120002320200100-2203311232001123-1312110132022312) |
| `volterra_certified_hw` | [volterra_certified_hw](resources--securemesh_site--reference--group-001.md#canonical-2201331321112131-2223032000131211-0031230323101122-2001022033121020-2021112031300120-3202221013313332-1131112323020110-3103110303311032) |
| `waf_signatures` | [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-1333331221112111-1230332223010200-1203221320210122-2111211122200321-3021230220001210-3132232112311312-0100312311231331-0000012200301310) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--securemesh_site--reference--group-004.md#canonical-3200220232003323-0003021002223332-1233031032022010-3131103322022310-2110223332132100-0200113312031002-3101103022302303-0030223102002320) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--securemesh_site--reference--group-004.md#canonical-2133132112202113-3222122211021301-1022132222111203-2212210111001332-3311123303100321-3210200311001013-2120231303000212-3002031110012222) |
| `worker_nodes` | [worker_nodes](resources--securemesh_site--reference--group-001.md#canonical-1220330132001101-3301000301213310-3232121012110020-1321011320000203-0121300200232311-0323213212203122-2320111301102230-0130130310023122) |

<a id="canonical-2031212010223203-3322320322202102-1202312033221122-3031032033030031-0313101211233000-3300230101000003-2303121311021212-0030020021212200"></a>

## Next pages — Property reference / 103230111112 / 15

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232)
- [coordinates](resources--securemesh_site--reference--group-001.md#canonical-1222331220131320-1233331310213200-1221203223323001-1313113213210333-2110020133102120-2131211230222000-2022131203310321-3311113111312312)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-2130322011122113-2023213101220320-0310300323120322-1213333313311002-3123113231010313-0110102302312320-1200031302010203-0311320121030301)
- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-1211212023033021-3013201123022023-0023321103121133-1222322023300323-0131132032303333-3332330303301122-2032220120223122-2021102112323110)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-1210231003210032-0210302031033031-1321120333323332-0113030213300300-3003201313113111-2321333000300322-1312122112223131-2303302010130323)
- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-0301322311211110-3313100231331120-3020131300021210-0031233031211302-0223101202102010-3322120202121023-2112032200010220-2203211233300002)
- [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-2030021313103121-0101111110133002-2130111211310100-2020102210122203-3100100030322221-2033211312031111-3022033203312230-3201003223312122)
- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-0333112101031113-3121130300222310-2130011201332013-2012000233233023-2231013313212021-1321312222230212-0323220201313111-3200133332031330)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100)
- [os](resources--securemesh_site--reference--group-004.md#canonical-0102133232320131-2231010312223023-0203311201032131-2210230021210322-0310110201031312-2321113100212222-0201202312033330-3000120030223322)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [sw](resources--securemesh_site--reference--group-004.md#canonical-1000012000123120-2210011233010203-3333022101301022-1133130210233321-1101033113332231-1333113032202102-2032212300223010-2110321233023213)
- [timeouts](resources--securemesh_site--reference--group-004.md#canonical-3231302133333333-3210323010223133-3130032303201002-1121332221230320-2013212123011031-2102312110221123-1333231321102102-3213311103101232)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102212222221130-2021233300000233-0001333200113332-0010021311310320-2001323120310103-0032111310230003-1210232003022301-0033120131321022"></a>

## blocked_services — blocked_services / 132031123302 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- blocked_services

<a id="canonical-1323331333032130-1212021222032303-1231200023110010-1223330121132013-3332113300121023-0203313201001010-3121311032130332-3100313123123303"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: blocked\_services, default\_blocked\_services; Default: default\_blocked\_services\]
Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-1323331333032130-1212021222032303-1231200023110010-1223330121132013-3332113300121023-0203313201001010-3121311032130332-3100313123123303)
- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-1303223332233202-3102111200020010-1101221113320021-3021211312123002-0022331320030330-2201001130120012-2332202001033010-0222100120332130)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103331130121332-3201310300232203-3213102332210022-1301121033202313-3331221223112133-3212213002130002-3033323201322002-0323303011102233"></a>

## Direct properties — blocked_services / 132031123302 / 3

- [blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223): complete subsection reference.

<a id="canonical-0231031322030203-3112032010102202-0212100331222221-1230013131032022-1221003012102232-0322203023333231-1012122012121133-3000311312111033"></a>

## Next pages — blocked_services / 132031123302 / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220201013122102-3131021113033000-1301012303122311-2213223100112102-1032330103333210-2303020130311212-0213010303000122-1112123030111311"></a>

## blocked_services.blocked_service — blocked_service / 010323313030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- blocked_services.blocked_service

<a id="canonical-0330011031130302-3233323031112113-3102033333121013-3010102123102302-2132201312331133-1222021231133002-3022210223032100-1320010323312032"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131202202200213-1303331000333002-1213130220010200-0321031133132030-0312231010200310-0300021031122032-2001000203222033-3220212120212023"></a>

## Direct properties — blocked_service / 010323313030 / 3

- [dns](resources--securemesh_site--reference--group-001.md#canonical-2021122001323320-0002212310010231-1212233322333201-3333312010323031-1312201330313332-2333102120330012-1120112303031101-2133233310300301): complete subsection reference.

<a id="canonical-2302100023010120-3130333210131211-1233300103321133-0213002231030120-1103121013220010-1122223010030300-2113101100202231-2120221312111030"></a>

<a id="canonical-0000313203031232-0202021221002222-3321021021012211-3100123303320332-3013320212122301-0321200200320321-0231001303121101-1312322113320232"></a>

## network_type property — blocked_service / 010323313030 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--securemesh_site--reference--group-001.md#canonical-0013103301301213-0120311013023020-2331001023323220-1212103211321220-3012022330000220-2203310110322113-2103103303233000-1232111213210111): complete subsection reference.

- [web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-0103302313130002-0102120211202113-2212131220002211-1311123213013103-3213311033101121-0002311302131300-2113302212012023-0203133330011222): complete subsection reference.

<a id="canonical-3000020212010012-3202120203001132-2323202112222233-0022303203202311-1111030221220231-3320220113031001-2030330133230203-1100113010233311"></a>

## Next pages — blocked_service / 010323313030 / 5

- [blocked_services.blocked_service.dns](resources--securemesh_site--reference--group-001.md#canonical-2021122001323320-0002212310010231-1212233322333201-3333312010323031-1312201330313332-2333102120330012-1120112303031101-2133233310300301)
- [blocked_services.blocked_service.ssh](resources--securemesh_site--reference--group-001.md#canonical-0013103301301213-0120311013023020-2331001023323220-1212103211321220-3012022330000220-2203310110322113-2103103303233000-1232111213210111)
- [blocked_services.blocked_service.web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-0103302313130002-0102120211202113-2212131220002211-1311123213013103-3213311033101121-0002311302131300-2113302212012023-0203133330011222)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2021122001323320-0002212310010231-1212233322333201-3333312010323031-1312201330313332-2333102120330012-1120112303031101-2133233310300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202121323222031-0320003113202200-0201132011211110-2300123213312103-3322020302103030-0102003322301122-2031000233320212-2002233321222110"></a>

## blocked_services.blocked_service.dns — dns / 211312333122 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- blocked_services.blocked_service.dns

<a id="canonical-0113110303022303-2203030323221213-2003233121310223-0023030001010010-1203010023010333-2023110031100332-3023112021213231-1321022200123112"></a>

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
dns = {}
```

<a id="canonical-2332213111322230-0111123322103101-1331333110123320-0011133222030313-2123220322220221-0023300200333122-1211232332322122-2122233131302331"></a>

## Direct properties — dns / 211312333122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203111312113331-3103013301012103-1133013100230211-0110322303131230-2102320010223211-1111210011323300-0233123120010002-2302110033321200"></a>

## Next pages — dns / 211312333122 / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0013103301301213-0120311013023020-2331001023323220-1212103211321220-3012022330000220-2203310110322113-2103103303233000-1232111213210111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331322102312133-1202103012310233-2210010310222103-3222000231121333-2001201120111322-0200302322202303-1330220132201033-2220032121122223"></a>

## blocked_services.blocked_service.ssh — ssh / 023212110123 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- blocked_services.blocked_service.ssh

<a id="canonical-0021010212112212-3331101013231022-0123033331001212-0032101223003132-3201302230103303-1132310221210121-3031212201103030-2022320220213122"></a>

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
ssh = {}
```

<a id="canonical-2202220230102330-2320330320111130-3131102120301033-2301311011211011-2003020211313211-3023020032233100-0223120312230133-1300212322132000"></a>

## Direct properties — ssh / 023212110123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323221001101323-1133102201231001-3330002330011331-1001033321000132-0220011202322230-1330100013020311-1203330102130131-2123323120200223"></a>

## Next pages — ssh / 023212110123 / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0103302313130002-0102120211202113-2212131220002211-1311123213013103-3213311033101121-0002311302131300-2113302212012023-0203133330011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032313131311122-0332110011201020-1022332010322023-3130023320311300-1002303201332131-1032210133233130-0301311130201133-0113111021111002"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 230202203030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0330033130303123-1003100130333232-2221022303222022-1022011211331033-2332222022032021-1213212200210103-2322220321132011-3233131222211210"></a>

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
web_user_interface = {}
```

<a id="canonical-0130323203333202-0311332311200020-2231311021211320-2022310203323023-1011123220013310-2231303100312302-3001330023021031-2022100101233222"></a>

## Direct properties — web_user_interface / 230202203030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031003131310203-1011303133131223-0103020100111033-3130311200222100-0303103322222103-2020001131000032-3211231210133322-2210301323032023"></a>

## Next pages — web_user_interface / 230202203030 / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3223132010333233-0331013311123332-3103313103100120-1222233031130131-1323233000200111-2202321110332200-0030211220313033-3331111231201223)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121303011331011-3032102123013231-2113300203202121-2202331213100233-3313330202321231-0213010113232333-1321121121323133-0323212300210220"></a>

## bond_device_list — bond_device_list / 321313123133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- bond_device_list

<a id="canonical-2200121111110131-1232233011211322-2200022100310221-3010233332001312-0230321231201201-2003033310002110-1113333310122311-0211303103120030"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
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

- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2200121111110131-1232233011211322-2200022100310221-3010233332001312-0230321231201201-2003033310002110-1113333310122311-0211303103120030)
- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-0113102101220123-2031110131133030-1302213332233013-2303211033233232-3233310230201330-3110102232321022-3310032031233310-1231321301312313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300303301120301-2213330130233313-1312013012133322-2133011102302132-0133321122103012-3020103322120022-2122030022021220-3303321301000023"></a>

## Direct properties — bond_device_list / 321313123133 / 3

- [bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200): complete subsection reference.

<a id="canonical-1101303023200332-0220011112000012-2133131300011021-0120301101312311-0322123210333201-1023231122123020-3321031002313123-3022330300001310"></a>

## Next pages — bond_device_list / 321313123133 / 4

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012101333320222-2112111213310131-1212020012221332-0111030302312301-1231210101213102-1211032332320313-1022110030333233-2330111013001132"></a>

## bond_device_list.bond_devices — bond_devices / 230012002333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232)
- bond_device_list.bond_devices

<a id="canonical-3002022202303211-1000130111031023-1233132202111210-1330202322233210-3212202113332210-3011200111211131-3221212100103303-3301232132320203"></a>

Type: `"object"`. list nested block, Optional.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingListObjectAttributes("active_backup",
    "lacp")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
bond_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133232202311301-0003312201223230-2030121303112300-1332332233310011-2123021311100033-2033132130303111-0003121030301013-0320110212133323"></a>

## Direct properties — bond_devices / 230012002333 / 3

- [active_backup](resources--securemesh_site--reference--group-001.md#canonical-3230002130021211-1230122212100121-3112010002021102-2301002201313123-1331221220111200-0211303031220302-1030011010332121-0131132010223022): complete subsection reference.

<a id="canonical-3300202223202001-1122012313201200-3031322000321111-1201322203221201-1101012132103113-2200223011110202-3222211312103012-2033032000112233"></a>

<a id="canonical-1101111221112112-3020031122032323-0320110311202012-3231310302100000-2202033123330312-3131022301332223-0030301330331322-0032121300023100"></a>

## devices property — bond_devices / 230012002333 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site--reference--group-001.md#canonical-2130022013330121-0121222312233022-3032200323232120-1301002202301202-0032301220312010-3102101030120311-3210033133320221-3101012023011311): complete subsection reference.

<a id="canonical-3110133020122311-1211301023211121-3331022120310121-2200222132211230-0012312212103130-0333313132300012-1013121102100321-2022110001333200"></a>

<a id="canonical-2211112311111112-3310010003122231-1001012102101100-1011211210010100-2122221310000101-2213322322012010-1133012233311222-1210210300112232"></a>

## link_polling_interval property — bond_devices / 230012002333 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1223222333323002-0013021331201222-0133311101102332-0021322302110312-0130310203003101-0232211003203310-0121032221333312-3020233111113132"></a>

<a id="canonical-0002302022231021-0202101233302102-0333133022320321-3201211133122230-3213302110021212-0120102112313112-3322222302122231-3221330011101310"></a>

## link_up_delay property — bond_devices / 230012002333 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-3132211300333023-1313000201333032-0211320231131102-2222013320100203-3300322011133100-2110003113030223-3031320202200010-0110220100130300"></a>

<a id="canonical-0133021003320103-2332121111010201-0122033101231331-0313203021011201-0021211122003123-2123002110013131-1101332231020013-2230121202010220"></a>

## name property — bond_devices / 230012002333 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3301310032222001-3123320010011230-2210031100131012-1103201010213302-1023331031211222-2010122112311122-0212031130233201-1131012021311201"></a>

## Next pages — bond_devices / 230012002333 / 8

- [bond_device_list.bond_devices.active_backup](resources--securemesh_site--reference--group-001.md#canonical-3230002130021211-1230122212100121-3112010002021102-2301002201313123-1331221220111200-0211303031220302-1030011010332121-0131132010223022)
- [bond_device_list.bond_devices.lacp](resources--securemesh_site--reference--group-001.md#canonical-2130022013330121-0121222312233022-3032200323232120-1301002202301202-0032301220312010-3102101030120311-3210033133320221-3101012023011311)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3230002130021211-1230122212100121-3112010002021102-2301002201313123-1331221220111200-0211303031220302-1030011010332121-0131132010223022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103231001312130-0131003000212001-1212000201203021-0320333033310300-3222330231121333-1033221033120323-3130011100220330-0110001123131313"></a>

## bond_device_list.bond_devices.active_backup — active_backup / 132320322330 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232)
- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200)
- bond_device_list.bond_devices.active_backup

<a id="canonical-1013010322301133-2213221213300231-0020123112203103-3322303112321123-3022203211023020-0021211121122113-1300322032112302-3102312102300323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-0021113313320220-0030021233130100-3113012320213033-1301023210020122-3211103231020030-2033111011210220-1030122321220012-3303220012310322"></a>

## Direct properties — active_backup / 132320322330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323020320313331-0131111032212220-1023000120310002-3310300001310033-3322211303311311-0002203332110012-0022122131022022-1322210323323103"></a>

## Next pages — active_backup / 132320322330 / 4

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2130022013330121-0121222312233022-3032200323232120-1301002202301202-0032301220312010-3102101030120311-3210033133320221-3101012023011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231311321333021-3002123101101233-3102021030133022-1303202212202203-0200202223001033-0013122330312232-1010312021123111-0131232222102112"></a>

## bond_device_list.bond_devices.lacp — lacp / 203001322211 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232)
- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200)
- bond_device_list.bond_devices.lacp

<a id="canonical-1311311100112200-0010303010312131-2013331001231102-3223321001222031-2022230211320212-2300013023323232-2130113120203302-2011300123222230"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300131230110131-0030202331033303-1101031310221331-0121223013232311-1013030322102312-1221231101003330-3332023012220223-0010222220110032"></a>

## Direct properties — lacp / 203001322211 / 3

<a id="canonical-1113311310223120-3223002013233131-3313101030033211-1323311000233021-0022301313220032-2030100132310022-1303011013120213-2330031222232322"></a>

<a id="canonical-1230011003131330-3022203113000331-0310233011100200-3231211230310121-0113112212010322-0300013123221323-0221300031131002-2311110232313302"></a>

## rate property — lacp / 203001322211 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1331100223023333-1023113012322113-2312120032120230-0202212030331230-1111010033112213-2032013121332000-3313110003101322-1032313223032233"></a>

## Next pages — lacp / 203001322211 / 5

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-2212332322110213-1123101002212023-1022123223203330-0212132232120110-3203102332130310-2301101333131011-0032233030021313-0331310110202200)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1222331220131320-1233331310213200-1221203223323001-1313113213210333-2110020133102120-2131211230222000-2022131203310321-3311113111312312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030002313001303-0213223321202232-0231130103011303-0230032333123032-1033230213023011-3112223101333332-2010330032012223-1222001121323033"></a>

## coordinates — coordinates / 013100232322 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- coordinates

<a id="canonical-2322032123200012-1300021322020212-0323232000212132-1133123020330113-2020310330010010-3221200210201201-1330002012213220-2331312022020113"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013102122101310-2223121322022332-0001131311002020-1301321133333201-3331012311123002-1301130022132231-0220210321022313-0321110131321313"></a>

## Direct properties — coordinates / 013100232322 / 3

<a id="canonical-3132202011221022-0322010002101201-2113123100300221-2332122111220333-3201301132202201-1031000131302300-1003100203002111-1223021113132330"></a>

<a id="canonical-1303023212031231-3232112110311221-1111000212001333-3301333103200210-3010202302021103-0222100123013301-2312213030010311-0123311103020102"></a>

## latitude property — coordinates / 013100232322 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-1012301111020231-3113010030110020-1033023300103230-3030032323030001-2221221222030330-2323313011200331-0300011230110323-3331311123022013"></a>

<a id="canonical-0301123020333223-0211013210110132-3223313000210210-2210121301232022-1313233023130332-2211122021131323-3020002000233023-0130323303211311"></a>

## longitude property — coordinates / 013100232322 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-2102210221300101-2121001200303301-3123203122213101-3210022323212110-2103030001233233-2123330310302302-3211203321010032-0111100231033311"></a>

## Next pages — coordinates / 013100232322 / 6

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301300121230220-1033300330101221-2200321012132010-1320022111211101-3313333123133022-3103330210021001-0010303130211033-3310132133221331"></a>

## custom_network_config — custom_network_config / 131121111133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- custom_network_config

<a id="canonical-1221311101112101-3010111123020230-0110121233010111-1231223120031232-0232321112013302-3020110211302320-2121031120023210-0031003232203110"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
SmsNetworkConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("default_config",
    "slo_config"),
  validators.ConflictingObjectAttributes("default_interface_config",
    "interface_list"),
  validators.ConflictingObjectAttributes("default_sli_config",
    "sli_config"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-interface_choice": "[\"default_interface_config\",\"interface_list\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

OneOf alternatives in this subsection:

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-1221311101112101-3010111123020230-0110121233010111-1231223120031232-0232321112013302-3020110211302320-2121031120023210-0031003232203110)
- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-2013013131311232-2133302221020332-3112133011223112-1223202233333032-3300330102101202-3320111211322111-0320131312220332-1201332032332311)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_network_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133211303001013-0120310332110310-0230120113023130-2130222133330220-3102022301032230-1022220023112100-1111100201330333-3233013002132323"></a>

## Direct properties — custom_network_config / 131121111133 / 3

- [active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-2002031330322012-1012233302212112-3013233021321320-0300100301002202-1321312102002212-0233302121100230-0001020100102311-1110011323111002): complete subsection reference.

- [active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-2201321013032120-0211131001321111-3110020111200102-1101220033300100-1300302002133120-0011321112312011-3010122313323122-0031322212023203): complete subsection reference.

- [active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-3023131331003221-2030010320001301-1121003103323033-2220220101231222-3320232031203110-3232001230223213-1002021211312311-3003233011102220): complete subsection reference.

- [default_config](resources--securemesh_site--reference--group-002.md#canonical-1003133003322303-2221311020000231-2132020123222101-2023312312123323-0130023221221313-2230022323203032-0201303102000312-2200122002000332): complete subsection reference.

- [default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-3022201322330102-2133020001321012-0320312322210022-3212030112002322-1110002223002100-1020111001131323-2132012322101210-0121001203020112): complete subsection reference.

- [default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-2221000332132000-3211222221213320-0312003330022103-3233322233133113-0012301330310133-0012213130011033-0113021130103333-0222302012203322): complete subsection reference.

- [forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-1032010001122222-1030130223330222-0330032002313232-1123210200131223-3232031003302033-3223312313301133-1013323200313311-2120020330132101): complete subsection reference.

- [global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203): complete subsection reference.

- [interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131): complete subsection reference.

- [no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-2303112131223200-0330102000132110-3111303323230203-0033123201010303-1032300122232203-0301333310233012-0102331232312221-0331010221101021): complete subsection reference.

- [no_global_network](resources--securemesh_site--reference--group-003.md#canonical-1321020212311300-1111231333313020-3332223223121213-3103332312322312-1011030011231203-3212022023111201-3230122000201132-0311311333310132): complete subsection reference.

- [no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-1001313103113102-1000033221100312-2111311202310323-2312033332301000-0123101211223010-1023230323131023-2220113220032212-3113332130103210): complete subsection reference.

- [sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202): complete subsection reference.

- [slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-2131101313011023-3202103200013000-3323310203031110-2210320233112330-0202233231312011-1312202132031201-0231331001012033-3000122001001103): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-2201302222100233-1133301323301033-1023210103301301-1330322123030131-2202111223230332-1121203213021131-1100102123030200-2000030033123120): complete subsection reference.

<a id="canonical-2001331121220031-0332022000321121-1330022013221020-3200011231302211-2332002223230301-0231232221111210-0122323013200300-3033031121020133"></a>

<a id="canonical-3310200301320011-0110130201120203-2322010211330223-2211111121313031-3202112310210022-1122021311323202-0003022303303231-3122232032021303"></a>

## tunnel_dead_timeout property — custom_network_config / 131121111133 / 4

Type: `"number"`. Optional.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 180000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="canonical-2223011101113110-2131120301113221-3102020332201023-1311123311211303-1131013012130033-2032122101010313-0232110320010013-3103002203131113"></a>

<a id="canonical-3233300102032331-0310130300020100-2003000300221203-0331023012232021-0303121121102202-2311002312010020-3012221132031311-2203122300011123"></a>

## vip_vrrp_mode property — custom_network_config / 131121111133 / 5

Type: `"string"`. Optional.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312130202021231-2000303210020023-2310011032103311-0131333232023300-2203333313231211-3302232022230000-2133000022013002-1332332130322221"></a>

## Next pages — custom_network_config / 131121111133 / 6

- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-2002031330322012-1012233302212112-3013233021321320-0300100301002202-1321312102002212-0233302121100230-0001020100102311-1110011323111002)
- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-2201321013032120-0211131001321111-3110020111200102-1101220033300100-1300302002133120-0011321112312011-3010122313323122-0031322212023203)
- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-3023131331003221-2030010320001301-1121003103323033-2220220101231222-3320232031203110-3232001230223213-1002021211312311-3003233011102220)
- [custom_network_config.default_config](resources--securemesh_site--reference--group-002.md#canonical-1003133003322303-2221311020000231-2132020123222101-2023312312123323-0130023221221313-2230022323203032-0201303102000312-2200122002000332)
- [custom_network_config.default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-3022201322330102-2133020001321012-0320312322210022-3212030112002322-1110002223002100-1020111001131323-2132012322101210-0121001203020112)
- [custom_network_config.default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-2221000332132000-3211222221213320-0312003330022103-3233322233133113-0012301330310133-0012213130011033-0113021130103333-0222302012203322)
- [custom_network_config.forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-1032010001122222-1030130223330222-0330032002313232-1123210200131223-3232031003302033-3223312313301133-1013323200313311-2120020330132101)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-2303112131223200-0330102000132110-3111303323230203-0033123201010303-1032300122232203-0301333310233012-0102331232312221-0331010221101021)
- [custom_network_config.no_global_network](resources--securemesh_site--reference--group-003.md#canonical-1321020212311300-1111231333313020-3332223223121213-3103332312322312-1011030011231203-3212022023111201-3230122000201132-0311311333310132)
- [custom_network_config.no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-1001313103113102-1000033221100312-2111311202310323-2312033332301000-0123101211223010-1023230323131023-2220113220032212-3113332130103210)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-2131101313011023-3202103200013000-3323310203031110-2210320233112330-0202233231312011-1312202132031201-0231331001012033-3000122001001103)
- [custom_network_config.sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-2201302222100233-1133301323301033-1023210103301301-1330322123030131-2202111223230332-1121203213021131-1100102123030200-2000030033123120)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2002031330322012-1012233302212112-3013233021321320-0300100301002202-1321312102002212-0233302121100230-0001020100102311-1110011323111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233110311323110-3122232123222012-1001233101023023-2022013210012211-1233033031013120-1312333032022213-2113100121212101-3220221203012013"></a>

## custom_network_config.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 102120021032 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.active_enhanced_firewall_policies

<a id="canonical-0021032211102020-2131030200222300-2020022311303020-2203101233312212-3110131210011221-1323320112220321-3211130221012013-2230212111031120"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333010221131222-2322221002030111-3233130203121030-3100203000103302-1231100021021320-0211320210102320-1313001103310021-1113321213313130"></a>

## Direct properties — active_enhanced_firewall_policies / 102120021032 / 3

- [enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-0123112003322103-1132131032000203-0233002213030133-3003223030111001-0023302010200331-1232113021202010-3232312101113333-3131331222223120): complete subsection reference.

<a id="canonical-0313032113303023-2220013110013120-0002131103222332-3103021032010102-0312030221212203-0101101332110132-0222202112130101-0310102012123201"></a>

## Next pages — active_enhanced_firewall_policies / 102120021032 / 4

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-0123112003322103-1132131032000203-0233002213030133-3003223030111001-0023302010200331-1232113021202010-3232312101113333-3131331222223120)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0123112003322103-1132131032000203-0233002213030133-3003223030111001-0023302010200331-1232113021202010-3232312101113333-3131331222223120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001222203021200-2313222313021230-2212212132312033-3023320313002021-3110310310103032-3101231132201131-3331203133030000-3320310002303032"></a>

## custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 210213031222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-2002031330322012-1012233302212112-3013233021321320-0300100301002202-1321312102002212-0233302121100230-0001020100102311-1110011323111002)
- custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0001313102020122-0023222332021222-2003202332110011-3111331030223330-3330213200110222-3003200233021200-0221310131032112-1330130222321123"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010130230033213-0222013111312023-0233220333221233-1202233121003200-0103331203311123-2000232211011011-2022310201210130-3012232000101313"></a>

## Direct properties — enhanced_firewall_policies / 210213031222 / 3

<a id="canonical-3330031332202000-2113030031231101-0003021112233121-2331113233213023-0012001003203223-2223122111233031-0103323200133203-2223003101103102"></a>

<a id="canonical-0222023103223013-0201002312022130-0303322231011021-2110211310233211-1000210113300222-2301303112000302-0022033021221022-2023333101320111"></a>

## name property — enhanced_firewall_policies / 210213031222 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1012222121232311-1303303010332003-3021133202331130-2221301011033233-2112220300221201-1120320201231112-2333033123300232-2100222111122011"></a>

<a id="canonical-2321301211221223-2002022300230210-1220331312310132-1312321231032233-3122031111001132-0033113222110130-3331312311020132-0023102211010011"></a>

## namespace property — enhanced_firewall_policies / 210213031222 / 5

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2030200023320111-2301020011221220-1110232212211103-1130333102001030-3113031310321223-1310321122000323-2101232101330220-2113030022321220"></a>

<a id="canonical-0310320000131220-2111022210021323-0302321200000101-1121130013101321-3120231023002111-3122301013031003-1310203300022221-1112232330121222"></a>

## tenant property — enhanced_firewall_policies / 210213031222 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2130113221330210-3132123301323012-1213211323312311-0301201103233031-3310203113201113-2111132131032000-2220012013120113-1021301121022322"></a>

## Next pages — enhanced_firewall_policies / 210213031222 / 7

- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-2002031330322012-1012233302212112-3013233021321320-0300100301002202-1321312102002212-0233302121100230-0001020100102311-1110011323111002)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2201321013032120-0211131001321111-3110020111200102-1101220033300100-1300302002133120-0011321112312011-3010122313323122-0031322212023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113110322023232-0233202032231212-3212223310311013-3332123232233001-3011200303001123-1110113203221330-3030212213310111-2232101310333121"></a>

## custom_network_config.active_forward_proxy_policies — active_forward_proxy_policies / 122213001022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.active_forward_proxy_policies

<a id="canonical-2003211123303112-0103300310123022-1232130213230100-0312310200332100-0030033332012200-2202013010302011-0323201110002202-0101000020123103"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233322210311100-1003013110101333-0010031221223030-0110022302102022-3133322003112100-3100122123001020-2031121201121213-1302222110332230"></a>

## Direct properties — active_forward_proxy_policies / 122213001022 / 3

- [forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-0031031222202300-0131321030231232-0301112023223302-1233012200322002-0112111210101010-2311200331110121-0001101110320311-3221000012233202): complete subsection reference.

<a id="canonical-0021132321120303-2021212113322031-2200100333113001-2010022102102133-3212202122301200-0233010230203311-2221230311000113-1113320020120200"></a>

## Next pages — active_forward_proxy_policies / 122213001022 / 4

- [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-0031031222202300-0131321030231232-0301112023223302-1233012200322002-0112111210101010-2311200331110121-0001101110320311-3221000012233202)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0031031222202300-0131321030231232-0301112023223302-1233012200322002-0112111210101010-2311200331110121-0001101110320311-3221000012233202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133322330102312-3123330323311120-0300002321013233-0311121312213132-2033330210302302-0220223132102320-2221212300301230-2301330101121221"></a>

## custom_network_config.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 110201313013 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-2201321013032120-0211131001321111-3110020111200102-1101220033300100-1300302002133120-0011321112312011-3010122313323122-0031322212023203)
- custom_network_config.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3113122231312030-1201021123121232-2302013330011002-2313112202121031-0231023022321111-1100330330130103-1310300032101111-3310032213002000"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310022021301212-0302130222213200-0231112003203033-3030112303300210-2103221311011031-3312223332120002-1232232311201132-0333111321110201"></a>

## Direct properties — forward_proxy_policies / 110201313013 / 3

<a id="canonical-0230030102013123-1202212132220101-1221130312222331-0020132312200112-2311111001112200-0322003220122033-1302302302221110-2000312122110323"></a>

<a id="canonical-0313323300201321-0231131201032320-1322202232222021-2201022002333023-2000231000112120-3323023120333302-3130020233121002-1302003032112032"></a>

## name property — forward_proxy_policies / 110201313013 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2231020310002301-0231122122210110-0232030303321203-2223100200300023-1012102003031122-1231001122002202-2221102003211032-1310111013010002"></a>

<a id="canonical-3012111100323012-3031122310220311-1122223033201101-2211003201100122-0022120311201021-1333112302023030-3322310313302113-2022212120102010"></a>

## namespace property — forward_proxy_policies / 110201313013 / 5

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2303123212221233-3012220231323300-0131321333131321-0010223312113130-0030210301203232-2221203103030101-2003210013301212-1121103103213130"></a>

<a id="canonical-1111012233211112-2313013331330202-1130323100330303-2310020021103013-3232231210101233-2331021110132000-0333311130112301-3203030113020310"></a>

## tenant property — forward_proxy_policies / 110201313013 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0322030132302300-3220003012310310-1330233022331321-1011110001113211-0232313312310112-3001231133022102-1131113120302103-1102202300323322"></a>

## Next pages — forward_proxy_policies / 110201313013 / 7

- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-2201321013032120-0211131001321111-3110020111200102-1101220033300100-1300302002133120-0011321112312011-3010122313323122-0031322212023203)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3023131331003221-2030010320001301-1121003103323033-2220220101231222-3320232031203110-3232001230223213-1002021211312311-3003233011102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320222110312303-0213233302031011-0033312223101023-0100331100200110-1312003231121003-2301331233033003-3211210302121033-0321110031321021"></a>

## custom_network_config.active_network_policies — active_network_policies / 233000221010 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.active_network_policies

<a id="canonical-2113313301002330-1121203111303013-0133313302021112-0110023221130023-2231202231230210-2121212331221122-2320321132302111-0230302321230112"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220203113321121-0002002001112312-0000310103301100-1133031233302232-3110233313132101-2023311100210111-0023210100012200-3001321311323132"></a>

## Direct properties — active_network_policies / 233000221010 / 3

- [network_policies](resources--securemesh_site--reference--group-001.md#canonical-2201212331223000-2100033031022020-2231212003033312-2313213032302001-0023313132212333-2310231211313130-1222303030102221-3020132113210302): complete subsection reference.

<a id="canonical-0210311311312222-2331220132112213-0322100020013320-3000333321202331-3320120000322221-1123000101201333-0302001022221223-1222012100112130"></a>

## Next pages — active_network_policies / 233000221010 / 4

- [custom_network_config.active_network_policies.network_policies](resources--securemesh_site--reference--group-001.md#canonical-2201212331223000-2100033031022020-2231212003033312-2313213032302001-0023313132212333-2310231211313130-1222303030102221-3020132113210302)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2201212331223000-2100033031022020-2231212003033312-2313213032302001-0023313132212333-2310231211313130-1222303030102221-3020132113210302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003132321223203-0001011033013320-1331130123010100-0310112302212202-3230013132222303-2203032301113111-0231223331123302-0120033033202000"></a>

## custom_network_config.active_network_policies.network_policies — network_policies / 012303233312 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-3023131331003221-2030010320001301-1121003103323033-2220220101231222-3320232031203110-3232001230223213-1002021211312311-3003233011102220)
- custom_network_config.active_network_policies.network_policies

<a id="canonical-1012030130033103-1113321221300112-3001312323123133-3331031201312333-0002210332002031-2131020333220223-3321001203321003-1033302003132012"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333202223300220-0222101301032332-0322201223130122-1113132201333100-3000100211323010-1130300102130210-3113300101010310-3021133202013001"></a>

## Direct properties — network_policies / 012303233312 / 3

<a id="canonical-3021003330011311-1333303212001111-2212311033210032-0322000022020120-2210331323310132-2022230310111201-3030002101031112-3133231022132211"></a>
